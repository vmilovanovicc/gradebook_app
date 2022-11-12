package grades

import (
	"fmt"
	"sync"
)

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Grades    []Grade
}

// Students - Modeling with an in-memory collection.
// The app does not have a backing database.
type Students []Student

var (
	students      Students
	studentsMutex sync.Mutex
)

type GradeType string

const (
	GradeTest     = GradeType("Test")
	GradeHomework = GradeType("Homework")
	GradeQuiz     = GradeType("Quiz")
)

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}

// Average calculates the grade average.
func (s Student) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score
	}
	return result / float32(len(s.Grades))
}

// GetByID returns a student if present in a collection.
func (s Students) GetByID(id int) (*Student, error) {
	for i := range s {
		if s[i].ID == id {
			return &s[i], nil
		}
	}
	return nil, fmt.Errorf("student with ID: %v not found", id)
}
