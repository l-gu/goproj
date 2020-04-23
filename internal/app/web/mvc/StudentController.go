package mvc

import (
	"net/http"

	"github.com/l-gu/goproject/internal/app/entities"
	"github.com/l-gu/goproject/internal/app/persistence"
)

//------------------------------------------------------------------------
// M V C  CONTROLLER
//------------------------------------------------------------------------
// MVC Controller type
type StudentController struct {
	MVCController // this controller extends MVCController
	dao           persistence.StudentDAO
}

// MVC Controller constructor
func NewStudentController() StudentController {
	c := new(StudentController)
	// abstract controller fields
	c.name = "StudentController"
	c.listTemplateName = "studentList" // template file name for list
	c.formTemplateName = "studentForm" // template file name for form
	c.adapter = c                      // set adapter with interface checking
	// concrete controller fields
	c.dao = persistence.GetStudentDAO()
	return *c
}

//------------------------------------------------------------------------
// BUILDERS
//------------------------------------------------------------------------
func (this *StudentController) buildEntity(r *http.Request) (*entities.Student, *StudentValidator) {
	this.log("buildEntity...")
	entity := entities.NewStudent()
	validator := NewStudentValidator()
	entity.Id = validator.getId(r)
	entity.FirstName = validator.getFirstName(r)
	entity.LastName = validator.getLastName(r)
	entity.Age = validator.getAge(r)
	entity.LanguageCode = validator.getLanguageCode(r)

	this.log("Entity built : ", entity.String())
	return &entity, validator
}

//------------------------------------------------------------------------
// ADPATER INTERFACE IMPLEMENTATION
//------------------------------------------------------------------------
func (this *StudentController) getListData(r *http.Request) interface{} {
	list := this.dao.FindAll()
	return NewStudentListData(list)
}
func (this *StudentController) getVoidFormData(r *http.Request) interface{} {
	this.log("StudentController - getVoidFormData :", r.URL.Path)
	return NewStudentFormData(nil, nil, CREATION_MODE)
}

func (this *StudentController) getEntityFormData(r *http.Request) interface{} {
	this.log("StudentController - getEntityFormData :", r.URL.Path)
	validator := NewStudentValidator()
	// get PK fields
	id := validator.getId(r)
	this.log("StudentController : validator.getId(r) : ", id)
	if validator.hasError() {
		// invalid PK fields
		this.log("StudentController - getEntityFormData : validator has error")
		return NewStudentFormData(nil, validator, CREATION_MODE)
	} else {
		this.log("StudentController : validator : no error")
		entity := this.dao.Find(id)
		if entity != nil {
			// found
			this.log("StudentController - getEntityFormData : found")
			formData := NewStudentFormData(entity, validator, UPDATE_MODE)
			this.log("FormData : ", formData.String())
			return formData
		} else {
			// not found
			this.log("StudentController - getEntityFormData : not found")
			formData := NewStudentFormData(nil, validator, CREATION_MODE)
			formData.addMessage("not found")
			return formData
		}
	}
}

func (this *StudentController) create(r *http.Request) interface{} {
	entity, validator := this.buildEntity(r)
	if validator.hasError() {
		// invalid fields => no creation => stay in CREATION mode
		return NewStudentFormData(entity, validator, CREATION_MODE)
	} else {
		this.dao.Create(entity)
		// entity created => switch to UPDATE mode
		return NewStudentFormData(entity, validator, UPDATE_MODE)
	}
}

func (this *StudentController) delete(r *http.Request) bool {
	validator := NewStudentValidator()
	// get PK fields
	id := validator.getId(r)
	if validator.hasError() {
		return false
	} else {
		return this.dao.Delete(id)
	}
}

func (this *StudentController) update(r *http.Request) interface{} {
	entity, validator := this.buildEntity(r)
	if !validator.hasError() {
		this.dao.Update(entity)
	}
	// always stay in UPDATE mode
	return NewStudentFormData(entity, validator, UPDATE_MODE)
}
