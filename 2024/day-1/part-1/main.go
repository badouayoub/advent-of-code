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

func parseDataToSortedArr(data string) ([]int, []int) {
	var list1 []int
	var list2 []int

	rows := strings.Split(data, "\n")
	for _, value := range rows {
		ids := strings.Split(value, "   ")
		for idx, id := range ids {
			intId, err := strconv.Atoi(id)
			if err != nil {
				continue
			} else if idx == 0 {
				list1 = append(list1, intId)
			} else if idx == 1 {
				list2 = append(list2, intId)
			}
		}
	}
	slices.Sort(list1)
	slices.Sort(list2)
	return list1, list2
}

func main() {
	data := fileToString("../input.txt")
	list1, list2 := parseDataToSortedArr(data)

	totalDistance := 0
	for i := range list1 {
		sum := list1[i] - list2[i]
		totalDistance += int(math.Abs(float64(sum)))
	}
	fmt.Println(totalDistance)
}
