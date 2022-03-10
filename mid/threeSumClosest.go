package main

import (
	"fmt"
	"math"
	"sort"
)

/*func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	var res int
	var tmp = 11000
	for first := 0; first < n; first++{
		if first > 0 && nums[first] == nums[first - 1] {
			first++
			continue
		}
		for second := first + 1; second < n; second++{
			if second > first + 1 && nums[second] == nums[second - 1] {
				second++
				continue
			}
			for third := n - 1; third >= 0; third--{
				if third == first || third == second || (third < n -1 && nums[third] == nums[third + 1]) {
					continue
				}
				res = nums[first] + nums[second] + nums[third]
				fmt.Println("res: ", res)
				if math.Abs(float64(target - res)) < math.Abs(float64(target - tmp)) {
					tmp = res
					if tmp == target {
						return 0
					}
				}
			}
		}
	}
	return tmp
}*/

/*func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	var tmp = 11000
	for first := 0; first < n - 2; first++{
		for second, third := first + 1, n - 1; second < third; {
			res := nums[first] + nums[second] + nums[third]
			fmt.Println("res: ", res)
			if math.Abs(float64(res - target)) < math.Abs(float64(tmp - target)) {
				tmp = res
				fmt.Println("tmp: ", tmp)
			}
			if res == target {
				return res
			} else if res > target{
				fmt.Println("------")
				third--
				for second < third {
					if nums[third] == nums[third + 1] && third < n - 1{
						third--
						continue
					}
					break
				}
			} else {
				second++
				for second < third {
					if nums[second] == nums[second - 1] && second > first + 1{
						second++
						continue
					}
					break
				}
			}
		}
	}
	return tmp
}*/


func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	var tmp = 11000
	for first := 0; first < n - 2; first++{
		for second, third := first + 1, n - 1; second < third; {
			res := nums[first] + nums[second] + nums[third]
			if math.Abs(float64(res - target)) < math.Abs(float64(tmp - target)) {
				tmp = res
			}
			if res == target {
				return res
			} else if res > target{
				third--
				for second < third {
					if nums[third] == nums[third + 1] && third < n - 1{
						third--
						continue
					}
					break
				}
			} else {
				second++
				for second < third {
					if nums[second] == nums[second - 1] && second > first + 1{
						second++
						continue
					}
					break
				}
			}
		}
	}
	return tmp
}

func main()  {
	res := threeSumClosest([]int{1, 2, 4, 8, 16, 32, 64, 128}, 82)
	fmt.Println(res)
}