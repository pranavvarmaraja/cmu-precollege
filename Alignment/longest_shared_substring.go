package Alignment

//LongestSharedSubstring takes two strings as input.  It returns the shared substring
//of the input strings of maximum length, via a brute-force algorithm.
//If multiple answers exist, it returns only one.
func LongestSharedSubstring(s1, s2 string) string {
	if len(s1) == 0 || len(s2) == 0 {
		return "" // or panic
	}

	longestString := ""

	substrings1 := AllSubstrings(s1)
	substrings2 := AllSubstrings(s2)

	// now let's range over all substrings and update longestString if we see something shared

	for i := range substrings1 {
		if len(substrings1[i]) > len(longestString) && Contains(substrings2, substrings1[i]) {
			longestString = substrings1[i]
		}
	}

	return longestString
}

func Contains(patterns []string, s string) bool {
	for _, pattern := range patterns {
		if pattern == s {
			// found match !
			return true
		}
	}
	// we survived all the checks
	return false
}

//AllSubstrings takes a string as input and returns a slice of all substrings.
func AllSubstrings(s string) []string {
	if len(s) == 0 {
		panic("Error: zero-length substring given to AllSubstrings.")
	}
	patterns := make([]string, 0)

	// range over all possible substrings. Have we seen it before? If not, add it to patterns
	for i := range s {
		// i is starting point of my substring
		for j := i + 1; j <= len(s); j++ {
			currString := s[i:j]
			if Contains(patterns, currString) == false {
				// now I know it is new and I should add to my patterns
				patterns = append(patterns, currString)
			}
		}
	}

	return patterns
}
