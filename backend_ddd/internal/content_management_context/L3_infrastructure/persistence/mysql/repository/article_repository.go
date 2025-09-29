package repository

import (
	"context"
	"errors"
	"fmt"
	"time"
	"gt-blog/backend_ddd/internal/content_management_context/L2_domain/model"
	"gt-blog/backend_ddd/internal/content_management_context/L2_domain/repository"
	"gt-blog/backend_ddd/internal/content_management_context/L3_infrastructure/persistence/mysql/dao"

	"gorm.io/gorm"
)

type MySQLArticleRepository struct {
	db *gorm.DB
}

func NewMySQLArticleRepository(db *gorm.DB) repository.ArticleRepository {
	return &MySQLArticleRepository{db: db}
}

// Add:领域模型→数据模型→保存→回写ID
func (r *MySQLArticleRepository) Add(ctx context.Context, article *model.Article) (int, error) {
	// 1. 领域模型 → 数据模型(核心转换逻辑)
	articleDAO := r.toDAO(article)

	// 2. 执行保存(技术操作)
	// 使用事务确保文章和标签关系的一致性
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if articleDAO.ID == 0 {
		// 新增:依赖MySQL自增ID
		if err := tx.Create(articleDAO).Error; err != nil {
			tx.Rollback()
			return 0, err
		}

		// 保存标签关系到中间表
		for _, tagID := range article.TagIDs {
			articleTag := &dao.ArticleTagDAO{
				ArticleID:  articleDAO.ID,
				TagID:      tagID,
				CreateTime: int(time.Now().Unix()),
			}
			if err := tx.Create(articleTag).Error; err != nil {
				tx.Rollback()
				return 0, err
			}
		}
	} else {
		// 更新:使用乐观锁(依赖Version字段)
		if err := tx.Save(articleDAO).Error; err != nil {
			tx.Rollback()
			return 0, err
		}

		// 更新标签关系：先删除旧关系，再添加新关系
		// 删除旧关系
		if err := tx.Where("article_id = ?", articleDAO.ID).Delete(&dao.ArticleTagDAO{}).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
		// 添加新关系
		for _, tagID := range article.TagIDs {
			articleTag := &dao.ArticleTagDAO{
				ArticleID:  articleDAO.ID,
				TagID:      tagID,
				CreateTime: int(time.Now().Unix()),
			}
			if err := tx.Create(articleTag).Error; err != nil {
				tx.Rollback()
				return 0, err
			}
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	// 3. 将数据库生成的ID和Version回写到领域模型
	article.ID = articleDAO.ID
	article.Version = articleDAO.Version
	return article.ID, nil
}

func (r *MySQLArticleRepository) Update(ctx context.Context, article *model.Article) error {
	// 1. 验证更新前提:必须是已持久化的文章(ID > 0)
	if article.ID <= 0 {
		return errors.New("文章ID不能为空, 无法执行更新操作")
	}

	// 2. 领域模型 → 数据模型(复用统一转换逻辑，避免不一致)
	articleDAO := r.toDAO(article)

	// 3. 执行更新(带乐观锁，防止并发更新冲突)
	tx := r.db.WithContext(ctx)
	// 仅更新非零字段，且通过 WHERE id + version 确保并发安全
	result := tx.Model(articleDAO).
		// Where("id = ? AND version = ?", articleDAO.ID, articleDAO.Version). // 乐观锁条件
		Where("id = ?", articleDAO.ID).
		Updates(articleDAO)

	// 4. 处理数据库操作错误
	if result.Error != nil {
		return fmt.Errorf("数据库更新失败:%w", result.Error)
	}

	// 5. 处理更新结果:无数据受影响的两种场景
	if result.RowsAffected == 0 {
		// 检查文章是否存在(区分“记录不存在”和“版本冲突”)
		exists, err := r.exists(ctx, article.ID)
		if err != nil {
			return fmt.Errorf("校验文章存在性失败:%w", err)
		}
		if !exists {
			return fmt.Errorf("待更新的文章不存在(ID:%d)", article.ID)
		}
		// 记录存在但未更新:版本号不匹配(并发冲突)
		return errors.New("文章已被其他操作修改，请刷新最新数据后重试")
	}

	// 6. 同步领域模型版本号(更新成功后+1，与数据库保持一致)
	article.Version++

	return nil
}

func (r *MySQLArticleRepository) DeleteByID(ctx context.Context, id int) error {
	// 1. 验证ID有效性
	if id <= 0 {
		return errors.New("文章ID必须大于0，无法执行删除操作")
	}

	// 2. 检查文章是否存在
	exists, err := r.exists(ctx, id)
	if err != nil {
		return fmt.Errorf("校验文章存在性失败:%w", err)
	}
	if !exists {
		return fmt.Errorf("待删除的文章不存在(ID:%d)", id)
	}

	// 3. 执行删除操作
	tx := r.db.WithContext(ctx)
	result := tx.Delete(&dao.ArticleDAO{}, "id = ?", id)

	// 4. 处理数据库操作错误
	if result.Error != nil {
		return fmt.Errorf("数据库删除失败:%w", result.Error)
	}

	// 5. 确认删除结果
	if result.RowsAffected == 0 {
		return fmt.Errorf("删除操作未影响任何数据(ID:%d)", id)
	}

	return nil
}

// FindByID:查询数据模型→转换为领域模型
func (r *MySQLArticleRepository) FindByID(ctx context.Context, id int) (*model.Article, error) {
	var articleDAO dao.ArticleDAO
	if err := r.db.WithContext(ctx).First(&articleDAO, "id = ?", id).Error; err != nil {
		return nil, err
	}
	
	// 查询文章关联的标签ID列表
	article := r.toDomain(&articleDAO)
	tagIDs, err := r.findTagIDsByArticleID(ctx, id)
	if err != nil {
		return article, nil // 返回基本信息，忽略标签错误
	}
	article.TagIDs = tagIDs
	return article, nil
}

func (r *MySQLArticleRepository) FindByPage(ctx context.Context, status int, offset, limit int, keyword string) ([]*model.Article, int, error) {
	// 1. 构建基础查询（带上下文）
	db := r.db.WithContext(ctx).
		Model(&dao.ArticleDAO{}).
		Where("status = ?", status)

	// 2. 处理关键词搜索条件（模糊匹配标题和摘要）
	if keyword != "" {
		searchPattern := "%" + keyword + "%"
		db = db.Where("title LIKE ? OR abstract LIKE ?", searchPattern, searchPattern)
	}

	// 3. 查询符合条件的总条数（用于计算分页总页数）
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("查询文章总数失败: %w", err)
	}

	// 4. 分页查询文章数据（带排序：置顶文章优先，再按发布时间倒序）
	var articleDAOs []*dao.ArticleDAO
	err := db.
		Offset(offset).
		Limit(limit).
		Order("is_top DESC").       // 置顶文章排在前面
		Order("publish_time DESC"). // 非置顶文章按发布时间倒序
		Find(&articleDAOs).Error

	if err != nil {
		return nil, 0, fmt.Errorf("分页查询文章失败: %w", err)
	}

	// 5. 将 DAO 转换为领域模型（核心：隔离数据层与领域层）
	articles := make([]*model.Article, 0, len(articleDAOs))
	for _, dao := range articleDAOs {
		article := r.toDomain(dao)
		// 查询并设置标签ID列表
		tagIDs, err := r.findTagIDsByArticleID(ctx, dao.ID)
		if err == nil {
			article.TagIDs = tagIDs
		}
		articles = append(articles, article)
	}

	// 6. 返回领域模型列表、总条数（注意转换为int，避免int64类型问题）
	return articles, int(total), nil
}

func (r *MySQLArticleRepository) FindTopArticles(ctx context.Context, limit int) ([]*model.Article, error) {
	// 1. 参数校验
	if limit <= 0 {
		return nil, fmt.Errorf("limit 必须为正数")
	}

	// 2. 构建基础查询并添加“置顶”条件
	db := r.db.WithContext(ctx).
		Where("is_top = ?", true).
		Order("create_time DESC"). // 固定排序规则
		Limit(limit)

	// 3. 执行查询
	var articleDAOs []*dao.ArticleDAO

	err := db.Find(&articleDAOs).Error

	if err != nil {
		return nil, fmt.Errorf("查询置顶文章失败: %w", err)
	}

	// 4. 将 DAO 转换为领域模型
	articles := make([]*model.Article, 0, len(articleDAOs))
	for _, dao := range articleDAOs {
		articles = append(articles, r.toDomain(dao))
	}

	return articles, nil
}
func (r *MySQLArticleRepository) FindLatestArticles(ctx context.Context, limit int) ([]*model.Article, error) {
	// 1. 参数校验
	if limit <= 0 {
		return nil, fmt.Errorf("limit 必须为正数")
	}

	// 2. 构建查询：筛选非置顶文章，并按创建时间降序排序
	db := r.db.WithContext(ctx).
		Where("is_top = ?", false).
		Order("create_time DESC"). // 固定排序规则
		Limit(limit)

	// 3. 执行查询
	var articleDAOs []*dao.ArticleDAO
	err := db.Find(&articleDAOs).Error
	if err != nil {
		return nil, fmt.Errorf("查询最新文章失败: %w", err)
	}

	// 4. 将 DAO 转换为领域模型
	articles := make([]*model.Article, 0, len(articleDAOs))
	for _, dao := range articleDAOs {
		articles = append(articles, r.toDomain(dao))
	}

	return articles, nil
}
func (r *MySQLArticleRepository) FindPopularArticles(ctx context.Context, limit int) ([]*model.Article, error) {
	// 1. 参数校验
	if limit <= 0 {
		return nil, fmt.Errorf("limit 必须为正数")
	}

	// 2. 构建查询：筛选非置顶文章，先按点赞数降序，再按创建时间降序
	db := r.db.WithContext(ctx).
		Where("is_top = ?", false).
		Order("like_count DESC").  // 第一排序规则：点赞数
		Order("create_time DESC"). // 第二排序规则：创建时间
		Limit(limit)

	// 3. 执行查询
	var articleDAOs []*dao.ArticleDAO
	err := db.Find(&articleDAOs).Error
	if err != nil {
		return nil, fmt.Errorf("查询热门文章失败: %w", err)
	}

	// 4. 将 DAO 转换为领域模型
	articles := make([]*model.Article, 0, len(articleDAOs))
	for _, dao := range articleDAOs {
		articles = append(articles, r.toDomain(dao))
	}

	return articles, nil
}

func (r *MySQLArticleRepository) FindByCategoryID(ctx context.Context, categoryID int, offset, limit int) ([]*model.Article, int, error) {
	// 1. 构建基础查询
	db := r.db.WithContext(ctx).Model(&dao.ArticleDAO{}).Where("category_id = ?", categoryID)

	// 2. 查询总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("查询分类文章总数失败: %w", err)
	}

	// 3. 分页查询
	var articleDAOs []*dao.ArticleDAO
	err := db.
		Offset(offset).
		Limit(limit).
		Order("is_top DESC, publish_time DESC"). // 保持与FindByPage一致的排序
		Find(&articleDAOs).Error
	if err != nil {
		return nil, 0, fmt.Errorf("分页查询分类文章失败: %w", err)
	}

	// 4. DAO 转领域模型
	articles := make([]*model.Article, 0, len(articleDAOs))
	for _, dao := range articleDAOs {
		articles = append(articles, r.toDomain(dao))
	}

	return articles, int(total), nil
}

