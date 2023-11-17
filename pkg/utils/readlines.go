package utils

import (
	"bufio"
	"os"
)

func ReadFileLines(filename string) ([]string, error) {
	// 파일을 열어서
	file, err := os.Open(filename)
	// 파일 열기 실패시 에러 리턴
	if err != nil {
		return nil, err
	}
	// 함수 종료시 파일 닫기
	defer file.Close()

	var lines []string
	// 줄 단위로 파일을 읽는 scanner객체 생성
	scanner := bufio.NewScanner(file)
	// 다음 줄에 도착할 때까지 파일 읽기
	for scanner.Scan() {
		// Text()는 읽은 파일 문자열로 반환
		// lines 슬라이스에 읽은 파일 문자열 줄 단위로 저장
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
