package creational

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//close 值为nil的channel会panic
//panic: close of nil channel
func TestCloseNilChannel(t *testing.T) {
	var dataStream chan interface{}
	close(dataStream)
}

// 关闭关闭的通道
// panic: close of closed channel
func TestCloseClosedChannel(t *testing.T) {
	var dataStream chan interface{}
	dataStream = make(chan interface{})
	close(dataStream)
	close(dataStream)
}

// write 关闭的通道
// panic: send on closed channel
func TestWriteClosedChannel(t *testing.T) {
	var dataStream chan interface{}
	dataStream = make(chan interface{})
	close(dataStream)
	dataStream <- struct{}{}
}

// read nil 通道 阻塞
// fatal error: all goroutines are asleep - deadlock!
func TestReadNilChannel(t *testing.T) {
	var dataStream chan interface{}
	dataStream = make(chan interface{}, 0)
	<-dataStream
}

// write nil 通道 阻塞
// fatal error: all goroutines are asleep - deadlock!
func TestWriteNilChannel(t *testing.T) {
	var dataStream chan interface{}
	dataStream <- struct{}{}
}

func TestObjectPool(t *testing.T) {
	pool := NewPool(5)
	for i := 0; i < 6; i++ {
		if v, err := pool.GetObj(1 * time.Second); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(v)
		}
	}
}

func TestObjectPool2(t *testing.T) {
	pool := NewPool(5)
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		index := i
		wg.Add(1)
		go func(pm *PoolManager, index int) {
			fmt.Printf("i = %d\n", index)
			pm.GetObj(1 * time.Second)
			wg.Done()
		}(pool, index)
	}
	wg.Wait()
}
