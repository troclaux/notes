
# CLI Tools

## CLI tools with its own markdown summaries

- [basic shell commands](./shell.md)
- [git](./git.md)
- [rsync](./backup.md)
- [docker](./docker.md)
- [ansible](./github.md)
- [psql](./postgresql.md)

## echo

> print strings to output

- use " instead of ' for variable expansion or command substitution

```shell

## Print a message with environment variables
echo "My path is $PATH"

## Print a message without the trailing newline
echo -n "Hello World"

```
## cat

## find

## grep

## wc

## xopen

## sed

## jq

## curl

> used to transfer data using various protocols

- downloads files
- interacts with APIs
- tests network performance
- can use many protocols:
  - HTTP
  - HTTPS
  - FTP
  - SCP
  - SMTP
  - POP3
  - etc

### use cases

1. run basic GET request

```shell
curl http://example.com
```

2. test API

```shell
curl http://localhost:5000/api/v1/tasks
```

3. download files

```shell
curl -O http://example.com/file.zip
```

4. upload files

```shell
curl -F "file=@/path/to/file.zip" http://example.com/upload
```

5. get file from github

```shell
curl -sL https://gist.githubusercontent.com/2KAbhishek/9c6d607e160b0439a186d4fbd1bd81df/raw/244284c0b3e40b2b67697665d2d61e537e0890fc/Shell_Keybindings.md
```
