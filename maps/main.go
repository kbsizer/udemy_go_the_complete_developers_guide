package main

import "fmt"

func main() {
	// Example: a map of string-type keys to string-type values
	// method #1: zero-valued map
	var colors1 map[string]string
	fmt.Printf("colors1 = %v\n", colors1)

	// method #2: using make()
	colors2 := make(map[string]string)
	colors2["green"] = "00ff00"
	colors2["yellow"] = "ffff00"
	fmt.Printf("colors2 = %v\n", colors2)

	// method #3: create and initialize map
	colors3 := map[string]string{
		"red":    "ff0000",
		"green":  "00ff00",
		"blue":   "0000ff",
		"cyan":   "00ffff",
		"yellow": "ffff00",
		"purple": "ff00ff",
	}
	printMap(colors3)

	// deleting key:value pairs
	delete(colors3, "red")
	delete(colors3, "yellow")
	printMap(colors3)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Printf("\t%-8s : %s\n", key, value)
	}
}
