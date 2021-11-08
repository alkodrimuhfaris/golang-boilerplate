package viperConfig

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type ConfigInterface interface {
	SetConfig() Config
}

func (c *Config) SetConfig() Config {
	err := godotenv.Load("./.env")
	if err != nil {
		if pathErr, ok := err.(*os.PathError); ok {
			if pathErr.Op == "open" {
				fmt.Println("Could not open configuration file .env, ignored")
			} else {
				log.Fatalf("Error loading .env file, %v", err)
			}
		} else {
			log.Fatalf("Error loading .env file, %v", err)
		}
	}

	viper.BindEnv("port", "PORT")
	viper.BindEnv("base_url", "BASE_URL")
	viper.BindEnv("keycloak_base_url", "KEYCLOAK_BASE_URL")
	viper.BindEnv("local_dev", "LOCAL_DEV")

	return Config{
		Port:            viper.GetString("port"),
		BaseURL:         viper.GetString("base_url"),
		KeycloakBaseURL: viper.GetString("keycloak_base_url"),
		LocalDev:        viper.GetString("local_dev"),
	}
}
