package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Interfaces can be composed of other interfaces.
// For example, the ReadCloser interface is defined as
// a Reader interface and a Closer interface.
func main() {
	log.Println("*** HTTP Reader Demo ***")
	targetURL := "http://google.com"
	resp, err := http.Get(targetURL)
	if err != nil {
		log.Fatal("At http.Get()", err)
	}
	// buff := make([]byte, 2<<14) // 2^14 = 16,384
	// bytesRead, err := resp.Body.Read(buff)   ??????????
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("At resp.Body.ReadAll()", err)
	}
	// log.Println("Read", bytesRead, "bytes from", targetURL)
	log.Println(string(body))
}
