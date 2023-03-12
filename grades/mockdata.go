package grades

func init() {
	students = []Student{
		Student{
			ID:        1,
			FirstName: "Petar",
			LastName:  "Petrovic",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				Grade{
					Title: "Homework 1",
					Type:  GradeHomework,
					Score: 94,
				},
				Grade{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 88,
				},
			},
		},
		Student{
			ID:        2,
			FirstName: "Ivan",
			LastName:  "Ivanovic",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 100,
				},
				Grade{
					Title: "Homework 1",
					Type:  GradeHomework,
					Score: 100,
				},
				Grade{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 88,
				},
			},
		},
		Student{
			ID:        3,
			FirstName: "Marko",
			LastName:  "Markovic",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 77,
				},
				Grade{
					Title: "Homework 1",
					Type:  GradeHomework,
					Score: 0,
				},
				Grade{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 65,
				},
			},
		},
		Student{
			ID:        4,
			FirstName: "Nikola",
			LastName:  "Nikolic",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 88,
				},
				Grade{
					Title: "Homework 1",
					Type:  GradeHomework,
					Score: 93,
				},
				Grade{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 84,
				},
			},
		},
		Student{
			ID:        5,
			FirstName: "Jovana",
			LastName:  "Jovanovic",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 95,
				},
				Grade{
					Title: "Homework 1",
					Type:  GradeHomework,
					Score: 65,
				},
				Grade{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 85,
				},
			},
		},
	}
}
