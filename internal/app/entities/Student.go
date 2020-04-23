package entities

import "fmt"

// Student type
type Student struct {
	Id           int // PK
	FirstName    string
	LastName     string
	Age          int
	LanguageCode string
}

// Student builder
func NewStudent() Student {
	// new Student with default values ( 'zero values' )
	return Student{}
}

// Stringer interface implementation
func (this Student) String() string {
	return fmt.Sprintf(
		"[%d : %s, %s, %d, %s]",
		this.Id,
		this.FirstName,
		this.LastName,
		this.Age,
		this.LanguageCode)
}
