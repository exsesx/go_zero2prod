package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct{}

func (h HealthController) Status(c *gin.Context) {
	c.Status(http.StatusOK)
}
