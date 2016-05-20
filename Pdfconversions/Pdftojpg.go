package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := "convert"
	args := []string{"-verbose", "-density 150", "-trim", "lect1.pdf", "qualtiy 100", "-sharpen 0x1.0", "ExampleSlide%d.jpg"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Successfully converted pdf to images")
}
