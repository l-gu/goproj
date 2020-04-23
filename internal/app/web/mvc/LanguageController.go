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
type LanguageController struct {
	MVCController // this type extends MVCController
	dao           persistence.LanguageDAO
}

// MVC Controller constructor
func NewLanguageController() LanguageController {
	c := new(LanguageController)
	// abstract controller fields
	c.name = "LanguageController"
	c.listTemplateName = "languageList" // template file name for list
	c.formTemplateName = "languageForm" // template file name for form
	c.adapter = c                       // set adapter with interface checking
	// concrete controller fields
	c.dao = persistence.GetLanguageDAO()
	return *c
}

//------------------------------------------------------------------------
// BUILDERS
//------------------------------------------------------------------------
// Entity builder
func (this *LanguageController) buildEntity(r *http.Request) (*entities.Language, *LanguageValidator) {
	this.log("buildEntity...")
	entity := entities.NewLanguage()
	validator := NewLanguageValidator()
	entity.Code = validator.getCode(r)
	entity.Name = validator.getName(r)

	this.log("Entity built : " + entity.String())
	return &entity, validator
}

//------------------------------------------------------------------------
// ADPATER INTERFACE IMPLEMENTATION
//------------------------------------------------------------------------
func (this *LanguageController) getListData(r *http.Request) interface{} {
	list := this.dao.FindAll()
	return NewLanguageListData(list)
}

func (this *LanguageController) getVoidFormData(r *http.Request) interface{} {
	this.log("getVoidFormData(r) :", r.URL.Path)
	return NewLanguageFormData(nil, nil, CREATION_MODE)
}
func (this *LanguageController) getEntityFormData(r *http.Request) interface{} {
	this.log("getEntityFormData() :", r.URL.Path)
	validator := NewLanguageValidator()
	// get PK fields
	code := validator.getCode(r)
	this.log("getEntityFormData() : code = ", code)
	if validator.hasError() {
		// invalid PK fields
		this.log("getEntityFormData() : validator has error")
		return NewLanguageFormData(nil, validator, CREATION_MODE)
	} else {
		this.log("getEntityFormData() : validator has no error -> find")
		entity := this.dao.Find(code)
		if entity != nil {
			// found
			this.log("getEntityFormData() : found")
			return NewLanguageFormData(entity, validator, UPDATE_MODE)
		} else {
			// not found
			this.log("getEntityFormData() : not found")
			formData := NewLanguageFormData(nil, validator, CREATION_MODE)
			formData.addMessage("not found")
			return formData
		}
	}
}

// create and stay on form
func (this *LanguageController) create(r *http.Request) interface{} {
	entity, validator := this.buildEntity(r)
	if validator.hasError() {
		// invalid fields => no creation => stay in CREATION mode
		return NewLanguageFormData(entity, validator, CREATION_MODE)
	} else {
		this.dao.Create(entity)
		// entity created => switch to UPDATE mode
		return NewLanguageFormData(entity, validator, UPDATE_MODE)
	}
}

// delete and go to list
func (this *LanguageController) delete(r *http.Request) bool {
	validator := NewLanguageValidator()
	// get PK fields
	code := validator.getCode(r)
	if validator.hasError() {
		return false
	} else {
		return this.dao.Delete(code)
	}
}

// update and stay on form
func (this *LanguageController) update(r *http.Request) interface{} {
	entity, validator := this.buildEntity(r)
	if !validator.hasError() {
		this.dao.Update(entity)
	}
	// always stay in UPDATE mode
	return NewLanguageFormData(entity, validator, UPDATE_MODE)
}
