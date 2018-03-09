package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "hello world"
	fmt.Println("\nThis is the functions topic")
	fmt.Println("\nResult from function", changeCase(text))

	bestFinish := bestGameFinishes(12, 34, 50, 20, 15, 16) /*Slice*/

	fmt.Println("\nBest finish is", bestFinish)
}

func changeCase(text string) string {
	fmt.Println("\nChange case function is called")
	return strings.Title(text)
}

/*Variadic Function*/
func bestGameFinishes(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes { /*Looping*/
		if i < best {
			best = i
		}
	}
	return best
}
