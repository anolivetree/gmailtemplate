package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"
)

type Mail struct {
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    []string
}

func createBookmark(m Mail) string {

	ret := "mailto:"

	if len(m.To) > 0 {
		ret += strings.Join(m.To, ",")
	}
	ret += "?"

	if len(m.Cc) > 0 {
		ret += "CC="
		ret += strings.Join(m.Cc, ",")
		ret += "&"
	}

	if len(m.Bcc) > 0 {
		ret += "BCC="
		ret += strings.Join(m.Bcc, ",")
		ret += "&"
	}

	if m.Subject != "" {
		ret += "Subject="
		ret += url.PathEscape(m.Subject)
		ret += "&"
	}

	if len(m.Body) > 0 {
		ret += "Body="
		ret += url.PathEscape(strings.Join(m.Body, "\n"))
		ret += "&"
	}

	return ret
}

func main() {

	r, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	var mail Mail
	err = json.NewDecoder(r).Decode(&mail)
	if err != nil {
		panic(err)
	}

	bookmark := createBookmark(mail)
	fmt.Printf("%v\n", bookmark)

}
