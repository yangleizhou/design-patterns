package structural

import "fmt"

//将一个类的接口转换成客户希望的另外一个接口，使得原本由于接口不兼容而不能一起工作的那些类能一起工作

//golang 实现的是 对象适配器模式

var _ Target = (*adapter)(nil)
var _ Target = (*MusicPlayer)(nil)
var _ Adaptee = (*SoundPlayer)(nil)

// Target 是适配的目标接口
type Target interface {
	PlayMusic()
}

// MusicPlayer 音乐播放器
type MusicPlayer struct {
	Src string
}

// PlayMusic 播放接口
func (mp *MusicPlayer) PlayMusic() {
	fmt.Printf("play music:%v\n", mp.Src)
}

// Adaptee 是被适配的目标接口
type Adaptee interface {
	PlaySound()
}

// SoundPlayer 音频播放器
type SoundPlayer struct {
	Src string
}

// PlaySound 播放接口
func (sp *SoundPlayer) PlaySound() {
	fmt.Printf("play sound:%v\n", sp.Src)
}

// adapter 适配器类,继承adapter&实现Target接口
// 继承Adaptee,golang 没有继承使用匿名组合
type adapter struct {
	Adaptee
}

// PlayMusic 实现Target接口
func (ap *adapter) PlayMusic() {
	ap.PlaySound()
}

//NewAdapter 是Adapter的工厂函数
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}
