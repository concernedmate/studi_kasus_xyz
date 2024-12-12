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

func processAuth(username string, grup string) (response entities.User, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"grup":     grup,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	access_token, err := token.SignedString([]byte(configs.GetAccessKey()))
	if err != nil {
		return entities.User{}, err
	}

	response = entities.User{
		AccessToken: access_token,
		Username:    username,
		Grup:        grup,
	}

	return response, nil
}

// PUBLIC
func Auth(data entities.Auth) (response entities.User, err error) {
	var username_db string
	var password_db string
	var grup_db string
	userRow := dbPool.QueryRow("SELECT Username, Password, Grup FROM users u LEFT JOIN grup g ON u.id_grup = g.id WHERE u.username = ?", data.Username)
	err = userRow.Scan(&username_db, &password_db, &grup_db)
	if err != nil {
		return entities.User{}, errors.New("user not found")
	}

	// auth
	if verifyPassword(data.Password, password_db) {
		result, err := processAuth(username_db, grup_db)
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
	userRow := dbPool.QueryRow("SELECT Username, Password FROM users WHERE username = ?", data.Username)
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

		_, err = dbPool.Query("UPDATE users SET password = ? WHERE username = ?", hashedPassword, data.Username)
		if err != nil {
			return err
		}

		return nil
	}
	return err
}
