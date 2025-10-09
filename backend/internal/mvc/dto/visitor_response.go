package dto

type AuthorInfoResp struct {
	ID             uint            `json:"id"`
	Nickname       string          `json:"nickname"`
	AvatarUrl      string          `json:"avatarUrl"`
	Bio            string          `json:"bio"`
	Email          string          `json:"email"`
	Wechat         string          `json:"wechat"`
	AboutMe        string          `json:"aboutMe"`
	AboutBlog      string          `json:"aboutBlog"`
	SocialAccounts []SocialAccount `json:"socialAccounts"`
}
