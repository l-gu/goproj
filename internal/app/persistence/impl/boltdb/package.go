package boltdb

import "github.com/l-gu/goproject/internal/app/config"

// Package implicit initialization - STEP 1 : constants
/*
const (
	Xxxx = xxx
)*/

// Package implicit initialization - STEP 2 : variables
var LogFlag bool = false
var boltDB *BoltDB = nil // Unique BOLT-DB instance

// Package implicit initialization - STEP 3 : init() functions
/*
func init() {
	boltDbFileName = config.GetBoltFileName()
}
*/

// Package explicit initialization (this function must be called explicitly)
func InitBoltDB() {
	// just used to trigger the package initialization
	// -> call all "init()" functions defined in each DAO
}

func GetBoltDB() *BoltDB {
	if boltDB == nil {
		boltDB = NewBoltDB(config.GetBoltFileName())
	}
	return boltDB
}

func CloseBoltDB() {
	if boltDB != nil {
		boltDB.Close()
	}
}
