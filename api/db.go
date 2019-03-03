package api

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

// Seconds-based time units
const (
	Minute = 60
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 30 * Day
	Year   = 12 * Month
)

var db *sql.DB

func GetDB() (*sql.DB) {
	if db == nil {
		connStr := "user=postgres password=1234 dbname=agd sslmode=disable"
		db, _ = sql.Open("postgres", connStr)
	}
	return db
}

//UpdateAGIDescription 更新awesome_go_info表description
func UpdateAGIDescription(description string, id int64)  {
	db := GetDB()
	sqlStr := `update awesome_go_info set description = $1, modify_time = CURRENT_TIMESTAMP where id = $2`
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	stmt.Exec(description, id)
}
//UpdateAGIGithubInfo 更新仓库github仓库相关信息
func UpdateAGIGithubInfo(agi *AGI, id int64)  {
	db := GetDB()
	sqlStr := `update awesome_go_info set 
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
	stmt.Exec(agi.RepoHtmlURL, agi.RepoDescription, agi.RepoPushedAt, agi.RepoHomepage, agi.RepoSize, 
			agi.RepoForksCount, agi.RepoStargazersCount, agi.RepoSubscribersCount,
			agi.RepoOpenIssuesCount, agi.RepoLicenseName, agi.RepoLicenseSpdxId,
			agi.RepoLicenseURL, agi.Name, agi.Description, agi.Homepage, agi.ParentId, id)
}
//GetAGI 获取awesome_go_info表数据
func GetAGI(name string, repo bool, category bool) (agi *AGI, err error) {
	db := GetDB()
	sqlStr := `select id, name from awesome_go_info where name = $1 and repo = $2 and category = $3`
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	agi = new(AGI)
	err = stmt.QueryRow(name, repo, category).Scan(&agi.Id, &agi.Name)
	return
}

//GetAGIByCategoryHtmlId 根据CategoryHtmlId获取awesome_go_info表分类数据(类别以锚点id为唯一标识)
func GetAGIByCategoryHtmlId(categoryHtmlId string) (agi *AGI, err error) {
	db := GetDB()
	sqlStr := `select id, name from awesome_go_info where category_html_id = $1 and category = true`
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	agi = new(AGI)
	err = stmt.QueryRow(categoryHtmlId).Scan(&agi.Id, &agi.Name)
	return
}
//ModifyAGIParentIdByCategoryHtmlId 更新awesome_go_info表parent_id
func ModifyAGIParentIdByCategoryHtmlId(parentId int64, categoryHtmlId string)  {
	db := GetDB()
	sqlStr := `update awesome_go_info set parent_id = $1, modify_time = CURRENT_TIMESTAMP where category_html_id = $2 and category = true`
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	stmt.Exec(parentId, categoryHtmlId)
}

//SaveAGI awesome_go_info表插入数据
func SaveAGI(agi *AGI)  {
	db := GetDB()
	sqlStr := `insert into awesome_go_info 
				(id, parent_id, repo_name, repo_full_name, repo_owner, repo_html_url, repo_description, repo_created_at, repo_pushed_at, repo_homepage, repo_size, repo_forks_count, repo_stargazers_count, repo_subscribers_count, repo_open_issues_count, repo_license_name, repo_license_spdx_id, repo_license_url, repo, category, name, description, homepage, category_html_id) 
				values 
				(DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23) 
				RETURNING id`
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	err := stmt.QueryRow(agi.ParentId, agi.RepoName, agi.RepoFullName, agi.RepoOwner, agi.RepoHtmlURL, agi.RepoDescription, agi.RepoCreatedAt, agi.RepoPushedAt, agi.RepoHomepage, agi.RepoSize, agi.RepoForksCount, agi.RepoStargazersCount, agi.RepoSubscribersCount, agi.RepoOpenIssuesCount, agi.RepoLicenseName, agi.RepoLicenseSpdxId, agi.RepoLicenseURL, agi.Repo, agi.Category, agi.Name, agi.Description, agi.Homepage, agi.CategoryHtmlId).Scan(&agi.Id)
	if err != nil {
		log.Println("插入仓库信息报错：", err)
	}
}
//GetAGITree 获取awesome_go_info信息树
func GetAGITree(all bool) (agi []AGI, err error) {
	db := GetDB()
	sqlStr := `WITH RECURSIVE subordinates AS (
 				SELECT id, parent_id, repo_name, repo_full_name, repo_owner, repo_html_url, repo_description, repo_created_at, repo_pushed_at, repo_homepage, repo_size, repo_forks_count, repo_stargazers_count, repo_subscribers_count, repo_open_issues_count, repo_license_name, repo_license_spdx_id, repo_license_url, repo, category, name, description, homepage, category_html_id, 1 depth,ARRAY[ID] AS path FROM awesome_go_info WHERE parent_id = 0 and category=true
 				UNION ALL
 				SELECT  r.id, r.parent_id, r.repo_name, r.repo_full_name, r.repo_owner, r.repo_html_url, r.repo_description, r.repo_created_at, r.repo_pushed_at, r.repo_homepage, r.repo_size, r.repo_forks_count, r.repo_stargazers_count, r.repo_subscribers_count, r.repo_open_issues_count, r.repo_license_name, r.repo_license_spdx_id, r.repo_license_url, r.repo, r.category, r.name, r.description, r.homepage, r.category_html_id, s.depth+1,s.path||r.id  depth FROM awesome_go_info r
				inner JOIN subordinates s ON r.parent_id = s.id  where r.category=true 
				) SELECT id, parent_id, repo_name, repo_full_name, repo_owner, repo_html_url, repo_description, repo_created_at, repo_pushed_at, repo_homepage, repo_size, repo_forks_count, repo_stargazers_count, repo_subscribers_count, repo_open_issues_count, repo_license_name, repo_license_spdx_id, repo_license_url, repo, category, name, description, homepage, category_html_id, depth FROM subordinates order by path`
	if all {
		sqlStr = `WITH RECURSIVE subordinates AS (
 				SELECT id, parent_id, repo_name, repo_full_name, repo_owner, repo_html_url, repo_description, repo_created_at, repo_pushed_at, repo_homepage, repo_size, repo_forks_count, repo_stargazers_count, repo_subscribers_count, repo_open_issues_count, repo_license_name, repo_license_spdx_id, repo_license_url, repo, category, name, description, homepage, category_html_id, 1 depth,ARRAY[ID] AS path FROM awesome_go_info WHERE parent_id = 0
 				UNION ALL
 				SELECT  r.id, r.parent_id, r.repo_name, r.repo_full_name, r.repo_owner, r.repo_html_url, r.repo_description, r.repo_created_at, r.repo_pushed_at, r.repo_homepage, r.repo_size, r.repo_forks_count, r.repo_stargazers_count, r.repo_subscribers_count, r.repo_open_issues_count, r.repo_license_name, r.repo_license_spdx_id, r.repo_license_url, r.repo, r.category, r.name, r.description, r.homepage, r.category_html_id, s.depth+1,s.path||r.id  depth FROM awesome_go_info r
				inner JOIN subordinates s ON r.parent_id = s.id
				) SELECT id, parent_id, repo_name, repo_full_name, repo_owner, repo_html_url, repo_description, repo_created_at, repo_pushed_at, repo_homepage, repo_size, repo_forks_count, repo_stargazers_count, repo_subscribers_count, repo_open_issues_count, repo_license_name, repo_license_spdx_id, repo_license_url, repo, category, name, description, homepage, category_html_id, depth FROM subordinates order by path`
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
		tmpAGI := new(AGI)
		scan := rows.Scan(&tmpAGI.Id, &tmpAGI.ParentId, &tmpAGI.RepoName, &tmpAGI.RepoFullName, &tmpAGI.RepoOwner, &tmpAGI.RepoHtmlURL, &tmpAGI.RepoDescription, &tmpAGI.RepoCreatedAt, &tmpAGI.RepoPushedAt, &tmpAGI.RepoHomepage, &tmpAGI.RepoSize, &tmpAGI.RepoForksCount, &tmpAGI.RepoStargazersCount, &tmpAGI.RepoSubscribersCount, &tmpAGI.RepoOpenIssuesCount, &tmpAGI.RepoLicenseName, &tmpAGI.RepoLicenseSpdxId, &tmpAGI.RepoLicenseURL, &tmpAGI.Repo, &tmpAGI.Category, &tmpAGI.Name, &tmpAGI.Description, &tmpAGI.Homepage, &tmpAGI.CategoryHtmlId, &tmpAGI.Depth)
		if scan != nil {
			log.Println(scan)
			continue
		}
		tmpAGI.Spaces = getSpace(tmpAGI.Depth-1)
		tmpAGI.TitleMarks = getTitleMarks(tmpAGI.Depth)
		tmpAGI.RepoCreatedAtStr = tmpAGI.RepoCreatedAt.Format("2006-01-02")
		tmpAGI.RepoPushedAtStr = tmpAGI.RepoPushedAt.Format("2006-01-02 15:04:05")
		tmpAGI.TimeSince = timeSince(tmpAGI.RepoPushedAt)
		agi = append(agi, *tmpAGI)
		if i > 0 && agi[i-1].Category && tmpAGI.Repo {
			agi[i-1].WithReposTable = true
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
// 代码复制自 https://github.com/gogs/gogs/blob/436dd6c0a4549e09069193879046a2a202d6291e/pkg/tool/tool.go
func timeSince(then time.Time) string {
	now := time.Now()

	lbl := "ago"
	diff := now.Unix() - then.Unix()
	if then.After(now) {
		lbl = "from now"
		diff = then.Unix() - now.Unix()
	}

	switch {
	case diff <= 0:
		return "now"
	case diff <= 2:
		return fmt.Sprintf("1 second %s", lbl)
	case diff < 1*Minute:
		return fmt.Sprintf("%d seconds %s", diff, lbl)

	case diff < 2*Minute:
		return fmt.Sprintf("1 minute %s", lbl)
	case diff < 1*Hour:
		return fmt.Sprintf("%d minutes %s", diff/Minute, lbl)

	case diff < 2*Hour:
		return fmt.Sprintf("1 hour %s", lbl)
	case diff < 1*Day:
		return fmt.Sprintf("%d hours %s", diff/Hour, lbl)

	case diff < 2*Day:
		return fmt.Sprintf("1 day %s", lbl)
	case diff < 1*Week:
		return fmt.Sprintf("%d days %s", diff/Day, lbl)

	case diff < 2*Week:
		return fmt.Sprintf("1 week %s", lbl)
	case diff < 1*Month:
		return fmt.Sprintf("%d weeks %s", diff/Week, lbl)

	case diff < 2*Month:
		return fmt.Sprintf("1 month %s", lbl)
	case diff < 1*Year:
		return fmt.Sprintf("%d months %s", diff/Month, lbl)

	case diff < 2*Year:
		return fmt.Sprintf("1 year %s", lbl)
	default:
		return fmt.Sprintf("%d years %s", diff/Year, lbl)
	}
}