package health

import (
	"github.com/gin-gonic/gin"
)

type HealthRouter struct {
	Handler *gin.Engine
}

func NewHealthRouter() *HealthRouter {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/", HealthHandler)

	return &HealthRouter{
		Handler: r,
	}
}
