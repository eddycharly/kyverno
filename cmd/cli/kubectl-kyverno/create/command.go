package create

import (
	"github.com/kyverno/kyverno/cmd/cli/kubectl-kyverno/create/test"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	cmd.AddCommand(test.Command())
	return cmd
}
