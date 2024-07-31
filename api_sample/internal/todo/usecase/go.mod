module api_sample/internal/todo/usecase

go 1.22.5

// base
replace api_sample/internal/base/domain => ../../base/domain

replace api_sample/internal/base/usecase => ../../base/usecase

// todo
replace api_sample/internal/todo/domain => ../domain

replace api_sample/internal/todo/usecase => ../usecase

require (
	api_sample/internal/base/domain v0.0.0-00010101000000-000000000000
	api_sample/internal/todo/domain v0.0.0-00010101000000-000000000000
)
