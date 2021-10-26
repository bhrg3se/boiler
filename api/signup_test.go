package api

import (
	"boiler/models"
	"boiler/store"
	"boiler/testutils"
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignup(t *testing.T) {
	mockStore:=new(store.MockStore)
	store.State=mockStore
	mockStore.On("CreateUser",mock.Anything).Return(nil).Times(1)
	type args struct {
		request *http.Request
	}
	tests := []struct {
		name string
		args args
		WantStatus int
		WantMessage string
	}{
		{
			name: "should return 400 when email is empty",
			args: struct{ request *http.Request }{request: testutils.GetPOSTRequest(models.User{
				Email:     "ssss",
				Name:      "SOME NAME",
				Password:  "12345678wertabsidashidsadj9uUI76T&G*(",
			},"")},
			WantStatus: 400,
			WantMessage: "invalid data: Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag",

		},
		{
			name: "should return 400 when password is weak",
			args: struct{ request *http.Request }{request: testutils.GetPOSTRequest(models.User{
				Email:     "",
				Name:      "SOME NAME",
				Password:  "12345678",
			},"")},
			WantStatus: 400,
			WantMessage: "password too weak",
		},

{
			name: "should return 200",
			args: struct{ request *http.Request }{request: testutils.GetPOSTRequest(models.User{
				Email:     "test@example.com",
				Name:      "SOME NAME",
				Password:  "12345678wertabsidashidsadj9uUI76T&G*(",
			},"")},
			WantStatus: 200,
			WantMessage: "",
		},





		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			writer:=httptest.NewRecorder()
			Signup(writer,tt.args.request)

			if tt.WantStatus!=writer.Code{
				t.Errorf("Signup() got status = %v, wantStatus %v", writer.Code, tt.WantStatus)
				return
			}
			if tt.WantStatus!=200{
				return
			}

			var resp struct {
				Status  bool        `json:"status"`
				Message string      `json:"message,omitempty"`
				Data    interface{} `json:"data,omitempty"`
			}
			err:=json.Unmarshal(writer.Body.Bytes(),&resp)
			if err != nil {
				t.Errorf("Unmarshal() got err: %v", err)
				return
			}

			if tt.WantMessage!=resp.Message{
				t.Errorf("Signup() got data.message = %v, wanted %v", resp.Message, tt.WantMessage)
			}


		})
	}
}
