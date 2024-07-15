package location

import (
	"boilerplate/internal/middleware"
	"boilerplate/internal/wrapper/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRoutes(api fiber.Router, handler handler.Handler) {
	api.Post("/location", middleware.AdminAuthMiddleware(), handler.Core.Location.Create)
	api.Get("/location", middleware.AdminAuthMiddleware(), handler.Core.Location.GetAllLocation)
	api.Get("/location/getbyid", middleware.AdminAuthMiddleware(), handler.Core.Location.GetLocationById)
	api.Get("/location/getbyname", middleware.AdminAuthMiddleware(), handler.Core.Location.GetLocationByName)
	api.Put("/location", middleware.AdminAuthMiddleware(), handler.Core.Location.Update)
	api.Delete("/location", middleware.AdminAuthMiddleware(), handler.Core.Location.Delete)
}
