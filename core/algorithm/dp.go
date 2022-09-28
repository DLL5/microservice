package algorithm

// 动态规划算法
func ClimbStairs3(n int) int {
	if n == 1 {
		return 1
	}

	i, j := 1, 2
	for n > 2 {
		i, j = j, i+j // bottom -> up
		n--
	}
	return j
}
