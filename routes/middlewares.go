package routes

import (
	"boiler/models"
	"boiler/store"
	"boiler/utils"
	"context"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Authentication middleware lets only authenticated user to access the API.
// It reads "auth" token from cookies, parses and verifies it.
// Then it fetches user details mapped to that token and adds it into context
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("auth")
		if err != nil {
			logrus.Debug(err)
			utils.ErrorResponse(w, "cookies not found", 401)
			return
		}

		parsedToken, err := utils.ParseAuthToken(c.Value, store.State.GetJWTPublicKey())
		if err != nil {
			logrus.Debug(err)
			utils.ErrorResponse(w, "failed to verify token", 401)
			return
		}

		user, err := store.State.FetchUser(parsedToken.UserID)
		if err != nil {
			logrus.Error("could not fetch user: ", err)
			utils.ErrorResponse(w, "could not fetch user", 500)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		next(w, r.WithContext(ctx))
	})

}

func AdminContent(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//user := r.Context().Value("user").(models.User)
		//if !user.IsAdmin {
		//	utils.ErrorResponse(w, "not authorised", 403)
		//	return
		//}
		next(w, r)
	})
}

func AuthenticationWS(next func(ws *websocket.Conn, user *models.User)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		parsedToken, err := utils.ParseAuthToken(token, store.State.GetJWTPublicKey())
		if err != nil {
			logrus.Error(err)
			utils.ErrorResponse(w, "failed to verify token", 401)
			return
		}
		user, err := store.State.FetchUser(parsedToken.UserID)
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

		next(conn, user)
	})

}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		//TODO
		return true
	},
	//Subprotocols: []string{"chat"},
}
