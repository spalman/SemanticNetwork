package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var baseTerms [50]string
var baseRelations [10]relation
var adjacencyMatrix [50][50]int

type relation struct {
	name    string
	relType int
}

func parseTxt(fileName string) {
	file, _ := os.Open(fileName)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() //should skip "#1"
	for scanner.Scan() {
		if scanner.Text() != "#2" {
			temp := strings.Split(scanner.Text(), ":")
			i, _ := strconv.Atoi(temp[0])
			baseTerms[i] = temp[1]
		} else {
			break
		}
	}
	for scanner.Scan() {
		if scanner.Text() != "#3" {
			temp := strings.Split(scanner.Text(), ":")
			i, _ := strconv.Atoi(temp[0])
			t, _ := strconv.Atoi(temp[2])
			baseRelations[i] = relation{name: temp[1], relType: t}
		} else {
			break
		}
	}
	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), ":")
		i, _ := strconv.Atoi(temp[0])
		j, _ := strconv.Atoi(temp[1])
		k, _ := strconv.Atoi(temp[2])
		adjacencyMatrix[i][k] = j
	}

}
func main() {
	parseTxt("1.txt")

}
