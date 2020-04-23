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
type FooBarController struct {
	MVCController // this controller extends MVCController
	dao           persistence.FooBarDAO
}

// MVC Controller constructor
func NewFooBarController() FooBarController {
	c := new(FooBarController)
	// abstract controller fields
	c.name = "FooBarController"
	c.listTemplateName = "foobarList" // template file name for list
	c.formTemplateName = "foobarForm" // template file name for form
	c.adapter = c                     // set adapter with interface checking
	// concrete controller fields
	c.dao = persistence.GetFooBarDAO()
	return *c
}

//------------------------------------------------------------------------
// BUILDERS
//------------------------------------------------------------------------
func (this *FooBarController) buildEntity(r *http.Request) (*entities.FooBar, *FooBarValidator) {
	this.log("buildEntity...")
	entity := entities.NewFooBar()
	validator := NewFooBarValidator()
	// populate entity from request data
	entity.Pk1 = validator.getPk1(r)
	entity.Pk2 = validator.getPk2(r)
	entity.Name = validator.getName(r)
	entity.Age = validator.getAge(r)
	entity.Wage = validator.getWage(r)
	entity.Weight = validator.getWeight(r)
	entity.Flag = validator.getFlag(r)
	entity.Count = validator.getCount(r)
	this.log("Entity built : ", entity.String())
	return &entity, validator
}

//------------------------------------------------------------------------
// ADPATER INTERFACE IMPLEMENTATION
//------------------------------------------------------------------------
func (this *FooBarController) getListData(r *http.Request) interface{} {
	list := this.dao.FindAll()
	return NewFooBarListData(list)
}
func (this *FooBarController) getVoidFormData(r *http.Request) interface{} {
	this.log("FooBarController - getVoidFormData :", r.URL.Path)
	return NewFooBarFormData(nil, nil, CREATION_MODE)
}

func (this *FooBarController) getEntityFormData(r *http.Request) interface{} {
	this.log("FooBarController - getEntityFormData :", r.URL.Path)
	validator := NewFooBarValidator()
	// get PK fields
	pk1 := validator.getPk1(r)
	pk2 := validator.getPk2(r)
	this.log("FooBarController : PK :", pk1, pk2)
	if validator.hasError() {
		// invalid PK parameters
		this.log("FooBarController - getEntityFormData : invalid key param (validator error)")
		return NewFooBarFormData(nil, validator, CREATION_MODE)
	} else {
		this.log("FooBarController : validator : no error")
		entity := this.dao.Find(pk1, pk2)
		if entity != nil {
			// found
			this.log("FooBarController - getEntityFormData : found")
			formData := NewFooBarFormData(entity, validator, UPDATE_MODE)
			this.log("FormData : ", formData.String())
			return formData
		} else {
			// not found
			this.log("FooBarController - getEntityFormData : not found")
			formData := NewFooBarFormData(nil, validator, CREATION_MODE)
			formData.addMessage("not found")
			return formData
		}
	}
}

func (this *FooBarController) create(r *http.Request) interface{} {
	entity, validator := this.buildEntity(r)
	if validator.hasError() {
		// invalid fields => no creation => stay in CREATION mode
		return NewFooBarFormData(entity, validator, CREATION_MODE)
	} else {
		this.dao.Create(entity)
		// entity created => switch to UPDATE mode
		return NewFooBarFormData(entity, validator, UPDATE_MODE)
	}
}

func (this *FooBarController) delete(r *http.Request) bool {
	validator := NewFooBarValidator()
	// get PK fields
	pk1 := validator.getPk1(r)
	pk2 := validator.getPk2(r)
	if validator.hasError() {
		return false
	} else {
		return this.dao.Delete(pk1, pk2)
	}
}

func (this *FooBarController) update(r *http.Request) interface{} {
	entity, validator := this.buildEntity(r)
	if !validator.hasError() {
		this.dao.Update(entity)
	}
	// always stay in UPDATE mode
	return NewFooBarFormData(entity, validator, UPDATE_MODE)
}
