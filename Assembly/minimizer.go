package main

//Minimizer takes a string text and an integer k as input.
//It returns the k-mer of text that is lexicographically minimum.
func Minimizer(text string, k int) string {

	ret := text[:k]

	for i := 1; i <= len(text)-k; i++ {
		if text[i:i+k] < ret {
			ret = text[i : i+k]
		}
	}
	return ret
}
