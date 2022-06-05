package registry

import "github.com/gba-3/tweat/handler"

type AppHandler struct {
	Th handler.TweatHandler
}

func NewAppHandler(th handler.TweatHandler) *AppHandler {
	return &AppHandler{th}
}
