package diff

func createLCSTable(a, b []string) [][]int {
	// LCS 테이블을 생성하기 위해 string 슬라이스의 길이 + 1의 크기들로 테이블을 초기화
	table := make([][]int, len(a)+1)
	for i := range table {
		table[i] = make([]int, len(b)+1)
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			// 값이 같으면 동일한 줄의 개수 +1
			if a[i-1] == b[j-1] {
				table[i][j] = table[i-1][j-1] + 1
				// 값이 다르다면 text1만 움직였을 때, text2만 움직였을 때의 동일한 줄의 개수를 비교하여
				// 더 큰 숫자를 기록한다
			} else {
				table[i][j] = max(table[i-1][j], table[i][j-1])
			}
		}
	}

	return table
}
