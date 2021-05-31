package five

import "math/bits"

// 1190.反转每对括号间的子串
// https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses/
func reverseParentheses(s string) string {
	stack := [][]byte{}
	str := []byte{}
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, str)
			str = []byte{}
		} else if s[i] == ')' {
			for j, n := 0, len(str); j < n/2; j++ {
				str[j], str[n-1-j] = str[n-1-j], str[j]
			}
			str = append(stack[len(stack)-1], str...)
			stack = stack[:len(stack)-1]
		} else {
			str = append(str, s[i])
		}
	}
	return string(str)
}

// 461.汉明距离
// https://leetcode-cn.com/problems/hamming-distance/
func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

// 477.汉明距离总和
// https://leetcode-cn.com/problems/total-hamming-distance/
func totalHammingDistance(nums []int) (ans int) {
	n := len(nums)
	for i := 0; i < 30; i++ {
		c := 0
		for _, val := range nums {
			c += val >> i & 1
		}
		ans += c * (n - c)
	}
	return
}

// 231.2的幂
// https://leetcode-cn.com/problems/power-of-two/
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return n&(n-1) == 0
}

func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n%3 == 1
}
