package main

import "fmt"

const max = 101

var Graph [max][max]int

func main() {
	var i, j, n, res int
	fmt.Scanf("%d", &n)
	for i = 1; i <= n; i++ {
		for j = 1; j <= i; j++ {
			fmt.Scanf("%d", &res)
			Graph[i][j] = res
		}
	}
	fmt.Println("Graph data wtite end ...")

	maxSum := make([]int, n+1)
	for i = n - 1; i >= 1; i-- {
		for j = 1; j <= i; j++ {
			maxSum[j] = maxInt(maxSum[j], maxSum[j+1]) + Graph[i][j]
		}
	}
	fmt.Println(maxSum[1])
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}
