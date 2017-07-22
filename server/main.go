package main

import (
	proto "github.com/pedrodelgallego/cloud-native-go-kit/proto"
	"golang.org/x/net/context"
	"fmt"
	"github.com/micro/go-micro"
)

type Greeter struct{}

func (*Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	
	fmt.Print("Responding with %s\n", rsp.Greeting)
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("1.0.1"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)
	
	// Init will parse the command line flags. Any flags set will
	// override the above settings. Options defined here will
	// override anything set on the command line.
	service.Init()
	
		// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))
	
	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
