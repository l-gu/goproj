package entities

import "fmt"

// Student type
type FooBar struct {
	Pk1    int    // PK
	Pk2    string // PK
	Name   string
	Age    int
	Wage   float32
	Weight float64
	Flag   bool
	Count  int64
}

// FooBar builder
func NewFooBar() FooBar {
	// new Student with default values ( 'zero values' )
	return FooBar{}
}

// Stringer interface implementation
func (this FooBar) String() string {
	return fmt.Sprintf(
		"[%v, %v : %v, %v, %v, %v, %v, %v]",
		this.Pk1,
		this.Pk2,
		this.Name,
		this.Age,
		this.Wage,
		this.Weight,
		this.Flag,
		this.Count)
}
