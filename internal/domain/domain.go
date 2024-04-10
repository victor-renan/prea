package domain

type IModel interface {
	Table() string
	Pk() string
}