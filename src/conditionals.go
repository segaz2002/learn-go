package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("Conditionals module")

	/*Initialized and scoped variables to the condition block*/
	if firstRank, secondRank := 39, 614; firstRank < secondRank {
		fmt.Println("\n First rank is greater tham that of the second")
	} else if secondRank < firstRank {
		fmt.Println("\n Second rank is greater tham that of the first")
	} else {
		fmt.Println("\n Something funny is happening")
	}

	/*switch*/
	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("\n Temp num is an even number", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("\n An Odd number was received", tmpNum)
	}

	file, err := os.Open("/Users/gabrielo/go/learn-go/src/env.go")
	if err != nil {
		fmt.Println("\n An error was encvountered on opening file", err)
	} else {
		fmt.Println("\n Openend file pointer", file)
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
