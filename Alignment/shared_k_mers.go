package Alignment

//CountSharedKmers takes two strings and an integer k. It returns the number of
//k-mers that are shared by the two strings.
func CountSharedKmers(str1, str2 string, k int) int {
	count := 0

	//form frequency map for each of the two strings

	freqMap1 := FrequencyMap(str1, k)
	freqMap2 := FrequencyMap(str2, k)

	for pattern := range freqMap1 {
		// just take the minimum of the two occurrences
		count += Min2(freqMap1[pattern], freqMap2[pattern])
	}

	return count
}

func Min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}
