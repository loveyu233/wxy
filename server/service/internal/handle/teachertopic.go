package handle

import (
	"github.com/gofiber/fiber/v2"
	"server/dao"
	"server/model"
	"server/util"
)

// AddTopic 教师操作 增加论文
func AddTopic(c *fiber.Ctx) error {
	type Req struct {
		ThesisName         string `json:"thesisName"`         //论文名称
		EnglishName        string `json:"englishName"`        //英文名称
		TeacherName        string `json:"teacherName"`        //指导教师
		Title              string `json:"title" `             //教师职称
		Unit               string `json:"unit"`               //单位
		ThesisType         string `json:"thesisType"`         //题目类别
		Nature             string `json:"nature"`             //题目性质
		Source             string `json:"source"`             //题目来源
		IsSelect           bool   `json:"isSelect"`           //是否选择
		Content            string `json:"content"`            //内容
		TaskBook           string `json:"TaskBook"`           //任务书
		SpecialRequest     string `json:"specialRequest"`     //特殊要求
		ReferenceMaterials string `json:"referenceMaterials"` //参考文献
	}
	var req Req

	if err := c.BodyParser(&req); err != nil {
		return util.Resp400(c, "参数绑定失败")
	}
	uid, st := c.Locals("userId").(string)
	if !st {
		return util.Resp400(c, "类型断言失败")
	}
	utype, st := c.Locals("type").(string)
	if !st || utype != "teacher" {
		return util.Resp400(c, "该操作没有权限")
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

	if err := dao.DB.Model(model.Topic{}).
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
	type Req struct {
		ID uint `json:"id"`
	}
	utype, st := c.Locals("type").(string)
	if !st || utype != "teacher" {
		return util.Resp400(c, "该操作没有权限")
	}
	var req Req
	if err := c.BodyParser(&req); err != nil {
		return util.Resp400(c, "参数绑定失败")
	}

	if err := dao.DB.Where("id=?", req.ID).
		Delete(model.Topic{}).
		Error; err != nil {
		return util.Resp400(c, "删除失败")
	}
	return util.Resp200(c, "删除成功")
}

// UpdateTopic 教师操作 修改论文
func UpdateTopic(c *fiber.Ctx) error {
	utype, st := c.Locals("type").(string)
	if !st || utype != "teacher" {
		return util.Resp400(c, "该操作没有权限")
	}
	type Req struct {
		ID                 uint   `json:"id"`
		ThesisName         string `json:"thesisName" gorm:"column:thesisName"`                 //论文名称
		EnglishName        string `json:"englishName" gorm:"column:englishName"`               //英文名称
		TeacherName        string `json:"teacherName" gorm:"column:teacherName"`               //指导教师
		Title              string `json:"title" gorm:"column:title"`                           //教师职称
		Unit               string `json:"unit" gorm:"column:unit"`                             //单位
		ThesisType         string `json:"thesisType" gorm:"column:thesisType"`                 //题目类别
		Nature             string `json:"nature" gorm:"column:nature"`                         //题目性质
		Source             string `json:"source" gorm:"column:source"`                         //题目来源
		IsSelect           bool   `json:"isSelect" gorm:"column:isSelect"`                     //是否选择
		Content            string `json:"content" gorm:"column:content"`                       //内容
		TaskBook           string `json:"TaskBook" gorm:"column:TaskBook"`                     //任务书
		SpecialRequest     string `json:"specialRequest" gorm:"column:specialRequest"`         //特殊要求
		ReferenceMaterials string `json:"referenceMaterials" gorm:"column:referenceMaterials"` //参考文献
	}
	var req Req

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
