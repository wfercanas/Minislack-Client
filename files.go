package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
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

	path := "./downloads/" + string(filename)
	ext := bytes.Split(filename, []byte("."))[1]
	name := bytes.TrimRight(bytes.TrimSuffix(filename, ext), ".")

	if !availablePath(path) {
		newName := string(name)
		var newPath string
		for i := 1; ; i++ {
			newName += "(" + strconv.Itoa(i) + ")."
			newPath = "./downloads/" + string(newName) + string(ext)
			if availablePath(newPath) {
				path = newPath
				break
			}
			newName = string(name)
		}
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("->> ERR: saving %s: %e\n", string(filename), err)
	}
	defer file.Close()

	n, err := file.Write(formattedBody)
	if err != nil {
		fmt.Printf("->> ERR: saving %s: %e\n", string(filename), err)
	} else {
		fmt.Printf("->> Filename %s saved in ./downloads folder (%d bytes)\n\n", string(filename), n)
	}
}

func availablePath(path string) bool {
	_, err := os.Stat(path)
	return err != nil
}
