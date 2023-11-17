package diff

import "fmt"

func printDiff(a, b []string, table [][]int, i, j int) {
	// text1, 2 차례대로 한 줄 읽었을 때,
	// 두 경우 모두 일치하는 줄의 개수가 동일하다면
	if i > 0 && j > 0 && a[i-1] == b[j-1] {
		// text1, 2 모두 한 줄 앞으로 가서 재귀(읽었을 때 둘의 차이가 없으므로)
		printDiff(a, b, table, i-1, j-1)
	} else {
		// i에는 일치하는 줄이 없고 j는 아직 줄이 남아있다면
		// 또는 text2의 한 줄을 읽었는데(j - 1), text 1의 이전 줄(i)과 동일하다면
		if j > 0 && (i == 0 || table[i][j-1] >= table[i-1][j]) {
			printDiff(a, b, table, i, j-1)
			// text2는 text1에서 한 줄(j)이 추가된 것
			fmt.Printf("%da%d\n< %s\n", i, j, b[j-1])
			// j에는 일치하는 줄이 없고 i에는 줄이 남아있다면
			// text1의 한 줄을 읽었는데(i - 1), text 2의 이전 줄(j)과 동일하다면
		} else if i > 0 && (j == 0 || table[i][j-1] < table[i-1][j]) {
			printDiff(a, b, table, i-1, j)
			// text2는 text1에서 한 줄(i - 1)이 제거된 것
			fmt.Printf("%dd%d\n> %s\n", i, j, a[i-1])
		}
	}
}
