package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

// go get -u github.com/ilyakaznacheev/cleanenv

// go mod tidy

func MustLoad() *Config {

	var configPath string

	// We need to provide this config path config\local.yml and this file is anywhere when we deploy on production.

	// To get that file value we use various methods

	// 1. By helping of enviromental varaible.
	configPath = os.Getenv("CONFIG_PATH")

	// Now we need to check that is this ` if configPath == "" ` empty string means that file not pass properly.
	if configPath == "" {
		// If configPath is empty "", then we try another method.
		// 2.Method is to Check in arguments and flag
		// When we run our program like this `go run cmd/students-api/main.go -config-path xyz`
		// SO if configPath not find under env file than we check here in argument

		flags := flag.String("config", "", "path to the configuration file")
		// Here we parse the flag value to configPath.
		// The command-line arguments passed to the program and fills in the values for any registered flags
		flag.Parse()

		configPath = *flags

		// If, after checking both the environment variable and the command-line flag, configPath is still empty, we have no way to find the config file.
		// After this if we again not get the configPath then
		// 3.Method then throw error

		if configPath == "" {
			log.Fatal("Config path is not set")
		}

	}

	// If we Not get any file from that file, then we need to check the file path.
	// 4.Method

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Cann't read config file: %s", err.Error())
	}

	return &cfg

}
