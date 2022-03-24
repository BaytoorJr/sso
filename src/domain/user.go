package domain

import (
	"github.com/google/uuid"
	"time"

)

type User struct {
	ID         string            `json:"ID"`
	Login      string            `json:"login"`
	Password   string            `json:"password"`
	Data       map[string]string `json:"data"`
	Created_at *time.Time        `json:"created_At"`
	Updated_at *time.Time        `json:"updated_At"`
}

func (u *User) Init(login, password string) {
	uid, _ := uuid.NewUUID()
	u.ID = uid.String()

	u.Login = login
	u.Password = password

	curTime := time.Now()
	u.Created_at = &curTime
	u.Updated_at = &curTime
}

func (u *User) Update(pwd string) {
	u.Password = pwd

	curTime := time.Now()
	u.Updated_at = &curTime
}
