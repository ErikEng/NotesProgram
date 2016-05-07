package main

import (
	"os"
    //"bytes"
    "fmt"
    "time"
    "log"
    "io/ioutil"
)

func main() {
	start := time.Now()
	d1 := []byte("<hmtl>\n<header> Subject, Date </header>\n<div> Lorem ipsum dolor sit amet. Nunc quis pellentesque nibh. Ut vitae <b>Aenean placerat </b> nisi eget iaculis facilisis.</div> \n <embed  src=\"Slides/lect1.pdf\" width=\"800px\" height=\"2100px\">")
	
	os.Create("GoTest.html")
    err := ioutil.WriteFile("GoTest.html", d1, 0644)
	//err := CreatePng("picture.html")
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