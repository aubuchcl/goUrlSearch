package main

import (
	"fmt"
	"os"

	"github.com/aubuchcl/httpParser/webcrawler"
)

func main() {

	fmt.Println(os.Args[1:])

	// new main function will be
	// grab url from os.args
	// var useURL string

	// //validate it
	// for _, u := range os.Args {
	// 	_, urlError := url.ParseRequestURI(u)
	// 	if urlError == nil {
	// 		useURL = u
	// 	}
	// }

	// if err != nil {
	// 	os.Exit(1)
	// }
	mfChar, mfCharNum := webcrawler.FormatIO("http://www.google.com")

	fmt.Println(mfChar, "occurs", mfCharNum, "times")

	webcrawler.ServeScan()

}
