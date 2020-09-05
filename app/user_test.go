package app

import (
	"chat/db"
	"chat/model"
	"reflect"
	"testing"
)

func TestAppGetUserByID(t *testing.T) {
	tApp := App{Database: &db.DatabaseTest{}}
	var userID uint = 1

	user, err := tApp.GetUserById(userID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if user.ID != userID {
		t.Errorf("User id should be: %v, got: %v", userID, user.ID)
	}
}

func TestCreateUser(t *testing.T) {
	tApp := App{Database: &db.DatabaseTest{}}
	tCtx := tApp.NewContext()
	//Valid user input data
	vUser := &model.User{
		Username: "valid_username",
		Password: "valid_pass",
	}
	if err := tCtx.CreateUser(vUser); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	errCases := []struct {
		user     *model.User
		errorStr string
	}{
		//Check validation, username and password can`t be blank
		{
			user: &model.User{
				Username: "",
				Password: "",
			},
			errorStr: "-: cannot be blank; username: cannot be blank.",
		},
		//Check validation, username length must be between 5 and 30
		{
			user: &model.User{
				Username: "invu",
				Password: "valid_pass",
			},
			errorStr: "username: the length must be between 5 and 30.",
		},
		//Check validation, password length must be between 8 and 30
		{
			user: &model.User{
				Username: "valid_username",
				Password: "invalid_password_invalid_password_invalid_password_invalid_password",
			},
			errorStr: "-: the length must be between 8 and 30.",
		},
	}

	for caseNum, item := range errCases {
		err := tCtx.CreateUser(item.user)
		if err == nil {
			t.Errorf("[%d] expected error, got nil", caseNum)
		}
		verr, ok := err.(*ValidationError)
		if !ok {
			t.Errorf("[%d] expected ValidationError, got %T", caseNum, err)
		}
		if verr.Message != item.errorStr {
			t.Errorf("[%d] got validation error: %s, want: %s", caseNum, verr.Message, item.errorStr)
		}
	}
}

//Testing creating user error - "already exist user with such username"
func TestCreateUserExistError(t *testing.T) {
	vUser := model.User{
		Model:    model.Model{ID: 1},
		Username: "valid_username",
		Password: "valid_pass",
	}
	tApp := App{Database: &db.DatabaseTest{
		Users: []model.User{vUser},
	}}
	tCtx := tApp.NewContext()

	errorStr := "username: already exist user with such username."
	err := tCtx.CreateUser(&vUser)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	verr, ok := err.(*ValidationError)
	if !ok {
		t.Errorf("expected ValidationError, got %T", err)
	}
	if verr.Message != errorStr {
		t.Errorf("got validation error: %s, want: %s", verr.Message, errorStr)
	}
}

func TestGetUserByID(t *testing.T) {
	var userID uint = 1
	user := model.User{
		Model: model.Model{ID: userID},
	}
	tApp := App{Database: &db.DatabaseTest{User: user}}
	tCtx := tApp.NewContext()

	tUser, err := tCtx.GetUserById(userID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if tUser.ID != userID {
		t.Errorf("User id should be: %v, got: %v", userID, tUser.ID)
	}
}

func TestAddFriendByID(t *testing.T) {
	var friendID uint = 2
	friend := model.User{
		Model: model.Model{ID: friendID},
	}
	tApp := App{Database: &db.DatabaseTest{Users: []model.User{friend}}}
	var userID uint = 1
	user := &model.User{
		Model: model.Model{ID: userID},
	}
	tCtx := tApp.NewContext().WithUser(user)

	if err := tCtx.AddFriendByID(friendID); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestAddFriendByIDErrors(t *testing.T) {
	errCases := []struct {
		userID            uint
		friendID          uint
		wantToAddFriendID uint
		errorStr          string
	}{
		//To add friend, he should be in database
		{
			userID:            0,
			friendID:          1,
			wantToAddFriendID: 2,
			errorStr:          "no such user",
		},
		//You can`t add yourself to friends
		{
			userID:            1,
			friendID:          1,
			wantToAddFriendID: 1,
			errorStr:          "you can`t add yourself to friends",
		},
	}

	for caseNum, item := range errCases {
		friend := model.User{
			Model: model.Model{ID: item.friendID},
		}
		tApp := App{Database: &db.DatabaseTest{Users: []model.User{friend}}}
		user := &model.User{
			Model: model.Model{ID: item.userID},
		}
		tCtx := tApp.NewContext().WithUser(user)

		err := tCtx.AddFriendByID(item.wantToAddFriendID)
		if err == nil {
			t.Errorf("[%d] expected error, got nil", caseNum)
			return
		}
		if err.Error() != item.errorStr {
			t.Errorf("[%d] got error: %#v, want: %#v", caseNum, err.Error(), item.errorStr)
		}
	}
}

//Testing adding friend error when friendship already exist
func TestAddFriendByIDFriendshipExistError(t *testing.T) {
	var friendID uint = 2
	friend := model.User{
		Model: model.Model{ID: friendID},
	}
	var userID uint = 1
	user := model.User{
		Model:   model.Model{ID: userID},
		Friends: []*model.User{&friend},
	}
	tApp := App{Database: &db.DatabaseTest{Users: []model.User{friend, user}}}
	tCtx := tApp.NewContext().WithUser(&user)
	errorStr := "friendship already exist"
	err := tCtx.AddFriendByID(friendID)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err.Error() != errorStr {
		t.Errorf("got error: %#v, want: %#v", err.Error(), errorStr)
	}
}

func TestDeleteFriendByID(t *testing.T) {
	var friendID uint = 2
	friend := model.User{
		Model: model.Model{ID: friendID},
	}
	tApp := App{Database: &db.DatabaseTest{Users: []model.User{friend}}}
	var userID uint = 1
	user := &model.User{
		Model: model.Model{ID: userID},
	}
	tCtx := tApp.NewContext().WithUser(user)

	if err := tCtx.DeleteFriendByID(friendID); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestGetUsers(t *testing.T) {
	tUsers := []model.User{
		{Model: model.Model{ID: 1}, Username: "user1", Password: "user1"},
		{Model: model.Model{ID: 2}, Username: "user2", Password: "user2"},
		{Model: model.Model{ID: 3}, Username: "user3", Password: "user3"},
	}
	tApp := App{Database: &db.DatabaseTest{Users: tUsers}}
	user := &model.User{
		Model: model.Model{ID: 4},
	}
	tCtx := tApp.NewContext().WithUser(user)

	users, err := tCtx.GetUsers()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(&tUsers, users) {
		t.Errorf("\nUnexpected result got:\n %+v \nwant:\n %+v", users, &tUsers)
	}
}

func TestGetUserFriends(t *testing.T) {
	tUsers := []model.User{
		{Model: model.Model{ID: 1}, Username: "user1", Password: "user1"},
		{Model: model.Model{ID: 2}, Username: "user2", Password: "user2"},
		{Model: model.Model{ID: 3}, Username: "user3", Password: "user3"},
	}
	tApp := App{Database: &db.DatabaseTest{Users: tUsers}}
	user := &model.User{
		Model: model.Model{ID: 4},
	}
	tCtx := tApp.NewContext().WithUser(user)

	users, err := tCtx.GetUserFriends()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(&tUsers, users) {
		t.Errorf("\nUnexpected result got:\n %+v \nwant:\n %+v", users, &tUsers)
	}
}
