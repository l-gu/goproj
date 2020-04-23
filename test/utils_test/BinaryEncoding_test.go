package memdb_test

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"

	"github.com/l-gu/goproject/internal/app/entities"
)

func Encode(e interface{}) []byte {
	buf := bytes.Buffer{}
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(e)
	if err != nil {
		panic("Encode error")
	}
	return buf.Bytes() // []byte
}

func Decode(binary []byte, e interface{}) {
	buf := bytes.Buffer{}
	buf.Write(binary)
	decoder := gob.NewDecoder(&buf)
	err := decoder.Decode(e) // struct POINTER expected
	if err != nil {
		panic("Decode error")
	}
}
func decode2(binary []byte, e interface{}) {
	Decode(binary, e)
}
func decode1(binary []byte, e interface{}) {
	decode2(binary, e)
}

func newLanguageUntyped() interface{} {
	/*
		var e entities.Language
		e = entities.Language{}
		//return e // NOT OK (decode => pointer expected)
		return &e // OK
	*/
	return &entities.Language{} // OK
}

func TestEncodeDecode2(t *testing.T) {

	e1 := entities.Language{Code: "J", Name: "Java"}
	binary := Encode(e1)
	fmt.Printf("Type %T \n", e1)

	//e2 := entities.Language{} // OK
	// Decode(binary, &e2)

	e2 := entities.NewLanguage()
	//Decode(binary, &e2) // OK
	decode1(binary, &e2) // OK

	var ei interface{}
	ei = &e2
	decode1(binary, ei)

	eu := newLanguageUntyped()
	//Decode(binary, &eu) // ERROR
	Decode(binary, eu) // ERROR

	/*
		e2 := eu.(entities.Language)

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
	*/
}

func TestEncodeDecode1(t *testing.T) {

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
