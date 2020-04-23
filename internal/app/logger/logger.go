package logger

import (
	"fmt"
)

// Log with "zero name"
func Log0(args ...interface{}) {
	fmt.Print("[LOG]")
	for _, x := range args {
		fmt.Print(" ", x)
	}
	fmt.Println()
}

// Log with a name to be added in the prefix
func Log(name string, args ...interface{}) {
	//fmt.Printf("[LOG %s]", name)
	fmt.Print("[LOG ", name, "]")
	for _, x := range args {
		fmt.Print(" ", x)
	}
	fmt.Println()
}
