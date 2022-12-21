package category

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	ecom "github.com/uacademy/e_commerce/catalog_service/proto-gen/e_commerce"
	"github.com/uacademy/e_commerce/catalog_service/storage"

	"context"
)

type categoryService struct {
	stg storage.StorageI
	ecom.UnimplementedCategoryServiceServer
}

// NewCategoryService ...
func NewCategoryService(stg storage.StorageI) *categoryService {
	return &categoryService{
		stg: stg,
	}
}

func (s *categoryService) CreateCategory(ctx context.Context, req *ecom.CreateCategoryRequest) (*ecom.Category, error) {
	id := uuid.New()

	err := s.stg.CreateCategory(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.CreateCategory: %s", err.Error())
	}

	category, err := s.stg.GetCategoryById(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryById: %s", err.Error())
	}

	return &ecom.Category{
		Id:            category.Id,
		CategoryTitle: category.CategoryTitle,
		CreatedAt:     category.CreatedAt,
		UpdatedAt:     category.UpdatedAt,
	}, nil
}

func (s *categoryService) GetCategoryList(ctx context.Context, req *ecom.GetCategoryListRequest) (*ecom.GetCategoryListResponse, error) {
	res, err := s.stg.GetCategoryList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryList: %s", err.Error())
	}

	return res, nil
}

func (s *categoryService) UpdateCategory(ctx context.Context, req *ecom.UpdateCategoryRequest) (*ecom.Category, error) {
	err := s.stg.UpdateCategory(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateCategory: %s", err.Error())
	}

	category, err := s.stg.GetCategoryById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryById: %s", err.Error())
	}

	return &ecom.Category{
		Id:            category.Id,
		CategoryTitle: category.CategoryTitle,
		CreatedAt:     category.CreatedAt,
		UpdatedAt:     category.UpdatedAt,
	}, nil
}

func (s *categoryService) DeleteCategory(ctx context.Context, req *ecom.DeleteCategoryRequest) (*ecom.Category, error) {
	category, err := s.stg.GetCategoryById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryById: %s", err.Error())
	}

	err = s.stg.DeleteCategory(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteCategory: %s", err.Error())
	}

	return &ecom.Category{
		Id:            category.Id,
		CategoryTitle: category.CategoryTitle,
		CreatedAt:     category.CreatedAt,
		UpdatedAt:     category.UpdatedAt,
	}, nil
}

func (s *categoryService) GetCategoryById(ctx context.Context, req *ecom.GetCategoryByIdRequest) (*ecom.GetCategoryByIdResponse, error) {
	category, err := s.stg.GetCategoryById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryById: %s", err.Error())
	}
	return category, nil
}
