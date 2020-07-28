# Awesome Go Info

go语言开源项目列表，项目分类及GitHub上的开源项目数据完全来自于[awesome-go](https://github.com/avelino/awesome-go) 的[README.md](https://github.com/avelino/awesome-go/blob/master/README.md)文件，通过调用GitHub的API获取仓库信息，展示项目的star数、watch数等，方便查看go语言开源项目的一些相关信息。

_该文件仅包含[awesome-go](https://github.com/avelino/awesome-go) [README.md](https://github.com/avelino/awesome-go/blob/master/README.md)文件中列出的在GitHub上开源的优秀项目，不罗列其它golang相关的网站_
_该文件中的GitHub仓库信息数据会在每天凌晨1点左右更新,当前数据更新于2020-07-29 00:58:22_

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
        - [Uncategorized](#uncategorized)
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
    - [WebAssembly](#webassembly)
    - [Files](#Files)
    - [File Handling](#file-handling)
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
    - [Style Guides](#style-guides)




# Awesome Go
        

## Audio and Music
        
*Libraries for manipulating audio.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [oto](https://github.com/hajimehoshi/oto) | 622 | 9 | 2017-05-04 | 3 days ago | A low-level library to play sound on multiple platforms. |
| [portaudio](https://github.com/gordonklaus/portaudio) | 384 | 17 | 2015-09-16 | 2 months ago | Go bindings for the PortAudio audio I/O library. |
| [music-theory](https://github.com/go-music-theory/music-theory) | 303 | 14 | 2016-03-17 | 5 days ago | Music theory models in Go. |
| [waveform](https://github.com/mdlayher/waveform) | 293 | 13 | 2014-09-13 | 4 months ago | Go package capable of generating waveform images from audio streams. |
| [portmidi](https://github.com/rakyll/portmidi) | 240 | 10 | 2013-11-10 | 3 months ago | Go bindings for PortMidi. |
| [id3v2](https://github.com/bogem/id3v2) | 161 | 5 | 2016-05-15 | 1 month ago | Fast and powerful ID3 decoding and encodung library written completely in Go. |
| [flac](https://github.com/mewkiz/flac) | 130 | 10 | 2012-11-01 | 7 months ago | Native Go FLAC encoder/decoder with support for FLAC streams. |
| [mix](https://github.com/go-mix/mix) | 123 | 3 | 2016-01-03 | 2 months ago | Sequence-based Go-native audio mixer for music apps. |
| [mp3](https://github.com/tcolgate/mp3) | 112 | 1 | 2015-02-26 | 3 years ago | Native Go MP3 decoder. |
| [go-sox](https://github.com/krig/go-sox) | 108 | 8 | 2013-10-08 | 2 years ago | libsox bindings for go. |
| [malgo](https://github.com/gen2brain/malgo) | 106 | 5 | 2017-11-09 | 6 days ago | Mini audio library. |
| [go-taglib](https://github.com/wtolson/go-taglib) | 74 | 6 | 2012-11-20 | 2 years ago | Go bindings for taglib. |
| [gaad](https://github.com/Comcast/gaad) | 72 | 10 | 2016-07-11 | 5 months ago | Native Go AAC bitstream parser. |
| [minimp3](https://github.com/tosone/minimp3) | 48 | 2 | 2018-01-26 | 4 months ago | Lightweight MP3 decoder library. |
| [EasyMIDI](https://github.com/algoGuy/EasyMIDI) | 31 | 3 | 2018-02-19 | 2 years ago | EasyMidi is a simple and reliable library for working with standard midi file (SMF). |
| [go_mediainfo](https://github.com/zhulik/go_mediainfo) | 31 | 1 | 2015-12-13 | 4 years ago | libmediainfo bindings for go. |
| [vorbis](https://github.com/mccoyst/vorbis) | 27 | 3 | 2013-07-12 | 1 year ago | "Native" Go Vorbis decoder (uses CGO, but has no dependencies). |
| [gosamplerate](https://github.com/dh1tw/gosamplerate) | 10 | 1 | 2016-11-20 | 2 weeks ago | libsamplerate bindings for go. |

## Authentication and OAuth
        
*Libraries for implementing authentications schemes.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [jwt-go](https://github.com/dgrijalva/jwt-go) | 8093 | 153 | 2012-04-18 | 1 day ago | Golang implementation of JSON Web Tokens (JWT). |
| [casbin](https://github.com/casbin/casbin) | 7317 | 205 | 2017-04-08 | 6 hours ago | Authorization library that supports access control models like ACL, RBAC, ABAC. |
| [oauth2](https://github.com/golang/oauth2) | 3009 | 102 | 2014-04-14 | 9 hours ago | Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support. |
| [goth](https://github.com/markbates/goth) | 2833 | 67 | 2014-10-14 | 1 month ago | provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box. |
| [authboss](https://github.com/volatiletech/authboss) | 2313 | 47 | 2015-01-03 | 1 week ago | Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time. |
| [loginsrv](https://github.com/tarent/loginsrv) | 1645 | 55 | 2016-11-11 | 1 month ago | JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam. |
| [osin](https://github.com/openshift/osin) | 1608 | 71 | 2013-09-10 | 1 month ago | Golang OAuth2 server library. |
| [go-jose](https://github.com/square/go-jose) | 1573 | 66 | 2014-11-14 | 6 days ago | Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs. |
| [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) | 1563 | 82 | 2015-11-01 | 2 months ago | Standalone, specification-compliant,  OAuth2 server written in Golang. |
| [gologin](https://github.com/dghubble/gologin) | 1238 | 28 | 2015-06-23 | 1 week ago | chainable handlers for login with OAuth1 and OAuth2 authentication providers. |
| [gorbac](https://github.com/mikespook/gorbac) | 1069 | 63 | 2013-12-26 | 1 year ago | provides a lightweight role-based access control (RBAC) implementation in Golang. |
| [scs](https://github.com/alexedwards/scs) | 739 | 19 | 2016-08-08 | 1 week ago | Session Manager for HTTP servers. |
| [permissions2](https://github.com/xyproto/permissions2) | 403 | 13 | 2014-11-19 | 4 months ago | Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt. |
| [paseto](https://github.com/o1egl/paseto) | 373 | 18 | 2018-01-23 | 1 month ago | Golang implementation of Platform-Agnostic Security Tokens (PASETO). |
| [jwt](https://github.com/pascaldekloe/jwt) | 205 | 11 | 2018-03-21 | 1 month ago | Lightweight JSON Web Token (JWT) library. |
| [httpauth](https://github.com/goji/httpauth) | 192 | 7 | 2014-05-26 | 4 years ago | HTTP Authentication middleware. |
| [jeff](https://github.com/abraithwaite/jeff) | 192 | 4 | 2018-08-02 | 3 months ago | Simple, flexible, secure and idiomatic web session management with pluggable backends. |
| [jwt-auth](https://github.com/adam-hanna/jwt-auth) | 185 | 12 | 2016-07-05 | 3 months ago | JWT middleware for Golang http servers with many configuration options. |
| [branca](https://github.com/hako/branca) | 127 | 7 | 2018-01-09 | 4 months ago | Golang implementation of Branca Tokens. |
| [session](https://github.com/icza/session) | 103 | 6 | 2016-02-08 | 1 year ago | Go session management for web servers (including support for Google App Engine - GAE). |
| [sessionup](https://github.com/swithek/sessionup) | 88 | 5 | 2019-07-23 | 4 months ago | Simple, yet effective HTTP session management and identification package. |
| [jwt](https://github.com/robbert229/jwt) | 81 | 6 | 2016-06-05 | 1 year ago | Clean and easy to use implementation of JSON Web Tokens (JWT). |
| [sjwt](https://github.com/brianvoe/sjwt) | 67 | 0 | 2019-06-20 | 10 months ago | Simple jwt generator and parser. |
| [rbac](https://github.com/zpatrick/rbac) | 60 | 3 | 2018-08-02 | 1 year ago | Minimalistic RBAC package for Go applications. |
| [go-guardian](https://github.com/shaj13/go-guardian) | 60 | 1 | 2020-05-14 | 2 weeks ago | Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication that supports LDAP, Basic, Bearer token and Certificate based authentication. |
| [sessions](https://github.com/adam-hanna/sessions) | 52 | 3 | 2017-04-29 | 3 months ago | Dead simple, highly performant, highly customizable sessions service for go http servers. |
| [securecookie](https://github.com/chmike/securecookie) | 39 | 6 | 2017-09-03 | 3 months ago | Efficient secure cookie encoding/decoding. |
| [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) | 8 | 2 | 2017-10-20 | 1 year ago | Go session management using the SessionGate Redis module. |
| [signedvalue](https://github.com/sashka/signedvalue) | 8 | 0 | 2018-01-06 | 10 months ago | Signed and timestamped strings compatible with [Tornado's](https://github.com/tornadoweb/tornado) `create_signed_value`, `decode_signed_value`, and therefore `set_secure_cookie` and `get_secure_cookie`. |
| [cookiestxt](https://github.com/mengzhuo/cookiestxt) | 4 | 1 | 2017-10-09 | 2 years ago | provides parser of cookies.txt file format. |
| [scope](https://github.com/SonicRoshan/scope) | 4 | 1 | 2019-09-23 | 6 months ago | Easily Manage OAuth2 Scopes In Go. |

## Bot Building
        
*Libraries for building and working with bots.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [olivia](https://github.com/olivia-ai/olivia) | 2613 | 64 | 2018-06-05 | 1 week ago | A chatbot built with an artificial neural network. |
| [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) | 2255 | 61 | 2015-06-25 | 1 day ago | Simple and clean Telegram bot client. |
| [telebot](https://github.com/tucnak/telebot) | 1386 | 41 | 2015-06-25 | 2 days ago | Telegram bot framework written in Go. |
| [bot](https://github.com/go-chat-bot/bot) | 591 | 40 | 2015-09-22 | 2 months ago | IRC, Slack & Telegram bot written in Go. |
| [slacker](https://github.com/shomali11/slacker) | 434 | 15 | 2017-05-20 | 6 hours ago | Easy to use framework to create Slack bots. |
| [kelp](https://github.com/stellar/kelp) | 345 | 38 | 2018-08-08 | 6 days ago | official trading and market-making bot for the [Stellar](https://www.stellar.org/) DEX. Works out-of-the-box, written in Golang, compatible with centralized exchanges and custom trading strategies. |
| [golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot) | 323 | 27 | 2017-05-14 | 1 month ago | A golang implementation of a console-based trading bot for cryptocurrency exchanges. |
| [tbot](https://github.com/yanzay/tbot) | 280 | 11 | 2015-09-11 | 2 months ago | Telegram bot server with API similar to net/http. |
| [tenyks](https://github.com/kyleterry/tenyks) | 167 | 15 | 2012-08-26 | 10 months ago | Service oriented IRC bot using Redis and JSON for messaging. |
| [go-sarah](https://github.com/oklahomer/go-sarah) | 165 | 7 | 2016-11-06 | 5 days ago | Framework to build bot for desired chat services including LINE, Slack, Gitter and more. |
| [hanu](https://github.com/sbstjn/hanu) | 126 | 5 | 2016-09-16 | 1 month ago | Framework for writing Slack bots. |
| [go-twitch-irc](https://github.com/gempir/go-twitch-irc) | 116 | 9 | 2017-03-23 | 1 month ago | Libary to write bots for twitch. |
| [go-tgbot](https://github.com/olebedev/go-tgbot) | 96 | 8 | 2016-12-11 | 2 years ago | Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware. |
| [margelet](https://github.com/zhulik/margelet) | 63 | 5 | 2015-11-21 | 3 years ago | Framework for building Telegram bots. |
| [slackscot](https://github.com/alexandre-normand/slackscot) | 32 | 1 | 2015-10-22 | 3 days ago | Another framework for building Slack bots. |
| [govkbot](https://github.com/nikepan/govkbot) | 29 | 3 | 2016-07-11 | 4 months ago | Simple Go [VK](https://vk.com) bot library. |
| [ephemeral-roles](https://github.com/ewohltman/ephemeral-roles) | 27 | 2 | 2017-12-19 | 1 day ago | A Discord bot for managing ephemeral roles based upon voice channel member presence. |
| [micha](https://github.com/onrik/micha) | 15 | 4 | 2016-04-14 | 5 months ago | Go Library for Telegram bot api. |

## Command Line
        
*Libraries for building Console Applications and Console User Interfaces.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [cobra](https://github.com/spf13/cobra) | 17929 | 324 | 2013-09-03 | 1 week ago | Commander for modern Go CLI interactions. |
| [cli](https://github.com/urfave/cli) | 14192 | 295 | 2013-07-13 | 5 days ago | Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli). |
| [termui](https://github.com/gizak/termui) | 10039 | 289 | 2015-02-03 | 1 week ago | Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib). |
| [gocui](https://github.com/jroimartin/gocui) | 6472 | 131 | 2014-01-04 | 3 weeks ago | Minimalist Go library aimed at creating Console User Interfaces. |
| [termbox-go](https://github.com/nsf/termbox-go) | 3819 | 101 | 2012-01-12 | 3 months ago | Termbox is a library for creating cross-platform text-based interfaces. |
| [go-prompt](https://github.com/c-bata/go-prompt) | 3100 | 47 | 2017-08-14 | 1 week ago | Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit). |
| [kingpin](https://github.com/alecthomas/kingpin) | 2891 | 56 | 2014-05-14 | 5 days ago | Command line and flag parser supporting sub commands. |
| [dnote](https://github.com/dnote/dnote) | 1789 | 31 | 2017-03-30 | 1 week ago | A simple command line notebook with multi-device sync. |
| [go-flags](https://github.com/jessevdk/go-flags) | 1719 | 29 | 2012-08-31 | 1 week ago | go command line option parser. |
| [uiprogress](https://github.com/gosuri/uiprogress) | 1677 | 35 | 2015-11-17 | 1 year ago | Flexible library to render progress bars in terminal applications. |
| [readline](https://github.com/chzyer/readline) | 1522 | 39 | 2015-09-20 | 8 months ago | Pure golang implementation that provides most features in GNU-Readline under MIT license. |
| [asciigraph](https://github.com/guptarohit/asciigraph) | 1393 | 27 | 2018-06-17 | 4 weeks ago | Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. |
| [docopt.go](https://github.com/docopt/docopt.go) | 1248 | 34 | 2013-08-25 | 10 months ago | Command-line arguments parser that will make you smile. |
| [termdash](https://github.com/mum4k/termdash) | 1228 | 25 | 2018-03-24 | 1 week ago | Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui). |
| [pflag](https://github.com/spf13/pflag) | 1184 | 28 | 2013-08-30 | 2 weeks ago | Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. |
| [progressbar](https://github.com/schollz/progressbar) | 1169 | 22 | 2017-10-26 | 2 weeks ago | Basic thread-safe progress bar that works in every OS. |
| [cli](https://github.com/mitchellh/cli) | 1165 | 22 | 2013-11-03 | 3 months ago | Go library for implementing command-line interfaces. |
| [uilive](https://github.com/gosuri/uilive) | 1165 | 18 | 2015-11-16 | 2 months ago | Library for updating terminal output in realtime. |
| [go-arg](https://github.com/alexflint/go-arg) | 929 | 14 | 2015-11-01 | 3 weeks ago | Struct-based argument parsing in Go. |
| [mpb](https://github.com/vbauerster/mpb) | 918 | 18 | 2016-12-14 | 1 week ago | Multi progress bar for terminal applications. |
| [gcli](https://github.com/tcnksm/gcli) | 898 | 26 | 2014-06-19 | 2 years ago | The easy way to start building Golang command line applications. |
| [aurora](https://github.com/logrusorgru/aurora) | 876 | 8 | 2016-11-06 | 1 week ago | ANSI terminal colors that supports fmt.Printf/Sprintf. |
| [complete](https://github.com/posener/complete) | 702 | 15 | 2017-05-05 | 4 weeks ago | Write bash completions in Go + Go command bash completion. |
| [liner](https://github.com/peterh/liner) | 698 | 23 | 2012-08-15 | 2 months ago | Go readline-like library for command-line interfaces. |
| [mow.cli](https://github.com/jawher/mow.cli) | 683 | 19 | 2014-12-18 | 1 week ago | Go library for building CLI applications with sophisticated flag and argument parsing and validation. |
| [flaggy](https://github.com/integrii/flaggy) | 669 | 18 | 2018-03-05 | 1 week ago | A robust and idiomatic flags package with excellent subcommand support. |
| [uitable](https://github.com/gosuri/uitable) | 570 | 15 | 2015-11-13 | 8 months ago | Library to improve readability in terminal apps using tabular data. |
| [cli](https://github.com/mkideal/cli) | 533 | 22 | 2016-02-26 | 2 weeks ago | Feature-rich and easy to use command-line package based on golang struct tags. |
| [ops](https://github.com/nanovms/ops) | 464 | 26 | 2018-09-10 | 4 days ago | Unikernel Builder/Orchestrator. |
| [go-colorable](https://github.com/mattn/go-colorable) | 448 | 19 | 2014-07-30 | 1 month ago | Colorable writer for windows. |
| [go-isatty](https://github.com/mattn/go-isatty) | 434 | 9 | 2014-04-01 | 6 months ago | isatty for golang. |
| [color](https://github.com/gookit/color) | 404 | 9 | 2018-07-01 | 4 days ago | Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows. |
| [chalk](https://github.com/ttacon/chalk) | 340 | 9 | 2014-07-18 | 11 months ago | Intuitive package for prettifying terminal/console output. |
| [tabby](https://github.com/cheynewallace/tabby) | 268 | 2 | 2018-12-17 | 1 year ago | A tiny library for super simple Golang tables. |
| [simpletable](https://github.com/alexeyco/simpletable) | 241 | 4 | 2017-03-29 | 2 weeks ago | Simple tables in terminal with Go. |
| [argparse](https://github.com/akamensky/argparse) | 211 | 10 | 2017-11-24 | 1 month ago | Command line argument parser inspired by Python's argparse module. |
| [go-colortext](https://github.com/daviddengcn/go-colortext) | 205 | 9 | 2013-01-23 | 4 months ago | Go library for color output in terminals. |
| [commandeer](https://github.com/jaffee/commandeer) | 116 | 7 | 2017-10-12 | 21 hours ago | Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags. |
| [sflags](https://github.com/octago/sflags) | 116 | 5 | 2016-12-04 | 2 months ago | Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries. |
| [wmenu](https://github.com/dixonwille/wmenu) | 115 | 2 | 2016-04-20 | 1 month ago | Easy to use menu structure for cli applications that prompts users to make choices. |
| [clif](https://github.com/ukautz/clif) | 105 | 2 | 2015-05-30 | 1 year ago | Small command line interface framework. |
| [flag](https://github.com/cosiner/flag) | 103 | 5 | 2016-10-05 | 3 months ago | Simple but powerful command line option parsing library for Go supporting subcommand. |
| [flag](https://github.com/zhuah/flag) | 102 | 5 | 2016-10-05 | 1 year ago | Simple but powerful command line option parsing library for Go supporting subcommand. |
| [yacspin](https://github.com/theckman/yacspin) | 84 | 4 | 2019-12-29 | 1 month ago | Yet Another CLi Spinner package, for working with terminal spinners. |
| [job](https://github.com/liujianping/job) | 79 | 1 | 2019-04-09 | 4 weeks ago | JOB, make your short-term command as a long-term job. |
| [cfmt](https://github.com/mingrammer/cfmt) | 76 | 3 | 2018-03-15 | 1 year ago | Contextual fmt inspired by bootstrap color classes. |
| [cli](https://github.com/teris-io/cli) | 75 | 1 | 2017-05-24 | 1 year ago | Simple and complete API for building command line interfaces in Go. |
| [1build](https://github.com/gopinath-langote/1build) | 70 | 6 | 2019-04-23 | 1 month ago | Command line tool to frictionlessly manage project-specific commands. |
| [env](https://github.com/codingconcepts/env) | 51 | 1 | 2017-06-14 | 1 year ago | Tag-based environment configuration for structs. |
| [gocmd](https://github.com/devfacet/gocmd) | 42 | 3 | 2018-01-08 | 1 year ago | Go library for building command line applications. |
| [wlog](https://github.com/dixonwille/wlog) | 42 | 1 | 2016-04-13 | 6 months ago | Simple logging interface that supports cross-platform color and concurrency. |
| [tabular](https://github.com/InVisionApp/tabular) | 41 | 85 | 2018-04-23 | 2 years ago | Print ASCII tables from command line utilities without the need to pass large sets of data to the API. |
| [cmdr](https://github.com/hedzr/cmdr) | 37 | 5 | 2019-05-15 | 1 day ago | A POSIX/GNU style, getopt-like command-line UI Go library. |
| [strumt](https://github.com/antham/strumt) | 37 | 1 | 2017-06-19 | 1 month ago | Library to create prompt chain. |
| [clir](https://github.com/leaanthony/clir) | 35 | 2 | 2019-11-18 | 2 months ago | A Simple and Clear CLI library. Dependency free. |
| [flagvar](https://github.com/sgreben/flagvar) | 33 | 2 | 2018-05-18 | 2 weeks ago | A collection of flag argument types for Go's standard `flag` package. |
| [argv](https://github.com/cosiner/argv) | 27 | 1 | 2017-01-22 | 3 months ago | Go library to split command line string as arguments array using the bash syntax. |
| [go-getoptions](https://github.com/DavidGamba/go-getoptions) | 25 | 4 | 2015-12-18 | 4 days ago | Go option parser inspired on the flexibility of Perl’s GetOpt::Long. |
| [ctc](https://github.com/wzshiming/ctc) | 25 | 1 | 2018-04-27 | 1 week ago | The non-invasive cross-platform terminal color library does not need to modify the Print method. |
| [colourize](https://github.com/TreyBastian/colourize) | 20 | 3 | 2015-05-11 | 4 years ago | Go library for ANSI colour text in terminals. |
| [argv](https://github.com/zhuah/argv) | 19 | 1 | 2017-01-22 | 1 year ago | Go library to split command line string as arguments array using the bash syntax. |
| [go-commander](https://github.com/yitsushi/go-commander) | 18 | 1 | 2016-10-10 | 2 months ago | Go library to simplify CLI workflow. |
| [cmd](https://github.com/posener/cmd) | 16 | 2 | 2019-10-29 | 3 months ago | Extends the standard `flag` package to support sub commands and more in idomatic way. |
| [ts](https://github.com/liujianping/ts) | 12 | 0 | 2019-06-25 | 1 year ago | Timestamp convert & compare tool. |
| [sand](https://github.com/Zaba505/sand) | 10 | 1 | 2018-11-18 | 1 year ago | Simple API for creating interpreters and so much more. |
| [go-ataman](https://github.com/workanator/go-ataman) | 8 | 2 | 2017-05-17 | 2 years ago | Go library for rendering ANSI colored text templates in terminals. |

## Configuration
        
*Libraries for configuration parsing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [viper](https://github.com/spf13/viper) | 12942 | 225 | 2014-04-02 | 1 week ago | Go configuration with fangs. |
| [envconfig](https://github.com/kelseyhightower/envconfig) | 3124 | 42 | 2013-11-06 | 1 month ago | Go library for managing configuration data from environment variables. |
| [godotenv](https://github.com/joho/godotenv) | 3036 | 37 | 2013-07-30 | 3 weeks ago | Go port of Ruby's dotenv library (Loads environment variables from `.env`). |
| [ini](https://github.com/go-ini/ini) | 2127 | 76 | 2014-12-18 | 2 months ago | Go package to read and write INI files. |
| [env](https://github.com/caarlos0/env) | 1566 | 19 | 2015-07-28 | 3 days ago | Parse environment variables to Go structs (with defaults). |
| [konfig](https://github.com/lalamove/konfig) | 570 | 13 | 2019-01-18 | 1 week ago | Composable, observable and performant config handling for Go for the distributed processing era. |
| [confita](https://github.com/heetch/confita) | 316 | 25 | 2017-12-21 | 1 month ago | Load configuration in cascade from multiple backends into a struct. |
| [store](https://github.com/tucnak/store) | 247 | 5 | 2015-10-03 | 2 years ago | Lightweight configuration manager for Go. |
| [config](https://github.com/JeremyLoy/config) | 244 | 1 | 2019-04-02 | 4 months ago | Cloud native application configuration. Bind ENV to structs in only two lines. |
| [config](https://github.com/olebedev/config) | 230 | 8 | 2014-04-21 | 1 year ago | JSON or YAML configuration wrapper with environment variables and flags parsing. |
| [hjson-go](https://github.com/hjson/hjson-go) | 208 | 9 | 2016-08-05 | 2 months ago | Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments. |
| [config](https://github.com/joshbetz/config) | 201 | 2 | 2017-04-02 | 9 months ago | Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP. |
| [koanf](https://github.com/knadh/koanf) | 199 | 7 | 2019-06-18 | 2 days ago | Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line. |
| [envconfig](https://github.com/vrischmann/envconfig) | 173 | 4 | 2015-04-21 | 1 month ago | Read your configuration from environment variables. |
| [config](https://github.com/gookit/config) | 162 | 5 | 2018-07-07 | 1 week ago | application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge. |
| [gcfg](https://github.com/go-gcfg/gcfg) | 131 | 7 | 2015-08-17 | 2 weeks ago | read INI-style configuration files into Go structs; supports user-defined types and subsections. |
| [cleanenv](https://github.com/ilyakaznacheev/cleanenv) | 129 | 3 | 2019-07-12 | 1 month ago | Minimalistic configuration reader (from files, ENV, and wherever you want). |
| [goconfig](https://github.com/crgimenes/goconfig) | 127 | 10 | 2016-12-18 | 2 months ago | Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. |
| [fig](https://github.com/kkyr/fig) | 120 | 4 | 2020-01-16 | 2 months ago | Tiny library for reading configuration from a file and from environment variables (with validation & defaults). |
| [envh](https://github.com/antham/envh) | 95 | 3 | 2017-01-12 | 2 weeks ago | Helpers to manage environment variables. |
| [envcfg](https://github.com/tomazk/envcfg) | 92 | 1 | 2014-11-29 | 3 years ago | Un-marshaling environment variables to Go structs. |
| [config](https://github.com/golobby/config) | 69 | 2 | 2019-10-15 | 5 days ago | A lightweight yet powerful config package for Go projects. |
| [harvester](https://github.com/beatlabs/harvester) | 60 | 12 | 2019-04-09 | 1 hour ago | Harvester, a easy to use static and dynamic configuration package supportig seeding, env vars and Consul integration. |
| [gofigure](https://github.com/ian-kent/gofigure) | 58 | 5 | 2014-11-25 | 10 months ago | Go application configuration made easy. |
| [configure](https://github.com/paked/configure) | 55 | 3 | 2015-06-14 | 1 year ago | Provides configuration through multiple sources, including JSON, flags and environment variables. |
| [xdg](https://github.com/OpenPeeDeeP/xdg) | 53 | 3 | 2017-07-20 | 8 months ago | Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html). |
| [configuro](https://github.com/sherifabdlnaby/configuro) | 53 | 4 | 2020-04-09 | 1 day ago | opinionated configuration loading & validation framework from ENV and Files focused towards 12-Factor compliant applications. |
| [go-aws-ssm](https://github.com/PaddleHQ/go-aws-ssm) | 32 | 16 | 2019-01-24 | 1 week ago | Go package that fetches parameters from AWS System Manager - Parameter Store. |
| [ingo](https://github.com/schachmat/ingo) | 31 | 1 | 2016-02-07 | 3 years ago | Flags persisted in an ini-like config file. |
| [configuration](https://github.com/BoRuDar/configuration) | 29 | 3 | 2019-11-27 | 3 months ago | Library for initializing configuration structs from env variables, files, flags and 'default' tag. |
| [go-up](https://github.com/ufoscout/go-up) | 28 | 1 | 2018-02-18 | 6 months ago | A simple configuration library with recursive placeholders resolution and no magic. |
| [mini](https://github.com/sasbury/mini) | 24 | 1 | 2015-04-29 | 1 year ago | Golang package for parsing ini-style configuration files. |
| [conflate](https://github.com/the4thamigo-uk/conflate) | 16 | 0 | 2018-02-01 | 1 year ago | Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema. |
| [genv](https://github.com/sakirsensoy/genv) | 13 | 1 | 2019-07-15 | 1 year ago | Read environment variables easily with dotenv support. |
| [hocon](https://github.com/gurkankaymak/hocon) | 12 | 0 | 2020-03-01 | 3 months ago | Configuration library for working with the HOCON(a human-friendly JSON superset) format, supports features like environment variables, referencing other values, comments and multiple files. |
| [envconf](https://github.com/ian-kent/envconf) | 9 | 1 | 2014-10-26 | 5 years ago | Configuration from environment. |
| [go-ssm-config](https://github.com/ianlopshire/go-ssm-config) | 8 | 1 | 2019-12-02 | 1 week ago | Go utility for loading configuration parameters from AWS SSM (Parameter Store). |
| [sprbox](https://github.com/oblq/sprbox) | 4 | 2 | 2018-07-17 | 3 months ago | Build-environment aware toolbox factory and agnostic config parser (YAML, TOML, JSON and Environment vars). |
| [env](https://github.com/nasermirzaei89/env) | 2 | 0 | 2019-07-24 | 2 months ago | Simple useful package for read environment variables. |
| [go-ini](https://github.com/subpop/go-ini) | 2 | 1 | 2019-09-11 | 2 months ago | A Go package that marshals and unmarshals INI-files. |
| [typenv](https://github.com/diegomarangoni/typenv) | 2 | 1 | 2020-06-30 | 6 days ago | Minimalistic, zero dependency, typed environment variables library. |
| [swap](https://github.com/oblq/swap) | 1 | 1 | 2020-04-12 | 2 months ago | Instantiate/configure structs recursively, based on build environment. (YAML, TOML, JSON and env). |

## Continuous Integration
        
*Tools for help with continuous integration.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [drone](https://github.com/drone/drone) | 21456 | 579 | 2014-02-07 | 3 days ago | Drone is a Continuous Integration platform built on Docker, written in Go. |
| [cds](https://github.com/ovh/cds) | 3001 | 74 | 2016-10-11 | 1 hour ago | Enterprise-Grade CI/CD and DevOps Automation Open Source Platform. |
| [goveralls](https://github.com/mattn/goveralls) | 651 | 13 | 2013-04-17 | 2 months ago | Go integration for Coveralls.io continuous code coverage tracking system. |
| [overalls](https://github.com/go-playground/overalls) | 104 | 3 | 2015-07-30 | 7 months ago | Multi-Package go project coverprofile for tools like goveralls. |
| [duci](https://github.com/duck8823/duci) | 57 | 3 | 2018-04-01 | 6 days ago | A simple ci server no needs domain specific languages. |
| [gomason](https://github.com/nikogura/gomason) | 49 | 1 | 2017-11-18 | 1 week ago | Test, Build, Sign, and Publish your go binaries from a clean workspace. |
| [roveralls](https://github.com/lawrencewoodman/roveralls) | 13 | 1 | 2016-06-26 | 2 years ago | Recursive coverage testing tool. |

## CSS Preprocessors
        
*Libraries for preprocessing CSS files.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gcss](https://github.com/yosssi/gcss) | 434 | 16 | 2014-09-04 | 5 years ago | Pure Go CSS Preprocessor. |
| [go-libsass](https://github.com/wellington/go-libsass) | 159 | 6 | 2015-04-19 | 1 year ago | Go wrapper to the 100% Sass compatible libsass project. |

## Data Structures
        
*Generic datastructures and algorithms in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gods](https://github.com/emirpasic/gods) | 8619 | 350 | 2015-03-04 | 2 weeks ago | Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc. |
| [go-datastructures](https://github.com/Workiva/go-datastructures) | 5665 | 322 | 2014-10-29 | 4 months ago | Collection of useful, performant, and thread-safe data structures. |
| [golang-set](https://github.com/deckarep/golang-set) | 1572 | 41 | 2013-07-03 | 7 months ago | Thread-Safe and Non-Thread-Safe high-performance sets for Go. |
| [gota](https://github.com/go-gota/gota) | 1304 | 69 | 2016-02-06 | 1 week ago | Implementation of dataframes, series, and data wrangling methods for Go. |
| [BoomFilters](https://github.com/tylertreat/BoomFilters) | 1276 | 40 | 2015-02-06 | 2 months ago | Probabilistic data structures for processing continuous, unbounded streams. |
| [roaring](https://github.com/RoaringBitmap/roaring) | 894 | 36 | 2014-07-10 | 3 days ago | Go package implementing compressed bitsets. |
| [bloom](https://github.com/willf/bloom) | 854 | 34 | 2011-05-21 | 9 months ago | Go package implementing Bloom filters. |
| [hyperloglog](https://github.com/axiomhq/hyperloglog) | 706 | 17 | 2017-06-18 | 8 months ago | HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction. |
| [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) | 650 | 19 | 2015-06-28 | 2 months ago | Cuckoo filter: a good alternative to a counting bloom filter implemented in Go. |
| [bitset](https://github.com/willf/bitset) | 582 | 31 | 2011-05-11 | 1 hour ago | Go package implementing bitsets. |
| [trie](https://github.com/derekparker/trie) | 488 | 19 | 2014-03-06 | 4 months ago | Trie implementation in Go. |
| [gocache](https://github.com/eko/gocache) | 476 | 13 | 2019-10-05 | 2 weeks ago | A complete Go cache library with mutiple stores (memory, memcache, redis, ...), chainable, loadable, metrics cache and more. |
| [algorithms](https://github.com/shady831213/algorithms) | 408 | 15 | 2018-01-31 | 1 year ago | Algorithms and data structures.CLRS study. |
| [go-geoindex](https://github.com/hailocab/go-geoindex) | 322 | 69 | 2015-01-22 | 2 years ago | In-memory geo index. |
| [mafsa](https://github.com/smartystreets-archives/mafsa) | 282 | 16 | 2014-11-07 | 1 year ago | MA-FSA implementation with Minimal Perfect Hashing. |
| [goskiplist](https://github.com/ryszard/goskiplist) | 216 | 15 | 2012-05-09 | 9 months ago | Skip list implementation in Go. |
| [merkletree](https://github.com/cbergoon/merkletree) | 206 | 4 | 2017-04-12 | 8 months ago | Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures. |
| [hilbert](https://github.com/google/hilbert) | 201 | 22 | 2015-08-06 | 1 year ago | Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves. |
| [deque](https://github.com/gammazero/deque) | 161 | 6 | 2018-04-24 | 6 days ago | Fast ring-buffer deque (double-ended queue). |
| [ttlcache](https://github.com/ReneKroon/ttlcache) | 153 | 10 | 2014-12-13 | 8 hours ago | In-memory LRU string-interface{} map with expiration for golang. |
| [binpacker](https://github.com/zhuangsirui/binpacker) | 143 | 11 | 2016-02-02 | 2 years ago | Binary packer and unpacker helps user build custom binary stream. |
| [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) | 142 | 7 | 2016-04-01 | 11 months ago | Go implementation of Adaptive Radix Tree. |
| [bloom](https://github.com/zentures/bloom) | 132 | 7 | 2013-09-03 | 2 years ago | Bloom filters implemented in Go. |
| [skiplist](https://github.com/MauriceGit/skiplist) | 126 | 5 | 2018-06-23 | 8 months ago | Very fast Go Skiplist implementation. |
| [gostl](https://github.com/liyue201/gostl) | 120 | 6 | 2019-10-12 | 5 days ago | Data structure and algorithm library for go, designed to provide functions similar to C++ STL. |
| [iter](https://github.com/disksing/iter) | 119 | 5 | 2019-10-20 | 7 months ago | Go implementation of C++ STL iterators and algorithms. |
| [ring](https://github.com/TheTannerRyan/ring) | 110 | 1 | 2019-01-27 | 9 months ago | Go implementation of a high performance, thread safe bloom filter. |
| [go-rquad](https://github.com/arl/go-rquad) | 105 | 3 | 2016-09-12 | 3 months ago | Region quadtrees with efficient point location and neighbour finding. |
| [encoding](https://github.com/zentures/encoding) | 102 | 6 | 2013-09-20 | 2 years ago | Integer Compression Libraries for Go. |
| [levenshtein](https://github.com/agnivade/levenshtein) | 98 | 3 | 2014-07-30 | 4 weeks ago | Implementation to calculate levenshtein distance in Go. |
| [bit](https://github.com/yourbasic/bit) | 90 | 6 | 2017-05-03 | 2 years ago | Golang set data structure with bonus bit-twiddling functions. |
| [conjungo](https://github.com/InVisionApp/conjungo) | 90 | 98 | 2016-12-29 | 1 year ago | A small, powerful and flexible merge library. |
| [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) | 86 | 1 | 2019-01-10 | 2 months ago | Concurrent FIFO queue. |
| [skiplist](https://github.com/gansidui/skiplist) | 71 | 7 | 2014-11-18 | 5 years ago | Skiplist implementation in Go. |
| [remember-go](https://github.com/rocketlaunchr/remember-go) | 56 | 4 | 2019-04-04 | 9 months ago | A universal interface for caching slow database queries (backed by redis, memcached, ristretto, or in-memory). |
| [bloom](https://github.com/yourbasic/bloom) | 54 | 2 | 2017-05-06 | 3 years ago | Golang Bloom filter implementation. |
| [go-mcache](https://github.com/OrlovEvgeny/go-mcache) | 54 | 3 | 2018-04-14 | 6 months ago | Fast in-memory key:value store/cache library. Pointer caches. |
| [count-min-log](https://github.com/seiflotfy/count-min-log) | 50 | 2 | 2015-08-16 | 3 years ago | Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory). |
| [levenshtein](https://github.com/agext/levenshtein) | 44 | 1 | 2016-04-08 | 4 months ago | Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. |
| [crunch](https://github.com/superwhiskers/crunch) | 40 | 6 | 2019-02-27 | 2 months ago | Go package implementing buffers for handling various datatypes easily. |
| [concurrent-writer](https://github.com/free/concurrent-writer) | 27 | 4 | 2017-09-18 | 2 years ago | Highly concurrent drop-in replacement for `bufio.Writer`. |
| [pipeline](https://github.com/hyfather/pipeline) | 24 | 1 | 2018-04-25 | 1 year ago | An implementation of pipelines with fan-in and fan-out. |
| [goset](https://github.com/zoumo/goset) | 23 | 1 | 2017-08-25 | 9 months ago | A useful Set collection implementation for Go. |
| [hide](https://github.com/emvi/hide) | 22 | 3 | 2019-01-16 | 1 month ago | ID type with marshalling to/from hash to prevent sending IDs to clients. |
| [typ](https://github.com/gurukami/typ) | 21 | 1 | 2019-03-03 | 11 months ago | Null Types, Safe primitive type conversion and fetching value from complex structures. |
| [deque](https://github.com/edwingeng/deque) | 20 | 2 | 2019-02-01 | 4 months ago | A highly optimized double-ended queue. |
| [dict](https://github.com/srfrog/dict) | 15 | 1 | 2019-04-23 | 1 year ago | Python-like dictionaries (dict) for Go. |
| [go-ef](https://github.com/amallia/go-ef) | 15 | 2 | 2017-09-22 | 2 years ago | A Go implementation of the Elias-Fano encoding. |
| [null](https://github.com/emvi/null) | 12 | 2 | 2018-07-04 | 1 month ago | Nullable Go types that can be marshalled/unmarshalled to/from JSON. |
| [timedmap](https://github.com/zekroTJA/timedmap) | 12 | 2 | 2019-01-30 | 2 months ago | Map with expiring key-value pairs. |
| [ptrie](https://github.com/viant/ptrie) | 10 | 7 | 2019-05-20 | 8 months ago | An implementation of prefix tree. |
| [mspm](https://github.com/BlackRabbitt/mspm) | 9 | 3 | 2018-05-17 | 2 years ago | Multi-String Pattern Matching Algorithm for information retrieval. |
| [treap](https://github.com/perdata/treap) | 9 | 2 | 2018-09-16 | 7 months ago | Persistent, fast ordered map using tree heaps. |
| [gofal](https://github.com/xxjwxc/gofal) | 7 | 1 | 2019-08-05 | 9 months ago | fractional api for Go. |
| [set](https://github.com/StudioSol/set) | 7 | 17 | 2018-07-20 | 1 year ago | Simple set data structure implementation in Go using LinkedHashMap. |
| [parsefields](https://github.com/MonaxGT/parsefields) | 4 | 1 | 2019-04-12 | 1 year ago | Tools for parse JSON-like logs for collecting unique fields and events. |
| [cmap](https://github.com/lrita/cmap) | 4 | 1 | 2019-11-26 | 2 months ago | a thread-safe concurrent map for go, support using `interface{}` as key and auto scale up shards. |

## Database
        
*SQL query builder, libraries for building and using SQL.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [prometheus](https://github.com/prometheus/prometheus) | 32112 | 1139 | 2012-11-24 | 4 hours ago | Monitoring system and time series database. |
| [tidb](https://github.com/pingcap/tidb) | 24428 | 1327 | 2015-09-06 | 57 minutes ago | TiDB is a distributed SQL database. Inspired by the design of Google F1. |
| [influxdb](https://github.com/influxdata/influxdb) | 19283 | 755 | 2013-09-26 | 1 hour ago | Scalable datastore for metrics, events, and real-time analytics. |
| [cockroach](https://github.com/cockroachdb/cockroach) | 18553 | 731 | 2014-02-06 | 52 minutes ago | Scalable, Geo-Replicated, Transactional Datastore. |
| [dgraph](https://github.com/dgraph-io/dgraph) | 13691 | 379 | 2015-08-25 | 58 minutes ago | Scalable, Distributed, Low Latency, High Throughput Graph Database. |
| [bolt](https://github.com/boltdb/bolt) | 10426 | 343 | 2013-12-20 | 2 years ago | Low-level key/value database for Go. |
| [vitess](https://github.com/vitessio/vitess) | 10399 | 510 | 2013-06-27 | 2 hours ago | vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. |
| [groupcache](https://github.com/golang/groupcache) | 8952 | 473 | 2013-07-22 | 2 months ago | Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. |
| [badger](https://github.com/dgraph-io/badger) | 7936 | 249 | 2017-01-26 | 2 hours ago | Fast key-value store in Go. |
| [pgweb](https://github.com/sosedoff/pgweb) | 6475 | 156 | 2014-10-09 | 1 week ago | Web-based PostgreSQL database browser. |
| [rqlite](https://github.com/rqlite/rqlite) | 6028 | 204 | 2014-08-23 | 5 days ago | The lightweight, distributed, relational database built on SQLite. |
| [kingshard](https://github.com/flike/kingshard) | 5292 | 407 | 2015-07-04 | 1 week ago | kingshard is a high performance proxy for MySQL powered by Golang. |
| [migrate](https://github.com/golang-migrate/migrate) | 4757 | 55 | 2018-01-19 | 1 day ago | Database migrations. CLI and Golang library. |
| [bigcache](https://github.com/allegro/bigcache) | 4138 | 113 | 2016-03-23 | 3 weeks ago | Efficient key/value cache for gigabytes of data. |
| [go-cache](https://github.com/patrickmn/go-cache) | 4076 | 105 | 2012-01-02 | 3 weeks ago | In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. |
| [goleveldb](https://github.com/syndtr/goleveldb) | 3817 | 173 | 2013-01-23 | 2 weeks ago | Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go. |
| [orchestrator](https://github.com/openark/orchestrator) | 3691 | 279 | 2016-11-30 | 9 hours ago | MySQL replication topology manager & visualizer. |
| [bbolt](https://github.com/etcd-io/bbolt) | 3479 | 108 | 2017-06-17 | 1 day ago | An embedded key/value database for Go. |
| [orchestrator](https://github.com/github/orchestrator) | 3362 | 299 | 2016-11-30 | 5 months ago | MySQL replication topology manager & visualizer. |
| [ledisdb](https://github.com/ledisdb/ledisdb) | 3361 | 185 | 2014-04-30 | 2 months ago | Ledisdb is a high performance NoSQL like Redis based on LevelDB. |
| [ledisdb](https://github.com/siddontang/ledisdb) | 3265 | 188 | 2014-04-30 | 3 months ago | Ledisdb is a high performance NoSQL like Redis based on LevelDB. |
| [squirrel](https://github.com/Masterminds/squirrel) | 3164 | 46 | 2014-01-18 | 2 months ago | Go library that helps you build SQL queries. |
| [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) | 3035 | 169 | 2015-01-15 | 7 months ago | Sync your MySQL data into Elasticsearch automatically. |
| [buntdb](https://github.com/tidwall/buntdb) | 2862 | 99 | 2016-07-19 | 1 month ago | Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support. |
| [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) | 2728 | 81 | 2018-09-30 | 4 hours ago | fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL. |
| [go-mysql](https://github.com/siddontang/go-mysql) | 2564 | 157 | 2014-02-21 | 1 week ago | Go toolset to handle MySQL protocol and replication. |
| [xo](https://github.com/xo/xo) | 2521 | 75 | 2016-02-05 | 1 month ago | Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. |
| [tiedot](https://github.com/HouzuoGuo/tiedot) | 2513 | 159 | 2013-05-26 | 3 months ago | Your NoSQL database powered by Golang. |
| [prest](https://github.com/prest/prest) | 2323 | 88 | 2016-11-22 | 4 days ago | Serve a RESTful API from any PostgreSQL database. |
| [sql-migrate](https://github.com/rubenv/sql-migrate) | 1820 | 33 | 2014-09-09 | 3 weeks ago | Database migration tool. Allows embedding migrations into the application using go-bindata. |
| [cache2go](https://github.com/muesli/cache2go) | 1330 | 65 | 2013-11-11 | 2 weeks ago | In-memory key:value cache which supports automatic invalidation based on timeouts. |
| [nutsdb](https://github.com/xujiajun/nutsdb) | 1297 | 40 | 2018-12-07 | 1 month ago | Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set. |
| [gcache](https://github.com/bluele/gcache) | 1191 | 39 | 2015-01-24 | 1 week ago | Cache library with support for expirable Cache, LFU, LRU and ARC. |
| [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) | 1062 | 68 | 2018-04-11 | 3 months ago | CovenantSQL is a SQL database on blockchain. |
| [gendry](https://github.com/didi/gendry) | 1032 | 60 | 2017-12-01 | 1 week ago | Non-invasive SQL builder and powerful data binder. |
| [diskv](https://github.com/peterbourgon/diskv) | 935 | 37 | 2012-03-21 | 2 months ago | Home-grown disk-backed key-value store. |
| [fastcache](https://github.com/VictoriaMetrics/fastcache) | 844 | 25 | 2018-11-22 | 4 months ago | fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead. |
| [goqu](https://github.com/doug-martin/goqu) | 812 | 29 | 2015-02-21 | 2 months ago | Idiomatic SQL builder and query library. |
| [moss](https://github.com/couchbase/moss) | 764 | 77 | 2016-02-06 | 3 months ago | Moss is a simple LSM key-value storage engine written in 100% Go. |
| [skeema](https://github.com/skeema/skeema) | 708 | 29 | 2016-10-31 | 16 hours ago | Pure-SQL schema management system for MySQL, with support for sharding and external online schema change tools. |
| [eliasdb](https://github.com/krotik/eliasdb) | 579 | 24 | 2016-08-13 | 1 week ago | Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language. |
| [pogreb](https://github.com/akrylysov/pogreb) | 573 | 21 | 2018-01-06 | 2 months ago | Embedded key-value store for read-heavy workloads. |
| [dotsql](https://github.com/gchaincl/dotsql) | 542 | 22 | 2014-11-20 | 3 months ago | Go library that helps you keep sql files in one place and use them with ease. |
| [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) | 483 | 25 | 2015-12-10 | 5 months ago | Powerful data retrieval methods as well as DB-agnostic query building capabilities. |
| [gormigrate](https://github.com/go-gormigrate/gormigrate) | 473 | 6 | 2016-08-31 | 5 days ago | Database schema migration helper for Gorm ORM. |
| [chproxy](https://github.com/Vertamedia/chproxy) | 441 | 25 | 2017-09-18 | 3 months ago | HTTP proxy for ClickHouse database. |
| [bitcask](https://github.com/prologic/bitcask) | 405 | 13 | 2019-03-12 | 1 day ago | Bitcask is an embeddable, persistent and fast key-value (KV) database written in pure Go with predictable read/write performance, low latency and high throughput thanks to the bitcask on-disk layout (LSM+WAL). |
| [levigo](https://github.com/jmhodges/levigo) | 389 | 24 | 2012-01-17 | 1 month ago | Levigo is a Go wrapper for LevelDB. |
| [jet](https://github.com/go-jet/jet) | 293 | 10 | 2019-03-02 | 1 month ago | Framework for writing type-safe SQL queries in Go, with ability to easily convert database query result into desired arbitrary object structure. |
| [pudge](https://github.com/recoilme/pudge) | 268 | 10 | 2018-11-20 | 3 months ago | Fast and simple  key/value store written using Go's standard library. |
| [sqrl](https://github.com/elgris/sqrl) | 212 | 8 | 2014-06-25 | 7 months ago | SQL query builder, fork of Squirrel with improved performance. |
| [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) | 211 | 7 | 2017-04-29 | 2 months ago | Collects small insterts and sends big requests to ClickHouse servers. |
| [dbq](https://github.com/rocketlaunchr/dbq) | 203 | 8 | 2019-07-11 | 5 days ago | Zero boilerplate database operations for Go. |
| [vasto](https://github.com/chrislusf/vasto) | 187 | 20 | 2018-01-16 | 1 year ago | A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption. |
| [kivik](https://github.com/go-kivik/kivik) | 182 | 9 | 2017-02-09 | 2 weeks ago | Kivik provides a common Go and GopherJS client library for CouchDB, PouchDB, and similar databases. |
| [piladb](https://github.com/fern4lvarez/piladb) | 177 | 11 | 2015-09-08 | 2 years ago | Lightweight RESTful database engine based on stack data structures. |
| [myreplication](https://github.com/2tvenom/myreplication) | 162 | 19 | 2015-02-04 | 1 year ago | MySql binary log replication listener. Supports statement and row based replication. |
| [goose](https://github.com/steinbacher/goose) | 141 | 4 | 2016-03-04 | 3 years ago | Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts. |
| [golang-scribble](https://github.com/nanobox-io/golang-scribble) | 104 | 3 | 2018-06-21 | 1 year ago | Tiny flat file JSON store. |
| [darwin](https://github.com/GuiaBolso/darwin) | 103 | 3 | 2016-04-05 | 7 months ago | Database schema evolution library for Go. |
| [slowpoke](https://github.com/recoilme/slowpoke) | 98 | 8 | 2018-02-19 | 10 months ago | Key-value store with persistence. |
| [migrator](https://github.com/lopezator/migrator) | 98 | 4 | 2019-02-04 | 3 months ago | Dead simple Go database migration library. |
| [octillery](https://github.com/knocknote/octillery) | 86 | 15 | 2018-11-26 | 1 week ago | Go package for sharding databases ( Supports every ORM or raw SQL ). |
| [igor](https://github.com/galeone/igor) | 82 | 7 | 2016-03-10 | 4 weeks ago | Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax. |
| [sqlingo](https://github.com/lqs/sqlingo) | 70 | 9 | 2018-11-18 | 6 hours ago | A lightweight DSL to build SQL in Go. |
| [databunker](https://github.com/paranoidguy/databunker) | 65 | 6 | 2019-12-08 | 2 days ago | Personally identifiable information (PII) storage service built to comply with GDPR and CCPA. |
| [cache](https://github.com/akyoto/cache) | 61 | 2 | 2019-05-11 | 5 months ago | In-memory key:value store with expiration time, 0 dependencies, <100 LoC, 100% coverage. |
| [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) | 56 | 1 | 2018-08-11 | 1 month ago | A Go package to help write migrations with go-pg/pg. |
| [bcache](https://github.com/iwanbk/bcache) | 54 | 2 | 2018-12-26 | 1 year ago | Eventually consistent distributed in-memory  cache Go library. |
| [godbal](https://github.com/xujiajun/godbal) | 51 | 4 | 2018-02-28 | 1 year ago | Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily. |
| [couchcache](https://github.com/codingsince1985/couchcache) | 47 | 3 | 2015-04-05 | 1 year ago | RESTful caching micro-service backed by Couchbase server. |
| [dbbench](https://github.com/sj14/dbbench) | 45 | 3 | 2018-11-24 | 1 month ago | Database benchmarking tool with support for several databases and scripts. |
| [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) | 35 | 6 | 2017-12-18 | 2 years ago | BigCache with clustering support and individual item expiration. |
| [unitdb](https://github.com/unit-io/unitdb) | 34 | 1 | 2019-08-29 | 18 hours ago | Fast timeseries database for IoT, realtime messaging  applications. Access unitdb with pubsub over tcp or websocket using github.com/unit-io/unitd application. |
| [gondolier](https://github.com/emvi/gondolier) | 33 | 4 | 2017-10-18 | 1 year ago | Database migration library using struct decorators. |
| [prep](https://github.com/hexdigest/prep) | 27 | 2 | 2017-12-11 | 2 years ago | Use prepared SQL statements without changing your code. |
| [go-fixtures](https://github.com/RichardKnop/go-fixtures) | 24 | 1 | 2015-12-24 | 7 months ago | Django style fixtures for Golang's excellent built-in database/sql library. |
| [pravasan](https://github.com/pravasan/pravasan) | 24 | 6 | 2015-01-03 | 1 year ago | Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc. |
| [datagen](https://github.com/codingconcepts/datagen) | 23 | 1 | 2019-04-18 | 1 month ago | A fast data generator that's multi-table aware and supports multi-row DML. |
| [coffer](https://github.com/claygod/coffer) | 20 | 3 | 2019-05-13 | 5 months ago | Simple ACID key-value database that supports transactions. |
| [buildsqlx](https://github.com/arthurkushman/buildsqlx) | 16 | 1 | 2019-08-18 | 2 months ago | Go database query builder library for PostgreSQL. |
| [tempdb](https://github.com/rafaeljesus/tempdb) | 14 | 3 | 2017-03-17 | 2 years ago | Key-value store for temporary items. |
| [avro](https://github.com/khezen/avro) | 14 | 1 | 2019-04-07 | 1 week ago | Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes. |
| [gorocksdb](https://github.com/kapitan-k/gorocksdb) | 12 | 1 | 2017-12-28 | 2 years ago | Gorocksdb is a wrapper for [RocksDB](https://rocksdb.org) written in Go. |
| [rwdb](https://github.com/andizzle/rwdb) | 11 | 2 | 2017-10-04 | 2 years ago | rwdb provides read replica capability for multiple database servers setup. |
| [sqlf](https://github.com/leporo/sqlf) | 10 | 2 | 2019-07-20 | 5 months ago | Fast SQL query builder. |
| [qry](https://github.com/HnH/qry) | 9 | 1 | 2019-08-20 | 6 months ago | Tool that generates constants from files with raw SQL queries. |
| [bucket](https://github.com/PumpkinSeed/bucket) | 6 | 3 | 2019-09-22 | 2 months ago | Optimized data structure framework for Couchbase specialized on one bucket usage. |
| [schema](https://github.com/adlio/schema) | 5 | 2 | 2019-09-24 | 2 months ago | Library to embed schema migrations for database/sql-compatible databases inside your Go binaries. |
| [gostore](https://github.com/twharmon/gostore) | 4 | 1 | 2020-01-08 | 5 months ago | Gostore is a simple, durable, embedded key-value storage engine written in Go. |
| [mpath-go](https://github.com/spacetab-io/mpath-go) | 4 | 3 | 2020-01-09 | 6 months ago | MPTT (Modified Preorder Tree Traversal) package for SQL records - materialized path realisation. |
| [gosql](https://github.com/twharmon/gosql) | 4 | 1 | 2020-01-08 | 2 weeks ago | SQL Query builder with better null values support. |
| [tracedb](https://github.com/unit-io/tracedb) | 3 | 1 | 2019-08-29 | 2 months ago | Fast timeseries database for IoT, realtime messaging  applications. Access tracedb with pubsub over tcp or websocket using github.com/unit-io/trace application. |
| [ormlite](https://github.com/pupizoid/ormlite) | 0 | 2 | 2018-06-28 | 2 months ago | Lightweight package containing some ORM-like features and helpers for sqlite databases. |

## Database Drivers
        
*Libraries for connecting and operating databases.*

### Relational Databases
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [mysql](https://github.com/go-sql-driver/mysql) | 9747 | 413 | 2012-12-09 | 8 hours ago | MySQL driver for Go. |
| [pq](https://github.com/lib/pq) | 6025 | 157 | 2012-03-12 | 18 hours ago | Pure Go Postgres driver for database/sql. |
| [go-sqlite3](https://github.com/mattn/go-sqlite3) | 4199 | 141 | 2011-11-11 | 5 days ago | SQLite3 driver for go that uses database/sql. |
| [pgx](https://github.com/jackc/pgx) | 2961 | 74 | 2013-03-30 | 6 days ago | PostgreSQL driver supporting features beyond those exposed by database/sql. |
| [go-mssqldb](https://github.com/denisenkom/go-mssqldb) | 1235 | 67 | 2013-12-16 | 1 month ago | Microsoft MSSQL driver for Go. |
| [go-oci8](https://github.com/mattn/go-oci8) | 480 | 40 | 2012-02-29 | 3 weeks ago | Oracle driver for go that uses database/sql. |
| [goracle](https://github.com/go-goracle/goracle) | 281 | 28 | 2015-03-25 | 7 months ago | Oracle driver for Go, using the ODPI-C driver. |
| [firebirdsql](https://github.com/nakagami/firebirdsql) | 127 | 18 | 2013-08-27 | 4 days ago | Firebird RDBMS SQL driver for Go. |
| [go-adodb](https://github.com/mattn/go-adodb) | 110 | 13 | 2011-11-14 | 5 months ago | Microsoft ActiveX Object DataBase driver for go that uses database/sql. |
| [gofreetds](https://github.com/minus5/gofreetds) | 98 | 22 | 2012-12-06 | 3 weeks ago | Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org). |
| [calcite-avatica-go](https://github.com/apache/calcite-avatica-go) | 54 | 25 | 2017-08-08 | 1 week ago | Apache Avatica/Phoenix SQL driver for database/sql. |
| [bgc](https://github.com/viant/bgc) | 14 | 10 | 2016-06-13 | 5 months ago | Datastore Connectivity for BigQuery for go. |
| [sqinn-go](https://github.com/cvilsmeier/sqinn-go) | 9 | 1 | 2020-06-06 | 1 month ago | SQLite with pure Go. |

### NoSQL Databases
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [cayley](https://github.com/cayleygraph/cayley) | 13490 | 612 | 2014-06-05 | 1 week ago | Graph database with support for multiple backends. |
| [redis](https://github.com/go-redis/redis) | 9315 | 241 | 2012-07-25 | 3 hours ago | Redis client for Golang. |
| [redigo](https://github.com/gomodule/redigo) | 7651 | 299 | 2012-04-14 | 1 month ago | Redigo is a Go client for the Redis database. |
| [bleve](https://github.com/blevesearch/bleve) | 6767 | 257 | 2014-04-17 | 6 days ago | Modern text indexing library for go. |
| [riot](https://github.com/go-ego/riot) | 5346 | 191 | 2017-06-21 | 1 month ago | Go Open Source, Distributed, Simple and efficient Search Engine. |
| [elastic](https://github.com/olivere/elastic) | 5185 | 168 | 2012-12-06 | 3 days ago | Elasticsearch client for Go. |
| [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) | 4651 | 141 | 2017-02-08 | 22 hours ago | Official MongoDB driver for the Go language. |
| [go-elasticsearch](https://github.com/elastic/go-elasticsearch) | 2652 | 281 | 2017-03-27 | 2 hours ago | Official Elasticsearch client for Go. |
| [mgo](https://github.com/globalsign/mgo) | 1799 | 65 | 2017-04-13 | 2 months ago | (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms. |
| [rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1533 | 48 | 2013-09-12 | 2 months ago | Go language driver for RethinkDB. |
| [elastigo](https://github.com/mattbaird/elastigo) | 950 | 48 | 2012-10-12 | 1 year ago | Elasticsearch client library. |
| [elasticsql](https://github.com/cch123/elasticsql) | 565 | 32 | 2016-08-24 | 3 weeks ago | Convert sql to elasticsearch dsl in Go. |
| [redeo](https://github.com/bsm/redeo) | 374 | 24 | 2014-03-06 | 11 months ago | Redis-protocol compatible TCP servers/services. |
| [neoism](https://github.com/jmcvetta/neoism) | 368 | 24 | 2012-07-12 | 5 months ago | Neo4j client for Golang. |
| [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) | 330 | 37 | 2014-07-26 | 3 weeks ago | Aerospike client in Go language. |
| [gocb](https://github.com/couchbase/gocb) | 318 | 65 | 2015-01-15 | 6 days ago | Official Couchbase Go SDK. |
| [go-couchbase](https://github.com/couchbase/go-couchbase) | 303 | 25 | 2012-01-19 | 1 day ago | Couchbase client in Go. |
| [gokv](https://github.com/philippgille/gokv) | 231 | 5 | 2018-10-08 | 4 months ago | Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more). |
| [mgm](https://github.com/Kamva/mgm) | 158 | 10 | 2019-12-27 | 4 weeks ago | MongoDB model-based ODM for Go (based on official MongoDB driver). |
| [go-rejson](https://github.com/nitishm/go-rejson) | 136 | 6 | 2018-04-23 | 1 week ago | Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease. |
| [cachego](https://github.com/faabiosr/cachego) | 132 | 6 | 2016-10-05 | 1 month ago | Golang Cache component for multiple drivers. |
| [godis](https://github.com/piaohao/godis) | 82 | 8 | 2019-06-14 | 2 months ago | redis client implement by golang, inspired by jedis. |
| [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) | 75 | 6 | 2011-06-04 | 2 years ago | Neo4j REST Client in golang. |
| [skizze](https://github.com/seiflotfy/skizze) | 74 | 7 | 2016-01-17 | 4 years ago | probabilistic data-structures service and storage. |
| [dynago](https://github.com/underarmour/dynago) | 69 | 123 | 2015-05-18 | 3 years ago | Dynago is a principle of least surprise client for DynamoDB. |
| [arangolite](https://github.com/solher/arangolite) | 66 | 5 | 2015-10-04 | 1 year ago | Lightweight golang driver for ArangoDB. |
| [go-pilosa](https://github.com/pilosa/go-pilosa) | 39 | 17 | 2016-09-30 | 4 months ago | Go client library for Pilosa. |
| [goforestdb](https://github.com/couchbase/goforestdb) | 28 | 41 | 2014-05-14 | 3 years ago | Go bindings for ForestDB. |
| [goriak](https://github.com/zegl/goriak) | 25 | 2 | 2016-10-05 | 1 year ago | Go language driver for Riak KV. |
| [neo4j](https://github.com/cihangir/neo4j) | 25 | 3 | 2013-05-18 | 5 years ago | Neo4j Rest API Bindings for Golang. |
| [goes](https://github.com/OwnLocal/goes) | 24 | 33 | 2015-12-28 | 3 years ago | Library to interact with Elasticsearch. |
| [dsc](https://github.com/viant/dsc) | 20 | 15 | 2016-06-13 | 4 months ago | Datastore connectivity for SQL, NoSQL, structured files. |
| [xredis](https://github.com/shomali11/xredis) | 11 | 1 | 2017-06-14 | 1 year ago | Typesafe, customizable, clean & easy to use Redis client. |
| [godscache](https://github.com/defcronyke/godscache) | 7 | 2 | 2018-05-08 | 1 year ago | A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached. |
| [asc](https://github.com/viant/asc) | 4 | 10 | 2016-06-13 | 1 year ago | Datastore Connectivity for Aerospike for go. |

## Date and Time
        
*Libraries for working with dates and times.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [now](https://github.com/jinzhu/now) | 2614 | 52 | 2013-11-18 | 9 months ago | Now is a time toolkit for golang. |
| [dateparse](https://github.com/araddon/dateparse) | 1076 | 17 | 2014-04-21 | 3 weeks ago | Parse date's without knowing format in advance. |
| [carbon](https://github.com/uniplaces/carbon) | 475 | 41 | 2016-08-03 | 1 year ago | Simple Time extension with a lot of util methods, ported from PHP Carbon library. |
| [durafmt](https://github.com/hako/durafmt) | 306 | 4 | 2016-05-20 | 2 weeks ago | Time duration formatting library for Go. |
| [timeutil](https://github.com/leekchan/timeutil) | 178 | 7 | 2015-08-02 | 1 year ago | Useful extensions (Timedelta, Strftime, ...) to the golang's time package. |
| [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) | 78 | 3 | 2016-01-31 | 8 months ago | The implementation of the Persian (Solar Hijri) Calendar in Go (golang). |
| [iso8601](https://github.com/relvacode/iso8601) | 78 | 3 | 2017-04-25 | 1 month ago | Efficiently parse ISO8601 date-times without regex. |
| [timespan](https://github.com/SaidinWoT/timespan) | 72 | 5 | 2014-10-07 | 1 year ago | For interacting with intervals of time, defined as a start time and a duration. |
| [date](https://github.com/rickb777/date) | 41 | 2 | 2015-11-23 | 2 weeks ago | Augments Time for working with dates, date ranges, time spans, periods, and time-of-day. |
| [feiertage](https://github.com/wlbr/feiertage) | 32 | 1 | 2015-11-04 | 2 months ago | Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving... |
| [go-sunrise](https://github.com/nathan-osman/go-sunrise) | 25 | 3 | 2017-06-15 | 2 years ago | Calculate the sunrise and sunset times for a given location. |
| [kair](https://github.com/GuilhermeCaruso/kair) | 15 | 0 | 2018-10-03 | 1 month ago | Date and Time - Golang Formatting Library. |
| [nulltime](https://github.com/kirillDanshin/nulltime) | 10 | 1 | 2016-03-06 | 3 years ago | Nullable `time.Time`. |
| [tuesday](https://github.com/osteele/tuesday) | 8 | 2 | 2017-08-10 | 2 years ago | Ruby-compatible Strftime function. |
| [strftime](https://github.com/awoodbeck/strftime) | 5 | 1 | 2018-02-10 | 2 years ago | C99-compatible strftime formatter. |
| [go-str2duration](https://github.com/xhit/go-str2duration) | 4 | 1 | 2020-02-02 | 2 hours ago | Convert string to duration. Support time.Duration returned string and more. |
| [cronrange](https://github.com/1set/cronrange) | 3 | 0 | 2019-11-10 | 8 months ago | Parses Cron-style time range expressions, checks if the given time is within any ranges. |
| [go-week](https://github.com/stoewer/go-week) | 3 | 3 | 2018-02-23 | 1 month ago | An efficient package to work with ISO8601 week dates. |

## Distributed Systems
        
*Packages that help with building Distributed Systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [kit](https://github.com/go-kit/kit) | 17584 | 674 | 2015-02-03 | 2 days ago | Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc. |
| [go-micro](https://github.com/micro/go-micro) | 13778 | 473 | 2015-01-13 | 1 hour ago | A distributed systems development framework. |
| [grpc-go](https://github.com/grpc/grpc-go) | 11890 | 496 | 2014-12-08 | 10 hours ago | The Go language implementation of gRPC. HTTP/2 based RPC. |
| [micro](https://github.com/micro/micro) | 8248 | 317 | 2015-01-16 | 1 hour ago | A distributed systems runtime for the cloud and beyond. |
| [nats-server](https://github.com/nats-io/nats-server) | 8034 | 375 | 2012-10-29 | 1 hour ago | Lightweight, high performance messaging system for microservices, IoT, and cloud native systems. |
| [rpcx](https://github.com/smallnest/rpcx) | 4843 | 331 | 2016-05-18 | 4 days ago | Distributed pluggable RPC service framework like alibaba Dubbo. |
| [raft](https://github.com/hashicorp/raft) | 3740 | 293 | 2013-11-05 | 5 days ago | Golang implementation of the Raft consensus protocol, by HashiCorp. |
| [tendermint](https://github.com/tendermint/tendermint) | 3683 | 259 | 2014-05-14 | 3 hours ago | High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols. |
| [torrent](https://github.com/anacrolix/torrent) | 3561 | 134 | 2015-01-08 | 17 hours ago | BitTorrent client package. |
| [dragonboat](https://github.com/lni/dragonboat) | 3139 | 136 | 2018-12-23 | 2 hours ago | A feature complete and high performance multi-group Raft library in Go. |
| [krakend](https://github.com/devopsfaith/krakend) | 2931 | 88 | 2016-11-04 | 23 hours ago | Ultra performant API Gateway framework with middlewares. |
| [glow](https://github.com/chrislusf/glow) | 2821 | 147 | 2015-06-14 | 1 year ago | Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go. |
| [gleam](https://github.com/chrislusf/gleam) | 2567 | 150 | 2016-08-26 | 3 months ago | Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed. |
| [emitter](https://github.com/emitter-io/emitter) | 2465 | 96 | 2016-10-29 | 1 month ago | High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love. |
| [liftbridge](https://github.com/liftbridge-io/liftbridge) | 1841 | 76 | 2017-10-13 | 4 days ago | Lightweight, fault-tolerant message streams for NATS. |
| [hprose-golang](https://github.com/hprose/hprose-golang) | 1108 | 95 | 2014-02-14 | 3 weeks ago | Very newbility RPC Library, support 25+ languages now. |
| [ringpop-go](https://github.com/uber/ringpop-go) | 621 | 2403 | 2015-06-05 | 1 year ago | Scalable, fault-tolerant application-layer sharding for Go applications. |
| [gorpc](https://github.com/valyala/gorpc) | 604 | 27 | 2014-11-20 | 10 months ago | Simple, fast and scalable RPC library for high load. |
| [go-health](https://github.com/InVisionApp/go-health) | 563 | 104 | 2017-11-29 | 6 months ago | Library for enabling asynchronous dependency health checks in your service. |
| [rain](https://github.com/cenkalti/rain) | 507 | 14 | 2014-05-21 | 1 week ago | BitTorrent client and library. |
| [digota](https://github.com/digota/digota) | 340 | 25 | 2017-08-14 | 1 year ago | grpc ecommerce microservice. |
| [sleuth](https://github.com/ursiform/sleuth) | 321 | 9 | 2016-04-23 | 2 years ago | Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)). |
| [go-sundheit](https://github.com/AppsFlyer/go-sundheit) | 316 | 8 | 2019-04-08 | 2 months ago | A library built to provide support for defining async service health checks for golang services. |
| [consistent](https://github.com/buraksezer/consistent) | 297 | 10 | 2018-03-25 | 9 months ago | Consistent hashing with bounded loads. |
| [go-jump](https://github.com/dgryski/go-jump) | 296 | 12 | 2014-06-15 | 2 years ago | Port of Google's "Jump" Consistent Hash function. |
| [redislock](https://github.com/bsm/redislock) | 184 | 3 | 2019-06-24 | 1 week ago | Simplified distributed locking implementation using Redis. |
| [dht](https://github.com/anacrolix/dht) | 164 | 12 | 2016-12-14 | 2 months ago | BitTorrent Kademlia DHT implementation. |
| [jsonrpc](https://github.com/ybbus/jsonrpc) | 133 | 9 | 2016-11-10 | 5 months ago | JSON-RPC 2.0 HTTP client implementation. |
| [jsonrpc](https://github.com/osamingo/jsonrpc) | 130 | 6 | 2016-10-28 | 5 months ago | The jsonrpc package helps implement of JSON-RPC 2.0. |
| [celeriac.v1](https://github.com/svcavallar/celeriac.v1) | 60 | 3 | 2015-10-10 | 1 year ago | Library for adding support for interacting and monitoring Celery workers, tasks and events in Go. |
| [doublejump](https://github.com/edwingeng/doublejump) | 54 | 3 | 2018-06-26 | 4 months ago | A revamped Google's jump consistent hash. |
| [semaphore](https://github.com/jexia/semaphore) | 35 | 16 | 2020-02-05 | 5 hours ago | A straightforward (micro) service orchestrator. |
| [maestro](https://github.com/jexia/maestro) | 34 | 16 | 2020-02-05 | 2 weeks ago | A straightforward (micro) service orchestrator. |
| [outboxer](https://github.com/italolelis/outboxer) | 33 | 0 | 2019-02-01 | 12 hours ago | Outboxer is a go library that implements the outbox pattern. |
| [flowgraph](https://github.com/vectaport/flowgraph) | 32 | 1 | 2018-08-29 | 1 month ago | flow-based programming package. |
| [drmaa](https://github.com/dgruber/drmaa) | 29 | 2 | 2013-03-17 | 1 month ago | Job submission library for cluster schedulers based on the DRMAA standard. |
| [go-pdu](https://github.com/pdupub/go-pdu) | 22 | 2 | 2018-10-08 | 1 month ago | A decentralized identity-based social network. |
| [dynatomic](https://github.com/tylfin/dynatomic) | 12 | 0 | 2019-02-08 | 1 year ago | A library for using DynamoDB as an atomic counter. |
| [micro](https://github.com/gmsec/micro) | 6 | 1 | 2020-05-03 | 1 month ago | A Go distributed systems development framework. |
| [go-mysql-lock](https://github.com/sanketplus/go-mysql-lock) | 6 | 1 | 2020-06-06 | 3 weeks ago | MySQL based distributed lock. |
| [consistenthash](https://github.com/mbrostami/consistenthash) | 5 | 1 | 2020-04-22 | 2 months ago | Consistent hashing with configurable replicas. |

## Dynamic DNS
        
*Tools for updating dynamic DNS records.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [godns](https://github.com/TimothyYe/godns) | 636 | 21 | 2014-05-11 | 4 hours ago | A dynamic DNS client tool, supports DNSPod & HE.net, written in Go. |
| [ddns](https://github.com/skibish/ddns) | 154 | 6 | 2017-03-13 | 1 week ago | Personal DDNS client with Digital Ocean Networking DNS as backend. |

## Email
        
*Libraries and tools that implement email creation and sending.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [MailHog](https://github.com/mailhog/MailHog) | 6736 | 128 | 2014-04-16 | 1 week ago | Email and SMTP testing with web and API interface. |
| [hermes](https://github.com/matcornic/hermes) | 2021 | 28 | 2017-03-25 | 1 month ago | Golang package that generates clean, responsive HTML e-mails. |
| [email](https://github.com/jordan-wright/email) | 1454 | 46 | 2013-12-12 | 2 weeks ago | A robust and flexible email library for Go. |
| [go-imap](https://github.com/emersion/go-imap) | 1033 | 34 | 2016-04-26 | 19 hours ago | IMAP library for clients and servers. |
| [sendgrid-go](https://github.com/sendgrid/sendgrid-go) | 624 | 186 | 2013-09-12 | 2 hours ago | SendGrid's Go library for sending email. |
| [mailgun-go](https://github.com/mailgun/mailgun-go) | 468 | 57 | 2014-02-28 | 1 month ago | Go library for sending mail with the Mailgun API. |
| [hectane](https://github.com/hectane/hectane) | 186 | 13 | 2015-08-28 | 11 months ago | Lightweight SMTP client providing an HTTP API. |
| [douceur](https://github.com/aymerick/douceur) | 182 | 2 | 2015-04-09 | 2 years ago | CSS inliner for your HTML emails. |
| [go-message](https://github.com/emersion/go-message) | 165 | 11 | 2016-12-31 | 1 month ago | Streaming library for the Internet Message Format and mail messages. |
| [go-dkim](https://github.com/toorop/go-dkim) | 61 | 2 | 2015-04-29 | 2 months ago | DKIM library, to sign & verify email. |
| [smtp](https://github.com/mailhog/smtp) | 56 | 10 | 2014-12-24 | 2 years ago | SMTP server protocol state machine. |
| [go-premailer](https://github.com/vanng822/go-premailer) | 51 | 2 | 2015-02-16 | 7 months ago | Inline styling for HTML mail in Go. |
| [mailchain](https://github.com/mailchain/mailchain) | 51 | 7 | 2019-04-11 | 22 hours ago | Send encrypted emails to blockchain addresses written in Go. |
| [go-simple-mail](https://github.com/xhit/go-simple-mail) | 33 | 1 | 2019-09-15 | 2 hours ago | Very simple package to send emails with SMTP Keep Alive and two timeouts: Connect and Send. |

## Embeddable Scripting Languages
        
*Embedding other languages inside your go code.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [otto](https://github.com/robertkrimen/otto) | 5441 | 185 | 2012-10-06 | 1 month ago | JavaScript interpreter written in Go. |
| [gopher-lua](https://github.com/yuin/gopher-lua) | 3640 | 148 | 2015-02-15 | 3 weeks ago | Lua 5.1 VM and compiler written in Go. |
| [tengo](https://github.com/d5/tengo) | 1954 | 44 | 2019-01-09 | 1 week ago | Bytecode compiled script language for Go. |
| [go-lua](https://github.com/Shopify/go-lua) | 1887 | 275 | 2013-12-20 | 1 week ago | Port of the Lua 5.2 VM to pure Go. |
| [expr](https://github.com/antonmedv/expr) | 1373 | 36 | 2018-07-14 | 2 months ago | Expression evaluation engine for Go: fast, non-Turing complete, dynamic typing, static typing. |
| [go-python](https://github.com/sbinet/go-python) | 1104 | 46 | 2012-07-09 | 4 months ago | naive go bindings to the CPython C-API. |
| [anko](https://github.com/mattn/anko) | 1038 | 46 | 2014-03-28 | 2 months ago | Scriptable interpreter written in Go. |
| [go-php](https://github.com/deuill/go-php) | 753 | 44 | 2015-09-17 | 1 year ago | PHP bindings for Go. |
| [go-duktape](https://github.com/olebedev/go-duktape) | 722 | 28 | 2015-01-08 | 1 month ago | Duktape JavaScript engine bindings for Go. |
| [cel-go](https://github.com/google/cel-go) | 503 | 27 | 2018-03-09 | 4 days ago | Fast, portable, non-Turing complete expression evaluation with gradual typing. |
| [golua](https://github.com/aarzilli/golua) | 481 | 33 | 2010-12-06 | 6 days ago | Go bindings for Lua C API. |
| [gisp](https://github.com/jcla1/gisp) | 439 | 20 | 2014-01-11 | 2 years ago | Simple LISP in Go. |
| [gval](https://github.com/PaesslerAG/gval) | 223 | 15 | 2017-09-27 | 1 year ago | A highly customizable expression language written in Go. |
| [gentee](https://github.com/gentee/gentee) | 51 | 3 | 2018-01-14 | 1 month ago | Embeddable scripting programming language. |
| [binder](https://github.com/alexeyco/binder) | 39 | 2 | 2017-04-02 | 2 years ago | Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua). |
| [purl](https://github.com/ian-kent/purl) | 29 | 2 | 2014-11-29 | 5 years ago | Perl 5.18.2 embedded in Go. |
| [ngaro](https://github.com/db47h/ngaro) | 20 | 1 | 2016-08-09 | 2 years ago | Embeddable Ngaro VM implementation enabling scripting in Retro. |

## Error Handling
        
*Libraries for handling errors.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [errors](https://github.com/pkg/errors) | 5971 | 112 | 2015-12-27 | 1 month ago | Package that provides simple error handling primitives. |
| [go-multierror](https://github.com/hashicorp/go-multierror) | 898 | 194 | 2014-12-15 | 4 weeks ago | Go (golang) package for representing a list of errors as a single error. |
| [eris](https://github.com/rotisserie/eris) | 629 | 9 | 2019-09-07 | 4 days ago | A better way to handle, trace, and log errors in Go. Compatible with the standard error library and github.com/pkg/errors. |
| [errorx](https://github.com/joomcode/errorx) | 629 | 60 | 2018-08-17 | 3 months ago | A feature rich error package with stack traces, composition of errors and more. |
| [tracerr](https://github.com/ztrue/tracerr) | 622 | 8 | 2019-02-06 | 1 year ago | Golang errors with stack trace and source fragments. |
| [errlog](https://github.com/snwfdhmp/errlog) | 343 | 6 | 2019-02-16 | 5 months ago | Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place. |
| [emperror](https://github.com/emperror/emperror) | 145 | 2 | 2017-06-13 | 3 months ago | Error handling tools and best practices for Go libraries and applications. |
| [errors](https://github.com/emperror/errors) | 58 | 3 | 2019-07-09 | 3 months ago | Drop-in replacement for the standard library errors package and github.com/pkg/errors. Provides various error handling primitives. |
| [werr](https://github.com/txgruppi/werr) | 13 | 0 | 2015-08-04 | 4 years ago | Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called. |
| [errors](https://github.com/bnkamalesh/errors) | 7 | 1 | 2020-07-17 | 1 week ago | Drop-in replacement for builting Go errors. This is a minimal error handling package with custom error types, user friendly messages, Unwrap & Is. With very easy to use and straightforward helper functions. |
| [falcon](https://github.com/SonicRoshan/falcon) | 5 | 2 | 2019-09-09 | 9 months ago | A Simple Yet Highly Powerful Package For Error Handling. |
| [errors](https://github.com/neuronlabs/errors) | 3 | 1 | 2019-07-26 | 1 year ago | Simple golang error handling with classification primitives. |
| [errors](https://github.com/PumpkinSeed/errors) | 2 | 1 | 2020-01-08 | 6 months ago | The most simple error wrapper with awesome performance and minimal memory overhead. |

## Files
        
*Libraries for handling files and file systems.*

## Financial
        
*Packages for accounting and finance.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [decimal](https://github.com/shopspring/decimal) | 2297 | 63 | 2015-02-25 | 3 weeks ago | Arbitrary-precision fixed-point decimal numbers. |
| [go-money](https://github.com/Rhymond/go-money) | 763 | 17 | 2017-03-20 | 2 weeks ago | Implementation of Fowler's Money pattern. |
| [accounting](https://github.com/leekchan/accounting) | 570 | 13 | 2015-08-10 | 7 months ago | money and currency formatting for golang. |
| [go-finance](https://github.com/FlashBoys/go-finance) | 537 | 27 | 2016-02-28 | 2 years ago | Comprehensive financial markets data in Go. |
| [techan](https://github.com/sdcoffey/techan) | 281 | 31 | 2017-03-08 | 7 months ago | Technical analysis library with advanced market analysis and trading strategies. |
| [orderbook](https://github.com/i25959341/orderbook) | 144 | 11 | 2018-04-24 | 1 year ago | Matching Engine for Limit Order Book in Golang. |
| [ofxgo](https://github.com/aclindsa/ofxgo) | 76 | 10 | 2015-11-08 | 3 months ago | Query OFX servers and/or parse the responses (with example command-line client). |
| [vat](https://github.com/dannyvankooten/vat) | 76 | 1 | 2016-06-18 | 3 months ago | VAT number validation & EU VAT rates. |
| [go-finance](https://github.com/alpeb/go-finance) | 74 | 4 | 2017-06-01 | 5 months ago | Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations. |
| [transaction](https://github.com/claygod/transaction) | 74 | 8 | 2017-10-11 | 3 weeks ago | Embedded transactional database of accounts, running in multithreaded mode. |
| [go-finnhub](https://github.com/m1/go-finnhub) | 36 | 3 | 2020-01-13 | 5 months ago | Client for stock market, forex and crypto data from finnhub.io. Access real-time financial market data from 60+ stock exchanges, 10 forex brokers, and 15+ crypto exchanges. |
| [currency](https://github.com/bnkamalesh/currency) | 27 | 5 | 2017-05-09 | 1 month ago | High performant & accurate currency computation package. |
| [go-finance](https://github.com/pieterclaerhout/go-finance) | 3 | 1 | 2019-09-30 | 9 months ago | Module to fetch exchange rates, check VAT numbers via VIES and check IBAN bank account numbers. |

## Forms
        
*Libraries for working with forms.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [nosurf](https://github.com/justinas/nosurf) | 1068 | 35 | 2013-08-22 | 7 months ago | CSRF protection middleware for Go. |
| [binding](https://github.com/mholt/binding) | 771 | 31 | 2014-05-20 | 2 years ago | Binds form and JSON data from net/http Request to struct. |
| [csrf](https://github.com/gorilla/csrf) | 553 | 23 | 2015-08-03 | 1 month ago | CSRF protection for Go web applications & services. |
| [form](https://github.com/go-playground/form) | 418 | 11 | 2016-05-26 | 8 months ago | Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support. |
| [conform](https://github.com/leebenson/conform) | 197 | 5 | 2016-01-05 | 1 month ago | Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags. |
| [formam](https://github.com/monoculum/formam) | 149 | 3 | 2014-10-25 | 2 months ago | decode form's values into a struct. |
| [forms](https://github.com/albrow/forms) | 111 | 7 | 2014-08-07 | 3 years ago | Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files. |
| [bind](https://github.com/robfig/bind) | 24 | 2 | 2014-08-06 | 6 years ago | Bind form data to any Go values. |
| [queryparam](https://github.com/TomWright/queryparam) | 3 | 1 | 2018-06-14 | 7 months ago | Decode `url.Values` into usable struct values of standard or custom types. |

## Functional
        
*Packages to support functional programming in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-underscore](https://github.com/tobyhede/go-underscore) | 1154 | 29 | 2014-07-02 | 1 year ago | Useful collection of helpfully functional Go collection utilities. |
| [fpGo](https://github.com/TeaEntityLab/fpGo) | 152 | 5 | 2018-05-24 | 2 years ago | Monad, Functional Programming features for Golang. |
| [fuego](https://github.com/seborama/fuego) | 68 | 2 | 2018-11-05 | 11 months ago | Functional Experiment in Go. |

## Game Development
        
*Awesome game development libraries.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [leaf](https://github.com/name5566/leaf) | 3613 | 322 | 2014-08-04 | 2 months ago | Lightweight game server framework. |
| [ebiten](https://github.com/hajimehoshi/ebiten) | 3085 | 95 | 2013-06-16 | 11 hours ago | dead simple 2D game library in Go. |
| [pixel](https://github.com/faiface/pixel) | 3068 | 99 | 2016-11-19 | 4 hours ago | Hand-crafted 2D game library in Go. |
| [goworld](https://github.com/xiaonanln/goworld) | 1542 | 123 | 2017-06-03 | 6 days ago | Scalable game server engine, featuring space-entity framework and hot-swapping. |
| [go-sdl2](https://github.com/veandco/go-sdl2) | 1385 | 44 | 2013-06-05 | 2 months ago | Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/). |
| [nano](https://github.com/lonng/nano) | 1328 | 58 | 2017-08-02 | 7 months ago | Lightweight, facility, high performance golang based game server framework. |
| [engo](https://github.com/EngoEngine/engo) | 1235 | 48 | 2014-11-12 | 3 weeks ago | Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm. |
| [termloop](https://github.com/JoelOtter/termloop) | 1149 | 32 | 2015-05-23 | 2 months ago | Terminal-based game engine for Go, built on top of Termbox. |
| [engine](https://github.com/g3n/engine) | 1116 | 68 | 2017-03-07 | 1 week ago | Go 3D Game Engine. |
| [gonet](https://github.com/xtaci/gonet) | 1101 | 134 | 2013-04-11 | 3 years ago | Game server skeleton implemented with golang. |
| [oak](https://github.com/oakmound/oak) | 738 | 44 | 2017-07-15 | 17 hours ago | Pure Go game engine. |
| [pitaya](https://github.com/topfreegames/pitaya) | 662 | 53 | 2018-03-19 | 3 days ago | Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. |
| [raylib-go](https://github.com/gen2brain/raylib-go) | 503 | 24 | 2017-01-27 | 1 month ago | Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming. |
| [engine](https://github.com/azul3d/engine) | 454 | 25 | 2016-02-29 | 9 months ago | 3D game engine written in Go. |
| [go-astar](https://github.com/beefsack/go-astar) | 376 | 10 | 2014-05-28 | 2 years ago | Go implementation of the A\* path finding algorithm. |
| [GarageEngine](https://github.com/vova616/GarageEngine) | 320 | 31 | 2012-08-07 | 11 months ago | 2d game engine written in Go working on OpenGL. |
| [go3d](https://github.com/ungerik/go3d) | 181 | 10 | 2011-06-27 | 5 months ago | Performance oriented 2D/3D math package for Go. |
| [glop](https://github.com/runningwild/glop) | 77 | 3 | 2011-04-20 | 4 years ago | Glop (Game Library Of Power) is a fairly simple cross-platform game library. |
| [prototype](https://github.com/gonutz/prototype) | 22 | 3 | 2015-03-04 | 1 month ago | Cross-platform (Windows/Linux/Mac) library for creating desktop games using a minimal API. |
| [go-collada](https://github.com/GlenKelley/go-collada) | 15 | 3 | 2013-09-19 | 6 years ago | Go package for working with the Collada file format. |

## Generation and Generics
        
*Tools to enhance the language with features like generics via code generation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-linq](https://github.com/ahmetb/go-linq) | 2136 | 71 | 2013-12-19 | 1 month ago | .NET LINQ-like query methods for Go. |
| [jennifer](https://github.com/dave/jennifer) | 1610 | 25 | 2016-12-04 | 2 months ago | Generate arbitrary Go code without templates. |
| [gen](https://github.com/clipperhouse/gen) | 1118 | 34 | 2013-10-13 | 6 months ago | Code generation tool for ‘generics’-like functionality. |
| [goderive](https://github.com/awalterschulze/goderive) | 814 | 24 | 2017-02-10 | 5 months ago | Derives functions from input types. |
| [gowrap](https://github.com/hexdigest/gowrap) | 375 | 11 | 2018-09-15 | 3 weeks ago | Generate decorators for Go interfaces using simple templates. |
| [interfaces](https://github.com/rjeczalik/interfaces) | 221 | 7 | 2015-12-06 | 1 month ago | Command line tool for generating interface definitions. |
| [go-enum](https://github.com/abice/go-enum) | 114 | 3 | 2017-08-10 | 6 months ago | Code generation for enums from code comments. |
| [pkgreflect](https://github.com/ungerik/pkgreflect) | 93 | 5 | 2012-09-03 | 2 years ago | Go preprocessor for package scoped reflection. |
| [efaceconv](https://github.com/t0pep0/efaceconv) | 45 | 4 | 2016-11-18 | 2 years ago | Code generation tool for high performance conversion from interface{} to immutable type without allocations. |
| [gotype](https://github.com/wzshiming/gotype) | 32 | 4 | 2017-12-05 | 2 months ago | Golang source code parsing, usage like reflect package. |
| [GENERIS](https://github.com/senselogic/GENERIS) | 21 | 1 | 2019-03-10 | 1 week ago | Code generation tool providing generics, free-form macros, conditional compilation and HTML templating. |
| [go-xray](https://github.com/pieterclaerhout/go-xray) | 12 | 2 | 2019-10-01 | 8 months ago | Helpers for making the use of reflection easier. |
| [typeregistry](https://github.com/xiaoxin01/typeregistry) | 4 | 1 | 2020-01-14 | 5 months ago | A library to create type dynamically. |

## Geographic
        
*Geographic tools and servers*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [tile38](https://github.com/tidwall/tile38) | 6971 | 205 | 2016-03-04 | 3 weeks ago | Geolocation DB with spatial index and realtime geofencing. |
| [geo](https://github.com/golang/geo) | 1072 | 77 | 2014-12-03 | 3 months ago | S2 geometry library in Go. |
| [mbtileserver](https://github.com/consbio/mbtileserver) | 163 | 13 | 2014-11-01 | 2 months ago | A simple Go-based server for map tiles stored in mbtiles format. |
| [geocache](https://github.com/melihmucuk/geocache) | 124 | 6 | 2016-06-21 | 4 years ago | In-memory cache that is suitable for geolocation based applications. |
| [osm](https://github.com/paulmach/osm) | 115 | 9 | 2016-02-02 | 1 month ago | Library for reading, writing and working with OpenStreetMap data and APIs. |
| [wgs84](https://github.com/wroge/wgs84) | 54 | 1 | 2019-06-08 | 5 months ago | Library for Coordinate Conversion and Transformation (ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM). |
| [geoserver](https://github.com/hishamkaram/geoserver) | 38 | 1 | 2018-03-26 | 1 month ago | geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API. |
| [gismanager](https://github.com/hishamkaram/gismanager) | 32 | 1 | 2018-09-29 | 1 year ago | Publish Your GIS Data(Vector Data) to PostGIS and Geoserver. |
| [pbf](https://github.com/maguro/pbf) | 23 | 2 | 2017-09-18 | 4 weeks ago | OpenStreetMap PBF golang encoder/decoder. |
| [s2-geojson](https://github.com/pantrif/s2-geojson) | 6 | 1 | 2020-03-27 | 3 months ago | Convert geojson to s2 cells & demonstrating some S2 geometry features on map. |

## Go Compilers
        
*Tools for compiling Go to other languages.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gopherjs](https://github.com/gopherjs/gopherjs) | 9518 | 263 | 2013-08-27 | 1 month ago | Compiler from Go to JavaScript. |
| [llgo](https://github.com/go-llvm/llgo) | 1086 | 64 | 2011-11-05 | 5 years ago | LLVM-based compiler for Go. |
| [tardisgo](https://github.com/tardisgo/tardisgo) | 399 | 28 | 2014-01-08 | 3 years ago | Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler. |
| [c4go](https://github.com/Konstantin8105/c4go) | 221 | 15 | 2018-03-28 | 2 months ago | Transpile C code to Go code. |
| [f4go](https://github.com/Konstantin8105/f4go) | 24 | 3 | 2018-07-08 | 8 months ago | Transpile FORTRAN 77 code to Go code. |

## Goroutines
        
*Tools for managing and working with Goroutines.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [ants](https://github.com/panjf2000/ants) | 3984 | 125 | 2018-05-19 | 1 week ago | A high-performance and low-cost goroutine pool in Go. |
| [goworker](https://github.com/benmanns/goworker) | 2433 | 76 | 2013-07-22 | 5 months ago | goworker is a Go-based background worker. |
| [tunny](https://github.com/Jeffail/tunny) | 1683 | 64 | 2014-04-02 | 10 months ago | Goroutine pool for golang. |
| [grpool](https://github.com/ivpusic/grpool) | 576 | 28 | 2015-07-22 | 1 year ago | Lightweight Goroutine pool. |
| [pool](https://github.com/go-playground/pool) | 573 | 14 | 2015-10-28 | 10 months ago | Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation. |
| [workerpool](https://github.com/gammazero/workerpool) | 327 | 9 | 2016-05-17 | 1 week ago | Goroutine pool that limits the concurrency of task execution, not the number of tasks queued. |
| [gowp](https://github.com/xxjwxc/gowp) | 213 | 12 | 2019-09-14 | 1 month ago | gowp is concurrency limiting goroutine pool. |
| [go-floc](https://github.com/workanator/go-floc) | 177 | 8 | 2017-07-03 | 2 years ago | Orchestrate goroutines with ease. |
| [go-flow](https://github.com/kamildrazkiewicz/go-flow) | 137 | 9 | 2016-09-25 | 1 year ago | Control goroutines execution order. |
| [GoSlaves](https://github.com/dgrr/GoSlaves) | 95 | 4 | 2017-09-17 | 11 months ago | Simple and Asynchronous Goroutine pool library. |
| [semaphore](https://github.com/marusama/semaphore) | 88 | 4 | 2017-11-22 | 1 month ago | Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations). |
| [semaphore](https://github.com/kamilsk/semaphore) | 83 | 1 | 2016-10-08 | 3 months ago | Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context. |
| [pond](https://github.com/alitto/pond) | 71 | 3 | 2020-03-21 | 1 month ago | Minimalistic and High-performance goroutine worker pool written in Go. |
| [gpool](https://github.com/sherifabdlnaby/gpool) | 68 | 1 | 2018-12-03 | 7 months ago | manages a resizeable pool of context-aware goroutines to bound concurrency. |
| [breaker](https://github.com/kamilsk/breaker) | 67 | 3 | 2019-02-15 | 1 month ago | Flexible mechanism to make execution flow interruptible. |
| [worker-pool](https://github.com/vardius/worker-pool) | 64 | 5 | 2017-10-04 | 4 months ago | goworker is a Go simple async worker pool. |
| [cyclicbarrier](https://github.com/marusama/cyclicbarrier) | 44 | 3 | 2018-01-11 | 4 weeks ago | CyclicBarrier for golang. |
| [async](https://github.com/StudioSol/async) | 43 | 9 | 2017-06-30 | 3 months ago | A safe way to execute functions asynchronously, recovering them in case of panic. |
| [artifex](https://github.com/mborders/artifex) | 39 | 3 | 2018-10-31 | 8 months ago | Simple in-memory job queue for Golang using worker-based dispatching. |
| [threadpool](https://github.com/shettyh/threadpool) | 38 | 3 | 2017-09-06 | 4 months ago | Golang threadpool implementation. |
| [gollback](https://github.com/vardius/gollback) | 32 | 1 | 2019-05-11 | 1 month ago | asynchronous simple function utilities, for managing execution of closures and callbacks. |
| [routine](https://github.com/x-mod/routine) | 30 | 1 | 2019-03-04 | 5 months ago | go routine control with context, support: Main, Go, Pool and some useful Executors. |
| [Hunch](https://github.com/AaronJan/Hunch) | 27 | 1 | 2019-06-05 | 1 year ago | Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive. |
| [parallel-fn](https://github.com/rafaeljesus/parallel-fn) | 27 | 3 | 2017-06-18 | 2 years ago | Run functions in parallel. |
| [kyoo](https://github.com/dirkaholic/kyoo) | 27 | 2 | 2020-01-06 | 4 months ago | Provides an unlimited job queue and concurrent worker pools. |
| [stl](https://github.com/ssgreg/stl) | 14 | 2 | 2018-06-19 | 4 days ago | Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism. |
| [nursery](https://github.com/arunsworld/nursery) | 14 | 3 | 2019-11-23 | 3 months ago | Structured concurrency in Go. |
| [go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup) | 12 | 1 | 2018-08-08 | 5 months ago | Like `sync.WaitGroup` with error handling and concurrency control. |
| [gohive](https://github.com/loveleshsharma/gohive) | 12 | 3 | 2019-05-31 | 9 months ago | A highly performant and easy to use Goroutine pool for Go. |
| [goccm](https://github.com/zenthangplus/goccm) | 11 | 1 | 2019-08-16 | 1 month ago | Go Concurrency Manager package limits the number of goroutines that allowed to run concurrently. |
| [go-trylock](https://github.com/subchen/go-trylock) | 10 | 1 | 2018-04-26 | 1 week ago | TryLock support on read-write lock for Golang. |
| [conexec](https://github.com/ITcathyh/conexec) | 6 | 2 | 2019-12-24 | 1 month ago | A concurrent toolkit to help execute funcs concurrently in an efficient and safe way.It supports specifying the overall timeout to avoid blocking and uses goroutine pool to improve efficiency. |
| [go-tools](https://github.com/nikhilsaraf/go-tools) | 5 | 2 | 2018-11-14 | 1 year ago | Manage a pool of goroutines using this lightweight library with a simple API. |
| [queue](https://github.com/AnikHasibul/queue) | 5 | 0 | 2018-12-21 | 1 year ago | Gives you a `sync.WaitGroup` like queue group accessibility. Helps you to throttle and limit goroutines, wait for the end of the all goroutines and much more. |
| [hands](https://github.com/duanckham/hands) | 4 | 1 | 2020-04-04 | 3 months ago | A process controller used to control the execution and return strategies of multiple goroutines. |

## GUI
        
*Interaction*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [fyne](https://github.com/fyne-io/fyne) | 10375 | 202 | 2018-02-04 | 7 hours ago | Cross platform native GUIs designed for Go based on Material Design. Supports: Linux, macOS, Windows, BSD, iOS and Android. |
| [qt](https://github.com/therecipe/qt) | 7591 | 304 | 2014-11-19 | 3 weeks ago | Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi). |
| [ui](https://github.com/andlabs/ui) | 7580 | 376 | 2014-02-17 | 1 month ago | Platform-native GUI library for Go. Cross platform. |
| [webview](https://github.com/webview/webview) | 6470 | 216 | 2017-08-19 | 4 days ago | Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux). |
| [webview](https://github.com/zserge/webview) | 6166 | 211 | 2017-08-19 | 2 months ago | Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux). |
| [robotgo](https://github.com/go-vgo/robotgo) | 6029 | 222 | 2016-09-26 | 3 days ago | Go Native cross-platform GUI system automation. Control the mouse, keyboard and other. |
| [walk](https://github.com/lxn/walk) | 4683 | 258 | 2010-09-16 | 2 weeks ago | Windows application library kit for Go. |
| [go-app](https://github.com/maxence-charriere/go-app) | 3873 | 118 | 2016-10-12 | 1 day ago | Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. |
| [go-astilectron](https://github.com/asticode/go-astilectron) | 3338 | 137 | 2017-04-22 | 1 week ago | Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron). |
| [app](https://github.com/maxence-charriere/app) | 3314 | 111 | 2016-10-12 | 5 months ago | Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. |
| [go-sciter](https://github.com/sciter-sdk/go-sciter) | 1806 | 120 | 2015-10-15 | 1 month ago | Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform. |
| [systray](https://github.com/getlantern/systray) | 1242 | 54 | 2014-11-12 | 5 days ago | Cross platform Go library to place an icon and menu in the notification area. |
| [gotk3](https://github.com/gotk3/gotk3) | 1140 | 52 | 2015-08-13 | 5 days ago | Go bindings for GTK3. |
| [gosx-notifier](https://github.com/deckarep/gosx-notifier) | 527 | 13 | 2013-11-25 | 5 months ago | OSX Desktop Notifications library for Go. |
| [gowd](https://github.com/dtylman/gowd) | 268 | 26 | 2017-03-29 | 1 year ago | Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform. |
| [trayhost](https://github.com/shurcooL/trayhost) | 192 | 4 | 2014-04-25 | 9 months ago | Cross-platform Go library to place an icon in the host operating system's taskbar. |
| [go-appindicator](https://github.com/dawidd6/go-appindicator) | 9 | 3 | 2019-05-04 | 4 months ago | Go bindings for libappindicator3 C library. |
| [activity-tracker](https://github.com/prashantgupta24/activity-tracker) | 6 | 1 | 2019-03-12 | 1 year ago | OSX library to notify about any (pluggable) activity on your machine. |
| [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) | 4 | 1 | 2019-03-30 | 1 year ago | OSX Sleep/Wake notifications in golang. |

## Hardware
        
*Libraries, tools, and tutorials for interacting with hardware.*

## Images
        
*Libraries for manipulating images.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gocv](https://github.com/hybridgroup/gocv) | 3385 | 132 | 2017-09-18 | 1 week ago | Go package for computer vision using OpenCV 3.3+. |
| [imaging](https://github.com/disintegration/imaging) | 3206 | 74 | 2012-12-06 | 1 month ago | Simple Go image processing package. |
| [imaginary](https://github.com/h2non/imaginary) | 3159 | 74 | 2015-03-04 | 2 weeks ago | Fast and simple HTTP microservice for image resizing. |
| [bild](https://github.com/anthonynsimon/bild) | 2903 | 62 | 2016-08-01 | 1 month ago | Collection of image processing algorithms in pure Go. |
| [ln](https://github.com/fogleman/ln) | 2704 | 87 | 2016-01-10 | 1 year ago | 3D line art rendering in Go. |
| [resize](https://github.com/nfnt/resize) | 2406 | 79 | 2012-08-02 | 2 years ago | Image resizing for Go with common interpolation methods. |
| [gg](https://github.com/fogleman/gg) | 2343 | 80 | 2016-02-18 | 2 months ago | 2D rendering in pure Go. |
| [pt](https://github.com/fogleman/pt) | 1868 | 59 | 2015-01-23 | 1 year ago | Path tracing engine written in Go. |
| [svgo](https://github.com/ajstarks/svgo) | 1504 | 50 | 2010-03-05 | 3 days ago | Go Language Library for SVG generation. |
| [smartcrop](https://github.com/muesli/smartcrop) | 1382 | 32 | 2014-04-07 | 3 months ago | Finds good crops for arbitrary images and crop sizes. |
| [gift](https://github.com/disintegration/gift) | 1351 | 51 | 2014-07-12 | 11 months ago | Package of image processing filters. |
| [picfit](https://github.com/thoas/picfit) | 1282 | 49 | 2014-12-06 | 5 days ago | An image resizing server written in Go. |
| [imagick](https://github.com/gographics/imagick) | 1194 | 55 | 2013-04-30 | 3 months ago | Go binding to ImageMagick's MagickWand C API. |
| [go-opencv](https://github.com/go-opencv/go-opencv) | 1184 | 67 | 2013-12-09 | 1 year ago | Go bindings for OpenCV. |
| [bimg](https://github.com/h2non/bimg) | 1147 | 34 | 2015-03-17 | 2 weeks ago | Small package for fast and efficient image processing using libvips. |
| [geopattern](https://github.com/pravj/geopattern) | 1085 | 21 | 2014-10-22 | 1 year ago | Create beautiful generative image patterns from a string. |
| [stegify](https://github.com/DimitarPetrov/stegify) | 856 | 21 | 2018-11-29 | 2 weeks ago | Go tool for LSB steganography, capable of hiding any file within an image. |
| [canvas](https://github.com/tdewolff/canvas) | 497 | 11 | 2017-05-20 | 3 days ago | Vector graphics to PDF, SVG or rasterized image. |
| [image2ascii](https://github.com/qeesung/image2ascii) | 423 | 9 | 2018-10-20 | 1 year ago | Convert image to ASCII. |
| [mort](https://github.com/aldor007/mort) | 406 | 19 | 2017-11-19 | 5 months ago | Storage and image processing server written in Go. |
| [draft](https://github.com/lucasepe/draft) | 392 | 10 | 2020-06-05 | 3 weeks ago | Generate High Level Microservice Architecture diagrams for GraphViz using simple YAML syntax. |
| [govatar](https://github.com/o1egl/govatar) | 381 | 7 | 2016-01-18 | 5 months ago | Library and CMD tool for generating funny avatars. |
| [goimagehash](https://github.com/corona10/goimagehash) | 325 | 9 | 2017-07-28 | 5 months ago | Go Perceptual image hashing package. |
| [go-nude](https://github.com/koyachi/go-nude) | 311 | 15 | 2014-05-02 | 1 year ago | Nudity detection with Go. |
| [rez](https://github.com/bamiaux/rez) | 201 | 9 | 2014-01-16 | 3 years ago | Image resizing in pure Go and SIMD. |
| [darkroom](https://github.com/gojek/darkroom) | 138 | 7 | 2019-07-01 | 2 weeks ago | An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency. |
| [img](https://github.com/hawx/img) | 133 | 5 | 2012-07-28 | 5 years ago | Selection of image manipulation tools. |
| [mergi](https://github.com/noelyahan/mergi) | 111 | 5 | 2018-09-24 | 1 month ago | Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate). |
| [go-cairo](https://github.com/ungerik/go-cairo) | 94 | 6 | 2012-08-22 | 9 months ago | Go binding for the cairo graphics library. |
| [gltf](https://github.com/qmuntal/gltf) | 73 | 1 | 2019-01-15 | 1 week ago | Efficient and robust glTF 2.0 reader, writer and validator. |
| [steganography](https://github.com/auyer/steganography) | 70 | 4 | 2018-05-21 | 5 months ago | Pure Go Library for LSB steganography. |
| [cameron](https://github.com/aofei/cameron) | 55 | 3 | 2018-05-05 | 4 days ago | An avatar generator for Go. |
| [go-gd](https://github.com/bolknote/go-gd) | 52 | 4 | 2011-05-12 | 2 years ago | Go binding for GD library. |
| [goimghdr](https://github.com/corona10/goimghdr) | 35 | 1 | 2018-02-25 | 1 year ago | The imghdr module determines the type of image contained in a file for Go. |
| [gridder](https://github.com/shomali11/gridder) | 26 | 2 | 2020-04-10 | 2 months ago | A Grid based 2D Graphics library. |
| [tga](https://github.com/ftrvxmtrx/tga) | 25 | 3 | 2012-10-08 | 5 years ago | Package tga is a TARGA image format decoder/encoder. |
| [go-webcolors](https://github.com/jyotiska/go-webcolors) | 24 | 2 | 2014-04-24 | 5 years ago | Port of webcolors library from Python to Go. |
| [mpo](https://github.com/donatj/mpo) | 6 | 1 | 2015-04-14 | 1 month ago | Decoder and conversion tool for MPO 3D Photos. |

## IoT
        
*Libraries for programming devices of the IoT.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [flogo](https://github.com/TIBCOSoftware/flogo) | 1507 | 137 | 2016-07-10 | 2 months ago | Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. |
| [mainflux](https://github.com/mainflux/mainflux) | 972 | 83 | 2015-07-06 | 7 hours ago | Industrial IoT Messaging and Device Management Server. |
| [gatt](https://github.com/paypal/gatt) | 904 | 53 | 2014-04-23 | 1 week ago | Gatt is a Go package for building Bluetooth Low Energy peripherals. |
| [devices](https://github.com/goiot/devices) | 235 | 17 | 2016-05-30 | 4 years ago | Suite of libraries for IoT devices, experimental for x/exp/io. |
| [heedy](https://github.com/heedy/heedy) | 225 | 21 | 2015-01-16 | 1 week ago | Open-Source Platform for Quantified Self & IoT. |
| [sensorbee](https://github.com/sensorbee/sensorbee) | 193 | 18 | 2016-02-19 | 8 months ago | Lightweight stream processing engine for IoT. |
| [huego](https://github.com/amimof/huego) | 153 | 3 | 2017-05-16 | 2 months ago | An extensive Philips Hue client library for Go. |
| [eywa](https://github.com/xcodersun/eywa) | 45 | 8 | 2016-02-20 | 3 years ago | Project Eywa is essentially a connection manager that keeps track of connected devices. |

## Job Scheduler
        
*Libraries for scheduling jobs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gron](https://github.com/roylee0704/gron) | 775 | 15 | 2016-06-04 | 1 week ago | Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly. |
| [jobrunner](https://github.com/bamzi/jobrunner) | 696 | 23 | 2015-10-21 | 7 months ago | Smart and featureful cron job scheduler with job queuing and live monitoring built in. |
| [jobs](https://github.com/albrow/jobs) | 474 | 19 | 2015-02-09 | 2 years ago | Persistent and flexible background jobs library. |
| [scheduler](https://github.com/carlescere/scheduler) | 339 | 15 | 2015-02-03 | 2 years ago | Cronjobs scheduling made easy. |
| [go-cron](https://github.com/rk/go-cron) | 194 | 9 | 2011-04-15 | 5 months ago | Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. |
| [leprechaun](https://github.com/kilgaloon/leprechaun) | 70 | 7 | 2018-04-08 | 5 months ago | Job scheduler that supports webhooks, crons and classic scheduling. |
| [clockwork](https://github.com/whiteShtef/clockwork) | 4 | 1 | 2018-04-23 | 5 months ago | Simple and intuitive job scheduling library in Go. |

## JSON
        
*Libraries for working with JSON.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gjson](https://github.com/tidwall/gjson) | 6769 | 142 | 2016-08-11 | 1 day ago | Get a JSON value with one line of code. |
| [gojson](https://github.com/ChimeraCoder/gojson) | 2238 | 44 | 2012-12-27 | 1 month ago | Automatically generate Go (golang) struct definitions from example JSON. |
| [kazaam](https://github.com/qntfy/kazaam) | 161 | 21 | 2016-07-19 | 3 months ago | API for arbitrary transformation of JSON documents. |
| [gojq](https://github.com/elgs/gojq) | 151 | 4 | 2015-12-30 | 2 years ago | JSON query in Golang. |
| [jsongo](https://github.com/ricardolonga/jsongo) | 94 | 1 | 2015-08-07 | 3 years ago | Fluent API to make it easier to create Json objects. |
| [jettison](https://github.com/wI2L/jettison) | 79 | 4 | 2019-08-30 | 4 months ago | High performance, reflection-less JSON encoder for Go. |
| [gjo](https://github.com/skanehira/gjo) | 76 | 8 | 2019-02-23 | 3 months ago | Small utility to create JSON objects. |
| [jaydiff](https://github.com/yazgazan/jaydiff) | 71 | 1 | 2017-04-24 | 10 months ago | JSON diff utility written in Go. |
| [jsonf](https://github.com/miolini/jsonf) | 57 | 3 | 2015-05-25 | 4 years ago | Console tool for highlighted formatting and struct query fetching JSON. |
| [json2go](https://github.com/m-zajac/json2go) | 52 | 2 | 2017-06-10 | 1 month ago | Advanced JSON to Go struct conversion. Provides package that can parse multiple JSON documents and create struct to fit them all. |
| [ajson](https://github.com/spyzhov/ajson) | 43 | 1 | 2019-03-07 | 1 week ago | Abstract JSON for golang with JSONPath support. |
| [mp](https://github.com/sanbornm/mp) | 40 | 2 | 2014-06-15 | 4 years ago | Simple cli email parser. It currently takes stdin and outputs JSON. |
| [go-respond](https://github.com/nicklaw5/go-respond) | 35 | 1 | 2017-03-12 | 1 year ago | Go package for handling common HTTP JSON responses. |
| [go-jsonerror](https://github.com/ddymko/go-jsonerror) | 9 | 1 | 2018-10-18 | 9 months ago | Go-JsonError is ment to allow us to easily create json response errors that follow the JsonApi spec. |
| [jsonhal](https://github.com/RichardKnop/jsonhal) | 9 | 2 | 2016-01-15 | 4 months ago | Simple Go package to make custom structs marshal into HAL compatible JSON responses. |
| [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) | 8 | 1 | 2016-07-08 | 3 years ago | Go bindings based on the JSON API errors reference. |
| [ej](https://github.com/lucassscaravelli/ej) | 4 | 1 | 2020-01-04 | 3 months ago | Write and read JSON from different sources succinctly. |
| [mapslice-json](https://github.com/mickep76/mapslice-json) | 1 | 1 | 2020-02-19 | 5 months ago | Go MapSlice for ordered marshal/ unmarshal of maps in JSON. |

## Logging
        
*Libraries for generating and working with log files.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [logrus](https://github.com/sirupsen/logrus) | 15468 | 313 | 2013-10-16 | 1 month ago | Structured logger for Go. |
| [zap](https://github.com/uber-go/zap) | 10288 | 241 | 2016-02-18 | 15 hours ago | Fast, structured, leveled logging in Go. |
| [go-spew](https://github.com/davecgh/go-spew) | 3969 | 61 | 2013-01-09 | 1 month ago | Implements a deep pretty printer for Go data structures to aid in debugging. |
| [zerolog](https://github.com/rs/zerolog) | 3597 | 54 | 2017-05-12 | 1 month ago | Zero-allocation JSON logger. |
| [glog](https://github.com/golang/glog) | 2567 | 88 | 2013-07-16 | 2 months ago | Leveled execution logs for Go. |
| [lumberjack](https://github.com/natefinch/lumberjack) | 2027 | 53 | 2014-06-14 | 1 month ago | Simple rolling logger, implements io.WriteCloser. |
| [tail](https://github.com/hpcloud/tail) | 1837 | 97 | 2013-02-05 | 3 months ago | Go package striving to emulate the features of the BSD tail program. |
| [seelog](https://github.com/cihub/seelog) | 1480 | 90 | 2011-11-17 | 1 year ago | Logging functionality with flexible dispatching, filtering, and formatting. |
| [log15](https://github.com/inconshreveable/log15) | 972 | 26 | 2014-05-20 | 3 months ago | Simple, powerful logging for Go. |
| [log](https://github.com/apex/log) | 965 | 36 | 2015-12-21 | 5 days ago | Structured logging package for Go. |
| [onelog](https://github.com/francoispqt/onelog) | 379 | 10 | 2018-05-06 | 1 year ago | Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenario. Also, it is one of the logger with the lowest allocation. |
| [logxi](https://github.com/mgutz/logxi) | 346 | 10 | 2015-03-01 | 3 months ago | 12-factor app logger that is fast and makes you happy. |
| [logutils](https://github.com/hashicorp/logutils) | 287 | 204 | 2013-10-09 | 1 month ago | Utilities for slightly better logging in Go (Golang) extending the standard logger. |
| [log](https://github.com/go-playground/log) | 277 | 10 | 2016-02-07 | 8 months ago | Simple, configurable and scalable Structured Logging for Go. |
| [go-logger](https://github.com/apsdehal/go-logger) | 259 | 7 | 2014-09-26 | 1 year ago | Simple logger of Go Programs, with level handlers. |
| [httpretty](https://github.com/henvic/httpretty) | 186 | 4 | 2020-01-24 | 4 months ago | Pretty-prints your regular HTTP requests on your terminal for debugging (similar to http.DumpRequest). |
| [rollingwriter](https://github.com/arthurkiller/rollingwriter) | 147 | 7 | 2017-02-12 | 5 months ago | RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation. |
| [logger](https://github.com/azer/logger) | 143 | 5 | 2014-09-30 | 4 months ago | Minimalistic logging library for Go. |
| [xlog](https://github.com/rs/xlog) | 134 | 8 | 2015-10-22 | 6 months ago | Structured logger for `net/context` aware HTTP handlers with flexible dispatching. |
| [ozzo-log](https://github.com/go-ozzo/ozzo-log) | 113 | 10 | 2015-10-22 | 2 years ago | High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail). |
| [sqldb-logger](https://github.com/simukti/sqldb-logger) | 107 | 3 | 2019-11-02 | 11 hours ago | A logger for Go SQL database driver without modify existing *sql.DB stdlib usage. |
| [logur](https://github.com/logur/logur) | 89 | 5 | 2018-12-09 | 3 months ago | An opinionated logger interface and collection of logging best practices with adapters and integrations for well-known libraries ([logrus](https://github.com/sirupsen/logrus), [go-kit log](https://github.com/go-kit/kit/tree/master/log), [zap](https://github.com/uber-go/zap), [zerolog](https://github.com/rs/zerolog), etc). |
| [logvoyage](https://github.com/firstrow/logvoyage) | 88 | 5 | 2015-03-29 | 3 years ago | Full-featured logging saas written in golang. |
| [glg](https://github.com/kpango/glg) | 78 | 5 | 2017-06-21 | 1 month ago | glg is simple and fast leveled logging library for Go. |
| [log](https://github.com/alexcesaro/log) | 44 | 6 | 2014-04-19 | 4 years ago | Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs. |
| [gologger](https://github.com/sadlil/gologger) | 38 | 5 | 2015-09-02 | 2 years ago | Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch. |
| [go-log](https://github.com/ian-kent/go-log) | 36 | 2 | 2014-05-02 | 2 years ago | Log4j implementation in Go. |
| [logex](https://github.com/chzyer/logex) | 35 | 7 | 2014-10-10 | 3 years ago | Golang log lib, supports tracking and level, wrap by standard log lib. |
| [go-cronowriter](https://github.com/utahta/go-cronowriter) | 29 | 1 | 2017-02-04 | 3 months ago | Simple writer that rotate log files automatically based on current date and time, like cronolog. |
| [go-log](https://github.com/siddontang/go-log) | 27 | 6 | 2014-05-18 | 1 year ago | Log lib supports level and multi handlers. |
| [logrusly](https://github.com/sebest/logrusly) | 25 | 4 | 2014-09-11 | 2 years ago | [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/). |
| [distillog](https://github.com/amoghe/distillog) | 23 | 1 | 2015-10-12 | 2 years ago | distilled levelled logging (think of it as stdlib + log levels). |
| [log](https://github.com/teris-io/log) | 23 | 1 | 2017-10-28 | 2 years ago | Structured log interface for Go cleanly separates logging facade from its implementation. |
| [journald](https://github.com/ssgreg/journald) | 22 | 2 | 2017-08-23 | 1 year ago | Go implementation of systemd Journal's native API for logging. |
| [mlog](https://github.com/jbrodriguez/mlog) | 22 | 1 | 2014-10-20 | 2 years ago | Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output. |
| [gomol](https://github.com/aphistic/gomol) | 16 | 2 | 2015-08-30 | 1 year ago | Multiple-output, structured logging for Go with extensible logging outputs. |
| [glo](https://github.com/lajosbencz/glo) | 13 | 1 | 2019-01-19 | 1 year ago | PHP Monolog inspired logging facility with identical severity levels. |
| [go-log](https://github.com/subchen/go-log) | 11 | 2 | 2017-05-07 | 2 years ago | Simple and configurable Logging in Go, with level, formatters and writers. |
| [logrusiowriter](https://github.com/cabify/logrusiowriter) | 10 | 87 | 2019-08-09 | 1 week ago | `io.Writer` implementation using [logrus](https://github.com/sirupsen/logrus) logger. |
| [logdump](https://github.com/ewwwwwqm/logdump) | 9 | 2 | 2017-01-13 | 2 years ago | Package for multi-level logging. |
| [logmatic](https://github.com/borderstech/logmatic) | 9 | 2 | 2018-11-07 | 1 year ago | Colorized logger for Golang with dynamic log level configuration. |
| [go-log](https://github.com/pieterclaerhout/go-log) | 8 | 1 | 2019-10-01 | 2 weeks ago | A logging library with strack traces, object dumping and optional timestamps. |
| [log](https://github.com/aerogo/log) | 7 | 1 | 2017-06-10 | 9 months ago | An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection). |
| [logo](https://github.com/mbndr/logo) | 6 | 1 | 2017-02-07 | 2 years ago | Golang logger to different configurable writers. |
| [xlog](https://github.com/xfxdev/xlog) | 6 | 1 | 2016-05-05 | 1 year ago | Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format. |

## Machine Learning
        
*Libraries for Machine Learning.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [golearn](https://github.com/sjwhitworth/golearn) | 7286 | 438 | 2013-12-26 | 2 days ago | General Machine Learning library for Go. |
| [gorgonia](https://github.com/gorgonia/gorgonia) | 3466 | 178 | 2016-09-14 | 5 days ago | graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. |
| [tfgo](https://github.com/galeone/tfgo) | 1390 | 47 | 2017-05-23 | 4 weeks ago | Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python. |
| [gosseract](https://github.com/otiai10/gosseract) | 1224 | 43 | 2013-10-11 | 2 months ago | Go package for OCR (Optical Character Recognition), by using Tesseract C++ library. |
| [goml](https://github.com/cdipaolo/goml) | 1115 | 75 | 2015-06-27 | 1 year ago | On-line Machine Learning in Go. |
| [gorse](https://github.com/zhenghaoz/gorse) | 996 | 42 | 2018-08-14 | 3 days ago | An offline recommender system backend based on collaborative filtering written in Go. |
| [eaopt](https://github.com/MaxHalford/eaopt) | 684 | 28 | 2016-01-31 | 5 months ago | An evolutionary optimization library. |
| [bayesian](https://github.com/jbrukh/bayesian) | 673 | 32 | 2011-11-23 | 3 days ago | Naive Bayesian Classification for Golang. |
| [CloudForest](https://github.com/ryanbressler/CloudForest) | 672 | 44 | 2012-10-22 | 1 year ago | Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go. |
| [gobrain](https://github.com/goml/gobrain) | 445 | 26 | 2014-04-29 | 1 month ago | Neural Networks written in go. |
| [ocrserver](https://github.com/otiai10/ocrserver) | 314 | 11 | 2015-11-15 | 1 month ago | A simple OCR API server, seriously easy to be deployed by Docker and Heroku. |
| [go-deep](https://github.com/patrikeh/go-deep) | 277 | 13 | 2017-12-09 | 7 months ago | A feature-rich neural network library in Go. |
| [regommend](https://github.com/muesli/regommend) | 273 | 16 | 2014-02-05 | 11 months ago | Recommendation & collaborative filtering engine. |
| [onnx-go](https://github.com/owulveryck/onnx-go) | 267 | 11 | 2018-08-28 | 1 month ago | Go Interface to Open Neural Network Exchange (ONNX). |
| [go-galib](https://github.com/thoj/go-galib) | 180 | 13 | 2009-11-30 | 4 years ago | Genetic Algorithms library written in Go / golang. |
| [goRecommend](https://github.com/timkaye11/goRecommend) | 160 | 9 | 2014-07-16 | 6 years ago | Recommendation Algorithms library written in Go. |
| [goptuna](https://github.com/c-bata/goptuna) | 137 | 5 | 2019-07-24 | 5 hours ago | Bayesian optimization framework for black-box functions written in Go. Everything will be optimized. |
| [shield](https://github.com/eaigner/shield) | 134 | 10 | 2013-04-10 | 4 months ago | Bayesian text classifier with flexible tokenizers and storage backends for Go. |
| [go-fann](https://github.com/vksnk/go-fann) | 102 | 7 | 2011-03-10 | 5 years ago | Go bindings for Fast Artificial Neural Networks(FANN) library. |
| [goga](https://github.com/tomcraven/goga) | 89 | 6 | 2015-10-20 | 3 years ago | Genetic algorithm library for Go. |
| [libsvm](https://github.com/datastream/libsvm) | 65 | 10 | 2012-07-31 | 4 years ago | libsvm golang version derived work based on LIBSVM 3.14. |
| [neural-go](https://github.com/schuyler/neural-go) | 62 | 2 | 2011-10-17 | 6 years ago | Multilayer perceptron network implemented in Go, with training via backpropagation. |
| [go-pr](https://github.com/daviddengcn/go-pr) | 59 | 6 | 2013-06-07 | 7 years ago | Pattern recognition package in Go lang. |
| [gonet](https://github.com/dathoangnd/gonet) | 57 | 4 | 2020-01-11 | 3 months ago | Neural Network for Go. |
| [neat](https://github.com/jinyeom/neat) | 56 | 12 | 2016-11-17 | 2 years ago | Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT). |
| [goscore](https://github.com/asafschers/goscore) | 53 | 5 | 2017-08-19 | 11 months ago | Go Scoring API for PMML. |
| [fonet](https://github.com/Fontinalis/fonet) | 41 | 3 | 2017-10-03 | 3 months ago | A Deep Neural Network library written in Go. |
| [golinear](https://github.com/danieldk/golinear) | 40 | 5 | 2013-04-05 | 1 year ago | liblinear bindings for Go. |
| [Varis](https://github.com/Xamber/Varis) | 29 | 5 | 2017-10-10 | 2 years ago | Golang Neural Network. |
| [go-cluster](https://github.com/e-XpertSolutions/go-cluster) | 26 | 5 | 2017-10-04 | 2 years ago | Go implementation of the k-modes and k-prototypes clustering algorithms. |
| [godist](https://github.com/e-dard/godist) | 25 | 2 | 2014-09-05 | 5 years ago | Various probability distributions, and associated methods. |
| [evoli](https://github.com/khezen/evoli) | 12 | 4 | 2015-06-12 | 2 months ago | Genetic Algorithm and Particle Swarm Optimization library. |
| [probab](https://github.com/ThePaw/probab) | 12 | 1 | 2015-09-14 | 4 years ago | Probability distribution functions. Bayesian inference. Written in pure Go. |
| [gomind](https://github.com/surenderthakran/gomind) | 7 | 2 | 2017-10-19 | 2 years ago | A simplistic Neural Network Library in Go. |
| [randomForest](https://github.com/malaschitz/randomForest) | 4 | 1 | 2018-10-25 | 7 months ago | Easy to use Random Forest library for Go. |

## Messaging
        
*Libraries that implement messaging systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [sarama](https://github.com/Shopify/sarama) | 6053 | 408 | 2013-07-05 | 21 hours ago | Go library for Apache Kafka. |
| [gorush](https://github.com/appleboy/gorush) | 4853 | 188 | 2016-03-22 | 1 week ago | Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm). |
| [centrifugo](https://github.com/centrifugal/centrifugo) | 4452 | 191 | 2015-03-31 | 3 days ago | Real-time messaging (Websockets or SockJS) server in Go. |
| [machinery](https://github.com/RichardKnop/machinery) | 4280 | 142 | 2015-04-05 | 1 week ago | Asynchronous task queue/job queue based on distributed message passing. |
| [go-socket.io](https://github.com/googollee/go-socket.io) | 3599 | 130 | 2013-07-13 | 4 hours ago | socket.io library for golang, a realtime application framework. |
| [nats.go](https://github.com/nats-io/nats.go) | 2963 | 162 | 2012-08-15 | 1 week ago | Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library. |
| [benthos](https://github.com/Jeffail/benthos) | 2416 | 70 | 2016-03-22 | 1 day ago | A message streaming bridge between a range of protocols. |
| [apns2](https://github.com/sideshow/apns2) | 2349 | 73 | 2016-01-05 | 6 months ago | HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps. |
| [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go) | 2162 | 259 | 2016-07-12 | 1 week ago | confluent-kafka-go is Confluent's Golang client for Apache Kafka and the Confluent Platform. |
| [mercure](https://github.com/dunglas/mercure) | 2086 | 58 | 2018-07-14 | 4 hours ago | Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events). |
| [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) | 1921 | 232 | 2013-12-27 | 3 years ago | gopush-cluster is a go push server cluster. |
| [melody](https://github.com/olahol/melody) | 1883 | 59 | 2015-05-13 | 6 months ago | Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling. |
| [go-nsq](https://github.com/nsqio/go-nsq) | 1727 | 65 | 2013-08-29 | 1 month ago | the official Go package for NSQ. |
| [mangos-v1](https://github.com/nanomsg/mangos-v1) | 1537 | 84 | 2014-10-25 | 8 months ago | Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability. |
| [uniqush-push](https://github.com/uniqush/uniqush-push) | 1182 | 78 | 2011-08-29 | 3 months ago | Redis backed unified push service for server-side notifications to mobile devices. |
| [zmq4](https://github.com/pebbe/zmq4) | 860 | 45 | 2013-10-18 | 1 month ago | Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2). |
| [Beaver](https://github.com/Clivern/Beaver) | 857 | 21 | 2018-10-20 | 1 week ago | A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps. |
| [gollum](https://github.com/trivago/gollum) | 853 | 39 | 2015-06-20 | 9 months ago | A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations. |
| [EventBus](https://github.com/asaskevich/EventBus) | 748 | 27 | 2014-12-19 | 3 months ago | The lightweight event bus with async compatibility. |
| [asynq](https://github.com/hibiken/asynq) | 627 | 14 | 2019-11-15 | 3 hours ago | A simple, reliable, and efficient distributed task queue for Go built on top of Redis. |
| [golongpoll](https://github.com/jcuga/golongpoll) | 473 | 23 | 2015-11-02 | 4 months ago | HTTP longpoll server library that makes web pub-sub simple. |
| [dbus](https://github.com/godbus/dbus) | 472 | 15 | 2014-03-27 | 2 weeks ago | Native Go bindings for D-Bus. |
| [emitter](https://github.com/olebedev/emitter) | 367 | 9 | 2015-11-10 | 5 months ago | Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins. |
| [glue](https://github.com/desertbit/glue) | 341 | 14 | 2015-06-07 | 2 months ago | Robust Go and Javascript Socket Library (Alternative to Socket.io). |
| [pubsub](https://github.com/cskr/pubsub) | 325 | 8 | 2012-04-01 | 1 week ago | Simple pubsub package for go. |
| [bus](https://github.com/mustafaturan/bus) | 153 | 5 | 2019-04-27 | 5 months ago | Minimalist message bus implementation for internal communication. |
| [guble](https://github.com/smancke/guble) | 144 | 13 | 2015-11-15 | 2 years ago | Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence. |
| [message-bus](https://github.com/vardius/message-bus) | 137 | 6 | 2017-10-04 | 3 weeks ago | messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD. |
| [rabtap](https://github.com/jandelgado/rabtap) | 131 | 9 | 2017-11-11 | 1 month ago | RabbitMQ swiss army knife cli app. |
| [oplog](https://github.com/dailymotion/oplog) | 104 | 92 | 2014-11-06 | 4 years ago | Generic oplog/replication system for REST APIs. |
| [rabbus](https://github.com/rafaeljesus/rabbus) | 73 | 9 | 2017-05-07 | 1 year ago | A tiny wrapper over amqp exchanges and queues. |
| [drone-line](https://github.com/appleboy/drone-line) | 68 | 5 | 2016-09-13 | 3 weeks ago | Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI. |
| [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) | 61 | 9 | 2017-01-15 | 2 years ago | A tiny wrapper around NSQ topic and channel. |
| [RapidMQ](https://github.com/sybrexsys/RapidMQ) | 59 | 5 | 2016-10-04 | 2 years ago | RapidMQ is a lightweight and reliable library for managing of the local messages queue. |
| [hub](https://github.com/leandro-lugaresi/hub) | 58 | 2 | 2018-04-13 | 3 days ago | A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges. |
| [go-notify](https://github.com/TheCreeper/go-notify) | 48 | 2 | 2015-03-01 | 1 year ago | Native implementation of the freedesktop notification spec. |
| [commander](https://github.com/jeroenrinzema/commander) | 47 | 1 | 2018-04-20 | 2 months ago | A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka. |
| [go-mq](https://github.com/cheshir/go-mq) | 40 | 4 | 2017-06-19 | 5 days ago | RabbitMQ client with declarative configuration. |
| [commander](https://github.com/CloudProud/commander) | 38 | 1 | 2018-04-20 | 7 months ago | A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka. |
| [event](https://github.com/agoalofalife/event) | 36 | 4 | 2017-07-02 | 2 years ago | Implementation of the pattern observer. |
| [go-res](https://github.com/jirenius/go-res) | 35 | 2 | 2018-07-15 | 1 month ago | Package for building REST/real-time services where clients are synchronized seamlessly, using NATS and Resgate. |
| [redisqueue](https://github.com/robinjoseph08/redisqueue) | 25 | 2 | 2019-07-07 | 2 months ago | redisqueue provides a producer and consumer of a queue that uses Redis streams. |
| [go-vitotrol](https://github.com/maxatome/go-vitotrol) | 14 | 6 | 2016-11-03 | 1 year ago | Client library to Viessmann Vitotrol web service. |
| [jazz](https://github.com/socifi/jazz) | 11 | 3 | 2018-10-22 | 1 year ago | A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages. |
| [rmqconn](https://github.com/sbabiv/rmqconn) | 9 | 0 | 2019-01-14 | 6 months ago | RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed. |
| [gaurun-client](https://github.com/osamingo/gaurun-client) | 8 | 1 | 2017-06-29 | 1 year ago | Gaurun Client written in Go. |
| [ami](https://github.com/kak-tus/ami) | 8 | 1 | 2018-10-27 | 3 months ago | Go client to reliable queues based on Redis Cluster Streams. |

## Microsoft Office
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [unioffice](https://github.com/unidoc/unioffice) | 2430 | 63 | 2017-08-29 | 1 week ago | Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents. |

### Microsoft Excel
        
*Libraries for working with Microsoft Excel.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [excelize](https://github.com/360EntSecGroup-Skylar/excelize) | 6529 | 172 | 2016-08-29 | 6 days ago | Golang library for reading and writing Microsoft Excel™ (XLSX) files. |
| [xlsx](https://github.com/tealeg/xlsx) | 4375 | 185 | 2011-06-28 | 1 week ago | Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs. |
| [xlsx](https://github.com/plandem/xlsx) | 118 | 12 | 2017-08-26 | 7 months ago | Fast and safe way to read/update your existing Microsoft Excel files in Go programs. |
| [go-excel](https://github.com/szyhf/go-excel) | 89 | 3 | 2017-09-03 | 6 months ago | A simple and light reader to read a relate-db-like excel as a table. |
| [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) | 15 | 2 | 2017-03-13 | 2 years ago | Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files. |

## Miscellaneous
        

### Dependency Injection
        
*Libraries for working with dependency injection.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [dig](https://github.com/uber-go/dig) | 1489 | 41 | 2017-03-21 | 1 month ago | A reflection based dependency injection toolkit for Go. |
| [fx](https://github.com/uber-go/fx) | 1394 | 56 | 2016-10-27 | 1 month ago | A dependency injection based application framework for Go (built on top of dig). |
| [container](https://github.com/golobby/container) | 94 | 3 | 2019-09-23 | 4 days ago | A powerful IoC Container with fluent and easy-to-use interface. |
| [dingo](https://github.com/i-love-flamingo/dingo) | 73 | 24 | 2018-10-29 | 1 month ago | A dependency injection toolkit for Go, based on Guice. |
| [inject](https://github.com/defval/inject) | 53 | 1 | 2019-02-03 | 5 months ago | A reflection based dependency injection container with simple interface. |
| [alice](https://github.com/magic003/alice) | 42 | 2 | 2017-04-08 | 3 years ago | Additive dependency injection container for Golang. |
| [wire](https://github.com/Fs02/wire) | 31 | 1 | 2018-07-05 | 5 months ago | Strict Runtime Dependency Injection for Golang. |
| [di](https://github.com/goava/di) | 29 | 4 | 2020-02-03 | 1 month ago | A dependency injection container for go programming language. |
| [linker](https://github.com/logrange/linker) | 24 | 3 | 2018-12-04 | 1 month ago | A reflection based dependency injection and inversion of control library with components lifecycle support. |
| [gocontainer](https://github.com/vardius/gocontainer) | 12 | 0 | 2019-06-06 | 4 months ago | Simple Dependency Injection Container. |
| [di](https://github.com/goioc/di) | 2 | 1 | 2020-06-11 | 1 week ago | Spring-inspired Dependency Injection Container. |

### Project Layout
        
*Unofficial set of patterns for structuring projects.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [project-layout](https://github.com/golang-standards/project-layout) | 16794 | 448 | 2017-09-09 | 4 hours ago | Set of common historical and emerging project layout patterns in the Go ecosystem. |
| [go-restful-api](https://github.com/qiangxue/go-restful-api) | 1038 | 52 | 2016-08-15 | 6 months ago | An idiomatic Go RESTful API starter kit following SOLID principles and Clean Architecture with a common project layout. |
| [modern-go-application](https://github.com/sagikazarmark/modern-go-application) | 698 | 12 | 2018-09-14 | 6 days ago | Go application boilerplate and example applying modern practices. |
| [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) | 370 | 7 | 2016-12-18 | 3 months ago | A Go application boilerplate template for quick starting projects following production best practices. |
| [scaffold](https://github.com/catchplay/scaffold) | 78 | 3 | 2018-12-11 | 1 year ago | Scaffold generates starter Go project layout. Lets you focus on business logic implemeted. |
| [go-sample](https://github.com/zitryss/go-sample) | 59 | 1 | 2019-01-24 | 1 year ago | A sample layout for Go application projects with the real code. |

### Strings
        
*Libraries for working with strings.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [xstrings](https://github.com/huandu/xstrings) | 779 | 24 | 2015-01-06 | 1 month ago | Collection of useful string functions ported from other languages. |
| [strutil](https://github.com/ozgio/strutil) | 100 | 1 | 2018-08-16 | 9 months ago | String utilities. |
| [stringy](https://github.com/gobeam/stringy) | 32 | 2 | 2020-04-03 | 1 week ago | String manipulation library to convert string to camel case, snake case, kebab case / slugify etc. |
| [Stringy](https://github.com/gobeam/Stringy) | 30 | 1 | 2020-04-03 | 3 months ago | String manipulation library to convert string to camel case, snake case, kebab case / slugify etc. |

### Uncategorized
        
*These libraries were placed here because none of the other categories seemed to fit.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gopsutil](https://github.com/shirou/gopsutil) | 5114 | 194 | 2014-04-18 | 6 hours ago | Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc). |
| [archiver](https://github.com/mholt/archiver) | 2862 | 50 | 2016-04-08 | 1 week ago | Library and command for making and extracting .zip and .tar.gz archives. |
| [gosms](https://github.com/haxpax/gosms) | 1295 | 56 | 2015-01-23 | 3 weeks ago | Your own local SMS gateway in Go that can be used to send SMS. |
| [gofakeit](https://github.com/brianvoe/gofakeit) | 1061 | 8 | 2015-04-24 | 3 weeks ago | Random data generator written in go. |
| [go-resiliency](https://github.com/eapache/go-resiliency) | 1038 | 28 | 2014-11-29 | 8 months ago | Resiliency patterns for golang. |
| [base64Captcha](https://github.com/mojocn/base64Captcha) | 884 | 47 | 2017-12-12 | 1 month ago | Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha. |
| [go-commons-pool](https://github.com/jolestar/go-commons-pool) | 820 | 47 | 2015-12-28 | 2 weeks ago | Generic object pool for Golang. |
| [shortid](https://github.com/teris-io/shortid) | 586 | 9 | 2016-01-04 | 3 months ago | Distributed generation of super short, unique, non-sequential, URL friendly IDs. |
| [llvm](https://github.com/llir/llvm) | 553 | 29 | 2014-09-19 | 1 month ago | Library for interacting with LLVM IR in pure Go. |
| [health](https://github.com/dimiro1/health) | 394 | 6 | 2016-03-08 | 9 months ago | Easy to use, extensible health check library. |
| [go-conv](https://github.com/cstockton/go-conv) | 357 | 8 | 2016-10-11 | 3 years ago | Package conv provides fast and intuitive conversions across Go types. |
| [banner](https://github.com/dimiro1/banner) | 278 | 5 | 2016-03-25 | 9 months ago | Add beautiful banners into your Go applications. |
| [gountries](https://github.com/pariz/gountries) | 262 | 8 | 2016-01-13 | 1 month ago | Package that exposes country and subdivision data. |
| [ghorg](https://github.com/gabrie30/ghorg) | 218 | 7 | 2018-03-29 | 1 week ago | Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, and Bitbucket. |
| [antch](https://github.com/antchfx/antch) | 185 | 13 | 2017-09-28 | 1 month ago | A fast, powerful and extensible web crawling & scraping framework. |
| [stateless](https://github.com/qmuntal/stateless) | 183 | 4 | 2019-09-11 | 1 month ago | A fluent library for creating state machines. |
| [ffmt](https://github.com/go-ffmt/ffmt) | 172 | 5 | 2015-02-14 | 1 week ago | Beautify data display for Humans. |
| [lk](https://github.com/hyperboloide/lk) | 161 | 6 | 2016-07-14 | 2 months ago | A simple licensing library for golang. |
| [battery](https://github.com/distatus/battery) | 151 | 3 | 2016-03-12 | 5 days ago | Cross-platform, normalized battery information library. |
| [stats](https://github.com/go-playground/stats) | 136 | 3 | 2015-09-14 | 3 years ago | Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... |
| [bitio](https://github.com/icza/bitio) | 127 | 6 | 2016-05-31 | 7 months ago | Highly optimized bit-level Reader and Writer for Go. |
| [healthcheck](https://github.com/etherlabsio/healthcheck) | 127 | 5 | 2017-08-18 | 7 months ago | An opinionated and concurrent health-check HTTP handler for RESTful services. |
| [turtle](https://github.com/hackebrot/turtle) | 105 | 1 | 2017-09-08 | 1 month ago | Emojis for Go. |
| [go-unarr](https://github.com/gen2brain/go-unarr) | 102 | 6 | 2015-11-01 | 5 months ago | Decompression library for RAR, TAR, ZIP and 7z archives. |
| [gommit](https://github.com/antham/gommit) | 89 | 3 | 2016-08-30 | 1 month ago | Analyze git commit messages to ensure they follow defined patterns. |
| [shoutrrr](https://github.com/containrrr/shoutrrr) | 84 | 5 | 2019-04-11 | 1 week ago | Notification library providing easy access to various messaging services like slack, mattermost, gotify and smtp among others. |
| [gotoprom](https://github.com/cabify/gotoprom) | 81 | 90 | 2018-10-10 | 6 months ago | Type-safe metrics builder wrapper library for the official Prometheus client. |
| [captcha](https://github.com/steambap/captcha) | 63 | 5 | 2017-09-12 | 1 week ago | Package captcha provides an easy to use, unopinionated API for captcha generation. |
| [indigo](https://github.com/osamingo/indigo) | 61 | 1 | 2016-08-31 | 7 months ago | Distributed unique ID generator of using Sonyflake and encoded by Base58. |
| [morse](https://github.com/alwindoss/morse) | 60 | 2 | 2018-08-15 | 1 year ago | Library to convert to and from morse code. |
| [xkg](https://github.com/go-xkg/xkg) | 50 | 1 | 2015-01-05 | 5 years ago | X Keyboard Grabber. |
| [pdfgen](https://github.com/hyperboloide/pdfgen) | 45 | 3 | 2015-11-30 | 2 years ago | HTTP service to generate PDF from Json requests. |
| [persian](https://github.com/mavihq/persian) | 41 | 1 | 2017-10-16 | 2 years ago | Some utilities for Persian language in go. |
| [datacounter](https://github.com/miolini/datacounter) | 33 | 1 | 2015-10-14 | 5 months ago | Go counters for readers/writer/http.ResponseWriter. |
| [browscap_go](https://github.com/digitalcrab/browscap_go) | 31 | 6 | 2014-09-18 | 7 months ago | GoLang Library for [Browser Capabilities Project](http://browscap.org/). |
| [autoflags](https://github.com/artyom/autoflags) | 29 | 4 | 2014-05-15 | 1 year ago | Go package to automatically define command line flags from struct fields. |
| [sandid](https://github.com/aofei/sandid) | 24 | 1 | 2018-06-12 | 4 days ago | Every grain of sand on earth has its own ID. |
| [gosh](https://github.com/osamingo/gosh) | 22 | 0 | 2018-05-25 | 7 months ago | Provide Go Statistics Handler, Struct, Measure Method. |
| [url-shortener](https://github.com/pantrif/url-shortener) | 22 | 1 | 2018-06-04 | 2 years ago | A modern, powerful, and robust URL shortener microservice with mysql support. |
| [xdg](https://github.com/rkoesters/xdg) | 21 | 1 | 2013-12-15 | 1 year ago | FreeDesktop.org (xdg) Specs implemented in Go. |
| [metrics](https://github.com/pascaldekloe/metrics) | 13 | 1 | 2019-01-29 | 5 months ago | Library for metrics instrumentation and Prometheus exposition. |
| [anagent](https://github.com/mudler/anagent) | 11 | 2 | 2017-12-29 | 1 year ago | Minimalistic, pluggable Golang evloop/timer handler with dependency-injection. |
| [shellwords](https://github.com/Wing924/shellwords) | 11 | 1 | 2017-09-28 | 2 years ago | A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell. |
| [avgRating](https://github.com/kirillDanshin/avgRating) | 10 | 1 | 2017-08-05 | 3 years ago | Calculate average score and rating based on Wilson Score Equation. |
| [hostutils](https://github.com/Wing924/hostutils) | 9 | 1 | 2017-09-26 | 1 year ago | A golang library for packing and unpacking FQDNs list. |
| [numa](https://github.com/lrita/numa) | 2 | 1 | 2018-12-10 | 1 year ago | NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code. |

## Natural Language Processing
        
*Libraries for working with human languages.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [prose](https://github.com/jdkato/prose) | 2553 | 59 | 2017-02-17 | 2 weeks ago | Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. English only. |
| [gse](https://github.com/go-ego/gse) | 1378 | 47 | 2017-06-23 | 1 hour ago | Go efficient text segmentation; support english, chinese, japanese and other. |
| [gojieba](https://github.com/yanyiwu/gojieba) | 1169 | 69 | 2015-09-12 | 1 month ago | This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm. |
| [when](https://github.com/olebedev/when) | 1085 | 24 | 2016-12-27 | 8 months ago | Natural EN and RU language date/time parser with pluggable rules. |
| [go-pinyin](https://github.com/mozillazg/go-pinyin) | 780 | 34 | 2014-11-09 | 1 month ago | CN Hanzi to Hanyu Pinyin converter. |
| [kagome](https://github.com/ikawaha/kagome) | 510 | 24 | 2014-06-26 | 2 days ago | JP morphological analyzer written in pure Go. |
| [whatlanggo](https://github.com/abadojack/whatlanggo) | 430 | 16 | 2017-02-20 | 1 year ago | Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc). |
| [nlp](https://github.com/shixzie/nlp) | 364 | 23 | 2017-01-25 | 2 years ago | Extract values from strings and fill your structs with nlp. |
| [sentences](https://github.com/neurosnap/sentences) | 282 | 13 | 2015-08-07 | 1 year ago | Sentence tokenizer:  converts text into a list of sentences. |
| [nlp](https://github.com/james-bowman/nlp) | 269 | 23 | 2017-03-15 | 3 months ago | Go Natural Language Processing library supporting LSA (Latent Semantic Analysis). |
| [getlang](https://github.com/rylans/getlang) | 99 | 3 | 2018-03-01 | 2 months ago | Fast natural language detection package. |
| [go-nlp](https://github.com/nuance/go-nlp) | 86 | 7 | 2011-05-02 | 8 years ago | Utilities for working with discrete probability distributions and other tools useful for doing NLP work. |
| [go-unidecode](https://github.com/mozillazg/go-unidecode) | 73 | 2 | 2016-07-08 | 1 year ago | ASCII transliterations of Unicode text. |
| [gounidecode](https://github.com/fiam/gounidecode) | 69 | 3 | 2012-05-01 | 4 years ago | Unicode transliterator (also known as unidecode) for Go. |
| [RAKE.Go](https://github.com/afjoseph/RAKE.Go) | 64 | 6 | 2016-12-17 | 5 months ago | Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE). |
| [textcat](https://github.com/pebbe/textcat) | 61 | 6 | 2012-09-21 | 2 years ago | Go package for n-gram based text categorization, with support for utf-8 and raw text. |
| [go-stem](https://github.com/agonopol/go-stem) | 58 | 3 | 2011-09-23 | 2 years ago | Implementation of the porter stemming algorithm. |
| [MMSEGO](https://github.com/awsong/MMSEGO) | 58 | 5 | 2012-04-18 | 8 years ago | This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm. |
| [stemmer](https://github.com/dchest/stemmer) | 49 | 4 | 2011-03-21 | 3 years ago | Stemmer packages for Go programming language. Includes English and German stemmers. |
| [porter2](https://github.com/zentures/porter2) | 38 | 2 | 2015-01-21 | 4 years ago | Really fast Porter 2 stemmer. |
| [go2vec](https://github.com/danieldk/go2vec) | 37 | 7 | 2015-01-27 | 1 year ago | Reader and utility functions for word2vec embeddings. |
| [petrovich](https://github.com/striker2000/petrovich) | 29 | 1 | 2016-12-26 | 2 months ago | Petrovich is the library which inflects Russian names to given grammatical case. |
| [snowball](https://github.com/goodsign/snowball) | 27 | 1 | 2012-12-11 | 3 years ago | Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/). |
| [paicehusk](https://github.com/rookii/paicehusk) | 26 | 3 | 2012-09-29 | 6 years ago | Golang implementation of the Paice/Husk Stemming Algorithm. |
| [mystem](https://github.com/dveselov/mystem) | 24 | 3 | 2016-08-30 | 3 years ago | CGo bindings to Yandex.Mystem - russian morphology analyzer. |
| [icu](https://github.com/goodsign/icu) | 20 | 0 | 2012-12-11 | 3 years ago | Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1. |
| [golibstemmer](https://github.com/rjohnsondev/golibstemmer) | 17 | 1 | 2012-08-06 | 6 years ago | Go bindings for the snowball libstemmer library including porter 2. |
| [iuliia-go](https://github.com/mehanizm/iuliia-go) | 14 | 1 | 2020-04-27 | 2 months ago | Transliterate Cyrillic → Latin in every possible way. |
| [shamoji](https://github.com/osamingo/shamoji) | 12 | 2 | 2017-07-23 | 7 months ago | The shamoji is word filtering package written in Go. |
| [libtextcat](https://github.com/goodsign/libtextcat) | 10 | 1 | 2012-12-10 | 7 years ago | Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2. |
| [gotokenizer](https://github.com/xujiajun/gotokenizer) | 9 | 1 | 2018-10-11 | 1 year ago | A tokenizer based on the dictionary and Bigram language models for Go. (Now only support chinese segmentation) |
| [porter](https://github.com/a2800276/porter) | 8 | 1 | 2013-09-17 | 6 years ago | This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm. |
| [go-localize](https://github.com/m1/go-localize) | 5 | 1 | 2019-12-23 | 6 months ago | Simple and easy to use i18n (Internationalization and localization) engine - used for translating locale strings. |
| [transliterator](https://github.com/alexsergivan/transliterator) | 5 | 1 | 2020-04-17 | 2 months ago | Provides one-way string transliteration with supporting of language-specific transliteration rules. |
| [detectlanguage-go](https://github.com/detectlanguage/detectlanguage-go) | 2 | 0 | 2019-12-14 | 7 months ago | Language Detection API Go Client. Supports batch requests, short phrase or single word language detection. |

## Networking
        
*Libraries for working with various layers of the network.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [fasthttp](https://github.com/valyala/fasthttp) | 13040 | 390 | 2015-10-18 | 3 hours ago | Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http. |
| [kcptun](https://github.com/xtaci/kcptun) | 11959 | 598 | 2016-02-26 | 3 weeks ago | Extremely simple & fast udp tunnel based on KCP protocol. |
| [webrtc](https://github.com/pion/webrtc) | 5084 | 196 | 2018-05-18 | 1 day ago | A pure Go implementation of the WebRTC API. |
| [dns](https://github.com/miekg/dns) | 4724 | 170 | 2010-08-03 | 5 hours ago | Go library for working with DNS. |
| [quic-go](https://github.com/lucas-clemente/quic-go) | 4254 | 176 | 2016-04-06 | 11 hours ago | An implementation of the QUIC protocol in pure Go. |
| [httplab](https://github.com/gchaincl/httplab) | 3583 | 66 | 2017-02-08 | 1 year ago | HTTPLabs let you inspect HTTP requests and forge responses. |
| [gopacket](https://github.com/google/gopacket) | 3541 | 141 | 2015-03-16 | 6 days ago | Go library for packet processing with libpcap bindings. |
| [gnet](https://github.com/panjf2000/gnet) | 2687 | 101 | 2019-02-24 | 2 days ago | `gnet` is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go. |
| [kcp-go](https://github.com/xtaci/kcp-go) | 2658 | 145 | 2015-06-16 | 3 weeks ago | KCP - Fast and Reliable ARQ Protocol. |
| [gobgp](https://github.com/osrg/gobgp) | 1963 | 124 | 2014-09-14 | 1 day ago | BGP implemented in the Go Programming Language. |
| [ssh](https://github.com/gliderlabs/ssh) | 1655 | 49 | 2016-10-03 | 1 week ago | Higher-level API for building SSH servers (wraps crypto/ssh). |
| [fortio](https://github.com/fortio/fortio) | 1377 | 38 | 2017-10-10 | 5 days ago | Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC. |
| [water](https://github.com/songgao/water) | 1070 | 43 | 2013-03-25 | 2 months ago | Simple TUN/TAP library. |
| [go-getter](https://github.com/hashicorp/go-getter) | 942 | 191 | 2015-10-12 | 1 week ago | Go library for downloading files or directories from various sources using a URL. |
| [gev](https://github.com/Allenxuxu/gev) | 880 | 32 | 2019-09-01 | 2 months ago | gev is a lightweight, fast non-blocking TCP network library based on Reactor mode. |
| [sftp](https://github.com/pkg/sftp) | 880 | 53 | 2013-11-05 | 3 days ago | Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt. |
| [nff-go](https://github.com/intel-go/nff-go) | 837 | 75 | 2017-03-29 | 21 hours ago | Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF). |
| [grab](https://github.com/cavaliercoder/grab) | 711 | 16 | 2016-01-05 | 3 weeks ago | Go package for managing file downloads. |
| [ftp](https://github.com/jlaffaye/ftp) | 691 | 25 | 2011-05-06 | 22 hours ago | Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959). |
| [mdns](https://github.com/hashicorp/mdns) | 663 | 205 | 2014-01-29 | 5 months ago | Simple mDNS (Multicast DNS) client/server library in Golang. |
| [vssh](https://github.com/yahoo/vssh) | 580 | 19 | 2020-06-09 | 3 days ago | Go library for building network and server automation over SSH protocol. |
| [lhttp](https://github.com/fanux/lhttp) | 576 | 57 | 2015-12-29 | 2 years ago | Powerful websocket framework, build your IM server more easily. |
| [gosnmp](https://github.com/soniah/gosnmp) | 546 | 44 | 2012-08-27 | 6 days ago | Native Go library for performing SNMP actions. |
| [cidranger](https://github.com/yl2chen/cidranger) | 503 | 13 | 2017-08-21 | 3 months ago | Fast IP to CIDR lookup for Go. |
| [gotcp](https://github.com/gansidui/gotcp) | 461 | 43 | 2014-04-13 | 3 years ago | Go package for quickly writing tcp applications. |
| [peerdiscovery](https://github.com/schollz/peerdiscovery) | 424 | 18 | 2018-04-22 | 1 week ago | Pure Go library for cross-platform local peer discovery using UDP multicast. |
| [gopcap](https://github.com/akrennmair/gopcap) | 395 | 22 | 2009-11-19 | 3 months ago | Go wrapper for libpcap. |
| [stun](https://github.com/gortc/stun) | 395 | 16 | 2016-04-24 | 2 months ago | Go implementation of RFC 5389 STUN protocol. |
| [go-stun](https://github.com/ccding/go-stun) | 374 | 11 | 2013-08-17 | 2 weeks ago | Go implementation of the STUN client (RFC 3489 and RFC 5389). |
| [raw](https://github.com/mdlayher/raw) | 362 | 12 | 2015-07-06 | 2 months ago | Package raw enables reading and writing data at the device driver level for a network interface. |
| [tcp_server](https://github.com/firstrow/tcp_server) | 332 | 18 | 2014-10-13 | 1 year ago | Go library for building tcp servers faster. |
| [winrm](https://github.com/masterzen/winrm) | 262 | 16 | 2013-12-30 | 1 month ago | Go WinRM client to remotely execute commands on Windows machines. |
| [buffstreams](https://github.com/StabbyCutyou/buffstreams) | 237 | 14 | 2015-06-29 | 3 years ago | Streaming protocolbuffer data over TCP made easy. |
| [arp](https://github.com/mdlayher/arp) | 221 | 10 | 2015-07-06 | 7 months ago | Package arp implements the ARP protocol, as described in RFC 826. |
| [ethernet](https://github.com/mdlayher/ethernet) | 203 | 13 | 2015-07-03 | 1 year ago | Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. |
| [gaio](https://github.com/xtaci/gaio) | 198 | 10 | 2019-12-20 | 1 month ago | High performance async-io networking for Golang in proactor mode. |
| [jazigo](https://github.com/udhos/jazigo) | 153 | 12 | 2016-06-07 | 10 months ago | Jazigo is a tool written in Go for retrieving configuration for multiple network devices. |
| [utp](https://github.com/anacrolix/utp) | 153 | 16 | 2015-03-20 | 2 years ago | Go uTP micro transport protocol implementation. |
| [gmqtt](https://github.com/DrmagicE/gmqtt) | 152 | 14 | 2018-09-16 | 2 days ago | Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1. |
| [canopus](https://github.com/zubairhamed/canopus) | 142 | 14 | 2015-02-24 | 2 years ago | CoAP Client/Server implementation (RFC 7252). |
| [gnxi](https://github.com/google/gnxi) | 139 | 26 | 2017-09-26 | 6 hours ago | A collection of tools for Network Management that use the gNMI and gNOI protocols. |
| [sslb](https://github.com/eduardonunesp/sslb) | 126 | 8 | 2015-10-18 | 10 months ago | It's a Super Simples Load Balancer, just a little project to achieve some kind of performance. |
| [xtcp](https://github.com/xfxdev/xtcp) | 102 | 14 | 2016-03-31 | 4 months ago | TCP Server Framework with simultaneous full duplex communication,graceful shutdown,custom protocol. |
| [dhcp6](https://github.com/mdlayher/dhcp6) | 67 | 4 | 2015-05-22 | 1 year ago | Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. |
| [ether](https://github.com/songgao/ether) | 66 | 3 | 2014-05-21 | 4 years ago | Cross-platform Go package for sending and receiving ethernet frames. |
| [packet](https://github.com/aerogo/packet) | 51 | 3 | 2017-10-29 | 8 months ago | Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed. |
| [linkio](https://github.com/ian-kent/linkio) | 47 | 2 | 2014-12-24 | 3 years ago | Network link speed simulation for Reader/Writer interfaces. |
| [portproxy](https://github.com/aybabtme/portproxy) | 44 | 0 | 2014-12-13 | 5 years ago | Simple TCP proxy which adds CORS support to API's which don't support it. |
| [graval](https://github.com/koofr/graval) | 26 | 8 | 2014-04-22 | 2 years ago | Experimental FTP server framework. |
| [publicip](https://github.com/polera/publicip) | 21 | 2 | 2016-12-28 | 3 years ago | Package publicip returns your public facing IPv4 address (internet egress). |
| [golibwireshark](https://github.com/sunwxg/golibwireshark) | 18 | 2 | 2015-11-16 | 2 years ago | Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data. |
| [go-powerdns](https://github.com/joeig/go-powerdns) | 17 | 2 | 2018-06-21 | 4 weeks ago | PowerDNS API bindings for Golang. |
| [goshark](https://github.com/sunwxg/goshark) | 10 | 1 | 2015-11-01 | 2 years ago | Package goshark use tshark to decode IP packet and create data struct to analyse packet. |
| [llb](https://github.com/kirillDanshin/llb) | 10 | 1 | 2016-02-21 | 4 years ago | It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response. |
| [tspool](https://github.com/two/tspool) | 6 | 0 | 2018-10-27 | 1 year ago | A TCP Library use worker pool to improve performance and protect your server. |
| [gosocsvr](https://github.com/Rakeki/gosocsvr) | 5 | 2 | 2019-11-12 | 7 months ago | Socket server made simple. |
| [httpproxy](https://github.com/wzshiming/httpproxy) | 3 | 1 | 2018-07-18 | 2 months ago | HTTP proxy handler and dialer. |

### HTTP Clients
        
*Libraries for making HTTP requests.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [resty](https://github.com/go-resty/resty) | 3128 | 74 | 2015-08-28 | 5 days ago | Simple HTTP and REST client for Go inspired by Ruby rest-client. |
| [grequests](https://github.com/levigross/grequests) | 1615 | 34 | 2015-06-11 | 6 months ago | A Go "clone" of the great and famous Requests library. |
| [heimdall](https://github.com/gojek/heimdall) | 1561 | 51 | 2018-01-19 | 1 month ago | An enchanced http client with retry and hystrix capabilities. |
| [sling](https://github.com/dghubble/sling) | 1157 | 34 | 2015-04-02 | 2 weeks ago | Sling is a Go HTTP client library for creating and sending API requests. |
| [gentleman](https://github.com/h2non/gentleman) | 808 | 19 | 2016-02-21 | 3 months ago | Full-featured plugin-driven HTTP client library. |
| [pester](https://github.com/sethgrid/pester) | 504 | 6 | 2015-05-20 | 1 month ago | Go HTTP client calls with retries, backoff, and concurrency. |
| [request](https://github.com/monaco-io/request) | 62 | 4 | 2020-03-25 | 1 month ago | HTTP client for golang. If you have experience about axios or requests, you will love it. No 3rd dependency. |
| [sreq](https://github.com/winterssy/sreq) | 50 | 0 | 2019-12-04 | 6 months ago | A simple, user-friendly and concurrent safe HTTP request library for Go. |
| [rq](https://github.com/ddo/rq) | 35 | 2 | 2017-12-26 | 11 months ago | A nicer interface for golang stdlib HTTP client. |
| [httpretry](https://github.com/ybbus/httpretry) | 8 | 2 | 2020-02-05 | 5 months ago | Enriches the default go HTTP client with retry functionality. |
| [go-http-client](https://github.com/bozd4g/go-http-client) | 8 | 1 | 2019-12-14 | 1 month ago | Make http calls simply and easily. |

## OpenGL
        
*Libraries for using OpenGL in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [glfw](https://github.com/go-gl/glfw) | 960 | 42 | 2013-05-19 | 3 weeks ago | Go bindings for GLFW 3. |
| [gl](https://github.com/go-gl/gl) | 735 | 40 | 2015-02-22 | 1 year ago | Go bindings for OpenGL (generated via glow). |
| [mathgl](https://github.com/go-gl/mathgl) | 337 | 25 | 2013-02-13 | 9 months ago | Pure Go math package specialized for 3D math, with inspiration from GLM. |
| [gl](https://github.com/goxjs/gl) | 139 | 15 | 2015-05-18 | 4 months ago | Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android). |
| [glfw](https://github.com/goxjs/glfw) | 65 | 6 | 2014-12-27 | 4 months ago | Go cross-platform glfw library for creating an OpenGL context and receiving events. |
| [go-glmatrix](https://github.com/technohippy/go-glmatrix) | 1 | 1 | 2020-07-02 | 2 weeks ago | Go port of [glMatrix](http://glmatrix.net/) library. |

## ORM
        
*Libraries that implement Object-Relational Mapping or datamapping techniques.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [xorm](https://github.com/go-xorm/xorm) | 6133 | 263 | 2013-05-09 | 3 months ago | Simple and powerful ORM for Go. |
| [pg](https://github.com/go-pg/pg) | 3932 | 85 | 2013-04-24 | 3 hours ago | PostgreSQL ORM with focus on PostgreSQL specific features and performance. |
| [gorp](https://github.com/go-gorp/gorp) | 3398 | 115 | 2012-01-04 | 1 month ago | Go Relational Persistence, ORM-ish library for Go. |
| [sqlboiler](https://github.com/volatiletech/sqlboiler) | 3243 | 79 | 2016-02-21 | 9 hours ago | ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema. |
| [ent](https://github.com/facebookincubator/ent) | 2963 | 65 | 2019-06-12 | 7 hours ago | An entity framework for Go. Simple, yet powerful ORM for modeling and querying data. |
| [db](https://github.com/upper/db) | 2220 | 58 | 2013-10-23 | 5 days ago | Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers. |
| [pop](https://github.com/gobuffalo/pop) | 921 | 25 | 2018-02-07 | 3 weeks ago | Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite. |
| [reform](https://github.com/go-reform/reform) | 876 | 23 | 2016-02-25 | 8 hours ago | Better ORM for Go, based on non-empty interfaces and code generation. |
| [gormt](https://github.com/xxjwxc/gormt) | 655 | 13 | 2019-05-05 | 1 day ago | Mysql database to golang gorm struct. |
| [qbs](https://github.com/coocood/qbs) | 552 | 45 | 2013-02-02 | 3 years ago | Stands for Query By Struct. A Go ORM. |
| [go-queryset](https://github.com/jirfag/go-queryset) | 541 | 19 | 2017-09-03 | 2 weeks ago | 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM. |
| [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) | 404 | 15 | 2017-12-27 | 12 hours ago | A flexible and powerful SQL string builder library plus a zero-config ORM. |
| [zoom](https://github.com/albrow/zoom) | 264 | 17 | 2013-07-17 | 2 months ago | Blazing-fast datastore and querying engine built on Redis. |
| [rel](https://github.com/Fs02/rel) | 157 | 5 | 2019-10-06 | 5 days ago | Golang SQL Repository Layer for Clean (Onion) Architecture. |
| [grimoire](https://github.com/Fs02/grimoire) | 138 | 5 | 2018-03-05 | 1 week ago | Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3). |
| [gosql](https://github.com/rushteam/gosql) | 100 | 5 | 2020-04-27 | 1 week ago | A easy ORM for mysql. |
| [go-store](https://github.com/gosuri/go-store) | 99 | 9 | 2015-03-22 | 3 years ago | Simple and fast Redis backed key-value store library for Go. |
| [marlow](https://github.com/dadleyy/marlow) | 81 | 5 | 2017-09-15 | 6 months ago | Generated ORM from project structs for compile time safety assurances. |
| [gorm](https://github.com/jinzhu/gorm) | 79 | 4 | 2013-10-25 | 5 days ago | The fantastic ORM library for Golang, aims to be developer friendly. |
| [go-firestorm](https://github.com/jschoedt/go-firestorm) | 14 | 1 | 2018-12-04 | 2 weeks ago | A simple ORM for Google/Firebase Cloud Firestore. |
| [lore](https://github.com/abrahambotros/lore) | 6 | 1 | 2017-04-29 | 2 years ago | Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go. |

## Package Management
        
*Unofficial libraries for package and dependency management.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [dep](https://github.com/golang/dep) | 13242 | 284 | 2016-10-07 | 1 week ago | Go dependency tool. |
| [glide](https://github.com/Masterminds/glide) | 8027 | 193 | 2014-07-09 | 3 months ago | Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip. |
| [godep](https://github.com/tools/godep) | 5632 | 149 | 2013-05-01 | 2 years ago | dependency tool for go, godep helps build packages reproducibly by fixing their dependencies. |
| [govendor](https://github.com/kardianos/govendor) | 5013 | 101 | 2015-04-12 | 4 months ago | Go Package Manager. Go vendor tool that works with the standard vendor file. |
| [gopm](https://github.com/gpmgo/gopm) | 2496 | 85 | 2013-05-15 | 1 year ago | Go Package Manager. |
| [gom](https://github.com/mattn/gom) | 1383 | 36 | 2013-09-11 | 1 year ago | Go Manager - bundle for go. |
| [gpm](https://github.com/pote/gpm) | 1202 | 31 | 2013-09-05 | 2 years ago | Barebones dependency manager for Go. |
| [goop](https://github.com/petejkim/goop) | 779 | 37 | 2014-06-18 | 4 years ago | Simple dependency manager for Go (golang), inspired by Bundler. |
| [nut](https://github.com/jingweno/nut) | 242 | 7 | 2015-01-23 | 5 years ago | Vendor Go dependencies. |
| [johnny-deps](https://github.com/VividCortex/johnny-deps) | 215 | 20 | 2013-07-19 | 4 weeks ago | Minimal dependency version using Git. |
| [gigo](https://github.com/LyricalSecurity/gigo) | 199 | 5 | 2015-05-01 | 1 year ago | PIP-like dependency tool for golang, with support for private repositories and hashes. |
| [VenGO](https://github.com/DamnWidget/VenGO) | 118 | 8 | 2014-10-17 | 4 years ago | create and manage exportable isolated go virtual environments. |
| [mvn-golang](https://github.com/raydac/mvn-golang) | 106 | 10 | 2016-03-24 | 5 days ago | plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure. |
| [gop](https://github.com/lunny/gop) | 50 | 7 | 2017-02-18 | 1 year ago | Build and manage your Go applications out of GOPATH. |

## Performance
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [jaeger](https://github.com/jaegertracing/jaeger) | 11418 | 333 | 2016-04-15 | 5 days ago | A distributed tracing system. |
| [profile](https://github.com/pkg/profile) | 1275 | 39 | 2014-10-22 | 1 month ago | Simple profiling support package for Go. |
| [pixie](https://github.com/pixie-labs/pixie) | 84 | 23 | 2020-02-27 | 3 days ago | No instrumentation tracing for Golang applications via eBPF. |
| [tracer](https://github.com/kamilsk/tracer) | 33 | 2 | 2019-06-22 | 1 month ago | Simple, lightweight tracing. |

## Query Language
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [graphql](https://github.com/graphql-go/graphql) | 6639 | 146 | 2015-07-19 | 59 minutes ago | Implementation of GraphQL for Go. |
| [graphql-go](https://github.com/graph-gophers/graphql-go) | 3397 | 89 | 2016-10-18 | 6 days ago | GraphQL server with a focus on ease of use. |
| [gojsonq](https://github.com/thedevsaddam/gojsonq) | 1467 | 30 | 2018-05-19 | 4 months ago | A simple Go package to Query over JSON Data. |
| [jsonql](https://github.com/elgs/jsonql) | 231 | 8 | 2015-12-29 | 4 months ago | JSON query expression library in Golang. |
| [rql](https://github.com/a8m/rql) | 158 | 7 | 2018-06-05 | 2 weeks ago | Resource Query Language for REST API. |
| [graphql](https://github.com/tmc/graphql) | 52 | 10 | 2015-04-18 | 3 years ago | graphql parser + utilities. |
| [jsonslice](https://github.com/bhmj/jsonslice) | 48 | 0 | 2018-05-02 | 2 months ago | Jsonpath queries with advanced filters. |
| [straf](https://github.com/SonicRoshan/straf) | 19 | 1 | 2019-08-16 | 2 months ago | Easily Convert Golang structs to GraphQL objects. |
| [api-fu](https://github.com/ccbrown/api-fu) | 13 | 2 | 2019-07-30 | 1 month ago | Comprehensive GraphQL implementation. |
| [rest-query-parser](https://github.com/timsolov/rest-query-parser) | 7 | 1 | 2020-02-10 | 2 months ago | Query Parser for REST API. Filtering, validations, both `AND`, `OR` operations are supported directly in the query. |
| [gws](https://github.com/Zaba505/gws) | 1 | 1 | 2020-06-08 | 3 weeks ago | Apollos' "GraphQL over Websocket" client and server implementation. |

## Resource Embedding
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [packr](https://github.com/gobuffalo/packr) | 2909 | 44 | 2017-03-15 | 2 months ago | The simple and easy way to embed static files into Go binaries. |
| [statik](https://github.com/rakyll/statik) | 2832 | 46 | 2014-02-04 | 1 month ago | Embeds static files into a Go executable. |
| [go.rice](https://github.com/GeertJohan/go.rice) | 2002 | 53 | 2013-10-23 | 3 weeks ago | go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy. |
| [vfsgen](https://github.com/shurcooL/vfsgen) | 860 | 22 | 2015-05-18 | 1 month ago | Generates a vfsdata.go file that statically implements the given virtual filesystem. |
| [esc](https://github.com/mjibson/esc) | 567 | 15 | 2014-01-26 | 8 months ago | Embeds files into Go programs and provides http.FileSystem interfaces to them. |
| [fileb0x](https://github.com/UnnoTed/fileb0x) | 551 | 16 | 2016-01-23 | 9 months ago | Simple tool to embed files in go with focus on "customization" and ease to use. |
| [go-resources](https://github.com/omeid/go-resources) | 170 | 7 | 2015-02-21 | 1 month ago | Unfancy resources embedding with Go. |
| [statics](https://github.com/go-playground/statics) | 58 | 3 | 2015-10-07 | 3 years ago | Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks. |
| [templify](https://github.com/wlbr/templify) | 25 | 2 | 2016-05-22 | 11 months ago | Embed external template files into Go code to create single file binaries. |
| [go-embed](https://github.com/pyros2097/go-embed) | 24 | 2 | 2015-12-06 | 2 months ago | Generates go code to embed resource files into your library or executable. |
| [mule](https://github.com/wlbr/mule) | 8 | 1 | 2020-01-17 | 2 months ago | Embed external resources like images, movies ... into Go source code to create single file binaries using `go generate`. Focussed on simplicity. |

## Science and Data Analysis
        
*Libraries for scientific computing and data analyzing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gonum](https://github.com/gonum/gonum) | 4031 | 109 | 2017-03-25 | 3 days ago | Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more. |
| [stats](https://github.com/montanaflynn/stats) | 1773 | 50 | 2014-12-16 | 5 months ago | Statistics package with common functions missing from the Golang standard library. |
| [plot](https://github.com/gonum/plot) | 1644 | 60 | 2013-07-23 | 2 days ago | gonum/plot provides an API for building and drawing plots in Go. |
| [gosl](https://github.com/cpmech/gosl) | 1488 | 68 | 2015-02-09 | 2 months ago | Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more. |
| [streamtools](https://github.com/nytlabs/streamtools) | 1315 | 72 | 2013-07-05 | 5 years ago | general purpose, graphical tool for dealing with streams of data. |
| [go-dsp](https://github.com/mjibson/go-dsp) | 690 | 28 | 2011-11-02 | 2 years ago | Digital Signal Processing for Go. |
| [chart](https://github.com/vdobler/chart) | 651 | 44 | 2011-06-27 | 1 month ago | Simple Chart Plotting library for Go. Supports many graphs types. |
| [goraph](https://github.com/gyuho/goraph) | 631 | 39 | 2014-02-27 | 2 years ago | Pure Go graph theory library(data structure, algorith visualization). |
| [graph](https://github.com/yourbasic/graph) | 357 | 18 | 2017-04-27 | 9 months ago | Library of basic graph algorithms. |
| [orb](https://github.com/paulmach/orb) | 308 | 19 | 2016-03-28 | 2 months ago | 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support. |
| [ewma](https://github.com/VividCortex/ewma) | 302 | 23 | 2013-07-05 | 4 weeks ago | Exponentially-weighted moving averages. |
| [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) | 292 | 17 | 2018-10-01 | 4 days ago | Dataframes for machine-learning and statistics (similar to pandas). |
| [calendarheatmap](https://github.com/nikolaydubina/calendarheatmap) | 214 | 7 | 2020-07-01 | 2 weeks ago | Calendar heatmap in plain Go inspired by Github contribution activity. |
| [gohistogram](https://github.com/VividCortex/gohistogram) | 144 | 15 | 2013-07-02 | 4 weeks ago | Approximate histograms for data streams. |
| [TextRank](https://github.com/DavidBelicza/TextRank) | 107 | 6 | 2018-01-09 | 4 months ago | TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support. |
| [sparse](https://github.com/james-bowman/sparse) | 98 | 7 | 2017-05-16 | 1 month ago | Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries. |
| [pagerank](https://github.com/alixaxel/pagerank) | 59 | 6 | 2015-08-06 | 3 months ago | Weighted PageRank algorithm implemented in Go. |
| [geom](https://github.com/skelterjohn/geom) | 45 | 4 | 2011-06-07 | 2 years ago | 2D geometry for golang. |
| [evaler](https://github.com/soniah/evaler) | 42 | 4 | 2012-09-04 | 2 years ago | Simple floating point arithmetic expression evaluator. |
| [goent](https://github.com/kzahedi/goent) | 20 | 1 | 2017-08-08 | 1 year ago | GO Implementation of Entropy Measures. |
| [decimal](https://github.com/db47h/decimal) | 19 | 1 | 2020-05-27 | 3 weeks ago | Package decimal implements arbitrary-precision decimal floating-point arithmetic. |
| [triangolatte](https://github.com/tchayen/triangolatte) | 13 | 1 | 2018-07-18 | 1 week ago | 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs. |
| [GoStats](https://github.com/OGFris/GoStats) | 12 | 2 | 2018-07-22 | 1 year ago | GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions. |
| [ode](https://github.com/ChristopherRabotin/ode) | 11 | 3 | 2016-11-11 | 3 years ago | Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions. |
| [piecewiselinear](https://github.com/sgreben/piecewiselinear) | 11 | 2 | 2018-10-21 | 6 months ago | Tiny linear interpolation library. |
| [PiHex](https://github.com/claygod/PiHex) | 10 | 3 | 2016-07-22 | 3 months ago | Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi. |
| [assocentity](https://github.com/ndabAP/assocentity) | 6 | 1 | 2018-12-21 | 2 months ago | Package assocentity returns the average distance from words to a given entity. |
| [go-gt](https://github.com/ThePaw/go-gt) | 5 | 0 | 2015-09-14 | 4 years ago | Graph theory algorithms written in "Go" language. |
| [rootfinding](https://github.com/khezen/rootfinding) | 3 | 2 | 2018-10-30 | 4 months ago | root-finding algorithms library for finding roots of quadratic functions. |
| [bradleyterry](https://github.com/seanhagen/bradleyterry) | 2 | 1 | 2019-04-30 | 1 year ago | Provides a Bradley-Terry Model for pairwise comparisons. |

## Security
        
*Libraries that are used to help make your application more secure.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [lego](https://github.com/go-acme/lego) | 4126 | 103 | 2015-06-08 | 1 week ago | Pure Go ACME client library and CLI tool (for use with Let's Encrypt). |
| [cameradar](https://github.com/Ullaakut/cameradar) | 2282 | 106 | 2016-05-20 | 2 months ago | Tool and library to remotely hack RTSP streams from surveillance cameras. |
| [memguard](https://github.com/awnumar/memguard) | 1786 | 48 | 2017-04-22 | 2 months ago | A pure Go library for handling sensitive values in memory. |
| [acmetool](https://github.com/hlandau/acmetool) | 1781 | 67 | 2015-11-15 | 3 weeks ago | ACME (Let's Encrypt) client tool with automatic renewal. |
| [secure](https://github.com/unrolled/secure) | 1527 | 36 | 2014-05-20 | 1 day ago | HTTP middleware for Go that facilitates some quick security wins. |
| [acra](https://github.com/cossacklabs/acra) | 613 | 36 | 2016-11-14 | 1 day ago | Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system. |
| [nacl](https://github.com/kevinburke/nacl) | 483 | 12 | 2017-07-20 | 11 months ago | Go implementation of the NaCL set of API's. |
| [badactor](https://github.com/jaredfolkins/badactor) | 282 | 9 | 2014-12-12 | 2 months ago | In-memory, application-driven jailer built in the spirit of fail2ban. |
| [ssh-vault](https://github.com/ssh-vault/ssh-vault) | 251 | 10 | 2016-09-29 | 5 months ago | encrypt/decrypt using ssh keys. |
| [passlib](https://github.com/hlandau/passlib) | 235 | 10 | 2014-12-21 | 1 year ago | Futureproof password hashing library. |
| [optimus-go](https://github.com/pjebs/optimus-go) | 221 | 5 | 2015-05-05 | 2 months ago | ID hashing and Obfuscation using Knuth's Algorithm. |
| [go-yara](https://github.com/hillu/go-yara) | 173 | 18 | 2015-01-25 | 1 month ago | Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)". |
| [simple-scrypt](https://github.com/elithrar/simple-scrypt) | 163 | 7 | 2015-04-14 | 1 year ago | Scrypt package with a simple, obvious API and automatic cost calibration built-in. |
| [argon2pw](https://github.com/raja/argon2pw) | 82 | 4 | 2018-03-13 | 1 year ago | Argon2 password hash generation with constant-time password comparison. |
| [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) | 37 | 1 | 2017-10-19 | 1 year ago | A probably paranoid package for securely hashing and encrypting passwords. |
| [certificates](https://github.com/mvmaasakkers/certificates) | 18 | 1 | 2019-03-04 | 1 month ago | An opinionated tool for generating tls certificates. |
| [goArgonPass](https://github.com/dwin/goArgonPass) | 12 | 2 | 2018-05-30 | 5 months ago | Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations. |
| [go-generate-password](https://github.com/m1/go-generate-password) | 11 | 1 | 2019-11-14 | 8 months ago | Password generator that can be used on the cli or as a library. |
| [sslmgr](https://github.com/adrianosela/sslmgr) | 9 | 2 | 2019-04-02 | 1 year ago | SSL certificates made easy with a high level wrapper around acme/autocert. |
| [secureio](https://github.com/xaionaro-go/secureio) | 8 | 1 | 2018-12-25 | 1 month ago | An keyexchanging+authenticating+encrypting wrapper and multiplexer for `io.ReadWriteCloser` based on XChaCha20-poly1305, ECDH and ED25519. |
| [argon2-hashing](https://github.com/andskur/argon2-hashing) | 8 | 1 | 2019-01-02 | 3 months ago | light wrapper around Go's argon2 package that closely mirrors with Go's standard library Bcrypt and simple-scrypt package. |
| [firewalld-rest](https://github.com/prashantgupta24/firewalld-rest) | 6 | 1 | 2020-06-12 | 1 week ago | A rest application to dynamically update firewalld rules on a linux server. |

## Serialization
        
*Libraries and tools for binary serialization.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go](https://github.com/json-iterator/go) | 7984 | 225 | 2016-11-30 | 4 days ago | High-performance 100% compatible drop-in replacement of "encoding/json". |
| [protobuf](https://github.com/golang/protobuf) | 6890 | 219 | 2014-11-23 | 2 months ago | Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers. |
| [protobuf](https://github.com/gogo/protobuf) | 3960 | 103 | 2014-12-03 | 3 days ago | Protocol Buffers for Go with Gadgets. |
| [mapstructure](https://github.com/mitchellh/mapstructure) | 3526 | 62 | 2013-05-20 | 6 days ago | Go library for decoding generic map values into native Go structures. |
| [go](https://github.com/ugorji/go) | 1436 | 55 | 2013-05-30 | 1 month ago | High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support. |
| [colfer](https://github.com/pascaldekloe/colfer) | 548 | 33 | 2015-09-05 | 2 weeks ago | Code generation for the Colfer binary format. |
| [csvutil](https://github.com/jszwec/csvutil) | 396 | 9 | 2017-10-30 | 3 days ago | High Performance, idiomatic CSV record encoding and decoding to native Go structures. |
| [go-capnproto](https://github.com/glycerine/go-capnproto) | 277 | 11 | 2013-11-07 | 6 months ago | Cap'n Proto library and parser for go. |
| [cbor](https://github.com/fxamacker/cbor) | 191 | 6 | 2019-05-15 | 2 months ago | Small, safe, and easy CBOR encoding and decoding library. |
| [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) | 140 | 9 | 2012-12-23 | 1 year ago | GoLang library for working with PHP session format and PHP Serialize/Unserialize functions. |
| [structomap](https://github.com/danhper/structomap) | 113 | 7 | 2015-05-13 | 1 year ago | Library to easily and dynamically generate maps from static structures. |
| [bambam](https://github.com/glycerine/bambam) | 60 | 4 | 2014-09-17 | 3 years ago | generator for Cap'n Proto schemas from go. |
| [asn1](https://github.com/Logicalis/asn1) | 46 | 8 | 2016-02-29 | 1 year ago | Asn.1 BER and DER encoding library for golang. |
| [binstruct](https://github.com/ghostiam/binstruct) | 19 | 1 | 2018-10-23 | 11 months ago | Golang binary decoder for mapping data into the structure. |
| [fwencoder](https://github.com/o1egl/fwencoder) | 13 | 1 | 2017-12-25 | 5 months ago | Fixed width file parser (encoding and decoding library) for Go. |
| [pletter](https://github.com/vimeda/pletter) | 11 | 0 | 2019-07-09 | 1 month ago | A standard way to wrap a proto message for message brokers. |
| [bel](https://github.com/csweichel/bel) | 10 | 2 | 2019-02-20 | 1 year ago | Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC. |
| [elastic](https://github.com/epiclabs-io/elastic) | 10 | 0 | 2020-02-25 | 5 months ago | Convert slices, maps or any other unknown value across different types at run-time, no matter what. |
| [bel](https://github.com/32leaves/bel) | 8 | 1 | 2019-02-20 | 1 year ago | Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC. |
| [fixedwidth](https://github.com/huydang284/fixedwidth) | 4 | 1 | 2019-08-11 | 7 months ago | Fixed-width text formatting (UTF-8 supported). |
| [go-lctree](https://github.com/sbourlon/go-lctree) | 0 | 1 | 2020-05-04 | 1 month ago | Provides a CLI and primitives to serialize and deserialize [LeetCode binary trees](https://support.leetcode.com/hc/en-us/articles/360011883654-What-does-1-null-2-3-mean-in-binary-tree-representation). |

## Server Applications
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [etcd](https://github.com/etcd-io/etcd) | 32127 | 1345 | 2013-07-06 | 42 minutes ago | Highly-available key value store for shared configuration and service discovery. |
| [caddy](https://github.com/caddyserver/caddy) | 29442 | 747 | 2015-01-13 | 1 day ago | Caddy is an alternative, HTTP/2 web server that's easy to configure and use. |
| [minio](https://github.com/minio/minio) | 23037 | 519 | 2015-01-14 | 49 minutes ago | Minio is a distributed object storage server. |
| [roadrunner](https://github.com/spiral/roadrunner) | 4453 | 156 | 2017-12-26 | 1 week ago | High-performance PHP application server, load-balancer and process manager. |
| [devd](https://github.com/cortesi/devd) | 3011 | 70 | 2015-09-27 | 3 months ago | Local webserver for developers. |
| [sftpgo](https://github.com/drakkan/sftpgo) | 1856 | 42 | 2019-07-20 | 3 days ago | Full featured and highly configurable SFTP server software. |
| [algernon](https://github.com/xyproto/algernon) | 1710 | 50 | 2015-03-10 | 2 days ago | HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber. |
| [flipt](https://github.com/markphelps/flipt) | 1270 | 16 | 2016-11-05 | 1 day ago | A self contained feature flag solution written in Go and Vue. |
| [fider](https://github.com/getfider/fider) | 1245 | 33 | 2017-01-17 | 1 week ago | Fider is an open platform to collect and organize customer feedback. |
| [flagr](https://github.com/checkr/flagr) | 1195 | 74 | 2017-10-03 | 5 days ago | Flagr is an open-source feature flagging and A/B testing service. |
| [trickster](https://github.com/tricksterproxy/trickster) | 1158 | 34 | 2018-03-29 | 2 weeks ago | HTTP reverse proxy cache and time series accelerator. |
| [trickster](https://github.com/Comcast/trickster) | 1053 | 35 | 2018-03-29 | 4 months ago | HTTP reverse proxy cache and time series accelerator. |
| [discovery](https://github.com/bilibili/discovery) | 981 | 51 | 2018-04-20 | 3 months ago | A registry for resilient mid-tier load balancing and failover. |
| [jackal](https://github.com/ortuman/jackal) | 844 | 38 | 2017-11-13 | 1 month ago | An XMPP server written in Go. |
| [dudeldu](https://github.com/krotik/dudeldu) | 115 | 3 | 2016-09-07 | 10 months ago | A simple SHOUTcast server. |
| [lets-proxy2](https://github.com/rekby/lets-proxy2) | 38 | 2 | 2019-04-12 | 1 week ago | Reverse proxy for handle https with issue certificates in fly from lets-encrypt. |
| [psql-streamer](https://github.com/blind-oracle/psql-streamer) | 19 | 4 | 2019-04-28 | 4 months ago | Stream database events from PostgreSQL to Kafka. |
| [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) | 14 | 1 | 2018-10-23 | 1 year ago | Nginx log parser and exporter to Prometheus. |
| [riemann-relay](https://github.com/blind-oracle/riemann-relay) | 0 | 1 | 2019-04-23 | 9 months ago | Relay to load-balance Riemann events and/or convert them to Carbon. |

## Stream Processing
        
*Libraries and tools for stream processing and reactive programming.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-streams](https://github.com/reugn/go-streams) | 375 | 12 | 2019-04-30 | 2 weeks ago | Go stream processing library. |

## Template Engines
        
*Libraries and tools for templating and lexing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gofpdf](https://github.com/jung-kurt/gofpdf) | 3581 | 95 | 2015-03-13 | 8 months ago | PDF document generator with high level support for text, drawing and images. |
| [quicktemplate](https://github.com/valyala/quicktemplate) | 1789 | 59 | 2016-03-06 | 2 weeks ago | Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it. |
| [pongo2](https://github.com/flosch/pongo2) | 1713 | 56 | 2013-08-23 | 1 month ago | Django-like template-engine for Go. |
| [hero](https://github.com/shiyanhui/hero) | 1374 | 44 | 2017-01-15 | 6 months ago | Hero is a handy, fast and powerful go template engine. |
| [mustache](https://github.com/hoisie/mustache) | 989 | 35 | 2009-12-30 | 2 weeks ago | Go implementation of the Mustache template language. |
| [amber](https://github.com/eknkc/amber) | 849 | 20 | 2012-10-31 | 1 year ago | Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade. |
| [ace](https://github.com/yosssi/ace) | 788 | 23 | 2014-07-13 | 2 years ago | Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold. |
| [gorazor](https://github.com/sipin/gorazor) | 748 | 55 | 2014-05-01 | 2 weeks ago | Razor view engine for Golang. |
| [jet](https://github.com/CloudyKit/jet) | 676 | 22 | 2016-03-31 | 5 days ago | Jet template engine. |
| [ego](https://github.com/benbjohnson/ego) | 449 | 16 | 2014-02-23 | 6 months ago | Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled. |
| [fasttemplate](https://github.com/valyala/fasttemplate) | 404 | 17 | 2015-08-19 | 3 weeks ago | Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/). |
| [raymond](https://github.com/aymerick/raymond) | 382 | 11 | 2015-04-22 | 1 month ago | Complete handlebars implementation in Go. |
| [maroto](https://github.com/johnfercher/maroto) | 194 | 7 | 2019-05-20 | 1 month ago | A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple. |
| [soy](https://github.com/robfig/soy) | 150 | 13 | 2013-12-15 | 1 week ago | Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/). |
| [goview](https://github.com/foolin/goview) | 143 | 4 | 2019-04-14 | 3 months ago | Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application. |
| [liquid](https://github.com/osteele/liquid) | 106 | 5 | 2017-06-26 | 7 months ago | Go implementation of Shopify Liquid templates. |
| [kasia.go](https://github.com/ziutek/kasia.go) | 71 | 2 | 2010-12-07 | 4 years ago | Templating system for HTML and other text documents - go implementation. |
| [velvet](https://github.com/gobuffalo/velvet) | 71 | 5 | 2016-12-29 | 3 years ago | Complete handlebars implementation in Go. |
| [damsel](https://github.com/dskinner/damsel) | 24 | 4 | 2012-05-02 | 4 years ago | Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others. |
| [extemplate](https://github.com/dannyvankooten/extemplate) | 23 | 3 | 2018-08-10 | 2 months ago | Tiny wrapper around html/template to allow for easy file-based template inheritance. |
| [gospin](https://github.com/m1/gospin) | 15 | 1 | 2019-02-22 | 2 months ago | Article spinning and spintax/spinning syntax engine, useful for A/B, testing pieces of text/articles and creating more natural conversations. |

## Testing
        
*Libraries for testing codebases and generating test data.*

### Testing Frameworks
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [testify](https://github.com/stretchr/testify) | 11033 | 166 | 2012-10-16 | 3 hours ago | Sacred extension to the standard go testing package. |
| [go-cmp](https://github.com/google/go-cmp) | 1841 | 27 | 2017-07-07 | 6 days ago | Package for comparing Go values in tests. |
| [httpexpect](https://github.com/gavv/httpexpect) | 1436 | 37 | 2016-04-29 | 1 day ago | Concise, declarative, and easy to use end-to-end HTTP and REST API testing. |
| [godog](https://github.com/cucumber/godog) | 1078 | 102 | 2015-06-10 | 3 days ago | Cucumber or Behat like BDD framework for Go. |
| [godog](https://github.com/DATA-DOG/godog) | 895 | 30 | 2015-06-10 | 5 months ago | Cucumber or Behat like BDD framework for Go. |
| [goblin](https://github.com/franela/goblin) | 695 | 14 | 2013-09-19 | 5 days ago | Mocha like testing framework fo Go. |
| [baloo](https://github.com/h2non/baloo) | 688 | 10 | 2016-05-29 | 1 year ago | Expressive and versatile end-to-end HTTP API testing made easy. |
| [testfixtures](https://github.com/go-testfixtures/testfixtures) | 534 | 5 | 2016-04-05 | 1 day ago | A helper for Rails' like test fixtures to test database applications. |
| [go-vcr](https://github.com/dnaeon/go-vcr) | 437 | 7 | 2015-12-14 | 2 months ago | Record and replay your HTTP interactions for fast, deterministic and accurate tests. |
| [go-mutesting](https://github.com/zimmski/go-mutesting) | 350 | 7 | 2014-12-26 | 9 months ago | Mutation testing for Go source code. |
| [gofight](https://github.com/appleboy/gofight) | 328 | 11 | 2016-03-29 | 6 months ago | API Handler Testing for Golang Router framework. |
| [frisby](https://github.com/hofstadter-io/frisby) | 259 | 8 | 2015-09-15 | 4 months ago | REST API testing framework. |
| [frisby](https://github.com/verdverm/frisby) | 257 | 8 | 2015-09-15 | 4 months ago | REST API testing framework. |
| [go-carpet](https://github.com/msoap/go-carpet) | 210 | 4 | 2016-02-28 | 4 months ago | Tool for viewing test coverage in terminal. |
| [charlatan](https://github.com/percolate/charlatan) | 190 | 45 | 2017-10-06 | 10 months ago | Tool to generate fake interface implementations for tests. |
| [gotest.tools](https://github.com/gotestyourself/gotest.tools) | 165 | 6 | 2017-08-08 | 1 week ago | A collection of packages to augment the go testing package and support common patterns. |
| [commander](https://github.com/SimonBaeumer/commander) | 154 | 9 | 2019-02-22 | 3 weeks ago | Tool for testing cli applications on windows, linux and osx. |
| [commander](https://github.com/commander-cli/commander) | 154 | 8 | 2019-02-22 | 13 hours ago | Tool for testing cli applications on windows, linux and osx. |
| [endly](https://github.com/viant/endly) | 152 | 15 | 2017-08-28 | 1 week ago | Declarative end to end functional testing. |
| [gospec](https://github.com/luontola/gospec) | 115 | 4 | 2009-11-24 | 6 years ago | BDD-style testing framework for the Go programming language. |
| [cupaloy](https://github.com/bradleyjkemp/cupaloy) | 113 | 2 | 2017-08-07 | 1 month ago | Simple snapshot testing addon for your test framework. |
| [dbcleaner](https://github.com/khaiql/dbcleaner) | 113 | 2 | 2017-01-17 | 4 months ago | Clean database for testing purpose, inspired by `database_cleaner` in Ruby. |
| [go-testdeep](https://github.com/maxatome/go-testdeep) | 106 | 1 | 2018-05-26 | 3 days ago | Extremely flexible golang deep comparison, extends the go testing package. |
| [wstest](https://github.com/posener/wstest) | 78 | 2 | 2017-03-31 | 4 months ago | Websocket client for unit-testing a websocket http.Handler. |
| [embedded-postgres](https://github.com/fergusstrange/embedded-postgres) | 65 | 3 | 2019-11-16 | 1 month ago | Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test. |
| [gospecify](https://github.com/stesla/gospecify) | 54 | 6 | 2009-11-20 | 8 years ago | This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec. |
| [restit](https://github.com/go-restit/restit) | 51 | 6 | 2014-06-25 | 9 months ago | Go micro framework to help writing RESTful API integration test. |
| [testcase](https://github.com/adamluzsi/testcase) | 47 | 4 | 2019-04-22 | 1 week ago | Idiomatic testing framework for Behavior Driven Development. |
| [jsonassert](https://github.com/kinbiko/jsonassert) | 43 | 0 | 2018-10-26 | 1 month ago | Package for verifying that your JSON payloads are serialized correctly. |
| [gomatch](https://github.com/jfilipczyk/gomatch) | 37 | 2 | 2019-01-27 | 1 year ago | library created for testing JSON against patterns. |
| [dsunit](https://github.com/viant/dsunit) | 33 | 9 | 2016-06-13 | 5 months ago | Datastore testing for SQL, NoSQL, structured files. |
| [hamcrest](https://github.com/rdrdr/hamcrest) | 27 | 3 | 2010-12-22 | 9 years ago | fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results. |
| [assert](https://github.com/go-playground/assert) | 20 | 1 | 2015-07-20 | 3 months ago | Basic Assertion Library used along side native go testing, with building blocks for custom assertions. |
| [flute](https://github.com/suzuki-shunsuke/flute) | 14 | 0 | 2019-07-06 | 1 week ago | HTTP client testing framework. |
| [schema](https://github.com/jgroeneveld/schema) | 11 | 3 | 2015-08-13 | 9 months ago | Quick and easy expression matching for JSON schemas used in requests and responses. |
| [badio](https://github.com/cavaliercoder/badio) | 9 | 1 | 2016-02-11 | 4 years ago | Extensions to Go's `testing/iotest` package. |
| [biff](https://github.com/fulldump/biff) | 9 | 1 | 2018-03-28 | 4 months ago | Bifurcation testing framework, BDD compatible. |
| [gocrest](https://github.com/corbym/gocrest) | 9 | 1 | 2017-12-23 | 2 years ago | Composable hamcrest-like matchers for Go assertions. |
| [gosuite](https://github.com/pavlo/gosuite) | 9 | 1 | 2016-10-15 | 3 years ago | Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests. |
| [testsql](https://github.com/zhulongcheng/testsql) | 9 | 2 | 2018-09-22 | 10 months ago | Generate test data from SQL files before testing and clear it after finished. |
| [tt](https://github.com/vcaesar/tt) | 9 | 1 | 2018-04-03 | 1 month ago | Simple and colorful test tools. |
| [gogiven](https://github.com/corbym/gogiven) | 8 | 4 | 2017-12-31 | 2 years ago | YATSPEC-like BDD testing framework for Go. |
| [trial](https://github.com/jgroeneveld/trial) | 4 | 1 | 2015-06-18 | 9 months ago | Quick and easy extendable assertions without introducing much boilerplate. |

### Mock
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [chromedp](https://github.com/chromedp/chromedp) | 5086 | 143 | 2017-01-24 | 18 hours ago | a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol. |
| [mock](https://github.com/golang/mock) | 4448 | 73 | 2015-06-12 | 2 days ago | Mocking framework for the Go programming language. |
| [go-fuzz](https://github.com/dvyukov/go-fuzz) | 3526 | 88 | 2015-04-15 | 4 months ago | Randomized testing system. |
| [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) | 2668 | 28 | 2014-02-07 | 1 month ago | Mock SQL driver for testing database interactions. |
| [selenoid](https://github.com/aerokube/selenoid) | 1662 | 92 | 2016-08-22 | 5 days ago | alternative Selenium hub server that launches browsers within containers. |
| [hoverfly](https://github.com/SpectoLabs/hoverfly) | 1610 | 61 | 2015-11-30 | 1 week ago | HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI. |
| [gock](https://github.com/h2non/gock) | 1036 | 17 | 2016-03-02 | 2 months ago | Versatile HTTP mocking made easy. |
| [gofuzz](https://github.com/google/gofuzz) | 818 | 24 | 2014-07-31 | 2 weeks ago | Library for populating go objects with random values. |
| [httpmock](https://github.com/jarcoal/httpmock) | 797 | 8 | 2014-02-24 | 4 months ago | Easy mocking of HTTP responses from external resources. |
| [rod](https://github.com/go-rod/rod) | 595 | 14 | 2020-01-21 | 23 minutes ago | A Devtools driver to make web automation and scraping easy. |
| [cdp](https://github.com/mafredri/cdp) | 454 | 18 | 2017-03-12 | 1 month ago | Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it. |
| [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) | 429 | 6 | 2014-05-21 | 1 month ago | Tool for generating self-contained mock objects. |
| [minimock](https://github.com/gojuno/minimock) | 319 | 10 | 2016-08-03 | 1 month ago | Mock generator for Go interfaces. |
| [go-txdb](https://github.com/DATA-DOG/go-txdb) | 283 | 6 | 2015-07-08 | 6 months ago | Single transaction based database driver mainly for testing purposes. |
| [rod](https://github.com/ysmood/rod) | 273 | 8 | 2020-01-21 | 1 month ago | A chrome devtools controller that is easy and safe to use. |
| [ggr](https://github.com/aerokube/ggr) | 258 | 26 | 2016-06-16 | 2 weeks ago | a lightweight server that routes and proxies Selenium WebDriver requests to multiple Selenium hubs. |
| [tavor](https://github.com/zimmski/tavor) | 221 | 12 | 2014-05-18 | 1 year ago | Generic fuzzing and delta-debugging framework. |
| [govcr](https://github.com/seborama/govcr) | 90 | 2 | 2016-07-10 | 10 months ago | HTTP mock for Golang: record and replay HTTP interactions for offline testing. |
| [timex](https://github.com/cabify/timex) | 36 | 74 | 2020-01-02 | 4 weeks ago | A test-friendly replacement for the native `time` package. |
| [mockhttp](https://github.com/tv42/mockhttp) | 22 | 1 | 2011-06-11 | 5 years ago | Mock object for Go http.ResponseWriter. |
| [go-localstack](https://github.com/elgohr/go-localstack) | 3 | 1 | 2020-03-18 | 20 hours ago | Tool for using localstack in AWS testing. |

### Fail injection
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [failpoint](https://github.com/pingcap/failpoint) | 520 | 79 | 2019-04-02 | 3 weeks ago | An implementation of [failpoints](http://www.freebsd.org/cgi/man.cgi?query=fail) for Golang. |

## Text Processing
        
*Libraries for parsing and manipulating texts.*

### Specific Formats
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [colly](https://github.com/gocolly/colly) | 11472 | 294 | 2017-09-29 | 3 days ago | Fast and Elegant Scraping Framework for Gophers. |
| [goquery](https://github.com/PuerkitoBio/goquery) | 9070 | 268 | 2012-08-29 | 3 months ago | GoQuery brings a syntax and a set of features similar to jQuery to the Go language. |
| [blackfriday](https://github.com/russross/blackfriday) | 4437 | 92 | 2011-05-27 | 2 weeks ago | Markdown processor in Go. |
| [toml](https://github.com/BurntSushi/toml) | 3280 | 85 | 2013-02-26 | 2 months ago | TOML configuration format (encoder/decoder with reflection). |
| [sh](https://github.com/mvdan/sh) | 2970 | 47 | 2016-01-16 | 1 week ago | Shell parser and formatter. |
| [go-humanize](https://github.com/dustin/go-humanize) | 2284 | 33 | 2012-01-13 | 3 months ago | Formatters for time, numbers, and memory size to human readable format. |
| [bluemonday](https://github.com/microcosm-cc/bluemonday) | 1608 | 31 | 2013-11-20 | 1 month ago | HTML Sanitizer. |
| [gofeed](https://github.com/mmcdole/gofeed) | 1318 | 39 | 2016-01-23 | 2 days ago | Parse RSS and Atom feeds in Go. |
| [inject](https://github.com/facebookarchive/inject) | 1248 | 44 | 2013-10-21 | 1 year ago | Package inject provides a reflect based injector. |
| [go-toml](https://github.com/pelletier/go-toml) | 795 | 33 | 2013-02-24 | 2 days ago | Go library for the TOML format with query support and handy cli tools. |
| [commonregex](https://github.com/mingrammer/commonregex) | 644 | 20 | 2017-03-23 | 8 months ago | A collection of common regular expressions for Go. |
| [slug](https://github.com/gosimple/slug) | 562 | 11 | 2014-03-31 | 6 months ago | URL-friendly slugify with multiple languages support. |
| [dataflowkit](https://github.com/slotix/dataflowkit) | 404 | 18 | 2017-02-09 | 1 month ago | Web scraping Framework to turn websites into structured data. |
| [mxj](https://github.com/clbanning/mxj) | 386 | 25 | 2014-02-03 | 1 week ago | Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages. |
| [gographviz](https://github.com/awalterschulze/gographviz) | 376 | 14 | 2015-03-14 | 1 month ago | Parses the Graphviz DOT language. |
| [go-runewidth](https://github.com/mattn/go-runewidth) | 276 | 11 | 2013-06-21 | 2 weeks ago | Functions to get fixed width of the character or string. |
| [gotext](https://github.com/leonelquinteros/gotext) | 272 | 6 | 2016-06-19 | 5 days ago | GNU gettext utilities for Go. |
| [htmlquery](https://github.com/antchfx/htmlquery) | 266 | 9 | 2017-12-05 | 3 months ago | An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression. |
| [goq](https://github.com/andrewstuart/goq) | 171 | 7 | 2017-02-20 | 1 year ago | Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery). |
| [goribot](https://github.com/zhshch2002/goribot) | 163 | 8 | 2019-09-08 | 1 week ago | A simple golang spider/scraping framework,build a spider in 3 lines. |
| [go-nmea](https://github.com/adrianmo/go-nmea) | 128 | 7 | 2015-07-22 | 1 week ago | NMEA parser library for the Go language. |
| [sdp](https://github.com/gortc/sdp) | 100 | 7 | 2016-05-13 | 2 months ago | SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)]. |
| [go-zero-width](https://github.com/trubitsyn/go-zero-width) | 79 | 1 | 2018-06-18 | 6 months ago | Zero-width character detection and removal for Go. |
| [podcast](https://github.com/eduncan911/podcast) | 74 | 3 | 2017-02-02 | 1 month ago | iTunes Compliant and RSS 2. |
| [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) | 71 | 8 | 2016-07-05 | 1 month ago | Editorconfig file parser and manipulator for Go. |
| [go-slugify](https://github.com/mozillazg/go-slugify) | 64 | 2 | 2016-07-16 | 2 months ago | Make pretty slug with multiple languages support. |
| [align](https://github.com/Guitarbum722/align) | 63 | 4 | 2017-04-29 | 8 months ago | A general purpose application that aligns text. |
| [genex](https://github.com/alixaxel/genex) | 58 | 3 | 2015-03-09 | 6 months ago | Count and expand Regular Expressions into all matching Strings. |
| [guesslanguage](https://github.com/endeveit/guesslanguage) | 47 | 1 | 2014-12-16 | 2 years ago | Functions to determine the natural language of a unicode text. |
| [go-vcard](https://github.com/emersion/go-vcard) | 45 | 3 | 2017-03-21 | 2 months ago | Parse and format vCard. |
| [allot](https://github.com/sbstjn/allot) | 44 | 1 | 2016-10-16 | 1 year ago | Placeholder and wildcard text parsing for CLI tools and bots. |
| [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) | 44 | 3 | 2017-11-15 | 3 months ago | Fixed-width text formatting (encoder/decoder with reflection). |
| [goregen](https://github.com/zach-klippenstein/goregen) | 44 | 2 | 2014-12-27 | 10 months ago | Library for generating random strings from regular expressions. |
| [did](https://github.com/ockam-network/did) | 35 | 9 | 2018-11-02 | 6 days ago | DID (Decentralized Identifiers) Parser and Stringer in Go. |
| [gonameparts](https://github.com/polera/gonameparts) | 31 | 0 | 2015-05-17 | 11 months ago | Parses human names into individual name parts. |
| [slugify](https://github.com/avelino/slugify) | 28 | 2 | 2015-04-13 | 2 years ago | Go slugify application that handles string. |
| [codetree](https://github.com/aerogo/codetree) | 15 | 2 | 2016-11-26 | 9 months ago | Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure. |
| [pagser](https://github.com/foolin/pagser) | 11 | 1 | 2020-04-19 | 2 months ago | Pagser is a simple, extensible, configurable parse and deserialize html page to struct based on goquery and struct tags for golang crawler. |
| [enca](https://github.com/endeveit/enca) | 8 | 1 | 2014-12-17 | 4 years ago | Minimal cgo bindings for [libenca](http://cihar.com/software/enca/). |
| [syndfeed](https://github.com/zhengchun/syndfeed) | 7 | 1 | 2017-04-07 | 2 years ago | A syndication feed for Atom 1.0 and RSS 2.0. |
| [bbConvert](https://github.com/CalebQ42/bbConvert) | 5 | 1 | 2016-04-15 | 3 years ago | Converts bbCode to HTML that allows you to add support for custom bbCode tags. |
| [ltsv](https://github.com/Wing924/ltsv) | 5 | 1 | 2019-05-12 | 1 year ago | High performance [LTSV (Labeled Tab Separated Value)](http://ltsv.org/) reader for Go. |
| [doi](https://github.com/hscells/doi) | 4 | 1 | 2017-08-02 | 2 years ago | Document object identifier (doi) parser in Go. |
| [encoding](https://github.com/mickep76/encoding) | 3 | 1 | 2018-04-06 | 8 months ago | Package provides a generic interface to encoders and decodersa. |

### Utility
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [xurls](https://github.com/mvdan/xurls) | 657 | 17 | 2015-01-12 | 5 days ago | Extract urls from text. |
| [gotabulate](https://github.com/bndr/gotabulate) | 242 | 7 | 2014-08-21 | 3 years ago | Easily pretty-print your tabular data with Go. |
| [radix](https://github.com/yourbasic/radix) | 160 | 5 | 2017-06-09 | 2 years ago | fast string sorting algorithm. |
| [parth](https://github.com/codemodus/parth) | 37 | 3 | 2015-04-06 | 1 year ago | URL path segmentation parsing. |
| [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) | 26 | 2 | 2018-09-09 | 1 year ago | A sanitization-based swear filter for Go. |
| [xj2go](https://github.com/stackerzzq/xj2go) | 19 | 2 | 2017-09-19 | 5 months ago | Convert xml or json to go struct. |
| [kace](https://github.com/codemodus/kace) | 14 | 1 | 2015-06-04 | 1 year ago | Common case conversions covering common initialisms. |
| [tagify](https://github.com/zoomio/tagify) | 11 | 1 | 2018-03-20 | 4 months ago | Produces a set of tags from given source. |
| [parseargs-go](https://github.com/txgruppi/parseargs-go) | 8 | 1 | 2016-02-24 | 3 years ago | string argument parser that understands quotes and backslashes. |
| [TySug](https://github.com/Dynom/TySug) | 6 | 1 | 2018-06-05 | 10 hours ago | Alternative suggestions with respect to keyboard layouts. |
| [textwrap](https://github.com/isbm/textwrap) | 1 | 1 | 2019-07-26 | 11 months ago | Implementation of `textwrap` module from Python. |

## Third-party APIs
        
*Libraries for accessing third party APIs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-github](https://github.com/google/go-github) | 6006 | 191 | 2013-05-24 | 2 days ago | Go library for accessing the GitHub REST API v3. |
| [aws-sdk-go](https://github.com/aws/aws-sdk-go) | 5929 | 255 | 2014-12-05 | 22 hours ago | The official AWS SDK for the Go programming language. |
| [google-api-go-client](https://github.com/googleapis/google-api-go-client) | 2369 | 136 | 2014-11-24 | 3 hours ago | Auto-generated Google APIs for Go. |
| [google-cloud-go](https://github.com/googleapis/google-cloud-go) | 2285 | 226 | 2014-05-09 | 1 hour ago | Google Cloud APIs Go Client Library. |
| [discordgo](https://github.com/bwmarrin/discordgo) | 1398 | 43 | 2015-11-01 | 14 hours ago | Go bindings for the Discord Chat API. |
| [stripe-go](https://github.com/stripe/stripe-go) | 1174 | 39 | 2014-06-05 | 19 hours ago | Go client for the Stripe API. |
| [anaconda](https://github.com/ChimeraCoder/anaconda) | 1052 | 21 | 2013-03-04 | 4 months ago | Go client library for the Twitter 1.1 API. |
| [minio-go](https://github.com/minio/minio-go) | 1045 | 42 | 2015-05-02 | 19 hours ago | Minio Go Library for Amazon S3 compatible cloud storage. |
| [go-twitter](https://github.com/dghubble/go-twitter) | 1033 | 31 | 2015-04-11 | 2 days ago | Go client library for the Twitter v1.1 APIs. |
| [facebook](https://github.com/huandu/facebook) | 892 | 98 | 2012-07-28 | 2 weeks ago | Go Library that supports the Facebook Graph API. |
| [githubv4](https://github.com/shurcooL/githubv4) | 668 | 20 | 2017-05-27 | 4 days ago | Go library for accessing the GitHub GraphQL API v4. |
| [webhooks](https://github.com/go-playground/webhooks) | 524 | 15 | 2015-10-25 | 5 days ago | Webhook receiver for GitHub and Bitbucket. |
| [paypal](https://github.com/plutov/paypal) | 371 | 15 | 2015-10-14 | 1 month ago | Wrapper for PayPal payment API. |
| [geo-golang](https://github.com/codingsince1985/geo-golang) | 353 | 12 | 2014-12-04 | 2 months ago | Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs. |
| [go-marathon](https://github.com/gambol99/go-marathon) | 192 | 13 | 2015-02-11 | 4 months ago | Go library for interacting with Mesosphere's Marathon PAAS. |
| [ethrpc](https://github.com/onrik/ethrpc) | 189 | 13 | 2017-01-24 | 1 day ago | Go bindings for Ethereum JSON RPC API. |
| [trello](https://github.com/adlio/trello) | 143 | 6 | 2016-09-24 | 2 weeks ago | Go wrapper for the Trello API. |
| [gostorm](https://github.com/jsgilmore/gostorm) | 124 | 11 | 2013-07-22 | 2 years ago | GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. |
| [medium-sdk-go](https://github.com/Medium/medium-sdk-go) | 124 | 122 | 2015-09-26 | 1 year ago | Golang SDK for Medium's OAuth2 API. |
| [hipchat](https://github.com/daneharrigan/hipchat) | 112 | 7 | 2013-04-28 | 3 years ago | A golang package to communicate with HipChat over XMPP. |
| [go-trending](https://github.com/andygrunwald/go-trending) | 109 | 7 | 2015-07-04 | 5 months ago | Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github. |
| [hipchat](https://github.com/andybons/hipchat) | 107 | 8 | 2012-10-20 | 4 years ago | This project implements a golang client library for the Hipchat API. |
| [wit-go](https://github.com/wit-ai/wit-go) | 90 | 15 | 2018-08-20 | 1 month ago | Go client for wit.ai HTTP API. |
| [cachet](https://github.com/andygrunwald/cachet) | 83 | 8 | 2015-10-31 | 2 years ago | Go client library for [Cachet (open source status page system)](https://cachethq.io/). |
| [pushover](https://github.com/gregdel/pushover) | 75 | 3 | 2015-02-19 | 3 months ago | Go wrapper for the Pushover API. |
| [igdb](https://github.com/Henry-Sarabia/igdb) | 60 | 2 | 2017-08-24 | 5 months ago | Go client for the [Internet Game Database API](https://api.igdb.com/). |
| [clarifai-go](https://github.com/Clarifai/clarifai-go) | 57 | 36 | 2015-09-28 | 2 years ago | Go client library for interfacing with the Clarifai API. |
| [megos](https://github.com/andygrunwald/megos) | 56 | 5 | 2015-10-02 | 2 years ago | Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster. |
| [go-circleci](https://github.com/jszwedko/go-circleci) | 53 | 3 | 2015-08-14 | 8 months ago | Go client library for interacting with CircleCI's API. |
| [gads](https://github.com/emiddleton/gads) | 47 | 7 | 2014-01-20 | 10 months ago | Google Adwords Unofficial API. |
| [gosip](https://github.com/koltyakov/gosip) | 43 | 2 | 2019-01-26 | 1 month ago | Go client library SharePoint API. |
| [simples3](https://github.com/rhnvrm/simples3) | 43 | 2 | 2018-12-06 | 3 months ago | Simple no frills AWS S3 Library using REST with V4 Signing written in Go. |
| [gomusicbrainz](https://github.com/michiwend/gomusicbrainz) | 41 | 7 | 2014-09-10 | 1 month ago | Go MusicBrainz WS2 client library. |
| [go-amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) | 40 | 1 | 2016-11-15 | 2 years ago | Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html). |
| [go-xkcd](https://github.com/nishanths/go-xkcd) | 40 | 4 | 2016-02-26 | 2 months ago | Go client for the xkcd API. |
| [uptimerobot](https://github.com/bitfield/uptimerobot) | 38 | 3 | 2018-05-29 | 2 months ago | Go wrapper and command-line client for the Uptime Robot v2 API. |
| [fcm](https://github.com/maddevsio/fcm) | 37 | 4 | 2017-01-06 | 4 months ago | Go library for Firebase Cloud Messaging. |
| [gogtrends](https://github.com/groovili/gogtrends) | 35 | 3 | 2018-12-27 | 1 month ago | Google Trends Unofficial API. |
| [golyrics](https://github.com/mamal72/golyrics) | 34 | 3 | 2016-11-18 | 2 years ago | Golyrics is a Go library to fetch music lyrics data from the Wikia website. |
| [mixpanel](https://github.com/dukex/mixpanel) | 33 | 3 | 2014-05-20 | 5 months ago | Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications. |
| [golang-tmdb](https://github.com/cyruzin/golang-tmdb) | 33 | 1 | 2019-01-11 | 6 days ago | Golang wrapper for The Movie Database API v3. |
| [go-unsplash](https://github.com/hbagdi/go-unsplash) | 30 | 1 | 2017-01-19 | 4 months ago | Go client library for the [Unsplash.com](https://unsplash.com) API. |
| [gcm](https://github.com/TheOrioli/gcm) | 29 | 3 | 2015-11-09 | 4 years ago | Go library for Google Cloud Messaging. |
| [translate](https://github.com/nuveo/translate) | 29 | 27 | 2015-07-13 | 4 years ago | Go online translation package. |
| [gami](https://github.com/bit4bit/gami) | 28 | 4 | 2014-05-14 | 2 years ago | Go library for Asterisk Manager Interface. |
| [ynab.go](https://github.com/brunomvsouza/ynab.go) | 27 | 1 | 2018-07-13 | 7 months ago | Go wrapper for the YNAB API. |
| [twitter-scraper](https://github.com/n0madic/twitter-scraper) | 27 | 4 | 2018-11-29 | 9 hours ago | Scrape the Twitter Frontend API without authentication and limits. |
| [go-spotify](https://github.com/rapito/go-spotify) | 26 | 2 | 2014-10-30 | 2 years ago | Go Library to access Spotify WEB API. |
| [go-steam](https://github.com/sostronk/go-steam) | 21 | 10 | 2014-11-23 | 6 months ago | Go Library to interact with Steam game servers. |
| [patreon-go](https://github.com/mxpv/patreon-go) | 20 | 4 | 2017-08-06 | 10 months ago | Go library for Patreon API. |
| [go-shopify](https://github.com/rapito/go-shopify) | 20 | 1 | 2014-10-28 | 2 years ago | Go Library to make CRUD request to the Shopify API. |
| [go-twitch](https://github.com/knspriggs/go-twitch) | 19 | 5 | 2016-06-28 | 2 years ago | Go client for interacting with the Twitch v3 API. |
| [brewerydb](https://github.com/naegelejd/brewerydb) | 17 | 3 | 2015-04-15 | 5 years ago | Go library for accessing the BreweryDB API. |
| [textbelt](https://github.com/farmergreg/textbelt) | 17 | 2 | 2015-09-01 | 4 years ago | Go client for the textbelt.com txt messaging API. |
| [codeship-go](https://github.com/codeship/codeship-go) | 16 | 23 | 2017-09-08 | 1 year ago | Go client library for interacting with Codeship's API v2. |
| [go-myanimelist](https://github.com/nstratos/go-myanimelist) | 16 | 2 | 2015-05-03 | 1 month ago | Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api). |
| [coinpaprika-api-go-client](https://github.com/coinpaprika/coinpaprika-api-go-client) | 12 | 8 | 2018-09-25 | 1 month ago | Go client library for interacting with Coinpaprika's API. |
| [go-google-analytics](https://github.com/chonthu/go-google-analytics) | 12 | 2 | 2015-06-01 | 5 years ago | Simple wrapper for easy google analytics reporting. |
| [lastpass-go](https://github.com/ansd/lastpass-go) | 12 | 0 | 2019-07-11 | 2 weeks ago | Go client library for the [LastPass](https://www.lastpass.com/) API. |
| [airtable](https://github.com/mehanizm/airtable) | 12 | 1 | 2020-04-12 | 3 months ago | Go client library for the [Airtable API](https://airtable.com/api). |
| [go-hacknews](https://github.com/PaulRosset/go-hacknews) | 11 | 2 | 2017-08-10 | 2 years ago | Tiny Go client for HackerNews API. |
| [smitego](https://github.com/sergiotapia/smitego) | 10 | 0 | 2013-12-11 | 6 years ago | Go package to wraps access to the Smite game API. |
| [go-here](https://github.com/abdullahselek/go-here) | 9 | 1 | 2019-07-07 | 1 month ago | Go client library around the HERE location based APIs. |
| [rrdaclient](https://github.com/Omie/rrdaclient) | 8 | 1 | 2014-09-15 | 5 years ago | Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP. |
| [go-sptrans](https://github.com/sergioaugrod/go-sptrans) | 7 | 3 | 2017-09-11 | 2 years ago | Go client library for the SPTrans Olho Vivo API. |
| [go-google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) | 7 | 0 | 2016-10-24 | 3 years ago | Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/). |
| [go-postman-collection](https://github.com/rbretecher/go-postman-collection) | 7 | 2 | 2019-11-16 | 6 days ago | Go module to work with [Postman Collections](https://learning.getpostman.com/docs/postman/collections/creating-collections/) (compatible with Insomnia). |
| [google-play-scraper](https://github.com/n0madic/google-play-scraper) | 7 | 1 | 2019-09-20 | 1 week ago | Get data from Google Play Store. |
| [go-sophos](https://github.com/esurdam/go-sophos) | 6 | 2 | 2018-09-05 | 9 months ago | Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies. |
| [slack](https://github.com/nlopes/slack) | 6 | 0 | 2015-01-24 | 5 months ago | Slack API in Go. |
| [gumblr](https://github.com/mattcunningham/gumblr) | 6 | 1 | 2015-07-09 | 3 years ago | Go wrapper for the Tumblr v2 API. |
| [gopaapi5](https://github.com/utekaravinash/gopaapi5) | 6 | 2 | 2020-02-15 | 3 months ago | Go Client Library for [Amazon Product Advertising API 5.0](https://webservices.amazon.com/paapi5/documentation/). |
| [go-zooz](https://github.com/gojuno/go-zooz) | 5 | 15 | 2017-07-04 | 2 weeks ago | Go client for the Zooz API. |
| [go-chronos](https://github.com/axelspringer/go-chronos) | 3 | 9 | 2017-10-23 | 2 years ago | Go library for interacting with the [Chronos](https://mesos.github. |
| [libgoffi](https://github.com/clevabit/libgoffi) | 2 | 7 | 2019-08-03 | 11 months ago | Library adapter toolbox for native [libffi](http://sourceware. |
| [playlyfe-go-sdk](https://github.com/playlyfe/playlyfe-go-sdk) | 1 | 4 | 2015-05-25 | 4 years ago | The Playlyfe Rest API Go SDK. |
| [tripadvisor-golang](https://github.com/mrbenosborne/tripadvisor-golang) | 1 | 1 | 2019-04-15 | 9 months ago | Go wrapper for the TripAdvisor API. |
| [vl-go](https://github.com/verifid/vl-go) | 1 | 1 | 2019-02-09 | 4 months ago | Go client library around the VerifID identity verification layer API. |

## Utilities
        
*General utilities and tools to make your life easier.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [fzf](https://github.com/junegunn/fzf) | 30906 | 379 | 2013-10-23 | 1 hour ago | Command-line fuzzy finder written in Go. |
| [hub](https://github.com/github/hub) | 19987 | 475 | 2009-12-05 | 23 hours ago | wrap git commands with additional functionality to interact with github from the terminal. |
| [delve](https://github.com/go-delve/delve) | 13355 | 387 | 2014-05-20 | 5 months ago | Go debugger. |
| [ctop](https://github.com/bcicen/ctop) | 10237 | 166 | 2016-12-27 | 1 day ago | [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics. |
| [wuzz](https://github.com/asciimoo/wuzz) | 9076 | 180 | 2017-01-30 | 1 month ago | Interactive cli tool for HTTP inspection. |
| [sqlx](https://github.com/jmoiron/sqlx) | 8720 | 184 | 2013-01-28 | 1 week ago | provides a set of extensions on top of the excellent built-in database/sql package. |
| [goreleaser](https://github.com/goreleaser/goreleaser) | 5971 | 83 | 2016-12-21 | 1 day ago | Deliver Go binaries as fast and easily as possible. |
| [peco](https://github.com/peco/peco) | 5953 | 137 | 2014-06-06 | 1 week ago | Simplistic interactive filtering tool. |
| [usql](https://github.com/xo/usql) | 5917 | 121 | 2017-03-02 | 1 day ago | usql is a universal command-line interface for SQL databases. |
| [godropbox](https://github.com/dropbox/godropbox) | 3891 | 246 | 2014-06-22 | 2 weeks ago | Common libraries for writing Go services/applications from Dropbox. |
| [realize](https://github.com/oxequa/realize) | 3737 | 71 | 2016-07-12 | 2 months ago | Go build system with file watchers and live reload. Run, build and watch file changes with custom paths. |
| [goreporter](https://github.com/360EntSecGroup-Skylar/goreporter) | 2712 | 99 | 2017-03-27 | 1 year ago | Golang tool that does static analysis, unit testing, code review and generate code quality report. |
| [hystrix-go](https://github.com/afex/hystrix-go) | 2701 | 89 | 2013-12-15 | 2 months ago | Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker. |
| [task](https://github.com/go-task/task) | 2467 | 41 | 2017-02-27 | 1 week ago | simple "Make" alternative. |
| [panicparse](https://github.com/maruel/panicparse) | 2296 | 38 | 2015-02-02 | 1 day ago | Groups similar goroutines and colorizes stack dump. |
| [minify](https://github.com/tdewolff/minify) | 2174 | 49 | 2014-05-21 | 1 day ago | Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats. |
| [go-funk](https://github.com/thoas/go-funk) | 1963 | 35 | 2016-12-30 | 1 day ago | Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...). |
| [storm](https://github.com/asdine/storm) | 1582 | 44 | 2016-01-10 | 1 week ago | Simple and powerful toolkit for BoltDB. |
| [mc](https://github.com/minio/mc) | 1572 | 50 | 2015-01-16 | 1 hour ago | Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems. |
| [mmake](https://github.com/tj/mmake) | 1509 | 28 | 2017-02-15 | 4 months ago | Modern Make. |
| [mole](https://github.com/davrodpin/mole) | 1372 | 32 | 2018-10-04 | 3 days ago | cli app to easily create ssh tunnels. |
| [mergo](https://github.com/imdario/mergo) | 1199 | 18 | 2013-03-11 | 4 days ago | Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements. |
| [spinner](https://github.com/briandowns/spinner) | 1175 | 14 | 2014-12-13 | 3 weeks ago | Go package to easily provide a terminal spinner with options. |
| [boilr](https://github.com/tmrts/boilr) | 1133 | 30 | 2015-12-19 | 2 months ago | Blazingly fast CLI tool for creating projects from boilerplate templates. |
| [filetype](https://github.com/h2non/filetype) | 1133 | 25 | 2015-09-24 | 3 weeks ago | Small package to infer the file type checking the magic numbers signature. |
| [circuitbreaker](https://github.com/rubyist/circuitbreaker) | 890 | 21 | 2014-07-17 | 9 months ago | Circuit Breakers in Go. |
| [jump](https://github.com/gsamokovarov/jump) | 864 | 15 | 2015-08-16 | 3 months ago | Jump helps you navigate faster by learning your habits. |
| [gtm](https://github.com/git-time-metric/gtm) | 809 | 28 | 2016-06-19 | 1 year ago | Simple, seamless, lightweight time tracking for Git. |
| [immortal](https://github.com/immortal/immortal) | 656 | 14 | 2016-06-30 | 1 month ago | \*nix cross-platform (OS agnostic) supervisor. |
| [htcat](https://github.com/htcat/htcat) | 517 | 18 | 2013-08-05 | 1 year ago | Parallel and Pipelined HTTP GET Utility. |
| [hostctl](https://github.com/guumaster/hostctl) | 510 | 9 | 2020-03-14 | 1 week ago | A CLI tool to manage /etc/hosts with easy commands. |
| [circuit](https://github.com/cep21/circuit) | 469 | 12 | 2017-12-23 | 2 months ago | An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern. |
| [go-dry](https://github.com/ungerik/go-dry) | 451 | 13 | 2014-02-28 | 2 years ago | DRY (don't repeat yourself) package for Go. |
| [godaemon](https://github.com/VividCortex/godaemon) | 442 | 29 | 2013-08-01 | 4 weeks ago | Utility to write daemons. |
| [gopencils](https://github.com/bndr/gopencils) | 432 | 14 | 2014-06-23 | 1 year ago | Small and simple package to easily consume REST APIs. |
| [koazee](https://github.com/wesovilabs/koazee) | 406 | 12 | 2018-11-09 | 4 months ago | Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays. |
| [ergo](https://github.com/cristianoliveira/ergo) | 392 | 6 | 2017-08-19 | 6 days ago | The management of multiple local services running over different ports made easy. |
| [request](https://github.com/mozillazg/request) | 384 | 15 | 2014-12-21 | 7 months ago | Go HTTP Requests for Humans™. |
| [go-rate](https://github.com/beefsack/go-rate) | 307 | 10 | 2014-08-25 | 2 years ago | Timed rate limiter for Go. |
| [mimetype](https://github.com/gabriel-vasile/mimetype) | 290 | 7 | 2018-07-02 | 1 week ago | Package for MIME type detection based on magic numbers. |
| [clockwork](https://github.com/jonboulle/clockwork) | 283 | 6 | 2014-09-09 | 3 weeks ago | A simple fake clock for golang. |
| [deepcopier](https://github.com/ulule/deepcopier) | 283 | 17 | 2015-07-24 | 2 months ago | Simple struct copying for Go. |
| [gubrak](https://github.com/novalagung/gubrak) | 270 | 5 | 2018-03-09 | 2 months ago | Golang utility library with syntactic sugar. It's like lodash, but for golang. |
| [gohper](https://github.com/cosiner/gohper) | 253 | 20 | 2015-03-23 | 3 years ago | Various tools/modules help for development. |
| [gohper](https://github.com/zhuah/gohper) | 252 | 20 | 2015-03-23 | 3 years ago | Various tools/modules help for development. |
| [retry](https://github.com/kamilsk/retry) | 246 | 3 | 2016-11-02 | 1 month ago | The most advanced functional mechanism to perform actions repetitively until successful. |
| [serve](https://github.com/syntaqx/serve) | 218 | 5 | 2019-01-10 | 3 weeks ago | A static http server anywhere you need. |
| [go-trigger](https://github.com/sadlil/go-trigger) | 202 | 12 | 2015-10-19 | 3 years ago | Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project. |
| [gotenv](https://github.com/subosito/gotenv) | 182 | 3 | 2013-08-27 | 10 months ago | Load environment variables from `.env` or any `io.Reader` in Go. |
| [util](https://github.com/shomali11/util) | 179 | 9 | 2017-05-24 | 4 months ago | Collection of useful utility functions. (strings, concurrency, manipulations, ...). |
| [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) | 165 | 4 | 2016-11-08 | 11 months ago | go:generate tool for wrapping symbols exported by golang plugins (1.8 only). |
| [delve](https://github.com/derekparker/delve) | 159 | 2 | 2020-02-18 | 2 months ago | Go debugger. |
| [death](https://github.com/vrecan/death) | 158 | 4 | 2015-03-09 | 3 weeks ago | Managing go application shutdown with signals. |
| [rerun](https://github.com/ivpusic/rerun) | 157 | 7 | 2014-12-10 | 2 years ago | Recompiling and rerunning go apps when source changes. |
| [moldova](https://github.com/StabbyCutyou/moldova) | 153 | 5 | 2016-01-30 | 2 years ago | Utility for generating random data based on an input template. |
| [toolbox](https://github.com/viant/toolbox) | 148 | 15 | 2016-06-13 | 22 hours ago | Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer. |
| [apm](https://github.com/topfreegames/apm) | 143 | 16 | 2015-11-18 | 3 years ago | Process manager for Golang applications with an HTTP API. |
| [robustly](https://github.com/VividCortex/robustly) | 143 | 16 | 2013-07-08 | 4 weeks ago | Runs functions resiliently, catching and restarting panics. |
| [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) | 131 | 6 | 2015-10-12 | 10 months ago | XML Sitemap generator written in Go. |
| [chyle](https://github.com/antham/chyle) | 128 | 6 | 2016-11-17 | 1 month ago | Changelog generator using a git repository with multiple configuration possibilities. |
| [onecache](https://github.com/adelowo/onecache) | 112 | 6 | 2017-04-14 | 2 months ago | Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc). |
| [lrserver](https://github.com/jaschaephraim/lrserver) | 106 | 5 | 2014-07-15 | 2 years ago | LiveReload server for Go. |
| [go-bsdiff](https://github.com/gabstv/go-bsdiff) | 102 | 1 | 2019-02-23 | 1 year ago | Pure Go bsdiff and bspatch libraries and CLI tools. |
| [xferspdy](https://github.com/monmohan/xferspdy) | 80 | 4 | 2015-05-22 | 3 years ago | Xferspdy provides binary diff and patch library in golang. |
| [mssqlx](https://github.com/linxGnu/mssqlx) | 79 | 9 | 2016-12-26 | 2 months ago | Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind. |
| [pm](https://github.com/VividCortex/pm) | 77 | 17 | 2013-11-17 | 4 weeks ago | Process (i.e. goroutine) manager with an HTTP API. |
| [go-health](https://github.com/Talento90/go-health) | 76 | 2 | 2018-02-13 | 2 years ago | Health package simplifies the way you add health check to your services. |
| [unis](https://github.com/esemplastic/unis) | 69 | 4 | 2017-05-06 | 3 years ago | Common Architecture™ for String Utilities in Go. |
| [nostromo](https://github.com/pokanop/nostromo) | 68 | 3 | 2019-07-13 | 1 month ago | CLI for building powerful aliases. |
| [cli](https://github.com/create-go-app/cli) | 68 | 3 | 2019-12-30 | 1 week ago | A powerful CLI for create a new production-ready project with backend (Golang), frontend (JavaScript, TypeScript) & deploy automation (Ansible, Docker) by running one command. |
| [netbug](https://github.com/e-dard/netbug) | 66 | 1 | 2015-03-05 | 4 years ago | Easy remote profiling of your services. |
| [repeat](https://github.com/ssgreg/repeat) | 65 | 5 | 2017-11-22 | 4 days ago | Go implementation of different backoff strategies useful for retrying operations and heartbeating. |
| [goseaweedfs](https://github.com/linxGnu/goseaweedfs) | 63 | 7 | 2017-07-20 | 4 months ago | SeaweedFS client library with almost full features. |
| [multitick](https://github.com/VividCortex/multitick) | 62 | 15 | 2013-12-10 | 4 weeks ago | Multiplexor for aligned tickers. |
| [sorty](https://github.com/jfcg/sorty) | 60 | 2 | 2019-02-18 | 1 week ago | Fast Concurrent / Parallel Sorting. |
| [handy](https://github.com/miguelpragier/handy) | 57 | 7 | 2018-06-13 | 3 weeks ago | Many utilities and helpers like string handlers/formatters and validators. |
| [minquery](https://github.com/icza/minquery) | 55 | 3 | 2016-11-16 | 4 months ago | MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off). |
| [mimemagic](https://github.com/zRedShift/mimemagic) | 54 | 1 | 2018-10-11 | 1 year ago | Pure Go ultra performant MIME sniffing library/utility. |
| [go-astitodo](https://github.com/asticode/go-astitodo) | 51 | 2 | 2016-10-17 | 1 year ago | Parse TODOs in your GO code. |
| [go-pattern-match](https://github.com/alexpantyukhin/go-pattern-match) | 50 | 2 | 2018-12-11 | 4 weeks ago | Pattern matching libray. |
| [golog](https://github.com/mlimaloureiro/golog) | 49 | 3 | 2016-01-09 | 1 year ago | Easy and lightweight CLI tool to time track your tasks. |
| [goreadability](https://github.com/philipjkim/goreadability) | 49 | 6 | 2016-04-20 | 1 year ago | Webpage summary extractor using Facebook Open Graph and arc90's readability. |
| [gaper](https://github.com/maxcnunes/gaper) | 44 | 0 | 2018-06-16 | 7 months ago | Builds and restarts a Go project when it crashes or some watched file changes. |
| [taskctl](https://github.com/taskctl/taskctl) | 44 | 2 | 2019-11-12 | 1 week ago | Concurrent task runner. |
| [countries](https://github.com/biter777/countries) | 43 | 4 | 2019-04-22 | 2 weeks ago | Full implementation of ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and IANA ccTLD standarts. |
| [copy-pasta](https://github.com/jutkko/copy-pasta) | 42 | 4 | 2017-01-28 | 1 month ago | Universal multi-workstation clipboard that uses S3 like backend for the storage. |
| [goback](https://github.com/carlescere/goback) | 42 | 1 | 2015-03-13 | 2 years ago | Go simple exponential backoff package. |
| [golarm](https://github.com/msempere/golarm) | 41 | 1 | 2015-08-14 | 5 years ago | Fire alarms with system events. |
| [intrinsic](https://github.com/mengzhuo/intrinsic) | 41 | 3 | 2017-06-13 | 3 years ago | Use x86 SIMD without writing any assembly code. |
| [pgo](https://github.com/arthurkushman/pgo) | 40 | 5 | 2018-12-26 | 6 days ago | Convenient functions for PHP community. |
| [retry](https://github.com/thedevsaddam/retry) | 40 | 1 | 2018-02-25 | 4 months ago | Simple and easy retry mechanism package for Go. |
| [retry-go](https://github.com/rafaeljesus/retry-go) | 37 | 1 | 2017-06-09 | 1 year ago | Retrying made simple and easy for golang. |
| [beyond](https://github.com/wesovilabs/beyond) | 36 | 1 | 2019-10-18 | 7 months ago | The Go tool that will drive you to the AOP world! |
| [gpath](https://github.com/tenntenn/gpath) | 36 | 3 | 2017-05-24 | 3 years ago | Library to simplify access struct fields with Go's expression in reflection. |
| [dbt](https://github.com/nikogura/dbt) | 33 | 2 | 2017-11-30 | 1 month ago | A framework for running self-updating signed binaries from a central, trusted repository. |
| [myhttp](https://github.com/inancgumus/myhttp) | 33 | 1 | 2017-09-13 | 2 years ago | Simple API to make HTTP GET requests with timeout support. |
| [scan](https://github.com/blockloop/scan) | 32 | 1 | 2017-11-27 | 1 month ago | Scan golang `sql.Rows` directly to structs, slices, or primitive types. |
| [rclient](https://github.com/zpatrick/rclient) | 31 | 3 | 2017-02-28 | 8 months ago | Readable, flexible, simple-to-use client for REST APIs. |
| [slice](https://github.com/psampaz/slice) | 30 | 1 | 2019-11-26 | 3 months ago | Type-safe functions for common Go slice operations. |
| [mongo-go-pagination](https://github.com/gobeam/mongo-go-pagination) | 27 | 1 | 2020-02-04 | 2 weeks ago | Mongodb Pagination for official mongodb/mongo-go-driver package which supports  both normal queries and Aggregation pipelines. |
| [filter](https://github.com/gookit/filter) | 26 | 5 | 2018-09-26 | 7 months ago | provide filtering, sanitizing, and conversion of Go data. |
| [cmd](https://github.com/commander-cli/cmd) | 26 | 3 | 2019-09-27 | 2 weeks ago | Library for executing shell commands on osx, windows and linux. |
| [cmd](https://github.com/SimonBaeumer/cmd) | 25 | 2 | 2019-09-27 | 3 months ago | Library for executing shell commands on osx, windows and linux. |
| [evaluator](https://github.com/nullne/evaluator) | 25 | 2 | 2017-04-27 | 1 month ago | Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend. |
| [generate](https://github.com/go-playground/generate) | 24 | 3 | 2015-11-15 | 3 years ago | runs go generate recursively on a specified path or environment variable and can filter by regex. |
| [gostrutils](https://github.com/ik5/gostrutils) | 23 | 3 | 2018-09-19 | 1 week ago | Collections of string manipulation and conversion functions. |
| [ugo](https://github.com/alxrm/ugo) | 22 | 1 | 2016-02-17 | 4 years ago | ugo is slice toolbox with concise syntax for Go. |
| [go-lock](https://github.com/viney-shih/go-lock) | 22 | 0 | 2020-04-30 | 2 months ago | go-lock is a lock library implementing read-write mutex and read-write trylock without starvation. |
| [goplaceholder](https://github.com/michiwend/goplaceholder) | 21 | 2 | 2014-10-12 | 4 years ago | a small golang lib to generate placeholder images. |
| [slicer](https://github.com/leaanthony/slicer) | 20 | 1 | 2019-01-10 | 6 months ago | Makes working with slices easier. |
| [r](https://github.com/is5/r) | 18 | 3 | 2020-02-20 | 5 months ago | Python-like `range()` experience for Go. |
| [go-httpheader](https://github.com/mozillazg/go-httpheader) | 16 | 2 | 2017-06-24 | 1 year ago | Go library for encoding structs into Header fields. |
| [rerate](https://github.com/abo/rerate) | 16 | 4 | 2016-05-24 | 3 years ago | Redis-based rate counter and rate limiter for Go. |
| [structs](https://github.com/PumpkinSeed/structs) | 16 | 2 | 2017-08-26 | 2 years ago | Implement simple functions to manipulate structs. |
| [ctxutil](https://github.com/posener/ctxutil) | 15 | 1 | 2018-07-30 | 4 months ago | A collection of utility functions for contexts. |
| [filler](https://github.com/yaronsumel/filler) | 15 | 1 | 2017-04-05 | 3 years ago | small utility to fill structs using "fill" tag. |
| [ghokin](https://github.com/antham/ghokin) | 15 | 2 | 2018-08-03 | 2 weeks ago | Parallelized formatter with no external dependencies for gherkin (cucumber, behat...). |
| [okrun](https://github.com/xta/okrun) | 15 | 2 | 2014-10-01 | 5 years ago | go run error steamroller. |
| [dlog](https://github.com/kirillDanshin/dlog) | 14 | 2 | 2016-07-04 | 3 years ago | Compile-time controlled logger to make your release smaller without removing debug calls. |
| [limiters](https://github.com/mennanov/limiters) | 14 | 1 | 2019-08-28 | 8 months ago | Rate limiters for distributed applications in Golang with configurable back-ends and distributed locks. |
| [equalizer](https://github.com/reugn/equalizer) | 14 | 1 | 2019-06-14 | 11 months ago | Quota manager and rate limiter collection for Go. |
| [rest-go](https://github.com/edermanoel94/rest-go) | 13 | 3 | 2019-07-29 | 4 months ago | A package that provide many helpful methods for working with rest api. |
| [jsend](https://github.com/clevergo/jsend) | 12 | 3 | 2020-01-14 | 3 weeks ago | JSend's implementation writen in Go. |
| [backscanner](https://github.com/icza/backscanner) | 11 | 2 | 2017-10-18 | 5 months ago | A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward. |
| [command](https://github.com/txgruppi/command) | 11 | 1 | 2015-08-24 | 4 years ago | Command pattern for Go with thread safe serial and parallel dispatcher. |
| [mimesniffer](https://github.com/aofei/mimesniffer) | 11 | 1 | 2018-12-20 | 4 days ago | A MIME type sniffer for Go. |
| [retry](https://github.com/shafreeck/retry) | 11 | 0 | 2018-07-18 | 5 months ago | A pretty simple library to ensure your work to be done. |
| [tome](https://github.com/cyruzin/tome) | 11 | 1 | 2019-04-12 | 3 months ago | Tome was designed to paginate simple RESTful APIs. |
| [shutdown](https://github.com/ztrue/shutdown) | 8 | 1 | 2018-11-17 | 1 year ago | App shutdown hooks for `os.Signal` handling. |
| [go-convert](https://github.com/Eun/go-convert) | 7 | 0 | 2019-06-07 | 3 months ago | Package go-convert enbles you to convert a value into another type. |
| [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) | 6 | 1 | 2019-05-16 | 5 months ago | Go package for working with Problem Details. |
| [retry](https://github.com/percolate/retry) | 6 | 34 | 2018-06-15 | 10 months ago | A simple but highly configurable retry package for Go. |
| [sliceconv](https://github.com/Henry-Sarabia/sliceconv) | 6 | 1 | 2019-02-15 | 5 months ago | Slice conversion between primitive types. |
| [ptr](https://github.com/gotidy/ptr) | 5 | 1 | 2019-12-25 | 7 months ago | Package that provide functions for simplified creation of pointers from constants of basic types. |
| [blank](https://github.com/Henry-Sarabia/blank) | 4 | 2 | 2019-02-13 | 1 year ago | Verify or remove blanks and whitespace from strings. |
| [silk](https://github.com/chrispassas/silk) | 4 | 1 | 2018-12-18 | 3 months ago | Read silk netflow files. |
| [olaf](https://github.com/btnguyen2k/olaf) | 1 | 1 | 2019-01-03 | 1 year ago | Twitter Snowflake implemented in Go. |
| [compare](https://github.com/posener/compare) | 1 | 1 | 2020-03-13 | 4 months ago | Enables more readable and easier comparison tasks. |
| [nfdump](https://github.com/chrispassas/nfdump) | 1 | 1 | 2020-04-08 | 3 months ago | Read nfdump netflow files. |

## UUID
        
*Libraries for working with UUIDs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [uuid](https://github.com/google/uuid) | 1994 | 48 | 2016-02-12 | 3 weeks ago | Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services. |
| [ulid](https://github.com/oklog/ulid) | 1963 | 41 | 2016-12-06 | 5 days ago | Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier). |
| [uuid](https://github.com/gofrs/uuid) | 817 | 19 | 2018-07-13 | 2 months ago | Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid. |
| [wuid](https://github.com/edwingeng/wuid) | 360 | 16 | 2018-01-27 | 1 week ago | An extremely fast unique number generator, 10-135 times faster than UUID. |
| [sno](https://github.com/muyo/sno) | 31 | 2 | 2019-05-26 | 3 months ago | Compact, sortable and fast unique IDs with embedded metadata. |
| [Goid](https://github.com/JakeHL/Goid) | 29 | 4 | 2017-05-19 | 1 year ago | Generate and Parse RFC4122 compliant V4 UUIDs. |
| [nanoid](https://github.com/aidarkhanov/nanoid) | 21 | 1 | 2019-07-02 | 1 month ago | A tiny and efficient Go unique string ID generator. |
| [uuid](https://github.com/agext/uuid) | 11 | 2 | 2016-02-03 | 4 months ago | Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier. |

## Validation
        
*Libraries for validation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [validator](https://github.com/go-playground/validator) | 5863 | 96 | 2015-02-12 | 2 hours ago | Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving. |
| [govalidator](https://github.com/asaskevich/govalidator) | 4284 | 103 | 2014-06-20 | 1 day ago | Validators and sanitizers for strings, numerics, slices and structs. |
| [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) | 1552 | 26 | 2016-06-22 | 1 hour ago | Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags. |
| [govalidator](https://github.com/thedevsaddam/govalidator) | 895 | 20 | 2017-09-13 | 1 month ago | Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. |
| [validate](https://github.com/gookit/validate) | 275 | 12 | 2018-07-16 | 1 week ago | Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features. |
| [checkdigit](https://github.com/osamingo/checkdigit) | 57 | 0 | 2019-04-05 | 7 months ago | Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.). |
| [jio](https://github.com/faceair/jio) | 44 | 2 | 2018-10-28 | 2 months ago | jio is a json schema validator similar to [joi](https://github.com/hapijs/joi). |
| [validate](https://github.com/gobuffalo/validate) | 37 | 6 | 2018-02-10 | 3 months ago | This package provides a framework for writing validations for Go applications. |
| [gody](https://github.com/guiferpa/gody) | 33 | 0 | 2018-11-01 | 3 months ago | :balloon: A lightweight struct validator for Go. |
| [terraform-validator](https://github.com/thazelart/terraform-validator) | 32 | 2 | 2019-05-29 | 6 months ago | A norms and conventions validator for Terraform. |
| [govalid](https://github.com/twharmon/govalid) | 17 | 1 | 2019-02-17 | 5 months ago | Fast, tag-based validation for structs. |

## Version Control
        
*Libraries for version control.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-git](https://github.com/src-d/go-git) | 5026 | 103 | 2015-10-22 | 3 months ago | highly extensible Git implementation in pure Go. |
| [git2go](https://github.com/libgit2/git2go) | 1509 | 53 | 2013-03-05 | 2 days ago | Go bindings for libgit2. |
| [hercules](https://github.com/src-d/hercules) | 1058 | 21 | 2016-12-12 | 3 months ago | gaining advanced insights from Git repository history. |
| [gh](https://github.com/rjeczalik/gh) | 75 | 6 | 2015-03-08 | 1 year ago | Scriptable server and net/http middleware for GitHub Webhooks. |
| [go-vcs](https://github.com/sourcegraph/go-vcs) | 73 | 42 | 2013-06-02 | 2 months ago | manipulate and inspect VCS repositories in Go. |
| [hgo](https://github.com/beyang/hgo) | 13 | 3 | 2014-06-18 | 4 years ago | Hgo is a collection of Go packages providing read-access to local Mercurial repositories. |

## Video
        
*Libraries for manipulating video.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [goav](https://github.com/giorgisio/goav) | 1291 | 45 | 2015-05-21 | 1 day ago | Comphrensive Go bindings for FFmpeg. |
| [m3u8](https://github.com/grafov/m3u8) | 732 | 36 | 2013-02-05 | 1 month ago | Parser and generator library of M3U8 playlists for Apple HLS. |
| [gmf](https://github.com/3d0c/gmf) | 616 | 31 | 2013-04-03 | 1 month ago | Go bindings for FFmpeg av\* libraries. |
| [go-astits](https://github.com/asticode/go-astits) | 321 | 16 | 2017-07-04 | 1 month ago | Parse and demux MPEG Transport Streams (.ts) natively in GO. |
| [go-astisub](https://github.com/asticode/go-astisub) | 262 | 7 | 2016-12-16 | 1 month ago | Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.). |
| [gst](https://github.com/ziutek/gst) | 156 | 9 | 2011-07-26 | 2 years ago | Go bindings for GStreamer. |
| [libvlc-go](https://github.com/adrg/libvlc-go) | 136 | 10 | 2015-01-06 | 1 week ago | Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player). |
| [go-m3u8](https://github.com/quangngotan95/go-m3u8) | 61 | 1 | 2018-11-06 | 2 months ago | Parser and generator library for Apple m3u8 playlists. |
| [v4l](https://github.com/korandiz/v4l) | 46 | 4 | 2016-10-25 | 2 years ago | Video capture library for Linux, written in Go. |
| [libgosubs](https://github.com/wargarblgarbl/libgosubs) | 14 | 1 | 2017-05-03 | 2 months ago | Subtitle format support for go. Supports .srt, .ttml, and .ass. |
| [go-mpd](https://github.com/unki2aut/go-mpd) | 3 | 1 | 2018-11-02 | 5 months ago | Parser and generator library for MPEG-DASH manifest files. |

## Web Frameworks
        
*Full stack web frameworks.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gin](https://github.com/gin-gonic/gin) | 40319 | 1283 | 2014-06-16 | 3 days ago | Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity. |
| [beego](https://github.com/astaxie/beego) | 24550 | 1256 | 2012-02-29 | 11 minutes ago | beego is an open-source, high-performance web framework for the Go programming language. |
| [echo](https://github.com/labstack/echo) | 17772 | 548 | 2015-03-01 | 2 days ago | High performance, minimalist Go web framework. |
| [revel](https://github.com/revel/revel) | 11780 | 547 | 2011-12-09 | 2 weeks ago | High-productivity web framework for the Go language. |
| [fiber](https://github.com/gofiber/fiber) | 6975 | 137 | 2020-01-16 | 12 hours ago | An Express.js inspired web framework build on Fasthttp. |
| [goa](https://github.com/goadesign/goa) | 3920 | 170 | 2014-12-05 | 2 days ago | Goa provides a holistic approach for developing remote APIs and microservices in Go. |
| [go-json-rest](https://github.com/ant0ine/go-json-rest) | 3420 | 160 | 2013-02-19 | 11 months ago | Quick and easy way to setup a RESTful JSON API. |
| [gizmo](https://github.com/nytimes/gizmo) | 3193 | 114 | 2015-12-15 | 6 hours ago | Microservice toolkit used by the New York Times. |
| [macaron](https://github.com/go-macaron/macaron) | 3008 | 145 | 2014-07-10 | 1 month ago | Macaron is a high productive and modular design web framework in Go. |
| [utron](https://github.com/gernest/utron) | 2160 | 69 | 2015-09-16 | 1 year ago | Lightweight MVC framework for Go(Golang). |
| [go-tigertonic](https://github.com/rcrowley/go-tigertonic) | 1001 | 45 | 2013-02-09 | 2 years ago | Go framework for building JSON web services inspired by Dropwizard. |
| [tango](https://github.com/lunny/tango) | 839 | 77 | 2014-12-17 | 1 year ago | Micro & pluggable web framework for Go. |
| [goyave](https://github.com/System-Glitch/goyave) | 442 | 17 | 2019-10-21 | 3 weeks ago | Feature-complete web framework aimed at clean code and fast development, with powerful built-in functionalities. |
| [gongular](https://github.com/mustafaakin/gongular) | 434 | 22 | 2016-06-22 | 3 weeks ago | Fast Go web framework with input mapping/validation and (DI) Dependency Injection. |
| [neo](https://github.com/ivpusic/neo) | 408 | 34 | 2015-02-04 | 2 years ago | Neo is minimal and fast Go Web Framework with extremely simple API. |
| [air](https://github.com/aofei/air) | 384 | 17 | 2016-07-20 | 4 days ago | An ideally refined web framework for Go. |
| [mango](https://github.com/paulbellamy/mango) | 351 | 21 | 2011-05-25 | 2 years ago | Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. |
| [gondola](https://github.com/rainycape/gondola) | 312 | 15 | 2014-07-25 | 1 year ago | The web framework for writing faster sites, faster. |
| [aero](https://github.com/aerogo/aero) | 292 | 17 | 2016-11-09 | 2 months ago | High-performance web framework for Go, reaches top scores in Lighthouse. |
| [golf](https://github.com/dinever/golf) | 243 | 18 | 2015-11-18 | 3 years ago | Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library. |
| [gearbox](https://github.com/gogearbox/gearbox) | 221 | 7 | 2020-04-25 | 15 hours ago | A web framework written in Go with a focus on high performance and memory optimization. |
| [gearbox](https://github.com/abahmed/gearbox) | 176 | 4 | 2020-04-25 | 2 months ago | A web framework written in Go with a focus on high performance and memory optimization. |
| [flamingo](https://github.com/i-love-flamingo/flamingo) | 153 | 22 | 2019-04-02 | 4 hours ago | Framework for pluggable web projects. Including a concept for modules and offering features for DI, Configareas, i18n, template engines, graphql, observability, security, events, routing & reverse routing etc. |
| [webgo](https://github.com/bnkamalesh/webgo) | 134 | 7 | 2015-12-16 | 1 week ago | A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc). |
| [hiboot](https://github.com/hidevopsio/hiboot) | 132 | 12 | 2018-03-16 | 6 days ago | hiboot is a high performance web application framework with auto configuration and dependency injection support. |
| [go-rest](https://github.com/ungerik/go-rest) | 120 | 10 | 2012-07-13 | 3 years ago | Small and evil REST framework for Go. |
| [flamingo-commerce](https://github.com/i-love-flamingo/flamingo-commerce) | 99 | 18 | 2019-04-02 | 9 hours ago | Providing e-commerce features using clean architecture like DDD and ports and adapters, that you can use to build flexible e-commerce applications. |
| [uadmin](https://github.com/uadmin/uadmin) | 97 | 9 | 2018-10-05 | 6 hours ago | Fully featured web framework for Golang, inspired by Django. |
| [ginrpc](https://github.com/xxjwxc/ginrpc) | 93 | 4 | 2019-06-22 | 1 week ago | Gin parameter automatic binding tool,gin rpc tools. |
| [microservice](https://github.com/claygod/microservice) | 73 | 9 | 2016-12-15 | 1 year ago | The framework for the creation of microservices, written in Golang. |
| [golax](https://github.com/fulldump/golax) | 72 | 7 | 2016-01-30 | 2 years ago | A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more. |
| [patron](https://github.com/beatlabs/patron) | 61 | 16 | 2019-01-30 | 2 hours ago | Patron is a microservice framework following best cloud practices with a focus on productivity. |
| [vox](https://github.com/aisk/vox) | 59 | 2 | 2014-12-24 | 2 months ago | A golang web framework for humans, inspired by Koa heavily. |
| [yarf](https://github.com/yarf-framework/yarf) | 56 | 3 | 2015-09-02 | 1 year ago | Fast micro-framework designed to build REST APIs and web services in a fast and simple way. |
| [fireball](https://github.com/zpatrick/fireball) | 51 | 4 | 2016-07-20 | 1 year ago | More "natural" feeling web framework. |
| [appy](https://github.com/appist/appy) | 47 | 2 | 2019-05-27 | 1 week ago | An opinionated productive web framework that helps scaling business easier. |
| [rux](https://github.com/gookit/rux) | 37 | 3 | 2018-08-05 | 3 days ago | Simple and fast web framework for build golang HTTP applications. |
| [goa](https://github.com/goa-go/goa) | 33 | 2 | 2019-07-26 | 7 months ago | goa is just like koajs for golang, it is a flexible, light, high-performance and extensible web framework based on middleware. |
| [api](https://github.com/resoursea/api) | 31 | 6 | 2015-01-24 | 5 years ago | REST framework for quickly writing resource based services. |
| [rex](https://github.com/goanywhere/rex) | 30 | 1 | 2014-10-16 | 2 years ago | Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`. |
| [banjo](https://github.com/nsheremet/banjo) | 11 | 1 | 2017-12-09 | 2 years ago | Very simple and fast web framework for Go. |
| [goweb](https://github.com/twharmon/goweb) | 6 | 1 | 2019-05-07 | 3 months ago | Web framework with routing, websockets, logging, middleware, static file server (optional gzip), and automatic TLS. |

### Middlewares
        

#### Actual middlewares
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [tollbooth](https://github.com/didip/tollbooth) | 1690 | 49 | 2015-05-17 | 2 months ago | Rate limit HTTP request handler. |
| [cors](https://github.com/rs/cors) | 1556 | 29 | 2014-10-25 | 1 month ago | Easily add CORS capabilities to your API. |
| [limiter](https://github.com/ulule/limiter) | 993 | 26 | 2015-10-02 | 1 week ago | Dead simple rate limit middleware for Go. |
| [go-server-timing](https://github.com/mitchellh/go-server-timing) | 779 | 19 | 2018-02-12 | 1 month ago | Add/parse Server-Timing header. |
| [go-fault](https://github.com/github/go-fault) | 308 | 86 | 2020-05-14 | 1 week ago | Fault injection middleware for Go. |
| [ln-paywall](https://github.com/philippgille/ln-paywall) | 106 | 4 | 2018-06-29 | 1 year ago | Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin). |
| [xff](https://github.com/sebest/xff) | 74 | 2 | 2014-12-22 | 1 year ago | Handle `X-Forwarded-For` header and friends. |
| [formjson](https://github.com/rs/formjson) | 34 | 2 | 2015-03-19 | 4 years ago | Transparently handle JSON input as a standard form POST. |
| [client-timing](https://github.com/posener/client-timing) | 18 | 1 | 2018-02-23 | 4 months ago | An HTTP client for Server-Timing header. |

#### Libraries for creating HTTP middlewares
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [negroni](https://github.com/urfave/negroni) | 6739 | 238 | 2014-05-18 | 1 week ago | Idiomatic HTTP middleware for Golang. |
| [alice](https://github.com/justinas/alice) | 2052 | 52 | 2014-05-25 | 4 months ago | Painless middleware chaining for Go. |
| [render](https://github.com/unrolled/render) | 1370 | 40 | 2014-06-10 | 3 months ago | Go package for easily rendering JSON, XML, and HTML template responses. |
| [stats](https://github.com/thoas/stats) | 566 | 16 | 2015-03-05 | 1 year ago | Go middleware that stores various information about your web application. |
| [interpose](https://github.com/carbocation/interpose) | 288 | 12 | 2014-07-20 | 3 years ago | Minimalist net/http middleware for golang. |
| [muxchain](https://github.com/stephens2424/muxchain) | 208 | 5 | 2014-05-03 | 1 year ago | Lightweight middleware for net/http. |
| [renderer](https://github.com/thedevsaddam/renderer) | 195 | 6 | 2017-11-07 | 1 year ago | Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go. |
| [rye](https://github.com/InVisionApp/rye) | 93 | 192 | 2016-10-06 | 1 year ago | Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context. |
| [gores](https://github.com/alioygur/gores) | 92 | 5 | 2015-12-25 | 2 months ago | Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs. |
| [chain](https://github.com/codemodus/chain) | 64 | 6 | 2015-05-14 | 1 year ago | Handler wrapper chaining with scoped data (net/context-based "middleware"). |
| [mediary](https://github.com/HereMobilityDevelopers/mediary) | 63 | 3 | 2020-03-23 | 1 month ago | add interceptors to `http.Client` to allow dumping/shaping/tracing/... of requests/responses. |
| [wrap](https://github.com/go-on/wrap) | 58 | 2 | 2014-02-16 | 1 year ago | Small middlewares package for net/http. |
| [catena](https://github.com/codemodus/catena) | 7 | 1 | 2015-07-30 | 1 year ago | http.Handler wrapper catenation (same API as "chain"). |

### Routers
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [mux](https://github.com/gorilla/mux) | 12385 | 302 | 2012-10-02 | 2 weeks ago | Powerful URL router and dispatcher for golang. |
| [httprouter](https://github.com/julienschmidt/httprouter) | 11550 | 309 | 2013-12-05 | 4 days ago | High performance router. Use this and the standard http handlers to form a very high performance web framework. |
| [chi](https://github.com/go-chi/chi) | 7893 | 186 | 2015-10-15 | 2 days ago | Small, fast and expressive HTTP router built on net/context. |
| [web](https://github.com/gocraft/web) | 1422 | 58 | 2013-11-16 | 1 year ago | Mux and middleware package in Go. |
| [bone](https://github.com/go-zoo/bone) | 1260 | 36 | 2014-11-19 | 1 year ago | Lightning Fast HTTP Multiplexer. |
| [fasthttprouter](https://github.com/buaazp/fasthttprouter) | 855 | 33 | 2015-12-13 | 1 year ago | High performance router forked from `httprouter`. The first router fit for `fasthttp`. |
| [goji](https://github.com/goji/goji) | 817 | 41 | 2015-11-16 | 1 year ago | Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`. |
| [gorouter](https://github.com/xujiajun/gorouter) | 481 | 15 | 2018-01-29 | 10 months ago | A simple and fast HTTP router for Go. |
| [httptreemux](https://github.com/dimfeld/httptreemux) | 434 | 23 | 2014-05-14 | 3 months ago | High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter. |
| [lars](https://github.com/go-playground/lars) | 383 | 15 | 2015-12-24 | 1 year ago | Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks. |
| [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) | 381 | 28 | 2015-10-27 | 5 months ago | An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs. |
| [siesta](https://github.com/VividCortex/siesta) | 352 | 27 | 2014-09-23 | 4 weeks ago | Composable framework to write middleware and handlers. |
| [vestigo](https://github.com/husobee/vestigo) | 263 | 16 | 2015-09-22 | 10 months ago | Performant, stand-alone, HTTP compliant URL Router for go web applications. |
| [router](https://github.com/gowww/router) | 157 | 6 | 2017-05-25 | 2 months ago | Lightning fast HTTP router fully compatible with the net/http.Handler interface. |
| [alien](https://github.com/gernest/alien) | 110 | 4 | 2016-01-30 | 1 year ago | Lightweight and fast http router from outer space. |
| [pure](https://github.com/go-playground/pure) | 103 | 5 | 2016-09-23 | 3 weeks ago | Is a lightweight HTTP router that sticks to the std "net/http" implementation. |
| [violetear](https://github.com/nbari/violetear) | 98 | 4 | 2015-06-19 | 8 months ago | Go HTTP router. |
| [Bxog](https://github.com/claygod/Bxog) | 95 | 8 | 2016-05-19 | 1 month ago | Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters. |
| [xmux](https://github.com/rs/xmux) | 88 | 6 | 2015-12-14 | 3 years ago | High performance muxer based on `httprouter` with `net/context` support. |
| [gorouter](https://github.com/vardius/gorouter) | 83 | 5 | 2016-07-14 | 1 week ago | GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`. |
| [bellt](https://github.com/GuilhermeCaruso/bellt) | 46 | 4 | 2019-02-21 | 1 month ago | A simple Go HTTP router. |
| [fastrouter](https://github.com/razonyang/fastrouter) | 19 | 2 | 2017-11-01 | 2 years ago | a fast, flexible HTTP router written in Go. |
| [route](https://github.com/goroute/route) | 6 | 3 | 2019-07-06 | 7 months ago | Simple yet powerful HTTP request multiplexer. |

## Windows
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-ole](https://github.com/go-ole/go-ole) | 648 | 38 | 2011-01-21 | 1 month ago | Win32 OLE implementation for golang. |
| [d3d9](https://github.com/gonutz/d3d9) | 105 | 7 | 2015-12-12 | 9 months ago | Go bindings for Direct3D9. |
| [gosddl](https://github.com/MonaxGT/gosddl) | 3 | 1 | 2018-12-04 | 1 year ago | Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL. |

## XML
        
*Libraries and tools for manipulating XML.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [zek](https://github.com/miku/zek) | 387 | 21 | 2017-11-23 | 2 months ago | Generate a Go struct from XML. |
| [xpath](https://github.com/antchfx/xpath) | 321 | 8 | 2016-10-09 | 2 weeks ago | XPath package for Go. |
| [xquery](https://github.com/antchfx/xquery) | 154 | 11 | 2016-10-09 | 2 years ago | XQuery lets you extract data from HTML/XML documents using XPath expression. |
| [xml2map](https://github.com/sbabiv/xml2map) | 24 | 1 | 2018-08-06 | 5 months ago | XML to MAP converter written Golang. |
| [XML-Comp](https://github.com/XML-Comp/XML-Comp) | 16 | 1 | 2016-10-25 | 2 years ago | Simple command line XML comparer that generates diffs of folders, files and tags. |
| [xmlwriter](https://github.com/shabbyrobe/xmlwriter) | 12 | 2 | 2017-04-11 | 5 months ago | Procedural XML generation API based on libxml2's xmlwriter module. |

## WebAssembly
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [tinygo](https://github.com/tinygo-org/tinygo) | 6569 | 143 | 2018-06-07 | 1 day ago | Go compiler for small places. Microcontrollers, WebAssembly, and command-line tools. Based on LLVM. |
| [dom](https://github.com/dennwc/dom) | 400 | 15 | 2018-06-30 | 10 months ago | DOM library. |
| [go-canvas](https://github.com/markfarnan/go-canvas) | 83 | 6 | 2019-05-05 | 4 days ago | Library to use HTML5 Canvas, with all drawing within go code. |
| [webapi](https://github.com/gowebapi/webapi) | 60 | 3 | 2019-02-08 | 4 months ago | Bindings for DOM and HTML generated from WebIDL. |
| [wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest) | 51 | 3 | 2018-07-14 | 2 months ago | Run Go WASM tests in your browser. |
| [vert](https://github.com/norunners/vert) | 43 | 4 | 2018-03-25 | 2 months ago | Interop between Go and JS values. |

## Files
        

## File Handling
        
*Libraries for handling files and file systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [afero](https://github.com/spf13/afero) | 3090 | 99 | 2014-10-28 | 1 week ago | FileSystem Abstraction System for Go. |
| [pdfcpu](https://github.com/pdfcpu/pdfcpu) | 1536 | 38 | 2017-06-18 | 1 week ago | PDF processor. |
| [notify](https://github.com/rjeczalik/notify) | 574 | 27 | 2014-09-08 | 1 week ago | File system event notification library with simple API, similar to os/signal. |
| [copy](https://github.com/otiai10/copy) | 173 | 4 | 2017-09-01 | 2 months ago | Copy directory recursively. |
| [bigfile](https://github.com/bigfile/bigfile) | 143 | 13 | 2019-07-15 | 5 months ago | A file transfer system, support to manage files with http api, rpc call and ftp client. |
| [go-csv-tag](https://github.com/artonge/go-csv-tag) | 73 | 1 | 2017-06-18 | 1 month ago | Load csv file using tag. |
| [afs](https://github.com/viant/afs) | 65 | 10 | 2019-08-19 | 5 days ago | Abstract File Storage (mem, scp, zip, tar, cloud: s3, gs) for Go. |
| [opc](https://github.com/qmuntal/opc) | 64 | 1 | 2018-11-06 | 6 days ago | Load Open Packaging Conventions (OPC) files for Go. |
| [vfs](https://github.com/C2FO/vfs) | 58 | 20 | 2017-08-01 | 2 weeks ago | A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS. |
| [skywalker](https://github.com/dixonwille/skywalker) | 57 | 4 | 2017-08-01 | 3 years ago | Package to allow one to concurrently go through a filesystem with ease. |
| [tarfs](https://github.com/posener/tarfs) | 41 | 2 | 2017-03-10 | 4 months ago | Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files. |
| [go-exiftool](https://github.com/barasher/go-exiftool) | 35 | 2 | 2019-05-12 | 6 months ago | Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...). |
| [checksum](https://github.com/codingsince1985/checksum) | 23 | 1 | 2014-11-05 | 2 months ago | Compute message digest, like MD5 and SHA256, for large files. |
| [go-gtfs](https://github.com/artonge/go-gtfs) | 21 | 2 | 2017-07-09 | 1 month ago | Load gtfs files in go. |
| [flop](https://github.com/homedepot/flop) | 16 | 14 | 2019-03-01 | 1 week ago | File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html). |
| [gut](https://github.com/1set/gut) | 14 | 1 | 2019-10-05 | 2 days ago | Simple and reliable package for file operations like copy/move/diff/list on files, directories and symbolic links. |
| [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) | 13 | 2 | 2018-10-16 | 6 months ago | Copy files for humans. |
| [parquet](https://github.com/parsyl/parquet) | 7 | 4 | 2019-01-29 | 9 months ago | Read and write [parquet](https://parquet.apache.org) files. |
| [baraka](https://github.com/xis/baraka) | 1 | 0 | 2020-07-12 | 1 week ago | A library for handling file uploads, simple, memory-friendly. |

# Tools
        
*Go software and plugins.*

## Code Analysis
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [lint](https://github.com/golang/lint) | 3593 | 112 | 2013-06-02 | 3 weeks ago | Golint is a linter for Go source code. |
| [errcheck](https://github.com/kisielk/errcheck) | 1547 | 26 | 2013-02-24 | 1 week ago | Errcheck is a program for checking for unchecked errors in Go programs. |
| [gcvis](https://github.com/davecheney/gcvis) | 984 | 35 | 2014-07-10 | 1 year ago | Visualise Go program GC trace data in real time. |
| [go-critic](https://github.com/go-critic/go-critic) | 765 | 22 | 2018-05-05 | 6 days ago | source code linter that brings checks that are currently not implemented in other linters. |
| [php-parser](https://github.com/z7zmey/php-parser) | 755 | 27 | 2017-11-07 | 1 day ago | A Parser for PHP written in Go. |
| [goast-viewer](https://github.com/yuroyoro/goast-viewer) | 455 | 16 | 2014-06-30 | 1 year ago | Web based Golang AST visualizer. |
| [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) | 421 | 7 | 2019-04-19 | 5 days ago | An easy way to find outdated dependencies of your Go projects. |
| [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) | 340 | 9 | 2017-04-12 | 2 months ago | go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects. |
| [unconvert](https://github.com/mdempsky/unconvert) | 297 | 8 | 2016-02-19 | 2 months ago | Remove unnecessary type conversions from Go source. |
| [gostatus](https://github.com/shurcooL/gostatus) | 243 | 7 | 2013-11-27 | 1 year ago | Command line tool, shows the status of repositories that contain Go packages. |
| [dupl](https://github.com/mibk/dupl) | 206 | 7 | 2015-05-20 | 1 year ago | Tool for code clone detection. |
| [tickgit](https://github.com/augmentable-dev/tickgit) | 197 | 7 | 2019-10-12 | 3 weeks ago | CLI and go package for surfacing code comment TODOs (in any language) and applying a `git blame`to identify the author. |
| [goplantuml](https://github.com/jfeliu007/goplantuml) | 192 | 7 | 2019-05-26 | 3 weeks ago | Library and CLI that generates text plantump class diagram containing information about structures and interfaces with the relationship among them. |
| [apicompat](https://github.com/bradleyfalzon/apicompat) | 171 | 8 | 2016-07-10 | 3 years ago | Checks recent changes to a Go project for backwards incompatible changes. |
| [checkstyle](https://github.com/qiniu/checkstyle) | 107 | 11 | 2014-01-01 | 7 months ago | checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments. |
| [lint](https://github.com/surullabs/lint) | 64 | 5 | 2016-07-09 | 1 year ago | Run linters as part of go test. |
| [validate](https://github.com/mccoyst/validate) | 61 | 6 | 2013-11-22 | 4 years ago | Automatically validates struct fields with tags. |
| [golines](https://github.com/segmentio/golines) | 46 | 14 | 2019-10-01 | 4 months ago | Formatter that automatically shortens long lines in Go code. |
| [go-outdated](https://github.com/firstrow/go-outdated) | 44 | 1 | 2015-06-29 | 1 year ago | Console application that displays outdated packages. |
| [blanket](https://github.com/verygoodsoftwarenotvirus/blanket) | 14 | 2 | 2017-09-04 | 2 years ago | tarp finds functions and methods without direct unit tests in Go source code. |

## Editor Plugins
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [vim-go](https://github.com/fatih/vim-go) | 12404 | 295 | 2014-03-24 | 2 days ago | Go development plugin for Vim. |
| [vscode-go](https://github.com/microsoft/vscode-go) | 5986 | 229 | 2015-10-14 | 1 month ago | Extension for Visual Studio Code (VS Code) which provides support for the Go language. |
| [gocode](https://github.com/nsf/gocode) | 4866 | 193 | 2010-07-05 | 2 months ago | Autocompletion daemon for the Go programming language. |
| [GoSublime](https://github.com/DisposaBoy/GoSublime) | 3361 | 123 | 2011-08-27 | 6 days ago | Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features. |
| [go-plus](https://github.com/joefitzgerald/go-plus) | 1508 | 46 | 2014-03-13 | 21 hours ago | Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting. |
| [go-mode.el](https://github.com/dominikh/go-mode.el) | 1084 | 52 | 2013-01-30 | 3 weeks ago | Go mode for GNU/Emacs. |
| [Watch](https://github.com/eaburns/Watch) | 174 | 12 | 2013-08-08 | 2 years ago | Runs a command in an acme win on file changes. |
| [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) | 85 | 5 | 2012-11-25 | 4 years ago | Vim plugin to highlight syntax errors on save. |
| [go-language-server](https://github.com/theia-ide/go-language-server) | 32 | 5 | 2017-11-21 | 1 year ago | A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol. |
| [gounit-vim](https://github.com/hexdigest/gounit-vim) | 20 | 2 | 2018-02-21 | 1 year ago | Vim plugin for generating Go tests based on the function's or method's signature. |
| [theia-go-extension](https://github.com/theia-ide/theia-go-extension) | 13 | 4 | 2017-11-30 | 1 year ago | Go language support for the Theia IDE. |

## Go Generate Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gotests](https://github.com/cweill/gotests) | 2779 | 74 | 2016-01-19 | 3 months ago | Generate Go tests from your source code. |
| [genny](https://github.com/cheekybits/genny) | 1247 | 24 | 2014-10-27 | 2 weeks ago | Elegant generics for Go. |
| [re2dfa](https://github.com/opennota/re2dfa) | 182 | 9 | 2015-06-20 | 1 year ago | Transform regular expressions into finite state machines and output Go source code. |
| [gocontracts](https://github.com/Parquery/gocontracts) | 61 | 6 | 2018-08-13 | 1 year ago | brings design-by-contract to Go by synchronizing the code with the documentation. |
| [hasgo](https://github.com/DylanMeeus/hasgo) | 45 | 4 | 2019-05-16 | 1 month ago | Generate Haskell inspired functions for your slices. |
| [gounit](https://github.com/hexdigest/gounit) | 40 | 4 | 2018-02-05 | 1 year ago | Generate Go tests using your own templates. |
| [xgen](https://github.com/xuri/xgen) | 35 | 5 | 2019-06-22 | 1 month ago | XSD (XML Schema Definition) parser and Go/C/Java/Rust/TypeScript code generator. |
| [generic](https://github.com/usk81/generic) | 32 | 3 | 2016-06-15 | 1 year ago | flexible data type for Go. |

## Go Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-swagger](https://github.com/go-swagger/go-swagger) | 5413 | 124 | 2014-11-16 | 3 days ago | Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API. |
| [OctoLinker](https://github.com/OctoLinker/OctoLinker) | 4428 | 90 | 2013-12-27 | 1 day ago | Navigate through go files efficiently with the OctoLinker browser extension for GitHub. |
| [go-callvis](https://github.com/ofabry/go-callvis) | 2671 | 79 | 2016-09-03 | 1 week ago | Visualize call graph of your Go program using dot format. |
| [go-callvis](https://github.com/TrueFurby/go-callvis) | 2408 | 70 | 2016-09-03 | 5 months ago | Visualize call graph of your Go program using dot format. |
| [depth](https://github.com/KyleBanks/depth) | 512 | 11 | 2017-03-04 | 5 months ago | Visualize dependency trees of any package by analyzing imports. |
| [richgo](https://github.com/kyoh86/richgo) | 472 | 5 | 2017-01-04 | 1 month ago | Enrich `go test` outputs with text decorations. |
| [rts](https://github.com/galeone/rts) | 204 | 3 | 2016-04-04 | 4 weeks ago | RTS: response to struct. Generates Go structs from server responses. |
| [godbg](https://github.com/tylerwince/godbg) | 168 | 4 | 2019-01-23 | 1 year ago | Implementation of Rusts `dbg!` macro for quick and easy debugging during development. |
| [typex](https://github.com/dtgorski/typex) | 124 | 2 | 2020-03-24 | 2 weeks ago | Examine Go types and their transitive dependencies, alternatively export results as TypeScript value objects (or types) declaration. |
| [colorgo](https://github.com/songgao/colorgo) | 105 | 6 | 2013-02-14 | 1 week ago | Wrapper around `go` command for colorized `go build` output. |
| [gothanks](https://github.com/psampaz/gothanks) | 93 | 3 | 2019-11-10 | 3 months ago | GoThanks automatically stars your go.mod github dependencies, sending this way some love to their maintainers. |
| [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) | 37 | 2 | 2015-05-22 | 2 years ago | Bash completion for go and wgo. |
| [igo](https://github.com/rocketlaunchr/igo) | 30 | 2 | 2018-11-17 | 3 months ago | Improved Go Syntax (transpiler) |
| [go-james](https://github.com/pieterclaerhout/go-james) | 27 | 2 | 2019-10-14 | 1 week ago | Go project skeleton creator, builds and tests your projects without the manual setup. |
| [generator-go-lang](https://github.com/axelspringer/generator-go-lang) | 21 | 13 | 2017-09-13 | 3 months ago | A [Yeoman](http://yeoman.io) generator to get new Go projects started. |

## Software Packages
        
*Software written in Go.*

### DevOps Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [kubernetes](https://github.com/kubernetes/kubernetes) | 68530 | 3201 | 2014-06-06 | 54 minutes ago | Container Cluster Manager from Google. |
| [moby](https://github.com/moby/moby) | 57691 | 3141 | 2013-01-18 | 2 hours ago | Collaborative project for the container ecosystem to assemble container-based systems. |
| [traefik](https://github.com/containous/traefik) | 29833 | 705 | 2015-09-13 | 32 minutes ago | Reverse proxy and load balancer with support for multiple backends. |
| [gitea](https://github.com/go-gitea/gitea) | 20646 | 459 | 2016-11-01 | 7 minutes ago | Fork of Gogs, entirely community driven. |
| [vegeta](https://github.com/tsenart/vegeta) | 15058 | 305 | 2013-08-13 | 3 days ago | HTTP load testing tool and library. It's over 9000! |
| [packer](https://github.com/hashicorp/packer) | 10453 | 406 | 2013-03-23 | 1 hour ago | Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. |
| [hey](https://github.com/rakyll/hey) | 8830 | 160 | 2016-09-02 | 1 week ago | Hey is a tiny program that sends some load to a web application. |
| [webhook](https://github.com/adnanh/webhook) | 5505 | 138 | 2015-01-12 | 22 hours ago | Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server. |
| [gvm](https://github.com/moovweb/gvm) | 5443 | 156 | 2011-12-03 | 2 months ago | GVM provides an interface to manage Go versions. |
| [gaia](https://github.com/gaia-pipeline/gaia) | 4150 | 107 | 2017-12-28 | 1 hour ago | Build powerful pipelines in any programming language. |
| [gox](https://github.com/mitchellh/gox) | 3798 | 74 | 2013-11-17 | 5 months ago | Dead simple, no frills Go cross compile tool. |
| [bosun](https://github.com/bosun-monitor/bosun) | 3025 | 155 | 2013-11-15 | 2 weeks ago | Time Series Alerting Framework. |
| [bombardier](https://github.com/codesenberg/bombardier) | 2169 | 54 | 2016-05-29 | 2 weeks ago | Fast cross-platform HTTP benchmarking tool. |
| [fac](https://github.com/mkchoi212/fac) | 1664 | 28 | 2017-12-29 | 9 months ago | Command-line user interface to fix git merge conflicts. |
| [goxc](https://github.com/laher/goxc) | 1648 | 49 | 2013-02-11 | 10 months ago | build tool for Go, with a focus on cross-compiling and packaging. |
| [pomerium](https://github.com/pomerium/pomerium) | 1598 | 20 | 2019-01-01 | 2 hours ago | Pomerium is an identity-aware access proxy. |
| [script](https://github.com/bitfield/script) | 1510 | 26 | 2019-04-20 | 3 weeks ago | Making it easy to write shell-like scripts in Go for DevOps and system administration tasks. |
| [kala](https://github.com/ajvb/kala) | 1489 | 65 | 2015-03-19 | 2 weeks ago | Simplistic, modern, and performant job scheduler. |
| [statusok](https://github.com/sanathp/statusok) | 1362 | 49 | 2015-08-26 | 1 week ago | Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected. |
| [s3gof3r](https://github.com/rlmcpherson/s3gof3r) | 1058 | 34 | 2013-08-02 | 5 months ago | Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3. |
| [go-selfupdate](https://github.com/sanbornm/go-selfupdate) | 747 | 28 | 2013-11-13 | 4 months ago | Enable your Go applications to self update. |
| [skm](https://github.com/TimothyYe/skm) | 638 | 19 | 2017-10-11 | 5 days ago | SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily! |
| [scaleway-cli](https://github.com/scaleway/scaleway-cli) | 607 | 31 | 2015-03-20 | 20 minutes ago | Manage BareMetal Servers from Command Line (as easily as with Docker). |
| [aurora](https://github.com/xuri/aurora) | 464 | 30 | 2016-10-09 | 2 months ago | Cross-platform web-based Beanstalkd queue server console. |
| [govvv](https://github.com/ahmetb/govvv) | 455 | 10 | 2016-08-02 | 5 months ago | “go build” wrapper to easily add version information into Go binaries. |
| [cassowary](https://github.com/rogerwelin/cassowary) | 437 | 5 | 2019-08-25 | 3 weeks ago | Modern cross-platform HTTP load-testing tool written in Go. |
| [s5cmd](https://github.com/peak/s5cmd) | 340 | 20 | 2016-11-16 | 5 hours ago | Blazing fast S3 and local filesystem execution tool. |
| [gonative](https://github.com/inconshreveable/gonative) | 321 | 7 | 2014-05-01 | 4 years ago | Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages. |
| [mora](https://github.com/emicklei/mora) | 273 | 23 | 2013-07-12 | 3 years ago | REST server for accessing MongoDB documents and meta data. |
| [trubka](https://github.com/xitonix/trubka) | 269 | 13 | 2019-07-05 | 2 days ago | A CLI tool to manage and troubleshoot Apache Kafka clusters with the ability of generically publishing/consuming protocol buffer and plain text events to/from Kafka. |
| [lstags](https://github.com/ivanilves/lstags) | 261 | 14 | 2017-08-15 | 1 month ago | Tool and API to sync Docker images across different registries. |
| [pewpew](https://github.com/bengadbois/pewpew) | 231 | 7 | 2016-10-12 | 1 month ago | Flexible HTTP command line stress tester. |
| [dogo](https://github.com/liudng/dogo) | 229 | 18 | 2014-11-19 | 1 year ago | Monitoring changes in the source file and automatically compile and run (restart). |
| [godbg](https://github.com/sirnewton01/godbg) | 221 | 18 | 2013-08-09 | 2 years ago | Web-based gdb front-end application. |
| [manssh](https://github.com/xwjdsh/manssh) | 220 | 4 | 2017-10-08 | 2 years ago | manssh is a command line tool for managing your ssh alias config easily. |
| [utask](https://github.com/ovh/utask) | 205 | 23 | 2019-11-05 | 1 day ago | Automation engine that models and executes business processes declared in yaml. |
| [jenkins-cli](https://github.com/jenkins-zh/jenkins-cli) | 187 | 10 | 2019-06-21 | 4 days ago | Jenkins CLI allows you manage your Jenkins as an easy way. |
| [blast](https://github.com/dave/blast) | 185 | 4 | 2017-10-21 | 2 years ago | A simple tool for API load testing and batch jobs. |
| [gobrew](https://github.com/cryptojuice/gobrew) | 177 | 5 | 2013-11-13 | 2 months ago | gobrew lets you easily switch between multiple versions of go. |
| [ostent](https://github.com/ostrost/ostent) | 166 | 6 | 2014-03-31 | 2 years ago | collects and displays system metrics and optionally relays to Graphite and/or InfluxDB. |
| [grapes](https://github.com/yaronsumel/grapes) | 151 | 6 | 2016-09-01 | 10 months ago | Lightweight tool designed to distribute commands over ssh with ease. |
| [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) | 148 | 5 | 2017-03-03 | 2 months ago | Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`. |
| [kcli](https://github.com/cswank/kcli) | 121 | 6 | 2017-03-25 | 6 months ago | Command line tool for inspecting kafka topics/partitions/messages. |
| [terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi) | 118 | 10 | 2017-10-17 | 1 week ago | Terraform provider plugin that dynamically configures itself at runtime based on an OpenAPI document (formerly known as swagger file) containing the definitions of the APIs exposed. |
| [winrm-cli](https://github.com/masterzen/winrm-cli) | 95 | 5 | 2016-05-23 | 2 months ago | Cli tool to remotely execute commands on Windows machines. |
| [dockerfile-generator](https://github.com/ozankasikci/dockerfile-generator) | 82 | 5 | 2019-08-14 | 6 months ago | A go library and an executable that produces valid Dockerfiles using various input channels. |
| [go-furnace](https://github.com/go-furnace/go-furnace) | 77 | 1 | 2016-10-09 | 10 months ago | Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean. |
| [drone-scp](https://github.com/appleboy/drone-scp) | 75 | 3 | 2016-10-16 | 1 month ago | Copy files and artifacts via SSH using a binary, docker or Drone CI. |
| [dropship](https://github.com/ChrisMcKenzie/dropship) | 52 | 3 | 2015-09-03 | 2 years ago | Tool for deploying code via cdn. |
| [rodent](https://github.com/alouche/rodent) | 31 | 2 | 2014-06-01 | 3 years ago | Rodent helps you manage Go versions, projects and track dependencies. |
| [drone-jenkins](https://github.com/appleboy/drone-jenkins) | 29 | 2 | 2016-10-15 | 1 week ago | Trigger downstream Jenkins jobs using a binary, docker or Drone CI. |
| [awsenv](https://github.com/soniah/awsenv) | 24 | 2 | 2015-08-05 | 2 years ago | Small binary that loads Amazon (AWS) environment variables for a profile. |
| [lwc](https://github.com/timdp/lwc) | 22 | 3 | 2018-04-22 | 2 months ago | A live-updating version of the UNIX wc command. |
| [s3-proxy](https://github.com/oxyno-zeta/s3-proxy) | 15 | 2 | 2019-09-22 | 3 days ago | S3 Proxy with GET, PUT and DELETE methods and authentication (OpenID Connect and Basic Auth). |
| [depcharge](https://github.com/centerorbit/depcharge) | 13 | 3 | 2018-07-25 | 5 months ago | Helps orchestrating the execution of commands across the many dependencies in larger projects. |
| [sg](https://github.com/ChristopherRabotin/sg) | 5 | 1 | 2015-08-19 | 3 years ago | Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response. |
| [aptly-fork](https://github.com/smira/aptly-fork) | 2 | 0 | 2019-07-04 | 10 months ago | aptly is a Debian repository management tool. |

### Other Software
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [goreplay](https://github.com/buger/goreplay) | 13077 | 465 | 2013-05-30 | 1 day ago | Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time. |
| [restic](https://github.com/restic/restic) | 10570 | 234 | 2014-04-27 | 20 hours ago | De-duplicating backup program. |
| [seaweedfs](https://github.com/chrislusf/seaweedfs) | 9981 | 519 | 2014-07-14 | 33 minutes ago | Fast, Simple and Scalable Distributed File System with O(1) disk seek. |
| [confd](https://github.com/kelseyhightower/confd) | 7103 | 261 | 2013-10-01 | 2 weeks ago | Manage local application configuration files using templates and data from etcd or consul. |
| [comcast](https://github.com/tylertreat/comcast) | 6686 | 152 | 2014-11-12 | 2 months ago | Simulate bad network connections. |
| [liteide](https://github.com/visualfc/liteide) | 6098 | 376 | 2012-11-19 | 3 days ago | LiteIDE is a simple, open source, cross-platform Go IDE. |
| [drive](https://github.com/odeke-em/drive) | 5531 | 203 | 2014-11-03 | 5 months ago | Google Drive client for the commandline. |
| [toxiproxy](https://github.com/Shopify/toxiproxy) | 4669 | 291 | 2014-09-04 | 2 months ago | Proxy to simulate network and system conditions for automated tests. |
| [nes](https://github.com/fogleman/nes) | 4524 | 151 | 2015-03-02 | 3 months ago | Nintendo Entertainment System (NES) emulator written in Go. |
| [duplicacy](https://github.com/gilbertchen/duplicacy) | 3370 | 95 | 2016-02-23 | 2 weeks ago | A cross-platform network and cloud backup tool based on the idea of lock-free deduplication. |
| [croc](https://github.com/schollz/croc) | 3305 | 59 | 2017-10-17 | 1 week ago | Easily and securely send files or folders from one computer to another. |
| [mylg](https://github.com/mehrdadrad/mylg) | 2337 | 109 | 2016-06-21 | 5 months ago | Command Line Network Diagnostic tool written in Go. |
| [goboy](https://github.com/Humpheh/goboy) | 2232 | 40 | 2017-08-20 | 1 month ago | Nintendo Game Boy Color emulator written in Go. |
| [sup](https://github.com/pressly/sup) | 2182 | 72 | 2015-02-23 | 5 months ago | Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers. |
| [lgo](https://github.com/yunabe/lgo) | 2048 | 49 | 2017-10-05 | 1 year ago | Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility. |
| [circuit](https://github.com/gocircuit/circuit) | 1839 | 142 | 2014-04-10 | 2 months ago | Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications. |
| [snap](https://github.com/intelsdi-x/snap) | 1798 | 146 | 2014-08-13 | 1 year ago | Powerful telemetry framework. |
| [scc](https://github.com/boyter/scc) | 1624 | 18 | 2018-03-01 | 3 weeks ago | Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates. |
| [borg](https://github.com/ok-borg/borg) | 1473 | 42 | 2016-09-10 | 2 years ago | Terminal based search engine for bash snippets. |
| [community](https://github.com/documize/community) | 1125 | 45 | 2016-04-29 | 2 months ago | Modern wiki software that integrates data from SaaS tools. |
| [Go-Package-Store](https://github.com/shurcooL/Go-Package-Store) | 890 | 19 | 2014-01-24 | 4 months ago | App that displays updates for the Go packages in your GOPATH. |
| [peg](https://github.com/pointlander/peg) | 707 | 31 | 2010-04-25 | 12 hours ago | Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator. |
| [vflow](https://github.com/VerizonDigital/vflow) | 683 | 83 | 2017-02-24 | 1 month ago | High-performance, scalable and reliable IPFIX, sFlow and Netflow collector. |
| [leaps](https://github.com/Jeffail/leaps) | 681 | 27 | 2014-06-19 | 1 year ago | Pair programming service using Operational Transforms. |
| [shell2http](https://github.com/msoap/shell2http) | 599 | 24 | 2015-03-11 | 1 month ago | Executing shell commands via http server (for prototyping or remote control). |
| [gfile](https://github.com/Antonito/gfile) | 564 | 12 | 2019-03-08 | 1 year ago | Securely transfer files between two computers, without any third party, over WebRTC. |
| [mockingjay-server](https://github.com/quii/mockingjay-server) | 457 | 10 | 2015-04-04 | 4 months ago | Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests. |
| [gocc](https://github.com/goccmack/gocc) | 415 | 23 | 2015-06-05 | 1 month ago | Gocc is a compiler kit for Go written in Go. |
| [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) | 407 | 18 | 2015-10-08 | 4 months ago | Video streaming torrent client. |
| [ipe](https://github.com/dimiro1/ipe) | 304 | 18 | 2015-01-13 | 4 months ago | Open source Pusher server implementation compatible with Pusher client libraries written in GO. |
| [wellington](https://github.com/wellington/wellington) | 295 | 13 | 2014-12-08 | 4 months ago | Sass project management tool, extends the language with sprite functions (like Compass). |
| [IDE](https://github.com/thestrukture/IDE) | 258 | 14 | 2017-09-09 | 11 months ago | Browser accessible IDE. Designed for Go with Go. |
| [cherry](https://github.com/rafael-santiago/cherry) | 222 | 12 | 2015-10-24 | 3 years ago | Tiny webchat server in Go. |
| [orange-cat](https://github.com/utatti/orange-cat) | 185 | 5 | 2014-11-01 | 1 year ago | Markdown previewer written in Go. |
| [joincap](https://github.com/assafmo/joincap) | 143 | 7 | 2018-05-31 | 3 months ago | Command-line utility for merging multiple pcap files together. |
| [orbit](https://github.com/gulien/orbit) | 138 | 8 | 2017-05-13 | 8 months ago | A simple tool for running commands and generating files from templates. |
| [boxed](https://github.com/tejo/boxed) | 71 | 2 | 2015-04-18 | 1 year ago | Dropbox based blog engine. |
| [dp](https://github.com/scryinfo/dp) | 60 | 8 | 2018-12-12 | 1 week ago | Through SDK for data exchange with blockchain, developers can get easy access to DAPP development. |
| [naclpipe](https://github.com/unix4fun/naclpipe) | 21 | 5 | 2015-05-05 | 1 year ago | Simple NaCL EC25519 based crypto pipe tool written in Go. |
| [term-quiz](https://github.com/crazcalm/term-quiz) | 17 | 1 | 2017-12-26 | 1 year ago | Quizzes for your terminal. |
| [snitch](https://github.com/lucasgomide/snitch) | 15 | 1 | 2017-04-06 | 2 years ago | Simple way to notify your team and many tools when someone has deployed any application via Tsuru. |
| [GoDocTooltip](https://github.com/diankong/GoDocTooltip) | 11 | 3 | 2016-01-21 | 3 months ago | Chrome extension for Go Doc sites, which shows function description as tooltip at function list. |

# Resources
        
*Where to discover new Go libraries.*

## Benchmarks
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) | 1398 | 59 | 2013-12-16 | 1 day ago | Go HTTP request router benchmark and comparison. |
| [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) | 1259 | 84 | 2016-04-06 | 15 hours ago | Go web framework benchmark. |
| [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) | 1016 | 38 | 2013-01-18 | 7 months ago | Benchmarks of Go serialization methods. |
| [skynet](https://github.com/atemerev/skynet) | 962 | 48 | 2016-02-14 | 1 year ago | Skynet 1M threads microbenchmark. |
| [speedtest-resize](https://github.com/fawick/speedtest-resize) | 193 | 7 | 2013-09-16 | 9 months ago | Compare various Image resize algorithms for the Go language. |
| [go-benchmarks](https://github.com/tylertreat/go-benchmarks) | 133 | 10 | 2016-02-25 | 4 years ago | Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches. |
| [gospeed](https://github.com/feyeleanor/gospeed) | 102 | 7 | 2011-05-23 | 1 year ago | Go micro-benchmarks for calculating the speed of language constructs. |
| [autobench](https://github.com/davecheney/autobench) | 90 | 8 | 2013-03-28 | 6 years ago | Framework to compare the performance between different Go versions. |
| [gocostmodel](https://github.com/mna/gocostmodel) | 55 | 5 | 2014-12-19 | 5 years ago | Benchmarks of common basic operations for the Go language. |
| [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) | 52 | 5 | 2014-09-24 | 2 years ago | Collection of benchmarks for popular Go database/SQL utilities. |
| [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) | 21 | 1 | 2017-01-24 | 3 years ago | Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results. |
| [kvbench](https://github.com/jimrobinson/kvbench) | 17 | 1 | 2014-04-15 | 10 months ago | Key/Value database benchmark. |

## Conferences
        

## E-Books
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [GoBooks](https://github.com/dariubs/GoBooks) | 8141 | 557 | 2015-05-05 | 1 week ago | A curated list of Go books. |
| [The-Golang-Standard-Library-by-Example](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example) | 6956 | 569 | 2013-04-14 | 6 days ago | Golang标准库。对于程序员而言，标准库与语言本身同样重要，它好比一个百宝箱，能为各种常见的任务提供完美的解决方案。以示例驱动的方式讲解Golang的标准库。 |
| [gosuccinctly](https://github.com/thedevsir/gosuccinctly) | 14 | 3 | 2018-09-02 | 1 year ago | in Persian. |

## Gophers
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gophers](https://github.com/ashleymcnamara/gophers) | 2220 | 95 | 2017-02-15 | 1 year ago | Gopher artworks by Ashley McNamara. |
| [gophers](https://github.com/egonelbre/gophers) | 2044 | 47 | 2015-06-03 | 1 month ago | Free gophers. |
| [free-gophers-pack](https://github.com/MariaLetta/free-gophers-pack) | 1955 | 52 | 2019-04-02 | 4 weeks ago | Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster. |
| [gophericons](https://github.com/shalakhin/gophericons) | 566 | 20 | 2015-08-22 | 2 years ago | 34 gopher images for Go developers community |
| [gopher-stickers](https://github.com/tenntenn/gopher-stickers) | 481 | 15 | 2014-11-09 | 7 months ago | gopher stickers |
| [gopherize.me](https://github.com/matryer/gopherize.me) | 434 | 9 | 2017-01-25 | 1 year ago | Gopherize yourself. |
| [gopher-vector](https://github.com/golang-samples/gopher-vector) | 364 | 14 | 2013-03-31 | 4 years ago | Vector data of gopher |
| [gopher-logos](https://github.com/GolangUA/gopher-logos) | 74 | 7 | 2017-07-27 | 2 years ago | adorable gopher logos. |
| [go-gopher](https://github.com/sillecelik/go-gopher) | 57 | 0 | 2018-03-28 | 3 months ago | Gopher amigurumi toy pattern. |
| [gophers](https://github.com/rogeralsing/gophers) | 53 | 2 | 2017-01-28 | 3 years ago | random gopher graphics. |
| [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) | 35 | 1 | 2014-09-03 | 2 years ago | Go gopher Vector Data [.ai, .svg]. |

## Meetups
        
*Add the group of your city/country here (send **PR**)*

## Twitter
        

## Websites
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) | 26364 | 1746 | 2014-07-08 | 3 weeks ago | List of other amazingly awesome lists. |
| [awesome-remote-job](https://github.com/lukasz-madon/awesome-remote-job) | 18943 | 1021 | 2015-01-02 | 1 week ago | Curated list of awesome remote jobs. A lot of them are looking for Go hackers. |
| [golang-graphics](https://github.com/mholt/golang-graphics) | 144 | 9 | 2014-03-24 | 4 years ago | Collection of Go images, graphics, and art. |
| [gocryforhelp](https://github.com/ninedraft/gocryforhelp) | 38 | 12 | 2016-05-09 | 2 years ago | Collection of Go projects that needs help. Good place to start your open-source way in Go. |

### Tutorials
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang) | 35280 | 2475 | 2012-08-02 | 15 hours ago | Golang ebook intro how to build a web app with golang. |
| [go-patterns](https://github.com/tmrts/go-patterns) | 13155 | 588 | 2015-12-14 | 2 weeks ago | Curated list of Go design patterns, recipes and idioms. |
| [learn-go-with-tests](https://github.com/quii/learn-go-with-tests) | 11501 | 273 | 2018-03-02 | 1 week ago | Learn Go with test-driven development. |
| [golang-cheat-sheet](https://github.com/a8m/golang-cheat-sheet) | 4664 | 184 | 2014-02-13 | 1 month ago | Go's reference card. |
| [golang-for-nodejs-developers](https://github.com/miguelmota/golang-for-nodejs-developers) | 1293 | 34 | 2019-01-03 | 2 months ago | Examples of Golang compared to Node.js for learning. |
| [working-with-go](https://github.com/mkaz/working-with-go) | 1162 | 50 | 2014-05-04 | 5 months ago | Intro to go for experienced programmers. |
| [ethereum-development-with-go-book](https://github.com/miguelmota/ethereum-development-with-go-book) | 630 | 39 | 2018-05-16 | 2 weeks ago | A little e-book on Ethereum Development with Go. |
| [goapp](https://github.com/bnkamalesh/goapp) | 146 | 10 | 2020-07-04 | 10 hours ago | An opinionated guideline to structure & develop a Go web application/service. |

## Style Guides
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-styleguide](https://github.com/bahlo/go-styleguide) | 936 | 28 | 2017-07-29 | 4 months ago | Opinionated Styleguide for the Go language |

> 该项目源码[Awesome Go Analysis](https://github.com/plholx/awesome-go-analysis)
> 更专业的go开源项目分析请移步 [Awesome Go](https://go.libhunt.com/)
