package grades

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// define server logic
func RegisterHandlers() {
	handler := new(studentsHandler)
	http.Handle("/students", handler)
	http.Handle("/students/", handler)
}

type studentsHandler struct{}

// ServeHttp handles a couple of patterns:
// /students - entire class ---> 0 empty, 1 students
// /students/{id} - a single student's record ---> 0 empty, 1 students, 2 id
// /students/{id}/grades - a single student's grades ---> 0 empty, 1 students, 2 id, 3 grades
func (sh studentsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	pathSegments := strings.Split(req.URL.Path, "/")
	switch len(pathSegments) {
	case 2:
		sh.getAll(w, req)
	case 3:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.getOne(w, req, id)
	case 4:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.addGrade(w, req, id)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// toJSON converts object to JSON
func (sh studentsHandler) toJSON(obj interface{}) ([]byte, error) {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	err := enc.Encode(obj)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize students: &q", err)
	}
	return b.Bytes(), nil
}
