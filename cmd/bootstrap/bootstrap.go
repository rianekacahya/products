package bootstrap

import (
	"database/sql"
	"fmt"
	"log"
	"service/pkg/consul"
	"service/pkg/postgres"
	"service/pkg/reconf"
)

func RegistrationService(){
	service := new(consul.ServiceRegister)
	service.ID = fmt.Sprintf("%v-%v", reconf.Config().GetString("service_name"), startedAt.Unix())
	service.Name = reconf.Config().GetString("service_name")
	service.Address = "host.docker.internal"
	service.Port = 8080
	service.Checks = new(consul.ServiceCheck)
	service.Checks.HTTP = "http://host.docker.internal:8080/v1/healthcheck"
	service.Checks.Interval = "5s"
	service.Checks.Timeout = "5s"
	service.Checks.DeregisterCriticalServiceAfter = "5s"

	err := service.Register()
	if err != nil {
		log.Fatalf("got an error while registration service to consul, error: %s", err)
	}
}

func InitPostgresRead() *sql.DB {
	db, err := postgres.New(
		reconf.Config().GetString("postgres.read.dsn"),
		reconf.Config().GetInt("postgres.read.max_open"),
		reconf.Config().GetInt("postgres.read.max_idle"),
		reconf.Config().GetInt("postgres.read.timeout"),
	)
	 if err != nil {
	 	log.Fatalf("got an error while connecting database server read, error: %s", err)
	 }

	return db
}

func InitPostgresWrite() *sql.DB {
	db, err := postgres.New(
		reconf.Config().GetString("postgres.write.dsn"),
		reconf.Config().GetInt("postgres.write.max_open"),
		reconf.Config().GetInt("postgres.write.max_idle"),
		reconf.Config().GetInt("postgres.write.timeout"),
	)
	if err != nil {
		log.Fatalf("got an error while connecting database server write, error: %s", err)
	}

	return db
}