package runner

import (
	"fmt"
	assert "probe/asserts"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/gookit/goutil/cliutil"
)

type Assert struct {
	stage *Stage `json:"-"`

	Status  int       `json:"status"`
	Headers []Headers `json:"headers"`
	Body    []Body    `json:"body"`
}

func (a *Assert) assertStatusCode(t *testing.T) {
	if a.Status > 0 {
		assert.Equal(t, a.Status, a.stage.resp.StatusCode)
	}
}

func (a *Assert) assertHeader(t *testing.T) {
	for _, h := range a.Headers {
		h := h // capture range variable
		h.assert = a
		t.Run(h.Select, func(t *testing.T) {
			t.Parallel()

			want := h.expectedValue()
			got := parseResponseHeaderValue(h.assert.stage.resp.Response, h.Select)

			if len(h.expectedValue()) > 0 {
				assert.Equal(t, want, got, h.getErrorMessage())
			}

			if len(h.ExportAs) > 0 {
				a.stage.test.probe.cache.Set(h.ExportAs, got)
			}
		})
	}
}

func (a *Assert) assertBody(t *testing.T) {
	for _, b := range a.Body {
		b := b // capture range variable
		b.assert = a

		t.Run(b.Select, func(t *testing.T) {
			// Do check in Parallel
			t.Parallel()

			// Validate JSON
			if strings.EqualFold(b.Constrain, "JSON") {
				expectedJSON, err := b.getExpectedValueAsCompactJSON()
				if err != nil {
					assert.NotNil(t, err)
				}
				gotJSON, err := b.getResponseBodyWithJQAsCompactJSON()
				if err != nil {
					assert.NotNil(t, err)
				}
				assert.JSONEq(t, expectedJSON, gotJSON, b.getErrorMessage())
				if len(b.ExportAs) > 0 {
					b.assert.stage.test.probe.cache.Set(b.ExportAs, gotJSON)
				}
				return
			}

			// Validate individual fields using jq
			gotValue, err := runJQOnResponseBody(b.assert.stage.resp.Bytes(), b.Select)
			if err != nil {
				cliutil.Redln(err)
				assert.NotNil(t, err)
			}
			expectedValue := b.getExpectedValueAsCorrectType()
			if expectedValue != nil {
				assert.Equal(t, expectedValue, gotValue, b.getErrorMessage())
			}
			if len(b.ExportAs) > 0 {
				b.assert.stage.test.probe.cache.Set(b.ExportAs, gotValue)
			}
		})
	}
}

func (h *Headers) getErrorMessage() string {
	testName := "test:  " + h.assert.stage.test.Name
	stageName := "stage: " + h.assert.stage.Name
	y, _ := yaml.Marshal(h)
	assert := "assert:"
	yml := strings.ReplaceAll(string(y), "\n", "\n\t")
	return fmt.Sprintf("%s\n%s\n%s\n\t%s", testName, stageName, assert, yml)
}

func (b *Body) getErrorMessage() string {
	testName := "test:  " + b.assert.stage.test.Name
	stageName := "stage: " + b.assert.stage.Name
	y, _ := yaml.Marshal(b)
	assert := "assert:"
	yml := strings.ReplaceAll(string(y), "\n", "\n\t")
	return fmt.Sprintf("%s\n%s\n%s\n\t%s", testName, stageName, assert, yml)
}
