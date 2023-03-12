package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("outfile.txt")

	barr := []byte{1, 2, 3}

	nb, err := f.Write(barr)
	nb, err = f.WriteString("Hi")

	fmt.Println(nb, err)

}
