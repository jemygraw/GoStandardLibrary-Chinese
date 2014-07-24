package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	zipFileName := "test.zip"

	rc, err := zip.OpenReader(zipFileName)
	if err != nil {
		log.Fatal(err)
	}

	zipFiles := rc.File
	for _, zipFile := range zipFiles {
		offset, _ := zipFile.DataOffset()
		log.Println("Compresszed:", zipFile.FileHeader.CompressedSize64, "Uncompressed:", zipFile.FileHeader.UncompressedSize64)
		log.Println(zipFile.FileHeader.Name, "\tDataOffset", offset)

		irc, err := zipFile.Open()
		if err != nil {
			log.Fatal(err)
		}

		if fileContent, err := ioutil.ReadAll(irc); err == nil {
			data := string(fileContent)
			fmt.Print(len(data), "\t", data)
		} else {
			log.Fatal(err)
		}
		irc.Close()
		fmt.Println("----------------------")
	}
}
