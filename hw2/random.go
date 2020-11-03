package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println(PeriodLength([]int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3}))
}

// =====================================
// Monte Carlo Relative Prime
// Write a function RelativelyPrimeProbability that takes three int as input(x,
// y, numPairs). Uniformly generate pairs of integers between x and y, inclusive
// (numPairs total), and return an estimate of probability of the pairs being
// relatively prime(type float64).

// Your program needs to panic if any of the input is non-positive.

// numPairs <= 100000. For all inputs, either 1 <= x, y <= 1000, or
// x = 1 and y <= 1000000.

// The PRNG will be seeded by the autograder and you are
// free to use it in whatever way you like. We will grade your implementation
// by checking that your output is within a certain amount of the correct answer.
// If your implementation is correct, it has at least 99.9999%
// probability to pass. If you believe your implementation
// is correct but autograder have it wrong, buy a lottery today as it's your
// (super) lucky day!

// Sample Input: (2, 4, 10)
// Sample Output: 0.5
// Note: You can get anything from 0 to 1. The actual probability for these x and y is 4/9 = .444.

// =====================================

func RelativelyPrimeProbability(x int, y int, numPairs int) float64 {
	if x <= 0 || y <= 0 {
		panic("input is negative")
	}
	countRelativePrime := 0
	for i := 0; i < numPairs; i++ {
		a := rand.Intn(y-x+1) + x
		b := rand.Intn(y-x+1) + x
		if EuclidGCD(a, b) == 1 {
			countRelativePrime++
		}
	}
	return float64(countRelativePrime) / float64(numPairs)
}

// =====================================
// Repeat Sequences
// Write a function HasRepeat that takes a slice of integers as input and
// returns a boolean indicating whether there is a repeated value (a value that
// appears twice or more in the slice).

// Length of slice <= 500. Elements in slice <= 100000000 in absolute value.
// Sample Input: [1, 2, 3, 4, 5, 6, 7, 1, 2, 3]
// Sample Output: true
// =====================================

func HasRepeat(arr []int) bool {

	checkRepeat := make(map[int]bool)

	if len(arr) <= 1 {
		return false
	}

	for _, val := range arr {
		if checkRepeat[val] == true {
			return true
		}
		checkRepeat[val] = true
	}
	return false

}

// =====================================
// Monte Carlo Birthday Paradox
// Write a function BirthdayParadox that takes two integers as input(numPeople, numTrials).
// Simulate numPeople random birthdays and estimate the probability that at least two
// of them are the same (sample numTrials times total). Return the estimate
// as a float64.

// Your program should panic if either of the inputs is non-positive.

// numPeople <= 50, numTrials <= 100000. We have the same guarantee regarding
// testing your answer for correctness as in the Monte Carlo relatively prime problem.

// Sample Input: numPeople = 2, numTrials = 1000
// Sample Output: 0.003
// Note: You can get anything from 0 to ~0.01. The actual probability is
// 1/365 = 0.0027...

// =====================================

func BirthdayParadox(numPeople int, numTrials int) float64 {

	if numPeople <= 0 || numTrials <= 0 {
		panic("input is negative")
	}
	countRepeats := 0
	for i := 0; i < numTrials; i++ {
		birthdays := make([]int, 0)
		for j := 0; j < numPeople; j++ {
			birthdays = append(birthdays, rand.Intn(365))
		}
		if HasRepeat(birthdays) {
			countRepeats++
		}
	}
	return float64(countRepeats) / float64(numTrials)

}

// =====================================
// Period Length
// Write a function ComputePeriodLength that takes a slice of int as input and
// check if you can decide the period of the input sequence. If no, return 0;
// If yes, return the period.

// Length of slice <= 500. Elements in slice <= 100000000 in absolute value.
// Sample Input: [1, 2, 3, 4, 5, 6, 7, 1, 2, 3]
// Sample Output: 7
// =====================================

func PeriodLength(arr []int) int {

	checkRepeat := make(map[int]bool)
	indexOfRepeat := 0
	firstIndex := 0

	if len(arr) <= 1 {
		return 0
	}

	for i, val := range arr {
		if checkRepeat[val] == true {
			indexOfRepeat = i
			break

		}
		checkRepeat[val] = true
	}
	for i, val := range arr {
		if val == arr[indexOfRepeat] {
			firstIndex = i
			break
		}
	}

	return indexOfRepeat - firstIndex

}

//=====================================
// Number of Digits
// Write a function CountNumDigits that takes an integer x(type int) as input,
// and return number of digits in x. By convention, 0 have 1 digit, and (-x)
// have same number of digits as x, for any positive integer x.

// You should not rely on any functions from imported packages for this problem.

// Input number <= 10^10.
// Sample Input: 12345
// Sample Output: 5
//=====================================

