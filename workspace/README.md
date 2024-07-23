# Tips

```bash
cd workspace
go work init .
go work use main
go work use ../modules/hello
go run ./main

# Test
go test ../modules/calc -v
go list -f '{{.Dir}}' -m | xargs go test
```
