package registry

import "sync"

// TODO: Define the registry service

const (
	ServerPort  = "3000"
	ServicesURL = "http://localhost" + ServerPort + "/services"
)

type registry struct {
	registrations []Registration
	mutex         *sync.Mutex
}

// Accept a new registration.
// Mutex allows to manipulate the registration slice without causing problems elsewhere i.e. thread safety.
func (r *registry) add(reg Registration) error {
	// allow to manipulate registration slice without casuing problems elsewhere
	r.mutex.Lock()
	r.registrations = append(r.registrations, reg)
	r.mutex.Unlock()
	return nil
}

// TODO: Create registry instance
var reg = registry{
	registrations: make([]Registration, 0),
	mutex:         new(sync.Mutex),
}
