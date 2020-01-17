# Awesome Go Data

go语言开源项目列表，项目分类及GitHub上的开源项目数据完全来自于[awesome-go](https://github.com/avelino/awesome-go) 的[README.md](https://github.com/avelino/awesome-go/blob/master/README.md)文件，通过调用GitHub的API获取仓库信息，展示项目的star数、watch数等，方便查看go语言开源项目的一些数据。

_该文件仅包含[awesome-go](https://github.com/avelino/awesome-go) [README.md](https://github.com/avelino/awesome-go/blob/master/README.md)文件中列出的在GitHub上开源的优秀项目，不罗列其它golang相关的网站_
_该文件中的GitHub仓库信息数据会在每天凌晨1点左右更新,当前数据更新于2019-10-29 00:53:15_

- [Awesome Go](#awesome-go)
    - [Audio and Music](#audio-and-music)
    - [Authentication and OAuth](#authentication-and-oauth)
    - [Bot Building](#bot-building)
    - [Command Line](#command-line)
    - [Configuration](#configuration)
    - [Continuous Integration](#continuous-integration)
    - [CSS Preprocessors](#css-preprocessors)
    - [Data Structures](#data-structures)
    - [Database](#database)
    - [Database Drivers](#database-drivers)
        - [Relational Databases](#relational-databases)
        - [NoSQL Databases](#nosql-databases)
    - [Date and Time](#date-and-time)
    - [Distributed Systems](#distributed-systems)
    - [Email](#email)
    - [Embeddable Scripting Languages](#embeddable-scripting-languages)
    - [Error Handling](#error-handling)
    - [Files](#files)
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
        - [Microsoft Excel](#microsoft-excel)
    - [Miscellaneous](#miscellaneous)
        - [Dependency Injection](#dependency-injection)
        - [Strings](#strings)
        - [Project Layout](#project-layout)
    - [Natural Language Processing](#natural-language-processing)
    - [Networking](#networking)
        - [HTTP Clients](#http-clients)
    - [OpenGL](#opengl)
    - [ORM](#orm)
    - [Package Management](#package-management)
    - [Query Language](#query-language)
    - [Resource Embedding](#resource-embedding)
    - [Science and Data Analysis](#science-and-data-analysis)
    - [Security](#security)
    - [Serialization](#serialization)
    - [Template Engines](#template-engines)
    - [Testing](#testing)
        - [Testing Frameworks](#testing-frameworks)
        - [Mock](#mock)
        - [Fail injection](#fail-injection)
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
        - [Middlewares](#middlewares)
            - [Actual middlewares](#actual-middlewares)
            - [Libraries for creating HTTP middlewares](#libraries-for-creating-http-middlewares)
        - [Routers](#routers)
    - [Windows](#windows)
    - [XML](#xml)
    - [Server Applications](#server-applications)
    - [Performance](#performance)
    - [Stream Processing](#stream-processing)
- [Tools](#tools)
    - [Code Analysis](#code-analysis)
    - [Editor Plugins](#editor-plugins)
    - [Go Generate Tools](#go-generate-tools)
    - [Go Tools](#go-tools)
    - [Software Packages](#software-packages)
        - [DevOps Tools](#devops-tools)
        - [Other Software](#other-software)
- [Resources](#resources)
    - [Benchmarks](#benchmarks)
    - [Conferences](#conferences)
    - [E-Books](#e-books)
    - [Gophers](#gophers)
    - [Meetups](#meetups)
    - [Twitter](#twitter)
    - [Websites](#websites)
        - [Tutorials](#tutorials)




# Awesome Go
        

## Audio and Music
        
*Libraries for manipulating audio.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [oto](https://github.com/hajimehoshi/oto) | 464 | 9 | 2017-05-04 | 1 day ago | A low-level library to play sound on multiple platforms. |
| [portaudio](https://github.com/gordonklaus/portaudio) | 313 | 12 | 2015-09-16 | 1 year ago | Go bindings for the PortAudio audio I/O library. |
| [waveform](https://github.com/mdlayher/waveform) | 262 | 13 | 2014-09-14 | 3 years ago | Go package capable of generating waveform images from audio streams. |
| [music-theory](https://github.com/go-music-theory/music-theory) | 262 | 12 | 2016-03-17 | 1 year ago | Music theory models in Go. |
| [portmidi](https://github.com/rakyll/portmidi) | 212 | 7 | 2013-11-10 | 2 years ago | Go bindings for PortMidi. |
| [id3v2](https://github.com/bogem/id3v2) | 120 | 5 | 2016-05-16 | 3 weeks ago | Fast and stable ID3 parsing and writing library for Go. |
| [flac](https://github.com/mewkiz/flac) | 105 | 9 | 2012-11-02 | 8 months ago | Native Go FLAC encoder/decoder with support for FLAC streams. |
| [mp3](https://github.com/tcolgate/mp3) | 100 | 1 | 2015-02-27 | 2 years ago | Native Go MP3 decoder. |
| [mix](https://github.com/go-mix/mix) | 100 | 3 | 2016-01-03 | 2 years ago | Sequence-based Go-native audio mixer for music apps. |
| [go-sox](https://github.com/krig/go-sox) | 95 | 8 | 2013-10-08 | 1 year ago | libsox bindings for go. |
| [flac](https://github.com/eaburns/flac) | 87 | 5 | 2013-08-21 | 2 years ago | No-frills native Go FLAC decoder that decodes FLAC files into byte slices. |
| [malgo](https://github.com/gen2brain/malgo) | 81 | 4 | 2017-11-10 | 3 weeks ago | Mini audio library. |
| [go-taglib](https://github.com/wtolson/go-taglib) | 68 | 6 | 2012-11-20 | 1 year ago | Go bindings for taglib. |
| [gaad](https://github.com/Comcast/gaad) | 60 | 11 | 2016-07-11 | 1 year ago | Native Go AAC bitstream parser. |
| [minimp3](https://github.com/tosone/minimp3) | 31 | 2 | 2018-01-26 | 3 months ago | Lightweight MP3 decoder library. |
| [go_mediainfo](https://github.com/zhulik/go_mediainfo) | 27 | 1 | 2015-12-14 | 3 years ago | libmediainfo bindings for go. |
| [EasyMIDI](https://github.com/algoGuy/EasyMIDI) | 23 | 3 | 2018-02-20 | 1 year ago | EasyMidi is a simple and reliable library for working with standard midi file (SMF). |
| [vorbis](https://github.com/mccoyst/vorbis) | 23 | 3 | 2013-07-12 | 7 months ago | "Native" Go Vorbis decoder (uses CGO, but has no dependencies). |
| [gosamplerate](https://github.com/dh1tw/gosamplerate) | 8 | 1 | 2016-11-21 | 2 weeks ago | libsamplerate bindings for go. |

## Authentication and OAuth
        
*Libraries for implementing authentications schemes.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [jwt-go](https://github.com/dgrijalva/jwt-go) | 6443 | 143 | 2012-04-18 | 1 month ago | Golang implementation of JSON Web Tokens (JWT). |
| [casbin](https://github.com/casbin/casbin) | 5453 | 172 | 2017-04-08 | 1 week ago | Authorization library that supports access control models like ACL, RBAC, ABAC. |
| [oauth2](https://github.com/golang/oauth2) | 2533 | 101 | 2014-04-14 | 2 weeks ago | Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support. |
| [goth](https://github.com/markbates/goth) | 2375 | 63 | 2014-10-15 | 3 days ago | provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box. |
| [authboss](https://github.com/volatiletech/authboss) | 2022 | 41 | 2015-01-03 | 6 days ago | Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time. |
| [osin](https://github.com/openshift/osin) | 1558 | 67 | 2013-09-11 | 3 days ago | Golang OAuth2 server library. |
| [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) | 1351 | 75 | 2015-11-01 | 1 week ago | Standalone, specification-compliant,  OAuth2 server written in Golang. |
| [go-jose](https://github.com/square/go-jose) | 1348 | 64 | 2014-11-15 | 5 days ago | Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs. |
| [gologin](https://github.com/dghubble/gologin) | 1103 | 28 | 2015-06-23 | 2 months ago | chainable handlers for login with OAuth1 and OAuth2 authentication providers. |
| [gorbac](https://github.com/mikespook/gorbac) | 948 | 58 | 2013-12-26 | 7 months ago | provides a lightweight role-based access control (RBAC) implementation in Golang. |
| [loginsrv](https://github.com/tarent/loginsrv) | 842 | 46 | 2016-11-11 | 2 days ago | JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam. |
| [scs](https://github.com/alexedwards/scs) | 613 | 19 | 2016-08-09 | 1 day ago | Session Manager for HTTP servers. |
| [permissions2](https://github.com/xyproto/permissions2) | 366 | 13 | 2014-11-19 | 3 weeks ago | Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt. |
| [paseto](https://github.com/o1egl/paseto) | 252 | 14 | 2018-01-23 | 2 weeks ago | Golang implementation of Platform-Agnostic Security Tokens (PASETO). |
| [httpauth](https://github.com/goji/httpauth) | 188 | 6 | 2014-05-27 | 3 years ago | HTTP Authentication middleware. |
| [jeff](https://github.com/abraithwaite/jeff) | 179 | 4 | 2018-08-03 | 2 weeks ago | Simple, flexible, secure and idiomatic web session management with pluggable backends. |
| [jwt-auth](https://github.com/adam-hanna/jwt-auth) | 161 | 11 | 2016-07-06 | 9 months ago | JWT middleware for Golang http servers with many configuration options. |
| [jwt](https://github.com/pascaldekloe/jwt) | 116 | 9 | 2018-03-21 | 2 weeks ago | Lightweight JSON Web Token (JWT) library. |
| [session](https://github.com/icza/session) | 93 | 6 | 2016-02-08 | 3 months ago | Go session management for web servers (including support for Google App Engine - GAE). |
| [branca](https://github.com/hako/branca) | 83 | 7 | 2018-01-09 | 1 year ago | Golang implementation of Branca Tokens. |
| [jwt](https://github.com/robbert229/jwt) | 76 | 6 | 2016-06-06 | 11 months ago | Clean and easy to use implementation of JSON Web Tokens (JWT). |
| [sessionup](https://github.com/swithek/sessionup) | 66 | 5 | 2019-07-24 | 1 week ago | Simple, yet effective HTTP session management and identification package. |
| [sessions](https://github.com/adam-hanna/sessions) | 47 | 3 | 2017-04-29 | 5 months ago | Dead simple, highly performant, highly customizable sessions service for go http servers. |
| [sjwt](https://github.com/brianvoe/sjwt) | 39 | 0 | 2019-06-20 | 1 month ago | Simple jwt generator and parser. |
| [rbac](https://github.com/zpatrick/rbac) | 35 | 3 | 2018-08-02 | 1 year ago | Minimalistic RBAC package for Go applications. |
| [securecookie](https://github.com/chmike/securecookie) | 34 | 5 | 2017-09-03 | 1 year ago | Efficient secure cookie encoding/decoding. |
| [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) | 8 | 2 | 2017-10-20 | 11 months ago | Go session management using the SessionGate Redis module. |
| [signedvalue](https://github.com/sashka/signedvalue) | 8 | 0 | 2018-01-06 | 1 month ago | Signed and timestamped strings compatible with [Tornado's](https://github.com/tornadoweb/tornado) `create_signed_value`, `decode_signed_value`, and therefore `set_secure_cookie` and `get_secure_cookie`. |
| [cookiestxt](https://github.com/mengzhuo/cookiestxt) | 2 | 1 | 2017-10-09 | 2 years ago | provides parser of cookies.txt file format. |
| [scope](https://github.com/SonicRoshan/scope) | 0 | 1 | 2019-09-23 | 1 month ago | Easily Manage OAuth2 Scopes In Go. |

## Bot Building
        
*Libraries for building and working with bots.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) | 1748 | 60 | 2015-06-25 | 3 days ago | Simple and clean Telegram bot client. |
| [telebot](https://github.com/tucnak/telebot) | 997 | 42 | 2015-06-26 | 3 weeks ago | Telegram bot framework written in Go. |
| [bot](https://github.com/go-chat-bot/bot) | 511 | 37 | 2015-09-23 | 6 days ago | IRC, Slack & Telegram bot written in Go. |
| [slacker](https://github.com/shomali11/slacker) | 332 | 12 | 2017-05-20 | 2 months ago | Easy to use framework to create Slack bots. |
| [golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot) | 252 | 24 | 2017-05-15 | 8 months ago | A golang implementation of a console-based trading bot for cryptocurrency exchanges. |
| [tbot](https://github.com/yanzay/tbot) | 233 | 10 | 2015-09-12 | 1 month ago | Telegram bot server with API similar to net/http. |
| [kelp](https://github.com/stellar/kelp) | 215 | 31 | 2018-08-09 | 2 days ago | official trading and market-making bot for the [Stellar](https://www.stellar.org/) DEX. Works out-of-the-box, written in Golang, compatible with centralized exchanges and custom trading strategies. |
| [tenyks](https://github.com/kyleterry/tenyks) | 168 | 14 | 2012-08-26 | 1 month ago | Service oriented IRC bot using Redis and JSON for messaging. |
| [go-sarah](https://github.com/oklahomer/go-sarah) | 148 | 7 | 2016-11-06 | 1 week ago | Framework to build bot for desired chat services including LINE, Slack, Gitter and more. |
| [hanu](https://github.com/sbstjn/hanu) | 114 | 5 | 2016-09-16 | 3 days ago | Framework for writing Slack bots. |
| [go-tgbot](https://github.com/olebedev/go-tgbot) | 89 | 8 | 2016-12-11 | 1 year ago | Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware. |
| [margelet](https://github.com/zhulik/margelet) | 60 | 5 | 2015-11-21 | 3 years ago | Framework for building Telegram bots. |
| [govkbot](https://github.com/nikepan/govkbot) | 29 | 3 | 2016-07-12 | 1 month ago | Simple Go [VK](https://vk.com) bot library. |
| [slackscot](https://github.com/alexandre-normand/slackscot) | 16 | 1 | 2015-10-22 | 1 week ago | Another framework for building Slack bots. |
| [micha](https://github.com/onrik/micha) | 10 | 3 | 2016-04-14 | 1 week ago | Go Library for Telegram bot api. |

## Command Line
        
*Libraries for building Console Applications and Console User Interfaces.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [cobra](https://github.com/spf13/cobra) | 14353 | 301 | 2013-09-04 | 3 days ago | Commander for modern Go CLI interactions. |
| [cli](https://github.com/urfave/cli) | 12165 | 285 | 2013-07-14 | 4 days ago | Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli). |
| [termui](https://github.com/gizak/termui) | 9205 | 283 | 2015-02-03 | 1 month ago | Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib). |
| [gocui](https://github.com/jroimartin/gocui) | 5640 | 118 | 2014-01-04 | 1 day ago | Minimalist Go library aimed at creating Console User Interfaces. |
| [termbox-go](https://github.com/nsf/termbox-go) | 3549 | 104 | 2012-01-13 | 3 weeks ago | Termbox is a library for creating cross-platform text-based interfaces. |
| [color](https://github.com/fatih/color) | 3107 | 56 | 2014-02-17 | 1 year ago | Versatile package for colored terminal output. |
| [kingpin](https://github.com/alecthomas/kingpin) | 2673 | 53 | 2014-05-15 | 1 week ago | Command line and flag parser supporting sub commands. |
| [go-prompt](https://github.com/c-bata/go-prompt) | 2599 | 45 | 2017-08-15 | 6 days ago | Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit). |
| [uiprogress](https://github.com/gosuri/uiprogress) | 1583 | 33 | 2015-11-17 | 3 months ago | Flexible library to render progress bars in terminal applications. |
| [go-flags](https://github.com/jessevdk/go-flags) | 1554 | 28 | 2012-08-31 | 3 weeks ago | go command line option parser. |
| [readline](https://github.com/chzyer/readline) | 1398 | 39 | 2015-09-20 | 3 months ago | Pure golang implementation that provides most features in GNU-Readline under MIT license. |
| [dnote](https://github.com/dnote/dnote) | 1237 | 27 | 2017-03-31 | 7 hours ago | A simple and end-to-end encrypted notebook for developers. |
| [asciigraph](https://github.com/guptarohit/asciigraph) | 1209 | 25 | 2018-06-17 | 3 weeks ago | Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. |
| [docopt.go](https://github.com/docopt/docopt.go) | 1193 | 34 | 2013-08-26 | 1 month ago | Command-line arguments parser that will make you smile. |
| [cli](https://github.com/mitchellh/cli) | 1025 | 24 | 2013-11-03 | 11 months ago | Go library for implementing command-line interfaces. |
| [uilive](https://github.com/gosuri/uilive) | 1016 | 13 | 2015-11-16 | 2 months ago | Library for updating terminal output in realtime. |
| [gcli](https://github.com/tcnksm/gcli) | 878 | 26 | 2014-06-19 | 1 year ago | The easy way to start building Golang command line applications. |
| [pflag](https://github.com/spf13/pflag) | 870 | 27 | 2013-08-30 | 4 days ago | Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. |
| [termdash](https://github.com/mum4k/termdash) | 860 | 19 | 2018-03-24 | 3 months ago | Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui). |
| [go-arg](https://github.com/alexflint/go-arg) | 787 | 15 | 2015-11-01 | 6 days ago | Struct-based argument parsing in Go. |
| [progressbar](https://github.com/schollz/progressbar) | 780 | 19 | 2017-10-27 | 1 month ago | Basic thread-safe progress bar that works in every OS. |
| [mpb](https://github.com/vbauerster/mpb) | 763 | 16 | 2016-12-14 | 3 days ago | Multi progress bar for terminal applications. |
| [aurora](https://github.com/logrusorgru/aurora) | 695 | 6 | 2016-11-07 | 1 week ago | ANSI terminal colors that supports fmt.Printf/Sprintf. |
| [complete](https://github.com/posener/complete) | 636 | 13 | 2017-05-06 | 4 days ago | Write bash completions in Go + Go command bash completion. |
| [mow.cli](https://github.com/jawher/mow.cli) | 632 | 17 | 2014-12-19 | 1 month ago | Go library for building CLI applications with sophisticated flag and argument parsing and validation. |
| [liner](https://github.com/peterh/liner) | 605 | 24 | 2012-08-16 | 2 months ago | Go readline-like library for command-line interfaces. |
| [uitable](https://github.com/gosuri/uitable) | 516 | 14 | 2015-11-14 | 5 months ago | Library to improve readability in terminal apps using tabular data. |
| [cli](https://github.com/mkideal/cli) | 489 | 22 | 2016-02-27 | 1 month ago | Feature-rich and easy to use command-line package based on golang struct tags. |
| [flaggy](https://github.com/integrii/flaggy) | 470 | 15 | 2018-03-05 | 13 hours ago | A robust and idiomatic flags package with excellent subcommand support. |
| [go-colorable](https://github.com/mattn/go-colorable) | 392 | 17 | 2014-07-30 | 3 weeks ago | Colorable writer for windows. |
| [go-isatty](https://github.com/mattn/go-isatty) | 362 | 8 | 2014-04-01 | 2 weeks ago | isatty for golang. |
| [chalk](https://github.com/ttacon/chalk) | 314 | 8 | 2014-07-19 | 2 months ago | Intuitive package for prettifying terminal/console output. |
| [ops](https://github.com/nanovms/ops) | 291 | 18 | 2018-09-11 | 1 month ago | Unikernel Builder/Orchestrator. |
| [tabby](https://github.com/cheynewallace/tabby) | 253 | 1 | 2018-12-18 | 7 months ago | A tiny library for super simple Golang tables. |
| [color](https://github.com/gookit/color) | 245 | 8 | 2018-07-01 | 1 month ago | Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows. |
| [termtables](https://github.com/apcera/termtables) | 214 | 71 | 2012-12-06 | 2 years ago | Go port of the Ruby library [terminal-tables](https://github.com/tj/terminal-table) for simple ASCII table generation as well as providing markdown and HTML output. |
| [go-colortext](https://github.com/daviddengcn/go-colortext) | 199 | 9 | 2013-01-23 | 1 year ago | Go library for color output in terminals. |
| [simpletable](https://github.com/alexeyco/simpletable) | 175 | 2 | 2017-03-29 | 5 days ago | Simple tables in terminal with Go. |
| [argparse](https://github.com/akamensky/argparse) | 134 | 7 | 2017-11-24 | 3 weeks ago | Command line argument parser inspired by Python's argparse module. |
| [flag](https://github.com/cosiner/flag) | 102 | 5 | 2016-10-06 | 7 months ago | Simple but powerful command line option parsing library for Go supporting subcommand. |
| [clif](https://github.com/ukautz/clif) | 99 | 2 | 2015-05-31 | 8 months ago | Small command line interface framework. |
| [sflags](https://github.com/octago/sflags) | 97 | 5 | 2016-12-04 | 2 months ago | Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries. |
| [commandeer](https://github.com/jaffee/commandeer) | 96 | 6 | 2017-10-12 | 2 weeks ago | Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags. |
| [wmenu](https://github.com/dixonwille/wmenu) | 92 | 2 | 2016-04-20 | 5 months ago | Easy to use menu structure for cli applications that prompts users to make choices. |
| [cfmt](https://github.com/mingrammer/cfmt) | 67 | 3 | 2018-03-16 | 10 months ago | Contextual fmt inspired by bootstrap color classes. |
| [job](https://github.com/liujianping/job) | 59 | 2 | 2019-04-09 | 5 months ago | JOB, make your short-term command as a long-term job. |
| [cli](https://github.com/teris-io/cli) | 58 | 1 | 2017-05-25 | 5 months ago | Simple and complete API for building command line interfaces in Go. |
| [env](https://github.com/codingconcepts/env) | 44 | 1 | 2017-06-15 | 4 months ago | Tag-based environment configuration for structs. |
| [wlog](https://github.com/dixonwille/wlog) | 37 | 1 | 2016-04-14 | 2 years ago | Simple logging interface that supports cross-platform color and concurrency. |
| [gocmd](https://github.com/devfacet/gocmd) | 36 | 3 | 2018-01-08 | 1 year ago | Go library for building command line applications. |
| [tabular](https://github.com/InVisionApp/tabular) | 34 | 4 | 2018-04-24 | 1 year ago | Print ASCII tables from command line utilities without the need to pass large sets of data to the API. |
| [strumt](https://github.com/antham/strumt) | 31 | 0 | 2017-06-20 | 2 months ago | Library to create prompt chain. |
| [flagvar](https://github.com/sgreben/flagvar) | 31 | 2 | 2018-05-19 | 4 months ago | A collection of flag argument types for Go's standard `flag` package. |
| [cmdr](https://github.com/hedzr/cmdr) | 21 | 4 | 2019-05-15 | 1 week ago | A POSIX/GNU style, getopt-like command-line UI Go library. |
| [argv](https://github.com/cosiner/argv) | 18 | 1 | 2017-01-22 | 4 months ago | Go library to split command line string as arguments array using the bash syntax. |
| [colourize](https://github.com/TreyBastian/colourize) | 16 | 3 | 2015-05-11 | 3 years ago | Go library for ANSI colour text in terminals. |
| [go-getoptions](https://github.com/DavidGamba/go-getoptions) | 16 | 3 | 2015-12-18 | 4 days ago | Go option parser inspired on the flexibility of Perl’s GetOpt::Long. |
| [go-commander](https://github.com/yitsushi/go-commander) | 15 | 1 | 2016-10-10 | 1 month ago | Go library to simplify CLI workflow. |
| [ctc](https://github.com/wzshiming/ctc) | 10 | 1 | 2018-04-28 | 1 day ago | The non-invasive cross-platform terminal color library does not need to modify the Print method. |
| [sand](https://github.com/Zaba505/sand) | 9 | 1 | 2018-11-19 | 11 months ago | Simple API for creating interpreters and so much more. |
| [go-ataman](https://github.com/workanator/go-ataman) | 8 | 2 | 2017-05-18 | 2 years ago | Go library for rendering ANSI colored text templates in terminals. |
| [ts](https://github.com/liujianping/ts) | 7 | 0 | 2019-06-25 | 3 months ago | Timestamp convert & compare tool. |

## Configuration
        
*Libraries for configuration parsing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [viper](https://github.com/spf13/viper) | 10118 | 197 | 2014-04-02 | 1 week ago | Go configuration with fangs. |
| [envconfig](https://github.com/kelseyhightower/envconfig) | 2530 | 40 | 2013-11-07 | 1 month ago | Go library for managing configuration data from environment variables. |
| [godotenv](https://github.com/joho/godotenv) | 2326 | 33 | 2013-07-30 | 1 week ago | Go port of Ruby's dotenv library (Loads environment variables from `.env`). |
| [ini](https://github.com/go-ini/ini) | 1699 | 67 | 2014-12-18 | 1 day ago | Go package to read and write INI files. |
| [env](https://github.com/caarlos0/env) | 1274 | 18 | 2015-07-28 | 2 weeks ago | Parse environment variables to Go structs (with defaults). |
| [konfig](https://github.com/lalamove/konfig) | 531 | 11 | 2019-01-19 | 6 hours ago | Composable, observable and performant config handling for Go for the distributed processing era. |
| [confita](https://github.com/heetch/confita) | 271 | 22 | 2017-12-21 | 1 week ago | Load configuration in cascade from multiple backends into a struct. |
| [store](https://github.com/tucnak/store) | 245 | 5 | 2015-10-04 | 2 years ago | Lightweight configuration manager for Go. |
| [config](https://github.com/olebedev/config) | 217 | 8 | 2014-04-21 | 5 months ago | JSON or YAML configuration wrapper with environment variables and flags parsing. |
| [config](https://github.com/JeremyLoy/config) | 199 | 1 | 2019-04-02 | 5 months ago | Cloud native application configuration. Bind ENV to structs in only two lines. |
| [config](https://github.com/joshbetz/config) | 197 | 2 | 2017-04-03 | 2 days ago | Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP. |
| [hjson-go](https://github.com/hjson/hjson-go) | 178 | 8 | 2016-08-06 | 2 months ago | Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments. |
| [envconfig](https://github.com/vrischmann/envconfig) | 154 | 4 | 2015-04-22 | 3 months ago | Read your configuration from environment variables. |
| [gcfg](https://github.com/go-gcfg/gcfg) | 118 | 6 | 2015-08-17 | 1 year ago | read INI-style configuration files into Go structs; supports user-defined types and subsections. |
| [goconfig](https://github.com/crgimenes/goconfig) | 115 | 11 | 2016-12-18 | 2 weeks ago | Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. |
| [config](https://github.com/gookit/config) | 104 | 5 | 2018-07-07 | 1 week ago | application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge. |
| [koanf](https://github.com/knadh/koanf) | 102 | 4 | 2019-06-18 | 3 weeks ago | Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line. |
| [envh](https://github.com/antham/envh) | 92 | 3 | 2017-01-12 | 2 months ago | Helpers to manage environment variables. |
| [envcfg](https://github.com/tomazk/envcfg) | 90 | 1 | 2014-11-29 | 2 years ago | Un-marshaling environment variables to Go structs. |
| [gofigure](https://github.com/ian-kent/gofigure) | 57 | 6 | 2014-11-25 | 1 month ago | Go application configuration made easy. |
| [configure](https://github.com/paked/configure) | 50 | 3 | 2015-06-14 | 8 months ago | Provides configuration through multiple sources, including JSON, flags and environment variables. |
| [harvester](https://github.com/beatlabs/harvester) | 44 | 11 | 2019-04-09 | 3 weeks ago | Harvester, a easy to use static and dynamic configuration package supportig seeding, env vars and Consul integration. |
| [xdg](https://github.com/OpenPeeDeeP/xdg) | 40 | 4 | 2017-07-20 | 7 months ago | Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html). |
| [go-up](https://github.com/ufoscout/go-up) | 26 | 1 | 2018-02-18 | 7 months ago | A simple configuration library with recursive placeholders resolution and no magic. |
| [ingo](https://github.com/schachmat/ingo) | 25 | 1 | 2016-02-08 | 2 years ago | Flags persisted in an ini-like config file. |
| [mini](https://github.com/sasbury/mini) | 20 | 2 | 2015-04-30 | 10 months ago | Golang package for parsing ini-style configuration files. |
| [conflate](https://github.com/the4thamigo-uk/conflate) | 9 | 0 | 2018-02-02 | 6 months ago | Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema. |
| [genv](https://github.com/sakirsensoy/genv) | 8 | 1 | 2019-07-15 | 3 months ago | Read environment variables easily with dotenv support. |
| [envconf](https://github.com/ian-kent/envconf) | 7 | 1 | 2014-10-26 | 5 years ago | Configuration from environment. |
| [sprbox](https://github.com/oblq/sprbox) | 4 | 2 | 2018-07-17 | 3 months ago | Build-environment aware toolbox factory and agnostic config parser (YAML, TOML, JSON and Environment vars). |
| [env](https://github.com/nasermirzaei89/env) | 2 | 0 | 2019-07-24 | 2 weeks ago | Simple useful package for read environment variables. |

## Continuous Integration
        
*Tools for help with continuous integration.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [drone](https://github.com/drone/drone) | 19688 | 583 | 2014-02-07 | 2 days ago | Drone is a Continuous Integration platform built on Docker, written in Go. |
| [goveralls](https://github.com/mattn/goveralls) | 602 | 16 | 2013-04-17 | 7 hours ago | Go integration for Coveralls.io continuous code coverage tracking system. |
| [overalls](https://github.com/go-playground/overalls) | 98 | 3 | 2015-07-30 | 1 year ago | Multi-Package go project coverprofile for tools like goveralls. |
| [duci](https://github.com/duck8823/duci) | 46 | 3 | 2018-04-01 | 1 week ago | A simple ci server no needs domain specific languages. |
| [gomason](https://github.com/nikogura/gomason) | 36 | 0 | 2017-11-18 | 4 months ago | Test, Build, Sign, and Publish your go binaries from a clean workspace. |
| [roveralls](https://github.com/lawrencewoodman/roveralls) | 12 | 1 | 2016-06-26 | 1 year ago | Recursive coverage testing tool. |

## CSS Preprocessors
        
*Libraries for preprocessing CSS files.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gcss](https://github.com/yosssi/gcss) | 426 | 16 | 2014-09-04 | 5 years ago | Pure Go CSS Preprocessor. |
| [go-libsass](https://github.com/wellington/go-libsass) | 141 | 6 | 2015-04-19 | 8 months ago | Go wrapper to the 100% Sass compatible libsass project. |

## Data Structures
        
*Generic datastructures and algorithms in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gods](https://github.com/emirpasic/gods) | 6904 | 309 | 2015-03-04 | 2 weeks ago | Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc. |
| [go-datastructures](https://github.com/Workiva/go-datastructures) | 5281 | 302 | 2014-10-29 | 2 months ago | Collection of useful, performant, and thread-safe data structures. |
| [golang-set](https://github.com/deckarep/golang-set) | 1258 | 36 | 2013-07-04 | 3 months ago | Thread-Safe and Non-Thread-Safe high-performance sets for Go. |
| [BoomFilters](https://github.com/tylertreat/BoomFilters) | 1191 | 40 | 2015-02-06 | 1 year ago | Probabilistic data structures for processing continuous, unbounded streams. |
| [gota](https://github.com/go-gota/gota) | 947 | 57 | 2016-02-07 | 4 weeks ago | Implementation of dataframes, series, and data wrangling methods for Go. |
| [roaring](https://github.com/RoaringBitmap/roaring) | 725 | 34 | 2014-07-11 | 2 weeks ago | Go package implementing compressed bitsets. |
| [bloom](https://github.com/willf/bloom) | 703 | 29 | 2011-05-21 | 2 weeks ago | Go package implementing Bloom filters. |
| [hyperloglog](https://github.com/axiomhq/hyperloglog) | 672 | 16 | 2017-06-18 | 4 months ago | HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction. |
| [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) | 534 | 18 | 2015-06-29 | 7 months ago | Cuckoo filter: a good alternative to a counting bloom filter implemented in Go. |
| [bitset](https://github.com/willf/bitset) | 506 | 29 | 2011-05-11 | 2 weeks ago | Go package implementing bitsets. |
| [trie](https://github.com/derekparker/trie) | 438 | 18 | 2014-03-07 | 2 months ago | Trie implementation in Go. |
| [go-geoindex](https://github.com/hailocab/go-geoindex) | 315 | 68 | 2015-01-22 | 1 year ago | In-memory geo index. |
| [algorithms](https://github.com/shady831213/algorithms) | 285 | 11 | 2018-01-31 | 6 months ago | Algorithms and data structures.CLRS study. |
| [mafsa](https://github.com/smartystreets-archives/mafsa) | 275 | 16 | 2014-11-08 | 4 months ago | MA-FSA implementation with Minimal Perfect Hashing. |
| [mafsa](https://github.com/smartystreets/mafsa) | 273 | 17 | 2014-11-08 | 4 months ago | MA-FSA implementation with Minimal Perfect Hashing. |
| [gocache](https://github.com/eko/gocache) | 235 | 9 | 2019-10-05 | 2 days ago | A complete Go cache library with mutiple stores (memory, memcache, redis, ...), chainable, loadable, metrics cache and more. |
| [goskiplist](https://github.com/ryszard/goskiplist) | 201 | 14 | 2012-05-09 | 2 years ago | Skip list implementation in Go. |
| [hilbert](https://github.com/google/hilbert) | 188 | 20 | 2015-08-06 | 11 months ago | Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves. |
| [merkletree](https://github.com/cbergoon/merkletree) | 162 | 5 | 2017-04-12 | 4 weeks ago | Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures. |
| [bloom](https://github.com/zentures/bloom) | 131 | 7 | 2013-09-03 | 1 year ago | Bloom filters implemented in Go. |
| [binpacker](https://github.com/zhuangsirui/binpacker) | 130 | 11 | 2016-02-02 | 1 year ago | Binary packer and unpacker helps user build custom binary stream. |
| [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) | 113 | 7 | 2016-04-01 | 2 months ago | Go implementation of Adaptive Radix Tree. |
| [ttlcache](https://github.com/ReneKroon/ttlcache) | 109 | 10 | 2014-12-13 | 19 hours ago | In-memory LRU string-interface{} map with expiration for golang. |
| [skiplist](https://github.com/MauriceGit/skiplist) | 105 | 5 | 2018-06-24 | 10 months ago | Very fast Go Skiplist implementation. |
| [go-rquad](https://github.com/arl/go-rquad) | 102 | 3 | 2016-09-13 | 1 year ago | Region quadtrees with efficient point location and neighbour finding. |
| [encoding](https://github.com/zentures/encoding) | 97 | 6 | 2013-09-21 | 1 year ago | Integer Compression Libraries for Go. |
| [ring](https://github.com/TheTannerRyan/ring) | 91 | 1 | 2019-01-27 | 2 weeks ago | Go implementation of a high performance, thread safe bloom filter. |
| [conjungo](https://github.com/InVisionApp/conjungo) | 84 | 16 | 2016-12-30 | 5 months ago | A small, powerful and flexible merge library. |
| [deque](https://github.com/gammazero/deque) | 80 | 6 | 2018-04-24 | 5 months ago | Fast ring-buffer deque (double-ended queue). |
| [skiplist](https://github.com/gansidui/skiplist) | 66 | 7 | 2014-11-19 | 5 years ago | Skiplist implementation in Go. |
| [bit](https://github.com/yourbasic/bit) | 64 | 5 | 2017-05-04 | 1 year ago | Golang set data structure with bonus bit-twiddling functions. |
| [levenshtein](https://github.com/agnivade/levenshtein) | 58 | 3 | 2014-07-30 | 2 days ago | Implementation to calculate levenshtein distance in Go. |
| [bloom](https://github.com/yourbasic/bloom) | 44 | 2 | 2017-05-07 | 2 years ago | Golang Bloom filter implementation. |
| [count-min-log](https://github.com/seiflotfy/count-min-log) | 43 | 3 | 2015-08-17 | 2 years ago | Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory). |
| [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) | 41 | 1 | 2019-01-11 | 2 months ago | Concurrent FIFO queue. |
| [go-mcache](https://github.com/OrlovEvgeny/go-mcache) | 38 | 3 | 2018-04-15 | 5 months ago | Fast in-memory key:value store/cache library. Pointer caches. |
| [levenshtein](https://github.com/agext/levenshtein) | 35 | 1 | 2016-04-08 | 3 months ago | Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. |
| [concurrent-writer](https://github.com/free/concurrent-writer) | 25 | 4 | 2017-09-18 | 1 year ago | Highly concurrent drop-in replacement for `bufio.Writer`. |
| [remember-go](https://github.com/rocketlaunchr/remember-go) | 25 | 2 | 2019-04-05 | 1 week ago | A universal interface for caching data (redis, memory, memcached etc). |
| [crunch](https://github.com/superwhiskers/crunch) | 19 | 3 | 2019-02-27 | 2 weeks ago | Go package implementing buffers for handling various datatypes easily. |
| [pipeline](https://github.com/hyfather/pipeline) | 18 | 1 | 2018-04-25 | 1 year ago | An implementation of pipelines with fan-in and fan-out. |
| [goset](https://github.com/zoumo/goset) | 17 | 1 | 2017-08-25 | 5 days ago | A useful Set collection implementation for Go. |
| [typ](https://github.com/gurukami/typ) | 14 | 1 | 2019-03-03 | 2 months ago | Null Types, Safe primitive type conversion and fetching value from complex structures. |
| [go-ef](https://github.com/amallia/go-ef) | 11 | 2 | 2017-09-22 | 2 years ago | A Go implementation of the Elias-Fano encoding. |
| [hide](https://github.com/emvi/hide) | 10 | 3 | 2019-01-16 | 7 months ago | ID type with marshalling to/from hash to prevent sending IDs to clients. |
| [deque](https://github.com/edwingeng/deque) | 9 | 1 | 2019-02-01 | 1 month ago | A highly optimized double-ended queue. |
| [dict](https://github.com/srfrog/dict) | 9 | 0 | 2019-04-23 | 4 months ago | Python-like dictionaries (dict) for Go. |
| [mspm](https://github.com/BlackRabbitt/mspm) | 8 | 3 | 2018-05-18 | 1 year ago | Multi-String Pattern Matching Algorithm for information retrieval. |
| [null](https://github.com/emvi/null) | 7 | 2 | 2018-07-05 | 1 month ago | Nullable Go types that can be marshalled/unmarshalled to/from JSON. |
| [treap](https://github.com/perdata/treap) | 7 | 2 | 2018-09-16 | 1 year ago | Persistent, fast ordered map using tree heaps. |
| [set](https://github.com/StudioSol/set) | 7 | 17 | 2018-07-21 | 3 months ago | Simple set data structure implementation in Go using LinkedHashMap. |
| [timedmap](https://github.com/zekroTJA/timedmap) | 3 | 1 | 2019-01-30 | 3 weeks ago | Map with expiring key-value pairs. |
| [parsefields](https://github.com/MonaxGT/parsefields) | 3 | 1 | 2019-04-13 | 5 months ago | Tools for parse JSON-like logs for collecting unique fields and events. |
| [gofal](https://github.com/xxjwxc/gofal) | 3 | 1 | 2019-08-05 | 2 weeks ago | fractional api for Go. |
| [ptrie](https://github.com/viant/ptrie) | 2 | 7 | 2019-05-20 | 2 months ago | An implementation of prefix tree. |

## Database
        
*SQL query builder, libraries for building and using SQL.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [prometheus](https://github.com/prometheus/prometheus) | 27209 | 1086 | 2012-11-24 | 1 hour ago | Monitoring system and time series database. |
| [tidb](https://github.com/pingcap/tidb) | 21067 | 1276 | 2015-09-06 | 1 hour ago | TiDB is a distributed SQL database. Inspired by the design of Google F1. |
| [influxdb](https://github.com/influxdata/influxdb) | 17640 | 760 | 2013-09-26 | 49 minutes ago | Scalable datastore for metrics, events, and real-time analytics. |
| [cockroach](https://github.com/cockroachdb/cockroach) | 17220 | 722 | 2014-02-06 | 1 hour ago | Scalable, Geo-Replicated, Transactional Datastore. |
| [dgraph](https://github.com/dgraph-io/dgraph) | 11480 | 358 | 2015-08-25 | 15 hours ago | Scalable, Distributed, Low Latency, High Throughput Graph Database. |
| [bolt](https://github.com/boltdb/bolt) | 10203 | 343 | 2013-12-21 | 1 year ago | Low-level key/value database for Go. |
| [vitess](https://github.com/vitessio/vitess) | 8770 | 486 | 2013-06-28 | 13 hours ago | vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. |
| [groupcache](https://github.com/golang/groupcache) | 7912 | 456 | 2013-07-23 | 19 hours ago | Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. |
| [badger](https://github.com/dgraph-io/badger) | 6701 | 242 | 2017-01-26 | 6 hours ago | Fast key-value store in Go. |
| [pgweb](https://github.com/sosedoff/pgweb) | 6110 | 152 | 2014-10-09 | 3 weeks ago | Web-based PostgreSQL database browser. |
| [rqlite](https://github.com/rqlite/rqlite) | 4900 | 190 | 2014-08-23 | 3 weeks ago | The lightweight, distributed, relational database built on SQLite. |
| [kingshard](https://github.com/flike/kingshard) | 4806 | 399 | 2015-07-04 | 2 months ago | kingshard is a high performance proxy for MySQL powered by Golang. |
| [goleveldb](https://github.com/syndtr/goleveldb) | 3337 | 168 | 2013-01-23 | 3 days ago | Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go. |
| [orchestrator](https://github.com/github/orchestrator) | 3185 | 216 | 2016-11-30 | 1 week ago | MySQL replication topology manager & visualizer. |
| [go-cache](https://github.com/patrickmn/go-cache) | 3181 | 93 | 2012-01-02 | 2 days ago | In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. |
| [ledisdb](https://github.com/siddontang/ledisdb) | 3147 | 187 | 2014-04-30 | 2 months ago | Ledisdb is a high performance NoSQL like Redis based on LevelDB. |
| [migrate](https://github.com/golang-migrate/migrate) | 3117 | 50 | 2018-01-19 | 6 days ago | Database migrations. CLI and Golang library. |
| [bigcache](https://github.com/allegro/bigcache) | 2629 | 82 | 2016-03-23 | 2 days ago | Efficient key/value cache for gigabytes of data. |
| [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) | 2569 | 160 | 2015-01-15 | 2 weeks ago | Sync your MySQL data into Elasticsearch automatically. |
| [buntdb](https://github.com/tidwall/buntdb) | 2507 | 95 | 2016-07-20 | 3 weeks ago | Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support. |
| [squirrel](https://github.com/Masterminds/squirrel) | 2482 | 39 | 2014-01-18 | 1 week ago | Go library that helps you build SQL queries. |
| [tiedot](https://github.com/HouzuoGuo/tiedot) | 2411 | 159 | 2013-05-26 | 9 months ago | Your NoSQL database powered by Golang. |
| [xo](https://github.com/xo/xo) | 2251 | 74 | 2016-02-05 | 1 month ago | Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. |
| [prest](https://github.com/prest/prest) | 2134 | 84 | 2016-11-22 | 5 months ago | Serve a RESTful API from any PostgreSQL database. |
| [go-mysql](https://github.com/siddontang/go-mysql) | 2034 | 150 | 2014-02-21 | 1 day ago | Go toolset to handle MySQL protocol and replication. |
| [sql-migrate](https://github.com/rubenv/sql-migrate) | 1484 | 31 | 2014-09-09 | 3 days ago | Database migration tool. Allows embedding migrations into the application using go-bindata. |
| [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) | 1300 | 39 | 2018-09-30 | 4 hours ago | fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL. |
| [cache2go](https://github.com/muesli/cache2go) | 1127 | 59 | 2013-11-11 | 1 week ago | In-memory key:value cache which supports automatic invalidation based on timeouts. |
| [gcache](https://github.com/bluele/gcache) | 974 | 37 | 2015-01-25 | 4 months ago | Cache library with support for expirable Cache, LFU, LRU and ARC. |
| [nutsdb](https://github.com/xujiajun/nutsdb) | 930 | 35 | 2018-12-07 | 1 month ago | Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set. |
| [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) | 922 | 62 | 2018-04-11 | 1 month ago | CovenantSQL is a SQL database on blockchain. |
| [gendry](https://github.com/didi/gendry) | 882 | 56 | 2017-12-01 | 2 days ago | Non-invasive SQL builder and powerful data binder. |
| [diskv](https://github.com/peterbourgon/diskv) | 801 | 32 | 2012-03-22 | 6 months ago | Home-grown disk-backed key-value store. |
| [moss](https://github.com/couchbase/moss) | 734 | 79 | 2016-02-07 | 3 days ago | Moss is a simple LSM key-value storage engine written in 100% Go. |
| [goqu](https://github.com/doug-martin/goqu) | 688 | 27 | 2015-02-21 | 3 weeks ago | Idiomatic SQL builder and query library. |
| [eliasdb](https://github.com/krotik/eliasdb) | 548 | 23 | 2016-08-13 | 1 month ago | Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language. |
| [fastcache](https://github.com/VictoriaMetrics/fastcache) | 545 | 19 | 2018-11-23 | 1 month ago | fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead. |
| [dotsql](https://github.com/gchaincl/dotsql) | 458 | 21 | 2014-11-20 | 2 months ago | Go library that helps you keep sql files in one place and use them with ease. |
| [skeema](https://github.com/skeema/skeema) | 449 | 32 | 2016-11-01 | 4 days ago | Pure-SQL schema management system for MySQL, with support for sharding and external online schema change tools. |
| [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) | 440 | 24 | 2015-12-11 | 5 months ago | Powerful data retrieval methods as well as DB-agnostic query building capabilities. |
| [levigo](https://github.com/jmhodges/levigo) | 372 | 22 | 2012-01-17 | 1 week ago | Levigo is a Go wrapper for LevelDB. |
| [gormigrate](https://github.com/go-gormigrate/gormigrate) | 366 | 3 | 2016-08-31 | 3 days ago | Database schema migration helper for Gorm ORM. |
| [chproxy](https://github.com/Vertamedia/chproxy) | 328 | 23 | 2017-09-18 | 1 month ago | HTTP proxy for ClickHouse database. |
| [pudge](https://github.com/recoilme/pudge) | 235 | 7 | 2018-11-20 | 4 weeks ago | Fast and simple  key/value store written using Go's standard library. |
| [bitcask](https://github.com/prologic/bitcask) | 201 | 7 | 2019-03-12 | 1 week ago | Bitcask is an embeddable, persistent and fast key-value (KV) database written in pure Go with predictable read/write performance, low latency and high throughput thanks to the bitcask on-disk layout (LSM+WAL). |
| [sqrl](https://github.com/elgris/sqrl) | 187 | 6 | 2014-06-25 | 1 month ago | SQL query builder, fork of Squirrel with improved performance. |
| [piladb](https://github.com/fern4lvarez/piladb) | 171 | 9 | 2015-09-09 | 1 year ago | Lightweight RESTful database engine based on stack data structures. |
| [vasto](https://github.com/chrislusf/vasto) | 164 | 19 | 2018-01-16 | 7 months ago | A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption. |
| [jet](https://github.com/go-jet/jet) | 162 | 9 | 2019-03-02 | 1 week ago | Framework for writing type-safe SQL queries in Go, with ability to easily convert database query result into desired arbitrary object structure. |
| [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) | 150 | 6 | 2017-04-29 | 5 days ago | Collects small insterts and sends big requests to ClickHouse servers. |
| [scaneo](https://github.com/variadico/scaneo) | 150 | 8 | 2015-04-30 | 1 year ago | Generate Go code to convert database rows into arbitrary structs. |
| [myreplication](https://github.com/2tvenom/myreplication) | 146 | 19 | 2015-02-05 | 1 year ago | MySql binary log replication listener. Supports statement and row based replication. |
| [kivik](https://github.com/go-kivik/kivik) | 124 | 5 | 2017-02-09 | 1 week ago | Kivik provides a common Go and GopherJS client library for CouchDB, PouchDB, and similar databases. |
| [goose](https://github.com/steinbacher/goose) | 123 | 4 | 2016-03-05 | 3 years ago | Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts. |
| [darwin](https://github.com/GuiaBolso/darwin) | 91 | 2 | 2016-04-05 | 3 weeks ago | Database schema evolution library for Go. |
| [slowpoke](https://github.com/recoilme/slowpoke) | 91 | 7 | 2018-02-19 | 4 weeks ago | Key-value store with persistence. |
| [migrator](https://github.com/lopezator/migrator) | 80 | 5 | 2019-02-04 | 4 weeks ago | Dead simple Go database migration library. |
| [igor](https://github.com/galeone/igor) | 78 | 7 | 2016-03-10 | 1 year ago | Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax. |
| [golang-scribble](https://github.com/nanobox-io/golang-scribble) | 73 | 3 | 2018-06-22 | 7 months ago | Tiny flat file JSON store. |
| [octillery](https://github.com/knocknote/octillery) | 65 | 14 | 2018-11-26 | 1 month ago | Go package for sharding databases ( Supports every ORM or raw SQL ). |
| [godbal](https://github.com/xujiajun/godbal) | 48 | 4 | 2018-02-28 | 9 months ago | Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily. |
| [dbq](https://github.com/rocketlaunchr/dbq) | 46 | 4 | 2019-07-11 | 2 weeks ago | Zero boilerplate database operations for Go. |
| [couchcache](https://github.com/codingsince1985/couchcache) | 44 | 3 | 2015-04-05 | 6 months ago | RESTful caching micro-service backed by Couchbase server. |
| [bcache](https://github.com/iwanbk/bcache) | 33 | 2 | 2018-12-26 | 6 months ago | Eventually consistent distributed in-memory  cache Go library. |
| [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) | 30 | 6 | 2017-12-18 | 1 year ago | BigCache with clustering support and individual item expiration. |
| [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) | 30 | 1 | 2018-08-11 | 1 day ago | A Go package to help write migrations with go-pg/pg. |
| [dbbench](https://github.com/sj14/dbbench) | 30 | 3 | 2018-11-24 | 8 months ago | Database benchmarking tool with support for several databases and scripts. |
| [gondolier](https://github.com/emvi/gondolier) | 26 | 4 | 2017-10-18 | 5 months ago | Database migration library using struct decorators. |
| [cache](https://github.com/akyoto/cache) | 24 | 2 | 2019-05-11 | 2 days ago | In-memory key:value store with expiration time, 0 dependencies, <100 LoC, 100% coverage. |
| [pravasan](https://github.com/pravasan/pravasan) | 24 | 6 | 2015-01-04 | 10 months ago | Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc. |
| [prep](https://github.com/hexdigest/prep) | 24 | 2 | 2017-12-12 | 1 year ago | Use prepared SQL statements without changing your code. |
| [go-fixtures](https://github.com/RichardKnop/go-fixtures) | 20 | 1 | 2015-12-24 | 1 year ago | Django style fixtures for Golang's excellent built-in database/sql library. |
| [tempdb](https://github.com/rafaeljesus/tempdb) | 13 | 3 | 2017-03-18 | 1 year ago | Key-value store for temporary items. |
| [datagen](https://github.com/codingconcepts/datagen) | 10 | 2 | 2019-04-19 | 1 month ago | A fast data generator that's multi-table aware and supports multi-row DML. |
| [gorocksdb](https://github.com/kapitan-k/gorocksdb) | 10 | 1 | 2017-12-28 | 1 year ago | Gorocksdb is a wrapper for [RocksDB](https://rocksdb.org) written in Go. |
| [rwdb](https://github.com/andizzle/rwdb) | 10 | 2 | 2017-10-04 | 1 year ago | rwdb provides read replica capability for multiple database servers setup. |
| [avro](https://github.com/khezen/avro) | 7 | 1 | 2019-04-07 | 1 week ago | Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes. |
| [bucket](https://github.com/PumpkinSeed/bucket) | 4 | 3 | 2019-09-22 | 6 days ago | Optimized data structure framework for Couchbase specialized on one bucket usage. |
| [schema](https://github.com/adlio/schema) | 3 | 1 | 2019-09-25 | 3 weeks ago | Library to embed schema migrations for database/sql-compatible databases inside your Go binaries. |
| [sqlf](https://github.com/leporo/sqlf) | 2 | 1 | 2019-07-20 | 2 weeks ago | Fast SQL query builder. |
| [ormlite](https://github.com/pupizoid/ormlite) | 0 | 1 | 2018-06-28 | 1 month ago | Lightweight package containing some ORM-like features and helpers for sqlite databases. |

## Database Drivers
        
*Libraries for connecting and operating databases.*

### Relational Databases
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [mysql](https://github.com/go-sql-driver/mysql) | 8508 | 414 | 2012-12-10 | 3 days ago | MySQL driver for Go. |
| [pq](https://github.com/lib/pq) | 5384 | 157 | 2012-03-13 | 1 week ago | Pure Go Postgres driver for database/sql. |
| [go-sqlite3](https://github.com/mattn/go-sqlite3) | 3571 | 127 | 2011-11-11 | 2 weeks ago | SQLite3 driver for go that uses database/sql. |
| [pgx](https://github.com/jackc/pgx) | 2133 | 61 | 2013-03-31 | 5 days ago | PostgreSQL driver supporting features beyond those exposed by database/sql. |
| [go-mssqldb](https://github.com/denisenkom/go-mssqldb) | 1079 | 60 | 2013-12-16 | 4 days ago | Microsoft MSSQL driver for Go. |
| [go-oci8](https://github.com/mattn/go-oci8) | 417 | 36 | 2012-02-29 | 3 weeks ago | Oracle driver for go that uses database/sql. |
| [goracle](https://github.com/go-goracle/goracle) | 266 | 29 | 2015-03-25 | 1 week ago | Oracle driver for Go, using the ODPI-C driver. |
| [firebirdsql](https://github.com/nakagami/firebirdsql) | 106 | 16 | 2013-08-27 | 3 weeks ago | Firebird RDBMS SQL driver for Go. |
| [gofreetds](https://github.com/minus5/gofreetds) | 93 | 22 | 2012-12-07 | 8 months ago | Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org). |
| [go-adodb](https://github.com/mattn/go-adodb) | 91 | 13 | 2011-11-14 | 4 months ago | Microsoft ActiveX Object DataBase driver for go that uses database/sql. |
| [calcite-avatica-go](https://github.com/apache/calcite-avatica-go) | 40 | 21 | 2017-08-08 | 1 month ago | Apache Avatica/Phoenix SQL driver for database/sql. |
| [bgc](https://github.com/viant/bgc) | 12 | 10 | 2016-06-14 | 1 month ago | Datastore Connectivity for BigQuery for go. |

### NoSQL Databases
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [cayley](https://github.com/cayleygraph/cayley) | 12931 | 614 | 2014-06-06 | 1 day ago | Graph database with support for multiple backends. |
| [redis](https://github.com/go-redis/redis) | 7158 | 233 | 2012-07-25 | 4 days ago | Redis client for Golang. |
| [redigo](https://github.com/gomodule/redigo) | 6600 | 296 | 2012-04-14 | 2 weeks ago | Redigo is a Go client for the Redis database. |
| [bleve](https://github.com/blevesearch/bleve) | 6018 | 243 | 2014-04-18 | 3 days ago | Modern text indexing library for go. |
| [riot](https://github.com/go-ego/riot) | 4856 | 180 | 2017-06-21 | 2 months ago | Go Open Source, Distributed, Simple and efficient Search Engine. |
| [elastic](https://github.com/olivere/elastic) | 4428 | 167 | 2012-12-07 | 22 hours ago | Elasticsearch client for Go. |
| [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) | 3532 | 143 | 2017-02-09 | 1 hour ago | Official MongoDB driver for the Go language. |
| [go-elasticsearch](https://github.com/elastic/go-elasticsearch) | 1878 | 283 | 2017-03-28 | 8 hours ago | Official Elasticsearch client for Go. |
| [mgo](https://github.com/globalsign/mgo) | 1692 | 73 | 2017-04-13 | 2 weeks ago | (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms. |
| [rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1479 | 47 | 2013-09-12 | 2 months ago | Go language driver for RethinkDB. |
| [elastigo](https://github.com/mattbaird/elastigo) | 951 | 47 | 2012-10-12 | 8 months ago | Elasticsearch client library. |
| [elasticsql](https://github.com/cch123/elasticsql) | 442 | 27 | 2016-08-24 | 7 months ago | Convert sql to elasticsearch dsl in Go. |
| [neoism](https://github.com/jmcvetta/neoism) | 362 | 23 | 2012-07-12 | 8 months ago | Neo4j client for Golang. |
| [redeo](https://github.com/bsm/redeo) | 358 | 24 | 2014-03-06 | 2 months ago | Redis-protocol compatible TCP servers/services. |
| [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) | 312 | 37 | 2014-07-26 | 1 week ago | Aerospike client in Go language. |
| [gocb](https://github.com/couchbase/gocb) | 302 | 67 | 2015-01-16 | 1 day ago | Official Couchbase Go SDK. |
| [go-couchbase](https://github.com/couchbase/go-couchbase) | 296 | 25 | 2012-01-20 | 2 days ago | Couchbase client in Go. |
| [cachego](https://github.com/faabiosr/cachego) | 111 | 6 | 2016-10-06 | 2 months ago | Golang Cache component for multiple drivers. |
| [gokv](https://github.com/philippgille/gokv) | 110 | 6 | 2018-10-09 | 1 day ago | Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more). |
| [cachego](https://github.com/fabiorphp/cachego) | 104 | 6 | 2016-10-06 | 1 year ago | Golang Cache component for multiple drivers. |
| [go-rejson](https://github.com/nitishm/go-rejson) | 100 | 5 | 2018-04-23 | 1 month ago | Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease. |
| [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) | 72 | 6 | 2011-06-05 | 1 year ago | Neo4j REST Client in golang. |
| [skizze](https://github.com/seiflotfy/skizze) | 68 | 6 | 2016-01-17 | 3 years ago | probabilistic data-structures service and storage. |
| [dynago](https://github.com/underarmour/dynago) | 66 | 128 | 2015-05-18 | 2 years ago | Dynago is a principle of least surprise client for DynamoDB. |
| [arangolite](https://github.com/solher/arangolite) | 66 | 5 | 2015-10-05 | 6 months ago | Lightweight golang driver for ArangoDB. |
| [godis](https://github.com/piaohao/godis) | 65 | 8 | 2019-06-14 | 2 months ago | redis client implement by golang, inspired by jedis. |
| [go-couchdb](https://github.com/fjl/go-couchdb) | 51 | 6 | 2013-10-28 | 3 months ago | Yet another CouchDB HTTP API wrapper for Go. |
| [go-pilosa](https://github.com/pilosa/go-pilosa) | 33 | 15 | 2016-10-01 | 2 days ago | Go client library for Pilosa. |
| [goforestdb](https://github.com/couchbase/goforestdb) | 29 | 42 | 2014-05-14 | 2 years ago | Go bindings for ForestDB. |
| [neo4j](https://github.com/cihangir/neo4j) | 24 | 3 | 2013-05-18 | 4 years ago | Neo4j Rest API Bindings for Golang. |
| [goriak](https://github.com/zegl/goriak) | 24 | 2 | 2016-10-06 | 10 months ago | Go language driver for Riak KV. |
| [goes](https://github.com/OwnLocal/goes) | 24 | 33 | 2015-12-29 | 2 years ago | Library to interact with Elasticsearch. |
| [dsc](https://github.com/viant/dsc) | 14 | 15 | 2016-06-14 | 1 month ago | Datastore connectivity for SQL, NoSQL, structured files. |
| [xredis](https://github.com/shomali11/xredis) | 9 | 1 | 2017-06-14 | 4 months ago | Typesafe, customizable, clean & easy to use Redis client. |
| [godscache](https://github.com/defcronyke/godscache) | 6 | 1 | 2018-05-09 | 8 months ago | A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached. |
| [asc](https://github.com/viant/asc) | 4 | 10 | 2016-06-14 | 6 months ago | Datastore Connectivity for Aerospike for go. |

## Date and Time
        
*Libraries for working with dates and times.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [now](https://github.com/jinzhu/now) | 2266 | 51 | 2013-11-18 | 3 days ago | Now is a time toolkit for golang. |
| [dateparse](https://github.com/araddon/dateparse) | 936 | 15 | 2014-04-21 | 1 month ago | Parse date's without knowing format in advance. |
| [carbon](https://github.com/uniplaces/carbon) | 364 | 38 | 2016-08-03 | 9 months ago | Simple Time extension with a lot of util methods, ported from PHP Carbon library. |
| [durafmt](https://github.com/hako/durafmt) | 262 | 4 | 2016-05-21 | 2 weeks ago | Time duration formatting library for Go. |
| [timeutil](https://github.com/leekchan/timeutil) | 171 | 6 | 2015-08-02 | 8 months ago | Useful extensions (Timedelta, Strftime, ...) to the golang's time package. |
| [iso8601](https://github.com/relvacode/iso8601) | 68 | 2 | 2017-04-25 | 10 months ago | Efficiently parse ISO8601 date-times without regex. |
| [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) | 67 | 3 | 2016-02-01 | 4 weeks ago | The implementation of the Persian (Solar Hijri) Calendar in Go (golang). |
| [timespan](https://github.com/SaidinWoT/timespan) | 65 | 5 | 2014-10-07 | 7 months ago | For interacting with intervals of time, defined as a start time and a duration. |
| [date](https://github.com/rickb777/date) | 31 | 2 | 2015-11-24 | 1 month ago | Augments Time for working with dates, date ranges, time spans, periods, and time-of-day. |
| [feiertage](https://github.com/wlbr/feiertage) | 23 | 1 | 2015-11-04 | 1 week ago | Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving... |
| [goweek](https://github.com/grsmv/goweek) | 19 | 3 | 2015-11-12 | 2 years ago | Library for working with week entity in golang. |
| [go-sunrise](https://github.com/nathan-osman/go-sunrise) | 14 | 3 | 2017-06-16 | 1 year ago | Calculate the sunrise and sunset times for a given location. |
| [kair](https://github.com/GuilhermeCaruso/kair) | 12 | 1 | 2018-10-03 | 8 months ago | Date and Time - Golang Formatting Library. |
| [nulltime](https://github.com/kirillDanshin/nulltime) | 9 | 1 | 2016-03-06 | 2 years ago | Nullable `time.Time`. |
| [tuesday](https://github.com/osteele/tuesday) | 7 | 1 | 2017-08-11 | 2 years ago | Ruby-compatible Strftime function. |
| [strftime](https://github.com/awoodbeck/strftime) | 4 | 1 | 2018-02-10 | 1 year ago | C99-compatible strftime formatter. |

## Distributed Systems
        
*Packages that help with building Distributed Systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [kit](https://github.com/go-kit/kit) | 15197 | 659 | 2015-02-03 | 1 week ago | Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc. |
| [grpc-go](https://github.com/grpc/grpc-go) | 9826 | 470 | 2014-12-09 | 2 days ago | The Go language implementation of gRPC. HTTP/2 based RPC. |
| [micro](https://github.com/micro/micro) | 6876 | 307 | 2015-01-17 | 1 hour ago | Pluggable microservice toolkit and distributed systems platform. |
| [nats-server](https://github.com/nats-io/nats-server) | 6771 | 352 | 2012-10-30 | 53 minutes ago | Lightweight, high performance messaging system for microservices, IoT, and cloud native systems. |
| [gnatsd](https://github.com/nats-io/gnatsd) | 5611 | 337 | 2012-10-30 | 5 months ago | Lightweight, high performance messaging system for microservices, IoT, and cloud native systems. |
| [rpcx](https://github.com/smallnest/rpcx) | 4064 | 310 | 2016-05-18 | 1 week ago | Distributed pluggable RPC service framework like alibaba Dubbo. |
| [tendermint](https://github.com/tendermint/tendermint) | 3296 | 254 | 2014-05-15 | 59 minutes ago | High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols. |
| [torrent](https://github.com/anacrolix/torrent) | 3093 | 126 | 2015-01-09 | 19 hours ago | BitTorrent client package. |
| [raft](https://github.com/hashicorp/raft) | 2993 | 227 | 2013-11-05 | 1 week ago | Golang implementation of the Raft consensus protocol, by HashiCorp. |
| [dragonboat](https://github.com/lni/dragonboat) | 2698 | 134 | 2018-12-23 | 1 day ago | A feature complete and high performance multi-group Raft library in Go. |
| [glow](https://github.com/chrislusf/glow) | 2663 | 139 | 2015-06-14 | 1 year ago | Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go. |
| [gleam](https://github.com/chrislusf/gleam) | 2240 | 136 | 2016-08-26 | 3 weeks ago | Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed. |
| [emitter](https://github.com/emitter-io/emitter) | 2070 | 84 | 2016-10-29 | 2 days ago | High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love. |
| [krakend](https://github.com/devopsfaith/krakend) | 2010 | 78 | 2016-11-05 | 4 weeks ago | Ultra performant API Gateway framework with middlewares. |
| [hprose-golang](https://github.com/hprose/hprose-golang) | 1037 | 96 | 2014-02-14 | 4 months ago | Very newbility RPC Library, support 25+ languages now. |
| [ringpop-go](https://github.com/uber/ringpop-go) | 581 | 2372 | 2015-06-06 | 8 months ago | Scalable, fault-tolerant application-layer sharding for Go applications. |
| [gorpc](https://github.com/valyala/gorpc) | 565 | 27 | 2014-11-21 | 1 month ago | Simple, fast and scalable RPC library for high load. |
| [go-health](https://github.com/InVisionApp/go-health) | 502 | 25 | 2017-11-30 | 3 weeks ago | Library for enabling asynchronous dependency health checks in your service. |
| [rain](https://github.com/cenkalti/rain) | 358 | 12 | 2014-05-21 | 1 week ago | BitTorrent client and library. |
| [digota](https://github.com/digota/digota) | 313 | 24 | 2017-08-14 | 1 year ago | grpc ecommerce microservice. |
| [sleuth](https://github.com/ursiform/sleuth) | 304 | 7 | 2016-04-23 | 1 year ago | Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)). |
| [go-sundheit](https://github.com/AppsFlyer/go-sundheit) | 265 | 8 | 2019-04-08 | 1 week ago | A library built to provide support for defining async service health checks for golang services. |
| [go-jump](https://github.com/dgryski/go-jump) | 264 | 12 | 2014-06-16 | 1 year ago | Port of Google's "Jump" Consistent Hash function. |
| [consistent](https://github.com/buraksezer/consistent) | 218 | 10 | 2018-03-25 | 3 weeks ago | Consistent hashing with bounded loads. |
| [redis-lock](https://github.com/bsm/redis-lock) | 146 | 7 | 2014-10-10 | 4 months ago | Simplified distributed locking implementation using Redis. |
| [dht](https://github.com/anacrolix/dht) | 135 | 11 | 2016-12-14 | 6 days ago | BitTorrent Kademlia DHT implementation. |
| [jsonrpc](https://github.com/osamingo/jsonrpc) | 114 | 6 | 2016-10-28 | 1 month ago | The jsonrpc package helps implement of JSON-RPC 2.0. |
| [jsonrpc](https://github.com/ybbus/jsonrpc) | 105 | 7 | 2016-11-10 | 1 week ago | JSON-RPC 2.0 HTTP client implementation. |
| [celeriac.v1](https://github.com/svcavallar/celeriac.v1) | 55 | 3 | 2015-10-10 | 10 months ago | Library for adding support for interacting and monitoring Celery workers, tasks and events in Go. |
| [redislock](https://github.com/bsm/redislock) | 54 | 1 | 2019-06-24 | 1 month ago | Simplified distributed locking implementation using Redis. |
| [doublejump](https://github.com/edwingeng/doublejump) | 45 | 2 | 2018-06-27 | 9 months ago | A revamped Google's jump consistent hash. |
| [drmaa](https://github.com/dgruber/drmaa) | 25 | 2 | 2013-03-17 | 1 year ago | Job submission library for cluster schedulers based on the DRMAA standard. |
| [flowgraph](https://github.com/vectaport/flowgraph) | 23 | 1 | 2018-08-30 | 2 months ago | flow-based programming package. |
| [outboxer](https://github.com/italolelis/outboxer) | 11 | 0 | 2019-02-01 | 9 hours ago | Outboxer is a go library that implements the outbox pattern. |
| [dynatomic](https://github.com/tylfin/dynatomic) | 10 | 0 | 2019-02-09 | 8 months ago | A library for using DynamoDB as an atomic counter. |

## Email
        
*Libraries and tools that implement email creation and sending.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [MailHog](https://github.com/mailhog/MailHog) | 5468 | 123 | 2014-04-17 | 3 days ago | Email and SMTP testing with web and API interface. |
| [hermes](https://github.com/matcornic/hermes) | 1678 | 27 | 2017-03-26 | 2 months ago | Golang package that generates clean, responsive HTML e-mails. |
| [email](https://github.com/jordan-wright/email) | 1138 | 38 | 2013-12-13 | 1 week ago | A robust and flexible email library for Go. |
| [go-imap](https://github.com/emersion/go-imap) | 804 | 33 | 2016-04-27 | 2 days ago | IMAP library for clients and servers. |
| [sendgrid-go](https://github.com/sendgrid/sendgrid-go) | 539 | 190 | 2013-09-12 | 1 week ago | SendGrid's Go library for sending email. |
| [mailgun-go](https://github.com/mailgun/mailgun-go) | 402 | 56 | 2014-02-28 | 3 days ago | Go library for sending mail with the Mailgun API. |
| [hectane](https://github.com/hectane/hectane) | 169 | 13 | 2015-08-28 | 2 months ago | Lightweight SMTP client providing an HTTP API. |
| [douceur](https://github.com/aymerick/douceur) | 163 | 3 | 2015-04-09 | 1 year ago | CSS inliner for your HTML emails. |
| [go-message](https://github.com/emersion/go-message) | 124 | 9 | 2016-12-31 | 1 week ago | Streaming library for the Internet Message Format and mail messages. |
| [smtp](https://github.com/mailhog/smtp) | 52 | 8 | 2014-12-25 | 1 year ago | SMTP server protocol state machine. |
| [go-dkim](https://github.com/toorop/go-dkim) | 50 | 2 | 2015-04-29 | 1 week ago | DKIM library, to sign & verify email. |
| [go-premailer](https://github.com/vanng822/go-premailer) | 38 | 2 | 2015-02-17 | 5 months ago | Inline styling for HTML mail in Go. |
| [go-simple-mail](https://github.com/xhit/go-simple-mail) | 6 | 1 | 2019-09-15 | 1 day ago | Very simple package to send emails with SMTP Keep Alive and two timeouts: Connect and Send. |

## Embeddable Scripting Languages
        
*Embedding other languages inside your go code.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [otto](https://github.com/robertkrimen/otto) | 4883 | 193 | 2012-10-06 | 4 months ago | JavaScript interpreter written in Go. |
| [gopher-lua](https://github.com/yuin/gopher-lua) | 3079 | 141 | 2015-02-15 | 2 months ago | Lua 5.1 VM and compiler written in Go. |
| [go-lua](https://github.com/Shopify/go-lua) | 1719 | 266 | 2013-12-21 | 1 month ago | Port of the Lua 5.2 VM to pure Go. |
| [tengo](https://github.com/d5/tengo) | 1432 | 33 | 2019-01-09 | 3 days ago | Bytecode compiled script language for Go. |
| [anko](https://github.com/mattn/anko) | 953 | 47 | 2014-03-28 | 3 days ago | Scriptable interpreter written in Go. |
| [go-python](https://github.com/sbinet/go-python) | 946 | 44 | 2012-07-09 | 3 months ago | naive go bindings to the CPython C-API. |
| [expr](https://github.com/antonmedv/expr) | 881 | 30 | 2018-07-14 | 5 hours ago | an engine that can evaluate expressions. |
| [go-php](https://github.com/deuill/go-php) | 703 | 42 | 2015-09-18 | 1 year ago | PHP bindings for Go. |
| [go-duktape](https://github.com/olebedev/go-duktape) | 662 | 26 | 2015-01-08 | 3 months ago | Duktape JavaScript engine bindings for Go. |
| [golua](https://github.com/aarzilli/golua) | 452 | 33 | 2010-12-07 | 3 months ago | Go bindings for Lua C API. |
| [gisp](https://github.com/jcla1/gisp) | 431 | 20 | 2014-01-11 | 2 years ago | Simple LISP in Go. |
| [agora](https://github.com/mna/agora) | 323 | 22 | 2013-06-19 | 4 years ago | Dynamically typed, embeddable programming language in Go. |
| [cel-go](https://github.com/google/cel-go) | 305 | 21 | 2018-03-10 | 2 days ago | Fast, portable, non-Turing complete expression evaluation with gradual typing. |
| [gval](https://github.com/PaesslerAG/gval) | 159 | 14 | 2017-09-27 | 2 months ago | A highly customizable expression language written in Go. |
| [gentee](https://github.com/gentee/gentee) | 36 | 3 | 2018-01-14 | 10 hours ago | Embeddable scripting programming language. |
| [binder](https://github.com/alexeyco/binder) | 32 | 2 | 2017-04-03 | 1 year ago | Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua). |
| [purl](https://github.com/ian-kent/purl) | 28 | 2 | 2014-11-30 | 4 years ago | Perl 5.18.2 embedded in Go. |
| [ngaro](https://github.com/db47h/ngaro) | 19 | 1 | 2016-08-09 | 1 year ago | Embeddable Ngaro VM implementation enabling scripting in Retro. |

## Error Handling
        
*Libraries for handling errors.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [errors](https://github.com/pkg/errors) | 5205 | 105 | 2015-12-27 | 2 months ago | Package that provides simple error handling primitives. |
| [go-multierror](https://github.com/hashicorp/go-multierror) | 761 | 144 | 2014-12-16 | 4 days ago | Go (golang) package for representing a list of errors as a single error. |
| [errorx](https://github.com/joomcode/errorx) | 597 | 49 | 2018-08-17 | 1 month ago | A feature rich error package with stack traces, composition of errors and more. |
| [tracerr](https://github.com/ztrue/tracerr) | 518 | 7 | 2019-02-07 | 7 months ago | Golang errors with stack trace and source fragments. |
| [errlog](https://github.com/snwfdhmp/errlog) | 198 | 5 | 2019-02-17 | 2 months ago | Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place. |
| [emperror](https://github.com/emperror/emperror) | 76 | 2 | 2017-06-13 | 1 month ago | Error handling tools and best practices for Go libraries and applications. |
| [errors](https://github.com/emperror/errors) | 27 | 3 | 2019-07-09 | 1 month ago | Drop-in replacement for the standard library errors package and github.com/pkg/errors. Provides various error handling primitives. |
| [werr](https://github.com/txgruppi/werr) | 11 | 0 | 2015-08-04 | 3 years ago | Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called. |
| [errors](https://github.com/neuronlabs/errors) | 3 | 1 | 2019-07-26 | 2 months ago | Simple golang error handling with classification primitives. |
| [falcon](https://github.com/SonicRoshan/falcon) | 2 | 1 | 2019-09-09 | 2 weeks ago | A Simple Yet Highly Powerful Package For Error Handling. |

## Files
        
*Libraries for handling files and file systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [afero](https://github.com/spf13/afero) | 2371 | 91 | 2014-10-28 | 3 weeks ago | FileSystem Abstraction System for Go. |
| [pdfcpu](https://github.com/pdfcpu/pdfcpu) | 1146 | 35 | 2017-06-19 | 3 days ago | PDF processor. |
| [pdfcpu](https://github.com/hhrutter/pdfcpu) | 948 | 25 | 2017-06-19 | 2 months ago | PDF processor. |
| [notify](https://github.com/rjeczalik/notify) | 510 | 28 | 2014-09-09 | 3 weeks ago | File system event notification library with simple API, similar to os/signal. |
| [bigfile](https://github.com/bigfile/bigfile) | 76 | 8 | 2019-07-15 | 1 week ago | A file transfer system, support to manage files with http api, rpc call and ftp client. |
| [opc](https://github.com/qmuntal/opc) | 63 | 1 | 2018-11-06 | 1 week ago | Load Open Packaging Conventions (OPC) files for Go. |
| [skywalker](https://github.com/dixonwille/skywalker) | 51 | 3 | 2017-08-02 | 2 years ago | Package to allow one to concurrently go through a filesystem with ease. |
| [go-csv-tag](https://github.com/artonge/go-csv-tag) | 50 | 1 | 2017-06-18 | 4 days ago | Load csv file using tag. |
| [tarfs](https://github.com/posener/tarfs) | 38 | 2 | 2017-03-11 | 2 years ago | Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files. |
| [vfs](https://github.com/C2FO/vfs) | 30 | 21 | 2017-08-02 | 1 week ago | A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS. |
| [go-gtfs](https://github.com/artonge/go-gtfs) | 16 | 2 | 2017-07-09 | 4 days ago | Load gtfs files in go. |
| [checksum](https://github.com/codingsince1985/checksum) | 12 | 1 | 2014-11-05 | 1 month ago | Compute message digest, like MD5 and SHA256, for large files. |
| [afs](https://github.com/viant/afs) | 12 | 7 | 2019-08-20 | 3 days ago | Abstract File Storage (mem, scp, zip, tar, cloud: s3, gs) for Go. |
| [flop](https://github.com/homedepot/flop) | 11 | 11 | 2019-03-01 | 1 month ago | File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html). |
| [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) | 11 | 1 | 2018-10-16 | 2 weeks ago | Copy files for humans. |
| [go-exiftool](https://github.com/barasher/go-exiftool) | 9 | 1 | 2019-05-13 | 1 week ago | Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...). |
| [parquet](https://github.com/parsyl/parquet) | 4 | 3 | 2019-01-30 | 2 weeks ago | Read and write [parquet](https://parquet.apache.org) files. |

## Financial
        
*Packages for accounting and finance.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [decimal](https://github.com/shopspring/decimal) | 1740 | 61 | 2015-02-26 | 2 weeks ago | Arbitrary-precision fixed-point decimal numbers. |
| [go-money](https://github.com/Rhymond/go-money) | 655 | 16 | 2017-03-21 | 1 month ago | Implementation of Fowler's Money pattern. |
| [go-finance](https://github.com/FlashBoys/go-finance) | 538 | 27 | 2016-02-28 | 1 year ago | Comprehensive financial markets data in Go. |
| [accounting](https://github.com/leekchan/accounting) | 506 | 12 | 2015-08-10 | 3 months ago | money and currency formatting for golang. |
| [techan](https://github.com/sdcoffey/techan) | 188 | 23 | 2017-03-08 | 1 month ago | Technical analysis library with advanced market analysis and trading strategies. |
| [orderbook](https://github.com/i25959341/orderbook) | 95 | 7 | 2018-04-25 | 5 months ago | Matching Engine for Limit Order Book in Golang. |
| [ofxgo](https://github.com/aclindsa/ofxgo) | 68 | 9 | 2015-11-08 | 3 months ago | Query OFX servers and/or parse the responses (with example command-line client). |
| [vat](https://github.com/dannyvankooten/vat) | 61 | 1 | 2016-06-19 | 1 year ago | VAT number validation & EU VAT rates. |
| [transaction](https://github.com/claygod/transaction) | 58 | 8 | 2017-10-11 | 1 week ago | Embedded transactional database of accounts, running in multithreaded mode. |
| [go-finance](https://github.com/alpeb/go-finance) | 45 | 2 | 2017-06-01 | 1 year ago | Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations. |
| [currency](https://github.com/bnkamalesh/currency) | 12 | 4 | 2017-05-09 | 4 months ago | High performant & accurate currency computation package. |

## Forms
        
*Libraries for working with forms.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [nosurf](https://github.com/justinas/nosurf) | 999 | 34 | 2013-08-23 | 6 months ago | CSRF protection middleware for Go. |
| [binding](https://github.com/mholt/binding) | 756 | 30 | 2014-05-21 | 1 year ago | Binds form and JSON data from net/http Request to struct. |
| [csrf](https://github.com/gorilla/csrf) | 468 | 21 | 2015-08-03 | 2 weeks ago | CSRF protection for Go web applications & services. |
| [form](https://github.com/go-playground/form) | 363 | 11 | 2016-05-26 | 1 week ago | Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support. |
| [conform](https://github.com/leebenson/conform) | 173 | 5 | 2016-01-06 | 2 months ago | Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags. |
| [formam](https://github.com/monoculum/formam) | 132 | 3 | 2014-10-25 | 1 month ago | decode form's values into a struct. |
| [forms](https://github.com/albrow/forms) | 105 | 6 | 2014-08-08 | 2 years ago | Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files. |
| [bind](https://github.com/robfig/bind) | 24 | 2 | 2014-08-06 | 5 years ago | Bind form data to any Go values. |

## Functional
        
*Packages to support functional programming in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-underscore](https://github.com/tobyhede/go-underscore) | 1095 | 27 | 2014-07-02 | 8 months ago | Useful collection of helpfully functional Go collection utilities. |
| [fpGo](https://github.com/TeaEntityLab/fpGo) | 120 | 3 | 2018-05-24 | 1 year ago | Monad, Functional Programming features for Golang. |
| [fuego](https://github.com/seborama/fuego) | 52 | 2 | 2018-11-06 | 2 months ago | Functional Experiment in Go. |

## Game Development
        
*Awesome game development libraries.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [leaf](https://github.com/name5566/leaf) | 3205 | 318 | 2014-08-04 | 6 hours ago | Lightweight game server framework. |
| [pixel](https://github.com/faiface/pixel) | 2568 | 93 | 2016-11-19 | 4 days ago | Hand-crafted 2D game library in Go. |
| [ebiten](https://github.com/hajimehoshi/ebiten) | 2063 | 80 | 2013-06-16 | 1 day ago | dead simple 2D game library in Go. |
| [goworld](https://github.com/xiaonanln/goworld) | 1272 | 116 | 2017-06-03 | 3 months ago | Scalable game server engine, featuring space-entity framework and hot-swapping. |
| [go-sdl2](https://github.com/veandco/go-sdl2) | 1200 | 41 | 2013-06-06 | 1 month ago | Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/). |
| [engo](https://github.com/EngoEngine/engo) | 1119 | 45 | 2014-11-12 | 1 month ago | Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm. |
| [nano](https://github.com/lonng/nano) | 1072 | 56 | 2017-08-02 | 1 month ago | Lightweight, facility, high performance golang based game server framework. |
| [gonet](https://github.com/xtaci/gonet) | 1068 | 132 | 2013-04-11 | 2 years ago | Game server skeleton implemented with golang. |
| [termloop](https://github.com/JoelOtter/termloop) | 1035 | 34 | 2015-05-24 | 4 months ago | Terminal-based game engine for Go, built on top of Termbox. |
| [engine](https://github.com/g3n/engine) | 825 | 58 | 2017-03-08 | 1 week ago | Go 3D Game Engine. |
| [oak](https://github.com/oakmound/oak) | 674 | 41 | 2017-07-16 | 6 days ago | Pure Go game engine. |
| [engine](https://github.com/azul3d/engine) | 431 | 23 | 2016-02-29 | 1 week ago | 3D game engine written in Go. |
| [raylib-go](https://github.com/gen2brain/raylib-go) | 401 | 20 | 2017-01-27 | 3 weeks ago | Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming. |
| [go-astar](https://github.com/beefsack/go-astar) | 335 | 9 | 2014-05-28 | 1 year ago | Go implementation of the A\* path finding algorithm. |
| [pitaya](https://github.com/topfreegames/pitaya) | 331 | 39 | 2018-03-20 | 2 hours ago | Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. |
| [GarageEngine](https://github.com/vova616/GarageEngine) | 314 | 30 | 2012-08-07 | 2 months ago | 2d game engine written in Go working on OpenGL. |
| [go3d](https://github.com/ungerik/go3d) | 169 | 10 | 2011-06-27 | 7 months ago | Performance oriented 2D/3D math package for Go. |
| [glop](https://github.com/runningwild/glop) | 76 | 3 | 2011-04-21 | 4 years ago | Glop (Game Library Of Power) is a fairly simple cross-platform game library. |
| [go-collada](https://github.com/GlenKelley/go-collada) | 14 | 3 | 2013-09-19 | 6 years ago | Go package for working with the Collada file format. |

## Generation and Generics
        
*Tools to enhance the language with features like generics via code generation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-linq](https://github.com/ahmetb/go-linq) | 1856 | 67 | 2013-12-19 | 7 hours ago | .NET LINQ-like query methods for Go. |
| [jennifer](https://github.com/dave/jennifer) | 1355 | 23 | 2016-12-05 | 1 month ago | Generate arbitrary Go code without templates. |
| [gen](https://github.com/clipperhouse/gen) | 1054 | 34 | 2013-10-14 | 1 year ago | Code generation tool for ‘generics’-like functionality. |
| [goderive](https://github.com/awalterschulze/goderive) | 767 | 25 | 2017-02-11 | 3 months ago | Derives functions from input types. |
| [gowrap](https://github.com/hexdigest/gowrap) | 294 | 11 | 2018-09-15 | 1 month ago | Generate decorators for Go interfaces using simple templates. |
| [interfaces](https://github.com/rjeczalik/interfaces) | 196 | 6 | 2015-12-06 | 1 week ago | Command line tool for generating interface definitions. |
| [go-enum](https://github.com/abice/go-enum) | 90 | 3 | 2017-08-11 | 3 weeks ago | Code generation for enums from code comments. |
| [pkgreflect](https://github.com/ungerik/pkgreflect) | 89 | 5 | 2012-09-03 | 2 years ago | Go preprocessor for package scoped reflection. |
| [efaceconv](https://github.com/t0pep0/efaceconv) | 44 | 4 | 2016-11-18 | 2 years ago | Code generation tool for high performance conversion from interface{} to immutable type without allocations. |
| [gotype](https://github.com/wzshiming/gotype) | 23 | 4 | 2017-12-05 | 2 weeks ago | Golang source code parsing, usage like reflect package. |
| [GENERIS](https://github.com/senselogic/GENERIS) | 19 | 1 | 2019-03-11 | 2 weeks ago | Code generation tool providing generics, free-form macros, conditional compilation and HTML templating. |

## Geographic
        
*Geographic tools and servers*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [tile38](https://github.com/tidwall/tile38) | 6485 | 198 | 2016-03-05 | 1 day ago | Geolocation DB with spatial index and realtime geofencing. |
| [geo](https://github.com/golang/geo) | 942 | 76 | 2014-12-04 | 1 month ago | S2 geometry library in Go. |
| [geocache](https://github.com/melihmucuk/geocache) | 115 | 6 | 2016-06-21 | 3 years ago | In-memory cache that is suitable for geolocation based applications. |
| [osm](https://github.com/paulmach/osm) | 80 | 7 | 2016-02-02 | 3 weeks ago | Library for reading, writing and working with OpenStreetMap data and APIs. |
| [wgs84](https://github.com/wroge/wgs84) | 42 | 1 | 2019-06-09 | 1 day ago | Library for Coordinate Conversion and Transformation (ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM). |
| [geoserver](https://github.com/hishamkaram/geoserver) | 25 | 1 | 2018-03-27 | 1 month ago | geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API. |
| [gismanager](https://github.com/hishamkaram/gismanager) | 20 | 1 | 2018-09-29 | 1 year ago | Publish Your GIS Data(Vector Data) to PostGIS and Geoserver. |
| [pbf](https://github.com/maguro/pbf) | 18 | 1 | 2017-09-19 | 1 year ago | OpenStreetMap PBF golang encoder/decoder. |

## Go Compilers
        
*Tools for compiling Go to other languages.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gopherjs](https://github.com/gopherjs/gopherjs) | 8786 | 253 | 2013-08-28 | 6 days ago | Compiler from Go to JavaScript. |
| [llgo](https://github.com/go-llvm/llgo) | 1002 | 62 | 2011-11-05 | 4 years ago | LLVM-based compiler for Go. |
| [tardisgo](https://github.com/tardisgo/tardisgo) | 391 | 25 | 2014-01-08 | 2 years ago | Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler. |
| [c4go](https://github.com/Konstantin8105/c4go) | 177 | 14 | 2018-03-28 | 5 days ago | Transpile C code to Go code. |
| [f4go](https://github.com/Konstantin8105/f4go) | 20 | 3 | 2018-07-09 | 6 days ago | Transpile FORTRAN 77 code to Go code. |

## Goroutines
        
*Tools for managing and working with Goroutines.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [ants](https://github.com/panjf2000/ants) | 2433 | 93 | 2018-05-19 | 4 days ago | A high-performance goroutine pool for golang. |
| [goworker](https://github.com/benmanns/goworker) | 2288 | 76 | 2013-07-23 | 2 months ago | goworker is a Go-based background worker. |
| [tunny](https://github.com/Jeffail/tunny) | 1423 | 61 | 2014-04-03 | 3 weeks ago | Goroutine pool for golang. |
| [grpool](https://github.com/ivpusic/grpool) | 524 | 29 | 2015-07-22 | 9 months ago | Lightweight Goroutine pool. |
| [pool](https://github.com/go-playground/pool) | 512 | 13 | 2015-10-29 | 1 month ago | Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation. |
| [go-floc](https://github.com/workanator/go-floc) | 173 | 8 | 2017-07-03 | 1 year ago | Orchestrate goroutines with ease. |
| [workerpool](https://github.com/gammazero/workerpool) | 171 | 7 | 2016-05-17 | 3 weeks ago | Goroutine pool that limits the concurrency of task execution, not the number of tasks queued. |
| [go-flow](https://github.com/kamildrazkiewicz/go-flow) | 115 | 8 | 2016-09-25 | 5 months ago | Control goroutines execution order. |
| [GoSlaves](https://github.com/dgrr/GoSlaves) | 89 | 4 | 2017-09-18 | 2 months ago | Simple and Asynchronous Goroutine pool library. |
| [semaphore](https://github.com/marusama/semaphore) | 77 | 4 | 2017-11-22 | 9 months ago | Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations). |
| [semaphore](https://github.com/kamilsk/semaphore) | 76 | 1 | 2016-10-08 | 2 months ago | Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context. |
| [gpool](https://github.com/sherifabdlnaby/gpool) | 57 | 1 | 2018-12-03 | 5 months ago | manages a resizeable pool of context-aware goroutines to bound concurrency. |
| [worker-pool](https://github.com/vardius/worker-pool) | 49 | 3 | 2017-10-04 | 2 months ago | goworker is a Go simple async worker pool. |
| [breaker](https://github.com/kamilsk/breaker) | 42 | 2 | 2019-02-15 | 1 month ago | Flexible mechanism to make execution flow interruptible. |
| [cyclicbarrier](https://github.com/marusama/cyclicbarrier) | 42 | 3 | 2018-01-11 | 1 year ago | CyclicBarrier for golang. |
| [gowp](https://github.com/xxjwxc/gowp) | 35 | 2 | 2019-09-14 | 5 days ago | gowp is concurrency limiting goroutine pool. |
| [gollback](https://github.com/vardius/gollback) | 28 | 1 | 2019-05-11 | 2 months ago | asynchronous simple function utilities, for managing execution of closures and callbacks. |
| [parallel-fn](https://github.com/rafaeljesus/parallel-fn) | 26 | 3 | 2017-06-18 | 1 year ago | Run functions in parallel. |
| [async](https://github.com/StudioSol/async) | 25 | 8 | 2017-07-01 | 1 week ago | A safe way to execute functions asynchronously, recovering them in case of panic. |
| [threadpool](https://github.com/shettyh/threadpool) | 21 | 2 | 2017-09-07 | 7 months ago | Golang threadpool implementation. |
| [artifex](https://github.com/mborders/artifex) | 17 | 3 | 2018-11-01 | 1 month ago | Simple in-memory job queue for Golang using worker-based dispatching. |
| [Hunch](https://github.com/AaronJan/Hunch) | 16 | 0 | 2019-06-05 | 3 months ago | Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive. |
| [artifex](https://github.com/borderstech/artifex) | 13 | 3 | 2018-11-01 | 11 months ago | Simple in-memory job queue for Golang using worker-based dispatching. |
| [stl](https://github.com/ssgreg/stl) | 11 | 1 | 2018-06-19 | 3 weeks ago | Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism. |
| [gohive](https://github.com/loveleshsharma/gohive) | 7 | 3 | 2019-05-31 | 2 weeks ago | A highly performant and easy to use Goroutine pool for Go. |
| [go-tools](https://github.com/nikhilsaraf/go-tools) | 5 | 2 | 2018-11-14 | 7 months ago | Manage a pool of goroutines using this lightweight library with a simple API. |
| [routine](https://github.com/x-mod/routine) | 4 | 1 | 2019-03-04 | 2 months ago | go routine control with context, support: Main, Go, Pool and some useful Executors. |
| [go-trylock](https://github.com/subchen/go-trylock) | 4 | 1 | 2018-04-26 | 1 year ago | TryLock support on read-write lock for Golang. |
| [queue](https://github.com/AnikHasibul/queue) | 2 | 0 | 2018-12-21 | 5 months ago | Gives you a `sync.WaitGroup` like queue group accessibility. Helps you to throttle and limit goroutines, wait for the end of the all goroutines and much more. |

## GUI
        
*Interaction*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [ui](https://github.com/andlabs/ui) | 7162 | 376 | 2014-02-18 | 4 weeks ago | Platform-native GUI library for Go. Cross platform. |
| [fyne](https://github.com/fyne-io/fyne) | 6768 | 161 | 2018-02-05 | 1 day ago | Cross platform native GUIs designed for Go, rendered using EFL. Supports: Linux, macOS, Windows. |
| [qt](https://github.com/therecipe/qt) | 6386 | 285 | 2014-11-19 | 5 days ago | Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi). |
| [webview](https://github.com/zserge/webview) | 4920 | 187 | 2017-08-19 | 3 weeks ago | Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux). |
| [robotgo](https://github.com/go-vgo/robotgo) | 4614 | 193 | 2016-09-27 | 2 days ago | Go Native cross-platform GUI system automation. Control the mouse, keyboard and other. |
| [walk](https://github.com/lxn/walk) | 3878 | 244 | 2010-09-16 | 3 hours ago | Windows application library kit for Go. |
| [app](https://github.com/maxence-charriere/app) | 3112 | 110 | 2016-10-12 | 16 hours ago | Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. |
| [go-astilectron](https://github.com/asticode/go-astilectron) | 2816 | 128 | 2017-04-22 | 3 hours ago | Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron). |
| [go-sciter](https://github.com/sciter-sdk/go-sciter) | 1523 | 118 | 2015-10-15 | 1 month ago | Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform. |
| [systray](https://github.com/getlantern/systray) | 854 | 43 | 2014-11-12 | 1 month ago | Cross platform Go library to place an icon and menu in the notification area. |
| [gotk3](https://github.com/gotk3/gotk3) | 833 | 46 | 2015-08-13 | 21 hours ago | Go bindings for GTK3. |
| [gosx-notifier](https://github.com/deckarep/gosx-notifier) | 493 | 14 | 2013-11-25 | 6 months ago | OSX Desktop Notifications library for Go. |
| [gowd](https://github.com/dtylman/gowd) | 220 | 25 | 2017-03-29 | 4 months ago | Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform. |
| [trayhost](https://github.com/shurcooL/trayhost) | 163 | 4 | 2014-04-25 | 2 days ago | Cross-platform Go library to place an icon in the host operating system's taskbar. |
| [go-appindicator](https://github.com/dawidd6/go-appindicator) | 4 | 2 | 2019-05-04 | 1 month ago | Go bindings for libappindicator3 C library. |
| [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) | 1 | 1 | 2019-03-30 | 4 months ago | OSX Sleep/Wake notifications in golang. |
| [activity-tracker](https://github.com/prashantgupta24/activity-tracker) | 1 | 1 | 2019-03-13 | 4 months ago | OSX library to notify about any (pluggable) activity on your machine. |

## Hardware
        
*Libraries, tools, and tutorials for interacting with hardware.*

## Images
        
*Libraries for manipulating images.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [imaging](https://github.com/disintegration/imaging) | 2728 | 72 | 2012-12-07 | 2 months ago | Simple Go image processing package. |
| [imaginary](https://github.com/h2non/imaginary) | 2698 | 75 | 2015-03-05 | 1 week ago | Fast and simple HTTP microservice for image resizing. |
| [gocv](https://github.com/hybridgroup/gocv) | 2686 | 127 | 2017-09-19 | 2 days ago | Go package for computer vision using OpenCV 3.3+. |
| [bild](https://github.com/anthonynsimon/bild) | 2656 | 59 | 2016-08-01 | 1 week ago | Collection of image processing algorithms in pure Go. |
| [ln](https://github.com/fogleman/ln) | 2537 | 87 | 2016-01-10 | 3 months ago | 3D line art rendering in Go. |
| [resize](https://github.com/nfnt/resize) | 2179 | 79 | 2012-08-03 | 1 year ago | Image resizing for Go with common interpolation methods. |
| [gg](https://github.com/fogleman/gg) | 1999 | 72 | 2016-02-19 | 2 months ago | 2D rendering in pure Go. |
| [pt](https://github.com/fogleman/pt) | 1793 | 58 | 2015-01-24 | 7 months ago | Path tracing engine written in Go. |
| [svgo](https://github.com/ajstarks/svgo) | 1368 | 50 | 2010-03-06 | 3 weeks ago | Go Language Library for SVG generation. |
| [smartcrop](https://github.com/muesli/smartcrop) | 1281 | 30 | 2014-04-08 | 3 weeks ago | Finds good crops for arbitrary images and crop sizes. |
| [gift](https://github.com/disintegration/gift) | 1257 | 52 | 2014-07-13 | 2 months ago | Package of image processing filters. |
| [go-opencv](https://github.com/go-opencv/go-opencv) | 1120 | 70 | 2013-12-09 | 5 months ago | Go bindings for OpenCV. |
| [picfit](https://github.com/thoas/picfit) | 1117 | 47 | 2014-12-07 | 1 month ago | An image resizing server written in Go. |
| [imagick](https://github.com/gographics/imagick) | 1021 | 56 | 2013-05-01 | 4 weeks ago | Go binding to ImageMagick's MagickWand C API. |
| [geopattern](https://github.com/pravj/geopattern) | 1019 | 21 | 2014-10-23 | 9 months ago | Create beautiful generative image patterns from a string. |
| [bimg](https://github.com/h2non/bimg) | 836 | 36 | 2015-03-17 | 2 days ago | Small package for fast and efficient image processing using libvips. |
| [stegify](https://github.com/DimitarPetrov/stegify) | 569 | 15 | 2018-11-30 | 1 week ago | Go tool for LSB steganography, capable of hiding any file within an image. |
| [mort](https://github.com/aldor007/mort) | 373 | 19 | 2017-11-19 | 5 days ago | Storage and image processing server written in Go. |
| [canvas](https://github.com/tdewolff/canvas) | 323 | 12 | 2017-05-21 | 1 day ago | Vector graphics to PDF, SVG or rasterized image. |
| [image2ascii](https://github.com/qeesung/image2ascii) | 321 | 7 | 2018-10-20 | 11 months ago | Convert image to ASCII. |
| [govatar](https://github.com/o1egl/govatar) | 320 | 5 | 2016-01-18 | 1 month ago | Library and CMD tool for generating funny avatars. |
| [go-nude](https://github.com/koyachi/go-nude) | 291 | 15 | 2014-05-02 | 11 months ago | Nudity detection with Go. |
| [goimagehash](https://github.com/corona10/goimagehash) | 246 | 8 | 2017-07-29 | 4 weeks ago | Go Perceptual image hashing package. |
| [rez](https://github.com/bamiaux/rez) | 192 | 9 | 2014-01-17 | 2 years ago | Image resizing in pure Go and SIMD. |
| [img](https://github.com/hawx/img) | 131 | 5 | 2012-07-29 | 4 years ago | Selection of image manipulation tools. |
| [darkroom](https://github.com/gojek/darkroom) | 92 | 7 | 2019-07-01 | 6 days ago | An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency. |
| [go-cairo](https://github.com/ungerik/go-cairo) | 88 | 6 | 2012-08-23 | 2 weeks ago | Go binding for the cairo graphics library. |
| [mergi](https://github.com/noelyahan/mergi) | 81 | 2 | 2018-09-24 | 5 months ago | Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate). |
| [go-gd](https://github.com/bolknote/go-gd) | 51 | 4 | 2011-05-12 | 1 year ago | Go binding for GD library. |
| [gltf](https://github.com/qmuntal/gltf) | 48 | 1 | 2019-01-16 | 3 weeks ago | Efficient and robust glTF 2.0 reader, writer and validator. |
| [cameron](https://github.com/aofei/cameron) | 35 | 1 | 2018-05-06 | 1 month ago | An avatar generator for Go. |
| [steganography](https://github.com/auyer/steganography) | 32 | 3 | 2018-05-22 | 2 weeks ago | Pure Go Library for LSB steganography. |
| [goimghdr](https://github.com/corona10/goimghdr) | 31 | 1 | 2018-02-25 | 4 months ago | The imghdr module determines the type of image contained in a file for Go. |
| [tga](https://github.com/ftrvxmtrx/tga) | 25 | 3 | 2012-10-08 | 4 years ago | Package tga is a TARGA image format decoder/encoder. |
| [go-webcolors](https://github.com/jyotiska/go-webcolors) | 24 | 1 | 2014-04-24 | 4 years ago | Port of webcolors library from Python to Go. |
| [mpo](https://github.com/donatj/mpo) | 6 | 1 | 2015-04-15 | 9 months ago | Decoder and conversion tool for MPO 3D Photos. |

## IoT
        
*Libraries for programming devices of the IoT.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [flogo](https://github.com/TIBCOSoftware/flogo) | 1217 | 127 | 2016-07-10 | 3 days ago | Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. |
| [gatt](https://github.com/paypal/gatt) | 835 | 56 | 2014-04-23 | 5 months ago | Gatt is a Go package for building Bluetooth Low Energy peripherals. |
| [mainflux](https://github.com/mainflux/mainflux) | 671 | 72 | 2015-07-07 | 5 hours ago | Industrial IoT Messaging and Device Management Server. |
| [devices](https://github.com/goiot/devices) | 227 | 17 | 2016-05-30 | 3 years ago | Suite of libraries for IoT devices, experimental for x/exp/io. |
| [sensorbee](https://github.com/sensorbee/sensorbee) | 184 | 17 | 2016-02-19 | 1 year ago | Lightweight stream processing engine for IoT. |
| [connectordb](https://github.com/connectordb/connectordb) | 179 | 19 | 2015-01-17 | 8 months ago | Open-Source Platform for Quantified Self & IoT. |
| [huego](https://github.com/amimof/huego) | 118 | 1 | 2017-05-16 | 1 day ago | An extensive Philips Hue client library for Go. |
| [eywa](https://github.com/xcodersun/eywa) | 40 | 8 | 2016-02-21 | 2 years ago | Project Eywa is essentially a connection manager that keeps track of connected devices. |

## Job Scheduler
        
*Libraries for scheduling jobs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gron](https://github.com/roylee0704/gron) | 656 | 16 | 2016-06-04 | 8 months ago | Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly. |
| [jobrunner](https://github.com/bamzi/jobrunner) | 614 | 17 | 2015-10-21 | 2 weeks ago | Smart and featureful cron job scheduler with job queuing and live monitoring built in. |
| [jobs](https://github.com/albrow/jobs) | 457 | 18 | 2015-02-10 | 1 year ago | Persistent and flexible background jobs library. |
| [scheduler](https://github.com/carlescere/scheduler) | 303 | 16 | 2015-02-04 | 1 year ago | Cronjobs scheduling made easy. |
| [go-cron](https://github.com/rk/go-cron) | 177 | 9 | 2011-04-15 | 4 years ago | Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. |
| [clockwork](https://github.com/whiteShtef/clockwork) | 88 | 3 | 2018-04-24 | 6 months ago | Simple and intuitive job scheduling library in Go. |
| [leprechaun](https://github.com/kilgaloon/leprechaun) | 46 | 5 | 2018-04-08 | 3 weeks ago | Job scheduler that supports webhooks, crons and classic scheduling. |

## JSON
        
*Libraries for working with JSON.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gjson](https://github.com/tidwall/gjson) | 5284 | 128 | 2016-08-11 | 2 months ago | Get a JSON value with one line of code. |
| [gojson](https://github.com/ChimeraCoder/gojson) | 2088 | 46 | 2012-12-28 | 4 months ago | Automatically generate Go (golang) struct definitions from example JSON. |
| [gojq](https://github.com/elgs/gojq) | 140 | 4 | 2015-12-30 | 1 year ago | JSON query in Golang. |
| [kazaam](https://github.com/qntfy/kazaam) | 139 | 20 | 2016-07-19 | 1 month ago | API for arbitrary transformation of JSON documents. |
| [jsongo](https://github.com/ricardolonga/jsongo) | 93 | 1 | 2015-08-08 | 2 years ago | Fluent API to make it easier to create Json objects. |
| [gjo](https://github.com/skanehira/gjo) | 63 | 7 | 2019-02-23 | 1 month ago | Small utility to create JSON objects. |
| [jsonf](https://github.com/miolini/jsonf) | 56 | 3 | 2015-05-25 | 3 years ago | Console tool for highlighted formatting and struct query fetching JSON. |
| [jaydiff](https://github.com/yazgazan/jaydiff) | 56 | 1 | 2017-04-25 | 3 weeks ago | JSON diff utility written in Go. |
| [jettison](https://github.com/wI2L/jettison) | 45 | 3 | 2019-08-30 | 5 days ago | High performance, reflection-less JSON encoder for Go. |
| [mp](https://github.com/sanbornm/mp) | 33 | 2 | 2014-06-16 | 3 years ago | Simple cli email parser. It currently takes stdin and outputs JSON. |
| [go-respond](https://github.com/nicklaw5/go-respond) | 27 | 1 | 2017-03-13 | 3 months ago | Go package for handling common HTTP JSON responses. |
| [json2go](https://github.com/m-zajac/json2go) | 24 | 2 | 2017-06-11 | 1 month ago | Advanced JSON to Go struct conversion. Provides package that can parse multiple JSON documents and create struct to fit them all. |
| [ajson](https://github.com/spyzhov/ajson) | 17 | 0 | 2019-03-08 | 4 months ago | Abstract JSON for golang with JSONPath support. |
| [jsonhal](https://github.com/RichardKnop/jsonhal) | 9 | 2 | 2016-01-15 | 1 year ago | Simple Go package to make custom structs marshal into HAL compatible JSON responses. |
| [go-jsonerror](https://github.com/ddymko/go-jsonerror) | 8 | 1 | 2018-10-18 | 2 weeks ago | Go-JsonError is ment to allow us to easily create json response errors that follow the JsonApi spec. |
| [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) | 6 | 1 | 2016-07-08 | 2 years ago | Go bindings based on the JSON API errors reference. |

## Logging
        
*Libraries for generating and working with log files.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [logrus](https://github.com/sirupsen/logrus) | 12790 | 286 | 2013-10-17 | 7 hours ago | Structured logger for Go. |
| [zap](https://github.com/uber-go/zap) | 8069 | 217 | 2016-02-19 | 4 days ago | Fast, structured, leveled logging in Go. |
| [go-spew](https://github.com/davecgh/go-spew) | 3459 | 56 | 2013-01-09 | 1 week ago | Implements a deep pretty printer for Go data structures to aid in debugging. |
| [zerolog](https://github.com/rs/zerolog) | 2518 | 45 | 2017-05-12 | 3 days ago | Zero-allocation JSON logger. |
| [glog](https://github.com/golang/glog) | 2398 | 85 | 2013-07-16 | 1 month ago | Leveled execution logs for Go. |
| [tail](https://github.com/hpcloud/tail) | 1615 | 98 | 2013-02-05 | 2 weeks ago | Go package striving to emulate the features of the BSD tail program. |
| [lumberjack](https://github.com/natefinch/lumberjack) | 1557 | 44 | 2014-06-14 | 1 month ago | Simple rolling logger, implements io.WriteCloser. |
| [seelog](https://github.com/cihub/seelog) | 1387 | 88 | 2011-11-17 | 7 months ago | Logging functionality with flexible dispatching, filtering, and formatting. |
| [log15](https://github.com/inconshreveable/log15) | 935 | 25 | 2014-05-20 | 2 months ago | Simple, powerful logging for Go. |
| [log](https://github.com/apex/log) | 752 | 34 | 2015-12-22 | 2 months ago | Structured logging package for Go. |
| [onelog](https://github.com/francoispqt/onelog) | 340 | 9 | 2018-05-06 | 7 months ago | Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenario. Also, it is one of the logger with the lowest allocation. |
| [logxi](https://github.com/mgutz/logxi) | 339 | 10 | 2015-03-02 | 4 months ago | 12-factor app logger that is fast and makes you happy. |
| [log](https://github.com/go-playground/log) | 269 | 10 | 2016-02-08 | 3 weeks ago | Simple, configurable and scalable Structured Logging for Go. |
| [logutils](https://github.com/hashicorp/logutils) | 259 | 148 | 2013-10-09 | 1 year ago | Utilities for slightly better logging in Go (Golang) extending the standard logger. |
| [go-logger](https://github.com/apsdehal/go-logger) | 241 | 7 | 2014-09-26 | 5 months ago | Simple logger of Go Programs, with level handlers. |
| [logger](https://github.com/azer/logger) | 135 | 4 | 2014-09-30 | 4 months ago | Minimalistic logging library for Go. |
| [xlog](https://github.com/rs/xlog) | 128 | 8 | 2015-10-22 | 4 months ago | Structured logger for `net/context` aware HTTP handlers with flexible dispatching. |
| [rollingwriter](https://github.com/arthurkiller/rollingwriter) | 110 | 7 | 2017-02-12 | 1 month ago | RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation. |
| [ozzo-log](https://github.com/go-ozzo/ozzo-log) | 108 | 9 | 2015-10-23 | 1 year ago | High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail). |
| [logvoyage](https://github.com/firstrow/logvoyage) | 83 | 5 | 2015-03-29 | 2 years ago | Full-featured logging saas written in golang. |
| [rollingWriter](https://github.com/arthurkiller/rollingWriter) | 65 | 6 | 2017-02-12 | 6 months ago | RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation. |
| [glg](https://github.com/kpango/glg) | 56 | 5 | 2017-06-21 | 1 week ago | glg is simple and fast leveled logging library for Go. |
| [log](https://github.com/alexcesaro/log) | 42 | 5 | 2014-04-19 | 4 years ago | Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs. |
| [gologger](https://github.com/sadlil/gologger) | 39 | 6 | 2015-09-02 | 1 year ago | Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch. |
| [go-log](https://github.com/ian-kent/go-log) | 34 | 2 | 2014-05-02 | 1 year ago | Log4j implementation in Go. |
| [logex](https://github.com/chzyer/logex) | 32 | 7 | 2014-10-10 | 2 years ago | Golang log lib, supports tracking and level, wrap by standard log lib. |
| [logrusly](https://github.com/sebest/logrusly) | 25 | 5 | 2014-09-12 | 1 year ago | [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/). |
| [log](https://github.com/teris-io/log) | 23 | 1 | 2017-10-29 | 1 year ago | Structured log interface for Go cleanly separates logging facade from its implementation. |
| [go-log](https://github.com/siddontang/go-log) | 23 | 4 | 2014-05-18 | 8 months ago | Log lib supports level and multi handlers. |
| [journald](https://github.com/ssgreg/journald) | 21 | 2 | 2017-08-23 | 10 months ago | Go implementation of systemd Journal's native API for logging. |
| [go-cronowriter](https://github.com/utahta/go-cronowriter) | 20 | 1 | 2017-02-04 | 1 month ago | Simple writer that rotate log files automatically based on current date and time, like cronolog. |
| [distillog](https://github.com/amoghe/distillog) | 20 | 1 | 2015-10-13 | 1 year ago | distilled levelled logging (think of it as stdlib + log levels). |
| [mlog](https://github.com/jbrodriguez/mlog) | 19 | 1 | 2014-10-20 | 1 year ago | Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output. |
| [gomol](https://github.com/aphistic/gomol) | 15 | 2 | 2015-08-30 | 7 months ago | Multiple-output, structured logging for Go with extensible logging outputs. |
| [logdump](https://github.com/ewwwwwqm/logdump) | 9 | 2 | 2017-01-13 | 1 year ago | Package for multi-level logging. |
| [go-log](https://github.com/subchen/go-log) | 9 | 2 | 2017-05-07 | 1 year ago | Simple and configurable Logging in Go, with level, formatters and writers. |
| [glo](https://github.com/lajosbencz/glo) | 8 | 1 | 2019-01-20 | 9 months ago | PHP Monolog inspired logging facility with identical severity levels. |
| [logmatic](https://github.com/borderstech/logmatic) | 6 | 1 | 2018-11-07 | 9 months ago | Colorized logger for Golang with dynamic log level configuration. |
| [xlog](https://github.com/xfxdev/xlog) | 6 | 1 | 2016-05-06 | 9 months ago | Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format. |
| [logrusiowriter](https://github.com/cabify/logrusiowriter) | 5 | 76 | 2019-08-09 | 3 weeks ago | `io.Writer` implementation using [logrus](https://github.com/sirupsen/logrus) logger. |
| [log](https://github.com/aerogo/log) | 5 | 1 | 2017-06-10 | 2 days ago | An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection). |
| [logo](https://github.com/mbndr/logo) | 5 | 1 | 2017-02-08 | 2 years ago | Golang logger to different configurable writers. |

## Machine Learning
        
*Libraries for Machine Learning.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [golearn](https://github.com/sjwhitworth/golearn) | 6771 | 429 | 2013-12-26 | 3 months ago | General Machine Learning library for Go. |
| [gorgonia](https://github.com/gorgonia/gorgonia) | 2856 | 171 | 2016-09-15 | 16 hours ago | graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. |
| [tfgo](https://github.com/galeone/tfgo) | 1237 | 48 | 2017-05-23 | 5 months ago | Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python. |
| [goml](https://github.com/cdipaolo/goml) | 1034 | 73 | 2015-06-27 | 6 months ago | On-line Machine Learning in Go. |
| [gosseract](https://github.com/otiai10/gosseract) | 950 | 34 | 2013-10-11 | 1 month ago | Go package for OCR (Optical Character Recognition), by using Tesseract C++ library. |
| [CloudForest](https://github.com/ryanbressler/CloudForest) | 653 | 44 | 2012-10-23 | 11 months ago | Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go. |
| [eaopt](https://github.com/MaxHalford/eaopt) | 639 | 28 | 2016-01-31 | 6 days ago | An evolutionary optimization library. |
| [bayesian](https://github.com/jbrukh/bayesian) | 638 | 31 | 2011-11-23 | 7 months ago | Naive Bayesian Classification for Golang. |
| [gorse](https://github.com/zhenghaoz/gorse) | 573 | 30 | 2018-08-14 | 1 day ago | A High Performance Recommender System Package based on Collaborative Filtering for Go. |
| [gobrain](https://github.com/goml/gobrain) | 399 | 23 | 2014-04-29 | 2 days ago | Neural Networks written in go. |
| [regommend](https://github.com/muesli/regommend) | 255 | 15 | 2014-02-06 | 2 months ago | Recommendation & collaborative filtering engine. |
| [ocrserver](https://github.com/otiai10/ocrserver) | 246 | 11 | 2015-11-15 | 7 months ago | A simple OCR API server, seriously easy to be deployed by Docker and Heroku. |
| [go-deep](https://github.com/patrikeh/go-deep) | 234 | 13 | 2017-12-09 | 1 year ago | A feature-rich neural network library in Go. |
| [onnx-go](https://github.com/owulveryck/onnx-go) | 215 | 8 | 2018-08-28 | 3 days ago | Go Interface to Open Neural Network Exchange (ONNX). |
| [go-galib](https://github.com/thoj/go-galib) | 173 | 14 | 2009-11-30 | 3 years ago | Genetic Algorithms library written in Go / golang. |
| [goRecommend](https://github.com/timkaye11/goRecommend) | 148 | 10 | 2014-07-16 | 5 years ago | Recommendation Algorithms library written in Go. |
| [shield](https://github.com/eaigner/shield) | 124 | 10 | 2013-04-11 | 3 years ago | Bayesian text classifier with flexible tokenizers and storage backends for Go. |
| [go-fann](https://github.com/white-pony/go-fann) | 101 | 7 | 2011-03-11 | 4 years ago | Go bindings for Fast Artificial Neural Networks(FANN) library. |
| [go-fann](https://github.com/vksnk/go-fann) | 101 | 7 | 2011-03-11 | 4 years ago | Go bindings for Fast Artificial Neural Networks(FANN) library. |
| [goptuna](https://github.com/c-bata/goptuna) | 95 | 2 | 2019-07-24 | 17 hours ago | Bayesian optimization framework for black-box functions written in Go. Everything will be optimized. |
| [goga](https://github.com/tomcraven/goga) | 78 | 7 | 2015-10-20 | 2 years ago | Genetic algorithm library for Go. |
| [libsvm](https://github.com/datastream/libsvm) | 63 | 10 | 2012-07-31 | 3 years ago | libsvm golang version derived work based on LIBSVM 3.14. |
| [neural-go](https://github.com/schuyler/neural-go) | 62 | 2 | 2011-10-17 | 5 years ago | Multilayer perceptron network implemented in Go, with training via backpropagation. |
| [go-pr](https://github.com/daviddengcn/go-pr) | 57 | 6 | 2013-06-07 | 6 years ago | Pattern recognition package in Go lang. |
| [neat](https://github.com/jinyeom/neat) | 56 | 11 | 2016-11-17 | 1 year ago | Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT). |
| [golinear](https://github.com/danieldk/golinear) | 39 | 5 | 2013-04-05 | 1 year ago | liblinear bindings for Go. |
| [goscore](https://github.com/asafschers/goscore) | 37 | 2 | 2017-08-19 | 2 months ago | Go Scoring API for PMML. |
| [fonet](https://github.com/Fontinalis/fonet) | 33 | 4 | 2017-10-03 | 1 month ago | A Deep Neural Network library written in Go. |
| [godist](https://github.com/e-dard/godist) | 24 | 2 | 2014-09-05 | 4 years ago | Various probability distributions, and associated methods. |
| [Varis](https://github.com/Xamber/Varis) | 23 | 5 | 2017-10-10 | 1 year ago | Golang Neural Network. |
| [go-cluster](https://github.com/e-XpertSolutions/go-cluster) | 22 | 5 | 2017-10-04 | 1 year ago | Go implementation of the k-modes and k-prototypes clustering algorithms. |
| [probab](https://github.com/ThePaw/probab) | 11 | 1 | 2015-09-14 | 4 years ago | Probability distribution functions. Bayesian inference. Written in pure Go. |
| [evoli](https://github.com/khezen/evoli) | 8 | 3 | 2015-06-12 | 5 months ago | Genetic Algorithm and Particle Swarm Optimization library. |
| [gomind](https://github.com/surenderthakran/gomind) | 6 | 2 | 2017-10-19 | 1 year ago | A simplistic Neural Network Library in Go. |
| [mlgo](https://github.com/NullHypothesis/mlgo) | 5 | 2 | 2015-12-07 | 3 years ago | This project aims to provide minimalistic machine learning algorithms in Go. |

## Messaging
        
*Libraries that implement messaging systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [sarama](https://github.com/Shopify/sarama) | 4958 | 381 | 2013-07-06 | 1 day ago | Go library for Apache Kafka. |
| [gorush](https://github.com/appleboy/gorush) | 3875 | 176 | 2016-03-22 | 1 day ago | Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm). |
| [centrifugo](https://github.com/centrifugal/centrifugo) | 3835 | 180 | 2015-04-01 | 1 day ago | Real-time messaging (Websockets or SockJS) server in Go. |
| [machinery](https://github.com/RichardKnop/machinery) | 3611 | 132 | 2015-04-06 | 4 days ago | Asynchronous task queue/job queue based on distributed message passing. |
| [go-socket.io](https://github.com/googollee/go-socket.io) | 3055 | 128 | 2013-07-13 | 5 days ago | socket.io library for golang, a realtime application framework. |
| [nats.go](https://github.com/nats-io/nats.go) | 2503 | 161 | 2012-08-15 | 1 week ago | Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library. |
| [go-nats](https://github.com/nats-io/go-nats) | 2247 | 159 | 2012-08-15 | 5 months ago | Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library. |
| [apns2](https://github.com/sideshow/apns2) | 2107 | 74 | 2016-01-05 | 2 weeks ago | HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps. |
| [benthos](https://github.com/Jeffail/benthos) | 2099 | 59 | 2016-03-22 | 1 day ago | A message streaming bridge between a range of protocols. |
| [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) | 1853 | 238 | 2013-12-27 | 2 years ago | gopush-cluster is a go push server cluster. |
| [mercure](https://github.com/dunglas/mercure) | 1702 | 56 | 2018-07-14 | 1 hour ago | Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events). |
| [melody](https://github.com/olahol/melody) | 1633 | 52 | 2015-05-14 | 2 months ago | Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling. |
| [mangos-v1](https://github.com/nanomsg/mangos-v1) | 1547 | 86 | 2014-10-25 | 2 months ago | Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability. |
| [go-nsq](https://github.com/nsqio/go-nsq) | 1520 | 60 | 2013-08-29 | 2 days ago | the official Go package for NSQ. |
| [uniqush-push](https://github.com/uniqush/uniqush-push) | 1114 | 74 | 2011-08-29 | 6 days ago | Redis backed unified push service for server-side notifications to mobile devices. |
| [zmq4](https://github.com/pebbe/zmq4) | 801 | 49 | 2013-10-18 | 3 weeks ago | Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2). |
| [gollum](https://github.com/trivago/gollum) | 789 | 37 | 2015-06-21 | 1 month ago | A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations. |
| [Beaver](https://github.com/Clivern/Beaver) | 753 | 21 | 2018-10-21 | 2 months ago | A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps. |
| [EventBus](https://github.com/asaskevich/EventBus) | 593 | 28 | 2014-12-20 | 3 weeks ago | The lightweight event bus with async compatibility. |
| [golongpoll](https://github.com/jcuga/golongpoll) | 436 | 23 | 2015-11-02 | 8 months ago | HTTP longpoll server library that makes web pub-sub simple. |
| [dbus](https://github.com/godbus/dbus) | 377 | 15 | 2014-03-28 | 2 weeks ago | Native Go bindings for D-Bus. |
| [emitter](https://github.com/olebedev/emitter) | 322 | 9 | 2015-11-11 | 9 months ago | Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins. |
| [glue](https://github.com/desertbit/glue) | 322 | 14 | 2015-06-07 | 4 months ago | Robust Go and Javascript Socket Library (Alternative to Socket.io). |
| [pubsub](https://github.com/cskr/pubsub) | 292 | 8 | 2012-04-01 | 6 months ago | Simple pubsub package for go. |
| [guble](https://github.com/smancke/guble) | 140 | 13 | 2015-11-16 | 2 years ago | Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence. |
| [bus](https://github.com/mustafaturan/bus) | 124 | 4 | 2019-04-27 | 2 weeks ago | Minimalist message bus implementation for internal communication. |
| [oplog](https://github.com/dailymotion/oplog) | 96 | 93 | 2014-11-06 | 4 years ago | Generic oplog/replication system for REST APIs. |
| [rabtap](https://github.com/jandelgado/rabtap) | 84 | 8 | 2017-11-11 | 3 weeks ago | RabbitMQ swiss army knife cli app. |
| [message-bus](https://github.com/vardius/message-bus) | 70 | 4 | 2017-10-04 | 2 months ago | messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD. |
| [rabbus](https://github.com/rafaeljesus/rabbus) | 65 | 8 | 2017-05-07 | 3 months ago | A tiny wrapper over amqp exchanges and queues. |
| [drone-line](https://github.com/appleboy/drone-line) | 62 | 5 | 2016-09-13 | 2 weeks ago | Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI. |
| [RapidMQ](https://github.com/sybrexsys/RapidMQ) | 56 | 4 | 2016-10-05 | 1 year ago | RapidMQ is a lightweight and reliable library for managing of the local messages queue. |
| [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) | 54 | 8 | 2017-01-16 | 1 year ago | A tiny wrapper around NSQ topic and channel. |
| [go-notify](https://github.com/TheCreeper/go-notify) | 47 | 2 | 2015-03-02 | 8 months ago | Native implementation of the freedesktop notification spec. |
| [goose](https://github.com/ian-kent/goose) | 34 | 1 | 2014-12-21 | 4 years ago | Server Sent Events in Go. |
| [commander](https://github.com/jeroenrinzema/commander) | 31 | 1 | 2018-04-20 | 1 day ago | A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka. |
| [event](https://github.com/agoalofalife/event) | 29 | 3 | 2017-07-02 | 1 year ago | Implementation of the pattern observer. |
| [hub](https://github.com/leandro-lugaresi/hub) | 25 | 1 | 2018-04-14 | 1 year ago | A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges. |
| [go-vitotrol](https://github.com/maxatome/go-vitotrol) | 12 | 6 | 2016-11-04 | 5 months ago | Client library to Viessmann Vitotrol web service. |
| [gaurun-client](https://github.com/osamingo/gaurun-client) | 8 | 1 | 2017-06-29 | 5 months ago | Gaurun Client written in Go. |
| [jazz](https://github.com/socifi/jazz) | 6 | 3 | 2018-10-22 | 7 months ago | A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages. |
| [redisqueue](https://github.com/robinjoseph08/redisqueue) | 5 | 2 | 2019-07-07 | 2 months ago | redisqueue provides a producer and consumer of a queue that uses Redis streams. |
| [rmqconn](https://github.com/sbabiv/rmqconn) | 0 | 0 | 2019-01-15 | 4 months ago | RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed. |

## Microsoft Office
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [unioffice](https://github.com/unidoc/unioffice) | 1876 | 59 | 2017-08-29 | 4 days ago | Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents. |

### Microsoft Excel
        
*Libraries for working with Microsoft Excel.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [excelize](https://github.com/360EntSecGroup-Skylar/excelize) | 4922 | 154 | 2016-08-29 | 1 day ago | Golang library for reading and writing Microsoft Excel™ (XLSX) files. |
| [xlsx](https://github.com/tealeg/xlsx) | 3603 | 183 | 2011-06-28 | 11 hours ago | Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs. |
| [xlsx](https://github.com/plandem/xlsx) | 87 | 9 | 2017-08-27 | 2 weeks ago | Fast and safe way to read/update your existing Microsoft Excel files in Go programs. |
| [xlsx2go](https://github.com/plandem/xlsx2go) | 86 | 9 | 2017-08-27 | 4 weeks ago | Fast and safe way to read/update your existing Microsoft Excel files in Go programs. |
| [go-excel](https://github.com/szyhf/go-excel) | 53 | 3 | 2017-09-03 | 2 months ago | A simple and light reader to read a relate-db-like excel as a table. |
| [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) | 14 | 1 | 2017-03-13 | 1 year ago | Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files. |

## Miscellaneous
        

### Dependency Injection
        
*Libraries for working with dependency injection.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [dig](https://github.com/uber-go/dig) | 1004 | 38 | 2017-03-22 | 4 months ago | A reflection based dependency injection toolkit for Go. |
| [fx](https://github.com/uber-go/fx) | 939 | 49 | 2016-10-27 | 1 day ago | A dependency injection based application framework for Go (built on top of dig). |
| [inject](https://github.com/defval/inject) | 40 | 2 | 2019-02-04 | 3 months ago | A reflection based dependency injection container with simple interface. |
| [container](https://github.com/golobby/container) | 36 | 1 | 2019-09-24 | 3 weeks ago | A powerful IoC Container with fluent and easy-to-use interface. |
| [alice](https://github.com/magic003/alice) | 35 | 1 | 2017-04-09 | 2 years ago | Additive dependency injection container for Golang. |
| [wire](https://github.com/Fs02/wire) | 20 | 1 | 2018-07-05 | 3 months ago | Strict Runtime Dependency Injection for Golang. |
| [linker](https://github.com/logrange/linker) | 11 | 1 | 2018-12-05 | 1 week ago | A reflection based dependency injection and inversion of control library with components lifecycle support. |
| [gocontainer](https://github.com/vardius/gocontainer) | 10 | 0 | 2019-06-06 | 2 months ago | Simple Dependency Injection Container. |

### Strings
        
*These libraries were placed here because none of the other categories seemed to fit.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gopsutil](https://github.com/shirou/gopsutil) | 4190 | 183 | 2014-04-18 | 4 days ago | Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc). |
| [archiver](https://github.com/mholt/archiver) | 2571 | 45 | 2016-04-09 | 23 hours ago | Library and command for making and extracting .zip and .tar.gz archives. |
| [gosms](https://github.com/haxpax/gosms) | 1236 | 56 | 2015-01-24 | 2 years ago | Your own local SMS gateway in Go that can be used to send SMS. |
| [go-resiliency](https://github.com/eapache/go-resiliency) | 905 | 28 | 2014-11-29 | 3 months ago | Resiliency patterns for golang. |
| [go-commons-pool](https://github.com/jolestar/go-commons-pool) | 717 | 44 | 2015-12-28 | 5 months ago | Generic object pool for Golang. |
| [base64Captcha](https://github.com/mojocn/base64Captcha) | 669 | 44 | 2017-12-12 | 2 months ago | Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha. |
| [xstrings](https://github.com/huandu/xstrings) | 638 | 26 | 2015-01-06 | 1 year ago | Collection of useful string functions ported from other languages. |
| [shortid](https://github.com/teris-io/shortid) | 480 | 6 | 2016-01-04 | 2 years ago | Distributed generation of super short, unique, non-sequential, URL friendly IDs. |
| [gofakeit](https://github.com/brianvoe/gofakeit) | 472 | 5 | 2015-04-24 | 2 weeks ago | Random data generator written in go. |
| [llvm](https://github.com/llir/llvm) | 445 | 30 | 2014-09-19 | 1 hour ago | Library for interacting with LLVM IR in pure Go. |
| [health](https://github.com/dimiro1/health) | 369 | 6 | 2016-03-09 | 1 week ago | Easy to use, extensible health check library. |
| [go-conv](https://github.com/cstockton/go-conv) | 342 | 8 | 2016-10-11 | 2 years ago | Package conv provides fast and intuitive conversions across Go types. |
| [banner](https://github.com/dimiro1/banner) | 239 | 4 | 2016-03-26 | 6 days ago | Add beautiful banners into your Go applications. |
| [gountries](https://github.com/pariz/gountries) | 219 | 7 | 2016-01-13 | 1 month ago | Package that exposes country and subdivision data. |
| [antch](https://github.com/antchfx/antch) | 157 | 12 | 2017-09-28 | 1 year ago | A fast, powerful and extensible web crawling & scraping framework. |
| [battery](https://github.com/distatus/battery) | 138 | 2 | 2016-03-13 | 3 weeks ago | Cross-platform, normalized battery information library. |
| [ffmt](https://github.com/go-ffmt/ffmt) | 130 | 4 | 2015-02-14 | 4 months ago | Beautify data display for Humans. |
| [lk](https://github.com/hyperboloide/lk) | 129 | 5 | 2016-07-15 | 5 months ago | A simple licensing library for golang. |
| [stats](https://github.com/go-playground/stats) | 127 | 3 | 2015-09-15 | 3 years ago | Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... |
| [stateless](https://github.com/qmuntal/stateless) | 126 | 3 | 2019-09-11 | 1 week ago | A fluent library for creating state machines. |
| [bitio](https://github.com/icza/bitio) | 103 | 6 | 2016-05-31 | 1 year ago | Highly optimized bit-level Reader and Writer for Go. |
| [healthcheck](https://github.com/etherlabsio/healthcheck) | 91 | 4 | 2017-08-18 | 5 months ago | An opinionated and concurrent health-check HTTP handler for RESTful services. |
| [turtle](https://github.com/hackebrot/turtle) | 78 | 1 | 2017-09-09 | 2 weeks ago | Emojis for Go. |
| [gommit](https://github.com/antham/gommit) | 78 | 3 | 2016-08-30 | 2 months ago | Analyze git commit messages to ensure they follow defined patterns. |
| [strutil](https://github.com/ozgio/strutil) | 74 | 1 | 2018-08-16 | 2 weeks ago | String utilities. |
| [go-unarr](https://github.com/gen2brain/go-unarr) | 74 | 5 | 2015-11-01 | 8 months ago | Decompression library for RAR, TAR, ZIP and 7z archives. |
| [ghorg](https://github.com/gabrie30/ghorg) | 60 | 2 | 2018-03-29 | 1 month ago | Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, and Bitbucket. |
| [indigo](https://github.com/osamingo/indigo) | 56 | 1 | 2016-08-31 | 1 month ago | Distributed unique ID generator of using Sonyflake and encoded by Base58. |
| [morse](https://github.com/alwindoss/morse) | 51 | 2 | 2018-08-15 | 8 months ago | Library to convert to and from morse code. |
| [captcha](https://github.com/steambap/captcha) | 44 | 4 | 2017-09-12 | 4 days ago | Package captcha provides an easy to use, unopinionated API for captcha generation. |
| [xkg](https://github.com/go-xkg/xkg) | 43 | 1 | 2015-01-05 | 4 years ago | X Keyboard Grabber. |
| [pdfgen](https://github.com/hyperboloide/pdfgen) | 34 | 3 | 2015-12-01 | 1 year ago | HTTP service to generate PDF from Json requests. |
| [persian](https://github.com/mavihq/persian) | 33 | 2 | 2017-10-17 | 1 year ago | Some utilities for Persian language in go. |
| [browscap_go](https://github.com/digitalcrab/browscap_go) | 31 | 6 | 2014-09-18 | 2 weeks ago | GoLang Library for [Browser Capabilities Project](http://browscap.org/). |
| [datacounter](https://github.com/miolini/datacounter) | 30 | 1 | 2015-10-15 | 3 months ago | Go counters for readers/writer/http.ResponseWriter. |
| [autoflags](https://github.com/artyom/autoflags) | 25 | 3 | 2014-05-16 | 9 months ago | Go package to automatically define command line flags from struct fields. |
| [gotoprom](https://github.com/cabify/gotoprom) | 23 | 80 | 2018-10-11 | 2 weeks ago | Type-safe metrics builder wrapper library for the official Prometheus client. |
| [xdg](https://github.com/rkoesters/xdg) | 20 | 1 | 2013-12-15 | 11 months ago | FreeDesktop.org (xdg) Specs implemented in Go. |
| [url-shortener](https://github.com/pantrif/url-shortener) | 18 | 1 | 2018-06-04 | 1 year ago | A modern, powerful, and robust URL shortener microservice with mysql support. |
| [gosh](https://github.com/osamingo/gosh) | 17 | 0 | 2018-05-25 | 1 month ago | Provide Go Statistics Handler, Struct, Measure Method. |
| [sandid](https://github.com/aofei/sandid) | 15 | 1 | 2018-06-12 | 1 month ago | Every grain of sand on earth has its own ID. |
| [anagent](https://github.com/mudler/anagent) | 10 | 2 | 2017-12-30 | 1 year ago | Minimalistic, pluggable Golang evloop/timer handler with dependency-injection. |
| [avgRating](https://github.com/kirillDanshin/avgRating) | 9 | 1 | 2017-08-06 | 2 years ago | Calculate average score and rating based on Wilson Score Equation. |
| [hostutils](https://github.com/Wing924/hostutils) | 8 | 1 | 2017-09-26 | 9 months ago | A golang library for packing and unpacking FQDNs list. |
| [shellwords](https://github.com/Wing924/shellwords) | 7 | 1 | 2017-09-28 | 2 years ago | A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell. |
| [metrics](https://github.com/pascaldekloe/metrics) | 4 | 1 | 2019-01-29 | 5 months ago | Library for metrics instrumentation and Prometheus exposition. |
| [numa](https://github.com/lrita/numa) | 2 | 1 | 2018-12-10 | 5 months ago | NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code. |

### Project Layout
        
*Unofficial set of patterns for structuring projects.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [project-layout](https://github.com/golang-standards/project-layout) | 10903 | 303 | 2017-09-10 | 1 week ago | Set of common historical and emerging project layout patterns in the Go ecosystem. |
| [modern-go-application](https://github.com/sagikazarmark/modern-go-application) | 388 | 11 | 2018-09-14 | 4 hours ago | Go application boilerplate and example applying modern practices. |
| [scaffold](https://github.com/catchplay/scaffold) | 35 | 2 | 2018-12-11 | 9 months ago | Scaffold generates starter Go project layout. Lets you focus on business logic implemeted. |
| [go-sample](https://github.com/zitryss/go-sample) | 33 | 1 | 2019-01-25 | 9 months ago | A sample layout for Go application projects with the real code. |

## Natural Language Processing
        
*Libraries for working with human languages.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [prose](https://github.com/jdkato/prose) | 2348 | 55 | 2017-02-18 | 3 weeks ago | Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. English only. |
| [gse](https://github.com/go-ego/gse) | 1118 | 43 | 2017-06-23 | 5 days ago | Go efficient text segmentation; support english, chinese, japanese and other. |
| [when](https://github.com/olebedev/when) | 935 | 21 | 2016-12-27 | 20 hours ago | Natural EN and RU language date/time parser with pluggable rules. |
| [gojieba](https://github.com/yanyiwu/gojieba) | 873 | 59 | 2015-09-12 | 4 days ago | This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm. |
| [go-pinyin](https://github.com/mozillazg/go-pinyin) | 566 | 30 | 2014-11-09 | 6 months ago | CN Hanzi to Hanyu Pinyin converter. |
| [kagome](https://github.com/ikawaha/kagome) | 450 | 22 | 2014-06-26 | 2 weeks ago | JP morphological analyzer written in pure Go. |
| [whatlanggo](https://github.com/abadojack/whatlanggo) | 365 | 13 | 2017-02-21 | 7 months ago | Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc). |
| [nlp](https://github.com/shixzie/nlp) | 357 | 24 | 2017-01-25 | 2 years ago | Extract values from strings and fill your structs with nlp. |
| [sentences](https://github.com/neurosnap/sentences) | 263 | 12 | 2015-08-07 | 6 months ago | Sentence tokenizer:  converts text into a list of sentences. |
| [nlp](https://github.com/james-bowman/nlp) | 224 | 22 | 2017-03-15 | 1 week ago | Go Natural Language Processing library supporting LSA (Latent Semantic Analysis). |
| [go-nlp](https://github.com/nuance/go-nlp) | 80 | 8 | 2011-05-02 | 8 years ago | Utilities for working with discrete probability distributions and other tools useful for doing NLP work. |
| [getlang](https://github.com/rylans/getlang) | 77 | 3 | 2018-03-02 | 4 months ago | Fast natural language detection package. |
| [gounidecode](https://github.com/fiam/gounidecode) | 68 | 3 | 2012-05-01 | 4 years ago | Unicode transliterator (also known as unidecode) for Go. |
| [go-unidecode](https://github.com/mozillazg/go-unidecode) | 64 | 2 | 2016-07-08 | 6 months ago | ASCII transliterations of Unicode text. |
| [textcat](https://github.com/pebbe/textcat) | 61 | 6 | 2012-09-21 | 1 year ago | Go package for n-gram based text categorization, with support for utf-8 and raw text. |
| [MMSEGO](https://github.com/awsong/MMSEGO) | 59 | 5 | 2012-04-18 | 7 years ago | This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm. |
| [go-stem](https://github.com/agonopol/go-stem) | 52 | 2 | 2011-09-24 | 1 year ago | Implementation of the porter stemming algorithm. |
| [stemmer](https://github.com/dchest/stemmer) | 48 | 4 | 2011-03-21 | 2 years ago | Stemmer packages for Go programming language. Includes English and German stemmers. |
| [RAKE.Go](https://github.com/afjoseph/RAKE.Go) | 46 | 5 | 2016-12-17 | 10 months ago | Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE). |
| [porter2](https://github.com/zentures/porter2) | 34 | 2 | 2015-01-21 | 4 years ago | Really fast Porter 2 stemmer. |
| [go2vec](https://github.com/danieldk/go2vec) | 32 | 6 | 2015-01-27 | 1 year ago | Reader and utility functions for word2vec embeddings. |
| [paicehusk](https://github.com/rookii/paicehusk) | 25 | 3 | 2012-09-30 | 5 years ago | Golang implementation of the Paice/Husk Stemming Algorithm. |
| [snowball](https://github.com/goodsign/snowball) | 24 | 0 | 2012-12-11 | 2 years ago | Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/). |
| [petrovich](https://github.com/striker2000/petrovich) | 24 | 1 | 2016-12-27 | 10 months ago | Petrovich is the library which inflects Russian names to given grammatical case. |
| [mystem](https://github.com/dveselov/mystem) | 23 | 3 | 2016-08-30 | 3 years ago | CGo bindings to Yandex.Mystem - russian morphology analyzer. |
| [icu](https://github.com/goodsign/icu) | 19 | 0 | 2012-12-11 | 2 years ago | Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1. |
| [golibstemmer](https://github.com/rjohnsondev/golibstemmer) | 15 | 1 | 2012-08-07 | 5 years ago | Go bindings for the snowball libstemmer library including porter 2. |
| [libtextcat](https://github.com/goodsign/libtextcat) | 10 | 1 | 2012-12-11 | 6 years ago | Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2. |
| [shamoji](https://github.com/osamingo/shamoji) | 10 | 2 | 2017-07-23 | 1 month ago | The shamoji is word filtering package written in Go. |
| [porter](https://github.com/a2800276/porter) | 8 | 1 | 2013-09-17 | 6 years ago | This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm. |
| [gotokenizer](https://github.com/xujiajun/gotokenizer) | 6 | 1 | 2018-10-11 | 6 months ago | A tokenizer based on the dictionary and Bigram language models for Go. (Now only support chinese segmentation) |
| [go-eco](https://github.com/ThePaw/go-eco) | 4 | 0 | 2015-09-14 | 4 years ago | Similarity, dissimilarity and distance matrices; diversity, equitability and inequality measures; species richness estimators; coenocline models. |

## Networking
        
*Libraries for working with various layers of the network.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [kcptun](https://github.com/xtaci/kcptun) | 11043 | 589 | 2016-02-26 | 2 days ago | Extremely simple & fast udp tunnel based on KCP protocol. |
| [fasthttp](https://github.com/valyala/fasthttp) | 10416 | 352 | 2015-10-19 | 1 hour ago | Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http. |
| [dns](https://github.com/miekg/dns) | 4006 | 164 | 2010-08-04 | 5 days ago | Go library for working with DNS. |
| [httplab](https://github.com/gchaincl/httplab) | 3457 | 67 | 2017-02-09 | 4 months ago | HTTPLabs let you inspect HTTP requests and forge responses. |
| [quic-go](https://github.com/lucas-clemente/quic-go) | 3301 | 158 | 2016-04-07 | 15 hours ago | An implementation of the QUIC protocol in pure Go. |
| [gopacket](https://github.com/google/gopacket) | 3064 | 135 | 2015-03-17 | 6 days ago | Go library for packet processing with libpcap bindings. |
| [webrtc](https://github.com/pion/webrtc) | 2593 | 123 | 2018-05-19 | 4 days ago | A pure Go implementation of the WebRTC API. |
| [kcp-go](https://github.com/xtaci/kcp-go) | 2354 | 139 | 2015-06-16 | 1 day ago | KCP - Fast and Reliable ARQ Protocol. |
| [gobgp](https://github.com/osrg/gobgp) | 1742 | 122 | 2014-09-14 | 1 day ago | BGP implemented in the Go Programming Language. |
| [webrtc](https://github.com/pions/webrtc) | 1423 | 77 | 2018-05-19 | 6 months ago | A pure Go implementation of the WebRTC API. |
| [gnet](https://github.com/panjf2000/gnet) | 1200 | 37 | 2019-02-24 | 3 hours ago | `gnet` is a high-performance, lightweight, nonblocking, event-loop networking library written in pure Go. |
| [ssh](https://github.com/gliderlabs/ssh) | 1184 | 39 | 2016-10-04 | 1 day ago | Higher-level API for building SSH servers (wraps crypto/ssh). |
| [fortio](https://github.com/fortio/fortio) | 977 | 32 | 2017-10-10 | 1 month ago | Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC. |
| [water](https://github.com/songgao/water) | 893 | 37 | 2013-03-26 | 2 months ago | Simple TUN/TAP library. |
| [sftp](https://github.com/pkg/sftp) | 782 | 51 | 2013-11-05 | 7 hours ago | Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt. |
| [go-getter](https://github.com/hashicorp/go-getter) | 779 | 143 | 2015-10-13 | 1 week ago | Go library for downloading files or directories from various sources using a URL. |
| [nff-go](https://github.com/intel-go/nff-go) | 717 | 69 | 2017-03-30 | 1 week ago | Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF). |
| [grab](https://github.com/cavaliercoder/grab) | 581 | 17 | 2016-01-05 | 1 day ago | Go package for managing file downloads. |
| [gev](https://github.com/Allenxuxu/gev) | 577 | 12 | 2019-09-01 | 5 hours ago | gev is a lightweight, fast non-blocking TCP network library based on Reactor mode. |
| [mdns](https://github.com/hashicorp/mdns) | 576 | 151 | 2014-01-30 | 1 month ago | Simple mDNS (Multicast DNS) client/server library in Golang. |
| [ftp](https://github.com/jlaffaye/ftp) | 567 | 25 | 2011-05-07 | 2 days ago | Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959). |
| [lhttp](https://github.com/fanux/lhttp) | 527 | 56 | 2015-12-29 | 1 year ago | Powerful websocket framework, build your IM server more easily. |
| [gosnmp](https://github.com/soniah/gosnmp) | 459 | 42 | 2012-08-27 | 1 week ago | Native Go library for performing SNMP actions. |
| [gotcp](https://github.com/gansidui/gotcp) | 432 | 42 | 2014-04-13 | 2 years ago | Go package for quickly writing tcp applications. |
| [cidranger](https://github.com/yl2chen/cidranger) | 407 | 11 | 2017-08-21 | 1 month ago | Fast IP to CIDR lookup for Go. |
| [peerdiscovery](https://github.com/schollz/peerdiscovery) | 386 | 15 | 2018-04-23 | 2 months ago | Pure Go library for cross-platform local peer discovery using UDP multicast. |
| [gopcap](https://github.com/akrennmair/gopcap) | 362 | 24 | 2009-11-19 | 2 months ago | Go wrapper for libpcap. |
| [go-stun](https://github.com/ccding/go-stun) | 339 | 11 | 2013-08-18 | 1 year ago | Go implementation of the STUN client (RFC 3489 and RFC 5389). |
| [raw](https://github.com/mdlayher/raw) | 325 | 11 | 2015-07-07 | 2 weeks ago | Package raw enables reading and writing data at the device driver level for a network interface. |
| [stun](https://github.com/gortc/stun) | 312 | 14 | 2016-04-25 | 3 weeks ago | Go implementation of RFC 5389 STUN protocol. |
| [tcp_server](https://github.com/firstrow/tcp_server) | 300 | 17 | 2014-10-14 | 6 months ago | Go library for building tcp servers faster. |
| [buffstreams](https://github.com/StabbyCutyou/buffstreams) | 234 | 14 | 2015-06-30 | 2 years ago | Streaming protocolbuffer data over TCP made easy. |
| [winrm](https://github.com/masterzen/winrm) | 228 | 15 | 2013-12-31 | 2 days ago | Go WinRM client to remotely execute commands on Windows machines. |
| [arp](https://github.com/mdlayher/arp) | 204 | 10 | 2015-07-07 | 1 week ago | Package arp implements the ARP protocol, as described in RFC 826. |
| [ethernet](https://github.com/mdlayher/ethernet) | 188 | 12 | 2015-07-03 | 4 months ago | Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. |
| [utp](https://github.com/anacrolix/utp) | 150 | 16 | 2015-03-20 | 1 year ago | Go uTP micro transport protocol implementation. |
| [canopus](https://github.com/zubairhamed/canopus) | 136 | 15 | 2015-02-24 | 1 year ago | CoAP Client/Server implementation (RFC 7252). |
| [jazigo](https://github.com/udhos/jazigo) | 132 | 11 | 2016-06-08 | 1 month ago | Jazigo is a tool written in Go for retrieving configuration for multiple network devices. |
| [sslb](https://github.com/eduardonunesp/sslb) | 117 | 7 | 2015-10-19 | 1 month ago | It's a Super Simples Load Balancer, just a little project to achieve some kind of performance. |
| [gmqtt](https://github.com/DrmagicE/gmqtt) | 112 | 9 | 2018-09-16 | 1 month ago | Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1. |
| [gnxi](https://github.com/google/gnxi) | 109 | 21 | 2017-09-26 | 1 week ago | A collection of tools for Network Management that use the gNMI and gNOI protocols. |
| [xtcp](https://github.com/xfxdev/xtcp) | 89 | 12 | 2016-04-01 | 9 months ago | TCP Server Framework with simultaneous full duplex communication,graceful shutdown,custom protocol. |
| [dhcp6](https://github.com/mdlayher/dhcp6) | 63 | 4 | 2015-05-22 | 7 months ago | Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. |
| [ether](https://github.com/songgao/ether) | 62 | 2 | 2014-05-21 | 3 years ago | Cross-platform Go package for sending and receiving ethernet frames. |
| [linkio](https://github.com/ian-kent/linkio) | 44 | 2 | 2014-12-24 | 2 years ago | Network link speed simulation for Reader/Writer interfaces. |
| [portproxy](https://github.com/aybabtme/portproxy) | 43 | 0 | 2014-12-13 | 4 years ago | Simple TCP proxy which adds CORS support to API's which don't support it. |
| [packet](https://github.com/aerogo/packet) | 36 | 3 | 2017-10-29 | 2 days ago | Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed. |
| [graval](https://github.com/koofr/graval) | 24 | 8 | 2014-04-23 | 1 year ago | Experimental FTP server framework. |
| [publicip](https://github.com/polera/publicip) | 18 | 1 | 2016-12-29 | 2 years ago | Package publicip returns your public facing IPv4 address (internet egress). |
| [golibwireshark](https://github.com/sunwxg/golibwireshark) | 16 | 2 | 2015-11-16 | 2 years ago | Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data. |
| [go-powerdns](https://github.com/joeig/go-powerdns) | 9 | 3 | 2018-06-22 | 3 weeks ago | PowerDNS API bindings for Golang. |
| [goshark](https://github.com/sunwxg/goshark) | 9 | 1 | 2015-11-01 | 2 years ago | Package goshark use tshark to decode IP packet and create data struct to analyse packet. |
| [llb](https://github.com/kirillDanshin/llb) | 9 | 1 | 2016-02-21 | 3 years ago | It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response. |
| [tspool](https://github.com/two/tspool) | 5 | 0 | 2018-10-27 | 1 year ago | A TCP Library use worker pool to improve performance and protect your server. |

### HTTP Clients
        
*Libraries for making HTTP requests.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [resty](https://github.com/go-resty/resty) | 2216 | 62 | 2015-08-29 | 3 days ago | Simple HTTP and REST client for Go inspired by Ruby rest-client. |
| [grequests](https://github.com/levigross/grequests) | 1469 | 31 | 2015-06-12 | 3 weeks ago | A Go "clone" of the great and famous Requests library. |
| [heimdall](https://github.com/gojek/heimdall) | 1150 | 46 | 2018-01-19 | 6 months ago | An enchanced http client with retry and hystrix capabilities. |
| [sling](https://github.com/dghubble/sling) | 1045 | 33 | 2015-04-02 | 1 month ago | Sling is a Go HTTP client library for creating and sending API requests. |
| [heimdall](https://github.com/gojektech/heimdall) | 1017 | 45 | 2018-01-19 | 6 months ago | An enchanced http client with retry and hystrix capabilities. |
| [gentleman](https://github.com/h2non/gentleman) | 693 | 17 | 2016-02-22 | 9 months ago | Full-featured plugin-driven HTTP client library. |
| [pester](https://github.com/sethgrid/pester) | 337 | 7 | 2015-05-20 | 9 months ago | Go HTTP client calls with retries, backoff, and concurrency. |
| [goreq](https://github.com/smallnest/goreq) | 101 | 7 | 2015-09-17 | 1 year ago | Enhanced simplified HTTP client based on gorequest. |
| [rq](https://github.com/ddo/rq) | 32 | 2 | 2017-12-26 | 2 months ago | A nicer interface for golang stdlib HTTP client. |
| [sreq](https://github.com/winterssy/sreq) | 21 | 0 | 2019-09-29 | 3 days ago | A simple, user-friendly and concurrent safe HTTP request library for Go. |

## OpenGL
        
*Libraries for using OpenGL in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [glfw](https://github.com/go-gl/glfw) | 779 | 39 | 2013-05-19 | 6 months ago | Go bindings for GLFW 3. |
| [gl](https://github.com/go-gl/gl) | 661 | 40 | 2015-02-22 | 7 months ago | Go bindings for OpenGL (generated via glow). |
| [mathgl](https://github.com/go-gl/mathgl) | 303 | 24 | 2013-02-13 | 1 week ago | Pure Go math package specialized for 3D math, with inspiration from GLM. |
| [gl](https://github.com/goxjs/gl) | 130 | 13 | 2015-05-18 | 11 months ago | Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android). |
| [glfw](https://github.com/goxjs/glfw) | 58 | 6 | 2014-12-28 | 11 months ago | Go cross-platform glfw library for creating an OpenGL context and receiving events. |

## ORM
        
*Libraries that implement Object-Relational Mapping or datamapping techniques.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gorm](https://github.com/jinzhu/gorm) | 15645 | 426 | 2013-10-25 | 7 hours ago | The fantastic ORM library for Golang, aims to be developer friendly. |
| [xorm](https://github.com/go-xorm/xorm) | 5490 | 259 | 2013-05-09 | 5 days ago | Simple and powerful ORM for Go. |
| [pg](https://github.com/go-pg/pg) | 3230 | 81 | 2013-04-24 | 1 day ago | PostgreSQL ORM with focus on PostgreSQL specific features and performance. |
| [gorp](https://github.com/go-gorp/gorp) | 3152 | 116 | 2012-01-05 | 1 day ago | Go Relational Persistence, ORM-ish library for Go. |
| [sqlboiler](https://github.com/volatiletech/sqlboiler) | 2436 | 72 | 2016-02-21 | 5 hours ago | ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema. |
| [db](https://github.com/upper/db) | 1964 | 60 | 2013-10-23 | 2 months ago | Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers. |
| [reform](https://github.com/go-reform/reform) | 822 | 28 | 2016-02-25 | 1 week ago | Better ORM for Go, based on non-empty interfaces and code generation. |
| [pop](https://github.com/gobuffalo/pop) | 725 | 22 | 2018-02-08 | 6 days ago | Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite. |
| [qbs](https://github.com/coocood/qbs) | 538 | 45 | 2013-02-02 | 2 years ago | Stands for Query By Struct. A Go ORM. |
| [go-queryset](https://github.com/jirfag/go-queryset) | 461 | 19 | 2017-09-04 | 1 week ago | 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM. |
| [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) | 265 | 14 | 2017-12-28 | 3 days ago | A flexible and powerful SQL string builder library plus a zero-config ORM. |
| [zoom](https://github.com/albrow/zoom) | 245 | 16 | 2013-07-17 | 1 year ago | Blazing-fast datastore and querying engine built on Redis. |
| [grimoire](https://github.com/Fs02/grimoire) | 120 | 5 | 2018-03-06 | 1 month ago | Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3). |
| [go-store](https://github.com/gosuri/go-store) | 96 | 9 | 2015-03-22 | 2 years ago | Simple and fast Redis backed key-value store library for Go. |
| [marlow](https://github.com/dadleyy/marlow) | 68 | 5 | 2017-09-15 | 1 month ago | Generated ORM from project structs for compile time safety assurances. |
| [gomodel](https://github.com/cosiner/gomodel) | 62 | 5 | 2015-04-16 | 2 years ago | Lightweight, fast, orm-like library helps interactive with database. |
| [go-firestorm](https://github.com/jschoedt/go-firestorm) | 5 | 1 | 2018-12-04 | 6 hours ago | A simple ORM for Google/Firebase Cloud Firestore. |
| [lore](https://github.com/abrahambotros/lore) | 5 | 1 | 2017-04-29 | 2 years ago | Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go. |

## Package Management
        
*Unofficial libraries for package and dependency management.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [dep](https://github.com/golang/dep) | 12813 | 293 | 2016-10-07 | 3 weeks ago | Go dependency tool. |
| [glide](https://github.com/Masterminds/glide) | 7856 | 203 | 2014-07-09 | 3 months ago | Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip. |
| [godep](https://github.com/tools/godep) | 5641 | 157 | 2013-05-01 | 1 year ago | dependency tool for go, godep helps build packages reproducibly by fixing their dependencies. |
| [govendor](https://github.com/kardianos/govendor) | 4864 | 104 | 2015-04-12 | 2 months ago | Go Package Manager. Go vendor tool that works with the standard vendor file. |
| [gopm](https://github.com/gpmgo/gopm) | 2404 | 83 | 2013-05-15 | 3 months ago | Go Package Manager. |
| [gom](https://github.com/mattn/gom) | 1357 | 37 | 2013-09-11 | 3 months ago | Go Manager - bundle for go. |
| [gpm](https://github.com/pote/gpm) | 1205 | 31 | 2013-09-05 | 2 years ago | Barebones dependency manager for Go. |
| [goop](https://github.com/petejkim/goop) | 778 | 37 | 2014-06-18 | 3 years ago | Simple dependency manager for Go (golang), inspired by Bundler. |
| [nut](https://github.com/jingweno/nut) | 245 | 8 | 2015-01-23 | 4 years ago | Vendor Go dependencies. |
| [johnny-deps](https://github.com/VividCortex/johnny-deps) | 214 | 22 | 2013-07-19 | 5 years ago | Minimal dependency version using Git. |
| [gigo](https://github.com/LyricalSecurity/gigo) | 197 | 5 | 2015-05-01 | 11 months ago | PIP-like dependency tool for golang, with support for private repositories and hashes. |
| [VenGO](https://github.com/DamnWidget/VenGO) | 116 | 8 | 2014-10-17 | 3 years ago | create and manage exportable isolated go virtual environments. |
| [mvn-golang](https://github.com/raydac/mvn-golang) | 90 | 8 | 2016-03-24 | 1 week ago | plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure. |
| [gop](https://github.com/lunny/gop) | 50 | 7 | 2017-02-18 | 7 months ago | Build and manage your Go applications out of GOPATH. |

## Query Language
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [graphql](https://github.com/graphql-go/graphql) | 5576 | 142 | 2015-07-19 | 5 days ago | Implementation of GraphQL for Go. |
| [graphql-go](https://github.com/graph-gophers/graphql-go) | 2951 | 85 | 2016-10-18 | 8 hours ago | GraphQL server with a focus on ease of use. |
| [gojsonq](https://github.com/thedevsaddam/gojsonq) | 1052 | 22 | 2018-05-20 | 1 week ago | A simple Go package to Query over JSON Data. |
| [jsonql](https://github.com/elgs/jsonql) | 207 | 8 | 2015-12-29 | 1 week ago | JSON query expression library in Golang. |
| [rql](https://github.com/a8m/rql) | 118 | 6 | 2018-06-06 | 1 month ago | Resource Query Language for REST API. |
| [graphql](https://github.com/tmc/graphql) | 51 | 10 | 2015-04-19 | 2 years ago | graphql parser + utilities. |
| [jsonslice](https://github.com/bhmj/jsonslice) | 26 | 0 | 2018-05-02 | 1 week ago | Jsonpath queries with advanced filters. |
| [straf](https://github.com/SonicRoshan/straf) | 3 | 1 | 2019-08-16 | 1 month ago | Easily Convert Golang structs to GraphQL objects. |

## Resource Embedding
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [packr](https://github.com/gobuffalo/packr) | 2391 | 44 | 2017-03-16 | 1 week ago | The simple and easy way to embed static files into Go binaries. |
| [statik](https://github.com/rakyll/statik) | 2288 | 38 | 2014-02-04 | 3 days ago | Embeds static files into a Go executable. |
| [go.rice](https://github.com/GeertJohan/go.rice) | 1745 | 49 | 2013-10-24 | 1 week ago | go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy. |
| [vfsgen](https://github.com/shurcooL/vfsgen) | 721 | 20 | 2015-05-18 | 8 months ago | Generates a vfsdata.go file that statically implements the given virtual filesystem. |
| [esc](https://github.com/mjibson/esc) | 494 | 13 | 2014-01-26 | 6 months ago | Embeds files into Go programs and provides http.FileSystem interfaces to them. |
| [fileb0x](https://github.com/UnnoTed/fileb0x) | 467 | 14 | 2016-01-24 | 1 week ago | Simple tool to embed files in go with focus on "customization" and ease to use. |
| [go-resources](https://github.com/omeid/go-resources) | 158 | 6 | 2015-02-21 | 7 months ago | Unfancy resources embedding with Go. |
| [statics](https://github.com/go-playground/statics) | 54 | 3 | 2015-10-07 | 3 years ago | Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks. |
| [templify](https://github.com/wlbr/templify) | 21 | 2 | 2016-05-23 | 2 months ago | Embed external template files into Go code to create single file binaries. |
| [go-embed](https://github.com/pyros2097/go-embed) | 18 | 1 | 2015-12-07 | 2 months ago | Generates go code to embed resource files into your library or executable. |

## Science and Data Analysis
        
*Libraries for scientific computing and data analyzing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gonum](https://github.com/gonum/gonum) | 3209 | 94 | 2017-03-25 | 5 hours ago | Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more. |
| [stats](https://github.com/montanaflynn/stats) | 1386 | 40 | 2014-12-16 | 1 month ago | Statistics package with common functions missing from the Golang standard library. |
| [gosl](https://github.com/cpmech/gosl) | 1346 | 65 | 2015-02-10 | 2 months ago | Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more. |
| [streamtools](https://github.com/nytlabs/streamtools) | 1316 | 75 | 2013-07-06 | 4 years ago | general purpose, graphical tool for dealing with streams of data. |
| [plot](https://github.com/gonum/plot) | 1278 | 53 | 2013-07-23 | 3 weeks ago | gonum/plot provides an API for building and drawing plots in Go. |
| [go-dsp](https://github.com/mjibson/go-dsp) | 640 | 29 | 2011-11-02 | 1 year ago | Digital Signal Processing for Go. |
| [goraph](https://github.com/gyuho/goraph) | 610 | 37 | 2014-02-27 | 2 years ago | Pure Go graph theory library(data structure, algorith visualization). |
| [chart](https://github.com/vdobler/chart) | 595 | 43 | 2011-06-27 | 3 months ago | Simple Chart Plotting library for Go. Supports many graphs types. |
| [ewma](https://github.com/VividCortex/ewma) | 273 | 26 | 2013-07-06 | 3 days ago | Exponentially-weighted moving averages. |
| [graph](https://github.com/yourbasic/graph) | 265 | 17 | 2017-04-28 | 1 year ago | Library of basic graph algorithms. |
| [orb](https://github.com/paulmach/orb) | 224 | 16 | 2016-03-28 | 1 month ago | 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support. |
| [gohistogram](https://github.com/VividCortex/gohistogram) | 131 | 16 | 2013-07-02 | 1 year ago | Approximate histograms for data streams. |
| [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) | 98 | 11 | 2018-10-01 | 1 week ago | Dataframes for Go for machine-learning and statistics (similar to pandas). |
| [TextRank](https://github.com/DavidBelicza/TextRank) | 80 | 5 | 2018-01-10 | 1 year ago | TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support. |
| [sparse](https://github.com/james-bowman/sparse) | 76 | 5 | 2017-05-16 | 6 months ago | Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries. |
| [pagerank](https://github.com/alixaxel/pagerank) | 52 | 6 | 2015-08-06 | 4 months ago | Weighted PageRank algorithm implemented in Go. |
| [geom](https://github.com/skelterjohn/geom) | 42 | 4 | 2011-06-08 | 1 year ago | 2D geometry for golang. |
| [evaler](https://github.com/soniah/evaler) | 40 | 4 | 2012-09-05 | 1 year ago | Simple floating point arithmetic expression evaluator. |
| [goent](https://github.com/kzahedi/goent) | 13 | 1 | 2017-08-08 | 6 months ago | GO Implementation of Entropy Measures. |
| [triangolatte](https://github.com/tchayen/triangolatte) | 12 | 1 | 2018-07-19 | 3 months ago | 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs. |
| [ode](https://github.com/ChristopherRabotin/ode) | 11 | 3 | 2016-11-12 | 2 years ago | Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions. |
| [go-fn](https://github.com/ematvey/go-fn) | 11 | 0 | 2015-04-28 | 4 years ago | Mathematical functions written in Go language, that are not covered by math pkg. |
| [GoStats](https://github.com/OGFris/GoStats) | 10 | 1 | 2018-07-23 | 9 months ago | GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions. |
| [piecewiselinear](https://github.com/sgreben/piecewiselinear) | 8 | 1 | 2018-10-21 | 1 week ago | Tiny linear interpolation library. |
| [PiHex](https://github.com/claygod/PiHex) | 8 | 2 | 2016-07-22 | 1 week ago | Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi. |
| [go-gt](https://github.com/ThePaw/go-gt) | 5 | 0 | 2015-09-14 | 4 years ago | Graph theory algorithms written in "Go" language. |
| [gocomplex](https://github.com/varver/gocomplex) | 4 | 0 | 2015-06-25 | 4 years ago | Complex number library for the Go programming language. |
| [assocentity](https://github.com/ndabAP/assocentity) | 4 | 1 | 2018-12-21 | 1 month ago | Package assocentity returns the average distance from words to a given entity. |
| [rootfinding](https://github.com/khezen/rootfinding) | 3 | 2 | 2018-10-31 | 9 months ago | root-finding algorithms library for finding roots of quadratic functions. |
| [bradleyterry](https://github.com/seanhagen/bradleyterry) | 1 | 1 | 2019-04-30 | 5 months ago | Provides a Bradley-Terry Model for pairwise comparisons. |

## Security
        
*Libraries that are used to help make your application more secure.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [lego](https://github.com/go-acme/lego) | 3669 | 101 | 2015-06-08 | 1 day ago | Pure Go ACME client library and CLI tool (for use with Let's Encrypt). |
| [lego](https://github.com/xenolf/lego) | 3181 | 89 | 2015-06-08 | 7 months ago | Pure Go ACME client library and CLI tool (for use with Let's Encrypt). |
| [cameradar](https://github.com/Ullaakut/cameradar) | 1919 | 94 | 2016-05-20 | 1 week ago | Tool and library to remotely hack RTSP streams from surveillance cameras. |
| [acmetool](https://github.com/hlandau/acmetool) | 1713 | 66 | 2015-11-15 | 1 week ago | ACME (Let's Encrypt) client tool with automatic renewal. |
| [acme](https://github.com/hlandau/acme) | 1706 | 67 | 2015-11-15 | 2 months ago | ACME (Let's Encrypt) client tool with automatic renewal. |
| [memguard](https://github.com/awnumar/memguard) | 1591 | 44 | 2017-04-22 | 1 hour ago | A pure Go library for handling sensitive values in memory. |
| [secure](https://github.com/unrolled/secure) | 1258 | 36 | 2014-05-21 | 1 month ago | HTTP middleware for Go that facilitates some quick security wins. |
| [acra](https://github.com/cossacklabs/acra) | 477 | 32 | 2016-11-15 | 1 week ago | Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system. |
| [nacl](https://github.com/kevinburke/nacl) | 456 | 11 | 2017-07-21 | 2 months ago | Go implementation of the NaCL set of API's. |
| [badactor](https://github.com/jaredfolkins/badactor) | 252 | 8 | 2014-12-13 | 2 months ago | In-memory, application-driven jailer built in the spirit of fail2ban. |
| [passlib](https://github.com/hlandau/passlib) | 230 | 10 | 2014-12-22 | 7 months ago | Futureproof password hashing library. |
| [ssh-vault](https://github.com/ssh-vault/ssh-vault) | 199 | 10 | 2016-09-29 | 10 months ago | encrypt/decrypt using ssh keys. |
| [simple-scrypt](https://github.com/elithrar/simple-scrypt) | 160 | 7 | 2015-04-14 | 3 months ago | Scrypt package with a simple, obvious API and automatic cost calibration built-in. |
| [go-yara](https://github.com/hillu/go-yara) | 141 | 17 | 2015-01-25 | 2 weeks ago | Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)". |
| [argon2pw](https://github.com/raja/argon2pw) | 75 | 3 | 2018-03-13 | 1 year ago | Argon2 password hash generation with constant-time password comparison. |
| [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) | 32 | 1 | 2017-10-20 | 8 months ago | A probably paranoid package for securely hashing and encrypting passwords. |
| [goArgonPass](https://github.com/dwin/goArgonPass) | 11 | 2 | 2018-05-30 | 7 months ago | Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations. |
| [certificates](https://github.com/mvmaasakkers/certificates) | 10 | 1 | 2019-03-04 | 4 months ago | An opinionated tool for generating tls certificates. |
| [sslmgr](https://github.com/adrianosela/sslmgr) | 7 | 2 | 2019-04-03 | 3 months ago | SSL certificates made easy with a high level wrapper around acme/autocert. |
| [jwc](https://github.com/khezen/jwc) | 5 | 1 | 2018-06-28 | 2 months ago | JSON Web Cryptography library. |

## Serialization
        
*Libraries and tools for binary serialization.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go](https://github.com/json-iterator/go) | 6320 | 204 | 2016-11-30 | 1 week ago | High-performance 100% compatible drop-in replacement of "encoding/json". |
| [protobuf](https://github.com/golang/protobuf) | 5565 | 199 | 2014-11-24 | 5 hours ago | Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers. |
| [protobuf](https://github.com/gogo/protobuf) | 3214 | 93 | 2014-12-03 | 6 days ago | Protocol Buffers for Go with Gadgets. |
| [mapstructure](https://github.com/mitchellh/mapstructure) | 2706 | 51 | 2013-05-20 | 1 week ago | Go library for decoding generic map values into native Go structures. |
| [go](https://github.com/ugorji/go) | 1291 | 51 | 2013-05-30 | 1 month ago | High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support. |
| [go-codec](https://github.com/ugorji/go-codec) | 1149 | 49 | 2013-05-30 | 6 months ago | High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support. |
| [colfer](https://github.com/pascaldekloe/colfer) | 482 | 33 | 2015-09-06 | 4 months ago | Code generation for the Colfer binary format. |
| [csvutil](https://github.com/jszwec/csvutil) | 323 | 8 | 2017-10-30 | 3 months ago | High Performance, idiomatic CSV record encoding and decoding to native Go structures. |
| [go-capnproto](https://github.com/glycerine/go-capnproto) | 274 | 11 | 2013-11-08 | 9 months ago | Cap'n Proto library and parser for go. |
| [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) | 126 | 9 | 2012-12-23 | 1 year ago | GoLang library for working with PHP session format and PHP Serialize/Unserialize functions. |
| [structomap](https://github.com/danhper/structomap) | 102 | 7 | 2015-05-13 | 5 months ago | Library to easily and dynamically generate maps from static structures. |
| [bambam](https://github.com/glycerine/bambam) | 60 | 4 | 2014-09-17 | 3 years ago | generator for Cap'n Proto schemas from go. |
| [cbor](https://github.com/fxamacker/cbor) | 53 | 4 | 2019-05-16 | 17 hours ago | Small, safe, and easy CBOR encoding and decoding library. |
| [asn1](https://github.com/Logicalis/asn1) | 42 | 8 | 2016-02-29 | 7 months ago | Asn.1 BER and DER encoding library for golang. |
| [fwencoder](https://github.com/o1egl/fwencoder) | 10 | 1 | 2017-12-25 | 1 year ago | Fixed width file parser (encoding and decoding library) for Go. |
| [binstruct](https://github.com/ghostiam/binstruct) | 10 | 1 | 2018-10-23 | 1 month ago | Golang binary decoder for mapping data into the structure. |
| [bel](https://github.com/32leaves/bel) | 6 | 1 | 2019-02-21 | 6 months ago | Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC. |
| [pletter](https://github.com/vimeda/pletter) | 6 | 0 | 2019-07-09 | 3 months ago | A standard way to wrap a proto message for message brokers. |

## Template Engines
        
*Libraries and tools for templating and lexing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gofpdf](https://github.com/jung-kurt/gofpdf) | 3358 | 90 | 2015-03-13 | 3 weeks ago | PDF document generator with high level support for text, drawing and images. |
| [pongo2](https://github.com/flosch/pongo2) | 1534 | 54 | 2013-08-23 | 1 month ago | Django-like template-engine for Go. |
| [quicktemplate](https://github.com/valyala/quicktemplate) | 1482 | 55 | 2016-03-07 | 5 days ago | Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it. |
| [hero](https://github.com/shiyanhui/hero) | 1234 | 42 | 2017-01-15 | 2 months ago | Hero is a handy, fast and powerful go template engine. |
| [mustache](https://github.com/hoisie/mustache) | 968 | 36 | 2009-12-31 | 1 year ago | Go implementation of the Mustache template language. |
| [amber](https://github.com/eknkc/amber) | 827 | 21 | 2012-11-01 | 1 year ago | Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade. |
| [ace](https://github.com/yosssi/ace) | 768 | 24 | 2014-07-13 | 1 year ago | Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold. |
| [gorazor](https://github.com/sipin/gorazor) | 716 | 59 | 2014-05-01 | 3 months ago | Razor view engine for Golang. |
| [jet](https://github.com/CloudyKit/jet) | 592 | 21 | 2016-04-01 | 2 months ago | Jet template engine. |
| [ego](https://github.com/benbjohnson/ego) | 420 | 16 | 2014-02-24 | 9 months ago | Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled. |
| [raymond](https://github.com/aymerick/raymond) | 349 | 11 | 2015-04-22 | 8 months ago | Complete handlebars implementation in Go. |
| [fasttemplate](https://github.com/valyala/fasttemplate) | 318 | 18 | 2015-08-19 | 3 weeks ago | Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/). |
| [soy](https://github.com/robfig/soy) | 145 | 12 | 2013-12-15 | 7 months ago | Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/). |
| [liquid](https://github.com/osteele/liquid) | 87 | 3 | 2017-06-26 | 6 months ago | Go implementation of Shopify Liquid templates. |
| [goview](https://github.com/foolin/goview) | 76 | 3 | 2019-04-14 | 3 months ago | Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application. |
| [kasia.go](https://github.com/ziutek/kasia.go) | 70 | 2 | 2010-12-07 | 4 years ago | Templating system for HTML and other text documents - go implementation. |
| [maroto](https://github.com/johnfercher/maroto) | 67 | 4 | 2019-05-21 | 3 days ago | A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple. |
| [velvet](https://github.com/gobuffalo/velvet) | 65 | 5 | 2016-12-30 | 2 years ago | Complete handlebars implementation in Go. |
| [damsel](https://github.com/dskinner/damsel) | 20 | 4 | 2012-05-03 | 3 years ago | Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others. |
| [extemplate](https://github.com/dannyvankooten/extemplate) | 15 | 2 | 2018-08-11 | 1 year ago | Tiny wrapper around html/template to allow for easy file-based template inheritance. |

## Testing
        
*Libraries for testing codebases and generating test data.*

### Testing Frameworks
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [testify](https://github.com/stretchr/testify) | 8746 | 139 | 2012-10-17 | 2 weeks ago | Sacred extension to the standard go testing package. |
| [go-cmp](https://github.com/google/go-cmp) | 1283 | 24 | 2017-07-08 | 22 hours ago | Package for comparing Go values in tests. |
| [httpexpect](https://github.com/gavv/httpexpect) | 1189 | 31 | 2016-04-30 | 3 weeks ago | Concise, declarative, and easy to use end-to-end HTTP and REST API testing. |
| [godog](https://github.com/DATA-DOG/godog) | 827 | 28 | 2015-06-10 | 1 month ago | Cucumber or Behat like BDD framework for Go. |
| [baloo](https://github.com/h2non/baloo) | 658 | 10 | 2016-05-30 | 9 months ago | Expressive and versatile end-to-end HTTP API testing made easy. |
| [goblin](https://github.com/franela/goblin) | 638 | 16 | 2013-09-19 | 7 months ago | Mocha like testing framework fo Go. |
| [testfixtures](https://github.com/go-testfixtures/testfixtures) | 352 | 6 | 2016-04-05 | 3 days ago | A helper for Rails' like test fixtures to test database applications. |
| [go-vcr](https://github.com/dnaeon/go-vcr) | 348 | 7 | 2015-12-14 | 2 weeks ago | Record and replay your HTTP interactions for fast, deterministic and accurate tests. |
| [go-mutesting](https://github.com/zimmski/go-mutesting) | 308 | 6 | 2014-12-27 | 1 week ago | Mutation testing for Go source code. |
| [gofight](https://github.com/appleboy/gofight) | 275 | 11 | 2016-03-29 | 1 week ago | API Handler Testing for Golang Router framework. |
| [frisby](https://github.com/verdverm/frisby) | 252 | 8 | 2015-09-15 | 4 months ago | REST API testing framework. |
| [go-carpet](https://github.com/msoap/go-carpet) | 200 | 4 | 2016-02-28 | 1 day ago | Tool for viewing test coverage in terminal. |
| [charlatan](https://github.com/percolate/charlatan) | 192 | 39 | 2017-10-07 | 1 month ago | Tool to generate fake interface implementations for tests. |
| [gotest.tools](https://github.com/gotestyourself/gotest.tools) | 133 | 6 | 2017-08-09 | 22 hours ago | A collection of packages to augment the go testing package and support common patterns. |
| [gospec](https://github.com/luontola/gospec) | 112 | 4 | 2009-11-24 | 5 years ago | BDD-style testing framework for the Go programming language. |
| [endly](https://github.com/viant/endly) | 106 | 14 | 2017-08-29 | 3 days ago | Declarative end to end functional testing. |
| [dbcleaner](https://github.com/khaiql/dbcleaner) | 92 | 3 | 2017-01-18 | 6 months ago | Clean database for testing purpose, inspired by `database_cleaner` in Ruby. |
| [cupaloy](https://github.com/bradleyjkemp/cupaloy) | 90 | 3 | 2017-08-08 | 2 months ago | Simple snapshot testing addon for your test framework. |
| [wstest](https://github.com/posener/wstest) | 65 | 2 | 2017-04-01 | 1 year ago | Websocket client for unit-testing a websocket http.Handler. |
| [go-testdeep](https://github.com/maxatome/go-testdeep) | 63 | 4 | 2018-05-26 | 3 weeks ago | Extremely flexible golang deep comparison, extends the go testing package. |
| [gospecify](https://github.com/stesla/gospecify) | 53 | 5 | 2009-11-20 | 8 years ago | This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec. |
| [restit](https://github.com/go-restit/restit) | 49 | 6 | 2014-06-25 | 1 week ago | Go micro framework to help writing RESTful API integration test. |
| [commander](https://github.com/SimonBaeumer/commander) | 38 | 3 | 2019-02-23 | 4 days ago | Tool for testing cli applications on windows, linux and osx. |
| [gomatch](https://github.com/jfilipczyk/gomatch) | 32 | 2 | 2019-01-28 | 3 months ago | library created for testing JSON against patterns. |
| [jsonassert](https://github.com/kinbiko/jsonassert) | 27 | 0 | 2018-10-27 | 1 month ago | Package for verifying that your JSON payloads are serialized correctly. |
| [hamcrest](https://github.com/rdrdr/hamcrest) | 27 | 3 | 2010-12-22 | 8 years ago | fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results. |
| [bro](https://github.com/marioidival/bro) | 25 | 3 | 2015-08-11 | 3 years ago | Watch files in directory and run tests for them. |
| [dsunit](https://github.com/viant/dsunit) | 25 | 9 | 2016-06-14 | 1 month ago | Datastore testing for SQL, NoSQL, structured files. |
| [assert](https://github.com/go-playground/assert) | 15 | 1 | 2015-07-21 | 1 week ago | Basic Assertion Library used along side native go testing, with building blocks for custom assertions. |
| [testcase](https://github.com/adamluzsi/testcase) | 12 | 1 | 2019-04-23 | 5 hours ago | Idiomatic testing framework for Behavior Driven Development. |
| [gosuite](https://github.com/pavlo/gosuite) | 9 | 1 | 2016-10-16 | 3 years ago | Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests. |
| [gocrest](https://github.com/corbym/gocrest) | 9 | 1 | 2017-12-24 | 1 year ago | Composable hamcrest-like matchers for Go assertions. |
| [badio](https://github.com/cavaliercoder/badio) | 9 | 1 | 2016-02-11 | 3 years ago | Extensions to Go's `testing/iotest` package. |
| [gogiven](https://github.com/corbym/gogiven) | 8 | 4 | 2018-01-01 | 1 year ago | YATSPEC-like BDD testing framework for Go. |
| [flute](https://github.com/suzuki-shunsuke/flute) | 8 | 2 | 2019-07-06 | 1 month ago | HTTP client testing framework. |
| [testsql](https://github.com/zhulongcheng/testsql) | 7 | 2 | 2018-09-22 | 1 month ago | Generate test data from SQL files before testing and clear it after finished. |
| [tt](https://github.com/vcaesar/tt) | 6 | 1 | 2018-04-03 | 1 day ago | Simple and colorful test tools. |
| [biff](https://github.com/fulldump/biff) | 6 | 1 | 2018-03-29 | 1 year ago | Bifurcation testing framework, BDD compatible. |
| [schema](https://github.com/jgroeneveld/schema) | 6 | 3 | 2015-08-13 | 2 weeks ago | Quick and easy expression matching for JSON schemas used in requests and responses. |
| [trial](https://github.com/jgroeneveld/trial) | 4 | 1 | 2015-06-18 | 2 weeks ago | Quick and easy extendable assertions without introducing much boilerplate. |

### Mock
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [chromedp](https://github.com/chromedp/chromedp) | 3884 | 137 | 2017-01-24 | 3 days ago | a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol. |
| [mock](https://github.com/golang/mock) | 3179 | 68 | 2015-06-13 | 4 days ago | Mocking framework for the Go programming language. |
| [go-fuzz](https://github.com/dvyukov/go-fuzz) | 3109 | 83 | 2015-04-15 | 6 days ago | Randomized testing system. |
| [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) | 1958 | 19 | 2014-02-07 | 1 month ago | Mock SQL driver for testing database interactions. |
| [hoverfly](https://github.com/SpectoLabs/hoverfly) | 1482 | 58 | 2015-12-01 | 2 weeks ago | HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI. |
| [selenoid](https://github.com/aerokube/selenoid) | 1345 | 96 | 2016-08-22 | 1 hour ago | alternative Selenium hub server that launches browsers within containers. |
| [gock](https://github.com/h2non/gock) | 881 | 17 | 2016-03-03 | 3 days ago | Versatile HTTP mocking made easy. |
| [httpmock](https://github.com/jarcoal/httpmock) | 638 | 7 | 2014-02-25 | 2 weeks ago | Easy mocking of HTTP responses from external resources. |
| [gofuzz](https://github.com/google/gofuzz) | 599 | 22 | 2014-08-01 | 1 month ago | Library for populating go objects with random values. |
| [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) | 378 | 7 | 2014-05-21 | 3 weeks ago | Tool for generating self-contained mock objects. |
| [cdp](https://github.com/mafredri/cdp) | 377 | 16 | 2017-03-12 | 3 weeks ago | Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it. |
| [minimock](https://github.com/gojuno/minimock) | 279 | 7 | 2016-08-04 | 4 weeks ago | Mock generator for Go interfaces. |
| [ggr](https://github.com/aerokube/ggr) | 217 | 25 | 2016-06-16 | 1 day ago | a lightweight server that routes and proxies Selenium WebDriver requests to multiple Selenium hubs. |
| [tavor](https://github.com/zimmski/tavor) | 213 | 11 | 2014-05-18 | 1 year ago | Generic fuzzing and delta-debugging framework. |
| [go-txdb](https://github.com/DATA-DOG/go-txdb) | 184 | 9 | 2015-07-08 | 2 months ago | Single transaction based database driver mainly for testing purposes. |
| [govcr](https://github.com/seborama/govcr) | 86 | 2 | 2016-07-11 | 1 month ago | HTTP mock for Golang: record and replay HTTP interactions for offline testing. |
| [mockhttp](https://github.com/tv42/mockhttp) | 22 | 1 | 2011-06-12 | 5 years ago | Mock object for Go http.ResponseWriter. |

### Fail injection
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [failpoint](https://github.com/pingcap/failpoint) | 426 | 74 | 2019-04-02 | 3 months ago | An implementation of [failpoints](http://www.freebsd.org/cgi/man.cgi?query=fail) for Golang. |

## Text Processing
        
*Libraries for parsing and manipulating texts.*

### Specific Formats
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [colly](https://github.com/gocolly/colly) | 9105 | 255 | 2017-09-29 | 4 days ago | Fast and Elegant Scraping Framework for Gophers. |
| [goquery](https://github.com/PuerkitoBio/goquery) | 7923 | 266 | 2012-08-29 | 1 week ago | GoQuery brings a syntax and a set of features similar to jQuery to the Go language. |
| [blackfriday](https://github.com/russross/blackfriday) | 4054 | 93 | 2011-05-28 | 2 weeks ago | Markdown processor in Go. |
| [toml](https://github.com/BurntSushi/toml) | 2895 | 81 | 2013-02-26 | 3 weeks ago | TOML configuration format (encoder/decoder with reflection). |
| [sh](https://github.com/mvdan/sh) | 2134 | 44 | 2016-01-16 | 1 week ago | Shell parser and formatter. |
| [go-humanize](https://github.com/dustin/go-humanize) | 1988 | 32 | 2012-01-13 | 3 months ago | Formatters for time, numbers, and memory size to human readable format. |
| [bluemonday](https://github.com/microcosm-cc/bluemonday) | 1309 | 32 | 2013-11-21 | 1 week ago | HTML Sanitizer. |
| [inject](https://github.com/facebookarchive/inject) | 1164 | 44 | 2013-10-21 | 9 months ago | Package inject provides a reflect based injector. |
| [gofeed](https://github.com/mmcdole/gofeed) | 1147 | 33 | 2016-01-23 | 3 weeks ago | Parse RSS and Atom feeds in Go. |
| [inject](https://github.com/facebookgo/inject) | 1098 | 44 | 2013-10-21 | 9 months ago | Package inject provides a reflect based injector. |
| [go-toml](https://github.com/pelletier/go-toml) | 642 | 24 | 2013-02-25 | 2 days ago | Go library for the TOML format with query support and handy cli tools. |
| [commonregex](https://github.com/mingrammer/commonregex) | 562 | 22 | 2017-03-23 | 2 weeks ago | A collection of common regular expressions for Go. |
| [slug](https://github.com/gosimple/slug) | 445 | 10 | 2014-03-31 | 5 days ago | URL-friendly slugify with multiple languages support. |
| [mxj](https://github.com/clbanning/mxj) | 345 | 21 | 2014-02-03 | 2 months ago | Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages. |
| [gographviz](https://github.com/awalterschulze/gographviz) | 322 | 13 | 2015-03-15 | 4 months ago | Parses the Graphviz DOT language. |
| [dataflowkit](https://github.com/slotix/dataflowkit) | 319 | 16 | 2017-02-09 | 3 weeks ago | Web scraping Framework to turn websites into structured data. |
| [gotext](https://github.com/leonelquinteros/gotext) | 238 | 5 | 2016-06-20 | 6 days ago | GNU gettext utilities for Go. |
| [go-runewidth](https://github.com/mattn/go-runewidth) | 223 | 11 | 2013-06-21 | 4 days ago | Functions to get fixed width of the character or string. |
| [htmlquery](https://github.com/antchfx/htmlquery) | 156 | 6 | 2017-12-05 | 1 day ago | An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression. |
| [goq](https://github.com/andrewstuart/goq) | 153 | 7 | 2017-02-20 | 4 months ago | Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery). |
| [go-nmea](https://github.com/adrianmo/go-nmea) | 105 | 6 | 2015-07-22 | 3 weeks ago | NMEA parser library for the Go language. |
| [sdp](https://github.com/gortc/sdp) | 74 | 6 | 2016-05-13 | 2 months ago | SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)]. |
| [align](https://github.com/Guitarbum722/align) | 60 | 5 | 2017-04-30 | 2 years ago | A general purpose application that aligns text. |
| [go-slugify](https://github.com/mozillazg/go-slugify) | 56 | 2 | 2016-07-16 | 3 years ago | Make pretty slug with multiple languages support. |
| [genex](https://github.com/alixaxel/genex) | 54 | 3 | 2015-03-10 | 4 months ago | Count and expand Regular Expressions into all matching Strings. |
| [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) | 49 | 7 | 2016-07-05 | 1 day ago | Editorconfig file parser and manipulator for Go. |
| [goribot](https://github.com/zhshch2002/goribot) | 48 | 4 | 2019-09-08 | 1 week ago | A simple golang spider/scraping framework,build a spider in 3 lines. |
| [guesslanguage](https://github.com/endeveit/guesslanguage) | 45 | 1 | 2014-12-16 | 1 year ago | Functions to determine the natural language of a unicode text. |
| [go-zero-width](https://github.com/trubitsyn/go-zero-width) | 42 | 1 | 2018-06-18 | 6 months ago | Zero-width character detection and removal for Go. |
| [goregen](https://github.com/zach-klippenstein/goregen) | 37 | 2 | 2014-12-27 | 1 month ago | Library for generating random strings from regular expressions. |
| [allot](https://github.com/sbstjn/allot) | 35 | 1 | 2016-10-16 | 5 months ago | Placeholder and wildcard text parsing for CLI tools and bots. |
| [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) | 31 | 3 | 2017-11-16 | 2 weeks ago | Fixed-width text formatting (encoder/decoder with reflection). |
| [gonameparts](https://github.com/polera/gonameparts) | 29 | 0 | 2015-05-17 | 2 months ago | Parses human names into individual name parts. |
| [slugify](https://github.com/avelino/slugify) | 27 | 2 | 2015-04-13 | 1 year ago | Go slugify application that handles string. |
| [go-vcard](https://github.com/emersion/go-vcard) | 26 | 3 | 2017-03-21 | 6 months ago | Parse and format vCard. |
| [did](https://github.com/ockam-network/did) | 25 | 6 | 2018-11-03 | 5 months ago | DID (Decentralized Identifiers) Parser and Stringer in Go. |
| [codetree](https://github.com/aerogo/codetree) | 13 | 2 | 2016-11-26 | 2 days ago | Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure. |
| [enca](https://github.com/endeveit/enca) | 8 | 1 | 2014-12-17 | 3 years ago | Minimal cgo bindings for [libenca](http://cihar.com/software/enca/). |
| [syndfeed](https://github.com/zhengchun/syndfeed) | 6 | 1 | 2017-04-07 | 1 year ago | A syndication feed for Atom 1.0 and RSS 2.0. |
| [bbConvert](https://github.com/CalebQ42/bbConvert) | 5 | 1 | 2016-04-15 | 3 years ago | Converts bbCode to HTML that allows you to add support for custom bbCode tags. |
| [doi](https://github.com/hscells/doi) | 4 | 1 | 2017-08-02 | 2 years ago | Document object identifier (doi) parser in Go. |
| [encoding](https://github.com/mickep76/encoding) | 3 | 0 | 2018-04-07 | 1 year ago | Package provides a generic interface to encoders and decodersa. |
| [ltsv](https://github.com/Wing924/ltsv) | 2 | 1 | 2019-05-12 | 4 months ago | High performance [LTSV (Labeled Tab Separeted Value)](http://ltsv.org/) reader for Go. |

### Utility
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [xurls](https://github.com/mvdan/xurls) | 540 | 13 | 2015-01-12 | 3 days ago | Extract urls from text. |
| [gotabulate](https://github.com/bndr/gotabulate) | 211 | 8 | 2014-08-21 | 2 years ago | Easily pretty-print your tabular data with Go. |
| [radix](https://github.com/yourbasic/radix) | 146 | 5 | 2017-06-09 | 1 year ago | fast string sorting algorithm. |
| [parth](https://github.com/codemodus/parth) | 32 | 3 | 2015-04-07 | 8 months ago | URL path segmentation parsing. |
| [xj2go](https://github.com/stackerzzq/xj2go) | 18 | 2 | 2017-09-19 | 10 months ago | Convert xml or json to go struct. |
| [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) | 15 | 1 | 2018-09-09 | 11 months ago | A sanitization-based swear filter for Go. |
| [kace](https://github.com/codemodus/kace) | 12 | 1 | 2015-06-05 | 1 year ago | Common case conversions covering common initialisms. |
| [parseargs-go](https://github.com/txgruppi/parseargs-go) | 7 | 1 | 2016-02-24 | 2 years ago | string argument parser that understands quotes and backslashes. |
| [TySug](https://github.com/Dynom/TySug) | 3 | 1 | 2018-06-06 | 6 months ago | Alternative suggestions with respect to keyboard layouts. |
| [tagify](https://github.com/zoomio/tagify) | 3 | 1 | 2018-03-20 | 5 hours ago | Produces a set of tags from given source. |
| [textwrap](https://github.com/isbm/textwrap) | 1 | 1 | 2019-07-27 | 2 months ago | Implementation of `textwrap` module from Python. |

## Third-party APIs
        
*Libraries for accessing third party APIs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [aws-sdk-go](https://github.com/aws/aws-sdk-go) | 5273 | 253 | 2014-12-05 | 10 hours ago | The official AWS SDK for the Go programming language. |
| [go-github](https://github.com/google/go-github) | 5151 | 185 | 2013-05-25 | 9 hours ago | Go library for accessing the GitHub REST API v3. |
| [slack](https://github.com/nlopes/slack) | 2567 | 46 | 2015-01-24 | 4 hours ago | Slack API in Go. |
| [google-api-go-client](https://github.com/googleapis/google-api-go-client) | 2019 | 132 | 2014-11-25 | 32 minutes ago | Auto-generated Google APIs for Go. |
| [google-cloud-go](https://github.com/googleapis/google-cloud-go) | 1895 | 213 | 2014-05-09 | 1 hour ago | Google Cloud APIs Go Client Library. |
| [discordgo](https://github.com/bwmarrin/discordgo) | 1036 | 51 | 2015-11-02 | 13 hours ago | Go bindings for the Discord Chat API. |
| [anaconda](https://github.com/ChimeraCoder/anaconda) | 1004 | 21 | 2013-03-05 | 4 months ago | Go client library for the Twitter 1.1 API. |
| [stripe-go](https://github.com/stripe/stripe-go) | 988 | 37 | 2014-06-06 | 3 days ago | Go client for the Stripe API. |
| [go-twitter](https://github.com/dghubble/go-twitter) | 842 | 28 | 2015-04-12 | 3 days ago | Go client library for the Twitter v1.1 APIs. |
| [facebook](https://github.com/huandu/facebook) | 799 | 84 | 2012-07-29 | 2 weeks ago | Go Library that supports the Facebook Graph API. |
| [minio-go](https://github.com/minio/minio-go) | 779 | 36 | 2015-05-02 | 5 hours ago | Minio Go Library for Amazon S3 compatible cloud storage. |
| [githubv4](https://github.com/shurcooL/githubv4) | 554 | 19 | 2017-05-27 | 3 weeks ago | Go library for accessing the GitHub GraphQL API v4. |
| [webhooks](https://github.com/go-playground/webhooks) | 399 | 13 | 2015-10-26 | 1 day ago | Webhook receiver for GitHub and Bitbucket. |
| [geo-golang](https://github.com/codingsince1985/geo-golang) | 321 | 11 | 2014-12-04 | 1 month ago | Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs. |
| [paypal](https://github.com/plutov/paypal) | 319 | 16 | 2015-10-14 | 2 weeks ago | Wrapper for PayPal payment API. |
| [PayPal-Go-SDK](https://github.com/logpacker/PayPal-Go-SDK) | 300 | 14 | 2015-10-14 | 2 months ago | Wrapper for PayPal payment API. |
| [go-marathon](https://github.com/gambol99/go-marathon) | 190 | 13 | 2015-02-11 | 4 days ago | Go library for interacting with Mesosphere's Marathon PAAS. |
| [ethrpc](https://github.com/onrik/ethrpc) | 177 | 10 | 2017-01-24 | 1 month ago | Go bindings for Ethereum JSON RPC API. |
| [gostorm](https://github.com/jsgilmore/gostorm) | 122 | 11 | 2013-07-22 | 2 years ago | GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. |
| [trello](https://github.com/adlio/trello) | 118 | 6 | 2016-09-24 | 2 days ago | Go wrapper for the Trello API. |
| [medium-sdk-go](https://github.com/Medium/medium-sdk-go) | 117 | 121 | 2015-09-27 | 1 year ago | Golang SDK for Medium's OAuth2 API. |
| [hipchat](https://github.com/daneharrigan/hipchat) | 113 | 7 | 2013-04-28 | 2 years ago | A golang package to communicate with HipChat over XMPP. |
| [hipchat](https://github.com/andybons/hipchat) | 109 | 8 | 2012-10-21 | 3 years ago | This project implements a golang client library for the Hipchat API. |
| [go-trending](https://github.com/andygrunwald/go-trending) | 103 | 7 | 2015-07-04 | 5 months ago | Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github. |
| [cachet](https://github.com/andygrunwald/cachet) | 71 | 5 | 2015-10-31 | 1 year ago | Go client library for [Cachet (open source status page system)](https://cachethq.io/). |
| [pushover](https://github.com/gregdel/pushover) | 59 | 3 | 2015-02-19 | 8 months ago | Go wrapper for the Pushover API. |
| [megos](https://github.com/andygrunwald/megos) | 57 | 5 | 2015-10-02 | 1 year ago | Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster. |
| [clarifai-go](https://github.com/Clarifai/clarifai-go) | 57 | 36 | 2015-09-29 | 2 years ago | Go client library for interfacing with the Clarifai API. |
| [wit-go](https://github.com/wit-ai/wit-go) | 53 | 9 | 2018-08-20 | 2 months ago | Go client for wit.ai HTTP API. |
| [igdb](https://github.com/Henry-Sarabia/igdb) | 53 | 2 | 2017-08-24 | 2 weeks ago | Go client for the [Internet Game Database API](https://api.igdb.com/). |
| [gads](https://github.com/emiddleton/gads) | 44 | 7 | 2014-01-20 | 1 month ago | Google Adwords Unofficial API. |
| [go-circleci](https://github.com/jszwedko/go-circleci) | 44 | 3 | 2015-08-15 | 1 week ago | Go client library for interacting with CircleCI's API. |
| [go-amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) | 40 | 1 | 2016-11-15 | 1 year ago | Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html). |
| [go-xkcd](https://github.com/nishanths/go-xkcd) | 38 | 3 | 2016-02-26 | 3 years ago | Go client for the xkcd API. |
| [gomusicbrainz](https://github.com/michiwend/gomusicbrainz) | 38 | 8 | 2014-09-11 | 1 month ago | Go MusicBrainz WS2 client library. |
| [uptimerobot](https://github.com/bitfield/uptimerobot) | 36 | 3 | 2018-05-29 | 3 days ago | Go wrapper and command-line client for the Uptime Robot v2 API. |
| [golyrics](https://github.com/mamal72/golyrics) | 33 | 3 | 2016-11-18 | 1 year ago | Golyrics is a Go library to fetch music lyrics data from the Wikia website. |
| [fcm](https://github.com/maddevsio/fcm) | 32 | 4 | 2017-01-06 | 7 months ago | Go library for Firebase Cloud Messaging. |
| [mixpanel](https://github.com/dukex/mixpanel) | 30 | 3 | 2014-05-20 | 1 year ago | Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications. |
| [gcm](https://github.com/TheOrioli/gcm) | 29 | 3 | 2015-11-10 | 3 years ago | Go library for Google Cloud Messaging. |
| [gcm](https://github.com/Aorioli/gcm) | 29 | 3 | 2015-11-10 | 3 years ago | Go library for Google Cloud Messaging. |
| [translate](https://github.com/nuveo/translate) | 27 | 22 | 2015-07-13 | 3 years ago | Go online translation package. |
| [gami](https://github.com/bit4bit/gami) | 26 | 4 | 2014-05-15 | 1 year ago | Go library for Asterisk Manager Interface. |
| [go-unsplash](https://github.com/hbagdi/go-unsplash) | 22 | 1 | 2017-01-19 | 1 week ago | Go client library for the [Unsplash.com](https://unsplash.com) API. |
| [patreon-go](https://github.com/mxpv/patreon-go) | 20 | 4 | 2017-08-07 | 1 month ago | Go library for Patreon API. |
| [go-shopify](https://github.com/rapito/go-shopify) | 19 | 1 | 2014-10-28 | 1 year ago | Go Library to make CRUD request to the Shopify API. |
| [go-spotify](https://github.com/rapito/go-spotify) | 19 | 1 | 2014-10-30 | 2 years ago | Go Library to access Spotify WEB API. |
| [go-twitch](https://github.com/knspriggs/go-twitch) | 17 | 5 | 2016-06-29 | 2 years ago | Go client for interacting with the Twitch v3 API. |
| [go-myanimelist](https://github.com/nstratos/go-myanimelist) | 16 | 2 | 2015-05-03 | 2 years ago | Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api). |
| [codeship-go](https://github.com/codeship/codeship-go) | 16 | 23 | 2017-09-09 | 6 months ago | Go client library for interacting with Codeship's API v2. |
| [brewerydb](https://github.com/naegelejd/brewerydb) | 16 | 3 | 2015-04-15 | 4 years ago | Go library for accessing the BreweryDB API. |
| [ynab.go](https://github.com/brunomvsouza/ynab.go) | 16 | 1 | 2018-07-13 | 2 months ago | Go wrapper for the YNAB API. |
| [go-steam](https://github.com/sostronk/go-steam) | 16 | 10 | 2014-11-24 | 1 year ago | Go Library to interact with Steam game servers. |
| [textbelt](https://github.com/farmergreg/textbelt) | 15 | 2 | 2015-09-02 | 4 years ago | Go client for the textbelt.com txt messaging API. |
| [simples3](https://github.com/rhnvrm/simples3) | 14 | 0 | 2018-12-06 | 2 days ago | Simple no frills AWS S3 Library using REST with V4 Signing written in Go. |
| [coinpaprika-api-go-client](https://github.com/coinpaprika/coinpaprika-api-go-client) | 12 | 7 | 2018-09-25 | 7 months ago | Go client library for interacting with Coinpaprika's API. |
| [go-tmdb](https://github.com/jbrodriguez/go-tmdb) | 12 | 1 | 2014-10-20 | 4 years ago | Simple golang package to communicate with [themoviedb.org](https://themoviedb.org). |
| [go-google-analytics](https://github.com/chonthu/go-google-analytics) | 11 | 2 | 2015-06-01 | 4 years ago | Simple wrapper for easy google analytics reporting. |
| [smitego](https://github.com/sergiotapia/smitego) | 10 | 0 | 2013-12-11 | 5 years ago | Go package to wraps access to the Smite game API. |
| [go-hacknews](https://github.com/PaulRosset/go-hacknews) | 9 | 2 | 2017-08-11 | 2 years ago | Tiny Go client for HackerNews API. |
| [rrdaclient](https://github.com/Omie/rrdaclient) | 8 | 1 | 2014-09-16 | 5 years ago | Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP. |
| [go-sptrans](https://github.com/sergioaugrod/go-sptrans) | 7 | 2 | 2017-09-11 | 2 years ago | Go client library for the SPTrans Olho Vivo API. |
| [go-google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) | 6 | 0 | 2016-10-24 | 3 years ago | Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/). |
| [lastpass-go](https://github.com/ansd/lastpass-go) | 6 | 0 | 2019-07-11 | 3 weeks ago | Go client library for the [LastPass](https://www.lastpass.com/) API. |
| [gumblr](https://github.com/mattcunningham/gumblr) | 6 | 1 | 2015-07-10 | 3 years ago | Go wrapper for the Tumblr v2 API. |
| [go-zooz](https://github.com/gojuno/go-zooz) | 5 | 16 | 2017-07-04 | 11 months ago | Go client for the Zooz API. |
| [go-here](https://github.com/abdullahselek/go-here) | 5 | 1 | 2019-07-07 | 2 weeks ago | Go client library around the HERE location based APIs. |
| [go-sophos](https://github.com/esurdam/go-sophos) | 5 | 1 | 2018-09-05 | 1 year ago | Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies. |
| [go-chronos](https://github.com/axelspringer/go-chronos) | 3 | 10 | 2017-10-23 | 1 year ago | Go library for interacting with the [Chronos](https://mesos.github. |
| [tripadvisor-golang](https://github.com/mrbenosborne/tripadvisor-golang) | 1 | 1 | 2019-04-16 | 5 days ago | Go wrapper for the TripAdvisor API. |
| [libgoffi](https://github.com/clevabit/libgoffi) | 1 | 2 | 2019-08-04 | 2 months ago | Library adapter toolbox for native [libffi](http://sourceware. |
| [vl-go](https://github.com/verifid/vl-go) | 1 | 1 | 2019-02-09 | 1 week ago | Go client library around the VerifID identity verification layer API. |
| [playlyfe-go-sdk](https://github.com/playlyfe/playlyfe-go-sdk) | 1 | 4 | 2015-05-25 | 3 years ago | The Playlyfe Rest API Go SDK. |

## Utilities
        
*General utilities and tools to make your life easier.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [fzf](https://github.com/junegunn/fzf) | 24762 | 348 | 2013-10-24 | 1 day ago | Command-line fuzzy finder written in Go. |
| [hub](https://github.com/github/hub) | 17731 | 355 | 2009-12-06 | 14 hours ago | wrap git commands with additional functionality to interact with github from the terminal. |
| [delve](https://github.com/go-delve/delve) | 12603 | 374 | 2014-05-21 | 8 hours ago | Go debugger. |
| [ctop](https://github.com/bcicen/ctop) | 9069 | 164 | 2016-12-27 | 1 week ago | [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics. |
| [wuzz](https://github.com/asciimoo/wuzz) | 8405 | 176 | 2017-01-31 | 6 days ago | Interactive cli tool for HTTP inspection. |
| [sqlx](https://github.com/jmoiron/sqlx) | 7166 | 177 | 2013-01-29 | 2 weeks ago | provides a set of extensions on top of the excellent built-in database/sql package. |
| [peco](https://github.com/peco/peco) | 5586 | 131 | 2014-06-06 | 14 hours ago | Simplistic interactive filtering tool. |
| [goreleaser](https://github.com/goreleaser/goreleaser) | 4766 | 74 | 2016-12-22 | 8 hours ago | Deliver Go binaries as fast and easily as possible. |
| [usql](https://github.com/xo/usql) | 4712 | 110 | 2017-03-02 | 1 week ago | usql is a universal command-line interface for SQL databases. |
| [godropbox](https://github.com/dropbox/godropbox) | 3768 | 254 | 2014-06-23 | 6 months ago | Common libraries for writing Go services/applications from Dropbox. |
| [go-torch](https://github.com/uber-archive/go-torch) | 3636 | 2360 | 2015-07-22 | 11 months ago | Stochastic flame graph profiler for Go programs. |
| [go-torch](https://github.com/uber/go-torch) | 3539 | 2403 | 2015-07-22 | 11 months ago | Stochastic flame graph profiler for Go programs. |
| [realize](https://github.com/oxequa/realize) | 3308 | 66 | 2016-07-12 | 1 month ago | Go build system with file watchers and live reload. Run, build and watch file changes with custom paths. |
| [goreporter](https://github.com/360EntSecGroup-Skylar/goreporter) | 2537 | 94 | 2017-03-27 | 1 year ago | Golang tool that does static analysis, unit testing, code review and generate code quality report. |
| [hystrix-go](https://github.com/afex/hystrix-go) | 2169 | 80 | 2013-12-15 | 4 months ago | Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker. |
| [panicparse](https://github.com/maruel/panicparse) | 2165 | 43 | 2015-02-02 | 1 month ago | Groups similar goroutines and colorizes stack dump. |
| [task](https://github.com/go-task/task) | 2055 | 38 | 2017-02-27 | 19 hours ago | simple "Make" alternative. |
| [minify](https://github.com/tdewolff/minify) | 1919 | 51 | 2014-05-21 | 2 weeks ago | Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats. |
| [mmake](https://github.com/tj/mmake) | 1465 | 32 | 2017-02-16 | 5 days ago | Modern Make. |
| [storm](https://github.com/asdine/storm) | 1416 | 47 | 2016-01-10 | 6 days ago | Simple and powerful toolkit for BoltDB. |
| [go-funk](https://github.com/thoas/go-funk) | 1348 | 37 | 2016-12-30 | 3 days ago | Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...). |
| [mole](https://github.com/davrodpin/mole) | 1334 | 33 | 2018-10-04 | 5 days ago | cli app to easily create ssh tunnels. |
| [mc](https://github.com/minio/mc) | 1202 | 46 | 2015-01-16 | 10 hours ago | Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems. |
| [filetype](https://github.com/h2non/filetype) | 992 | 23 | 2015-09-24 | 1 month ago | Small package to infer the file type checking the magic numbers signature. |
| [boilr](https://github.com/tmrts/boilr) | 983 | 28 | 2015-12-20 | 4 weeks ago | Blazingly fast CLI tool for creating projects from boilerplate templates. |
| [mergo](https://github.com/imdario/mergo) | 918 | 14 | 2013-03-12 | 2 weeks ago | Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements. |
| [spinner](https://github.com/briandowns/spinner) | 847 | 13 | 2014-12-13 | 3 weeks ago | Go package to easily provide a terminal spinner with options. |
| [circuitbreaker](https://github.com/rubyist/circuitbreaker) | 812 | 18 | 2014-07-18 | 1 week ago | Circuit Breakers in Go. |
| [gtm](https://github.com/git-time-metric/gtm) | 746 | 29 | 2016-06-20 | 2 months ago | Simple, seamless, lightweight time tracking for Git. |
| [jump](https://github.com/gsamokovarov/jump) | 694 | 14 | 2015-08-17 | 1 day ago | Jump helps you navigate faster by learning your habits. |
| [immortal](https://github.com/immortal/immortal) | 620 | 14 | 2016-07-01 | 3 days ago | \*nix cross-platform (OS agnostic) supervisor. |
| [htcat](https://github.com/htcat/htcat) | 493 | 17 | 2013-08-05 | 8 months ago | Parallel and Pipelined HTTP GET Utility. |
| [go-dry](https://github.com/ungerik/go-dry) | 435 | 14 | 2014-02-28 | 1 year ago | DRY (don't repeat yourself) package for Go. |
| [gopencils](https://github.com/bndr/gopencils) | 424 | 14 | 2014-06-23 | 8 months ago | Small and simple package to easily consume REST APIs. |
| [godaemon](https://github.com/VividCortex/godaemon) | 411 | 32 | 2013-08-02 | 7 months ago | Utility to write daemons. |
| [circuit](https://github.com/cep21/circuit) | 387 | 10 | 2017-12-24 | 6 days ago | An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern. |
| [request](https://github.com/mozillazg/request) | 361 | 14 | 2014-12-21 | 1 month ago | Go HTTP Requests for Humans™. |
| [ergo](https://github.com/cristianoliveira/ergo) | 332 | 5 | 2017-08-20 | 2 months ago | The management of multiple local services running over different ports made easy. |
| [koazee](https://github.com/wesovilabs/koazee) | 319 | 10 | 2018-11-09 | 1 week ago | Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays. |
| [go-rate](https://github.com/beefsack/go-rate) | 294 | 10 | 2014-08-25 | 1 year ago | Timed rate limiter for Go. |
| [gohper](https://github.com/cosiner/gohper) | 248 | 20 | 2015-03-24 | 2 years ago | Various tools/modules help for development. |
| [clockwork](https://github.com/jonboulle/clockwork) | 230 | 5 | 2014-09-10 | 2 months ago | A simple fake clock for golang. |
| [deepcopier](https://github.com/ulule/deepcopier) | 221 | 18 | 2015-07-25 | 2 years ago | Simple struct copying for Go. |
| [serve](https://github.com/syntaqx/serve) | 198 | 4 | 2019-01-11 | 7 months ago | A static http server anywhere you need. |
| [go-trigger](https://github.com/sadlil/go-trigger) | 188 | 12 | 2015-10-19 | 2 years ago | Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project. |
| [retry](https://github.com/kamilsk/retry) | 177 | 1 | 2016-11-03 | 1 month ago | The most advanced functional mechanism to perform actions repetitively until successful. |
| [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) | 162 | 3 | 2016-11-08 | 2 months ago | go:generate tool for wrapping symbols exported by golang plugins (1.8 only). |
| [gubrak](https://github.com/novalagung/gubrak) | 158 | 6 | 2018-03-09 | 3 weeks ago | Golang utility library with syntactic sugar. It's like lodash, but for golang. |
| [mimetype](https://github.com/gabriel-vasile/mimetype) | 155 | 3 | 2018-07-02 | 18 hours ago | Package for MIME type detection based on magic numbers. |
| [rerun](https://github.com/ivpusic/rerun) | 155 | 7 | 2014-12-10 | 1 year ago | Recompiling and rerunning go apps when source changes. |
| [gotenv](https://github.com/subosito/gotenv) | 153 | 4 | 2013-08-27 | 1 month ago | Load environment variables from `.env` or any `io.Reader` in Go. |
| [moldova](https://github.com/StabbyCutyou/moldova) | 147 | 5 | 2016-01-30 | 2 years ago | Utility for generating random data based on an input template. |
| [death](https://github.com/vrecan/death) | 142 | 4 | 2015-03-09 | 1 year ago | Managing go application shutdown with signals. |
| [util](https://github.com/shomali11/util) | 142 | 7 | 2017-05-24 | 4 months ago | Collection of useful utility functions. (strings, concurrency, manipulations, ...). |
| [robustly](https://github.com/VividCortex/robustly) | 139 | 19 | 2013-07-08 | 1 year ago | Runs functions resiliently, catching and restarting panics. |
| [apm](https://github.com/topfreegames/apm) | 130 | 16 | 2015-11-19 | 2 years ago | Process manager for Golang applications with an HTTP API. |
| [chyle](https://github.com/antham/chyle) | 110 | 7 | 2016-11-18 | 2 weeks ago | Changelog generator using a git repository with multiple configuration possibilities. |
| [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) | 108 | 5 | 2015-10-13 | 1 month ago | XML Sitemap generator written in Go. |
| [toolbox](https://github.com/viant/toolbox) | 106 | 13 | 2016-06-14 | 2 weeks ago | Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer. |
| [onecache](https://github.com/adelowo/onecache) | 102 | 7 | 2017-04-15 | 5 months ago | Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc). |
| [lrserver](https://github.com/jaschaephraim/lrserver) | 102 | 5 | 2014-07-15 | 1 year ago | LiveReload server for Go. |
| [go-bsdiff](https://github.com/gabstv/go-bsdiff) | 88 | 1 | 2019-02-24 | 7 months ago | Pure Go bsdiff and bspatch libraries and CLI tools. |
| [pm](https://github.com/VividCortex/pm) | 74 | 19 | 2013-11-18 | 3 days ago | Process (i.e. goroutine) manager with an HTTP API. |
| [unis](https://github.com/esemplastic/unis) | 69 | 4 | 2017-05-06 | 2 years ago | Common Architecture™ for String Utilities in Go. |
| [xferspdy](https://github.com/monmohan/xferspdy) | 69 | 3 | 2015-05-22 | 3 years ago | Xferspdy provides binary diff and patch library in golang. |
| [netbug](https://github.com/e-dard/netbug) | 65 | 1 | 2015-03-06 | 4 years ago | Easy remote profiling of your services. |
| [mssqlx](https://github.com/linxGnu/mssqlx) | 64 | 8 | 2016-12-26 | 7 months ago | Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind. |
| [go-health](https://github.com/Talento90/go-health) | 63 | 2 | 2018-02-14 | 1 year ago | Health package simplifies the way you add health check to your services. |
| [multitick](https://github.com/VividCortex/multitick) | 59 | 19 | 2013-12-11 | 3 days ago | Multiplexor for aligned tickers. |
| [repeat](https://github.com/ssgreg/repeat) | 59 | 4 | 2017-11-22 | 1 week ago | Go implementation of different backoff strategies useful for retrying operations and heartbeating. |
| [minquery](https://github.com/icza/minquery) | 51 | 3 | 2016-11-16 | 1 day ago | MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off). |
| [abutil](https://github.com/bahlo/abutil) | 51 | 3 | 2015-08-26 | 4 years ago | Collection of often-used Golang helpers. |
| [handy](https://github.com/miguelpragier/handy) | 48 | 4 | 2018-06-13 | 2 months ago | Many utilities and helpers like string handlers/formatters and validators. |
| [mimemagic](https://github.com/zRedShift/mimemagic) | 47 | 2 | 2018-10-12 | 10 months ago | Pure Go ultra performant MIME sniffing library/utility. |
| [go-astitodo](https://github.com/asticode/go-astitodo) | 46 | 2 | 2016-10-18 | 1 year ago | Parse TODOs in your GO code. |
| [golog](https://github.com/mlimaloureiro/golog) | 43 | 3 | 2016-01-09 | 9 months ago | Easy and lightweight CLI tool to time track your tasks. |
| [goback](https://github.com/carlescere/goback) | 42 | 1 | 2015-03-14 | 1 year ago | Go simple exponential backoff package. |
| [intrinsic](https://github.com/mengzhuo/intrinsic) | 40 | 3 | 2017-06-13 | 2 years ago | Use x86 SIMD without writing any assembly code. |
| [gaper](https://github.com/maxcnunes/gaper) | 40 | 0 | 2018-06-16 | 9 months ago | Builds and restarts a Go project when it crashes or some watched file changes. |
| [goreadability](https://github.com/philipjkim/goreadability) | 39 | 6 | 2016-04-20 | 6 months ago | Webpage summary extractor using Facebook Open Graph and arc90's readability. |
| [copy-pasta](https://github.com/jutkko/copy-pasta) | 38 | 4 | 2017-01-28 | 2 months ago | Universal multi-workstation clipboard that uses S3 like backend for the storage. |
| [golarm](https://github.com/msempere/golarm) | 35 | 1 | 2015-08-15 | 4 years ago | Fire alarms with system events. |
| [retry](https://github.com/thedevsaddam/retry) | 34 | 1 | 2018-02-26 | 1 year ago | Simple and easy retry mechanism package for Go. |
| [myhttp](https://github.com/inancgumus/myhttp) | 34 | 1 | 2017-09-13 | 1 year ago | Simple API to make HTTP GET requests with timeout support. |
| [gpath](https://github.com/tenntenn/gpath) | 32 | 2 | 2017-05-24 | 2 years ago | Library to simplify access struct fields with Go's expression in reflection. |
| [retry-go](https://github.com/rafaeljesus/retry-go) | 31 | 1 | 2017-06-10 | 1 year ago | Retrying made simple and easy for golang. |
| [goseaweedfs](https://github.com/linxGnu/goseaweedfs) | 31 | 6 | 2017-07-20 | 8 months ago | SeaweedFS client library with almost full features. |
| [rclient](https://github.com/zpatrick/rclient) | 29 | 3 | 2017-02-28 | 1 year ago | Readable, flexible, simple-to-use client for REST APIs. |
| [pgo](https://github.com/arthurkushman/pgo) | 26 | 4 | 2018-12-26 | 3 months ago | Convenient functions for PHP community. |
| [ugo](https://github.com/alxrm/ugo) | 22 | 1 | 2016-02-18 | 3 years ago | ugo is slice toolbox with concise syntax for Go. |
| [goplaceholder](https://github.com/michiwend/goplaceholder) | 21 | 2 | 2014-10-12 | 3 years ago | a small golang lib to generate placeholder images. |
| [generate](https://github.com/go-playground/generate) | 19 | 2 | 2015-11-15 | 2 years ago | runs go generate recursively on a specified path or environment variable and can filter by regex. |
| [evaluator](https://github.com/nullne/evaluator) | 18 | 1 | 2017-04-28 | 1 year ago | Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend. |
| [dbt](https://github.com/nikogura/dbt) | 18 | 1 | 2017-12-01 | 3 months ago | A framework for running self-updating signed binaries from a central, trusted repository. |
| [gostrutils](https://github.com/ik5/gostrutils) | 17 | 3 | 2018-09-19 | 1 month ago | Collections of string manipulation and conversion functions. |
| [scan](https://github.com/blockloop/scan) | 17 | 1 | 2017-11-28 | 2 months ago | Scan golang `sql.Rows` directly to structs, slices, or primitive types. |
| [filter](https://github.com/gookit/filter) | 17 | 4 | 2018-09-26 | 1 month ago | provide filtering, sanitizing, and conversion of Go data. |
| [dlog](https://github.com/kirillDanshin/dlog) | 15 | 2 | 2016-07-05 | 2 years ago | Compile-time controlled logger to make your release smaller without removing debug calls. |
| [okrun](https://github.com/xta/okrun) | 14 | 2 | 2014-10-01 | 5 years ago | go run error steamroller. |
| [filler](https://github.com/yaronsumel/filler) | 14 | 1 | 2017-04-05 | 2 years ago | small utility to fill structs using "fill" tag. |
| [go-httpheader](https://github.com/mozillazg/go-httpheader) | 14 | 2 | 2017-06-24 | 11 months ago | Go library for encoding structs into Header fields. |
| [structs](https://github.com/PumpkinSeed/structs) | 14 | 2 | 2017-08-26 | 2 years ago | Implement simple functions to manipulate structs. |
| [ghokin](https://github.com/antham/ghokin) | 13 | 2 | 2018-08-03 | 2 weeks ago | Parallelized formatter with no external dependencies for gherkin (cucumber, behat...). |
| [rest-go](https://github.com/edermanoel94/rest-go) | 12 | 3 | 2019-07-30 | 4 weeks ago | A package that provide many helpful methods for working with rest api. |
| [fastlz](https://github.com/digitalcrab/fastlz) | 12 | 1 | 2014-12-26 | 4 years ago | Wrap over [FastLz](http://fastlz.org/) (free, open-source, portable real-time compression library) for GoLang. |
| [rerate](https://github.com/abo/rerate) | 12 | 4 | 2016-05-24 | 2 years ago | Redis-based rate counter and rate limiter for Go. |
| [retry](https://github.com/shafreeck/retry) | 10 | 0 | 2018-07-18 | 3 months ago | A pretty simple library to ensure your work to be done. |
| [ctxutil](https://github.com/posener/ctxutil) | 10 | 1 | 2018-07-30 | 2 months ago | A collection of utility functions for contexts. |
| [command](https://github.com/txgruppi/command) | 9 | 2 | 2015-08-24 | 3 years ago | Command pattern for Go with thread safe serial and parallel dispatcher. |
| [mimesniffer](https://github.com/aofei/mimesniffer) | 9 | 1 | 2018-12-20 | 1 month ago | A MIME type sniffer for Go. |
| [backscanner](https://github.com/icza/backscanner) | 9 | 2 | 2017-10-18 | 1 year ago | A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward. |
| [sslice](https://github.com/yaa110/sslice) | 5 | 1 | 2018-11-17 | 3 months ago | Create a slice which is always sorted. |
| [shutdown](https://github.com/ztrue/shutdown) | 5 | 1 | 2018-11-18 | 8 months ago | App shutdown hooks for `os.Signal` handling. |
| [limiters](https://github.com/mennanov/limiters) | 4 | 1 | 2019-08-29 | 2 weeks ago | Rate limiters for distributed applications in Golang with configurable back-ends and distributed locks. |
| [silk](https://github.com/chrispassas/silk) | 4 | 1 | 2018-12-18 | 2 months ago | Read silk netflow files. |
| [slicer](https://github.com/leaanthony/slicer) | 4 | 1 | 2019-01-10 | 3 months ago | Makes working with slices easier. |
| [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) | 3 | 1 | 2019-05-16 | 5 months ago | Go package for working with Problem Details. |
| [sliceconv](https://github.com/Henry-Sarabia/sliceconv) | 3 | 1 | 2019-02-15 | 2 months ago | Slice conversion between primitive types. |
| [retry](https://github.com/percolate/retry) | 3 | 28 | 2018-06-16 | 1 month ago | A simple but highly configurable retry package for Go. |
| [blank](https://github.com/Henry-Sarabia/blank) | 1 | 2 | 2019-02-13 | 2 months ago | Verify or remove blanks and whitespace from strings. |
| [olaf](https://github.com/btnguyen2k/olaf) | 1 | 1 | 2019-01-03 | 6 months ago | Twitter Snowflake implemented in Go. |

## UUID
        
*Libraries for working with UUIDs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [ulid](https://github.com/oklog/ulid) | 1733 | 34 | 2016-12-06 | 1 month ago | Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier). |
| [uuid](https://github.com/gofrs/uuid) | 607 | 16 | 2018-07-13 | 6 days ago | Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid. |
| [wuid](https://github.com/edwingeng/wuid) | 309 | 15 | 2018-01-27 | 4 months ago | An extremely fast unique number generator, 10-135 times faster than UUID. |
| [Goid](https://github.com/JakeHL/Goid) | 25 | 4 | 2017-05-19 | 8 months ago | Generate and Parse RFC4122 compliant V4 UUIDs. |
| [sno](https://github.com/muyo/sno) | 20 | 1 | 2019-05-27 | 3 months ago | Compact, sortable and fast unique IDs with embedded metadata. |
| [nanoid](https://github.com/aidarkhanov/nanoid) | 11 | 1 | 2019-07-02 | 2 months ago | A tiny and efficient Go unique string ID generator. |
| [uuid](https://github.com/agext/uuid) | 10 | 1 | 2016-02-03 | 8 months ago | Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier. |

## Validation
        
*Libraries for validation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [validator](https://github.com/go-playground/validator) | 3889 | 73 | 2015-02-13 | 1 week ago | Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving. |
| [govalidator](https://github.com/asaskevich/govalidator) | 3796 | 96 | 2014-06-20 | 3 weeks ago | Validators and sanitizers for strings, numerics, slices and structs. |
| [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) | 1130 | 20 | 2016-06-22 | 1 month ago | Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags. |
| [govalidator](https://github.com/thedevsaddam/govalidator) | 765 | 17 | 2017-09-14 | 1 week ago | Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. |
| [validate](https://github.com/gookit/validate) | 141 | 7 | 2018-07-16 | 1 month ago | Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features. |
| [checkdigit](https://github.com/osamingo/checkdigit) | 46 | 0 | 2019-04-05 | 1 month ago | Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.). |
| [jio](https://github.com/faceair/jio) | 28 | 2 | 2018-10-28 | 5 months ago | jio is a json schema validator similar to [joi](https://github.com/hapijs/joi). |
| [validate](https://github.com/gobuffalo/validate) | 22 | 7 | 2018-02-11 | 3 weeks ago | This package provides a framework for writing validations for Go applications. |
| [terraform-validator](https://github.com/thazelart/terraform-validator) | 21 | 2 | 2019-05-29 | 1 month ago | A norms and conventions validator for Terraform. |

## Version Control
        
*Libraries for version control.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-git](https://github.com/src-d/go-git) | 4503 | 100 | 2015-10-22 | 3 days ago | highly extensible Git implementation in pure Go. |
| [git2go](https://github.com/libgit2/git2go) | 1394 | 53 | 2013-03-06 | 1 month ago | Go bindings for libgit2. |
| [hercules](https://github.com/src-d/hercules) | 689 | 22 | 2016-12-13 | 2 weeks ago | gaining advanced insights from Git repository history. |
| [go-vcs](https://github.com/sourcegraph/go-vcs) | 72 | 33 | 2013-06-02 | 1 month ago | manipulate and inspect VCS repositories in Go. |
| [gh](https://github.com/rjeczalik/gh) | 71 | 6 | 2015-03-09 | 1 year ago | Scriptable server and net/http middleware for GitHub Webhooks. |
| [hgo](https://github.com/beyang/hgo) | 12 | 4 | 2014-06-18 | 4 years ago | Hgo is a collection of Go packages providing read-access to local Mercurial repositories. |

## Video
        
*Libraries for manipulating video.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [goav](https://github.com/giorgisio/goav) | 883 | 42 | 2015-05-21 | 1 month ago | Comphrensive Go bindings for FFmpeg. |
| [gmf](https://github.com/3d0c/gmf) | 558 | 32 | 2013-04-03 | 3 months ago | Go bindings for FFmpeg av\* libraries. |
| [go-astits](https://github.com/asticode/go-astits) | 270 | 14 | 2017-07-04 | 1 month ago | Parse and demux MPEG Transport Streams (.ts) natively in GO. |
| [go-astisub](https://github.com/asticode/go-astisub) | 194 | 7 | 2016-12-16 | 1 month ago | Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.). |
| [gst](https://github.com/ziutek/gst) | 154 | 10 | 2011-07-26 | 1 year ago | Go bindings for GStreamer. |
| [libvlc-go](https://github.com/adrg/libvlc-go) | 73 | 7 | 2015-01-06 | 2 weeks ago | Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player). |
| [go-m3u8](https://github.com/quangngotan95/go-m3u8) | 42 | 1 | 2018-11-06 | 6 months ago | Parser and generator library for Apple m3u8 playlists. |
| [v4l](https://github.com/korandiz/v4l) | 33 | 4 | 2016-10-25 | 1 year ago | Video capture library for Linux, written in Go. |
| [libgosubs](https://github.com/wargarblgarbl/libgosubs) | 11 | 1 | 2017-05-04 | 10 months ago | Subtitle format support for go. Supports .srt, .ttml, and .ass. |

## Web Frameworks
        
*Full stack web frameworks.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gin](https://github.com/gin-gonic/gin) | 32367 | 1173 | 2014-06-17 | 5 hours ago | Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity. |
| [beego](https://github.com/astaxie/beego) | 22277 | 1262 | 2012-02-29 | 3 days ago | beego is an open-source, high-performance web framework for the Go programming language. |
| [echo](https://github.com/labstack/echo) | 15380 | 533 | 2015-03-02 | 11 hours ago | High performance, minimalist Go web framework. |
| [revel](https://github.com/revel/revel) | 11380 | 562 | 2011-12-09 | 3 months ago | High-productivity web framework for the Go language. |
| [goa](https://github.com/goadesign/goa) | 3584 | 166 | 2014-12-05 | 9 hours ago | Goa provides a holistic approach for developing remote APIs and microservices in Go. |
| [go-json-rest](https://github.com/ant0ine/go-json-rest) | 3357 | 163 | 2013-02-19 | 2 months ago | Quick and easy way to setup a RESTful JSON API. |
| [gizmo](https://github.com/nytimes/gizmo) | 2938 | 111 | 2015-12-16 | 2 weeks ago | Microservice toolkit used by the New York Times. |
| [macaron](https://github.com/go-macaron/macaron) | 2864 | 149 | 2014-07-10 | 1 week ago | Macaron is a high productive and modular design web framework in Go. |
| [utron](https://github.com/gernest/utron) | 2144 | 72 | 2015-09-16 | 1 year ago | Lightweight MVC framework for Go(Golang). |
| [go-tigertonic](https://github.com/rcrowley/go-tigertonic) | 996 | 45 | 2013-02-10 | 1 year ago | Go framework for building JSON web services inspired by Dropwizard. |
| [tango](https://github.com/lunny/tango) | 821 | 77 | 2014-12-17 | 5 months ago | Micro & pluggable web framework for Go. |
| [traffic](https://github.com/gravityblast/traffic) | 517 | 25 | 2013-08-09 | 3 years ago | Sinatra inspired regexp/pattern mux and web framework for Go. |
| [gongular](https://github.com/mustafaakin/gongular) | 416 | 21 | 2016-06-22 | 8 months ago | Fast Go web framework with input mapping/validation and (DI) Dependency Injection. |
| [neo](https://github.com/ivpusic/neo) | 395 | 33 | 2015-02-05 | 1 year ago | Neo is minimal and fast Go Web Framework with extremely simple API. |
| [air](https://github.com/aofei/air) | 349 | 23 | 2016-07-20 | 1 day ago | An ideally refined web framework for Go. |
| [mango](https://github.com/paulbellamy/mango) | 339 | 21 | 2011-05-25 | 2 years ago | Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. |
| [gondola](https://github.com/rainycape/gondola) | 314 | 15 | 2014-07-26 | 8 months ago | The web framework for writing faster sites, faster. |
| [golf](https://github.com/dinever/golf) | 239 | 20 | 2015-11-18 | 2 years ago | Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library. |
| [aero](https://github.com/aerogo/aero) | 189 | 14 | 2016-11-09 | 2 days ago | High-performance web framework for Go, reaches top scores in Lighthouse. |
| [gem](https://github.com/go-gem/gem) | 153 | 20 | 2016-11-10 | 2 years ago | Simple and fast web framework, friendly to REST API. |
| [go-rest](https://github.com/ungerik/go-rest) | 116 | 10 | 2012-07-13 | 2 years ago | Small and evil REST framework for Go. |
| [hiboot](https://github.com/hidevopsio/hiboot) | 100 | 10 | 2018-03-16 | 2 weeks ago | hiboot is a high performance web application framework with auto configuration and dependency injection support. |
| [webgo](https://github.com/bnkamalesh/webgo) | 82 | 3 | 2015-12-16 | 1 week ago | A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc). |
| [golax](https://github.com/fulldump/golax) | 72 | 9 | 2016-01-31 | 1 year ago | A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more. |
| [uadmin](https://github.com/uadmin/uadmin) | 63 | 6 | 2018-10-05 | 3 months ago | Fully featured web framework for Golang, inspired by Django. |
| [microservice](https://github.com/claygod/microservice) | 61 | 7 | 2016-12-15 | 4 months ago | The framework for the creation of microservices, written in Golang. |
| [yarf](https://github.com/yarf-framework/yarf) | 51 | 3 | 2015-09-02 | 7 months ago | Fast micro-framework designed to build REST APIs and web services in a fast and simple way. |
| [fireball](https://github.com/zpatrick/fireball) | 49 | 4 | 2016-07-20 | 1 year ago | More "natural" feeling web framework. |
| [vox](https://github.com/aisk/vox) | 43 | 1 | 2014-12-24 | 2 weeks ago | A golang web framework for humans, inspired by Koa heavily. |
| [patron](https://github.com/beatlabs/patron) | 40 | 14 | 2019-01-30 | 4 days ago | Patron is a microservice framework following best cloud practices with a focus on productivity. |
| [api](https://github.com/resoursea/api) | 29 | 5 | 2015-01-25 | 4 years ago | REST framework for quickly writing resource based services. |
| [rex](https://github.com/goanywhere/rex) | 27 | 1 | 2014-10-16 | 1 year ago | Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`. |
| [goa](https://github.com/goa-go/goa) | 25 | 1 | 2019-07-26 | 1 week ago | goa is just like koajs for golang, it is a flexible, light, high-performance and extensible web framework based on middleware. |
| [nio](https://github.com/go-nio/nio) | 21 | 3 | 2018-12-05 | 2 months ago | Modern, minimal and productive Go HTTP framework. |
| [rux](https://github.com/gookit/rux) | 15 | 1 | 2018-08-05 | 1 month ago | Simple and fast web framework for build golang HTTP applications. |
| [banjo](https://github.com/nsheremet/banjo) | 10 | 1 | 2017-12-09 | 1 year ago | Very simple and fast web framework for Go. |
| [ginrpc](https://github.com/xxjwxc/ginrpc) | 7 | 0 | 2019-06-22 | 2 weeks ago | Gin parameter automatic binding tool,gin rpc tools. |
| [sawsij](https://github.com/jaybill/sawsij) | 2 | 2 | 2016-06-09 | 3 years ago | lightweight, open-source web framework for building high-performance, data-driven web applications. |

### Middlewares
        

#### Actual middlewares
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [tollbooth](https://github.com/didip/tollbooth) | 1354 | 44 | 2015-05-17 | 1 month ago | Rate limit HTTP request handler. |
| [cors](https://github.com/rs/cors) | 1281 | 30 | 2014-10-25 | 2 weeks ago | Easily add CORS capabilities to your API. |
| [limiter](https://github.com/ulule/limiter) | 829 | 26 | 2015-10-02 | 2 weeks ago | Dead simple rate limit middleware for Go. |
| [go-server-timing](https://github.com/mitchellh/go-server-timing) | 760 | 20 | 2018-02-12 | 1 year ago | Add/parse Server-Timing header. |
| [ln-paywall](https://github.com/philippgille/ln-paywall) | 95 | 4 | 2018-06-30 | 8 months ago | Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin). |
| [xff](https://github.com/sebest/xff) | 72 | 2 | 2014-12-22 | 6 months ago | Handle `X-Forwarded-For` header and friends. |
| [formjson](https://github.com/rs/formjson) | 33 | 2 | 2015-03-20 | 3 years ago | Transparently handle JSON input as a standard form POST. |
| [client-timing](https://github.com/posener/client-timing) | 15 | 1 | 2018-02-23 | 8 months ago | An HTTP client for Server-Timing header. |

#### Libraries for creating HTTP middlewares
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [negroni](https://github.com/urfave/negroni) | 6428 | 238 | 2014-05-19 | 2 weeks ago | Idiomatic HTTP middleware for Golang. |
| [alice](https://github.com/justinas/alice) | 1883 | 55 | 2014-05-25 | 11 months ago | Painless middleware chaining for Go. |
| [render](https://github.com/unrolled/render) | 1298 | 40 | 2014-06-11 | 2 months ago | Go package for easily rendering JSON, XML, and HTML template responses. |
| [stats](https://github.com/thoas/stats) | 542 | 17 | 2015-03-06 | 6 months ago | Go middleware that stores various information about your web application. |
| [interpose](https://github.com/carbocation/interpose) | 288 | 13 | 2014-07-20 | 2 years ago | Minimalist net/http middleware for golang. |
| [muxchain](https://github.com/stephens2424/muxchain) | 209 | 5 | 2014-05-04 | 7 months ago | Lightweight middleware for net/http. |
| [renderer](https://github.com/thedevsaddam/renderer) | 175 | 6 | 2017-11-08 | 7 months ago | Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go. |
| [rye](https://github.com/InVisionApp/rye) | 94 | 185 | 2016-10-07 | 1 year ago | Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context. |
| [gores](https://github.com/alioygur/gores) | 85 | 5 | 2015-12-25 | 1 year ago | Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs. |
| [chain](https://github.com/codemodus/chain) | 64 | 6 | 2015-05-15 | 1 year ago | Handler wrapper chaining with scoped data (net/context-based "middleware"). |
| [wrap](https://github.com/go-on/wrap) | 59 | 3 | 2014-02-16 | 1 year ago | Small middlewares package for net/http. |
| [catena](https://github.com/codemodus/catena) | 7 | 1 | 2015-07-31 | 1 year ago | http.Handler wrapper catenation (same API as "chain"). |

### Routers
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [mux](https://github.com/gorilla/mux) | 10353 | 277 | 2012-10-03 | 4 days ago | Powerful URL router and dispatcher for golang. |
| [httprouter](https://github.com/julienschmidt/httprouter) | 10175 | 307 | 2013-12-05 | 3 weeks ago | High performance router. Use this and the standard http handlers to form a very high performance web framework. |
| [chi](https://github.com/go-chi/chi) | 6486 | 173 | 2015-10-16 | 1 week ago | Small, fast and expressive HTTP router built on net/context. |
| [web](https://github.com/gocraft/web) | 1399 | 56 | 2013-11-17 | 4 months ago | Mux and middleware package in Go. |
| [bone](https://github.com/go-zoo/bone) | 1235 | 36 | 2014-11-19 | 5 months ago | Lightning Fast HTTP Multiplexer. |
| [fasthttprouter](https://github.com/buaazp/fasthttprouter) | 777 | 34 | 2015-12-13 | 6 months ago | High performance router forked from `httprouter`. The first router fit for `fasthttp`. |
| [goji](https://github.com/goji/goji) | 774 | 40 | 2015-11-16 | 2 months ago | Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`. |
| [gorouter](https://github.com/xujiajun/gorouter) | 461 | 13 | 2018-01-29 | 1 month ago | A simple and fast HTTP router for Go. |
| [httptreemux](https://github.com/dimfeld/httptreemux) | 395 | 22 | 2014-05-15 | 1 month ago | High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter. |
| [lars](https://github.com/go-playground/lars) | 376 | 16 | 2015-12-25 | 5 months ago | Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks. |
| [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) | 363 | 28 | 2015-10-27 | 11 months ago | An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs. |
| [siesta](https://github.com/VividCortex/siesta) | 352 | 28 | 2014-09-23 | 3 days ago | Composable framework to write middleware and handlers. |
| [vestigo](https://github.com/husobee/vestigo) | 254 | 16 | 2015-09-22 | 1 month ago | Performant, stand-alone, HTTP compliant URL Router for go web applications. |
| [router](https://github.com/gowww/router) | 158 | 6 | 2017-05-25 | 1 year ago | Lightning fast HTTP router fully compatible with the net/http.Handler interface. |
| [alien](https://github.com/gernest/alien) | 109 | 4 | 2016-01-31 | 7 months ago | Lightweight and fast http router from outer space. |
| [violetear](https://github.com/nbari/violetear) | 95 | 3 | 2015-06-20 | 7 months ago | Go HTTP router. |
| [Bxog](https://github.com/claygod/Bxog) | 94 | 8 | 2016-05-19 | 1 week ago | Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters. |
| [pure](https://github.com/go-playground/pure) | 89 | 5 | 2016-09-24 | 5 months ago | Is a lightweight HTTP router that sticks to the std "net/http" implementation. |
| [xmux](https://github.com/rs/xmux) | 87 | 6 | 2015-12-15 | 2 years ago | High performance muxer based on `httprouter` with `net/context` support. |
| [gorouter](https://github.com/vardius/gorouter) | 52 | 3 | 2016-07-14 | 1 week ago | GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`. |
| [bellt](https://github.com/GuilhermeCaruso/bellt) | 40 | 5 | 2019-02-21 | 5 months ago | A simple Go HTTP router. |
| [fastrouter](https://github.com/razonyang/fastrouter) | 18 | 2 | 2017-11-01 | 2 years ago | a fast, flexible HTTP router written in Go. |
| [route](https://github.com/goroute/route) | 5 | 3 | 2019-07-07 | 1 month ago | Simple yet powerful HTTP request multiplexer. |

## Windows
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-ole](https://github.com/go-ole/go-ole) | 574 | 38 | 2011-01-21 | 1 month ago | Win32 OLE implementation for golang. |
| [d3d9](https://github.com/gonutz/d3d9) | 92 | 5 | 2015-12-13 | 10 months ago | Go bindings for Direct3D9. |
| [gosddl](https://github.com/MonaxGT/gosddl) | 1 | 1 | 2018-12-04 | 6 months ago | Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL. |

## XML
        
*Libraries and tools for manipulating XML.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [zek](https://github.com/miku/zek) | 289 | 17 | 2017-11-24 | 3 months ago | Generate a Go struct from XML. |
| [xpath](https://github.com/antchfx/xpath) | 188 | 9 | 2016-10-09 | 1 week ago | XPath package for Go. |
| [xquery](https://github.com/antchfx/xquery) | 149 | 11 | 2016-10-09 | 1 year ago | XQuery lets you extract data from HTML/XML documents using XPath expression. |
| [xml2map](https://github.com/sbabiv/xml2map) | 17 | 0 | 2018-08-07 | 6 months ago | XML to MAP converter written Golang. |
| [XML-Comp](https://github.com/XML-Comp/XML-Comp) | 16 | 1 | 2016-10-26 | 1 year ago | Simple command line XML comparer that generates diffs of folders, files and tags. |
| [xmlwriter](https://github.com/shabbyrobe/xmlwriter) | 7 | 1 | 2017-04-11 | 9 months ago | Procedural XML generation API based on libxml2's xmlwriter module. |

## Server Applications
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [etcd](https://github.com/etcd-io/etcd) | 27913 | 1303 | 2013-07-07 | 1 hour ago | Highly-available key value store for shared configuration and service discovery. |
| [caddy](https://github.com/caddyserver/caddy) | 24652 | 704 | 2015-01-14 | 19 hours ago | Caddy is an alternative, HTTP/2 web server that's easy to configure and use. |
| [caddy](https://github.com/mholt/caddy) | 22379 | 675 | 2015-01-14 | 4 months ago | Caddy is an alternative, HTTP/2 web server that's easy to configure and use. |
| [minio](https://github.com/minio/minio) | 18642 | 452 | 2015-01-15 | 23 minutes ago | Minio is a distributed object storage server. |
| [roadrunner](https://github.com/spiral/roadrunner) | 3577 | 143 | 2017-12-27 | 3 days ago | High-performance PHP application server, load-balancer and process manager. |
| [devd](https://github.com/cortesi/devd) | 2866 | 69 | 2015-09-28 | 2 months ago | Local webserver for developers. |
| [algernon](https://github.com/xyproto/algernon) | 1617 | 51 | 2015-03-10 | 3 weeks ago | HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber. |
| [sftpgo](https://github.com/drakkan/sftpgo) | 1101 | 24 | 2019-07-20 | 1 day ago | Full featured and highly configurable SFTP server software. |
| [flipt](https://github.com/markphelps/flipt) | 1031 | 10 | 2016-11-05 | 3 hours ago | A self contained feature flag solution written in Go and Vue. |
| [flagr](https://github.com/checkr/flagr) | 909 | 64 | 2017-10-04 | 6 days ago | Flagr is an open-source feature flagging and A/B testing service. |
| [fider](https://github.com/getfider/fider) | 859 | 21 | 2017-01-18 | 1 day ago | Fider is an open platform to collect and organize customer feedback. |
| [jackal](https://github.com/ortuman/jackal) | 742 | 36 | 2017-11-14 | 6 days ago | An XMPP server written in Go. |
| [discovery](https://github.com/bilibili/discovery) | 730 | 41 | 2018-04-20 | 13 hours ago | A registry for resilient mid-tier load balancing and failover. |
| [dudeldu](https://github.com/krotik/dudeldu) | 98 | 3 | 2016-09-08 | 1 month ago | A simple SHOUTcast server. |
| [psql-streamer](https://github.com/blind-oracle/psql-streamer) | 9 | 4 | 2019-04-29 | 5 months ago | Stream database events from PostgreSQL to Kafka. |
| [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) | 6 | 1 | 2018-10-23 | 6 months ago | Nginx log parser and exporter to Prometheus. |
| [riemann-relay](https://github.com/blind-oracle/riemann-relay) | 0 | 1 | 2019-04-23 | 1 week ago | Relay to load-balance Riemann events and/or convert them to Carbon. |

## Performance
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [jaeger](https://github.com/jaegertracing/jaeger) | 9322 | 314 | 2016-04-16 | 1 hour ago | A distributed tracing system. |
| [profile](https://github.com/pkg/profile) | 1096 | 35 | 2014-10-22 | 3 days ago | Simple profiling support package for Go. |
| [tracer](https://github.com/kamilsk/tracer) | 14 | 0 | 2019-06-22 | 1 month ago | Simple, lightweight tracing. |

## Stream Processing
        
*Libraries and tools for stream processing and reactive programming.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-streams](https://github.com/reugn/go-streams) | 217 | 9 | 2019-05-01 | 3 weeks ago | Go stream processing library. |

# Tools
        
*Go software and plugins.*

## Code Analysis
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gometalinter](https://github.com/alecthomas/gometalinter) | 3560 | 74 | 2014-08-05 | 7 months ago | Metalinter is a tool to automatically apply all static analysis tool and report their output in normalized form. |
| [lint](https://github.com/golang/lint) | 3265 | 111 | 2013-06-03 | 2 weeks ago | Golint is a linter for Go source code. |
| [errcheck](https://github.com/kisielk/errcheck) | 1352 | 20 | 2013-02-25 | 1 week ago | Errcheck is a program for checking for unchecked errors in Go programs. |
| [gcvis](https://github.com/davecheney/gcvis) | 931 | 34 | 2014-07-10 | 7 months ago | Visualise Go program GC trace data in real time. |
| [php-parser](https://github.com/z7zmey/php-parser) | 676 | 27 | 2017-11-07 | 17 hours ago | A Parser for PHP written in Go. |
| [go-critic](https://github.com/go-critic/go-critic) | 615 | 19 | 2018-05-05 | 5 hours ago | source code linter that brings checks that are currently not implemented in other linters. |
| [goast-viewer](https://github.com/yuroyoro/goast-viewer) | 393 | 13 | 2014-06-30 | 5 months ago | Web based Golang AST visualizer. |
| [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) | 293 | 8 | 2017-04-13 | 1 year ago | go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects. |
| [unconvert](https://github.com/mdempsky/unconvert) | 271 | 7 | 2016-02-20 | 1 month ago | Remove unnecessary type conversions from Go source. |
| [gostatus](https://github.com/shurcooL/gostatus) | 241 | 7 | 2013-11-27 | 8 months ago | Command line tool, shows the status of repositories that contain Go packages. |
| [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) | 225 | 5 | 2019-04-19 | 2 weeks ago | An easy way to find outdated dependencies of your Go projects. |
| [dupl](https://github.com/mibk/dupl) | 178 | 8 | 2015-05-20 | 1 year ago | Tool for code clone detection. |
| [apicompat](https://github.com/bradleyfalzon/apicompat) | 166 | 7 | 2016-07-10 | 2 years ago | Checks recent changes to a Go project for backwards incompatible changes. |
| [checkstyle](https://github.com/qiniu/checkstyle) | 95 | 10 | 2014-01-01 | 7 months ago | checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments. |
| [lint](https://github.com/surullabs/lint) | 63 | 5 | 2016-07-09 | 1 year ago | Run linters as part of go test. |
| [validate](https://github.com/mccoyst/validate) | 62 | 6 | 2013-11-22 | 3 years ago | Automatically validates struct fields with tags. |
| [goplantuml](https://github.com/jfeliu007/goplantuml) | 51 | 3 | 2019-05-26 | 9 minutes ago | Library and CLI that generates text plantump class diagram containing information about structures and interfaces with the relationship among them. |
| [go-outdated](https://github.com/firstrow/go-outdated) | 45 | 1 | 2015-06-29 | 9 months ago | Console application that displays outdated packages. |
| [blanket](https://github.com/verygoodsoftwarenotvirus/blanket) | 14 | 2 | 2017-09-04 | 1 year ago | tarp finds functions and methods without direct unit tests in Go source code. |

## Editor Plugins
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [vim-go](https://github.com/fatih/vim-go) | 11192 | 293 | 2014-03-24 | 4 days ago | Go development plugin for Vim. |
| [vscode-go](https://github.com/microsoft/vscode-go) | 5359 | 219 | 2015-10-15 | 2 days ago | Extension for Visual Studio Code (VS Code) which provides support for the Go language. |
| [vscode-go](https://github.com/Microsoft/vscode-go) | 4797 | 213 | 2015-10-15 | 5 months ago | Extension for Visual Studio Code (VS Code) which provides support for the Go language. |
| [gocode](https://github.com/nsf/gocode) | 4782 | 200 | 2010-07-05 | 8 months ago | Autocompletion daemon for the Go programming language. |
| [GoSublime](https://github.com/DisposaBoy/GoSublime) | 3270 | 118 | 2011-08-28 | 6 days ago | Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features. |
| [go-plus](https://github.com/joefitzgerald/go-plus) | 1491 | 45 | 2014-03-14 | 4 hours ago | Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting. |
| [go-mode.el](https://github.com/dominikh/go-mode.el) | 989 | 57 | 2013-01-31 | 4 days ago | Go mode for GNU/Emacs. |
| [Watch](https://github.com/eaburns/Watch) | 170 | 12 | 2013-08-09 | 1 year ago | Runs a command in an acme win on file changes. |
| [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) | 82 | 5 | 2012-11-26 | 3 years ago | Vim plugin to highlight syntax errors on save. |
| [go-language-server](https://github.com/theia-ide/go-language-server) | 32 | 4 | 2017-11-21 | 7 months ago | A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol. |
| [gounit-vim](https://github.com/hexdigest/gounit-vim) | 17 | 2 | 2018-02-22 | 1 year ago | Vim plugin for generating Go tests based on the function's or method's signature. |
| [velour](https://github.com/velour/velour) | 16 | 8 | 2014-06-27 | 1 year ago | IRC client for the acme editor. |
| [theia-go-extension](https://github.com/theia-ide/theia-go-extension) | 12 | 4 | 2017-11-30 | 7 months ago | Go language support for the Theia IDE. |

## Go Generate Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gotests](https://github.com/cweill/gotests) | 2342 | 64 | 2016-01-19 | 1 day ago | Generate Go tests from your source code. |
| [genny](https://github.com/cheekybits/genny) | 1035 | 21 | 2014-10-28 | 1 month ago | Elegant generics for Go. |
| [re2dfa](https://github.com/opennota/re2dfa) | 175 | 9 | 2015-06-20 | 1 year ago | Transform regular expressions into finite state machines and output Go source code. |
| [gocontracts](https://github.com/Parquery/gocontracts) | 54 | 5 | 2018-08-14 | 9 months ago | brings design-by-contract to Go by synchronizing the code with the documentation. |
| [generic](https://github.com/usk81/generic) | 31 | 3 | 2016-06-15 | 7 months ago | flexible data type for Go. |
| [gounit](https://github.com/hexdigest/gounit) | 31 | 4 | 2018-02-05 | 1 year ago | Generate Go tests using your own templates. |
| [hasgo](https://github.com/DylanMeeus/hasgo) | 26 | 4 | 2019-05-17 | 3 hours ago | Generate Haskell inspired functions for your slices. |

## Go Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-swagger](https://github.com/go-swagger/go-swagger) | 4330 | 114 | 2014-11-17 | 5 days ago | Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API. |
| [OctoLinker](https://github.com/OctoLinker/OctoLinker) | 3938 | 87 | 2013-12-28 | 4 hours ago | Navigate through go files efficiently with the OctoLinker browser extension for GitHub. |
| [go-callvis](https://github.com/TrueFurby/go-callvis) | 2141 | 68 | 2016-09-03 | 1 month ago | Visualize call graph of your Go program using dot format. |
| [richgo](https://github.com/kyoh86/richgo) | 407 | 4 | 2017-01-05 | 2 months ago | Enrich `go test` outputs with text decorations. |
| [depth](https://github.com/KyleBanks/depth) | 403 | 7 | 2017-03-04 | 3 months ago | Visualize dependency trees of any package by analyzing imports. |
| [rts](https://github.com/galeone/rts) | 189 | 3 | 2016-04-04 | 2 years ago | RTS: response to struct. Generates Go structs from server responses. |
| [godbg](https://github.com/tylerwince/godbg) | 160 | 5 | 2019-01-24 | 6 months ago | Implementation of Rusts `dbg!` macro for quick and easy debugging during development. |
| [colorgo](https://github.com/songgao/colorgo) | 99 | 6 | 2013-02-15 | 3 years ago | Wrapper around `go` command for colorized `go build` output. |
| [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) | 37 | 2 | 2015-05-22 | 1 year ago | Bash completion for go and wgo. |
| [generator-go-lang](https://github.com/axelspringer/generator-go-lang) | 14 | 12 | 2017-09-13 | 4 months ago | A [Yeoman](http://yeoman.io) generator to get new Go projects started. |

## Software Packages
        
*Software written in Go.*

### DevOps Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [kubernetes](https://github.com/kubernetes/kubernetes) | 59360 | 3049 | 2014-06-07 | 10 minutes ago | Container Cluster Manager from Google. |
| [moby](https://github.com/moby/moby) | 55333 | 3242 | 2013-01-19 | 1 hour ago | Collaborative project for the container ecosystem to assemble container-based systems. |
| [traefik](https://github.com/containous/traefik) | 25534 | 688 | 2015-09-14 | 3 minutes ago | Reverse proxy and load balancer with support for multiple backends. |
| [gitea](https://github.com/go-gitea/gitea) | 16597 | 437 | 2016-11-01 | 7 minutes ago | Fork of Gogs, entirely community driven. |
| [vegeta](https://github.com/tsenart/vegeta) | 12817 | 293 | 2013-08-13 | 5 hours ago | HTTP load testing tool and library. It's over 9000! |
| [packer](https://github.com/hashicorp/packer) | 9475 | 412 | 2013-03-23 | 1 hour ago | Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. |
| [hey](https://github.com/rakyll/hey) | 6895 | 147 | 2016-09-02 | 1 week ago | Hey is a tiny program that sends some load to a web application. |
| [gvm](https://github.com/moovweb/gvm) | 4682 | 151 | 2011-12-03 | 1 week ago | GVM provides an interface to manage Go versions. |
| [webhook](https://github.com/adnanh/webhook) | 4327 | 138 | 2015-01-13 | 1 week ago | Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server. |
| [gaia](https://github.com/gaia-pipeline/gaia) | 3823 | 100 | 2017-12-28 | 6 hours ago | Build powerful pipelines in any programming language. |
| [gox](https://github.com/mitchellh/gox) | 3479 | 70 | 2013-11-17 | 1 month ago | Dead simple, no frills Go cross compile tool. |
| [bosun](https://github.com/bosun-monitor/bosun) | 2878 | 155 | 2013-11-15 | 19 hours ago | Time Series Alerting Framework. |
| [bombardier](https://github.com/codesenberg/bombardier) | 1854 | 49 | 2016-05-29 | 7 months ago | Fast cross-platform HTTP benchmarking tool. |
| [aptly](https://github.com/aptly-dev/aptly) | 1721 | 96 | 2013-12-14 | 5 months ago | aptly is a Debian repository management tool. |
| [goxc](https://github.com/laher/goxc) | 1633 | 47 | 2013-02-11 | 4 weeks ago | build tool for Go, with a focus on cross-compiling and packaging. |
| [fac](https://github.com/mkchoi212/fac) | 1621 | 29 | 2017-12-30 | 2 weeks ago | Command-line user interface to fix git merge conflicts. |
| [kala](https://github.com/ajvb/kala) | 1371 | 64 | 2015-03-19 | 3 days ago | Simplistic, modern, and performant job scheduler. |
| [statusok](https://github.com/sanathp/statusok) | 1209 | 49 | 2015-08-27 | 3 weeks ago | Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected. |
| [script](https://github.com/bitfield/script) | 1202 | 22 | 2019-04-20 | 1 week ago | Making it easy to write shell-like scripts in Go for DevOps and system administration tasks. |
| [banshee](https://github.com/eleme/banshee) | 1053 | 59 | 2015-12-03 | 1 year ago | Anomalies detection system for periodic metrics. |
| [s3gof3r](https://github.com/rlmcpherson/s3gof3r) | 1016 | 34 | 2013-08-02 | 10 months ago | Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3. |
| [go-selfupdate](https://github.com/sanbornm/go-selfupdate) | 683 | 25 | 2013-11-13 | 2 months ago | Enable your Go applications to self update. |
| [pomerium](https://github.com/pomerium/pomerium) | 650 | 12 | 2019-01-01 | 1 week ago | Pomerium is an identity-aware access proxy. |
| [skm](https://github.com/TimothyYe/skm) | 564 | 20 | 2017-10-11 | 1 month ago | SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily! |
| [scaleway-cli](https://github.com/scaleway/scaleway-cli) | 542 | 43 | 2015-03-20 | 20 minutes ago | Manage BareMetal Servers from Command Line (as easily as with Docker). |
| [aurora](https://github.com/xuri/aurora) | 421 | 27 | 2016-10-09 | 6 days ago | Cross-platform web-based Beanstalkd queue server console. |
| [govvv](https://github.com/ahmetb/govvv) | 406 | 10 | 2016-08-03 | 1 year ago | “go build” wrapper to easily add version information into Go binaries. |
| [gonative](https://github.com/inconshreveable/gonative) | 313 | 7 | 2014-05-01 | 3 years ago | Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages. |
| [mora](https://github.com/emicklei/mora) | 267 | 23 | 2013-07-12 | 2 years ago | REST server for accessing MongoDB documents and meta data. |
| [lstags](https://github.com/ivanilves/lstags) | 235 | 14 | 2017-08-15 | 1 week ago | Tool and API to sync Docker images across different registries. |
| [dogo](https://github.com/liudng/dogo) | 220 | 19 | 2014-11-19 | 7 months ago | Monitoring changes in the source file and automatically compile and run (restart). |
| [godbg](https://github.com/sirnewton01/godbg) | 218 | 17 | 2013-08-09 | 1 year ago | Web-based gdb front-end application. |
| [pewpew](https://github.com/bengadbois/pewpew) | 213 | 7 | 2016-10-13 | 3 months ago | Flexible HTTP command line stress tester. |
| [manssh](https://github.com/xwjdsh/manssh) | 208 | 3 | 2017-10-08 | 1 year ago | manssh is a command line tool for managing your ssh alias config easily. |
| [gobrew](https://github.com/cryptojuice/gobrew) | 176 | 5 | 2013-11-13 | 2 years ago | gobrew lets you easily switch between multiple versions of go. |
| [blast](https://github.com/dave/blast) | 170 | 3 | 2017-10-22 | 1 year ago | A simple tool for API load testing and batch jobs. |
| [ostent](https://github.com/ostrost/ostent) | 165 | 5 | 2014-03-31 | 1 year ago | collects and displays system metrics and optionally relays to Graphite and/or InfluxDB. |
| [grapes](https://github.com/yaronsumel/grapes) | 141 | 6 | 2016-09-01 | 1 month ago | Lightweight tool designed to distribute commands over ssh with ease. |
| [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) | 108 | 3 | 2017-03-03 | 3 months ago | Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`. |
| [kcli](https://github.com/cswank/kcli) | 90 | 6 | 2017-03-26 | 1 week ago | Command line tool for inspecting kafka topics/partitions/messages. |
| [terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi) | 78 | 8 | 2017-10-17 | 3 days ago | Terraform provider plugin that dynamically configures itself at runtime based on an OpenAPI document (formerly known as swagger file) containing the definitions of the APIs exposed. |
| [winrm-cli](https://github.com/masterzen/winrm-cli) | 74 | 5 | 2016-05-23 | 6 months ago | Cli tool to remotely execute commands on Windows machines. |
| [go-furnace](https://github.com/go-furnace/go-furnace) | 72 | 1 | 2016-10-09 | 1 month ago | Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean. |
| [jenkins-cli](https://github.com/jenkins-zh/jenkins-cli) | 70 | 8 | 2019-06-21 | 1 day ago | Jenkins CLI allows you manage your Jenkins as an easy way. |
| [drone-scp](https://github.com/appleboy/drone-scp) | 64 | 2 | 2016-10-16 | 1 week ago | Copy files and artifacts via SSH using a binary, docker or Drone CI. |
| [dockerfile-generator](https://github.com/ozankasikci/dockerfile-generator) | 51 | 2 | 2019-08-15 | 2 weeks ago | A go library and an executable that produces valid Dockerfiles using various input channels. |
| [dropship](https://github.com/ChrisMcKenzie/dropship) | 46 | 3 | 2015-09-04 | 1 year ago | Tool for deploying code via cdn. |
| [rodent](https://github.com/alouche/rodent) | 30 | 2 | 2014-06-02 | 2 years ago | Rodent helps you manage Go versions, projects and track dependencies. |
| [drone-jenkins](https://github.com/appleboy/drone-jenkins) | 26 | 1 | 2016-10-15 | 1 week ago | Trigger downstream Jenkins jobs using a binary, docker or Drone CI. |
| [awsenv](https://github.com/soniah/awsenv) | 22 | 1 | 2015-08-05 | 1 year ago | Small binary that loads Amazon (AWS) environment variables for a profile. |
| [depcharge](https://github.com/centerorbit/depcharge) | 9 | 3 | 2018-07-25 | 5 days ago | Helps orchestrating the execution of commands across the many dependencies in larger projects. |
| [lwc](https://github.com/timdp/lwc) | 8 | 2 | 2018-04-22 | 1 year ago | A live-updating version of the UNIX wc command. |
| [sg](https://github.com/ChristopherRabotin/sg) | 5 | 1 | 2015-08-19 | 3 years ago | Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response. |
| [aptly-fork](https://github.com/smira/aptly-fork) | 0 | 0 | 2019-07-05 | 1 month ago | aptly is a Debian repository management tool. |

### Other Software
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [goreplay](https://github.com/buger/goreplay) | 11720 | 453 | 2013-05-30 | 6 days ago | Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time. |
| [rkt](https://github.com/rkt/rkt) | 8756 | 452 | 2014-11-12 | 2 months ago | App Container runtime that integrates with init systems, is compatible with other container formats like Docker, and supports alternative execution engines like KVM. |
| [seaweedfs](https://github.com/chrislusf/seaweedfs) | 8671 | 497 | 2014-07-15 | 1 day ago | Fast, Simple and Scalable Distributed File System with O(1) disk seek. |
| [restic](https://github.com/restic/restic) | 7966 | 230 | 2014-04-27 | 7 hours ago | De-duplicating backup program. |
| [confd](https://github.com/kelseyhightower/confd) | 6614 | 253 | 2013-10-01 | 2 weeks ago | Manage local application configuration files using templates and data from etcd or consul. |
| [comcast](https://github.com/tylertreat/comcast) | 6259 | 152 | 2014-11-12 | 3 months ago | Simulate bad network connections. |
| [liteide](https://github.com/visualfc/liteide) | 5653 | 377 | 2012-11-19 | 1 day ago | LiteIDE is a simple, open source, cross-platform Go IDE. |
| [drive](https://github.com/odeke-em/drive) | 5090 | 203 | 2014-11-03 | 4 months ago | Google Drive client for the commandline. |
| [nes](https://github.com/fogleman/nes) | 4241 | 149 | 2015-03-03 | 4 months ago | Nintendo Entertainment System (NES) emulator written in Go. |
| [toxiproxy](https://github.com/Shopify/toxiproxy) | 4133 | 283 | 2014-09-04 | 1 month ago | Proxy to simulate network and system conditions for automated tests. |
| [pipe](https://github.com/b3log/pipe) | 3717 | 114 | 2017-09-11 | 3 weeks ago | A small and beautiful blogging platform. |
| [duplicacy](https://github.com/gilbertchen/duplicacy) | 2756 | 93 | 2016-02-23 | 56 minutes ago | A cross-platform network and cloud backup tool based on the idea of lock-free deduplication. |
| [croc](https://github.com/schollz/croc) | 2570 | 53 | 2017-10-17 | 4 days ago | Easily and securely send files or folders from one computer to another. |
| [mylg](https://github.com/mehrdadrad/mylg) | 2226 | 107 | 2016-06-22 | 1 year ago | Command Line Network Diagnostic tool written in Go. |
| [goboy](https://github.com/Humpheh/goboy) | 2144 | 39 | 2017-08-20 | 4 months ago | Nintendo Game Boy Color emulator written in Go. |
| [sup](https://github.com/pressly/sup) | 2021 | 73 | 2015-02-24 | 3 weeks ago | Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers. |
| [lgo](https://github.com/yunabe/lgo) | 1865 | 46 | 2017-10-05 | 3 months ago | Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility. |
| [snap](https://github.com/intelsdi-x/snap) | 1803 | 145 | 2014-08-14 | 10 months ago | Powerful telemetry framework. |
| [circuit](https://github.com/gocircuit/circuit) | 1796 | 138 | 2014-04-11 | 5 months ago | Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications. |
| [borg](https://github.com/ok-borg/borg) | 1430 | 41 | 2016-09-11 | 1 year ago | Terminal based search engine for bash snippets. |
| [scc](https://github.com/boyter/scc) | 1058 | 16 | 2018-03-01 | 19 hours ago | Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates. |
| [community](https://github.com/documize/community) | 913 | 38 | 2016-04-30 | 1 hour ago | Modern wiki software that integrates data from SaaS tools. |
| [Go-Package-Store](https://github.com/shurcooL/Go-Package-Store) | 880 | 21 | 2014-01-24 | 2 months ago | App that displays updates for the Go packages in your GOPATH. |
| [leaps](https://github.com/Jeffail/leaps) | 646 | 26 | 2014-06-20 | 7 months ago | Pair programming service using Operational Transforms. |
| [peg](https://github.com/pointlander/peg) | 636 | 32 | 2010-04-26 | 3 months ago | Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator. |
| [vflow](https://github.com/VerizonDigital/vflow) | 618 | 81 | 2017-02-25 | 2 days ago | High-performance, scalable and reliable IPFIX, sFlow and Netflow collector. |
| [gfile](https://github.com/Antonito/gfile) | 514 | 10 | 2019-03-08 | 6 months ago | Securely transfer files between two computers, without any third party, over WebRTC. |
| [godns](https://github.com/TimothyYe/godns) | 486 | 22 | 2014-05-11 | 1 month ago | A dynamic DNS client tool, supports DNSPod & HE.net, written in Go. |
| [shell2http](https://github.com/msoap/shell2http) | 458 | 21 | 2015-03-12 | 1 day ago | Executing shell commands via http server (for prototyping or remote control). |
| [mockingjay-server](https://github.com/quii/mockingjay-server) | 424 | 11 | 2015-04-05 | 1 month ago | Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests. |
| [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) | 378 | 18 | 2015-10-09 | 2 months ago | Video streaming torrent client. |
| [gocc](https://github.com/goccmack/gocc) | 362 | 20 | 2015-06-05 | 5 days ago | Gocc is a compiler kit for Go written in Go. |
| [wellington](https://github.com/wellington/wellington) | 292 | 13 | 2014-12-09 | 1 year ago | Sass project management tool, extends the language with sprite functions (like Compass). |
| [ipe](https://github.com/dimiro1/ipe) | 285 | 16 | 2015-01-13 | 10 months ago | Open source Pusher server implementation compatible with Pusher client libraries written in GO. |
| [IDE](https://github.com/thestrukture/IDE) | 252 | 14 | 2017-09-10 | 2 months ago | Browser accessible IDE. Designed for Go with Go. |
| [cherry](https://github.com/rafael-santiago/cherry) | 198 | 12 | 2015-10-25 | 2 years ago | Tiny webchat server in Go. |
| [orange-cat](https://github.com/utatti/orange-cat) | 180 | 6 | 2014-11-01 | 9 months ago | Markdown previewer written in Go. |
| [orbit](https://github.com/gulien/orbit) | 131 | 8 | 2017-05-13 | 2 weeks ago | A simple tool for running commands and generating files from templates. |
| [joincap](https://github.com/assafmo/joincap) | 125 | 7 | 2018-06-01 | 5 months ago | Command-line utility for merging multiple pcap files together. |
| [ddns](https://github.com/skibish/ddns) | 111 | 5 | 2017-03-14 | 2 weeks ago | Personal DDNS client with Digital Ocean Networking DNS as backend. |
| [boxed](https://github.com/tejo/boxed) | 72 | 2 | 2015-04-19 | 1 year ago | Dropbox based blog engine. |
| [naclpipe](https://github.com/unix4fun/naclpipe) | 20 | 5 | 2015-05-06 | 11 months ago | Simple NaCL EC25519 based crypto pipe tool written in Go. |
| [term-quiz](https://github.com/crazcalm/term-quiz) | 17 | 1 | 2017-12-26 | 1 year ago | Quizzes for your terminal. |
| [snitch](https://github.com/lucasgomide/snitch) | 15 | 1 | 2017-04-07 | 1 year ago | Simple way to notify your team and many tools when someone has deployed any application via Tsuru. |
| [GoDocTooltip](https://github.com/diankong/GoDocTooltip) | 12 | 3 | 2016-01-21 | 3 years ago | Chrome extension for Go Doc sites, which shows function description as tooltip at function list. |

# Resources
        
*Where to discover new Go libraries.*

## Benchmarks
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) | 1286 | 59 | 2013-12-17 | 3 days ago | Go HTTP request router benchmark and comparison. |
| [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) | 1053 | 84 | 2016-04-06 | 1 week ago | Go web framework benchmark. |
| [skynet](https://github.com/atemerev/skynet) | 933 | 47 | 2016-02-14 | 5 months ago | Skynet 1M threads microbenchmark. |
| [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) | 901 | 40 | 2013-01-19 | 1 month ago | Benchmarks of Go serialization methods. |
| [speedtest-resize](https://github.com/fawick/speedtest-resize) | 176 | 8 | 2013-09-16 | 2 years ago | Compare various Image resize algorithms for the Go language. |
| [go-benchmarks](https://github.com/tylertreat/go-benchmarks) | 128 | 9 | 2016-02-25 | 3 years ago | Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches. |
| [gospeed](https://github.com/feyeleanor/gospeed) | 96 | 6 | 2011-05-24 | 1 year ago | Go micro-benchmarks for calculating the speed of language constructs. |
| [autobench](https://github.com/davecheney/autobench) | 89 | 8 | 2013-03-28 | 5 years ago | Framework to compare the performance between different Go versions. |
| [gocostmodel](https://github.com/mna/gocostmodel) | 52 | 5 | 2014-12-19 | 4 years ago | Benchmarks of common basic operations for the Go language. |
| [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) | 49 | 5 | 2014-09-25 | 1 year ago | Collection of benchmarks for popular Go database/SQL utilities. |
| [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) | 20 | 1 | 2017-01-24 | 2 years ago | Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results. |
| [golang-micro-benchmarks](https://github.com/amscanne/golang-micro-benchmarks) | 17 | 2 | 2014-04-05 | 8 months ago | Tiny collection of Go micro benchmarks. The intent is to compare some language features to others. |
| [kvbench](https://github.com/jimrobinson/kvbench) | 16 | 1 | 2014-04-15 | 1 month ago | Key/Value database benchmark. |

## Conferences
        

## E-Books
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [GoBooks](https://github.com/dariubs/GoBooks) | 7087 | 513 | 2015-05-05 | 1 week ago | A curated list of Go books. |
| [The-Golang-Standard-Library-by-Example](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example) | 5908 | 540 | 2013-04-14 | 2 weeks ago | Golang标准库。对于程序员而言，标准库与语言本身同样重要，它好比一个百宝箱，能为各种常见的任务提供完美的解决方案。以示例驱动的方式讲解Golang的标准库。 |
| [gosuccinctly](https://github.com/thedevsir/gosuccinctly) | 13 | 3 | 2018-09-02 | 1 year ago | in Persian. |

## Gophers
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gophers](https://github.com/ashleymcnamara/gophers) | 1931 | 85 | 2017-02-15 | 6 months ago | Gopher artworks by Ashley McNamara. |
| [gophers](https://github.com/egonelbre/gophers) | 1700 | 40 | 2015-06-03 | 7 months ago | Free gophers. |
| [free-gophers-pack](https://github.com/MariaLetta/free-gophers-pack) | 1658 | 41 | 2019-04-03 | 5 days ago | Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster. |
| [gophericons](https://github.com/shalakhin/gophericons) | 560 | 20 | 2015-08-22 | 1 year ago | 34 gopher images for Go developers community |
| [gopher-stickers](https://github.com/tenntenn/gopher-stickers) | 455 | 15 | 2014-11-10 | 3 years ago | gopher stickers |
| [gopher-vector](https://github.com/golang-samples/gopher-vector) | 345 | 13 | 2013-03-31 | 3 years ago | Vector data of gopher |
| [gopherize.me](https://github.com/matryer/gopherize.me) | 338 | 7 | 2017-01-25 | 1 year ago | Gopherize yourself. |
| [gopher-logos](https://github.com/GolangUA/gopher-logos) | 68 | 7 | 2017-07-27 | 1 year ago | adorable gopher logos. |
| [gophers](https://github.com/rogeralsing/gophers) | 50 | 2 | 2017-01-29 | 2 years ago | random gopher graphics. |
| [go-gopher](https://github.com/sillecelik/go-gopher) | 41 | 0 | 2018-03-29 | 3 weeks ago | Gopher amigurumi toy pattern. |
| [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) | 30 | 1 | 2014-09-04 | 1 year ago | Go gopher Vector Data [.ai, .svg]. |

## Meetups
        
*Add the group of your city/country here (send **PR**)*

## Twitter
        

## Websites
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) | 25122 | 1726 | 2014-07-08 | 6 days ago | List of other amazingly awesome lists. |
| [awesome-remote-job](https://github.com/lukasz-madon/awesome-remote-job) | 15132 | 832 | 2015-01-02 | 2 days ago | Curated list of awesome remote jobs. A lot of them are looking for Go hackers. |
| [golang-graphics](https://github.com/mholt/golang-graphics) | 143 | 9 | 2014-03-25 | 4 years ago | Collection of Go images, graphics, and art. |
| [gocryforhelp](https://github.com/ninedraft/gocryforhelp) | 36 | 12 | 2016-05-09 | 2 years ago | Collection of Go projects that needs help. Good place to start your open-source way in Go. |

### Tutorials
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang) | 32588 | 2440 | 2012-08-02 | 3 weeks ago | Golang ebook intro how to build a web app with golang. |
| [go-patterns](https://github.com/tmrts/go-patterns) | 11288 | 525 | 2015-12-15 | 6 days ago | Curated list of Go design patterns, recipes and idioms. |
| [learn-go-with-tests](https://github.com/quii/learn-go-with-tests) | 8497 | 232 | 2018-03-02 | 6 minutes ago | Learn Go with test-driven development. |
| [golang-cheat-sheet](https://github.com/a8m/golang-cheat-sheet) | 4213 | 168 | 2014-02-13 | 2 months ago | Go's reference card. |
| [go-lang-cheat-sheet](https://github.com/a8m/go-lang-cheat-sheet) | 4068 | 163 | 2014-02-13 | 2 months ago | Go's reference card. |
| [working-with-go](https://github.com/mkaz/working-with-go) | 1142 | 50 | 2014-05-05 | 2 weeks ago | Intro to go for experienced programmers. |
| [golang-for-nodejs-developers](https://github.com/miguelmota/golang-for-nodejs-developers) | 838 | 27 | 2019-01-03 | 4 months ago | Examples of Golang compared to Node.js for learning. |
| [ethereum-development-with-go-book](https://github.com/miguelmota/ethereum-development-with-go-book) | 488 | 30 | 2018-05-16 | 1 month ago | A little e-book on Ethereum Development with Go. |

> 这仅仅是个人练手学习go语言的一个小项目，代码low的地方欢迎指点 <plholx@126.com> ^_^ 

> 更专业的go开源项目分析请移步 [Awesome Go](https://go.libhunt.com/) 