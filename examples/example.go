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

	val, err := f.Get("TALB")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)
}
