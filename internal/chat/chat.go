package chat

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/FourLineCode/sttp/pkg/client"
	"github.com/sirupsen/logrus"
)

var (
	DefaultReader = os.Stdin
)

type ChatClient interface {
	Start() error
}

type chatClient struct {
	logger *logrus.Logger
	client client.Client
}

func NewChatClient() ChatClient {
	return chatClient{
		logger: logrus.New(),
		client: client.NewClient(),
	}
}

func (c chatClient) Start() error {
	var err chan error

	go c.readLoop(err)

	return <-err
}

func (c chatClient) readLoop(e chan error) {
	for {
		s, err := readString(DefaultReader)
		if err != nil {
			e <- err
			c.logger.Error("Error reading input ", err.Error())
			return
		}

		command := strings.Split(s, " ")

		port, err := strconv.Atoi(command[0])
		if err != nil || len(command) < 2 {
			continue
		}

		message := strings.Join(command[1:], " ")
		c.client.SendMessage(uint16(port), message)
	}
}

func readString(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	input := scanner.Text()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return input, nil
}
