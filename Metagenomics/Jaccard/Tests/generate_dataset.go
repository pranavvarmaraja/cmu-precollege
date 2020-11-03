package main

import ("math/rand"
		"io/ioutil"
		"strings"
		"time"
		"strconv")

func main() {
	rand.Seed(time.Now().UnixNano())
	makeAllSame("test1.txt")
	makeVariety("test2.txt")
	s := makeRealistic("test3.txt")
	makeSimilar("test4.txt", s)
	makeSimilar("test5.txt", s)
	makeSimilar("test6.txt", s)
	s2 := makeRealistic("test7.txt")
	makeSimilar("test8.txt", s2)
	makeSimilar("test9.txt", s2)
	makeSimilar("test10.txt", s2)
	c := makeRealistic("test11.txt")
	makeSimilar("test12.txt", c)
	makeSimilar("test13.txt", c)
	makeSimilar("test14.txt", c)
	return
}

func makeAllSame(filename string) {
	l := ""
	s := ""
	alphabet := [4]string{"A", "C", "T", "G"}
	for i := 0; i < 20; i++ {
		l += alphabet[rand.Intn(4)]
	}
	s = l + ": 20"

	ioutil.WriteFile(filename, []byte(s), 0644)
}

func makeVariety(filename string) {
	s := ""
	alphabet := [4]string{"A", "C", "T", "G"}
	for i := 0; i < 20; i++ {
		l := ""
		for j := 0; j < 20; j++ {
			l += alphabet[rand.Intn(4)]
		}
		s += l + "\n"
	}
	
	ioutil.WriteFile(filename, []byte(helper(s)), 0644)
}

func makeDiffLengths(filename string) {
	s := ""
	alphabet := [4]string{"A", "C", "T", "G"}
	for i := 0; i < 20; i++ {
		l := ""
		length := rand.Intn(20)
		for j := 0; j < length; j++ {
			l += alphabet[rand.Intn(4)]
		}
		s += l + "\n"
	}
	
	ioutil.WriteFile(filename, []byte(helper(s)), 0644)
}

func makeRealistic(filename string) string {
	s := ""
	alphabet := [4]string{"A", "C", "T", "G"}
	frequencies := make([]int, 0)
	totalFreq := 20

	for totalFreq > 1 {
		num := rand.Intn(totalFreq)
		frequencies = append(frequencies, num)
		totalFreq -= num
	}

	for _, f := range frequencies {
		l := ""
		for j := 0; j < 20; j++ {
			l += alphabet[rand.Intn(4)]
		}
		for i := 0; i < f; i++ {
			s += l + "\n"
		}
	}

	s2 := strings.Fields(s)
	rand.Shuffle(len(s2)-1, func(p, q int) {
		s2[p], s2[q] = s2[q], s2[p]
	})

	res := ""
	for _, val := range s2 {
		res += val + "\n"
	}

	ioutil.WriteFile(filename, []byte(helper(res)), 0644)
	return res
}

func makeSimilar(filename string, res string) {
	freqMap := makeFreqMap(res)
	r := ""

	for key,val := range freqMap {
		var newVal int
		if rand.Intn(2) == 0 {
			newVal = val + rand.Intn(10)
		} else {
			newVal = val - rand.Intn(10)
			if newVal < 0 {newVal = 0}
		}
		r += key + ": " + strconv.Itoa(newVal) + "\n"
	}

	ioutil.WriteFile(filename, []byte(r), 0644)
} 

func helper(s string) string {
	freqMap := makeFreqMap(s)

	res := ""
	for key,val := range freqMap {
		res += key + ": " + strconv.Itoa(val) + "\n"
	}
	return res
}

func makeFreqMap(s string) map[string]int{
	patterns := strings.Split(strings.TrimSpace(strings.Replace(s, "\r\n", "\n", -1)), "\n")
	freqMap := make(map[string]int)
	for _,val := range patterns {
		freqMap[val]++
	}
	return freqMap
}