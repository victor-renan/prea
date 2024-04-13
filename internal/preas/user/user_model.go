package user

import (
	"time"
)

type User struct {
	Id        string    `db:"id" db_type:"uuid" json:"id"`
	Name      string    `db:"name" db_type:"varchar" json:"name" form:"name" binding:"required"`
	Username  string    `db:"username" db_type:"varchar" json:"username" form:"username" binding:"required"`
	Email     string    `db:"email" db_type:"varchar" json:"email" form:"email" binding:"required"`
	Password  string    `db:"password" db_type:"varchar" json:"description" form:"password" binding:"required"`
	Profile   int       `db:"profile" db_type:"int" json:"profile" form:"profile" binding:"required"`
	Token     string    `db:"token" db_type:"text" json:"token" form:"token"`
	LastLogin time.Time `db:"lastlogin" db_type:"timestamp" json:"lastlogin"`
}

func (User) Table() string {
	return "users"
}

func (User) Pk() string {
	return "Id"
}
