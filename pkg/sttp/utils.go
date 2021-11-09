package sttp

import (
	"errors"
	"math"
	"strconv"
)

var (
	ErrorInvalidPort = errors.New("invalid port number")
)

func transformPort(port int) (string, error) {
	if port > math.MaxUint16 {
		return "", ErrorInvalidPort
	}

	portString := ":" + strconv.Itoa(port)
	return portString, nil
}
