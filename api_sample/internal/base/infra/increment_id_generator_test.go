package infra_test

import (
	"api_sample/internal/base/infra"
	"testing"
)

func TestGenerate(t *testing.T) {
	g := infra.NewIncrementIdGenerator()
	if g.Generate() != 1 {
		t.Fatal("generate() should be 1")
	}
	if g.Generate() != 2 {
		t.Fatal("generate() should be 2")
	}
}
