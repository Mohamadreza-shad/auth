package health

import (
	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	c.Writer.Write([]byte("OK"))
}
