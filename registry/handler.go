package registry

import "github.com/sugartr3e/tweat/handler"

type AppHandler struct {
	Th handler.TweatHandler
	Uh handler.UserHandler
}

func NewAppHandler(th handler.TweatHandler, uh handler.UserHandler) *AppHandler {
	return &AppHandler{th, uh}
}
