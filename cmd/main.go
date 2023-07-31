package main

import (
	"authentication-service/cmd/auth"
	"authentication-service/cmd/token"
	"authentication-service/config"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configFile = "config.yaml"
)

var root = &cobra.Command{
	Use:   "server",
	Short: "Short desc",
	Long:  "Long desc",
}

// Init config, log and cmd to run multiple server
func init() {
	loadEnv()
	configs := loadConfig()
	logger := initLogger()
	authServer := auth.NewServer(configs, logger)
	tokenServer := token.NewServer(configs, logger)
	root.AddCommand(authServer, tokenServer)
}

func main() {
	if err := root.Execute(); err != nil {
		log.Fatalf("Service start failed with error: %v\n", err)
	}
}

func initLogger() *logrus.Logger {
	// Can use other lib
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	return logger
}

func loadEnv() {
	_ = godotenv.Load()
}

func loadConfig() *config.AppConfigs {
	viper.SetConfigType("yaml")

	b, err := os.ReadFile(configFile)

	if err != nil {
		fmt.Printf("Load config file error: %v\n", err)
	}
	expandEnv := os.ExpandEnv(string(b))
	configReader := strings.NewReader(expandEnv)

	viper.AutomaticEnv()
	if err := viper.ReadConfig(configReader); err != nil {
		log.Fatalf("Read config error: %v\n", err)
	}

	configs := config.AppConfigs{}
	if err := viper.Unmarshal(&configs); err != nil {
		log.Fatalf("Unmarshal config error: %v\n", err)
	}

	return &configs
}
