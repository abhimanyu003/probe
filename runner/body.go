package runner

import (
	"bytes"
	"encoding/json"

	"github.com/abhimanyu003/probe/cast"
	"github.com/abhimanyu003/probe/jq"

	"github.com/ohler55/ojg/oj"
	spf "github.com/spf13/cast"
)

type Body struct {
	assert *Assert `json:"-"`

	Select    string `json:"select,omitempty" yaml:"select"`
	Constrain string `json:"constrain,omitempty" yaml:"constrain,omitempty"`
	Want      any    `json:"want,omitempty" yaml:"want,omitempty"`
	ExportAs  string `json:"exportAs,omitempty" yaml:"exportAs,omitempty"`
}

func (b *Body) getExpectedValueAsCorrectType() any {
	return cast.ToCorrectType(b.Want)
}

func (b *Body) getExpectedValueAsCompactJSON() (string, error) {
	buffer := new(bytes.Buffer)
	err := json.Compact(buffer, []byte(b.Want.(string)))
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func (b *Body) getResponseBodyAsCompactJSON() (string, error) {
	buffer := new(bytes.Buffer)
	err := json.Compact(buffer, b.assert.stage.resp.Bytes())
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

// getResponseBodyWithJQAsCompactJSON is only for string values.
func (b *Body) getResponseBodyWithJQAsCompactJSON() (string, error) {
	jqResult, err := runJQOnResponseBody(b.assert.stage.resp.Bytes(), b.Select)
	if err != nil {
		return "", err
	}

	switch jqResult.(type) {
	case int, float64, float32, bool, string:
		return spf.ToString(jqResult), nil
	}

	src, err := oj.Marshal(jqResult)
	if err != nil {
		return "", err
	}
	buffer := new(bytes.Buffer)
	err = json.Compact(buffer, src)

	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func runJQOnResponseBody(body []byte, query string) (any, error) {
	obj, err := oj.Parse(body)
	if err != nil {
		return nil, err
	}
	return jq.RunJq(query, obj)
}
