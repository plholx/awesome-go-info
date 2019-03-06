package api

import (
	"strings"
	"testing"
)

func TestReg(t *testing.T) {
	s := `* [go-openapi](https://github.com/go-openapi) - Collection of packages to parse and utilize open-api schemas.`
	//s = `* [go-commons-pool](https://github.com/jolestar/go-commons-pool) - Generic object pool for Golang.`
	s = `* [Go Projects](https://github.com/golang/go/wiki/Projects) - List of projects on the Go community wiki.`
	s = `* [Go Projects](https://github.com/golang/go) - List of projects on the Go community wiki.`
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
		t.Log("subMatchs长度", len(subMatchs))
		if len(subMatchs[3]) < 1 {
			t.Log("完整仓库地址")
		}
	} else {
		t.Error("匹配失败")
	}
}