package cmd

type CliOptions interface {
	Flags() (fss NamedFlagSets)
}
