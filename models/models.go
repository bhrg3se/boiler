package models

import "github.com/golang-jwt/jwt"

type Config struct {
	Database struct {
		User         string `toml:"user"`
		Password     string `toml:"password"`
		Host         string `toml:"host"`
		Port         string `toml:"port"`
		Name         string `toml:"name"`
		SSL          bool   `toml:"ssl"`
		CaCertPath   string `json:"caCertPath"`
		UserCertPath string `json:"userCertPath"`
		UserKeyPath  string `json:"userKeyPath"`
	} `toml:"database"`

	Logging struct {
		Level string `toml:"logging"`
	} `toml:"logging"`

	Server struct {
		Listen string `toml:"listen"`
	} `toml:"server"`

}

type JWTToken struct {
	jwt.StandardClaims
	UserID string `json:"userID"`
}
