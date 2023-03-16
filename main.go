package main

import (
	"context"
	"example/client"
	"example/server"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("參數數量錯誤")
	}

	arg := os.Args[1]

	switch arg {
	case "server":
		server.Start()
	case "client":
		var user string

		fmt.Print("Enter user (僅支援test): ")
		fmt.Scanf("%s", &user)

		var a, b int64

		fmt.Print("Enter a & b: ")
		fmt.Scanf("%d %d", &a, &b)

		result, err := client.GetSum(context.Background(), user, a, b)
		fmt.Printf("result = %d\nerr = %v\n", result, err)
	}
}
