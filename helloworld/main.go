package main

import (
	"github.com/ansg191/test-micro/helloworld/handler"
	pb "github.com/ansg191/test-micro/helloworld/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("helloworld"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterHelloworldHandler(srv.Server(), new(handler.Helloworld))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
