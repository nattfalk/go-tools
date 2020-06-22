package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"log"
	"strings"
	"unicode"
	"time"
)

func printOnly(r rune) rune {
	if unicode.IsPrint(r) {
		return r
	}
	return -1
}

func main() {
	startTime := time.Now()

	fmt.Println("Replacing bad XML chars in files")

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(wd)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.Mode().IsRegular() {
			// Skip folders
			continue
		}

		if !strings.HasSuffix(strings.ToLower(f.Name()), ".xml") {
			// Skip non XML files
			continue
		}
		fmt.Printf("    %s\n", f.Name())

		content, err := ioutil.ReadFile(f.Name())
		if err != nil {
			log.Fatal(err)
		}

		// Replace bad XML chars
		fixedContent := strings.ReplaceAll(string(content), "&#x2;", "(")
		fixedContent = strings.ReplaceAll(fixedContent, "&#x3;", ")")
		fixedContent = strings.Map(printOnly, fixedContent)

		err = ioutil.WriteFile(f.Name(), []byte(fixedContent), 0)
		if err != nil {
			log.Fatal(err)
		}	
	}

	fmt.Printf("Finished after %.2fs!\n", time.Since(startTime).Seconds())
}
