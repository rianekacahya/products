package consul

import (
	"github.com/hashicorp/consul/api"
)

type ServiceRegister struct {
	ID      string
	Name    string
	Tags    []string
	Address string
	Port    int
	Checks  *ServiceCheck
}

type ServiceCheck struct {
	HTTP                           string
	Interval                       string
	Timeout                        string
	DeregisterCriticalServiceAfter string
}

func (sr *ServiceRegister) Register() error {
	registration := new(api.AgentServiceRegistration)
	registration.ID = sr.ID
	registration.Name = sr.Name
	registration.Tags = sr.Tags
	registration.Address = sr.Address
	registration.Port = sr.Port
	registration.Check = new(api.AgentServiceCheck)
	registration.Check.HTTP = sr.Checks.HTTP
	registration.Check.Interval = sr.Checks.Interval
	registration.Check.Timeout = sr.Checks.Timeout
	registration.Check.DeregisterCriticalServiceAfter = sr.Checks.DeregisterCriticalServiceAfter

	if err := client().Agent().ServiceRegister(registration); err != nil {
		return err
	}

	return nil
}
