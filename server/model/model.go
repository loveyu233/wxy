package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId   string `json:"userId" gorm:"column:user_id"`    // 学号
	UserName string `json:"userName" gorm:"column:username"` //姓名
	PassWord string `json:"passWord" gorm:"column:password"` //密码
	Type     string `json:"type" gorm:"column:type"`         // teacher 老师 student 学生
}

type Topic struct {
	gorm.Model
	ThesisName         string `json:"thesis_name" gorm:"column:thesis_name"`                 //论文名称
	EnglishName        string `json:"english_name" gorm:"column:english_name"`               //英文名称
	TeacherName        string `json:"teacher_name" gorm:"column:teacher_name"`               //指导教师
	TeacherId          string `json:"teacher_id" gorm:"column:teacher_id"`                   //指导教师Id
	Title              string `json:"title" gorm:"column:title"`                             //教师职称
	Unit               string `json:"unit" gorm:"column:unit"`                               //单位
	ThesisType         string `json:"thesis_type" gorm:"column:thesis_type"`                 //题目类别
	Nature             string `json:"nature" gorm:"column:nature"`                           //题目性质
	Source             string `json:"source" gorm:"column:source"`                           //题目来源
	Content            string `json:"content" gorm:"column:content"`                         //内容
	TaskBook           string `json:"task_book" gorm:"column:task_book"`                     //任务书
	SpecialRequest     string `json:"special_request" gorm:"column:special_request"`         //特殊要求
	ReferenceMaterials string `json:"reference_materials" gorm:"column:reference_materials"` //参考文献
	IsSelect           bool   `json:"is_select" gorm:"column:is_select"`                     //是否已被选择
	StudentId          string `json:"student_id" gorm:"column:student_id"`                   //学生Id
}
