package main

import (
	"context"
	"fmt"
	_ "github.com/polarisink/distributed/config"
	"github.com/polarisink/distributed/registry"
	"log"
	"net/http"
)

func main() {
	registry.SetupRegistryService()
	http.Handle("/services", &registry.RegistryService{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("Registry service started. Press any key to stop.")
		var s string
		_, err := fmt.Scanln(&s)
		if err != nil {
			return
		}
		err = srv.Shutdown(ctx)
		if err != nil {
			return
		}
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("Shutting down registry service...")
}
