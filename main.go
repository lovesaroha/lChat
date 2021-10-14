package main

import (
	"log"
	"net/http"

	"./api"
)

// Main function.
func main() {

	var PORT = ":3003"

	// Serve static files.
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// All routes related to web sockets.
	http.HandleFunc("/ws", api.HandleWebSocket)

	// Routes related to users.
	http.HandleFunc("/api/sign-in", api.HandleSignIn)

	// Routes related to rooms.
	http.HandleFunc("/api/get-rooms", api.HandleGetRooms)
	http.HandleFunc("/api/create-room", api.HandleCreateRoom)
	http.HandleFunc("/api/join-room", api.HandleJoinRoom)
	http.HandleFunc("/api/leave-rooms", api.HandleLeaveRooms)
	http.HandleFunc("/api/create-message", api.HandleCreateMessage)

	log.Println("lchat server listening at " + PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
