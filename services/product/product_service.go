package product

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	ecom "github.com/uacademy/e_commerce/catalog_service/proto-gen/e_commerce"
	"github.com/uacademy/e_commerce/catalog_service/storage"

	"context"
)

type productService struct {
	stg storage.StorageI
	ecom.UnimplementedProductServiceServer
}

// NewProductService ...
func NewProductService(stg storage.StorageI) *productService {
	return &productService{
		stg: stg,
	}
}

func (s *productService) CreateProduct(ctx context.Context, req *ecom.CreateProductRequest) (*ecom.Product, error) {
	id := uuid.New()

	err := s.stg.CreateProduct(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.CreateProduct: %s", err.Error())
	}

	product, err := s.stg.GetProductById(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetProductById: %s", err.Error())
	}

	return &ecom.Product{
		Id:         product.Id,
		CategoryId: product.Category.Id,
		Title:      product.Title,
		Descrip:    product.Descrip,
		Price:      product.Price,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}, nil
}

func (s *productService) GetProductList(ctx context.Context, req *ecom.GetProductListRequest) (*ecom.GetProductListResponse, error) {
	res, err := s.stg.GetProductList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetProductList: %s", err.Error())
	}

	return res, nil
}

func (s *productService) UpdateProduct(ctx context.Context, req *ecom.UpdateProductRequest) (*ecom.Product, error) {
	err := s.stg.UpdateProduct(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateProduct: %s", err.Error())
	}

	product, err := s.stg.GetProductById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetProductById: %s", err.Error())
	}

	return &ecom.Product{
		Id:        product.Id,
		Title:     product.Title,
		Descrip:   product.Descrip,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}, nil
}

func (s *productService) DeleteProduct(ctx context.Context, req *ecom.DeleteProductRequest) (*ecom.Product, error) {
	product, err := s.stg.GetProductById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetProductById: %s", err.Error())
	}

	err = s.stg.DeleteProduct(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteProduct: %s", err.Error())
	}

	return &ecom.Product{
		Id:        product.Id,
		Title:     product.Title,
		Descrip:   product.Descrip,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}, nil
}

func (s *productService) GetProductById(ctx context.Context, req *ecom.GetProductByIdRequest) (*ecom.GetProductByIdResponse, error) {
	product, err := s.stg.GetProductById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetProductById: %s", err.Error())
	}
	return product, nil
}
