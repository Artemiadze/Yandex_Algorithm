package main

import (
	"fmt"
	"sort"
)

func minTotalLength(n int, coords []int) int {
	sort.Ints(coords) // Сортируем координаты

	dp := make([]int, n)
	dp[1] = coords[1] - coords[0] // Минимальное соединение первых двух

	if n > 2 {
		dp[2] = (coords[2] - coords[1]) + dp[1] // Соединяем вторую и третью
	}

	for i := 3; i < n; i++ {
		dp[i] = min(dp[i-1]+(coords[i]-coords[i-1]), dp[i-2]+(coords[i]-coords[i-1]))
	}

	return dp[n-1] // Итоговое минимальное расстояние
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scan(&n)

	coords := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&coords[i])
	}

	fmt.Println(minTotalLength(n, coords))
}
