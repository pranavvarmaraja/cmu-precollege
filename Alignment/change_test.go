package Alignment

import("testing"
	   "strconv")

type changetestpair struct {
	money int
	coins []int
	minNumCoins int
}

var changeTests = []changetestpair{
	{0,[]int{0},0},
	{0,[]int{1,2,7},0},
	{10,[]int{1,2,7},3},
	{21,[]int{1,2,7},3},
	{15,[]int{5,8,2},3},
	{8,[]int{2},4},
	{7,[]int{4},2},
	{10,[]int{3},4},
	{10,[]int{3,1},4}}

func TestChange(t *testing.T) {
	for _, pair := range changeTests {
		v := Change(pair.money, pair.coins)
		if v != pair.minNumCoins {
			t.Error(
				"For", pair.money,
				"and", pair.coins,
				"expected", strconv.Itoa(pair.minNumCoins),
				"got", strconv.Itoa(v),
			)
		}
	}
}
