package sttp

import (
	"bufio"
	"io"
	"net"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Conn net.Conn
type Addr net.Addr
type MessageHandler func(packet Message)

type Sttp interface {
	Listen() error
	SetDeadline(time.Duration)
	OnMessage(MessageHandler)
}

type sttp struct {
	port               string
	logger             *logrus.Logger
	connectionDeadline time.Duration
	onMessageHandlers  []MessageHandler
}

func NewServer(port uint16) Sttp {
	s := &sttp{
		port:               TransformPort(port),
		logger:             logrus.New(),
		connectionDeadline: DefaultConnectionDeadline,
		onMessageHandlers:  make([]MessageHandler, 0),
	}

	return s
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

		go s.handle(conn)
	}
}

func (s *sttp) handle(conn Conn) {
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(s.connectionDeadline))

	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				conn.Close()
			} else {
				s.logger.Error("Error while reading from connection ", err.Error())
			}
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

func (s *sttp) SetDeadline(deadline time.Duration) {
	s.connectionDeadline = deadline
}

func (s *sttp) OnMessage(handler MessageHandler) {
	s.onMessageHandlers = append(s.onMessageHandlers, handler)
}
