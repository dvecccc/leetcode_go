package main

import (
	"fmt"
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)
	dp[0] = 0
	sort.Ints(coins)
	for n := 1; n <= amount; n++{
		res := 0
		resMax := math.MaxInt64
		for _, v := range coins {
			if n - v >= 0 && dp[n - v] != -1{
				res = dp[n - v] + 1
				resMax = min(resMax, res)
			} else if n - v >= 0 {
				continue
			} else {
				break
			}
		}
		if resMax == math.MaxInt64 {
			dp[n] = -1
		} else {
			dp[n] = resMax
		}
	}
	return dp[amount]
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main()  {
	coins := []int{2}
	res := coinChange(coins, 11)
	fmt.Println(res)
}
