package models

type S3Config struct {
	S3AccessKey  string `mapstructure:"S3_ACCESS_KEY"`
	S3SecretKey  string `mapstructure:"S3_SECRET_KEY"`
	S3Region     string `mapstructure:"S3_REGION"`
	S3Bucket     string `mapstructure:"S3_BUCKET"`
	S3CloudFront string `mapstructure:"S3_CLOUDFRONT"`
}
