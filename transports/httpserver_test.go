package transports

import (
	"context"
	"kart-io/kart/transports/kart-http"
	options2 "kart-io/kart/transports/options"
	"testing"
)

type testKey struct{}

func Test_GinNewServer(t *testing.T) {

	opts := options2.NewInsecureServingOptions()

	config := kart_http.NewServerConfig()
	err := opts.ApplyTo(config)
	if err != nil {
		return
	}
	src := config.Complete().New()
	if src == nil {
		t.Error("Server is nil")
	}
	ctx := context.WithValue(context.Background(), testKey{}, "test")
	if err := src.Start(ctx); err != nil {
		t.Error(err)
	}
}