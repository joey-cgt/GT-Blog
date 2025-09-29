-- 创建文章表
CREATE TABLE IF NOT EXISTS articles (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL COMMENT '文章标题',
    abstract TEXT COMMENT '文章摘要',
    content LONGTEXT NOT NULL COMMENT '文章内容',
    column_id INT DEFAULT 0 COMMENT '专栏ID',
    category_id INT NOT NULL COMMENT '分类ID',
    tag_ids JSON COMMENT '标签ID数组',
    cover_url VARCHAR(500) COMMENT '封面图片URL',
    status TINYINT DEFAULT 0 COMMENT '文章状态：0-草稿，1-已发布',
    is_top BOOLEAN DEFAULT FALSE COMMENT '是否置顶',
    view_count INT DEFAULT 0 COMMENT '浏览量',
    like_count INT DEFAULT 0 COMMENT '点赞数',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    publish_time DATETIME NULL COMMENT '发布时间',
    update_time DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    version INT DEFAULT 0 COMMENT '版本号（用于乐观锁）',
    INDEX idx_status (status),
    INDEX idx_category (category_id),
    INDEX idx_column (column_id),
    INDEX idx_top (is_top),
    INDEX idx_publish_time (publish_time),
    INDEX idx_create_time (create_time)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文章表';

-- 创建分类表
CREATE TABLE IF NOT EXISTS categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL COMMENT '分类名称',
    description VARCHAR(500) COMMENT '分类描述',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    update_time DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY uk_name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文章分类表';

-- 创建专栏表
CREATE TABLE IF NOT EXISTS columns (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL COMMENT '专栏名称',
    description VARCHAR(500) COMMENT '专栏描述',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    update_time DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY uk_name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='专栏表';

-- 创建标签表
CREATE TABLE IF NOT EXISTS tags (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL COMMENT '标签名称',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    update_time DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY uk_name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='标签表';

-- 插入默认分类数据
INSERT IGNORE INTO categories (id, name, description) VALUES
(1, '前端开发', '前端技术相关文章'),
(2, '后端开发', '后端技术相关文章'),
(3, 'UI框架', 'UI框架相关文章'),
(4, '性能优化', '性能优化相关文章'),
(5, 'CSS', 'CSS技术相关文章'),
(6, 'JavaScript', 'JavaScript技术相关文章');

-- 插入默认专栏数据
INSERT IGNORE INTO columns (id, name, description) VALUES
(1, 'Vue.js 精讲', 'Vue.js 框架深入讲解'),
(2, 'React 实战', 'React 框架实战经验'),
(3, 'TypeScript 进阶', 'TypeScript 高级特性'),
(4, '前端性能优化', '前端性能优化技巧'),
(5, 'CSS 魔法', 'CSS 高级技巧'),
(6, 'Node.js 开发', 'Node.js 后端开发');

-- 插入默认标签数据
INSERT IGNORE INTO tags (id, name) VALUES
(1, 'Vue'),
(2, 'React'),
(3, 'TypeScript'),
(4, 'JavaScript'),
(5, 'CSS'),
(6, 'HTML'),
(7, 'Node.js'),
(8, 'Webpack');