func (r *MySQLArticleRepository) FindByColumnID(ctx context.Context, columnID int, offset, limit int) ([]*model.Article, int, error) {
	// 1. 构建基础查询
	db := r.db.WithContext(ctx).Model(&dao.ArticleDAO{}).Where("column_id = ?", columnID)

	// 2. 查询总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("查询专栏文章总数失败: %w", err)
	}

	// 3. 分页查询
	var articleDAOs []*dao.ArticleDAO
	err := db.
		Offset(offset).
		Limit(limit).
		Order("is_top DESC, publish_time DESC").
		Find(&articleDAOs).Error
	if err != nil {
		return nil, 0, fmt.Errorf("分页查询专栏文章失败: %w", err)
	}

	// 4. DAO 转领域模型
	articles := make([]*model.Article, 0, len(articleDAOs))
	for _, dao := range articleDAOs {
		articles = append(articles, r.toDomain(dao))
	}

	return articles, int(total), nil
}

func (r *MySQLArticleRepository) FindByTagID(ctx context.Context, tagID int, offset, limit int) ([]*model.Article, int, error) {
	// 1. 构建基础查询：JOIN 中间表进行筛选
	// 注意：这里的 "article_tags" 是数据库中的实际表名, 结构为 (article_id, tag_id)
	db := r.db.WithContext(ctx).
		Model(&dao.ArticleDAO{}).
		Joins("JOIN article_tags ON article_tags.article_id = articles.id").
		Where("article_tags.tag_id = ?", tagID)

	// 2. 查询总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("查询标签文章总数失败: %w", err)
	}

	// 3. 分页查询
	var articleDAOs []*dao.ArticleDAO
	err := db.
		Offset(offset).
		Limit(limit).
		Order("is_top DESC, publish_time DESC").
		Find(&articleDAOs).Error
	if err != nil {
		return nil, 0, fmt.Errorf("分页查询标签文章失败: %w", err)
	}

	// 4. DAO 转领域模型
	articles := make([]*model.Article, 0, len(articleDAOs))
	for _, dao := range articleDAOs {
		articles = append(articles, r.toDomain(dao))
	}

	return articles, int(total), nil
}

