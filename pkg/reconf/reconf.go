package reconf

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"log"
	"os"
	"sync"
	"time"
)

const (
	ConsulEnvironment = "CONSUL_HTTP_ADDR"
	ConsulConfigKey   = "CONSUL_CONFIG_KEY"
)

var (
	configuration *viper.Viper
	mutex         sync.Once
)

func Config() *viper.Viper {
	mutex.Do(func() {
		configuration = new()
	})

	return configuration
}

func Watch(ctx context.Context){
	for {
		// delay after each request
		time.Sleep(time.Second * 5)

		select {
		case <-ctx.Done():
			return
		default:
			err := Config().WatchRemoteConfig()
			if err != nil {
				log.Printf("unable to read remote config: %v\n", err)
				continue
			}
		}
	}
}

func new() *viper.Viper {
	// first lets load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("got an error while loading ENV, error: %s", err)
	}

	// init configuration
	configuration := viper.New()
	configuration.SetConfigType("json")
	err := configuration.AddRemoteProvider("consul", os.Getenv(ConsulEnvironment), os.Getenv(ConsulConfigKey))
	if err != nil {
		log.Fatalf("got an error while connecting consul config, error: %s", err)
	}

	// first time reading remote config
	if err := configuration.ReadRemoteConfig(); err != nil {
		log.Fatalf("got an error while reading consul config, error: %s", err)
	}

	return configuration
}
