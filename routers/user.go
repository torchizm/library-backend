package routers

import (
	"torchizm/library-backend/api/auth"
	"torchizm/library-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(ctx fiber.Router) {
	route := ctx.Group("/user")
	route.Post("/login", auth.Login)
	route.Post("/register", auth.Register)
	route.Post("/resend-confirmation", auth.ResendMail)
	route.Post("/activate", auth.ActivateAccount)
	route.Post("/logout", middlewares.IsAuth, auth.LogOut)
	route.Post("/forgot-password", auth.ForgotPassword)
	route.Post("/forgot-password-confirm", auth.ForgotPasswordConfirm)
}
