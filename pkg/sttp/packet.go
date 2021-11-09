package sttp

import "time"

type Packet struct {
	Text       string
	LocalAddr  Addr
	RemoteAddr Addr
	Time       time.Time
}
