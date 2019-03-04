package main

import (
	"flag"
	"log"
	"os/exec"

	//"awesome-go-data/api"
)

// memory usage optimizations
const (
	git = `git`
	add = `add`
	spot = `.`
	commit = `commit`
	m = `-m`
	msg = `测试git命令`
	push = `push`
)

var accessToken = flag.String("t", "", "GitHub API access_token, 必须输入")

func main() {
	flag.Parse() // Scans the arg list and sets up flags
	if *accessToken == "" {
		log.Fatal("GitHub API access_token为空")
	}
	//获取avelino/awesome-go项目中最新的README.md文件
	//readmeFilePath := api.DownloadReadmeFile()
	//解析awesome-go中的README.md文件,并存入数据库中
	//api.ParseReadmeFile(*accessToken, readmeFilePath)

	//api.GenerateMd()
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
