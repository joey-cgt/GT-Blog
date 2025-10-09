package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config 系统总配置结构体
type Config struct {
	Server ServerConfig `yaml:"server"` // 服务器配置
	MySQL  MySQLConfig  `yaml:"mysql"`  // MySQL数据库配置
	Redis  RedisConfig  `yaml:"redis"`  // Redis缓存配置
	Log    LogConfig    `yaml:"log"`    // 日志配置
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port         string `yaml:"port"`          // 监听端口，如"8080"
	Mode         string `yaml:"mode"`          // 运行模式：debug/release
	ReadTimeout  int    `yaml:"read_timeout"`  // HTTP读超时（秒）
	WriteTimeout int    `yaml:"write_timeout"` // HTTP写超时（秒）
}

// MySQLConfig 数据库配置
type MySQLConfig struct {
	Host            string `yaml:"host"`              // 数据库主机
	Port            int    `yaml:"port"`              // 数据库端口
	User            string `yaml:"user"`              // 用户名
	Password        string `yaml:"password"`          // 密码
	Dbname          string `yaml:"dbname"`            // 数据库名
	Charset         string `yaml:"charset"`           // 字符集
	ParseTime       bool   `yaml:"parseTime"`         // 是否解析时间
	MaxOpenConns    int    `yaml:"max_open_conns"`    // 最大打开连接数
	MaxIdleConns    int    `yaml:"max_idle_conns"`    // 最大空闲连接数
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"` // 连接最大存活时间（秒）
}

// RedisConfig 缓存配置
type RedisConfig struct {
	Enable       bool   `yaml:"enable"`         // 是否启用Redis
	Addr         string `yaml:"addr"`           // 地址：host:port
	Password     string `yaml:"password"`       // 密码
	DB           int    `yaml:"db"`             // 数据库编号
	PoolSize     int    `yaml:"pool_size"`      // 连接池大小
	MinIdleConns int    `yaml:"min_idle_conns"` // 最小空闲连接数
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `yaml:"level"`       // 日志级别：debug/info/warn/error
	OutputPath string `yaml:"output_path"` // 输出路径：stdout 或文件路径
	MaxSize    int    `yaml:"max_size"`    // 单个日志文件大小（MB）
	MaxBackups int    `yaml:"max_backups"` // 最大备份文件数
	MaxAge     int    `yaml:"max_age"`     // 日志保留天数
	Compress   bool   `yaml:"compress"`    // 是否压缩备份文件
}

// Load 从指定路径加载YAML配置文件
func Load(path string) (*Config, error) {
	// 解析绝对路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	// 读取文件内容
	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	// 解析YAML
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	// 设置默认值（避免配置文件缺失必填项）
	setDefaults(&cfg)

	return &cfg, nil
}

// setDefaults 为缺失的配置项设置默认值
func setDefaults(cfg *Config) {
	// 服务器默认配置
	if cfg.Server.Port == "" {
		cfg.Server.Port = "8080"
	}
	if cfg.Server.Mode == "" {
		cfg.Server.Mode = "debug"
	}
	if cfg.Server.ReadTimeout <= 0 {
		cfg.Server.ReadTimeout = 30
	}
	if cfg.Server.WriteTimeout <= 0 {
		cfg.Server.WriteTimeout = 30
	}

	// MySQL默认配置
	if cfg.MySQL.MaxOpenConns <= 0 {
		cfg.MySQL.MaxOpenConns = 100
	}
	if cfg.MySQL.MaxIdleConns <= 0 {
		cfg.MySQL.MaxIdleConns = 20
	}
	if cfg.MySQL.ConnMaxLifetime <= 0 {
		cfg.MySQL.ConnMaxLifetime = 3600
	}

	// Redis默认配置
	if cfg.Redis.PoolSize <= 0 {
		cfg.Redis.PoolSize = 10
	}
	if cfg.Redis.MinIdleConns <= 0 {
		cfg.Redis.MinIdleConns = 2
	}

	// 日志默认配置
	if cfg.Log.Level == "" {
		cfg.Log.Level = "info"
	}
	if cfg.Log.OutputPath == "" {
		cfg.Log.OutputPath = "stdout" // 默认输出到控制台
	}
	if cfg.Log.MaxSize <= 0 {
		cfg.Log.MaxSize = 100
	}
	if cfg.Log.MaxBackups <= 0 {
		cfg.Log.MaxBackups = 10
	}
	if cfg.Log.MaxAge <= 0 {
		cfg.Log.MaxAge = 30
	}
}
