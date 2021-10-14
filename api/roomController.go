/*  Love Saroha
    lovesaroha1994@gmail.com (email address)
    https://www.lovesaroha.com (website)
    https://github.com/lovesaroha  (github)
*/
package api

import (
	"encoding/json"
	"net/http"
)

// HandleGetRooms function listen to get /api/get-rooms and return rooms.
func HandleGetRooms(res http.ResponseWriter, req *http.Request) {
	var user = userObject{Token: req.Header.Get("X-Authentication-Token")}
	if invalidUserToken(&user) {
		http.Error(res, `{"error" : "INVALIDUSERTOKEN"}`, http.StatusUnauthorized)
		return
	}
	leaveRooms(user)
	res.Write([]byte(roomsJSON(rooms)))
}

// HandleCreateRoom function listen to post /api/create-room and return new room.
func HandleCreateRoom(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	var user = userObject{Token: req.Header.Get("X-Authentication-Token")}
	if invalidUserToken(&user) {
		http.Error(res, `{"error" : "INVALIDUSERTOKEN"}`, http.StatusUnauthorized)
		return
	}
	var room roomObject
	if json.NewDecoder(req.Body).Decode(&room) != nil {
		http.Error(res, `{"error" : "Invalid room parameters"}`, 400)
		return
	}
	if err := saveRoom(&room, user); err != nil {
		// Unable to create room.
		http.Error(res, `{"error" : "`+err.Error()+`"}`, 400)
		return
	}
	res.Write([]byte(roomJSON(room)))
}

// HandleJoinRoom function listen to post /api/join-room and return room.
func HandleJoinRoom(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	var user = userObject{Token: req.Header.Get("X-Authentication-Token")}
	if invalidUserToken(&user) {
		http.Error(res, `{"error" : "INVALIDUSERTOKEN"}`, http.StatusUnauthorized)
		return
	}
	var room roomObject
	if json.NewDecoder(req.Body).Decode(&room) != nil {
		http.Error(res, `{"error" : "Invalid room parameters"}`, 400)
		return
	}
	room, err := joinRoom(room, user)
	if err != nil {
		// Unable to join room.
		http.Error(res, `{"error" : "`+err.Error()+`"}`, 400)
		return
	}
	res.Write([]byte(roomJSON(room)))
}

// HandleLeaveRooms function listen to get /api/leave-rooms.
func HandleLeaveRooms(res http.ResponseWriter, req *http.Request) {
	var user = userObject{Token: req.Header.Get("X-Authentication-Token")}
	if invalidUserToken(&user) {
		http.Error(res, `{"error" : "INVALIDUSERTOKEN"}`, http.StatusUnauthorized)
		return
	}
	go leaveRooms(user)
	res.Write([]byte(`{"left" : true}`))
}

// HandleCreateMessage function listen to post /api/create-message.
func HandleCreateMessage(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	var user = userObject{Token: req.Header.Get("X-Authentication-Token")}
	if invalidUserToken(&user) {
		http.Error(res, `{"error" : "INVALIDUSERTOKEN"}`, http.StatusUnauthorized)
		return
	}
	var message messageObject
	if json.NewDecoder(req.Body).Decode(&message) != nil {
		http.Error(res, `{"error" : "Invalid message parameters"}`, 400)
		return
	}
	go saveAndSendMessage(message, user)
	res.Write([]byte(`{"created" : true}`))
}
