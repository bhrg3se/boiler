package main

import (
	"boiler/migrations"
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
	path := flag.String("c", "/etc/boiler/config", "config file location")
	writeToFile := flag.Bool("f", false, "write logs to file")
	flag.Parse()

	config := parseConfig(*path)
	level, err := logrus.ParseLevel(config.Logging.Level)
	if err != nil {
		level = logrus.ErrorLevel
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)

	if *writeToFile {
		f, err := os.OpenFile("/var/log/boiler.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
		logrus.SetOutput(f)
	}

	store.State = store.NewRealStore(config)

	migrations.Migrate(store.State.DB)

	r := routes.Routes()
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
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic("config file invalid: " + err.Error())
	}

	return config
}
