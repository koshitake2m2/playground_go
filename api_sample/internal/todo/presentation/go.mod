module api_sample/internal/todo/presentation

go 1.22.5

require (
	api_sample/internal/base/presentation v0.0.0-00010101000000-000000000000
	api_sample/internal/todo/domain v0.0.0-00010101000000-000000000000
	api_sample/internal/todo/usecase v0.0.0-00010101000000-000000000000
	github.com/labstack/echo/v4 v4.12.0
)

require (
	api_sample/internal/base/domain v0.0.0-00010101000000-000000000000 // indirect
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
replace api_sample/internal/base/domain => ../../base/domain

replace api_sample/internal/base/presentation => ../../base/presentation

replace api_sample/internal/base/usecase => ../../base/usecase

// todo
replace api_sample/internal/todo/domain => ../domain

replace api_sample/internal/todo/presentation => ../presentation

replace api_sample/internal/todo/usecase => ../usecase
