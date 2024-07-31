module api_sample/internal/todo/domain

go 1.22.5

// base
replace api_sample/internal/base/domain => ../../base/domain

// todo
replace api_sample/internal/todo/domain => ../domain

require api_sample/internal/base/domain v0.0.0-00010101000000-000000000000
