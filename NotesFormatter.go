package main

import (
	//"os"
//    "bytes"
    "fmt"
    "log"
    "io/ioutil"
    //"os/exec"
    "regexp"
    //"strings"
    "time"
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
	sourceText, _ := ioutil.ReadFile(sourceName)
	response += ".html"
	//formatedText := commandSearcher(sourceText)
	formatedByte := commandSearcher(sourceText)
	//fmt.Print(response)
	//fmt.Print(d1)
	//os.Create(response)
    ioutil.WriteFile(response, formatedByte, 0644)
   	//exec.Command("run", response).Run();
}

func commandSearcher(providedText []byte) (formated []byte){
	//words := strings.Fields(providedText) // Fields extracts the words into a slice.
	//var formated []byte
	for _, w := range providedText {
		if regexp.Match(":date:", w) {
		//replaces the codeword :date: with the current date
		currentTime:=time.Now().Format("2006-01-02 15:04:05")

		w =  []byte(currentTime)
		}
		//if strings.EqualFold("", w)

		append(formated, w...)
	}
	return formated
}
