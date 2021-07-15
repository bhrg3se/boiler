package users

import (
	"encoding/json"

	"boiler/models"
	"boiler/store"
	"boiler/utils"
	"net/http"
)

func PremiumContent(writer http.ResponseWriter, request *http.Request) {
	utils.SuccessResponse(writer, "", 200)
}

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		utils.ErrorResponse(writer, "json parse error", 400)
		return
	}

	err = saveUser(&store.State, &user)
	if err != nil {
		utils.ErrorResponse(writer, "could not save data", 500)
		return
	}
	utils.SuccessResponse(writer, user, 200)

}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	userID := request.URL.Query().Get("userID")

	user, err := FetchUser(&store.State, userID)
	if err != nil {
		utils.ErrorResponse(writer, "could not get user", 500)
		return
	}
	utils.SuccessResponse(writer, user, 200)

}
