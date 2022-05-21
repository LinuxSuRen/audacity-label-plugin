package pkg

import "fmt"
import "strings"
import "strconv"

func Convert(data string) (result string) {
	text := strings.ReplaceAll(data, "\r", "\n")
	tokens := strings.Split(text, "\n")

	for i := range tokens {
		token := tokens[i]
		lineTokens := strings.SplitN(token, "\t", 3)
		if len(lineTokens) != 3 {
			continue
		}
		result += fmt.Sprintf("%s\t%s\n", secondsToHuman(lineTokens[0]), lineTokens[2])
	}
	return
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
