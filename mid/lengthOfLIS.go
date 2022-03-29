package main

import (
	"fmt"
	"math"
)

//dp
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	res := math.MinInt32
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//二分
func lengthOfLIS2(nums []int) int {

	return 0
}


func main() {
	nums := []int{10,11,2,3,45,66,7,4,7,8,9,0,1,2,2,3,5}
	res := lengthOfLIS(nums)
	fmt.Println(res)
}