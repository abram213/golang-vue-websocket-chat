package api

import (
	"chat/app"
	"chat/model"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var (
	broadcast = make(chan model.Message)
	upgrader  = websocket.Upgrader{}
	clients   = map[string][]client{}
)

type client struct {
	Conn   *websocket.Conn
	UserID uint
}

func (a *API) HandleWSConnection(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	id, ok := r.URL.Query()["id"]
	if !ok || len(id) == 0 {
		return errors.New("url param 'id' is missing")
	}

	strUserID, ok := r.URL.Query()["user_id"]
	if !ok || len(strUserID) == 0 {
		return errors.New("url param 'user_id' is missing")
	}

	userID, err := strconv.Atoi(strUserID[0])
	if err != nil {
		return errors.New("err parse string to int")
	}

	clients[id[0]] = append(clients[id[0]], client{
		Conn:   ws,
		UserID: uint(userID),
	})

	for {
		var msg model.Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			//delete(clients[id[0]], uint(userID))
			return err
		}
		if err := ctx.CreateMessage(&msg); err != nil {
			return err
		}

		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
	return nil
}

func HandleMessages() {
	for {
		msg := <-broadcast
		for _, client := range clients[msg.ChatIdentifier] {
			if err := client.Conn.WriteJSON(msg); err != nil {
				fmt.Println(err)
				//delete(clients[msg.ChatIdentifier], uint())
				//chat.Clients[msg.ChatIdentifier] = append(chat.Clients[msg.ChatIdentifier][:i], chat.Clients[msg.ChatIdentifier][i+1:]...)
			}
		}
	}
}
