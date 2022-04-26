# TEST CLI
A test command-line interface created with Go

## Install Go 
- https://go.dev/doc/install

## Install Cobra CLI 
- https://github.com/spf13/cobra

```
 go install github.com/spf13/cobra-cli@latest
``` 

## Create Test CLI

```
 md testcli
 cd testcli
 go mod init testcli
 cobra-cli init
``` 

## Run Test CLI

Execute from source
```
 go run main.go
```

Exceute from binary
```
 go install testcli
 testcli
``` 
