# oapi_codegen_sample

## Tips

```shell
# Install tools
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
go get github.com/labstack/echo/v4
go get github.com/oapi-codegen/runtime
go get github.com/getkin/kin-openapi/openapi3
go get github.com/oapi-codegen/echo-middleware
go get github.com/swaggo/echo-swagge
go install github.com/swaggo/swag/cmd/swag@latest
go get github.com/swaggo/swag@latest

# Initialize swagger
swag init

# Generate code
rm petstore/petstore-expanded.gen.go
oapi-codegen -generate types,echo-server,spec -package petstore petstore-expanded.yaml > petstore/petstore-expanded.gen.go
```

## Swagger UI
- http://localhost:8080/swagger/index.html

## References
- https://github.com/oapi-codegen/oapi-codegen/blob/main/examples/petstore-expanded/petstore-expanded.yaml
