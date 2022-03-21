# shorturl

## About

This is something like a URL-shortener, that generates shorter links by using a hash function on
a URL. This hash is stored in a MongoDB database.
When you unroll the link, either by using localhost:3000/u/link format, program gets the
entry from the storage using the hash and returns you the link associated with it.

## Usage

1. Change the MongoDB URI for yours in internal/storage/database.go Init().
2. Compile either using Task and Taskfile.yml (task build or task build-release) or using go build main.go.
3. Launch the app using "bin/{mode}/main start-server "key".
4. Use localhost:3000/s/{link} for making short links, localhost:3000/u/{hash} for unrolling them.

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
go run main.go start-server dbkey
```

Basic syntax:

```shell
main start-server dbkey
```
