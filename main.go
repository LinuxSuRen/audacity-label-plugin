package main

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"github.com/linuxsuren/audacity-label-plugin/pkg"
	"os"
)

func main() {
	cmd := cobra.Command{
		Use:  "label",
		Args: cobra.ExactValidArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			file := args[0]
			var data []byte
			if data, err = ioutil.ReadFile(file); err != nil {
				return
			}

			cmd.Println(pkg.Convert(string(data)))
			return
		},
	}
	//cmd.AddCommand(version.NewVersionCmd("linuxsuren", "audacity-label-plugin", "label", nil))
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
