package main

import (
	//"os"
//    "bytes"
	"strings"
    "fmt"
    "log"
    "io/ioutil"
    //"os/exec"
    //"regexp"
    //"strings"
    "time"
    "bytes"
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

func commandSearcher(providedText []byte) (text []byte){
	//words := strings.Fields(providedText) // Fields extracts the words into a slice.
	//var formated []byte
	var buffer bytes.Buffer
	for _, word := range providedText {
		if strings.Compare(string(word), ":date:") == 0 {
			currentTime:=time.Now().Format("2006-01-02 15:04:05")
			buffer.WriteString(currentTime)
		} else {
			buffer.WriteString(string(word))
		}
	}
	text = []byte(buffer.String())
	return text
	}

	//re := regexp.MustCompile(datext)
	//dateCW := re.FindIndex([]byte(":date:")) //returns value and index of the match
	//if dateCW != nil {
		//replaces the codeword :date: with the current date
	//	currentTime:=time.Now().Format("2006-01-02 15:04:05")
		//re.ReplaceAllFunc()
	//	text[dateCW[0]:dateCW[1]] =  []byte("asdads") //reoplaces the codeword :date: with the current date
	//	}
		//if strings.EqualFold("", w)

		
	
	//return text
//}
