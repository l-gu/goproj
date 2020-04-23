package memdb_test

import (
	"fmt"
	"testing"

	"github.com/l-gu/goproject/internal/app/entities"
	"github.com/l-gu/goproject/internal/app/persistence/impl/memdb"
)

// Entity builder (used to deserialize)
func createLanguageEntityForDecoding() interface{} {
	return &entities.Language{}
}

func Test1(t *testing.T) {

	db := memdb.NewMemDB(createLanguageEntityForDecoding)

	db.Put("J", &entities.Language{Code: "J", Name: "Java"})

	e := db.Get("J")
	if e != nil {
		fmt.Println("OK, FOUND")
	} else {
		t.Error("NOT FOUND")
	}

	all := db.GetAll()
	if len(all) == 1 {
		fmt.Println("OK, LEN IS CORRECT")
	} else {
		t.Error("UNEXPECETD LEN")
	}

}
