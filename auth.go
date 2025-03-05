package auth

import (
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthServiceImpl struct{}

func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *DeliverTokenReq) (*DeliveryResp, error) {
	// 生成JWT令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   strconv.Itoa(int(req.UserId)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})
	signedToken, _ := token.SignedString([]byte("your_secret_key"))
	return &DeliveryResp{Token: signedToken}, nil
}

func (s *AuthServiceImpl) VerifyTokenByRPC(ctx context.Context, req *VerifyTokenReq) (*VerifyResp, error) {
	// 验证JWT令牌
	token, err := jwt.Parse(req.Token, func(t *jwt.Token) (interface{}, error) {
		return []byte("your_secret_key"), nil
	})
	return &VerifyResp{Res: token.Valid}, err
}
