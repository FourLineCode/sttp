package sttp

import (
	"bufio"
	"net"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Conn net.Conn
type Addr net.Addr
type Handler func(packet Packet)

const (
	MaxDeadline = time.Hour
)

type Sttp interface {
	Listen() error
	SetBufferSize(size int)
	SetDeadline(deadline time.Duration)
	OnMessage(handler Handler)
}

type sttp struct {
	port               string
	logger             *logrus.Logger
	maxBufferSize      int
	connectionDeadline time.Duration
	onMessageHandlers  []Handler
}

func NewServer(port int) (Sttp, error) {
	portString, err := transformPort(port)
	if err != nil {
		return nil, err
	}

	s := &sttp{
		port:               portString,
		logger:             logrus.New(),
		maxBufferSize:      1024,
		connectionDeadline: time.Second * 30,
		onMessageHandlers:  make([]Handler, 0),
	}

	return s, nil
}

func (s *sttp) Listen() error {
	listener, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}

	defer listener.Close()
	s.logger.Infof("\nSttp server running on port %v\n", s.port)
	s.logger.Info("Listening for connections ...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			s.logger.Error("Error establishing connection ", err.Error())
			conn.Close()
			continue
		}

		s.logger.Infof("Connection extablished: %v\n", conn.RemoteAddr().String())

		go s.handle(conn)
	}
}

func (s *sttp) handle(conn Conn) {
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(s.connectionDeadline))

	for {
		bytes, _, err := bufio.NewReader(conn).ReadLine()
		if err != nil {
			s.logger.Error("Error while reading from connection ", err.Error())
			return
		}

		conn.SetDeadline(time.Now().Add(s.connectionDeadline))

		packet := Packet{
			Text:       strings.TrimSpace(string(bytes)),
			LocalAddr:  conn.LocalAddr(),
			RemoteAddr: conn.RemoteAddr(),
			Time:       time.Now(),
		}

		for _, handler := range s.onMessageHandlers {
			go handler(packet)
		}
	}
}

func (s *sttp) SetBufferSize(size int) {
	s.maxBufferSize = size
}

func (s *sttp) SetDeadline(deadline time.Duration) {
	s.connectionDeadline = deadline
}

func (s *sttp) OnMessage(handler Handler) {
	s.onMessageHandlers = append(s.onMessageHandlers, handler)
}
