package main

import (
	"fmt"

	"github.com/l-gu/goproject/internal/app/entities"
	"github.com/l-gu/goproject/internal/app/persistence"
	"github.com/l-gu/goproject/internal/app/persistence/impl/boltdb"
)

func main() {
	fmt.Println("Test Bolt DB type ")

	//db := boltdb.NewBoltDB("test.db")
	db := boltdb.GetBoltDB()
	defer db.Close()

	fmt.Println("Bolt db path = " + db.Path())

	initLanguages()
	initStudents()
}

func initLanguages() {
	//dao := daoproviders.GetLanguageDAO()
	dao := persistence.GetSpecificLanguageDAO(persistence.BOLT_DAO)

	dao.Save(&entities.Language{Code: "J", Name: "Java"})
	dao.Save(&entities.Language{Code: "G", Name: "Golang"})
	dao.Save(&entities.Language{Code: "P", Name: "PHP"})

	fmt.Println("Find('J')")
	e := dao.Find("J")
	if e != nil {
		fmt.Println("Found")
	} else {
		fmt.Println("Not found")
	}
	dao.FindAll()

	r := dao.Exists("J")
	if r {
		fmt.Println("Exists")
	} else {
		fmt.Println("Does not exist")
	}

	dao.Delete("J")

}
func initStudents() {
	//	dao := daoproviders.GetStudentDAO()
	dao := persistence.GetSpecificStudentDAO(persistence.BOLT_DAO)

	dao.Save(&entities.Student{Id: 1, FirstName: "Bart", LastName: "Simpson"})
}
