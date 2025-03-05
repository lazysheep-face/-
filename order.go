package order

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderServiceImpl struct {
	DB *gorm.DB
}

func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *PlaceOrderReq) (*PlaceOrderResp, error) {
	orderID := uuid.New().String()
	order := Order{
		OrderID:     orderID,
		UserID:      req.UserId,
		TotalAmount: calculateTotal(req.OrderItems), // 计算总金额
		Status:      "pending",
	}
	if err := s.DB.Create(&order).Error; err != nil {
		return nil, err
	}
	return &PlaceOrderResp{Order: &OrderResult{OrderId: orderID}}, nil
}
