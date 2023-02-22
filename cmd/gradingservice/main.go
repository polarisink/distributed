package main

import (
	"context"
	"fmt"
	"github.com/polarisink/distributed/grades"
	"github.com/polarisink/distributed/registry"
	"github.com/polarisink/distributed/service"
	stlog "log"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration{
		ServiceName: registry.GradingService,
		ServiceURL:  serviceAddress,
		//RequiredServices: []registry.ServiceName{registry.LogService},
		//ServiceUpdateURL: serviceAddress + "/services",
		//HeartbeatURL:     serviceAddress + "/heartbeat",
	}
	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	/*if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at: %s\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}*/

	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
