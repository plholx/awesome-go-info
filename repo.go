package main

import (
	"awesome-go-data/api"
	"flag"
	"log"
	"os"
	"time"
)

var accessToken = flag.String("t", "", "GitHub API access_token, 必须输入")
var user = flag.String("u", "", "数据库用户名, 必须输入")
var password = flag.String("p", "", "数据库密码, 必须输入")

func main() {
	configLog()

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

	execute()
}

func execute()  {
	for {
		log.Println("业务开始执行：", time.Now().Format("2006-01-02 15:04:05"))

		//获取avelino/awesome-go项目中最新的README.md文件
		readmeFilePath := api.DownloadReadmeFile()
		//解析awesome-go中的README.md文件,并存入数据库中
		api.ParseReadmeFile(*accessToken, readmeFilePath)
		//生成README.md
		api.GenerateMd()
		//api.UpdateReadme()

		log.Println("业务执行结束：", time.Now().Format("2006-01-02 15:04:05"))
		log.Println("等待下次执行......")

		curTime := time.Now()
		nextTime := curTime.Add(time.Hour * 24)
		nextTime = time.Date(nextTime.Year(), nextTime.Month(), nextTime.Day(), 0, 0, 0, 0, nextTime.Location())
		//nextTime := curTime.Add(time.Second * 5)
		timer := time.NewTimer(nextTime.Sub(curTime))
		curTime = <- timer.C
	}
}

func configLog()  {
	logFile, err := os.OpenFile("log/agd.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
}
