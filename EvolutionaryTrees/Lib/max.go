package Lib

func MaxFloat(nums ...float64) float64 {
  m := 0.0
  // nums gets converted to an array
  for i, val := range nums {
    if val > m || i == 0 {
      // update m
      m = val
    }
  }
  return m
}
