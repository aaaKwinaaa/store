package routes

import (
	"store/app"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, ctrl *app.SetupController) (*gin.RouterGroup, error) {
	v1 := router.Group("/api/v1")
	return v1, nil

}
