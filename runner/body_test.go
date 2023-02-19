package runner

import (
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/imroc/req/v3"
	"github.com/stretchr/testify/assert"
)

func TestBody_getExpectedValueAsCorrectType(t *testing.T) {
	tests := []struct {
		name   string
		fields any
		want   any
	}{
		{
			name:   "int",
			fields: 1,
			want:   1,
		},
		{
			name:   "string",
			fields: "1",
			want:   "1",
		},
		{
			name:   "bool",
			fields: true,
			want:   true,
		},
		{
			name:   "float",
			fields: 12.12,
			want:   12.12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Body{
				Want: tt.fields,
			}
			assert.Exactly(t, tt.want, b.getExpectedValueAsCorrectType())
		})
	}
}

func TestBody_getExpectedValueAsCompactJSON(t *testing.T) {
	type fields struct {
		ExpectedValue any
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name:   "correct json",
			fields: fields{ExpectedValue: `{"name": "test"    }`},
			want:   `{"name":"test"}`,
		},
		{
			name:   "correct nested json",
			fields: fields{ExpectedValue: `{"name": "test", 		"array": [1,	2,		3]    }`},
			want:   `{"name":"test","array":[1,2,3]}`,
		},
		{
			name:    "incorrect json",
			fields:  fields{ExpectedValue: `{"name": "test   }`},
			want:    ``,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Body{
				Want: tt.fields.ExpectedValue,
			}
			got, err := b.getExpectedValueAsCompactJSON()
			if tt.wantErr {
				assert.NotNil(t, err)
			}
			assert.Equalf(t, tt.want, got, "getExpectedValueAsCompactJSON()")
		})
	}
}

func TestBody_runJQ(t *testing.T) {
	type fields struct {
		body   []byte
		Select string
	}
	tests := []struct {
		name    string
		fields  fields
		want    any
		wantErr bool
	}{
		{
			name: "should select field",
			fields: fields{
				body:   []byte(`{"name": "testUser"}`),
				Select: ".name",
			},
			want:    `testUser`,
			wantErr: false,
		},
		{
			name: "should select nested field",
			fields: fields{
				body: []byte(`{
  "name": "test",

  "obj": {
			    "insideObject": "1",
      "array": [1,    2,    3,    4]
  
},
  "id": 1
}`),
				Select: ".obj.insideObject",
			},
			want:    `1`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Body{
				Select: tt.fields.Select,
			}
			got, err := runJQOnResponseBody(tt.fields.body, b.Select)
			if tt.wantErr {
				assert.NotNil(t, err)
			}
			assert.Equalf(t, tt.want, got, "runJQOnResponseBody()")
		})
	}
}

func TestBody_getResponseBodyAsCompactJSON(t *testing.T) {
	client := req.C()
	defer gock.Off()
	gock.InterceptClient(client.GetClient())
	gock.New("https://example.com").
		Persist().
		Reply(http.StatusOK).
		JSON(`{	"name": 	"test"	, "id": 	1}`).
		SetHeaders(map[string]string{
			"x-token":       "abcd123",
			"second-header": "abcd123",
		})
	type fields struct {
		asserts *Assert
	}
	tests := []struct {
		name        string
		mockRequest func()
		fields      fields
		want        string
		wantErr     bool
	}{
		{
			name: "should return correct compact json",
			fields: fields{
				asserts: &Assert{
					stage: &Stage{
						Request: Request{
							Method: "GET",
							URL:    "https://example.com",
						},
					},
				},
			},
			want:    `{"name":"test","id":1}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Body{
				assert: tt.fields.asserts,
			}
			r, _ := client.R().Send("GET", tt.fields.asserts.stage.Request.URL)
			b.assert.stage.resp = r
			got, err := b.getResponseBodyAsCompactJSON()
			if tt.wantErr {
				assert.NotNil(t, err)
			}
			assert.Equalf(t, tt.want, got, "getResponseBodyAsCompactJSON()")
		})
	}
}

func TestBody_getResponseBodyJQAsCompactJSON(t *testing.T) {
	client := req.C()
	defer gock.Off()
	gock.InterceptClient(client.GetClient())
	gock.New("https://example.com").
		Persist().
		Reply(http.StatusOK).
		JSON(`{
  "name": "test",

  "obj": {
			    "insideObject": "1",
      "array": [1,    2,    3,    4]
  
},
  "id": 1
}`).SetHeaders(map[string]string{
		"x-token":       "abcd123",
		"second-header": "abcd123",
	})

	build := func() *Assert {
		r, _ := client.R().Send("GET", "https://example.com")
		return &Assert{
			stage: &Stage{
				resp: r,
			},
		}
	}

	type fields struct {
		asserts *Assert
		Select  string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should return object",
			fields: fields{
				asserts: build(),
				Select:  `.`,
			},
			want:    `{"name":"test","obj":{"array":[1,2,3,4],"insideObject":"1"},"id":1}`,
			wantErr: false,
		},
		{
			name: "should return array",
			fields: fields{
				asserts: build(),
				Select:  `.obj.array`,
			},
			want:    `[1,2,3,4]`,
			wantErr: false,
		},
		{
			// We are only worried about sting values here
			// As if given build() is exact int
			// It should not be used as `JSON` assert
			name: "should return array index value",
			fields: fields{
				asserts: build(),
				Select:  `.obj.array[0]`,
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "should return string value",
			fields: fields{
				asserts: build(),
				Select:  `.obj.insideObject`,
			},
			want:    "1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			b := &Body{
				assert: build(),
				Select: tt.fields.Select,
			}
			got, err := b.getResponseBodyWithJQAsCompactJSON()
			if tt.wantErr {
				assert.NotNil(t, err)
			}
			assert.JSONEq(t, tt.want, got, "getResponseBodyAsCompactJSON()")
			t.Cleanup(func() {
				build()
			})
		})
	}
}
