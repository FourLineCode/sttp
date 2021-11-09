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
type MessageHandler func(packet Message)

const (
	MaxDeadline = time.Hour
)

type Sttp interface {
	Listen() error
	SetBufferSize(size int)
	SetDeadline(deadline time.Duration)
	OnMessage(handler MessageHandler)
}

type sttp struct {
	port               string
	logger             *logrus.Logger
	maxBufferSize      int
	connectionDeadline time.Duration
	onMessageHandlers  []MessageHandler
}

func NewServer(port int) (Sttp, error) {
	portString, err := transformPort(port)
	if err != nil {
		return nil, err
	}

	s := &sttp{
		port:               portString,
		logger:             logrus.New(),
		maxBufferSize:      64 * 1024,
		connectionDeadline: time.Second * 30,
		onMessageHandlers:  make([]MessageHandler, 0),
	}

	return s, nil
}

func (s *sttp) Listen() error {
	listener, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}

	defer listener.Close()
	s.logger.Info("Sttp server running on port ", s.port)
	s.logger.Info("Listening for connections ...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			s.logger.Error("Error establishing connection ", err.Error())
			conn.Close()
			continue
		}

		s.logger.Info("Connection extablished: ", conn.RemoteAddr().String())

		go s.handle(conn)
	}
}

func (s *sttp) handle(conn Conn) {
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(s.connectionDeadline))

	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			s.logger.Error("Error while reading from connection ", err.Error())
			return
		}

		conn.SetDeadline(time.Now().Add(s.connectionDeadline))

		packet := Message{
			Body:       strings.TrimSpace(data),
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

func (s *sttp) OnMessage(handler MessageHandler) {
	s.onMessageHandlers = append(s.onMessageHandlers, handler)
}
