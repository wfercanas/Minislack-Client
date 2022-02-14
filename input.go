package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func handle(message []byte) ([]byte, error) {
	cmd := bytes.ToUpper(bytes.TrimSpace(bytes.Split(message, []byte(" "))[0]))
	args := bytes.TrimSpace(bytes.TrimPrefix(message, cmd))

	switch string(cmd) {
	case "SEND":
		return send(args)
	}
	return message, nil
}

func send(args []byte) ([]byte, error) {
	if args[0] != '#' {
		return []byte(""), fmt.Errorf("->> ERR: Recipient must be a channel ('#name')")
	}

	recipient := bytes.Split(args, []byte(" "))[0]
	if len(recipient) == 1 {
		return []byte(""), fmt.Errorf("->> ERR: Recipient must have a name ('#name')")
	}

	args = bytes.TrimSpace(bytes.TrimPrefix(args, recipient))
	filename := bytes.Split(args, []byte(" "))[0]
	if len(filename) == 1 {
		return []byte(""), fmt.Errorf("->> ERR: File must be saved with a name")
	}

	filepath := string(bytes.TrimSpace(bytes.TrimPrefix(args, filename)))
	filepathStat, err := os.Stat(filepath)
	if err != nil {
		return []byte(""), err
	}

	if !filepathStat.Mode().IsRegular() {
		return []byte(""), fmt.Errorf("->> ERR: %s is not a regular file", filepath)
	}

	buffer := make([]byte, 8)
	body := []byte(fmt.Sprintf(string(filename) + "\n"))

	file, err := os.Open(filepath)
	if err != nil {
		return []byte(""), err
	}

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return []byte(""), err
		}
		if n == 0 {
			break
		}
		body = append(body, buffer[:n]...)
	}

	command := []byte("SEND ")
	command = append(command, recipient...)
	command = append(command, []byte(" ")...)
	command = append(command, filename...)
	command = append(command, []byte(" ")...)
	command = append(command, body...)

	return command, nil
}
