package BrayCurtis

import ("strings"
		"testing"
		"path/filepath"
		"io/ioutil"
		"strconv"
		"math"
		"os")

type testpair struct {
	file1 string
	file2 string
	distance float64
}

var tests = []testpair{
	{filepath.Join("Tests","test1.txt"), filepath.Join("Tests","test1.txt"),0.000},
	{filepath.Join("Tests","test1.txt"), filepath.Join("Tests","test10.txt"),1.000},
	{filepath.Join("Tests","test1.txt"), filepath.Join("Tests","test11.txt"),1.000},
	{filepath.Join("Tests","test1.txt"), filepath.Join("Tests","test12.txt"),1.000},
	{filepath.Join("Tests","test1.txt"), filepath.Join("Tests","test13.txt"),1.000},
	{filepath.Join("Tests","test1.txt"), filepath.Join("Tests","test14.txt"),1.000},
	{filepath.Join("Tests","test1.txt"), filepath.Join("Tests","test2.txt"),1.000},
	{filepath.Join("Tests","test1.txt"), filepath.Join("Tests","test3.txt"),1.000},
	{filepath.Join("Tests","test1.txt"), filepath.Join("Tests","test4.txt"),1.000},
	{filepath.Join("Tests","test1.txt"), filepath.Join("Tests","test5.txt"),1.000},
	{filepath.Join("Tests","test1.txt"), filepath.Join("Tests","test6.txt"),1.000},
	{filepath.Join("Tests","test1.txt"), filepath.Join("Tests","test7.txt"),1.000},
	{filepath.Join("Tests","test1.txt"), filepath.Join("Tests","test8.txt"),1.000},
	{filepath.Join("Tests","test1.txt"), filepath.Join("Tests","test9.txt"),1.000},
	{filepath.Join("Tests","test10.txt"), filepath.Join("Tests","test1.txt"),1.000},
	{filepath.Join("Tests","test10.txt"), filepath.Join("Tests","test10.txt"),0.000},
	{filepath.Join("Tests","test10.txt"), filepath.Join("Tests","test11.txt"),1.000},
	{filepath.Join("Tests","test10.txt"), filepath.Join("Tests","test12.txt"),1.000},
	{filepath.Join("Tests","test10.txt"), filepath.Join("Tests","test13.txt"),1.000},
	{filepath.Join("Tests","test10.txt"), filepath.Join("Tests","test14.txt"),1.000},
	{filepath.Join("Tests","test10.txt"), filepath.Join("Tests","test2.txt"),1.000},
	{filepath.Join("Tests","test10.txt"), filepath.Join("Tests","test3.txt"),1.000},
	{filepath.Join("Tests","test10.txt"), filepath.Join("Tests","test4.txt"),1.000},
	{filepath.Join("Tests","test10.txt"), filepath.Join("Tests","test5.txt"),1.000},
	{filepath.Join("Tests","test10.txt"), filepath.Join("Tests","test6.txt"),1.000},
	{filepath.Join("Tests","test10.txt"), filepath.Join("Tests","test7.txt"),0.333},
	{filepath.Join("Tests","test10.txt"), filepath.Join("Tests","test8.txt"),0.490},
	{filepath.Join("Tests","test10.txt"), filepath.Join("Tests","test9.txt"),0.547},
	{filepath.Join("Tests","test11.txt"), filepath.Join("Tests","test1.txt"),1.000},
	{filepath.Join("Tests","test11.txt"), filepath.Join("Tests","test10.txt"),1.000},
	{filepath.Join("Tests","test11.txt"), filepath.Join("Tests","test11.txt"),0.000},
	{filepath.Join("Tests","test11.txt"), filepath.Join("Tests","test12.txt"),0.200},
	{filepath.Join("Tests","test11.txt"), filepath.Join("Tests","test13.txt"),0.209},
	{filepath.Join("Tests","test11.txt"), filepath.Join("Tests","test14.txt"),0.389},
	{filepath.Join("Tests","test11.txt"), filepath.Join("Tests","test2.txt"),1.000},
	{filepath.Join("Tests","test11.txt"), filepath.Join("Tests","test3.txt"),1.000},
	{filepath.Join("Tests","test11.txt"), filepath.Join("Tests","test4.txt"),1.000},
	{filepath.Join("Tests","test11.txt"), filepath.Join("Tests","test5.txt"),1.000},
	{filepath.Join("Tests","test11.txt"), filepath.Join("Tests","test6.txt"),1.000},
	{filepath.Join("Tests","test11.txt"), filepath.Join("Tests","test7.txt"),1.000},
	{filepath.Join("Tests","test11.txt"), filepath.Join("Tests","test8.txt"),1.000},
	{filepath.Join("Tests","test11.txt"), filepath.Join("Tests","test9.txt"),1.000},
	{filepath.Join("Tests","test12.txt"), filepath.Join("Tests","test1.txt"),1.000},
	{filepath.Join("Tests","test12.txt"), filepath.Join("Tests","test10.txt"),1.000},
	{filepath.Join("Tests","test12.txt"), filepath.Join("Tests","test11.txt"),0.200},
	{filepath.Join("Tests","test12.txt"), filepath.Join("Tests","test12.txt"),0.000},
	{filepath.Join("Tests","test12.txt"), filepath.Join("Tests","test13.txt"),0.300},
	{filepath.Join("Tests","test12.txt"), filepath.Join("Tests","test14.txt"),0.273},
	{filepath.Join("Tests","test12.txt"), filepath.Join("Tests","test2.txt"),1.000},
	{filepath.Join("Tests","test12.txt"), filepath.Join("Tests","test3.txt"),1.000},
	{filepath.Join("Tests","test12.txt"), filepath.Join("Tests","test4.txt"),1.000},
	{filepath.Join("Tests","test12.txt"), filepath.Join("Tests","test5.txt"),1.000},
	{filepath.Join("Tests","test12.txt"), filepath.Join("Tests","test6.txt"),1.000},
	{filepath.Join("Tests","test12.txt"), filepath.Join("Tests","test7.txt"),1.000},
	{filepath.Join("Tests","test12.txt"), filepath.Join("Tests","test8.txt"),1.000},
	{filepath.Join("Tests","test12.txt"), filepath.Join("Tests","test9.txt"),1.000},
	{filepath.Join("Tests","test13.txt"), filepath.Join("Tests","test1.txt"),1.000},
	{filepath.Join("Tests","test13.txt"), filepath.Join("Tests","test10.txt"),1.000},
	{filepath.Join("Tests","test13.txt"), filepath.Join("Tests","test11.txt"),0.209},
	{filepath.Join("Tests","test13.txt"), filepath.Join("Tests","test12.txt"),0.300},
	{filepath.Join("Tests","test13.txt"), filepath.Join("Tests","test13.txt"),0.000},
	{filepath.Join("Tests","test13.txt"), filepath.Join("Tests","test14.txt"),0.512},
	{filepath.Join("Tests","test13.txt"), filepath.Join("Tests","test2.txt"),1.000},
	{filepath.Join("Tests","test13.txt"), filepath.Join("Tests","test3.txt"),1.000},
	{filepath.Join("Tests","test13.txt"), filepath.Join("Tests","test4.txt"),1.000},
	{filepath.Join("Tests","test13.txt"), filepath.Join("Tests","test5.txt"),1.000},
	{filepath.Join("Tests","test13.txt"), filepath.Join("Tests","test6.txt"),1.000},
	{filepath.Join("Tests","test13.txt"), filepath.Join("Tests","test7.txt"),1.000},
	{filepath.Join("Tests","test13.txt"), filepath.Join("Tests","test8.txt"),1.000},
	{filepath.Join("Tests","test13.txt"), filepath.Join("Tests","test9.txt"),1.000},
	{filepath.Join("Tests","test14.txt"), filepath.Join("Tests","test1.txt"),1.000},
	{filepath.Join("Tests","test14.txt"), filepath.Join("Tests","test10.txt"),1.000},
	{filepath.Join("Tests","test14.txt"), filepath.Join("Tests","test11.txt"),0.389},
	{filepath.Join("Tests","test14.txt"), filepath.Join("Tests","test12.txt"),0.273},
	{filepath.Join("Tests","test14.txt"), filepath.Join("Tests","test13.txt"),0.512},
	{filepath.Join("Tests","test14.txt"), filepath.Join("Tests","test14.txt"),0.000},
	{filepath.Join("Tests","test14.txt"), filepath.Join("Tests","test2.txt"),1.000},
	{filepath.Join("Tests","test14.txt"), filepath.Join("Tests","test3.txt"),1.000},
	{filepath.Join("Tests","test14.txt"), filepath.Join("Tests","test4.txt"),1.000},
	{filepath.Join("Tests","test14.txt"), filepath.Join("Tests","test5.txt"),1.000},
	{filepath.Join("Tests","test14.txt"), filepath.Join("Tests","test6.txt"),1.000},
	{filepath.Join("Tests","test14.txt"), filepath.Join("Tests","test7.txt"),1.000},
	{filepath.Join("Tests","test14.txt"), filepath.Join("Tests","test8.txt"),1.000},
	{filepath.Join("Tests","test14.txt"), filepath.Join("Tests","test9.txt"),1.000},
	{filepath.Join("Tests","test2.txt"), filepath.Join("Tests","test1.txt"),1.000},
	{filepath.Join("Tests","test2.txt"), filepath.Join("Tests","test10.txt"),1.000},
	{filepath.Join("Tests","test2.txt"), filepath.Join("Tests","test11.txt"),1.000},
	{filepath.Join("Tests","test2.txt"), filepath.Join("Tests","test12.txt"),1.000},
	{filepath.Join("Tests","test2.txt"), filepath.Join("Tests","test13.txt"),1.000},
	{filepath.Join("Tests","test2.txt"), filepath.Join("Tests","test14.txt"),1.000},
	{filepath.Join("Tests","test2.txt"), filepath.Join("Tests","test2.txt"),0.000},
	{filepath.Join("Tests","test2.txt"), filepath.Join("Tests","test3.txt"),1.000},
	{filepath.Join("Tests","test2.txt"), filepath.Join("Tests","test4.txt"),1.000},
	{filepath.Join("Tests","test2.txt"), filepath.Join("Tests","test5.txt"),1.000},
	{filepath.Join("Tests","test2.txt"), filepath.Join("Tests","test6.txt"),1.000},
	{filepath.Join("Tests","test2.txt"), filepath.Join("Tests","test7.txt"),1.000},
	{filepath.Join("Tests","test2.txt"), filepath.Join("Tests","test8.txt"),1.000},
	{filepath.Join("Tests","test2.txt"), filepath.Join("Tests","test9.txt"),1.000},
	{filepath.Join("Tests","test3.txt"), filepath.Join("Tests","test1.txt"),1.000},
	{filepath.Join("Tests","test3.txt"), filepath.Join("Tests","test10.txt"),1.000},
	{filepath.Join("Tests","test3.txt"), filepath.Join("Tests","test11.txt"),1.000},
	{filepath.Join("Tests","test3.txt"), filepath.Join("Tests","test12.txt"),1.000},
	{filepath.Join("Tests","test3.txt"), filepath.Join("Tests","test13.txt"),1.000},
	{filepath.Join("Tests","test3.txt"), filepath.Join("Tests","test14.txt"),1.000},
	{filepath.Join("Tests","test3.txt"), filepath.Join("Tests","test2.txt"),1.000},
	{filepath.Join("Tests","test3.txt"), filepath.Join("Tests","test3.txt"),0.000},
	{filepath.Join("Tests","test3.txt"), filepath.Join("Tests","test4.txt"),0.255},
	{filepath.Join("Tests","test3.txt"), filepath.Join("Tests","test5.txt"),0.212},
	{filepath.Join("Tests","test3.txt"), filepath.Join("Tests","test6.txt"),0.027},
	{filepath.Join("Tests","test3.txt"), filepath.Join("Tests","test7.txt"),1.000},
	{filepath.Join("Tests","test3.txt"), filepath.Join("Tests","test8.txt"),1.000},
	{filepath.Join("Tests","test3.txt"), filepath.Join("Tests","test9.txt"),1.000},
	{filepath.Join("Tests","test4.txt"), filepath.Join("Tests","test1.txt"),1.000},
	{filepath.Join("Tests","test4.txt"), filepath.Join("Tests","test10.txt"),1.000},
	{filepath.Join("Tests","test4.txt"), filepath.Join("Tests","test11.txt"),1.000},
	{filepath.Join("Tests","test4.txt"), filepath.Join("Tests","test12.txt"),1.000},
	{filepath.Join("Tests","test4.txt"), filepath.Join("Tests","test13.txt"),1.000},
	{filepath.Join("Tests","test4.txt"), filepath.Join("Tests","test14.txt"),1.000},
	{filepath.Join("Tests","test4.txt"), filepath.Join("Tests","test2.txt"),1.000},
	{filepath.Join("Tests","test4.txt"), filepath.Join("Tests","test3.txt"),0.255},
	{filepath.Join("Tests","test4.txt"), filepath.Join("Tests","test4.txt"),0.000},
	{filepath.Join("Tests","test4.txt"), filepath.Join("Tests","test5.txt"),0.391},
	{filepath.Join("Tests","test4.txt"), filepath.Join("Tests","test6.txt"),0.280},
	{filepath.Join("Tests","test4.txt"), filepath.Join("Tests","test7.txt"),1.000},
	{filepath.Join("Tests","test4.txt"), filepath.Join("Tests","test8.txt"),1.000},
	{filepath.Join("Tests","test4.txt"), filepath.Join("Tests","test9.txt"),1.000},
	{filepath.Join("Tests","test5.txt"), filepath.Join("Tests","test1.txt"),1.000},
	{filepath.Join("Tests","test5.txt"), filepath.Join("Tests","test10.txt"),1.000},
	{filepath.Join("Tests","test5.txt"), filepath.Join("Tests","test11.txt"),1.000},
	{filepath.Join("Tests","test5.txt"), filepath.Join("Tests","test12.txt"),1.000},
	{filepath.Join("Tests","test5.txt"), filepath.Join("Tests","test13.txt"),1.000},
	{filepath.Join("Tests","test5.txt"), filepath.Join("Tests","test14.txt"),1.000},
	{filepath.Join("Tests","test5.txt"), filepath.Join("Tests","test2.txt"),1.000},
	{filepath.Join("Tests","test5.txt"), filepath.Join("Tests","test3.txt"),0.212},
	{filepath.Join("Tests","test5.txt"), filepath.Join("Tests","test4.txt"),0.391},
	{filepath.Join("Tests","test5.txt"), filepath.Join("Tests","test5.txt"),0.000},
	{filepath.Join("Tests","test5.txt"), filepath.Join("Tests","test6.txt"),0.250},
	{filepath.Join("Tests","test5.txt"), filepath.Join("Tests","test7.txt"),1.000},
	{filepath.Join("Tests","test5.txt"), filepath.Join("Tests","test8.txt"),1.000},
	{filepath.Join("Tests","test5.txt"), filepath.Join("Tests","test9.txt"),1.000},
	{filepath.Join("Tests","test6.txt"), filepath.Join("Tests","test1.txt"),1.000},
	{filepath.Join("Tests","test6.txt"), filepath.Join("Tests","test10.txt"),1.000},
	{filepath.Join("Tests","test6.txt"), filepath.Join("Tests","test11.txt"),1.000},
	{filepath.Join("Tests","test6.txt"), filepath.Join("Tests","test12.txt"),1.000},
	{filepath.Join("Tests","test6.txt"), filepath.Join("Tests","test13.txt"),1.000},
	{filepath.Join("Tests","test6.txt"), filepath.Join("Tests","test14.txt"),1.000},
	{filepath.Join("Tests","test6.txt"), filepath.Join("Tests","test2.txt"),1.000},
	{filepath.Join("Tests","test6.txt"), filepath.Join("Tests","test3.txt"),0.027},
	{filepath.Join("Tests","test6.txt"), filepath.Join("Tests","test4.txt"),0.280},
	{filepath.Join("Tests","test6.txt"), filepath.Join("Tests","test5.txt"),0.250},
	{filepath.Join("Tests","test6.txt"), filepath.Join("Tests","test6.txt"),0.000},
	{filepath.Join("Tests","test6.txt"), filepath.Join("Tests","test7.txt"),1.000},
	{filepath.Join("Tests","test6.txt"), filepath.Join("Tests","test8.txt"),1.000},
	{filepath.Join("Tests","test6.txt"), filepath.Join("Tests","test9.txt"),1.000},
	{filepath.Join("Tests","test7.txt"), filepath.Join("Tests","test1.txt"),1.000},
	{filepath.Join("Tests","test7.txt"), filepath.Join("Tests","test10.txt"),0.333},
	{filepath.Join("Tests","test7.txt"), filepath.Join("Tests","test11.txt"),1.000},
	{filepath.Join("Tests","test7.txt"), filepath.Join("Tests","test12.txt"),1.000},
	{filepath.Join("Tests","test7.txt"), filepath.Join("Tests","test13.txt"),1.000},
	{filepath.Join("Tests","test7.txt"), filepath.Join("Tests","test14.txt"),1.000},
	{filepath.Join("Tests","test7.txt"), filepath.Join("Tests","test2.txt"),1.000},
	{filepath.Join("Tests","test7.txt"), filepath.Join("Tests","test3.txt"),1.000},
	{filepath.Join("Tests","test7.txt"), filepath.Join("Tests","test4.txt"),1.000},
	{filepath.Join("Tests","test7.txt"), filepath.Join("Tests","test5.txt"),1.000},
	{filepath.Join("Tests","test7.txt"), filepath.Join("Tests","test6.txt"),1.000},
	{filepath.Join("Tests","test7.txt"), filepath.Join("Tests","test7.txt"),0.000},
	{filepath.Join("Tests","test7.txt"), filepath.Join("Tests","test8.txt"),0.421},
	{filepath.Join("Tests","test7.txt"), filepath.Join("Tests","test9.txt"),0.350},
	{filepath.Join("Tests","test8.txt"), filepath.Join("Tests","test1.txt"),1.000},
	{filepath.Join("Tests","test8.txt"), filepath.Join("Tests","test10.txt"),0.490},
	{filepath.Join("Tests","test8.txt"), filepath.Join("Tests","test11.txt"),1.000},
	{filepath.Join("Tests","test8.txt"), filepath.Join("Tests","test12.txt"),1.000},
	{filepath.Join("Tests","test8.txt"), filepath.Join("Tests","test13.txt"),1.000},
	{filepath.Join("Tests","test8.txt"), filepath.Join("Tests","test14.txt"),1.000},
	{filepath.Join("Tests","test8.txt"), filepath.Join("Tests","test2.txt"),1.000},
	{filepath.Join("Tests","test8.txt"), filepath.Join("Tests","test3.txt"),1.000},
	{filepath.Join("Tests","test8.txt"), filepath.Join("Tests","test4.txt"),1.000},
	{filepath.Join("Tests","test8.txt"), filepath.Join("Tests","test5.txt"),1.000},
	{filepath.Join("Tests","test8.txt"), filepath.Join("Tests","test6.txt"),1.000},
	{filepath.Join("Tests","test8.txt"), filepath.Join("Tests","test7.txt"),0.421},
	{filepath.Join("Tests","test8.txt"), filepath.Join("Tests","test8.txt"),0.000},
	{filepath.Join("Tests","test8.txt"), filepath.Join("Tests","test9.txt"),0.300},
	{filepath.Join("Tests","test9.txt"), filepath.Join("Tests","test1.txt"),1.000},
	{filepath.Join("Tests","test9.txt"), filepath.Join("Tests","test10.txt"),0.547},
	{filepath.Join("Tests","test9.txt"), filepath.Join("Tests","test11.txt"),1.000},
	{filepath.Join("Tests","test9.txt"), filepath.Join("Tests","test12.txt"),1.000},
	{filepath.Join("Tests","test9.txt"), filepath.Join("Tests","test13.txt"),1.000},
	{filepath.Join("Tests","test9.txt"), filepath.Join("Tests","test14.txt"),1.000},
	{filepath.Join("Tests","test9.txt"), filepath.Join("Tests","test2.txt"),1.000},
	{filepath.Join("Tests","test9.txt"), filepath.Join("Tests","test3.txt"),1.000},
	{filepath.Join("Tests","test9.txt"), filepath.Join("Tests","test4.txt"),1.000},
	{filepath.Join("Tests","test9.txt"), filepath.Join("Tests","test5.txt"),1.000},
	{filepath.Join("Tests","test9.txt"), filepath.Join("Tests","test6.txt"),1.000},
	{filepath.Join("Tests","test9.txt"), filepath.Join("Tests","test7.txt"),0.350},
	{filepath.Join("Tests","test9.txt"), filepath.Join("Tests","test8.txt"),0.300},
	{filepath.Join("Tests","test9.txt"), filepath.Join("Tests","test9.txt"),0.000}}

