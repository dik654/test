package diff

import (
	"fmt"
	"os"

	"github.com/dik654/test/pkg/utils"
)

// 두 텍스트간의 차이점을 확인시켜주는 함수
func diff(a, b []string) {
	// 두 텍스트 간의 공통 부분 개수를 테이블로 생성
	table := createLCSTable(a, b)
	// 테이블을 확인하며 두 텍스트간 차이점을 콘솔에 뿌리기
	printDiff(a, b, table, len(a), len(b))
}

func Diff() {
	// 비교할 두 파일 경로 입력 받기
	if len(os.Args) < 3 {
		fmt.Println("Usage: diff [file1] [file2]")
		os.Exit(1)
	}

	// 1번 파일 읽어서 []string 슬라이스로 변환
	a, err := utils.ReadFileLines(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file1:", err)
		os.Exit(1)
	}

	b, err := utils.ReadFileLines(os.Args[2])
	if err != nil {
		fmt.Println("Error reading file2:", err)
		os.Exit(1)
	}

	diff(a, b)
}
