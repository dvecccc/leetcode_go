package main

import (
	"fmt"
)

/*func platesBetweenCandles(s string, queries [][]int) []int {
	var res = make([]int, 0)
	var m = make(map[int]int)
	runes := []rune(s)
	var num int
	for i := 0; i < len(runes); i++ {
		if runes[i] =='*' {
			num++
		} else if runes[i] == '|' {
			m[i] = num
		}
	}
	for _, query := range queries {
		var count int
		for head, tail := query[0], query[1]; head < tail; {
			if runes[head] != '|' {
				head++
			}
			if runes[tail] != '|' {
				tail--
			}

			if runes[head] == '|' && runes[tail] == '|' {
				count = m[tail] - m[head]
				break
			}
		}
		res = append(res, count)
	}
	return res
}*/

func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	var num, l = 0, -1
	var res = make([]int, len(queries))
	var preSum = make([]int, n)
	var left = make([]int, n)
	for index, value := range s {
		if value == '*' {
			num++
		} else {
			l = index
		}
		preSum[index] = num
		left[index] = l
	}
	var right = make([]int, n)
	for index, r := n-1, -1; index >= 0 ; index-- {
		if s[index] == '|' {
			r = index
		}
		right[index] = r
	}
	for index, query := range queries {
		x, y := right[query[0]], left[query[1]]
		if x >= 0 && y >= 0 && x < y {
			res[index] = preSum[y] - preSum[x]
		}
	}
	return res
}

func main()  {
	queries := [][]int{{17,17}, {1,7}, {14,17}, {5,11}, {15,16}}
	res := platesBetweenCandles("***|**|*****|**||**|*", queries)
	fmt.Println(res)
}
