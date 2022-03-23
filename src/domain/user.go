package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID         string
	Login      string
	Password   string
	Data       map[string]string
	Created_at *time.Time
	Updated_at *time.Time
}

func (u *User) Init() {
	uid, _ := uuid.NewUUID()
	u.ID = uid.String()

	curTime := time.Now()
	u.Created_at = &curTime
	u.Updated_at = &curTime
}

func (u *User) Update(pwd string) {
	u.Password = pwd

	curTime := time.Now()
	u.Updated_at = &curTime
}
