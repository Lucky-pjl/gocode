package four

import "sort"

// 28.实现strStr() KMP算法
func strStr(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

// 91.解码方法 (动态规划)
// https://leetcode-cn.com/problems/decode-ways/
func numDecodings(s string) int {
	n := len(s)
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			f[i] += f[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			f[i] += f[i-2]
		}
	}
	return f[n]
}

// 368.最大整除子集 (动态规划)
// https://leetcode-cn.com/problems/largest-divisible-subset/
func largestDivisibleSubset(nums []int) (res []int) {
	sort.Ints(nums)
	// 1.动态规划找出最大子集的个数、最大子集中的最大整数
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	maxSize, maxVal := 1, 1
	for i := 1; i < n; i++ {
		for j, v := range nums[:i] {
			if nums[i]%v == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxSize {
			maxSize, maxVal = dp[i], nums[i]
		}
	}
	if maxSize == 1 {
		return []int{nums[0]}
	}
	// 2.倒推获得最大子集
	for i := n - 1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxVal%nums[i] == 0 {
			res = append(res, nums[i])
			maxVal = nums[i]
			maxSize--
		}
	}
	return
}

// 377.组合总和 IV
// https://leetcode-cn.com/problems/combination-sum-iv/
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

// 897.递增顺序搜索树
// https://leetcode-cn.com/problems/increasing-order-search-tree/
func increasingBST(root *TreeNode) *TreeNode {
	head := &TreeNode{}
	resNode := head
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		resNode.Right = node
		node.Left = nil
		resNode = node
		inorder(node.Right)
	}
	inorder(root)
	return head.Right
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
