package service

type LoginService interface{
	Login(username string, password string) bool
}

type loginService struct{
	authourizesUsername string
	authourizedPassword string
}

func NewLoginService() LoginService{
	return &loginService{
		authourizesUsername: "pragmatic",
		authourizedPassword: "reviews",
	}
}

func (service *loginService) Login(username string, password string) bool{
	return service.authourizesUsername == username && 
	service.authourizedPassword==password
}