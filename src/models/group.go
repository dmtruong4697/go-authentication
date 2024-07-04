package models

type Group struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Type      string `json:"type"`
	MemberIds []uint `json:"memberids"`
	Name      string `json:"name"`
}
