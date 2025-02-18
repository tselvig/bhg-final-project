//package main
package getFiles

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

//usage: go run getFiles.go /

var regexes = []*regexp.Regexp{
	regexp.MustCompile(`\.go`), //how it decides to pick files should be improved, which regexp option to use?
}

func walkFn(path string, f os.FileInfo, err error) error {
	for _, r := range regexes {
		if r.MatchString(path) {
			fmt.Printf("[+] HIT: %s\n", path)
		}
	}
	return nil
}

func Jazz() {
	root := os.Args[1]
	if err := filepath.Walk(root, walkFn); err != nil {
		log.Panicln(err)
	}
}
