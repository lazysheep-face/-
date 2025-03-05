package cart

import (
	"context"

	"gorm.io/gorm"
)

type CartServiceImpl struct {
	DB *gorm.DB
}

func (s *CartServiceImpl) AddItem(ctx context.Context, req *AddItemReq) (*AddItemResp, error) {
	cartItem := CartItem{
		UserID:    req.UserId,
		ProductID: req.Item.ProductId,
		Quantity:  req.Item.Quantity,
	}
	if err := s.DB.Create(&cartItem).Error; err != nil {
		return nil, err
	}
	return &AddItemResp{}, nil
}

func (s *CartServiceImpl) GetCart(ctx context.Context, req *GetCartReq) (*GetCartResp, error) {
	var items []CartItem
	s.DB.Where("user_id = ?", req.UserId).Find(&items)
	return &GetCartResp{Cart: &Cart{UserId: req.UserId, Items: items}}, nil
}
