package main

import (
	"fmt"
	"os"

	"github.com/makebyte/mutago/v2"
)

func main() {
	f, _ := os.Open("ok.mp3")
	z := v2.ParseHeader(f)
	fmt.Println(z.Size)

}
