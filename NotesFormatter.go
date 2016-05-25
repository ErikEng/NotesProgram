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
)

func main() {
	var nameOfFile string
	var sourceName string
	fmt.Print("Enter the name you want for your file:")
	_ , err := fmt.Scanln(&nameOfFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("What is the full name of your text?:")
	fmt.Scanln(&sourceName)
	sourceText, _ := ioutil.ReadFile(sourceName)
	nameOfFile += ".html"
	finnishedText := commandSearcher(sourceText,nameOfFile)
    ioutil.WriteFile(nameOfFile, finnishedText, 0644)
}

func commandSearcher(providedByte []byte, theTitle string) (text []byte){
	var buffer bytes.Buffer
	providedText := string(providedByte)
	w := strings.Fields(providedText)
	date, _ := regexp.Compile(":date:") //create regexp for date-command
	slide, _ := regexp.Compile(":slide:")
	title, _ := regexp.Compile(":title:")
	endWord, _ := regexp.Compile(":end:")
	video, _ := regexp.Compile(":video:")
	textFormating := ("<html><head><meta charset=\"utf-8\"><metaname=\"viewport\" content=\"width=device-width, initial-scale=1\"><title>"+ theTitle +"</title><style type=\"text/css\">body{margin:40px auto;max-width:650px;line-height:1.6;font-size:18px;color:#444;padding:0 10px}  </style></head><div>")
 	buffer.WriteString(textFormating)
 	n := 0
	for n < len(w) {
		word := w[n]
 
		if date.MatchString(word) {
			currentTime:=time.Now().Format("2006-01-02 15:04:05")
			buffer.WriteString(currentTime)

		} else if slide.MatchString(word) {
			if n >= len(w) {
				fmt.Println("No slide was given after the :slide: command")
				break
			}
			buffer.WriteString("</div>")
			slideSource := "Slides/"+w[n+1] + ".jpg"//takes out the source of the slide and its name 
			htmlExpression := "<embed  src=\"" + slideSource + "\" width=\"1100px\" height=\"800px\">"
			buffer.WriteString(htmlExpression) //inserts the slide into the html text

		} else if title.MatchString(word) {
			title := ""
			for n < len(w) && !endWord.MatchString(w[n+1]) {   // loop until you've encountered newline
				title += " " + w[n+1]
				n++  //increase n so title isn't looped through and printed again
			}
			n++ //ignore the :end: command
			divExpression := "<h1>" + title + "</h1>"
			buffer.WriteString(divExpression)
		} else if video.MatchString(word) {
			url := w[n+1]
			statement := "</div> <iframe title=\"YouTube video player\" class=\"youtube-player\" type=\"text/html\" width=\"640\"  height=\"390\" src=\""+ url +"\" frameborder=\"0\" allowFullScreen></iframe> <div>"
			buffer.WriteString(statement)
			n++
		
		} else {
			buffer.WriteString(word)
		}
		n++
		buffer.WriteString(" ")
	}
	buffer.WriteString("</html>")
	text = []byte(buffer.String())
	return text
}

func splitpdf(fp, lp, filename, newname string) {
	cmd := "pdfimages"
	args := []string{"-j", "-f", fp, "-l", lp, filename, newname}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Successfully converted pdf %w to images %w", filename, newname)
}
