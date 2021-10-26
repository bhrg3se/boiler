package routes

import (
	"boiler/store"
	"boiler/utils"
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Authentication middleware lets only authenticated user to access the API.
// It reads "auth" token from cookies, parses and verifies it.
// Then it fetches user details mapped to that token and adds it into context
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c,err:=r.Cookie("auth")
		if err != nil {
			logrus.Debug(err)
			utils.ErrorResponse(w, "cookies not found", 401)
			return
		}

		parsedToken, err := utils.ParseAuthToken(c.Value,store.State.GetJWTPublicKey())
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

