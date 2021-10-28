package store

import (
	"boiler/models"
	"crypto/rsa"
	"database/sql"
)

type RealStore struct {
	db     *sql.DB
	config models.Config
	jwtKey struct {
		public  *rsa.PublicKey
		private *rsa.PrivateKey
	}
}


func (rs RealStore) CreateUser(user *models.User) error {
	_,err := rs.db.Exec(`INSERT INTO users (id,email,password,name,created_at,updated_at) VALUES ($1,$2,$3,$4,now(),now())`,
		user.ID,user.Email,user.Password,user.Name)
	return err
}

func (rs RealStore) FetchUser(userID string) (*models.User, error) {
	var user models.User
	err := rs.db.QueryRow(`SELECT id,name,email,created_at,updated_at FROM users where id=$1`, userID).
		Scan(&user.ID, &user.Name, &user.Email,&user.CreatedAt,&user.UpdatedAt)
	return &user, err
}

func (rs RealStore) FetchUserWithPassword(email string) (*models.User, error) {
	var user models.User
	err := rs.db.QueryRow(`SELECT id,email,password,name FROM users where email=$1`, email).
		Scan(&user.ID, &user.Email, &user.Password, &user.Name)
	return &user, err
}


// GetConfig returns config
func (rs RealStore) GetConfig() models.Config {
	return rs.config
}

// GetJWTPrivateKey gets the private key used for generating JWT tokens
func (rs RealStore) GetJWTPrivateKey() *rsa.PrivateKey {
	return rs.jwtKey.private
}

// GetJWTPublicKey gets the private key used to verify JWT tokens
func (rs RealStore) GetJWTPublicKey() *rsa.PublicKey {
	return rs.jwtKey.public
}
