package app

import "testing"

func Test_CobraCommand(t *testing.T) {
	cmd := NewCommand("kart", "kart")
	cobra := cmd.CobraCommand()
	cobra.Execute()
}
