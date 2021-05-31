package five

import "math"

// 7.整数反转
// https://leetcode-cn.com/problems/reverse-integer/
func reverse(x int) int {
	res := 0
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		res = res*10 + digit
	}
	return res
}

// 1473.粉刷房子
// https://leetcode-cn.com/problems/paint-house-iii/
func minCost(houses []int, cost [][]int, m, n, target int) int {
	const inf = math.MaxInt64 / 2 // 防止整数相加溢出

	// 将颜色调整为从 0 开始编号，没有被涂色标记为 -1
	for i := range houses {
		houses[i]--
	}

	// dp 所有元素初始化为极大值
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, target)
			for k := range dp[i][j] {
				dp[i][j][k] = inf
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if houses[i] != -1 && houses[i] != j {
				continue
			}

			for k := 0; k < target; k++ {
				for j0 := 0; j0 < n; j0++ {
					if j == j0 {
						if i == 0 {
							if k == 0 {
								dp[i][j][k] = 0
							}
						} else {
							dp[i][j][k] = min(dp[i][j][k], dp[i-1][j][k])
						}
					} else if i > 0 && k > 0 {
						dp[i][j][k] = min(dp[i][j][k], dp[i-1][j0][k-1])
					}
				}

				if dp[i][j][k] != inf && houses[i] == -1 {
					dp[i][j][k] += cost[i][j]
				}
			}
		}
	}

	ans := inf
	for _, res := range dp[m-1] {
		ans = min(ans, res[target-1])
	}
	if ans == inf {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 740.删除并获得点数
// https://leetcode-cn.com/problems/delete-and-earn/
func deleteAndEarn(nums []int) int {
	maxVal := 0
	for _, val := range nums {
		maxVal = max(maxVal, val)
	}
	sum := make([]int, maxVal+1)
	for _, val := range nums {
		sum[val] += val
	}
	return rob(sum)
}

func rob(nums []int) int {
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 1720.解码异或后的数组
// https://leetcode-cn.com/problems/decode-xored-array/
func decode0(encoded []int, first int) []int {
	n := len(encoded) + 1
	res := make([]int, n)
	res[0] = first
	for i, e := range encoded {
		res[i+1] = res[i] ^ e
	}
	return res
}

// 1486.数组异或操作
// https://leetcode-cn.com/problems/xor-operation-in-an-array/
func xorOperation(n int, start int) int {
	res := 0
	for i := 0; i < n; i++ {
		tmp := start + 2*i
		res ^= tmp
	}
	return res
}
