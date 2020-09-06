package app

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"

	"chat/model"
)

func (app *App) AuthUser(username, password string) (*model.Tokens, error) {
	user, err := app.usernamePassMatch(username, password)
	if err != nil {
		return nil, err
	}
	return user.RefreshTokens(app.Auth)
}

func (app *App) AuthUserByToken(token *jwt.Token) (*model.Tokens, error) {
	var claims jwt.MapClaims
	if token != nil {
		if tokenClaims, ok := token.Claims.(jwt.MapClaims); ok {
			claims = tokenClaims
		} else {
			return nil, fmt.Errorf("jwtauth: unknown type of Claims: %T", token.Claims)
		}
	} else {
		return nil, errors.New("token == nil")
	}

	userID, err := strconv.Atoi(fmt.Sprintf("%s", claims["user_id"]))
	if err != nil {
		return nil, fmt.Errorf("bad user_id, err: %v", err)
	}
	user, err := app.GetUserById(uint(userID))
	if err != nil {
		return nil, fmt.Errorf("problem with user, err: %v", err)
	}

	return user.RefreshTokens(app.Auth)
}

func (app *App) usernamePassMatch(username, password string) (*model.User, *AuthError) {
	user, err := app.Database.GetUserByUsername(username)
	if err != nil {
		return nil, &AuthError{err.Error()}
	}
	if !user.PasswordIsValid(password) {
		return nil, &AuthError{"password is not valid"}
	}
	return user, nil
}
