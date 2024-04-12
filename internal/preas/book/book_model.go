package book

type Book struct {
	Id          int64  `json:"id"`
	Name        string `json:"name" form:"name" binding:"required"`
	Author      string `json:"author" form:"author" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
}

func (Book) Table() string {
	return "books"
}

func (Book) Pk() string {
	return "Id"
}
