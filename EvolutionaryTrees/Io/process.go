package Io

import (
  "fmt"
	"encoding/csv"
  "os"
  "strconv"
	"bufio"
	"strings"
	"EvolutionaryTrees/Lib"
	"EvolutionaryTrees/Datatypes"
)

// /Users/kunaljoshi/go/src/EvolutionaryTrees/Datasets/Input/SARS-Cov-2/distance.mtx

func WriteToDistanceMatrix(filePath string, fileDest string, setting string) {

	var mtx Datatypes.DistanceMatrix
	var labels []string

	switch setting {
		case "W":
			var dnaMap = ReadDNAStringsFromFile(filePath)
			var dnaLabels, dnaStrings = Lib.GetKeyValues(dnaMap)
			labels = dnaLabels
			mtx = Lib.CalcDistanceMatrix(dnaStrings)
		case "F":
			var dnaMap = Lib.CreateFrequencyDNAMap(ReadStringsFromFile(filePath))
			var dnaLabels, dnaStrings = Lib.GetKeyValues(dnaMap)
		  labels = dnaLabels
			mtx = Lib.CalcDistanceMatrix(dnaStrings)
		default:
			panic("Not valid setting.")
	}

  F, err := os.Create(fileDest + "/distance.mtx")
  if err != nil {
        fmt.Println(err)
        return
  }

  accum := ""
  for i := range mtx {
    accum += labels[i] + "\t"
    for j := range mtx {
      accum += fmt.Sprintf("%.2f", mtx[i][j]) + "\t"
    }
		accum = accum[0:len(accum)-1]
    accum += "\n"
  }

  l, err := F.WriteString(strconv.Itoa(len(mtx)) + "\n" + accum)
  if err != nil {
        fmt.Println(err)
        return
  }
	fmt.Println(l, "bytes written successfully")
  err = F.Close()
  if err != nil {
      fmt.Println(err)
      return
  }

}

func WriteAlignmentToFile(algn Datatypes.Alignment, labels []string, fileDest string, fileName string) {

	accum :=  ""
	for i, string := range algn {
		accum += ">" + labels[i] + "\n"
		accum += string + "\n"
	}

	F, err := os.Create(fileDest + "/" + fileName)
  if err != nil {
        fmt.Println(err)
        return
  }

  l, err := F.WriteString(accum)
  if err != nil {
        fmt.Println(err)
        return
  }
	fmt.Println(l, "bytes written successfully")
  err = F.Close()
  if err != nil {
      fmt.Println(err)
      return
  }

}

func WriteNewickToFile(t Datatypes.Tree, fileDest string, fileName string) {

	newickString := ToNewick(t)

	F, err := os.Create(fileDest + "/" + fileName)
  if err != nil {
        fmt.Println(err)
        return
  }

  l, err := F.WriteString(newickString)
  if err != nil {
        fmt.Println(err)
        return
  }
	fmt.Println(l, "bytes written successfully")
  err = F.Close()
  if err != nil {
      fmt.Println(err)
      return
  }

}

func WriteCSVToFile(t Datatypes.Tree, cats []string, fileDest string, fileName string) {

	csvString := CreateCSV(t, cats)

	F, err := os.Create(fileDest + "/" + fileName)
  if err != nil {
        fmt.Println(err)
        return
  }

  l, err := F.WriteString(csvString)
  if err != nil {
        fmt.Println(err)
        return
  }
	fmt.Println(l, "bytes written successfully")
  err = F.Close()
  if err != nil {
      fmt.Println(err)
      return
  }

}

func CreateDistanceMatrix (filePath string, setting string) (Datatypes.DistanceMatrix, []string) {

	var mtx Datatypes.DistanceMatrix
	var labels []string

	switch setting {
		case "W":
			var dnaMap = ReadDNAStringsFromFile(filePath)
			var dnaLabels, dnaStrings = Lib.GetKeyValues(dnaMap)
			labels = dnaLabels
			mtx = Lib.CalcDistanceMatrix(dnaStrings)
		case "F":
			var dnaMap = Lib.CreateFrequencyDNAMap(ReadStringsFromFile(filePath))
			var dnaLabels, dnaStrings = Lib.GetKeyValues(dnaMap)
		  labels = dnaLabels
			mtx = Lib.CalcDistanceMatrix(dnaStrings)
		default:
			panic("Not valid setting.")
	}

	return mtx, labels

}

func ProcessBacteria(filePath string) {

	file, err := os.Open(filePath)
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

	accum := ""
	freqDict := make(map[string]int, 0)
	for _, line := range lines {

		r := csv.NewReader(strings.NewReader(line))
	 	r.Comma = ' ' // space
	 	fields, err := r.Read()
	 	if err != nil {
			 fmt.Println(err)
			 return
	 	}

		label := fields[2] + "|" + fields[5]
    if label == "NA|NA" || strings.Contains(label, "|NA") {
			continue
		}

		num := 1
		_, exists := freqDict[label]
		if (exists) {
			num = freqDict[label]
			freqDict[label]++
		} else {
			freqDict[label] = 2
		}

		accum = accum + fields[2] + "|" + fields[5] + "|" + strconv.Itoa(num) + "\n"
		accum = accum + fields[0] + "\n"
	}
	fmt.Println(accum)

}
