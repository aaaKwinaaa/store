package main

import (
	"store/app"
	"store/config"

	publicRoutes "store/routes/public"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()

}

func main() {
	app := app.New()
	StartServer(app)
}

func StartServer(app *app.App) {
	r := gin.Default()

	if _, err := publicRoutes.SetupRoutes(&r.RouterGroup, &app.Controller); err != nil {
		panic(err)
	}

	if err := r.Run(); err != nil {
		panic("Failed to start the project")
	}
}
