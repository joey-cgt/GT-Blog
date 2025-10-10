package main

import (
	"context"
	errors "errors"
	"flag"
	"gt-blog/backend/internal/ddd/content_management_context/L0_interface/handler"
	"gt-blog/backend/internal/ddd/content_management_context/L0_interface/router"
	appservice "gt-blog/backend/internal/ddd/content_management_context/L1_application/service"
	domainservice "gt-blog/backend/internal/ddd/content_management_context/L2_domain/service"
	dao "gt-blog/backend/internal/ddd/content_management_context/L3_infrastructure/persistence/mysql/dao"
	mysqlrepository "gt-blog/backend/internal/ddd/content_management_context/L3_infrastructure/persistence/mysql/repository"
	"gt-blog/backend/internal/mvc/controller"
	"gt-blog/backend/internal/mvc/model"
	"gt-blog/backend/internal/mvc/repository"
	"gt-blog/backend/internal/mvc/route"
	"gt-blog/backend/internal/mvc/service"
	"gt-blog/backend/pkg/config"
	"gt-blog/backend/pkg/log"
	mysql "gt-blog/backend/pkg/msql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var configPath = flag.String("config", "pkg/config/config.yaml", "配置文件路径")

func main() {
	// 解析命令行参数（如指定配置文件路径）
	flag.Parse()

	// 加载配置
	cfg, err := config.Load(*configPath)
	if err != nil {
		// 配置加载失败，直接退出（此时日志可能未初始化，用标准输出）
		panic("加载配置失败: " + err.Error())
	}

	// 初始化日志（依赖配置中的日志设置）
	logger := log.Init(cfg.Log)
	defer logger.Sync() // 程序退出时刷新日志缓存
	logger.Info("日志初始化完成", log.Any("config", cfg.Log))

	// 初始化数据库连接（MySQL）
	db, err := mysql.Init(cfg.MySQL)
	if err != nil {
		logger.Fatal("数据库初始化失败", log.String("error", err.Error()))
	}
	defer mysql.Close(db) // 程序退出时关闭数据库连接
	logger.Info("数据库连接成功")

	// 自动迁移数据库表结构
	daoModels := []interface{}{
		&dao.ArticleDAO{},
		&dao.ArticleTagDAO{},
		&dao.CategoryDAO{},
		&dao.ColumnDAO{},
		&dao.TagDAO{},
		// 统计相关模型
		&dao.DailyBlogStatsDAO{},
		&dao.DailyArticleStatsDAO{},
		&dao.ArticleViewIncrementDAO{},

		// 添加admin相关模型
		&model.Admin{},
		&model.SocialAccount{},
	}
	if err := db.AutoMigrate(daoModels...); err != nil {
		logger.Fatal("数据库表迁移失败", log.String("error", err.Error()))
	}
	logger.Info("数据库表结构初始化完成")

	// 初始化admin表，添加默认管理员账号
	initAdminTable(db, logger)

	// 初始化仓储层（注入数据库/缓存依赖）
	articleRepo := mysqlrepository.NewMySQLArticleRepository(db)
	tagRepo := mysqlrepository.NewMySQLTagRepository(db)
	categoryRepo := mysqlrepository.NewMySQLCategoryRepository(db)
	articleTagRepo := mysqlrepository.NewMySQLArticleTagRepository(db)
	columnRepo := mysqlrepository.NewMySQLColumnRepository(db)
	adminRepo := repository.NewMySqlAdminRepository(db)
	visitorRepo := repository.NewMySQLVisitorRepository(db)

	// 初始化领域层（注入仓储依赖）
	articleDomainService := domainservice.NewArticleDomainService(
		articleRepo,
		articleTagRepo,
		tagRepo,
		columnRepo,
		categoryRepo,
	)

	// 初始化应用服务
	articleAppService := appservice.NewArticleAppService(
		articleDomainService,
		articleRepo,
		tagRepo,
		columnRepo,
		categoryRepo,
		articleTagRepo,
	)

	articleHandler := handler.NewArticleHandler(articleAppService)

	categoryDomainService := domainservice.NewCategoryDomainService(categoryRepo)
	categoryAppService := appservice.NewCategoryAppService(categoryRepo, categoryDomainService)
	categoryHander := handler.NewCategoryHandler(categoryAppService)

	columnDomainService := domainservice.NewColumnDomainService(columnRepo)
	columnAppService := appservice.NewColumnAppService(columnRepo, columnDomainService)
	columnHandler := handler.NewColumnHandler(columnAppService)

	tagDomainService := domainservice.NewTagDomainService(tagRepo)
	tagAppService := appservice.NewTagAppService(tagRepo, tagDomainService)
	tagHandler := handler.NewTagHandler(tagAppService)

	// 初始化admin相关服务和控制器
	adminService := service.NewAdminService(adminRepo)
	adminController := controller.NewAdminController(adminService)

	visitorService := service.NewVisitorService(visitorRepo)
	visitorController := controller.NewVisitorController(visitorService)

	statisticsRepo := mysqlrepository.NewMySQLStatisticsRepository(db)
	statAppService := appservice.NewBlogStatAppsService(articleRepo, categoryRepo, tagRepo, columnRepo, statisticsRepo)
	statisticsHandler := handler.NewStatisticsHandler(statAppService)

	// 注册路由
	r := gin.Default()

	// 关键：添加 CORS 中间件（放在所有路由之前）
	r.Use(cors.New(cors.Config{
		// 1. 允许的前端源：填写你的前端实际地址（不要用 *  if 前端需要带 Cookie）
		AllowOrigins: []string{"http://localhost:5173"},
		// 2. 允许的 HTTP 方法（包含预检请求的 OPTIONS 方法）
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		// 3. 允许的请求头（如 Content-Type、Authorization 等）
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		// 4. 允许前端读取的响应头（可选）
		ExposeHeaders: []string{"Content-Length"},
		// 5. 是否允许跨域请求带 Cookie（如前端需要登录态，设为 true）
		AllowCredentials: true,
		// 6. 预检请求结果的缓存时间（减少 OPTIONS 请求次数，单位：小时）
		MaxAge: 12 * time.Hour,
	}))

	api := r.Group("/api/v1")
	router.RegisterArticleRoutes(api, articleHandler)
	router.RegisterTagRoutes(api, tagHandler)
	router.RegisterCategoryRoutes(api, categoryHander)
	router.RegisterColumnRoutes(api, columnHandler)
	router.RegisterStatisticsRoutes(api, statisticsHandler)
	// 注册admin路由
	route.RegisterAdminRoutes(api, adminController)
	// 注册visitor路由
	route.RegisterVisitorRoutes(api, visitorController)

	// 启动每日统计任务
	dailyStatsTask := statAppService.CreateDailyStatsTask()
	// 启动HTTP服务（带优雅关闭）
	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: r,
	}

	// 异步启动服务器
	go func() {
		logger.Info("服务器启动中", log.String("port", cfg.Server.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("服务器启动失败", log.String("error", err.Error()))
		}
	}()

	// 12. 监听退出信号，实现优雅关闭
	quit := make(chan os.Signal, 1)
	// 接收 SIGINT（Ctrl+C）和 SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("开始优雅关闭服务器...")

	// 30秒超时关闭上下文
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("服务器关闭超时", log.String("error", err.Error()))
	}

	logger.Info("服务器已正常关闭")

	// 停止每日统计任务
	statAppService.StopDailyStatsTask(dailyStatsTask)
}

// initAdminTable 初始化admin表，添加默认管理员账号
func initAdminTable(db *gorm.DB, logger *zap.Logger) {
	// 创建adminRepository实例
	adminRepo := repository.NewMySqlAdminRepository(db)

	// 检查是否已经存在admin账号
	ctx := context.Background()
	_, err := adminRepo.GetAdminByUsername(ctx, "admin")

	if err == nil {
		// 账号已存在，无需初始化
		logger.Info("管理员账号已存在，跳过初始化")
		return
	}

	// 如果是因为记录不存在而报错，则创建默认管理员账号
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 创建默认管理员
		admin := &model.Admin{
			Username: "admin",
			Password: "666666", // 默认密码
			Nickname: "管理员",
			Email:    "admin@example.com",
			Bio:      "博客管理员",
		}

		// 插入管理员记录
		createErr := db.WithContext(ctx).Create(admin).Error
		if createErr != nil {
			logger.Error("创建默认管理员账号失败", log.String("error", err.Error()))
			return
		}

		logger.Info("默认管理员账号创建成功")
	} else {
		// 其他错误情况
		logger.Error("查询管理员账号时发生错误", log.String("error", err.Error()))
	}
}
