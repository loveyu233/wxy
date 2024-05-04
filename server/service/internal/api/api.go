package api

import (
	"TopicSelection/service/internal/handle"
	"TopicSelection/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Cors() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token, x-token")
		c.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		c.Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Set("Access-Control-Allow-Credentials", "true")

		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}

		return c.Next()
	}
}

func RouterInit() *fiber.App {
	app := fiber.New()
	app.Use(Cors())
	app.Use(logger.New())

	app.Post("/login", handle.UserLogin)
	app.Post("/update/password", util.JWTMiddleware(), handle.UpdatePassword)

	stu := app.Group("/student", util.JWTMiddleware())
	{
		stu.Post("/select", handle.Select)
		stu.Post("/show", handle.ShowTopic)
		stu.Post("mytopic", handle.MyTopic)
		stu.Post("/cancel", handle.Cancel)
	}

	teacher := app.Group("/teacher", util.JWTMiddleware())
	{
		teacher.Post("addtopic", handle.AddTopic)
		teacher.Post("searchMyTopic", handle.SearchMyTopic)
		teacher.Post("deleteTopic", handle.DeleteTopic)
		teacher.Post("updateTopic", handle.UpdateTopic)
	}
	return app
}
