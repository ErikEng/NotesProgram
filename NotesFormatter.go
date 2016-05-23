package main

import (
	"os"
//    "bytes"
    "fmt"
    "log"
    "io/ioutil"
    //"bufio"
)

func main() {
	var response string
	var sourceName string
	fmt.Print("Enter the name you want for your file:")
	_ , err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("What is the full name of your text?:")
	fmt.Scanln(&sourceName)
	sourceText:= ioutil.ReadFile(sourceName)
	//d1 := []byte("<hmtl>\n<header> Subject, Date </header>\n<div> Lorem ipsum dolor sit amet. Nunc quis pellentesque nibh. Ut vitae <b>Aenean placerat </b> nisi eget iaculis facilisis.</div> \n <embed  src=\"Slides/lect1.pdf\" width=\"800px\" height=\"2100px\">")
	response += ".html"
	d2 := []byte("momomomomo")
	fmt.Print(response)
	//fmt.Print(d1)

	os.Create(response)
    ioutil.WriteFile(response, d2, 0644)


//    ioutil.WriteFile(response, d2, 0644)
	//err := CreatePng("picture.html")
	//if err != nil {
	//		log.Fatal(err)
	//	}
	 
	//fmt.Println(time.Since(start))
}
