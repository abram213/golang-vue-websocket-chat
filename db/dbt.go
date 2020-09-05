package db

import (
	"chat/model"
)

//DatabaseTest struct for testing
type DatabaseTest struct {
	User  model.User
	Users []model.User
}

//User methods
func (db *DatabaseTest) GetUserById(id uint) (*model.User, error) {
	db.User.ID = id
	return &db.User, nil
}

func (db *DatabaseTest) GetUserByUsername(username string) (*model.User, error) {
	db.User.Username = username
	return &db.User, nil
}

func (db *DatabaseTest) CreateUser(user *model.User) error {
	return nil
}

func (db *DatabaseTest) UserExistByUsername(username string) bool {
	if db.Users != nil {
		for _, user := range db.Users {
			if user.Username == username {
				return true
			}
		}
	}
	return false
}

func (db *DatabaseTest) UserExistByID(id uint) bool {
	if db.Users != nil {
		for _, user := range db.Users {
			if user.ID == id {
				return true
			}
		}
	}
	return false
}

func (db *DatabaseTest) AddUserFriend(userID, friendID uint) error {
	return nil
}

func (db *DatabaseTest) DeleteUserFriend(userID, friendID uint) error {
	return nil
}

func (db *DatabaseTest) FriendshipExist(userID, friendID uint) bool {
	if db.Users != nil {
		for _, user := range db.Users {
			if user.ID == userID && user.Friends != nil {
				for _, friend := range user.Friends {
					if friend.ID == friendID {
						return true
					}
				}
			}
		}
	}
	return false
}

func (db *DatabaseTest) GetUsers() (*[]model.User, error) {
	return &db.Users, nil
}

func (db *DatabaseTest) GetUsersExcept(id uint) (*[]model.User, error) {
	return &db.Users, nil
}

func (db *DatabaseTest) GetUserFriends(id uint) (*[]model.User, error) {
	return &db.Users, nil
}

//Chat methods
func (db *DatabaseTest) CreateChat(chat *model.Chat) (string, error) {
	return chat.Identifier, nil
}

func (db *DatabaseTest) GetChatByID(id uint) (*model.Chat, error) {
	chat := &model.Chat{
		Model: model.Model{ID: id},
	}
	return chat, nil
}

func (db *DatabaseTest) GetChatByIdentifier(i string) (*model.Chat, error) {
	chat := &model.Chat{
		Identifier: i,
	}
	return chat, nil
}
func (db *DatabaseTest) GetUserChats(id uint) (*[]model.Chat, error) {
	chats := []model.Chat{
		{UserID: id},
		{UserID: id},
		{UserID: id},
	}
	return &chats, nil
}

func (db *DatabaseTest) DeleteChat(id uint) error {
	return nil
}

//Message methods
func (db *DatabaseTest) CreateMessage(m *model.Message) error {
	return nil
}

//Util methods
func (db *DatabaseTest) Migrate(values ...interface{}) {}

func (db *DatabaseTest) DropTables(values ...interface{}) {}

func (db *DatabaseTest) CloseDB() error {
	return nil
}
