package teacherportal

import "html/template"

var rootTemplate *template.Template

// ImportTemplates loads Go templates
func ImportTemplates() error {
	var err error
	rootTemplate, err = template.ParseFiles(
		"teacherportal/students.gohtml",
		"teacherportal/student.gohtml",
	)
	if err != nil {
		return err
	}
	return nil
}
