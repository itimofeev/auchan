package restapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/itimofeev/auchan/server/models"
	"github.com/itimofeev/auchan/server/restapi/operations/basket"
	"github.com/itimofeev/auchan/server/restapi/operations/goods"
	"github.com/itimofeev/auchan/server/restapi/operations/product"
	"github.com/itimofeev/auchan/server/restapi/operations/share"
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

var BasketGetAllBasketsHandler = basket.GetAllBasketsHandlerFunc(func(params basket.GetAllBasketsParams, principal interface{}) middleware.Responder {
	usr := principal.(*models.User)
	b, err := Service.GetUserBaskets(&store.User{ID: usr.ID})
	if err != nil {
		return util.ConvertHTTPErrorToResponse(err)
	}
	bs := make([]*models.Basket, 0, len(b))
	for _, bb := range b {
		bs = append(bs, &models.Basket{
			ID:   bb.ID,
			Name: bb.Name,
		})
	}
	return basket.NewGetAllBasketsOK().WithPayload(bs)
})

var ProductGetProductsByParamsHandler = product.GetProductsByParamsHandlerFunc(func(params product.GetProductsByParamsParams) middleware.Responder {
	products, err := Service.SearchProducts(*params.Name)
	if err != nil {
		return util.ConvertHTTPErrorToResponse(err)
	}

	var resp = make([]*models.Product, 0, len(products))
	for _, prod := range products {
		resp = append(resp, &models.Product{
			ID:         prod.ID,
			Name:       prod.Name,
			ImageURL:   prod.ImageURL,
			CategoryID: prod.CategoryID,
		})
	}

	return product.NewGetProductsByParamsOK().WithPayload(resp)
})

var GoodsGetAllGoodsInBasketHandler = goods.GetAllGoodsInBasketHandlerFunc(func(params goods.GetAllGoodsInBasketParams, principal interface{}) middleware.Responder {
	ggoods, err := Service.GetGoodsForBasket(&store.Basket{ID: params.BasketID})
	if err != nil {
		return util.ConvertHTTPErrorToResponse(err)
	}

	var resp = make([]*models.Goods, 0, len(ggoods))
	for _, g := range ggoods {
		resp = append(resp, &models.Goods{
			ID:        g.ID,
			Completed: g.Completed,
			Product:   nil,
			Price:     g.Price,
			Quantity:  g.Quantity,
			Unit:      g.Unit,
		})
	}

	return goods.NewGetAllGoodsInBasketOK().WithPayload(resp)
})

var ShareGetAllSharesForBasketHandler = share.GetAllSharesForBasketHandlerFunc(func(params share.GetAllSharesForBasketParams, principal interface{}) middleware.Responder {
	sshares, err := Service.GetSharesForBasket(&store.Basket{ID: params.BasketID})
	if err != nil {
		return util.ConvertHTTPErrorToResponse(err)
	}

	var resp = make([]*models.Share, 0, len(sshares))
	for _, g := range sshares {
		resp = append(resp, &models.Share{
			User: &models.User{
				ID:    g.UserID,
				Email: g.User.Email,
			},
		})
	}
	return share.NewGetAllSharesForBasketOK().WithPayload(resp)
})

var ShareAddUserToShareHandler = share.AddUserToShareHandlerFunc(func(params share.AddUserToShareParams, principal interface{}) middleware.Responder {
	sh, err := Service.AddUserToShare(&store.Basket{ID: params.BasketID}, *params.Share.Email)
	if err != nil {
		return util.ConvertHTTPErrorToResponse(err)
	}

	return share.NewAddUserToShareOK().WithPayload(&models.Share{User: &models.User{
		ID:    sh.User.ID,
		Email: sh.User.Email,
	}})
})
