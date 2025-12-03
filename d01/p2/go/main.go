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
		clicks, err := strconv.Atoi(instruction[1:])
		if err != nil {
			log.Fatalln("Error in Atoi")
		}
		if instruction[0] == 'L' {
			
			for idx := 0; idx < clicks; idx++ {
			
				dial--
				if dial < 0 {
					dial = 99
				}
				if dial == 0 {
					pw++
				}
			}
		
		} else if instruction[0] == 'R' {
			
			for idx := 0; idx < clicks; idx++ {
				
				dial++
				if dial > 99 {
					dial = 0
				}
				if dial == 0 {
					pw++
				}
			}
		}
		fmt.Println("Instruction[", i, "]: ", instruction, "Dial: ", dial, "  Total 0 points: ", pw)
	}
	fmt.Println("The password is ", pw)
}
