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
	"os/exec"
	"regexp"
	"strings"
	"text/template"
	"time"
)

const (
	sourceFileURL = "https://raw.githubusercontent.com/avelino/awesome-go/master/README.md"
	githubReposAPI = "https://api.github.com/repos/:owner/:repo?access_token=OAUTH-TOKEN"
	githubRateLimitAPI = "https://api.github.com/rate_limit?access_token=OAUTH-TOKEN"
	githubDomain = "https://github.com"
	README_PATH = "data/readmeFiles/README-FORMAT_DATE.md"

	README_TEMPLATES_PATH = "tmpl/tmpl.md"
	README_OUTEPUT_PATH = "README.md"

	git = `git`
	add = `add`
	spot = `.`
	commit = `commit`
	m = `-m`
	msg = `测试git命令`
	push = `push`
)

var (
	reCategoryLi          = regexp.MustCompile(`(\s*)- \[(.*)\]\(#(.*)\)`)
	reCategory            = regexp.MustCompile(`#+ (.+)`)
	reCategoryDescription = regexp.MustCompile(`(\s*)\*(.*)\*$`)
	reContainsLink        = regexp.MustCompile(`\* \[.*\]\(.*\)`)
	reOnlyLink            = regexp.MustCompile(`(\s*)\* \[(.*)\]\((.+)\)$`)
	reLinkWithDescription = regexp.MustCompile(`(\s*)\* \[(.*?)\]\((.+?)\) - (\S.*[\.\!])`)
	reLittleCategory      = regexp.MustCompile(`(\s*)\* ([a-zA-Z\s]*)$`)
	reGitHubURL           = regexp.MustCompile(`https://github.com/(.+?)/([a-zA-Z0-9_\-\.]+)(.*)$`)

	reSpecialCharacters   = regexp.MustCompile(`[^a-zA-Z0-9_\-\.]+`)
)

//下载awesome-go中的README.md文件
func DownloadReadmeFile() (readmeFilePath string) {
	readmeFilePath = strings.Replace(README_PATH, "FORMAT_DATE", time.Now().Format("20060102"), -1)
	_, err := os.Stat(readmeFilePath)
	//当天的README.md文件已下载就不再重复下载
	if err == nil || os.IsExist(err) {
		log.Println(readmeFilePath, "文件已存在")
		return
	}
	res, err := http.Get(sourceFileURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	bytes, e := ioutil.ReadAll(res.Body)
	if e != nil {
		log.Fatal(err)
	}
	outputFile, outputError := os.OpenFile(readmeFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if outputError != nil {
		log.Fatal(outputError)
	}
	defer outputFile.Close()
	writer := bufio.NewWriter(outputFile)
	writer.Write(bytes)
	writer.Flush()
	return 
}

//ParseReadmeFile 解析awesome-go中的README.md文件,并存入数据库中
func ParseReadmeFile(accessToken string, readmeFilePath string)  {
	input, err := ioutil.ReadFile(readmeFilePath)
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
			categoryHtmlId := subMatchs[3]
			spaceCount := len(subMatchs[1])

			agi, e := GetAGIByCategoryHtmlId(categoryHtmlId)
			if e != nil {
				agi = &AGI{
					ParentId: categoryIds[spaceCount-8],
					Repo: false,
					Category: true,
					Name: name,
					CategoryHtmlId: categoryHtmlId,
				}
				SaveAGI(agi)
			} else {
				ModifyAGIParentIdByCategoryHtmlId(categoryIds[spaceCount-8], categoryHtmlId)
			}
			categoryIds[spaceCount-4] = agi.Id
		} else if reCategory.MatchString(line) {//遇到分类
			subMatchs := reCategory.FindStringSubmatch(line)
			name = subMatchs[1]
			categoryHtmlId := getCategoryHtmlId(name)

			agi, e := GetAGIByCategoryHtmlId(categoryHtmlId)
			if e != nil {
				log.Printf("分类%s不存在", name)
				continue
			}
			tmpCategoryId = agi.Id
			linkCategoryId = tmpCategoryId
		} else if reCategoryDescription.MatchString(line) {//分类描述
			subMatchs := reCategoryDescription.FindStringSubmatch(line)
			description := subMatchs[2]
			UpdateAGIDescription(description, tmpCategoryId)
		} else if reLittleCategory.MatchString(line) {//小分类
			subMatchs := reLittleCategory.FindStringSubmatch(line)
			name = subMatchs[2]
			categoryHtmlId := getCategoryHtmlId(name)
			agi, e := GetAGIByCategoryHtmlId(categoryHtmlId)
			if e != nil {
				agi = &AGI{
					ParentId: tmpCategoryId,
					Repo: false,
					Category: true,
					Name: name,
					CategoryHtmlId: categoryHtmlId,
				}
				SaveAGI(agi)
			} else {
				ModifyAGIParentIdByCategoryHtmlId(tmpCategoryId, categoryHtmlId)
			}
			linkCategoryId = agi.Id
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
			if githubRepoLink != "" && reGitHubURL.MatchString(githubRepoLink) {
				subMatchs := reGitHubURL.FindStringSubmatch(githubRepoLink)
				extraStr := subMatchs[3]
				if len(extraStr) > 0 && strings.HasPrefix(extraStr, "/") {//非仓库地址
					continue
				}

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
				tmpAGI, err := GetRepoInfo(repoOwner, repoName, accessToken)
				if err != nil {
					log.Println(err)
					continue
				}
				tmpAGI.ParentId = linkCategoryId
				if repoDescription != "" {
					tmpAGI.Description = repoDescription
				}

				agi, e := GetAGI(name, true, false)
				if e != nil {
					SaveAGI(tmpAGI)
				} else {
					tmpAGI.Id = agi.Id
					UpdateAGIGithubInfo(tmpAGI)
				}

				SaveGRR(tmpAGI)
				log.Print("请求完成。")
			}
		}
	}
	log.Printf("解析%s文件完毕", readmeFilePath)
}


