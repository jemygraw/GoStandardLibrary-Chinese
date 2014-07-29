package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var src string
var dst string

func init() {
	flag.StringVar(&dst, "dst", "", "dst file: -dst filepath")
	flag.Parse()
}
func main() {

	var funcPattern string = `^func\s(\w+)\(.*?\)`
	var methodPattern string = `^func\s\(\w.*?\s(.*?)\)\s(\w+)`
	var typePattern string = `^type\s(.*?)\s\w+\s{`

	fmt.Println(src, dst)
	funcCompile, _ := regexp.Compile(funcPattern)
	methodCompile, _ := regexp.Compile(methodPattern)
	typeCompile, _ := regexp.Compile(typePattern)

	srcfile := os.Stdin

	dstfile, err := os.Create(dst)
	check(err)
	defer dstfile.Close()

	scanner := bufio.NewScanner(srcfile)
	defer srcfile.Close()
	mflag := false

	for scanner.Scan() {

		line := scanner.Text()
		if mflag {
			fmt.Fprintln(dstfile, line)
			if line == "}" {
				fmt.Fprint(dstfile, "```\n\n")
				mflag = false
				continue
			}
		}

		if ok := funcCompile.MatchString(line); ok {
			funcName := funcCompile.FindStringSubmatch(line)
			fmt.Fprintf(dstfile, "###%s\n", funcName[0])
			withgo(dstfile, line)
			fmt.Fprint(dstfile, "```\n\n")
		}

		if ok := methodCompile.MatchString(line); ok {
			methodName := methodCompile.FindStringSubmatch(line)
			fmt.Fprintf(dstfile, "###%s\n", methodName[0])
			withgo(dstfile, line)
			fmt.Fprint(dstfile, "```\n\n")
		}

		if ok := typeCompile.MatchString(line); ok {
			typeName := typeCompile.FindStringSubmatch(line)
			fmt.Fprintf(dstfile, "###%s\n", strings.Trim(typeName[0], "{"))
			withgo(dstfile, line)
			mflag = true
		}

	}
	fmt.Println("completed!")

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func withgo(f *os.File, line string) {
	fmt.Fprint(f, "```go\n")
	fmt.Fprintln(f, line)
}
