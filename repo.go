package main

import (
	"awesome-go-data/api"
	"flag"
	"log"
)

var accessToken = flag.String("t", "", "GitHub API access_token, 必须输入")
var user = flag.String("u", "", "数据库用户名, 必须输入")
var password = flag.String("p", "", "数据库密码, 必须输入")

func main() {
	flag.Parse() // Scans the arg list and sets up flags
	if *accessToken == "" {
		log.Fatal("GitHub API access_token为空")
	}
	if *user == "" {
		log.Fatal("数据库用户名为空")
	}
	if *password == "" {
		log.Fatal("数据库密码为空")
	}
	api.User = *user
	api.Password = *password
	//获取avelino/awesome-go项目中最新的README.md文件
	//readmeFilePath := api.DownloadReadmeFile()
	//解析awesome-go中的README.md文件,并存入数据库中
	//api.ParseReadmeFile(*accessToken, readmeFilePath)

	api.GenerateMd()
	
	//api.UpdateReadme()
}
