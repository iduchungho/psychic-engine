package application

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"smhome/pkg/middleware"
	"smhome/pkg/routes"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	r *fiber.App
}

var lock = &sync.Mutex{}
var application *App

func GetApplication() *App {
	// check app is already exist
	if application == nil {
		// Ensure that the instance hasn't yet been
		// initialized by another thread while this one
		// has been waiting for the lock's release.
		lock.Lock()
		defer lock.Unlock()
		if application == nil {
			application = &App{
				r: fiber.New(),
			}
		} else {
			return application
		}
	}
	return application
}

func (app *App) Run() {
	if app.r != nil {
		// define cors middleware
		app.r.Use(cors.New(cors.Config{
			AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language," +
				"Accept-Encoding,Connection,Access-Control-Allow-Origin",
			AllowOrigins:     "*",
			AllowCredentials: true,
			AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		}))

		// logger actions to server
		app.r.Use(logger.New())
		app.r.Use(middleware.Redirect)

		// generate collection
		//database.GenerateCollection()

		// routing services application
		route.SenSorRoute(app.r)
		route.UserRoute(app.r)
		route.ActionRoute(app.r)

		host := os.Getenv("PORT")
		if host != "" {
			err := app.r.Listen(":" + host)
			if err != nil {
				panic("Can't run fiber engine")
			}
		} else {
			err := app.r.Listen("0.0.0.0:8080")
			if err != nil {
				panic("Can't run fiber engine")
			}
		}

	} else {
		panic("fiber Engine not found")
	}
}
