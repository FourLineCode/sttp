package sttp

import "time"

type Message struct {
	Body       string
	LocalAddr  Addr
	RemoteAddr Addr
	Time       time.Time
}
