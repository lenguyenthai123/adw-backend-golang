package config

import (
	"log"
	"time"

	"github.com/MicahParks/keyfunc"
)

var JWKS *keyfunc.JWKS

// LoadJWKS tải JWKS từ URL
func LoadJWKS(jwksURL string) {
	var err error
	JWKS, err = keyfunc.Get(jwksURL, keyfunc.Options{
		RefreshInterval: time.Hour,
		RefreshErrorHandler: func(err error) {
			log.Printf("Failed to refresh JWKS: %v", err)
		},
		RefreshUnknownKID: true,
	})
	if err != nil {
		log.Fatalf("Failed to load JWKS: %v", err)
	}
}
