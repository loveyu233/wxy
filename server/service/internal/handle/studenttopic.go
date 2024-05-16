package handle

import (
	"github.com/gofiber/fiber/v2"
	"server/dao"
	"server/model"
	"server/util"
)

// Select 选择论文
func Select(c *fiber.Ctx) error {
	id := c.Query("id")
	uid, st := c.Locals("userId").(string)
	if !st {
		return util.Resp400(c, "类型断言失败")
	}

	var count int64
	dao.DB.Model(&model.Topic{}).Debug().Where("student_id = ?", uid).Count(&count)

	if count > 0 {
		return util.Resp400(c, "已选取不可重复选择")
	}

	if err := dao.DB.Table("topics").
		Where("id=?", id).
		UpdateColumn("student_id", uid).
		Error; err != nil {
		return util.Resp400(c, "选择失败")
	}

	return util.Resp200(c, "选择成功")
}

// ShowTopic 展示未被选择的论文
func ShowTopic(c *fiber.Ctx) error {
	type Req struct {
		PageNum  int64 `json:"pageNum" form:"pageNum"`
		PageSize int64 `json:"pageSize" form:"pageSize"`
	}

	var (
		topic []model.Topic
		num   int64
		err   error
		req   Req
	)
	c.BodyParser(&req)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}
	tx1 := dao.DB.Model(&model.Topic{}).
		Where("is_select = false and deleted_at is NULL")
	tx2 := dao.DB.Model(model.Topic{}).
		Where("is_select = false and deleted_at is NULL")
	if err = tx1.
		Offset(int((req.PageNum - 1) * req.PageSize)).Limit(int(req.PageSize)).
		Find(&topic).Error; err != nil {
		return util.Resp400(c, "查找论文失败")
	}
	tx2.Count(&num)
	return util.Resp200(c, fiber.Map{
		"topicList": topic,
		"total":     num,
	})
}

// MyTopic 我选择的论文
func MyTopic(c *fiber.Ctx) error {
	uid, st := c.Locals("userId").(string)
	if !st {
		return util.Resp400(c, "类型断言失败")
	}
	var (
		topic []*model.Topic
		err   error
	)
	if err = dao.DB.Model(model.Topic{}).
		Where("student_id=? and deleted_at is NULL", uid).
		Find(&topic).Error; err != nil {
		return util.Resp200(c, "未选择论文")
	}

	return util.Resp200(c, topic)
}

func Cancel(c *fiber.Ctx) error {
	id := c.Query("id")

	if err := dao.DB.Model(model.Topic{}).
		Where("id=? and deleted_at is NULL", id).
		UpdateColumn("student_id", nil).Error; err != nil {
		return util.Resp400(c, "取消失败")
	}

	return util.Resp200(c, "取消选择成功")
}
