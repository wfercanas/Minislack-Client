package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func read(connection net.Conn, inbound *bufio.Reader) {
	for {
		message, err := inbound.ReadString('\n')
		if err != nil {
			fmt.Print("->> ERR: Connection lost with server\n")
			return
		}
		fmt.Print(message)
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
	go read(connection, inbound)

	for {
		text, _ := outbound.ReadString('\n')
		fmt.Fprintf(connection, text+"\n")

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			connection.Close()
			return
		}
	}
}
