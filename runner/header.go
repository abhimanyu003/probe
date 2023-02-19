package runner

import (
	"fmt"
	"net/http"
	"probe/jq"
	"strings"

	cast2 "github.com/spf13/cast"
)

type Headers struct {
	assert *Assert `json:"-"`

	Select    string `json:"select,omitempty" yaml:"select"`
	Constrain string `json:"constrain,omitempty" yaml:"constrain,omitempty"`
	Want      any    `json:"want,omitempty" yaml:"want,omitempty"`
	ExportAs  string `json:"exportAs,omitempty" yaml:"exportAs,omitempty"`
}

func parseResponseHeaderValue(resp *http.Response, input string) string {
	// This block allows us to run jq query functions
	// Example: length, isnormal and other
	if strings.Contains(input, "|") {
		subStr := strings.Split(input, "|")

		headerValue := resp.Header.Get(strings.TrimSpace(subStr[0]))
		query := fmt.Sprintf(".header | %s", strings.TrimSpace(subStr[1]))

		input := map[string]interface{}{
			"header": headerValue,
		}
		value, _ := jq.RunJq(query, input)

		return cast2.ToString(value)
	}
	return resp.Header.Get(input)
}

func (h *Headers) expectedValue() string {
	return cast2.ToString(h.Want)
}
