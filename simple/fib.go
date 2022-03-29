package main

import "fmt"

//维持一个长度为n的切片
func fib(n int) int {
	dp := make([]int, n + 1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}
	return dp[n]
}

//只保留当前两个需要的值
func otherFib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := [2]int{0, 1}
	for i := 2; i <=n ; i++ {
		temp := dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = temp
	}
	return dp[1]
}

//递归剪枝
func otherOtherFib(n int) int {
	memo := make([]int, n + 1)
	var f func(int) int
	f = func(n int) int {
		if n == 0 || n == 1 {
			return n
		}
		if memo[n] != 0 {
			return memo[n]
		}
		memo[n] = f(n - 1) + f(n - 2)
		return memo[n]
	}
	f(n)
	return memo[n]
}

func main()  {
	res := otherOtherFib(30)
	fmt.Println(res)
}
