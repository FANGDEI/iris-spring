package main

import (
	"context"
	"iris-jwt/core"
	"iris-jwt/proto/getcode"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	core.ConsulFindServer("cloud-grpc-server-01")
	var ADDRESS = core.ServiceMap["cloud-grpc-server-01"]
	println(ADDRESS)
	// 建立连接
	conn, err := grpc.Dial(ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("failed to connect: ", err)
	}
	defer conn.Close()
	c := getcode.NewGetCodeClient(conn)

	r, err := c.GetCode(context.Background(), &getcode.CodeRequest{Email: "1639579565@qq.com"})
	if err != nil {
		log.Println("failed to send email: ", err)
		return
	}
	log.Printf("The response message: %s", r.Message)
}
