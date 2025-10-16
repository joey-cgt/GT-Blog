package util

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gopkg.in/yaml.v3"
)

// JWTConfig 定义JWT配置结构
var jwtConfig struct {
	JWT struct {
		Secret      string `yaml:"secret"`
		ExpireHours int    `yaml:"expire_hours"`
	} `yaml:"jwt"`
}

// init 初始化JWT配置
func init() {
	// 解析配置文件以获取JWT配置
	workDir, err := os.Getwd()
	if err != nil {
		panic("Failed to get work directory: " + err.Error())
	}
	// 拼接配置文件路径（工作目录 + 相对路径）
	configPath := filepath.Join(workDir, "pkg", "config", "config.yaml")
	configData, err := os.ReadFile(configPath)
	if err != nil {
		// 这里可以根据需要决定是否panic或者返回错误
		// 在初始化函数中通常会panic以确保配置正确加载
		panic("Failed to read JWT config: " + err.Error())
	}

	if err := yaml.Unmarshal(configData, &jwtConfig); err != nil {
		panic("Failed to parse JWT config: " + err.Error())
	}
}

// GenerateJWT 生成JWT令牌
func GenerateJWT(userID uint) (string, error) {
	// 设置令牌过期时间
	expirationTime := time.Now().Add(time.Duration(jwtConfig.JWT.ExpireHours) * time.Hour)

	// 创建claims
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    expirationTime.Unix(),
		"iat":    time.Now().Unix(),
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并获取完整的编码后的字符串令牌
	tokenString, err := token.SignedString([]byte(jwtConfig.JWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT 解析并验证JWT令牌
func ParseJWT(tokenString string) (uint, error) {
	// 解析令牌
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtConfig.JWT.Secret), nil
	})

	if err != nil {
		return 0, err
	}

	// 提取claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	// 获取userID
	userIDFloat, ok := claims["userID"].(float64)
	if !ok {
		return 0, errors.New("invalid userID in token")
	}

	return uint(userIDFloat), nil
}
