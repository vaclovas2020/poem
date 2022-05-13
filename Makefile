build:
	go build -o poem cmd/main/main.go
test:
	go test webimizer.dev/poem
	go test webimizer.dev/poem/runtime