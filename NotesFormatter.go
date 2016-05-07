package main

import (
	"os"
    //"bytes"
    "fmt"
    "time"
    "log"
)

func main() {
	start := time.Now()
	err := CreatePng("picture.html")
	if err != nil {
			log.Fatal(err)
		}
	
	fmt.Println(time.Since(start))
}

// CreatePng creates a PNG picture file with a Julia image of size n x n.
func CreatePng(filename string) (err error) {
	os.Create(filename)
	//if err != nil {
	//	return
	//}

	//defer file.Close()
	//err = png.Encode(file, Julia(f, n))
	return
}