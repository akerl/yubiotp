package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/yawn/ykoath"
)

func listRunner(_ *cobra.Command, _ []string) error {
	yubikeys, err := ykoath.NewSet()
	if err != nil {
		return err
	}

	for _, y := range yubikeys {
		serial, err := y.Serial()
		if err != nil {
			return nil
		}

		_, err = y.Select()
		defer y.Close()
		if err != nil {
			return err
		}

		otps, err := y.List()
		if err != nil {
			return err
		}
		for _, o := range otps {
			fmt.Printf("%s %s\n", o.Name, serial)
		}
	}

	return nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available MFA codes",
	RunE:  listRunner,
}

func init() {
	rootCmd.AddCommand(listCmd)
}
