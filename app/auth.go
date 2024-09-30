package app

import (
	"log"

	"github.com/arklib/ark/auth"
)

func (app *App) initAuth() {
	c := new(struct {
		Expire      int64 `default:"86400"`
		SecretKey   string
		TokenLookup string `default:"header: Authorization"`
	})

	err := app.BindConfig("auth", c)
	if err != nil {
		log.Fatalf("auth c: %v", err)
	}

	inst, err := auth.New(c.SecretKey, c.Expire, c.TokenLookup)
	if err != nil {
		log.Fatalf("auth: %v", err)
	}
	app.Logger.Debug("[app] init auth")

	app.Auth = inst
}
