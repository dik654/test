build:
	go build -o ./bin/diff main.go
run:
	go run main.go test/textfile1.txt test/textfile2.txt
docs:
	@go get golang.org/x/tools/cmd/godoc
	godoc -http=:6060