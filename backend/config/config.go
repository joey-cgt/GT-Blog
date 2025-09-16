package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	JWT      JWTConfig      `yaml:"jwt"`
	CORS     CORSConfig     `yaml:"cors"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
	Env  string `yaml:"env"`
}

type DatabaseConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	DBName    string `yaml:"dbname"`
	Charset   string `yaml:"charset"`
	ParseTime bool   `yaml:"parseTime"`
}

type JWTConfig struct {
	Secret      string `yaml:"secret"`
	ExpireHours int    `yaml:"expire_hours"`
}

type CORSConfig struct {
	AllowedOrigins []string `yaml:"allowed_origins"`
	AllowedMethods []string `yaml:"allowed_methods"`
	AllowedHeaders []string `yaml:"allowed_headers"`
}

// LoadConfig 从YAML文件加载配置
func LoadConfig() (*Config, error) {
	// 优先使用环境变量指定的配置文件路径
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config/config.yaml"
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析YAML
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	// 从环境变量覆盖敏感配置（可选）
	overrideFromEnv(&config)

	return &config, nil
}

// overrideFromEnv 从环境变量覆盖敏感配置
func overrideFromEnv(config *Config) {
	if dbPass := os.Getenv("DB_PASSWORD"); dbPass != "" {
		config.Database.Password = dbPass
	}
	if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
		config.JWT.Secret = jwtSecret
	}
}

// DSN 返回MySQL连接字符串
func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.DBName,
		d.Charset,
		d.ParseTime,
	)
}
