package controllers

import (
	"encoding/json"
	"go-authentication/src/database"
	"go-authentication/src/models"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var channels = make(map[string]map[*websocket.Conn]bool)
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

type Message struct {
	Channel    string `json:"channel"`
	SenderID   uint   `json:"senderid"`
	ReceiverID uint   `json:"receiverid"`
	Content    string `json:"content"`
}

type MessageRequestBody struct {
	Channel string `json:"channel"`
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	channel := query.Get("channel")

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer ws.Close()

	if channels[channel] == nil {
		channels[channel] = make(map[*websocket.Conn]bool)
	}
	channels[channel][ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(channels[channel], ws)
			break
		}

		database.DB.Create(&msg)

		for client := range channels[channel] {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(channels[channel], client)
			}
		}
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast

		channel := msg.Channel

		for client := range channels[channel] {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(channels[channel], client)
			}
		}
	}
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var messageRequestBody MessageRequestBody
	err := json.NewDecoder(r.Body).Decode(&messageRequestBody)
	if err != nil {
		http.Error(w, "Failed to decode message request body", http.StatusBadRequest)
		return
	}

	var messages []models.Message
	database.DB.Where("channel = ?", messageRequestBody.Channel).Find(&messages)
	json.NewEncoder(w).Encode(messages)
}
