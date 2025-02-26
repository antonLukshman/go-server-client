# Go TCP Server-Client

This is a simple TCP server-client application built using Go (Golang). The server listens for connections, and the client connects to the server, sends a message, and receives a response.

## Features

- TCP server that listens on port `8080`
- Client connects to the server and sends a message
- Server responds with a message
- Handles multiple clients using goroutines

## Prerequisites

- Go installed ([Download Go](https://go.dev/dl/))
- Git installed ([Download Git](https://git-scm.com/downloads))

## Installation & Usage

### 1. Clone the Repository

```sh
git clone https://github.com/your-username/go-server-client.git
cd go-server-client
```

### 2. Run the Server

```sh
go run server.go
```

The server will start listening on `localhost:8080`.

### 3. Run the Client

Open a new terminal and run:

```sh
go run client.go
```

Type a message in the client terminal, and the server will respond.

## Example Output

**Server:**

```
Server is listening on port 8080...
Client connected: 127.0.0.1:54321
Message received: Hello Server!
```

**Client:**

```
Connected to server. Type a message:
Hello Server!
Server response: Hello from server!
```

## Project Structure

```
/go-server-client
│── server.go  # TCP server
│── client.go  # TCP client
│── README.md  # Documentation
│── .gitignore # Git ignore file
```

## License

This project is open-source and available under the MIT License.

## Author

[Your Name](https://github.com/your-username)
