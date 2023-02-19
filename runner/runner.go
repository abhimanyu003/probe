package runner

import (
	"encoding/json"
	"flag"
	"fmt"
	assert "github.com/abhimanyu003/probe/asserts"
	"github.com/abhimanyu003/probe/cache"
	"github.com/abhimanyu003/probe/parser"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/fsutil/finder"
	"github.com/imroc/req/v3"
	"github.com/juju/errors"
)

type Flags struct {
	Verbose     bool
	DisableLogs bool
	FailFast    bool
	Parallel    uint
}

type Probe struct {
	directory string
	testFiles []string
	startTime time.Time
	tests     []testing.InternalTest
	cache     *cache.Cache
	client    *req.Client
	flags     Flags
}

func NewProbe(dir string, flags Flags, client *req.Client) *Probe {
	return &Probe{
		directory: dir,
		startTime: time.Now(),
		testFiles: getTestFiles(dir),
		flags:     flags,
		client:    client,
	}
}

func (p *Probe) Execute() {
	p.buildInternalTest()
	testing.Init()

	if p.flags.Verbose == true {
		flag.Lookup("test.v").Value.Set("true")
	}
	if p.flags.FailFast == true {
		flag.Lookup("test.failfast").Value.Set("true")
	}
	if p.flags.Parallel > 0 {
		flag.Lookup("test.parallel").Value.Set(fmt.Sprintf("%d", p.flags.Parallel))
	}
	tt := testing.MainStart(matchStringOnly(func(pat, str string) (bool, error) { return true, nil }),
		p.tests,
		nil,
		nil,
		nil)
	tt.Run()
}

func (p *Probe) buildInternalTest() {
	for _, file := range p.testFiles {
		test := p.NewTest(file)

		// only create log dir if logs are enabled.
		if p.flags.DisableLogs == false {
			test.createLogDir(0)
		}

		p.tests = append(p.tests, testing.InternalTest{
			Name: test.Name,
			F: func(t *testing.T) {
				t.Parallel()
				runStages(t, test)
			},
		})
	}
}

// NewTest will return instance of new Test
func (p *Probe) NewTest(file string) Test {
	test := Test{
		probe:    p,
		filePath: file,
	}
	// Having unique cache table name for each test will present variable name conflict.
	test.probe.cache = cache.NewCache(uuid.NewString())
	data, err := os.ReadFile(file)
	if err != nil {
		cliutil.Redln(errors.Annotate(err, "failed to read file "+test.filePath))
		os.Exit(1)
	}
	y, err := parser.ParseYaml(data)
	if err != nil {
		cliutil.Redln(errors.Annotate(err, "failed to parse yaml "+test.filePath))
		os.Exit(1)
	}
	if err = json.Unmarshal(y, &test); err != nil {
		cliutil.Redln(errors.Annotate(err, "failed to unmarshal json "+test.filePath))
		os.Exit(1)
	}
	return test
}

func (p *Probe) processEach(t *testing.T, parentTest Test, files []string) {
	for _, file := range files {
		test := p.NewTest(file)
		// if log dir path is already present that mean it's a child test.
		if len(parentTest.logDirPath) > 0 {
			test.logDirPath = parentTest.logDirPath
		}
		runStages(t, test)
	}
}

func getTestFiles(dir string) []string {
	var files []string

	if fsutil.IsFile(dir) {
		return []string{dir}
	}

	finder.EmptyFinder().
		AddDir(dir).
		NoDotFile().
		NoDotDir().
		ExcludeDir("logs").
		Find().
		Each(func(filePath string) {
			if strings.Contains(filePath, ".yaml") || strings.Contains(filePath, ".yml") {
				files = append(files, filePath)
			}
		})

	return files
}

func runStages(t *testing.T, data Test) {
	data.exportGlobalVariables()
	data.probe.processEach(t, data, data.BeforeAll)
	data.parseVariables()

	for _, stage := range data.Stages {
		stage := stage // capture range variable
		stage.test = &data
		stage.Assert.stage = &stage
		stage.parseVariables()

		data.probe.processEach(t, data, data.BeforeEach)
		if stage.Skip {
			t.Skip()
		}
		t.Run(stage.Name, func(t *testing.T) {
			if stage.test.probe.client == nil {
				stage.resp = sendHTTPRequest(stage, req.C())
			} else {
				stage.resp = sendHTTPRequest(stage, stage.test.probe.client)
			}
			assert.NotNil(t, stage.resp.Response)
			if stage.resp.Response != nil {
				stage.Assert.assertStatusCode(t)
				stage.Assert.assertHeader(t)
				stage.Assert.assertBody(t)
				stage.exportRunTimeVariables()
			}
		})
		data.probe.processEach(t, data, data.AfterEach)
		// Sleep after stage is done.
		time.Sleep(stage.SleepAfter * time.Millisecond)
	}
	data.probe.processEach(t, data, data.AfterAll)
}

// exportRunTimeVariables will export run time variables from header and body
//
//	export:
//	 header:
//		- select: "x-header-key"
//		  as: test
//	 body:
//	   	- select: ."user-agent"
//	      as: userAgent
func (s *Stage) exportRunTimeVariables() {
	for _, h := range s.Export.Headers {
		headerValue := parseResponseHeaderValue(s.resp.Response, h.Select)
		s.test.probe.cache.Set(h.As, headerValue)
	}
	for _, b := range s.Export.Body {
		bodyValue, _ := runJQOnResponseBody(s.resp.Bytes(), b.Select)
		s.test.probe.cache.Set(b.As, bodyValue)
	}
}
