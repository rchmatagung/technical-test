package server

import (
	"boilerplate/config"
	"boilerplate/internal/middleware"
	"boilerplate/internal/wrapper/handler"
	"boilerplate/internal/wrapper/repository"
	"boilerplate/internal/wrapper/usecase"
	"boilerplate/pkg/infra/db"
	"fmt"
	"log"

	coreLocation "boilerplate/internal/core/location"
	coreDepartment "boilerplate/internal/core/department"
	corePositions "boilerplate/internal/core/position"
	coreEmployee "boilerplate/internal/core/employee"
	coreAttendance "boilerplate/internal/core/attendance"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/sirupsen/logrus"
)

func Run(conf *config.Config, dbList *db.DatabaseList, appLoger *logrus.Logger) {

	//* Initial Engine
	engine := html.New("./views", ".html")

	//* Initial Fiber App
	app := fiber.New(fiber.Config{
		AppName:      conf.App.Name,
		ServerHeader: "Go Fiber",
		Views:        engine,
		BodyLimit:    conf.App.BodyLimit * 1024 * 1024,
	})

	//* Initial Data Middleware
	middleware.InitMiddlewareConfig(app, dbList, conf, appLoger)

	//* General Middleware
	middleware.CORSMiddleware()
	middleware.DefaultLimitterMiddleware()
	//middleware.RecoverMiddleware()

	//* Initial Wrapper

	if dbList.DatabaseApp == nil {
		log.Println("is nil")
	}

	repo := repository.NewRepository(conf, dbList, appLoger)
	usecase := usecase.NewUsecase(repo, conf, dbList, appLoger)
	handler := handler.NewHandler(usecase, conf, appLoger)

	//* Root Endpoint
	app.Get("/", handler.General.Root.GetRoot)

	//* Api Endpoint
	api := app.Group(conf.App.Endpoint)

	//* General Routes
	//generalEncyrption.NewRoutes(api, handler)

	//* Core Routes
	coreLocation.NewRoutes(api, handler)
	coreDepartment.NewRoutes(api, handler)
	corePositions.NewRoutes(api, handler)
	coreEmployee.NewRoutes(api, handler)
	coreAttendance.NewRoutes(api, handler)


	//* CMS Routes
	// cmsWorkOfType.NewRoutes(api, handler)

	//* Not found
	app.All("*", handler.General.NotFound.GetNotFound)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", conf.App.Port)))
}