//获取github仓库信息
func GetRepoInfo(repoOwner, repoName, accessToken string) (agi *AGI, err error) {
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
			return agi, fmt.Errorf("从GitHub获取%s仓库信息失败", repoFullName)
		}
		agi = &AGI{
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
			agi.RepoHomepage = homepage
			agi.Homepage = homepage
		}
		if description, ok := repoMap["description"].(string); ok {
			agi.RepoDescription = description
			agi.Description = description
		}
		//作者信息
		if ownerMap, ok := repoMap["owner"].(map[string]interface{}); ok {
			agi.RepoOwner = ownerMap["login"].(string)
		}
		//证书信息
		if licenseMap, ok := repoMap["license"].(map[string]interface{}); ok {
			if licenseName, ok := licenseMap["name"].(string); ok {
				agi.RepoLicenseName = licenseName
			}
			if licenseSpdxId, ok := licenseMap["spdx_id"].(string); ok {
				agi.RepoLicenseSpdxId = licenseSpdxId
			}
			if licenseURL, ok := licenseMap["url"].(string); ok {
				agi.RepoLicenseURL = licenseURL
			}
		}
		//日期时间信息
		if createAtStr, ok := repoMap["created_at"].(string); ok {
			createAt, err := time.Parse(time.RFC3339, createAtStr)
			if err == nil {
				agi.RepoCreatedAt = createAt
			}
		}
		if pushedAtStr, ok := repoMap["pushed_at"].(string); ok {
			pushedAt, err := time.Parse(time.RFC3339, pushedAtStr)
			if err == nil {
				agi.RepoPushedAt = pushedAt
			}
		}
		//size可能带小数
		if sizeFloat64, ok := repoMap["size"].(float64); ok {
			size := math.RoundToEven(sizeFloat64);
			agi.RepoSize = int64(size)
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
//UpdateReadme 更新项目中的README.md文件
func UpdateReadme() {
	out, err := exec.Command(git, add, spot).Output()
	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))

	out, err = exec.Command(git, commit, m, msg).Output()
	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))

	out, err = exec.Command(git, push).Output()
	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))
}

type data struct {
	Catagorys []AGI
	GoRepos []AGI
}



//GenerateMd 生成README.md文件
func GenerateMd()  {
	agis, err := GetAGITree(false)
	if err != nil {
		log.Println(err)
		return
	}
	allAGIs, err := GetAGITree(true)
	if err != nil {
		log.Println(err)
		return
	}

	t := template.Must(template.ParseFiles(README_TEMPLATES_PATH))
	f, _ := os.Create(README_OUTEPUT_PATH)
	data := &data{
		Catagorys: agis,
		GoRepos: allAGIs,
	}
	t.Execute(f, data)
}

func getCategoryHtmlId(categoryName string) (id string) {
	id = reSpecialCharacters.ReplaceAllString(categoryName, "-")
	id = strings.Trim(id, "-")
	return strings.ToLower(id)
}