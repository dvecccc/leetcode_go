package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var result = make([][]int, 0)
	for first := 0; first < n - 3; first++ {
		if first > 0 && nums[first] == nums[first - 1] || nums[first] + nums[n - 3] + nums[n - 2] + nums[n - 1] < target {
			continue
		}
		for second := first + 1; second < n - 2; second++ {
			if second > first + 1 && nums[second] == nums[second - 1] || nums[first] + nums[second] + nums[n - 2] + nums[n - 1] < target{
				continue
			}
			for third, fourth := second + 1, n - 1; third < fourth;  {
				if third > second + 1 && nums[third] == nums[third - 1] || nums[first] + nums[second] + nums[third] + nums[n -1] < target {
					third++
					continue
				}
				if fourth < n - 1 && nums[fourth] == nums[fourth + 1] {
					fourth--
					continue
				}
				res := nums[first] + nums[second] + nums[third] + nums[fourth]
				if res == target {
					result = append(result, []int{nums[first], nums[second], nums[third], nums[fourth]})
					third++
					fourth--
				} else if res > target {
					fourth--
				} else {
					third++
				}
			}
		}
	}
	return result
}

func main()  {
	//nums := []int{1, 0, -1, 0, -2, 2}
	nums2 := []int{1,0,-1,0,-2,2,3,4,5,-6,-3}
	res := fourSum(nums2, 0)
	fmt.Println(res)
}
