package sttp

import "time"

const (
	MaxBufferSize                    = 64 * 1024
	DefaultPort               uint16 = 6969
	DefaultConnectionDeadline        = time.Second * 30
)
