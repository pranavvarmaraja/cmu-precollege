package BrayCurtis

// Write the function BrayCurtisDistance() with the following input and output
// Input: two frequency maps sample1 and sample2
// Output: the Bray-Curtis distance between sample1 and sample2
func BrayCurtisDistance(map1, map2 map[string]int) float64 {
  sumMin := SumOfMinima(map1, map2)
  // now we need to take the average total

  s1 := SampleTotal(map1)
  s2 := SampleTotal(map2)

  //don't allow zero richness in either sample
  if s1 <= 0 || s2 <= 0 {
    panic("Error: sample total isn't positive in BrayCurtis distance...")
  }

  av := Average(float64(s1), float64(s2))

  return 1.0 - float64(sumMin)/av
}

//Average takes two decimal variables and returns their average.
func Average(x, y float64) float64 {
  return (x+y)/2.0
}

//SampleTotal takes a frequency map and return the sum of all its values.
//Assumes that the values are all nonnegative.
func SampleTotal(freq map[string]int) int {
  s := 0

  //range over values of freq, adding each one to s
  for _, val := range freq {
    // perhaps we panic if we hit a neg value or zero value
    if val > 0 {
      s += val
    }
  }

  return s
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
