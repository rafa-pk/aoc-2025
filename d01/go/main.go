package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func	main() {
	
	if len(os.Args) != 2 {
		log.Fatalln("Program requires one argument")
	}
	var pw int
	dial := 50
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln("Error reading from file", err)
	}
	moves := strings.Split(string(input), "\n")

	for idx, line := range moves {
		
	}
}

// take av1
// read file to buff
// split file
// for each line do operation
// if below 0, goes to 99
// if over 99, goes to 0
//+1 counter every time 0 is reached
//return count

