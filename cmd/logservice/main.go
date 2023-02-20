package main

import (
	"context"
	"fmt"
	"github.com/polarisink/distributed/log"
	"github.com/polarisink/distributed/registry"
	"github.com/polarisink/distributed/service"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "8000"
	r := registry.Registration{
		ServiceName: "Log Service",
		ServiceURL:  fmt.Sprintf("http://%s:%s", host, port),
	}
	ctx, err := service.Start(context.Background(), host, port, r, log.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()

	fmt.Println("Shutting down service")
}
