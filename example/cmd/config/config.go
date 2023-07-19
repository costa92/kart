package config

import "kart-io/kart/example/cmd/pkg"

type Config struct {
	*pkg.Options
}

// CreateConfigFromOptions creates a running configuration instance based
// on a given IAM pump command line or configuration file option.
func CreateConfigFromOptions(opts *pkg.Options) (*Config, error) {
	return &Config{opts}, nil
}
