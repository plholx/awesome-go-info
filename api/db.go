package api

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func GetDB() (*sql.DB) {
	if db == nil {
		connStr := "user=postgres password=1234 dbname=agd sslmode=disable"
		db, _ = sql.Open("postgres", connStr)
	}
	return db
}

//更新go_repo表description
func UpdateGoRepoDescription(description string, id int64)  {
	db := GetDB()
	sqlStr := `update go_repo set description = $1, modify_time = CURRENT_TIMESTAMP where id = $2`
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	stmt.Exec(description, id)
}
// 更新仓库github仓库相关信息
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
//GetGoRepo 获取go_repo表数据
func GetGoRepo(name string, repo bool, category bool) (goRepo *GoRepo, err error) {
	db := GetDB()
	sqlStr := `select id, name from go_repo where name = $1 and repo = $2 and category = $3`
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	goRepo = new(GoRepo)
	err = stmt.QueryRow(name, repo, category).Scan(&goRepo.Id, &goRepo.Name)
	return
}

//GetGoRepo 根据CategoryHtmlId获取go_repo表分类数据(类别以锚点id为唯一标识)
func GetGoRepoByCategoryHtmlId(categoryHtmlId string) (goRepo *GoRepo, err error) {
	db := GetDB()
	sqlStr := `select id, name from go_repo where category_html_id = $1 and category = true`
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	goRepo = new(GoRepo)
	err = stmt.QueryRow(categoryHtmlId).Scan(&goRepo.Id, &goRepo.Name)
	return
}
//更新go_repo表parent_id
func ModifyGoRepoParentIdByCategoryHtmlId(parentId int64, categoryHtmlId string)  {
	db := GetDB()
	sqlStr := `update go_repo set parent_id = $1, modify_time = CURRENT_TIMESTAMP where category_html_id = $2 and category = true`
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	stmt.Exec(parentId, categoryHtmlId)
}

