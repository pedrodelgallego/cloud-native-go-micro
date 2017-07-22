package main

import (
	proto "github.com/pedrodelgallego/cloud-native-go-kit/proto"
	"github.com/micro/go-micro"
	"time"
	"context"
	"fmt"
)

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)
	
	service.Init()
	
	greeter := proto.NewGreeterClient("greeter", service.Client())
	callEvery(3*time.Second, greeter, hello)
}

func hello(t time.Time, greeter proto.GreeterClient) {
	resp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{
		Name: "Pedro",
	})

	if err != nil {
		panic(err)
	}
	
	fmt.Print("%S\n", resp.Greeting)
}

func callEvery(d time.Duration, greeter proto.GreeterClient, f func(time.Time, proto.GreeterClient)) {
	for x := range time.Tick(d) {
		f(x, greeter)
	}
}