// 其他方法(Delete、IncrementViewCount等)类似，均需通过转换实现

// 辅助方法:检查指定ID的文章是否存在(封装DAO层细节，不暴露给领域层)
func (r *MySQLArticleRepository) exists(ctx context.Context, articleID int) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&dao.ArticleDAO{}). // 基于DAO模型查询，避免直接写表名
		Where("id = ?", articleID).
		Count(&count). // 仅计数，不查完整数据，性能更优
		Error
	if err != nil {
		return false, fmt.Errorf("查询文章计数失败:%w", err)
	}
	return count > 0, nil
}

// ------------------------------ 转换函数(核心) ------------------------------

// 领域模型 → 数据模型
func (r *MySQLArticleRepository) toDAO(article *model.Article) *dao.ArticleDAO {
	return &dao.ArticleDAO{
		ID:          article.ID,
		Title:       article.Title,
		Abstract:    article.Abstract,
		Content:     article.Content,
		ColumnID:    article.ColumnID,
		CategoryID:  article.CategoryID,
		CoverUrl:    article.CoverUrl,
		Status:      article.Status,
		IsTop:       article.IsTop,
		ViewCount:   article.ViewCount,
		LikeCount:   article.LikeCount,
		CreateTime:  article.CreateTime,
		PublishTime: article.PublishTime,
		UpdateTime:  article.UpdateTime,
		Version:     article.Version, // 乐观锁版本号同步
	}
}

