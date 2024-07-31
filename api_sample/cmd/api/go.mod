module api_sample/cmd/api

go 1.22.5

require (
	api_sample/internal/base/domain v0.0.0-00010101000000-000000000000
	api_sample/internal/base/infra v0.0.0-00010101000000-000000000000
	api_sample/internal/base/presentation v0.0.0-00010101000000-000000000000
	api_sample/internal/todo/domain v0.0.0-00010101000000-000000000000
	api_sample/internal/todo/infra v0.0.0-00010101000000-000000000000
	api_sample/internal/todo/presentation v0.0.0-00010101000000-000000000000
	api_sample/internal/todo/usecase v0.0.0-00010101000000-000000000000
	github.com/google/wire v0.6.0
	github.com/labstack/echo/v4 v4.12.0
)

require (
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

// base
replace api_sample/internal/base/domain => ../../internal/base/domain

replace api_sample/internal/base/infra => ../../internal/base/infra

replace api_sample/internal/base/presentation => ../../internal/base/presentation

replace api_sample/internal/base/usecase => ../../internal/base/usecase


// todo
replace api_sample/internal/todo/domain => ../../internal/todo/domain

replace api_sample/internal/todo/infra => ../../internal/todo/infra

replace api_sample/internal/todo/presentation => ../../internal/todo/presentation

replace api_sample/internal/todo/usecase => ../../internal/todo/usecase
