package main

import "fmt"

/*type result struct {
	score, index  int
}

func bestRotation(nums []int) int {
	n := len(nums)
	res := result{}

	for k := 0; k < n; k++ {
		var tmp = make([]int, 0)
		tmp = append(tmp, nums[k:]...)
		tmp = append(tmp, nums[:k]...)
		fmt.Println(tmp)
		v := calculateScore(tmp)
		if res.score < v {
			res.score = v
			res.index = k
		}
	}
	return res.index
}

func calculateScore(nums []int) int {
	var score int
	for index, value := range nums {
		if value <= index {
			score++
		}
	}
	return score
}*/

/*type result struct {
	maxScore, index int
}

func bestRotation(nums []int) int {
	n := len(nums)
	var score int
	for i := 0; i < n; i++ {
		if nums[i] <= i {
			score++
		}
	}
	var kMove = make([]int, n)
	for i := 0; i < n; i++ {
		if nums[i] <= i {
			kMove[i - nums[i]]++
		} else {
			kMove[i + n - nums[i]]++
		}
	}
	res := result{maxScore: score}
	for i := 1; i < n; i++ {
		score = score - kMove[i - 1] + 1
		if score > res.maxScore {
			res.index = i
			res.maxScore = score
		}
	}
	return  res.index
}*/

func bestRotation(nums []int) int {
	n := len(nums)
	diffs := make([]int, n)
	for i, num := range nums {
		low := (i + 1) % n
		high := (i - num + n + 1) % n
		diffs[low]++
		diffs[high]--
		if low >= high {
			diffs[0]++
		}
	}
	score, maxScore, idx := 0, 0, 0
	for i, diff := range diffs {
		score += diff
		if score > maxScore {
			maxScore, idx = score, i
		}
	}
	return idx
}


func main()  {
	nums := []int{2, 3, 1, 4, 0}
	res := bestRotation(nums)
	fmt.Println(res)
}
