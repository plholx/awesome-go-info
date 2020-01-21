# Awesome Go Info

go语言开源项目列表，项目分类及GitHub上的开源项目数据完全来自于[awesome-go](https://github.com/avelino/awesome-go) 的[README.md](https://github.com/avelino/awesome-go/blob/master/README.md)文件，通过调用GitHub的API获取仓库信息，展示项目的star数、watch数等，方便查看go语言开源项目的一些相关信息。

_该文件仅包含[awesome-go](https://github.com/avelino/awesome-go) [README.md](https://github.com/avelino/awesome-go/blob/master/README.md)文件中列出的在GitHub上开源的优秀项目，不罗列其它golang相关的网站_
_该文件中的GitHub仓库信息数据会在每天凌晨1点左右更新,当前数据更新于2020-01-22 01:05:51_

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
    - [Dynamic DNS](#dynamic-dns)
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
        - [Project Layout](#project-layout)
        - [Strings](#strings)
    - [Natural Language Processing](#natural-language-processing)
    - [Networking](#networking)
        - [HTTP Clients](#http-clients)
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
| [oto](https://github.com/hajimehoshi/oto) | 512 | 8 | 2017-05-04 | 3 weeks ago | A low-level library to play sound on multiple platforms. |
| [portaudio](https://github.com/gordonklaus/portaudio) | 331 | 14 | 2015-09-16 | 1 year ago | Go bindings for the PortAudio audio I/O library. |
| [music-theory](https://github.com/go-music-theory/music-theory) | 277 | 11 | 2016-03-17 | 1 month ago | Music theory models in Go. |
| [waveform](https://github.com/mdlayher/waveform) | 269 | 13 | 2014-09-13 | 3 years ago | Go package capable of generating waveform images from audio streams. |
| [portmidi](https://github.com/rakyll/portmidi) | 223 | 7 | 2013-11-10 | 2 months ago | Go bindings for PortMidi. |
| [id3v2](https://github.com/bogem/id3v2) | 133 | 5 | 2016-05-15 | 3 months ago | Fast and stable ID3 parsing and writing library for Go. |
| [flac](https://github.com/mewkiz/flac) | 112 | 9 | 2012-11-01 | 1 month ago | Native Go FLAC encoder/decoder with support for FLAC streams. |
| [mix](https://github.com/go-mix/mix) | 106 | 3 | 2016-01-03 | 1 month ago | Sequence-based Go-native audio mixer for music apps. |
| [mp3](https://github.com/tcolgate/mp3) | 105 | 1 | 2015-02-26 | 2 years ago | Native Go MP3 decoder. |
| [go-sox](https://github.com/krig/go-sox) | 102 | 8 | 2013-10-08 | 1 year ago | libsox bindings for go. |
| [malgo](https://github.com/gen2brain/malgo) | 92 | 5 | 2017-11-09 | 3 months ago | Mini audio library. |
| [go-taglib](https://github.com/wtolson/go-taglib) | 72 | 6 | 2012-11-20 | 1 year ago | Go bindings for taglib. |
| [gaad](https://github.com/Comcast/gaad) | 63 | 11 | 2016-07-11 | 1 year ago | Native Go AAC bitstream parser. |
| [minimp3](https://github.com/tosone/minimp3) | 35 | 2 | 2018-01-26 | 6 months ago | Lightweight MP3 decoder library. |
| [go_mediainfo](https://github.com/zhulik/go_mediainfo) | 28 | 1 | 2015-12-13 | 4 years ago | libmediainfo bindings for go. |
| [EasyMIDI](https://github.com/algoGuy/EasyMIDI) | 26 | 2 | 2018-02-19 | 1 year ago | EasyMidi is a simple and reliable library for working with standard midi file (SMF). |
| [vorbis](https://github.com/mccoyst/vorbis) | 24 | 3 | 2013-07-12 | 9 months ago | "Native" Go Vorbis decoder (uses CGO, but has no dependencies). |
| [gosamplerate](https://github.com/dh1tw/gosamplerate) | 9 | 1 | 2016-11-20 | 2 months ago | libsamplerate bindings for go. |

## Authentication and OAuth
        
*Libraries for implementing authentications schemes.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [jwt-go](https://github.com/dgrijalva/jwt-go) | 6908 | 140 | 2012-04-18 | 7 hours ago | Golang implementation of JSON Web Tokens (JWT). |
| [casbin](https://github.com/casbin/casbin) | 5962 | 180 | 2017-04-08 | 2 weeks ago | Authorization library that supports access control models like ACL, RBAC, ABAC. |
| [oauth2](https://github.com/golang/oauth2) | 2688 | 104 | 2014-04-14 | 1 week ago | Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support. |
| [goth](https://github.com/markbates/goth) | 2507 | 67 | 2014-10-14 | 3 days ago | provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box. |
| [authboss](https://github.com/volatiletech/authboss) | 2097 | 42 | 2015-01-03 | 1 month ago | Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time. |
| [osin](https://github.com/openshift/osin) | 1576 | 68 | 2013-09-10 | 1 month ago | Golang OAuth2 server library. |
| [go-jose](https://github.com/square/go-jose) | 1446 | 65 | 2014-11-14 | 1 week ago | Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs. |
| [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) | 1420 | 76 | 2015-11-01 | 3 months ago | Standalone, specification-compliant,  OAuth2 server written in Golang. |
| [gologin](https://github.com/dghubble/gologin) | 1142 | 28 | 2015-06-23 | 5 months ago | chainable handlers for login with OAuth1 and OAuth2 authentication providers. |
| [gorbac](https://github.com/mikespook/gorbac) | 975 | 59 | 2013-12-26 | 10 months ago | provides a lightweight role-based access control (RBAC) implementation in Golang. |
| [loginsrv](https://github.com/tarent/loginsrv) | 887 | 49 | 2016-11-11 | 1 week ago | JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam. |
| [scs](https://github.com/alexedwards/scs) | 657 | 18 | 2016-08-08 | 1 week ago | Session Manager for HTTP servers. |
| [permissions2](https://github.com/xyproto/permissions2) | 382 | 13 | 2014-11-19 | 1 month ago | Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt. |
| [paseto](https://github.com/o1egl/paseto) | 291 | 15 | 2018-01-23 | 3 days ago | Golang implementation of Platform-Agnostic Security Tokens (PASETO). |
| [httpauth](https://github.com/goji/httpauth) | 190 | 7 | 2014-05-26 | 3 years ago | HTTP Authentication middleware. |
| [jeff](https://github.com/abraithwaite/jeff) | 183 | 4 | 2018-08-02 | 3 months ago | Simple, flexible, secure and idiomatic web session management with pluggable backends. |
| [jwt-auth](https://github.com/adam-hanna/jwt-auth) | 165 | 11 | 2016-07-05 | 11 months ago | JWT middleware for Golang http servers with many configuration options. |
| [jwt](https://github.com/pascaldekloe/jwt) | 142 | 9 | 2018-03-21 | 2 weeks ago | Lightweight JSON Web Token (JWT) library. |
| [branca](https://github.com/hako/branca) | 100 | 7 | 2018-01-09 | 3 weeks ago | Golang implementation of Branca Tokens. |
| [session](https://github.com/icza/session) | 95 | 6 | 2016-02-08 | 6 months ago | Go session management for web servers (including support for Google App Engine - GAE). |
| [jwt](https://github.com/robbert229/jwt) | 78 | 6 | 2016-06-05 | 1 year ago | Clean and easy to use implementation of JSON Web Tokens (JWT). |
| [sessionup](https://github.com/swithek/sessionup) | 76 | 5 | 2019-07-23 | 3 months ago | Simple, yet effective HTTP session management and identification package. |
| [sjwt](https://github.com/brianvoe/sjwt) | 52 | 0 | 2019-06-20 | 4 months ago | Simple jwt generator and parser. |
| [sessions](https://github.com/adam-hanna/sessions) | 49 | 3 | 2017-04-29 | 8 months ago | Dead simple, highly performant, highly customizable sessions service for go http servers. |
| [rbac](https://github.com/zpatrick/rbac) | 44 | 3 | 2018-08-02 | 1 year ago | Minimalistic RBAC package for Go applications. |
| [securecookie](https://github.com/chmike/securecookie) | 34 | 5 | 2017-09-03 | 1 year ago | Efficient secure cookie encoding/decoding. |
| [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) | 8 | 2 | 2017-10-20 | 1 year ago | Go session management using the SessionGate Redis module. |
| [signedvalue](https://github.com/sashka/signedvalue) | 8 | 0 | 2018-01-06 | 4 months ago | Signed and timestamped strings compatible with [Tornado's](https://github.com/tornadoweb/tornado) `create_signed_value`, `decode_signed_value`, and therefore `set_secure_cookie` and `get_secure_cookie`. |
| [scope](https://github.com/SonicRoshan/scope) | 3 | 1 | 2019-09-23 | 1 week ago | Easily Manage OAuth2 Scopes In Go. |
| [cookiestxt](https://github.com/mengzhuo/cookiestxt) | 2 | 1 | 2017-10-09 | 2 years ago | provides parser of cookies.txt file format. |

## Bot Building
        
*Libraries for building and working with bots.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) | 1876 | 63 | 2015-06-25 | 1 day ago | Simple and clean Telegram bot client. |
| [telebot](https://github.com/tucnak/telebot) | 1084 | 42 | 2015-06-25 | 1 day ago | Telegram bot framework written in Go. |
| [bot](https://github.com/go-chat-bot/bot) | 536 | 38 | 2015-09-22 | 1 month ago | IRC, Slack & Telegram bot written in Go. |
| [slacker](https://github.com/shomali11/slacker) | 345 | 13 | 2017-05-20 | 1 week ago | Easy to use framework to create Slack bots. |
| [golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot) | 276 | 25 | 2017-05-14 | 11 months ago | A golang implementation of a console-based trading bot for cryptocurrency exchanges. |
| [kelp](https://github.com/stellar/kelp) | 251 | 32 | 2018-08-08 | 2 days ago | official trading and market-making bot for the [Stellar](https://www.stellar.org/) DEX. Works out-of-the-box, written in Golang, compatible with centralized exchanges and custom trading strategies. |
| [tbot](https://github.com/yanzay/tbot) | 244 | 11 | 2015-09-11 | 1 month ago | Telegram bot server with API similar to net/http. |
| [tenyks](https://github.com/kyleterry/tenyks) | 169 | 14 | 2012-08-26 | 4 months ago | Service oriented IRC bot using Redis and JSON for messaging. |
| [go-sarah](https://github.com/oklahomer/go-sarah) | 153 | 7 | 2016-11-06 | 1 month ago | Framework to build bot for desired chat services including LINE, Slack, Gitter and more. |
| [hanu](https://github.com/sbstjn/hanu) | 116 | 5 | 2016-09-16 | 2 months ago | Framework for writing Slack bots. |
| [go-tgbot](https://github.com/olebedev/go-tgbot) | 91 | 8 | 2016-12-11 | 1 year ago | Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware. |
| [margelet](https://github.com/zhulik/margelet) | 60 | 5 | 2015-11-21 | 3 years ago | Framework for building Telegram bots. |
| [govkbot](https://github.com/nikepan/govkbot) | 30 | 3 | 2016-07-11 | 4 months ago | Simple Go [VK](https://vk.com) bot library. |
| [slackscot](https://github.com/alexandre-normand/slackscot) | 21 | 1 | 2015-10-22 | 1 day ago | Another framework for building Slack bots. |
| [micha](https://github.com/onrik/micha) | 12 | 3 | 2016-04-14 | 3 months ago | Go Library for Telegram bot api. |

## Command Line
        
*Libraries for building Console Applications and Console User Interfaces.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [cobra](https://github.com/spf13/cobra) | 15510 | 315 | 2013-09-03 | 5 days ago | Commander for modern Go CLI interactions. |
| [cli](https://github.com/urfave/cli) | 12947 | 284 | 2013-07-13 | 19 hours ago | Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli). |
| [termui](https://github.com/gizak/termui) | 9441 | 285 | 2015-02-03 | 4 months ago | Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib). |
| [gocui](https://github.com/jroimartin/gocui) | 5897 | 123 | 2014-01-04 | 1 month ago | Minimalist Go library aimed at creating Console User Interfaces. |
| [termbox-go](https://github.com/nsf/termbox-go) | 3644 | 100 | 2012-01-12 | 3 weeks ago | Termbox is a library for creating cross-platform text-based interfaces. |
| [kingpin](https://github.com/alecthomas/kingpin) | 2787 | 55 | 2014-05-14 | 16 hours ago | Command line and flag parser supporting sub commands. |
| [go-prompt](https://github.com/c-bata/go-prompt) | 2764 | 45 | 2017-08-14 | 3 weeks ago | Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit). |
| [uiprogress](https://github.com/gosuri/uiprogress) | 1616 | 33 | 2015-11-17 | 6 months ago | Flexible library to render progress bars in terminal applications. |
| [go-flags](https://github.com/jessevdk/go-flags) | 1611 | 28 | 2012-08-31 | 1 month ago | go command line option parser. |
| [readline](https://github.com/chzyer/readline) | 1443 | 39 | 2015-09-20 | 2 months ago | Pure golang implementation that provides most features in GNU-Readline under MIT license. |
| [dnote](https://github.com/dnote/dnote) | 1439 | 26 | 2017-03-30 | 1 day ago | A simple and end-to-end encrypted notebook for developers. |
| [asciigraph](https://github.com/guptarohit/asciigraph) | 1273 | 27 | 2018-06-17 | 3 months ago | Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. |
| [docopt.go](https://github.com/docopt/docopt.go) | 1209 | 33 | 2013-08-25 | 4 months ago | Command-line arguments parser that will make you smile. |
| [uilive](https://github.com/gosuri/uilive) | 1074 | 17 | 2015-11-16 | 2 weeks ago | Library for updating terminal output in realtime. |
| [cli](https://github.com/mitchellh/cli) | 1069 | 23 | 2013-11-03 | 2 months ago | Go library for implementing command-line interfaces. |
| [pflag](https://github.com/spf13/pflag) | 965 | 29 | 2013-08-30 | 2 days ago | Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. |
| [termdash](https://github.com/mum4k/termdash) | 911 | 21 | 2018-03-24 | 6 months ago | Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui). |
| [gcli](https://github.com/tcnksm/gcli) | 889 | 26 | 2014-06-19 | 2 years ago | The easy way to start building Golang command line applications. |
| [progressbar](https://github.com/schollz/progressbar) | 836 | 20 | 2017-10-26 | 3 weeks ago | Basic thread-safe progress bar that works in every OS. |
| [go-arg](https://github.com/alexflint/go-arg) | 823 | 15 | 2015-11-01 | 1 day ago | Struct-based argument parsing in Go. |
| [mpb](https://github.com/vbauerster/mpb) | 823 | 16 | 2016-12-14 | 1 month ago | Multi progress bar for terminal applications. |
| [aurora](https://github.com/logrusorgru/aurora) | 764 | 6 | 2016-11-06 | 2 weeks ago | ANSI terminal colors that supports fmt.Printf/Sprintf. |
| [complete](https://github.com/posener/complete) | 663 | 14 | 2017-05-05 | 2 weeks ago | Write bash completions in Go + Go command bash completion. |
| [mow.cli](https://github.com/jawher/mow.cli) | 651 | 17 | 2014-12-18 | 1 month ago | Go library for building CLI applications with sophisticated flag and argument parsing and validation. |
| [liner](https://github.com/peterh/liner) | 633 | 23 | 2012-08-15 | 3 weeks ago | Go readline-like library for command-line interfaces. |
| [uitable](https://github.com/gosuri/uitable) | 539 | 15 | 2015-11-13 | 2 months ago | Library to improve readability in terminal apps using tabular data. |
| [flaggy](https://github.com/integrii/flaggy) | 537 | 17 | 2018-03-05 | 2 days ago | A robust and idiomatic flags package with excellent subcommand support. |
| [cli](https://github.com/mkideal/cli) | 509 | 22 | 2016-02-26 | 4 months ago | Feature-rich and easy to use command-line package based on golang struct tags. |
| [go-colorable](https://github.com/mattn/go-colorable) | 406 | 17 | 2014-07-30 | 1 week ago | Colorable writer for windows. |
| [go-isatty](https://github.com/mattn/go-isatty) | 389 | 8 | 2014-04-01 | 7 hours ago | isatty for golang. |
| [ops](https://github.com/nanovms/ops) | 353 | 21 | 2018-09-10 | 4 days ago | Unikernel Builder/Orchestrator. |
| [chalk](https://github.com/ttacon/chalk) | 326 | 8 | 2014-07-18 | 4 months ago | Intuitive package for prettifying terminal/console output. |
| [color](https://github.com/gookit/color) | 302 | 9 | 2018-07-01 | 5 days ago | Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows. |
| [tabby](https://github.com/cheynewallace/tabby) | 259 | 2 | 2018-12-17 | 10 months ago | A tiny library for super simple Golang tables. |
| [go-colortext](https://github.com/daviddengcn/go-colortext) | 200 | 9 | 2013-01-23 | 1 year ago | Go library for color output in terminals. |
| [simpletable](https://github.com/alexeyco/simpletable) | 184 | 2 | 2017-03-29 | 3 months ago | Simple tables in terminal with Go. |
| [argparse](https://github.com/akamensky/argparse) | 151 | 7 | 2017-11-24 | 5 hours ago | Command line argument parser inspired by Python's argparse module. |
| [commandeer](https://github.com/jaffee/commandeer) | 105 | 7 | 2017-10-12 | 2 months ago | Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags. |
| [flag](https://github.com/zhuah/flag) | 103 | 5 | 2016-10-05 | 10 months ago | Simple but powerful command line option parsing library for Go supporting subcommand. |
| [sflags](https://github.com/octago/sflags) | 103 | 5 | 2016-12-04 | 5 months ago | Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries. |
| [clif](https://github.com/ukautz/clif) | 98 | 2 | 2015-05-30 | 11 months ago | Small command line interface framework. |
| [wmenu](https://github.com/dixonwille/wmenu) | 95 | 2 | 2016-04-20 | 2 weeks ago | Easy to use menu structure for cli applications that prompts users to make choices. |
| [cfmt](https://github.com/mingrammer/cfmt) | 70 | 3 | 2018-03-15 | 1 year ago | Contextual fmt inspired by bootstrap color classes. |
| [job](https://github.com/liujianping/job) | 67 | 2 | 2019-04-09 | 8 months ago | JOB, make your short-term command as a long-term job. |
| [cli](https://github.com/teris-io/cli) | 61 | 1 | 2017-05-24 | 8 months ago | Simple and complete API for building command line interfaces in Go. |
| [env](https://github.com/codingconcepts/env) | 47 | 1 | 2017-06-14 | 7 months ago | Tag-based environment configuration for structs. |
| [1build](https://github.com/gopinath-langote/1build) | 45 | 4 | 2019-04-23 | 2 months ago | Command line tool to frictionlessly manage project-specific commands. |
| [wlog](https://github.com/dixonwille/wlog) | 40 | 1 | 2016-04-13 | 2 weeks ago | Simple logging interface that supports cross-platform color and concurrency. |
| [gocmd](https://github.com/devfacet/gocmd) | 37 | 3 | 2018-01-08 | 1 year ago | Go library for building command line applications. |
| [tabular](https://github.com/InVisionApp/tabular) | 36 | 4 | 2018-04-23 | 1 year ago | Print ASCII tables from command line utilities without the need to pass large sets of data to the API. |
| [flagvar](https://github.com/sgreben/flagvar) | 33 | 2 | 2018-05-18 | 2 weeks ago | A collection of flag argument types for Go's standard `flag` package. |
| [strumt](https://github.com/antham/strumt) | 32 | 0 | 2017-06-19 | 5 months ago | Library to create prompt chain. |
| [cmdr](https://github.com/hedzr/cmdr) | 28 | 4 | 2019-05-15 | 2 days ago | A POSIX/GNU style, getopt-like command-line UI Go library. |
| [argv](https://github.com/zhuah/argv) | 19 | 1 | 2017-01-22 | 7 months ago | Go library to split command line string as arguments array using the bash syntax. |
| [go-getoptions](https://github.com/DavidGamba/go-getoptions) | 19 | 3 | 2015-12-18 | 2 months ago | Go option parser inspired on the flexibility of Perl’s GetOpt::Long. |
| [colourize](https://github.com/TreyBastian/colourize) | 17 | 3 | 2015-05-11 | 3 years ago | Go library for ANSI colour text in terminals. |
| [go-commander](https://github.com/yitsushi/go-commander) | 15 | 1 | 2016-10-10 | 4 months ago | Go library to simplify CLI workflow. |
| [ctc](https://github.com/wzshiming/ctc) | 13 | 1 | 2018-04-27 | 1 week ago | The non-invasive cross-platform terminal color library does not need to modify the Print method. |
| [sand](https://github.com/Zaba505/sand) | 9 | 1 | 2018-11-18 | 1 year ago | Simple API for creating interpreters and so much more. |
| [go-ataman](https://github.com/workanator/go-ataman) | 8 | 2 | 2017-05-17 | 2 years ago | Go library for rendering ANSI colored text templates in terminals. |
| [ts](https://github.com/liujianping/ts) | 7 | 0 | 2019-06-25 | 6 months ago | Timestamp convert & compare tool. |
| [clir](https://github.com/leaanthony/clir) | 4 | 2 | 2019-11-18 | 3 weeks ago | A Simple and Clear CLI library. Dependency free. |
| [cmd](https://github.com/posener/cmd) | 3 | 1 | 2019-10-29 | 1 month ago | Extends the standard `flag` package to support sub commands and more in idomatic way. |

## Configuration
        
*Libraries for configuration parsing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [viper](https://github.com/spf13/viper) | 10960 | 208 | 2014-04-02 | 4 days ago | Go configuration with fangs. |
| [envconfig](https://github.com/kelseyhightower/envconfig) | 2731 | 41 | 2013-11-06 | 1 month ago | Go library for managing configuration data from environment variables. |
| [godotenv](https://github.com/joho/godotenv) | 2532 | 34 | 2013-07-30 | 3 months ago | Go port of Ruby's dotenv library (Loads environment variables from `.env`). |
| [ini](https://github.com/go-ini/ini) | 1865 | 68 | 2014-12-18 | 3 weeks ago | Go package to read and write INI files. |
| [env](https://github.com/caarlos0/env) | 1397 | 18 | 2015-07-28 | 3 months ago | Parse environment variables to Go structs (with defaults). |
| [konfig](https://github.com/lalamove/konfig) | 547 | 12 | 2019-01-18 | 2 months ago | Composable, observable and performant config handling for Go for the distributed processing era. |
| [confita](https://github.com/heetch/confita) | 287 | 23 | 2017-12-21 | 1 week ago | Load configuration in cascade from multiple backends into a struct. |
| [store](https://github.com/tucnak/store) | 248 | 5 | 2015-10-03 | 2 years ago | Lightweight configuration manager for Go. |
| [config](https://github.com/JeremyLoy/config) | 223 | 1 | 2019-04-02 | 1 month ago | Cloud native application configuration. Bind ENV to structs in only two lines. |
| [config](https://github.com/olebedev/config) | 223 | 8 | 2014-04-21 | 7 months ago | JSON or YAML configuration wrapper with environment variables and flags parsing. |
| [config](https://github.com/joshbetz/config) | 198 | 2 | 2017-04-02 | 2 months ago | Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP. |
| [hjson-go](https://github.com/hjson/hjson-go) | 186 | 9 | 2016-08-05 | 2 months ago | Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments. |
| [envconfig](https://github.com/vrischmann/envconfig) | 171 | 4 | 2015-04-21 | 6 months ago | Read your configuration from environment variables. |
| [koanf](https://github.com/knadh/koanf) | 148 | 5 | 2019-06-18 | 2 days ago | Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line. |
| [config](https://github.com/gookit/config) | 129 | 5 | 2018-07-07 | 1 month ago | application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge. |
| [gcfg](https://github.com/go-gcfg/gcfg) | 125 | 7 | 2015-08-17 | 2 months ago | read INI-style configuration files into Go structs; supports user-defined types and subsections. |
| [goconfig](https://github.com/crgimenes/goconfig) | 121 | 11 | 2016-12-18 | 3 months ago | Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. |
| [envh](https://github.com/antham/envh) | 92 | 3 | 2017-01-12 | 5 months ago | Helpers to manage environment variables. |
| [envcfg](https://github.com/tomazk/envcfg) | 90 | 1 | 2014-11-29 | 2 years ago | Un-marshaling environment variables to Go structs. |
| [gofigure](https://github.com/ian-kent/gofigure) | 58 | 6 | 2014-11-25 | 4 months ago | Go application configuration made easy. |
| [configure](https://github.com/paked/configure) | 53 | 3 | 2015-06-14 | 11 months ago | Provides configuration through multiple sources, including JSON, flags and environment variables. |
| [harvester](https://github.com/beatlabs/harvester) | 49 | 12 | 2019-04-09 | 5 hours ago | Harvester, a easy to use static and dynamic configuration package supportig seeding, env vars and Consul integration. |
| [xdg](https://github.com/OpenPeeDeeP/xdg) | 44 | 3 | 2017-07-20 | 1 month ago | Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html). |
| [config](https://github.com/golobby/config) | 39 | 2 | 2019-10-15 | 2 months ago | A lightweight yet powerful config package for Go projects. |
| [cleanenv](https://github.com/ilyakaznacheev/cleanenv) | 33 | 0 | 2019-07-12 | 1 week ago | Minimalistic configuration reader (from files, ENV, and wherever you want). |
| [ingo](https://github.com/schachmat/ingo) | 27 | 1 | 2016-02-07 | 2 years ago | Flags persisted in an ini-like config file. |
| [go-up](https://github.com/ufoscout/go-up) | 26 | 1 | 2018-02-18 | 1 week ago | A simple configuration library with recursive placeholders resolution and no magic. |
| [mini](https://github.com/sasbury/mini) | 20 | 1 | 2015-04-29 | 1 year ago | Golang package for parsing ini-style configuration files. |
| [conflate](https://github.com/the4thamigo-uk/conflate) | 11 | 0 | 2018-02-01 | 9 months ago | Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema. |
| [envconf](https://github.com/ian-kent/envconf) | 8 | 1 | 2014-10-26 | 5 years ago | Configuration from environment. |
| [genv](https://github.com/sakirsensoy/genv) | 8 | 1 | 2019-07-15 | 5 months ago | Read environment variables easily with dotenv support. |
| [sprbox](https://github.com/oblq/sprbox) | 5 | 2 | 2018-07-17 | 5 months ago | Build-environment aware toolbox factory and agnostic config parser (YAML, TOML, JSON and Environment vars). |
| [go-ssm-config](https://github.com/ianlopshire/go-ssm-config) | 3 | 1 | 2019-12-02 | 2 weeks ago | Go utility for loading configuration parameters from AWS SSM (Parameter Store). |
| [env](https://github.com/nasermirzaei89/env) | 2 | 0 | 2019-07-24 | 3 months ago | Simple useful package for read environment variables. |

## Continuous Integration
        
*Tools for help with continuous integration.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [drone](https://github.com/drone/drone) | 20335 | 580 | 2014-02-07 | 21 hours ago | Drone is a Continuous Integration platform built on Docker, written in Go. |
| [cds](https://github.com/ovh/cds) | 2614 | 68 | 2016-10-11 | 6 hours ago | Enterprise-Grade CI/CD and DevOps Automation Open Source Platform. |
| [goveralls](https://github.com/mattn/goveralls) | 623 | 15 | 2013-04-17 | 5 days ago | Go integration for Coveralls.io continuous code coverage tracking system. |
| [overalls](https://github.com/go-playground/overalls) | 99 | 3 | 2015-07-30 | 3 weeks ago | Multi-Package go project coverprofile for tools like goveralls. |
| [duci](https://github.com/duck8823/duci) | 51 | 3 | 2018-04-01 | 6 days ago | A simple ci server no needs domain specific languages. |
| [gomason](https://github.com/nikogura/gomason) | 42 | 1 | 2017-11-18 | 1 week ago | Test, Build, Sign, and Publish your go binaries from a clean workspace. |
| [roveralls](https://github.com/lawrencewoodman/roveralls) | 12 | 1 | 2016-06-26 | 2 years ago | Recursive coverage testing tool. |

## CSS Preprocessors
        
*Libraries for preprocessing CSS files.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gcss](https://github.com/yosssi/gcss) | 427 | 16 | 2014-09-04 | 5 years ago | Pure Go CSS Preprocessor. |
| [go-libsass](https://github.com/wellington/go-libsass) | 147 | 7 | 2015-04-19 | 11 months ago | Go wrapper to the 100% Sass compatible libsass project. |

## Data Structures
        
*Generic datastructures and algorithms in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gods](https://github.com/emirpasic/gods) | 7661 | 329 | 2015-03-04 | 1 week ago | Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc. |
| [go-datastructures](https://github.com/Workiva/go-datastructures) | 5408 | 321 | 2014-10-29 | 2 months ago | Collection of useful, performant, and thread-safe data structures. |
| [golang-set](https://github.com/deckarep/golang-set) | 1382 | 38 | 2013-07-03 | 1 month ago | Thread-Safe and Non-Thread-Safe high-performance sets for Go. |
| [BoomFilters](https://github.com/tylertreat/BoomFilters) | 1228 | 41 | 2015-02-06 | 1 year ago | Probabilistic data structures for processing continuous, unbounded streams. |
| [gota](https://github.com/go-gota/gota) | 1058 | 59 | 2016-02-06 | 11 hours ago | Implementation of dataframes, series, and data wrangling methods for Go. |
| [roaring](https://github.com/RoaringBitmap/roaring) | 773 | 36 | 2014-07-10 | 2 months ago | Go package implementing compressed bitsets. |
| [bloom](https://github.com/willf/bloom) | 738 | 29 | 2011-05-21 | 3 months ago | Go package implementing Bloom filters. |
| [hyperloglog](https://github.com/axiomhq/hyperloglog) | 682 | 16 | 2017-06-18 | 2 months ago | HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction. |
| [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) | 574 | 17 | 2015-06-28 | 2 weeks ago | Cuckoo filter: a good alternative to a counting bloom filter implemented in Go. |
| [bitset](https://github.com/willf/bitset) | 531 | 29 | 2011-05-11 | 3 months ago | Go package implementing bitsets. |
| [trie](https://github.com/derekparker/trie) | 450 | 18 | 2014-03-06 | 5 months ago | Trie implementation in Go. |
| [gocache](https://github.com/eko/gocache) | 342 | 10 | 2019-10-05 | 5 days ago | A complete Go cache library with mutiple stores (memory, memcache, redis, ...), chainable, loadable, metrics cache and more. |
| [algorithms](https://github.com/shady831213/algorithms) | 340 | 12 | 2018-01-31 | 9 months ago | Algorithms and data structures.CLRS study. |
| [go-geoindex](https://github.com/hailocab/go-geoindex) | 318 | 67 | 2015-01-22 | 1 year ago | In-memory geo index. |
| [mafsa](https://github.com/smartystreets-archives/mafsa) | 281 | 16 | 2014-11-07 | 7 months ago | MA-FSA implementation with Minimal Perfect Hashing. |
| [goskiplist](https://github.com/ryszard/goskiplist) | 215 | 15 | 2012-05-09 | 2 months ago | Skip list implementation in Go. |
| [hilbert](https://github.com/google/hilbert) | 193 | 21 | 2015-08-06 | 1 year ago | Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves. |
| [merkletree](https://github.com/cbergoon/merkletree) | 177 | 4 | 2017-04-12 | 2 months ago | Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures. |
| [binpacker](https://github.com/zhuangsirui/binpacker) | 133 | 11 | 2016-02-02 | 1 year ago | Binary packer and unpacker helps user build custom binary stream. |
| [bloom](https://github.com/zentures/bloom) | 131 | 7 | 2013-09-03 | 1 year ago | Bloom filters implemented in Go. |
| [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) | 126 | 7 | 2016-04-01 | 4 months ago | Go implementation of Adaptive Radix Tree. |
| [ttlcache](https://github.com/ReneKroon/ttlcache) | 122 | 10 | 2014-12-13 | 1 week ago | In-memory LRU string-interface{} map with expiration for golang. |
| [skiplist](https://github.com/MauriceGit/skiplist) | 111 | 5 | 2018-06-23 | 2 months ago | Very fast Go Skiplist implementation. |
| [iter](https://github.com/disksing/iter) | 105 | 5 | 2019-10-20 | 1 month ago | Go implementation of C++ STL iterators and algorithms. |
| [go-rquad](https://github.com/arl/go-rquad) | 101 | 3 | 2016-09-12 | 1 year ago | Region quadtrees with efficient point location and neighbour finding. |
| [ring](https://github.com/TheTannerRyan/ring) | 99 | 1 | 2019-01-27 | 3 months ago | Go implementation of a high performance, thread safe bloom filter. |
| [deque](https://github.com/gammazero/deque) | 98 | 7 | 2018-04-24 | 3 weeks ago | Fast ring-buffer deque (double-ended queue). |
| [encoding](https://github.com/zentures/encoding) | 97 | 6 | 2013-09-20 | 2 years ago | Integer Compression Libraries for Go. |
| [conjungo](https://github.com/InVisionApp/conjungo) | 85 | 18 | 2016-12-29 | 8 months ago | A small, powerful and flexible merge library. |
| [bit](https://github.com/yourbasic/bit) | 77 | 5 | 2017-05-03 | 1 year ago | Golang set data structure with bonus bit-twiddling functions. |
| [skiplist](https://github.com/gansidui/skiplist) | 67 | 7 | 2014-11-18 | 5 years ago | Skiplist implementation in Go. |
| [levenshtein](https://github.com/agnivade/levenshtein) | 66 | 3 | 2014-07-30 | 2 months ago | Implementation to calculate levenshtein distance in Go. |
| [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) | 58 | 1 | 2019-01-10 | 1 month ago | Concurrent FIFO queue. |
| [bloom](https://github.com/yourbasic/bloom) | 47 | 2 | 2017-05-06 | 2 years ago | Golang Bloom filter implementation. |
| [count-min-log](https://github.com/seiflotfy/count-min-log) | 46 | 2 | 2015-08-16 | 2 years ago | Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory). |
| [go-mcache](https://github.com/OrlovEvgeny/go-mcache) | 43 | 3 | 2018-04-14 | 4 hours ago | Fast in-memory key:value store/cache library. Pointer caches. |
| [levenshtein](https://github.com/agext/levenshtein) | 37 | 1 | 2016-04-08 | 6 months ago | Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. |
| [remember-go](https://github.com/rocketlaunchr/remember-go) | 33 | 3 | 2019-04-04 | 3 months ago | A universal interface for caching slow database queries (backed by redis, memcached, ristretto, or in-memory). |
| [concurrent-writer](https://github.com/free/concurrent-writer) | 26 | 4 | 2017-09-18 | 2 years ago | Highly concurrent drop-in replacement for `bufio.Writer`. |
| [crunch](https://github.com/superwhiskers/crunch) | 23 | 4 | 2019-02-27 | 3 months ago | Go package implementing buffers for handling various datatypes easily. |
| [goset](https://github.com/zoumo/goset) | 19 | 1 | 2017-08-25 | 3 months ago | A useful Set collection implementation for Go. |
| [pipeline](https://github.com/hyfather/pipeline) | 19 | 1 | 2018-04-25 | 1 year ago | An implementation of pipelines with fan-in and fan-out. |
| [typ](https://github.com/gurukami/typ) | 15 | 1 | 2019-03-03 | 5 months ago | Null Types, Safe primitive type conversion and fetching value from complex structures. |
| [deque](https://github.com/edwingeng/deque) | 14 | 1 | 2019-02-01 | 1 month ago | A highly optimized double-ended queue. |
| [dict](https://github.com/srfrog/dict) | 12 | 0 | 2019-04-23 | 7 months ago | Python-like dictionaries (dict) for Go. |
| [go-ef](https://github.com/amallia/go-ef) | 11 | 2 | 2017-09-22 | 2 years ago | A Go implementation of the Elias-Fano encoding. |
| [hide](https://github.com/emvi/hide) | 11 | 3 | 2019-01-16 | 10 months ago | ID type with marshalling to/from hash to prevent sending IDs to clients. |
| [mspm](https://github.com/BlackRabbitt/mspm) | 9 | 3 | 2018-05-17 | 1 year ago | Multi-String Pattern Matching Algorithm for information retrieval. |
| [null](https://github.com/emvi/null) | 8 | 2 | 2018-07-04 | 4 months ago | Nullable Go types that can be marshalled/unmarshalled to/from JSON. |
| [set](https://github.com/StudioSol/set) | 7 | 17 | 2018-07-20 | 6 months ago | Simple set data structure implementation in Go using LinkedHashMap. |
| [treap](https://github.com/perdata/treap) | 7 | 2 | 2018-09-16 | 1 month ago | Persistent, fast ordered map using tree heaps. |
| [gofal](https://github.com/xxjwxc/gofal) | 6 | 1 | 2019-08-05 | 3 months ago | fractional api for Go. |
| [timedmap](https://github.com/zekroTJA/timedmap) | 6 | 1 | 2019-01-30 | 2 months ago | Map with expiring key-value pairs. |
| [ptrie](https://github.com/viant/ptrie) | 5 | 7 | 2019-05-20 | 2 months ago | An implementation of prefix tree. |
| [parsefields](https://github.com/MonaxGT/parsefields) | 3 | 1 | 2019-04-12 | 8 months ago | Tools for parse JSON-like logs for collecting unique fields and events. |

## Database
        
*SQL query builder, libraries for building and using SQL.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [prometheus](https://github.com/prometheus/prometheus) | 28807 | 1112 | 2012-11-24 | 4 hours ago | Monitoring system and time series database. |
| [tidb](https://github.com/pingcap/tidb) | 22332 | 1286 | 2015-09-06 | 1 hour ago | TiDB is a distributed SQL database. Inspired by the design of Google F1. |
| [influxdb](https://github.com/influxdata/influxdb) | 18196 | 761 | 2013-09-26 | 2 hours ago | Scalable datastore for metrics, events, and real-time analytics. |
| [cockroach](https://github.com/cockroachdb/cockroach) | 17661 | 733 | 2014-02-06 | 49 minutes ago | Scalable, Geo-Replicated, Transactional Datastore. |
| [dgraph](https://github.com/dgraph-io/dgraph) | 12233 | 370 | 2015-08-25 | 51 minutes ago | Scalable, Distributed, Low Latency, High Throughput Graph Database. |
| [bolt](https://github.com/boltdb/bolt) | 10426 | 343 | 2013-12-20 | 1 year ago | Low-level key/value database for Go. |
| [vitess](https://github.com/vitessio/vitess) | 9406 | 500 | 2013-06-27 | 3 hours ago | vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. |
| [groupcache](https://github.com/golang/groupcache) | 8162 | 462 | 2013-07-22 | 12 hours ago | Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. |
| [badger](https://github.com/dgraph-io/badger) | 7193 | 251 | 2017-01-26 | 5 hours ago | Fast key-value store in Go. |
| [pgweb](https://github.com/sosedoff/pgweb) | 6239 | 153 | 2014-10-09 | 2 weeks ago | Web-based PostgreSQL database browser. |
| [rqlite](https://github.com/rqlite/rqlite) | 5497 | 195 | 2014-08-23 | 6 days ago | The lightweight, distributed, relational database built on SQLite. |
| [kingshard](https://github.com/flike/kingshard) | 4980 | 402 | 2015-07-04 | 3 weeks ago | kingshard is a high performance proxy for MySQL powered by Golang. |
| [migrate](https://github.com/golang-migrate/migrate) | 3550 | 51 | 2018-01-19 | 1 hour ago | Database migrations. CLI and Golang library. |
| [bigcache](https://github.com/allegro/bigcache) | 3534 | 101 | 2016-03-23 | 20 hours ago | Efficient key/value cache for gigabytes of data. |
| [goleveldb](https://github.com/syndtr/goleveldb) | 3516 | 171 | 2013-01-23 | 2 days ago | Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go. |
| [go-cache](https://github.com/patrickmn/go-cache) | 3420 | 98 | 2012-01-02 | 1 month ago | In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. |
| [orchestrator](https://github.com/github/orchestrator) | 3324 | 243 | 2016-11-30 | 4 days ago | MySQL replication topology manager & visualizer. |
| [ledisdb](https://github.com/siddontang/ledisdb) | 3204 | 187 | 2014-04-30 | 1 month ago | Ledisdb is a high performance NoSQL like Redis based on LevelDB. |
| [bbolt](https://github.com/etcd-io/bbolt) | 2909 | 95 | 2017-06-17 | 18 hours ago | An embedded key/value database for Go. |
| [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) | 2748 | 164 | 2015-01-15 | 4 weeks ago | Sync your MySQL data into Elasticsearch automatically. |
| [squirrel](https://github.com/Masterminds/squirrel) | 2697 | 42 | 2014-01-18 | 1 week ago | Go library that helps you build SQL queries. |
| [buntdb](https://github.com/tidwall/buntdb) | 2605 | 96 | 2016-07-19 | 2 months ago | Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support. |
| [tiedot](https://github.com/HouzuoGuo/tiedot) | 2433 | 158 | 2013-05-26 | 2 months ago | Your NoSQL database powered by Golang. |
| [xo](https://github.com/xo/xo) | 2354 | 75 | 2016-02-05 | 3 weeks ago | Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. |
| [go-mysql](https://github.com/siddontang/go-mysql) | 2186 | 151 | 2014-02-21 | 1 day ago | Go toolset to handle MySQL protocol and replication. |
| [prest](https://github.com/prest/prest) | 2180 | 86 | 2016-11-22 | 7 months ago | Serve a RESTful API from any PostgreSQL database. |
| [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) | 1671 | 52 | 2018-09-30 | 1 hour ago | fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL. |
| [sql-migrate](https://github.com/rubenv/sql-migrate) | 1583 | 30 | 2014-09-09 | 2 days ago | Database migration tool. Allows embedding migrations into the application using go-bindata. |
| [cache2go](https://github.com/muesli/cache2go) | 1192 | 62 | 2013-11-11 | 3 months ago | In-memory key:value cache which supports automatic invalidation based on timeouts. |
| [nutsdb](https://github.com/xujiajun/nutsdb) | 1074 | 37 | 2018-12-07 | 1 month ago | Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set. |
| [gcache](https://github.com/bluele/gcache) | 1039 | 39 | 2015-01-24 | 1 month ago | Cache library with support for expirable Cache, LFU, LRU and ARC. |
| [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) | 981 | 66 | 2018-04-11 | 4 months ago | CovenantSQL is a SQL database on blockchain. |
| [gendry](https://github.com/didi/gendry) | 931 | 55 | 2017-12-01 | 1 week ago | Non-invasive SQL builder and powerful data binder. |
| [diskv](https://github.com/peterbourgon/diskv) | 844 | 36 | 2012-03-21 | 9 months ago | Home-grown disk-backed key-value store. |
| [moss](https://github.com/couchbase/moss) | 746 | 77 | 2016-02-06 | 2 months ago | Moss is a simple LSM key-value storage engine written in 100% Go. |
| [goqu](https://github.com/doug-martin/goqu) | 717 | 26 | 2015-02-21 | 17 hours ago | Idiomatic SQL builder and query library. |
| [fastcache](https://github.com/VictoriaMetrics/fastcache) | 669 | 18 | 2018-11-22 | 2 weeks ago | fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead. |
| [eliasdb](https://github.com/krotik/eliasdb) | 555 | 24 | 2016-08-13 | 3 weeks ago | Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language. |
| [skeema](https://github.com/skeema/skeema) | 509 | 32 | 2016-10-31 | 1 week ago | Pure-SQL schema management system for MySQL, with support for sharding and external online schema change tools. |
| [dotsql](https://github.com/gchaincl/dotsql) | 497 | 22 | 2014-11-20 | 5 months ago | Go library that helps you keep sql files in one place and use them with ease. |
| [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) | 455 | 24 | 2015-12-10 | 1 week ago | Powerful data retrieval methods as well as DB-agnostic query building capabilities. |
| [gormigrate](https://github.com/go-gormigrate/gormigrate) | 396 | 5 | 2016-08-31 | 1 month ago | Database schema migration helper for Gorm ORM. |
| [levigo](https://github.com/jmhodges/levigo) | 377 | 22 | 2012-01-17 | 1 month ago | Levigo is a Go wrapper for LevelDB. |
| [chproxy](https://github.com/Vertamedia/chproxy) | 358 | 25 | 2017-09-18 | 1 day ago | HTTP proxy for ClickHouse database. |
| [bitcask](https://github.com/prologic/bitcask) | 260 | 8 | 2019-03-12 | 17 hours ago | Bitcask is an embeddable, persistent and fast key-value (KV) database written in pure Go with predictable read/write performance, low latency and high throughput thanks to the bitcask on-disk layout (LSM+WAL). |
| [pudge](https://github.com/recoilme/pudge) | 238 | 7 | 2018-11-20 | 1 month ago | Fast and simple  key/value store written using Go's standard library. |
| [jet](https://github.com/go-jet/jet) | 205 | 8 | 2019-03-02 | 1 month ago | Framework for writing type-safe SQL queries in Go, with ability to easily convert database query result into desired arbitrary object structure. |
| [sqrl](https://github.com/elgris/sqrl) | 197 | 7 | 2014-06-25 | 3 weeks ago | SQL query builder, fork of Squirrel with improved performance. |
| [piladb](https://github.com/fern4lvarez/piladb) | 173 | 10 | 2015-09-08 | 1 year ago | Lightweight RESTful database engine based on stack data structures. |
| [vasto](https://github.com/chrislusf/vasto) | 168 | 19 | 2018-01-16 | 10 months ago | A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption. |
| [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) | 166 | 6 | 2017-04-29 | 2 months ago | Collects small insterts and sends big requests to ClickHouse servers. |
| [myreplication](https://github.com/2tvenom/myreplication) | 155 | 19 | 2015-02-04 | 1 year ago | MySql binary log replication listener. Supports statement and row based replication. |
| [kivik](https://github.com/go-kivik/kivik) | 140 | 5 | 2017-02-09 | 1 week ago | Kivik provides a common Go and GopherJS client library for CouchDB, PouchDB, and similar databases. |
| [goose](https://github.com/steinbacher/goose) | 131 | 4 | 2016-03-04 | 3 years ago | Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts. |
| [darwin](https://github.com/GuiaBolso/darwin) | 96 | 2 | 2016-04-05 | 1 month ago | Database schema evolution library for Go. |
| [slowpoke](https://github.com/recoilme/slowpoke) | 93 | 7 | 2018-02-19 | 3 months ago | Key-value store with persistence. |
| [dbq](https://github.com/rocketlaunchr/dbq) | 89 | 6 | 2019-07-11 | 2 days ago | Zero boilerplate database operations for Go. |
| [migrator](https://github.com/lopezator/migrator) | 85 | 4 | 2019-02-04 | 2 months ago | Dead simple Go database migration library. |
| [golang-scribble](https://github.com/nanobox-io/golang-scribble) | 82 | 3 | 2018-06-21 | 10 months ago | Tiny flat file JSON store. |
| [igor](https://github.com/galeone/igor) | 78 | 7 | 2016-03-10 | 1 month ago | Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax. |
| [octillery](https://github.com/knocknote/octillery) | 71 | 15 | 2018-11-26 | 4 months ago | Go package for sharding databases ( Supports every ORM or raw SQL ). |
| [godbal](https://github.com/xujiajun/godbal) | 50 | 4 | 2018-02-28 | 11 months ago | Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily. |
| [couchcache](https://github.com/codingsince1985/couchcache) | 47 | 3 | 2015-04-05 | 9 months ago | RESTful caching micro-service backed by Couchbase server. |
| [bcache](https://github.com/iwanbk/bcache) | 41 | 2 | 2018-12-26 | 8 months ago | Eventually consistent distributed in-memory  cache Go library. |
| [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) | 40 | 1 | 2018-08-11 | 2 months ago | A Go package to help write migrations with go-pg/pg. |
| [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) | 33 | 6 | 2017-12-18 | 2 years ago | BigCache with clustering support and individual item expiration. |
| [dbbench](https://github.com/sj14/dbbench) | 32 | 3 | 2018-11-24 | 1 week ago | Database benchmarking tool with support for several databases and scripts. |
| [cache](https://github.com/akyoto/cache) | 31 | 2 | 2019-05-11 | 2 months ago | In-memory key:value store with expiration time, 0 dependencies, <100 LoC, 100% coverage. |
| [gondolier](https://github.com/emvi/gondolier) | 27 | 4 | 2017-10-18 | 8 months ago | Database migration library using struct decorators. |
| [prep](https://github.com/hexdigest/prep) | 25 | 2 | 2017-12-11 | 2 years ago | Use prepared SQL statements without changing your code. |
| [pravasan](https://github.com/pravasan/pravasan) | 24 | 6 | 2015-01-03 | 1 year ago | Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc. |
| [go-fixtures](https://github.com/RichardKnop/go-fixtures) | 22 | 1 | 2015-12-24 | 3 weeks ago | Django style fixtures for Golang's excellent built-in database/sql library. |
| [coffer](https://github.com/claygod/coffer) | 19 | 3 | 2019-05-13 | 1 month ago | Simple ACID key-value database that supports transactions. |
| [tempdb](https://github.com/rafaeljesus/tempdb) | 14 | 3 | 2017-03-17 | 1 year ago | Key-value store for temporary items. |
| [datagen](https://github.com/codingconcepts/datagen) | 13 | 2 | 2019-04-18 | 2 months ago | A fast data generator that's multi-table aware and supports multi-row DML. |
| [gorocksdb](https://github.com/kapitan-k/gorocksdb) | 12 | 1 | 2017-12-28 | 2 years ago | Gorocksdb is a wrapper for [RocksDB](https://rocksdb.org) written in Go. |
| [avro](https://github.com/khezen/avro) | 10 | 1 | 2019-04-07 | 1 month ago | Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes. |
| [rwdb](https://github.com/andizzle/rwdb) | 10 | 2 | 2017-10-04 | 2 years ago | rwdb provides read replica capability for multiple database servers setup. |
| [sqlf](https://github.com/leporo/sqlf) | 7 | 1 | 2019-07-20 | 3 months ago | Fast SQL query builder. |
| [tracedb](https://github.com/unit-io/tracedb) | 6 | 1 | 2019-08-29 | 15 hours ago | Fast timeseries database for IoT, realtime messaging  applications. Access tracedb with pubsub over tcp or websocket using github.com/unit-io/trace application. |
| [bucket](https://github.com/PumpkinSeed/bucket) | 6 | 3 | 2019-09-22 | 1 week ago | Optimized data structure framework for Couchbase specialized on one bucket usage. |
| [schema](https://github.com/adlio/schema) | 4 | 1 | 2019-09-24 | 1 month ago | Library to embed schema migrations for database/sql-compatible databases inside your Go binaries. |
| [qry](https://github.com/HnH/qry) | 4 | 1 | 2019-08-20 | 2 months ago | Tool that generates constants from files with raw SQL queries. |
| [ormlite](https://github.com/pupizoid/ormlite) | 0 | 2 | 2018-06-28 | 1 day ago | Lightweight package containing some ORM-like features and helpers for sqlite databases. |

## Database Drivers
        
*Libraries for connecting and operating databases.*

### Relational Databases
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [mysql](https://github.com/go-sql-driver/mysql) | 8927 | 414 | 2012-12-09 | 2 weeks ago | MySQL driver for Go. |
| [pq](https://github.com/lib/pq) | 5590 | 158 | 2012-03-12 | 4 days ago | Pure Go Postgres driver for database/sql. |
| [go-sqlite3](https://github.com/mattn/go-sqlite3) | 3764 | 132 | 2011-11-11 | 1 week ago | SQLite3 driver for go that uses database/sql. |
| [pgx](https://github.com/jackc/pgx) | 2307 | 65 | 2013-03-30 | 1 week ago | PostgreSQL driver supporting features beyond those exposed by database/sql. |
| [go-mssqldb](https://github.com/denisenkom/go-mssqldb) | 1130 | 62 | 2013-12-16 | 7 hours ago | Microsoft MSSQL driver for Go. |
| [go-oci8](https://github.com/mattn/go-oci8) | 427 | 36 | 2012-02-29 | 1 month ago | Oracle driver for go that uses database/sql. |
| [goracle](https://github.com/go-goracle/goracle) | 279 | 30 | 2015-03-25 | 1 month ago | Oracle driver for Go, using the ODPI-C driver. |
| [firebirdsql](https://github.com/nakagami/firebirdsql) | 114 | 18 | 2013-08-27 | 1 month ago | Firebird RDBMS SQL driver for Go. |
| [go-adodb](https://github.com/mattn/go-adodb) | 99 | 12 | 2011-11-14 | 5 days ago | Microsoft ActiveX Object DataBase driver for go that uses database/sql. |
| [gofreetds](https://github.com/minus5/gofreetds) | 94 | 22 | 2012-12-06 | 11 months ago | Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org). |
| [calcite-avatica-go](https://github.com/apache/calcite-avatica-go) | 46 | 21 | 2017-08-08 | 4 months ago | Apache Avatica/Phoenix SQL driver for database/sql. |
| [bgc](https://github.com/viant/bgc) | 13 | 10 | 2016-06-13 | 1 week ago | Datastore Connectivity for BigQuery for go. |

### NoSQL Databases
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [cayley](https://github.com/cayleygraph/cayley) | 13137 | 610 | 2014-06-05 | 2 days ago | Graph database with support for multiple backends. |
| [redis](https://github.com/go-redis/redis) | 7884 | 238 | 2012-07-25 | 1 day ago | Redis client for Golang. |
| [redigo](https://github.com/gomodule/redigo) | 6919 | 298 | 2012-04-14 | 5 hours ago | Redigo is a Go client for the Redis database. |
| [bleve](https://github.com/blevesearch/bleve) | 6265 | 244 | 2014-04-17 | 3 days ago | Modern text indexing library for go. |
| [riot](https://github.com/go-ego/riot) | 5016 | 181 | 2017-06-21 | 1 month ago | Go Open Source, Distributed, Simple and efficient Search Engine. |
| [elastic](https://github.com/olivere/elastic) | 4675 | 163 | 2012-12-06 | 1 week ago | Elasticsearch client for Go. |
| [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) | 3926 | 144 | 2017-02-08 | 4 days ago | Official MongoDB driver for the Go language. |
| [go-elasticsearch](https://github.com/elastic/go-elasticsearch) | 2113 | 283 | 2017-03-27 | 1 day ago | Official Elasticsearch client for Go. |
| [mgo](https://github.com/globalsign/mgo) | 1740 | 71 | 2017-04-13 | 1 month ago | (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms. |
| [rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1495 | 47 | 2013-09-12 | 4 weeks ago | Go language driver for RethinkDB. |
| [elastigo](https://github.com/mattbaird/elastigo) | 954 | 47 | 2012-10-12 | 11 months ago | Elasticsearch client library. |
| [elasticsql](https://github.com/cch123/elasticsql) | 480 | 31 | 2016-08-24 | 10 months ago | Convert sql to elasticsearch dsl in Go. |
| [redeo](https://github.com/bsm/redeo) | 363 | 24 | 2014-03-06 | 5 months ago | Redis-protocol compatible TCP servers/services. |
| [neoism](https://github.com/jmcvetta/neoism) | 360 | 24 | 2012-07-12 | 11 months ago | Neo4j client for Golang. |
| [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) | 315 | 37 | 2014-07-26 | 23 hours ago | Aerospike client in Go language. |
| [gocb](https://github.com/couchbase/gocb) | 307 | 66 | 2015-01-15 | 1 day ago | Official Couchbase Go SDK. |
| [go-couchbase](https://github.com/couchbase/go-couchbase) | 297 | 24 | 2012-01-19 | 1 month ago | Couchbase client in Go. |
| [gokv](https://github.com/philippgille/gokv) | 158 | 6 | 2018-10-08 | 1 day ago | Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more). |
| [cachego](https://github.com/faabiosr/cachego) | 114 | 6 | 2016-10-05 | 2 hours ago | Golang Cache component for multiple drivers. |
| [go-rejson](https://github.com/nitishm/go-rejson) | 109 | 5 | 2018-04-23 | 2 hours ago | Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease. |
| [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) | 74 | 6 | 2011-06-04 | 1 year ago | Neo4j REST Client in golang. |
| [skizze](https://github.com/seiflotfy/skizze) | 69 | 6 | 2016-01-17 | 3 years ago | probabilistic data-structures service and storage. |
| [godis](https://github.com/piaohao/godis) | 67 | 8 | 2019-06-14 | 5 months ago | redis client implement by golang, inspired by jedis. |
| [arangolite](https://github.com/solher/arangolite) | 66 | 5 | 2015-10-04 | 9 months ago | Lightweight golang driver for ArangoDB. |
| [dynago](https://github.com/underarmour/dynago) | 66 | 126 | 2015-05-18 | 2 years ago | Dynago is a principle of least surprise client for DynamoDB. |
| [mgm](https://github.com/Kamva/mgm) | 42 | 2 | 2019-12-27 | 1 day ago | MongoDB model-based ODM for Go (based on official MongoDB driver). |
| [go-pilosa](https://github.com/pilosa/go-pilosa) | 34 | 15 | 2016-09-30 | 1 month ago | Go client library for Pilosa. |
| [goforestdb](https://github.com/couchbase/goforestdb) | 28 | 41 | 2014-05-14 | 3 years ago | Go bindings for ForestDB. |
| [neo4j](https://github.com/cihangir/neo4j) | 25 | 3 | 2013-05-18 | 4 years ago | Neo4j Rest API Bindings for Golang. |
| [goriak](https://github.com/zegl/goriak) | 24 | 2 | 2016-10-05 | 1 year ago | Go language driver for Riak KV. |
| [goes](https://github.com/OwnLocal/goes) | 24 | 33 | 2015-12-28 | 2 years ago | Library to interact with Elasticsearch. |
| [dsc](https://github.com/viant/dsc) | 16 | 15 | 2016-06-13 | 4 months ago | Datastore connectivity for SQL, NoSQL, structured files. |
| [xredis](https://github.com/shomali11/xredis) | 10 | 1 | 2017-06-14 | 7 months ago | Typesafe, customizable, clean & easy to use Redis client. |
| [godscache](https://github.com/defcronyke/godscache) | 6 | 1 | 2018-05-08 | 11 months ago | A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached. |
| [asc](https://github.com/viant/asc) | 4 | 10 | 2016-06-13 | 9 months ago | Datastore Connectivity for Aerospike for go. |

## Date and Time
        
*Libraries for working with dates and times.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [now](https://github.com/jinzhu/now) | 2368 | 52 | 2013-11-18 | 2 months ago | Now is a time toolkit for golang. |
| [dateparse](https://github.com/araddon/dateparse) | 962 | 15 | 2014-04-21 | 4 months ago | Parse date's without knowing format in advance. |
| [carbon](https://github.com/uniplaces/carbon) | 384 | 39 | 2016-08-03 | 1 year ago | Simple Time extension with a lot of util methods, ported from PHP Carbon library. |
| [durafmt](https://github.com/hako/durafmt) | 273 | 4 | 2016-05-20 | 1 month ago | Time duration formatting library for Go. |
| [timeutil](https://github.com/leekchan/timeutil) | 174 | 7 | 2015-08-02 | 11 months ago | Useful extensions (Timedelta, Strftime, ...) to the golang's time package. |
| [iso8601](https://github.com/relvacode/iso8601) | 70 | 2 | 2017-04-25 | 1 year ago | Efficiently parse ISO8601 date-times without regex. |
| [timespan](https://github.com/SaidinWoT/timespan) | 67 | 5 | 2014-10-07 | 10 months ago | For interacting with intervals of time, defined as a start time and a duration. |
| [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) | 66 | 3 | 2016-01-31 | 2 months ago | The implementation of the Persian (Solar Hijri) Calendar in Go (golang). |
| [date](https://github.com/rickb777/date) | 34 | 2 | 2015-11-23 | 4 months ago | Augments Time for working with dates, date ranges, time spans, periods, and time-of-day. |
| [feiertage](https://github.com/wlbr/feiertage) | 24 | 1 | 2015-11-04 | 3 months ago | Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving... |
| [go-sunrise](https://github.com/nathan-osman/go-sunrise) | 17 | 3 | 2017-06-15 | 2 years ago | Calculate the sunrise and sunset times for a given location. |
| [kair](https://github.com/GuilhermeCaruso/kair) | 12 | 1 | 2018-10-03 | 11 months ago | Date and Time - Golang Formatting Library. |
| [nulltime](https://github.com/kirillDanshin/nulltime) | 9 | 1 | 2016-03-06 | 2 years ago | Nullable `time.Time`. |
| [cronrange](https://github.com/1set/cronrange) | 7 | 0 | 2019-11-10 | 2 months ago | Parses Cron-style time range expressions, checks if the given time is within any ranges. |
| [tuesday](https://github.com/osteele/tuesday) | 7 | 1 | 2017-08-10 | 2 years ago | Ruby-compatible Strftime function. |
| [strftime](https://github.com/awoodbeck/strftime) | 4 | 1 | 2018-02-10 | 1 year ago | C99-compatible strftime formatter. |
| [go-week](https://github.com/stoewer/go-week) | 2 | 3 | 2018-02-23 | 2 months ago | An efficient package to work with ISO8601 week dates. |

## Distributed Systems
        
*Packages that help with building Distributed Systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [kit](https://github.com/go-kit/kit) | 16076 | 675 | 2015-02-03 | 1 week ago | Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc. |
| [grpc-go](https://github.com/grpc/grpc-go) | 10522 | 468 | 2014-12-08 | 4 days ago | The Go language implementation of gRPC. HTTP/2 based RPC. |
| [micro](https://github.com/micro/micro) | 7448 | 319 | 2015-01-16 | 56 minutes ago | Pluggable microservice toolkit and distributed systems platform. |
| [nats-server](https://github.com/nats-io/nats-server) | 7163 | 362 | 2012-10-29 | 56 minutes ago | Lightweight, high performance messaging system for microservices, IoT, and cloud native systems. |
| [rpcx](https://github.com/smallnest/rpcx) | 4288 | 310 | 2016-05-18 | 2 weeks ago | Distributed pluggable RPC service framework like alibaba Dubbo. |
| [tendermint](https://github.com/tendermint/tendermint) | 3470 | 260 | 2014-05-14 | 4 hours ago | High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols. |
| [torrent](https://github.com/anacrolix/torrent) | 3292 | 128 | 2015-01-08 | 1 day ago | BitTorrent client package. |
| [raft](https://github.com/hashicorp/raft) | 3205 | 253 | 2013-11-05 | 4 days ago | Golang implementation of the Raft consensus protocol, by HashiCorp. |
| [dragonboat](https://github.com/lni/dragonboat) | 2857 | 135 | 2018-12-23 | 1 week ago | A feature complete and high performance multi-group Raft library in Go. |
| [glow](https://github.com/chrislusf/glow) | 2733 | 140 | 2015-06-14 | 1 year ago | Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go. |
| [gleam](https://github.com/chrislusf/gleam) | 2371 | 143 | 2016-08-26 | 3 months ago | Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed. |
| [krakend](https://github.com/devopsfaith/krakend) | 2289 | 84 | 2016-11-04 | 5 days ago | Ultra performant API Gateway framework with middlewares. |
| [emitter](https://github.com/emitter-io/emitter) | 2201 | 90 | 2016-10-29 | 1 week ago | High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love. |
| [liftbridge](https://github.com/liftbridge-io/liftbridge) | 1437 | 65 | 2017-10-13 | 1 day ago | Lightweight, fault-tolerant message streams for NATS. |
| [hprose-golang](https://github.com/hprose/hprose-golang) | 1064 | 97 | 2014-02-14 | 6 months ago | Very newbility RPC Library, support 25+ languages now. |
| [ringpop-go](https://github.com/uber/ringpop-go) | 601 | 2365 | 2015-06-05 | 11 months ago | Scalable, fault-tolerant application-layer sharding for Go applications. |
| [gorpc](https://github.com/valyala/gorpc) | 576 | 28 | 2014-11-20 | 4 months ago | Simple, fast and scalable RPC library for high load. |
| [go-health](https://github.com/InVisionApp/go-health) | 524 | 26 | 2017-11-29 | 1 week ago | Library for enabling asynchronous dependency health checks in your service. |
| [rain](https://github.com/cenkalti/rain) | 395 | 14 | 2014-05-21 | 1 week ago | BitTorrent client and library. |
| [digota](https://github.com/digota/digota) | 322 | 24 | 2017-08-14 | 1 year ago | grpc ecommerce microservice. |
| [sleuth](https://github.com/ursiform/sleuth) | 312 | 8 | 2016-04-23 | 1 year ago | Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)). |
| [go-sundheit](https://github.com/AppsFlyer/go-sundheit) | 281 | 9 | 2019-04-08 | 4 weeks ago | A library built to provide support for defining async service health checks for golang services. |
| [go-jump](https://github.com/dgryski/go-jump) | 279 | 12 | 2014-06-15 | 1 year ago | Port of Google's "Jump" Consistent Hash function. |
| [consistent](https://github.com/buraksezer/consistent) | 258 | 10 | 2018-03-25 | 3 months ago | Consistent hashing with bounded loads. |
| [dht](https://github.com/anacrolix/dht) | 145 | 11 | 2016-12-14 | 2 weeks ago | BitTorrent Kademlia DHT implementation. |
| [jsonrpc](https://github.com/osamingo/jsonrpc) | 118 | 6 | 2016-10-28 | 3 weeks ago | The jsonrpc package helps implement of JSON-RPC 2.0. |
| [jsonrpc](https://github.com/ybbus/jsonrpc) | 111 | 9 | 2016-11-10 | 6 days ago | JSON-RPC 2.0 HTTP client implementation. |
| [redislock](https://github.com/bsm/redislock) | 85 | 2 | 2019-06-24 | 1 month ago | Simplified distributed locking implementation using Redis. |
| [celeriac.v1](https://github.com/svcavallar/celeriac.v1) | 59 | 3 | 2015-10-10 | 1 year ago | Library for adding support for interacting and monitoring Celery workers, tasks and events in Go. |
| [doublejump](https://github.com/edwingeng/doublejump) | 48 | 2 | 2018-06-26 | 1 year ago | A revamped Google's jump consistent hash. |
| [drmaa](https://github.com/dgruber/drmaa) | 28 | 2 | 2013-03-17 | 1 year ago | Job submission library for cluster schedulers based on the DRMAA standard. |
| [flowgraph](https://github.com/vectaport/flowgraph) | 25 | 1 | 2018-08-29 | 5 months ago | flow-based programming package. |
| [outboxer](https://github.com/italolelis/outboxer) | 23 | 0 | 2019-02-01 | 12 hours ago | Outboxer is a go library that implements the outbox pattern. |
| [dynatomic](https://github.com/tylfin/dynatomic) | 10 | 0 | 2019-02-08 | 11 months ago | A library for using DynamoDB as an atomic counter. |

## Dynamic DNS
        
*Tools for updating dynamic DNS records.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [godns](https://github.com/TimothyYe/godns) | 521 | 23 | 2014-05-11 | 5 days ago | A dynamic DNS client tool, supports DNSPod & HE.net, written in Go. |
| [ddns](https://github.com/skibish/ddns) | 123 | 6 | 2017-03-13 | 3 months ago | Personal DDNS client with Digital Ocean Networking DNS as backend. |

## Email
        
*Libraries and tools that implement email creation and sending.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [MailHog](https://github.com/mailhog/MailHog) | 5818 | 124 | 2014-04-16 | 4 weeks ago | Email and SMTP testing with web and API interface. |
| [hermes](https://github.com/matcornic/hermes) | 1758 | 28 | 2017-03-25 | 4 weeks ago | Golang package that generates clean, responsive HTML e-mails. |
| [email](https://github.com/jordan-wright/email) | 1194 | 39 | 2013-12-12 | 3 hours ago | A robust and flexible email library for Go. |
| [go-imap](https://github.com/emersion/go-imap) | 862 | 33 | 2016-04-26 | 1 day ago | IMAP library for clients and servers. |
| [sendgrid-go](https://github.com/sendgrid/sendgrid-go) | 555 | 191 | 2013-09-12 | 3 days ago | SendGrid's Go library for sending email. |
| [mailgun-go](https://github.com/mailgun/mailgun-go) | 427 | 54 | 2014-02-28 | 2 weeks ago | Go library for sending mail with the Mailgun API. |
| [hectane](https://github.com/hectane/hectane) | 173 | 13 | 2015-08-28 | 5 months ago | Lightweight SMTP client providing an HTTP API. |
| [douceur](https://github.com/aymerick/douceur) | 167 | 3 | 2015-04-09 | 1 year ago | CSS inliner for your HTML emails. |
| [go-message](https://github.com/emersion/go-message) | 131 | 9 | 2016-12-31 | 1 day ago | Streaming library for the Internet Message Format and mail messages. |
| [smtp](https://github.com/mailhog/smtp) | 53 | 9 | 2014-12-24 | 1 year ago | SMTP server protocol state machine. |
| [go-dkim](https://github.com/toorop/go-dkim) | 52 | 2 | 2015-04-29 | 3 months ago | DKIM library, to sign & verify email. |
| [go-premailer](https://github.com/vanng822/go-premailer) | 43 | 2 | 2015-02-16 | 1 month ago | Inline styling for HTML mail in Go. |
| [mailchain](https://github.com/mailchain/mailchain) | 34 | 5 | 2019-04-11 | 1 week ago | Send encrypted emails to blockchain addresses written in Go. |
| [go-simple-mail](https://github.com/xhit/go-simple-mail) | 13 | 1 | 2019-09-15 | 1 month ago | Very simple package to send emails with SMTP Keep Alive and two timeouts: Connect and Send. |

## Embeddable Scripting Languages
        
*Embedding other languages inside your go code.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [otto](https://github.com/robertkrimen/otto) | 5071 | 191 | 2012-10-06 | 1 month ago | JavaScript interpreter written in Go. |
| [gopher-lua](https://github.com/yuin/gopher-lua) | 3236 | 144 | 2015-02-15 | 19 hours ago | Lua 5.1 VM and compiler written in Go. |
| [go-lua](https://github.com/Shopify/go-lua) | 1769 | 269 | 2013-12-20 | 2 months ago | Port of the Lua 5.2 VM to pure Go. |
| [tengo](https://github.com/d5/tengo) | 1531 | 34 | 2019-01-09 | 1 day ago | Bytecode compiled script language for Go. |
| [expr](https://github.com/antonmedv/expr) | 1062 | 30 | 2018-07-14 | 18 hours ago | an engine that can evaluate expressions. |
| [go-python](https://github.com/sbinet/go-python) | 991 | 45 | 2012-07-09 | 6 months ago | naive go bindings to the CPython C-API. |
| [anko](https://github.com/mattn/anko) | 977 | 47 | 2014-03-28 | 1 month ago | Scriptable interpreter written in Go. |
| [go-php](https://github.com/deuill/go-php) | 724 | 43 | 2015-09-17 | 1 year ago | PHP bindings for Go. |
| [go-duktape](https://github.com/olebedev/go-duktape) | 678 | 26 | 2015-01-08 | 6 months ago | Duktape JavaScript engine bindings for Go. |
| [golua](https://github.com/aarzilli/golua) | 457 | 33 | 2010-12-06 | 1 day ago | Go bindings for Lua C API. |
| [gisp](https://github.com/jcla1/gisp) | 432 | 20 | 2014-01-11 | 2 years ago | Simple LISP in Go. |
| [cel-go](https://github.com/google/cel-go) | 363 | 24 | 2018-03-09 | 3 days ago | Fast, portable, non-Turing complete expression evaluation with gradual typing. |
| [gval](https://github.com/PaesslerAG/gval) | 175 | 15 | 2017-09-27 | 5 months ago | A highly customizable expression language written in Go. |
| [gentee](https://github.com/gentee/gentee) | 40 | 3 | 2018-01-14 | 1 day ago | Embeddable scripting programming language. |
| [binder](https://github.com/alexeyco/binder) | 34 | 2 | 2017-04-02 | 1 year ago | Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua). |
| [purl](https://github.com/ian-kent/purl) | 29 | 2 | 2014-11-29 | 5 years ago | Perl 5.18.2 embedded in Go. |
| [ngaro](https://github.com/db47h/ngaro) | 19 | 1 | 2016-08-09 | 1 year ago | Embeddable Ngaro VM implementation enabling scripting in Retro. |

## Error Handling
        
*Libraries for handling errors.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [errors](https://github.com/pkg/errors) | 5472 | 109 | 2015-12-27 | 6 days ago | Package that provides simple error handling primitives. |
| [go-multierror](https://github.com/hashicorp/go-multierror) | 804 | 165 | 2014-12-15 | 3 weeks ago | Go (golang) package for representing a list of errors as a single error. |
| [errorx](https://github.com/joomcode/errorx) | 603 | 52 | 2018-08-17 | 4 months ago | A feature rich error package with stack traces, composition of errors and more. |
| [tracerr](https://github.com/ztrue/tracerr) | 580 | 8 | 2019-02-06 | 10 months ago | Golang errors with stack trace and source fragments. |
| [eris](https://github.com/rotisserie/eris) | 450 | 8 | 2019-09-07 | 3 days ago | A better way to handle, trace, and log errors in Go. Compatible with the standard error library and github.com/pkg/errors. |
| [errlog](https://github.com/snwfdhmp/errlog) | 275 | 6 | 2019-02-16 | 1 month ago | Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place. |
| [emperror](https://github.com/emperror/emperror) | 99 | 2 | 2017-06-13 | 5 days ago | Error handling tools and best practices for Go libraries and applications. |
| [errors](https://github.com/emperror/errors) | 40 | 3 | 2019-07-09 | 1 week ago | Drop-in replacement for the standard library errors package and github.com/pkg/errors. Provides various error handling primitives. |
| [werr](https://github.com/txgruppi/werr) | 12 | 0 | 2015-08-04 | 3 years ago | Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called. |
| [errors](https://github.com/neuronlabs/errors) | 4 | 1 | 2019-07-26 | 5 months ago | Simple golang error handling with classification primitives. |
| [falcon](https://github.com/SonicRoshan/falcon) | 2 | 1 | 2019-09-09 | 3 months ago | A Simple Yet Highly Powerful Package For Error Handling. |

## Files
        
*Libraries for handling files and file systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [afero](https://github.com/spf13/afero) | 2513 | 93 | 2014-10-28 | 2 weeks ago | FileSystem Abstraction System for Go. |
| [pdfcpu](https://github.com/pdfcpu/pdfcpu) | 1287 | 37 | 2017-06-18 | 1 week ago | PDF processor. |
| [notify](https://github.com/rjeczalik/notify) | 523 | 29 | 2014-09-08 | 2 weeks ago | File system event notification library with simple API, similar to os/signal. |
| [bigfile](https://github.com/bigfile/bigfile) | 101 | 9 | 2019-07-15 | 2 months ago | A file transfer system, support to manage files with http api, rpc call and ftp client. |
| [go-csv-tag](https://github.com/artonge/go-csv-tag) | 63 | 1 | 2017-06-18 | 1 month ago | Load csv file using tag. |
| [opc](https://github.com/qmuntal/opc) | 63 | 1 | 2018-11-06 | 3 months ago | Load Open Packaging Conventions (OPC) files for Go. |
| [skywalker](https://github.com/dixonwille/skywalker) | 54 | 3 | 2017-08-01 | 2 years ago | Package to allow one to concurrently go through a filesystem with ease. |
| [vfs](https://github.com/C2FO/vfs) | 41 | 23 | 2017-08-01 | 4 days ago | A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS. |
| [tarfs](https://github.com/posener/tarfs) | 39 | 2 | 2017-03-10 | 2 years ago | Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files. |
| [afs](https://github.com/viant/afs) | 30 | 8 | 2019-08-19 | 1 day ago | Abstract File Storage (mem, scp, zip, tar, cloud: s3, gs) for Go. |
| [go-exiftool](https://github.com/barasher/go-exiftool) | 21 | 2 | 2019-05-12 | 1 week ago | Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...). |
| [go-gtfs](https://github.com/artonge/go-gtfs) | 18 | 2 | 2017-07-09 | 2 months ago | Load gtfs files in go. |
| [checksum](https://github.com/codingsince1985/checksum) | 15 | 1 | 2014-11-05 | 4 months ago | Compute message digest, like MD5 and SHA256, for large files. |
| [flop](https://github.com/homedepot/flop) | 14 | 11 | 2019-03-01 | 4 months ago | File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html). |
| [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) | 14 | 2 | 2018-10-16 | 2 weeks ago | Copy files for humans. |
| [parquet](https://github.com/parsyl/parquet) | 5 | 4 | 2019-01-29 | 3 months ago | Read and write [parquet](https://parquet.apache.org) files. |

## Financial
        
*Packages for accounting and finance.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [decimal](https://github.com/shopspring/decimal) | 1886 | 62 | 2015-02-25 | 2 weeks ago | Arbitrary-precision fixed-point decimal numbers. |
| [go-money](https://github.com/Rhymond/go-money) | 686 | 17 | 2017-03-20 | 2 weeks ago | Implementation of Fowler's Money pattern. |
| [go-finance](https://github.com/FlashBoys/go-finance) | 537 | 27 | 2016-02-28 | 1 year ago | Comprehensive financial markets data in Go. |
| [accounting](https://github.com/leekchan/accounting) | 531 | 11 | 2015-08-10 | 4 weeks ago | money and currency formatting for golang. |
| [techan](https://github.com/sdcoffey/techan) | 215 | 23 | 2017-03-08 | 3 weeks ago | Technical analysis library with advanced market analysis and trading strategies. |
| [orderbook](https://github.com/i25959341/orderbook) | 116 | 9 | 2018-04-24 | 8 months ago | Matching Engine for Limit Order Book in Golang. |
| [ofxgo](https://github.com/aclindsa/ofxgo) | 72 | 9 | 2015-11-08 | 6 months ago | Query OFX servers and/or parse the responses (with example command-line client). |
| [vat](https://github.com/dannyvankooten/vat) | 64 | 1 | 2016-06-18 | 1 year ago | VAT number validation & EU VAT rates. |
| [transaction](https://github.com/claygod/transaction) | 63 | 8 | 2017-10-11 | 3 months ago | Embedded transactional database of accounts, running in multithreaded mode. |
| [go-finance](https://github.com/alpeb/go-finance) | 48 | 2 | 2017-06-01 | 1 year ago | Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations. |
| [currency](https://github.com/bnkamalesh/currency) | 19 | 5 | 2017-05-09 | 2 months ago | High performant & accurate currency computation package. |
| [go-finance](https://github.com/pieterclaerhout/go-finance) | 2 | 1 | 2019-09-30 | 3 months ago | Module to fetch exchange rates, check VAT numbers via VIES and check IBAN bank account numbers. |

## Forms
        
*Libraries for working with forms.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [nosurf](https://github.com/justinas/nosurf) | 1023 | 36 | 2013-08-22 | 4 weeks ago | CSRF protection middleware for Go. |
| [binding](https://github.com/mholt/binding) | 763 | 31 | 2014-05-20 | 1 year ago | Binds form and JSON data from net/http Request to struct. |
| [csrf](https://github.com/gorilla/csrf) | 491 | 20 | 2015-08-03 | 2 months ago | CSRF protection for Go web applications & services. |
| [form](https://github.com/go-playground/form) | 388 | 11 | 2016-05-26 | 2 months ago | Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support. |
| [conform](https://github.com/leebenson/conform) | 184 | 5 | 2016-01-05 | 5 months ago | Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags. |
| [formam](https://github.com/monoculum/formam) | 135 | 3 | 2014-10-25 | 3 weeks ago | decode form's values into a struct. |
| [forms](https://github.com/albrow/forms) | 106 | 6 | 2014-08-07 | 2 years ago | Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files. |
| [bind](https://github.com/robfig/bind) | 24 | 2 | 2014-08-06 | 5 years ago | Bind form data to any Go values. |
| [queryparam](https://github.com/TomWright/queryparam) | 1 | 1 | 2018-06-14 | 1 month ago | Decode `url.Values` into usable struct values of standard or custom types. |

## Functional
        
*Packages to support functional programming in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-underscore](https://github.com/tobyhede/go-underscore) | 1113 | 28 | 2014-07-02 | 11 months ago | Useful collection of helpfully functional Go collection utilities. |
| [fpGo](https://github.com/TeaEntityLab/fpGo) | 134 | 3 | 2018-05-24 | 1 year ago | Monad, Functional Programming features for Golang. |
| [fuego](https://github.com/seborama/fuego) | 60 | 2 | 2018-11-05 | 5 months ago | Functional Experiment in Go. |

## Game Development
        
*Awesome game development libraries.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [leaf](https://github.com/name5566/leaf) | 3360 | 320 | 2014-08-04 | 2 months ago | Lightweight game server framework. |
| [pixel](https://github.com/faiface/pixel) | 2698 | 95 | 2016-11-19 | 1 week ago | Hand-crafted 2D game library in Go. |
| [ebiten](https://github.com/hajimehoshi/ebiten) | 2444 | 85 | 2013-06-16 | 2 hours ago | dead simple 2D game library in Go. |
| [goworld](https://github.com/xiaonanln/goworld) | 1373 | 115 | 2017-06-03 | 3 weeks ago | Scalable game server engine, featuring space-entity framework and hot-swapping. |
| [go-sdl2](https://github.com/veandco/go-sdl2) | 1257 | 40 | 2013-06-05 | 2 days ago | Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/). |
| [nano](https://github.com/lonng/nano) | 1170 | 56 | 2017-08-02 | 1 month ago | Lightweight, facility, high performance golang based game server framework. |
| [engo](https://github.com/EngoEngine/engo) | 1146 | 46 | 2014-11-12 | 1 week ago | Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm. |
| [gonet](https://github.com/xtaci/gonet) | 1080 | 133 | 2013-04-11 | 2 years ago | Game server skeleton implemented with golang. |
| [termloop](https://github.com/JoelOtter/termloop) | 1077 | 33 | 2015-05-23 | 2 months ago | Terminal-based game engine for Go, built on top of Termbox. |
| [engine](https://github.com/g3n/engine) | 922 | 60 | 2017-03-07 | 3 days ago | Go 3D Game Engine. |
| [oak](https://github.com/oakmound/oak) | 696 | 41 | 2017-07-15 | 1 week ago | Pure Go game engine. |
| [pitaya](https://github.com/topfreegames/pitaya) | 443 | 42 | 2018-03-19 | 2 weeks ago | Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. |
| [engine](https://github.com/azul3d/engine) | 441 | 23 | 2016-02-29 | 3 months ago | 3D game engine written in Go. |
| [raylib-go](https://github.com/gen2brain/raylib-go) | 427 | 21 | 2017-01-27 | 1 day ago | Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming. |
| [go-astar](https://github.com/beefsack/go-astar) | 348 | 9 | 2014-05-28 | 1 year ago | Go implementation of the A\* path finding algorithm. |
| [GarageEngine](https://github.com/vova616/GarageEngine) | 318 | 31 | 2012-08-07 | 5 months ago | 2d game engine written in Go working on OpenGL. |
| [go3d](https://github.com/ungerik/go3d) | 171 | 10 | 2011-06-27 | 10 months ago | Performance oriented 2D/3D math package for Go. |
| [glop](https://github.com/runningwild/glop) | 76 | 3 | 2011-04-20 | 4 years ago | Glop (Game Library Of Power) is a fairly simple cross-platform game library. |
| [go-collada](https://github.com/GlenKelley/go-collada) | 15 | 3 | 2013-09-19 | 6 years ago | Go package for working with the Collada file format. |

## Generation and Generics
        
*Tools to enhance the language with features like generics via code generation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-linq](https://github.com/ahmetb/go-linq) | 1950 | 67 | 2013-12-19 | 2 months ago | .NET LINQ-like query methods for Go. |
| [jennifer](https://github.com/dave/jennifer) | 1428 | 23 | 2016-12-04 | 1 month ago | Generate arbitrary Go code without templates. |
| [gen](https://github.com/clipperhouse/gen) | 1081 | 34 | 2013-10-13 | 1 week ago | Code generation tool for ‘generics’-like functionality. |
| [goderive](https://github.com/awalterschulze/goderive) | 778 | 25 | 2017-02-10 | 5 months ago | Derives functions from input types. |
| [gowrap](https://github.com/hexdigest/gowrap) | 310 | 11 | 2018-09-15 | 4 months ago | Generate decorators for Go interfaces using simple templates. |
| [interfaces](https://github.com/rjeczalik/interfaces) | 201 | 7 | 2015-12-06 | 3 months ago | Command line tool for generating interface definitions. |
| [go-enum](https://github.com/abice/go-enum) | 94 | 3 | 2017-08-10 | 2 days ago | Code generation for enums from code comments. |
| [pkgreflect](https://github.com/ungerik/pkgreflect) | 92 | 5 | 2012-09-03 | 2 years ago | Go preprocessor for package scoped reflection. |
| [efaceconv](https://github.com/t0pep0/efaceconv) | 44 | 4 | 2016-11-18 | 2 years ago | Code generation tool for high performance conversion from interface{} to immutable type without allocations. |
| [gotype](https://github.com/wzshiming/gotype) | 28 | 4 | 2017-12-05 | 4 days ago | Golang source code parsing, usage like reflect package. |
| [GENERIS](https://github.com/senselogic/GENERIS) | 20 | 1 | 2019-03-10 | 2 months ago | Code generation tool providing generics, free-form macros, conditional compilation and HTML templating. |
| [go-xray](https://github.com/pieterclaerhout/go-xray) | 5 | 1 | 2019-10-01 | 2 months ago | Helpers for making the use of reflection easier. |

## Geographic
        
*Geographic tools and servers*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [tile38](https://github.com/tidwall/tile38) | 6628 | 205 | 2016-03-04 | 1 month ago | Geolocation DB with spatial index and realtime geofencing. |
| [geo](https://github.com/golang/geo) | 987 | 77 | 2014-12-03 | 4 months ago | S2 geometry library in Go. |
| [geocache](https://github.com/melihmucuk/geocache) | 117 | 6 | 2016-06-21 | 3 years ago | In-memory cache that is suitable for geolocation based applications. |
| [osm](https://github.com/paulmach/osm) | 95 | 8 | 2016-02-02 | 3 months ago | Library for reading, writing and working with OpenStreetMap data and APIs. |
| [wgs84](https://github.com/wroge/wgs84) | 43 | 1 | 2019-06-08 | 2 months ago | Library for Coordinate Conversion and Transformation (ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM). |
| [geoserver](https://github.com/hishamkaram/geoserver) | 27 | 1 | 2018-03-26 | 4 months ago | geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API. |
| [gismanager](https://github.com/hishamkaram/gismanager) | 22 | 1 | 2018-09-29 | 1 year ago | Publish Your GIS Data(Vector Data) to PostGIS and Geoserver. |
| [pbf](https://github.com/maguro/pbf) | 20 | 2 | 2017-09-18 | 2 weeks ago | OpenStreetMap PBF golang encoder/decoder. |

## Go Compilers
        
*Tools for compiling Go to other languages.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gopherjs](https://github.com/gopherjs/gopherjs) | 9078 | 257 | 2013-08-27 | 2 months ago | Compiler from Go to JavaScript. |
| [llgo](https://github.com/go-llvm/llgo) | 1029 | 63 | 2011-11-05 | 5 years ago | LLVM-based compiler for Go. |
| [tardisgo](https://github.com/tardisgo/tardisgo) | 394 | 27 | 2014-01-08 | 3 years ago | Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler. |
| [c4go](https://github.com/Konstantin8105/c4go) | 191 | 15 | 2018-03-28 | 3 weeks ago | Transpile C code to Go code. |
| [f4go](https://github.com/Konstantin8105/f4go) | 23 | 3 | 2018-07-08 | 2 months ago | Transpile FORTRAN 77 code to Go code. |

## Goroutines
        
*Tools for managing and working with Goroutines.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [ants](https://github.com/panjf2000/ants) | 2961 | 99 | 2018-05-19 | 4 days ago | A high-performance and low-cost goroutine pool in Go. |
| [goworker](https://github.com/benmanns/goworker) | 2324 | 77 | 2013-07-22 | 1 month ago | goworker is a Go-based background worker. |
| [tunny](https://github.com/Jeffail/tunny) | 1530 | 62 | 2014-04-02 | 3 months ago | Goroutine pool for golang. |
| [pool](https://github.com/go-playground/pool) | 539 | 14 | 2015-10-28 | 4 months ago | Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation. |
| [grpool](https://github.com/ivpusic/grpool) | 537 | 29 | 2015-07-22 | 11 months ago | Lightweight Goroutine pool. |
| [workerpool](https://github.com/gammazero/workerpool) | 238 | 8 | 2016-05-17 | 1 week ago | Goroutine pool that limits the concurrency of task execution, not the number of tasks queued. |
| [go-floc](https://github.com/workanator/go-floc) | 175 | 8 | 2017-07-03 | 2 years ago | Orchestrate goroutines with ease. |
| [gowp](https://github.com/xxjwxc/gowp) | 147 | 9 | 2019-09-14 | 2 weeks ago | gowp is concurrency limiting goroutine pool. |
| [go-flow](https://github.com/kamildrazkiewicz/go-flow) | 121 | 8 | 2016-09-25 | 8 months ago | Control goroutines execution order. |
| [GoSlaves](https://github.com/dgrr/GoSlaves) | 88 | 4 | 2017-09-17 | 5 months ago | Simple and Asynchronous Goroutine pool library. |
| [semaphore](https://github.com/kamilsk/semaphore) | 82 | 1 | 2016-10-08 | 1 week ago | Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context. |
| [semaphore](https://github.com/marusama/semaphore) | 80 | 4 | 2017-11-22 | 1 year ago | Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations). |
| [gpool](https://github.com/sherifabdlnaby/gpool) | 62 | 1 | 2018-12-03 | 1 month ago | manages a resizeable pool of context-aware goroutines to bound concurrency. |
| [breaker](https://github.com/kamilsk/breaker) | 51 | 3 | 2019-02-15 | 1 week ago | Flexible mechanism to make execution flow interruptible. |
| [worker-pool](https://github.com/vardius/worker-pool) | 51 | 4 | 2017-10-04 | 5 months ago | goworker is a Go simple async worker pool. |
| [cyclicbarrier](https://github.com/marusama/cyclicbarrier) | 43 | 3 | 2018-01-11 | 1 year ago | CyclicBarrier for golang. |
| [gollback](https://github.com/vardius/gollback) | 31 | 1 | 2019-05-11 | 5 months ago | asynchronous simple function utilities, for managing execution of closures and callbacks. |
| [async](https://github.com/StudioSol/async) | 29 | 9 | 2017-06-30 | 3 months ago | A safe way to execute functions asynchronously, recovering them in case of panic. |
| [parallel-fn](https://github.com/rafaeljesus/parallel-fn) | 26 | 3 | 2017-06-18 | 2 years ago | Run functions in parallel. |
| [threadpool](https://github.com/shettyh/threadpool) | 25 | 2 | 2017-09-06 | 10 months ago | Golang threadpool implementation. |
| [artifex](https://github.com/mborders/artifex) | 22 | 3 | 2018-10-31 | 2 months ago | Simple in-memory job queue for Golang using worker-based dispatching. |
| [Hunch](https://github.com/AaronJan/Hunch) | 21 | 0 | 2019-06-05 | 5 months ago | Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive. |
| [kyoo](https://github.com/dirkaholic/kyoo) | 17 | 1 | 2020-01-06 | 1 week ago | Provides an unlimited job queue and concurrent worker pools. |
| [stl](https://github.com/ssgreg/stl) | 11 | 1 | 2018-06-19 | 3 months ago | Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism. |
| [gohive](https://github.com/loveleshsharma/gohive) | 10 | 3 | 2019-05-31 | 3 months ago | A highly performant and easy to use Goroutine pool for Go. |
| [routine](https://github.com/x-mod/routine) | 6 | 1 | 2019-03-04 | 1 day ago | go routine control with context, support: Main, Go, Pool and some useful Executors. |
| [go-tools](https://github.com/nikhilsaraf/go-tools) | 5 | 2 | 2018-11-14 | 9 months ago | Manage a pool of goroutines using this lightweight library with a simple API. |
| [go-trylock](https://github.com/subchen/go-trylock) | 5 | 1 | 2018-04-26 | 3 weeks ago | TryLock support on read-write lock for Golang. |
| [conexec](https://github.com/ITcathyh/conexec) | 3 | 2 | 2019-12-24 | 4 days ago | A concurrent toolkit to help execute funcs concurrently in an efficient and safe way.It supports specifying the overall timeout to avoid blocking and uses goroutine pool to improve efficiency. |
| [go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup) | 3 | 1 | 2018-08-08 | 3 months ago | Like `sync.WaitGroup` with error handling and concurrency control. |
| [queue](https://github.com/AnikHasibul/queue) | 3 | 0 | 2018-12-21 | 8 months ago | Gives you a `sync.WaitGroup` like queue group accessibility. Helps you to throttle and limit goroutines, wait for the end of the all goroutines and much more. |

## GUI
        
*Interaction*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [fyne](https://github.com/fyne-io/fyne) | 7587 | 178 | 2018-02-04 | 18 hours ago | Cross platform native GUIs designed for Go based on Material Design. Supports: Linux, macOS, Windows, BSD, iOS and Android. |
| [ui](https://github.com/andlabs/ui) | 7303 | 373 | 2014-02-17 | 3 months ago | Platform-native GUI library for Go. Cross platform. |
| [qt](https://github.com/therecipe/qt) | 6892 | 295 | 2014-11-19 | 22 hours ago | Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi). |
| [webview](https://github.com/zserge/webview) | 5282 | 187 | 2017-08-19 | 3 hours ago | Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux). |
| [robotgo](https://github.com/go-vgo/robotgo) | 4801 | 197 | 2016-09-26 | 1 day ago | Go Native cross-platform GUI system automation. Control the mouse, keyboard and other. |
| [walk](https://github.com/lxn/walk) | 4202 | 249 | 2010-09-16 | 1 month ago | Windows application library kit for Go. |
| [app](https://github.com/maxence-charriere/app) | 3250 | 111 | 2016-10-12 | 2 months ago | Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. |
| [go-astilectron](https://github.com/asticode/go-astilectron) | 2976 | 133 | 2017-04-22 | 2 weeks ago | Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron). |
| [go-sciter](https://github.com/sciter-sdk/go-sciter) | 1616 | 123 | 2015-10-15 | 1 week ago | Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform. |
| [systray](https://github.com/getlantern/systray) | 952 | 47 | 2014-11-12 | 1 week ago | Cross platform Go library to place an icon and menu in the notification area. |
| [gotk3](https://github.com/gotk3/gotk3) | 892 | 50 | 2015-08-13 | 6 days ago | Go bindings for GTK3. |
| [gosx-notifier](https://github.com/deckarep/gosx-notifier) | 503 | 14 | 2013-11-25 | 8 months ago | OSX Desktop Notifications library for Go. |
| [gowd](https://github.com/dtylman/gowd) | 235 | 25 | 2017-03-29 | 7 months ago | Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform. |
| [trayhost](https://github.com/shurcooL/trayhost) | 168 | 4 | 2014-04-25 | 2 months ago | Cross-platform Go library to place an icon in the host operating system's taskbar. |
| [go-appindicator](https://github.com/dawidd6/go-appindicator) | 5 | 2 | 2019-05-04 | 4 months ago | Go bindings for libappindicator3 C library. |
| [activity-tracker](https://github.com/prashantgupta24/activity-tracker) | 2 | 1 | 2019-03-12 | 7 months ago | OSX library to notify about any (pluggable) activity on your machine. |
| [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) | 1 | 1 | 2019-03-30 | 7 months ago | OSX Sleep/Wake notifications in golang. |

## Hardware
        
*Libraries, tools, and tutorials for interacting with hardware.*

## Images
        
*Libraries for manipulating images.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gocv](https://github.com/hybridgroup/gocv) | 2887 | 132 | 2017-09-18 | 1 day ago | Go package for computer vision using OpenCV 3.3+. |
| [imaging](https://github.com/disintegration/imaging) | 2886 | 74 | 2012-12-06 | 2 months ago | Simple Go image processing package. |
| [imaginary](https://github.com/h2non/imaginary) | 2805 | 74 | 2015-03-04 | 1 month ago | Fast and simple HTTP microservice for image resizing. |
| [bild](https://github.com/anthonynsimon/bild) | 2749 | 59 | 2016-08-01 | 1 week ago | Collection of image processing algorithms in pure Go. |
| [ln](https://github.com/fogleman/ln) | 2588 | 88 | 2016-01-10 | 6 months ago | 3D line art rendering in Go. |
| [resize](https://github.com/nfnt/resize) | 2253 | 81 | 2012-08-02 | 1 year ago | Image resizing for Go with common interpolation methods. |
| [gg](https://github.com/fogleman/gg) | 2104 | 73 | 2016-02-18 | 4 months ago | 2D rendering in pure Go. |
| [pt](https://github.com/fogleman/pt) | 1824 | 57 | 2015-01-23 | 10 months ago | Path tracing engine written in Go. |
| [svgo](https://github.com/ajstarks/svgo) | 1414 | 50 | 2010-03-05 | 1 month ago | Go Language Library for SVG generation. |
| [smartcrop](https://github.com/muesli/smartcrop) | 1309 | 31 | 2014-04-07 | 3 months ago | Finds good crops for arbitrary images and crop sizes. |
| [gift](https://github.com/disintegration/gift) | 1288 | 50 | 2014-07-12 | 5 months ago | Package of image processing filters. |
| [picfit](https://github.com/thoas/picfit) | 1171 | 48 | 2014-12-06 | 4 weeks ago | An image resizing server written in Go. |
| [go-opencv](https://github.com/go-opencv/go-opencv) | 1142 | 70 | 2013-12-09 | 8 months ago | Go bindings for OpenCV. |
| [imagick](https://github.com/gographics/imagick) | 1094 | 56 | 2013-04-30 | 17 hours ago | Go binding to ImageMagick's MagickWand C API. |
| [geopattern](https://github.com/pravj/geopattern) | 1041 | 21 | 2014-10-22 | 1 year ago | Create beautiful generative image patterns from a string. |
| [bimg](https://github.com/h2non/bimg) | 869 | 35 | 2015-03-17 | 5 days ago | Small package for fast and efficient image processing using libvips. |
| [stegify](https://github.com/DimitarPetrov/stegify) | 592 | 16 | 2018-11-29 | 3 weeks ago | Go tool for LSB steganography, capable of hiding any file within an image. |
| [canvas](https://github.com/tdewolff/canvas) | 394 | 13 | 2017-05-20 | 3 days ago | Vector graphics to PDF, SVG or rasterized image. |
| [mort](https://github.com/aldor007/mort) | 381 | 19 | 2017-11-19 | 3 months ago | Storage and image processing server written in Go. |
| [image2ascii](https://github.com/qeesung/image2ascii) | 358 | 7 | 2018-10-20 | 1 year ago | Convert image to ASCII. |
| [govatar](https://github.com/o1egl/govatar) | 334 | 6 | 2016-01-18 | 1 week ago | Library and CMD tool for generating funny avatars. |
| [go-nude](https://github.com/koyachi/go-nude) | 296 | 15 | 2014-05-02 | 1 year ago | Nudity detection with Go. |
| [goimagehash](https://github.com/corona10/goimagehash) | 261 | 9 | 2017-07-28 | 1 month ago | Go Perceptual image hashing package. |
| [rez](https://github.com/bamiaux/rez) | 195 | 9 | 2014-01-16 | 2 years ago | Image resizing in pure Go and SIMD. |
| [img](https://github.com/hawx/img) | 132 | 5 | 2012-07-28 | 4 years ago | Selection of image manipulation tools. |
| [darkroom](https://github.com/gojek/darkroom) | 110 | 6 | 2019-07-01 | 2 days ago | An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency. |
| [go-cairo](https://github.com/ungerik/go-cairo) | 90 | 6 | 2012-08-22 | 3 months ago | Go binding for the cairo graphics library. |
| [mergi](https://github.com/noelyahan/mergi) | 85 | 3 | 2018-09-24 | 8 months ago | Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate). |
| [gltf](https://github.com/qmuntal/gltf) | 57 | 1 | 2019-01-15 | 1 month ago | Efficient and robust glTF 2.0 reader, writer and validator. |
| [go-gd](https://github.com/bolknote/go-gd) | 53 | 4 | 2011-05-12 | 1 year ago | Go binding for GD library. |
| [cameron](https://github.com/aofei/cameron) | 36 | 1 | 2018-05-05 | 4 months ago | An avatar generator for Go. |
| [steganography](https://github.com/auyer/steganography) | 36 | 3 | 2018-05-21 | 2 months ago | Pure Go Library for LSB steganography. |
| [goimghdr](https://github.com/corona10/goimghdr) | 31 | 1 | 2018-02-25 | 7 months ago | The imghdr module determines the type of image contained in a file for Go. |
| [tga](https://github.com/ftrvxmtrx/tga) | 25 | 3 | 2012-10-08 | 4 years ago | Package tga is a TARGA image format decoder/encoder. |
| [go-webcolors](https://github.com/jyotiska/go-webcolors) | 24 | 2 | 2014-04-24 | 4 years ago | Port of webcolors library from Python to Go. |
| [mpo](https://github.com/donatj/mpo) | 6 | 1 | 2015-04-14 | 1 year ago | Decoder and conversion tool for MPO 3D Photos. |

## IoT
        
*Libraries for programming devices of the IoT.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [flogo](https://github.com/TIBCOSoftware/flogo) | 1321 | 131 | 2016-07-10 | 1 month ago | Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. |
| [gatt](https://github.com/paypal/gatt) | 867 | 54 | 2014-04-23 | 1 month ago | Gatt is a Go package for building Bluetooth Low Energy peripherals. |
| [mainflux](https://github.com/mainflux/mainflux) | 763 | 78 | 2015-07-06 | 4 days ago | Industrial IoT Messaging and Device Management Server. |
| [devices](https://github.com/goiot/devices) | 229 | 17 | 2016-05-30 | 3 years ago | Suite of libraries for IoT devices, experimental for x/exp/io. |
| [heedy](https://github.com/heedy/heedy) | 193 | 20 | 2015-01-16 | 2 weeks ago | Open-Source Platform for Quantified Self & IoT. |
| [sensorbee](https://github.com/sensorbee/sensorbee) | 187 | 17 | 2016-02-19 | 2 months ago | Lightweight stream processing engine for IoT. |
| [huego](https://github.com/amimof/huego) | 130 | 3 | 2017-05-16 | 4 weeks ago | An extensive Philips Hue client library for Go. |
| [eywa](https://github.com/xcodersun/eywa) | 43 | 8 | 2016-02-20 | 2 years ago | Project Eywa is essentially a connection manager that keeps track of connected devices. |

## Job Scheduler
        
*Libraries for scheduling jobs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gron](https://github.com/roylee0704/gron) | 679 | 16 | 2016-06-04 | 11 months ago | Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly. |
| [jobrunner](https://github.com/bamzi/jobrunner) | 642 | 19 | 2015-10-21 | 1 month ago | Smart and featureful cron job scheduler with job queuing and live monitoring built in. |
| [jobs](https://github.com/albrow/jobs) | 461 | 18 | 2015-02-09 | 1 year ago | Persistent and flexible background jobs library. |
| [scheduler](https://github.com/carlescere/scheduler) | 313 | 16 | 2015-02-03 | 1 year ago | Cronjobs scheduling made easy. |
| [go-cron](https://github.com/rk/go-cron) | 178 | 9 | 2011-04-15 | 4 years ago | Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. |
| [clockwork](https://github.com/whiteShtef/clockwork) | 95 | 4 | 2018-04-23 | 2 months ago | Simple and intuitive job scheduling library in Go. |
| [leprechaun](https://github.com/kilgaloon/leprechaun) | 50 | 5 | 2018-04-08 | 1 month ago | Job scheduler that supports webhooks, crons and classic scheduling. |

## JSON
        
*Libraries for working with JSON.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gjson](https://github.com/tidwall/gjson) | 5762 | 134 | 2016-08-11 | 21 hours ago | Get a JSON value with one line of code. |
| [gojson](https://github.com/ChimeraCoder/gojson) | 2131 | 45 | 2012-12-27 | 2 months ago | Automatically generate Go (golang) struct definitions from example JSON. |
| [kazaam](https://github.com/qntfy/kazaam) | 143 | 20 | 2016-07-19 | 4 months ago | API for arbitrary transformation of JSON documents. |
| [gojq](https://github.com/elgs/gojq) | 141 | 4 | 2015-12-30 | 1 year ago | JSON query in Golang. |
| [jsongo](https://github.com/ricardolonga/jsongo) | 94 | 1 | 2015-08-07 | 3 years ago | Fluent API to make it easier to create Json objects. |
| [gjo](https://github.com/skanehira/gjo) | 71 | 7 | 2019-02-23 | 1 month ago | Small utility to create JSON objects. |
| [jaydiff](https://github.com/yazgazan/jaydiff) | 61 | 1 | 2017-04-24 | 3 months ago | JSON diff utility written in Go. |
| [jettison](https://github.com/wI2L/jettison) | 60 | 3 | 2019-08-30 | 4 days ago | High performance, reflection-less JSON encoder for Go. |
| [jsonf](https://github.com/miolini/jsonf) | 54 | 3 | 2015-05-25 | 3 years ago | Console tool for highlighted formatting and struct query fetching JSON. |
| [mp](https://github.com/sanbornm/mp) | 37 | 2 | 2014-06-15 | 3 years ago | Simple cli email parser. It currently takes stdin and outputs JSON. |
| [go-respond](https://github.com/nicklaw5/go-respond) | 29 | 1 | 2017-03-12 | 6 months ago | Go package for handling common HTTP JSON responses. |
| [json2go](https://github.com/m-zajac/json2go) | 27 | 2 | 2017-06-10 | 4 months ago | Advanced JSON to Go struct conversion. Provides package that can parse multiple JSON documents and create struct to fit them all. |
| [ajson](https://github.com/spyzhov/ajson) | 19 | 0 | 2019-03-07 | 7 months ago | Abstract JSON for golang with JSONPath support. |
| [go-jsonerror](https://github.com/ddymko/go-jsonerror) | 9 | 1 | 2018-10-18 | 3 months ago | Go-JsonError is ment to allow us to easily create json response errors that follow the JsonApi spec. |
| [jsonhal](https://github.com/RichardKnop/jsonhal) | 9 | 2 | 2016-01-15 | 1 year ago | Simple Go package to make custom structs marshal into HAL compatible JSON responses. |
| [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) | 7 | 1 | 2016-07-08 | 3 years ago | Go bindings based on the JSON API errors reference. |
| [ej](https://github.com/lucassscaravelli/ej) | 3 | 1 | 2020-01-04 | 1 week ago | Write and read JSON from different sources succinctly. |

## Logging
        
*Libraries for generating and working with log files.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [logrus](https://github.com/sirupsen/logrus) | 13693 | 302 | 2013-10-16 | 1 day ago | Structured logger for Go. |
| [zap](https://github.com/uber-go/zap) | 8844 | 219 | 2016-02-18 | 1 week ago | Fast, structured, leveled logging in Go. |
| [go-spew](https://github.com/davecgh/go-spew) | 3672 | 59 | 2013-01-09 | 6 days ago | Implements a deep pretty printer for Go data structures to aid in debugging. |
| [zerolog](https://github.com/rs/zerolog) | 2866 | 50 | 2017-05-12 | 5 days ago | Zero-allocation JSON logger. |
| [glog](https://github.com/golang/glog) | 2460 | 88 | 2013-07-16 | 1 month ago | Leveled execution logs for Go. |
| [lumberjack](https://github.com/natefinch/lumberjack) | 1743 | 50 | 2014-06-14 | 2 weeks ago | Simple rolling logger, implements io.WriteCloser. |
| [tail](https://github.com/hpcloud/tail) | 1689 | 98 | 2013-02-05 | 1 month ago | Go package striving to emulate the features of the BSD tail program. |
| [seelog](https://github.com/cihub/seelog) | 1421 | 89 | 2011-11-17 | 10 months ago | Logging functionality with flexible dispatching, filtering, and formatting. |
| [log15](https://github.com/inconshreveable/log15) | 944 | 26 | 2014-05-20 | 1 week ago | Simple, powerful logging for Go. |
| [log](https://github.com/apex/log) | 777 | 34 | 2015-12-21 | 4 months ago | Structured logging package for Go. |
| [onelog](https://github.com/francoispqt/onelog) | 359 | 9 | 2018-05-06 | 10 months ago | Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenario. Also, it is one of the logger with the lowest allocation. |
| [logxi](https://github.com/mgutz/logxi) | 340 | 10 | 2015-03-01 | 7 months ago | 12-factor app logger that is fast and makes you happy. |
| [log](https://github.com/go-playground/log) | 270 | 10 | 2016-02-07 | 2 months ago | Simple, configurable and scalable Structured Logging for Go. |
| [logutils](https://github.com/hashicorp/logutils) | 265 | 169 | 2013-10-09 | 1 month ago | Utilities for slightly better logging in Go (Golang) extending the standard logger. |
| [go-logger](https://github.com/apsdehal/go-logger) | 246 | 8 | 2014-09-26 | 8 months ago | Simple logger of Go Programs, with level handlers. |
| [logger](https://github.com/azer/logger) | 137 | 5 | 2014-09-30 | 6 months ago | Minimalistic logging library for Go. |
| [xlog](https://github.com/rs/xlog) | 131 | 8 | 2015-10-22 | 2 months ago | Structured logger for `net/context` aware HTTP handlers with flexible dispatching. |
| [rollingwriter](https://github.com/arthurkiller/rollingwriter) | 128 | 7 | 2017-02-12 | 1 month ago | RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation. |
| [ozzo-log](https://github.com/go-ozzo/ozzo-log) | 110 | 9 | 2015-10-22 | 1 year ago | High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail). |
| [logvoyage](https://github.com/firstrow/logvoyage) | 84 | 5 | 2015-03-29 | 2 years ago | Full-featured logging saas written in golang. |
| [glg](https://github.com/kpango/glg) | 62 | 5 | 2017-06-21 | 6 hours ago | glg is simple and fast leveled logging library for Go. |
| [log](https://github.com/alexcesaro/log) | 42 | 5 | 2014-04-19 | 4 years ago | Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs. |
| [gologger](https://github.com/sadlil/gologger) | 38 | 5 | 2015-09-02 | 2 years ago | Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch. |
| [go-log](https://github.com/ian-kent/go-log) | 34 | 2 | 2014-05-02 | 1 year ago | Log4j implementation in Go. |
| [logex](https://github.com/chzyer/logex) | 34 | 7 | 2014-10-10 | 2 years ago | Golang log lib, supports tracking and level, wrap by standard log lib. |
| [go-log](https://github.com/siddontang/go-log) | 26 | 5 | 2014-05-18 | 11 months ago | Log lib supports level and multi handlers. |
| [logrusly](https://github.com/sebest/logrusly) | 25 | 5 | 2014-09-11 | 1 year ago | [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/). |
| [go-cronowriter](https://github.com/utahta/go-cronowriter) | 24 | 1 | 2017-02-04 | 4 months ago | Simple writer that rotate log files automatically based on current date and time, like cronolog. |
| [log](https://github.com/teris-io/log) | 23 | 1 | 2017-10-28 | 2 years ago | Structured log interface for Go cleanly separates logging facade from its implementation. |
| [journald](https://github.com/ssgreg/journald) | 22 | 2 | 2017-08-23 | 1 year ago | Go implementation of systemd Journal's native API for logging. |
| [distillog](https://github.com/amoghe/distillog) | 21 | 1 | 2015-10-12 | 1 year ago | distilled levelled logging (think of it as stdlib + log levels). |
| [mlog](https://github.com/jbrodriguez/mlog) | 19 | 1 | 2014-10-20 | 1 year ago | Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output. |
| [gomol](https://github.com/aphistic/gomol) | 15 | 2 | 2015-08-30 | 10 months ago | Multiple-output, structured logging for Go with extensible logging outputs. |
| [go-log](https://github.com/subchen/go-log) | 9 | 2 | 2017-05-07 | 1 year ago | Simple and configurable Logging in Go, with level, formatters and writers. |
| [logdump](https://github.com/ewwwwwqm/logdump) | 9 | 2 | 2017-01-13 | 1 year ago | Package for multi-level logging. |
| [glo](https://github.com/lajosbencz/glo) | 8 | 1 | 2019-01-19 | 1 year ago | PHP Monolog inspired logging facility with identical severity levels. |
| [logrusiowriter](https://github.com/cabify/logrusiowriter) | 7 | 80 | 2019-08-09 | 1 week ago | `io.Writer` implementation using [logrus](https://github.com/sirupsen/logrus) logger. |
| [log](https://github.com/aerogo/log) | 6 | 1 | 2017-06-10 | 2 months ago | An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection). |
| [logmatic](https://github.com/borderstech/logmatic) | 6 | 1 | 2018-11-07 | 1 year ago | Colorized logger for Golang with dynamic log level configuration. |
| [xlog](https://github.com/xfxdev/xlog) | 6 | 1 | 2016-05-05 | 1 year ago | Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format. |
| [logo](https://github.com/mbndr/logo) | 5 | 1 | 2017-02-07 | 2 years ago | Golang logger to different configurable writers. |
| [go-log](https://github.com/pieterclaerhout/go-log) | 1 | 1 | 2019-10-01 | 1 month ago | A logging library with strack traces, object dumping and optional timestamps. |

## Machine Learning
        
*Libraries for Machine Learning.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [golearn](https://github.com/sjwhitworth/golearn) | 6940 | 430 | 2013-12-26 | 3 weeks ago | General Machine Learning library for Go. |
| [gorgonia](https://github.com/gorgonia/gorgonia) | 3087 | 175 | 2016-09-14 | 1 day ago | graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. |
| [tfgo](https://github.com/galeone/tfgo) | 1287 | 47 | 2017-05-23 | 1 month ago | Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python. |
| [goml](https://github.com/cdipaolo/goml) | 1055 | 74 | 2015-06-27 | 9 months ago | On-line Machine Learning in Go. |
| [gosseract](https://github.com/otiai10/gosseract) | 1035 | 36 | 2013-10-11 | 2 weeks ago | Go package for OCR (Optical Character Recognition), by using Tesseract C++ library. |
| [gorse](https://github.com/zhenghaoz/gorse) | 890 | 40 | 2018-08-14 | 3 weeks ago | An offline recommender system backend based on collaborative filtering written in Go. |
| [CloudForest](https://github.com/ryanbressler/CloudForest) | 662 | 44 | 2012-10-22 | 1 year ago | Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go. |
| [eaopt](https://github.com/MaxHalford/eaopt) | 655 | 29 | 2016-01-31 | 1 month ago | An evolutionary optimization library. |
| [bayesian](https://github.com/jbrukh/bayesian) | 649 | 32 | 2011-11-23 | 10 months ago | Naive Bayesian Classification for Golang. |
| [gobrain](https://github.com/goml/gobrain) | 422 | 26 | 2014-04-29 | 2 months ago | Neural Networks written in go. |
| [ocrserver](https://github.com/otiai10/ocrserver) | 266 | 11 | 2015-11-15 | 10 months ago | A simple OCR API server, seriously easy to be deployed by Docker and Heroku. |
| [regommend](https://github.com/muesli/regommend) | 263 | 15 | 2014-02-05 | 5 months ago | Recommendation & collaborative filtering engine. |
| [go-deep](https://github.com/patrikeh/go-deep) | 251 | 13 | 2017-12-09 | 1 month ago | A feature-rich neural network library in Go. |
| [onnx-go](https://github.com/owulveryck/onnx-go) | 237 | 9 | 2018-08-28 | 3 weeks ago | Go Interface to Open Neural Network Exchange (ONNX). |
| [go-galib](https://github.com/thoj/go-galib) | 176 | 13 | 2009-11-30 | 4 years ago | Genetic Algorithms library written in Go / golang. |
| [goRecommend](https://github.com/timkaye11/goRecommend) | 152 | 9 | 2014-07-16 | 5 years ago | Recommendation Algorithms library written in Go. |
| [shield](https://github.com/eaigner/shield) | 129 | 10 | 2013-04-10 | 3 years ago | Bayesian text classifier with flexible tokenizers and storage backends for Go. |
| [goptuna](https://github.com/c-bata/goptuna) | 107 | 5 | 2019-07-24 | 2 hours ago | Bayesian optimization framework for black-box functions written in Go. Everything will be optimized. |
| [go-fann](https://github.com/vksnk/go-fann) | 103 | 7 | 2011-03-10 | 5 years ago | Go bindings for Fast Artificial Neural Networks(FANN) library. |
| [goga](https://github.com/tomcraven/goga) | 83 | 7 | 2015-10-20 | 3 years ago | Genetic algorithm library for Go. |
| [libsvm](https://github.com/datastream/libsvm) | 64 | 10 | 2012-07-31 | 3 years ago | libsvm golang version derived work based on LIBSVM 3.14. |
| [neural-go](https://github.com/schuyler/neural-go) | 61 | 2 | 2011-10-17 | 5 years ago | Multilayer perceptron network implemented in Go, with training via backpropagation. |
| [go-pr](https://github.com/daviddengcn/go-pr) | 57 | 6 | 2013-06-07 | 6 years ago | Pattern recognition package in Go lang. |
| [neat](https://github.com/jinyeom/neat) | 56 | 12 | 2016-11-17 | 1 year ago | Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT). |
| [gonet](https://github.com/dathoangnd/gonet) | 50 | 4 | 2020-01-11 | 1 week ago | Neural Network for Go. |
| [goscore](https://github.com/asafschers/goscore) | 43 | 3 | 2017-08-19 | 5 months ago | Go Scoring API for PMML. |
| [golinear](https://github.com/danieldk/golinear) | 39 | 5 | 2013-04-05 | 1 year ago | liblinear bindings for Go. |
| [fonet](https://github.com/Fontinalis/fonet) | 35 | 4 | 2017-10-03 | 4 months ago | A Deep Neural Network library written in Go. |
| [Varis](https://github.com/Xamber/Varis) | 26 | 5 | 2017-10-10 | 1 year ago | Golang Neural Network. |
| [godist](https://github.com/e-dard/godist) | 25 | 2 | 2014-09-05 | 4 years ago | Various probability distributions, and associated methods. |
| [go-cluster](https://github.com/e-XpertSolutions/go-cluster) | 21 | 5 | 2017-10-04 | 1 year ago | Go implementation of the k-modes and k-prototypes clustering algorithms. |
| [probab](https://github.com/ThePaw/probab) | 11 | 1 | 2015-09-14 | 4 years ago | Probability distribution functions. Bayesian inference. Written in pure Go. |
| [evoli](https://github.com/khezen/evoli) | 10 | 3 | 2015-06-12 | 1 month ago | Genetic Algorithm and Particle Swarm Optimization library. |
| [gomind](https://github.com/surenderthakran/gomind) | 7 | 2 | 2017-10-19 | 1 year ago | A simplistic Neural Network Library in Go. |
| [randomForest](https://github.com/malaschitz/randomForest) | 2 | 1 | 2018-10-25 | 1 month ago | Easy to use Random Forest library for Go. |

## Messaging
        
*Libraries that implement messaging systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [sarama](https://github.com/Shopify/sarama) | 5319 | 394 | 2013-07-05 | 1 hour ago | Go library for Apache Kafka. |
| [gorush](https://github.com/appleboy/gorush) | 4094 | 181 | 2016-03-22 | 1 day ago | Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm). |
| [centrifugo](https://github.com/centrifugal/centrifugo) | 4017 | 181 | 2015-03-31 | 19 hours ago | Real-time messaging (Websockets or SockJS) server in Go. |
| [machinery](https://github.com/RichardKnop/machinery) | 3783 | 137 | 2015-04-05 | 2 weeks ago | Asynchronous task queue/job queue based on distributed message passing. |
| [go-socket.io](https://github.com/googollee/go-socket.io) | 3210 | 129 | 2013-07-13 | 4 days ago | socket.io library for golang, a realtime application framework. |
| [nats.go](https://github.com/nats-io/nats.go) | 2632 | 157 | 2012-08-15 | 6 days ago | Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library. |
| [apns2](https://github.com/sideshow/apns2) | 2232 | 73 | 2016-01-05 | 9 hours ago | HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps. |
| [benthos](https://github.com/Jeffail/benthos) | 2193 | 63 | 2016-03-22 | 2 hours ago | A message streaming bridge between a range of protocols. |
| [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) | 1886 | 237 | 2013-12-27 | 2 years ago | gopush-cluster is a go push server cluster. |
| [mercure](https://github.com/dunglas/mercure) | 1819 | 54 | 2018-07-14 | 5 days ago | Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events). |
| [melody](https://github.com/olahol/melody) | 1713 | 54 | 2015-05-13 | 1 week ago | Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling. |
| [go-nsq](https://github.com/nsqio/go-nsq) | 1572 | 61 | 2013-08-29 | 5 days ago | the official Go package for NSQ. |
| [mangos-v1](https://github.com/nanomsg/mangos-v1) | 1545 | 84 | 2014-10-25 | 2 months ago | Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability. |
| [uniqush-push](https://github.com/uniqush/uniqush-push) | 1129 | 75 | 2011-08-29 | 1 month ago | Redis backed unified push service for server-side notifications to mobile devices. |
| [gollum](https://github.com/trivago/gollum) | 825 | 37 | 2015-06-20 | 2 months ago | A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations. |
| [zmq4](https://github.com/pebbe/zmq4) | 822 | 48 | 2013-10-18 | 2 weeks ago | Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2). |
| [Beaver](https://github.com/Clivern/Beaver) | 778 | 17 | 2018-10-20 | 23 hours ago | A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps. |
| [EventBus](https://github.com/asaskevich/EventBus) | 637 | 28 | 2014-12-19 | 1 week ago | The lightweight event bus with async compatibility. |
| [golongpoll](https://github.com/jcuga/golongpoll) | 439 | 23 | 2015-11-02 | 11 months ago | HTTP longpoll server library that makes web pub-sub simple. |
| [dbus](https://github.com/godbus/dbus) | 406 | 16 | 2014-03-27 | 5 days ago | Native Go bindings for D-Bus. |
| [emitter](https://github.com/olebedev/emitter) | 331 | 9 | 2015-11-10 | 1 year ago | Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins. |
| [glue](https://github.com/desertbit/glue) | 325 | 14 | 2015-06-07 | 6 months ago | Robust Go and Javascript Socket Library (Alternative to Socket.io). |
| [pubsub](https://github.com/cskr/pubsub) | 306 | 8 | 2012-04-01 | 9 months ago | Simple pubsub package for go. |
| [guble](https://github.com/smancke/guble) | 142 | 13 | 2015-11-15 | 2 years ago | Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence. |
| [bus](https://github.com/mustafaturan/bus) | 133 | 5 | 2019-04-27 | 2 months ago | Minimalist message bus implementation for internal communication. |
| [oplog](https://github.com/dailymotion/oplog) | 100 | 92 | 2014-11-06 | 4 years ago | Generic oplog/replication system for REST APIs. |
| [rabtap](https://github.com/jandelgado/rabtap) | 99 | 8 | 2017-11-11 | 3 days ago | RabbitMQ swiss army knife cli app. |
| [message-bus](https://github.com/vardius/message-bus) | 79 | 4 | 2017-10-04 | 4 months ago | messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD. |
| [rabbus](https://github.com/rafaeljesus/rabbus) | 68 | 9 | 2017-05-07 | 6 months ago | A tiny wrapper over amqp exchanges and queues. |
| [drone-line](https://github.com/appleboy/drone-line) | 65 | 5 | 2016-09-13 | 2 months ago | Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI. |
| [RapidMQ](https://github.com/sybrexsys/RapidMQ) | 58 | 4 | 2016-10-04 | 2 years ago | RapidMQ is a lightweight and reliable library for managing of the local messages queue. |
| [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) | 57 | 9 | 2017-01-15 | 1 year ago | A tiny wrapper around NSQ topic and channel. |
| [go-notify](https://github.com/TheCreeper/go-notify) | 47 | 2 | 2015-03-01 | 11 months ago | Native implementation of the freedesktop notification spec. |
| [hub](https://github.com/leandro-lugaresi/hub) | 39 | 1 | 2018-04-13 | 1 week ago | A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges. |
| [commander](https://github.com/CloudProud/commander) | 38 | 1 | 2018-04-20 | 1 month ago | A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka. |
| [commander](https://github.com/jeroenrinzema/commander) | 38 | 1 | 2018-04-20 | 1 month ago | A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka. |
| [event](https://github.com/agoalofalife/event) | 32 | 3 | 2017-07-02 | 1 year ago | Implementation of the pattern observer. |
| [go-vitotrol](https://github.com/maxatome/go-vitotrol) | 12 | 6 | 2016-11-03 | 8 months ago | Client library to Viessmann Vitotrol web service. |
| [redisqueue](https://github.com/robinjoseph08/redisqueue) | 9 | 2 | 2019-07-07 | 5 months ago | redisqueue provides a producer and consumer of a queue that uses Redis streams. |
| [gaurun-client](https://github.com/osamingo/gaurun-client) | 8 | 1 | 2017-06-29 | 8 months ago | Gaurun Client written in Go. |
| [jazz](https://github.com/socifi/jazz) | 8 | 3 | 2018-10-22 | 10 months ago | A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages. |
| [rmqconn](https://github.com/sbabiv/rmqconn) | 2 | 0 | 2019-01-14 | 7 months ago | RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed. |
| [ami](https://github.com/kak-tus/ami) | 2 | 1 | 2018-10-27 | 1 week ago | Go client to reliable queues based on Redis Cluster Streams. |

## Microsoft Office
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [unioffice](https://github.com/unidoc/unioffice) | 2041 | 58 | 2017-08-29 | 2 hours ago | Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents. |

### Microsoft Excel
        
*Libraries for working with Microsoft Excel.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [excelize](https://github.com/360EntSecGroup-Skylar/excelize) | 5314 | 161 | 2016-08-29 | 1 hour ago | Golang library for reading and writing Microsoft Excel™ (XLSX) files. |
| [xlsx](https://github.com/tealeg/xlsx) | 3846 | 183 | 2011-06-28 | 1 week ago | Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs. |
| [xlsx](https://github.com/plandem/xlsx) | 103 | 11 | 2017-08-26 | 1 month ago | Fast and safe way to read/update your existing Microsoft Excel files in Go programs. |
| [go-excel](https://github.com/szyhf/go-excel) | 60 | 3 | 2017-09-03 | 2 weeks ago | A simple and light reader to read a relate-db-like excel as a table. |
| [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) | 14 | 1 | 2017-03-13 | 1 year ago | Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files. |

## Miscellaneous
        

### Dependency Injection
        
*Libraries for working with dependency injection.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [dig](https://github.com/uber-go/dig) | 1211 | 37 | 2017-03-21 | 2 weeks ago | A reflection based dependency injection toolkit for Go. |
| [fx](https://github.com/uber-go/fx) | 1119 | 52 | 2016-10-27 | 2 weeks ago | A dependency injection based application framework for Go (built on top of dig). |
| [container](https://github.com/golobby/container) | 65 | 2 | 2019-09-23 | 3 weeks ago | A powerful IoC Container with fluent and easy-to-use interface. |
| [inject](https://github.com/defval/inject) | 49 | 1 | 2019-02-03 | 3 weeks ago | A reflection based dependency injection container with simple interface. |
| [alice](https://github.com/magic003/alice) | 36 | 2 | 2017-04-08 | 2 years ago | Additive dependency injection container for Golang. |
| [dingo](https://github.com/i-love-flamingo/dingo) | 36 | 21 | 2018-10-29 | 1 month ago | A dependency injection toolkit for Go, based on Guice. |
| [wire](https://github.com/Fs02/wire) | 23 | 1 | 2018-07-05 | 6 months ago | Strict Runtime Dependency Injection for Golang. |
| [linker](https://github.com/logrange/linker) | 18 | 2 | 2018-12-04 | 4 days ago | A reflection based dependency injection and inversion of control library with components lifecycle support. |
| [gocontainer](https://github.com/vardius/gocontainer) | 11 | 0 | 2019-06-06 | 5 months ago | Simple Dependency Injection Container. |

### Project Layout
        
*Unofficial set of patterns for structuring projects.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [project-layout](https://github.com/golang-standards/project-layout) | 12660 | 357 | 2017-09-09 | 1 month ago | Set of common historical and emerging project layout patterns in the Go ecosystem. |
| [go-restful-api](https://github.com/qiangxue/go-restful-api) | 1020 | 52 | 2016-08-15 | 2 days ago | An idiomatic Go RESTful API starter kit following SOLID principles and Clean Architecture with a common project layout. |
| [modern-go-application](https://github.com/sagikazarmark/modern-go-application) | 491 | 13 | 2018-09-14 | 4 days ago | Go application boilerplate and example applying modern practices. |
| [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) | 312 | 7 | 2016-12-18 | 1 month ago | A Go application boilerplate template for quick starting projects following production best practices. |
| [scaffold](https://github.com/catchplay/scaffold) | 50 | 2 | 2018-12-11 | 1 year ago | Scaffold generates starter Go project layout. Lets you focus on business logic implemeted. |
| [go-sample](https://github.com/zitryss/go-sample) | 41 | 1 | 2019-01-24 | 1 year ago | A sample layout for Go application projects with the real code. |

### Strings
        
*These libraries were placed here because none of the other categories seemed to fit.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gopsutil](https://github.com/shirou/gopsutil) | 4466 | 190 | 2014-04-18 | 2 hours ago | Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc). |
| [archiver](https://github.com/mholt/archiver) | 2690 | 46 | 2016-04-08 | 1 month ago | Library and command for making and extracting .zip and .tar.gz archives. |
| [gosms](https://github.com/haxpax/gosms) | 1262 | 56 | 2015-01-23 | 2 years ago | Your own local SMS gateway in Go that can be used to send SMS. |
| [go-resiliency](https://github.com/eapache/go-resiliency) | 957 | 28 | 2014-11-29 | 1 month ago | Resiliency patterns for golang. |
| [gofakeit](https://github.com/brianvoe/gofakeit) | 804 | 7 | 2015-04-24 | 1 week ago | Random data generator written in go. |
| [go-commons-pool](https://github.com/jolestar/go-commons-pool) | 757 | 45 | 2015-12-28 | 2 months ago | Generic object pool for Golang. |
| [base64Captcha](https://github.com/mojocn/base64Captcha) | 745 | 47 | 2017-12-12 | 1 week ago | Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha. |
| [xstrings](https://github.com/huandu/xstrings) | 702 | 25 | 2015-01-06 | 1 week ago | Collection of useful string functions ported from other languages. |
| [shortid](https://github.com/teris-io/shortid) | 509 | 8 | 2016-01-04 | 2 years ago | Distributed generation of super short, unique, non-sequential, URL friendly IDs. |
| [llvm](https://github.com/llir/llvm) | 478 | 28 | 2014-09-19 | 5 days ago | Library for interacting with LLVM IR in pure Go. |
| [health](https://github.com/dimiro1/health) | 376 | 6 | 2016-03-08 | 3 months ago | Easy to use, extensible health check library. |
| [go-conv](https://github.com/cstockton/go-conv) | 348 | 8 | 2016-10-11 | 2 years ago | Package conv provides fast and intuitive conversions across Go types. |
| [banner](https://github.com/dimiro1/banner) | 248 | 4 | 2016-03-25 | 3 months ago | Add beautiful banners into your Go applications. |
| [gountries](https://github.com/pariz/gountries) | 228 | 7 | 2016-01-13 | 2 months ago | Package that exposes country and subdivision data. |
| [antch](https://github.com/antchfx/antch) | 166 | 13 | 2017-09-28 | 1 year ago | A fast, powerful and extensible web crawling & scraping framework. |
| [ffmt](https://github.com/go-ffmt/ffmt) | 148 | 4 | 2015-02-14 | 2 weeks ago | Beautify data display for Humans. |
| [lk](https://github.com/hyperboloide/lk) | 142 | 5 | 2016-07-14 | 1 month ago | A simple licensing library for golang. |
| [stateless](https://github.com/qmuntal/stateless) | 140 | 4 | 2019-09-11 | 2 months ago | A fluent library for creating state machines. |
| [battery](https://github.com/distatus/battery) | 138 | 2 | 2016-03-12 | 2 months ago | Cross-platform, normalized battery information library. |
| [stats](https://github.com/go-playground/stats) | 129 | 3 | 2015-09-14 | 3 years ago | Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... |
| [bitio](https://github.com/icza/bitio) | 111 | 6 | 2016-05-31 | 1 month ago | Highly optimized bit-level Reader and Writer for Go. |
| [ghorg](https://github.com/gabrie30/ghorg) | 100 | 5 | 2018-03-29 | 14 hours ago | Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, and Bitbucket. |
| [healthcheck](https://github.com/etherlabsio/healthcheck) | 95 | 4 | 2017-08-18 | 4 weeks ago | An opinionated and concurrent health-check HTTP handler for RESTful services. |
| [turtle](https://github.com/hackebrot/turtle) | 91 | 1 | 2017-09-08 | 2 months ago | Emojis for Go. |
| [strutil](https://github.com/ozgio/strutil) | 83 | 1 | 2018-08-16 | 3 months ago | String utilities. |
| [gommit](https://github.com/antham/gommit) | 83 | 3 | 2016-08-30 | 2 days ago | Analyze git commit messages to ensure they follow defined patterns. |
| [go-unarr](https://github.com/gen2brain/go-unarr) | 82 | 6 | 2015-11-01 | 11 months ago | Decompression library for RAR, TAR, ZIP and 7z archives. |
| [indigo](https://github.com/osamingo/indigo) | 58 | 1 | 2016-08-31 | 3 weeks ago | Distributed unique ID generator of using Sonyflake and encoded by Base58. |
| [morse](https://github.com/alwindoss/morse) | 55 | 2 | 2018-08-15 | 11 months ago | Library to convert to and from morse code. |
| [captcha](https://github.com/steambap/captcha) | 49 | 5 | 2017-09-12 | 1 week ago | Package captcha provides an easy to use, unopinionated API for captcha generation. |
| [xkg](https://github.com/go-xkg/xkg) | 46 | 1 | 2015-01-05 | 5 years ago | X Keyboard Grabber. |
| [gotoprom](https://github.com/cabify/gotoprom) | 45 | 81 | 2018-10-10 | 1 week ago | Type-safe metrics builder wrapper library for the official Prometheus client. |
| [persian](https://github.com/mavihq/persian) | 37 | 1 | 2017-10-16 | 1 year ago | Some utilities for Persian language in go. |
| [pdfgen](https://github.com/hyperboloide/pdfgen) | 36 | 3 | 2015-11-30 | 1 year ago | HTTP service to generate PDF from Json requests. |
| [browscap_go](https://github.com/digitalcrab/browscap_go) | 31 | 6 | 2014-09-18 | 1 month ago | GoLang Library for [Browser Capabilities Project](http://browscap.org/). |
| [datacounter](https://github.com/miolini/datacounter) | 30 | 1 | 2015-10-14 | 2 months ago | Go counters for readers/writer/http.ResponseWriter. |
| [autoflags](https://github.com/artyom/autoflags) | 25 | 3 | 2014-05-15 | 1 year ago | Go package to automatically define command line flags from struct fields. |
| [xdg](https://github.com/rkoesters/xdg) | 21 | 1 | 2013-12-15 | 1 year ago | FreeDesktop.org (xdg) Specs implemented in Go. |
| [gosh](https://github.com/osamingo/gosh) | 19 | 0 | 2018-05-25 | 3 weeks ago | Provide Go Statistics Handler, Struct, Measure Method. |
| [url-shortener](https://github.com/pantrif/url-shortener) | 18 | 1 | 2018-06-04 | 1 year ago | A modern, powerful, and robust URL shortener microservice with mysql support. |
| [sandid](https://github.com/aofei/sandid) | 14 | 1 | 2018-06-12 | 1 month ago | Every grain of sand on earth has its own ID. |
| [anagent](https://github.com/mudler/anagent) | 11 | 2 | 2017-12-29 | 1 year ago | Minimalistic, pluggable Golang evloop/timer handler with dependency-injection. |
| [avgRating](https://github.com/kirillDanshin/avgRating) | 10 | 1 | 2017-08-05 | 2 years ago | Calculate average score and rating based on Wilson Score Equation. |
| [hostutils](https://github.com/Wing924/hostutils) | 8 | 1 | 2017-09-26 | 1 year ago | A golang library for packing and unpacking FQDNs list. |
| [shellwords](https://github.com/Wing924/shellwords) | 7 | 1 | 2017-09-28 | 2 years ago | A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell. |
| [metrics](https://github.com/pascaldekloe/metrics) | 6 | 1 | 2019-01-29 | 1 week ago | Library for metrics instrumentation and Prometheus exposition. |
| [numa](https://github.com/lrita/numa) | 2 | 1 | 2018-12-10 | 8 months ago | NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code. |

## Natural Language Processing
        
*Libraries for working with human languages.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [prose](https://github.com/jdkato/prose) | 2405 | 58 | 2017-02-17 | 1 month ago | Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. English only. |
| [gse](https://github.com/go-ego/gse) | 1197 | 43 | 2017-06-23 | 1 week ago | Go efficient text segmentation; support english, chinese, japanese and other. |
| [when](https://github.com/olebedev/when) | 1051 | 24 | 2016-12-27 | 2 months ago | Natural EN and RU language date/time parser with pluggable rules. |
| [gojieba](https://github.com/yanyiwu/gojieba) | 955 | 62 | 2015-09-12 | 1 week ago | This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm. |
| [go-pinyin](https://github.com/mozillazg/go-pinyin) | 632 | 32 | 2014-11-09 | 6 days ago | CN Hanzi to Hanyu Pinyin converter. |
| [kagome](https://github.com/ikawaha/kagome) | 470 | 22 | 2014-06-26 | 2 months ago | JP morphological analyzer written in pure Go. |
| [whatlanggo](https://github.com/abadojack/whatlanggo) | 390 | 15 | 2017-02-20 | 10 months ago | Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc). |
| [nlp](https://github.com/shixzie/nlp) | 359 | 24 | 2017-01-25 | 2 years ago | Extract values from strings and fill your structs with nlp. |
| [sentences](https://github.com/neurosnap/sentences) | 264 | 12 | 2015-08-07 | 9 months ago | Sentence tokenizer:  converts text into a list of sentences. |
| [nlp](https://github.com/james-bowman/nlp) | 237 | 22 | 2017-03-15 | 1 month ago | Go Natural Language Processing library supporting LSA (Latent Semantic Analysis). |
| [go-nlp](https://github.com/nuance/go-nlp) | 82 | 7 | 2011-05-02 | 8 years ago | Utilities for working with discrete probability distributions and other tools useful for doing NLP work. |
| [getlang](https://github.com/rylans/getlang) | 80 | 3 | 2018-03-01 | 7 months ago | Fast natural language detection package. |
| [gounidecode](https://github.com/fiam/gounidecode) | 68 | 3 | 2012-05-01 | 4 years ago | Unicode transliterator (also known as unidecode) for Go. |
| [go-unidecode](https://github.com/mozillazg/go-unidecode) | 65 | 2 | 2016-07-08 | 9 months ago | ASCII transliterations of Unicode text. |
| [textcat](https://github.com/pebbe/textcat) | 62 | 6 | 2012-09-21 | 1 year ago | Go package for n-gram based text categorization, with support for utf-8 and raw text. |
| [MMSEGO](https://github.com/awsong/MMSEGO) | 59 | 5 | 2012-04-18 | 7 years ago | This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm. |
| [go-stem](https://github.com/agonopol/go-stem) | 55 | 3 | 2011-09-23 | 1 year ago | Implementation of the porter stemming algorithm. |
| [RAKE.Go](https://github.com/afjoseph/RAKE.Go) | 53 | 6 | 2016-12-17 | 2 months ago | Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE). |
| [stemmer](https://github.com/dchest/stemmer) | 48 | 4 | 2011-03-21 | 3 years ago | Stemmer packages for Go programming language. Includes English and German stemmers. |
| [porter2](https://github.com/zentures/porter2) | 36 | 2 | 2015-01-21 | 4 years ago | Really fast Porter 2 stemmer. |
| [go2vec](https://github.com/danieldk/go2vec) | 33 | 6 | 2015-01-27 | 1 year ago | Reader and utility functions for word2vec embeddings. |
| [petrovich](https://github.com/striker2000/petrovich) | 27 | 1 | 2016-12-26 | 2 months ago | Petrovich is the library which inflects Russian names to given grammatical case. |
| [paicehusk](https://github.com/rookii/paicehusk) | 26 | 3 | 2012-09-29 | 6 years ago | Golang implementation of the Paice/Husk Stemming Algorithm. |
| [snowball](https://github.com/goodsign/snowball) | 25 | 0 | 2012-12-11 | 2 years ago | Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/). |
| [mystem](https://github.com/dveselov/mystem) | 23 | 3 | 2016-08-30 | 3 years ago | CGo bindings to Yandex.Mystem - russian morphology analyzer. |
| [icu](https://github.com/goodsign/icu) | 19 | 0 | 2012-12-11 | 2 years ago | Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1. |
| [golibstemmer](https://github.com/rjohnsondev/golibstemmer) | 17 | 1 | 2012-08-06 | 5 years ago | Go bindings for the snowball libstemmer library including porter 2. |
| [shamoji](https://github.com/osamingo/shamoji) | 11 | 2 | 2017-07-23 | 3 weeks ago | The shamoji is word filtering package written in Go. |
| [libtextcat](https://github.com/goodsign/libtextcat) | 10 | 1 | 2012-12-10 | 7 years ago | Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2. |
| [porter](https://github.com/a2800276/porter) | 8 | 1 | 2013-09-17 | 6 years ago | This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm. |
| [gotokenizer](https://github.com/xujiajun/gotokenizer) | 6 | 1 | 2018-10-11 | 9 months ago | A tokenizer based on the dictionary and Bigram language models for Go. (Now only support chinese segmentation) |
| [go-localize](https://github.com/m1/go-localize) | 3 | 1 | 2019-12-23 | 2 weeks ago | Simple and easy to use i18n (Internationalization and localization) engine - used for translating locale strings. |
| [detectlanguage-go](https://github.com/detectlanguage/detectlanguage-go) | 1 | 0 | 2019-12-14 | 3 weeks ago | Language Detection API Go Client. Supports batch requests, short phrase or single word language detection. |

## Networking
        
*Libraries for working with various layers of the network.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [kcptun](https://github.com/xtaci/kcptun) | 11382 | 597 | 2016-02-26 | 4 hours ago | Extremely simple & fast udp tunnel based on KCP protocol. |
| [fasthttp](https://github.com/valyala/fasthttp) | 11255 | 365 | 2015-10-18 | 7 hours ago | Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http. |
| [dns](https://github.com/miekg/dns) | 4269 | 167 | 2010-08-03 | 1 week ago | Go library for working with DNS. |
| [quic-go](https://github.com/lucas-clemente/quic-go) | 3706 | 165 | 2016-04-06 | 7 hours ago | An implementation of the QUIC protocol in pure Go. |
| [httplab](https://github.com/gchaincl/httplab) | 3507 | 67 | 2017-02-08 | 7 months ago | HTTPLabs let you inspect HTTP requests and forge responses. |
| [gopacket](https://github.com/google/gopacket) | 3212 | 138 | 2015-03-16 | 4 days ago | Go library for packet processing with libpcap bindings. |
| [webrtc](https://github.com/pion/webrtc) | 2989 | 134 | 2018-05-18 | 11 hours ago | A pure Go implementation of the WebRTC API. |
| [kcp-go](https://github.com/xtaci/kcp-go) | 2433 | 141 | 2015-06-16 | 2 weeks ago | KCP - Fast and Reliable ARQ Protocol. |
| [gobgp](https://github.com/osrg/gobgp) | 1798 | 124 | 2014-09-14 | 3 days ago | BGP implemented in the Go Programming Language. |
| [gnet](https://github.com/panjf2000/gnet) | 1647 | 58 | 2019-02-24 | 2 days ago | `gnet` is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go. |
| [ssh](https://github.com/gliderlabs/ssh) | 1279 | 39 | 2016-10-03 | 4 weeks ago | Higher-level API for building SSH servers (wraps crypto/ssh). |
| [fortio](https://github.com/fortio/fortio) | 1120 | 34 | 2017-10-10 | 2 weeks ago | Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC. |
| [water](https://github.com/songgao/water) | 946 | 38 | 2013-03-25 | 5 months ago | Simple TUN/TAP library. |
| [go-getter](https://github.com/hashicorp/go-getter) | 838 | 164 | 2015-10-12 | 3 days ago | Go library for downloading files or directories from various sources using a URL. |
| [sftp](https://github.com/pkg/sftp) | 821 | 49 | 2013-11-05 | 1 week ago | Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt. |
| [gev](https://github.com/Allenxuxu/gev) | 772 | 24 | 2019-09-01 | 1 week ago | gev is a lightweight, fast non-blocking TCP network library based on Reactor mode. |
| [nff-go](https://github.com/intel-go/nff-go) | 758 | 74 | 2017-03-29 | 2 weeks ago | Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF). |
| [grab](https://github.com/cavaliercoder/grab) | 627 | 17 | 2016-01-05 | 2 months ago | Go package for managing file downloads. |
| [mdns](https://github.com/hashicorp/mdns) | 609 | 169 | 2014-01-29 | 4 months ago | Simple mDNS (Multicast DNS) client/server library in Golang. |
| [ftp](https://github.com/jlaffaye/ftp) | 606 | 25 | 2011-05-06 | 3 weeks ago | Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959). |
| [lhttp](https://github.com/fanux/lhttp) | 544 | 55 | 2015-12-29 | 1 year ago | Powerful websocket framework, build your IM server more easily. |
| [gosnmp](https://github.com/soniah/gosnmp) | 483 | 43 | 2012-08-27 | 2 weeks ago | Native Go library for performing SNMP actions. |
| [gotcp](https://github.com/gansidui/gotcp) | 448 | 43 | 2014-04-13 | 2 years ago | Go package for quickly writing tcp applications. |
| [cidranger](https://github.com/yl2chen/cidranger) | 431 | 11 | 2017-08-21 | 1 week ago | Fast IP to CIDR lookup for Go. |
| [peerdiscovery](https://github.com/schollz/peerdiscovery) | 395 | 16 | 2018-04-22 | 3 days ago | Pure Go library for cross-platform local peer discovery using UDP multicast. |
| [gopcap](https://github.com/akrennmair/gopcap) | 373 | 24 | 2009-11-19 | 1 month ago | Go wrapper for libpcap. |
| [go-stun](https://github.com/ccding/go-stun) | 342 | 11 | 2013-08-17 | 1 year ago | Go implementation of the STUN client (RFC 3489 and RFC 5389). |
| [raw](https://github.com/mdlayher/raw) | 338 | 11 | 2015-07-06 | 3 months ago | Package raw enables reading and writing data at the device driver level for a network interface. |
| [stun](https://github.com/gortc/stun) | 333 | 15 | 2016-04-24 | 1 week ago | Go implementation of RFC 5389 STUN protocol. |
| [tcp_server](https://github.com/firstrow/tcp_server) | 310 | 17 | 2014-10-13 | 9 months ago | Go library for building tcp servers faster. |
| [buffstreams](https://github.com/StabbyCutyou/buffstreams) | 236 | 14 | 2015-06-29 | 2 years ago | Streaming protocolbuffer data over TCP made easy. |
| [winrm](https://github.com/masterzen/winrm) | 234 | 14 | 2013-12-30 | 2 months ago | Go WinRM client to remotely execute commands on Windows machines. |
| [arp](https://github.com/mdlayher/arp) | 212 | 10 | 2015-07-06 | 1 month ago | Package arp implements the ARP protocol, as described in RFC 826. |
| [ethernet](https://github.com/mdlayher/ethernet) | 193 | 12 | 2015-07-03 | 7 months ago | Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. |
| [utp](https://github.com/anacrolix/utp) | 153 | 16 | 2015-03-20 | 1 year ago | Go uTP micro transport protocol implementation. |
| [jazigo](https://github.com/udhos/jazigo) | 140 | 12 | 2016-06-07 | 4 months ago | Jazigo is a tool written in Go for retrieving configuration for multiple network devices. |
| [canopus](https://github.com/zubairhamed/canopus) | 138 | 14 | 2015-02-24 | 1 year ago | CoAP Client/Server implementation (RFC 7252). |
| [sslb](https://github.com/eduardonunesp/sslb) | 122 | 8 | 2015-10-18 | 3 months ago | It's a Super Simples Load Balancer, just a little project to achieve some kind of performance. |
| [gmqtt](https://github.com/DrmagicE/gmqtt) | 119 | 9 | 2018-09-16 | 3 months ago | Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1. |
| [gnxi](https://github.com/google/gnxi) | 114 | 22 | 2017-09-26 | 3 months ago | A collection of tools for Network Management that use the gNMI and gNOI protocols. |
| [xtcp](https://github.com/xfxdev/xtcp) | 93 | 12 | 2016-03-31 | 1 month ago | TCP Server Framework with simultaneous full duplex communication,graceful shutdown,custom protocol. |
| [dhcp6](https://github.com/mdlayher/dhcp6) | 64 | 4 | 2015-05-22 | 10 months ago | Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. |
| [ether](https://github.com/songgao/ether) | 64 | 3 | 2014-05-21 | 3 years ago | Cross-platform Go package for sending and receiving ethernet frames. |
| [linkio](https://github.com/ian-kent/linkio) | 47 | 2 | 2014-12-24 | 2 years ago | Network link speed simulation for Reader/Writer interfaces. |
| [packet](https://github.com/aerogo/packet) | 42 | 3 | 2017-10-29 | 2 months ago | Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed. |
| [portproxy](https://github.com/aybabtme/portproxy) | 42 | 0 | 2014-12-13 | 5 years ago | Simple TCP proxy which adds CORS support to API's which don't support it. |
| [graval](https://github.com/koofr/graval) | 25 | 8 | 2014-04-22 | 1 year ago | Experimental FTP server framework. |
| [publicip](https://github.com/polera/publicip) | 18 | 1 | 2016-12-28 | 3 years ago | Package publicip returns your public facing IPv4 address (internet egress). |
| [golibwireshark](https://github.com/sunwxg/golibwireshark) | 17 | 2 | 2015-11-16 | 2 years ago | Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data. |
| [go-powerdns](https://github.com/joeig/go-powerdns) | 10 | 3 | 2018-06-21 | 3 weeks ago | PowerDNS API bindings for Golang. |
| [llb](https://github.com/kirillDanshin/llb) | 10 | 1 | 2016-02-21 | 3 years ago | It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response. |
| [goshark](https://github.com/sunwxg/goshark) | 9 | 1 | 2015-11-01 | 2 years ago | Package goshark use tshark to decode IP packet and create data struct to analyse packet. |
| [tspool](https://github.com/two/tspool) | 6 | 0 | 2018-10-27 | 1 year ago | A TCP Library use worker pool to improve performance and protect your server. |
| [gosocsvr](https://github.com/Rakeki/gosocsvr) | 4 | 2 | 2019-11-12 | 1 month ago | Socket server made simple. |
| [httpproxy](https://github.com/wzshiming/httpproxy) | 1 | 1 | 2018-07-18 | 2 weeks ago | HTTP proxy handler and dialer. |

### HTTP Clients
        
*Libraries for making HTTP requests.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [resty](https://github.com/go-resty/resty) | 2488 | 68 | 2015-08-28 | 6 days ago | Simple HTTP and REST client for Go inspired by Ruby rest-client. |
| [grequests](https://github.com/levigross/grequests) | 1510 | 32 | 2015-06-11 | 2 weeks ago | A Go "clone" of the great and famous Requests library. |
| [heimdall](https://github.com/gojek/heimdall) | 1209 | 46 | 2018-01-19 | 3 days ago | An enchanced http client with retry and hystrix capabilities. |
| [sling](https://github.com/dghubble/sling) | 1085 | 33 | 2015-04-02 | 1 month ago | Sling is a Go HTTP client library for creating and sending API requests. |
| [gentleman](https://github.com/h2non/gentleman) | 721 | 19 | 2016-02-21 | 2 months ago | Full-featured plugin-driven HTTP client library. |
| [pester](https://github.com/sethgrid/pester) | 347 | 7 | 2015-05-20 | 11 months ago | Go HTTP client calls with retries, backoff, and concurrency. |
| [sreq](https://github.com/winterssy/sreq) | 50 | 0 | 2019-12-04 | 16 hours ago | A simple, user-friendly and concurrent safe HTTP request library for Go. |
| [rq](https://github.com/ddo/rq) | 33 | 2 | 2017-12-26 | 4 months ago | A nicer interface for golang stdlib HTTP client. |

## OpenGL
        
*Libraries for using OpenGL in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [glfw](https://github.com/go-gl/glfw) | 839 | 41 | 2013-05-19 | 1 month ago | Go bindings for GLFW 3. |
| [gl](https://github.com/go-gl/gl) | 680 | 39 | 2015-02-22 | 10 months ago | Go bindings for OpenGL (generated via glow). |
| [mathgl](https://github.com/go-gl/mathgl) | 312 | 25 | 2013-02-13 | 3 months ago | Pure Go math package specialized for 3D math, with inspiration from GLM. |
| [gl](https://github.com/goxjs/gl) | 131 | 13 | 2015-05-18 | 1 year ago | Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android). |
| [glfw](https://github.com/goxjs/glfw) | 59 | 6 | 2014-12-27 | 1 month ago | Go cross-platform glfw library for creating an OpenGL context and receiving events. |

## ORM
        
*Libraries that implement Object-Relational Mapping or datamapping techniques.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gorm](https://github.com/jinzhu/gorm) | 16903 | 448 | 2013-10-25 | 4 days ago | The fantastic ORM library for Golang, aims to be developer friendly. |
| [xorm](https://github.com/go-xorm/xorm) | 5735 | 264 | 2013-05-09 | 3 months ago | Simple and powerful ORM for Go. |
| [pg](https://github.com/go-pg/pg) | 3461 | 84 | 2013-04-24 | 1 week ago | PostgreSQL ORM with focus on PostgreSQL specific features and performance. |
| [gorp](https://github.com/go-gorp/gorp) | 3267 | 115 | 2012-01-04 | 6 days ago | Go Relational Persistence, ORM-ish library for Go. |
| [sqlboiler](https://github.com/volatiletech/sqlboiler) | 2612 | 75 | 2016-02-21 | 15 hours ago | ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema. |
| [db](https://github.com/upper/db) | 2045 | 62 | 2013-10-23 | 2 months ago | Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers. |
| [reform](https://github.com/go-reform/reform) | 836 | 30 | 2016-02-25 | 3 months ago | Better ORM for Go, based on non-empty interfaces and code generation. |
| [pop](https://github.com/gobuffalo/pop) | 796 | 22 | 2018-02-07 | 2 days ago | Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite. |
| [qbs](https://github.com/coocood/qbs) | 544 | 45 | 2013-02-02 | 2 years ago | Stands for Query By Struct. A Go ORM. |
| [go-queryset](https://github.com/jirfag/go-queryset) | 481 | 19 | 2017-09-03 | 1 month ago | 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM. |
| [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) | 343 | 14 | 2017-12-27 | 2 months ago | A flexible and powerful SQL string builder library plus a zero-config ORM. |
| [gormt](https://github.com/xxjwxc/gormt) | 289 | 8 | 2019-05-05 | 2 hours ago | Mysql database to golang gorm struct. |
| [zoom](https://github.com/albrow/zoom) | 247 | 16 | 2013-07-17 | 1 year ago | Blazing-fast datastore and querying engine built on Redis. |
| [grimoire](https://github.com/Fs02/grimoire) | 123 | 5 | 2018-03-05 | 1 month ago | Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3). |
| [go-store](https://github.com/gosuri/go-store) | 100 | 9 | 2015-03-22 | 2 years ago | Simple and fast Redis backed key-value store library for Go. |
| [marlow](https://github.com/dadleyy/marlow) | 73 | 5 | 2017-09-15 | 2 weeks ago | Generated ORM from project structs for compile time safety assurances. |
| [rel](https://github.com/Fs02/rel) | 45 | 1 | 2019-10-06 | 1 day ago | Golang SQL Repository Layer for Clean (Onion) Architecture. |
| [go-firestorm](https://github.com/jschoedt/go-firestorm) | 7 | 1 | 2018-12-04 | 2 months ago | A simple ORM for Google/Firebase Cloud Firestore. |
| [lore](https://github.com/abrahambotros/lore) | 5 | 1 | 2017-04-29 | 2 years ago | Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go. |

## Package Management
        
*Unofficial libraries for package and dependency management.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [dep](https://github.com/golang/dep) | 13057 | 293 | 2016-10-07 | 3 weeks ago | Go dependency tool. |
| [glide](https://github.com/Masterminds/glide) | 7924 | 200 | 2014-07-09 | 3 weeks ago | Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip. |
| [godep](https://github.com/tools/godep) | 5648 | 153 | 2013-05-01 | 1 year ago | dependency tool for go, godep helps build packages reproducibly by fixing their dependencies. |
| [govendor](https://github.com/kardianos/govendor) | 4988 | 104 | 2015-04-12 | 3 weeks ago | Go Package Manager. Go vendor tool that works with the standard vendor file. |
| [gopm](https://github.com/gpmgo/gopm) | 2458 | 83 | 2013-05-15 | 5 months ago | Go Package Manager. |
| [gom](https://github.com/mattn/gom) | 1368 | 36 | 2013-09-11 | 5 months ago | Go Manager - bundle for go. |
| [gpm](https://github.com/pote/gpm) | 1204 | 31 | 2013-09-05 | 2 years ago | Barebones dependency manager for Go. |
| [goop](https://github.com/petejkim/goop) | 779 | 37 | 2014-06-18 | 4 years ago | Simple dependency manager for Go (golang), inspired by Bundler. |
| [nut](https://github.com/jingweno/nut) | 244 | 8 | 2015-01-23 | 4 years ago | Vendor Go dependencies. |
| [johnny-deps](https://github.com/VividCortex/johnny-deps) | 214 | 21 | 2013-07-19 | 5 years ago | Minimal dependency version using Git. |
| [gigo](https://github.com/LyricalSecurity/gigo) | 198 | 5 | 2015-05-01 | 1 year ago | PIP-like dependency tool for golang, with support for private repositories and hashes. |
| [VenGO](https://github.com/DamnWidget/VenGO) | 116 | 8 | 2014-10-17 | 3 years ago | create and manage exportable isolated go virtual environments. |
| [mvn-golang](https://github.com/raydac/mvn-golang) | 95 | 8 | 2016-03-24 | 5 days ago | plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure. |
| [gop](https://github.com/lunny/gop) | 50 | 7 | 2017-02-18 | 10 months ago | Build and manage your Go applications out of GOPATH. |

## Performance
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [jaeger](https://github.com/jaegertracing/jaeger) | 10067 | 320 | 2016-04-15 | 23 hours ago | A distributed tracing system. |
| [profile](https://github.com/pkg/profile) | 1163 | 36 | 2014-10-22 | 2 months ago | Simple profiling support package for Go. |
| [tracer](https://github.com/kamilsk/tracer) | 19 | 0 | 2019-06-22 | 1 week ago | Simple, lightweight tracing. |

## Query Language
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [graphql](https://github.com/graphql-go/graphql) | 5916 | 144 | 2015-07-19 | 1 month ago | Implementation of GraphQL for Go. |
| [graphql-go](https://github.com/graph-gophers/graphql-go) | 3134 | 89 | 2016-10-18 | 2 weeks ago | GraphQL server with a focus on ease of use. |
| [gojsonq](https://github.com/thedevsaddam/gojsonq) | 1208 | 26 | 2018-05-19 | 2 days ago | A simple Go package to Query over JSON Data. |
| [jsonql](https://github.com/elgs/jsonql) | 216 | 8 | 2015-12-29 | 3 months ago | JSON query expression library in Golang. |
| [rql](https://github.com/a8m/rql) | 130 | 6 | 2018-06-05 | 1 month ago | Resource Query Language for REST API. |
| [graphql](https://github.com/tmc/graphql) | 52 | 10 | 2015-04-18 | 2 years ago | graphql parser + utilities. |
| [jsonslice](https://github.com/bhmj/jsonslice) | 32 | 0 | 2018-05-02 | 4 weeks ago | Jsonpath queries with advanced filters. |
| [straf](https://github.com/SonicRoshan/straf) | 9 | 1 | 2019-08-16 | 3 months ago | Easily Convert Golang structs to GraphQL objects. |

## Resource Embedding
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [packr](https://github.com/gobuffalo/packr) | 2588 | 44 | 2017-03-15 | 3 days ago | The simple and easy way to embed static files into Go binaries. |
| [statik](https://github.com/rakyll/statik) | 2460 | 39 | 2014-02-04 | 2 months ago | Embeds static files into a Go executable. |
| [go.rice](https://github.com/GeertJohan/go.rice) | 1819 | 50 | 2013-10-23 | 1 month ago | go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy. |
| [vfsgen](https://github.com/shurcooL/vfsgen) | 783 | 21 | 2015-05-18 | 11 months ago | Generates a vfsdata.go file that statically implements the given virtual filesystem. |
| [esc](https://github.com/mjibson/esc) | 521 | 12 | 2014-01-26 | 2 months ago | Embeds files into Go programs and provides http.FileSystem interfaces to them. |
| [fileb0x](https://github.com/UnnoTed/fileb0x) | 497 | 15 | 2016-01-23 | 3 months ago | Simple tool to embed files in go with focus on "customization" and ease to use. |
| [go-resources](https://github.com/omeid/go-resources) | 162 | 7 | 2015-02-21 | 1 week ago | Unfancy resources embedding with Go. |
| [statics](https://github.com/go-playground/statics) | 54 | 3 | 2015-10-07 | 3 years ago | Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks. |
| [templify](https://github.com/wlbr/templify) | 22 | 2 | 2016-05-22 | 5 months ago | Embed external template files into Go code to create single file binaries. |
| [go-embed](https://github.com/pyros2097/go-embed) | 18 | 1 | 2015-12-06 | 5 months ago | Generates go code to embed resource files into your library or executable. |

## Science and Data Analysis
        
*Libraries for scientific computing and data analyzing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gonum](https://github.com/gonum/gonum) | 3458 | 98 | 2017-03-25 | 6 days ago | Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more. |
| [stats](https://github.com/montanaflynn/stats) | 1434 | 46 | 2014-12-16 | 21 hours ago | Statistics package with common functions missing from the Golang standard library. |
| [gosl](https://github.com/cpmech/gosl) | 1380 | 68 | 2015-02-09 | 1 month ago | Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more. |
| [plot](https://github.com/gonum/plot) | 1348 | 53 | 2013-07-23 | 1 week ago | gonum/plot provides an API for building and drawing plots in Go. |
| [streamtools](https://github.com/nytlabs/streamtools) | 1316 | 74 | 2013-07-05 | 4 years ago | general purpose, graphical tool for dealing with streams of data. |
| [go-dsp](https://github.com/mjibson/go-dsp) | 660 | 29 | 2011-11-02 | 1 year ago | Digital Signal Processing for Go. |
| [chart](https://github.com/vdobler/chart) | 619 | 44 | 2011-06-27 | 6 months ago | Simple Chart Plotting library for Go. Supports many graphs types. |
| [goraph](https://github.com/gyuho/goraph) | 619 | 38 | 2014-02-27 | 2 years ago | Pure Go graph theory library(data structure, algorith visualization). |
| [graph](https://github.com/yourbasic/graph) | 287 | 17 | 2017-04-27 | 2 months ago | Library of basic graph algorithms. |
| [ewma](https://github.com/VividCortex/ewma) | 285 | 25 | 2013-07-05 | 2 months ago | Exponentially-weighted moving averages. |
| [orb](https://github.com/paulmach/orb) | 248 | 18 | 2016-03-28 | 2 months ago | 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support. |
| [gohistogram](https://github.com/VividCortex/gohistogram) | 137 | 16 | 2013-07-02 | 1 year ago | Approximate histograms for data streams. |
| [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) | 134 | 12 | 2018-10-01 | 3 months ago | Dataframes for machine-learning and statistics (similar to pandas). |
| [TextRank](https://github.com/DavidBelicza/TextRank) | 90 | 6 | 2018-01-09 | 1 year ago | TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support. |
| [sparse](https://github.com/james-bowman/sparse) | 81 | 5 | 2017-05-16 | 1 week ago | Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries. |
| [pagerank](https://github.com/alixaxel/pagerank) | 53 | 6 | 2015-08-06 | 2 weeks ago | Weighted PageRank algorithm implemented in Go. |
| [geom](https://github.com/skelterjohn/geom) | 43 | 4 | 2011-06-07 | 2 years ago | 2D geometry for golang. |
| [evaler](https://github.com/soniah/evaler) | 40 | 4 | 2012-09-04 | 1 year ago | Simple floating point arithmetic expression evaluator. |
| [goent](https://github.com/kzahedi/goent) | 16 | 1 | 2017-08-08 | 9 months ago | GO Implementation of Entropy Measures. |
| [triangolatte](https://github.com/tchayen/triangolatte) | 12 | 1 | 2018-07-18 | 1 week ago | 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs. |
| [ode](https://github.com/ChristopherRabotin/ode) | 11 | 3 | 2016-11-11 | 2 years ago | Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions. |
| [GoStats](https://github.com/OGFris/GoStats) | 10 | 1 | 2018-07-22 | 1 year ago | GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions. |
| [piecewiselinear](https://github.com/sgreben/piecewiselinear) | 10 | 2 | 2018-10-21 | 2 weeks ago | Tiny linear interpolation library. |
| [PiHex](https://github.com/claygod/PiHex) | 8 | 2 | 2016-07-22 | 3 months ago | Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi. |
| [assocentity](https://github.com/ndabAP/assocentity) | 6 | 1 | 2018-12-21 | 1 week ago | Package assocentity returns the average distance from words to a given entity. |
| [go-gt](https://github.com/ThePaw/go-gt) | 5 | 0 | 2015-09-14 | 4 years ago | Graph theory algorithms written in "Go" language. |
| [rootfinding](https://github.com/khezen/rootfinding) | 3 | 2 | 2018-10-30 | 1 month ago | root-finding algorithms library for finding roots of quadratic functions. |
| [bradleyterry](https://github.com/seanhagen/bradleyterry) | 2 | 1 | 2019-04-30 | 8 months ago | Provides a Bradley-Terry Model for pairwise comparisons. |

## Security
        
*Libraries that are used to help make your application more secure.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [lego](https://github.com/go-acme/lego) | 3783 | 100 | 2015-06-08 | 2 days ago | Pure Go ACME client library and CLI tool (for use with Let's Encrypt). |
| [cameradar](https://github.com/Ullaakut/cameradar) | 2010 | 101 | 2016-05-20 | 9 hours ago | Tool and library to remotely hack RTSP streams from surveillance cameras. |
| [acmetool](https://github.com/hlandau/acmetool) | 1731 | 65 | 2015-11-15 | 21 hours ago | ACME (Let's Encrypt) client tool with automatic renewal. |
| [memguard](https://github.com/awnumar/memguard) | 1696 | 43 | 2017-04-22 | 3 weeks ago | A pure Go library for handling sensitive values in memory. |
| [secure](https://github.com/unrolled/secure) | 1400 | 37 | 2014-05-20 | 1 month ago | HTTP middleware for Go that facilitates some quick security wins. |
| [acra](https://github.com/cossacklabs/acra) | 526 | 33 | 2016-11-14 | 5 days ago | Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system. |
| [nacl](https://github.com/kevinburke/nacl) | 463 | 12 | 2017-07-20 | 4 months ago | Go implementation of the NaCL set of API's. |
| [badactor](https://github.com/jaredfolkins/badactor) | 262 | 9 | 2014-12-12 | 5 months ago | In-memory, application-driven jailer built in the spirit of fail2ban. |
| [passlib](https://github.com/hlandau/passlib) | 232 | 10 | 2014-12-21 | 9 months ago | Futureproof password hashing library. |
| [ssh-vault](https://github.com/ssh-vault/ssh-vault) | 215 | 10 | 2016-09-29 | 1 month ago | encrypt/decrypt using ssh keys. |
| [simple-scrypt](https://github.com/elithrar/simple-scrypt) | 163 | 7 | 2015-04-14 | 6 months ago | Scrypt package with a simple, obvious API and automatic cost calibration built-in. |
| [go-yara](https://github.com/hillu/go-yara) | 146 | 17 | 2015-01-25 | 1 month ago | Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)". |
| [argon2pw](https://github.com/raja/argon2pw) | 80 | 3 | 2018-03-13 | 1 year ago | Argon2 password hash generation with constant-time password comparison. |
| [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) | 34 | 1 | 2017-10-19 | 11 months ago | A probably paranoid package for securely hashing and encrypting passwords. |
| [goArgonPass](https://github.com/dwin/goArgonPass) | 12 | 2 | 2018-05-30 | 10 months ago | Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations. |
| [certificates](https://github.com/mvmaasakkers/certificates) | 11 | 1 | 2019-03-04 | 1 week ago | An opinionated tool for generating tls certificates. |
| [sslmgr](https://github.com/adrianosela/sslmgr) | 8 | 2 | 2019-04-02 | 5 months ago | SSL certificates made easy with a high level wrapper around acme/autocert. |
| [go-generate-password](https://github.com/m1/go-generate-password) | 6 | 1 | 2019-11-14 | 2 months ago | Password generator that can be used on the cli or as a library. |

## Serialization
        
*Libraries and tools for binary serialization.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go](https://github.com/json-iterator/go) | 7054 | 216 | 2016-11-30 | 3 days ago | High-performance 100% compatible drop-in replacement of "encoding/json". |
| [protobuf](https://github.com/golang/protobuf) | 6076 | 207 | 2014-11-23 | 4 days ago | Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers. |
| [protobuf](https://github.com/gogo/protobuf) | 3452 | 94 | 2014-12-03 | 5 days ago | Protocol Buffers for Go with Gadgets. |
| [mapstructure](https://github.com/mitchellh/mapstructure) | 3000 | 55 | 2013-05-20 | 2 months ago | Go library for decoding generic map values into native Go structures. |
| [go](https://github.com/ugorji/go) | 1343 | 53 | 2013-05-30 | 1 month ago | High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support. |
| [colfer](https://github.com/pascaldekloe/colfer) | 508 | 34 | 2015-09-05 | 2 months ago | Code generation for the Colfer binary format. |
| [csvutil](https://github.com/jszwec/csvutil) | 342 | 9 | 2017-10-30 | 20 hours ago | High Performance, idiomatic CSV record encoding and decoding to native Go structures. |
| [go-capnproto](https://github.com/glycerine/go-capnproto) | 276 | 11 | 2013-11-07 | 1 year ago | Cap'n Proto library and parser for go. |
| [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) | 133 | 9 | 2012-12-23 | 1 year ago | GoLang library for working with PHP session format and PHP Serialize/Unserialize functions. |
| [structomap](https://github.com/danhper/structomap) | 108 | 7 | 2015-05-13 | 8 months ago | Library to easily and dynamically generate maps from static structures. |
| [cbor](https://github.com/fxamacker/cbor) | 102 | 4 | 2019-05-15 | 2 days ago | Small, safe, and easy CBOR encoding and decoding library. |
| [bambam](https://github.com/glycerine/bambam) | 60 | 4 | 2014-09-17 | 3 years ago | generator for Cap'n Proto schemas from go. |
| [asn1](https://github.com/Logicalis/asn1) | 43 | 8 | 2016-02-29 | 10 months ago | Asn.1 BER and DER encoding library for golang. |
| [binstruct](https://github.com/ghostiam/binstruct) | 14 | 1 | 2018-10-23 | 4 months ago | Golang binary decoder for mapping data into the structure. |
| [fwencoder](https://github.com/o1egl/fwencoder) | 10 | 1 | 2017-12-25 | 2 years ago | Fixed width file parser (encoding and decoding library) for Go. |
| [bel](https://github.com/32leaves/bel) | 8 | 1 | 2019-02-20 | 9 months ago | Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC. |
| [pletter](https://github.com/vimeda/pletter) | 7 | 0 | 2019-07-09 | 6 months ago | A standard way to wrap a proto message for message brokers. |
| [fixedwidth](https://github.com/huydang284/fixedwidth) | 4 | 1 | 2019-08-11 | 1 month ago | Fixed-width text formatting (UTF-8 supported). |

## Server Applications
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [etcd](https://github.com/etcd-io/etcd) | 29286 | 1332 | 2013-07-06 | 37 minutes ago | Highly-available key value store for shared configuration and service discovery. |
| [caddy](https://github.com/caddyserver/caddy) | 25993 | 737 | 2015-01-13 | 1 day ago | Caddy is an alternative, HTTP/2 web server that's easy to configure and use. |
| [minio](https://github.com/minio/minio) | 19968 | 480 | 2015-01-14 | 2 hours ago | Minio is a distributed object storage server. |
| [roadrunner](https://github.com/spiral/roadrunner) | 3906 | 151 | 2017-12-26 | 3 days ago | High-performance PHP application server, load-balancer and process manager. |
| [devd](https://github.com/cortesi/devd) | 2933 | 69 | 2015-09-27 | 5 months ago | Local webserver for developers. |
| [algernon](https://github.com/xyproto/algernon) | 1656 | 49 | 2015-03-10 | 1 week ago | HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber. |
| [sftpgo](https://github.com/drakkan/sftpgo) | 1193 | 29 | 2019-07-20 | 7 hours ago | Full featured and highly configurable SFTP server software. |
| [flipt](https://github.com/markphelps/flipt) | 1146 | 14 | 2016-11-05 | 1 day ago | A self contained feature flag solution written in Go and Vue. |
| [flagr](https://github.com/checkr/flagr) | 975 | 69 | 2017-10-03 | 2 weeks ago | Flagr is an open-source feature flagging and A/B testing service. |
| [fider](https://github.com/getfider/fider) | 947 | 21 | 2017-01-17 | 2 days ago | Fider is an open platform to collect and organize customer feedback. |
| [discovery](https://github.com/bilibili/discovery) | 832 | 48 | 2018-04-20 | 1 month ago | A registry for resilient mid-tier load balancing and failover. |
| [jackal](https://github.com/ortuman/jackal) | 771 | 34 | 2017-11-13 | 1 week ago | An XMPP server written in Go. |
| [dudeldu](https://github.com/krotik/dudeldu) | 108 | 3 | 2016-09-07 | 4 months ago | A simple SHOUTcast server. |
| [psql-streamer](https://github.com/blind-oracle/psql-streamer) | 12 | 4 | 2019-04-28 | 8 months ago | Stream database events from PostgreSQL to Kafka. |
| [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) | 9 | 1 | 2018-10-23 | 8 months ago | Nginx log parser and exporter to Prometheus. |
| [riemann-relay](https://github.com/blind-oracle/riemann-relay) | 0 | 1 | 2019-04-23 | 2 months ago | Relay to load-balance Riemann events and/or convert them to Carbon. |

## Stream Processing
        
*Libraries and tools for stream processing and reactive programming.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-streams](https://github.com/reugn/go-streams) | 266 | 9 | 2019-04-30 | 4 days ago | Go stream processing library. |

## Template Engines
        
*Libraries and tools for templating and lexing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gofpdf](https://github.com/jung-kurt/gofpdf) | 3450 | 91 | 2015-03-13 | 2 months ago | PDF document generator with high level support for text, drawing and images. |
| [quicktemplate](https://github.com/valyala/quicktemplate) | 1601 | 56 | 2016-03-06 | 2 hours ago | Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it. |
| [pongo2](https://github.com/flosch/pongo2) | 1592 | 55 | 2013-08-23 | 3 months ago | Django-like template-engine for Go. |
| [hero](https://github.com/shiyanhui/hero) | 1264 | 41 | 2017-01-15 | 1 week ago | Hero is a handy, fast and powerful go template engine. |
| [mustache](https://github.com/hoisie/mustache) | 977 | 36 | 2009-12-30 | 4 weeks ago | Go implementation of the Mustache template language. |
| [amber](https://github.com/eknkc/amber) | 836 | 21 | 2012-10-31 | 1 year ago | Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade. |
| [ace](https://github.com/yosssi/ace) | 778 | 24 | 2014-07-13 | 1 year ago | Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold. |
| [gorazor](https://github.com/sipin/gorazor) | 728 | 58 | 2014-05-01 | 2 months ago | Razor view engine for Golang. |
| [jet](https://github.com/CloudyKit/jet) | 616 | 21 | 2016-03-31 | 6 days ago | Jet template engine. |
| [ego](https://github.com/benbjohnson/ego) | 425 | 16 | 2014-02-23 | 2 weeks ago | Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled. |
| [raymond](https://github.com/aymerick/raymond) | 359 | 11 | 2015-04-22 | 11 months ago | Complete handlebars implementation in Go. |
| [fasttemplate](https://github.com/valyala/fasttemplate) | 340 | 17 | 2015-08-19 | 4 days ago | Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/). |
| [soy](https://github.com/robfig/soy) | 147 | 12 | 2013-12-15 | 10 months ago | Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/). |
| [maroto](https://github.com/johnfercher/maroto) | 96 | 6 | 2019-05-20 | 1 week ago | A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple. |
| [liquid](https://github.com/osteele/liquid) | 94 | 3 | 2017-06-26 | 1 month ago | Go implementation of Shopify Liquid templates. |
| [goview](https://github.com/foolin/goview) | 91 | 3 | 2019-04-14 | 2 months ago | Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application. |
| [kasia.go](https://github.com/ziutek/kasia.go) | 71 | 2 | 2010-12-07 | 4 years ago | Templating system for HTML and other text documents - go implementation. |
| [velvet](https://github.com/gobuffalo/velvet) | 70 | 5 | 2016-12-29 | 2 years ago | Complete handlebars implementation in Go. |
| [damsel](https://github.com/dskinner/damsel) | 21 | 4 | 2012-05-02 | 3 years ago | Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others. |
| [extemplate](https://github.com/dannyvankooten/extemplate) | 18 | 2 | 2018-08-10 | 1 year ago | Tiny wrapper around html/template to allow for easy file-based template inheritance. |
| [gospin](https://github.com/m1/gospin) | 5 | 1 | 2019-02-22 | 1 month ago | Article spinning and spintax/spinning syntax engine, useful for A/B, testing pieces of text/articles and creating more natural conversations. |

## Testing
        
*Libraries for testing codebases and generating test data.*

### Testing Frameworks
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [testify](https://github.com/stretchr/testify) | 9508 | 150 | 2012-10-16 | 5 days ago | Sacred extension to the standard go testing package. |
| [go-cmp](https://github.com/google/go-cmp) | 1486 | 24 | 2017-07-07 | 1 week ago | Package for comparing Go values in tests. |
| [httpexpect](https://github.com/gavv/httpexpect) | 1264 | 32 | 2016-04-29 | 1 week ago | Concise, declarative, and easy to use end-to-end HTTP and REST API testing. |
| [godog](https://github.com/DATA-DOG/godog) | 890 | 28 | 2015-06-10 | 22 hours ago | Cucumber or Behat like BDD framework for Go. |
| [baloo](https://github.com/h2non/baloo) | 668 | 10 | 2016-05-29 | 1 year ago | Expressive and versatile end-to-end HTTP API testing made easy. |
| [goblin](https://github.com/franela/goblin) | 652 | 16 | 2013-09-19 | 2 weeks ago | Mocha like testing framework fo Go. |
| [testfixtures](https://github.com/go-testfixtures/testfixtures) | 428 | 5 | 2016-04-05 | 1 week ago | A helper for Rails' like test fixtures to test database applications. |
| [go-vcr](https://github.com/dnaeon/go-vcr) | 367 | 7 | 2015-12-14 | 1 month ago | Record and replay your HTTP interactions for fast, deterministic and accurate tests. |
| [go-mutesting](https://github.com/zimmski/go-mutesting) | 322 | 7 | 2014-12-26 | 3 months ago | Mutation testing for Go source code. |
| [gofight](https://github.com/appleboy/gofight) | 293 | 11 | 2016-03-29 | 2 weeks ago | API Handler Testing for Golang Router framework. |
| [frisby](https://github.com/verdverm/frisby) | 255 | 8 | 2015-09-15 | 7 months ago | REST API testing framework. |
| [go-carpet](https://github.com/msoap/go-carpet) | 201 | 4 | 2016-02-28 | 1 month ago | Tool for viewing test coverage in terminal. |
| [charlatan](https://github.com/percolate/charlatan) | 192 | 40 | 2017-10-06 | 4 months ago | Tool to generate fake interface implementations for tests. |
| [gotest.tools](https://github.com/gotestyourself/gotest.tools) | 142 | 7 | 2017-08-08 | 3 weeks ago | A collection of packages to augment the go testing package and support common patterns. |
| [endly](https://github.com/viant/endly) | 122 | 14 | 2017-08-28 | 1 hour ago | Declarative end to end functional testing. |
| [gospec](https://github.com/luontola/gospec) | 112 | 4 | 2009-11-24 | 5 years ago | BDD-style testing framework for the Go programming language. |
| [cupaloy](https://github.com/bradleyjkemp/cupaloy) | 98 | 3 | 2017-08-07 | 5 months ago | Simple snapshot testing addon for your test framework. |
| [dbcleaner](https://github.com/khaiql/dbcleaner) | 98 | 3 | 2017-01-17 | 9 months ago | Clean database for testing purpose, inspired by `database_cleaner` in Ruby. |
| [commander](https://github.com/SimonBaeumer/commander) | 83 | 3 | 2019-02-22 | 2 months ago | Tool for testing cli applications on windows, linux and osx. |
| [go-testdeep](https://github.com/maxatome/go-testdeep) | 75 | 2 | 2018-05-26 | 1 month ago | Extremely flexible golang deep comparison, extends the go testing package. |
| [wstest](https://github.com/posener/wstest) | 71 | 2 | 2017-03-31 | 4 weeks ago | Websocket client for unit-testing a websocket http.Handler. |
| [gospecify](https://github.com/stesla/gospecify) | 53 | 5 | 2009-11-20 | 8 years ago | This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec. |
| [restit](https://github.com/go-restit/restit) | 50 | 6 | 2014-06-25 | 3 months ago | Go micro framework to help writing RESTful API integration test. |
| [gomatch](https://github.com/jfilipczyk/gomatch) | 32 | 2 | 2019-01-27 | 6 months ago | library created for testing JSON against patterns. |
| [jsonassert](https://github.com/kinbiko/jsonassert) | 30 | 0 | 2018-10-26 | 6 days ago | Package for verifying that your JSON payloads are serialized correctly. |
| [dsunit](https://github.com/viant/dsunit) | 28 | 9 | 2016-06-13 | 6 days ago | Datastore testing for SQL, NoSQL, structured files. |
| [hamcrest](https://github.com/rdrdr/hamcrest) | 27 | 3 | 2010-12-22 | 9 years ago | fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results. |
| [assert](https://github.com/go-playground/assert) | 17 | 1 | 2015-07-20 | 3 months ago | Basic Assertion Library used along side native go testing, with building blocks for custom assertions. |
| [testcase](https://github.com/adamluzsi/testcase) | 12 | 1 | 2019-04-22 | 6 days ago | Idiomatic testing framework for Behavior Driven Development. |
| [flute](https://github.com/suzuki-shunsuke/flute) | 11 | 2 | 2019-07-06 | 1 month ago | HTTP client testing framework. |
| [badio](https://github.com/cavaliercoder/badio) | 9 | 1 | 2016-02-11 | 3 years ago | Extensions to Go's `testing/iotest` package. |
| [embedded-postgres](https://github.com/fergusstrange/embedded-postgres) | 9 | 1 | 2019-11-16 | 1 month ago | Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test. |
| [gocrest](https://github.com/corbym/gocrest) | 9 | 1 | 2017-12-23 | 1 year ago | Composable hamcrest-like matchers for Go assertions. |
| [gosuite](https://github.com/pavlo/gosuite) | 9 | 1 | 2016-10-15 | 3 years ago | Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests. |
| [gogiven](https://github.com/corbym/gogiven) | 8 | 4 | 2017-12-31 | 1 year ago | YATSPEC-like BDD testing framework for Go. |
| [schema](https://github.com/jgroeneveld/schema) | 8 | 3 | 2015-08-13 | 3 months ago | Quick and easy expression matching for JSON schemas used in requests and responses. |
| [biff](https://github.com/fulldump/biff) | 7 | 1 | 2018-03-28 | 1 year ago | Bifurcation testing framework, BDD compatible. |
| [testsql](https://github.com/zhulongcheng/testsql) | 7 | 2 | 2018-09-22 | 3 months ago | Generate test data from SQL files before testing and clear it after finished. |
| [tt](https://github.com/vcaesar/tt) | 6 | 1 | 2018-04-03 | 1 hour ago | Simple and colorful test tools. |
| [trial](https://github.com/jgroeneveld/trial) | 4 | 1 | 2015-06-18 | 3 months ago | Quick and easy extendable assertions without introducing much boilerplate. |

### Mock
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [chromedp](https://github.com/chromedp/chromedp) | 4232 | 138 | 2017-01-24 | 1 day ago | a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol. |
| [mock](https://github.com/golang/mock) | 3578 | 73 | 2015-06-12 | 21 hours ago | Mocking framework for the Go programming language. |
| [go-fuzz](https://github.com/dvyukov/go-fuzz) | 3281 | 89 | 2015-04-15 | 1 day ago | Randomized testing system. |
| [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) | 2176 | 23 | 2014-02-07 | 2 weeks ago | Mock SQL driver for testing database interactions. |
| [hoverfly](https://github.com/SpectoLabs/hoverfly) | 1512 | 58 | 2015-11-30 | 3 days ago | HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI. |
| [selenoid](https://github.com/aerokube/selenoid) | 1463 | 93 | 2016-08-22 | 4 days ago | alternative Selenium hub server that launches browsers within containers. |
| [gock](https://github.com/h2non/gock) | 942 | 17 | 2016-03-02 | 2 days ago | Versatile HTTP mocking made easy. |
| [httpmock](https://github.com/jarcoal/httpmock) | 693 | 8 | 2014-02-24 | 1 month ago | Easy mocking of HTTP responses from external resources. |
| [gofuzz](https://github.com/google/gofuzz) | 693 | 24 | 2014-07-31 | 2 months ago | Library for populating go objects with random values. |
| [cdp](https://github.com/mafredri/cdp) | 400 | 17 | 2017-03-12 | 1 month ago | Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it. |
| [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) | 394 | 7 | 2014-05-21 | 17 hours ago | Tool for generating self-contained mock objects. |
| [minimock](https://github.com/gojuno/minimock) | 288 | 9 | 2016-08-03 | 2 months ago | Mock generator for Go interfaces. |
| [ggr](https://github.com/aerokube/ggr) | 234 | 25 | 2016-06-16 | 1 month ago | a lightweight server that routes and proxies Selenium WebDriver requests to multiple Selenium hubs. |
| [go-txdb](https://github.com/DATA-DOG/go-txdb) | 219 | 9 | 2015-07-08 | 1 week ago | Single transaction based database driver mainly for testing purposes. |
| [tavor](https://github.com/zimmski/tavor) | 218 | 12 | 2014-05-18 | 1 year ago | Generic fuzzing and delta-debugging framework. |
| [govcr](https://github.com/seborama/govcr) | 87 | 2 | 2016-07-10 | 3 months ago | HTTP mock for Golang: record and replay HTTP interactions for offline testing. |
| [timex](https://github.com/cabify/timex) | 23 | 63 | 2020-01-02 | 1 week ago | A test-friendly replacement for the native `time` package. |
| [mockhttp](https://github.com/tv42/mockhttp) | 22 | 1 | 2011-06-11 | 5 years ago | Mock object for Go http.ResponseWriter. |

### Fail injection
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [failpoint](https://github.com/pingcap/failpoint) | 474 | 76 | 2019-04-02 | 6 days ago | An implementation of [failpoints](http://www.freebsd.org/cgi/man.cgi?query=fail) for Golang. |

## Text Processing
        
*Libraries for parsing and manipulating texts.*

### Specific Formats
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [colly](https://github.com/gocolly/colly) | 9864 | 272 | 2017-09-29 | 1 day ago | Fast and Elegant Scraping Framework for Gophers. |
| [goquery](https://github.com/PuerkitoBio/goquery) | 8281 | 269 | 2012-08-29 | 1 month ago | GoQuery brings a syntax and a set of features similar to jQuery to the Go language. |
| [blackfriday](https://github.com/russross/blackfriday) | 4202 | 98 | 2011-05-27 | 1 month ago | Markdown processor in Go. |
| [toml](https://github.com/BurntSushi/toml) | 3023 | 83 | 2013-02-26 | 2 days ago | TOML configuration format (encoder/decoder with reflection). |
| [sh](https://github.com/mvdan/sh) | 2406 | 45 | 2016-01-16 | 7 hours ago | Shell parser and formatter. |
| [go-humanize](https://github.com/dustin/go-humanize) | 2077 | 32 | 2012-01-13 | 1 month ago | Formatters for time, numbers, and memory size to human readable format. |
| [bluemonday](https://github.com/microcosm-cc/bluemonday) | 1383 | 32 | 2013-11-20 | 2 months ago | HTML Sanitizer. |
| [inject](https://github.com/facebookarchive/inject) | 1190 | 44 | 2013-10-21 | 1 year ago | Package inject provides a reflect based injector. |
| [gofeed](https://github.com/mmcdole/gofeed) | 1180 | 37 | 2016-01-23 | 1 month ago | Parse RSS and Atom feeds in Go. |
| [go-toml](https://github.com/pelletier/go-toml) | 694 | 27 | 2013-02-24 | 2 days ago | Go library for the TOML format with query support and handy cli tools. |
| [commonregex](https://github.com/mingrammer/commonregex) | 573 | 22 | 2017-03-23 | 2 months ago | A collection of common regular expressions for Go. |
| [slug](https://github.com/gosimple/slug) | 495 | 11 | 2014-03-31 | 5 days ago | URL-friendly slugify with multiple languages support. |
| [mxj](https://github.com/clbanning/mxj) | 357 | 23 | 2014-02-03 | 1 month ago | Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages. |
| [dataflowkit](https://github.com/slotix/dataflowkit) | 342 | 16 | 2017-02-09 | 3 months ago | Web scraping Framework to turn websites into structured data. |
| [gographviz](https://github.com/awalterschulze/gographviz) | 339 | 13 | 2015-03-14 | 7 months ago | Parses the Graphviz DOT language. |
| [gotext](https://github.com/leonelquinteros/gotext) | 248 | 5 | 2016-06-19 | 2 months ago | GNU gettext utilities for Go. |
| [go-runewidth](https://github.com/mattn/go-runewidth) | 240 | 10 | 2013-06-21 | 1 week ago | Functions to get fixed width of the character or string. |
| [htmlquery](https://github.com/antchfx/htmlquery) | 187 | 8 | 2017-12-05 | 1 week ago | An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression. |
| [goq](https://github.com/andrewstuart/goq) | 158 | 7 | 2017-02-20 | 7 months ago | Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery). |
| [go-nmea](https://github.com/adrianmo/go-nmea) | 109 | 7 | 2015-07-22 | 3 months ago | NMEA parser library for the Go language. |
| [sdp](https://github.com/gortc/sdp) | 80 | 7 | 2016-05-13 | 3 days ago | SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)]. |
| [go-zero-width](https://github.com/trubitsyn/go-zero-width) | 71 | 1 | 2018-06-18 | 2 days ago | Zero-width character detection and removal for Go. |
| [align](https://github.com/Guitarbum722/align) | 62 | 4 | 2017-04-29 | 2 months ago | A general purpose application that aligns text. |
| [genex](https://github.com/alixaxel/genex) | 57 | 3 | 2015-03-09 | 2 weeks ago | Count and expand Regular Expressions into all matching Strings. |
| [go-slugify](https://github.com/mozillazg/go-slugify) | 57 | 2 | 2016-07-16 | 3 years ago | Make pretty slug with multiple languages support. |
| [goribot](https://github.com/zhshch2002/goribot) | 57 | 5 | 2019-09-08 | 2 months ago | A simple golang spider/scraping framework,build a spider in 3 lines. |
| [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) | 54 | 8 | 2016-07-05 | 1 day ago | Editorconfig file parser and manipulator for Go. |
| [guesslanguage](https://github.com/endeveit/guesslanguage) | 45 | 1 | 2014-12-16 | 2 years ago | Functions to determine the natural language of a unicode text. |
| [allot](https://github.com/sbstjn/allot) | 38 | 1 | 2016-10-16 | 8 months ago | Placeholder and wildcard text parsing for CLI tools and bots. |
| [goregen](https://github.com/zach-klippenstein/goregen) | 38 | 2 | 2014-12-27 | 4 months ago | Library for generating random strings from regular expressions. |
| [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) | 33 | 3 | 2017-11-15 | 2 weeks ago | Fixed-width text formatting (encoder/decoder with reflection). |
| [go-vcard](https://github.com/emersion/go-vcard) | 32 | 3 | 2017-03-21 | 1 month ago | Parse and format vCard. |
| [gonameparts](https://github.com/polera/gonameparts) | 30 | 0 | 2015-05-17 | 5 months ago | Parses human names into individual name parts. |
| [slugify](https://github.com/avelino/slugify) | 27 | 2 | 2015-04-13 | 1 year ago | Go slugify application that handles string. |
| [did](https://github.com/ockam-network/did) | 26 | 7 | 2018-11-02 | 8 months ago | DID (Decentralized Identifiers) Parser and Stringer in Go. |
| [codetree](https://github.com/aerogo/codetree) | 14 | 2 | 2016-11-26 | 2 months ago | Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure. |
| [enca](https://github.com/endeveit/enca) | 8 | 1 | 2014-12-17 | 3 years ago | Minimal cgo bindings for [libenca](http://cihar.com/software/enca/). |
| [syndfeed](https://github.com/zhengchun/syndfeed) | 7 | 1 | 2017-04-07 | 1 year ago | A syndication feed for Atom 1.0 and RSS 2.0. |
| [bbConvert](https://github.com/CalebQ42/bbConvert) | 5 | 1 | 2016-04-15 | 3 years ago | Converts bbCode to HTML that allows you to add support for custom bbCode tags. |
| [doi](https://github.com/hscells/doi) | 4 | 1 | 2017-08-02 | 2 years ago | Document object identifier (doi) parser in Go. |
| [encoding](https://github.com/mickep76/encoding) | 3 | 1 | 2018-04-06 | 2 months ago | Package provides a generic interface to encoders and decodersa. |
| [ltsv](https://github.com/Wing924/ltsv) | 3 | 1 | 2019-05-12 | 7 months ago | High performance [LTSV (Labeled Tab Separeted Value)](http://ltsv.org/) reader for Go. |

### Utility
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [xurls](https://github.com/mvdan/xurls) | 573 | 14 | 2015-01-12 | 3 weeks ago | Extract urls from text. |
| [gotabulate](https://github.com/bndr/gotabulate) | 224 | 7 | 2014-08-21 | 2 years ago | Easily pretty-print your tabular data with Go. |
| [radix](https://github.com/yourbasic/radix) | 156 | 5 | 2017-06-09 | 1 year ago | fast string sorting algorithm. |
| [parth](https://github.com/codemodus/parth) | 34 | 3 | 2015-04-06 | 11 months ago | URL path segmentation parsing. |
| [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) | 19 | 1 | 2018-09-09 | 1 year ago | A sanitization-based swear filter for Go. |
| [xj2go](https://github.com/stackerzzq/xj2go) | 18 | 2 | 2017-09-19 | 1 week ago | Convert xml or json to go struct. |
| [kace](https://github.com/codemodus/kace) | 12 | 1 | 2015-06-04 | 1 year ago | Common case conversions covering common initialisms. |
| [parseargs-go](https://github.com/txgruppi/parseargs-go) | 7 | 1 | 2016-02-24 | 3 years ago | string argument parser that understands quotes and backslashes. |
| [tagify](https://github.com/zoomio/tagify) | 3 | 1 | 2018-03-20 | 2 days ago | Produces a set of tags from given source. |
| [TySug](https://github.com/Dynom/TySug) | 3 | 1 | 2018-06-05 | 2 weeks ago | Alternative suggestions with respect to keyboard layouts. |
| [textwrap](https://github.com/isbm/textwrap) | 1 | 1 | 2019-07-26 | 5 months ago | Implementation of `textwrap` module from Python. |

## Third-party APIs
        
*Libraries for accessing third party APIs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [aws-sdk-go](https://github.com/aws/aws-sdk-go) | 5514 | 258 | 2014-12-05 | 11 hours ago | The official AWS SDK for the Go programming language. |
| [go-github](https://github.com/google/go-github) | 5426 | 191 | 2013-05-24 | 14 hours ago | Go library for accessing the GitHub REST API v3. |
| [slack](https://github.com/nlopes/slack) | 2709 | 55 | 2015-01-24 | 17 hours ago | Slack API in Go. |
| [google-api-go-client](https://github.com/googleapis/google-api-go-client) | 2099 | 133 | 2014-11-24 | 16 hours ago | Auto-generated Google APIs for Go. |
| [google-cloud-go](https://github.com/googleapis/google-cloud-go) | 1988 | 219 | 2014-05-09 | 6 hours ago | Google Cloud APIs Go Client Library. |
| [discordgo](https://github.com/bwmarrin/discordgo) | 1105 | 52 | 2015-11-01 | 22 hours ago | Go bindings for the Discord Chat API. |
| [stripe-go](https://github.com/stripe/stripe-go) | 1069 | 39 | 2014-06-05 | 20 minutes ago | Go client for the Stripe API. |
| [anaconda](https://github.com/ChimeraCoder/anaconda) | 1012 | 21 | 2013-03-04 | 1 week ago | Go client library for the Twitter 1.1 API. |
| [go-twitter](https://github.com/dghubble/go-twitter) | 889 | 29 | 2015-04-11 | 4 days ago | Go client library for the Twitter v1.1 APIs. |
| [minio-go](https://github.com/minio/minio-go) | 851 | 39 | 2015-05-02 | 14 hours ago | Minio Go Library for Amazon S3 compatible cloud storage. |
| [facebook](https://github.com/huandu/facebook) | 827 | 84 | 2012-07-28 | 1 month ago | Go Library that supports the Facebook Graph API. |
| [githubv4](https://github.com/shurcooL/githubv4) | 588 | 19 | 2017-05-27 | 1 month ago | Go library for accessing the GitHub GraphQL API v4. |
| [webhooks](https://github.com/go-playground/webhooks) | 438 | 15 | 2015-10-25 | 2 months ago | Webhook receiver for GitHub and Bitbucket. |
| [geo-golang](https://github.com/codingsince1985/geo-golang) | 335 | 11 | 2014-12-04 | 4 days ago | Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs. |
| [paypal](https://github.com/plutov/paypal) | 335 | 15 | 2015-10-14 | 2 weeks ago | Wrapper for PayPal payment API. |
| [go-marathon](https://github.com/gambol99/go-marathon) | 192 | 13 | 2015-02-11 | 5 days ago | Go library for interacting with Mesosphere's Marathon PAAS. |
| [ethrpc](https://github.com/onrik/ethrpc) | 181 | 13 | 2017-01-24 | 4 months ago | Go bindings for Ethereum JSON RPC API. |
| [trello](https://github.com/adlio/trello) | 126 | 6 | 2016-09-24 | 1 month ago | Go wrapper for the Trello API. |
| [gostorm](https://github.com/jsgilmore/gostorm) | 124 | 11 | 2013-07-22 | 2 years ago | GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. |
| [medium-sdk-go](https://github.com/Medium/medium-sdk-go) | 122 | 123 | 2015-09-26 | 1 year ago | Golang SDK for Medium's OAuth2 API. |
| [hipchat](https://github.com/daneharrigan/hipchat) | 112 | 7 | 2013-04-28 | 2 years ago | A golang package to communicate with HipChat over XMPP. |
| [hipchat](https://github.com/andybons/hipchat) | 108 | 8 | 2012-10-20 | 3 years ago | This project implements a golang client library for the Hipchat API. |
| [go-trending](https://github.com/andygrunwald/go-trending) | 105 | 7 | 2015-07-04 | 8 months ago | Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github. |
| [cachet](https://github.com/andygrunwald/cachet) | 78 | 5 | 2015-10-31 | 1 year ago | Go client library for [Cachet (open source status page system)](https://cachethq.io/). |
| [pushover](https://github.com/gregdel/pushover) | 67 | 3 | 2015-02-19 | 11 months ago | Go wrapper for the Pushover API. |
| [wit-go](https://github.com/wit-ai/wit-go) | 59 | 11 | 2018-08-20 | 5 months ago | Go client for wit.ai HTTP API. |
| [clarifai-go](https://github.com/Clarifai/clarifai-go) | 57 | 36 | 2015-09-28 | 2 years ago | Go client library for interfacing with the Clarifai API. |
| [megos](https://github.com/andygrunwald/megos) | 57 | 5 | 2015-10-02 | 1 year ago | Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster. |
| [igdb](https://github.com/Henry-Sarabia/igdb) | 55 | 2 | 2017-08-24 | 4 weeks ago | Go client for the [Internet Game Database API](https://api.igdb.com/). |
| [gads](https://github.com/emiddleton/gads) | 47 | 7 | 2014-01-20 | 4 months ago | Google Adwords Unofficial API. |
| [go-circleci](https://github.com/jszwedko/go-circleci) | 46 | 3 | 2015-08-14 | 2 months ago | Go client library for interacting with CircleCI's API. |
| [go-xkcd](https://github.com/nishanths/go-xkcd) | 40 | 3 | 2016-02-26 | 3 years ago | Go client for the xkcd API. |
| [go-amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) | 39 | 1 | 2016-11-15 | 1 year ago | Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html). |
| [gomusicbrainz](https://github.com/michiwend/gomusicbrainz) | 37 | 8 | 2014-09-10 | 4 months ago | Go MusicBrainz WS2 client library. |
| [simples3](https://github.com/rhnvrm/simples3) | 36 | 1 | 2018-12-06 | 2 weeks ago | Simple no frills AWS S3 Library using REST with V4 Signing written in Go. |
| [uptimerobot](https://github.com/bitfield/uptimerobot) | 35 | 3 | 2018-05-29 | 2 months ago | Go wrapper and command-line client for the Uptime Robot v2 API. |
| [fcm](https://github.com/maddevsio/fcm) | 34 | 4 | 2017-01-06 | 10 months ago | Go library for Firebase Cloud Messaging. |
| [golyrics](https://github.com/mamal72/golyrics) | 33 | 3 | 2016-11-18 | 1 year ago | Golyrics is a Go library to fetch music lyrics data from the Wikia website. |
| [gosip](https://github.com/koltyakov/gosip) | 32 | 1 | 2019-01-26 | 1 day ago | Go client library SharePoint API. |
| [mixpanel](https://github.com/dukex/mixpanel) | 30 | 3 | 2014-05-20 | 6 days ago | Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications. |
| [gcm](https://github.com/TheOrioli/gcm) | 29 | 3 | 2015-11-09 | 4 years ago | Go library for Google Cloud Messaging. |
| [translate](https://github.com/nuveo/translate) | 29 | 27 | 2015-07-13 | 3 years ago | Go online translation package. |
| [gami](https://github.com/bit4bit/gami) | 27 | 4 | 2014-05-14 | 1 year ago | Go library for Asterisk Manager Interface. |
| [go-unsplash](https://github.com/hbagdi/go-unsplash) | 25 | 1 | 2017-01-19 | 1 month ago | Go client library for the [Unsplash.com](https://unsplash.com) API. |
| [go-spotify](https://github.com/rapito/go-spotify) | 21 | 1 | 2014-10-30 | 2 years ago | Go Library to access Spotify WEB API. |
| [ynab.go](https://github.com/brunomvsouza/ynab.go) | 21 | 1 | 2018-07-13 | 4 weeks ago | Go wrapper for the YNAB API. |
| [patreon-go](https://github.com/mxpv/patreon-go) | 19 | 4 | 2017-08-06 | 4 months ago | Go library for Patreon API. |
| [go-shopify](https://github.com/rapito/go-shopify) | 19 | 1 | 2014-10-28 | 2 years ago | Go Library to make CRUD request to the Shopify API. |
| [go-steam](https://github.com/sostronk/go-steam) | 19 | 10 | 2014-11-23 | 1 week ago | Go Library to interact with Steam game servers. |
| [go-twitch](https://github.com/knspriggs/go-twitch) | 18 | 5 | 2016-06-28 | 2 years ago | Go client for interacting with the Twitch v3 API. |
| [brewerydb](https://github.com/naegelejd/brewerydb) | 16 | 3 | 2015-04-15 | 4 years ago | Go library for accessing the BreweryDB API. |
| [codeship-go](https://github.com/codeship/codeship-go) | 16 | 23 | 2017-09-08 | 9 months ago | Go client library for interacting with Codeship's API v2. |
| [textbelt](https://github.com/farmergreg/textbelt) | 16 | 2 | 2015-09-01 | 4 years ago | Go client for the textbelt.com txt messaging API. |
| [go-myanimelist](https://github.com/nstratos/go-myanimelist) | 15 | 2 | 2015-05-03 | 2 years ago | Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api). |
| [coinpaprika-api-go-client](https://github.com/coinpaprika/coinpaprika-api-go-client) | 12 | 8 | 2018-09-25 | 10 months ago | Go client library for interacting with Coinpaprika's API. |
| [go-google-analytics](https://github.com/chonthu/go-google-analytics) | 11 | 2 | 2015-06-01 | 4 years ago | Simple wrapper for easy google analytics reporting. |
| [go-hacknews](https://github.com/PaulRosset/go-hacknews) | 10 | 2 | 2017-08-10 | 2 years ago | Tiny Go client for HackerNews API. |
| [smitego](https://github.com/sergiotapia/smitego) | 10 | 0 | 2013-12-11 | 5 years ago | Go package to wraps access to the Smite game API. |
| [lastpass-go](https://github.com/ansd/lastpass-go) | 9 | 0 | 2019-07-11 | 1 week ago | Go client library for the [LastPass](https://www.lastpass.com/) API. |
| [rrdaclient](https://github.com/Omie/rrdaclient) | 8 | 1 | 2014-09-15 | 5 years ago | Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP. |
| [go-sptrans](https://github.com/sergioaugrod/go-sptrans) | 7 | 2 | 2017-09-11 | 2 years ago | Go client library for the SPTrans Olho Vivo API. |
| [go-here](https://github.com/abdullahselek/go-here) | 6 | 1 | 2019-07-07 | 2 weeks ago | Go client library around the HERE location based APIs. |
| [go-google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) | 6 | 0 | 2016-10-24 | 3 years ago | Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/). |
| [gumblr](https://github.com/mattcunningham/gumblr) | 6 | 1 | 2015-07-09 | 3 years ago | Go wrapper for the Tumblr v2 API. |
| [go-sophos](https://github.com/esurdam/go-sophos) | 5 | 2 | 2018-09-05 | 2 months ago | Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies. |
| [go-zooz](https://github.com/gojuno/go-zooz) | 5 | 16 | 2017-07-04 | 1 year ago | Go client for the Zooz API. |
| [go-chronos](https://github.com/axelspringer/go-chronos) | 3 | 9 | 2017-10-23 | 2 years ago | Go library for interacting with the [Chronos](https://mesos.github. |
| [libgoffi](https://github.com/clevabit/libgoffi) | 1 | 2 | 2019-08-03 | 5 months ago | Library adapter toolbox for native [libffi](http://sourceware. |
| [playlyfe-go-sdk](https://github.com/playlyfe/playlyfe-go-sdk) | 1 | 4 | 2015-05-25 | 3 years ago | The Playlyfe Rest API Go SDK. |
| [tripadvisor-golang](https://github.com/mrbenosborne/tripadvisor-golang) | 1 | 1 | 2019-04-15 | 3 months ago | Go wrapper for the TripAdvisor API. |
| [vl-go](https://github.com/verifid/vl-go) | 1 | 1 | 2019-02-09 | 3 months ago | Go client library around the VerifID identity verification layer API. |

## Utilities
        
*General utilities and tools to make your life easier.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [fzf](https://github.com/junegunn/fzf) | 26517 | 354 | 2013-10-23 | 2 days ago | Command-line fuzzy finder written in Go. |
| [hub](https://github.com/github/hub) | 18371 | 390 | 2009-12-05 | 1 hour ago | wrap git commands with additional functionality to interact with github from the terminal. |
| [delve](https://github.com/go-delve/delve) | 13178 | 382 | 2014-05-20 | 2 hours ago | Go debugger. |
| [ctop](https://github.com/bcicen/ctop) | 9378 | 163 | 2016-12-27 | 2 weeks ago | [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics. |
| [wuzz](https://github.com/asciimoo/wuzz) | 8513 | 179 | 2017-01-30 | 2 weeks ago | Interactive cli tool for HTTP inspection. |
| [sqlx](https://github.com/jmoiron/sqlx) | 7687 | 178 | 2013-01-28 | 6 days ago | provides a set of extensions on top of the excellent built-in database/sql package. |
| [peco](https://github.com/peco/peco) | 5711 | 133 | 2014-06-06 | 1 week ago | Simplistic interactive filtering tool. |
| [usql](https://github.com/xo/usql) | 5556 | 118 | 2017-03-02 | 5 days ago | usql is a universal command-line interface for SQL databases. |
| [goreleaser](https://github.com/goreleaser/goreleaser) | 5034 | 77 | 2016-12-21 | 23 hours ago | Deliver Go binaries as fast and easily as possible. |
| [godropbox](https://github.com/dropbox/godropbox) | 3820 | 252 | 2014-06-22 | 8 months ago | Common libraries for writing Go services/applications from Dropbox. |
| [realize](https://github.com/oxequa/realize) | 3468 | 71 | 2016-07-12 | 1 month ago | Go build system with file watchers and live reload. Run, build and watch file changes with custom paths. |
| [goreporter](https://github.com/360EntSecGroup-Skylar/goreporter) | 2599 | 98 | 2017-03-27 | 1 year ago | Golang tool that does static analysis, unit testing, code review and generate code quality report. |
| [hystrix-go](https://github.com/afex/hystrix-go) | 2317 | 83 | 2013-12-15 | 7 months ago | Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker. |
| [task](https://github.com/go-task/task) | 2195 | 39 | 2017-02-27 | 1 day ago | simple "Make" alternative. |
| [panicparse](https://github.com/maruel/panicparse) | 2187 | 42 | 2015-02-02 | 4 months ago | Groups similar goroutines and colorizes stack dump. |
| [minify](https://github.com/tdewolff/minify) | 1982 | 47 | 2014-05-21 | 6 days ago | Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats. |
| [go-funk](https://github.com/thoas/go-funk) | 1539 | 38 | 2016-12-30 | 1 month ago | Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...). |
| [storm](https://github.com/asdine/storm) | 1474 | 46 | 2016-01-10 | 1 month ago | Simple and powerful toolkit for BoltDB. |
| [mmake](https://github.com/tj/mmake) | 1473 | 30 | 2017-02-15 | 1 week ago | Modern Make. |
| [mole](https://github.com/davrodpin/mole) | 1347 | 31 | 2018-10-04 | 3 months ago | cli app to easily create ssh tunnels. |
| [mc](https://github.com/minio/mc) | 1309 | 46 | 2015-01-16 | 16 hours ago | Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems. |
| [filetype](https://github.com/h2non/filetype) | 1029 | 24 | 2015-09-24 | 3 months ago | Small package to infer the file type checking the magic numbers signature. |
| [boilr](https://github.com/tmrts/boilr) | 1027 | 28 | 2015-12-19 | 3 months ago | Blazingly fast CLI tool for creating projects from boilerplate templates. |
| [mergo](https://github.com/imdario/mergo) | 1013 | 17 | 2013-03-11 | 1 day ago | Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements. |
| [spinner](https://github.com/briandowns/spinner) | 903 | 13 | 2014-12-13 | 2 months ago | Go package to easily provide a terminal spinner with options. |
| [circuitbreaker](https://github.com/rubyist/circuitbreaker) | 843 | 18 | 2014-07-17 | 3 months ago | Circuit Breakers in Go. |
| [gtm](https://github.com/git-time-metric/gtm) | 768 | 28 | 2016-06-19 | 5 months ago | Simple, seamless, lightweight time tracking for Git. |
| [jump](https://github.com/gsamokovarov/jump) | 767 | 13 | 2015-08-16 | 1 month ago | Jump helps you navigate faster by learning your habits. |
| [immortal](https://github.com/immortal/immortal) | 632 | 14 | 2016-06-30 | 1 month ago | \*nix cross-platform (OS agnostic) supervisor. |
| [htcat](https://github.com/htcat/htcat) | 503 | 17 | 2013-08-05 | 10 months ago | Parallel and Pipelined HTTP GET Utility. |
| [go-dry](https://github.com/ungerik/go-dry) | 444 | 14 | 2014-02-28 | 1 year ago | DRY (don't repeat yourself) package for Go. |
| [gopencils](https://github.com/bndr/gopencils) | 428 | 14 | 2014-06-23 | 11 months ago | Small and simple package to easily consume REST APIs. |
| [godaemon](https://github.com/VividCortex/godaemon) | 423 | 31 | 2013-08-01 | 10 months ago | Utility to write daemons. |
| [circuit](https://github.com/cep21/circuit) | 421 | 11 | 2017-12-23 | 3 months ago | An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern. |
| [request](https://github.com/mozillazg/request) | 370 | 14 | 2014-12-21 | 1 month ago | Go HTTP Requests for Humans™. |
| [koazee](https://github.com/wesovilabs/koazee) | 357 | 11 | 2018-11-09 | 2 months ago | Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays. |
| [ergo](https://github.com/cristianoliveira/ergo) | 354 | 5 | 2017-08-19 | 4 months ago | The management of multiple local services running over different ports made easy. |
| [go-rate](https://github.com/beefsack/go-rate) | 295 | 10 | 2014-08-25 | 1 year ago | Timed rate limiter for Go. |
| [gohper](https://github.com/zhuah/gohper) | 248 | 20 | 2015-03-23 | 2 years ago | Various tools/modules help for development. |
| [clockwork](https://github.com/jonboulle/clockwork) | 245 | 4 | 2014-09-09 | 2 days ago | A simple fake clock for golang. |
| [deepcopier](https://github.com/ulule/deepcopier) | 234 | 17 | 2015-07-24 | 4 days ago | Simple struct copying for Go. |
| [serve](https://github.com/syntaqx/serve) | 203 | 5 | 2019-01-10 | 2 weeks ago | A static http server anywhere you need. |
| [gubrak](https://github.com/novalagung/gubrak) | 198 | 6 | 2018-03-09 | 3 hours ago | Golang utility library with syntactic sugar. It's like lodash, but for golang. |
| [retry](https://github.com/kamilsk/retry) | 192 | 2 | 2016-11-02 | 2 weeks ago | The most advanced functional mechanism to perform actions repetitively until successful. |
| [go-trigger](https://github.com/sadlil/go-trigger) | 190 | 12 | 2015-10-19 | 2 years ago | Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project. |
| [mimetype](https://github.com/gabriel-vasile/mimetype) | 189 | 5 | 2018-07-02 | 7 hours ago | Package for MIME type detection based on magic numbers. |
| [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) | 164 | 3 | 2016-11-08 | 4 months ago | go:generate tool for wrapping symbols exported by golang plugins (1.8 only). |
| [gotenv](https://github.com/subosito/gotenv) | 162 | 4 | 2013-08-27 | 4 months ago | Load environment variables from `.env` or any `io.Reader` in Go. |
| [rerun](https://github.com/ivpusic/rerun) | 155 | 7 | 2014-12-10 | 1 year ago | Recompiling and rerunning go apps when source changes. |
| [util](https://github.com/shomali11/util) | 151 | 9 | 2017-05-24 | 7 months ago | Collection of useful utility functions. (strings, concurrency, manipulations, ...). |
| [death](https://github.com/vrecan/death) | 148 | 4 | 2015-03-09 | 1 year ago | Managing go application shutdown with signals. |
| [moldova](https://github.com/StabbyCutyou/moldova) | 148 | 5 | 2016-01-30 | 2 years ago | Utility for generating random data based on an input template. |
| [robustly](https://github.com/VividCortex/robustly) | 140 | 18 | 2013-07-08 | 1 year ago | Runs functions resiliently, catching and restarting panics. |
| [apm](https://github.com/topfreegames/apm) | 133 | 16 | 2015-11-18 | 3 years ago | Process manager for Golang applications with an HTTP API. |
| [chyle](https://github.com/antham/chyle) | 116 | 6 | 2016-11-17 | 8 hours ago | Changelog generator using a git repository with multiple configuration possibilities. |
| [toolbox](https://github.com/viant/toolbox) | 116 | 14 | 2016-06-13 | 20 hours ago | Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer. |
| [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) | 115 | 5 | 2015-10-12 | 4 months ago | XML Sitemap generator written in Go. |
| [onecache](https://github.com/adelowo/onecache) | 106 | 7 | 2017-04-14 | 8 months ago | Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc). |
| [lrserver](https://github.com/jaschaephraim/lrserver) | 105 | 5 | 2014-07-15 | 2 years ago | LiveReload server for Go. |
| [go-bsdiff](https://github.com/gabstv/go-bsdiff) | 90 | 1 | 2019-02-23 | 10 months ago | Pure Go bsdiff and bspatch libraries and CLI tools. |
| [pm](https://github.com/VividCortex/pm) | 75 | 18 | 2013-11-17 | 2 months ago | Process (i.e. goroutine) manager with an HTTP API. |
| [go-health](https://github.com/Talento90/go-health) | 72 | 2 | 2018-02-13 | 1 year ago | Health package simplifies the way you add health check to your services. |
| [xferspdy](https://github.com/monmohan/xferspdy) | 71 | 3 | 2015-05-22 | 3 years ago | Xferspdy provides binary diff and patch library in golang. |
| [unis](https://github.com/esemplastic/unis) | 69 | 4 | 2017-05-06 | 2 years ago | Common Architecture™ for String Utilities in Go. |
| [mssqlx](https://github.com/linxGnu/mssqlx) | 68 | 9 | 2016-12-26 | 2 months ago | Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind. |
| [multitick](https://github.com/VividCortex/multitick) | 61 | 18 | 2013-12-10 | 2 months ago | Multiplexor for aligned tickers. |
| [repeat](https://github.com/ssgreg/repeat) | 61 | 4 | 2017-11-22 | 2 months ago | Go implementation of different backoff strategies useful for retrying operations and heartbeating. |
| [minquery](https://github.com/icza/minquery) | 52 | 3 | 2016-11-16 | 2 months ago | MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off). |
| [mimemagic](https://github.com/zRedShift/mimemagic) | 51 | 1 | 2018-10-11 | 1 year ago | Pure Go ultra performant MIME sniffing library/utility. |
| [handy](https://github.com/miguelpragier/handy) | 49 | 5 | 2018-06-13 | 6 days ago | Many utilities and helpers like string handlers/formatters and validators. |
| [golog](https://github.com/mlimaloureiro/golog) | 47 | 3 | 2016-01-09 | 1 year ago | Easy and lightweight CLI tool to time track your tasks. |
| [go-astitodo](https://github.com/asticode/go-astitodo) | 46 | 2 | 2016-10-17 | 1 year ago | Parse TODOs in your GO code. |
| [goreadability](https://github.com/philipjkim/goreadability) | 44 | 6 | 2016-04-20 | 9 months ago | Webpage summary extractor using Facebook Open Graph and arc90's readability. |
| [goseaweedfs](https://github.com/linxGnu/goseaweedfs) | 43 | 6 | 2017-07-20 | 1 month ago | SeaweedFS client library with almost full features. |
| [goback](https://github.com/carlescere/goback) | 42 | 1 | 2015-03-13 | 1 year ago | Go simple exponential backoff package. |
| [intrinsic](https://github.com/mengzhuo/intrinsic) | 42 | 3 | 2017-06-13 | 2 years ago | Use x86 SIMD without writing any assembly code. |
| [gaper](https://github.com/maxcnunes/gaper) | 41 | 0 | 2018-06-16 | 1 month ago | Builds and restarts a Go project when it crashes or some watched file changes. |
| [copy-pasta](https://github.com/jutkko/copy-pasta) | 38 | 4 | 2017-01-28 | 5 months ago | Universal multi-workstation clipboard that uses S3 like backend for the storage. |
| [golarm](https://github.com/msempere/golarm) | 37 | 1 | 2015-08-14 | 4 years ago | Fire alarms with system events. |
| [retry](https://github.com/thedevsaddam/retry) | 35 | 1 | 2018-02-25 | 1 year ago | Simple and easy retry mechanism package for Go. |
| [pgo](https://github.com/arthurkushman/pgo) | 34 | 5 | 2018-12-26 | 2 days ago | Convenient functions for PHP community. |
| [myhttp](https://github.com/inancgumus/myhttp) | 33 | 1 | 2017-09-13 | 1 year ago | Simple API to make HTTP GET requests with timeout support. |
| [gpath](https://github.com/tenntenn/gpath) | 32 | 2 | 2017-05-24 | 2 years ago | Library to simplify access struct fields with Go's expression in reflection. |
| [rclient](https://github.com/zpatrick/rclient) | 31 | 3 | 2017-02-28 | 1 month ago | Readable, flexible, simple-to-use client for REST APIs. |
| [retry-go](https://github.com/rafaeljesus/retry-go) | 31 | 1 | 2017-06-09 | 1 year ago | Retrying made simple and easy for golang. |
| [beyond](https://github.com/wesovilabs/beyond) | 30 | 1 | 2019-10-18 | 1 month ago | The Go tool that will drive you to the AOP world! |
| [dbt](https://github.com/nikogura/dbt) | 24 | 2 | 2017-11-30 | 22 hours ago | A framework for running self-updating signed binaries from a central, trusted repository. |
| [gostrutils](https://github.com/ik5/gostrutils) | 22 | 3 | 2018-09-19 | 1 month ago | Collections of string manipulation and conversion functions. |
| [scan](https://github.com/blockloop/scan) | 22 | 1 | 2017-11-27 | 5 days ago | Scan golang `sql.Rows` directly to structs, slices, or primitive types. |
| [ugo](https://github.com/alxrm/ugo) | 22 | 1 | 2016-02-17 | 3 years ago | ugo is slice toolbox with concise syntax for Go. |
| [generate](https://github.com/go-playground/generate) | 21 | 2 | 2015-11-15 | 3 years ago | runs go generate recursively on a specified path or environment variable and can filter by regex. |
| [goplaceholder](https://github.com/michiwend/goplaceholder) | 21 | 2 | 2014-10-12 | 4 years ago | a small golang lib to generate placeholder images. |
| [evaluator](https://github.com/nullne/evaluator) | 20 | 1 | 2017-04-27 | 2 years ago | Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend. |
| [filter](https://github.com/gookit/filter) | 20 | 4 | 2018-09-26 | 1 month ago | provide filtering, sanitizing, and conversion of Go data. |
| [cmd](https://github.com/SimonBaeumer/cmd) | 16 | 2 | 2019-09-27 | 6 days ago | Library for executing shell commands on osx, windows and linux. |
| [slice](https://github.com/psampaz/slice) | 16 | 1 | 2019-11-26 | 1 day ago | Type-safe functions for common Go slice operations. |
| [dlog](https://github.com/kirillDanshin/dlog) | 15 | 2 | 2016-07-04 | 2 years ago | Compile-time controlled logger to make your release smaller without removing debug calls. |
| [go-httpheader](https://github.com/mozillazg/go-httpheader) | 15 | 2 | 2017-06-24 | 1 year ago | Go library for encoding structs into Header fields. |
| [filler](https://github.com/yaronsumel/filler) | 14 | 1 | 2017-04-05 | 2 years ago | small utility to fill structs using "fill" tag. |
| [okrun](https://github.com/xta/okrun) | 14 | 2 | 2014-10-01 | 5 years ago | go run error steamroller. |
| [rerate](https://github.com/abo/rerate) | 14 | 4 | 2016-05-24 | 2 years ago | Redis-based rate counter and rate limiter for Go. |
| [structs](https://github.com/PumpkinSeed/structs) | 14 | 2 | 2017-08-26 | 2 years ago | Implement simple functions to manipulate structs. |
| [ghokin](https://github.com/antham/ghokin) | 13 | 2 | 2018-08-03 | 3 days ago | Parallelized formatter with no external dependencies for gherkin (cucumber, behat...). |
| [rest-go](https://github.com/edermanoel94/rest-go) | 12 | 3 | 2019-07-29 | 3 months ago | A package that provide many helpful methods for working with rest api. |
| [slicer](https://github.com/leaanthony/slicer) | 12 | 1 | 2019-01-10 | 1 week ago | Makes working with slices easier. |
| [ctxutil](https://github.com/posener/ctxutil) | 11 | 1 | 2018-07-30 | 5 months ago | A collection of utility functions for contexts. |
| [mimesniffer](https://github.com/aofei/mimesniffer) | 11 | 1 | 2018-12-20 | 4 months ago | A MIME type sniffer for Go. |
| [retry](https://github.com/shafreeck/retry) | 10 | 0 | 2018-07-18 | 6 months ago | A pretty simple library to ensure your work to be done. |
| [backscanner](https://github.com/icza/backscanner) | 9 | 2 | 2017-10-18 | 1 year ago | A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward. |
| [command](https://github.com/txgruppi/command) | 9 | 2 | 2015-08-24 | 3 years ago | Command pattern for Go with thread safe serial and parallel dispatcher. |
| [tome](https://github.com/cyruzin/tome) | 9 | 1 | 2019-04-12 | 3 months ago | Tome was designed to paginate simple RESTful APIs. |
| [limiters](https://github.com/mennanov/limiters) | 8 | 1 | 2019-08-28 | 2 months ago | Rate limiters for distributed applications in Golang with configurable back-ends and distributed locks. |
| [shutdown](https://github.com/ztrue/shutdown) | 8 | 1 | 2018-11-17 | 11 months ago | App shutdown hooks for `os.Signal` handling. |
| [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) | 5 | 1 | 2019-05-16 | 8 months ago | Go package for working with Problem Details. |
| [retry](https://github.com/percolate/retry) | 5 | 29 | 2018-06-15 | 4 months ago | A simple but highly configurable retry package for Go. |
| [blank](https://github.com/Henry-Sarabia/blank) | 4 | 2 | 2019-02-13 | 5 months ago | Verify or remove blanks and whitespace from strings. |
| [silk](https://github.com/chrispassas/silk) | 4 | 1 | 2018-12-18 | 5 months ago | Read silk netflow files. |
| [sliceconv](https://github.com/Henry-Sarabia/sliceconv) | 4 | 1 | 2019-02-15 | 5 months ago | Slice conversion between primitive types. |
| [ptr](https://github.com/gotidy/ptr) | 3 | 1 | 2019-12-25 | 3 weeks ago | Package that provide functions for simplified creation of pointers from constants of basic types. |
| [olaf](https://github.com/btnguyen2k/olaf) | 1 | 1 | 2019-01-03 | 9 months ago | Twitter Snowflake implemented in Go. |
| [jsend](https://github.com/clevergo/jsend) | 1 | 1 | 2020-01-14 | 1 week ago | JSend's implementation writen in Go. |

## UUID
        
*Libraries for working with UUIDs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [ulid](https://github.com/oklog/ulid) | 1799 | 36 | 2016-12-06 | 4 months ago | Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier). |
| [uuid](https://github.com/gofrs/uuid) | 642 | 16 | 2018-07-13 | 3 months ago | Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid. |
| [wuid](https://github.com/edwingeng/wuid) | 323 | 16 | 2018-01-27 | 1 month ago | An extremely fast unique number generator, 10-135 times faster than UUID. |
| [Goid](https://github.com/JakeHL/Goid) | 26 | 4 | 2017-05-19 | 11 months ago | Generate and Parse RFC4122 compliant V4 UUIDs. |
| [sno](https://github.com/muyo/sno) | 23 | 1 | 2019-05-26 | 1 month ago | Compact, sortable and fast unique IDs with embedded metadata. |
| [nanoid](https://github.com/aidarkhanov/nanoid) | 16 | 1 | 2019-07-02 | 2 weeks ago | A tiny and efficient Go unique string ID generator. |
| [uuid](https://github.com/agext/uuid) | 10 | 1 | 2016-02-03 | 11 months ago | Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier. |

## Validation
        
*Libraries for validation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [validator](https://github.com/go-playground/validator) | 4611 | 82 | 2015-02-12 | 3 days ago | Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving. |
| [govalidator](https://github.com/asaskevich/govalidator) | 3986 | 100 | 2014-06-20 | 1 week ago | Validators and sanitizers for strings, numerics, slices and structs. |
| [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) | 1237 | 23 | 2016-06-22 | 1 hour ago | Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags. |
| [govalidator](https://github.com/thedevsaddam/govalidator) | 816 | 20 | 2017-09-13 | 2 months ago | Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. |
| [validate](https://github.com/gookit/validate) | 195 | 8 | 2018-07-16 | 1 month ago | Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features. |
| [checkdigit](https://github.com/osamingo/checkdigit) | 49 | 0 | 2019-04-05 | 3 weeks ago | Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.). |
| [jio](https://github.com/faceair/jio) | 30 | 2 | 2018-10-28 | 7 months ago | jio is a json schema validator similar to [joi](https://github.com/hapijs/joi). |
| [validate](https://github.com/gobuffalo/validate) | 27 | 7 | 2018-02-10 | 4 days ago | This package provides a framework for writing validations for Go applications. |
| [terraform-validator](https://github.com/thazelart/terraform-validator) | 25 | 2 | 2019-05-29 | 1 week ago | A norms and conventions validator for Terraform. |

## Version Control
        
*Libraries for version control.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-git](https://github.com/src-d/go-git) | 4870 | 103 | 2015-10-22 | 1 week ago | highly extensible Git implementation in pure Go. |
| [git2go](https://github.com/libgit2/git2go) | 1428 | 54 | 2013-03-05 | 38 minutes ago | Go bindings for libgit2. |
| [hercules](https://github.com/src-d/hercules) | 825 | 21 | 2016-12-12 | 6 days ago | gaining advanced insights from Git repository history. |
| [gh](https://github.com/rjeczalik/gh) | 72 | 6 | 2015-03-08 | 1 year ago | Scriptable server and net/http middleware for GitHub Webhooks. |
| [go-vcs](https://github.com/sourcegraph/go-vcs) | 72 | 38 | 2013-06-02 | 4 months ago | manipulate and inspect VCS repositories in Go. |
| [hgo](https://github.com/beyang/hgo) | 12 | 4 | 2014-06-18 | 4 years ago | Hgo is a collection of Go packages providing read-access to local Mercurial repositories. |

## Video
        
*Libraries for manipulating video.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [goav](https://github.com/giorgisio/goav) | 987 | 42 | 2015-05-21 | 1 day ago | Comphrensive Go bindings for FFmpeg. |
| [m3u8](https://github.com/grafov/m3u8) | 654 | 34 | 2013-02-05 | 6 days ago | Parser and generator library of M3U8 playlists for Apple HLS. |
| [gmf](https://github.com/3d0c/gmf) | 576 | 31 | 2013-04-03 | 6 months ago | Go bindings for FFmpeg av\* libraries. |
| [go-astits](https://github.com/asticode/go-astits) | 286 | 15 | 2017-07-04 | 1 week ago | Parse and demux MPEG Transport Streams (.ts) natively in GO. |
| [go-astisub](https://github.com/asticode/go-astisub) | 219 | 7 | 2016-12-16 | 1 week ago | Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.). |
| [gst](https://github.com/ziutek/gst) | 155 | 10 | 2011-07-26 | 1 year ago | Go bindings for GStreamer. |
| [libvlc-go](https://github.com/adrg/libvlc-go) | 83 | 8 | 2015-01-06 | 2 months ago | Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player). |
| [go-m3u8](https://github.com/quangngotan95/go-m3u8) | 44 | 1 | 2018-11-06 | 9 months ago | Parser and generator library for Apple m3u8 playlists. |
| [v4l](https://github.com/korandiz/v4l) | 36 | 4 | 2016-10-25 | 1 year ago | Video capture library for Linux, written in Go. |
| [libgosubs](https://github.com/wargarblgarbl/libgosubs) | 12 | 1 | 2017-05-03 | 1 year ago | Subtitle format support for go. Supports .srt, .ttml, and .ass. |

## Web Frameworks
        
*Full stack web frameworks.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gin](https://github.com/gin-gonic/gin) | 34931 | 1208 | 2014-06-16 | 12 minutes ago | Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity. |
| [beego](https://github.com/astaxie/beego) | 23084 | 1267 | 2012-02-29 | 4 days ago | beego is an open-source, high-performance web framework for the Go programming language. |
| [echo](https://github.com/labstack/echo) | 16267 | 545 | 2015-03-01 | 1 day ago | High performance, minimalist Go web framework. |
| [revel](https://github.com/revel/revel) | 11539 | 561 | 2011-12-09 | 6 months ago | High-productivity web framework for the Go language. |
| [goa](https://github.com/goadesign/goa) | 3690 | 169 | 2014-12-05 | 2 hours ago | Goa provides a holistic approach for developing remote APIs and microservices in Go. |
| [go-json-rest](https://github.com/ant0ine/go-json-rest) | 3387 | 163 | 2013-02-19 | 4 months ago | Quick and easy way to setup a RESTful JSON API. |
| [gizmo](https://github.com/nytimes/gizmo) | 3026 | 114 | 2015-12-15 | 1 month ago | Microservice toolkit used by the New York Times. |
| [macaron](https://github.com/go-macaron/macaron) | 2927 | 147 | 2014-07-10 | 3 months ago | Macaron is a high productive and modular design web framework in Go. |
| [utron](https://github.com/gernest/utron) | 2154 | 71 | 2015-09-16 | 1 year ago | Lightweight MVC framework for Go(Golang). |
| [go-tigertonic](https://github.com/rcrowley/go-tigertonic) | 1000 | 46 | 2013-02-09 | 1 year ago | Go framework for building JSON web services inspired by Dropwizard. |
| [tango](https://github.com/lunny/tango) | 833 | 77 | 2014-12-17 | 8 months ago | Micro & pluggable web framework for Go. |
| [gongular](https://github.com/mustafaakin/gongular) | 421 | 21 | 2016-06-22 | 11 months ago | Fast Go web framework with input mapping/validation and (DI) Dependency Injection. |
| [neo](https://github.com/ivpusic/neo) | 402 | 32 | 2015-02-04 | 2 years ago | Neo is minimal and fast Go Web Framework with extremely simple API. |
| [air](https://github.com/aofei/air) | 363 | 17 | 2016-07-20 | 3 weeks ago | An ideally refined web framework for Go. |
| [mango](https://github.com/paulbellamy/mango) | 344 | 22 | 2011-05-25 | 2 years ago | Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. |
| [gondola](https://github.com/rainycape/gondola) | 314 | 15 | 2014-07-25 | 11 months ago | The web framework for writing faster sites, faster. |
| [golf](https://github.com/dinever/golf) | 239 | 19 | 2015-11-18 | 2 years ago | Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library. |
| [aero](https://github.com/aerogo/aero) | 228 | 15 | 2016-11-09 | 1 month ago | High-performance web framework for Go, reaches top scores in Lighthouse. |
| [go-rest](https://github.com/ungerik/go-rest) | 116 | 10 | 2012-07-13 | 3 years ago | Small and evil REST framework for Go. |
| [hiboot](https://github.com/hidevopsio/hiboot) | 110 | 10 | 2018-03-16 | 1 week ago | hiboot is a high performance web application framework with auto configuration and dependency injection support. |
| [webgo](https://github.com/bnkamalesh/webgo) | 85 | 3 | 2015-12-16 | 3 days ago | A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc). |
| [flamingo](https://github.com/i-love-flamingo/flamingo) | 77 | 19 | 2019-04-02 | 4 days ago | Framework for pluggable web projects. Including a concept for modules and offering features for DI, Configareas, i18n, template engines, graphql, observability, security, events, routing & reverse routing etc. |
| [uadmin](https://github.com/uadmin/uadmin) | 74 | 7 | 2018-10-05 | 1 week ago | Fully featured web framework for Golang, inspired by Django. |
| [golax](https://github.com/fulldump/golax) | 71 | 8 | 2016-01-30 | 1 year ago | A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more. |
| [microservice](https://github.com/claygod/microservice) | 69 | 8 | 2016-12-15 | 7 months ago | The framework for the creation of microservices, written in Golang. |
| [yarf](https://github.com/yarf-framework/yarf) | 54 | 3 | 2015-09-02 | 10 months ago | Fast micro-framework designed to build REST APIs and web services in a fast and simple way. |
| [ginrpc](https://github.com/xxjwxc/ginrpc) | 52 | 2 | 2019-06-22 | 2 weeks ago | Gin parameter automatic binding tool,gin rpc tools. |
| [fireball](https://github.com/zpatrick/fireball) | 50 | 4 | 2016-07-20 | 1 year ago | More "natural" feeling web framework. |
| [vox](https://github.com/aisk/vox) | 47 | 2 | 2014-12-24 | 6 days ago | A golang web framework for humans, inspired by Koa heavily. |
| [patron](https://github.com/beatlabs/patron) | 46 | 16 | 2019-01-30 | 8 hours ago | Patron is a microservice framework following best cloud practices with a focus on productivity. |
| [flamingo-commerce](https://github.com/i-love-flamingo/flamingo-commerce) | 41 | 16 | 2019-04-02 | 47 minutes ago | Providing e-commerce features using clean architecture like DDD and ports and adapters, that you can use to build flexible e-commerce applications. |
| [api](https://github.com/resoursea/api) | 30 | 5 | 2015-01-24 | 5 years ago | REST framework for quickly writing resource based services. |
| [rex](https://github.com/goanywhere/rex) | 29 | 1 | 2014-10-16 | 2 years ago | Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`. |
| [goa](https://github.com/goa-go/goa) | 27 | 2 | 2019-07-26 | 1 month ago | goa is just like koajs for golang, it is a flexible, light, high-performance and extensible web framework based on middleware. |
| [rux](https://github.com/gookit/rux) | 19 | 2 | 2018-08-05 | 2 weeks ago | Simple and fast web framework for build golang HTTP applications. |
| [banjo](https://github.com/nsheremet/banjo) | 11 | 1 | 2017-12-09 | 2 years ago | Very simple and fast web framework for Go. |

### Middlewares
        

#### Actual middlewares
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [tollbooth](https://github.com/didip/tollbooth) | 1455 | 45 | 2015-05-17 | 4 months ago | Rate limit HTTP request handler. |
| [cors](https://github.com/rs/cors) | 1365 | 30 | 2014-10-25 | 1 month ago | Easily add CORS capabilities to your API. |
| [limiter](https://github.com/ulule/limiter) | 861 | 25 | 2015-10-02 | 6 days ago | Dead simple rate limit middleware for Go. |
| [go-server-timing](https://github.com/mitchellh/go-server-timing) | 767 | 20 | 2018-02-12 | 1 year ago | Add/parse Server-Timing header. |
| [ln-paywall](https://github.com/philippgille/ln-paywall) | 98 | 4 | 2018-06-29 | 11 months ago | Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin). |
| [xff](https://github.com/sebest/xff) | 72 | 2 | 2014-12-22 | 9 months ago | Handle `X-Forwarded-For` header and friends. |
| [formjson](https://github.com/rs/formjson) | 34 | 2 | 2015-03-19 | 4 years ago | Transparently handle JSON input as a standard form POST. |
| [client-timing](https://github.com/posener/client-timing) | 16 | 1 | 2018-02-23 | 10 months ago | An HTTP client for Server-Timing header. |

#### Libraries for creating HTTP middlewares
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [negroni](https://github.com/urfave/negroni) | 6516 | 238 | 2014-05-18 | 3 months ago | Idiomatic HTTP middleware for Golang. |
| [alice](https://github.com/justinas/alice) | 1939 | 52 | 2014-05-25 | 2 months ago | Painless middleware chaining for Go. |
| [render](https://github.com/unrolled/render) | 1328 | 41 | 2014-06-10 | 1 month ago | Go package for easily rendering JSON, XML, and HTML template responses. |
| [stats](https://github.com/thoas/stats) | 550 | 16 | 2015-03-05 | 9 months ago | Go middleware that stores various information about your web application. |
| [interpose](https://github.com/carbocation/interpose) | 288 | 12 | 2014-07-20 | 3 years ago | Minimalist net/http middleware for golang. |
| [muxchain](https://github.com/stephens2424/muxchain) | 209 | 5 | 2014-05-03 | 10 months ago | Lightweight middleware for net/http. |
| [renderer](https://github.com/thedevsaddam/renderer) | 181 | 6 | 2017-11-07 | 10 months ago | Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go. |
| [rye](https://github.com/InVisionApp/rye) | 93 | 186 | 2016-10-06 | 1 year ago | Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context. |
| [gores](https://github.com/alioygur/gores) | 88 | 5 | 2015-12-25 | 1 year ago | Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs. |
| [chain](https://github.com/codemodus/chain) | 64 | 6 | 2015-05-14 | 1 year ago | Handler wrapper chaining with scoped data (net/context-based "middleware"). |
| [wrap](https://github.com/go-on/wrap) | 59 | 3 | 2014-02-16 | 1 year ago | Small middlewares package for net/http. |
| [catena](https://github.com/codemodus/catena) | 7 | 1 | 2015-07-30 | 1 year ago | http.Handler wrapper catenation (same API as "chain"). |

### Routers
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [mux](https://github.com/gorilla/mux) | 11023 | 282 | 2012-10-02 | 1 week ago | Powerful URL router and dispatcher for golang. |
| [httprouter](https://github.com/julienschmidt/httprouter) | 10691 | 306 | 2013-12-05 | 1 week ago | High performance router. Use this and the standard http handlers to form a very high performance web framework. |
| [chi](https://github.com/go-chi/chi) | 6898 | 179 | 2015-10-15 | 1 week ago | Small, fast and expressive HTTP router built on net/context. |
| [web](https://github.com/gocraft/web) | 1406 | 58 | 2013-11-16 | 6 months ago | Mux and middleware package in Go. |
| [bone](https://github.com/go-zoo/bone) | 1245 | 36 | 2014-11-19 | 8 months ago | Lightning Fast HTTP Multiplexer. |
| [fasthttprouter](https://github.com/buaazp/fasthttprouter) | 812 | 34 | 2015-12-13 | 9 months ago | High performance router forked from `httprouter`. The first router fit for `fasthttp`. |
| [goji](https://github.com/goji/goji) | 794 | 40 | 2015-11-16 | 5 months ago | Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`. |
| [gorouter](https://github.com/xujiajun/gorouter) | 473 | 15 | 2018-01-29 | 3 months ago | A simple and fast HTTP router for Go. |
| [httptreemux](https://github.com/dimfeld/httptreemux) | 407 | 23 | 2014-05-14 | 3 months ago | High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter. |
| [lars](https://github.com/go-playground/lars) | 380 | 15 | 2015-12-24 | 8 months ago | Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks. |
| [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) | 371 | 28 | 2015-10-27 | 1 week ago | An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs. |
| [siesta](https://github.com/VividCortex/siesta) | 352 | 27 | 2014-09-23 | 1 month ago | Composable framework to write middleware and handlers. |
| [vestigo](https://github.com/husobee/vestigo) | 258 | 16 | 2015-09-22 | 3 months ago | Performant, stand-alone, HTTP compliant URL Router for go web applications. |
| [router](https://github.com/gowww/router) | 159 | 6 | 2017-05-25 | 1 year ago | Lightning fast HTTP router fully compatible with the net/http.Handler interface. |
| [alien](https://github.com/gernest/alien) | 109 | 4 | 2016-01-30 | 10 months ago | Lightweight and fast http router from outer space. |
| [violetear](https://github.com/nbari/violetear) | 96 | 3 | 2015-06-19 | 2 months ago | Go HTTP router. |
| [Bxog](https://github.com/claygod/Bxog) | 92 | 8 | 2016-05-19 | 2 days ago | Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters. |
| [pure](https://github.com/go-playground/pure) | 91 | 5 | 2016-09-23 | 2 months ago | Is a lightweight HTTP router that sticks to the std "net/http" implementation. |
| [xmux](https://github.com/rs/xmux) | 87 | 6 | 2015-12-14 | 2 years ago | High performance muxer based on `httprouter` with `net/context` support. |
| [gorouter](https://github.com/vardius/gorouter) | 59 | 4 | 2016-07-14 | 5 hours ago | GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`. |
| [bellt](https://github.com/GuilhermeCaruso/bellt) | 40 | 5 | 2019-02-21 | 8 months ago | A simple Go HTTP router. |
| [fastrouter](https://github.com/razonyang/fastrouter) | 19 | 2 | 2017-11-01 | 2 years ago | a fast, flexible HTTP router written in Go. |
| [route](https://github.com/goroute/route) | 6 | 3 | 2019-07-06 | 4 weeks ago | Simple yet powerful HTTP request multiplexer. |

## Windows
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-ole](https://github.com/go-ole/go-ole) | 599 | 38 | 2011-01-21 | 4 months ago | Win32 OLE implementation for golang. |
| [d3d9](https://github.com/gonutz/d3d9) | 96 | 6 | 2015-12-12 | 2 months ago | Go bindings for Direct3D9. |
| [gosddl](https://github.com/MonaxGT/gosddl) | 1 | 1 | 2018-12-04 | 8 months ago | Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL. |

## XML
        
*Libraries and tools for manipulating XML.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [zek](https://github.com/miku/zek) | 316 | 19 | 2017-11-23 | 6 months ago | Generate a Go struct from XML. |
| [xpath](https://github.com/antchfx/xpath) | 261 | 9 | 2016-10-09 | 1 week ago | XPath package for Go. |
| [xquery](https://github.com/antchfx/xquery) | 153 | 11 | 2016-10-09 | 1 year ago | XQuery lets you extract data from HTML/XML documents using XPath expression. |
| [xml2map](https://github.com/sbabiv/xml2map) | 19 | 1 | 2018-08-06 | 9 months ago | XML to MAP converter written Golang. |
| [XML-Comp](https://github.com/XML-Comp/XML-Comp) | 16 | 1 | 2016-10-25 | 1 year ago | Simple command line XML comparer that generates diffs of folders, files and tags. |
| [xmlwriter](https://github.com/shabbyrobe/xmlwriter) | 8 | 1 | 2017-04-11 | 1 week ago | Procedural XML generation API based on libxml2's xmlwriter module. |

# Tools
        
*Go software and plugins.*

## Code Analysis
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [lint](https://github.com/golang/lint) | 3351 | 112 | 2013-06-02 | 3 weeks ago | Golint is a linter for Go source code. |
| [errcheck](https://github.com/kisielk/errcheck) | 1390 | 23 | 2013-02-24 | 4 days ago | Errcheck is a program for checking for unchecked errors in Go programs. |
| [gcvis](https://github.com/davecheney/gcvis) | 953 | 35 | 2014-07-10 | 10 months ago | Visualise Go program GC trace data in real time. |
| [php-parser](https://github.com/z7zmey/php-parser) | 690 | 27 | 2017-11-07 | 1 week ago | A Parser for PHP written in Go. |
| [go-critic](https://github.com/go-critic/go-critic) | 656 | 21 | 2018-05-05 | 20 hours ago | source code linter that brings checks that are currently not implemented in other linters. |
| [goast-viewer](https://github.com/yuroyoro/goast-viewer) | 419 | 14 | 2014-06-30 | 7 months ago | Web based Golang AST visualizer. |
| [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) | 306 | 9 | 2017-04-12 | 3 weeks ago | go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects. |
| [unconvert](https://github.com/mdempsky/unconvert) | 277 | 8 | 2016-02-19 | 4 months ago | Remove unnecessary type conversions from Go source. |
| [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) | 253 | 5 | 2019-04-19 | 1 day ago | An easy way to find outdated dependencies of your Go projects. |
| [gostatus](https://github.com/shurcooL/gostatus) | 242 | 7 | 2013-11-27 | 11 months ago | Command line tool, shows the status of repositories that contain Go packages. |
| [dupl](https://github.com/mibk/dupl) | 187 | 8 | 2015-05-20 | 1 year ago | Tool for code clone detection. |
| [apicompat](https://github.com/bradleyfalzon/apicompat) | 170 | 8 | 2016-07-10 | 3 years ago | Checks recent changes to a Go project for backwards incompatible changes. |
| [tickgit](https://github.com/augmentable-dev/tickgit) | 104 | 4 | 2019-10-12 | 2 days ago | CLI and go package for surfacing code comment TODOs (in any language) and applying a `git blame`to identify the author. |
| [checkstyle](https://github.com/qiniu/checkstyle) | 103 | 11 | 2014-01-01 | 1 month ago | checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments. |
| [goplantuml](https://github.com/jfeliu007/goplantuml) | 86 | 6 | 2019-05-26 | 1 month ago | Library and CLI that generates text plantump class diagram containing information about structures and interfaces with the relationship among them. |
| [lint](https://github.com/surullabs/lint) | 64 | 5 | 2016-07-09 | 1 year ago | Run linters as part of go test. |
| [validate](https://github.com/mccoyst/validate) | 62 | 6 | 2013-11-22 | 3 years ago | Automatically validates struct fields with tags. |
| [go-outdated](https://github.com/firstrow/go-outdated) | 45 | 1 | 2015-06-29 | 1 year ago | Console application that displays outdated packages. |
| [golines](https://github.com/segmentio/golines) | 32 | 11 | 2019-10-01 | 1 week ago | Formatter that automatically shortens long lines in Go code. |
| [blanket](https://github.com/verygoodsoftwarenotvirus/blanket) | 14 | 2 | 2017-09-04 | 1 year ago | tarp finds functions and methods without direct unit tests in Go source code. |

## Editor Plugins
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [vim-go](https://github.com/fatih/vim-go) | 11540 | 295 | 2014-03-24 | 1 day ago | Go development plugin for Vim. |
| [vscode-go](https://github.com/microsoft/vscode-go) | 5593 | 226 | 2015-10-14 | 9 hours ago | Extension for Visual Studio Code (VS Code) which provides support for the Go language. |
| [gocode](https://github.com/nsf/gocode) | 4825 | 197 | 2010-07-05 | 2 weeks ago | Autocompletion daemon for the Go programming language. |
| [GoSublime](https://github.com/DisposaBoy/GoSublime) | 3308 | 122 | 2011-08-27 | 8 hours ago | Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features. |
| [go-plus](https://github.com/joefitzgerald/go-plus) | 1498 | 44 | 2014-03-13 | 4 days ago | Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting. |
| [go-mode.el](https://github.com/dominikh/go-mode.el) | 1023 | 55 | 2013-01-30 | 1 week ago | Go mode for GNU/Emacs. |
| [Watch](https://github.com/eaburns/Watch) | 171 | 12 | 2013-08-08 | 1 year ago | Runs a command in an acme win on file changes. |
| [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) | 81 | 5 | 2012-11-25 | 3 years ago | Vim plugin to highlight syntax errors on save. |
| [go-language-server](https://github.com/theia-ide/go-language-server) | 33 | 4 | 2017-11-21 | 10 months ago | A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol. |
| [gounit-vim](https://github.com/hexdigest/gounit-vim) | 18 | 2 | 2018-02-21 | 1 year ago | Vim plugin for generating Go tests based on the function's or method's signature. |
| [theia-go-extension](https://github.com/theia-ide/theia-go-extension) | 13 | 4 | 2017-11-30 | 10 months ago | Go language support for the Theia IDE. |

## Go Generate Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gotests](https://github.com/cweill/gotests) | 2492 | 72 | 2016-01-19 | 2 months ago | Generate Go tests from your source code. |
| [genny](https://github.com/cheekybits/genny) | 1112 | 23 | 2014-10-27 | 4 months ago | Elegant generics for Go. |
| [re2dfa](https://github.com/opennota/re2dfa) | 176 | 9 | 2015-06-20 | 1 year ago | Transform regular expressions into finite state machines and output Go source code. |
| [gocontracts](https://github.com/Parquery/gocontracts) | 55 | 6 | 2018-08-13 | 1 year ago | brings design-by-contract to Go by synchronizing the code with the documentation. |
| [gounit](https://github.com/hexdigest/gounit) | 35 | 4 | 2018-02-05 | 1 year ago | Generate Go tests using your own templates. |
| [generic](https://github.com/usk81/generic) | 32 | 3 | 2016-06-15 | 9 months ago | flexible data type for Go. |
| [hasgo](https://github.com/DylanMeeus/hasgo) | 31 | 4 | 2019-05-16 | 3 weeks ago | Generate Haskell inspired functions for your slices. |

## Go Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-swagger](https://github.com/go-swagger/go-swagger) | 4723 | 118 | 2014-11-16 | 5 days ago | Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API. |
| [OctoLinker](https://github.com/OctoLinker/OctoLinker) | 4081 | 89 | 2013-12-27 | 17 hours ago | Navigate through go files efficiently with the OctoLinker browser extension for GitHub. |
| [go-callvis](https://github.com/TrueFurby/go-callvis) | 2327 | 70 | 2016-09-03 | 1 week ago | Visualize call graph of your Go program using dot format. |
| [depth](https://github.com/KyleBanks/depth) | 440 | 9 | 2017-03-04 | 6 months ago | Visualize dependency trees of any package by analyzing imports. |
| [richgo](https://github.com/kyoh86/richgo) | 427 | 4 | 2017-01-04 | 1 month ago | Enrich `go test` outputs with text decorations. |
| [rts](https://github.com/galeone/rts) | 194 | 3 | 2016-04-04 | 3 years ago | RTS: response to struct. Generates Go structs from server responses. |
| [godbg](https://github.com/tylerwince/godbg) | 164 | 5 | 2019-01-23 | 9 months ago | Implementation of Rusts `dbg!` macro for quick and easy debugging during development. |
| [colorgo](https://github.com/songgao/colorgo) | 101 | 6 | 2013-02-14 | 3 years ago | Wrapper around `go` command for colorized `go build` output. |
| [gothanks](https://github.com/psampaz/gothanks) | 85 | 3 | 2019-11-10 | 1 day ago | GoThanks automatically stars your go.mod github dependencies, sending this way some love to their maintainers. |
| [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) | 37 | 2 | 2015-05-22 | 2 years ago | Bash completion for go and wgo. |
| [generator-go-lang](https://github.com/axelspringer/generator-go-lang) | 17 | 12 | 2017-09-13 | 2 months ago | A [Yeoman](http://yeoman.io) generator to get new Go projects started. |
| [go-james](https://github.com/pieterclaerhout/go-james) | 8 | 2 | 2019-10-14 | 1 month ago | Go project skeleton creator, builds and tests your projects without the manual setup. |

## Software Packages
        
*Software written in Go.*

### DevOps Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [kubernetes](https://github.com/kubernetes/kubernetes) | 62564 | 3133 | 2014-06-06 | 3 minutes ago | Container Cluster Manager from Google. |
| [moby](https://github.com/moby/moby) | 56184 | 3207 | 2013-01-18 | 17 minutes ago | Collaborative project for the container ecosystem to assemble container-based systems. |
| [traefik](https://github.com/containous/traefik) | 27046 | 686 | 2015-09-13 | 21 minutes ago | Reverse proxy and load balancer with support for multiple backends. |
| [gitea](https://github.com/go-gitea/gitea) | 17945 | 451 | 2016-11-01 | 27 minutes ago | Fork of Gogs, entirely community driven. |
| [vegeta](https://github.com/tsenart/vegeta) | 13689 | 296 | 2013-08-13 | 2 days ago | HTTP load testing tool and library. It's over 9000! |
| [packer](https://github.com/hashicorp/packer) | 9726 | 409 | 2013-03-23 | 5 minutes ago | Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. |
| [hey](https://github.com/rakyll/hey) | 7434 | 149 | 2016-09-02 | 3 days ago | Hey is a tiny program that sends some load to a web application. |
| [gvm](https://github.com/moovweb/gvm) | 4918 | 154 | 2011-12-03 | 3 weeks ago | GVM provides an interface to manage Go versions. |
| [webhook](https://github.com/adnanh/webhook) | 4823 | 138 | 2015-01-12 | 2 weeks ago | Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server. |
| [gaia](https://github.com/gaia-pipeline/gaia) | 3967 | 103 | 2017-12-28 | 1 month ago | Build powerful pipelines in any programming language. |
| [gox](https://github.com/mitchellh/gox) | 3587 | 71 | 2013-11-17 | 4 months ago | Dead simple, no frills Go cross compile tool. |
| [bosun](https://github.com/bosun-monitor/bosun) | 2929 | 158 | 2013-11-15 | 1 day ago | Time Series Alerting Framework. |
| [bombardier](https://github.com/codesenberg/bombardier) | 1942 | 50 | 2016-05-29 | 10 months ago | Fast cross-platform HTTP benchmarking tool. |
| [goxc](https://github.com/laher/goxc) | 1641 | 48 | 2013-02-11 | 3 months ago | build tool for Go, with a focus on cross-compiling and packaging. |
| [fac](https://github.com/mkchoi212/fac) | 1640 | 28 | 2017-12-29 | 3 months ago | Command-line user interface to fix git merge conflicts. |
| [kala](https://github.com/ajvb/kala) | 1397 | 63 | 2015-03-19 | 3 days ago | Simplistic, modern, and performant job scheduler. |
| [statusok](https://github.com/sanathp/statusok) | 1264 | 47 | 2015-08-26 | 2 weeks ago | Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected. |
| [script](https://github.com/bitfield/script) | 1260 | 23 | 2019-04-20 | 1 month ago | Making it easy to write shell-like scripts in Go for DevOps and system administration tasks. |
| [s3gof3r](https://github.com/rlmcpherson/s3gof3r) | 1031 | 34 | 2013-08-02 | 1 month ago | Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3. |
| [pomerium](https://github.com/pomerium/pomerium) | 962 | 13 | 2019-01-01 | 4 minutes ago | Pomerium is an identity-aware access proxy. |
| [go-selfupdate](https://github.com/sanbornm/go-selfupdate) | 706 | 27 | 2013-11-13 | 2 weeks ago | Enable your Go applications to self update. |
| [skm](https://github.com/TimothyYe/skm) | 579 | 20 | 2017-10-11 | 4 months ago | SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily! |
| [scaleway-cli](https://github.com/scaleway/scaleway-cli) | 561 | 39 | 2015-03-20 | 54 minutes ago | Manage BareMetal Servers from Command Line (as easily as with Docker). |
| [aurora](https://github.com/xuri/aurora) | 437 | 27 | 2016-10-09 | 3 months ago | Cross-platform web-based Beanstalkd queue server console. |
| [govvv](https://github.com/ahmetb/govvv) | 423 | 10 | 2016-08-02 | 1 year ago | “go build” wrapper to easily add version information into Go binaries. |
| [gonative](https://github.com/inconshreveable/gonative) | 314 | 7 | 2014-05-01 | 3 years ago | Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages. |
| [mora](https://github.com/emicklei/mora) | 269 | 22 | 2013-07-12 | 3 years ago | REST server for accessing MongoDB documents and meta data. |
| [lstags](https://github.com/ivanilves/lstags) | 241 | 14 | 2017-08-15 | 5 minutes ago | Tool and API to sync Docker images across different registries. |
| [dogo](https://github.com/liudng/dogo) | 222 | 19 | 2014-11-19 | 10 months ago | Monitoring changes in the source file and automatically compile and run (restart). |
| [godbg](https://github.com/sirnewton01/godbg) | 218 | 18 | 2013-08-09 | 1 year ago | Web-based gdb front-end application. |
| [pewpew](https://github.com/bengadbois/pewpew) | 216 | 7 | 2016-10-12 | 5 months ago | Flexible HTTP command line stress tester. |
| [manssh](https://github.com/xwjdsh/manssh) | 210 | 3 | 2017-10-08 | 1 year ago | manssh is a command line tool for managing your ssh alias config easily. |
| [gobrew](https://github.com/cryptojuice/gobrew) | 176 | 5 | 2013-11-13 | 2 years ago | gobrew lets you easily switch between multiple versions of go. |
| [blast](https://github.com/dave/blast) | 174 | 3 | 2017-10-21 | 1 year ago | A simple tool for API load testing and batch jobs. |
| [ostent](https://github.com/ostrost/ostent) | 165 | 5 | 2014-03-31 | 1 year ago | collects and displays system metrics and optionally relays to Graphite and/or InfluxDB. |
| [grapes](https://github.com/yaronsumel/grapes) | 145 | 6 | 2016-09-01 | 4 months ago | Lightweight tool designed to distribute commands over ssh with ease. |
| [utask](https://github.com/ovh/utask) | 128 | 23 | 2019-11-05 | 2 hours ago | Automation engine that models and executes business processes declared in yaml. |
| [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) | 123 | 4 | 2017-03-03 | 2 weeks ago | Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`. |
| [jenkins-cli](https://github.com/jenkins-zh/jenkins-cli) | 117 | 7 | 2019-06-21 | 14 hours ago | Jenkins CLI allows you manage your Jenkins as an easy way. |
| [kcli](https://github.com/cswank/kcli) | 103 | 6 | 2017-03-25 | 2 weeks ago | Command line tool for inspecting kafka topics/partitions/messages. |
| [terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi) | 90 | 9 | 2017-10-17 | 3 weeks ago | Terraform provider plugin that dynamically configures itself at runtime based on an OpenAPI document (formerly known as swagger file) containing the definitions of the APIs exposed. |
| [winrm-cli](https://github.com/masterzen/winrm-cli) | 82 | 5 | 2016-05-23 | 8 months ago | Cli tool to remotely execute commands on Windows machines. |
| [go-furnace](https://github.com/go-furnace/go-furnace) | 74 | 1 | 2016-10-09 | 4 months ago | Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean. |
| [drone-scp](https://github.com/appleboy/drone-scp) | 67 | 2 | 2016-10-16 | 1 day ago | Copy files and artifacts via SSH using a binary, docker or Drone CI. |
| [dockerfile-generator](https://github.com/ozankasikci/dockerfile-generator) | 62 | 3 | 2019-08-14 | 1 week ago | A go library and an executable that produces valid Dockerfiles using various input channels. |
| [dropship](https://github.com/ChrisMcKenzie/dropship) | 48 | 3 | 2015-09-03 | 1 year ago | Tool for deploying code via cdn. |
| [rodent](https://github.com/alouche/rodent) | 30 | 2 | 2014-06-01 | 2 years ago | Rodent helps you manage Go versions, projects and track dependencies. |
| [drone-jenkins](https://github.com/appleboy/drone-jenkins) | 26 | 1 | 2016-10-15 | 3 months ago | Trigger downstream Jenkins jobs using a binary, docker or Drone CI. |
| [awsenv](https://github.com/soniah/awsenv) | 22 | 1 | 2015-08-05 | 1 year ago | Small binary that loads Amazon (AWS) environment variables for a profile. |
| [lwc](https://github.com/timdp/lwc) | 10 | 2 | 2018-04-22 | 1 year ago | A live-updating version of the UNIX wc command. |
| [depcharge](https://github.com/centerorbit/depcharge) | 9 | 3 | 2018-07-25 | 2 months ago | Helps orchestrating the execution of commands across the many dependencies in larger projects. |
| [sg](https://github.com/ChristopherRabotin/sg) | 5 | 1 | 2015-08-19 | 3 years ago | Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response. |
| [aptly-fork](https://github.com/smira/aptly-fork) | 1 | 0 | 2019-07-04 | 3 months ago | aptly is a Debian repository management tool. |

### Other Software
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [goreplay](https://github.com/buger/goreplay) | 12159 | 462 | 2013-05-30 | 3 days ago | Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time. |
| [restic](https://github.com/restic/restic) | 9192 | 238 | 2014-04-27 | 1 day ago | De-duplicating backup program. |
| [seaweedfs](https://github.com/chrislusf/seaweedfs) | 9105 | 506 | 2014-07-14 | 12 hours ago | Fast, Simple and Scalable Distributed File System with O(1) disk seek. |
| [confd](https://github.com/kelseyhightower/confd) | 6775 | 257 | 2013-10-01 | 1 week ago | Manage local application configuration files using templates and data from etcd or consul. |
| [comcast](https://github.com/tylertreat/comcast) | 6366 | 151 | 2014-11-12 | 6 months ago | Simulate bad network connections. |
| [liteide](https://github.com/visualfc/liteide) | 5815 | 376 | 2012-11-19 | 6 days ago | LiteIDE is a simple, open source, cross-platform Go IDE. |
| [drive](https://github.com/odeke-em/drive) | 5229 | 205 | 2014-11-03 | 1 week ago | Google Drive client for the commandline. |
| [nes](https://github.com/fogleman/nes) | 4337 | 149 | 2015-03-02 | 7 months ago | Nintendo Entertainment System (NES) emulator written in Go. |
| [toxiproxy](https://github.com/Shopify/toxiproxy) | 4288 | 290 | 2014-09-04 | 2 months ago | Proxy to simulate network and system conditions for automated tests. |
| [duplicacy](https://github.com/gilbertchen/duplicacy) | 3136 | 97 | 2016-02-23 | 1 week ago | A cross-platform network and cloud backup tool based on the idea of lock-free deduplication. |
| [croc](https://github.com/schollz/croc) | 2792 | 61 | 2017-10-17 | 5 days ago | Easily and securely send files or folders from one computer to another. |
| [mylg](https://github.com/mehrdadrad/mylg) | 2263 | 107 | 2016-06-21 | 1 year ago | Command Line Network Diagnostic tool written in Go. |
| [goboy](https://github.com/Humpheh/goboy) | 2183 | 40 | 2017-08-20 | 1 week ago | Nintendo Game Boy Color emulator written in Go. |
| [sup](https://github.com/pressly/sup) | 2096 | 74 | 2015-02-23 | 1 month ago | Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers. |
| [lgo](https://github.com/yunabe/lgo) | 1944 | 46 | 2017-10-05 | 6 months ago | Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility. |
| [circuit](https://github.com/gocircuit/circuit) | 1816 | 139 | 2014-04-10 | 8 months ago | Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications. |
| [snap](https://github.com/intelsdi-x/snap) | 1804 | 145 | 2014-08-13 | 1 year ago | Powerful telemetry framework. |
| [borg](https://github.com/ok-borg/borg) | 1453 | 42 | 2016-09-10 | 1 year ago | Terminal based search engine for bash snippets. |
| [scc](https://github.com/boyter/scc) | 1258 | 18 | 2018-03-01 | 19 hours ago | Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates. |
| [community](https://github.com/documize/community) | 990 | 40 | 2016-04-29 | 1 month ago | Modern wiki software that integrates data from SaaS tools. |
| [Go-Package-Store](https://github.com/shurcooL/Go-Package-Store) | 881 | 21 | 2014-01-24 | 5 months ago | App that displays updates for the Go packages in your GOPATH. |
| [peg](https://github.com/pointlander/peg) | 662 | 32 | 2010-04-25 | 2 weeks ago | Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator. |
| [leaps](https://github.com/Jeffail/leaps) | 661 | 26 | 2014-06-19 | 10 months ago | Pair programming service using Operational Transforms. |
| [vflow](https://github.com/VerizonDigital/vflow) | 643 | 86 | 2017-02-24 | 2 months ago | High-performance, scalable and reliable IPFIX, sFlow and Netflow collector. |
| [gfile](https://github.com/Antonito/gfile) | 526 | 10 | 2019-03-08 | 9 months ago | Securely transfer files between two computers, without any third party, over WebRTC. |
| [shell2http](https://github.com/msoap/shell2http) | 494 | 20 | 2015-03-11 | 2 months ago | Executing shell commands via http server (for prototyping or remote control). |
| [mockingjay-server](https://github.com/quii/mockingjay-server) | 438 | 10 | 2015-04-04 | 4 months ago | Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests. |
| [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) | 391 | 18 | 2015-10-08 | 1 month ago | Video streaming torrent client. |
| [gocc](https://github.com/goccmack/gocc) | 373 | 21 | 2015-06-05 | 3 months ago | Gocc is a compiler kit for Go written in Go. |
| [wellington](https://github.com/wellington/wellington) | 292 | 13 | 2014-12-08 | 1 year ago | Sass project management tool, extends the language with sprite functions (like Compass). |
| [ipe](https://github.com/dimiro1/ipe) | 291 | 17 | 2015-01-13 | 1 year ago | Open source Pusher server implementation compatible with Pusher client libraries written in GO. |
| [IDE](https://github.com/thestrukture/IDE) | 256 | 14 | 2017-09-09 | 5 months ago | Browser accessible IDE. Designed for Go with Go. |
| [cherry](https://github.com/rafael-santiago/cherry) | 209 | 12 | 2015-10-24 | 2 years ago | Tiny webchat server in Go. |
| [orange-cat](https://github.com/utatti/orange-cat) | 182 | 6 | 2014-11-01 | 1 year ago | Markdown previewer written in Go. |
| [orbit](https://github.com/gulien/orbit) | 133 | 8 | 2017-05-13 | 1 month ago | A simple tool for running commands and generating files from templates. |
| [joincap](https://github.com/assafmo/joincap) | 128 | 7 | 2018-05-31 | 8 months ago | Command-line utility for merging multiple pcap files together. |
| [boxed](https://github.com/tejo/boxed) | 72 | 2 | 2015-04-18 | 1 year ago | Dropbox based blog engine. |
| [dp](https://github.com/scryinfo/dp) | 57 | 9 | 2018-12-12 | 1 week ago | Through SDK for data exchange with blockchain, developers can get easy access to DAPP development. |
| [naclpipe](https://github.com/unix4fun/naclpipe) | 20 | 5 | 2015-05-05 | 1 year ago | Simple NaCL EC25519 based crypto pipe tool written in Go. |
| [term-quiz](https://github.com/crazcalm/term-quiz) | 17 | 1 | 2017-12-26 | 1 year ago | Quizzes for your terminal. |
| [snitch](https://github.com/lucasgomide/snitch) | 15 | 1 | 2017-04-06 | 1 year ago | Simple way to notify your team and many tools when someone has deployed any application via Tsuru. |
| [GoDocTooltip](https://github.com/diankong/GoDocTooltip) | 11 | 3 | 2016-01-21 | 4 years ago | Chrome extension for Go Doc sites, which shows function description as tooltip at function list. |

# Resources
        
*Where to discover new Go libraries.*

## Benchmarks
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) | 1333 | 60 | 2013-12-16 | 1 month ago | Go HTTP request router benchmark and comparison. |
| [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) | 1105 | 85 | 2016-04-06 | 4 days ago | Go web framework benchmark. |
| [skynet](https://github.com/atemerev/skynet) | 946 | 47 | 2016-02-14 | 8 months ago | Skynet 1M threads microbenchmark. |
| [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) | 945 | 39 | 2013-01-18 | 4 weeks ago | Benchmarks of Go serialization methods. |
| [speedtest-resize](https://github.com/fawick/speedtest-resize) | 187 | 7 | 2013-09-16 | 2 months ago | Compare various Image resize algorithms for the Go language. |
| [go-benchmarks](https://github.com/tylertreat/go-benchmarks) | 130 | 9 | 2016-02-25 | 3 years ago | Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches. |
| [gospeed](https://github.com/feyeleanor/gospeed) | 100 | 7 | 2011-05-23 | 1 year ago | Go micro-benchmarks for calculating the speed of language constructs. |
| [autobench](https://github.com/davecheney/autobench) | 90 | 8 | 2013-03-28 | 5 years ago | Framework to compare the performance between different Go versions. |
| [gocostmodel](https://github.com/mna/gocostmodel) | 55 | 5 | 2014-12-19 | 5 years ago | Benchmarks of common basic operations for the Go language. |
| [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) | 50 | 5 | 2014-09-24 | 1 year ago | Collection of benchmarks for popular Go database/SQL utilities. |
| [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) | 20 | 1 | 2017-01-24 | 2 years ago | Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results. |
| [kvbench](https://github.com/jimrobinson/kvbench) | 16 | 1 | 2014-04-15 | 3 months ago | Key/Value database benchmark. |

## Conferences
        

## E-Books
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [GoBooks](https://github.com/dariubs/GoBooks) | 7388 | 528 | 2015-05-05 | 1 day ago | A curated list of Go books. |
| [The-Golang-Standard-Library-by-Example](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example) | 6321 | 548 | 2013-04-14 | 3 months ago | Golang标准库。对于程序员而言，标准库与语言本身同样重要，它好比一个百宝箱，能为各种常见的任务提供完美的解决方案。以示例驱动的方式讲解Golang的标准库。 |
| [gosuccinctly](https://github.com/thedevsir/gosuccinctly) | 13 | 3 | 2018-09-02 | 1 year ago | in Persian. |

## Gophers
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gophers](https://github.com/ashleymcnamara/gophers) | 2021 | 91 | 2017-02-15 | 9 months ago | Gopher artworks by Ashley McNamara. |
| [gophers](https://github.com/egonelbre/gophers) | 1848 | 45 | 2015-06-03 | 1 month ago | Free gophers. |
| [free-gophers-pack](https://github.com/MariaLetta/free-gophers-pack) | 1777 | 46 | 2019-04-02 | 1 day ago | Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster. |
| [gophericons](https://github.com/shalakhin/gophericons) | 563 | 20 | 2015-08-22 | 1 year ago | 34 gopher images for Go developers community |
| [gopher-stickers](https://github.com/tenntenn/gopher-stickers) | 462 | 15 | 2014-11-09 | 1 month ago | gopher stickers |
| [gopherize.me](https://github.com/matryer/gopherize.me) | 358 | 7 | 2017-01-25 | 1 year ago | Gopherize yourself. |
| [gopher-vector](https://github.com/golang-samples/gopher-vector) | 352 | 13 | 2013-03-31 | 3 years ago | Vector data of gopher |
| [gopher-logos](https://github.com/GolangUA/gopher-logos) | 69 | 7 | 2017-07-27 | 1 year ago | adorable gopher logos. |
| [gophers](https://github.com/rogeralsing/gophers) | 50 | 2 | 2017-01-28 | 2 years ago | random gopher graphics. |
| [go-gopher](https://github.com/sillecelik/go-gopher) | 42 | 0 | 2018-03-28 | 3 months ago | Gopher amigurumi toy pattern. |
| [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) | 31 | 1 | 2014-09-03 | 1 year ago | Go gopher Vector Data [.ai, .svg]. |

## Meetups
        
*Add the group of your city/country here (send **PR**)*

## Twitter
        

## Websites
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) | 25513 | 1733 | 2014-07-08 | 1 month ago | List of other amazingly awesome lists. |
| [awesome-remote-job](https://github.com/lukasz-madon/awesome-remote-job) | 15597 | 863 | 2015-01-02 | 10 hours ago | Curated list of awesome remote jobs. A lot of them are looking for Go hackers. |
| [golang-graphics](https://github.com/mholt/golang-graphics) | 142 | 9 | 2014-03-24 | 4 years ago | Collection of Go images, graphics, and art. |
| [gocryforhelp](https://github.com/ninedraft/gocryforhelp) | 38 | 12 | 2016-05-09 | 2 years ago | Collection of Go projects that needs help. Good place to start your open-source way in Go. |

### Tutorials
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang) | 33567 | 2457 | 2012-08-02 | 1 week ago | Golang ebook intro how to build a web app with golang. |
| [go-patterns](https://github.com/tmrts/go-patterns) | 11943 | 555 | 2015-12-14 | 3 weeks ago | Curated list of Go design patterns, recipes and idioms. |
| [learn-go-with-tests](https://github.com/quii/learn-go-with-tests) | 9886 | 257 | 2018-03-02 | 2 days ago | Learn Go with test-driven development. |
| [golang-cheat-sheet](https://github.com/a8m/golang-cheat-sheet) | 4353 | 175 | 2014-02-13 | 5 months ago | Go's reference card. |
| [working-with-go](https://github.com/mkaz/working-with-go) | 1149 | 51 | 2014-05-04 | 2 weeks ago | Intro to go for experienced programmers. |
| [golang-for-nodejs-developers](https://github.com/miguelmota/golang-for-nodejs-developers) | 1110 | 30 | 2019-01-03 | 1 week ago | Examples of Golang compared to Node.js for learning. |
| [ethereum-development-with-go-book](https://github.com/miguelmota/ethereum-development-with-go-book) | 540 | 33 | 2018-05-16 | 1 month ago | A little e-book on Ethereum Development with Go. |

> 这仅仅是个人练手学习go语言的一个小项目，欢迎指点 <plholx@126.com> ^_^
> 更专业的go开源项目分析请移步 [Awesome Go](https://go.libhunt.com/)
