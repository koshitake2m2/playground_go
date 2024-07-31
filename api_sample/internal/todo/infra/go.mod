module api_sample/internal/todo/infra

go 1.22.5

// base
replace api_sample/internal/base/domain => ../../base/domain

replace api_sample/internal/base/infra => ../../base/infra

replace api_sample/internal/base/presentation => ../../base/presentation

replace api_sample/internal/base/usecase => ../../base/usecase

// todo
replace api_sample/internal/todo/domain => ../domain

replace api_sample/internal/todo/presentation => ../presentation

replace api_sample/internal/todo/usecase => ../usecase

require api_sample/internal/todo/domain v0.0.0-00010101000000-000000000000

require api_sample/internal/base/domain v0.0.0-00010101000000-000000000000 // indirect
