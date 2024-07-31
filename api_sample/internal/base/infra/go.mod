module api_sample/internal/base/infra

go 1.22.5

replace api_sample/internal/base/domain => ../domain

replace api_sample/internal/base/presentation => ../presentation

replace api_sample/internal/base/usecase => ../usecase

require api_sample/internal/base/domain v0.0.0-00010101000000-000000000000
