package structural

import "testing"

func TestAdapter(t *testing.T) {
	mp := &MusicPlayer{Src: "music.mp4"}
	mp.PlayMusic()
	sp := &SoundPlayer{Src: "heis.m4a"}
	adpter := NewAdapter(sp)
	adpter.PlayMusic()
}
