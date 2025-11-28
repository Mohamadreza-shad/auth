package http

import (
	"context"
	"github.com/Mohamadreza-shad/auth/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TrackingIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		trackingID := uuid.NewString()
		ctx := context.WithValue(
			c.Request.Context(),
			utils.TrackingIDKey,
			trackingID,
		)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
