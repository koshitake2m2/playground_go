# Tips

```bash
# Setup
cd workspace_sample
go work init .
go work use main
go work use main2
go work use ./modules/calc
go work use ./modules/hello

# Run
go run ./main
go run ./main2

# Test
go test ./modules/calc -v
go list -f '{{.Dir}}' -m | xargs go test
```
