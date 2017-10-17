package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/aubuchcl/httpParser/ioformat"
)

func main() {

	// new main function will be
	//grab url from os.args
	var useURL string
	//validate it
	for _, u := range os.Args {
		_, urlError := url.ParseRequestURI(u)
		if urlError == nil {
			useURL = u
		}
	}

	mfChar, mfCharNum := ioformat.FormatIO(useURL)

	fmt.Println(mfChar, "occurs", mfCharNum, "times")

	serveScan()

}
