package helper

import (
	"strings"
)

func GetInnerSubstring(str string, prefix string, suffix string) string {
	var beginIndex, endIndex int
	beginIndex = strings.Index(str, prefix)
	if beginIndex == -1 {
		beginIndex = 0
		endIndex = 0
	} else if len(prefix) == 0 {
		beginIndex = 0
		endIndex = strings.Index(str, suffix)
		if endIndex == -1 || len(suffix) == 0 {
			endIndex = len(str)
		}
	} else {
		beginIndex += len(prefix)
		endIndex = strings.Index(str[beginIndex:], suffix)
		if endIndex == -1 {
			if strings.Index(str, suffix) < beginIndex {
				endIndex = beginIndex
			} else {
				endIndex = len(str)
			}
		} else {
			if len(suffix) == 0 {
				endIndex = len(str)
			} else {
				endIndex += beginIndex
			}
		}
	}

	return str[beginIndex:endIndex]
}
func GetInnerSubstringReverse(str string, suffix string, prefix string) string {
	var beginIndex, endIndex int
	endIndex = strings.Index(str, suffix)
	if endIndex == -1 {
		beginIndex = 0
		endIndex = len(str)
	} else if len(suffix) == 0 {
		endIndex = len(str)
	} else {
		beginIndex = strings.LastIndex(str[:endIndex], prefix)
		beginIndex += len(prefix)
	}
	return str[beginIndex:endIndex]
}
