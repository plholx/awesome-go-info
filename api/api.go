package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	sourceFileURL = "https://raw.githubusercontent.com/avelino/awesome-go/master/README.md"
	githubReposAPI = "https://api.github.com/repos/:owner/:repo?access_token=OAUTH-TOKEN"
	githubRateLimitAPI = "https://api.github.com/rate_limit?access_token=OAUTH-TOKEN"
	githubDomain = "https://github.com"
)

var (
	reCategoryLi          = regexp.MustCompile(`(\s*)- \[(.*)\]\(#(.*)\)`)
	reCategory            = regexp.MustCompile(`#+ (.+)`)
	reCategoryDescription = regexp.MustCompile(`(\s*)\*(.*)\*$`)
	reContainsLink        = regexp.MustCompile(`\* \[.*\]\(.*\)`)
	reOnlyLink            = regexp.MustCompile(`(\s*)\* \[(.*)\]\((.+)\)$`)
	reLinkWithDescription = regexp.MustCompile(`(\s*)\* \[(.*?)\]\((.+?)\) - (\S.*[\.\!])`)
	reLittleCategory      = regexp.MustCompile(`(\s*)\* ([a-zA-Z\s]*)$`)
	reGitHubURL           = regexp.MustCompile(`https://github.com/(.+?)/([a-zA-Z0-9_\-]+).*$`)
)

