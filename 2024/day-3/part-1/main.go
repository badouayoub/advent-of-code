package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func fileToString(pathToFile string) string {
	file, err := os.ReadFile(pathToFile)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func strToInt(s string) int {
	convInt, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return convInt
}

func main() {
	data := fileToString("../input.txt")

	// Loop through each character
	reFindValidMul := regexp.MustCompile("mul\\(\\d+,\\d+\\)")
	reExtractArgs := regexp.MustCompile("\\d+")
	matches := reFindValidMul.FindAllStringSubmatch(data, -1)

	sum := 0
	for _, match := range matches {
		args := reExtractArgs.FindAllStringSubmatch(match[0], -1)
		sum += strToInt(args[0][0]) * strToInt(args[1][0])
	}
	fmt.Println(sum)
}
