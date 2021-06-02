package main

import (
	"log"

	handler "supply/handler"
	pb "supply/proto"

	"github.com/micro/micro/v3/service"
)

func main() {

	// Create a new service. Optionally include some options here.
	service := service.New(
		service.Name("supply"),
		service.Version("latest"),
	)

	// Init will parse the command line flags.
	//service.Init()
	h := &handler.Handler{}

	// Register handler
	if err := pb.RegisterSupplyHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	// Run the server
	if err := service.Run(); err != nil {
		log.Panic(err)
	}

}
