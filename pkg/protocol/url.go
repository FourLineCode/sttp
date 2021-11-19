package protocol

import (
	"errors"
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
)

type Url struct {
	Host string
	Port uint16
}

var (
	ErrInvalidUrl     = errors.New("invalid sttp url format")
	ErrInvalidIp      = errors.New("invalid url ip address")
	ErrInvalidUrlPort = errors.New("invalid url port value")
)

func (u Url) String() string {
	return fmt.Sprintf("sttp://%v:%v", u.Host, u.Port)
}

func ParseUrl(url string) (Url, error) {
	ok, err := Validate(url)
	if !ok || err != nil {
		if err != nil {
			return Url{}, err
		}
		return Url{}, ErrInvalidUrl
	}

	urlParams := strings.Split(strings.TrimPrefix(url, "sttp://"), ":")

	host, port := urlParams[0], urlParams[1]
	portNum, _ := strconv.Atoi(port)

	return Url{Host: host, Port: uint16(portNum)}, nil
}

func Validate(url string) (bool, error) {
	url = strings.TrimSpace(url)

	if valid := strings.HasPrefix(url, "sttp://"); !valid {
		return false, ErrInvalidUrl
	}

	urlParams := strings.Split(strings.TrimPrefix(url, "sttp://"), ":")
	if len(urlParams) < 2 {
		return false, ErrInvalidUrl
	}

	host, port := urlParams[0], urlParams[1]

	if net.ParseIP(host) == nil {
		return false, ErrInvalidIp
	}

	portNum, err := strconv.Atoi(port)
	if err != nil || portNum > math.MaxUint16 {
		return false, ErrInvalidUrlPort
	}

	return true, nil
}
