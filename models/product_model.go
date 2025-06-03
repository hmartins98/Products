package models

type Product struct {
	Id          int64
	Name        string
	TypeId      int64
	Price       float64
	Description string
	ReviewScore float64
}
