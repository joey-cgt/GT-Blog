package dto

type GetAdminInfoResp struct {
	ID             uint            `json:"id"`
	Username       string          `json:"username"`
	Nickname       string          `json:"nickname"`
	AvatarUrl      string          `json:"avatarUrl"`
	Email          string          `json:"email"`
	Wechat         string          `json:"wechat"`
	Bio            string          `json:"bio"`
	AboutMe        string          `json:"aboutMe"`
	AboutBlog      string          `json:"aboutBlog"`
	SocialAccounts []SocialAccount `json:"socialAccounts"`
}

type LoginResp struct {
	Token string `json:"token"`
}
