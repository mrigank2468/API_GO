package models

import (
	"errors"

	"github.com/mrigank2468/API_GO/db"
	"github.com/mrigank2468/API_GO/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err

}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email=?"
	row := db.DB.QueryRow(query, u.Email)
	var retrivedPassword string
	err := row.Scan(&u.ID,&retrivedPassword)
	if err != nil {
		return err
	}
	passwordIsValid := utils.CheckPassword(u.Password, retrivedPassword)
	if !passwordIsValid {
		return errors.New("invalid password")
	}
	return nil
}
