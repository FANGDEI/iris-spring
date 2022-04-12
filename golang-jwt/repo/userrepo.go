package repo

import (
	"errors"
	"golang-jwt/database"
	"golang-jwt/model"
	"golang-jwt/model/dto"
)

var sqlDB = database.GetMySQLDB()

type UserRepository interface {
	SelectUserInfoById(id int32) (model.User, error)
	SelectUserPasswordByUsername(username string) (string, error)
	InsertUser(userDto dto.UserDto) (int64, error)
}

type userRepository struct {
}

// InsertUser implements UserRepository
func (userRepo *userRepository) InsertUser(userDto dto.UserDto) (int64, error) {
	result, err := sqlDB.Exec("insert into user (username, password, email, permission) values (?, ?, ?, ?)",
		userDto.UserName, userDto.Password, userDto.Email, 0)
	if err != nil {
		return 0, err
	}
	cnt, _ := result.RowsAffected()
	if cnt == 0 {
		return cnt, errors.New("something wrong when insert user")
	}
	return cnt, nil
}

// SelectUserInfoById implements UserRepository
func (userRepo *userRepository) SelectUserInfoById(id int32) (model.User, error) {
	row := sqlDB.QueryRow("select id, username, password, email, phone, permission from user where id = ?", id)
	var user model.User
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Permission)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// SelectUserPasswordByUsername implements UserRepository
func (userRepo *userRepository) SelectUserPasswordByUsername(username string) (string, error) {
	row := sqlDB.QueryRow("select password from user where username = ?", username)
	var password string
	err := row.Scan(&password)
	if err != nil {
		return "", nil
	}
	return password, nil
}

func NewBookRepository() UserRepository {
	return &userRepository{}
}
