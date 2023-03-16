package client

import (
	"context"
	"example/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// GetSum 呼叫 grpc 範例
func GetSum(ctx context.Context, user string, a, b int64) (int64, error) {
	conn, err := grpc.DialContext(
		ctx, "localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()), // 不使用 TLS/SSL 連線
		grpc.WithBlock(), // 有設定timeout時, 成功連線至server前需先block
	)

	if err != nil {
		//nolint:wrapcheck // 回傳原始錯誤, 以利後續判斷
		return 0, err
	}

	client := pb.NewServiceClient(conn)

	result, err := client.Sum(
		metadata.NewOutgoingContext(ctx, metadata.Pairs("user", user)), // 類似http的header
		&pb.SumRequest{
			A: a, B: b,
		},
	)

	if err != nil {
		//nolint:wrapcheck // 回傳原始錯誤, 以利後續判斷
		return 0, err
	}

	return result.Result, nil
}
