package auth

import (
	"boiler/models"
	"boiler/store"
	"boiler/utils"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		utils.ErrorResponse(writer, "invalid json data", 400)
		return
	}

	user, err := getAuthDetail(&store.State, req.Email)
	if err != nil {
		logrus.Debug(err)
		utils.ErrorResponse(writer, "user not found", 404)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		utils.ErrorResponse(writer, "incorrect password", 401)
		return
	}

	token, err := signData(user.ID)
	if err != nil {
		logrus.Debug(err)
		utils.ErrorResponse(writer, "could not generate token", 500)
		return
	}

	utils.SuccessResponse(writer, token, 200)
}

func signData(userID string) (string, error) {
	privateBytes := []byte(strings.TrimSpace(store.State.Config.Keys.Private))

	block, _ := pem.Decode(privateBytes)
	if block != nil {
		privateBytes = block.Bytes
	}

	signingKey, err := x509.ParsePKCS1PrivateKey(privateBytes)
	if err != nil {
		return "", fmt.Errorf("unable to load private key: %v", err)
	}

	// Create the Claims
	claims := models.JWTToken{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Hour * 48),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signed, err := token.SignedString(signingKey)
	if err != nil {
		logrus.Error("could not sign: ", err)
		return "", err
	}

	return signed, err
}
