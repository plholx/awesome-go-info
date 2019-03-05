# Awesome Go Data

go相关的开源项目列表，项目分类及GitHub上的开源项目数据完全来自于[awesome-go](https://github.com/avelino/awesome-go) 的[README.md](https://github.com/avelino/awesome-go/blob/master/README.md)文件，通过调用GitHub的API获取仓库信息，展示项目的star数、watch数等，方便选择适合的项目。

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
- [Tools](#tools)
    - [Code Analysis](#code-analysis)
    - [Editor Plugins](#editor-plugins)
    - [Go Generate Tools](#go-generate-tools)
    - [Go Tools](#go-tools)
    - [Software Packages](#software-packages)
        - [DevOps Tools](#devops-tools)
        - [Other Software](#other-software)
- [Server Applications](#server-applications)
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

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [portaudio](https://github.com/gordonklaus/portaudio) | 266 | 9 | 2015-09-16 | 6 months ago | Go bindings for the PortAudio audio I/O library |
| [music-theory](https://github.com/go-music-theory/music-theory) | 233 | 10 | 2016-03-17 | 8 months ago | Go models of Note, Scale, Chord and Key |
| [waveform](https://github.com/mdlayher/waveform) | 229 | 13 | 2014-09-14 | 2 years ago | Go package capable of generating waveform images from audio streams. MIT Licensed. |
| [portmidi](https://github.com/rakyll/portmidi) | 188 | 6 | 2013-11-10 | 1 year ago | Go bindings for libportmidi |
| [flac](https://github.com/mewkiz/flac) | 92 | 6 | 2012-11-02 | 1 week ago | Package flac provides access to FLAC (Free Lossless Audio Codec) streams. |
| [mix](https://github.com/go-mix/mix) | 90 | 3 | 2016-01-03 | 1 year ago | Sequence-based Go-native audio mixer for music apps |
| [id3v2](https://github.com/bogem/id3v2) | 90 | 5 | 2016-05-16 | 3 weeks ago | 🎵 ID3 parsing and writing library for Go |
| [go-sox](https://github.com/krig/go-sox) | 87 | 8 | 2013-10-08 | 8 months ago | libsox bindings for go |
| [mp3](https://github.com/tcolgate/mp3) | 85 | 1 | 2015-02-27 | 1 year ago | golang mp3 frame parser |
| [flac](https://github.com/eaburns/flac) | 81 | 5 | 2013-08-21 | 1 year ago | A Free Lossless Audio Codec decoder in Go |
| [go-taglib](https://github.com/wtolson/go-taglib) | 64 | 6 | 2012-11-20 | 7 months ago | Go wrapper for taglib |
| [malgo](https://github.com/gen2brain/malgo) | 55 | 4 | 2017-11-10 | 1 month ago | Mini audio library |
| [gaad](https://github.com/Comcast/gaad) | 50 | 10 | 2016-07-11 | 1 year ago | GAAD (Go Advanced Audio Decoder) |
| [go_mediainfo](https://github.com/zhulik/go_mediainfo) | 22 | 1 | 2015-12-14 | 3 years ago | Golang bindings for libmediainfo |
| [vorbis](https://github.com/mccoyst/vorbis) | 21 | 3 | 2013-07-12 | 1 year ago | A "native" ogg vorbis decoder for Go (uses inline stb_vorbis) |
| [minimp3](https://github.com/tosone/minimp3) | 19 | 2 | 2018-01-26 | 3 days ago | Decode mp3 base on https://github.com/lieff/minimp3 |
| [EasyMIDI](https://github.com/algoGuy/EasyMIDI) | 17 | 3 | 2018-02-20 | 11 months ago | EasyMidi is a simple and reliable library for working with standard midi file (SMF) |
| [gosamplerate](https://github.com/dh1tw/gosamplerate) | 8 | 1 | 2016-11-21 | 9 months ago | Go Bindings for libsamplerate |

## Authentication and OAuth
        
*Libraries for implementing authentications schemes.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [jwt-go](https://github.com/dgrijalva/jwt-go) | 5025 | 130 | 2012-04-18 | 1 week ago | Golang implementation of JSON Web Tokens (JWT) |
| [casbin](https://github.com/casbin/casbin) | 3874 | 140 | 2017-04-08 | 2 weeks ago | Authorization library that supports access control models like ACL, RBAC, ABAC. |
| [casbin](https://github.com/casbin/casbin) | 3865 | 140 | 2017-04-08 | 2 weeks ago | Authorization library that supports access control models like ACL, RBAC, ABAC. |
| [casbin](https://github.com/casbin/casbin) | 3864 | 140 | 2017-04-08 | 2 weeks ago | Authorization library that supports access control models like ACL, RBAC, ABAC. |
| [oauth2](https://github.com/golang/oauth2) | 2116 | 92 | 2014-04-14 | 6 days ago | Go OAuth2 |
| [goth](https://github.com/markbates/goth) | 2082 | 58 | 2014-10-15 | 1 week ago | Package goth provides a simple, clean, and idiomatic way to write authentication packages for Go web applications. |
| [authboss](https://github.com/volatiletech/authboss) | 1742 | 38 | 2015-01-03 | 3 weeks ago | The boss of http auth. |
| [osin](https://github.com/openshift/osin) | 1507 | 69 | 2013-09-11 | 4 months ago | Golang OAuth2 server library |
| [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) | 1119 | 72 | 2015-11-01 | 1 week ago | A standalone, specification-compliant,  OAuth2 server written in Golang. |
| [go-jose](https://github.com/square/go-jose) | 1016 | 61 | 2014-11-15 | 6 days ago | An implementation of JOSE standards (JWE, JWS, JWT) in Go |
| [gologin](https://github.com/dghubble/gologin) | 960 | 27 | 2015-06-23 | 4 hours ago | Go login handlers for authentication providers (OAuth1, OAuth2) |
| [gorbac](https://github.com/mikespook/gorbac) | 831 | 55 | 2013-12-26 | 1 month ago | goRBAC provides a lightweight role-based access control (RBAC) implementation in Golang. |
| [loginsrv](https://github.com/tarent/loginsrv) | 735 | 45 | 2016-11-11 | 3 days ago | JWT login microservice with plugable backends such as OAuth2, Google, Github, htpasswd, osiam, .. |
| [permissions2](https://github.com/xyproto/permissions2) | 300 | 12 | 2014-11-19 | 1 week ago |   :closed_lock_with_key: Middleware for keeping track of users, login states and permissions |
| [paseto](https://github.com/o1egl/paseto) | 194 | 9 | 2018-01-23 | 3 months ago | Platform-Agnostic Security Tokens implementation in GO (Golang) |
| [httpauth](https://github.com/goji/httpauth) | 167 | 5 | 2014-05-27 | 2 years ago | HTTP Authentication middlewares |
| [jwt-auth](https://github.com/adam-hanna/jwt-auth) | 144 | 9 | 2016-07-06 | 1 month ago | This package provides json web token (jwt) middleware for goLang http servers |
| [session](https://github.com/icza/session) | 82 | 6 | 2016-02-08 | 5 months ago | Go session management for web servers (including support for Google App Engine - GAE). |
| [jwt](https://github.com/robbert229/jwt) | 60 | 6 | 2016-06-06 | 3 months ago | This is an implementation of JWT in golang! |
| [branca](https://github.com/hako/branca) | 50 | 5 | 2018-01-09 | 6 months ago | :key: Secure alternative to JWT. Authenticated Encrypted API Tokens for Go. |
| [sessions](https://github.com/adam-hanna/sessions) | 41 | 3 | 2017-04-29 | 1 year ago | A dead simple, highly performant, highly customizable sessions middleware for go http servers. |
| [jwt](https://github.com/pascaldekloe/jwt) | 40 | 2 | 2018-03-21 | 1 day ago | JSON Web Token library |
| [securecookie](https://github.com/chmike/securecookie) | 26 | 4 | 2017-09-03 | 6 months ago | Fast, secure and efficient secure cookie encoder/decoder  |
| [rbac](https://github.com/zpatrick/rbac) | 20 | 3 | 2018-08-02 | 6 months ago | Minimalistic RBAC package for Go applications |
| [signedvalue](https://github.com/sashka/signedvalue) | 8 | 0 | 2018-01-06 | 1 month ago | Compatibility layer for tornado's signed values (and secure cookies consequently) |
| [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) | 6 | 2 | 2017-10-20 | 3 months ago | Driver for the SessionGate Redis module for easy session management in the Go language. |
| [cookiestxt](https://github.com/mengzhuo/cookiestxt) | 2 | 1 | 2017-10-09 | 1 year ago | cookiestxt implement parser of cookies txt format |

## Bot Building
        
*Libraries for building and working with bots.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) | 1387 | 56 | 2015-06-25 | 3 days ago | Simple and clean Telegram bot client. |
| [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) | 1385 | 56 | 2015-06-25 | 3 days ago | Simple and clean Telegram bot client. |
| [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) | 1385 | 56 | 2015-06-25 | 3 days ago | Simple and clean Telegram bot client. |
| [telebot](https://github.com/tucnak/telebot) | 827 | 34 | 2015-06-26 | 2 days ago | Telebot is a Telegram bot framework in Go. |
| [bot](https://github.com/go-chat-bot/bot) | 414 | 34 | 2015-09-23 | 2 weeks ago | IRC, Slack, Telegram and RocketChat bot written in go |
| [slacker](https://github.com/shomali11/slacker) | 256 | 13 | 2017-05-20 | 1 week ago | Slack Bot Framework |
| [tbot](https://github.com/yanzay/tbot) | 184 | 8 | 2015-09-12 | 2 months ago | Telegram Bot Server |
| [tenyks](https://github.com/kyleterry/tenyks) | 166 | 14 | 2012-08-26 | 2 years ago | The Tenyks IRC bot. |
| [golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot) | 148 | 19 | 2017-05-15 | 3 weeks ago | A golang implementation of a console-based trading bot for cryptocurrency exchanges |
| [go-sarah](https://github.com/oklahomer/go-sarah) | 111 | 3 | 2016-11-06 | 7 months ago | Simple yet customizable bot framework written in Go. |
| [hanu](https://github.com/sbstjn/hanu) | 90 | 6 | 2016-09-16 | 6 months ago | Golang Framework for writing Slack bots |
| [go-tgbot](https://github.com/olebedev/go-tgbot) | 80 | 7 | 2016-12-11 | 8 months ago | Golang  telegram bot API wrapper, session-based router and middleware |
| [margelet](https://github.com/zhulik/margelet) | 51 | 5 | 2015-11-21 | 2 years ago | Telegram Bot Framework for Go |
| [govkbot](https://github.com/nikepan/govkbot) | 19 | 2 | 2016-07-12 | 15 hours ago | VK bot package for Go |
| [micha](https://github.com/onrik/micha) | 9 | 3 | 2016-04-14 | 1 year ago | Client lib for Telegram bot api |

## Command Line
        
*Libraries for building Console Applications and Console User Interfaces.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [cobra](https://github.com/spf13/cobra) | 10791 | 255 | 2013-09-04 | 1 week ago | A Commander for modern Go CLI interactions |
| [cli](https://github.com/urfave/cli) | 10090 | 271 | 2013-07-14 | 3 weeks ago | A simple, fast, and fun package for building command line apps in Go |
| [termui](https://github.com/gizak/termui) | 8362 | 275 | 2015-02-03 | 3 days ago | Golang terminal dashboard |
| [gocui](https://github.com/jroimartin/gocui) | 3971 | 102 | 2014-01-04 | 1 day ago | Minimalist Go package aimed at creating Console User Interfaces. |
| [termbox-go](https://github.com/nsf/termbox-go) | 3293 | 103 | 2012-01-13 | 2 weeks ago | Pure Go termbox implementation |
| [color](https://github.com/fatih/color) | 2866 | 56 | 2014-02-17 | 4 months ago | Color package for Go (golang) |
| [kingpin](https://github.com/alecthomas/kingpin) | 2301 | 54 | 2014-05-15 | 1 month ago | A Go (golang) command line and flag parser |
| [go-prompt](https://github.com/c-bata/go-prompt) | 2081 | 35 | 2017-08-15 | 1 week ago | Building powerful interactive prompts in Go, inspired by python-prompt-toolkit. |
| [readline](https://github.com/chzyer/readline) | 1307 | 39 | 2015-09-20 | 3 weeks ago | Readline is a pure go(golang) implementation for GNU-Readline kind library |
| [go-flags](https://github.com/jessevdk/go-flags) | 1294 | 28 | 2012-08-31 | 1 week ago | go command line option parser |
| [uiprogress](https://github.com/gosuri/uiprogress) | 1229 | 27 | 2015-11-17 | 4 months ago | A go library to render progress bars in terminal applications |
| [docopt.go](https://github.com/docopt/docopt.go) | 1038 | 36 | 2013-08-26 | 5 months ago | A command-line arguments parser that will make you smile. |
| [asciigraph](https://github.com/guptarohit/asciigraph) | 1022 | 23 | 2018-06-17 | 1 month ago | Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. |
| [cli](https://github.com/mitchellh/cli) | 928 | 25 | 2013-11-03 | 3 months ago | A Go library for implementing command-line interfaces. |
| [gcli](https://github.com/tcnksm/gcli) | 856 | 25 | 2014-06-19 | 1 year ago | The easy way to build Golang command-line application. |
| [uilive](https://github.com/gosuri/uilive) | 724 | 11 | 2015-11-16 | 3 months ago | uilive is a go library for updating terminal output in realtime |
| [pflag](https://github.com/spf13/pflag) | 627 | 24 | 2013-08-30 | 3 weeks ago | Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. |
| [mow.cli](https://github.com/jawher/mow.cli) | 590 | 19 | 2014-12-19 | 1 week ago | A versatile library for building CLI applications in Go |
| [go-arg](https://github.com/alexflint/go-arg) | 571 | 15 | 2015-11-01 | 2 months ago | Struct-based argument parsing in Go |
| [complete](https://github.com/posener/complete) | 542 | 12 | 2017-05-06 | 1 week ago | bash completion written in go + bash completion for go command |
| [liner](https://github.com/peterh/liner) | 522 | 22 | 2012-08-16 | 10 hours ago | Pure Go line editor with history, inspired by linenoise |
| [progressbar](https://github.com/schollz/progressbar) | 494 | 13 | 2017-10-27 | 1 month ago | A really basic thread-safe progress bar for Golang applications |
| [uitable](https://github.com/gosuri/uitable) | 464 | 15 | 2015-11-14 | 1 year ago | A go library to improve readability in terminal apps using tabular data |
| [aurora](https://github.com/logrusorgru/aurora) | 428 | 5 | 2016-11-07 | 5 months ago | Golang ultimate ANSI-colors that supports Printf/Sprintf methods |
| [cli](https://github.com/mkideal/cli) | 427 | 20 | 2016-02-27 | 5 days ago | CLI - A package for building command line app with go |
| [mpb](https://github.com/vbauerster/mpb) | 419 | 8 | 2016-12-14 | 1 week ago | multi progress bar for Go cli applications |
| [flaggy](https://github.com/integrii/flaggy) | 410 | 11 | 2018-03-05 | 1 week ago | Idiomatic Go input parsing with subcommands, positional values, and flags at any position. No required project or package layout and no external dependencies. |
| [go-colorable](https://github.com/mattn/go-colorable) | 334 | 16 | 2014-07-30 | 1 week ago |  |
| [go-isatty](https://github.com/mattn/go-isatty) | 299 | 7 | 2014-04-01 | 1 week ago |  |
| [chalk](https://github.com/ttacon/chalk) | 286 | 11 | 2014-07-19 | 2 years ago | Intuitive package for prettifying terminal/console output. http://godoc.org/github.com/ttacon/chalk |
| [termtables](https://github.com/apcera/termtables) | 203 | 68 | 2012-12-06 | 1 year ago | Go ASCII Table Generator, ported from the Ruby terminal-tables library |
| [go-colortext](https://github.com/daviddengcn/go-colortext) | 188 | 9 | 2013-01-23 | 10 months ago | Change the color of console text. |
| [color](https://github.com/gookit/color) | 130 | 4 | 2018-07-01 | 1 week ago | Terminal color rendering tool library, support 8/16 colors, 256 colors, RGB color rendering output, compatible with Windows. CLI 控制台颜色渲染工具库, 拥有简洁的使用API，支持16色，256色，RGB色彩渲染输出，兼容 Windows 环境 |
| [simpletable](https://github.com/alexeyco/simpletable) | 121 | 2 | 2017-03-29 | 1 week ago | Simple tables in terminal with Go |
| [clif](https://github.com/ukautz/clif) | 95 | 2 | 2015-05-31 | 2 weeks ago | Another CLI framework for Go. It works on my machine. |
| [flag](https://github.com/cosiner/flag) | 90 | 5 | 2016-10-06 | 1 week ago | Flag is a simple but powerful command line option parsing library for Go support infinite level subcommand |
| [termdash](https://github.com/mum4k/termdash) | 83 | 6 | 2018-03-24 | 1 day ago | Terminal based dashboard. |
| [argparse](https://github.com/akamensky/argparse) | 74 | 5 | 2017-11-24 | 1 month ago | Argparse for golang. Just because `flag` sucks |
| [sflags](https://github.com/octago/sflags) | 71 | 5 | 2016-12-04 | 1 month ago | Generate flags by parsing structures |
| [commandeer](https://github.com/jaffee/commandeer) | 70 | 4 | 2017-10-12 | 3 days ago | Automatically sets up command line flags based on struct fields and tags. |
| [wmenu](https://github.com/dixonwille/wmenu) | 67 | 2 | 2016-04-20 | 1 week ago | An easy to use menu structure for cli applications that prompts users to make choices. |
| [cfmt](https://github.com/mingrammer/cfmt) | 53 | 3 | 2018-03-16 | 2 months ago | :art: Contextual fmt inspired by bootstrap color classes |
| [cli](https://github.com/teris-io/cli) | 43 | 1 | 2017-05-25 | 8 months ago | Simple and complete API for building command line applications in Go |
| [env](https://github.com/codingconcepts/env) | 37 | 1 | 2017-06-15 | 4 months ago | Tag-based environment configuration for structs |
| [wlog](https://github.com/dixonwille/wlog) | 29 | 1 | 2016-04-14 | 1 year ago | A simple logging interface that supports cross-platform color and concurrency. |
| [gocmd](https://github.com/devfacet/gocmd) | 28 | 3 | 2018-01-08 | 6 months ago | A Go library for building command line applications |
| [flagvar](https://github.com/sgreben/flagvar) | 25 | 1 | 2018-05-19 | 4 months ago | A collection of CLI argument types for the Go `flag` package. |
| [tabular](https://github.com/InVisionApp/tabular) | 23 | 3 | 2018-04-24 | 9 months ago | Tabular simplifies printing ASCII tables from command line utilities |
| [strumt](https://github.com/antham/strumt) | 22 | 0 | 2017-06-20 | 3 days ago | Strumt is a library to create prompt chain |
| [colourize](https://github.com/TreyBastian/colourize) | 15 | 3 | 2015-05-11 | 2 years ago | An ANSI colour terminal package for Go |
| [argv](https://github.com/cosiner/argv) | 11 | 1 | 2017-01-22 | 1 month ago |  |
| [go-commander](https://github.com/yitsushi/go-commander) | 8 | 1 | 2016-10-10 | 1 week ago | Go library to simplify CLI workflow |
| [go-ataman](https://github.com/workanator/go-ataman) | 8 | 2 | 2017-05-18 | 1 year ago | Another Text Attribute Manupulator |
| [ctc](https://github.com/wzshiming/ctc) | 6 | 1 | 2018-04-28 | 4 months ago | Console Text Colors - The non-invasive cross-platform terminal color library does not need to modify the Print method |
| [go-getoptions](https://github.com/DavidGamba/go-getoptions) | 5 | 1 | 2015-12-18 | 2 days ago | Go (golang) command line option parser inspired on the flexibility of Perl’s GetOpt::Long. |
| [sand](https://github.com/Zaba505/sand) | 2 | 1 | 2018-11-19 | 3 months ago | Package for creating interpreters |

## Configuration
        
*Libraries for configuration parsing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [viper](https://github.com/spf13/viper) | 7673 | 165 | 2014-04-02 | 4 days ago | Go configuration with fangs |
| [envconfig](https://github.com/kelseyhightower/envconfig) | 2127 | 37 | 2013-11-07 | 2 weeks ago | Golang library for managing configuration data from environment variables |
| [godotenv](https://github.com/joho/godotenv) | 1750 | 30 | 2013-07-30 | 1 week ago | A Go port of Ruby's dotenv library (Loads environment variables from `.env`.) |
| [ini](https://github.com/go-ini/ini) | 1332 | 62 | 2014-12-18 | 2 weeks ago | Package ini provides INI file read and write functionality in Go. |
| [env](https://github.com/caarlos0/env) | 702 | 16 | 2015-07-28 | 2 days ago | A KISS way to deal with environment variables in Go. |
| [store](https://github.com/tucnak/store) | 239 | 5 | 2015-10-04 | 1 year ago | A dead simple configuration manager for Go applications |
| [confita](https://github.com/heetch/confita) | 208 | 25 | 2017-12-21 | 2 weeks ago | Load configuration in cascade from multiple backends into a struct |
| [config](https://github.com/joshbetz/config) | 193 | 3 | 2017-04-03 | 1 year ago | A configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP |
| [config](https://github.com/olebedev/config) | 190 | 8 | 2014-04-21 | 5 months ago | JSON or YAML configuration wrapper with convenient access methods. |
| [hjson-go](https://github.com/hjson/hjson-go) | 163 | 7 | 2016-08-06 | 3 weeks ago | Hjson for Go |
| [envconfig](https://github.com/vrischmann/envconfig) | 136 | 4 | 2015-04-22 | 2 months ago | Small library to read your configuration from environment variables |
| [gcfg](https://github.com/go-gcfg/gcfg) | 102 | 4 | 2015-08-17 | 9 months ago | read INI-style configuration files into Go structs; supports user-defined types and subsections |
| [goconfig](https://github.com/crgimenes/goconfig) | 93 | 11 | 2016-12-18 | 3 weeks ago | Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. |
| [goconfig](https://github.com/crgimenes/goconfig) | 92 | 11 | 2016-12-18 | 3 weeks ago | Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. |
| [envh](https://github.com/antham/envh) | 92 | 3 | 2017-01-12 | 3 days ago | Go helpers to manage environment variables |
| [goconfig](https://github.com/crgimenes/goconfig) | 92 | 11 | 2016-12-18 | 3 weeks ago | Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. |
| [envcfg](https://github.com/tomazk/envcfg) | 89 | 1 | 2014-11-29 | 1 year ago | Un-marshaling environment variables to Go structs |
| [gofigure](https://github.com/ian-kent/gofigure) | 56 | 6 | 2014-11-25 | 1 year ago | Go configuration made easy! |
| [config](https://github.com/gookit/config) | 45 | 3 | 2018-07-07 | 1 month ago | Go config manage(load,get,set). support JSON, YAML, TOML, INI, HCL, ENV and Flags. Multi file load, data override merge, parse ENV var. Go应用配置加载管理，支持多种格式，多文件加载，远程文件加载，支持数据合并，解析环境变量名 |
| [configure](https://github.com/paked/configure) | 44 | 3 | 2015-06-14 | 2 weeks ago | Configure is a Go package that gives you easy configuration of your project through redundancy |
| [ingo](https://github.com/schachmat/ingo) | 22 | 1 | 2016-02-08 | 1 year ago | persistent storage for flags in go |
| [go-up](https://github.com/ufoscout/go-up) | 21 | 1 | 2018-02-18 | 1 day ago | go-up! A simple configuration library with recursive placeholders resolution and no magic. |
| [xdg](https://github.com/OpenPeeDeeP/xdg) | 19 | 3 | 2017-07-20 | 1 week ago | A cross platform package that follows the XDG Standard |
| [mini](https://github.com/sasbury/mini) | 19 | 2 | 2015-04-30 | 2 months ago | A golang package for parsing ini-style configuration files |
| [envconf](https://github.com/ian-kent/envconf) | 7 | 1 | 2014-10-26 | 4 years ago | Configure Go applications from the environment |
| [conflate](https://github.com/the4thamigo-uk/conflate) | 5 | 0 | 2018-02-02 | 3 weeks ago | Library providing routines to merge and validate JSON, YAML and/or TOML files |
| [sprbox](https://github.com/oblq/sprbox) | 3 | 2 | 2018-07-17 | 4 months ago | Build-environment aware toolbox factory & agnostic, layered, config parser (supporting YAML, TOML, JSON and Environment vars). |

## Continuous Integration
        
*Tools for help with continuous integration.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [drone](https://github.com/drone/drone) | 17364 | 585 | 2014-02-07 | 6 hours ago | Drone is a Container-Native, Continuous Delivery Platform |
| [goveralls](https://github.com/mattn/goveralls) | 540 | 16 | 2013-04-17 | 1 month ago |  |
| [overalls](https://github.com/go-playground/overalls) | 94 | 4 | 2015-07-30 | 6 months ago | :jeans:Multi-Package go project coverprofile for tools like goveralls |
| [duci](https://github.com/duck8823/duci) | 32 | 3 | 2018-04-01 | 1 day ago | The simple ci server  |
| [gomason](https://github.com/nikogura/gomason) | 22 | 0 | 2017-11-18 | 3 weeks ago | A tool for testing, building, signing, and publishing go binaries from a clean workspace. |
| [roveralls](https://github.com/lawrencewoodman/roveralls) | 11 | 1 | 2016-06-26 | 1 year ago | Recursive coverage testing tool. |
| [roveralls](https://github.com/lawrencewoodman/roveralls) | 11 | 1 | 2016-06-26 | 1 year ago | Recursive coverage testing tool. |
| [roveralls](https://github.com/lawrencewoodman/roveralls) | 11 | 1 | 2016-06-26 | 1 year ago | Recursive coverage testing tool. |

## CSS Preprocessors
        
*Libraries for preprocessing CSS files.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gcss](https://github.com/yosssi/gcss) | 409 | 16 | 2014-09-04 | 4 years ago | Pure Go CSS Preprocessor |
| [go-libsass](https://github.com/wellington/go-libsass) | 116 | 4 | 2015-04-19 | 3 weeks ago | Go wrapper for libsass, the only Sass 3.5 compiler for Go |

## Data Structures
        
*Generic datastructures and algorithms in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gods](https://github.com/emirpasic/gods) | 5354 | 255 | 2015-03-04 | 2 days ago | GoDS (Go Data Structures). Containers (Sets, Lists, Stacks, Maps, Trees), Sets (HashSet, TreeSet, LinkedHashSet), Lists (ArrayList, SinglyLinkedList, DoublyLinkedList), Stacks (LinkedListStack, ArrayStack), Maps (HashMap, TreeMap, HashBidiMap, TreeBidiMap, LinkedHashMap), Trees (RedBlackTree, AVLTree, BTree, BinaryHeap), Comparators, Iterators, Enumerables, Sort, JSON |
| [go-datastructures](https://github.com/Workiva/go-datastructures) | 4783 | 292 | 2014-10-29 | 3 months ago |  |
| [BoomFilters](https://github.com/tylertreat/BoomFilters) | 1081 | 36 | 2015-02-06 | 4 months ago | Probabilistic data structures for processing continuous, unbounded streams. |
| [golang-set](https://github.com/deckarep/golang-set) | 977 | 27 | 2013-07-04 | 5 months ago | A simple set type for the Go language. Also used by Docker, 1Password, Ethereum. |
| [gota](https://github.com/go-gota/gota) | 738 | 57 | 2016-02-07 | 6 days ago | Implementation of dataframes, series, and data wrangling methods for Go. |
| [gota](https://github.com/go-gota/gota) | 737 | 57 | 2016-02-07 | 6 days ago | Implementation of dataframes, series, and data wrangling methods for Go. |
| [gota](https://github.com/go-gota/gota) | 737 | 57 | 2016-02-07 | 6 days ago | Implementation of dataframes, series, and data wrangling methods for Go. |
| [hyperloglog](https://github.com/axiomhq/hyperloglog) | 634 | 18 | 2017-06-18 | 2 days ago | HyperLogLog with lots of sugar (Sparse, LogLog-Beta bias correction and TailCut space reduction) |
| [bloom](https://github.com/willf/bloom) | 589 | 30 | 2011-05-21 | 4 days ago | Go package implementing Bloom filters |
| [roaring](https://github.com/RoaringBitmap/roaring) | 567 | 34 | 2014-07-11 | 1 month ago | Roaring bitmaps in Go (golang) |
| [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) | 445 | 17 | 2015-06-29 | 2 days ago | Cuckoo Filter: Practically Better Than Bloom |
| [bitset](https://github.com/willf/bitset) | 441 | 27 | 2011-05-11 | 4 days ago | Go package implementing bitsets |
| [trie](https://github.com/derekparker/trie) | 363 | 20 | 2014-03-07 | 9 months ago | Data structure and relevant algorithms for extremely fast prefix/fuzzy string searching. |
| [go-geoindex](https://github.com/hailocab/go-geoindex) | 297 | 67 | 2015-01-22 | 1 year ago | Go native library for fast point tracking and K-Nearest queries |
| [mafsa](https://github.com/smartystreets/mafsa) | 272 | 18 | 2014-11-08 | 2 years ago | Package mafsa implements Minimal Acyclic Finite State Automata in Go, essentially a high-speed, memory-efficient, Unicode-friendly set of strings. |
| [goskiplist](https://github.com/ryszard/goskiplist) | 178 | 12 | 2012-05-09 | 2 years ago | A skip list implementation in Go |
| [hilbert](https://github.com/google/hilbert) | 171 | 19 | 2015-08-06 | 3 months ago | Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves. |
| [algorithms](https://github.com/shady831213/algorithms) | 155 | 6 | 2018-01-31 | 1 month ago | CLRS study. Codes are written with golang. |
| [bloom](https://github.com/zentures/bloom) | 125 | 8 | 2013-09-03 | 10 months ago | Bloom filters implemented in Go. |
| [bloom](https://github.com/zentures/bloom) | 125 | 8 | 2013-09-03 | 10 months ago | Bloom filters implemented in Go. |
| [bloom](https://github.com/zentures/bloom) | 125 | 8 | 2013-09-03 | 10 months ago | Bloom filters implemented in Go. |
| [merkletree](https://github.com/cbergoon/merkletree) | 119 | 4 | 2017-04-12 | 1 day ago | A Merkle Tree implementation written in Go. |
| [binpacker](https://github.com/zhuangsirui/binpacker) | 110 | 10 | 2016-02-02 | 8 months ago | A binary stream packer and unpacker |
| [go-rquad](https://github.com/arl/go-rquad) | 94 | 3 | 2016-09-13 | 8 months ago | Region quadtrees with efficient point location and neighbour finding. |
| [go-rquad](https://github.com/arl/go-rquad) | 94 | 3 | 2016-09-13 | 8 months ago | Region quadtrees with efficient point location and neighbour finding. |
| [go-rquad](https://github.com/arl/go-rquad) | 94 | 3 | 2016-09-13 | 8 months ago | Region quadtrees with efficient point location and neighbour finding. |
| [encoding](https://github.com/zentures/encoding) | 92 | 7 | 2013-09-21 | 1 year ago | Integer Compression Libraries for Go. |
| [encoding](https://github.com/zentures/encoding) | 92 | 7 | 2013-09-21 | 1 year ago | Integer Compression Libraries for Go. |
| [encoding](https://github.com/zentures/encoding) | 92 | 7 | 2013-09-21 | 1 year ago | Integer Compression Libraries for Go. |
| [ring](https://github.com/TheTannerRyan/ring) | 83 | 1 | 2019-01-27 | 2 weeks ago | Package ring provides a high performance and thread safe Go implementation of a bloom filter. |
| [skiplist](https://github.com/MauriceGit/skiplist) | 80 | 5 | 2018-06-24 | 2 months ago | A Go library for an efficient implementation of a skip list: https://godoc.org/github.com/MauriceGit/skiplist |
| [ttlcache](https://github.com/ReneKroon/ttlcache) | 78 | 7 | 2014-12-13 | 1 week ago | In-memory LRU string-interface{} map with expiration for golang. |
| [ttlcache](https://github.com/ReneKroon/ttlcache) | 78 | 7 | 2014-12-13 | 1 week ago | In-memory LRU string-interface{} map with expiration for golang. |
| [ttlcache](https://github.com/ReneKroon/ttlcache) | 78 | 7 | 2014-12-13 | 1 week ago | In-memory LRU string-interface{} map with expiration for golang. |
| [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) | 63 | 5 | 2016-04-01 | 2 months ago | Adaptive Radix Trees implemented in Go |
| [skiplist](https://github.com/gansidui/skiplist) | 58 | 7 | 2014-11-19 | 4 years ago | skiplist for golang |
| [conjungo](https://github.com/InVisionApp/conjungo) | 58 | 14 | 2016-12-30 | 2 weeks ago | A small flexible merge library in go |
| [levenshtein](https://github.com/agnivade/levenshtein) | 42 | 3 | 2014-07-30 | 5 days ago | Go implementation to calculate Levenshtein Distance. |
| [count-min-log](https://github.com/seiflotfy/count-min-log) | 40 | 4 | 2015-08-17 | 2 years ago | Go implementation of Count-Min-Log |
| [deque](https://github.com/gammazero/deque) | 34 | 4 | 2018-04-24 | 1 month ago | Fast ring-buffer deque (double-ended queue) |
| [bit](https://github.com/yourbasic/bit) | 29 | 3 | 2017-05-04 | 11 months ago | Bitset data structure for positive numbers |
| [levenshtein](https://github.com/agext/levenshtein) | 25 | 1 | 2016-04-08 | 1 week ago | Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. |
| [go-mcache](https://github.com/OrlovEvgeny/go-mcache) | 23 | 2 | 2018-04-15 | 3 months ago | Fast in-memory key:value store/cache  library for Golang |
| [concurrent-writer](https://github.com/free/concurrent-writer) | 15 | 4 | 2017-09-18 | 1 year ago | Highly concurrent drop-in replacement for bufio.Writer |
| [bloom](https://github.com/yourbasic/bloom) | 13 | 1 | 2017-05-07 | 1 year ago | Probabilistic set data structure |
| [goset](https://github.com/zoumo/goset) | 12 | 1 | 2017-08-25 | 1 year ago | Set is a useful collection but there is no built-in implementation in Go lang. |
| [pipeline](https://github.com/hyfather/pipeline) | 9 | 1 | 2018-04-25 | 6 months ago | Pipelines using goroutines |
| [go-ef](https://github.com/amallia/go-ef) | 8 | 1 | 2017-09-22 | 1 year ago | A Go implementation of the Elias-Fano encoding |
| [set](https://github.com/StudioSol/set) | 5 | 17 | 2018-07-21 | 4 months ago | A simple Set data structure implementation in Go (Golang) using LinkedHashMap. |
| [mspm](https://github.com/BlackRabbitt/mspm) | 4 | 3 | 2018-05-18 | 9 months ago | Multi-String Pattern Matching Algorithm Using TrieHashNode |
| [null](https://github.com/emvi/null) | 2 | 2 | 2018-07-05 | 1 week ago | Nullable Go types that can be marshalled/unmarshalled to/from JSON. |
| [hide](https://github.com/emvi/hide) | 2 | 2 | 2019-01-16 | 1 week ago | ID type with marshalling to/from hash to prevent sending IDs to clients. |
| [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) | 2 | 1 | 2019-01-11 | 1 month ago | Go concurrent safe queue |
| [deque](https://github.com/edwingeng/deque) | 1 | 1 | 2019-02-01 | 2 weeks ago | A highly optimized double-ended queue |
| [timedmap](https://github.com/zekroTJA/timedmap) | 0 | 1 | 2019-01-30 | 1 week ago | A map which has expiring key-value pairs |

## Database
        
*SQL query builder, libraries for building and using SQL.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [prometheus](https://github.com/prometheus/prometheus) | 22424 | 1007 | 2012-11-24 | 1 hour ago | The Prometheus monitoring system and time series database. |
| [tidb](https://github.com/pingcap/tidb) | 17571 | 1197 | 2015-09-06 | 1 hour ago | TiDB is a distributed HTAP database compatible with the MySQL protocol  |
| [influxdb](https://github.com/influxdata/influxdb) | 15707 | 749 | 2013-09-26 | 8 hours ago | Scalable datastore for metrics, events, and real-time analytics. |
| [influxdb](https://github.com/influxdata/influxdb) | 15705 | 748 | 2013-09-26 | 11 hours ago | Scalable datastore for metrics, events, and real-time analytics. |
| [influxdb](https://github.com/influxdata/influxdb) | 15705 | 749 | 2013-09-26 | 1 day ago | Scalable datastore for metrics, events, and real-time analytics. |
| [cockroach](https://github.com/cockroachdb/cockroach) | 15535 | 694 | 2014-02-06 | 2 hours ago | CockroachDB - the open source, cloud-native SQL database. |
| [bolt](https://github.com/boltdb/bolt) | 9498 | 338 | 2013-12-21 | 1 year ago | An embedded key/value database for Go. |
| [dgraph](https://github.com/dgraph-io/dgraph) | 8739 | 330 | 2015-08-25 | 8 hours ago | Fast, Distributed Graph DB |
| [vitess](https://github.com/vitessio/vitess) | 7613 | 464 | 2013-06-28 | 11 hours ago | vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. |
| [vitess](https://github.com/vitessio/vitess) | 7611 | 464 | 2013-06-28 | 11 hours ago | vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. |
| [vitess](https://github.com/vitessio/vitess) | 7608 | 464 | 2013-06-28 | 1 day ago | vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. |
| [groupcache](https://github.com/golang/groupcache) | 7103 | 443 | 2013-07-23 | 2 weeks ago | groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. |
| [pgweb](https://github.com/sosedoff/pgweb) | 5760 | 152 | 2014-10-09 | 1 week ago | Cross-platform client for PostgreSQL databases |
| [badger](https://github.com/dgraph-io/badger) | 5391 | 220 | 2017-01-26 | 1 day ago | Fast key-value DB in Go. |
| [rqlite](https://github.com/rqlite/rqlite) | 4260 | 182 | 2014-08-23 | 6 days ago | The lightweight, distributed relational database built on SQLite. |
| [kingshard](https://github.com/flike/kingshard) | 4193 | 388 | 2015-07-04 | 2 weeks ago | A high-performance MySQL proxy |
| [ledisdb](https://github.com/siddontang/ledisdb) | 2890 | 184 | 2014-04-30 | 4 weeks ago | a high performance NoSQL powered by Go |
| [goleveldb](https://github.com/syndtr/goleveldb) | 2800 | 151 | 2013-01-23 | 5 hours ago | LevelDB key/value database in Go. |
| [orchestrator](https://github.com/github/orchestrator) | 2602 | 181 | 2016-11-30 | 5 hours ago | MySQL replication topology management and HA |
| [go-cache](https://github.com/patrickmn/go-cache) | 2424 | 84 | 2012-01-02 | 6 days ago | In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. |
| [go-cache](https://github.com/patrickmn/go-cache) | 2422 | 84 | 2012-01-02 | 6 days ago | In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. |
| [go-cache](https://github.com/patrickmn/go-cache) | 2422 | 84 | 2012-01-02 | 6 days ago | In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. |
| [tiedot](https://github.com/HouzuoGuo/tiedot) | 2306 | 161 | 2013-05-26 | 1 month ago | Your NoSQL database powered by Golang |
| [buntdb](https://github.com/tidwall/buntdb) | 2269 | 94 | 2016-07-20 | 1 month ago | BuntDB is an embeddable, in-memory key/value database for Go with custom indexing and geospatial support |
| [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) | 2037 | 149 | 2015-01-15 | 2 days ago | Sync MySQL data into elasticsearch  |
| [prest](https://github.com/prest/prest) | 1963 | 82 | 2016-11-22 | 6 months ago | Serve a RESTful API from any PostgreSQL database. |
| [prest](https://github.com/prest/prest) | 1962 | 82 | 2016-11-22 | 6 months ago | Serve a RESTful API from any PostgreSQL database. |
| [prest](https://github.com/prest/prest) | 1962 | 82 | 2016-11-22 | 6 months ago | Serve a RESTful API from any PostgreSQL database. |
| [xo](https://github.com/xo/xo) | 1947 | 69 | 2016-02-05 | 2 days ago | Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. |
| [xo](https://github.com/xo/xo) | 1946 | 70 | 2016-02-05 | 2 days ago | Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. |
| [xo](https://github.com/xo/xo) | 1946 | 70 | 2016-02-05 | 2 days ago | Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. |
| [squirrel](https://github.com/Masterminds/squirrel) | 1940 | 33 | 2014-01-18 | 6 days ago | Fluent SQL generation for golang |
| [bigcache](https://github.com/allegro/bigcache) | 1881 | 64 | 2016-03-23 | 2 weeks ago | Efficient cache for gigabytes of data written in Go. |
| [migrate](https://github.com/golang-migrate/migrate) | 1715 | 37 | 2018-01-19 | 3 days ago | Database migrations. CLI and Golang library. |
| [go-mysql](https://github.com/siddontang/go-mysql) | 1568 | 132 | 2014-02-21 | 2 days ago | a powerful mysql toolset with Go |
| [sql-migrate](https://github.com/rubenv/sql-migrate) | 1246 | 29 | 2014-09-09 | 2 weeks ago | SQL schema migration tool for Go. |
| [cache2go](https://github.com/muesli/cache2go) | 817 | 58 | 2013-11-11 | 23 hours ago | Concurrency-safe Go caching library with expiration capabilities and access counters |
| [diskv](https://github.com/peterbourgon/diskv) | 695 | 29 | 2012-03-22 | 6 months ago | A disk-backed key-value store. |
| [moss](https://github.com/couchbase/moss) | 685 | 79 | 2016-02-07 | 1 week ago | moss - a simple, fast, ordered, persistable, key-val storage library for golang |
| [gcache](https://github.com/bluele/gcache) | 670 | 29 | 2015-01-25 | 4 days ago | Cache library for golang. It supports expirable Cache, LFU, LRU and ARC. |
| [gendry](https://github.com/didi/gendry) | 576 | 46 | 2017-12-01 | 2 months ago | a golang library for sql builder |
| [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) | 564 | 46 | 2018-04-11 | 1 hour ago | a GDPR-compliant SQL database running on Open Internet |
| [eliasdb](https://github.com/krotik/eliasdb) | 505 | 23 | 2016-08-13 | 2 weeks ago | EliasDB is a graph-based database. |
| [goqu](https://github.com/doug-martin/goqu) | 492 | 25 | 2015-02-21 | 2 weeks ago | SQL builder and query library for golang |
| [dotsql](https://github.com/gchaincl/dotsql) | 407 | 21 | 2014-11-20 | 4 months ago | A Golang library for using SQL. |
| [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) | 394 | 26 | 2015-12-11 | 6 months ago | A Go (golang) package that enhances the standard database/sql package by providing powerful data retrieval methods as well as DB-agnostic query building capabilities. |
| [fastcache](https://github.com/VictoriaMetrics/fastcache) | 350 | 13 | 2018-11-23 | 1 week ago | Fast thread-safe inmemory cache for big number of entries in Go. Minimzes GC overhead |
| [levigo](https://github.com/jmhodges/levigo) | 349 | 24 | 2012-01-17 | 5 days ago | levigo is a Go wrapper for LevelDB |
| [gormigrate](https://github.com/go-gormigrate/gormigrate) | 257 | 3 | 2016-08-31 | 4 weeks ago | Minimalistic database migration helper for Gorm ORM |
| [chproxy](https://github.com/Vertamedia/chproxy) | 241 | 20 | 2017-09-18 | 1 month ago | ClickHouse http proxy and load balancer |
| [pudge](https://github.com/recoilme/pudge) | 181 | 6 | 2018-11-20 | 6 days ago | Fast and simple key/value store written using Go's standard library |
| [piladb](https://github.com/fern4lvarez/piladb) | 167 | 8 | 2015-09-09 | 11 months ago | Lightweight RESTful database engine based on stack data structures |
| [scaneo](https://github.com/variadico/scaneo) | 141 | 8 | 2015-04-30 | 11 months ago | Generate Go code to convert database rows into arbitrary structs. |
| [sqrl](https://github.com/elgris/sqrl) | 141 | 5 | 2014-06-25 | 1 month ago | Fluent SQL generation for golang |
| [myreplication](https://github.com/2tvenom/myreplication) | 129 | 17 | 2015-02-05 | 5 months ago | Golang MySql binary log replication listener |
| [vasto](https://github.com/chrislusf/vasto) | 120 | 16 | 2018-01-16 | 4 months ago | A distributed key-value store. On Disk. Able to grow or shrink without service interruption. |
| [goose](https://github.com/steinbacher/goose) | 103 | 4 | 2016-03-05 | 2 years ago | Go database migration tool |
| [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) | 100 | 6 | 2017-04-29 | 3 hours ago | Collects many small insterts to ClickHouse and send in big inserts |
| [darwin](https://github.com/GuiaBolso/darwin) | 77 | 2 | 2016-04-05 | 3 months ago | Database schema evolution library for Go |
| [slowpoke](https://github.com/recoilme/slowpoke) | 74 | 6 | 2018-02-19 | 2 months ago | Low-level key/value store in pure Go.  |
| [igor](https://github.com/galeone/igor) | 73 | 6 | 2016-03-10 | 8 months ago | igor is an abstraction layer for PostgreSQL with a gorm like syntax. |
| [godbal](https://github.com/xujiajun/godbal) | 48 | 3 | 2018-02-28 | 1 month ago | Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily  (now only support mysql) |
| [octillery](https://github.com/knocknote/octillery) | 43 | 14 | 2018-11-26 | 1 month ago | Go package for sharding databases ( Supports every ORM or raw SQL ) |
| [couchcache](https://github.com/codingsince1985/couchcache) | 39 | 3 | 2015-04-05 | 1 year ago | A RESTful caching micro-service in Go backed by Couchbase |
| [golang-scribble](https://github.com/nanobox-io/golang-scribble) | 32 | 3 | 2018-06-22 | 4 months ago | A tiny Golang JSON database |
| [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) | 25 | 5 | 2017-12-18 | 1 year ago | golang bigcache with clustering as a library. |
| [dbbench](https://github.com/sj14/dbbench) | 25 | 2 | 2018-11-24 | 1 month ago | dbbench is a simple database benchmarking tool which supports several databases and own scripts |
| [prep](https://github.com/hexdigest/prep) | 23 | 3 | 2017-12-12 | 1 year ago | Prep finds all SQL statements in a Go package and instruments db connection with prepared statements |
| [pravasan](https://github.com/pravasan/pravasan) | 23 | 6 | 2015-01-04 | 2 months ago | Simple Migration Tool - written in Go |
| [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) | 19 | 1 | 2018-08-11 | 2 months ago | A Go package to help write migrations with go-pg/pg. |
| [go-fixtures](https://github.com/RichardKnop/go-fixtures) | 17 | 1 | 2015-12-24 | 4 months ago | Django style fixtures for Golang's excellent built-in database/sql library. |
| [tempdb](https://github.com/rafaeljesus/tempdb) | 13 | 3 | 2017-03-18 | 1 year ago | Key-value store for temporary items :memo: |
| [rwdb](https://github.com/andizzle/rwdb) | 10 | 2 | 2017-10-04 | 1 year ago | Database wrapper that manage read write connections |
| [gorocksdb](https://github.com/kapitan-k/gorocksdb) | 6 | 1 | 2017-12-28 | 1 year ago | gorocksdb is a Go wrapper for RocksDB with additional features |
| [ormlite](https://github.com/pupizoid/ormlite) | 0 | 1 | 2018-06-28 | 2 weeks ago | Lightweight package containing some ORM-like features and helpers for sqlite databases. |

## Database Drivers
        
*Libraries for connecting and operating databases.*

### Relational Databases
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [mysql](https://github.com/go-sql-driver/mysql) | 7053 | 393 | 2012-12-10 | 23 hours ago | Go MySQL Driver is a MySQL driver for Go's (golang) database/sql package |
| [pq](https://github.com/lib/pq) | 4685 | 157 | 2012-03-13 | 1 day ago | Pure Go Postgres driver for database/sql |
| [go-sqlite3](https://github.com/mattn/go-sqlite3) | 3092 | 125 | 2011-11-11 | 2 weeks ago | sqlite3 driver for go using database/sql |
| [pgx](https://github.com/jackc/pgx) | 1650 | 66 | 2013-03-31 | 4 days ago | PostgreSQL driver and toolkit for Go |
| [go-mssqldb](https://github.com/denisenkom/go-mssqldb) | 911 | 59 | 2013-12-16 | 2 days ago | Microsoft SQL server driver written in go language |
| [go-oci8](https://github.com/mattn/go-oci8) | 361 | 34 | 2012-02-29 | 4 weeks ago | oracle driver for go that using database/sql |
| [goracle](https://github.com/go-goracle/goracle) | 181 | 28 | 2015-03-25 | 1 week ago | Oracle driver for Go, using the ODPI-C driver |
| [firebirdsql](https://github.com/nakagami/firebirdsql) | 96 | 14 | 2013-08-27 | 3 days ago | Firebird RDBMS sql driver for Go (golang) |
| [go-adodb](https://github.com/mattn/go-adodb) | 84 | 12 | 2011-11-14 | 1 week ago | Microsoft ActiveX Object DataBase driver for go that using exp/sql |
| [gofreetds](https://github.com/minus5/gofreetds) | 84 | 21 | 2012-12-07 | 1 week ago | Go Sql Server database driver. |
| [calcite-avatica-go](https://github.com/apache/calcite-avatica-go) | 26 | 16 | 2017-08-08 | 1 month ago | Mirror of Apache Calcite - Avatica Go SQL Driver |
| [bgc](https://github.com/viant/bgc) | 10 | 9 | 2016-06-14 | 1 month ago | Datastore Connectivity for BigQuery in go |

### NoSQL Databases
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [cayley](https://github.com/cayleygraph/cayley) | 12256 | 633 | 2014-06-06 | 1 week ago | Graph database with support for multiple backends. |
| [cayley](https://github.com/cayleygraph/cayley) | 12256 | 633 | 2014-06-06 | 1 week ago | Graph database with support for multiple backends. |
| [cayley](https://github.com/cayleygraph/cayley) | 12255 | 633 | 2014-06-06 | 1 week ago | Graph database with support for multiple backends. |
| [redigo](https://github.com/gomodule/redigo) | 5551 | 280 | 2012-04-14 | 4 days ago | Go client for Redis |
| [redis](https://github.com/go-redis/redis) | 5154 | 200 | 2012-07-25 | 2 days ago | Type-safe Redis client for Golang |
| [bleve](https://github.com/blevesearch/bleve) | 5005 | 232 | 2014-04-18 | 1 day ago | A modern text indexing library for go |
| [riot](https://github.com/go-ego/riot) | 4304 | 171 | 2017-06-21 | 1 month ago | Go Open Source, Distributed, Simple and efficient Search Engine  |
| [elastic](https://github.com/olivere/elastic) | 3514 | 160 | 2012-12-07 | 1 day ago | Elasticsearch client for Go. |
| [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) | 2030 | 137 | 2017-02-09 | 16 hours ago | The Go driver for MongoDB |
| [mgo](https://github.com/globalsign/mgo) | 1428 | 77 | 2017-04-13 | 4 days ago | The MongoDB driver for Go |
| [rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1418 | 45 | 2013-09-12 | 1 month ago | Go language driver for RethinkDB. |
| [rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1418 | 45 | 2013-09-12 | 1 month ago | Go language driver for RethinkDB. |
| [rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1418 | 45 | 2013-09-12 | 1 month ago | Go language driver for RethinkDB. |
| [gomemcache](https://github.com/bradfitz/gomemcache) | 1011 | 51 | 2011-06-29 | 1 week ago | Go Memcached client library #golang |
| [elastigo](https://github.com/mattbaird/elastigo) | 937 | 49 | 2012-10-12 | 3 weeks ago | A Go (golang) based Elasticsearch client library. |
| [neoism](https://github.com/jmcvetta/neoism) | 350 | 23 | 2012-07-12 | 2 weeks ago | Neo4j client for Golang |
| [elasticsql](https://github.com/cch123/elasticsql) | 295 | 21 | 2016-08-24 | 1 week ago | convert sql to elasticsearch DSL in golang(go) |
| [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) | 289 | 39 | 2014-07-26 | 1 week ago | Aerospike Client Go  |
| [go-couchbase](https://github.com/couchbase/go-couchbase) | 286 | 26 | 2012-01-20 | 18 hours ago | Couchbase client in Go |
| [gocb](https://github.com/couchbase/gocb) | 281 | 70 | 2015-01-16 | 2 weeks ago | The Couchbase Go SDK |
| [redeo](https://github.com/bsm/redeo) | 240 | 24 | 2014-03-06 | 3 months ago | High-performance framework for building redis-protocol compatible TCP servers/services |
| [cachego](https://github.com/fabiorphp/cachego) | 103 | 6 | 2016-10-06 | 9 months ago | Golang Cache component - Multiple drivers |
| [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) | 72 | 6 | 2011-06-05 | 8 months ago | Neo4j REST Client in golang |
| [go-rejson](https://github.com/nitishm/go-rejson) | 67 | 4 | 2018-04-23 | 1 month ago | Golang client for redislabs' ReJSON module with support for multilple redis clients (redigo, go-redis) |
| [arangolite](https://github.com/solher/arangolite) | 64 | 5 | 2015-10-05 | 1 year ago | Lightweight Golang driver for ArangoDB |
| [dynago](https://github.com/underarmour/dynago) | 63 | 131 | 2015-05-18 | 1 year ago | A DynamoDB client for Go |
| [skizze](https://github.com/seiflotfy/skizze) | 59 | 6 | 2016-01-17 | 2 years ago | A probabilistic data structure service and storage |
| [go-couchdb](https://github.com/fjl/go-couchdb) | 51 | 6 | 2013-10-28 | 6 months ago | Yet another CouchDB HTTP API wrapper for Go |
| [gokv](https://github.com/philippgille/gokv) | 47 | 4 | 2018-10-09 | 1 day ago | Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more) |
| [goforestdb](https://github.com/couchbase/goforestdb) | 29 | 44 | 2014-05-14 | 2 years ago | Go bindings for ForestDB |
| [go-pilosa](https://github.com/pilosa/go-pilosa) | 27 | 13 | 2016-10-01 | 22 hours ago | Go client library for Pilosa |
| [goriak](https://github.com/zegl/goriak) | 25 | 2 | 2016-10-06 | 2 months ago | goriak - Go language driver for Riak KV |
| [goes](https://github.com/OwnLocal/goes) | 25 | 34 | 2015-12-29 | 2 years ago | A library to interact with Elasticsearch in Go! |
| [neo4j](https://github.com/cihangir/neo4j) | 24 | 3 | 2013-05-18 | 3 years ago | Neo4j Rest API Client for Go lang |
| [xredis](https://github.com/shomali11/xredis) | 9 | 1 | 2017-06-14 | 9 months ago | Go Redis Client |
| [dsc](https://github.com/viant/dsc) | 8 | 17 | 2016-06-14 | 1 week ago | Datastore Connectivity in go |
| [asc](https://github.com/viant/asc) | 4 | 10 | 2016-06-14 | 3 months ago | Datastore Connectivity for Aerospike for go |
| [godscache](https://github.com/defcronyke/godscache) | 4 | 2 | 2018-05-09 | 3 weeks ago | An unofficial Google Cloud Platform Go Datastore wrapper that adds caching using memcached. For App Engine Flexible, Compute Engine, Kubernetes Engine, and more. |

## Date and Time
        
*Libraries for working with dates and times.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [now](https://github.com/jinzhu/now) | 1982 | 51 | 2013-11-18 | 2 weeks ago | Now is a time toolkit for golang |
| [dateparse](https://github.com/araddon/dateparse) | 807 | 14 | 2014-04-21 | 1 week ago | GoLang Parse many date strings without knowing format in advance. |
| [carbon](https://github.com/uniplaces/carbon) | 306 | 38 | 2016-08-03 | 1 month ago | Carbon for Golang, an extension for Time |
| [durafmt](https://github.com/hako/durafmt) | 208 | 4 | 2016-05-21 | 6 months ago | :clock8: Better time duration formatting in Go!  |
| [timeutil](https://github.com/leekchan/timeutil) | 156 | 6 | 2015-08-02 | 1 month ago | timeutil - useful extensions (Timedelta, Strftime, ...) to the golang's time package |
| [timespan](https://github.com/SaidinWoT/timespan) | 60 | 5 | 2014-10-07 | 1 month ago | Golang package to manipulate time intervals. |
| [iso8601](https://github.com/relvacode/iso8601) | 60 | 2 | 2017-04-25 | 2 months ago | A fast ISO8601 date parser for Go |
| [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) | 48 | 3 | 2016-02-01 | 1 month ago | The implementation of the Persian (Solar Hijri) Calendar in Go (golang) |
| [date](https://github.com/rickb777/date) | 23 | 2 | 2015-11-24 | 3 months ago | A Go package for working with dates |
| [goweek](https://github.com/grsmv/goweek) | 18 | 3 | 2015-11-12 | 2 years ago | ISO 8601 compatible library for working with week entities for Go |
| [feiertage](https://github.com/wlbr/feiertage) | 17 | 1 | 2015-11-04 | 2 weeks ago | Gesetzliche Feiertage und mehr in Deutschland und Österreich (Bank holidays/public holidays in Austria and Germany) |
| [go-sunrise](https://github.com/nathan-osman/go-sunrise) | 10 | 3 | 2017-06-16 | 1 year ago | Go package for calculating the sunrise and sunset times for a given location |
| [kair](https://github.com/GuilhermeCaruso/kair) | 9 | 1 | 2018-10-03 | 1 week ago | :clock1: Date and Time - Golang Formatting Library |
| [nulltime](https://github.com/kirillDanshin/nulltime) | 9 | 1 | 2016-03-06 | 1 year ago |  |
| [tuesday](https://github.com/osteele/tuesday) | 7 | 1 | 2017-08-11 | 1 year ago | Ruby-compatible strftime for golang |
| [strftime](https://github.com/awoodbeck/strftime) | 5 | 1 | 2018-02-10 | 1 year ago | C99-compatible strftime formatter for use with Go time.Time instances. |

## Distributed Systems
        
*Packages that help with building Distributed Systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [etcd](https://github.com/etcd-io/etcd) | 23279 | 1228 | 2013-07-07 | 20 hours ago | Go implementation of the Raft consensus protocol, by CoreOS. |
| [etcd](https://github.com/etcd-io/etcd) | 23258 | 1228 | 2013-07-07 | 20 hours ago | Go implementation of the Raft consensus protocol, by CoreOS. |
| [etcd](https://github.com/etcd-io/etcd) | 23253 | 1228 | 2013-07-07 | 1 day ago | Go implementation of the Raft consensus protocol, by CoreOS. |
| [kit](https://github.com/go-kit/kit) | 12787 | 641 | 2015-02-03 | 14 hours ago | A standard library for microservices. |
| [grpc-go](https://github.com/grpc/grpc-go) | 7593 | 437 | 2014-12-09 | 14 hours ago | The Go language implementation of gRPC. HTTP/2 based RPC |
| [jaeger](https://github.com/jaegertracing/jaeger) | 7346 | 278 | 2016-04-16 | 4 hours ago | CNCF Jaeger, a Distributed Tracing System |
| [micro](https://github.com/micro/micro) | 5613 | 291 | 2015-01-17 | 5 days ago | A microservice toolkit |
| [gnatsd](https://github.com/nats-io/gnatsd) | 5314 | 327 | 2012-10-30 | 11 hours ago | High-Performance server for NATS, the cloud native messaging system. |
| [rpcx](https://github.com/smallnest/rpcx) | 3242 | 271 | 2016-05-18 | 4 days ago | Faster multil-language  bidirectional RPC framework in Go, like alibaba Dubbo, but with more features, Scale easily. |
| [tendermint](https://github.com/tendermint/tendermint) | 2713 | 240 | 2014-05-15 | 2 hours ago | ⟁ Tendermint Core (BFT Consensus) in Go |
| [torrent](https://github.com/anacrolix/torrent) | 2613 | 121 | 2015-01-09 | 2 days ago | Full-featured BitTorrent client package and utilities |
| [raft](https://github.com/hashicorp/raft) | 2504 | 177 | 2013-11-05 | 16 hours ago | Golang implementation of the Raft consensus protocol |
| [glow](https://github.com/chrislusf/glow) | 2354 | 141 | 2015-06-14 | 4 months ago | Glow is an easy-to-use distributed computation system written in Go, similar to Hadoop Map Reduce, Spark, Flink, Storm, etc. I am also working on another similar pure Go system, https://github.com/chrislusf/gleam , which is more flexible and more performant. |
| [gleam](https://github.com/chrislusf/gleam) | 1835 | 138 | 2016-08-26 | 2 months ago | Fast, efficient, and scalable distributed map/reduce system, DAG execution, in memory or on disk, written in pure Go, runs standalone or distributedly. |
| [emitter](https://github.com/emitter-io/emitter) | 1657 | 81 | 2016-10-29 | 1 day ago | High performance, distributed and low latency publish-subscribe platform. |
| [krakend](https://github.com/devopsfaith/krakend) | 1281 | 67 | 2016-11-05 | 22 hours ago | Ultra performant API Gateway with middlewares |
| [hprose-golang](https://github.com/hprose/hprose-golang) | 921 | 89 | 2014-02-14 | 9 months ago | Hprose is a cross-language RPC. This project is Hprose 2.0 for Golang. |
| [ringpop-go](https://github.com/uber/ringpop-go) | 531 | 2313 | 2015-06-06 | 2 weeks ago | Scalable, fault-tolerant application-layer sharding for Go applications |
| [gorpc](https://github.com/valyala/gorpc) | 522 | 26 | 2014-11-21 | 1 year ago | Simple, fast and scalable golang rpc library for high load |
| [go-health](https://github.com/InVisionApp/go-health) | 437 | 20 | 2017-11-30 | 3 months ago | Library for enabling asynchronous health checks in your service |
| [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) | 356 | 18 | 2015-10-09 | 1 week ago | Go Peerflix |
| [sleuth](https://github.com/ursiform/sleuth) | 292 | 7 | 2016-04-23 | 11 months ago | A Go library for master-less peer-to-peer autodiscovery and RPC between HTTP services |
| [digota](https://github.com/digota/digota) | 270 | 24 | 2017-08-14 | 4 months ago | ecommerce microservice |
| [go-jump](https://github.com/dgryski/go-jump) | 231 | 13 | 2014-06-16 | 1 year ago | go-jump: Jump consistent hashing |
| [consistent](https://github.com/buraksezer/consistent) | 152 | 9 | 2018-03-25 | 1 week ago | Consistent hashing with bounded loads in Golang |
| [redis-lock](https://github.com/bsm/redis-lock) | 109 | 6 | 2014-10-10 | 3 months ago |  |
| [jsonrpc](https://github.com/osamingo/jsonrpc) | 104 | 6 | 2016-10-28 | 2 months ago | The jsonrpc package helps implement of JSON-RPC 2.0 |
| [jsonrpc](https://github.com/ybbus/jsonrpc) | 79 | 6 | 2016-11-10 | 3 months ago | A simple go implementation of json rpc 2.0 client over http |
| [celeriac.v1](https://github.com/svcavallar/celeriac.v1) | 47 | 3 | 2015-10-10 | 2 months ago | Golang client library for adding support for interacting and monitoring Celery workers, tasks and events. |
| [doublejump](https://github.com/edwingeng/doublejump) | 28 | 3 | 2018-06-27 | 2 months ago | A revamped Google's jump consistent hash |
| [drmaa](https://github.com/dgruber/drmaa) | 23 | 2 | 2013-03-17 | 9 months ago | Compute cluster (HPC) job submission library for Go (#golang) based on the open DRMAA standard. |
| [flowgraph](https://github.com/vectaport/flowgraph) | 9 | 1 | 2018-08-30 | 3 days ago | Flowgraph package for scalable asynchronous system development |
| [dynatomic](https://github.com/tylfin/dynatomic) | 6 | 0 | 2019-02-09 | 1 week ago | Dynatomic is a library for using dynamodb as an atomic counter |
| [outboxer](https://github.com/italolelis/outboxer) | 0 | 0 | 2019-02-01 | 1 week ago | A library that implements the outboxer pattern in go |

## Email
        
*Libraries and tools that implement email creation and sending.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [MailHog](https://github.com/mailhog/MailHog) | 4517 | 125 | 2014-04-17 | 3 days ago | Web and API based SMTP testing |
| [gomail](https://github.com/go-gomail/gomail) | 2139 | 65 | 2014-10-15 | 2 weeks ago | The best way to send emails in Go. |
| [hermes](https://github.com/matcornic/hermes) | 1480 | 24 | 2017-03-26 | 6 days ago | Golang package that generates clean, responsive HTML e-mails for sending transactional mail |
| [email](https://github.com/jordan-wright/email) | 1007 | 35 | 2013-12-13 | 2 weeks ago | Robust and flexible email library for Go |
| [go-imap](https://github.com/emersion/go-imap) | 625 | 31 | 2016-04-27 | 13 hours ago |  :inbox_tray: An IMAP library for clients and servers |
| [sendgrid-go](https://github.com/sendgrid/sendgrid-go) | 469 | 181 | 2013-09-12 | 2 days ago | The Official SendGrid Led, Community Driven Golang API Library |
| [hectane](https://github.com/hectane/hectane) | 160 | 12 | 2015-08-28 | 10 months ago | Lightweight SMTP client written in Go |
| [douceur](https://github.com/aymerick/douceur) | 142 | 4 | 2015-04-09 | 11 months ago | A simple CSS parser and inliner in Go |
| [go-message](https://github.com/emersion/go-message) | 79 | 6 | 2016-12-31 | 1 month ago | :envelope: A streaming Go library for the Internet Message Format and mail messages |
| [smtp](https://github.com/mailhog/smtp) | 49 | 9 | 2014-12-25 | 8 months ago | MailHog SMTP Protocol |
| [go-dkim](https://github.com/toorop/go-dkim) | 47 | 3 | 2015-04-29 | 1 year ago | DKIM package for golang |

## Embeddable Scripting Languages
        
*Embedding other languages inside your go code.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [otto](https://github.com/robertkrimen/otto) | 4409 | 188 | 2012-10-06 | 1 week ago | A JavaScript interpreter in Go (golang) |
| [gopher-lua](https://github.com/yuin/gopher-lua) | 2612 | 137 | 2015-02-15 | 3 weeks ago | GopherLua: VM and compiler for Lua in Go |
| [go-lua](https://github.com/Shopify/go-lua) | 1542 | 252 | 2013-12-21 | 3 months ago | A Lua VM in Go |
| [tengo](https://github.com/d5/tengo) | 1056 | 25 | 2019-01-09 | 17 hours ago | A fast script language for Go |
| [anko](https://github.com/mattn/anko) | 843 | 47 | 2014-03-28 | 1 week ago | Scriptable interpreter written in golang |
| [go-python](https://github.com/sbinet/go-python) | 809 | 43 | 2012-07-09 | 1 month ago | naive go bindings to the CPython C-API |
| [go-duktape](https://github.com/olebedev/go-duktape) | 628 | 23 | 2015-01-08 | 2 weeks ago | Duktape JavaScript engine bindings for Go |
| [go-php](https://github.com/deuill/go-php) | 610 | 39 | 2015-09-18 | 4 months ago | PHP bindings for the Go programming language (Golang) |
| [golua](https://github.com/aarzilli/golua) | 423 | 35 | 2010-12-07 | 6 months ago | Go bindings for Lua C API - in progress |
| [gisp](https://github.com/jcla1/gisp) | 415 | 20 | 2014-01-11 | 1 year ago | Simple LISP in Go |
| [agora](https://github.com/mna/agora) | 317 | 22 | 2013-06-19 | 4 years ago | Dynamically typed, embeddable programming language in Go. |
| [agora](https://github.com/mna/agora) | 317 | 22 | 2013-06-19 | 4 years ago | Dynamically typed, embeddable programming language in Go. |
| [agora](https://github.com/mna/agora) | 317 | 22 | 2013-06-19 | 4 years ago | Dynamically typed, embeddable programming language in Go. |
| [expr](https://github.com/antonmedv/expr) | 297 | 16 | 2018-07-14 | 1 month ago | Expr is an engine that can evaluate expressions. |
| [gval](https://github.com/PaesslerAG/gval) | 85 | 12 | 2017-09-27 | 2 weeks ago | Expression evaluation in golang |
| [purl](https://github.com/ian-kent/purl) | 25 | 2 | 2014-11-30 | 4 years ago | Perl, but fluffy like a cat! |
| [binder](https://github.com/alexeyco/binder) | 24 | 2 | 2017-04-03 | 7 months ago | High level go to Lua binder. Write less, do more. |
| [ngaro](https://github.com/db47h/ngaro) | 16 | 1 | 2016-08-09 | 9 months ago | An embeddable implementation of the Ngaro Virtual Machine for Go programs |

## Error Handling
        
*Libraries for handling errors.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [errors](https://github.com/pkg/errors) | 4168 | 92 | 2015-12-27 | 6 days ago | Simple error handling primitives |
| [go-multierror](https://github.com/hashicorp/go-multierror) | 558 | 83 | 2014-12-16 | 2 months ago | A Go (golang) package for representing a list of errors as a single error. |
| [errorx](https://github.com/joomcode/errorx) | 507 | 38 | 2018-08-17 | 2 weeks ago | A comprehensive error handling library for Go |
| [tracerr](https://github.com/ztrue/tracerr) | 174 | 6 | 2019-02-07 | 2 weeks ago | Golang errors with stack trace and source fragments. |
| [werr](https://github.com/txgruppi/werr) | 10 | 0 | 2015-08-04 | 3 years ago | Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called. |

## Files
        
*Libraries for  handling files and file systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [afero](https://github.com/spf13/afero) | 1944 | 84 | 2014-10-28 | 1 hour ago | A FileSystem Abstraction System for Go |
| [pdfcpu](https://github.com/hhrutter/pdfcpu) | 797 | 23 | 2017-06-19 | 13 hours ago | A PDF processor written in Go. |
| [notify](https://github.com/rjeczalik/notify) | 452 | 26 | 2014-09-09 | 2 weeks ago | File system event notification library on steroids. |
| [opc](https://github.com/qmuntal/opc) | 50 | 1 | 2018-11-06 | 1 week ago | Golang  implementation of the Open Packaging Conventions (OPC) |
| [skywalker](https://github.com/dixonwille/skywalker) | 41 | 3 | 2017-08-02 | 1 year ago | A package to allow one to concurrently go through a filesystem with ease |
| [go-csv-tag](https://github.com/artonge/go-csv-tag) | 39 | 1 | 2017-06-18 | 4 months ago | Read csv file from go using tags |
| [tarfs](https://github.com/posener/tarfs) | 31 | 2 | 2017-03-11 | 1 year ago | An implementation of the FileSystem interface for tar files. |
| [go-gtfs](https://github.com/artonge/go-gtfs) | 13 | 2 | 2017-07-09 | 1 month ago | Load GTFS files in golang |
| [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) | 11 | 1 | 2018-10-16 | 4 months ago | copy files for humans |

## Financial
        
*Packages for accounting and finance.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [decimal](https://github.com/shopspring/decimal) | 1339 | 56 | 2015-02-26 | 1 day ago | Arbitrary-precision fixed-point decimal numbers in go |
| [go-money](https://github.com/Rhymond/go-money) | 552 | 15 | 2017-03-21 | 1 month ago | Implementation of Fowler's Money pattern. |
| [go-money](https://github.com/Rhymond/go-money) | 552 | 15 | 2017-03-21 | 1 month ago | Implementation of Fowler's Money pattern. |
| [go-money](https://github.com/Rhymond/go-money) | 552 | 15 | 2017-03-21 | 1 month ago | Implementation of Fowler's Money pattern. |
| [go-finance](https://github.com/FlashBoys/go-finance) | 538 | 28 | 2016-02-28 | 1 year ago | :warning: Deprecrated in favor of https://github.com/piquette/finance-go :warning: |
| [accounting](https://github.com/leekchan/accounting) | 443 | 12 | 2015-08-10 | 1 month ago | money and currency formatting for golang |
| [techan](https://github.com/sdcoffey/techan) | 108 | 17 | 2017-03-08 | 2 months ago | Technical Analysis Library for Golang |
| [vat](https://github.com/dannyvankooten/vat) | 54 | 1 | 2016-06-19 | 5 months ago | Go package for dealing with EU VAT. Does VAT number validation & rates retrieval. |
| [ofxgo](https://github.com/aclindsa/ofxgo) | 53 | 5 | 2015-11-08 | 2 days ago | Golang library for querying and parsing OFX |
| [transaction](https://github.com/claygod/transaction) | 46 | 8 | 2017-10-11 | 6 months ago | Embedded database for accounts transactions. |
| [go-finance](https://github.com/alpeb/go-finance) | 33 | 2 | 2017-06-01 | 10 months ago | Go library containing a collection of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations. |
| [orderbook](https://github.com/i25959341/orderbook) | 22 | 3 | 2018-04-25 | 4 days ago | Matching Engine for Limit Order Book in Golang |

## Forms
        
*Libraries for working with forms.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [nosurf](https://github.com/justinas/nosurf) | 920 | 35 | 2013-08-23 | 1 month ago | CSRF protection middleware for Go. |
| [binding](https://github.com/mholt/binding) | 731 | 30 | 2014-05-21 | 11 months ago | Reflectionless data binding for Go's net/http (not yet a stable 1.0, but not likely to change much either) |
| [csrf](https://github.com/gorilla/csrf) | 371 | 19 | 2015-08-03 | 2 months ago | gorilla/csrf provides Cross Site Request Forgery (CSRF) prevention middleware for Go web applications & services. |
| [form](https://github.com/go-playground/form) | 324 | 8 | 2016-05-26 | 1 month ago | :steam_locomotive: Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support. |
| [conform](https://github.com/leebenson/conform) | 163 | 5 | 2016-01-06 | 8 months ago | Trims, sanitizes & scrubs data based on struct tags (go, golang) |
| [formam](https://github.com/monoculum/formam) | 111 | 2 | 2014-10-25 | 6 months ago | a package for decode form's values into struct in Go |
| [forms](https://github.com/albrow/forms) | 96 | 6 | 2014-08-08 | 1 year ago | A lightweight go library for parsing form data or json from an http.Request. |
| [bind](https://github.com/robfig/bind) | 23 | 2 | 2014-08-06 | 4 years ago |  |

## Functional
        
*Packages to support functional programming in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [go-underscore](https://github.com/tobyhede/go-underscore) | 1011 | 26 | 2014-07-02 | 2 weeks ago |  Helpfully Functional Go -  A useful collection of Go utilities. Designed for programmer happiness.  |
| [fpGo](https://github.com/TeaEntityLab/fpGo) | 71 | 1 | 2018-05-24 | 7 months ago | Monad, Functional Programming features for Golang |
| [fuego](https://github.com/seborama/fuego) | 16 | 1 | 2018-11-06 | 14 hours ago | Functional Experiment in Go |

## Game Development
        
*Awesome game development libraries.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [leaf](https://github.com/name5566/leaf) | 2725 | 302 | 2014-08-04 | 2 months ago | A game server framework in Go (golang) |
| [pixel](https://github.com/faiface/pixel) | 2118 | 92 | 2016-11-19 | 1 week ago | A hand-crafted 2D game library in Go |
| [ebiten](https://github.com/hajimehoshi/ebiten) | 1492 | 65 | 2013-06-16 | 20 hours ago | A dead simple 2D game library in Go |
| [go-sdl2](https://github.com/veandco/go-sdl2) | 1051 | 42 | 2013-06-06 | 1 day ago | SDL2 binding for Go |
| [gonet](https://github.com/xtaci/gonet) | 998 | 125 | 2013-04-11 | 1 year ago | A Game Server Skeleton in golang. |
| [termloop](https://github.com/JoelOtter/termloop) | 977 | 35 | 2015-05-24 | 2 days ago | Terminal-based game engine for Go, built on top of Termbox |
| [engo](https://github.com/EngoEngine/engo) | 954 | 45 | 2014-11-12 | 3 days ago | Engo is an open-source 2D game engine written in Go. |
| [goworld](https://github.com/xiaonanln/goworld) | 947 | 100 | 2017-06-03 | 2 months ago | Scalable Distributed Game Server Engine with Hot Swapping in Golang |
| [nano](https://github.com/lonng/nano) | 846 | 52 | 2017-08-02 | 1 month ago | Lightweight, facility, high performance golang based game server framework |
| [engine](https://github.com/g3n/engine) | 596 | 59 | 2017-03-08 | 2 weeks ago | Go 3D Game Engine |
| [oak](https://github.com/oakmound/oak) | 573 | 34 | 2017-07-16 | 1 day ago | A pure Go game engine |
| [engine](https://github.com/azul3d/engine) | 406 | 23 | 2016-02-29 | 8 months ago | Azul3D - A 3D game engine written in Go! |
| [raylib-go](https://github.com/gen2brain/raylib-go) | 341 | 18 | 2017-01-27 | 3 months ago | Go bindings for raylib, a simple and easy-to-use library to learn videogames programming. |
| [go-astar](https://github.com/beefsack/go-astar) | 307 | 10 | 2014-05-28 | 11 months ago | Go implementation of the A* search algorithm |
| [GarageEngine](https://github.com/vova616/GarageEngine) | 303 | 31 | 2012-08-07 | 5 years ago | Game engine written in Go (golang). |
| [pitaya](https://github.com/topfreegames/pitaya) | 204 | 26 | 2018-03-20 | 4 days ago | Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. |
| [go3d](https://github.com/ungerik/go3d) | 150 | 10 | 2011-06-27 | 4 days ago | A performance oriented 2D/3D math package for Go |
| [glop](https://github.com/runningwild/glop) | 76 | 4 | 2011-04-21 | 3 years ago | Bare-bones osx alternative to sdl |
| [go-collada](https://github.com/GlenKelley/go-collada) | 12 | 3 | 2013-09-19 | 5 years ago | Go package for working with the Collada file format. |

## Generation and Generics
        
*Tools to enhance the language with features like generics via code generation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [go-linq](https://github.com/ahmetb/go-linq) | 1661 | 72 | 2013-12-19 | 3 months ago | .NET LINQ-like query methods for Go. |
| [go-linq](https://github.com/ahmetb/go-linq) | 1661 | 72 | 2013-12-19 | 3 months ago | .NET LINQ-like query methods for Go. |
| [go-linq](https://github.com/ahmetb/go-linq) | 1661 | 72 | 2013-12-19 | 3 months ago | .NET LINQ-like query methods for Go. |
| [jennifer](https://github.com/dave/jennifer) | 1138 | 21 | 2016-12-05 | 1 month ago | Jennifer is a code generator for Go |
| [gen](https://github.com/clipperhouse/gen) | 977 | 36 | 2013-10-14 | 8 months ago | Type-driven code generation for Go |
| [goderive](https://github.com/awalterschulze/goderive) | 671 | 22 | 2017-02-11 | 1 week ago | Functional programming for Golang |
| [gowrap](https://github.com/hexdigest/gowrap) | 206 | 7 | 2018-09-15 | 1 week ago | GoWrap is a command line tool for generating decorators for Go interfaces |
| [interfaces](https://github.com/rjeczalik/interfaces) | 160 | 5 | 2015-12-06 | 2 months ago | Code generation tools for Go. |
| [pkgreflect](https://github.com/ungerik/pkgreflect) | 78 | 5 | 2012-09-03 | 1 year ago | A Go preprocessor for package scoped reflection |
| [go-enum](https://github.com/abice/go-enum) | 66 | 4 | 2017-08-11 | 2 weeks ago | An enum generator for go |
| [efaceconv](https://github.com/t0pep0/efaceconv) | 40 | 4 | 2016-11-18 | 1 year ago |  |
| [gotype](https://github.com/wzshiming/gotype) | 18 | 3 | 2017-12-05 | 1 week ago | Golang source code parsing, usage like reflect package |

## Geographic
        
*Geographic tools and servers*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [tile38](https://github.com/tidwall/tile38) | 5655 | 181 | 2016-03-05 | 19 hours ago | Tile38 is a geospatial database, spatial index, and realtime geofence. 🌐 |
| [geo](https://github.com/golang/geo) | 803 | 73 | 2014-12-04 | 1 week ago | S2 geometry library in Go |
| [geocache](https://github.com/melihmucuk/geocache) | 95 | 5 | 2016-06-21 | 2 years ago | Geocache is an in-memory cache that is suitable for geolocation based applications. |
| [osm](https://github.com/paulmach/osm) | 41 | 5 | 2016-02-02 | 2 months ago | General purpose library for reading, writing and working with OpenStreetMap data |
| [geoserver](https://github.com/hishamkaram/geoserver) | 19 | 2 | 2018-03-27 | 3 months ago | geoserver is a Go library for manipulating a GeoServer instance via the GeoServer REST API. |
| [pbf](https://github.com/maguro/pbf) | 10 | 1 | 2017-09-19 | 4 months ago | OpenStreetMap PBF golang parser |
| [gismanager](https://github.com/hishamkaram/gismanager) | 9 | 1 | 2018-09-29 | 4 months ago | Publish Your GIS Data(Vector Data) to PostGIS and Geoserver |

## Go Compilers
        
*Tools for compiling Go to other languages.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gopherjs](https://github.com/gopherjs/gopherjs) | 8000 | 249 | 2013-08-28 | 2 hours ago | A compiler from Go to JavaScript for running Go code in a browser |
| [llgo](https://github.com/go-llvm/llgo) | 952 | 61 | 2011-11-05 | 4 years ago | LLVM-based compiler for Go |
| [tardisgo](https://github.com/tardisgo/tardisgo) | 382 | 27 | 2014-01-08 | 2 years ago | Golang->Haxe->CPP/CSharp/Java/JavaScript transpiler   |
| [c4go](https://github.com/Konstantin8105/c4go) | 69 | 8 | 2018-03-28 | 6 hours ago | Transpiling C code to Go code |
| [f4go](https://github.com/Konstantin8105/f4go) | 9 | 2 | 2018-07-09 | 2 months ago | Transpiling fortran code to golang code |

## Goroutines
        
*Tools for managing and working with Goroutines.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [goworker](https://github.com/benmanns/goworker) | 2135 | 70 | 2013-07-23 | 3 weeks ago | goworker is a Go-based background worker that runs 10 to 100,000* times faster than Ruby-based workers. |
| [ants](https://github.com/panjf2000/ants) | 1365 | 54 | 2018-05-19 | 1 week ago | 🐜⚡️A high-performance goroutine pool for Go, inspired by fasthttp. |
| [tunny](https://github.com/Jeffail/tunny) | 1138 | 51 | 2014-04-03 | 3 months ago | A goroutine pool for Go |
| [grpool](https://github.com/ivpusic/grpool) | 460 | 29 | 2015-07-22 | 1 month ago | Lightweight Goroutine pool |
| [pool](https://github.com/go-playground/pool) | 452 | 12 | 2015-10-29 | 2 years ago | :speedboat: a limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation |
| [go-floc](https://github.com/workanator/go-floc) | 164 | 8 | 2017-07-03 | 1 year ago | Floc: Orchestrate goroutines with ease. |
| [go-flow](https://github.com/kamildrazkiewicz/go-flow) | 94 | 7 | 2016-09-25 | 1 year ago | Simply way to control goroutines execution order based on dependencies |
| [workerpool](https://github.com/gammazero/workerpool) | 80 | 5 | 2016-05-17 | 2 months ago | Concurrency limiting goroutine pool |
| [semaphore](https://github.com/marusama/semaphore) | 64 | 4 | 2017-11-22 | 1 month ago | Fast resizable golang semaphore |
| [semaphore](https://github.com/kamilsk/semaphore) | 61 | 1 | 2016-10-08 | 3 days ago | 🚦 Semaphore pattern implementation with timeout of lock/unlock operations. |
| [GoSlaves](https://github.com/dgrr/GoSlaves) | 60 | 4 | 2017-09-18 | 7 months ago | Simple and Asynchronous Goroutine pool library. |
| [GoSlaves](https://github.com/dgrr/GoSlaves) | 60 | 4 | 2017-09-18 | 7 months ago | Simple and Asynchronous Goroutine pool library. |
| [GoSlaves](https://github.com/dgrr/GoSlaves) | 60 | 4 | 2017-09-18 | 7 months ago | Simple and Asynchronous Goroutine pool library. |
| [gpool](https://github.com/sherifabdlnaby/gpool) | 47 | 1 | 2018-12-03 | 1 day ago | manages a resizeable pool of context-aware goroutines to bound concurrency. |
| [gpool](https://github.com/sherifabdlnaby/gpool) | 47 | 1 | 2018-12-03 | 1 day ago | manages a resizeable pool of context-aware goroutines to bound concurrency. |
| [gpool](https://github.com/sherifabdlnaby/gpool) | 46 | 1 | 2018-12-03 | 1 day ago | manages a resizeable pool of context-aware goroutines to bound concurrency. |
| [worker-pool](https://github.com/vardius/worker-pool) | 32 | 3 | 2017-10-04 | 4 weeks ago | Go simple async worker pool |
| [parallel-fn](https://github.com/rafaeljesus/parallel-fn) | 24 | 3 | 2017-06-18 | 1 year ago | Run functions in parallel :comet: |
| [cyclicbarrier](https://github.com/marusama/cyclicbarrier) | 23 | 3 | 2018-01-11 | 4 months ago | CyclicBarrier golang implementation |
| [async](https://github.com/StudioSol/async) | 14 | 9 | 2017-07-01 | 5 months ago | A safe way to execute functions asynchronously, recovering them in case of panic. |
| [async](https://github.com/StudioSol/async) | 14 | 9 | 2017-07-01 | 5 months ago | A safe way to execute functions asynchronously, recovering them in case of panic. |
| [async](https://github.com/StudioSol/async) | 14 | 9 | 2017-07-01 | 5 months ago | A safe way to execute functions asynchronously, recovering them in case of panic. |
| [threadpool](https://github.com/shettyh/threadpool) | 12 | 2 | 2017-09-07 | 1 year ago | Golang simple thread pool implementation |
| [breaker](https://github.com/kamilsk/breaker) | 7 | 1 | 2019-02-15 | 1 day ago | 🚧 Flexible mechanism to make your code breakable. |
| [artifex](https://github.com/borderstech/artifex) | 6 | 2 | 2018-11-01 | 4 months ago | Simple in-memory job queue for Golang using worker-based dispatching |
| [stl](https://github.com/ssgreg/stl) | 5 | 1 | 2018-06-19 | 5 months ago | Software Transactional Locks |
| [go-trylock](https://github.com/subchen/go-trylock) | 4 | 1 | 2018-04-26 | 9 months ago | TryLock support on read-write lock for Golang |

## GUI
        
*Interaction*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [ui](https://github.com/andlabs/ui) | 6460 | 355 | 2014-02-18 | 3 weeks ago | Platform-native GUI library for Go. |
| [qt](https://github.com/therecipe/qt) | 5269 | 273 | 2014-11-19 | 1 month ago | Qt binding for Go (Golang) with support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi / AsteroidOS / Ubuntu Touch / JavaScript / WebAssembly |
| [robotgo](https://github.com/go-vgo/robotgo) | 3919 | 176 | 2016-09-27 | 22 hours ago | RobotGo, Go Native cross-platform GUI automation |
| [webview](https://github.com/zserge/webview) | 3899 | 173 | 2017-08-19 | 2 days ago | Tiny cross-platform webview library for C/C++/Golang. Uses WebKit (Gtk/Cocoa) and MSHTML (Windows) |
| [walk](https://github.com/lxn/walk) | 3279 | 240 | 2010-09-16 | 5 days ago | A Windows GUI toolkit for the Go Programming Language |
| [app](https://github.com/maxence-charriere/app) | 2752 | 106 | 2016-10-12 | 4 days ago | Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. |
| [app](https://github.com/maxence-charriere/app) | 2751 | 106 | 2016-10-12 | 4 days ago | Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. |
| [app](https://github.com/maxence-charriere/app) | 2750 | 106 | 2016-10-12 | 4 days ago | Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. |
| [go-astilectron](https://github.com/asticode/go-astilectron) | 2363 | 112 | 2017-04-22 | 1 month ago | Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron) |
| [go-sciter](https://github.com/sciter-sdk/go-sciter) | 1298 | 114 | 2015-10-15 | 4 months ago | Golang bindings of Sciter: the Embeddable HTML/CSS/script engine for modern UI development |
| [fyne](https://github.com/fyne-io/fyne) | 776 | 42 | 2018-02-05 | 1 day ago | Cross platform GUI in Go based on Material Design |
| [systray](https://github.com/getlantern/systray) | 637 | 38 | 2014-11-12 | 2 weeks ago | a cross platfrom Go library to place an icon and menu in the notification area |
| [gotk3](https://github.com/gotk3/gotk3) | 623 | 43 | 2015-08-13 | 9 hours ago | Go bindings for GTK3 |
| [gosx-notifier](https://github.com/deckarep/gosx-notifier) | 476 | 15 | 2013-11-25 | 1 year ago | gosx-notifier is a Go framework for sending desktop notifications to OSX 10.8 or higher |
| [gowd](https://github.com/dtylman/gowd) | 173 | 19 | 2017-03-29 | 4 months ago | Build cross platform GUI apps with GO and HTML/JS/CSS (powered by nwjs) |
| [trayhost](https://github.com/shurcooL/trayhost) | 143 | 4 | 2014-04-25 | 4 months ago | Cross-platform Go library to place an icon in the host operating system's taskbar. |

## Hardware
        
*Libraries, tools, and tutorials for interacting with hardware.*

## Images
        
*Libraries for manipulating images.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [ln](https://github.com/fogleman/ln) | 2378 | 89 | 2016-01-10 | 1 year ago | 3D line art engine. |
| [imaginary](https://github.com/h2non/imaginary) | 2313 | 70 | 2015-03-05 | 1 week ago | Fast, simple, scalable HTTP microservice for high-level image processing with first-class Docker support |
| [imaging](https://github.com/disintegration/imaging) | 2190 | 63 | 2012-12-07 | 1 month ago | Imaging is a simple image processing package for Go |
| [resize](https://github.com/nfnt/resize) | 2024 | 73 | 2012-08-03 | 1 year ago | Pure golang image resizing  |
| [gocv](https://github.com/hybridgroup/gocv) | 1999 | 103 | 2017-09-19 | 1 week ago | Go package for computer vision using OpenCV 4 and beyond. |
| [bild](https://github.com/anthonynsimon/bild) | 1931 | 55 | 2016-08-01 | 5 months ago | A collection of parallel image processing algorithms in pure Go |
| [pt](https://github.com/fogleman/pt) | 1718 | 62 | 2015-01-24 | 1 year ago | A path tracer written in Go. |
| [gg](https://github.com/fogleman/gg) | 1711 | 68 | 2016-02-19 | 1 week ago | Go Graphics - 2D rendering in Go with a simple API. |
| [svgo](https://github.com/ajstarks/svgo) | 1222 | 47 | 2010-03-06 | 5 months ago | Go Language Library for SVG generation |
| [smartcrop](https://github.com/muesli/smartcrop) | 1203 | 30 | 2014-04-08 | 4 months ago | smartcrop finds good image crops for arbitrary crop sizes |
| [gift](https://github.com/disintegration/gift) | 1157 | 49 | 2014-07-13 | 5 months ago | Go Image Filtering Toolkit |
| [go-opencv](https://github.com/go-opencv/go-opencv) | 1028 | 68 | 2013-12-09 | 1 month ago | Go bindings for OpenCV. |
| [go-opencv](https://github.com/go-opencv/go-opencv) | 1027 | 68 | 2013-12-09 | 1 month ago | Go bindings for OpenCV. |
| [go-opencv](https://github.com/go-opencv/go-opencv) | 1027 | 68 | 2013-12-09 | 1 month ago | Go bindings for OpenCV. |
| [geopattern](https://github.com/pravj/geopattern) | 983 | 20 | 2014-10-23 | 1 month ago | :triangular_ruler: Create beautiful generative image patterns from a string in golang. |
| [picfit](https://github.com/thoas/picfit) | 968 | 41 | 2014-12-07 | 1 month ago | An image resizing server written in Go |
| [imagick](https://github.com/gographics/imagick) | 884 | 54 | 2013-05-01 | 1 month ago | Go binding to ImageMagick's MagickWand C API |
| [bimg](https://github.com/h2non/bimg) | 693 | 35 | 2015-03-17 | 2 days ago | Small Go package for fast high-level image processing powered by libvips C library |
| [mort](https://github.com/aldor007/mort) | 343 | 17 | 2017-11-19 | 2 weeks ago | Storage and image processing server written in Go |
| [stegify](https://github.com/DimitarPetrov/stegify) | 281 | 13 | 2018-11-30 | 2 weeks ago | Go tool for LSB steganography, capable of hiding any file within an image. |
| [govatar](https://github.com/o1egl/govatar) | 280 | 5 | 2016-01-18 | 10 months ago | Avatar generator library for GO language |
| [go-nude](https://github.com/koyachi/go-nude) | 272 | 17 | 2014-05-02 | 3 months ago | Nudity detection with Go. |
| [image2ascii](https://github.com/qeesung/image2ascii) | 238 | 7 | 2018-10-20 | 3 months ago | :foggy: Convert image to ASCII |
| [goimagehash](https://github.com/corona10/goimagehash) | 168 | 9 | 2017-07-29 | 3 weeks ago | Go Perceptual image hashing package |
| [rez](https://github.com/bamiaux/rez) | 156 | 9 | 2014-01-17 | 1 year ago | Image resizing in pure Go and SIMD |
| [img](https://github.com/hawx/img) | 127 | 5 | 2012-07-29 | 3 years ago | A selection of image manipulation tools |
| [go-cairo](https://github.com/ungerik/go-cairo) | 83 | 6 | 2012-08-23 | 5 months ago | Go binding for the cairo graphics library |
| [go-gd](https://github.com/bolknote/go-gd) | 47 | 4 | 2011-05-12 | 10 months ago | Go bingings for GD (http://www.boutell.com/gd/) |
| [mergi](https://github.com/noelyahan/mergi) | 41 | 2 | 2018-09-24 | 1 week ago | Fun and lightweight image programming tool + library (merge, crop, resize, watermark, animate, easing) |
| [go-webcolors](https://github.com/jyotiska/go-webcolors) | 23 | 1 | 2014-04-24 | 3 years ago | Port of webcolors library from Python to Go |
| [tga](https://github.com/ftrvxmtrx/tga) | 21 | 2 | 2012-10-08 | 3 years ago | Go package for decoding and encoding TARGA image format |
| [cameron](https://github.com/aofei/cameron) | 17 | 1 | 2018-05-06 | 1 week ago | An avatar generator for Go. |
| [steganography](https://github.com/auyer/steganography) | 12 | 3 | 2018-05-22 | 4 months ago | Pure Golang Library that allows simple LSB steganography on images |
| [mpo](https://github.com/donatj/mpo) | 5 | 1 | 2015-04-15 | 1 month ago | JPEG-MPO Decoder / Converter Library and CLI Tool |

## IoT
        
*Libraries for programming devices of the IoT.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gobot](https://github.com/hybridgroup/gobot) | 5204 | 298 | 2013-09-21 | 2 weeks ago | Golang framework for robotics, drones, and the Internet of Things (IoT) |
| [flogo](https://github.com/TIBCOSoftware/flogo) | 959 | 113 | 2016-07-10 | 4 days ago | Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. |
| [flogo](https://github.com/TIBCOSoftware/flogo) | 955 | 113 | 2016-07-10 | 4 days ago | Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. |
| [flogo](https://github.com/TIBCOSoftware/flogo) | 954 | 113 | 2016-07-10 | 4 days ago | Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. |
| [gatt](https://github.com/paypal/gatt) | 770 | 57 | 2014-04-23 | 3 months ago | Gatt is a Go package for building Bluetooth Low Energy peripherals |
| [mainflux](https://github.com/mainflux/mainflux) | 366 | 54 | 2015-07-07 | 15 hours ago | Industrial IoT Messaging and Device Management Server. |
| [mainflux](https://github.com/mainflux/mainflux) | 365 | 53 | 2015-07-07 | 15 hours ago | Industrial IoT Messaging and Device Management Server. |
| [mainflux](https://github.com/mainflux/mainflux) | 363 | 53 | 2015-07-07 | 1 day ago | Industrial IoT Messaging and Device Management Server. |
| [devices](https://github.com/goiot/devices) | 221 | 17 | 2016-05-30 | 2 years ago | Suite of libraries for IoT devices (written in Go), experimental for x/exp/io |
| [sensorbee](https://github.com/sensorbee/sensorbee) | 165 | 19 | 2016-02-19 | 1 year ago | Lightweight stream processing engine for IoT |
| [connectordb](https://github.com/connectordb/connectordb) | 152 | 17 | 2015-01-17 | 4 days ago | An Open-Source Platform for Quantified Self & IoT |
| [huego](https://github.com/amimof/huego) | 85 | 1 | 2017-05-16 | 6 days ago | An extensive Philips Hue client library for Go with an emphasis on simplicity |
| [eywa](https://github.com/xcodersun/eywa) | 30 | 8 | 2016-02-21 | 1 year ago | Make IoT a lot more fun with data.  |
| [iot](https://github.com/vaelen/iot) | 28 | 4 | 2018-03-08 | 10 months ago | A Golang client for Google IoT Core |

## Job Scheduler
        
*Libraries for scheduling jobs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gron](https://github.com/roylee0704/gron) | 552 | 16 | 2016-06-04 | 1 week ago | gron, Cron Jobs in Go. |
| [jobrunner](https://github.com/bamzi/jobrunner) | 518 | 17 | 2015-10-21 | 2 years ago | Framework for performing work asynchronously, outside of the request flow |
| [jobs](https://github.com/albrow/jobs) | 438 | 18 | 2015-02-10 | 8 months ago | A persistent and flexible background jobs library for go. |
| [scheduler](https://github.com/carlescere/scheduler) | 266 | 16 | 2015-02-04 | 8 months ago | Job scheduling made easy. |
| [go-cron](https://github.com/rk/go-cron) | 169 | 9 | 2011-04-15 | 3 years ago | A simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. |
| [clockwork](https://github.com/whiteShtef/clockwork) | 51 | 1 | 2018-04-24 | 2 weeks ago | ⏰ A simple and intuitive scheduling library in Go. |
| [leprechaun](https://github.com/kilgaloon/leprechaun) | 24 | 4 | 2018-04-08 | 5 days ago | You had one job, or more then one, which can be done in steps |

## JSON
        
*Libraries for working with JSON.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gjson](https://github.com/tidwall/gjson) | 4095 | 112 | 2016-08-11 | 2 weeks ago | Get JSON values quickly  - JSON Parser for Go |
| [gojson](https://github.com/ChimeraCoder/gojson) | 1921 | 45 | 2012-12-28 | 2 weeks ago | Automatically generate Go (golang) struct definitions from example JSON |
| [gojq](https://github.com/elgs/gojq) | 132 | 5 | 2015-12-30 | 9 months ago | JSON query in Golang |
| [kazaam](https://github.com/qntfy/kazaam) | 106 | 15 | 2016-07-19 | 6 months ago | API for arbitrary transformation of JSON documents. |
| [kazaam](https://github.com/qntfy/kazaam) | 106 | 15 | 2016-07-19 | 6 months ago | API for arbitrary transformation of JSON documents. |
| [kazaam](https://github.com/qntfy/kazaam) | 106 | 15 | 2016-07-19 | 6 months ago | API for arbitrary transformation of JSON documents. |
| [jsongo](https://github.com/ricardolonga/jsongo) | 85 | 1 | 2015-08-08 | 2 years ago | Fluent API to make it easier to create Json objects. |
| [jsonf](https://github.com/miolini/jsonf) | 53 | 3 | 2015-05-25 | 2 years ago | Console JSON formatter with query feature |
| [jaydiff](https://github.com/yazgazan/jaydiff) | 33 | 1 | 2017-04-25 | 2 months ago | A JSON diff utility |
| [mp](https://github.com/sanbornm/mp) | 29 | 2 | 2014-06-16 | 2 years ago | Simple Email Parser |
| [go-respond](https://github.com/nicklaw5/go-respond) | 18 | 1 | 2017-03-13 | 4 months ago | A Go package for handling common HTTP JSON responses. |
| [jsonhal](https://github.com/RichardKnop/jsonhal) | 8 | 2 | 2016-01-15 | 4 months ago | A simple Go package to make custom structs marshal into HAL compatible JSON responses. |
| [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) | 5 | 1 | 2016-07-08 | 2 years ago | Go bindings based on the JSON API errors reference |

## Logging
        
*Libraries for generating and working with log files.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [logrus](https://github.com/sirupsen/logrus) | 10007 | 267 | 2013-10-17 | 18 hours ago | Structured logger for Go. |
| [logrus](https://github.com/sirupsen/logrus) | 9997 | 267 | 2013-10-17 | 18 hours ago | Structured logger for Go. |
| [logrus](https://github.com/sirupsen/logrus) | 9986 | 267 | 2013-10-17 | 2 days ago | Structured logger for Go. |
| [zap](https://github.com/uber-go/zap) | 6151 | 199 | 2016-02-19 | 10 hours ago | Blazing fast, structured, leveled logging in Go. |
| [go-spew](https://github.com/davecgh/go-spew) | 2979 | 59 | 2013-01-09 | 1 month ago | Implements a deep pretty printer for Go data structures to aid in debugging |
| [glog](https://github.com/golang/glog) | 2146 | 84 | 2013-07-16 | 3 months ago | Leveled execution logs for Go |
| [zerolog](https://github.com/rs/zerolog) | 1770 | 47 | 2017-05-12 | 13 hours ago | Zero Allocation JSON Logger |
| [tail](https://github.com/hpcloud/tail) | 1360 | 101 | 2013-02-05 | 5 days ago | Go package for reading from continously updated files (tail -f) |
| [seelog](https://github.com/cihub/seelog) | 1271 | 89 | 2011-11-17 | 1 day ago | Seelog is a native Go logging library that provides flexible asynchronous dispatching, filtering, and formatting. |
| [lumberjack](https://github.com/natefinch/lumberjack) | 1188 | 37 | 2014-06-14 | 3 days ago | lumberjack is a log rolling package for Go |
| [log15](https://github.com/inconshreveable/log15) | 843 | 23 | 2014-05-20 | 4 months ago | Structured, composable logging for Go |
| [log](https://github.com/apex/log) | 676 | 35 | 2015-12-22 | 4 months ago | Structured logging package for Go. |
| [logxi](https://github.com/mgutz/logxi) | 328 | 10 | 2015-03-02 | 1 year ago | A 12-factor app logger built for performance and happy development |
| [onelog](https://github.com/francoispqt/onelog) | 304 | 9 | 2018-05-06 | 1 day ago | Dead simple, super fast, zero allocation and modular logger for Golang |
| [log](https://github.com/go-playground/log) | 257 | 10 | 2016-02-08 | 2 weeks ago | :green_book: Simple, configurable and scalable Structured Logging for Go. |
| [logutils](https://github.com/hashicorp/logutils) | 239 | 88 | 2013-10-09 | 6 months ago | Utilities for slightly better logging in Go (Golang). |
| [go-logger](https://github.com/apsdehal/go-logger) | 212 | 8 | 2014-09-26 | 5 months ago | Simple logger for Go programs. Allows custom formats for messages. |
| [logger](https://github.com/azer/logger) | 130 | 4 | 2014-09-30 | 5 months ago | Minimalistic logging library for Go. |
| [xlog](https://github.com/rs/xlog) | 126 | 7 | 2015-10-22 | 11 months ago | xlog is a logger for net/context aware HTTP applications |
| [ozzo-log](https://github.com/go-ozzo/ozzo-log) | 102 | 11 | 2015-10-23 | 11 months ago | A Go (golang) package providing high-performance asynchronous logging, message filtering by severity and category, and multiple message targets. |
| [logvoyage](https://github.com/firstrow/logvoyage) | 83 | 5 | 2015-03-29 | 1 year ago | LogVoyage - logging SaaS written in GoLang |
| [log](https://github.com/alexcesaro/log) | 44 | 4 | 2014-04-19 | 3 years ago | Logging packages for Go |
| [gologger](https://github.com/sadlil/gologger) | 38 | 6 | 2015-09-02 | 1 year ago | The Simplest and worst logging library ever written |
| [go-log](https://github.com/ian-kent/go-log) | 33 | 2 | 2014-05-02 | 11 months ago | A logger, for Go |
| [logex](https://github.com/chzyer/logex) | 32 | 7 | 2014-10-10 | 1 year ago | An golang log lib, supports tracking and level, wrap by standard log lib |
| [glg](https://github.com/kpango/glg) | 32 | 3 | 2017-06-21 | 1 day ago | Simple and fast lockfree logging library for golang |
| [logrusly](https://github.com/sebest/logrusly) | 24 | 5 | 2014-09-12 | 11 months ago | Loggly Hooks for GO Logrus logger |
| [gone](https://github.com/One-com/gone) | 24 | 7 | 2016-09-05 | 3 weeks ago | Golang packages for writing small daemons and servers. |
| [go-log](https://github.com/siddontang/go-log) | 22 | 3 | 2014-05-18 | 1 week ago | a golang log lib supports level and multi handlers |
| [log](https://github.com/teris-io/log) | 21 | 1 | 2017-10-29 | 1 year ago | Structured log interface |
| [go-cronowriter](https://github.com/utahta/go-cronowriter) | 17 | 1 | 2017-02-04 | 23 hours ago | Time based rotating file writer |
| [journald](https://github.com/ssgreg/journald) | 17 | 2 | 2017-08-23 | 2 months ago | Go implementation of systemd Journal's native API for logging |
| [mlog](https://github.com/jbrodriguez/mlog) | 16 | 2 | 2014-10-20 | 7 months ago | A simple logging module for go, with a rotating file feature and console logging. |
| [distillog](https://github.com/amoghe/distillog) | 15 | 1 | 2015-10-13 | 7 months ago | Logging, distilled |
| [gomol](https://github.com/aphistic/gomol) | 12 | 2 | 2015-08-30 | 9 months ago | Gomol is a library for structured, multiple-output logging for Go with extensible logging outputs |
| [go-log](https://github.com/subchen/go-log) | 8 | 1 | 2017-05-07 | 9 months ago | Simple and configurable Logging in Go, with level, formatters and writers |
| [logdump](https://github.com/ewwwwwqm/logdump) | 7 | 2 | 2017-01-13 | 11 months ago | Package for multi-level logging |
| [xlog](https://github.com/xfxdev/xlog) | 7 | 1 | 2016-05-06 | 1 month ago | plugin architecture and flexible log system for golang |
| [logmatic](https://github.com/borderstech/logmatic) | 5 | 1 | 2018-11-07 | 1 month ago | Colorized logger for Golang with dynamic log level configuration |
| [glo](https://github.com/lajosbencz/glo) | 5 | 1 | 2019-01-20 | 1 month ago | Logging library for Golang |
| [logo](https://github.com/mbndr/logo) | 4 | 1 | 2017-02-08 | 1 year ago | Golang logger to different configurable writers. |

## Machine Learning
        
*Libraries for Machine Learning.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [golearn](https://github.com/sjwhitworth/golearn) | 6311 | 429 | 2013-12-26 | 3 weeks ago | Machine Learning for Go |
| [gorgonia](https://github.com/gorgonia/gorgonia) | 2447 | 162 | 2016-09-15 | 1 day ago | graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. |
| [gorgonia](https://github.com/gorgonia/gorgonia) | 2446 | 162 | 2016-09-15 | 1 day ago | graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. |
| [gorgonia](https://github.com/gorgonia/gorgonia) | 2445 | 162 | 2016-09-15 | 1 day ago | graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. |
| [tfgo](https://github.com/galeone/tfgo) | 1072 | 48 | 2017-05-23 | 2 weeks ago | Tensorflow + Go, the gopher way |
| [goml](https://github.com/cdipaolo/goml) | 963 | 73 | 2015-06-27 | 2 years ago | On-line Machine Learning in Go (and so much more) |
| [gosseract](https://github.com/otiai10/gosseract) | 745 | 31 | 2013-10-11 | 1 week ago | Go package for OCR (Optical Character Recognition), by using Tesseract C++ library |
| [CloudForest](https://github.com/ryanbressler/CloudForest) | 628 | 43 | 2012-10-23 | 3 months ago | Ensembles of decision trees in go/golang. |
| [bayesian](https://github.com/jbrukh/bayesian) | 596 | 30 | 2011-11-23 | 2 weeks ago | Naive Bayesian Classification for Golang. |
| [eaopt](https://github.com/MaxHalford/eaopt) | 579 | 29 | 2016-01-31 | 1 week ago | :four_leaf_clover: Evolutionary optimization library for Go (genetic algorithm, partical swarm optimization, differential evolution) |
| [gorse](https://github.com/zhenghaoz/gorse) | 344 | 16 | 2018-08-14 | 1 hour ago | A High Performance Recommender System Engine over SQL Database based on Collaborative Filtering written in Go (in progress) |
| [gobrain](https://github.com/goml/gobrain) | 337 | 23 | 2014-04-29 | 2 months ago | Neural Networks written in go |
| [regommend](https://github.com/muesli/regommend) | 233 | 14 | 2014-02-06 | 3 months ago | Recommendation engine for Go |
| [go-deep](https://github.com/patrikeh/go-deep) | 201 | 11 | 2017-12-09 | 5 months ago | Artificial Neural Network |
| [go-galib](https://github.com/thoj/go-galib) | 164 | 14 | 2009-11-30 | 3 years ago | Genetic Algorithms library written in Go / golang |
| [goRecommend](https://github.com/timkaye11/goRecommend) | 136 | 10 | 2014-07-16 | 4 years ago | Collaborative Filtering (CF) Algorithms in Go!  |
| [shield](https://github.com/eaigner/shield) | 119 | 10 | 2013-04-11 | 3 years ago | Bayesian text classifier with flexible tokenizers and storage backends for Go |
| [go-fann](https://github.com/white-pony/go-fann) | 98 | 8 | 2011-03-11 | 4 years ago | Go bindings for FANN, library for artificial neural networks |
| [goga](https://github.com/tomcraven/goga) | 72 | 8 | 2015-10-20 | 2 years ago | Golang Genetic Algorithm |
| [neural-go](https://github.com/schuyler/neural-go) | 60 | 2 | 2011-10-17 | 4 years ago | A multilayer perceptron network implemented in Go, with training via backpropagation. |
| [libsvm](https://github.com/datastream/libsvm) | 59 | 10 | 2012-07-31 | 2 years ago | libsvm go version |
| [go-pr](https://github.com/daviddengcn/go-pr) | 56 | 6 | 2013-06-07 | 5 years ago | Pattern recognition package in Go lang. |
| [neat](https://github.com/jinyeom/neat) | 51 | 11 | 2016-11-17 | 8 months ago | NEAT (NeuroEvolution of Augmenting Topologies) implemented in Go |
| [golinear](https://github.com/danieldk/golinear) | 38 | 5 | 2013-04-05 | 6 months ago | liblinear bindings for Go |
| [goscore](https://github.com/asafschers/goscore) | 27 | 2 | 2017-08-19 | 6 months ago | Go Scoring API for PMML |
| [fonet](https://github.com/Fontinalis/fonet) | 26 | 3 | 2017-10-03 | 10 months ago | fonet is a deep neural network package for Go. |
| [godist](https://github.com/e-dard/godist) | 22 | 2 | 2014-09-05 | 3 years ago | Probability distributions and associated methods in Go |
| [Varis](https://github.com/Xamber/Varis) | 21 | 5 | 2017-10-10 | 7 months ago | Golang Neural Network  |
| [go-cluster](https://github.com/e-XpertSolutions/go-cluster) | 18 | 5 | 2017-10-04 | 7 months ago | k-modes and k-prototypes clustering algorithms implementation in Go |
| [probab](https://github.com/ThePaw/probab) | 9 | 2 | 2015-09-14 | 3 years ago | Automatically exported from code.google.com/p/probab |
| [gomind](https://github.com/surenderthakran/gomind) | 5 | 2 | 2017-10-19 | 7 months ago | A simplistic Neural Network Library in Go |
| [evoli](https://github.com/khezen/evoli) | 5 | 3 | 2015-06-12 | 1 month ago | Genetic Algorithm and Particle Swarm Optimization |
| [mlgo](https://github.com/NullHypothesis/mlgo) | 4 | 2 | 2015-12-07 | 3 years ago | Automatically exported from code.google.com/p/mlgo |

## Messaging
        
*Libraries that implement messaging systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [sarama](https://github.com/Shopify/sarama) | 3945 | 361 | 2013-07-06 | 7 hours ago | Sarama is a Go library for Apache Kafka 0.8, and up. |
| [centrifugo](https://github.com/centrifugal/centrifugo) | 3353 | 174 | 2015-04-01 | 5 days ago | Language-agnostic real-time messaging server (Websocket and SockJS) |
| [gorush](https://github.com/appleboy/gorush) | 3278 | 151 | 2016-03-22 | 5 days ago | A push notification server written in Go (Golang). |
| [machinery](https://github.com/RichardKnop/machinery) | 2949 | 131 | 2015-04-06 | 15 hours ago | Machinery is an asynchronous task queue/job queue based on distributed message passing. |
| [go-socket.io](https://github.com/googollee/go-socket.io) | 2628 | 122 | 2013-07-13 | 1 week ago | socket.io library for golang, a realtime application framework. |
| [go-nats](https://github.com/nats-io/go-nats) | 2130 | 159 | 2012-08-15 | 1 day ago | Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library. |
| [go-nats](https://github.com/nats-io/go-nats) | 2130 | 158 | 2012-08-15 | 1 day ago | Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library. |
| [go-nats](https://github.com/nats-io/go-nats) | 2130 | 159 | 2012-08-15 | 7 hours ago | Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library. |
| [apns2](https://github.com/sideshow/apns2) | 1933 | 72 | 2016-01-05 | 23 hours ago | ⚡ HTTP/2 Apple Push Notification Service (APNs) push provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps, using the APNs HTTP/2 protocol. |
| [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) | 1771 | 239 | 2013-12-27 | 1 year ago | Golang push server cluster |
| [benthos](https://github.com/Jeffail/benthos) | 1727 | 54 | 2016-03-22 | 10 hours ago | A stream processor for dull stuff written in Go |
| [mangos](https://github.com/nanomsg/mangos) | 1521 | 89 | 2014-10-25 | 20 hours ago | Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability. |
| [mangos](https://github.com/nanomsg/mangos) | 1521 | 89 | 2014-10-25 | 20 hours ago | Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability. |
| [mangos](https://github.com/nanomsg/mangos) | 1521 | 89 | 2014-10-25 | 1 day ago | Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability. |
| [melody](https://github.com/olahol/melody) | 1329 | 43 | 2015-05-14 | 4 months ago | :notes: Minimalist websocket framework for Go |
| [go-nsq](https://github.com/nsqio/go-nsq) | 1321 | 61 | 2013-08-29 | 2 weeks ago | The official Go package for NSQ |
| [uniqush-push](https://github.com/uniqush/uniqush-push) | 1056 | 71 | 2011-08-29 | 4 months ago | Uniqush is a free and open source software system which provides a unified push service for server side notification to apps on mobile devices. |
| [mercure](https://github.com/dunglas/mercure) | 896 | 42 | 2018-07-14 | 3 days ago | Server-sent live updates: protocol and reference implementation |
| [zmq4](https://github.com/pebbe/zmq4) | 729 | 47 | 2013-10-18 | 1 month ago | A Go interface to ZeroMQ version 4 |
| [gollum](https://github.com/trivago/gollum) | 653 | 35 | 2015-06-21 | 3 weeks ago | An n:m message multiplexer written in Go |
| [EventBus](https://github.com/asaskevich/EventBus) | 481 | 27 | 2014-12-20 | 1 month ago | [Go] Lightweight eventbus with async compatibility for Go |
| [golongpoll](https://github.com/jcuga/golongpoll) | 397 | 21 | 2015-11-02 | 2 weeks ago | golang long polling library.  Makes web pub-sub easy via HTTP long-poll server :smiley: :coffee: :computer: |
| [dbus](https://github.com/godbus/dbus) | 315 | 15 | 2014-03-28 | 2 hours ago | Native Go bindings for D-Bus |
| [glue](https://github.com/desertbit/glue) | 299 | 13 | 2015-06-07 | 1 year ago | Glue - Robust Go and Javascript Socket Library (Alternative to Socket.io) |
| [emitter](https://github.com/olebedev/emitter) | 283 | 10 | 2015-11-11 | 1 month ago | Emits events in Go way, with wildcard, predicates, cancellation possibilities and many other good wins |
| [pubsub](https://github.com/cskr/pubsub) | 245 | 8 | 2012-04-01 | 3 weeks ago | Simple pubsub package for go. |
| [pubsub](https://github.com/cskr/pubsub) | 244 | 8 | 2012-04-01 | 3 weeks ago | Simple pubsub package for go. |
| [pubsub](https://github.com/cskr/pubsub) | 243 | 8 | 2012-04-01 | 3 weeks ago | Simple pubsub package for go. |
| [guble](https://github.com/smancke/guble) | 132 | 13 | 2015-11-16 | 1 year ago | websocket based messaging server written in golang |
| [oplog](https://github.com/dailymotion/oplog) | 94 | 97 | 2014-11-06 | 3 years ago | A generic oplog/replication system for microservices |
| [drone-line](https://github.com/appleboy/drone-line) | 57 | 5 | 2016-09-13 | 3 months ago | Sending line notifications using a binary, docker or Drone CI. |
| [rabbus](https://github.com/rafaeljesus/rabbus) | 55 | 8 | 2017-05-07 | 1 month ago | A tiny wrapper over amqp exchanges and queues 🚌 ✨ |
| [rabtap](https://github.com/jandelgado/rabtap) | 54 | 6 | 2017-11-11 | 3 days ago | RabbitMQ wire tap and swiss army knife |
| [RapidMQ](https://github.com/sybrexsys/RapidMQ) | 52 | 3 | 2016-10-05 | 1 year ago | RapidMQ is a pure, extremely productive, lightweight and reliable library for managing of the local messages queue |
| [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) | 46 | 7 | 2017-01-16 | 1 year ago | A tiny wrapper around NSQ topic and channel :rocket: |
| [go-notify](https://github.com/TheCreeper/go-notify) | 44 | 2 | 2015-03-02 | 3 weeks ago | Package notify provides an implementation of the Gnome DBus Notifications Specification. |
| [goose](https://github.com/ian-kent/goose) | 35 | 1 | 2014-12-21 | 4 years ago | Server-Sent Events in Go |
| [message-bus](https://github.com/vardius/message-bus) | 32 | 3 | 2017-10-04 | 4 weeks ago | Go simple async message bus |
| [event](https://github.com/agoalofalife/event) | 20 | 2 | 2017-07-02 | 1 year ago | The implementation of the pattern observer |
| [hub](https://github.com/leandro-lugaresi/hub) | 19 | 1 | 2018-04-14 | 10 months ago | :incoming_envelope: A fast Message/Event Hub using publish/subscribe pattern with support for topics like* rabbitMQ exchanges for Go applications |
| [go-vitotrol](https://github.com/maxatome/go-vitotrol) | 11 | 6 | 2016-11-04 | 4 months ago | golang client library to Viessmann Vitotrol web service |
| [gaurun-client](https://github.com/osamingo/gaurun-client) | 7 | 1 | 2017-06-29 | 7 months ago | Gaurun Client written in Go |
| [jazz](https://github.com/socifi/jazz) | 4 | 4 | 2018-10-22 | 4 weeks ago | Abstraction layer for simple rabbitMQ connection, messaging and administration |

## Microsoft Office
        

### Microsoft Excel
        
*Libraries for working with Microsoft Excel.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [excelize](https://github.com/360EntSecGroup-Skylar/excelize) | 3360 | 127 | 2016-08-29 | 1 week ago | Golang library for reading and writing Microsoft Excel™ (XLSX) files. |
| [xlsx](https://github.com/tealeg/xlsx) | 3017 | 175 | 2011-06-28 | 1 week ago | Google Go (golang) library for reading and writing XLSX files.  You should probably also checkout: https://github.com/360EntSecGroup-Skylar/excelize |
| [xlsx](https://github.com/plandem/xlsx) | 35 | 3 | 2017-08-27 | 1 month ago | Fast and safe way to read/update your existing Excel xlsx files |
| [go-excel](https://github.com/szyhf/go-excel) | 32 | 2 | 2017-09-03 | 2 months ago | A simple and light excel file reader to read a standard excel as a table faster | 一个轻量级的Excel数据读取库，用一种更`关系数据库`的方式解析Excel。 |
| [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) | 11 | 1 | 2017-03-13 | 7 months ago | Golang bindings for libxlsxwriter for writing XLSX files |

## Miscellaneous
        

### Dependency Injection
        
*Libraries for working with dependency injection.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [dig](https://github.com/uber-go/dig) | 615 | 35 | 2017-03-22 | 1 week ago | A reflection based dependency injection toolkit for Go. |
| [alice](https://github.com/magic003/alice) | 31 | 1 | 2017-04-09 | 1 year ago | An additive dependency injection container for Golang. |
| [wire](https://github.com/Fs02/wire) | 13 | 1 | 2018-07-05 | 1 month ago | Strict Runtime Dependency Injection for Golang |

### Strings
        
*These libraries were placed here because none of the other categories seemed to fit.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gopsutil](https://github.com/shirou/gopsutil) | 3430 | 172 | 2014-04-18 | 12 hours ago | Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc). |
| [archiver](https://github.com/mholt/archiver) | 2093 | 43 | 2016-04-09 | 3 days ago | Easily create and extract .zip, .tar, .tar.gz, .tar.bz2, .tar.xz, .tar.lz4, .tar.sz, and .rar (extract-only) files with Go |
| [gosms](https://github.com/haxpax/gosms) | 1199 | 54 | 2015-01-24 | 1 year ago | Your own local SMS gateway in Go that can be used to send SMS. |
| [go-resiliency](https://github.com/eapache/go-resiliency) | 764 | 29 | 2014-11-29 | 3 weeks ago | Resiliency patterns for golang. |
| [go-commons-pool](https://github.com/jolestar/go-commons-pool) | 593 | 43 | 2015-12-28 | 6 months ago | a generic object pool for golang |
| [xstrings](https://github.com/huandu/xstrings) | 550 | 24 | 2015-01-06 | 5 months ago | Package xstrings: A collection of useful string functions in Go. |
| [base64Captcha](https://github.com/mojocn/base64Captcha) | 486 | 33 | 2017-12-12 | 3 months ago | golang base64-captcha supports digits, numbers,alphabet, arithmetic, audio and digit-alphabet captcha |
| [shortid](https://github.com/teris-io/shortid) | 402 | 6 | 2016-01-04 | 1 year ago | Distributed generation of super short, unique, non-sequential, URL friendly IDs. |
| [llvm](https://github.com/llir/llvm) | 351 | 29 | 2014-09-19 | 2 days ago | Library for interacting with LLVM IR in pure Go. |
| [health](https://github.com/dimiro1/health) | 343 | 6 | 2016-03-09 | 7 months ago | Easy to use, extensible health check library. |
| [go-conv](https://github.com/cstockton/go-conv) | 336 | 10 | 2016-10-11 | 1 year ago | Fast conversions across various Go types with a simple API. |
| [gofakeit](https://github.com/brianvoe/gofakeit) | 242 | 4 | 2015-04-24 | 1 month ago | Random data generator written in go. |
| [banner](https://github.com/dimiro1/banner) | 205 | 4 | 2016-03-26 | 2 years ago | An easy way to add useful startup banners into your Go applications |
| [gountries](https://github.com/pariz/gountries) | 198 | 6 | 2016-01-13 | 3 months ago | Package that exposes country and subdivision data. |
| [battery](https://github.com/distatus/battery) | 122 | 2 | 2016-03-13 | 5 days ago | cross-platform, normalized battery information library |
| [antch](https://github.com/antchfx/antch) | 121 | 10 | 2017-09-28 | 7 months ago | Antch, a fast, powerful and extensible web crawling & scraping framework for Go |
| [stats](https://github.com/go-playground/stats) | 115 | 3 | 2015-09-15 | 2 years ago | Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... |
| [ffmt](https://github.com/go-ffmt/ffmt) | 108 | 4 | 2015-02-14 | 3 months ago | Golang beautify data display for Humans |
| [lk](https://github.com/hyperboloide/lk) | 89 | 5 | 2016-07-15 | 1 month ago | A simple licensing library for golang. |
| [bitio](https://github.com/icza/bitio) | 80 | 5 | 2016-05-31 | 1 year ago | Highly optimized bit-level Reader and Writer for Go. |
| [turtle](https://github.com/hackebrot/turtle) | 69 | 1 | 2017-09-09 | 1 year ago | Emojis for Go. |
| [healthcheck](https://github.com/etherlabsio/healthcheck) | 62 | 4 | 2017-08-18 | 1 year ago | An opinionated and concurrent health-check HTTP handler for RESTful services. |
| [gommit](https://github.com/antham/gommit) | 61 | 3 | 2016-08-30 | 8 hours ago | Analyze git commit messages to ensure they follow defined patterns. |
| [go-unarr](https://github.com/gen2brain/go-unarr) | 56 | 5 | 2015-11-01 | 1 month ago | Decompression library for RAR, TAR, ZIP and 7z archives. |
| [indigo](https://github.com/osamingo/indigo) | 47 | 1 | 2016-08-31 | 4 months ago | Distributed unique ID generator of using Sonyflake and encoded by Base58. |
| [morse](https://github.com/alwindoss/morse) | 45 | 2 | 2018-08-15 | 1 week ago | Library to convert to and from morse code. |
| [strutil](https://github.com/ozgio/strutil) | 44 | 1 | 2018-08-16 | 6 months ago | String utilities for Go |
| [xkg](https://github.com/go-xkg/xkg) | 37 | 1 | 2015-01-05 | 4 years ago | X Keyboard Grabber. |
| [captcha](https://github.com/steambap/captcha) | 33 | 4 | 2017-09-12 | 5 days ago | :sunglasses:Package captcha provides an easy to use, unopinionated API for captcha generation |
| [persian](https://github.com/mavihq/persian) | 27 | 2 | 2017-10-17 | 7 months ago | Some utilities for Persian language in go. |
| [browscap_go](https://github.com/digitalcrab/browscap_go) | 27 | 6 | 2014-09-18 | 1 month ago | GoLang Library for Browser Capabilities Project |
| [datacounter](https://github.com/miolini/datacounter) | 26 | 1 | 2015-10-15 | 10 months ago | Golang counters for readers/writers |
| [autoflags](https://github.com/artyom/autoflags) | 23 | 2 | 2014-05-16 | 1 month ago | Populate go command line app flags from config struct |
| [pdfgen](https://github.com/hyperboloide/pdfgen) | 21 | 3 | 2015-12-01 | 1 year ago | HTTP service to generate PDF from Json requests. |
| [ghorg](https://github.com/gabrie30/ghorg) | 17 | 2 | 2018-03-29 | 2 months ago | Quickly clone an entire GitHub Org into one directory  |
| [xdg](https://github.com/rkoesters/xdg) | 13 | 1 | 2013-12-15 | 3 months ago | FreeDesktop.org (xdg) Specs implemented in Go. |
| [gosh](https://github.com/osamingo/gosh) | 13 | 0 | 2018-05-25 | 1 month ago | Provide Go Statistics Handler, Struct, Measure Method. |
| [url-shortener](https://github.com/pantrif/url-shortener) | 12 | 1 | 2018-06-04 | 8 months ago | A modern, powerful, and robust URL shortener microservice with mysql support. |
| [anagent](https://github.com/mudler/anagent) | 9 | 2 | 2017-12-30 | 6 months ago | Minimalistic, pluggable Golang evloop/timer handler with dependency-injection |
| [hostutils](https://github.com/Wing924/hostutils) | 8 | 1 | 2017-09-26 | 1 month ago | A golang library for packing and unpacking FQDNs list. |
| [avgRating](https://github.com/kirillDanshin/avgRating) | 8 | 1 | 2017-08-06 | 1 year ago | Calculate average score and rating based on Wilson Score Equation |
| [sandid](https://github.com/aofei/sandid) | 8 | 1 | 2018-06-12 | 1 week ago | Every grain of sand on earth has its own ID. |
| [shellwords](https://github.com/Wing924/shellwords) | 5 | 1 | 2017-09-28 | 1 year ago | A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell. |
| [generators](https://github.com/azr/generators) | 3 | 1 | 2016-02-29 | 2 years ago | Generate boilerplate http input and output handling. |

## Natural Language Processing
        
*Libraries for working with human languages.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [prose](https://github.com/jdkato/prose) | 1924 | 46 | 2017-02-18 | 4 weeks ago | Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. |
| [when](https://github.com/olebedev/when) | 907 | 21 | 2016-12-27 | 1 day ago | Natural EN and RU language date/time parser with pluggable rules. |
| [go-i18n](https://github.com/nicksnyder/go-i18n) | 852 | 17 | 2012-01-15 | 1 month ago | Package and an accompanying tool to work with localized text. |
| [gse](https://github.com/go-ego/gse) | 793 | 40 | 2017-06-23 | 1 day ago | Go efficient text segmentation; support english, chinese, japanese and other. |
| [gojieba](https://github.com/yanyiwu/gojieba) | 719 | 59 | 2015-09-12 | 5 months ago | This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm. |
| [go-pinyin](https://github.com/mozillazg/go-pinyin) | 446 | 28 | 2014-11-09 | 6 months ago | CN Hanzi to Hanyu Pinyin converter. |
| [kagome](https://github.com/ikawaha/kagome) | 356 | 21 | 2014-06-26 | 2 weeks ago | JP morphological analyzer written in pure Go. |
| [nlp](https://github.com/shixzie/nlp) | 347 | 23 | 2017-01-25 | 1 year ago | Extract values from strings and fill your structs with nlp. |
| [whatlanggo](https://github.com/abadojack/whatlanggo) | 320 | 14 | 2017-02-21 | 1 day ago | Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc). |
| [sentences](https://github.com/neurosnap/sentences) | 248 | 12 | 2015-08-07 | 3 weeks ago | Sentence tokenizer:  converts text into a list of sentences. |
| [nlp](https://github.com/james-bowman/nlp) | 189 | 20 | 2017-03-15 | 3 days ago | Go Natural Language Processing library supporting LSA (Latent Semantic Analysis). |
| [go-nlp](https://github.com/nuance/go-nlp) | 78 | 8 | 2011-05-02 | 7 years ago | Utilities for working with discrete probability distributions and other tools useful for doing NLP work. |
| [gounidecode](https://github.com/fiam/gounidecode) | 62 | 3 | 2012-05-01 | 3 years ago | Unicode transliterator (also known as unidecode) for Go. |
| [textcat](https://github.com/pebbe/textcat) | 60 | 6 | 2012-09-21 | 7 months ago | Go package for n-gram based text categorization, with support for utf-8 and raw text. |
| [MMSEGO](https://github.com/awsong/MMSEGO) | 58 | 5 | 2012-04-18 | 6 years ago | This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm. |
| [getlang](https://github.com/rylans/getlang) | 52 | 3 | 2018-03-02 | 1 month ago | Fast natural language detection package. |
| [go-stem](https://github.com/agonopol/go-stem) | 51 | 2 | 2011-09-24 | 8 months ago | Implementation of the porter stemming algorithm. |
| [go-unidecode](https://github.com/mozillazg/go-unidecode) | 45 | 2 | 2016-07-08 | 2 years ago | ASCII transliterations of Unicode text. |
| [stemmer](https://github.com/dchest/stemmer) | 43 | 4 | 2011-03-21 | 2 years ago | Stemmer packages for Go programming language. Includes English and German stemmers. |
| [RAKE.Go](https://github.com/afjoseph/RAKE.Go) | 39 | 5 | 2016-12-17 | 2 months ago | Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE). |
| [porter2](https://github.com/zentures/porter2) | 33 | 3 | 2015-01-21 | 3 years ago | Really fast Porter 2 stemmer. |
| [go2vec](https://github.com/danieldk/go2vec) | 30 | 6 | 2015-01-27 | 6 months ago | Reader and utility functions for word2vec embeddings. |
| [paicehusk](https://github.com/rookii/paicehusk) | 24 | 3 | 2012-09-30 | 5 years ago | Golang implementation of the Paice/Husk Stemming Algorithm. |
| [snowball](https://github.com/goodsign/snowball) | 23 | 0 | 2012-12-11 | 1 year ago | Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/). |
| [mystem](https://github.com/dveselov/mystem) | 21 | 3 | 2016-08-30 | 2 years ago | CGo bindings to Yandex.Mystem - russian morphology analyzer. |
| [petrovich](https://github.com/striker2000/petrovich) | 20 | 1 | 2016-12-27 | 2 months ago | Petrovich is the library which inflects Russian names to given grammatical case. |
| [icu](https://github.com/goodsign/icu) | 18 | 0 | 2012-12-11 | 1 year ago | Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1. |
| [golibstemmer](https://github.com/rjohnsondev/golibstemmer) | 15 | 1 | 2012-08-07 | 4 years ago | Go bindings for the snowball libstemmer library including porter 2. |
| [libtextcat](https://github.com/goodsign/libtextcat) | 10 | 1 | 2012-12-11 | 6 years ago | Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2. |
| [shamoji](https://github.com/osamingo/shamoji) | 10 | 2 | 2017-07-23 | 1 month ago | The shamoji is word filtering package written in Go. |
| [porter](https://github.com/a2800276/porter) | 8 | 1 | 2013-09-17 | 5 years ago | This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm. |
| [gotokenizer](https://github.com/xujiajun/gotokenizer) | 4 | 1 | 2018-10-11 | 1 month ago | A tokenizer based on the dictionary and Bigram language models for Golang. (Now only support chinese segmentation) |
| [go-eco](https://github.com/ThePaw/go-eco) | 4 | 0 | 2015-09-14 | 3 years ago | Similarity, dissimilarity and distance matrices; diversity, equitability and inequality measures; species richness estimators; coenocline models. |

## Networking
        
*Libraries for working with various layers of the network.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [kcptun](https://github.com/xtaci/kcptun) | 9770 | 581 | 2016-02-26 | 2 weeks ago | Extremely simple & fast udp tunnel based on KCP protocol. |
| [fasthttp](https://github.com/valyala/fasthttp) | 8156 | 329 | 2015-10-19 | 2 days ago | Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http. |
| [dns](https://github.com/miekg/dns) | 3459 | 155 | 2010-08-04 | 1 day ago | Go library for working with DNS. |
| [httplab](https://github.com/gchaincl/httplab) | 3306 | 66 | 2017-02-09 | 4 months ago | HTTPLabs let you inspect HTTP requests and forge responses. |
| [gopacket](https://github.com/google/gopacket) | 2526 | 129 | 2015-03-17 | 1 week ago | Go library for packet processing with libpcap bindings. |
| [quic-go](https://github.com/lucas-clemente/quic-go) | 2444 | 141 | 2016-04-07 | 4 hours ago | An implementation of the QUIC protocol in pure Go. |
| [gobgp](https://github.com/osrg/gobgp) | 1557 | 123 | 2014-09-14 | 2 days ago | BGP implemented in the Go Programming Language. |
| [webrtc](https://github.com/pions/webrtc) | 1125 | 64 | 2018-05-19 | 2 hours ago | A pure Go implementation of the WebRTC API. |
| [ssh](https://github.com/gliderlabs/ssh) | 973 | 38 | 2016-10-04 | 1 week ago | Higher-level API for building SSH servers (wraps crypto/ssh). |
| [water](https://github.com/songgao/water) | 751 | 34 | 2013-03-26 | 3 days ago | Simple TUN/TAP library. |
| [sftp](https://github.com/pkg/sftp) | 654 | 47 | 2013-11-05 | 3 weeks ago | Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt. |
| [fortio](https://github.com/fortio/fortio) | 646 | 28 | 2017-10-10 | 1 day ago | Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC. |
| [go-getter](https://github.com/hashicorp/go-getter) | 626 | 89 | 2015-10-13 | 2 hours ago | Go library for downloading files or directories from various sources using a URL. |
| [nff-go](https://github.com/intel-go/nff-go) | 580 | 66 | 2017-03-30 | 3 days ago | Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF). |
| [mdns](https://github.com/hashicorp/mdns) | 496 | 92 | 2014-01-30 | 1 month ago | Simple mDNS (Multicast DNS) client/server library in Golang. |
| [lhttp](https://github.com/fanux/lhttp) | 473 | 55 | 2015-12-29 | 11 months ago | Powerful websocket framework, build your IM server more easily. |
| [grab](https://github.com/cavaliercoder/grab) | 464 | 14 | 2016-01-05 | 1 week ago | Go package for managing file downloads. |
| [ftp](https://github.com/jlaffaye/ftp) | 441 | 23 | 2011-05-07 | 4 days ago | Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959). |
| [gosnmp](https://github.com/soniah/gosnmp) | 398 | 37 | 2012-08-27 | 1 week ago | Native Go library for performing SNMP actions. |
| [gotcp](https://github.com/gansidui/gotcp) | 390 | 42 | 2014-04-13 | 1 year ago | Go package for quickly writing tcp applications. |
| [cidranger](https://github.com/yl2chen/cidranger) | 350 | 11 | 2017-08-21 | 1 month ago | Fast IP to CIDR lookup for Go. |
| [gopcap](https://github.com/akrennmair/gopcap) | 338 | 25 | 2009-11-19 | 2 years ago | Go wrapper for libpcap. |
| [peerdiscovery](https://github.com/schollz/peerdiscovery) | 335 | 13 | 2018-04-23 | 1 week ago | Pure Go library for cross-platform local peer discovery using UDP multicast. |
| [go-stun](https://github.com/ccding/go-stun) | 294 | 13 | 2013-08-18 | 7 months ago | Go implementation of the STUN client (RFC 3489 and RFC 5389). |
| [raw](https://github.com/mdlayher/raw) | 252 | 9 | 2015-07-07 | 1 day ago | Package raw enables reading and writing data at the device driver level for a network interface. |
| [tcp_server](https://github.com/firstrow/tcp_server) | 251 | 16 | 2014-10-14 | 6 months ago | Go library for building tcp servers faster. |
| [buffstreams](https://github.com/StabbyCutyou/buffstreams) | 217 | 13 | 2015-06-30 | 1 year ago | Streaming protocolbuffer data over TCP made easy. |
| [stun](https://github.com/gortc/stun) | 193 | 12 | 2016-04-25 | 2 weeks ago | Go implementation of RFC 5389 STUN protocol. |
| [winrm](https://github.com/masterzen/winrm) | 193 | 18 | 2013-12-31 | 1 week ago | Go WinRM client to remotely execute commands on Windows machines. |
| [ethernet](https://github.com/mdlayher/ethernet) | 168 | 11 | 2015-07-03 | 4 months ago | Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. |
| [utp](https://github.com/anacrolix/utp) | 147 | 15 | 2015-03-20 | 1 year ago | Go uTP micro transport protocol implementation. |
| [canopus](https://github.com/zubairhamed/canopus) | 128 | 15 | 2015-02-24 | 11 months ago | CoAP Client/Server implementation (RFC 7252). |
| [sslb](https://github.com/eduardonunesp/sslb) | 110 | 7 | 2015-10-19 | 1 year ago | It's a Super Simples Load Balancer, just a little project to achieve some kind of performance. |
| [jazigo](https://github.com/udhos/jazigo) | 108 | 11 | 2016-06-08 | 6 months ago | Jazigo is a tool written in Go for retrieving configuration for multiple network devices. |
| [gnxi](https://github.com/google/gnxi) | 81 | 20 | 2017-09-26 | 1 day ago | A collection of tools for Network Management that use the gNMI and gNOI protocols. |
| [xtcp](https://github.com/xfxdev/xtcp) | 68 | 10 | 2016-04-01 | 1 month ago | TCP Server Framework with simultaneous full duplex communication,graceful shutdown,custom protocol. |
| [ether](https://github.com/songgao/ether) | 58 | 2 | 2014-05-21 | 2 years ago | Cross-platform Go package for sending and receiving ethernet frames. |
| [dhcp6](https://github.com/mdlayher/dhcp6) | 55 | 4 | 2015-05-22 | 3 months ago | Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. |
| [portproxy](https://github.com/aybabtme/portproxy) | 39 | 0 | 2014-12-13 | 4 years ago | Simple TCP proxy which adds CORS support to API's which don't support it. |
| [linkio](https://github.com/ian-kent/linkio) | 39 | 2 | 2014-12-24 | 1 year ago | Network link speed simulation for Reader/Writer interfaces. |
| [gmqtt](https://github.com/DrmagicE/gmqtt) | 37 | 5 | 2018-09-16 | 2 months ago | Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1. |
| [graval](https://github.com/koofr/graval) | 23 | 7 | 2014-04-23 | 9 months ago | Experimental FTP server framework. |
| [publicip](https://github.com/polera/publicip) | 17 | 1 | 2016-12-29 | 2 years ago | Package publicip returns your public facing IPv4 address (internet egress). |
| [golibwireshark](https://github.com/sunwxg/golibwireshark) | 14 | 2 | 2015-11-16 | 1 year ago | Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data. |
| [packet](https://github.com/aerogo/packet) | 10 | 2 | 2017-10-29 | 2 months ago | Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed. |
| [goshark](https://github.com/sunwxg/goshark) | 8 | 1 | 2015-11-01 | 1 year ago | Package goshark use tshark to decode IP packet and create data struct to analyse packet. |
| [llb](https://github.com/kirillDanshin/llb) | 7 | 1 | 2016-02-21 | 2 years ago | It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response. |
| [tspool](https://github.com/two/tspool) | 5 | 0 | 2018-10-27 | 4 months ago | A TCP Library use worker pool to improve performance and protect your server. |

### HTTP Clients
        
*Libraries for making HTTP requests.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [grequests](https://github.com/levigross/grequests) | 1297 | 30 | 2015-06-12 | 2 weeks ago | A Go "clone" of the great and famous Requests library. |
| [heimdall](https://github.com/gojektech/heimdall) | 813 | 41 | 2018-01-19 | 2 weeks ago | An enchanced http client with retry and hystrix capabilities. |
| [sling](https://github.com/dghubble/sling) | 796 | 28 | 2015-04-02 | 4 hours ago | Sling is a Go HTTP client library for creating and sending API requests. |
| [gentleman](https://github.com/h2non/gentleman) | 612 | 16 | 2016-02-22 | 1 month ago | Full-featured plugin-driven HTTP client library. |
| [pester](https://github.com/sethgrid/pester) | 301 | 5 | 2015-05-20 | 1 month ago | Go HTTP client calls with retries, backoff, and concurrency. |
| [goreq](https://github.com/smallnest/goreq) | 92 | 7 | 2015-09-17 | 7 months ago | Enhanced simplified HTTP client based on gorequest. |
| [rq](https://github.com/ddo/rq) | 24 | 2 | 2017-12-26 | 4 months ago | A nicer interface for golang stdlib HTTP client. |

## OpenGL
        
*Libraries for using OpenGL in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [glfw](https://github.com/go-gl/glfw) | 597 | 35 | 2013-05-19 | 2 weeks ago | Go bindings for GLFW 3. |
| [gl](https://github.com/go-gl/gl) | 579 | 40 | 2015-02-22 | 2 months ago | Go bindings for OpenGL (generated via glow). |
| [mathgl](https://github.com/go-gl/mathgl) | 266 | 22 | 2013-02-13 | 4 days ago | Pure Go math package specialized for 3D math, with inspiration from GLM. |
| [gl](https://github.com/goxjs/gl) | 130 | 13 | 2015-05-18 | 4 months ago | Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android). |
| [glfw](https://github.com/goxjs/glfw) | 59 | 5 | 2014-12-28 | 4 months ago | Go cross-platform glfw library for creating an OpenGL context and receiving events. |

## ORM
        
*Libraries that implement Object-Relational Mapping or datamapping techniques.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gorm](https://github.com/jinzhu/gorm) | 12403 | 392 | 2013-10-25 | 3 days ago | The fantastic ORM library for Golang, aims to be developer friendly. |
| [xorm](https://github.com/go-xorm/xorm) | 4448 | 241 | 2013-05-09 | 2 weeks ago | Simple and powerful ORM for Go. |
| [gorp](https://github.com/go-gorp/gorp) | 2975 | 116 | 2012-01-05 | 3 months ago | Go Relational Persistence, ORM-ish library for Go. |
| [pg](https://github.com/go-pg/pg) | 2449 | 76 | 2013-04-24 | 53 minutes ago | PostgreSQL ORM with focus on PostgreSQL specific features and performance. |
| [sqlboiler](https://github.com/volatiletech/sqlboiler) | 1925 | 66 | 2016-02-21 | 1 week ago | ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema. |
| [db](https://github.com/upper/db) | 1660 | 59 | 2013-10-23 | 1 month ago | Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers. |
| [reform](https://github.com/go-reform/reform) | 737 | 31 | 2016-02-25 | 3 weeks ago | Better ORM for Go, based on non-empty interfaces and code generation. |
| [pop](https://github.com/gobuffalo/pop) | 584 | 24 | 2018-02-08 | 1 day ago | A Tasty Treat For All Your Database Needs |
| [qbs](https://github.com/coocood/qbs) | 526 | 45 | 2013-02-02 | 1 year ago | Stands for Query By Struct. A Go ORM. |
| [go-queryset](https://github.com/jirfag/go-queryset) | 400 | 19 | 2017-09-04 | 1 month ago | 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM. |
| [zoom](https://github.com/albrow/zoom) | 230 | 14 | 2013-07-17 | 8 months ago | Blazing-fast datastore and querying engine built on Redis. |
| [grimoire](https://github.com/Fs02/grimoire) | 104 | 3 | 2018-03-06 | 6 hours ago | Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3). |
| [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) | 99 | 8 | 2017-12-28 | 3 months ago | A flexible and powerful SQL string builder library plus a zero-config ORM. |
| [go-store](https://github.com/gosuri/go-store) | 92 | 10 | 2015-03-22 | 2 years ago | Simple and fast Redis backed key-value store library for Go. |
| [gomodel](https://github.com/cosiner/gomodel) | 62 | 5 | 2015-04-16 | 1 year ago | Lightweight, fast, orm-like library helps interactive with database. |
| [marlow](https://github.com/dadleyy/marlow) | 47 | 4 | 2017-09-15 | 5 months ago | Generated ORM from project structs for compile time safety assurances. |
| [lore](https://github.com/abrahambotros/lore) | 4 | 1 | 2017-04-29 | 1 year ago | Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go. |

## Package Management
        
*Unofficial libraries for package and dependency management.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [dep](https://github.com/golang/dep) | 11636 | 304 | 2016-10-07 | 14 hours ago | Go dependency tool. |
| [glide](https://github.com/Masterminds/glide) | 7499 | 197 | 2014-07-09 | 1 week ago | Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip. |
| [godep](https://github.com/tools/godep) | 5602 | 159 | 2013-05-01 | 10 months ago | dependency tool for go, godep helps build packages reproducibly by fixing their dependencies. |
| [govendor](https://github.com/kardianos/govendor) | 4353 | 106 | 2015-04-12 | 3 months ago | Go Package Manager. Go vendor tool that works with the standard vendor file. |
| [gopm](https://github.com/gpmgo/gopm) | 2146 | 77 | 2013-05-15 | 1 week ago | Go Package Manager. |
| [gom](https://github.com/mattn/gom) | 1346 | 40 | 2013-09-11 | 9 months ago | Go Manager - bundle for go. |
| [gpm](https://github.com/pote/gpm) | 1201 | 32 | 2013-09-05 | 1 year ago | Barebones dependency manager for Go. |
| [goop](https://github.com/petejkim/goop) | 773 | 37 | 2014-06-18 | 3 years ago | Simple dependency manager for Go (golang), inspired by Bundler. |
| [nut](https://github.com/jingweno/nut) | 248 | 9 | 2015-01-23 | 3 years ago | Vendor Go dependencies. |
| [johnny-deps](https://github.com/VividCortex/johnny-deps) | 213 | 23 | 2013-07-19 | 4 years ago | Minimal dependency version using Git. |
| [gigo](https://github.com/LyricalSecurity/gigo) | 196 | 5 | 2015-05-01 | 3 months ago | PIP-like dependency tool for golang, with support for private repositories and hashes. |
| [VenGO](https://github.com/DamnWidget/VenGO) | 115 | 8 | 2014-10-17 | 2 years ago | create and manage exportable isolated go virtual environments. |
| [mvn-golang](https://github.com/raydac/mvn-golang) | 76 | 7 | 2016-03-24 | 3 days ago | plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure. |
| [gop](https://github.com/lunny/gop) | 50 | 7 | 2017-02-18 | 2 weeks ago | Build and manage your Go applications out of GOPATH. |

## Query Language
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [graphql](https://github.com/graphql-go/graphql) | 4443 | 142 | 2015-07-19 | 1 week ago | Implementation of GraphQL for Go. |
| [graphql-go](https://github.com/graph-gophers/graphql-go) | 2437 | 91 | 2016-10-18 | 1 week ago | GraphQL server with a focus on ease of use. |
| [gojsonq](https://github.com/thedevsaddam/gojsonq) | 728 | 20 | 2018-05-20 | 1 week ago | A simple Go package to Query over JSON Data. |
| [jsonql](https://github.com/elgs/jsonql) | 193 | 8 | 2015-12-29 | 3 months ago | JSON query expression library in Golang. |
| [rql](https://github.com/a8m/rql) | 86 | 5 | 2018-06-06 | 1 day ago | Resource Query Language for REST API. |
| [graphql](https://github.com/tmc/graphql) | 50 | 10 | 2015-04-19 | 1 year ago | graphql parser + utilities. |
| [jsonslice](https://github.com/bhmj/jsonslice) | 16 | 1 | 2018-05-02 | 3 days ago | Jsonpath queries with advanced filters. |

## Resource Embedding
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [statik](https://github.com/rakyll/statik) | 1749 | 32 | 2014-02-04 | 19 hours ago | Embeds static files into a Go executable. |
| [packr](https://github.com/gobuffalo/packr) | 1684 | 37 | 2017-03-16 | 5 days ago | The simple and easy way to embed static files into Go binaries. |
| [go.rice](https://github.com/GeertJohan/go.rice) | 1513 | 43 | 2013-10-24 | 1 day ago | go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy. |
| [vfsgen](https://github.com/shurcooL/vfsgen) | 490 | 19 | 2015-05-18 | 2 weeks ago | Generates a vfsdata.go file that statically implements the given virtual filesystem. |
| [esc](https://github.com/mjibson/esc) | 407 | 12 | 2014-01-26 | 3 weeks ago | Embeds files into Go programs and provides http.FileSystem interfaces to them. |
| [fileb0x](https://github.com/UnnoTed/fileb0x) | 388 | 14 | 2016-01-24 | 1 month ago | Simple tool to embed files in go with focus on "customization" and ease to use. |
| [go-resources](https://github.com/omeid/go-resources) | 148 | 6 | 2015-02-21 | 1 year ago | Unfancy resources embedding with Go. |
| [statics](https://github.com/go-playground/statics) | 51 | 3 | 2015-10-07 | 2 years ago | Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks. |
| [templify](https://github.com/wlbr/templify) | 15 | 2 | 2016-05-23 | 1 month ago | Embed external template files into Go code to create single file binaries. |
| [go-embed](https://github.com/pyros2097/go-embed) | 12 | 1 | 2015-12-07 | 2 years ago | Generates go code to embed resource files into your library or executable. |

## Science and Data Analysis
        
*Libraries for scientific computing and data analyzing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gonum](https://github.com/gonum/gonum) | 2448 | 86 | 2017-03-25 | 1 hour ago | Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more. |
| [streamtools](https://github.com/nytlabs/streamtools) | 1311 | 77 | 2013-07-06 | 3 years ago | general purpose, graphical tool for dealing with streams of data. |
| [stats](https://github.com/montanaflynn/stats) | 1250 | 36 | 2014-12-16 | 1 month ago | Statistics package with common functions missing from the Golang standard library. |
| [gosl](https://github.com/cpmech/gosl) | 1232 | 62 | 2015-02-10 | 5 days ago | Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more. |
| [plot](https://github.com/gonum/plot) | 1069 | 50 | 2013-07-23 | 45 minutes ago | gonum/plot provides an API for building and drawing plots in Go. |
| [go-dsp](https://github.com/mjibson/go-dsp) | 604 | 28 | 2011-11-02 | 10 months ago | Digital Signal Processing for Go. |
| [goraph](https://github.com/gyuho/goraph) | 575 | 38 | 2014-02-27 | 1 year ago | Pure Go graph theory library(data structure, algorith visualization). |
| [chart](https://github.com/vdobler/chart) | 544 | 41 | 2011-06-27 | 2 weeks ago | Simple Chart Plotting library for Go. Supports many graphs types. |
| [ewma](https://github.com/VividCortex/ewma) | 254 | 30 | 2013-07-06 | 1 year ago | Exponentially-weighted moving averages. |
| [graph](https://github.com/yourbasic/graph) | 188 | 13 | 2017-04-28 | 1 year ago | Library of basic graph algorithms. |
| [gohistogram](https://github.com/VividCortex/gohistogram) | 117 | 18 | 2013-07-02 | 11 months ago | Approximate histograms for data streams. |
| [orb](https://github.com/paulmach/orb) | 115 | 13 | 2016-03-28 | 2 weeks ago | 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support. |
| [TextRank](https://github.com/DavidBelicza/TextRank) | 59 | 5 | 2018-01-10 | 6 months ago | TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support. |
| [sparse](https://github.com/james-bowman/sparse) | 56 | 4 | 2017-05-16 | 2 months ago | Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries. |
| [pagerank](https://github.com/alixaxel/pagerank) | 42 | 6 | 2015-08-06 | 3 years ago | Weighted PageRank algorithm implemented in Go. |
| [geom](https://github.com/skelterjohn/geom) | 40 | 4 | 2011-06-08 | 1 year ago | 2D geometry for golang. |
| [evaler](https://github.com/soniah/evaler) | 38 | 5 | 2012-09-05 | 7 months ago | Simple floating point arithmetic expression evaluator. |
| [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) | 33 | 8 | 2018-10-01 | 4 months ago | Dataframes for Go for machine-learning and statistics (similar to pandas). |
| [goent](https://github.com/kzahedi/goent) | 12 | 1 | 2017-08-08 | 2 months ago | GO Implementation of Entropy Measures. |
| [go-fn](https://github.com/ematvey/go-fn) | 11 | 0 | 2015-04-28 | 3 years ago | Mathematical functions written in Go language, that are not covered by math pkg. |
| [ode](https://github.com/ChristopherRabotin/ode) | 9 | 3 | 2016-11-12 | 1 year ago | Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions. |
| [GoStats](https://github.com/OGFris/GoStats) | 8 | 1 | 2018-07-23 | 1 month ago | GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions. |
| [triangolatte](https://github.com/tchayen/triangolatte) | 7 | 1 | 2018-07-19 | 6 months ago | 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs. |
| [PiHex](https://github.com/claygod/PiHex) | 6 | 2 | 2016-07-22 | 1 year ago | Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi. |
| [piecewiselinear](https://github.com/sgreben/piecewiselinear) | 5 | 1 | 2018-10-21 | 3 months ago | Tiny linear interpolation library. |
| [go-gt](https://github.com/ThePaw/go-gt) | 5 | 0 | 2015-09-14 | 3 years ago | Graph theory algorithms written in "Go" language. |
| [gocomplex](https://github.com/varver/gocomplex) | 5 | 0 | 2015-06-25 | 3 years ago | Complex number library for the Go programming language. |
| [rootfinding](https://github.com/khezen/rootfinding) | 2 | 1 | 2018-10-31 | 1 month ago | root-finding algorithms library for finding roots of quadratic functions. |

## Security
        
*Libraries that are used to help make your application more secure.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [lego](https://github.com/xenolf/lego) | 3176 | 89 | 2015-06-08 | 19 hours ago | Pure Go ACME client library and CLI tool (for use with Let's Encrypt). |
| [acme](https://github.com/hlandau/acme) | 1653 | 68 | 2015-11-15 | 6 months ago | ACME (Let's Encrypt) client tool with automatic renewal. |
| [cameradar](https://github.com/Ullaakut/cameradar) | 1616 | 95 | 2016-05-20 | 1 month ago | Tool and library to remotely hack RTSP streams from surveillance cameras. |
| [secure](https://github.com/unrolled/secure) | 1113 | 33 | 2014-05-21 | 2 weeks ago | HTTP middleware for Go that facilitates some quick security wins. |
| [memguard](https://github.com/awnumar/memguard) | 903 | 31 | 2017-04-22 | 1 week ago | A pure Go library for handling sensitive values in memory. |
| [nacl](https://github.com/kevinburke/nacl) | 443 | 13 | 2017-07-21 | 2 months ago | Go implementation of the NaCL set of API's. |
| [acra](https://github.com/cossacklabs/acra) | 374 | 28 | 2016-11-15 | 1 hour ago | Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system. |
| [badactor](https://github.com/jaredfolkins/badactor) | 237 | 8 | 2014-12-13 | 1 year ago | In-memory, application-driven jailer built in the spirit of fail2ban. |
| [passlib](https://github.com/hlandau/passlib) | 212 | 11 | 2014-12-22 | 2 months ago | Futureproof password hashing library. |
| [ssh-vault](https://github.com/ssh-vault/ssh-vault) | 191 | 10 | 2016-09-29 | 2 months ago | encrypt/decrypt using ssh keys. |
| [simple-scrypt](https://github.com/elithrar/simple-scrypt) | 147 | 6 | 2015-04-14 | 9 months ago | Scrypt package with a simple, obvious API and automatic cost calibration built-in. |
| [go-yara](https://github.com/hillu/go-yara) | 122 | 17 | 2015-01-25 | 6 days ago | Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)". |
| [argon2pw](https://github.com/raja/argon2pw) | 66 | 3 | 2018-03-13 | 6 months ago | Argon2 password hash generation with constant-time password comparison. |
| [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) | 25 | 1 | 2017-10-20 | 2 weeks ago | A probably paranoid package for securely hashing and encrypting passwords. |
| [goArgonPass](https://github.com/dwin/goArgonPass) | 7 | 2 | 2018-05-30 | 1 week ago | Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations. |
| [jwc](https://github.com/khezen/jwc) | 3 | 1 | 2018-06-28 | 1 month ago | JSON Web Cryptography library. |

## Serialization
        
*Libraries and tools for binary serialization.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [go](https://github.com/json-iterator/go) | 4454 | 178 | 2016-11-30 | 6 days ago | High-performance 100% compatible drop-in replacement of "encoding/json". |
| [protobuf](https://github.com/golang/protobuf) | 4150 | 187 | 2014-11-24 | 2 days ago | Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers. |
| [protobuf](https://github.com/gogo/protobuf) | 2418 | 80 | 2014-12-03 | 2 weeks ago | Protocol Buffers for Go with Gadgets. |
| [mapstructure](https://github.com/mitchellh/mapstructure) | 2046 | 44 | 2013-05-20 | 4 days ago | Go library for decoding generic map values into native Go structures. |
| [go](https://github.com/ugorji/go) | 1118 | 47 | 2013-05-30 | 1 week ago | High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support. |
| [colfer](https://github.com/pascaldekloe/colfer) | 441 | 34 | 2015-09-06 | 6 months ago | Code generation for the Colfer binary format. |
| [go-capnproto](https://github.com/glycerine/go-capnproto) | 269 | 11 | 2013-11-08 | 1 month ago | Cap'n Proto library and parser for go. |
| [csvutil](https://github.com/jszwec/csvutil) | 266 | 7 | 2017-10-30 | 6 days ago | High Performance, idiomatic CSV record encoding and decoding to native Go structures. |
| [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) | 114 | 9 | 2012-12-23 | 4 months ago | GoLang library for working with PHP session format and PHP Serialize/Unserialize functions. |
| [structomap](https://github.com/danhper/structomap) | 84 | 7 | 2015-05-13 | 3 years ago | Library to easily and dynamically generate maps from static structures. |
| [bambam](https://github.com/glycerine/bambam) | 60 | 3 | 2014-09-17 | 2 years ago | generator for Cap'n Proto schemas from go. |
| [asn1](https://github.com/Logicalis/asn1) | 36 | 8 | 2016-02-29 | 1 year ago | Asn.1 BER and DER encoding library for golang. |
| [fwencoder](https://github.com/o1egl/fwencoder) | 6 | 1 | 2017-12-25 | 1 year ago | Fixed width file parser (encoding and decoding library) for Go. |
| [bel](https://github.com/32leaves/bel) | 0 | 1 | 2019-02-21 | 6 days ago | Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC. |

## Template Engines
        
*Libraries and tools for templating and lexing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gofpdf](https://github.com/jung-kurt/gofpdf) | 2746 | 79 | 2015-03-13 | 4 days ago | PDF document generator with high level support for text, drawing and images. |
| [pongo2](https://github.com/flosch/pongo2) | 1388 | 55 | 2013-08-23 | 2 weeks ago | Django-like template-engine for Go. |
| [quicktemplate](https://github.com/valyala/quicktemplate) | 1217 | 49 | 2016-03-07 | 1 month ago | Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it. |
| [hero](https://github.com/shiyanhui/hero) | 1133 | 38 | 2017-01-15 | 4 months ago | Hero is a handy, fast and powerful go template engine. |
| [mustache](https://github.com/hoisie/mustache) | 950 | 35 | 2009-12-31 | 9 months ago | Go implementation of the Mustache template language. |
| [amber](https://github.com/eknkc/amber) | 803 | 23 | 2012-11-01 | 6 months ago | Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade. |
| [ace](https://github.com/yosssi/ace) | 749 | 23 | 2014-07-13 | 8 months ago | Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold. |
| [gorazor](https://github.com/sipin/gorazor) | 648 | 57 | 2014-05-01 | 5 months ago | Razor view engine for Golang. |
| [jet](https://github.com/CloudyKit/jet) | 546 | 22 | 2016-04-01 | 22 hours ago | Jet template engine. |
| [ego](https://github.com/benbjohnson/ego) | 401 | 16 | 2014-02-24 | 1 month ago | Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled. |
| [raymond](https://github.com/aymerick/raymond) | 312 | 12 | 2015-04-22 | 1 month ago | Complete handlebars implementation in Go. |
| [fasttemplate](https://github.com/valyala/fasttemplate) | 250 | 17 | 2015-08-19 | 2 days ago | Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/). |
| [soy](https://github.com/robfig/soy) | 141 | 13 | 2013-12-15 | 3 days ago | Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/). |
| [liquid](https://github.com/osteele/liquid) | 73 | 4 | 2017-06-26 | 4 months ago | Go implementation of Shopify Liquid templates. |
| [kasia.go](https://github.com/ziutek/kasia.go) | 70 | 2 | 2010-12-07 | 3 years ago | Templating system for HTML and other text documents - go implementation. |
| [velvet](https://github.com/gobuffalo/velvet) | 63 | 6 | 2016-12-30 | 1 year ago | Complete handlebars implementation in Go. |
| [damsel](https://github.com/dskinner/damsel) | 20 | 4 | 2012-05-03 | 2 years ago | Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others. |
| [extemplate](https://github.com/dannyvankooten/extemplate) | 10 | 1 | 2018-08-11 | 6 months ago | Tiny wrapper around html/template to allow for easy file-based template inheritance. |

## Testing
        
*Libraries for testing codebases and generating test data.*

### Testing Frameworks
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [testify](https://github.com/stretchr/testify) | 6877 | 134 | 2012-10-17 | 1 day ago | Sacred extension to the standard go testing package. |
| [goconvey](https://github.com/smartystreets/goconvey) | 4290 | 143 | 2013-08-21 | 1 day ago | BDD-style framework with web UI and live reload. |
| [httpexpect](https://github.com/gavv/httpexpect) | 1035 | 32 | 2016-04-30 | 7 months ago | Concise, declarative, and easy to use end-to-end HTTP and REST API testing. |
| [go-cmp](https://github.com/google/go-cmp) | 813 | 21 | 2017-07-08 | 4 days ago | Package for comparing Go values in tests. |
| [baloo](https://github.com/h2non/baloo) | 601 | 12 | 2016-05-30 | 1 month ago | Expressive and versatile end-to-end HTTP API testing made easy. |
| [goblin](https://github.com/franela/goblin) | 592 | 16 | 2013-09-19 | 5 months ago | Mocha like testing framework fo Go. |
| [godog](https://github.com/DATA-DOG/godog) | 567 | 27 | 2015-06-10 | 6 days ago | Cucumber or Behat like BDD framework for Go. |
| [go-vcr](https://github.com/dnaeon/go-vcr) | 273 | 7 | 2015-12-14 | 2 months ago | Record and replay your HTTP interactions for fast, deterministic and accurate tests. |
| [testfixtures](https://github.com/go-testfixtures/testfixtures) | 260 | 5 | 2016-04-05 | 2 months ago | A helper for Rails' like test fixtures to test database applications. |
| [frisby](https://github.com/verdverm/frisby) | 240 | 8 | 2015-09-15 | 11 months ago | REST API testing framework. |
| [gofight](https://github.com/appleboy/gofight) | 231 | 13 | 2016-03-29 | 1 month ago | API Handler Testing for Golang Router framework. |
| [go-mutesting](https://github.com/zimmski/go-mutesting) | 226 | 7 | 2014-12-27 | 3 months ago | Mutation testing for Go source code. |
| [go-carpet](https://github.com/msoap/go-carpet) | 189 | 5 | 2016-02-28 | 2 days ago | Tool for viewing test coverage in terminal. |
| [charlatan](https://github.com/percolate/charlatan) | 185 | 39 | 2017-10-07 | 5 months ago | Tool to generate fake interface implementations for tests. |
| [gospec](https://github.com/luontola/gospec) | 110 | 4 | 2009-11-24 | 4 years ago | BDD-style testing framework for the Go programming language. |
| [gotest.tools](https://github.com/gotestyourself/gotest.tools) | 83 | 5 | 2017-08-09 | 5 hours ago | A collection of packages to augment the go testing package and support common patterns. |
| [endly](https://github.com/viant/endly) | 74 | 11 | 2017-08-29 | 3 days ago | Declarative end to end functional testing. |
| [cupaloy](https://github.com/bradleyjkemp/cupaloy) | 64 | 3 | 2017-08-08 | 1 month ago | Simple snapshot testing addon for your test framework. |
| [dbcleaner](https://github.com/khaiql/dbcleaner) | 62 | 3 | 2017-01-18 | 1 week ago | Clean database for testing purpose, inspired by `database_cleaner` in Ruby. |
| [wstest](https://github.com/posener/wstest) | 56 | 2 | 2017-04-01 | 1 year ago | Websocket client for unit-testing a websocket http.Handler. |
| [gospecify](https://github.com/stesla/gospecify) | 50 | 5 | 2009-11-20 | 7 years ago | This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec. |
| [restit](https://github.com/go-restit/restit) | 49 | 6 | 2014-06-25 | 2 years ago | Go micro framework to help writing RESTful API integration test. |
| [go-testdeep](https://github.com/maxatome/go-testdeep) | 41 | 2 | 2018-05-26 | 1 day ago | Extremely flexible golang deep comparison, extends the go testing package. |
| [hamcrest](https://github.com/rdrdr/hamcrest) | 26 | 3 | 2010-12-22 | 8 years ago | fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results. |
| [gomatch](https://github.com/jfilipczyk/gomatch) | 25 | 1 | 2019-01-28 | 1 month ago | library created for testing JSON against patterns. |
| [bro](https://github.com/marioidival/bro) | 24 | 3 | 2015-08-11 | 2 years ago | Watch files in directory and run tests for them. |
| [dsunit](https://github.com/viant/dsunit) | 22 | 10 | 2016-06-14 | 2 weeks ago | Datastore testing for SQL, NoSQL, structured files. |
| [jsonassert](https://github.com/kinbiko/jsonassert) | 17 | 0 | 2018-10-27 | 3 weeks ago | Package for verifying that your JSON payloads are serialized correctly. |
| [assert](https://github.com/go-playground/assert) | 12 | 1 | 2015-07-21 | 3 years ago | Basic Assertion Library used along side native go testing, with building blocks for custom assertions. |
| [gosuite](https://github.com/pavlo/gosuite) | 8 | 1 | 2016-10-16 | 2 years ago | Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests. |
| [gocrest](https://github.com/corbym/gocrest) | 8 | 1 | 2017-12-24 | 1 year ago | Composable hamcrest-like matchers for Go assertions. |
| [gogiven](https://github.com/corbym/gogiven) | 7 | 4 | 2018-01-01 | 1 year ago | YATSPEC-like BDD testing framework for Go. |
| [badio](https://github.com/cavaliercoder/badio) | 7 | 1 | 2016-02-11 | 3 years ago | Extensions to Go's `testing/iotest` package. |
| [testsql](https://github.com/zhulongcheng/testsql) | 6 | 1 | 2018-09-22 | 4 months ago | Generate test data from SQL files before testing and clear it after finished. |
| [biff](https://github.com/fulldump/biff) | 5 | 1 | 2018-03-29 | 9 months ago | Bifurcation testing framework, BDD compatible. |
| [tt](https://github.com/vcaesar/tt) | 3 | 1 | 2018-04-03 | 1 month ago | Simple and colorful test tools. |

### Mock
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [chromedp](https://github.com/chromedp/chromedp) | 2785 | 112 | 2017-01-24 | 14 hours ago | a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol. |
| [go-fuzz](https://github.com/dvyukov/go-fuzz) | 2579 | 81 | 2015-04-15 | 11 hours ago | Randomized testing system. |
| [mock](https://github.com/golang/mock) | 2163 | 53 | 2015-06-13 | 6 days ago | Mocking framework for the Go programming language. |
| [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) | 1325 | 15 | 2014-02-07 | 3 days ago | Mock SQL driver for testing database interactions. |
| [hoverfly](https://github.com/SpectoLabs/hoverfly) | 1311 | 57 | 2015-12-01 | 4 days ago | HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI. |
| [selenoid](https://github.com/aerokube/selenoid) | 953 | 85 | 2016-08-22 | 1 day ago | alternative Selenium hub server that launches browsers within containers. |
| [gock](https://github.com/h2non/gock) | 681 | 16 | 2016-03-03 | 4 days ago | Versatile HTTP mocking made easy. |
| [gofuzz](https://github.com/google/gofuzz) | 475 | 19 | 2014-08-01 | 1 year ago | Library for populating go objects with random values. |
| [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) | 315 | 11 | 2014-05-21 | 2 weeks ago | Tool for generating self-contained mock objects. |
| [cdp](https://github.com/mafredri/cdp) | 265 | 13 | 2017-03-12 | 3 weeks ago | Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it. |
| [tavor](https://github.com/zimmski/tavor) | 202 | 12 | 2014-05-18 | 4 months ago | Generic fuzzing and delta-debugging framework. |
| [ggr](https://github.com/aerokube/ggr) | 189 | 22 | 2016-06-16 | 4 days ago | a lightweight server that routes and proxies Selenium Wedriver requests to multiple Selenium hubs. |
| [minimock](https://github.com/gojuno/minimock) | 162 | 5 | 2016-08-04 | 5 days ago | Mock generator for Go interfaces. |
| [go-txdb](https://github.com/DATA-DOG/go-txdb) | 122 | 8 | 2015-07-08 | 3 weeks ago | Single transaction based database driver mainly for testing purposes. |
| [govcr](https://github.com/seborama/govcr) | 67 | 2 | 2016-07-11 | 2 days ago | HTTP mock for Golang: record and replay HTTP interactions for offline testing. |
| [mockhttp](https://github.com/tv42/mockhttp) | 22 | 1 | 2011-06-12 | 4 years ago | Mock object for Go http.ResponseWriter. |

## Text Processing
        
*Libraries for parsing and manipulating texts.*

### Specific Formats
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [colly](https://github.com/gocolly/colly) | 6990 | 215 | 2017-09-29 | 1 day ago | Fast and Elegant Scraping Framework for Gophers. |
| [goquery](https://github.com/PuerkitoBio/goquery) | 6868 | 248 | 2012-08-29 | 1 month ago | GoQuery brings a syntax and a set of features similar to jQuery to the Go language. |
| [blackfriday](https://github.com/russross/blackfriday) | 3579 | 86 | 2011-05-28 | 2 weeks ago | Markdown processor in Go. |
| [toml](https://github.com/BurntSushi/toml) | 2471 | 76 | 2013-02-26 | 3 weeks ago | TOML configuration format (encoder/decoder with reflection). |
| [go-humanize](https://github.com/dustin/go-humanize) | 1721 | 29 | 2012-01-13 | 1 month ago | Formatters for time, numbers, and memory size to human readable format. |
| [sh](https://github.com/mvdan/sh) | 1697 | 38 | 2016-01-16 | 6 days ago | Shell parser and formatter. |
| [bluemonday](https://github.com/microcosm-cc/bluemonday) | 1106 | 30 | 2013-11-21 | 2 months ago | HTML Sanitizer. |
| [inject](https://github.com/facebookgo/inject) | 1086 | 42 | 2013-10-21 | 1 month ago | Package inject provides a reflect based injector. |
| [gofeed](https://github.com/mmcdole/gofeed) | 998 | 33 | 2016-01-23 | 2 weeks ago | Parse RSS and Atom feeds in Go. |
| [commonregex](https://github.com/mingrammer/commonregex) | 532 | 20 | 2017-03-23 | 2 months ago | A collection of common regular expressions for Go. |
| [slug](https://github.com/gosimple/slug) | 327 | 9 | 2014-03-31 | 1 month ago | URL-friendly slugify with multiple languages support. |
| [mxj](https://github.com/clbanning/mxj) | 314 | 22 | 2014-02-03 | 3 hours ago | Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages. |
| [gommon](https://github.com/labstack/gommon) | 271 | 19 | 2015-03-13 | 1 month ago | Common packages for Go |
| [gographviz](https://github.com/awalterschulze/gographviz) | 245 | 10 | 2015-03-15 | 1 week ago | Parses the Graphviz DOT language. |
| [dataflowkit](https://github.com/slotix/dataflowkit) | 226 | 15 | 2017-02-09 | 16 hours ago | Web scraping Framework to turn websites into structured data. |
| [gotext](https://github.com/leonelquinteros/gotext) | 209 | 5 | 2016-06-20 | 2 weeks ago | GNU gettext utilities for Go. |
| [go-runewidth](https://github.com/mattn/go-runewidth) | 182 | 12 | 2013-06-21 | 2 months ago | Functions to get fixed width of the character or string. |
| [goq](https://github.com/andrewstuart/goq) | 128 | 7 | 2017-02-20 | 9 months ago | Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery). |
| [htmlquery](https://github.com/antchfx/htmlquery) | 87 | 3 | 2017-12-05 | 3 weeks ago | An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression. |
| [go-nmea](https://github.com/adrianmo/go-nmea) | 77 | 3 | 2015-07-22 | 2 weeks ago | NMEA parser library for the Go language. |
| [sdp](https://github.com/gortc/sdp) | 52 | 7 | 2016-05-13 | 2 weeks ago | SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)]. |
| [genex](https://github.com/alixaxel/genex) | 48 | 3 | 2015-03-10 | 3 years ago | Count and expand Regular Expressions into all matching Strings. |
| [align](https://github.com/Guitarbum722/align) | 46 | 5 | 2017-04-30 | 1 year ago | A general purpose application that aligns text. |
| [guesslanguage](https://github.com/endeveit/guesslanguage) | 41 | 1 | 2014-12-16 | 1 year ago | Functions to determine the natural language of a unicode text. |
| [go-zero-width](https://github.com/trubitsyn/go-zero-width) | 40 | 1 | 2018-06-18 | 2 months ago | Zero-width character detection and removal for Go. |
| [allot](https://github.com/sbstjn/allot) | 34 | 1 | 2016-10-16 | 3 months ago | Placeholder and wildcard text parsing for CLI tools and bots. |
| [goregen](https://github.com/zach-klippenstein/goregen) | 33 | 2 | 2014-12-27 | 3 years ago | Library for generating random strings from regular expressions. |
| [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) | 30 | 7 | 2016-07-05 | 3 months ago | Editorconfig file parser and manipulator for Go. |
| [gonameparts](https://github.com/polera/gonameparts) | 28 | 0 | 2015-05-17 | 1 year ago | Parses human names into individual name parts. |
| [go-slugify](https://github.com/mozillazg/go-slugify) | 26 | 3 | 2016-07-16 | 2 years ago | Make pretty slug with multiple languages support. |
| [slugify](https://github.com/avelino/slugify) | 24 | 2 | 2015-04-13 | 10 months ago | Go slugify application that handles string. |
| [go-vcard](https://github.com/emersion/go-vcard) | 19 | 2 | 2017-03-21 | 1 month ago | Parse and format vCard. |
| [did](https://github.com/ockam-network/did) | 17 | 3 | 2018-11-03 | 3 months ago | DID (Decentralized Identifiers) Parser and Stringer in Go. |
| [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) | 14 | 1 | 2017-11-16 | 6 months ago | Fixed-width text formatting (encoder/decoder with reflection). |
| [bbConvert](https://github.com/CalebQ42/bbConvert) | 5 | 1 | 2016-04-15 | 2 years ago | Converts bbCode to HTML that allows you to add support for custom bbCode tags. |
| [enca](https://github.com/endeveit/enca) | 5 | 1 | 2014-12-17 | 3 years ago | Minimal cgo bindings for [libenca](http://cihar.com/software/enca/). |
| [doi](https://github.com/hscells/doi) | 4 | 1 | 2017-08-02 | 1 year ago | Document object identifier (doi) parser in Go. |
| [syndfeed](https://github.com/zhengchun/syndfeed) | 4 | 1 | 2017-04-07 | 11 months ago | A syndication feed for Atom 1.0 and RSS 2.0. |
| [encoding](https://github.com/mickep76/encoding) | 3 | 1 | 2018-04-07 | 5 months ago | Package provides a generic interface to encoders and decodersa. |

### Utility
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [xurls](https://github.com/mvdan/xurls) | 407 | 13 | 2015-01-12 | 2 weeks ago | Extract urls from text. |
| [gotabulate](https://github.com/bndr/gotabulate) | 192 | 8 | 2014-08-21 | 1 year ago | Easily pretty-print your tabular data with Go. |
| [radix](https://github.com/yourbasic/radix) | 61 | 2 | 2017-06-09 | 1 year ago | fast string sorting algorithm. |
| [parth](https://github.com/codemodus/parth) | 27 | 3 | 2015-04-07 | 1 month ago | URL path segmentation parsing. |
| [xj2go](https://github.com/stackerzzq/xj2go) | 15 | 1 | 2017-09-19 | 2 months ago | Convert xml or json to go struct. |
| [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) | 10 | 1 | 2018-09-09 | 3 months ago | A sanitization-based swear filter for Go. |
| [kace](https://github.com/codemodus/kace) | 9 | 1 | 2015-06-05 | 6 months ago | Common case conversions covering common initialisms. |
| [parseargs-go](https://github.com/txgruppi/parseargs-go) | 4 | 1 | 2016-02-24 | 2 years ago | string argument parser that understands quotes and backslashes. |

## Third-party APIs
        
*Libraries for accessing third party APIs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [aws-sdk-go](https://github.com/aws/aws-sdk-go) | 4566 | 254 | 2014-12-05 | 15 hours ago | The official AWS SDK for the Go programming language. |
| [go-github](https://github.com/google/go-github) | 4370 | 175 | 2013-05-25 | 2 days ago | Go library for accessing the GitHub REST API v3. |
| [slack](https://github.com/nlopes/slack) | 2100 | 43 | 2015-01-24 | 2 days ago | Slack API in Go. |
| [google-api-go-client](https://github.com/googleapis/google-api-go-client) | 1744 | 127 | 2014-11-25 | 2 days ago | Auto-generated Google APIs for Go. |
| [google-cloud-go](https://github.com/googleapis/google-cloud-go) | 1558 | 210 | 2014-05-09 | 4 days ago | Google Cloud APIs Go Client Library. |
| [anaconda](https://github.com/ChimeraCoder/anaconda) | 953 | 23 | 2013-03-05 | 3 weeks ago | Go client library for the Twitter 1.1 API. |
| [stripe-go](https://github.com/stripe/stripe-go) | 847 | 35 | 2014-06-06 | 17 hours ago | Go client for the Stripe API. |
| [discordgo](https://github.com/bwmarrin/discordgo) | 803 | 50 | 2015-11-02 | 17 hours ago | Go bindings for the Discord Chat API. |
| [facebook](https://github.com/huandu/facebook) | 724 | 86 | 2012-07-29 | 2 months ago | Go Library that supports the Facebook Graph API. |
| [minio-go](https://github.com/minio/minio-go) | 629 | 41 | 2015-05-02 | 17 hours ago | Minio Go Library for Amazon S3 compatible cloud storage. |
| [go-twitter](https://github.com/dghubble/go-twitter) | 611 | 27 | 2015-04-12 | 4 hours ago | Go client library for the Twitter v1.1 APIs. |
| [githubv4](https://github.com/shurcooL/githubv4) | 402 | 18 | 2017-05-27 | 1 month ago | Go library for accessing the GitHub GraphQL API v4. |
| [webhooks](https://github.com/go-playground/webhooks) | 280 | 12 | 2015-10-26 | 3 hours ago | Webhook receiver for GitHub and Bitbucket. |
| [geo-golang](https://github.com/codingsince1985/geo-golang) | 280 | 11 | 2014-12-04 | 1 month ago | Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs. |
| [PayPal-Go-SDK](https://github.com/logpacker/PayPal-Go-SDK) | 267 | 16 | 2015-10-14 | 4 days ago | Wrapper for PayPal payment API. |
| [go-marathon](https://github.com/gambol99/go-marathon) | 184 | 12 | 2015-02-11 | 4 days ago | Go library for interacting with Mesosphere's Marathon PAAS. |
| [ethrpc](https://github.com/onrik/ethrpc) | 156 | 9 | 2017-01-24 | 2 hours ago | Go bindings for Ethereum JSON RPC API. |
| [gostorm](https://github.com/jsgilmore/gostorm) | 117 | 11 | 2013-07-22 | 1 year ago | GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. |
| [hipchat](https://github.com/daneharrigan/hipchat) | 112 | 6 | 2013-04-28 | 1 year ago | A golang package to communicate with HipChat over XMPP. |
| [hipchat](https://github.com/andybons/hipchat) | 110 | 8 | 2012-10-21 | 2 years ago | This project implements a golang client library for the Hipchat API. |
| [medium-sdk-go](https://github.com/Medium/medium-sdk-go) | 108 | 104 | 2015-09-27 | 4 months ago | Golang SDK for Medium's OAuth2 API. |
| [go-trending](https://github.com/andygrunwald/go-trending) | 96 | 7 | 2015-07-04 | 4 months ago | Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github. |
| [trello](https://github.com/adlio/trello) | 86 | 5 | 2016-09-24 | 2 days ago | Go wrapper for the Trello API. |
| [cachet](https://github.com/andygrunwald/cachet) | 63 | 5 | 2015-10-31 | 11 months ago | Go client library for [Cachet (open source status page system)](https://cachethq.io/). |
| [clarifai-go](https://github.com/Clarifai/clarifai-go) | 56 | 37 | 2015-09-29 | 1 year ago | Go client library for interfacing with the Clarifai API. |
| [megos](https://github.com/andygrunwald/megos) | 56 | 5 | 2015-10-02 | 10 months ago | Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster. |
| [pushover](https://github.com/gregdel/pushover) | 53 | 2 | 2015-02-19 | 2 weeks ago | Go wrapper for the Pushover API. |
| [gads](https://github.com/emiddleton/gads) | 41 | 8 | 2014-01-20 | 2 years ago | Google Adwords Unofficial API. |
| [go-amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) | 37 | 1 | 2016-11-15 | 11 months ago | Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html). |
| [go-circleci](https://github.com/jszwedko/go-circleci) | 36 | 3 | 2015-08-15 | 1 month ago | Go client library for interacting with CircleCI's API. |
| [go-xkcd](https://github.com/nishanths/go-xkcd) | 36 | 3 | 2016-02-26 | 2 years ago | Go client for the xkcd API. |
| [gomusicbrainz](https://github.com/michiwend/gomusicbrainz) | 33 | 8 | 2014-09-11 | 4 months ago | Go MusicBrainz WS2 client library. |
| [fcm](https://github.com/maddevsio/fcm) | 32 | 4 | 2017-01-06 | 3 months ago | Go library for Firebase Cloud Messaging. |
| [gcm](https://github.com/Aorioli/gcm) | 29 | 3 | 2015-11-10 | 3 years ago | Go library for Google Cloud Messaging. |
| [uptimerobot](https://github.com/bitfield/uptimerobot) | 28 | 3 | 2018-05-29 | 2 months ago | Go wrapper and command-line client for the Uptime Robot v2 API. |
| [golyrics](https://github.com/mamal72/golyrics) | 28 | 3 | 2016-11-18 | 8 months ago | Golyrics is a Go library to fetch music lyrics data from the Wikia website. |
| [wit-go](https://github.com/wit-ai/wit-go) | 27 | 6 | 2018-08-20 | 1 month ago | Go client for wit.ai HTTP API. |
| [mixpanel](https://github.com/dukex/mixpanel) | 27 | 3 | 2014-05-20 | 5 months ago | Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications. |
| [gami](https://github.com/bit4bit/gami) | 26 | 4 | 2014-05-15 | 8 months ago | Go library for Asterisk Manager Interface. |
| [translate](https://github.com/nuveo/translate) | 25 | 20 | 2015-07-13 | 3 years ago | Go online translation package. |
| [igdb](https://github.com/Henry-Sarabia/igdb) | 24 | 1 | 2017-08-24 | 10 hours ago | Go client for the [Internet Game Database API](https://api.igdb.com/). |
| [go-unsplash](https://github.com/hbagdi/go-unsplash) | 19 | 1 | 2017-01-19 | 7 months ago | Go client library for the [Unsplash.com](https://unsplash.com) API. |
| [go-shopify](https://github.com/rapito/go-shopify) | 19 | 1 | 2014-10-28 | 1 year ago | Go Library to make CRUD request to the Shopify API. |
| [go-spotify](https://github.com/rapito/go-spotify) | 16 | 1 | 2014-10-30 | 1 year ago | Go Library to access Spotify WEB API. |
| [patreon-go](https://github.com/mxpv/patreon-go) | 15 | 4 | 2017-08-07 | 1 month ago | Go library for Patreon API. |
| [go-twitch](https://github.com/knspriggs/go-twitch) | 15 | 4 | 2016-06-29 | 1 year ago | Go client for interacting with the Twitch v3 API. |
| [go-myanimelist](https://github.com/nstratos/go-myanimelist) | 15 | 2 | 2015-05-03 | 2 years ago | Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api). |
| [brewerydb](https://github.com/naegelejd/brewerydb) | 15 | 2 | 2015-04-15 | 3 years ago | Go library for accessing the BreweryDB API. |
| [codeship-go](https://github.com/codeship/codeship-go) | 14 | 18 | 2017-09-09 | 1 month ago | Go client library for interacting with Codeship's API v2. |
| [go-steam](https://github.com/sostronk/go-steam) | 13 | 10 | 2014-11-24 | 11 months ago | Go Library to interact with Steam game servers. |
| [textbelt](https://github.com/farmergreg/textbelt) | 12 | 2 | 2015-09-02 | 3 years ago | Go client for the textbelt.com txt messaging API. |
| [go-google-analytics](https://github.com/chonthu/go-google-analytics) | 11 | 2 | 2015-06-01 | 3 years ago | Simple wrapper for easy google analytics reporting. |
| [go-tmdb](https://github.com/jbrodriguez/go-tmdb) | 11 | 1 | 2014-10-20 | 3 years ago | Simple golang package to communicate with [themoviedb.org](https://themoviedb.org). |
| [go-hacknews](https://github.com/PaulRosset/go-hacknews) | 9 | 2 | 2017-08-11 | 1 year ago | Tiny Go client for HackerNews API. |
| [simples3](https://github.com/rhnvrm/simples3) | 9 | 0 | 2018-12-06 | 1 month ago | Simple no frills AWS S3 Library using REST with V4 Signing written in Go. |
| [smitego](https://github.com/sergiotapia/smitego) | 9 | 0 | 2013-12-11 | 4 years ago | Go package to wraps access to the Smite game API. |
| [go-sptrans](https://github.com/sergioaugrod/go-sptrans) | 8 | 2 | 2017-09-11 | 1 year ago | Go client library for the SPTrans Olho Vivo API. |
| [rrdaclient](https://github.com/Omie/rrdaclient) | 8 | 1 | 2014-09-16 | 4 years ago | Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP. |
| [coinpaprika-api-go-client](https://github.com/coinpaprika/coinpaprika-api-go-client) | 8 | 6 | 2018-09-25 | 1 month ago | Go client library for interacting with Coinpaprika's API. |
| [ynab.go](https://github.com/brunomvsouza/ynab.go) | 7 | 1 | 2018-07-13 | 1 week ago | Go wrapper for the YNAB API. |
| [go-google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) | 6 | 0 | 2016-10-24 | 2 years ago | Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/). |
| [gumblr](https://github.com/mattcunningham/gumblr) | 6 | 1 | 2015-07-10 | 2 years ago | Go wrapper for the Tumblr v2 API. |
| [go-zooz](https://github.com/gojuno/go-zooz) | 5 | 15 | 2017-07-04 | 3 months ago | Go client for the Zooz API. |
| [go-sophos](https://github.com/esurdam/go-sophos) | 4 | 1 | 2018-09-05 | 5 months ago | Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies. |
| [go-chronos](https://github.com/axelspringer/go-chronos) | 3 | 10 | 2017-10-23 | 1 year ago | Go library for interacting with the [Chronos](https://mesos.github. |
| [playlyfe-go-sdk](https://github.com/playlyfe/playlyfe-go-sdk) | 1 | 4 | 2015-05-25 | 3 years ago | The Playlyfe Rest API Go SDK. |

## Utilities
        
*General utilities and tools to make your life easier.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [fzf](https://github.com/junegunn/fzf) | 19975 | 316 | 2013-10-24 | 4 hours ago | Command-line fuzzy finder written in Go. |
| [hub](https://github.com/github/hub) | 14923 | 310 | 2009-12-06 | 2 days ago | wrap git commands with additional functionality to interact with github from the terminal. |
| [delve](https://github.com/go-delve/delve) | 10679 | 353 | 2014-05-21 | 3 days ago | Go debugger. |
| [ctop](https://github.com/bcicen/ctop) | 8120 | 153 | 2016-12-27 | 1 month ago | [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics. |
| [wuzz](https://github.com/asciimoo/wuzz) | 7948 | 176 | 2017-01-31 | 2 months ago | Interactive cli tool for HTTP inspection. |
| [sqlx](https://github.com/jmoiron/sqlx) | 5809 | 164 | 2013-01-29 | 1 week ago | provides a set of extensions on top of the excellent built-in database/sql package. |
| [peco](https://github.com/peco/peco) | 5189 | 131 | 2014-06-06 | 6 days ago | Simplistic interactive filtering tool. |
| [usql](https://github.com/xo/usql) | 4539 | 111 | 2017-03-02 | 1 week ago | usql is a universal command-line interface for SQL databases. |
| [goreleaser](https://github.com/goreleaser/goreleaser) | 3811 | 61 | 2016-12-22 | 3 hours ago | Deliver Go binaries as fast and easily as possible. |
| [godropbox](https://github.com/dropbox/godropbox) | 3643 | 250 | 2014-06-23 | 1 week ago | Common libraries for writing Go services/applications from Dropbox. |
| [go-torch](https://github.com/uber/go-torch) | 3534 | 2405 | 2015-07-22 | 3 months ago | Stochastic flame graph profiler for Go programs. |
| [realize](https://github.com/oxequa/realize) | 2864 | 68 | 2016-07-12 | 1 month ago | Go build system with file watchers and live reload. Run, build and watch file changes with custom paths. |
| [goreporter](https://github.com/360EntSecGroup-Skylar/goreporter) | 2328 | 93 | 2017-03-27 | 4 months ago | Golang tool that does static analysis, unit testing, code review and generate code quality report. |
| [panicparse](https://github.com/maruel/panicparse) | 1825 | 37 | 2015-02-02 | 7 months ago | Groups similar goroutines and colorizes stack dump. |
| [task](https://github.com/go-task/task) | 1704 | 43 | 2017-02-27 | 11 hours ago | simple "Make" alternative. |
| [minify](https://github.com/tdewolff/minify) | 1694 | 52 | 2014-05-21 | 1 month ago | Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats. |
| [hystrix-go](https://github.com/afex/hystrix-go) | 1639 | 73 | 2013-12-15 | 3 months ago | Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker. |
| [resty](https://github.com/go-resty/resty) | 1592 | 59 | 2015-08-29 | 5 days ago | Simple HTTP and REST client for Go inspired by Ruby rest-client. |
| [mmake](https://github.com/tj/mmake) | 1439 | 34 | 2017-02-16 | 1 year ago | Modern Make. |
| [storm](https://github.com/asdine/storm) | 1255 | 45 | 2016-01-10 | 2 weeks ago | Simple and powerful toolkit for BoltDB. |
| [mole](https://github.com/davrodpin/mole) | 1242 | 27 | 2018-10-04 | 12 hours ago | cli app to easily create ssh tunnels. |
| [mc](https://github.com/minio/mc) | 942 | 53 | 2015-01-16 | 2 hours ago | Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems. |
| [profile](https://github.com/pkg/profile) | 889 | 33 | 2014-10-22 | 4 months ago | Simple profiling support package for Go. |
| [boilr](https://github.com/tmrts/boilr) | 859 | 27 | 2015-12-20 | 7 months ago | Blazingly fast CLI tool for creating projects from boilerplate templates. |
| [filetype](https://github.com/h2non/filetype) | 842 | 17 | 2015-09-24 | 1 week ago | Small package to infer the file type checking the magic numbers signature. |
| [circuitbreaker](https://github.com/rubyist/circuitbreaker) | 733 | 19 | 2014-07-18 | 1 month ago | Circuit Breakers in Go. |
| [go-funk](https://github.com/thoas/go-funk) | 726 | 25 | 2016-12-30 | 3 months ago | Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...). |
| [mergo](https://github.com/imdario/mergo) | 703 | 11 | 2013-03-12 | 13 hours ago | Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements. |
| [spinner](https://github.com/briandowns/spinner) | 677 | 15 | 2014-12-13 | 2 weeks ago | Go package to easily provide a terminal spinner with options. |
| [gtm](https://github.com/git-time-metric/gtm) | 667 | 28 | 2016-06-20 | 8 months ago | Simple, seamless, lightweight time tracking for Git. |
| [jump](https://github.com/gsamokovarov/jump) | 585 | 15 | 2015-08-17 | 4 days ago | Jump helps you navigate faster by learning your habits. |
| [immortal](https://github.com/immortal/immortal) | 574 | 15 | 2016-07-01 | 2 months ago | \*nix cross-platform (OS agnostic) supervisor. |
| [htcat](https://github.com/htcat/htcat) | 472 | 16 | 2013-08-05 | 6 days ago | Parallel and Pipelined HTTP GET Utility. |
| [gopencils](https://github.com/bndr/gopencils) | 419 | 14 | 2014-06-23 | 2 weeks ago | Small and simple package to easily consume REST APIs. |
| [go-dry](https://github.com/ungerik/go-dry) | 418 | 15 | 2014-02-28 | 10 months ago | DRY (don't repeat yourself) package for Go. |
| [godaemon](https://github.com/VividCortex/godaemon) | 382 | 32 | 2013-08-02 | 3 years ago | Utility to write daemons. |
| [request](https://github.com/mozillazg/request) | 340 | 13 | 2014-12-21 | 8 months ago | Go HTTP Requests for Humans™. |
| [go-rate](https://github.com/beefsack/go-rate) | 277 | 12 | 2014-08-25 | 11 months ago | Timed rate limiter for Go. |
| [ergo](https://github.com/cristianoliveira/ergo) | 276 | 4 | 2017-08-20 | 3 months ago | The management of multiple local services running over different ports made easy. |
| [circuit](https://github.com/cep21/circuit) | 262 | 11 | 2017-12-24 | 3 days ago | An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern. |
| [gohper](https://github.com/cosiner/gohper) | 248 | 19 | 2015-03-24 | 1 year ago | Various tools/modules help for development. |
| [koazee](https://github.com/wesovilabs/koazee) | 230 | 9 | 2018-11-09 | 5 hours ago | Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays. |
| [clockwork](https://github.com/jonboulle/clockwork) | 195 | 5 | 2014-09-10 | 1 month ago | A simple fake clock for golang. |
| [deepcopier](https://github.com/ulule/deepcopier) | 188 | 20 | 2015-07-25 | 1 year ago | Simple struct copying for Go. |
| [serve](https://github.com/syntaqx/serve) | 170 | 6 | 2019-01-11 | 1 month ago | A static http server anywhere you need. |
| [go-trigger](https://github.com/sadlil/go-trigger) | 167 | 12 | 2015-10-19 | 1 year ago | Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project. |
| [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) | 157 | 3 | 2016-11-08 | 1 year ago | go:generate tool for wrapping symbols exported by golang plugins (1.8 only). |
| [rerun](https://github.com/ivpusic/rerun) | 150 | 7 | 2014-12-10 | 11 months ago | Recompiling and rerunning go apps when source changes. |
| [moldova](https://github.com/StabbyCutyou/moldova) | 148 | 5 | 2016-01-30 | 1 year ago | Utility for generating random data based on an input template. |
| [robustly](https://github.com/VividCortex/robustly) | 129 | 21 | 2013-07-08 | 11 months ago | Runs functions resiliently, catching and restarting panics. |
| [gotenv](https://github.com/subosito/gotenv) | 125 | 4 | 2013-08-27 | 9 months ago | Load environment variables from `.env` or any `io.Reader` in Go. |
| [death](https://github.com/vrecan/death) | 121 | 4 | 2015-03-09 | 4 months ago | Managing go application shutdown with signals. |
| [apm](https://github.com/topfreegames/apm) | 119 | 14 | 2015-11-19 | 2 years ago | Process manager for Golang applications with an HTTP API. |
| [util](https://github.com/shomali11/util) | 107 | 6 | 2017-05-24 | 9 months ago | Collection of useful utility functions. (strings, concurrency, manipulations, ...). |
| [chyle](https://github.com/antham/chyle) | 102 | 6 | 2016-11-18 | 3 days ago | Changelog generator using a git repository with multiple configuration possibilities. |
| [retry](https://github.com/kamilsk/retry) | 98 | 0 | 2016-11-03 | 7 hours ago | Functional mechanism based on context to perform actions repetitively until successful. |
| [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) | 98 | 5 | 2015-10-13 | 1 month ago | XML Sitemap generator written in Go. |
| [lrserver](https://github.com/jaschaephraim/lrserver) | 98 | 5 | 2014-07-15 | 1 year ago | LiveReload server for Go. |
| [onecache](https://github.com/adelowo/onecache) | 86 | 6 | 2017-04-15 | 3 days ago | Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc). |
| [pm](https://github.com/VividCortex/pm) | 67 | 20 | 2013-11-18 | 2 months ago | Process (i.e. goroutine) manager with an HTTP API. |
| [unis](https://github.com/esemplastic/unis) | 67 | 5 | 2017-05-06 | 1 year ago | Common Architecture™ for String Utilities in Go. |
| [netbug](https://github.com/e-dard/netbug) | 65 | 1 | 2015-03-06 | 3 years ago | Easy remote profiling of your services. |
| [xferspdy](https://github.com/monmohan/xferspdy) | 65 | 4 | 2015-05-22 | 2 years ago | Xferspdy provides binary diff and patch library in golang. |
| [mimetype](https://github.com/gabriel-vasile/mimetype) | 61 | 6 | 2018-07-02 | 4 days ago | Package for MIME type detection based on magic numbers. |
| [toolbox](https://github.com/viant/toolbox) | 60 | 15 | 2016-06-14 | 3 days ago | Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer. |
| [go-health](https://github.com/Talento90/go-health) | 59 | 2 | 2018-02-14 | 8 months ago | Health package simplifies the way you add health check to your services. |
| [multitick](https://github.com/VividCortex/multitick) | 58 | 20 | 2013-12-11 | 2 years ago | Multiplexor for aligned tickers. |
| [repeat](https://github.com/ssgreg/repeat) | 53 | 4 | 2017-11-22 | 1 month ago | Go implementation of different backoff strategies useful for retrying operations and heartbeating. |
| [mssqlx](https://github.com/linxGnu/mssqlx) | 52 | 5 | 2016-12-26 | 1 day ago | Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind. |
| [abutil](https://github.com/bahlo/abutil) | 51 | 3 | 2015-08-26 | 3 years ago | Collection of often-used Golang helpers. |
| [minquery](https://github.com/icza/minquery) | 44 | 3 | 2016-11-16 | 4 months ago | MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off). |
| [go-astitodo](https://github.com/asticode/go-astitodo) | 43 | 2 | 2016-10-18 | 5 months ago | Parse TODOs in your GO code. |
| [golog](https://github.com/mlimaloureiro/golog) | 42 | 3 | 2016-01-09 | 1 month ago | Easy and lightweight CLI tool to time track your tasks. |
| [mimemagic](https://github.com/zRedShift/mimemagic) | 39 | 2 | 2018-10-12 | 2 months ago | Pure Go ultra performant MIME sniffing library/utility. |
| [goback](https://github.com/carlescere/goback) | 39 | 1 | 2015-03-14 | 1 year ago | Go simple exponential backoff package. |
| [intrinsic](https://github.com/mengzhuo/intrinsic) | 37 | 4 | 2017-06-13 | 1 year ago | Use x86 SIMD without writing any assembly code. |
| [handy](https://github.com/miguelpragier/handy) | 37 | 3 | 2018-06-13 | 1 week ago | Many utilities and helpers like string handlers/formatters and validators. |
| [gaper](https://github.com/maxcnunes/gaper) | 34 | 0 | 2018-06-16 | 1 month ago | Builds and restarts a Go project when it crashes or some watched file changes. |
| [copy-pasta](https://github.com/jutkko/copy-pasta) | 34 | 4 | 2017-01-28 | 6 months ago | Universal multi-workstation clipboard that uses S3 like backend for the storage. |
| [golarm](https://github.com/msempere/golarm) | 32 | 1 | 2015-08-15 | 3 years ago | Fire alarms with system events. |
| [retry](https://github.com/thedevsaddam/retry) | 31 | 1 | 2018-02-26 | 11 months ago | Simple and easy retry mechanism package for Go. |
| [myhttp](https://github.com/inancgumus/myhttp) | 30 | 1 | 2017-09-13 | 10 months ago | Simple API to make HTTP GET requests with timeout support. |
| [gpath](https://github.com/tenntenn/gpath) | 25 | 2 | 2017-05-24 | 1 year ago | Library to simplify access struct fields with Go's expression in reflection. |
| [retry-go](https://github.com/rafaeljesus/retry-go) | 24 | 1 | 2017-06-10 | 4 months ago | Retrying made simple and easy for golang. |
| [goplaceholder](https://github.com/michiwend/goplaceholder) | 21 | 2 | 2014-10-12 | 3 years ago | a small golang lib to generate placeholder images. |
| [rclient](https://github.com/zpatrick/rclient) | 21 | 2 | 2017-02-28 | 6 months ago | Readable, flexible, simple-to-use client for REST APIs. |
| [goseaweedfs](https://github.com/linxGnu/goseaweedfs) | 19 | 3 | 2017-07-20 | 3 weeks ago | SeaweedFS client library with almost full features. |
| [ugo](https://github.com/alxrm/ugo) | 19 | 1 | 2016-02-18 | 2 years ago | ugo is slice toolbox with concise syntax for Go. |
| [dlog](https://github.com/kirillDanshin/dlog) | 15 | 2 | 2016-07-05 | 1 year ago | Compile-time controlled logger to make your release smaller without removing debug calls. |
| [generate](https://github.com/go-playground/generate) | 15 | 2 | 2015-11-15 | 2 years ago | runs go generate recursively on a specified path or environment variable and can filter by regex. |
| [filler](https://github.com/yaronsumel/filler) | 14 | 1 | 2017-04-05 | 1 year ago | small utility to fill structs using "fill" tag. |
| [go-httpheader](https://github.com/mozillazg/go-httpheader) | 14 | 2 | 2017-06-24 | 4 months ago | Go library for encoding structs into Header fields. |
| [evaluator](https://github.com/nullne/evaluator) | 13 | 1 | 2017-04-28 | 1 year ago | Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend. |
| [okrun](https://github.com/xta/okrun) | 13 | 2 | 2014-10-01 | 4 years ago | go run error steamroller. |
| [rerate](https://github.com/abo/rerate) | 12 | 4 | 2016-05-24 | 1 year ago | Redis-based rate counter and rate limiter for Go. |
| [gostrutils](https://github.com/ik5/gostrutils) | 12 | 3 | 2018-09-19 | 1 month ago | Collections of string manipulation and conversion functions. |
| [fastlz](https://github.com/digitalcrab/fastlz) | 11 | 1 | 2014-12-26 | 4 years ago | Wrap over [FastLz](http://fastlz.org/) (free, open-source, portable real-time compression library) for GoLang. |
| [structs](https://github.com/PumpkinSeed/structs) | 11 | 1 | 2017-08-26 | 1 year ago | Implement simple functions to manipulate structs. |
| [pgo](https://github.com/arthurkushman/pgo) | 10 | 2 | 2018-12-26 | 1 week ago | Convenient functions for PHP community. |
| [command](https://github.com/txgruppi/command) | 9 | 2 | 2015-08-24 | 2 years ago | Command pattern for Go with thread safe serial and parallel dispatcher. |
| [retry](https://github.com/shafreeck/retry) | 8 | 1 | 2018-07-18 | 6 months ago | A pretty simple library to ensure your work to be done. |
| [backscanner](https://github.com/icza/backscanner) | 7 | 2 | 2017-10-18 | 4 months ago | A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward. |
| [ghokin](https://github.com/antham/ghokin) | 6 | 2 | 2018-08-03 | 4 days ago | Parallelized formatter with no external dependencies for gherkin (cucumber, behat...). |
| [filter](https://github.com/gookit/filter) | 6 | 2 | 2018-09-26 | 2 months ago | provide filtering, sanitizing, and conversion of Go data. |
| [silk](https://github.com/chrispassas/silk) | 4 | 1 | 2018-12-18 | 1 month ago | Read silk netflow files. |
| [mimesniffer](https://github.com/aofei/mimesniffer) | 3 | 1 | 2018-12-20 | 1 week ago | A MIME type sniffer for Go. |
| [sslice](https://github.com/yaa110/sslice) | 2 | 1 | 2018-11-17 | 3 months ago | Create a slice which is always sorted. |
| [retry](https://github.com/percolate/retry) | 2 | 30 | 2018-06-16 | 4 months ago | A simple but highly configurable retry package for Go. |
| [olaf](https://github.com/btnguyen2k/olaf) | 1 | 1 | 2019-01-03 | 1 month ago | Twitter Snowflake implemented in Go. |
| [ctxutil](https://github.com/posener/ctxutil) | 1 | 0 | 2018-07-30 | 1 day ago | A collection of utility functions for contexts. |
| [slicer](https://github.com/leaanthony/slicer) | 1 | 1 | 2019-01-10 | 2 days ago | Makes working with slices easier. |

## UUID
        
*Libraries for working with UUIDs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [uuid](https://github.com/gofrs/uuid) | 440 | 17 | 2018-07-13 | 4 days ago | Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid. |
| [wuid](https://github.com/edwingeng/wuid) | 235 | 13 | 2018-01-27 | 8 months ago | An extremely fast unique number generator, 10-135 times faster than UUID. |
| [Goid](https://github.com/JakeHL/Goid) | 18 | 3 | 2017-05-19 | 2 weeks ago | Generate and Parse RFC4122 compliant V4 UUIDs. |
| [uuid](https://github.com/agext/uuid) | 9 | 1 | 2016-02-03 | 1 week ago | Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier. |

## Validation
        
*Libraries for validation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [govalidator](https://github.com/asaskevich/govalidator) | 3179 | 92 | 2014-06-20 | 1 day ago | Validators and sanitizers for strings, numerics, slices and structs. |
| [validator](https://github.com/go-playground/validator) | 2752 | 64 | 2015-02-13 | 2 weeks ago | Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving. |
| [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) | 909 | 19 | 2016-06-22 | 3 months ago | Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags. |
| [govalidator](https://github.com/thedevsaddam/govalidator) | 569 | 16 | 2017-09-14 | 3 weeks ago | Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. |
| [validate](https://github.com/gookit/validate) | 40 | 3 | 2018-07-16 | 1 day ago | Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features. |
| [validate](https://github.com/gobuffalo/validate) | 11 | 7 | 2018-02-11 | 4 days ago | This package provides a framework for writing validations for Go applications. |

## Version Control
        
*Libraries for version control.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [git2go](https://github.com/libgit2/git2go) | 1279 | 52 | 2013-03-06 | 1 week ago | Go bindings for libgit2. |
| [gh](https://github.com/rjeczalik/gh) | 67 | 5 | 2015-03-09 | 4 months ago | Scriptable server and net/http middleware for GitHub Webhooks. |
| [go-vcs](https://github.com/sourcegraph/go-vcs) | 65 | 28 | 2013-06-02 | 1 month ago | manipulate and inspect VCS repositories in Go. |
| [hgo](https://github.com/beyang/hgo) | 12 | 4 | 2014-06-18 | 3 years ago | Hgo is a collection of Go packages providing read-access to local Mercurial repositories. |

## Video
        
*Libraries for manipulating video.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [goav](https://github.com/giorgisio/goav) | 624 | 40 | 2015-05-21 | 2 months ago | Comphrensive Go bindings for FFmpeg. |
| [gmf](https://github.com/3d0c/gmf) | 462 | 30 | 2013-04-03 | 2 weeks ago | Go bindings for FFmpeg av\* libraries. |
| [go-astits](https://github.com/asticode/go-astits) | 228 | 15 | 2017-07-04 | 1 month ago | Parse and demux MPEG Transport Streams (.ts) natively in GO. |
| [go-astisub](https://github.com/asticode/go-astisub) | 151 | 7 | 2016-12-16 | 2 months ago | Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.). |
| [gst](https://github.com/ziutek/gst) | 148 | 10 | 2011-07-26 | 7 months ago | Go bindings for GStreamer. |
| [libvlc-go](https://github.com/adrg/libvlc-go) | 39 | 5 | 2015-01-06 | 4 months ago | Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player). |
| [go-m3u8](https://github.com/quangngotan95/go-m3u8) | 24 | 1 | 2018-11-06 | 3 months ago | Parser and generator library for Apple m3u8 playlists. |
| [v4l](https://github.com/korandiz/v4l) | 23 | 4 | 2016-10-25 | 9 months ago | Video capture library for Linux, written in Go. |
| [libgosubs](https://github.com/wargarblgarbl/libgosubs) | 11 | 1 | 2017-05-04 | 3 months ago | Subtitle format support for go. Supports .srt, .ttml, and .ass. |

## Web Frameworks
        
*Full stack web frameworks.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gin](https://github.com/gin-gonic/gin) | 24823 | 1025 | 2014-06-17 | 1 hour ago | Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity. |
| [beego](https://github.com/astaxie/beego) | 19130 | 1223 | 2012-02-29 | 1 day ago | beego is an open-source, high-performance web framework for the Go programming language. |
| [echo](https://github.com/labstack/echo) | 13080 | 528 | 2015-03-02 | 22 hours ago | High performance, minimalist Go web framework. |
| [revel](https://github.com/revel/revel) | 10786 | 560 | 2011-12-09 | 3 hours ago | High-productivity web framework for the Go language. |
| [goa](https://github.com/goadesign/goa) | 3293 | 164 | 2014-12-05 | 1 hour ago | Framework for developing microservices based on the design of Ruby's Praxis. |
| [go-json-rest](https://github.com/ant0ine/go-json-rest) | 3263 | 163 | 2013-02-19 | 2 weeks ago | Quick and easy way to setup a RESTful JSON API. |
| [macaron](https://github.com/go-macaron/macaron) | 2676 | 152 | 2014-07-10 | 2 months ago | Macaron is a high productive and modular design web framework in Go. |
| [gizmo](https://github.com/nytimes/gizmo) | 2573 | 107 | 2015-12-16 | 19 hours ago | Microservice toolkit used by the New York Times. |
| [utron](https://github.com/gernest/utron) | 2122 | 74 | 2015-09-16 | 4 months ago | Lightweight MVC framework for Go(Golang). |
| [go-tigertonic](https://github.com/rcrowley/go-tigertonic) | 993 | 45 | 2013-02-10 | 7 months ago | Go framework for building JSON web services inspired by Dropwizard. |
| [tango](https://github.com/lunny/tango) | 774 | 77 | 2014-12-17 | 2 weeks ago | Micro & pluggable web framework for Go. |
| [traffic](https://github.com/gravityblast/traffic) | 512 | 25 | 2013-08-09 | 3 years ago | Sinatra inspired regexp/pattern mux and web framework for Go. |
| [gongular](https://github.com/mustafaakin/gongular) | 410 | 23 | 2016-06-22 | 2 weeks ago | Fast Go web framework with input mapping/validation and (DI) Dependency Injection. |
| [neo](https://github.com/ivpusic/neo) | 387 | 33 | 2015-02-05 | 1 year ago | Neo is minimal and fast Go Web Framework with extremely simple API. |
| [mango](https://github.com/paulbellamy/mango) | 334 | 22 | 2011-05-25 | 1 year ago | Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. |
| [gondola](https://github.com/rainycape/gondola) | 312 | 15 | 2014-07-26 | 2 weeks ago | The web framework for writing faster sites, faster. |
| [air](https://github.com/aofei/air) | 272 | 13 | 2016-07-20 | 1 day ago | An ideally refined web framework for Go. |
| [golf](https://github.com/dinever/golf) | 229 | 19 | 2015-11-18 | 2 years ago | Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library. |
| [gem](https://github.com/go-gem/gem) | 151 | 20 | 2016-11-10 | 1 year ago | Simple and fast web framework, friendly to REST API. |
| [go-rest](https://github.com/ungerik/go-rest) | 111 | 9 | 2012-07-13 | 2 years ago | Small and evil REST framework for Go. |
| [aero](https://github.com/aerogo/aero) | 91 | 13 | 2016-11-09 | 6 hours ago | High-performance web framework for Go, reaches top scores in Lighthouse. |
| [golax](https://github.com/fulldump/golax) | 66 | 9 | 2016-01-31 | 9 months ago | A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more. |
| [webgo](https://github.com/bnkamalesh/webgo) | 63 | 3 | 2015-12-16 | 1 month ago | A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc). |
| [hiboot](https://github.com/hidevopsio/hiboot) | 61 | 6 | 2018-03-16 | 1 week ago | hiboot is a high performance web and cli application framework with dependency injection support |
| [microservice](https://github.com/claygod/microservice) | 51 | 6 | 2016-12-15 | 1 year ago | The framework for the creation of microservices, written in Golang. |
| [yarf](https://github.com/yarf-framework/yarf) | 46 | 3 | 2015-09-02 | 1 year ago | Fast micro-framework designed to build REST APIs and web services in a fast and simple way. |
| [fireball](https://github.com/zpatrick/fireball) | 44 | 4 | 2016-07-20 | 5 months ago | More "natural" feeling web framework. |
| [uadmin](https://github.com/uadmin/uadmin) | 36 | 6 | 2018-10-05 | 2 months ago | Fully featured web framework for Golang, inspired by Django. |
| [api](https://github.com/resoursea/api) | 29 | 5 | 2015-01-25 | 4 years ago | REST framework for quickly writing resource based services. |
| [rex](https://github.com/goanywhere/rex) | 24 | 1 | 2014-10-16 | 1 year ago | Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`. |
| [vox](https://github.com/aisk/vox) | 17 | 1 | 2014-12-24 | 2 months ago | A golang web framework for humans, inspired by Koa heavily. |
| [nio](https://github.com/go-nio/nio) | 14 | 2 | 2018-12-05 | 1 week ago | Modern, minimal and productive Go HTTP framework. |
| [banjo](https://github.com/nsheremet/banjo) | 4 | 1 | 2017-12-09 | 1 year ago | Very simple and fast web framework for Go. |
| [sawsij](https://github.com/jaybill/sawsij) | 2 | 2 | 2016-06-09 | 2 years ago | lightweight, open-source web framework for building high-performance, data-driven web applications. |

### Middlewares
        

#### Actual middlewares
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [tollbooth](https://github.com/didip/tollbooth) | 1082 | 37 | 2015-05-17 | 3 months ago | Rate limit HTTP request handler. |
| [cors](https://github.com/rs/cors) | 1052 | 26 | 2014-10-25 | 2 weeks ago | Easily add CORS capabilities to your API. |
| [go-server-timing](https://github.com/mitchellh/go-server-timing) | 731 | 19 | 2018-02-12 | 6 months ago | Add/parse Server-Timing header. |
| [limiter](https://github.com/ulule/limiter) | 479 | 21 | 2015-10-02 | 1 month ago | Dead simple rate limit middleware for Go. |
| [ln-paywall](https://github.com/philippgille/ln-paywall) | 80 | 2 | 2018-06-30 | 1 week ago | Go middleware for monetizing your API on a per-request basis with Bitcoin and Lightning ⚡️ |
| [xff](https://github.com/sebest/xff) | 69 | 2 | 2014-12-22 | 2 years ago | Handle `X-Forwarded-For` header and friends. |
| [formjson](https://github.com/rs/formjson) | 30 | 2 | 2015-03-20 | 3 years ago | Transparently handle JSON input as a standard form POST. |
| [client-timing](https://github.com/posener/client-timing) | 10 | 1 | 2018-02-23 | 5 days ago | An HTTP client for Server-Timing header. |

#### Libraries for creating HTTP middlewares
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [negroni](https://github.com/urfave/negroni) | 6083 | 236 | 2014-05-19 | 2 weeks ago | Idiomatic HTTP middleware for Golang. |
| [alice](https://github.com/justinas/alice) | 1736 | 54 | 2014-05-25 | 3 months ago | Painless middleware chaining for Go. |
| [render](https://github.com/unrolled/render) | 1198 | 39 | 2014-06-11 | 2 weeks ago | Go package for easily rendering JSON, XML, and HTML template responses. |
| [stats](https://github.com/thoas/stats) | 519 | 16 | 2015-03-06 | 2 weeks ago | Go middleware that stores various information about your web application. |
| [interpose](https://github.com/carbocation/interpose) | 288 | 13 | 2014-07-20 | 2 years ago | Minimalist net/http middleware for golang. |
| [muxchain](https://github.com/stephens2424/muxchain) | 205 | 5 | 2014-05-04 | 1 week ago | Lightweight middleware for net/http. |
| [renderer](https://github.com/thedevsaddam/renderer) | 144 | 5 | 2017-11-08 | 2 months ago | Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go. |
| [rye](https://github.com/InVisionApp/rye) | 90 | 176 | 2016-10-07 | 5 months ago | Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context. |
| [gores](https://github.com/alioygur/gores) | 78 | 5 | 2015-12-25 | 4 months ago | Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs. |
| [chain](https://github.com/codemodus/chain) | 64 | 7 | 2015-05-15 | 6 months ago | Handler wrapper chaining with scoped data (net/context-based "middleware"). |
| [wrap](https://github.com/go-on/wrap) | 55 | 3 | 2014-02-16 | 6 months ago | Small middlewares package for net/http. |
| [catena](https://github.com/codemodus/catena) | 7 | 1 | 2015-07-31 | 6 months ago | http.Handler wrapper catenation (same API as "chain"). |

### Routers
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [httprouter](https://github.com/julienschmidt/httprouter) | 8716 | 286 | 2013-12-05 | 3 weeks ago | High performance router. Use this and the standard http handlers to form a very high performance web framework. |
| [mux](https://github.com/gorilla/mux) | 8263 | 256 | 2012-10-03 | 20 hours ago | Powerful URL router and dispatcher for golang. |
| [chi](https://github.com/go-chi/chi) | 5118 | 151 | 2015-10-16 | 5 hours ago | Small, fast and expressive HTTP router built on net/context. |
| [web](https://github.com/gocraft/web) | 1365 | 55 | 2013-11-17 | 3 weeks ago | Mux and middleware package in Go. |
| [bone](https://github.com/go-zoo/bone) | 1191 | 36 | 2014-11-19 | 1 month ago | Lightning Fast HTTP Multiplexer. |
| [goji](https://github.com/goji/goji) | 712 | 42 | 2015-11-16 | 2 weeks ago | Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`. |
| [fasthttprouter](https://github.com/buaazp/fasthttprouter) | 660 | 32 | 2015-12-13 | 1 month ago | High performance router forked from `httprouter`. The first router fit for `fasthttp`. |
| [gorouter](https://github.com/xujiajun/gorouter) | 385 | 12 | 2018-01-29 | 1 month ago | A simple and fast HTTP router for Go. |
| [lars](https://github.com/go-playground/lars) | 374 | 18 | 2015-12-25 | 1 year ago | Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks. |
| [httptreemux](https://github.com/dimfeld/httptreemux) | 374 | 21 | 2014-05-15 | 2 months ago | High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter. |
| [siesta](https://github.com/VividCortex/siesta) | 348 | 32 | 2014-09-23 | 2 months ago | Composable framework to write middleware and handlers. |
| [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) | 332 | 29 | 2015-10-27 | 3 months ago | An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs. |
| [vestigo](https://github.com/husobee/vestigo) | 246 | 18 | 2015-09-22 | 1 month ago | Performant, stand-alone, HTTP compliant URL Router for go web applications. |
| [router](https://github.com/gowww/router) | 155 | 5 | 2017-05-25 | 11 months ago | Lightning fast HTTP router fully compatible with the net/http.Handler interface. |
| [alien](https://github.com/gernest/alien) | 97 | 3 | 2016-01-31 | 1 year ago | Lightweight and fast http router from outer space. |
| [violetear](https://github.com/nbari/violetear) | 92 | 4 | 2015-06-20 | 5 months ago | Go HTTP router. |
| [Bxog](https://github.com/claygod/Bxog) | 92 | 8 | 2016-05-19 | 2 months ago | Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters. |
| [xmux](https://github.com/rs/xmux) | 87 | 6 | 2015-12-15 | 1 year ago | High performance muxer based on `httprouter` with `net/context` support. |
| [pure](https://github.com/go-playground/pure) | 75 | 4 | 2016-09-24 | 1 year ago | Is a lightweight HTTP router that sticks to the std "net/http" implementation. |
| [gorouter](https://github.com/vardius/gorouter) | 37 | 3 | 2016-07-14 | 4 weeks ago | GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`. |
| [fastrouter](https://github.com/razonyang/fastrouter) | 18 | 2 | 2017-11-01 | 1 year ago | a fast, flexible HTTP router written in Go. |

## Windows
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [go-ole](https://github.com/go-ole/go-ole) | 488 | 35 | 2011-01-21 | 5 hours ago | Win32 OLE implementation for golang. |
| [d3d9](https://github.com/gonutz/d3d9) | 80 | 5 | 2015-12-13 | 2 months ago | Go bindings for Direct3D9. |

## XML
        
*Libraries and tools for manipulating XML.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [xquery](https://github.com/antchfx/xquery) | 141 | 11 | 2016-10-09 | 9 months ago | XQuery lets you extract data from HTML/XML documents using XPath expression. |
| [xpath](https://github.com/antchfx/xpath) | 121 | 8 | 2016-10-09 | 1 day ago | XPath package for Go. |
| [XML-Comp](https://github.com/XML-Comp/XML-Comp) | 15 | 1 | 2016-10-26 | 7 months ago | Simple command line XML comparer that generates diffs of folders, files and tags. |
| [xml2map](https://github.com/sbabiv/xml2map) | 13 | 0 | 2018-08-07 | 1 month ago | XML to MAP converter written Golang. |
| [xmlwriter](https://github.com/shabbyrobe/xmlwriter) | 6 | 1 | 2017-04-11 | 1 month ago | Procedural XML generation API based on libxml2's xmlwriter module. |

# Tools
        
*Go software and plugins.*

## Code Analysis
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gometalinter](https://github.com/alecthomas/gometalinter) | 3549 | 80 | 2014-08-05 | 5 days ago | Metalinter is a tool to automatically apply all static analysis tool and report their output in normalized form. |
| [lint](https://github.com/golang/lint) | 2887 | 100 | 2013-06-03 | 3 days ago | Golint is a linter for Go source code. |
| [go-tools](https://github.com/dominikh/go-tools) | 2224 | 66 | 2017-01-25 | 2 weeks ago | Staticcheck – a collection of static analysis tools for working with Go code |
| [errcheck](https://github.com/kisielk/errcheck) | 1235 | 20 | 2013-02-25 | 2 months ago | Errcheck is a program for checking for unchecked errors in Go programs. |
| [gcvis](https://github.com/davecheney/gcvis) | 878 | 34 | 2014-07-10 | 1 year ago | Visualise Go program GC trace data in real time. |
| [php-parser](https://github.com/z7zmey/php-parser) | 578 | 27 | 2017-11-07 | 1 week ago | A Parser for PHP written in Go. |
| [go-critic](https://github.com/go-critic/go-critic) | 483 | 18 | 2018-05-05 | 3 days ago | source code linter that brings checks that are currently not implemented in other linters. |
| [goast-viewer](https://github.com/yuroyoro/goast-viewer) | 328 | 13 | 2014-06-30 | 1 year ago | Web based Golang AST visualizer. |
| [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) | 247 | 7 | 2017-04-13 | 9 months ago | go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects. |
| [unconvert](https://github.com/mdempsky/unconvert) | 244 | 9 | 2016-02-20 | 1 month ago | Remove unnecessary type conversions from Go source. |
| [gostatus](https://github.com/shurcooL/gostatus) | 233 | 6 | 2013-11-27 | 4 weeks ago | Command line tool, shows the status of repositories that contain Go packages. |
| [apicompat](https://github.com/bradleyfalzon/apicompat) | 161 | 7 | 2016-07-10 | 2 years ago | Checks recent changes to a Go project for backwards incompatible changes. |
| [dupl](https://github.com/mibk/dupl) | 145 | 8 | 2015-05-20 | 4 months ago | Tool for code clone detection. |
| [checkstyle](https://github.com/qiniu/checkstyle) | 92 | 10 | 2014-01-01 | 3 months ago | checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style refered to some points in Go Code Review Comments. |
| [validate](https://github.com/mccoyst/validate) | 61 | 6 | 2013-11-22 | 2 years ago | Automatically validates struct fields with tags. |
| [lint](https://github.com/surullabs/lint) | 61 | 6 | 2016-07-09 | 4 months ago | Run linters as part of go test. |
| [go-outdated](https://github.com/firstrow/go-outdated) | 46 | 1 | 2015-06-29 | 1 month ago | Console application that displays outdated packages. |
| [blanket](https://github.com/verygoodsoftwarenotvirus/blanket) | 11 | 2 | 2017-09-04 | 7 months ago | tarp finds functions and methods without direct unit tests in Go source code. |

## Editor Plugins
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [vim-go](https://github.com/fatih/vim-go) | 9941 | 279 | 2014-03-24 | 4 days ago | Go development plugin for Vim. |
| [gocode](https://github.com/nsf/gocode) | 4601 | 203 | 2010-07-05 | 3 days ago | Autocompletion daemon for the Go programming language. |
| [vscode-go](https://github.com/Microsoft/vscode-go) | 4558 | 211 | 2015-10-15 | 13 hours ago | Extension for Visual Studio Code (VS Code) which provides support for the Go language. |
| [GoSublime](https://github.com/DisposaBoy/GoSublime) | 3136 | 115 | 2011-08-28 | 1 week ago | Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features. |
| [go-plus](https://github.com/joefitzgerald/go-plus) | 1451 | 48 | 2014-03-14 | 7 hours ago | Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting. |
| [go-mode.el](https://github.com/dominikh/go-mode.el) | 883 | 58 | 2013-01-31 | 1 month ago | Go mode for GNU/Emacs. |
| [Watch](https://github.com/eaburns/Watch) | 157 | 11 | 2013-08-09 | 11 months ago | Runs a command in an acme win on file changes. |
| [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) | 78 | 5 | 2012-11-26 | 2 years ago | Vim plugin to highlight syntax errors on save. |
| [go-language-server](https://github.com/theia-ide/go-language-server) | 27 | 4 | 2017-11-21 | 6 days ago | A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol. |
| [gounit-vim](https://github.com/hexdigest/gounit-vim) | 16 | 2 | 2018-02-22 | 4 months ago | Vim plugin for generating Go tests based on the function's or method's signature. |
| [velour](https://github.com/velour/velour) | 16 | 8 | 2014-06-27 | 4 months ago | IRC client for the acme editor. |
| [theia-go-extension](https://github.com/theia-ide/theia-go-extension) | 10 | 4 | 2017-11-30 | 6 days ago | Go language support for the Theia IDE. |

## Go Generate Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gotests](https://github.com/cweill/gotests) | 1852 | 56 | 2016-01-19 | 4 days ago | Generate Go tests from your source code. |
| [genny](https://github.com/cheekybits/genny) | 811 | 19 | 2014-10-28 | 4 days ago | Elegant generics for Go. |
| [re2dfa](https://github.com/opennota/re2dfa) | 162 | 10 | 2015-06-20 | 5 months ago | Transform regular expressions into finite state machines and output Go source code. |
| [gocontracts](https://github.com/Parquery/gocontracts) | 45 | 4 | 2018-08-14 | 1 month ago | brings design-by-contract to Go by synchronizing the code with the documentation. |
| [generic](https://github.com/usk81/generic) | 25 | 2 | 2016-06-15 | 6 months ago | flexible data type for Go. |
| [gounit](https://github.com/hexdigest/gounit) | 21 | 4 | 2018-02-05 | 6 months ago | Generate Go tests using your own templates. |

## Go Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [OctoLinker](https://github.com/OctoLinker/OctoLinker) | 3346 | 85 | 2013-12-28 | 3 days ago | Navigate through go files efficiently with the OctoLinker browser extension for GitHub. |
| [go-swagger](https://github.com/go-swagger/go-swagger) | 3295 | 97 | 2014-11-17 | 4 hours ago | Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API. |
| [go-callvis](https://github.com/TrueFurby/go-callvis) | 1675 | 58 | 2016-09-03 | 1 month ago | Visualize call graph of your Go program using dot format. |
| [richgo](https://github.com/kyoh86/richgo) | 339 | 4 | 2017-01-05 | 1 week ago | Enrich `go test` outputs with text decorations. |
| [depth](https://github.com/KyleBanks/depth) | 312 | 7 | 2017-03-04 | 2 days ago | Visualize dependency trees of any package by analyzing imports. |
| [rts](https://github.com/galeone/rts) | 180 | 3 | 2016-04-04 | 2 years ago | RTS: response to struct. Generates Go structs from server responses. |
| [godbg](https://github.com/tylerwince/godbg) | 147 | 4 | 2019-01-24 | 1 month ago | Implementation of Rusts `dbg!` macro for quick and easy debugging during development. |
| [colorgo](https://github.com/songgao/colorgo) | 95 | 5 | 2013-02-15 | 2 years ago | Wrapper around `go` command for colorized `go build` output. |
| [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) | 37 | 2 | 2015-05-22 | 1 year ago | Bash completion for go and wgo. |
| [generator-go-lang](https://github.com/axelspringer/generator-go-lang) | 10 | 12 | 2017-09-13 | 1 year ago | A [Yeoman](http://yeoman.io) generator to get new Go projects started. |

## Software Packages
        
*Software written in Go.*

### DevOps Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [moby](https://github.com/moby/moby) | 52387 | 3287 | 2013-01-19 | 9 hours ago | Collaborative project for the container ecosystem to assemble container-based systems. |
| [kubernetes](https://github.com/kubernetes/kubernetes) | 49092 | 2913 | 2014-06-07 | 45 minutes ago | Container Cluster Manager from Google. |
| [traefik](https://github.com/containous/traefik) | 20764 | 623 | 2015-09-14 | 7 hours ago | Reverse proxy and load balancer with support for multiple backends. |
| [gitea](https://github.com/go-gitea/gitea) | 12655 | 424 | 2016-11-01 | 2 hours ago | Fork of Gogs, entirely community driven. |
| [vegeta](https://github.com/tsenart/vegeta) | 10799 | 279 | 2013-08-13 | 1 hour ago | HTTP load testing tool and library. It's over 9000! |
| [packer](https://github.com/hashicorp/packer) | 8614 | 422 | 2013-03-23 | 5 hours ago | Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. |
| [hey](https://github.com/rakyll/hey) | 5066 | 118 | 2016-09-02 | 3 weeks ago | Hey is a tiny program that sends some load to a web application. |
| [gvm](https://github.com/moovweb/gvm) | 4146 | 142 | 2011-12-03 | 2 months ago | GVM provides an interface to manage Go versions. |
| [webhook](https://github.com/adnanh/webhook) | 3589 | 132 | 2015-01-13 | 4 days ago | Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server. |
| [gaia](https://github.com/gaia-pipeline/gaia) | 3515 | 94 | 2017-12-28 | 41 minutes ago | Build powerful pipelines in any programming language. |
| [gox](https://github.com/mitchellh/gox) | 3119 | 69 | 2013-11-17 | 6 days ago | Dead simple, no frills Go cross compile tool. |
| [bosun](https://github.com/bosun-monitor/bosun) | 2787 | 153 | 2013-11-15 | 1 week ago | Time Series Alerting Framework. |
| [aptly](https://github.com/aptly-dev/aptly) | 1674 | 97 | 2013-12-14 | 1 month ago | aptly is a Debian repository management tool. |
| [goxc](https://github.com/laher/goxc) | 1601 | 45 | 2013-02-11 | 8 months ago | build tool for Go, with a focus on cross-compiling and packaging. |
| [fac](https://github.com/mkchoi212/fac) | 1542 | 31 | 2017-12-30 | 5 months ago | Command-line user interface to fix git merge conflicts. |
| [bombardier](https://github.com/codesenberg/bombardier) | 1445 | 38 | 2016-05-29 | 1 day ago | Fast cross-platform HTTP benchmarking tool. |
| [kala](https://github.com/ajvb/kala) | 1304 | 62 | 2015-03-19 | 3 weeks ago | Simplistic, modern, and performant job scheduler. |
| [statusok](https://github.com/sanathp/statusok) | 1063 | 43 | 2015-08-27 | 1 month ago | Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected. |
| [banshee](https://github.com/eleme/banshee) | 1036 | 56 | 2015-12-03 | 4 months ago | Anomalies detection system for periodic metrics. |
| [s3gof3r](https://github.com/rlmcpherson/s3gof3r) | 969 | 32 | 2013-08-02 | 2 months ago | Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3. |
| [go-selfupdate](https://github.com/sanbornm/go-selfupdate) | 629 | 23 | 2013-11-13 | 2 years ago | Enable your Go applications to self update. |
| [skm](https://github.com/TimothyYe/skm) | 494 | 20 | 2017-10-11 | 6 days ago | SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily! |
| [scaleway-cli](https://github.com/scaleway/scaleway-cli) | 489 | 50 | 2015-03-20 | 19 hours ago | Manage BareMetal Servers from Command Line (as easily as with Docker). |
| [aurora](https://github.com/xuri/aurora) | 362 | 26 | 2016-10-09 | 1 month ago | Cross-platform web-based Beanstalkd queue server console. |
| [govvv](https://github.com/ahmetb/govvv) | 339 | 11 | 2016-08-03 | 10 months ago | “go build” wrapper to easily add version information into Go binaries. |
| [gonative](https://github.com/inconshreveable/gonative) | 301 | 8 | 2014-05-01 | 2 years ago | Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages. |
| [mora](https://github.com/emicklei/mora) | 260 | 23 | 2013-07-12 | 2 years ago | REST server for accessing MongoDB documents and meta data. |
| [godbg](https://github.com/sirnewton01/godbg) | 217 | 17 | 2013-08-09 | 7 months ago | Web-based gdb front-end application. |
| [lstags](https://github.com/ivanilves/lstags) | 212 | 13 | 2017-08-15 | 1 month ago | Tool and API to sync Docker images across different registries. |
| [dogo](https://github.com/liudng/dogo) | 205 | 18 | 2014-11-19 | 2 years ago | Monitoring changes in the source file and automatically compile and run (restart). |
| [manssh](https://github.com/xwjdsh/manssh) | 190 | 4 | 2017-10-08 | 8 months ago | manssh is a command line tool for managing your ssh alias config easily. |
| [pomerium](https://github.com/pomerium/pomerium) | 187 | 12 | 2019-01-01 | 1 day ago | Pomerium is an identity-aware access proxy. |
| [pewpew](https://github.com/bengadbois/pewpew) | 186 | 7 | 2016-10-13 | 4 months ago | Flexible HTTP command line stress tester. |
| [gobrew](https://github.com/cryptojuice/gobrew) | 173 | 5 | 2013-11-13 | 1 year ago | gobrew lets you easily switch between multiple versions of go. |
| [ostent](https://github.com/ostrost/ostent) | 161 | 5 | 2014-03-31 | 11 months ago | collects and displays system metrics and optionally relays to Graphite and/or InfluxDB. |
| [blast](https://github.com/dave/blast) | 160 | 3 | 2017-10-22 | 1 year ago | A simple tool for API load testing and batch jobs. |
| [grapes](https://github.com/yaronsumel/grapes) | 116 | 6 | 2016-09-01 | 4 weeks ago | Lightweight tool designed to distribute commands over ssh with ease. |
| [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) | 78 | 2 | 2017-03-03 | 3 months ago | Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`. |
| [kcli](https://github.com/cswank/kcli) | 59 | 6 | 2017-03-26 | 1 month ago | Command line tool for inspecting kafka topics/partitions/messages. |
| [winrm-cli](https://github.com/masterzen/winrm-cli) | 55 | 5 | 2016-05-23 | 1 month ago | Cli tool to remotely execute commands on Windows machines. |
| [go-furnace](https://github.com/go-furnace/go-furnace) | 52 | 2 | 2016-10-09 | 2 months ago | Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean. |
| [drone-scp](https://github.com/appleboy/drone-scp) | 45 | 2 | 2016-10-16 | 23 hours ago | Copy files and artifacts via SSH using a binary, docker or Drone CI. |
| [dropship](https://github.com/ChrisMcKenzie/dropship) | 43 | 3 | 2015-09-04 | 7 months ago | Tool for deploying code via cdn. |
| [rodent](https://github.com/alouche/rodent) | 30 | 2 | 2014-06-02 | 1 year ago | Rodent helps you manage Go versions, projects and track dependencies. |
| [drone-jenkins](https://github.com/appleboy/drone-jenkins) | 20 | 1 | 2016-10-15 | 23 hours ago | Trigger downstream Jenkins jobs using a binary, docker or Drone CI. |
| [awsenv](https://github.com/soniah/awsenv) | 18 | 1 | 2015-08-05 | 7 months ago | Small binary that loads Amazon (AWS) environment variables for a profile. |
| [depcharge](https://github.com/centerorbit/depcharge) | 8 | 3 | 2018-07-25 | 3 months ago | Helps orchestrating the execution of commands across the many dependencies in larger projects. |
| [lwc](https://github.com/timdp/lwc) | 7 | 2 | 2018-04-22 | 8 months ago | A live-updating version of the UNIX wc command. |
| [sg](https://github.com/ChristopherRabotin/sg) | 5 | 1 | 2015-08-19 | 2 years ago | Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response. |

### Other Software
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [goreplay](https://github.com/buger/goreplay) | 10483 | 441 | 2013-05-30 | 1 week ago | Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time. |
| [rkt](https://github.com/rkt/rkt) | 8476 | 452 | 2014-11-12 | 1 day ago | App Container runtime that integrates with init systems, is compatible with other container formats like Docker, and supports alternative execution engines like KVM. |
| [seaweedfs](https://github.com/chrislusf/seaweedfs) | 7405 | 471 | 2014-07-15 | 16 hours ago | Fast, Simple and Scalable Distributed File System with O(1) disk seek. |
| [restic](https://github.com/restic/restic) | 6183 | 213 | 2014-04-27 | 2 days ago | De-duplicating backup program. |
| [comcast](https://github.com/tylertreat/comcast) | 5974 | 150 | 2014-11-12 | 4 months ago | Simulate bad network connections. |
| [confd](https://github.com/kelseyhightower/confd) | 5910 | 237 | 2013-10-01 | 1 week ago | Manage local application configuration files using templates and data from etcd or consul. |
| [liteide](https://github.com/visualfc/liteide) | 5133 | 372 | 2012-11-19 | 1 day ago | LiteIDE is a simple, open source, cross-platform Go IDE. |
| [drive](https://github.com/odeke-em/drive) | 4711 | 212 | 2014-11-03 | 4 weeks ago | Google Drive client for the commandline. |
| [nes](https://github.com/fogleman/nes) | 3961 | 139 | 2015-03-03 | 1 month ago | Nintendo Entertainment System (NES) emulator written in Go. |
| [toxiproxy](https://github.com/Shopify/toxiproxy) | 3529 | 261 | 2014-09-04 | 1 week ago | Proxy to simulate network and system conditions for automated tests. |
| [duplicacy](https://github.com/gilbertchen/duplicacy) | 2543 | 88 | 2016-02-23 | 17 hours ago | A cross-platform network and cloud backup tool based on the idea of lock-free deduplication. |
| [pipe](https://github.com/b3log/pipe) | 2150 | 88 | 2017-09-11 | 1 day ago | A small and beautiful blogging platform. |
| [mylg](https://github.com/mehrdadrad/mylg) | 2141 | 104 | 2016-06-22 | 4 months ago | Command Line Network Diagnostic tool written in Go. |
| [goboy](https://github.com/Humpheh/goboy) | 1994 | 37 | 2017-08-20 | 3 weeks ago | Nintendo Game Boy Color emulator written in Go. |
| [sup](https://github.com/pressly/sup) | 1883 | 75 | 2015-02-24 | 2 months ago | Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers. |
| [snap](https://github.com/intelsdi-x/snap) | 1791 | 154 | 2014-08-14 | 2 months ago | Powerful telemetry framework. |
| [circuit](https://github.com/gocircuit/circuit) | 1741 | 139 | 2014-04-11 | 2 years ago | Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications. |
| [lgo](https://github.com/yunabe/lgo) | 1613 | 42 | 2017-10-05 | 4 days ago | Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility. |
| [borg](https://github.com/ok-borg/borg) | 1403 | 43 | 2016-09-11 | 1 year ago | Terminal based search engine for bash snippets. |
| [Go-Package-Store](https://github.com/shurcooL/Go-Package-Store) | 876 | 22 | 2014-01-24 | 2 months ago | App that displays updates for the Go packages in your GOPATH. |
| [scc](https://github.com/boyter/scc) | 652 | 15 | 2018-03-01 | 1 day ago | Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates. |
| [leaps](https://github.com/Jeffail/leaps) | 631 | 26 | 2014-06-20 | 2 days ago | Pair programming service using Operational Transforms. |
| [community](https://github.com/documize/community) | 611 | 40 | 2016-04-30 | 1 hour ago | Modern wiki software that integrates data from SaaS tools. |
| [peg](https://github.com/pointlander/peg) | 552 | 30 | 2010-04-26 | 1 day ago | Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator. |
| [vflow](https://github.com/VerizonDigital/vflow) | 526 | 72 | 2017-02-25 | 3 weeks ago | High-performance, scalable and reliable IPFIX, sFlow and Netflow collector. |
| [mockingjay-server](https://github.com/quii/mockingjay-server) | 393 | 14 | 2015-04-05 | 3 months ago | Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests. |
| [godns](https://github.com/TimothyYe/godns) | 349 | 19 | 2014-05-11 | 6 days ago | A dynamic DNS client tool, supports DNSPod & HE.net, written in Go. |
| [shell2http](https://github.com/msoap/shell2http) | 328 | 18 | 2015-03-12 | 3 days ago | Executing shell commands via http server (for prototyping or remote control). |
| [gocc](https://github.com/goccmack/gocc) | 311 | 22 | 2015-06-05 | 2 days ago | Gocc is a compiler kit for Go written in Go. |
| [wellington](https://github.com/wellington/wellington) | 286 | 13 | 2014-12-09 | 4 months ago | Sass project management tool, extends the language with sprite functions (like Compass). |
| [ipe](https://github.com/dimiro1/ipe) | 251 | 14 | 2015-01-13 | 2 months ago | Open source Pusher server implementation compatible with Pusher client libraries written in GO. |
| [IDE](https://github.com/thestrukture/IDE) | 243 | 14 | 2017-09-10 | 2 months ago | Browser accessible IDE. Designed for Go with Go. |
| [cherry](https://github.com/rafael-santiago/cherry) | 185 | 12 | 2015-10-25 | 1 year ago | Tiny webchat server in Go. |
| [orange-cat](https://github.com/utatti/orange-cat) | 174 | 6 | 2014-11-01 | 1 month ago | Markdown previewer written in Go. |
| [orbit](https://github.com/gulien/orbit) | 126 | 9 | 2017-05-13 | 8 months ago | A simple tool for running commands and generating files from templates. |
| [joincap](https://github.com/assafmo/joincap) | 116 | 7 | 2018-06-01 | 1 week ago | Command-line utility for merging multiple pcap files together. |
| [ddns](https://github.com/skibish/ddns) | 79 | 5 | 2017-03-14 | 3 months ago | Personal DDNS client with Digital Ocean Networking DNS as backend. |
| [boxed](https://github.com/tejo/boxed) | 72 | 2 | 2015-04-19 | 6 months ago | Dropbox based blog engine. |
| [naclpipe](https://github.com/unix4fun/naclpipe) | 20 | 5 | 2015-05-06 | 3 months ago | Simple NaCL EC25519 based crypto pipe tool written in Go. |
| [snitch](https://github.com/lucasgomide/snitch) | 15 | 1 | 2017-04-07 | 7 months ago | Simple way to notify your team and many tools when someone has deployed any application via Tsuru. |
| [term-quiz](https://github.com/crazcalm/term-quiz) | 15 | 1 | 2017-12-26 | 4 months ago | Quizzes for your terminal. |
| [GoDocTooltip](https://github.com/diankong/GoDocTooltip) | 12 | 3 | 2016-01-21 | 3 years ago | Chrome extension for Go Doc sites, which shows function description as tooltip at function list. |

# Server Applications
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [etcd](https://github.com/etcd-io/etcd) | 23278 | 1228 | 2013-07-07 | 20 hours ago | Highly-available key value store for shared configuration and service discovery. |
| [caddy](https://github.com/mholt/caddy) | 20788 | 668 | 2015-01-14 | 11 hours ago | Caddy is an alternative, HTTP/2 web server that's easy to configure and use. |
| [minio](https://github.com/minio/minio) | 14951 | 415 | 2015-01-15 | 2 hours ago | Minio is a distributed object storage server. |
| [devd](https://github.com/cortesi/devd) | 2732 | 68 | 2015-09-28 | 4 weeks ago | Local webserver for developers. |
| [roadrunner](https://github.com/spiral/roadrunner) | 2621 | 104 | 2017-12-27 | 1 week ago | High-performance PHP application server, load-balancer and process manager. |
| [fider](https://github.com/getfider/fider) | 669 | 14 | 2017-01-18 | 15 hours ago | Fider is an open platform to collect and organize customer feedback. |
| [jackal](https://github.com/ortuman/jackal) | 637 | 38 | 2017-11-14 | 2 days ago | An XMPP server written in Go. |
| [flagr](https://github.com/checkr/flagr) | 555 | 45 | 2017-10-04 | 4 days ago | Flagr is an open-source feature flagging and A/B testing service. |
| [algernon](https://github.com/xyproto/algernon) | 523 | 36 | 2015-03-10 | 1 week ago | HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber. |
| [discovery](https://github.com/bilibili/discovery) | 349 | 28 | 2018-04-20 | 1 week ago | A registry for resilient mid-tier load balancing and failover. |
| [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) | 4 | 1 | 2018-10-23 | 1 month ago | Nginx log parser and exporter to Prometheus. |

# Resources
        
*Where to discover new Go libraries.*

## Benchmarks
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) | 1204 | 54 | 2013-12-17 | 10 months ago | Go HTTP request router benchmark and comparison. |
| [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) | 897 | 75 | 2016-04-06 | 1 month ago | Go web framework benchmark. |
| [skynet](https://github.com/atemerev/skynet) | 888 | 46 | 2016-02-14 | 4 months ago | Skynet 1M threads microbenchmark. |
| [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) | 766 | 34 | 2013-01-19 | 1 month ago | Benchmarks of Go serialization methods. |
| [speedtest-resize](https://github.com/fawick/speedtest-resize) | 163 | 8 | 2013-09-16 | 2 years ago | Compare various Image resize algorithms for the Go language. |
| [go-benchmarks](https://github.com/tylertreat/go-benchmarks) | 112 | 9 | 2016-02-25 | 3 years ago | Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches. |
| [autobench](https://github.com/davecheney/autobench) | 89 | 8 | 2013-03-28 | 4 years ago | Framework to compare the performance between different Go versions. |
| [gospeed](https://github.com/feyeleanor/gospeed) | 88 | 7 | 2011-05-24 | 5 months ago | Go micro-benchmarks for calculating the speed of language constructs. |
| [gocostmodel](https://github.com/mna/gocostmodel) | 51 | 5 | 2014-12-19 | 4 years ago | Benchmarks of common basic operations for the Go language. |
| [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) | 48 | 5 | 2014-09-25 | 11 months ago | Collection of benchmarks for popular Go database/SQL utilities. |
| [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) | 19 | 1 | 2017-01-24 | 1 year ago | Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results. |
| [golang-micro-benchmarks](https://github.com/amscanne/golang-micro-benchmarks) | 16 | 2 | 2014-04-05 | 3 days ago | Tiny collection of Go micro benchmarks. The intent is to compare some language features to others. |
| [kvbench](https://github.com/jimrobinson/kvbench) | 14 | 1 | 2014-04-15 | 4 years ago | Key/Value database benchmark. |

## Conferences
        

## E-Books
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [GoBooks](https://github.com/dariubs/GoBooks) | 6098 | 487 | 2015-05-05 | 1 week ago | A curated list of Go books. |
| [web-dev-golang-anti-textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook) | 2220 | 90 | 2016-01-01 | 1 week ago | Learn how to write webapps without a framework in Go. |

## Gophers
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gophers](https://github.com/ashleymcnamara/gophers) | 1614 | 82 | 2017-02-15 | 4 days ago | Gopher artworks by Ashley McNamara. |
| [gophers](https://github.com/egonelbre/gophers) | 1425 | 30 | 2015-06-03 | 1 month ago | Free gophers. |
| [gophericons](https://github.com/shalakhin/gophericons) | 554 | 20 | 2015-08-22 | 11 months ago | 34 gopher images for Go developers community |
| [gopher-stickers](https://github.com/tenntenn/gopher-stickers) | 404 | 14 | 2014-11-10 | 2 years ago | gopher stickers |
| [gopher-vector](https://github.com/golang-samples/gopher-vector) | 331 | 13 | 2013-03-31 | 2 years ago | Vector data of gopher |
| [gopherize.me](https://github.com/matryer/gopherize.me) | 269 | 6 | 2017-01-25 | 4 months ago | Gopherize yourself. |
| [gopher-logos](https://github.com/GolangUA/gopher-logos) | 59 | 7 | 2017-07-27 | 8 months ago | adorable gopher logos. |
| [gophers](https://github.com/rogeralsing/gophers) | 49 | 2 | 2017-01-29 | 1 year ago | random gopher graphics. |
| [go-gopher](https://github.com/sillecelik/go-gopher) | 35 | 0 | 2018-03-29 | 1 week ago | Gopher amigurumi toy pattern. |
| [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) | 29 | 1 | 2014-09-04 | 1 year ago | Go gopher Vector Data [.ai, .svg]. |

## Meetups
        
*Add the group of your city/country here (send **PR**)*

## Twitter
        

## Websites
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [go](https://github.com/golang/go) | 54003 | 3259 | 2014-08-19 | 1 hour ago | List of projects on the Go community wiki. |
| [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) | 23790 | 1710 | 2014-07-08 | 5 days ago | List of other amazingly awesome lists. |
| [awesome-remote-job](https://github.com/lukasz-madon/awesome-remote-job) | 13589 | 789 | 2015-01-02 | 6 days ago | Curated list of awesome remote jobs. A lot of them are looking for Go hackers. |
| [golang-graphics](https://github.com/mholt/golang-graphics) | 141 | 10 | 2014-03-25 | 3 years ago | Collection of Go images, graphics, and art. |
| [gocryforhelp](https://github.com/ninedraft/gocryforhelp) | 32 | 10 | 2016-05-09 | 1 year ago | Collection of Go projects that needs help. Good place to start your open-source way in Go. |

### Tutorials
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang) | 28260 | 2316 | 2012-08-02 | 2 days ago | Golang ebook intro how to build a web app with golang. |
| [go-lang-cheat-sheet](https://github.com/a8m/go-lang-cheat-sheet) | 3761 | 155 | 2014-02-13 | 2 weeks ago | Go's reference card. |
| [learn-go-with-tests](https://github.com/quii/learn-go-with-tests) | 3089 | 124 | 2018-03-02 | 6 hours ago | Learn Go with test-driven development. |
| [working-with-go](https://github.com/mkaz/working-with-go) | 1087 | 50 | 2014-05-05 | 2 weeks ago | Intro to go for experienced programmers. |
| [ethereum-development-with-go-book](https://github.com/miguelmota/ethereum-development-with-go-book) | 359 | 26 | 2018-05-16 | 4 days ago | A little e-book on Ethereum Development with Go. |
| [golang-for-nodejs-developers](https://github.com/miguelmota/golang-for-nodejs-developers) | 240 | 10 | 2019-01-03 | 2 days ago | Examples of Golang compared to Node.js for learning. |