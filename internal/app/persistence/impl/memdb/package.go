package memdb

// Package implicit initialization - STEP 1 : constants
/*
const (
	XXX  = XX
)
*/

// Package implicit initialization - STEP 2 : variables
var LogFlag bool = false

// Package implicit initialization - STEP 3 : init() functions
/*
func init() {
}
*/

// Package explicit initialization (this function must be called explicitly)
func InitMemDB() {
	// just used to trigger the package initialization
	// -> call all "init()" functions defined in each DAO
}
