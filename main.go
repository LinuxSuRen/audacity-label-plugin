package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	cmd := cobra.Command{
		Use: "label",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			file := args[0]
			var data []byte
			data, err = ioutil.ReadFile(file)
			tokens := strings.Split(string(data), "\n")
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
	err := cmd.Execute()
	if err != nil {
		panic(err)
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
