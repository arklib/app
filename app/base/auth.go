package base

import (
	"log"
	"time"

	"github.com/arklib/ark/auth"
)

func (base *Base) initAuth() {
	base.Logger.Info("[app] Init auth")
	config := new(struct {
		Expires     time.Duration `default:"72h"`
		SecretKey   []byte
		TokenLookup string `default:"header: Authorization"`
	})

	err := base.BindConfig("auth", config)
	if err != nil {
		log.Fatalf("auth config: %v", err)
	}

	authInst, err := auth.New(config.SecretKey, config.Expires, config.TokenLookup)
	if err != nil {
		log.Fatalf("auth: %v", err)
	}
	base.Auth = authInst
}
