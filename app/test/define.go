package test

import (
	"demo/app"
	"demo/app/test/api"
	. "github.com/arklib/ark"
)

func Define(app *app.App) {
	// add api
	apiFn := api.New(app)
	app.AddApi("test", apiFn)

	// auth middleware
	authMw := app.Auth.HttpMiddleware("user")

	// add routes
	router := app.HttpServer.Group("api/test")
	router.AddRoutes(HttpRoutes{
		{
			Method:  "GET",
			Path:    "ping",
			Handler: ApiHandler[api.PingIn, api.PingOut](apiFn.Ping),
			ApiMiddlewares: ApiMiddlewares{
				func(p *ApiPayload) error {
					err := p.Next()
					if err != nil {
						return err
					}
					out := p.Out.(*api.PingOut)
					out.Message = "(proxy) " + out.Message
					return nil
				},
			},
		},
		{
			Method:  "GET",
			Path:    "cache",
			Handler: ApiHandler[api.CacheGetIn, api.CacheGetOut](apiFn.CacheGet)},
		{
			Method:  "GET",
			Path:    "lock",
			Handler: ApiHandler[api.LockApplyIn, api.LockApplyOut](apiFn.LockApply),
		},
		{
			Path:    "validate",
			Handler: ApiHandler[api.ValidateIn, api.ValidateOut](apiFn.Validate),
		},
		{
			Path:    "upload",
			Handler: ApiHandler[api.UploadIn, api.UploadOut](apiFn.Upload),
		},
		{
			Path:    "sse",
			Handler: ApiHandler[api.SSEIn, api.SSEOut](apiFn.SSE),
		},
		{
			Path:    "sse.req",
			Handler: ApiHandler[api.SSEReqIn, api.SSEReqOut](apiFn.SSEReq),
		},
		{
			Path:    "shop.item.rpc",
			Handler: ApiHandler[api.ShopItemRPCIn, api.ShopItemRPCOut](apiFn.ShopItemRPC),
		},
		{
			Path:    "error",
			Handler: ApiHandler[api.ErrorIn, api.ErrorOut](apiFn.Error)},
		{
			Path:    "token/create",
			Handler: ApiHandler[api.TokenCreateIn, api.TokenCreateOut](apiFn.TokenCreate),
		},
		{
			Path:            "token/auth",
			Handler:         ApiHandler[api.TokenAuthIn, api.TokenAuthOut](apiFn.TokenAuth),
			HttpMiddlewares: HttpMiddlewares{authMw},
		},
	})

	// add rpc
	rpc := app.RPCServer.Group("test")
	rpc.AddRoutes(RPCRoutes{
		{
			Path:    "get",
			Handler: ApiHandler[api.ShopItemGetIn, api.ShopItemGetOut](apiFn.ShopItemGet),
			ApiMiddlewares: ApiMiddlewares{
				func(payload *ApiPayload) error {
					err := payload.Next()
					if err != nil {
						return err
					}

					// dump
					app.Dump(payload.Path, payload.In, payload)

					// rewrite output
					out := payload.Out.(*api.ShopItemGetOut)
					out.Title = "(proxy) " + out.Title
					return nil
				},
			},
		},
	})
}
