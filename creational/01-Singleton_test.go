package creational

import (
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	s1 := GetInstance()
	s2 := GetInstance()
	if s1 != s2 {
		t.Fatalf("instance is not equal s1 addr = %p ,s2 addr = %p", s1, s2)
	}
	t.Logf("s1 addr = %p ,s2 addr = %p", s1, s2)
}

func TestParallelSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(50)
	instances := [50]*Singleton{}
	for i := 0; i < 50; i++ {
		go func(index int) {
			defer wg.Done()
			instances[index] = GetInstance()
		}(i)
	}
	wg.Wait()
	for i := 1; i < 50; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
