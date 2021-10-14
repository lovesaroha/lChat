/*  Love Saroha
    lovesaroha1994@gmail.com (email address)
    https://www.lovesaroha.com (website)
    https://github.com/lovesaroha  (github)
*/
package api

import (
	"errors"
	"net/http"

	"github.com/gorilla/websocket"
)

var users = make(map[string]*websocket.Conn)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(req *http.Request) bool { return true },
}

// User socket handler.
func userSocketHandler(emailAddress string, conn *websocket.Conn) {
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			go removeUserSocket(emailAddress)
			return
		}
	}
}

// Save user socket.
func saveUserSocket(emailAddress string, conn *websocket.Conn) error {
	if _, ok := users[emailAddress]; ok {
		return errors.New("already connected")
	}
	users[emailAddress] = conn
	return nil
}

// Remove user socekt.
func removeUserSocket(emailAddress string) {
	delete(users, emailAddress)
	leaveRooms(userObject{EmailAddress: emailAddress})
}

// Send message to all users.
func sendMessageToAll(message map[string]interface{}) {
	for _, connection := range users {
		connection.WriteJSON(message)
	}
}
