package Lib

import (
	"strconv"
)

/************************************************
  CORE LOGIC
************************************************/

/* CreateFrequencyDNAMap
   REQUIRES : true
	 ENSURES  : \result is a string dictionary, keys are [len(dnaStrings)]
	 DESCRIP  : Given a slice of strings, produces an dictionary where dict[i]
	            corresponds to slice[i] in the slice. Essentially provides dummy
							labels for unannotated species.
 */
func CreateFrequencyDNAMap(dnaStrings []string) map[string]string {
	var freqMap = CreateFrequencyMap(dnaStrings)
	var keys, _ = GetKeyValuesI(freqMap)

	var dnaMap = make(map[string]string)
	for i := 0 ; i < len(keys) ; i++ {
		dnaMap[strconv.Itoa(i)] = keys[i]
	}
	return dnaMap
}

/* CreateFrequencyMap
   REQUIRES : true
	 ENSURES  : \result is a valid frequency map
	 DESCRIP  : Given list of strings, creates standard frequency map.
 */
func CreateFrequencyMap(patterns []string) map[string]int {
	freqMap := make(map[string]int)
	for _,val := range patterns {
		freqMap[val]++
	}
	return freqMap
}
