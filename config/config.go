package config

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var envFileName = "./.env"

func init() {
	viper.SetConfigFile(envFileName)
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Connont find config file, %s", err)
	}
}

func NewConfig() (defConfig *Config, err error) {
	defConfig = &Config{}
	appEnv := os.Getenv(`APP_ENV`)
	appGRPCPort, _ := strconv.Atoi(os.Getenv(`APP_GRPC_PORT`))
	appHTTPPort, _ := strconv.Atoi(os.Getenv(`APP_HTTP_PORT`))
	debugString := os.Getenv(`APP_DEBUG`)
	debug := false

	if debugString == "true" {
		debug = true
	}

	dbHost := os.Getenv(`DB_HOST`)
	dbPort, _ := strconv.Atoi(os.Getenv(`DB_PORT`))
	dbUser := os.Getenv(`DB_USER`)
	dbPassword := os.Getenv(`DB_PASS`)
	dbName := os.Getenv(`DB_NAME`)
	driverName := os.Getenv(`DB_DRIVER_NAME`)

	if appEnv == "" || appGRPCPort == 0 || appHTTPPort == 0 {
		err = fmt.Errorf("[CONFIG][Critical] Please check section APP on %s", envFileName)
		return
	}

	defConfig.AppEnv = appEnv
	defConfig.AppGRPCPort = appGRPCPort
	defConfig.AppHTTPPort = appHTTPPort
	defConfig.Debug = debug

	if dbHost == "" || dbPort == 0 || dbUser == "" || dbName == "" || driverName == "" {
		err = fmt.Errorf("[CONFIG][Critical] Please check section DB on %s", envFileName)
		return
	}

	dbConfig := &DB{
		Host:       dbHost,
		Port:       dbPort,
		Username:   dbUser,
		Password:   dbPassword,
		Name:       dbName,
		DriverName: driverName,
	}

	defConfig.DB = dbConfig

	return defConfig, nil
}
