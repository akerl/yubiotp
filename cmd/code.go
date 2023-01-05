package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/yawn/ykoath"
)

func codeRunner(_ *cobra.Command, argv []string) error {
	if len(argv) != 1 {
		return fmt.Errorf("provide only one arg")
	}
	name := argv[0]

	yubikeys, err := ykoath.NewSet()
	if err != nil {
		return err
	}

	for _, y := range yubikeys {
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
			if o.Name == name {
				code, err := y.Calculate(name, func(_ string) error {
					fmt.Fprintln(os.Stderr, "Touch the yubikey to use OTP")
					return nil
				})
				if err != nil {
					return err
				}
				fmt.Println(code)
				return nil
			}
		}
	}

	return fmt.Errorf("no matching code found")
}

var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "Fetch an MFA code from connected yubikeys",
	RunE:  codeRunner,
}

func init() {
	rootCmd.AddCommand(codeCmd)
}
