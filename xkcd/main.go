// see: https://github.com/nishanths/go-xkcd/blob/master/README.md
package main

import (
	"fmt"
	"log"

	"github.com/nishanths/go-xkcd"
)

func main() {
	client := xkcd.NewClient()
	comic, err := client.Latest()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n\n", comic)
	fmt.Println(comic.ImageURL)
}
