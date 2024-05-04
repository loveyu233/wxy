package util

import (
	"github.com/gofiber/fiber/v2"
)

// JWTMiddleware JWT身份验证的中间件
func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		err := VerifyToken(token)
		if err != nil {
			return Resp400(c, err.Error())
		}
		userid, username, utype, err := GetPayload(token)
		if err != nil {
			return Resp400(c, err.Error())
		}
		c.Locals("userId", userid)
		c.Locals("username", username)
		c.Locals("type", utype)
		return c.Next()
	}
}
