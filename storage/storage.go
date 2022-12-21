package storage

import ecom "github.com/uacademy/e_commerce/catalog_service/proto-gen/e_commerce"

type StorageI interface {
	CreateProduct(id string, input *ecom.CreateProductRequest) error
	GetProductList(offset, limit int, search string) (resp *ecom.GetProductListResponse, err error)
	UpdateProduct(input *ecom.UpdateProductRequest) error
	GetProductById(id string) (resp *ecom.GetProductByIdResponse, err error)
	DeleteProduct(id string) error

	CreateCategory(id string, input *ecom.CreateCategoryRequest) error
	GetCategoryList(offset, limit int, search string) (resp *ecom.GetCategoryListResponse, err error)
	UpdateCategory(input *ecom.UpdateCategoryRequest) error
	GetCategoryById(id string) (resp *ecom.GetCategoryByIdResponse, err error)
	DeleteCategory(id string) error
}
