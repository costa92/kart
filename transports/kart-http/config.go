package kart_http

import "github.com/gin-gonic/gin"

type HttpConfig struct {
	Name            string   `yaml:"name" json:"name" toml:"name"`
	Port            int      `yaml:"port" json:"port" toml:"port"`
	Mode            string   `yaml:"mode" json:"mode" toml:"mode"`
	ReadTimeout     int      `json:"read_timeout" yaml:"read_timeout" toml:"read_timeout"`
	WriteTimeout    int      `json:"write_timeout" yaml:"write_timeout" toml:"write_timeout"`
	Healthz         bool     `json:"healthz" yaml:"healthz" toml:"healthz"`
	EnableMetrics   bool     `json:"enable-metrics" yaml:"enable-metrics" toml:"enable-metrics"`
	Middlewares     []string `json:"middlewares" yaml:"middlewares" toml:"middlewares"`
	EnableProfiling bool     `json:"enable_profiling" yaml:"enable_profiling"`
}

type ServerConfig struct {
	Mode            string
	Middlewares     []string
	Healthz         bool
	EnableProfiling bool
	EnableMetrics   bool
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Healthz:         true,
		Mode:            gin.ReleaseMode,
		Middlewares:     []string{},
		EnableProfiling: true,
		EnableMetrics:   true,
	}
}

// CompletedConfig is the completed configuration for GenericAPIServer.
type CompletedConfig struct {
	*ServerConfig
}

func (c *ServerConfig) Complete() CompletedConfig {
	return CompletedConfig{c}
}

func (c CompletedConfig) New() {
}
