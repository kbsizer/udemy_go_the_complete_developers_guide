package main

import "fmt"

func main() {
	ints := make([]int, 11)
	for i := 0; i < cap(ints); i++ {
		ints[i] = i
	}

	for _, v := range ints {
		if v%2 == 0 {
			fmt.Println(v, "is even")
		} else {
			fmt.Println(v, "is odd")
		}
	}
}
