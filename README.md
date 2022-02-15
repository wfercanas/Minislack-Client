# Minislack - Client

## Project Description

This project creates the Client side of a Command Line version of Slack. Users get connected to a server where users and channels are hosted to exchange messages and files.

## Technical Details

- This project uses the net package of Golang to create a connection with the server and exchange commands and information.
- The data flows in bytes in both directions.
- Mainly, the client sends commands and receives notifications from the server to be printed in the CLI. The exceptions are:
  - Sending files: in this case the client has to open the file, read it and send the bytes in the command to the server.
  - Getting files: in this case the client receives the bytes of the requested files, creates a file and saves it in the _./download/_ folder
- **Custom Protocol**: the protocol used to support the communication between client and server was designed for this project and is based in TCP. To check the protocol please refer to the server-side repository [here](https://github.com/wfercanas/Minislack-Server).

## How to use

1. First of all you need to have Go installed in your machine. For info about this point go to the [Download](https://go.dev/dl/) page of the official site.
2. Complete the step by step of [how to run the server](https://github.com/wfercanas/Minislack-Server) first. If the server isn't running when you try the next steps, you won't be able to establish a connection and use the software.
3. Clone this repo in your computer: `git clone`.
4. For each client you must open a new session in your terminal. Use `go run .` to start your client. This automatically tries a connection to **127.0.0.1:3000**, so remember that the server should be up by now in order to create a proper connection. From this moment you are now free to register yourself and use minislack :D. Remember to check the procotol in the server repository.
