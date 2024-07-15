package employee

import (
	"boilerplate/internal/middleware"
	"boilerplate/internal/wrapper/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRoutes(api fiber.Router, handler handler.Handler) {
	api.Post("/employee", middleware.AdminAuthMiddleware(), handler.Core.Employee.Create)
	api.Get("/employee", middleware.EmployeeAuthMiddleware(), handler.Core.Employee.GetAllEmployee)
	api.Get("/employee/getbyid", middleware.EmployeeAuthMiddleware(), handler.Core.Employee.GetEmployeeById)
	api.Get("/employee/getbyname", middleware.EmployeeAuthMiddleware(), handler.Core.Employee.GetEmployeeByName)
	api.Put("/employee", middleware.AdminAuthMiddleware(), handler.Core.Employee.Update)
	api.Delete("/employee", middleware.AdminAuthMiddleware(), handler.Core.Employee.Delete)

	api.Post("/employee/login", handler.Core.Employee.Login)
	api.Put("/employee/forgotpassword", middleware.EmployeeAuthMiddleware(), handler.Core.Employee.ForgotPassword)
}
