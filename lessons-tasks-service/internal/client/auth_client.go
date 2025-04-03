// lessons-tasks-service/internal/client/auth_client.go
package client

import (
	"context"
)

type AuthClient struct{}

func NewAuthClient() *AuthClient {
	return &AuthClient{}
}

func (c *AuthClient) VerifyToken(ctx context.Context, token string) (string, error) {
	// Заглушка для тестирования
	return "test-user-id", nil
}
