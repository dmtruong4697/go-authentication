package models

import "time"

type Channel struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatorID uint      `json:"creatorid"`
	CreatedAt time.Time `json:"createdat"`
	MemberIDs []uint    `gorm:"type:uint"`
}
