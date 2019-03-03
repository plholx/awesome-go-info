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
    - [Text Processing](#text-processing)
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
| [EasyMIDI](https://github.com/algoGuy/EasyMIDI) | 17 | 3 | 2018-02-20 | 11 months ago | EasyMidi is a simple and reliable library for working with standard midi file (SMF). |
| [flac](https://github.com/eaburns/flac) | 81 | 5 | 2013-08-21 | 1 year ago | No-frills native Go FLAC decoder that decodes FLAC files into byte slices. |
| [flac](https://github.com/mewkiz/flac) | 92 | 6 | 2012-11-02 | 1 week ago | Native Go FLAC encoder/decoder with support for FLAC streams. |
| [gaad](https://github.com/Comcast/gaad) | 50 | 10 | 2016-07-11 | 1 year ago | Native Go AAC bitstream parser. |
| [go-sox](https://github.com/krig/go-sox) | 87 | 8 | 2013-10-08 | 8 months ago | libsox bindings for go. |
| [go_mediainfo](https://github.com/zhulik/go_mediainfo) | 22 | 1 | 2015-12-14 | 3 years ago | libmediainfo bindings for go. |
| [gosamplerate](https://github.com/dh1tw/gosamplerate) | 8 | 1 | 2016-11-21 | 9 months ago | libsamplerate bindings for go. |
| [id3v2](https://github.com/bogem/id3v2) | 90 | 5 | 2016-05-16 | 3 weeks ago | Fast and stable ID3 parsing and writing library for Go. |
| [malgo](https://github.com/gen2brain/malgo) | 55 | 4 | 2017-11-10 | 4 weeks ago | Mini audio library. |
| [minimp3](https://github.com/tosone/minimp3) | 19 | 2 | 2018-01-26 | 1 day ago | Lightweight MP3 decoder library. |
| [mix](https://github.com/go-mix/mix) | 90 | 3 | 2016-01-03 | 1 year ago | Sequence-based Go-native audio mixer for music apps. |
| [mp3](https://github.com/tcolgate/mp3) | 85 | 1 | 2015-02-27 | 1 year ago | Native Go MP3 decoder. |
| [music-theory](https://github.com/go-music-theory/music-theory) | 233 | 10 | 2016-03-17 | 8 months ago | Music theory models in Go. |
| [portaudio](https://github.com/gordonklaus/portaudio) | 262 | 9 | 2015-09-16 | 6 months ago | Go bindings for the PortAudio audio I/O library. |
| [portmidi](https://github.com/rakyll/portmidi) | 187 | 6 | 2013-11-10 | 1 year ago | Go bindings for PortMidi. |
| [go-taglib](https://github.com/wtolson/go-taglib) | 64 | 6 | 2012-11-20 | 7 months ago | Go bindings for taglib. |
| [vorbis](https://github.com/mccoyst/vorbis) | 21 | 3 | 2013-07-12 | 1 year ago | "Native" Go Vorbis decoder (uses CGO, but has no dependencies). |
| [waveform](https://github.com/mdlayher/waveform) | 228 | 13 | 2014-09-14 | 2 years ago | Go package capable of generating waveform images from audio streams. |

## Authentication and OAuth
        
*Libraries for implementing authentications schemes.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [authboss](https://github.com/volatiletech/authboss) | 1737 | 38 | 2015-01-03 | 3 weeks ago | Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time. |
| [branca](https://github.com/hako/branca) | 50 | 5 | 2018-01-09 | 6 months ago | Golang implementation of Branca Tokens. |
| [casbin](https://github.com/casbin/casbin) | 3854 | 140 | 2017-04-08 | 2 weeks ago | Authorization library that supports access control models like ACL, RBAC, ABAC. |
| [cookiestxt](https://github.com/mengzhuo/cookiestxt) | 2 | 1 | 2017-10-09 | 1 year ago | provides parser of cookies.txt file format. |
| [go-jose](https://github.com/square/go-jose) | 1015 | 61 | 2014-11-15 | 4 days ago | Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs. |
| [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) | 1115 | 71 | 2015-11-01 | 1 week ago | Standalone, specification-compliant,  OAuth2 server written in Golang. |
| [gologin](https://github.com/dghubble/gologin) | 960 | 27 | 2015-06-23 | 2 weeks ago | chainable handlers for login with OAuth1 and OAuth2 authentication providers. |
| [gorbac](https://github.com/mikespook/gorbac) | 828 | 55 | 2013-12-26 | 1 month ago | provides a lightweight role-based access control (RBAC) implementation in Golang. |
| [goth](https://github.com/markbates/goth) | 2081 | 58 | 2014-10-15 | 1 week ago | provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box. |
| [httpauth](https://github.com/goji/httpauth) | 167 | 5 | 2014-05-27 | 2 years ago | HTTP Authentication middleware. |
| [jwt](https://github.com/robbert229/jwt) | 60 | 6 | 2016-06-06 | 3 months ago | Clean and easy to use implementation of JSON Web Tokens (JWT). |
| [jwt](https://github.com/pascaldekloe/jwt) | 38 | 2 | 2018-03-21 | 1 month ago | Lightweight JSON Web Token (JWT) library. |
| [jwt-auth](https://github.com/adam-hanna/jwt-auth) | 144 | 9 | 2016-07-06 | 1 month ago | JWT middleware for Golang http servers with many configuration options. |
| [jwt-go](https://github.com/dgrijalva/jwt-go) | 5014 | 130 | 2012-04-18 | 1 week ago | Golang implementation of JSON Web Tokens (JWT). |
| [loginsrv](https://github.com/tarent/loginsrv) | 734 | 45 | 2016-11-11 | 1 day ago | JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam. |
| [oauth2](https://github.com/golang/oauth2) | 2109 | 92 | 2014-04-14 | 4 days ago | Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support. |
| [osin](https://github.com/openshift/osin) | 1507 | 69 | 2013-09-11 | 4 months ago | Golang OAuth2 server library. |
| [paseto](https://github.com/o1egl/paseto) | 194 | 9 | 2018-01-23 | 3 months ago | Golang implementation of Platform-Agnostic Security Tokens (PASETO). |
| [permissions2](https://github.com/xyproto/permissions2) | 298 | 12 | 2014-11-19 | 1 week ago | Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt. |
| [rbac](https://github.com/zpatrick/rbac) | 20 | 3 | 2018-08-02 | 6 months ago | Minimalistic RBAC package for Go applications. |
| [securecookie](https://github.com/chmike/securecookie) | 26 | 4 | 2017-09-03 | 6 months ago | Efficient secure cookie encoding/decoding. |
| [session](https://github.com/icza/session) | 82 | 6 | 2016-02-08 | 5 months ago | Go session management for web servers (including support for Google App Engine - GAE). |
| [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) | 6 | 2 | 2017-10-20 | 3 months ago | Go session management using the SessionGate Redis module. |
| [sessions](https://github.com/adam-hanna/sessions) | 41 | 3 | 2017-04-29 | 1 year ago | Dead simple, highly performant, highly customizable sessions service for go http servers. |
| [signedvalue](https://github.com/sashka/signedvalue) | 6 | 0 | 2018-01-06 | 1 month ago | Signed and timestamped strings compatible with [Tornado's](https://github.com/tornadoweb/tornado) `create_signed_value`, `decode_signed_value`, and therefore `set_secure_cookie` and `get_secure_cookie`. |

## Bot Building
        
*Libraries for building and working with bots.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [bot](https://github.com/go-chat-bot/bot) | 412 | 34 | 2015-09-23 | 1 week ago | IRC, Slack & Telegram bot written in Go. |
| [go-sarah](https://github.com/oklahomer/go-sarah) | 111 | 3 | 2016-11-06 | 7 months ago | Framework to build bot for desired chat services including LINE, Slack, Gitter and more. |
| [go-tgbot](https://github.com/olebedev/go-tgbot) | 80 | 7 | 2016-12-11 | 8 months ago | Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware. |
| [golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot) | 147 | 19 | 2017-05-15 | 2 weeks ago | A golang implementation of a console-based trading bot for cryptocurrency exchanges. |
| [govkbot](https://github.com/nikepan/govkbot) | 19 | 2 | 2016-07-12 | 2 hours ago | Simple Go [VK](https://vk.com) bot library. |
| [hanu](https://github.com/sbstjn/hanu) | 90 | 6 | 2016-09-16 | 5 months ago | Framework for writing Slack bots. |
| [margelet](https://github.com/zhulik/margelet) | 51 | 5 | 2015-11-21 | 2 years ago | Framework for building Telegram bots. |
| [micha](https://github.com/onrik/micha) | 9 | 3 | 2016-04-14 | 1 year ago | Go Library for Telegram bot api. |
| [slacker](https://github.com/shomali11/slacker) | 256 | 13 | 2017-05-20 | 1 week ago | Easy to use framework to create Slack bots. |
| [tbot](https://github.com/yanzay/tbot) | 184 | 8 | 2015-09-12 | 2 months ago | Telegram bot server with API similar to net/http. |
| [telebot](https://github.com/tucnak/telebot) | 825 | 34 | 2015-06-26 | 7 hours ago | Telegram bot framework written in Go. |
| [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) | 1385 | 56 | 2015-06-25 | 1 day ago | Simple and clean Telegram bot client. |
| [tenyks](https://github.com/kyleterry/tenyks) | 166 | 14 | 2012-08-26 | 2 years ago | Service oriented IRC bot using Redis and JSON for messaging. |

## Command Line
        
*Libraries for building Console Applications and Console User Interfaces.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [argparse](https://github.com/akamensky/argparse) | 74 | 5 | 2017-11-24 | 1 month ago | Command line argument parser inspired by Python's argparse module. |
| [argv](https://github.com/cosiner/argv) | 11 | 1 | 2017-01-22 | 1 month ago | Go library to split command line string as arguments array using the bash syntax. |
| [cli](https://github.com/mkideal/cli) | 426 | 20 | 2016-02-27 | 3 days ago | Feature-rich and easy to use command-line package based on golang struct tags. |
| [cli](https://github.com/teris-io/cli) | 43 | 1 | 2017-05-25 | 8 months ago | Simple and complete API for building command line interfaces in Go. |
| [gcli](https://github.com/tcnksm/gcli) | 856 | 25 | 2014-06-19 | 1 year ago | The easy way to start building Golang command line applications. |
| [cobra](https://github.com/spf13/cobra) | 10761 | 255 | 2013-09-04 | 1 week ago | Commander for modern Go CLI interactions. |
| [commandeer](https://github.com/jaffee/commandeer) | 70 | 4 | 2017-10-12 | 1 day ago | Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags. |
| [complete](https://github.com/posener/complete) | 541 | 12 | 2017-05-06 | 1 week ago | Write bash completions in Go + Go command bash completion. |
| [docopt.go](https://github.com/docopt/docopt.go) | 1036 | 36 | 2013-08-26 | 5 months ago | Command-line arguments parser that will make you smile. |
| [env](https://github.com/codingconcepts/env) | 37 | 1 | 2017-06-15 | 4 months ago | Tag-based environment configuration for structs. |
| [flag](https://github.com/cosiner/flag) | 90 | 5 | 2016-10-06 | 1 week ago | Simple but powerful command line option parsing library for Go supporting subcommand. |
| [flaggy](https://github.com/integrii/flaggy) | 410 | 11 | 2018-03-05 | 1 week ago | A robust and idiomatic flags package with excellent subcommand support. |
| [flagvar](https://github.com/sgreben/flagvar) | 26 | 1 | 2018-05-19 | 4 months ago | A collection of flag argument types for Go's standard `flag` package. |
| [go-arg](https://github.com/alexflint/go-arg) | 571 | 15 | 2015-11-01 | 2 months ago | Struct-based argument parsing in Go. |
| [go-commander](https://github.com/yitsushi/go-commander) | 8 | 1 | 2016-10-10 | 5 days ago | Go library to simplify CLI workflow. |
| [go-flags](https://github.com/jessevdk/go-flags) | 1290 | 28 | 2012-08-31 | 6 days ago | go command line option parser. |
| [go-getoptions](https://github.com/DavidGamba/go-getoptions) | 5 | 1 | 2015-12-18 | 15 hours ago | Go option parser inspired on the flexibility of Perl’s GetOpt::Long. |
| [gocmd](https://github.com/devfacet/gocmd) | 28 | 3 | 2018-01-08 | 5 months ago | Go library for building command line applications. |
| [hiboot](https://github.com/hidevopsio/hiboot) | 59 | 6 | 2018-03-16 | 5 days ago | cli application framework with auto configuration and dependency injection. |
| [kingpin](https://github.com/alecthomas/kingpin) | 2297 | 54 | 2014-05-15 | 1 month ago | Command line and flag parser supporting sub commands. |
| [liner](https://github.com/peterh/liner) | 521 | 22 | 2012-08-16 | 2 weeks ago | Go readline-like library for command-line interfaces. |
| [cli](https://github.com/mitchellh/cli) | 925 | 25 | 2013-11-03 | 3 months ago | Go library for implementing command-line interfaces. |
| [mow.cli](https://github.com/jawher/mow.cli) | 590 | 19 | 2014-12-19 | 5 days ago | Go library for building CLI applications with sophisticated flag and argument parsing and validation. |
| [pflag](https://github.com/spf13/pflag) | 624 | 24 | 2013-08-30 | 2 weeks ago | Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. |
| [readline](https://github.com/chzyer/readline) | 1306 | 39 | 2015-09-20 | 3 weeks ago | Pure golang implementation that provides most features in GNU-Readline under MIT license. |
| [sand](https://github.com/Zaba505/sand) | 2 | 1 | 2018-11-19 | 3 months ago | Simple API for creating interpreters and so much more. |
| [sflags](https://github.com/octago/sflags) | 71 | 5 | 2016-12-04 | 4 weeks ago | Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries. |
| [strumt](https://github.com/antham/strumt) | 22 | 0 | 2017-06-20 | 1 day ago | Library to create prompt chain. |
| [clif](https://github.com/ukautz/clif) | 95 | 2 | 2015-05-31 | 1 week ago | Small command line interface framework. |
| [cli](https://github.com/urfave/cli) | 10074 | 271 | 2013-07-14 | 3 weeks ago | Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli). |
| [wlog](https://github.com/dixonwille/wlog) | 29 | 1 | 2016-04-14 | 1 year ago | Simple logging interface that supports cross-platform color and concurrency. |
| [wmenu](https://github.com/dixonwille/wmenu) | 67 | 2 | 2016-04-20 | 1 week ago | Easy to use menu structure for cli applications that prompts users to make choices. |
| [asciigraph](https://github.com/guptarohit/asciigraph) | 1019 | 23 | 2018-06-17 | 1 month ago | Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. |
| [aurora](https://github.com/logrusorgru/aurora) | 428 | 5 | 2016-11-07 | 5 months ago | ANSI terminal colors that supports fmt.Printf/Sprintf. |
| [cfmt](https://github.com/mingrammer/cfmt) | 53 | 3 | 2018-03-16 | 2 months ago | Contextual fmt inspired by bootstrap color classes. |
| [chalk](https://github.com/ttacon/chalk) | 286 | 11 | 2014-07-19 | 2 years ago | Intuitive package for prettifying terminal/console output. |
| [color](https://github.com/fatih/color) | 2863 | 56 | 2014-02-17 | 4 months ago | Versatile package for colored terminal output. |
| [colourize](https://github.com/TreyBastian/colourize) | 15 | 3 | 2015-05-11 | 2 years ago | Go library for ANSI colour text in terminals. |
| [ctc](https://github.com/wzshiming/ctc) | 6 | 1 | 2018-04-28 | 4 months ago | The non-invasive cross-platform terminal color library does not need to modify the Print method. |
| [go-ataman](https://github.com/workanator/go-ataman) | 8 | 2 | 2017-05-18 | 1 year ago | Go library for rendering ANSI colored text templates in terminals. |
| [go-colorable](https://github.com/mattn/go-colorable) | 334 | 16 | 2014-07-30 | 1 week ago | Colorable writer for windows. |
| [go-colortext](https://github.com/daviddengcn/go-colortext) | 188 | 9 | 2013-01-23 | 10 months ago | Go library for color output in terminals. |
| [go-isatty](https://github.com/mattn/go-isatty) | 298 | 7 | 2014-04-01 | 5 days ago | isatty for golang. |
| [go-prompt](https://github.com/c-bata/go-prompt) | 2081 | 35 | 2017-08-15 | 1 week ago | Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit). |
| [gocui](https://github.com/jroimartin/gocui) | 3966 | 102 | 2014-01-04 | 1 week ago | Minimalist Go library aimed at creating Console User Interfaces. |
| [gommon](https://github.com/labstack/gommon) | 269 | 19 | 2015-03-13 | 1 month ago | Style terminal text. |
| [color](https://github.com/gookit/color) | 114 | 3 | 2018-07-01 | 1 week ago | Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows. |
| [mpb](https://github.com/vbauerster/mpb) | 419 | 8 | 2016-12-14 | 5 days ago | Multi progress bar for terminal applications. |
| [progressbar](https://github.com/schollz/progressbar) | 493 | 13 | 2017-10-27 | 1 month ago | Basic thread-safe progress bar that works in every OS. |
| [simpletable](https://github.com/alexeyco/simpletable) | 121 | 2 | 2017-03-29 | 1 week ago | Simple tables in terminal with Go. |
| [tabular](https://github.com/InVisionApp/tabular) | 23 | 3 | 2018-04-24 | 9 months ago | Print ASCII tables from command line utilities without the need to pass large sets of data to the API. |
| [termbox-go](https://github.com/nsf/termbox-go) | 3289 | 103 | 2012-01-13 | 2 weeks ago | Termbox is a library for creating cross-platform text-based interfaces. |
| [termdash](https://github.com/mum4k/termdash) | 81 | 6 | 2018-03-24 | 4 hours ago | Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui). |
| [termtables](https://github.com/apcera/termtables) | 203 | 68 | 2012-12-06 | 1 year ago | Go port of the Ruby library [terminal-tables](https://github.com/tj/terminal-table) for simple ASCII table generation as well as providing markdown and HTML output. |
| [termui](https://github.com/gizak/termui) | 8358 | 275 | 2015-02-03 | 1 day ago | Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib). |
| [uilive](https://github.com/gosuri/uilive) | 723 | 11 | 2015-11-16 | 2 months ago | Library for updating terminal output in realtime. |
| [uiprogress](https://github.com/gosuri/uiprogress) | 1228 | 27 | 2015-11-17 | 4 months ago | Flexible library to render progress bars in terminal applications. |
| [uitable](https://github.com/gosuri/uitable) | 464 | 15 | 2015-11-14 | 1 year ago | Library to improve readability in terminal apps using tabular data. |

## Configuration
        
*Libraries for configuration parsing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [config](https://github.com/olebedev/config) | 190 | 8 | 2014-04-21 | 5 months ago | JSON or YAML configuration wrapper with environment variables and flags parsing. |
| [configure](https://github.com/paked/configure) | 43 | 3 | 2015-06-14 | 1 week ago | Provides configuration through multiple sources, including JSON, flags and environment variables. |
| [confita](https://github.com/heetch/confita) | 208 | 25 | 2017-12-21 | 2 weeks ago | Load configuration in cascade from multiple backends into a struct. |
| [conflate](https://github.com/the4thamigo-uk/conflate) | 3 | 0 | 2018-02-02 | 3 weeks ago | Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema. |
| [env](https://github.com/caarlos0/env) | 698 | 16 | 2015-07-28 | 16 hours ago | Parse environment variables to Go structs (with defaults). |
| [envcfg](https://github.com/tomazk/envcfg) | 89 | 1 | 2014-11-29 | 1 year ago | Un-marshaling environment variables to Go structs. |
| [envconf](https://github.com/ian-kent/envconf) | 7 | 1 | 2014-10-26 | 4 years ago | Configuration from environment. |
| [envconfig](https://github.com/vrischmann/envconfig) | 136 | 4 | 2015-04-22 | 2 months ago | Read your configuration from environment variables. |
| [envh](https://github.com/antham/envh) | 91 | 3 | 2017-01-12 | 1 day ago | Helpers to manage environment variables. |
| [gcfg](https://github.com/go-gcfg/gcfg) | 102 | 4 | 2015-08-17 | 9 months ago | read INI-style configuration files into Go structs; supports user-defined types and subsections. |
| [go-up](https://github.com/ufoscout/go-up) | 19 | 1 | 2018-02-18 | 1 day ago | A simple configuration library with recursive placeholders resolution and no magic. |
| [goconfig](https://github.com/crgimenes/goconfig) | 91 | 11 | 2016-12-18 | 3 weeks ago | Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. |
| [godotenv](https://github.com/joho/godotenv) | 1745 | 29 | 2013-07-30 | 5 days ago | Go port of Ruby's dotenv library (Loads environment variables from `.env`). |
| [gofigure](https://github.com/ian-kent/gofigure) | 56 | 6 | 2014-11-25 | 1 year ago | Go application configuration made easy. |
| [gone](https://github.com/One-com/gone) | 24 | 7 | 2016-09-05 | 3 weeks ago | Modular JSON configuration. Keep you config structs along with the code they configure and delegate parsing to submodules without sacrificing full config serialization. |
| [config](https://github.com/gookit/config) | 42 | 3 | 2018-07-07 | 1 month ago | application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge. |
| [hjson-go](https://github.com/hjson/hjson-go) | 162 | 7 | 2016-08-06 | 3 weeks ago | Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments. |
| [ingo](https://github.com/schachmat/ingo) | 22 | 1 | 2016-02-08 | 1 year ago | Flags persisted in an ini-like config file. |
| [ini](https://github.com/go-ini/ini) | 1325 | 62 | 2014-12-18 | 1 week ago | Go package to read and write INI files. |
| [config](https://github.com/joshbetz/config) | 193 | 3 | 2017-04-03 | 1 year ago | Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP. |
| [envconfig](https://github.com/kelseyhightower/envconfig) | 2121 | 37 | 2013-11-07 | 2 weeks ago | Go library for managing configuration data from environment variables. |
| [mini](https://github.com/sasbury/mini) | 19 | 2 | 2015-04-30 | 2 months ago | Golang package for parsing ini-style configuration files. |
| [sprbox](https://github.com/oblq/sprbox) | 3 | 2 | 2018-07-17 | 4 months ago | Build-environment aware toolbox factory and agnostic config parser (YAML, TOML, JSON and Environment vars). |
| [store](https://github.com/tucnak/store) | 239 | 5 | 2015-10-04 | 1 year ago | Lightweight configuration manager for Go. |
| [viper](https://github.com/spf13/viper) | 7657 | 165 | 2014-04-02 | 1 day ago | Go configuration with fangs. |
| [xdg](https://github.com/OpenPeeDeeP/xdg) | 19 | 3 | 2017-07-20 | 1 week ago | Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html). |

## Continuous Integration
        
*Tools for help with continuous integration.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [drone](https://github.com/drone/drone) | 17344 | 584 | 2014-02-07 | 1 day ago | Drone is a Continuous Integration platform built on Docker, written in Go. |
| [duci](https://github.com/duck8823/duci) | 30 | 3 | 2018-04-01 | 3 hours ago | A simple ci server no needs domain specific languages. |
| [gomason](https://github.com/nikogura/gomason) | 20 | 0 | 2017-11-18 | 3 weeks ago | Test, Build, Sign, and Publish your go binaries from a clean workspace. |
| [goveralls](https://github.com/mattn/goveralls) | 540 | 16 | 2013-04-17 | 4 weeks ago | Go integration for Coveralls.io continuous code coverage tracking system. |
| [overalls](https://github.com/go-playground/overalls) | 94 | 4 | 2015-07-30 | 6 months ago | Multi-Package go project coverprofile for tools like goveralls. |
| [roveralls](https://github.com/lawrencewoodman/roveralls) | 11 | 1 | 2016-06-26 | 1 year ago | Recursive coverage testing tool. |

## CSS Preprocessors
        
*Libraries for preprocessing CSS files.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gcss](https://github.com/yosssi/gcss) | 408 | 16 | 2014-09-04 | 4 years ago | Pure Go CSS Preprocessor. |
| [go-libsass](https://github.com/wellington/go-libsass) | 116 | 4 | 2015-04-19 | 2 weeks ago | Go wrapper to the 100% Sass compatible libsass project. |

## Data Structures
        
*Generic datastructures and algorithms in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [algorithms](https://github.com/shady831213/algorithms) | 155 | 6 | 2018-01-31 | 1 month ago | Algorithms and data structures.CLRS study. |
| [binpacker](https://github.com/zhuangsirui/binpacker) | 110 | 10 | 2016-02-02 | 8 months ago | Binary packer and unpacker helps user build custom binary stream. |
| [bit](https://github.com/yourbasic/bit) | 29 | 3 | 2017-05-04 | 11 months ago | Golang set data structure with bonus bit-twiddling functions. |
| [bitset](https://github.com/willf/bitset) | 439 | 27 | 2011-05-11 | 2 days ago | Go package implementing bitsets. |
| [bloom](https://github.com/zentures/bloom) | 125 | 8 | 2013-09-03 | 10 months ago | Bloom filters implemented in Go. |
| [bloom](https://github.com/yourbasic/bloom) | 13 | 1 | 2017-05-07 | 1 year ago | Golang Bloom filter implementation. |
| [BoomFilters](https://github.com/tylertreat/BoomFilters) | 1080 | 36 | 2015-02-06 | 4 months ago | Probabilistic data structures for processing continuous, unbounded streams. |
| [concurrent-writer](https://github.com/free/concurrent-writer) | 15 | 4 | 2017-09-18 | 1 year ago | Highly concurrent drop-in replacement for `bufio.Writer`. |
| [conjungo](https://github.com/InVisionApp/conjungo) | 57 | 14 | 2016-12-30 | 2 weeks ago | A small, powerful and flexible merge library. |
| [count-min-log](https://github.com/seiflotfy/count-min-log) | 40 | 4 | 2015-08-17 | 2 years ago | Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory). |
| [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) | 443 | 17 | 2015-06-29 | 14 hours ago | Cuckoo filter: a good alternative to a counting bloom filter implemented in Go. |
| [deque](https://github.com/edwingeng/deque) | 1 | 1 | 2019-02-01 | 2 weeks ago | A highly optimized double-ended queue. |
| [deque](https://github.com/gammazero/deque) | 32 | 4 | 2018-04-24 | 1 month ago | Fast ring-buffer deque (double-ended queue). |
| [encoding](https://github.com/zentures/encoding) | 92 | 7 | 2013-09-21 | 1 year ago | Integer Compression Libraries for Go. |
| [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) | 63 | 5 | 2016-04-01 | 2 months ago | Go implementation of Adaptive Radix Tree. |
| [go-datastructures](https://github.com/Workiva/go-datastructures) | 4778 | 292 | 2014-10-29 | 3 months ago | Collection of useful, performant, and thread-safe data structures. |
| [go-ef](https://github.com/amallia/go-ef) | 8 | 1 | 2017-09-22 | 1 year ago | A Go implementation of the Elias-Fano encoding. |
| [go-geoindex](https://github.com/hailocab/go-geoindex) | 297 | 67 | 2015-01-22 | 1 year ago | In-memory geo index. |
| [go-mcache](https://github.com/OrlovEvgeny/go-mcache) | 23 | 2 | 2018-04-15 | 3 months ago | Fast in-memory key:value store/cache library. Pointer caches. |
| [go-rquad](https://github.com/arl/go-rquad) | 94 | 3 | 2016-09-13 | 8 months ago | Region quadtrees with efficient point location and neighbour finding. |
| [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) | 2 | 1 | 2019-01-11 | 1 month ago | Concurrent FIFO queue. |
| [gods](https://github.com/emirpasic/gods) | 5346 | 254 | 2015-03-04 | 12 hours ago | Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc. |
| [golang-set](https://github.com/deckarep/golang-set) | 975 | 27 | 2013-07-04 | 5 months ago | Thread-Safe and Non-Thread-Safe high-performance sets for Go. |
| [goset](https://github.com/zoumo/goset) | 12 | 1 | 2017-08-25 | 1 year ago | A useful Set collection implementation for Go. |
| [goskiplist](https://github.com/ryszard/goskiplist) | 178 | 12 | 2012-05-09 | 2 years ago | Skip list implementation in Go. |
| [gota](https://github.com/go-gota/gota) | 735 | 57 | 2016-02-07 | 4 days ago | Implementation of dataframes, series, and data wrangling methods for Go. |
| [hide](https://github.com/emvi/hide) | 2 | 2 | 2019-01-16 | 1 week ago | ID type with marshalling to/from hash to prevent sending IDs to clients. |
| [hilbert](https://github.com/google/hilbert) | 171 | 19 | 2015-08-06 | 3 months ago | Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves. |
| [hyperloglog](https://github.com/axiomhq/hyperloglog) | 634 | 18 | 2017-06-18 | 2 months ago | HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction. |
| [levenshtein](https://github.com/agext/levenshtein) | 25 | 1 | 2016-04-08 | 1 week ago | Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. |
| [levenshtein](https://github.com/agnivade/levenshtein) | 42 | 3 | 2014-07-30 | 3 days ago | Implementation to calculate levenshtein distance in Go. |
| [mafsa](https://github.com/smartystreets/mafsa) | 272 | 18 | 2014-11-08 | 2 years ago | MA-FSA implementation with Minimal Perfect Hashing. |
| [merkletree](https://github.com/cbergoon/merkletree) | 119 | 4 | 2017-04-12 | 2 months ago | Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures. |
| [mspm](https://github.com/BlackRabbitt/mspm) | 4 | 3 | 2018-05-18 | 9 months ago | Multi-String Pattern Matching Algorithm for information retrieval. |
| [null](https://github.com/emvi/null) | 2 | 2 | 2018-07-05 | 1 week ago | Nullable Go types that can be marshalled/unmarshalled to/from JSON. |
| [pipeline](https://github.com/hyfather/pipeline) | 9 | 1 | 2018-04-25 | 6 months ago | An implementation of pipelines with fan-in and fan-out. |
| [ring](https://github.com/TheTannerRyan/ring) | 83 | 1 | 2019-01-27 | 1 week ago | Go implementation of a high performance, thread safe bloom filter. |
| [roaring](https://github.com/RoaringBitmap/roaring) | 565 | 34 | 2014-07-11 | 1 month ago | Go package implementing compressed bitsets. |
| [set](https://github.com/StudioSol/set) | 5 | 17 | 2018-07-21 | 4 months ago | Simple set data structure implementation in Go using LinkedHashMap. |
| [skiplist](https://github.com/MauriceGit/skiplist) | 80 | 5 | 2018-06-24 | 2 months ago | Very fast Go Skiplist implementation. |
| [skiplist](https://github.com/gansidui/skiplist) | 58 | 7 | 2014-11-19 | 4 years ago | Skiplist implementation in Go. |
| [timedmap](https://github.com/zekroTJA/timedmap) | 0 | 1 | 2019-01-30 | 1 week ago | Map with expiring key-value pairs. |
| [trie](https://github.com/derekparker/trie) | 360 | 20 | 2014-03-07 | 9 months ago | Trie implementation in Go. |
| [ttlcache](https://github.com/ReneKroon/ttlcache) | 78 | 7 | 2014-12-13 | 5 days ago | In-memory LRU string-interface{} map with expiration for golang. |
| [bloom](https://github.com/willf/bloom) | 586 | 30 | 2011-05-21 | 2 days ago | Go package implementing Bloom filters. |

## Database
        
*Databases implemented in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [badger](https://github.com/dgraph-io/badger) | 5378 | 219 | 2017-01-26 | 1 day ago | Fast key-value store in Go. |
| [bigcache](https://github.com/allegro/bigcache) | 1879 | 64 | 2016-03-23 | 1 week ago | Efficient key/value cache for gigabytes of data. |
| [bolt](https://github.com/boltdb/bolt) | 9491 | 338 | 2013-12-21 | 1 year ago | Low-level key/value database for Go. |
| [buntdb](https://github.com/tidwall/buntdb) | 2262 | 93 | 2016-07-20 | 4 weeks ago | Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support. |
| [cache2go](https://github.com/muesli/cache2go) | 812 | 57 | 2013-11-11 | 4 months ago | In-memory key:value cache which supports automatic invalidation based on timeouts. |
| [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) | 25 | 5 | 2017-12-18 | 1 year ago | BigCache with clustering support and individual item expiration. |
| [cockroach](https://github.com/cockroachdb/cockroach) | 15523 | 693 | 2014-02-06 | 7 hours ago | Scalable, Geo-Replicated, Transactional Datastore. |
| [couchcache](https://github.com/codingsince1985/couchcache) | 39 | 3 | 2015-04-05 | 1 year ago | RESTful caching micro-service backed by Couchbase server. |
| [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) | 525 | 45 | 2018-04-11 | 2 days ago | CovenantSQL is a SQL database on blockchain. |
| [dgraph](https://github.com/dgraph-io/dgraph) | 8702 | 329 | 2015-08-25 | 6 hours ago | Scalable, Distributed, Low Latency, High Throughput Graph Database. |
| [diskv](https://github.com/peterbourgon/diskv) | 695 | 29 | 2012-03-22 | 6 months ago | Home-grown disk-backed key-value store. |
| [eliasdb](https://github.com/krotik/eliasdb) | 503 | 23 | 2016-08-13 | 2 weeks ago | Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language. |
| [fastcache](https://github.com/VictoriaMetrics/fastcache) | 348 | 12 | 2018-11-23 | 6 days ago | fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead. |
| [gcache](https://github.com/bluele/gcache) | 665 | 29 | 2015-01-25 | 2 days ago | Cache library with support for expirable Cache, LFU, LRU and ARC. |

## Database Drivers
        

## Date and Time
        

## Distributed Systems
        

## Email
        

## Embeddable Scripting Languages
        

## Error Handling
        

## Files
        

## Financial
        

## Forms
        

## Functional
        

## Game Development
        

## Generation and Generics
        

## Geographic
        

## Go Compilers
        

## Goroutines
        

## GUI
        

## Hardware
        

## Images
        

## IoT
        

## Job Scheduler
        

## JSON
        

## Logging
        

## Machine Learning
        

## Messaging
        

## Microsoft Office
        

### Microsoft Excel
        

## Miscellaneous
        

### Dependency Injection
        

### Strings
        

## Natural Language Processing
        

## Networking
        

### HTTP Clients
        

## OpenGL
        

## ORM
        

## Package Management
        

## Query Language
        

## Resource Embedding
        

## Science and Data Analysis
        

## Security
        

## Serialization
        

## Template Engines
        

## Testing
        

## Text Processing
        

## Third-party APIs
        

## Utilities
        

## UUID
        

## Validation
        

## Version Control
        

## Video
        

## Web Frameworks
        

### Middlewares
        

#### Actual middlewares
        

#### Libraries for creating HTTP middlewares
        

### Routers
        

## Windows
        

## XML
        

# Tools
        

## Code Analysis
        

## Editor Plugins
        

## Go Generate Tools
        

## Go Tools
        

## Software Packages
        

### DevOps Tools
        

### Other Software
        

# Server Applications
        

# Resources
        

## Benchmarks
        

## Conferences
        

## E-Books
        

## Gophers
        

## Meetups
        

## Twitter
        

## Websites
        

### Tutorials
        