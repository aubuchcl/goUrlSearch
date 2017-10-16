package urlstring

import (
	"net/url"
)

//Urlstring placeholder string type that has single property url(string)
type Urlstring struct {
	Url string
}

// NewUrlString allows urlstring struct to be initialized in main
func NewUrlString() Urlstring {
	urlstruct := Urlstring{
		Url: "",
	}
	return urlstruct
}

//IsValidURL use this to check if a url is valid
func IsValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		return true
	}
}

//FindURL setter method for url
func (s *Urlstring) FindURL(u string) {
	(*s).Url = u
	return
}
