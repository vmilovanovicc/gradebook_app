package main

import (
	"context"
	"fmt"
	"gradebook_app/log"
	"gradebook_app/service"
	stlog "log"
)

func main() {
	log.Run("./app.log")
	host, port := "localhost", "4000"
	ctx, err := service.Start(
		context.Background(),
		"Log Service",
		host,
		port,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down log service.")

}
