package product

import (
	"context"

	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	DB *gorm.DB
}

func (s *ProductServiceImpl) ListProducts(ctx context.Context, req *ListProductsReq) (*ListProductsResp, error) {
	var products []Product
	query := s.DB.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))
	if req.CategoryName != "" {
		query = query.Where("categories LIKE ?", "%"+req.CategoryName+"%")
	}
	query.Find(&products)
	return &ListProductsResp{Products: products}, nil
}

func (s *ProductServiceImpl) GetProduct(ctx context.Context, req *GetProductReq) (*GetProductResp, error) {
	var product Product
	if err := s.DB.First(&product, req.Id).Error; err != nil {
		return nil, err
	}
	return &GetProductResp{Product: &product}, nil
}
