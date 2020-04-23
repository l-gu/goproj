package persistence

// Package implicit initialization - STEP 1 : constants
const (
	MEM_DAO  = 1 // defaut type ( see config package )
	BOLT_DAO = 2
)

// Package implicit initialization - STEP 2 : variables
var defaultDAOType int = MEM_DAO

// Package implicit initialization - STEP 3 : init() functions
/*
func init() {
}
*/

// Package explicit initialization (this function must be called explicitly)
func InitPersistence(daoType int) {
	defaultDAOType = daoType
}
