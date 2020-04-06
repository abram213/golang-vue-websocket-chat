package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/tidwall/gjson"

	"chat/app"
	"chat/model"
)

func (a *API) authRouter() http.Handler {
	r := chi.NewRouter()
	r.Method("POST", "/sign_up", a.handler(a.SignUp))
	r.Method("POST", "/sign_in", a.handler(a.SignIn))
	r.Method("POST", "/refresh_token", a.handler(a.RefreshToken))
	return r
}

type authInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *API) SignIn(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var input authInput
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	userTokens, err := a.App.AuthUser(input.Username, input.Password)
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(userTokens)
	return nil
}

func (a *API) SignUp(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var input authInput

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	user := &model.User{
		Username: input.Username,
		Password: input.Password,
	}
	if err := ctx.CreateUser(user); err != nil {
		return err
	}
	userTokens, err := a.App.AuthUser(input.Username, input.Password)
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(userTokens)
	return nil
}

func (a *API) RefreshToken(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	token, err := jwtauth.VerifyRequest(a.App.Auth, r, TokenFromRequestBody)
	if err != nil {
		return err
	}

	userTokens, err := a.App.AuthUserByToken(token)
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(userTokens)
	return nil
}

//TokenAuth route middleware. Receives JWT token from header, verify, validate and pass it to context
func TokenAuth(ja *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfunc := func(w http.ResponseWriter, r *http.Request) {
			token, err := jwtauth.VerifyRequest(ja, r, jwtauth.TokenFromHeader)
			if err != nil {
				http.Error(w, err.Error(), 401)
				return
			}

			if token == nil || !token.Valid {
				http.Error(w, http.StatusText(401), 401)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, jwtauth.TokenCtxKey, token)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(hfunc)
	}
}

//TokenFromContext get token from context. return token, claims and claimsOk if token isn`t nil and claims is valid
func TokenFromContext(ctx context.Context) (*jwt.Token, jwt.MapClaims, bool) {
	token, _ := ctx.Value(jwtauth.TokenCtxKey).(*jwt.Token)

	var claims jwt.MapClaims
	var claimsOk bool
	if token != nil {
		if tokenClaims, ok := token.Claims.(jwt.MapClaims); ok {
			claims = tokenClaims
			claimsOk = true
		} else {
			panic(fmt.Sprintf("jwtauth: unknown type of Claims: %T", token.Claims))
		}
	} else {
		claims = jwt.MapClaims{}
	}
	return token, claims, claimsOk
}

// TokenFromRequestBody tries to retreive the token string from the "jwt"
// request body.
func TokenFromRequestBody(r *http.Request) string {
	body, _ := ioutil.ReadAll(r.Body)
	return gjson.Get(string(body), "jwt").Str
}
