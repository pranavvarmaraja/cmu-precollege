package main

//StringIndex is a type that will map a minimizer string to its list of indices
//in a read set corresponding to reads with this minimizer.
type StringIndex map[string][]int

//MapToMinimizer takes a collection of reads, integer k and integer numIndices.
//It returns a map of k-mers to the indices of reads in the following fashion.
//It splits the k-mers of a read into numIndices approximately equal pieces.
//For each of these pieces, it adds the index of the read (in our read set)
//to the minimizer's list of indexes.
func MapToMinimizer(reads []string, k int, numIndices int) StringIndex {
	dict := make(StringIndex)
	for i, read := range reads {
		patterns := SplitKmer(read, k, numIndices)
		for _, pattern := range patterns {
			min := Minimizer(pattern, k)
			if !ArrayContains(i, dict[min]) {
				dict[min] = append(dict[min], i)
			}
		}
	}

	return dict
}

func SplitKmer(str string, k, numIndices int) []string {
	ret := make([]string, 0)
	n := len(str) / numIndices
	for i := 0; i < numIndices; i++ {
		if i == 0 {
			ret = append(ret, str[0:n])
		} else if i == numIndices-1 {
			ret = append(ret, str[i*n-k:])

		} else {
			ret = append(ret, str[i*n-k:(i+1)*n])

		}
	}
	return ret
}

func ArrayContains(val int, mtx []int) bool {
	for _, num := range mtx {
		if val == num {
			return true
		}
	}
	return false
}

//note: if n is the length of the current string divided by the number of indices,
//then the pieces we divide into should be:
//[0:n], [n-k:2n], [2n-k:3n], etc.
//the last substring we should consider is the terminal substring [(numIndices*n-k):], the
