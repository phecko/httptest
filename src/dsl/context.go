package dsl

import (
	"fmt"
	"regexp"
	"strings"
)

type Context struct {
	varibles map[string]interface{}
}

func SplitString(str string) []string {
	result := []string{}
	curStr := ""
	inQuote := ""
	quoteTmpl := "'\"|`"
	regx := regexp.MustCompile("\\s")
	for _, i := range str {
		s := string(i)
		fmt.Println(s)

		if len(inQuote) == 0 {
			if regx.MatchString(s) {
				if len(curStr) > 0 {
					result = append(result, curStr)
					curStr = ""
				}
				continue
			}
			if strings.Index(quoteTmpl, s) > -1 {
				inQuote = s
				continue
			}
			curStr += s
		} else {
			if inQuote == s {
				result = append(result, curStr)
				curStr = ""
			}
			curStr += s
		}
		fmt.Println(curStr)
	}
	if len(curStr) > 0 {
		result = append(result, curStr)
	}
	return result
}
