package five

// 872.叶子相似的树
// https://leetcode-cn.com/problems/leaf-similar-trees/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var l1, l2 []int
	var leafList func(*TreeNode)
	leafList = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			l1 = append(l1, root.Val)
			return
		}
		leafList(root.Left)
		leafList(root.Right)
	}
	leafList(root1)
	l2 = append(l2, l1...)
	l1 = []int{}
	leafList(root2)
	if len(l1) != len(l2) {
		return false
	}
	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

// 1734.解码异或后的排列
// https://leetcode-cn.com/problems/decode-xored-permutation/
func decode(encoded []int) []int {
	n := len(encoded)
	total := 0
	for i := 1; i <= n+1; i++ {
		total ^= i
	}
	odd := 0
	for i := 1; i < n; i += 2 {
		odd ^= encoded[i]
	}
	perm := make([]int, n+1)
	perm[0] = total ^ odd
	for i, v := range encoded {
		perm[i+1] = perm[i] ^ v
	}
	return perm
}

// 1310.子数组异或查询
// https://leetcode-cn.com/problems/xor-queries-of-a-subarray/
func xorQueries(arr []int, queries [][]int) []int {
	// 前缀异或
	xors := make([]int, len(arr)+1)
	for i, v := range arr {
		xors[i+1] = xors[i] ^ v
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		res[i] = xors[q[0]] ^ xors[q[1]+1]
	}
	return res
}

// 1269.停在原地的方案数
// https://leetcode-cn.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/
func numWays(steps, arrLen int) int {
	const mod = 1e9 + 7
	maxColumn := min(arrLen-1, steps)
	dp := make([]int, maxColumn+1)
	dp[0] = 1
	for i := 1; i <= steps; i++ {
		dpNext := make([]int, maxColumn+1)
		for j := 0; j <= maxColumn; j++ {
			dpNext[j] = dp[j]
			if j-1 >= 0 {
				dpNext[j] = (dpNext[j] + dp[j-1]) % (mod)
			}
			if j+1 <= maxColumn {
				dpNext[j] = (dpNext[j] + dp[j+1]) % (mod)
			}
		}
		dp = dpNext
	}
	return dp[0]
}

// 12.整数转罗马数字
// https://leetcode-cn.com/problems/integer-to-roman/
func intToRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strs := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	ans := []byte{}
	for i := 0; num > 0; {
		if num >= nums[i] {
			ans = append(ans, strs[i]...)
			num -= nums[i]
		} else {
			i++
		}
	}
	return string(ans)
}

func romanToInt(s string) (ans int) {
	mmap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	n := len(s)
	for i := range s {
		val := mmap[s[i]]
		if i < n-1 && val < mmap[s[i+1]] {
			ans -= val
		} else {
			ans += val
		}
	}
	return
}

// 421.数组中两个数的最大异或值
// https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array/
func findMaximumXOR(nums []int) (x int) {
	const highBit = 30 // 最高位的二进制位编号为 30
	for k := highBit; k >= 0; k-- {
		// 将所有的 pre^k(a_j) 放入哈希表中
		seen := map[int]bool{}
		for _, num := range nums {
			// 如果只想保留从最高位开始到第 k 个二进制位为止的部分
			// 只需将其右移 k 位
			seen[num>>k] = true
		}

		// 目前 x 包含从最高位开始到第 k+1 个二进制位为止的部分
		// 我们将 x 的第 k 个二进制位置为 1，即为 x = x*2+1
		xNext := x*2 + 1
		found := false

		// 枚举 i
		for _, num := range nums {
			if seen[num>>k^xNext] {
				found = true
				break
			}
		}

		if found {
			x = xNext
		} else {
			// 如果没有找到满足等式的 a_i 和 a_j，那么 x 的第 k 个二进制位只能为 0
			// 即为 x = x*2
			x = xNext - 1
		}
	}
	return
}
