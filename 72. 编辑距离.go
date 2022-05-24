package main

/*
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符


示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
*/

func minDistance(word1 string, word2 string) int {
	var n = len(word1)
	var m = len(word2)
	res := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		res[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		res[i][0] = i
	}

	for i := 0; i <= m; i++ {
		res[0][i] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				res[i][j] = res[i-1][j-1]
			} else {
				res[i][j] = 1 + min(min(res[i-1][j], res[i][j-1]), res[i-1][j-1])
			}
		}
	}

	return res[n][m]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
