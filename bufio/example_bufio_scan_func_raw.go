package main

import (
	"bufio"
	"fmt"
)

func main() {
	var str = "hello world\ni am jemy\r"

	var buf = []byte(str)
	fmt.Println("----------ScanBytes----------")
	for {
		advance, token, err := bufio.ScanBytes(buf, true)
		if advance == 0 {
			break
		}
		fmt.Println(advance, token, err)
		if advance <= len(buf) {
			buf = buf[advance:]
		}
	}

	fmt.Println("----------ScanLines----------")
	buf = []byte(str)
	for {
		advance, token, err := bufio.ScanLines(buf, true)
		if advance == 0 {
			break
		}
		fmt.Print(advance, string(token), err)
		fmt.Println()
		if advance <= len(buf) {
			buf = buf[advance:]
		}
	}

	fmt.Println("----------ScanRunes----------")
	buf = []byte(str)
	for {
		advance, token, err := bufio.ScanRunes(buf, true)
		if advance == 0 {
			break
		}
		fmt.Print(advance, string(token), len(token), err)
		fmt.Println()
		if advance <= len(buf) {
			buf = buf[advance:]
		}
	}

	fmt.Println("----------ScanWords----------")
	buf = []byte(str)
	for {
		advance, token, err := bufio.ScanWords(buf, true)
		if advance == 0 {
			break
		}
		fmt.Print(advance, string(token), len(token), err)
		fmt.Println()
		if advance <= len(buf) {
			buf = buf[advance:]
		}
	}
}
