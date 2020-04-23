package memdb_test

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"

	"github.com/l-gu/goproject/internal/app/entities"
)

func Test1(t *testing.T) {

	var err error
	e1 := entities.Language{Code: "J", Name: "Java"}

	buf := bytes.Buffer{}
	encoder := gob.NewEncoder(&buf)
	err = encoder.Encode(e1)
	if err != nil {
		t.Error("Encode EROR")
	}
	var binary []byte
	binary = buf.Bytes()

	buf2 := bytes.Buffer{}
	buf2.Write(binary)
	decoder := gob.NewDecoder(&buf2)
	e2 := entities.Language{}
	err = decoder.Decode(&e2)
	if err != nil {
		t.Error("Encode EROR")
	}

	if e1.Code == e2.Code {
		fmt.Println("Code OK")
	} else {
		t.Error("Code !=")
	}

	if e1.Name == e2.Name {
		fmt.Println("Name OK")
	} else {
		t.Error("Name !=")
	}
}
