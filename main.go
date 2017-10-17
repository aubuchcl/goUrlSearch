package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sort"
)

func main() {

	var useURL string

	client := &http.Client{}
	cliArgs := os.Args

	for _, u := range cliArgs {
		if isValidURL(u) == true {
			useURL = u
		}
	}

	bs := make([]byte, 100000)

	resp, err := client.Get(useURL)
	//resp.Body.Read(bs)
	if err != nil {
		fmt.Println("you broke it")
	}
	_, ioErr := io.ReadFull(resp.Body, bs)
	//	fmt.Println(b, err, bs)
	if ioErr != nil {
		fmt.Println("you broke it")
	}

	regxp, err := regexp.Compile(`<(?:[^>=]|='[^']*'|="[^"]*"|=[^'"][^\s>]*)*>`)
	strippedHTML := regxp.ReplaceAllString(string(bs), "")

	// var xyz []character.Character
	// for _, c := range strippedHTML {
	// 	if c == 0 {
	// 		continue
	// 	} else {
	// 		z := strings.Count(strippedHTML, string(c))
	// 		xyz = append(xyz, character.Character{string(c), z})

	// 	}
	// }

	chars := make(map[string]int)

	for _, v := range strippedHTML {

		if v == 0 {
			continue
		}
		if _, ok := chars[string(v)]; !ok {
			chars[string(v)] = 1
		} else {
			chars[string(v)]++
		}
	}

	fmt.Println(chars)

	// charSliceSort := character.CharSort(xyz)

	// mostChar := charSliceSort[0].Char
	// numChar := charSliceSort[0].Count
	// fmt.Println(mostChar, "occurs", numChar, "times")

	serveScan()
	resp.Body.Close()
}

//IsValidURL use this to check if a url is valid
func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	return true

}

//CharSort use to sort stripped and itemized character slices
func CharSort(slc map[string]int) map[string]int {
	//redo this function with regex
	sort.SliceStable(slc, func(i, j int) bool {
		return slc[i].Count > slc[j].Count
	})
	//fmt.Println("By Char:", slc)
	return slc
}
