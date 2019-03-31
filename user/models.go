package user

import (
	"database"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	ID           int
	Email        string
	PasswordHash string
}

func saveUser(m *UserModel) error {
	db := database.GetDB()
	sqlStatement := `INSERT INTO users (email,password_hash)
	VALUES ($2,$3)`
	_, err := db.Exec(sqlStatement, m.Email, m.PasswordHash)

	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	m.PasswordHash = string(passwordHash)
	return nil
}

func getUser(email string) (UserModel, error) {
	db := database.GetDB()
	sqlStatement := `SELECT email,password_hash FROM users WHERE email=$1`
	user, err := db.Exec(sqlStatement, email)
	return user, err
}

func (m *UserModel) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(m.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
