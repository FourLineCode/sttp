package chat

import "github.com/FourLineCode/sttp/pkg/protocol"

type chatState struct {
	initialized bool
	username    string
	roomUrl     protocol.Url
}

func (s *chatState) SetUsername(u string) {
	s.username = u
}

func (s *chatState) SetUrl(url protocol.Url) {
	s.initialized = true
	s.roomUrl = url
}

func (s *chatState) ClearState() {
	s.initialized = false
	s.roomUrl = protocol.Url{}
	s.username = ""
}
