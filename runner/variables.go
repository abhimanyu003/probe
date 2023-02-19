package runner

import (
	"github.com/abhimanyu003/probe/cache"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cast"
)

func parseVariables(input string, cache *cache.Cache) string {
	pattern := regexp.MustCompile(`\$\{.*[a-zA-Z0-9].*\}`)
	matches := pattern.FindAllString(input, -1)
	for _, match := range matches {
		var finalValue string

		str := strings.TrimSpace(match)
		str = strings.ReplaceAll(str, "${", "")
		str = strings.ReplaceAll(str, "}", "")
		str = strings.ReplaceAll(str, "\t", "")
		str = strings.TrimSpace(str)
		if strings.Contains(str, "env:") || strings.Contains(str, "ENV:") {
			str = strings.ReplaceAll(str, "env:", "")
			str = strings.ReplaceAll(str, "ENV:", "")
			finalValue = os.Getenv(str)
			if finalValue == "" {
				finalValue = match
			}
		} else {
			cacheValue, err := cache.Get(str)
			if err != nil {
				finalValue = match
			} else {
				finalValue = cast.ToString(cacheValue)
			}
		}
		input = strings.ReplaceAll(input, match, finalValue)
	}

	return input
}
