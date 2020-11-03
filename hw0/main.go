package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(NextPerfectNumber(9000))

}

func Factorial(n int) int {

	ret := 1
	for i := 1; i <= n; i++ {
		ret *= i
	}
	return ret
}

func Combination(n, k int) int {

	return ProductUntil(n, k) / Factorial(n-k)
}

func Permutation(n, k int) int {
	return ProductUntil(n, k)
}

func ProductUntil(n, k int) int {
	product := 1
	for ; n > k; n-- {
		product *= n

	}
	return product
}

func FactorialArray(n int) []int {
	factorialArray := make([]int, n+1)
	factorialArray[0] = 1
	for i := 1; i < len(factorialArray); i++ {
		factorialArray[i] = i * factorialArray[i-1]
	}
	return factorialArray
}

func FibonacciArray(n int) []int {
	fibArray := make([]int, n+1)
	fibArray[0] = 1
	fibArray[1] = 1
	for i := 2; i < len(fibArray); i++ {
		fibArray[i] = fibArray[i-1] + fibArray[i-2]
	}
	return fibArray
}

func MinArray(arr []int) int {
	min := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func GCDArray(arr []int) int {
	min := MinArray(arr)
	d := 1

	for i := 2; i < min; i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j]%i != 0 {
				break
			} else if j == len(arr)-1 && arr[j]%i == 0 {
				d = i
			}
		}

	}
	return d
}

func IsPerfect(n int) bool {
	sumDivisors := 1
	for i := 2; float64(i) < math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			sumDivisors += i
			sumDivisors += n / i
		}
	}
	return n == sumDivisors
}

func NextPerfectNumber(n int) int {
	power := 1
	for Exponent(2, power)-1 <= n {
		power++
	}
	for !isPrime(Exponent(2, power) - 1) {
		power++
	}
	return Exponent(2, power-1) * (Exponent(2, power) - 1)
}

func Exponent(base, power int) int {

	product := 1
	for i := 1; i <= power; i++ {
		product *= base
	}
	return product
}

func isPrime(n int) bool {
	for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func NextTwinPrimes(n int) (int, int) {
	n++
	for !isPrime(n) && !isPrime(n+2) {
		n++
	}
	return n, n + 2

}
