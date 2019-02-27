package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

const (
	sourceFileURL = "https://raw.githubusercontent.com/avelino/awesome-go/master/README.md"
	githubReposAPI = "https://api.github.com/repos/:owner/:repo?access_token=OAUTH-TOKEN"
	githubRateLimitAPI = "https://api.github.com/rate_limit?access_token=OAUTH-TOKEN"
	githubDomain = "https://github.com"
)
//下载awesome-go中的README.md文件
func downloadReadmeFile()  {
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
var (
	reCategoryLi          = regexp.MustCompile(`(\s*)- \[(.*)\]\(#(.*)\)`)
	reCategory            = regexp.MustCompile(`#+ (.+)`)
	reCategoryDescription = regexp.MustCompile(`(\s*)\*(.*)\*$`)
	reContainsLink        = regexp.MustCompile(`\* \[.*\]\(.*\)`)
	reOnlyLink            = regexp.MustCompile(`(\s*)\* \[(.*)\]\((.+)\)$`)
	reLinkWithDescription = regexp.MustCompile(`(\s*)\* \[(.*?)\]\((.+?)\) - (\S.*[\.\!])`)
	reLittleCategory      = regexp.MustCompile(`(\s*)\* ([a-zA-Z\s]*)$`)
	reGitHubURL           = regexp.MustCompile(`https://github.com/(.+?)/(.+?)\b`)
)
//解析awesome-go中的README.md文件
func parseReadmeFile()  {
	input, err := ioutil.ReadFile("data/readmeFiles/README.md")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")
	categoryIds := make(map[int]int64)
	var tmpCategoryId int64 = 0
	var linkCategoryId int64 = 0
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
				log.Println("开始请求仓库信息", name)
				repo, err := GetRepoInfo(repoOwner, repoName)
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
//GoRepo github项目仓库信息结构体
type GoRepo struct {
	Id int64 `json:"id,omitempty"`
	ParentId int64 `json:"parent_id,omitempty"`
	RepoName string `json:"repo_name,omitempty"`
	RepoFullName string `json:"repo_full_name,omitempty"`
	RepoOwner string `json:"repo_owner,omitempty"`
	RepoHtmlURL string `json:"repo_html_url,omitempty"`
	RepoDescription string `json:"repo_description,omitempty"`
	RepoCreatedAt time.Time `json:"repo_created_at,omitempty"`
	RepoPushedAt time.Time `json:"repo_pushed_at,omitempty"`
	RepoHomepage string `json:"repo_homepage,omitempty"`
	RepoSize int64 `json:"repo_size,omitempty"`
	RepoForksCount int64 `json:"repo_forks_count,omitempty"`
	RepoStargazersCount int64 `json:"repo_stargazers_count,omitempty"`
	RepoSubscribersCount int64 `json:"repo_subscribers_count,omitempty"`
	RepoOpenIssuesCount int64 `json:"repo_open_issues_count,omitempty"`
	RepoLicenseName string `json:"repo_license_name,omitempty"`
	RepoLicenseSpdxId string `json:"repo_license_spdx_id,omitempty"`
	RepoLicenseURL string `json:"repo_license_url,omitempty"`
	AddTime time.Time `json:"add_time,omitempty"`
	ModifyTime time.Time `json:"modify_time,omitempty"`
	Repo bool `json:"repo,omitempty"`
	Category bool `json:"category,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Homepage string `json:"homepage,omitempty"`
}

var db *sql.DB

func GetDB() (*sql.DB) {
	if db == nil {
		connStr := "user=postgres password=1234 dbname=agd sslmode=disable"
		db, _ = sql.Open("postgres", connStr)
	}
	return db
}

//保存goRepo
func SaveGoRepo(goRepo *GoRepo)  {
	db := GetDB()
	sqlStr := `insert into go_repo 
				(id, parent_id, repo_name, repo_full_name, repo_owner, repo_html_url, repo_description, repo_created_at, repo_pushed_at, repo_homepage, repo_size, repo_forks_count, repo_stargazers_count, repo_subscribers_count, repo_open_issues_count, repo_license_name, repo_license_spdx_id, repo_license_url, repo, category, name, description, homepage) 
				values 
				(DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22) 
				RETURNING id`
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	err := stmt.QueryRow(goRepo.ParentId, goRepo.RepoName, goRepo.RepoFullName, goRepo.RepoOwner, goRepo.RepoHtmlURL, goRepo.RepoDescription, goRepo.RepoCreatedAt, goRepo.RepoPushedAt, goRepo.RepoHomepage, goRepo.RepoSize, goRepo.RepoForksCount, goRepo.RepoStargazersCount, goRepo.RepoSubscribersCount, goRepo.RepoOpenIssuesCount, goRepo.RepoLicenseName, goRepo.RepoLicenseSpdxId, goRepo.RepoLicenseURL, goRepo.Repo, goRepo.Category, goRepo.Name, goRepo.Description, goRepo.Homepage).Scan(&goRepo.Id)
	if err != nil {
		log.Println("插入仓库信息报错：", err)
	}
}
func UpdateGoRepoParentId(parentId int64, name string, repo bool, category bool)  {
	db := GetDB()
	sqlStr := `update go_repo set parent_id = $1, modify_time = CURRENT_TIMESTAMP where name = $2 and repo = $3 and category = $4`
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	stmt.Exec(parentId, name, repo, category)
}
func UpdateGoRepoDescription(description string, id int64)  {
	db := GetDB()
	sqlStr := `update go_repo set description = $1, modify_time = CURRENT_TIMESTAMP where id = $2`
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	stmt.Exec(description, id)
}
// 更新仓库github信息
func UpdateGoRepoGithubInfo(repo *GoRepo, id int64)  {
	db := GetDB()
	sqlStr := `update go_repo set 
				modify_time = CURRENT_TIMESTAMP,
				repo_html_url = $1,
				repo_description = $2,
				repo_pushed_at = $3,
				repo_homepage = $4,
				repo_size = $5,
				repo_forks_count = $6,
				repo_stargazers_count = $7,
				repo_subscribers_count = $8,
				repo_open_issues_count = $9,
				repo_license_name = $10,
				repo_license_spdx_id = $11,
				repo_license_url = $12,
				name = $13,
				description = $14,
				homepage = $15,
				parent_id = $16
				where id = $17`
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	stmt.Exec(repo.RepoHtmlURL, repo.RepoDescription, repo.RepoPushedAt, repo.RepoHomepage, repo.RepoSize, 
			repo.RepoForksCount, repo.RepoStargazersCount, repo.RepoSubscribersCount,
			repo.RepoOpenIssuesCount, repo.RepoLicenseName, repo.RepoLicenseSpdxId,
			repo.RepoLicenseURL, repo.Name, repo.Description, repo.Homepage, repo.ParentId, id)
}
func GetGoRepo(name string, repo bool, category bool) (goRepo *GoRepo, err error) {
	db := GetDB()
	sqlStr := `select id, name from go_repo where name = $1 and repo = $2 and category = $3`
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	goRepo = new(GoRepo)
	err = stmt.QueryRow(name, repo, category).Scan(&goRepo.Id, &goRepo.Name)
	return
}
//获取github仓库信息
func GetRepoInfo(repoOwner, repoName string) (repo *GoRepo, err error) {
	repoFullName := repoOwner + "/" + repoName
	apiURL := strings.Replace(githubReposAPI, ":owner/:repo", repoFullName, -1)
	apiURL = strings.Replace(apiURL, "OAUTH-TOKEN", *accessToken, -1)
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
func ControlGitHubAPIReq()  {
	//通过访问rate_limit API，校验access_token是否有效，及通过剩余次数控制API的访问时机
	apiURL := strings.Replace(githubRateLimitAPI, "OAUTH-TOKEN", *accessToken, -1)
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
		log.Println(rateLimitMap)
	}
}
var accessToken = flag.String("t", "", "GitHub API access_token, 必须输入")

func main() {
	flag.Parse() // Scans the arg list and sets up flags
	if *accessToken == "" {
		log.Fatal("GitHub API access_token为空")
	}

	parseReadmeFile()

	// repo, err := GetRepoInfo("avelino", "awesome-go")
	// log.Println(repo)
	// log.Println(err)

	//goRepo, err := GetGoRepo("awesome-go2", true, false)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(goRepo)

	//log.Println("%", strings.Trim("  x   ", " "), "%")
	//log.Println("%", strings.TrimSpace("  x   "), "%")

	// log.Println(reCategoryLi.MatchString("    - [Audio and Music](#audio-and-music)"))
	// log.Println(reOnlyLink.MatchString("* [gopher-stickers](https://github.com/tenntenn/gopher-stickers)"))
	// log.Println(reLinkWithDescription.MatchString("* [mix](https://github.com/go-mix/mix) - Sequence-based Go-native audio mixer for music apps."))
	// log.Println(reLittleCategory.MatchString(`* [gmf](https://github.com/3d0c/gmf) - Go bindings for FFmpeg av\* libraries.`))
	// subMatchs := reGitHubURL.FindStringSubmatch("https://github.com/rakyll/portmidi?")
	// log.Println(subMatchs[1], "%", subMatchs[2])

	//接口中的时间获取后转换有问题，后续需要解决
	// createAt, _ := time.Parse(time.RFC3339, "2014-07-06T13:42:15Z")
	// log.Println(createAt)
	// log.Println(time.Now())

	//reCategoryLi          = regexp.MustCompile(`(\s*)- \[(.*)\]\(#(.*)\)`)
	// line := "    - [Audio and Music](#audio-and-music)"
	// if reCategoryLi.MatchString(line) {
	// 	subMatchs := reCategoryLi.FindStringSubmatch(line)
	// 	log.Println(subMatchs)
		
	// 	// log.Println(subMatchs[1][1])
	// }
}
