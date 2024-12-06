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
	// for i,j index in matrix
	//  if character == "X"
	//      SearchFunctionsCardinalDir(matrix, i, j) -> bool
	//          N, NW, W, SW, S, SE, E, NE
	//          if !isOOB
	//              sum+=1 when case true
	//
	//

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
			if matrix[i][j] == "X" {
				if isValidN(matrix, i, j) {
					sum++
				}
				if isValidNE(matrix, i, j) {
					sum++
				}
				if isValidE(matrix, i, j) {
					sum++
				}
				if isValidSE(matrix, i, j) {
					sum++
				}
				if isValidS(matrix, i, j) {
					sum++
				}
				if isValidSW(matrix, i, j) {
					sum++
				}
				if isValidW(matrix, i, j) {
					sum++
				}
				if isValidNW(matrix, i, j) {
					sum++
				}
			}
		}
	}
	fmt.Println(sum)

}

func isValidN(matrix [][]string, i, j int) bool {
	if i > 2 {
		word := matrix[i][j] + matrix[i-1][j] + matrix[i-2][j] + matrix[i-3][j]
		if word == "XMAS" {
			return true
		}
	}
	return false
}

func isValidNE(matrix [][]string, i, j int) bool {
	if (j < len(matrix[0])-2-1) && (i > 2) {
		word := matrix[i][j] + matrix[i-1][j+1] + matrix[i-2][j+2] + matrix[i-3][j+3]
		if word == "XMAS" {
			return true
		}
	}
	return false
}

func isValidE(matrix [][]string, i, j int) bool {
	if j < len(matrix[0])-2-1 {
		word := matrix[i][j] + matrix[i][j+1] + matrix[i][j+2] + matrix[i][j+3]
		if word == "XMAS" {
			return true
		}
	}
	return false
}

func isValidSE(matrix [][]string, i, j int) bool {
	if (j < len(matrix[0])-2-1) && (i < len(matrix)-1-2) {
		word := matrix[i][j] + matrix[i+1][j+1] + matrix[i+2][j+2] + matrix[i+3][j+3]
		if word == "XMAS" {
			return true

		}
	}
	return false
}

func isValidS(matrix [][]string, i, j int) bool {
	if i < len(matrix)-1-2 {
		word := matrix[i][j] + matrix[i+1][j] + matrix[i+2][j] + matrix[i+3][j]
		if word == "XMAS" {
			return true
		}
	}
	return false
}

func isValidSW(matrix [][]string, i, j int) bool {
	if (i < len(matrix)-1-2) && (j > 2) {
		word := matrix[i][j] + matrix[i+1][j-1] + matrix[i+2][j-2] + matrix[i+3][j-3]
		if word == "XMAS" {
			return true
		}
	}
	return false
}

func isValidW(matrix [][]string, i, j int) bool {
	if j > 2 {
		word := matrix[i][j] + matrix[i][j-1] + matrix[i][j-2] + matrix[i][j-3]
		if word == "XMAS" {
			return true
		}
	}
	return false
}

func isValidNW(matrix [][]string, i, j int) bool {
	if (j > 2) && (i > 2) {
		word := matrix[i][j] + matrix[i-1][j-1] + matrix[i-2][j-2] + matrix[i-3][j-3]
		if word == "XMAS" {
			return true
		}
	}
	return false
}
