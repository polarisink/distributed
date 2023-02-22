package grades

import (
	"fmt"
	"sync"
)

const (
	GradeQuiz = GradeType("Quiz")
	GradeTest = GradeType("Test")
	GradeExam = GradeType("Exam")
)

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}

type GradeType string

type Students []Student

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Grades    []Grade
}

var (
	students      Students
	studentsMutex sync.Mutex
)

func (ss Students) GetById(id int) (*Student, error) {
	for i := range ss {
		if ss[i].ID == id {
			return &ss[i], nil
		}
	}
	return nil, fmt.Errorf("student with ID %d not found", id)
}

func (ss Students) GetByID(id int) (*Student, error) {
	for i := range ss {
		if ss[i].ID == id {
			return &ss[i], nil
		}
	}

	return nil, fmt.Errorf("student with ID %d not found", id)
}

func (s Student) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score
	}
	return result / float32(len(s.Grades))
}
