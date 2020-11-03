package FrequencyMap

import ("strings"
		"testing"
		"path/filepath"
		"io/ioutil"
		"reflect"
		"strconv")

type testpair struct {
	filename string
	freqMap map[string]int
}

var f1 = filepath.Join("Tests","test1.txt")
var f2 = filepath.Join("Tests","test2.txt")
var f3 = filepath.Join("Tests","test3.txt")
var f4 = filepath.Join("Tests","test4.txt")
var f5 = filepath.Join("Tests","test5.txt")
var f6 = filepath.Join("Tests","test6.txt")
var f7 = filepath.Join("Tests","test7.txt")

var f1sol = filepath.Join("Tests","test1sol.txt")
var f2sol = filepath.Join("Tests","test2sol.txt")
var f3sol = filepath.Join("Tests","test3sol.txt")
var f4sol = filepath.Join("Tests","test4sol.txt")
var f5sol = filepath.Join("Tests","test5sol.txt")
var f6sol = filepath.Join("Tests","test6sol.txt")
var f7sol = filepath.Join("Tests","test7sol.txt")

var tests = []testpair{{f1, FreqMapFromFile(f1sol)}, 
{f2, FreqMapFromFile(f2sol)},
{f3, FreqMapFromFile(f3sol)},
{f4, FreqMapFromFile(f4sol)},
{f5, FreqMapFromFile(f5sol)},
{f6, FreqMapFromFile(f6sol)},
{f7, FreqMapFromFile(f7sol)}}

func TestFrequencyMap(t *testing.T) {
	for _, pair := range tests {
		v := FrequencyMap(Read(pair.filename))
		if !reflect.DeepEqual(v, pair.freqMap) {
			t.Error(
				"For", pair.filename,
				"expected", pair.freqMap,
				"got", v,
			) 
		}
	}
}

func Read(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error reading file.")
	}
	return strings.Split(strings.TrimSpace(strings.Replace(string(data), "\r\n", "\n", -1)), "\n")
}

func FreqMapFromFile(filename string) map[string]int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(strings.Replace(string(data), "\r\n", "\n", -1)), "\n")
	freqMap := make(map[string]int)
	for _,line := range lines {
		l := strings.Split(line, ":")
		key,val := l[0], l[1]
		freqMap[key],_ = strconv.Atoi(val)
	}
	return freqMap
}