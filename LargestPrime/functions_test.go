package LargestPrime

import (
	"fmt"
	"testing"
)

func TestLargestPrimeFactor(t *testing.T) {
	var tests = []struct {
		n    int
		want int
	}{
		{2, 2},
		{23, 23},
		{24, 3},
		{34, 17},
		{25, 5},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.n)
		t.Run(testname, func(t *testing.T) {
			ans := LargestPrimeFactor(tt.n)
			if ans != tt.want {
				t.Errorf("Error! For n = %d, your code gives %d, and the largest prime factor of n is %d", tt.n, ans, tt.want)
			} else {
				fmt.Println("Correct! When n is", tt.n, "then the largest prime factor of n is", tt.want)
			}
		})
	}
}
