package users

import (
	"boiler/models"
	"boiler/store"
)

func FetchUser(state *store.Store, userID string) (models.User, error) {
	var user models.User
	err := state.DB.QueryRow(`SELECT id,username,email FROM users where id=$1`, userID).
		Scan(&user.ID, &user.Username, &user.Email)
	return user, err
}

func saveUser(state *store.Store, user *models.User) error {
	_, err := state.DB.Exec(`INSERT INTO users (id, username, email,password) VALUES ($1,$2,$3,$4)`, user.ID, user.Username, user.Email, user.Password)
	return err
}
