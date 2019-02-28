package api

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func GetDB() (*sql.DB) {
	if db == nil {
		connStr := "user=postgres password=1234 dbname=agd sslmode=disable"
		db, _ = sql.Open("postgres", connStr)
	}
	return db
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