package structural

import "testing"

func TestProxy(t *testing.T) {
	var s Subject = &ProxyObject{RS: &RealSubject{}}
	s.Do()
}
