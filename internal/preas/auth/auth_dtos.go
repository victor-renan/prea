package auth

import "time"

type LoginPayload struct {
	Username  string     `json:"username"`
	Profile   int8       `json:"profile"`
	LastLogin *time.Time `time_format:"2006-01-02 01:01:10" json:"lastlogin"`
}

type LoginResponse struct {
	Token   string       `json:"token"`
	Payload LoginPayload `json:"payload"`
}

type LoginBody struct {
	Username string `binding:"required" form:"username"`
	Password string `binding:"required" form:"password"`
}
