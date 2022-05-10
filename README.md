# poem
Go-powered CMS for poets and writers (based on gRPC services)

[![Go Reference](https://pkg.go.dev/badge/webimizer.dev/poem.svg)](https://pkg.go.dev/webimizer.dev/poem)

![Poem](web/assets/img/poem_logo.png "Poem")

Build application:
```sh
make
```

Test application:
```sh
make test
```

Install on Docker:
1. Build image form Dockerfile:
```sh
docker build . -t poem
```
2. Rename file ```docker-compose.example.yml``` to ```docker-compose.yml``` and change environment variables.
3. Build Docker containers with docker-compose:
```sh
 docker-compose up -d
```
4. Install CMS database with:
```sh
docker-compose exec poems_rpc /go/poem install
```