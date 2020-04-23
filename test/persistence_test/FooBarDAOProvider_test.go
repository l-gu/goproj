package persistence_test

import (
	"fmt"
	"testing"

	"github.com/l-gu/goproject/internal/app/persistence"
	"github.com/l-gu/goproject/internal/app/persistence/impl/memdb"
)

func TestGetFooBarDAO(t *testing.T) {

	memdb.InitPackage()
	var dao persistence.FooBarDAO
	dao = persistence.GetFooBarDAO()

	if dao != nil {
		e := dao.Find(1, "A")
		if e != nil {
			fmt.Println("Found")
		} else {
			fmt.Println("Not found")
		}
		dao.Exists(1, "X")
	} else {
		t.Error("DAO NOT FOUND")
	}

}
func TestGetSpecificFooBarDAOForMemDB(t *testing.T) {
	var dao persistence.FooBarDAO
	dao = persistence.GetSpecificFooBarDAO(persistence.MEM_DAO)
	if dao != nil {
		dao.Exists(1, "X")
	} else {
		t.Error("DAO NOT FOUND")
	}
}

/*
func TestGetSpecificFooBarDAOForBoltDB(t *testing.T) {
	var dao persistence.FooBarDAO
	dao = persistence.GetSpecificFooBarDAO(persistence.BOLT_DAO)
	if dao != nil {
		dao.Exists(1, "X")
	} else {
		t.Error("DAO NOT FOUND")
	}
}
*/
