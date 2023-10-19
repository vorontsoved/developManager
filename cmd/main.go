package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	// if err := godotenv.Load("../.env"); err != nil {
	// 	logrus.Fatalf("error loading env variables: %s", err.Error())
	// }

}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
