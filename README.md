# vetra
vetra is a tool for static analysis of Go programs.

## Usage
```shell
go install github.com/sepulsa/vetra
go vet -vettool=$(go env GOPATH)/bin/vetra ./...
```

## Registered analyzers
```shell
viper        check for viper.Sub and viper.UnmarshalKey usage
```
