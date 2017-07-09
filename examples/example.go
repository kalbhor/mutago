package main

import (
	"fmt"

	"github.com/makebyte/mutago"
)

func main() {
	f, err := mutago.Open("sample.mp3")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	title, _ := f.Title() // Get title value
	album, _ := f.Album()
	artist, _ := f.Artist()

	fmt.Println(title, album, artist)

	tags := f.List() // Get all available tags identifiers (Not values)
	for _, y := range tags {
		fmt.Println(y) // Print identifiers
	}

	text, _ := f.Get("TXXX") // Getting tag with identifier TXXX
	fmt.Println(text)
}
