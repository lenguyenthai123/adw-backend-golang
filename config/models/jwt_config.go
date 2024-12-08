package models

type JWTConfig struct {
	JWTSecretKey       string `mapstructure:"JWT_SECRET_KEY"`
	AccessTokenExpiry  int    `mapstructure:"ACCESS_TOKEN_EXPIRY"`
	RefreshTokenExpiry int    `mapstructure:"REFRESH_TOKEN_EXPIRY"`
}
