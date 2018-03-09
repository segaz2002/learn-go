package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10.000

	/*Pointers*/
	ptr := &a

	fmt.Println("\nA is type", reflect.TypeOf(a), "and A is type", reflect.TypeOf(a))

	fmt.Println("\nPointer value (memory address) of A", ptr, "And the value stored in the address is", *ptr)
}
