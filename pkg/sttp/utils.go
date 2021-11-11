package sttp

import (
	"errors"
	"strconv"
)

var (
	ErrorInvalidPort = errors.New("invalid port number")
)

func TransformPort(port uint16) string {
	portString := ":" + strconv.Itoa(int(port))
	return portString
}
