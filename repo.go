package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
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
	githubReposAPI = "https://api.github.com/repos/:owner/:repo"
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
	reCategoryDescription = regexp.MustCompile(`\*.+\*`)
	reContainsLink        = regexp.MustCompile(`\* \[.*\]\(.*\)`)
	reOnlyLink            = regexp.MustCompile(`(\s*)\* \[(.*)\]\((.+)\)`)
	reLinkWithDescription = regexp.MustCompile(`(\s*)\* \[(.*)\]\((.+)\) - (\S.*[\.\!])`)
	reLittleCategory      = regexp.MustCompile(`(\s*)\* (.*)`)
)
//解析awesome-go中的README.md文件
func parseReadmeFile()  {
	input, err := ioutil.ReadFile("data/readmeFiles/README.md")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		//分类
		if line == "### Contents" {

		}
		//line = strings.Trim(line, " ")
		/*containsLink = reContainsLink.MatchString(line)
		if containsLink {
			noDescription = reOnlyLink.MatchString(line)
			if noDescription {
				continue
			}

			matched = reLinkWithDescription.MatchString(line)
			if !matched {
				t.Errorf("expected entry to be in form of `* [link] - description.`, got '%s'", line)
			}
		}*/
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
	//goRepo := new(GoRepo)
	db := GetDB()
	sqlStr := `insert into go_repo 
				(id, parent_id, repo_name, repo_full_name, repo_owner, repo_html_url, repo_description, repo_created_at, repo_pushed_at, repo_homepage, repo_size, repo_forks_count, repo_stargazers_count, repo_subscribers_count, repo_open_issues_count, repo_license_name, repo_license_spdx_id, repo_license_url, repo, category, name, description, homepage) 
				values 
				(DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22) 
				RETURNING id`
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	stmt.QueryRow(goRepo.ParentId, goRepo.RepoName, goRepo.RepoFullName, goRepo.RepoOwner, goRepo.RepoHtmlURL, goRepo.RepoDescription, goRepo.RepoCreatedAt, goRepo.RepoPushedAt, goRepo.RepoHomepage, goRepo.RepoSize, goRepo.RepoForksCount, goRepo.RepoStargazersCount, goRepo.RepoSubscribersCount, goRepo.RepoOpenIssuesCount, goRepo.RepoLicenseName, goRepo.RepoLicenseSpdxId, goRepo.RepoLicenseURL, goRepo.Repo, goRepo.Category, goRepo.Name, goRepo.Description, goRepo.Homepage).Scan(goRepo.Id)
}
//获取github仓库信息并保存
func GetRepoInfoAndSave() {
	apiURL := strings.Replace(githubReposAPI, ":owner/:repo", "avelino/awesome-go", -1)
	res, err := http.Get(apiURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	bytes, e := ioutil.ReadAll(res.Body)
	if e != nil {
		log.Fatal(err)
	}
	var v interface{}
	json.Unmarshal(bytes, &v)
	if repoMap, ok := v.(map[string]interface{}); ok {
		repo := &GoRepo{
			ParentId:0, 
			RepoName:repoMap["name"].(string), 
			RepoFullName:repoMap["full_name"].(string), 
			RepoHtmlURL:repoMap["html_url"].(string), 
			RepoDescription:repoMap["description"].(string), 
			RepoHomepage:repoMap["homepage"].(string), 
			RepoForksCount:int64(repoMap["forks_count"].(float64)), 
			RepoStargazersCount:int64(repoMap["stargazers_count"].(float64)), 
			RepoSubscribersCount:int64(repoMap["subscribers_count"].(float64)), 
			RepoOpenIssuesCount:int64(repoMap["open_issues_count"].(float64)), 
			Repo:true, 
			Category:false, 
			Name:repoMap["name"].(string), 
			Description:repoMap["description"].(string), 
			Homepage:repoMap["homepage"].(string),
		}
		//作者信息
		if ownerMap, ok := repoMap["owner"].(map[string]interface{}); ok {
			repo.RepoOwner = ownerMap["login"].(string)
		}
		//证书信息
		if licenseMap, ok := repoMap["license"].(map[string]interface{}); ok {
			repo.RepoLicenseName = licenseMap["name"].(string)
			repo.RepoLicenseSpdxId = licenseMap["spdx_id"].(string)
			repo.RepoLicenseURL = licenseMap["url"].(string)
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
		SaveGoRepo(repo)
	}

}
func main() {
	//GetRepoInfoAndSave()

	//log.Println("%", strings.Trim("  x   ", " "), "%")
	//log.Println("%", strings.TrimSpace("  x   "), "%")

	log.Println(reCategoryLi.MatchString("    - [Audio and Music](#audio-and-music)"))
	log.Println(reOnlyLink.MatchString("* [gopher-stickers](https://github.com/tenntenn/gopher-stickers)"))
	log.Println(reLinkWithDescription.MatchString("* [mix](https://github.com/go-mix/mix) - Sequence-based Go-native audio mixer for music apps."))
	log.Println(reLittleCategory.MatchString("* Testing Frameworks"))

	//接口中的时间获取后转换有问题，后续需要解决
	// createAt, _ := time.Parse(time.RFC3339, "2014-07-06T13:42:15Z")
	// log.Println(createAt)
	// log.Println(time.Now())
}
