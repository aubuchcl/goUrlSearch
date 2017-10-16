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
	innerHTML := strip.StripTags(string(bs))
	//replace this with a regex if you have time.
	innerHTML = strings.Replace(innerHTML, "\n", "", -1)
	innerHTML = strings.Replace(innerHTML, ",", "", -1)
	innerHTML = strings.Replace(innerHTML, "\t", "", -1)
	innerHTML = strings.Replace(innerHTML, ".", "", -1)
	innerHTML = strings.Replace(innerHTML, ",", "", -1)
	innerHTML = strings.Replace(innerHTML, " ", "", -1)
	var xyz []character
	for _, c := range innerHTML {
		if c == 0 {
			continue
		} else {
			z := strings.Count(innerHTML, string(c))
			//fmt.Println(reflect.TypeOf(c), string(c))
			xyz = append(xyz, character{string(c), z})

		}
	}
	//fmt.Println(xyz)
	// stringSlice := stripResponse(innerHTML)
	// charOnlySlice := charSlice(stringSlice)
	charSliceSort := charSort(xyz)
	fmt.Println(charSliceSort)

	// if 12 == 23 {
	// 	fmt.Println(stringSlice)
	// 	fmt.Println(charSliceSort)
	// 	fmt.Println(innerHTML)
	// }
	//fmt.Println(charOnlySlice)

}

// func stripResponse(responseString string) []string {
// 	htmlSlice := strings.Split(responseString, " ")
// 	return htmlSlice
// }

// //function that takes a slice and makes a slice of chars out of it.
// func charSlice(slc []string) []character {

// 	var newCharSlice []character
// 	for _, s := range slc {
// 		for _, c := range string(s) {
// 			newCharSlice = append(newCharSlice, character{string(c), 0})
// 		}
// 	}

// 	return newCharSlice
// }

// //looking to compare newCharSlice[index].char
func charSort(slc []character) []character {
	sort.SliceStable(slc, func(i, j int) bool {
		return slc[i].count > slc[j].count
	})
	//fmt.Println("By Char:", slc)
	return slc
}