//SaveGoRepo go_repo表插入数据
func SaveGoRepo(goRepo *GoRepo)  {
	db := GetDB()
	sqlStr := `insert into go_repo 
				(id, parent_id, repo_name, repo_full_name, repo_owner, repo_html_url, repo_description, repo_created_at, repo_pushed_at, repo_homepage, repo_size, repo_forks_count, repo_stargazers_count, repo_subscribers_count, repo_open_issues_count, repo_license_name, repo_license_spdx_id, repo_license_url, repo, category, name, description, homepage, category_html_id) 
				values 
				(DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23) 
				RETURNING id`
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	err := stmt.QueryRow(goRepo.ParentId, goRepo.RepoName, goRepo.RepoFullName, goRepo.RepoOwner, goRepo.RepoHtmlURL, goRepo.RepoDescription, goRepo.RepoCreatedAt, goRepo.RepoPushedAt, goRepo.RepoHomepage, goRepo.RepoSize, goRepo.RepoForksCount, goRepo.RepoStargazersCount, goRepo.RepoSubscribersCount, goRepo.RepoOpenIssuesCount, goRepo.RepoLicenseName, goRepo.RepoLicenseSpdxId, goRepo.RepoLicenseURL, goRepo.Repo, goRepo.Category, goRepo.Name, goRepo.Description, goRepo.Homepage, goRepo.CategoryHtmlId).Scan(&goRepo.Id)
	if err != nil {
		log.Println("插入仓库信息报错：", err)
	}
}
//GetRepoTree 获取仓库信息树
func GetRepoTree(all bool) (repos []GoRepo, err error) {
	db := GetDB()
	sqlStr := `WITH RECURSIVE subordinates AS (
 				SELECT id, parent_id, repo_name, repo_full_name, repo_owner, repo_html_url, repo_description, repo_created_at, repo_pushed_at, repo_homepage, repo_size, repo_forks_count, repo_stargazers_count, repo_subscribers_count, repo_open_issues_count, repo_license_name, repo_license_spdx_id, repo_license_url, repo, category, name, description, homepage,1 depth,ARRAY[ID] AS path FROM go_repo WHERE parent_id = 0 and category=true
 				UNION ALL
 				SELECT  r.id, r.parent_id, r.repo_name, r.repo_full_name, r.repo_owner, r.repo_html_url, r.repo_description, r.repo_created_at, r.repo_pushed_at, r.repo_homepage, r.repo_size, r.repo_forks_count, r.repo_stargazers_count, r.repo_subscribers_count, r.repo_open_issues_count, r.repo_license_name, r.repo_license_spdx_id, r.repo_license_url, r.repo, r.category, r.name, r.description, r.homepage, s.depth+1,s.path||r.id  depth FROM go_repo r
				inner JOIN subordinates s ON r.parent_id = s.id  where r.category=true 
				) SELECT id, parent_id, repo_name, repo_full_name, repo_owner, repo_html_url, repo_description, repo_created_at, repo_pushed_at, repo_homepage, repo_size, repo_forks_count, repo_stargazers_count, repo_subscribers_count, repo_open_issues_count, repo_license_name, repo_license_spdx_id, repo_license_url, repo, category, name, description, homepage, depth FROM subordinates order by path`
	if all {
		sqlStr = `WITH RECURSIVE subordinates AS (
 				SELECT id, parent_id, repo_name, repo_full_name, repo_owner, repo_html_url, repo_description, repo_created_at, repo_pushed_at, repo_homepage, repo_size, repo_forks_count, repo_stargazers_count, repo_subscribers_count, repo_open_issues_count, repo_license_name, repo_license_spdx_id, repo_license_url, repo, category, name, description, homepage,1 depth,ARRAY[ID] AS path FROM go_repo WHERE parent_id = 0
 				UNION ALL
 				SELECT  r.id, r.parent_id, r.repo_name, r.repo_full_name, r.repo_owner, r.repo_html_url, r.repo_description, r.repo_created_at, r.repo_pushed_at, r.repo_homepage, r.repo_size, r.repo_forks_count, r.repo_stargazers_count, r.repo_subscribers_count, r.repo_open_issues_count, r.repo_license_name, r.repo_license_spdx_id, r.repo_license_url, r.repo, r.category, r.name, r.description, r.homepage, s.depth+1,s.path||r.id  depth FROM go_repo r
				inner JOIN subordinates s ON r.parent_id = s.id
				) SELECT id, parent_id, repo_name, repo_full_name, repo_owner, repo_html_url, repo_description, repo_created_at, repo_pushed_at, repo_homepage, repo_size, repo_forks_count, repo_stargazers_count, repo_subscribers_count, repo_open_issues_count, repo_license_name, repo_license_spdx_id, repo_license_url, repo, category, name, description, homepage, depth FROM subordinates order by path`
	}
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
		return
	}
	var i int64 = 0
	for rows.Next() {
		tmpRepo := new(GoRepo)
		scan := rows.Scan(&tmpRepo.Id, &tmpRepo.ParentId, &tmpRepo.RepoName, &tmpRepo.RepoFullName, &tmpRepo.RepoOwner, &tmpRepo.RepoHtmlURL, &tmpRepo.RepoDescription, &tmpRepo.RepoCreatedAt, &tmpRepo.RepoPushedAt, &tmpRepo.RepoHomepage, &tmpRepo.RepoSize, &tmpRepo.RepoForksCount, &tmpRepo.RepoStargazersCount, &tmpRepo.RepoSubscribersCount, &tmpRepo.RepoOpenIssuesCount, &tmpRepo.RepoLicenseName, &tmpRepo.RepoLicenseSpdxId, &tmpRepo.RepoLicenseURL, &tmpRepo.Repo, &tmpRepo.Category, &tmpRepo.Name, &tmpRepo.Description, &tmpRepo.Homepage, &tmpRepo.Depth)
		if scan != nil {
			log.Println(scan)
			continue
		}
		tmpRepo.Spaces = getSpace(tmpRepo.Depth-1)
		tmpRepo.TitleMarks = getTitleMarks(tmpRepo.Depth)
		tmpRepo.RepoCreatedAtStr = tmpRepo.RepoCreatedAt.Format("2006-01-02 15:04:05")
		tmpRepo.RepoPushedAtStr = tmpRepo.RepoPushedAt.Format("2006-01-02 15:04:05")
		repos = append(repos, *tmpRepo)
		if i > 0 && repos[i-1].Category && tmpRepo.Repo {
			repos[i-1].WithReposTable = true
		}
		i ++
	}
	return
}

func getSpace(count int64) (s string) {
	for i:=int64(0); i<count; i++ {
		s += "    "
	}
	return s
}
func getTitleMarks(count int64) (s string) {
	for i:=int64(0); i<count; i++ {
		s += "#"
	}
	return s
}
