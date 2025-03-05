package payment

import (
	"context"

	"github.com/google/uuid"
)

type PaymentServiceImpl struct{}

func (s *PaymentServiceImpl) Charge(ctx context.Context, req *ChargeReq) (*ChargeResp, error) {
	// 模拟支付处理
	return &ChargeResp{
		TransactionId: uuid.New().String(),
	}, nil
}
