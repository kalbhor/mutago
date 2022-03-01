package main

import (
	"fmt"

	"github.com/kalbhor/mutago"
)

func main() {
	fmt.Println("\n")

	m, err := mutago.Open("sample.mp3")
	if err != nil {
		panic(err)
	}
	if err := m.Parse(); err != nil {
		panic(err)
	}

	/* Basic methods (Title, album, artist, album art) */
	title, _ := m.Title()   // string
	album, _ := m.Album()   // string
	artist, _ := m.Artist() // string

	/* albumart struct containing:
	   - Mime (format of image) (string)
	   - TextEncoding (encoding of Description) (Byte)
	   - Description (Description of cover) (string)
	   - PictureType (According to ID3) (Byte)
	*/

	fmt.Println("Song title :", title)
	fmt.Println("Song album : ", album)
	fmt.Println("Song artist : ", artist)

}
