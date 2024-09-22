package main

import (
	"fmt"
	"runtime"
)

func main() {
	if runtime.GOOS != "windows" {
		fmt.Println("This program can only be used on Windows.")
		return
	}

	initializeApp()
}
