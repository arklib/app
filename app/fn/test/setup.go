package test

import (
	. "github.com/arklib/ark"

	"demo/app/base"
)

type Fn struct{ *base.Base }

func Setup(base *base.Base) {
	// add fn
	fn := &Fn{base}
	base.AddFn("test", fn)

	// add api
	router := base.HttpServer
	api := router.Group("api/test")
	api.AddRoutes(HttpRoutes{
		{
			Method:  "GET",
			Path:    "ping",
			Handler: ApiHandler[ApiPingIn, ApiPingOut](fn.ApiPing),
			ApiMiddlewares: ApiMiddlewares{
				func(p *ApiPayload) error {
					err := p.Next()
					if err != nil {
						return err
					}
					out := p.Out.(*ApiPingOut)
					out.Message = "(proxy) " + out.Message
					return nil
				},
			},
		},
		{Method: "GET", Path: "cache", Handler: ApiHandler[ApiCacheGetIn, ApiCacheGetOut](fn.ApiCacheGet)},
		{Method: "GET", Path: "lock", Handler: ApiHandler[ApiLockApplyIn, ApiLockApplyOut](fn.ApiLockApply)},
		{Path: "validate", Handler: ApiHandler[ApiValidateIn, ApiValidateOut](fn.ApiValidate)},
		{Path: "upload", Handler: ApiHandler[ApiUploadIn, ApiUploadOut](fn.ApiUpload)},
		{Path: "sse", Handler: ApiHandler[ApiSSEIn, ApiSSEOut](fn.ApiSSE)},
		{Path: "sse.req", Handler: ApiHandler[ApiSSEReqIn, ApiSSEReqOut](fn.ApiSSEReq)},
		{Path: "shop.item.rpc", Handler: ApiHandler[ApiShopItemRPCIn, ApiShopItemRPCOut](fn.ApiShopItemRPC)},
		{Path: "error", Handler: ApiHandler[ApiErrorIn, ApiErrorOut](fn.ApiError)},
		{Path: "token/create", Handler: ApiHandler[ApiTokenCreateIn, ApiTokenCreateOut](fn.ApiTokenCreate)},
		{
			Path:    "token/auth",
			Handler: ApiHandler[ApiTokenAuthIn, ApiTokenAuthOut](fn.ApiTokenAuth),
			Middlewares: HttpMiddlewares{
				base.Auth.HttpMiddleware("user"),
			},
		},
	})

	// add rpc
	rpc := base.RPCServer.Group("item")
	rpc.AddRoutes(RPCRoutes{
		{
			Path:    "get",
			Handler: ApiHandler[ApiShopItemGetIn, ApiShopItemGetOut](fn.ApiShopItemGet),
			ApiMiddlewares: ApiMiddlewares{
				func(payload *ApiPayload) error {
					err := payload.Next()
					if err != nil {
						return err
					}

					// dump
					base.Dump(payload.Path, payload.In, payload)

					// rewrite output
					out := payload.Out.(*ApiShopItemGetOut)
					out.Title = "(proxy) " + out.Title
					return nil
				},
			},
		},
	})
}