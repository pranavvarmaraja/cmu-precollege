package LargestPrime

import "math"

//Write a function to find a largest prime number.

//LargestPrime takes an integer n as input.
//It panics if n < 2. Otherwise, it returns the largest prime
//factor of n.
func LargestPrimeFactor(n int) int {
	largest := 0
	if n < 2 {
		panic(string(n) + " does not have any prime factors")
	}
	if isPrime(n) {
		return n
	}
	for i := 2; i <= n/2; i++ {
		if isPrime(i) && n%i == 0 {
			largest = i
		}
	}
	return largest
}

func isPrime(n int) bool {

	if n == 2 {
		return true
	}

	for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {

		if n%i == 0 {
			return false
		}
	}
	return true
}
