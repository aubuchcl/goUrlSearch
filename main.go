package main

import (
	"fmt"
	"net/http"
	"sort"
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
	bs := make([]byte, 32*500)

	resp, err := client.Get("http://golang.org")
	//resp, err := client.Get("http://www.lipsum.com")

	resp.Body.Read(bs)
	fmt.Println(err)

	stringSlice := stripResponse(string(bs))
	charOnlySlice := charSlice(stringSlice)
	charSliceSort := charSort(charOnlySlice)

	if 12 == 23 {
		fmt.Println(stringSlice)
		fmt.Println(charSliceSort)
	}
	//fmt.Println(charOnlySlice)

}

func stripResponse(responseString string) []string {

	//replace this with a regex if you have time.
	innerHTML := strip.StripTags(responseString)
	innerHTML = strings.Replace(innerHTML, "\n", "", -1)
	innerHTML = strings.Replace(innerHTML, ",", "", -1)
	innerHTML = strings.Replace(innerHTML, "\t", "", -1)
	innerHTML = strings.Replace(innerHTML, ".", "", -1)
	innerHTML = strings.Replace(innerHTML, ",", "", -1)
	//innerHTML = strings.Replace(innerHTML, " ", "", -1)

	htmlSlice := strings.Split(innerHTML, " ")
	fmt.Println(htmlSlice)

	return htmlSlice
}

//function that takes a slice and makes a slice of chars out of it.
func charSlice(slc []string) []character {

	var newCharSlice []character
	for _, s := range slc {
		for _, c := range string(s) {
			newCharSlice = append(newCharSlice, character{string(c), 0})
		}
	}

	return newCharSlice
}

//looking to compare newCharSlice[index].char
func charSort(slc []character) []character {
	sort.SliceStable(slc, func(i, j int) bool {
		return slc[i].char < slc[j].char
	})
	fmt.Println("By Char:", slc)
	return slc
}
