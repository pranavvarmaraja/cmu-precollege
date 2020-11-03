package Io

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"EvolutionaryTrees/Datatypes"
)

func SequenceOrder(T Datatypes.Tree) []string {
	return ReturnSequenceOrder(T[len(T)-1])
}

func ReturnSequenceOrder(node *Datatypes.Node) []string {
	if (node.Child1 == nil) {
		return []string{node.Label}
	} else {
		return append(ReturnSequenceOrder(node.Child1), ReturnSequenceOrder(node.Child2)...)
	}
}

func CreateCSV(tree Datatypes.Tree, categories []string) string {
	var freqDict = make(map[string]int, 0)
	i, c := 1, 0;
	var count = &c
	for _, item := range categories {
		freqDict[item] = i
		i += 1
	}
	return "##,continent,color\n" + subtreeCSV(tree[len(tree) - 1], freqDict, count)
}

func subtreeCSV (node *Datatypes.Node, freqDict map[string]int, count *int) string {
	if (node.Child1 == nil) {
		*count = *count + 1
		var category = getCatFASTA(node.Label)
		return strconv.Itoa(*count) + "," + category + "," + strconv.Itoa(freqDict[category]) + "\n"
	} else {
		return subtreeCSV(node.Child1, freqDict, count) + subtreeCSV(node.Child2, freqDict, count)
	}
}

func getCatFASTA(annotation string) string {
	var bars = strings.Split(annotation, "|")
	return bars[0]
}

func ToNewick(tree Datatypes.Tree) string {
	return "(" + subtreeNewick(tree[len(tree) - 1]) + ");"
}

func ToNewickL(tree Datatypes.Tree) string {
	return "(" + subtreeNewickL(tree[len(tree) - 1]) + ");"
}

func ToNewickAges(tree Datatypes.Tree) string {
	return "(" + subtreeNewickAges(tree[len(tree) - 1]) + ");"
}

func subtreeNewick (node *Datatypes.Node) string {
	if node.Child1 == nil {
		return node.Label
	} else {
		return "(" + subtreeNewick(node.Child1) + "," + subtreeNewick(node.Child2) + ")"
	}
}

func subtreeNewickL (node *Datatypes.Node) string {
	if node.Child1 == nil {
		return node.Label
	} else {
		return "(" + subtreeNewickL(node.Child2) + "," + subtreeNewickL(node.Child1) + ")"
	}
}

func subtreeNewickAges (node *Datatypes.Node) string {
	if node.Child1 == nil {
		return node.Label + ":" + fmt.Sprintf("%.2f", node.Age)
	} else {
		return "(" + subtreeNewickAges(node.Child1) + "," + subtreeNewickAges(node.Child2) + "):" + fmt.Sprintf("%.2f", node.Age)
	}
}


func ReadDNAStringsFromFile(fileName string) (map[string]string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: couldn't open the file")
		os.Exit(1)
	}
	var lines []string = make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println("Sorry: there was some kind of error during the file reading")
		os.Exit(1)
	}
	file.Close()

	var dnaMap = make(map[string]string, 0)
	var curLabel = ""

	for idx, str := range lines {
		if idx % 2 == 0 {
			curLabel = str
		} else {
			dnaMap[curLabel] = str
		}
	}
	return dnaMap
}

func ReadStringsFromFile(fileName string) ([]string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: couldn't open the file")
		os.Exit(1)
	}
	var lines []string = make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println("Sorry: there was some kind of error during the file reading")
		os.Exit(1)
	}
	file.Close()

	var dnaMap = make([]string, 0)
	for _, str := range lines {
		dnaMap = append(dnaMap, str)
	}
	return dnaMap
}

//ReadMatrixFromFile takes a file name and reads the information in this file to produce
//a distance matrix and a slice of strings holding the species names.  The first line of the
//file should contain the number of species.  Each other line contains a species name
//and its distance to each other species.
func ReadMatrixFromFile(fileName string) (Datatypes.DistanceMatrix, []string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: couldn't open the file")
		os.Exit(1)
	}
	var lines []string = make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println("Sorry: there was some kind of error during the file reading")
		os.Exit(1)
	}
	file.Close()

	mtx := make(Datatypes.DistanceMatrix, 0)
	speciesNames := make([]string, 0)

	for idx, _ := range lines {
		if idx >= 1 {
			row := make([]float64, 0)
			nums := strings.Split(lines[idx], "\t")
			for i, num := range nums {
				if i == 0 {
					speciesNames = append(speciesNames, num)
				} else {
					n, err := strconv.ParseFloat(num, 64)
					if err != nil {
						fmt.Println("Error: Wrong format of matrix!")
						os.Exit(1)
					}
					row = append(row, n)
				}
			}
			mtx = append(mtx, row)
		}
	}
	return mtx, speciesNames
}

// PrintGraphViz prints the tree in GraphViz format, where directed = true
// if we desire to print a directed graph and directed = false for an
// undirected graph.
func PrintGraphViz(t Datatypes.Tree) {
	fmt.Println("strict digraph {")
	for i := range t {
		if t[i].Child1 != nil && t[i].Child2 != nil {
			//print first edge
			fmt.Print("\"", t[i].Label, "\"")
			fmt.Print("->")
			fmt.Print("\"", t[i].Child1.Label, "\"")
			fmt.Print("[label = \"")
			fmt.Printf("%.2f", t[i].Age-t[i].Child1.Age)
			fmt.Print("\"]")
			fmt.Println()

			//print second edge
			fmt.Print("\"", t[i].Label, "\"")
			fmt.Print("->")
			fmt.Print("\"", t[i].Child2.Label, "\"")
			fmt.Print("[label = \"")
			fmt.Printf("%.2f", t[i].Age-t[i].Child2.Age)
			fmt.Print("\"]")
			fmt.Println()
		}
	}
	fmt.Println("}")
}
