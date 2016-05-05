package main

import (
	"bytes"
	"os"
	"os/signal"
	"syscall"

	log "github.com/Sirupsen/logrus"
	"github.com/iron-io/titan/common"
	drivercommon "github.com/iron-io/titan/runner/drivers/common"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

type Config struct {
	ApiUrl       string `json:"api_url"`
	Concurrency  int    `json:"concurrency"`
	DriverConfig *drivercommon.Config
	Registries   map[string]*Registry `json:"registries"`
}

// Registry holds auth for a registry
type Registry struct {
	Auth     string `json:"auth"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func InitConfig() *Config {
	config := &Config{}
	err := viper.Unmarshal(config)
	if err != nil {
		log.WithError(err).Fatalln("could not unmarshal registries from config")
	}
	config.ApiUrl = viper.GetString("api_url")
	config.Concurrency = viper.GetInt("concurrency")

	dconfig := &drivercommon.Config{}
	dconfig.JobsDir = viper.GetString("jobs_dir")
	dconfig.Memory = int64(viper.GetInt("memory"))
	dconfig.CPUShares = int64(viper.GetInt("cpu_shares"))
	dconfig.DefaultTimeout = uint(viper.GetInt("timeout"))
	dconfig.Defaults()
	config.DriverConfig = dconfig

	return config
}

// Put generated client inside ./client and call the API package 'titan'
//go:generate swagger generate client -t ./client -f ../jobserver/swagger/api.yml -c titan

func main() {
	viper.SetDefault("concurrency", 5)
	viper.SetEnvPrefix("titan")
	viper.SetDefault("api_url", "http://localhost:8080")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.UnsupportedConfigError); ok {
			log.Infoln("Couldn't read config file, this is fine, it's not required.", err)
			// ignore
		} else {
			log.WithError(err).Fatalln("Error reading config file")
		}
	}
	if os.Getenv("CONFIG") != "" {
		viper.SetConfigType("json")
		err = viper.ReadConfig(bytes.NewBufferString(os.Getenv("CONFIG")))
		if err != nil {
			log.WithError(err).Fatalln("Error reading CONFIG from env")
		}
	}
	config := InitConfig()
	common.SetLogLevel()

	log.Debugf("Config: %+v", config)
	// viper.Debug()

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		sig := <-c
		log.Info("received signal", sig)
		cancel()
	}()

	hostname, err := os.Hostname()
	if err != nil {
		log.WithError(err).Fatal("couldn't resolve hostname")
	}
	l := log.WithFields(log.Fields{
		"hostname": hostname,
	})
	tasker := NewTasker(config, l)

	Run(config, tasker, BoxTime{}, ctx)
}