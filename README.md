# lChat
A simple group chat app implemented using WebSockets.<br>
Demo [lchat.lovesaroha](https://lchat.lovesaroha.com)

## Features
- Create chat rooms.
- Make room private with password.
- Search available rooms in lobby.
- Send messages in room.

## Requirements
- Go 1.9 or higher. We aim to support the 3 latest versions of Go.
- Gorilla web socket.
- JWT go.

## Docker Go SDK
Simple install the package to your [$GOPATH](https://github.com/golang/go/wiki/GOPATH "GOPATH") with the [go tool](https://golang.org/cmd/go/ "go command") from shell:
```bash
go get -u "github.com/gorilla/websocket"
```
```bash
go get -u "github.com/dgrijalva/jwt-go"
```
Make sure [Git is installed](https://git-scm.com/downloads) on your machine and in your system's `PATH`.

## Usage
```bash
go run main.go
```

![image](https://raw.githubusercontent.com/lovesaroha/gimages/main/103.png)

---

![image](https://raw.githubusercontent.com/lovesaroha/gimages/main/104.png)

---

![image](https://raw.githubusercontent.com/lovesaroha/gimages/main/105.png)

---

![image](https://raw.githubusercontent.com/lovesaroha/gimages/main/106.png)

---

![image](https://raw.githubusercontent.com/lovesaroha/gimages/main/107.png)