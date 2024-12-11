package models

type KeycloakConfig struct {
	JwksURL string `mapstructure:"KEYCLOAK_JWKS_URL"`
}
