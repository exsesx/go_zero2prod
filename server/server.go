package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	c "go_zero2prod/config"
	"log"
)

func Init() {
	config := c.GetConfig()
	gin.SetMode(config.GetString("GIN_MODE"))

	r := NewRouter()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("Error setting trusted proxies: %s", err)
	}

	err = r.Run(fmt.Sprintf(":%s", config.GetString("PORT")))
	if err != nil {
		log.Fatalf("Error running server: %s", err)
	}
}
