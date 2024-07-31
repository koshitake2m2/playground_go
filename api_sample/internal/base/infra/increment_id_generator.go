package infra

import (
	"api_sample/internal/base/domain"
	"math"
	"sync"
)

type IncrementIdGenerator struct {
	count int
	lock  sync.Mutex
}

func NewIncrementIdGenerator() *IncrementIdGenerator {
	return &IncrementIdGenerator{count: 0, lock: sync.Mutex{}}
}

func NewIncrementIdGeneratorAsIdGenerator() domain.IdGenerator {
	return &IncrementIdGenerator{count: 0, lock: sync.Mutex{}}
}

func (i *IncrementIdGenerator) Generate() int {
	i.lock.Lock()
	i.count++
	if i.count == math.MaxInt {
		i.count = 0
	}
	i.lock.Unlock()

	return i.count
}
