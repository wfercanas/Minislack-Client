package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func read(inbound *bufio.Reader) {
	for {
		message, err := inbound.ReadBytes('\n')
		if err != nil {
			fmt.Print("->> ERR: Connection lost with server\nUse CTRL+C to quit\n")
			return
		}

		if isFile(message) {
			saveFile(message)
		} else {
			fmt.Print(string(message))
		}
	}
}

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println(err)
		return
	}

	outbound := bufio.NewReader(os.Stdin)
	inbound := bufio.NewReader(connection)
	go read(inbound)

	for {
		msg, _ := outbound.ReadBytes('\n')

		text, err := handle(msg)
		if err != nil {
			fmt.Println(err)
		} else {
			connection.Write(text)
		}
	}
}
