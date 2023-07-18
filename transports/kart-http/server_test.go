package kart_http

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
)

type testKey struct{}

func Test_GinNewServer(t *testing.T) {
	handler := gin.Default()
	config := &HttpConfig{
		Port:          "8080",
		Healthz:       true,
		EnableMetrics: true,
	}
	src := NewServer(
		WithGinEngin(handler),
		WithConfig(config),
	)
	if src == nil {
		t.Error("Server is nil")
	}
	ctx := context.WithValue(context.Background(), testKey{}, "test")
	if err := src.Start(ctx); err != nil {
		t.Error(err)
	}
}
