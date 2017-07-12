package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/makebyte/mutago"
)

func main() {
	fmt.Println("\n")

	f, err := mutago.Open("sample.mp3")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	/* Basic methods (Title, album, artist, album art) */
	title, _ := f.Title()   // string
	album, _ := f.Album()   // string
	artist, _ := f.Artist() // string

	/* albumart struct containing:
	   - Mime (format of image) (string)
	   - TextEncoding (encoding of Description) (Byte)
	   - Description (Description of cover) (string)
	   - PictureType (According to ID3) (Byte)
	*/
	albumart := f.Albumart()

	reader := bytes.NewReader(albumart.Data)
	img, _ := os.Create("albumart.png")
	defer img.Close()

	io.Copy(img, reader)

	fmt.Println("Song title :", title)
	fmt.Println("Song album : ", album)
	fmt.Println("Song artist : ", artist)
	fmt.Println("Albumart image type : ", albumart.Mime)
	fmt.Println("Albumart description : ", albumart.Description)

	fmt.Println("\n\nSaved albumart.png")

}