func NumDigits(x int) int {

	countDigits := 0
	for x != 0 {
		x = x / 10
		countDigits++
	}
	return countDigits

}

// =====================================
// SquareMiddle
// Write a function SquareMiddle that takes two ints: x and numDigits, and
// return the number formed by middle numDigits of x squared.
// Additionally, we ask you to perform sanity checks. If numDigits is not even,
// or if x has more than 2*numDigits of digit, or either input is negative,
// your function should panic.

// For all valid data: 0 <= x < 10^8, 0 < numDigits <= 8
// Your function need to panic if the either of the input falls outside
// of valid range, OR if x is no less than 10^numDigits, OR numDigits is not even.

// Sample Input 1: x = 55, numDigits = 2
// Sample Output 1: 2 (55 * 55 = 3025)
// Sample Input 2: x = 55, numDigits = 3
// Sample Output 2: None(You should panic!)
// =====================================

func SquareMiddle(x int, numDigits int) int {
	if numDigits%2 != 0 || x <= 0 || numDigits <= 0 || NumDigits(x) > 2*numDigits {
		panic("incorrect input")
	}

	x = x * x

	num := strconv.Itoa(x)
	num = num[len(num)/2-numDigits/2 : len(num)/2+numDigits/2]
	ret, _ := strconv.Atoi(num)
	return ret

}

// =====================================
// SquareMiddle Generator
// Now that you have every tool you need, let's implement the actual generator.
// Write a function GenerateMiddleSquareSequence that takes two ints: seed and
// numDigits, and output the generated sequence right after first repeated
// element appears(stop at this point; This is when you know the period of your
// sequence).

// 0 <= seed < 10^numDigits. 2 <= numDigits <= 6. We guarantee the output
// sequence has less than 500 elements.
// Sample Input: seed = 70, numDigits = 2
// Sample Output: [70, 90, 10, 10]
// =====================================

func SquareMiddleGenerator(seed int, numDigits int) []int {

	ret := make([]int, 0)

	ret = append(ret, seed)
	for !HasRepeat(ret) {
		ret = append(ret, SquareMiddle(ret[len(ret)-1], numDigits))
	}
	return ret

}

func WeightedDie() int {

	random := rand.Float64()
	if random < 0.5 {
		return 3
	} else if random < 0.6 {
		return 1
	} else if random < 0.7 {
		return 2
	} else if random < 0.8 {
		return 4
	} else if random < 0.9 {
		return 5
	} else {
		return 6
	}
}

func TrivialGCD(a int, b int) int {
	d := 1
	min := Min2(a, b)

	for i := 1; i <= min; i++ {
		if a%i == 0 && b%i == 0 {
			d = i
		}
	}
	return d
}

func EuclidGCD(a int, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a

}

func Min2(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

func TrivialRangeGCD(a int, b int) float64 {

	start := time.Now()
	for i := 0; i < 10; i++ {

		x := rand.Intn(AbsoluteSubtract(a, b)) + Min2(a, b)
		y := rand.Intn(AbsoluteSubtract(a, b)) + Min2(a, b)
		TrivialGCD(x, y)
	}
	elapsed := time.Since(start)
	return float64(elapsed.Nanoseconds()) / float64(10)

}

func EuclidRangeGCD(a int, b int) float64 {

	start := time.Now()
	for i := 0; i < 10; i++ {

		x := rand.Intn(AbsoluteSubtract(a, b)) + Min2(a, b)
		y := rand.Intn(AbsoluteSubtract(a, b)) + Min2(a, b)
		EuclidGCD(x, y)
	}
	elapsed := time.Since(start)
	return float64(elapsed.Nanoseconds()) / float64(10)

}

func AbsoluteSubtract(a int, b int) int {

	if a > b {
		return a - b
	} else if b > a {
		return b - a
	}
	return 0

}

func TableGCDTimes() map[string][]float64 {
	ret := make(map[string][]float64)

	ret["1000-2000"] = []float64{TrivialRangeGCD(1000, 2000), EuclidRangeGCD(1000, 2000)}
	ret["10000-20000"] = []float64{TrivialRangeGCD(10000, 20000), EuclidRangeGCD(10000, 20000)}
	ret["100000-200000"] = []float64{TrivialRangeGCD(100000, 200000), EuclidRangeGCD(100000, 200000)}
	ret["1000000-2000000"] = []float64{TrivialRangeGCD(1000000, 2000000), EuclidRangeGCD(1000000, 2000000)}

	return ret

}

func GenerateLinearCongruenceSequence(seed int, a int, c int, m int) []int {

	ret := make([]int, 0)
	ret = append(ret, seed)

	for !HasRepeat(ret) {
		seed = (a*seed + c) % m
		ret = append(ret, seed)
	}

	return ret

}
