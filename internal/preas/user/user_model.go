package user

import (
	"time"
)

type User struct {
	Id        string     `json:"id" form:"-"`
	Name      string     `json:"name" form:"name" binding:"required"`
	Username  string     `json:"username" form:"username" binding:"required"`
	Email     string     `json:"email" form:"email" binding:"required"`
	Password  string     `json:"-" form:"password" binding:"required"`
	Profile   int8       `json:"profile" form:"profile" binding:"required"`
	Token     *string    `json:"token" form:"token"`
	LastLogin *time.Time `json:"lastlogin" time_format:"2006-01-01 00:00:00"`
}

func (User) Table() string {
	return "users"
}
