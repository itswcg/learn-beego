package models

import (
	"path/filepath"
	"strings"
)

var (
	NotPV []string = []string{"css", "js", "class", "gif", "png"}
)

const big = 0xFFFFFF

func LogPV(urls string) bool {
	ext := filepath.Ext(urls)
	if ext == "" {
		return true
	}
	for _, v := range NotPV {
		if v == strings.ToLower(ext) {
			return false
		}
	}
	return true
}
