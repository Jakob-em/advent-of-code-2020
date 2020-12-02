package utils

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReadLinesFromFile(filename string, separator string) []string {
	content, err := ioutil.ReadFile(filename)

	content = removeNewlineFromEnd(content)

	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(content), separator)
}

func removeNewlineFromEnd(content []byte) []byte {
	if string(content[len(content)-1]) == "\n" {
		content = content[:len(content)-1]
	}
	return content
}
