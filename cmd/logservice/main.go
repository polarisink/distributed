package main

import (
	"context"
	"fmt"
	"github.com/polarisink/distributed/log"
	"github.com/polarisink/distributed/registry"
	"github.com/polarisink/distributed/service"
	"github.com/polarisink/distributed/tools"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "8000"
	serviceAddress := tools.ParseAddress(host, port)

	r := registry.Registration{
		ServiceName:      registry.LogService,
		ServiceURL:       serviceAddress,
		RequiredServices: make([]registry.ServiceName, 0),
		ServiceUpdateURL: serviceAddress + "/services",
		HeartbeatURL:     serviceAddress + "/heartbeat",
	}
	ctx, err := service.Start(context.Background(), host, port, r, log.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()

	fmt.Println("Shutting down service")
}
