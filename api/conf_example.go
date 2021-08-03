// +build !prod

package api

import (
	"regexp"
)

var articleIDRegex = regexp.MustCompile(`^/index\.php/archives/\d+/$`)
var mgmtToken = "token..."
var dingdingURL = "https://oapi.dingtalk.com/robot/send?access_token=xxxx"
var refererAllowDomains = []string{"127.0.0.1", "example.me"}
var baseURL = "https://example.me"
