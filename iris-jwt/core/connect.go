package core

import (
	"fmt"
	"iris-jwt/config"
	"log"

	consulapi "github.com/hashicorp/consul/api"
)

var (
	consulAgentAddress = config.Configuration.ConsulAddr
	ServiceMap         map[string]string
)

func init() {
	ServiceMap = make(map[string]string)
}

// 从consul中发现服务
func ConsulFindServer(serverName string) {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = consulAgentAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Println("consul client error : ", err)
	}

	// 获取指定service
	service, _, err := client.Agent().Service(serverName, nil)
	if err != nil {
		log.Println("failed to get service")
	}

	ServiceMap[serverName] = service.Address + ":" + fmt.Sprintf("%d", service.Port)
	// 只获取健康的service
	// serviceHealthy, _, err := client.Health().Service("service337", "", true, nil)
	// if err == nil {
	// 	log.Println(serviceHealthy[0].Service.Address)
	// }
}
