package command

import "kart-io/kart/pkg/cliflag"

type CliOptions interface {
	Flags() (fss cliflag.NamedFlagSets)
}
