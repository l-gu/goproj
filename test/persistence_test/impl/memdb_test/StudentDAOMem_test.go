package memdb_test

import (
	"fmt"
	"testing"

	"github.com/l-gu/goproject/internal/app/entities"
	"github.com/l-gu/goproject/internal/app/persistence/impl/memdb"
)

func TestStudentDaoMem(t *testing.T) {

	dao := memdb.NewStudentDAOMem()

	// create for the first time
	if dao.Create(&entities.Student{Id: 123, FirstName: "John", LastName: "Doe"}) {
		fmt.Println("OK, CREATED")
	} else {
		t.Error("NOT CREATED")
	}

	// try to create again
	if dao.Create(&entities.Student{Id: 123}) {
		t.Error("NOT SUPPOSED TO BE CREATED")
	} else {
		fmt.Println("OK, NOT CREATED")
	}

	// find
	student := dao.Find(123)
	if student != nil {
		fmt.Println("OK, FOUND : " + student.String())
	} else {
		t.Error("NOT FOUND")
	}

	// update existing entity
	r := dao.Update(&entities.Student{Id: 123, FirstName: "Bob", LastName: "Doe"})
	if r {
		fmt.Println("OK, UPDATED")
	} else {
		t.Error("NOT UPDATED")
	}

	// check existence
	if dao.Exists(123) {
		fmt.Println("OK, EXISTS")
	} else {
		t.Error("DOES NOT EXIST")
	}

	all := dao.FindAll()
	if len(all) != 1 {
		t.Error("UNEXPECTED LENGTH")
	}

	// delete existing entity
	if dao.Delete(123) {
		fmt.Println("OK, DELETED")
	} else {
		t.Error("NOT DELETED")
	}

	// delete nonexistent  entity
	if dao.Delete(123) {
		t.Error("NOT SUPPOSED TO BE DELETED")
	} else {
		fmt.Println("OK, NOT DELETED")
	}
}
