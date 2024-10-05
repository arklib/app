package app

import (
	"log"

	"github.com/arklib/ark/auth"
)

func (app *App) initAuth() {
	var c struct {
		Expire      int64 `default:"86400"`
		SecretKey   string
		TokenLookup string `default:"header: Authorization"`
	}

	err := app.BindConfig("auth", &c)
	if err != nil {
		log.Fatalf("auth config: %v", err)
	}

	authInst, err := auth.New(c.SecretKey, c.Expire, c.TokenLookup)
	if err != nil {
		log.Fatalf("auth: %v", err)
	}
	app.Auth = authInst
}
