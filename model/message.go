package model

type Message struct {
	Model
	Body   string `json:"body" gorm:"type:text"`
	UserID uint   `json:"user_id"`
	ChatID uint   `json:"chat_id"`
	User   User   `json:"user"`
}
