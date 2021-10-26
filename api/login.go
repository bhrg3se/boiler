package api

import (
	"boiler/models"
	"boiler/store"
	"boiler/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// Login using email and password
func Login(writer http.ResponseWriter, request *http.Request) {
	var req models.User

	err := utils.ParseAndValidateRequest(request,&req)
	if err != nil {
		utils.ErrorResponse(writer, "invalid data: "+err.Error(), 400)
		return
	}

	// get hashed password from database
	user, err := store.State.FetchUserWithPassword(req.Email)
	if err != nil {
		logrus.Debug(err)
		utils.ErrorResponse(writer, "user not found", 404)
		return
	}

	// check password by hashing it
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		utils.ErrorResponse(writer, "incorrect password", 401)
		return
	}

	//generate JWT token by signing userID
	token, err := utils.GenerateAuthToken(user.ID,store.State.GetJWTPrivateKey())
	if err != nil {
		logrus.Debug(err)
		utils.ErrorResponse(writer, "could not generate token", 500)
		return
	}

	// set cookies
	c:=http.Cookie{
		Name:       "auth",
		Value:      token,
		Path: "/",
		//Domain:     "",
		Expires:    time.Now().Add(time.Hour*24*7),
		// TODO make it secure
		//Secure:     true,
		HttpOnly:   true,
	}

	http.SetCookie(writer, &c)

	utils.SuccessResponse(writer, "login successful", 200)
}

// Logout clears cookies
func Logout(writer http.ResponseWriter, request *http.Request) {
	// set cookies
	c:=http.Cookie{
		Name:       "auth",
		Value:      "",
		Path: "/",
		//Domain:     "",
		Expires:    time.Now(),
		// TODO make it secure
		//Secure:     true,
		HttpOnly:   true,
	}

	http.SetCookie(writer, &c)

	utils.SuccessResponse(writer, "logout successful", 200)
}

