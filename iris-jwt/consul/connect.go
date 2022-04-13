package consul

import (
	"log"

	consulapi "github.com/hashicorp/consul/api"
)

const (
	consulAgentAddress = "124.222.35.20:8500"
)

// 从consul中发现服务
func ConsulFindServer() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = consulAgentAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Println("consul client error : ", err)
	}

	// 获取指定service
	service, _, err := client.Agent().Service("cloud-grpc-server-01", nil)
	if err == nil {
		log.Println(service.Address)
		log.Println(service.Port)
	}

	// 只获取健康的service
	// serviceHealthy, _, err := client.Health().Service("service337", "", true, nil)
	// if err == nil {
	// 	fmt.Println(serviceHealthy[0].Service.Address)
	// }

}
