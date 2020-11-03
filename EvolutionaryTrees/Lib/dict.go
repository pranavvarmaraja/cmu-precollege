package Lib

/************************************************
  CORE LOGIC
************************************************/

/* GetKeyValues  : map[string]string -> ([]string, []string)
   REQUIRES      : true
	 ENSURES       : (\result1, \result2) are (keys, values) slices of
	                 dnaMap.
	 DESCRIP       : Given a string dictionary, returns slices of keys
	                 and values in the order they were inserted.
 */
func GetKeyValues (dnaMap map[string]string) ([]string, []string) {
	var keys = make([]string, 0, len(dnaMap))
	var values = make([]string, 0, len(dnaMap))
	for k, v := range dnaMap {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

/* GetKeyValuesI  : map[string]string -> ([]string, []string)
   REQUIRES       : true
	 ENSURES        : (\result1, \result2) are (keys, values) slices of
	                 dnaMap.
	 DESCRIP        : Given a string,int dictionary, returns slices of keys
	                 and values in the order they were inserted.
 */
func GetKeyValuesI (dnaMap map[string]int) ([]string, []int) {
	var keys = make([]string, 0, len(dnaMap))
	var values = make([]int, 0, len(dnaMap))
	for k, v := range dnaMap {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}
