package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/**
A simple program to read input from stdin
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
