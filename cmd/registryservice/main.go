package registryservice

import (
	"context"
	"fmt"
	"gradebook_app/registry"
	"log"
	"net/http"
)

func main() {
	http.Handle("/services", &registry.RegistryService{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var srv http.Server
	srv.Addr = registry.ServerPort

	// Start the server.
	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()
	// Shut down the server.
	go func() {
		fmt.Println("Registry service started. Press any key to stop.")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("Shutting down the registry service.")

}
