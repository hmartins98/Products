package models

import (
	"database/sql"

	entities "Products/entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) CreateProduct(prod *entities.Product) (bool, error) {
	_, err := productModel.Db.Query("CALL public.\"CreateProduct\"(?, ?, ?, ?)", prod.Name, prod.TypeId, prod.Price, prod.Description)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (productModel ProductModel) ReadProduct(productId int64) (*entities.Product, error) {
	rows, err := productModel.Db.Query("SELECT * FROM public.\"ReadProduct\"(?)", productId)
	if err != nil {
		return nil, err
	} else {
		product := &entities.Product{}
		for rows.Next() {
			var _id int64
			var _name string
			var _type_id int32
			var _price float32
			var _description string
			err2 := rows.Scan(&_id, &_name, &_type_id, &_price, &_description)
			if err2 != nil {
				return &entities.Product{}, err2
			} else {
				product = &entities.Product{Id: _id, Name: _name, TypeId: _type_id, Price: _price, Description: _description, ReviewScore: 0.0}
			}
		}
		return product, nil
	}
}

func (productModel ProductModel) UpdateProduct(min, max float64) (bool, error) {
	// rows, err := productModel.Db.Query("CALL UpdateProduct(?, ?)", min, max)
	// if err != nil {
	// 	return nil, err
	// } else {
	// 	products := []entities.Product{}
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
	// 			product := entities.Product{id, name, price, quantity, status}
	// 			products = append(products, product)
	// 		}
	// 	}
	// 	return products, nil
	// }
	return false, nil
}

func (productModel ProductModel) DeleteProduct(min, max float64) (bool, error) {
	// rows, err := productModel.Db.Query("CALL DeleteProduct(?, ?)", min, max)
	// if err != nil {
	// 	return nil, err
	// } else {
	// 	products := []entities.Product{}
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
	// 			product := entities.Product{id, name, price, quantity, status}
	// 			products = append(products, product)
	// 		}
	// 	}
	// 	return products, nil
	// }
	return false, nil
}
