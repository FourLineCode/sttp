package protocol

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type Conn net.Conn
type Addr net.Addr

type Packet struct {
	Body       string
	Host       string
	Port       uint16
	LocalAddr  Addr
	RemoteAddr Addr
	Time       time.Time
}

var (
	ErrInvalidHost    = errors.New("invalid host for protocol message")
	ErrInvalidPort    = errors.New("invalid port for protocol message")
	ErrInvalidMessgae = errors.New("invalid message for protocol message")
	ErrInvalidLocal   = errors.New("invalid local address for protocol message")
	ErrInvalidRemote  = errors.New("invalid remote address for protocol message")
	ErrInvalidPacket  = errors.New("couldn't parse invalid protocol message")
)

func Marshal(p Packet) (string, error) {
	if p.Host == "" {
		return "", ErrInvalidHost
	} else if p.Port == 0 {
		return "", ErrInvalidPort
	} else if p.LocalAddr == nil {
		return "", ErrInvalidLocal
	} else if p.RemoteAddr == nil {
		return "", ErrInvalidRemote
	} else if len(strings.TrimSpace(p.Body)) == 0 {
		return "", ErrInvalidMessgae
	}

	// Format: <host>;<port>;<local>;<remote>;<time>;<body>
	return fmt.Sprintf("%v;%v;%v;%v;%v;%v\n", p.Host, p.Port, p.LocalAddr.String(), p.RemoteAddr.String(), p.Time.Format(time.RFC822), p.Body), nil
}

func Unmarshal(s string) (Packet, error) {
	parts := strings.Split(s, ";")
	if len(parts) < 6 {
		return Packet{}, ErrInvalidPacket
	}

	host, portString, localString, remoteString, timeString := parts[0], parts[1], parts[2], parts[3], parts[4]

	port, err := strconv.Atoi(portString)
	if err != nil {
		return Packet{}, err
	}

	local, err := net.ResolveTCPAddr("tcp", localString)
	if err != nil {
		return Packet{}, err
	}

	remote, err := net.ResolveTCPAddr("tcp", remoteString)
	if err != nil {
		return Packet{}, err
	}

	time, err := time.Parse(time.RFC822, timeString)
	if err != nil {
		return Packet{}, err
	}

	packet := Packet{
		Host:       host,
		Port:       uint16(port),
		LocalAddr:  local,
		RemoteAddr: remote,
		Body:       strings.TrimSpace(strings.Join(parts[5:], ";")),
		Time:       time,
	}

	return packet, nil
}
