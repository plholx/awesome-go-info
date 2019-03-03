package api

import (
	"time"
)

//awesome_go_info表对应结构体
type AGI struct {
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
	CategoryHtmlId string `json:"category_html_id,omitempty"`

	Depth int64 `json:"depth,omitempty"`
	Spaces string `json:"spaces,omitempty"`
	TitleMarks string `json:"title_marks,omitempty"`
	RepoCreatedAtStr string `json:"repo_created_at_str,omitempty"`
	RepoPushedAtStr string `json:"repo_pushed_at_str,omitempty"`
	WithReposTable bool `json:"with_repos_table,omitempty`
	TimeSince string `json:"time_since,omitempty"`
}