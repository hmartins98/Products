package services

import (
	context "context"

	models "hmartins98/Products/models"
	repository "hmartins98/Products/repositories"

	contract "hmartins98/Products/contracts"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductService struct {
	repo repository.ProductRepo
	contract.UnimplementedProductsContractServer
}

func NewProductService(repo repository.ProductRepo) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *contract.Product) (*emptypb.Empty, error) {
	productEntity := &models.Product{
		Id:          req.Id,
		Name:        req.Name,
		TypeId:      req.TypeId,
		Price:       req.Price,
		Description: req.Description,
		ReviewScore: req.ReviewScore,
	}

	err := s.repo.CreateProduct(productEntity)
	return &emptypb.Empty{}, err
}

func (s *ProductService) ReadProduct(ctx context.Context, req *contract.ProductId) (*contract.Product, error) {
	productEntity, err := s.repo.ReadProduct(req.Id)

	productResult := &contract.Product{
		Id:          productEntity.Id,
		Name:        productEntity.Name,
		TypeId:      productEntity.TypeId,
		Price:       productEntity.Price,
		Description: productEntity.Description,
		ReviewScore: productEntity.ReviewScore,
	}

	return productResult, err
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *contract.Product) (*emptypb.Empty, error) {

	// db, err := config.GetPostgresDB()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	ProductRepository := models.ProductRepository{
	// 		Db: db,
	// 	}
	// 	fmt.Println("Product List")
	// 	products, err2 := ProductRepository.Search(250, 500)
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

	return &emptypb.Empty{}, nil
}

func (*ProductService) DeleteProduct(ctx context.Context, req *contract.ProductId) (*emptypb.Empty, error) {

	// db, err := config.GetPostgresDB()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	ProductRepository := models.ProductRepository{
	// 		Db: db,
	// 	}
	// 	fmt.Println("Product List")
	// 	products, err2 := ProductRepository.Search(250, 500)
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

	return &emptypb.Empty{}, nil
}
