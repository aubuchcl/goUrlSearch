package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
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

	bs := make([]byte, 1000000)

	resp, err := client.Get(useURL)
	//resp.Body.Read(bs)
	if err != nil {
		fmt.Println("you broke it")
	}
	_, ioErr := io.ReadFull(resp.Body, bs)
	//	fmt.Println(b, err, bs)
	if ioErr != nil {
		fmt.Println("you broke it from IO", ioErr)
	}

	regxp, err := regexp.Compile(`<(?:[^>=]|='[^']*'|="[^"]*"|=[^'"][^\s>]*)*>`)
	strippedHTML := regxp.ReplaceAllString(string(bs), "")

	regxpClean, _ := regexp.Compile(`/[^a-zA-Z 0-9]+/g`)
	strippedHTML = regxpClean.ReplaceAllString(strippedHTML, "")
	//fmt.Println(strippedHTML)

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

	//fmt.Println(chars)

	mfChar, mfCharNum := findBig(chars)

	fmt.Println(mfChar, "occurs", mfCharNum, "times")

	serveScan()
	if ioErr == nil {
		resp.Body.Close()
	}
}

//IsValidURL use this to check if a url is valid
func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	return true

}

// //CharSort use to sort stripped and itemized character slices
// func CharSort(slc map[string]int) map[string]int {
// 	//redo this function with regex
// 	sort.SliceStable(slc, func(i, j int) bool {
// 		return slc[i].Count > slc[j].Count
// 	})
// 	//fmt.Println("By Char:", slc)
// 	return slc
// }
//hashMap find
func findBig(hm map[string]int) (string, int) {
	s := ""
	i := 0
	abcs := "abcdefghijklmnopqrstuvwxyz"
	for key, x := range hm {
		matched, _ := regexp.MatchString(key, abcs)
		if matched == false {
			continue
		}
		if x > i {
			i = x
			s = key
		}
	}
	return s, i
}