func TestBrayCurtisDistance(t *testing.T) {
	for _, pair := range tests {
		v := round(BrayCurtisDistance(Read(pair.file1, pair.file2)))
		if v != pair.distance {
			t.Error(
				"For", pair.file1, 
				"and", pair.file2,
				"expected", strconv.FormatFloat(pair.distance, 'f', 3, 64),
				"got", strconv.FormatFloat(v, 'f', 3, 64),
			) 
		}
		/*p1 := strings.Split(pair.file1, "\\")
		p2 := strings.Split(pair.file2, "\\")
		fmt.Printf("{filepath.Join(\"%s\",\"%s\"), filepath.Join(\"%s\",\"%s\"),%.3f},\n", 
		p1[0], p1[1], p2[0], p2[1],round(pair.distance))*/
	}
}

func round(dist float64) float64 {
	return math.Round(dist * float64(1000)) / float64(1000)
}


func makeTests() []testpair {
	names := make([]string, 0) 	
	tests := make([]testpair, 0)
	
	err := filepath.Walk("./Tests", 
		func(path string, info os.FileInfo, err error) error {
			name := info.Name()
			if strings.Contains(name, ".txt") {
				names = append(names, path)
			}
			return nil
		})
	if err != nil {panic(err)}

	for _, file1 := range names {
		for _, file2 := range names {
			bc := BrayCurtisDistance(Read(file1, file2))
			tests = append(tests, testpair{file1, file2, bc})
		}
	}
	return tests
}


func Read(f1 string, f2 string) (map[string]int,map[string]int) {
	map1 := ReadPatternsToMap(f1)
	map2 := ReadPatternsToMap(f2)
	return map1, map2
}

// Assuming well formed input such that each frequency map is formatted as
// string: count
func ReadPatternsToMap(filename string) map[string]int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error reading file.")
	}
	res := make(map[string]int)

	split := strings.Split(strings.Replace(strings.TrimSpace(string(data)), "\r\n", "\n", -1), "\n")

	for _,s := range split {
		splitLine := strings.Split(s, ": ")
		num,_ := strconv.Atoi(splitLine[1])
		if num != 0 {res[splitLine[0]] = num}
	}
	return res
}

