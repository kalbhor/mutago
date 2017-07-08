package main

import (
	"fmt"
	"os"

	"../v2"
)

func main() {
	f, _ := os.Open("ok.mp3")
	z := v2.ParseHeader(f)
	r, _ := f.Seek(0, 1)
	for int(r) < int(z.Size) {
		x := v2.ParseFrame(f)
		fmt.Println(x.ID)
		fmt.Println(x.Info)
		r, _ = f.Seek(0, 1)
		fmt.Println("\n\n")
	} // Print all tags

	/*
		Note after APIC there might be junk values since they're part of
		the album art
	*/

}
