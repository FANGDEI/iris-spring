package handlers

import (
	"fmt"
	"iris-jwt/constant"
	"log"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/kataras/iris/v12"
)

func AdminService(ctx iris.Context) {
	ctx.HTML("<h1> this is a service for admin </h1>")
}

func RegisterAdminService() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = constant.CONSUL_ADDRESS
	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	// 创建注册到consul的服务到
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "1"
	registration.Name = "IrisService"
	registration.Port = 7777
	registration.Tags = []string{"adminService"}
	registration.Address = "110.42.159.47"

	// 增加consul健康检查回调函数
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d/check", registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalln(err)
	}
}
