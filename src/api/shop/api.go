package shop

import (
	. "github.com/arklib/ark"

	. "app/src"
)

type Api struct {
	*App
}

func New(app *App) *Api {
	api := &Api{app}
	router := app.RPCServer

	item := router.Group("item")
	item.AddRoutes(RPCRoutes{
		{
			Path:    "get",
			Handler: ApiHandler[ItemGetIn, ItemGetOut](api.ItemGet),
			ApiMiddlewares: ApiMiddlewares{
				func(payload *ApiPayload) error {
					err := payload.Next()
					if err != nil {
						return err
					}

					// dump
					app.Dump(payload.Path, payload.In, payload.Out)

					// rewrite output
					out := payload.Out.(*ItemGetOut)
					out.Title = "(proxy) " + out.Title
					return nil
				},
			},
		},
	})

	return api
}
