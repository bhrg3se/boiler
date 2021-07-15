package auth

import (
	"boiler/models"
	"boiler/store"
)

func getAuthDetail(state *store.Store, email string) (models.User, error) {
	var user models.User
	err := state.DB.QueryRow(`SELECT id,email,password,username FROM users where email=$1`, email).
		Scan(&user.ID, &user.Email, &user.Password, &user.Username)
	return user, err
}
