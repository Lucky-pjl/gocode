package four

import (
	"math"
	"sort"
)

// 1011.在D天内送达包裹的能力 (二分查找转化为判定问题)
// https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/
func shipWithinDays(weights []int, D int) int {
	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}

	return left + sort.Search(right-left, func(x int) bool {
		x += left
		day := 1
		sum := 0
		for _, w := range weights {
			if sum+w > x {
				day++
				sum = 0
			}
			sum += w
		}
		return day <= D
	})
}

// 938.二叉搜索时的范围和
// https://leetcode-cn.com/problems/range-sum-of-bst/
func rangeSumBST(root *TreeNode, low, high int) int {
	if root == nil {
		return 0
	}
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

// 633.平方数之和
// https://leetcode-cn.com/problems/sum-of-square-numbers/
func judgeSquareSum(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum > c {
			right--
		} else {
			left++
		}
	}
	return false
}

// 403.青蛙过河
// https://leetcode-cn.com/problems/frog-jump/
func canCross(stones []int) bool {
	n := len(stones)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true
	for i := 1; i < n; i++ {
		if stones[i]-stones[i-1] > i {
			return false
		}
	}
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			k := stones[i] - stones[j]
			if k > j+1 {
				break
			}
			dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
			if i == n-1 && dp[i][k] {
				return true
			}
		}
	}
	return false
}

// 137.只出现一次的数字 II
// https://leetcode-cn.com/problems/single-number-ii/
func singleNumber(nums []int) int {
	coumap := make(map[int]int)
	for _, v := range nums {
		coumap[v]++
	}
	for num, cou := range coumap {
		if cou == 1 {
			return num
		}
	}
	return 0
}

// 690.员工的重要性
// https://leetcode-cn.com/problems/employee-importance/
type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {

	mp := map[int]*Employee{}
	for _, employee := range employees {
		mp[employee.Id] = employee
	}
	sum := 0
	var dfs func(int)
	dfs = func(id int) {
		employee := mp[id]
		sum += employee.Importance
		for _, v := range employee.Subordinates {
			dfs(v)
		}
	}
	dfs(id)
	return sum
}

// 554.砖墙
// https://leetcode-cn.com/problems/brick-wall/
func leastBricks(wall [][]int) int {
	cnt := map[int]int{}
	for _, widths := range wall {
		sum := 0
		for _, width := range widths[:len(widths)-1] {
			sum += width
			cnt[sum]++
		}
	}
	max := 0
	for _, c := range cnt {
		if c > max {
			max = c
		}
	}
	return len(wall) - max
}

// 1482.制作m束花所需的最少天数
// https://leetcode-cn.com/problems/minimum-number-of-days-to-make-m-bouquets/
func minDays(bloomDay []int, m int, k int) int {
	if k*m > len(bloomDay) {
		return -1
	}
	maxDay := 0
	for _, day := range bloomDay {
		if day > maxDay {
			maxDay = day
		}
	}
	return sort.Search(maxDay, func(days int) bool {
		flowers, bouquets := 0, 0
		for _, d := range bloomDay {
			if d > days {
				flowers = 0
			} else {
				flowers++
				if flowers == k {
					bouquets++
					flowers = 0
				}
			}
		}
		return bouquets >= m
	})
}
