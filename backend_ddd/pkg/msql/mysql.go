package mysql

import (
	"context"
	"fmt"
	"gt-blog/backend_ddd/pkg/config"
	"gt-blog/backend_ddd/pkg/log"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Init 初始化MySQL连接，返回gorm.DB实例
// 接收MySQL配置，返回数据库连接和错误
func Init(cfg config.MySQLConfig) (*gorm.DB, error) {
	// 1. 构建DSN（数据源名称）
	dsn := buildDSN(cfg)
	log.Info("MySQL DSN构建完成", log.String("dsn", maskPassword(dsn))) // 日志隐藏密码

	// 2. 配置GORM选项
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 日志级别：Info（可根据环境调整）
	}

	// 3. 打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("打开MySQL连接失败: %w", err)
	}

	// 4. 配置连接池（关键：避免高并发下的连接耗尽）
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("获取底层SQL连接失败: %w", err)
	}

	// 设置连接池参数（从配置读取，确保灵活性）
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)                                    // 最大打开连接数
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)                                    // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second) // 连接最大存活时间

	// 5. 测试连接（验证配置有效性）
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("MySQL连接测试失败: %w", err)
	}

	log.Info("MySQL连接初始化成功",
		log.Int("max_open_conns", cfg.MaxOpenConns),
		log.Int("max_idle_conns", cfg.MaxIdleConns),
	)
	return db, nil
}

// Close 关闭MySQL连接（释放资源）
func Close(db *gorm.DB) error {
	if db == nil {
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// 构建DSN字符串（根据配置拼接）
func buildDSN(cfg config.MySQLConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname,
		cfg.Charset,
		cfg.ParseTime,
	)
}

// 日志中隐藏密码（安全考虑）
func maskPassword(dsn string) string {
	// 简单实现：替换密码部分为***
	// 实际可根据需要完善正则匹配
	parts := strings.Split(dsn, ":")
	if len(parts) >= 2 {
		parts[1] = "***" // 密码在第二部分（user:password@...）
		return strings.Join(parts, ":")
	}
	return dsn
}
