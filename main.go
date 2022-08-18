package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func main() {
	resp, err := http.Get("http://tools.2345.com/quwcs/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if(resp.StatusCode != http.StatusOK) {
		 fmt.Println("error: status code", resp.StatusCode)
	     return
	}
	// charset.DetermineEncoding(resp.Body)
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
    bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}      
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
