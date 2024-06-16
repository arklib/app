package top

import (
	. "app/src"
	. "github.com/arklib/ark"
)

type Api struct {
	*App
}

func New(app *App) *Api {
	api := &Api{app}
	router := app.HttpServer

	// cacheMW := cache.ApiMiddleware(app.Redis, "api", 5*time.Minute)
	top := router.Group("api/top")
	top.AddRoutes(HttpRoutes{
		{
			Title:   "用户创建",
			Path:    "user/create",
			Handler: ApiHandler[UserCreateIn, UserCreateOut](api.UserCreate),
		},
		{
			Title:   "用户登陆",
			Path:    "user/login",
			Handler: ApiHandler[UserLoginIn, UserLoginOut](api.UserLogin),
		},
	})

	topAuth := router.Group("api/top")
	topAuth.AddRoutes(HttpRoutes{
		{
			Title:   "用户获取",
			Path:    "user/get",
			Handler: ApiHandler[UserGetIn, UserGetOut](api.UserGet),
		},
		{
			Title:   "用户搜索",
			Path:    "user/search",
			Handler: ApiHandler[UserSearchIn, UserSearchOut](api.UserSearch),
			// ApiMiddlewares: ApiMiddlewares{cacheMW},
		},
	},
		app.Auth.HttpMiddleware("user"),
	)

	return api
}
