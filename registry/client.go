package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Allow web service (e.g. log service) to register itself with the service registry.
// Buffer holds data for the POST request.
func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		return err
	}
	res, err := http.Post(ServicesURL, "application/json", buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to register service. Registry service responded with code %v", res.StatusCode)
	}
	return nil
}

// Allow web service to deregister itself with the service registry.
func ShutdownService(serviceURL string) error {
	req, err := http.NewRequest(http.MethodDelete, ServicesURL, bytes.NewBuffer([]byte(serviceURL)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "text/plain")
	// Send delete request to registry service.
	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to deregister service. Registry service responded with code %v", res.StatusCode)
	}
	return err
}
