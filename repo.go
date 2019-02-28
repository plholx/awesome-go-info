package main

import (
	"awesome-go-data/api"
)

var accessToken = flag.String("t", "", "GitHub API access_token, 必须输入")

func main() {
	flag.Parse() // Scans the arg list and sets up flags
	if *accessToken == "" {
		log.Fatal("GitHub API access_token为空")
	}
	downloadReadmeFile()
}
