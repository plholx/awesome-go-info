package api

import (
	"strings"
	"testing"
)

func TestReg(t *testing.T) {
	s := `* [go-openapi](https://github.com/go-openapi) - Collection of packages to parse and utilize open-api schemas.`
	//s = `* [go-commons-pool](https://github.com/jolestar/go-commons-pool) - Generic object pool for Golang.`
	if reLinkWithDescription.MatchString(s){
		subMatchs := reLinkWithDescription.FindStringSubmatch(s)
		t.Log(strings.Join(subMatchs, "#"))
		url := subMatchs[3]
		t.Log(url)
		if !reGitHubURL.MatchString(url) {
			t.Log("url匹配失败")
		}
		subMatchs = reGitHubURL.FindStringSubmatch(url)
		t.Log(strings.Join(subMatchs, "%"))
	} else {
		t.Error("匹配失败")
	}
}