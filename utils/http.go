package utils

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validatorInstance *validator.Validate

func init() {
	validatorInstance = validator.New()
}

func SuccessResponse(writer http.ResponseWriter, data interface{}, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	resp := apiResponse{
		Status:  true,
		Message: "",
		Data:    data,
	}
	marshalledResp, _ := json.Marshal(resp)
	writer.Write(marshalledResp)
}

func ErrorResponse(writer http.ResponseWriter, message string, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	data := apiResponse{Status: false, Message: message}
	marshalledData, _ := json.Marshal(data)
	writer.Write(marshalledData)
}

//ParseAndValidateRequest unmarshalls request body into given struct and also verify json fields
func ParseAndValidateRequest(r *http.Request, reqStruct interface{}) error {

	if err := json.NewDecoder(r.Body).Decode(reqStruct); err != nil {
		return err
	}
	if err := validatorInstance.Struct(reqStruct); err != nil {
		return err
	}
	return nil
}

type apiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
