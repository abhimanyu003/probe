package runner

import (
	"github.com/abhimanyu003/probe/cache"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/h2non/gock"
	"github.com/imroc/req/v3"
	"github.com/stretchr/testify/assert"
)

func Test_runStages_Headers_HappyPath(t *testing.T) {
	defer gock.Off()
	client := req.C()
	gock.InterceptClient(client.GetClient())
	gock.New("https://example.com").
		Times(1).
		Reply(http.StatusOK).
		SetHeaders(map[string]string{
			"x-token":       "abcd123",
			"second-header": "abcd123",
		})
	gock.New("https://example-2.com").
		Times(1).
		Reply(http.StatusBadGateway)

	probe := Probe{
		cache:  cache.NewCache("testing"),
		client: client,
	}
	data := Test{
		probe: &probe,
		Stages: []Stage{
			{
				Name: "Should select correct fields",
				Request: Request{
					Method: "GET",
					URL:    "https://example.com",
				},
				Assert: Assert{
					Headers: []Headers{
						{
							Select: "x-token",
							Want:   "abcd123",
						},
						{
							Select: "x-token | length",
							Want:   7,
						},
						{
							Select: "second-header",
							Want:   "abcd123",
						},
					},
				},
			},
			{
				Name: "Should assert StatusBadGateway",
				Request: Request{
					Method: "GET",
					URL:    "https://example-2.com",
				},
				Assert: Assert{
					Status: http.StatusBadGateway,
				},
			},
		},
	}

	t.Run("Test_runStages_Headers_HappyPath", func(t *testing.T) {
		runStages(t, data)
	})
}

func Test_runStages_Get_Request_HappyPath(t *testing.T) {
	defer gock.Off()
	client := req.C()
	gock.InterceptClient(client.GetClient())
	gock.New("https://example.com").
		Times(2).
		Reply(http.StatusOK).
		JSON(`{
  "id": 1,
  "title": "iPhone 9",
  "description": "An apple mobile which is nothing like apple",
  "price": 549,
  "discountPercentage": 12.96,
  "rating": 4.69,
  "stock": 94,
  "brand": "Apple",
  "category": "smartphones",
  "thumbnail": "https://example.com/data/products/1/thumbnail.jpg",
  "images": [
    "https://example.com/data/products/1/1.jpg",
    "https://example.com/data/products/1/2.jpg",
    "https://example.com/data/products/1/3.jpg",
    "https://example.com/data/products/1/4.jpg",
    "https://example.com/data/products/1/thumbnail.jpg"
  ]
}`)
	probe := Probe{
		cache:  cache.NewCache("testing"),
		client: client,
	}
	data := Test{
		probe: &probe,
		Stages: []Stage{
			{
				Name: "Should select correct fields",
				Request: Request{
					Method: "GET",
					URL:    "https://example.com",
				},
				Assert: Assert{
					Body: []Body{
						{
							Select: ".id",
							Want:   1,
						},
						{
							Select: ".title",
							Want:   "iPhone 9",
						},
						{
							Select: ".description | length",
							Want:   43,
						},
					},
				},
			},
			{
				Name: "Send Second Request",
				Request: Request{
					Method: "GET",
					URL:    "https://example.com",
				},
				Assert: Assert{
					Body: []Body{
						{
							Select: ".images | length",
							Want:   5,
						},
						{
							Select: `has("images")`,
							Want:   true,
						},
						{
							Select:    ".images",
							Constrain: "json",
							Want: `[
  "https://example.com/data/products/1/1.jpg",
  "https://example.com/data/products/1/2.jpg",
  "https://example.com/data/products/1/3.jpg",
  "https://example.com/data/products/1/4.jpg",
  "https://example.com/data/products/1/thumbnail.jpg"
]`,
						},
					},
				},
			},
		},
	}

	t.Run("Test_runStages_Get_Request_HappyPath", func(t *testing.T) {
		runStages(t, data)
	})
}

