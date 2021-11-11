package sttp

import (
	"bufio"
	"io"
	"net"
	"time"

	"github.com/FourLineCode/sttp/pkg/protocol"
	"github.com/sirupsen/logrus"
)

type Conn protocol.Conn
type Addr protocol.Addr
type Message protocol.Packet
type Url protocol.Url
type MessageHandler func(packet Message)

type Sttp interface {
	Listen() error
	SetDeadline(time.Duration)
	OnMessage(MessageHandler)
}

type sttp struct {
	port               uint16
	logger             *logrus.Logger
	connectionDeadline time.Duration
	onMessageHandlers  []MessageHandler
}

func NewServer(port uint16) Sttp {
	s := &sttp{
		port:               port,
		logger:             logrus.New(),
		connectionDeadline: DefaultConnectionDeadline,
		onMessageHandlers:  make([]MessageHandler, 0),
	}

	return s
}

func (s *sttp) Listen() error {
	listener, err := net.Listen("tcp", TransformPort(s.port))
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

		packet, err := protocol.Unmarshal(data)
		if err != nil {
			s.logger.Error("Recieved invalid message ", err.Error())
		}

		for _, handler := range s.onMessageHandlers {
			go handler(Message(packet))
		}
	}
}

func (s *sttp) SetDeadline(deadline time.Duration) {
	s.connectionDeadline = deadline
}

func (s *sttp) OnMessage(handler MessageHandler) {
	s.onMessageHandlers = append(s.onMessageHandlers, handler)
}
