package database

import (
	"context"
	"fmt"
	"log"
)

type User struct {
	Id       int64
	Username string
	Password string
}

// GetUsername gets the users password from the id
func (dbConn DatabaseConnection) GetPasswordById(id int64) (string, error) {
	var password string
	err := dbConn.conn.QueryRow(context.Background(),
		"SELECT password FROM users WHERE id = $1;", id).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

// GetUsername gets the users password from the id
func (dbConn DatabaseConnection) GetPasswordByUsername(username string) (string, error) {
	var password string
	err := dbConn.conn.QueryRow(context.Background(),
		"SELECT password FROM users WHERE username = $1;", username).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

// GetUsername gets the users password from the id
func (dbConn DatabaseConnection) GetIdFromUsername(username string) int64 {
	var id int64
	err := dbConn.conn.QueryRow(context.Background(),
		"SELECT id FROM users WHERE username = $1;", username).Scan(&id)
	if err != nil {
		log.Printf("Couldn't find username: %v", err)
		return 0
	}
	return id
}

// GetUsername gets the users username from the id
func (dbConn DatabaseConnection) GetUsername(id int64) (string, error) {
	var username string
	err := dbConn.conn.QueryRow(context.Background(),
		"SELECT username FROM users WHERE id = $1;", id).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}

// Verifys the password of a user
func (dbConn DatabaseConnection) VerifyPasswordByUsername(username, password string) bool {
	realPassword, err := dbConn.GetPasswordByUsername(username)
	if err != nil || realPassword == "" {
		return false // couldn't find the user?
	}
	return realPassword == password
}

// checks if a user exists
func (dbConn DatabaseConnection) CheckIfUserExists(username string) bool {
	row := dbConn.conn.QueryRow(context.Background(), "SELECT username FROM users WHERE username = $1;", username)
	var dbUser string
	_ = row.Scan(&dbUser)
	log.Printf("username:%v", dbUser)
	if dbUser == "" {
		return false
	}
	return true
}

// checks if a user exists
func (dbConn DatabaseConnection) CreateUser(username, password string) {
	row := dbConn.conn.QueryRow(context.Background(), "INSERT INTO users VALUES(DEFAULT, $1, $2)",
		username, password)
	err := row.Scan()
	if err != nil {
		fmt.Println(err)
	}
}