func Test_runStages_BasicAuth_HappyPath(t *testing.T) {
	defer gock.Off()
	client := req.C()
	gock.InterceptClient(client.GetClient())
	gock.New("https://example.com").
		Times(1).
		MatchHeader("Authorization", "Basic dXNlcm5hbWU6cGFzc3dvcmQ=").
		Reply(http.StatusOK)

	probe := Probe{
		cache:  cache.NewCache("testing"),
		client: client,
	}
	data := Test{
		probe: &probe,
		Stages: []Stage{
			{
				Name: "Should select correct fields",
				Request: Request{
					Method: "GET",
					BasicAuth: BasicAuth{
						User:     "username",
						Password: "password",
					},
					URL: "https://example.com",
				},
				Assert: Assert{
					Status: 200,
				},
			},
		},
	}

	t.Run("Test_runStages_BasicAuth_HappyPath", func(t *testing.T) {
		runStages(t, data)
	})
}

func Test_runStages_ExportVariables_HappyPath(t *testing.T) {
	defer gock.Off()
	client := req.C()
	gock.InterceptClient(client.GetClient())
	gock.New("https://example.com").
		Times(1).
		Reply(http.StatusOK).
		SetHeaders(map[string]string{
			"token":   "abcd123",
			"runTime": "run-time-1234",
		}).
		JSON(`{"name":"abhimanyu", "tool": "probe"}`)

	gock.New("https://exported.com").
		Times(2).
		MatchHeader("username", "abhimanyu").
		MatchHeader("runTime", "run-time-1234").
		MatchHeader("tool", "probe").
		MatchHeader("globalVariable", "globalVariable").
		Reply(http.StatusOK)

	probe := Probe{
		cache:  cache.NewCache("testing"),
		client: client,
	}
	data := Test{
		Variables: map[string]any{
			"globalVariable": "globalVariable",
		},
		probe: &probe,
		Stages: []Stage{
			{
				Name: "Should select correct fields",
				Request: Request{
					Method: "GET",
					URL:    "https://example.com",
				},
				Assert: Assert{
					Headers: []Headers{
						{
							Select:   "runTime",
							ExportAs: "runTime",
						},
					},
					Body: []Body{
						{
							Select:   ".tool",
							ExportAs: "tool",
						},
					},
				},
				Export: Export{
					Headers: []ExportHeaders{
						{
							Select: "token",
							As:     "token",
						},
					},
					Body: []ExportBody{
						{
							Select: ".name",
							As:     "exportedName",
						},
					},
				},
			},
			{
				Name: "exported variables should be parsed",
				Request: Request{
					Method: "GET",
					URL:    "https://exported.com",
					Headers: map[string]string{
						"username":       "${exportedName}",
						"runTime":        "${runTime}",
						"tool":           "${tool}",
						"globalVariable": "${globalVariable}",
					},
				},
				Assert: Assert{
					Status: http.StatusOK,
				},
			},
			{
				Name: "parse variables with multiple spaces",
				Request: Request{
					Method: "GET",
					URL:    "https://exported.com",
					Headers: map[string]string{
						"username":       "${ exportedName  }",
						"runTime":        "${   runTime     }",
						"tool":           "${   tool}",
						"globalVariable": "${globalVariable   }",
					},
				},
				Assert: Assert{
					Status: http.StatusOK,
				},
			},
		},
	}

	t.Run("Test_runStages_ExportVariables_HappyPath", func(t *testing.T) {
		runStages(t, data)
	})
	value, err := data.probe.cache.Get("exportedName")
	assert.Nil(t, err)
	assert.Exactly(t, value, "abhimanyu")

	token, err := data.probe.cache.Get("token")
	assert.Nil(t, err)
	assert.Exactly(t, token, "abcd123")
}

func Test_runStages_AllowInsecureRequest_HappyPath(t *testing.T) {
	defer gock.Off()
	client := req.C()
	gock.InterceptClient(client.GetClient())
	gock.New("https://example.com").
		Times(1).
		Reply(http.StatusOK)

	probe := Probe{
		cache:  cache.NewCache("testing"),
		client: client,
	}
	data := Test{
		probe: &probe,
		Stages: []Stage{
			{
				Name: "Should allow insecure request",
				Request: Request{
					Method:        "GET",
					AllowInsecure: true,
					URL:           "https://example.com",
				},
				Assert: Assert{
					Status: 200,
				},
			},
		},
	}

	t.Run("Test_runStages_AllowInsecureRequest_HappyPath", func(t *testing.T) {
		runStages(t, data)
	})
}

