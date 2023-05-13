package teacherportal

import (
	"bytes"
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
	case 4: // /students/{:id}/grades
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if strings.ToLower(pathSegments[3]) != "grades" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.renderGrades(w, req, id)

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

// renderStudent retrieves a single student
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

// Handle the POST request:
// 1) Only allow handling of the POST requests
// 2) Defer function to redirect back to the student page no matter what happens
// 3) Parse the form received by this POST request
// 4) Construct a grade object
// 5) Convert the grade object to a format that can be sent to the Grading Service.
// 6) Request the Grading Service.
// 7) Post the request off to the grading service and check the status code.

// renderGrades adds a new grade.
func (studentsHandler) renderGrades(w http.ResponseWriter, req *http.Request, id int) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer func() {
		w.Header().Add("location", fmt.Sprintf("/students/%v", id))
		w.WriteHeader(http.StatusTemporaryRedirect)
	}()
	title := req.FormValue("Title")
	gradeType := req.FormValue("Type")
	score, err := strconv.ParseFloat(req.FormValue("Score"), 32)
	if err != nil {
		log.Println("Failed to parse score: ", err)
		return
	}
	g := grades.Grade{
		Title: title,
		Type:  grades.GradeType(gradeType),
		Score: float32(score),
	}
	data, err := json.Marshal(g)
	if err != nil {
		log.Println("Failed to convert grade to JSON: ", g, err)
	}
	serviceURL, err := registry.GetProvider(registry.GradingService)
	if err != nil {
		log.Println("Failed to retrieve instance of Grading service.", err)
		return
	}
	res, err := http.Post(fmt.Sprintf("%v/students/%v/grades", serviceURL, id), "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Println("Failed to save grade to Grading Service.", err)
		return
	}
	if res.StatusCode != http.StatusCreated {
		log.Println("Failed to save grade to Grading Service. Status: ", res.StatusCode)
	}
}
