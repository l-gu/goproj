package mvc

import (
	"net/http"
)

type LanguageValidator struct {
	MVCValidator // this type extends MVCValidator
}

func NewLanguageValidator() *LanguageValidator {
	//return new(LanguageValidator)
	v := new(LanguageValidator)
	// abstract controller fields
	v.name = "LanguageValidator"
	return v
}

func (this *LanguageValidator) getCode(r *http.Request) string {
	const name = "code"
	value, _ := this.getStringParam(r, name, true)
	//this.checkNotEmpty(value)
	this.checkNotBlank(name, value)
	this.checkMinLength(name, value, 1)
	this.checkMaxLength(name, value, 2)
	return value
}

func (this *LanguageValidator) getName(r *http.Request) string {
	const name = "name"
	value, _ := this.getStringParam(r, name, true)
	return value
}
