package creational

import (
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	go createJob()

	dispatcher := NewDispatcher(10)
	dispatcher.Run()

	time.Sleep(5 * time.Second)
}
