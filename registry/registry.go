package registry

import "github.com/sugartr3e/tweat/registry/container"

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
