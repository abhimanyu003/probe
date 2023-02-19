package runner

import (
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestHeaders_expectedValue(t *testing.T) {
	type fields struct {
		ExpectedValue any
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "int",
			fields: fields{ExpectedValue: 1},
			want:   "1",
		},
		{
			name:   "bool-true",
			fields: fields{ExpectedValue: true},
			want:   "true",
		},
		{
			name:   "bool-false",
			fields: fields{ExpectedValue: false},
			want:   "false",
		},
		{
			name:   "float",
			fields: fields{ExpectedValue: 12.62},
			want:   "12.62",
		},
		{
			name:   "negative",
			fields: fields{ExpectedValue: -12.62},
			want:   "-12.62",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Headers{
				Want: tt.fields.ExpectedValue,
			}
			assert.Equalf(t, tt.want, h.expectedValue(), "expectedValue()")
		})
	}
}

func Test_parseResponseHeaderValue(t *testing.T) {
	defer gock.Off()
	gock.New("https://example.com").
		Get("/").
		Reply(200).
		SetHeaders(map[string]string{
			"x-test-header": "header value",
		})

	req, _ := http.NewRequest("GET", "https://example.com", nil)
	res, _ := (&http.Client{}).Do(req)

	type args struct {
		resp  *http.Response
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return correct header value",
			args: args{
				resp:  res,
				input: "x-test-header",
			},
			want: "header value",
		},
		{
			name: "should run jq contains function",
			args: args{
				resp:  res,
				input: `x-test-header | contains("header")`,
			},
			want: "true",
		},
		{
			name: "should run jq length function",
			args: args{
				resp:  res,
				input: "x-test-header | length",
			},
			want: "12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equalf(t, tt.want, parseResponseHeaderValue(tt.args.resp, tt.args.input),
				"parseResponseHeaderValue(%v, %v)", tt.args.resp, tt.args.input)
		})
	}
}
