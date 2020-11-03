package main

import ("math/rand"
		"io/ioutil"
		"strings"
		"time")

func main () {
	rand.Seed(time.Now().UnixNano())
	makeAllSame("test1.txt")
	makeVariety("test2.txt")
	makeDiffLengths("test3.txt")
	makeRealistic("test4.txt")
	makeRealistic("test5.txt")
	makeRealistic("test6.txt")
	makeRealistic("test7.txt")
}

func makeAllSame(filename string) {
	l := ""
	s := ""
	alphabet := [4]string{"A", "C", "T", "G"}
	for i := 0; i < 20; i++ {
		l += alphabet[rand.Intn(4)]
	}
	for i := 0; i < 20; i++ {
		s += l + "\n"
	}
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
	
	ioutil.WriteFile(filename, []byte(s), 0644)
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
	
	ioutil.WriteFile(filename, []byte(s), 0644)
}

func makeRealistic(filename string) {
	s := ""
	alphabet := [4]string{"A", "C", "T", "G"}
	frequencies := make([]int, 0)
	totalFreq := 20

	for totalFreq > 1 {
		num := rand.Intn(totalFreq)
		if num < 15 {
			frequencies = append(frequencies, num)
			totalFreq -= num
		}
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

	ioutil.WriteFile(filename, []byte(res), 0644)
}
