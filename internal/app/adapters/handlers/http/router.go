package http

import (
	"errors"
	"github.com/Mohamadreza-shad/auth/pkg/i18n"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Handler *gin.Engine
}

func NewRouter(h *GoHexagonalHttpHandler, l i18n.I18n) *Router {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(TrackingIDMiddleware())

	r.NoRoute(func(c *gin.Context) {
		MakeErrorLocalizedResponse(c, errors.New("url not found"), l, "en")
	})

	return &Router{Handler: r}
}
