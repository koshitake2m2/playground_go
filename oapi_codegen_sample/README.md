# oapi_codegen_sample

## Tips

```shell
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
go get github.com/labstack/echo/v4
go get github.com/oapi-codegen/runtime
oapi-codegen -generate types,echo-server -package petstore petstore-expanded.yaml > petstore-expanded.gen.go
```

## References
- https://github.com/oapi-codegen/oapi-codegen/blob/main/examples/petstore-expanded/petstore-expanded.yaml
