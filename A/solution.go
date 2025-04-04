package main

import (
	"fmt"
)

func countSequences(n int) int {
	if n == 1 {
		return 2
	} else if n == 2 {
		return 4
	} else if n == 3 {
		return 7
	}

	dp := make([]int, n+1)
	dp[1], dp[2], dp[3] = 2, 4, 7

	for i := 4; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	return dp[n]
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(countSequences(n))
}
