package runner

import (
	"time"

	"github.com/imroc/req/v3"
)

func sendHTTPRequest(stage Stage, c *req.Client) *req.Response {
	client := c

	// only enable logger if logs are enabled.
	if stage.test.probe.flags.DisableLogs == false {
		client = c.EnableTraceAll().EnableDumpAllToFile(stage.getLogFileName(0))
	}

	client.SetUserAgent("probe ( https://github.com/abhimanyu003/probe )")
	if len(stage.Request.UserAgent) > 0 {
		client.SetUserAgent(stage.Request.UserAgent)
	}
	if stage.Request.Timeout > 0 {
		client.SetTimeout(stage.Request.Timeout * time.Millisecond)
	}
	if len(stage.Request.Certificates.CertFile) > 0 || len(stage.Request.Certificates.KeyFile) > 0 {
		client.SetCertFromFile(stage.Request.Certificates.CertFile, stage.Request.Certificates.KeyFile)
	}
	if stage.Request.AllowInsecure {
		client.EnableInsecureSkipVerify()
	}
	r := client.R()
	r.SetHeaders(stage.Request.Headers)
	r.SetQueryParams(stage.Request.QueryParams)
	r.SetFormDataAnyType(stage.Request.FormData)

	for _, file := range stage.Request.Files {
		r.SetFile(file.Name, file.Path)
	}
	if len(stage.Request.Body) > 0 {
		r.SetBody(stage.Request.Body)
	}
	if len(stage.Request.BodyJson) > 0 {
		r.SetBodyJsonString(stage.Request.BodyJson)
	}
	if len(stage.Request.BasicAuth.User) > 0 {
		r.SetBasicAuth(stage.Request.BasicAuth.User, stage.Request.BasicAuth.Password)
	}
	if len(stage.Request.BearerAuthToken) > 0 {
		r.SetBearerAuthToken(stage.Request.BearerAuthToken)
	}
	if stage.Request.Retry.Count > 0 {
		r.SetRetryCount(stage.Request.Retry.Count).
			SetRetryFixedInterval(stage.Request.Retry.After).
			AddRetryCondition(func(resp *req.Response, err error) bool {
				return err != nil || resp.StatusCode >= 500
			})
	}

	var resp *req.Response
	resp, _ = r.Send(stage.Request.Method, stage.Request.URL)
	time.Sleep(stage.Request.SleepAfter * time.Millisecond)
	for i := 2; i <= int(stage.Request.Times); i++ {
		resp, _ = r.Send(stage.Request.Method, stage.Request.URL)
		time.Sleep(time.Duration(int(stage.Request.SleepAfter)) * time.Millisecond)
	}

	return resp
}
