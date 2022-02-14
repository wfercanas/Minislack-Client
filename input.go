package main

import (
	"bytes"
)

func handle(message []byte) []byte {
	cmd := bytes.ToUpper(bytes.TrimSpace(bytes.Split(message, []byte(" "))[0]))
	args := bytes.TrimSpace(bytes.TrimPrefix(message, cmd))

	switch string(cmd) {
	case "SEND":
		if err := send(args); err != nil {
			return []byte("Error sending file")
		}
	}
	return message
}

func send(args []byte) error {
	return nil
}
