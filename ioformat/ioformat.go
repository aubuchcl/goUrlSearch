package ioformat

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//FormatIO formats the read data from the passed url
func FormatIO(s string) (string, uint) {
	client := &http.Client{}
	resp, err := client.Get(s)
	if err != nil {
		fmt.Println("you broke it")
	}

	bs, ioErr := ioutil.ReadAll(resp.Body)
	if ioErr != nil {
		fmt.Println("you broke it from IO", ioErr)
		fmt.Println("b was ", bs)
	}

	strippedHTML := stripResponse(bs)
	mappedChars := mapChars(strippedHTML)
	freqChar, freqCharCount := sortChars(mappedChars)

	resp.Body.Close()
	return freqChar, freqCharCount

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
