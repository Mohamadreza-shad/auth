package http

import "github.com/Mohamadreza-shad/auth/pkg/logging"

type GoHexagonalHttpHandler struct {
	logger logging.Logger
}

func NewGoHexagonalHttpHandler(l logging.Logger) *GoHexagonalHttpHandler {
	return &GoHexagonalHttpHandler{logger: l}
}
