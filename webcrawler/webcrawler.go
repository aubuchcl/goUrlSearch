package webcrawler

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

//FormatIO formats the read data from the passed url
func FormatIO(s string) (string, uint) {
	client := &http.Client{}
	resp, err := client.Get(s)
	if err != nil {
		fmt.Println("you broke it")
	}

	bs := readURL(resp.Body)

	strippedHTML := stripResponse(bs)
	mappedChars := mapChars(strippedHTML)
	freqChar, freqCharCount := sortChars(mappedChars)

	resp.Body.Close()
	return freqChar, freqCharCount

}

func readURL(rc io.ReadCloser) []byte {
	bs, err := ioutil.ReadAll(rc)
	if err != nil {
		fmt.Println("Body could not be read ", err)
	}
	return bs
}

func stripResponse(bs []byte) string {
	regxp, err := regexp.Compile(`<(?:[^>=]|='[^']*'|="[^"]*"|=[^'"][^\s>]*)*>`)
	if err != nil {
		fmt.Println("there was a regexp compile error ", err)
	}

	strippedHTML := regxp.ReplaceAllString(string(bs), "")
	return strippedHTML
}

func mapChars(s string) map[string]uint {
	chars := make(map[string]uint)
	for _, v := range s {
		if v == 0 {
			continue
		}
		if _, ok := chars[string(v)]; !ok {
			chars[string(v)] = 1
		} else {
			chars[string(v)]++
		}
	}
	return chars
}

func sortChars(m map[string]uint) (string, uint) {
	s := ""
	i := uint(0)
	abcs := "abcdefghijklmnopqrstuvwxyz"
	for key, x := range m {
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

//ServeScan will serve up the opportunity to input from command line until you type in close
func ServeScan() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "close" {
			os.Exit(1)
		}
		fmt.Println(scanner.Text())
	}
}
