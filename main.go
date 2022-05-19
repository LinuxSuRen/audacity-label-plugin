package main

import (
	"fmt"
	"github.com/linuxsuren/cobra-extension/version"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	cmd := cobra.Command{
		Use:  "label",
		Args: cobra.ExactValidArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			file := args[0]
			var data []byte
			data, err = ioutil.ReadFile(file)

			text := strings.ReplaceAll(string(data), "\r", "\n")
			tokens := strings.Split(text, "\n")
			for i := range tokens {
				token := tokens[i]
				lineTokens := strings.SplitN(token, "\t", 3)
				if len(lineTokens) != 3 {
					continue
				}
				cmd.Printf("%s\t%s\n", secondsToHuman(lineTokens[0]), lineTokens[2])
			}
			return
		},
	}
	cmd.AddCommand(version.NewVersionCmd("linuxsuren", "audacity-label-plugin", "label", nil))
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func secondsToHuman(seconds string) (result string) {
	num, err := strconv.ParseFloat(seconds, 8)
	if err == nil {
		sec := int(num) % 60
		min := int(num) / 60
		result = fmt.Sprintf("* %s:%s", twoDigitalNum(min), twoDigitalNum(sec))
	}
	return
}

func twoDigitalNum(num int) string {
	if num >= 10 {
		return fmt.Sprintf("%d", num)
	}
	return fmt.Sprintf("0%d", num)
}
