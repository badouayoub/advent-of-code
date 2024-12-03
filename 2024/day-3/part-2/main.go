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

	reFindInstructions := regexp.MustCompile("((mul\\(\\d+,\\d+\\))|(don't)|(do))")
	reExtractArgs := regexp.MustCompile("\\d+")
	matches := reFindInstructions.FindAllStringSubmatch(data, -1)

	sum := 0
	enabled := true
	for _, match := range matches {
		instruction := match[0]
		if instruction == "don't" {
			enabled = false
			continue
		} else if instruction == "do" {
			enabled = true
			continue
		}
		for enabled {
			args := reExtractArgs.FindAllStringSubmatch(instruction, -1)
			sum += strToInt(args[0][0]) * strToInt(args[1][0])
			break
		}
	}
	fmt.Println(sum)
}
