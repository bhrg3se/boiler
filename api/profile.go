package api

import (
	"boiler/models"
	"boiler/utils"
	"net/http"
)


func GetProfile(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value("user").(*models.User)

	if user == nil {
		utils.ErrorResponse(writer, "could not get user", 500)
		return
	}
	utils.SuccessResponse(writer, user, 200)
}
