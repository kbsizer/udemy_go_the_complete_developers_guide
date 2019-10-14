package main

import (
	"io"
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

	// // version 1: using ReadAll
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal("At resp.Body.ReadAll()", err)
	// }
	// log.Println(string(body))

	// Version 2: using Read with a byte slice buffer
	buff := make([]byte, 2<<14) // 2^14 = 16,384
	bytesRead, err := resp.Body.Read(buff)
	log.Println("resp.Body.Read(buff) : err = ", err)
	if err != nil && err != io.EOF {
		log.Fatal("At Body.Read()", err)
	}

	log.Println("Read", bytesRead, "bytes from", targetURL)
}
