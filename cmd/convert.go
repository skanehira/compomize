package cmd

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh/terminal"
	"k8s.io/client-go/kubernetes/scheme"

	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use: "convert",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			input []byte
			err   error
		)
		if terminal.IsTerminal(int(os.Stdin.Fd())) {
			if len(args) < 1 {
				return errors.New("required file name")
			}
			input, err = ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}
		} else {
			input, err = ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
		}
		obj, _, err := scheme.Codecs.UniversalDeserializer().Decode(input, nil, nil)
		if err != nil {
			return err
		}

		if err := json.NewEncoder(os.Stdout).Encode(&obj); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
}
