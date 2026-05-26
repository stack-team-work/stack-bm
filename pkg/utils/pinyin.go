package utils

import (
	"strings"

	"github.com/mozillazg/go-pinyin"
)

func ToPinYinMark(name string) string {
	args := pinyin.NewArgs()
	args.Style = pinyin.FirstLetter
	letters := make([]string, 0)
	for _, r := range name {
		result := pinyin.Pinyin(string(r), args)
		if len(result) > 0 && len(result[0]) > 0 {
			letters = append(letters, strings.ToUpper(result[0][0]))
		} else {
			letters = append(letters, strings.ToUpper(string(r)))
		}
	}
	return strings.Join(letters, "")
}
