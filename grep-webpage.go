package main

import (
	"net/http"
	"os"
//	"os/exec"
	"io/ioutil"
	"fmt"
	"regexp"
)

func main () {
	args := os.Args[1:]
	firstArg := ""
	if (args[0] != "") {
		firstArg = args[0]
	} else
	{
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

