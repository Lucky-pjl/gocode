package five

import (
	"sort"
)

// 993.二叉树的唐兄弟节点
// https://leetcode-cn.com/problems/cousins-in-binary-tree/
func isCousins(root *TreeNode, x int, y int) bool {
	var xP, yP *TreeNode
	var xDepth, yDepth int
	var xF, yF bool

	var dfs func(node, parent *TreeNode, depth int)
	dfs = func(node, parent *TreeNode, depth int) {
		if node == nil {
			return
		}
		if node.Val == x {
			xP, xDepth, xF = parent, depth, true
		} else if node.Val == y {
			yP, yDepth, yF = parent, depth, true
		}
		if xF && yF {
			return
		}
		dfs(node.Left, node, depth+1)
		if xF && yF {
			return
		}
		dfs(node.Right, node, depth+1)
	}
	dfs(root, nil, 0)
	return xDepth == yDepth && xP != yP
}

// 1442.形成两个异或相等数组的三元组数目
// https://leetcode-cn.com/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/
func countTriplets(arr []int) (ans int) {
	cnt := map[int]int{}
	total := map[int]int{}
	s := 0
	for k, val := range arr {
		if m, has := cnt[s^val]; has {
			ans += m*k - total[s^val]
		}
		cnt[s]++
		total[s] += k
		s ^= val
	}
	return
}

// 1738.找出第K大的异或坐标值 (前缀和)
// https://leetcode-cn.com/problems/find-kth-largest-xor-coordinate-value/
func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	results := make([]int, 0, m*n)
	pre := make([][]int, m+1)
	pre[0] = make([]int, n+1)
	for i, row := range matrix {
		pre[i+1] = make([]int, n+1)
		for j, val := range row {
			pre[i+1][j+1] = pre[i+1][j] ^ pre[i][j+1] ^ pre[i][j] ^ val
			results = append(results, pre[i+1][j+1])
		}
	}
	// 可以改成使用快速排序找第K大的数
	sort.Sort(sort.Reverse(sort.IntSlice(results)))
	return results[k-1]
}

// 692.前K个高频单词
// https://leetcode-cn.com/problems/top-k-frequent-words/
func topKFrequent(words []string, k int) []string {
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w]++
	}
	uniqueWords := make([]string, 0, len(cnt))
	for w := range cnt {
		uniqueWords = append(uniqueWords, w)
	}
	sort.Slice(uniqueWords, func(i, j int) bool {
		s, t := uniqueWords[i], uniqueWords[j]
		return cnt[s] > cnt[t] || cnt[s] == cnt[t] && s < t
	})
	return uniqueWords[:k]
}

// 1035.不相交的线
// https://leetcode-cn.com/problems/uncrossed-lines/
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, v := range nums1 {
		for j, w := range nums2 {
			if v == w {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}
