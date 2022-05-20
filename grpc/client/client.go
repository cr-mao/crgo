package client

import (
	"context"
	"crgo/grpc/biz/helloworld"
	"google.golang.org/grpc"
	"log"
)

func SayHello(client helloworld.GreeterClient) error {
	resp, err := client.SayHello(context.Background(), &helloworld.HelloRequest{
		Name: "crmao",
	})
	if err != nil {
		return err
	}
	log.Printf("client.Sayhello resp:%s", resp.Message)
	return nil
}

func Do() error {
	conn, err := grpc.Dial(":8001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("client conn err :%v", err)
		return err
	}
	defer conn.Close()
	client := helloworld.NewGreeterClient(conn)
	if err := SayHello(client); err != nil {
		log.Fatalf("sayhello err :%v", err)
		return err
	}
	return nil
}
