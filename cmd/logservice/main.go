package main

import (
	"context"
	"fmt"
	"gradebook_app/log"
	"gradebook_app/registry"
	"gradebook_app/service"
	stlog "log"
)

func main() {
	log.Run("./app.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	// Create registration object.
	var r registry.Registration
	r.ServiceName = registry.LogService
	r.ServiceURL = serviceAddress
	r.RequiredServices = make([]registry.ServiceName, 0)
	r.ServiceUpdateURL = r.ServiceURL + "/services"

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
	)

	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down log service.")

}
