package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/aubuchcl/httpParser/webcrawler"
)

func main() {
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

	mfChar, mfCharNum := webcrawler.FormatIO(webURL)
	fmt.Println(mfChar, "occurs", mfCharNum, "times")
	webcrawler.ServeScan()

}
