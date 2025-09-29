package dto

type ProfileUpdateRequest struct {
	AvatarURL   string   `json:"avatarUrl"`
	Nickname    string   `json:"nickname"`
	Bio         string   `json:"bio"`
	Email       string   `json:"email"`
	Wechat      string   `json:"wechat"`
	AboutMe     string   `json:"aboutMe"`
	AboutBlog   string   `json:"aboutBlog"`
	SocialMedia []string `json:"socialMedia"`
}
