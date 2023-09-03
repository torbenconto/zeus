package zeus

func StringDifference(a, b string) string {
	return longestCommonSubsequence(a, b)
}

func longestCommonSubsequence(s1, s2 string) string {
	m := len(s1)
	n := len(s2)

	lcsLengths := make([][]int, m+1)
	for i := range lcsLengths {
		lcsLengths[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				lcsLengths[i][j] = lcsLengths[i-1][j-1] + 1
			} else {
				lcsLengths[i][j] = max(lcsLengths[i-1][j], lcsLengths[i][j-1])
			}
		}
	}

	lcs := make([]byte, 0, lcsLengths[m][n])
	i, j := m, n
	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			lcs = append(lcs, s1[i-1])
			i--
			j--
		} else if lcsLengths[i-1][j] > lcsLengths[i][j-1] {
			i--
		} else {
			j--
		}
	}

	reverseString(lcs)

	return string(lcs)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
