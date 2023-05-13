package teacherportal

import (
	"encoding/json"
	"fmt"
	"gradebook_app/grades"
	"gradebook_app/registry"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type studentsHandler struct{}

// RegisterHandlers registers the request handlers.
func RegisterHandlers() {
	http.Handle("/", http.RedirectHandler("/students", http.StatusPermanentRedirect))
	h := new(studentsHandler)

	http.Handle("/students", h)
	http.Handle("/students/", h)
}

// ServeHTTP handles various redirection use cases.
func (sh studentsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	pathSegments := strings.Split(req.URL.Path, "/")
	switch len(pathSegments) {
	case 2: // /students
		sh.renderStudents(w, req)
	case 3: // /students/{:id}
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.renderStudent(w, req, id)

	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

// renderStudents retrieves all students
func (studentsHandler) renderStudents(w http.ResponseWriter, req *http.Request) {
	var err error
	// Centralized error handling.
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error retrieving students: ", err)
		}
	}()
	// API Gateway
	serviceURL, err := registry.GetProvider(registry.GradingService)
	if err != nil {
		return
	}
	res, err := http.Get(serviceURL + "/students")
	if err != nil {
		return
	}
	var s grades.Students
	err = json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		return
	}
	rootTemplate.Lookup("students.gohtml").Execute(w, s)

}

// renderStudent
func (studentsHandler) renderStudent(w http.ResponseWriter, req *http.Request, id int) {
	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error retrieving a student: ", err)
		}
	}()

	serviceURL, err := registry.GetProvider(registry.GradingService)
	if err != nil {
		return
	}
	res, err := http.Get(fmt.Sprintf("%v/students/%v", serviceURL, id))
	if err != nil {
		return
	}
	var s grades.Student
	err = json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		return
	}

	rootTemplate.Lookup("student.gohtml").Execute(w, s)
}
