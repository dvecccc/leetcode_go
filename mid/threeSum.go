package main

import (
	"fmt"
	"sort"
)

/*func threeSum(nums []int) [][]int {
	m := make(map[int]int)
	con := make([]int, 0)
	arr := make([][]int, 0)
	deleteRepetedMap := make(map[string]struct{})
	var target int
	for _, value := range nums {
		m[value]++
	}
	for key, _ := range m {
		con = append(con, key)
	}
	sort.Ints(con)
	fmt.Println(con)
	var head int
	for head <= len(con) - 1 && con[head] <= 0{
		for tail := len(con) - 1; head <= tail; tail-- {
			if con[tail] <= 0 {
				break
			}
			m[con[head]]--
			m[con[tail]]--
			target = 0 - (con[tail] + con[head])
			if value, ok := m[target]; ok && value > 0 && m[con[head]] > 0 {
				l := []int{con[head], con[tail], target}
				sort.Ints(l)
				str := strconv.Itoa(l[0]) + strconv.Itoa(l[1]) + strconv.Itoa(l[2])
				_, ok := deleteRepetedMap[str]
				if !ok {
					deleteRepetedMap[str] = struct{}{}
					arr = append(arr, []int{con[head], target, con[tail]})
				}
			}
			m[con[head]]++
			m[con[tail]]++
		}
		head++
	}
	return arr
}*/

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first - 1] {
			continue
		}
		third := n - 1
		target := -1 * nums[first]
		for second := first + 1; second < n; second++ {
			if second > first + 1 && nums[second] == nums[second - 1] {
				continue
			}
			for second < third && nums[second] + nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second] + nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func main()  {
	nums := []int{1, 2, -2, -1}
	arr := threeSum(nums)
	for _, v := range arr {
		fmt.Println(v)
	}
}

/*func binarySearch(nums []int, key int) int {
	low, high := 0, len(nums) - 1
	var mid int
	for low <= high {
		mid = (low + high) / 2
		if nums[mid] == key {
			return mid
		} else if nums[mid] > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}*/
