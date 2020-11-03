package Richness

import ("strings"
		"testing"
		"path/filepath"
		"io/ioutil"
		"strconv"
)

type testpair struct {
	filename string
	distance int
}

var tests = []testpair{{filepath.Join("Tests","test1.txt"), 1},
{filepath.Join("Tests","test2.txt"), 20},
{filepath.Join("Tests","test3.txt"), 19},
{filepath.Join("Tests","test4.txt"), 3},
{filepath.Join("Tests","test5.txt"), 4},
{filepath.Join("Tests","test6.txt"), 7},
{filepath.Join("Tests","test7.txt"), 4},
}

func TestRichness(t *testing.T) {
	for _, pair := range tests {
		v := Richness(FreqMapFromFile(pair.filename))
		if v != pair.distance {
			t.Error(
				"For", pair.filename,
				"expected", strconv.Itoa(pair.distance),
				"got", strconv.Itoa(v),
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
