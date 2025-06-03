package repositories

import (
	"database/sql"

	models "hmartins98/Products/models"
)

type ProductRepo interface {
	CreateProduct(prod *models.Product) error
	ReadProduct(productId int64) (*models.Product, error)
	UpdateProduct(min, max float64) error
	DeleteProduct(min, max float64) error
}

type PostgresProductRepo struct {
	db *sql.DB
}

func NewProductPostgresRepo(db *sql.DB) *PostgresProductRepo {
	return &PostgresProductRepo{db: db}
}

func (p PostgresProductRepo) CreateProduct(prod *models.Product) error {
	_, err := p.db.Query("CALL public.\"CreateProduct\"(?, ?, ?, ?)", prod.Name, prod.TypeId, prod.Price, prod.Description)
	return err
}

func (p PostgresProductRepo) ReadProduct(productId int64) (*models.Product, error) {
	rows, err := p.db.Query("SELECT * FROM public.\"ReadProduct\"(?)", productId)
	if err != nil {
		return nil, err
	} else {
		product := &models.Product{}
		for rows.Next() {
			var _id int64
			var _name string
			var _type_id int64
			var _price float64
			var _description string
			err2 := rows.Scan(&_id, &_name, &_type_id, &_price, &_description)
			if err2 != nil {
				return &models.Product{}, err2
			} else {
				product = &models.Product{Id: _id, Name: _name, TypeId: _type_id, Price: _price, Description: _description, ReviewScore: 0.0}
			}
		}
		return product, nil
	}
}

func (p PostgresProductRepo) UpdateProduct(min, max float64) error {
	// rows, err := productModel.Db.Query("CALL UpdateProduct(?, ?)", min, max)
	// if err != nil {
	// 	return nil, err
	// } else {
	// 	products := []models.Product{}
	// 	for rows.Next() {
	// 		var id int64
	// 		var name string
	// 		var price float32
	// 		var quantity int
	// 		var status bool
	// 		err2 := rows.Scan(&id, &name, &price, &quantity, &status)
	// 		if err2 != nil {
	// 			return nil, err2
	// 		} else {
	// 			product := models.Product{id, name, price, quantity, status}
	// 			products = append(products, product)
	// 		}
	// 	}
	// 	return products, nil
	// }
	return nil
}

func (p PostgresProductRepo) DeleteProduct(min, max float64) error {
	// rows, err := productModel.Db.Query("CALL DeleteProduct(?, ?)", min, max)
	// if err != nil {
	// 	return nil, err
	// } else {
	// 	products := []models.Product{}
	// 	for rows.Next() {
	// 		var id int64
	// 		var name string
	// 		var price float32
	// 		var quantity int
	// 		var status bool
	// 		err2 := rows.Scan(&id, &name, &price, &quantity, &status)
	// 		if err2 != nil {
	// 			return nil, err2
	// 		} else {
	// 			product := models.Product{id, name, price, quantity, status}
	// 			products = append(products, product)
	// 		}
	// 	}
	// 	return products, nil
	// }
	return nil
}
