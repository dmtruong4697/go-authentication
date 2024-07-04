package controllers

import (
	"encoding/json"
	"go-authentication/src/database"
	"go-authentication/src/models"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var channels = make(map[uint64]map[*websocket.Conn]bool)
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

type Message struct {
	ChannelID  uint   `json:"channelid"`
	SenderID   uint   `json:"senderid"`
	ReceiverID uint   `json:"receiverid"`
	Content    string `json:"content"`
}

type MessageRequestBody struct {
	ChannelID uint `json:"channelid"`
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	channelIDStr := query.Get("channelid")

	channelID, err := strconv.ParseUint(channelIDStr, 10, 64)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer ws.Close()

	if channels[channelID] == nil {
		channels[channelID] = make(map[*websocket.Conn]bool)
	}
	channels[channelID][ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(channels[channelID], ws)
			break
		}

		database.DB.Create(&msg)

		for client := range channels[channelID] {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(channels[channelID], client)
			}
		}
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast

		channel := msg.ChannelID

		channelID64 := uint64(channel)

		for client := range channels[channelID64] {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(channels[channelID64], client)
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
	database.DB.Where("channel_id = ?", messageRequestBody.ChannelID).Find(&messages)
	json.NewEncoder(w).Encode(messages)
}
