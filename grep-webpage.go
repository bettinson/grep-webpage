package main

import (
	"net/http"
	"os"
	"os/exec"
	"io/ioutil"
	"fmt"
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
		//fmt.Printf("%s\n", string(contents))
		grepCmd := exec.Command("grep title").Output()
		fmt.Println(grepCmd)
	}
	//fmt.Println(resp, err)
	//fmt.Println(args)
}

