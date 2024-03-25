package services

import (
	"arw-api/db"
	"arw-api/models"
)

func CreateUser(reqUser models.FullUser) (models.User, error) {
	var user models.User

	err := db.GetDB().QueryRow("INSERT INTO users(email, password) VALUES($1, $2) RETURNING id, email", reqUser.Email, reqUser.Password).Scan(&user.ID, &user.Email)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetFullUserByEmail(email string) (models.FullUser, error) {
	var fullUser models.FullUser

	err := db.GetDB().QueryRow("SELECT id, email, password FROM users WHERE email = $1", email).Scan(&fullUser.ID, &fullUser.Email, &fullUser.Password)
	if err != nil {
		return models.FullUser{}, err
	}

	return fullUser, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	err := db.GetDB().QueryRow("SELECT id, email FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
