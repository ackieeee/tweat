package registry

import "github.com/gba-3/tweat/handler"

type AppHandler struct {
	Th handler.TweatHandler
	Uh handler.UserHandler
}

func NewAppHandler(th handler.TweatHandler, uh handler.UserHandler) *AppHandler {
	return &AppHandler{th, uh}
}
