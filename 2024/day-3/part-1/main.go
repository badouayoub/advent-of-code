package main

import (
	"fmt"
	"os"
)

func fileToString(pathToFile string) string {
	file, err := os.ReadFile(pathToFile)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func main() {
	data := fileToString("../input.txt")
	fmt.Println(data)

	// Loop through each character
}
