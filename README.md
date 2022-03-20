# shorturl

## About
This is something like a URL-shortener, that generates shorter links by using a hash function on 
a URL. This hash is stored in a local file (not in the database for now) in a JSON format.
When you unroll the link, either by using localhost:3000/u/link format or by CLI, program gets the
entry from the storage using the hash and returns you the link associated with it.

## Usage
1. Compile either using Task and Taskfile.yml (task build or task build-release) or using go build main.go.
2. Go to bin/{mode}/ and run the executable.
3. Follow the CLI for further instructions.

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

Basic syntax:
```shell
main short https://github.com
main unroll somelink
main start-server
```


## Note
Currently, links are stored in a json file in C:\Go\Saved\data.json
The file loads into memory every time you launch the program and links are remembered.
Using the shortened links in the browser requires the server to run for the redirect,
other operations do not require the server to run.
