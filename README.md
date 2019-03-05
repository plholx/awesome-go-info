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
| [portaudio](https://github.com/gordonklaus/portaudio) | 264 | 9 | 2015-09-16 | 6 months ago | Go bindings for the PortAudio audio I/O library. |
| [music-theory](https://github.com/go-music-theory/music-theory) | 233 | 10 | 2016-03-17 | 8 months ago | Music theory models in Go. |
| [waveform](https://github.com/mdlayher/waveform) | 228 | 13 | 2014-09-14 | 2 years ago | Go package capable of generating waveform images from audio streams. |
| [portmidi](https://github.com/rakyll/portmidi) | 187 | 6 | 2013-11-10 | 1 year ago | Go bindings for PortMidi. |
| [flac](https://github.com/mewkiz/flac) | 92 | 6 | 2012-11-02 | 1 week ago | Native Go FLAC encoder/decoder with support for FLAC streams. |
| [mix](https://github.com/go-mix/mix) | 90 | 3 | 2016-01-03 | 1 year ago | Sequence-based Go-native audio mixer for music apps. |
| [id3v2](https://github.com/bogem/id3v2) | 90 | 5 | 2016-05-16 | 3 weeks ago | Fast and stable ID3 parsing and writing library for Go. |
| [go-sox](https://github.com/krig/go-sox) | 87 | 8 | 2013-10-08 | 8 months ago | libsox bindings for go. |
| [mp3](https://github.com/tcolgate/mp3) | 85 | 1 | 2015-02-27 | 1 year ago | Native Go MP3 decoder. |
| [flac](https://github.com/eaburns/flac) | 81 | 5 | 2013-08-21 | 1 year ago | No-frills native Go FLAC decoder that decodes FLAC files into byte slices. |
| [go-taglib](https://github.com/wtolson/go-taglib) | 64 | 6 | 2012-11-20 | 7 months ago | Go bindings for taglib. |
| [malgo](https://github.com/gen2brain/malgo) | 55 | 4 | 2017-11-10 | 1 month ago | Mini audio library. |
| [gaad](https://github.com/Comcast/gaad) | 50 | 10 | 2016-07-11 | 1 year ago | Native Go AAC bitstream parser. |
| [go_mediainfo](https://github.com/zhulik/go_mediainfo) | 22 | 1 | 2015-12-14 | 3 years ago | libmediainfo bindings for go. |
| [vorbis](https://github.com/mccoyst/vorbis) | 21 | 3 | 2013-07-12 | 1 year ago | "Native" Go Vorbis decoder (uses CGO, but has no dependencies). |
| [minimp3](https://github.com/tosone/minimp3) | 19 | 2 | 2018-01-26 | 2 days ago | Lightweight MP3 decoder library. |
| [EasyMIDI](https://github.com/algoGuy/EasyMIDI) | 17 | 3 | 2018-02-20 | 11 months ago | EasyMidi is a simple and reliable library for working with standard midi file (SMF). |
| [gosamplerate](https://github.com/dh1tw/gosamplerate) | 8 | 1 | 2016-11-21 | 9 months ago | libsamplerate bindings for go. |

## Authentication and OAuth
        
*Libraries for implementing authentications schemes.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [jwt-go](https://github.com/dgrijalva/jwt-go) | 5019 | 130 | 2012-04-18 | 1 week ago | Golang implementation of JSON Web Tokens (JWT). |
| [casbin](https://github.com/casbin/casbin) | 3864 | 140 | 2017-04-08 | 2 weeks ago | Authorization library that supports access control models like ACL, RBAC, ABAC. |
| [oauth2](https://github.com/golang/oauth2) | 2114 | 92 | 2014-04-14 | 6 days ago | Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support. |
| [goth](https://github.com/markbates/goth) | 2082 | 58 | 2014-10-15 | 1 week ago | provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box. |
| [authboss](https://github.com/volatiletech/authboss) | 1739 | 38 | 2015-01-03 | 3 weeks ago | Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time. |
| [osin](https://github.com/openshift/osin) | 1507 | 69 | 2013-09-11 | 4 months ago | Golang OAuth2 server library. |
| [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) | 1117 | 72 | 2015-11-01 | 1 week ago | Standalone, specification-compliant,  OAuth2 server written in Golang. |
| [go-jose](https://github.com/square/go-jose) | 1015 | 61 | 2014-11-15 | 6 days ago | Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs. |
| [gologin](https://github.com/dghubble/gologin) | 960 | 27 | 2015-06-23 | 2 weeks ago | chainable handlers for login with OAuth1 and OAuth2 authentication providers. |
| [gorbac](https://github.com/mikespook/gorbac) | 829 | 55 | 2013-12-26 | 1 month ago | provides a lightweight role-based access control (RBAC) implementation in Golang. |
| [loginsrv](https://github.com/tarent/loginsrv) | 735 | 45 | 2016-11-11 | 2 days ago | JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam. |
| [permissions2](https://github.com/xyproto/permissions2) | 299 | 12 | 2014-11-19 | 1 week ago | Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt. |
| [paseto](https://github.com/o1egl/paseto) | 194 | 9 | 2018-01-23 | 3 months ago | Golang implementation of Platform-Agnostic Security Tokens (PASETO). |
| [httpauth](https://github.com/goji/httpauth) | 167 | 5 | 2014-05-27 | 2 years ago | HTTP Authentication middleware. |
| [jwt-auth](https://github.com/adam-hanna/jwt-auth) | 144 | 9 | 2016-07-06 | 1 month ago | JWT middleware for Golang http servers with many configuration options. |
| [session](https://github.com/icza/session) | 82 | 6 | 2016-02-08 | 5 months ago | Go session management for web servers (including support for Google App Engine - GAE). |
| [jwt](https://github.com/robbert229/jwt) | 60 | 6 | 2016-06-06 | 3 months ago | Clean and easy to use implementation of JSON Web Tokens (JWT). |
| [branca](https://github.com/hako/branca) | 50 | 5 | 2018-01-09 | 6 months ago | Golang implementation of Branca Tokens. |
| [sessions](https://github.com/adam-hanna/sessions) | 41 | 3 | 2017-04-29 | 1 year ago | Dead simple, highly performant, highly customizable sessions service for go http servers. |
| [jwt](https://github.com/pascaldekloe/jwt) | 39 | 2 | 2018-03-21 | 12 hours ago | Lightweight JSON Web Token (JWT) library. |
| [securecookie](https://github.com/chmike/securecookie) | 26 | 4 | 2017-09-03 | 6 months ago | Efficient secure cookie encoding/decoding. |
| [rbac](https://github.com/zpatrick/rbac) | 20 | 3 | 2018-08-02 | 6 months ago | Minimalistic RBAC package for Go applications. |
| [signedvalue](https://github.com/sashka/signedvalue) | 7 | 0 | 2018-01-06 | 1 month ago | Signed and timestamped strings compatible with [Tornado's](https://github.com/tornadoweb/tornado) `create_signed_value`, `decode_signed_value`, and therefore `set_secure_cookie` and `get_secure_cookie`. |
| [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) | 6 | 2 | 2017-10-20 | 3 months ago | Go session management using the SessionGate Redis module. |
| [cookiestxt](https://github.com/mengzhuo/cookiestxt) | 2 | 1 | 2017-10-09 | 1 year ago | provides parser of cookies.txt file format. |

## Bot Building
        
*Libraries for building and working with bots.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) | 1385 | 56 | 2015-06-25 | 3 days ago | Simple and clean Telegram bot client. |
| [telebot](https://github.com/tucnak/telebot) | 827 | 34 | 2015-06-26 | 1 day ago | Telegram bot framework written in Go. |
| [bot](https://github.com/go-chat-bot/bot) | 414 | 34 | 2015-09-23 | 1 week ago | IRC, Slack & Telegram bot written in Go. |
| [slacker](https://github.com/shomali11/slacker) | 256 | 13 | 2017-05-20 | 1 week ago | Easy to use framework to create Slack bots. |
| [tbot](https://github.com/yanzay/tbot) | 184 | 8 | 2015-09-12 | 2 months ago | Telegram bot server with API similar to net/http. |
| [tenyks](https://github.com/kyleterry/tenyks) | 166 | 14 | 2012-08-26 | 2 years ago | Service oriented IRC bot using Redis and JSON for messaging. |
| [golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot) | 148 | 19 | 2017-05-15 | 3 weeks ago | A golang implementation of a console-based trading bot for cryptocurrency exchanges. |
| [go-sarah](https://github.com/oklahomer/go-sarah) | 111 | 3 | 2016-11-06 | 7 months ago | Framework to build bot for desired chat services including LINE, Slack, Gitter and more. |
| [hanu](https://github.com/sbstjn/hanu) | 90 | 6 | 2016-09-16 | 6 months ago | Framework for writing Slack bots. |
| [go-tgbot](https://github.com/olebedev/go-tgbot) | 80 | 7 | 2016-12-11 | 8 months ago | Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware. |
| [margelet](https://github.com/zhulik/margelet) | 51 | 5 | 2015-11-21 | 2 years ago | Framework for building Telegram bots. |
| [govkbot](https://github.com/nikepan/govkbot) | 19 | 2 | 2016-07-12 | 14 hours ago | Simple Go [VK](https://vk.com) bot library. |
| [micha](https://github.com/onrik/micha) | 9 | 3 | 2016-04-14 | 1 year ago | Go Library for Telegram bot api. |

## Command Line
        
*Libraries for building Console Applications and Console User Interfaces.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [cobra](https://github.com/spf13/cobra) | 10778 | 255 | 2013-09-04 | 1 week ago | Commander for modern Go CLI interactions. |
| [cli](https://github.com/urfave/cli) | 10083 | 271 | 2013-07-14 | 3 weeks ago | Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli). |
| [termui](https://github.com/gizak/termui) | 8362 | 275 | 2015-02-03 | 3 days ago | Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib). |
| [gocui](https://github.com/jroimartin/gocui) | 3970 | 102 | 2014-01-04 | 15 hours ago | Minimalist Go library aimed at creating Console User Interfaces. |
| [termbox-go](https://github.com/nsf/termbox-go) | 3292 | 103 | 2012-01-13 | 2 weeks ago | Termbox is a library for creating cross-platform text-based interfaces. |
| [color](https://github.com/fatih/color) | 2864 | 56 | 2014-02-17 | 4 months ago | Versatile package for colored terminal output. |
| [kingpin](https://github.com/alecthomas/kingpin) | 2298 | 54 | 2014-05-15 | 1 month ago | Command line and flag parser supporting sub commands. |
| [go-prompt](https://github.com/c-bata/go-prompt) | 2081 | 35 | 2017-08-15 | 1 week ago | Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit). |
| [readline](https://github.com/chzyer/readline) | 1307 | 39 | 2015-09-20 | 3 weeks ago | Pure golang implementation that provides most features in GNU-Readline under MIT license. |
| [go-flags](https://github.com/jessevdk/go-flags) | 1292 | 28 | 2012-08-31 | 1 week ago | go command line option parser. |
| [uiprogress](https://github.com/gosuri/uiprogress) | 1229 | 27 | 2015-11-17 | 4 months ago | Flexible library to render progress bars in terminal applications. |
| [docopt.go](https://github.com/docopt/docopt.go) | 1036 | 36 | 2013-08-26 | 5 months ago | Command-line arguments parser that will make you smile. |
| [asciigraph](https://github.com/guptarohit/asciigraph) | 1022 | 23 | 2018-06-17 | 1 month ago | Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. |
| [cli](https://github.com/mitchellh/cli) | 927 | 25 | 2013-11-03 | 3 months ago | Go library for implementing command-line interfaces. |
| [gcli](https://github.com/tcnksm/gcli) | 856 | 25 | 2014-06-19 | 1 year ago | The easy way to start building Golang command line applications. |
| [uilive](https://github.com/gosuri/uilive) | 724 | 11 | 2015-11-16 | 3 months ago | Library for updating terminal output in realtime. |
| [pflag](https://github.com/spf13/pflag) | 625 | 24 | 2013-08-30 | 3 weeks ago | Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. |
| [mow.cli](https://github.com/jawher/mow.cli) | 590 | 19 | 2014-12-19 | 6 days ago | Go library for building CLI applications with sophisticated flag and argument parsing and validation. |
| [go-arg](https://github.com/alexflint/go-arg) | 571 | 15 | 2015-11-01 | 2 months ago | Struct-based argument parsing in Go. |
| [complete](https://github.com/posener/complete) | 542 | 12 | 2017-05-06 | 1 week ago | Write bash completions in Go + Go command bash completion. |
| [liner](https://github.com/peterh/liner) | 521 | 22 | 2012-08-16 | 3 weeks ago | Go readline-like library for command-line interfaces. |
| [progressbar](https://github.com/schollz/progressbar) | 494 | 13 | 2017-10-27 | 1 month ago | Basic thread-safe progress bar that works in every OS. |
| [uitable](https://github.com/gosuri/uitable) | 464 | 15 | 2015-11-14 | 1 year ago | Library to improve readability in terminal apps using tabular data. |
| [aurora](https://github.com/logrusorgru/aurora) | 428 | 5 | 2016-11-07 | 5 months ago | ANSI terminal colors that supports fmt.Printf/Sprintf. |
| [cli](https://github.com/mkideal/cli) | 427 | 20 | 2016-02-27 | 4 days ago | Feature-rich and easy to use command-line package based on golang struct tags. |
| [mpb](https://github.com/vbauerster/mpb) | 419 | 8 | 2016-12-14 | 6 days ago | Multi progress bar for terminal applications. |
| [flaggy](https://github.com/integrii/flaggy) | 410 | 11 | 2018-03-05 | 1 week ago | A robust and idiomatic flags package with excellent subcommand support. |
| [go-colorable](https://github.com/mattn/go-colorable) | 334 | 16 | 2014-07-30 | 1 week ago | Colorable writer for windows. |
| [go-isatty](https://github.com/mattn/go-isatty) | 298 | 7 | 2014-04-01 | 1 week ago | isatty for golang. |
| [chalk](https://github.com/ttacon/chalk) | 286 | 11 | 2014-07-19 | 2 years ago | Intuitive package for prettifying terminal/console output. |
| [gommon](https://github.com/labstack/gommon) | 270 | 19 | 2015-03-13 | 1 month ago | Style terminal text. |
| [termtables](https://github.com/apcera/termtables) | 203 | 68 | 2012-12-06 | 1 year ago | Go port of the Ruby library [terminal-tables](https://github.com/tj/terminal-table) for simple ASCII table generation as well as providing markdown and HTML output. |
| [go-colortext](https://github.com/daviddengcn/go-colortext) | 188 | 9 | 2013-01-23 | 10 months ago | Go library for color output in terminals. |
| [color](https://github.com/gookit/color) | 129 | 4 | 2018-07-01 | 1 week ago | Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows. |
| [simpletable](https://github.com/alexeyco/simpletable) | 121 | 2 | 2017-03-29 | 1 week ago | Simple tables in terminal with Go. |
| [clif](https://github.com/ukautz/clif) | 95 | 2 | 2015-05-31 | 2 weeks ago | Small command line interface framework. |
| [flag](https://github.com/cosiner/flag) | 90 | 5 | 2016-10-06 | 1 week ago | Simple but powerful command line option parsing library for Go supporting subcommand. |
| [termdash](https://github.com/mum4k/termdash) | 82 | 6 | 2018-03-24 | 1 day ago | Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui). |
| [argparse](https://github.com/akamensky/argparse) | 74 | 5 | 2017-11-24 | 1 month ago | Command line argument parser inspired by Python's argparse module. |
| [sflags](https://github.com/octago/sflags) | 71 | 5 | 2016-12-04 | 1 month ago | Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries. |
| [commandeer](https://github.com/jaffee/commandeer) | 70 | 4 | 2017-10-12 | 3 days ago | Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags. |
| [wmenu](https://github.com/dixonwille/wmenu) | 67 | 2 | 2016-04-20 | 1 week ago | Easy to use menu structure for cli applications that prompts users to make choices. |
| [hiboot](https://github.com/hidevopsio/hiboot) | 61 | 6 | 2018-03-16 | 6 days ago | cli application framework with auto configuration and dependency injection. |
| [cfmt](https://github.com/mingrammer/cfmt) | 53 | 3 | 2018-03-16 | 2 months ago | Contextual fmt inspired by bootstrap color classes. |
| [cli](https://github.com/teris-io/cli) | 43 | 1 | 2017-05-25 | 8 months ago | Simple and complete API for building command line interfaces in Go. |
| [env](https://github.com/codingconcepts/env) | 37 | 1 | 2017-06-15 | 4 months ago | Tag-based environment configuration for structs. |
| [wlog](https://github.com/dixonwille/wlog) | 29 | 1 | 2016-04-14 | 1 year ago | Simple logging interface that supports cross-platform color and concurrency. |
| [gocmd](https://github.com/devfacet/gocmd) | 28 | 3 | 2018-01-08 | 6 months ago | Go library for building command line applications. |
| [flagvar](https://github.com/sgreben/flagvar) | 25 | 1 | 2018-05-19 | 4 months ago | A collection of flag argument types for Go's standard `flag` package. |
| [tabular](https://github.com/InVisionApp/tabular) | 23 | 3 | 2018-04-24 | 9 months ago | Print ASCII tables from command line utilities without the need to pass large sets of data to the API. |
| [strumt](https://github.com/antham/strumt) | 22 | 0 | 2017-06-20 | 2 days ago | Library to create prompt chain. |
| [colourize](https://github.com/TreyBastian/colourize) | 15 | 3 | 2015-05-11 | 2 years ago | Go library for ANSI colour text in terminals. |
| [argv](https://github.com/cosiner/argv) | 11 | 1 | 2017-01-22 | 1 month ago | Go library to split command line string as arguments array using the bash syntax. |
| [go-commander](https://github.com/yitsushi/go-commander) | 8 | 1 | 2016-10-10 | 6 days ago | Go library to simplify CLI workflow. |
| [go-ataman](https://github.com/workanator/go-ataman) | 8 | 2 | 2017-05-18 | 1 year ago | Go library for rendering ANSI colored text templates in terminals. |
| [ctc](https://github.com/wzshiming/ctc) | 6 | 1 | 2018-04-28 | 4 months ago | The non-invasive cross-platform terminal color library does not need to modify the Print method. |
| [go-getoptions](https://github.com/DavidGamba/go-getoptions) | 5 | 1 | 2015-12-18 | 2 days ago | Go option parser inspired on the flexibility of Perl’s GetOpt::Long. |
| [sand](https://github.com/Zaba505/sand) | 2 | 1 | 2018-11-19 | 3 months ago | Simple API for creating interpreters and so much more. |

## Configuration
        
*Libraries for configuration parsing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [viper](https://github.com/spf13/viper) | 7663 | 165 | 2014-04-02 | 3 days ago | Go configuration with fangs. |
| [envconfig](https://github.com/kelseyhightower/envconfig) | 2123 | 37 | 2013-11-07 | 2 weeks ago | Go library for managing configuration data from environment variables. |
| [godotenv](https://github.com/joho/godotenv) | 1747 | 30 | 2013-07-30 | 6 days ago | Go port of Ruby's dotenv library (Loads environment variables from `.env`). |
| [ini](https://github.com/go-ini/ini) | 1327 | 62 | 2014-12-18 | 2 weeks ago | Go package to read and write INI files. |
| [env](https://github.com/caarlos0/env) | 700 | 16 | 2015-07-28 | 2 days ago | Parse environment variables to Go structs (with defaults). |
| [store](https://github.com/tucnak/store) | 239 | 5 | 2015-10-04 | 1 year ago | Lightweight configuration manager for Go. |
| [confita](https://github.com/heetch/confita) | 208 | 25 | 2017-12-21 | 2 weeks ago | Load configuration in cascade from multiple backends into a struct. |
| [config](https://github.com/joshbetz/config) | 193 | 3 | 2017-04-03 | 1 year ago | Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP. |
| [config](https://github.com/olebedev/config) | 190 | 8 | 2014-04-21 | 5 months ago | JSON or YAML configuration wrapper with environment variables and flags parsing. |
| [hjson-go](https://github.com/hjson/hjson-go) | 163 | 7 | 2016-08-06 | 3 weeks ago | Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments. |
| [envconfig](https://github.com/vrischmann/envconfig) | 136 | 4 | 2015-04-22 | 2 months ago | Read your configuration from environment variables. |
| [gcfg](https://github.com/go-gcfg/gcfg) | 102 | 4 | 2015-08-17 | 9 months ago | read INI-style configuration files into Go structs; supports user-defined types and subsections. |
| [envh](https://github.com/antham/envh) | 92 | 3 | 2017-01-12 | 3 days ago | Helpers to manage environment variables. |
| [goconfig](https://github.com/crgimenes/goconfig) | 92 | 11 | 2016-12-18 | 3 weeks ago | Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. |
| [envcfg](https://github.com/tomazk/envcfg) | 89 | 1 | 2014-11-29 | 1 year ago | Un-marshaling environment variables to Go structs. |
| [gofigure](https://github.com/ian-kent/gofigure) | 56 | 6 | 2014-11-25 | 1 year ago | Go application configuration made easy. |
| [configure](https://github.com/paked/configure) | 44 | 3 | 2015-06-14 | 2 weeks ago | Provides configuration through multiple sources, including JSON, flags and environment variables. |
| [config](https://github.com/gookit/config) | 43 | 3 | 2018-07-07 | 1 month ago | application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge. |
| [ingo](https://github.com/schachmat/ingo) | 22 | 1 | 2016-02-08 | 1 year ago | Flags persisted in an ini-like config file. |
| [go-up](https://github.com/ufoscout/go-up) | 20 | 1 | 2018-02-18 | 1 day ago | A simple configuration library with recursive placeholders resolution and no magic. |
| [xdg](https://github.com/OpenPeeDeeP/xdg) | 19 | 3 | 2017-07-20 | 1 week ago | Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html). |
| [mini](https://github.com/sasbury/mini) | 19 | 2 | 2015-04-30 | 2 months ago | Golang package for parsing ini-style configuration files. |
| [envconf](https://github.com/ian-kent/envconf) | 7 | 1 | 2014-10-26 | 4 years ago | Configuration from environment. |
| [conflate](https://github.com/the4thamigo-uk/conflate) | 4 | 0 | 2018-02-02 | 3 weeks ago | Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema. |
| [sprbox](https://github.com/oblq/sprbox) | 3 | 2 | 2018-07-17 | 4 months ago | Build-environment aware toolbox factory and agnostic config parser (YAML, TOML, JSON and Environment vars). |

## Continuous Integration
        
*Tools for help with continuous integration.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [drone](https://github.com/drone/drone) | 17355 | 584 | 2014-02-07 | 3 days ago | Drone is a Continuous Integration platform built on Docker, written in Go. |
| [goveralls](https://github.com/mattn/goveralls) | 540 | 16 | 2013-04-17 | 1 month ago | Go integration for Coveralls.io continuous code coverage tracking system. |
| [overalls](https://github.com/go-playground/overalls) | 94 | 4 | 2015-07-30 | 6 months ago | Multi-Package go project coverprofile for tools like goveralls. |
| [duci](https://github.com/duck8823/duci) | 31 | 3 | 2018-04-01 | 1 day ago | A simple ci server no needs domain specific languages. |
| [gomason](https://github.com/nikogura/gomason) | 21 | 0 | 2017-11-18 | 3 weeks ago | Test, Build, Sign, and Publish your go binaries from a clean workspace. |
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
| [gods](https://github.com/emirpasic/gods) | 5353 | 254 | 2015-03-04 | 2 days ago | Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc. |
| [go-datastructures](https://github.com/Workiva/go-datastructures) | 4780 | 292 | 2014-10-29 | 3 months ago | Collection of useful, performant, and thread-safe data structures. |
| [BoomFilters](https://github.com/tylertreat/BoomFilters) | 1081 | 36 | 2015-02-06 | 4 months ago | Probabilistic data structures for processing continuous, unbounded streams. |
| [golang-set](https://github.com/deckarep/golang-set) | 976 | 27 | 2013-07-04 | 5 months ago | Thread-Safe and Non-Thread-Safe high-performance sets for Go. |
| [gota](https://github.com/go-gota/gota) | 737 | 57 | 2016-02-07 | 6 days ago | Implementation of dataframes, series, and data wrangling methods for Go. |
| [hyperloglog](https://github.com/axiomhq/hyperloglog) | 634 | 18 | 2017-06-18 | 1 day ago | HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction. |
| [bloom](https://github.com/willf/bloom) | 588 | 30 | 2011-05-21 | 4 days ago | Go package implementing Bloom filters. |
| [roaring](https://github.com/RoaringBitmap/roaring) | 567 | 34 | 2014-07-11 | 1 month ago | Go package implementing compressed bitsets. |
| [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) | 444 | 17 | 2015-06-29 | 2 days ago | Cuckoo filter: a good alternative to a counting bloom filter implemented in Go. |
| [bitset](https://github.com/willf/bitset) | 440 | 27 | 2011-05-11 | 4 days ago | Go package implementing bitsets. |
| [trie](https://github.com/derekparker/trie) | 363 | 20 | 2014-03-07 | 9 months ago | Trie implementation in Go. |
| [go-geoindex](https://github.com/hailocab/go-geoindex) | 297 | 67 | 2015-01-22 | 1 year ago | In-memory geo index. |
| [mafsa](https://github.com/smartystreets/mafsa) | 272 | 18 | 2014-11-08 | 2 years ago | MA-FSA implementation with Minimal Perfect Hashing. |
| [goskiplist](https://github.com/ryszard/goskiplist) | 178 | 12 | 2012-05-09 | 2 years ago | Skip list implementation in Go. |
| [hilbert](https://github.com/google/hilbert) | 171 | 19 | 2015-08-06 | 3 months ago | Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves. |
| [algorithms](https://github.com/shady831213/algorithms) | 155 | 6 | 2018-01-31 | 1 month ago | Algorithms and data structures.CLRS study. |
| [bloom](https://github.com/zentures/bloom) | 125 | 8 | 2013-09-03 | 10 months ago | Bloom filters implemented in Go. |
| [merkletree](https://github.com/cbergoon/merkletree) | 119 | 4 | 2017-04-12 | 14 hours ago | Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures. |
| [binpacker](https://github.com/zhuangsirui/binpacker) | 110 | 10 | 2016-02-02 | 8 months ago | Binary packer and unpacker helps user build custom binary stream. |
| [go-rquad](https://github.com/arl/go-rquad) | 94 | 3 | 2016-09-13 | 8 months ago | Region quadtrees with efficient point location and neighbour finding. |
| [encoding](https://github.com/zentures/encoding) | 92 | 7 | 2013-09-21 | 1 year ago | Integer Compression Libraries for Go. |
| [ring](https://github.com/TheTannerRyan/ring) | 83 | 1 | 2019-01-27 | 1 week ago | Go implementation of a high performance, thread safe bloom filter. |
| [skiplist](https://github.com/MauriceGit/skiplist) | 80 | 5 | 2018-06-24 | 2 months ago | Very fast Go Skiplist implementation. |
| [ttlcache](https://github.com/ReneKroon/ttlcache) | 78 | 7 | 2014-12-13 | 1 week ago | In-memory LRU string-interface{} map with expiration for golang. |
| [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) | 63 | 5 | 2016-04-01 | 2 months ago | Go implementation of Adaptive Radix Tree. |
| [skiplist](https://github.com/gansidui/skiplist) | 58 | 7 | 2014-11-19 | 4 years ago | Skiplist implementation in Go. |
| [conjungo](https://github.com/InVisionApp/conjungo) | 57 | 14 | 2016-12-30 | 2 weeks ago | A small, powerful and flexible merge library. |
| [levenshtein](https://github.com/agnivade/levenshtein) | 42 | 3 | 2014-07-30 | 5 days ago | Implementation to calculate levenshtein distance in Go. |
| [count-min-log](https://github.com/seiflotfy/count-min-log) | 40 | 4 | 2015-08-17 | 2 years ago | Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory). |
| [deque](https://github.com/gammazero/deque) | 34 | 4 | 2018-04-24 | 1 month ago | Fast ring-buffer deque (double-ended queue). |
| [bit](https://github.com/yourbasic/bit) | 29 | 3 | 2017-05-04 | 11 months ago | Golang set data structure with bonus bit-twiddling functions. |
| [levenshtein](https://github.com/agext/levenshtein) | 25 | 1 | 2016-04-08 | 1 week ago | Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. |
| [go-mcache](https://github.com/OrlovEvgeny/go-mcache) | 23 | 2 | 2018-04-15 | 3 months ago | Fast in-memory key:value store/cache library. Pointer caches. |
| [concurrent-writer](https://github.com/free/concurrent-writer) | 15 | 4 | 2017-09-18 | 1 year ago | Highly concurrent drop-in replacement for `bufio.Writer`. |
| [bloom](https://github.com/yourbasic/bloom) | 13 | 1 | 2017-05-07 | 1 year ago | Golang Bloom filter implementation. |
| [goset](https://github.com/zoumo/goset) | 12 | 1 | 2017-08-25 | 1 year ago | A useful Set collection implementation for Go. |
| [pipeline](https://github.com/hyfather/pipeline) | 9 | 1 | 2018-04-25 | 6 months ago | An implementation of pipelines with fan-in and fan-out. |
| [go-ef](https://github.com/amallia/go-ef) | 8 | 1 | 2017-09-22 | 1 year ago | A Go implementation of the Elias-Fano encoding. |
| [set](https://github.com/StudioSol/set) | 5 | 17 | 2018-07-21 | 4 months ago | Simple set data structure implementation in Go using LinkedHashMap. |
| [mspm](https://github.com/BlackRabbitt/mspm) | 4 | 3 | 2018-05-18 | 9 months ago | Multi-String Pattern Matching Algorithm for information retrieval. |
| [hide](https://github.com/emvi/hide) | 2 | 2 | 2019-01-16 | 1 week ago | ID type with marshalling to/from hash to prevent sending IDs to clients. |
| [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) | 2 | 1 | 2019-01-11 | 1 month ago | Concurrent FIFO queue. |
| [null](https://github.com/emvi/null) | 2 | 2 | 2018-07-05 | 1 week ago | Nullable Go types that can be marshalled/unmarshalled to/from JSON. |
| [deque](https://github.com/edwingeng/deque) | 1 | 1 | 2019-02-01 | 2 weeks ago | A highly optimized double-ended queue. |
| [timedmap](https://github.com/zekroTJA/timedmap) | 0 | 1 | 2019-01-30 | 1 week ago | Map with expiring key-value pairs. |

## Database
        
*SQL query builder, libraries for building and using SQL.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [prometheus](https://github.com/prometheus/prometheus) | 22394 | 1005 | 2012-11-24 | 12 hours ago | Monitoring system and time series database. |
| [tidb](https://github.com/pingcap/tidb) | 17553 | 1196 | 2015-09-06 | 11 hours ago | TiDB is a distributed SQL database. Inspired by the design of Google F1. |
| [influxdb](https://github.com/influxdata/influxdb) | 15705 | 749 | 2013-09-26 | 16 hours ago | Scalable datastore for metrics, events, and real-time analytics. |
| [cockroach](https://github.com/cockroachdb/cockroach) | 15530 | 694 | 2014-02-06 | 11 hours ago | Scalable, Geo-Replicated, Transactional Datastore. |
| [bolt](https://github.com/boltdb/bolt) | 9493 | 338 | 2013-12-21 | 1 year ago | Low-level key/value database for Go. |
| [dgraph](https://github.com/dgraph-io/dgraph) | 8728 | 331 | 2015-08-25 | 19 hours ago | Scalable, Distributed, Low Latency, High Throughput Graph Database. |
| [vitess](https://github.com/vitessio/vitess) | 7608 | 464 | 2013-06-28 | 15 hours ago | vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. |
| [groupcache](https://github.com/golang/groupcache) | 7098 | 443 | 2013-07-23 | 2 weeks ago | Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. |
| [pgweb](https://github.com/sosedoff/pgweb) | 5757 | 152 | 2014-10-09 | 6 days ago | Web-based PostgreSQL database browser. |
| [badger](https://github.com/dgraph-io/badger) | 5385 | 220 | 2017-01-26 | 1 day ago | Fast key-value store in Go. |
| [rqlite](https://github.com/rqlite/rqlite) | 4259 | 182 | 2014-08-23 | 6 days ago | The lightweight, distributed, relational database built on SQLite. |
| [kingshard](https://github.com/flike/kingshard) | 4185 | 388 | 2015-07-04 | 1 week ago | kingshard is a high performance proxy for MySQL powered by Golang. |
| [ledisdb](https://github.com/siddontang/ledisdb) | 2889 | 184 | 2014-04-30 | 3 weeks ago | Ledisdb is a high performance NoSQL like Redis based on LevelDB. |
| [goleveldb](https://github.com/syndtr/goleveldb) | 2795 | 151 | 2013-01-23 | 13 hours ago | Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go. |
| [orchestrator](https://github.com/github/orchestrator) | 2596 | 181 | 2016-11-30 | 19 hours ago | MySQL replication topology manager & visualizer. |
| [go-cache](https://github.com/patrickmn/go-cache) | 2422 | 84 | 2012-01-02 | 6 days ago | In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. |
| [tiedot](https://github.com/HouzuoGuo/tiedot) | 2306 | 161 | 2013-05-26 | 1 month ago | Your NoSQL database powered by Golang. |
| [buntdb](https://github.com/tidwall/buntdb) | 2265 | 93 | 2016-07-20 | 1 month ago | Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support. |
| [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) | 2035 | 149 | 2015-01-15 | 1 day ago | Sync your MySQL data into Elasticsearch automatically. |
| [prest](https://github.com/prest/prest) | 1962 | 82 | 2016-11-22 | 6 months ago | Serve a RESTful API from any PostgreSQL database. |
| [xo](https://github.com/xo/xo) | 1946 | 70 | 2016-02-05 | 2 days ago | Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. |
| [squirrel](https://github.com/Masterminds/squirrel) | 1938 | 33 | 2014-01-18 | 6 days ago | Go library that helps you build SQL queries. |
| [bigcache](https://github.com/allegro/bigcache) | 1879 | 64 | 2016-03-23 | 2 weeks ago | Efficient key/value cache for gigabytes of data. |
| [migrate](https://github.com/golang-migrate/migrate) | 1707 | 37 | 2018-01-19 | 2 days ago | Database migrations. CLI and Golang library. |
| [go-mysql](https://github.com/siddontang/go-mysql) | 1566 | 132 | 2014-02-21 | 1 day ago | Go toolset to handle MySQL protocol and replication. |
| [sql-migrate](https://github.com/rubenv/sql-migrate) | 1243 | 29 | 2014-09-09 | 2 weeks ago | Database migration tool. Allows embedding migrations into the application using go-bindata. |
| [cache2go](https://github.com/muesli/cache2go) | 813 | 57 | 2013-11-11 | 11 hours ago | In-memory key:value cache which supports automatic invalidation based on timeouts. |
| [diskv](https://github.com/peterbourgon/diskv) | 695 | 29 | 2012-03-22 | 6 months ago | Home-grown disk-backed key-value store. |
| [moss](https://github.com/couchbase/moss) | 684 | 79 | 2016-02-07 | 1 week ago | Moss is a simple LSM key-value storage engine written in 100% Go. |
| [gcache](https://github.com/bluele/gcache) | 666 | 29 | 2015-01-25 | 3 days ago | Cache library with support for expirable Cache, LFU, LRU and ARC. |
| [pop](https://github.com/gobuffalo/pop) | 584 | 24 | 2018-02-08 | 1 day ago | Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite. |
| [gendry](https://github.com/didi/gendry) | 573 | 46 | 2017-12-01 | 2 months ago | Non-invasive SQL builder and powerful data binder. |
| [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) | 551 | 45 | 2018-04-11 | 15 hours ago | CovenantSQL is a SQL database on blockchain. |
| [eliasdb](https://github.com/krotik/eliasdb) | 504 | 23 | 2016-08-13 | 2 weeks ago | Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language. |
| [goqu](https://github.com/doug-martin/goqu) | 492 | 25 | 2015-02-21 | 2 weeks ago | Idiomatic SQL builder and query library. |
| [dotsql](https://github.com/gchaincl/dotsql) | 407 | 21 | 2014-11-20 | 4 months ago | Go library that helps you keep sql files in one place and use them with ease. |
| [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) | 385 | 26 | 2015-12-11 | 6 months ago | Powerful data retrieval methods as well as DB-agnostic query building capabilities. |
| [levigo](https://github.com/jmhodges/levigo) | 348 | 24 | 2012-01-17 | 4 days ago | Levigo is a Go wrapper for LevelDB. |
| [fastcache](https://github.com/VictoriaMetrics/fastcache) | 348 | 12 | 2018-11-23 | 1 week ago | fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead. |
| [gormigrate](https://github.com/go-gormigrate/gormigrate) | 257 | 3 | 2016-08-31 | 4 weeks ago | Database schema migration helper for Gorm ORM. |
| [chproxy](https://github.com/Vertamedia/chproxy) | 241 | 20 | 2017-09-18 | 1 month ago | HTTP proxy for ClickHouse database. |
| [pudge](https://github.com/recoilme/pudge) | 181 | 6 | 2018-11-20 | 5 days ago | Fast and simple  key/value store written using Go's standard library. |
| [piladb](https://github.com/fern4lvarez/piladb) | 167 | 8 | 2015-09-09 | 11 months ago | Lightweight RESTful database engine based on stack data structures. |
| [sqrl](https://github.com/elgris/sqrl) | 141 | 5 | 2014-06-25 | 1 month ago | SQL query builder, fork of Squirrel with improved performance. |
| [scaneo](https://github.com/variadico/scaneo) | 141 | 8 | 2015-04-30 | 11 months ago | Generate Go code to convert database rows into arbitrary structs. |
| [myreplication](https://github.com/2tvenom/myreplication) | 129 | 17 | 2015-02-05 | 5 months ago | MySql binary log replication listener. Supports statement and row based replication. |
| [vasto](https://github.com/chrislusf/vasto) | 120 | 16 | 2018-01-16 | 4 months ago | A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption. |
| [goose](https://github.com/steinbacher/goose) | 103 | 4 | 2016-03-05 | 2 years ago | Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts. |
| [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) | 99 | 6 | 2017-04-29 | 1 week ago | Collects small insterts and sends big requests to ClickHouse servers. |
| [darwin](https://github.com/GuiaBolso/darwin) | 77 | 2 | 2016-04-05 | 3 months ago | Database schema evolution library for Go. |
| [slowpoke](https://github.com/recoilme/slowpoke) | 74 | 6 | 2018-02-19 | 2 months ago | Key-value store with persistence. |
| [igor](https://github.com/galeone/igor) | 73 | 6 | 2016-03-10 | 8 months ago | Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax. |
| [godbal](https://github.com/xujiajun/godbal) | 48 | 3 | 2018-02-28 | 1 month ago | Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily. |
| [octillery](https://github.com/knocknote/octillery) | 43 | 14 | 2018-11-26 | 1 month ago | Go package for sharding databases ( Supports every ORM or raw SQL ). |
| [couchcache](https://github.com/codingsince1985/couchcache) | 39 | 3 | 2015-04-05 | 1 year ago | RESTful caching micro-service backed by Couchbase server. |
| [golang-scribble](https://github.com/nanobox-io/golang-scribble) | 32 | 3 | 2018-06-22 | 4 months ago | Tiny flat file JSON store. |
| [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) | 25 | 5 | 2017-12-18 | 1 year ago | BigCache with clustering support and individual item expiration. |
| [dbbench](https://github.com/sj14/dbbench) | 24 | 2 | 2018-11-24 | 1 month ago | Database benchmarking tool with support for several databases and scripts. |
| [pravasan](https://github.com/pravasan/pravasan) | 23 | 6 | 2015-01-04 | 2 months ago | Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc. |
| [prep](https://github.com/hexdigest/prep) | 22 | 2 | 2017-12-12 | 1 year ago | Use prepared SQL statements without changing your code. |
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
| [mysql](https://github.com/go-sql-driver/mysql) | 7048 | 392 | 2012-12-10 | 11 hours ago | MySQL driver for Go. |
| [pq](https://github.com/lib/pq) | 4678 | 157 | 2012-03-13 | 20 hours ago | Pure Go Postgres driver for database/sql. |
| [go-sqlite3](https://github.com/mattn/go-sqlite3) | 3091 | 125 | 2011-11-11 | 2 weeks ago | SQLite3 driver for go that uses database/sql. |
| [pgx](https://github.com/jackc/pgx) | 1650 | 66 | 2013-03-31 | 4 days ago | PostgreSQL driver supporting features beyond those exposed by database/sql. |
| [go-mssqldb](https://github.com/denisenkom/go-mssqldb) | 910 | 59 | 2013-12-16 | 1 day ago | Microsoft MSSQL driver for Go. |
| [go-oci8](https://github.com/mattn/go-oci8) | 360 | 34 | 2012-02-29 | 3 weeks ago | Oracle driver for go that uses database/sql. |
| [goracle](https://github.com/go-goracle/goracle) | 181 | 28 | 2015-03-25 | 1 week ago | Oracle driver for Go, using the ODPI-C driver. |
| [firebirdsql](https://github.com/nakagami/firebirdsql) | 96 | 14 | 2013-08-27 | 2 days ago | Firebird RDBMS SQL driver for Go. |
| [go-adodb](https://github.com/mattn/go-adodb) | 84 | 12 | 2011-11-14 | 6 days ago | Microsoft ActiveX Object DataBase driver for go that uses database/sql. |
| [gofreetds](https://github.com/minus5/gofreetds) | 84 | 21 | 2012-12-07 | 1 week ago | Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org). |
| [calcite-avatica-go](https://github.com/apache/calcite-avatica-go) | 26 | 16 | 2017-08-08 | 1 month ago | Apache Avatica/Phoenix SQL driver for database/sql. |
| [bgc](https://github.com/viant/bgc) | 10 | 9 | 2016-06-14 | 1 month ago | Datastore Connectivity for BigQuery for go. |

### NoSQL Databases
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [cayley](https://github.com/cayleygraph/cayley) | 12255 | 633 | 2014-06-06 | 1 week ago | Graph database with support for multiple backends. |
| [redigo](https://github.com/gomodule/redigo) | 5546 | 280 | 2012-04-14 | 3 days ago | Redigo is a Go client for the Redis database. |
| [redis](https://github.com/go-redis/redis) | 5148 | 200 | 2012-07-25 | 1 day ago | Redis client for Golang. |
| [bleve](https://github.com/blevesearch/bleve) | 5002 | 232 | 2014-04-18 | 14 hours ago | Modern text indexing library for go. |
| [riot](https://github.com/go-ego/riot) | 4301 | 171 | 2017-06-21 | 1 month ago | Go Open Source, Distributed, Simple and efficient Search Engine. |
| [elastic](https://github.com/olivere/elastic) | 3511 | 160 | 2012-12-07 | 1 day ago | Elasticsearch client for Go. |
| [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) | 2027 | 137 | 2017-02-09 | 4 days ago | Official MongoDB driver for the Go language. |
| [mgo](https://github.com/globalsign/mgo) | 1428 | 77 | 2017-04-13 | 4 days ago | MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms. |
| [rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1418 | 45 | 2013-09-12 | 1 month ago | Go language driver for RethinkDB. |
| [gomemcache](https://github.com/bradfitz/gomemcache) | 1011 | 51 | 2011-06-29 | 1 week ago | memcache client library for the Go programming language. |
| [elastigo](https://github.com/mattbaird/elastigo) | 937 | 49 | 2012-10-12 | 3 weeks ago | Elasticsearch client library. |
| [neoism](https://github.com/jmcvetta/neoism) | 350 | 23 | 2012-07-12 | 2 weeks ago | Neo4j client for Golang. |
| [elasticsql](https://github.com/cch123/elasticsql) | 293 | 21 | 2016-08-24 | 1 week ago | Convert sql to elasticsearch dsl in Go. |
| [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) | 288 | 39 | 2014-07-26 | 1 week ago | Aerospike client in Go language. |
| [go-couchbase](https://github.com/couchbase/go-couchbase) | 286 | 26 | 2012-01-20 | 2 days ago | Couchbase client in Go. |
| [gocb](https://github.com/couchbase/gocb) | 281 | 70 | 2015-01-16 | 1 week ago | Official Couchbase Go SDK. |
| [redeo](https://github.com/bsm/redeo) | 240 | 24 | 2014-03-06 | 3 months ago | Redis-protocol compatible TCP servers/services. |
| [cachego](https://github.com/fabiorphp/cachego) | 103 | 6 | 2016-10-06 | 9 months ago | Golang Cache component for multiple drivers. |
| [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) | 72 | 6 | 2011-06-05 | 8 months ago | Neo4j REST Client in golang. |
| [go-rejson](https://github.com/nitishm/go-rejson) | 67 | 4 | 2018-04-23 | 1 month ago | Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease. |
| [arangolite](https://github.com/solher/arangolite) | 64 | 5 | 2015-10-05 | 1 year ago | Lightweight golang driver for ArangoDB. |
| [dynago](https://github.com/underarmour/dynago) | 63 | 131 | 2015-05-18 | 1 year ago | Dynago is a principle of least surprise client for DynamoDB. |
| [skizze](https://github.com/seiflotfy/skizze) | 59 | 6 | 2016-01-17 | 2 years ago | probabilistic data-structures service and storage. |
| [go-couchdb](https://github.com/fjl/go-couchdb) | 51 | 6 | 2013-10-28 | 6 months ago | Yet another CouchDB HTTP API wrapper for Go. |
| [gokv](https://github.com/philippgille/gokv) | 45 | 4 | 2018-10-09 | 1 day ago | Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more) |
| [goforestdb](https://github.com/couchbase/goforestdb) | 29 | 44 | 2014-05-14 | 2 years ago | Go bindings for ForestDB. |
| [go-pilosa](https://github.com/pilosa/go-pilosa) | 27 | 13 | 2016-10-01 | 3 days ago | Go client library for Pilosa. |
| [goriak](https://github.com/zegl/goriak) | 25 | 2 | 2016-10-06 | 2 months ago | Go language driver for Riak KV. |
| [goes](https://github.com/OwnLocal/goes) | 25 | 34 | 2015-12-29 | 2 years ago | Library to interact with Elasticsearch. |
| [neo4j](https://github.com/cihangir/neo4j) | 24 | 3 | 2013-05-18 | 3 years ago | Neo4j Rest API Bindings for Golang. |
| [xredis](https://github.com/shomali11/xredis) | 9 | 1 | 2017-06-14 | 9 months ago | Typesafe, customizable, clean & easy to use Redis client. |
| [dsc](https://github.com/viant/dsc) | 7 | 17 | 2016-06-14 | 1 week ago | Datastore connectivity for SQL, NoSQL, structured files. |
| [asc](https://github.com/viant/asc) | 4 | 10 | 2016-06-14 | 3 months ago | Datastore Connectivity for Aerospike for go. |
| [godscache](https://github.com/defcronyke/godscache) | 4 | 2 | 2018-05-09 | 3 weeks ago | A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached. |

## Date and Time
        
*Libraries for working with dates and times.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [now](https://github.com/jinzhu/now) | 1981 | 51 | 2013-11-18 | 2 weeks ago | Now is a time toolkit for golang. |
| [dateparse](https://github.com/araddon/dateparse) | 804 | 14 | 2014-04-21 | 1 week ago | Parse date's without knowing format in advance. |
| [carbon](https://github.com/uniplaces/carbon) | 306 | 38 | 2016-08-03 | 1 month ago | Simple Time extension with a lot of util methods, ported from PHP Carbon library. |
| [durafmt](https://github.com/hako/durafmt) | 208 | 4 | 2016-05-21 | 6 months ago | Time duration formatting library for Go. |
| [timeutil](https://github.com/leekchan/timeutil) | 156 | 6 | 2015-08-02 | 4 weeks ago | Useful extensions (Timedelta, Strftime, ...) to the golang's time package. |
| [timespan](https://github.com/SaidinWoT/timespan) | 60 | 5 | 2014-10-07 | 1 month ago | For interacting with intervals of time, defined as a start time and a duration. |
| [iso8601](https://github.com/relvacode/iso8601) | 60 | 2 | 2017-04-25 | 2 months ago | Efficiently parse ISO8601 date-times without regex. |
| [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) | 48 | 3 | 2016-02-01 | 1 month ago | The implementation of the Persian (Solar Hijri) Calendar in Go (golang). |
| [date](https://github.com/rickb777/date) | 23 | 2 | 2015-11-24 | 3 months ago | Augments Time for working with dates, date ranges, time spans, periods, and time-of-day. |
| [goweek](https://github.com/grsmv/goweek) | 18 | 3 | 2015-11-12 | 2 years ago | Library for working with week entity in golang. |
| [feiertage](https://github.com/wlbr/feiertage) | 17 | 1 | 2015-11-04 | 2 weeks ago | Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving... |
| [go-sunrise](https://github.com/nathan-osman/go-sunrise) | 10 | 3 | 2017-06-16 | 1 year ago | Calculate the sunrise and sunset times for a given location. |
| [kair](https://github.com/GuilhermeCaruso/kair) | 9 | 1 | 2018-10-03 | 1 week ago | Date and Time - Golang Formatting Library. |
| [nulltime](https://github.com/kirillDanshin/nulltime) | 9 | 1 | 2016-03-06 | 1 year ago | Nullable `time.Time`. |
| [tuesday](https://github.com/osteele/tuesday) | 7 | 1 | 2017-08-11 | 1 year ago | Ruby-compatible Strftime function. |
| [strftime](https://github.com/awoodbeck/strftime) | 5 | 1 | 2018-02-10 | 1 year ago | C99-compatible strftime formatter. |

## Distributed Systems
        
*Packages that help with building Distributed Systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [etcd](https://github.com/etcd-io/etcd) | 23253 | 1228 | 2013-07-07 | 1 day ago | Go implementation of the Raft consensus protocol, by CoreOS. |
| [kit](https://github.com/go-kit/kit) | 12778 | 641 | 2015-02-03 | 21 hours ago | Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc. |
| [grpc-go](https://github.com/grpc/grpc-go) | 7585 | 437 | 2014-12-09 | 17 hours ago | The Go language implementation of gRPC. HTTP/2 based RPC. |
| [jaeger](https://github.com/jaegertracing/jaeger) | 7338 | 277 | 2016-04-16 | 11 hours ago | A distributed tracing system. |
| [micro](https://github.com/micro/micro) | 5604 | 290 | 2015-01-17 | 4 days ago | Pluggable microservice toolkit and distributed systems platform. |
| [gnatsd](https://github.com/nats-io/gnatsd) | 5314 | 326 | 2012-10-30 | 1 day ago | Lightweight, high performance messaging system for microservices, IoT, and cloud native systems. |
| [rpcx](https://github.com/smallnest/rpcx) | 3237 | 271 | 2016-05-18 | 3 days ago | Distributed pluggable RPC service framework like alibaba Dubbo. |
| [tendermint](https://github.com/tendermint/tendermint) | 2713 | 240 | 2014-05-15 | 12 hours ago | High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols. |
| [torrent](https://github.com/anacrolix/torrent) | 2613 | 121 | 2015-01-09 | 1 day ago | BitTorrent client package. |
| [raft](https://github.com/hashicorp/raft) | 2499 | 177 | 2013-11-05 | 3 days ago | Golang implementation of the Raft consensus protocol, by HashiCorp. |
| [glow](https://github.com/chrislusf/glow) | 2354 | 141 | 2015-06-14 | 4 months ago | Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go. |
| [gleam](https://github.com/chrislusf/gleam) | 1834 | 138 | 2016-08-26 | 2 months ago | Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed. |
| [emitter](https://github.com/emitter-io/emitter) | 1655 | 81 | 2016-10-29 | 1 day ago | High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love. |
| [krakend](https://github.com/devopsfaith/krakend) | 1279 | 67 | 2016-11-05 | 2 days ago | Ultra performant API Gateway framework with middlewares. |
| [hprose-golang](https://github.com/hprose/hprose-golang) | 919 | 89 | 2014-02-14 | 9 months ago | Very newbility RPC Library, support 25+ languages now. |
| [ringpop-go](https://github.com/uber/ringpop-go) | 531 | 2310 | 2015-06-06 | 1 week ago | Scalable, fault-tolerant application-layer sharding for Go applications. |
| [gorpc](https://github.com/valyala/gorpc) | 522 | 26 | 2014-11-21 | 1 year ago | Simple, fast and scalable RPC library for high load. |
| [go-health](https://github.com/InVisionApp/go-health) | 437 | 20 | 2017-11-30 | 3 months ago | Library for enabling asynchronous dependency health checks in your service. |
| [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) | 355 | 18 | 2015-10-09 | 1 week ago | Video streaming torrent client. |
| [sleuth](https://github.com/ursiform/sleuth) | 292 | 7 | 2016-04-23 | 11 months ago | Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)). |
| [digota](https://github.com/digota/digota) | 270 | 24 | 2017-08-14 | 4 months ago | grpc ecommerce microservice. |
| [go-jump](https://github.com/dgryski/go-jump) | 231 | 13 | 2014-06-16 | 1 year ago | Port of Google's "Jump" Consistent Hash function. |
| [consistent](https://github.com/buraksezer/consistent) | 152 | 9 | 2018-03-25 | 1 week ago | Consistent hashing with bounded loads. |
| [redis-lock](https://github.com/bsm/redis-lock) | 108 | 6 | 2014-10-10 | 3 months ago | Simplified distributed locking implementation using Redis. |
| [jsonrpc](https://github.com/osamingo/jsonrpc) | 104 | 6 | 2016-10-28 | 2 months ago | The jsonrpc package helps implement of JSON-RPC 2.0. |
| [jsonrpc](https://github.com/ybbus/jsonrpc) | 79 | 6 | 2016-11-10 | 3 months ago | JSON-RPC 2.0 HTTP client implementation. |
| [celeriac.v1](https://github.com/svcavallar/celeriac.v1) | 47 | 3 | 2015-10-10 | 2 months ago | Library for adding support for interacting and monitoring Celery workers, tasks and events in Go. |
| [doublejump](https://github.com/edwingeng/doublejump) | 27 | 3 | 2018-06-27 | 2 months ago | A revamped Google's jump consistent hash. |
| [drmaa](https://github.com/dgruber/drmaa) | 23 | 2 | 2013-03-17 | 9 months ago | Job submission library for cluster schedulers based on the DRMAA standard. |
| [flowgraph](https://github.com/vectaport/flowgraph) | 8 | 1 | 2018-08-30 | 3 days ago | flow-based programming package. |
| [dynatomic](https://github.com/tylfin/dynatomic) | 6 | 0 | 2019-02-09 | 1 week ago | A library for using DynamoDB as an atomic counter. |
| [outboxer](https://github.com/italolelis/outboxer) | 0 | 0 | 2019-02-01 | 1 week ago | Outboxer is a go library that implements the outbox pattern. |

## Email
        
*Libraries and tools that implement email creation and sending.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [MailHog](https://github.com/mailhog/MailHog) | 4512 | 125 | 2014-04-17 | 3 days ago | Email and SMTP testing with web and API interface. |
| [gomail](https://github.com/go-gomail/gomail) | 2136 | 65 | 2014-10-15 | 2 weeks ago | Gomail is a very simple and powerful package to send emails. |
| [hermes](https://github.com/matcornic/hermes) | 1476 | 24 | 2017-03-26 | 6 days ago | Golang package that generates clean, responsive HTML e-mails. |
| [email](https://github.com/jordan-wright/email) | 1007 | 35 | 2013-12-13 | 2 weeks ago | A robust and flexible email library for Go. |
| [go-imap](https://github.com/emersion/go-imap) | 626 | 31 | 2016-04-27 | 1 day ago | IMAP library for clients and servers. |
| [sendgrid-go](https://github.com/sendgrid/sendgrid-go) | 468 | 181 | 2013-09-12 | 1 day ago | SendGrid's Go library for sending email. |
| [hectane](https://github.com/hectane/hectane) | 160 | 12 | 2015-08-28 | 10 months ago | Lightweight SMTP client providing an HTTP API. |
| [douceur](https://github.com/aymerick/douceur) | 142 | 4 | 2015-04-09 | 11 months ago | CSS inliner for your HTML emails. |
| [go-message](https://github.com/emersion/go-message) | 79 | 6 | 2016-12-31 | 1 month ago | Streaming library for the Internet Message Format and mail messages. |
| [smtp](https://github.com/mailhog/smtp) | 49 | 9 | 2014-12-25 | 8 months ago | SMTP server protocol state machine. |
| [go-dkim](https://github.com/toorop/go-dkim) | 47 | 3 | 2015-04-29 | 1 year ago | DKIM library, to sign & verify email. |

## Embeddable Scripting Languages
        
*Embedding other languages inside your go code.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [otto](https://github.com/robertkrimen/otto) | 4406 | 187 | 2012-10-06 | 1 week ago | JavaScript interpreter written in Go. |
| [gopher-lua](https://github.com/yuin/gopher-lua) | 2610 | 136 | 2015-02-15 | 3 weeks ago | Lua 5.1 VM and compiler written in Go. |
| [go-lua](https://github.com/Shopify/go-lua) | 1541 | 253 | 2013-12-21 | 3 months ago | Port of the Lua 5.2 VM to pure Go. |
| [tengo](https://github.com/d5/tengo) | 1056 | 25 | 2019-01-09 | 15 hours ago | Bytecode compiled script language for Go. |
| [anko](https://github.com/mattn/anko) | 842 | 47 | 2014-03-28 | 6 days ago | Scriptable interpreter written in Go. |
| [go-python](https://github.com/sbinet/go-python) | 808 | 43 | 2012-07-09 | 1 month ago | naive go bindings to the CPython C-API. |
| [go-duktape](https://github.com/olebedev/go-duktape) | 628 | 23 | 2015-01-08 | 2 weeks ago | Duktape JavaScript engine bindings for Go. |
| [go-php](https://github.com/deuill/go-php) | 609 | 39 | 2015-09-18 | 4 months ago | PHP bindings for Go. |
| [golua](https://github.com/aarzilli/golua) | 422 | 35 | 2010-12-07 | 6 months ago | Go bindings for Lua C API. |
| [gisp](https://github.com/jcla1/gisp) | 415 | 20 | 2014-01-11 | 1 year ago | Simple LISP in Go. |
| [agora](https://github.com/mna/agora) | 317 | 22 | 2013-06-19 | 4 years ago | Dynamically typed, embeddable programming language in Go. |
| [expr](https://github.com/antonmedv/expr) | 296 | 16 | 2018-07-14 | 1 month ago | an engine that can evaluate expressions. |
| [gval](https://github.com/PaesslerAG/gval) | 85 | 12 | 2017-09-27 | 1 week ago | A highly customizable expression language written in Go. |
| [purl](https://github.com/ian-kent/purl) | 25 | 2 | 2014-11-30 | 4 years ago | Perl 5.18.2 embedded in Go. |
| [binder](https://github.com/alexeyco/binder) | 24 | 2 | 2017-04-03 | 7 months ago | Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua). |
| [ngaro](https://github.com/db47h/ngaro) | 16 | 1 | 2016-08-09 | 9 months ago | Embeddable Ngaro VM implementation enabling scripting in Retro. |

## Error Handling
        
*Libraries for handling errors.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [errors](https://github.com/pkg/errors) | 4164 | 92 | 2015-12-27 | 6 days ago | Package that provides simple error handling primitives. |
| [go-multierror](https://github.com/hashicorp/go-multierror) | 558 | 82 | 2014-12-16 | 2 months ago | Go (golang) package for representing a list of errors as a single error. |
| [errorx](https://github.com/joomcode/errorx) | 507 | 37 | 2018-08-17 | 2 weeks ago | A feature rich error package with stack traces, composition of errors and more. |
| [tracerr](https://github.com/ztrue/tracerr) | 174 | 6 | 2019-02-07 | 2 weeks ago | Golang errors with stack trace and source fragments. |
| [werr](https://github.com/txgruppi/werr) | 10 | 0 | 2015-08-04 | 3 years ago | Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called. |

## Files
        
*Libraries for  handling files and file systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [afero](https://github.com/spf13/afero) | 1941 | 84 | 2014-10-28 | 2 days ago | FileSystem Abstraction System for Go. |
| [pdfcpu](https://github.com/hhrutter/pdfcpu) | 796 | 23 | 2017-06-19 | 1 day ago | PDF processor. |
| [notify](https://github.com/rjeczalik/notify) | 452 | 26 | 2014-09-09 | 2 weeks ago | File system event notification library with simple API, similar to os/signal. |
| [opc](https://github.com/qmuntal/opc) | 49 | 1 | 2018-11-06 | 6 days ago | Load Open Packaging Conventions (OPC) files for Go. |
| [skywalker](https://github.com/dixonwille/skywalker) | 41 | 3 | 2017-08-02 | 1 year ago | Package to allow one to concurrently go through a filesystem with ease. |
| [go-csv-tag](https://github.com/artonge/go-csv-tag) | 39 | 1 | 2017-06-18 | 4 months ago | Load csv file using tag. |
| [tarfs](https://github.com/posener/tarfs) | 31 | 2 | 2017-03-11 | 1 year ago | Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files. |
| [go-gtfs](https://github.com/artonge/go-gtfs) | 12 | 2 | 2017-07-09 | 1 month ago | Load gtfs files in go. |
| [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) | 11 | 1 | 2018-10-16 | 4 months ago | Copy files for humans. |

## Financial
        
*Packages for accounting and finance.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [decimal](https://github.com/shopspring/decimal) | 1336 | 56 | 2015-02-26 | 1 day ago | Arbitrary-precision fixed-point decimal numbers. |
| [go-money](https://github.com/Rhymond/go-money) | 552 | 15 | 2017-03-21 | 1 month ago | Implementation of Fowler's Money pattern. |
| [go-finance](https://github.com/FlashBoys/go-finance) | 537 | 28 | 2016-02-28 | 1 year ago | Comprehensive financial markets data in Go. |
| [accounting](https://github.com/leekchan/accounting) | 442 | 12 | 2015-08-10 | 1 month ago | money and currency formatting for golang. |
| [techan](https://github.com/sdcoffey/techan) | 108 | 17 | 2017-03-08 | 2 months ago | Technical analysis library with advanced market analysis and trading strategies. |
| [vat](https://github.com/dannyvankooten/vat) | 54 | 1 | 2016-06-19 | 5 months ago | VAT number validation & EU VAT rates. |
| [ofxgo](https://github.com/aclindsa/ofxgo) | 53 | 5 | 2015-11-08 | 2 days ago | Query OFX servers and/or parse the responses (with example command-line client). |
| [transaction](https://github.com/claygod/transaction) | 44 | 8 | 2017-10-11 | 6 months ago | Embedded transactional database of accounts, running in multithreaded mode. |
| [go-finance](https://github.com/alpeb/go-finance) | 32 | 2 | 2017-06-01 | 10 months ago | Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations. |
| [orderbook](https://github.com/i25959341/orderbook) | 22 | 3 | 2018-04-25 | 3 days ago | Matching Engine for Limit Order Book in Golang. |

## Forms
        
*Libraries for working with forms.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [nosurf](https://github.com/justinas/nosurf) | 920 | 35 | 2013-08-23 | 1 month ago | CSRF protection middleware for Go. |
| [binding](https://github.com/mholt/binding) | 730 | 30 | 2014-05-21 | 11 months ago | Binds form and JSON data from net/http Request to struct. |
| [csrf](https://github.com/gorilla/csrf) | 369 | 19 | 2015-08-03 | 2 months ago | CSRF protection for Go web applications & services. |
| [form](https://github.com/go-playground/form) | 324 | 8 | 2016-05-26 | 4 weeks ago | Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support. |
| [conform](https://github.com/leebenson/conform) | 163 | 5 | 2016-01-06 | 8 months ago | Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags. |
| [formam](https://github.com/monoculum/formam) | 111 | 2 | 2014-10-25 | 6 months ago | decode form's values into a struct. |
| [forms](https://github.com/albrow/forms) | 96 | 6 | 2014-08-08 | 1 year ago | Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files. |
| [bind](https://github.com/robfig/bind) | 23 | 2 | 2014-08-06 | 4 years ago | Bind form data to any Go values. |

## Functional
        
*Packages to support functional programming in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [go-underscore](https://github.com/tobyhede/go-underscore) | 1010 | 26 | 2014-07-02 | 2 weeks ago | Useful collection of helpfully functional Go collection utilities. |
| [fpGo](https://github.com/TeaEntityLab/fpGo) | 70 | 1 | 2018-05-24 | 7 months ago | Monad, Functional Programming features for Golang. |
| [fuego](https://github.com/seborama/fuego) | 14 | 1 | 2018-11-06 | 3 days ago | Functional Experiment in Go. |

## Game Development
        
*Awesome game development libraries.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [leaf](https://github.com/name5566/leaf) | 2723 | 302 | 2014-08-04 | 2 months ago | Lightweight game server framework. |
| [pixel](https://github.com/faiface/pixel) | 2113 | 92 | 2016-11-19 | 1 week ago | Hand-crafted 2D game library in Go. |
| [ebiten](https://github.com/hajimehoshi/ebiten) | 1406 | 64 | 2013-06-16 | 11 hours ago | dead simple 2D game library in Go. |
| [go-sdl2](https://github.com/veandco/go-sdl2) | 1052 | 42 | 2013-06-06 | 1 day ago | Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/). |
| [gonet](https://github.com/xtaci/gonet) | 998 | 125 | 2013-04-11 | 1 year ago | Game server skeleton implemented with golang. |
| [termloop](https://github.com/JoelOtter/termloop) | 977 | 35 | 2015-05-24 | 2 days ago | Terminal-based game engine for Go, built on top of Termbox. |
| [engo](https://github.com/EngoEngine/engo) | 953 | 45 | 2014-11-12 | 3 days ago | Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm. |
| [goworld](https://github.com/xiaonanln/goworld) | 947 | 99 | 2017-06-03 | 2 months ago | Scalable game server engine, featuring space-entity framework and hot-swapping. |
| [nano](https://github.com/lonng/nano) | 844 | 53 | 2017-08-02 | 1 month ago | Lightweight, facility, high performance golang based game server framework. |
| [engine](https://github.com/g3n/engine) | 593 | 58 | 2017-03-08 | 1 week ago | Go 3D Game Engine. |
| [oak](https://github.com/oakmound/oak) | 573 | 34 | 2017-07-16 | 1 day ago | Pure Go game engine. |
| [engine](https://github.com/azul3d/engine) | 406 | 23 | 2016-02-29 | 8 months ago | 3D game engine written in Go. |
| [raylib-go](https://github.com/gen2brain/raylib-go) | 340 | 18 | 2017-01-27 | 3 months ago | Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming. |
| [go-astar](https://github.com/beefsack/go-astar) | 307 | 10 | 2014-05-28 | 11 months ago | Go implementation of the A\* path finding algorithm. |
| [GarageEngine](https://github.com/vova616/GarageEngine) | 303 | 31 | 2012-08-07 | 5 years ago | 2d game engine written in Go working on OpenGL. |
| [pitaya](https://github.com/topfreegames/pitaya) | 202 | 26 | 2018-03-20 | 4 days ago | Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. |
| [go3d](https://github.com/ungerik/go3d) | 150 | 10 | 2011-06-27 | 3 days ago | Performance oriented 2D/3D math package for Go. |
| [glop](https://github.com/runningwild/glop) | 76 | 4 | 2011-04-21 | 3 years ago | Glop (Game Library Of Power) is a fairly simple cross-platform game library. |
| [go-collada](https://github.com/GlenKelley/go-collada) | 12 | 3 | 2013-09-19 | 5 years ago | Go package for working with the Collada file format. |

## Generation and Generics
        
*Tools to enhance the language with features like generics via code generation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [go-linq](https://github.com/ahmetb/go-linq) | 1661 | 72 | 2013-12-19 | 3 months ago | .NET LINQ-like query methods for Go. |
| [jennifer](https://github.com/dave/jennifer) | 1138 | 21 | 2016-12-05 | 1 month ago | Generate arbitrary Go code without templates. |
| [gen](https://github.com/clipperhouse/gen) | 977 | 36 | 2013-10-14 | 8 months ago | Code generation tool for ‘generics’-like functionality. |
| [goderive](https://github.com/awalterschulze/goderive) | 670 | 22 | 2017-02-11 | 1 week ago | Derives functions from input types. |
| [gowrap](https://github.com/hexdigest/gowrap) | 205 | 6 | 2018-09-15 | 6 days ago | Generate decorators for Go interfaces using simple templates. |
| [interfaces](https://github.com/rjeczalik/interfaces) | 160 | 5 | 2015-12-06 | 2 months ago | Command line tool for generating interface definitions. |
| [pkgreflect](https://github.com/ungerik/pkgreflect) | 78 | 5 | 2012-09-03 | 1 year ago | Go preprocessor for package scoped reflection. |
| [go-enum](https://github.com/abice/go-enum) | 66 | 4 | 2017-08-11 | 2 weeks ago | Code generation for enums from code comments. |
| [efaceconv](https://github.com/t0pep0/efaceconv) | 40 | 4 | 2016-11-18 | 1 year ago | Code generation tool for high performance conversion from interface{} to immutable type without allocations. |
| [gotype](https://github.com/wzshiming/gotype) | 18 | 3 | 2017-12-05 | 1 week ago | Golang source code parsing, usage like reflect package. |

## Geographic
        
*Geographic tools and servers*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [tile38](https://github.com/tidwall/tile38) | 5318 | 176 | 2016-03-05 | 3 days ago | Geolocation DB with spatial index and realtime geofencing. |
| [geo](https://github.com/golang/geo) | 803 | 73 | 2014-12-04 | 1 week ago | S2 geometry library in Go. |
| [geocache](https://github.com/melihmucuk/geocache) | 95 | 5 | 2016-06-21 | 2 years ago | In-memory cache that is suitable for geolocation based applications. |
| [osm](https://github.com/paulmach/osm) | 41 | 5 | 2016-02-02 | 2 months ago | Library for reading, writing and working with OpenStreetMap data and APIs. |
| [geoserver](https://github.com/hishamkaram/geoserver) | 19 | 2 | 2018-03-27 | 3 months ago | geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API. |
| [pbf](https://github.com/maguro/pbf) | 10 | 1 | 2017-09-19 | 4 months ago | OpenStreetMap PBF golang encoder/decoder. |
| [gismanager](https://github.com/hishamkaram/gismanager) | 9 | 1 | 2018-09-29 | 4 months ago | Publish Your GIS Data(Vector Data) to PostGIS and Geoserver. |

## Go Compilers
        
*Tools for compiling Go to other languages.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gopherjs](https://github.com/gopherjs/gopherjs) | 7992 | 249 | 2013-08-28 | 1 day ago | Compiler from Go to JavaScript. |
| [llgo](https://github.com/go-llvm/llgo) | 952 | 61 | 2011-11-05 | 4 years ago | LLVM-based compiler for Go. |
| [tardisgo](https://github.com/tardisgo/tardisgo) | 382 | 27 | 2014-01-08 | 2 years ago | Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler. |
| [c4go](https://github.com/Konstantin8105/c4go) | 63 | 7 | 2018-03-28 | 12 hours ago | Transpile C code to Go code. |
| [f4go](https://github.com/Konstantin8105/f4go) | 9 | 2 | 2018-07-09 | 2 months ago | Transpile FORTRAN 77 code to Go code. |

## Goroutines
        
*Tools for managing and working with Goroutines.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [goworker](https://github.com/benmanns/goworker) | 2135 | 70 | 2013-07-23 | 3 weeks ago | goworker is a Go-based background worker. |
| [ants](https://github.com/panjf2000/ants) | 1361 | 54 | 2018-05-19 | 1 week ago | A high-performance goroutine pool for golang. |
| [tunny](https://github.com/Jeffail/tunny) | 1136 | 51 | 2014-04-03 | 3 months ago | Goroutine pool for golang. |
| [grpool](https://github.com/ivpusic/grpool) | 460 | 29 | 2015-07-22 | 1 month ago | Lightweight Goroutine pool. |
| [pool](https://github.com/go-playground/pool) | 452 | 12 | 2015-10-29 | 2 years ago | Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation. |
| [go-floc](https://github.com/workanator/go-floc) | 164 | 8 | 2017-07-03 | 1 year ago | Orchestrate goroutines with ease. |
| [go-flow](https://github.com/kamildrazkiewicz/go-flow) | 94 | 7 | 2016-09-25 | 1 year ago | Control goroutines execution order. |
| [workerpool](https://github.com/gammazero/workerpool) | 80 | 5 | 2016-05-17 | 2 months ago | Goroutine pool that limits the concurrency of task execution, not the number of tasks queued. |
| [semaphore](https://github.com/marusama/semaphore) | 63 | 4 | 2017-11-22 | 1 month ago | Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations). |
| [semaphore](https://github.com/kamilsk/semaphore) | 60 | 1 | 2016-10-08 | 3 days ago | Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context. |
| [GoSlaves](https://github.com/dgrr/GoSlaves) | 60 | 4 | 2017-09-18 | 7 months ago | Simple and Asynchronous Goroutine pool library. |
| [gpool](https://github.com/sherifabdlnaby/gpool) | 46 | 1 | 2018-12-03 | 1 day ago | manages a resizeable pool of context-aware goroutines to bound concurrency. |
| [worker-pool](https://github.com/vardius/worker-pool) | 31 | 3 | 2017-10-04 | 3 weeks ago | goworker is a Go simple async worker pool. |
| [parallel-fn](https://github.com/rafaeljesus/parallel-fn) | 24 | 3 | 2017-06-18 | 1 year ago | Run functions in parallel. |
| [cyclicbarrier](https://github.com/marusama/cyclicbarrier) | 23 | 3 | 2018-01-11 | 4 months ago | CyclicBarrier for golang. |
| [async](https://github.com/StudioSol/async) | 14 | 9 | 2017-07-01 | 5 months ago | A safe way to execute functions asynchronously, recovering them in case of panic. |
| [threadpool](https://github.com/shettyh/threadpool) | 12 | 2 | 2017-09-07 | 1 year ago | Golang threadpool implementation. |
| [breaker](https://github.com/kamilsk/breaker) | 6 | 1 | 2019-02-15 | 1 day ago | 🚧 Flexible mechanism to make your code breakable. |
| [artifex](https://github.com/borderstech/artifex) | 6 | 2 | 2018-11-01 | 4 months ago | Simple in-memory job queue for Golang using worker-based dispatching. |
| [stl](https://github.com/ssgreg/stl) | 5 | 1 | 2018-06-19 | 5 months ago | Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism. |
| [go-trylock](https://github.com/subchen/go-trylock) | 4 | 1 | 2018-04-26 | 9 months ago | TryLock support on read-write lock for Golang. |

## GUI
        
*Interaction*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [ui](https://github.com/andlabs/ui) | 6455 | 355 | 2014-02-18 | 3 weeks ago | Platform-native GUI library for Go. Cross platform. |
| [qt](https://github.com/therecipe/qt) | 5264 | 273 | 2014-11-19 | 1 month ago | Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi). |
| [robotgo](https://github.com/go-vgo/robotgo) | 3917 | 176 | 2016-09-27 | 3 days ago | Go Native cross-platform GUI system automation. Control the mouse, keyboard and other. |
| [webview](https://github.com/zserge/webview) | 3897 | 173 | 2017-08-19 | 1 day ago | Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux). |
| [walk](https://github.com/lxn/walk) | 3275 | 239 | 2010-09-16 | 4 days ago | Windows application library kit for Go. |
| [app](https://github.com/maxence-charriere/app) | 2750 | 106 | 2016-10-12 | 3 days ago | Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. |
| [go-astilectron](https://github.com/asticode/go-astilectron) | 2358 | 112 | 2017-04-22 | 1 month ago | Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron). |
| [go-sciter](https://github.com/sciter-sdk/go-sciter) | 1297 | 113 | 2015-10-15 | 4 months ago | Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform. |
| [fyne](https://github.com/fyne-io/fyne) | 773 | 41 | 2018-02-05 | 17 hours ago | Cross platform native GUIs designed for Go, rendered using EFL. Supports: Linux, macOS, Windows. |
| [systray](https://github.com/getlantern/systray) | 635 | 38 | 2014-11-12 | 2 weeks ago | Cross platform Go library to place an icon and menu in the notification area. |
| [gotk3](https://github.com/gotk3/gotk3) | 621 | 43 | 2015-08-13 | 15 hours ago | Go bindings for GTK3. |
| [gosx-notifier](https://github.com/deckarep/gosx-notifier) | 476 | 15 | 2013-11-25 | 1 year ago | OSX Desktop Notifications library for Go. |
| [gowd](https://github.com/dtylman/gowd) | 172 | 19 | 2017-03-29 | 4 months ago | Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform. |
| [trayhost](https://github.com/shurcooL/trayhost) | 143 | 4 | 2014-04-25 | 4 months ago | Cross-platform Go library to place an icon in the host operating system's taskbar. |

## Hardware
        
*Libraries, tools, and tutorials for interacting with hardware.*

## Images
        
*Libraries for manipulating images.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [ln](https://github.com/fogleman/ln) | 2378 | 89 | 2016-01-10 | 1 year ago | 3D line art rendering in Go. |
| [imaginary](https://github.com/h2non/imaginary) | 2312 | 70 | 2015-03-05 | 1 week ago | Fast and simple HTTP microservice for image resizing. |
| [imaging](https://github.com/disintegration/imaging) | 2190 | 63 | 2012-12-07 | 1 month ago | Simple Go image processing package. |
| [resize](https://github.com/nfnt/resize) | 2024 | 73 | 2012-08-03 | 1 year ago | Image resizing for Go with common interpolation methods. |
| [gocv](https://github.com/hybridgroup/gocv) | 1995 | 102 | 2017-09-19 | 1 week ago | Go package for computer vision using OpenCV 3.3+. |
| [bild](https://github.com/anthonynsimon/bild) | 1931 | 55 | 2016-08-01 | 5 months ago | Collection of image processing algorithms in pure Go. |
| [pt](https://github.com/fogleman/pt) | 1717 | 62 | 2015-01-24 | 1 year ago | Path tracing engine written in Go. |
| [gg](https://github.com/fogleman/gg) | 1711 | 68 | 2016-02-19 | 1 week ago | 2D rendering in pure Go. |
| [svgo](https://github.com/ajstarks/svgo) | 1222 | 47 | 2010-03-06 | 5 months ago | Go Language Library for SVG generation. |
| [smartcrop](https://github.com/muesli/smartcrop) | 1203 | 30 | 2014-04-08 | 4 months ago | Finds good crops for arbitrary images and crop sizes. |
| [gift](https://github.com/disintegration/gift) | 1156 | 49 | 2014-07-13 | 5 months ago | Package of image processing filters. |
| [go-opencv](https://github.com/go-opencv/go-opencv) | 1027 | 68 | 2013-12-09 | 1 month ago | Go bindings for OpenCV. |
| [geopattern](https://github.com/pravj/geopattern) | 982 | 20 | 2014-10-23 | 1 month ago | Create beautiful generative image patterns from a string. |
| [picfit](https://github.com/thoas/picfit) | 967 | 41 | 2014-12-07 | 1 month ago | An image resizing server written in Go. |
| [imagick](https://github.com/gographics/imagick) | 884 | 54 | 2013-05-01 | 1 month ago | Go binding to ImageMagick's MagickWand C API. |
| [bimg](https://github.com/h2non/bimg) | 692 | 35 | 2015-03-17 | 1 day ago | Small package for fast and efficient image processing using libvips. |
| [mort](https://github.com/aldor007/mort) | 342 | 17 | 2017-11-19 | 2 weeks ago | Storage and image processing server written in Go. |
| [stegify](https://github.com/DimitarPetrov/stegify) | 281 | 13 | 2018-11-30 | 2 weeks ago | Go tool for LSB steganography, capable of hiding any file within an image. |
| [govatar](https://github.com/o1egl/govatar) | 280 | 5 | 2016-01-18 | 10 months ago | Library and CMD tool for generating funny avatars. |
| [go-nude](https://github.com/koyachi/go-nude) | 272 | 17 | 2014-05-02 | 3 months ago | Nudity detection with Go. |
| [image2ascii](https://github.com/qeesung/image2ascii) | 238 | 7 | 2018-10-20 | 3 months ago | Convert image to ASCII. |
| [goimagehash](https://github.com/corona10/goimagehash) | 167 | 9 | 2017-07-29 | 3 weeks ago | Go Perceptual image hashing package. |
| [rez](https://github.com/bamiaux/rez) | 156 | 9 | 2014-01-17 | 1 year ago | Image resizing in pure Go and SIMD. |
| [img](https://github.com/hawx/img) | 127 | 5 | 2012-07-29 | 3 years ago | Selection of image manipulation tools. |
| [go-cairo](https://github.com/ungerik/go-cairo) | 83 | 6 | 2012-08-23 | 5 months ago | Go binding for the cairo graphics library. |
| [go-gd](https://github.com/bolknote/go-gd) | 47 | 4 | 2011-05-12 | 10 months ago | Go binding for GD library. |
| [mergi](https://github.com/noelyahan/mergi) | 41 | 2 | 2018-09-24 | 1 week ago | Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate). |
| [go-webcolors](https://github.com/jyotiska/go-webcolors) | 23 | 1 | 2014-04-24 | 3 years ago | Port of webcolors library from Python to Go. |
| [tga](https://github.com/ftrvxmtrx/tga) | 21 | 2 | 2012-10-08 | 3 years ago | Package tga is a TARGA image format decoder/encoder. |
| [cameron](https://github.com/aofei/cameron) | 17 | 1 | 2018-05-06 | 6 days ago | An avatar generator for Go. |
| [steganography](https://github.com/auyer/steganography) | 12 | 3 | 2018-05-22 | 4 months ago | Pure Go Library for LSB steganography. |
| [mpo](https://github.com/donatj/mpo) | 5 | 1 | 2015-04-15 | 1 month ago | Decoder and conversion tool for MPO 3D Photos. |

## IoT
        
*Libraries for programming devices of the IoT.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gobot](https://github.com/hybridgroup/gobot) | 5201 | 298 | 2013-09-21 | 2 weeks ago | Gobot is a framework for robotics, physical computing, and the Internet of Things. |
| [flogo](https://github.com/TIBCOSoftware/flogo) | 954 | 113 | 2016-07-10 | 4 days ago | Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. |
| [gatt](https://github.com/paypal/gatt) | 769 | 57 | 2014-04-23 | 3 months ago | Gatt is a Go package for building Bluetooth Low Energy peripherals. |
| [mainflux](https://github.com/mainflux/mainflux) | 363 | 53 | 2015-07-07 | 16 hours ago | Industrial IoT Messaging and Device Management Server. |
| [devices](https://github.com/goiot/devices) | 221 | 17 | 2016-05-30 | 2 years ago | Suite of libraries for IoT devices, experimental for x/exp/io. |
| [sensorbee](https://github.com/sensorbee/sensorbee) | 165 | 19 | 2016-02-19 | 1 year ago | Lightweight stream processing engine for IoT. |
| [connectordb](https://github.com/connectordb/connectordb) | 152 | 17 | 2015-01-17 | 4 days ago | Open-Source Platform for Quantified Self & IoT. |
| [huego](https://github.com/amimof/huego) | 83 | 1 | 2017-05-16 | 6 days ago | An extensive Philips Hue client library for Go. |
| [eywa](https://github.com/xcodersun/eywa) | 30 | 8 | 2016-02-21 | 1 year ago | Project Eywa is essentially a connection manager that keeps track of connected devices. |
| [iot](https://github.com/vaelen/iot) | 28 | 4 | 2018-03-08 | 10 months ago | IoT is a simple framework for implementing a Google IoT Core device. |

## Job Scheduler
        
*Libraries for scheduling jobs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gron](https://github.com/roylee0704/gron) | 552 | 16 | 2016-06-04 | 1 week ago | Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly. |
| [jobrunner](https://github.com/bamzi/jobrunner) | 517 | 17 | 2015-10-21 | 2 years ago | Smart and featureful cron job scheduler with job queuing and live monitoring built in. |
| [jobs](https://github.com/albrow/jobs) | 438 | 18 | 2015-02-10 | 8 months ago | Persistent and flexible background jobs library. |
| [scheduler](https://github.com/carlescere/scheduler) | 266 | 16 | 2015-02-04 | 8 months ago | Cronjobs scheduling made easy. |
| [go-cron](https://github.com/rk/go-cron) | 169 | 9 | 2011-04-15 | 3 years ago | Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. |
| [clockwork](https://github.com/whiteShtef/clockwork) | 50 | 1 | 2018-04-24 | 2 weeks ago | Simple and intuitive job scheduling library in Go. |
| [leprechaun](https://github.com/kilgaloon/leprechaun) | 23 | 4 | 2018-04-08 | 4 days ago | Job scheduler that supports webhooks, crons and classic scheduling. |

## JSON
        
*Libraries for working with JSON.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [gjson](https://github.com/tidwall/gjson) | 4086 | 112 | 2016-08-11 | 2 weeks ago | Get a JSON value with one line of code. |
| [gojson](https://github.com/ChimeraCoder/gojson) | 1920 | 45 | 2012-12-28 | 2 weeks ago | Automatically generate Go (golang) struct definitions from example JSON. |
| [gojq](https://github.com/elgs/gojq) | 132 | 5 | 2015-12-30 | 9 months ago | JSON query in Golang. |
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
| [logrus](https://github.com/sirupsen/logrus) | 9986 | 267 | 2013-10-17 | 1 day ago | Structured logger for Go. |
| [zap](https://github.com/uber-go/zap) | 6140 | 199 | 2016-02-19 | 1 day ago | Fast, structured, leveled logging in Go. |
| [go-spew](https://github.com/davecgh/go-spew) | 2978 | 59 | 2013-01-09 | 1 month ago | Implements a deep pretty printer for Go data structures to aid in debugging. |
| [glog](https://github.com/golang/glog) | 2146 | 84 | 2013-07-16 | 3 months ago | Leveled execution logs for Go. |
| [zerolog](https://github.com/rs/zerolog) | 1769 | 47 | 2017-05-12 | 2 days ago | Zero-allocation JSON logger. |
| [tail](https://github.com/hpcloud/tail) | 1360 | 101 | 2013-02-05 | 5 days ago | Go package striving to emulate the features of the BSD tail program. |
| [seelog](https://github.com/cihub/seelog) | 1271 | 89 | 2011-11-17 | 18 hours ago | Logging functionality with flexible dispatching, filtering, and formatting. |
| [lumberjack](https://github.com/natefinch/lumberjack) | 1188 | 37 | 2014-06-14 | 3 days ago | Simple rolling logger, implements io.WriteCloser. |
| [log15](https://github.com/inconshreveable/log15) | 843 | 23 | 2014-05-20 | 4 months ago | Simple, powerful logging for Go. |
| [log](https://github.com/apex/log) | 676 | 35 | 2015-12-22 | 4 months ago | Structured logging package for Go. |
| [logxi](https://github.com/mgutz/logxi) | 328 | 10 | 2015-03-02 | 1 year ago | 12-factor app logger that is fast and makes you happy. |
| [onelog](https://github.com/francoispqt/onelog) | 303 | 9 | 2018-05-06 | 16 hours ago | Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenario. Also, it is one of the logger with the lowest allocation. |
| [log](https://github.com/go-playground/log) | 257 | 10 | 2016-02-08 | 2 weeks ago | Simple, configurable and scalable Structured Logging for Go. |
| [logutils](https://github.com/hashicorp/logutils) | 239 | 88 | 2013-10-09 | 6 months ago | Utilities for slightly better logging in Go (Golang) extending the standard logger. |
| [go-logger](https://github.com/apsdehal/go-logger) | 212 | 8 | 2014-09-26 | 5 months ago | Simple logger of Go Programs, with level handlers. |
| [logger](https://github.com/azer/logger) | 130 | 4 | 2014-09-30 | 5 months ago | Minimalistic logging library for Go. |
| [xlog](https://github.com/rs/xlog) | 126 | 7 | 2015-10-22 | 11 months ago | Structured logger for `net/context` aware HTTP handlers with flexible dispatching. |
| [ozzo-log](https://github.com/go-ozzo/ozzo-log) | 102 | 11 | 2015-10-23 | 11 months ago | High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail). |
| [logvoyage](https://github.com/firstrow/logvoyage) | 83 | 5 | 2015-03-29 | 1 year ago | Full-featured logging saas written in golang. |
| [log](https://github.com/alexcesaro/log) | 44 | 4 | 2014-04-19 | 3 years ago | Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs. |
| [gologger](https://github.com/sadlil/gologger) | 38 | 6 | 2015-09-02 | 1 year ago | Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch. |
| [go-log](https://github.com/ian-kent/go-log) | 33 | 2 | 2014-05-02 | 11 months ago | Log4j implementation in Go. |
| [logex](https://github.com/chzyer/logex) | 32 | 7 | 2014-10-10 | 1 year ago | Golang log lib, supports tracking and level, wrap by standard log lib. |
| [glg](https://github.com/kpango/glg) | 31 | 3 | 2017-06-21 | 15 hours ago | glg is simple and fast leveled logging library for Go. |
| [gone](https://github.com/One-com/gone) | 24 | 7 | 2016-09-05 | 3 weeks ago | Golang packages for writing small daemons and servers. |
| [logrusly](https://github.com/sebest/logrusly) | 24 | 5 | 2014-09-12 | 11 months ago | [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/). |
| [log](https://github.com/teris-io/log) | 21 | 1 | 2017-10-29 | 1 year ago | Structured log interface for Go cleanly separates logging facade from its implementation. |
| [go-log](https://github.com/siddontang/go-log) | 21 | 3 | 2014-05-18 | 1 week ago | Log lib supports level and multi handlers. |
| [go-cronowriter](https://github.com/utahta/go-cronowriter) | 17 | 1 | 2017-02-04 | 11 hours ago | Simple writer that rotate log files automatically based on current date and time, like cronolog. |
| [mlog](https://github.com/jbrodriguez/mlog) | 16 | 2 | 2014-10-20 | 7 months ago | Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output. |
| [distillog](https://github.com/amoghe/distillog) | 15 | 1 | 2015-10-13 | 7 months ago | distilled levelled logging (think of it as stdlib + log levels). |
| [gomol](https://github.com/aphistic/gomol) | 12 | 2 | 2015-08-30 | 9 months ago | Multiple-output, structured logging for Go with extensible logging outputs. |
| [go-log](https://github.com/subchen/go-log) | 8 | 1 | 2017-05-07 | 9 months ago | Simple and configurable Logging in Go, with level, formatters and writers. |
| [logdump](https://github.com/ewwwwwqm/logdump) | 7 | 2 | 2017-01-13 | 11 months ago | Package for multi-level logging. |
| [xlog](https://github.com/xfxdev/xlog) | 5 | 1 | 2016-05-06 | 1 month ago | Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format. |
| [glo](https://github.com/lajosbencz/glo) | 4 | 1 | 2019-01-20 | 1 month ago | PHP Monolog inspired logging facility with identical severity levels. |
| [logmatic](https://github.com/borderstech/logmatic) | 4 | 1 | 2018-11-07 | 1 month ago | Colorized logger for Golang with dynamic log level configuration. |
| [logo](https://github.com/mbndr/logo) | 4 | 1 | 2017-02-08 | 1 year ago | Golang logger to different configurable writers. |

## Machine Learning
        
*Libraries for Machine Learning.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [golearn](https://github.com/sjwhitworth/golearn) | 6307 | 429 | 2013-12-26 | 3 weeks ago | General Machine Learning library for Go. |
| [gorgonia](https://github.com/gorgonia/gorgonia) | 2445 | 162 | 2016-09-15 | 1 day ago | graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. |
| [tfgo](https://github.com/galeone/tfgo) | 1071 | 48 | 2017-05-23 | 2 weeks ago | Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python. |
| [goml](https://github.com/cdipaolo/goml) | 962 | 73 | 2015-06-27 | 2 years ago | On-line Machine Learning in Go. |
| [gosseract](https://github.com/otiai10/gosseract) | 745 | 31 | 2013-10-11 | 1 week ago | Go package for OCR (Optical Character Recognition), by using Tesseract C++ library. |
| [CloudForest](https://github.com/ryanbressler/CloudForest) | 628 | 43 | 2012-10-23 | 3 months ago | Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go. |
| [bayesian](https://github.com/jbrukh/bayesian) | 596 | 30 | 2011-11-23 | 2 weeks ago | Naive Bayesian Classification for Golang. |
| [eaopt](https://github.com/MaxHalford/eaopt) | 579 | 29 | 2016-01-31 | 1 week ago | An evolutionary optimization library. |
| [gorse](https://github.com/zhenghaoz/gorse) | 338 | 17 | 2018-08-14 | 22 hours ago | A High Performance Recommender System Package based on Collaborative Filtering for Go. |
| [gobrain](https://github.com/goml/gobrain) | 336 | 23 | 2014-04-29 | 2 months ago | Neural Networks written in go. |
| [regommend](https://github.com/muesli/regommend) | 233 | 14 | 2014-02-06 | 3 months ago | Recommendation & collaborative filtering engine. |
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
| [sarama](https://github.com/Shopify/sarama) | 3943 | 362 | 2013-07-06 | 12 hours ago | Go library for Apache Kafka. |
| [centrifugo](https://github.com/centrifugal/centrifugo) | 3351 | 174 | 2015-04-01 | 4 days ago | Real-time messaging (Websockets or SockJS) server in Go. |
| [gorush](https://github.com/appleboy/gorush) | 3276 | 151 | 2016-03-22 | 4 days ago | Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm). |
| [machinery](https://github.com/RichardKnop/machinery) | 2948 | 131 | 2015-04-06 | 15 hours ago | Asynchronous task queue/job queue based on distributed message passing. |
| [go-socket.io](https://github.com/googollee/go-socket.io) | 2624 | 122 | 2013-07-13 | 1 week ago | socket.io library for golang, a realtime application framework. |
| [go-nats](https://github.com/nats-io/go-nats) | 2130 | 158 | 2012-08-15 | 18 hours ago | Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library. |
| [apns2](https://github.com/sideshow/apns2) | 1932 | 72 | 2016-01-05 | 11 hours ago | HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps. |
| [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) | 1771 | 239 | 2013-12-27 | 1 year ago | gopush-cluster is a go push server cluster. |
| [benthos](https://github.com/Jeffail/benthos) | 1727 | 54 | 2016-03-22 | 13 hours ago | A message streaming bridge between a range of protocols. |
| [mangos](https://github.com/nanomsg/mangos) | 1521 | 89 | 2014-10-25 | 20 hours ago | Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability. |
| [melody](https://github.com/olahol/melody) | 1329 | 43 | 2015-05-14 | 4 months ago | Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling. |
| [go-nsq](https://github.com/nsqio/go-nsq) | 1321 | 61 | 2013-08-29 | 1 week ago | the official Go package for NSQ. |
| [uniqush-push](https://github.com/uniqush/uniqush-push) | 1056 | 71 | 2011-08-29 | 4 months ago | Redis backed unified push service for server-side notifications to mobile devices. |
| [mercure](https://github.com/dunglas/mercure) | 892 | 42 | 2018-07-14 | 3 days ago | Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events). |
| [zmq4](https://github.com/pebbe/zmq4) | 728 | 48 | 2013-10-18 | 1 month ago | Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2). |
| [gollum](https://github.com/trivago/gollum) | 650 | 35 | 2015-06-21 | 3 weeks ago | A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations. |
| [EventBus](https://github.com/asaskevich/EventBus) | 481 | 27 | 2014-12-20 | 1 month ago | The lightweight event bus with async compatibility. |
| [golongpoll](https://github.com/jcuga/golongpoll) | 397 | 21 | 2015-11-02 | 2 weeks ago | HTTP longpoll server library that makes web pub-sub simple. |
| [dbus](https://github.com/godbus/dbus) | 312 | 15 | 2014-03-28 | 1 week ago | Native Go bindings for D-Bus. |
| [glue](https://github.com/desertbit/glue) | 299 | 13 | 2015-06-07 | 1 year ago | Robust Go and Javascript Socket Library (Alternative to Socket.io). |
| [emitter](https://github.com/olebedev/emitter) | 281 | 10 | 2015-11-11 | 1 month ago | Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins. |
| [pubsub](https://github.com/cskr/pubsub) | 243 | 8 | 2012-04-01 | 3 weeks ago | Simple pubsub package for go. |
| [guble](https://github.com/smancke/guble) | 132 | 13 | 2015-11-16 | 1 year ago | Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence. |
| [oplog](https://github.com/dailymotion/oplog) | 94 | 97 | 2014-11-06 | 3 years ago | Generic oplog/replication system for REST APIs. |
| [drone-line](https://github.com/appleboy/drone-line) | 57 | 5 | 2016-09-13 | 3 months ago | Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI. |
| [RapidMQ](https://github.com/sybrexsys/RapidMQ) | 52 | 3 | 2016-10-05 | 1 year ago | RapidMQ is a lightweight and reliable library for managing of the local messages queue. |
| [rabtap](https://github.com/jandelgado/rabtap) | 52 | 6 | 2017-11-11 | 3 days ago | RabbitMQ swiss army knife cli app. |
| [rabbus](https://github.com/rafaeljesus/rabbus) | 52 | 8 | 2017-05-07 | 1 month ago | A tiny wrapper over amqp exchanges and queues. |
| [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) | 46 | 7 | 2017-01-16 | 1 year ago | A tiny wrapper around NSQ topic and channel. |
| [go-notify](https://github.com/TheCreeper/go-notify) | 44 | 2 | 2015-03-02 | 3 weeks ago | Native implementation of the freedesktop notification spec. |
| [goose](https://github.com/ian-kent/goose) | 35 | 1 | 2014-12-21 | 4 years ago | Server Sent Events in Go. |
| [message-bus](https://github.com/vardius/message-bus) | 30 | 3 | 2017-10-04 | 3 weeks ago | messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD. |
| [event](https://github.com/agoalofalife/event) | 20 | 2 | 2017-07-02 | 1 year ago | Implementation of the pattern observer. |
| [hub](https://github.com/leandro-lugaresi/hub) | 19 | 1 | 2018-04-14 | 10 months ago | A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges. |
| [go-vitotrol](https://github.com/maxatome/go-vitotrol) | 11 | 6 | 2016-11-04 | 4 months ago | Client library to Viessmann Vitotrol web service. |
| [gaurun-client](https://github.com/osamingo/gaurun-client) | 7 | 1 | 2017-06-29 | 7 months ago | Gaurun Client written in Go. |
| [jazz](https://github.com/socifi/jazz) | 2 | 4 | 2018-10-22 | 3 weeks ago | A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages. |

## Microsoft Office
        

### Microsoft Excel
        
*Libraries for working with Microsoft Excel.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [excelize](https://github.com/360EntSecGroup-Skylar/excelize) | 3356 | 127 | 2016-08-29 | 6 days ago | Golang library for reading and writing Microsoft Excel™ (XLSX) files. |
| [xlsx](https://github.com/tealeg/xlsx) | 3014 | 175 | 2011-06-28 | 1 week ago | Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs. |
| [xlsx](https://github.com/plandem/xlsx) | 35 | 3 | 2017-08-27 | 1 month ago | Fast and safe way to read/update your existing Microsoft Excel files in Go programs. |
| [go-excel](https://github.com/szyhf/go-excel) | 32 | 2 | 2017-09-03 | 2 months ago | A simple and light reader to read a relate-db-like excel as a table. |
| [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) | 11 | 1 | 2017-03-13 | 7 months ago | Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files. |

## Miscellaneous
        

### Dependency Injection
        
*Libraries for working with dependency injection.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [dig](https://github.com/uber-go/dig) | 614 | 35 | 2017-03-22 | 1 week ago | A reflection based dependency injection toolkit for Go. |
| [alice](https://github.com/magic003/alice) | 31 | 1 | 2017-04-09 | 1 year ago | Additive dependency injection container for Golang. |
| [wire](https://github.com/Fs02/wire) | 13 | 1 | 2018-07-05 | 1 month ago | Strict Runtime Dependency Injection for Golang. |

### Strings
        
*These libraries were placed here because none of the other categories seemed to fit.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_commit | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:| ---------:|
| [archiver](https://github.com/mholt/archiver) | 2090 | 43 | 2016-04-09 | 3 days ago | Library and command for making and extracting .zip and .tar.gz archives. |
| [go-commons-pool](https://github.com/jolestar/go-commons-pool) | 592 | 43 | 2015-12-28 | 6 months ago | Generic object pool for Golang. |
| [xstrings](https://github.com/huandu/xstrings) | 550 | 24 | 2015-01-06 | 5 months ago | Collection of useful string functions ported from other languages. |
| [base64Captcha](https://github.com/mojocn/base64Captcha) | 485 | 33 | 2017-12-12 | 3 months ago | Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha. |
| [go-conv](https://github.com/cstockton/go-conv) | 336 | 10 | 2016-10-11 | 1 year ago | Package conv provides fast and intuitive conversions across Go types. |
| [banner](https://github.com/dimiro1/banner) | 205 | 4 | 2016-03-26 | 2 years ago | Add beautiful banners into your Go applications. |
| [antch](https://github.com/antchfx/antch) | 121 | 10 | 2017-09-28 | 7 months ago | A fast, powerful and extensible web crawling & scraping framework. |
| [battery](https://github.com/distatus/battery) | 120 | 2 | 2016-03-13 | 4 days ago | Cross-platform, normalized battery information library. |
| [ffmt](https://github.com/go-ffmt/ffmt) | 108 | 4 | 2015-02-14 | 3 months ago | Beautify data display for Humans. |
| [bitio](https://github.com/icza/bitio) | 80 | 5 | 2016-05-31 | 1 year ago | Highly optimized bit-level Reader and Writer for Go. |
| [strutil](https://github.com/ozgio/strutil) | 44 | 1 | 2018-08-16 | 6 months ago | String utilities. |
| [captcha](https://github.com/steambap/captcha) | 31 | 4 | 2017-09-12 | 4 days ago | Package captcha provides an easy to use, unopinionated API for captcha generation. |
| [browscap_go](https://github.com/digitalcrab/browscap_go) | 27 | 6 | 2014-09-18 | 1 month ago | GoLang Library for [Browser Capabilities Project](http://browscap.org/). |
| [datacounter](https://github.com/miolini/datacounter) | 26 | 1 | 2015-10-15 | 10 months ago | Go counters for readers/writer/http.ResponseWriter. |
| [autoflags](https://github.com/artyom/autoflags) | 21 | 2 | 2014-05-16 | 1 month ago | Go package to automatically define command line flags from struct fields. |
| [ghorg](https://github.com/gabrie30/ghorg) | 17 | 2 | 2018-03-29 | 2 months ago | Clone all repos from a GitHub org into a single directory. |
| [anagent](https://github.com/mudler/anagent) | 9 | 2 | 2017-12-30 | 6 months ago | Minimalistic, pluggable Golang evloop/timer handler with dependency-injection. |
| [avgRating](https://github.com/kirillDanshin/avgRating) | 8 | 1 | 2017-08-06 | 1 year ago | Calculate average score and rating based on Wilson Score Equation. |

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
        