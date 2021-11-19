package chat

import (
	"errors"
	"net"
	"strings"
	"time"

	"github.com/FourLineCode/sttp/pkg/logger"
	"github.com/FourLineCode/sttp/pkg/protocol"
	"github.com/FourLineCode/sttp/pkg/sttp"
)

type commandType string

const (
	CommandHelp    commandType = "/help"
	CommandJoin    commandType = "/join"
	CommandLeave   commandType = "/leave"
	CommandMessage commandType = "/msg"
)

type command struct {
	Type commandType
	Args []string
}

var (
	ErrNotValidCommand      = errors.New("provided string is not a valid command")
	ErrCommandNotRecognized = errors.New("command not recognized")
	ErrInvalidJoinArgs      = errors.New("invalid arguement for command \"join\" (usage: \"/join <url>\")")
	ErrInvalidMsgArgs       = errors.New("invalid arguement for command \"msg\" (usage: \"/msg <url> <body>\")")
	ErrConnectionRefused    = errors.New("cannot establish connection to url")
)

func parseCommand(s string) (command, error) {
	if !strings.HasPrefix(s, "/") {
		return command{}, ErrNotValidCommand
	}

	var cmdType commandType
	args := strings.Split(strings.TrimSpace(s), " ")
	if len(args) < 1 {
		return command{}, ErrNotValidCommand
	}

	switch args[0] {
	case string(CommandHelp):
		cmdType = CommandHelp
	case string(CommandJoin):
		cmdType = CommandJoin
	case string(CommandLeave):
		cmdType = CommandLeave
	case string(CommandMessage):
		cmdType = CommandMessage
	default:
		return command{}, ErrCommandNotRecognized
	}

	return command{Type: cmdType, Args: args[1:]}, nil
}

func execCommand(c command, client *chatClient) error {
	switch c.Type {
	case CommandHelp:
		return helpCommand()
	case CommandJoin:
		return joinCommand(c, client)
	case CommandLeave:
		return leaveCommand(c, client)
	case CommandMessage:
		return messageCommand(c, client)
	default:
		return ErrCommandNotRecognized
	}
}

func helpCommand() error {
	logger.Info("Here is a list of all the commands -")

	logger.Info("\t/help\t- lists all the available commands\n\n")

	logger.Info("\t/msg\t- directly send message to a client")
	logger.Info("\t\tusage: /msg <url> <body>\n\n")

	logger.Info("\t/join\t- join a users room")
	logger.Info("\t\tusage: /join <url>\n\n")

	logger.Info("\t/leave\t- leave a users room")
	logger.Info("\t\tusage: /leave")

	return nil
}

func joinCommand(c command, chat *chatClient) error {
	if len(c.Args) < 1 {
		return ErrInvalidJoinArgs
	}

	url, err := protocol.ParseUrl(c.Args[0])
	if err != nil {
		return err
	}

	conn, err := net.DialTimeout("tcp", url.Host+sttp.TransformPort(url.Port), time.Second)
	if err != nil {
		return ErrConnectionRefused
	}
	conn.Close()

	if chat.state.initialized {
		leaveCommand(c, chat)
	}

	chat.state.SetUrl(url)
	logger.Success("Successfully joined room - %v", chat.state.roomUrl.String())

	return nil
}

func leaveCommand(c command, chat *chatClient) error {
	if !chat.state.initialized {
		logger.Warn("You are not joined in a room")
		return nil
	}

	logger.Success("Successfully left room - %v", chat.state.roomUrl.String())
	chat.state.ClearState()

	return nil
}

func messageCommand(c command, chat *chatClient) error {
	if len(c.Args) < 2 {
		return ErrInvalidMsgArgs
	}

	packet := protocol.Packet{
		Body: strings.Join(c.Args[1:], " "),
		Host: chat.host.Host,
		Port: chat.host.Port,
	}

	url, err := protocol.ParseUrl(c.Args[0])
	if err != nil {
		return err
	}

	if err := chat.client.SendMessage(url, packet); err != nil {
		return err
	}

	return nil
}
