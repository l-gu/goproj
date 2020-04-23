package main

import (
	"github.com/l-gu/goproject/internal/app/persistence"
	"github.com/l-gu/goproject/internal/app/persistence/impl/boltdb"
	"github.com/l-gu/goproject/internal/app/persistence/impl/memdb"
)

func initPersistence(daoType int) {

	log("Persistence initialization - Start")

	persistence.InitPersistence(daoType)

	// Init all DAO implementations (for DAO registration)
	memdb.InitMemDB()
	memdb.LogFlag = true

	boltdb.InitBoltDB()
	boltdb.LogFlag = true

	log("Persistence initialization - End")
}
