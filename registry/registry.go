package registry

import "github.com/gba-3/tweat/registry/container"

func NewRegistry() *AppHandler {
	c := container.Container{}
	return NewAppHandler(
		c.GetTweatHandler(
			c.GetTweatUsecase(
				c.GetTweatRepository(),
			),
		),
		c.GetUserHandler(
			c.GetUserUsecase(
				c.GetUserRepository(),
			),
		),
	)
}
