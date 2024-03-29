package grades

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

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
		return nil, fmt.Errorf("failed to serialize students: %q", err)
	}
	return b.Bytes(), nil
}

func (sh studentsHandler) getAll(w http.ResponseWriter, req *http.Request) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()
	data, err := sh.toJSON(students)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (sh studentsHandler) getOne(w http.ResponseWriter, req *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()
	student, err := students.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}

	data, err := sh.toJSON(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(fmt.Errorf("failed to serialize student: %q", err))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (sh studentsHandler) addGrade(w http.ResponseWriter, req *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()
	// Return a student.
	student, err := students.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}
	// Decode grades from the POST body.
	var g Grade
	dec := json.NewDecoder(req.Body)
	err1 := dec.Decode(&g)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err1)
		return
	}
	// Add a new grade onto the student's grades slice.
	student.Grades = append(student.Grades, g)
	w.WriteHeader(http.StatusCreated)

	data, err2 := sh.toJSON(g)
	if err2 != nil {
		log.Println(err2)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