//下载awesome-go中的README.md文件
func DownloadReadmeFile()  {
	res, err := http.Get(sourceFileURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	bytes, e := ioutil.ReadAll(res.Body)
	if e != nil {
		log.Fatal(err)
	}
	outputFile, outputError := os.OpenFile("data/readmeFiles/README.md", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if outputError != nil {
		log.Fatal(outputError)
	}
	defer outputFile.Close()
	writer := bufio.NewWriter(outputFile)
	writer.Write(bytes)
	writer.Flush()
}

//解析awesome-go中的README.md文件
func parseReadmeFile(accessToken string)  {
	input, err := ioutil.ReadFile("data/readmeFiles/README.md")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")
	categoryIds := make(map[int]int64)
	var tmpCategoryId int64 = 0
	var linkCategoryId int64 = 0
	var count int64 = 0
	name := ""
	for _, line := range lines {
		if reCategoryLi.MatchString(line) {//分类目录
			subMatchs := reCategoryLi.FindStringSubmatch(line)
			name = subMatchs[2]
			spaceCount := len(subMatchs[1])

			goRepo, e := GetGoRepo(name, false, true)
			if e != nil {
				goRepo = &GoRepo{
					ParentId: categoryIds[spaceCount-8],
					Repo: false,
					Category: true,
					Name: name,
				}
				SaveGoRepo(goRepo)
			} else {
				UpdateGoRepoParentId(categoryIds[spaceCount-8], name, false, true)
			}
			categoryIds[spaceCount-4] = goRepo.Id
		} else if reCategory.MatchString(line) {//遇到分类
			subMatchs := reCategory.FindStringSubmatch(line)
			name = subMatchs[1]
			goRepo, e := GetGoRepo(name, false, true)
			if e != nil {
				log.Printf("分类%s不存在", name)
				continue
			}
			tmpCategoryId = goRepo.Id
			linkCategoryId = tmpCategoryId
		} else if reCategoryDescription.MatchString(line) {//分类描述
			subMatchs := reCategoryDescription.FindStringSubmatch(line)
			description := subMatchs[2]
			UpdateGoRepoDescription(description, tmpCategoryId)
		} else if reLittleCategory.MatchString(line) {//小分类
			subMatchs := reLittleCategory.FindStringSubmatch(line)
			name = subMatchs[2]
			goRepo, e := GetGoRepo(name, false, true)
			if e != nil {
				goRepo = &GoRepo{
					ParentId: tmpCategoryId,
					Repo: false,
					Category: true,
					Name: name,
				}
				SaveGoRepo(goRepo)
			} else {
				UpdateGoRepoParentId(tmpCategoryId, name, false, true)
			}
			linkCategoryId = goRepo.Id
		} else if reContainsLink.MatchString(line) && strings.Contains(line, githubDomain) {//含有链接,且为GitHub仓库
			githubRepoLink := ""
			repoDescription := ""
			if reOnlyLink.MatchString(line){
				subMatchs := reOnlyLink.FindStringSubmatch(line)
				githubRepoLink = subMatchs[3]
			} else if reLinkWithDescription.MatchString(line) {
				subMatchs := reLinkWithDescription.FindStringSubmatch(line)
				githubRepoLink = subMatchs[3]
				repoDescription = subMatchs[4]
			}
			if githubRepoLink != "" {
				subMatchs := reGitHubURL.FindStringSubmatch(githubRepoLink)
				repoOwner, repoName := subMatchs[1], subMatchs[2]
				name = repoOwner + "/" + repoName
				ok, err := GitHubAPIReqControl(accessToken)
				if err != nil {
					log.Println(err)
					continue
				}
				if !ok {
					log.Println("API请求速率限制")
					continue
				}
				count++
				log.Println(count, " 开始请求仓库信息", name)
				repo, err := GetRepoInfo(repoOwner, repoName, accessToken)
				if err != nil {
					log.Println(err)
					continue
				}
				repo.ParentId = linkCategoryId
				goRepo, e := GetGoRepo(name, true, false)
				if e != nil {
					if repoDescription != "" {
						repo.Description = repoDescription
					}
					SaveGoRepo(repo)
				} else {
					UpdateGoRepoGithubInfo(repo, goRepo.Id)
				}
			}
		}
	}
}


//获取github仓库信息
func GetRepoInfo(repoOwner, repoName, accessToken string) (repo *GoRepo, err error) {
	repoFullName := repoOwner + "/" + repoName
	apiURL := strings.Replace(githubReposAPI, ":owner/:repo", repoFullName, -1)
	apiURL = strings.Replace(apiURL, "OAUTH-TOKEN", accessToken, -1)
	res, err := http.Get(apiURL)
	if err != nil {
		return
	}
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 
	}
	var v interface{}
	err = json.Unmarshal(bytes, &v)
	if err != nil {
		return
	}
	if repoMap, ok := v.(map[string]interface{}); ok {
		if fullName, ok := repoMap["full_name"].(string); !ok || fullName == "" {
			return repo, fmt.Errorf("从GitHub获取%s仓库信息失败", repoFullName)
		}
		repo = &GoRepo{
			RepoName:repoMap["name"].(string), 
			RepoFullName:repoMap["full_name"].(string), 
			RepoHtmlURL:repoMap["html_url"].(string), 
			RepoForksCount:int64(repoMap["forks_count"].(float64)), 
			RepoStargazersCount:int64(repoMap["stargazers_count"].(float64)), 
			RepoSubscribersCount:int64(repoMap["subscribers_count"].(float64)), 
			RepoOpenIssuesCount:int64(repoMap["open_issues_count"].(float64)), 
			Repo:true, 
			Category:false, 
			Name:repoMap["full_name"].(string), 
		}
		//可能不存在的一些信息，需要进行判断
		if homepage, ok := repoMap["homepage"].(string); ok {
			repo.RepoHomepage = homepage
			repo.Homepage = homepage
		}
		if description, ok := repoMap["description"].(string); ok {
			repo.RepoDescription = description
			repo.Description = description
		}
		//作者信息
		if ownerMap, ok := repoMap["owner"].(map[string]interface{}); ok {
			repo.RepoOwner = ownerMap["login"].(string)
		}
		//证书信息
		if licenseMap, ok := repoMap["license"].(map[string]interface{}); ok {
			if licenseName, ok := licenseMap["name"].(string); ok {
				repo.RepoLicenseName = licenseName
			}
			if licenseSpdxId, ok := licenseMap["spdx_id"].(string); ok {
				repo.RepoLicenseSpdxId = licenseSpdxId
			}
			if licenseURL, ok := licenseMap["url"].(string); ok {
				repo.RepoLicenseURL = licenseURL
			}
		}
		//日期时间信息
		if createAtStr, ok := repoMap["created_at"].(string); ok {
			createAt, err := time.Parse(time.RFC3339, createAtStr)
			if err == nil {
				repo.RepoCreatedAt = createAt
			}
		}
		if pushedAtStr, ok := repoMap["pushed_at"].(string); ok {
			pushedAt, err := time.Parse(time.RFC3339, pushedAtStr)
			if err == nil {
				repo.RepoPushedAt = pushedAt
			}
		}
		//size可能带小数
		if sizeFloat64, ok := repoMap["size"].(float64); ok {
			size := math.RoundToEven(sizeFloat64);
			repo.RepoSize = int64(size)
		}
	}
	return
}
//控制是否可以进行GitHub的API调用
func GitHubAPIReqControl(accessToken string) (ok bool, err error) {
	//通过访问rate_limit API，校验access_token是否有效，及通过剩余次数控制API的访问时机
	apiURL := strings.Replace(githubRateLimitAPI, "OAUTH-TOKEN", accessToken, -1)
	res, err := http.Get(apiURL)
	if err != nil {
		return
	}
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	var v interface{}
	err = json.Unmarshal(bytes, &v)
	if err != nil {
		return
	}
	if rateLimitMap, ok := v.(map[string]interface{}); ok {
		//resources信息
		if resourcesMap, ok := rateLimitMap["resources"].(map[string]interface{}); ok {
			if coreMap, ok := resourcesMap["core"].(map[string]interface{}); ok {
				var remaining float64
				var limit float64
				if remainingFloat64, ok := coreMap["remaining"].(float64); ok {
					remaining = math.RoundToEven(remainingFloat64);
				}
				if limitFloat64, ok := coreMap["limit"].(float64); ok {
					limit = math.RoundToEven(limitFloat64);
				}
				if remaining > 0 && limit > 0 {
					time.Sleep(3600*1e9/time.Duration(limit))
					return true, nil
				}
			}
		}
	}
	return
}