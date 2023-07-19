package app

import "kart-io/kart/example/cmd/cflag"

type CliOptions interface {
	Flags() (fss cflag.NamedFlagSets)
}
