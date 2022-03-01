package main

import (
	"log"
	"os"

	v1 "github.com/kalbhor/mutago/v1"
)

func main() {
	f, err := os.OpenFile("sample.mp3", os.O_RDWR, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}

	m := v1.New(f)

	if err := m.Parse(); err != nil {
		log.Fatal(err)
	}

	if err := m.SetArtist("New Artist"); err != nil {
		log.Fatal(err)
	}

	if err := m.Parse(); err != nil {
		log.Fatal(err)
	}
}
