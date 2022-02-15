package main

import (
	"bytes"
	"fmt"
	"os"
)

var BREAK_LINE_DELIMITER = []byte("//")

func isFile(message []byte) bool {
	cmd := bytes.ToUpper(bytes.TrimSpace(bytes.Split(message, []byte(" "))[0]))
	return string(cmd) == "FILE"
}

func saveFile(message []byte) {
	cmd := bytes.ToUpper(bytes.TrimSpace(bytes.Split(message, []byte(" "))[0]))
	args := bytes.TrimSpace(bytes.TrimPrefix(message, cmd))
	filename := bytes.TrimSpace(bytes.Split(args, []byte(" "))[0])
	body := bytes.TrimSpace(bytes.TrimPrefix(args, filename))
	splittedBody := bytes.Split(body, []byte("//"))
	formattedBody := bytes.Join(splittedBody, []byte("\n"))

	file, err := os.Create("./downloads/" + string(filename))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	n, err := file.Write(formattedBody)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Filename %s saved in ./downloads folder (%d bytes)\n\n", string(filename), n)
}
