package attendance

import (
	"boilerplate/internal/middleware"
	"boilerplate/internal/wrapper/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRoutes(api fiber.Router, handler handler.Handler) {
	api.Post("/absentin", middleware.EmployeeAuthMiddleware(), handler.Core.Attendance.AttendanceAbsentIn)
	api.Put("/absentout", middleware.EmployeeAuthMiddleware(), handler.Core.Attendance.AttendanceAbsentOut)
	api.Get("/absent", middleware.AdminAuthMiddleware(), handler.Core.Attendance.GetAllAttendance)
	api.Get("/absent/getbyid", middleware.AdminAuthMiddleware(), handler.Core.Attendance.GetAttendanceById)
	api.Get("/absent/getbyemployeeid", middleware.AdminAuthMiddleware(), handler.Core.Attendance.GetAttendanceByEmployeeId)
	api.Delete("/absent", middleware.AdminAuthMiddleware(), handler.Core.Attendance.Delete)

}
