# Awesome Go Data

go语言开源项目列表，项目分类及GitHub上的开源项目数据完全来自于[awesome-go](https://github.com/avelino/awesome-go) 的[README.md](https://github.com/avelino/awesome-go/blob/master/README.md)文件，通过调用GitHub的API获取仓库信息，展示项目的star数、watch数等，方便更直观的查看go语言开源项目的一些数据。

{{range .Catagorys}}{{.Spaces}}- [{{.Name}}](#{{.CategoryHtmlId}})
{{end}}

{{range .GoRepos}}
    {{- if .Category}}
        {{- "\n\n"}}{{.TitleMarks}} {{.Name}}
        {{if .Description}}
            {{- "\n*"}}{{.Description}}*
        {{- end}}
        {{- if .WithReposTable}}
            {{- "\n\n|"}} Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
            {{- "\n|"}} :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
        {{- end}}
    {{- else if .Repo}}
         {{- "\n|"}} [{{.RepoName}}]({{.RepoHtmlURL -}}) | {{.RepoStargazersCount}} | {{.RepoSubscribersCount}} | {{.RepoCreatedAtStr}} | {{.TimeSince}} | {{.Description}} |
    {{- end}}
{{- end}}