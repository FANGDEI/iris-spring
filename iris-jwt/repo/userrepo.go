package repo

import (
	"database/sql"
	"iris-jwt/database"
	"iris-jwt/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB = database.GetMySQLDB()

func SelectUserInfoById(id int) model.User {
	var user model.User
	row := DB.QueryRow("select id, username, password, email, phone, permission from user where id = ?", id)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Permission)
	if err != nil {
		log.Println("get the user information err: ", err)
	}
	return user
}

func SelectUserInformationByUsername(username string) model.User {
	var user model.User
	row := DB.QueryRow("select id, username, password, email, phone, permission from user where username = ?", username)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Permission)
	if err != nil {
		log.Println("get the user information err: ", err)
	}
	return user
}
