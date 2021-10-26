package api

import (
	"boiler/models"
	"boiler/store"
	"boiler/testutils"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"strings"
	"testing"
)

func  TestAPI(t *testing.T) {
	mockStore:=new(store.MockStore)
	store.State=mockStore



	// Signup
	signupData:=models.User{
		Email:     testutils.MockUser1.Email,
		Name:      testutils.MockUser1.Name,
		Password:  testutils.MockUser1Pass,
	}
	mockStore.On("CreateUser",mock.Anything).Return(nil).Times(1)
	signUpReq:=testutils.GetPOSTRequest(signupData,"")
	recorder:=httptest.NewRecorder()
	Signup(recorder,signUpReq)
	if recorder.Code!=200 {
		t.Error("status code in signup is not 200")
		return
	}


	// Login
	recorder=httptest.NewRecorder()
	loginData:=models.User{
		Email:     testutils.MockUser1.Email,
		Password:  testutils.MockUser1Pass,
	}
	mockStore.On("FetchUserWithPassword",loginData.Email).Return(&testutils.MockUser1,nil).Times(1)
	mockStore.On("GetJWTPrivateKey").Return(testutils.GetMockPrivateKey1()).Times(1)

	loginReq:=testutils.GetPOSTRequest(loginData,"")
	Login(recorder,loginReq)

	if recorder.Code!=200 {
		t.Errorf("status code in login is not 200, got: %v",recorder.Code)
		return
	}

	cookie:=recorder.Header().Get("Set-Cookie")
	token:=strings.Split(strings.Split(cookie,";")[0],"=")[1]

	if token==""{
		t.Error("token is empty")
	}

}
