package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

var dst string

func init() {
	flag.StringVar(&dst, "dst", "", "dst file: -dst filepath")
	flag.Parse()
}
func main() {

	var funcPattern string = `^func\s(\w+)\(.*?\).*?`
	var methodPattern string = `^func\s\(\w+\s(\*?.*?)\)\s(\w+)\(.*?\).*?`
	var typePattern string = `^type\s(.*?)\s(\w+)\s{`

	fmt.Println(dst)
	funcCompile, _ := regexp.Compile(funcPattern)
	methodCompile, _ := regexp.Compile(methodPattern)
	typeCompile, _ := regexp.Compile(typePattern)

	dstfile, err := os.Create(dst)
	check(err)
	defer dstfile.Close()

	scanner := bufio.NewScanner(os.Stdin)
	defer os.Stdin.Close()
	mflag := false

	for scanner.Scan() {

		line := scanner.Text()
		//struct和interface内容单独判断
		if mflag {
			fmt.Fprintln(dstfile, line)
			if line == "}" {
				fmt.Fprint(dstfile, "```\n\n")
				mflag = false
				continue
			}
		}
		//匹配普通函数func
		if ok := funcCompile.MatchString(line); ok {
			funcName := funcCompile.FindStringSubmatch(line)
			fmt.Fprintf(dstfile, "###func %s\n", funcName[1])
			withgo(dstfile, line)
			fmt.Fprint(dstfile, "```\n\n")
		}
		//匹配method
		if ok := methodCompile.MatchString(line); ok {
			methodName := methodCompile.FindStringSubmatch(line)
			fmt.Fprintf(dstfile, "###func (%s) %s\n", methodName[1], methodName[2])
			withgo(dstfile, line)
			fmt.Fprint(dstfile, "```\n\n")
		}
		//匹配struct和interface
		if ok := typeCompile.MatchString(line); ok {
			typeName := typeCompile.FindStringSubmatch(line)
			fmt.Fprintf(dstfile, "###type %s %s\n", typeName[1], typeName[2])
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
