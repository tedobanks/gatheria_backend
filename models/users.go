package models

import (
	"errors"
	"fmt"
	"time"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	Id             int64     `json:"user_id"`
	Email          string    `json:"email" binding:"required,email"`
	Password       string    `json:"password" binding:"required"`
	UserName       string    `json:"username"`
	FirstName      string    `json:"firstname"`
	LastName       string    `json:"lastname"`
	VerifiedStatus bool      `json:"is_verified"`
	Role           string    `json:"role"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (user *User) Save() (int64, error) {
	query := `
		INSERT INTO users(email, password, username, firstname, lastname, isverified, role, created, updatedat)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("unable to prepare query")
		return 0, err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		fmt.Println("hashing unsuccessful")
		return 0, err
	}

	result, err := stmt.Exec(user.Email, hashedPassword, user.UserName, user.FirstName, user.LastName, user.VerifiedStatus, user.Role, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		fmt.Println("failed at execution")
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("failed to retrieve last insert id")
		return 0, err
	}

	return lastId, err
}

func GetAllUsers() ([]User, error) {
	var users []User
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Email, &user.Password, &user.UserName, &user.FirstName, &user.LastName, &user.VerifiedStatus, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUser(userId int64) (*User, error) {
	var user User
	query := "SELECT * FROM users WHERE id = ?"
	row := db.DB.QueryRow(query, userId)

	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.UserName, &user.FirstName, &user.LastName, &user.VerifiedStatus, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		fmt.Println("failed at scan")
		return nil, err
	}

	return &user, nil
}

func (user *User) LoginUser() error {
	query := "SELECT id,password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.Id, &retrievedPassword)
	if err != nil {
		fmt.Println("failed at scan")
		fmt.Println(err)
		return err
	}

	isValidPassword := utils.CheckPasswordHash(retrievedPassword, user.Password)

	if !isValidPassword {
		fmt.Println("INVALID CREDENTIALS")
		return errors.New("INVALID CREDENTIALS")
	}

	return nil
}
