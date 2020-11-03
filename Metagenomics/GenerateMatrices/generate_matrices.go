package main

import (
	"Metagenomics/MultipleBetaDiversity"
	"Metagenomics/MultipleEvenness"
	"Metagenomics/MultipleRichness"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	path, _ := os.Getwd()
	pathF := strings.Replace(path, "GenerateMatrices", filepath.Join("GenerateMatrices/Data/", "FallSamples"), 1)
	pathW := strings.Replace(path, "GenerateMatrices", filepath.Join("GenerateMatrices/Data/", "WinterSamples"), 1)
	pathSp := strings.Replace(path, "GenerateMatrices", filepath.Join("GenerateMatrices/Data/", "SpringSamples"), 1)
	pathSu := strings.Replace(path, "GenerateMatrices", filepath.Join("GenerateMatrices/Data/", "SummerSamples"), 1)

	mapsF, namesF := ReadFreqMapsFromDirectory(pathF)
	mtxBC := MultipleBetaDiversity.BetaDiversityMatrix(mapsF, "Bray-Curtis")
	WriteBDMatrixToFile(mtxBC, namesF, "Bray-Curtis", "Output/Bray-Curtis Fall.txt")
	mtxJ := MultipleBetaDiversity.BetaDiversityMatrix(mapsF, "Jaccard")
	WriteBDMatrixToFile(mtxJ, namesF, "Jaccard", "Output/Jaccard Fall.txt")

	mapsW, namesW := ReadFreqMapsFromDirectory(pathW)
	mtxBC = MultipleBetaDiversity.BetaDiversityMatrix(mapsW, "Bray-Curtis")
	WriteBDMatrixToFile(mtxBC, namesW, "Bray-Curtis", "Output/Bray-Curtis Winter.txt")
	mtxJ = MultipleBetaDiversity.BetaDiversityMatrix(mapsW, "Jaccard")
	WriteBDMatrixToFile(mtxJ, namesW, "Jaccard", "Output/Jaccard Winter.txt")

	mapsSp, namesSp := ReadFreqMapsFromDirectory(pathSp)
	mtxBC = MultipleBetaDiversity.BetaDiversityMatrix(mapsSp, "Bray-Curtis")
	WriteBDMatrixToFile(mtxBC, namesSp, "Bray-Curtis", "Output/Bray-Curtis Spring.txt")
	mtxJ = MultipleBetaDiversity.BetaDiversityMatrix(mapsSp, "Jaccard")
	WriteBDMatrixToFile(mtxJ, namesSp, "Jaccard", "Output/Jaccard Spring.txt")

	mapsSu, namesSu := ReadFreqMapsFromDirectory(pathSu)
	mtxBC = MultipleBetaDiversity.BetaDiversityMatrix(mapsSu, "Bray-Curtis")
	WriteBDMatrixToFile(mtxBC, namesSu, "Bray-Curtis", "Output/Bray-Curtis Summer.txt")
	mtxJ = MultipleBetaDiversity.BetaDiversityMatrix(mapsSu, "Jaccard")
	WriteBDMatrixToFile(mtxJ, namesSu, "Jaccard", "Output/Jaccard Summer.txt")

	mtxE := MultipleEvenness.SimpsonsMatrix(mapsF)
	WriteEMatrixToFile(mtxE, namesF, "Output/Evenness Fall.txt")
	mtxE = MultipleEvenness.SimpsonsMatrix(mapsW)
	WriteEMatrixToFile(mtxE, namesW, "Output/Evenness Winter.txt")
	mtxE = MultipleEvenness.SimpsonsMatrix(mapsSp)
	WriteEMatrixToFile(mtxE, namesSp, "Output/Evenness Spring.txt")
	mtxE = MultipleEvenness.SimpsonsMatrix(mapsSu)
	WriteEMatrixToFile(mtxE, namesSu, "Output/Evenness Summer.txt")

	mtxR := MultipleRichness.RichnessMatrix(mapsF)
	WriteRMatrixToFile(mtxR, namesF, "Output/Richness Fall.txt")
	mtxR = MultipleRichness.RichnessMatrix(mapsW)
	WriteRMatrixToFile(mtxR, namesW, "Output/Richness Winter.txt")
	mtxR = MultipleRichness.RichnessMatrix(mapsSp)
	WriteRMatrixToFile(mtxR, namesSp, "Output/Richness Spring.txt")
	mtxR = MultipleRichness.RichnessMatrix(mapsSu)
	WriteRMatrixToFile(mtxR, namesSu, "Output/Richness Summer.txt")
}

// ReadFreqMapsFromDirectory with the following input and output
// Input: a collection of filenames. Each file has one string on each line
// Output: a collection of frequency maps and the associated sample names
// NOTE: Provided with ReadFromFile() which takes in a file name and returns the
// frequency map for the file's contents
func ReadFreqMapsFromDirectory(directory string) ([](map[string]int), []string) {

	dirContents, err := ioutil.ReadDir(directory)
	if err != nil {
		panic("Error reading directory!")
	}

	sampleNames := make([]string, 0)
	allMaps := make([](map[string]int), 0)

	for _, fileData := range dirContents {
		// Remove the file extension
		sampleNames = append(sampleNames, (strings.Replace(fileData.Name(), ".txt", "", 1)))
		allMaps = append(allMaps, ReadFreqMapFromFile(filepath.Join(directory, fileData.Name())))
	}

	return allMaps, sampleNames
}

////////////////////////////////////////////////////////////////////////////////
// CODE PROVIDED BY INSTRUCTORS
////////////////////////////////////////////////////////////////////////////////

//ReadFreqMapFromFile
// Input: name of file which contains one string on each line
// Output: a frequency map corresponding to the strings in the file
func ReadFreqMapFromFile(filename string) map[string]int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	patterns := strings.Split(strings.Replace(strings.TrimSpace(string(data)), "\r\n", "\n", -1), "\n")
	freqMap := make(map[string]int)
	for _, val := range patterns {
		freqMap[val]++
	}
	return freqMap
}

func WriteBDMatrixToFile(matrix [][]float64, sampleNames []string, distMetric string, filename string) {
	res := ""
	for _, name := range sampleNames {
		res += name + "\t"
	}
	res = strings.TrimSpace(res)

	for i, _ := range sampleNames {
		res += "\n"
		for j, _ := range sampleNames {
			res = res + strconv.FormatFloat(matrix[i][j], 'f', 3, 64) + "\t"
		}
		res = strings.TrimSpace(res)
	}
	ioutil.WriteFile(filename, []byte(res), 0644)
}

func WriteEMatrixToFile(matrix []float64, sampleNames []string, filename string) {
	res := ""
	for _, sample := range sampleNames {
		res = res + sample + "\t"
	}
	res = res + "\n"
	for _, evenness := range matrix {
		res = res + strconv.FormatFloat(evenness, 'f', 3, 64) + "\t"
	}
	ioutil.WriteFile(filename, []byte(res), 0644)
}

func WriteRMatrixToFile(matrix []int, sampleNames []string, filename string) {
	res := ""
	for _, sample := range sampleNames {
		res = res + sample + "\t"
	}
	res = res + "\n"
	for _, richness := range matrix {
		res = res + strconv.Itoa(richness) + "\t"
	}
	ioutil.WriteFile(filename, []byte(res), 0644)
}
