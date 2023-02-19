package runner

import (
	"fmt"
	"time"

	"github.com/gookit/goutil/fsutil"
	"github.com/imroc/req/v3"
	"gopkg.in/yaml.v3"
)

type Stage struct {
	test *Test         `json:"-"`
	resp *req.Response `json:"-"`

	Name       string        `json:"name"`
	Key        string        `json:"key"`
	Skip       bool          `json:"skip"`
	Request    Request       `json:"request"`
	Assert     Assert        `json:"assert"`
	Export     Export        `json:"export"`
	SleepAfter time.Duration `json:"sleepAfter"`
}

type ExportHeaders struct {
	Select string `json:"select"`
	As     string `json:"as"`
}
type ExportBody struct {
	Select string `json:"select"`
	As     string `json:"as"`
}
type Export struct {
	Headers []ExportHeaders `json:"headers"`
	Body    []ExportBody    `json:"body"`
}

func (s *Stage) getLogFileName(prefix uint) string {
	var logFileName string
	if prefix == 0 {
		logFileName = fmt.Sprintf("%s/%s.log", s.test.logDirPath, s.Name)
	} else {
		logFileName = fmt.Sprintf("%s/%s-%d.log", s.test.logDirPath, s.Name, prefix)
	}
	if fsutil.IsFile(logFileName) {
		return s.getLogFileName(prefix + 1)
	}

	return logFileName
}

func (s *Stage) parseVariables() {
	y, _ := yaml.Marshal(s)
	outStr := parseVariables(string(y), s.test.probe.cache)
	yaml.Unmarshal([]byte(outStr), &s)
}
