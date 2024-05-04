package handle

import (
	"TopicSelection/bcrypt"
	"TopicSelection/dao"
	"TopicSelection/model"
	"TopicSelection/util"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// UserLogin 登录逻辑函数
func UserLogin(c *fiber.Ctx) error {
	type Req struct {
		UserId   int    `json:"user_id"`
		PassWord string `json:"password"`
	}
	var (
		user model.User
		req  Req
	)
	if err := c.BodyParser(&req); err != nil {
		return util.Resp400(c, "参数绑定失败")
	}
	// 根据userID查询用户信息
	err := dao.DB.Model(model.User{}).Where("user_id = ?", fmt.Sprintf("%d", req.UserId)).First(&user).Error
	if err != nil {
		return util.Resp400(c, "查询失败")
	}
	if user.PassWord == "123456" {
		token, err := util.GenToken(user.UserId, user.UserName, user.Type)
		if err != nil {
			return util.Resp400(c, fmt.Errorf("生成token失败：%v", err))
		}
		return util.Resp200(c, token)
	}
	// 验证密码是否正确
	if !bcrypt.ComparePasswords(req.PassWord, user.PassWord) {
		return util.Resp400(c, "密码错误")
	}

	// 生成token
	token, err := util.GenToken(user.UserId, user.UserName, user.Type)
	if err != nil {
		return util.Resp400(c, fmt.Errorf("生成token失败：%v", err))
	}

	logrus.Info("用户登录成功")
	return util.Resp200(c, fiber.Map{
		"msg":   "登陆成功",
		"token": token,
		"info":  user,
	})
}

// UpdatePassword 修改密码
func UpdatePassword(c *fiber.Ctx) error {
	type Req struct {
		Value string `json:"value"`
	}
	var (
		err error
		req Req
	)
	userID, st := c.Locals("userId").(string)
	if !st {
		return errors.New("类型断言失败")
	}
	err = c.BodyParser(&req)
	if err != nil {
		return util.Resp400(c, "参数绑定失败")
	}
	newPwd, err := bcrypt.HashPassword(req.Value)
	if err != nil {
		return util.Resp500(c, "加密失败")
	}
	err = dao.DB.Table("users").
		Where("user_id = ?", userID).
		UpdateColumn("password", newPwd).
		Error
	if err != nil {
		return util.Resp500(c, "修改密码失败")
	}
	return util.Resp200(c, "修改成功")
}
