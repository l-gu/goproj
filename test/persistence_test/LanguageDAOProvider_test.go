package persistence_test

import (
	"testing"

	"github.com/l-gu/goproject/internal/app/persistence"
)

func TestGetLanguageDAO(t *testing.T) {

	var dao persistence.LanguageDAO = persistence.GetLanguageDAO()

	//dao = persistence.GetLanguageDAO()
	if dao != nil {
		dao.Find("J")
		dao.Exists("X")
	} else {
		t.Error("DAO NOT FOUND")
	}
}
func TestGetSpecificLanguageDAOForMemDB(t *testing.T) {

	var dao persistence.LanguageDAO
	dao = persistence.GetSpecificLanguageDAO(persistence.MEM_DAO)
	if dao != nil {
		dao.Find("J")
		dao.Exists("X")
	} else {
		t.Error("DAO NOT FOUND")
	}
}
func TestGetSpecificLanguageDAOForBoltDB(t *testing.T) {

	var dao persistence.LanguageDAO
	dao = persistence.GetSpecificLanguageDAO(persistence.BOLT_DAO)
	if dao != nil {
		dao.Find("J")
		dao.Exists("X")
	} else {
		t.Error("DAO NOT FOUND")
	}
}
