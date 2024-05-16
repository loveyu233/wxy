package handle

import (
	"github.com/gofiber/fiber/v2"
	"server/dao"
	"server/model"
	"server/util"
)

// AddTopic 教师操作 增加论文
func AddTopic(c *fiber.Ctx) error {
	var req model.Topic

	if err := c.BodyParser(&req); err != nil {
		return util.Resp400(c, "参数绑定失败")
	}
	uid, st := c.Locals("userId").(string)
	if !st {
		return util.Resp400(c, "类型断言失败")
	}
	username, st := c.Locals("username").(string)
	if !st {
		return util.Resp400(c, "类型断言失败")
	}
	topic := &model.Topic{
		ThesisName:         req.ThesisName,
		EnglishName:        req.EnglishName,
		TeacherName:        username,
		TeacherId:          uid,
		Title:              req.Title,
		Unit:               req.Unit,
		ThesisType:         req.ThesisType,
		Nature:             req.Nature,
		Source:             req.Source,
		IsSelect:           req.IsSelect,
		Content:            req.Content,
		TaskBook:           req.TaskBook,
		SpecialRequest:     req.SpecialRequest,
		ReferenceMaterials: req.ReferenceMaterials,
	}

	if err := dao.DB.Model(&model.Topic{}).
		Create(topic).
		Error; err != nil {
		return util.Resp400(c, "添加数据库失败")
	}

	return util.Resp200(c, "添加论文成功")
}

// SearchMyTopic 查找自己发布的论文
func SearchMyTopic(c *fiber.Ctx) error {
	uid, st := c.Locals("userId").(string)
	if !st {
		return util.Resp400(c, "类型断言失败")
	}
	utype, st := c.Locals("type").(string)
	if !st || utype != "teacher" {
		return util.Resp400(c, "该操作没有权限")
	}
	var topic []model.Topic
	if err := dao.DB.Model(model.Topic{}).
		Where("teacher_id=?", uid).
		Find(&topic).
		Error; err != nil {
		return util.Resp400(c, "查找失败")
	}

	return util.Resp200(c, topic)
}

// DeleteTopic 教师操作 删除论文
func DeleteTopic(c *fiber.Ctx) error {
	query := c.QueryInt("id")

	if err := dao.DB.Where("id=?", query).
		Delete(&model.Topic{}).
		Error; err != nil {
		return util.Resp400(c, "删除失败")
	}
	return util.Resp200(c, "删除成功")
}

// UpdateTopic 教师操作 修改论文
func UpdateTopic(c *fiber.Ctx) error {
	var req *model.Topic

	if err := c.BodyParser(&req); err != nil {
		return util.Resp400(c, "参数绑定失败")
	}
	uid, st := c.Locals("userId").(string)
	if !st {
		return util.Resp400(c, "类型断言失败")
	}
	topic := model.Topic{
		ThesisName:         req.ThesisName,
		EnglishName:        req.EnglishName,
		TeacherName:        req.TeacherName,
		TeacherId:          uid,
		Title:              req.Title,
		Unit:               req.Unit,
		ThesisType:         req.ThesisType,
		Nature:             req.Nature,
		Source:             req.Source,
		IsSelect:           req.IsSelect,
		Content:            req.Content,
		TaskBook:           req.TaskBook,
		SpecialRequest:     req.SpecialRequest,
		ReferenceMaterials: req.ReferenceMaterials,
	}
	if err := dao.DB.Model(model.Topic{}).
		Where("id=?", req.ID).
		Updates(topic).
		Error; err != nil {
		return util.Resp400(c, "论文信息修改失败")
	}
	return util.Resp200(c, "修改成功")
}
