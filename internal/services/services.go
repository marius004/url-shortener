package services

type Services struct {
	UserService UserService
}

func New(userService UserService) *Services {
	return &Services{
		UserService: userService,
	}
}
