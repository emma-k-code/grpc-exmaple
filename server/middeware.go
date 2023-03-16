package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// CheckUser 檢查帶入使用者資料
func CheckUser() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		var user string

		if mt, ok := metadata.FromIncomingContext(ctx); ok {
			if data, exist := mt["user"]; exist {
				for _, v := range data {
					user = v
					break
				}
			}
		}

		if user == "" {
			//nolint:wrapcheck // 回傳 grpc 錯誤
			return nil, status.Error(codes.Code(101), "無帶入使用者資料")
		}

		if user != "test" {
			//nolint:wrapcheck // 回傳 grpc 錯誤
			return nil, status.Error(codes.Code(102), "非支援使用者")
		}

		return handler(ctx, req)
	}
}
