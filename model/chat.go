package model

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/jinzhu/gorm"
	"time"
)

type Chat struct {
	Model
	Identifier string `json:"identifier"`
	Title    string    `json:"title"`
	UserID   uint      `json:"user_id"`
	Messages []Message `json:"messages"`
	Users    []*User   `gorm:"many2many:chat_users;association_autoupdate:false;association_autocreate:false" json:"users"`
}

type ChatUser struct {
	ChatID uint
	UserID uint
}

func (c *Chat) BeforeCreate(scope *gorm.Scope) error {
	h := sha1.New()
	h.Write([]byte(time.Now().String()))
	scope.SetColumn("Identifier", hex.EncodeToString(h.Sum(nil)))
	return nil
}