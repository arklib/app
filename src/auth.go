package src

import (
	"log"
	"time"

	"github.com/arklib/ark/auth"
)

func initAuth(app *App) *auth.Auth {
	app.Logger.Info("[app] init auth")
	config := new(struct {
		Expires     time.Duration `default:"72h"`
		SecretKey   []byte
		TokenLookup string `default:"header: Authorization"`
	})

	err := app.BindConfig("auth", config)
	if err != nil {
		log.Fatalf("auth config: %v", err)
	}

	newAuth, err := auth.New(config.SecretKey, config.Expires, config.TokenLookup)
	if err != nil {
		log.Fatalf("auth: %v", err)
	}
	return newAuth
}
