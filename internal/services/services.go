package services

type Services struct {
	UserService     UserService
	ShortUrlService ShortUrlService
}

func New(userService UserService, shortUrlService ShortUrlService) *Services {
	return &Services{
		UserService:     userService,
		ShortUrlService: shortUrlService,
	}
}
