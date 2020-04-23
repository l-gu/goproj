package mvc

import (
	"net/http"
)

type StudentValidator struct {
	MVCValidator // this type extends MVCValidator
}

func NewStudentValidator() *StudentValidator {
	//return new(StudentValidator)
	v := new(StudentValidator)
	// abstract controller fields
	v.name = "StudentValidator"
	return v
}

func (this *StudentValidator) getId(r *http.Request) int {
	value, _ := this.getIntParam(r, "id", true)
	// insert other checking here
	return value
}
func (this *StudentValidator) getFirstName(r *http.Request) string {
	value, _ := this.getStringParam(r, "firstName", true)
	return value
}
func (this *StudentValidator) getLastName(r *http.Request) string {
	value, _ := this.getStringParam(r, "lastName", true)
	if len(value) == 0 {
		this.addError("invalid lastName : cannot be void")
	}
	return value
}
func (this *StudentValidator) getAge(r *http.Request) int {
	value, ok := this.getIntParam(r, "age", true)
	if !ok {
		// specific message
	} else {
		if value <= 0 {
			this.addError("invalid age : must be > 0 ")
		}
	}
	return value
}
func (this *StudentValidator) getLanguageCode(r *http.Request) string {
	value, _ := this.getStringParam(r, "languageCode", false)
	return value
}
