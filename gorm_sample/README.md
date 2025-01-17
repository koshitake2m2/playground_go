# gorm_sample

## Depends on

playground_infra/postgres

## Setup

```shell
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u gorm.io/driver/postgres
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Tips migration

```shell
# create migration file
migrate create -ext sql -dir migrations -seq create_articles_table

# migration
migrate --path migrations --database 'postgresql://postgres:postgres@localhost:5432/mydb?sslmode=disable' -verbose up
```

## Tips postgres

```shell
# postgres cli
pgcli -p 5432 -U postgres -d mydb -h localhost

# dump
/opt/homebrew/opt/postgresql@15/bin/pg_dump -p 5432 -U postgres -d mydb -h localhost > dump.sql
```

## References

- https://github.com/go-gorm/postgres
