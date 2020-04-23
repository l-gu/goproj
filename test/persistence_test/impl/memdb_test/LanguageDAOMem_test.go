package memdb_test

import (
	"fmt"
	"testing"

	"github.com/l-gu/goproject/internal/app/entities"
	"github.com/l-gu/goproject/internal/app/persistence/impl/memdb"
)

func TestLanguageDaoMem(t *testing.T) {

	dao := memdb.NewLanguageDAOMem()

	// create for the first time
	if dao.Create(&entities.Language{Code: "J", Name: "Java"}) {
		fmt.Println("OK, CREATED")
	} else {
		t.Error("NOT CREATED")
	}

	// try to create again
	if dao.Create(&entities.Language{Code: "J", Name: "Java"}) {
		t.Error("NOT SUPPOSED TO BE CREATED")
	} else {
		fmt.Println("OK, NOT CREATED")
	}

	// find
	language := dao.Find("J")
	if language != nil {
		fmt.Println("OK, FOUND : " + language.String())
	} else {
		t.Error("NOT FOUND")
	}

	// update existing entity
	r := dao.Update(&entities.Language{Code: "J", Name: "Java"})
	if r {
		fmt.Println("OK, UPDATED")
	} else {
		t.Error("NOT UPDATED")
	}

	// check existence
	if dao.Exists("J") {
		fmt.Println("OK, EXISTS")
	} else {
		t.Error("DOES NOT EXIST")
	}

	all := dao.FindAll()
	if len(all) != 1 {
		t.Error("UNEXPECTED LENGTH")
	}

	// delete existing entity
	if dao.Delete("J") {
		fmt.Println("OK, DELETED")
	} else {
		t.Error("NOT DELETED")
	}

	// delete nonexistent  entity
	if dao.Delete("J") {
		t.Error("NOT SUPPOSED TO BE DELETED")
	} else {
		fmt.Println("OK, NOT DELETED")
	}

	// check existence
	if dao.Exists("J") {
		t.Error("NOT SUPPOSED TO EXIST")
	} else {
		fmt.Println("OK, NOT FOUND")
	}

	// find
	if dao.Find("J") != nil {
		t.Error("NOT SUPPOSED TO EXIST")
	} else {
		fmt.Println("OK, NOT FOUND")
	}
}
