package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"server/service/internal/handle"
	"server/util"
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

	apiGroup := app.Group("/api")
	apiGroup.Post("/login", handle.UserLogin)
	apiGroup.Get("/verify", handle.Verify())
	apiGroup.Post("/update/password", util.JWTMiddleware(), handle.UpdatePassword)

	stu := apiGroup.Group("/student", util.JWTMiddleware())
	{
		stu.Get("/select", handle.Select)
		stu.Post("/show", handle.ShowTopic)
		stu.Get("/info", handle.MyTopic)
		stu.Get("/cancel", handle.Cancel)
	}

	teacher := apiGroup.Group("/teacher", util.JWTMiddleware())
	{
		teacher.Post("/add", handle.AddTopic)
		teacher.Post("/search", handle.SearchMyTopic)
		teacher.Post("/delete", handle.DeleteTopic)
		teacher.Post("/edit", handle.UpdateTopic)
	}
	return app
}
