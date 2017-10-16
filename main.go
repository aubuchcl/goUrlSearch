package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/aubuchcl/httpParser/strip"
)


type character struct {
	char  string
	count int
}

func main() {

	fmt.Println("hello world")

	client := &http.Client{
	//CheckRedirect: http.Redirect(w ResponseWriter, r *Request, url string, code int),
	}
	bs := make([]byte, 99999)

	resp, err := client.Get("http://golang.org")
	//resp, err := client.Get("http://www.lipsum.com")

	resp.Body.Read(bs)
	fmt.Println(err)

	stringSlice := stripResponse(string(bs))
	//charOnlySlice := charSlice(stringSlice)
	fmt.Println(stringSlice)
	//fmt.Println(charOnlySlice)
}

func stripResponse(responseString string) []string {
	var formattedSlice []string

	//replace this with a regex if you have time.
	innerHTML := strip.StripTags(responseString)
	innerHTML = strings.Replace(innerHTML, "\n", "", -1)
	innerHTML = strings.Replace(innerHTML, ",", "", -1)
	innerHTML = strings.Replace(innerHTML, "\t", "", -1)
	innerHTML = strings.Replace(innerHTML, ".", "", -1)
	innerHTML = strings.Replace(innerHTML, ",", "", -1)
	//innerHTML = strings.Replace(innerHTML, " ", "", -1)

	htmlSlice := strings.Split(innerHTML, " ")
	for _, piece := range htmlSlice {
		if piece != "" || false {
			formattedSlice = append(formattedSlice, piece)
		}
	}
	//whitespace was not being extracted in first loop this is
	//a workaround that gets the text from doc without junk
	noWhitespaceString := strings.Join(formattedSlice, ",")

	noWhitespaceString = strings.Replace(noWhitespaceString, " ", "", -1)
	noWhitespaceString = strings.Replace(noWhitespaceString, ",", "", -1)
	htmlSlice = strings.Split(noWhitespaceString, "")

	fmt.Println(htmlSlice)

	return formattedSlice
}

//function that takes a slice and makes a slice of chars out of it.
func charSlice(slc []string) []character {

	var newCharSlice []character
	for _, s := range slc {
		if s != "" {
			newCharSlice = append(newCharSlice, character{s, 0})
		}
	}

	return newCharSlice
}

//looking to compare newCharSlice[index].char
func charSort(slc []character) []character {
	return slc
}