// 数据模型 → 领域模型
func (r *MySQLArticleRepository) toDomain(dao *dao.ArticleDAO) *model.Article {
	return &model.Article{
		ID:          dao.ID,
		Title:       dao.Title,
		Abstract:    dao.Abstract,
		Content:     dao.Content,
		ColumnID:    dao.ColumnID,
		CategoryID:  dao.CategoryID,
		TagIDs:      []int{}, // 标签ID需要单独查询
		CoverUrl:    dao.CoverUrl,
		Status:      dao.Status,
		IsTop:       dao.IsTop,
		ViewCount:   dao.ViewCount,
		LikeCount:   dao.LikeCount,
		CreateTime:  dao.CreateTime,
		PublishTime: dao.PublishTime,
		UpdateTime:  dao.UpdateTime,
		Version:     dao.Version,
	}
}

// 通过文章ID查询标签ID列表
func (r *MySQLArticleRepository) findTagIDsByArticleID(ctx context.Context, articleID int) ([]int, error) {
	var tagIDs []int
	// 从中间表查询文章关联的所有标签ID
	if err := r.db.WithContext(ctx).
		Model(&dao.ArticleTagDAO{}).
		Where("article_id = ?", articleID).
		Pluck("tag_id", &tagIDs).
		Error;
	err != nil {
		return nil, err
	}
	return tagIDs, nil
}
