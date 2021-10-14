/*  Love Saroha
    lovesaroha1994@gmail.com (email address)
    https://www.lovesaroha.com (website)
    https://github.com/lovesaroha  (github)
*/
package api

import (
	"encoding/json"
	"errors"
	"time"
)

// Room object.
type roomObject struct {
	ID       string   `json:"_id,omitempty" bson:"_id,omitempty"`
	Password string   `json:"password,omitempty" bson:"password,omitempty"`
	Members  []string `json:"members,omitempty" bson:"members,omitempty"`
}

// Message object.
type messageObject struct {
	To        string    `json:"to,omitempty" bson:"to,omitempty"`
	Content   string    `json:"content,omitempty" bson:"content,omitempty"`
	From      string    `json:"from,omitempty" bson:"from,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

var rooms []roomObject

// This function return json string of room.
func roomJSON(room roomObject) string {
	s, _ := json.Marshal(room)
	return string(s)
}

// This function return json string of rooms.
func roomsJSON(rooms []roomObject) string {
	if len(rooms) < 1 {
		return "[]"
	}
	s, err := json.Marshal(rooms)
	if err != nil {
		return "[]"
	}
	return string(s)
}

// Save and send message.
func saveAndSendMessage(message messageObject, user userObject) {
	if !isRoomMember(message.To, user.EmailAddress) {
		return
	}
	message.From = user.EmailAddress
	message.CreatedAt = time.Now().UTC()
	go sendMessageToRoom(message.To, map[string]interface{}{"type": "message", "message": message})
}

// Save room.
func saveRoom(room *roomObject, user userObject) error {
	room.ID = generateID()
	room.Password = getSHA1Hash(room.Password)
	room.Members = append(room.Members, user.EmailAddress)
	rooms = append(rooms, *room)
	return nil
}

// Join room.
func joinRoom(room roomObject, user userObject) (roomObject, error) {
	password := getSHA1Hash(room.Password)
	for i := 0; i < len(rooms); i++ {
		if rooms[i].ID == room.ID && rooms[i].Password == password {
			// Add user.
			go sendMessageToRoom(rooms[i].ID, map[string]interface{}{"type": "userJoined", "roomID": room.ID, "emailAddress": user.EmailAddress})
			rooms[i].Members = append(rooms[i].Members, user.EmailAddress)
			return rooms[i], nil
		}
	}
	return roomObject{}, errors.New("room not found")
}

// Leave room.
func leaveRooms(user userObject) {
	for i := len(rooms) - 1; i >= 0; i-- {
		for j := len(rooms[i].Members) - 1; j >= 0; j-- {
			if rooms[i].Members[j] == user.EmailAddress {
				var id = rooms[i].ID
				// Remove member.
				rooms[i].Members[j] = rooms[i].Members[len(rooms[i].Members)-1]
				rooms[i].Members = rooms[i].Members[:len(rooms[i].Members)-1]
				if len(rooms[i].Members) == 0 {
					// Remove room.
					removeRoomByID(id)
					go sendMessageToAll(map[string]interface{}{"type": "roomUnavailable", "_id": id})
				} else {
					go sendMessageToRoom(id, map[string]interface{}{"type": "userLeft", "roomID": id, "emailAddress": user.EmailAddress})
				}
			}
		}
	}
}

// Find room by id.
func findRoomByID(id string) *roomObject {
	for i := 0; i < len(rooms); i++ {
		if rooms[i].ID == id {
			return &rooms[i]
		}
	}
	return &roomObject{}
}

// Remove room.
func removeRoomByID(id string) {
	if len(rooms) == 1 {
		rooms = []roomObject{}
		return
	}
	for i := 0; i < len(rooms); i++ {
		if rooms[i].ID == id {
			rooms[i] = rooms[len(rooms)-1]
			rooms = rooms[:len(rooms)-1]
			return
		}
	}
}

// This function checks is room member.
func isRoomMember(roomID string, emailAddress string) bool {
	for i := 0; i < len(rooms); i++ {
		if rooms[i].ID == roomID {
			for j := len(rooms[i].Members) - 1; j >= 0; j-- {
				if rooms[i].Members[j] == emailAddress {
					return true
				}
			}
		}
	}
	return false
}

// Send message to room.
func sendMessageToRoom(id string, message map[string]interface{}) {
	room := findRoomByID(id)
	if len(room.ID) < 4 {
		return
	}
	for i := 0; i < len(room.Members); i++ {
		if connection, ok := users[room.Members[i]]; ok {
			// Send message.
			connection.WriteJSON(message)
		}
	}
}
