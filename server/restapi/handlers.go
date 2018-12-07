package restapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/itimofeev/auchan/server/models"
	"github.com/itimofeev/auchan/server/restapi/operations/basket"
	"github.com/itimofeev/auchan/server/restapi/operations/user"
	"github.com/itimofeev/auchan/service"
	"github.com/itimofeev/auchan/store"
	"github.com/itimofeev/auchan/util"
)

var Service *service.Service

var AuthFunc = func(token string) (interface{}, error) {
	user, err := Service.GetByTokenEnsure(token)
	if err != nil {
		e := util.ConvertHTTPErrorToResponse(err)
		t, _ := e.(error)
		return nil, t

	}
	return user, nil
}

var UserLoginUserHandler = user.LoginUserHandlerFunc(func(params user.LoginUserParams) middleware.Responder {
	token, err := Service.LoginUser(params.Email, params.Password)
	if err != nil {
		return util.ConvertHTTPErrorToResponse(err)
	}
	return user.NewLoginUserOK().WithXAuthToken(token)
})

var UserGetCurrentUserHandler = user.GetCurrentUserHandlerFunc(func(params user.GetCurrentUserParams, principal interface{}) middleware.Responder {
	usr1 := principal.(*models.User)
	usr, err := Service.GetUserByID(usr1.ID)
	if err != nil {
		return util.ConvertHTTPErrorToResponse(err)
	}

	return user.NewGetCurrentUserOK().WithPayload(&models.User{
		ID:    usr.ID,
		Email: usr.Email,
	})
})

var BasketCreateBasketHandler = basket.CreateBasketHandlerFunc(func(params basket.CreateBasketParams, principal interface{}) middleware.Responder {
	usr := principal.(*models.User)
	created, err := Service.CreateBasket(&store.User{ID: usr.ID}, *params.Basket.Name)
	if err != nil {
		return util.ConvertHTTPErrorToResponse(err)
	}
	return basket.NewCreateBasketOK().WithPayload(&models.Basket{
		ID:   created.ID,
		Name: created.Name,
	})
})
