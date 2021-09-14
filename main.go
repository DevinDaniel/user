package main

import (
	"usetemp/handler"
	pb "usetemp/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("usetemp"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterUsetempHandler(srv.Server(), new(handler.Usetemp))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
