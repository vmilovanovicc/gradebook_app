package log

import (
	"io"
	stlog "log"
	"net/http"
	"os"
)

// create Logger to handle logging for the app
var log *stlog.Logger

type fileLog string

// custom writer to handle the actual writing to the filesystem
func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

// point logger
func Run(destination string) {
	log = stlog.New(fileLog(destination), "", stlog.LstdFlags)
}

// register http endpoints
func RegisterHandlers() {
	http.HandleFunc("/log", func(writer http.ResponseWriter, request *http.Request) {
		msg, err := io.ReadAll(request.Body)
		if err != nil || len(msg) == 0 {
			writer.WriteHeader(http.StatusBadRequest)
		}
		write(string(msg))
	})
}

func write(message string) {
	log.Printf("%v\n", message)
}
