package restapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/itimofeev/auchan/server/restapi/operations/user"
	"github.com/itimofeev/auchan/service"
)

var Service *service.Service

var UserLoginUserHandler = user.LoginUserHandlerFunc(func(params user.LoginUserParams) middleware.Responder {
	Service.CheckPassword(params.Email, params.Password)
	return user.NewLoginUserOK().WithXAuthToken("hello")
})
