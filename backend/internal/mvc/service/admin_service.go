package service

import (
	"errors"
	"gt-blog/backend/internal/mvc/dto"
	"gt-blog/backend/internal/mvc/model"
	"gt-blog/backend/internal/mvc/repository"
	"gt-blog/backend/pkg/util"

	"github.com/gin-gonic/gin"
)

type AdminService struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(adminRepo repository.AdminRepository) *AdminService {
	return &AdminService{
		adminRepo: adminRepo,
	}
}

func (as *AdminService) GetAdminInfo(c *gin.Context, req *dto.GetAdminInfoReq) (*dto.GetAdminInfoResp, error) {
	admin, err := as.adminRepo.GetAdminByID(c, req.ID)
	if err != nil {
		return nil, err
	}
	// 创建DTO社交账号数组
	socialAccounts := make([]dto.SocialAccount, 0, len(admin.SocialAccounts))
	// 将model的SocialAccounts转换为dto的SocialAccounts
	for _, sa := range admin.SocialAccounts {
		socialAccounts = append(socialAccounts, dto.SocialAccount{
			ID:       sa.ID,
			UserID:   sa.UserID,
			Platform: sa.Platform,
			Url:      sa.Url,
		})
	}
	resp := &dto.GetAdminInfoResp{
		ID:             admin.ID,
		Username:       admin.Username,
		Nickname:       admin.Nickname,
		AvatarUrl:      admin.AvatarUrl,
		Email:          admin.Email,
		Wechat:         admin.Wechat,
		Bio:            admin.Bio,
		AboutMe:        admin.AboutMe,
		AboutBlog:      admin.AboutBlog,
		SocialAccounts: socialAccounts,
	}
	return resp, nil
}

func (as *AdminService) UpdateAdminInfo(c *gin.Context, req *dto.UpdateAdminInfoReq) error {
	// 先查询完整的管理员信息
	admin, err := as.adminRepo.GetAdminByID(c, req.ID)
	if err != nil {
		return err
	}

	// 只更新请求中提供的字段
	admin.Nickname = req.Nickname
	admin.AvatarUrl = req.AvatarUrl
	admin.Email = req.Email
	admin.Wechat = req.Wechat
	admin.Bio = req.Bio
	admin.AboutMe = req.AboutMe
	admin.AboutBlog = req.AboutBlog

	// 处理社交账号
	if req.SocialAccounts != nil {
		admin.SocialAccounts = []model.SocialAccount{}
		for _, sa := range req.SocialAccounts {
			admin.SocialAccounts = append(admin.SocialAccounts, model.SocialAccount{
				ID:       sa.ID,
				UserID:   admin.ID,
				Platform: sa.Platform,
				Url:      sa.Url,
			})
		}
	}

	return as.adminRepo.UpdateAdmin(c, admin)
}

// ChangePassword 修改管理员密码
// 使用JWT后，从context中获取用户ID，不再依赖请求中的ID
func (as *AdminService) ChangePassword(c *gin.Context, req *dto.ChangePasswordReq) error {
	// 从context中获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		return errors.New("无法获取用户信息")
	}

	// 类型断言
	id, ok := userID.(uint)
	if !ok {
		return errors.New("用户ID格式错误")
	}

	// 使用从JWT中获取的用户ID查询管理员信息
	admin, err := as.adminRepo.GetAdminByID(c, id)
	if err != nil {
		return err
	}

	// 验证当前密码
	match, err := util.ComparePassword(admin.Password, req.CurrentPassword)
	if err != nil {
		return err
	}
	if !match {
		return errors.New("当前密码不正确")
	}

	// 检查新密码是否与当前密码相同
	if req.CurrentPassword == req.NewPassword {
		return errors.New("新密码不能与当前密码相同")
	}

	// 哈希新密码
	hashedPassword, err := util.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}
	// 更新密码字段为哈希值
	admin.Password = hashedPassword

	return as.adminRepo.UpdateAdmin(c, admin)
}

func (as *AdminService) Login(c *gin.Context, req *dto.LoginReq) (*dto.LoginResp, error) {
	// 先查询完整的管理员信息
	admin, err := as.adminRepo.GetAdminByUsername(c, req.Username)
	if err != nil {
		return nil, err
	}

	// 验证密码
	match, err := util.ComparePassword(admin.Password, req.Password)
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, errors.New("password is incorrect")
	}

	// 生成JWT token，使用公共工具函数
	token, err := util.GenerateJWT(admin.ID)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResp{
		Token: token,
	}, nil
}
