package models

type Message struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Channel    string `json:"channel"`
	SenderID   uint   `json:"senderid"`
	ReceiverID uint   `json:"receiverid"`
	Content    string `json:"content"`
}
