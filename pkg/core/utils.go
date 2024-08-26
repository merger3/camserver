package core

import (
	"net/http"

	"github.com/gempir/go-twitch-irc/v4"
)

type Command struct {
	Channel string
	Command string `query:"command"`
}

func SendCommand(chatter *twitch.Client, cmd Command) int {
	chatter.Say(cmd.Channel, cmd.Command)

	return http.StatusOK
}
