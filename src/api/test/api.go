package test

import (
	. "github.com/arklib/ark"

	"app/src"
)

type Api struct {
	*src.App
}

func New(app *src.App) *Api {
	api := &Api{app}
	router := app.HttpServer

	// auth middleware
	authMw := app.Auth.HttpMiddleware("user")

	test := router.Group("api/test")
	test.AddRoutes(HttpRoutes{
		{
			Method:  "GET",
			Path:    "ping",
			Handler: ApiHandler[PingIn, PingOut](api.Ping),
			ApiMiddlewares: ApiMiddlewares{
				func(p *ApiPayload) error {
					err := p.Next()
					if err != nil {
						return err
					}
					out := p.Out.(*PingOut)
					out.Message = "(proxy) " + out.Message
					return nil
				},
			},
		},
		{
			Path:    "validate",
			Handler: ApiHandler[ValidateIn, ValidateOut](api.Validate),
		},
		{
			Path:    "upload",
			Handler: ApiHandler[UploadIn, UploadOut](api.Upload),
		},
		{
			Path:    "sse",
			Handler: ApiHandler[SSEIn, SSEOut](api.SSE),
		},
		{
			Path:    "sse.req",
			Handler: ApiHandler[SSEReqIn, SSEReqOut](api.SSEReq),
		},
		{
			Path:    "rpc",
			Handler: ApiHandler[RPCIn, RPCOut](api.RPC),
		},
		{
			Path:    "error",
			Handler: ApiHandler[ErrorIn, ErrorOut](api.Error),
		},
		{
			Method:  "GET",
			Path:    "cache",
			Handler: ApiHandler[CacheIn, CacheOut](api.Cache),
		},
		{
			Method:  "GET",
			Path:    "lock",
			Handler: ApiHandler[LockIn, LockOut](api.Lock),
		},
		{
			Path:    "token/create",
			Handler: ApiHandler[TokenCreateIn, TokenCreateOut](api.TokenCreate),
		},
		{
			Path:        "token/auth",
			Middlewares: HttpMiddlewares{authMw},
			Handler:     ApiHandler[TokenAuthIn, TokenAuthOut](api.TokenAuth),
		},
	})
	return api
}
