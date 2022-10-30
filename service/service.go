package service

import (
	"context"
	"fmt"
	"gradebook_app/registry"
	"log"
	"net/http"
)

// Start the web service.
func Start(ctx context.Context, host, port string, reg registry.Registration, registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()
	ctx = startService(ctx, reg.ServiceName, host, port)

	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	var srv http.Server
	srv.Addr = ":" + port

	// Start the server.
	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	// Shut down the server.
	go func() {
		fmt.Printf("%v started. Press any key to stop.\n", serviceName)
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()
	return ctx
}
