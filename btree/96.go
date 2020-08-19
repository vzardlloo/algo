/*
 * Author:   Cailei Lu
 * Email:    lucailei@foxmail.com
 * Time:     2020/8/19
 */
package btree

// [96] 不同的二叉搜索树
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}
