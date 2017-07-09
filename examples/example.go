package main

import (
	"fmt"

	"github.com/makebyte/mutago"
)

func main() {
	f, err := mutago.Open("ok.mp3")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	val, err := f.Get("TIT2") // Get title value
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)

	tags := f.List() // Get all available tags (Not values)
	for _, y := range tags {
		fmt.Println(y) // Print all tags
	}
}
