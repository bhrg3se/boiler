package ping

import (
	"boiler/models"
	"boiler/utils"
	"net/http"
)

func Ping(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value("user").(models.User)
	msg := PingResponse{Message: "ok " + user.Email}
	utils.SuccessResponse(writer, msg, 200)
}
