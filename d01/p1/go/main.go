package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
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

	fmt.Println(" Dial: ", dial)
	for i, instruction := range moves {
		
		if instruction == "" {
			continue
		}
		if instruction[0] == 'L' {
			clicks, err := strconv.Atoi(instruction[1:])
			if err != nil {
				log.Fatalln("Error in Atoi", err)
			}
			dial = (dial - clicks) % 100
			for dial < 0 {
				dial += 100
			}
		} else if instruction[0] == 'R' {
			clicks, err:= strconv.Atoi(instruction[1:])
			if err != nil {
				log.Fatalln("Error in Atoi", err)
			}
			dial = (dial + clicks) % 100
		}
		fmt.Println("Instruction[", i, "]: ", instruction, "  Dial: ", dial)
		if dial == 0 {
			pw++
		}
	}
	fmt.Println("The password is", pw)
}

