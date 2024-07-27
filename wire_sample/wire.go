//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func provideDog() *Dog {
	dog := Dog{Name: "Doggy"}
	return &dog
}

var animalSet = wire.NewSet(
	provideDog,
	wire.Bind(new(Animal), new(*Dog)),
)

func provideMySingleton() MySingleton {
	return MySingleton{Name: "Bob"}
}

var mySingletonSet = wire.NewSet(
	provideMySingleton,
)

// 以下の代わりに wire.Struct を使う
// func NewAnimalService(animal Animal, mySingleton MySingleton) AnimalService {
// 	return AnimalService{
// 		Animal:      animal,
// 		MySingleton: mySingleton,
// 	}
// }

type DependenciesSet struct {
	AnimalService AnimalService
}

func Initialize() (*DependenciesSet, error) {
	wire.Build(
		animalSet,
		mySingletonSet,
		// NewAnimalService,
		wire.Struct(new(AnimalService), "*"),
		wire.Struct(new(DependenciesSet), "*"),
	)
	return nil, nil
}
