// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"api_sample/internal/base/domain"
	"api_sample/internal/base/infra"
	"api_sample/internal/base/presentation"
	domain2 "api_sample/internal/todo/domain"
	infra2 "api_sample/internal/todo/infra"
	presentation2 "api_sample/internal/todo/presentation"
	"api_sample/internal/todo/usecase"
	"github.com/google/wire"
)

// Injectors from wire.go:

func Initialize() (*Dependencies, error) {
	dummyAuthenticator := provideDummyAuthenticator()
	dummyUserRepository := provideDummyUserRepository()
	authenticationHelper := presentation.AuthenticationHelper{
		Authenticator:  dummyAuthenticator,
		UserRepository: dummyUserRepository,
	}
	idGenerator := provideIdGenerator()
	todoRepository := provideTodoRepository()
	todoUsecase := usecase.TodoUsecase{
		IdGenerator:    idGenerator,
		TodoRepository: todoRepository,
	}
	todoController := presentation2.TodoController{
		AuthenticationHelper: authenticationHelper,
		TodoUsecase:          todoUsecase,
	}
	dependencies := &Dependencies{
		TodoController: todoController,
	}
	return dependencies, nil
}

// wire.go:

func provideDummyAuthenticator() infra.DummyAuthenticator {
	return infra.DummyAuthenticator{}
}

var dummyAuthenticatorSet = wire.NewSet(
	provideDummyAuthenticator, wire.Bind(new(domain.Authenticator), new(infra.DummyAuthenticator)),
)

func provideDummyUserRepository() infra.DummyUserRepository {
	return infra.DummyUserRepository{}
}

var dummyUserRepositorySet = wire.NewSet(
	provideDummyUserRepository, wire.Bind(new(domain.UserRepository), new(infra.DummyUserRepository)),
)

//	func provideIncrementIdGenerator() *baseInfra.IncrementIdGenerator {
//		return baseInfra.NewIncrementIdGenerator()
//	}
func provideIdGenerator() domain.IdGenerator {
	return infra.NewIncrementIdGeneratorAsIdGenerator()
}

var incrementIdGeneratorSet = wire.NewSet(
	provideIdGenerator,
)

func provideTodoRepository() domain2.TodoRepository {
	return infra2.DummyTodoRepository{}
}

type Dependencies struct {
	TodoController presentation2.TodoController
}
