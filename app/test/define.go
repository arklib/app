package test

import (
	"demo/app"
	"demo/app/test/api"
	. "github.com/arklib/ark"
)

func Define(app *app.App) {
	// add api
	testApi := api.New(app)

	// auth middleware
	authMw := app.Auth.HttpMiddleware("user")

	// add routes
	router := app.HttpServer.Group("api/test")
	router.AddRoutes(HttpRoutes{
		{
			Method:  "GET",
			Path:    "ping",
			Handler: ApiHandler[api.PingIn, api.PingOut](testApi.Ping),
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
			Path:    "cache",
			Handler: ApiHandler[api.CacheIn, api.CacheOut](testApi.Cache),
		},
		{
			Path:    "lock",
			Handler: ApiHandler[api.LockIn, api.LockOut](testApi.Lock),
		},
		{
			Path:    "validate",
			Handler: ApiHandler[api.ValidateIn, api.ValidateOut](testApi.Validate),
		},
		{
			Path:    "upload",
			Handler: ApiHandler[api.UploadIn, api.UploadOut](testApi.Upload),
		},
		{
			Path:    "sse",
			Handler: ApiHandler[api.SSEIn, api.SSEOut](testApi.SSE),
		},
		{
			Path:    "sse.req",
			Handler: ApiHandler[api.SSEReqIn, api.SSEReqOut](testApi.SSEReq),
		},
		{
			Path:    "shop.item.rpc",
			Handler: ApiHandler[api.ShopItemRPCIn, api.ShopItemRPCOut](testApi.ShopItemRPC),
		},
		{
			Path:    "error",
			Handler: ApiHandler[api.ErrorIn, api.ErrorOut](testApi.Error)},
		{
			Path:    "token/create",
			Handler: ApiHandler[api.TokenCreateIn, api.TokenCreateOut](testApi.TokenCreate),
		},
		{
			Path:            "token/auth",
			Handler:         ApiHandler[api.TokenAuthIn, api.TokenAuthOut](testApi.TokenAuth),
			HttpMiddlewares: HttpMiddlewares{authMw},
		},
	})

	// add rpc
	rpc := app.RPCServer.Group("test")
	rpc.AddRoutes(RPCRoutes{
		{
			Path:    "get",
			Handler: ApiHandler[api.ShopItemGetIn, api.ShopItemGetOut](testApi.ShopItemGet),
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
