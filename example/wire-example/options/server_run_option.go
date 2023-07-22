package options

type ServerRunOption struct {
	Name        string   `json:"name" mapstructure:"name"`
	Mode        string   `json:"mode"        mapstructure:"mode"`
	Healthz     bool     `json:"healthz"     mapstructure:"healthz"`
	Middlewares []string `json:"middlewares" mapstructure:"middlewares"`
}

func NewServerRunOption() *ServerRunOption {
	return &ServerRunOption{}
}
