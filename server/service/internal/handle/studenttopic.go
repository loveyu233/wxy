package handle

import (
	"TopicSelection/dao"
	"TopicSelection/model"
	"TopicSelection/util"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// Select 选择论文
func Select(c *fiber.Ctx) error {
	type Req struct {
		ID uint `json:"id"`
	}
	var req Req
	if err := c.BodyParser(&req); err != nil {
		return util.Resp400(c, "参数绑定失败")
	}
	uid, st := c.Locals("userId").(string)
	if !st {
		return util.Resp400(c, "类型断言失败")
	}
	tx := dao.DB.Model(model.Topic{}).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	dao.DB.Clauses()
	if err := dao.DB.Table("topics").
		Where("id=?", req.ID).
		UpdateColumn("student_id", uid).
		Error; err != nil {
		tx.Rollback()
		return util.Resp400(c, "选择失败")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("数据库提交失败 ：%v", err)
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
	tx1 := dao.DB.Model(&model.Topic{}).
		Where("student_id = ? and deleted_at is NULL", "")
	tx2 := dao.DB.Model(model.Topic{}).
		Where("student_id = ? and deleted_at is NULL", "")
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
		topic model.Topic
		err   error
	)
	if err = dao.DB.Model(model.Topic{}).
		Where("student_id=? and deleted_at is not NULL", uid).
		First(&topic).Error; err != nil {
		return util.Resp200(c, "未选择论文")
	}

	return util.Resp200(c, topic)
}

func Cancel(c *fiber.Ctx) error {
	id := c.Query("id")

	if err := dao.DB.Model(model.Topic{}).
		Where("id=? and deleted_at is not NULL", id).
		UpdateColumn("student_id", nil).Error; err != nil {
		return util.Resp400(c, "取消失败")
	}

	return util.Resp200(c, "取消选择成功")
}
