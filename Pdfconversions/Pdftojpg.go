package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main(fp, lp, filename, newname string) {
	cmd := "pdfimages"
	args := []string{"-j", "-f", fp, "-l", lp, filename, newname}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Successfully converted pdf %w to images %w", filename, newname)
}
