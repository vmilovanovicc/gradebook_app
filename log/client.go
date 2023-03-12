package log

import (
	"bytes"
	"fmt"
	"gradebook_app/registry"
	stlog "log"
	"net/http"
)

type clientLogger struct {
	url string
}

// SetClientLogger is used by consumers of this.
func SetClientLogger(serviceURL string, clientService registry.ServiceName) {
	stlog.SetPrefix(fmt.Sprintf("[%v] - ", clientService))
	stlog.SetFlags(0)
	stlog.SetOutput(&clientLogger{url: serviceURL})
}

// Write sends the POST request to the LogService.
func (cl clientLogger) Write(data []byte) (int, error) {
	b := bytes.NewBuffer(data)
	res, err := http.Post(cl.url+"/log", "text/plain", b)
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to send log message. Service responded with %v - %v", res.StatusCode, res.Status)
	}
	return len(data), nil
}
