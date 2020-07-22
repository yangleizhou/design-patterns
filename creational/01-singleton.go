package creational

import "sync"

//单类模式严格一个类只有一个实例，并提供一个全局的访问接口
//1.声明一个全局变量
//2.多线程考虑线程安全，引入sync.Once

//Singleton 是单例模式类
type Singleton struct{}

var (
	once      sync.Once
	singleton *Singleton
)

// GetInstance 获取单例
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
