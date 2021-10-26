package api

import (
	"boiler/models"
	"boiler/store"
	"boiler/utils"
	"github.com/google/uuid"
	"github.com/nbutton23/zxcvbn-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)


// Signup registers new user into system
func Signup(writer http.ResponseWriter, request *http.Request) {
	var user models.User

	err := utils.ParseAndValidateRequest(request,&user)
	if err != nil {
		utils.ErrorResponse(writer, "invalid data: "+err.Error(), 400)
		return
	}
	passwordStrength := zxcvbn.PasswordStrength(user.Password, nil)
	if passwordStrength.Score<2 {
		utils.ErrorResponse(writer, "password too weak", 400)
		return
	}


	user.ID=uuid.New().String()
	hashedPass,err:=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(writer, "could not hash password", 500)
		return
	}
	user.Password=string(hashedPass)

	err=store.State.CreateUser(&user)
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(writer, "could not create user", 500)
		return
	}

	// TODO verify email

	utils.SuccessResponse(writer, "user created", 200)
}
