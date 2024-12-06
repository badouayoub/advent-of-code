package main

import (
	"fmt"
	"os"
	"strings"
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

	var matrix [][]string
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		chars := strings.Split(line, "")
		if len(chars) > 0 {
			matrix = append(matrix, chars)
		}
	}

	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "A" {
				if isValidNE(matrix, i, j) && isValidSE(matrix, i, j) {
					sum++
				}
			}
		}
	}
	fmt.Println(sum)

}

func boundTopLeft(i, j int) bool {
	if i > 0 && j > 0 {
		return true
	}
	return false
}
func boundTopRight(matrix [][]string, i, j int) bool {
	if i > 0 && j < len(matrix[0])-1 {
		return true
	}
	return false
}

func boundBottomRight(matrix [][]string, i, j int) bool {
	if i < len(matrix)-1 && j < len(matrix[0])-1 {
		return true
	}
	return false
}

func boundBottomLeft(matrix [][]string, i, j int) bool {
	if i < len(matrix)-1 && j > 0 {
		return true
	}
	return false
}

func isInBounds(matrix [][]string, i, j int) bool {
	if boundTopLeft(i, j) && boundTopRight(matrix, i, j) && boundBottomLeft(matrix, i, j) && boundBottomRight(matrix, i, j) {
		return true
	}
	return false
}

func isValidNE(matrix [][]string, i, j int) bool {
	if isInBounds(matrix, i, j) {
		word := matrix[i+1][j-1] + matrix[i][j] + matrix[i-1][j+1]
		if (word == "MAS") || (word == "SAM") {
			return true
		}
	}
	return false
}

func isValidSE(matrix [][]string, i, j int) bool {
	if isInBounds(matrix, i, j) {
		word := matrix[i-1][j-1] + matrix[i][j] + matrix[i+1][j+1]
		if (word == "MAS") || (word == "SAM") {
			return true
		}
	}
	return false
}
