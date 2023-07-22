package options

import (
	"kart-io/kart/pkg/cliflag"
)

type Options struct {
	ServerRunOption *ServerRunOption `json:"server"  mapstructure:"server"`
}

func NewOptions() *Options {
	o := Options{
		ServerRunOption: NewServerRunOption(),
	}
	return &o
}

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	return fss
}
