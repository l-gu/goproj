package mvc

import (
	"fmt"

	"github.com/l-gu/goproject/internal/app/entities"
	"github.com/l-gu/goproject/internal/app/persistence"
)

//-----------------------------------------------------------------------------
// LIST DATA
//-----------------------------------------------------------------------------
type StudentListData struct {
	MVCListData // this type extends MVCListData
	// List to be used in page
	List []entities.Student
}

func NewStudentListData(list []entities.Student) StudentListData {
	// New form data (void)
	var listData StudentListData
	listData.init("Student", "student")
	listData.List = list
	return listData
}

//-----------------------------------------------------------------------------
// FORM DATA
//-----------------------------------------------------------------------------
// Data structure for form template
type StudentFormData struct {
	MVCFormData // this type extends MVCFormData
	// one string for each form field value (string to allow void value)
	// "zero value" is empty ("") for strings.
	Id           string // void by default
	FirstName    string // void by default
	LastName     string // void by default
	Age          string // void by default
	LanguageCode string // void by default

	// further data
	Languages []entities.Language
}

func NewStudentFormData(entity *entities.Student, validator *StudentValidator, mode int) StudentFormData {
	log("StudentFormData", "NewStudentFormData")
	// New form data (void)
	var formData StudentFormData
	formData.init("Student", "student", mode, validator)

	// Further data
	formData.Languages = persistence.GetLanguageDAO().FindAll() // use cache

	if entity != nil {
		// Entity fields used in form
		formData.Id = intToString(entity.Id)
		formData.FirstName = entity.FirstName
		formData.LastName = entity.LastName
		formData.Age = intToString(entity.Age)
		formData.LanguageCode = entity.LanguageCode
	} else {
		// set default values
		formData.Id = "0"
	}

	// Return form data
	return formData
}

// Stringer interface implementation
func (this *StudentFormData) String() string {
	return fmt.Sprintf(
		"[%s : %s, %s, %s, %s]",
		this.Id,
		this.FirstName,
		this.LastName,
		this.Age,
		this.LanguageCode)
}
