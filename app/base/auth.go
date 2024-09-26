package base

import (
	"log"

	"github.com/arklib/ark/auth"
)

func (base *Base) initAuth() *auth.Auth {
	if base.Auth != nil {
		return base.Auth
	}

	c := new(struct {
		Expire      int64 `default:"86400"`
		SecretKey   string
		TokenLookup string `default:"header: Authorization"`
	})

	err := base.BindConfig("auth", c)
	if err != nil {
		log.Fatalf("auth c: %v", err)
	}

	authInst, err := auth.New(c.SecretKey, c.Expire, c.TokenLookup)
	if err != nil {
		log.Fatalf("auth: %v", err)
	}

	base.Auth = authInst
	return authInst
}
