package persistence_test

import (
	"fmt"

	"github.com/l-gu/goproject/internal/app/persistence"
	"github.com/l-gu/goproject/internal/app/persistence/impl/boltdb"
	"github.com/l-gu/goproject/internal/app/persistence/impl/memdb"
)

func init() {
	fmt.Println("-------------------------")
	fmt.Println("persistence_test - init()")
	fmt.Println("-------------------------")
	persistence.InitPersistence(persistence.MEM_DAO)
	memdb.InitMemDB()
	boltdb.InitBoltDB()

}
