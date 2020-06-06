package sycli

import (
	"encoding/json"
	"net"
)

type SendWrapper struct {
	Command   Command `json:"command"`
	RequestID int     `json:"request_id,omitempty"`
	Async     bool    `json:"async,omitempty"`
}

type ReceiveWrapper struct {
	Error     string      `json:"error"`
	RequestID int         `json:"request_id"`
	Data      interface{} `json:"data"`
}

type Command struct {
	CommandName string
	Params      []interface{}
}

func (c Command) MarshalJSON() ([]byte, error) {
	all := append([]interface{}{c.CommandName}, c.Params...)
	return json.Marshal(all)
}

// Send a Command to the mpv unix socket
func SendCommand(cmd Command, conn net.Conn) (int, error) {
	wrapper := SendWrapper{Command: cmd}
	b, err := json.Marshal(wrapper)
	if err != nil {
		return 0, err
	}
	return conn.Write([]byte(string(b) + "\n"))
}

// Utility function for quickly creating Command structs for setting
// a boolean property of the player
func SetBoolPropertyCommand(property string, val bool) Command {
	return Command{
		CommandName: "set_property",
		Params:      []interface{}{property, val},
	}
}

// Utility function for quickly creating Command structs for setting
// an integer property of the player
func SetIntPropertyCommand(property string, val int) Command {
	return Command{
		CommandName: "set_property",
		Params:      []interface{}{property, val},
	}
}
