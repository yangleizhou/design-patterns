package structural

import "testing"

func TestBridge(t *testing.T) {
	http := NewMessageHTTP()
	tcp := NewMessageTCP()
	rm := NewRefindMessage()
	rm.SetMessageImplementer(tcp)
	rm.SendMessage("hello", "heis")
	rm.SetMessageImplementer(http)
	rm.SendMessage("hello", "heis")
}
