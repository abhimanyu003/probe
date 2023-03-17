package runner

import (
	"fmt"
	"time"

	"github.com/gookit/goutil/fsutil"
	"gopkg.in/yaml.v3"
)

type Test struct {
	probe      *Probe `json:"-"`
	logDirPath string `json:"-"`
	filePath   string `json:"-"`

	Stages     []Stage        `json:"stages"`
	Skip       bool           `json:"skip"`
	Key        string         `json:"key"`
	Name       string         `json:"name"`
	BeforeAll  []string       `json:"beforeAll"`
	AfterAll   []string       `json:"afterAll"`
	BeforeEach []string       `json:"beforeEach"`
	AfterEach  []string       `json:"afterEach"`
	Variables  map[string]any `json:"variables"`
	Request    Request        `json:"request"`
}

func (t *Test) createLogDir(prefix int) {
	var logDirPath string
	if prefix == 0 {
		logDirPath = fmt.Sprintf("logs/%s/%s", t.probe.startTime.Format(time.RFC3339), t.Name)
	} else {
		logDirPath = fmt.Sprintf("logs/%s/%s-%d", t.probe.startTime.Format(time.RFC3339), t.Name, prefix)
	}
	if !fsutil.IsDir(logDirPath) {
		fsutil.Mkdir(logDirPath, 0700)
	} else {
		t.createLogDir(prefix + 1)
	}
	t.logDirPath = logDirPath
}

type BasicAuth struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type Retry struct {
	After   time.Duration `json:"after"`
	Count   int           `json:"count"`
	RetryOn []int         `json:"retryOn"`
}

type Certificates struct {
	CertFile string `json:"certFile"`
	KeyFile  string `json:"keyFile"`
}

type Files struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Request struct {
	AllowInsecure   bool              `json:"allowInsecure"`
	BasicAuth       BasicAuth         `json:"basicAuth"`
	Body            string            `json:"body"`
	BodyJSON        string            `json:"bodyJson"`
	Certificates    Certificates      `json:"certificates"`
	FormData        map[string]any    `json:"formData"`
	Headers         map[string]string `json:"headers"`
	Method          string            `json:"method"`
	QueryParams     map[string]string `json:"queryParams"`
	Retry           Retry             `json:"retry"`
	Files           []Files           `json:"files"`
	SleepAfter      time.Duration     `json:"sleepAfter"`
	Timeout         time.Duration     `json:"timeout"`
	Times           time.Duration     `json:"times"`
	URL             string            `json:"url"`
	UserAgent       string            `json:"userAgent"`
	BearerAuthToken string            `json:"bearerAuthToken"`
}

func (t *Test) exportGlobalVariables() {
	// Set global state variables
	for key, value := range t.Variables {
		t.probe.cache.Set(key, value)
	}
}

func (t *Test) parseVariables() {
	y, _ := yaml.Marshal(t)
	outStr := parseVariables(string(y), t.probe.cache)
	yaml.Unmarshal([]byte(outStr), &t)
}
