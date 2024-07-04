package models

import "time"

type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ChannelID uint      `json:"channelid"`
	SenderID  uint      `json:"senderid"`
	Content   string    `json:"content"`
	CreateAt  time.Time `json:"createat"`
}
