package chat

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/FourLineCode/sttp/pkg/client"
	"github.com/FourLineCode/sttp/pkg/logger"
	"github.com/FourLineCode/sttp/pkg/protocol"
)

var (
	DefaultReader = os.Stdin
)

type chatClient struct {
	client client.Client
	host   protocol.Url
	state  *chatState
}

func StartClient(r io.Reader, url protocol.Url) {
	client := chatClient{
		client: client.NewClient(),
		host:   url,
		state:  &chatState{},
	}

	err := client.readLoop(r)
	if err != nil {
		logger.Panic("Error while running chat client %v", err.Error())
	}
}

func (c *chatClient) readLoop(r io.Reader) error {
	for {
		s, err := readString(r)
		if err != nil {
			return err
		}

		if strings.HasPrefix(s, "/") {
			cmd, err := parseCommand(s)
			if err != nil {
				logger.Error("Error parsing command %v", err.Error())
				continue
			}

			if err := execCommand(cmd, c); err != nil {
				logger.Error("Error executing command \"%v\"", cmd.Type)
				logger.Error(err.Error())
				continue
			}
		} else {
			if strings.TrimSpace(s) == "" {
				continue
			}

			if !c.state.initialized {
				logger.Warn("Cannot send message without joining room")
				logger.Info("Use command \"/join <url>\" to join a room")
				logger.Info("Or use command \"/msg <url> <body>\" to directly send a message")
				continue
			}

			// TODO: add username here later
			packet := protocol.Packet{
				Body: s,
				Host: c.host.Host,
				Port: c.host.Port,
			}

			if err := c.client.SendMessage(c.state.roomUrl, packet); err != nil {
				logger.Error("Couldn't send message %v", err.Error())
				continue
			}
		}
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
