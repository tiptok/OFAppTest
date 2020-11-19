package leetcode

import (
	"strconv"
	"testing"
)

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxK := 2
	if len(prices)/2 < maxK {
		maxK = len(prices) / 2
	}
	dptable := map[string]int{}
	for i := 0; i < len(prices); i++ {
		dptable[strconv.Itoa(i)+"-0-0"] = 0
		for k := maxK; k > 0; k-- {
			if i == 0 {
				dptable["0-"+strconv.Itoa(k)+"-0"] = 0
				dptable["0-"+strconv.Itoa(k)+"-1"] = -prices[i]
			} else {
				dptable[strconv.Itoa(i)+"-"+strconv.Itoa(k)+"-0"] = max(dptable[strconv.Itoa(i-1)+"-"+strconv.Itoa(k)+"-0"], dptable[strconv.Itoa(i-1)+"-"+strconv.Itoa(k)+"-1"]+prices[i])
				dptable[strconv.Itoa(i)+"-"+strconv.Itoa(k)+"-1"] = max(dptable[strconv.Itoa(i-1)+"-"+strconv.Itoa(k)+"-1"], dptable[strconv.Itoa(i-1)+"-"+strconv.Itoa(k-1)+"-0"]-prices[i])
			}
		}
	}
	return dptable[strconv.Itoa(len(prices)-1)+"-"+strconv.Itoa(maxK)+"-0"]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Test_maxProfit(t *testing.T) {
	input := [][]int{{518, 498, 505, 505, 506, 501}}
	for i := range input {
		p := maxProfit(input[i])
		t.Log("input:", input[i], " profit:", p)
	}
}
