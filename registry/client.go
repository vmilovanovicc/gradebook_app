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
		return fmt.Errorf("Failed to register service. Registry service responded with code %v.", res.StatusCode)
	}
	return nil
}
