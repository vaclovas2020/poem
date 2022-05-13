build:
	go build -o poem cmd/main/main.go
test:
	go test webimizer.dev/poem
	go test webimizer.dev/poem/runtime
	go test webimizer.dev/poem/cmd/subcommands/install
	go test webimizer.dev/poem/cmd/subcommands/adminserver