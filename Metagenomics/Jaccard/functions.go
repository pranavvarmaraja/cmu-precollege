package Jaccard

// Write the function JaccardDistance() with the following input and output
// Input: two frequency maps sample1 and sample2
// Output: the Jaccard distance between sample1 and sample2
func JaccardDistance(map1, map2 map[string]int) float64 {
  sumMin := SumOfMinima(map1, map2)
  sumMax := SumOfMaxima(map1, map2)
  return 1 - float64(sumMin)/float64(sumMax)
}

//SumOfMaxima takes two frequency maps and returns the max value for every
//key present in either map.
func SumOfMaxima(map1, map2 map[string]int) int {
  c := 0

  //range over maps, take max for every key in either map that we see
  for pattern := range map1 {
    // check if pattern is a key in map2
    _, exists := map2[pattern]
      if exists == true {
        // we have a shared key between the two maps
        c += Max2(map1[pattern], map2[pattern])
      } else { // we don't see this key in map2
        c += map1[pattern]
      }
  }

  //now we need to account for keys of map2 that aren't present in map1
  for pattern := range map2 {
    // check if pattern is a key in map1
    _, exists := map1[pattern]
    if !exists {
      // add the appropriate value of map2
      c += map2[pattern]
    }
  }

  return c
}

//SumOfMinima takes two freq. maps and returns the sum of minimum values for all shared keys.
func SumOfMinima(map1, map2 map[string]int) int {
  c := 0

  //range over maps, finding the shared values, and grabbing the min each time
  for pattern := range map1 {
    // check if pattern is a key in map2
    _, exists := map2[pattern]
      if exists == true {
        // we have a shared value
        // add the minimum of the two values to c
        c += Min2(map1[pattern], map2[pattern])
      }
  }

  return c
}

func Min2(x, y int) int {
  if x < y {
    return x
  }
  return y
}

func Max2(x, y int) int {
  if x > y {
    return x
  }
  return y
}
