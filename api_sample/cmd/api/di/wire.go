//go:build wireinject
// +build wireinject

package di

import (
	baseDomain "api_sample/internal/base/domain"
	baseInfra "api_sample/internal/base/infra"
	basePresentation "api_sample/internal/base/presentation"
	todoDomain "api_sample/internal/todo/domain"
	todoInfra "api_sample/internal/todo/infra"
	todoPresentation "api_sample/internal/todo/presentation"
	todoUsecase "api_sample/internal/todo/usecase"

	"github.com/google/wire"
)

////////////
// base
////////////

func provideDummyAuthenticator() baseInfra.DummyAuthenticator {
	return baseInfra.DummyAuthenticator{}
}

var dummyAuthenticatorSet = wire.NewSet(
	provideDummyAuthenticator,
	wire.Bind(new(baseDomain.Authenticator), new(baseInfra.DummyAuthenticator)),
)

func provideDummyUserRepository() baseInfra.DummyUserRepository {
	return baseInfra.DummyUserRepository{}
}

var dummyUserRepositorySet = wire.NewSet(
	provideDummyUserRepository,
	wire.Bind(new(baseDomain.UserRepository), new(baseInfra.DummyUserRepository)),
)

//	func provideIncrementIdGenerator() *baseInfra.IncrementIdGenerator {
//		return baseInfra.NewIncrementIdGenerator()
//	}
func provideIdGenerator() baseDomain.IdGenerator {
	return baseInfra.NewIncrementIdGeneratorAsIdGenerator()
}

var incrementIdGeneratorSet = wire.NewSet(
	provideIdGenerator,
	// provideIncrementIdGenerator,
	// wire.Bind(new(baseDomain.IdGenerator), new(*infra.IncrementIdGenerator)),
)

// type BaseDependencies struct {
// 	Authenticator        baseDomain.Authenticator
// 	UserRepository       baseDomain.UserRepository
// 	IdGenerator          baseDomain.IdGenerator
// 	AuthenticationHelper basePresentation.AuthenticationHelper
// }

////////////
// todo
////////////

func provideTodoRepository() todoDomain.TodoRepository {
	return todoInfra.DummyTodoRepository{}
}

////////////
// init
////////////

type Dependencies struct {
	TodoController todoPresentation.TodoController
}

func Initialize() (*Dependencies, error) {
	wire.Build(
		dummyAuthenticatorSet,
		dummyUserRepositorySet,
		incrementIdGeneratorSet,
		wire.Struct(new(basePresentation.AuthenticationHelper), "*"),
		provideTodoRepository,
		wire.Struct(new(todoUsecase.TodoUsecase), "*"),
		wire.Struct(new(todoPresentation.TodoController), "*"),
		wire.Struct(new(Dependencies), "*"),
	)
	return nil, nil
}
