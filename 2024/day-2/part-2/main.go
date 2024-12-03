package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func fileToString(pathToFile string) string {
	file, err := os.ReadFile(pathToFile)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func strToIntArr(arr []string) []int {
	var levelsInt []int
	for _, v := range arr {
		convertedInt, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		levelsInt = append(levelsInt, convertedInt)
	}
	return levelsInt
}

func boolToInt(b bool) int {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}

func main() {
	data := fileToString("../input.txt")
	reports := strings.Split(data, "\n")
	safeCount := 0
	for _, report := range reports {
		levelsStr := strings.Split(report, " ")
		levelsInt := strToIntArr(levelsStr)

		if len(levelsInt) > 0 {
			if isValid(levelsInt) {
				safeCount++
			} else if canBeValid(levelsInt) {
				safeCount++
			}
		}
	}
	fmt.Println(safeCount)
}

func canBeValid(report []int) bool {
	for i := range report {
		newReport := make([]int, len(report))
		copy(newReport, report)
		slices.Delete(newReport, i, i+1)
		newReport = newReport[:len(newReport)-1]
		if isValid(newReport) {
			return true
		}
	}
	return false
}

func isValid(report []int) bool {
	safe := true
	for i := 0; i < len(report)-1; i++ {
		diff := math.Abs(float64(report[i]) - float64(report[i+1]))
		if diff < 1 || diff > 3 {
			safe = false
		}
	}

	if !isIncreasing(report) && !isDecreasing(report) {
		safe = false
	}
	return safe
}

func isIncreasing(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		if report[i] <= report[i+1] {
			return false
		}
	}
	return true
}

func isDecreasing(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		if report[i] >= report[i+1] {
			return false
		}
	}
	return true
}
