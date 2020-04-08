package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	Username string  `json:"username" gorm:"unique;"`
	Password string  `json:"-"`
	Friends  []*User `gorm:"many2many:friendships;association_jointable_foreignkey:friend_id"`
}

type Friendship struct {
	UserID   uint
	FriendID uint
}

type Tokens struct {
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	AccessExpiresIn int64  `json:"access_expires_in"`
}

//HashPassword hashes user password
func (u *User) HashPassword() error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}
	u.Password = string(hashedPass)
	return nil
}

//PasswordIsValid check user password
func (u *User) PasswordIsValid(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) RefreshTokens(tokenAuth *jwtauth.JWTAuth) (*Tokens, error) {
	//set access claims
	accessClaims := jwt.MapClaims{"user_id": u.ID}
	accessExpInTime := time.Now().Add(time.Minute * 60)
	jwtauth.SetExpiry(accessClaims, accessExpInTime)

	//set refresh claims
	refreshClaims := jwt.MapClaims{"user_id": u.ID}
	refreshExpInTime := time.Now().Add(time.Hour * 24 * 30)
	jwtauth.SetExpiry(refreshClaims, refreshExpInTime)

	_, accessToken, err := tokenAuth.Encode(accessClaims)
	if err != nil {
		return nil, err
	}
	_, refreshToken, err := tokenAuth.Encode(refreshClaims)
	if err != nil {
		return nil, err
	}
	return &Tokens{
		AccessToken:     accessToken,
		RefreshToken:    refreshToken,
		AccessExpiresIn: accessExpInTime.Unix(),
	}, nil
}
