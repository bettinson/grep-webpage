package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	//	"os/exec"
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	args := os.Args[1:]
	firstArg := ""
	if args[0] != "" {
		firstArg = args[0]
	} else {
		os.Exit(1)

	}
	r, _ := regexp.Compile(`/<\s*\w.*?>/g`)

	resp, err := http.Get(firstArg)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Println(string(contents))
		fmt.Println(r.FindAllStringSubmatchIndex(string(contents), -1))
		//fmt.Println(grepCmd)
	}
	//fmt.Println(resp, err)
	//fmt.Println(args)
}

func searchReader(contents io.Reader, predicate func(string) bool) []string {
	buff := bufio.NewReader(contents)
	lines := make([]string, 0)
	for {
		line, err := buff.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error when reading buffer: %s", err.Error())
		}
		if predicate(line) {
			lines = append(lines, line)
		}
	}
	return lines
}
