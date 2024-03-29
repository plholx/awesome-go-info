# Awesome Go Info

go语言开源项目列表，项目分类及GitHub上的开源项目数据完全来自于 [awesome-go](https://github.com/avelino/awesome-go) 的 [README.md](https://github.com/avelino/awesome-go/blob/master/README.md) 文件，通过调用GitHub的API获取仓库信息，展示项目的star数、watch数等，方便查看go语言开源项目的一些相关信息。

_该文件仅包含 [awesome-go](https://github.com/avelino/awesome-go) 仓库中 [README.md](https://github.com/avelino/awesome-go/blob/master/README.md) 文件中列出的在GitHub上开源的优秀项目，不罗列其它golang相关的网站_
_该文件中的GitHub仓库信息数据会在每天凌晨1点左右更新,当前数据更新于2022-02-16 00:00:01_

- [Awesome Go](#awesome-go)
    - [Advanced Console UIs](#advanced-console-uis)
    - [Dependency Injection](#dependency-injection)
    - [DevOps Tools](#devops-tools)
    - [HTTP Clients](#http-clients)
    - [Microsoft Excel](#microsoft-excel)
    - [Middlewares](#middlewares)
    - [Other Software](#other-software)
    - [Project Layout](#project-layout)
    - [Routers](#routers)
    - [Standard CLI](#standard-cli)
    - [Strings](#strings)
    - [Uncategorized](#uncategorized)
- [Audio and Music](#audio-and-music)
- [Authentication and OAuth](#authentication-and-oauth)
- [Blockchain](#blockchain)
- [Bot Building](#bot-building)
- [Build Automation](#build-automation)
- [Command Line](#command-line)
- [Configuration](#configuration)
- [Continuous Integration](#continuous-integration)
- [CSS Preprocessors](#css-preprocessors)
- [Data Structures](#data-structures)
- [Database](#database)
- [Database Drivers](#database-drivers)
    - [NoSQL Databases](#nosql-databases)
    - [Relational Databases](#relational-databases)
- [Date and Time](#date-and-time)
- [Distributed Systems](#distributed-systems)
- [Dynamic DNS](#dynamic-dns)
- [Email](#email)
- [Embeddable Scripting Languages](#embeddable-scripting-languages)
- [Error Handling](#error-handling)
- [File Handling](#file-handling)
- [Financial](#financial)
- [Forms](#forms)
- [Functional](#functional)
- [Game Development](#game-development)
- [Generation and Generics](#generation-and-generics)
- [Geographic](#geographic)
- [Go Compilers](#go-compilers)
- [Goroutines](#goroutines)
- [GUI](#gui)
- [Hardware](#hardware)
- [Images](#images)
- [IoT](#iot-internet-of-things)
- [Job Scheduler](#job-scheduler)
- [JSON](#json)
- [Logging](#logging)
- [Machine Learning](#machine-learning)
- [Messaging](#messaging)
- [Microsoft Office](#microsoft-office)
- [Miscellaneous](#miscellaneous)
- [Natural Language Processing](#natural-language-processing)
- [Networking](#networking)
- [OpenGL](#opengl)
- [ORM](#orm)
- [Package Management](#package-management)
- [Performance](#performance)
- [Query Language](#query-language)
- [Resource Embedding](#resource-embedding)
- [Science and Data Analysis](#science-and-data-analysis)
- [Security](#security)
- [Serialization](#serialization)
- [Server Applications](#server-applications)
- [Stream Processing](#stream-processing)
- [Template Engines](#template-engines)
- [Testing](#testing)
    - [Fail injection](#fail-injection)
    - [Mock](#mock)
    - [Testing Frameworks](#testing-frameworks)
- [Text Processing](#text-processing)
    - [Specific Formats](#specific-formats)
    - [Utility](#utility)
- [Third-party APIs](#third-party-apis)
- [Utilities](#utilities)
- [UUID](#uuid)
- [Validation](#validation)
- [Version Control](#version-control)
- [Video](#video)
- [Web Frameworks](#web-frameworks)
    - [Actual middlewares](#actual-middlewares)
    - [Libraries for creating HTTP middlewares](#libraries-for-creating-http-middlewares)
- [WebAssembly](#webassembly)
- [Windows](#windows)
- [XML](#xml)
- [Zero Trust](#zero-trust)
- [Tools](#tools)
- [Code Analysis](#code-analysis)
- [Editor Plugins](#editor-plugins)
- [Go Generate Tools](#go-generate-tools)
- [Go Tools](#go-tools)
- [Software Packages](#software-packages)
- [Resources](#resources)
    - [E-books for purchase](#e-books-for-purchase)
    - [Free e-books](#free-e-books)
    - [Reddit](#reddit)
    - [Tutorials](#tutorials)
    - [Twitter](#twitter)
- [Benchmarks](#benchmarks)
- [Conferences](#conferences)
- [E-Books](#e-books)
- [Gophers](#gophers)
- [Meetups](#meetups)
- [Style Guides](#style-guides)
- [Websites](#websites)
- [Social Media](#social-media)



**[⬆ back to top](#awesome-go-info)**

# Awesome Go
        
**Contents:**
**[⬆ back to top](#awesome-go-info)**

## Advanced Console UIs
        
**[⬆ back to top](#awesome-go-info)**

## Dependency Injection
        
*Libraries for working with dependency injection.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [wire](https://github.com/google/wire) | 7390 | 93 | 2018/11/28 | 2 weeks ago | Automated Initialization in Go. |
| [fx](https://github.com/uber-go/fx) | 2491 | 65 | 2016/10/27 | 1 week ago | A dependency injection based application framework for Go (built on top of dig). |
| [dig](https://github.com/uber-go/dig) | 2329 | 48 | 2017/03/21 | 2 weeks ago | A reflection based dependency injection toolkit for Go. |
| [container](https://github.com/golobby/container) | 298 | 5 | 2019/09/23 | 2 weeks ago | A powerful IoC Container with fluent and easy-to-use interface. |
| [di](https://github.com/goioc/di) | 149 | 6 | 2020/06/11 | 1 month ago | Spring-inspired Dependency Injection Container. |
| [di](https://github.com/goava/di) | 136 | 9 | 2020/02/03 | 2 months ago | A dependency injection container for go programming language. |
| [dingo](https://github.com/i-love-flamingo/dingo) | 122 | 26 | 2018/10/29 | 9 months ago | A dependency injection toolkit for Go, based on Guice. |
| [alice](https://github.com/magic003/alice) | 43 | 2 | 2017/04/08 | 4 years ago | Additive dependency injection container for Golang. |
| [wire](https://github.com/Fs02/wire) | 34 | 2 | 2018/07/05 | 5 months ago | Strict Runtime Dependency Injection for Golang. |
| [linker](https://github.com/logrange/linker) | 32 | 4 | 2018/12/04 | 1 year ago | A reflection based dependency injection and inversion of control library with components lifecycle support. |
| [gocontainer](https://github.com/vardius/gocontainer) | 14 | 0 | 2019/06/06 | 1 year ago | Simple Dependency Injection Container. |
| [kinit](https://github.com/go-kata/kinit) | 5 | 3 | 2021/01/24 | 8 months ago | Customizable dependency injection container with the global mode, cascade initialization and panic-safe finalization. |
| [nject](https://github.com/muir/nject) | 5 | 1 | 2021/09/15 | 3 weeks ago | A type safe, reflective framework for libraries, tests, and http endpoints, and service startup. |
| [di](https://github.com/HnH/di) | 3 | 1 | 2021/10/13 | 1 month ago | DI container library that is focused on clean API and flexibility. |
**[⬆ back to top](#awesome-go-info)**

## DevOps Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [kubernetes](https://github.com/kubernetes/kubernetes) | 85116 | 3279 | 2014/06/06 | 1 week ago | Container Cluster Manager from Google. |
| [moby](https://github.com/moby/moby) | 62137 | 3038 | 2013/01/18 | 1 week ago | Collaborative project for the container ecosystem to assemble container-based systems. |
| [traefik](https://github.com/traefik/traefik) | 36694 | 691 | 2015/09/13 | 1 week ago | Reverse proxy and load balancer with support for multiple backends. |
| [gitea](https://github.com/go-gitea/gitea) | 28137 | 484 | 2016/11/01 | 1 week ago | Fork of Gogs, entirely community driven. |
| [vegeta](https://github.com/tsenart/vegeta) | 18994 | 318 | 2013/08/13 | 4 months ago | HTTP load testing tool and library. It's over 9000! |
| [packer](https://github.com/hashicorp/packer) | 13488 | 481 | 2013/03/23 | 1 week ago | Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. |
| [hey](https://github.com/rakyll/hey) | 12810 | 172 | 2016/09/02 | 1 month ago | Hey is a tiny program that sends some load to a web application. |
| [webhook](https://github.com/adnanh/webhook) | 7401 | 148 | 2015/01/12 | 1 month ago | Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server. |
| [gvm](https://github.com/moovweb/gvm) | 6985 | 154 | 2011/12/03 | 1 week ago | GVM provides an interface to manage Go versions. |
| [gaia](https://github.com/gaia-pipeline/gaia) | 4588 | 104 | 2017/12/28 | 1 month ago | Build powerful pipelines in any programming language. |
| [gox](https://github.com/mitchellh/gox) | 4187 | 76 | 2013/11/17 | 11 months ago | Dead simple, no frills Go cross compile tool. |
| [ddosify](https://github.com/ddosify/ddosify) | 3279 | 27 | 2021/08/04 | 2 weeks ago | High-performance load testing tool, written in Golang. |
| [bosun](https://github.com/bosun-monitor/bosun) | 3244 | 146 | 2013/11/15 | 4 months ago | Time Series Alerting Framework. |
| [bombardier](https://github.com/codesenberg/bombardier) | 3016 | 67 | 2016/05/29 | 1 month ago | Fast cross-platform HTTP benchmarking tool. |
| [pomerium](https://github.com/pomerium/pomerium) | 2926 | 37 | 2019/01/01 | 1 week ago | Pomerium is an identity-aware access proxy. |
| [script](https://github.com/bitfield/script) | 2009 | 34 | 2019/04/20 | 1 month ago | Making it easy to write shell-like scripts in Go for DevOps and system administration tasks. |
| [kala](https://github.com/ajvb/kala) | 1750 | 63 | 2015/03/19 | 3 weeks ago | Simplistic, modern, and performant job scheduler. |
| [fac](https://github.com/mkchoi212/fac) | 1742 | 31 | 2017/12/29 | 2 years ago | Command-line user interface to fix git merge conflicts. |
| [goxc](https://github.com/laher/goxc) | 1675 | 49 | 2013/02/11 | 2 years ago | build tool for Go, with a focus on cross-compiling and packaging. |
| [statusok](https://github.com/sanathp/statusok) | 1526 | 50 | 2015/08/26 | 6 months ago | Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected. |
| [s3gof3r](https://github.com/rlmcpherson/s3gof3r) | 1116 | 33 | 2013/08/02 | 5 months ago | Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3. |
| [go-selfupdate](https://github.com/sanbornm/go-selfupdate) | 918 | 31 | 2013/11/13 | 4 months ago | Enable your Go applications to self update. |
| [s5cmd](https://github.com/peak/s5cmd) | 916 | 24 | 2016/11/16 | 1 week ago | Blazing fast S3 and local filesystem execution tool. |
| [skm](https://github.com/TimothyYe/skm) | 750 | 20 | 2017/10/11 | 1 week ago | SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily! |
| [scaleway-cli](https://github.com/scaleway/scaleway-cli) | 728 | 33 | 2015/03/20 | 1 week ago | Manage BareMetal Servers from Command Line (as easily as with Docker). |
| [ghorg](https://github.com/gabrie30/ghorg) | 698 | 17 | 2018/03/29 | 3 weeks ago | Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, Gitea, and Bitbucket. |
| [utask](https://github.com/ovh/utask) | 600 | 29 | 2019/11/05 | 3 weeks ago | Automation engine that models and executes business processes declared in yaml. |
| [kool](https://github.com/kool-dev/kool) | 585 | 11 | 2020/07/06 | 1 month ago | Command line tool for managing Docker environments as an easy way. |
| [cassowary](https://github.com/rogerwelin/cassowary) | 555 | 5 | 2019/08/25 | 2 months ago | Modern cross-platform HTTP load-testing tool written in Go. |
| [aurora](https://github.com/xuri/aurora) | 549 | 31 | 2016/10/09 | 5 months ago | Cross-platform web-based Beanstalkd queue server console. |
| [govvv](https://github.com/ahmetb/govvv) | 519 | 10 | 2016/08/02 | 2 years ago | “go build” wrapper to easily add version information into Go binaries. |
| [kwatch](https://github.com/abahmed/kwatch) | 345 | 7 | 2021/11/20 | 1 week ago | Monitor & detect crashes in your Kubernetes(K8s) cluster instantly. |
| [gonative](https://github.com/inconshreveable/gonative) | 328 | 8 | 2014/05/01 | 5 years ago | Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages. |
| [pewpew](https://github.com/bengadbois/pewpew) | 313 | 10 | 2016/10/12 | 1 month ago | Flexible HTTP command line stress tester. |
| [trubka](https://github.com/xitonix/trubka) | 313 | 13 | 2019/07/05 | 1 month ago | A CLI tool to manage and troubleshoot Apache Kafka clusters with the ability of generically publishing/consuming protocol buffer and plain text events to/from Kafka. |
| [jenkins-cli](https://github.com/jenkins-zh/jenkins-cli) | 307 | 13 | 2019/06/21 | 4 weeks ago | Jenkins CLI allows you manage your Jenkins as an easy way. |
| [mora](https://github.com/emicklei/mora) | 298 | 25 | 2013/07/12 | 10 months ago | REST server for accessing MongoDB documents and meta data. |
| [lstags](https://github.com/ivanilves/lstags) | 286 | 11 | 2017/08/15 | 6 months ago | Tool and API to sync Docker images across different registries. |
| [balerter](https://github.com/balerter/balerter) | 248 | 7 | 2019/12/30 | 1 month ago | A self-hosted script-based alerting manager. |
| [manssh](https://github.com/xwjdsh/manssh) | 240 | 5 | 2017/10/08 | 2 weeks ago | manssh is a command line tool for managing your ssh alias config easily. |
| [dogo](https://github.com/liudng/dogo) | 239 | 19 | 2014/11/19 | 2 years ago | Monitoring changes in the source file and automatically compile and run (restart). |
| [godbg](https://github.com/sirnewton01/godbg) | 225 | 17 | 2013/08/09 | 3 years ago | Web-based gdb front-end application. |
| [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) | 223 | 9 | 2017/03/03 | 2 months ago | Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`. |
| [blast](https://github.com/dave/blast) | 203 | 5 | 2017/10/21 | 4 years ago | A simple tool for API load testing and batch jobs. |
| [gobrew](https://github.com/cryptojuice/gobrew) | 188 | 5 | 2013/11/13 | 1 year ago | gobrew lets you easily switch between multiple versions of go. |
| [terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi) | 186 | 11 | 2017/10/17 | 3 weeks ago | Terraform provider plugin that dynamically configures itself at runtime based on an OpenAPI document (formerly known as swagger file) containing the definitions of the APIs exposed. |
| [abbreviate](https://github.com/dnnrly/abbreviate) | 176 | 5 | 2018/11/23 | 4 months ago | abbreviate is a tool turning long strings in to shorter ones with configurable seperaters, for example to embed branch names in to deployment stack IDs. |
| [ostent](https://github.com/ostrost/ostent) | 172 | 7 | 2014/03/31 | 3 years ago | collects and displays system metrics and optionally relays to Graphite and/or InfluxDB. |
| [kcli](https://github.com/cswank/kcli) | 167 | 6 | 2017/03/25 | 2 years ago | Command line tool for inspecting kafka topics/partitions/messages. |
| [grapes](https://github.com/yaronsumel/grapes) | 153 | 6 | 2016/09/01 | 1 year ago | Lightweight tool designed to distribute commands over ssh with ease. |
| [winrm-cli](https://github.com/masterzen/winrm-cli) | 139 | 5 | 2016/05/23 | 1 month ago | Cli tool to remotely execute commands on Windows machines. |
| [dockerfile-generator](https://github.com/ozankasikci/dockerfile-generator) | 121 | 5 | 2019/08/14 | 2 years ago | A go library and an executable that produces valid Dockerfiles using various input channels. |
| [drone-scp](https://github.com/appleboy/drone-scp) | 93 | 3 | 2016/10/16 | 3 months ago | Copy files and artifacts via SSH using a binary, docker or Drone CI. |
| [go-furnace](https://github.com/go-furnace/go-furnace) | 86 | 2 | 2016/10/09 | 3 months ago | Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean. |
| [s3-proxy](https://github.com/oxyno-zeta/s3-proxy) | 79 | 2 | 2019/09/22 | 1 week ago | S3 Proxy with GET, PUT and DELETE methods and authentication (OpenID Connect and Basic Auth). |
| [dropship](https://github.com/ChrisMcKenzie/dropship) | 57 | 3 | 2015/09/03 | 3 years ago | Tool for deploying code via cdn. |
| [drone-jenkins](https://github.com/appleboy/drone-jenkins) | 33 | 3 | 2016/10/15 | 4 months ago | Trigger downstream Jenkins jobs using a binary, docker or Drone CI. |
| [rodent](https://github.com/alouche/rodent) | 32 | 3 | 2014/06/01 | 4 years ago | Rodent helps you manage Go versions, projects and track dependencies. |
| [awsenv](https://github.com/soniah/awsenv) | 28 | 2 | 2015/08/05 | 3 years ago | Small binary that loads Amazon (AWS) environment variables for a profile. |
| [docker-go-mingw](https://github.com/x1unix/docker-go-mingw) | 27 | 3 | 2020/09/16 | 1 month ago | Docker image for building Go binaries for Windows with MinGW toolchain. |
| [lwc](https://github.com/timdp/lwc) | 26 | 4 | 2018/04/22 | 1 year ago | A live-updating version of the UNIX wc command. |
| [depcharge](https://github.com/centerorbit/depcharge) | 19 | 3 | 2018/07/25 | 1 month ago | Helps orchestrating the execution of commands across the many dependencies in larger projects. |
| [httpref](https://github.com/dnnrly/httpref) | 18 | 3 | 2020/01/10 | 3 months ago | httpref is a handy CLI reference for HTTP methods, status codes, headers, and TCP and UDP ports. |
| [sg](https://github.com/ChristopherRabotin/sg) | 6 | 2 | 2015/08/19 | 5 years ago | Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response. |
| [aptly-fork](https://github.com/smira/aptly-fork) | 4 | 0 | 2019/07/04 | 2 years ago | aptly is a Debian repository management tool. |
**[⬆ back to top](#awesome-go-info)**

## HTTP Clients
        
*Libraries for making HTTP requests.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [resty](https://github.com/go-resty/resty) | 5588 | 90 | 2015/08/28 | 1 month ago | Simple HTTP and REST client for Go inspired by Ruby rest-client. |
| [heimdall](https://github.com/gojek/heimdall) | 2173 | 52 | 2018/01/19 | 2 weeks ago | An enhanced http client with retry and hystrix capabilities. |
| [grequests](https://github.com/levigross/grequests) | 1857 | 35 | 2015/06/11 | 1 year ago | A Go "clone" of the great and famous Requests library. |
| [sling](https://github.com/dghubble/sling) | 1411 | 29 | 2015/04/02 | 3 months ago | Sling is a Go HTTP client library for creating and sending API requests. |
| [go-retryablehttp](https://github.com/hashicorp/go-retryablehttp) | 1095 | 246 | 2015/12/07 | 1 month ago | Retryable HTTP client in Go. |
| [gentleman](https://github.com/h2non/gentleman) | 935 | 20 | 2016/02/21 | 1 year ago | Full-featured plugin-driven HTTP client library. |
| [pester](https://github.com/sethgrid/pester) | 570 | 6 | 2015/05/20 | 1 month ago | Go HTTP client calls with retries, backoff, and concurrency. |
| [requests](https://github.com/carlmjohnson/requests) | 270 | 8 | 2021/05/20 | 3 weeks ago | HTTP requests for Gophers. Uses context.Context and doesn't hide the underlying net/http.Client, making it compatible with standard Go APIs. Also includes testing tools. |
| [request](https://github.com/monaco-io/request) | 181 | 11 | 2020/03/25 | 1 month ago | HTTP client for golang. If you have experience about axios or requests, you will love it. No 3rd dependency. |
| [rq](https://github.com/ddo/rq) | 40 | 3 | 2017/12/26 | 2 years ago | A nicer interface for golang stdlib HTTP client. |
| [go-http-client](https://github.com/bozd4g/go-http-client) | 34 | 2 | 2019/12/14 | 9 months ago | Make http calls simply and easily. |
| [httpretry](https://github.com/ybbus/httpretry) | 16 | 3 | 2020/02/05 | 2 years ago | Enriches the default go HTTP client with retry functionality. |
| [go-req](https://github.com/wenerme/go-req) | 13 | 1 | 2021/07/11 | 5 months ago | Declarative golang HTTP client. |
| [httpc](https://github.com/valord577/httpc) | 4 | 1 | 2021/08/11 | 2 months ago | A customizable and simple HTTP client library. Only depend on the stdlib HTTP client. |
**[⬆ back to top](#awesome-go-info)**

## Microsoft Excel
        
*Libraries for working with Microsoft Excel.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [excelize](https://github.com/qax-os/excelize) | 10776 | 210 | 2016/08/29 | 1 week ago | Golang library for reading and writing Microsoft Excel&trade; (XLSX) files. |
| [xlsx](https://github.com/tealeg/xlsx) | 5225 | 171 | 2011/06/28 | 2 weeks ago | Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs. |
| [xlsx](https://github.com/plandem/xlsx) | 147 | 13 | 2017/08/26 | 1 year ago | Fast and safe way to read/update your existing Microsoft Excel files in Go programs. |
| [go-excel](https://github.com/szyhf/go-excel) | 134 | 3 | 2017/09/03 | 2 months ago | A simple and light reader to read a relate-db-like excel as a table. |
| [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) | 18 | 2 | 2017/03/13 | 3 years ago | Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files. |
**[⬆ back to top](#awesome-go-info)**

## Middlewares
        
**[⬆ back to top](#awesome-go-info)**

## Other Software
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [croc](https://github.com/schollz/croc) | 18348 | 231 | 2017/10/17 | 1 week ago | Easily and securely send files or folders from one computer to another. |
| [restic](https://github.com/restic/restic) | 15586 | 240 | 2014/04/27 | 1 week ago | De-duplicating backup program. |
| [goreplay](https://github.com/buger/goreplay) | 15186 | 467 | 2013/05/30 | 1 week ago | Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time. |
| [seaweedfs](https://github.com/chrislusf/seaweedfs) | 13777 | 529 | 2014/07/14 | 1 week ago | Fast, Simple and Scalable Distributed File System with O(1) disk seek. |
| [comcast](https://github.com/tylertreat/comcast) | 7754 | 151 | 2014/11/12 | 1 month ago | Simulate bad network connections. |
| [toxiproxy](https://github.com/Shopify/toxiproxy) | 7745 | 391 | 2014/09/04 | 2 weeks ago | Proxy to simulate network and system conditions for automated tests. |
| [confd](https://github.com/kelseyhightower/confd) | 7712 | 251 | 2013/10/01 | 4 months ago | Manage local application configuration files using templates and data from etcd or consul. |
| [liteide](https://github.com/visualfc/liteide) | 6712 | 371 | 2012/11/19 | 2 weeks ago | LiteIDE is a simple, open source, cross-platform Go IDE. |
| [drive](https://github.com/odeke-em/drive) | 6306 | 186 | 2014/11/03 | 1 year ago | Google Drive client for the commandline. |
| [nes](https://github.com/fogleman/nes) | 4974 | 145 | 2015/03/02 | 8 months ago | Nintendo Entertainment System (NES) emulator written in Go. |
| [duplicacy](https://github.com/gilbertchen/duplicacy) | 4036 | 95 | 2016/02/23 | 1 month ago | A cross-platform network and cloud backup tool based on the idea of lock-free deduplication. |
| [scc](https://github.com/boyter/scc) | 3121 | 28 | 2018/03/01 | 1 week ago | Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates. |
| [mylg](https://github.com/mehrdadrad/mylg) | 2546 | 111 | 2016/06/21 | 1 year ago | Command Line Network Diagnostic tool written in Go. |
| [goboy](https://github.com/Humpheh/goboy) | 2417 | 44 | 2017/08/20 | 1 year ago | Nintendo Game Boy Color emulator written in Go. |
| [sup](https://github.com/pressly/sup) | 2318 | 68 | 2015/02/23 | 3 weeks ago | Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers. |
| [lgo](https://github.com/yunabe/lgo) | 2229 | 48 | 2017/10/05 | 1 year ago | Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility. |
| [circuit](https://github.com/gocircuit/circuit) | 1930 | 138 | 2014/04/10 | 1 year ago | Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications. |
| [snap](https://github.com/intelsdi-x/snap) | 1798 | 142 | 2014/08/13 | 3 years ago | Powerful telemetry framework. |
| [borg](https://github.com/ok-borg/borg) | 1532 | 41 | 2016/09/10 | 4 years ago | Terminal based search engine for bash snippets. |
| [community](https://github.com/documize/community) | 1471 | 55 | 2016/04/29 | 1 month ago | Modern wiki software that integrates data from SaaS tools. |
| [blocky](https://github.com/0xERR0R/blocky) | 1310 | 22 | 2019/11/06 | 1 week ago | Fast and lightweight DNS proxy as ad-blocker for local network with many features. |
| [Go-Package-Store](https://github.com/shurcooL/Go-Package-Store) | 886 | 22 | 2014/01/24 | 1 year ago | App that displays updates for the Go packages in your GOPATH. |
| [shell2http](https://github.com/msoap/shell2http) | 880 | 23 | 2015/03/11 | 3 months ago | Executing shell commands via http server (for prototyping or remote control). |
| [vflow](https://github.com/EdgeCast/vflow) | 861 | 86 | 2017/02/24 | 1 week ago | High-performance, scalable and reliable IPFIX, sFlow and Netflow collector. |
| [peg](https://github.com/pointlander/peg) | 839 | 27 | 2010/04/25 | 5 months ago | Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator. |
| [plik](https://github.com/root-gg/plik) | 829 | 32 | 2015/04/19 | 1 week ago | Plik is a temporary file upload system (Wetransfer like) in Go. |
| [leaps](https://github.com/Jeffail/leaps) | 714 | 29 | 2014/06/19 | 11 months ago | Pair programming service using Operational Transforms. |
| [gfile](https://github.com/Antonito/gfile) | 644 | 12 | 2019/03/08 | 11 months ago | Securely transfer files between two computers, without any third party, over WebRTC. |
| [guora](https://github.com/argoodies/guora) | 576 | 15 | 2020/08/13 | 1 year ago | A self-hosted Quora like web application written in Go. |
| [guora](https://github.com/meloalright/guora) | 573 | 14 | 2020/08/13 | 1 year ago | A self-hosted Quora like web application written in Go. |
| [gebug](https://github.com/moshebe/gebug) | 560 | 5 | 2020/07/20 | 1 week ago | A tool that makes debugging of Dockerized Go applications super easy by enabling Debugger and Hot-Reload features, seamlessly. |
| [gocc](https://github.com/goccmack/gocc) | 512 | 22 | 2015/06/05 | 2 months ago | Gocc is a compiler kit for Go written in Go. |
| [mockingjay-server](https://github.com/quii/mockingjay-server) | 504 | 9 | 2015/04/04 | 1 year ago | Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests. |
| [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) | 443 | 20 | 2015/10/08 | 6 months ago | Video streaming torrent client. |
| [ipe](https://github.com/dimiro1/ipe) | 344 | 18 | 2015/01/13 | 10 months ago | Open source Pusher server implementation compatible with Pusher client libraries written in GO. |
| [IDE](https://github.com/thestrukture/IDE) | 325 | 19 | 2017/09/09 | 6 months ago | Browser accessible IDE. Designed for Go with Go. |
| [tcpprobe](https://github.com/mehrdadrad/tcpprobe) | 318 | 9 | 2020/10/26 | 11 months ago | TCP tool for network performance and path monitoring, including socket statistics. |
| [woke](https://github.com/get-woke/woke) | 302 | 7 | 2020/08/31 | 1 month ago | Detect non-inclusive language in your source code. |
| [wellington](https://github.com/wellington/wellington) | 300 | 13 | 2014/12/08 | 1 year ago | Sass project management tool, extends the language with sprite functions (like Compass). |
| [cherry](https://github.com/rafael-santiago/cherry) | 269 | 13 | 2015/10/24 | 4 years ago | Tiny webchat server in Go. |
| [Neo-cowsay](https://github.com/Code-Hex/Neo-cowsay) | 176 | 7 | 2016/11/05 | 2 months ago | 🐮 cowsay is reborn. for a New Era. |
| [tcpdog](https://github.com/mehrdadrad/tcpdog) | 173 | 10 | 2020/12/30 | 6 months ago | eBPF based TCP observability. |
| [joincap](https://github.com/assafmo/joincap) | 171 | 9 | 2018/05/31 | 11 months ago | Command-line utility for merging multiple pcap files together. |
| [orbit](https://github.com/gulien/orbit) | 162 | 8 | 2017/05/13 | 1 year ago | A simple tool for running commands and generating files from templates. |
| [vaku](https://github.com/lingrino/vaku) | 128 | 3 | 2018/04/24 | 2 weeks ago | CLI & API for folder-based functions in Vault like copy, move, and search. |
| [dp](https://github.com/scryinfo/dp) | 83 | 12 | 2018/12/12 | 1 month ago | Through SDK for data exchange with blockchain, developers can get easy access to DAPP development. |
| [boxed](https://github.com/tejo/boxed) | 75 | 3 | 2015/04/18 | 3 years ago | Dropbox based blog engine. |
| [crawley](https://github.com/s0rg/crawley) | 45 | 1 | 2021/10/27 | 1 month ago | Web scraper/crawler for cli. |
| [naclpipe](https://github.com/unix4fun/naclpipe) | 22 | 6 | 2015/05/05 | 3 years ago | Simple NaCL EC25519 based crypto pipe tool written in Go. |
| [term-quiz](https://github.com/crazcalm/term-quiz) | 21 | 1 | 2017/12/26 | 3 years ago | Quizzes for your terminal. |
| [snitch](https://github.com/lucasgomide/snitch) | 15 | 1 | 2017/04/06 | 3 years ago | Simple way to notify your team and many tools when someone has deployed any application via Tsuru. |
| [GoDocTooltip](https://github.com/diankong/GoDocTooltip) | 12 | 3 | 2016/01/21 | 1 month ago | Chrome extension for Go Doc sites, which shows function description as tooltip at function list. |
| [protoncheck](https://github.com/servusdei2018/protoncheck) | 3 | 1 | 2021/12/26 | 1 month ago | ProtonMail module for waybar/polybar/yabar/i3blocks. |
| [hoofli](https://github.com/dnnrly/hoofli) | 1 | 2 | 2021/04/23 | 4 months ago | Generate PlantUML diagrams from Chrome or Firefox network inspections. |
**[⬆ back to top](#awesome-go-info)**

## Project Layout
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [project-layout](https://github.com/golang-standards/project-layout) | 29328 | 549 | 2017/09/09 | 2 weeks ago | Set of common historical and emerging project layout patterns in the Go ecosystem. Note: despite the org-name they do not represent official golang standards, see [this issue](https://github.com/golang-standards/project-layout/issues/117) for more information. Nonetheless, some may find the layout useful. |
| [service](https://github.com/ardanlabs/service) | 2168 | 96 | 2017/11/20 | 1 week ago | A [starter kit](https://github.com/ardanlabs/service/wiki) for building production grade scalable web service applications. |
| [modern-go-application](https://github.com/sagikazarmark/modern-go-application) | 1130 | 26 | 2018/09/14 | 1 month ago | Go application boilerplate and example applying modern practices. |
| [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) | 497 | 10 | 2016/12/18 | 1 year ago | A Go application boilerplate template for quick starting projects following production best practices. |
| [seed](https://github.com/golang-templates/seed) | 240 | 5 | 2020/04/30 | 2 weeks ago | Go application GitHub repository template. |
| [go-starter](https://github.com/allaboutapps/go-starter) | 120 | 11 | 2020/05/08 | 1 week ago | An opinionated production-ready RESTful JSON backend template, highly integrated with VSCode DevContainers. |
| [go-todo-backend](https://github.com/Fs02/go-todo-backend) | 117 | 4 | 2020/06/25 | 1 month ago | Go Todo Backend example using modular project layout for product microservice. |
| [scaffold](https://github.com/catchplay/scaffold) | 106 | 5 | 2018/12/11 | 3 years ago | Scaffold generates a starter Go project layout. Lets you focus on business logic implemented. |
| [go-sample](https://github.com/zitryss/go-sample) | 96 | 1 | 2019/01/24 | 3 years ago | A sample layout for Go application projects with the real code. |
| [gobase](https://github.com/wajox/gobase) | 19 | 3 | 2020/12/15 | 4 months ago | A simple skeleton for golang application with basic setup for real golang application. |
| [go-project-layout](https://github.com/wangyoucao577/go-project-layout) | 14 | 2 | 2019/10/06 | 9 months ago | Set of practices and discussions on how to structure Go project layout. |
| [inizio](https://github.com/insidieux/inizio) | 9 | 1 | 2021/03/02 | 9 months ago | Golang project layout generator with plugins. |
**[⬆ back to top](#awesome-go-info)**

## Routers
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [mux](https://github.com/gorilla/mux) | 15955 | 313 | 2012/10/02 | 4 weeks ago | Powerful URL router and dispatcher for golang. |
| [httprouter](https://github.com/julienschmidt/httprouter) | 13652 | 326 | 2013/12/05 | 4 weeks ago | High performance router. Use this and the standard http handlers to form a very high performance web framework. |
| [chi](https://github.com/go-chi/chi) | 10874 | 180 | 2015/10/15 | 4 weeks ago | Small, fast and expressive HTTP router built on net/context. |
| [web](https://github.com/gocraft/web) | 1458 | 59 | 2013/11/16 | 1 year ago | Mux and middleware package in Go. |
| [bone](https://github.com/go-zoo/bone) | 1279 | 36 | 2014/11/19 | 2 years ago | Lightning Fast HTTP Multiplexer. |
| [goji](https://github.com/goji/goji) | 895 | 42 | 2015/11/16 | 2 years ago | Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`. |
| [fasthttprouter](https://github.com/buaazp/fasthttprouter) | 870 | 34 | 2015/12/13 | 2 years ago | High performance router forked from `httprouter`. The first router fit for `fasthttp`. |
| [httptreemux](https://github.com/dimfeld/httptreemux) | 522 | 23 | 2014/05/14 | 3 months ago | High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter. |
| [gorouter](https://github.com/xujiajun/gorouter) | 514 | 16 | 2018/01/29 | 2 years ago | A simple and fast HTTP router for Go. |
| [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) | 426 | 28 | 2015/10/27 | 5 months ago | An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs. |
| [lars](https://github.com/go-playground/lars) | 385 | 15 | 2015/12/24 | 2 years ago | Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks. |
| [siesta](https://github.com/VividCortex/siesta) | 352 | 28 | 2014/09/23 | 9 months ago | Composable framework to write middleware and handlers. |
| [vestigo](https://github.com/husobee/vestigo) | 266 | 15 | 2015/09/22 | 1 year ago | Performant, stand-alone, HTTP compliant URL Router for go web applications. |
| [router](https://github.com/gowww/router) | 160 | 7 | 2017/05/25 | 1 year ago | Lightning fast HTTP router fully compatible with the net/http.Handler interface. |
| [pure](https://github.com/go-playground/pure) | 123 | 7 | 2016/09/23 | 1 year ago | Is a lightweight HTTP router that sticks to the std "net/http" implementation. |
| [alien](https://github.com/gernest/alien) | 121 | 4 | 2016/01/30 | 2 years ago | Lightweight and fast http router from outer space. |
| [gorouter](https://github.com/vardius/gorouter) | 106 | 6 | 2016/07/14 | 1 month ago | GoRouter is a Server/API micro framework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`. |
| [violetear](https://github.com/nbari/violetear) | 104 | 4 | 2015/06/19 | 8 months ago | Go HTTP router. |
| [Bxog](https://github.com/claygod/Bxog) | 102 | 9 | 2016/05/19 | 1 year ago | Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters. |
| [xmux](https://github.com/rs/xmux) | 94 | 6 | 2015/12/14 | 4 years ago | High performance muxer based on `httprouter` with `net/context` support. |
| [bellt](https://github.com/GuilhermeCaruso/bellt) | 53 | 6 | 2019/02/21 | 1 year ago | A simple Go HTTP router. |
| [ngamux](https://github.com/ngamux/ngamux) | 46 | 1 | 2021/08/22 | 3 months ago | Simple HTTP router for Go. |
| [fastrouter](https://github.com/razonyang/fastrouter) | 20 | 2 | 2017/11/01 | 4 years ago | a fast, flexible HTTP router written in Go. |
| [route](https://github.com/goroute/route) | 7 | 4 | 2019/07/06 | 2 years ago | Simple yet powerful HTTP request multiplexer. |
**[⬆ back to top](#awesome-go-info)**

## Standard CLI
        
**[⬆ back to top](#awesome-go-info)**

## Strings
        
*Libraries for working with strings.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [xstrings](https://github.com/huandu/xstrings) | 983 | 25 | 2015/01/06 | 1 month ago | Collection of useful string functions ported from other languages. |
| [sttr](https://github.com/abhimanyu003/sttr) | 346 | 5 | 2021/09/18 | 1 month ago | cross-platform, cli app to perform various operations on string. |
| [strutil](https://github.com/ozgio/strutil) | 151 | 4 | 2018/08/16 | 3 months ago | String utilities. |
| [stringy](https://github.com/gobeam/stringy) | 106 | 4 | 2020/04/03 | 3 months ago | String manipulation library to convert string to camel case, snake case, kebab case / slugify etc. |
| [bexp](https://github.com/mkungla/bexp) | 4 | 1 | 2020/12/15 | 4 months ago | Go implementation of Brace Expansion mechanism to generate arbitrary strings. |
**[⬆ back to top](#awesome-go-info)**

## Uncategorized
        
*These libraries were placed here because none of the other categories seemed to fit.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gopsutil](https://github.com/shirou/gopsutil) | 7292 | 209 | 2014/04/18 | 1 week ago | Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc). |
| [archiver](https://github.com/mholt/archiver) | 3461 | 51 | 2016/04/08 | 2 weeks ago | Library and command for making and extracting .zip and .tar.gz archives. |
| [gofakeit](https://github.com/brianvoe/gofakeit) | 2267 | 17 | 2015/04/24 | 1 week ago | Random data generator written in go. |
| [gatus](https://github.com/TwiN/gatus) | 2134 | 22 | 2019/09/04 | 1 week ago | Automated service health dashboard. |
| [gosms](https://github.com/haxpax/gosms) | 1367 | 57 | 2015/01/23 | 1 year ago | Your own local SMS gateway in Go that can be used to send SMS. |
| [base64Captcha](https://github.com/mojocn/base64Captcha) | 1340 | 48 | 2017/12/12 | 2 months ago | Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha. |
| [go-resiliency](https://github.com/eapache/go-resiliency) | 1319 | 26 | 2014/11/29 | 5 months ago | Resiliency patterns for golang. |
| [go-commons-pool](https://github.com/jolestar/go-commons-pool) | 1006 | 49 | 2015/12/28 | 9 months ago | Generic object pool for Golang. |
| [llvm](https://github.com/llir/llvm) | 850 | 32 | 2014/09/19 | 3 weeks ago | Library for interacting with LLVM IR in pure Go. |
| [shortid](https://github.com/teris-io/shortid) | 738 | 10 | 2016/01/04 | 1 year ago | Distributed generation of super short, unique, non-sequential, URL friendly IDs. |
| [health](https://github.com/alexliesenfeld/health) | 483 | 6 | 2021/07/02 | 1 month ago | A simple and flexible health check library for Go. |
| [health](https://github.com/dimiro1/health) | 427 | 6 | 2016/03/08 | 2 years ago | Easy to use, extensible health check library. |
| [stateless](https://github.com/qmuntal/stateless) | 398 | 10 | 2019/09/11 | 4 months ago | A fluent library for creating state machines. |
| [banner](https://github.com/dimiro1/banner) | 378 | 6 | 2016/03/25 | 1 year ago | Add beautiful banners into your Go applications. |
| [go-conv](https://github.com/cstockton/go-conv) | 375 | 8 | 2016/10/11 | 5 months ago | Package conv provides fast and intuitive conversions across Go types. |
| [gountries](https://github.com/pariz/gountries) | 332 | 11 | 2016/01/13 | 2 months ago | Package that exposes country and subdivision data. |
| [shoutrrr](https://github.com/containrrr/shoutrrr) | 296 | 8 | 2019/04/11 | 2 weeks ago | Notification library providing easy access to various messaging services like slack, mattermost, gotify and smtp among others. |
| [ffmt](https://github.com/go-ffmt/ffmt) | 250 | 7 | 2015/02/14 | 2 months ago | Beautify data display for Humans. |
| [lk](https://github.com/hyperboloide/lk) | 233 | 7 | 2016/07/14 | 1 year ago | A simple licensing library for golang. |
| [antch](https://github.com/antchfx/antch) | 225 | 16 | 2017/09/28 | 1 year ago | A fast, powerful and extensible web crawling & scraping framework. |
| [healthcheck](https://github.com/etherlabsio/healthcheck) | 209 | 10 | 2017/08/18 | 8 months ago | An opinionated and concurrent health-check HTTP handler for RESTful services. |
| [battery](https://github.com/distatus/battery) | 201 | 5 | 2016/03/12 | 1 month ago | Cross-platform, normalized battery information library. |
| [bitio](https://github.com/icza/bitio) | 178 | 8 | 2016/05/31 | 3 weeks ago | Highly optimized bit-level Reader and Writer for Go. |
| [go-unarr](https://github.com/gen2brain/go-unarr) | 168 | 7 | 2015/11/01 | 2 weeks ago | Decompression library for RAR, TAR, ZIP and 7z archives. |
| [stats](https://github.com/go-playground/stats) | 158 | 3 | 2015/09/14 | 5 years ago | Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... |
| [turtle](https://github.com/hackebrot/turtle) | 131 | 2 | 2017/09/08 | 4 months ago | Emojis for Go. |
| [gommit](https://github.com/antham/gommit) | 96 | 4 | 2016/08/30 | 2 months ago | Analyze git commit messages to ensure they follow defined patterns. |
| [indigo](https://github.com/osamingo/indigo) | 93 | 1 | 2016/08/31 | 1 year ago | Distributed unique ID generator of using Sonyflake and encoded by Base58. |
| [gotoprom](https://github.com/cabify/gotoprom) | 92 | 93 | 2018/10/10 | 2 years ago | Type-safe metrics builder wrapper library for the official Prometheus client. |
| [captcha](https://github.com/steambap/captcha) | 91 | 5 | 2017/09/12 | 1 month ago | Package captcha provides an easy to use, unopinionated API for captcha generation. |
| [morse](https://github.com/alwindoss/morse) | 74 | 2 | 2018/08/15 | 5 months ago | Library to convert to and from morse code. |
| [persian](https://github.com/mavihq/persian) | 61 | 6 | 2017/10/16 | 8 months ago | Some utilities for Persian language in go. |
| [faker](https://github.com/pioz/faker) | 57 | 4 | 2020/07/22 | 4 months ago | Random fake data and struct generator for Go. |
| [pdfgen](https://github.com/hyperboloide/pdfgen) | 56 | 4 | 2015/11/30 | 4 years ago | HTTP service to generate PDF from Json requests. |
| [xkg](https://github.com/go-xkg/xkg) | 52 | 2 | 2015/01/05 | 7 years ago | X Keyboard Grabber. |
| [browscap_go](https://github.com/digitalcrab/browscap_go) | 40 | 6 | 2014/09/18 | 5 months ago | GoLang Library for [Browser Capabilities Project](https://browscap.org/). |
| [datacounter](https://github.com/miolini/datacounter) | 37 | 1 | 2015/10/14 | 2 years ago | Go counters for readers/writer/http.ResponseWriter. |
| [autoflags](https://github.com/artyom/autoflags) | 36 | 5 | 2014/05/15 | 9 months ago | Go package to automatically define command line flags from struct fields. |
| [url-shortener](https://github.com/pantrif/url-shortener) | 35 | 1 | 2018/06/04 | 3 years ago | A modern, powerful, and robust URL shortener microservice with mysql support. |
| [sandid](https://github.com/aofei/sandid) | 33 | 1 | 2018/06/12 | 3 months ago | Every grain of sand on earth has its own ID. |
| [gtree](https://github.com/ddddddO/gtree) | 31 | 2 | 2021/05/30 | 2 weeks ago | Provide CLI and Package for tree output and directories creation from Markdown or programmatically. |
| [gosh](https://github.com/osamingo/gosh) | 29 | 0 | 2018/05/25 | 1 year ago | Provide Go Statistics Handler, Struct, Measure Method. |
| [xdg](https://github.com/rkoesters/xdg) | 28 | 3 | 2013/12/15 | 7 months ago | FreeDesktop.org (xdg) Specs implemented in Go. |
| [metrics](https://github.com/pascaldekloe/metrics) | 23 | 1 | 2019/01/29 | 11 months ago | Library for metrics instrumentation and Prometheus exposition. |
| [shellwords](https://github.com/Wing924/shellwords) | 16 | 2 | 2017/09/28 | 4 years ago | A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell. |
| [anagent](https://github.com/mudler/anagent) | 14 | 3 | 2017/12/29 | 3 years ago | Minimalistic, pluggable Golang evloop/timer handler with dependency-injection. |
| [avgRating](https://github.com/kirillDanshin/avgRating) | 11 | 2 | 2017/08/05 | 4 years ago | Calculate average score and rating based on Wilson Score Equation. |
| [hostutils](https://github.com/Wing924/hostutils) | 10 | 2 | 2017/09/26 | 3 weeks ago | A golang library for packing and unpacking FQDNs list. |
| [numa](https://github.com/lrita/numa) | 6 | 4 | 2018/12/10 | 1 year ago | NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code. |
| [go-commandbus](https://github.com/lana/go-commandbus) | 4 | 24 | 2019/10/03 | 2 weeks ago | A slight and pluggable command-bus for Go. |
| [varint](https://github.com/chmike/varint) | 2 | 1 | 2021/11/30 | 2 months ago | A faster varying length integer encoder/decoder than the one provided in the standard library. |
| [faker](https://github.com/neotoolkit/faker) | 2 | 0 | 2022/01/23 | 1 week ago | Fake data generator. |
| [basexx](https://github.com/bobg/basexx) | 1 | 1 | 2019/06/08 | 4 months ago | Convert to, from, and between digit strings in various number bases. |
**[⬆ back to top](#awesome-go-info)**

# Audio and Music
        
*Libraries for manipulating audio.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [oto](https://github.com/hajimehoshi/oto) | 949 | 13 | 2017/05/04 | 1 month ago | A low-level library to play sound on multiple platforms. |
| [portaudio](https://github.com/gordonklaus/portaudio) | 491 | 12 | 2015/09/16 | 1 year ago | Go bindings for the PortAudio audio I/O library. |
| [music-theory](https://github.com/go-music-theory/music-theory) | 381 | 20 | 2016/03/17 | 1 year ago | Music theory models in Go. |
| [portmidi](https://github.com/rakyll/portmidi) | 263 | 11 | 2013/11/10 | 1 year ago | Go bindings for PortMidi. |
| [id3v2](https://github.com/bogem/id3v2) | 232 | 7 | 2016/05/15 | 2 weeks ago | ID3 decoding and encoding library for Go. |
| [GoAudio](https://github.com/DylanMeeus/GoAudio) | 205 | 7 | 2020/07/05 | 2 months ago | Native Go Audio Processing Library. |
| [flac](https://github.com/mewkiz/flac) | 197 | 11 | 2012/11/01 | 1 month ago | Native Go FLAC encoder/decoder with support for FLAC streams. |
| [malgo](https://github.com/gen2brain/malgo) | 163 | 6 | 2017/11/09 | 4 months ago | Mini audio library. |
| [mix](https://github.com/go-mix/mix) | 146 | 3 | 2016/01/03 | 1 year ago | Sequence-based Go-native audio mixer for music apps. |
| [gaad](https://github.com/Comcast/gaad) | 93 | 11 | 2016/07/11 | 4 months ago | Native Go AAC bitstream parser. |
| [minimp3](https://github.com/tosone/minimp3) | 73 | 4 | 2018/01/26 | 10 months ago | Lightweight MP3 decoder library. |
| [vorbis](https://github.com/mccoyst/vorbis) | 30 | 3 | 2013/07/12 | 2 years ago | "Native" Go Vorbis decoder (uses CGO, but has no dependencies). |
| [gosamplerate](https://github.com/dh1tw/gosamplerate) | 16 | 1 | 2016/11/20 | 1 year ago | libsamplerate bindings for go. |
**[⬆ back to top](#awesome-go-info)**

# Authentication and OAuth
        
*Libraries for implementing authentications schemes.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [casbin](https://github.com/casbin/casbin) | 11207 | 224 | 2017/04/08 | 2 weeks ago | Authorization library that supports access control models like ACL, RBAC, ABAC. |
| [oauth2](https://github.com/golang/oauth2) | 3988 | 101 | 2014/04/14 | 2 weeks ago | Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support. |
| [goth](https://github.com/markbates/goth) | 3499 | 60 | 2014/10/14 | 3 weeks ago | provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box. |
| [authboss](https://github.com/volatiletech/authboss) | 2988 | 56 | 2015/01/03 | 6 months ago | Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time. |
| [go-jose](https://github.com/square/go-jose) | 1866 | 58 | 2014/11/14 | 3 months ago | Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs. |
| [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) | 1866 | 86 | 2015/11/01 | 1 year ago | Standalone, specification-compliant,  OAuth2 server written in Golang. |
| [loginsrv](https://github.com/tarent/loginsrv) | 1844 | 53 | 2016/11/11 | 11 months ago | JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam. |
| [osin](https://github.com/openshift/osin) | 1712 | 71 | 2013/09/10 | 2 months ago | Golang OAuth2 server library. |
| [gologin](https://github.com/dghubble/gologin) | 1480 | 27 | 2015/06/23 | 1 week ago | chainable handlers for login with OAuth1 and OAuth2 authentication providers. |
| [gorbac](https://github.com/mikespook/gorbac) | 1246 | 63 | 2013/12/26 | 9 months ago | provides a lightweight role-based access control (RBAC) implementation in Golang. |
| [scs](https://github.com/alexedwards/scs) | 1030 | 27 | 2016/08/08 | 1 week ago | Session Manager for HTTP servers. |
| [paseto](https://github.com/o1egl/paseto) | 576 | 23 | 2018/01/23 | 4 months ago | Golang implementation of Platform-Agnostic Security Tokens (PASETO). |
| [permissions2](https://github.com/xyproto/permissions2) | 443 | 14 | 2014/11/19 | 4 weeks ago | Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt. |
| [go-guardian](https://github.com/shaj13/go-guardian) | 338 | 5 | 2020/05/14 | 1 month ago | Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication that supports LDAP, Basic, Bearer token and Certificate based authentication. |
| [jwt](https://github.com/cristalhq/jwt) | 316 | 12 | 2019/07/20 | 3 weeks ago | Safe, simple and fast JSON Web Tokens for Go. |
| [jwt](https://github.com/pascaldekloe/jwt) | 277 | 13 | 2018/03/21 | 3 months ago | Lightweight JSON Web Token (JWT) library. |
| [jeff](https://github.com/abraithwaite/jeff) | 234 | 4 | 2018/08/02 | 7 months ago | Simple, flexible, secure and idiomatic web session management with pluggable backends. |
| [httpauth](https://github.com/goji/httpauth) | 213 | 8 | 2014/05/26 | 4 months ago | HTTP Authentication middleware. |
| [jwt-auth](https://github.com/adam-hanna/jwt-auth) | 213 | 12 | 2016/07/05 | 6 months ago | JWT middleware for Golang http servers with many configuration options. |
| [branca](https://github.com/hako/branca) | 167 | 8 | 2018/01/09 | 1 year ago | Golang implementation of Branca Tokens. |
| [otpgen](https://github.com/RijulGulati/otpgen) | 116 | 2 | 2021/06/01 | 6 months ago | Library to generate TOTP/HOTP codes. |
| [sessionup](https://github.com/swithek/sessionup) | 114 | 7 | 2019/07/23 | 5 months ago | Simple, yet effective HTTP session management and identification package. |
| [session](https://github.com/icza/session) | 106 | 7 | 2016/02/08 | 6 months ago | Go session management for web servers (including support for Google App Engine - GAE). |
| [jwt](https://github.com/robbert229/jwt) | 97 | 9 | 2016/06/05 | 1 year ago | Clean and easy to use implementation of JSON Web Tokens (JWT). |
| [sjwt](https://github.com/brianvoe/sjwt) | 97 | 1 | 2019/06/20 | 2 years ago | Simple jwt generator and parser. |
| [rbac](https://github.com/zpatrick/rbac) | 90 | 3 | 2018/08/02 | 3 years ago | Minimalistic RBAC package for Go applications. |
| [sessions](https://github.com/adam-hanna/sessions) | 62 | 3 | 2017/04/29 | 1 year ago | Dead simple, highly performant, highly customizable sessions service for go http servers. |
| [securecookie](https://github.com/chmike/securecookie) | 55 | 5 | 2017/09/03 | 4 months ago | Efficient secure cookie encoding/decoding. |
| [go-email-normalizer](https://github.com/dimuska139/go-email-normalizer) | 37 | 1 | 2020/08/21 | 5 months ago | Golang library for providing a canonical representation of email address. |
| [otpgo](https://github.com/jltorresm/otpgo) | 27 | 3 | 2020/08/19 | 11 months ago | Time-Based One-Time Password (TOTP) and HMAC-Based One-Time Password (HOTP) library for Go. |
| [branca](https://github.com/essentialkaos/branca) | 27 | 4 | 2018/12/29 | 5 months ago | branca token [specification implementation](https://github.com/tuupola/branca-spec) for Golang 1.15+. |
| [scope](https://github.com/ThundR67/scope) | 19 | 1 | 2019/09/23 | 8 months ago | Easily Manage OAuth2 Scopes In Go. |
| [scope](https://github.com/SonicRoshan/scope) | 17 | 1 | 2019/09/23 | 8 months ago | Easily Manage OAuth2 Scopes In Go. |
| [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) | 10 | 2 | 2017/10/20 | 3 years ago | Go session management using the SessionGate Redis module. |
| [cookiestxt](https://github.com/mengzhuo/cookiestxt) | 9 | 1 | 2017/10/09 | 11 months ago | provides parser of cookies.txt file format. |
**[⬆ back to top](#awesome-go-info)**

# Blockchain
        
*Tools for building blockchains.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-ethereum](https://github.com/ethereum/go-ethereum) | 35264 | 2135 | 2013/12/26 | 1 week ago | Official Go implementation of the Ethereum protocol. |
| [tendermint](https://github.com/tendermint/tendermint) | 4625 | 264 | 2014/05/14 | 1 week ago | High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols. |
| [cosmos-sdk](https://github.com/cosmos/cosmos-sdk) | 3358 | 212 | 2016/02/06 | 1 week ago | A Framework for Building Public Blockchains in the Cosmos Ecosystem. |
| [gossamer](https://github.com/ChainSafe/gossamer) | 301 | 17 | 2019/01/28 | 1 week ago | A Go implementation of the Polkadot Host. |
| [solana-go](https://github.com/gagliardetto/solana-go) | 196 | 9 | 2021/06/29 | 1 week ago | Go library to interface with Solana JSON RPC and WebSocket interfaces. |
**[⬆ back to top](#awesome-go-info)**

# Bot Building
        
*Libraries for building and working with bots.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) | 3366 | 78 | 2015/06/25 | 1 month ago | Simple and clean Telegram bot client. |
| [olivia](https://github.com/olivia-ai/olivia) | 3190 | 83 | 2018/06/05 | 1 month ago | A chatbot built with an artificial neural network. |
| [telebot](https://github.com/tucnak/telebot) | 2280 | 55 | 2015/06/25 | 2 weeks ago | Telegram bot framework written in Go. |
| [kelp](https://github.com/stellar/kelp) | 845 | 53 | 2018/08/08 | 2 months ago | official trading and market-making bot for the [Stellar](https://www.stellar.org/) DEX. Works out-of-the-box, written in Golang, compatible with centralized exchanges and custom trading strategies. |
| [bot](https://github.com/go-chat-bot/bot) | 719 | 47 | 2015/09/22 | 2 weeks ago | IRC, Slack & Telegram bot written in Go. |
| [golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot) | 694 | 40 | 2017/05/14 | 1 week ago | A golang implementation of a console-based trading bot for cryptocurrency exchanges. |
| [slacker](https://github.com/shomali11/slacker) | 577 | 15 | 2017/05/20 | 2 weeks ago | Easy to use framework to create Slack bots. |
| [tbot](https://github.com/yanzay/tbot) | 323 | 11 | 2015/09/11 | 10 months ago | Telegram bot server with API similar to net/http. |
| [go-sarah](https://github.com/oklahomer/go-sarah) | 216 | 7 | 2016/11/06 | 1 month ago | Framework to build bot for desired chat services including LINE, Slack, Gitter and more. |
| [go-twitch-irc](https://github.com/gempir/go-twitch-irc) | 216 | 11 | 2017/03/23 | 2 weeks ago | Library to write bots for twitch. |
| [tenyks](https://github.com/kyleterry/tenyks) | 172 | 14 | 2012/08/26 | 2 years ago | Service oriented IRC bot using Redis and JSON for messaging. |
| [hanu](https://github.com/sbstjn/hanu) | 135 | 7 | 2016/09/16 | 8 months ago | Framework for writing Slack bots. |
| [go-tgbot](https://github.com/olebedev/go-tgbot) | 112 | 9 | 2016/12/11 | 3 years ago | Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware. |
| [slack-bot](https://github.com/innogames/slack-bot) | 78 | 4 | 2019/07/19 | 1 week ago | Ready to use Slack Bot for lazy developers: Custom commands, Jenkins, Jira, Bitbucket, Github... |
| [margelet](https://github.com/zhulik/margelet) | 67 | 4 | 2015/11/21 | 5 years ago | Framework for building Telegram bots. |
| [echotron](https://github.com/NicoNex/echotron) | 60 | 5 | 2019/07/22 | 2 weeks ago | An elegant and concurrent library for Telegram Bots in Go. |
| [ephemeral-roles](https://github.com/ewohltman/ephemeral-roles) | 54 | 5 | 2017/12/19 | 2 weeks ago | A Discord bot for managing ephemeral roles based upon voice channel member presence. |
| [slackscot](https://github.com/alexandre-normand/slackscot) | 49 | 2 | 2015/10/22 | 2 months ago | Another framework for building Slack bots. |
| [govkbot](https://github.com/nikepan/govkbot) | 38 | 3 | 2016/07/11 | 6 months ago | Simple Go [VK](https://vk.com) bot library. |
| [larry](https://github.com/ezeoleaf/larry) | 34 | 1 | 2020/11/16 | 3 weeks ago | Larry 🐦 is a really simple Twitter bot generator that tweets random repositories from Github built in Go. |
| [micha](https://github.com/onrik/micha) | 18 | 4 | 2016/04/14 | 8 months ago | Go Library for Telegram bot api. |
| [telego](https://github.com/mymmrac/telego) | 10 | 1 | 2021/06/27 | 1 week ago | Telegram Bot API library for Golang with full one-to-one API implementation. |
| [teleterm](https://github.com/alfiankan/teleterm) | 5 | 1 | 2020/12/31 | 3 months ago | Telegram Bot Exec Terminal Command. |
**[⬆ back to top](#awesome-go-info)**

# Build Automation
        
*Libraries and tools helping with build automation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [task](https://github.com/go-task/task) | 4578 | 61 | 2017/02/27 | 1 week ago | simple "Make" alternative. |
| [realize](https://github.com/oxequa/realize) | 4202 | 72 | 2016/07/12 | 9 months ago | Go build system with file watchers and live reload. Run, build and watch file changes with custom paths. |
| [mage](https://github.com/magefile/mage) | 2830 | 44 | 2017/09/20 | 3 weeks ago | Mage is a make/rake-like build tool using Go. |
| [mmake](https://github.com/tj/mmake) | 1598 | 29 | 2017/02/15 | 1 year ago | Modern Make. |
| [goyek](https://github.com/goyek/goyek) | 279 | 4 | 2020/10/11 | 2 weeks ago | Create build pipelines in Go. |
| [taskctl](https://github.com/taskctl/taskctl) | 153 | 7 | 2019/11/12 | 8 months ago | Concurrent task runner. |
| [1build](https://github.com/gopinath-langote/1build) | 134 | 8 | 2019/04/23 | 2 weeks ago | Command line tool to frictionlessly manage project-specific commands. |
| [gaper](https://github.com/maxcnunes/gaper) | 51 | 0 | 2018/06/16 | 1 month ago | Builds and restarts a Go project when it crashes or some watched file changes. |
| [anko](https://github.com/GuilhermeCaruso/anko) | 21 | 3 | 2021/03/02 | 10 months ago | Simple application watcher for multiple programming languages. |
**[⬆ back to top](#awesome-go-info)**

# Command Line
        
**[⬆ back to top](#awesome-go-info)**

# Configuration
        
*Libraries for configuration parsing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [viper](https://github.com/spf13/viper) | 18148 | 242 | 2014/04/02 | 2 weeks ago | Go configuration with fangs. |
| [godotenv](https://github.com/joho/godotenv) | 4533 | 42 | 2013/07/30 | 1 month ago | Go port of Ruby's dotenv library (Loads environment variables from `.env`). |
| [envconfig](https://github.com/kelseyhightower/envconfig) | 3966 | 39 | 2013/11/06 | 2 months ago | Go library for managing configuration data from environment variables. |
| [ini](https://github.com/go-ini/ini) | 2838 | 77 | 2014/12/18 | 3 weeks ago | Go package to read and write INI files. |
| [env](https://github.com/caarlos0/env) | 2261 | 24 | 2015/07/28 | 1 month ago | Parse environment variables to Go structs (with defaults). |
| [kong](https://github.com/alecthomas/kong) | 830 | 15 | 2018/04/10 | 1 week ago | Command-line parser with support for arbitrarily complex command-line structures and additional sources of configuration such as YAML, JSON, TOML, etc (succesor to `kingpin`). |
| [koanf](https://github.com/knadh/koanf) | 748 | 13 | 2019/06/18 | 2 weeks ago | Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line. |
| [konfig](https://github.com/lalamove/konfig) | 619 | 14 | 2019/01/18 | 1 year ago | Composable, observable and performant config handling for Go for the distributed processing era. |
| [cleanenv](https://github.com/ilyakaznacheev/cleanenv) | 462 | 6 | 2019/07/12 | 3 weeks ago | Minimalistic configuration reader (from files, ENV, and wherever you want). |
| [confita](https://github.com/heetch/confita) | 422 | 24 | 2017/12/21 | 6 months ago | Load configuration in cascade from multiple backends into a struct. |
| [aconfig](https://github.com/cristalhq/aconfig) | 337 | 5 | 2020/06/26 | 3 weeks ago | Simple, useful and opinionated config loader. |
| [config](https://github.com/gookit/config) | 328 | 13 | 2018/07/07 | 2 months ago | application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge. |
| [config](https://github.com/JeremyLoy/config) | 290 | 2 | 2019/04/02 | 2 months ago | Cloud native application configuration. Bind ENV to structs in only two lines. |
| [store](https://github.com/tucnak/store) | 260 | 6 | 2015/10/03 | 4 years ago | Lightweight configuration manager for Go. |
| [hjson-go](https://github.com/hjson/hjson-go) | 257 | 9 | 2016/08/05 | 3 weeks ago | Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments. |
| [config](https://github.com/golobby/config) | 244 | 5 | 2019/10/15 | 1 week ago | A lightweight yet powerful config package for Go projects. |
| [config](https://github.com/olebedev/config) | 242 | 9 | 2014/04/21 | 2 months ago | JSON or YAML configuration wrapper with environment variables and flags parsing. |
| [envconfig](https://github.com/vrischmann/envconfig) | 221 | 6 | 2015/04/21 | 3 months ago | Read your configuration from environment variables. |
| [config](https://github.com/joshbetz/config) | 209 | 2 | 2017/04/02 | 3 months ago | Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP. |
| [fig](https://github.com/kkyr/fig) | 190 | 5 | 2020/01/16 | 1 month ago | Tiny library for reading configuration from a file and from environment variables (with validation & defaults). |
| [xdg](https://github.com/adrg/xdg) | 175 | 4 | 2014/08/22 | 1 month ago | Go implementation of the [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html) and [XDG user directories](https://wiki.archlinux.org/index.php/XDG_user_directories). |
| [gcfg](https://github.com/go-gcfg/gcfg) | 158 | 8 | 2015/08/17 | 7 months ago | read INI-style configuration files into Go structs; supports user-defined types and subsections. |
| [goconfig](https://github.com/gosidekick/goconfig) | 150 | 13 | 2016/12/18 | 3 months ago | Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. |
| [envcfg](https://github.com/tomazk/envcfg) | 96 | 2 | 2014/11/29 | 4 years ago | Un-marshaling environment variables to Go structs. |
| [envh](https://github.com/antham/envh) | 96 | 4 | 2017/01/12 | 9 months ago | Helpers to manage environment variables. |
| [onion](https://github.com/goraz/onion) | 95 | 6 | 2015/07/22 | 5 months ago | Layer based configuration for Go, Supports JSON, TOML, YAML, properties, etcd, env, and encryption using PGP. |
| [harvester](https://github.com/beatlabs/harvester) | 94 | 13 | 2019/04/09 | 3 weeks ago | Harvester, a easy to use static and dynamic configuration package supportig seeding, env vars and Consul integration. |
| [configuro](https://github.com/sherifabdlnaby/configuro) | 80 | 4 | 2020/04/09 | 11 months ago | opinionated configuration loading & validation framework from ENV and Files focused towards 12-Factor compliant applications. |
| [xdg](https://github.com/OpenPeeDeeP/xdg) | 66 | 3 | 2017/07/20 | 1 year ago | Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html). |
| [gofigure](https://github.com/ian-kent/gofigure) | 61 | 5 | 2014/11/25 | 2 years ago | Go application configuration made easy. |
| [configure](https://github.com/paked/configure) | 55 | 4 | 2015/06/14 | 3 years ago | Provides configuration through multiple sources, including JSON, flags and environment variables. |
| [configuration](https://github.com/BoRuDar/configuration) | 45 | 2 | 2019/11/27 | 3 weeks ago | Library for initializing configuration structs from env variables, files, flags and 'default' tag. |
| [go-aws-ssm](https://github.com/PaddleHQ/go-aws-ssm) | 45 | 32 | 2019/01/24 | 11 months ago | Go package that fetches parameters from AWS System Manager - Parameter Store. |
| [uconfig](https://github.com/omeid/uconfig) | 40 | 2 | 2017/05/11 | 2 months ago | Lightweight, zero-dependency, and extendable configuration management. |
| [hocon](https://github.com/gurkankaymak/hocon) | 37 | 1 | 2020/03/01 | 2 months ago | Configuration library for working with the HOCON(a human-friendly JSON superset) format, supports features like environment variables, referencing other values, comments and multiple files. |
| [go-up](https://github.com/ufoscout/go-up) | 36 | 1 | 2018/02/18 | 2 years ago | A simple configuration library with recursive placeholders resolution and no magic. |
| [ingo](https://github.com/schachmat/ingo) | 35 | 1 | 2016/02/07 | 4 years ago | Flags persisted in an ini-like config file. |
| [mini](https://github.com/sasbury/mini) | 29 | 1 | 2015/04/29 | 3 years ago | Golang package for parsing ini-style configuration files. |
| [genv](https://github.com/sakirsensoy/genv) | 26 | 2 | 2019/07/15 | 2 years ago | Read environment variables easily with dotenv support. |
| [conflate](https://github.com/the4thamigo-uk/conflate) | 22 | 0 | 2018/02/01 | 1 year ago | Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema. |
| [env](https://github.com/junk1tm/env) | 12 | 2 | 2022/01/10 | 1 week ago | A lightweight package for loading environment variables into structs. |
| [go-ssm-config](https://github.com/ianlopshire/go-ssm-config) | 11 | 1 | 2019/12/02 | 1 year ago | Go utility for loading configuration parameters from AWS SSM (Parameter Store). |
| [envconf](https://github.com/ian-kent/envconf) | 10 | 1 | 2014/10/26 | 7 years ago | Configuration from environment. |
| [ini](https://github.com/wlevene/ini) | 9 | 1 | 2021/08/13 | 2 months ago | INI Parser & Write Library, Unmarshal to Struct,Marshal to Json,Write File,watch file. |
| [env](https://github.com/nasermirzaei89/env) | 9 | 0 | 2019/07/24 | 1 month ago | Simple useful package for read environment variables. |
| [go-ini](https://github.com/subpop/go-ini) | 7 | 2 | 2019/09/11 | 10 months ago | A Go package that marshals and unmarshals INI-files. |
| [swap](https://github.com/oblq/swap) | 5 | 2 | 2020/04/12 | 3 months ago | Instantiate/configure structs recursively, based on build environment. (YAML, TOML, JSON and env). |
| [typenv](https://github.com/diegomarangoni/typenv) | 5 | 1 | 2020/06/30 | 1 year ago | Minimalistic, zero dependency, typed environment variables library. |
| [piper](https://github.com/Yiling-J/piper) | 4 | 1 | 2021/11/17 | 2 months ago | Viper wrapper with config inheritance and key generation. |
| [gonfig](https://github.com/miladabc/gonfig) | 3 | 1 | 2021/01/21 | 6 months ago | Tag-based configuration parser which loads values from different providers into typesafe struct. |
| [go-conf](https://github.com/ThomasObenaus/go-conf) | 2 | 1 | 2021/01/27 | 3 months ago | Simple library for application configuration based on annotated structs. It supports reading the configuration from environment variables, config files and command line parameters. |
**[⬆ back to top](#awesome-go-info)**

# Continuous Integration
        
*Tools for help with continuous integration.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [drone](https://github.com/harness/drone) | 24471 | 553 | 2014/02/07 | 2 weeks ago | Drone is a Continuous Integration platform built on Docker, written in Go. |
| [cds](https://github.com/ovh/cds) | 3722 | 85 | 2016/10/11 | 1 week ago | Enterprise-Grade CI/CD and DevOps Automation Open Source Platform. |
| [goveralls](https://github.com/mattn/goveralls) | 710 | 13 | 2013/04/17 | 1 month ago | Go integration for Coveralls.io continuous code coverage tracking system. |
| [gotestfmt](https://github.com/haveyoudebuggedit/gotestfmt) | 191 | 3 | 2021/04/29 | 1 month ago | go test output for humans. |
| [overalls](https://github.com/go-playground/overalls) | 108 | 4 | 2015/07/30 | 2 years ago | Multi-Package go project coverprofile for tools like goveralls. |
| [duci](https://github.com/duck8823/duci) | 70 | 3 | 2018/04/01 | 3 weeks ago | A simple ci server no needs domain specific languages. |
| [gomason](https://github.com/nikogura/gomason) | 51 | 1 | 2017/11/18 | 1 month ago | Test, Build, Sign, and Publish your go binaries from a clean workspace. |
| [roveralls](https://github.com/lawrencewoodman/roveralls) | 15 | 2 | 2016/06/26 | 4 years ago | Recursive coverage testing tool. |
**[⬆ back to top](#awesome-go-info)**

# CSS Preprocessors
        
*Libraries for preprocessing CSS files.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gcss](https://github.com/yosssi/gcss) | 444 | 16 | 2014/09/04 | 7 years ago | Pure Go CSS Preprocessor. |
| [go-libsass](https://github.com/wellington/go-libsass) | 185 | 8 | 2015/04/19 | 1 year ago | Go wrapper to the 100% Sass compatible libsass project. |
**[⬆ back to top](#awesome-go-info)**

# Data Structures
        
*Generic datastructures and algorithms in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gods](https://github.com/emirpasic/gods) | 11024 | 341 | 2015/03/04 | 1 week ago | Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc. |
| [go-datastructures](https://github.com/Workiva/go-datastructures) | 6349 | 323 | 2014/10/29 | 9 months ago | Collection of useful, performant, and thread-safe data structures. |
| [golang-set](https://github.com/deckarep/golang-set) | 2236 | 44 | 2013/07/03 | 1 month ago | Thread-Safe and Non-Thread-Safe high-performance sets for Go. |
| [gota](https://github.com/go-gota/gota) | 1958 | 78 | 2016/02/06 | 2 months ago | Implementation of dataframes, series, and data wrangling methods for Go. |
| [roaring](https://github.com/RoaringBitmap/roaring) | 1478 | 40 | 2014/07/10 | 3 weeks ago | Go package implementing compressed bitsets. |
| [bloom](https://github.com/bits-and-blooms/bloom) | 1421 | 35 | 2011/05/21 | 4 months ago | Go package implementing Bloom filters. |
| [BoomFilters](https://github.com/tylertreat/BoomFilters) | 1405 | 39 | 2015/02/06 | 11 months ago | Probabilistic data structures for processing continuous, unbounded streams. |
| [gocache](https://github.com/eko/gocache) | 1054 | 19 | 2019/10/05 | 1 week ago | A complete Go cache library with mutiple stores (memory, memcache, redis, ...), chainable, loadable, metrics cache and more. |
| [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) | 866 | 19 | 2015/06/28 | 6 months ago | Cuckoo filter: a good alternative to a counting bloom filter implemented in Go. |
| [bitset](https://github.com/bits-and-blooms/bitset) | 830 | 29 | 2011/05/11 | 1 month ago | Go package implementing bitsets. |
| [hyperloglog](https://github.com/axiomhq/hyperloglog) | 768 | 17 | 2017/06/18 | 1 month ago | HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction. |
| [algorithms](https://github.com/shady831213/algorithms) | 610 | 19 | 2018/01/31 | 11 months ago | Algorithms and data structures.CLRS study. |
| [gostl](https://github.com/liyue201/gostl) | 586 | 14 | 2019/10/12 | 1 week ago | Data structure and algorithm library for go, designed to provide functions similar to C++ STL. |
| [trie](https://github.com/derekparker/trie) | 571 | 22 | 2014/03/06 | 1 year ago | Trie implementation in Go. |
| [ttlcache](https://github.com/ReneKroon/ttlcache) | 339 | 15 | 2014/12/13 | 1 month ago | In-memory string-interface{} cache with various time-based expiration options and callbacks. |
| [go-geoindex](https://github.com/hailocab/go-geoindex) | 335 | 65 | 2015/01/22 | 4 years ago | In-memory geo index. |
| [go-edlib](https://github.com/hbollon/go-edlib) | 308 | 11 | 2020/08/18 | 2 weeks ago | Go string comparison and edit distance algorithms library (Levenshtein, LCS, Hamming, Damerau levenshtein, Jaro-Winkler, etc.) compatible with Unicode. |
| [merkletree](https://github.com/cbergoon/merkletree) | 306 | 8 | 2017/04/12 | 2 years ago | Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures. |
| [deque](https://github.com/gammazero/deque) | 300 | 5 | 2018/04/24 | 2 months ago | Fast ring-buffer deque (double-ended queue). |
| [hilbert](https://github.com/google/hilbert) | 241 | 22 | 2015/08/06 | 3 years ago | Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves. |
| [goskiplist](https://github.com/ryszard/goskiplist) | 231 | 16 | 2012/05/09 | 2 years ago | Skip list implementation in Go. |
| [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) | 220 | 8 | 2016/04/01 | 1 year ago | Go implementation of Adaptive Radix Tree. |
| [cuckoo-filter](https://github.com/linvon/cuckoo-filter) | 198 | 5 | 2021/02/19 | 4 months ago | Cuckoo filter: a comprehensive cuckoo filter, which is configurable and space optimized compared with other implements, and all features mentioned in original paper is available. |
| [skiplist](https://github.com/MauriceGit/skiplist) | 183 | 6 | 2018/06/23 | 1 week ago | Very fast Go Skiplist implementation. |
| [binpacker](https://github.com/zhuangsirui/binpacker) | 181 | 13 | 2016/02/02 | 4 months ago | Binary packer and unpacker helps user build custom binary stream. |
| [levenshtein](https://github.com/agnivade/levenshtein) | 175 | 4 | 2014/07/30 | 9 months ago | Implementation to calculate levenshtein distance in Go. |
| [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) | 157 | 3 | 2019/01/10 | 1 year ago | Concurrent FIFO queue. |
| [bloom](https://github.com/zentures/bloom) | 146 | 8 | 2013/09/03 | 3 years ago | Bloom filters implemented in Go. |
| [iter](https://github.com/disksing/iter) | 143 | 5 | 2019/10/20 | 2 years ago | Go implementation of C++ STL iterators and algorithms. |
| [ring](https://github.com/tannerryan/ring) | 123 | 1 | 2019/01/27 | 1 year ago | Go implementation of a high performance, thread safe bloom filter. |
| [go-rquad](https://github.com/arl/go-rquad) | 121 | 4 | 2016/09/12 | 1 year ago | Region quadtrees with efficient point location and neighbour finding. |
| [bitmap](https://github.com/kelindar/bitmap) | 119 | 4 | 2021/05/28 | 3 weeks ago | Dense, zero-allocation, SIMD-enabled bitmap/bitset in Go. |
| [encoding](https://github.com/zentures/encoding) | 113 | 7 | 2013/09/20 | 4 years ago | Integer Compression Libraries for Go. |
| [bit](https://github.com/yourbasic/bit) | 110 | 8 | 2017/05/03 | 3 years ago | Golang set data structure with bonus bit-twiddling functions. |
| [remember-go](https://github.com/rocketlaunchr/remember-go) | 108 | 5 | 2019/04/04 | 10 months ago | A universal interface for caching slow database queries (backed by redis, memcached, ristretto, or in-memory). |
| [conjungo](https://github.com/InVisionApp/conjungo) | 103 | 109 | 2016/12/29 | 1 year ago | A small, powerful and flexible merge library. |
| [go-mcache](https://github.com/OrlovEvgeny/go-mcache) | 77 | 4 | 2018/04/14 | 2 years ago | Fast in-memory key:value store/cache library. Pointer caches. |
| [skiplist](https://github.com/gansidui/skiplist) | 77 | 7 | 2014/11/18 | 7 years ago | Skiplist implementation in Go. |
| [bloom](https://github.com/yourbasic/bloom) | 66 | 3 | 2017/05/06 | 4 years ago | Golang Bloom filter implementation. |
| [levenshtein](https://github.com/agext/levenshtein) | 60 | 2 | 2016/04/08 | 1 year ago | Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. |
| [count-min-log](https://github.com/seiflotfy/count-min-log) | 55 | 2 | 2015/08/16 | 5 years ago | Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory). |
| [crunch](https://github.com/superwhiskers/crunch) | 50 | 6 | 2019/02/27 | 1 month ago | Go package implementing buffers for handling various datatypes easily. |
| [nan](https://github.com/kak-tus/nan) | 48 | 3 | 2020/05/05 | 3 weeks ago | Zero allocation Nullable structures in one library with handy conversion functions, marshallers and unmarshallers. |
| [goset](https://github.com/zoumo/goset) | 44 | 1 | 2017/08/25 | 1 year ago | A useful Set collection implementation for Go. |
| [hide](https://github.com/emvi/hide) | 42 | 4 | 2019/01/16 | 3 months ago | ID type with marshalling to/from hash to prevent sending IDs to clients. |
| [pipeline](https://github.com/hyfather/pipeline) | 39 | 2 | 2018/04/25 | 3 months ago | An implementation of pipelines with fan-in and fan-out. |
| [concurrent-writer](https://github.com/free/concurrent-writer) | 38 | 5 | 2017/09/18 | 4 years ago | Highly concurrent drop-in replacement for `bufio.Writer`. |
| [deque](https://github.com/edwingeng/deque) | 38 | 4 | 2019/02/01 | 9 months ago | A highly optimized double-ended queue. |
| [timedmap](https://github.com/zekroTJA/timedmap) | 38 | 2 | 2019/01/30 | 2 months ago | Map with expiring key-value pairs. |
| [memlog](https://github.com/embano1/memlog) | 32 | 1 | 2022/01/03 | 1 week ago | An easy to use, lightweight, thread-safe and append-only in-memory data structure inspired by Apache Kafka. |
| [typ](https://github.com/gurukami/typ) | 30 | 2 | 2019/03/03 | 4 months ago | Null Types, Safe primitive type conversion and fetching value from complex structures. |
| [cmap](https://github.com/lrita/cmap) | 24 | 2 | 2019/11/26 | 1 year ago | a thread-safe concurrent map for go, support using `interface{}` as key and auto scale up shards. |
| [null](https://github.com/emvi/null) | 24 | 3 | 2018/07/04 | 3 months ago | Nullable Go types that can be marshalled/unmarshalled to/from JSON. |
| [dict](https://github.com/srfrog/dict) | 23 | 2 | 2019/04/23 | 1 year ago | Python-like dictionaries (dict) for Go. |
| [go-ef](https://github.com/amallia/go-ef) | 19 | 3 | 2017/09/22 | 4 years ago | A Go implementation of the Elias-Fano encoding. |
| [ptrie](https://github.com/viant/ptrie) | 18 | 8 | 2019/05/20 | 1 year ago | An implementation of prefix tree. |
| [set](https://github.com/StudioSol/set) | 18 | 23 | 2018/07/20 | 4 months ago | Simple set data structure implementation in Go using LinkedHashMap. |
| [mspm](https://github.com/BlackRabbitt/mspm) | 17 | 3 | 2018/05/17 | 3 years ago | Multi-String Pattern Matching Algorithm for information retrieval. |
| [fsm](https://github.com/cocoonspace/fsm) | 14 | 1 | 2021/10/11 | 4 months ago | Finite-State Machine package. |
| [parapipe](https://github.com/nazar256/parapipe) | 14 | 1 | 2021/04/09 | 8 months ago | FIFO Pipeline which parallels execution on each stage while maintaining the order of messages and results. |
| [treap](https://github.com/perdata/treap) | 14 | 3 | 2018/09/16 | 2 years ago | Persistent, fast ordered map using tree heaps. |
| [gofal](https://github.com/xxjwxc/gofal) | 12 | 2 | 2019/08/05 | 2 years ago | fractional api for Go. |
| [ordered-concurrently](https://github.com/tejzpr/ordered-concurrently) | 12 | 1 | 2021/02/28 | 2 months ago | Go module that processes work concurrently and returns output in a channel in the order of input. |
| [bloomfilter](https://github.com/OldPanda/bloomfilter) | 9 | 1 | 2021/01/01 | 7 months ago | Yet another Bloomfilter implementation in Go, compatible with Java's Guava library. |
| [gdcache](https://github.com/ulovecode/gdcache) | 7 | 1 | 2021/07/20 | 4 months ago | A pure non-intrusive cache library implemented by golang, you can use it to implement your own distributed cache. |
| [goterator](https://github.com/yaa110/goterator) | 7 | 2 | 2020/08/12 | 1 year ago | Iterator implementation to provide map and reduce functionalities. |
| [slices](https://github.com/srfrog/slices) | 7 | 3 | 2020/07/02 | 1 year ago | Functions that operate on slices; like `package strings` but adapted to work with slices. |
| [dsu](https://github.com/ihebu/dsu) | 6 | 1 | 2021/04/27 | 2 weeks ago | Disjoint Set data structure implementation in Go. |
| [parsefields](https://github.com/MonaxGT/parsefields) | 6 | 2 | 2019/04/12 | 2 years ago | Tools for parse JSON-like logs for collecting unique fields and events. |
| [merkle](https://github.com/bobg/merkle) | 2 | 2 | 2018/10/13 | 6 months ago | Space-efficient computation of Merkle root hashes and inclusion proofs. |
| [bingo](https://github.com/iancmcc/bingo) | 0 | 1 | 2021/08/22 | 1 week ago | Fast, zero-allocation, lexicographical-order-preserving packing of native types to bytes. |
**[⬆ back to top](#awesome-go-info)**

# Database
        
*SQL query builder, libraries for building and using SQL.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [prometheus](https://github.com/prometheus/prometheus) | 40896 | 1139 | 2012/11/24 | 1 week ago | Monitoring system and time series database. |
| [tidb](https://github.com/pingcap/tidb) | 30319 | 1316 | 2015/09/06 | 1 week ago | TiDB is a distributed SQL database. Inspired by the design of Google F1. |
| [cockroach](https://github.com/cockroachdb/cockroach) | 23344 | 729 | 2014/02/06 | 1 week ago | Scalable, Geo-Replicated, Transactional Datastore. |
| [influxdb](https://github.com/influxdata/influxdb) | 22825 | 751 | 2013/09/26 | 1 week ago | Scalable datastore for metrics, events, and real-time analytics. |
| [dgraph](https://github.com/dgraph-io/dgraph) | 17575 | 374 | 2015/08/25 | 2 weeks ago | Scalable, Distributed, Low Latency, High Throughput Graph Database. |
| [vitess](https://github.com/vitessio/vitess) | 13286 | 513 | 2013/06/27 | 1 week ago | vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. |
| [groupcache](https://github.com/golang/groupcache) | 11108 | 498 | 2013/07/22 | 1 month ago | Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. |
| [badger](https://github.com/dgraph-io/badger) | 10397 | 247 | 2017/01/26 | 1 week ago | Fast key-value store in Go. |
| [rqlite](https://github.com/rqlite/rqlite) | 9559 | 220 | 2014/08/23 | 1 week ago | The lightweight, distributed, relational database built on SQLite. |
| [milvus](https://github.com/milvus-io/milvus) | 9338 | 188 | 2019/09/16 | 1 week ago | Milvus is a vector database for embedding management, analytics and search. |
| [migrate](https://github.com/golang-migrate/migrate) | 7995 | 80 | 2018/01/19 | 1 week ago | Database migrations. CLI and Golang library. |
| [pgweb](https://github.com/sosedoff/pgweb) | 7186 | 147 | 2014/10/09 | 2 weeks ago | Web-based PostgreSQL database browser. |
| [immudb](https://github.com/codenotary/immudb) | 7089 | 85 | 2019/11/07 | 1 week ago | immudb is a lightweight, high-speed immutable database for systems and applications written in Go. |
| [kingshard](https://github.com/flike/kingshard) | 5929 | 404 | 2015/07/04 | 8 months ago | kingshard is a high performance proxy for MySQL powered by Golang. |
| [go-cache](https://github.com/patrickmn/go-cache) | 5816 | 116 | 2012/01/02 | 1 month ago | In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. |
| [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) | 5776 | 112 | 2018/09/30 | 1 week ago | fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL. |
| [bigcache](https://github.com/allegro/bigcache) | 5424 | 114 | 2016/03/23 | 4 weeks ago | Efficient key/value cache for gigabytes of data. |
| [bbolt](https://github.com/etcd-io/bbolt) | 5231 | 120 | 2017/06/17 | 2 weeks ago | An embedded key/value database for Go. |
| [goleveldb](https://github.com/syndtr/goleveldb) | 4889 | 180 | 2013/01/23 | 1 month ago | Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go. |
| [dtm](https://github.com/dtm-labs/dtm) | 4580 | 62 | 2021/05/16 | 1 week ago | A distributed transaction manager. Support XA, TCC, SAGA, Reliable Messages. |
| [squirrel](https://github.com/Masterminds/squirrel) | 4487 | 49 | 2014/01/18 | 1 month ago | Go library that helps you build SQL queries. |
| [orchestrator](https://github.com/openark/orchestrator) | 4392 | 264 | 2016/11/30 | 2 weeks ago | MySQL replication topology manager & visualizer. |
| [dtm](https://github.com/yedf/dtm) | 4207 | 92 | 2021/05/16 | 1 month ago | A distributed transaction manager. Support XA, TCC, SAGA, Reliable Messages. |
| [ledisdb](https://github.com/ledisdb/ledisdb) | 3802 | 185 | 2014/04/30 | 2 weeks ago | Ledisdb is a high performance NoSQL like Redis based on LevelDB. |
| [go-mysql-elasticsearch](https://github.com/go-mysql-org/go-mysql-elasticsearch) | 3715 | 179 | 2015/01/15 | 1 year ago | Sync your MySQL data into Elasticsearch automatically. |
| [buntdb](https://github.com/tidwall/buntdb) | 3636 | 99 | 2016/07/19 | 1 month ago | Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support. |
| [go-mysql](https://github.com/go-mysql-org/go-mysql) | 3438 | 158 | 2014/02/21 | 1 week ago | Go toolset to handle MySQL protocol and replication. |
| [prest](https://github.com/prest/prest) | 3089 | 82 | 2016/11/22 | 2 weeks ago | Simplify and accelerate development, ⚡ instant, realtime, high-performance on any Postgres application, existing or new. |
| [xo](https://github.com/xo/xo) | 3005 | 69 | 2016/02/05 | 2 weeks ago | Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. |
| [tiedot](https://github.com/HouzuoGuo/tiedot) | 2651 | 159 | 2013/05/26 | 5 months ago | Your NoSQL database powered by Golang. |
| [sql-migrate](https://github.com/rubenv/sql-migrate) | 2379 | 33 | 2014/09/09 | 1 week ago | Database migration tool. Allows embedding migrations into the application using go-bindata. |
| [rosedb](https://github.com/flower-corp/rosedb) | 2256 | 29 | 2020/12/06 | 1 month ago | An embedded k-v database based on LSM+WAL, supports string, list, hash, set, zset. |
| [goose](https://github.com/pressly/goose) | 2250 | 48 | 2016/02/25 | 1 week ago | Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts. |
| [nutsdb](https://github.com/xujiajun/nutsdb) | 1867 | 50 | 2018/12/07 | 2 months ago | Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set. |
| [gcache](https://github.com/bluele/gcache) | 1811 | 44 | 2015/01/24 | 1 month ago | Cache library with support for expirable Cache, LFU, LRU and ARC. |
| [cache2go](https://github.com/muesli/cache2go) | 1669 | 68 | 2013/11/11 | 1 month ago | In-memory key:value cache which supports automatic invalidation based on timeouts. |
| [godis](https://github.com/HDT3213/godis) | 1647 | 24 | 2019/06/01 | 1 month ago | A Golang implemented high-performance Redis server and cluster. |
| [goqu](https://github.com/doug-martin/goqu) | 1391 | 36 | 2015/02/21 | 4 weeks ago | Idiomatic SQL builder and query library. |
| [gendry](https://github.com/didi/gendry) | 1344 | 64 | 2017/12/01 | 3 weeks ago | Non-invasive SQL builder and powerful data binder. |
| [fastcache](https://github.com/VictoriaMetrics/fastcache) | 1340 | 30 | 2018/11/22 | 2 months ago | fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead. |
| [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) | 1275 | 72 | 2018/04/11 | 5 months ago | CovenantSQL is a SQL database on blockchain. |
| [diskv](https://github.com/peterbourgon/diskv) | 1132 | 41 | 2012/03/21 | 3 months ago | Home-grown disk-backed key-value store. |
| [atlas](https://github.com/ariga/atlas) | 1097 | 15 | 2021/04/30 | 1 week ago | A Database Toolkit. A CLI designed to help companies better work with their data. |
| [skeema](https://github.com/skeema/skeema) | 980 | 29 | 2016/10/31 | 3 weeks ago | Pure-SQL schema management system for MySQL, with support for sharding and external online schema change tools. |
| [databunker](https://github.com/securitybunker/databunker) | 952 | 27 | 2019/12/08 | 2 weeks ago | Personally identifiable information (PII) storage service built to comply with GDPR and CCPA. |
| [column](https://github.com/kelindar/column) | 871 | 14 | 2021/05/26 | 1 month ago | High-performance, columnar, embeddable in-memory store with bitmap indexing and transactions. |
| [eliasdb](https://github.com/krotik/eliasdb) | 871 | 25 | 2016/08/13 | 1 month ago | Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language. |
| [pogreb](https://github.com/akrylysov/pogreb) | 861 | 26 | 2018/01/06 | 5 months ago | Embedded key-value store for read-heavy workloads. |
| [moss](https://github.com/couchbase/moss) | 856 | 76 | 2016/02/06 | 2 months ago | Moss is a simple LSM key-value storage engine written in 100% Go. |
| [chproxy](https://github.com/ContentSquare/chproxy) | 813 | 32 | 2017/09/18 | 2 weeks ago | HTTP proxy for ClickHouse database. |
| [goavro](https://github.com/linkedin/goavro) | 777 | 23 | 2015/02/23 | 2 weeks ago | A Go package that encodes and decodes Avro data. |
| [gormigrate](https://github.com/go-gormigrate/gormigrate) | 720 | 7 | 2016/08/31 | 1 week ago | Database schema migration helper for Gorm ORM. |
| [jet](https://github.com/go-jet/jet) | 618 | 15 | 2019/03/02 | 1 week ago | Framework for writing type-safe SQL queries in Go, with ability to easily convert database query result into desired arbitrary object structure. |
| [dotsql](https://github.com/qustavo/dotsql) | 614 | 25 | 2014/11/20 | 6 months ago | Go library that helps you keep sql files in one place and use them with ease. |
| [pg_timetable](https://github.com/cybertec-postgresql/pg_timetable) | 607 | 17 | 2018/12/19 | 1 week ago | Advanced scheduling for PostgreSQL. |
| [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) | 550 | 28 | 2015/12/10 | 2 weeks ago | Powerful data retrieval methods as well as DB-agnostic query building capabilities. |
| [levigo](https://github.com/jmhodges/levigo) | 403 | 23 | 2012/01/17 | 2 months ago | Levigo is a Go wrapper for LevelDB. |
| [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) | 339 | 7 | 2017/04/29 | 4 months ago | Collects small insterts and sends big requests to ClickHouse servers. |
| [dbq](https://github.com/rocketlaunchr/dbq) | 329 | 9 | 2019/07/11 | 11 months ago | Zero boilerplate database operations for Go. |
| [pudge](https://github.com/recoilme/pudge) | 309 | 11 | 2018/11/20 | 7 months ago | Fast and simple  key/value store written using Go's standard library. |
| [sqrl](https://github.com/elgris/sqrl) | 238 | 10 | 2014/06/25 | 3 months ago | SQL query builder, fork of Squirrel with improved performance. |
| [vasto](https://github.com/chrislusf/vasto) | 236 | 19 | 2018/01/16 | 2 years ago | A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption. |
| [kivik](https://github.com/go-kivik/kivik) | 224 | 5 | 2017/02/09 | 1 month ago | Kivik provides a common Go and GopherJS client library for CouchDB, PouchDB, and similar databases. |
| [sqlingo](https://github.com/lqs/sqlingo) | 208 | 13 | 2018/11/18 | 2 weeks ago | A lightweight DSL to build SQL in Go. |
| [piladb](https://github.com/fern4lvarez/piladb) | 194 | 12 | 2015/09/08 | 1 year ago | Lightweight RESTful database engine based on stack data structures. |
| [myreplication](https://github.com/2tvenom/myreplication) | 178 | 21 | 2015/02/04 | 3 years ago | MySql binary log replication listener. Supports statement and row based replication. |
| [octillery](https://github.com/blastrain/octillery) | 153 | 19 | 2018/11/26 | 8 months ago | Go package for sharding databases ( Supports every ORM or raw SQL ). |
| [go-structured-query](https://github.com/bokwoon95/go-structured-query) | 146 | 1 | 2020/05/30 | 3 months ago | Type-safe SQL builder and struct mapper for Go. |
| [golang-scribble](https://github.com/nanobox-io/golang-scribble) | 145 | 5 | 2018/06/21 | 2 years ago | Tiny flat file JSON store. |
| [darwin](https://github.com/GuiaBolso/darwin) | 127 | 4 | 2016/04/05 | 10 months ago | Database schema evolution library for Go. |
| [migrator](https://github.com/lopezator/migrator) | 118 | 5 | 2019/02/04 | 1 year ago | Dead simple Go database migration library. |
| [lotusdb](https://github.com/flower-corp/lotusdb) | 106 | 2 | 2021/12/14 | 2 weeks ago | Fast k/v database compatible with lsm and b+tree. |
| [cache](https://github.com/akyoto/cache) | 101 | 3 | 2019/05/11 | 2 years ago | In-memory key:value store with expiration time, 0 dependencies, <100 LoC, 100% coverage. |
| [slowpoke](https://github.com/recoilme/slowpoke) | 99 | 8 | 2018/02/19 | 2 years ago | Key-value store with persistence. |
| [igor](https://github.com/galeone/igor) | 85 | 7 | 2016/03/10 | 1 year ago | Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax. |
| [unitdb](https://github.com/unit-io/unitdb) | 83 | 8 | 2019/08/29 | 3 months ago | Fast timeseries database for IoT, realtime messaging  applications. Access unitdb with pubsub over tcp or websocket using github.com/unit-io/unitd application. |
| [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) | 82 | 1 | 2018/08/11 | 5 months ago | A Go package to help write migrations with go-pg/pg. |
| [bcache](https://github.com/iwanbk/bcache) | 78 | 4 | 2018/12/26 | 2 years ago | Eventually consistent distributed in-memory  cache Go library. |
| [dbbench](https://github.com/sj14/dbbench) | 63 | 4 | 2018/11/24 | 1 week ago | Database benchmarking tool with support for several databases and scripts. |
| [couchcache](https://github.com/codingsince1985/couchcache) | 54 | 4 | 2015/04/05 | 4 months ago | RESTful caching micro-service backed by Couchbase server. |
| [godbal](https://github.com/xujiajun/godbal) | 53 | 4 | 2018/02/28 | 3 years ago | Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily. |
| [hare](https://github.com/jameycribbs/hare) | 50 | 6 | 2016/10/05 | 11 months ago | A simple database management system that stores each table as a text file of line-delimited JSON. |
| [datagen](https://github.com/codingconcepts/datagen) | 46 | 1 | 2019/04/18 | 1 year ago | A fast data generator that's multi-table aware and supports multi-row DML. |
| [buildsqlx](https://github.com/arthurkushman/buildsqlx) | 45 | 1 | 2019/08/18 | 3 weeks ago | Go database query builder library for PostgreSQL. |
| [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) | 40 | 6 | 2017/12/18 | 4 years ago | BigCache with clustering support and individual item expiration. |
| [sqlf](https://github.com/leporo/sqlf) | 39 | 4 | 2019/07/20 | 1 month ago | Fast SQL query builder. |
| [prep](https://github.com/hexdigest/prep) | 31 | 3 | 2017/12/11 | 4 years ago | Use prepared SQL statements without changing your code. |
| [avro](https://github.com/khezen/avro) | 29 | 3 | 2019/04/07 | 1 month ago | Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes. |
| [coffer](https://github.com/claygod/coffer) | 28 | 5 | 2019/05/13 | 3 weeks ago | Simple ACID key-value database that supports transactions. |
| [sqlize](https://github.com/sunary/sqlize) | 28 | 1 | 2020/09/08 | 1 month ago | Database migration generator. Allows generate sql migration from model and existing sql by differ them. |
| [go-fixtures](https://github.com/RichardKnop/go-fixtures) | 27 | 1 | 2015/12/24 | 2 years ago | Django style fixtures for Golang's excellent built-in database/sql library. |
| [pravasan](https://github.com/pravasan/pravasan) | 27 | 7 | 2015/01/03 | 3 years ago | Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc. |
| [bqb](https://github.com/nullism/bqb) | 27 | 1 | 2021/07/31 | 2 months ago | Lightweight and easy to learn query builder. |
| [qry](https://github.com/HnH/qry) | 20 | 3 | 2019/08/20 | 4 months ago | Tool that generates constants from files with raw SQL queries. |
| [gosql](https://github.com/twharmon/gosql) | 18 | 1 | 2020/01/08 | 1 year ago | SQL Query builder with better null values support. |
| [tempdb](https://github.com/rafaeljesus/tempdb) | 16 | 3 | 2017/03/17 | 4 years ago | Key-value store for temporary items. |
| [migrator](https://github.com/larapulse/migrator) | 14 | 2 | 2020/06/27 | 9 months ago | MySQL database migrator designed to run migrations to your features and manage database schema update with intuitive go code. |
| [rwdb](https://github.com/andizzle/rwdb) | 14 | 2 | 2017/10/04 | 4 years ago | rwdb provides read replica capability for multiple database servers setup. |
| [schema](https://github.com/adlio/schema) | 11 | 3 | 2019/09/24 | 2 months ago | Library to embed schema migrations for database/sql-compatible databases inside your Go binaries. |
| [mpath-go](https://github.com/spacetab-io/mpath-go) | 11 | 4 | 2020/01/09 | 2 years ago | MPTT (Modified Preorder Tree Traversal) package for SQL records - materialized path realisation. |
| [go-pg-migrate](https://github.com/lawzava/go-pg-migrate) | 7 | 1 | 2021/01/16 | 2 months ago | CLI-friendly package for go-pg migrations management. |
| [ttlcache](https://github.com/cheshir/ttlcache) | 5 | 2 | 2021/01/06 | 11 months ago | In-memory key value storage with TTL for each record. |
| [ormlite](https://github.com/pupizoid/ormlite) | 2 | 2 | 2018/06/28 | 1 year ago | Lightweight package containing some ORM-like features and helpers for sqlite databases. |
**[⬆ back to top](#awesome-go-info)**

# Database Drivers
        
*Libraries for connecting and operating databases.*
**[⬆ back to top](#awesome-go-info)**

## NoSQL Databases
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [cayley](https://github.com/cayleygraph/cayley) | 14066 | 599 | 2014/06/05 | 3 weeks ago | Graph database with support for multiple backends. |
| [redis](https://github.com/go-redis/redis) | 13427 | 246 | 2012/07/25 | 1 week ago | Redis client for Golang. |
| [redigo](https://github.com/gomodule/redigo) | 8861 | 290 | 2012/04/14 | 2 weeks ago | Redigo is a Go client for the Redis database. |
| [bleve](https://github.com/blevesearch/bleve) | 8167 | 247 | 2014/04/17 | 3 weeks ago | Modern text indexing library for go. |
| [elastic](https://github.com/olivere/elastic) | 6531 | 170 | 2012/12/06 | 2 weeks ago | Elasticsearch client for Go. |
| [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) | 6412 | 129 | 2017/02/08 | 1 week ago | Official MongoDB driver for the Go language. |
| [riot](https://github.com/go-ego/riot) | 6040 | 200 | 2017/06/21 | 1 year ago | Go Open Source, Distributed, Simple and efficient Search Engine. |
| [go-elasticsearch](https://github.com/elastic/go-elasticsearch) | 3877 | 316 | 2017/03/27 | 1 week ago | Official Elasticsearch client for Go. |
| [mgo](https://github.com/globalsign/mgo) | 1933 | 62 | 2017/04/13 | 3 months ago | (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms. |
| [rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1590 | 50 | 2013/09/12 | 4 months ago | Go language driver for RethinkDB. |
| [elastigo](https://github.com/mattbaird/elastigo) | 947 | 47 | 2012/10/12 | 3 years ago | Elasticsearch client library. |
| [elasticsql](https://github.com/cch123/elasticsql) | 855 | 33 | 2016/08/24 | 3 months ago | Convert sql to elasticsearch dsl in Go. |
| [qmgo](https://github.com/qiniu/qmgo) | 791 | 21 | 2020/08/04 | 2 months ago | The MongoDB driver for Go. It‘s based on official MongoDB driver but easier to use like Mgo. |
| [mgm](https://github.com/Kamva/mgm) | 452 | 17 | 2019/12/27 | 1 week ago | MongoDB model-based ODM for Go (based on official MongoDB driver). |
| [redeo](https://github.com/bsm/redeo) | 414 | 26 | 2014/03/06 | 1 year ago | Redis-protocol compatible TCP servers/services. |
| [gokv](https://github.com/philippgille/gokv) | 397 | 10 | 2018/10/08 | 5 months ago | Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more). |
| [neoism](https://github.com/jmcvetta/neoism) | 383 | 24 | 2012/07/12 | 2 years ago | Neo4j client for Golang. |
| [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) | 375 | 42 | 2014/07/26 | 1 week ago | Aerospike client in Go language. |
| [gocb](https://github.com/couchbase/gocb) | 331 | 68 | 2015/01/15 | 1 week ago | Official Couchbase Go SDK. |
| [go-couchbase](https://github.com/couchbase/go-couchbase) | 312 | 25 | 2012/01/19 | 1 month ago | Couchbase client in Go. |
| [go-rejson](https://github.com/nitishm/go-rejson) | 239 | 7 | 2018/04/23 | 2 weeks ago | Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease. |
| [cachego](https://github.com/faabiosr/cachego) | 160 | 7 | 2016/10/05 | 1 month ago | Golang Cache component for multiple drivers. |
| [godis](https://github.com/piaohao/godis) | 102 | 10 | 2019/06/14 | 1 year ago | redis client implement by golang, inspired by jedis. |
| [skizze](https://github.com/seiflotfy/skizze) | 82 | 6 | 2016/01/17 | 5 years ago | probabilistic data-structures service and storage. |
| [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) | 77 | 7 | 2011/06/04 | 3 years ago | Neo4j REST Client in golang. |
| [arangolite](https://github.com/solher/arangolite) | 70 | 6 | 2015/10/04 | 11 months ago | Lightweight golang driver for ArangoDB. |
| [go-pilosa](https://github.com/pilosa/go-pilosa) | 51 | 21 | 2016/09/30 | 1 year ago | Go client library for Pilosa. |
| [goforestdb](https://github.com/couchbase/goforestdb) | 33 | 39 | 2014/05/14 | 5 years ago | Go bindings for ForestDB. |
| [goriak](https://github.com/zegl/goriak) | 27 | 4 | 2016/10/05 | 5 months ago | Go language driver for Riak KV. |
| [neo4j](https://github.com/cihangir/neo4j) | 27 | 4 | 2013/05/18 | 6 years ago | Neo4j Rest API Bindings for Golang. |
| [goes](https://github.com/OwnLocal/goes) | 26 | 34 | 2015/12/28 | 1 year ago | Library to interact with Elasticsearch. |
| [dsc](https://github.com/viant/dsc) | 24 | 16 | 2016/06/13 | 1 year ago | Datastore connectivity for SQL, NoSQL, structured files. |
| [xredis](https://github.com/shomali11/xredis) | 18 | 2 | 2017/06/14 | 2 years ago | Typesafe, customizable, clean & easy to use Redis client. |
| [godscache](https://github.com/defcronyke/godscache) | 10 | 3 | 2018/05/08 | 3 years ago | A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached. |
| [asc](https://github.com/viant/asc) | 8 | 11 | 2016/06/13 | 2 years ago | Datastore Connectivity for Aerospike for go. |
| [gocosmos](https://github.com/btnguyen2k/gocosmos) | 6 | 2 | 2020/12/06 | 7 months ago | REST client and standard `database/sql` driver for Azure Cosmos DB. |
**[⬆ back to top](#awesome-go-info)**

## Relational Databases
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [mysql](https://github.com/go-sql-driver/mysql) | 11864 | 407 | 2012/12/09 | 3 weeks ago | MySQL driver for Go. |
| [pq](https://github.com/lib/pq) | 7056 | 153 | 2012/03/12 | 1 month ago | Pure Go Postgres driver for database/sql. |
| [go-sqlite3](https://github.com/mattn/go-sqlite3) | 5417 | 150 | 2011/11/11 | 2 weeks ago | SQLite3 driver for go that uses database/sql. |
| [pgx](https://github.com/jackc/pgx) | 4999 | 89 | 2013/03/30 | 1 week ago | PostgreSQL driver supporting features beyond those exposed by database/sql. |
| [go-mssqldb](https://github.com/denisenkom/go-mssqldb) | 1509 | 65 | 2013/12/16 | 3 weeks ago | Microsoft MSSQL driver for Go. |
| [go-oci8](https://github.com/mattn/go-oci8) | 576 | 41 | 2012/02/29 | 3 months ago | Oracle driver for go that uses database/sql. |
| [sqlhooks](https://github.com/qustavo/sqlhooks) | 514 | 6 | 2016/04/20 | 1 month ago | Attach hooks to any database/sql driver. |
| [godror](https://github.com/godror/godror) | 324 | 23 | 2019/11/21 | 2 weeks ago | Oracle driver for Go, using the ODPI-C driver. |
| [firebirdsql](https://github.com/nakagami/firebirdsql) | 164 | 17 | 2013/08/27 | 1 month ago | Firebird RDBMS SQL driver for Go. |
| [go-adodb](https://github.com/mattn/go-adodb) | 126 | 12 | 2011/11/14 | 2 years ago | Microsoft ActiveX Object DataBase driver for go that uses database/sql. |
| [gofreetds](https://github.com/minus5/gofreetds) | 106 | 23 | 2012/12/06 | 1 year ago | Microsoft MSSQL driver. Go wrapper over [FreeTDS](https://www.freetds.org). |
| [sqinn-go](https://github.com/cvilsmeier/sqinn-go) | 100 | 2 | 2020/06/06 | 8 months ago | SQLite with pure Go. |
| [calcite-avatica-go](https://github.com/apache/calcite-avatica-go) | 88 | 25 | 2017/08/08 | 1 year ago | Apache Avatica/Phoenix SQL driver for database/sql. |
| [bgc](https://github.com/viant/bgc) | 16 | 11 | 2016/06/13 | 2 years ago | Datastore Connectivity for BigQuery for go. |
| [pig](https://github.com/alexeyco/pig) | 7 | 1 | 2021/04/15 | 10 months ago | Simple [pgx](https://github.com/jackc/pgx) wrapper to execute and [scan](https://github.com/georgysavva/scany) query results easily. |
**[⬆ back to top](#awesome-go-info)**

# Date and Time
        
*Libraries for working with dates and times.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [now](https://github.com/jinzhu/now) | 3550 | 67 | 2013/11/18 | 2 months ago | Now is a time toolkit for golang. |
| [dateparse](https://github.com/araddon/dateparse) | 1627 | 23 | 2014/04/21 | 2 weeks ago | Parse date's without knowing format in advance. |
| [carbon](https://github.com/golang-module/carbon) | 1600 | 13 | 2020/09/07 | 3 months ago | A simple, semantic and developer-friendly golang package for datetime. |
| [carbon](https://github.com/uniplaces/carbon) | 678 | 43 | 2016/08/03 | 3 months ago | Simple Time extension with a lot of util methods, ported from PHP Carbon library. |
| [durafmt](https://github.com/hako/durafmt) | 417 | 6 | 2016/05/20 | 8 months ago | Time duration formatting library for Go. |
| [timeutil](https://github.com/leekchan/timeutil) | 187 | 8 | 2015/08/02 | 3 years ago | Useful extensions (Timedelta, Strftime, ...) to the golang's time package. |
| [gostradamus](https://github.com/bykof/gostradamus) | 159 | 5 | 2020/04/07 | 2 months ago | A Go package for working with dates. |
| [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) | 114 | 5 | 2016/01/31 | 2 months ago | The implementation of the Persian (Solar Hijri) Calendar in Go (golang). |
| [iso8601](https://github.com/relvacode/iso8601) | 98 | 4 | 2017/04/25 | 7 months ago | Efficiently parse ISO8601 date-times without regex. |
| [date](https://github.com/rickb777/date) | 87 | 3 | 2015/11/23 | 1 week ago | Augments Time for working with dates, date ranges, time spans, periods, and time-of-day. |
| [timespan](https://github.com/SaidinWoT/timespan) | 78 | 6 | 2014/10/07 | 2 years ago | For interacting with intervals of time, defined as a start time and a duration. |
| [feiertage](https://github.com/wlbr/feiertage) | 41 | 3 | 2015/11/04 | 6 months ago | Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving... |
| [go-sunrise](https://github.com/nathan-osman/go-sunrise) | 41 | 5 | 2017/06/15 | 8 months ago | Calculate the sunrise and sunset times for a given location. |
| [go-str2duration](https://github.com/xhit/go-str2duration) | 37 | 3 | 2020/02/02 | 1 year ago | Convert string to duration. Support time.Duration returned string and more. |
| [kair](https://github.com/GuilhermeCaruso/kair) | 22 | 2 | 2018/10/03 | 1 year ago | Date and Time - Golang Formatting Library. |
| [cronrange](https://github.com/1set/cronrange) | 14 | 2 | 2019/11/10 | 1 week ago | Parses Cron-style time range expressions, checks if the given time is within any ranges. |
| [nulltime](https://github.com/kirillDanshin/nulltime) | 12 | 2 | 2016/03/06 | 4 years ago | Nullable `time.Time`. |
| [tuesday](https://github.com/osteele/tuesday) | 11 | 3 | 2017/08/10 | 8 months ago | Ruby-compatible Strftime function. |
| [strftime](https://github.com/awoodbeck/strftime) | 8 | 2 | 2018/02/10 | 4 years ago | C99-compatible strftime formatter. |
| [go-week](https://github.com/stoewer/go-week) | 7 | 5 | 2018/02/23 | 3 months ago | An efficient package to work with ISO8601 week dates. |
**[⬆ back to top](#awesome-go-info)**

# Distributed Systems
        
*Packages that help with building Distributed Systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [kit](https://github.com/go-kit/kit) | 22265 | 689 | 2015/02/03 | 2 weeks ago | Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc. |
| [go-micro](https://github.com/asim/go-micro) | 17605 | 506 | 2015/01/13 | 4 weeks ago | A distributed systems development framework. |
| [kratos](https://github.com/go-kratos/kratos) | 16362 | 419 | 2019/01/10 | 2 weeks ago | A modular-designed and easy-to-use microservices framework in Go. |
| [grpc-go](https://github.com/grpc/grpc-go) | 15308 | 481 | 2014/12/08 | 1 week ago | The Go language implementation of gRPC. HTTP/2 based RPC. |
| [go-zero](https://github.com/zeromicro/go-zero) | 14408 | 245 | 2020/08/07 | 1 week ago | A web and rpc framework. It's born to ensure the stability of the busy sites with resilient design. Builtin goctl greatly improves the development productivity. |
| [micro](https://github.com/micro/micro) | 10888 | 331 | 2015/01/16 | 1 week ago | A distributed systems runtime for the cloud and beyond. |
| [nats-server](https://github.com/nats-io/nats-server) | 10447 | 378 | 2012/10/29 | 1 week ago | Lightweight, high performance messaging system for microservices, IoT, and cloud native systems. |
| [rpcx](https://github.com/smallnest/rpcx) | 6672 | 348 | 2016/05/18 | 2 weeks ago | Distributed pluggable RPC service framework like alibaba Dubbo. |
| [raft](https://github.com/hashicorp/raft) | 5637 | 375 | 2013/11/05 | 1 week ago | Golang implementation of the Raft consensus protocol, by HashiCorp. |
| [lura](https://github.com/luraproject/lura) | 4843 | 124 | 2016/11/04 | 4 weeks ago | Ultra performant API Gateway framework with middlewares. |
| [torrent](https://github.com/anacrolix/torrent) | 4196 | 130 | 2015/01/08 | 1 week ago | BitTorrent client package. |
| [dragonboat](https://github.com/lni/dragonboat) | 4094 | 145 | 2018/12/23 | 1 week ago | A feature complete and high performance multi-group Raft library in Go. |
| [emitter](https://github.com/emitter-io/emitter) | 3169 | 105 | 2016/10/29 | 1 month ago | High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love. |
| [glow](https://github.com/chrislusf/glow) | 3026 | 144 | 2015/06/14 | 3 years ago | Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go. |
| [gleam](https://github.com/chrislusf/gleam) | 2976 | 149 | 2016/08/26 | 9 months ago | Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed. |
| [liftbridge](https://github.com/liftbridge-io/liftbridge) | 2200 | 68 | 2017/10/13 | 1 week ago | Lightweight, fault-tolerant message streams for NATS. |
| [hprose-golang](https://github.com/hprose/hprose-golang) | 1192 | 91 | 2014/02/14 | 2 months ago | Very newbility RPC Library, support 25+ languages now. |
| [ringpop-go](https://github.com/uber/ringpop-go) | 721 | 2471 | 2015/06/05 | 11 months ago | Scalable, fault-tolerant application-layer sharding for Go applications. |
| [rain](https://github.com/cenkalti/rain) | 667 | 18 | 2014/05/21 | 2 months ago | BitTorrent client and library. |
| [gorpc](https://github.com/valyala/gorpc) | 650 | 24 | 2014/11/20 | 2 years ago | Simple, fast and scalable RPC library for high load. |
| [go-health](https://github.com/InVisionApp/go-health) | 616 | 118 | 2017/11/29 | 2 years ago | Library for enabling asynchronous dependency health checks in your service. |
| [redislock](https://github.com/bsm/redislock) | 578 | 9 | 2019/06/24 | 1 month ago | Simplified distributed locking implementation using Redis. |
| [arpc](https://github.com/lesismal/arpc) | 464 | 19 | 2020/05/19 | 1 week ago | More effective network communication, support two-way-calling, notify, broadcast. |
| [go-sundheit](https://github.com/AppsFlyer/go-sundheit) | 452 | 10 | 2019/04/08 | 6 months ago | A library built to provide support for defining async service health checks for golang services. |
| [consistent](https://github.com/buraksezer/consistent) | 450 | 15 | 2018/03/25 | 8 months ago | Consistent hashing with bounded loads. |
| [digota](https://github.com/digota/digota) | 424 | 29 | 2017/08/14 | 1 year ago | grpc ecommerce microservice. |
| [sleuth](https://github.com/ursiform/sleuth) | 342 | 10 | 2016/04/23 | 3 years ago | Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)). |
| [go-jump](https://github.com/dgryski/go-jump) | 336 | 16 | 2014/06/15 | 3 months ago | Port of Google's "Jump" Consistent Hash function. |
| [go-doudou](https://github.com/unionj-cloud/go-doudou) | 244 | 13 | 2021/02/24 | 1 week ago | A gossip protocol and OpenAPI 3.0 spec based decentralized microservice framework. Built-in go-doudou cli focusing on low-code and rapid dev can power up your productivity. |
| [dht](https://github.com/anacrolix/dht) | 222 | 10 | 2016/12/14 | 1 week ago | BitTorrent Kademlia DHT implementation. |
| [jsonrpc](https://github.com/ybbus/jsonrpc) | 204 | 9 | 2016/11/10 | 5 months ago | JSON-RPC 2.0 HTTP client implementation. |
| [jsonrpc](https://github.com/osamingo/jsonrpc) | 158 | 5 | 2016/10/28 | 4 months ago | The jsonrpc package helps implement of JSON-RPC 2.0. |
| [celeriac.v1](https://github.com/svcavallar/celeriac.v1) | 70 | 4 | 2015/10/10 | 1 year ago | Library for adding support for interacting and monitoring Celery workers, tasks and events in Go. |
| [doublejump](https://github.com/edwingeng/doublejump) | 68 | 5 | 2018/06/26 | 6 months ago | A revamped Google's jump consistent hash. |
| [semaphore](https://github.com/jexia/semaphore) | 67 | 14 | 2020/02/05 | 8 months ago | A straightforward (micro) service orchestrator. |
| [outboxer](https://github.com/italolelis/outboxer) | 61 | 0 | 2019/02/01 | 2 weeks ago | Outboxer is a go library that implements the outbox pattern. |
| [flowgraph](https://github.com/vectaport/flowgraph) | 45 | 1 | 2018/08/29 | 9 months ago | flow-based programming package. |
| [drmaa](https://github.com/dgruber/drmaa) | 36 | 3 | 2013/03/17 | 1 year ago | Job submission library for cluster schedulers based on the DRMAA standard. |
| [go-mysql-lock](https://github.com/sanketplus/go-mysql-lock) | 36 | 1 | 2020/06/06 | 6 months ago | MySQL based distributed lock. |
| [go-pdu](https://github.com/pdupub/go-pdu) | 34 | 5 | 2018/10/08 | 2 weeks ago | A decentralized identity-based social network. |
| [micro](https://github.com/gmsec/micro) | 17 | 3 | 2020/05/03 | 3 months ago | A Go distributed systems development framework. |
| [dynatomic](https://github.com/tylfin/dynatomic) | 14 | 1 | 2019/02/08 | 1 year ago | A library for using DynamoDB as an atomic counter. |
| [consistenthash](https://github.com/mbrostami/consistenthash) | 10 | 1 | 2020/04/22 | 1 year ago | Consistent hashing with configurable replicas. |
| [failured](https://github.com/andy2046/failured) | 3 | 1 | 2021/07/26 | 6 months ago | adaptive accrual failure detector for distributed systems. |
**[⬆ back to top](#awesome-go-info)**

# Dynamic DNS
        
*Tools for updating dynamic DNS records.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [godns](https://github.com/TimothyYe/godns) | 955 | 33 | 2014/05/11 | 1 week ago | A dynamic DNS client tool, supports DNSPod & HE.net, written in Go. |
| [ddns](https://github.com/skibish/ddns) | 202 | 8 | 2017/03/13 | 4 months ago | Personal DDNS client with Digital Ocean Networking DNS as backend. |
**[⬆ back to top](#awesome-go-info)**

# Email
        
*Libraries and tools that implement email creation and sending.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [MailHog](https://github.com/mailhog/MailHog) | 9627 | 147 | 2014/04/16 | 1 month ago | Email and SMTP testing with web and API interface. |
| [hermes](https://github.com/matcornic/hermes) | 2420 | 28 | 2017/03/25 | 2 months ago | Golang package that generates clean, responsive HTML e-mails. |
| [email](https://github.com/jordan-wright/email) | 1931 | 51 | 2013/12/12 | 2 months ago | A robust and flexible email library for Go. |
| [go-imap](https://github.com/emersion/go-imap) | 1452 | 45 | 2016/04/26 | 3 weeks ago | IMAP library for clients and servers. |
| [sendgrid-go](https://github.com/sendgrid/sendgrid-go) | 791 | 199 | 2013/09/12 | 1 week ago | SendGrid's Go library for sending email. |
| [mailgun-go](https://github.com/mailgun/mailgun-go) | 577 | 73 | 2014/02/28 | 3 weeks ago | Go library for sending mail with the Mailgun API. |
| [email-verifier](https://github.com/AfterShip/email-verifier) | 366 | 25 | 2020/12/18 | 1 week ago | A Go library for email verification without sending any emails. |
| [go-simple-mail](https://github.com/xhit/go-simple-mail) | 267 | 5 | 2019/09/15 | 3 weeks ago | Very simple package to send emails with SMTP Keep Alive and two timeouts: Connect and Send. |
| [go-message](https://github.com/emersion/go-message) | 236 | 14 | 2016/12/31 | 2 weeks ago | Streaming library for the Internet Message Format and mail messages. |
| [hectane](https://github.com/hectane/hectane) | 214 | 14 | 2015/08/28 | 1 year ago | Lightweight SMTP client providing an HTTP API. |
| [douceur](https://github.com/aymerick/douceur) | 208 | 3 | 2015/04/09 | 8 months ago | CSS inliner for your HTML emails. |
| [mailchain](https://github.com/mailchain/mailchain) | 107 | 7 | 2019/04/11 | 2 weeks ago | Send encrypted emails to blockchain addresses written in Go. |
| [go-premailer](https://github.com/vanng822/go-premailer) | 81 | 2 | 2015/02/16 | 11 months ago | Inline styling for HTML mail in Go. |
| [go-dkim](https://github.com/toorop/go-dkim) | 75 | 4 | 2015/04/29 | 1 year ago | DKIM library, to sign & verify email. |
| [smtp](https://github.com/mailhog/smtp) | 68 | 9 | 2014/12/24 | 3 months ago | SMTP server protocol state machine. |
| [go-smtp-mock](https://github.com/mocktools/go-smtp-mock) | 26 | 1 | 2021/08/31 | 2 weeks ago | Lightweight configurable multithreaded fake SMTP server. Mimic any SMTP behaviour for your test environment. |
| [go-email-validator](https://github.com/go-email-validator/go-email-validator) | 19 | 4 | 2020/12/10 | 1 week ago | Modular email validator for syntax, disposable, smtp, etc... checking. |
**[⬆ back to top](#awesome-go-info)**

# Embeddable Scripting Languages
        
*Embedding other languages inside your go code.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gopher-lua](https://github.com/yuin/gopher-lua) | 4552 | 151 | 2015/02/15 | 1 month ago | Lua 5.1 VM and compiler written in Go. |
| [goja](https://github.com/dop251/goja) | 2695 | 65 | 2016/11/04 | 3 weeks ago | ECMAScript 5.1(+) implementation in Go. |
| [tengo](https://github.com/d5/tengo) | 2611 | 55 | 2019/01/09 | 2 weeks ago | Bytecode compiled script language for Go. |
| [expr](https://github.com/antonmedv/expr) | 2382 | 46 | 2018/07/14 | 1 week ago | Expression evaluation engine for Go: fast, non-Turing complete, dynamic typing, static typing. |
| [go-lua](https://github.com/Shopify/go-lua) | 2250 | 368 | 2013/12/20 | 3 weeks ago | Port of the Lua 5.2 VM to pure Go. |
| [go-python](https://github.com/sbinet/go-python) | 1347 | 45 | 2012/07/09 | 10 months ago | naive go bindings to the CPython C-API. |
| [anko](https://github.com/mattn/anko) | 1203 | 47 | 2014/03/28 | 2 months ago | Scriptable interpreter written in Go. |
| [cel-go](https://github.com/google/cel-go) | 1045 | 31 | 2018/03/09 | 2 months ago | Fast, portable, non-Turing complete expression evaluation with gradual typing. |
| [core](https://github.com/metacall/core) | 880 | 19 | 2018/12/26 | 1 week ago | Cross-platform Polyglot Runtime which supports NodeJS, JavaScript, TypeScript, Python, Ruby, C#, WebAssembly, Java, Cobol and more. |
| [go-php](https://github.com/deuill/go-php) | 835 | 43 | 2015/09/17 | 2 months ago | PHP bindings for Go. |
| [go-duktape](https://github.com/olebedev/go-duktape) | 779 | 27 | 2015/01/08 | 4 months ago | Duktape JavaScript engine bindings for Go. |
| [golua](https://github.com/aarzilli/golua) | 581 | 35 | 2010/12/06 | 2 months ago | Go bindings for Lua C API. |
| [gisp](https://github.com/jcla1/gisp) | 479 | 21 | 2014/01/11 | 4 years ago | Simple LISP in Go. |
| [gval](https://github.com/PaesslerAG/gval) | 446 | 17 | 2017/09/27 | 2 months ago | A highly customizable expression language written in Go. |
| [prolog](https://github.com/ichiban/prolog) | 334 | 9 | 2020/11/03 | 2 weeks ago | Embeddable Prolog. |
| [gentee](https://github.com/gentee/gentee) | 88 | 3 | 2018/01/14 | 3 weeks ago | Embeddable scripting programming language. |
| [binder](https://github.com/alexeyco/binder) | 55 | 2 | 2017/04/02 | 3 years ago | Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua). |
| [purl](https://github.com/ian-kent/purl) | 34 | 3 | 2014/11/29 | 7 years ago | Perl 5.18.2 embedded in Go. |
| [ngaro](https://github.com/db47h/ngaro) | 22 | 2 | 2016/08/09 | 3 years ago | Embeddable Ngaro VM implementation enabling scripting in Retro. |
| [ecal](https://github.com/krotik/ecal) | 17 | 2 | 2020/11/30 | 8 months ago | A simple embeddable scripting language which supports concurrent event processing. |
**[⬆ back to top](#awesome-go-info)**

# Error Handling
        
*Libraries for handling errors.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [errors](https://github.com/pkg/errors) | 7597 | 109 | 2015/12/27 | 3 months ago | Package that provides simple error handling primitives. |
| [go-multierror](https://github.com/hashicorp/go-multierror) | 1493 | 256 | 2014/12/15 | 3 weeks ago | Go (golang) package for representing a list of errors as a single error. |
| [eris](https://github.com/rotisserie/eris) | 910 | 12 | 2019/09/07 | 6 months ago | A better way to handle, trace, and log errors in Go. Compatible with the standard error library and github.com/pkg/errors. |
| [errorx](https://github.com/joomcode/errorx) | 806 | 75 | 2018/08/17 | 4 weeks ago | A feature rich error package with stack traces, composition of errors and more. |
| [tracerr](https://github.com/ztrue/tracerr) | 702 | 11 | 2019/02/06 | 2 years ago | Golang errors with stack trace and source fragments. |
| [errlog](https://github.com/snwfdhmp/errlog) | 401 | 5 | 2019/02/16 | 1 year ago | Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place. |
| [emperror](https://github.com/emperror/emperror) | 242 | 4 | 2017/06/13 | 1 year ago | Error handling tools and best practices for Go libraries and applications. |
| [errors](https://github.com/emperror/errors) | 113 | 5 | 2019/07/09 | 2 weeks ago | Drop-in replacement for the standard library errors package and github.com/pkg/errors. Provides various error handling primitives. |
| [errors](https://github.com/bnkamalesh/errors) | 27 | 2 | 2020/07/17 | 2 months ago | Drop-in replacement for builting Go errors. This is a minimal error handling package with custom error types, user friendly messages, Unwrap & Is. With very easy to use and straightforward helper functions. |
| [falcon](https://github.com/SonicRoshan/falcon) | 7 | 2 | 2019/09/09 | 2 years ago | A Simple Yet Highly Powerful Package For Error Handling. |
| [falcon](https://github.com/ThundR67/falcon) | 7 | 2 | 2019/09/09 | 2 years ago | A Simple Yet Highly Powerful Package For Error Handling. |
| [errors](https://github.com/neuronlabs/errors) | 3 | 2 | 2019/07/26 | 2 years ago | Simple golang error handling with classification primitives. |
| [errors](https://github.com/PumpkinSeed/errors) | 3 | 1 | 2020/01/08 | 2 years ago | The most simple error wrapper with awesome performance and minimal memory overhead. |
**[⬆ back to top](#awesome-go-info)**

# File Handling
        
*Libraries for handling files and file systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [afero](https://github.com/spf13/afero) | 4237 | 87 | 2014/10/28 | 1 month ago | FileSystem Abstraction System for Go. |
| [pdfcpu](https://github.com/pdfcpu/pdfcpu) | 2956 | 67 | 2017/06/18 | 3 weeks ago | PDF processor. |
| [gdu](https://github.com/dundee/gdu) | 1597 | 19 | 2018/02/24 | 1 week ago | Disk usage analyzer with console interface. |
| [notify](https://github.com/rjeczalik/notify) | 717 | 30 | 2014/09/08 | 6 months ago | File system event notification library with simple API, similar to os/signal. |
| [copy](https://github.com/otiai10/copy) | 388 | 7 | 2017/09/01 | 1 month ago | Copy directory recursively. |
| [bigfile](https://github.com/bigfile/bigfile) | 215 | 16 | 2019/07/15 | 2 years ago | A file transfer system, support to manage files with http api, rpc call and ftp client. |
| [afs](https://github.com/viant/afs) | 173 | 12 | 2019/08/19 | 3 months ago | Abstract File Storage (mem, scp, zip, tar, cloud: s3, gs) for Go. |
| [vfs](https://github.com/C2FO/vfs) | 150 | 22 | 2017/08/01 | 1 month ago | A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS. |
| [go-exiftool](https://github.com/barasher/go-exiftool) | 100 | 4 | 2019/05/12 | 1 week ago | Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...). |
| [go-csv-tag](https://github.com/artonge/go-csv-tag) | 95 | 1 | 2017/06/18 | 3 months ago | Load csv file using tag. |
| [opc](https://github.com/qmuntal/opc) | 70 | 3 | 2018/11/06 | 11 months ago | Load Open Packaging Conventions (OPC) files for Go. |
| [skywalker](https://github.com/dixonwille/skywalker) | 69 | 4 | 2017/08/01 | 5 months ago | Package to allow one to concurrently go through a filesystem with ease. |
| [checksum](https://github.com/codingsince1985/checksum) | 50 | 3 | 2014/11/05 | 2 months ago | Compute message digest, like MD5, SHA256, SHA1, CRC or BLAKE2s, for large files. |
| [tarfs](https://github.com/posener/tarfs) | 49 | 2 | 2017/03/10 | 1 year ago | Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files. |
| [parquet](https://github.com/parsyl/parquet) | 47 | 7 | 2019/01/29 | 4 months ago | Read and write [parquet](https://parquet.apache.org) files. |
| [baraka](https://github.com/xis/baraka) | 38 | 2 | 2020/07/12 | 5 months ago | A library to process http file uploads easily. |
| [go-gtfs](https://github.com/artonge/go-gtfs) | 31 | 2 | 2017/07/09 | 1 year ago | Load gtfs files in go. |
| [flop](https://github.com/homedepot/flop) | 30 | 18 | 2019/03/01 | 2 months ago | File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html). |
| [gut](https://github.com/1set/gut) | 22 | 3 | 2019/10/05 | 1 year ago | Simple and reliable package for file operations like copy/move/diff/list on files, directories and symbolic links. |
| [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) | 15 | 3 | 2018/10/16 | 2 years ago | Copy files for humans. |
| [todotxt](https://github.com/1set/todotxt) | 12 | 3 | 2020/11/06 | 2 weeks ago | Go library for Gina Trapani's [*todo.txt*](http://todotxt.org/) files, supports parsing and manipulating of task lists in the [*todo.txt* format](https://github.com/todotxt/todo.txt). |
| [higgs](https://github.com/dastoori/higgs) | 8 | 1 | 2020/12/13 | 2 weeks ago | A tiny cross-platform Go library to hide/unhide files and directories. |
| [pathtype](https://github.com/jonchun/pathtype) | 8 | 2 | 2021/08/03 | 6 months ago | Treat paths as their own type instead of using strings. |
**[⬆ back to top](#awesome-go-info)**

# Financial
        
*Packages for accounting and finance.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [ticker](https://github.com/achannarasappa/ticker) | 4095 | 63 | 2021/01/24 | 2 weeks ago | Terminal stock watcher and stock position tracker. |
| [decimal](https://github.com/shopspring/decimal) | 3703 | 63 | 2015/02/25 | 1 month ago | Arbitrary-precision fixed-point decimal numbers. |
| [go-money](https://github.com/Rhymond/go-money) | 1045 | 17 | 2017/03/20 | 1 week ago | Implementation of Fowler's Money pattern. |
| [accounting](https://github.com/leekchan/accounting) | 704 | 15 | 2015/08/10 | 1 month ago | money and currency formatting for golang. |
| [techan](https://github.com/sdcoffey/techan) | 589 | 48 | 2017/03/08 | 2 months ago | Technical analysis library with advanced market analysis and trading strategies. |
| [go-finance](https://github.com/FlashBoys/go-finance) | 535 | 27 | 2016/02/28 | 3 years ago | Comprehensive financial markets data in Go. |
| [ach](https://github.com/moov-io/ach) | 284 | 19 | 2016/12/14 | 2 weeks ago | A reader, writer, and valdiator for Automated Clearing House (ACH) files. |
| [currency](https://github.com/bojanz/currency) | 271 | 6 | 2020/04/16 | 2 weeks ago | Handles currency amounts, provides currency information and formatting. |
| [orderbook](https://github.com/i25959341/orderbook) | 243 | 20 | 2018/04/24 | 9 months ago | Matching Engine for Limit Order Book in Golang. |
| [go-finance](https://github.com/alpeb/go-finance) | 122 | 8 | 2017/06/01 | 2 months ago | Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations. |
| [transaction](https://github.com/claygod/transaction) | 102 | 10 | 2017/10/11 | 7 months ago | Embedded transactional database of accounts, running in multithreaded mode. |
| [ofxgo](https://github.com/aclindsa/ofxgo) | 96 | 10 | 2015/11/08 | 4 months ago | Query OFX servers and/or parse the responses (with example command-line client). |
| [vat](https://github.com/dannyvankooten/vat) | 86 | 3 | 2016/06/18 | 2 weeks ago | VAT number validation & EU VAT rates. |
| [sleet](https://github.com/BoltApp/sleet) | 82 | 54 | 2019/11/13 | 1 week ago | One unified interface for multiple Payment Service Providers (PsP) to process online payment. |
| [go-finnhub](https://github.com/m1/go-finnhub) | 64 | 6 | 2020/01/13 | 2 years ago | Client for stock market, forex and crypto data from finnhub.io. Access real-time financial market data from 60+ stock exchanges, 10 forex brokers, and 15+ crypto exchanges. |
| [currency](https://github.com/bnkamalesh/currency) | 44 | 6 | 2017/05/09 | 3 months ago | High performant & accurate currency computation package. |
| [fastme](https://github.com/newity/fastme) | 27 | 4 | 2020/10/29 | 4 months ago | Fast extensible matching engine Go implementation. |
| [payme](https://github.com/jovandeginste/payme) | 10 | 1 | 2021/05/03 | 2 months ago | QR code generator (ASCII & PNG) for SEPA payments. |
| [go-finance](https://github.com/pieterclaerhout/go-finance) | 6 | 2 | 2019/09/30 | 2 years ago | Module to fetch exchange rates, check VAT numbers via VIES and check IBAN bank account numbers. |
**[⬆ back to top](#awesome-go-info)**

# Forms
        
*Libraries for working with forms.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [nosurf](https://github.com/justinas/nosurf) | 1248 | 35 | 2013/08/22 | 1 year ago | CSRF protection middleware for Go. |
| [binding](https://github.com/mholt/binding) | 786 | 32 | 2014/05/20 | 3 years ago | Binds form and JSON data from net/http Request to struct. |
| [csrf](https://github.com/gorilla/csrf) | 748 | 23 | 2015/08/03 | 3 weeks ago | CSRF protection for Go web applications & services. |
| [form](https://github.com/go-playground/form) | 508 | 13 | 2016/05/26 | 7 months ago | Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support. |
| [conform](https://github.com/leebenson/conform) | 248 | 5 | 2016/01/05 | 4 months ago | Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags. |
| [formam](https://github.com/monoculum/formam) | 166 | 5 | 2014/10/25 | 4 months ago | decode form's values into a struct. |
| [forms](https://github.com/albrow/forms) | 129 | 8 | 2014/08/07 | 4 years ago | Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files. |
| [qs](https://github.com/sonh/qs) | 60 | 3 | 2020/10/02 | 8 months ago | Go module for encoding structs into URL query parameters. |
| [bind](https://github.com/robfig/bind) | 27 | 4 | 2014/08/06 | 7 years ago | Bind form data to any Go values. |
| [queryparam](https://github.com/TomWright/queryparam) | 12 | 2 | 2018/06/14 | 1 year ago | Decode `url.Values` into usable struct values of standard or custom types. |
**[⬆ back to top](#awesome-go-info)**

# Functional
        
*Packages to support functional programming in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-underscore](https://github.com/tobyhede/go-underscore) | 1225 | 30 | 2014/07/02 | 3 years ago | Useful collection of helpfully functional Go collection utilities. |
| [fpGo](https://github.com/TeaEntityLab/fpGo) | 230 | 7 | 2018/05/24 | 1 month ago | Monad, Functional Programming features for Golang. |
| [gofp](https://github.com/rbrahul/gofp) | 106 | 3 | 2021/02/19 | 11 months ago | A lodash like powerful utility library for Golang. |
| [fuego](https://github.com/seborama/fuego) | 103 | 3 | 2018/11/05 | 1 year ago | Functional Experiment in Go. |
**[⬆ back to top](#awesome-go-info)**

# Game Development
        
*Awesome game development libraries.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [ebiten](https://github.com/hajimehoshi/ebiten) | 5987 | 122 | 2013/06/16 | 1 week ago | dead simple 2D game library in Go. |
| [leaf](https://github.com/name5566/leaf) | 4285 | 321 | 2014/08/04 | 7 months ago | Lightweight game server framework. |
| [pixel](https://github.com/faiface/pixel) | 3808 | 100 | 2016/11/19 | 4 months ago | Hand-crafted 2D game library in Go. |
| [goworld](https://github.com/xiaonanln/goworld) | 2043 | 131 | 2017/06/03 | 7 months ago | Scalable game server engine, featuring space-entity framework and hot-swapping. |
| [nano](https://github.com/lonng/nano) | 1934 | 67 | 2017/08/02 | 7 months ago | Lightweight, facility, high performance golang based game server framework. |
| [engine](https://github.com/g3n/engine) | 1823 | 83 | 2017/03/07 | 2 weeks ago | Go 3D Game Engine. |
| [go-sdl2](https://github.com/veandco/go-sdl2) | 1698 | 45 | 2013/06/05 | 3 weeks ago | Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/). |
| [engo](https://github.com/EngoEngine/engo) | 1468 | 48 | 2014/11/12 | 2 months ago | Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm. |
| [pitaya](https://github.com/topfreegames/pitaya) | 1321 | 73 | 2018/03/19 | 2 weeks ago | Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. |
| [termloop](https://github.com/JoelOtter/termloop) | 1261 | 30 | 2015/05/23 | 6 months ago | Terminal-based game engine for Go, built on top of Termbox. |
| [gonet](https://github.com/xtaci/gonet) | 1165 | 135 | 2013/04/11 | 4 years ago | Game server skeleton implemented with golang. |
| [oak](https://github.com/oakmound/oak) | 1009 | 44 | 2017/07/15 | 2 weeks ago | Pure Go game engine. |
| [raylib-go](https://github.com/gen2brain/raylib-go) | 727 | 17 | 2017/01/27 | 4 weeks ago | Go bindings for [raylib](https://www.raylib.com/), a simple and easy-to-use library to learn videogames programming. |
| [engine](https://github.com/azul3d/engine) | 528 | 25 | 2016/02/29 | 3 months ago | 3D game engine written in Go. |
| [go-astar](https://github.com/beefsack/go-astar) | 481 | 10 | 2014/05/28 | 2 weeks ago | Go implementation of the A\* path finding algorithm. |
| [go3d](https://github.com/ungerik/go3d) | 221 | 9 | 2011/06/27 | 1 month ago | Performance oriented 2D/3D math package for Go. |
| [prototype](https://github.com/gonutz/prototype) | 68 | 3 | 2015/03/04 | 2 months ago | Cross-platform (Windows/Linux/Mac) library for creating desktop games using a minimal API. |
| [tile](https://github.com/kelindar/tile) | 41 | 1 | 2020/08/19 | 1 month ago | Data-oriented and cache-friendly 2D Grid library (TileMap), includes pathfinding, observers and import/export. |
**[⬆ back to top](#awesome-go-info)**

# Generation and Generics
        
*Tools to enhance the language with features like generics via code generation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-linq](https://github.com/ahmetb/go-linq) | 2845 | 74 | 2013/12/19 | 2 months ago | .NET LINQ-like query methods for Go. |
| [jennifer](https://github.com/dave/jennifer) | 2301 | 32 | 2016/12/04 | 1 month ago | Generate arbitrary Go code without templates. |
| [gen](https://github.com/clipperhouse/gen) | 1351 | 33 | 2013/10/13 | 2 years ago | Code generation tool for ‘generics’-like functionality. |
| [goderive](https://github.com/awalterschulze/goderive) | 959 | 17 | 2017/02/10 | 1 month ago | Derives functions from input types. |
| [gowrap](https://github.com/hexdigest/gowrap) | 565 | 12 | 2018/09/15 | 2 weeks ago | Generate decorators for Go interfaces using simple templates. |
| [interfaces](https://github.com/rjeczalik/interfaces) | 330 | 6 | 2015/12/06 | 9 months ago | Command line tool for generating interface definitions. |
| [go-enum](https://github.com/abice/go-enum) | 280 | 4 | 2017/08/10 | 2 weeks ago | Code generation for enums from code comments. |
| [goverter](https://github.com/jmattheis/goverter) | 101 | 3 | 2021/03/09 | 1 month ago | Generate converters by defining an interface. |
| [pkgreflect](https://github.com/ungerik/pkgreflect) | 99 | 6 | 2012/09/03 | 4 years ago | Go preprocessor for package scoped reflection. |
| [efaceconv](https://github.com/t0pep0/efaceconv) | 51 | 5 | 2016/11/18 | 4 years ago | Code generation tool for high performance conversion from interface{} to immutable type without allocations. |
| [gotype](https://github.com/wzshiming/gotype) | 36 | 4 | 2017/12/05 | 6 months ago | Golang source code parsing, usage like reflect package. |
| [GENERIS](https://github.com/senselogic/GENERIS) | 31 | 1 | 2019/03/10 | 6 months ago | Code generation tool providing generics, free-form macros, conditional compilation and HTML templating. |
| [go-xray](https://github.com/pieterclaerhout/go-xray) | 22 | 4 | 2019/10/01 | 2 years ago | Helpers for making the use of reflection easier. |
| [typeregistry](https://github.com/xiaoxin01/typeregistry) | 12 | 1 | 2020/01/14 | 2 years ago | A library to create type dynamically. |
**[⬆ back to top](#awesome-go-info)**

# Geographic
        
*Geographic tools and servers*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [tile38](https://github.com/tidwall/tile38) | 7954 | 209 | 2016/03/04 | 1 month ago | Geolocation DB with spatial index and realtime geofencing. |
| [geo](https://github.com/golang/geo) | 1320 | 79 | 2014/12/03 | 2 months ago | S2 geometry library in Go. |
| [mbtileserver](https://github.com/consbio/mbtileserver) | 304 | 14 | 2014/11/01 | 2 weeks ago | A simple Go-based server for map tiles stored in mbtiles format. |
| [osm](https://github.com/paulmach/osm) | 190 | 12 | 2016/02/02 | 5 months ago | Library for reading, writing and working with OpenStreetMap data and APIs. |
| [wgs84](https://github.com/wroge/wgs84) | 69 | 1 | 2019/06/08 | 2 months ago | Library for Coordinate Conversion and Transformation (ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM). |
| [geoserver](https://github.com/hishamkaram/geoserver) | 67 | 2 | 2018/03/26 | 2 months ago | geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API. |
| [godal](https://github.com/airbusgeo/godal) | 61 | 4 | 2021/02/05 | 1 week ago | Go wrapper for GDAL. |
| [simplefeatures](https://github.com/peterstace/simplefeatures) | 42 | 4 | 2019/06/07 | 1 week ago | simplesfeatures is a 2D geometry library that provides Go types that model geometries, as well as algorithms that operate on them. |
| [gismanager](https://github.com/hishamkaram/gismanager) | 41 | 0 | 2018/09/29 | 3 years ago | Publish Your GIS Data(Vector Data) to PostGIS and Geoserver. |
| [pbf](https://github.com/maguro/pbf) | 31 | 4 | 2017/09/18 | 10 months ago | OpenStreetMap PBF golang encoder/decoder. |
| [s2-geojson](https://github.com/pantrif/s2-geojson) | 16 | 1 | 2020/03/27 | 1 year ago | Convert geojson to s2 cells & demonstrating some S2 geometry features on map. |
**[⬆ back to top](#awesome-go-info)**

# Go Compilers
        
*Tools for compiling Go to other languages.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gopherjs](https://github.com/gopherjs/gopherjs) | 10864 | 250 | 2013/08/27 | 1 month ago | Compiler from Go to JavaScript. |
| [tardisgo](https://github.com/tardisgo/tardisgo) | 416 | 30 | 2014/01/08 | 5 years ago | Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler. |
| [c4go](https://github.com/Konstantin8105/c4go) | 294 | 18 | 2018/03/28 | 3 months ago | Transpile C code to Go code. |
| [esp32-transpiler](https://github.com/andygeiss/esp32-transpiler) | 38 | 3 | 2018/03/14 | 7 months ago | Transpile Go into Arduino code. |
| [f4go](https://github.com/Konstantin8105/f4go) | 31 | 5 | 2018/07/08 | 2 months ago | Transpile FORTRAN 77 code to Go code. |
**[⬆ back to top](#awesome-go-info)**

# Goroutines
        
*Tools for managing and working with Goroutines.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [ants](https://github.com/panjf2000/ants) | 7600 | 160 | 2018/05/19 | 2 weeks ago | A high-performance and low-cost goroutine pool in Go. |
| [tunny](https://github.com/Jeffail/tunny) | 2834 | 74 | 2014/04/02 | 2 weeks ago | Goroutine pool for golang. |
| [goworker](https://github.com/benmanns/goworker) | 2630 | 73 | 2013/07/22 | 2 months ago | goworker is a Go-based background worker. |
| [workerpool](https://github.com/gammazero/workerpool) | 724 | 16 | 2016/05/17 | 3 months ago | Goroutine pool that limits the concurrency of task execution, not the number of tasks queued. |
| [grpool](https://github.com/ivpusic/grpool) | 675 | 29 | 2015/07/22 | 3 years ago | Lightweight Goroutine pool. |
| [pool](https://github.com/go-playground/pool) | 670 | 15 | 2015/10/28 | 7 months ago | Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation. |
| [pond](https://github.com/alitto/pond) | 446 | 11 | 2020/03/21 | 1 month ago | Minimalistic and High-performance goroutine worker pool written in Go. |
| [gowp](https://github.com/xxjwxc/gowp) | 360 | 18 | 2019/09/14 | 9 months ago | gowp is concurrency limiting goroutine pool. |
| [go-floc](https://github.com/workanator/go-floc) | 225 | 7 | 2017/07/03 | 6 months ago | Orchestrate goroutines with ease. |
| [go-flow](https://github.com/kamildrazkiewicz/go-flow) | 184 | 11 | 2016/09/25 | 2 years ago | Control goroutines execution order. |
| [semaphore](https://github.com/marusama/semaphore) | 139 | 2 | 2017/11/22 | 10 months ago | Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations). |
| [go-workers](https://github.com/catmullet/go-workers) | 133 | 4 | 2020/10/06 | 1 month ago | Easily and safely run workers for large data processing pipelines. |
| [artifex](https://github.com/mborders/artifex) | 126 | 7 | 2018/10/31 | 1 year ago | Simple in-memory job queue for Golang using worker-based dispatching. |
| [errgroup](https://github.com/neilotoole/errgroup) | 111 | 3 | 2020/06/26 | 4 months ago | Drop-in alternative to `sync/errgroup`, limited to a pool of N worker goroutines. |
| [async](https://github.com/StudioSol/async) | 99 | 13 | 2017/06/30 | 1 year ago | A safe way to execute functions asynchronously, recovering them in case of panic. |
| [cyclicbarrier](https://github.com/marusama/cyclicbarrier) | 93 | 2 | 2018/01/11 | 1 year ago | CyclicBarrier for golang. |
| [semaphore](https://github.com/kamilsk/semaphore) | 90 | 1 | 2016/10/08 | 1 year ago | Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context. |
| [gpool](https://github.com/sherifabdlnaby/gpool) | 83 | 1 | 2018/12/03 | 2 years ago | manages a resizeable pool of context-aware goroutines to bound concurrency. |
| [worker-pool](https://github.com/vardius/worker-pool) | 81 | 5 | 2017/10/04 | 1 year ago | goworker is a Go simple async worker pool. |
| [gollback](https://github.com/vardius/gollback) | 77 | 1 | 2019/05/11 | 1 month ago | asynchronous simple function utilities, for managing execution of closures and callbacks. |
| [Hunch](https://github.com/AaronJan/Hunch) | 74 | 1 | 2019/06/05 | 1 year ago | Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive. |
| [threadpool](https://github.com/shettyh/threadpool) | 67 | 2 | 2017/09/06 | 1 year ago | Golang threadpool implementation. |
| [routine](https://github.com/x-mod/routine) | 48 | 3 | 2019/03/04 | 1 year ago | go routine control with context, support: Main, Go, Pool and some useful Executors. |
| [goccm](https://github.com/zenthangplus/goccm) | 42 | 1 | 2019/08/16 | 4 months ago | Go Concurrency Manager package limits the number of goroutines that allowed to run concurrently. |
| [nursery](https://github.com/arunsworld/nursery) | 39 | 4 | 2019/11/23 | 7 months ago | Structured concurrency in Go. |
| [async](https://github.com/reugn/async) | 37 | 2 | 2019/12/28 | 2 months ago | An alternative sync library for Go (Future, Promise, Locks). |
| [kyoo](https://github.com/dirkaholic/kyoo) | 37 | 2 | 2020/01/06 | 1 year ago | Provides an unlimited job queue and concurrent worker pools. |
| [parallel-fn](https://github.com/rafaeljesus/parallel-fn) | 32 | 3 | 2017/06/18 | 4 years ago | Run functions in parallel. |
| [gohive](https://github.com/loveleshsharma/gohive) | 29 | 3 | 2019/05/31 | 2 months ago | A highly performant and easy to use Goroutine pool for Go. |
| [go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup) | 26 | 2 | 2018/08/08 | 2 years ago | Like `sync.WaitGroup` with error handling and concurrency control. |
| [go-trylock](https://github.com/subchen/go-trylock) | 24 | 1 | 2018/04/26 | 9 months ago | TryLock support on read-write lock for Golang. |
| [stl](https://github.com/ssgreg/stl) | 23 | 2 | 2018/06/19 | 1 year ago | Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism. |
| [channelify](https://github.com/ddelizia/channelify) | 17 | 1 | 2020/10/05 | 11 months ago | Transform your function to return channels for easy and powerful parallel processing. |
| [gowl](https://github.com/hamed-yousefi/gowl) | 14 | 1 | 2021/04/12 | 6 months ago | Gowl is a process management and process monitoring tool at once. An infinite worker pool gives you the ability to control the pool and processes and monitor their status. |
| [conexec](https://github.com/ITcathyh/conexec) | 12 | 2 | 2019/12/24 | 1 year ago | A concurrent toolkit to help execute funcs concurrently in an efficient and safe way. It supports specifying the overall timeout to avoid blocking and uses goroutine pool to improve efficiency. |
| [queue](https://github.com/AnikHasibul/queue) | 11 | 0 | 2018/12/21 | 2 years ago | Gives you a `sync.WaitGroup` like queue group accessibility. Helps you to throttle and limit goroutines, wait for the end of the all goroutines and much more. |
| [execpool](https://github.com/hexdigest/execpool) | 8 | 1 | 2021/06/17 | 7 months ago | A pool built around exec.Cmd that spins up a given number of processes in advance and attaches stdin and stdout to them when needed. Very similar to FastCGI or Apache Prefork MPM but works for any command. |
| [hands](https://github.com/duanckham/hands) | 8 | 1 | 2020/04/04 | 1 year ago | A process controller used to control the execution and return strategies of multiple goroutines. |
| [concurrency-limiter](https://github.com/vivek-ng/concurrency-limiter) | 7 | 1 | 2020/11/22 | 1 year ago | Concurrency limiter with support for timeouts , dynamic priority and context cancellation of goroutines. |
| [go-tools](https://github.com/nikhilsaraf/go-tools) | 6 | 2 | 2018/11/14 | 2 years ago | Manage a pool of goroutines using this lightweight library with a simple API. |
| [breaker](https://github.com/kamilsk/breaker) | 3 | 0 | 2021/07/11 | 7 months ago | Flexible mechanism to make execution flow interruptible. |
**[⬆ back to top](#awesome-go-info)**

# GUI
        
*Interaction*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [fyne](https://github.com/fyne-io/fyne) | 15580 | 229 | 2018/02/04 | 1 week ago | Cross platform native GUIs designed for Go based on Material Design. Supports: Linux, macOS, Windows, BSD, iOS and Android. |
| [webview](https://github.com/webview/webview) | 9497 | 225 | 2017/08/19 | 2 weeks ago | Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux). |
| [qt](https://github.com/therecipe/qt) | 9041 | 319 | 2014/11/19 | 1 year ago | Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi). |
| [ui](https://github.com/andlabs/ui) | 8064 | 368 | 2014/02/17 | 7 months ago | Platform-native GUI library for Go. Cross platform. |
| [robotgo](https://github.com/go-vgo/robotgo) | 7219 | 236 | 2016/09/26 | 1 week ago | Go Native cross-platform GUI system automation. Control the mouse, keyboard and other. |
| [walk](https://github.com/lxn/walk) | 5912 | 261 | 2010/09/16 | 7 months ago | Windows application library kit for Go. |
| [go-app](https://github.com/maxence-charriere/go-app) | 5837 | 152 | 2016/10/12 | 1 month ago | Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. |
| [go-astilectron](https://github.com/asticode/go-astilectron) | 4223 | 134 | 2017/04/22 | 2 months ago | Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron). |
| [go-sciter](https://github.com/sciter-sdk/go-sciter) | 2322 | 130 | 2015/10/15 | 3 months ago | Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform. |
| [systray](https://github.com/getlantern/systray) | 2129 | 66 | 2014/11/12 | 9 months ago | Cross platform Go library to place an icon and menu in the notification area. |
| [gotk3](https://github.com/gotk3/gotk3) | 1682 | 65 | 2015/08/13 | 1 month ago | Go bindings for GTK3. |
| [gosx-notifier](https://github.com/deckarep/gosx-notifier) | 559 | 16 | 2013/11/25 | 1 year ago | OSX Desktop Notifications library for Go. |
| [gowd](https://github.com/dtylman/gowd) | 343 | 26 | 2017/03/29 | 2 years ago | Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform. |
| [trayhost](https://github.com/shurcooL/trayhost) | 221 | 7 | 2014/04/25 | 2 years ago | Cross-platform Go library to place an icon in the host operating system's taskbar. |
| [zenity](https://github.com/ncruces/zenity) | 219 | 6 | 2019/12/10 | 4 months ago | Cross-platform Go library and CLI to create simple dialogs that interact graphically with the user. |
| [go-appindicator](https://github.com/dawidd6/go-appindicator) | 19 | 5 | 2019/05/04 | 1 year ago | Go bindings for libappindicator3 C library. |
| [activity-tracker](https://github.com/prashantgupta24/activity-tracker) | 14 | 3 | 2019/03/12 | 2 years ago | OSX library to notify about any (pluggable) activity on your machine. |
| [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) | 12 | 4 | 2019/03/30 | 2 years ago | OSX Sleep/Wake notifications in golang. |
**[⬆ back to top](#awesome-go-info)**

# Hardware
        
*Libraries, tools, and tutorials for interacting with hardware.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [arduino-cli](https://github.com/arduino/arduino-cli) | 3237 | 104 | 2018/08/08 | 1 week ago | Official Arduino CLI and library. Can run standalone, or be incorporated into larger Go projects. |
| [go-rpio](https://github.com/stianeikeland/go-rpio) | 1800 | 61 | 2013/07/30 | 3 weeks ago | GPIO for Go, doesn't require cgo. |
| [ghw](https://github.com/jaypipes/ghw) | 1102 | 27 | 2017/05/26 | 2 weeks ago | Golang hardware discovery/inspection library. |
| [emgo](https://github.com/ziutek/emgo) | 959 | 35 | 2014/07/09 | 2 months ago | Go-like language for programming embedded systems (e.g. STM32 MCU). |
| [sysinfo](https://github.com/zcalusic/sysinfo) | 318 | 14 | 2016/08/22 | 2 weeks ago | A pure Go library providing Linux OS / kernel / hardware system information. |
| [goroslib](https://github.com/aler9/goroslib) | 157 | 12 | 2020/01/19 | 1 month ago | Robot Operating System (ROS) library for Go. |
| [go-osc](https://github.com/hypebeast/go-osc) | 140 | 8 | 2013/08/26 | 1 week ago | Open Sound Control (OSC) bindings for Go. |
| [joystick](https://github.com/0xcafed00d/joystick) | 30 | 2 | 2015/07/24 | 2 years ago | a polled API to read the state of an attached joystick. |
**[⬆ back to top](#awesome-go-info)**

# Images
        
*Libraries for manipulating images.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gocv](https://github.com/hybridgroup/gocv) | 4602 | 141 | 2017/09/18 | 2 weeks ago | Go package for computer vision using OpenCV 3.3+. |
| [imaginary](https://github.com/h2non/imaginary) | 4206 | 75 | 2015/03/04 | 1 month ago | Fast and simple HTTP microservice for image resizing. |
| [imaging](https://github.com/disintegration/imaging) | 4097 | 78 | 2012/12/06 | 1 year ago | Simple Go image processing package. |
| [bild](https://github.com/anthonynsimon/bild) | 3468 | 73 | 2016/08/01 | 2 months ago | Collection of image processing algorithms in pure Go. |
| [gg](https://github.com/fogleman/gg) | 3249 | 90 | 2016/02/18 | 3 months ago | 2D rendering in pure Go. |
| [ln](https://github.com/fogleman/ln) | 3032 | 93 | 2016/01/10 | 2 years ago | 3D line art rendering in Go. |
| [resize](https://github.com/nfnt/resize) | 2781 | 79 | 2012/08/02 | 1 year ago | Image resizing for Go with common interpolation methods. |
| [pt](https://github.com/fogleman/pt) | 1991 | 58 | 2015/01/23 | 2 years ago | Path tracing engine written in Go. |
| [svgo](https://github.com/ajstarks/svgo) | 1800 | 50 | 2010/03/05 | 3 months ago | Go Language Library for SVG generation. |
| [bimg](https://github.com/h2non/bimg) | 1788 | 37 | 2015/03/17 | 2 weeks ago | Small package for fast and efficient image processing using libvips. |
| [picfit](https://github.com/thoas/picfit) | 1606 | 52 | 2014/12/06 | 3 weeks ago | An image resizing server written in Go. |
| [smartcrop](https://github.com/muesli/smartcrop) | 1600 | 33 | 2014/04/07 | 3 weeks ago | Finds good crops for arbitrary images and crop sizes. |
| [gift](https://github.com/disintegration/gift) | 1497 | 48 | 2014/07/12 | 1 year ago | Package of image processing filters. |
| [imagick](https://github.com/gographics/imagick) | 1397 | 52 | 2013/04/30 | 5 months ago | Go binding to ImageMagick's MagickWand C API. |
| [gowitness](https://github.com/sensepost/gowitness) | 1354 | 38 | 2017/10/31 | 2 weeks ago | Screenshoting webpages using go and headless chrome on command line. |
| [go-opencv](https://github.com/go-opencv/go-opencv) | 1261 | 62 | 2013/12/09 | 2 years ago | Go bindings for OpenCV. |
| [geopattern](https://github.com/pravj/geopattern) | 1164 | 22 | 2014/10/22 | 3 years ago | Create beautiful generative image patterns from a string. |
| [stegify](https://github.com/DimitarPetrov/stegify) | 982 | 22 | 2018/11/29 | 1 year ago | Go tool for LSB steganography, capable of hiding any file within an image. |
| [canvas](https://github.com/tdewolff/canvas) | 905 | 20 | 2017/05/20 | 2 weeks ago | Vector graphics to PDF, SVG or rasterized image. |
| [govips](https://github.com/davidbyttow/govips) | 628 | 13 | 2016/12/25 | 1 week ago | A lightning fast image processing and resizing library for Go. |
| [image2ascii](https://github.com/qeesung/image2ascii) | 627 | 9 | 2018/10/20 | 6 months ago | Convert image to ASCII. |
| [draft](https://github.com/lucasepe/draft) | 519 | 11 | 2020/06/05 | 5 months ago | Generate High Level Microservice Architecture diagrams for GraphViz using simple YAML syntax. |
| [govatar](https://github.com/o1egl/govatar) | 478 | 10 | 2016/01/18 | 11 months ago | Library and CMD tool for generating funny avatars. |
| [goimagehash](https://github.com/corona10/goimagehash) | 467 | 11 | 2017/07/28 | 1 year ago | Go Perceptual image hashing package. |
| [mort](https://github.com/aldor007/mort) | 444 | 18 | 2017/11/19 | 3 months ago | Storage and image processing server written in Go. |
| [go-nude](https://github.com/koyachi/go-nude) | 343 | 16 | 2014/05/02 | 3 years ago | Nudity detection with Go. |
| [rez](https://github.com/bamiaux/rez) | 203 | 10 | 2014/01/16 | 4 years ago | Image resizing in pure Go and SIMD. |
| [darkroom](https://github.com/gojek/darkroom) | 184 | 9 | 2019/07/01 | 7 months ago | An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency. |
| [mergi](https://github.com/noelyahan/mergi) | 162 | 7 | 2018/09/24 | 1 year ago | Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate). |
| [gltf](https://github.com/qmuntal/gltf) | 148 | 5 | 2019/01/15 | 1 week ago | Efficient and robust glTF 2.0 reader, writer and validator. |
| [img](https://github.com/hawx/img) | 138 | 5 | 2012/07/28 | 6 years ago | Selection of image manipulation tools. |
| [steganography](https://github.com/auyer/steganography) | 127 | 6 | 2018/05/21 | 6 months ago | Pure Go Library for LSB steganography. |
| [go-cairo](https://github.com/ungerik/go-cairo) | 119 | 5 | 2012/08/22 | 1 month ago | Go binding for the cairo graphics library. |
| [cameron](https://github.com/aofei/cameron) | 78 | 4 | 2018/05/05 | 3 months ago | An avatar generator for Go. |
| [go-gd](https://github.com/bolknote/go-gd) | 53 | 4 | 2011/05/12 | 3 years ago | Go binding for GD library. |
| [gridder](https://github.com/shomali11/gridder) | 51 | 4 | 2020/04/10 | 4 months ago | A Grid based 2D Graphics library. |
| [go-webp](https://github.com/kolesa-team/go-webp) | 42 | 5 | 2020/02/18 | 5 months ago | Library for encode and decode webp pictures, using libwebp. |
| [goimghdr](https://github.com/corona10/goimghdr) | 38 | 1 | 2018/02/25 | 2 years ago | The imghdr module determines the type of image contained in a file for Go. |
| [webp-server](https://github.com/mehdipourfar/webp-server) | 33 | 1 | 2020/11/22 | 1 year ago | Simple and minimal image server capable of storing, resizing, converting and caching images. |
| [tga](https://github.com/ftrvxmtrx/tga) | 29 | 3 | 2012/10/08 | 6 years ago | Package tga is a TARGA image format decoder/encoder. |
| [go-webcolors](https://github.com/jyotiska/go-webcolors) | 26 | 2 | 2014/04/24 | 6 years ago | Port of webcolors library from Python to Go. |
| [mpo](https://github.com/donatj/mpo) | 8 | 2 | 2015/04/14 | 1 year ago | Decoder and conversion tool for MPO 3D Photos. |
| [scout](https://github.com/jonoton/scout) | 2 | 1 | 2020/09/25 | 4 months ago | Scout is a standalone open source software solution for DIY video security. |
**[⬆ back to top](#awesome-go-info)**

# IoT
        
*Libraries for programming devices of the IoT.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [flogo](https://github.com/TIBCOSoftware/flogo) | 1918 | 155 | 2016/07/10 | 1 year ago | Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. |
| [mainflux](https://github.com/mainflux/mainflux) | 1685 | 101 | 2015/07/06 | 1 week ago | Industrial IoT Messaging and Device Management Server. |
| [gatt](https://github.com/paypal/gatt) | 1007 | 55 | 2014/04/23 | 1 year ago | Gatt is a Go package for building Bluetooth Low Energy peripherals. |
| [heedy](https://github.com/heedy/heedy) | 323 | 24 | 2015/01/16 | 2 weeks ago | Open-Source Platform for Quantified Self & IoT. |
| [devices](https://github.com/goiot/devices) | 251 | 16 | 2016/05/30 | 5 years ago | Suite of libraries for IoT devices, experimental for x/exp/io. |
| [sensorbee](https://github.com/sensorbee/sensorbee) | 210 | 19 | 2016/02/19 | 2 years ago | Lightweight stream processing engine for IoT. |
| [huego](https://github.com/amimof/huego) | 197 | 4 | 2017/05/16 | 1 month ago | An extensive Philips Hue client library for Go. |
| [eywa](https://github.com/xcodersun/eywa) | 52 | 8 | 2016/02/20 | 4 years ago | Project Eywa is essentially a connection manager that keeps track of connected devices. |
**[⬆ back to top](#awesome-go-info)**

# Job Scheduler
        
*Libraries for scheduling jobs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gocron](https://github.com/go-co-op/gocron) | 1553 | 21 | 2020/03/20 | 3 weeks ago | Easy and fluent Go job scheduling. This is an actively maintained fork of [jasonlvhit/gocron](https://github.com/jasonlvhit/gocron). |
| [jobrunner](https://github.com/bamzi/jobrunner) | 900 | 26 | 2015/10/21 | 1 year ago | Smart and featureful cron job scheduler with job queuing and live monitoring built in. |
| [gron](https://github.com/roylee0704/gron) | 883 | 15 | 2016/06/04 | 1 year ago | Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly. |
| [go-quartz](https://github.com/reugn/go-quartz) | 584 | 13 | 2019/04/14 | 1 month ago | Simple, zero-dependency scheduling library for Go. |
| [jobs](https://github.com/albrow/jobs) | 486 | 19 | 2015/02/09 | 3 years ago | Persistent and flexible background jobs library. |
| [scheduler](https://github.com/carlescere/scheduler) | 378 | 14 | 2015/02/03 | 1 year ago | Cronjobs scheduling made easy. |
| [go-cron](https://github.com/rk/go-cron) | 209 | 10 | 2011/04/15 | 2 years ago | Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. |
| [gronx](https://github.com/adhocore/gronx) | 183 | 3 | 2021/04/21 | 4 months ago | Cron expression parser, task runner and daemon consuming crontab like task list. |
| [clockwerk](https://github.com/onatm/clockwerk) | 116 | 2 | 2017/04/09 | 2 years ago | Go package to schedule periodic jobs using a simple, fluent syntax. |
| [leprechaun](https://github.com/kilgaloon/leprechaun) | 84 | 8 | 2018/04/08 | 3 months ago | Job scheduler that supports webhooks, crons and classic scheduling. |
| [tasks](https://github.com/madflojo/tasks) | 70 | 3 | 2019/12/24 | 1 week ago | An easy to use in-process scheduler for recurring tasks in Go. |
| [sched](https://github.com/romshark/sched) | 22 | 1 | 2021/06/19 | 7 months ago | A job scheduler with the ability to fast-forward time. |
| [cheek](https://github.com/datarootsio/cheek) | 21 | 3 | 2021/12/01 | 1 week ago | A simple crontab like scheduler that aims to offer a KISS approach to job scheduling. |
| [cronticker](https://github.com/krayzpipes/cronticker) | 2 | 1 | 2020/11/28 | 1 year ago | A ticker implementation to support cron schedules. |
**[⬆ back to top](#awesome-go-info)**

# JSON
        
*Libraries for working with JSON.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gjson](https://github.com/tidwall/gjson) | 9710 | 154 | 2016/08/11 | 1 week ago | Get a JSON value with one line of code. |
| [gojson](https://github.com/ChimeraCoder/gojson) | 2457 | 48 | 2012/12/27 | 6 months ago | Automatically generate Go (golang) struct definitions from example JSON. |
| [fastjson](https://github.com/valyala/fastjson) | 1445 | 26 | 2018/05/28 | 7 months ago | Fast JSON parser and validator for Go. No custom structs, no code generation, no reflection. |
| [ojg](https://github.com/ohler55/ojg) | 452 | 6 | 2020/04/12 | 1 month ago | Optimized JSON for Go is a high performance parser with a variety of additional JSON tools including JSONPath. |
| [kazaam](https://github.com/qntfy/kazaam) | 218 | 21 | 2016/07/19 | 7 months ago | API for arbitrary transformation of JSON documents. |
| [gojq](https://github.com/elgs/gojq) | 181 | 5 | 2015/12/30 | 1 year ago | JSON query in Golang. |
| [jsondiff](https://github.com/wI2L/jsondiff) | 149 | 2 | 2020/11/28 | 5 months ago | JSON diff library for Go based on RFC6902 (JSON Patch). |
| [jettison](https://github.com/wI2L/jettison) | 117 | 6 | 2019/08/30 | 1 month ago | Fast and flexible JSON encoder for Go. |
| [ajson](https://github.com/spyzhov/ajson) | 105 | 3 | 2019/03/07 | 2 weeks ago | Abstract JSON for golang with JSONPath support. |
| [gjo](https://github.com/skanehira/gjo) | 99 | 8 | 2019/02/23 | 10 months ago | Small utility to create JSON objects. |
| [jsongo](https://github.com/ricardolonga/jsongo) | 99 | 1 | 2015/08/07 | 4 months ago | Fluent API to make it easier to create Json objects. |
| [json2go](https://github.com/m-zajac/json2go) | 95 | 3 | 2017/06/10 | 2 months ago | Advanced JSON to Go struct conversion. Provides package that can parse multiple JSON documents and create struct to fit them all. |
| [jaydiff](https://github.com/yazgazan/jaydiff) | 84 | 2 | 2017/04/24 | 1 year ago | JSON diff utility written in Go. |
| [jsonf](https://github.com/miolini/jsonf) | 63 | 3 | 2015/05/25 | 1 year ago | Console tool for highlighted formatting and struct query fetching JSON. |
| [ujson](https://github.com/olvrng/ujson) | 53 | 1 | 2019/02/27 | 6 months ago | Fast and minimal JSON parser and transformer that works on unstructured JSON. |
| [go-respond](https://github.com/nicklaw5/go-respond) | 45 | 1 | 2017/03/12 | 4 months ago | Go package for handling common HTTP JSON responses. |
| [mp](https://github.com/sanbornm/mp) | 44 | 2 | 2014/06/15 | 5 years ago | Simple cli email parser. It currently takes stdin and outputs JSON. |
| [jsoncolor](https://github.com/neilotoole/jsoncolor) | 27 | 1 | 2021/09/13 | 4 months ago | Drop-in replacement for `encoding/json` that outputs colorized JSON. |
| [vjson](https://github.com/miladibra10/vjson) | 26 | 1 | 2021/04/29 | 3 months ago | Go package for validating JSON objects with declaring a JSON schema with fluent API. |
| [ask](https://github.com/simonnilsson/ask) | 14 | 1 | 2020/09/13 | 1 year ago | Easy access to nested values in maps and slices. Works in combination with encoding/json and other packages that "Unmarshal" arbitrary data into Go data-types. |
| [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) | 12 | 2 | 2016/07/08 | 5 years ago | Go bindings based on the JSON API errors reference. |
| [go-jsonerror](https://github.com/ddymko/go-jsonerror) | 11 | 2 | 2018/10/18 | 2 years ago | Go-JsonError is ment to allow us to easily create json response errors that follow the JsonApi spec. |
| [jscan](https://github.com/romshark/jscan) | 11 | 1 | 2022/01/08 | 3 weeks ago | High performance zero-allocation JSON iterator. |
| [dynjson](https://github.com/cocoonspace/dynjson) | 10 | 2 | 2020/05/06 | 4 months ago | Client-customizable JSON formats for dynamic APIs. |
| [jsonhal](https://github.com/RichardKnop/jsonhal) | 10 | 2 | 2016/01/15 | 1 year ago | Simple Go package to make custom structs marshal into HAL compatible JSON responses. |
| [mapslice-json](https://github.com/ake-persson/mapslice-json) | 9 | 1 | 2020/02/19 | 7 months ago | Go MapSlice for ordered marshal/ unmarshal of maps in JSON. |
| [epoch](https://github.com/vtopc/epoch) | 8 | 1 | 2019/12/15 | 10 months ago | Contains primitives for marshaling/unmarshaling Unix timestamp/epoch to/from build-in time.Time type in JSON. |
| [ej](https://github.com/lucassscaravelli/ej) | 7 | 2 | 2020/01/04 | 1 year ago | Write and read JSON from different sources succinctly. |
| [jzon](https://github.com/zerosnake0/jzon) | 6 | 1 | 2019/11/12 | 11 months ago | JSON library with standard compatible API/behavior. |
| [jsonic](https://github.com/sinhashubham95/jsonic) | 5 | 1 | 2021/01/09 | 1 year ago | Utilities to handle and query JSON without defining structs in a type safe manner. |
| [omg.jsonparser](https://github.com/dedalqq/omg.jsonparser) | 3 | 1 | 2021/07/08 | 4 months ago | Simple JSON parser with validation by condition via golang struct fields tags. |
**[⬆ back to top](#awesome-go-info)**

# Logging
        
*Libraries for generating and working with log files.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [logrus](https://github.com/sirupsen/logrus) | 19791 | 311 | 2013/10/16 | 2 weeks ago | Structured logger for Go. |
| [zap](https://github.com/uber-go/zap) | 14830 | 248 | 2016/02/18 | 1 week ago | Fast, structured, leveled logging in Go. |
| [zerolog](https://github.com/rs/zerolog) | 5828 | 57 | 2017/05/12 | 1 week ago | Zero-allocation JSON logger. |
| [go-spew](https://github.com/davecgh/go-spew) | 4850 | 67 | 2013/01/09 | 1 month ago | Implements a deep pretty printer for Go data structures to aid in debugging. |
| [glog](https://github.com/golang/glog) | 3108 | 90 | 2013/07/16 | 5 months ago | Leveled execution logs for Go. |
| [lumberjack](https://github.com/natefinch/lumberjack) | 3068 | 60 | 2014/06/14 | 2 weeks ago | Simple rolling logger, implements io.WriteCloser. |
| [tail](https://github.com/hpcloud/tail) | 2261 | 103 | 2013/02/05 | 1 month ago | Go package striving to emulate the features of the BSD tail program. |
| [seelog](https://github.com/cihub/seelog) | 1585 | 92 | 2011/11/17 | 2 years ago | Logging functionality with flexible dispatching, filtering, and formatting. |
| [log](https://github.com/apex/log) | 1203 | 35 | 2015/12/21 | 3 months ago | Structured logging package for Go. |
| [log15](https://github.com/inconshreveable/log15) | 1035 | 25 | 2014/05/20 | 3 months ago | Simple, powerful logging for Go. |
| [log](https://github.com/phuslu/log) | 414 | 17 | 2019/07/07 | 2 weeks ago | Structured Logging Made Easy. |
| [onelog](https://github.com/francoispqt/onelog) | 400 | 9 | 2018/05/06 | 2 years ago | Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenarios. Also, it is one of the logger with the lowest allocation. |
| [logxi](https://github.com/mgutz/logxi) | 348 | 10 | 2015/03/01 | 1 year ago | 12-factor app logger that is fast and makes you happy. |
| [logutils](https://github.com/hashicorp/logutils) | 306 | 268 | 2013/10/09 | 3 months ago | Utilities for slightly better logging in Go (Golang) extending the standard logger. |
| [log](https://github.com/go-playground/log) | 276 | 11 | 2016/02/07 | 2 years ago | Simple, configurable and scalable Structured Logging for Go. |
| [go-logger](https://github.com/apsdehal/go-logger) | 272 | 7 | 2014/09/26 | 2 years ago | Simple logger of Go Programs, with level handlers. |
| [httpretty](https://github.com/henvic/httpretty) | 248 | 6 | 2020/01/24 | 1 year ago | Pretty-prints your regular HTTP requests on your terminal for debugging (similar to http.DumpRequest). |
| [rollingwriter](https://github.com/arthurkiller/rollingwriter) | 219 | 8 | 2017/02/12 | 4 weeks ago | RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation. |
| [sqldb-logger](https://github.com/simukti/sqldb-logger) | 210 | 2 | 2019/11/02 | 1 week ago | A logger for Go SQL database driver without modify existing *sql.DB stdlib usage. |
| [logur](https://github.com/logur/logur) | 150 | 7 | 2018/12/09 | 1 year ago | An opinionated logger interface and collection of logging best practices with adapters and integrations for well-known libraries ([logrus](https://github.com/sirupsen/logrus), [go-kit log](https://github.com/go-kit/kit/tree/master/log), [zap](https://github.com/uber-go/zap), [zerolog](https://github.com/rs/zerolog), etc). |
| [logger](https://github.com/azer/logger) | 148 | 6 | 2014/09/30 | 2 months ago | Minimalistic logging library for Go. |
| [glg](https://github.com/kpango/glg) | 138 | 6 | 2017/06/21 | 1 month ago | glg is simple and fast leveled logging library for Go. |
| [xlog](https://github.com/rs/xlog) | 135 | 8 | 2015/10/22 | 1 year ago | Structured logger for `net/context` aware HTTP handlers with flexible dispatching. |
| [ozzo-log](https://github.com/go-ozzo/ozzo-log) | 118 | 12 | 2015/10/22 | 1 year ago | High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail). |
| [logvoyage](https://github.com/firstrow/logvoyage) | 90 | 5 | 2015/03/29 | 4 years ago | Full-featured logging saas written in golang. |
| [go-cronowriter](https://github.com/utahta/go-cronowriter) | 46 | 2 | 2017/02/04 | 11 months ago | Simple writer that rotate log files automatically based on current date and time, like cronolog. |
| [log](https://github.com/alexcesaro/log) | 45 | 6 | 2014/04/19 | 6 years ago | Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs. |
| [gologger](https://github.com/sadlil/gologger) | 40 | 6 | 2015/09/02 | 4 years ago | Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch. |
| [go-log](https://github.com/ian-kent/go-log) | 38 | 2 | 2014/05/02 | 3 years ago | Log4j implementation in Go. |
| [logex](https://github.com/chzyer/logex) | 38 | 9 | 2014/10/10 | 1 month ago | Golang log lib, supports tracking and level, wrap by standard log lib. |
| [noodlog](https://github.com/gyozatech/noodlog) | 35 | 6 | 2021/04/09 | 4 months ago | Parametrized JSON logging library which lets you obfuscate sensitive data and marshal any kind of content. No more printed pointers instead of values, nor escape chars for the JSON strings. |
| [go-log](https://github.com/siddontang/go-log) | 29 | 6 | 2014/05/18 | 3 years ago | Log lib supports level and multi handlers. |
| [journald](https://github.com/ssgreg/journald) | 29 | 3 | 2017/08/23 | 11 months ago | Go implementation of systemd Journal's native API for logging. |
| [logrusly](https://github.com/sebest/logrusly) | 27 | 5 | 2014/09/11 | 6 months ago | [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/). |
| [distillog](https://github.com/amoghe/distillog) | 25 | 2 | 2015/10/12 | 3 years ago | distilled levelled logging (think of it as stdlib + log levels). |
| [log](https://github.com/teris-io/log) | 24 | 2 | 2017/10/28 | 4 years ago | Structured log interface for Go cleanly separates logging facade from its implementation. |
| [mlog](https://github.com/jbrodriguez/mlog) | 24 | 1 | 2014/10/20 | 3 years ago | Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output. |
| [gomol](https://github.com/aphistic/gomol) | 18 | 2 | 2015/08/30 | 2 years ago | Multiple-output, structured logging for Go with extensible logging outputs. |
| [zkits-logger](https://github.com/edoger/zkits-logger) | 15 | 1 | 2020/03/31 | 1 month ago | A powerful zero-dependency JSON logger. |
| [glo](https://github.com/lajosbencz/glo) | 14 | 1 | 2019/01/19 | 3 years ago | PHP Monolog inspired logging facility with identical severity levels. |
| [logrusiowriter](https://github.com/cabify/logrusiowriter) | 13 | 90 | 2019/08/09 | 1 year ago | `io.Writer` implementation using [logrus](https://github.com/sirupsen/logrus) logger. |
| [logmatic](https://github.com/mborders/logmatic) | 12 | 2 | 2018/11/07 | 1 year ago | Colorized logger for Golang with dynamic log level configuration. |
| [go-log](https://github.com/subchen/go-log) | 11 | 2 | 2017/05/07 | 3 years ago | Simple and configurable Logging in Go, with level, formatters and writers. |
| [log](https://github.com/aerogo/log) | 9 | 2 | 2017/06/10 | 2 years ago | An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection). |
| [logdump](https://github.com/ewwwwwqm/logdump) | 9 | 2 | 2017/01/13 | 3 years ago | Package for multi-level logging. |
| [logo](https://github.com/mbndr/logo) | 9 | 2 | 2017/02/07 | 1 year ago | Golang logger to different configurable writers. |
| [go-log](https://github.com/pieterclaerhout/go-log) | 8 | 3 | 2019/10/01 | 1 year ago | A logging library with strack traces, object dumping and optional timestamps. |
| [kemba](https://github.com/clok/kemba) | 7 | 1 | 2020/07/13 | 1 month ago | A tiny debug logging tool inspired by [debug](https://github.com/visionmedia/debug), great for CLI tools and applications. |
| [xlog](https://github.com/xfxdev/xlog) | 6 | 1 | 2016/05/05 | 3 years ago | Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format. |
| [log](https://github.com/structy/log) | 4 | 0 | 2022/01/26 | 2 weeks ago | A simple to use log system, minimalist but with features for debugging and differentiation of messages. |
| [yell](https://github.com/jfcg/yell) | 0 | 2 | 2021/02/07 | 6 months ago | Yet another minimalistic logging library. |
**[⬆ back to top](#awesome-go-info)**

# Machine Learning
        
*Libraries for Machine Learning.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [golearn](https://github.com/sjwhitworth/golearn) | 8194 | 429 | 2013/12/26 | 4 weeks ago | General Machine Learning library for Go. |
| [gorse](https://github.com/zhenghaoz/gorse) | 5165 | 51 | 2018/08/14 | 1 week ago | An offline recommender system backend based on collaborative filtering written in Go. |
| [gorgonia](https://github.com/gorgonia/gorgonia) | 4368 | 193 | 2016/09/14 | 2 weeks ago | graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. |
| [m2cgen](https://github.com/BayesWitnesses/m2cgen) | 2007 | 44 | 2019/01/13 | 1 week ago | A CLI tool to transpile trained classic ML models into a native Go code with zero dependencies, written in Python with Go language support. |
| [tfgo](https://github.com/galeone/tfgo) | 1884 | 60 | 2017/05/23 | 5 months ago | Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python. |
| [gosseract](https://github.com/otiai10/gosseract) | 1654 | 47 | 2013/10/11 | 2 months ago | Go package for OCR (Optical Character Recognition), by using Tesseract C++ library. |
| [goml](https://github.com/cdipaolo/goml) | 1303 | 74 | 2015/06/27 | 3 months ago | On-line Machine Learning in Go. |
| [eaopt](https://github.com/MaxHalford/eaopt) | 771 | 30 | 2016/01/31 | 10 months ago | An evolutionary optimization library. |
| [bayesian](https://github.com/jbrukh/bayesian) | 726 | 35 | 2011/11/23 | 1 year ago | Naive Bayesian Classification for Golang. |
| [CloudForest](https://github.com/ryanbressler/CloudForest) | 705 | 43 | 2012/10/22 | 1 year ago | Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go. |
| [gobrain](https://github.com/goml/gobrain) | 509 | 27 | 2014/04/29 | 1 year ago | Neural Networks written in go. |
| [ocrserver](https://github.com/otiai10/ocrserver) | 455 | 15 | 2015/11/15 | 6 months ago | A simple OCR API server, seriously easy to be deployed by Docker and Heroku. |
| [onnx-go](https://github.com/owulveryck/onnx-go) | 371 | 12 | 2018/08/28 | 3 months ago | Go Interface to Open Neural Network Exchange (ONNX). |
| [go-deep](https://github.com/patrikeh/go-deep) | 344 | 16 | 2017/12/09 | 2 weeks ago | A feature-rich neural network library in Go. |
| [regommend](https://github.com/muesli/regommend) | 299 | 16 | 2014/02/05 | 2 years ago | Recommendation & collaborative filtering engine. |
| [goptuna](https://github.com/c-bata/goptuna) | 206 | 9 | 2019/07/24 | 1 month ago | Bayesian optimization framework for black-box functions written in Go. Everything will be optimized. |
| [go-galib](https://github.com/thoj/go-galib) | 188 | 15 | 2009/11/30 | 6 years ago | Genetic Algorithms library written in Go / golang. |
| [goRecommend](https://github.com/timkaye11/goRecommend) | 183 | 10 | 2014/07/16 | 7 years ago | Recommendation Algorithms library written in Go. |
| [shield](https://github.com/eaigner/shield) | 149 | 11 | 2013/04/10 | 1 year ago | Bayesian text classifier with flexible tokenizers and storage backends for Go. |
| [goga](https://github.com/tomcraven/goga) | 147 | 10 | 2015/10/20 | 1 month ago | Genetic algorithm library for Go. |
| [go-fann](https://github.com/vksnk/go-fann) | 106 | 9 | 2011/03/10 | 7 years ago | Go bindings for Fast Artificial Neural Networks(FANN) library. |
| [goscore](https://github.com/asafschers/goscore) | 74 | 7 | 2017/08/19 | 2 years ago | Go Scoring API for PMML. |
| [gonet](https://github.com/dathoangnd/gonet) | 71 | 5 | 2020/01/11 | 1 year ago | Neural Network for Go. |
| [libsvm](https://github.com/datastream/libsvm) | 71 | 11 | 2012/07/31 | 5 years ago | libsvm golang version derived work based on LIBSVM 3.14. |
| [go-featureprocessing](https://github.com/nikolaydubina/go-featureprocessing) | 69 | 1 | 2020/12/18 | 1 month ago | Fast and convenient feature processing for low latency machine learning in Go. |
| [neural-go](https://github.com/schuyler/neural-go) | 62 | 3 | 2011/10/17 | 1 year ago | Multilayer perceptron network implemented in Go, with training via backpropagation. |
| [go-pr](https://github.com/daviddengcn/go-pr) | 61 | 7 | 2013/06/07 | 8 years ago | Pattern recognition package in Go lang. |
| [neat](https://github.com/jinyeom/neat) | 61 | 13 | 2016/11/17 | 3 years ago | Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT). |
| [fonet](https://github.com/Fontinalis/fonet) | 58 | 5 | 2017/10/03 | 8 months ago | A Deep Neural Network library written in Go. |
| [golinear](https://github.com/danieldk/golinear) | 44 | 6 | 2013/04/05 | 3 years ago | liblinear bindings for Go. |
| [Varis](https://github.com/Xamber/Varis) | 42 | 8 | 2017/10/10 | 3 years ago | Golang Neural Network. |
| [godist](https://github.com/e-dard/godist) | 32 | 4 | 2014/09/05 | 6 years ago | Various probability distributions, and associated methods. |
| [go-cluster](https://github.com/e-XpertSolutions/go-cluster) | 30 | 8 | 2017/10/04 | 3 years ago | Go implementation of the k-modes and k-prototypes clustering algorithms. |
| [evoli](https://github.com/khezen/evoli) | 21 | 6 | 2015/06/12 | 3 months ago | Genetic Algorithm and Particle Swarm Optimization library. |
| [gomind](https://github.com/surenderthakran/gomind) | 21 | 4 | 2017/10/19 | 3 years ago | A simplistic Neural Network Library in Go. |
| [randomForest](https://github.com/malaschitz/randomForest) | 19 | 4 | 2018/10/25 | 4 months ago | Easy to use Random Forest library for Go. |
| [ddt](https://github.com/sgrodriguez/ddt) | 17 | 1 | 2020/05/20 | 11 months ago | Dynamic decision tree, create trees defining customizable rules. |
| [probab](https://github.com/ThePaw/probab) | 17 | 2 | 2015/09/14 | 6 years ago | Probability distribution functions. Bayesian inference. Written in pure Go. |
**[⬆ back to top](#awesome-go-info)**

# Messaging
        
*Libraries that implement messaging systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [sarama](https://github.com/Shopify/sarama) | 8123 | 512 | 2013/07/05 | 2 weeks ago | Go library for Apache Kafka. |
| [gorush](https://github.com/appleboy/gorush) | 6070 | 191 | 2016/03/22 | 1 week ago | Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm). |
| [machinery](https://github.com/RichardKnop/machinery) | 5950 | 152 | 2015/04/05 | 2 weeks ago | Asynchronous task queue/job queue based on distributed message passing. |
| [centrifugo](https://github.com/centrifugal/centrifugo) | 5789 | 197 | 2015/03/31 | 2 weeks ago | Real-time messaging (Websockets or SockJS) server in Go. |
| [go-socket.io](https://github.com/googollee/go-socket.io) | 4462 | 134 | 2013/07/13 | 3 weeks ago | socket.io library for golang, a realtime application framework. |
| [benthos](https://github.com/Jeffail/benthos) | 3958 | 91 | 2016/03/22 | 1 week ago | A message streaming bridge between a range of protocols. |
| [nats.go](https://github.com/nats-io/nats.go) | 3769 | 164 | 2012/08/15 | 1 week ago | Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library. |
| [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go) | 3117 | 270 | 2016/07/12 | 3 weeks ago | confluent-kafka-go is Confluent's Golang client for Apache Kafka and the Confluent Platform. |
| [mercure](https://github.com/dunglas/mercure) | 2649 | 54 | 2018/07/14 | 3 weeks ago | Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events). |
| [asynq](https://github.com/hibiken/asynq) | 2624 | 51 | 2019/11/15 | 2 weeks ago | A simple, reliable, and efficient distributed task queue for Go built on top of Redis. |
| [apns2](https://github.com/sideshow/apns2) | 2587 | 76 | 2016/01/05 | 4 months ago | HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps. |
| [melody](https://github.com/olahol/melody) | 2343 | 58 | 2015/05/13 | 9 months ago | Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling. |
| [go-nsq](https://github.com/nsqio/go-nsq) | 2079 | 64 | 2013/08/29 | 2 months ago | the official Go package for NSQ. |
| [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) | 2022 | 231 | 2013/12/27 | 4 years ago | gopush-cluster is a go push server cluster. |
| [Beaver](https://github.com/Clivern/Beaver) | 1309 | 30 | 2018/10/20 | 1 week ago | A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps. |
| [uniqush-push](https://github.com/uniqush/uniqush-push) | 1300 | 79 | 2011/08/29 | 1 year ago | Redis backed unified push service for server-side notifications to mobile devices. |
| [EventBus](https://github.com/asaskevich/EventBus) | 1084 | 29 | 2014/12/19 | 7 months ago | The lightweight event bus with async compatibility. |
| [zmq4](https://github.com/pebbe/zmq4) | 957 | 44 | 2013/10/18 | 4 months ago | Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2). |
| [gollum](https://github.com/trivago/gollum) | 913 | 39 | 2015/06/20 | 7 months ago | A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations. |
| [chanify](https://github.com/chanify/chanify) | 744 | 6 | 2021/02/25 | 3 weeks ago | A push notification server send message to your iOS devices. |
| [dbus](https://github.com/godbus/dbus) | 671 | 23 | 2014/03/27 | 1 week ago | Native Go bindings for D-Bus. |
| [golongpoll](https://github.com/jcuga/golongpoll) | 590 | 22 | 2015/11/02 | 9 months ago | HTTP longpoll server library that makes web pub-sub simple. |
| [mangos](https://github.com/nanomsg/mangos) | 477 | 22 | 2018/10/12 | 2 weeks ago | Pure go implementation of the Nanomsg ("Scalability Protocols") with transport interoperability. |
| [emitter](https://github.com/olebedev/emitter) | 417 | 10 | 2015/11/10 | 2 years ago | Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins. |
| [glue](https://github.com/desertbit/glue) | 393 | 17 | 2015/06/07 | 1 year ago | Robust Go and Javascript Socket Library (Alternative to Socket.io). |
| [pubsub](https://github.com/cskr/pubsub) | 366 | 9 | 2012/04/01 | 3 weeks ago | Simple pubsub package for go. |
| [bus](https://github.com/mustafaturan/bus) | 244 | 3 | 2019/04/27 | 9 months ago | Minimalist message bus implementation for internal communication. |
| [amqp091-go](https://github.com/rabbitmq/amqp091-go) | 221 | 23 | 2021/06/09 | 2 weeks ago | Go RabbitMQ Client Library. |
| [rabtap](https://github.com/jandelgado/rabtap) | 206 | 8 | 2017/11/11 | 2 months ago | RabbitMQ swiss army knife cli app. |
| [message-bus](https://github.com/vardius/message-bus) | 204 | 8 | 2017/10/04 | 1 year ago | messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD. |
| [guble](https://github.com/smancke/guble) | 151 | 13 | 2015/11/15 | 4 years ago | Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence. |
| [hub](https://github.com/leandro-lugaresi/hub) | 115 | 2 | 2018/04/13 | 1 year ago | A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges. |
| [oplog](https://github.com/dailymotion/oplog) | 111 | 92 | 2014/11/06 | 6 years ago | Generic oplog/replication system for REST APIs. |
| [rabbus](https://github.com/rafaeljesus/rabbus) | 94 | 8 | 2017/05/07 | 2 years ago | A tiny wrapper over amqp exchanges and queues. |
| [drone-line](https://github.com/appleboy/drone-line) | 76 | 5 | 2016/09/13 | 8 months ago | Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI. |
| [go-mq](https://github.com/cheshir/go-mq) | 74 | 7 | 2017/06/19 | 2 months ago | RabbitMQ client with declarative configuration. |
| [redisqueue](https://github.com/robinjoseph08/redisqueue) | 72 | 3 | 2019/07/07 | 1 week ago | redisqueue provides a producer and consumer of a queue that uses Redis streams. |
| [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) | 71 | 8 | 2017/01/15 | 4 years ago | A tiny wrapper around NSQ topic and channel. |
| [RapidMQ](https://github.com/sybrexsys/RapidMQ) | 64 | 5 | 2016/10/04 | 4 years ago | RapidMQ is a lightweight and reliable library for managing of the local messages queue. |
| [commander](https://github.com/jeroenrinzema/commander) | 58 | 1 | 2018/04/20 | 9 months ago | A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka. |
| [go-notify](https://github.com/TheCreeper/go-notify) | 58 | 2 | 2015/03/01 | 1 year ago | Native implementation of the freedesktop notification spec. |
| [go-res](https://github.com/jirenius/go-res) | 54 | 2 | 2018/07/15 | 4 weeks ago | Package for building REST/real-time services where clients are synchronized seamlessly, using NATS and Resgate. |
| [event](https://github.com/agoalofalife/event) | 45 | 4 | 2017/07/02 | 4 years ago | Implementation of the pattern observer. |
| [hare](https://github.com/leozz37/hare) | 34 | 2 | 2020/12/01 | 1 month ago | A user friendly library for sending messages and listening to TCP sockets. |
| [ami](https://github.com/kak-tus/ami) | 22 | 1 | 2018/10/27 | 1 year ago | Go client to reliable queues based on Redis Cluster Streams. |
| [gosd](https://github.com/alexsniffin/gosd) | 19 | 1 | 2020/05/17 | 1 year ago | A library for scheduling when to dispatch a message to a channel. |
| [go-vitotrol](https://github.com/maxatome/go-vitotrol) | 17 | 7 | 2016/11/03 | 1 year ago | Client library to Viessmann Vitotrol web service. |
| [rmqconn](https://github.com/sbabiv/rmqconn) | 16 | 1 | 2019/01/14 | 2 years ago | RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed. |
| [jazz](https://github.com/socifi/jazz) | 14 | 4 | 2018/10/22 | 2 years ago | A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages. |
| [gaurun-client](https://github.com/osamingo/gaurun-client) | 10 | 1 | 2017/06/29 | 6 months ago | Gaurun Client written in Go. |
**[⬆ back to top](#awesome-go-info)**

# Microsoft Office
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [unioffice](https://github.com/unidoc/unioffice) | 3166 | 75 | 2017/08/29 | 1 month ago | Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents. |
**[⬆ back to top](#awesome-go-info)**

# Miscellaneous
        
**[⬆ back to top](#awesome-go-info)**

# Natural Language Processing
        
*Libraries for working with human languages.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [prose](https://github.com/jdkato/prose) | 2866 | 60 | 2017/02/17 | 4 months ago | Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. English only. |
| [gojieba](https://github.com/yanyiwu/gojieba) | 1772 | 68 | 2015/09/12 | 2 weeks ago | This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm. |
| [gse](https://github.com/go-ego/gse) | 1742 | 59 | 2017/06/23 | 1 month ago | Go efficient text segmentation; support english, chinese, japanese and other. |
| [when](https://github.com/olebedev/when) | 1139 | 25 | 2016/12/27 | 2 months ago | Natural EN and RU language date/time parser with pluggable rules. |
| [spago](https://github.com/nlpodyssey/spago) | 1089 | 31 | 2020/01/05 | 1 week ago | Self-contained Machine Learning and Natural Language Processing library in Go. |
| [go-pinyin](https://github.com/mozillazg/go-pinyin) | 1076 | 38 | 2014/11/09 | 2 months ago | CN Hanzi to Hanyu Pinyin converter. |
| [kagome](https://github.com/ikawaha/kagome) | 601 | 23 | 2014/06/26 | 2 months ago | JP morphological analyzer written in pure Go. |
| [whatlanggo](https://github.com/abadojack/whatlanggo) | 522 | 15 | 2017/02/20 | 1 year ago | Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc). |
| [nlp](https://github.com/shixzie/nlp) | 378 | 22 | 2017/01/25 | 4 years ago | Extract values from strings and fill your structs with nlp. |
| [nlp](https://github.com/nymiun/nlp) | 378 | 22 | 2017/01/25 | 4 years ago | Extract values from strings and fill your structs with nlp. |
| [nlp](https://github.com/james-bowman/nlp) | 354 | 24 | 2017/03/15 | 9 months ago | Go Natural Language Processing library supporting LSA (Latent Semantic Analysis). |
| [sentences](https://github.com/neurosnap/sentences) | 316 | 14 | 2015/08/07 | 8 months ago | Sentence tokenizer:  converts text into a list of sentences. |
| [getlang](https://github.com/rylans/getlang) | 127 | 4 | 2018/03/01 | 1 year ago | Fast natural language detection package. |
| [go-nlp](https://github.com/nuance/go-nlp) | 91 | 7 | 2011/05/02 | 10 years ago | Utilities for working with discrete probability distributions and other tools useful for doing NLP work. |
| [go-unidecode](https://github.com/mozillazg/go-unidecode) | 91 | 2 | 2016/07/08 | 9 months ago | ASCII transliterations of Unicode text. |
| [RAKE.Go](https://github.com/afjoseph/RAKE.Go) | 89 | 7 | 2016/12/17 | 1 year ago | Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE). |
| [gounidecode](https://github.com/fiam/gounidecode) | 74 | 5 | 2012/05/01 | 6 years ago | Unicode transliterator (also known as unidecode) for Go. |
| [textcat](https://github.com/pebbe/textcat) | 66 | 7 | 2012/09/21 | 1 year ago | Go package for n-gram based text categorization, with support for utf-8 and raw text. |
| [go-stem](https://github.com/agonopol/go-stem) | 65 | 4 | 2011/09/23 | 3 years ago | Implementation of the porter stemming algorithm. |
| [MMSEGO](https://github.com/awsong/MMSEGO) | 60 | 5 | 2012/04/18 | 9 years ago | This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm. |
| [stemmer](https://github.com/dchest/stemmer) | 51 | 4 | 2011/03/21 | 5 years ago | Stemmer packages for Go programming language. Includes English and German stemmers. |
| [go2vec](https://github.com/danieldk/go2vec) | 44 | 7 | 2015/01/27 | 3 years ago | Reader and utility functions for word2vec embeddings. |
| [address](https://github.com/bojanz/address) | 43 | 3 | 2020/10/07 | 1 week ago | Handles address representation, validation and formatting. |
| [porter2](https://github.com/zentures/porter2) | 42 | 3 | 2015/01/21 | 1 year ago | Really fast Porter 2 stemmer. |
| [petrovich](https://github.com/striker2000/petrovich) | 37 | 2 | 2016/12/26 | 11 months ago | Petrovich is the library which inflects Russian names to given grammatical case. |
| [go-localize](https://github.com/m1/go-localize) | 30 | 2 | 2019/12/23 | 3 months ago | Simple and easy to use i18n (Internationalization and localization) engine - used for translating locale strings. |
| [snowball](https://github.com/goodsign/snowball) | 30 | 1 | 2012/12/11 | 4 years ago | Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/). |
| [mystem](https://github.com/dveselov/mystem) | 28 | 3 | 2016/08/30 | 5 years ago | CGo bindings to Yandex.Mystem - russian morphology analyzer. |
| [iuliia-go](https://github.com/mehanizm/iuliia-go) | 28 | 2 | 2020/04/27 | 8 months ago | Transliterate Cyrillic → Latin in every possible way. |
| [paicehusk](https://github.com/rookii/paicehusk) | 28 | 3 | 2012/09/29 | 8 years ago | Golang implementation of the Paice/Husk Stemming Algorithm. |
| [golibstemmer](https://github.com/rjohnsondev/golibstemmer) | 20 | 2 | 2012/08/06 | 7 years ago | Go bindings for the snowball libstemmer library including porter 2. |
| [icu](https://github.com/goodsign/icu) | 20 | 0 | 2012/12/11 | 4 years ago | Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1. |
| [transliterator](https://github.com/alexsergivan/transliterator) | 20 | 1 | 2020/04/17 | 1 year ago | Provides one-way string transliteration with supporting of language-specific transliteration rules. |
| [govader](https://github.com/jonreiter/govader) | 17 | 1 | 2020/01/19 | 1 week ago | Go implementation of [VADER Sentiment Analysis](https://github.com/cjhutto/vaderSentiment). |
| [gotokenizer](https://github.com/xujiajun/gotokenizer) | 14 | 1 | 2018/10/11 | 2 years ago | A tokenizer based on the dictionary and Bigram language models for Go. (Now only support chinese segmentation) |
| [detectlanguage-go](https://github.com/detectlanguage/detectlanguage-go) | 13 | 2 | 2019/12/14 | 1 year ago | Language Detection API Go Client. Supports batch requests, short phrase or single word language detection. |
| [shamoji](https://github.com/osamingo/shamoji) | 12 | 2 | 2017/07/23 | 1 year ago | The shamoji is word filtering package written in Go. |
| [libtextcat](https://github.com/goodsign/libtextcat) | 11 | 2 | 2012/12/10 | 9 years ago | Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2. |
| [porter](https://github.com/a2800276/porter) | 9 | 1 | 2013/09/17 | 8 years ago | This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm. |
| [gosentiwordnet](https://github.com/dinopuguh/gosentiwordnet) | 8 | 1 | 2020/04/21 | 11 months ago | Sentiment analyzer using sentiwordnet lexicon in Go. |
| [t](https://github.com/youthlin/t) | 8 | 2 | 2021/06/04 | 3 months ago | Another i18n pkg for golang, which follows GNU gettext style and supports .po/.mo files: `t.T (gettext)`, `t.N (ngettext)`, etc. And it contains a cmd tool [xtemplate](https://github.com/youthlin/t/blob/main/cmd/xtemplate), which can extract messages as a pot file from text/html template. |
**[⬆ back to top](#awesome-go-info)**

# Networking
        
*Libraries for working with various layers of the network.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [fasthttp](https://github.com/valyala/fasthttp) | 16856 | 396 | 2015/10/18 | 1 week ago | Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http. |
| [kcptun](https://github.com/xtaci/kcptun) | 12734 | 597 | 2016/02/26 | 1 month ago | Extremely simple & fast udp tunnel based on KCP protocol. |
| [webrtc](https://github.com/pion/webrtc) | 8693 | 250 | 2018/05/18 | 1 week ago | A pure Go implementation of the WebRTC API. |
| [quic-go](https://github.com/lucas-clemente/quic-go) | 6262 | 203 | 2016/04/06 | 2 weeks ago | An implementation of the QUIC protocol in pure Go. |
| [dns](https://github.com/miekg/dns) | 6077 | 179 | 2010/08/03 | 1 week ago | Go library for working with DNS. |
| [gnet](https://github.com/panjf2000/gnet) | 6047 | 157 | 2019/02/24 | 2 weeks ago | `gnet` is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go. |
| [gopacket](https://github.com/google/gopacket) | 4630 | 141 | 2015/03/16 | 1 week ago | Go library for packet processing with libpcap bindings. |
| [httplab](https://github.com/qustavo/httplab) | 3768 | 63 | 2017/02/08 | 2 years ago | HTTPLabs let you inspect HTTP requests and forge responses. |
| [kcp-go](https://github.com/xtaci/kcp-go) | 3218 | 151 | 2015/06/16 | 3 weeks ago | KCP - Fast and Reliable ARQ Protocol. |
| [gobgp](https://github.com/osrg/gobgp) | 2420 | 118 | 2014/09/14 | 3 weeks ago | BGP implemented in the Go Programming Language. |
| [ssh](https://github.com/gliderlabs/ssh) | 2346 | 56 | 2016/10/03 | 1 month ago | Higher-level API for building SSH servers (wraps crypto/ssh). |
| [netpoll](https://github.com/cloudwego/netpoll) | 2308 | 56 | 2021/02/25 | 1 week ago | A high-performance non-blocking I/O networking framework, which focused on RPC scenarios, developed by ByteDance. |
| [fortio](https://github.com/fortio/fortio) | 2261 | 46 | 2017/10/10 | 1 month ago | Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC. |
| [water](https://github.com/songgao/water) | 1368 | 46 | 2013/03/25 | 2 weeks ago | Simple TUN/TAP library. |
| [gev](https://github.com/Allenxuxu/gev) | 1344 | 39 | 2019/09/01 | 3 weeks ago | gev is a lightweight, fast non-blocking TCP network library based on Reactor mode. |
| [go-getter](https://github.com/hashicorp/go-getter) | 1283 | 258 | 2015/10/12 | 2 weeks ago | Go library for downloading files or directories from various sources using a URL. |
| [nff-go](https://github.com/intel-go/nff-go) | 1157 | 80 | 2017/03/29 | 5 months ago | Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF). |
| [sftp](https://github.com/pkg/sftp) | 1093 | 53 | 2013/11/05 | 1 week ago | Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt. |
| [grab](https://github.com/cavaliergopher/grab) | 981 | 17 | 2016/01/05 | 1 month ago | Go package for managing file downloads. |
| [ftp](https://github.com/jlaffaye/ftp) | 881 | 26 | 2011/05/06 | 1 week ago | Package ftp implements a FTP client as described in [RFC 959](https://tools.ietf.org/html/rfc959). |
| [mdns](https://github.com/hashicorp/mdns) | 845 | 263 | 2014/01/29 | 1 month ago | Simple mDNS (Multicast DNS) client/server library in Golang. |
| [gosnmp](https://github.com/gosnmp/gosnmp) | 800 | 47 | 2012/08/27 | 1 week ago | Native Go library for performing SNMP actions. |
| [vssh](https://github.com/yahoo/vssh) | 770 | 20 | 2020/06/09 | 1 year ago | Go library for building network and server automation over SSH protocol. |
| [cidranger](https://github.com/yl2chen/cidranger) | 668 | 18 | 2017/08/21 | 3 weeks ago | Fast IP to CIDR lookup for Go. |
| [lhttp](https://github.com/fanux/lhttp) | 660 | 59 | 2015/12/29 | 3 years ago | Powerful websocket framework, build your IM server more easily. |
| [gmqtt](https://github.com/DrmagicE/gmqtt) | 565 | 30 | 2018/09/16 | 1 month ago | Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1. |
| [peerdiscovery](https://github.com/schollz/peerdiscovery) | 526 | 19 | 2018/04/22 | 1 week ago | Pure Go library for cross-platform local peer discovery using UDP multicast. |
| [gotcp](https://github.com/gansidui/gotcp) | 487 | 39 | 2014/04/13 | 4 years ago | Go package for quickly writing tcp applications. |
| [stun](https://github.com/gortc/stun) | 479 | 21 | 2016/04/24 | 9 months ago | Go implementation of RFC 5389 STUN protocol. |
| [nbio](https://github.com/lesismal/nbio) | 473 | 13 | 2020/01/25 | 1 month ago | Pure Go 1000k+ connections solution, support tls/http1.x/websocket and basically compatible with net/http, with high-performance and low memory cost, non-blocking, event-driven, easy-to-use. |
| [go-stun](https://github.com/ccding/go-stun) | 471 | 15 | 2013/08/17 | 2 weeks ago | Go implementation of the STUN client (RFC 3489 and RFC 5389). |
| [gopcap](https://github.com/akrennmair/gopcap) | 439 | 23 | 2009/11/19 | 9 months ago | Go wrapper for libpcap. |
| [raw](https://github.com/mdlayher/raw) | 417 | 13 | 2015/07/06 | 2 months ago | Package raw enables reading and writing data at the device driver level for a network interface. |
| [gaio](https://github.com/xtaci/gaio) | 414 | 17 | 2019/12/20 | 5 months ago | High performance async-io networking for Golang in proactor mode. |
| [tcp_server](https://github.com/firstrow/tcp_server) | 409 | 19 | 2014/10/13 | 3 months ago | Go library for building tcp servers faster. |
| [winrm](https://github.com/masterzen/winrm) | 350 | 18 | 2013/12/30 | 1 month ago | Go WinRM client to remotely execute commands on Windows machines. |
| [ftpserverlib](https://github.com/fclairamb/ftpserverlib) | 283 | 11 | 2016/09/25 | 1 week ago | Fully featured FTP server library. |
| [arp](https://github.com/mdlayher/arp) | 275 | 10 | 2015/07/06 | 2 years ago | Package arp implements the ARP protocol, as described in RFC 826. |
| [easytcp](https://github.com/DarthPestilane/easytcp) | 252 | 1 | 2021/04/26 | 2 weeks ago | A light-weight TCP framework written in Go (Golang), built with message router. EasyTCP helps you build a TCP server easily fast and less painful. |
| [buffstreams](https://github.com/StabbyCutyou/buffstreams) | 247 | 14 | 2015/06/29 | 1 year ago | Streaming protocolbuffer data over TCP made easy. |
| [ethernet](https://github.com/mdlayher/ethernet) | 238 | 12 | 2015/07/03 | 2 years ago | Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. |
| [gnxi](https://github.com/google/gnxi) | 202 | 31 | 2017/09/26 | 2 weeks ago | A collection of tools for Network Management that use the gNMI and gNOI protocols. |
| [jazigo](https://github.com/udhos/jazigo) | 175 | 14 | 2016/06/07 | 2 years ago | Jazigo is a tool written in Go for retrieving configuration for multiple network devices. |
| [utp](https://github.com/anacrolix/utp) | 163 | 19 | 2015/03/20 | 1 year ago | Go uTP micro transport protocol implementation. |
| [canopus](https://github.com/zubairhamed/canopus) | 148 | 14 | 2015/02/24 | 3 years ago | CoAP Client/Server implementation (RFC 7252). |
| [dnsmonster](https://github.com/mosajjal/dnsmonster) | 137 | 8 | 2020/02/09 | 1 week ago | Passive DNS Capture/Monitoring Framework. |
| [sslb](https://github.com/eduardonunesp/sslb) | 135 | 9 | 2015/10/18 | 2 years ago | It's a Super Simples Load Balancer, just a little project to achieve some kind of performance. |
| [xtcp](https://github.com/xfxdev/xtcp) | 129 | 14 | 2016/03/31 | 1 year ago | TCP Server Framework with simultaneous full duplex communication, graceful shutdown, and custom protocol. |
| [ether](https://github.com/songgao/ether) | 74 | 4 | 2014/05/21 | 5 years ago | Cross-platform Go package for sending and receiving ethernet frames. |
| [dhcp6](https://github.com/mdlayher/dhcp6) | 73 | 4 | 2015/05/22 | 2 years ago | Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. |
| [packet](https://github.com/aerogo/packet) | 62 | 5 | 2017/10/29 | 2 years ago | Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed. |
| [go-powerdns](https://github.com/joeig/go-powerdns) | 52 | 4 | 2018/06/21 | 1 month ago | PowerDNS API bindings for Golang. |
| [linkio](https://github.com/ian-kent/linkio) | 50 | 2 | 2014/12/24 | 4 years ago | Network link speed simulation for Reader/Writer interfaces. |
| [portproxy](https://github.com/aybabtme/portproxy) | 46 | 0 | 2014/12/13 | 7 years ago | Simple TCP proxy which adds CORS support to API's which don't support it. |
| [panoptes-stream](https://github.com/yahoo/panoptes-stream) | 32 | 8 | 2020/10/09 | 11 months ago | A cloud native distributed streaming network telemetry (gNMI, Juniper JTI and Cisco MDT). |
| [graval](https://github.com/koofr/graval) | 28 | 8 | 2014/04/22 | 1 year ago | Experimental FTP server framework. |
| [publicip](https://github.com/polera/publicip) | 25 | 4 | 2016/12/28 | 5 years ago | Package publicip returns your public facing IPv4 address (internet egress). |
| [golibwireshark](https://github.com/sunwxg/golibwireshark) | 23 | 2 | 2015/11/16 | 4 years ago | Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data. |
| [gohooks](https://github.com/averageflow/gohooks) | 16 | 1 | 2020/10/30 | 7 months ago | GoHooks make it easy to send and consume secured web-hooks from a Go application. Inspired by Spatie's Laravel Webhook Client and Server. |
| [goshark](https://github.com/sunwxg/goshark) | 14 | 1 | 2015/11/01 | 4 years ago | Package goshark use tshark to decode IP packet and create data struct to analyse packet. |
| [llb](https://github.com/kirillDanshin/llb) | 12 | 1 | 2016/02/21 | 5 years ago | It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response. |
| [tspool](https://github.com/two/tspool) | 12 | 1 | 2018/10/27 | 3 years ago | A TCP Library use worker pool to improve performance and protect your server. |
| [httpproxy](https://github.com/wzshiming/httpproxy) | 11 | 2 | 2018/07/18 | 3 months ago | HTTP proxy handler and dialer. |
**[⬆ back to top](#awesome-go-info)**

# OpenGL
        
*Libraries for using OpenGL in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [glfw](https://github.com/go-gl/glfw) | 1211 | 37 | 2013/05/19 | 3 weeks ago | Go bindings for GLFW 3. |
| [gl](https://github.com/go-gl/gl) | 871 | 40 | 2015/02/22 | 2 months ago | Go bindings for OpenGL (generated via glow). |
| [mathgl](https://github.com/go-gl/mathgl) | 409 | 28 | 2013/02/13 | 1 year ago | Pure Go math package specialized for 3D math, with inspiration from GLM. |
| [gl](https://github.com/goxjs/gl) | 155 | 15 | 2015/05/18 | 1 year ago | Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android). |
| [glfw](https://github.com/goxjs/glfw) | 72 | 7 | 2014/12/27 | 3 weeks ago | Go cross-platform glfw library for creating an OpenGL context and receiving events. |
| [go-glmatrix](https://github.com/technohippy/go-glmatrix) | 3 | 1 | 2020/07/02 | 1 year ago | Go port of [glMatrix](https://glmatrix.net/) library. |
**[⬆ back to top](#awesome-go-info)**

# ORM
        
*Libraries that implement Object-Relational Mapping or datamapping techniques.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [beego](https://github.com/beego/beego) | 27662 | 1230 | 2012/02/29 | 2 weeks ago | Powerful orm framework for go. Support: pq/mysql/sqlite3. |
| [gorm](https://github.com/go-gorm/gorm) | 26721 | 496 | 2013/10/25 | 1 week ago | The fantastic ORM library for Golang, aims to be developer friendly. |
| [ent](https://github.com/ent/ent) | 9812 | 146 | 2019/06/12 | 1 week ago | An entity framework for Go. Simple, yet powerful ORM for modeling and querying data. |
| [pg](https://github.com/go-pg/pg) | 4970 | 87 | 2013/04/24 | 1 month ago | PostgreSQL ORM with focus on PostgreSQL specific features and performance. |
| [sqlboiler](https://github.com/volatiletech/sqlboiler) | 4617 | 80 | 2016/02/21 | 2 weeks ago | ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema. |
| [gorp](https://github.com/go-gorp/gorp) | 3548 | 111 | 2012/01/04 | 11 months ago | Go Relational Persistence, ORM-ish library for Go. |
| [db](https://github.com/upper/db) | 2882 | 63 | 2013/10/23 | 1 week ago | Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers. |
| [gormt](https://github.com/xxjwxc/gormt) | 1680 | 21 | 2019/05/05 | 4 weeks ago | Mysql database to golang gorm struct. |
| [reform](https://github.com/go-reform/reform) | 1228 | 24 | 2016/02/25 | 2 weeks ago | Better ORM for Go, based on non-empty interfaces and code generation. |
| [pop](https://github.com/gobuffalo/pop) | 1167 | 24 | 2018/02/07 | 1 week ago | Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite. |
| [prisma-client-go](https://github.com/prisma/prisma-client-go) | 1133 | 23 | 2019/09/24 | 1 week ago | Prisma Client Go, Typesafe database access for Go. |
| [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) | 698 | 13 | 2017/12/27 | 2 months ago | A flexible and powerful SQL string builder library plus a zero-config ORM. |
| [go-queryset](https://github.com/jirfag/go-queryset) | 649 | 24 | 2017/09/03 | 7 months ago | 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM. |
| [rel](https://github.com/go-rel/rel) | 479 | 10 | 2019/10/06 | 3 months ago | Modern Database Access Layer for Golang - Testable, Extendable and Crafted Into a Clean and Elegant API. |
| [zoom](https://github.com/albrow/zoom) | 286 | 20 | 2013/07/17 | 1 year ago | Blazing-fast datastore and querying engine built on Redis. |
| [gosql](https://github.com/rushteam/gosql) | 156 | 7 | 2020/04/27 | 7 months ago | A easy ORM for mysql. |
| [grimoire](https://github.com/Fs02/grimoire) | 156 | 6 | 2018/03/05 | 3 months ago | Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3). |
| [go-store](https://github.com/gosuri/go-store) | 108 | 9 | 2015/03/22 | 5 years ago | Simple and fast Redis backed key-value store library for Go. |
| [go-firestorm](https://github.com/jschoedt/go-firestorm) | 28 | 1 | 2018/12/04 | 2 months ago | A simple ORM for Google/Firebase Cloud Firestore. |
| [cacheme-go](https://github.com/Yiling-J/cacheme-go) | 18 | 1 | 2021/10/03 | 1 month ago | Schema based, typed Redis caching/memoize framework for Go. |
| [lore](https://github.com/abrahambotros/lore) | 11 | 1 | 2017/04/29 | 4 years ago | Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go. |
| [marlow](https://github.com/marlow/marlow) | 10 | 3 | 2020/08/11 | 1 year ago | Generated ORM from project structs for compile time safety assurances. |
**[⬆ back to top](#awesome-go-info)**

# Package Management
        
*Unofficial libraries for package and dependency management.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [dep](https://github.com/golang/dep) | 13090 | 260 | 2016/10/07 | 1 year ago | Go dependency tool. |
| [glide](https://github.com/Masterminds/glide) | 8148 | 189 | 2014/07/09 | 1 year ago | Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip. |
| [godep](https://github.com/tools/godep) | 5603 | 149 | 2013/05/01 | 3 years ago | dependency tool for go, godep helps build packages reproducibly by fixing their dependencies. |
| [govendor](https://github.com/kardianos/govendor) | 4985 | 98 | 2015/04/12 | 1 year ago | Go Package Manager. Go vendor tool that works with the standard vendor file. |
| [gopm](https://github.com/gpmgo/gopm) | 2493 | 85 | 2013/05/15 | 2 years ago | Go Package Manager. |
| [gom](https://github.com/mattn/gom) | 1398 | 35 | 2013/09/11 | 2 years ago | Go Manager - bundle for go. |
| [gpm](https://github.com/pote/gpm) | 1201 | 32 | 2013/09/05 | 4 years ago | Barebones dependency manager for Go. |
| [goop](https://github.com/petejkim/goop) | 778 | 36 | 2014/06/18 | 6 years ago | Simple dependency manager for Go (golang), inspired by Bundler. |
| [modgv](https://github.com/lucasepe/modgv) | 400 | 7 | 2020/09/12 | 1 year ago | Converts 'go mod graph' output into Graphviz's DOT language. |
| [nut](https://github.com/owenthereal/nut) | 238 | 6 | 2015/01/23 | 6 years ago | Vendor Go dependencies. |
| [johnny-deps](https://github.com/VividCortex/johnny-deps) | 212 | 21 | 2013/07/19 | 1 year ago | Minimal dependency version using Git. |
| [mvn-golang](https://github.com/raydac/mvn-golang) | 131 | 11 | 2016/03/24 | 4 weeks ago | plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure. |
| [VenGO](https://github.com/DamnWidget/VenGO) | 120 | 9 | 2014/10/17 | 5 years ago | create and manage exportable isolated go virtual environments. |
| [gop](https://github.com/lunny/gop) | 48 | 8 | 2017/02/18 | 2 years ago | Build and manage your Go applications out of GOPATH. |
**[⬆ back to top](#awesome-go-info)**

# Performance
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [jaeger](https://github.com/jaegertracing/jaeger) | 15114 | 344 | 2016/04/15 | 1 week ago | A distributed tracing system. |
| [pixie](https://github.com/pixie-io/pixie) | 2863 | 63 | 2020/02/27 | 1 week ago | No instrumentation tracing for Golang applications via eBPF. |
| [statsviz](https://github.com/arl/statsviz) | 1760 | 21 | 2020/08/14 | 1 month ago | Live visualization of your Go application runtime statistics. |
| [profile](https://github.com/pkg/profile) | 1648 | 40 | 2014/10/22 | 3 months ago | Simple profiling support package for Go. |
| [tracer](https://github.com/kamilsk/tracer) | 63 | 4 | 2019/06/22 | 11 months ago | Simple, lightweight tracing. |
**[⬆ back to top](#awesome-go-info)**

# Query Language
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [graphql](https://github.com/graphql-go/graphql) | 8289 | 150 | 2015/07/19 | 1 month ago | Implementation of GraphQL for Go. |
| [gqlgen](https://github.com/99designs/gqlgen) | 7073 | 151 | 2018/02/11 | 1 week ago | go generate based graphql server library. |
| [graphql-go](https://github.com/graph-gophers/graphql-go) | 4040 | 86 | 2016/10/18 | 2 weeks ago | GraphQL server with a focus on ease of use. |
| [gojsonq](https://github.com/thedevsaddam/gojsonq) | 1811 | 32 | 2018/05/19 | 2 weeks ago | A simple Go package to Query over JSON Data. |
| [dasel](https://github.com/TomWright/dasel) | 1797 | 10 | 2020/09/22 | 1 month ago | Query and update data structures using selectors from the command line. Comparable to jq/yq but supports JSON, YAML, TOML and XML with zero runtime dependencies. |
| [jsonql](https://github.com/elgs/jsonql) | 249 | 11 | 2015/12/29 | 1 year ago | JSON query expression library in Golang. |
| [rql](https://github.com/a8m/rql) | 246 | 11 | 2018/06/05 | 1 month ago | Resource Query Language for REST API. |
| [jsonslice](https://github.com/bhmj/jsonslice) | 59 | 0 | 2018/05/02 | 1 month ago | Jsonpath queries with advanced filters. |
| [graphql](https://github.com/tmc/graphql) | 54 | 11 | 2015/04/18 | 4 years ago | graphql parser + utilities. |
| [api-fu](https://github.com/ccbrown/api-fu) | 38 | 3 | 2019/07/30 | 1 month ago | Comprehensive GraphQL implementation. |
| [straf](https://github.com/ThundR67/straf) | 33 | 1 | 2019/08/16 | 1 year ago | Easily Convert Golang structs to GraphQL objects. |
| [straf](https://github.com/SonicRoshan/straf) | 32 | 1 | 2019/08/16 | 1 year ago | Easily Convert Golang structs to GraphQL objects. |
| [rest-query-parser](https://github.com/timsolov/rest-query-parser) | 28 | 2 | 2020/02/10 | 2 months ago | Query Parser for REST API. Filtering, validations, both `AND`, `OR` operations are supported directly in the query. |
| [jsonpath](https://github.com/AsaiYusuke/jsonpath) | 8 | 2 | 2020/11/29 | 3 months ago | A query library for retrieving part of JSON based on JSONPath syntax. |
| [gws](https://github.com/Zaba505/gws) | 4 | 2 | 2020/06/08 | 1 year ago | Apollos' "GraphQL over Websocket" client and server implementation. |
**[⬆ back to top](#awesome-go-info)**

# Resource Embedding
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [statik](https://github.com/rakyll/statik) | 3363 | 52 | 2014/02/04 | 1 year ago | Embeds static files into a Go executable. |
| [packr](https://github.com/gobuffalo/packr) | 3354 | 47 | 2017/03/15 | 2 months ago | The simple and easy way to embed static files into Go binaries. |
| [go.rice](https://github.com/GeertJohan/go.rice) | 2270 | 55 | 2013/10/23 | 3 months ago | go.rice is a Go package that makes working with resources such as HTML, JS, CSS, images, and templates very easy. |
| [vfsgen](https://github.com/shurcooL/vfsgen) | 940 | 19 | 2015/05/18 | 1 month ago | Generates a vfsdata.go file that statically implements the given virtual filesystem. |
| [esc](https://github.com/mjibson/esc) | 607 | 15 | 2014/01/26 | 2 years ago | Embeds files into Go programs and provides http.FileSystem interfaces to them. |
| [fileb0x](https://github.com/UnnoTed/fileb0x) | 599 | 19 | 2016/01/23 | 1 year ago | Simple tool to embed files in go with focus on "customization" and ease to use. |
| [go-resources](https://github.com/omeid/go-resources) | 174 | 9 | 2015/02/21 | 8 months ago | Unfancy resources embedding with Go. |
| [statics](https://github.com/go-playground/statics) | 65 | 3 | 2015/10/07 | 5 years ago | Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks. |
| [templify](https://github.com/wlbr/templify) | 27 | 3 | 2016/05/22 | 6 months ago | Embed external template files into Go code to create single file binaries. |
| [rebed](https://github.com/soypat/rebed) | 20 | 2 | 2021/02/17 | 10 months ago | Recreate folder structures and files from Go 1.16's `embed. |
| [debme](https://github.com/leaanthony/debme) | 15 | 1 | 2021/04/16 | 8 months ago | Create an `embed.FS` from an existing `embed.FS` subdirectory. |
| [mule](https://github.com/wlbr/mule) | 11 | 2 | 2020/01/17 | 6 months ago | Embed external resources like images, movies ... into Go source code to create single file binaries using `go generate`. Focussed on simplicity. |
**[⬆ back to top](#awesome-go-info)**

# Science and Data Analysis
        
*Libraries for scientific computing and data analyzing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gonum](https://github.com/gonum/gonum) | 5500 | 116 | 2017/03/25 | 1 week ago | Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more. |
| [stats](https://github.com/montanaflynn/stats) | 2272 | 52 | 2014/12/16 | 3 months ago | Statistics package with common functions missing from the Golang standard library. |
| [plot](https://github.com/gonum/plot) | 2100 | 57 | 2013/07/23 | 3 weeks ago | gonum/plot provides an API for building and drawing plots in Go. |
| [gosl](https://github.com/cpmech/gosl) | 1640 | 72 | 2015/02/09 | 2 weeks ago | Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more. |
| [streamtools](https://github.com/nytlabs/streamtools) | 1317 | 70 | 2013/07/05 | 6 years ago | general purpose, graphical tool for dealing with streams of data. |
| [go-dsp](https://github.com/mjibson/go-dsp) | 764 | 29 | 2011/11/02 | 2 weeks ago | Digital Signal Processing for Go. |
| [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) | 737 | 33 | 2018/10/01 | 3 months ago | Dataframes for machine-learning and statistics (similar to pandas). |
| [chart](https://github.com/vdobler/chart) | 712 | 44 | 2011/06/27 | 8 months ago | Simple Chart Plotting library for Go. Supports many graphs types. |
| [goraph](https://github.com/gyuho/goraph) | 653 | 39 | 2014/02/27 | 4 years ago | Pure Go graph theory library(data structure, algorithm visualization). |
| [graph](https://github.com/yourbasic/graph) | 519 | 22 | 2017/04/27 | 4 months ago | Library of basic graph algorithms. |
| [orb](https://github.com/paulmach/orb) | 475 | 23 | 2016/03/28 | 1 month ago | 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support. |
| [ewma](https://github.com/VividCortex/ewma) | 345 | 24 | 2013/07/05 | 6 months ago | Exponentially-weighted moving averages. |
| [calendarheatmap](https://github.com/nikolaydubina/calendarheatmap) | 336 | 9 | 2020/07/01 | 1 month ago | Calendar heatmap in plain Go inspired by Github contribution activity. |
| [gohistogram](https://github.com/VividCortex/gohistogram) | 159 | 17 | 2013/07/02 | 1 year ago | Approximate histograms for data streams. |
| [TextRank](https://github.com/DavidBelicza/TextRank) | 151 | 8 | 2018/01/09 | 7 months ago | TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support. |
| [sparse](https://github.com/james-bowman/sparse) | 125 | 8 | 2017/05/16 | 6 months ago | Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries. |
| [go-estimate](https://github.com/milosgajdos/go-estimate) | 90 | 5 | 2018/11/04 | 5 months ago | State estimation and filtering algorithms in Go. |
| [pagerank](https://github.com/alixaxel/pagerank) | 73 | 8 | 2015/08/06 | 8 months ago | Weighted PageRank algorithm implemented in Go. |
| [jsonl-graph](https://github.com/nikolaydubina/jsonl-graph) | 55 | 2 | 2021/06/26 | 1 month ago | Tool to manipulate JSONL graphs with graphviz support. |
| [geom](https://github.com/skelterjohn/geom) | 50 | 5 | 2011/06/07 | 4 years ago | 2D geometry for golang. |
| [evaler](https://github.com/soniah/evaler) | 47 | 5 | 2012/09/04 | 3 years ago | Simple floating point arithmetic expression evaluator. |
| [goent](https://github.com/kzahedi/goent) | 26 | 1 | 2017/08/08 | 2 years ago | GO Implementation of Entropy Measures. |
| [triangolatte](https://github.com/tchayen/triangolatte) | 25 | 2 | 2018/07/18 | 6 months ago | 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs. |
| [decimal](https://github.com/db47h/decimal) | 24 | 4 | 2020/05/27 | 1 year ago | Package decimal implements arbitrary-precision decimal floating-point arithmetic. |
| [piecewiselinear](https://github.com/sgreben/piecewiselinear) | 21 | 4 | 2018/10/21 | 1 year ago | Tiny linear interpolation library. |
| [GoStats](https://github.com/OGFris/GoStats) | 19 | 3 | 2018/07/22 | 3 years ago | GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions. |
| [godesim](https://github.com/soypat/godesim) | 17 | 2 | 2020/12/16 | 2 months ago | Extended/multivariable ODE solver framework for event-based simulations with simple API. |
| [PiHex](https://github.com/claygod/PiHex) | 15 | 4 | 2016/07/22 | 1 year ago | Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi. |
| [ode](https://github.com/ChristopherRabotin/ode) | 14 | 5 | 2016/11/11 | 4 years ago | Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions. |
| [assocentity](https://github.com/ndabAP/assocentity) | 7 | 2 | 2018/12/21 | 1 year ago | Package assocentity returns the average distance from words to a given entity. |
| [rootfinding](https://github.com/khezen/rootfinding) | 7 | 5 | 2018/10/30 | 1 year ago | root-finding algorithms library for finding roots of quadratic functions. |
| [go-gt](https://github.com/ThePaw/go-gt) | 6 | 0 | 2015/09/14 | 6 years ago | Graph theory algorithms written in "Go" language. |
| [bradleyterry](https://github.com/seanhagen/bradleyterry) | 5 | 2 | 2019/04/30 | 2 years ago | Provides a Bradley-Terry Model for pairwise comparisons. |
**[⬆ back to top](#awesome-go-info)**

# Security
        
*Libraries that are used to help make your application more secure.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [age](https://github.com/FiloSottile/age) | 9881 | 118 | 2019/05/18 | 3 weeks ago | A simple, modern and secure encryption tool (and Go library) with small explicit keys, no config options, and UNIX-style composability. |
| [lego](https://github.com/go-acme/lego) | 5078 | 93 | 2015/06/08 | 1 week ago | Pure Go ACME client library and CLI tool (for use with Let's Encrypt). |
| [certmagic](https://github.com/caddyserver/certmagic) | 3905 | 52 | 2018/12/10 | 1 week ago | Mature, robust, and powerful ACME client integration for fully-managed TLS certificate issuance and renewal. |
| [cameradar](https://github.com/Ullaakut/cameradar) | 2817 | 130 | 2016/05/20 | 3 months ago | Tool and library to remotely hack RTSP streams from surveillance cameras. |
| [memguard](https://github.com/awnumar/memguard) | 2070 | 50 | 2017/04/22 | 11 months ago | A pure Go library for handling sensitive values in memory. |
| [acmetool](https://github.com/hlandau/acmetool) | 1891 | 63 | 2015/11/15 | 10 months ago | ACME (Let's Encrypt) client tool with automatic renewal. |
| [secure](https://github.com/unrolled/secure) | 1876 | 36 | 2014/05/20 | 3 weeks ago | HTTP middleware for Go that facilitates some quick security wins. |
| [themis](https://github.com/cossacklabs/themis) | 1430 | 42 | 2015/05/06 | 3 weeks ago | high-level cryptographic library for solving typical data security tasks (secure data storage, secure messaging, zero-knowledge proof authentication), available for 14 languages, best fit for multi-platform apps. |
| [acra](https://github.com/cossacklabs/acra) | 930 | 42 | 2016/11/14 | 1 week ago | Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system. |
| [nacl](https://github.com/kevinburke/nacl) | 519 | 13 | 2017/07/20 | 10 months ago | Go implementation of the NaCL set of API's. |
| [ssh-vault](https://github.com/ssh-vault/ssh-vault) | 335 | 10 | 2016/09/29 | 7 months ago | encrypt/decrypt using ssh keys. |
| [firewalld-rest](https://github.com/prashantgupta24/firewalld-rest) | 313 | 9 | 2020/06/12 | 1 year ago | A rest application to dynamically update firewalld rules on a linux server. |
| [go-password-validator](https://github.com/wagslane/go-password-validator) | 312 | 9 | 2020/10/14 | 2 months ago | Password validator based on raw cryptographic entropy values. |
| [badactor](https://github.com/jaredfolkins/badactor) | 305 | 8 | 2014/12/12 | 1 year ago | In-memory, application-driven jailer built in the spirit of fail2ban. |
| [optimus-go](https://github.com/pjebs/optimus-go) | 298 | 5 | 2015/05/05 | 1 year ago | ID hashing and Obfuscation using Knuth's Algorithm. |
| [passlib](https://github.com/hlandau/passlib) | 261 | 12 | 2014/12/21 | 10 months ago | Futureproof password hashing library. |
| [go-yara](https://github.com/hillu/go-yara) | 243 | 21 | 2015/01/25 | 1 month ago | Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)". |
| [simple-scrypt](https://github.com/elithrar/simple-scrypt) | 174 | 7 | 2015/04/14 | 10 months ago | Scrypt package with a simple, obvious API and automatic cost calibration built-in. |
| [argon2pw](https://github.com/raja/argon2pw) | 89 | 4 | 2018/03/13 | 5 months ago | Argon2 password hash generation with constant-time password comparison. |
| [dongle](https://github.com/golang-module/dongle) | 71 | 2 | 2021/08/11 | 5 months ago | A simple, semantic and developer-friendly golang package for encoding&decoding and encryption&decryption. |
| [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) | 45 | 2 | 2017/10/19 | 1 year ago | A probably paranoid package for securely hashing and encrypting passwords. |
| [go-generate-password](https://github.com/m1/go-generate-password) | 31 | 2 | 2019/11/14 | 6 months ago | Password generator that can be used on the cli or as a library. |
| [certificates](https://github.com/mvmaasakkers/certificates) | 23 | 0 | 2019/03/04 | 1 year ago | An opinionated tool for generating tls certificates. |
| [secureio](https://github.com/xaionaro-go/secureio) | 23 | 3 | 2018/12/25 | 1 year ago | An keyexchanging+authenticating+encrypting wrapper and multiplexer for `io.ReadWriteCloser` based on XChaCha20-poly1305, ECDH and ED25519. |
| [go-htpasswd](https://github.com/tg123/go-htpasswd) | 21 | 1 | 2015/06/18 | 3 months ago | Apache htpasswd Parser for Go. |
| [argon2-hashing](https://github.com/andskur/argon2-hashing) | 16 | 2 | 2019/01/02 | 1 year ago | light wrapper around Go's argon2 package that closely mirrors with Go's standard library Bcrypt and simple-scrypt package. |
| [goArgonPass](https://github.com/dwin/goArgonPass) | 14 | 2 | 2018/05/30 | 1 year ago | Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations. |
| [sslmgr](https://github.com/adrianosela/sslmgr) | 12 | 2 | 2019/04/02 | 2 years ago | SSL certificates made easy with a high level wrapper around acme/autocert. |
| [secret](https://github.com/rsjethani/secret) | 6 | 1 | 2022/01/10 | 3 weeks ago | Prevent your secrets from leaking into logs, std* etc. |
**[⬆ back to top](#awesome-go-info)**

# Serialization
        
*Libraries and tools for binary serialization.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go](https://github.com/json-iterator/go) | 10421 | 238 | 2016/11/30 | 1 month ago | High-performance 100% compatible drop-in replacement of "encoding/json". |
| [protobuf](https://github.com/golang/protobuf) | 8220 | 214 | 2014/11/23 | 1 month ago | Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers. |
| [mapstructure](https://github.com/mitchellh/mapstructure) | 5336 | 72 | 2013/05/20 | 3 weeks ago | Go library for decoding generic map values into native Go structures. |
| [protobuf](https://github.com/gogo/protobuf) | 5059 | 105 | 2014/12/03 | 4 weeks ago | Protocol Buffers for Go with Gadgets. |
| [go](https://github.com/ugorji/go) | 1625 | 55 | 2013/05/30 | 3 months ago | High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support. |
| [csvutil](https://github.com/jszwec/csvutil) | 670 | 10 | 2017/10/30 | 2 months ago | High Performance, idiomatic CSV record encoding and decoding to native Go structures. |
| [colfer](https://github.com/pascaldekloe/colfer) | 649 | 33 | 2015/09/05 | 5 months ago | Code generation for the Colfer binary format. |
| [cbor](https://github.com/fxamacker/cbor) | 375 | 9 | 2019/05/15 | 2 weeks ago | Small, safe, and easy CBOR encoding and decoding library. |
| [go-capnproto](https://github.com/glycerine/go-capnproto) | 280 | 11 | 2013/11/07 | 2 years ago | Cap'n Proto library and parser for go. |
| [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) | 155 | 10 | 2012/12/23 | 3 years ago | GoLang library for working with PHP session format and PHP Serialize/Unserialize functions. |
| [structomap](https://github.com/danhper/structomap) | 131 | 7 | 2015/05/13 | 2 years ago | Library to easily and dynamically generate maps from static structures. |
| [bambam](https://github.com/glycerine/bambam) | 64 | 3 | 2014/09/17 | 5 years ago | generator for Cap'n Proto schemas from go. |
| [asn1](https://github.com/Logicalis/asn1) | 48 | 9 | 2016/02/29 | 2 years ago | Asn.1 BER and DER encoding library for golang. |
| [binstruct](https://github.com/ghostiam/binstruct) | 42 | 2 | 2018/10/23 | 1 month ago | Golang binary decoder for mapping data into the structure. |
| [bel](https://github.com/csweichel/bel) | 18 | 2 | 2019/02/20 | 1 year ago | Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC. |
| [fwencoder](https://github.com/o1egl/fwencoder) | 17 | 2 | 2017/12/25 | 2 years ago | Fixed width file parser (encoding and decoding library) for Go. |
| [pletter](https://github.com/vimeda/pletter) | 17 | 6 | 2019/07/09 | 4 months ago | A standard way to wrap a proto message for message brokers. |
| [elastic](https://github.com/epiclabs-io/elastic) | 16 | 1 | 2020/02/25 | 9 months ago | Convert slices, maps or any other unknown value across different types at run-time, no matter what. |
| [fixedwidth](https://github.com/huydang284/fixedwidth) | 6 | 1 | 2019/08/11 | 2 years ago | Fixed-width text formatting (UTF-8 supported). |
| [unitpacking](https://github.com/recolude/unitpacking) | 4 | 2 | 2021/01/17 | 10 months ago | Library to pack unit vectors into as fewest bytes as possible. |
| [go-lctree](https://github.com/sbourlon/go-lctree) | 3 | 1 | 2020/05/04 | 1 year ago | Provides a CLI and primitives to serialize and deserialize [LeetCode binary trees](https://support.leetcode.com/hc/en-us/articles/360011883654-What-does-1-null-2-3-mean-in-binary-tree-representation). |
**[⬆ back to top](#awesome-go-info)**

# Server Applications
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [etcd](https://github.com/etcd-io/etcd) | 38692 | 1357 | 2013/07/06 | 1 week ago | Highly-available key value store for shared configuration and service discovery. |
| [caddy](https://github.com/caddyserver/caddy) | 36769 | 765 | 2015/01/13 | 1 week ago | Caddy is an alternative, HTTP/2 web server that's easy to configure and use. |
| [minio](https://github.com/minio/minio) | 31403 | 590 | 2015/01/14 | 1 week ago | Minio is a distributed object storage server. |
| [roadrunner](https://github.com/roadrunner-server/roadrunner) | 6233 | 154 | 2017/12/26 | 1 week ago | High-performance PHP application server, load-balancer and process manager. |
| [roadrunner](https://github.com/spiral/roadrunner) | 6169 | 156 | 2017/12/26 | 1 month ago | High-performance PHP application server, load-balancer and process manager. |
| [easegress](https://github.com/megaease/easegress) | 4200 | 99 | 2021/05/28 | 1 week ago | A cloud native high availability/performance traffic orchestration system with observability and extensibility. |
| [sftpgo](https://github.com/drakkan/sftpgo) | 3743 | 69 | 2019/07/20 | 1 week ago | Fully featured and highly configurable SFTP server with optional FTP/S and WebDAV support. It can serve local filesystem and Cloud Storage backends such as S3 and Google Cloud Storage. |
| [devd](https://github.com/cortesi/devd) | 3204 | 72 | 2015/09/27 | 5 months ago | Local webserver for developers. |
| [algernon](https://github.com/xyproto/algernon) | 1906 | 51 | 2015/03/10 | 1 week ago | HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber. |
| [fider](https://github.com/getfider/fider) | 1822 | 35 | 2017/01/17 | 1 week ago | Fider is an open platform to collect and organize customer feedback. |
| [flagr](https://github.com/checkr/flagr) | 1800 | 75 | 2017/10/03 | 2 weeks ago | Flagr is an open-source feature flagging and A/B testing service. |
| [flipt](https://github.com/markphelps/flipt) | 1726 | 15 | 2016/11/05 | 1 week ago | A self contained feature flag solution written in Go and Vue. |
| [trickster](https://github.com/trickstercache/trickster) | 1613 | 44 | 2018/03/29 | 2 weeks ago | HTTP reverse proxy cache and time series accelerator. |
| [discovery](https://github.com/bilibili/discovery) | 1597 | 61 | 2018/04/20 | 3 months ago | A registry for resilient mid-tier load balancing and failover. |
| [jackal](https://github.com/ortuman/jackal) | 1199 | 39 | 2017/11/13 | 1 week ago | An XMPP server written in Go. |
| [go-feature-flag](https://github.com/thomaspoignant/go-feature-flag) | 386 | 1 | 2020/12/11 | 1 week ago | A feature flag solution, with only a YAML file in the backend (S3, GitHub, HTTP, local file ...), no server to install, just add a file in a central system and refer to it. |
| [euterpe](https://github.com/ironsmile/euterpe) | 375 | 9 | 2014/01/01 | 1 month ago | Self-hosted music streaming server with built-in web UI and REST API. |
| [dudeldu](https://github.com/krotik/dudeldu) | 133 | 3 | 2016/09/07 | 2 years ago | A simple SHOUTcast server. |
| [lets-proxy2](https://github.com/rekby/lets-proxy2) | 59 | 5 | 2019/04/12 | 2 months ago | Reverse proxy for handle https with issue certificates in fly from lets-encrypt. |
| [go-proxy-cache](https://github.com/fabiocicerchia/go-proxy-cache) | 39 | 1 | 2020/11/12 | 1 month ago | Simple Reverse Proxy with Caching, written in Go, using Redis. |
| [psql-streamer](https://github.com/blind-oracle/psql-streamer) | 37 | 4 | 2019/04/28 | 1 year ago | Stream database events from PostgreSQL to Kafka. |
| [cortex-tenant](https://github.com/blind-oracle/cortex-tenant) | 32 | 2 | 2020/10/06 | 2 weeks ago | Prometheus remote write proxy that adds add Cortex tenant ID header based on metric labels. |
| [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) | 28 | 1 | 2018/10/23 | 1 year ago | Nginx log parser and exporter to Prometheus. |
| [protoxy](https://github.com/camgraff/protoxy) | 22 | 2 | 2020/09/03 | 1 year ago | A proxy server that converts JSON request bodies to Protocol Buffers. |
| [simple-jwt-provider](https://github.com/leberKleber/simple-jwt-provider) | 21 | 2 | 2019/12/18 | 2 months ago | Simple and lightweight provider which exhibits JWTs, supports login, password-reset (via mail) and user management. |
| [moxy](https://github.com/sinhashubham95/moxy) | 5 | 1 | 2021/07/17 | 2 months ago | Moxy is a simple mocker and proxy application server, you can create mock endpoints as well as proxy requests in case no mock exists for the endpoint. |
| [riemann-relay](https://github.com/blind-oracle/riemann-relay) | 1 | 1 | 2019/04/23 | 2 years ago | Relay to load-balance Riemann events and/or convert them to Carbon. |
**[⬆ back to top](#awesome-go-info)**

# Stream Processing
        
*Libraries and tools for stream processing and reactive programming.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-streams](https://github.com/reugn/go-streams) | 834 | 22 | 2019/04/30 | 1 month ago | Go stream processing library. |
| [machine](https://github.com/whitaker-io/machine) | 101 | 6 | 2020/10/13 | 1 month ago | Go library for writing and generating stream workers with built in metrics and traceability. |
| [stream](https://github.com/youthlin/stream) | 50 | 2 | 2020/11/12 | 1 year ago | Go Stream, like Java 8 Stream: Filter/Map/FlatMap/Peek/Sorted/ForEach/Reduce... |
**[⬆ back to top](#awesome-go-info)**

# Template Engines
        
*Libraries and tools for templating and lexing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gofpdf](https://github.com/jung-kurt/gofpdf) | 3933 | 105 | 2015/03/13 | 3 months ago | PDF document generator with high level support for text, drawing and images. |
| [sprig](https://github.com/Masterminds/sprig) | 2775 | 36 | 2013/11/22 | 2 weeks ago | Useful template functions for Go templates. |
| [quicktemplate](https://github.com/valyala/quicktemplate) | 2321 | 59 | 2016/03/06 | 5 months ago | Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it. |
| [pongo2](https://github.com/flosch/pongo2) | 2159 | 64 | 2013/08/23 | 1 week ago | Django-like template-engine for Go. |
| [hero](https://github.com/shiyanhui/hero) | 1489 | 43 | 2017/01/15 | 2 years ago | Hero is a handy, fast and powerful go template engine. |
| [mustache](https://github.com/hoisie/mustache) | 1032 | 35 | 2009/12/30 | 1 month ago | Go implementation of the Mustache template language. |
| [amber](https://github.com/eknkc/amber) | 889 | 19 | 2012/10/31 | 1 year ago | Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade. |
| [jet](https://github.com/CloudyKit/jet) | 880 | 23 | 2016/03/31 | 3 months ago | Jet template engine. |
| [ace](https://github.com/yosssi/ace) | 810 | 22 | 2014/07/13 | 3 years ago | Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold. |
| [gorazor](https://github.com/sipin/gorazor) | 794 | 57 | 2014/05/01 | 1 year ago | Razor view engine for Golang. |
| [fasttemplate](https://github.com/valyala/fasttemplate) | 567 | 20 | 2015/08/19 | 1 year ago | Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](https://golang.org/pkg/text/template/). |
| [maroto](https://github.com/johnfercher/maroto) | 551 | 11 | 2019/05/20 | 3 weeks ago | A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple. |
| [ego](https://github.com/benbjohnson/ego) | 516 | 14 | 2014/02/23 | 2 months ago | Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled. |
| [raymond](https://github.com/aymerick/raymond) | 461 | 11 | 2015/04/22 | 3 months ago | Complete handlebars implementation in Go. |
| [goview](https://github.com/foolin/goview) | 268 | 6 | 2019/04/14 | 1 month ago | Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application. |
| [soy](https://github.com/robfig/soy) | 159 | 13 | 2013/12/15 | 4 months ago | Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/). |
| [liquid](https://github.com/osteele/liquid) | 147 | 6 | 2017/06/26 | 2 weeks ago | Go implementation of Shopify Liquid templates. |
| [velvet](https://github.com/gobuffalo/velvet) | 74 | 7 | 2016/12/29 | 4 years ago | Complete handlebars implementation in Go. |
| [kasia.go](https://github.com/ziutek/kasia.go) | 73 | 2 | 2010/12/07 | 6 years ago | Templating system for HTML and other text documents - go implementation. |
| [extemplate](https://github.com/dannyvankooten/extemplate) | 42 | 4 | 2018/08/10 | 8 months ago | Tiny wrapper around html/template to allow for easy file-based template inheritance. |
| [gospin](https://github.com/m1/gospin) | 35 | 4 | 2019/02/22 | 9 months ago | Article spinning and spintax/spinning syntax engine, useful for A/B, testing pieces of text/articles and creating more natural conversations. |
| [damsel](https://github.com/dskinner/damsel) | 25 | 4 | 2012/05/02 | 5 years ago | Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others. |
| [tbd](https://github.com/lucasepe/tbd) | 17 | 1 | 2021/05/21 | 5 months ago | A really simple way to create text templates with placeholders - exposes extra builtin Git repo metadata. |
**[⬆ back to top](#awesome-go-info)**

# Testing
        
*Libraries for testing codebases and generating test data.*
**[⬆ back to top](#awesome-go-info)**

## Fail injection
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [failpoint](https://github.com/pingcap/failpoint) | 653 | 102 | 2019/04/02 | 3 weeks ago | An implementation of [failpoints](https://www.freebsd.org/cgi/man.cgi?query=fail) for Golang. |
**[⬆ back to top](#awesome-go-info)**

## Mock
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [chromedp](https://github.com/chromedp/chromedp) | 7216 | 154 | 2017/01/24 | 1 week ago | a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol. |
| [mock](https://github.com/golang/mock) | 6730 | 85 | 2015/06/12 | 1 month ago | Mocking framework for the Go programming language. |
| [go-fuzz](https://github.com/dvyukov/go-fuzz) | 4250 | 84 | 2015/04/15 | 2 months ago | Randomized testing system. |
| [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) | 4128 | 42 | 2014/02/07 | 2 months ago | Mock SQL driver for testing database interactions. |
| [mockery](https://github.com/vektra/mockery) | 3246 | 30 | 2014/09/02 | 3 weeks ago | Tool to generate Go interfaces. |
| [rod](https://github.com/go-rod/rod) | 2115 | 33 | 2020/01/21 | 1 week ago | A Devtools driver to make web automation and scraping easy. |
| [selenoid](https://github.com/aerokube/selenoid) | 2095 | 97 | 2016/08/22 | 2 months ago | alternative Selenium hub server that launches browsers within containers. |
| [hoverfly](https://github.com/SpectoLabs/hoverfly) | 1829 | 60 | 2015/11/30 | 2 months ago | HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI. |
| [gock](https://github.com/h2non/gock) | 1568 | 19 | 2016/03/02 | 6 months ago | Versatile HTTP mocking made easy. |
| [httpmock](https://github.com/jarcoal/httpmock) | 1301 | 8 | 2014/02/24 | 1 month ago | Easy mocking of HTTP responses from external resources. |
| [gofuzz](https://github.com/google/gofuzz) | 1220 | 29 | 2014/07/31 | 5 months ago | Library for populating go objects with random values. |
| [playwright-go](https://github.com/playwright-community/playwright-go) | 617 | 19 | 2020/08/16 | 2 weeks ago | browser automation library to control Chromium, Firefox and WebKit with a single API. |
| [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) | 614 | 10 | 2014/05/21 | 2 weeks ago | Tool for generating self-contained mock objects. |
| [playwright-go](https://github.com/mxschmitt/playwright-go) | 610 | 18 | 2020/08/16 | 2 months ago | browser automation library to control Chromium, Firefox and WebKit with a single API. |
| [cdp](https://github.com/mafredri/cdp) | 598 | 19 | 2017/03/12 | 7 months ago | Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it. |
| [minimock](https://github.com/gojuno/minimock) | 446 | 10 | 2016/08/03 | 4 months ago | Mock generator for Go interfaces. |
| [go-txdb](https://github.com/DATA-DOG/go-txdb) | 418 | 7 | 2015/07/08 | 1 month ago | Single transaction based database driver mainly for testing purposes. |
| [ggr](https://github.com/aerokube/ggr) | 279 | 25 | 2016/06/16 | 2 weeks ago | a lightweight server that routes and proxies Selenium WebDriver requests to multiple Selenium hubs. |
| [tavor](https://github.com/zimmski/tavor) | 230 | 13 | 2014/05/18 | 3 years ago | Generic fuzzing and delta-debugging framework. |
| [govcr](https://github.com/seborama/govcr) | 100 | 2 | 2016/07/10 | 2 years ago | HTTP mock for Golang: record and replay HTTP interactions for offline testing. |
| [timex](https://github.com/cabify/timex) | 58 | 79 | 2020/01/02 | 1 year ago | A test-friendly replacement for the native `time` package. |
| [go-localstack](https://github.com/elgohr/go-localstack) | 33 | 1 | 2020/03/18 | 1 week ago | Tool for using localstack in AWS testing. |
| [mockhttp](https://github.com/tv42/mockhttp) | 22 | 2 | 2011/06/11 | 7 years ago | Mock object for Go http.ResponseWriter. |
| [mockit](https://github.com/pasdam/mockit) | 8 | 1 | 2020/01/01 | 1 month ago | Allows functions and method easy mocking, without defining new types; it's similar to Mockito for Java. |
**[⬆ back to top](#awesome-go-info)**

## Testing Frameworks
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [testify](https://github.com/stretchr/testify) | 15364 | 175 | 2012/10/16 | 3 weeks ago | Sacred extension to the standard go testing package. |
| [go-cmp](https://github.com/google/go-cmp) | 2688 | 28 | 2017/07/07 | 3 weeks ago | Package for comparing Go values in tests. |
| [httpexpect](https://github.com/gavv/httpexpect) | 1846 | 37 | 2016/04/29 | 5 months ago | Concise, declarative, and easy to use end-to-end HTTP and REST API testing. |
| [godog](https://github.com/cucumber/godog) | 1560 | 98 | 2015/06/10 | 2 weeks ago | Cucumber or Behat like BDD framework for Go. |
| [is](https://github.com/matryer/is) | 1332 | 18 | 2016/12/06 | 2 months ago | Professional lightweight testing mini-framework for Go. |
| [go-vcr](https://github.com/dnaeon/go-vcr) | 846 | 14 | 2015/12/14 | 4 months ago | Record and replay your HTTP interactions for fast, deterministic and accurate tests. |
| [goblin](https://github.com/franela/goblin) | 828 | 16 | 2013/09/19 | 4 months ago | Mocha like testing framework fo Go. |
| [testfixtures](https://github.com/go-testfixtures/testfixtures) | 737 | 6 | 2016/04/05 | 3 weeks ago | A helper for Rails' like test fixtures to test database applications. |
| [baloo](https://github.com/h2non/baloo) | 717 | 11 | 2016/05/29 | 1 month ago | Expressive and versatile end-to-end HTTP API testing made easy. |
| [gnomock](https://github.com/orlangure/gnomock) | 648 | 13 | 2020/01/31 | 1 week ago | integration testing with real dependencies (database, cache, even Kubernetes or AWS) running in Docker, without mocks. |
| [go-mutesting](https://github.com/zimmski/go-mutesting) | 505 | 12 | 2014/12/26 | 2 months ago | Mutation testing for Go source code. |
| [goc](https://github.com/qiniu/goc) | 483 | 21 | 2020/05/07 | 3 weeks ago | Goc is a comprehensive coverage testing system for The Go Programming Language. |
| [gofight](https://github.com/appleboy/gofight) | 398 | 14 | 2016/03/29 | 7 months ago | API Handler Testing for Golang Router framework. |
| [embedded-postgres](https://github.com/fergusstrange/embedded-postgres) | 340 | 4 | 2019/11/16 | 4 weeks ago | Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test. |
| [testza](https://github.com/MarvinJWendt/testza) | 312 | 3 | 2021/07/05 | 1 week ago | Full-featured test framework with nice colorized output. |
| [gotest.tools](https://github.com/gotestyourself/gotest.tools) | 281 | 8 | 2017/08/08 | 1 month ago | A collection of packages to augment the go testing package and support common patterns. |
| [frisby](https://github.com/hofstadter-io/frisby) | 270 | 10 | 2015/09/15 | 1 year ago | REST API testing framework. |
| [go-testdeep](https://github.com/maxatome/go-testdeep) | 257 | 3 | 2018/05/26 | 1 week ago | Extremely flexible golang deep comparison, extends the go testing package. |
| [go-carpet](https://github.com/msoap/go-carpet) | 218 | 4 | 2016/02/28 | 1 month ago | Tool for viewing test coverage in terminal. |
| [endly](https://github.com/viant/endly) | 202 | 17 | 2017/08/28 | 2 months ago | Declarative end to end functional testing. |
| [cupaloy](https://github.com/bradleyjkemp/cupaloy) | 201 | 3 | 2017/08/07 | 3 months ago | Simple snapshot testing addon for your test framework. |
| [charlatan](https://github.com/percolate/charlatan) | 194 | 52 | 2017/10/06 | 2 years ago | Tool to generate fake interface implementations for tests. |
| [commander](https://github.com/commander-cli/commander) | 193 | 7 | 2019/02/22 | 1 month ago | Tool for testing cli applications on windows, linux and osx. |
| [dbcleaner](https://github.com/khaiql/dbcleaner) | 134 | 4 | 2017/01/17 | 3 months ago | Clean database for testing purpose, inspired by `database_cleaner` in Ruby. |
| [gospec](https://github.com/luontola/gospec) | 114 | 4 | 2009/11/24 | 7 years ago | BDD-style testing framework for the Go programming language. |
| [wstest](https://github.com/posener/wstest) | 86 | 2 | 2017/03/31 | 1 year ago | Websocket client for unit-testing a websocket http.Handler. |
| [gocrest](https://github.com/corbym/gocrest) | 81 | 4 | 2017/12/23 | 1 year ago | Composable hamcrest-like matchers for Go assertions. |
| [testcase](https://github.com/adamluzsi/testcase) | 80 | 3 | 2019/04/22 | 1 week ago | Idiomatic testing framework for Behavior Driven Development. |
| [jsonassert](https://github.com/kinbiko/jsonassert) | 76 | 2 | 2018/10/26 | 2 months ago | Package for verifying that your JSON payloads are serialized correctly. |
| [go-hit](https://github.com/Eun/go-hit) | 66 | 1 | 2019/06/04 | 2 weeks ago | Hit is an http integration test framework written in golang. |
| [restit](https://github.com/go-restit/restit) | 55 | 7 | 2014/06/25 | 2 years ago | Go micro framework to help writing RESTful API integration test. |
| [gospecify](https://github.com/stesla/gospecify) | 52 | 6 | 2009/11/20 | 10 years ago | This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec. |
| [covergates](https://github.com/covergates/covergates) | 47 | 5 | 2020/05/29 | 1 year ago | Self-hosted code coverage report review and management service. |
| [gomatch](https://github.com/jfilipczyk/gomatch) | 41 | 2 | 2019/01/27 | 1 year ago | library created for testing JSON against patterns. |
| [dsunit](https://github.com/viant/dsunit) | 39 | 10 | 2016/06/13 | 2 years ago | Datastore testing for SQL, NoSQL, structured files. |
| [assert](https://github.com/go-playground/assert) | 37 | 2 | 2015/07/20 | 1 year ago | Basic Assertion Library used along side native go testing, with building blocks for custom assertions. |
| [hamcrest](https://github.com/rdrdr/hamcrest) | 27 | 3 | 2010/12/22 | 1 year ago | fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results. |
| [schema](https://github.com/jgroeneveld/schema) | 16 | 1 | 2015/08/13 | 2 years ago | Quick and easy expression matching for JSON schemas used in requests and responses. |
| [flute](https://github.com/suzuki-shunsuke/flute) | 15 | 0 | 2019/07/06 | 3 weeks ago | HTTP client testing framework. |
| [gogiven](https://github.com/corbym/gogiven) | 11 | 5 | 2017/12/31 | 6 months ago | YATSPEC-like BDD testing framework for Go. |
| [testsql](https://github.com/zhulongcheng/testsql) | 11 | 2 | 2018/09/22 | 2 years ago | Generate test data from SQL files before testing and clear it after finished. |
| [gosuite](https://github.com/pavlo/gosuite) | 10 | 2 | 2016/10/15 | 5 years ago | Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests. |
| [badio](https://github.com/cavaliergopher/badio) | 9 | 0 | 2016/02/11 | 6 years ago | Extensions to Go's `testing/iotest` package. |
| [biff](https://github.com/fulldump/biff) | 9 | 1 | 2018/03/28 | 7 months ago | Bifurcation testing framework, BDD compatible. |
| [stop-and-go](https://github.com/elgohr/stop-and-go) | 6 | 1 | 2020/11/06 | 5 months ago | Testing helper for concurrency. |
| [trial](https://github.com/jgroeneveld/trial) | 5 | 0 | 2015/06/18 | 2 years ago | Quick and easy extendable assertions without introducing much boilerplate. |
| [fixenv](https://github.com/rekby/fixenv) | 4 | 1 | 2021/08/27 | 5 months ago | Fixture manage engine, inspired by pytest fixtures. |
| [tt](https://github.com/vcaesar/tt) | 4 | 1 | 2018/04/03 | 2 weeks ago | Simple and colorful test tools. |
| [go-testpredicate](https://github.com/maargenton/go-testpredicate) | 2 | 2 | 2018/11/23 | 2 months ago | Test predicate style assertions library with extensive diagnostics output. |
| [omg.testingtools](https://github.com/dedalqq/omg.testingtools) | 0 | 1 | 2021/10/13 | 4 months ago | The simple library for change a values of private fields for testing. |
**[⬆ back to top](#awesome-go-info)**

# Text Processing
        
*Libraries for parsing and manipulating texts.*
**[⬆ back to top](#awesome-go-info)**

## Specific Formats
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [colly](https://github.com/gocolly/colly) | 15878 | 325 | 2017/09/29 | 1 week ago | Fast and Elegant Scraping Framework for Gophers. |
| [goquery](https://github.com/PuerkitoBio/goquery) | 11076 | 255 | 2012/08/29 | 1 month ago | GoQuery brings a syntax and a set of features similar to jQuery to the Go language. |
| [blackfriday](https://github.com/russross/blackfriday) | 4870 | 90 | 2011/05/27 | 4 months ago | Markdown processor in Go. |
| [sh](https://github.com/mvdan/sh) | 4444 | 53 | 2016/01/16 | 1 week ago | Shell parser and formatter. |
| [toml](https://github.com/BurntSushi/toml) | 3752 | 84 | 2013/02/26 | 1 month ago | TOML configuration format (encoder/decoder with reflection). |
| [go-humanize](https://github.com/dustin/go-humanize) | 3033 | 34 | 2012/01/13 | 1 week ago | Formatters for time, numbers, and memory size to human readable format. |
| [bluemonday](https://github.com/microcosm-cc/bluemonday) | 2176 | 30 | 2013/11/20 | 1 month ago | HTML Sanitizer. |
| [gofeed](https://github.com/mmcdole/gofeed) | 1803 | 42 | 2016/01/23 | 1 month ago | Parse RSS and Atom feeds in Go. |
| [inject](https://github.com/facebookarchive/inject) | 1363 | 45 | 2013/10/21 | 3 years ago | Package inject provides a reflect based injector. |
| [go-toml](https://github.com/pelletier/go-toml) | 1157 | 30 | 2013/02/24 | 3 weeks ago | Go library for the TOML format with query support and handy cli tools. |
| [commonregex](https://github.com/mingrammer/commonregex) | 795 | 21 | 2017/03/23 | 2 years ago | A collection of common regular expressions for Go. |
| [slug](https://github.com/gosimple/slug) | 783 | 13 | 2014/03/31 | 1 month ago | URL-friendly slugify with multiple languages support. |
| [dataflowkit](https://github.com/slotix/dataflowkit) | 508 | 22 | 2017/02/09 | 1 year ago | Web scraping Framework to turn websites into structured data. |
| [mxj](https://github.com/clbanning/mxj) | 505 | 27 | 2014/02/03 | 1 month ago | Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages. |
| [gographviz](https://github.com/awalterschulze/gographviz) | 462 | 12 | 2015/03/14 | 5 months ago | Parses the Graphviz DOT language. |
| [htmlquery](https://github.com/antchfx/htmlquery) | 451 | 13 | 2017/12/05 | 2 months ago | An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression. |
| [omniparser](https://github.com/jf-tech/omniparser) | 412 | 9 | 2020/08/16 | 2 months ago | A versatile ETL library that parses text input (CSV/txt/JSON/XML/EDI/X12/EDIFACT/etc) in streaming fashion and transforms data into JSON output using data-driven schema. |
| [go-runewidth](https://github.com/mattn/go-runewidth) | 408 | 14 | 2013/06/21 | 2 months ago | Functions to get fixed width of the character or string. |
| [gotext](https://github.com/leonelquinteros/gotext) | 319 | 7 | 2016/06/19 | 1 month ago | GNU gettext utilities for Go. |
| [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) | 316 | 10 | 2018/05/15 | 3 weeks ago | Convert HTML to Markdown. Even works with entire websites and can be extended through rules. |
| [goq](https://github.com/andrewstuart/goq) | 210 | 7 | 2017/02/20 | 5 months ago | Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery). |
| [goribot](https://github.com/zhshch2002/goribot) | 205 | 11 | 2019/09/08 | 1 year ago | [Crawler/Scraper for Golang]🕷A lightweight distributed friendly Golang crawler framework.一个轻量的分布式友好的 Golang 爬虫框架。 |
| [go-nmea](https://github.com/adrianmo/go-nmea) | 166 | 8 | 2015/07/22 | 1 month ago | NMEA parser library for the Go language. |
| [sdp](https://github.com/gortc/sdp) | 112 | 8 | 2016/05/13 | 1 year ago | SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)]. |
| [podcast](https://github.com/eduncan911/podcast) | 104 | 5 | 2017/02/02 | 1 year ago | iTunes Compliant and RSS 2. |
| [go-zero-width](https://github.com/trubitsyn/go-zero-width) | 97 | 1 | 2018/06/18 | 1 year ago | Zero-width character detection and removal for Go. |
| [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) | 91 | 10 | 2016/07/05 | 3 weeks ago | Editorconfig file parser and manipulator for Go. |
| [align](https://github.com/Guitarbum722/align) | 73 | 5 | 2017/04/29 | 5 months ago | A general purpose application that aligns text. |
| [go-slugify](https://github.com/mozillazg/go-slugify) | 72 | 2 | 2016/07/16 | 1 year ago | Make pretty slug with multiple languages support. |
| [go-vcard](https://github.com/emersion/go-vcard) | 66 | 4 | 2017/03/21 | 9 months ago | Parse and format vCard. |
| [genex](https://github.com/alixaxel/genex) | 63 | 3 | 2015/03/09 | 2 years ago | Count and expand Regular Expressions into all matching Strings. |
| [goregen](https://github.com/zach-klippenstein/goregen) | 61 | 3 | 2014/12/27 | 7 months ago | Library for generating random strings from regular expressions. |
| [did](https://github.com/ockam-network/did) | 58 | 15 | 2018/11/02 | 1 year ago | DID (Decentralized Identifiers) Parser and Stringer in Go. |
| [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) | 58 | 3 | 2017/11/15 | 1 month ago | Fixed-width text formatting (encoder/decoder with reflection). |
| [guesslanguage](https://github.com/endeveit/guesslanguage) | 53 | 1 | 2014/12/16 | 4 years ago | Functions to determine the natural language of a unicode text. |
| [allot](https://github.com/sbstjn/allot) | 51 | 2 | 2016/10/16 | 2 weeks ago | Placeholder and wildcard text parsing for CLI tools and bots. |
| [pagser](https://github.com/foolin/pagser) | 50 | 2 | 2020/04/19 | 1 month ago | Pagser is a simple, extensible, configurable parse and deserialize html page to struct based on goquery and struct tags for golang crawler. |
| [bafi](https://github.com/mmalcek/bafi) | 48 | 4 | 2021/07/13 | 2 months ago | Universal JSON, BSON, YAML, XML translator to ANY format using templates. |
| [gonameparts](https://github.com/polera/gonameparts) | 34 | 1 | 2015/05/17 | 2 years ago | Parses human names into individual name parts. |
| [slugify](https://github.com/avelino/slugify) | 31 | 3 | 2015/04/13 | 3 years ago | Go slugify application that handles string. |
| [normalize](https://github.com/avito-tech/normalize) | 24 | 3 | 2021/03/22 | 10 months ago | Sanitize, normalize and compare fuzzy text. |
| [codetree](https://github.com/aerogo/codetree) | 19 | 3 | 2016/11/26 | 2 years ago | Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure. |
| [go-wildcard](https://github.com/IGLOU-EU/go-wildcard) | 12 | 3 | 2021/03/28 | 2 weeks ago | Simple and lightweight wildcard pattern matching. |
| [enca](https://github.com/endeveit/enca) | 11 | 1 | 2014/12/17 | 6 years ago | Minimal cgo bindings for [libenca](https://cihar.com/software/enca/). |
| [syndfeed](https://github.com/zhengchun/syndfeed) | 8 | 1 | 2017/04/07 | 3 years ago | A syndication feed for Atom 1.0 and RSS 2.0. |
| [ltsv](https://github.com/Wing924/ltsv) | 7 | 2 | 2019/05/12 | 2 years ago | High performance [LTSV (Labeled Tab Separated Value)](http://ltsv.org/) reader for Go. |
| [bbConvert](https://github.com/CalebQ42/bbConvert) | 6 | 1 | 2016/04/15 | 5 years ago | Converts bbCode to HTML that allows you to add support for custom bbCode tags. |
| [doi](https://github.com/hscells/doi) | 6 | 1 | 2017/08/02 | 4 years ago | Document object identifier (doi) parser in Go. |
| [encoding](https://github.com/ake-persson/encoding) | 5 | 1 | 2018/04/06 | 2 years ago | Package provides a generic interface to encoders and decodersa. |
| [go-output-format](https://github.com/drewstinnett/go-output-format) | 5 | 1 | 2021/04/08 | 3 months ago | Output go structures into multiple formats (YAML/JSON/etc) in your command line app. |
**[⬆ back to top](#awesome-go-info)**

## Utility
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [xurls](https://github.com/mvdan/xurls) | 858 | 18 | 2015/01/12 | 3 weeks ago | Extract urls from text. |
| [gotabulate](https://github.com/bndr/gotabulate) | 274 | 9 | 2014/08/21 | 1 year ago | Easily pretty-print your tabular data with Go. |
| [radix](https://github.com/yourbasic/radix) | 172 | 7 | 2017/06/09 | 4 years ago | fast string sorting algorithm. |
| [regroup](https://github.com/oriser/regroup) | 100 | 1 | 2020/09/08 | 6 months ago | Match regex expression named groups into go struct using struct tags and automatic parsing. |
| [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) | 42 | 2 | 2018/09/09 | 7 months ago | A sanitization-based swear filter for Go. |
| [parth](https://github.com/codemodus/parth) | 40 | 4 | 2015/04/06 | 3 years ago | URL path segmentation parsing. |
| [xj2go](https://github.com/wk30/xj2go) | 22 | 2 | 2017/09/19 | 4 months ago | Convert xml or json to go struct. |
| [tagify](https://github.com/zoomio/tagify) | 19 | 2 | 2018/03/20 | 1 month ago | Produces a set of tags from given source. |
| [kace](https://github.com/codemodus/kace) | 17 | 2 | 2015/06/04 | 3 years ago | Common case conversions covering common initialisms. |
| [TySug](https://github.com/Dynom/TySug) | 10 | 1 | 2018/06/05 | 1 year ago | Alternative suggestions with respect to keyboard layouts. |
| [parseargs-go](https://github.com/txgruppi/parseargs-go) | 9 | 0 | 2016/02/24 | 5 years ago | string argument parser that understands quotes and backslashes. |
| [textwrap](https://github.com/isbm/textwrap) | 2 | 2 | 2019/07/26 | 2 years ago | Implementation of `textwrap` module from Python. |
**[⬆ back to top](#awesome-go-info)**

# Third-party APIs
        
*Libraries for accessing third party APIs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-github](https://github.com/google/go-github) | 8145 | 213 | 2013/05/24 | 1 week ago | Go library for accessing the GitHub REST API v3. |
| [aws-sdk-go](https://github.com/aws/aws-sdk-go) | 7357 | 255 | 2014/12/05 | 1 week ago | The official AWS SDK for the Go programming language. |
| [slack](https://github.com/slack-go/slack) | 3790 | 59 | 2015/01/24 | 1 week ago | Slack API in Go. |
| [google-api-go-client](https://github.com/googleapis/google-api-go-client) | 2897 | 174 | 2014/11/24 | 1 week ago | Auto-generated Google APIs for Go. |
| [google-cloud-go](https://github.com/googleapis/google-cloud-go) | 2750 | 249 | 2014/05/09 | 1 week ago | Google Cloud APIs Go Client Library. |
| [discordgo](https://github.com/bwmarrin/discordgo) | 2709 | 57 | 2015/11/01 | 1 week ago | Go bindings for the Discord Chat API. |
| [minio-go](https://github.com/minio/minio-go) | 1506 | 48 | 2015/05/02 | 2 weeks ago | Minio Go Library for Amazon S3 compatible cloud storage. |
| [stripe-go](https://github.com/stripe/stripe-go) | 1498 | 41 | 2014/06/05 | 1 week ago | Go client for the Stripe API. |
| [go-twitter](https://github.com/dghubble/go-twitter) | 1388 | 28 | 2015/04/11 | 3 months ago | Go client library for the Twitter v1.1 APIs. |
| [anaconda](https://github.com/ChimeraCoder/anaconda) | 1108 | 21 | 2013/03/04 | 6 months ago | Go client library for the Twitter 1.1 API. |
| [facebook](https://github.com/huandu/facebook) | 1032 | 117 | 2012/07/28 | 1 month ago | Go Library that supports the Facebook Graph API. |
| [githubv4](https://github.com/shurcooL/githubv4) | 842 | 20 | 2017/05/27 | 1 month ago | Go library for accessing the GitHub GraphQL API v4. |
| [webhooks](https://github.com/go-playground/webhooks) | 693 | 15 | 2015/10/25 | 1 week ago | Webhook receiver for GitHub and Bitbucket. |
| [paypal](https://github.com/plutov/paypal) | 463 | 26 | 2015/10/14 | 2 months ago | Wrapper for PayPal payment API. |
| [geo-golang](https://github.com/codingsince1985/geo-golang) | 425 | 13 | 2014/12/04 | 6 months ago | Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](https://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](https://opencagedata.com/api), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs. |
| [ethrpc](https://github.com/onrik/ethrpc) | 222 | 14 | 2017/01/24 | 1 year ago | Go bindings for Ethereum JSON RPC API. |
| [go-marathon](https://github.com/gambol99/go-marathon) | 196 | 14 | 2015/02/11 | 1 year ago | Go library for interacting with Mesosphere's Marathon PAAS. |
| [trello](https://github.com/adlio/trello) | 193 | 8 | 2016/09/24 | 8 months ago | Go wrapper for the Trello API. |
| [twitter-scraper](https://github.com/n0madic/twitter-scraper) | 178 | 5 | 2018/11/29 | 3 weeks ago | Scrape the Twitter Frontend API without authentication and limits. |
| [medium-sdk-go](https://github.com/Medium/medium-sdk-go) | 131 | 139 | 2015/09/26 | 3 years ago | Golang SDK for Medium's OAuth2 API. |
| [gostorm](https://github.com/jsgilmore/gostorm) | 128 | 12 | 2013/07/22 | 4 years ago | GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. |
| [lark](https://github.com/chyroc/lark) | 121 | 3 | 2021/04/21 | 3 weeks ago | [Feishu](https://open.feishu.cn/)/[Lark](https://open.larksuite.com/) Open API Go SDK, Support ALL Open API and Event Callback. |
| [go-trending](https://github.com/andygrunwald/go-trending) | 117 | 7 | 2015/07/04 | 3 months ago | Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github. |
| [wit-go](https://github.com/wit-ai/wit-go) | 112 | 20 | 2018/08/20 | 5 months ago | Go client for wit.ai HTTP API. |
| [hipchat](https://github.com/daneharrigan/hipchat) | 110 | 7 | 2013/04/28 | 4 years ago | A golang package to communicate with HipChat over XMPP. |
| [pushover](https://github.com/gregdel/pushover) | 109 | 5 | 2015/02/19 | 3 months ago | Go wrapper for the Pushover API. |
| [hipchat](https://github.com/andybons/hipchat) | 104 | 6 | 2012/10/20 | 5 years ago | This project implements a golang client library for the Hipchat API. |
| [cachet](https://github.com/andygrunwald/cachet) | 90 | 7 | 2015/10/31 | 7 months ago | Go client library for [Cachet (open source status page system)](https://cachethq.io/). |
| [simples3](https://github.com/rhnvrm/simples3) | 75 | 2 | 2018/12/06 | 1 month ago | Simple no frills AWS S3 Library using REST with V4 Signing written in Go. |
| [lark](https://github.com/go-lark/lark) | 72 | 1 | 2021/04/20 | 3 weeks ago | An easy-to-use unofficial SDK for [Feishu](https://open.feishu.cn/) and [Lark](https://open.larksuite.com/) Open Platform. |
| [gosip](https://github.com/koltyakov/gosip) | 71 | 4 | 2019/01/26 | 1 month ago | Go client library SharePoint API. |
| [igdb](https://github.com/Henry-Sarabia/igdb) | 71 | 2 | 2017/08/24 | 11 months ago | Go client for the [Internet Game Database API](https://api.igdb.com/). |
| [go-circleci](https://github.com/jszwedko/go-circleci) | 61 | 3 | 2015/08/14 | 2 years ago | Go client library for interacting with CircleCI's API. |
| [go-unsplash](https://github.com/hbagdi/go-unsplash) | 58 | 2 | 2017/01/19 | 10 months ago | Go client library for the [Unsplash.com](https://unsplash.com) API. |
| [gogtrends](https://github.com/groovili/gogtrends) | 57 | 1 | 2018/12/27 | 5 months ago | Google Trends Unofficial API. |
| [clarifai-go](https://github.com/Clarifai/clarifai-go) | 55 | 37 | 2015/09/28 | 4 years ago | Go client library for interfacing with the Clarifai API. |
| [megos](https://github.com/andygrunwald/megos) | 54 | 5 | 2015/10/02 | 7 months ago | Client library for accessing an [Apache Mesos](https://mesos.apache.org/) cluster. |
| [go-amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) | 51 | 1 | 2016/11/15 | 3 years ago | Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html). |
| [gads](https://github.com/emiddleton/gads) | 49 | 7 | 2014/01/20 | 2 months ago | Google Adwords Unofficial API. |
| [ynab.go](https://github.com/brunomvsouza/ynab.go) | 49 | 2 | 2018/07/13 | 5 months ago | Go wrapper for the YNAB API. |
| [golang-tmdb](https://github.com/cyruzin/golang-tmdb) | 48 | 1 | 2019/01/11 | 2 weeks ago | Golang wrapper for The Movie Database API v3. |
| [uptimerobot](https://github.com/bitfield/uptimerobot) | 45 | 4 | 2018/05/29 | 1 year ago | Go wrapper and command-line client for the Uptime Robot v2 API. |
| [fcm](https://github.com/maddevsio/fcm) | 44 | 5 | 2017/01/06 | 1 year ago | Go library for Firebase Cloud Messaging. |
| [gomusicbrainz](https://github.com/michiwend/gomusicbrainz) | 44 | 6 | 2014/09/10 | 1 year ago | Go MusicBrainz WS2 client library. |
| [go-xkcd](https://github.com/nishanths/go-xkcd) | 43 | 3 | 2016/02/26 | 3 months ago | Go client for the xkcd API. |
| [mixpanel](https://github.com/dukex/mixpanel) | 42 | 4 | 2014/05/20 | 4 weeks ago | Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications. |
| [go-spotify](https://github.com/rapito/go-spotify) | 42 | 3 | 2014/10/30 | 1 year ago | Go Library to access Spotify WEB API. |
| [go-postman-collection](https://github.com/rbretecher/go-postman-collection) | 39 | 2 | 2019/11/16 | 1 week ago | Go module to work with [Postman Collections](https://learning.getpostman.com/docs/postman/collections/creating-collections/) (compatible with Insomnia). |
| [golyrics](https://github.com/mamal72/golyrics) | 36 | 4 | 2016/11/18 | 3 years ago | Golyrics is a Go library to fetch music lyrics data from the Wikia website. |
| [airtable](https://github.com/mehanizm/airtable) | 32 | 1 | 2020/04/12 | 1 month ago | Go client library for the [Airtable API](https://airtable.com/api). |
| [translate](https://github.com/nuveo/translate) | 31 | 31 | 2015/07/13 | 6 years ago | Go online translation package. |
| [gami](https://github.com/bit4bit/gami) | 30 | 4 | 2014/05/14 | 3 years ago | Go library for Asterisk Manager Interface. |
| [gcm](https://github.com/TheOrioli/gcm) | 30 | 3 | 2015/11/09 | 6 years ago | Go library for Google Cloud Messaging. |
| [patreon-go](https://github.com/mxpv/patreon-go) | 27 | 4 | 2017/08/06 | 2 years ago | Go library for Patreon API. |
| [go-myanimelist](https://github.com/nstratos/go-myanimelist) | 25 | 1 | 2015/05/03 | 2 months ago | Go client library for accessing the [MyAnimeList API](https://myanimelist.net/apiconfig/references/api/v2). |
| [go-steam](https://github.com/sostronk/go-steam) | 24 | 11 | 2014/11/23 | 5 months ago | Go Library to interact with Steam game servers. |
| [google-play-scraper](https://github.com/n0madic/google-play-scraper) | 23 | 1 | 2019/09/20 | 1 month ago | Get data from Google Play Store. |
| [lastpass-go](https://github.com/ansd/lastpass-go) | 23 | 4 | 2019/07/11 | 2 weeks ago | Go client library for the [LastPass](https://www.lastpass.com/) API. |
| [go-shopify](https://github.com/rapito/go-shopify) | 22 | 2 | 2014/10/28 | 1 year ago | Go Library to make CRUD request to the Shopify API. |
| [go-twitch](https://github.com/knspriggs/go-twitch) | 21 | 5 | 2016/06/28 | 4 years ago | Go client for interacting with the Twitch v3 API. |
| [codeship-go](https://github.com/codeship/codeship-go) | 18 | 20 | 2017/09/08 | 1 year ago | Go client library for interacting with Codeship's API v2. |
| [brewerydb](https://github.com/naegelejd/brewerydb) | 17 | 3 | 2015/04/15 | 6 years ago | Go library for accessing the BreweryDB API. |
| [textbelt](https://github.com/farmergreg/textbelt) | 17 | 2 | 2015/09/01 | 6 years ago | Go client for the textbelt.com txt messaging API. |
| [coinpaprika-api-go-client](https://github.com/coinpaprika/coinpaprika-api-go-client) | 14 | 8 | 2018/09/25 | 1 year ago | Go client library for interacting with Coinpaprika's API. |
| [go-hacknews](https://github.com/PaulRosset/go-hacknews) | 14 | 2 | 2017/08/10 | 4 years ago | Tiny Go client for HackerNews API. |
| [go-aws-news](https://github.com/circa10a/go-aws-news) | 12 | 3 | 2020/01/08 | 2 months ago | Go application and library to fetch what's new from AWS. |
| [go-google-analytics](https://github.com/chonthu/go-google-analytics) | 12 | 3 | 2015/06/01 | 6 years ago | Simple wrapper for easy google analytics reporting. |
| [jokeapi](https://github.com/Icelain/jokeapi) | 12 | 2 | 2020/11/22 | 4 weeks ago | Go client for [JokeAPI](https://sv443.net/jokeapi/v2/). |
| [gopaapi5](https://github.com/utekaravinash/gopaapi5) | 11 | 3 | 2020/02/15 | 1 year ago | Go Client Library for [Amazon Product Advertising API 5.0](https://webservices.amazon.com/paapi5/documentation/). |
| [device-check-go](https://github.com/rinchsan/device-check-go) | 10 | 1 | 2019/04/11 | 4 months ago | Go client library for interacting with [iOS DeviceCheck API](https://developer.apple.com/documentation/devicecheck) v1. |
| [go-here](https://github.com/abdullahselek/go-here) | 10 | 0 | 2019/07/07 | 1 year ago | Go client library around the HERE location based APIs. |
| [go-openproject](https://github.com/manuelbcd/go-openproject) | 10 | 2 | 2021/02/13 | 10 months ago | Go client library for interacting with [OpenProject](https://docs.openproject.org/api/) API. |
| [smitego](https://github.com/sergiotapia/smitego) | 10 | 0 | 2013/12/11 | 7 years ago | Go package to wraps access to the Smite game API. |
| [bqwriter](https://github.com/OTA-Insight/bqwriter) | 8 | 2 | 2021/10/12 | 2 weeks ago | High Level Go Library to write data into [Google BigQuery](https://cloud.google.com/bigquery) at a high throughout. |
| [go-sophos](https://github.com/esurdam/go-sophos) | 8 | 2 | 2018/09/05 | 1 year ago | Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies. |
| [rrdaclient](https://github.com/Omie/rrdaclient) | 8 | 2 | 2014/09/15 | 7 years ago | Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP. |
| [go-google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) | 7 | 0 | 2016/10/24 | 5 years ago | Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/). |
| [libgoffi](https://github.com/clevabit/libgoffi) | 7 | 12 | 2019/08/03 | 1 year ago | Library adapter toolbox for native [libffi](https://sourceware. |
| [go-zooz](https://github.com/gojuno/go-zooz) | 7 | 14 | 2017/07/04 | 1 month ago | Go client for the Zooz API. |
| [go-sptrans](https://github.com/sergioaugrod/go-sptrans) | 6 | 3 | 2017/09/11 | 1 year ago | Go client library for the SPTrans Olho Vivo API. |
| [gumblr](https://github.com/mattcunningham/gumblr) | 6 | 1 | 2015/07/09 | 5 years ago | Go wrapper for the Tumblr v2 API. |
| [go-swagger-ui](https://github.com/esurdam/go-swagger-ui) | 5 | 1 | 2021/05/25 | 8 months ago | Go library containing precompiled [Swagger UI](https://swagger.io/tools/swagger-ui/) for serving swagger json. |
| [go-chronos](https://github.com/axelspringer/go-chronos) | 4 | 9 | 2017/10/23 | 4 years ago | Go library for interacting with the [Chronos](https://mesos.github. |
| [rawg-sdk-go](https://github.com/dimuska139/rawg-sdk-go) | 4 | 1 | 2020/10/16 | 1 month ago | Go library for the [RAWG Video Games Database](https://rawg. |
| [kanka](https://github.com/Henry-Sarabia/kanka) | 3 | 3 | 2019/12/26 | 1 year ago | Go client for the [Kanka API](https://kanka.io/en-US/docs/1.0). |
| [appstore-sdk-go](https://github.com/Kachit/appstore-sdk-go) | 2 | 1 | 2020/06/11 | 11 months ago | Unofficial Golang SDK for AppStore Connect API. |
| [go-restcountries](https://github.com/chriscross0/go-restcountries) | 2 | 1 | 2021/08/01 | 3 months ago | Go library for the [REST Countries API](https://restcountries.eu/). |
| [playlyfe-go-sdk](https://github.com/playlyfe/playlyfe-go-sdk) | 1 | 5 | 2015/05/25 | 6 years ago | The Playlyfe Rest API Go SDK. |
| [tripadvisor-golang](https://github.com/mrbenosborne/tripadvisor-golang) | 1 | 2 | 2019/04/15 | 2 years ago | Go wrapper for the TripAdvisor API. |
| [vl-go](https://github.com/verifid/vl-go) | 1 | 2 | 2019/02/09 | 8 months ago | Go client library around the VerifID identity verification layer API. |
**[⬆ back to top](#awesome-go-info)**

# Utilities
        
*General utilities and tools to make your life easier.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [fzf](https://github.com/junegunn/fzf) | 41891 | 395 | 2013/10/23 | 1 week ago | Command-line fuzzy finder written in Go. |
| [hub](https://github.com/github/hub) | 21510 | 518 | 2009/12/05 | 1 month ago | wrap git commands with additional functionality to interact with github from the terminal. |
| [ctop](https://github.com/bcicen/ctop) | 12320 | 163 | 2016/12/27 | 2 weeks ago | [Top-like](https://ctop.sh) interface (e.g. htop) for container metrics. |
| [sqlx](https://github.com/jmoiron/sqlx) | 11400 | 191 | 2013/01/28 | 2 weeks ago | provides a set of extensions on top of the excellent built-in database/sql package. |
| [wuzz](https://github.com/asciimoo/wuzz) | 9890 | 168 | 2017/01/30 | 1 month ago | Interactive cli tool for HTTP inspection. |
| [goreleaser](https://github.com/goreleaser/goreleaser) | 9575 | 109 | 2016/12/21 | 1 week ago | Deliver Go binaries as fast and easily as possible. |
| [usql](https://github.com/xo/usql) | 6936 | 117 | 2017/03/02 | 4 weeks ago | usql is a universal command-line interface for SQL databases. |
| [peco](https://github.com/peco/peco) | 6734 | 136 | 2014/06/06 | 6 months ago | Simplistic interactive filtering tool. |
| [godropbox](https://github.com/dropbox/godropbox) | 4006 | 245 | 2014/06/22 | 1 year ago | Common libraries for writing Go services/applications from Dropbox. |
| [hystrix-go](https://github.com/afex/hystrix-go) | 3527 | 92 | 2013/12/15 | 5 months ago | Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker. |
| [go-funk](https://github.com/thoas/go-funk) | 3213 | 40 | 2016/12/30 | 1 month ago | Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...). |
| [panicparse](https://github.com/maruel/panicparse) | 3016 | 37 | 2015/02/02 | 1 month ago | Groups similar goroutines and colorizes stack dump. |
| [goreporter](https://github.com/qax-os/goreporter) | 2965 | 100 | 2017/03/27 | 3 years ago | Golang tool that does static analysis, unit testing, code review and generate code quality report. |
| [minify](https://github.com/tdewolff/minify) | 2862 | 54 | 2014/05/21 | 2 weeks ago | Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats. |
| [mc](https://github.com/minio/mc) | 2022 | 52 | 2015/01/16 | 1 week ago | Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems. |
| [mergo](https://github.com/imdario/mergo) | 1858 | 25 | 2013/03/11 | 3 weeks ago | Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements. |
| [storm](https://github.com/asdine/storm) | 1845 | 42 | 2016/01/10 | 9 months ago | Simple and powerful toolkit for BoltDB. |
| [mole](https://github.com/davrodpin/mole) | 1538 | 29 | 2018/10/04 | 3 weeks ago | cli app to easily create ssh tunnels. |
| [filetype](https://github.com/h2non/filetype) | 1481 | 28 | 2015/09/24 | 1 month ago | Small package to infer the file type checking the magic numbers signature. |
| [boilr](https://github.com/tmrts/boilr) | 1426 | 28 | 2015/12/19 | 2 months ago | Blazingly fast CLI tool for creating projects from boilerplate templates. |
| [jump](https://github.com/gsamokovarov/jump) | 1297 | 20 | 2015/08/16 | 2 months ago | Jump helps you navigate faster by learning your habits. |
| [cli](https://github.com/create-go-app/cli) | 1264 | 19 | 2019/12/30 | 3 weeks ago | A powerful CLI for create a new production-ready project with backend (Golang), frontend (JavaScript, TypeScript) & deploy automation (Ansible, Docker) by running one command. |
| [circuitbreaker](https://github.com/rubyist/circuitbreaker) | 975 | 21 | 2014/07/17 | 2 years ago | Circuit Breakers in Go. |
| [gtm](https://github.com/git-time-metric/gtm) | 891 | 29 | 2016/06/19 | 2 weeks ago | Simple, seamless, lightweight time tracking for Git. |
| [immortal](https://github.com/immortal/immortal) | 723 | 19 | 2016/06/30 | 1 year ago | \*nix cross-platform (OS agnostic) supervisor. |
| [hostctl](https://github.com/guumaster/hostctl) | 709 | 9 | 2020/03/14 | 2 months ago | A CLI tool to manage /etc/hosts with easy commands. |
| [mimetype](https://github.com/gabriel-vasile/mimetype) | 662 | 10 | 2018/07/02 | 1 week ago | Package for MIME type detection based on magic numbers. |
| [circuit](https://github.com/cep21/circuit) | 625 | 14 | 2017/12/23 | 2 months ago | An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern. |
| [htcat](https://github.com/htcat/htcat) | 545 | 18 | 2013/08/05 | 3 years ago | Parallel and Pipelined HTTP GET Utility. |
| [ergo](https://github.com/cristianoliveira/ergo) | 495 | 5 | 2017/08/19 | 2 months ago | The management of multiple local services running over different ports made easy. |
| [scany](https://github.com/georgysavva/scany) | 489 | 7 | 2020/07/02 | 1 week ago | Library for scanning data from a database into Go structs and more. |
| [godaemon](https://github.com/VividCortex/godaemon) | 486 | 28 | 2013/08/01 | 7 months ago | Utility to write daemons. |
| [koazee](https://github.com/wesovilabs/koazee) | 484 | 11 | 2018/11/09 | 1 year ago | Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays. |
| [go-dry](https://github.com/ungerik/go-dry) | 476 | 14 | 2014/02/28 | 1 month ago | DRY (don't repeat yourself) package for Go. |
| [delve](https://github.com/derekparker/delve) | 451 | 10 | 2020/02/18 | 1 week ago | Go debugger. |
| [gopencils](https://github.com/bndr/gopencils) | 437 | 14 | 2014/06/23 | 3 years ago | Small and simple package to easily consume REST APIs. |
| [request](https://github.com/mozillazg/request) | 407 | 15 | 2014/12/21 | 2 years ago | Go HTTP Requests for Humans™. |
| [gubrak](https://github.com/novalagung/gubrak) | 384 | 6 | 2018/03/09 | 1 year ago | Golang utility library with syntactic sugar. It's like lodash, but for golang. |
| [deepcopier](https://github.com/ulule/deepcopier) | 378 | 20 | 2015/07/24 | 1 year ago | Simple struct copying for Go. |
| [clockwork](https://github.com/jonboulle/clockwork) | 370 | 6 | 2014/09/09 | 2 months ago | A simple fake clock for golang. |
| [go-rate](https://github.com/beefsack/go-rate) | 347 | 11 | 2014/08/25 | 1 year ago | Timed rate limiter for Go. |
| [retry](https://github.com/kamilsk/retry) | 313 | 6 | 2016/11/02 | 11 months ago | The most advanced functional mechanism to perform actions repetitively until successful. |
| [serve](https://github.com/syntaqx/serve) | 259 | 6 | 2019/01/10 | 1 month ago | A static http server anywhere you need. |
| [gohper](https://github.com/cosiner/gohper) | 255 | 21 | 2015/03/23 | 4 years ago | Various tools/modules help for development. |
| [scan](https://github.com/blockloop/scan) | 248 | 6 | 2017/11/27 | 3 months ago | Scan golang `sql.Rows` directly to structs, slices, or primitive types. |
| [util](https://github.com/shomali11/util) | 240 | 12 | 2017/05/24 | 1 year ago | Collection of useful utility functions. (strings, concurrency, manipulations, ...). |
| [go-trigger](https://github.com/sadlil/go-trigger) | 225 | 14 | 2015/10/19 | 4 years ago | Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project. |
| [gotenv](https://github.com/subosito/gotenv) | 214 | 3 | 2013/08/27 | 4 months ago | Load environment variables from `.env` or any `io.Reader` in Go. |
| [grofer](https://github.com/pesos/grofer) | 190 | 8 | 2020/08/01 | 1 month ago | A system and resource monitoring tool written in Golang! |
| [death](https://github.com/vrecan/death) | 176 | 4 | 2015/03/09 | 1 week ago | Managing go application shutdown with signals. |
| [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) | 176 | 5 | 2016/11/08 | 2 years ago | go:generate tool for wrapping symbols exported by golang plugins (1.8 only). |
| [mani](https://github.com/alajmo/mani) | 174 | 5 | 2019/10/22 | 1 month ago | CLI tool to help you manage multiple repositories. |
| [toolbox](https://github.com/viant/toolbox) | 171 | 15 | 2016/06/13 | 6 months ago | Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer. |
| [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) | 170 | 6 | 2015/10/12 | 1 month ago | XML Sitemap generator written in Go. |
| [go-pattern-match](https://github.com/alexpantyukhin/go-pattern-match) | 168 | 2 | 2018/12/11 | 1 year ago | Pattern matching libray. |
| [rerun](https://github.com/ivpusic/rerun) | 162 | 7 | 2014/12/10 | 3 years ago | Recompiling and rerunning go apps when source changes. |
| [moldova](https://github.com/StabbyCutyou/moldova) | 160 | 5 | 2016/01/30 | 4 years ago | Utility for generating random data based on an input template. |
| [clipboard](https://github.com/golang-design/clipboard) | 158 | 4 | 2020/11/19 | 2 weeks ago | 📋 cross-platform clipboard package in Go. |
| [apm](https://github.com/topfreegames/apm) | 154 | 19 | 2015/11/18 | 5 years ago | Process manager for Golang applications with an HTTP API. |
| [robustly](https://github.com/VividCortex/robustly) | 151 | 16 | 2013/07/08 | 9 months ago | Runs functions resiliently, catching and restarting panics. |
| [changie](https://github.com/miniscruff/changie) | 150 | 3 | 2020/12/05 | 1 week ago | Automated changelog tool for preparing releases with lots of customization options. |
| [countries](https://github.com/biter777/countries) | 143 | 6 | 2019/04/22 | 2 months ago | Full implementation of ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and IANA ccTLD standarts. |
| [chyle](https://github.com/antham/chyle) | 141 | 7 | 2016/11/17 | 1 week ago | Changelog generator using a git repository with multiple configuration possibilities. |
| [onecache](https://github.com/adelowo/onecache) | 125 | 7 | 2017/04/14 | 1 year ago | Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc). |
| [go-bsdiff](https://github.com/gabstv/go-bsdiff) | 121 | 3 | 2019/02/23 | 2 years ago | Pure Go bsdiff and bspatch libraries and CLI tools. |
| [lrserver](https://github.com/jaschaephraim/lrserver) | 120 | 5 | 2014/07/15 | 4 years ago | LiveReload server for Go. |
| [nostromo](https://github.com/pokanop/nostromo) | 107 | 5 | 2019/07/13 | 1 week ago | CLI for building powerful aliases. |
| [rospo](https://github.com/ferama/rospo) | 102 | 4 | 2021/04/02 | 4 weeks ago | Simple and reliable ssh tunnels with embedded ssh server in Golang. |
| [cryptgo](https://github.com/Gituser143/cryptgo) | 101 | 2 | 2021/05/20 | 4 months ago | Crytpgo is a TUI based application written purely in Go to monitor and observe cryptocurrency prices in real time! |
| [sorty](https://github.com/jfcg/sorty) | 94 | 4 | 2019/02/18 | 1 month ago | Fast Concurrent / Parallel Sorting. |
| [goseaweedfs](https://github.com/linxGnu/goseaweedfs) | 90 | 8 | 2017/07/20 | 3 months ago | SeaweedFS client library with almost full features. |
| [xferspdy](https://github.com/monmohan/xferspdy) | 90 | 4 | 2015/05/22 | 10 months ago | Xferspdy provides binary diff and patch library in golang. |
| [mssqlx](https://github.com/linxGnu/mssqlx) | 89 | 9 | 2016/12/26 | 5 months ago | Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind. |
| [mongo-go-pagination](https://github.com/gobeam/mongo-go-pagination) | 88 | 3 | 2020/02/04 | 3 months ago | Mongodb Pagination for official mongodb/mongo-go-driver package which supports  both normal queries and Aggregation pipelines. |
| [go-health](https://github.com/Talento90/go-health) | 85 | 3 | 2018/02/13 | 3 weeks ago | Health package simplifies the way you add health check to your services. |
| [cmd](https://github.com/commander-cli/cmd) | 82 | 6 | 2019/09/27 | 1 month ago | Library for executing shell commands on osx, windows and linux. |
| [limiters](https://github.com/mennanov/limiters) | 80 | 3 | 2019/08/28 | 1 month ago | Rate limiters for distributed applications in Golang with configurable back-ends and distributed locks. |
| [pm](https://github.com/VividCortex/pm) | 77 | 17 | 2013/11/17 | 1 year ago | Process (i.e. goroutine) manager with an HTTP API. |
| [repeat](https://github.com/ssgreg/repeat) | 77 | 6 | 2017/11/22 | 1 year ago | Go implementation of different backoff strategies useful for retrying operations and heartbeating. |
| [mimemagic](https://github.com/zRedShift/mimemagic) | 69 | 2 | 2018/10/11 | 2 months ago | Pure Go ultra performant MIME sniffing library/utility. |
| [netbug](https://github.com/e-dard/netbug) | 69 | 2 | 2015/03/05 | 6 years ago | Easy remote profiling of your services. |
| [unis](https://github.com/esemplastic/unis) | 67 | 5 | 2017/05/06 | 4 years ago | Common Architecture™ for String Utilities in Go. |
| [handy](https://github.com/miguelpragier/handy) | 66 | 8 | 2018/06/13 | 1 year ago | Many utilities and helpers like string handlers/formatters and validators. |
| [multitick](https://github.com/VividCortex/multitick) | 65 | 15 | 2013/12/10 | 9 months ago | Multiplexor for aligned tickers. |
| [goval](https://github.com/maja42/goval) | 63 | 5 | 2018/06/17 | 1 year ago | Evaluate arbitrary expressions in Go. |
| [pgo](https://github.com/arthurkushman/pgo) | 63 | 6 | 2018/12/26 | 3 weeks ago | Convenient functions for PHP community. |
| [goreadability](https://github.com/philipjkim/goreadability) | 62 | 6 | 2016/04/20 | 2 years ago | Webpage summary extractor using Facebook Open Graph and arc90's readability. |
| [minquery](https://github.com/icza/minquery) | 59 | 2 | 2016/11/16 | 6 months ago | MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off). |
| [go-astitodo](https://github.com/asticode/go-astitodo) | 58 | 2 | 2016/10/17 | 1 year ago | Parse TODOs in your GO code. |
| [go-lock](https://github.com/viney-shih/go-lock) | 56 | 1 | 2020/04/30 | 6 months ago | go-lock is a lock library implementing read-write mutex and read-write trylock without starvation. |
| [golog](https://github.com/mlimaloureiro/golog) | 56 | 3 | 2016/01/09 | 3 years ago | Easy and lightweight CLI tool to time track your tasks. |
| [filter](https://github.com/gookit/filter) | 53 | 6 | 2018/09/26 | 3 weeks ago | provide filtering, sanitizing, and conversion of Go data. |
| [retry](https://github.com/thedevsaddam/retry) | 53 | 1 | 2018/02/25 | 1 month ago | Simple and easy retry mechanism package for Go. |
| [copy-pasta](https://github.com/jutkko/copy-pasta) | 49 | 5 | 2017/01/28 | 1 year ago | Universal multi-workstation clipboard that uses S3 like backend for the storage. |
| [beyond](https://github.com/wesovilabs/beyond) | 47 | 1 | 2019/10/18 | 4 months ago | The Go tool that will drive you to the AOP world! |
| [slice](https://github.com/psampaz/slice) | 47 | 1 | 2019/11/26 | 1 year ago | Type-safe functions for common Go slice operations. |
| [golarm](https://github.com/msempere/golarm) | 46 | 2 | 2015/08/14 | 6 years ago | Fire alarms with system events. |
| [dbt](https://github.com/nikogura/dbt) | 45 | 1 | 2017/11/30 | 11 months ago | A framework for running self-updating signed binaries from a central, trusted repository. |
| [goback](https://github.com/carlescere/goback) | 44 | 1 | 2015/03/13 | 11 months ago | Go simple exponential backoff package. |
| [intrinsic](https://github.com/mengzhuo/intrinsic) | 43 | 4 | 2017/06/13 | 4 years ago | Use x86 SIMD without writing any assembly code. |
| [retry-go](https://github.com/rafaeljesus/retry-go) | 43 | 2 | 2017/06/09 | 3 years ago | Retrying made simple and easy for golang. |
| [gpath](https://github.com/tenntenn/gpath) | 40 | 3 | 2017/05/24 | 4 years ago | Library to simplify access struct fields with Go's expression in reflection. |
| [go-httpheader](https://github.com/mozillazg/go-httpheader) | 37 | 3 | 2017/06/24 | 2 months ago | Go library for encoding structs into Header fields. |
| [myhttp](https://github.com/inancgumus/myhttp) | 35 | 0 | 2017/09/13 | 3 years ago | Simple API to make HTTP GET requests with timeout support. |
| [backscanner](https://github.com/icza/backscanner) | 34 | 4 | 2017/10/18 | 4 months ago | A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward. |
| [equalizer](https://github.com/reugn/equalizer) | 34 | 2 | 2019/06/14 | 1 year ago | Quota manager and rate limiter collection for Go. |
| [evaluator](https://github.com/nullne/evaluator) | 33 | 3 | 2017/04/27 | 6 months ago | Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend. |
| [gostrutils](https://github.com/ik5/gostrutils) | 32 | 3 | 2018/09/19 | 5 months ago | Collections of string manipulation and conversion functions. |
| [rclient](https://github.com/zpatrick/rclient) | 32 | 3 | 2017/02/28 | 2 years ago | Readable, flexible, simple-to-use client for REST APIs. |
| [slicer](https://github.com/leaanthony/slicer) | 30 | 2 | 2019/01/10 | 6 months ago | Makes working with slices easier. |
| [set](https://github.com/nofeaturesonlybugs/set) | 29 | 4 | 2020/12/16 | 2 weeks ago | Performant and flexible struct mapping and loose type conversion. |
| [tome](https://github.com/cyruzin/tome) | 29 | 1 | 2019/04/12 | 1 year ago | Tome was designed to paginate simple RESTful APIs. |
| [throttle](https://github.com/yudppp/throttle) | 28 | 2 | 2019/10/25 | 5 months ago | Throttle is an object that will perform exactly one action per duration. |
| [sshman](https://github.com/shoobyban/sshman) | 27 | 2 | 2021/08/27 | 2 weeks ago | SSH Manager for authorized_keys files on multiple remote servers. |
| [shutdown](https://github.com/ztrue/shutdown) | 26 | 1 | 2018/11/17 | 1 month ago | App shutdown hooks for `os.Signal` handling. |
| [ugo](https://github.com/alxrm/ugo) | 26 | 2 | 2016/02/17 | 5 years ago | ugo is slice toolbox with concise syntax for Go. |
| [generate](https://github.com/go-playground/generate) | 24 | 3 | 2015/11/15 | 5 years ago | runs go generate recursively on a specified path or environment variable and can filter by regex. |
| [ghokin](https://github.com/antham/ghokin) | 23 | 3 | 2018/08/03 | 1 month ago | Parallelized formatter with no external dependencies for gherkin (cucumber, behat...). |
| [goplaceholder](https://github.com/michiwend/goplaceholder) | 22 | 2 | 2014/10/12 | 6 years ago | a small golang lib to generate placeholder images. |
| [rerate](https://github.com/abo/rerate) | 21 | 4 | 2016/05/24 | 4 years ago | Redis-based rate counter and rate limiter for Go. |
| [mimesniffer](https://github.com/aofei/mimesniffer) | 19 | 1 | 2018/12/20 | 3 months ago | A MIME type sniffer for Go. |
| [ctxutil](https://github.com/posener/ctxutil) | 18 | 1 | 2018/07/30 | 1 year ago | A collection of utility functions for contexts. |
| [structs](https://github.com/PumpkinSeed/structs) | 18 | 3 | 2017/08/26 | 4 years ago | Implement simple functions to manipulate structs. |
| [filler](https://github.com/yaronsumel/filler) | 16 | 1 | 2017/04/05 | 4 years ago | small utility to fill structs using "fill" tag. |
| [rest-go](https://github.com/edermanoel94/rest-go) | 16 | 3 | 2019/07/29 | 1 year ago | A package that provide many helpful methods for working with rest api. |
| [dlog](https://github.com/kirillDanshin/dlog) | 15 | 2 | 2016/07/04 | 4 years ago | Compile-time controlled logger to make your release smaller without removing debug calls. |
| [go-convert](https://github.com/Eun/go-convert) | 15 | 1 | 2019/06/07 | 9 months ago | Package go-convert enables you to convert a value into another type. |
| [okrun](https://github.com/xta/okrun) | 15 | 3 | 2014/10/01 | 7 years ago | go run error steamroller. |
| [copy](https://github.com/gotidy/copy) | 14 | 3 | 2020/10/09 | 1 year ago | Package for fast copying structs of different types. |
| [jsend](https://github.com/clevergo/jsend) | 14 | 4 | 2020/01/14 | 7 months ago | JSend's implementation writen in Go. |
| [command](https://github.com/txgruppi/command) | 13 | 1 | 2015/08/24 | 5 years ago | Command pattern for Go with thread safe serial and parallel dispatcher. |
| [go-types](https://github.com/mikekonan/go-types) | 13 | 1 | 2021/04/21 | 2 months ago | Library providing Go types for store/validation and transfer of ISO-4217, ISO-3166, and other types. |
| [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) | 12 | 1 | 2019/05/16 | 2 years ago | Go package for working with Problem Details. |
| [cvt](https://github.com/shockerli/cvt) | 12 | 1 | 2021/03/09 | 1 month ago | Easy and safe convert any value to another type. |
| [ptr](https://github.com/gotidy/ptr) | 11 | 2 | 2019/12/25 | 1 month ago | Package that provide functions for simplified creation of pointers from constants of basic types. |
| [silk](https://github.com/chrispassas/silk) | 11 | 1 | 2018/12/18 | 1 year ago | Read silk netflow files. |
| [retry](https://github.com/shafreeck/retry) | 10 | 0 | 2018/07/18 | 2 years ago | A pretty simple library to ensure your work to be done. |
| [go-countries](https://github.com/mikekonan/go-countries) | 9 | 3 | 2020/10/27 | 1 year ago | Lightweight lookup over ISO-3166 codes. |
| [statiks](https://github.com/janiltonmaciel/statiks) | 9 | 0 | 2018/06/26 | 1 year ago | Fast, zero-configuration, static HTTP filer server. |
| [sliceconv](https://github.com/Henry-Sarabia/sliceconv) | 8 | 1 | 2019/02/15 | 2 years ago | Slice conversion between primitive types. |
| [blank](https://github.com/Henry-Sarabia/blank) | 7 | 2 | 2019/02/13 | 2 years ago | Verify or remove blanks and whitespace from strings. |
| [go-clip](https://github.com/prashantgupta24/go-clip) | 7 | 2 | 2020/11/18 | 1 year ago | A minimalistic clipboard manager for Mac. |
| [retry](https://github.com/percolate/retry) | 7 | 41 | 2018/06/15 | 2 years ago | A simple but highly configurable retry package for Go. |
| [nfdump](https://github.com/chrispassas/nfdump) | 6 | 1 | 2020/04/08 | 6 months ago | Read nfdump netflow files. |
| [go-actuator](https://github.com/sinhashubham95/go-actuator) | 5 | 1 | 2021/07/17 | 5 months ago | Production ready features for Go based web frameworks. |
| [go-pkg](https://github.com/chenquan/go-pkg) | 5 | 1 | 2021/11/28 | 1 week ago | A go toolkit. |
| [go-safe](https://github.com/kenkyu392/go-safe) | 4 | 0 | 2019/10/29 | 2 months ago | Panic-safe sandbox. |
| [lets-go](https://github.com/aplescia/lets-go) | 3 | 1 | 2020/02/19 | 9 months ago | Go module that provides common utilities for Cloud Native REST API development. Also contains AWS Specific utilities. |
| [olaf](https://github.com/btnguyen2k/olaf) | 3 | 1 | 2019/01/03 | 2 years ago | Twitter Snowflake implemented in Go. |
| [tik](https://github.com/andy2046/tik) | 3 | 1 | 2020/07/04 | 1 year ago | Simple and easy timing wheel package for Go. |
| [bleep](https://github.com/sinhashubham95/bleep) | 2 | 1 | 2021/01/02 | 1 year ago | Perform any number of actions on any set of OS signals in Go. |
| [goctx](https://github.com/zerosnake0/goctx) | 2 | 1 | 2020/11/14 | 1 year ago | Get your context value with high performance. |
**[⬆ back to top](#awesome-go-info)**

# UUID
        
*Libraries for working with UUIDs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [uuid](https://github.com/google/uuid) | 3353 | 46 | 2016/02/12 | 2 months ago | Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services. |
| [ulid](https://github.com/oklog/ulid) | 2534 | 46 | 2016/12/06 | 3 months ago | Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier). |
| [xid](https://github.com/rs/xid) | 2396 | 33 | 2015/11/10 | 1 month ago | Xid is a globally unique id generator library, ready to be safely used directly in your server code. |
| [uuid](https://github.com/gofrs/uuid) | 1062 | 19 | 2018/07/13 | 2 months ago | Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid. |
| [wuid](https://github.com/edwingeng/wuid) | 447 | 17 | 2018/01/27 | 6 months ago | An extremely fast unique number generator, 10-135 times faster than UUID. |
| [sno](https://github.com/muyo/sno) | 56 | 3 | 2019/05/26 | 3 months ago | Compact, sortable and fast unique IDs with embedded metadata. |
| [nanoid](https://github.com/aidarkhanov/nanoid) | 44 | 2 | 2019/07/02 | 5 months ago | A tiny and efficient Go unique string ID generator. |
| [Goid](https://github.com/JakeHL/Goid) | 30 | 5 | 2017/05/19 | 3 years ago | Generate and Parse RFC4122 compliant V4 UUIDs. |
| [uuid](https://github.com/agext/uuid) | 14 | 3 | 2016/02/03 | 1 year ago | Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier. |
| [gouid](https://github.com/twharmon/gouid) | 12 | 1 | 2020/10/08 | 2 weeks ago | Generate cryptographically secure random string IDs with just one allocation. |
| [goflake](https://github.com/Hart87/goflake) | 9 | 1 | 2021/05/03 | 9 months ago | A small, scalable, & serverless unique ID generator for use in distributed systems. Inspired by Twitters Snowflake. |
**[⬆ back to top](#awesome-go-info)**

# Validation
        
*Libraries for validation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [validator](https://github.com/go-playground/validator) | 9568 | 101 | 2015/02/12 | 2 weeks ago | Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving. |
| [govalidator](https://github.com/asaskevich/govalidator) | 5230 | 95 | 2014/06/20 | 1 week ago | Validators and sanitizers for strings, numerics, slices and structs. |
| [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) | 2542 | 29 | 2016/06/22 | 3 weeks ago | Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags. |
| [govalidator](https://github.com/thedevsaddam/govalidator) | 1054 | 22 | 2017/09/13 | 2 weeks ago | Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. |
| [validate](https://github.com/gookit/validate) | 517 | 17 | 2018/07/16 | 3 weeks ago | Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features. |
| [checkdigit](https://github.com/osamingo/checkdigit) | 88 | 0 | 2019/04/05 | 1 year ago | Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.). |
| [terraform-validator](https://github.com/thazelart/terraform-validator) | 76 | 3 | 2019/05/29 | 1 year ago | A norms and conventions validator for Terraform. |
| [validate](https://github.com/gobuffalo/validate) | 66 | 10 | 2018/02/10 | 3 months ago | This package provides a framework for writing validations for Go applications. |
| [jio](https://github.com/faceair/jio) | 64 | 2 | 2018/10/28 | 1 year ago | jio is a json schema validator similar to [joi](https://github.com/hapijs/joi). |
| [gody](https://github.com/guiferpa/gody) | 53 | 0 | 2018/11/01 | 1 year ago | :balloon: A lightweight struct validator for Go. |
| [govalid](https://github.com/twharmon/govalid) | 24 | 1 | 2019/02/17 | 4 months ago | Fast, tag-based validation for structs. |
**[⬆ back to top](#awesome-go-info)**

# Version Control
        
*Libraries for version control.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-git](https://github.com/go-git/go-git) | 3072 | 37 | 2019/12/19 | 1 week ago | highly extensible Git implementation in pure Go. |
| [glab](https://github.com/profclems/glab) | 1716 | 20 | 2020/07/24 | 2 weeks ago | An open-source GitLab command line tool bringing GitLab's cool features to your command line. |
| [git2go](https://github.com/libgit2/git2go) | 1691 | 50 | 2013/03/05 | 2 weeks ago | Go bindings for libgit2. |
| [hercules](https://github.com/src-d/hercules) | 1664 | 22 | 2016/12/12 | 3 months ago | gaining advanced insights from Git repository history. |
| [gh](https://github.com/rjeczalik/gh) | 76 | 6 | 2015/03/08 | 3 years ago | Scriptable server and net/http middleware for GitHub Webhooks. |
| [go-vcs](https://github.com/sourcegraph/go-vcs) | 74 | 80 | 2013/06/02 | 10 months ago | manipulate and inspect VCS repositories in Go. |
| [Githooks](https://github.com/gabyx/Githooks) | 32 | 0 | 2019/06/28 | 4 weeks ago | Per-repo and shared Git hooks with version control and auto update. |
| [hgo](https://github.com/beyang/hgo) | 13 | 4 | 2014/06/18 | 6 years ago | Hgo is a collection of Go packages providing read-access to local Mercurial repositories. |
| [froggit-go](https://github.com/jfrog/froggit-go) | 13 | 6 | 2021/08/31 | 2 weeks ago | Froggit-Go is a Go library, allowing to perform actions on VCS providers. |
**[⬆ back to top](#awesome-go-info)**

# Video
        
*Libraries for manipulating video.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [goav](https://github.com/giorgisio/goav) | 1803 | 48 | 2015/05/21 | 8 months ago | Comprehensive Go bindings for FFmpeg. |
| [m3u8](https://github.com/grafov/m3u8) | 895 | 40 | 2013/02/05 | 2 weeks ago | Parser and generator library of M3U8 playlists for Apple HLS. |
| [gmf](https://github.com/3d0c/gmf) | 740 | 31 | 2013/04/03 | 3 weeks ago | Go bindings for FFmpeg av\* libraries. |
| [go-astits](https://github.com/asticode/go-astits) | 390 | 15 | 2017/07/04 | 3 weeks ago | Parse and demux MPEG Transport Streams (.ts) natively in GO. |
| [go-astisub](https://github.com/asticode/go-astisub) | 368 | 7 | 2016/12/16 | 2 weeks ago | Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.). |
| [libvlc-go](https://github.com/adrg/libvlc-go) | 274 | 12 | 2015/01/06 | 4 months ago | Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player). |
| [gortsplib](https://github.com/aler9/gortsplib) | 185 | 8 | 2020/01/20 | 1 week ago | Pure Go RTSP server and client library. |
| [gst](https://github.com/ziutek/gst) | 166 | 10 | 2011/07/26 | 1 year ago | Go bindings for GStreamer. |
| [go-m3u8](https://github.com/quangngotan95/go-m3u8) | 87 | 2 | 2018/11/06 | 1 year ago | Parser and generator library for Apple m3u8 playlists. |
| [v4l](https://github.com/korandiz/v4l) | 66 | 7 | 2016/10/25 | 1 month ago | Video capture library for Linux, written in Go. |
| [libgosubs](https://github.com/wargarblgarbl/libgosubs) | 17 | 2 | 2017/05/03 | 1 year ago | Subtitle format support for go. Supports .srt, .ttml, and .ass. |
| [go-mpd](https://github.com/unki2aut/go-mpd) | 11 | 1 | 2018/11/02 | 1 year ago | Parser and generator library for MPEG-DASH manifest files. |
**[⬆ back to top](#awesome-go-info)**

# Web Frameworks
        
*Full stack web frameworks.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gin](https://github.com/gin-gonic/gin) | 55223 | 1355 | 2014/06/16 | 1 week ago | Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity. |
| [echo](https://github.com/labstack/echo) | 21568 | 530 | 2015/03/01 | 1 week ago | High performance, minimalist Go web framework. |
| [fiber](https://github.com/gofiber/fiber) | 18026 | 248 | 2020/01/16 | 1 week ago | An Express.js inspired web framework build on Fasthttp. |
| [revel](https://github.com/revel/revel) | 12485 | 536 | 2011/12/09 | 3 months ago | High-productivity web framework for the Go language. |
| [goa](https://github.com/goadesign/goa) | 4564 | 164 | 2014/12/05 | 2 weeks ago | Goa provides a holistic approach for developing remote APIs and microservices in Go. |
| [gizmo](https://github.com/nytimes/gizmo) | 3586 | 121 | 2015/12/15 | 6 months ago | Microservice toolkit used by the New York Times. |
| [go-json-rest](https://github.com/ant0ine/go-json-rest) | 3493 | 160 | 2013/02/19 | 1 year ago | Quick and easy way to setup a RESTful JSON API. |
| [macaron](https://github.com/go-macaron/macaron) | 3246 | 145 | 2014/07/10 | 2 weeks ago | Macaron is a high productive and modular design web framework in Go. |
| [utron](https://github.com/gernest/utron) | 2210 | 69 | 2015/09/16 | 3 years ago | Lightweight MVC framework for Go(Golang). |
| [go-tigertonic](https://github.com/rcrowley/go-tigertonic) | 998 | 46 | 2013/02/09 | 3 years ago | Go framework for building JSON web services inspired by Dropwizard. |
| [goyave](https://github.com/go-goyave/goyave) | 945 | 26 | 2019/10/21 | 1 month ago | Feature-complete REST API framework aimed at clean code and fast development, with powerful built-in functionalities. |
| [tango](https://github.com/lunny/tango) | 835 | 75 | 2014/12/17 | 2 years ago | Micro & pluggable web framework for Go. |
| [gearbox](https://github.com/gogearbox/gearbox) | 595 | 17 | 2020/04/25 | 2 weeks ago | A web framework written in Go with a focus on high performance and memory optimization. |
| [gongular](https://github.com/mustafaakin/gongular) | 448 | 21 | 2016/06/22 | 1 year ago | Fast Go web framework with input mapping/validation and (DI) Dependency Injection. |
| [aero](https://github.com/aerogo/aero) | 429 | 21 | 2016/11/09 | 2 months ago | High-performance web framework for Go, reaches top scores in Lighthouse. |
| [neo](https://github.com/ivpusic/neo) | 415 | 34 | 2015/02/04 | 4 years ago | Neo is minimal and fast Go Web Framework with extremely simple API. |
| [air](https://github.com/aofei/air) | 409 | 19 | 2016/07/20 | 10 months ago | An ideally refined web framework for Go. |
| [beego](https://github.com/astaxie/beego) | 366 | 11 | 2020/12/13 | 5 months ago | beego is an open-source, high-performance web framework for the Go programming language. |
| [mango](https://github.com/paulbellamy/mango) | 360 | 21 | 2011/05/25 | 4 years ago | Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. |
| [gondola](https://github.com/rainycape/gondola) | 309 | 15 | 2014/07/25 | 3 years ago | The web framework for writing faster sites, faster. |
| [flamingo](https://github.com/i-love-flamingo/flamingo) | 276 | 28 | 2019/04/02 | 1 week ago | Framework for pluggable web projects. Including a concept for modules and offering features for DI, Configareas, i18n, template engines, graphql, observability, security, events, routing & reverse routing etc. |
| [flamingo-commerce](https://github.com/i-love-flamingo/flamingo-commerce) | 257 | 24 | 2019/04/02 | 2 weeks ago | Providing e-commerce features using clean architecture like DDD and ports and adapters, that you can use to build flexible e-commerce applications. |
| [golf](https://github.com/dinever/golf) | 252 | 16 | 2015/11/18 | 5 months ago | Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library. |
| [webgo](https://github.com/bnkamalesh/webgo) | 217 | 9 | 2015/12/16 | 4 months ago | A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc). |
| [ginrpc](https://github.com/xxjwxc/ginrpc) | 205 | 7 | 2019/06/22 | 2 months ago | Gin parameter automatic binding tool,gin rpc tools. |
| [uadmin](https://github.com/uadmin/uadmin) | 174 | 15 | 2018/10/05 | 3 weeks ago | Fully featured web framework for Golang, inspired by Django. |
| [hiboot](https://github.com/hidevopsio/hiboot) | 163 | 14 | 2018/03/16 | 1 month ago | hiboot is a high performance web application framework with auto configuration and dependency injection support. |
| [rk-boot](https://github.com/rookie-ninja/rk-boot) | 137 | 4 | 2020/07/31 | 1 week ago | A bootstrapper library for building enterprise go microservice with Gin and gRPC quickly and easily. |
| [go-rest](https://github.com/ungerik/go-rest) | 125 | 10 | 2012/07/13 | 5 years ago | Small and evil REST framework for Go. |
| [appy](https://github.com/appist/appy) | 115 | 4 | 2019/05/27 | 2 months ago | An opinionated productive web framework that helps scaling business easier. |
| [patron](https://github.com/beatlabs/patron) | 92 | 16 | 2019/01/30 | 1 week ago | Patron is a microservice framework following best cloud practices with a focus on productivity. |
| [microservice](https://github.com/claygod/microservice) | 89 | 9 | 2016/12/15 | 3 weeks ago | The framework for the creation of microservices, written in Golang. |
| [vox](https://github.com/aisk/vox) | 76 | 2 | 2014/12/24 | 8 months ago | A golang web framework for humans, inspired by Koa heavily. |
| [golax](https://github.com/fulldump/golax) | 73 | 7 | 2016/01/30 | 1 week ago | A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more. |
| [rux](https://github.com/gookit/rux) | 71 | 4 | 2018/08/05 | 3 weeks ago | Simple and fast web framework for build golang HTTP applications. |
| [yarf](https://github.com/yarf-framework/yarf) | 65 | 4 | 2015/09/02 | 2 years ago | Fast micro-framework designed to build REST APIs and web services in a fast and simple way. |
| [fireball](https://github.com/zpatrick/fireball) | 56 | 4 | 2016/07/20 | 3 years ago | More "natural" feeling web framework. |
| [goa](https://github.com/goa-go/goa) | 48 | 4 | 2019/07/26 | 2 years ago | goa is just like koajs for golang, it is a flexible, light, high-performance and extensible web framework based on middleware. |
| [gotuna](https://github.com/gotuna/gotuna) | 39 | 1 | 2021/04/08 | 6 months ago | Minimalistic web framework for Go with mux router, middlewares, sessions, templates, embedded views and static files. |
| [api](https://github.com/resoursea/api) | 32 | 6 | 2015/01/24 | 7 years ago | REST framework for quickly writing resource based services. |
| [rex](https://github.com/goanywhere/rex) | 32 | 2 | 2014/10/16 | 4 years ago | Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`. |
| [goweb](https://github.com/twharmon/goweb) | 26 | 1 | 2019/05/07 | 1 month ago | Web framework with routing, websockets, logging, middleware, static file server (optional gzip), and automatic TLS. |
| [banjo](https://github.com/n4Zz2/banjo) | 19 | 1 | 2017/12/09 | 4 years ago | Very simple and fast web framework for Go. |
**[⬆ back to top](#awesome-go-info)**

## Actual middlewares
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [tollbooth](https://github.com/didip/tollbooth) | 2099 | 48 | 2015/05/17 | 2 months ago | Rate limit HTTP request handler. |
| [cors](https://github.com/rs/cors) | 2012 | 30 | 2014/10/25 | 2 weeks ago | Easily add CORS capabilities to your API. |
| [limiter](https://github.com/ulule/limiter) | 1492 | 28 | 2015/10/02 | 2 weeks ago | Dead simple rate limit middleware for Go. |
| [go-server-timing](https://github.com/mitchellh/go-server-timing) | 826 | 18 | 2018/02/12 | 1 year ago | Add/parse Server-Timing header. |
| [go-fault](https://github.com/github/go-fault) | 421 | 137 | 2020/05/14 | 5 months ago | Fault injection middleware for Go. |
| [ln-paywall](https://github.com/philippgille/ln-paywall) | 120 | 5 | 2018/06/29 | 3 years ago | Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin). |
| [xff](https://github.com/sebest/xff) | 87 | 3 | 2014/12/22 | 3 weeks ago | Handle `X-Forwarded-For` header and friends. |
| [formjson](https://github.com/rs/formjson) | 36 | 2 | 2015/03/19 | 6 years ago | Transparently handle JSON input as a standard form POST. |
| [rk-gin](https://github.com/rookie-ninja/rk-gin) | 24 | 3 | 2020/10/12 | 2 weeks ago | Middleware for Gin framework with logging, metrics, auth, tracing etc. |
| [rk-grpc](https://github.com/rookie-ninja/rk-grpc) | 24 | 2 | 2020/07/25 | 2 weeks ago | Middleware for gRPC with logging, metrics, auth, tracing etc. |
| [client-timing](https://github.com/posener/client-timing) | 20 | 1 | 2018/02/23 | 1 year ago | An HTTP client for Server-Timing header. |
| [mid](https://github.com/bobg/mid) | 4 | 1 | 2020/07/13 | 6 months ago | Miscellaneous HTTP middleware features: idiomatic error return from handlers; receive/respond with JSON data; request tracing; and more. |
**[⬆ back to top](#awesome-go-info)**

## Libraries for creating HTTP middlewares
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [negroni](https://github.com/urfave/negroni) | 7144 | 236 | 2014/05/18 | 1 week ago | Idiomatic HTTP middleware for Golang. |
| [alice](https://github.com/justinas/alice) | 2468 | 48 | 2014/05/25 | 3 months ago | Painless middleware chaining for Go. |
| [render](https://github.com/unrolled/render) | 1573 | 36 | 2014/06/10 | 3 months ago | Go package for easily rendering JSON, XML, and HTML template responses. |
| [stats](https://github.com/thoas/stats) | 583 | 16 | 2015/03/05 | 2 years ago | Go middleware that stores various information about your web application. |
| [interpose](https://github.com/carbocation/interpose) | 296 | 12 | 2014/07/20 | 5 years ago | Minimalist net/http middleware for golang. |
| [renderer](https://github.com/thedevsaddam/renderer) | 232 | 7 | 2017/11/07 | 1 year ago | Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go. |
| [muxchain](https://github.com/stephens2424/muxchain) | 210 | 5 | 2014/05/03 | 2 years ago | Lightweight middleware for net/http. |
| [gores](https://github.com/alioygur/gores) | 97 | 5 | 2015/12/25 | 1 year ago | Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs. |
| [rye](https://github.com/InVisionApp/rye) | 96 | 201 | 2016/10/06 | 3 years ago | Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context. |
| [mediary](https://github.com/HereMobilityDevelopers/mediary) | 78 | 5 | 2020/03/23 | 1 year ago | add interceptors to `http.Client` to allow dumping/shaping/tracing/... of requests/responses. |
| [chain](https://github.com/codemodus/chain) | 64 | 7 | 2015/05/14 | 3 years ago | Handler wrapper chaining with scoped data (net/context-based "middleware"). |
| [wrap](https://github.com/go-on/wrap) | 60 | 3 | 2014/02/16 | 3 years ago | Small middlewares package for net/http. |
| [catena](https://github.com/codemodus/catena) | 8 | 2 | 2015/07/30 | 3 years ago | http.Handler wrapper catenation (same API as "chain"). |
**[⬆ back to top](#awesome-go-info)**

# WebAssembly
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [tinygo](https://github.com/tinygo-org/tinygo) | 9368 | 160 | 2018/06/07 | 1 week ago | Go compiler for small places. Microcontrollers, WebAssembly, and command-line tools. Based on LLVM. |
| [dom](https://github.com/dennwc/dom) | 441 | 17 | 2018/06/30 | 2 years ago | DOM library. |
| [go-canvas](https://github.com/markfarnan/go-canvas) | 156 | 7 | 2019/05/05 | 1 year ago | Library to use HTML5 Canvas, with all drawing within go code. |
| [wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest) | 112 | 3 | 2018/07/14 | 3 months ago | Run Go WASM tests in your browser. |
| [webapi](https://github.com/gowebapi/webapi) | 100 | 8 | 2019/02/08 | 1 month ago | Bindings for DOM and HTML generated from WebIDL. |
| [vert](https://github.com/norunners/vert) | 59 | 6 | 2018/03/25 | 1 month ago | Interop between Go and JS values. |
**[⬆ back to top](#awesome-go-info)**

# Windows
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-ole](https://github.com/go-ole/go-ole) | 835 | 41 | 2011/01/21 | 2 months ago | Win32 OLE implementation for golang. |
| [d3d9](https://github.com/gonutz/d3d9) | 129 | 8 | 2015/12/12 | 2 months ago | Go bindings for Direct3D9. |
| [gosddl](https://github.com/MonaxGT/gosddl) | 8 | 2 | 2018/12/04 | 2 years ago | Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL. |
**[⬆ back to top](#awesome-go-info)**

# XML
        
*Libraries and tools for manipulating XML.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [zek](https://github.com/miku/zek) | 532 | 20 | 2017/11/23 | 4 months ago | Generate a Go struct from XML. |
| [xpath](https://github.com/antchfx/xpath) | 453 | 12 | 2016/10/09 | 1 month ago | XPath package for Go. |
| [xquery](https://github.com/antchfx/xquery) | 155 | 11 | 2016/10/09 | 3 years ago | XQuery lets you extract data from HTML/XML documents using XPath expression. |
| [xml2map](https://github.com/sbabiv/xml2map) | 38 | 3 | 2018/08/06 | 2 months ago | XML to MAP converter written Golang. |
| [xmlwriter](https://github.com/shabbyrobe/xmlwriter) | 21 | 2 | 2017/04/11 | 10 months ago | Procedural XML generation API based on libxml2's xmlwriter module. |
| [XML-Comp](https://github.com/XML-Comp/XML-Comp) | 16 | 2 | 2016/10/25 | 3 years ago | Simple command line XML comparer that generates diffs of folders, files and tags. |
**[⬆ back to top](#awesome-go-info)**

# Zero Trust
        
*Libraries and tools to implement Zero Trust architectures.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [cosign](https://github.com/sigstore/cosign) | 1573 | 39 | 2021/02/04 | 1 week ago | Container Signing, Verification and Storage in an OCI registry. |
| [spire](https://github.com/spiffe/spire) | 1040 | 81 | 2017/08/11 | 1 week ago | SPIRE (the SPIFFE Runtime Environment) is a toolchain of APIs for establishing trust between software systems across a wide variety of hosting platforms. |
| [in-toto-golang](https://github.com/in-toto/in-toto-golang) | 48 | 10 | 2018/10/15 | 1 month ago | Go implementation of the in-toto (provides a framework to protect the integrity of the software supply chain) python reference implementation. |
| [spiffe-vault](https://github.com/philips-labs/spiffe-vault) | 10 | 1 | 2021/08/26 | 2 weeks ago | Utilizes Spiffe JWT authentication with Hashicorp Vault for secretless authentication. |
**[⬆ back to top](#awesome-go-info)**

# Tools
        
*Go software and plugins.*
**[⬆ back to top](#awesome-go-info)**

# Code Analysis
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [lint](https://github.com/golang/lint) | 3929 | 103 | 2013/06/02 | 9 months ago | Golint is a linter for Go source code. |
| [errcheck](https://github.com/kisielk/errcheck) | 1786 | 26 | 2013/02/24 | 1 month ago | Errcheck is a program for checking for unchecked errors in Go programs. |
| [go-critic](https://github.com/go-critic/go-critic) | 1229 | 20 | 2018/05/05 | 1 week ago | source code linter that brings checks that are currently not implemented in other linters. |
| [gcvis](https://github.com/davecheney/gcvis) | 1053 | 35 | 2014/07/10 | 2 years ago | Visualise Go program GC trace data in real time. |
| [php-parser](https://github.com/z7zmey/php-parser) | 851 | 29 | 2017/11/07 | 9 months ago | A Parser for PHP written in Go. |
| [goplantuml](https://github.com/jfeliu007/goplantuml) | 659 | 12 | 2019/05/26 | 2 weeks ago | Library and CLI that generates text plantump class diagram containing information about structures and interfaces with the relationship among them. |
| [goast-viewer](https://github.com/yuroyoro/goast-viewer) | 587 | 16 | 2014/06/30 | 2 years ago | Web based Golang AST visualizer. |
| [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) | 557 | 11 | 2017/04/12 | 3 months ago | go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects. |
| [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) | 553 | 5 | 2019/04/19 | 1 month ago | An easy way to find outdated dependencies of your Go projects. |
| [todocheck](https://github.com/preslavmihaylov/todocheck) | 368 | 5 | 2020/07/18 | 4 months ago | Static code analyser which links TODO comments in code with issues in your issue tracker. |
| [unconvert](https://github.com/mdempsky/unconvert) | 315 | 9 | 2016/02/19 | 1 year ago | Remove unnecessary type conversions from Go source. |
| [golines](https://github.com/segmentio/golines) | 310 | 15 | 2019/10/01 | 1 month ago | Formatter that automatically shortens long lines in Go code. |
| [tickgit](https://github.com/augmentable-dev/tickgit) | 277 | 8 | 2019/10/12 | 1 month ago | CLI and go package for surfacing code comment TODOs (in any language) and applying a `git blame`to identify the author. |
| [dupl](https://github.com/mibk/dupl) | 267 | 8 | 2015/05/20 | 1 year ago | Tool for code clone detection. |
| [gostatus](https://github.com/shurcooL/gostatus) | 245 | 7 | 2013/11/27 | 3 years ago | Command line tool, shows the status of repositories that contain Go packages. |
| [apicompat](https://github.com/bradleyfalzon/apicompat) | 176 | 7 | 2016/07/10 | 5 years ago | Checks recent changes to a Go project for backwards incompatible changes. |
| [checkstyle](https://github.com/qiniu/checkstyle) | 118 | 12 | 2014/01/01 | 11 months ago | checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments. |
| [lint](https://github.com/surullabs/lint) | 66 | 5 | 2016/07/09 | 3 years ago | Run linters as part of go test. |
| [validate](https://github.com/mccoyst/validate) | 60 | 6 | 2013/11/22 | 5 years ago | Automatically validates struct fields with tags. |
| [go-outdated](https://github.com/firstrow/go-outdated) | 44 | 1 | 2015/06/29 | 3 years ago | Console application that displays outdated packages. |
| [chainjacking](https://github.com/Checkmarx/chainjacking) | 18 | 5 | 2021/11/16 | 3 months ago | Find which of your Go lang direct GitHub dependencies is susceptible to ChainJacking attack. |
| [blanket](https://github.com/verygoodsoftwarenotvirus/blanket) | 14 | 3 | 2017/09/04 | 3 years ago | tarp finds functions and methods without direct unit tests in Go source code. |
| [golang-ifood-sdk](https://github.com/arxdsilva/golang-ifood-sdk) | 8 | 1 | 2021/03/13 | 3 months ago | iFood API SDK. |
**[⬆ back to top](#awesome-go-info)**

# Editor Plugins
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [vim-go](https://github.com/fatih/vim-go) | 14229 | 269 | 2014/03/24 | 2 weeks ago | Go development plugin for Vim. |
| [gocode](https://github.com/nsf/gocode) | 4944 | 191 | 2010/07/05 | 3 months ago | Autocompletion daemon for the Go programming language. |
| [GoSublime](https://github.com/DisposaBoy/GoSublime) | 3420 | 117 | 2011/08/27 | 1 year ago | Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features. |
| [vscode-go](https://github.com/golang/vscode-go) | 2390 | 56 | 2020/03/06 | 1 week ago | Extension for Visual Studio Code (VS Code) which provides support for the Go language. |
| [go-plus](https://github.com/joefitzgerald/go-plus) | 1516 | 44 | 2014/03/13 | 9 months ago | Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting. |
| [go-mode.el](https://github.com/dominikh/go-mode.el) | 1217 | 53 | 2013/01/30 | 1 month ago | Go mode for GNU/Emacs. |
| [coc-go](https://github.com/josa42/coc-go) | 421 | 4 | 2019/04/25 | 1 week ago | This plugin adds [gopls](https://github.com/golang/tools/blob/master/gopls/README.md) features to Vim/Neovim. |
| [Watch](https://github.com/eaburns/Watch) | 188 | 13 | 2013/08/08 | 3 years ago | Runs a command in an acme win on file changes. |
| [goimports-reviser](https://github.com/incu6us/goimports-reviser) | 107 | 5 | 2020/04/08 | 1 month ago | Formatting tool for imports. |
| [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) | 88 | 5 | 2012/11/25 | 5 years ago | Vim plugin to highlight syntax errors on save. |
| [go-language-server](https://github.com/theia-ide/go-language-server) | 31 | 4 | 2017/11/21 | 2 years ago | A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol. |
| [gounit-vim](https://github.com/hexdigest/gounit-vim) | 22 | 2 | 2018/02/21 | 3 years ago | Vim plugin for generating Go tests based on the function's or method's signature. |
| [theia-go-extension](https://github.com/theia-ide/theia-go-extension) | 16 | 4 | 2017/11/30 | 2 years ago | Go language support for the Theia IDE. |
| [vscode-go-doc](https://github.com/msyrus/vscode-go-doc) | 5 | 2 | 2018/03/15 | 8 months ago | A Visual Studio Code extension for showing definition in output and generating go doc. |
**[⬆ back to top](#awesome-go-info)**

# Go Generate Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gotests](https://github.com/cweill/gotests) | 3703 | 76 | 2016/01/19 | 2 weeks ago | Generate Go tests from your source code. |
| [genny](https://github.com/cheekybits/genny) | 1625 | 25 | 2014/10/27 | 5 months ago | Elegant generics for Go. |
| [re2dfa](https://github.com/opennota/re2dfa) | 190 | 9 | 2015/06/20 | 3 years ago | Transform regular expressions into finite state machines and output Go source code. |
| [xgen](https://github.com/xuri/xgen) | 135 | 11 | 2019/06/22 | 3 weeks ago | XSD (XML Schema Definition) parser and Go/C/Java/Rust/TypeScript code generator. |
| [gonerics](https://github.com/bouk/gonerics) | 113 | 4 | 2014/09/29 | 7 years ago | Idiomatic Generics in Go. |
| [hasgo](https://github.com/DylanMeeus/hasgo) | 110 | 6 | 2019/05/16 | 9 months ago | Generate Haskell inspired functions for your slices. |
| [gocontracts](https://github.com/Parquery/gocontracts) | 79 | 8 | 2018/08/13 | 3 years ago | brings design-by-contract to Go by synchronizing the code with the documentation. |
| [gounit](https://github.com/hexdigest/gounit) | 61 | 5 | 2018/02/05 | 3 years ago | Generate Go tests using your own templates. |
| [generic](https://github.com/usk81/generic) | 41 | 3 | 2016/06/15 | 1 year ago | flexible data type for Go. |
| [godal](https://github.com/mafulong/godal) | 10 | 1 | 2021/03/16 | 3 months ago | Generate orm models corresponding to golang by specifying sql ddl file, which can be used by gorm. |
**[⬆ back to top](#awesome-go-info)**

# Go Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-swagger](https://github.com/go-swagger/go-swagger) | 7188 | 121 | 2014/11/16 | 1 week ago | Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API. |
| [OctoLinker](https://github.com/OctoLinker/OctoLinker) | 4882 | 90 | 2013/12/27 | 2 weeks ago | Navigate through go files efficiently with the OctoLinker browser extension for GitHub. |
| [go-callvis](https://github.com/ofabry/go-callvis) | 3817 | 75 | 2016/09/03 | 2 months ago | Visualize call graph of your Go program using dot format. |
| [depth](https://github.com/KyleBanks/depth) | 675 | 13 | 2017/03/04 | 2 months ago | Visualize dependency trees of any package by analyzing imports. |
| [richgo](https://github.com/kyoh86/richgo) | 645 | 5 | 2017/01/04 | 3 weeks ago | Enrich `go test` outputs with text decorations. |
| [rts](https://github.com/galeone/rts) | 227 | 3 | 2016/04/04 | 4 months ago | RTS: response to struct. Generates Go structs from server responses. |
| [godbg](https://github.com/tylerwince/godbg) | 182 | 4 | 2019/01/23 | 2 years ago | Implementation of Rusts `dbg!` macro for quick and easy debugging during development. |
| [typex](https://github.com/dtgorski/typex) | 136 | 3 | 2020/03/24 | 1 year ago | Examine Go types and their transitive dependencies, alternatively export results as TypeScript value objects (or types) declaration. |
| [colorgo](https://github.com/songgao/colorgo) | 108 | 4 | 2013/02/14 | 1 year ago | Wrapper around `go` command for colorized `go build` output. |
| [gothanks](https://github.com/psampaz/gothanks) | 107 | 3 | 2019/11/10 | 11 months ago | GoThanks automatically stars your go.mod github dependencies, sending this way some love to their maintainers. |
| [roumon](https://github.com/becheran/roumon) | 74 | 3 | 2021/03/02 | 10 months ago | Monitor current state of all active goroutines via a command line interface. |
| [igo](https://github.com/rocketlaunchr/igo) | 51 | 3 | 2018/11/17 | 1 year ago | Improved Go Syntax (transpiler) |
| [go-james](https://github.com/pieterclaerhout/go-james) | 50 | 3 | 2019/10/14 | 1 month ago | Go project skeleton creator, builds and tests your projects without the manual setup. |
| [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) | 39 | 2 | 2015/05/22 | 4 years ago | Bash completion for go and wgo. |
| [generator-go-lang](https://github.com/axelspringer/generator-go-lang) | 24 | 13 | 2017/09/13 | 1 year ago | A [Yeoman](https://yeoman.io) generator to get new Go projects started. |
| [docs](https://github.com/go-oas/docs) | 11 | 2 | 2021/01/28 | 11 months ago | Automatically generate RESTful API documentation for GO projects - aligned with Open API Specification standard. |
| [modver](https://github.com/bobg/modver) | 0 | 1 | 2021/07/17 | 1 week ago | Compare two versions of a Go module to check the version-number change required (major, minor, or patchlevel), according to [semver](https://semver.org/) rules. |
**[⬆ back to top](#awesome-go-info)**

# Software Packages
        
*Software written in Go.*
**[⬆ back to top](#awesome-go-info)**

# Resources
        
*Where to discover new Go libraries.*
**[⬆ back to top](#awesome-go-info)**

## E-books for purchase
        
**[⬆ back to top](#awesome-go-info)**

## Free e-books
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [GoBooks](https://github.com/dariubs/GoBooks) | 11653 | 610 | 2015/05/05 | 2 weeks ago | A curated list of Go books. |
| [The-Golang-Standard-Library-by-Example](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example) | 8378 | 586 | 2013/04/14 | 2 months ago | Golang标准库。对于程序员而言，标准库与语言本身同样重要，它好比一个百宝箱，能为各种常见的任务提供完美的解决方案。以示例驱动的方式讲解Golang的标准库。 |
| [gosuccinctly](https://github.com/thedevsir/gosuccinctly) | 21 | 4 | 2018/09/02 | 3 years ago | in Persian. |
**[⬆ back to top](#awesome-go-info)**

## Reddit
        
**[⬆ back to top](#awesome-go-info)**

## Tutorials
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang) | 39568 | 2445 | 2012/08/02 | 1 week ago | Golang ebook intro how to build a web app with golang. |
| [go-patterns](https://github.com/tmrts/go-patterns) | 18293 | 616 | 2015/12/14 | 6 months ago | Curated list of Go design patterns, recipes and idioms. |
| [learn-go-with-tests](https://github.com/quii/learn-go-with-tests) | 16509 | 313 | 2018/03/02 | 2 weeks ago | Learn Go with test-driven development. |
| [learngo](https://github.com/inancgumus/learngo) | 12772 | 284 | 2018/10/15 | 2 weeks ago | Learn Go with thousands of examples, exercises, and quizzes. |
| [golang-cheat-sheet](https://github.com/a8m/golang-cheat-sheet) | 6214 | 198 | 2014/02/13 | 2 months ago | Go's reference card. |
| [golang-for-nodejs-developers](https://github.com/miguelmota/golang-for-nodejs-developers) | 2778 | 42 | 2019/01/03 | 2 weeks ago | Examples of Golang compared to Node.js for learning. |
| [ethereum-development-with-go-book](https://github.com/miguelmota/ethereum-development-with-go-book) | 1172 | 52 | 2018/05/16 | 2 weeks ago | A little e-book on Ethereum Development with Go. |
| [working-with-go](https://github.com/mkaz/working-with-go) | 1158 | 48 | 2014/05/04 | 2 years ago | Intro to go for experienced programmers. |
| [go-clean-template](https://github.com/evrone/go-clean-template) | 543 | 22 | 2021/01/18 | 1 week ago | Clean Architecture template for Golang services. |
| [goapp](https://github.com/bnkamalesh/goapp) | 332 | 10 | 2020/07/04 | 3 months ago | An opinionated guideline to structure & develop a Go web application/service. |
| [design-patterns](https://github.com/shubhamzanwar/design-patterns) | 67 | 4 | 2020/09/24 | 1 year ago | Collection of programming design patterns implemented in Go. |
| [go-patterns](https://github.com/haveyoudebuggedit/go-patterns) | 1 | 1 | 2021/06/25 | 7 months ago | Advanced Go patterns with ready-to-run examples. |
**[⬆ back to top](#awesome-go-info)**

## Twitter
        
**[⬆ back to top](#awesome-go-info)**

# Benchmarks
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) | 1629 | 85 | 2016/04/06 | 6 months ago | Go web framework benchmark. |
| [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) | 1542 | 60 | 2013/12/16 | 2 months ago | Go HTTP request router benchmark and comparison. |
| [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) | 1251 | 38 | 2013/01/18 | 4 months ago | Benchmarks of Go serialization methods. |
| [skynet](https://github.com/atemerev/skynet) | 1007 | 49 | 2016/02/14 | 8 months ago | Skynet 1M threads microbenchmark. |
| [speedtest-resize](https://github.com/fawick/speedtest-resize) | 214 | 7 | 2013/09/16 | 1 year ago | Compare various Image resize algorithms for the Go language. |
| [go-benchmarks](https://github.com/tylertreat/go-benchmarks) | 142 | 11 | 2016/02/25 | 6 years ago | Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches. |
| [gospeed](https://github.com/feyeleanor/gospeed) | 108 | 7 | 2011/05/23 | 11 months ago | Go micro-benchmarks for calculating the speed of language constructs. |
| [autobench](https://github.com/davecheney/autobench) | 92 | 9 | 2013/03/28 | 7 years ago | Framework to compare the performance between different Go versions. |
| [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) | 60 | 5 | 2014/09/24 | 3 years ago | Collection of benchmarks for popular Go database/SQL utilities. |
| [gocostmodel](https://github.com/mna/gocostmodel) | 57 | 6 | 2014/12/19 | 9 months ago | Benchmarks of common basic operations for the Go language. |
| [kvbench](https://github.com/jimrobinson/kvbench) | 24 | 1 | 2014/04/15 | 2 years ago | Key/Value database benchmark. |
| [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) | 22 | 2 | 2017/01/24 | 4 years ago | Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results. |
| [go-ml-benchmarks](https://github.com/nikolaydubina/go-ml-benchmarks) | 19 | 1 | 2021/02/09 | 1 month ago | benchmarks for machine learning inference in Go. |
| [go-json-benchmark](https://github.com/zerosnake0/go-json-benchmark) | 6 | 2 | 2019/11/10 | 1 year ago | Go JSON benchmark. |
**[⬆ back to top](#awesome-go-info)**

# Conferences
        
**[⬆ back to top](#awesome-go-info)**

# E-Books
        
**[⬆ back to top](#awesome-go-info)**

# Gophers
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gophers](https://github.com/egonelbre/gophers) | 2602 | 59 | 2015/06/03 | 1 year ago | Free gophers. |
| [gophers](https://github.com/ashleymcnamara/gophers) | 2560 | 98 | 2017/02/15 | 2 years ago | Gopher artworks by Ashley McNamara. |
| [free-gophers-pack](https://github.com/MariaLetta/free-gophers-pack) | 2389 | 59 | 2019/04/02 | 1 year ago | Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster. |
| [gophericons](https://github.com/shalakhin/gophericons) | 592 | 21 | 2015/08/22 | 3 years ago | 34 gopher images for Go developers community |
| [gopherize.me](https://github.com/matryer/gopherize.me) | 548 | 8 | 2017/01/25 | 5 months ago | Gopherize yourself. |
| [gopher-stickers](https://github.com/tenntenn/gopher-stickers) | 521 | 14 | 2014/11/09 | 2 years ago | gopher stickers |
| [gopher-vector](https://github.com/golang-samples/gopher-vector) | 401 | 12 | 2013/03/31 | 5 years ago | Vector data of gopher |
| [go-gopher](https://github.com/sillecelik/go-gopher) | 99 | 0 | 2018/03/28 | 6 months ago | Gopher amigurumi toy pattern. |
| [gopher-logos](https://github.com/GolangUA/gopher-logos) | 97 | 9 | 2017/07/27 | 7 months ago | adorable gopher logos. |
| [gophers](https://github.com/rogeralsing/gophers) | 53 | 2 | 2017/01/28 | 1 year ago | random gopher graphics. |
| [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) | 43 | 2 | 2014/09/03 | 4 years ago | Go gopher Vector Data [.ai, .svg]. |
| [gophers](https://github.com/scraly/gophers) | 13 | 2 | 2021/06/23 | 3 months ago | Gophers by Aurélie Vache. |
**[⬆ back to top](#awesome-go-info)**

# Meetups
        
*Add the group of your city/country here (send **PR**)*
**[⬆ back to top](#awesome-go-info)**

# Style Guides
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-styleguide](https://github.com/bahlo/go-styleguide) | 1203 | 31 | 2017/07/29 | 3 weeks ago | 🏆 Opinionated Styleguide for the Go language |
**[⬆ back to top](#awesome-go-info)**

# Websites
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) | 28493 | 1716 | 2014/07/08 | 1 month ago | List of other amazingly awesome lists. |
| [awesome-remote-job](https://github.com/lukasz-madon/awesome-remote-job) | 21670 | 996 | 2015/01/02 | 2 weeks ago | Curated list of awesome remote jobs. A lot of them are looking for Go hackers. |
| [awesome-golang-workshops](https://github.com/amit-davidson/awesome-golang-workshops) | 456 | 17 | 2021/06/27 | 7 months ago | A curated list of awesome golang workshops. |
| [golang-graphics](https://github.com/mholt/golang-graphics) | 138 | 8 | 2014/03/24 | 6 years ago | Collection of Go images, graphics, and art. |
| [gocryforhelp](https://github.com/ninedraft/gocryforhelp) | 40 | 11 | 2016/05/09 | 4 years ago | Collection of Go projects that needs help. Good place to start your open-source way in Go. |
| [awesome-go-extra](https://github.com/xwjdsh/awesome-go-extra) | 19 | 2 | 2021/06/01 | 1 week ago | Parse awesome-go README file and generate a new README file with repo info. |
**[⬆ back to top](#awesome-go-info)**

# Social Media
        

> 该项目源码[Awesome Go Analysis](https://github.com/plholx/awesome-go-analysis)
> 更专业的go开源项目分析请移步 [Awesome Go](https://go.libhunt.com/)
