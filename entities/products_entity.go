package entities

type Product struct {
	Id          int64
	Name        string
	TypeId      int32
	Price       float32
	Description string
	ReviewScore float32
}
