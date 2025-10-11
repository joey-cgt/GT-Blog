package util

import "golang.org/x/crypto/bcrypt"

// HashPassword 使用bcrypt算法对密码进行哈希处理
// 返回哈希后的密码字符串和可能的错误
func HashPassword(password string) (string, error) {
	// GenerateFromPassword会自动生成盐值并返回哈希后的密码
	// 第二个参数是cost值，值越大越安全但计算时间越长，通常取10-14
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// 转换为字符串返回
	return string(hashedPassword), nil
}

// ComparePassword 比较明文密码与哈希密码是否匹配
// 返回是否匹配以及可能的错误
func ComparePassword(hashedPassword, password string) (bool, error) {
	// CompareHashAndPassword比较哈希密码与明文密码
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		// 如果密码不匹配，返回false但不返回错误
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		// 如果有其他错误，返回错误
		return false, err
	}
	// 密码匹配
	return true, nil
}