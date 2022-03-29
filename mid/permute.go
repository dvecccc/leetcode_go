package main

import "fmt"

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	track := make([]int, 0)

	var f func([]int, []bool)
	f = func(sl[]int, used []bool){
		if len(track) == len(sl) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, track)
			return
		}
		for idx, v := range sl {
			if used[idx]{
				continue
			}
			used[idx] = true
			track = append(track, v)
			f(sl, used)
			track = track[:len(track) - 1]
			used[idx] = false
		}
	}
	f(nums, used)
	return res
}

func main()  {
	nums := []int{1,2,3}
	res := permute(nums)
	fmt.Println(res)
}