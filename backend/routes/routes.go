package routes

import (
	"jobsforgobe/database"
	"jobsforgobe/handler"

	"github.com/labstack/echo/v4"
)

func RouteSetup(e *echo.Echo) {
	jobsController := handler.JobsController(database.DBAccess)
	authController := handler.AuthController(database.DBAccess)


	e.Static("/img-jobs", "image")

	e.POST("/v1/signup", authController.SignUpHandler)
	e.POST("/v1/signin", authController.SignInHandler)

	e.GET("/v1/jobs", jobsController.FindAllJobHandler)
	e.GET("/v1/jobs/:category", jobsController.FindAllJobByCategoryHandler)

}
