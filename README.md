# poem
Go-powered CMS for poets and writers (based on gRPC services)

[![Go Reference](https://pkg.go.dev/badge/webimizer.dev/poem.svg)](https://pkg.go.dev/webimizer.dev/poem)

![Poem](cmd/subcommands/adminfrontend/assets/img/poem_logo.png "Poem")

This project is based on Go and use gRPC framework for backend web services. 
By developing this project, I learned:
1. How to create gRPC services using Go and protoc (Protocol Buffer Compiler). [Go to source code](https://github.com/vaclovas2020/poem/blob/main/poems/poems.proto)
2. How to write Go tests. [Go to source code](https://github.com/vaclovas2020/poem/blob/main/poem_test.go)
3. How to create CLI subcommands. [Go to source code](https://github.com/vaclovas2020/poem/blob/main/cmd/subcommands/subcommands.go)
4. How to create generic function to handle multiple types. [Sql ExecDB generic function](https://github.com/vaclovas2020/poem/blob/main/runtime/db_exec.go) [Sql QueryRowDB generic function](https://github.com/vaclovas2020/poem/blob/main/runtime/db_query_row.go) [Sql QueryDB generic function](https://github.com/vaclovas2020/poem/blob/main/runtime/db_query.go) [Template generic function](https://github.com/vaclovas2020/poem/blob/main/runtime/generic_template.go)

Test coverage:

webimizer.dev/poem/runtime	0.004s	coverage: 100.0% of statements

webimizer.dev/poem/cmd/subcommands/install	4.481s	coverage: 55.2% of statements

Build application:
```sh
make
```

Test application:
```sh
make test
```

Install on Docker:
1. Build image from Dockerfile:
```sh
docker build . -t poem
```
2. Copy file ```docker-compose.example.yml``` to ```docker-compose.yml``` and change environment variables (if needed).

```sh
cp docker-compose.example.yml docker-compose.yml
```

3. Build Docker containers with docker-compose:
```sh
 docker-compose up -d
```
4. Install CMS database with:
```sh
docker-compose exec poems_rpc /go/poem install
```