package models

import (
	"errors"
	"studi_kasus_xyz/configs"
	"studi_kasus_xyz/entities"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Private
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func verifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func processAuth(id int, username string) (response entities.User, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	access_token, err := token.SignedString([]byte(configs.GetAccessKey()))
	if err != nil {
		return entities.User{}, err
	}

	response = entities.User{
		Id:          id,
		AccessToken: access_token,
		Username:    username,
	}

	return response, nil
}

// PUBLIC
func Auth(data entities.Auth) (response entities.User, err error) {
	var id_db int
	var username_db string
	var password_db string
	userRow := dbPool.QueryRow("SELECT id, username, password FROM user WHERE username = ?", data.Username)
	err = userRow.Scan(&id_db, &username_db, &password_db)
	if err != nil {
		return entities.User{}, errors.New("user not found")
	}

	// auth
	if verifyPassword(data.Password, password_db) {
		result, err := processAuth(id_db, username_db)
		if err != nil {
			return entities.User{}, err
		}
		return result, nil
	}
	return entities.User{}, errors.New("user not found")
}

func ChangePass(data entities.ChangePass) error {
	var username_db string
	var password_db string
	userRow := dbPool.QueryRow("SELECT username, password FROM user WHERE username = ?", data.Username)
	err := userRow.Scan(&username_db, &password_db)
	if err != nil {
		return err
	}

	// auth
	if verifyPassword(data.Password, password_db) {
		hashedPassword, err := hashPassword(data.NewPassword)
		if err != nil {
			return err
		}

		_, err = dbPool.Query("UPDATE user SET password = ? WHERE username = ?", hashedPassword, data.Username)
		if err != nil {
			return err
		}

		return nil
	}
	return err
}
