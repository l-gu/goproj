package mvc

import (
	"fmt"

	"github.com/l-gu/goproject/internal/app/entities"
)

//-----------------------------------------------------------------------------
// LIST DATA
//-----------------------------------------------------------------------------
type FooBarListData struct {
	MVCListData // this type extends MVCListData
	// List to be used in page
	List []entities.FooBar
}

func NewFooBarListData(list []entities.FooBar) FooBarListData {
	// New form data (void)
	var listData FooBarListData
	listData.init("FooBar", "foobar")
	listData.List = list
	return listData
}

//-----------------------------------------------------------------------------
// FORM DATA
//-----------------------------------------------------------------------------
// Data structure for form template
type FooBarFormData struct {
	MVCFormData // this type extends MVCFormData
	// one string for each form field value (string to allow void value)
	// "zero value" is empty ("") for strings.
	Pk1    string // void by default
	Pk2    string // void by default
	Name   string // void by default
	Age    string // void by default
	Wage   string // void by default
	Weight string // void by default
	Flag   string // void by default
	Count  string // void by default

}

func NewFooBarFormData(entity *entities.FooBar, validator *FooBarValidator, mode int) FooBarFormData {
	log("FooBarFormData", "NewFooBarFormData")
	// New form data (void)
	var formData FooBarFormData

	formData.init("FooBar", "foobar", mode, validator)

	if entity != nil {
		// Entity fields used in form
		formData.Pk1 = intToString(entity.Pk1)
		formData.Pk2 = entity.Pk2
		formData.Name = entity.Name
		formData.Age = intToString(entity.Age)
		formData.Wage = float32ToString(entity.Wage)
		formData.Weight = float64ToString(entity.Weight)
		formData.Flag = boolToString(entity.Flag)
		formData.Count = int64ToString(entity.Count)
	} else {
		// set default values
		formData.Pk1 = "0"
		formData.Pk2 = ""
	}
	// Return form data
	return formData
}

// Stringer interface implementation
func (this *FooBarFormData) String() string {
	return fmt.Sprintf(
		"[%s, %s : %s, %s, %s, %s, %s, %s]",
		this.Pk1,
		this.Pk2,
		this.Name,
		this.Age,
		this.Wage,
		this.Weight,
		this.Flag,
		this.Count)
}
