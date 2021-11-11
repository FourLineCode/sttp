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

type chatClient struct {
	logger *logrus.Logger
	client client.Client
}

func StartClient(r io.Reader) {
	client := chatClient{
		logger: logrus.New(),
		client: client.NewClient(),
	}

	err := client.readLoop(r)
	if err != nil {
		client.logger.Panic("Error while running chat client ", err.Error())
	}
}

func (c chatClient) readLoop(r io.Reader) error {
	for {
		s, err := readString(r)
		if err != nil {
			return err
		}

		command := strings.Split(s, " ")

		port, err := strconv.Atoi(command[0])
		if err != nil || len(command) < 2 {
			// TODO: change this after protocol implementation
			c.logger.Warning("Invalid message (usage: <port> <message>)")
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
