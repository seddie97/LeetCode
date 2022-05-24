package main

import (
	"sync"
	"testing"
)

type SingTon struct {
}

var instance *SingTon
var mu sync.Mutex
var one sync.Once

func TestLazy(t *testing.T) *SingTon {
	if instance == nil {
		instance = &SingTon{}
	}

	return instance
}

func TestSingTon(t *testing.T) *SingTon {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &SingTon{}
	}

	return instance
}

func Once() *SingTon {
	one.Do(func() {
		instance = &SingTon{}
	})
	return instance
}
