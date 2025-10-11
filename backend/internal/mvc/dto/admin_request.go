package dto

// GetAdminInfoReq 定义了获取管理员信息的请求参数
// 使用JWT后，ID字段由后端从JWT令牌中自动填充，不需要前端传递

type GetAdminInfoReq struct {
	ID uint `form:"id"`
}

type SocialAccount struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"useId"`
	Platform string `json:"platform"`
	Url      string `json:"url"`
}

type UpdateAdminInfoReq struct {
	ID             uint            `json:"id"`
	Nickname       string          `json:"nickname"`
	AvatarUrl      string          `json:"avatarUrl"`
	Email          string          `json:"email"`
	Wechat         string          `json:"wechat"`
	Bio            string          `json:"bio"`
	AboutMe        string          `json:"aboutMe"`
	AboutBlog      string          `json:"aboutBlog"`
	SocialAccounts []SocialAccount `json:"socialAccounts"`
}

// ChangePasswordReq 定义了修改密码请求参数
// 使用JWT后，不需要再从请求中传递ID
// 而是从JWT令牌中获取登录用户的ID

type ChangePasswordReq struct {
	// ID uint `json:"id"`  // 不再需要，从JWT中获取
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
