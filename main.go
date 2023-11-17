package main

import (
	"fmt"
	"os"

	"github.com/dik654/test/pkg/diff"
	"github.com/dik654/test/pkg/utils"
)

func main() {
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

	// diff 실행
	diff.Diff(a, b)
}
