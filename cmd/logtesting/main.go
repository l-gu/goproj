package main

import (
	"errors"
	"fmt"

	"github.com/l-gu/goproject/internal/app/logger"
)

func log(v ...interface{}) {
	logger.Log("main", v...)
}

func main() {
	fmt.Println("Start ...")

	fmt.Print("Test Print", "aa", 123, true, "bb")
	fmt.Println("Test Println", "aa", 123, true, "bb")

	logger.Log("foo", "aa", "bb")
	logger.Log("foo")
	logger.Log("foo", "aa", 12, true, 45.68, "bb")
	logger.Log0("foo", "aa", 12, true, 45.68, "bb")
	logger.Log("foo", "ERROR", errors.New("My error"))

	log("foo", "aa", 12, true, 45.68, "bb")
	log("single arg")
}

/*
func Log1(s string) {
	fmt.Println("[LOG] " + s)
}
func Log0(args ...interface{}) {
	fmt.Print("[LOG]")
	for _, x := range args {
		fmt.Print(" ", x)
	}
	fmt.Println()
}
func Log(name string, args ...interface{}) {
	//fmt.Printf("[LOG %s]", name)
	fmt.Print("[LOG ", name, "]")
	for _, x := range args {
		fmt.Print(" ", x)
	}
	fmt.Println()
}
*/
