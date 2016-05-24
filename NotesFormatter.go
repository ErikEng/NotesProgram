package main

import (
	"os"
    "bytes"
    "fmt"
    "log"
    "io/ioutil"
    "strings"
    "time"
    "os/exec"
    "regexp"
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
	//d1 := []byte("<hmtl>\n<header> Subject, Date </header>\n<div> Lorem ipsum dolor sit amet. Nunc quis pellentesque nibh. Ut vitae <b>Aenean placerat </b> nisi eget iaculis facilisis.</div> \n <embed  src=\"Slides/lect1.pdf\" width=\"800px\" height=\"2100px\">")
	response += ".html"
	finnishedText := commandSearcher(sourceText)
	//d2 := []byte("momomomomo")
	//fmt.Print(response)
	//fmt.Print(d1)

	//os.Create(response)
    ioutil.WriteFile(response, finnishedText, 0644)
}

//    ioutil.WriteFile(response, d2, 0644)
	//err := CreatePng("picture.html")
	//if err != nil {
	//		log.Fatal(err)
	//	}
func commandSearcher(providedByte []byte) (text []byte){
	//words := strings.Fields(providedText) // Fields extracts the words into a slice.
	//var formated []byte
	var buffer bytes.Buffer
	providedText := string(providedByte)
	w := strings.Fields(providedText)
	date, _ := regexp.Compile(":date:") //create regexp for date-command
	slide, _ := regexp.Compile(":slide:")
	for _, word := range w {
		
 
		if date.MatchString(word) {
		//if /[:date:]/ {	
			currentTime:=time.Now().Format("2006-01-02 15:04:05")
			//fmt.Print("success" + currentTime)
			buffer.WriteString(currentTime)
		} else if slide.MatchString(word) {
			slideSource := "Slides/"+word[5:len(word)] //takes out the source of the slide and its name 
			htmlExpression=("<embed  src="+slideSource+" width="800px" height="2100px">")
			buffer.WriteString(htmlExpression) //inserts the slide into the html text
		} else {
			buffer.WriteString(word)
		}
		buffer.WriteString(" ")
	}
	text = []byte(buffer.String())
	//fmt.Println("why")
	return text
}

	//fmt.Println(time.Since(start))
func splitpdf(fp, lp, filename, newname string) {
	cmd := "pdfimages"
	args := []string{"-j", "-f", fp, "-l", lp, filename, newname}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Successfully converted pdf %w to images %w", filename, newname)
}
