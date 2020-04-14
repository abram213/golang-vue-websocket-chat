package model

type Message struct {
	Model
	Body   string `json:"body" gorm:"type:text"`
	UserID uint   `json:"user_id"`
	ChatIdentifier string   `json:"chat_identifier"`
	User   User   `json:"user"`
}
