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
func FormatIO(s string) {
	//init http client
	client := &http.Client{}

	//get raw info from s(url)
	resp, err := client.Get(s)

	//check to see if request was successful
	if err != nil {
		fmt.Println("there was an error handling your get request: ", err)
	}

	//capture the response body
	bs := readURL(resp.Body)

	//strip the response of HTML
	strippedHTML := stripResponse(bs)

	//make a hashmap of a-z chars from stripped response
	mappedChars := mapChars(strippedHTML)

	//find the most frequently used char and its frequency
	freqChar, freqCharCount := sortChars(mappedChars)

	//close response body
	resp.Body.Close()

	//notify cli user of char frequency
	fmt.Println(freqChar, "occurs", freqCharCount, "times")
	return
}

//take in an io.ReadCloser and return a byteslice of its contents
func readURL(rc io.ReadCloser) []byte {
	bs, err := ioutil.ReadAll(rc)
	if err != nil {
		fmt.Println("ioutil errored when trying to read: ", rc)
		fmt.Println("error recieved was: ", err)
	}
	return bs
}

//use regexp to remove HTML and convert byteslice to string
func stripResponse(bs []byte) string {
	regxp, err := regexp.Compile(`<(?:[^>=]|='[^']*'|="[^"]*"|=[^'"][^\s>]*)*>`)
	if err != nil {
		fmt.Println("there was a regexp compile error ", err)
	}

	strippedHTML := regxp.ReplaceAllString(string(bs), "")
	return strippedHTML
}

//turn the string into a hashmap of character frequency
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

//look through the hashmap and return the char with the highest use
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
			fmt.Println("Sure.  Closing now")
			os.Exit(0)
		}

	}
}
