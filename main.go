package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/aubuchcl/httpParser/strip"
)

type character struct {
	char  string
	count int
}

func main() {
	cliArgs := os.Args

	inputArgIndex := len(cliArgs) - 1

	urlArg := cliArgs[inputArgIndex]

	//fmt.Println("hello world")

	client := &http.Client{
	//CheckRedirect: http.Redirect(w ResponseWriter, r *Request, url string, code int),
	}
	bs := make([]byte, 32*500)

	//resp, err := client.Get("http://golang.org")
	// resp, err := client.Get("http://www.lipsum.com")
	resp, err := client.Get(urlArg)

	resp.Body.Read(bs)
	if err != nil {
		fmt.Println(err, "please enter a valid url")
	}
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

	charSliceSort := charSort(xyz)

	mostChar := charSliceSort[0].char
	numChar := charSliceSort[0].count
	fmt.Println(mostChar, "occurs", numChar, "times")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "close" {
			os.Exit(1)
		}
		fmt.Println(scanner.Text())
	}

}

// //looking to compare newCharSlice[index].char
func charSort(slc []character) []character {
	sort.SliceStable(slc, func(i, j int) bool {
		return slc[i].count > slc[j].count
	})
	//fmt.Println("By Char:", slc)
	return slc
}
