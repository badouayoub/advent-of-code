package main

import (
	"fmt"
	"math"
	"os"
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
	unsafeCount := 0
	for _, report := range reports {
		levelsStr := strings.Split(report, " ")
		levelsInt := strToIntArr(levelsStr)

		sign := 0
		if len(levelsInt) > 0 {
			for i := 0; i < len(levelsInt)-1; i++ {
				diff := levelsInt[i] - levelsInt[i+1]
				if i == 0 {
					sign = boolToInt(diff > 0) - boolToInt(diff < 0)
				}
				diffAbs := int(math.Abs(float64(diff)))
				if (sign * diff) < 0 {
					unsafeCount++
					break
				} else if diffAbs < 1 || diffAbs > 3 {
					unsafeCount++
					break
				}
			}
		}
	}
	fmt.Println(len(reports) - 1 - unsafeCount)
}
