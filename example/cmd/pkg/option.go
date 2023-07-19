package pkg

import "kart-io/kart/example/cmd/cflag"

type Options struct {
}

func NewOptions() *Options {
	o := Options{}
	return &o
}

func (o *Options) Flags() (fss cflag.NamedFlagSets) {
	return fss
}
