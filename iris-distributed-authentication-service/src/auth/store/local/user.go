package local

import (
	"time"
)

type User struct {
	ID       int64     `json:"id" gorm:"primaryKey;column:id"`
	Username string    `json:"username" gorm:"unique"`
	Password string    `json:"-"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

var user = "t_user"

func (m *Manager) InsertUser(u User) error {
	return m.handler.Table(user).Create(&u).Error
}

func (m *Manager) DeleteUserByUsername(username string) error {
	return m.handler.Table(user).Where("username = ?", username).Delete(&User{}).Error
}

func (m *Manager) SelectUser(username string) (User, error) {
	var u User
	err := m.handler.Table(user).Where("username = ?", username).Take(&u).Error
	return u, err
}

func (m *Manager) UpdateUser(u User) error {
	return m.handler.Table(user).Where("id = ? and username = ?", u.ID, u.Username).Updates(&u).Error
}
