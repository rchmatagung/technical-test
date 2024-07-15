package position

import (
	"boilerplate/internal/middleware"
	"boilerplate/internal/wrapper/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRoutes(api fiber.Router, handler handler.Handler) {
	api.Post("/positions", middleware.AdminAuthMiddleware(), handler.Core.Position.Create)
	api.Get("/positions", middleware.AdminAuthMiddleware(), handler.Core.Position.GetAllPosition)
	api.Get("/positions/getbyid", middleware.AdminAuthMiddleware(), handler.Core.Position.GetPositionById)
	api.Get("/positions/getbyname", middleware.AdminAuthMiddleware(), handler.Core.Position.GetPositionByName)
	api.Put("/positions", middleware.AdminAuthMiddleware(), handler.Core.Position.Update)
	api.Delete("/positions", middleware.AdminAuthMiddleware(), handler.Core.Position.Delete)
}
