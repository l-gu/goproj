package memdb_test

import (
	"fmt"

	"github.com/l-gu/goproject/internal/app/persistence/impl/memdb"
)

func init() {
	fmt.Println("-------------------------")
	fmt.Println("memdb - init()")
	fmt.Println("-------------------------")

	memdb.InitMemDB()
}
