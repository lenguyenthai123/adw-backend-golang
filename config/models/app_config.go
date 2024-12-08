package models

type AppConfig struct {
	BuildEnv       string `mapstructure:"BUILD_ENV"`
	Version        string `mapstructure:"VERSION"`
	ServiceName    string `mapstructure:"SERVICE_NAME"`
	ServiceHost    string `mapstructure:"SERVICE_HOST"`
	ServicePort    int    `mapstructure:"SERVICE_PORT"`
	ServiceTimeout int    `mapstructure:"SERVICE_TIMEOUT"`
}
