package mvc

import (
	"net/http"
)

type FooBarValidator struct {
	MVCValidator // this type extends MVCValidator
}

func NewFooBarValidator() *FooBarValidator {
	//return new(FooBarValidator)
	v := new(FooBarValidator)
	// abstract controller fields
	v.name = "FooBarValidator"
	return v
}

func (this *FooBarValidator) getPk1(r *http.Request) int {
	value, _ := this.getIntParam(r, "pk1", true)
	// insert other checking here
	return value
}
func (this *FooBarValidator) getPk2(r *http.Request) string {
	value, _ := this.getStringParam(r, "pk2", true)
	return value
}
func (this *FooBarValidator) getName(r *http.Request) string {
	value, _ := this.getStringParam(r, "name", true)
	return value
}
func (this *FooBarValidator) getAge(r *http.Request) int {
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
func (this *FooBarValidator) getWage(r *http.Request) float32 {
	value, _ := this.getFloat32Param(r, "wage", false)
	return value
}
func (this *FooBarValidator) getWeight(r *http.Request) float64 {
	value, _ := this.getFloat64Param(r, "weight", false)
	return value
}
func (this *FooBarValidator) getFlag(r *http.Request) bool {
	value, _ := this.getBoolParam(r, "flag", false)
	return value
}
func (this *FooBarValidator) getCount(r *http.Request) int64 {
	value, _ := this.getInt64Param(r, "count", false)
	return value
}
