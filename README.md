# shorturl

## About

This is something like a URL-shortener, that generates shorter links by using a hash function on
a URL. This hash is stored in a MongoDB database.
When you unroll the link, program gets the entry from the database using the hash and returns
you the link associated with it.

## Usage

1. Change the MongoDB URI for yours in internal/storage/database.go Init().
2. Change the MONGO_KEY in main.go to yours.
3. Compile either using Task and Taskfile.yml (task build or task build-release) or using go build main.go.
4. Launch the app.
5. Use localhost:8080/s/{link} for making short links, localhost:8080/u/{hash} for unrolling them.

Compilation commands:
Task:

```shell
task build-release
task build
task run
```

GoBuild:

```shell
go build main.go
go run main.go
```

Docker:

```shell
docker-compose up --build
```
