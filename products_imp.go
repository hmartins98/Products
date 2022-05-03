package main

import (
	context "context"

	entities "Products/entities"
	models "Products/models"

	CustomTypes "github.com/hmartins98/Contracts/CustomTypes"
	Products "github.com/hmartins98/Contracts/Products"
)

type server struct {
	Products.UnimplementedProductsContractServer
}

func (*server) CreateProduct(ctx context.Context, req *Products.Product) (*CustomTypes.BOOL, error) {
	productModel := models.ProductModel{
		Db: db,
	}

	productEntity := &entities.Product{
		Id:          req.Id,
		Name:        req.Name,
		TypeId:      req.TypeId,
		Price:       req.Price,
		Description: req.Description,
		ReviewScore: req.ReviewScore,
	}

	result, err := productModel.CreateProduct(productEntity)
	return &CustomTypes.BOOL{Value: result}, err
}

func (*server) ReadProduct(ctx context.Context, req *Products.ProductId) (*Products.Product, error) {
	productModel := models.ProductModel{
		Db: db,
	}

	productEntity, err := productModel.ReadProduct(req.Id)

	productResult := &Products.Product{
		Id:          productEntity.Id,
		Name:        productEntity.Name,
		TypeId:      productEntity.TypeId,
		Price:       productEntity.Price,
		Description: productEntity.Description,
		ReviewScore: productEntity.ReviewScore,
	}

	return productResult, err
}

func (*server) UpdateProduct(ctx context.Context, req *Products.Product) (*CustomTypes.BOOL, error) {

	// db, err := config.GetPostgresDB()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	productModel := models.ProductModel{
	// 		Db: db,
	// 	}
	// 	fmt.Println("Product List")
	// 	products, err2 := productModel.Search(250, 500)
	// 	if err2 != nil {
	// 		fmt.Println(err2)
	// 	} else {
	// 		fmt.Print("Products: ", len(products), "\n")
	// 		for _, product := range products {
	// 			fmt.Println("Id:", product.Id)
	// 			fmt.Println("Name:", product.Name)
	// 			fmt.Println("Price:", product.Price)
	// 			fmt.Println("Quantity:", product.Quantity)
	// 			fmt.Println("Status:", product.Status)
	// 			fmt.Println("----------------------------")
	// 		}
	// 	}
	// }

	return &CustomTypes.BOOL{Value: false}, nil
}

func (*server) DeleteProduct(ctx context.Context, req *Products.ProductId) (*CustomTypes.BOOL, error) {

	// db, err := config.GetPostgresDB()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	productModel := models.ProductModel{
	// 		Db: db,
	// 	}
	// 	fmt.Println("Product List")
	// 	products, err2 := productModel.Search(250, 500)
	// 	if err2 != nil {
	// 		fmt.Println(err2)
	// 	} else {
	// 		fmt.Print("Products: ", len(products), "\n")
	// 		for _, product := range products {
	// 			fmt.Println("Id:", product.Id)
	// 			fmt.Println("Name:", product.Name)
	// 			fmt.Println("Price:", product.Price)
	// 			fmt.Println("Quantity:", product.Quantity)
	// 			fmt.Println("Status:", product.Status)
	// 			fmt.Println("----------------------------")
	// 		}
	// 	}
	// }

	return &CustomTypes.BOOL{Value: false}, nil
}
