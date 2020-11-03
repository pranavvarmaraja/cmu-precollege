package Alignment

//Change takes an amount of money along with a collection of denominations as a []int.
//It returns the minimum number of coins needed to change the money given the denominations.
func Change(money int, coins []int) int {
	minNumCoins := make([]int, money+1)

	// range over all relevant values of k
	for k := 1; k <= money; k++ {
		// take minimum of all relevant values at each step
		var currentMin int
		for i := range coins {
			// make sure that current coin isn't too big
			if k-coins[i] >= 0 {
				if minNumCoins[k-coins[i]] == 0 && k-coins[i] != 0 {
					panic("Bad")
				}
				//we're OK
				if i == 0 || minNumCoins[k-coins[i]] < currentMin {
					currentMin = minNumCoins[k-coins[i]]
				}
			}
		}
		// update my minNumCoins[k] value
		minNumCoins[k] = currentMin + 1
	}

	return minNumCoins[money]
}
