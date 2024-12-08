package models

type DBConfig struct {
	BuildEnv                    string `mapstructure:"BUILD_ENV"`
	DBHost                      string `mapstructure:"DB_HOST"`
	DBPort                      string `mapstructure:"DB_PORT"`
	DBName                      string `mapstructure:"DB_NAME"`
	DBUserName                  string `mapstructure:"DB_USERNAME"`
	DBPassword                  string `mapstructure:"DB_PASSWORD"`
	SSLCertAuthorityCertificate string `mapstructure:"SSL_CERT_AUTH"`
	MaxOpenConnections          int    `mapstructure:"MAX_OPEN_CONNECTIONS"`
	MaxIdleConnections          int    `mapstructure:"MAX_IDLE_CONNECTIONS"`
	ConnectionMaxIdleTime       int    `mapstructure:"CONNECTION_MAX_IDLE_TIME"`
	ConnectionMaxLifeTime       int    `mapstructure:"CONNECTION_MAX_LIFE_TIME"`
	ConnectionTimeout           int    `mapstructure:"CONNECTION_TIMEOUT"`
}
