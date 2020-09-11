package consul

import (
	"github.com/hashicorp/consul/api"
	"log"
	"os"
	"sync"
)

const (
	ConsulEnvironment = "CONSUL_HTTP_ADDR"
)

var (
	connection *api.Client
	mutex      sync.Once
)

func client() *api.Client {
	mutex.Do(func() {
		connection = newConsul()
	})

	return connection
}

func newConsul() *api.Client {
	consul, err := api.NewClient(&api.Config{
		Address: os.Getenv(ConsulEnvironment),
	})
	if err != nil {
		log.Fatalf("got an error while connecting to consul server, error: %s", err)
	}

	return consul
}
