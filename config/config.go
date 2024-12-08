package config

import (
	"backend-golang/pkgs/log"
	"backend-golang/utils"
	"fmt"
	"github.com/joho/godotenv"
	"os"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func LoadConfig(model interface{}) interface{} {
	// Load the .env file from the current directory
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		panic(err)
	}

	// Get the CONFIG_PATH from the environment variables
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("CONFIG_PATH is not set in the environment")
	}

	// Load the configuration from the file (you need to implement this function)
	config, err := loadConfigFromFile(configPath)
	if err != nil {
		fmt.Println("Error loading config file:", err)
		panic(err)
	}

	// Configure the decoder to allow weakly typed input and map it to the model
	customConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           model, // Pass the model directly
	}

	// Create a new decoder with the custom configuration
	decoder, err := mapstructure.NewDecoder(customConfig)
	if err != nil {
		fmt.Println("Error creating decoder:", err)
		panic(err)
	}

	// Decode the configuration into the model
	err = decoder.Decode(config)
	if err != nil {
		fmt.Println("Error decoding config:", err)
		panic(err)
	}

	return model
}

// LoadConfigFromFile loads a configuration file from the given path and returns the loaded configuration.
// It expects the path to the configuration file and the name of the configuration file as parameters.
func loadConfigFromFile(configPath string) (config interface{}, err error) {
	var dirPath = utils.GetDirectoryPath(configPath)
	var fileName = utils.GetFileName(configPath)

	viper.AddConfigPath(dirPath)
	viper.SetConfigName(fileName)
	viper.SetConfigType("env")

	// Automatically read in environment variables that match the configuration keys.
	viper.AutomaticEnv()

	// Read in the configuration file.
	err = viper.ReadInConfig()
	if err != nil {
		log.JsonLogger.Error(err.Error())
		panic(err)
	}

	// Unmarshal the configuration file into the given config interface.
	err = viper.Unmarshal(&config)
	if err != nil {
		log.JsonLogger.Error(err.Error())
		panic(err)
	}

	return
}
