package handlers

import (
	"github.com/DaZZler12/MyRestServer/pkg/service"
)

// type MyUser struct {
// 	models.User
// }

type Handler struct {
	Service service.UserService
}

var userService service.UserService

func SetUserService(us service.UserService) {
	userService = us
}

func New() *Handler {
	return &Handler{
		Service: userService,
	}
}
