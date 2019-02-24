package main

import (
	"bufio"
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
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
func parseReadmeFile()  {
	file, err := os.Open("data/readmeFiles/README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

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
}

var db *sql.DB

func GetDB() (*sql.DB) {
	if db == nil {
		connStr := "user=postgres password=1234 dbname=agd sslmode=disable"
		db, _ = sql.Open("postgres", connStr)
	}
	return db
}

//操作数据库
func opDb()  {

}
func main() {
	downloadReadmeFile()
}
