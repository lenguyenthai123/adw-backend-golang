package models

type MailConfig struct {
	MailFrom   string `mapstructure:"MAIL_FROM"`
	MailServer string `mapstructure:"MAIL_SERVER"`
	MailPort   int    `mapstructure:"MAIL_PORT"`
	MailPass   string `mapstructure:"MAIL_PASS"`
}
