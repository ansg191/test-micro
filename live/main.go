package main

import (
	"github.com/ansg191/test-micro/live/handler"
	pb "github.com/ansg191/test-micro/live/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("live"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterLiveHandler(srv.Server(), new(handler.Live))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
