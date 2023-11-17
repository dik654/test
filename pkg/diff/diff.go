package diff

// 두 텍스트간의 차이점을 확인시켜주는 함수
func Diff(a, b []string) {
	// 두 텍스트 간의 공통 부분 개수를 테이블로 생성
	table := createLCSTable(a, b)
	// 테이블을 확인하며 두 텍스트간 차이점을 콘솔에 뿌리기
	printDiff(a, b, table, len(a), len(b))
}