package cmd

import (
	"github.com/spf13/cobra"
)

//RootCmd ...
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "CLI TODO LIST",
}
