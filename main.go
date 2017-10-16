package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/aubuchcl/httpParser/character"
)

func main() {

	//useURL := urlstring.Urlstring{}
	var useURL map[string]string

	client := &http.Client{}
	cliArgs := os.Args

	for _, u := range cliArgs {
		if isValidURL(u) == true {
			useURL["url"] = u
		}
	}

	bs := make([]byte, 32*500)

	//resp, err := client.Get("http://golang.org")
	// resp, err := client.Get("http://www.lipsum.com")
	resp, err := client.Get(useURL["url"])
	resp.Body.Read(bs)
	//fmt.Println(string(bs))

	regxp, err := regexp.Compile(`<(?:[^>=]|='[^']*'|="[^"]*"|=[^'"][^\s>]*)*>`)
	strippedHTML := regxp.ReplaceAllString(string(bs), "")

	if err != nil {
		fmt.Println("you broke it")
	}

	var xyz []character.Character
	for _, c := range strippedHTML {
		if c == 0 {
			continue
		} else {
			z := strings.Count(strippedHTML, string(c))
			//fmt.Println(reflect.TypeOf(c), string(c))
			xyz = append(xyz, character.Character{string(c), z})

		}
	}

	charSliceSort := character.CharSort(xyz)

	mostChar := charSliceSort[0].Char
	numChar := charSliceSort[0].Count
	fmt.Println(mostChar, "occurs", numChar, "times")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "close" {
			resp.Body.Close()
			os.Exit(1)
		}
		fmt.Println(scanner.Text())
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
