package entities

import "fmt"

// Language type
type Language struct {
	Code string // PK
	Name string
}

// Language builder
func NewLanguage() Language {
	// new structure with default values ( 'zero values' )
	return Language{}
}

// Stringer interface implementation
func (this Language) String() string {
	return fmt.Sprintf(
		"[%s : %s]",
		this.Code,
		this.Name)
}
