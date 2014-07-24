package main

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// 新建一个buffer
	buf := new(bytes.Buffer)
	// 创建一个zip的writer
	zipWriter := zip.NewWriter(buf)

	// 这里可以使用两种方式提供文件路径
	//var files = []struct {
	//	Name, Path string
	//}{
	//	{"a.txt", "/Users/jemy/Temp/go_standard_library/zip_demo/a.txt"},
	//	{"b.txt", "/Users/jemy/Temp/go_standard_library/zip_demo/b.txt"},
	//	{"c.txt", "/Users/jemy/Temp/go_standard_library/zip_demo/c.txt"},
	//}
	var files = []struct {
		Name, Path string
	}{
		{"a.txt", "a.txt"},
		{"b.txt", "b.txt"},
		{"c.txt", "c.txt"},
	}

	for _, file := range files {
		// 将每个文件添加到zip中
		fw, err := zipWriter.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		// 读取文件的内容，然后再写入zip文件
		if fileContent, err := ioutil.ReadFile(file.Path); err != nil {
			log.Fatal(err)
		} else {
			_, err = fw.Write(fileContent)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// 关闭zip的writer
	err := zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 将buffer数据写入文件
	err = ioutil.WriteFile("test.zip", buf.Bytes(), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
