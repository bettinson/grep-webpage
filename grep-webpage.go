package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {
	args := os.Args[1:]
	var pattern, url string
	if len(args) >= 2 {
		pattern = args[0]
		url = args[1]
	} else {
		log.Fatal("Too few args. Fist arg: regex pattern, second url")
	}

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
		lines := searchReader(resp.Body, regexPredicate(pattern))
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}

func searchReader(contents io.Reader, predicate func(string) bool) []string {
	buff := bufio.NewReader(contents)
	lines := make([]string, 0)
	for {
		line, err := buff.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = line[:len(line)-1]
		if err != nil {
			log.Fatalf("Error when reading buffer: %s", err.Error())
		}
		if predicate(line) {
			lines = append(lines, line)
		}
	}
	return lines
}

func regexPredicate(pattern string) func(string) bool {
	reg := regexp.MustCompile(pattern)
	return func(s string) bool {
		return reg.MatchString(s)
	}
}
