package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	file := args[1]
	open, err := os.Open(file)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, open) //opened file is type File, that has https://pkg.go.dev/os#File.Read function Read therefore can be classified as member of Reader interface
}
