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

func process() {
	for i := range adjacencyMatrix {
		for j := range adjacencyMatrix[i] {
			if adjacencyMatrix[i][j] == 2 || adjacencyMatrix[i][j] == 1 {
				addRelations(i, j, baseRelations[adjacencyMatrix[i][j]].relType)
			}
		}
	}
}

func addRelations(upperEdge, lowerEdge, relType int) {

	switch relType { //adjacencyMatrix[lowerEdge][i] == 2 || adjacencyMatrix[lowerEdge][i] == 1 {
	case 1:
		for i := range adjacencyMatrix[lowerEdge] {
			if adjacencyMatrix[lowerEdge][i] != 0 && baseRelations[adjacencyMatrix[lowerEdge][i]].relType != 0 {
				adjacencyMatrix[upperEdge][i] = adjacencyMatrix[lowerEdge][i]
			}
		}
	case 2:
		for i := range adjacencyMatrix[lowerEdge] {
			if adjacencyMatrix[lowerEdge][i] == 1 {
				adjacencyMatrix[upperEdge][i] = adjacencyMatrix[lowerEdge][i]
			}
		}
	}
}

func requestHandler(request string) {
	temp := strings.Split(request, ":")
	//fmt.Printf("temp: %s %s %s", temp[0], temp[1], temp[2])
	temp[2] = temp[2][0 : len(temp[2])-1]

	switch {
	// True/false request
	case temp[0] != "?" && temp[1] != "?" && temp[2] != "?":
		i0, _ := strconv.Atoi(temp[0])
		i1, _ := strconv.Atoi(temp[1])
		i2, _ := strconv.Atoi(temp[2])
		if adjacencyMatrix[i0][i2] == i1 {
			fmt.Printf("\nTrue")
		} else {
			fmt.Printf("\nFalse")
		}

		// First param unknown
		//Shows objects that has selected relation with another one
	case temp[0] == "?" && temp[1] != "?" && temp[2] != "?":
		i1, _ := strconv.Atoi(temp[1])
		i2, _ := strconv.Atoi(temp[2])
		for i := range adjacencyMatrix {
			if adjacencyMatrix[i][i2] == i1 {
				fmt.Printf("%s %s %s\n", baseTerms[i], baseRelations[i1].name, baseTerms[i2])
			}
		}

		//Fist & second unknown
		//Finds relations with second object
	case temp[0] == "?" && temp[1] == "?" && temp[2] != "?":
		i2, _ := strconv.Atoi(temp[2])
		for i := range adjacencyMatrix {
			if adjacencyMatrix[i][i2] != 0 {
				fmt.Printf("%s %s %s\n", baseTerms[i], baseRelations[adjacencyMatrix[i][i2]].name, baseTerms[i2])
			}
		}

		//Second & third param unknown
	case temp[0] != "?" && temp[1] == "?" && temp[2] == "?":
		i0, _ := strconv.Atoi(temp[0])
		for i := range adjacencyMatrix[i0] {
			if adjacencyMatrix[i0][i] != 0 {
				fmt.Printf("%s %s %s\n", baseTerms[i0], baseRelations[adjacencyMatrix[i0][i]].name, baseTerms[i])
			}
		}

		//Second param unknown
		//Shows relation between two objects
	case temp[0] != "?" && temp[1] == "?" && temp[2] != "?":
		i0, _ := strconv.Atoi(temp[0])
		i2, _ := strconv.Atoi(temp[2])
		if adjacencyMatrix[i0][i2] != 0 {
			fmt.Printf("%s %s %s\n", baseTerms[i0], baseRelations[adjacencyMatrix[i0][i2]].name, baseTerms[i2])

		} else {
			fmt.Printf("No relations between this objects\n")
		}
	//Third param unknown
	//Finds relations with first object
	case temp[0] != "?" && temp[1] != "?" && temp[2] == "?":
		i0, _ := strconv.Atoi(temp[0])
		i1, _ := strconv.Atoi(temp[1])
		for i := range adjacencyMatrix[i0] {
			if adjacencyMatrix[i0][i] == i1 {
				fmt.Printf("%s %s %s\n", baseTerms[i0], baseRelations[i1].name, baseTerms[i])
			}
		}
	//First && third unknown
	case temp[0] == "?" && temp[1] != "?" && temp[2] == "?":
		i1, _ := strconv.Atoi(temp[1])
		for i := range adjacencyMatrix {
			for j := range adjacencyMatrix {
				if adjacencyMatrix[i][j] == i1 {
					fmt.Printf("%s %s %s\n", baseTerms[i], baseRelations[adjacencyMatrix[i][j]].name, baseTerms[j])
				}
			}
		}
	//All params "?"
	//Print all data
	case temp[0] == "?" && temp[1] == "?" && temp[2] == "?":
		for i := range adjacencyMatrix {
			for j := range adjacencyMatrix[i] {
				if adjacencyMatrix[i][j] != 0 {
					fmt.Printf("%s %s %s\n", baseTerms[i], baseRelations[adjacencyMatrix[i][j]].name, baseTerms[j])
				}
			}
		}
	}
}

func main() {
	parseTxt("TestData/1.txt")
	process()
	for i := range adjacencyMatrix {
		b := false
		for j := range adjacencyMatrix[i] {
			if adjacencyMatrix[i][j] != 0 {
				fmt.Printf("(%d; %d)=%d ", i, j, adjacencyMatrix[i][j])
				b = true
			}
		}
		if b {
			fmt.Printf("\n")
		}
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	requestHandler(text)

}
