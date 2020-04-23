package mvc

import (
	"fmt"

	"github.com/l-gu/goproject/internal/app/entities"
)

//-----------------------------------------------------------------------------
// LIST DATA
//-----------------------------------------------------------------------------
type LanguageListData struct {
	MVCListData // this type extends MVCListData
	// List to be used in page
	List []entities.Language
}

func NewLanguageListData(list []entities.Language) LanguageListData {
	// New form data (void)
	var listData LanguageListData
	listData.init("Language", "language")
	listData.List = list
	return listData
}

//-----------------------------------------------------------------------------
// FORM DATA
//-----------------------------------------------------------------------------
type LanguageFormData struct {
	MVCFormData // this type extends MVCFormData
	// Entity fields
	Code string // void by default
	Name string // void by default
}

func NewLanguageFormData(entity *entities.Language, validator *LanguageValidator, mode int) LanguageFormData {
	log("NewLanguageFormData()")
	// New form data (void)
	var formData LanguageFormData
	formData.init("Language", "language", mode, validator)

	if entity != nil {
		// Entity fields used in the form
		formData.Code = entity.Code
		formData.Name = entity.Name
	} else {
		// set default values
		// formData.Code = ""
	}

	return formData
}

// Stringer interface implementation
func (this *LanguageFormData) String() string {
	return fmt.Sprintf(
		"[%s : %s]",
		this.Code,
		this.Name)
}
