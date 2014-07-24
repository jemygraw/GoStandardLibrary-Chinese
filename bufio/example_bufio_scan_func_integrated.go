package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str = "hello world\ni am jemy\r"
	scanner := bufio.NewScanner(strings.NewReader(str))

	fmt.Println("----------ScanBytes----------")
	scanner.Split(bufio.ScanBytes)
	count := 0
	for scanner.Scan() {
		count++
		fmt.Print(scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Println("Byte Count:", count)

	fmt.Println("----------ScanLines----------")
	scanner = bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanLines)
	count = 0
	for scanner.Scan() {
		count++
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Println("Line Count:", count)

	fmt.Println("----------ScanRunes----------")
	scanner = bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanRunes)
	count = 0
	for scanner.Scan() {
		count++
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Println("Rune Count:", count)

	fmt.Println("----------ScanWords----------")
	scanner = bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)
	count = 0
	for scanner.Scan() {
		count++
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Println("Word Count:", count)
}
