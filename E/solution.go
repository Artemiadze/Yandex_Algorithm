package main

import (
	"fmt"
	"sort"
)

const INF = 1_000_000_000

func main() {
	var N int
	fmt.Scan(&N)

	prices := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&prices[i])
	}

	if N == 0 {
		fmt.Println("0\n0 0")
		return
	}

	// dp[i][j] - минимальная сумма денег на i дней с j купонами в конце i-го дня
	dp := make([][]int, N+1)
	from := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, N+1)
		from[i] = make([]int, N+1)
		for j := range dp[i] {
			dp[i][j] = INF // условно бесконечность
		}
	}
	dp[0][0] = 0

	// Используем динамическое программирование
	for i := 0; i < N; i++ {
		for j := 0; j <= N; j++ {
			if dp[i][j] == INF {
				continue
			}
			// Если покупаем за деньги
			if prices[i] > 100 {
				if dp[i+1][j+1] > dp[i][j]+prices[i] {
					dp[i+1][j+1] = dp[i][j] + prices[i]
					from[i+1][j+1] = j
				}
			} else {
				if dp[i+1][j] > dp[i][j]+prices[i] {
					dp[i+1][j] = dp[i][j] + prices[i]
					from[i+1][j] = j
				}
			}
			// Если используем купон
			if j > 0 {
				if dp[i+1][j-1] > dp[i][j] {
					dp[i+1][j-1] = dp[i][j]
					from[i+1][j-1] = j
				}
			}
		}
	}

	// Определяем минимальную стоимость и максимальное количество оставшихся купонов
	minCost := INF
	remainingCoupons := 0
	for j := N; j >= 0; j-- { // Ищем сначала с максимальным j
		if dp[N][j] < minCost {
			minCost = dp[N][j]
			remainingCoupons = j
		}
	}

	// Восстанавливаем путь
	usedCoupons := []int{}
	j := remainingCoupons
	for i := N; i > 0; i-- {
		prevJ := from[i][j]
		if prevJ > j {
			usedCoupons = append(usedCoupons, i)
		}
		j = prevJ
	}

	// Вывод результата
	fmt.Println(minCost)
	fmt.Println(remainingCoupons, len(usedCoupons))

	sort.Ints(usedCoupons)
	for _, day := range usedCoupons {
		fmt.Println(day)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
