package websocket

import "golang.org/x/net/websocket"

type messageType struct {
	// ...
}

func WebSocketHandler(conn *websocket.Conn) {
	defer conn.Close()
	for {
		message := messageType{}
		if err := websocket.JSON.Receive(conn, &message); err != nil {
			// Log error, optionally send a close message
			break // Exit the loop to close the connection
		}
		// Process message...
		responseMessage := messageType{}
		if err := websocket.JSON.Send(conn, responseMessage); err != nil {
			// Handle send error
			break
		}
	}
}
