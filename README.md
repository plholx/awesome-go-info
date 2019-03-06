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
| [portaudio](https://github.com/gordonklaus/portaudio) | 266 | 9 | 2015-09-16 | 6 months ago | Go bindings for the PortAudio audio I/O library. |
| [music-theory](https://github.com/go-music-theory/music-theory) | 233 | 10 | 2016-03-17 | 8 months ago | Music theory models in Go. |
| [waveform](https://github.com/mdlayher/waveform) | 229 | 13 | 2014-09-14 | 2 years ago | Go package capable of generating waveform images from audio streams. |
| [portmidi](https://github.com/rakyll/portmidi) | 188 | 6 | 2013-11-10 | 1 year ago | Go bindings for PortMidi. |
| [flac](https://github.com/mewkiz/flac) | 92 | 6 | 2012-11-02 | 1 week ago | Native Go FLAC encoder/decoder with support for FLAC streams. |
| [id3v2](https://github.com/bogem/id3v2) | 90 | 5 | 2016-05-16 | 3 weeks ago | Fast and stable ID3 parsing and writing library for Go. |
| [mix](https://github.com/go-mix/mix) | 90 | 3 | 2016-01-03 | 1 year ago | Sequence-based Go-native audio mixer for music apps. |
| [go-sox](https://github.com/krig/go-sox) | 87 | 8 | 2013-10-08 | 8 months ago | libsox bindings for go. |
| [mp3](https://github.com/tcolgate/mp3) | 85 | 1 | 2015-02-27 | 1 year ago | Native Go MP3 decoder. |
| [flac](https://github.com/eaburns/flac) | 81 | 5 | 2013-08-21 | 1 year ago | No-frills native Go FLAC decoder that decodes FLAC files into byte slices. |
| [go-taglib](https://github.com/wtolson/go-taglib) | 64 | 6 | 2012-11-20 | 7 months ago | Go bindings for taglib. |
| [malgo](https://github.com/gen2brain/malgo) | 55 | 4 | 2017-11-10 | 1 month ago | Mini audio library. |
| [gaad](https://github.com/Comcast/gaad) | 51 | 10 | 2016-07-11 | 1 year ago | Native Go AAC bitstream parser. |
| [go_mediainfo](https://github.com/zhulik/go_mediainfo) | 22 | 1 | 2015-12-14 | 3 years ago | libmediainfo bindings for go. |
| [vorbis](https://github.com/mccoyst/vorbis) | 21 | 3 | 2013-07-12 | 1 year ago | "Native" Go Vorbis decoder (uses CGO, but has no dependencies). |
| [minimp3](https://github.com/tosone/minimp3) | 19 | 2 | 2018-01-26 | 4 days ago | Lightweight MP3 decoder library. |
| [EasyMIDI](https://github.com/algoGuy/EasyMIDI) | 17 | 3 | 2018-02-20 | 11 months ago | EasyMidi is a simple and reliable library for working with standard midi file (SMF). |
| [gosamplerate](https://github.com/dh1tw/gosamplerate) | 8 | 1 | 2016-11-21 | 9 months ago | libsamplerate bindings for go. |

## Authentication and OAuth
        
*Libraries for implementing authentications schemes.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [jwt-go](https://github.com/dgrijalva/jwt-go) | 5035 | 130 | 2012-04-18 | 1 week ago | Golang implementation of JSON Web Tokens (JWT). |
| [casbin](https://github.com/casbin/casbin) | 3883 | 139 | 2017-04-08 | 3 weeks ago | Authorization library that supports access control models like ACL, RBAC, ABAC. |
| [casbin](https://github.com/casbin/casbin) | 3874 | 140 | 2017-04-08 | 3 weeks ago | Authorization library that supports access control models like ACL, RBAC, ABAC. |
| [casbin](https://github.com/casbin/casbin) | 3865 | 140 | 2017-04-08 | 3 weeks ago | Authorization library that supports access control models like ACL, RBAC, ABAC. |
| [casbin](https://github.com/casbin/casbin) | 3864 | 140 | 2017-04-08 | 3 weeks ago | Authorization library that supports access control models like ACL, RBAC, ABAC. |
| [oauth2](https://github.com/golang/oauth2) | 2116 | 92 | 2014-04-14 | 1 week ago | Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support. |
| [goth](https://github.com/markbates/goth) | 2083 | 58 | 2014-10-15 | 1 week ago | provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box. |
| [authboss](https://github.com/volatiletech/authboss) | 1746 | 38 | 2015-01-03 | 3 weeks ago | Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time. |
| [osin](https://github.com/openshift/osin) | 1509 | 69 | 2013-09-11 | 4 months ago | Golang OAuth2 server library. |
| [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) | 1121 | 71 | 2015-11-01 | 1 week ago | Standalone, specification-compliant,  OAuth2 server written in Golang. |
| [go-jose](https://github.com/square/go-jose) | 1016 | 61 | 2014-11-15 | 1 week ago | Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs. |
| [gologin](https://github.com/dghubble/gologin) | 961 | 27 | 2015-06-23 | 1 day ago | chainable handlers for login with OAuth1 and OAuth2 authentication providers. |
| [gorbac](https://github.com/mikespook/gorbac) | 833 | 55 | 2013-12-26 | 1 month ago | provides a lightweight role-based access control (RBAC) implementation in Golang. |
| [loginsrv](https://github.com/tarent/loginsrv) | 735 | 45 | 2016-11-11 | 4 days ago | JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam. |
| [permissions2](https://github.com/xyproto/permissions2) | 301 | 12 | 2014-11-19 | 1 week ago | Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt. |
| [paseto](https://github.com/o1egl/paseto) | 195 | 9 | 2018-01-23 | 3 months ago | Golang implementation of Platform-Agnostic Security Tokens (PASETO). |
| [httpauth](https://github.com/goji/httpauth) | 167 | 5 | 2014-05-27 | 2 years ago | HTTP Authentication middleware. |
| [jwt-auth](https://github.com/adam-hanna/jwt-auth) | 144 | 9 | 2016-07-06 | 1 month ago | JWT middleware for Golang http servers with many configuration options. |
| [session](https://github.com/icza/session) | 82 | 6 | 2016-02-08 | 5 months ago | Go session management for web servers (including support for Google App Engine - GAE). |
| [jwt](https://github.com/robbert229/jwt) | 60 | 6 | 2016-06-06 | 3 months ago | Clean and easy to use implementation of JSON Web Tokens (JWT). |
| [branca](https://github.com/hako/branca) | 50 | 5 | 2018-01-09 | 7 months ago | Golang implementation of Branca Tokens. |
| [sessions](https://github.com/adam-hanna/sessions) | 41 | 3 | 2017-04-29 | 1 year ago | Dead simple, highly performant, highly customizable sessions service for go http servers. |
| [jwt](https://github.com/pascaldekloe/jwt) | 40 | 2 | 2018-03-21 | 1 day ago | Lightweight JSON Web Token (JWT) library. |
| [securecookie](https://github.com/chmike/securecookie) | 26 | 4 | 2017-09-03 | 6 months ago | Efficient secure cookie encoding/decoding. |
| [rbac](https://github.com/zpatrick/rbac) | 20 | 3 | 2018-08-02 | 6 months ago | Minimalistic RBAC package for Go applications. |
| [signedvalue](https://github.com/sashka/signedvalue) | 8 | 0 | 2018-01-06 | 1 month ago | Signed and timestamped strings compatible with [Tornado's](https://github.com/tornadoweb/tornado) `create_signed_value`, `decode_signed_value`, and therefore `set_secure_cookie` and `get_secure_cookie`. |
| [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) | 6 | 2 | 2017-10-20 | 3 months ago | Go session management using the SessionGate Redis module. |
| [cookiestxt](https://github.com/mengzhuo/cookiestxt) | 2 | 1 | 2017-10-09 | 1 year ago | provides parser of cookies.txt file format. |

## Bot Building
        
*Libraries for building and working with bots.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) | 1389 | 56 | 2015-06-25 | 4 days ago | Simple and clean Telegram bot client. |
| [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) | 1387 | 56 | 2015-06-25 | 4 days ago | Simple and clean Telegram bot client. |
| [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) | 1385 | 56 | 2015-06-25 | 4 days ago | Simple and clean Telegram bot client. |
| [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) | 1385 | 56 | 2015-06-25 | 4 days ago | Simple and clean Telegram bot client. |
| [telebot](https://github.com/tucnak/telebot) | 828 | 33 | 2015-06-26 | 3 days ago | Telegram bot framework written in Go. |
| [bot](https://github.com/go-chat-bot/bot) | 414 | 34 | 2015-09-23 | 2 weeks ago | IRC, Slack & Telegram bot written in Go. |
| [slacker](https://github.com/shomali11/slacker) | 256 | 13 | 2017-05-20 | 1 week ago | Easy to use framework to create Slack bots. |
| [tbot](https://github.com/yanzay/tbot) | 184 | 8 | 2015-09-12 | 2 months ago | Telegram bot server with API similar to net/http. |
| [tenyks](https://github.com/kyleterry/tenyks) | 166 | 14 | 2012-08-26 | 2 years ago | Service oriented IRC bot using Redis and JSON for messaging. |
| [golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot) | 148 | 19 | 2017-05-15 | 3 weeks ago | A golang implementation of a console-based trading bot for cryptocurrency exchanges. |
| [go-sarah](https://github.com/oklahomer/go-sarah) | 111 | 3 | 2016-11-06 | 7 months ago | Framework to build bot for desired chat services including LINE, Slack, Gitter and more. |
| [hanu](https://github.com/sbstjn/hanu) | 90 | 6 | 2016-09-16 | 6 months ago | Framework for writing Slack bots. |
| [go-tgbot](https://github.com/olebedev/go-tgbot) | 80 | 7 | 2016-12-11 | 8 months ago | Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware. |
| [margelet](https://github.com/zhulik/margelet) | 51 | 5 | 2015-11-21 | 2 years ago | Framework for building Telegram bots. |
| [govkbot](https://github.com/nikepan/govkbot) | 18 | 2 | 2016-07-12 | 19 hours ago | Simple Go [VK](https://vk.com) bot library. |
| [micha](https://github.com/onrik/micha) | 9 | 3 | 2016-04-14 | 1 year ago | Go Library for Telegram bot api. |

## Command Line
        
*Libraries for building Console Applications and Console User Interfaces.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [cobra](https://github.com/spf13/cobra) | 10799 | 255 | 2013-09-04 | 1 week ago | Commander for modern Go CLI interactions. |
| [cli](https://github.com/urfave/cli) | 10103 | 271 | 2013-07-14 | 3 weeks ago | Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli). |
| [termui](https://github.com/gizak/termui) | 8365 | 275 | 2015-02-03 | 7 hours ago | Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib). |
| [gocui](https://github.com/jroimartin/gocui) | 3974 | 102 | 2014-01-04 | 2 days ago | Minimalist Go library aimed at creating Console User Interfaces. |
| [termbox-go](https://github.com/nsf/termbox-go) | 3294 | 103 | 2012-01-13 | 3 weeks ago | Termbox is a library for creating cross-platform text-based interfaces. |
| [color](https://github.com/fatih/color) | 2865 | 56 | 2014-02-17 | 4 months ago | Versatile package for colored terminal output. |
| [kingpin](https://github.com/alecthomas/kingpin) | 2302 | 54 | 2014-05-15 | 14 hours ago | Command line and flag parser supporting sub commands. |
| [go-prompt](https://github.com/c-bata/go-prompt) | 2082 | 35 | 2017-08-15 | 1 week ago | Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit). |
| [readline](https://github.com/chzyer/readline) | 1307 | 39 | 2015-09-20 | 3 weeks ago | Pure golang implementation that provides most features in GNU-Readline under MIT license. |
| [go-flags](https://github.com/jessevdk/go-flags) | 1294 | 28 | 2012-08-31 | 1 week ago | go command line option parser. |
| [uiprogress](https://github.com/gosuri/uiprogress) | 1229 | 27 | 2015-11-17 | 4 months ago | Flexible library to render progress bars in terminal applications. |
| [docopt.go](https://github.com/docopt/docopt.go) | 1039 | 36 | 2013-08-26 | 5 months ago | Command-line arguments parser that will make you smile. |
| [asciigraph](https://github.com/guptarohit/asciigraph) | 1023 | 23 | 2018-06-17 | 1 month ago | Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. |
| [cli](https://github.com/mitchellh/cli) | 930 | 25 | 2013-11-03 | 3 months ago | Go library for implementing command-line interfaces. |
| [gcli](https://github.com/tcnksm/gcli) | 856 | 25 | 2014-06-19 | 1 year ago | The easy way to start building Golang command line applications. |
| [uilive](https://github.com/gosuri/uilive) | 724 | 11 | 2015-11-16 | 3 months ago | Library for updating terminal output in realtime. |
| [pflag](https://github.com/spf13/pflag) | 628 | 24 | 2013-08-30 | 3 weeks ago | Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. |
| [mow.cli](https://github.com/jawher/mow.cli) | 590 | 19 | 2014-12-19 | 1 week ago | Go library for building CLI applications with sophisticated flag and argument parsing and validation. |
| [go-arg](https://github.com/alexflint/go-arg) | 571 | 15 | 2015-11-01 | 2 months ago | Struct-based argument parsing in Go. |
| [complete](https://github.com/posener/complete) | 542 | 12 | 2017-05-06 | 2 weeks ago | Write bash completions in Go + Go command bash completion. |
| [liner](https://github.com/peterh/liner) | 522 | 22 | 2012-08-16 | 16 hours ago | Go readline-like library for command-line interfaces. |
| [progressbar](https://github.com/schollz/progressbar) | 493 | 13 | 2017-10-27 | 1 month ago | Basic thread-safe progress bar that works in every OS. |
| [uitable](https://github.com/gosuri/uitable) | 464 | 15 | 2015-11-14 | 1 year ago | Library to improve readability in terminal apps using tabular data. |
| [aurora](https://github.com/logrusorgru/aurora) | 429 | 5 | 2016-11-07 | 5 months ago | ANSI terminal colors that supports fmt.Printf/Sprintf. |
| [cli](https://github.com/mkideal/cli) | 427 | 20 | 2016-02-27 | 6 days ago | Feature-rich and easy to use command-line package based on golang struct tags. |
| [mpb](https://github.com/vbauerster/mpb) | 419 | 8 | 2016-12-14 | 1 week ago | Multi progress bar for terminal applications. |
| [flaggy](https://github.com/integrii/flaggy) | 410 | 11 | 2018-03-05 | 2 weeks ago | A robust and idiomatic flags package with excellent subcommand support. |
| [go-colorable](https://github.com/mattn/go-colorable) | 336 | 16 | 2014-07-30 | 1 week ago | Colorable writer for windows. |
| [go-isatty](https://github.com/mattn/go-isatty) | 299 | 7 | 2014-04-01 | 1 week ago | isatty for golang. |
| [chalk](https://github.com/ttacon/chalk) | 286 | 11 | 2014-07-19 | 2 years ago | Intuitive package for prettifying terminal/console output. |
| [termtables](https://github.com/apcera/termtables) | 203 | 68 | 2012-12-06 | 1 year ago | Go port of the Ruby library [terminal-tables](https://github.com/tj/terminal-table) for simple ASCII table generation as well as providing markdown and HTML output. |
| [go-colortext](https://github.com/daviddengcn/go-colortext) | 188 | 9 | 2013-01-23 | 11 months ago | Go library for color output in terminals. |
| [color](https://github.com/gookit/color) | 130 | 4 | 2018-07-01 | 1 week ago | Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows. |
| [simpletable](https://github.com/alexeyco/simpletable) | 121 | 2 | 2017-03-29 | 1 week ago | Simple tables in terminal with Go. |
| [clif](https://github.com/ukautz/clif) | 95 | 2 | 2015-05-31 | 2 weeks ago | Small command line interface framework. |
| [flag](https://github.com/cosiner/flag) | 90 | 5 | 2016-10-06 | 1 week ago | Simple but powerful command line option parsing library for Go supporting subcommand. |
| [termdash](https://github.com/mum4k/termdash) | 83 | 6 | 2018-03-24 | 2 days ago | Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui). |
| [argparse](https://github.com/akamensky/argparse) | 74 | 5 | 2017-11-24 | 1 month ago | Command line argument parser inspired by Python's argparse module. |
| [sflags](https://github.com/octago/sflags) | 71 | 5 | 2016-12-04 | 1 month ago | Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries. |
| [commandeer](https://github.com/jaffee/commandeer) | 70 | 4 | 2017-10-12 | 4 days ago | Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags. |
| [wmenu](https://github.com/dixonwille/wmenu) | 67 | 2 | 2016-04-20 | 1 week ago | Easy to use menu structure for cli applications that prompts users to make choices. |
| [cfmt](https://github.com/mingrammer/cfmt) | 53 | 3 | 2018-03-16 | 2 months ago | Contextual fmt inspired by bootstrap color classes. |
| [cli](https://github.com/teris-io/cli) | 43 | 1 | 2017-05-25 | 8 months ago | Simple and complete API for building command line interfaces in Go. |
| [env](https://github.com/codingconcepts/env) | 37 | 1 | 2017-06-15 | 4 months ago | Tag-based environment configuration for structs. |
| [wlog](https://github.com/dixonwille/wlog) | 29 | 1 | 2016-04-14 | 1 year ago | Simple logging interface that supports cross-platform color and concurrency. |
| [gocmd](https://github.com/devfacet/gocmd) | 28 | 3 | 2018-01-08 | 6 months ago | Go library for building command line applications. |
| [flagvar](https://github.com/sgreben/flagvar) | 26 | 2 | 2018-05-19 | 4 months ago | A collection of flag argument types for Go's standard `flag` package. |
| [tabular](https://github.com/InVisionApp/tabular) | 23 | 3 | 2018-04-24 | 9 months ago | Print ASCII tables from command line utilities without the need to pass large sets of data to the API. |
| [strumt](https://github.com/antham/strumt) | 22 | 0 | 2017-06-20 | 4 days ago | Library to create prompt chain. |
| [colourize](https://github.com/TreyBastian/colourize) | 15 | 3 | 2015-05-11 | 2 years ago | Go library for ANSI colour text in terminals. |
| [argv](https://github.com/cosiner/argv) | 11 | 1 | 2017-01-22 | 1 month ago | Go library to split command line string as arguments array using the bash syntax. |
| [go-commander](https://github.com/yitsushi/go-commander) | 8 | 1 | 2016-10-10 | 1 week ago | Go library to simplify CLI workflow. |
| [go-ataman](https://github.com/workanator/go-ataman) | 8 | 2 | 2017-05-18 | 1 year ago | Go library for rendering ANSI colored text templates in terminals. |
| [ctc](https://github.com/wzshiming/ctc) | 6 | 1 | 2018-04-28 | 4 months ago | The non-invasive cross-platform terminal color library does not need to modify the Print method. |
| [go-getoptions](https://github.com/DavidGamba/go-getoptions) | 5 | 1 | 2015-12-18 | 3 days ago | Go option parser inspired on the flexibility of Perl’s GetOpt::Long. |
| [sand](https://github.com/Zaba505/sand) | 2 | 1 | 2018-11-19 | 3 months ago | Simple API for creating interpreters and so much more. |

## Configuration
        
*Libraries for configuration parsing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [viper](https://github.com/spf13/viper) | 7678 | 165 | 2014-04-02 | 1 hour ago | Go configuration with fangs. |
| [envconfig](https://github.com/kelseyhightower/envconfig) | 2131 | 37 | 2013-11-07 | 2 weeks ago | Go library for managing configuration data from environment variables. |
| [godotenv](https://github.com/joho/godotenv) | 1752 | 30 | 2013-07-30 | 1 week ago | Go port of Ruby's dotenv library (Loads environment variables from `.env`). |
| [ini](https://github.com/go-ini/ini) | 1335 | 62 | 2014-12-18 | 2 weeks ago | Go package to read and write INI files. |
| [env](https://github.com/caarlos0/env) | 704 | 16 | 2015-07-28 | 3 days ago | Parse environment variables to Go structs (with defaults). |
| [store](https://github.com/tucnak/store) | 239 | 5 | 2015-10-04 | 1 year ago | Lightweight configuration manager for Go. |
| [confita](https://github.com/heetch/confita) | 209 | 25 | 2017-12-21 | 7 hours ago | Load configuration in cascade from multiple backends into a struct. |
| [config](https://github.com/joshbetz/config) | 193 | 3 | 2017-04-03 | 1 year ago | Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP. |
| [config](https://github.com/olebedev/config) | 190 | 8 | 2014-04-21 | 5 months ago | JSON or YAML configuration wrapper with environment variables and flags parsing. |
| [hjson-go](https://github.com/hjson/hjson-go) | 163 | 7 | 2016-08-06 | 3 weeks ago | Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments. |
| [envconfig](https://github.com/vrischmann/envconfig) | 136 | 4 | 2015-04-22 | 2 months ago | Read your configuration from environment variables. |
| [gcfg](https://github.com/go-gcfg/gcfg) | 102 | 4 | 2015-08-17 | 9 months ago | read INI-style configuration files into Go structs; supports user-defined types and subsections. |
| [goconfig](https://github.com/crgimenes/goconfig) | 93 | 11 | 2016-12-18 | 4 weeks ago | Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. |
| [goconfig](https://github.com/crgimenes/goconfig) | 93 | 11 | 2016-12-18 | 4 weeks ago | Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. |
| [goconfig](https://github.com/crgimenes/goconfig) | 92 | 11 | 2016-12-18 | 4 weeks ago | Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. |
| [goconfig](https://github.com/crgimenes/goconfig) | 92 | 11 | 2016-12-18 | 4 weeks ago | Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. |
| [envh](https://github.com/antham/envh) | 92 | 3 | 2017-01-12 | 4 days ago | Helpers to manage environment variables. |
| [envcfg](https://github.com/tomazk/envcfg) | 89 | 1 | 2014-11-29 | 1 year ago | Un-marshaling environment variables to Go structs. |
| [gofigure](https://github.com/ian-kent/gofigure) | 56 | 6 | 2014-11-25 | 1 year ago | Go application configuration made easy. |
| [config](https://github.com/gookit/config) | 46 | 3 | 2018-07-07 | 1 month ago | application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge. |
| [configure](https://github.com/paked/configure) | 44 | 3 | 2015-06-14 | 2 weeks ago | Provides configuration through multiple sources, including JSON, flags and environment variables. |
| [ingo](https://github.com/schachmat/ingo) | 22 | 1 | 2016-02-08 | 1 year ago | Flags persisted in an ini-like config file. |
| [go-up](https://github.com/ufoscout/go-up) | 21 | 1 | 2018-02-18 | 2 days ago | A simple configuration library with recursive placeholders resolution and no magic. |
| [xdg](https://github.com/OpenPeeDeeP/xdg) | 19 | 3 | 2017-07-20 | 1 week ago | Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html). |
| [mini](https://github.com/sasbury/mini) | 19 | 2 | 2015-04-30 | 2 months ago | Golang package for parsing ini-style configuration files. |
| [envconf](https://github.com/ian-kent/envconf) | 7 | 1 | 2014-10-26 | 4 years ago | Configuration from environment. |
| [conflate](https://github.com/the4thamigo-uk/conflate) | 5 | 0 | 2018-02-02 | 3 weeks ago | Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema. |
| [sprbox](https://github.com/oblq/sprbox) | 3 | 2 | 2018-07-17 | 4 months ago | Build-environment aware toolbox factory and agnostic config parser (YAML, TOML, JSON and Environment vars). |

## Continuous Integration
        
*Tools for help with continuous integration.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [drone](https://github.com/drone/drone) | 17377 | 587 | 2014-02-07 | 4 hours ago | Drone is a Continuous Integration platform built on Docker, written in Go. |
| [goveralls](https://github.com/mattn/goveralls) | 540 | 16 | 2013-04-17 | 1 month ago | Go integration for Coveralls.io continuous code coverage tracking system. |
| [overalls](https://github.com/go-playground/overalls) | 94 | 4 | 2015-07-30 | 6 months ago | Multi-Package go project coverprofile for tools like goveralls. |
| [duci](https://github.com/duck8823/duci) | 32 | 3 | 2018-04-01 | 2 days ago | A simple ci server no needs domain specific languages. |
| [gomason](https://github.com/nikogura/gomason) | 22 | 0 | 2017-11-18 | 3 weeks ago | Test, Build, Sign, and Publish your go binaries from a clean workspace. |
| [roveralls](https://github.com/lawrencewoodman/roveralls) | 11 | 1 | 2016-06-26 | 1 year ago | Recursive coverage testing tool. |
| [roveralls](https://github.com/lawrencewoodman/roveralls) | 11 | 1 | 2016-06-26 | 1 year ago | Recursive coverage testing tool. |
| [roveralls](https://github.com/lawrencewoodman/roveralls) | 11 | 1 | 2016-06-26 | 1 year ago | Recursive coverage testing tool. |
| [roveralls](https://github.com/lawrencewoodman/roveralls) | 11 | 1 | 2016-06-26 | 1 year ago | Recursive coverage testing tool. |

## CSS Preprocessors
        
*Libraries for preprocessing CSS files.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gcss](https://github.com/yosssi/gcss) | 409 | 16 | 2014-09-04 | 4 years ago | Pure Go CSS Preprocessor. |
| [go-libsass](https://github.com/wellington/go-libsass) | 116 | 4 | 2015-04-19 | 3 weeks ago | Go wrapper to the 100% Sass compatible libsass project. |

## Data Structures
        
*Generic datastructures and algorithms in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gods](https://github.com/emirpasic/gods) | 5356 | 255 | 2015-03-04 | 3 days ago | Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc. |
| [go-datastructures](https://github.com/Workiva/go-datastructures) | 4788 | 292 | 2014-10-29 | 4 hours ago | Collection of useful, performant, and thread-safe data structures. |
| [BoomFilters](https://github.com/tylertreat/BoomFilters) | 1082 | 36 | 2015-02-06 | 4 months ago | Probabilistic data structures for processing continuous, unbounded streams. |
| [golang-set](https://github.com/deckarep/golang-set) | 978 | 27 | 2013-07-04 | 5 months ago | Thread-Safe and Non-Thread-Safe high-performance sets for Go. |
| [gota](https://github.com/go-gota/gota) | 739 | 57 | 2016-02-07 | 1 week ago | Implementation of dataframes, series, and data wrangling methods for Go. |
| [gota](https://github.com/go-gota/gota) | 738 | 57 | 2016-02-07 | 1 week ago | Implementation of dataframes, series, and data wrangling methods for Go. |
| [gota](https://github.com/go-gota/gota) | 737 | 57 | 2016-02-07 | 1 week ago | Implementation of dataframes, series, and data wrangling methods for Go. |
| [gota](https://github.com/go-gota/gota) | 737 | 57 | 2016-02-07 | 1 week ago | Implementation of dataframes, series, and data wrangling methods for Go. |
| [hyperloglog](https://github.com/axiomhq/hyperloglog) | 634 | 18 | 2017-06-18 | 2 days ago | HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction. |
| [bloom](https://github.com/willf/bloom) | 591 | 30 | 2011-05-21 | 5 days ago | Go package implementing Bloom filters. |
| [roaring](https://github.com/RoaringBitmap/roaring) | 569 | 34 | 2014-07-11 | 1 month ago | Go package implementing compressed bitsets. |
| [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) | 447 | 17 | 2015-06-29 | 3 days ago | Cuckoo filter: a good alternative to a counting bloom filter implemented in Go. |
| [bitset](https://github.com/willf/bitset) | 441 | 27 | 2011-05-11 | 5 days ago | Go package implementing bitsets. |
| [trie](https://github.com/derekparker/trie) | 363 | 20 | 2014-03-07 | 9 months ago | Trie implementation in Go. |
| [go-geoindex](https://github.com/hailocab/go-geoindex) | 297 | 67 | 2015-01-22 | 1 year ago | In-memory geo index. |
| [mafsa](https://github.com/smartystreets/mafsa) | 272 | 18 | 2014-11-08 | 2 years ago | MA-FSA implementation with Minimal Perfect Hashing. |
| [goskiplist](https://github.com/ryszard/goskiplist) | 178 | 12 | 2012-05-09 | 2 years ago | Skip list implementation in Go. |
| [hilbert](https://github.com/google/hilbert) | 171 | 19 | 2015-08-06 | 3 months ago | Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves. |
| [algorithms](https://github.com/shady831213/algorithms) | 155 | 6 | 2018-01-31 | 1 month ago | Algorithms and data structures.CLRS study. |
| [bloom](https://github.com/zentures/bloom) | 125 | 8 | 2013-09-03 | 10 months ago | Bloom filters implemented in Go. |
| [bloom](https://github.com/zentures/bloom) | 125 | 8 | 2013-09-03 | 10 months ago | Bloom filters implemented in Go. |
| [bloom](https://github.com/zentures/bloom) | 125 | 8 | 2013-09-03 | 10 months ago | Bloom filters implemented in Go. |
| [bloom](https://github.com/zentures/bloom) | 125 | 8 | 2013-09-03 | 10 months ago | Bloom filters implemented in Go. |
| [merkletree](https://github.com/cbergoon/merkletree) | 119 | 4 | 2017-04-12 | 2 days ago | Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures. |
| [binpacker](https://github.com/zhuangsirui/binpacker) | 110 | 10 | 2016-02-02 | 8 months ago | Binary packer and unpacker helps user build custom binary stream. |
| [go-rquad](https://github.com/arl/go-rquad) | 94 | 3 | 2016-09-13 | 8 months ago | Region quadtrees with efficient point location and neighbour finding. |
| [go-rquad](https://github.com/arl/go-rquad) | 94 | 3 | 2016-09-13 | 8 months ago | Region quadtrees with efficient point location and neighbour finding. |
| [go-rquad](https://github.com/arl/go-rquad) | 94 | 3 | 2016-09-13 | 8 months ago | Region quadtrees with efficient point location and neighbour finding. |
| [go-rquad](https://github.com/arl/go-rquad) | 94 | 3 | 2016-09-13 | 8 months ago | Region quadtrees with efficient point location and neighbour finding. |
| [encoding](https://github.com/zentures/encoding) | 92 | 7 | 2013-09-21 | 1 year ago | Integer Compression Libraries for Go. |
| [encoding](https://github.com/zentures/encoding) | 92 | 7 | 2013-09-21 | 1 year ago | Integer Compression Libraries for Go. |
| [encoding](https://github.com/zentures/encoding) | 92 | 7 | 2013-09-21 | 1 year ago | Integer Compression Libraries for Go. |
| [encoding](https://github.com/zentures/encoding) | 92 | 7 | 2013-09-21 | 1 year ago | Integer Compression Libraries for Go. |
| [ring](https://github.com/TheTannerRyan/ring) | 84 | 1 | 2019-01-27 | 2 weeks ago | Go implementation of a high performance, thread safe bloom filter. |
| [skiplist](https://github.com/MauriceGit/skiplist) | 80 | 5 | 2018-06-24 | 2 months ago | Very fast Go Skiplist implementation. |
| [ttlcache](https://github.com/ReneKroon/ttlcache) | 78 | 7 | 2014-12-13 | 1 week ago | In-memory LRU string-interface{} map with expiration for golang. |
| [ttlcache](https://github.com/ReneKroon/ttlcache) | 78 | 7 | 2014-12-13 | 1 week ago | In-memory LRU string-interface{} map with expiration for golang. |
| [ttlcache](https://github.com/ReneKroon/ttlcache) | 78 | 7 | 2014-12-13 | 1 week ago | In-memory LRU string-interface{} map with expiration for golang. |
| [ttlcache](https://github.com/ReneKroon/ttlcache) | 78 | 7 | 2014-12-13 | 1 week ago | In-memory LRU string-interface{} map with expiration for golang. |
| [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) | 64 | 5 | 2016-04-01 | 3 months ago | Go implementation of Adaptive Radix Tree. |
| [conjungo](https://github.com/InVisionApp/conjungo) | 59 | 14 | 2016-12-30 | 3 weeks ago | A small, powerful and flexible merge library. |
| [skiplist](https://github.com/gansidui/skiplist) | 58 | 7 | 2014-11-19 | 4 years ago | Skiplist implementation in Go. |
| [levenshtein](https://github.com/agnivade/levenshtein) | 42 | 3 | 2014-07-30 | 6 days ago | Implementation to calculate levenshtein distance in Go. |
| [count-min-log](https://github.com/seiflotfy/count-min-log) | 40 | 4 | 2015-08-17 | 2 years ago | Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory). |
| [deque](https://github.com/gammazero/deque) | 36 | 4 | 2018-04-24 | 1 month ago | Fast ring-buffer deque (double-ended queue). |
| [bit](https://github.com/yourbasic/bit) | 29 | 3 | 2017-05-04 | 11 months ago | Golang set data structure with bonus bit-twiddling functions. |
| [levenshtein](https://github.com/agext/levenshtein) | 25 | 1 | 2016-04-08 | 1 week ago | Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. |
| [go-mcache](https://github.com/OrlovEvgeny/go-mcache) | 23 | 2 | 2018-04-15 | 3 months ago | Fast in-memory key:value store/cache library. Pointer caches. |
| [concurrent-writer](https://github.com/free/concurrent-writer) | 16 | 4 | 2017-09-18 | 1 year ago | Highly concurrent drop-in replacement for `bufio.Writer`. |
| [bloom](https://github.com/yourbasic/bloom) | 13 | 1 | 2017-05-07 | 1 year ago | Golang Bloom filter implementation. |
| [goset](https://github.com/zoumo/goset) | 12 | 1 | 2017-08-25 | 1 year ago | A useful Set collection implementation for Go. |
| [pipeline](https://github.com/hyfather/pipeline) | 9 | 1 | 2018-04-25 | 6 months ago | An implementation of pipelines with fan-in and fan-out. |
| [go-ef](https://github.com/amallia/go-ef) | 8 | 1 | 2017-09-22 | 1 year ago | A Go implementation of the Elias-Fano encoding. |
| [set](https://github.com/StudioSol/set) | 5 | 17 | 2018-07-21 | 4 months ago | Simple set data structure implementation in Go using LinkedHashMap. |
| [mspm](https://github.com/BlackRabbitt/mspm) | 4 | 3 | 2018-05-18 | 9 months ago | Multi-String Pattern Matching Algorithm for information retrieval. |
| [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) | 2 | 1 | 2019-01-11 | 1 month ago | Concurrent FIFO queue. |
| [null](https://github.com/emvi/null) | 2 | 2 | 2018-07-05 | 1 week ago | Nullable Go types that can be marshalled/unmarshalled to/from JSON. |
| [hide](https://github.com/emvi/hide) | 2 | 2 | 2019-01-16 | 1 week ago | ID type with marshalling to/from hash to prevent sending IDs to clients. |
| [deque](https://github.com/edwingeng/deque) | 1 | 1 | 2019-02-01 | 3 weeks ago | A highly optimized double-ended queue. |
| [timedmap](https://github.com/zekroTJA/timedmap) | 0 | 1 | 2019-01-30 | 1 week ago | Map with expiring key-value pairs. |

## Database
        
*SQL query builder, libraries for building and using SQL.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [prometheus](https://github.com/prometheus/prometheus) | 22446 | 1006 | 2012-11-24 | 1 hour ago | Monitoring system and time series database. |
| [tidb](https://github.com/pingcap/tidb) | 17592 | 1196 | 2015-09-06 | 39 minutes ago | TiDB is a distributed SQL database. Inspired by the design of Google F1. |
| [influxdb](https://github.com/influxdata/influxdb) | 15715 | 749 | 2013-09-26 | 57 minutes ago | Scalable datastore for metrics, events, and real-time analytics. |
| [influxdb](https://github.com/influxdata/influxdb) | 15707 | 749 | 2013-09-26 | 1 day ago | Scalable datastore for metrics, events, and real-time analytics. |
| [influxdb](https://github.com/influxdata/influxdb) | 15705 | 749 | 2013-09-26 | 2 days ago | Scalable datastore for metrics, events, and real-time analytics. |
| [influxdb](https://github.com/influxdata/influxdb) | 15705 | 748 | 2013-09-26 | 1 day ago | Scalable datastore for metrics, events, and real-time analytics. |
| [cockroach](https://github.com/cockroachdb/cockroach) | 15544 | 695 | 2014-02-06 | 7 hours ago | Scalable, Geo-Replicated, Transactional Datastore. |
| [bolt](https://github.com/boltdb/bolt) | 9511 | 338 | 2013-12-21 | 1 year ago | Low-level key/value database for Go. |
| [dgraph](https://github.com/dgraph-io/dgraph) | 8747 | 329 | 2015-08-25 | 4 hours ago | Scalable, Distributed, Low Latency, High Throughput Graph Database. |
| [vitess](https://github.com/vitessio/vitess) | 7618 | 464 | 2013-06-28 | 3 hours ago | vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. |
| [vitess](https://github.com/vitessio/vitess) | 7613 | 464 | 2013-06-28 | 1 day ago | vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. |
| [vitess](https://github.com/vitessio/vitess) | 7611 | 464 | 2013-06-28 | 1 day ago | vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. |
| [vitess](https://github.com/vitessio/vitess) | 7608 | 464 | 2013-06-28 | 2 days ago | vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. |
| [groupcache](https://github.com/golang/groupcache) | 7105 | 443 | 2013-07-23 | 3 weeks ago | Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. |
| [pgweb](https://github.com/sosedoff/pgweb) | 5761 | 152 | 2014-10-09 | 1 week ago | Web-based PostgreSQL database browser. |
| [badger](https://github.com/dgraph-io/badger) | 5395 | 220 | 2017-01-26 | 17 hours ago | Fast key-value store in Go. |
| [rqlite](https://github.com/rqlite/rqlite) | 4260 | 182 | 2014-08-23 | 1 week ago | The lightweight, distributed, relational database built on SQLite. |
| [kingshard](https://github.com/flike/kingshard) | 4194 | 388 | 2015-07-04 | 2 weeks ago | kingshard is a high performance proxy for MySQL powered by Golang. |
| [ledisdb](https://github.com/siddontang/ledisdb) | 2890 | 184 | 2014-04-30 | 4 weeks ago | Ledisdb is a high performance NoSQL like Redis based on LevelDB. |
| [goleveldb](https://github.com/syndtr/goleveldb) | 2805 | 152 | 2013-01-23 | 1 day ago | Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go. |
| [orchestrator](https://github.com/github/orchestrator) | 2607 | 182 | 2016-11-30 | 2 hours ago | MySQL replication topology manager & visualizer. |
| [go-cache](https://github.com/patrickmn/go-cache) | 2430 | 84 | 2012-01-02 | 1 week ago | In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. |
| [go-cache](https://github.com/patrickmn/go-cache) | 2424 | 84 | 2012-01-02 | 1 week ago | In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. |
| [go-cache](https://github.com/patrickmn/go-cache) | 2422 | 84 | 2012-01-02 | 1 week ago | In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. |
| [go-cache](https://github.com/patrickmn/go-cache) | 2422 | 84 | 2012-01-02 | 1 week ago | In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. |
| [tiedot](https://github.com/HouzuoGuo/tiedot) | 2307 | 161 | 2013-05-26 | 1 month ago | Your NoSQL database powered by Golang. |
| [buntdb](https://github.com/tidwall/buntdb) | 2273 | 94 | 2016-07-20 | 1 month ago | Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support. |
| [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) | 2039 | 149 | 2015-01-15 | 2 days ago | Sync your MySQL data into Elasticsearch automatically. |
| [prest](https://github.com/prest/prest) | 1963 | 82 | 2016-11-22 | 6 months ago | Serve a RESTful API from any PostgreSQL database. |
| [prest](https://github.com/prest/prest) | 1963 | 82 | 2016-11-22 | 6 months ago | Serve a RESTful API from any PostgreSQL database. |
| [prest](https://github.com/prest/prest) | 1962 | 82 | 2016-11-22 | 6 months ago | Serve a RESTful API from any PostgreSQL database. |
| [prest](https://github.com/prest/prest) | 1962 | 82 | 2016-11-22 | 6 months ago | Serve a RESTful API from any PostgreSQL database. |
| [xo](https://github.com/xo/xo) | 1949 | 69 | 2016-02-05 | 3 days ago | Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. |
| [xo](https://github.com/xo/xo) | 1947 | 69 | 2016-02-05 | 3 days ago | Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. |
| [xo](https://github.com/xo/xo) | 1946 | 70 | 2016-02-05 | 3 days ago | Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. |
| [xo](https://github.com/xo/xo) | 1946 | 70 | 2016-02-05 | 3 days ago | Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. |
| [squirrel](https://github.com/Masterminds/squirrel) | 1944 | 33 | 2014-01-18 | 1 week ago | Go library that helps you build SQL queries. |
| [bigcache](https://github.com/allegro/bigcache) | 1881 | 64 | 2016-03-23 | 2 weeks ago | Efficient key/value cache for gigabytes of data. |
| [migrate](https://github.com/golang-migrate/migrate) | 1727 | 37 | 2018-01-19 | 4 days ago | Database migrations. CLI and Golang library. |
| [go-mysql](https://github.com/siddontang/go-mysql) | 1571 | 132 | 2014-02-21 | 2 days ago | Go toolset to handle MySQL protocol and replication. |
| [sql-migrate](https://github.com/rubenv/sql-migrate) | 1247 | 29 | 2014-09-09 | 2 weeks ago | Database migration tool. Allows embedding migrations into the application using go-bindata. |
| [cache2go](https://github.com/muesli/cache2go) | 817 | 58 | 2013-11-11 | 1 day ago | In-memory key:value cache which supports automatic invalidation based on timeouts. |
| [diskv](https://github.com/peterbourgon/diskv) | 695 | 29 | 2012-03-22 | 6 months ago | Home-grown disk-backed key-value store. |
| [moss](https://github.com/couchbase/moss) | 686 | 79 | 2016-02-07 | 22 hours ago | Moss is a simple LSM key-value storage engine written in 100% Go. |
| [gcache](https://github.com/bluele/gcache) | 669 | 29 | 2015-01-25 | 5 days ago | Cache library with support for expirable Cache, LFU, LRU and ARC. |
| [gendry](https://github.com/didi/gendry) | 577 | 46 | 2017-12-01 | 2 months ago | Non-invasive SQL builder and powerful data binder. |
| [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) | 576 | 46 | 2018-04-11 | 2 hours ago | CovenantSQL is a SQL database on blockchain. |
| [eliasdb](https://github.com/krotik/eliasdb) | 505 | 23 | 2016-08-13 | 19 hours ago | Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language. |
| [goqu](https://github.com/doug-martin/goqu) | 492 | 25 | 2015-02-21 | 2 weeks ago | Idiomatic SQL builder and query library. |
| [dotsql](https://github.com/gchaincl/dotsql) | 407 | 21 | 2014-11-20 | 4 months ago | Go library that helps you keep sql files in one place and use them with ease. |
| [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) | 397 | 26 | 2015-12-11 | 6 months ago | Powerful data retrieval methods as well as DB-agnostic query building capabilities. |
| [levigo](https://github.com/jmhodges/levigo) | 352 | 24 | 2012-01-17 | 6 days ago | Levigo is a Go wrapper for LevelDB. |
| [fastcache](https://github.com/VictoriaMetrics/fastcache) | 350 | 13 | 2018-11-23 | 1 week ago | fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead. |
| [gormigrate](https://github.com/go-gormigrate/gormigrate) | 257 | 3 | 2016-08-31 | 1 month ago | Database schema migration helper for Gorm ORM. |
| [chproxy](https://github.com/Vertamedia/chproxy) | 242 | 20 | 2017-09-18 | 1 month ago | HTTP proxy for ClickHouse database. |
| [pudge](https://github.com/recoilme/pudge) | 182 | 6 | 2018-11-20 | 1 week ago | Fast and simple  key/value store written using Go's standard library. |
| [piladb](https://github.com/fern4lvarez/piladb) | 167 | 8 | 2015-09-09 | 11 months ago | Lightweight RESTful database engine based on stack data structures. |
| [scaneo](https://github.com/variadico/scaneo) | 141 | 8 | 2015-04-30 | 11 months ago | Generate Go code to convert database rows into arbitrary structs. |
| [sqrl](https://github.com/elgris/sqrl) | 141 | 5 | 2014-06-25 | 1 month ago | SQL query builder, fork of Squirrel with improved performance. |
| [myreplication](https://github.com/2tvenom/myreplication) | 130 | 17 | 2015-02-05 | 5 months ago | MySql binary log replication listener. Supports statement and row based replication. |
| [vasto](https://github.com/chrislusf/vasto) | 121 | 16 | 2018-01-16 | 8 hours ago | A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption. |
| [goose](https://github.com/steinbacher/goose) | 103 | 4 | 2016-03-05 | 2 years ago | Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts. |
| [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) | 100 | 6 | 2017-04-29 | 1 day ago | Collects small insterts and sends big requests to ClickHouse servers. |
| [darwin](https://github.com/GuiaBolso/darwin) | 77 | 2 | 2016-04-05 | 3 months ago | Database schema evolution library for Go. |
| [slowpoke](https://github.com/recoilme/slowpoke) | 74 | 6 | 2018-02-19 | 2 months ago | Key-value store with persistence. |
| [igor](https://github.com/galeone/igor) | 73 | 6 | 2016-03-10 | 8 months ago | Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax. |
| [godbal](https://github.com/xujiajun/godbal) | 48 | 3 | 2018-02-28 | 1 month ago | Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily. |
| [octillery](https://github.com/knocknote/octillery) | 43 | 14 | 2018-11-26 | 1 month ago | Go package for sharding databases ( Supports every ORM or raw SQL ). |
| [couchcache](https://github.com/codingsince1985/couchcache) | 38 | 3 | 2015-04-05 | 1 year ago | RESTful caching micro-service backed by Couchbase server. |
| [golang-scribble](https://github.com/nanobox-io/golang-scribble) | 32 | 3 | 2018-06-22 | 4 months ago | Tiny flat file JSON store. |
| [dbbench](https://github.com/sj14/dbbench) | 27 | 3 | 2018-11-24 | 1 month ago | Database benchmarking tool with support for several databases and scripts. |
| [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) | 25 | 5 | 2017-12-18 | 1 year ago | BigCache with clustering support and individual item expiration. |
| [prep](https://github.com/hexdigest/prep) | 23 | 3 | 2017-12-12 | 1 year ago | Use prepared SQL statements without changing your code. |
| [pravasan](https://github.com/pravasan/pravasan) | 23 | 6 | 2015-01-04 | 2 months ago | Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc. |
| [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) | 19 | 1 | 2018-08-11 | 2 months ago | A Go package to help write migrations with go-pg/pg. |
| [go-fixtures](https://github.com/RichardKnop/go-fixtures) | 17 | 1 | 2015-12-24 | 4 months ago | Django style fixtures for Golang's excellent built-in database/sql library. |
| [tempdb](https://github.com/rafaeljesus/tempdb) | 13 | 3 | 2017-03-18 | 1 year ago | Key-value store for temporary items. |
| [rwdb](https://github.com/andizzle/rwdb) | 10 | 2 | 2017-10-04 | 1 year ago | rwdb provides read replica capability for multiple database servers setup. |
| [gorocksdb](https://github.com/kapitan-k/gorocksdb) | 6 | 1 | 2017-12-28 | 1 year ago | Gorocksdb is a wrapper for [RocksDB](https://rocksdb.org) written in Go. |
| [ormlite](https://github.com/pupizoid/ormlite) | 0 | 1 | 2018-06-28 | 2 weeks ago | Lightweight package containing some ORM-like features and helpers for sqlite databases. |

## Database Drivers
        
*Libraries for connecting and operating databases.*

### Relational Databases
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [mysql](https://github.com/go-sql-driver/mysql) | 7059 | 392 | 2012-12-10 | 1 day ago | MySQL driver for Go. |
| [pq](https://github.com/lib/pq) | 4690 | 157 | 2012-03-13 | 15 hours ago | Pure Go Postgres driver for database/sql. |
| [go-sqlite3](https://github.com/mattn/go-sqlite3) | 3093 | 125 | 2011-11-11 | 2 weeks ago | SQLite3 driver for go that uses database/sql. |
| [pgx](https://github.com/jackc/pgx) | 1651 | 66 | 2013-03-31 | 5 days ago | PostgreSQL driver supporting features beyond those exposed by database/sql. |
| [go-mssqldb](https://github.com/denisenkom/go-mssqldb) | 911 | 59 | 2013-12-16 | 3 days ago | Microsoft MSSQL driver for Go. |
| [go-oci8](https://github.com/mattn/go-oci8) | 361 | 34 | 2012-02-29 | 4 weeks ago | Oracle driver for go that uses database/sql. |
| [goracle](https://github.com/go-goracle/goracle) | 182 | 28 | 2015-03-25 | 2 weeks ago | Oracle driver for Go, using the ODPI-C driver. |
| [firebirdsql](https://github.com/nakagami/firebirdsql) | 97 | 14 | 2013-08-27 | 6 hours ago | Firebird RDBMS SQL driver for Go. |
| [go-adodb](https://github.com/mattn/go-adodb) | 84 | 12 | 2011-11-14 | 1 week ago | Microsoft ActiveX Object DataBase driver for go that uses database/sql. |
| [gofreetds](https://github.com/minus5/gofreetds) | 84 | 21 | 2012-12-07 | 2 weeks ago | Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org). |
| [calcite-avatica-go](https://github.com/apache/calcite-avatica-go) | 26 | 16 | 2017-08-08 | 1 month ago | Apache Avatica/Phoenix SQL driver for database/sql. |
| [bgc](https://github.com/viant/bgc) | 10 | 9 | 2016-06-14 | 1 month ago | Datastore Connectivity for BigQuery for go. |

### NoSQL Databases
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [cayley](https://github.com/cayleygraph/cayley) | 12257 | 632 | 2014-06-06 | 1 week ago | Graph database with support for multiple backends. |
| [cayley](https://github.com/cayleygraph/cayley) | 12256 | 633 | 2014-06-06 | 1 week ago | Graph database with support for multiple backends. |
| [cayley](https://github.com/cayleygraph/cayley) | 12256 | 633 | 2014-06-06 | 1 week ago | Graph database with support for multiple backends. |
| [cayley](https://github.com/cayleygraph/cayley) | 12255 | 633 | 2014-06-06 | 1 week ago | Graph database with support for multiple backends. |
| [redigo](https://github.com/gomodule/redigo) | 5554 | 281 | 2012-04-14 | 5 days ago | Redigo is a Go client for the Redis database. |
| [redis](https://github.com/go-redis/redis) | 5162 | 200 | 2012-07-25 | 22 hours ago | Redis client for Golang. |
| [bleve](https://github.com/blevesearch/bleve) | 5007 | 232 | 2014-04-18 | 2 days ago | Modern text indexing library for go. |
| [riot](https://github.com/go-ego/riot) | 4305 | 170 | 2017-06-21 | 1 month ago | Go Open Source, Distributed, Simple and efficient Search Engine. |
| [elastic](https://github.com/olivere/elastic) | 3525 | 160 | 2012-12-07 | 2 days ago | Elasticsearch client for Go. |
| [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) | 2038 | 138 | 2017-02-09 | 9 hours ago | Official MongoDB driver for the Go language. |
| [mgo](https://github.com/globalsign/mgo) | 1430 | 77 | 2017-04-13 | 5 days ago | MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms. |
| [rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1418 | 45 | 2013-09-12 | 1 month ago | Go language driver for RethinkDB. |
| [rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1418 | 45 | 2013-09-12 | 1 month ago | Go language driver for RethinkDB. |
| [rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1418 | 45 | 2013-09-12 | 1 month ago | Go language driver for RethinkDB. |
| [rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1418 | 45 | 2013-09-12 | 1 month ago | Go language driver for RethinkDB. |
| [gomemcache](https://github.com/bradfitz/gomemcache) | 1010 | 51 | 2011-06-29 | 1 week ago | memcache client library for the Go programming language. |
| [elastigo](https://github.com/mattbaird/elastigo) | 936 | 49 | 2012-10-12 | 4 weeks ago | Elasticsearch client library. |
| [neoism](https://github.com/jmcvetta/neoism) | 350 | 23 | 2012-07-12 | 2 weeks ago | Neo4j client for Golang. |
| [elasticsql](https://github.com/cch123/elasticsql) | 296 | 21 | 2016-08-24 | 1 week ago | Convert sql to elasticsearch dsl in Go. |
| [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) | 289 | 39 | 2014-07-26 | 1 week ago | Aerospike client in Go language. |
| [go-couchbase](https://github.com/couchbase/go-couchbase) | 286 | 26 | 2012-01-20 | 1 day ago | Couchbase client in Go. |
| [gocb](https://github.com/couchbase/gocb) | 281 | 70 | 2015-01-16 | 2 weeks ago | Official Couchbase Go SDK. |
| [redeo](https://github.com/bsm/redeo) | 240 | 24 | 2014-03-06 | 3 months ago | Redis-protocol compatible TCP servers/services. |
| [cachego](https://github.com/fabiorphp/cachego) | 103 | 6 | 2016-10-06 | 9 months ago | Golang Cache component for multiple drivers. |
| [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) | 72 | 6 | 2011-06-05 | 8 months ago | Neo4j REST Client in golang. |
| [go-rejson](https://github.com/nitishm/go-rejson) | 67 | 4 | 2018-04-23 | 1 month ago | Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease. |
| [arangolite](https://github.com/solher/arangolite) | 64 | 5 | 2015-10-05 | 1 year ago | Lightweight golang driver for ArangoDB. |
| [dynago](https://github.com/underarmour/dynago) | 63 | 131 | 2015-05-18 | 1 year ago | Dynago is a principle of least surprise client for DynamoDB. |
| [skizze](https://github.com/seiflotfy/skizze) | 59 | 6 | 2016-01-17 | 2 years ago | probabilistic data-structures service and storage. |
| [go-couchdb](https://github.com/fjl/go-couchdb) | 51 | 6 | 2013-10-28 | 6 months ago | Yet another CouchDB HTTP API wrapper for Go. |
| [gokv](https://github.com/philippgille/gokv) | 48 | 4 | 2018-10-09 | 2 days ago | Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more) |
| [goforestdb](https://github.com/couchbase/goforestdb) | 29 | 44 | 2014-05-14 | 2 years ago | Go bindings for ForestDB. |
| [go-pilosa](https://github.com/pilosa/go-pilosa) | 27 | 13 | 2016-10-01 | 1 day ago | Go client library for Pilosa. |
| [goriak](https://github.com/zegl/goriak) | 25 | 2 | 2016-10-06 | 2 months ago | Go language driver for Riak KV. |
| [goes](https://github.com/OwnLocal/goes) | 25 | 34 | 2015-12-29 | 2 years ago | Library to interact with Elasticsearch. |
| [neo4j](https://github.com/cihangir/neo4j) | 24 | 3 | 2013-05-18 | 3 years ago | Neo4j Rest API Bindings for Golang. |
| [xredis](https://github.com/shomali11/xredis) | 9 | 1 | 2017-06-14 | 9 months ago | Typesafe, customizable, clean & easy to use Redis client. |
| [dsc](https://github.com/viant/dsc) | 8 | 17 | 2016-06-14 | 1 week ago | Datastore connectivity for SQL, NoSQL, structured files. |
| [asc](https://github.com/viant/asc) | 4 | 10 | 2016-06-14 | 3 months ago | Datastore Connectivity for Aerospike for go. |
| [godscache](https://github.com/defcronyke/godscache) | 4 | 2 | 2018-05-09 | 3 weeks ago | A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached. |

## Date and Time
        
*Libraries for working with dates and times.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [now](https://github.com/jinzhu/now) | 1982 | 51 | 2013-11-18 | 2 weeks ago | Now is a time toolkit for golang. |
| [dateparse](https://github.com/araddon/dateparse) | 807 | 14 | 2014-04-21 | 1 week ago | Parse date's without knowing format in advance. |
| [carbon](https://github.com/uniplaces/carbon) | 306 | 38 | 2016-08-03 | 1 month ago | Simple Time extension with a lot of util methods, ported from PHP Carbon library. |
| [durafmt](https://github.com/hako/durafmt) | 208 | 4 | 2016-05-21 | 6 months ago | Time duration formatting library for Go. |
| [timeutil](https://github.com/leekchan/timeutil) | 156 | 6 | 2015-08-02 | 1 month ago | Useful extensions (Timedelta, Strftime, ...) to the golang's time package. |
| [timespan](https://github.com/SaidinWoT/timespan) | 60 | 5 | 2014-10-07 | 1 month ago | For interacting with intervals of time, defined as a start time and a duration. |
| [iso8601](https://github.com/relvacode/iso8601) | 60 | 2 | 2017-04-25 | 2 months ago | Efficiently parse ISO8601 date-times without regex. |
| [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) | 48 | 3 | 2016-02-01 | 1 month ago | The implementation of the Persian (Solar Hijri) Calendar in Go (golang). |
| [date](https://github.com/rickb777/date) | 23 | 2 | 2015-11-24 | 3 months ago | Augments Time for working with dates, date ranges, time spans, periods, and time-of-day. |
| [goweek](https://github.com/grsmv/goweek) | 18 | 3 | 2015-11-12 | 2 years ago | Library for working with week entity in golang. |
| [feiertage](https://github.com/wlbr/feiertage) | 17 | 1 | 2015-11-04 | 2 weeks ago | Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving... |
| [go-sunrise](https://github.com/nathan-osman/go-sunrise) | 10 | 3 | 2017-06-16 | 1 year ago | Calculate the sunrise and sunset times for a given location. |
| [nulltime](https://github.com/kirillDanshin/nulltime) | 9 | 1 | 2016-03-06 | 1 year ago | Nullable `time.Time`. |
| [kair](https://github.com/GuilhermeCaruso/kair) | 9 | 1 | 2018-10-03 | 1 week ago | Date and Time - Golang Formatting Library. |
| [tuesday](https://github.com/osteele/tuesday) | 7 | 1 | 2017-08-11 | 1 year ago | Ruby-compatible Strftime function. |
| [strftime](https://github.com/awoodbeck/strftime) | 5 | 1 | 2018-02-10 | 1 year ago | C99-compatible strftime formatter. |

## Distributed Systems
        
*Packages that help with building Distributed Systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [etcd](https://github.com/etcd-io/etcd) | 23301 | 1229 | 2013-07-07 | 1 hour ago | Go implementation of the Raft consensus protocol, by CoreOS. |
| [etcd](https://github.com/etcd-io/etcd) | 23279 | 1228 | 2013-07-07 | 1 day ago | Go implementation of the Raft consensus protocol, by CoreOS. |
| [etcd](https://github.com/etcd-io/etcd) | 23258 | 1228 | 2013-07-07 | 1 day ago | Go implementation of the Raft consensus protocol, by CoreOS. |
| [etcd](https://github.com/etcd-io/etcd) | 23253 | 1228 | 2013-07-07 | 2 days ago | Go implementation of the Raft consensus protocol, by CoreOS. |
| [kit](https://github.com/go-kit/kit) | 12798 | 640 | 2015-02-03 | 14 hours ago | Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc. |
| [grpc-go](https://github.com/grpc/grpc-go) | 7603 | 436 | 2014-12-09 | 1 day ago | The Go language implementation of gRPC. HTTP/2 based RPC. |
| [jaeger](https://github.com/jaegertracing/jaeger) | 7354 | 277 | 2016-04-16 | 17 hours ago | A distributed tracing system. |
| [micro](https://github.com/micro/micro) | 5623 | 291 | 2015-01-17 | 21 hours ago | Pluggable microservice toolkit and distributed systems platform. |
| [gnatsd](https://github.com/nats-io/gnatsd) | 5316 | 329 | 2012-10-30 | 13 hours ago | Lightweight, high performance messaging system for microservices, IoT, and cloud native systems. |
| [rpcx](https://github.com/smallnest/rpcx) | 3244 | 270 | 2016-05-18 | 5 days ago | Distributed pluggable RPC service framework like alibaba Dubbo. |
| [tendermint](https://github.com/tendermint/tendermint) | 2719 | 242 | 2014-05-15 | 53 minutes ago | High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols. |
| [torrent](https://github.com/anacrolix/torrent) | 2613 | 121 | 2015-01-09 | 3 days ago | BitTorrent client package. |
| [raft](https://github.com/hashicorp/raft) | 2507 | 177 | 2013-11-05 | 1 day ago | Golang implementation of the Raft consensus protocol, by HashiCorp. |
| [glow](https://github.com/chrislusf/glow) | 2353 | 141 | 2015-06-14 | 4 months ago | Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go. |
| [gleam](https://github.com/chrislusf/gleam) | 1837 | 138 | 2016-08-26 | 2 months ago | Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed. |
| [emitter](https://github.com/emitter-io/emitter) | 1660 | 81 | 2016-10-29 | 2 days ago | High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love. |
| [krakend](https://github.com/devopsfaith/krakend) | 1287 | 67 | 2016-11-05 | 1 day ago | Ultra performant API Gateway framework with middlewares. |
| [hprose-golang](https://github.com/hprose/hprose-golang) | 922 | 89 | 2014-02-14 | 9 months ago | Very newbility RPC Library, support 25+ languages now. |
| [ringpop-go](https://github.com/uber/ringpop-go) | 531 | 2311 | 2015-06-06 | 2 weeks ago | Scalable, fault-tolerant application-layer sharding for Go applications. |
| [gorpc](https://github.com/valyala/gorpc) | 522 | 26 | 2014-11-21 | 1 year ago | Simple, fast and scalable RPC library for high load. |
| [go-health](https://github.com/InVisionApp/go-health) | 438 | 20 | 2017-11-30 | 3 months ago | Library for enabling asynchronous dependency health checks in your service. |
| [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) | 356 | 18 | 2015-10-09 | 1 week ago | Video streaming torrent client. |
| [sleuth](https://github.com/ursiform/sleuth) | 292 | 7 | 2016-04-23 | 11 months ago | Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)). |
| [digota](https://github.com/digota/digota) | 270 | 24 | 2017-08-14 | 4 months ago | grpc ecommerce microservice. |
| [go-jump](https://github.com/dgryski/go-jump) | 231 | 13 | 2014-06-16 | 1 year ago | Port of Google's "Jump" Consistent Hash function. |
| [consistent](https://github.com/buraksezer/consistent) | 152 | 9 | 2018-03-25 | 1 week ago | Consistent hashing with bounded loads. |
| [redis-lock](https://github.com/bsm/redis-lock) | 109 | 6 | 2014-10-10 | 3 months ago | Simplified distributed locking implementation using Redis. |
| [jsonrpc](https://github.com/osamingo/jsonrpc) | 103 | 6 | 2016-10-28 | 2 months ago | The jsonrpc package helps implement of JSON-RPC 2.0. |
| [jsonrpc](https://github.com/ybbus/jsonrpc) | 79 | 6 | 2016-11-10 | 3 months ago | JSON-RPC 2.0 HTTP client implementation. |
| [celeriac.v1](https://github.com/svcavallar/celeriac.v1) | 47 | 3 | 2015-10-10 | 2 months ago | Library for adding support for interacting and monitoring Celery workers, tasks and events in Go. |
| [doublejump](https://github.com/edwingeng/doublejump) | 27 | 3 | 2018-06-27 | 2 months ago | A revamped Google's jump consistent hash. |
| [drmaa](https://github.com/dgruber/drmaa) | 23 | 2 | 2013-03-17 | 9 months ago | Job submission library for cluster schedulers based on the DRMAA standard. |
| [flowgraph](https://github.com/vectaport/flowgraph) | 10 | 1 | 2018-08-30 | 11 hours ago | flow-based programming package. |
| [dynatomic](https://github.com/tylfin/dynatomic) | 6 | 0 | 2019-02-09 | 1 week ago | A library for using DynamoDB as an atomic counter. |
| [outboxer](https://github.com/italolelis/outboxer) | 0 | 0 | 2019-02-01 | 1 week ago | Outboxer is a go library that implements the outbox pattern. |

## Email
        
*Libraries and tools that implement email creation and sending.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [MailHog](https://github.com/mailhog/MailHog) | 4521 | 125 | 2014-04-17 | 4 days ago | Email and SMTP testing with web and API interface. |
| [gomail](https://github.com/go-gomail/gomail) | 2140 | 65 | 2014-10-15 | 3 weeks ago | Gomail is a very simple and powerful package to send emails. |
| [hermes](https://github.com/matcornic/hermes) | 1480 | 24 | 2017-03-26 | 1 week ago | Golang package that generates clean, responsive HTML e-mails. |
| [email](https://github.com/jordan-wright/email) | 1008 | 35 | 2013-12-13 | 2 weeks ago | A robust and flexible email library for Go. |
| [go-imap](https://github.com/emersion/go-imap) | 625 | 31 | 2016-04-27 | 2 hours ago | IMAP library for clients and servers. |
| [sendgrid-go](https://github.com/sendgrid/sendgrid-go) | 469 | 181 | 2013-09-12 | 3 days ago | SendGrid's Go library for sending email. |
| [hectane](https://github.com/hectane/hectane) | 160 | 12 | 2015-08-28 | 10 months ago | Lightweight SMTP client providing an HTTP API. |
| [douceur](https://github.com/aymerick/douceur) | 142 | 4 | 2015-04-09 | 11 months ago | CSS inliner for your HTML emails. |
| [go-message](https://github.com/emersion/go-message) | 79 | 6 | 2016-12-31 | 1 month ago | Streaming library for the Internet Message Format and mail messages. |
| [smtp](https://github.com/mailhog/smtp) | 49 | 9 | 2014-12-25 | 8 months ago | SMTP server protocol state machine. |
| [go-dkim](https://github.com/toorop/go-dkim) | 47 | 3 | 2015-04-29 | 1 year ago | DKIM library, to sign & verify email. |

## Embeddable Scripting Languages
        
*Embedding other languages inside your go code.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [otto](https://github.com/robertkrimen/otto) | 4411 | 188 | 2012-10-06 | 1 week ago | JavaScript interpreter written in Go. |
| [gopher-lua](https://github.com/yuin/gopher-lua) | 2613 | 137 | 2015-02-15 | 3 weeks ago | Lua 5.1 VM and compiler written in Go. |
| [go-lua](https://github.com/Shopify/go-lua) | 1542 | 251 | 2013-12-21 | 3 months ago | Port of the Lua 5.2 VM to pure Go. |
| [tengo](https://github.com/d5/tengo) | 1071 | 25 | 2019-01-09 | 9 hours ago | Bytecode compiled script language for Go. |
| [anko](https://github.com/mattn/anko) | 843 | 47 | 2014-03-28 | 1 week ago | Scriptable interpreter written in Go. |
| [go-python](https://github.com/sbinet/go-python) | 809 | 43 | 2012-07-09 | 1 month ago | naive go bindings to the CPython C-API. |
| [go-duktape](https://github.com/olebedev/go-duktape) | 628 | 23 | 2015-01-08 | 2 weeks ago | Duktape JavaScript engine bindings for Go. |
| [go-php](https://github.com/deuill/go-php) | 611 | 39 | 2015-09-18 | 4 months ago | PHP bindings for Go. |
| [golua](https://github.com/aarzilli/golua) | 423 | 35 | 2010-12-07 | 6 months ago | Go bindings for Lua C API. |
| [gisp](https://github.com/jcla1/gisp) | 414 | 20 | 2014-01-11 | 1 year ago | Simple LISP in Go. |
| [agora](https://github.com/mna/agora) | 317 | 22 | 2013-06-19 | 4 years ago | Dynamically typed, embeddable programming language in Go. |
| [agora](https://github.com/mna/agora) | 317 | 22 | 2013-06-19 | 4 years ago | Dynamically typed, embeddable programming language in Go. |
| [agora](https://github.com/mna/agora) | 317 | 22 | 2013-06-19 | 4 years ago | Dynamically typed, embeddable programming language in Go. |
| [agora](https://github.com/mna/agora) | 317 | 22 | 2013-06-19 | 4 years ago | Dynamically typed, embeddable programming language in Go. |
| [expr](https://github.com/antonmedv/expr) | 297 | 16 | 2018-07-14 | 1 month ago | an engine that can evaluate expressions. |
| [gval](https://github.com/PaesslerAG/gval) | 86 | 12 | 2017-09-27 | 2 weeks ago | A highly customizable expression language written in Go. |
| [purl](https://github.com/ian-kent/purl) | 25 | 2 | 2014-11-30 | 4 years ago | Perl 5.18.2 embedded in Go. |
| [binder](https://github.com/alexeyco/binder) | 24 | 2 | 2017-04-03 | 7 months ago | Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua). |
| [ngaro](https://github.com/db47h/ngaro) | 16 | 1 | 2016-08-09 | 9 months ago | Embeddable Ngaro VM implementation enabling scripting in Retro. |

## Error Handling
        
*Libraries for handling errors.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [errors](https://github.com/pkg/errors) | 4171 | 92 | 2015-12-27 | 1 week ago | Package that provides simple error handling primitives. |
| [go-multierror](https://github.com/hashicorp/go-multierror) | 559 | 83 | 2014-12-16 | 2 months ago | Go (golang) package for representing a list of errors as a single error. |
| [errorx](https://github.com/joomcode/errorx) | 507 | 38 | 2018-08-17 | 2 weeks ago | A feature rich error package with stack traces, composition of errors and more. |
| [tracerr](https://github.com/ztrue/tracerr) | 174 | 6 | 2019-02-07 | 2 weeks ago | Golang errors with stack trace and source fragments. |
| [werr](https://github.com/txgruppi/werr) | 10 | 0 | 2015-08-04 | 3 years ago | Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called. |

## Files
        
*Libraries for  handling files and file systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [afero](https://github.com/spf13/afero) | 1948 | 84 | 2014-10-28 | 1 day ago | FileSystem Abstraction System for Go. |
| [pdfcpu](https://github.com/hhrutter/pdfcpu) | 798 | 23 | 2017-06-19 | 13 hours ago | PDF processor. |
| [notify](https://github.com/rjeczalik/notify) | 452 | 26 | 2014-09-09 | 2 weeks ago | File system event notification library with simple API, similar to os/signal. |
| [opc](https://github.com/qmuntal/opc) | 50 | 1 | 2018-11-06 | 1 week ago | Load Open Packaging Conventions (OPC) files for Go. |
| [skywalker](https://github.com/dixonwille/skywalker) | 41 | 3 | 2017-08-02 | 1 year ago | Package to allow one to concurrently go through a filesystem with ease. |
| [go-csv-tag](https://github.com/artonge/go-csv-tag) | 39 | 1 | 2017-06-18 | 4 months ago | Load csv file using tag. |
| [tarfs](https://github.com/posener/tarfs) | 31 | 2 | 2017-03-11 | 1 year ago | Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files. |
| [go-gtfs](https://github.com/artonge/go-gtfs) | 13 | 2 | 2017-07-09 | 1 month ago | Load gtfs files in go. |
| [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) | 11 | 1 | 2018-10-16 | 4 months ago | Copy files for humans. |

## Financial
        
*Packages for accounting and finance.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [decimal](https://github.com/shopspring/decimal) | 1341 | 56 | 2015-02-26 | 2 days ago | Arbitrary-precision fixed-point decimal numbers. |
| [go-money](https://github.com/Rhymond/go-money) | 552 | 15 | 2017-03-21 | 1 month ago | Implementation of Fowler's Money pattern. |
| [go-money](https://github.com/Rhymond/go-money) | 552 | 15 | 2017-03-21 | 1 month ago | Implementation of Fowler's Money pattern. |
| [go-money](https://github.com/Rhymond/go-money) | 552 | 15 | 2017-03-21 | 1 month ago | Implementation of Fowler's Money pattern. |
| [go-money](https://github.com/Rhymond/go-money) | 552 | 15 | 2017-03-21 | 1 month ago | Implementation of Fowler's Money pattern. |
| [go-finance](https://github.com/FlashBoys/go-finance) | 537 | 28 | 2016-02-28 | 1 year ago | Comprehensive financial markets data in Go. |
| [accounting](https://github.com/leekchan/accounting) | 443 | 12 | 2015-08-10 | 1 month ago | money and currency formatting for golang. |
| [techan](https://github.com/sdcoffey/techan) | 108 | 17 | 2017-03-08 | 2 months ago | Technical analysis library with advanced market analysis and trading strategies. |
| [vat](https://github.com/dannyvankooten/vat) | 54 | 1 | 2016-06-19 | 5 months ago | VAT number validation & EU VAT rates. |
| [ofxgo](https://github.com/aclindsa/ofxgo) | 53 | 5 | 2015-11-08 | 41 minutes ago | Query OFX servers and/or parse the responses (with example command-line client). |
| [transaction](https://github.com/claygod/transaction) | 46 | 8 | 2017-10-11 | 6 months ago | Embedded transactional database of accounts, running in multithreaded mode. |
| [go-finance](https://github.com/alpeb/go-finance) | 33 | 2 | 2017-06-01 | 10 months ago | Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations. |
| [orderbook](https://github.com/i25959341/orderbook) | 22 | 3 | 2018-04-25 | 5 days ago | Matching Engine for Limit Order Book in Golang. |

## Forms
        
*Libraries for working with forms.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [nosurf](https://github.com/justinas/nosurf) | 920 | 35 | 2013-08-23 | 1 month ago | CSRF protection middleware for Go. |
| [binding](https://github.com/mholt/binding) | 731 | 30 | 2014-05-21 | 11 months ago | Binds form and JSON data from net/http Request to struct. |
| [csrf](https://github.com/gorilla/csrf) | 371 | 19 | 2015-08-03 | 2 months ago | CSRF protection for Go web applications & services. |
| [form](https://github.com/go-playground/form) | 324 | 8 | 2016-05-26 | 1 month ago | Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support. |
| [conform](https://github.com/leebenson/conform) | 163 | 5 | 2016-01-06 | 8 months ago | Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags. |
| [formam](https://github.com/monoculum/formam) | 111 | 2 | 2014-10-25 | 6 months ago | decode form's values into a struct. |
| [forms](https://github.com/albrow/forms) | 96 | 6 | 2014-08-08 | 1 year ago | Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files. |
| [bind](https://github.com/robfig/bind) | 23 | 2 | 2014-08-06 | 4 years ago | Bind form data to any Go values. |

## Functional
        
*Packages to support functional programming in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [go-underscore](https://github.com/tobyhede/go-underscore) | 1014 | 26 | 2014-07-02 | 2 weeks ago | Useful collection of helpfully functional Go collection utilities. |
| [fpGo](https://github.com/TeaEntityLab/fpGo) | 72 | 1 | 2018-05-24 | 7 months ago | Monad, Functional Programming features for Golang. |
| [fuego](https://github.com/seborama/fuego) | 17 | 1 | 2018-11-06 | 1 day ago | Functional Experiment in Go. |

## Game Development
        
*Awesome game development libraries.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [leaf](https://github.com/name5566/leaf) | 2729 | 303 | 2014-08-04 | 2 months ago | Lightweight game server framework. |
| [pixel](https://github.com/faiface/pixel) | 2118 | 92 | 2016-11-19 | 1 week ago | Hand-crafted 2D game library in Go. |
| [ebiten](https://github.com/hajimehoshi/ebiten) | 1509 | 65 | 2013-06-16 | 1 day ago | dead simple 2D game library in Go. |
| [go-sdl2](https://github.com/veandco/go-sdl2) | 1051 | 42 | 2013-06-06 | 2 days ago | Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/). |
| [gonet](https://github.com/xtaci/gonet) | 998 | 125 | 2013-04-11 | 1 year ago | Game server skeleton implemented with golang. |
| [termloop](https://github.com/JoelOtter/termloop) | 977 | 35 | 2015-05-24 | 3 days ago | Terminal-based game engine for Go, built on top of Termbox. |
| [engo](https://github.com/EngoEngine/engo) | 955 | 45 | 2014-11-12 | 18 hours ago | Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm. |
| [goworld](https://github.com/xiaonanln/goworld) | 947 | 100 | 2017-06-03 | 2 months ago | Scalable game server engine, featuring space-entity framework and hot-swapping. |
| [nano](https://github.com/lonng/nano) | 849 | 53 | 2017-08-02 | 1 month ago | Lightweight, facility, high performance golang based game server framework. |
| [engine](https://github.com/g3n/engine) | 598 | 58 | 2017-03-08 | 2 weeks ago | Go 3D Game Engine. |
| [oak](https://github.com/oakmound/oak) | 574 | 34 | 2017-07-16 | 2 days ago | Pure Go game engine. |
| [engine](https://github.com/azul3d/engine) | 406 | 23 | 2016-02-29 | 8 months ago | 3D game engine written in Go. |
| [raylib-go](https://github.com/gen2brain/raylib-go) | 341 | 18 | 2017-01-27 | 3 months ago | Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming. |
| [go-astar](https://github.com/beefsack/go-astar) | 307 | 10 | 2014-05-28 | 11 months ago | Go implementation of the A\* path finding algorithm. |
| [GarageEngine](https://github.com/vova616/GarageEngine) | 303 | 31 | 2012-08-07 | 5 years ago | 2d game engine written in Go working on OpenGL. |
| [pitaya](https://github.com/topfreegames/pitaya) | 205 | 26 | 2018-03-20 | 5 days ago | Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. |
| [go3d](https://github.com/ungerik/go3d) | 150 | 10 | 2011-06-27 | 5 days ago | Performance oriented 2D/3D math package for Go. |
| [glop](https://github.com/runningwild/glop) | 76 | 4 | 2011-04-21 | 3 years ago | Glop (Game Library Of Power) is a fairly simple cross-platform game library. |
| [go-collada](https://github.com/GlenKelley/go-collada) | 12 | 3 | 2013-09-19 | 5 years ago | Go package for working with the Collada file format. |

## Generation and Generics
        
*Tools to enhance the language with features like generics via code generation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [go-linq](https://github.com/ahmetb/go-linq) | 1661 | 72 | 2013-12-19 | 3 months ago | .NET LINQ-like query methods for Go. |
| [go-linq](https://github.com/ahmetb/go-linq) | 1661 | 72 | 2013-12-19 | 3 months ago | .NET LINQ-like query methods for Go. |
| [go-linq](https://github.com/ahmetb/go-linq) | 1661 | 72 | 2013-12-19 | 3 months ago | .NET LINQ-like query methods for Go. |
| [go-linq](https://github.com/ahmetb/go-linq) | 1661 | 72 | 2013-12-19 | 3 months ago | .NET LINQ-like query methods for Go. |
| [jennifer](https://github.com/dave/jennifer) | 1140 | 21 | 2016-12-05 | 1 month ago | Generate arbitrary Go code without templates. |
| [gen](https://github.com/clipperhouse/gen) | 978 | 36 | 2013-10-14 | 8 months ago | Code generation tool for ‘generics’-like functionality. |
| [goderive](https://github.com/awalterschulze/goderive) | 672 | 22 | 2017-02-11 | 1 week ago | Derives functions from input types. |
| [gowrap](https://github.com/hexdigest/gowrap) | 206 | 7 | 2018-09-15 | 1 week ago | Generate decorators for Go interfaces using simple templates. |
| [interfaces](https://github.com/rjeczalik/interfaces) | 160 | 5 | 2015-12-06 | 2 months ago | Command line tool for generating interface definitions. |
| [pkgreflect](https://github.com/ungerik/pkgreflect) | 78 | 5 | 2012-09-03 | 1 year ago | Go preprocessor for package scoped reflection. |
| [go-enum](https://github.com/abice/go-enum) | 66 | 4 | 2017-08-11 | 2 weeks ago | Code generation for enums from code comments. |
| [efaceconv](https://github.com/t0pep0/efaceconv) | 40 | 4 | 2016-11-18 | 1 year ago | Code generation tool for high performance conversion from interface{} to immutable type without allocations. |
| [gotype](https://github.com/wzshiming/gotype) | 18 | 3 | 2017-12-05 | 2 weeks ago | Golang source code parsing, usage like reflect package. |

## Geographic
        
*Geographic tools and servers*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [tile38](https://github.com/tidwall/tile38) | 5856 | 186 | 2016-03-05 | 13 hours ago | Geolocation DB with spatial index and realtime geofencing. |
| [geo](https://github.com/golang/geo) | 802 | 73 | 2014-12-04 | 1 week ago | S2 geometry library in Go. |
| [geocache](https://github.com/melihmucuk/geocache) | 94 | 5 | 2016-06-21 | 2 years ago | In-memory cache that is suitable for geolocation based applications. |
| [osm](https://github.com/paulmach/osm) | 41 | 5 | 2016-02-02 | 2 months ago | Library for reading, writing and working with OpenStreetMap data and APIs. |
| [geoserver](https://github.com/hishamkaram/geoserver) | 19 | 2 | 2018-03-27 | 3 months ago | geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API. |
| [pbf](https://github.com/maguro/pbf) | 10 | 1 | 2017-09-19 | 4 months ago | OpenStreetMap PBF golang encoder/decoder. |
| [gismanager](https://github.com/hishamkaram/gismanager) | 9 | 1 | 2018-09-29 | 4 months ago | Publish Your GIS Data(Vector Data) to PostGIS and Geoserver. |

## Go Compilers
        
*Tools for compiling Go to other languages.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gopherjs](https://github.com/gopherjs/gopherjs) | 8009 | 249 | 2013-08-28 | 21 hours ago | Compiler from Go to JavaScript. |
| [llgo](https://github.com/go-llvm/llgo) | 952 | 61 | 2011-11-05 | 4 years ago | LLVM-based compiler for Go. |
| [tardisgo](https://github.com/tardisgo/tardisgo) | 382 | 27 | 2014-01-08 | 2 years ago | Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler. |
| [c4go](https://github.com/Konstantin8105/c4go) | 73 | 9 | 2018-03-28 | 1 day ago | Transpile C code to Go code. |
| [f4go](https://github.com/Konstantin8105/f4go) | 9 | 2 | 2018-07-09 | 2 months ago | Transpile FORTRAN 77 code to Go code. |

## Goroutines
        
*Tools for managing and working with Goroutines.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [goworker](https://github.com/benmanns/goworker) | 2136 | 70 | 2013-07-23 | 4 weeks ago | goworker is a Go-based background worker. |
| [ants](https://github.com/panjf2000/ants) | 1367 | 54 | 2018-05-19 | 1 week ago | A high-performance goroutine pool for golang. |
| [tunny](https://github.com/Jeffail/tunny) | 1138 | 51 | 2014-04-03 | 3 months ago | Goroutine pool for golang. |
| [grpool](https://github.com/ivpusic/grpool) | 460 | 29 | 2015-07-22 | 1 month ago | Lightweight Goroutine pool. |
| [pool](https://github.com/go-playground/pool) | 452 | 12 | 2015-10-29 | 2 years ago | Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation. |
| [go-floc](https://github.com/workanator/go-floc) | 164 | 8 | 2017-07-03 | 1 year ago | Orchestrate goroutines with ease. |
| [go-flow](https://github.com/kamildrazkiewicz/go-flow) | 94 | 7 | 2016-09-25 | 1 year ago | Control goroutines execution order. |
| [workerpool](https://github.com/gammazero/workerpool) | 80 | 5 | 2016-05-17 | 2 months ago | Goroutine pool that limits the concurrency of task execution, not the number of tasks queued. |
| [semaphore](https://github.com/marusama/semaphore) | 64 | 4 | 2017-11-22 | 1 month ago | Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations). |
| [semaphore](https://github.com/kamilsk/semaphore) | 61 | 1 | 2016-10-08 | 7 hours ago | Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context. |
| [GoSlaves](https://github.com/dgrr/GoSlaves) | 60 | 4 | 2017-09-18 | 7 months ago | Simple and Asynchronous Goroutine pool library. |
| [GoSlaves](https://github.com/dgrr/GoSlaves) | 60 | 4 | 2017-09-18 | 7 months ago | Simple and Asynchronous Goroutine pool library. |
| [GoSlaves](https://github.com/dgrr/GoSlaves) | 60 | 4 | 2017-09-18 | 7 months ago | Simple and Asynchronous Goroutine pool library. |
| [GoSlaves](https://github.com/dgrr/GoSlaves) | 60 | 4 | 2017-09-18 | 7 months ago | Simple and Asynchronous Goroutine pool library. |
| [gpool](https://github.com/sherifabdlnaby/gpool) | 47 | 1 | 2018-12-03 | 2 days ago | manages a resizeable pool of context-aware goroutines to bound concurrency. |
| [gpool](https://github.com/sherifabdlnaby/gpool) | 47 | 1 | 2018-12-03 | 2 days ago | manages a resizeable pool of context-aware goroutines to bound concurrency. |
| [gpool](https://github.com/sherifabdlnaby/gpool) | 47 | 1 | 2018-12-03 | 2 days ago | manages a resizeable pool of context-aware goroutines to bound concurrency. |
| [gpool](https://github.com/sherifabdlnaby/gpool) | 46 | 1 | 2018-12-03 | 2 days ago | manages a resizeable pool of context-aware goroutines to bound concurrency. |
| [worker-pool](https://github.com/vardius/worker-pool) | 32 | 3 | 2017-10-04 | 4 weeks ago | goworker is a Go simple async worker pool. |
| [parallel-fn](https://github.com/rafaeljesus/parallel-fn) | 24 | 3 | 2017-06-18 | 1 year ago | Run functions in parallel. |
| [cyclicbarrier](https://github.com/marusama/cyclicbarrier) | 23 | 3 | 2018-01-11 | 4 months ago | CyclicBarrier for golang. |
| [async](https://github.com/StudioSol/async) | 14 | 9 | 2017-07-01 | 5 months ago | A safe way to execute functions asynchronously, recovering them in case of panic. |
| [async](https://github.com/StudioSol/async) | 14 | 9 | 2017-07-01 | 5 months ago | A safe way to execute functions asynchronously, recovering them in case of panic. |
| [async](https://github.com/StudioSol/async) | 14 | 9 | 2017-07-01 | 5 months ago | A safe way to execute functions asynchronously, recovering them in case of panic. |
| [async](https://github.com/StudioSol/async) | 14 | 9 | 2017-07-01 | 5 months ago | A safe way to execute functions asynchronously, recovering them in case of panic. |
| [threadpool](https://github.com/shettyh/threadpool) | 12 | 2 | 2017-09-07 | 1 year ago | Golang threadpool implementation. |
| [breaker](https://github.com/kamilsk/breaker) | 7 | 1 | 2019-02-15 | 7 hours ago | 🚧 Flexible mechanism to make your code breakable. |
| [artifex](https://github.com/borderstech/artifex) | 6 | 2 | 2018-11-01 | 4 months ago | Simple in-memory job queue for Golang using worker-based dispatching. |
| [stl](https://github.com/ssgreg/stl) | 5 | 1 | 2018-06-19 | 5 months ago | Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism. |
| [go-trylock](https://github.com/subchen/go-trylock) | 4 | 1 | 2018-04-26 | 9 months ago | TryLock support on read-write lock for Golang. |

## GUI
        
*Interaction*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [ui](https://github.com/andlabs/ui) | 6463 | 356 | 2014-02-18 | 3 weeks ago | Platform-native GUI library for Go. Cross platform. |
| [qt](https://github.com/therecipe/qt) | 5269 | 273 | 2014-11-19 | 1 month ago | Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi). |
| [robotgo](https://github.com/go-vgo/robotgo) | 3920 | 176 | 2016-09-27 | 21 hours ago | Go Native cross-platform GUI system automation. Control the mouse, keyboard and other. |
| [webview](https://github.com/zserge/webview) | 3904 | 173 | 2017-08-19 | 3 days ago | Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux). |
| [walk](https://github.com/lxn/walk) | 3279 | 240 | 2010-09-16 | 6 days ago | Windows application library kit for Go. |
| [app](https://github.com/maxence-charriere/app) | 2753 | 106 | 2016-10-12 | 5 days ago | Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. |
| [app](https://github.com/maxence-charriere/app) | 2752 | 106 | 2016-10-12 | 5 days ago | Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. |
| [app](https://github.com/maxence-charriere/app) | 2751 | 106 | 2016-10-12 | 5 days ago | Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. |
| [app](https://github.com/maxence-charriere/app) | 2750 | 106 | 2016-10-12 | 5 days ago | Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. |
| [go-astilectron](https://github.com/asticode/go-astilectron) | 2371 | 113 | 2017-04-22 | 1 month ago | Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron). |
| [go-sciter](https://github.com/sciter-sdk/go-sciter) | 1299 | 114 | 2015-10-15 | 4 months ago | Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform. |
| [fyne](https://github.com/fyne-io/fyne) | 781 | 42 | 2018-02-05 | 20 hours ago | Cross platform native GUIs designed for Go, rendered using EFL. Supports: Linux, macOS, Windows. |
| [systray](https://github.com/getlantern/systray) | 637 | 38 | 2014-11-12 | 2 weeks ago | Cross platform Go library to place an icon and menu in the notification area. |
| [gotk3](https://github.com/gotk3/gotk3) | 623 | 43 | 2015-08-13 | 1 day ago | Go bindings for GTK3. |
| [gosx-notifier](https://github.com/deckarep/gosx-notifier) | 476 | 15 | 2013-11-25 | 1 year ago | OSX Desktop Notifications library for Go. |
| [gowd](https://github.com/dtylman/gowd) | 173 | 19 | 2017-03-29 | 4 months ago | Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform. |
| [trayhost](https://github.com/shurcooL/trayhost) | 143 | 4 | 2014-04-25 | 4 months ago | Cross-platform Go library to place an icon in the host operating system's taskbar. |

## Hardware
        
*Libraries, tools, and tutorials for interacting with hardware.*

## Images
        
*Libraries for manipulating images.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [ln](https://github.com/fogleman/ln) | 2378 | 89 | 2016-01-10 | 1 year ago | 3D line art rendering in Go. |
| [imaginary](https://github.com/h2non/imaginary) | 2314 | 70 | 2015-03-05 | 1 week ago | Fast and simple HTTP microservice for image resizing. |
| [imaging](https://github.com/disintegration/imaging) | 2190 | 64 | 2012-12-07 | 15 hours ago | Simple Go image processing package. |
| [resize](https://github.com/nfnt/resize) | 2025 | 73 | 2012-08-03 | 1 year ago | Image resizing for Go with common interpolation methods. |
| [gocv](https://github.com/hybridgroup/gocv) | 2002 | 103 | 2017-09-19 | 2 weeks ago | Go package for computer vision using OpenCV 3.3+. |
| [bild](https://github.com/anthonynsimon/bild) | 1931 | 55 | 2016-08-01 | 5 months ago | Collection of image processing algorithms in pure Go. |
| [pt](https://github.com/fogleman/pt) | 1717 | 62 | 2015-01-24 | 1 year ago | Path tracing engine written in Go. |
| [gg](https://github.com/fogleman/gg) | 1710 | 68 | 2016-02-19 | 1 week ago | 2D rendering in pure Go. |
| [svgo](https://github.com/ajstarks/svgo) | 1223 | 47 | 2010-03-06 | 5 months ago | Go Language Library for SVG generation. |
| [smartcrop](https://github.com/muesli/smartcrop) | 1204 | 30 | 2014-04-08 | 4 months ago | Finds good crops for arbitrary images and crop sizes. |
| [gift](https://github.com/disintegration/gift) | 1157 | 49 | 2014-07-13 | 5 months ago | Package of image processing filters. |
| [go-opencv](https://github.com/go-opencv/go-opencv) | 1028 | 68 | 2013-12-09 | 1 month ago | Go bindings for OpenCV. |
| [go-opencv](https://github.com/go-opencv/go-opencv) | 1028 | 68 | 2013-12-09 | 1 month ago | Go bindings for OpenCV. |
| [go-opencv](https://github.com/go-opencv/go-opencv) | 1027 | 68 | 2013-12-09 | 1 month ago | Go bindings for OpenCV. |
| [go-opencv](https://github.com/go-opencv/go-opencv) | 1027 | 68 | 2013-12-09 | 1 month ago | Go bindings for OpenCV. |
| [geopattern](https://github.com/pravj/geopattern) | 984 | 20 | 2014-10-23 | 1 month ago | Create beautiful generative image patterns from a string. |
| [picfit](https://github.com/thoas/picfit) | 969 | 41 | 2014-12-07 | 1 month ago | An image resizing server written in Go. |
| [imagick](https://github.com/gographics/imagick) | 886 | 54 | 2013-05-01 | 1 month ago | Go binding to ImageMagick's MagickWand C API. |
| [bimg](https://github.com/h2non/bimg) | 694 | 35 | 2015-03-17 | 3 days ago | Small package for fast and efficient image processing using libvips. |
| [mort](https://github.com/aldor007/mort) | 343 | 17 | 2017-11-19 | 2 weeks ago | Storage and image processing server written in Go. |
| [stegify](https://github.com/DimitarPetrov/stegify) | 282 | 13 | 2018-11-30 | 3 weeks ago | Go tool for LSB steganography, capable of hiding any file within an image. |
| [govatar](https://github.com/o1egl/govatar) | 281 | 5 | 2016-01-18 | 10 months ago | Library and CMD tool for generating funny avatars. |
| [go-nude](https://github.com/koyachi/go-nude) | 272 | 17 | 2014-05-02 | 3 months ago | Nudity detection with Go. |
| [image2ascii](https://github.com/qeesung/image2ascii) | 238 | 7 | 2018-10-20 | 3 months ago | Convert image to ASCII. |
| [goimagehash](https://github.com/corona10/goimagehash) | 168 | 9 | 2017-07-29 | 3 weeks ago | Go Perceptual image hashing package. |
| [rez](https://github.com/bamiaux/rez) | 156 | 9 | 2014-01-17 | 1 year ago | Image resizing in pure Go and SIMD. |
| [img](https://github.com/hawx/img) | 127 | 5 | 2012-07-29 | 3 years ago | Selection of image manipulation tools. |
| [go-cairo](https://github.com/ungerik/go-cairo) | 83 | 6 | 2012-08-23 | 5 months ago | Go binding for the cairo graphics library. |
| [go-gd](https://github.com/bolknote/go-gd) | 47 | 4 | 2011-05-12 | 10 months ago | Go binding for GD library. |
| [mergi](https://github.com/noelyahan/mergi) | 42 | 2 | 2018-09-24 | 1 week ago | Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate). |
| [go-webcolors](https://github.com/jyotiska/go-webcolors) | 23 | 1 | 2014-04-24 | 3 years ago | Port of webcolors library from Python to Go. |
| [tga](https://github.com/ftrvxmtrx/tga) | 21 | 2 | 2012-10-08 | 3 years ago | Package tga is a TARGA image format decoder/encoder. |
| [cameron](https://github.com/aofei/cameron) | 17 | 1 | 2018-05-06 | 1 week ago | An avatar generator for Go. |
| [steganography](https://github.com/auyer/steganography) | 12 | 3 | 2018-05-22 | 4 months ago | Pure Go Library for LSB steganography. |
| [mpo](https://github.com/donatj/mpo) | 5 | 1 | 2015-04-15 | 1 month ago | Decoder and conversion tool for MPO 3D Photos. |

## IoT
        
*Libraries for programming devices of the IoT.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gobot](https://github.com/hybridgroup/gobot) | 5210 | 298 | 2013-09-21 | 2 weeks ago | Gobot is a framework for robotics, physical computing, and the Internet of Things. |
| [flogo](https://github.com/TIBCOSoftware/flogo) | 961 | 113 | 2016-07-10 | 5 days ago | Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. |
| [flogo](https://github.com/TIBCOSoftware/flogo) | 959 | 113 | 2016-07-10 | 5 days ago | Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. |
| [flogo](https://github.com/TIBCOSoftware/flogo) | 955 | 113 | 2016-07-10 | 5 days ago | Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. |
| [flogo](https://github.com/TIBCOSoftware/flogo) | 954 | 113 | 2016-07-10 | 5 days ago | Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. |
| [gatt](https://github.com/paypal/gatt) | 770 | 57 | 2014-04-23 | 3 months ago | Gatt is a Go package for building Bluetooth Low Energy peripherals. |
| [mainflux](https://github.com/mainflux/mainflux) | 368 | 54 | 2015-07-07 | 1 hour ago | Industrial IoT Messaging and Device Management Server. |
| [mainflux](https://github.com/mainflux/mainflux) | 366 | 54 | 2015-07-07 | 1 day ago | Industrial IoT Messaging and Device Management Server. |
| [mainflux](https://github.com/mainflux/mainflux) | 365 | 53 | 2015-07-07 | 1 day ago | Industrial IoT Messaging and Device Management Server. |
| [mainflux](https://github.com/mainflux/mainflux) | 363 | 53 | 2015-07-07 | 2 days ago | Industrial IoT Messaging and Device Management Server. |
| [devices](https://github.com/goiot/devices) | 221 | 17 | 2016-05-30 | 2 years ago | Suite of libraries for IoT devices, experimental for x/exp/io. |
| [sensorbee](https://github.com/sensorbee/sensorbee) | 165 | 19 | 2016-02-19 | 1 year ago | Lightweight stream processing engine for IoT. |
| [connectordb](https://github.com/connectordb/connectordb) | 152 | 17 | 2015-01-17 | 5 days ago | Open-Source Platform for Quantified Self & IoT. |
| [huego](https://github.com/amimof/huego) | 85 | 1 | 2017-05-16 | 1 week ago | An extensive Philips Hue client library for Go. |
| [eywa](https://github.com/xcodersun/eywa) | 30 | 8 | 2016-02-21 | 1 year ago | Project Eywa is essentially a connection manager that keeps track of connected devices. |
| [iot](https://github.com/vaelen/iot) | 28 | 4 | 2018-03-08 | 10 months ago | IoT is a simple framework for implementing a Google IoT Core device. |

## Job Scheduler
        
*Libraries for scheduling jobs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gron](https://github.com/roylee0704/gron) | 552 | 16 | 2016-06-04 | 1 week ago | Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly. |
| [jobrunner](https://github.com/bamzi/jobrunner) | 521 | 17 | 2015-10-21 | 2 years ago | Smart and featureful cron job scheduler with job queuing and live monitoring built in. |
| [jobs](https://github.com/albrow/jobs) | 439 | 18 | 2015-02-10 | 8 months ago | Persistent and flexible background jobs library. |
| [scheduler](https://github.com/carlescere/scheduler) | 266 | 16 | 2015-02-04 | 8 months ago | Cronjobs scheduling made easy. |
| [go-cron](https://github.com/rk/go-cron) | 169 | 9 | 2011-04-15 | 3 years ago | Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. |
| [clockwork](https://github.com/whiteShtef/clockwork) | 52 | 1 | 2018-04-24 | 3 weeks ago | Simple and intuitive job scheduling library in Go. |
| [leprechaun](https://github.com/kilgaloon/leprechaun) | 24 | 4 | 2018-04-08 | 6 days ago | Job scheduler that supports webhooks, crons and classic scheduling. |

## JSON
        
*Libraries for working with JSON.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gjson](https://github.com/tidwall/gjson) | 4103 | 113 | 2016-08-11 | 2 weeks ago | Get a JSON value with one line of code. |
| [gojson](https://github.com/ChimeraCoder/gojson) | 1920 | 45 | 2012-12-28 | 3 weeks ago | Automatically generate Go (golang) struct definitions from example JSON. |
| [gojq](https://github.com/elgs/gojq) | 132 | 5 | 2015-12-30 | 9 months ago | JSON query in Golang. |
| [kazaam](https://github.com/qntfy/kazaam) | 106 | 15 | 2016-07-19 | 6 months ago | API for arbitrary transformation of JSON documents. |
| [kazaam](https://github.com/qntfy/kazaam) | 106 | 15 | 2016-07-19 | 6 months ago | API for arbitrary transformation of JSON documents. |
| [kazaam](https://github.com/qntfy/kazaam) | 106 | 15 | 2016-07-19 | 6 months ago | API for arbitrary transformation of JSON documents. |
| [kazaam](https://github.com/qntfy/kazaam) | 106 | 15 | 2016-07-19 | 6 months ago | API for arbitrary transformation of JSON documents. |
| [jsongo](https://github.com/ricardolonga/jsongo) | 85 | 1 | 2015-08-08 | 2 years ago | Fluent API to make it easier to create Json objects. |
| [jsonf](https://github.com/miolini/jsonf) | 53 | 3 | 2015-05-25 | 2 years ago | Console tool for highlighted formatting and struct query fetching JSON. |
| [jaydiff](https://github.com/yazgazan/jaydiff) | 33 | 1 | 2017-04-25 | 2 months ago | JSON diff utility written in Go. |
| [mp](https://github.com/sanbornm/mp) | 29 | 2 | 2014-06-16 | 2 years ago | Simple cli email parser. It currently takes stdin and outputs JSON. |
| [go-respond](https://github.com/nicklaw5/go-respond) | 18 | 1 | 2017-03-13 | 4 months ago | Go package for handling common HTTP JSON responses. |
| [jsonhal](https://github.com/RichardKnop/jsonhal) | 8 | 2 | 2016-01-15 | 4 months ago | Simple Go package to make custom structs marshal into HAL compatible JSON responses. |
| [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) | 5 | 1 | 2016-07-08 | 2 years ago | Go bindings based on the JSON API errors reference. |

## Logging
        
*Libraries for generating and working with log files.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [logrus](https://github.com/sirupsen/logrus) | 10019 | 267 | 2013-10-17 | 1 day ago | Structured logger for Go. |
| [logrus](https://github.com/sirupsen/logrus) | 10007 | 267 | 2013-10-17 | 1 day ago | Structured logger for Go. |
| [logrus](https://github.com/sirupsen/logrus) | 9997 | 267 | 2013-10-17 | 1 day ago | Structured logger for Go. |
| [logrus](https://github.com/sirupsen/logrus) | 9986 | 267 | 2013-10-17 | 3 days ago | Structured logger for Go. |
| [zap](https://github.com/uber-go/zap) | 6160 | 198 | 2016-02-19 | 17 hours ago | Fast, structured, leveled logging in Go. |
| [go-spew](https://github.com/davecgh/go-spew) | 2981 | 59 | 2013-01-09 | 1 month ago | Implements a deep pretty printer for Go data structures to aid in debugging. |
| [glog](https://github.com/golang/glog) | 2145 | 84 | 2013-07-16 | 2 hours ago | Leveled execution logs for Go. |
| [zerolog](https://github.com/rs/zerolog) | 1773 | 47 | 2017-05-12 | 1 day ago | Zero-allocation JSON logger. |
| [tail](https://github.com/hpcloud/tail) | 1361 | 101 | 2013-02-05 | 6 days ago | Go package striving to emulate the features of the BSD tail program. |
| [seelog](https://github.com/cihub/seelog) | 1270 | 89 | 2011-11-17 | 2 days ago | Logging functionality with flexible dispatching, filtering, and formatting. |
| [lumberjack](https://github.com/natefinch/lumberjack) | 1188 | 37 | 2014-06-14 | 4 days ago | Simple rolling logger, implements io.WriteCloser. |
| [log15](https://github.com/inconshreveable/log15) | 843 | 23 | 2014-05-20 | 4 months ago | Simple, powerful logging for Go. |
| [log](https://github.com/apex/log) | 676 | 35 | 2015-12-22 | 4 months ago | Structured logging package for Go. |
| [logxi](https://github.com/mgutz/logxi) | 328 | 9 | 2015-03-02 | 1 year ago | 12-factor app logger that is fast and makes you happy. |
| [onelog](https://github.com/francoispqt/onelog) | 305 | 9 | 2018-05-06 | 7 hours ago | Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenario. Also, it is one of the logger with the lowest allocation. |
| [log](https://github.com/go-playground/log) | 257 | 10 | 2016-02-08 | 2 weeks ago | Simple, configurable and scalable Structured Logging for Go. |
| [logutils](https://github.com/hashicorp/logutils) | 239 | 88 | 2013-10-09 | 6 months ago | Utilities for slightly better logging in Go (Golang) extending the standard logger. |
| [go-logger](https://github.com/apsdehal/go-logger) | 213 | 8 | 2014-09-26 | 5 months ago | Simple logger of Go Programs, with level handlers. |
| [logger](https://github.com/azer/logger) | 130 | 4 | 2014-09-30 | 5 months ago | Minimalistic logging library for Go. |
| [xlog](https://github.com/rs/xlog) | 126 | 7 | 2015-10-22 | 11 months ago | Structured logger for `net/context` aware HTTP handlers with flexible dispatching. |
| [ozzo-log](https://github.com/go-ozzo/ozzo-log) | 102 | 11 | 2015-10-23 | 11 months ago | High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail). |
| [logvoyage](https://github.com/firstrow/logvoyage) | 83 | 5 | 2015-03-29 | 1 year ago | Full-featured logging saas written in golang. |
| [log](https://github.com/alexcesaro/log) | 44 | 4 | 2014-04-19 | 3 years ago | Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs. |
| [gologger](https://github.com/sadlil/gologger) | 38 | 6 | 2015-09-02 | 1 year ago | Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch. |
| [go-log](https://github.com/ian-kent/go-log) | 33 | 2 | 2014-05-02 | 11 months ago | Log4j implementation in Go. |
| [glg](https://github.com/kpango/glg) | 32 | 3 | 2017-06-21 | 2 days ago | glg is simple and fast leveled logging library for Go. |
| [logex](https://github.com/chzyer/logex) | 32 | 7 | 2014-10-10 | 1 year ago | Golang log lib, supports tracking and level, wrap by standard log lib. |
| [logrusly](https://github.com/sebest/logrusly) | 24 | 5 | 2014-09-12 | 11 months ago | [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/). |
| [gone](https://github.com/One-com/gone) | 24 | 7 | 2016-09-05 | 3 weeks ago | Fast, extendable, full-featured, std-lib source compatible log library. |
| [go-log](https://github.com/siddontang/go-log) | 22 | 3 | 2014-05-18 | 1 week ago | Log lib supports level and multi handlers. |
| [log](https://github.com/teris-io/log) | 21 | 1 | 2017-10-29 | 1 year ago | Structured log interface for Go cleanly separates logging facade from its implementation. |
| [go-cronowriter](https://github.com/utahta/go-cronowriter) | 17 | 1 | 2017-02-04 | 1 day ago | Simple writer that rotate log files automatically based on current date and time, like cronolog. |
| [journald](https://github.com/ssgreg/journald) | 17 | 2 | 2017-08-23 | 2 months ago | Go implementation of systemd Journal's native API for logging. |
| [mlog](https://github.com/jbrodriguez/mlog) | 16 | 2 | 2014-10-20 | 7 months ago | Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output. |
| [distillog](https://github.com/amoghe/distillog) | 15 | 1 | 2015-10-13 | 7 months ago | distilled levelled logging (think of it as stdlib + log levels). |
| [gomol](https://github.com/aphistic/gomol) | 12 | 2 | 2015-08-30 | 9 months ago | Multiple-output, structured logging for Go with extensible logging outputs. |
| [go-log](https://github.com/subchen/go-log) | 8 | 1 | 2017-05-07 | 9 months ago | Simple and configurable Logging in Go, with level, formatters and writers. |
| [logdump](https://github.com/ewwwwwqm/logdump) | 7 | 2 | 2017-01-13 | 11 months ago | Package for multi-level logging. |
| [xlog](https://github.com/xfxdev/xlog) | 7 | 1 | 2016-05-06 | 1 month ago | Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format. |
| [logmatic](https://github.com/borderstech/logmatic) | 5 | 1 | 2018-11-07 | 1 month ago | Colorized logger for Golang with dynamic log level configuration. |
| [glo](https://github.com/lajosbencz/glo) | 5 | 1 | 2019-01-20 | 1 month ago | PHP Monolog inspired logging facility with identical severity levels. |
| [logo](https://github.com/mbndr/logo) | 4 | 1 | 2017-02-08 | 1 year ago | Golang logger to different configurable writers. |

## Machine Learning
        
*Libraries for Machine Learning.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [golearn](https://github.com/sjwhitworth/golearn) | 6316 | 429 | 2013-12-26 | 3 weeks ago | General Machine Learning library for Go. |
| [gorgonia](https://github.com/gorgonia/gorgonia) | 2447 | 162 | 2016-09-15 | 2 days ago | graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. |
| [gorgonia](https://github.com/gorgonia/gorgonia) | 2447 | 161 | 2016-09-15 | 2 days ago | graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. |
| [gorgonia](https://github.com/gorgonia/gorgonia) | 2446 | 162 | 2016-09-15 | 2 days ago | graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. |
| [gorgonia](https://github.com/gorgonia/gorgonia) | 2445 | 162 | 2016-09-15 | 2 days ago | graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. |
| [tfgo](https://github.com/galeone/tfgo) | 1073 | 48 | 2017-05-23 | 2 weeks ago | Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python. |
| [goml](https://github.com/cdipaolo/goml) | 967 | 73 | 2015-06-27 | 2 years ago | On-line Machine Learning in Go. |
| [gosseract](https://github.com/otiai10/gosseract) | 746 | 31 | 2013-10-11 | 1 week ago | Go package for OCR (Optical Character Recognition), by using Tesseract C++ library. |
| [CloudForest](https://github.com/ryanbressler/CloudForest) | 627 | 43 | 2012-10-23 | 3 months ago | Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go. |
| [bayesian](https://github.com/jbrukh/bayesian) | 597 | 30 | 2011-11-23 | 2 weeks ago | Naive Bayesian Classification for Golang. |
| [eaopt](https://github.com/MaxHalford/eaopt) | 579 | 29 | 2016-01-31 | 2 weeks ago | An evolutionary optimization library. |
| [gorse](https://github.com/zhenghaoz/gorse) | 346 | 16 | 2018-08-14 | 7 hours ago | A High Performance Recommender System Package based on Collaborative Filtering for Go. |
| [gobrain](https://github.com/goml/gobrain) | 337 | 23 | 2014-04-29 | 2 months ago | Neural Networks written in go. |
| [regommend](https://github.com/muesli/regommend) | 232 | 14 | 2014-02-06 | 3 months ago | Recommendation & collaborative filtering engine. |
| [go-deep](https://github.com/patrikeh/go-deep) | 201 | 11 | 2017-12-09 | 5 months ago | A feature-rich neural network library in Go. |
| [go-galib](https://github.com/thoj/go-galib) | 164 | 14 | 2009-11-30 | 3 years ago | Genetic Algorithms library written in Go / golang. |
| [goRecommend](https://github.com/timkaye11/goRecommend) | 136 | 10 | 2014-07-16 | 4 years ago | Recommendation Algorithms library written in Go. |
| [shield](https://github.com/eaigner/shield) | 119 | 10 | 2013-04-11 | 3 years ago | Bayesian text classifier with flexible tokenizers and storage backends for Go. |
| [go-fann](https://github.com/white-pony/go-fann) | 98 | 8 | 2011-03-11 | 4 years ago | Go bindings for Fast Artificial Neural Networks(FANN) library. |
| [goga](https://github.com/tomcraven/goga) | 72 | 8 | 2015-10-20 | 2 years ago | Genetic algorithm library for Go. |
| [neural-go](https://github.com/schuyler/neural-go) | 60 | 2 | 2011-10-17 | 4 years ago | Multilayer perceptron network implemented in Go, with training via backpropagation. |
| [libsvm](https://github.com/datastream/libsvm) | 59 | 10 | 2012-07-31 | 2 years ago | libsvm golang version derived work based on LIBSVM 3.14. |
| [go-pr](https://github.com/daviddengcn/go-pr) | 56 | 6 | 2013-06-07 | 5 years ago | Pattern recognition package in Go lang. |
| [neat](https://github.com/jinyeom/neat) | 51 | 11 | 2016-11-17 | 8 months ago | Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT). |
| [golinear](https://github.com/danieldk/golinear) | 38 | 5 | 2013-04-05 | 6 months ago | liblinear bindings for Go. |
| [goscore](https://github.com/asafschers/goscore) | 27 | 2 | 2017-08-19 | 6 months ago | Go Scoring API for PMML. |
| [fonet](https://github.com/Fontinalis/fonet) | 26 | 3 | 2017-10-03 | 10 months ago | A Deep Neural Network library written in Go. |
| [godist](https://github.com/e-dard/godist) | 22 | 2 | 2014-09-05 | 3 years ago | Various probability distributions, and associated methods. |
| [Varis](https://github.com/Xamber/Varis) | 21 | 5 | 2017-10-10 | 7 months ago | Golang Neural Network. |
| [go-cluster](https://github.com/e-XpertSolutions/go-cluster) | 18 | 5 | 2017-10-04 | 7 months ago | Go implementation of the k-modes and k-prototypes clustering algorithms. |
| [probab](https://github.com/ThePaw/probab) | 9 | 2 | 2015-09-14 | 3 years ago | Probability distribution functions. Bayesian inference. Written in pure Go. |
| [evoli](https://github.com/khezen/evoli) | 5 | 3 | 2015-06-12 | 1 month ago | Genetic Algorithm and Particle Swarm Optimization library. |
| [gomind](https://github.com/surenderthakran/gomind) | 5 | 2 | 2017-10-19 | 7 months ago | A simplistic Neural Network Library in Go. |
| [mlgo](https://github.com/NullHypothesis/mlgo) | 4 | 2 | 2015-12-07 | 3 years ago | This project aims to provide minimalistic machine learning algorithms in Go. |

## Messaging
        
*Libraries that implement messaging systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [sarama](https://github.com/Shopify/sarama) | 3952 | 361 | 2013-07-06 | 1 day ago | Go library for Apache Kafka. |
| [centrifugo](https://github.com/centrifugal/centrifugo) | 3356 | 174 | 2015-04-01 | 6 days ago | Real-time messaging (Websockets or SockJS) server in Go. |
| [gorush](https://github.com/appleboy/gorush) | 3280 | 151 | 2016-03-22 | 6 days ago | Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm). |
| [machinery](https://github.com/RichardKnop/machinery) | 2950 | 131 | 2015-04-06 | 1 day ago | Asynchronous task queue/job queue based on distributed message passing. |
| [go-socket.io](https://github.com/googollee/go-socket.io) | 2629 | 122 | 2013-07-13 | 3 hours ago | socket.io library for golang, a realtime application framework. |
| [go-nats](https://github.com/nats-io/go-nats) | 2134 | 159 | 2012-08-15 | 16 hours ago | Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library. |
| [go-nats](https://github.com/nats-io/go-nats) | 2130 | 158 | 2012-08-15 | 2 days ago | Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library. |
| [go-nats](https://github.com/nats-io/go-nats) | 2130 | 159 | 2012-08-15 | 2 days ago | Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library. |
| [go-nats](https://github.com/nats-io/go-nats) | 2130 | 159 | 2012-08-15 | 1 day ago | Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library. |
| [apns2](https://github.com/sideshow/apns2) | 1936 | 72 | 2016-01-05 | 1 day ago | HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps. |
| [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) | 1771 | 239 | 2013-12-27 | 1 year ago | gopush-cluster is a go push server cluster. |
| [benthos](https://github.com/Jeffail/benthos) | 1729 | 54 | 2016-03-22 | 15 hours ago | A message streaming bridge between a range of protocols. |
| [mangos](https://github.com/nanomsg/mangos) | 1521 | 89 | 2014-10-25 | 1 day ago | Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability. |
| [mangos](https://github.com/nanomsg/mangos) | 1521 | 89 | 2014-10-25 | 2 days ago | Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability. |
| [mangos](https://github.com/nanomsg/mangos) | 1521 | 89 | 2014-10-25 | 5 hours ago | Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability. |
| [mangos](https://github.com/nanomsg/mangos) | 1521 | 89 | 2014-10-25 | 1 day ago | Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability. |
| [melody](https://github.com/olahol/melody) | 1329 | 43 | 2015-05-14 | 4 months ago | Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling. |
| [go-nsq](https://github.com/nsqio/go-nsq) | 1321 | 61 | 2013-08-29 | 2 weeks ago | the official Go package for NSQ. |
| [uniqush-push](https://github.com/uniqush/uniqush-push) | 1056 | 71 | 2011-08-29 | 4 months ago | Redis backed unified push service for server-side notifications to mobile devices. |
| [mercure](https://github.com/dunglas/mercure) | 899 | 42 | 2018-07-14 | 2 hours ago | Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events). |
| [zmq4](https://github.com/pebbe/zmq4) | 729 | 47 | 2013-10-18 | 1 month ago | Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2). |
| [gollum](https://github.com/trivago/gollum) | 655 | 35 | 2015-06-21 | 4 weeks ago | A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations. |
| [EventBus](https://github.com/asaskevich/EventBus) | 482 | 27 | 2014-12-20 | 1 month ago | The lightweight event bus with async compatibility. |
| [golongpoll](https://github.com/jcuga/golongpoll) | 397 | 21 | 2015-11-02 | 2 weeks ago | HTTP longpoll server library that makes web pub-sub simple. |
| [dbus](https://github.com/godbus/dbus) | 316 | 15 | 2014-03-28 | 19 hours ago | Native Go bindings for D-Bus. |
| [glue](https://github.com/desertbit/glue) | 299 | 13 | 2015-06-07 | 1 year ago | Robust Go and Javascript Socket Library (Alternative to Socket.io). |
| [emitter](https://github.com/olebedev/emitter) | 283 | 10 | 2015-11-11 | 1 month ago | Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins. |
| [pubsub](https://github.com/cskr/pubsub) | 246 | 8 | 2012-04-01 | 3 weeks ago | Simple pubsub package for go. |
| [pubsub](https://github.com/cskr/pubsub) | 245 | 8 | 2012-04-01 | 3 weeks ago | Simple pubsub package for go. |
| [pubsub](https://github.com/cskr/pubsub) | 244 | 8 | 2012-04-01 | 3 weeks ago | Simple pubsub package for go. |
| [pubsub](https://github.com/cskr/pubsub) | 243 | 8 | 2012-04-01 | 3 weeks ago | Simple pubsub package for go. |
| [guble](https://github.com/smancke/guble) | 132 | 13 | 2015-11-16 | 1 year ago | Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence. |
| [oplog](https://github.com/dailymotion/oplog) | 94 | 97 | 2014-11-06 | 3 years ago | Generic oplog/replication system for REST APIs. |
| [drone-line](https://github.com/appleboy/drone-line) | 57 | 5 | 2016-09-13 | 3 months ago | Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI. |
| [rabbus](https://github.com/rafaeljesus/rabbus) | 55 | 8 | 2017-05-07 | 1 month ago | A tiny wrapper over amqp exchanges and queues. |
| [rabtap](https://github.com/jandelgado/rabtap) | 54 | 6 | 2017-11-11 | 4 days ago | RabbitMQ swiss army knife cli app. |
| [RapidMQ](https://github.com/sybrexsys/RapidMQ) | 52 | 3 | 2016-10-05 | 1 year ago | RapidMQ is a lightweight and reliable library for managing of the local messages queue. |
| [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) | 46 | 7 | 2017-01-16 | 1 year ago | A tiny wrapper around NSQ topic and channel. |
| [go-notify](https://github.com/TheCreeper/go-notify) | 44 | 2 | 2015-03-02 | 3 weeks ago | Native implementation of the freedesktop notification spec. |
| [goose](https://github.com/ian-kent/goose) | 35 | 1 | 2014-12-21 | 4 years ago | Server Sent Events in Go. |
| [message-bus](https://github.com/vardius/message-bus) | 33 | 3 | 2017-10-04 | 4 weeks ago | messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD. |
| [hub](https://github.com/leandro-lugaresi/hub) | 20 | 1 | 2018-04-14 | 10 months ago | A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges. |
| [event](https://github.com/agoalofalife/event) | 20 | 2 | 2017-07-02 | 1 year ago | Implementation of the pattern observer. |
| [go-vitotrol](https://github.com/maxatome/go-vitotrol) | 11 | 6 | 2016-11-04 | 4 months ago | Client library to Viessmann Vitotrol web service. |
| [gaurun-client](https://github.com/osamingo/gaurun-client) | 7 | 1 | 2017-06-29 | 7 months ago | Gaurun Client written in Go. |
| [jazz](https://github.com/socifi/jazz) | 4 | 4 | 2018-10-22 | 4 weeks ago | A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages. |

## Microsoft Office
        

### Microsoft Excel
        
*Libraries for working with Microsoft Excel.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [excelize](https://github.com/360EntSecGroup-Skylar/excelize) | 3363 | 127 | 2016-08-29 | 1 hour ago | Golang library for reading and writing Microsoft Excel™ (XLSX) files. |
| [xlsx](https://github.com/tealeg/xlsx) | 3020 | 175 | 2011-06-28 | 1 week ago | Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs. |
| [xlsx](https://github.com/plandem/xlsx) | 35 | 3 | 2017-08-27 | 1 month ago | Fast and safe way to read/update your existing Microsoft Excel files in Go programs. |
| [go-excel](https://github.com/szyhf/go-excel) | 32 | 2 | 2017-09-03 | 2 months ago | A simple and light reader to read a relate-db-like excel as a table. |
| [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) | 11 | 1 | 2017-03-13 | 7 months ago | Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files. |

## Miscellaneous
        

### Dependency Injection
        
*Libraries for working with dependency injection.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [dig](https://github.com/uber-go/dig) | 615 | 35 | 2017-03-22 | 2 weeks ago | A reflection based dependency injection toolkit for Go. |
| [alice](https://github.com/magic003/alice) | 31 | 1 | 2017-04-09 | 1 year ago | Additive dependency injection container for Golang. |
| [wire](https://github.com/Fs02/wire) | 13 | 1 | 2018-07-05 | 1 month ago | Strict Runtime Dependency Injection for Golang. |

### Strings
        
*These libraries were placed here because none of the other categories seemed to fit.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gopsutil](https://github.com/shirou/gopsutil) | 3434 | 172 | 2014-04-18 | 1 day ago | Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc). |
| [archiver](https://github.com/mholt/archiver) | 2096 | 43 | 2016-04-09 | 4 days ago | Library and command for making and extracting .zip and .tar.gz archives. |
| [gosms](https://github.com/haxpax/gosms) | 1201 | 54 | 2015-01-24 | 1 year ago | Your own local SMS gateway in Go that can be used to send SMS. |
| [go-resiliency](https://github.com/eapache/go-resiliency) | 765 | 29 | 2014-11-29 | 3 weeks ago | Resiliency patterns for golang. |
| [go-commons-pool](https://github.com/jolestar/go-commons-pool) | 592 | 43 | 2015-12-28 | 6 months ago | Generic object pool for Golang. |
| [xstrings](https://github.com/huandu/xstrings) | 552 | 24 | 2015-01-06 | 6 months ago | Collection of useful string functions ported from other languages. |
| [base64Captcha](https://github.com/mojocn/base64Captcha) | 486 | 33 | 2017-12-12 | 3 months ago | Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha. |
| [shortid](https://github.com/teris-io/shortid) | 402 | 6 | 2016-01-04 | 1 year ago | Distributed generation of super short, unique, non-sequential, URL friendly IDs. |
| [llvm](https://github.com/llir/llvm) | 351 | 29 | 2014-09-19 | 3 days ago | Library for interacting with LLVM IR in pure Go. |
| [health](https://github.com/dimiro1/health) | 343 | 6 | 2016-03-09 | 7 months ago | Easy to use, extensible health check library. |
| [go-conv](https://github.com/cstockton/go-conv) | 336 | 10 | 2016-10-11 | 1 year ago | Package conv provides fast and intuitive conversions across Go types. |
| [gofakeit](https://github.com/brianvoe/gofakeit) | 242 | 4 | 2015-04-24 | 1 month ago | Random data generator written in go. |
| [banner](https://github.com/dimiro1/banner) | 205 | 4 | 2016-03-26 | 2 years ago | Add beautiful banners into your Go applications. |
| [gountries](https://github.com/pariz/gountries) | 197 | 6 | 2016-01-13 | 3 months ago | Package that exposes country and subdivision data. |
| [battery](https://github.com/distatus/battery) | 122 | 2 | 2016-03-13 | 5 days ago | Cross-platform, normalized battery information library. |
| [antch](https://github.com/antchfx/antch) | 121 | 10 | 2017-09-28 | 7 months ago | A fast, powerful and extensible web crawling & scraping framework. |
| [stats](https://github.com/go-playground/stats) | 115 | 3 | 2015-09-15 | 2 years ago | Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... |
| [ffmt](https://github.com/go-ffmt/ffmt) | 108 | 4 | 2015-02-14 | 3 months ago | Beautify data display for Humans. |
| [lk](https://github.com/hyperboloide/lk) | 89 | 5 | 2016-07-15 | 1 month ago | A simple licensing library for golang. |
| [bitio](https://github.com/icza/bitio) | 80 | 5 | 2016-05-31 | 1 year ago | Highly optimized bit-level Reader and Writer for Go. |
| [turtle](https://github.com/hackebrot/turtle) | 69 | 1 | 2017-09-09 | 1 year ago | Emojis for Go. |
| [healthcheck](https://github.com/etherlabsio/healthcheck) | 62 | 4 | 2017-08-18 | 1 year ago | An opinionated and concurrent health-check HTTP handler for RESTful services. |
| [gommit](https://github.com/antham/gommit) | 61 | 3 | 2016-08-30 | 5 hours ago | Analyze git commit messages to ensure they follow defined patterns. |
| [go-unarr](https://github.com/gen2brain/go-unarr) | 56 | 5 | 2015-11-01 | 1 month ago | Decompression library for RAR, TAR, ZIP and 7z archives. |
| [indigo](https://github.com/osamingo/indigo) | 47 | 1 | 2016-08-31 | 4 months ago | Distributed unique ID generator of using Sonyflake and encoded by Base58. |
| [morse](https://github.com/alwindoss/morse) | 45 | 2 | 2018-08-15 | 1 week ago | Library to convert to and from morse code. |
| [strutil](https://github.com/ozgio/strutil) | 45 | 1 | 2018-08-16 | 6 months ago | String utilities. |
| [xkg](https://github.com/go-xkg/xkg) | 37 | 1 | 2015-01-05 | 4 years ago | X Keyboard Grabber. |
| [captcha](https://github.com/steambap/captcha) | 33 | 4 | 2017-09-12 | 6 days ago | Package captcha provides an easy to use, unopinionated API for captcha generation. |
| [browscap_go](https://github.com/digitalcrab/browscap_go) | 27 | 6 | 2014-09-18 | 1 month ago | GoLang Library for [Browser Capabilities Project](http://browscap.org/). |
| [persian](https://github.com/mavihq/persian) | 27 | 2 | 2017-10-17 | 7 months ago | Some utilities for Persian language in go. |
| [datacounter](https://github.com/miolini/datacounter) | 26 | 1 | 2015-10-15 | 11 months ago | Go counters for readers/writer/http.ResponseWriter. |
| [autoflags](https://github.com/artyom/autoflags) | 23 | 2 | 2014-05-16 | 1 month ago | Go package to automatically define command line flags from struct fields. |
| [pdfgen](https://github.com/hyperboloide/pdfgen) | 21 | 3 | 2015-12-01 | 1 year ago | HTTP service to generate PDF from Json requests. |
| [ghorg](https://github.com/gabrie30/ghorg) | 17 | 2 | 2018-03-29 | 2 months ago | Clone all repos from a GitHub org into a single directory. |
| [xdg](https://github.com/rkoesters/xdg) | 14 | 1 | 2013-12-15 | 3 months ago | FreeDesktop.org (xdg) Specs implemented in Go. |
| [gosh](https://github.com/osamingo/gosh) | 13 | 0 | 2018-05-25 | 1 month ago | Provide Go Statistics Handler, Struct, Measure Method. |
| [url-shortener](https://github.com/pantrif/url-shortener) | 12 | 1 | 2018-06-04 | 8 months ago | A modern, powerful, and robust URL shortener microservice with mysql support. |
| [anagent](https://github.com/mudler/anagent) | 9 | 2 | 2017-12-30 | 6 months ago | Minimalistic, pluggable Golang evloop/timer handler with dependency-injection. |
| [avgRating](https://github.com/kirillDanshin/avgRating) | 8 | 1 | 2017-08-06 | 1 year ago | Calculate average score and rating based on Wilson Score Equation. |
| [sandid](https://github.com/aofei/sandid) | 8 | 1 | 2018-06-12 | 1 week ago | Every grain of sand on earth has its own ID. |
| [hostutils](https://github.com/Wing924/hostutils) | 8 | 1 | 2017-09-26 | 1 month ago | A golang library for packing and unpacking FQDNs list. |
| [shellwords](https://github.com/Wing924/shellwords) | 5 | 1 | 2017-09-28 | 1 year ago | A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell. |
| [generators](https://github.com/azr/generators) | 3 | 1 | 2016-02-29 | 2 years ago | Generate boilerplate http input and output handling. |

## Natural Language Processing
        
*Libraries for working with human languages.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [prose](https://github.com/jdkato/prose) | 1926 | 46 | 2017-02-18 | 1 month ago | Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. |
| [when](https://github.com/olebedev/when) | 908 | 21 | 2016-12-27 | 2 days ago | Natural EN and RU language date/time parser with pluggable rules. |
| [go-i18n](https://github.com/nicksnyder/go-i18n) | 852 | 17 | 2012-01-15 | 1 month ago | Package and an accompanying tool to work with localized text. |
| [gse](https://github.com/go-ego/gse) | 793 | 40 | 2017-06-23 | 2 days ago | Go efficient text segmentation; support english, chinese, japanese and other. |
| [gojieba](https://github.com/yanyiwu/gojieba) | 721 | 59 | 2015-09-12 | 5 months ago | This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm. |
| [go-pinyin](https://github.com/mozillazg/go-pinyin) | 446 | 28 | 2014-11-09 | 6 months ago | CN Hanzi to Hanyu Pinyin converter. |
| [kagome](https://github.com/ikawaha/kagome) | 356 | 21 | 2014-06-26 | 2 weeks ago | JP morphological analyzer written in pure Go. |
| [nlp](https://github.com/shixzie/nlp) | 348 | 23 | 2017-01-25 | 1 year ago | Extract values from strings and fill your structs with nlp. |
| [nlp](https://github.com/shixzie/nlp) | 347 | 23 | 2017-01-25 | 1 year ago | Extract values from strings and fill your structs with nlp. |
| [whatlanggo](https://github.com/abadojack/whatlanggo) | 322 | 14 | 2017-02-21 | 3 hours ago | Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc). |
| [sentences](https://github.com/neurosnap/sentences) | 248 | 12 | 2015-08-07 | 3 weeks ago | Sentence tokenizer:  converts text into a list of sentences. |
| [nlp](https://github.com/james-bowman/nlp) | 190 | 20 | 2017-03-15 | 4 days ago | Go Natural Language Processing library supporting LSA (Latent Semantic Analysis). |
| [go-nlp](https://github.com/nuance/go-nlp) | 78 | 8 | 2011-05-02 | 7 years ago | Utilities for working with discrete probability distributions and other tools useful for doing NLP work. |
| [gounidecode](https://github.com/fiam/gounidecode) | 62 | 3 | 2012-05-01 | 3 years ago | Unicode transliterator (also known as unidecode) for Go. |
| [textcat](https://github.com/pebbe/textcat) | 60 | 6 | 2012-09-21 | 7 months ago | Go package for n-gram based text categorization, with support for utf-8 and raw text. |
| [MMSEGO](https://github.com/awsong/MMSEGO) | 58 | 5 | 2012-04-18 | 6 years ago | This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm. |
| [getlang](https://github.com/rylans/getlang) | 52 | 3 | 2018-03-02 | 1 month ago | Fast natural language detection package. |
| [go-stem](https://github.com/agonopol/go-stem) | 51 | 2 | 2011-09-24 | 8 months ago | Implementation of the porter stemming algorithm. |
| [go-unidecode](https://github.com/mozillazg/go-unidecode) | 45 | 2 | 2016-07-08 | 2 years ago | ASCII transliterations of Unicode text. |
| [stemmer](https://github.com/dchest/stemmer) | 43 | 4 | 2011-03-21 | 2 years ago | Stemmer packages for Go programming language. Includes English and German stemmers. |
| [RAKE.Go](https://github.com/afjoseph/RAKE.Go) | 39 | 5 | 2016-12-17 | 2 months ago | Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE). |
| [RAKE.Go](https://github.com/afjoseph/RAKE.Go) | 39 | 5 | 2016-12-17 | 2 months ago | Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE). |
| [porter2](https://github.com/zentures/porter2) | 33 | 3 | 2015-01-21 | 3 years ago | Really fast Porter 2 stemmer. |
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
| [go-eco](https://github.com/ThePaw/go-eco) | 4 | 0 | 2015-09-14 | 3 years ago | Similarity, dissimilarity and distance matrices; diversity, equitability and inequality measures; species richness estimators; coenocline models. |
| [gotokenizer](https://github.com/xujiajun/gotokenizer) | 4 | 1 | 2018-10-11 | 1 month ago | A tokenizer based on the dictionary and Bigram language models for Golang. (Now only support chinese segmentation) |

## Networking
        
*Libraries for working with various layers of the network.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [kcptun](https://github.com/xtaci/kcptun) | 9780 | 581 | 2016-02-26 | 5 hours ago | Extremely simple & fast udp tunnel based on KCP protocol. |
| [fasthttp](https://github.com/valyala/fasthttp) | 8160 | 329 | 2015-10-19 | 3 days ago | Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http. |
| [dns](https://github.com/miekg/dns) | 3463 | 155 | 2010-08-04 | 12 hours ago | Go library for working with DNS. |
| [httplab](https://github.com/gchaincl/httplab) | 3305 | 66 | 2017-02-09 | 4 months ago | HTTPLabs let you inspect HTTP requests and forge responses. |
| [gopacket](https://github.com/google/gopacket) | 2530 | 130 | 2015-03-17 | 46 minutes ago | Go library for packet processing with libpcap bindings. |
| [quic-go](https://github.com/lucas-clemente/quic-go) | 2447 | 141 | 2016-04-07 | 4 hours ago | An implementation of the QUIC protocol in pure Go. |
| [kcp-go](https://github.com/xtaci/kcp-go) | 1910 | 133 | 2015-06-16 | 1 month ago | KCP - Fast and Reliable ARQ Protocol. |
| [gobgp](https://github.com/osrg/gobgp) | 1557 | 124 | 2014-09-14 | 12 hours ago | BGP implemented in the Go Programming Language. |
| [webrtc](https://github.com/pions/webrtc) | 1131 | 64 | 2018-05-19 | 6 hours ago | A pure Go implementation of the WebRTC API. |
| [ssh](https://github.com/gliderlabs/ssh) | 974 | 38 | 2016-10-04 | 1 week ago | Higher-level API for building SSH servers (wraps crypto/ssh). |
| [water](https://github.com/songgao/water) | 751 | 34 | 2013-03-26 | 4 days ago | Simple TUN/TAP library. |
| [sftp](https://github.com/pkg/sftp) | 654 | 47 | 2013-11-05 | 4 weeks ago | Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt. |
| [fortio](https://github.com/fortio/fortio) | 649 | 28 | 2017-10-10 | 2 days ago | Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC. |
| [go-getter](https://github.com/hashicorp/go-getter) | 627 | 89 | 2015-10-13 | 2 hours ago | Go library for downloading files or directories from various sources using a URL. |
| [nff-go](https://github.com/intel-go/nff-go) | 579 | 66 | 2017-03-30 | 18 hours ago | Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF). |
| [mdns](https://github.com/hashicorp/mdns) | 496 | 92 | 2014-01-30 | 1 month ago | Simple mDNS (Multicast DNS) client/server library in Golang. |
| [lhttp](https://github.com/fanux/lhttp) | 474 | 55 | 2015-12-29 | 11 months ago | Powerful websocket framework, build your IM server more easily. |
| [grab](https://github.com/cavaliercoder/grab) | 464 | 14 | 2016-01-05 | 1 week ago | Go package for managing file downloads. |
| [ftp](https://github.com/jlaffaye/ftp) | 441 | 23 | 2011-05-07 | 5 days ago | Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959). |
| [gosnmp](https://github.com/soniah/gosnmp) | 398 | 37 | 2012-08-27 | 2 weeks ago | Native Go library for performing SNMP actions. |
| [gotcp](https://github.com/gansidui/gotcp) | 389 | 42 | 2014-04-13 | 1 year ago | Go package for quickly writing tcp applications. |
| [cidranger](https://github.com/yl2chen/cidranger) | 350 | 11 | 2017-08-21 | 1 month ago | Fast IP to CIDR lookup for Go. |
| [gopcap](https://github.com/akrennmair/gopcap) | 338 | 25 | 2009-11-19 | 2 years ago | Go wrapper for libpcap. |
| [peerdiscovery](https://github.com/schollz/peerdiscovery) | 335 | 13 | 2018-04-23 | 1 week ago | Pure Go library for cross-platform local peer discovery using UDP multicast. |
| [go-stun](https://github.com/ccding/go-stun) | 294 | 13 | 2013-08-18 | 7 months ago | Go implementation of the STUN client (RFC 3489 and RFC 5389). |
| [raw](https://github.com/mdlayher/raw) | 252 | 9 | 2015-07-07 | 2 days ago | Package raw enables reading and writing data at the device driver level for a network interface. |
| [tcp_server](https://github.com/firstrow/tcp_server) | 251 | 16 | 2014-10-14 | 6 months ago | Go library for building tcp servers faster. |
| [buffstreams](https://github.com/StabbyCutyou/buffstreams) | 217 | 13 | 2015-06-30 | 1 year ago | Streaming protocolbuffer data over TCP made easy. |
| [buffstreams](https://github.com/StabbyCutyou/buffstreams) | 217 | 13 | 2015-06-30 | 1 year ago | Streaming protocolbuffer data over TCP made easy. |
| [stun](https://github.com/gortc/stun) | 193 | 12 | 2016-04-25 | 14 hours ago | Go implementation of RFC 5389 STUN protocol. |
| [stun](https://github.com/gortc/stun) | 193 | 12 | 2016-04-25 | 2 weeks ago | Go implementation of RFC 5389 STUN protocol. |
| [winrm](https://github.com/masterzen/winrm) | 193 | 18 | 2013-12-31 | 1 week ago | Go WinRM client to remotely execute commands on Windows machines. |
| [arp](https://github.com/mdlayher/arp) | 174 | 11 | 2015-07-07 | 4 months ago | Package arp implements the ARP protocol, as described in RFC 826. |
| [ethernet](https://github.com/mdlayher/ethernet) | 168 | 11 | 2015-07-03 | 4 months ago | Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. |
| [utp](https://github.com/anacrolix/utp) | 147 | 15 | 2015-03-20 | 1 year ago | Go uTP micro transport protocol implementation. |
| [canopus](https://github.com/zubairhamed/canopus) | 128 | 15 | 2015-02-24 | 11 months ago | CoAP Client/Server implementation (RFC 7252). |
| [sslb](https://github.com/eduardonunesp/sslb) | 110 | 7 | 2015-10-19 | 1 year ago | It's a Super Simples Load Balancer, just a little project to achieve some kind of performance. |
| [jazigo](https://github.com/udhos/jazigo) | 108 | 11 | 2016-06-08 | 6 months ago | Jazigo is a tool written in Go for retrieving configuration for multiple network devices. |
| [gnxi](https://github.com/google/gnxi) | 81 | 20 | 2017-09-26 | 2 days ago | A collection of tools for Network Management that use the gNMI and gNOI protocols. |
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
| [heimdall](https://github.com/gojektech/heimdall) | 815 | 41 | 2018-01-19 | 2 weeks ago | An enchanced http client with retry and hystrix capabilities. |
| [sling](https://github.com/dghubble/sling) | 796 | 28 | 2015-04-02 | 1 day ago | Sling is a Go HTTP client library for creating and sending API requests. |
| [gentleman](https://github.com/h2non/gentleman) | 612 | 16 | 2016-02-22 | 1 month ago | Full-featured plugin-driven HTTP client library. |
| [pester](https://github.com/sethgrid/pester) | 302 | 5 | 2015-05-20 | 1 month ago | Go HTTP client calls with retries, backoff, and concurrency. |
| [goreq](https://github.com/smallnest/goreq) | 93 | 7 | 2015-09-17 | 7 months ago | Enhanced simplified HTTP client based on gorequest. |
| [rq](https://github.com/ddo/rq) | 24 | 2 | 2017-12-26 | 4 months ago | A nicer interface for golang stdlib HTTP client. |

## OpenGL
        
*Libraries for using OpenGL in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [glfw](https://github.com/go-gl/glfw) | 599 | 35 | 2013-05-19 | 2 weeks ago | Go bindings for GLFW 3. |
| [gl](https://github.com/go-gl/gl) | 579 | 40 | 2015-02-22 | 2 months ago | Go bindings for OpenGL (generated via glow). |
| [mathgl](https://github.com/go-gl/mathgl) | 266 | 22 | 2013-02-13 | 5 days ago | Pure Go math package specialized for 3D math, with inspiration from GLM. |
| [gl](https://github.com/goxjs/gl) | 130 | 13 | 2015-05-18 | 4 months ago | Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android). |
| [glfw](https://github.com/goxjs/glfw) | 59 | 5 | 2014-12-28 | 4 months ago | Go cross-platform glfw library for creating an OpenGL context and receiving events. |

## ORM
        
*Libraries that implement Object-Relational Mapping or datamapping techniques.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gorm](https://github.com/jinzhu/gorm) | 12426 | 393 | 2013-10-25 | 4 days ago | The fantastic ORM library for Golang, aims to be developer friendly. |
| [xorm](https://github.com/go-xorm/xorm) | 4452 | 241 | 2013-05-09 | 2 weeks ago | Simple and powerful ORM for Go. |
| [gorp](https://github.com/go-gorp/gorp) | 2975 | 116 | 2012-01-05 | 3 months ago | Go Relational Persistence, ORM-ish library for Go. |
| [pg](https://github.com/go-pg/pg) | 2454 | 76 | 2013-04-24 | 23 hours ago | PostgreSQL ORM with focus on PostgreSQL specific features and performance. |
| [sqlboiler](https://github.com/volatiletech/sqlboiler) | 1926 | 66 | 2016-02-21 | 1 week ago | ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema. |
| [db](https://github.com/upper/db) | 1661 | 59 | 2013-10-23 | 1 month ago | Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers. |
| [reform](https://github.com/go-reform/reform) | 737 | 31 | 2016-02-25 | 3 weeks ago | Better ORM for Go, based on non-empty interfaces and code generation. |
| [pop](https://github.com/gobuffalo/pop) | 584 | 24 | 2018-02-08 | 1 hour ago | Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite. |
| [qbs](https://github.com/coocood/qbs) | 526 | 45 | 2013-02-02 | 1 year ago | Stands for Query By Struct. A Go ORM. |
| [go-queryset](https://github.com/jirfag/go-queryset) | 400 | 19 | 2017-09-04 | 1 month ago | 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM. |
| [zoom](https://github.com/albrow/zoom) | 230 | 14 | 2013-07-17 | 8 months ago | Blazing-fast datastore and querying engine built on Redis. |
| [grimoire](https://github.com/Fs02/grimoire) | 104 | 3 | 2018-03-06 | 1 day ago | Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3). |
| [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) | 100 | 8 | 2017-12-28 | 3 months ago | A flexible and powerful SQL string builder library plus a zero-config ORM. |
| [go-store](https://github.com/gosuri/go-store) | 92 | 10 | 2015-03-22 | 2 years ago | Simple and fast Redis backed key-value store library for Go. |
| [gomodel](https://github.com/cosiner/gomodel) | 62 | 5 | 2015-04-16 | 1 year ago | Lightweight, fast, orm-like library helps interactive with database. |
| [marlow](https://github.com/dadleyy/marlow) | 47 | 4 | 2017-09-15 | 5 months ago | Generated ORM from project structs for compile time safety assurances. |
| [lore](https://github.com/abrahambotros/lore) | 4 | 1 | 2017-04-29 | 1 year ago | Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go. |

## Package Management
        
*Unofficial libraries for package and dependency management.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [dep](https://github.com/golang/dep) | 11638 | 302 | 2016-10-07 | 1 day ago | Go dependency tool. |
| [glide](https://github.com/Masterminds/glide) | 7504 | 197 | 2014-07-09 | 1 week ago | Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip. |
| [godep](https://github.com/tools/godep) | 5602 | 158 | 2013-05-01 | 10 months ago | dependency tool for go, godep helps build packages reproducibly by fixing their dependencies. |
| [govendor](https://github.com/kardianos/govendor) | 4356 | 107 | 2015-04-12 | 3 months ago | Go Package Manager. Go vendor tool that works with the standard vendor file. |
| [gopm](https://github.com/gpmgo/gopm) | 2146 | 77 | 2013-05-15 | 1 week ago | Go Package Manager. |
| [gom](https://github.com/mattn/gom) | 1346 | 40 | 2013-09-11 | 9 months ago | Go Manager - bundle for go. |
| [gpm](https://github.com/pote/gpm) | 1201 | 32 | 2013-09-05 | 1 year ago | Barebones dependency manager for Go. |
| [goop](https://github.com/petejkim/goop) | 773 | 37 | 2014-06-18 | 3 years ago | Simple dependency manager for Go (golang), inspired by Bundler. |
| [goop](https://github.com/petejkim/goop) | 773 | 37 | 2014-06-18 | 3 years ago | Simple dependency manager for Go (golang), inspired by Bundler. |
| [nut](https://github.com/jingweno/nut) | 248 | 9 | 2015-01-23 | 3 years ago | Vendor Go dependencies. |
| [johnny-deps](https://github.com/VividCortex/johnny-deps) | 213 | 23 | 2013-07-19 | 4 years ago | Minimal dependency version using Git. |
| [gigo](https://github.com/LyricalSecurity/gigo) | 196 | 5 | 2015-05-01 | 3 months ago | PIP-like dependency tool for golang, with support for private repositories and hashes. |
| [VenGO](https://github.com/DamnWidget/VenGO) | 115 | 8 | 2014-10-17 | 2 years ago | create and manage exportable isolated go virtual environments. |
| [mvn-golang](https://github.com/raydac/mvn-golang) | 76 | 7 | 2016-03-24 | 4 days ago | plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure. |
| [gop](https://github.com/lunny/gop) | 50 | 7 | 2017-02-18 | 3 weeks ago | Build and manage your Go applications out of GOPATH. |

## Query Language
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [graphql](https://github.com/graphql-go/graphql) | 4448 | 142 | 2015-07-19 | 1 week ago | Implementation of GraphQL for Go. |
| [graphql-go](https://github.com/graph-gophers/graphql-go) | 2438 | 91 | 2016-10-18 | 1 week ago | GraphQL server with a focus on ease of use. |
| [graphql-go](https://github.com/graph-gophers/graphql-go) | 2437 | 91 | 2016-10-18 | 1 week ago | GraphQL server with a focus on ease of use. |
| [gojsonq](https://github.com/thedevsaddam/gojsonq) | 729 | 20 | 2018-05-20 | 1 week ago | A simple Go package to Query over JSON Data. |
| [jsonql](https://github.com/elgs/jsonql) | 194 | 8 | 2015-12-29 | 4 months ago | JSON query expression library in Golang. |
| [rql](https://github.com/a8m/rql) | 87 | 5 | 2018-06-06 | 2 days ago | Resource Query Language for REST API. |
| [graphql](https://github.com/tmc/graphql) | 50 | 10 | 2015-04-19 | 1 year ago | graphql parser + utilities. |
| [jsonslice](https://github.com/bhmj/jsonslice) | 16 | 1 | 2018-05-02 | 4 days ago | Jsonpath queries with advanced filters. |

## Resource Embedding
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [statik](https://github.com/rakyll/statik) | 1751 | 32 | 2014-02-04 | 1 day ago | Embeds static files into a Go executable. |
| [packr](https://github.com/gobuffalo/packr) | 1687 | 37 | 2017-03-16 | 6 days ago | The simple and easy way to embed static files into Go binaries. |
| [go.rice](https://github.com/GeertJohan/go.rice) | 1515 | 43 | 2013-10-24 | 19 hours ago | go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy. |
| [vfsgen](https://github.com/shurcooL/vfsgen) | 490 | 19 | 2015-05-18 | 2 weeks ago | Generates a vfsdata.go file that statically implements the given virtual filesystem. |
| [esc](https://github.com/mjibson/esc) | 407 | 12 | 2014-01-26 | 3 weeks ago | Embeds files into Go programs and provides http.FileSystem interfaces to them. |
| [fileb0x](https://github.com/UnnoTed/fileb0x) | 387 | 14 | 2016-01-24 | 1 month ago | Simple tool to embed files in go with focus on "customization" and ease to use. |
| [go-resources](https://github.com/omeid/go-resources) | 148 | 6 | 2015-02-21 | 1 year ago | Unfancy resources embedding with Go. |
| [statics](https://github.com/go-playground/statics) | 51 | 3 | 2015-10-07 | 2 years ago | Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks. |
| [templify](https://github.com/wlbr/templify) | 15 | 2 | 2016-05-23 | 1 month ago | Embed external template files into Go code to create single file binaries. |
| [go-embed](https://github.com/pyros2097/go-embed) | 12 | 1 | 2015-12-07 | 2 years ago | Generates go code to embed resource files into your library or executable. |

## Science and Data Analysis
        
*Libraries for scientific computing and data analyzing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gonum](https://github.com/gonum/gonum) | 2455 | 86 | 2017-03-25 | 1 hour ago | Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more. |
| [streamtools](https://github.com/nytlabs/streamtools) | 1311 | 77 | 2013-07-06 | 3 years ago | general purpose, graphical tool for dealing with streams of data. |
| [stats](https://github.com/montanaflynn/stats) | 1250 | 36 | 2014-12-16 | 1 month ago | Statistics package with common functions missing from the Golang standard library. |
| [gosl](https://github.com/cpmech/gosl) | 1234 | 62 | 2015-02-10 | 6 days ago | Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more. |
| [plot](https://github.com/gonum/plot) | 1072 | 50 | 2013-07-23 | 14 hours ago | gonum/plot provides an API for building and drawing plots in Go. |
| [go-dsp](https://github.com/mjibson/go-dsp) | 604 | 28 | 2011-11-02 | 10 months ago | Digital Signal Processing for Go. |
| [goraph](https://github.com/gyuho/goraph) | 574 | 38 | 2014-02-27 | 1 year ago | Pure Go graph theory library(data structure, algorith visualization). |
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
| [gocomplex](https://github.com/varver/gocomplex) | 5 | 0 | 2015-06-25 | 3 years ago | Complex number library for the Go programming language. |
| [go-gt](https://github.com/ThePaw/go-gt) | 5 | 0 | 2015-09-14 | 3 years ago | Graph theory algorithms written in "Go" language. |
| [rootfinding](https://github.com/khezen/rootfinding) | 2 | 1 | 2018-10-31 | 1 month ago | root-finding algorithms library for finding roots of quadratic functions. |

## Security
        
*Libraries that are used to help make your application more secure.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [lego](https://github.com/xenolf/lego) | 3178 | 89 | 2015-06-08 | 17 hours ago | Pure Go ACME client library and CLI tool (for use with Let's Encrypt). |
| [acme](https://github.com/hlandau/acme) | 1654 | 68 | 2015-11-15 | 6 months ago | ACME (Let's Encrypt) client tool with automatic renewal. |
| [cameradar](https://github.com/Ullaakut/cameradar) | 1616 | 95 | 2016-05-20 | 1 month ago | Tool and library to remotely hack RTSP streams from surveillance cameras. |
| [secure](https://github.com/unrolled/secure) | 1117 | 33 | 2014-05-21 | 2 weeks ago | HTTP middleware for Go that facilitates some quick security wins. |
| [memguard](https://github.com/awnumar/memguard) | 905 | 31 | 2017-04-22 | 1 week ago | A pure Go library for handling sensitive values in memory. |
| [nacl](https://github.com/kevinburke/nacl) | 443 | 13 | 2017-07-21 | 2 months ago | Go implementation of the NaCL set of API's. |
| [acra](https://github.com/cossacklabs/acra) | 375 | 28 | 2016-11-15 | 10 hours ago | Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system. |
| [badactor](https://github.com/jaredfolkins/badactor) | 237 | 8 | 2014-12-13 | 1 year ago | In-memory, application-driven jailer built in the spirit of fail2ban. |
| [passlib](https://github.com/hlandau/passlib) | 212 | 11 | 2014-12-22 | 2 months ago | Futureproof password hashing library. |
| [ssh-vault](https://github.com/ssh-vault/ssh-vault) | 191 | 10 | 2016-09-29 | 2 months ago | encrypt/decrypt using ssh keys. |
| [simple-scrypt](https://github.com/elithrar/simple-scrypt) | 147 | 6 | 2015-04-14 | 9 months ago | Scrypt package with a simple, obvious API and automatic cost calibration built-in. |
| [go-yara](https://github.com/hillu/go-yara) | 123 | 17 | 2015-01-25 | 1 week ago | Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)". |
| [argon2pw](https://github.com/raja/argon2pw) | 67 | 3 | 2018-03-13 | 6 months ago | Argon2 password hash generation with constant-time password comparison. |
| [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) | 25 | 1 | 2017-10-20 | 2 weeks ago | A probably paranoid package for securely hashing and encrypting passwords. |
| [goArgonPass](https://github.com/dwin/goArgonPass) | 7 | 2 | 2018-05-30 | 1 week ago | Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations. |
| [jwc](https://github.com/khezen/jwc) | 4 | 1 | 2018-06-28 | 1 month ago | JSON Web Cryptography library. |

## Serialization
        
*Libraries and tools for binary serialization.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [go](https://github.com/json-iterator/go) | 4461 | 179 | 2016-11-30 | 1 week ago | High-performance 100% compatible drop-in replacement of "encoding/json". |
| [protobuf](https://github.com/golang/protobuf) | 4155 | 188 | 2014-11-24 | 3 days ago | Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers. |
| [protobuf](https://github.com/gogo/protobuf) | 2419 | 80 | 2014-12-03 | 4 hours ago | Protocol Buffers for Go with Gadgets. |
| [mapstructure](https://github.com/mitchellh/mapstructure) | 2047 | 44 | 2013-05-20 | 5 days ago | Go library for decoding generic map values into native Go structures. |
| [go](https://github.com/ugorji/go) | 1118 | 47 | 2013-05-30 | 2 weeks ago | High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support. |
| [colfer](https://github.com/pascaldekloe/colfer) | 441 | 34 | 2015-09-06 | 6 months ago | Code generation for the Colfer binary format. |
| [go-capnproto](https://github.com/glycerine/go-capnproto) | 270 | 11 | 2013-11-08 | 1 month ago | Cap'n Proto library and parser for go. |
| [csvutil](https://github.com/jszwec/csvutil) | 267 | 7 | 2017-10-30 | 1 week ago | High Performance, idiomatic CSV record encoding and decoding to native Go structures. |
| [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) | 114 | 9 | 2012-12-23 | 4 months ago | GoLang library for working with PHP session format and PHP Serialize/Unserialize functions. |
| [structomap](https://github.com/danhper/structomap) | 84 | 7 | 2015-05-13 | 3 years ago | Library to easily and dynamically generate maps from static structures. |
| [structomap](https://github.com/danhper/structomap) | 84 | 7 | 2015-05-13 | 3 years ago | Library to easily and dynamically generate maps from static structures. |
| [bambam](https://github.com/glycerine/bambam) | 60 | 3 | 2014-09-17 | 2 years ago | generator for Cap'n Proto schemas from go. |
| [asn1](https://github.com/Logicalis/asn1) | 36 | 8 | 2016-02-29 | 1 year ago | Asn.1 BER and DER encoding library for golang. |
| [asn1](https://github.com/Logicalis/asn1) | 36 | 8 | 2016-02-29 | 1 year ago | Asn.1 BER and DER encoding library for golang. |
| [fwencoder](https://github.com/o1egl/fwencoder) | 6 | 1 | 2017-12-25 | 1 year ago | Fixed width file parser (encoding and decoding library) for Go. |
| [bel](https://github.com/32leaves/bel) | 0 | 1 | 2019-02-21 | 1 week ago | Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC. |

## Template Engines
        
*Libraries and tools for templating and lexing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gofpdf](https://github.com/jung-kurt/gofpdf) | 2749 | 79 | 2015-03-13 | 5 days ago | PDF document generator with high level support for text, drawing and images. |
| [pongo2](https://github.com/flosch/pongo2) | 1390 | 55 | 2013-08-23 | 2 weeks ago | Django-like template-engine for Go. |
| [quicktemplate](https://github.com/valyala/quicktemplate) | 1218 | 49 | 2016-03-07 | 1 month ago | Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it. |
| [hero](https://github.com/shiyanhui/hero) | 1135 | 38 | 2017-01-15 | 4 months ago | Hero is a handy, fast and powerful go template engine. |
| [mustache](https://github.com/hoisie/mustache) | 950 | 35 | 2009-12-31 | 9 months ago | Go implementation of the Mustache template language. |
| [amber](https://github.com/eknkc/amber) | 803 | 23 | 2012-11-01 | 6 months ago | Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade. |
| [ace](https://github.com/yosssi/ace) | 750 | 23 | 2014-07-13 | 8 months ago | Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold. |
| [gorazor](https://github.com/sipin/gorazor) | 649 | 57 | 2014-05-01 | 5 months ago | Razor view engine for Golang. |
| [jet](https://github.com/CloudyKit/jet) | 547 | 22 | 2016-04-01 | 1 day ago | Jet template engine. |
| [ego](https://github.com/benbjohnson/ego) | 401 | 16 | 2014-02-24 | 1 month ago | Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled. |
| [raymond](https://github.com/aymerick/raymond) | 313 | 12 | 2015-04-22 | 1 month ago | Complete handlebars implementation in Go. |
| [fasttemplate](https://github.com/valyala/fasttemplate) | 250 | 17 | 2015-08-19 | 22 hours ago | Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/). |
| [soy](https://github.com/robfig/soy) | 141 | 13 | 2013-12-15 | 4 days ago | Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/). |
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
| [testify](https://github.com/stretchr/testify) | 6888 | 134 | 2012-10-17 | 2 days ago | Sacred extension to the standard go testing package. |
| [goconvey](https://github.com/smartystreets/goconvey) | 4295 | 143 | 2013-08-21 | 2 days ago | BDD-style framework with web UI and live reload. |
| [httpexpect](https://github.com/gavv/httpexpect) | 1035 | 32 | 2016-04-30 | 7 months ago | Concise, declarative, and easy to use end-to-end HTTP and REST API testing. |
| [go-cmp](https://github.com/google/go-cmp) | 814 | 21 | 2017-07-08 | 5 days ago | Package for comparing Go values in tests. |
| [baloo](https://github.com/h2non/baloo) | 601 | 12 | 2016-05-30 | 1 month ago | Expressive and versatile end-to-end HTTP API testing made easy. |
| [goblin](https://github.com/franela/goblin) | 592 | 16 | 2013-09-19 | 5 months ago | Mocha like testing framework fo Go. |
| [godog](https://github.com/DATA-DOG/godog) | 566 | 27 | 2015-06-10 | 6 days ago | Cucumber or Behat like BDD framework for Go. |
| [go-vcr](https://github.com/dnaeon/go-vcr) | 273 | 7 | 2015-12-14 | 2 months ago | Record and replay your HTTP interactions for fast, deterministic and accurate tests. |
| [testfixtures](https://github.com/go-testfixtures/testfixtures) | 260 | 5 | 2016-04-05 | 2 months ago | A helper for Rails' like test fixtures to test database applications. |
| [frisby](https://github.com/verdverm/frisby) | 240 | 8 | 2015-09-15 | 11 months ago | REST API testing framework. |
| [gofight](https://github.com/appleboy/gofight) | 231 | 13 | 2016-03-29 | 1 month ago | API Handler Testing for Golang Router framework. |
| [go-mutesting](https://github.com/zimmski/go-mutesting) | 226 | 7 | 2014-12-27 | 3 months ago | Mutation testing for Go source code. |
| [go-carpet](https://github.com/msoap/go-carpet) | 189 | 5 | 2016-02-28 | 3 days ago | Tool for viewing test coverage in terminal. |
| [charlatan](https://github.com/percolate/charlatan) | 185 | 39 | 2017-10-07 | 5 months ago | Tool to generate fake interface implementations for tests. |
| [gospec](https://github.com/luontola/gospec) | 110 | 4 | 2009-11-24 | 4 years ago | BDD-style testing framework for the Go programming language. |
| [gospec](https://github.com/luontola/gospec) | 110 | 4 | 2009-11-24 | 4 years ago | BDD-style testing framework for the Go programming language. |
| [gotest.tools](https://github.com/gotestyourself/gotest.tools) | 84 | 5 | 2017-08-09 | 1 day ago | A collection of packages to augment the go testing package and support common patterns. |
| [endly](https://github.com/viant/endly) | 74 | 11 | 2017-08-29 | 4 days ago | Declarative end to end functional testing. |
| [cupaloy](https://github.com/bradleyjkemp/cupaloy) | 64 | 3 | 2017-08-08 | 1 month ago | Simple snapshot testing addon for your test framework. |
| [dbcleaner](https://github.com/khaiql/dbcleaner) | 62 | 3 | 2017-01-18 | 1 week ago | Clean database for testing purpose, inspired by `database_cleaner` in Ruby. |
| [wstest](https://github.com/posener/wstest) | 56 | 2 | 2017-04-01 | 1 year ago | Websocket client for unit-testing a websocket http.Handler. |
| [gospecify](https://github.com/stesla/gospecify) | 50 | 5 | 2009-11-20 | 7 years ago | This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec. |
| [restit](https://github.com/go-restit/restit) | 49 | 6 | 2014-06-25 | 2 years ago | Go micro framework to help writing RESTful API integration test. |
| [restit](https://github.com/go-restit/restit) | 49 | 6 | 2014-06-25 | 2 years ago | Go micro framework to help writing RESTful API integration test. |
| [go-testdeep](https://github.com/maxatome/go-testdeep) | 41 | 2 | 2018-05-26 | 2 days ago | Extremely flexible golang deep comparison, extends the go testing package. |
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
| [chromedp](https://github.com/chromedp/chromedp) | 2786 | 112 | 2017-01-24 | 1 day ago | a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol. |
| [chromedp](https://github.com/chromedp/chromedp) | 2785 | 112 | 2017-01-24 | 1 day ago | a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol. |
| [go-fuzz](https://github.com/dvyukov/go-fuzz) | 2581 | 81 | 2015-04-15 | 12 hours ago | Randomized testing system. |
| [mock](https://github.com/golang/mock) | 2168 | 53 | 2015-06-13 | 1 week ago | Mocking framework for the Go programming language. |
| [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) | 1328 | 15 | 2014-02-07 | 4 days ago | Mock SQL driver for testing database interactions. |
| [hoverfly](https://github.com/SpectoLabs/hoverfly) | 1311 | 57 | 2015-12-01 | 5 days ago | HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI. |
| [selenoid](https://github.com/aerokube/selenoid) | 954 | 85 | 2016-08-22 | 2 days ago | alternative Selenium hub server that launches browsers within containers. |
| [gock](https://github.com/h2non/gock) | 682 | 16 | 2016-03-03 | 5 days ago | Versatile HTTP mocking made easy. |
| [gofuzz](https://github.com/google/gofuzz) | 475 | 19 | 2014-08-01 | 1 year ago | Library for populating go objects with random values. |
| [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) | 316 | 11 | 2014-05-21 | 2 weeks ago | Tool for generating self-contained mock objects. |
| [cdp](https://github.com/mafredri/cdp) | 266 | 13 | 2017-03-12 | 3 weeks ago | Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it. |
| [tavor](https://github.com/zimmski/tavor) | 202 | 12 | 2014-05-18 | 4 months ago | Generic fuzzing and delta-debugging framework. |
| [ggr](https://github.com/aerokube/ggr) | 189 | 22 | 2016-06-16 | 5 days ago | a lightweight server that routes and proxies Selenium Wedriver requests to multiple Selenium hubs. |
| [minimock](https://github.com/gojuno/minimock) | 162 | 5 | 2016-08-04 | 6 days ago | Mock generator for Go interfaces. |
| [go-txdb](https://github.com/DATA-DOG/go-txdb) | 122 | 8 | 2015-07-08 | 3 weeks ago | Single transaction based database driver mainly for testing purposes. |
| [govcr](https://github.com/seborama/govcr) | 67 | 2 | 2016-07-11 | 2 days ago | HTTP mock for Golang: record and replay HTTP interactions for offline testing. |
| [mockhttp](https://github.com/tv42/mockhttp) | 22 | 1 | 2011-06-12 | 4 years ago | Mock object for Go http.ResponseWriter. |

## Text Processing
        
*Libraries for parsing and manipulating texts.*

### Specific Formats
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [colly](https://github.com/gocolly/colly) | 6999 | 216 | 2017-09-29 | 2 days ago | Fast and Elegant Scraping Framework for Gophers. |
| [colly](https://github.com/gocolly/colly) | 6990 | 215 | 2017-09-29 | 2 days ago | Fast and Elegant Scraping Framework for Gophers. |
| [goquery](https://github.com/PuerkitoBio/goquery) | 6868 | 248 | 2012-08-29 | 1 month ago | GoQuery brings a syntax and a set of features similar to jQuery to the Go language. |
| [blackfriday](https://github.com/russross/blackfriday) | 3580 | 86 | 2011-05-28 | 3 weeks ago | Markdown processor in Go. |
| [toml](https://github.com/BurntSushi/toml) | 2471 | 76 | 2013-02-26 | 3 weeks ago | TOML configuration format (encoder/decoder with reflection). |
| [go-humanize](https://github.com/dustin/go-humanize) | 1723 | 29 | 2012-01-13 | 1 month ago | Formatters for time, numbers, and memory size to human readable format. |
| [sh](https://github.com/mvdan/sh) | 1700 | 38 | 2016-01-16 | 1 week ago | Shell parser and formatter. |
| [bluemonday](https://github.com/microcosm-cc/bluemonday) | 1107 | 30 | 2013-11-21 | 2 months ago | HTML Sanitizer. |
| [inject](https://github.com/facebookgo/inject) | 1086 | 42 | 2013-10-21 | 1 month ago | Package inject provides a reflect based injector. |
| [gofeed](https://github.com/mmcdole/gofeed) | 998 | 33 | 2016-01-23 | 2 weeks ago | Parse RSS and Atom feeds in Go. |
| [commonregex](https://github.com/mingrammer/commonregex) | 532 | 20 | 2017-03-23 | 2 months ago | A collection of common regular expressions for Go. |
| [slug](https://github.com/gosimple/slug) | 327 | 9 | 2014-03-31 | 1 month ago | URL-friendly slugify with multiple languages support. |
| [mxj](https://github.com/clbanning/mxj) | 314 | 22 | 2014-02-03 | 1 day ago | Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages. |
| [gommon](https://github.com/labstack/gommon) | 272 | 19 | 2015-03-13 | 1 month ago | Format bytes to string. |
| [gographviz](https://github.com/awalterschulze/gographviz) | 245 | 10 | 2015-03-15 | 1 week ago | Parses the Graphviz DOT language. |
| [dataflowkit](https://github.com/slotix/dataflowkit) | 227 | 15 | 2017-02-09 | 1 day ago | Web scraping Framework to turn websites into structured data. |
| [gotext](https://github.com/leonelquinteros/gotext) | 210 | 5 | 2016-06-20 | 2 weeks ago | GNU gettext utilities for Go. |
| [go-runewidth](https://github.com/mattn/go-runewidth) | 182 | 12 | 2013-06-21 | 2 months ago | Functions to get fixed width of the character or string. |
| [goq](https://github.com/andrewstuart/goq) | 128 | 7 | 2017-02-20 | 9 months ago | Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery). |
| [htmlquery](https://github.com/antchfx/htmlquery) | 86 | 3 | 2017-12-05 | 3 weeks ago | An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression. |
| [go-nmea](https://github.com/adrianmo/go-nmea) | 77 | 3 | 2015-07-22 | 2 weeks ago | NMEA parser library for the Go language. |
| [sdp](https://github.com/gortc/sdp) | 51 | 7 | 2016-05-13 | 2 weeks ago | SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)]. |
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
| [syndfeed](https://github.com/zhengchun/syndfeed) | 4 | 1 | 2017-04-07 | 11 months ago | A syndication feed for Atom 1.0 and RSS 2.0. |
| [doi](https://github.com/hscells/doi) | 4 | 1 | 2017-08-02 | 1 year ago | Document object identifier (doi) parser in Go. |
| [encoding](https://github.com/mickep76/encoding) | 3 | 1 | 2018-04-07 | 5 months ago | Package provides a generic interface to encoders and decodersa. |
| [encoding](https://github.com/mickep76/encoding) | 3 | 1 | 2018-04-07 | 5 months ago | Package provides a generic interface to encoders and decodersa. |

### Utility
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [xurls](https://github.com/mvdan/xurls) | 408 | 13 | 2015-01-12 | 2 weeks ago | Extract urls from text. |
| [gotabulate](https://github.com/bndr/gotabulate) | 192 | 8 | 2014-08-21 | 2 years ago | Easily pretty-print your tabular data with Go. |
| [radix](https://github.com/yourbasic/radix) | 61 | 2 | 2017-06-09 | 1 year ago | fast string sorting algorithm. |
| [parth](https://github.com/codemodus/parth) | 27 | 3 | 2015-04-07 | 1 month ago | URL path segmentation parsing. |
| [xj2go](https://github.com/stackerzzq/xj2go) | 15 | 1 | 2017-09-19 | 2 months ago | Convert xml or json to go struct. |
| [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) | 11 | 1 | 2018-09-09 | 3 months ago | A sanitization-based swear filter for Go. |
| [kace](https://github.com/codemodus/kace) | 9 | 1 | 2015-06-05 | 6 months ago | Common case conversions covering common initialisms. |
| [parseargs-go](https://github.com/txgruppi/parseargs-go) | 4 | 1 | 2016-02-24 | 2 years ago | string argument parser that understands quotes and backslashes. |
| [parseargs-go](https://github.com/txgruppi/parseargs-go) | 4 | 1 | 2016-02-24 | 2 years ago | string argument parser that understands quotes and backslashes. |

## Third-party APIs
        
*Libraries for accessing third party APIs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [aws-sdk-go](https://github.com/aws/aws-sdk-go) | 4572 | 254 | 2014-12-05 | 15 hours ago | The official AWS SDK for the Go programming language. |
| [go-github](https://github.com/google/go-github) | 4370 | 175 | 2013-05-25 | 3 days ago | Go library for accessing the GitHub REST API v3. |
| [slack](https://github.com/nlopes/slack) | 2102 | 43 | 2015-01-24 | 19 hours ago | Slack API in Go. |
| [google-api-go-client](https://github.com/googleapis/google-api-go-client) | 1744 | 127 | 2014-11-25 | 16 hours ago | Auto-generated Google APIs for Go. |
| [google-api-go-client](https://github.com/googleapis/google-api-go-client) | 1744 | 127 | 2014-11-25 | 3 days ago | Auto-generated Google APIs for Go. |
| [google-cloud-go](https://github.com/googleapis/google-cloud-go) | 1558 | 210 | 2014-05-09 | 5 days ago | Google Cloud APIs Go Client Library. |
| [google-cloud-go](https://github.com/googleapis/google-cloud-go) | 1558 | 212 | 2014-05-09 | 14 hours ago | Google Cloud APIs Go Client Library. |
| [anaconda](https://github.com/ChimeraCoder/anaconda) | 954 | 23 | 2013-03-05 | 3 weeks ago | Go client library for the Twitter 1.1 API. |
| [stripe-go](https://github.com/stripe/stripe-go) | 847 | 35 | 2014-06-06 | 11 hours ago | Go client for the Stripe API. |
| [discordgo](https://github.com/bwmarrin/discordgo) | 805 | 50 | 2015-11-02 | 1 day ago | Go bindings for the Discord Chat API. |
| [facebook](https://github.com/huandu/facebook) | 724 | 86 | 2012-07-29 | 2 months ago | Go Library that supports the Facebook Graph API. |
| [minio-go](https://github.com/minio/minio-go) | 629 | 41 | 2015-05-02 | 1 day ago | Minio Go Library for Amazon S3 compatible cloud storage. |
| [go-twitter](https://github.com/dghubble/go-twitter) | 612 | 27 | 2015-04-12 | 1 day ago | Go client library for the Twitter v1.1 APIs. |
| [githubv4](https://github.com/shurcooL/githubv4) | 402 | 18 | 2017-05-27 | 1 month ago | Go library for accessing the GitHub GraphQL API v4. |
| [githubv4](https://github.com/shurcooL/githubv4) | 402 | 18 | 2017-05-27 | 1 month ago | Go library for accessing the GitHub GraphQL API v4. |
| [webhooks](https://github.com/go-playground/webhooks) | 280 | 12 | 2015-10-26 | 1 day ago | Webhook receiver for GitHub and Bitbucket. |
| [geo-golang](https://github.com/codingsince1985/geo-golang) | 280 | 11 | 2014-12-04 | 1 month ago | Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs. |
| [PayPal-Go-SDK](https://github.com/logpacker/PayPal-Go-SDK) | 268 | 16 | 2015-10-14 | 5 days ago | Wrapper for PayPal payment API. |
| [go-marathon](https://github.com/gambol99/go-marathon) | 184 | 12 | 2015-02-11 | 5 days ago | Go library for interacting with Mesosphere's Marathon PAAS. |
| [ethrpc](https://github.com/onrik/ethrpc) | 156 | 9 | 2017-01-24 | 1 day ago | Go bindings for Ethereum JSON RPC API. |
| [gostorm](https://github.com/jsgilmore/gostorm) | 117 | 11 | 2013-07-22 | 1 year ago | GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. |
| [hipchat](https://github.com/daneharrigan/hipchat) | 112 | 6 | 2013-04-28 | 1 year ago | A golang package to communicate with HipChat over XMPP. |
| [hipchat](https://github.com/andybons/hipchat) | 110 | 8 | 2012-10-21 | 2 years ago | This project implements a golang client library for the Hipchat API. |
| [medium-sdk-go](https://github.com/Medium/medium-sdk-go) | 108 | 104 | 2015-09-27 | 4 months ago | Golang SDK for Medium's OAuth2 API. |
| [go-trending](https://github.com/andygrunwald/go-trending) | 96 | 7 | 2015-07-04 | 4 months ago | Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github. |
| [trello](https://github.com/adlio/trello) | 86 | 5 | 2016-09-24 | 3 days ago | Go wrapper for the Trello API. |
| [cachet](https://github.com/andygrunwald/cachet) | 63 | 5 | 2015-10-31 | 11 months ago | Go client library for [Cachet (open source status page system)](https://cachethq.io/). |
| [megos](https://github.com/andygrunwald/megos) | 56 | 5 | 2015-10-02 | 10 months ago | Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster. |
| [clarifai-go](https://github.com/Clarifai/clarifai-go) | 56 | 37 | 2015-09-29 | 1 year ago | Go client library for interfacing with the Clarifai API. |
| [clarifai-go](https://github.com/Clarifai/clarifai-go) | 56 | 37 | 2015-09-29 | 1 year ago | Go client library for interfacing with the Clarifai API. |
| [pushover](https://github.com/gregdel/pushover) | 53 | 2 | 2015-02-19 | 2 weeks ago | Go wrapper for the Pushover API. |
| [gads](https://github.com/emiddleton/gads) | 41 | 8 | 2014-01-20 | 2 years ago | Google Adwords Unofficial API. |
| [go-amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) | 37 | 1 | 2016-11-15 | 11 months ago | Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html). |
| [go-xkcd](https://github.com/nishanths/go-xkcd) | 36 | 3 | 2016-02-26 | 2 years ago | Go client for the xkcd API. |
| [go-circleci](https://github.com/jszwedko/go-circleci) | 36 | 3 | 2015-08-15 | 1 month ago | Go client library for interacting with CircleCI's API. |
| [gomusicbrainz](https://github.com/michiwend/gomusicbrainz) | 33 | 8 | 2014-09-11 | 4 months ago | Go MusicBrainz WS2 client library. |
| [fcm](https://github.com/maddevsio/fcm) | 32 | 4 | 2017-01-06 | 3 months ago | Go library for Firebase Cloud Messaging. |
| [gcm](https://github.com/Aorioli/gcm) | 29 | 3 | 2015-11-10 | 3 years ago | Go library for Google Cloud Messaging. |
| [golyrics](https://github.com/mamal72/golyrics) | 28 | 3 | 2016-11-18 | 8 months ago | Golyrics is a Go library to fetch music lyrics data from the Wikia website. |
| [uptimerobot](https://github.com/bitfield/uptimerobot) | 28 | 3 | 2018-05-29 | 2 months ago | Go wrapper and command-line client for the Uptime Robot v2 API. |
| [mixpanel](https://github.com/dukex/mixpanel) | 27 | 3 | 2014-05-20 | 5 months ago | Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications. |
| [wit-go](https://github.com/wit-ai/wit-go) | 27 | 6 | 2018-08-20 | 1 month ago | Go client for wit.ai HTTP API. |
| [gami](https://github.com/bit4bit/gami) | 26 | 4 | 2014-05-15 | 8 months ago | Go library for Asterisk Manager Interface. |
| [translate](https://github.com/nuveo/translate) | 25 | 20 | 2015-07-13 | 3 years ago | Go online translation package. |
| [translate](https://github.com/nuveo/translate) | 25 | 20 | 2015-07-13 | 3 years ago | Go online translation package. |
| [igdb](https://github.com/Henry-Sarabia/igdb) | 24 | 1 | 2017-08-24 | 1 day ago | Go client for the [Internet Game Database API](https://api.igdb.com/). |
| [go-shopify](https://github.com/rapito/go-shopify) | 19 | 1 | 2014-10-28 | 1 year ago | Go Library to make CRUD request to the Shopify API. |
| [go-unsplash](https://github.com/hbagdi/go-unsplash) | 19 | 1 | 2017-01-19 | 7 months ago | Go client library for the [Unsplash.com](https://unsplash.com) API. |
| [go-spotify](https://github.com/rapito/go-spotify) | 16 | 1 | 2014-10-30 | 1 year ago | Go Library to access Spotify WEB API. |
| [patreon-go](https://github.com/mxpv/patreon-go) | 15 | 4 | 2017-08-07 | 2 months ago | Go library for Patreon API. |
| [go-myanimelist](https://github.com/nstratos/go-myanimelist) | 15 | 2 | 2015-05-03 | 2 years ago | Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api). |
| [brewerydb](https://github.com/naegelejd/brewerydb) | 15 | 2 | 2015-04-15 | 3 years ago | Go library for accessing the BreweryDB API. |
| [go-twitch](https://github.com/knspriggs/go-twitch) | 15 | 4 | 2016-06-29 | 1 year ago | Go client for interacting with the Twitch v3 API. |
| [codeship-go](https://github.com/codeship/codeship-go) | 14 | 18 | 2017-09-09 | 1 month ago | Go client library for interacting with Codeship's API v2. |
| [go-steam](https://github.com/sostronk/go-steam) | 13 | 10 | 2014-11-24 | 11 months ago | Go Library to interact with Steam game servers. |
| [textbelt](https://github.com/farmergreg/textbelt) | 12 | 2 | 2015-09-02 | 3 years ago | Go client for the textbelt.com txt messaging API. |
| [textbelt](https://github.com/farmergreg/textbelt) | 12 | 2 | 2015-09-02 | 3 years ago | Go client for the textbelt.com txt messaging API. |
| [go-google-analytics](https://github.com/chonthu/go-google-analytics) | 11 | 2 | 2015-06-01 | 3 years ago | Simple wrapper for easy google analytics reporting. |
| [go-tmdb](https://github.com/jbrodriguez/go-tmdb) | 11 | 1 | 2014-10-20 | 3 years ago | Simple golang package to communicate with [themoviedb.org](https://themoviedb.org). |
| [smitego](https://github.com/sergiotapia/smitego) | 9 | 0 | 2013-12-11 | 4 years ago | Go package to wraps access to the Smite game API. |
| [go-hacknews](https://github.com/PaulRosset/go-hacknews) | 9 | 2 | 2017-08-11 | 1 year ago | Tiny Go client for HackerNews API. |
| [simples3](https://github.com/rhnvrm/simples3) | 9 | 0 | 2018-12-06 | 1 month ago | Simple no frills AWS S3 Library using REST with V4 Signing written in Go. |
| [rrdaclient](https://github.com/Omie/rrdaclient) | 8 | 1 | 2014-09-16 | 4 years ago | Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP. |
| [go-sptrans](https://github.com/sergioaugrod/go-sptrans) | 8 | 2 | 2017-09-11 | 1 year ago | Go client library for the SPTrans Olho Vivo API. |
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
| [fzf](https://github.com/junegunn/fzf) | 20005 | 316 | 2013-10-24 | 2 hours ago | Command-line fuzzy finder written in Go. |
| [hub](https://github.com/github/hub) | 14934 | 310 | 2009-12-06 | 3 days ago | wrap git commands with additional functionality to interact with github from the terminal. |
| [delve](https://github.com/go-delve/delve) | 10695 | 353 | 2014-05-21 | 8 hours ago | Go debugger. |
| [delve](https://github.com/go-delve/delve) | 10679 | 353 | 2014-05-21 | 4 days ago | Go debugger. |
| [ctop](https://github.com/bcicen/ctop) | 8125 | 153 | 2016-12-27 | 1 month ago | [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics. |
| [wuzz](https://github.com/asciimoo/wuzz) | 7950 | 176 | 2017-01-31 | 2 months ago | Interactive cli tool for HTTP inspection. |
| [sqlx](https://github.com/jmoiron/sqlx) | 5815 | 164 | 2013-01-29 | 1 week ago | provides a set of extensions on top of the excellent built-in database/sql package. |
| [peco](https://github.com/peco/peco) | 5194 | 131 | 2014-06-06 | 1 week ago | Simplistic interactive filtering tool. |
| [usql](https://github.com/xo/usql) | 4543 | 111 | 2017-03-02 | 1 week ago | usql is a universal command-line interface for SQL databases. |
| [usql](https://github.com/xo/usql) | 4539 | 111 | 2017-03-02 | 1 week ago | usql is a universal command-line interface for SQL databases. |
| [goreleaser](https://github.com/goreleaser/goreleaser) | 3814 | 61 | 2016-12-22 | 1 day ago | Deliver Go binaries as fast and easily as possible. |
| [godropbox](https://github.com/dropbox/godropbox) | 3645 | 250 | 2014-06-23 | 1 week ago | Common libraries for writing Go services/applications from Dropbox. |
| [go-torch](https://github.com/uber/go-torch) | 3539 | 2403 | 2015-07-22 | 3 months ago | Stochastic flame graph profiler for Go programs. |
| [realize](https://github.com/oxequa/realize) | 2866 | 68 | 2016-07-12 | 1 month ago | Go build system with file watchers and live reload. Run, build and watch file changes with custom paths. |
| [realize](https://github.com/oxequa/realize) | 2864 | 68 | 2016-07-12 | 1 month ago | Go build system with file watchers and live reload. Run, build and watch file changes with custom paths. |
| [goreporter](https://github.com/360EntSecGroup-Skylar/goreporter) | 2330 | 93 | 2017-03-27 | 4 months ago | Golang tool that does static analysis, unit testing, code review and generate code quality report. |
| [goreporter](https://github.com/360EntSecGroup-Skylar/goreporter) | 2328 | 93 | 2017-03-27 | 4 months ago | Golang tool that does static analysis, unit testing, code review and generate code quality report. |
| [panicparse](https://github.com/maruel/panicparse) | 1827 | 37 | 2015-02-02 | 7 months ago | Groups similar goroutines and colorizes stack dump. |
| [task](https://github.com/go-task/task) | 1706 | 43 | 2017-02-27 | 1 day ago | simple "Make" alternative. |
| [minify](https://github.com/tdewolff/minify) | 1697 | 52 | 2014-05-21 | 1 month ago | Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats. |
| [hystrix-go](https://github.com/afex/hystrix-go) | 1650 | 73 | 2013-12-15 | 3 months ago | Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker. |
| [resty](https://github.com/go-resty/resty) | 1593 | 59 | 2015-08-29 | 6 days ago | Simple HTTP and REST client for Go inspired by Ruby rest-client. |
| [mmake](https://github.com/tj/mmake) | 1439 | 34 | 2017-02-16 | 1 year ago | Modern Make. |
| [storm](https://github.com/asdine/storm) | 1257 | 45 | 2016-01-10 | 2 weeks ago | Simple and powerful toolkit for BoltDB. |
| [mole](https://github.com/davrodpin/mole) | 1242 | 27 | 2018-10-04 | 6 hours ago | cli app to easily create ssh tunnels. |
| [mc](https://github.com/minio/mc) | 942 | 53 | 2015-01-16 | 1 hour ago | Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems. |
| [profile](https://github.com/pkg/profile) | 890 | 33 | 2014-10-22 | 4 months ago | Simple profiling support package for Go. |
| [boilr](https://github.com/tmrts/boilr) | 859 | 27 | 2015-12-20 | 7 months ago | Blazingly fast CLI tool for creating projects from boilerplate templates. |
| [filetype](https://github.com/h2non/filetype) | 843 | 17 | 2015-09-24 | 1 week ago | Small package to infer the file type checking the magic numbers signature. |
| [circuitbreaker](https://github.com/rubyist/circuitbreaker) | 732 | 19 | 2014-07-18 | 1 month ago | Circuit Breakers in Go. |
| [go-funk](https://github.com/thoas/go-funk) | 727 | 25 | 2016-12-30 | 3 months ago | Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...). |
| [mergo](https://github.com/imdario/mergo) | 706 | 11 | 2013-03-12 | 12 hours ago | Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements. |
| [spinner](https://github.com/briandowns/spinner) | 677 | 15 | 2014-12-13 | 3 weeks ago | Go package to easily provide a terminal spinner with options. |
| [gtm](https://github.com/git-time-metric/gtm) | 667 | 28 | 2016-06-20 | 8 months ago | Simple, seamless, lightweight time tracking for Git. |
| [jump](https://github.com/gsamokovarov/jump) | 585 | 15 | 2015-08-17 | 5 days ago | Jump helps you navigate faster by learning your habits. |
| [immortal](https://github.com/immortal/immortal) | 576 | 15 | 2016-07-01 | 2 months ago | \*nix cross-platform (OS agnostic) supervisor. |
| [htcat](https://github.com/htcat/htcat) | 472 | 16 | 2013-08-05 | 1 week ago | Parallel and Pipelined HTTP GET Utility. |
| [gopencils](https://github.com/bndr/gopencils) | 419 | 14 | 2014-06-23 | 2 weeks ago | Small and simple package to easily consume REST APIs. |
| [go-dry](https://github.com/ungerik/go-dry) | 417 | 15 | 2014-02-28 | 10 months ago | DRY (don't repeat yourself) package for Go. |
| [godaemon](https://github.com/VividCortex/godaemon) | 383 | 32 | 2013-08-02 | 3 years ago | Utility to write daemons. |
| [request](https://github.com/mozillazg/request) | 340 | 13 | 2014-12-21 | 8 months ago | Go HTTP Requests for Humans™. |
| [go-rate](https://github.com/beefsack/go-rate) | 277 | 12 | 2014-08-25 | 11 months ago | Timed rate limiter for Go. |
| [ergo](https://github.com/cristianoliveira/ergo) | 276 | 4 | 2017-08-20 | 3 months ago | The management of multiple local services running over different ports made easy. |
| [circuit](https://github.com/cep21/circuit) | 262 | 11 | 2017-12-24 | 4 days ago | An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern. |
| [gohper](https://github.com/cosiner/gohper) | 248 | 19 | 2015-03-24 | 1 year ago | Various tools/modules help for development. |
| [koazee](https://github.com/wesovilabs/koazee) | 231 | 9 | 2018-11-09 | 1 day ago | Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays. |
| [clockwork](https://github.com/jonboulle/clockwork) | 195 | 5 | 2014-09-10 | 1 month ago | A simple fake clock for golang. |
| [deepcopier](https://github.com/ulule/deepcopier) | 188 | 20 | 2015-07-25 | 1 year ago | Simple struct copying for Go. |
| [serve](https://github.com/syntaqx/serve) | 171 | 6 | 2019-01-11 | 1 month ago | A static http server anywhere you need. |
| [go-trigger](https://github.com/sadlil/go-trigger) | 168 | 12 | 2015-10-19 | 1 year ago | Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project. |
| [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) | 157 | 3 | 2016-11-08 | 1 year ago | go:generate tool for wrapping symbols exported by golang plugins (1.8 only). |
| [rerun](https://github.com/ivpusic/rerun) | 151 | 7 | 2014-12-10 | 11 months ago | Recompiling and rerunning go apps when source changes. |
| [moldova](https://github.com/StabbyCutyou/moldova) | 148 | 5 | 2016-01-30 | 1 year ago | Utility for generating random data based on an input template. |
| [robustly](https://github.com/VividCortex/robustly) | 130 | 21 | 2013-07-08 | 11 months ago | Runs functions resiliently, catching and restarting panics. |
| [gotenv](https://github.com/subosito/gotenv) | 125 | 4 | 2013-08-27 | 9 months ago | Load environment variables from `.env` or any `io.Reader` in Go. |
| [death](https://github.com/vrecan/death) | 121 | 4 | 2015-03-09 | 4 months ago | Managing go application shutdown with signals. |
| [apm](https://github.com/topfreegames/apm) | 119 | 14 | 2015-11-19 | 2 years ago | Process manager for Golang applications with an HTTP API. |
| [util](https://github.com/shomali11/util) | 107 | 6 | 2017-05-24 | 9 months ago | Collection of useful utility functions. (strings, concurrency, manipulations, ...). |
| [chyle](https://github.com/antham/chyle) | 102 | 6 | 2016-11-18 | 4 days ago | Changelog generator using a git repository with multiple configuration possibilities. |
| [retry](https://github.com/kamilsk/retry) | 98 | 0 | 2016-11-03 | 7 hours ago | Functional mechanism based on context to perform actions repetitively until successful. |
| [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) | 98 | 5 | 2015-10-13 | 1 month ago | XML Sitemap generator written in Go. |
| [lrserver](https://github.com/jaschaephraim/lrserver) | 98 | 5 | 2014-07-15 | 1 year ago | LiveReload server for Go. |
| [onecache](https://github.com/adelowo/onecache) | 86 | 6 | 2017-04-15 | 4 days ago | Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc). |
| [pm](https://github.com/VividCortex/pm) | 67 | 20 | 2013-11-18 | 2 months ago | Process (i.e. goroutine) manager with an HTTP API. |
| [unis](https://github.com/esemplastic/unis) | 67 | 5 | 2017-05-06 | 1 year ago | Common Architecture™ for String Utilities in Go. |
| [netbug](https://github.com/e-dard/netbug) | 65 | 1 | 2015-03-06 | 3 years ago | Easy remote profiling of your services. |
| [xferspdy](https://github.com/monmohan/xferspdy) | 65 | 4 | 2015-05-22 | 2 years ago | Xferspdy provides binary diff and patch library in golang. |
| [mimetype](https://github.com/gabriel-vasile/mimetype) | 61 | 6 | 2018-07-02 | 5 days ago | Package for MIME type detection based on magic numbers. |
| [go-health](https://github.com/Talento90/go-health) | 60 | 2 | 2018-02-14 | 8 months ago | Health package simplifies the way you add health check to your services. |
| [toolbox](https://github.com/viant/toolbox) | 60 | 15 | 2016-06-14 | 4 days ago | Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer. |
| [multitick](https://github.com/VividCortex/multitick) | 58 | 20 | 2013-12-11 | 2 years ago | Multiplexor for aligned tickers. |
| [repeat](https://github.com/ssgreg/repeat) | 53 | 4 | 2017-11-22 | 1 month ago | Go implementation of different backoff strategies useful for retrying operations and heartbeating. |
| [mssqlx](https://github.com/linxGnu/mssqlx) | 52 | 5 | 2016-12-26 | 2 days ago | Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind. |
| [abutil](https://github.com/bahlo/abutil) | 51 | 3 | 2015-08-26 | 3 years ago | Collection of often-used Golang helpers. |
| [minquery](https://github.com/icza/minquery) | 44 | 3 | 2016-11-16 | 4 months ago | MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off). |
| [go-astitodo](https://github.com/asticode/go-astitodo) | 43 | 2 | 2016-10-18 | 5 months ago | Parse TODOs in your GO code. |
| [golog](https://github.com/mlimaloureiro/golog) | 42 | 3 | 2016-01-09 | 1 month ago | Easy and lightweight CLI tool to time track your tasks. |
| [mimemagic](https://github.com/zRedShift/mimemagic) | 39 | 2 | 2018-10-12 | 2 months ago | Pure Go ultra performant MIME sniffing library/utility. |
| [goback](https://github.com/carlescere/goback) | 39 | 1 | 2015-03-14 | 1 year ago | Go simple exponential backoff package. |
| [intrinsic](https://github.com/mengzhuo/intrinsic) | 37 | 4 | 2017-06-13 | 1 year ago | Use x86 SIMD without writing any assembly code. |
| [handy](https://github.com/miguelpragier/handy) | 37 | 3 | 2018-06-13 | 12 minutes ago | Many utilities and helpers like string handlers/formatters and validators. |
| [copy-pasta](https://github.com/jutkko/copy-pasta) | 34 | 4 | 2017-01-28 | 6 months ago | Universal multi-workstation clipboard that uses S3 like backend for the storage. |
| [gaper](https://github.com/maxcnunes/gaper) | 34 | 0 | 2018-06-16 | 1 month ago | Builds and restarts a Go project when it crashes or some watched file changes. |
| [golarm](https://github.com/msempere/golarm) | 32 | 1 | 2015-08-15 | 3 years ago | Fire alarms with system events. |
| [retry](https://github.com/thedevsaddam/retry) | 31 | 1 | 2018-02-26 | 11 months ago | Simple and easy retry mechanism package for Go. |
| [myhttp](https://github.com/inancgumus/myhttp) | 30 | 1 | 2017-09-13 | 10 months ago | Simple API to make HTTP GET requests with timeout support. |
| [gpath](https://github.com/tenntenn/gpath) | 25 | 2 | 2017-05-24 | 1 year ago | Library to simplify access struct fields with Go's expression in reflection. |
| [retry-go](https://github.com/rafaeljesus/retry-go) | 24 | 1 | 2017-06-10 | 4 months ago | Retrying made simple and easy for golang. |
| [goplaceholder](https://github.com/michiwend/goplaceholder) | 21 | 2 | 2014-10-12 | 3 years ago | a small golang lib to generate placeholder images. |
| [rclient](https://github.com/zpatrick/rclient) | 21 | 2 | 2017-02-28 | 6 months ago | Readable, flexible, simple-to-use client for REST APIs. |
| [goseaweedfs](https://github.com/linxGnu/goseaweedfs) | 19 | 3 | 2017-07-20 | 3 weeks ago | SeaweedFS client library with almost full features. |
| [ugo](https://github.com/alxrm/ugo) | 19 | 1 | 2016-02-18 | 2 years ago | ugo is slice toolbox with concise syntax for Go. |
| [generate](https://github.com/go-playground/generate) | 15 | 2 | 2015-11-15 | 2 years ago | runs go generate recursively on a specified path or environment variable and can filter by regex. |
| [dlog](https://github.com/kirillDanshin/dlog) | 15 | 2 | 2016-07-05 | 1 year ago | Compile-time controlled logger to make your release smaller without removing debug calls. |
| [go-httpheader](https://github.com/mozillazg/go-httpheader) | 14 | 2 | 2017-06-24 | 4 months ago | Go library for encoding structs into Header fields. |
| [filler](https://github.com/yaronsumel/filler) | 14 | 1 | 2017-04-05 | 1 year ago | small utility to fill structs using "fill" tag. |
| [okrun](https://github.com/xta/okrun) | 13 | 2 | 2014-10-01 | 4 years ago | go run error steamroller. |
| [evaluator](https://github.com/nullne/evaluator) | 13 | 1 | 2017-04-28 | 1 year ago | Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend. |
| [gostrutils](https://github.com/ik5/gostrutils) | 12 | 3 | 2018-09-19 | 1 month ago | Collections of string manipulation and conversion functions. |
| [rerate](https://github.com/abo/rerate) | 12 | 4 | 2016-05-24 | 1 year ago | Redis-based rate counter and rate limiter for Go. |
| [structs](https://github.com/PumpkinSeed/structs) | 11 | 1 | 2017-08-26 | 1 year ago | Implement simple functions to manipulate structs. |
| [fastlz](https://github.com/digitalcrab/fastlz) | 11 | 1 | 2014-12-26 | 4 years ago | Wrap over [FastLz](http://fastlz.org/) (free, open-source, portable real-time compression library) for GoLang. |
| [pgo](https://github.com/arthurkushman/pgo) | 10 | 2 | 2018-12-26 | 1 week ago | Convenient functions for PHP community. |
| [command](https://github.com/txgruppi/command) | 9 | 2 | 2015-08-24 | 2 years ago | Command pattern for Go with thread safe serial and parallel dispatcher. |
| [retry](https://github.com/shafreeck/retry) | 8 | 1 | 2018-07-18 | 6 months ago | A pretty simple library to ensure your work to be done. |
| [backscanner](https://github.com/icza/backscanner) | 7 | 2 | 2017-10-18 | 4 months ago | A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward. |
| [filter](https://github.com/gookit/filter) | 6 | 2 | 2018-09-26 | 2 months ago | provide filtering, sanitizing, and conversion of Go data. |
| [ghokin](https://github.com/antham/ghokin) | 6 | 2 | 2018-08-03 | 5 days ago | Parallelized formatter with no external dependencies for gherkin (cucumber, behat...). |
| [silk](https://github.com/chrispassas/silk) | 4 | 1 | 2018-12-18 | 1 month ago | Read silk netflow files. |
| [mimesniffer](https://github.com/aofei/mimesniffer) | 3 | 1 | 2018-12-20 | 1 week ago | A MIME type sniffer for Go. |
| [sslice](https://github.com/yaa110/sslice) | 2 | 1 | 2018-11-17 | 3 months ago | Create a slice which is always sorted. |
| [retry](https://github.com/percolate/retry) | 2 | 29 | 2018-06-16 | 4 months ago | A simple but highly configurable retry package for Go. |
| [ctxutil](https://github.com/posener/ctxutil) | 1 | 0 | 2018-07-30 | 2 days ago | A collection of utility functions for contexts. |
| [slicer](https://github.com/leaanthony/slicer) | 1 | 1 | 2019-01-10 | 3 days ago | Makes working with slices easier. |
| [olaf](https://github.com/btnguyen2k/olaf) | 1 | 1 | 2019-01-03 | 1 month ago | Twitter Snowflake implemented in Go. |

## UUID
        
*Libraries for working with UUIDs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [uuid](https://github.com/gofrs/uuid) | 440 | 17 | 2018-07-13 | 5 days ago | Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid. |
| [wuid](https://github.com/edwingeng/wuid) | 237 | 13 | 2018-01-27 | 8 months ago | An extremely fast unique number generator, 10-135 times faster than UUID. |
| [Goid](https://github.com/JakeHL/Goid) | 18 | 3 | 2017-05-19 | 2 weeks ago | Generate and Parse RFC4122 compliant V4 UUIDs. |
| [Goid](https://github.com/JakeHL/Goid) | 18 | 3 | 2017-05-19 | 2 weeks ago | Generate and Parse RFC4122 compliant V4 UUIDs. |
| [uuid](https://github.com/agext/uuid) | 9 | 1 | 2016-02-03 | 1 week ago | Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier. |

## Validation
        
*Libraries for validation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [govalidator](https://github.com/asaskevich/govalidator) | 3190 | 91 | 2014-06-20 | 2 days ago | Validators and sanitizers for strings, numerics, slices and structs. |
| [validator](https://github.com/go-playground/validator) | 2756 | 65 | 2015-02-13 | 11 hours ago | Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving. |
| [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) | 910 | 20 | 2016-06-22 | 3 months ago | Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags. |
| [govalidator](https://github.com/thedevsaddam/govalidator) | 570 | 16 | 2017-09-14 | 3 weeks ago | Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. |
| [validate](https://github.com/gookit/validate) | 40 | 3 | 2018-07-16 | 2 days ago | Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features. |
| [validate](https://github.com/gobuffalo/validate) | 11 | 7 | 2018-02-11 | 5 days ago | This package provides a framework for writing validations for Go applications. |

## Version Control
        
*Libraries for version control.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [git2go](https://github.com/libgit2/git2go) | 1280 | 52 | 2013-03-06 | 1 week ago | Go bindings for libgit2. |
| [gh](https://github.com/rjeczalik/gh) | 67 | 5 | 2015-03-09 | 4 months ago | Scriptable server and net/http middleware for GitHub Webhooks. |
| [go-vcs](https://github.com/sourcegraph/go-vcs) | 65 | 28 | 2013-06-02 | 1 month ago | manipulate and inspect VCS repositories in Go. |
| [hgo](https://github.com/beyang/hgo) | 12 | 4 | 2014-06-18 | 3 years ago | Hgo is a collection of Go packages providing read-access to local Mercurial repositories. |

## Video
        
*Libraries for manipulating video.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [goav](https://github.com/giorgisio/goav) | 625 | 40 | 2015-05-21 | 2 months ago | Comphrensive Go bindings for FFmpeg. |
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
| [gin](https://github.com/gin-gonic/gin) | 24853 | 1025 | 2014-06-17 | 8 hours ago | Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity. |
| [beego](https://github.com/astaxie/beego) | 19150 | 1224 | 2012-02-29 | 10 hours ago | beego is an open-source, high-performance web framework for the Go programming language. |
| [echo](https://github.com/labstack/echo) | 13092 | 528 | 2015-03-02 | 20 hours ago | High performance, minimalist Go web framework. |
| [revel](https://github.com/revel/revel) | 10793 | 560 | 2011-12-09 | 1 day ago | High-productivity web framework for the Go language. |
| [goa](https://github.com/goadesign/goa) | 3294 | 164 | 2014-12-05 | 23 minutes ago | Framework for developing microservices based on the design of Ruby's Praxis. |
| [goa](https://github.com/goadesign/goa) | 3293 | 164 | 2014-12-05 | 1 day ago | Framework for developing microservices based on the design of Ruby's Praxis. |
| [go-json-rest](https://github.com/ant0ine/go-json-rest) | 3263 | 163 | 2013-02-19 | 2 weeks ago | Quick and easy way to setup a RESTful JSON API. |
| [macaron](https://github.com/go-macaron/macaron) | 2676 | 153 | 2014-07-10 | 2 months ago | Macaron is a high productive and modular design web framework in Go. |
| [gizmo](https://github.com/nytimes/gizmo) | 2575 | 107 | 2015-12-16 | 20 hours ago | Microservice toolkit used by the New York Times. |
| [gizmo](https://github.com/nytimes/gizmo) | 2573 | 107 | 2015-12-16 | 1 day ago | Microservice toolkit used by the New York Times. |
| [utron](https://github.com/gernest/utron) | 2122 | 74 | 2015-09-16 | 4 months ago | Lightweight MVC framework for Go(Golang). |
| [go-tigertonic](https://github.com/rcrowley/go-tigertonic) | 993 | 45 | 2013-02-10 | 7 months ago | Go framework for building JSON web services inspired by Dropwizard. |
| [tango](https://github.com/lunny/tango) | 774 | 77 | 2014-12-17 | 3 weeks ago | Micro & pluggable web framework for Go. |
| [traffic](https://github.com/gravityblast/traffic) | 512 | 25 | 2013-08-09 | 3 years ago | Sinatra inspired regexp/pattern mux and web framework for Go. |
| [traffic](https://github.com/gravityblast/traffic) | 512 | 25 | 2013-08-09 | 3 years ago | Sinatra inspired regexp/pattern mux and web framework for Go. |
| [gongular](https://github.com/mustafaakin/gongular) | 410 | 23 | 2016-06-22 | 2 weeks ago | Fast Go web framework with input mapping/validation and (DI) Dependency Injection. |
| [neo](https://github.com/ivpusic/neo) | 387 | 33 | 2015-02-05 | 1 year ago | Neo is minimal and fast Go Web Framework with extremely simple API. |
| [mango](https://github.com/paulbellamy/mango) | 334 | 22 | 2011-05-25 | 1 year ago | Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. |
| [gondola](https://github.com/rainycape/gondola) | 312 | 15 | 2014-07-26 | 2 weeks ago | The web framework for writing faster sites, faster. |
| [air](https://github.com/aofei/air) | 273 | 13 | 2016-07-20 | 2 days ago | An ideally refined web framework for Go. |
| [golf](https://github.com/dinever/golf) | 229 | 19 | 2015-11-18 | 2 years ago | Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library. |
| [gem](https://github.com/go-gem/gem) | 151 | 20 | 2016-11-10 | 1 year ago | Simple and fast web framework, friendly to REST API. |
| [go-rest](https://github.com/ungerik/go-rest) | 111 | 9 | 2012-07-13 | 2 years ago | Small and evil REST framework for Go. |
| [aero](https://github.com/aerogo/aero) | 91 | 13 | 2016-11-09 | 1 day ago | High-performance web framework for Go, reaches top scores in Lighthouse. |
| [golax](https://github.com/fulldump/golax) | 66 | 9 | 2016-01-31 | 9 months ago | A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more. |
| [webgo](https://github.com/bnkamalesh/webgo) | 63 | 3 | 2015-12-16 | 1 month ago | A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc). |
| [hiboot](https://github.com/hidevopsio/hiboot) | 61 | 6 | 2018-03-16 | 1 week ago | hiboot is a high performance web application framework with auto configuration and dependency injection support. |
| [microservice](https://github.com/claygod/microservice) | 51 | 6 | 2016-12-15 | 1 year ago | The framework for the creation of microservices, written in Golang. |
| [yarf](https://github.com/yarf-framework/yarf) | 46 | 3 | 2015-09-02 | 1 year ago | Fast micro-framework designed to build REST APIs and web services in a fast and simple way. |
| [fireball](https://github.com/zpatrick/fireball) | 44 | 4 | 2016-07-20 | 5 months ago | More "natural" feeling web framework. |
| [uadmin](https://github.com/uadmin/uadmin) | 36 | 6 | 2018-10-05 | 2 months ago | Fully featured web framework for Golang, inspired by Django. |
| [api](https://github.com/resoursea/api) | 29 | 5 | 2015-01-25 | 4 years ago | REST framework for quickly writing resource based services. |
| [rex](https://github.com/goanywhere/rex) | 24 | 1 | 2014-10-16 | 1 year ago | Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`. |
| [vox](https://github.com/aisk/vox) | 17 | 1 | 2014-12-24 | 2 months ago | A golang web framework for humans, inspired by Koa heavily. |
| [nio](https://github.com/go-nio/nio) | 15 | 2 | 2018-12-05 | 1 week ago | Modern, minimal and productive Go HTTP framework. |
| [banjo](https://github.com/nsheremet/banjo) | 4 | 1 | 2017-12-09 | 1 year ago | Very simple and fast web framework for Go. |
| [sawsij](https://github.com/jaybill/sawsij) | 2 | 2 | 2016-06-09 | 2 years ago | lightweight, open-source web framework for building high-performance, data-driven web applications. |

### Middlewares
        

#### Actual middlewares
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [tollbooth](https://github.com/didip/tollbooth) | 1082 | 37 | 2015-05-17 | 3 months ago | Rate limit HTTP request handler. |
| [cors](https://github.com/rs/cors) | 1053 | 26 | 2014-10-25 | 2 weeks ago | Easily add CORS capabilities to your API. |
| [go-server-timing](https://github.com/mitchellh/go-server-timing) | 731 | 19 | 2018-02-12 | 6 months ago | Add/parse Server-Timing header. |
| [limiter](https://github.com/ulule/limiter) | 479 | 21 | 2015-10-02 | 1 month ago | Dead simple rate limit middleware for Go. |
| [ln-paywall](https://github.com/philippgille/ln-paywall) | 80 | 2 | 2018-06-30 | 1 week ago | Go middleware for monetizing your API on a per-request basis with Bitcoin and Lightning ⚡️ |
| [xff](https://github.com/sebest/xff) | 69 | 2 | 2014-12-22 | 2 years ago | Handle `X-Forwarded-For` header and friends. |
| [formjson](https://github.com/rs/formjson) | 30 | 2 | 2015-03-20 | 3 years ago | Transparently handle JSON input as a standard form POST. |
| [client-timing](https://github.com/posener/client-timing) | 10 | 1 | 2018-02-23 | 6 days ago | An HTTP client for Server-Timing header. |

#### Libraries for creating HTTP middlewares
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [negroni](https://github.com/urfave/negroni) | 6082 | 236 | 2014-05-19 | 2 weeks ago | Idiomatic HTTP middleware for Golang. |
| [alice](https://github.com/justinas/alice) | 1737 | 54 | 2014-05-25 | 3 months ago | Painless middleware chaining for Go. |
| [render](https://github.com/unrolled/render) | 1198 | 39 | 2014-06-11 | 2 weeks ago | Go package for easily rendering JSON, XML, and HTML template responses. |
| [stats](https://github.com/thoas/stats) | 520 | 16 | 2015-03-06 | 2 weeks ago | Go middleware that stores various information about your web application. |
| [interpose](https://github.com/carbocation/interpose) | 288 | 13 | 2014-07-20 | 2 years ago | Minimalist net/http middleware for golang. |
| [muxchain](https://github.com/stephens2424/muxchain) | 206 | 5 | 2014-05-04 | 1 week ago | Lightweight middleware for net/http. |
| [renderer](https://github.com/thedevsaddam/renderer) | 144 | 5 | 2017-11-08 | 2 months ago | Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go. |
| [rye](https://github.com/InVisionApp/rye) | 91 | 176 | 2016-10-07 | 5 months ago | Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context. |
| [gores](https://github.com/alioygur/gores) | 78 | 5 | 2015-12-25 | 4 months ago | Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs. |
| [chain](https://github.com/codemodus/chain) | 64 | 7 | 2015-05-15 | 6 months ago | Handler wrapper chaining with scoped data (net/context-based "middleware"). |
| [wrap](https://github.com/go-on/wrap) | 55 | 3 | 2014-02-16 | 6 months ago | Small middlewares package for net/http. |
| [catena](https://github.com/codemodus/catena) | 7 | 1 | 2015-07-31 | 6 months ago | http.Handler wrapper catenation (same API as "chain"). |

### Routers
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [httprouter](https://github.com/julienschmidt/httprouter) | 8717 | 285 | 2013-12-05 | 3 weeks ago | High performance router. Use this and the standard http handlers to form a very high performance web framework. |
| [mux](https://github.com/gorilla/mux) | 8273 | 256 | 2012-10-03 | 1 day ago | Powerful URL router and dispatcher for golang. |
| [chi](https://github.com/go-chi/chi) | 5124 | 150 | 2015-10-16 | 1 day ago | Small, fast and expressive HTTP router built on net/context. |
| [web](https://github.com/gocraft/web) | 1365 | 56 | 2013-11-17 | 3 weeks ago | Mux and middleware package in Go. |
| [bone](https://github.com/go-zoo/bone) | 1191 | 36 | 2014-11-19 | 1 month ago | Lightning Fast HTTP Multiplexer. |
| [goji](https://github.com/goji/goji) | 712 | 42 | 2015-11-16 | 2 weeks ago | Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`. |
| [fasthttprouter](https://github.com/buaazp/fasthttprouter) | 660 | 32 | 2015-12-13 | 1 month ago | High performance router forked from `httprouter`. The first router fit for `fasthttp`. |
| [gorouter](https://github.com/xujiajun/gorouter) | 391 | 12 | 2018-01-29 | 1 month ago | A simple and fast HTTP router for Go. |
| [lars](https://github.com/go-playground/lars) | 375 | 18 | 2015-12-25 | 1 year ago | Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks. |
| [httptreemux](https://github.com/dimfeld/httptreemux) | 374 | 21 | 2014-05-15 | 2 months ago | High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter. |
| [siesta](https://github.com/VividCortex/siesta) | 348 | 32 | 2014-09-23 | 2 months ago | Composable framework to write middleware and handlers. |
| [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) | 332 | 29 | 2015-10-27 | 3 months ago | An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs. |
| [vestigo](https://github.com/husobee/vestigo) | 246 | 18 | 2015-09-22 | 1 month ago | Performant, stand-alone, HTTP compliant URL Router for go web applications. |
| [router](https://github.com/gowww/router) | 154 | 5 | 2017-05-25 | 11 months ago | Lightning fast HTTP router fully compatible with the net/http.Handler interface. |
| [alien](https://github.com/gernest/alien) | 97 | 3 | 2016-01-31 | 1 year ago | Lightweight and fast http router from outer space. |
| [Bxog](https://github.com/claygod/Bxog) | 92 | 8 | 2016-05-19 | 2 months ago | Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters. |
| [violetear](https://github.com/nbari/violetear) | 92 | 4 | 2015-06-20 | 5 months ago | Go HTTP router. |
| [xmux](https://github.com/rs/xmux) | 87 | 6 | 2015-12-15 | 1 year ago | High performance muxer based on `httprouter` with `net/context` support. |
| [pure](https://github.com/go-playground/pure) | 75 | 4 | 2016-09-24 | 1 year ago | Is a lightweight HTTP router that sticks to the std "net/http" implementation. |
| [gorouter](https://github.com/vardius/gorouter) | 37 | 3 | 2016-07-14 | 4 weeks ago | GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`. |
| [fastrouter](https://github.com/razonyang/fastrouter) | 18 | 2 | 2017-11-01 | 1 year ago | a fast, flexible HTTP router written in Go. |

## Windows
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [go-ole](https://github.com/go-ole/go-ole) | 489 | 35 | 2011-01-21 | 1 day ago | Win32 OLE implementation for golang. |
| [d3d9](https://github.com/gonutz/d3d9) | 79 | 5 | 2015-12-13 | 2 months ago | Go bindings for Direct3D9. |

## XML
        
*Libraries and tools for manipulating XML.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [xquery](https://github.com/antchfx/xquery) | 141 | 11 | 2016-10-09 | 9 months ago | XQuery lets you extract data from HTML/XML documents using XPath expression. |
| [xpath](https://github.com/antchfx/xpath) | 121 | 8 | 2016-10-09 | 2 days ago | XPath package for Go. |
| [XML-Comp](https://github.com/XML-Comp/XML-Comp) | 15 | 1 | 2016-10-26 | 7 months ago | Simple command line XML comparer that generates diffs of folders, files and tags. |
| [XML-Comp](https://github.com/XML-Comp/XML-Comp) | 15 | 1 | 2016-10-26 | 7 months ago | Simple command line XML comparer that generates diffs of folders, files and tags. |
| [xml2map](https://github.com/sbabiv/xml2map) | 14 | 0 | 2018-08-07 | 1 month ago | XML to MAP converter written Golang. |
| [xmlwriter](https://github.com/shabbyrobe/xmlwriter) | 6 | 1 | 2017-04-11 | 1 month ago | Procedural XML generation API based on libxml2's xmlwriter module. |

# Tools
        
*Go software and plugins.*

## Code Analysis
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gometalinter](https://github.com/alecthomas/gometalinter) | 3551 | 80 | 2014-08-05 | 6 days ago | Metalinter is a tool to automatically apply all static analysis tool and report their output in normalized form. |
| [lint](https://github.com/golang/lint) | 2888 | 100 | 2013-06-03 | 4 days ago | Golint is a linter for Go source code. |
| [go-tools](https://github.com/dominikh/go-tools) | 2229 | 66 | 2017-01-25 | 2 weeks ago | unused checks Go code for unused constants, variables, functions and types. |
| [errcheck](https://github.com/kisielk/errcheck) | 1237 | 20 | 2013-02-25 | 2 months ago | Errcheck is a program for checking for unchecked errors in Go programs. |
| [gcvis](https://github.com/davecheney/gcvis) | 878 | 34 | 2014-07-10 | 1 year ago | Visualise Go program GC trace data in real time. |
| [php-parser](https://github.com/z7zmey/php-parser) | 579 | 27 | 2017-11-07 | 3 hours ago | A Parser for PHP written in Go. |
| [go-critic](https://github.com/go-critic/go-critic) | 484 | 18 | 2018-05-05 | 4 days ago | source code linter that brings checks that are currently not implemented in other linters. |
| [goast-viewer](https://github.com/yuroyoro/goast-viewer) | 329 | 13 | 2014-06-30 | 1 year ago | Web based Golang AST visualizer. |
| [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) | 247 | 7 | 2017-04-13 | 9 months ago | go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects. |
| [unconvert](https://github.com/mdempsky/unconvert) | 244 | 9 | 2016-02-20 | 1 month ago | Remove unnecessary type conversions from Go source. |
| [gostatus](https://github.com/shurcooL/gostatus) | 233 | 6 | 2013-11-27 | 1 month ago | Command line tool, shows the status of repositories that contain Go packages. |
| [apicompat](https://github.com/bradleyfalzon/apicompat) | 161 | 7 | 2016-07-10 | 2 years ago | Checks recent changes to a Go project for backwards incompatible changes. |
| [dupl](https://github.com/mibk/dupl) | 145 | 8 | 2015-05-20 | 4 months ago | Tool for code clone detection. |
| [checkstyle](https://github.com/qiniu/checkstyle) | 92 | 10 | 2014-01-01 | 3 months ago | checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style refered to some points in Go Code Review Comments. |
| [lint](https://github.com/surullabs/lint) | 61 | 6 | 2016-07-09 | 4 months ago | Run linters as part of go test. |
| [validate](https://github.com/mccoyst/validate) | 61 | 6 | 2013-11-22 | 2 years ago | Automatically validates struct fields with tags. |
| [go-outdated](https://github.com/firstrow/go-outdated) | 46 | 1 | 2015-06-29 | 1 month ago | Console application that displays outdated packages. |
| [blanket](https://github.com/verygoodsoftwarenotvirus/blanket) | 11 | 2 | 2017-09-04 | 7 months ago | tarp finds functions and methods without direct unit tests in Go source code. |
| [blanket](https://github.com/verygoodsoftwarenotvirus/blanket) | 11 | 2 | 2017-09-04 | 7 months ago | tarp finds functions and methods without direct unit tests in Go source code. |

## Editor Plugins
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [vim-go](https://github.com/fatih/vim-go) | 9946 | 279 | 2014-03-24 | 5 days ago | Go development plugin for Vim. |
| [gocode](https://github.com/nsf/gocode) | 4600 | 202 | 2010-07-05 | 4 days ago | Autocompletion daemon for the Go programming language. |
| [vscode-go](https://github.com/Microsoft/vscode-go) | 4561 | 210 | 2015-10-15 | 8 hours ago | Extension for Visual Studio Code (VS Code) which provides support for the Go language. |
| [GoSublime](https://github.com/DisposaBoy/GoSublime) | 3137 | 115 | 2011-08-28 | 1 week ago | Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features. |
| [go-plus](https://github.com/joefitzgerald/go-plus) | 1452 | 48 | 2014-03-14 | 1 day ago | Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting. |
| [go-mode.el](https://github.com/dominikh/go-mode.el) | 885 | 58 | 2013-01-31 | 1 month ago | Go mode for GNU/Emacs. |
| [Watch](https://github.com/eaburns/Watch) | 157 | 11 | 2013-08-09 | 11 months ago | Runs a command in an acme win on file changes. |
| [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) | 78 | 5 | 2012-11-26 | 2 years ago | Vim plugin to highlight syntax errors on save. |
| [go-language-server](https://github.com/theia-ide/go-language-server) | 27 | 4 | 2017-11-21 | 1 week ago | A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol. |
| [velour](https://github.com/velour/velour) | 16 | 8 | 2014-06-27 | 4 months ago | IRC client for the acme editor. |
| [gounit-vim](https://github.com/hexdigest/gounit-vim) | 16 | 2 | 2018-02-22 | 4 months ago | Vim plugin for generating Go tests based on the function's or method's signature. |
| [theia-go-extension](https://github.com/theia-ide/theia-go-extension) | 10 | 4 | 2017-11-30 | 1 week ago | Go language support for the Theia IDE. |

## Go Generate Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gotests](https://github.com/cweill/gotests) | 1853 | 56 | 2016-01-19 | 5 days ago | Generate Go tests from your source code. |
| [genny](https://github.com/cheekybits/genny) | 811 | 19 | 2014-10-28 | 5 days ago | Elegant generics for Go. |
| [re2dfa](https://github.com/opennota/re2dfa) | 162 | 10 | 2015-06-20 | 5 months ago | Transform regular expressions into finite state machines and output Go source code. |
| [gocontracts](https://github.com/Parquery/gocontracts) | 45 | 4 | 2018-08-14 | 1 month ago | brings design-by-contract to Go by synchronizing the code with the documentation. |
| [generic](https://github.com/usk81/generic) | 25 | 2 | 2016-06-15 | 6 months ago | flexible data type for Go. |
| [gounit](https://github.com/hexdigest/gounit) | 21 | 4 | 2018-02-05 | 6 months ago | Generate Go tests using your own templates. |

## Go Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [OctoLinker](https://github.com/OctoLinker/OctoLinker) | 3346 | 85 | 2013-12-28 | 4 days ago | Navigate through go files efficiently with the OctoLinker browser extension for GitHub. |
| [OctoLinker](https://github.com/OctoLinker/OctoLinker) | 3346 | 85 | 2013-12-28 | 4 days ago | Navigate through go files efficiently with the OctoLinker browser extension for GitHub. |
| [go-swagger](https://github.com/go-swagger/go-swagger) | 3299 | 98 | 2014-11-17 | 1 day ago | Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API. |
| [go-callvis](https://github.com/TrueFurby/go-callvis) | 1677 | 58 | 2016-09-03 | 1 month ago | Visualize call graph of your Go program using dot format. |
| [richgo](https://github.com/kyoh86/richgo) | 339 | 4 | 2017-01-05 | 1 week ago | Enrich `go test` outputs with text decorations. |
| [depth](https://github.com/KyleBanks/depth) | 312 | 7 | 2017-03-04 | 3 days ago | Visualize dependency trees of any package by analyzing imports. |
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
| [moby](https://github.com/moby/moby) | 52398 | 3287 | 2013-01-19 | 6 hours ago | Collaborative project for the container ecosystem to assemble container-based systems. |
| [kubernetes](https://github.com/kubernetes/kubernetes) | 49132 | 2915 | 2014-06-07 | 6 minutes ago | Container Cluster Manager from Google. |
| [traefik](https://github.com/containous/traefik) | 20784 | 623 | 2015-09-14 | 4 hours ago | Reverse proxy and load balancer with support for multiple backends. |
| [gitea](https://github.com/go-gitea/gitea) | 12669 | 424 | 2016-11-01 | 1 hour ago | Fork of Gogs, entirely community driven. |
| [vegeta](https://github.com/tsenart/vegeta) | 10810 | 279 | 2013-08-13 | 16 hours ago | HTTP load testing tool and library. It's over 9000! |
| [packer](https://github.com/hashicorp/packer) | 8622 | 422 | 2013-03-23 | 1 hour ago | Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. |
| [packer](https://github.com/hashicorp/packer) | 8614 | 422 | 2013-03-23 | 1 day ago | Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. |
| [hey](https://github.com/rakyll/hey) | 5072 | 118 | 2016-09-02 | 3 weeks ago | Hey is a tiny program that sends some load to a web application. |
| [gvm](https://github.com/moovweb/gvm) | 4147 | 142 | 2011-12-03 | 2 months ago | GVM provides an interface to manage Go versions. |
| [webhook](https://github.com/adnanh/webhook) | 3592 | 132 | 2015-01-13 | 5 days ago | Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server. |
| [gaia](https://github.com/gaia-pipeline/gaia) | 3519 | 93 | 2017-12-28 | 16 hours ago | Build powerful pipelines in any programming language. |
| [gox](https://github.com/mitchellh/gox) | 3122 | 70 | 2013-11-17 | 1 week ago | Dead simple, no frills Go cross compile tool. |
| [bosun](https://github.com/bosun-monitor/bosun) | 2788 | 153 | 2013-11-15 | 6 hours ago | Time Series Alerting Framework. |
| [aptly](https://github.com/aptly-dev/aptly) | 1675 | 97 | 2013-12-14 | 1 month ago | aptly is a Debian repository management tool. |
| [aptly](https://github.com/aptly-dev/aptly) | 1674 | 97 | 2013-12-14 | 1 month ago | aptly is a Debian repository management tool. |
| [goxc](https://github.com/laher/goxc) | 1601 | 45 | 2013-02-11 | 8 months ago | build tool for Go, with a focus on cross-compiling and packaging. |
| [fac](https://github.com/mkchoi212/fac) | 1542 | 31 | 2017-12-30 | 5 months ago | Command-line user interface to fix git merge conflicts. |
| [bombardier](https://github.com/codesenberg/bombardier) | 1447 | 39 | 2016-05-29 | 2 days ago | Fast cross-platform HTTP benchmarking tool. |
| [kala](https://github.com/ajvb/kala) | 1304 | 62 | 2015-03-19 | 3 weeks ago | Simplistic, modern, and performant job scheduler. |
| [statusok](https://github.com/sanathp/statusok) | 1063 | 43 | 2015-08-27 | 1 month ago | Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected. |
| [banshee](https://github.com/eleme/banshee) | 1036 | 56 | 2015-12-03 | 4 months ago | Anomalies detection system for periodic metrics. |
| [s3gof3r](https://github.com/rlmcpherson/s3gof3r) | 969 | 32 | 2013-08-02 | 2 months ago | Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3. |
| [go-selfupdate](https://github.com/sanbornm/go-selfupdate) | 631 | 23 | 2013-11-13 | 2 years ago | Enable your Go applications to self update. |
| [skm](https://github.com/TimothyYe/skm) | 494 | 20 | 2017-10-11 | 1 week ago | SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily! |
| [scaleway-cli](https://github.com/scaleway/scaleway-cli) | 490 | 50 | 2015-03-20 | 1 day ago | Manage BareMetal Servers from Command Line (as easily as with Docker). |
| [aurora](https://github.com/xuri/aurora) | 362 | 26 | 2016-10-09 | 1 month ago | Cross-platform web-based Beanstalkd queue server console. |
| [govvv](https://github.com/ahmetb/govvv) | 339 | 11 | 2016-08-03 | 11 months ago | “go build” wrapper to easily add version information into Go binaries. |
| [govvv](https://github.com/ahmetb/govvv) | 339 | 11 | 2016-08-03 | 11 months ago | “go build” wrapper to easily add version information into Go binaries. |
| [gonative](https://github.com/inconshreveable/gonative) | 302 | 8 | 2014-05-01 | 2 years ago | Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages. |
| [mora](https://github.com/emicklei/mora) | 261 | 23 | 2013-07-12 | 2 years ago | REST server for accessing MongoDB documents and meta data. |
| [godbg](https://github.com/sirnewton01/godbg) | 217 | 17 | 2013-08-09 | 7 months ago | Web-based gdb front-end application. |
| [lstags](https://github.com/ivanilves/lstags) | 212 | 13 | 2017-08-15 | 1 month ago | Tool and API to sync Docker images across different registries. |
| [dogo](https://github.com/liudng/dogo) | 204 | 18 | 2014-11-19 | 2 years ago | Monitoring changes in the source file and automatically compile and run (restart). |
| [manssh](https://github.com/xwjdsh/manssh) | 190 | 4 | 2017-10-08 | 8 months ago | manssh is a command line tool for managing your ssh alias config easily. |
| [pomerium](https://github.com/pomerium/pomerium) | 188 | 12 | 2019-01-01 | 8 hours ago | Pomerium is an identity-aware access proxy. |
| [pewpew](https://github.com/bengadbois/pewpew) | 187 | 7 | 2016-10-13 | 4 months ago | Flexible HTTP command line stress tester. |
| [gobrew](https://github.com/cryptojuice/gobrew) | 173 | 5 | 2013-11-13 | 1 year ago | gobrew lets you easily switch between multiple versions of go. |
| [ostent](https://github.com/ostrost/ostent) | 161 | 5 | 2014-03-31 | 11 months ago | collects and displays system metrics and optionally relays to Graphite and/or InfluxDB. |
| [blast](https://github.com/dave/blast) | 160 | 3 | 2017-10-22 | 1 year ago | A simple tool for API load testing and batch jobs. |
| [grapes](https://github.com/yaronsumel/grapes) | 117 | 6 | 2016-09-01 | 1 month ago | Lightweight tool designed to distribute commands over ssh with ease. |
| [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) | 79 | 2 | 2017-03-03 | 4 hours ago | Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`. |
| [kcli](https://github.com/cswank/kcli) | 59 | 6 | 2017-03-26 | 1 month ago | Command line tool for inspecting kafka topics/partitions/messages. |
| [winrm-cli](https://github.com/masterzen/winrm-cli) | 55 | 5 | 2016-05-23 | 1 month ago | Cli tool to remotely execute commands on Windows machines. |
| [go-furnace](https://github.com/go-furnace/go-furnace) | 53 | 2 | 2016-10-09 | 2 months ago | Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean. |
| [drone-scp](https://github.com/appleboy/drone-scp) | 45 | 2 | 2016-10-16 | 8 hours ago | Copy files and artifacts via SSH using a binary, docker or Drone CI. |
| [dropship](https://github.com/ChrisMcKenzie/dropship) | 43 | 3 | 2015-09-04 | 7 months ago | Tool for deploying code via cdn. |
| [dropship](https://github.com/ChrisMcKenzie/dropship) | 43 | 3 | 2015-09-04 | 7 months ago | Tool for deploying code via cdn. |
| [rodent](https://github.com/alouche/rodent) | 30 | 2 | 2014-06-02 | 1 year ago | Rodent helps you manage Go versions, projects and track dependencies. |
| [drone-jenkins](https://github.com/appleboy/drone-jenkins) | 20 | 1 | 2016-10-15 | 1 day ago | Trigger downstream Jenkins jobs using a binary, docker or Drone CI. |
| [awsenv](https://github.com/soniah/awsenv) | 18 | 1 | 2015-08-05 | 7 months ago | Small binary that loads Amazon (AWS) environment variables for a profile. |
| [depcharge](https://github.com/centerorbit/depcharge) | 8 | 3 | 2018-07-25 | 3 months ago | Helps orchestrating the execution of commands across the many dependencies in larger projects. |
| [lwc](https://github.com/timdp/lwc) | 7 | 2 | 2018-04-22 | 8 months ago | A live-updating version of the UNIX wc command. |
| [sg](https://github.com/ChristopherRabotin/sg) | 5 | 1 | 2015-08-19 | 2 years ago | Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response. |

### Other Software
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [goreplay](https://github.com/buger/goreplay) | 10491 | 441 | 2013-05-30 | 1 week ago | Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time. |
| [goreplay](https://github.com/buger/goreplay) | 10483 | 441 | 2013-05-30 | 1 week ago | Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time. |
| [rkt](https://github.com/rkt/rkt) | 8480 | 452 | 2014-11-12 | 2 days ago | App Container runtime that integrates with init systems, is compatible with other container formats like Docker, and supports alternative execution engines like KVM. |
| [rkt](https://github.com/rkt/rkt) | 8476 | 452 | 2014-11-12 | 2 days ago | App Container runtime that integrates with init systems, is compatible with other container formats like Docker, and supports alternative execution engines like KVM. |
| [seaweedfs](https://github.com/chrislusf/seaweedfs) | 7412 | 471 | 2014-07-15 | 1 day ago | Fast, Simple and Scalable Distributed File System with O(1) disk seek. |
| [restic](https://github.com/restic/restic) | 6191 | 216 | 2014-04-27 | 3 days ago | De-duplicating backup program. |
| [comcast](https://github.com/tylertreat/comcast) | 5977 | 150 | 2014-11-12 | 4 months ago | Simulate bad network connections. |
| [comcast](https://github.com/tylertreat/comcast) | 5974 | 150 | 2014-11-12 | 4 months ago | Simulate bad network connections. |
| [confd](https://github.com/kelseyhightower/confd) | 5913 | 238 | 2013-10-01 | 2 weeks ago | Manage local application configuration files using templates and data from etcd or consul. |
| [liteide](https://github.com/visualfc/liteide) | 5136 | 372 | 2012-11-19 | 2 days ago | LiteIDE is a simple, open source, cross-platform Go IDE. |
| [drive](https://github.com/odeke-em/drive) | 4717 | 212 | 2014-11-03 | 4 weeks ago | Google Drive client for the commandline. |
| [nes](https://github.com/fogleman/nes) | 3963 | 138 | 2015-03-03 | 1 month ago | Nintendo Entertainment System (NES) emulator written in Go. |
| [toxiproxy](https://github.com/Shopify/toxiproxy) | 3531 | 261 | 2014-09-04 | 1 week ago | Proxy to simulate network and system conditions for automated tests. |
| [toxiproxy](https://github.com/Shopify/toxiproxy) | 3529 | 261 | 2014-09-04 | 1 week ago | Proxy to simulate network and system conditions for automated tests. |
| [duplicacy](https://github.com/gilbertchen/duplicacy) | 2544 | 88 | 2016-02-23 | 1 day ago | A cross-platform network and cloud backup tool based on the idea of lock-free deduplication. |
| [pipe](https://github.com/b3log/pipe) | 2152 | 88 | 2017-09-11 | 9 hours ago | A small and beautiful blogging platform. |
| [mylg](https://github.com/mehrdadrad/mylg) | 2141 | 104 | 2016-06-22 | 4 months ago | Command Line Network Diagnostic tool written in Go. |
| [goboy](https://github.com/Humpheh/goboy) | 1996 | 37 | 2017-08-20 | 3 weeks ago | Nintendo Game Boy Color emulator written in Go. |
| [sup](https://github.com/pressly/sup) | 1883 | 75 | 2015-02-24 | 2 months ago | Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers. |
| [snap](https://github.com/intelsdi-x/snap) | 1791 | 154 | 2014-08-14 | 2 months ago | Powerful telemetry framework. |
| [circuit](https://github.com/gocircuit/circuit) | 1741 | 139 | 2014-04-11 | 2 years ago | Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications. |
| [lgo](https://github.com/yunabe/lgo) | 1616 | 43 | 2017-10-05 | 5 days ago | Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility. |
| [borg](https://github.com/ok-borg/borg) | 1403 | 43 | 2016-09-11 | 1 year ago | Terminal based search engine for bash snippets. |
| [borg](https://github.com/ok-borg/borg) | 1403 | 43 | 2016-09-11 | 1 year ago | Terminal based search engine for bash snippets. |
| [Go-Package-Store](https://github.com/shurcooL/Go-Package-Store) | 875 | 22 | 2014-01-24 | 2 months ago | App that displays updates for the Go packages in your GOPATH. |
| [scc](https://github.com/boyter/scc) | 653 | 15 | 2018-03-01 | 5 hours ago | Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates. |
| [leaps](https://github.com/Jeffail/leaps) | 631 | 26 | 2014-06-20 | 3 days ago | Pair programming service using Operational Transforms. |
| [leaps](https://github.com/Jeffail/leaps) | 631 | 26 | 2014-06-20 | 3 days ago | Pair programming service using Operational Transforms. |
| [community](https://github.com/documize/community) | 612 | 40 | 2016-04-30 | 23 hours ago | Modern wiki software that integrates data from SaaS tools. |
| [peg](https://github.com/pointlander/peg) | 553 | 30 | 2010-04-26 | 2 days ago | Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator. |
| [vflow](https://github.com/VerizonDigital/vflow) | 526 | 72 | 2017-02-25 | 3 weeks ago | High-performance, scalable and reliable IPFIX, sFlow and Netflow collector. |
| [mockingjay-server](https://github.com/quii/mockingjay-server) | 393 | 14 | 2015-04-05 | 3 months ago | Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests. |
| [godns](https://github.com/TimothyYe/godns) | 350 | 19 | 2014-05-11 | 1 week ago | A dynamic DNS client tool, supports DNSPod & HE.net, written in Go. |
| [godns](https://github.com/TimothyYe/godns) | 349 | 19 | 2014-05-11 | 1 week ago | A dynamic DNS client tool, supports DNSPod & HE.net, written in Go. |
| [shell2http](https://github.com/msoap/shell2http) | 328 | 18 | 2015-03-12 | 3 days ago | Executing shell commands via http server (for prototyping or remote control). |
| [gocc](https://github.com/goccmack/gocc) | 312 | 22 | 2015-06-05 | 4 hours ago | Gocc is a compiler kit for Go written in Go. |
| [wellington](https://github.com/wellington/wellington) | 286 | 13 | 2014-12-09 | 4 months ago | Sass project management tool, extends the language with sprite functions (like Compass). |
| [ipe](https://github.com/dimiro1/ipe) | 251 | 14 | 2015-01-13 | 2 months ago | Open source Pusher server implementation compatible with Pusher client libraries written in GO. |
| [IDE](https://github.com/thestrukture/IDE) | 243 | 14 | 2017-09-10 | 2 months ago | Browser accessible IDE. Designed for Go with Go. |
| [IDE](https://github.com/thestrukture/IDE) | 243 | 14 | 2017-09-10 | 2 months ago | Browser accessible IDE. Designed for Go with Go. |
| [cherry](https://github.com/rafael-santiago/cherry) | 185 | 12 | 2015-10-25 | 1 year ago | Tiny webchat server in Go. |
| [orange-cat](https://github.com/utatti/orange-cat) | 174 | 6 | 2014-11-01 | 1 month ago | Markdown previewer written in Go. |
| [orange-cat](https://github.com/utatti/orange-cat) | 174 | 6 | 2014-11-01 | 1 month ago | Markdown previewer written in Go. |
| [orbit](https://github.com/gulien/orbit) | 126 | 9 | 2017-05-13 | 8 months ago | A simple tool for running commands and generating files from templates. |
| [joincap](https://github.com/assafmo/joincap) | 116 | 7 | 2018-06-01 | 2 weeks ago | Command-line utility for merging multiple pcap files together. |
| [ddns](https://github.com/skibish/ddns) | 81 | 5 | 2017-03-14 | 3 months ago | Personal DDNS client with Digital Ocean Networking DNS as backend. |
| [boxed](https://github.com/tejo/boxed) | 72 | 2 | 2015-04-19 | 6 months ago | Dropbox based blog engine. |
| [naclpipe](https://github.com/unix4fun/naclpipe) | 20 | 5 | 2015-05-06 | 3 months ago | Simple NaCL EC25519 based crypto pipe tool written in Go. |
| [snitch](https://github.com/lucasgomide/snitch) | 15 | 1 | 2017-04-07 | 7 months ago | Simple way to notify your team and many tools when someone has deployed any application via Tsuru. |
| [term-quiz](https://github.com/crazcalm/term-quiz) | 15 | 1 | 2017-12-26 | 4 months ago | Quizzes for your terminal. |
| [GoDocTooltip](https://github.com/diankong/GoDocTooltip) | 12 | 3 | 2016-01-21 | 3 years ago | Chrome extension for Go Doc sites, which shows function description as tooltip at function list. |

# Server Applications
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [etcd](https://github.com/etcd-io/etcd) | 23303 | 1229 | 2013-07-07 | 27 minutes ago | Highly-available key value store for shared configuration and service discovery. |
| [etcd](https://github.com/etcd-io/etcd) | 23278 | 1228 | 2013-07-07 | 1 day ago | Highly-available key value store for shared configuration and service discovery. |
| [caddy](https://github.com/mholt/caddy) | 20806 | 668 | 2015-01-14 | 19 hours ago | Caddy is an alternative, HTTP/2 web server that's easy to configure and use. |
| [minio](https://github.com/minio/minio) | 14966 | 417 | 2015-01-15 | 1 hour ago | Minio is a distributed object storage server. |
| [devd](https://github.com/cortesi/devd) | 2732 | 68 | 2015-09-28 | 4 weeks ago | Local webserver for developers. |
| [roadrunner](https://github.com/spiral/roadrunner) | 2624 | 104 | 2017-12-27 | 1 week ago | High-performance PHP application server, load-balancer and process manager. |
| [fider](https://github.com/getfider/fider) | 670 | 14 | 2017-01-18 | 15 hours ago | Fider is an open platform to collect and organize customer feedback. |
| [jackal](https://github.com/ortuman/jackal) | 637 | 38 | 2017-11-14 | 3 days ago | An XMPP server written in Go. |
| [flagr](https://github.com/checkr/flagr) | 555 | 45 | 2017-10-04 | 5 days ago | Flagr is an open-source feature flagging and A/B testing service. |
| [algernon](https://github.com/xyproto/algernon) | 523 | 36 | 2015-03-10 | 1 week ago | HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber. |
| [discovery](https://github.com/bilibili/discovery) | 349 | 28 | 2018-04-20 | 1 week ago | A registry for resilient mid-tier load balancing and failover. |
| [discovery](https://github.com/bilibili/discovery) | 349 | 28 | 2018-04-20 | 1 week ago | A registry for resilient mid-tier load balancing and failover. |
| [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) | 4 | 1 | 2018-10-23 | 1 month ago | Nginx log parser and exporter to Prometheus. |

# Resources
        
*Where to discover new Go libraries.*

## Benchmarks
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) | 1203 | 54 | 2013-12-17 | 10 months ago | Go HTTP request router benchmark and comparison. |
| [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) | 897 | 75 | 2016-04-06 | 1 month ago | Go web framework benchmark. |
| [skynet](https://github.com/atemerev/skynet) | 888 | 46 | 2016-02-14 | 4 months ago | Skynet 1M threads microbenchmark. |
| [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) | 767 | 34 | 2013-01-19 | 1 month ago | Benchmarks of Go serialization methods. |
| [speedtest-resize](https://github.com/fawick/speedtest-resize) | 163 | 8 | 2013-09-16 | 2 years ago | Compare various Image resize algorithms for the Go language. |
| [go-benchmarks](https://github.com/tylertreat/go-benchmarks) | 112 | 9 | 2016-02-25 | 3 years ago | Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches. |
| [autobench](https://github.com/davecheney/autobench) | 89 | 8 | 2013-03-28 | 4 years ago | Framework to compare the performance between different Go versions. |
| [gospeed](https://github.com/feyeleanor/gospeed) | 89 | 7 | 2011-05-24 | 5 months ago | Go micro-benchmarks for calculating the speed of language constructs. |
| [gospeed](https://github.com/feyeleanor/gospeed) | 88 | 7 | 2011-05-24 | 5 months ago | Go micro-benchmarks for calculating the speed of language constructs. |
| [gocostmodel](https://github.com/mna/gocostmodel) | 51 | 5 | 2014-12-19 | 4 years ago | Benchmarks of common basic operations for the Go language. |
| [gocostmodel](https://github.com/mna/gocostmodel) | 51 | 5 | 2014-12-19 | 4 years ago | Benchmarks of common basic operations for the Go language. |
| [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) | 48 | 5 | 2014-09-25 | 11 months ago | Collection of benchmarks for popular Go database/SQL utilities. |
| [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) | 19 | 1 | 2017-01-24 | 1 year ago | Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results. |
| [golang-micro-benchmarks](https://github.com/amscanne/golang-micro-benchmarks) | 16 | 2 | 2014-04-05 | 4 days ago | Tiny collection of Go micro benchmarks. The intent is to compare some language features to others. |
| [kvbench](https://github.com/jimrobinson/kvbench) | 14 | 1 | 2014-04-15 | 4 years ago | Key/Value database benchmark. |

## Conferences
        

## E-Books
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [GoBooks](https://github.com/dariubs/GoBooks) | 6103 | 487 | 2015-05-05 | 1 week ago | A curated list of Go books. |
| [web-dev-golang-anti-textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook) | 2220 | 90 | 2016-01-01 | 1 week ago | Learn how to write webapps without a framework in Go. |

## Gophers
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gophers](https://github.com/ashleymcnamara/gophers) | 1615 | 82 | 2017-02-15 | 5 days ago | Gopher artworks by Ashley McNamara. |
| [gophers](https://github.com/egonelbre/gophers) | 1427 | 31 | 2015-06-03 | 1 month ago | Free gophers. |
| [gophericons](https://github.com/shalakhin/gophericons) | 554 | 20 | 2015-08-22 | 11 months ago | 34 gopher images for Go developers community |
| [gopher-stickers](https://github.com/tenntenn/gopher-stickers) | 404 | 14 | 2014-11-10 | 2 years ago | gopher stickers |
| [gopher-vector](https://github.com/golang-samples/gopher-vector) | 331 | 13 | 2013-03-31 | 2 years ago | Vector data of gopher |
| [gopherize.me](https://github.com/matryer/gopherize.me) | 270 | 6 | 2017-01-25 | 4 months ago | Gopherize yourself. |
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
| [go](https://github.com/golang/go) | 54056 | 3257 | 2014-08-19 | 1 hour ago | List of projects on the Go community wiki. |
| [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) | 23795 | 1711 | 2014-07-08 | 6 days ago | List of other amazingly awesome lists. |
| [awesome-remote-job](https://github.com/lukasz-madon/awesome-remote-job) | 13596 | 791 | 2015-01-02 | 1 week ago | Curated list of awesome remote jobs. A lot of them are looking for Go hackers. |
| [golang-graphics](https://github.com/mholt/golang-graphics) | 141 | 10 | 2014-03-25 | 3 years ago | Collection of Go images, graphics, and art. |
| [gocryforhelp](https://github.com/ninedraft/gocryforhelp) | 32 | 10 | 2016-05-09 | 1 year ago | Collection of Go projects that needs help. Good place to start your open-source way in Go. |

### Tutorials
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang) | 28282 | 2316 | 2012-08-02 | 21 hours ago | Golang ebook intro how to build a web app with golang. |
| [go-lang-cheat-sheet](https://github.com/a8m/go-lang-cheat-sheet) | 3763 | 156 | 2014-02-13 | 2 weeks ago | Go's reference card. |
| [learn-go-with-tests](https://github.com/quii/learn-go-with-tests) | 3093 | 124 | 2018-03-02 | 1 day ago | Learn Go with test-driven development. |
| [working-with-go](https://github.com/mkaz/working-with-go) | 1087 | 50 | 2014-05-05 | 2 weeks ago | Intro to go for experienced programmers. |
| [ethereum-development-with-go-book](https://github.com/miguelmota/ethereum-development-with-go-book) | 359 | 26 | 2018-05-16 | 5 days ago | A little e-book on Ethereum Development with Go. |
| [golang-for-nodejs-developers](https://github.com/miguelmota/golang-for-nodejs-developers) | 242 | 10 | 2019-01-03 | 3 days ago | Examples of Golang compared to Node.js for learning. |