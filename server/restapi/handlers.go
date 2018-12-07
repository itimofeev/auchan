package restapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/itimofeev/auchan/server/models"
	"github.com/itimofeev/auchan/server/restapi/operations/user"
	"github.com/itimofeev/auchan/service"
	"github.com/itimofeev/auchan/util"
)

var Service *service.Service

var UserLoginUserHandler = user.LoginUserHandlerFunc(func(params user.LoginUserParams) middleware.Responder {
	token, err := Service.LoginUser(params.Email, params.Password)
	if err != nil {
		return util.ConvertHTTPErrorToResponse(err)
	}
	return user.NewLoginUserOK().WithXAuthToken(token)
})

var UserGetUserByNameHandler = user.GetUserByNameHandlerFunc(func(params user.GetUserByNameParams, principal interface{}) middleware.Responder {
	usr, err := Service.GetUserByEmail(params.Email)
	if err != nil {
		return util.ConvertHTTPErrorToResponse(err)
	}

	return user.NewGetUserByNameOK().WithPayload(&models.User{
		ID:    usr.ID,
		Email: usr.Email,
	})
})

var AuthFunc = func(token string) (interface{}, error) {
	user, err := Service.GetByTokenEnsure(token)
	if err != nil {
		e := util.ConvertHTTPErrorToResponse(err)
		t, _ := e.(error)
		return nil, t

	}
	return user, nil
}
