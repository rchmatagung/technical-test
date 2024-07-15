package department

import (
	"boilerplate/internal/middleware"
	"boilerplate/internal/wrapper/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRoutes(api fiber.Router, handler handler.Handler) {
	api.Post("/department", middleware.AdminAuthMiddleware(), handler.Core.Department.Create)
	api.Get("/department", middleware.AdminAuthMiddleware(), handler.Core.Department.GetAllDepartment)
	api.Get("/department/getbyid", middleware.AdminAuthMiddleware(), handler.Core.Department.GetDepartmentById)
	api.Get("/department/getbyname", middleware.AdminAuthMiddleware(), handler.Core.Department.GetDepartmentByName)
	api.Put("/department", middleware.AdminAuthMiddleware(), handler.Core.Department.Update)
	api.Delete("/department", middleware.AdminAuthMiddleware(), handler.Core.Department.Delete)
}
