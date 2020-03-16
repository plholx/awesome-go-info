# Awesome Go Info

go语言开源项目列表，项目分类及GitHub上的开源项目数据完全来自于[awesome-go](https://github.com/avelino/awesome-go) 的[README.md](https://github.com/avelino/awesome-go/blob/master/README.md)文件，通过调用GitHub的API获取仓库信息，展示项目的star数、watch数等，方便查看go语言开源项目的一些相关信息。

_该文件仅包含[awesome-go](https://github.com/avelino/awesome-go) [README.md](https://github.com/avelino/awesome-go/blob/master/README.md)文件中列出的在GitHub上开源的优秀项目，不罗列其它golang相关的网站_
_该文件中的GitHub仓库信息数据会在每天凌晨1点左右更新,当前数据更新于2020-03-17 01:15:14_

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
| [oto](https://github.com/hajimehoshi/oto) | 548 | 8 | 2017-05-04 | 1 week ago | A low-level library to play sound on multiple platforms. |
| [portaudio](https://github.com/gordonklaus/portaudio) | 339 | 16 | 2015-09-16 | 1 year ago | Go bindings for the PortAudio audio I/O library. |
| [music-theory](https://github.com/go-music-theory/music-theory) | 285 | 12 | 2016-03-17 | 2 months ago | Music theory models in Go. |
| [waveform](https://github.com/mdlayher/waveform) | 276 | 13 | 2014-09-13 | 3 years ago | Go package capable of generating waveform images from audio streams. |
| [portmidi](https://github.com/rakyll/portmidi) | 226 | 8 | 2013-11-10 | 4 months ago | Go bindings for PortMidi. |
| [id3v2](https://github.com/bogem/id3v2) | 146 | 4 | 2016-05-15 | 6 days ago | Fast and stable ID3 parsing and writing library for Go. |
| [flac](https://github.com/mewkiz/flac) | 118 | 9 | 2012-11-01 | 2 months ago | Native Go FLAC encoder/decoder with support for FLAC streams. |
| [mix](https://github.com/go-mix/mix) | 112 | 3 | 2016-01-03 | 2 months ago | Sequence-based Go-native audio mixer for music apps. |
| [mp3](https://github.com/tcolgate/mp3) | 107 | 1 | 2015-02-26 | 2 years ago | Native Go MP3 decoder. |
| [go-sox](https://github.com/krig/go-sox) | 105 | 8 | 2013-10-08 | 1 year ago | libsox bindings for go. |
| [malgo](https://github.com/gen2brain/malgo) | 96 | 5 | 2017-11-09 | 5 months ago | Mini audio library. |
| [go-taglib](https://github.com/wtolson/go-taglib) | 72 | 6 | 2012-11-20 | 1 year ago | Go bindings for taglib. |
| [gaad](https://github.com/Comcast/gaad) | 70 | 11 | 2016-07-11 | 1 month ago | Native Go AAC bitstream parser. |
| [minimp3](https://github.com/tosone/minimp3) | 38 | 2 | 2018-01-26 | 7 months ago | Lightweight MP3 decoder library. |
| [go_mediainfo](https://github.com/zhulik/go_mediainfo) | 28 | 1 | 2015-12-13 | 4 years ago | libmediainfo bindings for go. |
| [EasyMIDI](https://github.com/algoGuy/EasyMIDI) | 27 | 2 | 2018-02-19 | 2 years ago | EasyMidi is a simple and reliable library for working with standard midi file (SMF). |
| [vorbis](https://github.com/mccoyst/vorbis) | 25 | 3 | 2013-07-12 | 11 months ago | "Native" Go Vorbis decoder (uses CGO, but has no dependencies). |
| [gosamplerate](https://github.com/dh1tw/gosamplerate) | 9 | 1 | 2016-11-20 | 1 month ago | libsamplerate bindings for go. |

## Authentication and OAuth
        
*Libraries for implementing authentications schemes.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [jwt-go](https://github.com/dgrijalva/jwt-go) | 7222 | 148 | 2012-04-18 | 1 week ago | Golang implementation of JSON Web Tokens (JWT). |
| [casbin](https://github.com/casbin/casbin) | 6254 | 184 | 2017-04-08 | 4 days ago | Authorization library that supports access control models like ACL, RBAC, ABAC. |
| [oauth2](https://github.com/golang/oauth2) | 2770 | 103 | 2014-04-14 | 5 days ago | Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support. |
| [goth](https://github.com/markbates/goth) | 2581 | 67 | 2014-10-14 | 3 days ago | provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box. |
| [authboss](https://github.com/volatiletech/authboss) | 2167 | 41 | 2015-01-03 | 1 month ago | Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time. |
| [osin](https://github.com/openshift/osin) | 1585 | 68 | 2013-09-10 | 2 months ago | Golang OAuth2 server library. |
| [go-jose](https://github.com/square/go-jose) | 1487 | 65 | 2014-11-14 | 1 week ago | Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs. |
| [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) | 1463 | 78 | 2015-11-01 | 4 months ago | Standalone, specification-compliant,  OAuth2 server written in Golang. |
| [gologin](https://github.com/dghubble/gologin) | 1169 | 28 | 2015-06-23 | 7 months ago | chainable handlers for login with OAuth1 and OAuth2 authentication providers. |
| [gorbac](https://github.com/mikespook/gorbac) | 995 | 60 | 2013-12-26 | 1 year ago | provides a lightweight role-based access control (RBAC) implementation in Golang. |
| [loginsrv](https://github.com/tarent/loginsrv) | 919 | 49 | 2016-11-11 | 4 days ago | JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam. |
| [scs](https://github.com/alexedwards/scs) | 683 | 18 | 2016-08-08 | 2 weeks ago | Session Manager for HTTP servers. |
| [permissions2](https://github.com/xyproto/permissions2) | 393 | 13 | 2014-11-19 | 3 weeks ago | Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt. |
| [paseto](https://github.com/o1egl/paseto) | 318 | 17 | 2018-01-23 | 4 days ago | Golang implementation of Platform-Agnostic Security Tokens (PASETO). |
| [httpauth](https://github.com/goji/httpauth) | 191 | 7 | 2014-05-26 | 3 years ago | HTTP Authentication middleware. |
| [jeff](https://github.com/abraithwaite/jeff) | 190 | 4 | 2018-08-02 | 5 months ago | Simple, flexible, secure and idiomatic web session management with pluggable backends. |
| [jwt-auth](https://github.com/adam-hanna/jwt-auth) | 172 | 11 | 2016-07-05 | 1 year ago | JWT middleware for Golang http servers with many configuration options. |
| [jwt](https://github.com/pascaldekloe/jwt) | 159 | 11 | 2018-03-21 | 23 hours ago | Lightweight JSON Web Token (JWT) library. |
| [branca](https://github.com/hako/branca) | 103 | 6 | 2018-01-09 | 3 days ago | Golang implementation of Branca Tokens. |
| [session](https://github.com/icza/session) | 99 | 6 | 2016-02-08 | 8 months ago | Go session management for web servers (including support for Google App Engine - GAE). |
| [sessionup](https://github.com/swithek/sessionup) | 82 | 6 | 2019-07-23 | 1 week ago | Simple, yet effective HTTP session management and identification package. |
| [jwt](https://github.com/robbert229/jwt) | 78 | 6 | 2016-06-05 | 1 year ago | Clean and easy to use implementation of JSON Web Tokens (JWT). |
| [sjwt](https://github.com/brianvoe/sjwt) | 61 | 0 | 2019-06-20 | 5 months ago | Simple jwt generator and parser. |
| [sessions](https://github.com/adam-hanna/sessions) | 48 | 3 | 2017-04-29 | 10 months ago | Dead simple, highly performant, highly customizable sessions service for go http servers. |
| [rbac](https://github.com/zpatrick/rbac) | 46 | 3 | 2018-08-02 | 1 year ago | Minimalistic RBAC package for Go applications. |
| [securecookie](https://github.com/chmike/securecookie) | 33 | 5 | 2017-09-03 | 2 weeks ago | Efficient secure cookie encoding/decoding. |
| [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) | 8 | 2 | 2017-10-20 | 1 year ago | Go session management using the SessionGate Redis module. |
| [signedvalue](https://github.com/sashka/signedvalue) | 8 | 0 | 2018-01-06 | 5 months ago | Signed and timestamped strings compatible with [Tornado's](https://github.com/tornadoweb/tornado) `create_signed_value`, `decode_signed_value`, and therefore `set_secure_cookie` and `get_secure_cookie`. |
| [scope](https://github.com/SonicRoshan/scope) | 4 | 1 | 2019-09-23 | 2 months ago | Easily Manage OAuth2 Scopes In Go. |
| [cookiestxt](https://github.com/mengzhuo/cookiestxt) | 2 | 1 | 2017-10-09 | 2 years ago | provides parser of cookies.txt file format. |

## Bot Building
        
*Libraries for building and working with bots.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) | 1975 | 62 | 2015-06-25 | 5 days ago | Simple and clean Telegram bot client. |
| [telebot](https://github.com/tucnak/telebot) | 1146 | 41 | 2015-06-25 | 2 weeks ago | Telegram bot framework written in Go. |
| [bot](https://github.com/go-chat-bot/bot) | 555 | 39 | 2015-09-22 | 3 months ago | IRC, Slack & Telegram bot written in Go. |
| [slacker](https://github.com/shomali11/slacker) | 360 | 13 | 2017-05-20 | 1 month ago | Easy to use framework to create Slack bots. |
| [golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot) | 290 | 27 | 2017-05-14 | 1 year ago | A golang implementation of a console-based trading bot for cryptocurrency exchanges. |
| [kelp](https://github.com/stellar/kelp) | 266 | 33 | 2018-08-08 | 2 days ago | official trading and market-making bot for the [Stellar](https://www.stellar.org/) DEX. Works out-of-the-box, written in Golang, compatible with centralized exchanges and custom trading strategies. |
| [tbot](https://github.com/yanzay/tbot) | 251 | 11 | 2015-09-11 | 3 months ago | Telegram bot server with API similar to net/http. |
| [tenyks](https://github.com/kyleterry/tenyks) | 167 | 14 | 2012-08-26 | 6 months ago | Service oriented IRC bot using Redis and JSON for messaging. |
| [go-sarah](https://github.com/oklahomer/go-sarah) | 156 | 7 | 2016-11-06 | 3 months ago | Framework to build bot for desired chat services including LINE, Slack, Gitter and more. |
| [hanu](https://github.com/sbstjn/hanu) | 119 | 5 | 2016-09-16 | 4 months ago | Framework for writing Slack bots. |
| [go-twitch-irc](https://github.com/gempir/go-twitch-irc) | 95 | 10 | 2017-03-23 | 2 days ago | Libary to write bots for twitch. |
| [go-tgbot](https://github.com/olebedev/go-tgbot) | 91 | 8 | 2016-12-11 | 1 year ago | Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware. |
| [margelet](https://github.com/zhulik/margelet) | 60 | 5 | 2015-11-21 | 3 years ago | Framework for building Telegram bots. |
| [govkbot](https://github.com/nikepan/govkbot) | 30 | 3 | 2016-07-11 | 1 week ago | Simple Go [VK](https://vk.com) bot library. |
| [slackscot](https://github.com/alexandre-normand/slackscot) | 24 | 1 | 2015-10-22 | 2 days ago | Another framework for building Slack bots. |
| [micha](https://github.com/onrik/micha) | 12 | 3 | 2016-04-14 | 2 weeks ago | Go Library for Telegram bot api. |

## Command Line
        
*Libraries for building Console Applications and Console User Interfaces.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [cobra](https://github.com/spf13/cobra) | 16116 | 313 | 2013-09-03 | 4 days ago | Commander for modern Go CLI interactions. |
| [cli](https://github.com/urfave/cli) | 13324 | 289 | 2013-07-13 | 3 days ago | Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli). |
| [termui](https://github.com/gizak/termui) | 9618 | 286 | 2015-02-03 | 3 weeks ago | Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib). |
| [gocui](https://github.com/jroimartin/gocui) | 6042 | 123 | 2014-01-04 | 1 month ago | Minimalist Go library aimed at creating Console User Interfaces. |
| [termbox-go](https://github.com/nsf/termbox-go) | 3701 | 99 | 2012-01-12 | 1 month ago | Termbox is a library for creating cross-platform text-based interfaces. |
| [go-prompt](https://github.com/c-bata/go-prompt) | 2878 | 47 | 2017-08-14 | 6 days ago | Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit). |
| [kingpin](https://github.com/alecthomas/kingpin) | 2818 | 57 | 2014-05-14 | 3 weeks ago | Command line and flag parser supporting sub commands. |
| [go-flags](https://github.com/jessevdk/go-flags) | 1649 | 28 | 2012-08-31 | 2 weeks ago | go command line option parser. |
| [uiprogress](https://github.com/gosuri/uiprogress) | 1647 | 33 | 2015-11-17 | 8 months ago | Flexible library to render progress bars in terminal applications. |
| [dnote](https://github.com/dnote/dnote) | 1609 | 28 | 2017-03-30 | 9 hours ago | A simple command line notebook with multi-device sync. |
| [readline](https://github.com/chzyer/readline) | 1471 | 39 | 2015-09-20 | 3 months ago | Pure golang implementation that provides most features in GNU-Readline under MIT license. |
| [asciigraph](https://github.com/guptarohit/asciigraph) | 1315 | 27 | 2018-06-17 | 5 months ago | Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. |
| [docopt.go](https://github.com/docopt/docopt.go) | 1225 | 33 | 2013-08-25 | 5 months ago | Command-line arguments parser that will make you smile. |
| [uilive](https://github.com/gosuri/uilive) | 1110 | 17 | 2015-11-16 | 1 month ago | Library for updating terminal output in realtime. |
| [cli](https://github.com/mitchellh/cli) | 1099 | 23 | 2013-11-03 | 3 months ago | Go library for implementing command-line interfaces. |
| [pflag](https://github.com/spf13/pflag) | 1023 | 29 | 2013-08-30 | 1 week ago | Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. |
| [termdash](https://github.com/mum4k/termdash) | 964 | 21 | 2018-03-24 | 1 week ago | Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui). |
| [gcli](https://github.com/tcnksm/gcli) | 894 | 26 | 2014-06-19 | 2 years ago | The easy way to start building Golang command line applications. |
| [progressbar](https://github.com/schollz/progressbar) | 867 | 20 | 2017-10-26 | 1 month ago | Basic thread-safe progress bar that works in every OS. |
| [mpb](https://github.com/vbauerster/mpb) | 852 | 15 | 2016-12-14 | 9 hours ago | Multi progress bar for terminal applications. |
| [go-arg](https://github.com/alexflint/go-arg) | 845 | 15 | 2015-11-01 | 2 weeks ago | Struct-based argument parsing in Go. |
| [aurora](https://github.com/logrusorgru/aurora) | 789 | 6 | 2016-11-06 | 2 months ago | ANSI terminal colors that supports fmt.Printf/Sprintf. |
| [complete](https://github.com/posener/complete) | 675 | 15 | 2017-05-05 | 6 days ago | Write bash completions in Go + Go command bash completion. |
| [liner](https://github.com/peterh/liner) | 667 | 24 | 2012-08-15 | 1 month ago | Go readline-like library for command-line interfaces. |
| [mow.cli](https://github.com/jawher/mow.cli) | 658 | 17 | 2014-12-18 | 2 months ago | Go library for building CLI applications with sophisticated flag and argument parsing and validation. |
| [flaggy](https://github.com/integrii/flaggy) | 633 | 17 | 2018-03-05 | 1 week ago | A robust and idiomatic flags package with excellent subcommand support. |
| [uitable](https://github.com/gosuri/uitable) | 547 | 15 | 2015-11-13 | 4 months ago | Library to improve readability in terminal apps using tabular data. |
| [cli](https://github.com/mkideal/cli) | 518 | 22 | 2016-02-26 | 6 months ago | Feature-rich and easy to use command-line package based on golang struct tags. |
| [go-colorable](https://github.com/mattn/go-colorable) | 419 | 18 | 2014-07-30 | 2 weeks ago | Colorable writer for windows. |
| [go-isatty](https://github.com/mattn/go-isatty) | 405 | 8 | 2014-04-01 | 1 month ago | isatty for golang. |
| [ops](https://github.com/nanovms/ops) | 387 | 22 | 2018-09-10 | 1 week ago | Unikernel Builder/Orchestrator. |
| [chalk](https://github.com/ttacon/chalk) | 329 | 8 | 2014-07-18 | 6 months ago | Intuitive package for prettifying terminal/console output. |
| [color](https://github.com/gookit/color) | 329 | 8 | 2018-07-01 | 1 week ago | Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows. |
| [tabby](https://github.com/cheynewallace/tabby) | 264 | 2 | 2018-12-17 | 11 months ago | A tiny library for super simple Golang tables. |
| [simpletable](https://github.com/alexeyco/simpletable) | 216 | 4 | 2017-03-29 | 1 month ago | Simple tables in terminal with Go. |
| [go-colortext](https://github.com/daviddengcn/go-colortext) | 201 | 9 | 2013-01-23 | 1 year ago | Go library for color output in terminals. |
| [argparse](https://github.com/akamensky/argparse) | 176 | 7 | 2017-11-24 | 1 week ago | Command line argument parser inspired by Python's argparse module. |
| [commandeer](https://github.com/jaffee/commandeer) | 109 | 7 | 2017-10-12 | 4 months ago | Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags. |
| [sflags](https://github.com/octago/sflags) | 109 | 5 | 2016-12-04 | 7 months ago | Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries. |
| [flag](https://github.com/zhuah/flag) | 102 | 5 | 2016-10-05 | 1 year ago | Simple but powerful command line option parsing library for Go supporting subcommand. |
| [wmenu](https://github.com/dixonwille/wmenu) | 102 | 2 | 2016-04-20 | 2 months ago | Easy to use menu structure for cli applications that prompts users to make choices. |
| [clif](https://github.com/ukautz/clif) | 101 | 2 | 2015-05-30 | 1 year ago | Small command line interface framework. |
| [cfmt](https://github.com/mingrammer/cfmt) | 72 | 3 | 2018-03-15 | 1 year ago | Contextual fmt inspired by bootstrap color classes. |
| [job](https://github.com/liujianping/job) | 70 | 2 | 2019-04-09 | 9 months ago | JOB, make your short-term command as a long-term job. |
| [cli](https://github.com/teris-io/cli) | 65 | 1 | 2017-05-24 | 9 months ago | Simple and complete API for building command line interfaces in Go. |
| [1build](https://github.com/gopinath-langote/1build) | 54 | 4 | 2019-04-23 | 3 weeks ago | Command line tool to frictionlessly manage project-specific commands. |
| [env](https://github.com/codingconcepts/env) | 48 | 1 | 2017-06-14 | 9 months ago | Tag-based environment configuration for structs. |
| [wlog](https://github.com/dixonwille/wlog) | 42 | 1 | 2016-04-13 | 2 months ago | Simple logging interface that supports cross-platform color and concurrency. |
| [gocmd](https://github.com/devfacet/gocmd) | 37 | 3 | 2018-01-08 | 1 year ago | Go library for building command line applications. |
| [tabular](https://github.com/InVisionApp/tabular) | 36 | 4 | 2018-04-23 | 1 year ago | Print ASCII tables from command line utilities without the need to pass large sets of data to the API. |
| [strumt](https://github.com/antham/strumt) | 34 | 0 | 2017-06-19 | 3 weeks ago | Library to create prompt chain. |
| [flagvar](https://github.com/sgreben/flagvar) | 33 | 2 | 2018-05-18 | 2 months ago | A collection of flag argument types for Go's standard `flag` package. |
| [cmdr](https://github.com/hedzr/cmdr) | 30 | 4 | 2019-05-15 | 1 week ago | A POSIX/GNU style, getopt-like command-line UI Go library. |
| [clir](https://github.com/leaanthony/clir) | 21 | 2 | 2019-11-18 | 3 weeks ago | A Simple and Clear CLI library. Dependency free. |
| [argv](https://github.com/zhuah/argv) | 19 | 1 | 2017-01-22 | 9 months ago | Go library to split command line string as arguments array using the bash syntax. |
| [go-getoptions](https://github.com/DavidGamba/go-getoptions) | 19 | 3 | 2015-12-18 | 4 months ago | Go option parser inspired on the flexibility of Perl’s GetOpt::Long. |
| [colourize](https://github.com/TreyBastian/colourize) | 17 | 3 | 2015-05-11 | 3 years ago | Go library for ANSI colour text in terminals. |
| [ctc](https://github.com/wzshiming/ctc) | 17 | 1 | 2018-04-27 | 2 months ago | The non-invasive cross-platform terminal color library does not need to modify the Print method. |
| [go-commander](https://github.com/yitsushi/go-commander) | 15 | 1 | 2016-10-10 | 6 months ago | Go library to simplify CLI workflow. |
| [sand](https://github.com/Zaba505/sand) | 9 | 1 | 2018-11-18 | 1 year ago | Simple API for creating interpreters and so much more. |
| [ts](https://github.com/liujianping/ts) | 8 | 0 | 2019-06-25 | 8 months ago | Timestamp convert & compare tool. |
| [go-ataman](https://github.com/workanator/go-ataman) | 8 | 2 | 2017-05-17 | 2 years ago | Go library for rendering ANSI colored text templates in terminals. |
| [cmd](https://github.com/posener/cmd) | 5 | 1 | 2019-10-29 | 1 week ago | Extends the standard `flag` package to support sub commands and more in idomatic way. |

## Configuration
        
*Libraries for configuration parsing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [viper](https://github.com/spf13/viper) | 11481 | 212 | 2014-04-02 | 3 hours ago | Go configuration with fangs. |
| [envconfig](https://github.com/kelseyhightower/envconfig) | 2819 | 42 | 2013-11-06 | 3 weeks ago | Go library for managing configuration data from environment variables. |
| [godotenv](https://github.com/joho/godotenv) | 2676 | 35 | 2013-07-30 | 2 weeks ago | Go port of Ruby's dotenv library (Loads environment variables from `.env`). |
| [ini](https://github.com/go-ini/ini) | 1959 | 69 | 2014-12-18 | 1 day ago | Go package to read and write INI files. |
| [env](https://github.com/caarlos0/env) | 1448 | 18 | 2015-07-28 | 2 weeks ago | Parse environment variables to Go structs (with defaults). |
| [konfig](https://github.com/lalamove/konfig) | 548 | 12 | 2019-01-18 | 4 months ago | Composable, observable and performant config handling for Go for the distributed processing era. |
| [confita](https://github.com/heetch/confita) | 296 | 28 | 2017-12-21 | 2 months ago | Load configuration in cascade from multiple backends into a struct. |
| [store](https://github.com/tucnak/store) | 248 | 5 | 2015-10-03 | 2 years ago | Lightweight configuration manager for Go. |
| [config](https://github.com/JeremyLoy/config) | 231 | 1 | 2019-04-02 | 1 week ago | Cloud native application configuration. Bind ENV to structs in only two lines. |
| [config](https://github.com/olebedev/config) | 225 | 8 | 2014-04-21 | 9 months ago | JSON or YAML configuration wrapper with environment variables and flags parsing. |
| [config](https://github.com/joshbetz/config) | 198 | 2 | 2017-04-02 | 4 months ago | Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP. |
| [hjson-go](https://github.com/hjson/hjson-go) | 195 | 9 | 2016-08-05 | 1 hour ago | Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments. |
| [envconfig](https://github.com/vrischmann/envconfig) | 172 | 4 | 2015-04-21 | 8 months ago | Read your configuration from environment variables. |
| [koanf](https://github.com/knadh/koanf) | 170 | 5 | 2019-06-18 | 3 weeks ago | Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line. |
| [config](https://github.com/gookit/config) | 139 | 5 | 2018-07-07 | 3 months ago | application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge. |
| [gcfg](https://github.com/go-gcfg/gcfg) | 127 | 7 | 2015-08-17 | 3 months ago | read INI-style configuration files into Go structs; supports user-defined types and subsections. |
| [goconfig](https://github.com/crgimenes/goconfig) | 124 | 10 | 2016-12-18 | 5 days ago | Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. |
| [envh](https://github.com/antham/envh) | 94 | 3 | 2017-01-12 | 3 weeks ago | Helpers to manage environment variables. |
| [envcfg](https://github.com/tomazk/envcfg) | 91 | 1 | 2014-11-29 | 2 years ago | Un-marshaling environment variables to Go structs. |
| [cleanenv](https://github.com/ilyakaznacheev/cleanenv) | 58 | 3 | 2019-07-12 | 3 weeks ago | Minimalistic configuration reader (from files, ENV, and wherever you want). |
| [gofigure](https://github.com/ian-kent/gofigure) | 58 | 6 | 2014-11-25 | 6 months ago | Go application configuration made easy. |
| [configure](https://github.com/paked/configure) | 54 | 3 | 2015-06-14 | 1 year ago | Provides configuration through multiple sources, including JSON, flags and environment variables. |
| [harvester](https://github.com/beatlabs/harvester) | 52 | 12 | 2019-04-09 | 23 hours ago | Harvester, a easy to use static and dynamic configuration package supportig seeding, env vars and Consul integration. |
| [config](https://github.com/golobby/config) | 51 | 2 | 2019-10-15 | 1 month ago | A lightweight yet powerful config package for Go projects. |
| [xdg](https://github.com/OpenPeeDeeP/xdg) | 48 | 3 | 2017-07-20 | 3 months ago | Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html). |
| [ingo](https://github.com/schachmat/ingo) | 28 | 1 | 2016-02-07 | 2 years ago | Flags persisted in an ini-like config file. |
| [go-up](https://github.com/ufoscout/go-up) | 26 | 1 | 2018-02-18 | 2 months ago | A simple configuration library with recursive placeholders resolution and no magic. |
| [mini](https://github.com/sasbury/mini) | 22 | 1 | 2015-04-29 | 1 year ago | Golang package for parsing ini-style configuration files. |
| [conflate](https://github.com/the4thamigo-uk/conflate) | 13 | 0 | 2018-02-01 | 11 months ago | Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema. |
| [envconf](https://github.com/ian-kent/envconf) | 9 | 1 | 2014-10-26 | 5 years ago | Configuration from environment. |
| [genv](https://github.com/sakirsensoy/genv) | 9 | 1 | 2019-07-15 | 7 months ago | Read environment variables easily with dotenv support. |
| [sprbox](https://github.com/oblq/sprbox) | 5 | 2 | 2018-07-17 | 7 months ago | Build-environment aware toolbox factory and agnostic config parser (YAML, TOML, JSON and Environment vars). |
| [go-ssm-config](https://github.com/ianlopshire/go-ssm-config) | 3 | 1 | 2019-12-02 | 2 months ago | Go utility for loading configuration parameters from AWS SSM (Parameter Store). |
| [env](https://github.com/nasermirzaei89/env) | 2 | 0 | 2019-07-24 | 5 months ago | Simple useful package for read environment variables. |

## Continuous Integration
        
*Tools for help with continuous integration.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [drone](https://github.com/drone/drone) | 20672 | 577 | 2014-02-07 | 2 days ago | Drone is a Continuous Integration platform built on Docker, written in Go. |
| [cds](https://github.com/ovh/cds) | 2739 | 71 | 2016-10-11 | 1 hour ago | Enterprise-Grade CI/CD and DevOps Automation Open Source Platform. |
| [goveralls](https://github.com/mattn/goveralls) | 634 | 15 | 2013-04-17 | 1 month ago | Go integration for Coveralls.io continuous code coverage tracking system. |
| [overalls](https://github.com/go-playground/overalls) | 101 | 3 | 2015-07-30 | 2 months ago | Multi-Package go project coverprofile for tools like goveralls. |
| [duci](https://github.com/duck8823/duci) | 53 | 3 | 2018-04-01 | 2 weeks ago | A simple ci server no needs domain specific languages. |
| [gomason](https://github.com/nikogura/gomason) | 44 | 1 | 2017-11-18 | 1 week ago | Test, Build, Sign, and Publish your go binaries from a clean workspace. |
| [roveralls](https://github.com/lawrencewoodman/roveralls) | 12 | 1 | 2016-06-26 | 2 years ago | Recursive coverage testing tool. |

## CSS Preprocessors
        
*Libraries for preprocessing CSS files.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gcss](https://github.com/yosssi/gcss) | 431 | 16 | 2014-09-04 | 5 years ago | Pure Go CSS Preprocessor. |
| [go-libsass](https://github.com/wellington/go-libsass) | 155 | 7 | 2015-04-19 | 1 year ago | Go wrapper to the 100% Sass compatible libsass project. |

## Data Structures
        
*Generic datastructures and algorithms in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gods](https://github.com/emirpasic/gods) | 7910 | 334 | 2015-03-04 | 1 month ago | Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc. |
| [go-datastructures](https://github.com/Workiva/go-datastructures) | 5480 | 319 | 2014-10-29 | 1 week ago | Collection of useful, performant, and thread-safe data structures. |
| [golang-set](https://github.com/deckarep/golang-set) | 1418 | 39 | 2013-07-03 | 3 months ago | Thread-Safe and Non-Thread-Safe high-performance sets for Go. |
| [BoomFilters](https://github.com/tylertreat/BoomFilters) | 1236 | 40 | 2015-02-06 | 1 year ago | Probabilistic data structures for processing continuous, unbounded streams. |
| [gota](https://github.com/go-gota/gota) | 1105 | 62 | 2016-02-06 | 1 month ago | Implementation of dataframes, series, and data wrangling methods for Go. |
| [roaring](https://github.com/RoaringBitmap/roaring) | 796 | 38 | 2014-07-10 | 4 months ago | Go package implementing compressed bitsets. |
| [bloom](https://github.com/willf/bloom) | 757 | 29 | 2011-05-21 | 5 months ago | Go package implementing Bloom filters. |
| [hyperloglog](https://github.com/axiomhq/hyperloglog) | 691 | 16 | 2017-06-18 | 4 months ago | HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction. |
| [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) | 590 | 17 | 2015-06-28 | 2 months ago | Cuckoo filter: a good alternative to a counting bloom filter implemented in Go. |
| [bitset](https://github.com/willf/bitset) | 539 | 29 | 2011-05-11 | 5 months ago | Go package implementing bitsets. |
| [trie](https://github.com/derekparker/trie) | 463 | 18 | 2014-03-06 | 7 months ago | Trie implementation in Go. |
| [gocache](https://github.com/eko/gocache) | 373 | 11 | 2019-10-05 | 12 hours ago | A complete Go cache library with mutiple stores (memory, memcache, redis, ...), chainable, loadable, metrics cache and more. |
| [algorithms](https://github.com/shady831213/algorithms) | 364 | 13 | 2018-01-31 | 11 months ago | Algorithms and data structures.CLRS study. |
| [go-geoindex](https://github.com/hailocab/go-geoindex) | 320 | 67 | 2015-01-22 | 2 years ago | In-memory geo index. |
| [mafsa](https://github.com/smartystreets-archives/mafsa) | 281 | 16 | 2014-11-07 | 9 months ago | MA-FSA implementation with Minimal Perfect Hashing. |
| [goskiplist](https://github.com/ryszard/goskiplist) | 217 | 15 | 2012-05-09 | 4 months ago | Skip list implementation in Go. |
| [hilbert](https://github.com/google/hilbert) | 197 | 22 | 2015-08-06 | 1 year ago | Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves. |
| [merkletree](https://github.com/cbergoon/merkletree) | 182 | 4 | 2017-04-12 | 3 months ago | Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures. |
| [binpacker](https://github.com/zhuangsirui/binpacker) | 135 | 11 | 2016-02-02 | 1 year ago | Binary packer and unpacker helps user build custom binary stream. |
| [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) | 133 | 7 | 2016-04-01 | 6 months ago | Go implementation of Adaptive Radix Tree. |
| [bloom](https://github.com/zentures/bloom) | 131 | 7 | 2013-09-03 | 1 year ago | Bloom filters implemented in Go. |
| [ttlcache](https://github.com/ReneKroon/ttlcache) | 129 | 10 | 2014-12-13 | 1 month ago | In-memory LRU string-interface{} map with expiration for golang. |
| [iter](https://github.com/disksing/iter) | 118 | 5 | 2019-10-20 | 3 months ago | Go implementation of C++ STL iterators and algorithms. |
| [skiplist](https://github.com/MauriceGit/skiplist) | 117 | 5 | 2018-06-23 | 3 months ago | Very fast Go Skiplist implementation. |
| [deque](https://github.com/gammazero/deque) | 107 | 6 | 2018-04-24 | 5 days ago | Fast ring-buffer deque (double-ended queue). |
| [ring](https://github.com/TheTannerRyan/ring) | 105 | 1 | 2019-01-27 | 5 months ago | Go implementation of a high performance, thread safe bloom filter. |
| [go-rquad](https://github.com/arl/go-rquad) | 101 | 3 | 2016-09-12 | 1 year ago | Region quadtrees with efficient point location and neighbour finding. |
| [encoding](https://github.com/zentures/encoding) | 100 | 6 | 2013-09-20 | 2 years ago | Integer Compression Libraries for Go. |
| [conjungo](https://github.com/InVisionApp/conjungo) | 87 | 18 | 2016-12-29 | 10 months ago | A small, powerful and flexible merge library. |
| [bit](https://github.com/yourbasic/bit) | 84 | 6 | 2017-05-03 | 2 years ago | Golang set data structure with bonus bit-twiddling functions. |
| [gostl](https://github.com/liyue201/gostl) | 77 | 4 | 2019-10-12 | 1 month ago | Data structure and algorithm library for go, designed to provide functions similar to C++ STL. |
| [levenshtein](https://github.com/agnivade/levenshtein) | 72 | 3 | 2014-07-30 | 4 months ago | Implementation to calculate levenshtein distance in Go. |
| [skiplist](https://github.com/gansidui/skiplist) | 69 | 7 | 2014-11-18 | 5 years ago | Skiplist implementation in Go. |
| [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) | 64 | 1 | 2019-01-10 | 3 months ago | Concurrent FIFO queue. |
| [bloom](https://github.com/yourbasic/bloom) | 49 | 2 | 2017-05-06 | 2 years ago | Golang Bloom filter implementation. |
| [go-mcache](https://github.com/OrlovEvgeny/go-mcache) | 48 | 3 | 2018-04-14 | 1 month ago | Fast in-memory key:value store/cache library. Pointer caches. |
| [count-min-log](https://github.com/seiflotfy/count-min-log) | 46 | 2 | 2015-08-16 | 3 years ago | Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory). |
| [levenshtein](https://github.com/agext/levenshtein) | 39 | 1 | 2016-04-08 | 3 days ago | Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. |
| [remember-go](https://github.com/rocketlaunchr/remember-go) | 38 | 3 | 2019-04-04 | 5 months ago | A universal interface for caching slow database queries (backed by redis, memcached, ristretto, or in-memory). |
| [concurrent-writer](https://github.com/free/concurrent-writer) | 27 | 4 | 2017-09-18 | 2 years ago | Highly concurrent drop-in replacement for `bufio.Writer`. |
| [crunch](https://github.com/superwhiskers/crunch) | 23 | 4 | 2019-02-27 | 20 hours ago | Go package implementing buffers for handling various datatypes easily. |
| [pipeline](https://github.com/hyfather/pipeline) | 22 | 1 | 2018-04-25 | 1 year ago | An implementation of pipelines with fan-in and fan-out. |
| [goset](https://github.com/zoumo/goset) | 20 | 1 | 2017-08-25 | 4 months ago | A useful Set collection implementation for Go. |
| [typ](https://github.com/gurukami/typ) | 16 | 1 | 2019-03-03 | 6 months ago | Null Types, Safe primitive type conversion and fetching value from complex structures. |
| [deque](https://github.com/edwingeng/deque) | 15 | 1 | 2019-02-01 | 2 days ago | A highly optimized double-ended queue. |
| [hide](https://github.com/emvi/hide) | 14 | 3 | 2019-01-16 | 1 year ago | ID type with marshalling to/from hash to prevent sending IDs to clients. |
| [dict](https://github.com/srfrog/dict) | 11 | 0 | 2019-04-23 | 9 months ago | Python-like dictionaries (dict) for Go. |
| [go-ef](https://github.com/amallia/go-ef) | 11 | 2 | 2017-09-22 | 2 years ago | A Go implementation of the Elias-Fano encoding. |
| [mspm](https://github.com/BlackRabbitt/mspm) | 9 | 3 | 2018-05-17 | 1 year ago | Multi-String Pattern Matching Algorithm for information retrieval. |
| [null](https://github.com/emvi/null) | 9 | 2 | 2018-07-04 | 6 months ago | Nullable Go types that can be marshalled/unmarshalled to/from JSON. |
| [ptrie](https://github.com/viant/ptrie) | 7 | 7 | 2019-05-20 | 4 months ago | An implementation of prefix tree. |
| [set](https://github.com/StudioSol/set) | 7 | 17 | 2018-07-20 | 8 months ago | Simple set data structure implementation in Go using LinkedHashMap. |
| [treap](https://github.com/perdata/treap) | 7 | 2 | 2018-09-16 | 2 months ago | Persistent, fast ordered map using tree heaps. |
| [gofal](https://github.com/xxjwxc/gofal) | 6 | 1 | 2019-08-05 | 5 months ago | fractional api for Go. |
| [timedmap](https://github.com/zekroTJA/timedmap) | 6 | 1 | 2019-01-30 | 4 months ago | Map with expiring key-value pairs. |
| [parsefields](https://github.com/MonaxGT/parsefields) | 3 | 1 | 2019-04-12 | 10 months ago | Tools for parse JSON-like logs for collecting unique fields and events. |

## Database
        
*SQL query builder, libraries for building and using SQL.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [prometheus](https://github.com/prometheus/prometheus) | 29623 | 1108 | 2012-11-24 | 1 hour ago | Monitoring system and time series database. |
| [tidb](https://github.com/pingcap/tidb) | 22812 | 1299 | 2015-09-06 | 52 minutes ago | TiDB is a distributed SQL database. Inspired by the design of Google F1. |
| [influxdb](https://github.com/influxdata/influxdb) | 18478 | 754 | 2013-09-26 | 1 hour ago | Scalable datastore for metrics, events, and real-time analytics. |
| [cockroach](https://github.com/cockroachdb/cockroach) | 17919 | 735 | 2014-02-06 | 54 minutes ago | Scalable, Geo-Replicated, Transactional Datastore. |
| [dgraph](https://github.com/dgraph-io/dgraph) | 12678 | 374 | 2015-08-25 | 1 hour ago | Scalable, Distributed, Low Latency, High Throughput Graph Database. |
| [bolt](https://github.com/boltdb/bolt) | 10426 | 343 | 2013-12-20 | 2 years ago | Low-level key/value database for Go. |
| [vitess](https://github.com/vitessio/vitess) | 9672 | 504 | 2013-06-27 | 1 hour ago | vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. |
| [groupcache](https://github.com/golang/groupcache) | 8461 | 464 | 2013-07-22 | 4 days ago | Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. |
| [badger](https://github.com/dgraph-io/badger) | 7395 | 249 | 2017-01-26 | 1 hour ago | Fast key-value store in Go. |
| [pgweb](https://github.com/sosedoff/pgweb) | 6312 | 152 | 2014-10-09 | 15 hours ago | Web-based PostgreSQL database browser. |
| [rqlite](https://github.com/rqlite/rqlite) | 5724 | 205 | 2014-08-23 | 3 weeks ago | The lightweight, distributed, relational database built on SQLite. |
| [kingshard](https://github.com/flike/kingshard) | 5044 | 403 | 2015-07-04 | 2 weeks ago | kingshard is a high performance proxy for MySQL powered by Golang. |
| [migrate](https://github.com/golang-migrate/migrate) | 3839 | 52 | 2018-01-19 | 2 hours ago | Database migrations. CLI and Golang library. |
| [bigcache](https://github.com/allegro/bigcache) | 3651 | 103 | 2016-03-23 | 2 days ago | Efficient key/value cache for gigabytes of data. |
| [goleveldb](https://github.com/syndtr/goleveldb) | 3603 | 174 | 2013-01-23 | 3 weeks ago | Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go. |
| [go-cache](https://github.com/patrickmn/go-cache) | 3586 | 99 | 2012-01-02 | 4 days ago | In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. |
| [orchestrator](https://github.com/openark/orchestrator) | 3420 | 290 | 2016-11-30 | 3 days ago | MySQL replication topology manager & visualizer. |
| [orchestrator](https://github.com/github/orchestrator) | 3362 | 299 | 2016-11-30 | 1 month ago | MySQL replication topology manager & visualizer. |
| [ledisdb](https://github.com/siddontang/ledisdb) | 3242 | 188 | 2014-04-30 | 3 months ago | Ledisdb is a high performance NoSQL like Redis based on LevelDB. |
| [bbolt](https://github.com/etcd-io/bbolt) | 3041 | 101 | 2017-06-17 | 2 days ago | An embedded key/value database for Go. |
| [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) | 2819 | 165 | 2015-01-15 | 2 months ago | Sync your MySQL data into Elasticsearch automatically. |
| [squirrel](https://github.com/Masterminds/squirrel) | 2795 | 43 | 2014-01-18 | 5 days ago | Go library that helps you build SQL queries. |
| [buntdb](https://github.com/tidwall/buntdb) | 2660 | 97 | 2016-07-19 | 1 month ago | Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support. |
| [tiedot](https://github.com/HouzuoGuo/tiedot) | 2463 | 157 | 2013-05-26 | 4 months ago | Your NoSQL database powered by Golang. |
| [xo](https://github.com/xo/xo) | 2393 | 74 | 2016-02-05 | 1 month ago | Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. |
| [go-mysql](https://github.com/siddontang/go-mysql) | 2263 | 155 | 2014-02-21 | 1 day ago | Go toolset to handle MySQL protocol and replication. |
| [prest](https://github.com/prest/prest) | 2219 | 85 | 2016-11-22 | 9 months ago | Serve a RESTful API from any PostgreSQL database. |
| [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) | 1938 | 58 | 2018-09-30 | 2 days ago | fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL. |
| [sql-migrate](https://github.com/rubenv/sql-migrate) | 1640 | 30 | 2014-09-09 | 1 month ago | Database migration tool. Allows embedding migrations into the application using go-bindata. |
| [cache2go](https://github.com/muesli/cache2go) | 1228 | 62 | 2013-11-11 | 4 months ago | In-memory key:value cache which supports automatic invalidation based on timeouts. |
| [nutsdb](https://github.com/xujiajun/nutsdb) | 1109 | 35 | 2018-12-07 | 2 weeks ago | Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set. |
| [gcache](https://github.com/bluele/gcache) | 1081 | 39 | 2015-01-24 | 2 months ago | Cache library with support for expirable Cache, LFU, LRU and ARC. |
| [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) | 1003 | 69 | 2018-04-11 | 6 months ago | CovenantSQL is a SQL database on blockchain. |
| [gendry](https://github.com/didi/gendry) | 953 | 56 | 2017-12-01 | 1 week ago | Non-invasive SQL builder and powerful data binder. |
| [diskv](https://github.com/peterbourgon/diskv) | 860 | 38 | 2012-03-21 | 10 months ago | Home-grown disk-backed key-value store. |
| [moss](https://github.com/couchbase/moss) | 755 | 77 | 2016-02-06 | 4 months ago | Moss is a simple LSM key-value storage engine written in 100% Go. |
| [goqu](https://github.com/doug-martin/goqu) | 739 | 26 | 2015-02-21 | 2 days ago | Idiomatic SQL builder and query library. |
| [fastcache](https://github.com/VictoriaMetrics/fastcache) | 723 | 20 | 2018-11-22 | 1 week ago | fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead. |
| [skeema](https://github.com/skeema/skeema) | 635 | 33 | 2016-10-31 | 1 week ago | Pure-SQL schema management system for MySQL, with support for sharding and external online schema change tools. |
| [eliasdb](https://github.com/krotik/eliasdb) | 561 | 24 | 2016-08-13 | 1 week ago | Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language. |
| [dotsql](https://github.com/gchaincl/dotsql) | 507 | 22 | 2014-11-20 | 7 months ago | Go library that helps you keep sql files in one place and use them with ease. |
| [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) | 461 | 23 | 2015-12-10 | 1 month ago | Powerful data retrieval methods as well as DB-agnostic query building capabilities. |
| [gormigrate](https://github.com/go-gormigrate/gormigrate) | 419 | 5 | 2016-08-31 | 3 months ago | Database schema migration helper for Gorm ORM. |
| [chproxy](https://github.com/Vertamedia/chproxy) | 382 | 25 | 2017-09-18 | 1 hour ago | HTTP proxy for ClickHouse database. |
| [levigo](https://github.com/jmhodges/levigo) | 378 | 22 | 2012-01-17 | 3 months ago | Levigo is a Go wrapper for LevelDB. |
| [bitcask](https://github.com/prologic/bitcask) | 303 | 10 | 2019-03-12 | 3 days ago | Bitcask is an embeddable, persistent and fast key-value (KV) database written in pure Go with predictable read/write performance, low latency and high throughput thanks to the bitcask on-disk layout (LSM+WAL). |
| [pudge](https://github.com/recoilme/pudge) | 249 | 9 | 2018-11-20 | 3 months ago | Fast and simple  key/value store written using Go's standard library. |
| [jet](https://github.com/go-jet/jet) | 230 | 8 | 2019-03-02 | 1 week ago | Framework for writing type-safe SQL queries in Go, with ability to easily convert database query result into desired arbitrary object structure. |
| [sqrl](https://github.com/elgris/sqrl) | 202 | 7 | 2014-06-25 | 2 months ago | SQL query builder, fork of Squirrel with improved performance. |
| [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) | 182 | 7 | 2017-04-29 | 1 week ago | Collects small insterts and sends big requests to ClickHouse servers. |
| [piladb](https://github.com/fern4lvarez/piladb) | 173 | 10 | 2015-09-08 | 1 year ago | Lightweight RESTful database engine based on stack data structures. |
| [vasto](https://github.com/chrislusf/vasto) | 171 | 19 | 2018-01-16 | 1 year ago | A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption. |
| [myreplication](https://github.com/2tvenom/myreplication) | 158 | 19 | 2015-02-04 | 1 year ago | MySql binary log replication listener. Supports statement and row based replication. |
| [kivik](https://github.com/go-kivik/kivik) | 155 | 6 | 2017-02-09 | 19 hours ago | Kivik provides a common Go and GopherJS client library for CouchDB, PouchDB, and similar databases. |
| [goose](https://github.com/steinbacher/goose) | 136 | 4 | 2016-03-04 | 3 years ago | Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts. |
| [darwin](https://github.com/GuiaBolso/darwin) | 99 | 2 | 2016-04-05 | 2 months ago | Database schema evolution library for Go. |
| [dbq](https://github.com/rocketlaunchr/dbq) | 98 | 6 | 2019-07-11 | 10 hours ago | Zero boilerplate database operations for Go. |
| [slowpoke](https://github.com/recoilme/slowpoke) | 94 | 7 | 2018-02-19 | 5 months ago | Key-value store with persistence. |
| [migrator](https://github.com/lopezator/migrator) | 90 | 4 | 2019-02-04 | 4 months ago | Dead simple Go database migration library. |
| [golang-scribble](https://github.com/nanobox-io/golang-scribble) | 88 | 3 | 2018-06-21 | 1 year ago | Tiny flat file JSON store. |
| [igor](https://github.com/galeone/igor) | 81 | 7 | 2016-03-10 | 3 months ago | Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax. |
| [octillery](https://github.com/knocknote/octillery) | 75 | 15 | 2018-11-26 | 7 hours ago | Go package for sharding databases ( Supports every ORM or raw SQL ). |
| [godbal](https://github.com/xujiajun/godbal) | 50 | 4 | 2018-02-28 | 1 year ago | Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily. |
| [couchcache](https://github.com/codingsince1985/couchcache) | 47 | 3 | 2015-04-05 | 10 months ago | RESTful caching micro-service backed by Couchbase server. |
| [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) | 45 | 1 | 2018-08-11 | 3 months ago | A Go package to help write migrations with go-pg/pg. |
| [bcache](https://github.com/iwanbk/bcache) | 44 | 2 | 2018-12-26 | 10 months ago | Eventually consistent distributed in-memory  cache Go library. |
| [cache](https://github.com/akyoto/cache) | 35 | 2 | 2019-05-11 | 2 weeks ago | In-memory key:value store with expiration time, 0 dependencies, <100 LoC, 100% coverage. |
| [dbbench](https://github.com/sj14/dbbench) | 35 | 3 | 2018-11-24 | 1 week ago | Database benchmarking tool with support for several databases and scripts. |
| [databunker](https://github.com/paranoidguy/databunker) | 34 | 6 | 2019-12-08 | 1 week ago | Personally identifiable information (PII) storage service built to comply with GDPR and CCPA. |
| [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) | 33 | 6 | 2017-12-18 | 2 years ago | BigCache with clustering support and individual item expiration. |
| [gondolier](https://github.com/emvi/gondolier) | 28 | 4 | 2017-10-18 | 10 months ago | Database migration library using struct decorators. |
| [prep](https://github.com/hexdigest/prep) | 25 | 2 | 2017-12-11 | 2 years ago | Use prepared SQL statements without changing your code. |
| [pravasan](https://github.com/pravasan/pravasan) | 24 | 6 | 2015-01-03 | 1 year ago | Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc. |
| [go-fixtures](https://github.com/RichardKnop/go-fixtures) | 23 | 1 | 2015-12-24 | 2 months ago | Django style fixtures for Golang's excellent built-in database/sql library. |
| [coffer](https://github.com/claygod/coffer) | 20 | 3 | 2019-05-13 | 1 month ago | Simple ACID key-value database that supports transactions. |
| [datagen](https://github.com/codingconcepts/datagen) | 15 | 2 | 2019-04-18 | 1 week ago | A fast data generator that's multi-table aware and supports multi-row DML. |
| [tempdb](https://github.com/rafaeljesus/tempdb) | 14 | 3 | 2017-03-17 | 2 years ago | Key-value store for temporary items. |
| [gorocksdb](https://github.com/kapitan-k/gorocksdb) | 12 | 1 | 2017-12-28 | 2 years ago | Gorocksdb is a wrapper for [RocksDB](https://rocksdb.org) written in Go. |
| [avro](https://github.com/khezen/avro) | 12 | 1 | 2019-04-07 | 1 week ago | Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes. |
| [rwdb](https://github.com/andizzle/rwdb) | 11 | 2 | 2017-10-04 | 2 years ago | rwdb provides read replica capability for multiple database servers setup. |
| [tracedb](https://github.com/unit-io/tracedb) | 9 | 1 | 2019-08-29 | 17 hours ago | Fast timeseries database for IoT, realtime messaging  applications. Access tracedb with pubsub over tcp or websocket using github.com/unit-io/trace application. |
| [sqlf](https://github.com/leporo/sqlf) | 7 | 2 | 2019-07-20 | 1 month ago | Fast SQL query builder. |
| [bucket](https://github.com/PumpkinSeed/bucket) | 6 | 3 | 2019-09-22 | 2 months ago | Optimized data structure framework for Couchbase specialized on one bucket usage. |
| [qry](https://github.com/HnH/qry) | 6 | 1 | 2019-08-20 | 1 month ago | Tool that generates constants from files with raw SQL queries. |
| [mpath-go](https://github.com/spacetab-io/mpath-go) | 5 | 3 | 2020-01-09 | 2 months ago | MPTT (Modified Preorder Tree Traversal) package for SQL records - materialized path realisation. |
| [schema](https://github.com/adlio/schema) | 4 | 1 | 2019-09-24 | 3 months ago | Library to embed schema migrations for database/sql-compatible databases inside your Go binaries. |
| [gostore](https://github.com/twharmon/gostore) | 4 | 1 | 2020-01-08 | 4 weeks ago | Gostore is a simple, durable, embedded key-value storage engine written in Go. |
| [ormlite](https://github.com/pupizoid/ormlite) | 0 | 2 | 2018-06-28 | 1 month ago | Lightweight package containing some ORM-like features and helpers for sqlite databases. |

## Database Drivers
        
*Libraries for connecting and operating databases.*

### Relational Databases
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [mysql](https://github.com/go-sql-driver/mysql) | 9128 | 412 | 2012-12-09 | 5 days ago | MySQL driver for Go. |
| [pq](https://github.com/lib/pq) | 5717 | 157 | 2012-03-12 | 1 day ago | Pure Go Postgres driver for database/sql. |
| [go-sqlite3](https://github.com/mattn/go-sqlite3) | 3880 | 134 | 2011-11-11 | 2 weeks ago | SQLite3 driver for go that uses database/sql. |
| [pgx](https://github.com/jackc/pgx) | 2431 | 63 | 2013-03-30 | 1 week ago | PostgreSQL driver supporting features beyond those exposed by database/sql. |
| [go-mssqldb](https://github.com/denisenkom/go-mssqldb) | 1153 | 65 | 2013-12-16 | 3 weeks ago | Microsoft MSSQL driver for Go. |
| [go-oci8](https://github.com/mattn/go-oci8) | 447 | 36 | 2012-02-29 | 1 week ago | Oracle driver for go that uses database/sql. |
| [goracle](https://github.com/go-goracle/goracle) | 280 | 29 | 2015-03-25 | 3 months ago | Oracle driver for Go, using the ODPI-C driver. |
| [firebirdsql](https://github.com/nakagami/firebirdsql) | 119 | 18 | 2013-08-27 | 6 days ago | Firebird RDBMS SQL driver for Go. |
| [go-adodb](https://github.com/mattn/go-adodb) | 102 | 12 | 2011-11-14 | 3 weeks ago | Microsoft ActiveX Object DataBase driver for go that uses database/sql. |
| [gofreetds](https://github.com/minus5/gofreetds) | 98 | 22 | 2012-12-06 | 1 year ago | Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org). |
| [calcite-avatica-go](https://github.com/apache/calcite-avatica-go) | 47 | 21 | 2017-08-08 | 6 months ago | Apache Avatica/Phoenix SQL driver for database/sql. |
| [bgc](https://github.com/viant/bgc) | 13 | 10 | 2016-06-13 | 1 month ago | Datastore Connectivity for BigQuery for go. |

### NoSQL Databases
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [cayley](https://github.com/cayleygraph/cayley) | 13263 | 606 | 2014-06-05 | 2 weeks ago | Graph database with support for multiple backends. |
| [redis](https://github.com/go-redis/redis) | 8243 | 242 | 2012-07-25 | 3 days ago | Redis client for Golang. |
| [redigo](https://github.com/gomodule/redigo) | 7075 | 299 | 2012-04-14 | 1 week ago | Redigo is a Go client for the Redis database. |
| [bleve](https://github.com/blevesearch/bleve) | 6410 | 253 | 2014-04-17 | 2 hours ago | Modern text indexing library for go. |
| [riot](https://github.com/go-ego/riot) | 5088 | 183 | 2017-06-21 | 6 days ago | Go Open Source, Distributed, Simple and efficient Search Engine. |
| [elastic](https://github.com/olivere/elastic) | 4800 | 166 | 2012-12-06 | 2 weeks ago | Elasticsearch client for Go. |
| [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) | 4131 | 144 | 2017-02-08 | 1 hour ago | Official MongoDB driver for the Go language. |
| [go-elasticsearch](https://github.com/elastic/go-elasticsearch) | 2245 | 288 | 2017-03-27 | 3 days ago | Official Elasticsearch client for Go. |
| [mgo](https://github.com/globalsign/mgo) | 1759 | 71 | 2017-04-13 | 3 months ago | (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms. |
| [rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1502 | 46 | 2013-09-12 | 1 day ago | Go language driver for RethinkDB. |
| [elastigo](https://github.com/mattbaird/elastigo) | 952 | 47 | 2012-10-12 | 1 year ago | Elasticsearch client library. |
| [elasticsql](https://github.com/cch123/elasticsql) | 495 | 31 | 2016-08-24 | 1 year ago | Convert sql to elasticsearch dsl in Go. |
| [redeo](https://github.com/bsm/redeo) | 369 | 24 | 2014-03-06 | 7 months ago | Redis-protocol compatible TCP servers/services. |
| [neoism](https://github.com/jmcvetta/neoism) | 363 | 24 | 2012-07-12 | 4 weeks ago | Neo4j client for Golang. |
| [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) | 316 | 37 | 2014-07-26 | 2 days ago | Aerospike client in Go language. |
| [gocb](https://github.com/couchbase/gocb) | 307 | 64 | 2015-01-15 | 4 days ago | Official Couchbase Go SDK. |
| [go-couchbase](https://github.com/couchbase/go-couchbase) | 298 | 24 | 2012-01-19 | 6 days ago | Couchbase client in Go. |
| [gokv](https://github.com/philippgille/gokv) | 178 | 6 | 2018-10-08 | 8 hours ago | Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more). |
| [cachego](https://github.com/faabiosr/cachego) | 122 | 6 | 2016-10-05 | 18 hours ago | Golang Cache component for multiple drivers. |
| [go-rejson](https://github.com/nitishm/go-rejson) | 117 | 5 | 2018-04-23 | 1 month ago | Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease. |
| [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) | 74 | 6 | 2011-06-04 | 1 year ago | Neo4j REST Client in golang. |
| [godis](https://github.com/piaohao/godis) | 72 | 8 | 2019-06-14 | 7 months ago | redis client implement by golang, inspired by jedis. |
| [skizze](https://github.com/seiflotfy/skizze) | 72 | 6 | 2016-01-17 | 3 years ago | probabilistic data-structures service and storage. |
| [arangolite](https://github.com/solher/arangolite) | 66 | 5 | 2015-10-04 | 10 months ago | Lightweight golang driver for ArangoDB. |
| [dynago](https://github.com/underarmour/dynago) | 66 | 125 | 2015-05-18 | 2 years ago | Dynago is a principle of least surprise client for DynamoDB. |
| [mgm](https://github.com/Kamva/mgm) | 60 | 3 | 2019-12-27 | 1 week ago | MongoDB model-based ODM for Go (based on official MongoDB driver). |
| [go-pilosa](https://github.com/pilosa/go-pilosa) | 34 | 16 | 2016-09-30 | 1 week ago | Go client library for Pilosa. |
| [goforestdb](https://github.com/couchbase/goforestdb) | 28 | 41 | 2014-05-14 | 3 years ago | Go bindings for ForestDB. |
| [neo4j](https://github.com/cihangir/neo4j) | 25 | 3 | 2013-05-18 | 5 years ago | Neo4j Rest API Bindings for Golang. |
| [goriak](https://github.com/zegl/goriak) | 24 | 2 | 2016-10-05 | 1 year ago | Go language driver for Riak KV. |
| [goes](https://github.com/OwnLocal/goes) | 24 | 33 | 2015-12-28 | 3 years ago | Library to interact with Elasticsearch. |
| [dsc](https://github.com/viant/dsc) | 18 | 15 | 2016-06-13 | 1 month ago | Datastore connectivity for SQL, NoSQL, structured files. |
| [xredis](https://github.com/shomali11/xredis) | 10 | 1 | 2017-06-14 | 9 months ago | Typesafe, customizable, clean & easy to use Redis client. |
| [godscache](https://github.com/defcronyke/godscache) | 7 | 2 | 2018-05-08 | 1 year ago | A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached. |
| [asc](https://github.com/viant/asc) | 4 | 10 | 2016-06-13 | 11 months ago | Datastore Connectivity for Aerospike for go. |

## Date and Time
        
*Libraries for working with dates and times.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [now](https://github.com/jinzhu/now) | 2436 | 52 | 2013-11-18 | 4 months ago | Now is a time toolkit for golang. |
| [dateparse](https://github.com/araddon/dateparse) | 992 | 15 | 2014-04-21 | 6 months ago | Parse date's without knowing format in advance. |
| [carbon](https://github.com/uniplaces/carbon) | 437 | 40 | 2016-08-03 | 1 year ago | Simple Time extension with a lot of util methods, ported from PHP Carbon library. |
| [durafmt](https://github.com/hako/durafmt) | 287 | 4 | 2016-05-20 | 3 months ago | Time duration formatting library for Go. |
| [timeutil](https://github.com/leekchan/timeutil) | 176 | 7 | 2015-08-02 | 1 year ago | Useful extensions (Timedelta, Strftime, ...) to the golang's time package. |
| [iso8601](https://github.com/relvacode/iso8601) | 71 | 3 | 2017-04-25 | 1 year ago | Efficiently parse ISO8601 date-times without regex. |
| [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) | 68 | 3 | 2016-01-31 | 4 months ago | The implementation of the Persian (Solar Hijri) Calendar in Go (golang). |
| [timespan](https://github.com/SaidinWoT/timespan) | 68 | 5 | 2014-10-07 | 1 year ago | For interacting with intervals of time, defined as a start time and a duration. |
| [date](https://github.com/rickb777/date) | 35 | 2 | 2015-11-23 | 6 months ago | Augments Time for working with dates, date ranges, time spans, periods, and time-of-day. |
| [feiertage](https://github.com/wlbr/feiertage) | 25 | 1 | 2015-11-04 | 1 month ago | Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving... |
| [go-sunrise](https://github.com/nathan-osman/go-sunrise) | 21 | 3 | 2017-06-15 | 2 years ago | Calculate the sunrise and sunset times for a given location. |
| [kair](https://github.com/GuilhermeCaruso/kair) | 15 | 1 | 2018-10-03 | 1 month ago | Date and Time - Golang Formatting Library. |
| [cronrange](https://github.com/1set/cronrange) | 13 | 0 | 2019-11-10 | 3 months ago | Parses Cron-style time range expressions, checks if the given time is within any ranges. |
| [nulltime](https://github.com/kirillDanshin/nulltime) | 10 | 1 | 2016-03-06 | 3 years ago | Nullable `time.Time`. |
| [tuesday](https://github.com/osteele/tuesday) | 8 | 1 | 2017-08-10 | 2 years ago | Ruby-compatible Strftime function. |
| [strftime](https://github.com/awoodbeck/strftime) | 4 | 1 | 2018-02-10 | 2 years ago | C99-compatible strftime formatter. |
| [go-week](https://github.com/stoewer/go-week) | 2 | 3 | 2018-02-23 | 4 months ago | An efficient package to work with ISO8601 week dates. |
| [go-str2duration](https://github.com/xhit/go-str2duration) | 2 | 1 | 2020-02-02 | 1 month ago | Convert string to duration. Support time.Duration returned string and more. |

## Distributed Systems
        
*Packages that help with building Distributed Systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [kit](https://github.com/go-kit/kit) | 16431 | 673 | 2015-02-03 | 6 days ago | Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc. |
| [grpc-go](https://github.com/grpc/grpc-go) | 10815 | 468 | 2014-12-08 | 51 minutes ago | The Go language implementation of gRPC. HTTP/2 based RPC. |
| [micro](https://github.com/micro/micro) | 7733 | 315 | 2015-01-16 | 2 hours ago | Pluggable microservice toolkit and distributed systems platform. |
| [nats-server](https://github.com/nats-io/nats-server) | 7383 | 368 | 2012-10-29 | 2 days ago | Lightweight, high performance messaging system for microservices, IoT, and cloud native systems. |
| [rpcx](https://github.com/smallnest/rpcx) | 4449 | 318 | 2016-05-18 | 6 days ago | Distributed pluggable RPC service framework like alibaba Dubbo. |
| [tendermint](https://github.com/tendermint/tendermint) | 3538 | 258 | 2014-05-14 | 2 hours ago | High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols. |
| [torrent](https://github.com/anacrolix/torrent) | 3374 | 128 | 2015-01-08 | 11 hours ago | BitTorrent client package. |
| [raft](https://github.com/hashicorp/raft) | 3298 | 276 | 2013-11-05 | 1 week ago | Golang implementation of the Raft consensus protocol, by HashiCorp. |
| [dragonboat](https://github.com/lni/dragonboat) | 2915 | 134 | 2018-12-23 | 1 day ago | A feature complete and high performance multi-group Raft library in Go. |
| [glow](https://github.com/chrislusf/glow) | 2757 | 141 | 2015-06-14 | 1 year ago | Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go. |
| [krakend](https://github.com/devopsfaith/krakend) | 2450 | 90 | 2016-11-04 | 1 week ago | Ultra performant API Gateway framework with middlewares. |
| [gleam](https://github.com/chrislusf/gleam) | 2407 | 147 | 2016-08-26 | 1 month ago | Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed. |
| [emitter](https://github.com/emitter-io/emitter) | 2276 | 92 | 2016-10-29 | 1 week ago | High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love. |
| [liftbridge](https://github.com/liftbridge-io/liftbridge) | 1480 | 63 | 2017-10-13 | 1 day ago | Lightweight, fault-tolerant message streams for NATS. |
| [hprose-golang](https://github.com/hprose/hprose-golang) | 1081 | 96 | 2014-02-14 | 1 day ago | Very newbility RPC Library, support 25+ languages now. |
| [ringpop-go](https://github.com/uber/ringpop-go) | 611 | 2368 | 2015-06-05 | 1 year ago | Scalable, fault-tolerant application-layer sharding for Go applications. |
| [gorpc](https://github.com/valyala/gorpc) | 579 | 28 | 2014-11-20 | 6 months ago | Simple, fast and scalable RPC library for high load. |
| [go-health](https://github.com/InVisionApp/go-health) | 534 | 26 | 2017-11-29 | 2 months ago | Library for enabling asynchronous dependency health checks in your service. |
| [rain](https://github.com/cenkalti/rain) | 482 | 12 | 2014-05-21 | 1 month ago | BitTorrent client and library. |
| [digota](https://github.com/digota/digota) | 325 | 24 | 2017-08-14 | 1 year ago | grpc ecommerce microservice. |
| [sleuth](https://github.com/ursiform/sleuth) | 312 | 8 | 2016-04-23 | 2 years ago | Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)). |
| [go-sundheit](https://github.com/AppsFlyer/go-sundheit) | 291 | 9 | 2019-04-08 | 1 month ago | A library built to provide support for defining async service health checks for golang services. |
| [go-jump](https://github.com/dgryski/go-jump) | 281 | 12 | 2014-06-15 | 2 years ago | Port of Google's "Jump" Consistent Hash function. |
| [consistent](https://github.com/buraksezer/consistent) | 266 | 10 | 2018-03-25 | 5 months ago | Consistent hashing with bounded loads. |
| [dht](https://github.com/anacrolix/dht) | 150 | 11 | 2016-12-14 | 15 hours ago | BitTorrent Kademlia DHT implementation. |
| [jsonrpc](https://github.com/osamingo/jsonrpc) | 125 | 6 | 2016-10-28 | 3 weeks ago | The jsonrpc package helps implement of JSON-RPC 2.0. |
| [jsonrpc](https://github.com/ybbus/jsonrpc) | 117 | 9 | 2016-11-10 | 1 month ago | JSON-RPC 2.0 HTTP client implementation. |
| [redislock](https://github.com/bsm/redislock) | 99 | 2 | 2019-06-24 | 1 week ago | Simplified distributed locking implementation using Redis. |
| [celeriac.v1](https://github.com/svcavallar/celeriac.v1) | 60 | 3 | 2015-10-10 | 1 year ago | Library for adding support for interacting and monitoring Celery workers, tasks and events in Go. |
| [doublejump](https://github.com/edwingeng/doublejump) | 50 | 3 | 2018-06-26 | 3 weeks ago | A revamped Google's jump consistent hash. |
| [drmaa](https://github.com/dgruber/drmaa) | 29 | 2 | 2013-03-17 | 1 year ago | Job submission library for cluster schedulers based on the DRMAA standard. |
| [outboxer](https://github.com/italolelis/outboxer) | 27 | 0 | 2019-02-01 | 12 hours ago | Outboxer is a go library that implements the outbox pattern. |
| [flowgraph](https://github.com/vectaport/flowgraph) | 25 | 1 | 2018-08-29 | 6 months ago | flow-based programming package. |
| [go-pdu](https://github.com/pdupub/go-pdu) | 12 | 2 | 2018-10-08 | 5 days ago | A decentralized identity-based social network. |
| [dynatomic](https://github.com/tylfin/dynatomic) | 10 | 0 | 2019-02-08 | 1 year ago | A library for using DynamoDB as an atomic counter. |

## Dynamic DNS
        
*Tools for updating dynamic DNS records.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [godns](https://github.com/TimothyYe/godns) | 554 | 23 | 2014-05-11 | 1 week ago | A dynamic DNS client tool, supports DNSPod & HE.net, written in Go. |
| [ddns](https://github.com/skibish/ddns) | 135 | 6 | 2017-03-13 | 5 months ago | Personal DDNS client with Digital Ocean Networking DNS as backend. |

## Email
        
*Libraries and tools that implement email creation and sending.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [MailHog](https://github.com/mailhog/MailHog) | 6080 | 125 | 2014-04-16 | 1 week ago | Email and SMTP testing with web and API interface. |
| [hermes](https://github.com/matcornic/hermes) | 1906 | 29 | 2017-03-25 | 2 weeks ago | Golang package that generates clean, responsive HTML e-mails. |
| [email](https://github.com/jordan-wright/email) | 1262 | 42 | 2013-12-12 | 5 days ago | A robust and flexible email library for Go. |
| [go-imap](https://github.com/emersion/go-imap) | 921 | 35 | 2016-04-26 | 1 week ago | IMAP library for clients and servers. |
| [sendgrid-go](https://github.com/sendgrid/sendgrid-go) | 569 | 190 | 2013-09-12 | 1 week ago | SendGrid's Go library for sending email. |
| [mailgun-go](https://github.com/mailgun/mailgun-go) | 443 | 55 | 2014-02-28 | 6 days ago | Go library for sending mail with the Mailgun API. |
| [hectane](https://github.com/hectane/hectane) | 181 | 13 | 2015-08-28 | 6 months ago | Lightweight SMTP client providing an HTTP API. |
| [douceur](https://github.com/aymerick/douceur) | 174 | 2 | 2015-04-09 | 2 years ago | CSS inliner for your HTML emails. |
| [go-message](https://github.com/emersion/go-message) | 143 | 10 | 2016-12-31 | 2 weeks ago | Streaming library for the Internet Message Format and mail messages. |
| [smtp](https://github.com/mailhog/smtp) | 54 | 9 | 2014-12-24 | 1 year ago | SMTP server protocol state machine. |
| [go-dkim](https://github.com/toorop/go-dkim) | 52 | 2 | 2015-04-29 | 4 months ago | DKIM library, to sign & verify email. |
| [go-premailer](https://github.com/vanng822/go-premailer) | 45 | 2 | 2015-02-16 | 3 months ago | Inline styling for HTML mail in Go. |
| [mailchain](https://github.com/mailchain/mailchain) | 42 | 6 | 2019-04-11 | 1 day ago | Send encrypted emails to blockchain addresses written in Go. |
| [go-simple-mail](https://github.com/xhit/go-simple-mail) | 19 | 1 | 2019-09-15 | 3 months ago | Very simple package to send emails with SMTP Keep Alive and two timeouts: Connect and Send. |

## Embeddable Scripting Languages
        
*Embedding other languages inside your go code.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [otto](https://github.com/robertkrimen/otto) | 5165 | 191 | 2012-10-06 | 1 month ago | JavaScript interpreter written in Go. |
| [gopher-lua](https://github.com/yuin/gopher-lua) | 3334 | 144 | 2015-02-15 | 6 days ago | Lua 5.1 VM and compiler written in Go. |
| [go-lua](https://github.com/Shopify/go-lua) | 1792 | 267 | 2013-12-20 | 1 month ago | Port of the Lua 5.2 VM to pure Go. |
| [tengo](https://github.com/d5/tengo) | 1571 | 36 | 2019-01-09 | 22 hours ago | Bytecode compiled script language for Go. |
| [expr](https://github.com/antonmedv/expr) | 1149 | 31 | 2018-07-14 | 3 hours ago | an engine that can evaluate expressions. |
| [go-python](https://github.com/sbinet/go-python) | 1011 | 45 | 2012-07-09 | 8 months ago | naive go bindings to the CPython C-API. |
| [anko](https://github.com/mattn/anko) | 982 | 46 | 2014-03-28 | 1 week ago | Scriptable interpreter written in Go. |
| [go-php](https://github.com/deuill/go-php) | 729 | 44 | 2015-09-17 | 1 year ago | PHP bindings for Go. |
| [go-duktape](https://github.com/olebedev/go-duktape) | 686 | 27 | 2015-01-08 | 1 week ago | Duktape JavaScript engine bindings for Go. |
| [golua](https://github.com/aarzilli/golua) | 469 | 33 | 2010-12-06 | 3 weeks ago | Go bindings for Lua C API. |
| [gisp](https://github.com/jcla1/gisp) | 435 | 19 | 2014-01-11 | 2 years ago | Simple LISP in Go. |
| [cel-go](https://github.com/google/cel-go) | 380 | 23 | 2018-03-09 | 5 days ago | Fast, portable, non-Turing complete expression evaluation with gradual typing. |
| [gval](https://github.com/PaesslerAG/gval) | 184 | 15 | 2017-09-27 | 7 months ago | A highly customizable expression language written in Go. |
| [gentee](https://github.com/gentee/gentee) | 45 | 3 | 2018-01-14 | 3 weeks ago | Embeddable scripting programming language. |
| [binder](https://github.com/alexeyco/binder) | 35 | 2 | 2017-04-02 | 1 year ago | Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua). |
| [purl](https://github.com/ian-kent/purl) | 29 | 2 | 2014-11-29 | 5 years ago | Perl 5.18.2 embedded in Go. |
| [ngaro](https://github.com/db47h/ngaro) | 19 | 1 | 2016-08-09 | 1 year ago | Embeddable Ngaro VM implementation enabling scripting in Retro. |

## Error Handling
        
*Libraries for handling errors.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [errors](https://github.com/pkg/errors) | 5594 | 109 | 2015-12-27 | 2 months ago | Package that provides simple error handling primitives. |
| [go-multierror](https://github.com/hashicorp/go-multierror) | 816 | 187 | 2014-12-15 | 2 months ago | Go (golang) package for representing a list of errors as a single error. |
| [errorx](https://github.com/joomcode/errorx) | 611 | 53 | 2018-08-17 | 1 week ago | A feature rich error package with stack traces, composition of errors and more. |
| [tracerr](https://github.com/ztrue/tracerr) | 595 | 8 | 2019-02-06 | 1 year ago | Golang errors with stack trace and source fragments. |
| [eris](https://github.com/rotisserie/eris) | 579 | 8 | 2019-09-07 | 1 week ago | A better way to handle, trace, and log errors in Go. Compatible with the standard error library and github.com/pkg/errors. |
| [errlog](https://github.com/snwfdhmp/errlog) | 303 | 6 | 2019-02-16 | 1 month ago | Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place. |
| [emperror](https://github.com/emperror/emperror) | 115 | 2 | 2017-06-13 | 1 month ago | Error handling tools and best practices for Go libraries and applications. |
| [errors](https://github.com/emperror/errors) | 43 | 3 | 2019-07-09 | 1 month ago | Drop-in replacement for the standard library errors package and github.com/pkg/errors. Provides various error handling primitives. |
| [werr](https://github.com/txgruppi/werr) | 13 | 0 | 2015-08-04 | 4 years ago | Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called. |
| [errors](https://github.com/neuronlabs/errors) | 3 | 1 | 2019-07-26 | 7 months ago | Simple golang error handling with classification primitives. |
| [falcon](https://github.com/SonicRoshan/falcon) | 2 | 1 | 2019-09-09 | 5 months ago | A Simple Yet Highly Powerful Package For Error Handling. |
| [errors](https://github.com/PumpkinSeed/errors) | 2 | 1 | 2020-01-08 | 2 months ago | The most simple error wrapper with awesome performance and minimal memory overhead. |

## Files
        
*Libraries for handling files and file systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [afero](https://github.com/spf13/afero) | 2613 | 94 | 2014-10-28 | 1 month ago | FileSystem Abstraction System for Go. |
| [pdfcpu](https://github.com/pdfcpu/pdfcpu) | 1349 | 37 | 2017-06-18 | 4 weeks ago | PDF processor. |
| [notify](https://github.com/rjeczalik/notify) | 547 | 28 | 2014-09-08 | 2 months ago | File system event notification library with simple API, similar to os/signal. |
| [copy](https://github.com/otiai10/copy) | 142 | 4 | 2017-09-01 | 5 days ago | Copy directory recursively. |
| [bigfile](https://github.com/bigfile/bigfile) | 116 | 9 | 2019-07-15 | 2 weeks ago | A file transfer system, support to manage files with http api, rpc call and ftp client. |
| [go-csv-tag](https://github.com/artonge/go-csv-tag) | 64 | 1 | 2017-06-18 | 3 weeks ago | Load csv file using tag. |
| [opc](https://github.com/qmuntal/opc) | 63 | 1 | 2018-11-06 | 4 weeks ago | Load Open Packaging Conventions (OPC) files for Go. |
| [skywalker](https://github.com/dixonwille/skywalker) | 55 | 3 | 2017-08-01 | 2 years ago | Package to allow one to concurrently go through a filesystem with ease. |
| [vfs](https://github.com/C2FO/vfs) | 47 | 21 | 2017-08-01 | 5 days ago | A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS. |
| [afs](https://github.com/viant/afs) | 43 | 9 | 2019-08-19 | 4 days ago | Abstract File Storage (mem, scp, zip, tar, cloud: s3, gs) for Go. |
| [tarfs](https://github.com/posener/tarfs) | 39 | 2 | 2017-03-10 | 2 days ago | Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files. |
| [go-exiftool](https://github.com/barasher/go-exiftool) | 23 | 2 | 2019-05-12 | 2 months ago | Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...). |
| [go-gtfs](https://github.com/artonge/go-gtfs) | 20 | 2 | 2017-07-09 | 3 weeks ago | Load gtfs files in go. |
| [gut](https://github.com/1set/gut) | 20 | 1 | 2019-10-05 | 2 weeks ago | Simple and reliable package for file operations like copy/move/diff/list on files, directories and symbolic links. |
| [checksum](https://github.com/codingsince1985/checksum) | 17 | 1 | 2014-11-05 | 6 months ago | Compute message digest, like MD5 and SHA256, for large files. |
| [flop](https://github.com/homedepot/flop) | 15 | 11 | 2019-03-01 | 6 months ago | File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html). |
| [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) | 14 | 2 | 2018-10-16 | 2 months ago | Copy files for humans. |
| [parquet](https://github.com/parsyl/parquet) | 6 | 4 | 2019-01-29 | 5 months ago | Read and write [parquet](https://parquet.apache.org) files. |

## Financial
        
*Packages for accounting and finance.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [decimal](https://github.com/shopspring/decimal) | 1981 | 65 | 2015-02-25 | 2 weeks ago | Arbitrary-precision fixed-point decimal numbers. |
| [go-money](https://github.com/Rhymond/go-money) | 719 | 18 | 2017-03-20 | 1 week ago | Implementation of Fowler's Money pattern. |
| [accounting](https://github.com/leekchan/accounting) | 542 | 11 | 2015-08-10 | 2 months ago | money and currency formatting for golang. |
| [go-finance](https://github.com/FlashBoys/go-finance) | 539 | 27 | 2016-02-28 | 2 years ago | Comprehensive financial markets data in Go. |
| [techan](https://github.com/sdcoffey/techan) | 238 | 28 | 2017-03-08 | 2 months ago | Technical analysis library with advanced market analysis and trading strategies. |
| [orderbook](https://github.com/i25959341/orderbook) | 122 | 9 | 2018-04-24 | 10 months ago | Matching Engine for Limit Order Book in Golang. |
| [ofxgo](https://github.com/aclindsa/ofxgo) | 73 | 9 | 2015-11-08 | 8 months ago | Query OFX servers and/or parse the responses (with example command-line client). |
| [transaction](https://github.com/claygod/transaction) | 65 | 8 | 2017-10-11 | 5 months ago | Embedded transactional database of accounts, running in multithreaded mode. |
| [vat](https://github.com/dannyvankooten/vat) | 65 | 1 | 2016-06-18 | 1 year ago | VAT number validation & EU VAT rates. |
| [go-finance](https://github.com/alpeb/go-finance) | 56 | 3 | 2017-06-01 | 1 month ago | Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations. |
| [go-finnhub](https://github.com/m1/go-finnhub) | 21 | 3 | 2020-01-13 | 1 month ago | Client for stock market, forex and crypto data from finnhub.io. Access real-time financial market data from 60+ stock exchanges, 10 forex brokers, and 15+ crypto exchanges. |
| [currency](https://github.com/bnkamalesh/currency) | 19 | 5 | 2017-05-09 | 4 months ago | High performant & accurate currency computation package. |
| [go-finance](https://github.com/pieterclaerhout/go-finance) | 3 | 1 | 2019-09-30 | 4 months ago | Module to fetch exchange rates, check VAT numbers via VIES and check IBAN bank account numbers. |

## Forms
        
*Libraries for working with forms.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [nosurf](https://github.com/justinas/nosurf) | 1037 | 35 | 2013-08-22 | 2 months ago | CSRF protection middleware for Go. |
| [binding](https://github.com/mholt/binding) | 768 | 31 | 2014-05-20 | 1 year ago | Binds form and JSON data from net/http Request to struct. |
| [csrf](https://github.com/gorilla/csrf) | 508 | 21 | 2015-08-03 | 1 month ago | CSRF protection for Go web applications & services. |
| [form](https://github.com/go-playground/form) | 398 | 12 | 2016-05-26 | 3 months ago | Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support. |
| [conform](https://github.com/leebenson/conform) | 184 | 5 | 2016-01-05 | 6 months ago | Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags. |
| [formam](https://github.com/monoculum/formam) | 139 | 3 | 2014-10-25 | 2 months ago | decode form's values into a struct. |
| [forms](https://github.com/albrow/forms) | 104 | 6 | 2014-08-07 | 2 years ago | Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files. |
| [bind](https://github.com/robfig/bind) | 24 | 2 | 2014-08-06 | 5 years ago | Bind form data to any Go values. |
| [queryparam](https://github.com/TomWright/queryparam) | 2 | 1 | 2018-06-14 | 3 months ago | Decode `url.Values` into usable struct values of standard or custom types. |

## Functional
        
*Packages to support functional programming in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-underscore](https://github.com/tobyhede/go-underscore) | 1130 | 28 | 2014-07-02 | 1 year ago | Useful collection of helpfully functional Go collection utilities. |
| [fpGo](https://github.com/TeaEntityLab/fpGo) | 136 | 3 | 2018-05-24 | 1 year ago | Monad, Functional Programming features for Golang. |
| [fuego](https://github.com/seborama/fuego) | 63 | 2 | 2018-11-05 | 7 months ago | Functional Experiment in Go. |

## Game Development
        
*Awesome game development libraries.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [leaf](https://github.com/name5566/leaf) | 3439 | 320 | 2014-08-04 | 4 months ago | Lightweight game server framework. |
| [pixel](https://github.com/faiface/pixel) | 2790 | 98 | 2016-11-19 | 1 month ago | Hand-crafted 2D game library in Go. |
| [ebiten](https://github.com/hajimehoshi/ebiten) | 2564 | 85 | 2013-06-16 | 2 hours ago | dead simple 2D game library in Go. |
| [goworld](https://github.com/xiaonanln/goworld) | 1425 | 114 | 2017-06-03 | 1 month ago | Scalable game server engine, featuring space-entity framework and hot-swapping. |
| [go-sdl2](https://github.com/veandco/go-sdl2) | 1290 | 41 | 2013-06-05 | 5 days ago | Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/). |
| [nano](https://github.com/lonng/nano) | 1212 | 55 | 2017-08-02 | 2 months ago | Lightweight, facility, high performance golang based game server framework. |
| [engo](https://github.com/EngoEngine/engo) | 1176 | 46 | 2014-11-12 | 1 week ago | Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm. |
| [termloop](https://github.com/JoelOtter/termloop) | 1109 | 33 | 2015-05-23 | 4 months ago | Terminal-based game engine for Go, built on top of Termbox. |
| [gonet](https://github.com/xtaci/gonet) | 1088 | 133 | 2013-04-11 | 2 years ago | Game server skeleton implemented with golang. |
| [engine](https://github.com/g3n/engine) | 966 | 64 | 2017-03-07 | 2 weeks ago | Go 3D Game Engine. |
| [oak](https://github.com/oakmound/oak) | 714 | 41 | 2017-07-15 | 2 days ago | Pure Go game engine. |
| [pitaya](https://github.com/topfreegames/pitaya) | 494 | 44 | 2018-03-19 | 3 days ago | Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. |
| [raylib-go](https://github.com/gen2brain/raylib-go) | 444 | 20 | 2017-01-27 | 1 month ago | Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming. |
| [engine](https://github.com/azul3d/engine) | 443 | 23 | 2016-02-29 | 4 months ago | 3D game engine written in Go. |
| [go-astar](https://github.com/beefsack/go-astar) | 357 | 10 | 2014-05-28 | 2 years ago | Go implementation of the A\* path finding algorithm. |
| [GarageEngine](https://github.com/vova616/GarageEngine) | 319 | 31 | 2012-08-07 | 7 months ago | 2d game engine written in Go working on OpenGL. |
| [go3d](https://github.com/ungerik/go3d) | 173 | 10 | 2011-06-27 | 3 weeks ago | Performance oriented 2D/3D math package for Go. |
| [glop](https://github.com/runningwild/glop) | 77 | 3 | 2011-04-20 | 4 years ago | Glop (Game Library Of Power) is a fairly simple cross-platform game library. |
| [go-collada](https://github.com/GlenKelley/go-collada) | 15 | 3 | 2013-09-19 | 6 years ago | Go package for working with the Collada file format. |

## Generation and Generics
        
*Tools to enhance the language with features like generics via code generation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-linq](https://github.com/ahmetb/go-linq) | 1996 | 68 | 2013-12-19 | 2 weeks ago | .NET LINQ-like query methods for Go. |
| [jennifer](https://github.com/dave/jennifer) | 1469 | 23 | 2016-12-04 | 3 months ago | Generate arbitrary Go code without templates. |
| [gen](https://github.com/clipperhouse/gen) | 1093 | 34 | 2013-10-13 | 2 months ago | Code generation tool for ‘generics’-like functionality. |
| [goderive](https://github.com/awalterschulze/goderive) | 789 | 25 | 2017-02-10 | 3 weeks ago | Derives functions from input types. |
| [gowrap](https://github.com/hexdigest/gowrap) | 331 | 11 | 2018-09-15 | 1 week ago | Generate decorators for Go interfaces using simple templates. |
| [interfaces](https://github.com/rjeczalik/interfaces) | 208 | 7 | 2015-12-06 | 1 month ago | Command line tool for generating interface definitions. |
| [go-enum](https://github.com/abice/go-enum) | 101 | 3 | 2017-08-10 | 1 month ago | Code generation for enums from code comments. |
| [pkgreflect](https://github.com/ungerik/pkgreflect) | 92 | 5 | 2012-09-03 | 2 years ago | Go preprocessor for package scoped reflection. |
| [efaceconv](https://github.com/t0pep0/efaceconv) | 44 | 4 | 2016-11-18 | 2 years ago | Code generation tool for high performance conversion from interface{} to immutable type without allocations. |
| [gotype](https://github.com/wzshiming/gotype) | 30 | 4 | 2017-12-05 | 1 month ago | Golang source code parsing, usage like reflect package. |
| [GENERIS](https://github.com/senselogic/GENERIS) | 21 | 2 | 2019-03-10 | 4 months ago | Code generation tool providing generics, free-form macros, conditional compilation and HTML templating. |
| [go-xray](https://github.com/pieterclaerhout/go-xray) | 6 | 1 | 2019-10-01 | 3 months ago | Helpers for making the use of reflection easier. |
| [typeregistry](https://github.com/xiaoxin01/typeregistry) | 3 | 1 | 2020-01-14 | 3 weeks ago | A library to create type dynamically. |

## Geographic
        
*Geographic tools and servers*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [tile38](https://github.com/tidwall/tile38) | 6725 | 208 | 2016-03-04 | 1 week ago | Geolocation DB with spatial index and realtime geofencing. |
| [geo](https://github.com/golang/geo) | 1016 | 75 | 2014-12-03 | 4 weeks ago | S2 geometry library in Go. |
| [mbtileserver](https://github.com/consbio/mbtileserver) | 132 | 12 | 2014-11-01 | 3 weeks ago | A simple Go-based server for map tiles stored in mbtiles format. |
| [geocache](https://github.com/melihmucuk/geocache) | 119 | 6 | 2016-06-21 | 3 years ago | In-memory cache that is suitable for geolocation based applications. |
| [osm](https://github.com/paulmach/osm) | 101 | 8 | 2016-02-02 | 5 months ago | Library for reading, writing and working with OpenStreetMap data and APIs. |
| [wgs84](https://github.com/wroge/wgs84) | 50 | 1 | 2019-06-08 | 3 weeks ago | Library for Coordinate Conversion and Transformation (ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM). |
| [geoserver](https://github.com/hishamkaram/geoserver) | 33 | 1 | 2018-03-26 | 1 month ago | geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API. |
| [gismanager](https://github.com/hishamkaram/gismanager) | 26 | 1 | 2018-09-29 | 1 year ago | Publish Your GIS Data(Vector Data) to PostGIS and Geoserver. |
| [pbf](https://github.com/maguro/pbf) | 20 | 2 | 2017-09-18 | 2 months ago | OpenStreetMap PBF golang encoder/decoder. |

## Go Compilers
        
*Tools for compiling Go to other languages.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gopherjs](https://github.com/gopherjs/gopherjs) | 9231 | 259 | 2013-08-27 | 3 weeks ago | Compiler from Go to JavaScript. |
| [llgo](https://github.com/go-llvm/llgo) | 1042 | 63 | 2011-11-05 | 5 years ago | LLVM-based compiler for Go. |
| [tardisgo](https://github.com/tardisgo/tardisgo) | 396 | 27 | 2014-01-08 | 3 years ago | Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler. |
| [c4go](https://github.com/Konstantin8105/c4go) | 203 | 15 | 2018-03-28 | 1 month ago | Transpile C code to Go code. |
| [f4go](https://github.com/Konstantin8105/f4go) | 23 | 3 | 2018-07-08 | 3 months ago | Transpile FORTRAN 77 code to Go code. |

## Goroutines
        
*Tools for managing and working with Goroutines.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [ants](https://github.com/panjf2000/ants) | 3187 | 102 | 2018-05-19 | 4 days ago | A high-performance and low-cost goroutine pool in Go. |
| [goworker](https://github.com/benmanns/goworker) | 2361 | 77 | 2013-07-22 | 3 weeks ago | goworker is a Go-based background worker. |
| [tunny](https://github.com/Jeffail/tunny) | 1575 | 63 | 2014-04-02 | 5 months ago | Goroutine pool for golang. |
| [pool](https://github.com/go-playground/pool) | 547 | 14 | 2015-10-28 | 6 months ago | Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation. |
| [grpool](https://github.com/ivpusic/grpool) | 544 | 29 | 2015-07-22 | 1 year ago | Lightweight Goroutine pool. |
| [workerpool](https://github.com/gammazero/workerpool) | 270 | 8 | 2016-05-17 | 4 days ago | Goroutine pool that limits the concurrency of task execution, not the number of tasks queued. |
| [go-floc](https://github.com/workanator/go-floc) | 177 | 7 | 2017-07-03 | 2 years ago | Orchestrate goroutines with ease. |
| [gowp](https://github.com/xxjwxc/gowp) | 168 | 10 | 2019-09-14 | 2 months ago | gowp is concurrency limiting goroutine pool. |
| [go-flow](https://github.com/kamildrazkiewicz/go-flow) | 127 | 8 | 2016-09-25 | 10 months ago | Control goroutines execution order. |
| [GoSlaves](https://github.com/dgrr/GoSlaves) | 91 | 4 | 2017-09-17 | 6 months ago | Simple and Asynchronous Goroutine pool library. |
| [semaphore](https://github.com/kamilsk/semaphore) | 82 | 1 | 2016-10-08 | 2 months ago | Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context. |
| [semaphore](https://github.com/marusama/semaphore) | 81 | 4 | 2017-11-22 | 1 year ago | Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations). |
| [gpool](https://github.com/sherifabdlnaby/gpool) | 66 | 1 | 2018-12-03 | 3 months ago | manages a resizeable pool of context-aware goroutines to bound concurrency. |
| [breaker](https://github.com/kamilsk/breaker) | 56 | 3 | 2019-02-15 | 2 weeks ago | Flexible mechanism to make execution flow interruptible. |
| [worker-pool](https://github.com/vardius/worker-pool) | 53 | 4 | 2017-10-04 | 1 month ago | goworker is a Go simple async worker pool. |
| [cyclicbarrier](https://github.com/marusama/cyclicbarrier) | 44 | 3 | 2018-01-11 | 1 year ago | CyclicBarrier for golang. |
| [gollback](https://github.com/vardius/gollback) | 32 | 1 | 2019-05-11 | 1 month ago | asynchronous simple function utilities, for managing execution of closures and callbacks. |
| [async](https://github.com/StudioSol/async) | 30 | 9 | 2017-06-30 | 1 month ago | A safe way to execute functions asynchronously, recovering them in case of panic. |
| [parallel-fn](https://github.com/rafaeljesus/parallel-fn) | 26 | 3 | 2017-06-18 | 2 years ago | Run functions in parallel. |
| [threadpool](https://github.com/shettyh/threadpool) | 26 | 2 | 2017-09-06 | 1 year ago | Golang threadpool implementation. |
| [artifex](https://github.com/mborders/artifex) | 25 | 3 | 2018-10-31 | 4 months ago | Simple in-memory job queue for Golang using worker-based dispatching. |
| [Hunch](https://github.com/AaronJan/Hunch) | 23 | 0 | 2019-06-05 | 7 months ago | Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive. |
| [kyoo](https://github.com/dirkaholic/kyoo) | 23 | 1 | 2020-01-06 | 2 months ago | Provides an unlimited job queue and concurrent worker pools. |
| [gohive](https://github.com/loveleshsharma/gohive) | 11 | 3 | 2019-05-31 | 5 months ago | A highly performant and easy to use Goroutine pool for Go. |
| [stl](https://github.com/ssgreg/stl) | 11 | 1 | 2018-06-19 | 5 months ago | Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism. |
| [routine](https://github.com/x-mod/routine) | 9 | 1 | 2019-03-04 | 3 weeks ago | go routine control with context, support: Main, Go, Pool and some useful Executors. |
| [go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup) | 7 | 1 | 2018-08-08 | 3 weeks ago | Like `sync.WaitGroup` with error handling and concurrency control. |
| [go-tools](https://github.com/nikhilsaraf/go-tools) | 5 | 2 | 2018-11-14 | 11 months ago | Manage a pool of goroutines using this lightweight library with a simple API. |
| [go-trylock](https://github.com/subchen/go-trylock) | 5 | 1 | 2018-04-26 | 2 months ago | TryLock support on read-write lock for Golang. |
| [conexec](https://github.com/ITcathyh/conexec) | 4 | 2 | 2019-12-24 | 1 month ago | A concurrent toolkit to help execute funcs concurrently in an efficient and safe way.It supports specifying the overall timeout to avoid blocking and uses goroutine pool to improve efficiency. |
| [queue](https://github.com/AnikHasibul/queue) | 4 | 0 | 2018-12-21 | 10 months ago | Gives you a `sync.WaitGroup` like queue group accessibility. Helps you to throttle and limit goroutines, wait for the end of the all goroutines and much more. |

## GUI
        
*Interaction*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [fyne](https://github.com/fyne-io/fyne) | 9000 | 191 | 2018-02-04 | 1 hour ago | Cross platform native GUIs designed for Go based on Material Design. Supports: Linux, macOS, Windows, BSD, iOS and Android. |
| [ui](https://github.com/andlabs/ui) | 7395 | 371 | 2014-02-17 | 5 months ago | Platform-native GUI library for Go. Cross platform. |
| [qt](https://github.com/therecipe/qt) | 7086 | 297 | 2014-11-19 | 1 month ago | Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi). |
| [webview](https://github.com/zserge/webview) | 5590 | 203 | 2017-08-19 | 1 week ago | Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux). |
| [robotgo](https://github.com/go-vgo/robotgo) | 4877 | 195 | 2016-09-26 | 1 day ago | Go Native cross-platform GUI system automation. Control the mouse, keyboard and other. |
| [walk](https://github.com/lxn/walk) | 4309 | 251 | 2010-09-16 | 3 months ago | Windows application library kit for Go. |
| [go-app](https://github.com/maxence-charriere/go-app) | 3464 | 110 | 2016-10-12 | 10 hours ago | Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. |
| [app](https://github.com/maxence-charriere/app) | 3314 | 111 | 2016-10-12 | 3 weeks ago | Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. |
| [go-astilectron](https://github.com/asticode/go-astilectron) | 3071 | 133 | 2017-04-22 | 2 weeks ago | Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron). |
| [go-sciter](https://github.com/sciter-sdk/go-sciter) | 1650 | 122 | 2015-10-15 | 2 months ago | Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform. |
| [systray](https://github.com/getlantern/systray) | 1029 | 49 | 2014-11-12 | 4 hours ago | Cross platform Go library to place an icon and menu in the notification area. |
| [gotk3](https://github.com/gotk3/gotk3) | 943 | 50 | 2015-08-13 | 2 days ago | Go bindings for GTK3. |
| [gosx-notifier](https://github.com/deckarep/gosx-notifier) | 512 | 13 | 2013-11-25 | 2 weeks ago | OSX Desktop Notifications library for Go. |
| [gowd](https://github.com/dtylman/gowd) | 243 | 25 | 2017-03-29 | 9 months ago | Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform. |
| [trayhost](https://github.com/shurcooL/trayhost) | 175 | 4 | 2014-04-25 | 4 months ago | Cross-platform Go library to place an icon in the host operating system's taskbar. |
| [go-appindicator](https://github.com/dawidd6/go-appindicator) | 6 | 2 | 2019-05-04 | 6 months ago | Go bindings for libappindicator3 C library. |
| [activity-tracker](https://github.com/prashantgupta24/activity-tracker) | 5 | 1 | 2019-03-12 | 9 months ago | OSX library to notify about any (pluggable) activity on your machine. |
| [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) | 2 | 1 | 2019-03-30 | 9 months ago | OSX Sleep/Wake notifications in golang. |

## Hardware
        
*Libraries, tools, and tutorials for interacting with hardware.*

## Images
        
*Libraries for manipulating images.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gocv](https://github.com/hybridgroup/gocv) | 3013 | 128 | 2017-09-18 | 4 days ago | Go package for computer vision using OpenCV 3.3+. |
| [imaging](https://github.com/disintegration/imaging) | 2973 | 74 | 2012-12-06 | 1 month ago | Simple Go image processing package. |
| [imaginary](https://github.com/h2non/imaginary) | 2901 | 75 | 2015-03-04 | 2 weeks ago | Fast and simple HTTP microservice for image resizing. |
| [bild](https://github.com/anthonynsimon/bild) | 2797 | 59 | 2016-08-01 | 1 month ago | Collection of image processing algorithms in pure Go. |
| [ln](https://github.com/fogleman/ln) | 2630 | 87 | 2016-01-10 | 8 months ago | 3D line art rendering in Go. |
| [resize](https://github.com/nfnt/resize) | 2300 | 80 | 2012-08-02 | 2 years ago | Image resizing for Go with common interpolation methods. |
| [gg](https://github.com/fogleman/gg) | 2178 | 77 | 2016-02-18 | 3 weeks ago | 2D rendering in pure Go. |
| [pt](https://github.com/fogleman/pt) | 1838 | 56 | 2015-01-23 | 1 year ago | Path tracing engine written in Go. |
| [svgo](https://github.com/ajstarks/svgo) | 1438 | 50 | 2010-03-05 | 1 month ago | Go Language Library for SVG generation. |
| [smartcrop](https://github.com/muesli/smartcrop) | 1338 | 31 | 2014-04-07 | 5 months ago | Finds good crops for arbitrary images and crop sizes. |
| [gift](https://github.com/disintegration/gift) | 1304 | 50 | 2014-07-12 | 6 months ago | Package of image processing filters. |
| [picfit](https://github.com/thoas/picfit) | 1203 | 49 | 2014-12-06 | 3 weeks ago | An image resizing server written in Go. |
| [go-opencv](https://github.com/go-opencv/go-opencv) | 1161 | 68 | 2013-12-09 | 9 months ago | Go bindings for OpenCV. |
| [imagick](https://github.com/gographics/imagick) | 1124 | 57 | 2013-04-30 | 1 month ago | Go binding to ImageMagick's MagickWand C API. |
| [geopattern](https://github.com/pravj/geopattern) | 1054 | 21 | 2014-10-22 | 1 year ago | Create beautiful generative image patterns from a string. |
| [bimg](https://github.com/h2non/bimg) | 904 | 34 | 2015-03-17 | 1 month ago | Small package for fast and efficient image processing using libvips. |
| [stegify](https://github.com/DimitarPetrov/stegify) | 797 | 21 | 2018-11-29 | 1 day ago | Go tool for LSB steganography, capable of hiding any file within an image. |
| [canvas](https://github.com/tdewolff/canvas) | 419 | 13 | 2017-05-20 | 2 hours ago | Vector graphics to PDF, SVG or rasterized image. |
| [mort](https://github.com/aldor007/mort) | 388 | 19 | 2017-11-19 | 1 month ago | Storage and image processing server written in Go. |
| [image2ascii](https://github.com/qeesung/image2ascii) | 381 | 7 | 2018-10-20 | 1 year ago | Convert image to ASCII. |
| [govatar](https://github.com/o1egl/govatar) | 347 | 6 | 2016-01-18 | 1 month ago | Library and CMD tool for generating funny avatars. |
| [go-nude](https://github.com/koyachi/go-nude) | 300 | 15 | 2014-05-02 | 1 year ago | Nudity detection with Go. |
| [goimagehash](https://github.com/corona10/goimagehash) | 278 | 9 | 2017-07-28 | 1 month ago | Go Perceptual image hashing package. |
| [rez](https://github.com/bamiaux/rez) | 200 | 9 | 2014-01-16 | 2 years ago | Image resizing in pure Go and SIMD. |
| [img](https://github.com/hawx/img) | 133 | 5 | 2012-07-28 | 4 years ago | Selection of image manipulation tools. |
| [darkroom](https://github.com/gojek/darkroom) | 120 | 6 | 2019-07-01 | 1 week ago | An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency. |
| [go-cairo](https://github.com/ungerik/go-cairo) | 92 | 6 | 2012-08-22 | 5 months ago | Go binding for the cairo graphics library. |
| [mergi](https://github.com/noelyahan/mergi) | 89 | 4 | 2018-09-24 | 10 months ago | Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate). |
| [gltf](https://github.com/qmuntal/gltf) | 61 | 1 | 2019-01-15 | 1 month ago | Efficient and robust glTF 2.0 reader, writer and validator. |
| [go-gd](https://github.com/bolknote/go-gd) | 52 | 4 | 2011-05-12 | 1 year ago | Go binding for GD library. |
| [steganography](https://github.com/auyer/steganography) | 44 | 3 | 2018-05-21 | 1 month ago | Pure Go Library for LSB steganography. |
| [cameron](https://github.com/aofei/cameron) | 40 | 1 | 2018-05-05 | 2 weeks ago | An avatar generator for Go. |
| [goimghdr](https://github.com/corona10/goimghdr) | 34 | 1 | 2018-02-25 | 9 months ago | The imghdr module determines the type of image contained in a file for Go. |
| [tga](https://github.com/ftrvxmtrx/tga) | 26 | 3 | 2012-10-08 | 4 years ago | Package tga is a TARGA image format decoder/encoder. |
| [go-webcolors](https://github.com/jyotiska/go-webcolors) | 24 | 2 | 2014-04-24 | 4 years ago | Port of webcolors library from Python to Go. |
| [mpo](https://github.com/donatj/mpo) | 6 | 1 | 2015-04-14 | 1 year ago | Decoder and conversion tool for MPO 3D Photos. |

## IoT
        
*Libraries for programming devices of the IoT.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [flogo](https://github.com/TIBCOSoftware/flogo) | 1376 | 135 | 2016-07-10 | 1 month ago | Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. |
| [gatt](https://github.com/paypal/gatt) | 884 | 55 | 2014-04-23 | 3 months ago | Gatt is a Go package for building Bluetooth Low Energy peripherals. |
| [mainflux](https://github.com/mainflux/mainflux) | 814 | 77 | 2015-07-06 | 1 day ago | Industrial IoT Messaging and Device Management Server. |
| [devices](https://github.com/goiot/devices) | 229 | 17 | 2016-05-30 | 3 years ago | Suite of libraries for IoT devices, experimental for x/exp/io. |
| [heedy](https://github.com/heedy/heedy) | 203 | 21 | 2015-01-16 | 2 days ago | Open-Source Platform for Quantified Self & IoT. |
| [sensorbee](https://github.com/sensorbee/sensorbee) | 191 | 17 | 2016-02-19 | 4 months ago | Lightweight stream processing engine for IoT. |
| [huego](https://github.com/amimof/huego) | 136 | 2 | 2017-05-16 | 3 weeks ago | An extensive Philips Hue client library for Go. |
| [eywa](https://github.com/xcodersun/eywa) | 44 | 8 | 2016-02-20 | 2 years ago | Project Eywa is essentially a connection manager that keeps track of connected devices. |

## Job Scheduler
        
*Libraries for scheduling jobs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gron](https://github.com/roylee0704/gron) | 695 | 16 | 2016-06-04 | 7 hours ago | Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly. |
| [jobrunner](https://github.com/bamzi/jobrunner) | 663 | 21 | 2015-10-21 | 3 months ago | Smart and featureful cron job scheduler with job queuing and live monitoring built in. |
| [jobs](https://github.com/albrow/jobs) | 466 | 18 | 2015-02-09 | 1 year ago | Persistent and flexible background jobs library. |
| [scheduler](https://github.com/carlescere/scheduler) | 319 | 16 | 2015-02-03 | 1 year ago | Cronjobs scheduling made easy. |
| [go-cron](https://github.com/rk/go-cron) | 184 | 9 | 2011-04-15 | 1 month ago | Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. |
| [leprechaun](https://github.com/kilgaloon/leprechaun) | 61 | 6 | 2018-04-08 | 3 weeks ago | Job scheduler that supports webhooks, crons and classic scheduling. |
| [clockwork](https://github.com/whiteShtef/clockwork) | 1 | 1 | 2018-04-23 | 3 weeks ago | Simple and intuitive job scheduling library in Go. |

## JSON
        
*Libraries for working with JSON.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gjson](https://github.com/tidwall/gjson) | 5993 | 134 | 2016-08-11 | 2 weeks ago | Get a JSON value with one line of code. |
| [gojson](https://github.com/ChimeraCoder/gojson) | 2153 | 44 | 2012-12-27 | 3 months ago | Automatically generate Go (golang) struct definitions from example JSON. |
| [kazaam](https://github.com/qntfy/kazaam) | 146 | 20 | 2016-07-19 | 6 months ago | API for arbitrary transformation of JSON documents. |
| [gojq](https://github.com/elgs/gojq) | 143 | 4 | 2015-12-30 | 1 year ago | JSON query in Golang. |
| [jsongo](https://github.com/ricardolonga/jsongo) | 94 | 1 | 2015-08-07 | 3 years ago | Fluent API to make it easier to create Json objects. |
| [jettison](https://github.com/wI2L/jettison) | 74 | 3 | 2019-08-30 | 3 days ago | High performance, reflection-less JSON encoder for Go. |
| [gjo](https://github.com/skanehira/gjo) | 71 | 7 | 2019-02-23 | 3 months ago | Small utility to create JSON objects. |
| [jaydiff](https://github.com/yazgazan/jaydiff) | 63 | 1 | 2017-04-24 | 5 months ago | JSON diff utility written in Go. |
| [jsonf](https://github.com/miolini/jsonf) | 55 | 3 | 2015-05-25 | 3 years ago | Console tool for highlighted formatting and struct query fetching JSON. |
| [mp](https://github.com/sanbornm/mp) | 38 | 2 | 2014-06-15 | 3 years ago | Simple cli email parser. It currently takes stdin and outputs JSON. |
| [json2go](https://github.com/m-zajac/json2go) | 33 | 2 | 2017-06-10 | 5 months ago | Advanced JSON to Go struct conversion. Provides package that can parse multiple JSON documents and create struct to fit them all. |
| [go-respond](https://github.com/nicklaw5/go-respond) | 30 | 1 | 2017-03-12 | 7 months ago | Go package for handling common HTTP JSON responses. |
| [ajson](https://github.com/spyzhov/ajson) | 26 | 0 | 2019-03-07 | 2 weeks ago | Abstract JSON for golang with JSONPath support. |
| [go-jsonerror](https://github.com/ddymko/go-jsonerror) | 9 | 1 | 2018-10-18 | 5 months ago | Go-JsonError is ment to allow us to easily create json response errors that follow the JsonApi spec. |
| [jsonhal](https://github.com/RichardKnop/jsonhal) | 9 | 2 | 2016-01-15 | 1 year ago | Simple Go package to make custom structs marshal into HAL compatible JSON responses. |
| [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) | 8 | 1 | 2016-07-08 | 3 years ago | Go bindings based on the JSON API errors reference. |
| [ej](https://github.com/lucassscaravelli/ej) | 3 | 1 | 2020-01-04 | 2 months ago | Write and read JSON from different sources succinctly. |
| [mapslice-json](https://github.com/mickep76/mapslice-json) | 1 | 1 | 2020-02-19 | 3 weeks ago | Go MapSlice for ordered marshal/ unmarshal of maps in JSON. |

## Logging
        
*Libraries for generating and working with log files.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [logrus](https://github.com/sirupsen/logrus) | 14163 | 306 | 2013-10-16 | 2 days ago | Structured logger for Go. |
| [zap](https://github.com/uber-go/zap) | 9178 | 224 | 2016-02-18 | 1 day ago | Fast, structured, leveled logging in Go. |
| [go-spew](https://github.com/davecgh/go-spew) | 3769 | 60 | 2013-01-09 | 3 weeks ago | Implements a deep pretty printer for Go data structures to aid in debugging. |
| [zerolog](https://github.com/rs/zerolog) | 3025 | 48 | 2017-05-12 | 5 days ago | Zero-allocation JSON logger. |
| [glog](https://github.com/golang/glog) | 2488 | 90 | 2013-07-16 | 1 month ago | Leveled execution logs for Go. |
| [lumberjack](https://github.com/natefinch/lumberjack) | 1809 | 50 | 2014-06-14 | 2 months ago | Simple rolling logger, implements io.WriteCloser. |
| [tail](https://github.com/hpcloud/tail) | 1735 | 97 | 2013-02-05 | 6 days ago | Go package striving to emulate the features of the BSD tail program. |
| [seelog](https://github.com/cihub/seelog) | 1443 | 90 | 2011-11-17 | 1 year ago | Logging functionality with flexible dispatching, filtering, and formatting. |
| [log15](https://github.com/inconshreveable/log15) | 951 | 26 | 2014-05-20 | 1 month ago | Simple, powerful logging for Go. |
| [log](https://github.com/apex/log) | 811 | 34 | 2015-12-21 | 1 month ago | Structured logging package for Go. |
| [onelog](https://github.com/francoispqt/onelog) | 364 | 9 | 2018-05-06 | 1 year ago | Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenario. Also, it is one of the logger with the lowest allocation. |
| [logxi](https://github.com/mgutz/logxi) | 343 | 10 | 2015-03-01 | 9 months ago | 12-factor app logger that is fast and makes you happy. |
| [log](https://github.com/go-playground/log) | 273 | 10 | 2016-02-07 | 4 months ago | Simple, configurable and scalable Structured Logging for Go. |
| [logutils](https://github.com/hashicorp/logutils) | 272 | 190 | 2013-10-09 | 3 months ago | Utilities for slightly better logging in Go (Golang) extending the standard logger. |
| [go-logger](https://github.com/apsdehal/go-logger) | 248 | 8 | 2014-09-26 | 10 months ago | Simple logger of Go Programs, with level handlers. |
| [logger](https://github.com/azer/logger) | 138 | 5 | 2014-09-30 | 8 months ago | Minimalistic logging library for Go. |
| [xlog](https://github.com/rs/xlog) | 133 | 8 | 2015-10-22 | 1 month ago | Structured logger for `net/context` aware HTTP handlers with flexible dispatching. |
| [rollingwriter](https://github.com/arthurkiller/rollingwriter) | 130 | 7 | 2017-02-12 | 3 weeks ago | RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation. |
| [ozzo-log](https://github.com/go-ozzo/ozzo-log) | 112 | 9 | 2015-10-22 | 1 year ago | High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail). |
| [logvoyage](https://github.com/firstrow/logvoyage) | 84 | 5 | 2015-03-29 | 2 years ago | Full-featured logging saas written in golang. |
| [glg](https://github.com/kpango/glg) | 67 | 5 | 2017-06-21 | 1 day ago | glg is simple and fast leveled logging library for Go. |
| [log](https://github.com/alexcesaro/log) | 42 | 5 | 2014-04-19 | 4 years ago | Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs. |
| [gologger](https://github.com/sadlil/gologger) | 38 | 5 | 2015-09-02 | 2 years ago | Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch. |
| [logex](https://github.com/chzyer/logex) | 35 | 7 | 2014-10-10 | 3 years ago | Golang log lib, supports tracking and level, wrap by standard log lib. |
| [go-log](https://github.com/ian-kent/go-log) | 34 | 2 | 2014-05-02 | 1 year ago | Log4j implementation in Go. |
| [go-cronowriter](https://github.com/utahta/go-cronowriter) | 27 | 1 | 2017-02-04 | 6 months ago | Simple writer that rotate log files automatically based on current date and time, like cronolog. |
| [go-log](https://github.com/siddontang/go-log) | 26 | 5 | 2014-05-18 | 1 year ago | Log lib supports level and multi handlers. |
| [logrusly](https://github.com/sebest/logrusly) | 25 | 5 | 2014-09-11 | 2 years ago | [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/). |
| [log](https://github.com/teris-io/log) | 23 | 1 | 2017-10-28 | 2 years ago | Structured log interface for Go cleanly separates logging facade from its implementation. |
| [distillog](https://github.com/amoghe/distillog) | 22 | 1 | 2015-10-12 | 1 year ago | distilled levelled logging (think of it as stdlib + log levels). |
| [journald](https://github.com/ssgreg/journald) | 22 | 2 | 2017-08-23 | 1 year ago | Go implementation of systemd Journal's native API for logging. |
| [mlog](https://github.com/jbrodriguez/mlog) | 19 | 1 | 2014-10-20 | 1 year ago | Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output. |
| [gomol](https://github.com/aphistic/gomol) | 16 | 2 | 2015-08-30 | 1 year ago | Multiple-output, structured logging for Go with extensible logging outputs. |
| [go-log](https://github.com/subchen/go-log) | 10 | 2 | 2017-05-07 | 1 year ago | Simple and configurable Logging in Go, with level, formatters and writers. |
| [glo](https://github.com/lajosbencz/glo) | 9 | 1 | 2019-01-19 | 1 year ago | PHP Monolog inspired logging facility with identical severity levels. |
| [logdump](https://github.com/ewwwwwqm/logdump) | 9 | 2 | 2017-01-13 | 1 year ago | Package for multi-level logging. |
| [logrusiowriter](https://github.com/cabify/logrusiowriter) | 7 | 86 | 2019-08-09 | 2 months ago | `io.Writer` implementation using [logrus](https://github.com/sirupsen/logrus) logger. |
| [log](https://github.com/aerogo/log) | 6 | 1 | 2017-06-10 | 4 months ago | An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection). |
| [logmatic](https://github.com/borderstech/logmatic) | 6 | 1 | 2018-11-07 | 1 year ago | Colorized logger for Golang with dynamic log level configuration. |
| [xlog](https://github.com/xfxdev/xlog) | 6 | 1 | 2016-05-05 | 1 year ago | Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format. |
| [logo](https://github.com/mbndr/logo) | 5 | 1 | 2017-02-07 | 2 years ago | Golang logger to different configurable writers. |
| [go-log](https://github.com/pieterclaerhout/go-log) | 3 | 1 | 2019-10-01 | 3 weeks ago | A logging library with strack traces, object dumping and optional timestamps. |

## Machine Learning
        
*Libraries for Machine Learning.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [golearn](https://github.com/sjwhitworth/golearn) | 7097 | 430 | 2013-12-26 | 1 week ago | General Machine Learning library for Go. |
| [gorgonia](https://github.com/gorgonia/gorgonia) | 3201 | 172 | 2016-09-14 | 2 weeks ago | graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. |
| [tfgo](https://github.com/galeone/tfgo) | 1324 | 47 | 2017-05-23 | 2 months ago | Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python. |
| [gosseract](https://github.com/otiai10/gosseract) | 1083 | 39 | 2013-10-11 | 3 weeks ago | Go package for OCR (Optical Character Recognition), by using Tesseract C++ library. |
| [goml](https://github.com/cdipaolo/goml) | 1082 | 75 | 2015-06-27 | 11 months ago | On-line Machine Learning in Go. |
| [gorse](https://github.com/zhenghaoz/gorse) | 934 | 42 | 2018-08-14 | 2 weeks ago | An offline recommender system backend based on collaborative filtering written in Go. |
| [CloudForest](https://github.com/ryanbressler/CloudForest) | 667 | 44 | 2012-10-22 | 1 year ago | Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go. |
| [eaopt](https://github.com/MaxHalford/eaopt) | 667 | 29 | 2016-01-31 | 1 month ago | An evolutionary optimization library. |
| [bayesian](https://github.com/jbrukh/bayesian) | 655 | 32 | 2011-11-23 | 1 year ago | Naive Bayesian Classification for Golang. |
| [gobrain](https://github.com/goml/gobrain) | 433 | 26 | 2014-04-29 | 4 months ago | Neural Networks written in go. |
| [ocrserver](https://github.com/otiai10/ocrserver) | 274 | 10 | 2015-11-15 | 1 year ago | A simple OCR API server, seriously easy to be deployed by Docker and Heroku. |
| [regommend](https://github.com/muesli/regommend) | 266 | 15 | 2014-02-05 | 7 months ago | Recommendation & collaborative filtering engine. |
| [go-deep](https://github.com/patrikeh/go-deep) | 256 | 13 | 2017-12-09 | 3 months ago | A feature-rich neural network library in Go. |
| [onnx-go](https://github.com/owulveryck/onnx-go) | 250 | 10 | 2018-08-28 | 3 weeks ago | Go Interface to Open Neural Network Exchange (ONNX). |
| [go-galib](https://github.com/thoj/go-galib) | 177 | 13 | 2009-11-30 | 4 years ago | Genetic Algorithms library written in Go / golang. |
| [goRecommend](https://github.com/timkaye11/goRecommend) | 155 | 9 | 2014-07-16 | 5 years ago | Recommendation Algorithms library written in Go. |
| [shield](https://github.com/eaigner/shield) | 129 | 10 | 2013-04-10 | 1 week ago | Bayesian text classifier with flexible tokenizers and storage backends for Go. |
| [goptuna](https://github.com/c-bata/goptuna) | 119 | 3 | 2019-07-24 | 12 hours ago | Bayesian optimization framework for black-box functions written in Go. Everything will be optimized. |
| [go-fann](https://github.com/vksnk/go-fann) | 103 | 7 | 2011-03-10 | 5 years ago | Go bindings for Fast Artificial Neural Networks(FANN) library. |
| [goga](https://github.com/tomcraven/goga) | 86 | 7 | 2015-10-20 | 3 years ago | Genetic algorithm library for Go. |
| [libsvm](https://github.com/datastream/libsvm) | 64 | 10 | 2012-07-31 | 3 years ago | libsvm golang version derived work based on LIBSVM 3.14. |
| [neural-go](https://github.com/schuyler/neural-go) | 61 | 2 | 2011-10-17 | 6 years ago | Multilayer perceptron network implemented in Go, with training via backpropagation. |
| [go-pr](https://github.com/daviddengcn/go-pr) | 59 | 6 | 2013-06-07 | 6 years ago | Pattern recognition package in Go lang. |
| [gonet](https://github.com/dathoangnd/gonet) | 58 | 4 | 2020-01-11 | 1 month ago | Neural Network for Go. |
| [neat](https://github.com/jinyeom/neat) | 56 | 12 | 2016-11-17 | 1 year ago | Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT). |
| [goscore](https://github.com/asafschers/goscore) | 45 | 3 | 2017-08-19 | 6 months ago | Go Scoring API for PMML. |
| [golinear](https://github.com/danieldk/golinear) | 39 | 5 | 2013-04-05 | 1 year ago | liblinear bindings for Go. |
| [fonet](https://github.com/Fontinalis/fonet) | 36 | 4 | 2017-10-03 | 6 months ago | A Deep Neural Network library written in Go. |
| [Varis](https://github.com/Xamber/Varis) | 28 | 5 | 2017-10-10 | 1 year ago | Golang Neural Network. |
| [godist](https://github.com/e-dard/godist) | 25 | 2 | 2014-09-05 | 4 years ago | Various probability distributions, and associated methods. |
| [go-cluster](https://github.com/e-XpertSolutions/go-cluster) | 22 | 5 | 2017-10-04 | 1 year ago | Go implementation of the k-modes and k-prototypes clustering algorithms. |
| [probab](https://github.com/ThePaw/probab) | 12 | 1 | 2015-09-14 | 4 years ago | Probability distribution functions. Bayesian inference. Written in pure Go. |
| [evoli](https://github.com/khezen/evoli) | 9 | 3 | 2015-06-12 | 5 days ago | Genetic Algorithm and Particle Swarm Optimization library. |
| [gomind](https://github.com/surenderthakran/gomind) | 7 | 2 | 2017-10-19 | 1 year ago | A simplistic Neural Network Library in Go. |
| [randomForest](https://github.com/malaschitz/randomForest) | 3 | 1 | 2018-10-25 | 3 months ago | Easy to use Random Forest library for Go. |

## Messaging
        
*Libraries that implement messaging systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [sarama](https://github.com/Shopify/sarama) | 5493 | 394 | 2013-07-05 | 2 days ago | Go library for Apache Kafka. |
| [gorush](https://github.com/appleboy/gorush) | 4242 | 175 | 2016-03-22 | 1 day ago | Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm). |
| [centrifugo](https://github.com/centrifugal/centrifugo) | 4132 | 184 | 2015-03-31 | 1 day ago | Real-time messaging (Websockets or SockJS) server in Go. |
| [machinery](https://github.com/RichardKnop/machinery) | 3892 | 139 | 2015-04-05 | 1 week ago | Asynchronous task queue/job queue based on distributed message passing. |
| [go-socket.io](https://github.com/googollee/go-socket.io) | 3299 | 133 | 2013-07-13 | 1 week ago | socket.io library for golang, a realtime application framework. |
| [nats.go](https://github.com/nats-io/nats.go) | 2709 | 157 | 2012-08-15 | 3 days ago | Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library. |
| [apns2](https://github.com/sideshow/apns2) | 2269 | 72 | 2016-01-05 | 1 month ago | HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps. |
| [benthos](https://github.com/Jeffail/benthos) | 2251 | 65 | 2016-03-22 | 9 hours ago | A message streaming bridge between a range of protocols. |
| [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) | 1895 | 236 | 2013-12-27 | 2 years ago | gopush-cluster is a go push server cluster. |
| [mercure](https://github.com/dunglas/mercure) | 1890 | 53 | 2018-07-14 | 3 weeks ago | Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events). |
| [melody](https://github.com/olahol/melody) | 1763 | 54 | 2015-05-13 | 2 months ago | Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling. |
| [go-nsq](https://github.com/nsqio/go-nsq) | 1602 | 61 | 2013-08-29 | 2 weeks ago | the official Go package for NSQ. |
| [mangos-v1](https://github.com/nanomsg/mangos-v1) | 1537 | 84 | 2014-10-25 | 4 months ago | Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability. |
| [uniqush-push](https://github.com/uniqush/uniqush-push) | 1143 | 75 | 2011-08-29 | 3 months ago | Redis backed unified push service for server-side notifications to mobile devices. |
| [zmq4](https://github.com/pebbe/zmq4) | 837 | 47 | 2013-10-18 | 2 weeks ago | Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2). |
| [gollum](https://github.com/trivago/gollum) | 835 | 38 | 2015-06-20 | 4 months ago | A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations. |
| [Beaver](https://github.com/Clivern/Beaver) | 814 | 18 | 2018-10-20 | 3 hours ago | A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps. |
| [EventBus](https://github.com/asaskevich/EventBus) | 656 | 28 | 2014-12-19 | 2 months ago | The lightweight event bus with async compatibility. |
| [golongpoll](https://github.com/jcuga/golongpoll) | 445 | 23 | 2015-11-02 | 1 week ago | HTTP longpoll server library that makes web pub-sub simple. |
| [dbus](https://github.com/godbus/dbus) | 429 | 16 | 2014-03-27 | 1 month ago | Native Go bindings for D-Bus. |
| [asynq](https://github.com/hibiken/asynq) | 419 | 6 | 2019-11-15 | 13 hours ago | A simple, reliable, and efficient distributed task queue for Go built on top of Redis. |
| [emitter](https://github.com/olebedev/emitter) | 344 | 9 | 2015-11-10 | 1 month ago | Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins. |
| [glue](https://github.com/desertbit/glue) | 332 | 14 | 2015-06-07 | 8 months ago | Robust Go and Javascript Socket Library (Alternative to Socket.io). |
| [pubsub](https://github.com/cskr/pubsub) | 310 | 8 | 2012-04-01 | 11 months ago | Simple pubsub package for go. |
| [guble](https://github.com/smancke/guble) | 143 | 13 | 2015-11-15 | 2 years ago | Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence. |
| [bus](https://github.com/mustafaturan/bus) | 142 | 5 | 2019-04-27 | 1 month ago | Minimalist message bus implementation for internal communication. |
| [rabtap](https://github.com/jandelgado/rabtap) | 111 | 8 | 2017-11-11 | 1 month ago | RabbitMQ swiss army knife cli app. |
| [oplog](https://github.com/dailymotion/oplog) | 102 | 93 | 2014-11-06 | 4 years ago | Generic oplog/replication system for REST APIs. |
| [message-bus](https://github.com/vardius/message-bus) | 99 | 4 | 2017-10-04 | 2 days ago | messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD. |
| [rabbus](https://github.com/rafaeljesus/rabbus) | 69 | 9 | 2017-05-07 | 7 months ago | A tiny wrapper over amqp exchanges and queues. |
| [drone-line](https://github.com/appleboy/drone-line) | 67 | 5 | 2016-09-13 | 1 month ago | Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI. |
| [RapidMQ](https://github.com/sybrexsys/RapidMQ) | 58 | 4 | 2016-10-04 | 2 years ago | RapidMQ is a lightweight and reliable library for managing of the local messages queue. |
| [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) | 57 | 9 | 2017-01-15 | 2 years ago | A tiny wrapper around NSQ topic and channel. |
| [go-notify](https://github.com/TheCreeper/go-notify) | 48 | 2 | 2015-03-01 | 1 year ago | Native implementation of the freedesktop notification spec. |
| [hub](https://github.com/leandro-lugaresi/hub) | 48 | 1 | 2018-04-13 | 2 months ago | A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges. |
| [commander](https://github.com/jeroenrinzema/commander) | 40 | 1 | 2018-04-20 | 1 month ago | A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka. |
| [commander](https://github.com/CloudProud/commander) | 38 | 1 | 2018-04-20 | 3 months ago | A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka. |
| [event](https://github.com/agoalofalife/event) | 33 | 3 | 2017-07-02 | 2 years ago | Implementation of the pattern observer. |
| [go-vitotrol](https://github.com/maxatome/go-vitotrol) | 13 | 6 | 2016-11-03 | 10 months ago | Client library to Viessmann Vitotrol web service. |
| [redisqueue](https://github.com/robinjoseph08/redisqueue) | 12 | 2 | 2019-07-07 | 7 months ago | redisqueue provides a producer and consumer of a queue that uses Redis streams. |
| [jazz](https://github.com/socifi/jazz) | 9 | 3 | 2018-10-22 | 1 year ago | A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages. |
| [gaurun-client](https://github.com/osamingo/gaurun-client) | 8 | 1 | 2017-06-29 | 10 months ago | Gaurun Client written in Go. |
| [ami](https://github.com/kak-tus/ami) | 6 | 1 | 2018-10-27 | 1 week ago | Go client to reliable queues based on Redis Cluster Streams. |
| [rmqconn](https://github.com/sbabiv/rmqconn) | 5 | 0 | 2019-01-14 | 1 month ago | RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed. |

## Microsoft Office
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [unioffice](https://github.com/unidoc/unioffice) | 2147 | 60 | 2017-08-29 | 3 days ago | Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents. |

### Microsoft Excel
        
*Libraries for working with Microsoft Excel.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [excelize](https://github.com/360EntSecGroup-Skylar/excelize) | 5823 | 168 | 2016-08-29 | 1 day ago | Golang library for reading and writing Microsoft Excel™ (XLSX) files. |
| [xlsx](https://github.com/tealeg/xlsx) | 3952 | 184 | 2011-06-28 | 5 days ago | Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs. |
| [xlsx](https://github.com/plandem/xlsx) | 106 | 11 | 2017-08-26 | 3 months ago | Fast and safe way to read/update your existing Microsoft Excel files in Go programs. |
| [go-excel](https://github.com/szyhf/go-excel) | 71 | 3 | 2017-09-03 | 2 months ago | A simple and light reader to read a relate-db-like excel as a table. |
| [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) | 15 | 2 | 2017-03-13 | 1 year ago | Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files. |

## Miscellaneous
        

### Dependency Injection
        
*Libraries for working with dependency injection.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [dig](https://github.com/uber-go/dig) | 1302 | 37 | 2017-03-21 | 1 week ago | A reflection based dependency injection toolkit for Go. |
| [fx](https://github.com/uber-go/fx) | 1194 | 54 | 2016-10-27 | 3 weeks ago | A dependency injection based application framework for Go (built on top of dig). |
| [container](https://github.com/golobby/container) | 77 | 2 | 2019-09-23 | 2 months ago | A powerful IoC Container with fluent and easy-to-use interface. |
| [inject](https://github.com/defval/inject) | 53 | 1 | 2019-02-03 | 3 weeks ago | A reflection based dependency injection container with simple interface. |
| [dingo](https://github.com/i-love-flamingo/dingo) | 41 | 21 | 2018-10-29 | 2 weeks ago | A dependency injection toolkit for Go, based on Guice. |
| [alice](https://github.com/magic003/alice) | 38 | 2 | 2017-04-08 | 2 years ago | Additive dependency injection container for Golang. |
| [wire](https://github.com/Fs02/wire) | 29 | 1 | 2018-07-05 | 1 month ago | Strict Runtime Dependency Injection for Golang. |
| [linker](https://github.com/logrange/linker) | 21 | 2 | 2018-12-04 | 5 days ago | A reflection based dependency injection and inversion of control library with components lifecycle support. |
| [gocontainer](https://github.com/vardius/gocontainer) | 12 | 0 | 2019-06-06 | 1 month ago | Simple Dependency Injection Container. |

### Project Layout
        
*Unofficial set of patterns for structuring projects.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [project-layout](https://github.com/golang-standards/project-layout) | 13672 | 370 | 2017-09-09 | 3 months ago | Set of common historical and emerging project layout patterns in the Go ecosystem. |
| [go-restful-api](https://github.com/qiangxue/go-restful-api) | 1038 | 52 | 2016-08-15 | 1 month ago | An idiomatic Go RESTful API starter kit following SOLID principles and Clean Architecture with a common project layout. |
| [modern-go-application](https://github.com/sagikazarmark/modern-go-application) | 564 | 13 | 2018-09-14 | 2 weeks ago | Go application boilerplate and example applying modern practices. |
| [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) | 324 | 7 | 2016-12-18 | 1 week ago | A Go application boilerplate template for quick starting projects following production best practices. |
| [scaffold](https://github.com/catchplay/scaffold) | 58 | 2 | 2018-12-11 | 1 year ago | Scaffold generates starter Go project layout. Lets you focus on business logic implemeted. |
| [go-sample](https://github.com/zitryss/go-sample) | 46 | 1 | 2019-01-24 | 1 year ago | A sample layout for Go application projects with the real code. |

### Strings
        
*These libraries were placed here because none of the other categories seemed to fit.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gopsutil](https://github.com/shirou/gopsutil) | 4652 | 193 | 2014-04-18 | 2 days ago | Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc). |
| [archiver](https://github.com/mholt/archiver) | 2753 | 46 | 2016-04-08 | 4 days ago | Library and command for making and extracting .zip and .tar.gz archives. |
| [gosms](https://github.com/haxpax/gosms) | 1273 | 56 | 2015-01-23 | 2 years ago | Your own local SMS gateway in Go that can be used to send SMS. |
| [go-resiliency](https://github.com/eapache/go-resiliency) | 982 | 28 | 2014-11-29 | 3 months ago | Resiliency patterns for golang. |
| [gofakeit](https://github.com/brianvoe/gofakeit) | 837 | 7 | 2015-04-24 | 3 weeks ago | Random data generator written in go. |
| [base64Captcha](https://github.com/mojocn/base64Captcha) | 785 | 45 | 2017-12-12 | 13 hours ago | Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha. |
| [go-commons-pool](https://github.com/jolestar/go-commons-pool) | 775 | 44 | 2015-12-28 | 4 months ago | Generic object pool for Golang. |
| [xstrings](https://github.com/huandu/xstrings) | 719 | 25 | 2015-01-06 | 2 months ago | Collection of useful string functions ported from other languages. |
| [shortid](https://github.com/teris-io/shortid) | 530 | 8 | 2016-01-04 | 2 years ago | Distributed generation of super short, unique, non-sequential, URL friendly IDs. |
| [llvm](https://github.com/llir/llvm) | 503 | 30 | 2014-09-19 | 4 days ago | Library for interacting with LLVM IR in pure Go. |
| [health](https://github.com/dimiro1/health) | 380 | 6 | 2016-03-08 | 4 months ago | Easy to use, extensible health check library. |
| [go-conv](https://github.com/cstockton/go-conv) | 350 | 8 | 2016-10-11 | 2 years ago | Package conv provides fast and intuitive conversions across Go types. |
| [banner](https://github.com/dimiro1/banner) | 256 | 4 | 2016-03-25 | 4 months ago | Add beautiful banners into your Go applications. |
| [gountries](https://github.com/pariz/gountries) | 235 | 7 | 2016-01-13 | 6 days ago | Package that exposes country and subdivision data. |
| [antch](https://github.com/antchfx/antch) | 169 | 13 | 2017-09-28 | 1 year ago | A fast, powerful and extensible web crawling & scraping framework. |
| [ffmt](https://github.com/go-ffmt/ffmt) | 155 | 4 | 2015-02-14 | 2 months ago | Beautify data display for Humans. |
| [stateless](https://github.com/qmuntal/stateless) | 148 | 4 | 2019-09-11 | 4 months ago | A fluent library for creating state machines. |
| [ghorg](https://github.com/gabrie30/ghorg) | 145 | 6 | 2018-03-29 | 2 weeks ago | Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, and Bitbucket. |
| [lk](https://github.com/hyperboloide/lk) | 144 | 5 | 2016-07-14 | 3 months ago | A simple licensing library for golang. |
| [battery](https://github.com/distatus/battery) | 143 | 2 | 2016-03-12 | 4 months ago | Cross-platform, normalized battery information library. |
| [stats](https://github.com/go-playground/stats) | 131 | 3 | 2015-09-14 | 3 years ago | Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... |
| [bitio](https://github.com/icza/bitio) | 116 | 6 | 2016-05-31 | 3 months ago | Highly optimized bit-level Reader and Writer for Go. |
| [healthcheck](https://github.com/etherlabsio/healthcheck) | 107 | 5 | 2017-08-18 | 2 months ago | An opinionated and concurrent health-check HTTP handler for RESTful services. |
| [turtle](https://github.com/hackebrot/turtle) | 95 | 1 | 2017-09-08 | 4 months ago | Emojis for Go. |
| [go-unarr](https://github.com/gen2brain/go-unarr) | 87 | 5 | 2015-11-01 | 3 weeks ago | Decompression library for RAR, TAR, ZIP and 7z archives. |
| [strutil](https://github.com/ozgio/strutil) | 86 | 1 | 2018-08-16 | 5 months ago | String utilities. |
| [gommit](https://github.com/antham/gommit) | 85 | 3 | 2016-08-30 | 3 weeks ago | Analyze git commit messages to ensure they follow defined patterns. |
| [gotoprom](https://github.com/cabify/gotoprom) | 80 | 87 | 2018-10-10 | 1 month ago | Type-safe metrics builder wrapper library for the official Prometheus client. |
| [indigo](https://github.com/osamingo/indigo) | 58 | 1 | 2016-08-31 | 2 months ago | Distributed unique ID generator of using Sonyflake and encoded by Base58. |
| [morse](https://github.com/alwindoss/morse) | 57 | 2 | 2018-08-15 | 1 year ago | Library to convert to and from morse code. |
| [captcha](https://github.com/steambap/captcha) | 54 | 5 | 2017-09-12 | 3 weeks ago | Package captcha provides an easy to use, unopinionated API for captcha generation. |
| [xkg](https://github.com/go-xkg/xkg) | 48 | 1 | 2015-01-05 | 5 years ago | X Keyboard Grabber. |
| [persian](https://github.com/mavihq/persian) | 40 | 1 | 2017-10-16 | 1 year ago | Some utilities for Persian language in go. |
| [pdfgen](https://github.com/hyperboloide/pdfgen) | 39 | 3 | 2015-11-30 | 2 years ago | HTTP service to generate PDF from Json requests. |
| [browscap_go](https://github.com/digitalcrab/browscap_go) | 31 | 6 | 2014-09-18 | 3 months ago | GoLang Library for [Browser Capabilities Project](http://browscap.org/). |
| [datacounter](https://github.com/miolini/datacounter) | 31 | 1 | 2015-10-14 | 1 month ago | Go counters for readers/writer/http.ResponseWriter. |
| [autoflags](https://github.com/artyom/autoflags) | 26 | 3 | 2014-05-15 | 1 year ago | Go package to automatically define command line flags from struct fields. |
| [xdg](https://github.com/rkoesters/xdg) | 21 | 1 | 2013-12-15 | 1 year ago | FreeDesktop.org (xdg) Specs implemented in Go. |
| [gosh](https://github.com/osamingo/gosh) | 20 | 0 | 2018-05-25 | 2 months ago | Provide Go Statistics Handler, Struct, Measure Method. |
| [url-shortener](https://github.com/pantrif/url-shortener) | 19 | 1 | 2018-06-04 | 1 year ago | A modern, powerful, and robust URL shortener microservice with mysql support. |
| [sandid](https://github.com/aofei/sandid) | 15 | 1 | 2018-06-12 | 2 weeks ago | Every grain of sand on earth has its own ID. |
| [anagent](https://github.com/mudler/anagent) | 11 | 2 | 2017-12-29 | 1 year ago | Minimalistic, pluggable Golang evloop/timer handler with dependency-injection. |
| [avgRating](https://github.com/kirillDanshin/avgRating) | 10 | 1 | 2017-08-05 | 2 years ago | Calculate average score and rating based on Wilson Score Equation. |
| [hostutils](https://github.com/Wing924/hostutils) | 8 | 1 | 2017-09-26 | 1 year ago | A golang library for packing and unpacking FQDNs list. |
| [shellwords](https://github.com/Wing924/shellwords) | 8 | 1 | 2017-09-28 | 2 years ago | A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell. |
| [metrics](https://github.com/pascaldekloe/metrics) | 6 | 1 | 2019-01-29 | 3 weeks ago | Library for metrics instrumentation and Prometheus exposition. |
| [numa](https://github.com/lrita/numa) | 2 | 1 | 2018-12-10 | 10 months ago | NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code. |

## Natural Language Processing
        
*Libraries for working with human languages.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [prose](https://github.com/jdkato/prose) | 2449 | 61 | 2017-02-17 | 3 months ago | Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. English only. |
| [gse](https://github.com/go-ego/gse) | 1220 | 45 | 2017-06-23 | 21 hours ago | Go efficient text segmentation; support english, chinese, japanese and other. |
| [when](https://github.com/olebedev/when) | 1064 | 25 | 2016-12-27 | 3 months ago | Natural EN and RU language date/time parser with pluggable rules. |
| [gojieba](https://github.com/yanyiwu/gojieba) | 999 | 65 | 2015-09-12 | 5 days ago | This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm. |
| [go-pinyin](https://github.com/mozillazg/go-pinyin) | 661 | 33 | 2014-11-09 | 4 days ago | CN Hanzi to Hanyu Pinyin converter. |
| [kagome](https://github.com/ikawaha/kagome) | 474 | 22 | 2014-06-26 | 1 week ago | JP morphological analyzer written in pure Go. |
| [whatlanggo](https://github.com/abadojack/whatlanggo) | 398 | 15 | 2017-02-20 | 1 year ago | Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc). |
| [nlp](https://github.com/shixzie/nlp) | 363 | 24 | 2017-01-25 | 2 years ago | Extract values from strings and fill your structs with nlp. |
| [sentences](https://github.com/neurosnap/sentences) | 268 | 12 | 2015-08-07 | 11 months ago | Sentence tokenizer:  converts text into a list of sentences. |
| [nlp](https://github.com/james-bowman/nlp) | 246 | 22 | 2017-03-15 | 3 months ago | Go Natural Language Processing library supporting LSA (Latent Semantic Analysis). |
| [getlang](https://github.com/rylans/getlang) | 86 | 3 | 2018-03-01 | 9 months ago | Fast natural language detection package. |
| [go-nlp](https://github.com/nuance/go-nlp) | 82 | 7 | 2011-05-02 | 8 years ago | Utilities for working with discrete probability distributions and other tools useful for doing NLP work. |
| [gounidecode](https://github.com/fiam/gounidecode) | 68 | 3 | 2012-05-01 | 4 years ago | Unicode transliterator (also known as unidecode) for Go. |
| [go-unidecode](https://github.com/mozillazg/go-unidecode) | 67 | 2 | 2016-07-08 | 11 months ago | ASCII transliterations of Unicode text. |
| [textcat](https://github.com/pebbe/textcat) | 61 | 6 | 2012-09-21 | 1 year ago | Go package for n-gram based text categorization, with support for utf-8 and raw text. |
| [MMSEGO](https://github.com/awsong/MMSEGO) | 58 | 5 | 2012-04-18 | 8 years ago | This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm. |
| [go-stem](https://github.com/agonopol/go-stem) | 56 | 3 | 2011-09-23 | 1 year ago | Implementation of the porter stemming algorithm. |
| [RAKE.Go](https://github.com/afjoseph/RAKE.Go) | 56 | 6 | 2016-12-17 | 2 weeks ago | Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE). |
| [stemmer](https://github.com/dchest/stemmer) | 49 | 4 | 2011-03-21 | 3 years ago | Stemmer packages for Go programming language. Includes English and German stemmers. |
| [porter2](https://github.com/zentures/porter2) | 37 | 2 | 2015-01-21 | 4 years ago | Really fast Porter 2 stemmer. |
| [go2vec](https://github.com/danieldk/go2vec) | 34 | 6 | 2015-01-27 | 1 year ago | Reader and utility functions for word2vec embeddings. |
| [petrovich](https://github.com/striker2000/petrovich) | 29 | 1 | 2016-12-26 | 1 month ago | Petrovich is the library which inflects Russian names to given grammatical case. |
| [paicehusk](https://github.com/rookii/paicehusk) | 26 | 3 | 2012-09-29 | 6 years ago | Golang implementation of the Paice/Husk Stemming Algorithm. |
| [snowball](https://github.com/goodsign/snowball) | 25 | 0 | 2012-12-11 | 2 years ago | Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/). |
| [mystem](https://github.com/dveselov/mystem) | 23 | 3 | 2016-08-30 | 3 years ago | CGo bindings to Yandex.Mystem - russian morphology analyzer. |
| [icu](https://github.com/goodsign/icu) | 19 | 0 | 2012-12-11 | 3 years ago | Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1. |
| [golibstemmer](https://github.com/rjohnsondev/golibstemmer) | 17 | 1 | 2012-08-06 | 5 years ago | Go bindings for the snowball libstemmer library including porter 2. |
| [shamoji](https://github.com/osamingo/shamoji) | 11 | 2 | 2017-07-23 | 2 months ago | The shamoji is word filtering package written in Go. |
| [libtextcat](https://github.com/goodsign/libtextcat) | 10 | 1 | 2012-12-10 | 7 years ago | Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2. |
| [porter](https://github.com/a2800276/porter) | 8 | 1 | 2013-09-17 | 6 years ago | This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm. |
| [gotokenizer](https://github.com/xujiajun/gotokenizer) | 7 | 1 | 2018-10-11 | 11 months ago | A tokenizer based on the dictionary and Bigram language models for Go. (Now only support chinese segmentation) |
| [go-localize](https://github.com/m1/go-localize) | 4 | 1 | 2019-12-23 | 2 months ago | Simple and easy to use i18n (Internationalization and localization) engine - used for translating locale strings. |
| [detectlanguage-go](https://github.com/detectlanguage/detectlanguage-go) | 1 | 0 | 2019-12-14 | 2 months ago | Language Detection API Go Client. Supports batch requests, short phrase or single word language detection. |

## Networking
        
*Libraries for working with various layers of the network.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [fasthttp](https://github.com/valyala/fasthttp) | 11728 | 368 | 2015-10-18 | 11 hours ago | Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http. |
| [kcptun](https://github.com/xtaci/kcptun) | 11609 | 596 | 2016-02-26 | 1 week ago | Extremely simple & fast udp tunnel based on KCP protocol. |
| [dns](https://github.com/miekg/dns) | 4388 | 168 | 2010-08-03 | 5 hours ago | Go library for working with DNS. |
| [quic-go](https://github.com/lucas-clemente/quic-go) | 3853 | 163 | 2016-04-06 | 3 days ago | An implementation of the QUIC protocol in pure Go. |
| [httplab](https://github.com/gchaincl/httplab) | 3533 | 67 | 2017-02-08 | 9 months ago | HTTPLabs let you inspect HTTP requests and forge responses. |
| [gopacket](https://github.com/google/gopacket) | 3296 | 134 | 2015-03-16 | 2 days ago | Go library for packet processing with libpcap bindings. |
| [webrtc](https://github.com/pion/webrtc) | 3222 | 142 | 2018-05-18 | 1 day ago | A pure Go implementation of the WebRTC API. |
| [kcp-go](https://github.com/xtaci/kcp-go) | 2494 | 142 | 2015-06-16 | 1 day ago | KCP - Fast and Reliable ARQ Protocol. |
| [gobgp](https://github.com/osrg/gobgp) | 1827 | 122 | 2014-09-14 | 20 hours ago | BGP implemented in the Go Programming Language. |
| [gnet](https://github.com/panjf2000/gnet) | 1761 | 65 | 2019-02-24 | 1 day ago | `gnet` is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go. |
| [ssh](https://github.com/gliderlabs/ssh) | 1541 | 43 | 2016-10-03 | 1 month ago | Higher-level API for building SSH servers (wraps crypto/ssh). |
| [fortio](https://github.com/fortio/fortio) | 1183 | 35 | 2017-10-10 | 4 days ago | Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC. |
| [water](https://github.com/songgao/water) | 978 | 40 | 2013-03-25 | 4 hours ago | Simple TUN/TAP library. |
| [go-getter](https://github.com/hashicorp/go-getter) | 865 | 185 | 2015-10-12 | 9 hours ago | Go library for downloading files or directories from various sources using a URL. |
| [sftp](https://github.com/pkg/sftp) | 841 | 50 | 2013-11-05 | 1 day ago | Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt. |
| [gev](https://github.com/Allenxuxu/gev) | 791 | 25 | 2019-09-01 | 2 months ago | gev is a lightweight, fast non-blocking TCP network library based on Reactor mode. |
| [nff-go](https://github.com/intel-go/nff-go) | 774 | 74 | 2017-03-29 | 4 days ago | Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF). |
| [grab](https://github.com/cavaliercoder/grab) | 646 | 17 | 2016-01-05 | 4 months ago | Go package for managing file downloads. |
| [mdns](https://github.com/hashicorp/mdns) | 624 | 190 | 2014-01-29 | 4 weeks ago | Simple mDNS (Multicast DNS) client/server library in Golang. |
| [ftp](https://github.com/jlaffaye/ftp) | 622 | 25 | 2011-05-06 | 3 days ago | Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959). |
| [lhttp](https://github.com/fanux/lhttp) | 549 | 56 | 2015-12-29 | 1 year ago | Powerful websocket framework, build your IM server more easily. |
| [gosnmp](https://github.com/soniah/gosnmp) | 496 | 43 | 2012-08-27 | 2 weeks ago | Native Go library for performing SNMP actions. |
| [cidranger](https://github.com/yl2chen/cidranger) | 454 | 11 | 2017-08-21 | 3 days ago | Fast IP to CIDR lookup for Go. |
| [gotcp](https://github.com/gansidui/gotcp) | 451 | 43 | 2014-04-13 | 2 years ago | Go package for quickly writing tcp applications. |
| [peerdiscovery](https://github.com/schollz/peerdiscovery) | 405 | 17 | 2018-04-22 | 1 month ago | Pure Go library for cross-platform local peer discovery using UDP multicast. |
| [gopcap](https://github.com/akrennmair/gopcap) | 377 | 22 | 2009-11-19 | 2 months ago | Go wrapper for libpcap. |
| [raw](https://github.com/mdlayher/raw) | 345 | 12 | 2015-07-06 | 1 week ago | Package raw enables reading and writing data at the device driver level for a network interface. |
| [go-stun](https://github.com/ccding/go-stun) | 344 | 12 | 2013-08-17 | 1 year ago | Go implementation of the STUN client (RFC 3489 and RFC 5389). |
| [stun](https://github.com/gortc/stun) | 342 | 15 | 2016-04-24 | 2 months ago | Go implementation of RFC 5389 STUN protocol. |
| [tcp_server](https://github.com/firstrow/tcp_server) | 322 | 17 | 2014-10-13 | 10 months ago | Go library for building tcp servers faster. |
| [winrm](https://github.com/masterzen/winrm) | 242 | 14 | 2013-12-30 | 4 months ago | Go WinRM client to remotely execute commands on Windows machines. |
| [buffstreams](https://github.com/StabbyCutyou/buffstreams) | 234 | 14 | 2015-06-29 | 2 years ago | Streaming protocolbuffer data over TCP made easy. |
| [arp](https://github.com/mdlayher/arp) | 216 | 10 | 2015-07-06 | 3 months ago | Package arp implements the ARP protocol, as described in RFC 826. |
| [ethernet](https://github.com/mdlayher/ethernet) | 196 | 13 | 2015-07-03 | 9 months ago | Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. |
| [utp](https://github.com/anacrolix/utp) | 152 | 16 | 2015-03-20 | 2 years ago | Go uTP micro transport protocol implementation. |
| [jazigo](https://github.com/udhos/jazigo) | 144 | 12 | 2016-06-07 | 6 months ago | Jazigo is a tool written in Go for retrieving configuration for multiple network devices. |
| [canopus](https://github.com/zubairhamed/canopus) | 138 | 14 | 2015-02-24 | 2 years ago | CoAP Client/Server implementation (RFC 7252). |
| [gmqtt](https://github.com/DrmagicE/gmqtt) | 133 | 9 | 2018-09-16 | 2 weeks ago | Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1. |
| [sslb](https://github.com/eduardonunesp/sslb) | 124 | 8 | 2015-10-18 | 5 months ago | It's a Super Simples Load Balancer, just a little project to achieve some kind of performance. |
| [gnxi](https://github.com/google/gnxi) | 122 | 22 | 2017-09-26 | 5 months ago | A collection of tools for Network Management that use the gNMI and gNOI protocols. |
| [gaio](https://github.com/xtaci/gaio) | 111 | 6 | 2019-12-20 | 1 day ago | High performance async-io networking for Golang in proactor mode. |
| [xtcp](https://github.com/xfxdev/xtcp) | 97 | 13 | 2016-03-31 | 2 weeks ago | TCP Server Framework with simultaneous full duplex communication,graceful shutdown,custom protocol. |
| [dhcp6](https://github.com/mdlayher/dhcp6) | 65 | 4 | 2015-05-22 | 1 year ago | Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. |
| [ether](https://github.com/songgao/ether) | 65 | 3 | 2014-05-21 | 4 years ago | Cross-platform Go package for sending and receiving ethernet frames. |
| [linkio](https://github.com/ian-kent/linkio) | 47 | 2 | 2014-12-24 | 2 years ago | Network link speed simulation for Reader/Writer interfaces. |
| [packet](https://github.com/aerogo/packet) | 44 | 3 | 2017-10-29 | 3 months ago | Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed. |
| [portproxy](https://github.com/aybabtme/portproxy) | 43 | 0 | 2014-12-13 | 5 years ago | Simple TCP proxy which adds CORS support to API's which don't support it. |
| [graval](https://github.com/koofr/graval) | 25 | 8 | 2014-04-22 | 1 year ago | Experimental FTP server framework. |
| [publicip](https://github.com/polera/publicip) | 18 | 1 | 2016-12-28 | 3 years ago | Package publicip returns your public facing IPv4 address (internet egress). |
| [golibwireshark](https://github.com/sunwxg/golibwireshark) | 17 | 2 | 2015-11-16 | 2 years ago | Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data. |
| [go-powerdns](https://github.com/joeig/go-powerdns) | 12 | 3 | 2018-06-21 | 1 week ago | PowerDNS API bindings for Golang. |
| [llb](https://github.com/kirillDanshin/llb) | 10 | 1 | 2016-02-21 | 4 years ago | It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response. |
| [goshark](https://github.com/sunwxg/goshark) | 9 | 1 | 2015-11-01 | 2 years ago | Package goshark use tshark to decode IP packet and create data struct to analyse packet. |
| [tspool](https://github.com/two/tspool) | 6 | 0 | 2018-10-27 | 1 year ago | A TCP Library use worker pool to improve performance and protect your server. |
| [gosocsvr](https://github.com/Rakeki/gosocsvr) | 5 | 2 | 2019-11-12 | 3 months ago | Socket server made simple. |
| [httpproxy](https://github.com/wzshiming/httpproxy) | 2 | 1 | 2018-07-18 | 1 month ago | HTTP proxy handler and dialer. |

### HTTP Clients
        
*Libraries for making HTTP requests.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [resty](https://github.com/go-resty/resty) | 2620 | 69 | 2015-08-28 | 4 days ago | Simple HTTP and REST client for Go inspired by Ruby rest-client. |
| [grequests](https://github.com/levigross/grequests) | 1530 | 33 | 2015-06-11 | 2 months ago | A Go "clone" of the great and famous Requests library. |
| [heimdall](https://github.com/gojek/heimdall) | 1259 | 48 | 2018-01-19 | 3 weeks ago | An enchanced http client with retry and hystrix capabilities. |
| [sling](https://github.com/dghubble/sling) | 1105 | 34 | 2015-04-02 | 3 months ago | Sling is a Go HTTP client library for creating and sending API requests. |
| [gentleman](https://github.com/h2non/gentleman) | 736 | 20 | 2016-02-21 | 4 weeks ago | Full-featured plugin-driven HTTP client library. |
| [pester](https://github.com/sethgrid/pester) | 354 | 7 | 2015-05-20 | 1 week ago | Go HTTP client calls with retries, backoff, and concurrency. |
| [sreq](https://github.com/winterssy/sreq) | 50 | 0 | 2019-12-04 | 1 month ago | A simple, user-friendly and concurrent safe HTTP request library for Go. |
| [rq](https://github.com/ddo/rq) | 34 | 2 | 2017-12-26 | 6 months ago | A nicer interface for golang stdlib HTTP client. |
| [httpretry](https://github.com/ybbus/httpretry) | 4 | 2 | 2020-02-05 | 1 month ago | Enriches the default go HTTP client with retry functionality. |

## OpenGL
        
*Libraries for using OpenGL in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [glfw](https://github.com/go-gl/glfw) | 872 | 41 | 2013-05-19 | 4 days ago | Go bindings for GLFW 3. |
| [gl](https://github.com/go-gl/gl) | 695 | 39 | 2015-02-22 | 1 year ago | Go bindings for OpenGL (generated via glow). |
| [mathgl](https://github.com/go-gl/mathgl) | 322 | 25 | 2013-02-13 | 5 months ago | Pure Go math package specialized for 3D math, with inspiration from GLM. |
| [gl](https://github.com/goxjs/gl) | 132 | 13 | 2015-05-18 | 1 month ago | Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android). |
| [glfw](https://github.com/goxjs/glfw) | 61 | 6 | 2014-12-27 | 3 weeks ago | Go cross-platform glfw library for creating an OpenGL context and receiving events. |

## ORM
        
*Libraries that implement Object-Relational Mapping or datamapping techniques.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gorm](https://github.com/jinzhu/gorm) | 17504 | 457 | 2013-10-25 | 4 days ago | The fantastic ORM library for Golang, aims to be developer friendly. |
| [xorm](https://github.com/go-xorm/xorm) | 5842 | 263 | 2013-05-09 | 4 months ago | Simple and powerful ORM for Go. |
| [pg](https://github.com/go-pg/pg) | 3605 | 86 | 2013-04-24 | 2 weeks ago | PostgreSQL ORM with focus on PostgreSQL specific features and performance. |
| [gorp](https://github.com/go-gorp/gorp) | 3307 | 115 | 2012-01-04 | 1 month ago | Go Relational Persistence, ORM-ish library for Go. |
| [sqlboiler](https://github.com/volatiletech/sqlboiler) | 2706 | 74 | 2016-02-21 | 2 days ago | ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema. |
| [db](https://github.com/upper/db) | 2100 | 61 | 2013-10-23 | 2 weeks ago | Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers. |
| [reform](https://github.com/go-reform/reform) | 847 | 30 | 2016-02-25 | 1 week ago | Better ORM for Go, based on non-empty interfaces and code generation. |
| [pop](https://github.com/gobuffalo/pop) | 834 | 25 | 2018-02-07 | 1 week ago | Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite. |
| [qbs](https://github.com/coocood/qbs) | 549 | 45 | 2013-02-02 | 2 years ago | Stands for Query By Struct. A Go ORM. |
| [go-queryset](https://github.com/jirfag/go-queryset) | 494 | 19 | 2017-09-03 | 3 months ago | 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM. |
| [gormt](https://github.com/xxjwxc/gormt) | 387 | 9 | 2019-05-05 | 1 week ago | Mysql database to golang gorm struct. |
| [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) | 360 | 14 | 2017-12-27 | 3 weeks ago | A flexible and powerful SQL string builder library plus a zero-config ORM. |
| [zoom](https://github.com/albrow/zoom) | 250 | 16 | 2013-07-17 | 1 year ago | Blazing-fast datastore and querying engine built on Redis. |
| [grimoire](https://github.com/Fs02/grimoire) | 125 | 5 | 2018-03-05 | 3 months ago | Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3). |
| [go-store](https://github.com/gosuri/go-store) | 99 | 9 | 2015-03-22 | 3 years ago | Simple and fast Redis backed key-value store library for Go. |
| [marlow](https://github.com/dadleyy/marlow) | 76 | 5 | 2017-09-15 | 2 months ago | Generated ORM from project structs for compile time safety assurances. |
| [rel](https://github.com/Fs02/rel) | 59 | 4 | 2019-10-06 | 2 days ago | Golang SQL Repository Layer for Clean (Onion) Architecture. |
| [go-firestorm](https://github.com/jschoedt/go-firestorm) | 12 | 1 | 2018-12-04 | 4 months ago | A simple ORM for Google/Firebase Cloud Firestore. |
| [lore](https://github.com/abrahambotros/lore) | 5 | 1 | 2017-04-29 | 2 years ago | Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go. |

## Package Management
        
*Unofficial libraries for package and dependency management.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [dep](https://github.com/golang/dep) | 13120 | 291 | 2016-10-07 | 1 month ago | Go dependency tool. |
| [glide](https://github.com/Masterminds/glide) | 7962 | 201 | 2014-07-09 | 2 months ago | Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip. |
| [godep](https://github.com/tools/godep) | 5647 | 152 | 2013-05-01 | 1 year ago | dependency tool for go, godep helps build packages reproducibly by fixing their dependencies. |
| [govendor](https://github.com/kardianos/govendor) | 5013 | 103 | 2015-04-12 | 2 weeks ago | Go Package Manager. Go vendor tool that works with the standard vendor file. |
| [gopm](https://github.com/gpmgo/gopm) | 2470 | 83 | 2013-05-15 | 7 months ago | Go Package Manager. |
| [gom](https://github.com/mattn/gom) | 1380 | 36 | 2013-09-11 | 7 months ago | Go Manager - bundle for go. |
| [gpm](https://github.com/pote/gpm) | 1203 | 31 | 2013-09-05 | 2 years ago | Barebones dependency manager for Go. |
| [goop](https://github.com/petejkim/goop) | 780 | 37 | 2014-06-18 | 4 years ago | Simple dependency manager for Go (golang), inspired by Bundler. |
| [nut](https://github.com/jingweno/nut) | 244 | 8 | 2015-01-23 | 4 years ago | Vendor Go dependencies. |
| [johnny-deps](https://github.com/VividCortex/johnny-deps) | 214 | 21 | 2013-07-19 | 5 years ago | Minimal dependency version using Git. |
| [gigo](https://github.com/LyricalSecurity/gigo) | 198 | 5 | 2015-05-01 | 1 year ago | PIP-like dependency tool for golang, with support for private repositories and hashes. |
| [VenGO](https://github.com/DamnWidget/VenGO) | 117 | 8 | 2014-10-17 | 3 years ago | create and manage exportable isolated go virtual environments. |
| [mvn-golang](https://github.com/raydac/mvn-golang) | 98 | 9 | 2016-03-24 | 2 weeks ago | plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure. |
| [gop](https://github.com/lunny/gop) | 49 | 7 | 2017-02-18 | 1 year ago | Build and manage your Go applications out of GOPATH. |

## Performance
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [jaeger](https://github.com/jaegertracing/jaeger) | 10419 | 316 | 2016-04-15 | 1 hour ago | A distributed tracing system. |
| [profile](https://github.com/pkg/profile) | 1203 | 37 | 2014-10-22 | 3 months ago | Simple profiling support package for Go. |
| [tracer](https://github.com/kamilsk/tracer) | 23 | 1 | 2019-06-22 | 2 months ago | Simple, lightweight tracing. |

## Query Language
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [graphql](https://github.com/graphql-go/graphql) | 6118 | 145 | 2015-07-19 | 1 week ago | Implementation of GraphQL for Go. |
| [graphql-go](https://github.com/graph-gophers/graphql-go) | 3216 | 88 | 2016-10-18 | 6 days ago | GraphQL server with a focus on ease of use. |
| [gojsonq](https://github.com/thedevsaddam/gojsonq) | 1322 | 28 | 2018-05-19 | 5 days ago | A simple Go package to Query over JSON Data. |
| [jsonql](https://github.com/elgs/jsonql) | 219 | 8 | 2015-12-29 | 5 months ago | JSON query expression library in Golang. |
| [rql](https://github.com/a8m/rql) | 136 | 7 | 2018-06-05 | 3 weeks ago | Resource Query Language for REST API. |
| [graphql](https://github.com/tmc/graphql) | 52 | 10 | 2015-04-18 | 2 years ago | graphql parser + utilities. |
| [jsonslice](https://github.com/bhmj/jsonslice) | 39 | 0 | 2018-05-02 | 2 months ago | Jsonpath queries with advanced filters. |
| [straf](https://github.com/SonicRoshan/straf) | 10 | 1 | 2019-08-16 | 5 months ago | Easily Convert Golang structs to GraphQL objects. |

## Resource Embedding
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [packr](https://github.com/gobuffalo/packr) | 2665 | 45 | 2017-03-15 | 3 days ago | The simple and easy way to embed static files into Go binaries. |
| [statik](https://github.com/rakyll/statik) | 2559 | 40 | 2014-02-04 | 2 weeks ago | Embeds static files into a Go executable. |
| [go.rice](https://github.com/GeertJohan/go.rice) | 1868 | 50 | 2013-10-23 | 3 months ago | go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy. |
| [vfsgen](https://github.com/shurcooL/vfsgen) | 813 | 21 | 2015-05-18 | 1 month ago | Generates a vfsdata.go file that statically implements the given virtual filesystem. |
| [esc](https://github.com/mjibson/esc) | 534 | 12 | 2014-01-26 | 4 months ago | Embeds files into Go programs and provides http.FileSystem interfaces to them. |
| [fileb0x](https://github.com/UnnoTed/fileb0x) | 502 | 16 | 2016-01-23 | 4 months ago | Simple tool to embed files in go with focus on "customization" and ease to use. |
| [go-resources](https://github.com/omeid/go-resources) | 165 | 7 | 2015-02-21 | 2 weeks ago | Unfancy resources embedding with Go. |
| [statics](https://github.com/go-playground/statics) | 54 | 3 | 2015-10-07 | 3 years ago | Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks. |
| [templify](https://github.com/wlbr/templify) | 25 | 2 | 2016-05-22 | 6 months ago | Embed external template files into Go code to create single file binaries. |
| [go-embed](https://github.com/pyros2097/go-embed) | 21 | 1 | 2015-12-06 | 3 weeks ago | Generates go code to embed resource files into your library or executable. |

## Science and Data Analysis
        
*Libraries for scientific computing and data analyzing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gonum](https://github.com/gonum/gonum) | 3624 | 99 | 2017-03-25 | 3 hours ago | Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more. |
| [stats](https://github.com/montanaflynn/stats) | 1641 | 50 | 2014-12-16 | 3 weeks ago | Statistics package with common functions missing from the Golang standard library. |
| [plot](https://github.com/gonum/plot) | 1493 | 57 | 2013-07-23 | 4 days ago | gonum/plot provides an API for building and drawing plots in Go. |
| [gosl](https://github.com/cpmech/gosl) | 1401 | 67 | 2015-02-09 | 1 month ago | Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more. |
| [streamtools](https://github.com/nytlabs/streamtools) | 1317 | 72 | 2013-07-05 | 4 years ago | general purpose, graphical tool for dealing with streams of data. |
| [go-dsp](https://github.com/mjibson/go-dsp) | 667 | 29 | 2011-11-02 | 1 year ago | Digital Signal Processing for Go. |
| [chart](https://github.com/vdobler/chart) | 626 | 44 | 2011-06-27 | 7 months ago | Simple Chart Plotting library for Go. Supports many graphs types. |
| [goraph](https://github.com/gyuho/goraph) | 619 | 38 | 2014-02-27 | 2 years ago | Pure Go graph theory library(data structure, algorith visualization). |
| [graph](https://github.com/yourbasic/graph) | 306 | 17 | 2017-04-27 | 4 months ago | Library of basic graph algorithms. |
| [ewma](https://github.com/VividCortex/ewma) | 291 | 25 | 2013-07-05 | 4 months ago | Exponentially-weighted moving averages. |
| [orb](https://github.com/paulmach/orb) | 268 | 19 | 2016-03-28 | 1 month ago | 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support. |
| [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) | 165 | 15 | 2018-10-01 | 1 week ago | Dataframes for machine-learning and statistics (similar to pandas). |
| [gohistogram](https://github.com/VividCortex/gohistogram) | 140 | 17 | 2013-07-02 | 2 years ago | Approximate histograms for data streams. |
| [TextRank](https://github.com/DavidBelicza/TextRank) | 96 | 6 | 2018-01-09 | 1 year ago | TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support. |
| [sparse](https://github.com/james-bowman/sparse) | 82 | 5 | 2017-05-16 | 1 month ago | Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries. |
| [pagerank](https://github.com/alixaxel/pagerank) | 55 | 6 | 2015-08-06 | 2 months ago | Weighted PageRank algorithm implemented in Go. |
| [geom](https://github.com/skelterjohn/geom) | 44 | 4 | 2011-06-07 | 2 years ago | 2D geometry for golang. |
| [evaler](https://github.com/soniah/evaler) | 39 | 4 | 2012-09-04 | 1 year ago | Simple floating point arithmetic expression evaluator. |
| [goent](https://github.com/kzahedi/goent) | 17 | 1 | 2017-08-08 | 11 months ago | GO Implementation of Entropy Measures. |
| [triangolatte](https://github.com/tchayen/triangolatte) | 12 | 1 | 2018-07-18 | 20 hours ago | 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs. |
| [ode](https://github.com/ChristopherRabotin/ode) | 11 | 3 | 2016-11-11 | 3 years ago | Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions. |
| [piecewiselinear](https://github.com/sgreben/piecewiselinear) | 11 | 2 | 2018-10-21 | 2 months ago | Tiny linear interpolation library. |
| [GoStats](https://github.com/OGFris/GoStats) | 10 | 1 | 2018-07-22 | 1 year ago | GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions. |
| [PiHex](https://github.com/claygod/PiHex) | 8 | 2 | 2016-07-22 | 5 months ago | Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi. |
| [assocentity](https://github.com/ndabAP/assocentity) | 6 | 1 | 2018-12-21 | 1 month ago | Package assocentity returns the average distance from words to a given entity. |
| [go-gt](https://github.com/ThePaw/go-gt) | 5 | 0 | 2015-09-14 | 4 years ago | Graph theory algorithms written in "Go" language. |
| [rootfinding](https://github.com/khezen/rootfinding) | 3 | 2 | 2018-10-30 | 1 month ago | root-finding algorithms library for finding roots of quadratic functions. |
| [bradleyterry](https://github.com/seanhagen/bradleyterry) | 2 | 1 | 2019-04-30 | 10 months ago | Provides a Bradley-Terry Model for pairwise comparisons. |

## Security
        
*Libraries that are used to help make your application more secure.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [lego](https://github.com/go-acme/lego) | 3901 | 103 | 2015-06-08 | 18 hours ago | Pure Go ACME client library and CLI tool (for use with Let's Encrypt). |
| [cameradar](https://github.com/Ullaakut/cameradar) | 2153 | 106 | 2016-05-20 | 1 month ago | Tool and library to remotely hack RTSP streams from surveillance cameras. |
| [acmetool](https://github.com/hlandau/acmetool) | 1747 | 68 | 2015-11-15 | 1 month ago | ACME (Let's Encrypt) client tool with automatic renewal. |
| [memguard](https://github.com/awnumar/memguard) | 1726 | 45 | 2017-04-22 | 4 days ago | A pure Go library for handling sensitive values in memory. |
| [secure](https://github.com/unrolled/secure) | 1429 | 37 | 2014-05-20 | 3 months ago | HTTP middleware for Go that facilitates some quick security wins. |
| [acra](https://github.com/cossacklabs/acra) | 546 | 36 | 2016-11-14 | 6 hours ago | Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system. |
| [nacl](https://github.com/kevinburke/nacl) | 476 | 12 | 2017-07-20 | 6 months ago | Go implementation of the NaCL set of API's. |
| [badactor](https://github.com/jaredfolkins/badactor) | 265 | 9 | 2014-12-12 | 6 months ago | In-memory, application-driven jailer built in the spirit of fail2ban. |
| [passlib](https://github.com/hlandau/passlib) | 233 | 10 | 2014-12-21 | 11 months ago | Futureproof password hashing library. |
| [ssh-vault](https://github.com/ssh-vault/ssh-vault) | 231 | 10 | 2016-09-29 | 1 month ago | encrypt/decrypt using ssh keys. |
| [optimus-go](https://github.com/pjebs/optimus-go) | 193 | 4 | 2015-05-05 | 1 month ago | ID hashing and Obfuscation using Knuth's Algorithm. |
| [simple-scrypt](https://github.com/elithrar/simple-scrypt) | 163 | 7 | 2015-04-14 | 8 months ago | Scrypt package with a simple, obvious API and automatic cost calibration built-in. |
| [go-yara](https://github.com/hillu/go-yara) | 153 | 19 | 2015-01-25 | 8 hours ago | Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)". |
| [argon2pw](https://github.com/raja/argon2pw) | 81 | 3 | 2018-03-13 | 1 year ago | Argon2 password hash generation with constant-time password comparison. |
| [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) | 35 | 1 | 2017-10-19 | 1 year ago | A probably paranoid package for securely hashing and encrypting passwords. |
| [certificates](https://github.com/mvmaasakkers/certificates) | 13 | 1 | 2019-03-04 | 1 month ago | An opinionated tool for generating tls certificates. |
| [goArgonPass](https://github.com/dwin/goArgonPass) | 11 | 2 | 2018-05-30 | 1 month ago | Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations. |
| [sslmgr](https://github.com/adrianosela/sslmgr) | 8 | 2 | 2019-04-02 | 7 months ago | SSL certificates made easy with a high level wrapper around acme/autocert. |
| [go-generate-password](https://github.com/m1/go-generate-password) | 7 | 1 | 2019-11-14 | 4 months ago | Password generator that can be used on the cli or as a library. |

## Serialization
        
*Libraries and tools for binary serialization.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go](https://github.com/json-iterator/go) | 7289 | 218 | 2016-11-30 | 1 week ago | High-performance 100% compatible drop-in replacement of "encoding/json". |
| [protobuf](https://github.com/golang/protobuf) | 6311 | 211 | 2014-11-23 | 2 days ago | Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers. |
| [protobuf](https://github.com/gogo/protobuf) | 3559 | 97 | 2014-12-03 | 1 month ago | Protocol Buffers for Go with Gadgets. |
| [mapstructure](https://github.com/mitchellh/mapstructure) | 3127 | 53 | 2013-05-20 | 3 months ago | Go library for decoding generic map values into native Go structures. |
| [go](https://github.com/ugorji/go) | 1374 | 54 | 2013-05-30 | 1 month ago | High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support. |
| [colfer](https://github.com/pascaldekloe/colfer) | 517 | 35 | 2015-09-05 | 4 months ago | Code generation for the Colfer binary format. |
| [csvutil](https://github.com/jszwec/csvutil) | 353 | 9 | 2017-10-30 | 2 weeks ago | High Performance, idiomatic CSV record encoding and decoding to native Go structures. |
| [go-capnproto](https://github.com/glycerine/go-capnproto) | 277 | 11 | 2013-11-07 | 1 month ago | Cap'n Proto library and parser for go. |
| [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) | 135 | 9 | 2012-12-23 | 1 year ago | GoLang library for working with PHP session format and PHP Serialize/Unserialize functions. |
| [cbor](https://github.com/fxamacker/cbor) | 116 | 5 | 2019-05-15 | 2 days ago | Small, safe, and easy CBOR encoding and decoding library. |
| [structomap](https://github.com/danhper/structomap) | 110 | 7 | 2015-05-13 | 9 months ago | Library to easily and dynamically generate maps from static structures. |
| [bambam](https://github.com/glycerine/bambam) | 60 | 4 | 2014-09-17 | 3 years ago | generator for Cap'n Proto schemas from go. |
| [asn1](https://github.com/Logicalis/asn1) | 44 | 8 | 2016-02-29 | 1 year ago | Asn.1 BER and DER encoding library for golang. |
| [binstruct](https://github.com/ghostiam/binstruct) | 17 | 1 | 2018-10-23 | 6 months ago | Golang binary decoder for mapping data into the structure. |
| [fwencoder](https://github.com/o1egl/fwencoder) | 10 | 1 | 2017-12-25 | 1 month ago | Fixed width file parser (encoding and decoding library) for Go. |
| [bel](https://github.com/32leaves/bel) | 8 | 1 | 2019-02-20 | 11 months ago | Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC. |
| [bel](https://github.com/csweichel/bel) | 8 | 1 | 2019-02-20 | 11 months ago | Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC. |
| [pletter](https://github.com/vimeda/pletter) | 7 | 0 | 2019-07-09 | 3 days ago | A standard way to wrap a proto message for message brokers. |
| [elastic](https://github.com/epiclabs-io/elastic) | 5 | 0 | 2020-02-25 | 2 weeks ago | Convert slices, maps or any other unknown value across different types at run-time, no matter what. |
| [fixedwidth](https://github.com/huydang284/fixedwidth) | 4 | 1 | 2019-08-11 | 2 months ago | Fixed-width text formatting (UTF-8 supported). |

## Server Applications
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [etcd](https://github.com/etcd-io/etcd) | 30031 | 1339 | 2013-07-06 | 1 hour ago | Highly-available key value store for shared configuration and service discovery. |
| [caddy](https://github.com/caddyserver/caddy) | 26735 | 745 | 2015-01-13 | 13 hours ago | Caddy is an alternative, HTTP/2 web server that's easy to configure and use. |
| [minio](https://github.com/minio/minio) | 20724 | 497 | 2015-01-14 | 2 hours ago | Minio is a distributed object storage server. |
| [roadrunner](https://github.com/spiral/roadrunner) | 4092 | 153 | 2017-12-26 | 2 days ago | High-performance PHP application server, load-balancer and process manager. |
| [devd](https://github.com/cortesi/devd) | 2955 | 69 | 2015-09-27 | 4 weeks ago | Local webserver for developers. |
| [algernon](https://github.com/xyproto/algernon) | 1674 | 50 | 2015-03-10 | 2 weeks ago | HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber. |
| [sftpgo](https://github.com/drakkan/sftpgo) | 1488 | 35 | 2019-07-20 | 18 hours ago | Full featured and highly configurable SFTP server software. |
| [flipt](https://github.com/markphelps/flipt) | 1173 | 14 | 2016-11-05 | 6 hours ago | A self contained feature flag solution written in Go and Vue. |
| [trickster](https://github.com/tricksterproxy/trickster) | 1058 | 36 | 2018-03-29 | 23 hours ago | HTTP reverse proxy cache and time series accelerator. |
| [fider](https://github.com/getfider/fider) | 1056 | 23 | 2017-01-17 | 2 days ago | Fider is an open platform to collect and organize customer feedback. |
| [trickster](https://github.com/Comcast/trickster) | 1053 | 35 | 2018-03-29 | 4 days ago | HTTP reverse proxy cache and time series accelerator. |
| [flagr](https://github.com/checkr/flagr) | 1024 | 72 | 2017-10-03 | 4 days ago | Flagr is an open-source feature flagging and A/B testing service. |
| [discovery](https://github.com/bilibili/discovery) | 869 | 47 | 2018-04-20 | 5 days ago | A registry for resilient mid-tier load balancing and failover. |
| [jackal](https://github.com/ortuman/jackal) | 797 | 34 | 2017-11-13 | 1 day ago | An XMPP server written in Go. |
| [dudeldu](https://github.com/krotik/dudeldu) | 109 | 3 | 2016-09-07 | 5 months ago | A simple SHOUTcast server. |
| [psql-streamer](https://github.com/blind-oracle/psql-streamer) | 15 | 4 | 2019-04-28 | 6 days ago | Stream database events from PostgreSQL to Kafka. |
| [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) | 9 | 1 | 2018-10-23 | 10 months ago | Nginx log parser and exporter to Prometheus. |
| [riemann-relay](https://github.com/blind-oracle/riemann-relay) | 0 | 1 | 2019-04-23 | 4 months ago | Relay to load-balance Riemann events and/or convert them to Carbon. |

## Stream Processing
        
*Libraries and tools for stream processing and reactive programming.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-streams](https://github.com/reugn/go-streams) | 287 | 11 | 2019-04-30 | 1 month ago | Go stream processing library. |

## Template Engines
        
*Libraries and tools for templating and lexing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gofpdf](https://github.com/jung-kurt/gofpdf) | 3491 | 92 | 2015-03-13 | 3 months ago | PDF document generator with high level support for text, drawing and images. |
| [quicktemplate](https://github.com/valyala/quicktemplate) | 1654 | 58 | 2016-03-06 | 2 weeks ago | Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it. |
| [pongo2](https://github.com/flosch/pongo2) | 1633 | 56 | 2013-08-23 | 2 weeks ago | Django-like template-engine for Go. |
| [hero](https://github.com/shiyanhui/hero) | 1292 | 42 | 2017-01-15 | 2 months ago | Hero is a handy, fast and powerful go template engine. |
| [mustache](https://github.com/hoisie/mustache) | 980 | 36 | 2009-12-30 | 3 weeks ago | Go implementation of the Mustache template language. |
| [amber](https://github.com/eknkc/amber) | 840 | 21 | 2012-10-31 | 1 year ago | Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade. |
| [ace](https://github.com/yosssi/ace) | 785 | 24 | 2014-07-13 | 1 year ago | Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold. |
| [gorazor](https://github.com/sipin/gorazor) | 733 | 57 | 2014-05-01 | 4 months ago | Razor view engine for Golang. |
| [jet](https://github.com/CloudyKit/jet) | 634 | 23 | 2016-03-31 | 3 days ago | Jet template engine. |
| [ego](https://github.com/benbjohnson/ego) | 429 | 16 | 2014-02-23 | 2 months ago | Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled. |
| [fasttemplate](https://github.com/valyala/fasttemplate) | 368 | 17 | 2015-08-19 | 1 month ago | Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/). |
| [raymond](https://github.com/aymerick/raymond) | 364 | 11 | 2015-04-22 | 1 year ago | Complete handlebars implementation in Go. |
| [soy](https://github.com/robfig/soy) | 148 | 12 | 2013-12-15 | 1 year ago | Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/). |
| [maroto](https://github.com/johnfercher/maroto) | 122 | 6 | 2019-05-20 | 3 weeks ago | A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple. |
| [goview](https://github.com/foolin/goview) | 109 | 3 | 2019-04-14 | 2 weeks ago | Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application. |
| [liquid](https://github.com/osteele/liquid) | 101 | 4 | 2017-06-26 | 2 months ago | Go implementation of Shopify Liquid templates. |
| [kasia.go](https://github.com/ziutek/kasia.go) | 71 | 2 | 2010-12-07 | 4 years ago | Templating system for HTML and other text documents - go implementation. |
| [velvet](https://github.com/gobuffalo/velvet) | 71 | 5 | 2016-12-29 | 3 years ago | Complete handlebars implementation in Go. |
| [damsel](https://github.com/dskinner/damsel) | 22 | 4 | 2012-05-02 | 3 years ago | Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others. |
| [extemplate](https://github.com/dannyvankooten/extemplate) | 19 | 3 | 2018-08-10 | 1 year ago | Tiny wrapper around html/template to allow for easy file-based template inheritance. |
| [gospin](https://github.com/m1/gospin) | 8 | 1 | 2019-02-22 | 1 week ago | Article spinning and spintax/spinning syntax engine, useful for A/B, testing pieces of text/articles and creating more natural conversations. |

## Testing
        
*Libraries for testing codebases and generating test data.*

### Testing Frameworks
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [testify](https://github.com/stretchr/testify) | 9873 | 153 | 2012-10-16 | 1 day ago | Sacred extension to the standard go testing package. |
| [go-cmp](https://github.com/google/go-cmp) | 1555 | 25 | 2017-07-07 | 2 weeks ago | Package for comparing Go values in tests. |
| [httpexpect](https://github.com/gavv/httpexpect) | 1300 | 32 | 2016-04-29 | 2 months ago | Concise, declarative, and easy to use end-to-end HTTP and REST API testing. |
| [godog](https://github.com/cucumber/godog) | 957 | 101 | 2015-06-10 | 4 hours ago | Cucumber or Behat like BDD framework for Go. |
| [godog](https://github.com/DATA-DOG/godog) | 895 | 30 | 2015-06-10 | 1 month ago | Cucumber or Behat like BDD framework for Go. |
| [baloo](https://github.com/h2non/baloo) | 675 | 10 | 2016-05-29 | 1 year ago | Expressive and versatile end-to-end HTTP API testing made easy. |
| [goblin](https://github.com/franela/goblin) | 663 | 16 | 2013-09-19 | 2 months ago | Mocha like testing framework fo Go. |
| [testfixtures](https://github.com/go-testfixtures/testfixtures) | 453 | 5 | 2016-04-05 | 2 weeks ago | A helper for Rails' like test fixtures to test database applications. |
| [go-vcr](https://github.com/dnaeon/go-vcr) | 394 | 8 | 2015-12-14 | 1 week ago | Record and replay your HTTP interactions for fast, deterministic and accurate tests. |
| [go-mutesting](https://github.com/zimmski/go-mutesting) | 330 | 7 | 2014-12-26 | 4 months ago | Mutation testing for Go source code. |
| [gofight](https://github.com/appleboy/gofight) | 303 | 11 | 2016-03-29 | 2 months ago | API Handler Testing for Golang Router framework. |
| [frisby](https://github.com/verdverm/frisby) | 255 | 8 | 2015-09-15 | 1 week ago | REST API testing framework. |
| [go-carpet](https://github.com/msoap/go-carpet) | 202 | 4 | 2016-02-28 | 3 months ago | Tool for viewing test coverage in terminal. |
| [charlatan](https://github.com/percolate/charlatan) | 192 | 41 | 2017-10-06 | 6 months ago | Tool to generate fake interface implementations for tests. |
| [gotest.tools](https://github.com/gotestyourself/gotest.tools) | 148 | 7 | 2017-08-08 | 1 day ago | A collection of packages to augment the go testing package and support common patterns. |
| [endly](https://github.com/viant/endly) | 133 | 14 | 2017-08-28 | 2 weeks ago | Declarative end to end functional testing. |
| [commander](https://github.com/SimonBaeumer/commander) | 119 | 6 | 2019-02-22 | 22 hours ago | Tool for testing cli applications on windows, linux and osx. |
| [gospec](https://github.com/luontola/gospec) | 114 | 4 | 2009-11-24 | 5 years ago | BDD-style testing framework for the Go programming language. |
| [cupaloy](https://github.com/bradleyjkemp/cupaloy) | 103 | 2 | 2017-08-07 | 3 weeks ago | Simple snapshot testing addon for your test framework. |
| [dbcleaner](https://github.com/khaiql/dbcleaner) | 101 | 2 | 2017-01-17 | 5 days ago | Clean database for testing purpose, inspired by `database_cleaner` in Ruby. |
| [go-testdeep](https://github.com/maxatome/go-testdeep) | 84 | 2 | 2018-05-26 | 2 weeks ago | Extremely flexible golang deep comparison, extends the go testing package. |
| [wstest](https://github.com/posener/wstest) | 73 | 2 | 2017-03-31 | 2 days ago | Websocket client for unit-testing a websocket http.Handler. |
| [gospecify](https://github.com/stesla/gospecify) | 53 | 5 | 2009-11-20 | 8 years ago | This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec. |
| [restit](https://github.com/go-restit/restit) | 50 | 6 | 2014-06-25 | 5 months ago | Go micro framework to help writing RESTful API integration test. |
| [jsonassert](https://github.com/kinbiko/jsonassert) | 34 | 0 | 2018-10-26 | 2 months ago | Package for verifying that your JSON payloads are serialized correctly. |
| [gomatch](https://github.com/jfilipczyk/gomatch) | 32 | 2 | 2019-01-27 | 8 months ago | library created for testing JSON against patterns. |
| [dsunit](https://github.com/viant/dsunit) | 29 | 9 | 2016-06-13 | 1 month ago | Datastore testing for SQL, NoSQL, structured files. |
| [hamcrest](https://github.com/rdrdr/hamcrest) | 27 | 3 | 2010-12-22 | 9 years ago | fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results. |
| [assert](https://github.com/go-playground/assert) | 17 | 1 | 2015-07-20 | 5 months ago | Basic Assertion Library used along side native go testing, with building blocks for custom assertions. |
| [testcase](https://github.com/adamluzsi/testcase) | 17 | 1 | 2019-04-22 | 3 days ago | Idiomatic testing framework for Behavior Driven Development. |
| [embedded-postgres](https://github.com/fergusstrange/embedded-postgres) | 12 | 1 | 2019-11-16 | 3 months ago | Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test. |
| [flute](https://github.com/suzuki-shunsuke/flute) | 11 | 2 | 2019-07-06 | 3 months ago | HTTP client testing framework. |
| [badio](https://github.com/cavaliercoder/badio) | 9 | 1 | 2016-02-11 | 4 years ago | Extensions to Go's `testing/iotest` package. |
| [gocrest](https://github.com/corbym/gocrest) | 9 | 1 | 2017-12-23 | 2 years ago | Composable hamcrest-like matchers for Go assertions. |
| [gosuite](https://github.com/pavlo/gosuite) | 9 | 1 | 2016-10-15 | 3 years ago | Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests. |
| [schema](https://github.com/jgroeneveld/schema) | 9 | 3 | 2015-08-13 | 5 months ago | Quick and easy expression matching for JSON schemas used in requests and responses. |
| [biff](https://github.com/fulldump/biff) | 8 | 1 | 2018-03-28 | 1 week ago | Bifurcation testing framework, BDD compatible. |
| [gogiven](https://github.com/corbym/gogiven) | 8 | 4 | 2017-12-31 | 2 years ago | YATSPEC-like BDD testing framework for Go. |
| [testsql](https://github.com/zhulongcheng/testsql) | 7 | 2 | 2018-09-22 | 5 months ago | Generate test data from SQL files before testing and clear it after finished. |
| [tt](https://github.com/vcaesar/tt) | 6 | 1 | 2018-04-03 | 2 weeks ago | Simple and colorful test tools. |
| [trial](https://github.com/jgroeneveld/trial) | 4 | 1 | 2015-06-18 | 5 months ago | Quick and easy extendable assertions without introducing much boilerplate. |

### Mock
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [chromedp](https://github.com/chromedp/chromedp) | 4419 | 138 | 2017-01-24 | 1 week ago | a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol. |
| [mock](https://github.com/golang/mock) | 3868 | 76 | 2015-06-12 | 2 hours ago | Mocking framework for the Go programming language. |
| [go-fuzz](https://github.com/dvyukov/go-fuzz) | 3338 | 88 | 2015-04-15 | 1 month ago | Randomized testing system. |
| [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) | 2310 | 26 | 2014-02-07 | 1 month ago | Mock SQL driver for testing database interactions. |
| [hoverfly](https://github.com/SpectoLabs/hoverfly) | 1532 | 59 | 2015-11-30 | 3 days ago | HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI. |
| [selenoid](https://github.com/aerokube/selenoid) | 1525 | 90 | 2016-08-22 | 4 days ago | alternative Selenium hub server that launches browsers within containers. |
| [gock](https://github.com/h2non/gock) | 962 | 17 | 2016-03-02 | 1 month ago | Versatile HTTP mocking made easy. |
| [gofuzz](https://github.com/google/gofuzz) | 737 | 23 | 2014-07-31 | 1 month ago | Library for populating go objects with random values. |
| [httpmock](https://github.com/jarcoal/httpmock) | 719 | 8 | 2014-02-24 | 2 weeks ago | Easy mocking of HTTP responses from external resources. |
| [cdp](https://github.com/mafredri/cdp) | 407 | 18 | 2017-03-12 | 3 months ago | Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it. |
| [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) | 401 | 7 | 2014-05-21 | 1 week ago | Tool for generating self-contained mock objects. |
| [minimock](https://github.com/gojuno/minimock) | 294 | 10 | 2016-08-03 | 2 weeks ago | Mock generator for Go interfaces. |
| [ggr](https://github.com/aerokube/ggr) | 243 | 24 | 2016-06-16 | 3 months ago | a lightweight server that routes and proxies Selenium WebDriver requests to multiple Selenium hubs. |
| [go-txdb](https://github.com/DATA-DOG/go-txdb) | 238 | 9 | 2015-07-08 | 2 months ago | Single transaction based database driver mainly for testing purposes. |
| [tavor](https://github.com/zimmski/tavor) | 219 | 12 | 2014-05-18 | 1 year ago | Generic fuzzing and delta-debugging framework. |
| [govcr](https://github.com/seborama/govcr) | 88 | 2 | 2016-07-10 | 5 months ago | HTTP mock for Golang: record and replay HTTP interactions for offline testing. |
| [timex](https://github.com/cabify/timex) | 26 | 72 | 2020-01-02 | 2 months ago | A test-friendly replacement for the native `time` package. |
| [mockhttp](https://github.com/tv42/mockhttp) | 22 | 1 | 2011-06-11 | 5 years ago | Mock object for Go http.ResponseWriter. |
| [rod](https://github.com/ysmood/rod) | 12 | 2 | 2020-01-21 | 3 days ago | A chrome devtools controller that is easy and safe to use. |

### Fail injection
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [failpoint](https://github.com/pingcap/failpoint) | 486 | 76 | 2019-04-02 | 1 week ago | An implementation of [failpoints](http://www.freebsd.org/cgi/man.cgi?query=fail) for Golang. |

## Text Processing
        
*Libraries for parsing and manipulating texts.*

### Specific Formats
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [colly](https://github.com/gocolly/colly) | 10293 | 276 | 2017-09-29 | 1 day ago | Fast and Elegant Scraping Framework for Gophers. |
| [goquery](https://github.com/PuerkitoBio/goquery) | 8498 | 270 | 2012-08-29 | 2 weeks ago | GoQuery brings a syntax and a set of features similar to jQuery to the Go language. |
| [blackfriday](https://github.com/russross/blackfriday) | 4281 | 96 | 2011-05-27 | 3 days ago | Markdown processor in Go. |
| [toml](https://github.com/BurntSushi/toml) | 3094 | 85 | 2013-02-26 | 1 month ago | TOML configuration format (encoder/decoder with reflection). |
| [sh](https://github.com/mvdan/sh) | 2563 | 45 | 2016-01-16 | 2 days ago | Shell parser and formatter. |
| [go-humanize](https://github.com/dustin/go-humanize) | 2120 | 33 | 2012-01-13 | 2 weeks ago | Formatters for time, numbers, and memory size to human readable format. |
| [bluemonday](https://github.com/microcosm-cc/bluemonday) | 1482 | 33 | 2013-11-20 | 3 months ago | HTML Sanitizer. |
| [inject](https://github.com/facebookarchive/inject) | 1208 | 44 | 2013-10-21 | 1 year ago | Package inject provides a reflect based injector. |
| [gofeed](https://github.com/mmcdole/gofeed) | 1206 | 37 | 2016-01-23 | 1 day ago | Parse RSS and Atom feeds in Go. |
| [go-toml](https://github.com/pelletier/go-toml) | 717 | 29 | 2013-02-24 | 2 hours ago | Go library for the TOML format with query support and handy cli tools. |
| [commonregex](https://github.com/mingrammer/commonregex) | 585 | 21 | 2017-03-23 | 4 months ago | A collection of common regular expressions for Go. |
| [slug](https://github.com/gosimple/slug) | 517 | 11 | 2014-03-31 | 2 months ago | URL-friendly slugify with multiple languages support. |
| [dataflowkit](https://github.com/slotix/dataflowkit) | 366 | 17 | 2017-02-09 | 5 months ago | Web scraping Framework to turn websites into structured data. |
| [mxj](https://github.com/clbanning/mxj) | 365 | 24 | 2014-02-03 | 4 weeks ago | Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages. |
| [gographviz](https://github.com/awalterschulze/gographviz) | 348 | 14 | 2015-03-14 | 8 months ago | Parses the Graphviz DOT language. |
| [gotext](https://github.com/leonelquinteros/gotext) | 253 | 4 | 2016-06-19 | 1 week ago | GNU gettext utilities for Go. |
| [go-runewidth](https://github.com/mattn/go-runewidth) | 247 | 11 | 2013-06-21 | 4 days ago | Functions to get fixed width of the character or string. |
| [htmlquery](https://github.com/antchfx/htmlquery) | 211 | 8 | 2017-12-05 | 2 months ago | An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression. |
| [goq](https://github.com/andrewstuart/goq) | 163 | 7 | 2017-02-20 | 9 months ago | Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery). |
| [go-nmea](https://github.com/adrianmo/go-nmea) | 114 | 6 | 2015-07-22 | 3 weeks ago | NMEA parser library for the Go language. |
| [sdp](https://github.com/gortc/sdp) | 87 | 7 | 2016-05-13 | 1 month ago | SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)]. |
| [go-zero-width](https://github.com/trubitsyn/go-zero-width) | 77 | 1 | 2018-06-18 | 1 month ago | Zero-width character detection and removal for Go. |
| [align](https://github.com/Guitarbum722/align) | 63 | 4 | 2017-04-29 | 3 months ago | A general purpose application that aligns text. |
| [goribot](https://github.com/zhshch2002/goribot) | 63 | 5 | 2019-09-08 | 4 months ago | A simple golang spider/scraping framework,build a spider in 3 lines. |
| [podcast](https://github.com/eduncan911/podcast) | 61 | 3 | 2017-02-02 | 1 month ago | iTunes Compliant and RSS 2. |
| [genex](https://github.com/alixaxel/genex) | 58 | 3 | 2015-03-09 | 2 months ago | Count and expand Regular Expressions into all matching Strings. |
| [go-slugify](https://github.com/mozillazg/go-slugify) | 58 | 2 | 2016-07-16 | 3 years ago | Make pretty slug with multiple languages support. |
| [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) | 57 | 8 | 2016-07-05 | 1 day ago | Editorconfig file parser and manipulator for Go. |
| [guesslanguage](https://github.com/endeveit/guesslanguage) | 45 | 1 | 2014-12-16 | 2 years ago | Functions to determine the natural language of a unicode text. |
| [goregen](https://github.com/zach-klippenstein/goregen) | 40 | 2 | 2014-12-27 | 5 months ago | Library for generating random strings from regular expressions. |
| [allot](https://github.com/sbstjn/allot) | 38 | 1 | 2016-10-16 | 9 months ago | Placeholder and wildcard text parsing for CLI tools and bots. |
| [go-vcard](https://github.com/emersion/go-vcard) | 37 | 3 | 2017-03-21 | 2 months ago | Parse and format vCard. |
| [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) | 33 | 3 | 2017-11-15 | 11 hours ago | Fixed-width text formatting (encoder/decoder with reflection). |
| [gonameparts](https://github.com/polera/gonameparts) | 31 | 0 | 2015-05-17 | 7 months ago | Parses human names into individual name parts. |
| [did](https://github.com/ockam-network/did) | 30 | 7 | 2018-11-02 | 6 days ago | DID (Decentralized Identifiers) Parser and Stringer in Go. |
| [slugify](https://github.com/avelino/slugify) | 27 | 2 | 2015-04-13 | 1 year ago | Go slugify application that handles string. |
| [codetree](https://github.com/aerogo/codetree) | 14 | 2 | 2016-11-26 | 4 months ago | Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure. |
| [enca](https://github.com/endeveit/enca) | 8 | 1 | 2014-12-17 | 4 years ago | Minimal cgo bindings for [libenca](http://cihar.com/software/enca/). |
| [syndfeed](https://github.com/zhengchun/syndfeed) | 7 | 1 | 2017-04-07 | 2 years ago | A syndication feed for Atom 1.0 and RSS 2.0. |
| [bbConvert](https://github.com/CalebQ42/bbConvert) | 5 | 1 | 2016-04-15 | 3 years ago | Converts bbCode to HTML that allows you to add support for custom bbCode tags. |
| [doi](https://github.com/hscells/doi) | 4 | 1 | 2017-08-02 | 2 years ago | Document object identifier (doi) parser in Go. |
| [ltsv](https://github.com/Wing924/ltsv) | 4 | 1 | 2019-05-12 | 8 months ago | High performance [LTSV (Labeled Tab Separeted Value)](http://ltsv.org/) reader for Go. |
| [encoding](https://github.com/mickep76/encoding) | 3 | 1 | 2018-04-06 | 4 months ago | Package provides a generic interface to encoders and decodersa. |

### Utility
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [xurls](https://github.com/mvdan/xurls) | 607 | 15 | 2015-01-12 | 1 week ago | Extract urls from text. |
| [gotabulate](https://github.com/bndr/gotabulate) | 232 | 7 | 2014-08-21 | 3 years ago | Easily pretty-print your tabular data with Go. |
| [radix](https://github.com/yourbasic/radix) | 157 | 5 | 2017-06-09 | 2 years ago | fast string sorting algorithm. |
| [parth](https://github.com/codemodus/parth) | 36 | 3 | 2015-04-06 | 1 year ago | URL path segmentation parsing. |
| [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) | 22 | 1 | 2018-09-09 | 1 year ago | A sanitization-based swear filter for Go. |
| [xj2go](https://github.com/stackerzzq/xj2go) | 18 | 2 | 2017-09-19 | 3 weeks ago | Convert xml or json to go struct. |
| [kace](https://github.com/codemodus/kace) | 12 | 1 | 2015-06-04 | 1 year ago | Common case conversions covering common initialisms. |
| [tagify](https://github.com/zoomio/tagify) | 10 | 1 | 2018-03-20 | 1 week ago | Produces a set of tags from given source. |
| [parseargs-go](https://github.com/txgruppi/parseargs-go) | 9 | 1 | 2016-02-24 | 3 years ago | string argument parser that understands quotes and backslashes. |
| [TySug](https://github.com/Dynom/TySug) | 4 | 1 | 2018-06-05 | 2 months ago | Alternative suggestions with respect to keyboard layouts. |
| [textwrap](https://github.com/isbm/textwrap) | 1 | 1 | 2019-07-26 | 7 months ago | Implementation of `textwrap` module from Python. |

## Third-party APIs
        
*Libraries for accessing third party APIs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [aws-sdk-go](https://github.com/aws/aws-sdk-go) | 5652 | 257 | 2014-12-05 | 2 days ago | The official AWS SDK for the Go programming language. |
| [go-github](https://github.com/google/go-github) | 5610 | 193 | 2013-05-24 | 8 hours ago | Go library for accessing the GitHub REST API v3. |
| [google-api-go-client](https://github.com/googleapis/google-api-go-client) | 2163 | 134 | 2014-11-24 | 17 hours ago | Auto-generated Google APIs for Go. |
| [google-cloud-go](https://github.com/googleapis/google-cloud-go) | 2061 | 218 | 2014-05-09 | 57 minutes ago | Google Cloud APIs Go Client Library. |
| [discordgo](https://github.com/bwmarrin/discordgo) | 1156 | 51 | 2015-11-01 | 4 days ago | Go bindings for the Discord Chat API. |
| [stripe-go](https://github.com/stripe/stripe-go) | 1096 | 39 | 2014-06-05 | 2 days ago | Go client for the Stripe API. |
| [anaconda](https://github.com/ChimeraCoder/anaconda) | 1020 | 21 | 2013-03-04 | 2 days ago | Go client library for the Twitter 1.1 API. |
| [go-twitter](https://github.com/dghubble/go-twitter) | 922 | 30 | 2015-04-11 | 1 month ago | Go client library for the Twitter v1.1 APIs. |
| [minio-go](https://github.com/minio/minio-go) | 890 | 38 | 2015-05-02 | 8 hours ago | Minio Go Library for Amazon S3 compatible cloud storage. |
| [facebook](https://github.com/huandu/facebook) | 846 | 83 | 2012-07-28 | 1 week ago | Go Library that supports the Facebook Graph API. |
| [githubv4](https://github.com/shurcooL/githubv4) | 609 | 19 | 2017-05-27 | 3 months ago | Go library for accessing the GitHub GraphQL API v4. |
| [webhooks](https://github.com/go-playground/webhooks) | 466 | 15 | 2015-10-25 | 14 hours ago | Webhook receiver for GitHub and Bitbucket. |
| [paypal](https://github.com/plutov/paypal) | 343 | 16 | 2015-10-14 | 1 month ago | Wrapper for PayPal payment API. |
| [geo-golang](https://github.com/codingsince1985/geo-golang) | 341 | 11 | 2014-12-04 | 1 month ago | Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs. |
| [go-marathon](https://github.com/gambol99/go-marathon) | 192 | 13 | 2015-02-11 | 4 days ago | Go library for interacting with Mesosphere's Marathon PAAS. |
| [ethrpc](https://github.com/onrik/ethrpc) | 180 | 13 | 2017-01-24 | 2 weeks ago | Go bindings for Ethereum JSON RPC API. |
| [trello](https://github.com/adlio/trello) | 128 | 6 | 2016-09-24 | 6 days ago | Go wrapper for the Trello API. |
| [gostorm](https://github.com/jsgilmore/gostorm) | 124 | 11 | 2013-07-22 | 2 years ago | GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. |
| [medium-sdk-go](https://github.com/Medium/medium-sdk-go) | 123 | 124 | 2015-09-26 | 1 year ago | Golang SDK for Medium's OAuth2 API. |
| [hipchat](https://github.com/daneharrigan/hipchat) | 112 | 7 | 2013-04-28 | 2 years ago | A golang package to communicate with HipChat over XMPP. |
| [go-trending](https://github.com/andygrunwald/go-trending) | 108 | 7 | 2015-07-04 | 1 month ago | Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github. |
| [hipchat](https://github.com/andybons/hipchat) | 108 | 8 | 2012-10-20 | 4 years ago | This project implements a golang client library for the Hipchat API. |
| [cachet](https://github.com/andygrunwald/cachet) | 78 | 5 | 2015-10-31 | 1 year ago | Go client library for [Cachet (open source status page system)](https://cachethq.io/). |
| [pushover](https://github.com/gregdel/pushover) | 70 | 3 | 2015-02-19 | 1 year ago | Go wrapper for the Pushover API. |
| [wit-go](https://github.com/wit-ai/wit-go) | 65 | 12 | 2018-08-20 | 2 weeks ago | Go client for wit.ai HTTP API. |
| [clarifai-go](https://github.com/Clarifai/clarifai-go) | 57 | 37 | 2015-09-28 | 2 years ago | Go client library for interfacing with the Clarifai API. |
| [igdb](https://github.com/Henry-Sarabia/igdb) | 57 | 2 | 2017-08-24 | 1 month ago | Go client for the [Internet Game Database API](https://api.igdb.com/). |
| [megos](https://github.com/andygrunwald/megos) | 56 | 5 | 2015-10-02 | 1 year ago | Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster. |
| [go-circleci](https://github.com/jszwedko/go-circleci) | 49 | 3 | 2015-08-14 | 3 months ago | Go client library for interacting with CircleCI's API. |
| [gads](https://github.com/emiddleton/gads) | 47 | 7 | 2014-01-20 | 5 months ago | Google Adwords Unofficial API. |
| [go-amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) | 40 | 1 | 2016-11-15 | 1 year ago | Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html). |
| [go-xkcd](https://github.com/nishanths/go-xkcd) | 40 | 4 | 2016-02-26 | 3 years ago | Go client for the xkcd API. |
| [gomusicbrainz](https://github.com/michiwend/gomusicbrainz) | 40 | 8 | 2014-09-10 | 6 months ago | Go MusicBrainz WS2 client library. |
| [fcm](https://github.com/maddevsio/fcm) | 36 | 4 | 2017-01-06 | 1 week ago | Go library for Firebase Cloud Messaging. |
| [simples3](https://github.com/rhnvrm/simples3) | 36 | 1 | 2018-12-06 | 2 months ago | Simple no frills AWS S3 Library using REST with V4 Signing written in Go. |
| [uptimerobot](https://github.com/bitfield/uptimerobot) | 36 | 3 | 2018-05-29 | 4 months ago | Go wrapper and command-line client for the Uptime Robot v2 API. |
| [gosip](https://github.com/koltyakov/gosip) | 35 | 2 | 2019-01-26 | 1 week ago | Go client library SharePoint API. |
| [golyrics](https://github.com/mamal72/golyrics) | 34 | 3 | 2016-11-18 | 1 year ago | Golyrics is a Go library to fetch music lyrics data from the Wikia website. |
| [mixpanel](https://github.com/dukex/mixpanel) | 32 | 3 | 2014-05-20 | 1 month ago | Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications. |
| [gcm](https://github.com/TheOrioli/gcm) | 29 | 3 | 2015-11-09 | 4 years ago | Go library for Google Cloud Messaging. |
| [translate](https://github.com/nuveo/translate) | 29 | 27 | 2015-07-13 | 4 years ago | Go online translation package. |
| [golang-tmdb](https://github.com/cyruzin/golang-tmdb) | 28 | 1 | 2019-01-11 | 1 week ago | Golang wrapper for The Movie Database API v3. |
| [gami](https://github.com/bit4bit/gami) | 27 | 4 | 2014-05-14 | 1 year ago | Go library for Asterisk Manager Interface. |
| [go-unsplash](https://github.com/hbagdi/go-unsplash) | 27 | 1 | 2017-01-19 | 4 days ago | Go client library for the [Unsplash.com](https://unsplash.com) API. |
| [ynab.go](https://github.com/brunomvsouza/ynab.go) | 24 | 1 | 2018-07-13 | 2 months ago | Go wrapper for the YNAB API. |
| [go-spotify](https://github.com/rapito/go-spotify) | 21 | 1 | 2014-10-30 | 2 years ago | Go Library to access Spotify WEB API. |
| [go-twitch](https://github.com/knspriggs/go-twitch) | 19 | 5 | 2016-06-28 | 2 years ago | Go client for interacting with the Twitch v3 API. |
| [patreon-go](https://github.com/mxpv/patreon-go) | 19 | 4 | 2017-08-06 | 6 months ago | Go library for Patreon API. |
| [go-shopify](https://github.com/rapito/go-shopify) | 19 | 1 | 2014-10-28 | 2 years ago | Go Library to make CRUD request to the Shopify API. |
| [go-steam](https://github.com/sostronk/go-steam) | 19 | 10 | 2014-11-23 | 2 months ago | Go Library to interact with Steam game servers. |
| [brewerydb](https://github.com/naegelejd/brewerydb) | 17 | 3 | 2015-04-15 | 4 years ago | Go library for accessing the BreweryDB API. |
| [codeship-go](https://github.com/codeship/codeship-go) | 17 | 23 | 2017-09-08 | 11 months ago | Go client library for interacting with Codeship's API v2. |
| [go-myanimelist](https://github.com/nstratos/go-myanimelist) | 16 | 2 | 2015-05-03 | 3 years ago | Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api). |
| [textbelt](https://github.com/farmergreg/textbelt) | 16 | 2 | 2015-09-01 | 4 years ago | Go client for the textbelt.com txt messaging API. |
| [coinpaprika-api-go-client](https://github.com/coinpaprika/coinpaprika-api-go-client) | 12 | 8 | 2018-09-25 | 1 year ago | Go client library for interacting with Coinpaprika's API. |
| [go-google-analytics](https://github.com/chonthu/go-google-analytics) | 11 | 2 | 2015-06-01 | 4 years ago | Simple wrapper for easy google analytics reporting. |
| [go-hacknews](https://github.com/PaulRosset/go-hacknews) | 10 | 2 | 2017-08-10 | 2 years ago | Tiny Go client for HackerNews API. |
| [lastpass-go](https://github.com/ansd/lastpass-go) | 10 | 0 | 2019-07-11 | 1 month ago | Go client library for the [LastPass](https://www.lastpass.com/) API. |
| [smitego](https://github.com/sergiotapia/smitego) | 10 | 0 | 2013-12-11 | 5 years ago | Go package to wraps access to the Smite game API. |
| [rrdaclient](https://github.com/Omie/rrdaclient) | 8 | 1 | 2014-09-15 | 5 years ago | Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP. |
| [go-sptrans](https://github.com/sergioaugrod/go-sptrans) | 7 | 2 | 2017-09-11 | 2 years ago | Go client library for the SPTrans Olho Vivo API. |
| [go-here](https://github.com/abdullahselek/go-here) | 6 | 1 | 2019-07-07 | 1 week ago | Go client library around the HERE location based APIs. |
| [go-google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) | 6 | 0 | 2016-10-24 | 3 years ago | Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/). |
| [gumblr](https://github.com/mattcunningham/gumblr) | 6 | 1 | 2015-07-09 | 3 years ago | Go wrapper for the Tumblr v2 API. |
| [go-sophos](https://github.com/esurdam/go-sophos) | 5 | 1 | 2018-09-05 | 4 months ago | Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies. |
| [go-zooz](https://github.com/gojuno/go-zooz) | 5 | 16 | 2017-07-04 | 1 year ago | Go client for the Zooz API. |
| [go-postman-collection](https://github.com/rbretecher/go-postman-collection) | 5 | 1 | 2019-11-16 | 1 month ago | Go module to work with [Postman Collections](https://learning.getpostman.com/docs/postman/collections/creating-collections/) (compatible with Insomnia). |
| [go-chronos](https://github.com/axelspringer/go-chronos) | 3 | 9 | 2017-10-23 | 2 years ago | Go library for interacting with the [Chronos](https://mesos.github. |
| [twitter-scraper](https://github.com/n0madic/twitter-scraper) | 3 | 1 | 2018-11-29 | 1 week ago | Scrape the Twitter Frontend API without authentication and limits. |
| [google-play-scraper](https://github.com/n0madic/google-play-scraper) | 2 | 1 | 2019-09-20 | 8 hours ago | Get data from Google Play Store. |
| [gopaapi5](https://github.com/utekaravinash/gopaapi5) | 2 | 1 | 2020-02-15 | 2 days ago | Go Client Library for [Amazon Product Advertising API 5.0](https://webservices.amazon.com/paapi5/documentation/). |
| [libgoffi](https://github.com/clevabit/libgoffi) | 1 | 2 | 2019-08-03 | 7 months ago | Library adapter toolbox for native [libffi](http://sourceware. |
| [playlyfe-go-sdk](https://github.com/playlyfe/playlyfe-go-sdk) | 1 | 4 | 2015-05-25 | 4 years ago | The Playlyfe Rest API Go SDK. |
| [tripadvisor-golang](https://github.com/mrbenosborne/tripadvisor-golang) | 1 | 1 | 2019-04-15 | 4 months ago | Go wrapper for the TripAdvisor API. |
| [vl-go](https://github.com/verifid/vl-go) | 1 | 1 | 2019-02-09 | 1 week ago | Go client library around the VerifID identity verification layer API. |
| [slack](https://github.com/nlopes/slack) | 0 | 0 | 2015-01-24 | 3 weeks ago | Slack API in Go. |

## Utilities
        
*General utilities and tools to make your life easier.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [fzf](https://github.com/junegunn/fzf) | 27710 | 363 | 2013-10-23 | 39 minutes ago | Command-line fuzzy finder written in Go. |
| [hub](https://github.com/github/hub) | 19221 | 473 | 2009-12-05 | 2 hours ago | wrap git commands with additional functionality to interact with github from the terminal. |
| [delve](https://github.com/go-delve/delve) | 13355 | 387 | 2014-05-20 | 4 weeks ago | Go debugger. |
| [ctop](https://github.com/bcicen/ctop) | 9602 | 163 | 2016-12-27 | 2 months ago | [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics. |
| [wuzz](https://github.com/asciimoo/wuzz) | 8574 | 177 | 2017-01-30 | 1 month ago | Interactive cli tool for HTTP inspection. |
| [sqlx](https://github.com/jmoiron/sqlx) | 7941 | 177 | 2013-01-28 | 1 week ago | provides a set of extensions on top of the excellent built-in database/sql package. |
| [usql](https://github.com/xo/usql) | 5801 | 118 | 2017-03-02 | 1 week ago | usql is a universal command-line interface for SQL databases. |
| [peco](https://github.com/peco/peco) | 5785 | 135 | 2014-06-06 | 2 months ago | Simplistic interactive filtering tool. |
| [goreleaser](https://github.com/goreleaser/goreleaser) | 5291 | 83 | 2016-12-21 | 1 hour ago | Deliver Go binaries as fast and easily as possible. |
| [godropbox](https://github.com/dropbox/godropbox) | 3846 | 248 | 2014-06-22 | 2 weeks ago | Common libraries for writing Go services/applications from Dropbox. |
| [realize](https://github.com/oxequa/realize) | 3561 | 70 | 2016-07-12 | 2 months ago | Go build system with file watchers and live reload. Run, build and watch file changes with custom paths. |
| [goreporter](https://github.com/360EntSecGroup-Skylar/goreporter) | 2622 | 97 | 2017-03-27 | 1 year ago | Golang tool that does static analysis, unit testing, code review and generate code quality report. |
| [hystrix-go](https://github.com/afex/hystrix-go) | 2384 | 84 | 2013-12-15 | 9 months ago | Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker. |
| [task](https://github.com/go-task/task) | 2273 | 40 | 2017-02-27 | 2 weeks ago | simple "Make" alternative. |
| [panicparse](https://github.com/maruel/panicparse) | 2224 | 41 | 2015-02-02 | 20 hours ago | Groups similar goroutines and colorizes stack dump. |
| [minify](https://github.com/tdewolff/minify) | 2047 | 49 | 2014-05-21 | 2 days ago | Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats. |
| [go-funk](https://github.com/thoas/go-funk) | 1649 | 37 | 2016-12-30 | 3 days ago | Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...). |
| [storm](https://github.com/asdine/storm) | 1494 | 47 | 2016-01-10 | 2 months ago | Simple and powerful toolkit for BoltDB. |
| [mmake](https://github.com/tj/mmake) | 1493 | 29 | 2017-02-15 | 2 weeks ago | Modern Make. |
| [mc](https://github.com/minio/mc) | 1378 | 46 | 2015-01-16 | 11 hours ago | Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems. |
| [mole](https://github.com/davrodpin/mole) | 1356 | 32 | 2018-10-04 | 2 weeks ago | cli app to easily create ssh tunnels. |
| [boilr](https://github.com/tmrts/boilr) | 1058 | 29 | 2015-12-19 | 5 months ago | Blazingly fast CLI tool for creating projects from boilerplate templates. |
| [filetype](https://github.com/h2non/filetype) | 1058 | 26 | 2015-09-24 | 1 week ago | Small package to infer the file type checking the magic numbers signature. |
| [mergo](https://github.com/imdario/mergo) | 1051 | 18 | 2013-03-11 | 2 weeks ago | Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements. |
| [spinner](https://github.com/briandowns/spinner) | 1045 | 15 | 2014-12-13 | 1 month ago | Go package to easily provide a terminal spinner with options. |
| [circuitbreaker](https://github.com/rubyist/circuitbreaker) | 859 | 19 | 2014-07-17 | 4 months ago | Circuit Breakers in Go. |
| [jump](https://github.com/gsamokovarov/jump) | 811 | 15 | 2015-08-16 | 1 month ago | Jump helps you navigate faster by learning your habits. |
| [gtm](https://github.com/git-time-metric/gtm) | 786 | 27 | 2016-06-19 | 7 months ago | Simple, seamless, lightweight time tracking for Git. |
| [immortal](https://github.com/immortal/immortal) | 635 | 14 | 2016-06-30 | 3 months ago | \*nix cross-platform (OS agnostic) supervisor. |
| [htcat](https://github.com/htcat/htcat) | 511 | 17 | 2013-08-05 | 1 year ago | Parallel and Pipelined HTTP GET Utility. |
| [go-dry](https://github.com/ungerik/go-dry) | 446 | 14 | 2014-02-28 | 1 year ago | DRY (don't repeat yourself) package for Go. |
| [circuit](https://github.com/cep21/circuit) | 441 | 12 | 2017-12-23 | 4 months ago | An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern. |
| [godaemon](https://github.com/VividCortex/godaemon) | 431 | 31 | 2013-08-01 | 1 year ago | Utility to write daemons. |
| [gopencils](https://github.com/bndr/gopencils) | 431 | 14 | 2014-06-23 | 1 year ago | Small and simple package to easily consume REST APIs. |
| [request](https://github.com/mozillazg/request) | 371 | 14 | 2014-12-21 | 3 months ago | Go HTTP Requests for Humans™. |
| [koazee](https://github.com/wesovilabs/koazee) | 368 | 11 | 2018-11-09 | 1 week ago | Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays. |
| [ergo](https://github.com/cristianoliveira/ergo) | 363 | 6 | 2017-08-19 | 6 months ago | The management of multiple local services running over different ports made easy. |
| [go-rate](https://github.com/beefsack/go-rate) | 300 | 10 | 2014-08-25 | 1 year ago | Timed rate limiter for Go. |
| [clockwork](https://github.com/jonboulle/clockwork) | 258 | 4 | 2014-09-09 | 1 month ago | A simple fake clock for golang. |
| [gohper](https://github.com/zhuah/gohper) | 252 | 20 | 2015-03-23 | 2 years ago | Various tools/modules help for development. |
| [deepcopier](https://github.com/ulule/deepcopier) | 246 | 17 | 2015-07-24 | 1 month ago | Simple struct copying for Go. |
| [mimetype](https://github.com/gabriel-vasile/mimetype) | 219 | 5 | 2018-07-02 | 2 hours ago | Package for MIME type detection based on magic numbers. |
| [retry](https://github.com/kamilsk/retry) | 216 | 3 | 2016-11-02 | 1 week ago | The most advanced functional mechanism to perform actions repetitively until successful. |
| [gubrak](https://github.com/novalagung/gubrak) | 209 | 6 | 2018-03-09 | 1 month ago | Golang utility library with syntactic sugar. It's like lodash, but for golang. |
| [serve](https://github.com/syntaqx/serve) | 207 | 5 | 2019-01-10 | 2 months ago | A static http server anywhere you need. |
| [go-trigger](https://github.com/sadlil/go-trigger) | 193 | 12 | 2015-10-19 | 3 years ago | Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project. |
| [gotenv](https://github.com/subosito/gotenv) | 169 | 4 | 2013-08-27 | 6 months ago | Load environment variables from `.env` or any `io.Reader` in Go. |
| [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) | 165 | 3 | 2016-11-08 | 6 months ago | go:generate tool for wrapping symbols exported by golang plugins (1.8 only). |
| [util](https://github.com/shomali11/util) | 158 | 9 | 2017-05-24 | 9 months ago | Collection of useful utility functions. (strings, concurrency, manipulations, ...). |
| [rerun](https://github.com/ivpusic/rerun) | 155 | 7 | 2014-12-10 | 2 years ago | Recompiling and rerunning go apps when source changes. |
| [death](https://github.com/vrecan/death) | 149 | 4 | 2015-03-09 | 1 year ago | Managing go application shutdown with signals. |
| [moldova](https://github.com/StabbyCutyou/moldova) | 149 | 5 | 2016-01-30 | 2 years ago | Utility for generating random data based on an input template. |
| [robustly](https://github.com/VividCortex/robustly) | 140 | 18 | 2013-07-08 | 2 years ago | Runs functions resiliently, catching and restarting panics. |
| [apm](https://github.com/topfreegames/apm) | 137 | 16 | 2015-11-18 | 3 years ago | Process manager for Golang applications with an HTTP API. |
| [toolbox](https://github.com/viant/toolbox) | 126 | 15 | 2016-06-13 | 3 weeks ago | Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer. |
| [chyle](https://github.com/antham/chyle) | 117 | 6 | 2016-11-17 | 2 weeks ago | Changelog generator using a git repository with multiple configuration possibilities. |
| [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) | 116 | 5 | 2015-10-12 | 6 months ago | XML Sitemap generator written in Go. |
| [onecache](https://github.com/adelowo/onecache) | 106 | 7 | 2017-04-14 | 9 months ago | Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc). |
| [lrserver](https://github.com/jaschaephraim/lrserver) | 105 | 5 | 2014-07-15 | 2 years ago | LiveReload server for Go. |
| [go-bsdiff](https://github.com/gabstv/go-bsdiff) | 91 | 1 | 2019-02-23 | 1 year ago | Pure Go bsdiff and bspatch libraries and CLI tools. |
| [pm](https://github.com/VividCortex/pm) | 78 | 18 | 2013-11-17 | 4 months ago | Process (i.e. goroutine) manager with an HTTP API. |
| [go-health](https://github.com/Talento90/go-health) | 74 | 2 | 2018-02-13 | 1 year ago | Health package simplifies the way you add health check to your services. |
| [xferspdy](https://github.com/monmohan/xferspdy) | 73 | 4 | 2015-05-22 | 3 years ago | Xferspdy provides binary diff and patch library in golang. |
| [mssqlx](https://github.com/linxGnu/mssqlx) | 72 | 8 | 2016-12-26 | 1 month ago | Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind. |
| [unis](https://github.com/esemplastic/unis) | 69 | 4 | 2017-05-06 | 2 years ago | Common Architecture™ for String Utilities in Go. |
| [multitick](https://github.com/VividCortex/multitick) | 63 | 18 | 2013-12-10 | 4 months ago | Multiplexor for aligned tickers. |
| [repeat](https://github.com/ssgreg/repeat) | 63 | 4 | 2017-11-22 | 1 month ago | Go implementation of different backoff strategies useful for retrying operations and heartbeating. |
| [minquery](https://github.com/icza/minquery) | 53 | 3 | 2016-11-16 | 4 months ago | MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off). |
| [mimemagic](https://github.com/zRedShift/mimemagic) | 52 | 1 | 2018-10-11 | 1 year ago | Pure Go ultra performant MIME sniffing library/utility. |
| [handy](https://github.com/miguelpragier/handy) | 50 | 5 | 2018-06-13 | 1 week ago | Many utilities and helpers like string handlers/formatters and validators. |
| [golog](https://github.com/mlimaloureiro/golog) | 48 | 3 | 2016-01-09 | 1 year ago | Easy and lightweight CLI tool to time track your tasks. |
| [go-astitodo](https://github.com/asticode/go-astitodo) | 46 | 2 | 2016-10-17 | 1 year ago | Parse TODOs in your GO code. |
| [goreadability](https://github.com/philipjkim/goreadability) | 46 | 6 | 2016-04-20 | 10 months ago | Webpage summary extractor using Facebook Open Graph and arc90's readability. |
| [goseaweedfs](https://github.com/linxGnu/goseaweedfs) | 45 | 6 | 2017-07-20 | 6 days ago | SeaweedFS client library with almost full features. |
| [pgo](https://github.com/arthurkushman/pgo) | 45 | 6 | 2018-12-26 | 1 day ago | Convenient functions for PHP community. |
| [gaper](https://github.com/maxcnunes/gaper) | 42 | 0 | 2018-06-16 | 3 months ago | Builds and restarts a Go project when it crashes or some watched file changes. |
| [goback](https://github.com/carlescere/goback) | 42 | 1 | 2015-03-13 | 2 years ago | Go simple exponential backoff package. |
| [intrinsic](https://github.com/mengzhuo/intrinsic) | 41 | 3 | 2017-06-13 | 2 years ago | Use x86 SIMD without writing any assembly code. |
| [copy-pasta](https://github.com/jutkko/copy-pasta) | 38 | 4 | 2017-01-28 | 7 months ago | Universal multi-workstation clipboard that uses S3 like backend for the storage. |
| [golarm](https://github.com/msempere/golarm) | 38 | 1 | 2015-08-14 | 4 years ago | Fire alarms with system events. |
| [retry](https://github.com/thedevsaddam/retry) | 36 | 1 | 2018-02-25 | 2 years ago | Simple and easy retry mechanism package for Go. |
| [go-pattern-match](https://github.com/alexpantyukhin/go-pattern-match) | 36 | 2 | 2018-12-11 | 3 weeks ago | Pattern matching libray. |
| [beyond](https://github.com/wesovilabs/beyond) | 34 | 1 | 2019-10-18 | 3 months ago | The Go tool that will drive you to the AOP world! |
| [retry-go](https://github.com/rafaeljesus/retry-go) | 34 | 1 | 2017-06-09 | 1 year ago | Retrying made simple and easy for golang. |
| [gpath](https://github.com/tenntenn/gpath) | 33 | 2 | 2017-05-24 | 2 years ago | Library to simplify access struct fields with Go's expression in reflection. |
| [myhttp](https://github.com/inancgumus/myhttp) | 33 | 1 | 2017-09-13 | 1 year ago | Simple API to make HTTP GET requests with timeout support. |
| [rclient](https://github.com/zpatrick/rclient) | 33 | 3 | 2017-02-28 | 3 months ago | Readable, flexible, simple-to-use client for REST APIs. |
| [dbt](https://github.com/nikogura/dbt) | 26 | 2 | 2017-11-30 | 15 hours ago | A framework for running self-updating signed binaries from a central, trusted repository. |
| [scan](https://github.com/blockloop/scan) | 25 | 1 | 2017-11-27 | 3 weeks ago | Scan golang `sql.Rows` directly to structs, slices, or primitive types. |
| [countries](https://github.com/biter777/countries) | 25 | 3 | 2019-04-22 | 5 hours ago | Full implementation of ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and IANA ccTLD standarts. |
| [generate](https://github.com/go-playground/generate) | 23 | 2 | 2015-11-15 | 3 years ago | runs go generate recursively on a specified path or environment variable and can filter by regex. |
| [slice](https://github.com/psampaz/slice) | 23 | 1 | 2019-11-26 | 2 weeks ago | Type-safe functions for common Go slice operations. |
| [gostrutils](https://github.com/ik5/gostrutils) | 22 | 3 | 2018-09-19 | 3 months ago | Collections of string manipulation and conversion functions. |
| [ugo](https://github.com/alxrm/ugo) | 22 | 1 | 2016-02-17 | 3 years ago | ugo is slice toolbox with concise syntax for Go. |
| [cmd](https://github.com/SimonBaeumer/cmd) | 21 | 2 | 2019-09-27 | 2 months ago | Library for executing shell commands on osx, windows and linux. |
| [goplaceholder](https://github.com/michiwend/goplaceholder) | 21 | 2 | 2014-10-12 | 4 years ago | a small golang lib to generate placeholder images. |
| [evaluator](https://github.com/nullne/evaluator) | 20 | 1 | 2017-04-27 | 2 years ago | Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend. |
| [filter](https://github.com/gookit/filter) | 20 | 4 | 2018-09-26 | 3 months ago | provide filtering, sanitizing, and conversion of Go data. |
| [delve](https://github.com/derekparker/delve) | 19 | 0 | 2020-02-18 | 2 days ago | Go debugger. |
| [go-httpheader](https://github.com/mozillazg/go-httpheader) | 16 | 2 | 2017-06-24 | 1 year ago | Go library for encoding structs into Header fields. |
| [dlog](https://github.com/kirillDanshin/dlog) | 15 | 2 | 2016-07-04 | 2 years ago | Compile-time controlled logger to make your release smaller without removing debug calls. |
| [filler](https://github.com/yaronsumel/filler) | 15 | 1 | 2017-04-05 | 2 years ago | small utility to fill structs using "fill" tag. |
| [slicer](https://github.com/leaanthony/slicer) | 15 | 1 | 2019-01-10 | 2 months ago | Makes working with slices easier. |
| [okrun](https://github.com/xta/okrun) | 14 | 2 | 2014-10-01 | 5 years ago | go run error steamroller. |
| [rerate](https://github.com/abo/rerate) | 14 | 4 | 2016-05-24 | 3 years ago | Redis-based rate counter and rate limiter for Go. |
| [structs](https://github.com/PumpkinSeed/structs) | 14 | 2 | 2017-08-26 | 2 years ago | Implement simple functions to manipulate structs. |
| [r](https://github.com/is5/r) | 14 | 1 | 2020-02-20 | 3 weeks ago | Python-like `range()` experience for Go. |
| [ghokin](https://github.com/antham/ghokin) | 13 | 2 | 2018-08-03 | 3 weeks ago | Parallelized formatter with no external dependencies for gherkin (cucumber, behat...). |
| [rest-go](https://github.com/edermanoel94/rest-go) | 13 | 3 | 2019-07-29 | 1 week ago | A package that provide many helpful methods for working with rest api. |
| [limiters](https://github.com/mennanov/limiters) | 12 | 1 | 2019-08-28 | 4 months ago | Rate limiters for distributed applications in Golang with configurable back-ends and distributed locks. |
| [ctxutil](https://github.com/posener/ctxutil) | 11 | 1 | 2018-07-30 | 2 weeks ago | A collection of utility functions for contexts. |
| [mimesniffer](https://github.com/aofei/mimesniffer) | 11 | 1 | 2018-12-20 | 2 weeks ago | A MIME type sniffer for Go. |
| [retry](https://github.com/shafreeck/retry) | 11 | 0 | 2018-07-18 | 1 month ago | A pretty simple library to ensure your work to be done. |
| [command](https://github.com/txgruppi/command) | 10 | 1 | 2015-08-24 | 3 years ago | Command pattern for Go with thread safe serial and parallel dispatcher. |
| [backscanner](https://github.com/icza/backscanner) | 9 | 2 | 2017-10-18 | 1 month ago | A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward. |
| [shutdown](https://github.com/ztrue/shutdown) | 9 | 1 | 2018-11-17 | 1 year ago | App shutdown hooks for `os.Signal` handling. |
| [tome](https://github.com/cyruzin/tome) | 9 | 1 | 2019-04-12 | 5 months ago | Tome was designed to paginate simple RESTful APIs. |
| [jsend](https://github.com/clevergo/jsend) | 8 | 3 | 2020-01-14 | 6 days ago | JSend's implementation writen in Go. |
| [retry](https://github.com/percolate/retry) | 6 | 30 | 2018-06-15 | 6 months ago | A simple but highly configurable retry package for Go. |
| [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) | 5 | 1 | 2019-05-16 | 4 weeks ago | Go package for working with Problem Details. |
| [sliceconv](https://github.com/Henry-Sarabia/sliceconv) | 5 | 1 | 2019-02-15 | 1 month ago | Slice conversion between primitive types. |
| [blank](https://github.com/Henry-Sarabia/blank) | 4 | 2 | 2019-02-13 | 7 months ago | Verify or remove blanks and whitespace from strings. |
| [silk](https://github.com/chrispassas/silk) | 4 | 1 | 2018-12-18 | 7 months ago | Read silk netflow files. |
| [ptr](https://github.com/gotidy/ptr) | 3 | 1 | 2019-12-25 | 2 months ago | Package that provide functions for simplified creation of pointers from constants of basic types. |
| [olaf](https://github.com/btnguyen2k/olaf) | 1 | 1 | 2019-01-03 | 11 months ago | Twitter Snowflake implemented in Go. |
| [compare](https://github.com/posener/compare) | 1 | 1 | 2020-03-13 | 12 hours ago | Enables more readable and easier comparison tasks. |

## UUID
        
*Libraries for working with UUIDs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [ulid](https://github.com/oklog/ulid) | 1845 | 38 | 2016-12-06 | 6 months ago | Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier). |
| [uuid](https://github.com/gofrs/uuid) | 662 | 16 | 2018-07-13 | 4 months ago | Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid. |
| [wuid](https://github.com/edwingeng/wuid) | 332 | 16 | 2018-01-27 | 3 months ago | An extremely fast unique number generator, 10-135 times faster than UUID. |
| [Goid](https://github.com/JakeHL/Goid) | 27 | 4 | 2017-05-19 | 1 year ago | Generate and Parse RFC4122 compliant V4 UUIDs. |
| [sno](https://github.com/muyo/sno) | 22 | 1 | 2019-05-26 | 2 weeks ago | Compact, sortable and fast unique IDs with embedded metadata. |
| [nanoid](https://github.com/aidarkhanov/nanoid) | 16 | 1 | 2019-07-02 | 2 months ago | A tiny and efficient Go unique string ID generator. |
| [uuid](https://github.com/agext/uuid) | 10 | 1 | 2016-02-03 | 3 days ago | Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier. |

## Validation
        
*Libraries for validation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [validator](https://github.com/go-playground/validator) | 4890 | 83 | 2015-02-12 | 6 days ago | Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving. |
| [govalidator](https://github.com/asaskevich/govalidator) | 4077 | 101 | 2014-06-20 | 1 week ago | Validators and sanitizers for strings, numerics, slices and structs. |
| [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) | 1309 | 26 | 2016-06-22 | 1 month ago | Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags. |
| [govalidator](https://github.com/thedevsaddam/govalidator) | 845 | 20 | 2017-09-13 | 3 months ago | Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. |
| [validate](https://github.com/gookit/validate) | 205 | 10 | 2018-07-16 | 2 months ago | Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features. |
| [checkdigit](https://github.com/osamingo/checkdigit) | 51 | 0 | 2019-04-05 | 2 months ago | Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.). |
| [validate](https://github.com/gobuffalo/validate) | 31 | 7 | 2018-02-10 | 1 hour ago | This package provides a framework for writing validations for Go applications. |
| [jio](https://github.com/faceair/jio) | 30 | 2 | 2018-10-28 | 9 months ago | jio is a json schema validator similar to [joi](https://github.com/hapijs/joi). |
| [terraform-validator](https://github.com/thazelart/terraform-validator) | 26 | 2 | 2019-05-29 | 1 month ago | A norms and conventions validator for Terraform. |
| [gody](https://github.com/guiferpa/gody) | 19 | 0 | 2018-11-01 | 2 weeks ago | :balloon: A lightweight struct validator for Go. |
| [govalid](https://github.com/twharmon/govalid) | 13 | 1 | 2019-02-17 | 4 weeks ago | Fast, tag-based validation for structs. |

## Version Control
        
*Libraries for version control.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-git](https://github.com/src-d/go-git) | 5031 | 103 | 2015-10-22 | 7 hours ago | highly extensible Git implementation in pure Go. |
| [git2go](https://github.com/libgit2/git2go) | 1459 | 57 | 2013-03-05 | 1 week ago | Go bindings for libgit2. |
| [hercules](https://github.com/src-d/hercules) | 907 | 22 | 2016-12-12 | 3 weeks ago | gaining advanced insights from Git repository history. |
| [gh](https://github.com/rjeczalik/gh) | 72 | 6 | 2015-03-08 | 1 year ago | Scriptable server and net/http middleware for GitHub Webhooks. |
| [go-vcs](https://github.com/sourcegraph/go-vcs) | 72 | 37 | 2013-06-02 | 5 months ago | manipulate and inspect VCS repositories in Go. |
| [hgo](https://github.com/beyang/hgo) | 12 | 4 | 2014-06-18 | 4 years ago | Hgo is a collection of Go packages providing read-access to local Mercurial repositories. |

## Video
        
*Libraries for manipulating video.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [goav](https://github.com/giorgisio/goav) | 1085 | 43 | 2015-05-21 | 3 weeks ago | Comphrensive Go bindings for FFmpeg. |
| [m3u8](https://github.com/grafov/m3u8) | 676 | 34 | 2013-02-05 | 3 hours ago | Parser and generator library of M3U8 playlists for Apple HLS. |
| [gmf](https://github.com/3d0c/gmf) | 586 | 32 | 2013-04-03 | 7 months ago | Go bindings for FFmpeg av\* libraries. |
| [go-astits](https://github.com/asticode/go-astits) | 297 | 16 | 2017-07-04 | 2 months ago | Parse and demux MPEG Transport Streams (.ts) natively in GO. |
| [go-astisub](https://github.com/asticode/go-astisub) | 230 | 7 | 2016-12-16 | 2 months ago | Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.). |
| [gst](https://github.com/ziutek/gst) | 155 | 10 | 2011-07-26 | 1 year ago | Go bindings for GStreamer. |
| [libvlc-go](https://github.com/adrg/libvlc-go) | 97 | 8 | 2015-01-06 | 20 minutes ago | Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player). |
| [go-m3u8](https://github.com/quangngotan95/go-m3u8) | 47 | 1 | 2018-11-06 | 10 months ago | Parser and generator library for Apple m3u8 playlists. |
| [v4l](https://github.com/korandiz/v4l) | 38 | 4 | 2016-10-25 | 1 year ago | Video capture library for Linux, written in Go. |
| [libgosubs](https://github.com/wargarblgarbl/libgosubs) | 13 | 1 | 2017-05-03 | 1 year ago | Subtitle format support for go. Supports .srt, .ttml, and .ass. |
| [go-mpd](https://github.com/unki2aut/go-mpd) | 0 | 1 | 2018-11-02 | 2 weeks ago | Parser and generator library for MPEG-DASH manifest files. |

## Web Frameworks
        
*Full stack web frameworks.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gin](https://github.com/gin-gonic/gin) | 36140 | 1217 | 2014-06-16 | 23 minutes ago | Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity. |
| [beego](https://github.com/astaxie/beego) | 23434 | 1269 | 2012-02-29 | 1 week ago | beego is an open-source, high-performance web framework for the Go programming language. |
| [echo](https://github.com/labstack/echo) | 16689 | 546 | 2015-03-01 | 6 days ago | High performance, minimalist Go web framework. |
| [revel](https://github.com/revel/revel) | 11612 | 554 | 2011-12-09 | 1 month ago | High-productivity web framework for the Go language. |
| [goa](https://github.com/goadesign/goa) | 3756 | 166 | 2014-12-05 | 1 week ago | Goa provides a holistic approach for developing remote APIs and microservices in Go. |
| [go-json-rest](https://github.com/ant0ine/go-json-rest) | 3392 | 162 | 2013-02-19 | 6 months ago | Quick and easy way to setup a RESTful JSON API. |
| [fiber](https://github.com/gofiber/fiber) | 3364 | 58 | 2020-01-16 | 56 minutes ago | An Express.js inspired web framework build on Fasthttp. |
| [gizmo](https://github.com/nytimes/gizmo) | 3067 | 113 | 2015-12-15 | 1 hour ago | Microservice toolkit used by the New York Times. |
| [macaron](https://github.com/go-macaron/macaron) | 2943 | 147 | 2014-07-10 | 1 week ago | Macaron is a high productive and modular design web framework in Go. |
| [utron](https://github.com/gernest/utron) | 2160 | 70 | 2015-09-16 | 1 year ago | Lightweight MVC framework for Go(Golang). |
| [go-tigertonic](https://github.com/rcrowley/go-tigertonic) | 1000 | 46 | 2013-02-09 | 1 year ago | Go framework for building JSON web services inspired by Dropwizard. |
| [tango](https://github.com/lunny/tango) | 833 | 77 | 2014-12-17 | 10 months ago | Micro & pluggable web framework for Go. |
| [gongular](https://github.com/mustafaakin/gongular) | 425 | 22 | 2016-06-22 | 1 year ago | Fast Go web framework with input mapping/validation and (DI) Dependency Injection. |
| [neo](https://github.com/ivpusic/neo) | 405 | 32 | 2015-02-04 | 2 years ago | Neo is minimal and fast Go Web Framework with extremely simple API. |
| [air](https://github.com/aofei/air) | 373 | 17 | 2016-07-20 | 5 days ago | An ideally refined web framework for Go. |
| [mango](https://github.com/paulbellamy/mango) | 346 | 22 | 2011-05-25 | 2 years ago | Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. |
| [gondola](https://github.com/rainycape/gondola) | 315 | 15 | 2014-07-25 | 1 year ago | The web framework for writing faster sites, faster. |
| [aero](https://github.com/aerogo/aero) | 250 | 15 | 2016-11-09 | 1 month ago | High-performance web framework for Go, reaches top scores in Lighthouse. |
| [golf](https://github.com/dinever/golf) | 242 | 19 | 2015-11-18 | 3 years ago | Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library. |
| [goyave](https://github.com/System-Glitch/goyave) | 227 | 12 | 2019-10-21 | 4 days ago | Feature-complete web framework aimed at clean code and fast development, with powerful built-in functionalities. |
| [hiboot](https://github.com/hidevopsio/hiboot) | 119 | 12 | 2018-03-16 | 3 weeks ago | hiboot is a high performance web application framework with auto configuration and dependency injection support. |
| [go-rest](https://github.com/ungerik/go-rest) | 118 | 10 | 2012-07-13 | 3 years ago | Small and evil REST framework for Go. |
| [flamingo](https://github.com/i-love-flamingo/flamingo) | 109 | 21 | 2019-04-02 | 4 hours ago | Framework for pluggable web projects. Including a concept for modules and offering features for DI, Configareas, i18n, template engines, graphql, observability, security, events, routing & reverse routing etc. |
| [webgo](https://github.com/bnkamalesh/webgo) | 89 | 3 | 2015-12-16 | 1 month ago | A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc). |
| [uadmin](https://github.com/uadmin/uadmin) | 83 | 10 | 2018-10-05 | 4 days ago | Fully featured web framework for Golang, inspired by Django. |
| [golax](https://github.com/fulldump/golax) | 71 | 7 | 2016-01-30 | 1 year ago | A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more. |
| [microservice](https://github.com/claygod/microservice) | 71 | 8 | 2016-12-15 | 9 months ago | The framework for the creation of microservices, written in Golang. |
| [ginrpc](https://github.com/xxjwxc/ginrpc) | 61 | 2 | 2019-06-22 | 1 day ago | Gin parameter automatic binding tool,gin rpc tools. |
| [yarf](https://github.com/yarf-framework/yarf) | 55 | 3 | 2015-09-02 | 1 year ago | Fast micro-framework designed to build REST APIs and web services in a fast and simple way. |
| [flamingo-commerce](https://github.com/i-love-flamingo/flamingo-commerce) | 51 | 17 | 2019-04-02 | 1 hour ago | Providing e-commerce features using clean architecture like DDD and ports and adapters, that you can use to build flexible e-commerce applications. |
| [patron](https://github.com/beatlabs/patron) | 51 | 16 | 2019-01-30 | 1 hour ago | Patron is a microservice framework following best cloud practices with a focus on productivity. |
| [fireball](https://github.com/zpatrick/fireball) | 50 | 4 | 2016-07-20 | 1 year ago | More "natural" feeling web framework. |
| [vox](https://github.com/aisk/vox) | 48 | 2 | 2014-12-24 | 1 month ago | A golang web framework for humans, inspired by Koa heavily. |
| [api](https://github.com/resoursea/api) | 31 | 5 | 2015-01-24 | 5 years ago | REST framework for quickly writing resource based services. |
| [goa](https://github.com/goa-go/goa) | 29 | 2 | 2019-07-26 | 3 months ago | goa is just like koajs for golang, it is a flexible, light, high-performance and extensible web framework based on middleware. |
| [rex](https://github.com/goanywhere/rex) | 29 | 1 | 2014-10-16 | 2 years ago | Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`. |
| [rux](https://github.com/gookit/rux) | 24 | 2 | 2018-08-05 | 2 months ago | Simple and fast web framework for build golang HTTP applications. |
| [banjo](https://github.com/nsheremet/banjo) | 11 | 1 | 2017-12-09 | 2 years ago | Very simple and fast web framework for Go. |
| [goweb](https://github.com/twharmon/goweb) | 3 | 1 | 2019-05-07 | 4 weeks ago | Web framework with routing, websockets, logging, middleware, static file server (optional gzip), and automatic TLS. |

### Middlewares
        

#### Actual middlewares
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [tollbooth](https://github.com/didip/tollbooth) | 1504 | 45 | 2015-05-17 | 6 months ago | Rate limit HTTP request handler. |
| [cors](https://github.com/rs/cors) | 1410 | 30 | 2014-10-25 | 4 days ago | Easily add CORS capabilities to your API. |
| [limiter](https://github.com/ulule/limiter) | 885 | 25 | 2015-10-02 | 1 week ago | Dead simple rate limit middleware for Go. |
| [go-server-timing](https://github.com/mitchellh/go-server-timing) | 773 | 20 | 2018-02-12 | 1 year ago | Add/parse Server-Timing header. |
| [ln-paywall](https://github.com/philippgille/ln-paywall) | 101 | 4 | 2018-06-29 | 1 year ago | Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin). |
| [xff](https://github.com/sebest/xff) | 74 | 2 | 2014-12-22 | 11 months ago | Handle `X-Forwarded-For` header and friends. |
| [formjson](https://github.com/rs/formjson) | 34 | 2 | 2015-03-19 | 4 years ago | Transparently handle JSON input as a standard form POST. |
| [client-timing](https://github.com/posener/client-timing) | 16 | 1 | 2018-02-23 | 2 days ago | An HTTP client for Server-Timing header. |

#### Libraries for creating HTTP middlewares
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [negroni](https://github.com/urfave/negroni) | 6572 | 240 | 2014-05-18 | 5 months ago | Idiomatic HTTP middleware for Golang. |
| [alice](https://github.com/justinas/alice) | 1975 | 51 | 2014-05-25 | 4 months ago | Painless middleware chaining for Go. |
| [render](https://github.com/unrolled/render) | 1340 | 40 | 2014-06-10 | 1 month ago | Go package for easily rendering JSON, XML, and HTML template responses. |
| [stats](https://github.com/thoas/stats) | 555 | 16 | 2015-03-05 | 11 months ago | Go middleware that stores various information about your web application. |
| [interpose](https://github.com/carbocation/interpose) | 288 | 12 | 2014-07-20 | 3 years ago | Minimalist net/http middleware for golang. |
| [muxchain](https://github.com/stephens2424/muxchain) | 209 | 5 | 2014-05-03 | 1 year ago | Lightweight middleware for net/http. |
| [renderer](https://github.com/thedevsaddam/renderer) | 186 | 7 | 2017-11-07 | 1 year ago | Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go. |
| [rye](https://github.com/InVisionApp/rye) | 93 | 187 | 2016-10-06 | 1 year ago | Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context. |
| [gores](https://github.com/alioygur/gores) | 88 | 5 | 2015-12-25 | 1 year ago | Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs. |
| [chain](https://github.com/codemodus/chain) | 64 | 6 | 2015-05-14 | 1 year ago | Handler wrapper chaining with scoped data (net/context-based "middleware"). |
| [wrap](https://github.com/go-on/wrap) | 59 | 3 | 2014-02-16 | 1 year ago | Small middlewares package for net/http. |
| [catena](https://github.com/codemodus/catena) | 7 | 1 | 2015-07-30 | 1 year ago | http.Handler wrapper catenation (same API as "chain"). |

### Routers
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [mux](https://github.com/gorilla/mux) | 11388 | 291 | 2012-10-02 | 3 weeks ago | Powerful URL router and dispatcher for golang. |
| [httprouter](https://github.com/julienschmidt/httprouter) | 10907 | 305 | 2013-12-05 | 1 week ago | High performance router. Use this and the standard http handlers to form a very high performance web framework. |
| [chi](https://github.com/go-chi/chi) | 7180 | 178 | 2015-10-15 | 20 hours ago | Small, fast and expressive HTTP router built on net/context. |
| [web](https://github.com/gocraft/web) | 1413 | 58 | 2013-11-16 | 8 months ago | Mux and middleware package in Go. |
| [bone](https://github.com/go-zoo/bone) | 1251 | 36 | 2014-11-19 | 10 months ago | Lightning Fast HTTP Multiplexer. |
| [fasthttprouter](https://github.com/buaazp/fasthttprouter) | 826 | 35 | 2015-12-13 | 10 months ago | High performance router forked from `httprouter`. The first router fit for `fasthttp`. |
| [goji](https://github.com/goji/goji) | 803 | 40 | 2015-11-16 | 7 months ago | Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`. |
| [gorouter](https://github.com/xujiajun/gorouter) | 473 | 15 | 2018-01-29 | 5 months ago | A simple and fast HTTP router for Go. |
| [httptreemux](https://github.com/dimfeld/httptreemux) | 413 | 23 | 2014-05-14 | 5 months ago | High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter. |
| [lars](https://github.com/go-playground/lars) | 380 | 15 | 2015-12-24 | 10 months ago | Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks. |
| [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) | 376 | 28 | 2015-10-27 | 1 month ago | An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs. |
| [siesta](https://github.com/VividCortex/siesta) | 353 | 27 | 2014-09-23 | 3 months ago | Composable framework to write middleware and handlers. |
| [vestigo](https://github.com/husobee/vestigo) | 261 | 17 | 2015-09-22 | 5 months ago | Performant, stand-alone, HTTP compliant URL Router for go web applications. |
| [router](https://github.com/gowww/router) | 159 | 6 | 2017-05-25 | 1 year ago | Lightning fast HTTP router fully compatible with the net/http.Handler interface. |
| [alien](https://github.com/gernest/alien) | 109 | 4 | 2016-01-30 | 11 months ago | Lightweight and fast http router from outer space. |
| [violetear](https://github.com/nbari/violetear) | 97 | 3 | 2015-06-19 | 4 months ago | Go HTTP router. |
| [Bxog](https://github.com/claygod/Bxog) | 93 | 8 | 2016-05-19 | 1 month ago | Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters. |
| [pure](https://github.com/go-playground/pure) | 92 | 5 | 2016-09-23 | 4 months ago | Is a lightweight HTTP router that sticks to the std "net/http" implementation. |
| [xmux](https://github.com/rs/xmux) | 88 | 6 | 2015-12-14 | 2 years ago | High performance muxer based on `httprouter` with `net/context` support. |
| [gorouter](https://github.com/vardius/gorouter) | 72 | 5 | 2016-07-14 | 3 weeks ago | GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`. |
| [bellt](https://github.com/GuilhermeCaruso/bellt) | 41 | 5 | 2019-02-21 | 9 months ago | A simple Go HTTP router. |
| [fastrouter](https://github.com/razonyang/fastrouter) | 19 | 2 | 2017-11-01 | 2 years ago | a fast, flexible HTTP router written in Go. |
| [route](https://github.com/goroute/route) | 5 | 3 | 2019-07-06 | 2 months ago | Simple yet powerful HTTP request multiplexer. |

## Windows
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-ole](https://github.com/go-ole/go-ole) | 604 | 38 | 2011-01-21 | 5 months ago | Win32 OLE implementation for golang. |
| [d3d9](https://github.com/gonutz/d3d9) | 97 | 6 | 2015-12-12 | 4 months ago | Go bindings for Direct3D9. |
| [gosddl](https://github.com/MonaxGT/gosddl) | 2 | 1 | 2018-12-04 | 10 months ago | Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL. |

## XML
        
*Libraries and tools for manipulating XML.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [zek](https://github.com/miku/zek) | 332 | 21 | 2017-11-23 | 1 day ago | Generate a Go struct from XML. |
| [xpath](https://github.com/antchfx/xpath) | 288 | 9 | 2016-10-09 | 1 month ago | XPath package for Go. |
| [xquery](https://github.com/antchfx/xquery) | 154 | 11 | 2016-10-09 | 1 year ago | XQuery lets you extract data from HTML/XML documents using XPath expression. |
| [xml2map](https://github.com/sbabiv/xml2map) | 20 | 1 | 2018-08-06 | 2 weeks ago | XML to MAP converter written Golang. |
| [XML-Comp](https://github.com/XML-Comp/XML-Comp) | 16 | 1 | 2016-10-25 | 1 year ago | Simple command line XML comparer that generates diffs of folders, files and tags. |
| [xmlwriter](https://github.com/shabbyrobe/xmlwriter) | 9 | 2 | 2017-04-11 | 1 month ago | Procedural XML generation API based on libxml2's xmlwriter module. |

# Tools
        
*Go software and plugins.*

## Code Analysis
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [lint](https://github.com/golang/lint) | 3416 | 111 | 2013-06-02 | 1 week ago | Golint is a linter for Go source code. |
| [errcheck](https://github.com/kisielk/errcheck) | 1415 | 23 | 2013-02-24 | 2 weeks ago | Errcheck is a program for checking for unchecked errors in Go programs. |
| [gcvis](https://github.com/davecheney/gcvis) | 962 | 35 | 2014-07-10 | 1 year ago | Visualise Go program GC trace data in real time. |
| [php-parser](https://github.com/z7zmey/php-parser) | 704 | 28 | 2017-11-07 | 3 days ago | A Parser for PHP written in Go. |
| [go-critic](https://github.com/go-critic/go-critic) | 686 | 21 | 2018-05-05 | 3 weeks ago | source code linter that brings checks that are currently not implemented in other linters. |
| [goast-viewer](https://github.com/yuroyoro/goast-viewer) | 433 | 14 | 2014-06-30 | 9 months ago | Web based Golang AST visualizer. |
| [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) | 342 | 5 | 2019-04-19 | 2 weeks ago | An easy way to find outdated dependencies of your Go projects. |
| [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) | 317 | 9 | 2017-04-12 | 2 months ago | go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects. |
| [unconvert](https://github.com/mdempsky/unconvert) | 283 | 8 | 2016-02-19 | 2 weeks ago | Remove unnecessary type conversions from Go source. |
| [gostatus](https://github.com/shurcooL/gostatus) | 241 | 7 | 2013-11-27 | 1 year ago | Command line tool, shows the status of repositories that contain Go packages. |
| [dupl](https://github.com/mibk/dupl) | 193 | 8 | 2015-05-20 | 1 year ago | Tool for code clone detection. |
| [apicompat](https://github.com/bradleyfalzon/apicompat) | 172 | 8 | 2016-07-10 | 3 years ago | Checks recent changes to a Go project for backwards incompatible changes. |
| [tickgit](https://github.com/augmentable-dev/tickgit) | 154 | 7 | 2019-10-12 | 4 days ago | CLI and go package for surfacing code comment TODOs (in any language) and applying a `git blame`to identify the author. |
| [goplantuml](https://github.com/jfeliu007/goplantuml) | 105 | 6 | 2019-05-26 | 3 months ago | Library and CLI that generates text plantump class diagram containing information about structures and interfaces with the relationship among them. |
| [checkstyle](https://github.com/qiniu/checkstyle) | 104 | 11 | 2014-01-01 | 3 months ago | checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments. |
| [lint](https://github.com/surullabs/lint) | 64 | 5 | 2016-07-09 | 1 year ago | Run linters as part of go test. |
| [validate](https://github.com/mccoyst/validate) | 62 | 6 | 2013-11-22 | 4 years ago | Automatically validates struct fields with tags. |
| [go-outdated](https://github.com/firstrow/go-outdated) | 45 | 1 | 2015-06-29 | 1 year ago | Console application that displays outdated packages. |
| [golines](https://github.com/segmentio/golines) | 37 | 15 | 2019-10-01 | 1 week ago | Formatter that automatically shortens long lines in Go code. |
| [blanket](https://github.com/verygoodsoftwarenotvirus/blanket) | 15 | 2 | 2017-09-04 | 1 year ago | tarp finds functions and methods without direct unit tests in Go source code. |

## Editor Plugins
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [vim-go](https://github.com/fatih/vim-go) | 11798 | 295 | 2014-03-24 | 1 day ago | Go development plugin for Vim. |
| [vscode-go](https://github.com/microsoft/vscode-go) | 5730 | 227 | 2015-10-14 | 13 hours ago | Extension for Visual Studio Code (VS Code) which provides support for the Go language. |
| [gocode](https://github.com/nsf/gocode) | 4847 | 196 | 2010-07-05 | 1 hour ago | Autocompletion daemon for the Go programming language. |
| [GoSublime](https://github.com/DisposaBoy/GoSublime) | 3327 | 122 | 2011-08-27 | 1 week ago | Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features. |
| [go-plus](https://github.com/joefitzgerald/go-plus) | 1504 | 44 | 2014-03-13 | 1 hour ago | Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting. |
| [go-mode.el](https://github.com/dominikh/go-mode.el) | 1038 | 53 | 2013-01-30 | 1 week ago | Go mode for GNU/Emacs. |
| [Watch](https://github.com/eaburns/Watch) | 171 | 12 | 2013-08-08 | 2 years ago | Runs a command in an acme win on file changes. |
| [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) | 84 | 5 | 2012-11-25 | 3 years ago | Vim plugin to highlight syntax errors on save. |
| [go-language-server](https://github.com/theia-ide/go-language-server) | 32 | 5 | 2017-11-21 | 11 months ago | A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol. |
| [gounit-vim](https://github.com/hexdigest/gounit-vim) | 19 | 2 | 2018-02-21 | 1 year ago | Vim plugin for generating Go tests based on the function's or method's signature. |
| [theia-go-extension](https://github.com/theia-ide/theia-go-extension) | 13 | 4 | 2017-11-30 | 1 year ago | Go language support for the Theia IDE. |

## Go Generate Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gotests](https://github.com/cweill/gotests) | 2575 | 74 | 2016-01-19 | 2 weeks ago | Generate Go tests from your source code. |
| [genny](https://github.com/cheekybits/genny) | 1148 | 23 | 2014-10-27 | 5 days ago | Elegant generics for Go. |
| [re2dfa](https://github.com/opennota/re2dfa) | 178 | 9 | 2015-06-20 | 1 year ago | Transform regular expressions into finite state machines and output Go source code. |
| [gocontracts](https://github.com/Parquery/gocontracts) | 60 | 6 | 2018-08-13 | 1 year ago | brings design-by-contract to Go by synchronizing the code with the documentation. |
| [hasgo](https://github.com/DylanMeeus/hasgo) | 39 | 4 | 2019-05-16 | 2 weeks ago | Generate Haskell inspired functions for your slices. |
| [gounit](https://github.com/hexdigest/gounit) | 35 | 4 | 2018-02-05 | 1 year ago | Generate Go tests using your own templates. |
| [generic](https://github.com/usk81/generic) | 32 | 3 | 2016-06-15 | 11 months ago | flexible data type for Go. |

## Go Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-swagger](https://github.com/go-swagger/go-swagger) | 4884 | 120 | 2014-11-16 | 1 day ago | Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API. |
| [OctoLinker](https://github.com/OctoLinker/OctoLinker) | 4227 | 91 | 2013-12-27 | 4 hours ago | Navigate through go files efficiently with the OctoLinker browser extension for GitHub. |
| [go-callvis](https://github.com/ofabry/go-callvis) | 2425 | 71 | 2016-09-03 | 1 week ago | Visualize call graph of your Go program using dot format. |
| [go-callvis](https://github.com/TrueFurby/go-callvis) | 2408 | 70 | 2016-09-03 | 3 weeks ago | Visualize call graph of your Go program using dot format. |
| [depth](https://github.com/KyleBanks/depth) | 464 | 9 | 2017-03-04 | 1 month ago | Visualize dependency trees of any package by analyzing imports. |
| [richgo](https://github.com/kyoh86/richgo) | 436 | 5 | 2017-01-04 | 3 months ago | Enrich `go test` outputs with text decorations. |
| [rts](https://github.com/galeone/rts) | 195 | 3 | 2016-04-04 | 3 years ago | RTS: response to struct. Generates Go structs from server responses. |
| [godbg](https://github.com/tylerwince/godbg) | 168 | 5 | 2019-01-23 | 11 months ago | Implementation of Rusts `dbg!` macro for quick and easy debugging during development. |
| [colorgo](https://github.com/songgao/colorgo) | 103 | 6 | 2013-02-14 | 3 years ago | Wrapper around `go` command for colorized `go build` output. |
| [gothanks](https://github.com/psampaz/gothanks) | 91 | 3 | 2019-11-10 | 2 weeks ago | GoThanks automatically stars your go.mod github dependencies, sending this way some love to their maintainers. |
| [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) | 38 | 2 | 2015-05-22 | 2 years ago | Bash completion for go and wgo. |
| [generator-go-lang](https://github.com/axelspringer/generator-go-lang) | 17 | 12 | 2017-09-13 | 2 days ago | A [Yeoman](http://yeoman.io) generator to get new Go projects started. |
| [go-james](https://github.com/pieterclaerhout/go-james) | 11 | 2 | 2019-10-14 | 1 month ago | Go project skeleton creator, builds and tests your projects without the manual setup. |

## Software Packages
        
*Software written in Go.*

### DevOps Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [kubernetes](https://github.com/kubernetes/kubernetes) | 64102 | 3158 | 2014-06-06 | 32 minutes ago | Container Cluster Manager from Google. |
| [moby](https://github.com/moby/moby) | 56588 | 3185 | 2013-01-18 | 8 hours ago | Collaborative project for the container ecosystem to assemble container-based systems. |
| [traefik](https://github.com/containous/traefik) | 27862 | 702 | 2015-09-13 | 36 minutes ago | Reverse proxy and load balancer with support for multiple backends. |
| [gitea](https://github.com/go-gitea/gitea) | 18866 | 457 | 2016-11-01 | 28 minutes ago | Fork of Gogs, entirely community driven. |
| [vegeta](https://github.com/tsenart/vegeta) | 14086 | 292 | 2013-08-13 | 1 day ago | HTTP load testing tool and library. It's over 9000! |
| [packer](https://github.com/hashicorp/packer) | 9886 | 406 | 2013-03-23 | 20 minutes ago | Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. |
| [hey](https://github.com/rakyll/hey) | 7771 | 151 | 2016-09-02 | 3 days ago | Hey is a tiny program that sends some load to a web application. |
| [gvm](https://github.com/moovweb/gvm) | 5115 | 151 | 2011-12-03 | 4 weeks ago | GVM provides an interface to manage Go versions. |
| [webhook](https://github.com/adnanh/webhook) | 4997 | 142 | 2015-01-12 | 1 month ago | Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server. |
| [gaia](https://github.com/gaia-pipeline/gaia) | 4009 | 101 | 2017-12-28 | 5 days ago | Build powerful pipelines in any programming language. |
| [gox](https://github.com/mitchellh/gox) | 3646 | 73 | 2013-11-17 | 1 month ago | Dead simple, no frills Go cross compile tool. |
| [bosun](https://github.com/bosun-monitor/bosun) | 2955 | 157 | 2013-11-15 | 3 hours ago | Time Series Alerting Framework. |
| [bombardier](https://github.com/codesenberg/bombardier) | 2013 | 52 | 2016-05-29 | 2 weeks ago | Fast cross-platform HTTP benchmarking tool. |
| [fac](https://github.com/mkchoi212/fac) | 1651 | 28 | 2017-12-29 | 5 months ago | Command-line user interface to fix git merge conflicts. |
| [goxc](https://github.com/laher/goxc) | 1644 | 49 | 2013-02-11 | 5 months ago | build tool for Go, with a focus on cross-compiling and packaging. |
| [kala](https://github.com/ajvb/kala) | 1416 | 64 | 2015-03-19 | 1 month ago | Simplistic, modern, and performant job scheduler. |
| [script](https://github.com/bitfield/script) | 1297 | 23 | 2019-04-20 | 3 weeks ago | Making it easy to write shell-like scripts in Go for DevOps and system administration tasks. |
| [statusok](https://github.com/sanathp/statusok) | 1297 | 48 | 2015-08-26 | 1 month ago | Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected. |
| [pomerium](https://github.com/pomerium/pomerium) | 1052 | 14 | 2019-01-01 | 1 hour ago | Pomerium is an identity-aware access proxy. |
| [s3gof3r](https://github.com/rlmcpherson/s3gof3r) | 1044 | 34 | 2013-08-02 | 1 month ago | Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3. |
| [go-selfupdate](https://github.com/sanbornm/go-selfupdate) | 717 | 27 | 2013-11-13 | 3 days ago | Enable your Go applications to self update. |
| [skm](https://github.com/TimothyYe/skm) | 595 | 20 | 2017-10-11 | 6 months ago | SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily! |
| [scaleway-cli](https://github.com/scaleway/scaleway-cli) | 578 | 36 | 2015-03-20 | 3 minutes ago | Manage BareMetal Servers from Command Line (as easily as with Docker). |
| [aurora](https://github.com/xuri/aurora) | 439 | 27 | 2016-10-09 | 1 month ago | Cross-platform web-based Beanstalkd queue server console. |
| [govvv](https://github.com/ahmetb/govvv) | 428 | 10 | 2016-08-02 | 1 month ago | “go build” wrapper to easily add version information into Go binaries. |
| [gonative](https://github.com/inconshreveable/gonative) | 316 | 7 | 2014-05-01 | 3 years ago | Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages. |
| [mora](https://github.com/emicklei/mora) | 271 | 23 | 2013-07-12 | 3 years ago | REST server for accessing MongoDB documents and meta data. |
| [lstags](https://github.com/ivanilves/lstags) | 252 | 14 | 2017-08-15 | 1 week ago | Tool and API to sync Docker images across different registries. |
| [dogo](https://github.com/liudng/dogo) | 223 | 19 | 2014-11-19 | 1 year ago | Monitoring changes in the source file and automatically compile and run (restart). |
| [godbg](https://github.com/sirnewton01/godbg) | 220 | 18 | 2013-08-09 | 1 year ago | Web-based gdb front-end application. |
| [pewpew](https://github.com/bengadbois/pewpew) | 218 | 7 | 2016-10-12 | 7 months ago | Flexible HTTP command line stress tester. |
| [manssh](https://github.com/xwjdsh/manssh) | 214 | 3 | 2017-10-08 | 1 year ago | manssh is a command line tool for managing your ssh alias config easily. |
| [blast](https://github.com/dave/blast) | 183 | 4 | 2017-10-21 | 2 years ago | A simple tool for API load testing and batch jobs. |
| [gobrew](https://github.com/cryptojuice/gobrew) | 176 | 5 | 2013-11-13 | 2 years ago | gobrew lets you easily switch between multiple versions of go. |
| [ostent](https://github.com/ostrost/ostent) | 166 | 6 | 2014-03-31 | 1 year ago | collects and displays system metrics and optionally relays to Graphite and/or InfluxDB. |
| [utask](https://github.com/ovh/utask) | 166 | 22 | 2019-11-05 | 25 minutes ago | Automation engine that models and executes business processes declared in yaml. |
| [grapes](https://github.com/yaronsumel/grapes) | 146 | 6 | 2016-09-01 | 6 months ago | Lightweight tool designed to distribute commands over ssh with ease. |
| [jenkins-cli](https://github.com/jenkins-zh/jenkins-cli) | 137 | 6 | 2019-06-21 | 17 hours ago | Jenkins CLI allows you manage your Jenkins as an easy way. |
| [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) | 134 | 6 | 2017-03-03 | 1 month ago | Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`. |
| [kcli](https://github.com/cswank/kcli) | 107 | 6 | 2017-03-25 | 2 months ago | Command line tool for inspecting kafka topics/partitions/messages. |
| [terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi) | 99 | 9 | 2017-10-17 | 2 months ago | Terraform provider plugin that dynamically configures itself at runtime based on an OpenAPI document (formerly known as swagger file) containing the definitions of the APIs exposed. |
| [winrm-cli](https://github.com/masterzen/winrm-cli) | 85 | 6 | 2016-05-23 | 1 month ago | Cli tool to remotely execute commands on Windows machines. |
| [go-furnace](https://github.com/go-furnace/go-furnace) | 75 | 1 | 2016-10-09 | 5 months ago | Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean. |
| [drone-scp](https://github.com/appleboy/drone-scp) | 71 | 3 | 2016-10-16 | 2 weeks ago | Copy files and artifacts via SSH using a binary, docker or Drone CI. |
| [dockerfile-generator](https://github.com/ozankasikci/dockerfile-generator) | 69 | 4 | 2019-08-14 | 2 months ago | A go library and an executable that produces valid Dockerfiles using various input channels. |
| [dropship](https://github.com/ChrisMcKenzie/dropship) | 50 | 3 | 2015-09-03 | 1 year ago | Tool for deploying code via cdn. |
| [rodent](https://github.com/alouche/rodent) | 31 | 2 | 2014-06-01 | 2 years ago | Rodent helps you manage Go versions, projects and track dependencies. |
| [drone-jenkins](https://github.com/appleboy/drone-jenkins) | 27 | 1 | 2016-10-15 | 1 month ago | Trigger downstream Jenkins jobs using a binary, docker or Drone CI. |
| [awsenv](https://github.com/soniah/awsenv) | 23 | 2 | 2015-08-05 | 1 year ago | Small binary that loads Amazon (AWS) environment variables for a profile. |
| [lwc](https://github.com/timdp/lwc) | 21 | 2 | 2018-04-22 | 1 year ago | A live-updating version of the UNIX wc command. |
| [depcharge](https://github.com/centerorbit/depcharge) | 12 | 3 | 2018-07-25 | 1 month ago | Helps orchestrating the execution of commands across the many dependencies in larger projects. |
| [s3-proxy](https://github.com/oxyno-zeta/s3-proxy) | 6 | 2 | 2019-09-22 | 18 hours ago | S3 Proxy with GET, PUT and DELETE methods and authentication (OpenID Connect and Basic Auth). |
| [sg](https://github.com/ChristopherRabotin/sg) | 5 | 1 | 2015-08-19 | 3 years ago | Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response. |
| [aptly-fork](https://github.com/smira/aptly-fork) | 2 | 0 | 2019-07-04 | 5 months ago | aptly is a Debian repository management tool. |

### Other Software
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [goreplay](https://github.com/buger/goreplay) | 12401 | 464 | 2013-05-30 | 5 days ago | Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time. |
| [restic](https://github.com/restic/restic) | 9634 | 233 | 2014-04-27 | 2 hours ago | De-duplicating backup program. |
| [seaweedfs](https://github.com/chrislusf/seaweedfs) | 9292 | 517 | 2014-07-14 | 10 hours ago | Fast, Simple and Scalable Distributed File System with O(1) disk seek. |
| [confd](https://github.com/kelseyhightower/confd) | 6868 | 256 | 2013-10-01 | 57 minutes ago | Manage local application configuration files using templates and data from etcd or consul. |
| [comcast](https://github.com/tylertreat/comcast) | 6496 | 150 | 2014-11-12 | 8 months ago | Simulate bad network connections. |
| [liteide](https://github.com/visualfc/liteide) | 5905 | 379 | 2012-11-19 | 3 days ago | LiteIDE is a simple, open source, cross-platform Go IDE. |
| [drive](https://github.com/odeke-em/drive) | 5326 | 204 | 2014-11-03 | 1 month ago | Google Drive client for the commandline. |
| [toxiproxy](https://github.com/Shopify/toxiproxy) | 4380 | 288 | 2014-09-04 | 4 weeks ago | Proxy to simulate network and system conditions for automated tests. |
| [nes](https://github.com/fogleman/nes) | 4377 | 149 | 2015-03-02 | 9 months ago | Nintendo Entertainment System (NES) emulator written in Go. |
| [duplicacy](https://github.com/gilbertchen/duplicacy) | 3199 | 95 | 2016-02-23 | 2 hours ago | A cross-platform network and cloud backup tool based on the idea of lock-free deduplication. |
| [croc](https://github.com/schollz/croc) | 2878 | 59 | 2017-10-17 | 1 week ago | Easily and securely send files or folders from one computer to another. |
| [mylg](https://github.com/mehrdadrad/mylg) | 2286 | 107 | 2016-06-21 | 2 weeks ago | Command Line Network Diagnostic tool written in Go. |
| [goboy](https://github.com/Humpheh/goboy) | 2199 | 40 | 2017-08-20 | 1 month ago | Nintendo Game Boy Color emulator written in Go. |
| [sup](https://github.com/pressly/sup) | 2115 | 73 | 2015-02-23 | 3 weeks ago | Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers. |
| [lgo](https://github.com/yunabe/lgo) | 1988 | 47 | 2017-10-05 | 8 months ago | Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility. |
| [circuit](https://github.com/gocircuit/circuit) | 1817 | 143 | 2014-04-10 | 10 months ago | Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications. |
| [snap](https://github.com/intelsdi-x/snap) | 1803 | 145 | 2014-08-13 | 1 year ago | Powerful telemetry framework. |
| [borg](https://github.com/ok-borg/borg) | 1464 | 42 | 2016-09-10 | 2 years ago | Terminal based search engine for bash snippets. |
| [scc](https://github.com/boyter/scc) | 1366 | 19 | 2018-03-01 | 1 week ago | Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates. |
| [community](https://github.com/documize/community) | 1030 | 41 | 2016-04-29 | 1 week ago | Modern wiki software that integrates data from SaaS tools. |
| [Go-Package-Store](https://github.com/shurcooL/Go-Package-Store) | 888 | 20 | 2014-01-24 | 1 week ago | App that displays updates for the Go packages in your GOPATH. |
| [peg](https://github.com/pointlander/peg) | 681 | 32 | 2010-04-25 | 1 month ago | Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator. |
| [leaps](https://github.com/Jeffail/leaps) | 668 | 27 | 2014-06-19 | 1 year ago | Pair programming service using Operational Transforms. |
| [vflow](https://github.com/VerizonDigital/vflow) | 659 | 86 | 2017-02-24 | 4 months ago | High-performance, scalable and reliable IPFIX, sFlow and Netflow collector. |
| [gfile](https://github.com/Antonito/gfile) | 535 | 11 | 2019-03-08 | 11 months ago | Securely transfer files between two computers, without any third party, over WebRTC. |
| [shell2http](https://github.com/msoap/shell2http) | 524 | 20 | 2015-03-11 | 4 months ago | Executing shell commands via http server (for prototyping or remote control). |
| [mockingjay-server](https://github.com/quii/mockingjay-server) | 443 | 10 | 2015-04-04 | 1 week ago | Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests. |
| [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) | 397 | 18 | 2015-10-08 | 2 months ago | Video streaming torrent client. |
| [gocc](https://github.com/goccmack/gocc) | 391 | 20 | 2015-06-05 | 4 months ago | Gocc is a compiler kit for Go written in Go. |
| [ipe](https://github.com/dimiro1/ipe) | 299 | 17 | 2015-01-13 | 3 days ago | Open source Pusher server implementation compatible with Pusher client libraries written in GO. |
| [wellington](https://github.com/wellington/wellington) | 290 | 13 | 2014-12-08 | 1 year ago | Sass project management tool, extends the language with sprite functions (like Compass). |
| [IDE](https://github.com/thestrukture/IDE) | 255 | 14 | 2017-09-09 | 7 months ago | Browser accessible IDE. Designed for Go with Go. |
| [cherry](https://github.com/rafael-santiago/cherry) | 214 | 12 | 2015-10-24 | 2 years ago | Tiny webchat server in Go. |
| [orange-cat](https://github.com/utatti/orange-cat) | 184 | 6 | 2014-11-01 | 1 year ago | Markdown previewer written in Go. |
| [orbit](https://github.com/gulien/orbit) | 135 | 8 | 2017-05-13 | 3 months ago | A simple tool for running commands and generating files from templates. |
| [joincap](https://github.com/assafmo/joincap) | 131 | 7 | 2018-05-31 | 2 weeks ago | Command-line utility for merging multiple pcap files together. |
| [boxed](https://github.com/tejo/boxed) | 72 | 2 | 2015-04-18 | 1 year ago | Dropbox based blog engine. |
| [dp](https://github.com/scryinfo/dp) | 57 | 9 | 2018-12-12 | 1 day ago | Through SDK for data exchange with blockchain, developers can get easy access to DAPP development. |
| [naclpipe](https://github.com/unix4fun/naclpipe) | 20 | 5 | 2015-05-05 | 1 year ago | Simple NaCL EC25519 based crypto pipe tool written in Go. |
| [term-quiz](https://github.com/crazcalm/term-quiz) | 17 | 1 | 2017-12-26 | 1 year ago | Quizzes for your terminal. |
| [snitch](https://github.com/lucasgomide/snitch) | 15 | 1 | 2017-04-06 | 1 year ago | Simple way to notify your team and many tools when someone has deployed any application via Tsuru. |
| [GoDocTooltip](https://github.com/diankong/GoDocTooltip) | 11 | 3 | 2016-01-21 | 4 years ago | Chrome extension for Go Doc sites, which shows function description as tooltip at function list. |

# Resources
        
*Where to discover new Go libraries.*

## Benchmarks
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) | 1356 | 61 | 2013-12-16 | 3 months ago | Go HTTP request router benchmark and comparison. |
| [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) | 1142 | 85 | 2016-04-06 | 1 week ago | Go web framework benchmark. |
| [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) | 956 | 39 | 2013-01-18 | 2 months ago | Benchmarks of Go serialization methods. |
| [skynet](https://github.com/atemerev/skynet) | 946 | 46 | 2016-02-14 | 10 months ago | Skynet 1M threads microbenchmark. |
| [speedtest-resize](https://github.com/fawick/speedtest-resize) | 189 | 7 | 2013-09-16 | 4 months ago | Compare various Image resize algorithms for the Go language. |
| [go-benchmarks](https://github.com/tylertreat/go-benchmarks) | 132 | 9 | 2016-02-25 | 4 years ago | Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches. |
| [gospeed](https://github.com/feyeleanor/gospeed) | 101 | 7 | 2011-05-23 | 1 year ago | Go micro-benchmarks for calculating the speed of language constructs. |
| [autobench](https://github.com/davecheney/autobench) | 90 | 8 | 2013-03-28 | 5 years ago | Framework to compare the performance between different Go versions. |
| [gocostmodel](https://github.com/mna/gocostmodel) | 55 | 5 | 2014-12-19 | 5 years ago | Benchmarks of common basic operations for the Go language. |
| [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) | 52 | 5 | 2014-09-24 | 2 years ago | Collection of benchmarks for popular Go database/SQL utilities. |
| [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) | 20 | 1 | 2017-01-24 | 3 years ago | Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results. |
| [kvbench](https://github.com/jimrobinson/kvbench) | 16 | 1 | 2014-04-15 | 5 months ago | Key/Value database benchmark. |

## Conferences
        

## E-Books
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [GoBooks](https://github.com/dariubs/GoBooks) | 7618 | 536 | 2015-05-05 | 1 month ago | A curated list of Go books. |
| [The-Golang-Standard-Library-by-Example](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example) | 6483 | 551 | 2013-04-14 | 1 month ago | Golang标准库。对于程序员而言，标准库与语言本身同样重要，它好比一个百宝箱，能为各种常见的任务提供完美的解决方案。以示例驱动的方式讲解Golang的标准库。 |
| [gosuccinctly](https://github.com/thedevsir/gosuccinctly) | 13 | 3 | 2018-09-02 | 1 year ago | in Persian. |

## Gophers
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gophers](https://github.com/ashleymcnamara/gophers) | 2087 | 94 | 2017-02-15 | 11 months ago | Gopher artworks by Ashley McNamara. |
| [gophers](https://github.com/egonelbre/gophers) | 1895 | 45 | 2015-06-03 | 4 weeks ago | Free gophers. |
| [free-gophers-pack](https://github.com/MariaLetta/free-gophers-pack) | 1823 | 45 | 2019-04-02 | 1 month ago | Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster. |
| [gophericons](https://github.com/shalakhin/gophericons) | 565 | 20 | 2015-08-22 | 2 years ago | 34 gopher images for Go developers community |
| [gopher-stickers](https://github.com/tenntenn/gopher-stickers) | 467 | 15 | 2014-11-09 | 3 months ago | gopher stickers |
| [gopherize.me](https://github.com/matryer/gopherize.me) | 370 | 7 | 2017-01-25 | 1 year ago | Gopherize yourself. |
| [gopher-vector](https://github.com/golang-samples/gopher-vector) | 356 | 13 | 2013-03-31 | 3 years ago | Vector data of gopher |
| [gopher-logos](https://github.com/GolangUA/gopher-logos) | 74 | 7 | 2017-07-27 | 1 year ago | adorable gopher logos. |
| [gophers](https://github.com/rogeralsing/gophers) | 52 | 2 | 2017-01-28 | 3 years ago | random gopher graphics. |
| [go-gopher](https://github.com/sillecelik/go-gopher) | 44 | 1 | 2018-03-28 | 5 months ago | Gopher amigurumi toy pattern. |
| [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) | 33 | 1 | 2014-09-03 | 2 years ago | Go gopher Vector Data [.ai, .svg]. |

## Meetups
        
*Add the group of your city/country here (send **PR**)*

## Twitter
        

## Websites
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) | 25732 | 1733 | 2014-07-08 | 4 hours ago | List of other amazingly awesome lists. |
| [awesome-remote-job](https://github.com/lukasz-madon/awesome-remote-job) | 17128 | 919 | 2015-01-02 | 10 minutes ago | Curated list of awesome remote jobs. A lot of them are looking for Go hackers. |
| [golang-graphics](https://github.com/mholt/golang-graphics) | 142 | 9 | 2014-03-24 | 4 years ago | Collection of Go images, graphics, and art. |
| [gocryforhelp](https://github.com/ninedraft/gocryforhelp) | 38 | 12 | 2016-05-09 | 2 years ago | Collection of Go projects that needs help. Good place to start your open-source way in Go. |

### Tutorials
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang) | 34021 | 2462 | 2012-08-02 | 3 weeks ago | Golang ebook intro how to build a web app with golang. |
| [go-patterns](https://github.com/tmrts/go-patterns) | 12262 | 567 | 2015-12-14 | 1 month ago | Curated list of Go design patterns, recipes and idioms. |
| [learn-go-with-tests](https://github.com/quii/learn-go-with-tests) | 10313 | 253 | 2018-03-02 | 1 week ago | Learn Go with test-driven development. |
| [golang-cheat-sheet](https://github.com/a8m/golang-cheat-sheet) | 4432 | 177 | 2014-02-13 | 7 months ago | Go's reference card. |
| [golang-for-nodejs-developers](https://github.com/miguelmota/golang-for-nodejs-developers) | 1196 | 30 | 2019-01-03 | 2 months ago | Examples of Golang compared to Node.js for learning. |
| [working-with-go](https://github.com/mkaz/working-with-go) | 1155 | 50 | 2014-05-04 | 1 month ago | Intro to go for experienced programmers. |
| [ethereum-development-with-go-book](https://github.com/miguelmota/ethereum-development-with-go-book) | 556 | 37 | 2018-05-16 | 2 months ago | A little e-book on Ethereum Development with Go. |

> 该项目源码[Awesome Go Analysis](https://github.com/plholx/awesome-go-analysis)
> 更专业的go开源项目分析请移步 [Awesome Go](https://go.libhunt.com/)
