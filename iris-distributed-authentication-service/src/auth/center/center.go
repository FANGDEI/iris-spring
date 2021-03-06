package center

import (
	consul "github.com/hashicorp/consul/api"
	"log"
)

var client *consul.Client

func init() {
	var err error
	config := consul.DefaultConfig()
	config.Address = consulAddr
	client, err = consul.NewClient(config)
	if err != nil {
		log.Fatalln("[INIT CONSUL]", err)
	}
}

func GetValue(key string) ([]byte, error) {
	res, _, err := client.KV().Get(key, nil)
	if err != nil {
		return nil, err
	}
	return res.Value, nil
}

func Register(reg consul.AgentServiceRegistration) error {
	agent := client.Agent()

	return agent.ServiceRegister(&reg)
}
