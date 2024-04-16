package user

type UserCreateDAO struct {
	Name     string `form:"name" binding:"required"`
	Username string `form:"username" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
	Profile  int8   `form:"profile" binding:"required"`
}

type UserUpdateDAO struct {
	Name     string `form:"name"`
	Username string `form:"username"`
}