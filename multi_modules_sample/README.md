# multi modules sample

## Tips

Set up bbb module. bbb module depends on aaa module. 

```bash
cd bbb
go mod init example.com/bbb
go mod edit -replace example.com/aaa=../aaa
go mod tidy
```

## Run

```bash
cd main
go run .
```
