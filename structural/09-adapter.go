package structural

import "fmt"

//将一个类的接口转换成客户希望的另外一个接口，使得原本由于接口不兼容而不能一起工作的那些类能一起工作

var _ Target = (*adapter)(nil)
var _ Target = (*MusicPlayer)(nil)
var _ Adaptee = (*SoundPlayer)(nil)

// Target 是适配的目标接口
type Target interface {
	PlayMusic()
}

// MusicPlayer 音乐播放器
type MusicPlayer struct {
	src string
}

// PlayMusic 播放接口
func (mp *MusicPlayer) PlayMusic() {
	fmt.Printf("play music:%v\n", mp.src)
}

// Adaptee 是被适配的目标接口
type Adaptee interface {
	PlaySound()
}

// SoundPlayer 音频播放器
type SoundPlayer struct {
	scr string
}

// PlaySound 播放接口
func (sp *SoundPlayer) PlaySound() {
	fmt.Printf("play sound:%v\n", sp.scr)
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
