package Evenness

import ("strings"
		"testing"
		"path/filepath"
		"io/ioutil"
		"strconv"
		"math")

type testpair struct {
	filename string
	distance float64
}

var tests = []testpair{{filepath.Join("Tests","test1.txt"), 1.000},
{filepath.Join("Tests","test2.txt"), 0.050},
{filepath.Join("Tests","test3.txt"), 0.055},
{filepath.Join("Tests","test4.txt"), 0.579},
{filepath.Join("Tests","test5.txt"), 0.446},
{filepath.Join("Tests","test6.txt"), 0.191},
{filepath.Join("Tests","test7.txt"), 0.507}}

func TestEvenness(t *testing.T) {
	for _, pair := range tests {
		v := round(SimpsonsIndex(FreqMapFromFile(pair.filename)))
		if v != pair.distance {
			t.Error(
				"For", pair.filename,
				"expected", strconv.FormatFloat(pair.distance, 'f', 3, 64),
				"got", strconv.FormatFloat(v, 'f', 3, 64),
			) 
		}
	}
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
		freqMap[key],_ = strconv.Atoi(strings.TrimSpace(val))
	}
	return freqMap
}

func round(dist float64) float64 {
	return math.Round(dist * float64(1000)) / float64(1000)
}