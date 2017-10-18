package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/aubuchcl/httpParser/webcrawler"
)

func main() {
	//cap and validate args from command line
	cliArgs := os.Args
	if len(cliArgs) != 2 {
		os.Exit(1)
	}
	webURL := cliArgs[1]
	_, err := url.ParseRequestURI(webURL)
	if err != nil {
		fmt.Println("please enter a Valid URL and try again")
		os.Exit(1)
	}

	//give url to webcrawler so it can find most freq chars
	webcrawler.FormatIO(webURL)
	//webcrawler watches cli for a command to close program
	webcrawler.ServeScan()

}
