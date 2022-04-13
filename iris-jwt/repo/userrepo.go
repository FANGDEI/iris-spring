package repo

import (
	"database/sql"
	"errors"
	"iris-jwt/database"
	"iris-jwt/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB = database.GetMySQLDB()

func SelectUserInfoById(id int) (model.User, error) {
	var user model.User
	row := DB.QueryRow("select id, username, password, email, phone, permission from user where id = ?", id)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Permission)
	if err != nil {
		log.Println("get the user information err: ", err)
		return model.User{}, nil
	}
	return user, nil
}

func SelectUserInformationByUsername(username string) (model.User, error) {
	var user model.User
	row := DB.QueryRow("select id, username, password, email, phone, permission from user where username = ?", username)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Permission)
	if err != nil {
		log.Println("get the user information err: ", err)
		return model.User{}, nil
	}
	return user, nil
}

func InsertUser(username, password, email string) (int64, error) {
	result, err := DB.Exec("insert into user (username, password, email, permission) values (?, ?, ?, ?)",
		username, password, email, 0)
	if err != nil {
		return 0, err
	}
	cnt, _ := result.RowsAffected()
	if cnt == 0 {
		return cnt, errors.New("something wrong when insert user")
	}
	return cnt, nil
}
