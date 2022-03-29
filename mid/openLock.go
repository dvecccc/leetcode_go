package main

import (
	"fmt"
	"strconv"
)

func openLock(deadends []string, target string) int {
	queue := make([][]int, 0)
	queue = append(queue, []int{0, 0, 0, 0})
	m := make(map[string]bool)
	for _, v := range deadends {
		m[v] = true
	}
	step := 0
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			temp := queue[0]
			str := convert2String(temp)
			fmt.Println("temp", str)
			if m[str]{
				queue = queue[1:]
				continue
			}
			if str == target {
				return step
			}
			queue = append(queue, add(temp, 0, 1))
			queue = append(queue, add(temp, 0, -1))
			queue = append(queue, add(temp, 1, 1))
			queue = append(queue, add(temp, 1, -1))
			queue = append(queue, add(temp, 2, 1))
			queue = append(queue, add(temp, 2, -1))
			queue = append(queue, add(temp, 3, 1))
			queue = append(queue, add(temp, 3, -1))
			m[str] = true
			queue = queue[1:]
		}
		println("=====================================================")
		step++
	}
	return -1
}

func add(temp []int, idx, num int) []int {
	var res = make([]int, 4)
	switch num {
	case 1:
		temp[idx] = temp[idx] + 1
		if temp[idx] == 10 {
			temp[idx] = 0
		}
	case -1:
		temp[idx] = temp[idx] - 1
		if temp[idx] < 0 {
			temp[idx] = 9
		}
	}
	copy(res, temp)
	return res
}

func convert2String(nums []int) string {
	var str = ""
	for _, v := range nums {
		str += strconv.Itoa(v)
	}
	return str
}

func main()  {
	nums := []string{"0201","0101","0102","1212","2002"}
	res := openLock(nums, "0202")
	fmt.Println(res)
}
