package require

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

// NoArgs returns an error if any args are included.
func NoArgs(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("%q accepts no arguments\n\nUsage: %s",
			cmd.CommandPath(),
			cmd.UseLine(),
		)
	}
	return nil
}

func SpecificArg(cmd *cobra.Command, arg string, validArgs []string) error {
	if slices.Contains(validArgs, arg) {
		return nil

	}
	return fmt.Errorf("error")

}

// ExactArgs returns an error if there are not exactly n args.
func ExactArgs(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) != n {
			return fmt.Errorf(
				"%q requires %d %s\n\nUsage:  %s",
				cmd.CommandPath(),
				n,
				pluralize("argument", n),
				cmd.UseLine(),
			)
		}
		return nil
	}
}

// MaximumNArgs returns an error if there are more than N args.
func MaximumNArgs(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) > n {
			return fmt.Errorf(
				"%q accepts at most %d %s\n\nUsage:  %s",
				cmd.CommandPath(),
				n,
				pluralize("argument", n),
				cmd.UseLine(),
			)
		}
		return nil
	}
}

// MinimumNArgs returns an error if there is not at least N args.
func MinimumNArgs(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < n {
			return fmt.Errorf(
				"%q requires at least %d %s\n\nUsage:  %s",
				cmd.CommandPath(),
				n,
				pluralize("argument", n),
				cmd.UseLine(),
			)
		}
		return nil
	}
}

func pluralize(word string, n int) string {
	if n == 1 {
		return word
	}
	return word + "s"
}
