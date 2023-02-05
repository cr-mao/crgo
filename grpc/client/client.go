package client

import (
	"context"
	"crgo/grpc/biz/helloworld"
	"crgo/infra/errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
)

func SayHello(client helloworld.GreeterClient) error {
	resp, err := client.SayHello(context.Background(), &helloworld.HelloRequest{
		Name: "crmao",
	})
	if err != nil {

		s, ok := status.FromError(err)

		if !ok {
			fmt.Println("is not stardand grpc error")
			return err
		}
		fmt.Println(111)
		fmt.Println(s.Code())

		//grpc 错误转为内部错误
		e := errors.FromGrpcError(err)
		Coder := errors.ParseCoder(e)
		fmt.Println(2222)
		fmt.Println(Coder.Code())
		fmt.Println(Coder.HTTPStatus())

		return err
	}

	log.Printf("client.Sayhello resp:%s", resp.Message)
	return nil
}

func Do() error {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
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
