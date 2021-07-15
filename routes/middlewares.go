package routes

import (
	"boiler/api/users"
	"boiler/models"
	"boiler/store"
	"boiler/utils"
	"context"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")
		if len(token) > 7 {
			token = token[7:]
		}

		parsedToken, err := parseJWTToken(token)
		if err != nil {
			logrus.Error(err)
			utils.ErrorResponse(w, "failed to verify token", 401)
			return
		}

		user, err := users.FetchUser(&store.State, parsedToken.UserID)
		if err != nil {
			logrus.Error("invalid user id in jwt token: ", err)
			utils.ErrorResponse(w, "invalid user_id", 500)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		next(w, r.WithContext(ctx))
	})

}
func AuthenticationWS(next func(ws *websocket.Conn, user *models.User)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		parsedToken, err := parseJWTToken(token)
		if err != nil {
			logrus.Error(err)
			utils.ErrorResponse(w, "failed to verify token", 401)
			return
		}

		user, err := users.FetchUser(&store.State, parsedToken.UserID)
		//user, err := store.DBState.GetUser(parsedToken.UserID)
		if err != nil {
			logrus.Error("invalid user id in jwt token, ", err)
			utils.ErrorResponse(w, "invalid user_id", 500)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logrus.Error(err)
			return
		}

		next(conn, &user)
	})

}

func PremiumUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(models.User)
		if !user.IsAdmin {
			utils.ErrorResponse(w, "not authorised for non premium users", 403)
			return
		}
		next(w, r)
	})
}

func parseJWTToken(signedData string) (*models.JWTToken, error) {
	var tokenStruct models.JWTToken

	pubKey := []byte(store.State.Config.Keys.Public)

	block, _ := pem.Decode(pubKey)
	if block != nil {
		pubKey = block.Bytes
	}

	verificationKey, err := x509.ParsePKIXPublicKey(pubKey)
	if err != nil {
		return nil, fmt.Errorf("unable to load public key: %s", err)
	}

	token, err := jwt.Parse(signedData, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return verificationKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not parse signed data: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tokenStruct.UserID = claims["userId"].(string)
	} else {
		return nil, errors.New("invalid token")
	}

	return &tokenStruct, nil
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		//TODO
		return true
	},
	//Subprotocols: []string{"chat"},
}
