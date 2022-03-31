package server

import (
	"github.com/gin-gonic/gin"
	"go_zero2prod/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	return router

}
