package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
)

/*
var (
	name   string
	course string
	module float64
)
*/

/* Type inference*/
var (
	name, course, module = "Gabriel", "Go Language", 3.2
)

func main() {
	name := os.Getenv("USER")
	moduleNo := module * 0.1
	fmt.Println("Hello from", runtime.GOOS)
	fmt.Println("Name is set to", name)
	fmt.Println("Module is set to", module)
	fmt.Println("Module is set to", reflect.TypeOf(module))
	fmt.Println("Module Number is set to", reflect.TypeOf(moduleNo))
	fmt.Println("Module is set to", reflect.TypeOf(name))
}