func Test_runStages_RequestTimeout_HappyPath(t *testing.T) {
	defer gock.Off()
	client := req.C()
	gock.InterceptClient(client.GetClient())
	gock.New("https://example.com").
		Times(1).
		Reply(http.StatusOK).
		Delay(time.Second * 2)

	probe := Probe{
		cache:  cache.NewCache("testing"),
		client: client,
	}
	data := Test{
		probe: &probe,
		Stages: []Stage{
			{
				Name: "Test_runStages_RequestTimeout_HappyPath",
				Request: Request{
					Method:        "GET",
					AllowInsecure: true,
					Timeout:       4000,
					URL:           "https://example.com",
				},
				Assert: Assert{
					Status: 200,
				},
			},
		},
	}

	t.Run("Test_runStages_RequestTimeout_HappyPath", func(t *testing.T) {
		runStages(t, data)
	})
}

func Test_runStages_Body_HappyPath(t *testing.T) {
	defer gock.Off()
	client := req.C()
	gock.InterceptClient(client.GetClient())
	gock.New("https://example.com").
		Body(strings.NewReader("123")).
		Times(1).
		Reply(http.StatusOK)

	probe := Probe{
		cache:  cache.NewCache("testing"),
		client: client,
	}
	data := Test{
		probe: &probe,
		Stages: []Stage{
			{
				Name: "Test_runStages_Body_HappyPath",
				Request: Request{
					Method: "POST",
					Body:   "123",
					URL:    "https://example.com",
				},
				Assert: Assert{
					Status: 200,
				},
			},
		},
	}

	t.Run("Test_runStages_RequestTimeout_HappyPath", func(t *testing.T) {
		runStages(t, data)
	})
}

func Test_runStages_UserAgent_HappyPath(t *testing.T) {
	defer gock.Off()
	client := req.C()
	gock.InterceptClient(client.GetClient())
	gock.New("https://example.com").
		MatchHeader("user-agent", "test-agent").
		Times(1).
		Reply(http.StatusOK)

	probe := Probe{
		cache:  cache.NewCache("testing"),
		client: client,
	}
	data := Test{
		probe: &probe,
		Stages: []Stage{
			{
				Name: "Should custom user agent",
				Request: Request{
					UserAgent: "test-agent",
					Method:    "GET",
					URL:       "https://example.com",
				},
				Assert: Assert{
					Status: 200,
				},
			},
		},
	}

	t.Run("Test_runStages_UserAgent_HappyPath", func(t *testing.T) {
		runStages(t, data)
	})
}

func Test_runStages_BearerToken_HappyPath(t *testing.T) {
	defer gock.Off()
	client := req.C()
	gock.InterceptClient(client.GetClient())
	gock.New("https://example.com").
		MatchHeader("Authorization", "Bearer 123").
		Times(1).
		Reply(http.StatusOK)

	probe := Probe{
		cache:  cache.NewCache("testing"),
		client: client,
	}
	data := Test{
		probe: &probe,
		Stages: []Stage{
			{
				Name: "Should allow custom bearer token",
				Request: Request{
					UserAgent:       "test-agent",
					BearerAuthToken: "123",
					URL:             "https://example.com",
				},
				Assert: Assert{
					Status: 200,
				},
			},
		},
	}

	t.Run("Test_runStages_BearerToken_HappyPath", func(t *testing.T) {
		runStages(t, data)
	})
}

func Test_runStages_BodyJSON_HappyPath(t *testing.T) {
	defer gock.Off()
	client := req.C()
	gock.InterceptClient(client.GetClient())
	gock.New("https://example.com").
		MatchHeader("content-type", "application/json; charset=utf-8").
		Body(strings.NewReader(`{"bodyJSON": "JSON"}`)).
		Times(1).
		Reply(http.StatusOK)

	probe := Probe{
		cache:  cache.NewCache("testing"),
		client: client,
	}
	data := Test{
		probe: &probe,
		Stages: []Stage{
			{
				Name: "Test_runStages_Body_HappyPath",
				Request: Request{
					Method:   "POST",
					BodyJSON: `{"bodyJSON": "JSON"}`,
					URL:      "https://example.com",
				},
				Assert: Assert{
					Status: 200,
				},
			},
		},
	}

	t.Run("Test_runStages_BodyJSON_HappyPath", func(t *testing.T) {
		runStages(t, data)
	})
}
