package restapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/itimofeev/auchan/server/models"
	"github.com/itimofeev/auchan/server/restapi/operations"
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

	return user.NewGetCurrentUserOK().WithPayload(convertUser(usr))
})

var BasketCreateBasketHandler = basket.CreateBasketHandlerFunc(func(params basket.CreateBasketParams, principal interface{}) middleware.Responder {
	usr := principal.(*models.User)
	created, err := Service.CreateBasket(&store.User{ID: usr.ID}, *params.Basket.Name)
	if err != nil {
		return util.ConvertHTTPErrorToResponse(err)
	}
	return basket.NewCreateBasketOK().WithPayload(convertBasket(created))
})

var BasketGetAllBasketsHandler = basket.GetAllBasketsHandlerFunc(func(params basket.GetAllBasketsParams, principal interface{}) middleware.Responder {
	usr := principal.(*models.User)
	b, err := Service.GetUserBaskets(&store.User{ID: usr.ID})
	if err != nil {
		return util.ConvertHTTPErrorToResponse(err)
	}
	bs := make([]*models.Basket, 0, len(b))
	for _, bb := range b {
		bs = append(bs, convertBasket(bb))
	}
	return basket.NewGetAllBasketsOK().WithPayload(bs)
})

var ProductGetProductsByParamsHandler = product.GetProductsByParamsHandlerFunc(func(params product.GetProductsByParamsParams) middleware.Responder {
	products, err := Service.SearchProducts(params.Name)
	if err != nil {
		return util.ConvertHTTPErrorToResponse(err)
	}

	var resp = make([]*models.Product, 0, len(products))
	for _, prod := range products {
		resp = append(resp, convertProduct(prod))
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
		resp = append(resp, convertGoods(g))
	}

	return goods.NewGetAllGoodsInBasketOK().WithPayload(resp)
})

func convertProduct(g *store.Product) *models.Product {
	return &models.Product{
		ID:           g.ID,
		Name:         g.Name,
		CategoryName: g.CategoryName,
		ImageURL:     g.ImageURL,
	}
}

func convertUser(g *store.User) *models.User {
	return &models.User{
		ID:    g.ID,
		Email: g.Email,
	}
}

func convertGoods(g *store.Goods) *models.Goods {
	return &models.Goods{
		ID:        g.ID,
		Completed: g.Completed,
		Product:   convertProduct(g.Product),
		User:      convertUser(g.User),
		Price:     g.Price,
		Quantity:  g.Quantity,
		Unit:      g.Unit,
	}
}

func convertBasket(g *store.Basket) *models.Basket {
	return &models.Basket{
		ID:   g.ID,
		Name: g.Name,
	}
}

var ShareGetAllSharesForBasketHandler = share.GetAllSharesForBasketHandlerFunc(func(params share.GetAllSharesForBasketParams, principal interface{}) middleware.Responder {
	sshares, err := Service.GetSharesForBasket(&store.Basket{ID: params.BasketID})
	if err != nil {
		return util.ConvertHTTPErrorToResponse(err)
	}

	var resp = make([]*models.Share, 0, len(sshares))
	for _, g := range sshares {
		resp = append(resp, &models.Share{
			User: convertUser(g.User),
		})
	}
	return share.NewGetAllSharesForBasketOK().WithPayload(resp)
})

var ShareAddUserToShareHandler = share.AddUserToShareHandlerFunc(func(params share.AddUserToShareParams, principal interface{}) middleware.Responder {
	sh, err := Service.AddUserToShare(&store.Basket{ID: params.BasketID}, *params.Share.Email)
	if err != nil {
		return util.ConvertHTTPErrorToResponse(err)
	}

	return share.NewAddUserToShareOK().WithPayload(&models.Share{User: convertUser(sh.User)})
})

var GoodsAddGoodsToBasketHandler = goods.AddGoodsToBasketHandlerFunc(func(params goods.AddGoodsToBasketParams, principal interface{}) middleware.Responder {
	usr1 := principal.(*models.User)

	gds, err := Service.UpdateGoodsInBasket(&store.User{ID: usr1.ID, Email: usr1.Email}, &store.Basket{ID: params.BasketID}, params.Goods.ProductID, params.Goods.Quantity)
	if err != nil {
		return util.ConvertHTTPErrorToResponse(err)
	}
	return goods.NewAddGoodsToBasketOK().WithPayload(convertGoods(gds))
})

var HelloHandler = operations.HelloHandlerFunc(func(params operations.HelloParams) middleware.Responder {
	return operations.NewHelloOK().WithPayload("hi, there!")
})
