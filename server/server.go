package server

import (
	"boiler/models"
	"boiler/routes"
	"boiler/store"
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"path/filepath"
)

func StartServer() {
	writeToFile := flag.Bool("f", false, "write logs to file")
	flag.Parse()

	// parse config file
	config := parseConfig("/etc/boiler/")

	//set log level
	level, err := logrus.ParseLevel(config.Logging.Level)
	if err != nil {
		level = logrus.ErrorLevel
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)

	// create log file if enabled
	if *writeToFile {
		f, err := os.OpenFile("/var/log/boiler.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
		logrus.SetOutput(f)
	}

	// create database connections and other dependencies
	store.State = store.NewRealStore(config)

	//initialize API routes
	r := routes.Routes()

	// start server
	logrus.Infof("starting server at: %s", config.Server.Listen)
	logrus.Error(http.ListenAndServe(config.Server.Listen, r))
}

// parseConfig uses viper to parse config file.
func parseConfig(path string) models.Config {
	var config models.Config
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic("config file not found in " + filepath.Join(path))
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(absPath)

	if err = viper.ReadInConfig(); err != nil {
		logrus.Fatalf("could not read config file: %v",err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		logrus.Fatalf("config file invalid: %v" , err)
	}

	return config
}
