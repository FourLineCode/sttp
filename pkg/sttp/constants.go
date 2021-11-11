package sttp

import "time"

const (
	MaxBufferSize             = 64 * 1024
	DefaultPort               = 6969
	DefaultConnectionDeadline = time.Second * 30
)
