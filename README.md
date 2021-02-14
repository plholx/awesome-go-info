# Awesome Go Info

go语言开源项目列表，项目分类及GitHub上的开源项目数据完全来自于[awesome-go](https://github.com/avelino/awesome-go) 的[README.md](https://github.com/avelino/awesome-go/blob/master/README.md)文件，通过调用GitHub的API获取仓库信息，展示项目的star数、watch数等，方便查看go语言开源项目的一些相关信息。

_该文件仅包含[awesome-go](https://github.com/avelino/awesome-go) [README.md](https://github.com/avelino/awesome-go/blob/master/README.md)文件中列出的在GitHub上开源的优秀项目，不罗列其它golang相关的网站_
_该文件中的GitHub仓库信息数据会在每天凌晨1点左右更新,当前数据更新于2021-02-15 01:10:25_

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
    - [Build Automation](#build-automation)
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
| [oto](https://github.com/hajimehoshi/oto) | 769 | 10 | 2017/05/04 | 6 days ago | A low-level library to play sound on multiple platforms. |
| [portaudio](https://github.com/gordonklaus/portaudio) | 431 | 14 | 2015/09/16 | 5 months ago | Go bindings for the PortAudio audio I/O library. |
| [music-theory](https://github.com/go-music-theory/music-theory) | 331 | 16 | 2016/03/17 | 6 months ago | Music theory models in Go. |
| [waveform](https://github.com/mdlayher/waveform) | 324 | 12 | 2014/09/13 | 10 months ago | Go package capable of generating waveform images from audio streams. |
| [portmidi](https://github.com/rakyll/portmidi) | 252 | 11 | 2013/11/10 | 3 months ago | Go bindings for PortMidi. |
| [id3v2](https://github.com/bogem/id3v2) | 177 | 6 | 2016/05/15 | 3 weeks ago | ID3 decoding and encoding library for Go. |
| [flac](https://github.com/mewkiz/flac) | 144 | 10 | 2012/11/01 | 2 weeks ago | Native Go FLAC encoder/decoder with support for FLAC streams. |
| [mix](https://github.com/go-mix/mix) | 136 | 3 | 2016/01/03 | 9 months ago | Sequence-based Go-native audio mixer for music apps. |
| [malgo](https://github.com/gen2brain/malgo) | 134 | 5 | 2017/11/09 | 1 month ago | Mini audio library. |
| [go-sox](https://github.com/krig/go-sox) | 118 | 8 | 2013/10/08 | 2 years ago | libsox bindings for go. |
| [GoAudio](https://github.com/DylanMeeus/GoAudio) | 116 | 6 | 2020/07/05 | 2 weeks ago | Native Go Audio Processing Library. |
| [mp3](https://github.com/tcolgate/mp3) | 112 | 2 | 2015/02/26 | 3 years ago | Native Go MP3 decoder. |
| [gaad](https://github.com/Comcast/gaad) | 78 | 10 | 2016/07/11 | 1 year ago | Native Go AAC bitstream parser. |
| [go-taglib](https://github.com/wtolson/go-taglib) | 74 | 6 | 2012/11/20 | 2 years ago | Go bindings for taglib. |
| [minimp3](https://github.com/tosone/minimp3) | 52 | 3 | 2018/01/26 | 5 months ago | Lightweight MP3 decoder library. |
| [EasyMIDI](https://github.com/algoGuy/EasyMIDI) | 35 | 3 | 2018/02/19 | 2 years ago | EasyMidi is a simple and reliable library for working with standard midi file (SMF). |
| [go_mediainfo](https://github.com/zhulik/go_mediainfo) | 32 | 1 | 2015/12/13 | 5 years ago | libmediainfo bindings for go. |
| [vorbis](https://github.com/mccoyst/vorbis) | 28 | 3 | 2013/07/12 | 1 year ago | "Native" Go Vorbis decoder (uses CGO, but has no dependencies). |
| [gosamplerate](https://github.com/dh1tw/gosamplerate) | 10 | 1 | 2016/11/20 | 7 months ago | libsamplerate bindings for go. |

## Authentication and OAuth
        
*Libraries for implementing authentications schemes.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [jwt-go](https://github.com/dgrijalva/jwt-go) | 9044 | 152 | 2012/04/18 | 1 week ago | Golang implementation of JSON Web Tokens (JWT). |
| [casbin](https://github.com/casbin/casbin) | 8627 | 210 | 2017/04/08 | 2 hours ago | Authorization library that supports access control models like ACL, RBAC, ABAC. |
| [oauth2](https://github.com/golang/oauth2) | 3502 | 107 | 2014/04/14 | 1 day ago | Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support. |
| [goth](https://github.com/markbates/goth) | 3067 | 64 | 2014/10/14 | 2 weeks ago | provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box. |
| [authboss](https://github.com/volatiletech/authboss) | 2555 | 52 | 2015/01/03 | 2 months ago | Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time. |
| [loginsrv](https://github.com/tarent/loginsrv) | 1743 | 54 | 2016/11/11 | 3 days ago | JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam. |
| [go-jose](https://github.com/square/go-jose) | 1704 | 63 | 2014/11/14 | 2 months ago | Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs. |
| [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) | 1698 | 84 | 2015/11/01 | 1 month ago | Standalone, specification-compliant,  OAuth2 server written in Golang. |
| [osin](https://github.com/openshift/osin) | 1646 | 74 | 2013/09/10 | 2 days ago | Golang OAuth2 server library. |
| [gologin](https://github.com/dghubble/gologin) | 1328 | 28 | 2015/06/23 | 4 months ago | chainable handlers for login with OAuth1 and OAuth2 authentication providers. |
| [gorbac](https://github.com/mikespook/gorbac) | 1134 | 61 | 2013/12/26 | 1 year ago | provides a lightweight role-based access control (RBAC) implementation in Golang. |
| [scs](https://github.com/alexedwards/scs) | 835 | 21 | 2016/08/08 | 2 weeks ago | Session Manager for HTTP servers. |
| [paseto](https://github.com/o1egl/paseto) | 435 | 21 | 2018/01/23 | 8 months ago | Golang implementation of Platform-Agnostic Security Tokens (PASETO). |
| [permissions2](https://github.com/xyproto/permissions2) | 419 | 13 | 2014/11/19 | 5 months ago | Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt. |
| [jwt](https://github.com/pascaldekloe/jwt) | 235 | 12 | 2018/03/21 | 6 months ago | Lightweight JSON Web Token (JWT) library. |
| [jeff](https://github.com/abraithwaite/jeff) | 224 | 4 | 2018/08/02 | 6 months ago | Simple, flexible, secure and idiomatic web session management with pluggable backends. |
| [jwt](https://github.com/cristalhq/jwt) | 206 | 10 | 2019/07/20 | 1 week ago | Safe, simple and fast JSON Web Tokens for Go. |
| [jwt-auth](https://github.com/adam-hanna/jwt-auth) | 205 | 13 | 2016/07/05 | 3 months ago | JWT middleware for Golang http servers with many configuration options. |
| [httpauth](https://github.com/goji/httpauth) | 199 | 6 | 2014/05/26 | 6 months ago | HTTP Authentication middleware. |
| [go-guardian](https://github.com/shaj13/go-guardian) | 177 | 4 | 2020/05/14 | 1 hour ago | Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication that supports LDAP, Basic, Bearer token and Certificate based authentication. |
| [branca](https://github.com/hako/branca) | 146 | 7 | 2018/01/09 | 6 months ago | Golang implementation of Branca Tokens. |
| [session](https://github.com/icza/session) | 105 | 6 | 2016/02/08 | 1 year ago | Go session management for web servers (including support for Google App Engine - GAE). |
| [sessionup](https://github.com/swithek/sessionup) | 104 | 6 | 2019/07/23 | 2 months ago | Simple, yet effective HTTP session management and identification package. |
| [jwt](https://github.com/robbert229/jwt) | 90 | 8 | 2016/06/05 | 2 months ago | Clean and easy to use implementation of JSON Web Tokens (JWT). |
| [sjwt](https://github.com/brianvoe/sjwt) | 83 | 1 | 2019/06/20 | 1 year ago | Simple jwt generator and parser. |
| [rbac](https://github.com/zpatrick/rbac) | 73 | 3 | 2018/08/02 | 2 years ago | Minimalistic RBAC package for Go applications. |
| [sessions](https://github.com/adam-hanna/sessions) | 56 | 3 | 2017/04/29 | 10 months ago | Dead simple, highly performant, highly customizable sessions service for go http servers. |
| [securecookie](https://github.com/chmike/securecookie) | 44 | 6 | 2017/09/03 | 6 months ago | Efficient secure cookie encoding/decoding. |
| [otpgo](https://github.com/jltorresm/otpgo) | 15 | 3 | 2020/08/19 | 5 months ago | Time-Based One-Time Password (TOTP) and HMAC-Based One-Time Password (HOTP) library for Go. |
| [scope](https://github.com/SonicRoshan/scope) | 10 | 1 | 2019/09/23 | 1 year ago | Easily Manage OAuth2 Scopes In Go. |
| [go-email-normalizer](https://github.com/dimuska139/go-email-normalizer) | 10 | 1 | 2020/08/21 | 2 months ago | Golang library for providing a canonical representation of email address. |
| [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) | 9 | 2 | 2017/10/20 | 2 years ago | Go session management using the SessionGate Redis module. |
| [signedvalue](https://github.com/sashka/signedvalue) | 8 | 0 | 2018/01/06 | 1 year ago | Signed and timestamped strings compatible with [Tornado's](https://github.com/tornadoweb/tornado) `create_signed_value`, `decode_signed_value`, and therefore `set_secure_cookie` and `get_secure_cookie`. |
| [cookiestxt](https://github.com/mengzhuo/cookiestxt) | 6 | 1 | 2017/10/09 | 3 years ago | provides parser of cookies.txt file format. |

## Bot Building
        
*Libraries for building and working with bots.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [olivia](https://github.com/olivia-ai/olivia) | 2835 | 73 | 2018/06/05 | 21 hours ago | A chatbot built with an artificial neural network. |
| [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) | 2658 | 75 | 2015/06/25 | 20 hours ago | Simple and clean Telegram bot client. |
| [telebot](https://github.com/tucnak/telebot) | 1664 | 46 | 2015/06/25 | 2 days ago | Telegram bot framework written in Go. |
| [bot](https://github.com/go-chat-bot/bot) | 652 | 46 | 2015/09/22 | 4 months ago | IRC, Slack & Telegram bot written in Go. |
| [kelp](https://github.com/stellar/kelp) | 534 | 45 | 2018/08/08 | 6 hours ago | official trading and market-making bot for the [Stellar](https://www.stellar.org/) DEX. Works out-of-the-box, written in Golang, compatible with centralized exchanges and custom trading strategies. |
| [slacker](https://github.com/shomali11/slacker) | 489 | 17 | 2017/05/20 | 3 weeks ago | Easy to use framework to create Slack bots. |
| [golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot) | 428 | 29 | 2017/05/14 | 2 weeks ago | A golang implementation of a console-based trading bot for cryptocurrency exchanges. |
| [tbot](https://github.com/yanzay/tbot) | 297 | 11 | 2015/09/11 | 5 months ago | Telegram bot server with API similar to net/http. |
| [go-sarah](https://github.com/oklahomer/go-sarah) | 188 | 8 | 2016/11/06 | 4 months ago | Framework to build bot for desired chat services including LINE, Slack, Gitter and more. |
| [tenyks](https://github.com/kyleterry/tenyks) | 170 | 15 | 2012/08/26 | 1 year ago | Service oriented IRC bot using Redis and JSON for messaging. |
| [go-twitch-irc](https://github.com/gempir/go-twitch-irc) | 144 | 10 | 2017/03/23 | 1 month ago | Libary to write bots for twitch. |
| [hanu](https://github.com/sbstjn/hanu) | 127 | 5 | 2016/09/16 | 8 months ago | Framework for writing Slack bots. |
| [go-tgbot](https://github.com/olebedev/go-tgbot) | 102 | 9 | 2016/12/11 | 2 years ago | Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware. |
| [margelet](https://github.com/zhulik/margelet) | 64 | 4 | 2015/11/21 | 4 years ago | Framework for building Telegram bots. |
| [slackscot](https://github.com/alexandre-normand/slackscot) | 39 | 1 | 2015/10/22 | 1 month ago | Another framework for building Slack bots. |
| [ephemeral-roles](https://github.com/ewohltman/ephemeral-roles) | 38 | 5 | 2017/12/19 | 6 days ago | A Discord bot for managing ephemeral roles based upon voice channel member presence. |
| [govkbot](https://github.com/nikepan/govkbot) | 36 | 3 | 2016/07/11 | 11 months ago | Simple Go [VK](https://vk.com) bot library. |
| [micha](https://github.com/onrik/micha) | 18 | 4 | 2016/04/14 | 5 months ago | Go Library for Telegram bot api. |

## Command Line
        
*Libraries for building Console Applications and Console User Interfaces.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [cobra](https://github.com/spf13/cobra) | 20418 | 335 | 2013/09/03 | 2 hours ago | Commander for modern Go CLI interactions. |
| [cli](https://github.com/urfave/cli) | 15283 | 299 | 2013/07/13 | 1 day ago | Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli). |
| [termui](https://github.com/gizak/termui) | 10606 | 293 | 2015/02/03 | 2 weeks ago | Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib). |
| [gocui](https://github.com/jroimartin/gocui) | 6937 | 126 | 2014/01/04 | 1 month ago | Minimalist Go library aimed at creating Console User Interfaces. |
| [termbox-go](https://github.com/nsf/termbox-go) | 3959 | 101 | 2012/01/12 | 1 month ago | Termbox is a library for creating cross-platform text-based interfaces. |
| [go-prompt](https://github.com/c-bata/go-prompt) | 3703 | 48 | 2017/08/14 | 1 month ago | Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit). |
| [kingpin](https://github.com/alecthomas/kingpin) | 3000 | 56 | 2014/05/14 | 6 months ago | Command line and flag parser supporting sub commands. |
| [dnote](https://github.com/dnote/dnote) | 2008 | 32 | 2017/03/30 | 6 days ago | A simple command line notebook with multi-device sync. |
| [go-flags](https://github.com/jessevdk/go-flags) | 1865 | 28 | 2012/08/31 | 1 month ago | go command line option parser. |
| [uiprogress](https://github.com/gosuri/uiprogress) | 1757 | 36 | 2015/11/17 | 3 months ago | Flexible library to render progress bars in terminal applications. |
| [asciigraph](https://github.com/guptarohit/asciigraph) | 1573 | 31 | 2018/06/17 | 1 month ago | Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. |
| [readline](https://github.com/chzyer/readline) | 1555 | 40 | 2015/09/20 | 1 year ago | Pure golang implementation that provides most features in GNU-Readline under MIT license. |
| [progressbar](https://github.com/schollz/progressbar) | 1548 | 22 | 2017/10/26 | 1 week ago | Basic thread-safe progress bar that works in every OS. |
| [termdash](https://github.com/mum4k/termdash) | 1537 | 26 | 2018/03/24 | 1 month ago | Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui). |
| [pflag](https://github.com/spf13/pflag) | 1360 | 33 | 2013/08/30 | 1 month ago | Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. |
| [cli](https://github.com/mitchellh/cli) | 1281 | 23 | 2013/11/03 | 1 month ago | Go library for implementing command-line interfaces. |
| [uilive](https://github.com/gosuri/uilive) | 1267 | 18 | 2015/11/16 | 3 months ago | Library for updating terminal output in realtime. |
| [docopt.go](https://github.com/docopt/docopt.go) | 1266 | 35 | 2013/08/25 | 3 months ago | Command-line arguments parser that will make you smile. |
| [mpb](https://github.com/vbauerster/mpb) | 1200 | 20 | 2016/12/14 | 4 weeks ago | Multi progress bar for terminal applications. |
| [go-arg](https://github.com/alexflint/go-arg) | 1041 | 13 | 2015/11/01 | 1 week ago | Struct-based argument parsing in Go. |
| [aurora](https://github.com/logrusorgru/aurora) | 1034 | 8 | 2016/11/06 | 4 days ago | ANSI terminal colors that supports fmt.Printf/Sprintf. |
| [gcli](https://github.com/tcnksm/gcli) | 907 | 26 | 2014/06/19 | 3 years ago | The easy way to start building Golang command line applications. |
| [color](https://github.com/gookit/color) | 773 | 12 | 2018/07/01 | 1 week ago | Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows. |
| [liner](https://github.com/peterh/liner) | 756 | 23 | 2012/08/15 | 3 weeks ago | Go readline-like library for command-line interfaces. |
| [complete](https://github.com/posener/complete) | 751 | 15 | 2017/05/05 | 2 weeks ago | Write bash completions in Go + Go command bash completion. |
| [mow.cli](https://github.com/jawher/mow.cli) | 718 | 19 | 2014/12/18 | 4 months ago | Go library for building CLI applications with sophisticated flag and argument parsing and validation. |
| [flaggy](https://github.com/integrii/flaggy) | 709 | 18 | 2018/03/05 | 7 months ago | A robust and idiomatic flags package with excellent subcommand support. |
| [uitable](https://github.com/gosuri/uitable) | 611 | 14 | 2015/11/13 | 3 months ago | Library to improve readability in terminal apps using tabular data. |
| [cli](https://github.com/mkideal/cli) | 554 | 22 | 2016/02/26 | 4 months ago | Feature-rich and easy to use command-line package based on golang struct tags. |
| [ops](https://github.com/nanovms/ops) | 545 | 24 | 2018/09/10 | 1 day ago | Unikernel Builder/Orchestrator. |
| [go-colorable](https://github.com/mattn/go-colorable) | 501 | 19 | 2014/07/30 | 2 weeks ago | Colorable writer for windows. |
| [go-isatty](https://github.com/mattn/go-isatty) | 486 | 10 | 2014/04/01 | 3 weeks ago | isatty for golang. |
| [pterm](https://github.com/pterm/pterm) | 399 | 13 | 2020/09/17 | 1 week ago | A library to beautify console output on every platform with many combinable components. |
| [chalk](https://github.com/ttacon/chalk) | 358 | 9 | 2014/07/18 | 1 year ago | Intuitive package for prettifying terminal/console output. |
| [argparse](https://github.com/akamensky/argparse) | 287 | 14 | 2017/11/24 | 1 month ago | Command line argument parser inspired by Python's argparse module. |
| [simpletable](https://github.com/alexeyco/simpletable) | 284 | 4 | 2017/03/29 | 6 months ago | Simple tables in terminal with Go. |
| [tabby](https://github.com/cheynewallace/tabby) | 279 | 3 | 2018/12/17 | 1 month ago | A tiny library for super simple Golang tables. |
| [go-colortext](https://github.com/daviddengcn/go-colortext) | 206 | 9 | 2013/01/23 | 10 months ago | Go library for color output in terminals. |
| [commandeer](https://github.com/jaffee/commandeer) | 142 | 6 | 2017/10/12 | 6 months ago | Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags. |
| [yacspin](https://github.com/theckman/yacspin) | 136 | 5 | 2019/12/29 | 8 months ago | Yet Another CLi Spinner package, for working with terminal spinners. |
| [wmenu](https://github.com/dixonwille/wmenu) | 127 | 2 | 2016/04/20 | 8 months ago | Easy to use menu structure for cli applications that prompts users to make choices. |
| [sflags](https://github.com/octago/sflags) | 123 | 5 | 2016/12/04 | 1 day ago | Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries. |
| [flag](https://github.com/cosiner/flag) | 113 | 5 | 2016/10/05 | 1 month ago | Simple but powerful command line option parsing library for Go supporting subcommand. |
| [clif](https://github.com/ukautz/clif) | 106 | 3 | 2015/05/30 | 2 years ago | Small command line interface framework. |
| [flag](https://github.com/zhuah/flag) | 102 | 5 | 2016/10/05 | 1 year ago | Simple but powerful command line option parsing library for Go supporting subcommand. |
| [job](https://github.com/liujianping/job) | 94 | 1 | 2019/04/09 | 7 months ago | JOB, make your short-term command as a long-term job. |
| [cli](https://github.com/teris-io/cli) | 78 | 1 | 2017/05/24 | 1 year ago | Simple and complete API for building command line interfaces in Go. |
| [cfmt](https://github.com/mingrammer/cfmt) | 77 | 3 | 2018/03/15 | 2 years ago | Contextual fmt inspired by bootstrap color classes. |
| [env](https://github.com/codingconcepts/env) | 73 | 1 | 2017/06/14 | 5 months ago | Tag-based environment configuration for structs. |
| [clir](https://github.com/leaanthony/clir) | 61 | 2 | 2019/11/18 | 2 months ago | A Simple and Clear CLI library. Dependency free. |
| [cmdr](https://github.com/hedzr/cmdr) | 60 | 5 | 2019/05/15 | 2 days ago | A POSIX/GNU style, getopt-like command-line UI Go library. |
| [tabular](https://github.com/InVisionApp/tabular) | 50 | 92 | 2018/04/23 | 2 years ago | Print ASCII tables from command line utilities without the need to pass large sets of data to the API. |
| [gocmd](https://github.com/devfacet/gocmd) | 48 | 3 | 2018/01/08 | 2 years ago | Go library for building command line applications. |
| [wlog](https://github.com/dixonwille/wlog) | 46 | 1 | 2016/04/13 | 1 year ago | Simple logging interface that supports cross-platform color and concurrency. |
| [strumt](https://github.com/antham/strumt) | 41 | 1 | 2017/06/19 | 1 month ago | Library to create prompt chain. |
| [flagvar](https://github.com/sgreben/flagvar) | 36 | 2 | 2018/05/18 | 7 months ago | A collection of flag argument types for Go's standard `flag` package. |
| [cmd](https://github.com/posener/cmd) | 32 | 2 | 2019/10/29 | 4 months ago | Extends the standard `flag` package to support sub commands and more in idomatic way. |
| [go-getoptions](https://github.com/DavidGamba/go-getoptions) | 30 | 4 | 2015/12/18 | 2 months ago | Go option parser inspired on the flexibility of Perl’s GetOpt::Long. |
| [ctc](https://github.com/wzshiming/ctc) | 29 | 0 | 2018/04/27 | 7 months ago | The non-invasive cross-platform terminal color library does not need to modify the Print method. |
| [argv](https://github.com/cosiner/argv) | 27 | 1 | 2017/01/22 | 10 months ago | Go library to split command line string as arguments array using the bash syntax. |
| [colourize](https://github.com/TreyBastian/colourize) | 23 | 3 | 2015/05/11 | 4 years ago | Go library for ANSI colour text in terminals. |
| [go-commander](https://github.com/yitsushi/go-commander) | 21 | 1 | 2016/10/10 | 8 months ago | Go library to simplify CLI workflow. |
| [argv](https://github.com/zhuah/argv) | 19 | 1 | 2017/01/22 | 1 year ago | Go library to split command line string as arguments array using the bash syntax. |
| [cfmt](https://github.com/i582/cfmt) | 19 | 1 | 2020/11/13 | 2 months ago | Simple and convenient formatted stylized output fully compatible with fmt library. |
| [sand](https://github.com/Zaba505/sand) | 13 | 1 | 2018/11/18 | 2 years ago | Simple API for creating interpreters and so much more. |
| [ts](https://github.com/liujianping/ts) | 13 | 0 | 2019/06/25 | 1 year ago | Timestamp convert & compare tool. |
| [go-ataman](https://github.com/workanator/go-ataman) | 9 | 2 | 2017/05/17 | 1 month ago | Go library for rendering ANSI colored text templates in terminals. |
| [table](https://github.com/tomlazar/table) | 5 | 1 | 2020/09/22 | 1 week ago | Small library for terminal color based tables . |

## Configuration
        
*Libraries for configuration parsing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [viper](https://github.com/spf13/viper) | 14728 | 228 | 2014/04/02 | 3 days ago | Go configuration with fangs. |
| [godotenv](https://github.com/joho/godotenv) | 3572 | 42 | 2013/07/30 | 1 week ago | Go port of Ruby's dotenv library (Loads environment variables from `.env`). |
| [envconfig](https://github.com/kelseyhightower/envconfig) | 3450 | 41 | 2013/11/06 | 2 months ago | Go library for managing configuration data from environment variables. |
| [ini](https://github.com/go-ini/ini) | 2376 | 78 | 2014/12/18 | 4 weeks ago | Go package to read and write INI files. |
| [env](https://github.com/caarlos0/env) | 1778 | 24 | 2015/07/28 | 6 days ago | Parse environment variables to Go structs (with defaults). |
| [konfig](https://github.com/lalamove/konfig) | 593 | 14 | 2019/01/18 | 3 months ago | Composable, observable and performant config handling for Go for the distributed processing era. |
| [koanf](https://github.com/knadh/koanf) | 431 | 10 | 2019/06/18 | 6 days ago | Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line. |
| [confita](https://github.com/heetch/confita) | 340 | 25 | 2017/12/21 | 7 months ago | Load configuration in cascade from multiple backends into a struct. |
| [store](https://github.com/tucnak/store) | 254 | 5 | 2015/10/03 | 3 years ago | Lightweight configuration manager for Go. |
| [config](https://github.com/JeremyLoy/config) | 252 | 1 | 2019/04/02 | 6 months ago | Cloud native application configuration. Bind ENV to structs in only two lines. |
| [config](https://github.com/olebedev/config) | 235 | 8 | 2014/04/21 | 1 year ago | JSON or YAML configuration wrapper with environment variables and flags parsing. |
| [hjson-go](https://github.com/hjson/hjson-go) | 230 | 8 | 2016/08/05 | 3 months ago | Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments. |
| [cleanenv](https://github.com/ilyakaznacheev/cleanenv) | 225 | 7 | 2019/07/12 | 2 months ago | Minimalistic configuration reader (from files, ENV, and wherever you want). |
| [config](https://github.com/gookit/config) | 214 | 8 | 2018/07/07 | 4 days ago | application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge. |
| [config](https://github.com/joshbetz/config) | 203 | 2 | 2017/04/02 | 1 year ago | Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP. |
| [envconfig](https://github.com/vrischmann/envconfig) | 192 | 5 | 2015/04/21 | 1 month ago | Read your configuration from environment variables. |
| [aconfig](https://github.com/cristalhq/aconfig) | 176 | 7 | 2020/06/26 | 4 weeks ago | Simple, useful and opinionated config loader. |
| [goconfig](https://github.com/gosidekick/goconfig) | 145 | 12 | 2016/12/18 | 5 months ago | Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. |
| [gcfg](https://github.com/go-gcfg/gcfg) | 140 | 7 | 2015/08/17 | 7 months ago | read INI-style configuration files into Go structs; supports user-defined types and subsections. |
| [fig](https://github.com/kkyr/fig) | 140 | 4 | 2020/01/16 | 9 months ago | Tiny library for reading configuration from a file and from environment variables (with validation & defaults). |
| [goconfig](https://github.com/crgimenes/goconfig) | 131 | 11 | 2016/12/18 | 9 months ago | Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. |
| [config](https://github.com/golobby/config) | 122 | 4 | 2019/10/15 | 2 days ago | A lightweight yet powerful config package for Go projects. |
| [envh](https://github.com/antham/envh) | 94 | 3 | 2017/01/12 | 1 month ago | Helpers to manage environment variables. |
| [envcfg](https://github.com/tomazk/envcfg) | 93 | 1 | 2014/11/29 | 3 years ago | Un-marshaling environment variables to Go structs. |
| [xdg](https://github.com/adrg/xdg) | 81 | 3 | 2014/08/22 | 3 weeks ago | Go implementation of the [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html) and [XDG user directories](https://wiki.archlinux.org/index.php/XDG_user_directories). |
| [harvester](https://github.com/beatlabs/harvester) | 77 | 12 | 2019/04/09 | 3 weeks ago | Harvester, a easy to use static and dynamic configuration package supportig seeding, env vars and Consul integration. |
| [configuro](https://github.com/sherifabdlnaby/configuro) | 69 | 4 | 2020/04/09 | 1 month ago | opinionated configuration loading & validation framework from ENV and Files focused towards 12-Factor compliant applications. |
| [xdg](https://github.com/OpenPeeDeeP/xdg) | 61 | 3 | 2017/07/20 | 3 months ago | Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html). |
| [gofigure](https://github.com/ian-kent/gofigure) | 59 | 5 | 2014/11/25 | 1 year ago | Go application configuration made easy. |
| [configure](https://github.com/paked/configure) | 55 | 3 | 2015/06/14 | 2 years ago | Provides configuration through multiple sources, including JSON, flags and environment variables. |
| [configuration](https://github.com/BoRuDar/configuration) | 37 | 2 | 2019/11/27 | 6 months ago | Library for initializing configuration structs from env variables, files, flags and 'default' tag. |
| [go-aws-ssm](https://github.com/PaddleHQ/go-aws-ssm) | 37 | 17 | 2019/01/24 | 4 months ago | Go package that fetches parameters from AWS System Manager - Parameter Store. |
| [ingo](https://github.com/schachmat/ingo) | 32 | 1 | 2016/02/07 | 3 years ago | Flags persisted in an ini-like config file. |
| [go-up](https://github.com/ufoscout/go-up) | 29 | 1 | 2018/02/18 | 1 year ago | A simple configuration library with recursive placeholders resolution and no magic. |
| [mini](https://github.com/sasbury/mini) | 28 | 1 | 2015/04/29 | 2 years ago | Golang package for parsing ini-style configuration files. |
| [genv](https://github.com/sakirsensoy/genv) | 20 | 1 | 2019/07/15 | 1 year ago | Read environment variables easily with dotenv support. |
| [hocon](https://github.com/gurkankaymak/hocon) | 20 | 1 | 2020/03/01 | 3 months ago | Configuration library for working with the HOCON(a human-friendly JSON superset) format, supports features like environment variables, referencing other values, comments and multiple files. |
| [conflate](https://github.com/the4thamigo-uk/conflate) | 19 | 0 | 2018/02/01 | 4 months ago | Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema. |
| [envconf](https://github.com/ian-kent/envconf) | 10 | 1 | 2014/10/26 | 6 years ago | Configuration from environment. |
| [go-ssm-config](https://github.com/ianlopshire/go-ssm-config) | 8 | 1 | 2019/12/02 | 2 months ago | Go utility for loading configuration parameters from AWS SSM (Parameter Store). |
| [sprbox](https://github.com/oblq/sprbox) | 4 | 2 | 2018/07/17 | 9 months ago | Build-environment aware toolbox factory and agnostic config parser (YAML, TOML, JSON and Environment vars). |
| [swap](https://github.com/oblq/swap) | 4 | 1 | 2020/04/12 | 9 months ago | Instantiate/configure structs recursively, based on build environment. (YAML, TOML, JSON and env). |
| [typenv](https://github.com/diegomarangoni/typenv) | 4 | 1 | 2020/06/30 | 6 months ago | Minimalistic, zero dependency, typed environment variables library. |
| [env](https://github.com/nasermirzaei89/env) | 3 | 0 | 2019/07/24 | 4 months ago | Simple useful package for read environment variables. |
| [go-ini](https://github.com/subpop/go-ini) | 3 | 1 | 2019/09/11 | 3 months ago | A Go package that marshals and unmarshals INI-files. |

## Continuous Integration
        
*Tools for help with continuous integration.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [drone](https://github.com/drone/drone) | 22522 | 570 | 2014/02/07 | 1 hour ago | Drone is a Continuous Integration platform built on Docker, written in Go. |
| [cds](https://github.com/ovh/cds) | 3274 | 81 | 2016/10/11 | 2 days ago | Enterprise-Grade CI/CD and DevOps Automation Open Source Platform. |
| [goveralls](https://github.com/mattn/goveralls) | 678 | 13 | 2013/04/17 | 1 month ago | Go integration for Coveralls.io continuous code coverage tracking system. |
| [overalls](https://github.com/go-playground/overalls) | 105 | 3 | 2015/07/30 | 1 year ago | Multi-Package go project coverprofile for tools like goveralls. |
| [duci](https://github.com/duck8823/duci) | 62 | 3 | 2018/04/01 | 1 week ago | A simple ci server no needs domain specific languages. |
| [gomason](https://github.com/nikogura/gomason) | 50 | 1 | 2017/11/18 | 2 months ago | Test, Build, Sign, and Publish your go binaries from a clean workspace. |
| [roveralls](https://github.com/lawrencewoodman/roveralls) | 14 | 1 | 2016/06/26 | 3 years ago | Recursive coverage testing tool. |

## CSS Preprocessors
        
*Libraries for preprocessing CSS files.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gcss](https://github.com/yosssi/gcss) | 440 | 16 | 2014/09/04 | 6 years ago | Pure Go CSS Preprocessor. |
| [go-libsass](https://github.com/wellington/go-libsass) | 167 | 8 | 2015/04/19 | 3 months ago | Go wrapper to the 100% Sass compatible libsass project. |

## Data Structures
        
*Generic datastructures and algorithms in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gods](https://github.com/emirpasic/gods) | 9593 | 353 | 2015/03/04 | 2 days ago | Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc. |
| [go-datastructures](https://github.com/Workiva/go-datastructures) | 5917 | 323 | 2014/10/29 | 10 months ago | Collection of useful, performant, and thread-safe data structures. |
| [golang-set](https://github.com/deckarep/golang-set) | 1813 | 43 | 2013/07/03 | 1 month ago | Thread-Safe and Non-Thread-Safe high-performance sets for Go. |
| [gota](https://github.com/go-gota/gota) | 1514 | 74 | 2016/02/06 | 1 day ago | Implementation of dataframes, series, and data wrangling methods for Go. |
| [BoomFilters](https://github.com/tylertreat/BoomFilters) | 1320 | 40 | 2015/02/06 | 9 months ago | Probabilistic data structures for processing continuous, unbounded streams. |
| [roaring](https://github.com/RoaringBitmap/roaring) | 1095 | 41 | 2014/07/10 | 2 weeks ago | Go package implementing compressed bitsets. |
| [bloom](https://github.com/willf/bloom) | 957 | 35 | 2011/05/21 | 5 months ago | Go package implementing Bloom filters. |
| [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) | 753 | 17 | 2015/06/28 | 1 month ago | Cuckoo filter: a good alternative to a counting bloom filter implemented in Go. |
| [gocache](https://github.com/eko/gocache) | 751 | 19 | 2019/10/05 | 1 week ago | A complete Go cache library with mutiple stores (memory, memcache, redis, ...), chainable, loadable, metrics cache and more. |
| [hyperloglog](https://github.com/axiomhq/hyperloglog) | 737 | 17 | 2017/06/18 | 1 year ago | HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction. |
| [bitset](https://github.com/willf/bitset) | 647 | 30 | 2011/05/11 | 6 months ago | Go package implementing bitsets. |
| [trie](https://github.com/derekparker/trie) | 519 | 20 | 2014/03/06 | 6 months ago | Trie implementation in Go. |
| [algorithms](https://github.com/shady831213/algorithms) | 475 | 19 | 2018/01/31 | 1 year ago | Algorithms and data structures.CLRS study. |
| [go-geoindex](https://github.com/hailocab/go-geoindex) | 327 | 66 | 2015/01/22 | 3 years ago | In-memory geo index. |
| [mafsa](https://github.com/smartystreets-archives/mafsa) | 283 | 16 | 2014/11/07 | 1 year ago | MA-FSA implementation with Minimal Perfect Hashing. |
| [go-edlib](https://github.com/hbollon/go-edlib) | 247 | 9 | 2020/08/18 | 1 week ago | Go string comparison and edit distance algorithms library (Levenshtein, LCS, Hamming, Damerau levenshtein, Jaro-Winkler, etc.) compatible with Unicode. |
| [gostl](https://github.com/liyue201/gostl) | 245 | 7 | 2019/10/12 | 1 month ago | Data structure and algorithm library for go, designed to provide functions similar to C++ STL. |
| [merkletree](https://github.com/cbergoon/merkletree) | 233 | 6 | 2017/04/12 | 1 year ago | Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures. |
| [hilbert](https://github.com/google/hilbert) | 224 | 21 | 2015/08/06 | 2 years ago | Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves. |
| [goskiplist](https://github.com/ryszard/goskiplist) | 223 | 15 | 2012/05/09 | 1 year ago | Skip list implementation in Go. |
| [ttlcache](https://github.com/ReneKroon/ttlcache) | 207 | 11 | 2014/12/13 | 1 week ago | In-memory string-interface{} cache with various time-based expiration options and callbacks. |
| [deque](https://github.com/gammazero/deque) | 201 | 6 | 2018/04/24 | 4 months ago | Fast ring-buffer deque (double-ended queue). |
| [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) | 172 | 7 | 2016/04/01 | 6 months ago | Go implementation of Adaptive Radix Tree. |
| [binpacker](https://github.com/zhuangsirui/binpacker) | 163 | 13 | 2016/02/02 | 2 years ago | Binary packer and unpacker helps user build custom binary stream. |
| [skiplist](https://github.com/MauriceGit/skiplist) | 138 | 5 | 2018/06/23 | 1 year ago | Very fast Go Skiplist implementation. |
| [bloom](https://github.com/zentures/bloom) | 137 | 7 | 2013/09/03 | 2 years ago | Bloom filters implemented in Go. |
| [iter](https://github.com/disksing/iter) | 130 | 5 | 2019/10/20 | 1 year ago | Go implementation of C++ STL iterators and algorithms. |
| [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) | 128 | 2 | 2019/01/10 | 9 months ago | Concurrent FIFO queue. |
| [levenshtein](https://github.com/agnivade/levenshtein) | 121 | 3 | 2014/07/30 | 4 months ago | Implementation to calculate levenshtein distance in Go. |
| [go-rquad](https://github.com/arl/go-rquad) | 112 | 3 | 2016/09/12 | 10 months ago | Region quadtrees with efficient point location and neighbour finding. |
| [ring](https://github.com/tannerryan/ring) | 112 | 1 | 2019/01/27 | 5 months ago | Go implementation of a high performance, thread safe bloom filter. |
| [ring](https://github.com/TheTannerRyan/ring) | 110 | 1 | 2019/01/27 | 1 year ago | Go implementation of a high performance, thread safe bloom filter. |
| [encoding](https://github.com/zentures/encoding) | 104 | 6 | 2013/09/20 | 3 years ago | Integer Compression Libraries for Go. |
| [bit](https://github.com/yourbasic/bit) | 98 | 6 | 2017/05/03 | 2 years ago | Golang set data structure with bonus bit-twiddling functions. |
| [conjungo](https://github.com/InVisionApp/conjungo) | 92 | 103 | 2016/12/29 | 3 months ago | A small, powerful and flexible merge library. |
| [remember-go](https://github.com/rocketlaunchr/remember-go) | 76 | 4 | 2019/04/04 | 4 months ago | A universal interface for caching slow database queries (backed by redis, memcached, ristretto, or in-memory). |
| [skiplist](https://github.com/gansidui/skiplist) | 71 | 8 | 2014/11/18 | 6 years ago | Skiplist implementation in Go. |
| [go-mcache](https://github.com/OrlovEvgeny/go-mcache) | 60 | 3 | 2018/04/14 | 1 year ago | Fast in-memory key:value store/cache library. Pointer caches. |
| [bloom](https://github.com/yourbasic/bloom) | 58 | 2 | 2017/05/06 | 3 years ago | Golang Bloom filter implementation. |
| [levenshtein](https://github.com/agext/levenshtein) | 56 | 1 | 2016/04/08 | 4 months ago | Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. |
| [count-min-log](https://github.com/seiflotfy/count-min-log) | 52 | 2 | 2015/08/16 | 4 years ago | Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory). |
| [crunch](https://github.com/superwhiskers/crunch) | 42 | 6 | 2019/02/27 | 3 weeks ago | Go package implementing buffers for handling various datatypes easily. |
| [nan](https://github.com/kak-tus/nan) | 34 | 3 | 2020/05/05 | 3 months ago | Zero allocation Nullable structures in one library with handy conversion functions, marshallers and unmarshallers. |
| [concurrent-writer](https://github.com/free/concurrent-writer) | 30 | 4 | 2017/09/18 | 3 years ago | Highly concurrent drop-in replacement for `bufio.Writer`. |
| [hide](https://github.com/emvi/hide) | 30 | 3 | 2019/01/16 | 1 week ago | ID type with marshalling to/from hash to prevent sending IDs to clients. |
| [goset](https://github.com/zoumo/goset) | 29 | 1 | 2017/08/25 | 2 months ago | A useful Set collection implementation for Go. |
| [pipeline](https://github.com/hyfather/pipeline) | 29 | 1 | 2018/04/25 | 2 years ago | An implementation of pipelines with fan-in and fan-out. |
| [deque](https://github.com/edwingeng/deque) | 23 | 3 | 2019/02/01 | 1 month ago | A highly optimized double-ended queue. |
| [timedmap](https://github.com/zekroTJA/timedmap) | 23 | 2 | 2019/01/30 | 1 month ago | Map with expiring key-value pairs. |
| [typ](https://github.com/gurukami/typ) | 23 | 1 | 2019/03/03 | 1 year ago | Null Types, Safe primitive type conversion and fetching value from complex structures. |
| [dict](https://github.com/srfrog/dict) | 18 | 0 | 2019/04/23 | 3 months ago | Python-like dictionaries (dict) for Go. |
| [go-ef](https://github.com/amallia/go-ef) | 17 | 2 | 2017/09/22 | 3 years ago | A Go implementation of the Elias-Fano encoding. |
| [mspm](https://github.com/BlackRabbitt/mspm) | 15 | 3 | 2018/05/17 | 2 years ago | Multi-String Pattern Matching Algorithm for information retrieval. |
| [null](https://github.com/emvi/null) | 15 | 2 | 2018/07/04 | 1 week ago | Nullable Go types that can be marshalled/unmarshalled to/from JSON. |
| [set](https://github.com/StudioSol/set) | 12 | 16 | 2018/07/20 | 4 months ago | Simple set data structure implementation in Go using LinkedHashMap. |
| [ptrie](https://github.com/viant/ptrie) | 11 | 7 | 2019/05/20 | 5 months ago | An implementation of prefix tree. |
| [treap](https://github.com/perdata/treap) | 10 | 2 | 2018/09/16 | 1 year ago | Persistent, fast ordered map using tree heaps. |
| [cmap](https://github.com/lrita/cmap) | 10 | 1 | 2019/11/26 | 6 months ago | a thread-safe concurrent map for go, support using `interface{}` as key and auto scale up shards. |
| [gofal](https://github.com/xxjwxc/gofal) | 9 | 2 | 2019/08/05 | 1 year ago | fractional api for Go. |
| [parsefields](https://github.com/MonaxGT/parsefields) | 5 | 1 | 2019/04/12 | 1 year ago | Tools for parse JSON-like logs for collecting unique fields and events. |
| [goterator](https://github.com/yaa110/goterator) | 3 | 1 | 2020/08/12 | 2 months ago | Iterator implementation to provide map and reduce functionalities. |
| [slices](https://github.com/srfrog/slices) | 2 | 1 | 2020/07/02 | 3 months ago | Functions that operate on slices; like `package strings` but adapted to work with slices. |

## Database
        
*SQL query builder, libraries for building and using SQL.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [prometheus](https://github.com/prometheus/prometheus) | 35278 | 1167 | 2012/11/24 | 2 hours ago | Monitoring system and time series database. |
| [tidb](https://github.com/pingcap/tidb) | 26800 | 1355 | 2015/09/06 | 1 hour ago | TiDB is a distributed SQL database. Inspired by the design of Google F1. |
| [influxdb](https://github.com/influxdata/influxdb) | 20486 | 754 | 2013/09/26 | 2 days ago | Scalable datastore for metrics, events, and real-time analytics. |
| [cockroach](https://github.com/cockroachdb/cockroach) | 19917 | 716 | 2014/02/06 | 1 hour ago | Scalable, Geo-Replicated, Transactional Datastore. |
| [dgraph](https://github.com/dgraph-io/dgraph) | 14822 | 370 | 2015/08/25 | 1 day ago | Scalable, Distributed, Low Latency, High Throughput Graph Database. |
| [vitess](https://github.com/vitessio/vitess) | 11398 | 510 | 2013/06/27 | 1 hour ago | vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. |
| [bolt](https://github.com/boltdb/bolt) | 10426 | 343 | 2013/12/20 | 2 years ago | Low-level key/value database for Go. |
| [groupcache](https://github.com/golang/groupcache) | 9806 | 482 | 2013/07/22 | 3 months ago | Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. |
| [badger](https://github.com/dgraph-io/badger) | 8831 | 250 | 2017/01/26 | 1 day ago | Fast key-value store in Go. |
| [rqlite](https://github.com/rqlite/rqlite) | 7656 | 218 | 2014/08/23 | 5 days ago | The lightweight, distributed, relational database built on SQLite. |
| [pgweb](https://github.com/sosedoff/pgweb) | 6762 | 149 | 2014/10/09 | 1 month ago | Web-based PostgreSQL database browser. |
| [migrate](https://github.com/golang-migrate/migrate) | 5851 | 70 | 2018/01/19 | 4 days ago | Database migrations. CLI and Golang library. |
| [kingshard](https://github.com/flike/kingshard) | 5556 | 409 | 2015/07/04 | 1 week ago | kingshard is a high performance proxy for MySQL powered by Golang. |
| [go-cache](https://github.com/patrickmn/go-cache) | 4705 | 108 | 2012/01/02 | 6 days ago | In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. |
| [bigcache](https://github.com/allegro/bigcache) | 4651 | 113 | 2016/03/23 | 2 days ago | Efficient key/value cache for gigabytes of data. |
| [goleveldb](https://github.com/syndtr/goleveldb) | 4139 | 177 | 2013/01/23 | 1 week ago | Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go. |
| [bbolt](https://github.com/etcd-io/bbolt) | 4093 | 115 | 2017/06/17 | 3 weeks ago | An embedded key/value database for Go. |
| [orchestrator](https://github.com/openark/orchestrator) | 3941 | 271 | 2016/11/30 | 6 days ago | MySQL replication topology manager & visualizer. |
| [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) | 3755 | 100 | 2018/09/30 | 2 days ago | fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL. |
| [squirrel](https://github.com/Masterminds/squirrel) | 3649 | 47 | 2014/01/18 | 3 weeks ago | Go library that helps you build SQL queries. |
| [ledisdb](https://github.com/ledisdb/ledisdb) | 3527 | 185 | 2014/04/30 | 9 months ago | Ledisdb is a high performance NoSQL like Redis based on LevelDB. |
| [orchestrator](https://github.com/github/orchestrator) | 3362 | 299 | 2016/11/30 | 1 year ago | MySQL replication topology manager & visualizer. |
| [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) | 3341 | 176 | 2015/01/15 | 1 week ago | Sync your MySQL data into Elasticsearch automatically. |
| [ledisdb](https://github.com/siddontang/ledisdb) | 3265 | 188 | 2014/04/30 | 9 months ago | Ledisdb is a high performance NoSQL like Redis based on LevelDB. |
| [buntdb](https://github.com/tidwall/buntdb) | 3103 | 101 | 2016/07/19 | 6 days ago | Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support. |
| [go-mysql](https://github.com/siddontang/go-mysql) | 2839 | 160 | 2014/02/21 | 1 week ago | Go toolset to handle MySQL protocol and replication. |
| [xo](https://github.com/xo/xo) | 2691 | 74 | 2016/02/05 | 2 months ago | Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. |
| [tiedot](https://github.com/HouzuoGuo/tiedot) | 2572 | 157 | 2013/05/26 | 10 months ago | Your NoSQL database powered by Golang. |
| [prest](https://github.com/prest/prest) | 2562 | 86 | 2016/11/22 | 1 day ago | Simplify and accelerate development, ⚡ instant, realtime, high-performance on any Postgres application, existing or new. |
| [sql-migrate](https://github.com/rubenv/sql-migrate) | 2037 | 36 | 2014/09/09 | 1 day ago | Database migration tool. Allows embedding migrations into the application using go-bindata. |
| [immudb](https://github.com/codenotary/immudb) | 1734 | 37 | 2019/11/07 | 2 days ago | immudb is a lightweight, high-speed immutable database for systems and applications written in Go. |
| [goose](https://github.com/pressly/goose) | 1562 | 45 | 2016/02/25 | 1 day ago | Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts. |
| [nutsdb](https://github.com/xujiajun/nutsdb) | 1497 | 48 | 2018/12/07 | 1 week ago | Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set. |
| [cache2go](https://github.com/muesli/cache2go) | 1465 | 66 | 2013/11/11 | 1 month ago | In-memory key:value cache which supports automatic invalidation based on timeouts. |
| [gcache](https://github.com/bluele/gcache) | 1367 | 45 | 2015/01/24 | 3 weeks ago | Cache library with support for expirable Cache, LFU, LRU and ARC. |
| [gendry](https://github.com/didi/gendry) | 1160 | 63 | 2017/12/01 | 1 month ago | Non-invasive SQL builder and powerful data binder. |
| [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) | 1140 | 69 | 2018/04/11 | 9 months ago | CovenantSQL is a SQL database on blockchain. |
| [diskv](https://github.com/peterbourgon/diskv) | 1006 | 37 | 2012/03/21 | 2 weeks ago | Home-grown disk-backed key-value store. |
| [fastcache](https://github.com/VictoriaMetrics/fastcache) | 995 | 26 | 2018/11/22 | 3 days ago | fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead. |
| [goqu](https://github.com/doug-martin/goqu) | 958 | 34 | 2015/02/21 | 2 weeks ago | Idiomatic SQL builder and query library. |
| [skeema](https://github.com/skeema/skeema) | 842 | 28 | 2016/10/31 | 1 day ago | Pure-SQL schema management system for MySQL, with support for sharding and external online schema change tools. |
| [moss](https://github.com/couchbase/moss) | 793 | 74 | 2016/02/06 | 1 month ago | Moss is a simple LSM key-value storage engine written in 100% Go. |
| [pogreb](https://github.com/akrylysov/pogreb) | 692 | 25 | 2018/01/06 | 20 hours ago | Embedded key-value store for read-heavy workloads. |
| [bitcask](https://github.com/prologic/bitcask) | 607 | 18 | 2019/03/12 | 1 month ago | Bitcask is an embeddable, persistent and fast key-value (KV) database written in pure Go with predictable read/write performance, low latency and high throughput thanks to the bitcask on-disk layout (LSM+WAL). |
| [eliasdb](https://github.com/krotik/eliasdb) | 606 | 23 | 2016/08/13 | 2 months ago | Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language. |
| [dotsql](https://github.com/gchaincl/dotsql) | 573 | 25 | 2014/11/20 | 10 months ago | Go library that helps you keep sql files in one place and use them with ease. |
| [chproxy](https://github.com/Vertamedia/chproxy) | 557 | 29 | 2017/09/18 | 2 days ago | HTTP proxy for ClickHouse database. |
| [gormigrate](https://github.com/go-gormigrate/gormigrate) | 556 | 6 | 2016/08/31 | 3 months ago | Database schema migration helper for Gorm ORM. |
| [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) | 519 | 27 | 2015/12/10 | 3 weeks ago | Powerful data retrieval methods as well as DB-agnostic query building capabilities. |
| [levigo](https://github.com/jmhodges/levigo) | 390 | 23 | 2012/01/17 | 8 months ago | Levigo is a Go wrapper for LevelDB. |
| [pg_timetable](https://github.com/cybertec-postgresql/pg_timetable) | 365 | 19 | 2018/12/19 | 2 weeks ago | Advanced scheduling for PostgreSQL. |
| [jet](https://github.com/go-jet/jet) | 359 | 10 | 2019/03/02 | 2 weeks ago | Framework for writing type-safe SQL queries in Go, with ability to easily convert database query result into desired arbitrary object structure. |
| [pudge](https://github.com/recoilme/pudge) | 284 | 11 | 2018/11/20 | 1 week ago | Fast and simple  key/value store written using Go's standard library. |
| [dbq](https://github.com/rocketlaunchr/dbq) | 263 | 9 | 2019/07/11 | 5 months ago | Zero boilerplate database operations for Go. |
| [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) | 257 | 7 | 2017/04/29 | 3 weeks ago | Collects small insterts and sends big requests to ClickHouse servers. |
| [sqrl](https://github.com/elgris/sqrl) | 217 | 8 | 2014/06/25 | 4 days ago | SQL query builder, fork of Squirrel with improved performance. |
| [vasto](https://github.com/chrislusf/vasto) | 205 | 19 | 2018/01/16 | 1 year ago | A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption. |
| [kivik](https://github.com/go-kivik/kivik) | 201 | 6 | 2017/02/09 | 1 month ago | Kivik provides a common Go and GopherJS client library for CouchDB, PouchDB, and similar databases. |
| [piladb](https://github.com/fern4lvarez/piladb) | 183 | 12 | 2015/09/08 | 3 months ago | Lightweight RESTful database engine based on stack data structures. |
| [myreplication](https://github.com/2tvenom/myreplication) | 172 | 20 | 2015/02/04 | 2 years ago | MySql binary log replication listener. Supports statement and row based replication. |
| [goose](https://github.com/steinbacher/goose) | 144 | 4 | 2016/03/04 | 4 years ago | Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts. |
| [sqlingo](https://github.com/lqs/sqlingo) | 137 | 12 | 2018/11/18 | 1 month ago | A lightweight DSL to build SQL in Go. |
| [octillery](https://github.com/blastrain/octillery) | 122 | 17 | 2018/11/26 | 1 month ago | Go package for sharding databases ( Supports every ORM or raw SQL ). |
| [golang-scribble](https://github.com/nanobox-io/golang-scribble) | 116 | 4 | 2018/06/21 | 1 year ago | Tiny flat file JSON store. |
| [darwin](https://github.com/GuiaBolso/darwin) | 111 | 2 | 2016/04/05 | 1 year ago | Database schema evolution library for Go. |
| [databunker](https://github.com/securitybunker/databunker) | 111 | 10 | 2019/12/08 | 1 week ago | Personally identifiable information (PII) storage service built to comply with GDPR and CCPA. |
| [migrator](https://github.com/lopezator/migrator) | 106 | 5 | 2019/02/04 | 4 months ago | Dead simple Go database migration library. |
| [slowpoke](https://github.com/recoilme/slowpoke) | 98 | 8 | 2018/02/19 | 1 year ago | Key-value store with persistence. |
| [go-structured-query](https://github.com/bokwoon95/go-structured-query) | 96 | 1 | 2020/05/30 | 2 months ago | Type-safe SQL builder and struct mapper for Go. |
| [databunker](https://github.com/paranoidguy/databunker) | 93 | 9 | 2019/12/08 | 2 months ago | Personally identifiable information (PII) storage service built to comply with GDPR and CCPA. |
| [octillery](https://github.com/knocknote/octillery) | 92 | 16 | 2018/11/26 | 6 months ago | Go package for sharding databases ( Supports every ORM or raw SQL ). |
| [igor](https://github.com/galeone/igor) | 83 | 7 | 2016/03/10 | 7 months ago | Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax. |
| [cache](https://github.com/akyoto/cache) | 79 | 3 | 2019/05/11 | 11 months ago | In-memory key:value store with expiration time, 0 dependencies, <100 LoC, 100% coverage. |
| [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) | 72 | 1 | 2018/08/11 | 3 months ago | A Go package to help write migrations with go-pg/pg. |
| [bcache](https://github.com/iwanbk/bcache) | 61 | 2 | 2018/12/26 | 1 year ago | Eventually consistent distributed in-memory  cache Go library. |
| [unitdb](https://github.com/unit-io/unitdb) | 56 | 4 | 2019/08/29 | 2 weeks ago | Fast timeseries database for IoT, realtime messaging  applications. Access unitdb with pubsub over tcp or websocket using github.com/unit-io/unitd application. |
| [couchcache](https://github.com/codingsince1985/couchcache) | 52 | 3 | 2015/04/05 | 5 months ago | RESTful caching micro-service backed by Couchbase server. |
| [dbbench](https://github.com/sj14/dbbench) | 52 | 2 | 2018/11/24 | 4 days ago | Database benchmarking tool with support for several databases and scripts. |
| [godbal](https://github.com/xujiajun/godbal) | 51 | 4 | 2018/02/28 | 2 years ago | Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily. |
| [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) | 36 | 6 | 2017/12/18 | 3 years ago | BigCache with clustering support and individual item expiration. |
| [gondolier](https://github.com/emvi/gondolier) | 31 | 4 | 2017/10/18 | 1 year ago | Database migration library using struct decorators. |
| [datagen](https://github.com/codingconcepts/datagen) | 29 | 1 | 2019/04/18 | 7 months ago | A fast data generator that's multi-table aware and supports multi-row DML. |
| [prep](https://github.com/hexdigest/prep) | 27 | 3 | 2017/12/11 | 3 years ago | Use prepared SQL statements without changing your code. |
| [buildsqlx](https://github.com/arthurkushman/buildsqlx) | 26 | 1 | 2019/08/18 | 2 months ago | Go database query builder library for PostgreSQL. |
| [go-fixtures](https://github.com/RichardKnop/go-fixtures) | 25 | 1 | 2015/12/24 | 1 year ago | Django style fixtures for Golang's excellent built-in database/sql library. |
| [pravasan](https://github.com/pravasan/pravasan) | 24 | 7 | 2015/01/03 | 2 years ago | Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc. |
| [coffer](https://github.com/claygod/coffer) | 23 | 4 | 2019/05/13 | 4 months ago | Simple ACID key-value database that supports transactions. |
| [avro](https://github.com/khezen/avro) | 21 | 2 | 2019/04/07 | 7 months ago | Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes. |
| [qry](https://github.com/HnH/qry) | 20 | 2 | 2019/08/20 | 5 months ago | Tool that generates constants from files with raw SQL queries. |
| [sqlf](https://github.com/leporo/sqlf) | 16 | 3 | 2019/07/20 | 1 year ago | Fast SQL query builder. |
| [gosql](https://github.com/twharmon/gosql) | 15 | 1 | 2020/01/08 | 5 months ago | SQL Query builder with better null values support. |
| [tempdb](https://github.com/rafaeljesus/tempdb) | 14 | 3 | 2017/03/17 | 3 years ago | Key-value store for temporary items. |
| [gorocksdb](https://github.com/kapitan-k/gorocksdb) | 12 | 1 | 2017/12/28 | 3 years ago | Gorocksdb is a wrapper for [RocksDB](https://rocksdb.org) written in Go. |
| [rwdb](https://github.com/andizzle/rwdb) | 11 | 2 | 2017/10/04 | 3 years ago | rwdb provides read replica capability for multiple database servers setup. |
| [schema](https://github.com/adlio/schema) | 7 | 3 | 2019/09/24 | 1 month ago | Library to embed schema migrations for database/sql-compatible databases inside your Go binaries. |
| [mpath-go](https://github.com/spacetab-io/mpath-go) | 7 | 3 | 2020/01/09 | 1 year ago | MPTT (Modified Preorder Tree Traversal) package for SQL records - materialized path realisation. |
| [bucket](https://github.com/PumpkinSeed/bucket) | 6 | 3 | 2019/09/22 | 8 months ago | Optimized data structure framework for Couchbase specialized on one bucket usage. |
| [gostore](https://github.com/twharmon/gostore) | 5 | 2 | 2020/01/08 | 1 year ago | Gostore is a simple, durable, embedded key-value storage engine written in Go. |
| [go-pg-migrate](https://github.com/lawzava/go-pg-migrate) | 4 | 1 | 2021/01/16 | 3 weeks ago | CLI-friendly package for go-pg migrations management. |
| [tracedb](https://github.com/unit-io/tracedb) | 3 | 1 | 2019/08/29 | 9 months ago | Fast timeseries database for IoT, realtime messaging  applications. Access tracedb with pubsub over tcp or websocket using github.com/unit-io/trace application. |
| [ttlcache](https://github.com/cheshir/ttlcache) | 1 | 1 | 2021/01/06 | 1 month ago | In-memory key value storage with TTL for each record. |
| [ormlite](https://github.com/pupizoid/ormlite) | 0 | 2 | 2018/06/28 | 1 month ago | Lightweight package containing some ORM-like features and helpers for sqlite databases. |

## Database Drivers
        
*Libraries for connecting and operating databases.*

### Relational Databases
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [mysql](https://github.com/go-sql-driver/mysql) | 10502 | 415 | 2012/12/09 | 1 week ago | MySQL driver for Go. |
| [pq](https://github.com/lib/pq) | 6376 | 160 | 2012/03/12 | 1 day ago | Pure Go Postgres driver for database/sql. |
| [go-sqlite3](https://github.com/mattn/go-sqlite3) | 4616 | 150 | 2011/11/11 | 4 days ago | SQLite3 driver for go that uses database/sql. |
| [pgx](https://github.com/jackc/pgx) | 3710 | 78 | 2013/03/30 | 20 hours ago | PostgreSQL driver supporting features beyond those exposed by database/sql. |
| [go-mssqldb](https://github.com/denisenkom/go-mssqldb) | 1341 | 66 | 2013/12/16 | 2 hours ago | Microsoft MSSQL driver for Go. |
| [go-oci8](https://github.com/mattn/go-oci8) | 531 | 41 | 2012/02/29 | 1 month ago | Oracle driver for go that uses database/sql. |
| [goracle](https://github.com/go-goracle/goracle) | 283 | 28 | 2015/03/25 | 1 year ago | Oracle driver for Go, using the ODPI-C driver. |
| [godror](https://github.com/godror/godror) | 221 | 16 | 2019/11/21 | 6 days ago | Oracle driver for Go, using the ODPI-C driver. |
| [firebirdsql](https://github.com/nakagami/firebirdsql) | 137 | 18 | 2013/08/27 | 2 days ago | Firebird RDBMS SQL driver for Go. |
| [go-adodb](https://github.com/mattn/go-adodb) | 117 | 12 | 2011/11/14 | 1 year ago | Microsoft ActiveX Object DataBase driver for go that uses database/sql. |
| [gofreetds](https://github.com/minus5/gofreetds) | 105 | 22 | 2012/12/06 | 2 months ago | Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org). |
| [calcite-avatica-go](https://github.com/apache/calcite-avatica-go) | 63 | 25 | 2017/08/08 | 4 months ago | Apache Avatica/Phoenix SQL driver for database/sql. |
| [sqinn-go](https://github.com/cvilsmeier/sqinn-go) | 33 | 1 | 2020/06/06 | 3 months ago | SQLite with pure Go. |
| [bgc](https://github.com/viant/bgc) | 15 | 10 | 2016/06/13 | 1 year ago | Datastore Connectivity for BigQuery for go. |

### NoSQL Databases
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [cayley](https://github.com/cayleygraph/cayley) | 13749 | 614 | 2014/06/05 | 1 week ago | Graph database with support for multiple backends. |
| [redis](https://github.com/go-redis/redis) | 10759 | 242 | 2012/07/25 | 2 days ago | Redis client for Golang. |
| [redigo](https://github.com/gomodule/redigo) | 8177 | 297 | 2012/04/14 | 1 day ago | Redigo is a Go client for the Redis database. |
| [bleve](https://github.com/blevesearch/bleve) | 7388 | 256 | 2014/04/17 | 1 day ago | Modern text indexing library for go. |
| [elastic](https://github.com/olivere/elastic) | 5703 | 172 | 2012/12/06 | 6 days ago | Elasticsearch client for Go. |
| [riot](https://github.com/go-ego/riot) | 5624 | 200 | 2017/06/21 | 4 months ago | Go Open Source, Distributed, Simple and efficient Search Engine. |
| [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) | 5309 | 135 | 2017/02/08 | 1 day ago | Official MongoDB driver for the Go language. |
| [go-elasticsearch](https://github.com/elastic/go-elasticsearch) | 3103 | 289 | 2017/03/27 | 3 days ago | Official Elasticsearch client for Go. |
| [mgo](https://github.com/globalsign/mgo) | 1869 | 66 | 2017/04/13 | 8 months ago | (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms. |
| [rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1553 | 49 | 2013/09/12 | 5 months ago | Go language driver for RethinkDB. |
| [elastigo](https://github.com/mattbaird/elastigo) | 946 | 47 | 2012/10/12 | 2 years ago | Elasticsearch client library. |
| [elasticsql](https://github.com/cch123/elasticsql) | 668 | 34 | 2016/08/24 | 3 months ago | Convert sql to elasticsearch dsl in Go. |
| [qmgo](https://github.com/qiniu/qmgo) | 408 | 6 | 2020/08/04 | 5 days ago | The MongoDB driver for Go. It‘s based on official MongoDB driver but easier to use like Mgo. |
| [redeo](https://github.com/bsm/redeo) | 390 | 25 | 2014/03/06 | 2 months ago | Redis-protocol compatible TCP servers/services. |
| [neoism](https://github.com/jmcvetta/neoism) | 373 | 25 | 2012/07/12 | 1 year ago | Neo4j client for Golang. |
| [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) | 350 | 40 | 2014/07/26 | 2 hours ago | Aerospike client in Go language. |
| [gocb](https://github.com/couchbase/gocb) | 321 | 60 | 2015/01/15 | 2 days ago | Official Couchbase Go SDK. |
| [go-couchbase](https://github.com/couchbase/go-couchbase) | 309 | 24 | 2012/01/19 | 2 days ago | Couchbase client in Go. |
| [gokv](https://github.com/philippgille/gokv) | 304 | 7 | 2018/10/08 | 3 months ago | Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more). |
| [mgm](https://github.com/Kamva/mgm) | 247 | 13 | 2019/12/27 | 2 months ago | MongoDB model-based ODM for Go (based on official MongoDB driver). |
| [go-rejson](https://github.com/nitishm/go-rejson) | 156 | 7 | 2018/04/23 | 3 weeks ago | Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease. |
| [cachego](https://github.com/faabiosr/cachego) | 144 | 6 | 2016/10/05 | 5 months ago | Golang Cache component for multiple drivers. |
| [godis](https://github.com/piaohao/godis) | 87 | 10 | 2019/06/14 | 9 months ago | redis client implement by golang, inspired by jedis. |
| [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) | 76 | 7 | 2011/06/04 | 2 years ago | Neo4j REST Client in golang. |
| [skizze](https://github.com/seiflotfy/skizze) | 75 | 7 | 2016/01/17 | 4 years ago | probabilistic data-structures service and storage. |
| [dynago](https://github.com/underarmour/dynago) | 71 | 124 | 2015/05/18 | 3 years ago | Dynago is a principle of least surprise client for DynamoDB. |
| [arangolite](https://github.com/solher/arangolite) | 69 | 6 | 2015/10/04 | 1 year ago | Lightweight golang driver for ArangoDB. |
| [go-pilosa](https://github.com/pilosa/go-pilosa) | 44 | 21 | 2016/09/30 | 11 months ago | Go client library for Pilosa. |
| [goforestdb](https://github.com/couchbase/goforestdb) | 29 | 39 | 2014/05/14 | 4 years ago | Go bindings for ForestDB. |
| [goriak](https://github.com/zegl/goriak) | 26 | 3 | 2016/10/05 | 2 years ago | Go language driver for Riak KV. |
| [neo4j](https://github.com/cihangir/neo4j) | 26 | 4 | 2013/05/18 | 5 years ago | Neo4j Rest API Bindings for Golang. |
| [goes](https://github.com/OwnLocal/goes) | 24 | 33 | 2015/12/28 | 3 months ago | Library to interact with Elasticsearch. |
| [dsc](https://github.com/viant/dsc) | 21 | 16 | 2016/06/13 | 11 months ago | Datastore connectivity for SQL, NoSQL, structured files. |
| [xredis](https://github.com/shomali11/xredis) | 13 | 2 | 2017/06/14 | 1 year ago | Typesafe, customizable, clean & easy to use Redis client. |
| [godscache](https://github.com/defcronyke/godscache) | 8 | 3 | 2018/05/08 | 2 years ago | A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached. |
| [asc](https://github.com/viant/asc) | 6 | 10 | 2016/06/13 | 1 year ago | Datastore Connectivity for Aerospike for go. |
| [gocosmos](https://github.com/btnguyen2k/gocosmos) | 2 | 1 | 2020/12/06 | 1 month ago | REST client and standard `database/sql` driver for Azure Cosmos DB. |

## Date and Time
        
*Libraries for working with dates and times.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [now](https://github.com/jinzhu/now) | 2907 | 59 | 2013/11/18 | 1 day ago | Now is a time toolkit for golang. |
| [dateparse](https://github.com/araddon/dateparse) | 1296 | 21 | 2014/04/21 | 1 week ago | Parse date's without knowing format in advance. |
| [carbon](https://github.com/uniplaces/carbon) | 545 | 42 | 2016/08/03 | 17 hours ago | Simple Time extension with a lot of util methods, ported from PHP Carbon library. |
| [durafmt](https://github.com/hako/durafmt) | 342 | 5 | 2016/05/20 | 1 week ago | Time duration formatting library for Go. |
| [timeutil](https://github.com/leekchan/timeutil) | 183 | 8 | 2015/08/02 | 2 years ago | Useful extensions (Timedelta, Strftime, ...) to the golang's time package. |
| [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) | 92 | 4 | 2016/01/31 | 4 months ago | The implementation of the Persian (Solar Hijri) Calendar in Go (golang). |
| [iso8601](https://github.com/relvacode/iso8601) | 84 | 4 | 2017/04/25 | 8 months ago | Efficiently parse ISO8601 date-times without regex. |
| [timespan](https://github.com/SaidinWoT/timespan) | 74 | 6 | 2014/10/07 | 1 year ago | For interacting with intervals of time, defined as a start time and a duration. |
| [date](https://github.com/rickb777/date) | 59 | 3 | 2015/11/23 | 1 month ago | Augments Time for working with dates, date ranges, time spans, periods, and time-of-day. |
| [feiertage](https://github.com/wlbr/feiertage) | 35 | 2 | 2015/11/04 | 4 months ago | Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving... |
| [go-sunrise](https://github.com/nathan-osman/go-sunrise) | 31 | 4 | 2017/06/15 | 3 months ago | Calculate the sunrise and sunset times for a given location. |
| [go-str2duration](https://github.com/xhit/go-str2duration) | 19 | 2 | 2020/02/02 | 6 months ago | Convert string to duration. Support time.Duration returned string and more. |
| [kair](https://github.com/GuilhermeCaruso/kair) | 18 | 1 | 2018/10/03 | 8 months ago | Date and Time - Golang Formatting Library. |
| [nulltime](https://github.com/kirillDanshin/nulltime) | 11 | 2 | 2016/03/06 | 3 years ago | Nullable `time.Time`. |
| [tuesday](https://github.com/osteele/tuesday) | 9 | 3 | 2017/08/10 | 3 years ago | Ruby-compatible Strftime function. |
| [cronrange](https://github.com/1set/cronrange) | 8 | 1 | 2019/11/10 | 1 year ago | Parses Cron-style time range expressions, checks if the given time is within any ranges. |
| [strftime](https://github.com/awoodbeck/strftime) | 7 | 2 | 2018/02/10 | 3 years ago | C99-compatible strftime formatter. |
| [go-week](https://github.com/stoewer/go-week) | 4 | 4 | 2018/02/23 | 7 months ago | An efficient package to work with ISO8601 week dates. |

## Distributed Systems
        
*Packages that help with building Distributed Systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [kit](https://github.com/go-kit/kit) | 19353 | 687 | 2015/02/03 | 4 days ago | Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc. |
| [go-micro](https://github.com/asim/go-micro) | 15379 | 504 | 2015/01/13 | 4 days ago | A distributed systems development framework. |
| [nitro](https://github.com/gonitro/nitro) | 14946 | 498 | 2015/01/13 | 2 months ago | A distributed systems development framework. |
| [nitro](https://github.com/asim/nitro) | 14898 | 498 | 2015/01/13 | 2 months ago | A distributed systems development framework. |
| [go-micro](https://github.com/micro/go-micro) | 14444 | 485 | 2015/01/13 | 4 months ago | A distributed systems development framework. |
| [grpc-go](https://github.com/grpc/grpc-go) | 13120 | 494 | 2014/12/08 | 1 day ago | The Go language implementation of gRPC. HTTP/2 based RPC. |
| [micro](https://github.com/micro/micro) | 9700 | 339 | 2015/01/16 | 6 hours ago | A distributed systems runtime for the cloud and beyond. |
| [nats-server](https://github.com/nats-io/nats-server) | 8868 | 379 | 2012/10/29 | 18 hours ago | Lightweight, high performance messaging system for microservices, IoT, and cloud native systems. |
| [rpcx](https://github.com/smallnest/rpcx) | 5439 | 335 | 2016/05/18 | 5 days ago | Distributed pluggable RPC service framework like alibaba Dubbo. |
| [raft](https://github.com/hashicorp/raft) | 4316 | 327 | 2013/11/05 | 1 week ago | Golang implementation of the Raft consensus protocol, by HashiCorp. |
| [tendermint](https://github.com/tendermint/tendermint) | 3930 | 259 | 2014/05/14 | 1 day ago | High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols. |
| [torrent](https://github.com/anacrolix/torrent) | 3785 | 133 | 2015/01/08 | 3 days ago | BitTorrent client package. |
| [krakend](https://github.com/devopsfaith/krakend) | 3548 | 108 | 2016/11/04 | 2 days ago | Ultra performant API Gateway framework with middlewares. |
| [dragonboat](https://github.com/lni/dragonboat) | 3487 | 140 | 2018/12/23 | 1 week ago | A feature complete and high performance multi-group Raft library in Go. |
| [glow](https://github.com/chrislusf/glow) | 2926 | 148 | 2015/06/14 | 2 years ago | Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go. |
| [gleam](https://github.com/chrislusf/gleam) | 2739 | 149 | 2016/08/26 | 3 weeks ago | Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed. |
| [emitter](https://github.com/emitter-io/emitter) | 2738 | 100 | 2016/10/29 | 11 hours ago | High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love. |
| [liftbridge](https://github.com/liftbridge-io/liftbridge) | 1997 | 72 | 2017/10/13 | 1 day ago | Lightweight, fault-tolerant message streams for NATS. |
| [hprose-golang](https://github.com/hprose/hprose-golang) | 1133 | 91 | 2014/02/14 | 1 day ago | Very newbility RPC Library, support 25+ languages now. |
| [ringpop-go](https://github.com/uber/ringpop-go) | 661 | 2391 | 2015/06/05 | 1 week ago | Scalable, fault-tolerant application-layer sharding for Go applications. |
| [gorpc](https://github.com/valyala/gorpc) | 624 | 26 | 2014/11/20 | 1 year ago | Simple, fast and scalable RPC library for high load. |
| [go-health](https://github.com/InVisionApp/go-health) | 583 | 111 | 2017/11/29 | 1 year ago | Library for enabling asynchronous dependency health checks in your service. |
| [rain](https://github.com/cenkalti/rain) | 564 | 15 | 2014/05/21 | 1 month ago | BitTorrent client and library. |
| [digota](https://github.com/digota/digota) | 379 | 26 | 2017/08/14 | 2 years ago | grpc ecommerce microservice. |
| [consistent](https://github.com/buraksezer/consistent) | 358 | 14 | 2018/03/25 | 1 year ago | Consistent hashing with bounded loads. |
| [redislock](https://github.com/bsm/redislock) | 355 | 5 | 2019/06/24 | 3 months ago | Simplified distributed locking implementation using Redis. |
| [go-sundheit](https://github.com/AppsFlyer/go-sundheit) | 352 | 8 | 2019/04/08 | 5 days ago | A library built to provide support for defining async service health checks for golang services. |
| [sleuth](https://github.com/ursiform/sleuth) | 331 | 10 | 2016/04/23 | 2 years ago | Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)). |
| [go-jump](https://github.com/dgryski/go-jump) | 310 | 14 | 2014/06/15 | 3 years ago | Port of Google's "Jump" Consistent Hash function. |
| [dht](https://github.com/anacrolix/dht) | 182 | 13 | 2016/12/14 | 2 weeks ago | BitTorrent Kademlia DHT implementation. |
| [jsonrpc](https://github.com/ybbus/jsonrpc) | 164 | 9 | 2016/11/10 | 1 month ago | JSON-RPC 2.0 HTTP client implementation. |
| [arpc](https://github.com/lesismal/arpc) | 154 | 11 | 2020/05/19 | 7 hours ago | More effective network communication, support two-way-calling, notify, broadcast. |
| [jsonrpc](https://github.com/osamingo/jsonrpc) | 141 | 5 | 2016/10/28 | 2 months ago | The jsonrpc package helps implement of JSON-RPC 2.0. |
| [celeriac.v1](https://github.com/svcavallar/celeriac.v1) | 62 | 4 | 2015/10/10 | 4 months ago | Library for adding support for interacting and monitoring Celery workers, tasks and events in Go. |
| [doublejump](https://github.com/edwingeng/doublejump) | 57 | 4 | 2018/06/26 | 10 months ago | A revamped Google's jump consistent hash. |
| [semaphore](https://github.com/jexia/semaphore) | 50 | 15 | 2020/02/05 | 2 months ago | A straightforward (micro) service orchestrator. |
| [outboxer](https://github.com/italolelis/outboxer) | 47 | 0 | 2019/02/01 | 2 days ago | Outboxer is a go library that implements the outbox pattern. |
| [flowgraph](https://github.com/vectaport/flowgraph) | 38 | 1 | 2018/08/29 | 7 months ago | flow-based programming package. |
| [maestro](https://github.com/jexia/maestro) | 34 | 16 | 2020/02/05 | 7 months ago | A straightforward (micro) service orchestrator. |
| [drmaa](https://github.com/dgruber/drmaa) | 29 | 2 | 2013/03/17 | 4 months ago | Job submission library for cluster schedulers based on the DRMAA standard. |
| [go-pdu](https://github.com/pdupub/go-pdu) | 24 | 3 | 2018/10/08 | 4 days ago | A decentralized identity-based social network. |
| [go-mysql-lock](https://github.com/sanketplus/go-mysql-lock) | 19 | 1 | 2020/06/06 | 3 days ago | MySQL based distributed lock. |
| [dynatomic](https://github.com/tylfin/dynatomic) | 14 | 1 | 2019/02/08 | 3 months ago | A library for using DynamoDB as an atomic counter. |
| [micro](https://github.com/gmsec/micro) | 9 | 1 | 2020/05/03 | 2 months ago | A Go distributed systems development framework. |
| [consistenthash](https://github.com/mbrostami/consistenthash) | 5 | 1 | 2020/04/22 | 9 months ago | Consistent hashing with configurable replicas. |

## Dynamic DNS
        
*Tools for updating dynamic DNS records.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [godns](https://github.com/TimothyYe/godns) | 768 | 34 | 2014/05/11 | 1 week ago | A dynamic DNS client tool, supports DNSPod & HE.net, written in Go. |
| [ddns](https://github.com/skibish/ddns) | 175 | 7 | 2017/03/13 | 3 weeks ago | Personal DDNS client with Digital Ocean Networking DNS as backend. |

## Email
        
*Libraries and tools that implement email creation and sending.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [MailHog](https://github.com/mailhog/MailHog) | 7900 | 137 | 2014/04/16 | 1 month ago | Email and SMTP testing with web and API interface. |
| [hermes](https://github.com/matcornic/hermes) | 2181 | 29 | 2017/03/25 | 2 months ago | Golang package that generates clean, responsive HTML e-mails. |
| [email](https://github.com/jordan-wright/email) | 1645 | 49 | 2013/12/12 | 1 month ago | A robust and flexible email library for Go. |
| [go-imap](https://github.com/emersion/go-imap) | 1195 | 38 | 2016/04/26 | 2 days ago | IMAP library for clients and servers. |
| [sendgrid-go](https://github.com/sendgrid/sendgrid-go) | 704 | 192 | 2013/09/12 | 1 day ago | SendGrid's Go library for sending email. |
| [mailgun-go](https://github.com/mailgun/mailgun-go) | 499 | 58 | 2014/02/28 | 1 day ago | Go library for sending mail with the Mailgun API. |
| [hectane](https://github.com/hectane/hectane) | 197 | 12 | 2015/08/28 | 2 months ago | Lightweight SMTP client providing an HTTP API. |
| [douceur](https://github.com/aymerick/douceur) | 189 | 3 | 2015/04/09 | 2 years ago | CSS inliner for your HTML emails. |
| [go-message](https://github.com/emersion/go-message) | 189 | 12 | 2016/12/31 | 3 weeks ago | Streaming library for the Internet Message Format and mail messages. |
| [email-verifier](https://github.com/AfterShip/email-verifier) | 154 | 20 | 2020/12/18 | 2 weeks ago | A Go library for email verification without sending any emails. |
| [go-simple-mail](https://github.com/xhit/go-simple-mail) | 100 | 2 | 2019/09/15 | 1 week ago | Very simple package to send emails with SMTP Keep Alive and two timeouts: Connect and Send. |
| [mailchain](https://github.com/mailchain/mailchain) | 71 | 6 | 2019/04/11 | 17 hours ago | Send encrypted emails to blockchain addresses written in Go. |
| [go-premailer](https://github.com/vanng822/go-premailer) | 68 | 2 | 2015/02/16 | 5 months ago | Inline styling for HTML mail in Go. |
| [go-dkim](https://github.com/toorop/go-dkim) | 65 | 2 | 2015/04/29 | 3 months ago | DKIM library, to sign & verify email. |
| [smtp](https://github.com/mailhog/smtp) | 59 | 10 | 2014/12/24 | 6 months ago | SMTP server protocol state machine. |
| [go-email-validator](https://github.com/go-email-validator/go-email-validator) | 8 | 3 | 2020/12/10 | 1 week ago | Modular email validator for syntax, disposable, smtp, etc... checking. |

## Embeddable Scripting Languages
        
*Embedding other languages inside your go code.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [otto](https://github.com/robertkrimen/otto) | 5508 | 183 | 2012/10/06 | 5 months ago | JavaScript interpreter written in Go. |
| [gopher-lua](https://github.com/yuin/gopher-lua) | 3932 | 151 | 2015/02/15 | 3 months ago | Lua 5.1 VM and compiler written in Go. |
| [tengo](https://github.com/d5/tengo) | 2161 | 47 | 2019/01/09 | 3 months ago | Bytecode compiled script language for Go. |
| [go-lua](https://github.com/Shopify/go-lua) | 2018 | 282 | 2013/12/20 | 2 months ago | Port of the Lua 5.2 VM to pure Go. |
| [goja](https://github.com/dop251/goja) | 1980 | 64 | 2016/11/04 | 4 days ago | ECMAScript 5.1(+) implementation in Go. |
| [expr](https://github.com/antonmedv/expr) | 1624 | 38 | 2018/07/14 | 1 week ago | Expression evaluation engine for Go: fast, non-Turing complete, dynamic typing, static typing. |
| [go-python](https://github.com/sbinet/go-python) | 1230 | 43 | 2012/07/09 | 10 months ago | naive go bindings to the CPython C-API. |
| [anko](https://github.com/mattn/anko) | 1113 | 47 | 2014/03/28 | 1 month ago | Scriptable interpreter written in Go. |
| [go-php](https://github.com/deuill/go-php) | 779 | 43 | 2015/09/17 | 2 years ago | PHP bindings for Go. |
| [go-duktape](https://github.com/olebedev/go-duktape) | 741 | 29 | 2015/01/08 | 8 months ago | Duktape JavaScript engine bindings for Go. |
| [cel-go](https://github.com/google/cel-go) | 693 | 28 | 2018/03/09 | 2 days ago | Fast, portable, non-Turing complete expression evaluation with gradual typing. |
| [golua](https://github.com/aarzilli/golua) | 514 | 32 | 2010/12/06 | 1 month ago | Go bindings for Lua C API. |
| [gisp](https://github.com/jcla1/gisp) | 445 | 21 | 2014/01/11 | 3 years ago | Simple LISP in Go. |
| [gval](https://github.com/PaesslerAG/gval) | 285 | 15 | 2017/09/27 | 1 month ago | A highly customizable expression language written in Go. |
| [gentee](https://github.com/gentee/gentee) | 70 | 3 | 2018/01/14 | 1 month ago | Embeddable scripting programming language. |
| [binder](https://github.com/alexeyco/binder) | 48 | 2 | 2017/04/02 | 2 years ago | Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua). |
| [purl](https://github.com/ian-kent/purl) | 30 | 3 | 2014/11/29 | 6 years ago | Perl 5.18.2 embedded in Go. |
| [ngaro](https://github.com/db47h/ngaro) | 20 | 1 | 2016/08/09 | 2 years ago | Embeddable Ngaro VM implementation enabling scripting in Retro. |
| [ecal](https://github.com/krotik/ecal) | 6 | 1 | 2020/11/30 | 1 month ago | A simple embeddable scripting language which supports concurrent event processing. |

## Error Handling
        
*Libraries for handling errors.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [errors](https://github.com/pkg/errors) | 6537 | 113 | 2015/12/27 | 2 months ago | Package that provides simple error handling primitives. |
| [go-multierror](https://github.com/hashicorp/go-multierror) | 1130 | 220 | 2014/12/15 | 2 weeks ago | Go (golang) package for representing a list of errors as a single error. |
| [eris](https://github.com/rotisserie/eris) | 745 | 10 | 2019/09/07 | 2 weeks ago | A better way to handle, trace, and log errors in Go. Compatible with the standard error library and github.com/pkg/errors. |
| [errorx](https://github.com/joomcode/errorx) | 663 | 70 | 2018/08/17 | 5 months ago | A feature rich error package with stack traces, composition of errors and more. |
| [tracerr](https://github.com/ztrue/tracerr) | 642 | 9 | 2019/02/06 | 1 year ago | Golang errors with stack trace and source fragments. |
| [errlog](https://github.com/snwfdhmp/errlog) | 391 | 6 | 2019/02/16 | 2 months ago | Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place. |
| [emperror](https://github.com/emperror/emperror) | 198 | 4 | 2017/06/13 | 4 months ago | Error handling tools and best practices for Go libraries and applications. |
| [errors](https://github.com/emperror/errors) | 88 | 3 | 2019/07/09 | 4 months ago | Drop-in replacement for the standard library errors package and github.com/pkg/errors. Provides various error handling primitives. |
| [errors](https://github.com/bnkamalesh/errors) | 21 | 1 | 2020/07/17 | 1 day ago | Drop-in replacement for builting Go errors. This is a minimal error handling package with custom error types, user friendly messages, Unwrap & Is. With very easy to use and straightforward helper functions. |
| [werr](https://github.com/txgruppi/werr) | 13 | 0 | 2015/08/04 | 5 years ago | Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called. |
| [falcon](https://github.com/SonicRoshan/falcon) | 6 | 2 | 2019/09/09 | 1 year ago | A Simple Yet Highly Powerful Package For Error Handling. |
| [errors](https://github.com/neuronlabs/errors) | 3 | 1 | 2019/07/26 | 1 year ago | Simple golang error handling with classification primitives. |
| [errors](https://github.com/PumpkinSeed/errors) | 2 | 1 | 2020/01/08 | 1 year ago | The most simple error wrapper with awesome performance and minimal memory overhead. |

## Files
        
*Libraries for handling files and file systems.*

## Financial
        
*Packages for accounting and finance.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [decimal](https://github.com/shopspring/decimal) | 2784 | 67 | 2015/02/25 | 1 hour ago | Arbitrary-precision fixed-point decimal numbers. |
| [go-money](https://github.com/Rhymond/go-money) | 873 | 16 | 2017/03/20 | 1 week ago | Implementation of Fowler's Money pattern. |
| [accounting](https://github.com/leekchan/accounting) | 624 | 14 | 2015/08/10 | 1 week ago | money and currency formatting for golang. |
| [go-finance](https://github.com/FlashBoys/go-finance) | 535 | 26 | 2016/02/28 | 2 years ago | Comprehensive financial markets data in Go. |
| [techan](https://github.com/sdcoffey/techan) | 374 | 31 | 2017/03/08 | 2 months ago | Technical analysis library with advanced market analysis and trading strategies. |
| [currency](https://github.com/bojanz/currency) | 231 | 6 | 2020/04/16 | 2 days ago | Handles currency amounts, provides currency information and formatting. |
| [orderbook](https://github.com/i25959341/orderbook) | 178 | 15 | 2018/04/24 | 1 year ago | Matching Engine for Limit Order Book in Golang. |
| [go-finance](https://github.com/alpeb/go-finance) | 90 | 6 | 2017/06/01 | 2 months ago | Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations. |
| [transaction](https://github.com/claygod/transaction) | 83 | 9 | 2017/10/11 | 7 months ago | Embedded transactional database of accounts, running in multithreaded mode. |
| [ofxgo](https://github.com/aclindsa/ofxgo) | 81 | 10 | 2015/11/08 | 1 month ago | Query OFX servers and/or parse the responses (with example command-line client). |
| [vat](https://github.com/dannyvankooten/vat) | 81 | 2 | 2016/06/18 | 1 week ago | VAT number validation & EU VAT rates. |
| [go-finnhub](https://github.com/m1/go-finnhub) | 50 | 5 | 2020/01/13 | 1 year ago | Client for stock market, forex and crypto data from finnhub.io. Access real-time financial market data from 60+ stock exchanges, 10 forex brokers, and 15+ crypto exchanges. |
| [currency](https://github.com/bnkamalesh/currency) | 38 | 6 | 2017/05/09 | 8 months ago | High performant & accurate currency computation package. |
| [sleet](https://github.com/BoltApp/sleet) | 30 | 22 | 2019/11/13 | 5 days ago | One unified interface for multiple Payment Service Providers (PsP) to process online payment. |
| [fastme](https://github.com/newity/fastme) | 19 | 3 | 2020/10/29 | 1 week ago | Fast extensible matching engine Go implementation. |
| [go-finance](https://github.com/pieterclaerhout/go-finance) | 4 | 1 | 2019/09/30 | 1 year ago | Module to fetch exchange rates, check VAT numbers via VIES and check IBAN bank account numbers. |

## Forms
        
*Libraries for working with forms.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [nosurf](https://github.com/justinas/nosurf) | 1122 | 34 | 2013/08/22 | 3 months ago | CSRF protection middleware for Go. |
| [binding](https://github.com/mholt/binding) | 782 | 32 | 2014/05/20 | 2 years ago | Binds form and JSON data from net/http Request to struct. |
| [csrf](https://github.com/gorilla/csrf) | 618 | 21 | 2015/08/03 | 1 week ago | CSRF protection for Go web applications & services. |
| [form](https://github.com/go-playground/form) | 446 | 12 | 2016/05/26 | 1 week ago | Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support. |
| [conform](https://github.com/leebenson/conform) | 213 | 6 | 2016/01/05 | 8 months ago | Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags. |
| [formam](https://github.com/monoculum/formam) | 156 | 4 | 2014/10/25 | 2 weeks ago | decode form's values into a struct. |
| [forms](https://github.com/albrow/forms) | 117 | 7 | 2014/08/07 | 3 years ago | Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files. |
| [qs](https://github.com/sonh/qs) | 55 | 3 | 2020/10/02 | 3 months ago | Go module for encoding structs into URL query parameters. |
| [bind](https://github.com/robfig/bind) | 25 | 3 | 2014/08/06 | 6 years ago | Bind form data to any Go values. |
| [queryparam](https://github.com/TomWright/queryparam) | 7 | 2 | 2018/06/14 | 4 months ago | Decode `url.Values` into usable struct values of standard or custom types. |

## Functional
        
*Packages to support functional programming in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-underscore](https://github.com/tobyhede/go-underscore) | 1184 | 30 | 2014/07/02 | 2 years ago | Useful collection of helpfully functional Go collection utilities. |
| [fpGo](https://github.com/TeaEntityLab/fpGo) | 162 | 5 | 2018/05/24 | 2 years ago | Monad, Functional Programming features for Golang. |
| [fuego](https://github.com/seborama/fuego) | 83 | 3 | 2018/11/05 | 3 months ago | Functional Experiment in Go. |

## Game Development
        
*Awesome game development libraries.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [ebiten](https://github.com/hajimehoshi/ebiten) | 4075 | 114 | 2013/06/16 | 4 hours ago | dead simple 2D game library in Go. |
| [leaf](https://github.com/name5566/leaf) | 3851 | 320 | 2014/08/04 | 9 months ago | Lightweight game server framework. |
| [pixel](https://github.com/faiface/pixel) | 3357 | 102 | 2016/11/19 | 2 weeks ago | Hand-crafted 2D game library in Go. |
| [goworld](https://github.com/xiaonanln/goworld) | 1749 | 128 | 2017/06/03 | 5 months ago | Scalable game server engine, featuring space-entity framework and hot-swapping. |
| [nano](https://github.com/lonng/nano) | 1580 | 65 | 2017/08/02 | 2 months ago | Lightweight, facility, high performance golang based game server framework. |
| [go-sdl2](https://github.com/veandco/go-sdl2) | 1492 | 46 | 2013/06/05 | 10 hours ago | Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/). |
| [engo](https://github.com/EngoEngine/engo) | 1341 | 48 | 2014/11/12 | 23 hours ago | Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm. |
| [engine](https://github.com/g3n/engine) | 1320 | 73 | 2017/03/07 | 5 days ago | Go 3D Game Engine. |
| [termloop](https://github.com/JoelOtter/termloop) | 1185 | 32 | 2015/05/23 | 2 months ago | Terminal-based game engine for Go, built on top of Termbox. |
| [gonet](https://github.com/xtaci/gonet) | 1122 | 136 | 2013/04/11 | 3 years ago | Game server skeleton implemented with golang. |
| [pitaya](https://github.com/topfreegames/pitaya) | 881 | 65 | 2018/03/19 | 3 weeks ago | Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. |
| [oak](https://github.com/oakmound/oak) | 834 | 46 | 2017/07/15 | 1 week ago | Pure Go game engine. |
| [raylib-go](https://github.com/gen2brain/raylib-go) | 568 | 21 | 2017/01/27 | 1 week ago | Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming. |
| [engine](https://github.com/azul3d/engine) | 472 | 25 | 2016/02/29 | 1 year ago | 3D game engine written in Go. |
| [go-astar](https://github.com/beefsack/go-astar) | 409 | 10 | 2014/05/28 | 5 months ago | Go implementation of the A\* path finding algorithm. |
| [GarageEngine](https://github.com/vova616/GarageEngine) | 320 | 31 | 2012/08/07 | 1 year ago | 2d game engine written in Go working on OpenGL. |
| [go3d](https://github.com/ungerik/go3d) | 191 | 10 | 2011/06/27 | 1 year ago | Performance oriented 2D/3D math package for Go. |
| [glop](https://github.com/runningwild/glop) | 77 | 3 | 2011/04/20 | 5 years ago | Glop (Game Library Of Power) is a fairly simple cross-platform game library. |
| [prototype](https://github.com/gonutz/prototype) | 42 | 3 | 2015/03/04 | 4 months ago | Cross-platform (Windows/Linux/Mac) library for creating desktop games using a minimal API. |
| [go-collada](https://github.com/GlenKelley/go-collada) | 15 | 3 | 2013/09/19 | 7 years ago | Go package for working with the Collada file format. |
| [tile](https://github.com/kelindar/tile) | 15 | 1 | 2020/08/19 | 5 months ago | Data-oriented and cache-friendly 2D Grid library (TileMap), includes pathfinding, observers and import/export. |

## Generation and Generics
        
*Tools to enhance the language with features like generics via code generation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-linq](https://github.com/ahmetb/go-linq) | 2393 | 73 | 2013/12/19 | 1 month ago | .NET LINQ-like query methods for Go. |
| [jennifer](https://github.com/dave/jennifer) | 1841 | 27 | 2016/12/04 | 4 weeks ago | Generate arbitrary Go code without templates. |
| [gen](https://github.com/clipperhouse/gen) | 1235 | 33 | 2013/10/13 | 1 year ago | Code generation tool for ‘generics’-like functionality. |
| [goderive](https://github.com/awalterschulze/goderive) | 841 | 24 | 2017/02/10 | 3 months ago | Derives functions from input types. |
| [gowrap](https://github.com/hexdigest/gowrap) | 446 | 11 | 2018/09/15 | 1 month ago | Generate decorators for Go interfaces using simple templates. |
| [interfaces](https://github.com/rjeczalik/interfaces) | 276 | 7 | 2015/12/06 | 1 month ago | Command line tool for generating interface definitions. |
| [go-enum](https://github.com/abice/go-enum) | 165 | 3 | 2017/08/10 | 1 month ago | Code generation for enums from code comments. |
| [pkgreflect](https://github.com/ungerik/pkgreflect) | 96 | 6 | 2012/09/03 | 3 years ago | Go preprocessor for package scoped reflection. |
| [efaceconv](https://github.com/t0pep0/efaceconv) | 47 | 4 | 2016/11/18 | 3 years ago | Code generation tool for high performance conversion from interface{} to immutable type without allocations. |
| [gotype](https://github.com/wzshiming/gotype) | 34 | 3 | 2017/12/05 | 9 months ago | Golang source code parsing, usage like reflect package. |
| [GENERIS](https://github.com/senselogic/GENERIS) | 23 | 1 | 2019/03/10 | 4 months ago | Code generation tool providing generics, free-form macros, conditional compilation and HTML templating. |
| [go-xray](https://github.com/pieterclaerhout/go-xray) | 16 | 2 | 2019/10/01 | 1 year ago | Helpers for making the use of reflection easier. |
| [typeregistry](https://github.com/xiaoxin01/typeregistry) | 9 | 1 | 2020/01/14 | 1 year ago | A library to create type dynamically. |

## Geographic
        
*Geographic tools and servers*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [tile38](https://github.com/tidwall/tile38) | 7258 | 204 | 2016/03/04 | 6 days ago | Geolocation DB with spatial index and realtime geofencing. |
| [geo](https://github.com/golang/geo) | 1164 | 79 | 2014/12/03 | 2 days ago | S2 geometry library in Go. |
| [mbtileserver](https://github.com/consbio/mbtileserver) | 213 | 15 | 2014/11/01 | 1 month ago | A simple Go-based server for map tiles stored in mbtiles format. |
| [osm](https://github.com/paulmach/osm) | 151 | 11 | 2016/02/02 | 1 week ago | Library for reading, writing and working with OpenStreetMap data and APIs. |
| [geocache](https://github.com/melihmucuk/geocache) | 128 | 6 | 2016/06/21 | 4 years ago | In-memory cache that is suitable for geolocation based applications. |
| [wgs84](https://github.com/wroge/wgs84) | 59 | 1 | 2019/06/08 | 2 months ago | Library for Coordinate Conversion and Transformation (ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM). |
| [geoserver](https://github.com/hishamkaram/geoserver) | 45 | 2 | 2018/03/26 | 3 days ago | geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API. |
| [gismanager](https://github.com/hishamkaram/gismanager) | 37 | 1 | 2018/09/29 | 2 years ago | Publish Your GIS Data(Vector Data) to PostGIS and Geoserver. |
| [pbf](https://github.com/maguro/pbf) | 24 | 2 | 2017/09/18 | 4 days ago | OpenStreetMap PBF golang encoder/decoder. |
| [s2-geojson](https://github.com/pantrif/s2-geojson) | 10 | 1 | 2020/03/27 | 10 months ago | Convert geojson to s2 cells & demonstrating some S2 geometry features on map. |

## Go Compilers
        
*Tools for compiling Go to other languages.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gopherjs](https://github.com/gopherjs/gopherjs) | 9936 | 267 | 2013/08/27 | 1 week ago | Compiler from Go to JavaScript. |
| [llgo](https://github.com/go-llvm/llgo) | 1095 | 63 | 2011/11/05 | 6 years ago | LLVM-based compiler for Go. |
| [tardisgo](https://github.com/tardisgo/tardisgo) | 405 | 28 | 2014/01/08 | 4 years ago | Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler. |
| [c4go](https://github.com/Konstantin8105/c4go) | 248 | 18 | 2018/03/28 | 6 months ago | Transpile C code to Go code. |
| [f4go](https://github.com/Konstantin8105/f4go) | 23 | 3 | 2018/07/08 | 1 month ago | Transpile FORTRAN 77 code to Go code. |

## Goroutines
        
*Tools for managing and working with Goroutines.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [ants](https://github.com/panjf2000/ants) | 4972 | 150 | 2018/05/19 | 2 weeks ago | A high-performance and low-cost goroutine pool in Go. |
| [goworker](https://github.com/benmanns/goworker) | 2513 | 74 | 2013/07/22 | 1 week ago | goworker is a Go-based background worker. |
| [tunny](https://github.com/Jeffail/tunny) | 1858 | 68 | 2014/04/02 | 2 weeks ago | Goroutine pool for golang. |
| [grpool](https://github.com/ivpusic/grpool) | 608 | 30 | 2015/07/22 | 2 years ago | Lightweight Goroutine pool. |
| [pool](https://github.com/go-playground/pool) | 589 | 14 | 2015/10/28 | 1 year ago | Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation. |
| [workerpool](https://github.com/gammazero/workerpool) | 459 | 15 | 2016/05/17 | 3 months ago | Goroutine pool that limits the concurrency of task execution, not the number of tasks queued. |
| [gowp](https://github.com/xxjwxc/gowp) | 256 | 14 | 2019/09/14 | 8 months ago | gowp is concurrency limiting goroutine pool. |
| [go-floc](https://github.com/workanator/go-floc) | 184 | 7 | 2017/07/03 | 1 month ago | Orchestrate goroutines with ease. |
| [pond](https://github.com/alitto/pond) | 179 | 5 | 2020/03/21 | 1 month ago | Minimalistic and High-performance goroutine worker pool written in Go. |
| [go-flow](https://github.com/kamildrazkiewicz/go-flow) | 152 | 9 | 2016/09/25 | 1 year ago | Control goroutines execution order. |
| [semaphore](https://github.com/marusama/semaphore) | 106 | 2 | 2017/11/22 | 8 months ago | Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations). |
| [artifex](https://github.com/mborders/artifex) | 97 | 6 | 2018/10/31 | 5 months ago | Simple in-memory job queue for Golang using worker-based dispatching. |
| [GoSlaves](https://github.com/dgrr/GoSlaves) | 95 | 4 | 2017/09/17 | 1 year ago | Simple and Asynchronous Goroutine pool library. |
| [breaker](https://github.com/kamilsk/breaker) | 92 | 3 | 2019/02/15 | 2 weeks ago | Flexible mechanism to make execution flow interruptible. |
| [go-workers](https://github.com/catmullet/go-workers) | 85 | 3 | 2020/10/06 | 2 days ago | Easily and safely run workers for large data processing pipelines. |
| [semaphore](https://github.com/kamilsk/semaphore) | 83 | 1 | 2016/10/08 | 10 months ago | Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context. |
| [worker-pool](https://github.com/vardius/worker-pool) | 73 | 5 | 2017/10/04 | 4 weeks ago | goworker is a Go simple async worker pool. |
| [gpool](https://github.com/sherifabdlnaby/gpool) | 72 | 1 | 2018/12/03 | 1 year ago | manages a resizeable pool of context-aware goroutines to bound concurrency. |
| [async](https://github.com/StudioSol/async) | 71 | 12 | 2017/06/30 | 2 months ago | A safe way to execute functions asynchronously, recovering them in case of panic. |
| [errgroup](https://github.com/neilotoole/errgroup) | 66 | 1 | 2020/06/26 | 6 months ago | Drop-in alternative to `sync/errgroup`, limited to a pool of N worker goroutines. |
| [cyclicbarrier](https://github.com/marusama/cyclicbarrier) | 59 | 2 | 2018/01/11 | 7 months ago | CyclicBarrier for golang. |
| [threadpool](https://github.com/shettyh/threadpool) | 53 | 3 | 2017/09/06 | 10 months ago | Golang threadpool implementation. |
| [gollback](https://github.com/vardius/gollback) | 51 | 1 | 2019/05/11 | 8 months ago | asynchronous simple function utilities, for managing execution of closures and callbacks. |
| [Hunch](https://github.com/AaronJan/Hunch) | 42 | 1 | 2019/06/05 | 4 months ago | Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive. |
| [routine](https://github.com/x-mod/routine) | 39 | 2 | 2019/03/04 | 4 months ago | go routine control with context, support: Main, Go, Pool and some useful Executors. |
| [kyoo](https://github.com/dirkaholic/kyoo) | 32 | 2 | 2020/01/06 | 10 months ago | Provides an unlimited job queue and concurrent worker pools. |
| [parallel-fn](https://github.com/rafaeljesus/parallel-fn) | 29 | 3 | 2017/06/18 | 3 years ago | Run functions in parallel. |
| [nursery](https://github.com/arunsworld/nursery) | 23 | 4 | 2019/11/23 | 2 weeks ago | Structured concurrency in Go. |
| [async](https://github.com/reugn/async) | 20 | 1 | 2019/12/28 | 5 months ago | An alternative sync library for Go (Future, Promise, Locks). |
| [stl](https://github.com/ssgreg/stl) | 18 | 2 | 2018/06/19 | 6 months ago | Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism. |
| [goccm](https://github.com/zenthangplus/goccm) | 18 | 1 | 2019/08/16 | 8 months ago | Go Concurrency Manager package limits the number of goroutines that allowed to run concurrently. |
| [go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup) | 16 | 1 | 2018/08/08 | 11 months ago | Like `sync.WaitGroup` with error handling and concurrency control. |
| [go-trylock](https://github.com/subchen/go-trylock) | 13 | 1 | 2018/04/26 | 7 months ago | TryLock support on read-write lock for Golang. |
| [gohive](https://github.com/loveleshsharma/gohive) | 13 | 3 | 2019/05/31 | 1 year ago | A highly performant and easy to use Goroutine pool for Go. |
| [conexec](https://github.com/ITcathyh/conexec) | 10 | 2 | 2019/12/24 | 7 months ago | A concurrent toolkit to help execute funcs concurrently in an efficient and safe way. It supports specifying the overall timeout to avoid blocking and uses goroutine pool to improve efficiency. |
| [queue](https://github.com/AnikHasibul/queue) | 9 | 0 | 2018/12/21 | 1 year ago | Gives you a `sync.WaitGroup` like queue group accessibility. Helps you to throttle and limit goroutines, wait for the end of the all goroutines and much more. |
| [channelify](https://github.com/ddelizia/channelify) | 7 | 1 | 2020/10/05 | 4 months ago | Transform your function to return channels for easy and powerful parallel processing. |
| [hands](https://github.com/duanckham/hands) | 6 | 1 | 2020/04/04 | 10 months ago | A process controller used to control the execution and return strategies of multiple goroutines. |
| [go-tools](https://github.com/nikhilsaraf/go-tools) | 5 | 2 | 2018/11/14 | 1 year ago | Manage a pool of goroutines using this lightweight library with a simple API. |
| [concurrency-limiter](https://github.com/vivek-ng/concurrency-limiter) | 4 | 1 | 2020/11/22 | 2 months ago | Concurrency limiter with support for timeouts , dynamic priority and context cancellation of goroutines. |

## GUI
        
*Interaction*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [fyne](https://github.com/fyne-io/fyne) | 12329 | 221 | 2018/02/04 | 15 hours ago | Cross platform native GUIs designed for Go based on Material Design. Supports: Linux, macOS, Windows, BSD, iOS and Android. |
| [qt](https://github.com/therecipe/qt) | 8285 | 313 | 2014/11/19 | 1 month ago | Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi). |
| [ui](https://github.com/andlabs/ui) | 7793 | 378 | 2014/02/17 | 3 days ago | Platform-native GUI library for Go. Cross platform. |
| [webview](https://github.com/webview/webview) | 7631 | 222 | 2017/08/19 | 5 days ago | Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux). |
| [robotgo](https://github.com/go-vgo/robotgo) | 6446 | 231 | 2016/09/26 | 3 days ago | Go Native cross-platform GUI system automation. Control the mouse, keyboard and other. |
| [webview](https://github.com/zserge/webview) | 6166 | 211 | 2017/08/19 | 8 months ago | Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux). |
| [walk](https://github.com/lxn/walk) | 5220 | 262 | 2010/09/16 | 2 days ago | Windows application library kit for Go. |
| [go-app](https://github.com/maxence-charriere/go-app) | 4404 | 136 | 2016/10/12 | 6 hours ago | Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. |
| [go-astilectron](https://github.com/asticode/go-astilectron) | 3646 | 138 | 2017/04/22 | 1 week ago | Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron). |
| [app](https://github.com/maxence-charriere/app) | 3314 | 111 | 2016/10/12 | 11 months ago | Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. |
| [go-sciter](https://github.com/sciter-sdk/go-sciter) | 1984 | 121 | 2015/10/15 | 4 days ago | Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform. |
| [systray](https://github.com/getlantern/systray) | 1651 | 57 | 2014/11/12 | 2 months ago | Cross platform Go library to place an icon and menu in the notification area. |
| [gotk3](https://github.com/gotk3/gotk3) | 1401 | 63 | 2015/08/13 | 22 hours ago | Go bindings for GTK3. |
| [gosx-notifier](https://github.com/deckarep/gosx-notifier) | 531 | 15 | 2013/11/25 | 11 months ago | OSX Desktop Notifications library for Go. |
| [gowd](https://github.com/dtylman/gowd) | 303 | 28 | 2017/03/29 | 1 year ago | Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform. |
| [trayhost](https://github.com/shurcooL/trayhost) | 201 | 6 | 2014/04/25 | 1 year ago | Cross-platform Go library to place an icon in the host operating system's taskbar. |
| [go-appindicator](https://github.com/dawidd6/go-appindicator) | 14 | 5 | 2019/05/04 | 1 month ago | Go bindings for libappindicator3 C library. |
| [activity-tracker](https://github.com/prashantgupta24/activity-tracker) | 9 | 2 | 2019/03/12 | 1 year ago | OSX library to notify about any (pluggable) activity on your machine. |
| [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) | 7 | 2 | 2019/03/30 | 1 year ago | OSX Sleep/Wake notifications in golang. |

## Hardware
        
*Libraries, tools, and tutorials for interacting with hardware.*

## Images
        
*Libraries for manipulating images.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gocv](https://github.com/hybridgroup/gocv) | 3809 | 135 | 2017/09/18 | 6 days ago | Go package for computer vision using OpenCV 3.3+. |
| [imaging](https://github.com/disintegration/imaging) | 3520 | 76 | 2012/12/06 | 1 month ago | Simple Go image processing package. |
| [imaginary](https://github.com/h2non/imaginary) | 3466 | 75 | 2015/03/04 | 2 days ago | Fast and simple HTTP microservice for image resizing. |
| [bild](https://github.com/anthonynsimon/bild) | 3026 | 64 | 2016/08/01 | 1 day ago | Collection of image processing algorithms in pure Go. |
| [ln](https://github.com/fogleman/ln) | 2837 | 90 | 2016/01/10 | 1 year ago | 3D line art rendering in Go. |
| [gg](https://github.com/fogleman/gg) | 2649 | 85 | 2016/02/18 | 1 week ago | 2D rendering in pure Go. |
| [resize](https://github.com/nfnt/resize) | 2578 | 81 | 2012/08/02 | 2 months ago | Image resizing for Go with common interpolation methods. |
| [pt](https://github.com/fogleman/pt) | 1916 | 58 | 2015/01/23 | 1 year ago | Path tracing engine written in Go. |
| [svgo](https://github.com/ajstarks/svgo) | 1598 | 49 | 2010/03/05 | 6 months ago | Go Language Library for SVG generation. |
| [smartcrop](https://github.com/muesli/smartcrop) | 1494 | 32 | 2014/04/07 | 9 months ago | Finds good crops for arbitrary images and crop sizes. |
| [gift](https://github.com/disintegration/gift) | 1412 | 51 | 2014/07/12 | 2 months ago | Package of image processing filters. |
| [picfit](https://github.com/thoas/picfit) | 1411 | 52 | 2014/12/06 | 1 month ago | An image resizing server written in Go. |
| [bimg](https://github.com/h2non/bimg) | 1348 | 37 | 2015/03/17 | 2 months ago | Small package for fast and efficient image processing using libvips. |
| [imagick](https://github.com/gographics/imagick) | 1260 | 52 | 2013/04/30 | 2 months ago | Go binding to ImageMagick's MagickWand C API. |
| [go-opencv](https://github.com/go-opencv/go-opencv) | 1225 | 65 | 2013/12/09 | 1 year ago | Go bindings for OpenCV. |
| [geopattern](https://github.com/pravj/geopattern) | 1109 | 21 | 2014/10/22 | 2 years ago | Create beautiful generative image patterns from a string. |
| [stegify](https://github.com/DimitarPetrov/stegify) | 922 | 21 | 2018/11/29 | 7 months ago | Go tool for LSB steganography, capable of hiding any file within an image. |
| [canvas](https://github.com/tdewolff/canvas) | 614 | 16 | 2017/05/20 | 4 days ago | Vector graphics to PDF, SVG or rasterized image. |
| [image2ascii](https://github.com/qeesung/image2ascii) | 483 | 9 | 2018/10/20 | 2 years ago | Convert image to ASCII. |
| [draft](https://github.com/lucasepe/draft) | 477 | 14 | 2020/06/05 | 1 month ago | Generate High Level Microservice Architecture diagrams for GraphViz using simple YAML syntax. |
| [govips](https://github.com/davidbyttow/govips) | 429 | 9 | 2016/12/25 | 2 days ago | A lightning fast image processing and resizing library for Go. |
| [mort](https://github.com/aldor007/mort) | 418 | 20 | 2017/11/19 | 1 month ago | Storage and image processing server written in Go. |
| [govatar](https://github.com/o1egl/govatar) | 410 | 8 | 2016/01/18 | 1 year ago | Library and CMD tool for generating funny avatars. |
| [goimagehash](https://github.com/corona10/goimagehash) | 384 | 10 | 2017/07/28 | 1 month ago | Go Perceptual image hashing package. |
| [go-nude](https://github.com/koyachi/go-nude) | 321 | 16 | 2014/05/02 | 2 years ago | Nudity detection with Go. |
| [rez](https://github.com/bamiaux/rez) | 200 | 9 | 2014/01/16 | 3 years ago | Image resizing in pure Go and SIMD. |
| [darkroom](https://github.com/gojek/darkroom) | 160 | 8 | 2019/07/01 | 2 days ago | An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency. |
| [img](https://github.com/hawx/img) | 137 | 5 | 2012/07/28 | 5 years ago | Selection of image manipulation tools. |
| [mergi](https://github.com/noelyahan/mergi) | 127 | 7 | 2018/09/24 | 8 months ago | Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate). |
| [go-cairo](https://github.com/ungerik/go-cairo) | 97 | 6 | 2012/08/22 | 1 year ago | Go binding for the cairo graphics library. |
| [steganography](https://github.com/auyer/steganography) | 90 | 4 | 2018/05/21 | 1 year ago | Pure Go Library for LSB steganography. |
| [gltf](https://github.com/qmuntal/gltf) | 89 | 2 | 2019/01/15 | 2 weeks ago | Efficient and robust glTF 2.0 reader, writer and validator. |
| [cameron](https://github.com/aofei/cameron) | 66 | 4 | 2018/05/05 | 5 months ago | An avatar generator for Go. |
| [go-gd](https://github.com/bolknote/go-gd) | 52 | 4 | 2011/05/12 | 2 years ago | Go binding for GD library. |
| [goimghdr](https://github.com/corona10/goimghdr) | 36 | 1 | 2018/02/25 | 1 year ago | The imghdr module determines the type of image contained in a file for Go. |
| [gridder](https://github.com/shomali11/gridder) | 36 | 2 | 2020/04/10 | 9 months ago | A Grid based 2D Graphics library. |
| [tga](https://github.com/ftrvxmtrx/tga) | 26 | 3 | 2012/10/08 | 5 years ago | Package tga is a TARGA image format decoder/encoder. |
| [go-webcolors](https://github.com/jyotiska/go-webcolors) | 25 | 2 | 2014/04/24 | 5 years ago | Port of webcolors library from Python to Go. |
| [mpo](https://github.com/donatj/mpo) | 6 | 1 | 2015/04/14 | 8 months ago | Decoder and conversion tool for MPO 3D Photos. |
| [webp-server](https://github.com/mehdipourfar/webp-server) | 6 | 1 | 2020/11/22 | 1 month ago | Simple and minimal image server capable of storing, resizing, converting and caching images. |

## IoT
        
*Libraries for programming devices of the IoT.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [flogo](https://github.com/TIBCOSoftware/flogo) | 1672 | 148 | 2016/07/10 | 2 months ago | Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. |
| [mainflux](https://github.com/mainflux/mainflux) | 1298 | 97 | 2015/07/06 | 2 days ago | Industrial IoT Messaging and Device Management Server. |
| [gatt](https://github.com/paypal/gatt) | 961 | 54 | 2014/04/23 | 7 months ago | Gatt is a Go package for building Bluetooth Low Energy peripherals. |
| [heedy](https://github.com/heedy/heedy) | 260 | 23 | 2015/01/16 | 1 week ago | Open-Source Platform for Quantified Self & IoT. |
| [devices](https://github.com/goiot/devices) | 239 | 16 | 2016/05/30 | 4 years ago | Suite of libraries for IoT devices, experimental for x/exp/io. |
| [sensorbee](https://github.com/sensorbee/sensorbee) | 200 | 18 | 2016/02/19 | 1 year ago | Lightweight stream processing engine for IoT. |
| [huego](https://github.com/amimof/huego) | 175 | 3 | 2017/05/16 | 1 week ago | An extensive Philips Hue client library for Go. |
| [eywa](https://github.com/xcodersun/eywa) | 47 | 8 | 2016/02/20 | 3 years ago | Project Eywa is essentially a connection manager that keeps track of connected devices. |

## Job Scheduler
        
*Libraries for scheduling jobs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gron](https://github.com/roylee0704/gron) | 830 | 14 | 2016/06/04 | 1 month ago | Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly. |
| [jobrunner](https://github.com/bamzi/jobrunner) | 828 | 27 | 2015/10/21 | 3 months ago | Smart and featureful cron job scheduler with job queuing and live monitoring built in. |
| [gocron](https://github.com/go-co-op/gocron) | 525 | 15 | 2020/03/20 | 4 days ago | Easy and fluent Go job scheduling. This is an actively maintained fork of [jasonlvhit/gocron](https://github.com/jasonlvhit/gocron). |
| [jobs](https://github.com/albrow/jobs) | 480 | 18 | 2015/02/09 | 2 years ago | Persistent and flexible background jobs library. |
| [scheduler](https://github.com/carlescere/scheduler) | 356 | 15 | 2015/02/03 | 1 month ago | Cronjobs scheduling made easy. |
| [go-cron](https://github.com/rk/go-cron) | 202 | 8 | 2011/04/15 | 1 year ago | Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. |
| [go-quartz](https://github.com/reugn/go-quartz) | 108 | 7 | 2019/04/14 | 2 months ago | Simple, zero-dependency scheduling library for Go. |
| [leprechaun](https://github.com/kilgaloon/leprechaun) | 79 | 7 | 2018/04/08 | 1 month ago | Job scheduler that supports webhooks, crons and classic scheduling. |
| [clockwork](https://github.com/whiteShtef/clockwork) | 25 | 1 | 2018/04/23 | 4 months ago | Simple and intuitive job scheduling library in Go. |
| [cronticker](https://github.com/krayzpipes/cronticker) | 1 | 1 | 2020/11/28 | 1 month ago | A ticker implementation to support cron schedules. |

## JSON
        
*Libraries for working with JSON.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gjson](https://github.com/tidwall/gjson) | 7769 | 149 | 2016/08/11 | 2 weeks ago | Get a JSON value with one line of code. |
| [gojson](https://github.com/ChimeraCoder/gojson) | 2323 | 45 | 2012/12/27 | 7 months ago | Automatically generate Go (golang) struct definitions from example JSON. |
| [fastjson](https://github.com/valyala/fastjson) | 1143 | 27 | 2018/05/28 | 1 month ago | Fast JSON parser and validator for Go. No custom structs, no code generation, no reflection. |
| [kazaam](https://github.com/qntfy/kazaam) | 178 | 21 | 2016/07/19 | 9 months ago | API for arbitrary transformation of JSON documents. |
| [gojq](https://github.com/elgs/gojq) | 161 | 4 | 2015/12/30 | 2 months ago | JSON query in Golang. |
| [jsondiff](https://github.com/wI2L/jsondiff) | 103 | 2 | 2020/11/28 | 1 month ago | JSON diff library for Go based on RFC6902 (JSON Patch). |
| [jettison](https://github.com/wI2L/jettison) | 98 | 5 | 2019/08/30 | 2 months ago | Fast and flexible JSON encoder for Go. |
| [jsongo](https://github.com/ricardolonga/jsongo) | 95 | 1 | 2015/08/07 | 4 years ago | Fluent API to make it easier to create Json objects. |
| [gjo](https://github.com/skanehira/gjo) | 86 | 8 | 2019/02/23 | 1 month ago | Small utility to create JSON objects. |
| [jaydiff](https://github.com/yazgazan/jaydiff) | 80 | 1 | 2017/04/24 | 2 weeks ago | JSON diff utility written in Go. |
| [json2go](https://github.com/m-zajac/json2go) | 73 | 2 | 2017/06/10 | 2 months ago | Advanced JSON to Go struct conversion. Provides package that can parse multiple JSON documents and create struct to fit them all. |
| [jsonf](https://github.com/miolini/jsonf) | 60 | 3 | 2015/05/25 | 2 months ago | Console tool for highlighted formatting and struct query fetching JSON. |
| [ajson](https://github.com/spyzhov/ajson) | 54 | 2 | 2019/03/07 | 7 months ago | Abstract JSON for golang with JSONPath support. |
| [mp](https://github.com/sanbornm/mp) | 44 | 2 | 2014/06/15 | 4 years ago | Simple cli email parser. It currently takes stdin and outputs JSON. |
| [go-respond](https://github.com/nicklaw5/go-respond) | 40 | 1 | 2017/03/12 | 1 year ago | Go package for handling common HTTP JSON responses. |
| [go-jsonerror](https://github.com/ddymko/go-jsonerror) | 10 | 1 | 2018/10/18 | 1 year ago | Go-JsonError is ment to allow us to easily create json response errors that follow the JsonApi spec. |
| [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) | 9 | 1 | 2016/07/08 | 4 years ago | Go bindings based on the JSON API errors reference. |
| [jsonhal](https://github.com/RichardKnop/jsonhal) | 9 | 2 | 2016/01/15 | 10 months ago | Simple Go package to make custom structs marshal into HAL compatible JSON responses. |
| [dynjson](https://github.com/cocoonspace/dynjson) | 6 | 1 | 2020/05/06 | 6 months ago | Client-customizable JSON formats for dynamic APIs. |
| [ej](https://github.com/lucassscaravelli/ej) | 5 | 2 | 2020/01/04 | 10 months ago | Write and read JSON from different sources succinctly. |
| [epoch](https://github.com/vtopc/epoch) | 4 | 1 | 2019/12/15 | 3 months ago | Contains primitives for marshaling/unmarshaling Unix timestamp/epoch to/from build-in time.Time type in JSON. |
| [mapslice-json](https://github.com/mickep76/mapslice-json) | 3 | 1 | 2020/02/19 | 1 year ago | Go MapSlice for ordered marshal/ unmarshal of maps in JSON. |
| [jzon](https://github.com/zerosnake0/jzon) | 3 | 1 | 2019/11/12 | 2 months ago | JSON library with standard compatible API/behavior. |
| [jsonic](https://github.com/sinhashubham95/jsonic) | 1 | 1 | 2021/01/09 | 1 month ago | Utilities to handle and query JSON without defining structs in a type safe manner. |

## Logging
        
*Libraries for generating and working with log files.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [logrus](https://github.com/sirupsen/logrus) | 17128 | 312 | 2013/10/16 | 1 hour ago | Structured logger for Go. |
| [zap](https://github.com/uber-go/zap) | 11863 | 250 | 2016/02/18 | 2 hours ago | Fast, structured, leveled logging in Go. |
| [zerolog](https://github.com/rs/zerolog) | 4338 | 59 | 2017/05/12 | 3 weeks ago | Zero-allocation JSON logger. |
| [go-spew](https://github.com/davecgh/go-spew) | 4219 | 61 | 2013/01/09 | 2 months ago | Implements a deep pretty printer for Go data structures to aid in debugging. |
| [glog](https://github.com/golang/glog) | 2648 | 89 | 2013/07/16 | 6 days ago | Leveled execution logs for Go. |
| [lumberjack](https://github.com/natefinch/lumberjack) | 2383 | 56 | 2014/06/14 | 1 week ago | Simple rolling logger, implements io.WriteCloser. |
| [tail](https://github.com/hpcloud/tail) | 2000 | 97 | 2013/02/05 | 1 month ago | Go package striving to emulate the features of the BSD tail program. |
| [seelog](https://github.com/cihub/seelog) | 1529 | 90 | 2011/11/17 | 1 year ago | Logging functionality with flexible dispatching, filtering, and formatting. |
| [log](https://github.com/apex/log) | 1075 | 34 | 2015/12/21 | 3 weeks ago | Structured logging package for Go. |
| [log15](https://github.com/inconshreveable/log15) | 995 | 26 | 2014/05/20 | 3 months ago | Simple, powerful logging for Go. |
| [onelog](https://github.com/francoispqt/onelog) | 388 | 10 | 2018/05/06 | 1 year ago | Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenarios. Also, it is one of the logger with the lowest allocation. |
| [logxi](https://github.com/mgutz/logxi) | 346 | 10 | 2015/03/01 | 10 months ago | 12-factor app logger that is fast and makes you happy. |
| [log](https://github.com/phuslu/log) | 317 | 15 | 2019/07/07 | 6 hours ago | Structured Logging Made Easy. |
| [logutils](https://github.com/hashicorp/logutils) | 290 | 229 | 2013/10/09 | 8 months ago | Utilities for slightly better logging in Go (Golang) extending the standard logger. |
| [log](https://github.com/go-playground/log) | 276 | 10 | 2016/02/07 | 1 year ago | Simple, configurable and scalable Structured Logging for Go. |
| [go-logger](https://github.com/apsdehal/go-logger) | 261 | 7 | 2014/09/26 | 1 year ago | Simple logger of Go Programs, with level handlers. |
| [httpretty](https://github.com/henvic/httpretty) | 209 | 5 | 2020/01/24 | 1 month ago | Pretty-prints your regular HTTP requests on your terminal for debugging (similar to http.DumpRequest). |
| [rollingwriter](https://github.com/arthurkiller/rollingwriter) | 165 | 8 | 2017/02/12 | 4 months ago | RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation. |
| [sqldb-logger](https://github.com/simukti/sqldb-logger) | 153 | 2 | 2019/11/02 | 2 months ago | A logger for Go SQL database driver without modify existing *sql.DB stdlib usage. |
| [logger](https://github.com/azer/logger) | 147 | 5 | 2014/09/30 | 3 months ago | Minimalistic logging library for Go. |
| [xlog](https://github.com/rs/xlog) | 134 | 8 | 2015/10/22 | 2 days ago | Structured logger for `net/context` aware HTTP handlers with flexible dispatching. |
| [ozzo-log](https://github.com/go-ozzo/ozzo-log) | 115 | 10 | 2015/10/22 | 1 month ago | High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail). |
| [logur](https://github.com/logur/logur) | 108 | 5 | 2018/12/09 | 4 months ago | An opinionated logger interface and collection of logging best practices with adapters and integrations for well-known libraries ([logrus](https://github.com/sirupsen/logrus), [go-kit log](https://github.com/go-kit/kit/tree/master/log), [zap](https://github.com/uber-go/zap), [zerolog](https://github.com/rs/zerolog), etc). |
| [glg](https://github.com/kpango/glg) | 100 | 6 | 2017/06/21 | 10 hours ago | glg is simple and fast leveled logging library for Go. |
| [logvoyage](https://github.com/firstrow/logvoyage) | 88 | 5 | 2015/03/29 | 3 years ago | Full-featured logging saas written in golang. |
| [log](https://github.com/alexcesaro/log) | 45 | 6 | 2014/04/19 | 5 years ago | Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs. |
| [go-cronowriter](https://github.com/utahta/go-cronowriter) | 40 | 1 | 2017/02/04 | 1 month ago | Simple writer that rotate log files automatically based on current date and time, like cronolog. |
| [gologger](https://github.com/sadlil/gologger) | 39 | 6 | 2015/09/02 | 3 years ago | Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch. |
| [go-log](https://github.com/ian-kent/go-log) | 37 | 2 | 2014/05/02 | 2 years ago | Log4j implementation in Go. |
| [logex](https://github.com/chzyer/logex) | 36 | 7 | 2014/10/10 | 3 years ago | Golang log lib, supports tracking and level, wrap by standard log lib. |
| [go-log](https://github.com/siddontang/go-log) | 28 | 6 | 2014/05/18 | 2 years ago | Log lib supports level and multi handlers. |
| [distillog](https://github.com/amoghe/distillog) | 27 | 1 | 2015/10/12 | 2 years ago | distilled levelled logging (think of it as stdlib + log levels). |
| [journald](https://github.com/ssgreg/journald) | 26 | 2 | 2017/08/23 | 2 years ago | Go implementation of systemd Journal's native API for logging. |
| [logrusly](https://github.com/sebest/logrusly) | 26 | 4 | 2014/09/11 | 6 months ago | [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/). |
| [log](https://github.com/teris-io/log) | 24 | 1 | 2017/10/28 | 3 years ago | Structured log interface for Go cleanly separates logging facade from its implementation. |
| [mlog](https://github.com/jbrodriguez/mlog) | 22 | 1 | 2014/10/20 | 2 years ago | Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output. |
| [gomol](https://github.com/aphistic/gomol) | 16 | 2 | 2015/08/30 | 1 year ago | Multiple-output, structured logging for Go with extensible logging outputs. |
| [zkits-logger](https://github.com/edoger/zkits-logger) | 15 | 1 | 2020/03/31 | 1 week ago | A powerful zero-dependency JSON logger. |
| [glo](https://github.com/lajosbencz/glo) | 13 | 1 | 2019/01/19 | 2 years ago | PHP Monolog inspired logging facility with identical severity levels. |
| [go-log](https://github.com/subchen/go-log) | 11 | 2 | 2017/05/07 | 2 years ago | Simple and configurable Logging in Go, with level, formatters and writers. |
| [logrusiowriter](https://github.com/cabify/logrusiowriter) | 10 | 91 | 2019/08/09 | 7 months ago | `io.Writer` implementation using [logrus](https://github.com/sirupsen/logrus) logger. |
| [logdump](https://github.com/ewwwwwqm/logdump) | 9 | 2 | 2017/01/13 | 2 years ago | Package for multi-level logging. |
| [logmatic](https://github.com/borderstech/logmatic) | 9 | 2 | 2018/11/07 | 2 years ago | Colorized logger for Golang with dynamic log level configuration. |
| [logo](https://github.com/mbndr/logo) | 9 | 1 | 2017/02/07 | 1 month ago | Golang logger to different configurable writers. |
| [logmatic](https://github.com/mborders/logmatic) | 9 | 2 | 2018/11/07 | 1 month ago | Colorized logger for Golang with dynamic log level configuration. |
| [log](https://github.com/aerogo/log) | 8 | 1 | 2017/06/10 | 1 year ago | An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection). |
| [go-log](https://github.com/pieterclaerhout/go-log) | 7 | 1 | 2019/10/01 | 7 months ago | A logging library with strack traces, object dumping and optional timestamps. |
| [xlog](https://github.com/xfxdev/xlog) | 6 | 1 | 2016/05/05 | 2 years ago | Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format. |
| [kemba](https://github.com/clok/kemba) | 4 | 0 | 2020/07/13 | 1 week ago | A tiny debug logging tool inspired by [debug](https://github.com/visionmedia/debug), great for CLI tools and applications. |

## Machine Learning
        
*Libraries for Machine Learning.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [golearn](https://github.com/sjwhitworth/golearn) | 7670 | 439 | 2013/12/26 | 4 weeks ago | General Machine Learning library for Go. |
| [gorgonia](https://github.com/gorgonia/gorgonia) | 3853 | 186 | 2016/09/14 | 6 hours ago | graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. |
| [tfgo](https://github.com/galeone/tfgo) | 1644 | 55 | 2017/05/23 | 1 week ago | Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python. |
| [gosseract](https://github.com/otiai10/gosseract) | 1398 | 46 | 2013/10/11 | 3 weeks ago | Go package for OCR (Optical Character Recognition), by using Tesseract C++ library. |
| [goml](https://github.com/cdipaolo/goml) | 1162 | 74 | 2015/06/27 | 1 year ago | On-line Machine Learning in Go. |
| [gorse](https://github.com/zhenghaoz/gorse) | 1128 | 44 | 2018/08/14 | 3 days ago | An offline recommender system backend based on collaborative filtering written in Go. |
| [eaopt](https://github.com/MaxHalford/eaopt) | 712 | 28 | 2016/01/31 | 2 days ago | An evolutionary optimization library. |
| [bayesian](https://github.com/jbrukh/bayesian) | 699 | 34 | 2011/11/23 | 6 months ago | Naive Bayesian Classification for Golang. |
| [CloudForest](https://github.com/ryanbressler/CloudForest) | 680 | 44 | 2012/10/22 | 2 months ago | Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go. |
| [gobrain](https://github.com/goml/gobrain) | 465 | 26 | 2014/04/29 | 2 months ago | Neural Networks written in go. |
| [ocrserver](https://github.com/otiai10/ocrserver) | 353 | 13 | 2015/11/15 | 1 month ago | A simple OCR API server, seriously easy to be deployed by Docker and Heroku. |
| [onnx-go](https://github.com/owulveryck/onnx-go) | 302 | 12 | 2018/08/28 | 4 months ago | Go Interface to Open Neural Network Exchange (ONNX). |
| [go-deep](https://github.com/patrikeh/go-deep) | 299 | 16 | 2017/12/09 | 3 months ago | A feature-rich neural network library in Go. |
| [regommend](https://github.com/muesli/regommend) | 288 | 16 | 2014/02/05 | 1 year ago | Recommendation & collaborative filtering engine. |
| [go-galib](https://github.com/thoj/go-galib) | 183 | 14 | 2009/11/30 | 5 years ago | Genetic Algorithms library written in Go / golang. |
| [goRecommend](https://github.com/timkaye11/goRecommend) | 169 | 10 | 2014/07/16 | 6 years ago | Recommendation Algorithms library written in Go. |
| [goptuna](https://github.com/c-bata/goptuna) | 168 | 8 | 2019/07/24 | 4 days ago | Bayesian optimization framework for black-box functions written in Go. Everything will be optimized. |
| [shield](https://github.com/eaigner/shield) | 137 | 11 | 2013/04/10 | 11 months ago | Bayesian text classifier with flexible tokenizers and storage backends for Go. |
| [go-fann](https://github.com/vksnk/go-fann) | 104 | 8 | 2011/03/10 | 6 years ago | Go bindings for Fast Artificial Neural Networks(FANN) library. |
| [goga](https://github.com/tomcraven/goga) | 100 | 8 | 2015/10/20 | 4 years ago | Genetic algorithm library for Go. |
| [libsvm](https://github.com/datastream/libsvm) | 67 | 11 | 2012/07/31 | 4 years ago | libsvm golang version derived work based on LIBSVM 3.14. |
| [goscore](https://github.com/asafschers/goscore) | 63 | 7 | 2017/08/19 | 1 year ago | Go Scoring API for PMML. |
| [gonet](https://github.com/dathoangnd/gonet) | 63 | 5 | 2020/01/11 | 10 months ago | Neural Network for Go. |
| [neural-go](https://github.com/schuyler/neural-go) | 62 | 3 | 2011/10/17 | 5 months ago | Multilayer perceptron network implemented in Go, with training via backpropagation. |
| [go-pr](https://github.com/daviddengcn/go-pr) | 58 | 7 | 2013/06/07 | 7 years ago | Pattern recognition package in Go lang. |
| [neat](https://github.com/jinyeom/neat) | 58 | 13 | 2016/11/17 | 2 years ago | Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT). |
| [go-featureprocessing](https://github.com/nikolaydubina/go-featureprocessing) | 51 | 1 | 2020/12/18 | 49 minutes ago | Fast and convenient feature processing for low latency machine leraning in Go. |
| [fonet](https://github.com/Fontinalis/fonet) | 43 | 5 | 2017/10/03 | 9 months ago | A Deep Neural Network library written in Go. |
| [golinear](https://github.com/danieldk/golinear) | 41 | 6 | 2013/04/05 | 2 years ago | liblinear bindings for Go. |
| [Varis](https://github.com/Xamber/Varis) | 34 | 6 | 2017/10/10 | 2 years ago | Golang Neural Network. |
| [godist](https://github.com/e-dard/godist) | 27 | 3 | 2014/09/05 | 5 years ago | Various probability distributions, and associated methods. |
| [go-cluster](https://github.com/e-XpertSolutions/go-cluster) | 26 | 6 | 2017/10/04 | 2 years ago | Go implementation of the k-modes and k-prototypes clustering algorithms. |
| [evoli](https://github.com/khezen/evoli) | 15 | 5 | 2015/06/12 | 9 months ago | Genetic Algorithm and Particle Swarm Optimization library. |
| [probab](https://github.com/ThePaw/probab) | 14 | 2 | 2015/09/14 | 5 years ago | Probability distribution functions. Bayesian inference. Written in pure Go. |
| [gomind](https://github.com/surenderthakran/gomind) | 10 | 3 | 2017/10/19 | 2 years ago | A simplistic Neural Network Library in Go. |
| [randomForest](https://github.com/malaschitz/randomForest) | 7 | 3 | 2018/10/25 | 1 month ago | Easy to use Random Forest library for Go. |

## Messaging
        
*Libraries that implement messaging systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [sarama](https://github.com/Shopify/sarama) | 6804 | 428 | 2013/07/05 | 2 weeks ago | Go library for Apache Kafka. |
| [gorush](https://github.com/appleboy/gorush) | 5238 | 192 | 2016/03/22 | 3 weeks ago | Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm). |
| [machinery](https://github.com/RichardKnop/machinery) | 4911 | 147 | 2015/04/05 | 2 hours ago | Asynchronous task queue/job queue based on distributed message passing. |
| [centrifugo](https://github.com/centrifugal/centrifugo) | 4877 | 195 | 2015/03/31 | 2 days ago | Real-time messaging (Websockets or SockJS) server in Go. |
| [go-socket.io](https://github.com/googollee/go-socket.io) | 3922 | 128 | 2013/07/13 | 3 days ago | socket.io library for golang, a realtime application framework. |
| [nats.go](https://github.com/nats-io/nats.go) | 3221 | 161 | 2012/08/15 | 3 days ago | Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library. |
| [benthos](https://github.com/Jeffail/benthos) | 2853 | 82 | 2016/03/22 | 5 hours ago | A message streaming bridge between a range of protocols. |
| [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go) | 2504 | 266 | 2016/07/12 | 1 week ago | confluent-kafka-go is Confluent's Golang client for Apache Kafka and the Confluent Platform. |
| [apns2](https://github.com/sideshow/apns2) | 2435 | 78 | 2016/01/05 | 3 weeks ago | HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps. |
| [mercure](https://github.com/dunglas/mercure) | 2303 | 61 | 2018/07/14 | 7 hours ago | Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events). |
| [melody](https://github.com/olahol/melody) | 2060 | 58 | 2015/05/13 | 1 year ago | Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling. |
| [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) | 1961 | 232 | 2013/12/27 | 3 years ago | gopush-cluster is a go push server cluster. |
| [go-nsq](https://github.com/nsqio/go-nsq) | 1842 | 66 | 2013/08/29 | 4 weeks ago | the official Go package for NSQ. |
| [mangos-v1](https://github.com/nanomsg/mangos-v1) | 1532 | 83 | 2014/10/25 | 1 year ago | Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability. |
| [uniqush-push](https://github.com/uniqush/uniqush-push) | 1233 | 79 | 2011/08/29 | 10 months ago | Redis backed unified push service for server-side notifications to mobile devices. |
| [Beaver](https://github.com/Clivern/Beaver) | 1039 | 27 | 2018/10/20 | 3 days ago | A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps. |
| [asynq](https://github.com/hibiken/asynq) | 900 | 25 | 2019/11/15 | 1 week ago | A simple, reliable, and efficient distributed task queue for Go built on top of Redis. |
| [zmq4](https://github.com/pebbe/zmq4) | 890 | 42 | 2013/10/18 | 3 days ago | Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2). |
| [gollum](https://github.com/trivago/gollum) | 882 | 37 | 2015/06/20 | 1 year ago | A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations. |
| [EventBus](https://github.com/asaskevich/EventBus) | 867 | 27 | 2014/12/19 | 3 months ago | The lightweight event bus with async compatibility. |
| [dbus](https://github.com/godbus/dbus) | 549 | 17 | 2014/03/27 | 1 month ago | Native Go bindings for D-Bus. |
| [golongpoll](https://github.com/jcuga/golongpoll) | 503 | 23 | 2015/11/02 | 14 hours ago | HTTP longpoll server library that makes web pub-sub simple. |
| [emitter](https://github.com/olebedev/emitter) | 391 | 10 | 2015/11/10 | 1 year ago | Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins. |
| [mangos](https://github.com/nanomsg/mangos) | 371 | 23 | 2018/10/12 | 4 weeks ago | Pure go implementation of the Nanomsg ("Scalability Protocols") with transport interoperability. |
| [glue](https://github.com/desertbit/glue) | 362 | 16 | 2015/06/07 | 9 months ago | Robust Go and Javascript Socket Library (Alternative to Socket.io). |
| [pubsub](https://github.com/cskr/pubsub) | 340 | 8 | 2012/04/01 | 7 months ago | Simple pubsub package for go. |
| [bus](https://github.com/mustafaturan/bus) | 178 | 5 | 2019/04/27 | 1 month ago | Minimalist message bus implementation for internal communication. |
| [rabtap](https://github.com/jandelgado/rabtap) | 167 | 11 | 2017/11/11 | 2 months ago | RabbitMQ swiss army knife cli app. |
| [message-bus](https://github.com/vardius/message-bus) | 161 | 7 | 2017/10/04 | 1 month ago | messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD. |
| [guble](https://github.com/smancke/guble) | 146 | 13 | 2015/11/15 | 3 years ago | Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence. |
| [oplog](https://github.com/dailymotion/oplog) | 106 | 90 | 2014/11/06 | 5 years ago | Generic oplog/replication system for REST APIs. |
| [hub](https://github.com/leandro-lugaresi/hub) | 91 | 2 | 2018/04/13 | 3 months ago | A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges. |
| [rabbus](https://github.com/rafaeljesus/rabbus) | 85 | 9 | 2017/05/07 | 1 year ago | A tiny wrapper over amqp exchanges and queues. |
| [drone-line](https://github.com/appleboy/drone-line) | 73 | 5 | 2016/09/13 | 4 months ago | Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI. |
| [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) | 66 | 9 | 2017/01/15 | 3 years ago | A tiny wrapper around NSQ topic and channel. |
| [RapidMQ](https://github.com/sybrexsys/RapidMQ) | 60 | 5 | 2016/10/04 | 3 years ago | RapidMQ is a lightweight and reliable library for managing of the local messages queue. |
| [go-mq](https://github.com/cheshir/go-mq) | 54 | 6 | 2017/06/19 | 5 days ago | RabbitMQ client with declarative configuration. |
| [commander](https://github.com/jeroenrinzema/commander) | 52 | 1 | 2018/04/20 | 4 months ago | A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka. |
| [go-notify](https://github.com/TheCreeper/go-notify) | 51 | 2 | 2015/03/01 | 2 months ago | Native implementation of the freedesktop notification spec. |
| [redisqueue](https://github.com/robinjoseph08/redisqueue) | 50 | 2 | 2019/07/07 | 4 months ago | redisqueue provides a producer and consumer of a queue that uses Redis streams. |
| [go-res](https://github.com/jirenius/go-res) | 45 | 2 | 2018/07/15 | 2 months ago | Package for building REST/real-time services where clients are synchronized seamlessly, using NATS and Resgate. |
| [commander](https://github.com/CloudProud/commander) | 38 | 1 | 2018/04/20 | 1 year ago | A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka. |
| [event](https://github.com/agoalofalife/event) | 36 | 4 | 2017/07/02 | 3 years ago | Implementation of the pattern observer. |
| [gosd](https://github.com/alexsniffin/gosd) | 17 | 1 | 2020/05/17 | 3 months ago | A library for scheduling when to dispatch a message to a channel. |
| [go-vitotrol](https://github.com/maxatome/go-vitotrol) | 16 | 7 | 2016/11/03 | 1 year ago | Client library to Viessmann Vitotrol web service. |
| [ami](https://github.com/kak-tus/ami) | 14 | 1 | 2018/10/27 | 10 months ago | Go client to reliable queues based on Redis Cluster Streams. |
| [hare](https://github.com/leozz37/hare) | 14 | 1 | 2020/12/01 | 2 weeks ago | A user friendly library for sending messages and listening to TCP sockets. |
| [jazz](https://github.com/socifi/jazz) | 13 | 3 | 2018/10/22 | 1 year ago | A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages. |
| [rmqconn](https://github.com/sbabiv/rmqconn) | 13 | 1 | 2019/01/14 | 1 year ago | RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed. |
| [gaurun-client](https://github.com/osamingo/gaurun-client) | 9 | 1 | 2017/06/29 | 1 year ago | Gaurun Client written in Go. |

## Microsoft Office
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [unioffice](https://github.com/unidoc/unioffice) | 2727 | 67 | 2017/08/29 | 1 month ago | Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents. |

### Microsoft Excel
        
*Libraries for working with Microsoft Excel.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [excelize](https://github.com/360EntSecGroup-Skylar/excelize) | 8023 | 191 | 2016/08/29 | 1 hour ago | Golang library for reading and writing Microsoft Excel™ (XLSX) files. |
| [xlsx](https://github.com/tealeg/xlsx) | 4769 | 178 | 2011/06/28 | 5 days ago | Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs. |
| [xlsx](https://github.com/plandem/xlsx) | 130 | 12 | 2017/08/26 | 3 months ago | Fast and safe way to read/update your existing Microsoft Excel files in Go programs. |
| [go-excel](https://github.com/szyhf/go-excel) | 113 | 4 | 2017/09/03 | 1 month ago | A simple and light reader to read a relate-db-like excel as a table. |
| [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) | 15 | 2 | 2017/03/13 | 2 years ago | Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files. |

## Miscellaneous
        

### Dependency Injection
        
*Libraries for working with dependency injection.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [dig](https://github.com/uber-go/dig) | 1776 | 50 | 2017/03/21 | 4 weeks ago | A reflection based dependency injection toolkit for Go. |
| [fx](https://github.com/uber-go/fx) | 1772 | 57 | 2016/10/27 | 5 days ago | A dependency injection based application framework for Go (built on top of dig). |
| [container](https://github.com/golobby/container) | 141 | 4 | 2019/09/23 | 3 weeks ago | A powerful IoC Container with fluent and easy-to-use interface. |
| [dingo](https://github.com/i-love-flamingo/dingo) | 96 | 24 | 2018/10/29 | 5 months ago | A dependency injection toolkit for Go, based on Guice. |
| [di](https://github.com/goioc/di) | 66 | 3 | 2020/06/11 | 2 weeks ago | Spring-inspired Dependency Injection Container. |
| [di](https://github.com/goava/di) | 64 | 7 | 2020/02/03 | 5 hours ago | A dependency injection container for go programming language. |
| [inject](https://github.com/defval/inject) | 53 | 1 | 2019/02/03 | 1 year ago | A reflection based dependency injection container with simple interface. |
| [alice](https://github.com/magic003/alice) | 44 | 2 | 2017/04/08 | 3 years ago | Additive dependency injection container for Golang. |
| [wire](https://github.com/Fs02/wire) | 33 | 1 | 2018/07/05 | 1 year ago | Strict Runtime Dependency Injection for Golang. |
| [linker](https://github.com/logrange/linker) | 29 | 3 | 2018/12/04 | 7 months ago | A reflection based dependency injection and inversion of control library with components lifecycle support. |
| [gocontainer](https://github.com/vardius/gocontainer) | 14 | 0 | 2019/06/06 | 10 months ago | Simple Dependency Injection Container. |

### Project Layout
        
*Unofficial set of patterns for structuring projects.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [project-layout](https://github.com/golang-standards/project-layout) | 21115 | 519 | 2017/09/09 | 1 week ago | Set of common historical and emerging project layout patterns in the Go ecosystem. |
| [go-restful-api](https://github.com/qiangxue/go-restful-api) | 1038 | 52 | 2016/08/15 | 1 year ago | An idiomatic Go RESTful API starter kit following SOLID principles and Clean Architecture with a common project layout. |
| [modern-go-application](https://github.com/sagikazarmark/modern-go-application) | 873 | 18 | 2018/09/14 | 1 week ago | Go application boilerplate and example applying modern practices. |
| [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) | 428 | 9 | 2016/12/18 | 1 week ago | A Go application boilerplate template for quick starting projects following production best practices. |
| [seed](https://github.com/golang-templates/seed) | 98 | 3 | 2020/04/30 | 6 days ago | Go application GitHub repository template. |
| [scaffold](https://github.com/catchplay/scaffold) | 95 | 4 | 2018/12/11 | 2 years ago | Scaffold generates a starter Go project layout. Lets you focus on business logic implemented. |
| [go-sample](https://github.com/zitryss/go-sample) | 81 | 1 | 2019/01/24 | 2 years ago | A sample layout for Go application projects with the real code. |
| [go-todo-backend](https://github.com/Fs02/go-todo-backend) | 49 | 3 | 2020/06/25 | 3 months ago | Go Todo Backend example using modular project layout for product microservice. |
| [gobase](https://github.com/wajox/gobase) | 2 | 1 | 2020/12/15 | 1 month ago | A simple skeleton for golang application with basic setup for real golang application. |

### Strings
        
*Libraries for working with strings.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [xstrings](https://github.com/huandu/xstrings) | 880 | 24 | 2015/01/06 | 1 month ago | Collection of useful string functions ported from other languages. |
| [strutil](https://github.com/ozgio/strutil) | 111 | 2 | 2018/08/16 | 1 year ago | String utilities. |
| [stringy](https://github.com/gobeam/stringy) | 47 | 2 | 2020/04/03 | 7 months ago | String manipulation library to convert string to camel case, snake case, kebab case / slugify etc. |
| [Stringy](https://github.com/gobeam/Stringy) | 30 | 1 | 2020/04/03 | 9 months ago | String manipulation library to convert string to camel case, snake case, kebab case / slugify etc. |

### Uncategorized
        
*These libraries were placed here because none of the other categories seemed to fit.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gopsutil](https://github.com/shirou/gopsutil) | 5880 | 201 | 2014/04/18 | 5 days ago | Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc). |
| [archiver](https://github.com/mholt/archiver) | 3083 | 52 | 2016/04/08 | 2 hours ago | Library and command for making and extracting .zip and .tar.gz archives. |
| [gofakeit](https://github.com/brianvoe/gofakeit) | 1660 | 15 | 2015/04/24 | 11 hours ago | Random data generator written in go. |
| [gosms](https://github.com/haxpax/gosms) | 1325 | 56 | 2015/01/23 | 1 week ago | Your own local SMS gateway in Go that can be used to send SMS. |
| [go-resiliency](https://github.com/eapache/go-resiliency) | 1145 | 28 | 2014/11/29 | 3 months ago | Resiliency patterns for golang. |
| [gatus](https://github.com/TwinProduction/gatus) | 1139 | 16 | 2019/09/04 | 1 day ago | Automated service health dashboard. |
| [base64Captcha](https://github.com/mojocn/base64Captcha) | 1033 | 47 | 2017/12/12 | 5 months ago | Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha. |
| [go-commons-pool](https://github.com/jolestar/go-commons-pool) | 911 | 49 | 2015/12/28 | 1 month ago | Generic object pool for Golang. |
| [llvm](https://github.com/llir/llvm) | 650 | 28 | 2014/09/19 | 2 hours ago | Library for interacting with LLVM IR in pure Go. |
| [shortid](https://github.com/teris-io/shortid) | 646 | 9 | 2016/01/04 | 2 months ago | Distributed generation of super short, unique, non-sequential, URL friendly IDs. |
| [health](https://github.com/dimiro1/health) | 410 | 6 | 2016/03/08 | 1 year ago | Easy to use, extensible health check library. |
| [go-conv](https://github.com/cstockton/go-conv) | 365 | 9 | 2016/10/11 | 3 years ago | Package conv provides fast and intuitive conversions across Go types. |
| [ghorg](https://github.com/gabrie30/ghorg) | 349 | 7 | 2018/03/29 | 4 hours ago | Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, and Bitbucket. |
| [banner](https://github.com/dimiro1/banner) | 328 | 6 | 2016/03/25 | 1 month ago | Add beautiful banners into your Go applications. |
| [gountries](https://github.com/pariz/gountries) | 299 | 10 | 2016/01/13 | 5 months ago | Package that exposes country and subdivision data. |
| [stateless](https://github.com/qmuntal/stateless) | 242 | 4 | 2019/09/11 | 3 weeks ago | A fluent library for creating state machines. |
| [ffmt](https://github.com/go-ffmt/ffmt) | 216 | 5 | 2015/02/14 | 6 months ago | Beautify data display for Humans. |
| [antch](https://github.com/antchfx/antch) | 196 | 14 | 2017/09/28 | 8 months ago | A fast, powerful and extensible web crawling & scraping framework. |
| [lk](https://github.com/hyperboloide/lk) | 184 | 6 | 2016/07/14 | 9 months ago | A simple licensing library for golang. |
| [battery](https://github.com/distatus/battery) | 167 | 4 | 2016/03/12 | 1 month ago | Cross-platform, normalized battery information library. |
| [shoutrrr](https://github.com/containrrr/shoutrrr) | 166 | 7 | 2019/04/11 | 6 days ago | Notification library providing easy access to various messaging services like slack, mattermost, gotify and smtp among others. |
| [healthcheck](https://github.com/etherlabsio/healthcheck) | 158 | 6 | 2017/08/18 | 1 year ago | An opinionated and concurrent health-check HTTP handler for RESTful services. |
| [stats](https://github.com/go-playground/stats) | 146 | 2 | 2015/09/14 | 4 years ago | Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... |
| [bitio](https://github.com/icza/bitio) | 142 | 7 | 2016/05/31 | 1 year ago | Highly optimized bit-level Reader and Writer for Go. |
| [go-unarr](https://github.com/gen2brain/go-unarr) | 121 | 6 | 2015/11/01 | 4 months ago | Decompression library for RAR, TAR, ZIP and 7z archives. |
| [turtle](https://github.com/hackebrot/turtle) | 116 | 2 | 2017/09/08 | 8 months ago | Emojis for Go. |
| [gommit](https://github.com/antham/gommit) | 92 | 3 | 2016/08/30 | 1 month ago | Analyze git commit messages to ensure they follow defined patterns. |
| [gotoprom](https://github.com/cabify/gotoprom) | 88 | 95 | 2018/10/10 | 1 year ago | Type-safe metrics builder wrapper library for the official Prometheus client. |
| [indigo](https://github.com/osamingo/indigo) | 79 | 1 | 2016/08/31 | 1 week ago | Distributed unique ID generator of using Sonyflake and encoded by Base58. |
| [captcha](https://github.com/steambap/captcha) | 71 | 5 | 2017/09/12 | 1 month ago | Package captcha provides an easy to use, unopinionated API for captcha generation. |
| [morse](https://github.com/alwindoss/morse) | 63 | 3 | 2018/08/15 | 2 years ago | Library to convert to and from morse code. |
| [xkg](https://github.com/go-xkg/xkg) | 50 | 1 | 2015/01/05 | 6 years ago | X Keyboard Grabber. |
| [pdfgen](https://github.com/hyperboloide/pdfgen) | 49 | 3 | 2015/11/30 | 3 years ago | HTTP service to generate PDF from Json requests. |
| [persian](https://github.com/mavihq/persian) | 49 | 3 | 2017/10/16 | 1 month ago | Some utilities for Persian language in go. |
| [browscap_go](https://github.com/digitalcrab/browscap_go) | 36 | 6 | 2014/09/18 | 1 year ago | GoLang Library for [Browser Capabilities Project](http://browscap.org/). |
| [datacounter](https://github.com/miolini/datacounter) | 36 | 1 | 2015/10/14 | 1 year ago | Go counters for readers/writer/http.ResponseWriter. |
| [faker](https://github.com/pioz/faker) | 34 | 2 | 2020/07/22 | 2 months ago | Random fake data and struct generator for Go. |
| [autoflags](https://github.com/artyom/autoflags) | 33 | 5 | 2014/05/15 | 2 years ago | Go package to automatically define command line flags from struct fields. |
| [sandid](https://github.com/aofei/sandid) | 26 | 1 | 2018/06/12 | 5 months ago | Every grain of sand on earth has its own ID. |
| [gosh](https://github.com/osamingo/gosh) | 25 | 0 | 2018/05/25 | 1 month ago | Provide Go Statistics Handler, Struct, Measure Method. |
| [url-shortener](https://github.com/pantrif/url-shortener) | 25 | 1 | 2018/06/04 | 2 years ago | A modern, powerful, and robust URL shortener microservice with mysql support. |
| [xdg](https://github.com/rkoesters/xdg) | 24 | 2 | 2013/12/15 | 2 years ago | FreeDesktop.org (xdg) Specs implemented in Go. |
| [metrics](https://github.com/pascaldekloe/metrics) | 19 | 1 | 2019/01/29 | 11 months ago | Library for metrics instrumentation and Prometheus exposition. |
| [shellwords](https://github.com/Wing924/shellwords) | 13 | 1 | 2017/09/28 | 3 years ago | A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell. |
| [anagent](https://github.com/mudler/anagent) | 12 | 2 | 2017/12/29 | 2 years ago | Minimalistic, pluggable Golang evloop/timer handler with dependency-injection. |
| [avgRating](https://github.com/kirillDanshin/avgRating) | 10 | 2 | 2017/08/05 | 3 years ago | Calculate average score and rating based on Wilson Score Equation. |
| [hostutils](https://github.com/Wing924/hostutils) | 9 | 1 | 2017/09/26 | 2 years ago | A golang library for packing and unpacking FQDNs list. |
| [numa](https://github.com/lrita/numa) | 4 | 2 | 2018/12/10 | 2 months ago | NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code. |

## Natural Language Processing
        
*Libraries for working with human languages.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [prose](https://github.com/jdkato/prose) | 2682 | 62 | 2017/02/17 | 3 weeks ago | Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. English only. |
| [gse](https://github.com/go-ego/gse) | 1526 | 55 | 2017/06/23 | 1 week ago | Go efficient text segmentation; support english, chinese, japanese and other. |
| [gojieba](https://github.com/yanyiwu/gojieba) | 1395 | 67 | 2015/09/12 | 1 month ago | This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm. |
| [when](https://github.com/olebedev/when) | 1108 | 25 | 2016/12/27 | 5 months ago | Natural EN and RU language date/time parser with pluggable rules. |
| [go-pinyin](https://github.com/mozillazg/go-pinyin) | 897 | 36 | 2014/11/09 | 2 weeks ago | CN Hanzi to Hanyu Pinyin converter. |
| [spago](https://github.com/nlpodyssey/spago) | 827 | 24 | 2020/01/05 | 2 hours ago | Self-contained Machine Learning and Natural Language Processing library in Go. |
| [kagome](https://github.com/ikawaha/kagome) | 552 | 23 | 2014/06/26 | 2 weeks ago | JP morphological analyzer written in pure Go. |
| [whatlanggo](https://github.com/abadojack/whatlanggo) | 477 | 15 | 2017/02/20 | 1 month ago | Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc). |
| [nlp](https://github.com/shixzie/nlp) | 368 | 23 | 2017/01/25 | 3 years ago | Extract values from strings and fill your structs with nlp. |
| [nlp](https://github.com/james-bowman/nlp) | 304 | 24 | 2017/03/15 | 6 months ago | Go Natural Language Processing library supporting LSA (Latent Semantic Analysis). |
| [sentences](https://github.com/neurosnap/sentences) | 293 | 14 | 2015/08/07 | 1 year ago | Sentence tokenizer:  converts text into a list of sentences. |
| [getlang](https://github.com/rylans/getlang) | 110 | 4 | 2018/03/01 | 1 month ago | Fast natural language detection package. |
| [go-nlp](https://github.com/nuance/go-nlp) | 89 | 7 | 2011/05/02 | 9 years ago | Utilities for working with discrete probability distributions and other tools useful for doing NLP work. |
| [go-unidecode](https://github.com/mozillazg/go-unidecode) | 84 | 2 | 2016/07/08 | 1 year ago | ASCII transliterations of Unicode text. |
| [RAKE.Go](https://github.com/afjoseph/RAKE.Go) | 75 | 7 | 2016/12/17 | 11 months ago | Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE). |
| [gounidecode](https://github.com/fiam/gounidecode) | 72 | 3 | 2012/05/01 | 5 years ago | Unicode transliterator (also known as unidecode) for Go. |
| [textcat](https://github.com/pebbe/textcat) | 64 | 6 | 2012/09/21 | 2 years ago | Go package for n-gram based text categorization, with support for utf-8 and raw text. |
| [go-stem](https://github.com/agonopol/go-stem) | 60 | 4 | 2011/09/23 | 2 years ago | Implementation of the porter stemming algorithm. |
| [MMSEGO](https://github.com/awsong/MMSEGO) | 59 | 5 | 2012/04/18 | 8 years ago | This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm. |
| [stemmer](https://github.com/dchest/stemmer) | 49 | 4 | 2011/03/21 | 4 years ago | Stemmer packages for Go programming language. Includes English and German stemmers. |
| [go2vec](https://github.com/danieldk/go2vec) | 41 | 7 | 2015/01/27 | 2 years ago | Reader and utility functions for word2vec embeddings. |
| [porter2](https://github.com/zentures/porter2) | 39 | 2 | 2015/01/21 | 4 months ago | Really fast Porter 2 stemmer. |
| [petrovich](https://github.com/striker2000/petrovich) | 33 | 1 | 2016/12/26 | 9 months ago | Petrovich is the library which inflects Russian names to given grammatical case. |
| [snowball](https://github.com/goodsign/snowball) | 28 | 1 | 2012/12/11 | 3 years ago | Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/). |
| [paicehusk](https://github.com/rookii/paicehusk) | 26 | 3 | 2012/09/29 | 7 years ago | Golang implementation of the Paice/Husk Stemming Algorithm. |
| [address](https://github.com/bojanz/address) | 26 | 2 | 2020/10/07 | 1 month ago | Handles address representation, validation and formatting. |
| [mystem](https://github.com/dveselov/mystem) | 25 | 3 | 2016/08/30 | 4 years ago | CGo bindings to Yandex.Mystem - russian morphology analyzer. |
| [icu](https://github.com/goodsign/icu) | 19 | 0 | 2012/12/11 | 3 years ago | Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1. |
| [golibstemmer](https://github.com/rjohnsondev/golibstemmer) | 18 | 2 | 2012/08/06 | 6 years ago | Go bindings for the snowball libstemmer library including porter 2. |
| [iuliia-go](https://github.com/mehanizm/iuliia-go) | 18 | 2 | 2020/04/27 | 8 months ago | Transliterate Cyrillic → Latin in every possible way. |
| [go-localize](https://github.com/m1/go-localize) | 15 | 1 | 2019/12/23 | 3 months ago | Simple and easy to use i18n (Internationalization and localization) engine - used for translating locale strings. |
| [shamoji](https://github.com/osamingo/shamoji) | 12 | 2 | 2017/07/23 | 1 month ago | The shamoji is word filtering package written in Go. |
| [libtextcat](https://github.com/goodsign/libtextcat) | 10 | 1 | 2012/12/10 | 8 years ago | Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2. |
| [govader](https://github.com/jonreiter/govader) | 10 | 1 | 2020/01/19 | 4 months ago | Go implementation of [VADER Sentiment Analysis](https://github.com/cjhutto/vaderSentiment). |
| [gotokenizer](https://github.com/xujiajun/gotokenizer) | 9 | 1 | 2018/10/11 | 1 year ago | A tokenizer based on the dictionary and Bigram language models for Go. (Now only support chinese segmentation) |
| [transliterator](https://github.com/alexsergivan/transliterator) | 9 | 1 | 2020/04/17 | 9 months ago | Provides one-way string transliteration with supporting of language-specific transliteration rules. |
| [porter](https://github.com/a2800276/porter) | 8 | 1 | 2013/09/17 | 7 years ago | This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm. |
| [detectlanguage-go](https://github.com/detectlanguage/detectlanguage-go) | 7 | 0 | 2019/12/14 | 4 months ago | Language Detection API Go Client. Supports batch requests, short phrase or single word language detection. |
| [gosentiwordnet](https://github.com/dinopuguh/gosentiwordnet) | 5 | 1 | 2020/04/21 | 2 months ago | Sentiment analyzer using sentiwordnet lexicon in Go. |

## Networking
        
*Libraries for working with various layers of the network.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [fasthttp](https://github.com/valyala/fasthttp) | 14423 | 399 | 2015/10/18 | 18 hours ago | Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http. |
| [kcptun](https://github.com/xtaci/kcptun) | 12296 | 592 | 2016/02/26 | 1 month ago | Extremely simple & fast udp tunnel based on KCP protocol. |
| [webrtc](https://github.com/pion/webrtc) | 6528 | 228 | 2018/05/18 | 16 hours ago | A pure Go implementation of the WebRTC API. |
| [dns](https://github.com/miekg/dns) | 5192 | 174 | 2010/08/03 | 13 hours ago | Go library for working with DNS. |
| [quic-go](https://github.com/lucas-clemente/quic-go) | 4891 | 191 | 2016/04/06 | 1 day ago | An implementation of the QUIC protocol in pure Go. |
| [gopacket](https://github.com/google/gopacket) | 3915 | 143 | 2015/03/16 | 4 days ago | Go library for packet processing with libpcap bindings. |
| [gnet](https://github.com/panjf2000/gnet) | 3667 | 129 | 2019/02/24 | 5 days ago | `gnet` is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go. |
| [httplab](https://github.com/gchaincl/httplab) | 3650 | 62 | 2017/02/08 | 1 year ago | HTTPLabs let you inspect HTTP requests and forge responses. |
| [kcp-go](https://github.com/xtaci/kcp-go) | 2864 | 145 | 2015/06/16 | 6 days ago | KCP - Fast and Reliable ARQ Protocol. |
| [gobgp](https://github.com/osrg/gobgp) | 2135 | 123 | 2014/09/14 | 1 day ago | BGP implemented in the Go Programming Language. |
| [ssh](https://github.com/gliderlabs/ssh) | 1860 | 52 | 2016/10/03 | 1 week ago | Higher-level API for building SSH servers (wraps crypto/ssh). |
| [fortio](https://github.com/fortio/fortio) | 1760 | 39 | 2017/10/10 | 2 weeks ago | Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC. |
| [water](https://github.com/songgao/water) | 1185 | 43 | 2013/03/25 | 4 months ago | Simple TUN/TAP library. |
| [go-getter](https://github.com/hashicorp/go-getter) | 1096 | 220 | 2015/10/12 | 2 weeks ago | Go library for downloading files or directories from various sources using a URL. |
| [gev](https://github.com/Allenxuxu/gev) | 1032 | 35 | 2019/09/01 | 3 weeks ago | gev is a lightweight, fast non-blocking TCP network library based on Reactor mode. |
| [nff-go](https://github.com/intel-go/nff-go) | 1029 | 79 | 2017/03/29 | 5 months ago | Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF). |
| [sftp](https://github.com/pkg/sftp) | 956 | 55 | 2013/11/05 | 3 days ago | Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt. |
| [grab](https://github.com/cavaliercoder/grab) | 799 | 18 | 2016/01/05 | 1 week ago | Go package for managing file downloads. |
| [ftp](https://github.com/jlaffaye/ftp) | 774 | 25 | 2011/05/06 | 17 hours ago | Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959). |
| [mdns](https://github.com/hashicorp/mdns) | 742 | 224 | 2014/01/29 | 2 weeks ago | Simple mDNS (Multicast DNS) client/server library in Golang. |
| [vssh](https://github.com/yahoo/vssh) | 701 | 21 | 2020/06/09 | 2 months ago | Go library for building network and server automation over SSH protocol. |
| [gosnmp](https://github.com/gosnmp/gosnmp) | 676 | 49 | 2012/08/27 | 1 day ago | Native Go library for performing SNMP actions. |
| [lhttp](https://github.com/fanux/lhttp) | 603 | 58 | 2015/12/29 | 2 years ago | Powerful websocket framework, build your IM server more easily. |
| [cidranger](https://github.com/yl2chen/cidranger) | 572 | 14 | 2017/08/21 | 4 months ago | Fast IP to CIDR lookup for Go. |
| [gosnmp](https://github.com/soniah/gosnmp) | 563 | 46 | 2012/08/27 | 4 months ago | Native Go library for performing SNMP actions. |
| [gotcp](https://github.com/gansidui/gotcp) | 477 | 41 | 2014/04/13 | 3 years ago | Go package for quickly writing tcp applications. |
| [peerdiscovery](https://github.com/schollz/peerdiscovery) | 464 | 19 | 2018/04/22 | 1 week ago | Pure Go library for cross-platform local peer discovery using UDP multicast. |
| [stun](https://github.com/gortc/stun) | 443 | 18 | 2016/04/24 | 6 days ago | Go implementation of RFC 5389 STUN protocol. |
| [go-stun](https://github.com/ccding/go-stun) | 411 | 13 | 2013/08/17 | 2 weeks ago | Go implementation of the STUN client (RFC 3489 and RFC 5389). |
| [gopcap](https://github.com/akrennmair/gopcap) | 407 | 23 | 2009/11/19 | 10 months ago | Go wrapper for libpcap. |
| [raw](https://github.com/mdlayher/raw) | 382 | 12 | 2015/07/06 | 9 months ago | Package raw enables reading and writing data at the device driver level for a network interface. |
| [tcp_server](https://github.com/firstrow/tcp_server) | 348 | 18 | 2014/10/13 | 1 year ago | Go library for building tcp servers faster. |
| [winrm](https://github.com/masterzen/winrm) | 290 | 20 | 2013/12/30 | 1 week ago | Go WinRM client to remotely execute commands on Windows machines. |
| [gaio](https://github.com/xtaci/gaio) | 265 | 12 | 2019/12/20 | 2 weeks ago | High performance async-io networking for Golang in proactor mode. |
| [gmqtt](https://github.com/DrmagicE/gmqtt) | 261 | 22 | 2018/09/16 | 2 days ago | Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1. |
| [buffstreams](https://github.com/StabbyCutyou/buffstreams) | 242 | 13 | 2015/06/29 | 6 months ago | Streaming protocolbuffer data over TCP made easy. |
| [arp](https://github.com/mdlayher/arp) | 239 | 10 | 2015/07/06 | 1 year ago | Package arp implements the ARP protocol, as described in RFC 826. |
| [ethernet](https://github.com/mdlayher/ethernet) | 214 | 13 | 2015/07/03 | 1 year ago | Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. |
| [gnxi](https://github.com/google/gnxi) | 171 | 31 | 2017/09/26 | 1 month ago | A collection of tools for Network Management that use the gNMI and gNOI protocols. |
| [jazigo](https://github.com/udhos/jazigo) | 163 | 13 | 2016/06/07 | 1 year ago | Jazigo is a tool written in Go for retrieving configuration for multiple network devices. |
| [utp](https://github.com/anacrolix/utp) | 155 | 16 | 2015/03/20 | 2 weeks ago | Go uTP micro transport protocol implementation. |
| [canopus](https://github.com/zubairhamed/canopus) | 142 | 14 | 2015/02/24 | 2 years ago | CoAP Client/Server implementation (RFC 7252). |
| [sslb](https://github.com/eduardonunesp/sslb) | 130 | 8 | 2015/10/18 | 1 year ago | It's a Super Simples Load Balancer, just a little project to achieve some kind of performance. |
| [xtcp](https://github.com/xfxdev/xtcp) | 114 | 15 | 2016/03/31 | 11 months ago | TCP Server Framework with simultaneous full duplex communication, graceful shutdown, and custom protocol. |
| [dhcp6](https://github.com/mdlayher/dhcp6) | 68 | 4 | 2015/05/22 | 1 year ago | Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. |
| [ether](https://github.com/songgao/ether) | 68 | 4 | 2014/05/21 | 4 years ago | Cross-platform Go package for sending and receiving ethernet frames. |
| [packet](https://github.com/aerogo/packet) | 56 | 4 | 2017/10/29 | 1 year ago | Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed. |
| [linkio](https://github.com/ian-kent/linkio) | 50 | 2 | 2014/12/24 | 3 years ago | Network link speed simulation for Reader/Writer interfaces. |
| [portproxy](https://github.com/aybabtme/portproxy) | 43 | 0 | 2014/12/13 | 6 years ago | Simple TCP proxy which adds CORS support to API's which don't support it. |
| [graval](https://github.com/koofr/graval) | 26 | 8 | 2014/04/22 | 4 months ago | Experimental FTP server framework. |
| [go-powerdns](https://github.com/joeig/go-powerdns) | 23 | 2 | 2018/06/21 | 1 month ago | PowerDNS API bindings for Golang. |
| [publicip](https://github.com/polera/publicip) | 21 | 2 | 2016/12/28 | 4 years ago | Package publicip returns your public facing IPv4 address (internet egress). |
| [golibwireshark](https://github.com/sunwxg/golibwireshark) | 19 | 2 | 2015/11/16 | 3 years ago | Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data. |
| [panoptes-stream](https://github.com/yahoo/panoptes-stream) | 18 | 9 | 2020/10/09 | 3 weeks ago | A cloud native distributed streaming network telemetry (gNMI, Juniper JTI and Cisco MDT). |
| [gohooks](https://github.com/averageflow/gohooks) | 12 | 1 | 2020/10/30 | 1 month ago | GoHooks make it easy to send and consume secured web-hooks from a Go application. Inspired by Spatie's Laravel Webhook Client and Server. |
| [llb](https://github.com/kirillDanshin/llb) | 11 | 1 | 2016/02/21 | 4 years ago | It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response. |
| [goshark](https://github.com/sunwxg/goshark) | 10 | 1 | 2015/11/01 | 3 years ago | Package goshark use tshark to decode IP packet and create data struct to analyse packet. |
| [tspool](https://github.com/two/tspool) | 8 | 0 | 2018/10/27 | 2 years ago | A TCP Library use worker pool to improve performance and protect your server. |
| [httpproxy](https://github.com/wzshiming/httpproxy) | 7 | 0 | 2018/07/18 | 1 month ago | HTTP proxy handler and dialer. |
| [gosocsvr](https://github.com/Rakeki/gosocsvr) | 5 | 2 | 2019/11/12 | 1 year ago | Socket server made simple. |

### HTTP Clients
        
*Libraries for making HTTP requests.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [resty](https://github.com/go-resty/resty) | 3895 | 81 | 2015/08/28 | 2 days ago | Simple HTTP and REST client for Go inspired by Ruby rest-client. |
| [heimdall](https://github.com/gojek/heimdall) | 1769 | 49 | 2018/01/19 | 1 month ago | An enchanced http client with retry and hystrix capabilities. |
| [grequests](https://github.com/levigross/grequests) | 1734 | 35 | 2015/06/11 | 2 months ago | A Go "clone" of the great and famous Requests library. |
| [sling](https://github.com/dghubble/sling) | 1245 | 32 | 2015/04/02 | 3 months ago | Sling is a Go HTTP client library for creating and sending API requests. |
| [gentleman](https://github.com/h2non/gentleman) | 874 | 20 | 2016/02/21 | 1 day ago | Full-featured plugin-driven HTTP client library. |
| [pester](https://github.com/sethgrid/pester) | 536 | 6 | 2015/05/20 | 8 months ago | Go HTTP client calls with retries, backoff, and concurrency. |
| [request](https://github.com/monaco-io/request) | 95 | 8 | 2020/03/25 | 3 hours ago | HTTP client for golang. If you have experience about axios or requests, you will love it. No 3rd dependency. |
| [sreq](https://github.com/winterssy/sreq) | 50 | 0 | 2019/12/04 | 1 year ago | A simple, user-friendly and concurrent safe HTTP request library for Go. |
| [rq](https://github.com/ddo/rq) | 37 | 2 | 2017/12/26 | 1 year ago | A nicer interface for golang stdlib HTTP client. |
| [go-http-client](https://github.com/bozd4g/go-http-client) | 16 | 1 | 2019/12/14 | 2 months ago | Make http calls simply and easily. |
| [httpretry](https://github.com/ybbus/httpretry) | 11 | 2 | 2020/02/05 | 1 year ago | Enriches the default go HTTP client with retry functionality. |

## OpenGL
        
*Libraries for using OpenGL in Go.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [glfw](https://github.com/go-gl/glfw) | 1051 | 41 | 2013/05/19 | 3 months ago | Go bindings for GLFW 3. |
| [gl](https://github.com/go-gl/gl) | 779 | 38 | 2015/02/22 | 1 year ago | Go bindings for OpenGL (generated via glow). |
| [mathgl](https://github.com/go-gl/mathgl) | 356 | 25 | 2013/02/13 | 4 months ago | Pure Go math package specialized for 3D math, with inspiration from GLM. |
| [gl](https://github.com/goxjs/gl) | 147 | 15 | 2015/05/18 | 1 month ago | Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android). |
| [glfw](https://github.com/goxjs/glfw) | 68 | 6 | 2014/12/27 | 10 months ago | Go cross-platform glfw library for creating an OpenGL context and receiving events. |
| [go-glmatrix](https://github.com/technohippy/go-glmatrix) | 1 | 1 | 2020/07/02 | 1 week ago | Go port of [glMatrix](http://glmatrix.net/) library. |

## ORM
        
*Libraries that implement Object-Relational Mapping or datamapping techniques.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gorm](https://github.com/go-gorm/gorm) | 22754 | 505 | 2013/10/25 | 16 hours ago | The fantastic ORM library for Golang, aims to be developer friendly. |
| [ent](https://github.com/ent/ent) | 6645 | 122 | 2019/06/12 | 1 day ago | An entity framework for Go. Simple, yet powerful ORM for modeling and querying data. |
| [ent](https://github.com/facebook/ent) | 6615 | 123 | 2019/06/12 | 5 days ago | An entity framework for Go. Simple, yet powerful ORM for modeling and querying data. |
| [xorm](https://github.com/go-xorm/xorm) | 6186 | 262 | 2013/05/09 | 10 months ago | Simple and powerful ORM for Go. |
| [pg](https://github.com/go-pg/pg) | 4392 | 83 | 2013/04/24 | 1 day ago | PostgreSQL ORM with focus on PostgreSQL specific features and performance. |
| [sqlboiler](https://github.com/volatiletech/sqlboiler) | 3652 | 79 | 2016/02/21 | 4 hours ago | ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema. |
| [gorp](https://github.com/go-gorp/gorp) | 3470 | 114 | 2012/01/04 | 2 months ago | Go Relational Persistence, ORM-ish library for Go. |
| [ent](https://github.com/facebookincubator/ent) | 3067 | 65 | 2019/06/12 | 6 months ago | An entity framework for Go. Simple, yet powerful ORM for modeling and querying data. |
| [db](https://github.com/upper/db) | 2467 | 61 | 2013/10/23 | 1 month ago | Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers. |
| [reform](https://github.com/go-reform/reform) | 1069 | 25 | 2016/02/25 | 6 days ago | Better ORM for Go, based on non-empty interfaces and code generation. |
| [pop](https://github.com/gobuffalo/pop) | 1025 | 24 | 2018/02/07 | 3 days ago | Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite. |
| [gormt](https://github.com/xxjwxc/gormt) | 1004 | 18 | 2019/05/05 | 1 week ago | Mysql database to golang gorm struct. |
| [go-queryset](https://github.com/jirfag/go-queryset) | 585 | 20 | 2017/09/03 | 7 months ago | 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM. |
| [qbs](https://github.com/coocood/qbs) | 552 | 45 | 2013/02/02 | 3 years ago | Stands for Query By Struct. A Go ORM. |
| [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) | 527 | 15 | 2017/12/27 | 1 week ago | A flexible and powerful SQL string builder library plus a zero-config ORM. |
| [rel](https://github.com/go-rel/rel) | 302 | 8 | 2019/10/06 | 2 months ago | Modern Database Access Layer for Golang - Testable, Extendable and Crafted Into a Clean and Elegant API. |
| [zoom](https://github.com/albrow/zoom) | 271 | 17 | 2013/07/17 | 9 months ago | Blazing-fast datastore and querying engine built on Redis. |
| [grimoire](https://github.com/Fs02/grimoire) | 150 | 5 | 2018/03/05 | 6 months ago | Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3). |
| [gosql](https://github.com/rushteam/gosql) | 140 | 5 | 2020/04/27 | 2 months ago | A easy ORM for mysql. |
| [gorm](https://github.com/jinzhu/gorm) | 137 | 10 | 2013/10/25 | 5 months ago | The fantastic ORM library for Golang, aims to be developer friendly. |
| [go-store](https://github.com/gosuri/go-store) | 100 | 9 | 2015/03/22 | 4 years ago | Simple and fast Redis backed key-value store library for Go. |
| [marlow](https://github.com/dadleyy/marlow) | 83 | 5 | 2017/09/15 | 6 months ago | Generated ORM from project structs for compile time safety assurances. |
| [go-firestorm](https://github.com/jschoedt/go-firestorm) | 22 | 1 | 2018/12/04 | 1 month ago | A simple ORM for Google/Firebase Cloud Firestore. |
| [lore](https://github.com/abrahambotros/lore) | 6 | 1 | 2017/04/29 | 3 years ago | Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go. |
| [marlow](https://github.com/marlow/marlow) | 4 | 2 | 2020/08/11 | 6 months ago | Generated ORM from project structs for compile time safety assurances. |
| [rel](https://github.com/Fs02/rel) | 1 | 0 | 2019/10/06 | 4 months ago | Golang SQL Repository Layer for Clean (Onion) Architecture. |

## Package Management
        
*Unofficial libraries for package and dependency management.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [dep](https://github.com/golang/dep) | 13208 | 273 | 2016/10/07 | 5 months ago | Go dependency tool. |
| [glide](https://github.com/Masterminds/glide) | 8079 | 194 | 2014/07/09 | 10 months ago | Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip. |
| [godep](https://github.com/tools/godep) | 5628 | 148 | 2013/05/01 | 2 years ago | dependency tool for go, godep helps build packages reproducibly by fixing their dependencies. |
| [govendor](https://github.com/kardianos/govendor) | 5010 | 101 | 2015/04/12 | 11 months ago | Go Package Manager. Go vendor tool that works with the standard vendor file. |
| [gopm](https://github.com/gpmgo/gopm) | 2497 | 84 | 2013/05/15 | 1 year ago | Go Package Manager. |
| [gom](https://github.com/mattn/gom) | 1390 | 35 | 2013/09/11 | 1 year ago | Go Manager - bundle for go. |
| [gpm](https://github.com/pote/gpm) | 1200 | 31 | 2013/09/05 | 3 years ago | Barebones dependency manager for Go. |
| [goop](https://github.com/petejkim/goop) | 780 | 37 | 2014/06/18 | 5 years ago | Simple dependency manager for Go (golang), inspired by Bundler. |
| [modgv](https://github.com/lucasepe/modgv) | 357 | 7 | 2020/09/12 | 4 months ago | Converts 'go mod graph' output into Graphviz's DOT language. |
| [nut](https://github.com/jingweno/nut) | 242 | 6 | 2015/01/23 | 5 years ago | Vendor Go dependencies. |
| [nut](https://github.com/owenthereal/nut) | 241 | 6 | 2015/01/23 | 5 years ago | Vendor Go dependencies. |
| [johnny-deps](https://github.com/VividCortex/johnny-deps) | 215 | 22 | 2013/07/19 | 2 months ago | Minimal dependency version using Git. |
| [gigo](https://github.com/LyricalSecurity/gigo) | 199 | 5 | 2015/05/01 | 2 years ago | PIP-like dependency tool for golang, with support for private repositories and hashes. |
| [VenGO](https://github.com/DamnWidget/VenGO) | 121 | 9 | 2014/10/17 | 4 years ago | create and manage exportable isolated go virtual environments. |
| [mvn-golang](https://github.com/raydac/mvn-golang) | 115 | 11 | 2016/03/24 | 3 hours ago | plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure. |
| [gop](https://github.com/lunny/gop) | 49 | 7 | 2017/02/18 | 1 year ago | Build and manage your Go applications out of GOPATH. |

## Performance
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [jaeger](https://github.com/jaegertracing/jaeger) | 12807 | 336 | 2016/04/15 | 8 hours ago | A distributed tracing system. |
| [profile](https://github.com/pkg/profile) | 1390 | 39 | 2014/10/22 | 3 months ago | Simple profiling support package for Go. |
| [statsviz](https://github.com/arl/statsviz) | 995 | 15 | 2020/08/14 | 30 minutes ago | Live visualization of your Go application runtime statistics. |
| [pixie](https://github.com/pixie-labs/pixie) | 542 | 32 | 2020/02/27 | 3 days ago | No instrumentation tracing for Golang applications via eBPF. |
| [tracer](https://github.com/kamilsk/tracer) | 41 | 3 | 2019/06/22 | 8 months ago | Simple, lightweight tracing. |

## Query Language
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [graphql](https://github.com/graphql-go/graphql) | 7356 | 152 | 2015/07/19 | 1 week ago | Implementation of GraphQL for Go. |
| [graphql-go](https://github.com/graph-gophers/graphql-go) | 3605 | 88 | 2016/10/18 | 2 days ago | GraphQL server with a focus on ease of use. |
| [gojsonq](https://github.com/thedevsaddam/gojsonq) | 1619 | 33 | 2018/05/19 | 3 weeks ago | A simple Go package to Query over JSON Data. |
| [dasel](https://github.com/TomWright/dasel) | 723 | 5 | 2020/09/22 | 2 days ago | Query and update data structures using selectors from the command line. Comparable to jq/yq but supports JSON, YAML, TOML and XML with zero runtime dependencies. |
| [jsonql](https://github.com/elgs/jsonql) | 243 | 9 | 2015/12/29 | 2 months ago | JSON query expression library in Golang. |
| [rql](https://github.com/a8m/rql) | 187 | 7 | 2018/06/05 | 5 months ago | Resource Query Language for REST API. |
| [graphql](https://github.com/tmc/graphql) | 52 | 10 | 2015/04/18 | 3 years ago | graphql parser + utilities. |
| [jsonslice](https://github.com/bhmj/jsonslice) | 50 | 0 | 2018/05/02 | 4 months ago | Jsonpath queries with advanced filters. |
| [api-fu](https://github.com/ccbrown/api-fu) | 24 | 3 | 2019/07/30 | 3 weeks ago | Comprehensive GraphQL implementation. |
| [straf](https://github.com/SonicRoshan/straf) | 22 | 1 | 2019/08/16 | 9 months ago | Easily Convert Golang structs to GraphQL objects. |
| [rest-query-parser](https://github.com/timsolov/rest-query-parser) | 15 | 1 | 2020/02/10 | 1 month ago | Query Parser for REST API. Filtering, validations, both `AND`, `OR` operations are supported directly in the query. |
| [gws](https://github.com/Zaba505/gws) | 4 | 1 | 2020/06/08 | 5 months ago | Apollos' "GraphQL over Websocket" client and server implementation. |
| [jsonpath](https://github.com/AsaiYusuke/jsonpath) | 4 | 1 | 2020/11/29 | 1 week ago | A query library for retrieving part of JSON based on JSONPath syntax. |

## Resource Embedding
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [packr](https://github.com/gobuffalo/packr) | 3164 | 46 | 2017/03/15 | 1 month ago | The simple and easy way to embed static files into Go binaries. |
| [statik](https://github.com/rakyll/statik) | 3112 | 49 | 2014/02/04 | 3 months ago | Embeds static files into a Go executable. |
| [go.rice](https://github.com/GeertJohan/go.rice) | 2159 | 55 | 2013/10/23 | 2 months ago | go.rice is a Go package that makes working with resources such as HTML, JS, CSS, images, and templates very easy. |
| [vfsgen](https://github.com/shurcooL/vfsgen) | 922 | 21 | 2015/05/18 | 5 months ago | Generates a vfsdata.go file that statically implements the given virtual filesystem. |
| [esc](https://github.com/mjibson/esc) | 589 | 15 | 2014/01/26 | 1 year ago | Embeds files into Go programs and provides http.FileSystem interfaces to them. |
| [fileb0x](https://github.com/UnnoTed/fileb0x) | 580 | 17 | 2016/01/23 | 4 hours ago | Simple tool to embed files in go with focus on "customization" and ease to use. |
| [go-resources](https://github.com/omeid/go-resources) | 172 | 7 | 2015/02/21 | 8 months ago | Unfancy resources embedding with Go. |
| [statics](https://github.com/go-playground/statics) | 62 | 2 | 2015/10/07 | 4 years ago | Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks. |
| [go-embed](https://github.com/pyros2097/go-embed) | 28 | 3 | 2015/12/06 | 1 month ago | Generates go code to embed resource files into your library or executable. |
| [templify](https://github.com/wlbr/templify) | 26 | 2 | 2016/05/22 | 1 year ago | Embed external template files into Go code to create single file binaries. |
| [mule](https://github.com/wlbr/mule) | 9 | 1 | 2020/01/17 | 6 months ago | Embed external resources like images, movies ... into Go source code to create single file binaries using `go generate`. Focussed on simplicity. |

## Science and Data Analysis
        
*Libraries for scientific computing and data analyzing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gonum](https://github.com/gonum/gonum) | 4593 | 114 | 2017/03/25 | 2 days ago | Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more. |
| [stats](https://github.com/montanaflynn/stats) | 1910 | 49 | 2014/12/16 | 1 month ago | Statistics package with common functions missing from the Golang standard library. |
| [plot](https://github.com/gonum/plot) | 1816 | 60 | 2013/07/23 | 2 days ago | gonum/plot provides an API for building and drawing plots in Go. |
| [gosl](https://github.com/cpmech/gosl) | 1557 | 74 | 2015/02/09 | 3 weeks ago | Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more. |
| [streamtools](https://github.com/nytlabs/streamtools) | 1314 | 71 | 2013/07/05 | 5 years ago | general purpose, graphical tool for dealing with streams of data. |
| [go-dsp](https://github.com/mjibson/go-dsp) | 708 | 29 | 2011/11/02 | 2 years ago | Digital Signal Processing for Go. |
| [chart](https://github.com/vdobler/chart) | 666 | 45 | 2011/06/27 | 8 months ago | Simple Chart Plotting library for Go. Supports many graphs types. |
| [goraph](https://github.com/gyuho/goraph) | 637 | 40 | 2014/02/27 | 3 years ago | Pure Go graph theory library(data structure, algorith visualization). |
| [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) | 443 | 23 | 2018/10/01 | 3 days ago | Dataframes for machine-learning and statistics (similar to pandas). |
| [graph](https://github.com/yourbasic/graph) | 419 | 20 | 2017/04/27 | 3 weeks ago | Library of basic graph algorithms. |
| [orb](https://github.com/paulmach/orb) | 366 | 21 | 2016/03/28 | 1 week ago | 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support. |
| [ewma](https://github.com/VividCortex/ewma) | 320 | 24 | 2013/07/05 | 2 months ago | Exponentially-weighted moving averages. |
| [calendarheatmap](https://github.com/nikolaydubina/calendarheatmap) | 251 | 8 | 2020/07/01 | 1 month ago | Calendar heatmap in plain Go inspired by Github contribution activity. |
| [gohistogram](https://github.com/VividCortex/gohistogram) | 149 | 17 | 2013/07/02 | 2 months ago | Approximate histograms for data streams. |
| [TextRank](https://github.com/DavidBelicza/TextRank) | 123 | 7 | 2018/01/09 | 11 months ago | TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support. |
| [sparse](https://github.com/james-bowman/sparse) | 108 | 8 | 2017/05/16 | 2 months ago | Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries. |
| [pagerank](https://github.com/alixaxel/pagerank) | 67 | 8 | 2015/08/06 | 3 weeks ago | Weighted PageRank algorithm implemented in Go. |
| [geom](https://github.com/skelterjohn/geom) | 46 | 5 | 2011/06/07 | 3 years ago | 2D geometry for golang. |
| [evaler](https://github.com/soniah/evaler) | 45 | 5 | 2012/09/04 | 2 years ago | Simple floating point arithmetic expression evaluator. |
| [goent](https://github.com/kzahedi/goent) | 23 | 1 | 2017/08/08 | 1 year ago | GO Implementation of Entropy Measures. |
| [triangolatte](https://github.com/tchayen/triangolatte) | 19 | 2 | 2018/07/18 | 1 month ago | 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs. |
| [decimal](https://github.com/db47h/decimal) | 19 | 2 | 2020/05/27 | 7 months ago | Package decimal implements arbitrary-precision decimal floating-point arithmetic. |
| [GoStats](https://github.com/OGFris/GoStats) | 14 | 3 | 2018/07/22 | 2 years ago | GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions. |
| [piecewiselinear](https://github.com/sgreben/piecewiselinear) | 14 | 4 | 2018/10/21 | 2 months ago | Tiny linear interpolation library. |
| [ode](https://github.com/ChristopherRabotin/ode) | 13 | 4 | 2016/11/11 | 3 years ago | Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions. |
| [PiHex](https://github.com/claygod/PiHex) | 13 | 3 | 2016/07/22 | 5 months ago | Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi. |
| [assocentity](https://github.com/ndabAP/assocentity) | 6 | 1 | 2018/12/21 | 3 months ago | Package assocentity returns the average distance from words to a given entity. |
| [go-gt](https://github.com/ThePaw/go-gt) | 5 | 0 | 2015/09/14 | 5 years ago | Graph theory algorithms written in "Go" language. |
| [bradleyterry](https://github.com/seanhagen/bradleyterry) | 4 | 2 | 2019/04/30 | 1 year ago | Provides a Bradley-Terry Model for pairwise comparisons. |
| [rootfinding](https://github.com/khezen/rootfinding) | 4 | 3 | 2018/10/30 | 10 months ago | root-finding algorithms library for finding roots of quadratic functions. |

## Security
        
*Libraries that are used to help make your application more secure.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [lego](https://github.com/go-acme/lego) | 4450 | 101 | 2015/06/08 | 1 hour ago | Pure Go ACME client library and CLI tool (for use with Let's Encrypt). |
| [cameradar](https://github.com/Ullaakut/cameradar) | 2468 | 119 | 2016/05/20 | 3 months ago | Tool and library to remotely hack RTSP streams from surveillance cameras. |
| [memguard](https://github.com/awnumar/memguard) | 1858 | 46 | 2017/04/22 | 6 months ago | A pure Go library for handling sensitive values in memory. |
| [acmetool](https://github.com/hlandau/acmetool) | 1831 | 69 | 2015/11/15 | 7 months ago | ACME (Let's Encrypt) client tool with automatic renewal. |
| [secure](https://github.com/unrolled/secure) | 1613 | 37 | 2014/05/20 | 3 days ago | HTTP middleware for Go that facilitates some quick security wins. |
| [themis](https://github.com/cossacklabs/themis) | 1213 | 42 | 2015/05/06 | 1 day ago | high-level cryptographic library for solving typical data security tasks (secure data storage, secure messaging, zero-knowledge proof authentication), available for 14 languages, best fit for multi-platform apps. |
| [acra](https://github.com/cossacklabs/acra) | 707 | 38 | 2016/11/14 | 3 days ago | Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system. |
| [nacl](https://github.com/kevinburke/nacl) | 494 | 13 | 2017/07/20 | 1 week ago | Go implementation of the NaCL set of API's. |
| [firewalld-rest](https://github.com/prashantgupta24/firewalld-rest) | 299 | 8 | 2020/06/12 | 5 months ago | A rest application to dynamically update firewalld rules on a linux server. |
| [badactor](https://github.com/jaredfolkins/badactor) | 292 | 8 | 2014/12/12 | 8 months ago | In-memory, application-driven jailer built in the spirit of fail2ban. |
| [ssh-vault](https://github.com/ssh-vault/ssh-vault) | 274 | 10 | 2016/09/29 | 1 year ago | encrypt/decrypt using ssh keys. |
| [go-password-validator](https://github.com/lane-c-wagner/go-password-validator) | 257 | 8 | 2020/10/14 | 1 month ago | Password validator based on raw cryptographic entropy values. |
| [optimus-go](https://github.com/pjebs/optimus-go) | 255 | 5 | 2015/05/05 | 9 months ago | ID hashing and Obfuscation using Knuth's Algorithm. |
| [passlib](https://github.com/hlandau/passlib) | 240 | 11 | 2014/12/21 | 1 year ago | Futureproof password hashing library. |
| [go-yara](https://github.com/hillu/go-yara) | 196 | 19 | 2015/01/25 | 1 week ago | Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)". |
| [simple-scrypt](https://github.com/elithrar/simple-scrypt) | 166 | 7 | 2015/04/14 | 5 months ago | Scrypt package with a simple, obvious API and automatic cost calibration built-in. |
| [argon2pw](https://github.com/raja/argon2pw) | 85 | 4 | 2018/03/13 | 2 years ago | Argon2 password hash generation with constant-time password comparison. |
| [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) | 39 | 1 | 2017/10/19 | 2 months ago | A probably paranoid package for securely hashing and encrypting passwords. |
| [certificates](https://github.com/mvmaasakkers/certificates) | 20 | 1 | 2019/03/04 | 2 months ago | An opinionated tool for generating tls certificates. |
| [go-generate-password](https://github.com/m1/go-generate-password) | 17 | 1 | 2019/11/14 | 3 months ago | Password generator that can be used on the cli or as a library. |
| [secureio](https://github.com/xaionaro-go/secureio) | 15 | 2 | 2018/12/25 | 7 months ago | An keyexchanging+authenticating+encrypting wrapper and multiplexer for `io.ReadWriteCloser` based on XChaCha20-poly1305, ECDH and ED25519. |
| [goArgonPass](https://github.com/dwin/goArgonPass) | 14 | 2 | 2018/05/30 | 2 months ago | Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations. |
| [argon2-hashing](https://github.com/andskur/argon2-hashing) | 12 | 2 | 2019/01/02 | 10 months ago | light wrapper around Go's argon2 package that closely mirrors with Go's standard library Bcrypt and simple-scrypt package. |
| [sslmgr](https://github.com/adrianosela/sslmgr) | 10 | 2 | 2019/04/02 | 1 year ago | SSL certificates made easy with a high level wrapper around acme/autocert. |

## Serialization
        
*Libraries and tools for binary serialization.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go](https://github.com/json-iterator/go) | 8885 | 228 | 2016/11/30 | 3 weeks ago | High-performance 100% compatible drop-in replacement of "encoding/json". |
| [protobuf](https://github.com/golang/protobuf) | 7435 | 217 | 2014/11/23 | 2 months ago | Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers. |
| [protobuf](https://github.com/gogo/protobuf) | 4384 | 95 | 2014/12/03 | 2 weeks ago | Protocol Buffers for Go with Gadgets. |
| [mapstructure](https://github.com/mitchellh/mapstructure) | 4195 | 65 | 2013/05/20 | 3 weeks ago | Go library for decoding generic map values into native Go structures. |
| [go](https://github.com/ugorji/go) | 1501 | 51 | 2013/05/30 | 1 day ago | High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support. |
| [colfer](https://github.com/pascaldekloe/colfer) | 590 | 34 | 2015/09/05 | 2 weeks ago | Code generation for the Colfer binary format. |
| [csvutil](https://github.com/jszwec/csvutil) | 470 | 7 | 2017/10/30 | 4 days ago | High Performance, idiomatic CSV record encoding and decoding to native Go structures. |
| [go-capnproto](https://github.com/glycerine/go-capnproto) | 277 | 11 | 2013/11/07 | 1 year ago | Cap'n Proto library and parser for go. |
| [cbor](https://github.com/fxamacker/cbor) | 236 | 6 | 2019/05/15 | 2 months ago | Small, safe, and easy CBOR encoding and decoding library. |
| [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) | 145 | 9 | 2012/12/23 | 2 years ago | GoLang library for working with PHP session format and PHP Serialize/Unserialize functions. |
| [structomap](https://github.com/danhper/structomap) | 117 | 7 | 2015/05/13 | 1 year ago | Library to easily and dynamically generate maps from static structures. |
| [bambam](https://github.com/glycerine/bambam) | 61 | 4 | 2014/09/17 | 4 years ago | generator for Cap'n Proto schemas from go. |
| [asn1](https://github.com/Logicalis/asn1) | 47 | 8 | 2016/02/29 | 1 year ago | Asn.1 BER and DER encoding library for golang. |
| [binstruct](https://github.com/ghostiam/binstruct) | 28 | 1 | 2018/10/23 | 1 year ago | Golang binary decoder for mapping data into the structure. |
| [fwencoder](https://github.com/o1egl/fwencoder) | 14 | 1 | 2017/12/25 | 1 year ago | Fixed width file parser (encoding and decoding library) for Go. |
| [elastic](https://github.com/epiclabs-io/elastic) | 13 | 0 | 2020/02/25 | 11 months ago | Convert slices, maps or any other unknown value across different types at run-time, no matter what. |
| [pletter](https://github.com/vimeda/pletter) | 12 | 4 | 2019/07/09 | 4 months ago | A standard way to wrap a proto message for message brokers. |
| [bel](https://github.com/csweichel/bel) | 12 | 2 | 2019/02/20 | 6 months ago | Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC. |
| [bel](https://github.com/32leaves/bel) | 8 | 1 | 2019/02/20 | 1 year ago | Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC. |
| [fixedwidth](https://github.com/huydang284/fixedwidth) | 5 | 1 | 2019/08/11 | 1 year ago | Fixed-width text formatting (UTF-8 supported). |
| [go-lctree](https://github.com/sbourlon/go-lctree) | 2 | 1 | 2020/05/04 | 8 months ago | Provides a CLI and primitives to serialize and deserialize [LeetCode binary trees](https://support.leetcode.com/hc/en-us/articles/360011883654-What-does-1-null-2-3-mean-in-binary-tree-representation). |
| [unitpacking](https://github.com/recolude/unitpacking) | 1 | 1 | 2021/01/17 | 2 weeks ago | Library to pack unit vectors into as fewest bytes as possible. |

## Server Applications
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [etcd](https://github.com/etcd-io/etcd) | 34756 | 1365 | 2013/07/06 | 1 day ago | Highly-available key value store for shared configuration and service discovery. |
| [caddy](https://github.com/caddyserver/caddy) | 31919 | 747 | 2015/01/13 | 2 days ago | Caddy is an alternative, HTTP/2 web server that's easy to configure and use. |
| [minio](https://github.com/minio/minio) | 26064 | 568 | 2015/01/14 | 12 hours ago | Minio is a distributed object storage server. |
| [roadrunner](https://github.com/spiral/roadrunner) | 4825 | 141 | 2017/12/26 | 2 days ago | High-performance PHP application server, load-balancer and process manager. |
| [devd](https://github.com/cortesi/devd) | 3089 | 69 | 2015/09/27 | 9 months ago | Local webserver for developers. |
| [sftpgo](https://github.com/drakkan/sftpgo) | 2296 | 58 | 2019/07/20 | 1 day ago | Fully featured and highly configurable SFTP server with optional FTP/S and WebDAV support. It can serve local filesystem and Cloud Storage backends such as S3 and Google Cloud Storage. |
| [algernon](https://github.com/xyproto/algernon) | 1780 | 47 | 2015/03/10 | 3 weeks ago | HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber. |
| [flagr](https://github.com/checkr/flagr) | 1479 | 76 | 2017/10/03 | 2 weeks ago | Flagr is an open-source feature flagging and A/B testing service. |
| [fider](https://github.com/getfider/fider) | 1446 | 34 | 2017/01/17 | 1 week ago | Fider is an open platform to collect and organize customer feedback. |
| [flipt](https://github.com/markphelps/flipt) | 1402 | 16 | 2016/11/05 | 2 days ago | A self contained feature flag solution written in Go and Vue. |
| [trickster](https://github.com/tricksterproxy/trickster) | 1288 | 38 | 2018/03/29 | 2 days ago | HTTP reverse proxy cache and time series accelerator. |
| [discovery](https://github.com/bilibili/discovery) | 1281 | 58 | 2018/04/20 | 1 month ago | A registry for resilient mid-tier load balancing and failover. |
| [trickster](https://github.com/Comcast/trickster) | 1053 | 35 | 2018/03/29 | 11 months ago | HTTP reverse proxy cache and time series accelerator. |
| [jackal](https://github.com/ortuman/jackal) | 874 | 40 | 2017/11/13 | 4 weeks ago | An XMPP server written in Go. |
| [dudeldu](https://github.com/krotik/dudeldu) | 120 | 3 | 2016/09/07 | 1 year ago | A simple SHOUTcast server. |
| [lets-proxy2](https://github.com/rekby/lets-proxy2) | 47 | 4 | 2019/04/12 | 2 months ago | Reverse proxy for handle https with issue certificates in fly from lets-encrypt. |
| [go-feature-flag](https://github.com/thomaspoignant/go-feature-flag) | 29 | 0 | 2020/12/11 | 1 week ago | A feature flag solution, with only a YAML file in the backend (S3, GitHub, HTTP, local file ...), no server to install, just add a file in a central system and refer to it. |
| [psql-streamer](https://github.com/blind-oracle/psql-streamer) | 26 | 4 | 2019/04/28 | 11 months ago | Stream database events from PostgreSQL to Kafka. |
| [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) | 18 | 1 | 2018/10/23 | 5 months ago | Nginx log parser and exporter to Prometheus. |
| [protoxy](https://github.com/camgraff/protoxy) | 12 | 2 | 2020/09/03 | 3 months ago | A proxy server that converts JSON request bodies to Protocol Buffers. |
| [simple-jwt-provider](https://github.com/leberKleber/simple-jwt-provider) | 12 | 2 | 2019/12/18 | 3 weeks ago | Simple and lightweight provider which exhibits JWTs, supports login, password-reset (via mail) and user management. |
| [go-proxy-cache](https://github.com/fabiocicerchia/go-proxy-cache) | 6 | 1 | 2020/11/12 | 1 week ago | Simple Reverse Proxy with Caching, written in Go, using Redis. |
| [riemann-relay](https://github.com/blind-oracle/riemann-relay) | 0 | 1 | 2019/04/23 | 1 year ago | Relay to load-balance Riemann events and/or convert them to Carbon. |

## Stream Processing
        
*Libraries and tools for stream processing and reactive programming.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-streams](https://github.com/reugn/go-streams) | 599 | 18 | 2019/04/30 | 1 month ago | Go stream processing library. |
| [machine](https://github.com/whitaker-io/machine) | 76 | 3 | 2020/10/13 | 3 days ago | Go library for writing and generating stream workers with built in metrics and traceability. |
| [stream](https://github.com/youthlin/stream) | 25 | 2 | 2020/11/12 | 2 months ago | Go Stream, like Java 8 Stream: Filter/Map/FlatMap/Peek/Sorted/ForEach/Reduce... |

## Template Engines
        
*Libraries and tools for templating and lexing.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gofpdf](https://github.com/jung-kurt/gofpdf) | 3743 | 101 | 2015/03/13 | 1 year ago | PDF document generator with high level support for text, drawing and images. |
| [sprig](https://github.com/Masterminds/sprig) | 2156 | 30 | 2013/11/22 | 2 days ago | Useful template functions for Go templates. |
| [quicktemplate](https://github.com/valyala/quicktemplate) | 1978 | 56 | 2016/03/06 | 4 months ago | Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it. |
| [pongo2](https://github.com/flosch/pongo2) | 1885 | 59 | 2013/08/23 | 2 months ago | Django-like template-engine for Go. |
| [hero](https://github.com/shiyanhui/hero) | 1435 | 43 | 2017/01/15 | 1 year ago | Hero is a handy, fast and powerful go template engine. |
| [mustache](https://github.com/hoisie/mustache) | 1006 | 35 | 2009/12/30 | 1 month ago | Go implementation of the Mustache template language. |
| [amber](https://github.com/eknkc/amber) | 867 | 19 | 2012/10/31 | 4 months ago | Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade. |
| [ace](https://github.com/yosssi/ace) | 797 | 23 | 2014/07/13 | 2 years ago | Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold. |
| [gorazor](https://github.com/sipin/gorazor) | 772 | 57 | 2014/05/01 | 2 months ago | Razor view engine for Golang. |
| [jet](https://github.com/CloudyKit/jet) | 748 | 22 | 2016/03/31 | 3 days ago | Jet template engine. |
| [ego](https://github.com/benbjohnson/ego) | 472 | 17 | 2014/02/23 | 6 days ago | Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled. |
| [fasttemplate](https://github.com/valyala/fasttemplate) | 463 | 18 | 2015/08/19 | 1 month ago | Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/). |
| [raymond](https://github.com/aymerick/raymond) | 414 | 12 | 2015/04/22 | 5 months ago | Complete handlebars implementation in Go. |
| [maroto](https://github.com/johnfercher/maroto) | 321 | 9 | 2019/05/20 | 1 week ago | A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple. |
| [goview](https://github.com/foolin/goview) | 203 | 5 | 2019/04/14 | 1 month ago | Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application. |
| [soy](https://github.com/robfig/soy) | 155 | 12 | 2013/12/15 | 1 week ago | Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/). |
| [liquid](https://github.com/osteele/liquid) | 115 | 6 | 2017/06/26 | 1 month ago | Go implementation of Shopify Liquid templates. |
| [velvet](https://github.com/gobuffalo/velvet) | 73 | 5 | 2016/12/29 | 3 years ago | Complete handlebars implementation in Go. |
| [kasia.go](https://github.com/ziutek/kasia.go) | 71 | 2 | 2010/12/07 | 5 years ago | Templating system for HTML and other text documents - go implementation. |
| [extemplate](https://github.com/dannyvankooten/extemplate) | 32 | 4 | 2018/08/10 | 9 months ago | Tiny wrapper around html/template to allow for easy file-based template inheritance. |
| [damsel](https://github.com/dskinner/damsel) | 24 | 4 | 2012/05/02 | 4 years ago | Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others. |
| [gospin](https://github.com/m1/gospin) | 21 | 2 | 2019/02/22 | 9 months ago | Article spinning and spintax/spinning syntax engine, useful for A/B, testing pieces of text/articles and creating more natural conversations. |

## Testing
        
*Libraries for testing codebases and generating test data.*

### Testing Frameworks
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [testify](https://github.com/stretchr/testify) | 12523 | 176 | 2012/10/16 | 3 days ago | Sacred extension to the standard go testing package. |
| [go-cmp](https://github.com/google/go-cmp) | 2140 | 27 | 2017/07/07 | 1 week ago | Package for comparing Go values in tests. |
| [httpexpect](https://github.com/gavv/httpexpect) | 1614 | 38 | 2016/04/29 | 1 day ago | Concise, declarative, and easy to use end-to-end HTTP and REST API testing. |
| [godog](https://github.com/cucumber/godog) | 1257 | 96 | 2015/06/10 | 5 days ago | Cucumber or Behat like BDD framework for Go. |
| [godog](https://github.com/DATA-DOG/godog) | 895 | 30 | 2015/06/10 | 1 year ago | Cucumber or Behat like BDD framework for Go. |
| [goblin](https://github.com/franela/goblin) | 755 | 15 | 2013/09/19 | 1 month ago | Mocha like testing framework fo Go. |
| [baloo](https://github.com/h2non/baloo) | 699 | 11 | 2016/05/29 | 2 years ago | Expressive and versatile end-to-end HTTP API testing made easy. |
| [testfixtures](https://github.com/go-testfixtures/testfixtures) | 633 | 5 | 2016/04/05 | 1 month ago | A helper for Rails' like test fixtures to test database applications. |
| [go-vcr](https://github.com/dnaeon/go-vcr) | 484 | 8 | 2015/12/14 | 3 months ago | Record and replay your HTTP interactions for fast, deterministic and accurate tests. |
| [gnomock](https://github.com/orlangure/gnomock) | 385 | 10 | 2020/01/31 | 2 weeks ago | integration testing with real dependencies (database, cache, even Kubernetes or AWS) running in Docker, without mocks. |
| [go-mutesting](https://github.com/zimmski/go-mutesting) | 382 | 8 | 2014/12/26 | 1 year ago | Mutation testing for Go source code. |
| [gofight](https://github.com/appleboy/gofight) | 360 | 13 | 2016/03/29 | 2 weeks ago | API Handler Testing for Golang Router framework. |
| [goc](https://github.com/qiniu/goc) | 298 | 13 | 2020/05/07 | 5 days ago | Goc is a comprehensive coverage testing system for The Go Programming Language. |
| [frisby](https://github.com/hofstadter-io/frisby) | 265 | 8 | 2015/09/15 | 11 months ago | REST API testing framework. |
| [frisby](https://github.com/verdverm/frisby) | 257 | 8 | 2015/09/15 | 11 months ago | REST API testing framework. |
| [go-carpet](https://github.com/msoap/go-carpet) | 210 | 4 | 2016/02/28 | 2 months ago | Tool for viewing test coverage in terminal. |
| [gotest.tools](https://github.com/gotestyourself/gotest.tools) | 198 | 6 | 2017/08/08 | 1 week ago | A collection of packages to augment the go testing package and support common patterns. |
| [charlatan](https://github.com/percolate/charlatan) | 195 | 44 | 2017/10/06 | 1 year ago | Tool to generate fake interface implementations for tests. |
| [commander](https://github.com/commander-cli/commander) | 182 | 7 | 2019/02/22 | 2 days ago | Tool for testing cli applications on windows, linux and osx. |
| [endly](https://github.com/viant/endly) | 174 | 15 | 2017/08/28 | 3 months ago | Declarative end to end functional testing. |
| [commander](https://github.com/SimonBaeumer/commander) | 154 | 9 | 2019/02/22 | 7 months ago | Tool for testing cli applications on windows, linux and osx. |
| [cupaloy](https://github.com/bradleyjkemp/cupaloy) | 129 | 2 | 2017/08/07 | 1 month ago | Simple snapshot testing addon for your test framework. |
| [dbcleaner](https://github.com/khaiql/dbcleaner) | 123 | 2 | 2017/01/17 | 11 months ago | Clean database for testing purpose, inspired by `database_cleaner` in Ruby. |
| [go-testdeep](https://github.com/maxatome/go-testdeep) | 116 | 2 | 2018/05/26 | 5 hours ago | Extremely flexible golang deep comparison, extends the go testing package. |
| [gospec](https://github.com/luontola/gospec) | 114 | 4 | 2009/11/24 | 6 years ago | BDD-style testing framework for the Go programming language. |
| [embedded-postgres](https://github.com/fergusstrange/embedded-postgres) | 93 | 3 | 2019/11/16 | 2 weeks ago | Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test. |
| [wstest](https://github.com/posener/wstest) | 83 | 2 | 2017/03/31 | 1 month ago | Websocket client for unit-testing a websocket http.Handler. |
| [gocrest](https://github.com/corbym/gocrest) | 77 | 3 | 2017/12/23 | 1 month ago | Composable hamcrest-like matchers for Go assertions. |
| [testcase](https://github.com/adamluzsi/testcase) | 58 | 3 | 2019/04/22 | 1 day ago | Idiomatic testing framework for Behavior Driven Development. |
| [jsonassert](https://github.com/kinbiko/jsonassert) | 57 | 0 | 2018/10/26 | 8 months ago | Package for verifying that your JSON payloads are serialized correctly. |
| [gospecify](https://github.com/stesla/gospecify) | 53 | 6 | 2009/11/20 | 9 years ago | This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec. |
| [restit](https://github.com/go-restit/restit) | 52 | 6 | 2014/06/25 | 1 year ago | Go micro framework to help writing RESTful API integration test. |
| [gomatch](https://github.com/jfilipczyk/gomatch) | 40 | 2 | 2019/01/27 | 1 month ago | library created for testing JSON against patterns. |
| [dsunit](https://github.com/viant/dsunit) | 38 | 9 | 2016/06/13 | 1 year ago | Datastore testing for SQL, NoSQL, structured files. |
| [covergates](https://github.com/covergates/covergates) | 36 | 4 | 2020/05/29 | 1 month ago | Self-hosted code coverage report review and management service. |
| [assert](https://github.com/go-playground/assert) | 28 | 1 | 2015/07/20 | 3 months ago | Basic Assertion Library used along side native go testing, with building blocks for custom assertions. |
| [go-hit](https://github.com/Eun/go-hit) | 27 | 1 | 2019/06/04 | 3 days ago | Hit is an http integration test framework written in golang. |
| [hamcrest](https://github.com/rdrdr/hamcrest) | 26 | 3 | 2010/12/22 | 1 month ago | fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results. |
| [flute](https://github.com/suzuki-shunsuke/flute) | 15 | 0 | 2019/07/06 | 1 week ago | HTTP client testing framework. |
| [schema](https://github.com/jgroeneveld/schema) | 14 | 1 | 2015/08/13 | 1 year ago | Quick and easy expression matching for JSON schemas used in requests and responses. |
| [badio](https://github.com/cavaliercoder/badio) | 10 | 1 | 2016/02/11 | 5 years ago | Extensions to Go's `testing/iotest` package. |
| [biff](https://github.com/fulldump/biff) | 10 | 1 | 2018/03/28 | 6 months ago | Bifurcation testing framework, BDD compatible. |
| [testsql](https://github.com/zhulongcheng/testsql) | 10 | 2 | 2018/09/22 | 1 year ago | Generate test data from SQL files before testing and clear it after finished. |
| [gogiven](https://github.com/corbym/gogiven) | 9 | 4 | 2017/12/31 | 3 years ago | YATSPEC-like BDD testing framework for Go. |
| [gosuite](https://github.com/pavlo/gosuite) | 9 | 1 | 2016/10/15 | 4 years ago | Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests. |
| [tt](https://github.com/vcaesar/tt) | 5 | 1 | 2018/04/03 | 1 month ago | Simple and colorful test tools. |
| [stop-and-go](https://github.com/elgohr/stop-and-go) | 5 | 0 | 2020/11/06 | 2 weeks ago | Testing helper for concurrency. |
| [trial](https://github.com/jgroeneveld/trial) | 4 | 0 | 2015/06/18 | 1 year ago | Quick and easy extendable assertions without introducing much boilerplate. |

### Mock
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [chromedp](https://github.com/chromedp/chromedp) | 5890 | 149 | 2017/01/24 | 3 weeks ago | a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol. |
| [mock](https://github.com/golang/mock) | 5167 | 80 | 2015/06/12 | 14 hours ago | Mocking framework for the Go programming language. |
| [go-fuzz](https://github.com/dvyukov/go-fuzz) | 3836 | 90 | 2015/04/15 | 1 week ago | Randomized testing system. |
| [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) | 3176 | 34 | 2014/02/07 | 4 weeks ago | Mock SQL driver for testing database interactions. |
| [selenoid](https://github.com/aerokube/selenoid) | 1818 | 96 | 2016/08/22 | 2 weeks ago | alternative Selenium hub server that launches browsers within containers. |
| [hoverfly](https://github.com/SpectoLabs/hoverfly) | 1691 | 63 | 2015/11/30 | 1 week ago | HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI. |
| [rod](https://github.com/go-rod/rod) | 1342 | 22 | 2020/01/21 | 6 days ago | A Devtools driver to make web automation and scraping easy. |
| [gock](https://github.com/h2non/gock) | 1154 | 17 | 2016/03/02 | 3 weeks ago | Versatile HTTP mocking made easy. |
| [gofuzz](https://github.com/google/gofuzz) | 973 | 27 | 2014/07/31 | 1 month ago | Library for populating go objects with random values. |
| [httpmock](https://github.com/jarcoal/httpmock) | 972 | 10 | 2014/02/24 | 2 weeks ago | Easy mocking of HTTP responses from external resources. |
| [cdp](https://github.com/mafredri/cdp) | 514 | 18 | 2017/03/12 | 4 months ago | Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it. |
| [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) | 502 | 8 | 2014/05/21 | 1 week ago | Tool for generating self-contained mock objects. |
| [minimock](https://github.com/gojuno/minimock) | 344 | 11 | 2016/08/03 | 3 weeks ago | Mock generator for Go interfaces. |
| [go-txdb](https://github.com/DATA-DOG/go-txdb) | 337 | 6 | 2015/07/08 | 2 months ago | Single transaction based database driver mainly for testing purposes. |
| [rod](https://github.com/ysmood/rod) | 273 | 8 | 2020/01/21 | 7 months ago | A chrome devtools controller that is easy and safe to use. |
| [ggr](https://github.com/aerokube/ggr) | 268 | 24 | 2016/06/16 | 1 week ago | a lightweight server that routes and proxies Selenium WebDriver requests to multiple Selenium hubs. |
| [playwright-go](https://github.com/mxschmitt/playwright-go) | 251 | 13 | 2020/08/16 | 3 weeks ago | browser automation library to control Chromium, Firefox and WebKit with a single API. |
| [tavor](https://github.com/zimmski/tavor) | 226 | 13 | 2014/05/18 | 2 years ago | Generic fuzzing and delta-debugging framework. |
| [govcr](https://github.com/seborama/govcr) | 92 | 2 | 2016/07/10 | 1 year ago | HTTP mock for Golang: record and replay HTTP interactions for offline testing. |
| [timex](https://github.com/cabify/timex) | 53 | 78 | 2020/01/02 | 6 months ago | A test-friendly replacement for the native `time` package. |
| [mockhttp](https://github.com/tv42/mockhttp) | 22 | 1 | 2011/06/11 | 6 years ago | Mock object for Go http.ResponseWriter. |
| [go-localstack](https://github.com/elgohr/go-localstack) | 8 | 0 | 2020/03/18 | 1 day ago | Tool for using localstack in AWS testing. |
| [mockit](https://github.com/pasdam/mockit) | 5 | 1 | 2020/01/01 | 2 weeks ago | Allows functions and method easy mocking, without defining new types; it's similar to Mockito for Java. |

### Fail injection
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [failpoint](https://github.com/pingcap/failpoint) | 561 | 78 | 2019/04/02 | 4 months ago | An implementation of [failpoints](http://www.freebsd.org/cgi/man.cgi?query=fail) for Golang. |

## Text Processing
        
*Libraries for parsing and manipulating texts.*

### Specific Formats
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [colly](https://github.com/gocolly/colly) | 13099 | 306 | 2017/09/29 | 50 minutes ago | Fast and Elegant Scraping Framework for Gophers. |
| [goquery](https://github.com/PuerkitoBio/goquery) | 9794 | 262 | 2012/08/29 | 3 weeks ago | GoQuery brings a syntax and a set of features similar to jQuery to the Go language. |
| [blackfriday](https://github.com/russross/blackfriday) | 4642 | 93 | 2011/05/27 | 1 week ago | Markdown processor in Go. |
| [sh](https://github.com/mvdan/sh) | 3480 | 55 | 2016/01/16 | 1 week ago | Shell parser and formatter. |
| [toml](https://github.com/BurntSushi/toml) | 3412 | 88 | 2013/02/26 | 3 months ago | TOML configuration format (encoder/decoder with reflection). |
| [go-humanize](https://github.com/dustin/go-humanize) | 2508 | 34 | 2012/01/13 | 2 months ago | Formatters for time, numbers, and memory size to human readable format. |
| [bluemonday](https://github.com/microcosm-cc/bluemonday) | 1767 | 29 | 2013/11/20 | 1 month ago | HTML Sanitizer. |
| [gofeed](https://github.com/mmcdole/gofeed) | 1550 | 43 | 2016/01/23 | 1 day ago | Parse RSS and Atom feeds in Go. |
| [inject](https://github.com/facebookarchive/inject) | 1305 | 45 | 2013/10/21 | 2 years ago | Package inject provides a reflect based injector. |
| [go-toml](https://github.com/pelletier/go-toml) | 921 | 33 | 2013/02/24 | 21 hours ago | Go library for the TOML format with query support and handy cli tools. |
| [commonregex](https://github.com/mingrammer/commonregex) | 719 | 20 | 2017/03/23 | 1 year ago | A collection of common regular expressions for Go. |
| [slug](https://github.com/gosimple/slug) | 626 | 11 | 2014/03/31 | 3 months ago | URL-friendly slugify with multiple languages support. |
| [dataflowkit](https://github.com/slotix/dataflowkit) | 453 | 22 | 2017/02/09 | 8 months ago | Web scraping Framework to turn websites into structured data. |
| [mxj](https://github.com/clbanning/mxj) | 444 | 25 | 2014/02/03 | 1 week ago | Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages. |
| [gographviz](https://github.com/awalterschulze/gographviz) | 411 | 12 | 2015/03/14 | 2 months ago | Parses the Graphviz DOT language. |
| [go-runewidth](https://github.com/mattn/go-runewidth) | 334 | 12 | 2013/06/21 | 3 weeks ago | Functions to get fixed width of the character or string. |
| [htmlquery](https://github.com/antchfx/htmlquery) | 328 | 11 | 2017/12/05 | 6 months ago | An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression. |
| [gotext](https://github.com/leonelquinteros/gotext) | 290 | 6 | 2016/06/19 | 3 months ago | GNU gettext utilities for Go. |
| [goq](https://github.com/andrewstuart/goq) | 186 | 7 | 2017/02/20 | 1 year ago | Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery). |
| [goribot](https://github.com/zhshch2002/goribot) | 185 | 11 | 2019/09/08 | 6 months ago | A simple golang spider/scraping framework,build a spider in 3 lines. |
| [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) | 144 | 6 | 2018/05/15 | 4 months ago | Convert HTML to Markdown. Even works with entire websites and can be extended through rules. |
| [omniparser](https://github.com/jf-tech/omniparser) | 140 | 6 | 2020/08/16 | 1 month ago | A versatile ETL library that parses text input (CSV/txt/JSON/XML/EDI/X12/EDIFACT/etc) in streaming fashion and transforms data into JSON output using data-driven schema. |
| [go-nmea](https://github.com/adrianmo/go-nmea) | 138 | 7 | 2015/07/22 | 5 months ago | NMEA parser library for the Go language. |
| [sdp](https://github.com/gortc/sdp) | 107 | 7 | 2016/05/13 | 9 months ago | SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)]. |
| [go-zero-width](https://github.com/trubitsyn/go-zero-width) | 93 | 1 | 2018/06/18 | 6 months ago | Zero-width character detection and removal for Go. |
| [podcast](https://github.com/eduncan911/podcast) | 88 | 5 | 2017/02/02 | 3 months ago | iTunes Compliant and RSS 2. |
| [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) | 80 | 8 | 2016/07/05 | 20 hours ago | Editorconfig file parser and manipulator for Go. |
| [align](https://github.com/Guitarbum722/align) | 68 | 4 | 2017/04/29 | 5 days ago | A general purpose application that aligns text. |
| [go-slugify](https://github.com/mozillazg/go-slugify) | 65 | 2 | 2016/07/16 | 9 months ago | Make pretty slug with multiple languages support. |
| [genex](https://github.com/alixaxel/genex) | 60 | 3 | 2015/03/09 | 1 year ago | Count and expand Regular Expressions into all matching Strings. |
| [go-vcard](https://github.com/emersion/go-vcard) | 53 | 3 | 2017/03/21 | 9 months ago | Parse and format vCard. |
| [goregen](https://github.com/zach-klippenstein/goregen) | 53 | 2 | 2014/12/27 | 1 year ago | Library for generating random strings from regular expressions. |
| [guesslanguage](https://github.com/endeveit/guesslanguage) | 51 | 1 | 2014/12/16 | 3 years ago | Functions to determine the natural language of a unicode text. |
| [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) | 50 | 2 | 2017/11/15 | 10 months ago | Fixed-width text formatting (encoder/decoder with reflection). |
| [allot](https://github.com/sbstjn/allot) | 46 | 1 | 2016/10/16 | 1 year ago | Placeholder and wildcard text parsing for CLI tools and bots. |
| [did](https://github.com/ockam-network/did) | 46 | 10 | 2018/11/02 | 1 month ago | DID (Decentralized Identifiers) Parser and Stringer in Go. |
| [gonameparts](https://github.com/polera/gonameparts) | 33 | 0 | 2015/05/17 | 1 year ago | Parses human names into individual name parts. |
| [slugify](https://github.com/avelino/slugify) | 28 | 2 | 2015/04/13 | 2 years ago | Go slugify application that handles string. |
| [pagser](https://github.com/foolin/pagser) | 24 | 1 | 2020/04/19 | 9 months ago | Pagser is a simple, extensible, configurable parse and deserialize html page to struct based on goquery and struct tags for golang crawler. |
| [codetree](https://github.com/aerogo/codetree) | 17 | 2 | 2016/11/26 | 1 year ago | Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure. |
| [enca](https://github.com/endeveit/enca) | 11 | 1 | 2014/12/17 | 4 years ago | Minimal cgo bindings for [libenca](http://cihar.com/software/enca/). |
| [syndfeed](https://github.com/zhengchun/syndfeed) | 7 | 1 | 2017/04/07 | 2 years ago | A syndication feed for Atom 1.0 and RSS 2.0. |
| [bbConvert](https://github.com/CalebQ42/bbConvert) | 6 | 1 | 2016/04/15 | 4 years ago | Converts bbCode to HTML that allows you to add support for custom bbCode tags. |
| [doi](https://github.com/hscells/doi) | 6 | 1 | 2017/08/02 | 3 years ago | Document object identifier (doi) parser in Go. |
| [ltsv](https://github.com/Wing924/ltsv) | 5 | 1 | 2019/05/12 | 1 year ago | High performance [LTSV (Labeled Tab Separated Value)](http://ltsv.org/) reader for Go. |
| [encoding](https://github.com/mickep76/encoding) | 3 | 1 | 2018/04/06 | 1 year ago | Package provides a generic interface to encoders and decodersa. |

### Utility
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [xurls](https://github.com/mvdan/xurls) | 757 | 19 | 2015/01/12 | 1 month ago | Extract urls from text. |
| [gotabulate](https://github.com/bndr/gotabulate) | 259 | 8 | 2014/08/21 | 5 days ago | Easily pretty-print your tabular data with Go. |
| [radix](https://github.com/yourbasic/radix) | 166 | 5 | 2017/06/09 | 2 years ago | fast string sorting algorithm. |
| [regroup](https://github.com/oriser/regroup) | 73 | 1 | 2020/09/08 | 3 months ago | Match regex expression named groups into go struct using struct tags and automatic parsing. |
| [parth](https://github.com/codemodus/parth) | 39 | 3 | 2015/04/06 | 2 years ago | URL path segmentation parsing. |
| [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) | 31 | 2 | 2018/09/09 | 1 month ago | A sanitization-based swear filter for Go. |
| [xj2go](https://github.com/wk30/xj2go) | 20 | 2 | 2017/09/19 | 6 months ago | Convert xml or json to go struct. |
| [xj2go](https://github.com/stackerzzq/xj2go) | 19 | 2 | 2017/09/19 | 1 year ago | Convert xml or json to go struct. |
| [kace](https://github.com/codemodus/kace) | 16 | 1 | 2015/06/04 | 2 years ago | Common case conversions covering common initialisms. |
| [tagify](https://github.com/zoomio/tagify) | 13 | 1 | 2018/03/20 | 11 months ago | Produces a set of tags from given source. |
| [TySug](https://github.com/Dynom/TySug) | 9 | 1 | 2018/06/05 | 6 months ago | Alternative suggestions with respect to keyboard layouts. |
| [parseargs-go](https://github.com/txgruppi/parseargs-go) | 8 | 1 | 2016/02/24 | 4 years ago | string argument parser that understands quotes and backslashes. |
| [textwrap](https://github.com/isbm/textwrap) | 2 | 1 | 2019/07/26 | 1 year ago | Implementation of `textwrap` module from Python. |

## Third-party APIs
        
*Libraries for accessing third party APIs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-github](https://github.com/google/go-github) | 7172 | 222 | 2013/05/24 | 1 day ago | Go library for accessing the GitHub REST API v3. |
| [aws-sdk-go](https://github.com/aws/aws-sdk-go) | 6471 | 261 | 2014/12/05 | 1 day ago | The official AWS SDK for the Go programming language. |
| [google-api-go-client](https://github.com/googleapis/google-api-go-client) | 2553 | 160 | 2014/11/24 | 4 hours ago | Auto-generated Google APIs for Go. |
| [google-cloud-go](https://github.com/googleapis/google-cloud-go) | 2470 | 244 | 2014/05/09 | 1 day ago | Google Cloud APIs Go Client Library. |
| [discordgo](https://github.com/bwmarrin/discordgo) | 1813 | 52 | 2015/11/01 | 3 hours ago | Go bindings for the Discord Chat API. |
| [stripe-go](https://github.com/stripe/stripe-go) | 1308 | 39 | 2014/06/05 | 19 hours ago | Go client for the Stripe API. |
| [minio-go](https://github.com/minio/minio-go) | 1200 | 43 | 2015/05/02 | 3 days ago | Minio Go Library for Amazon S3 compatible cloud storage. |
| [go-twitter](https://github.com/dghubble/go-twitter) | 1156 | 31 | 2015/04/11 | 1 month ago | Go client library for the Twitter v1.1 APIs. |
| [anaconda](https://github.com/ChimeraCoder/anaconda) | 1077 | 21 | 2013/03/04 | 2 months ago | Go client library for the Twitter 1.1 API. |
| [facebook](https://github.com/huandu/facebook) | 944 | 111 | 2012/07/28 | 4 weeks ago | Go Library that supports the Facebook Graph API. |
| [githubv4](https://github.com/shurcooL/githubv4) | 749 | 20 | 2017/05/27 | 2 weeks ago | Go library for accessing the GitHub GraphQL API v4. |
| [webhooks](https://github.com/go-playground/webhooks) | 591 | 15 | 2015/10/25 | 3 weeks ago | Webhook receiver for GitHub and Bitbucket. |
| [paypal](https://github.com/plutov/paypal) | 405 | 21 | 2015/10/14 | 1 week ago | Wrapper for PayPal payment API. |
| [geo-golang](https://github.com/codingsince1985/geo-golang) | 389 | 13 | 2014/12/04 | 2 days ago | Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs. |
| [go-marathon](https://github.com/gambol99/go-marathon) | 196 | 13 | 2015/02/11 | 4 months ago | Go library for interacting with Mesosphere's Marathon PAAS. |
| [ethrpc](https://github.com/onrik/ethrpc) | 192 | 13 | 2017/01/24 | 5 months ago | Go bindings for Ethereum JSON RPC API. |
| [trello](https://github.com/adlio/trello) | 161 | 5 | 2016/09/24 | 1 month ago | Go wrapper for the Trello API. |
| [medium-sdk-go](https://github.com/Medium/medium-sdk-go) | 131 | 128 | 2015/09/26 | 2 years ago | Golang SDK for Medium's OAuth2 API. |
| [gostorm](https://github.com/jsgilmore/gostorm) | 127 | 11 | 2013/07/22 | 3 years ago | GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. |
| [hipchat](https://github.com/daneharrigan/hipchat) | 112 | 7 | 2013/04/28 | 3 years ago | A golang package to communicate with HipChat over XMPP. |
| [go-trending](https://github.com/andygrunwald/go-trending) | 111 | 7 | 2015/07/04 | 5 days ago | Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github. |
| [hipchat](https://github.com/andybons/hipchat) | 106 | 8 | 2012/10/20 | 4 years ago | This project implements a golang client library for the Hipchat API. |
| [wit-go](https://github.com/wit-ai/wit-go) | 100 | 16 | 2018/08/20 | 1 month ago | Go client for wit.ai HTTP API. |
| [pushover](https://github.com/gregdel/pushover) | 90 | 3 | 2015/02/19 | 3 months ago | Go wrapper for the Pushover API. |
| [cachet](https://github.com/andygrunwald/cachet) | 85 | 7 | 2015/10/31 | 2 years ago | Go client library for [Cachet (open source status page system)](https://cachethq.io/). |
| [twitter-scraper](https://github.com/n0madic/twitter-scraper) | 84 | 4 | 2018/11/29 | 1 week ago | Scrape the Twitter Frontend API without authentication and limits. |
| [igdb](https://github.com/Henry-Sarabia/igdb) | 65 | 2 | 2017/08/24 | 2 months ago | Go client for the [Internet Game Database API](https://api.igdb.com/). |
| [clarifai-go](https://github.com/Clarifai/clarifai-go) | 57 | 37 | 2015/09/28 | 3 years ago | Go client library for interfacing with the Clarifai API. |
| [go-circleci](https://github.com/jszwedko/go-circleci) | 56 | 3 | 2015/08/14 | 1 year ago | Go client library for interacting with CircleCI's API. |
| [megos](https://github.com/andygrunwald/megos) | 56 | 5 | 2015/10/02 | 2 years ago | Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster. |
| [gosip](https://github.com/koltyakov/gosip) | 53 | 3 | 2019/01/26 | 3 weeks ago | Go client library SharePoint API. |
| [simples3](https://github.com/rhnvrm/simples3) | 50 | 2 | 2018/12/06 | 5 days ago | Simple no frills AWS S3 Library using REST with V4 Signing written in Go. |
| [gads](https://github.com/emiddleton/gads) | 49 | 7 | 2014/01/20 | 1 year ago | Google Adwords Unofficial API. |
| [gogtrends](https://github.com/groovili/gogtrends) | 48 | 3 | 2018/12/27 | 1 month ago | Google Trends Unofficial API. |
| [go-amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) | 46 | 1 | 2016/11/15 | 2 years ago | Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html). |
| [ynab.go](https://github.com/brunomvsouza/ynab.go) | 45 | 2 | 2018/07/13 | 1 year ago | Go wrapper for the YNAB API. |
| [go-xkcd](https://github.com/nishanths/go-xkcd) | 43 | 3 | 2016/02/26 | 8 months ago | Go client for the xkcd API. |
| [uptimerobot](https://github.com/bitfield/uptimerobot) | 43 | 3 | 2018/05/29 | 1 month ago | Go wrapper and command-line client for the Uptime Robot v2 API. |
| [gomusicbrainz](https://github.com/michiwend/gomusicbrainz) | 42 | 6 | 2014/09/10 | 4 days ago | Go MusicBrainz WS2 client library. |
| [fcm](https://github.com/maddevsio/fcm) | 40 | 4 | 2017/01/06 | 11 months ago | Go library for Firebase Cloud Messaging. |
| [mixpanel](https://github.com/dukex/mixpanel) | 37 | 3 | 2014/05/20 | 1 year ago | Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications. |
| [golang-tmdb](https://github.com/cyruzin/golang-tmdb) | 37 | 1 | 2019/01/11 | 1 month ago | Golang wrapper for The Movie Database API v3. |
| [golyrics](https://github.com/mamal72/golyrics) | 36 | 4 | 2016/11/18 | 2 years ago | Golyrics is a Go library to fetch music lyrics data from the Wikia website. |
| [go-unsplash](https://github.com/hbagdi/go-unsplash) | 34 | 1 | 2017/01/19 | 11 months ago | Go client library for the [Unsplash.com](https://unsplash.com) API. |
| [go-spotify](https://github.com/rapito/go-spotify) | 33 | 2 | 2014/10/30 | 2 months ago | Go Library to access Spotify WEB API. |
| [translate](https://github.com/nuveo/translate) | 31 | 29 | 2015/07/13 | 5 years ago | Go online translation package. |
| [gami](https://github.com/bit4bit/gami) | 29 | 4 | 2014/05/14 | 2 years ago | Go library for Asterisk Manager Interface. |
| [gcm](https://github.com/TheOrioli/gcm) | 29 | 3 | 2015/11/09 | 5 years ago | Go library for Google Cloud Messaging. |
| [patreon-go](https://github.com/mxpv/patreon-go) | 23 | 4 | 2017/08/06 | 1 year ago | Go library for Patreon API. |
| [go-shopify](https://github.com/rapito/go-shopify) | 22 | 1 | 2014/10/28 | 2 months ago | Go Library to make CRUD request to the Shopify API. |
| [go-steam](https://github.com/sostronk/go-steam) | 22 | 10 | 2014/11/23 | 1 year ago | Go Library to interact with Steam game servers. |
| [go-twitch](https://github.com/knspriggs/go-twitch) | 20 | 5 | 2016/06/28 | 3 years ago | Go client for interacting with the Twitch v3 API. |
| [go-myanimelist](https://github.com/nstratos/go-myanimelist) | 19 | 2 | 2015/05/03 | 1 month ago | Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api). |
| [lastpass-go](https://github.com/ansd/lastpass-go) | 19 | 1 | 2019/07/11 | 2 days ago | Go client library for the [LastPass](https://www.lastpass.com/) API. |
| [brewerydb](https://github.com/naegelejd/brewerydb) | 17 | 3 | 2015/04/15 | 5 years ago | Go library for accessing the BreweryDB API. |
| [textbelt](https://github.com/farmergreg/textbelt) | 17 | 2 | 2015/09/01 | 5 years ago | Go client for the textbelt.com txt messaging API. |
| [airtable](https://github.com/mehanizm/airtable) | 17 | 1 | 2020/04/12 | 1 week ago | Go client library for the [Airtable API](https://airtable.com/api). |
| [codeship-go](https://github.com/codeship/codeship-go) | 16 | 21 | 2017/09/08 | 3 months ago | Go client library for interacting with Codeship's API v2. |
| [go-postman-collection](https://github.com/rbretecher/go-postman-collection) | 15 | 2 | 2019/11/16 | 1 month ago | Go module to work with [Postman Collections](https://learning.getpostman.com/docs/postman/collections/creating-collections/) (compatible with Insomnia). |
| [go-hacknews](https://github.com/PaulRosset/go-hacknews) | 12 | 2 | 2017/08/10 | 3 years ago | Tiny Go client for HackerNews API. |
| [go-google-analytics](https://github.com/chonthu/go-google-analytics) | 12 | 2 | 2015/06/01 | 5 years ago | Simple wrapper for easy google analytics reporting. |
| [coinpaprika-api-go-client](https://github.com/coinpaprika/coinpaprika-api-go-client) | 11 | 8 | 2018/09/25 | 5 months ago | Go client library for interacting with Coinpaprika's API. |
| [smitego](https://github.com/sergiotapia/smitego) | 10 | 0 | 2013/12/11 | 6 years ago | Go package to wraps access to the Smite game API. |
| [google-play-scraper](https://github.com/n0madic/google-play-scraper) | 10 | 1 | 2019/09/20 | 3 months ago | Get data from Google Play Store. |
| [go-aws-news](https://github.com/circa10a/go-aws-news) | 10 | 3 | 2020/01/08 | 1 month ago | Go application and library to fetch what's new from AWS. |
| [device-check-go](https://github.com/rinchsan/device-check-go) | 9 | 1 | 2019/04/11 | 15 hours ago | Go client library for interacting with [iOS DeviceCheck API](https://developer.apple.com/documentation/devicecheck) v1. |
| [go-here](https://github.com/abdullahselek/go-here) | 8 | 0 | 2019/07/07 | 7 months ago | Go client library around the HERE location based APIs. |
| [go-sophos](https://github.com/esurdam/go-sophos) | 8 | 2 | 2018/09/05 | 4 months ago | Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies. |
| [rrdaclient](https://github.com/Omie/rrdaclient) | 8 | 2 | 2014/09/15 | 6 years ago | Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP. |
| [slack](https://github.com/nlopes/slack) | 8 | 0 | 2015/01/24 | 3 months ago | Slack API in Go. |
| [go-google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) | 7 | 0 | 2016/10/24 | 4 years ago | Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/). |
| [gopaapi5](https://github.com/utekaravinash/gopaapi5) | 7 | 3 | 2020/02/15 | 10 months ago | Go Client Library for [Amazon Product Advertising API 5.0](https://webservices.amazon.com/paapi5/documentation/). |
| [go-sptrans](https://github.com/sergioaugrod/go-sptrans) | 6 | 3 | 2017/09/11 | 5 months ago | Go client library for the SPTrans Olho Vivo API. |
| [gumblr](https://github.com/mattcunningham/gumblr) | 6 | 1 | 2015/07/09 | 4 years ago | Go wrapper for the Tumblr v2 API. |
| [go-zooz](https://github.com/gojuno/go-zooz) | 6 | 13 | 2017/07/04 | 3 months ago | Go client for the Zooz API. |
| [go-chronos](https://github.com/axelspringer/go-chronos) | 4 | 8 | 2017/10/23 | 3 years ago | Go library for interacting with the [Chronos](https://mesos.github. |
| [libgoffi](https://github.com/clevabit/libgoffi) | 3 | 11 | 2019/08/03 | 5 months ago | Library adapter toolbox for native [libffi](http://sourceware. |
| [kanka](https://github.com/Henry-Sarabia/kanka) | 2 | 2 | 2019/12/26 | 6 months ago | Go client for the [Kanka API](https://kanka.io/en-US/docs/1.0). |
| [rawg-sdk-go](https://github.com/dimuska139/rawg-sdk-go) | 2 | 1 | 2020/10/16 | 1 month ago | Go library for the [RAWG Video Games Database](https://rawg. |
| [playlyfe-go-sdk](https://github.com/playlyfe/playlyfe-go-sdk) | 1 | 4 | 2015/05/25 | 5 years ago | The Playlyfe Rest API Go SDK. |
| [tripadvisor-golang](https://github.com/mrbenosborne/tripadvisor-golang) | 1 | 1 | 2019/04/15 | 1 year ago | Go wrapper for the TripAdvisor API. |
| [vl-go](https://github.com/verifid/vl-go) | 1 | 2 | 2019/02/09 | 11 months ago | Go client library around the VerifID identity verification layer API. |

## Utilities
        
*General utilities and tools to make your life easier.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [fzf](https://github.com/junegunn/fzf) | 34829 | 386 | 2013/10/23 | 2 days ago | Command-line fuzzy finder written in Go. |
| [hub](https://github.com/github/hub) | 20724 | 471 | 2009/12/05 | 6 days ago | wrap git commands with additional functionality to interact with github from the terminal. |
| [delve](https://github.com/go-delve/delve) | 13355 | 387 | 2014/05/20 | 1 year ago | Go debugger. |
| [ctop](https://github.com/bcicen/ctop) | 11065 | 159 | 2016/12/27 | 1 week ago | [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics. |
| [sqlx](https://github.com/jmoiron/sqlx) | 9726 | 188 | 2013/01/28 | 2 weeks ago | provides a set of extensions on top of the excellent built-in database/sql package. |
| [wuzz](https://github.com/asciimoo/wuzz) | 9492 | 170 | 2017/01/30 | 3 weeks ago | Interactive cli tool for HTTP inspection. |
| [goreleaser](https://github.com/goreleaser/goreleaser) | 7416 | 100 | 2016/12/21 | 1 day ago | Deliver Go binaries as fast and easily as possible. |
| [peco](https://github.com/peco/peco) | 6204 | 138 | 2014/06/06 | 1 month ago | Simplistic interactive filtering tool. |
| [usql](https://github.com/xo/usql) | 6097 | 115 | 2017/03/02 | 1 day ago | usql is a universal command-line interface for SQL databases. |
| [godropbox](https://github.com/dropbox/godropbox) | 3926 | 247 | 2014/06/22 | 7 months ago | Common libraries for writing Go services/applications from Dropbox. |
| [hystrix-go](https://github.com/afex/hystrix-go) | 3053 | 90 | 2013/12/15 | 2 months ago | Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker. |
| [goreporter](https://github.com/360EntSecGroup-Skylar/goreporter) | 2814 | 101 | 2017/03/27 | 2 years ago | Golang tool that does static analysis, unit testing, code review and generate code quality report. |
| [minify](https://github.com/tdewolff/minify) | 2572 | 53 | 2014/05/21 | 4 days ago | Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats. |
| [panicparse](https://github.com/maruel/panicparse) | 2522 | 38 | 2015/02/02 | 1 month ago | Groups similar goroutines and colorizes stack dump. |
| [go-funk](https://github.com/thoas/go-funk) | 2356 | 36 | 2016/12/30 | 1 week ago | Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...). |
| [mc](https://github.com/minio/mc) | 1731 | 50 | 2015/01/16 | 7 hours ago | Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems. |
| [storm](https://github.com/asdine/storm) | 1681 | 43 | 2016/01/10 | 5 months ago | Simple and powerful toolkit for BoltDB. |
| [mole](https://github.com/davrodpin/mole) | 1434 | 29 | 2018/10/04 | 1 month ago | cli app to easily create ssh tunnels. |
| [mergo](https://github.com/imdario/mergo) | 1400 | 23 | 2013/03/11 | 2 weeks ago | Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements. |
| [spinner](https://github.com/briandowns/spinner) | 1362 | 17 | 2014/12/13 | 1 week ago | Go package to easily provide a terminal spinner with options. |
| [boilr](https://github.com/tmrts/boilr) | 1250 | 29 | 2015/12/19 | 3 months ago | Blazingly fast CLI tool for creating projects from boilerplate templates. |
| [filetype](https://github.com/h2non/filetype) | 1250 | 29 | 2015/09/24 | 1 week ago | Small package to infer the file type checking the magic numbers signature. |
| [jump](https://github.com/gsamokovarov/jump) | 1007 | 16 | 2015/08/16 | 4 days ago | Jump helps you navigate faster by learning your habits. |
| [circuitbreaker](https://github.com/rubyist/circuitbreaker) | 925 | 20 | 2014/07/17 | 1 year ago | Circuit Breakers in Go. |
| [gtm](https://github.com/git-time-metric/gtm) | 858 | 28 | 2016/06/19 | 4 months ago | Simple, seamless, lightweight time tracking for Git. |
| [immortal](https://github.com/immortal/immortal) | 697 | 18 | 2016/06/30 | 7 months ago | \*nix cross-platform (OS agnostic) supervisor. |
| [hostctl](https://github.com/guumaster/hostctl) | 637 | 10 | 2020/03/14 | 6 months ago | A CLI tool to manage /etc/hosts with easy commands. |
| [circuit](https://github.com/cep21/circuit) | 537 | 15 | 2017/12/23 | 2 months ago | An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern. |
| [htcat](https://github.com/htcat/htcat) | 528 | 18 | 2013/08/05 | 1 year ago | Parallel and Pipelined HTTP GET Utility. |
| [godaemon](https://github.com/VividCortex/godaemon) | 465 | 30 | 2013/08/01 | 3 weeks ago | Utility to write daemons. |
| [go-dry](https://github.com/ungerik/go-dry) | 459 | 15 | 2014/02/28 | 5 days ago | DRY (don't repeat yourself) package for Go. |
| [ergo](https://github.com/cristianoliveira/ergo) | 448 | 6 | 2017/08/19 | 2 weeks ago | The management of multiple local services running over different ports made easy. |
| [koazee](https://github.com/wesovilabs/koazee) | 443 | 12 | 2018/11/09 | 2 months ago | Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays. |
| [gopencils](https://github.com/bndr/gopencils) | 435 | 14 | 2014/06/23 | 2 years ago | Small and simple package to easily consume REST APIs. |
| [mimetype](https://github.com/gabriel-vasile/mimetype) | 435 | 8 | 2018/07/02 | 1 week ago | Package for MIME type detection based on magic numbers. |
| [request](https://github.com/mozillazg/request) | 393 | 15 | 2014/12/21 | 1 year ago | Go HTTP Requests for Humans™. |
| [deepcopier](https://github.com/ulule/deepcopier) | 333 | 20 | 2015/07/24 | 9 months ago | Simple struct copying for Go. |
| [gubrak](https://github.com/novalagung/gubrak) | 326 | 6 | 2018/03/09 | 8 months ago | Golang utility library with syntactic sugar. It's like lodash, but for golang. |
| [go-rate](https://github.com/beefsack/go-rate) | 322 | 10 | 2014/08/25 | 5 months ago | Timed rate limiter for Go. |
| [clockwork](https://github.com/jonboulle/clockwork) | 318 | 6 | 2014/09/09 | 4 months ago | A simple fake clock for golang. |
| [retry](https://github.com/kamilsk/retry) | 278 | 5 | 2016/11/02 | 2 weeks ago | The most advanced functional mechanism to perform actions repetitively until successful. |
| [delve](https://github.com/derekparker/delve) | 276 | 7 | 2020/02/18 | 1 month ago | Go debugger. |
| [cli](https://github.com/create-go-app/cli) | 267 | 7 | 2019/12/30 | 3 days ago | A powerful CLI for create a new production-ready project with backend (Golang), frontend (JavaScript, TypeScript) & deploy automation (Ansible, Docker) by running one command. |
| [gohper](https://github.com/cosiner/gohper) | 253 | 20 | 2015/03/23 | 3 years ago | Various tools/modules help for development. |
| [gohper](https://github.com/zhuah/gohper) | 252 | 20 | 2015/03/23 | 3 years ago | Various tools/modules help for development. |
| [serve](https://github.com/syntaqx/serve) | 232 | 7 | 2019/01/10 | 7 months ago | A static http server anywhere you need. |
| [go-trigger](https://github.com/sadlil/go-trigger) | 211 | 12 | 2015/10/19 | 3 years ago | Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project. |
| [scany](https://github.com/georgysavva/scany) | 207 | 9 | 2020/07/02 | 3 months ago | Library for scanning data from a database into Go structs and more. |
| [util](https://github.com/shomali11/util) | 196 | 9 | 2017/05/24 | 10 months ago | Collection of useful utility functions. (strings, concurrency, manipulations, ...). |
| [gotenv](https://github.com/subosito/gotenv) | 189 | 3 | 2013/08/27 | 2 months ago | Load environment variables from `.env` or any `io.Reader` in Go. |
| [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) | 169 | 4 | 2016/11/08 | 1 year ago | go:generate tool for wrapping symbols exported by golang plugins (1.8 only). |
| [death](https://github.com/vrecan/death) | 165 | 4 | 2015/03/09 | 7 months ago | Managing go application shutdown with signals. |
| [rerun](https://github.com/ivpusic/rerun) | 161 | 7 | 2014/12/10 | 2 years ago | Recompiling and rerunning go apps when source changes. |
| [moldova](https://github.com/StabbyCutyou/moldova) | 159 | 5 | 2016/01/30 | 3 years ago | Utility for generating random data based on an input template. |
| [toolbox](https://github.com/viant/toolbox) | 155 | 15 | 2016/06/13 | 4 months ago | Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer. |
| [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) | 151 | 6 | 2015/10/12 | 1 year ago | XML Sitemap generator written in Go. |
| [apm](https://github.com/topfreegames/apm) | 149 | 18 | 2015/11/18 | 4 years ago | Process manager for Golang applications with an HTTP API. |
| [robustly](https://github.com/VividCortex/robustly) | 148 | 17 | 2013/07/08 | 3 weeks ago | Runs functions resiliently, catching and restarting panics. |
| [chyle](https://github.com/antham/chyle) | 136 | 6 | 2016/11/17 | 5 days ago | Changelog generator using a git repository with multiple configuration possibilities. |
| [onecache](https://github.com/adelowo/onecache) | 115 | 6 | 2017/04/14 | 8 months ago | Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc). |
| [lrserver](https://github.com/jaschaephraim/lrserver) | 113 | 5 | 2014/07/15 | 3 years ago | LiveReload server for Go. |
| [go-bsdiff](https://github.com/gabstv/go-bsdiff) | 110 | 1 | 2019/02/23 | 1 year ago | Pure Go bsdiff and bspatch libraries and CLI tools. |
| [xferspdy](https://github.com/monmohan/xferspdy) | 85 | 4 | 2015/05/22 | 2 months ago | Xferspdy provides binary diff and patch library in golang. |
| [mssqlx](https://github.com/linxGnu/mssqlx) | 83 | 8 | 2016/12/26 | 3 months ago | Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind. |
| [countries](https://github.com/biter777/countries) | 82 | 4 | 2019/04/22 | 3 months ago | Full implementation of ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and IANA ccTLD standarts. |
| [nostromo](https://github.com/pokanop/nostromo) | 81 | 3 | 2019/07/13 | 2 months ago | CLI for building powerful aliases. |
| [pm](https://github.com/VividCortex/pm) | 79 | 19 | 2013/11/17 | 2 months ago | Process (i.e. goroutine) manager with an HTTP API. |
| [goseaweedfs](https://github.com/linxGnu/goseaweedfs) | 78 | 8 | 2017/07/20 | 1 month ago | SeaweedFS client library with almost full features. |
| [go-health](https://github.com/Talento90/go-health) | 76 | 3 | 2018/02/13 | 2 years ago | Health package simplifies the way you add health check to your services. |
| [sorty](https://github.com/jfcg/sorty) | 73 | 2 | 2019/02/18 | 2 weeks ago | Fast Concurrent / Parallel Sorting. |
| [repeat](https://github.com/ssgreg/repeat) | 72 | 5 | 2017/11/22 | 6 months ago | Go implementation of different backoff strategies useful for retrying operations and heartbeating. |
| [scan](https://github.com/blockloop/scan) | 69 | 1 | 2017/11/27 | 4 weeks ago | Scan golang `sql.Rows` directly to structs, slices, or primitive types. |
| [unis](https://github.com/esemplastic/unis) | 68 | 4 | 2017/05/06 | 3 years ago | Common Architecture™ for String Utilities in Go. |
| [netbug](https://github.com/e-dard/netbug) | 68 | 1 | 2015/03/05 | 5 years ago | Easy remote profiling of your services. |
| [go-pattern-match](https://github.com/alexpantyukhin/go-pattern-match) | 67 | 2 | 2018/12/11 | 7 months ago | Pattern matching libray. |
| [multitick](https://github.com/VividCortex/multitick) | 64 | 17 | 2013/12/10 | 2 months ago | Multiplexor for aligned tickers. |
| [handy](https://github.com/miguelpragier/handy) | 63 | 8 | 2018/06/13 | 4 months ago | Many utilities and helpers like string handlers/formatters and validators. |
| [mimemagic](https://github.com/zRedShift/mimemagic) | 58 | 1 | 2018/10/11 | 1 month ago | Pure Go ultra performant MIME sniffing library/utility. |
| [minquery](https://github.com/icza/minquery) | 57 | 2 | 2016/11/16 | 10 months ago | MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off). |
| [go-astitodo](https://github.com/asticode/go-astitodo) | 56 | 2 | 2016/10/17 | 6 months ago | Parse TODOs in your GO code. |
| [goreadability](https://github.com/philipjkim/goreadability) | 56 | 6 | 2016/04/20 | 1 year ago | Webpage summary extractor using Facebook Open Graph and arc90's readability. |
| [pgo](https://github.com/arthurkushman/pgo) | 52 | 6 | 2018/12/26 | 3 months ago | Convenient functions for PHP community. |
| [mongo-go-pagination](https://github.com/gobeam/mongo-go-pagination) | 52 | 1 | 2020/02/04 | 4 months ago | Mongodb Pagination for official mongodb/mongo-go-driver package which supports  both normal queries and Aggregation pipelines. |
| [cmd](https://github.com/commander-cli/cmd) | 51 | 6 | 2019/09/27 | 7 months ago | Library for executing shell commands on osx, windows and linux. |
| [golog](https://github.com/mlimaloureiro/golog) | 50 | 3 | 2016/01/09 | 2 years ago | Easy and lightweight CLI tool to time track your tasks. |
| [copy-pasta](https://github.com/jutkko/copy-pasta) | 47 | 5 | 2017/01/28 | 7 months ago | Universal multi-workstation clipboard that uses S3 like backend for the storage. |
| [retry](https://github.com/thedevsaddam/retry) | 46 | 1 | 2018/02/25 | 3 weeks ago | Simple and easy retry mechanism package for Go. |
| [filter](https://github.com/gookit/filter) | 43 | 5 | 2018/09/26 | 2 weeks ago | provide filtering, sanitizing, and conversion of Go data. |
| [goback](https://github.com/carlescere/goback) | 43 | 1 | 2015/03/13 | 3 years ago | Go simple exponential backoff package. |
| [golarm](https://github.com/msempere/golarm) | 43 | 2 | 2015/08/14 | 5 years ago | Fire alarms with system events. |
| [intrinsic](https://github.com/mengzhuo/intrinsic) | 42 | 4 | 2017/06/13 | 3 years ago | Use x86 SIMD without writing any assembly code. |
| [dbt](https://github.com/nikogura/dbt) | 40 | 1 | 2017/11/30 | 1 month ago | A framework for running self-updating signed binaries from a central, trusted repository. |
| [gpath](https://github.com/tenntenn/gpath) | 40 | 3 | 2017/05/24 | 3 years ago | Library to simplify access struct fields with Go's expression in reflection. |
| [retry-go](https://github.com/rafaeljesus/retry-go) | 40 | 2 | 2017/06/09 | 2 years ago | Retrying made simple and easy for golang. |
| [beyond](https://github.com/wesovilabs/beyond) | 39 | 1 | 2019/10/18 | 1 year ago | The Go tool that will drive you to the AOP world! |
| [slice](https://github.com/psampaz/slice) | 38 | 0 | 2019/11/26 | 10 months ago | Type-safe functions for common Go slice operations. |
| [go-httpheader](https://github.com/mozillazg/go-httpheader) | 34 | 3 | 2017/06/24 | 2 weeks ago | Go library for encoding structs into Header fields. |
| [myhttp](https://github.com/inancgumus/myhttp) | 33 | 1 | 2017/09/13 | 2 years ago | Simple API to make HTTP GET requests with timeout support. |
| [go-lock](https://github.com/viney-shih/go-lock) | 33 | 0 | 2020/04/30 | 4 months ago | go-lock is a lock library implementing read-write mutex and read-write trylock without starvation. |
| [rclient](https://github.com/zpatrick/rclient) | 32 | 3 | 2017/02/28 | 1 year ago | Readable, flexible, simple-to-use client for REST APIs. |
| [equalizer](https://github.com/reugn/equalizer) | 31 | 1 | 2019/06/14 | 1 month ago | Quota manager and rate limiter collection for Go. |
| [evaluator](https://github.com/nullne/evaluator) | 30 | 2 | 2017/04/27 | 7 months ago | Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend. |
| [gostrutils](https://github.com/ik5/gostrutils) | 30 | 2 | 2018/09/19 | 6 months ago | Collections of string manipulation and conversion functions. |
| [changie](https://github.com/miniscruff/changie) | 28 | 2 | 2020/12/05 | 9 hours ago | Automated changelog tool for preparing releases with lots of customization options. |
| [tome](https://github.com/cyruzin/tome) | 27 | 1 | 2019/04/12 | 10 months ago | Tome was designed to paginate simple RESTful APIs. |
| [cmd](https://github.com/SimonBaeumer/cmd) | 25 | 2 | 2019/09/27 | 10 months ago | Library for executing shell commands on osx, windows and linux. |
| [ugo](https://github.com/alxrm/ugo) | 24 | 1 | 2016/02/17 | 4 years ago | ugo is slice toolbox with concise syntax for Go. |
| [generate](https://github.com/go-playground/generate) | 23 | 3 | 2015/11/15 | 4 years ago | runs go generate recursively on a specified path or environment variable and can filter by regex. |
| [limiters](https://github.com/mennanov/limiters) | 23 | 1 | 2019/08/28 | 5 months ago | Rate limiters for distributed applications in Golang with configurable back-ends and distributed locks. |
| [slicer](https://github.com/leaanthony/slicer) | 23 | 1 | 2019/01/10 | 3 months ago | Makes working with slices easier. |
| [goplaceholder](https://github.com/michiwend/goplaceholder) | 22 | 2 | 2014/10/12 | 5 years ago | a small golang lib to generate placeholder images. |
| [rerate](https://github.com/abo/rerate) | 18 | 4 | 2016/05/24 | 3 years ago | Redis-based rate counter and rate limiter for Go. |
| [r](https://github.com/is5/r) | 18 | 3 | 2020/02/20 | 11 months ago | Python-like `range()` experience for Go. |
| [ghokin](https://github.com/antham/ghokin) | 17 | 2 | 2018/08/03 | 3 days ago | Parallelized formatter with no external dependencies for gherkin (cucumber, behat...). |
| [ctxutil](https://github.com/posener/ctxutil) | 16 | 1 | 2018/07/30 | 11 months ago | A collection of utility functions for contexts. |
| [dlog](https://github.com/kirillDanshin/dlog) | 16 | 2 | 2016/07/04 | 3 years ago | Compile-time controlled logger to make your release smaller without removing debug calls. |
| [filler](https://github.com/yaronsumel/filler) | 16 | 1 | 2017/04/05 | 3 years ago | small utility to fill structs using "fill" tag. |
| [shutdown](https://github.com/ztrue/shutdown) | 16 | 1 | 2018/11/17 | 2 years ago | App shutdown hooks for `os.Signal` handling. |
| [structs](https://github.com/PumpkinSeed/structs) | 16 | 2 | 2017/08/26 | 3 years ago | Implement simple functions to manipulate structs. |
| [mimesniffer](https://github.com/aofei/mimesniffer) | 15 | 1 | 2018/12/20 | 1 week ago | A MIME type sniffer for Go. |
| [okrun](https://github.com/xta/okrun) | 15 | 2 | 2014/10/01 | 6 years ago | go run error steamroller. |
| [rest-go](https://github.com/edermanoel94/rest-go) | 14 | 3 | 2019/07/29 | 6 months ago | A package that provide many helpful methods for working with rest api. |
| [command](https://github.com/txgruppi/command) | 13 | 1 | 2015/08/24 | 4 years ago | Command pattern for Go with thread safe serial and parallel dispatcher. |
| [jsend](https://github.com/clevergo/jsend) | 13 | 3 | 2020/01/14 | 2 months ago | JSend's implementation writen in Go. |
| [backscanner](https://github.com/icza/backscanner) | 12 | 2 | 2017/10/18 | 1 year ago | A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward. |
| [retry](https://github.com/shafreeck/retry) | 11 | 0 | 2018/07/18 | 1 year ago | A pretty simple library to ensure your work to be done. |
| [go-convert](https://github.com/Eun/go-convert) | 11 | 0 | 2019/06/07 | 6 months ago | Package go-convert enbles you to convert a value into another type. |
| [go-countries](https://github.com/mikekonan/go-countries) | 9 | 3 | 2020/10/27 | 1 month ago | Lightweight lookup over ISO-3166 codes. |
| [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) | 7 | 1 | 2019/05/16 | 1 year ago | Go package for working with Problem Details. |
| [retry](https://github.com/percolate/retry) | 7 | 33 | 2018/06/15 | 1 year ago | A simple but highly configurable retry package for Go. |
| [statiks](https://github.com/janiltonmaciel/statiks) | 7 | 0 | 2018/06/26 | 4 months ago | Fast, zero-configuration, static HTTP filer server. |
| [blank](https://github.com/Henry-Sarabia/blank) | 6 | 2 | 2019/02/13 | 1 year ago | Verify or remove blanks and whitespace from strings. |
| [ptr](https://github.com/gotidy/ptr) | 6 | 1 | 2019/12/25 | 3 months ago | Package that provide functions for simplified creation of pointers from constants of basic types. |
| [sliceconv](https://github.com/Henry-Sarabia/sliceconv) | 6 | 1 | 2019/02/15 | 1 year ago | Slice conversion between primitive types. |
| [silk](https://github.com/chrispassas/silk) | 4 | 1 | 2018/12/18 | 3 months ago | Read silk netflow files. |
| [nfdump](https://github.com/chrispassas/nfdump) | 4 | 1 | 2020/04/08 | 9 months ago | Read nfdump netflow files. |
| [go-safe](https://github.com/kenkyu392/go-safe) | 3 | 0 | 2019/10/29 | 10 months ago | Panic-safe sandbox. |
| [copy](https://github.com/gotidy/copy) | 3 | 2 | 2020/10/09 | 1 month ago | Package for fast copying structs of different types. |
| [olaf](https://github.com/btnguyen2k/olaf) | 2 | 1 | 2019/01/03 | 1 year ago | Twitter Snowflake implemented in Go. |
| [lets-go](https://github.com/aplescia-chwy/lets-go) | 2 | 1 | 2020/02/19 | 4 months ago | Go module that provides common utilities for Cloud Native REST API development. Also contains AWS Specific utilities. |
| [tik](https://github.com/andy2046/tik) | 2 | 1 | 2020/07/04 | 4 months ago | Simple and easy timing wheel package for Go. |
| [goctx](https://github.com/zerosnake0/goctx) | 2 | 1 | 2020/11/14 | 2 months ago | Get your context value with high performance. |
| [compare](https://github.com/posener/compare) | 1 | 1 | 2020/03/13 | 10 months ago | Enables more readable and easier comparison tasks. |
| [bleep](https://github.com/sinhashubham95/bleep) | 1 | 1 | 2021/01/02 | 1 month ago | Perform any number of actions on any set of OS signals in Go. |

## UUID
        
*Libraries for working with UUIDs.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [uuid](https://github.com/google/uuid) | 2440 | 49 | 2016/02/12 | 4 days ago | Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services. |
| [ulid](https://github.com/oklog/ulid) | 2144 | 42 | 2016/12/06 | 2 weeks ago | Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier). |
| [uuid](https://github.com/gofrs/uuid) | 883 | 18 | 2018/07/13 | 1 month ago | Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid. |
| [wuid](https://github.com/edwingeng/wuid) | 389 | 16 | 2018/01/27 | 1 month ago | An extremely fast unique number generator, 10-135 times faster than UUID. |
| [sno](https://github.com/muyo/sno) | 43 | 3 | 2019/05/26 | 10 months ago | Compact, sortable and fast unique IDs with embedded metadata. |
| [Goid](https://github.com/JakeHL/Goid) | 29 | 4 | 2017/05/19 | 2 years ago | Generate and Parse RFC4122 compliant V4 UUIDs. |
| [nanoid](https://github.com/aidarkhanov/nanoid) | 29 | 1 | 2019/07/02 | 7 months ago | A tiny and efficient Go unique string ID generator. |
| [uuid](https://github.com/agext/uuid) | 12 | 2 | 2016/02/03 | 11 months ago | Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier. |
| [gouid](https://github.com/twharmon/gouid) | 5 | 1 | 2020/10/08 | 1 month ago | Generate cryptographically secure random string IDs with just one allocation. |

## Validation
        
*Libraries for validation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [validator](https://github.com/go-playground/validator) | 7056 | 97 | 2015/02/12 | 1 week ago | Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving. |
| [govalidator](https://github.com/asaskevich/govalidator) | 4624 | 103 | 2014/06/20 | 4 days ago | Validators and sanitizers for strings, numerics, slices and structs. |
| [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) | 1898 | 27 | 2016/06/22 | 2 weeks ago | Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags. |
| [govalidator](https://github.com/thedevsaddam/govalidator) | 960 | 21 | 2017/09/13 | 3 weeks ago | Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. |
| [validate](https://github.com/gookit/validate) | 365 | 14 | 2018/07/16 | 1 week ago | Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features. |
| [checkdigit](https://github.com/osamingo/checkdigit) | 74 | 0 | 2019/04/05 | 1 month ago | Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.). |
| [validate](https://github.com/gobuffalo/validate) | 57 | 7 | 2018/02/10 | 2 months ago | This package provides a framework for writing validations for Go applications. |
| [jio](https://github.com/faceair/jio) | 55 | 2 | 2018/10/28 | 9 months ago | jio is a json schema validator similar to [joi](https://github.com/hapijs/joi). |
| [gody](https://github.com/guiferpa/gody) | 48 | 0 | 2018/11/01 | 1 week ago | :balloon: A lightweight struct validator for Go. |
| [terraform-validator](https://github.com/thazelart/terraform-validator) | 43 | 2 | 2019/05/29 | 4 months ago | A norms and conventions validator for Terraform. |
| [govalid](https://github.com/twharmon/govalid) | 22 | 1 | 2019/02/17 | 3 months ago | Fast, tag-based validation for structs. |

## Version Control
        
*Libraries for version control.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-git](https://github.com/src-d/go-git) | 5013 | 100 | 2015/10/22 | 10 months ago | highly extensible Git implementation in pure Go. |
| [go-git](https://github.com/go-git/go-git) | 1992 | 33 | 2019/12/19 | 4 days ago | highly extensible Git implementation in pure Go. |
| [git2go](https://github.com/libgit2/git2go) | 1588 | 52 | 2013/03/05 | 1 week ago | Go bindings for libgit2. |
| [hercules](https://github.com/src-d/hercules) | 1255 | 19 | 2016/12/12 | 4 weeks ago | gaining advanced insights from Git repository history. |
| [gh](https://github.com/rjeczalik/gh) | 75 | 6 | 2015/03/08 | 2 years ago | Scriptable server and net/http middleware for GitHub Webhooks. |
| [go-vcs](https://github.com/sourcegraph/go-vcs) | 75 | 47 | 2013/06/02 | 3 months ago | manipulate and inspect VCS repositories in Go. |
| [hgo](https://github.com/beyang/hgo) | 13 | 3 | 2014/06/18 | 5 years ago | Hgo is a collection of Go packages providing read-access to local Mercurial repositories. |

## Video
        
*Libraries for manipulating video.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [goav](https://github.com/giorgisio/goav) | 1511 | 49 | 2015/05/21 | 1 month ago | Comprehensive Go bindings for FFmpeg. |
| [m3u8](https://github.com/grafov/m3u8) | 789 | 36 | 2013/02/05 | 1 day ago | Parser and generator library of M3U8 playlists for Apple HLS. |
| [gmf](https://github.com/3d0c/gmf) | 655 | 32 | 2013/04/03 | 8 months ago | Go bindings for FFmpeg av\* libraries. |
| [go-astits](https://github.com/asticode/go-astits) | 338 | 15 | 2017/07/04 | 1 week ago | Parse and demux MPEG Transport Streams (.ts) natively in GO. |
| [go-astisub](https://github.com/asticode/go-astisub) | 300 | 8 | 2016/12/16 | 1 week ago | Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.). |
| [libvlc-go](https://github.com/adrg/libvlc-go) | 182 | 10 | 2015/01/06 | 6 days ago | Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player). |
| [gst](https://github.com/ziutek/gst) | 159 | 10 | 2011/07/26 | 1 month ago | Go bindings for GStreamer. |
| [go-m3u8](https://github.com/quangngotan95/go-m3u8) | 69 | 2 | 2018/11/06 | 9 months ago | Parser and generator library for Apple m3u8 playlists. |
| [v4l](https://github.com/korandiz/v4l) | 49 | 5 | 2016/10/25 | 2 months ago | Video capture library for Linux, written in Go. |
| [libgosubs](https://github.com/wargarblgarbl/libgosubs) | 13 | 2 | 2017/05/03 | 9 months ago | Subtitle format support for go. Supports .srt, .ttml, and .ass. |
| [go-mpd](https://github.com/unki2aut/go-mpd) | 5 | 1 | 2018/11/02 | 6 months ago | Parser and generator library for MPEG-DASH manifest files. |

## Web Frameworks
        
*Full stack web frameworks.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gin](https://github.com/gin-gonic/gin) | 45659 | 1341 | 2014/06/16 | 4 hours ago | Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity. |
| [echo](https://github.com/labstack/echo) | 19172 | 543 | 2015/03/01 | 2 days ago | High performance, minimalist Go web framework. |
| [revel](https://github.com/revel/revel) | 12137 | 546 | 2011/12/09 | 1 week ago | High-productivity web framework for the Go language. |
| [fiber](https://github.com/gofiber/fiber) | 11661 | 213 | 2020/01/16 | 1 day ago | An Express.js inspired web framework build on Fasthttp. |
| [goa](https://github.com/goadesign/goa) | 4132 | 169 | 2014/12/05 | 5 days ago | Goa provides a holistic approach for developing remote APIs and microservices in Go. |
| [go-json-rest](https://github.com/ant0ine/go-json-rest) | 3457 | 160 | 2013/02/19 | 3 weeks ago | Quick and easy way to setup a RESTful JSON API. |
| [gizmo](https://github.com/nytimes/gizmo) | 3345 | 124 | 2015/12/15 | 2 days ago | Microservice toolkit used by the New York Times. |
| [macaron](https://github.com/go-macaron/macaron) | 3120 | 149 | 2014/07/10 | 3 months ago | Macaron is a high productive and modular design web framework in Go. |
| [utron](https://github.com/gernest/utron) | 2181 | 69 | 2015/09/16 | 2 years ago | Lightweight MVC framework for Go(Golang). |
| [go-tigertonic](https://github.com/rcrowley/go-tigertonic) | 1001 | 46 | 2013/02/09 | 2 years ago | Go framework for building JSON web services inspired by Dropwizard. |
| [tango](https://github.com/lunny/tango) | 837 | 77 | 2014/12/17 | 1 year ago | Micro & pluggable web framework for Go. |
| [goyave](https://github.com/System-Glitch/goyave) | 797 | 25 | 2019/10/21 | 2 days ago | Feature-complete web framework aimed at clean code and fast development, with powerful built-in functionalities. |
| [gongular](https://github.com/mustafaakin/gongular) | 438 | 21 | 2016/06/22 | 7 months ago | Fast Go web framework with input mapping/validation and (DI) Dependency Injection. |
| [gearbox](https://github.com/gogearbox/gearbox) | 428 | 11 | 2020/04/25 | 3 days ago | A web framework written in Go with a focus on high performance and memory optimization. |
| [neo](https://github.com/ivpusic/neo) | 409 | 33 | 2015/02/04 | 3 years ago | Neo is minimal and fast Go Web Framework with extremely simple API. |
| [air](https://github.com/aofei/air) | 400 | 19 | 2016/07/20 | 3 weeks ago | An ideally refined web framework for Go. |
| [mango](https://github.com/paulbellamy/mango) | 353 | 22 | 2011/05/25 | 3 years ago | Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. |
| [aero](https://github.com/aerogo/aero) | 347 | 18 | 2016/11/09 | 9 months ago | High-performance web framework for Go, reaches top scores in Lighthouse. |
| [gondola](https://github.com/rainycape/gondola) | 311 | 15 | 2014/07/25 | 2 years ago | The web framework for writing faster sites, faster. |
| [golf](https://github.com/dinever/golf) | 246 | 17 | 2015/11/18 | 4 years ago | Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library. |
| [flamingo](https://github.com/i-love-flamingo/flamingo) | 193 | 23 | 2019/04/02 | 3 days ago | Framework for pluggable web projects. Including a concept for modules and offering features for DI, Configareas, i18n, template engines, graphql, observability, security, events, routing & reverse routing etc. |
| [gearbox](https://github.com/abahmed/gearbox) | 176 | 4 | 2020/04/25 | 8 months ago | A web framework written in Go with a focus on high performance and memory optimization. |
| [webgo](https://github.com/bnkamalesh/webgo) | 157 | 8 | 2015/12/16 | 3 hours ago | A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc). |
| [ginrpc](https://github.com/xxjwxc/ginrpc) | 155 | 6 | 2019/06/22 | 5 months ago | Gin parameter automatic binding tool,gin rpc tools. |
| [hiboot](https://github.com/hidevopsio/hiboot) | 149 | 13 | 2018/03/16 | 3 weeks ago | hiboot is a high performance web application framework with auto configuration and dependency injection support. |
| [flamingo-commerce](https://github.com/i-love-flamingo/flamingo-commerce) | 147 | 22 | 2019/04/02 | 3 days ago | Providing e-commerce features using clean architecture like DDD and ports and adapters, that you can use to build flexible e-commerce applications. |
| [go-rest](https://github.com/ungerik/go-rest) | 125 | 10 | 2012/07/13 | 4 years ago | Small and evil REST framework for Go. |
| [uadmin](https://github.com/uadmin/uadmin) | 122 | 10 | 2018/10/05 | 5 months ago | Fully featured web framework for Golang, inspired by Django. |
| [appy](https://github.com/appist/appy) | 80 | 1 | 2019/05/27 | 3 days ago | An opinionated productive web framework that helps scaling business easier. |
| [vox](https://github.com/aisk/vox) | 74 | 2 | 2014/12/24 | 2 weeks ago | A golang web framework for humans, inspired by Koa heavily. |
| [golax](https://github.com/fulldump/golax) | 73 | 7 | 2016/01/30 | 2 years ago | A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more. |
| [microservice](https://github.com/claygod/microservice) | 73 | 9 | 2016/12/15 | 1 year ago | The framework for the creation of microservices, written in Golang. |
| [beego](https://github.com/astaxie/beego) | 71 | 4 | 2012/02/29 | 2 weeks ago | beego is an open-source, high-performance web framework for the Go programming language. |
| [patron](https://github.com/beatlabs/patron) | 69 | 16 | 2019/01/30 | 1 day ago | Patron is a microservice framework following best cloud practices with a focus on productivity. |
| [yarf](https://github.com/yarf-framework/yarf) | 61 | 3 | 2015/09/02 | 1 year ago | Fast micro-framework designed to build REST APIs and web services in a fast and simple way. |
| [rux](https://github.com/gookit/rux) | 59 | 3 | 2018/08/05 | 1 week ago | Simple and fast web framework for build golang HTTP applications. |
| [fireball](https://github.com/zpatrick/fireball) | 56 | 4 | 2016/07/20 | 2 years ago | More "natural" feeling web framework. |
| [goa](https://github.com/goa-go/goa) | 44 | 2 | 2019/07/26 | 1 year ago | goa is just like koajs for golang, it is a flexible, light, high-performance and extensible web framework based on middleware. |
| [api](https://github.com/resoursea/api) | 31 | 6 | 2015/01/24 | 6 years ago | REST framework for quickly writing resource based services. |
| [rex](https://github.com/goanywhere/rex) | 31 | 1 | 2014/10/16 | 3 years ago | Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`. |
| [goweb](https://github.com/twharmon/goweb) | 20 | 1 | 2019/05/07 | 3 days ago | Web framework with routing, websockets, logging, middleware, static file server (optional gzip), and automatic TLS. |
| [banjo](https://github.com/nsheremet/banjo) | 17 | 1 | 2017/12/09 | 3 years ago | Very simple and fast web framework for Go. |

### Middlewares
        

#### Actual middlewares
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [tollbooth](https://github.com/didip/tollbooth) | 1854 | 51 | 2015/05/17 | 1 month ago | Rate limit HTTP request handler. |
| [cors](https://github.com/rs/cors) | 1741 | 31 | 2014/10/25 | 2 months ago | Easily add CORS capabilities to your API. |
| [limiter](https://github.com/ulule/limiter) | 1217 | 27 | 2015/10/02 | 5 days ago | Dead simple rate limit middleware for Go. |
| [go-server-timing](https://github.com/mitchellh/go-server-timing) | 802 | 19 | 2018/02/12 | 3 months ago | Add/parse Server-Timing header. |
| [go-fault](https://github.com/github/go-fault) | 388 | 96 | 2020/05/14 | 2 weeks ago | Fault injection middleware for Go. |
| [ln-paywall](https://github.com/philippgille/ln-paywall) | 110 | 5 | 2018/06/29 | 2 years ago | Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin). |
| [xff](https://github.com/sebest/xff) | 77 | 2 | 2014/12/22 | 1 month ago | Handle `X-Forwarded-For` header and friends. |
| [formjson](https://github.com/rs/formjson) | 36 | 2 | 2015/03/19 | 5 years ago | Transparently handle JSON input as a standard form POST. |
| [client-timing](https://github.com/posener/client-timing) | 18 | 1 | 2018/02/23 | 11 months ago | An HTTP client for Server-Timing header. |

#### Libraries for creating HTTP middlewares
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [negroni](https://github.com/urfave/negroni) | 6896 | 237 | 2014/05/18 | 3 weeks ago | Idiomatic HTTP middleware for Golang. |
| [alice](https://github.com/justinas/alice) | 2205 | 49 | 2014/05/25 | 2 weeks ago | Painless middleware chaining for Go. |
| [render](https://github.com/unrolled/render) | 1407 | 36 | 2014/06/10 | 2 months ago | Go package for easily rendering JSON, XML, and HTML template responses. |
| [stats](https://github.com/thoas/stats) | 574 | 16 | 2015/03/05 | 1 year ago | Go middleware that stores various information about your web application. |
| [interpose](https://github.com/carbocation/interpose) | 291 | 12 | 2014/07/20 | 4 years ago | Minimalist net/http middleware for golang. |
| [renderer](https://github.com/thedevsaddam/renderer) | 219 | 7 | 2017/11/07 | 3 weeks ago | Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go. |
| [muxchain](https://github.com/stephens2424/muxchain) | 209 | 5 | 2014/05/03 | 1 year ago | Lightweight middleware for net/http. |
| [rye](https://github.com/InVisionApp/rye) | 95 | 196 | 2016/10/06 | 2 years ago | Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context. |
| [gores](https://github.com/alioygur/gores) | 94 | 5 | 2015/12/25 | 1 month ago | Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs. |
| [mediary](https://github.com/HereMobilityDevelopers/mediary) | 71 | 3 | 2020/03/23 | 7 months ago | add interceptors to `http.Client` to allow dumping/shaping/tracing/... of requests/responses. |
| [chain](https://github.com/codemodus/chain) | 63 | 6 | 2015/05/14 | 2 years ago | Handler wrapper chaining with scoped data (net/context-based "middleware"). |
| [wrap](https://github.com/go-on/wrap) | 59 | 2 | 2014/02/16 | 2 years ago | Small middlewares package for net/http. |
| [catena](https://github.com/codemodus/catena) | 7 | 1 | 2015/07/30 | 2 years ago | http.Handler wrapper catenation (same API as "chain"). |

### Routers
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [mux](https://github.com/gorilla/mux) | 13722 | 306 | 2012/10/02 | 2 weeks ago | Powerful URL router and dispatcher for golang. |
| [httprouter](https://github.com/julienschmidt/httprouter) | 12352 | 315 | 2013/12/05 | 3 months ago | High performance router. Use this and the standard http handlers to form a very high performance web framework. |
| [chi](https://github.com/go-chi/chi) | 8870 | 187 | 2015/10/15 | 3 days ago | Small, fast and expressive HTTP router built on net/context. |
| [web](https://github.com/gocraft/web) | 1439 | 59 | 2013/11/16 | 4 months ago | Mux and middleware package in Go. |
| [bone](https://github.com/go-zoo/bone) | 1266 | 36 | 2014/11/19 | 1 year ago | Lightning Fast HTTP Multiplexer. |
| [fasthttprouter](https://github.com/buaazp/fasthttprouter) | 871 | 33 | 2015/12/13 | 1 year ago | High performance router forked from `httprouter`. The first router fit for `fasthttp`. |
| [goji](https://github.com/goji/goji) | 844 | 42 | 2015/11/16 | 1 year ago | Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`. |
| [gorouter](https://github.com/xujiajun/gorouter) | 493 | 15 | 2018/01/29 | 1 year ago | A simple and fast HTTP router for Go. |
| [httptreemux](https://github.com/dimfeld/httptreemux) | 477 | 24 | 2014/05/14 | 4 months ago | High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter. |
| [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) | 396 | 28 | 2015/10/27 | 1 month ago | An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs. |
| [lars](https://github.com/go-playground/lars) | 382 | 15 | 2015/12/24 | 1 year ago | Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks. |
| [siesta](https://github.com/VividCortex/siesta) | 350 | 29 | 2014/09/23 | 2 months ago | Composable framework to write middleware and handlers. |
| [vestigo](https://github.com/husobee/vestigo) | 265 | 16 | 2015/09/22 | 4 months ago | Performant, stand-alone, HTTP compliant URL Router for go web applications. |
| [router](https://github.com/gowww/router) | 157 | 6 | 2017/05/25 | 9 months ago | Lightning fast HTTP router fully compatible with the net/http.Handler interface. |
| [alien](https://github.com/gernest/alien) | 115 | 4 | 2016/01/30 | 1 year ago | Lightweight and fast http router from outer space. |
| [pure](https://github.com/go-playground/pure) | 111 | 5 | 2016/09/23 | 2 months ago | Is a lightweight HTTP router that sticks to the std "net/http" implementation. |
| [Bxog](https://github.com/claygod/Bxog) | 101 | 8 | 2016/05/19 | 8 months ago | Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters. |
| [violetear](https://github.com/nbari/violetear) | 100 | 4 | 2015/06/19 | 1 year ago | Go HTTP router. |
| [gorouter](https://github.com/vardius/gorouter) | 91 | 5 | 2016/07/14 | 2 months ago | GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`. |
| [xmux](https://github.com/rs/xmux) | 88 | 6 | 2015/12/14 | 3 years ago | High performance muxer based on `httprouter` with `net/context` support. |
| [bellt](https://github.com/GuilhermeCaruso/bellt) | 48 | 6 | 2019/02/21 | 8 months ago | A simple Go HTTP router. |
| [fastrouter](https://github.com/razonyang/fastrouter) | 19 | 2 | 2017/11/01 | 3 years ago | a fast, flexible HTTP router written in Go. |
| [route](https://github.com/goroute/route) | 7 | 3 | 2019/07/06 | 1 year ago | Simple yet powerful HTTP request multiplexer. |

## Windows
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-ole](https://github.com/go-ole/go-ole) | 711 | 40 | 2011/01/21 | 1 month ago | Win32 OLE implementation for golang. |
| [d3d9](https://github.com/gonutz/d3d9) | 115 | 7 | 2015/12/12 | 4 months ago | Go bindings for Direct3D9. |
| [gosddl](https://github.com/MonaxGT/gosddl) | 5 | 1 | 2018/12/04 | 1 year ago | Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL. |

## XML
        
*Libraries and tools for manipulating XML.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [zek](https://github.com/miku/zek) | 446 | 20 | 2017/11/23 | 3 months ago | Generate a Go struct from XML. |
| [xpath](https://github.com/antchfx/xpath) | 369 | 9 | 2016/10/09 | 1 month ago | XPath package for Go. |
| [xquery](https://github.com/antchfx/xquery) | 155 | 11 | 2016/10/09 | 2 years ago | XQuery lets you extract data from HTML/XML documents using XPath expression. |
| [xml2map](https://github.com/sbabiv/xml2map) | 29 | 1 | 2018/08/06 | 2 weeks ago | XML to MAP converter written Golang. |
| [xmlwriter](https://github.com/shabbyrobe/xmlwriter) | 18 | 2 | 2017/04/11 | 1 year ago | Procedural XML generation API based on libxml2's xmlwriter module. |
| [XML-Comp](https://github.com/XML-Comp/XML-Comp) | 15 | 1 | 2016/10/25 | 2 years ago | Simple command line XML comparer that generates diffs of folders, files and tags. |

## WebAssembly
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [tinygo](https://github.com/tinygo-org/tinygo) | 7507 | 148 | 2018/06/07 | 1 hour ago | Go compiler for small places. Microcontrollers, WebAssembly, and command-line tools. Based on LLVM. |
| [dom](https://github.com/dennwc/dom) | 420 | 16 | 2018/06/30 | 1 year ago | DOM library. |
| [go-canvas](https://github.com/markfarnan/go-canvas) | 113 | 4 | 2019/05/05 | 2 months ago | Library to use HTML5 Canvas, with all drawing within go code. |
| [webapi](https://github.com/gowebapi/webapi) | 76 | 5 | 2019/02/08 | 2 weeks ago | Bindings for DOM and HTML generated from WebIDL. |
| [wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest) | 71 | 3 | 2018/07/14 | 3 months ago | Run Go WASM tests in your browser. |
| [vert](https://github.com/norunners/vert) | 45 | 4 | 2018/03/25 | 6 months ago | Interop between Go and JS values. |

## Files
        

## File Handling
        
*Libraries for handling files and file systems.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [afero](https://github.com/spf13/afero) | 3524 | 95 | 2014/10/28 | 2 days ago | FileSystem Abstraction System for Go. |
| [pdfcpu](https://github.com/pdfcpu/pdfcpu) | 1929 | 44 | 2017/06/18 | 22 hours ago | PDF processor. |
| [notify](https://github.com/rjeczalik/notify) | 615 | 30 | 2014/09/08 | 2 days ago | File system event notification library with simple API, similar to os/signal. |
| [copy](https://github.com/otiai10/copy) | 245 | 6 | 2017/09/01 | 5 days ago | Copy directory recursively. |
| [bigfile](https://github.com/bigfile/bigfile) | 181 | 13 | 2019/07/15 | 11 months ago | A file transfer system, support to manage files with http api, rpc call and ftp client. |
| [afs](https://github.com/viant/afs) | 120 | 12 | 2019/08/19 | 2 weeks ago | Abstract File Storage (mem, scp, zip, tar, cloud: s3, gs) for Go. |
| [vfs](https://github.com/C2FO/vfs) | 86 | 21 | 2017/08/01 | 6 days ago | A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS. |
| [go-csv-tag](https://github.com/artonge/go-csv-tag) | 77 | 1 | 2017/06/18 | 4 months ago | Load csv file using tag. |
| [opc](https://github.com/qmuntal/opc) | 66 | 3 | 2018/11/06 | 2 weeks ago | Load Open Packaging Conventions (OPC) files for Go. |
| [skywalker](https://github.com/dixonwille/skywalker) | 61 | 3 | 2017/08/01 | 3 years ago | Package to allow one to concurrently go through a filesystem with ease. |
| [go-exiftool](https://github.com/barasher/go-exiftool) | 46 | 3 | 2019/05/12 | 2 months ago | Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...). |
| [tarfs](https://github.com/posener/tarfs) | 45 | 2 | 2017/03/10 | 11 months ago | Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files. |
| [checksum](https://github.com/codingsince1985/checksum) | 31 | 2 | 2014/11/05 | 5 months ago | Compute message digest, like MD5 and SHA256, for large files. |
| [flop](https://github.com/homedepot/flop) | 24 | 16 | 2019/03/01 | 2 weeks ago | File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html). |
| [go-gtfs](https://github.com/artonge/go-gtfs) | 24 | 3 | 2017/07/09 | 4 months ago | Load gtfs files in go. |
| [baraka](https://github.com/xis/baraka) | 18 | 2 | 2020/07/12 | 3 days ago | A library to process http file uploads easily. |
| [gut](https://github.com/1set/gut) | 16 | 2 | 2019/10/05 | 2 months ago | Simple and reliable package for file operations like copy/move/diff/list on files, directories and symbolic links. |
| [parquet](https://github.com/parsyl/parquet) | 15 | 5 | 2019/01/29 | 6 days ago | Read and write [parquet](https://parquet.apache.org) files. |
| [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) | 13 | 3 | 2018/10/16 | 1 year ago | Copy files for humans. |
| [todotxt](https://github.com/1set/todotxt) | 7 | 1 | 2020/11/06 | 2 months ago | Go library for Gina Trapani's [*todo.txt*](http://todotxt.org/) files, supports parsing and manipulating of task lists in the [*todo.txt* format](https://github.com/todotxt/todo.txt). |

## Build Automation
        
*Libraries and tools helping with build automation.*

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [realize](https://github.com/oxequa/realize) | 3918 | 72 | 2016/07/12 | 3 months ago | Go build system with file watchers and live reload. Run, build and watch file changes with custom paths. |
| [task](https://github.com/go-task/task) | 3096 | 53 | 2017/02/27 | 5 days ago | simple "Make" alternative. |
| [mmake](https://github.com/tj/mmake) | 1550 | 28 | 2017/02/15 | 11 months ago | Modern Make. |
| [1build](https://github.com/gopinath-langote/1build) | 86 | 6 | 2019/04/23 | 6 months ago | Command line tool to frictionlessly manage project-specific commands. |
| [taskctl](https://github.com/taskctl/taskctl) | 84 | 5 | 2019/11/12 | 1 month ago | Concurrent task runner. |
| [gaper](https://github.com/maxcnunes/gaper) | 45 | 0 | 2018/06/16 | 1 year ago | Builds and restarts a Go project when it crashes or some watched file changes. |
| [taskflow](https://github.com/pellared/taskflow) | 45 | 2 | 2020/10/11 | 6 days ago | Create build pipelines in Go. |

# Tools
        
*Go software and plugins.*

## Code Analysis
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [lint](https://github.com/golang/lint) | 3813 | 109 | 2013/06/02 | 5 days ago | Golint is a linter for Go source code. |
| [errcheck](https://github.com/kisielk/errcheck) | 1643 | 25 | 2013/02/24 | 1 month ago | Errcheck is a program for checking for unchecked errors in Go programs. |
| [gcvis](https://github.com/davecheney/gcvis) | 1003 | 36 | 2014/07/10 | 1 year ago | Visualise Go program GC trace data in real time. |
| [go-critic](https://github.com/go-critic/go-critic) | 851 | 20 | 2018/05/05 | 1 week ago | source code linter that brings checks that are currently not implemented in other linters. |
| [php-parser](https://github.com/z7zmey/php-parser) | 782 | 25 | 2017/11/07 | 19 hours ago | A Parser for PHP written in Go. |
| [goast-viewer](https://github.com/yuroyoro/goast-viewer) | 491 | 17 | 2014/06/30 | 1 year ago | Web based Golang AST visualizer. |
| [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) | 463 | 6 | 2019/04/19 | 4 months ago | An easy way to find outdated dependencies of your Go projects. |
| [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) | 414 | 12 | 2017/04/12 | 9 months ago | go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects. |
| [goplantuml](https://github.com/jfeliu007/goplantuml) | 330 | 10 | 2019/05/26 | 5 days ago | Library and CLI that generates text plantump class diagram containing information about structures and interfaces with the relationship among them. |
| [unconvert](https://github.com/mdempsky/unconvert) | 305 | 8 | 2016/02/19 | 9 months ago | Remove unnecessary type conversions from Go source. |
| [gostatus](https://github.com/shurcooL/gostatus) | 244 | 7 | 2013/11/27 | 2 years ago | Command line tool, shows the status of repositories that contain Go packages. |
| [tickgit](https://github.com/augmentable-dev/tickgit) | 243 | 7 | 2019/10/12 | 7 months ago | CLI and go package for surfacing code comment TODOs (in any language) and applying a `git blame`to identify the author. |
| [dupl](https://github.com/mibk/dupl) | 224 | 7 | 2015/05/20 | 1 month ago | Tool for code clone detection. |
| [apicompat](https://github.com/bradleyfalzon/apicompat) | 171 | 8 | 2016/07/10 | 4 years ago | Checks recent changes to a Go project for backwards incompatible changes. |
| [checkstyle](https://github.com/qiniu/checkstyle) | 111 | 12 | 2014/01/01 | 2 months ago | checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments. |
| [golines](https://github.com/segmentio/golines) | 98 | 12 | 2019/10/01 | 1 week ago | Formatter that automatically shortens long lines in Go code. |
| [lint](https://github.com/surullabs/lint) | 66 | 5 | 2016/07/09 | 2 years ago | Run linters as part of go test. |
| [validate](https://github.com/mccoyst/validate) | 60 | 6 | 2013/11/22 | 4 years ago | Automatically validates struct fields with tags. |
| [go-outdated](https://github.com/firstrow/go-outdated) | 45 | 1 | 2015/06/29 | 2 years ago | Console application that displays outdated packages. |
| [blanket](https://github.com/verygoodsoftwarenotvirus/blanket) | 14 | 2 | 2017/09/04 | 2 years ago | tarp finds functions and methods without direct unit tests in Go source code. |

## Editor Plugins
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [vim-go](https://github.com/fatih/vim-go) | 13110 | 286 | 2014/03/24 | 10 hours ago | Go development plugin for Vim. |
| [vscode-go](https://github.com/microsoft/vscode-go) | 5978 | 226 | 2015/10/14 | 8 months ago | Extension for Visual Studio Code (VS Code) which provides support for the Go language. |
| [gocode](https://github.com/nsf/gocode) | 4910 | 194 | 2010/07/05 | 3 months ago | Autocompletion daemon for the Go programming language. |
| [GoSublime](https://github.com/DisposaBoy/GoSublime) | 3395 | 122 | 2011/08/27 | 6 months ago | Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features. |
| [go-plus](https://github.com/joefitzgerald/go-plus) | 1513 | 46 | 2014/03/13 | 3 days ago | Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting. |
| [vscode-go](https://github.com/golang/vscode-go) | 1464 | 44 | 2020/03/06 | 2 days ago | Extension for Visual Studio Code (VS Code) which provides support for the Go language. |
| [go-mode.el](https://github.com/dominikh/go-mode.el) | 1141 | 54 | 2013/01/30 | 1 week ago | Go mode for GNU/Emacs. |
| [Watch](https://github.com/eaburns/Watch) | 179 | 14 | 2013/08/08 | 2 years ago | Runs a command in an acme win on file changes. |
| [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) | 88 | 5 | 2012/11/25 | 4 years ago | Vim plugin to highlight syntax errors on save. |
| [go-language-server](https://github.com/theia-ide/go-language-server) | 33 | 4 | 2017/11/21 | 1 year ago | A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol. |
| [goimports-reviser](https://github.com/incu6us/goimports-reviser) | 26 | 2 | 2020/04/08 | 1 week ago | Formatting tool for imports. |
| [gounit-vim](https://github.com/hexdigest/gounit-vim) | 19 | 2 | 2018/02/21 | 2 years ago | Vim plugin for generating Go tests based on the function's or method's signature. |
| [theia-go-extension](https://github.com/theia-ide/theia-go-extension) | 15 | 5 | 2017/11/30 | 1 year ago | Go language support for the Theia IDE. |

## Go Generate Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gotests](https://github.com/cweill/gotests) | 3048 | 77 | 2016/01/19 | 1 month ago | Generate Go tests from your source code. |
| [genny](https://github.com/cheekybits/genny) | 1394 | 26 | 2014/10/27 | 2 months ago | Elegant generics for Go. |
| [re2dfa](https://github.com/opennota/re2dfa) | 181 | 9 | 2015/06/20 | 2 years ago | Transform regular expressions into finite state machines and output Go source code. |
| [hasgo](https://github.com/DylanMeeus/hasgo) | 81 | 5 | 2019/05/16 | 4 months ago | Generate Haskell inspired functions for your slices. |
| [xgen](https://github.com/xuri/xgen) | 68 | 6 | 2019/06/22 | 2 weeks ago | XSD (XML Schema Definition) parser and Go/C/Java/Rust/TypeScript code generator. |
| [gocontracts](https://github.com/Parquery/gocontracts) | 65 | 6 | 2018/08/13 | 2 years ago | brings design-by-contract to Go by synchronizing the code with the documentation. |
| [gounit](https://github.com/hexdigest/gounit) | 45 | 4 | 2018/02/05 | 2 years ago | Generate Go tests using your own templates. |
| [generic](https://github.com/usk81/generic) | 36 | 2 | 2016/06/15 | 1 month ago | flexible data type for Go. |

## Go Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-swagger](https://github.com/go-swagger/go-swagger) | 6099 | 122 | 2014/11/16 | 4 days ago | Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API. |
| [OctoLinker](https://github.com/OctoLinker/OctoLinker) | 4630 | 86 | 2013/12/27 | 6 days ago | Navigate through go files efficiently with the OctoLinker browser extension for GitHub. |
| [go-callvis](https://github.com/ofabry/go-callvis) | 3038 | 80 | 2016/09/03 | 2 weeks ago | Visualize call graph of your Go program using dot format. |
| [go-callvis](https://github.com/TrueFurby/go-callvis) | 2408 | 70 | 2016/09/03 | 11 months ago | Visualize call graph of your Go program using dot format. |
| [depth](https://github.com/KyleBanks/depth) | 580 | 13 | 2017/03/04 | 1 year ago | Visualize dependency trees of any package by analyzing imports. |
| [richgo](https://github.com/kyoh86/richgo) | 535 | 4 | 2017/01/04 | 3 weeks ago | Enrich `go test` outputs with text decorations. |
| [rts](https://github.com/galeone/rts) | 206 | 3 | 2016/04/04 | 7 months ago | RTS: response to struct. Generates Go structs from server responses. |
| [godbg](https://github.com/tylerwince/godbg) | 172 | 4 | 2019/01/23 | 1 year ago | Implementation of Rusts `dbg!` macro for quick and easy debugging during development. |
| [typex](https://github.com/dtgorski/typex) | 128 | 3 | 2020/03/24 | 3 weeks ago | Examine Go types and their transitive dependencies, alternatively export results as TypeScript value objects (or types) declaration. |
| [colorgo](https://github.com/songgao/colorgo) | 108 | 5 | 2013/02/14 | 7 months ago | Wrapper around `go` command for colorized `go build` output. |
| [gothanks](https://github.com/psampaz/gothanks) | 97 | 2 | 2019/11/10 | 1 month ago | GoThanks automatically stars your go.mod github dependencies, sending this way some love to their maintainers. |
| [igo](https://github.com/rocketlaunchr/igo) | 40 | 2 | 2018/11/17 | 10 months ago | Improved Go Syntax (transpiler) |
| [go-james](https://github.com/pieterclaerhout/go-james) | 38 | 2 | 2019/10/14 | 5 months ago | Go project skeleton creator, builds and tests your projects without the manual setup. |
| [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) | 37 | 2 | 2015/05/22 | 3 years ago | Bash completion for go and wgo. |
| [generator-go-lang](https://github.com/axelspringer/generator-go-lang) | 22 | 12 | 2017/09/13 | 10 months ago | A [Yeoman](http://yeoman.io) generator to get new Go projects started. |

## Software Packages
        
*Software written in Go.*

### DevOps Tools
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [kubernetes](https://github.com/kubernetes/kubernetes) | 74338 | 3292 | 2014/06/06 | 40 minutes ago | Container Cluster Manager from Google. |
| [moby](https://github.com/moby/moby) | 59602 | 3109 | 2013/01/18 | 22 hours ago | Collaborative project for the container ecosystem to assemble container-based systems. |
| [traefik](https://github.com/traefik/traefik) | 32602 | 700 | 2015/09/13 | 1 day ago | Reverse proxy and load balancer with support for multiple backends. |
| [traefik](https://github.com/containous/traefik) | 30724 | 719 | 2015/09/13 | 5 months ago | Reverse proxy and load balancer with support for multiple backends. |
| [gitea](https://github.com/go-gitea/gitea) | 23523 | 480 | 2016/11/01 | 21 minutes ago | Fork of Gogs, entirely community driven. |
| [vegeta](https://github.com/tsenart/vegeta) | 16541 | 315 | 2013/08/13 | 2 weeks ago | HTTP load testing tool and library. It's over 9000! |
| [packer](https://github.com/hashicorp/packer) | 12570 | 503 | 2013/03/23 | 5 hours ago | Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. |
| [hey](https://github.com/rakyll/hey) | 10280 | 174 | 2016/09/02 | 1 month ago | Hey is a tiny program that sends some load to a web application. |
| [webhook](https://github.com/adnanh/webhook) | 6207 | 142 | 2015/01/12 | 2 weeks ago | Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server. |
| [gvm](https://github.com/moovweb/gvm) | 5971 | 152 | 2011/12/03 | 1 week ago | GVM provides an interface to manage Go versions. |
| [gaia](https://github.com/gaia-pipeline/gaia) | 4322 | 106 | 2017/12/28 | 1 month ago | Build powerful pipelines in any programming language. |
| [gox](https://github.com/mitchellh/gox) | 3970 | 74 | 2013/11/17 | 1 year ago | Dead simple, no frills Go cross compile tool. |
| [bosun](https://github.com/bosun-monitor/bosun) | 3101 | 156 | 2013/11/15 | 3 months ago | Time Series Alerting Framework. |
| [bombardier](https://github.com/codesenberg/bombardier) | 2493 | 60 | 2016/05/29 | 3 months ago | Fast cross-platform HTTP benchmarking tool. |
| [pomerium](https://github.com/pomerium/pomerium) | 2335 | 28 | 2019/01/01 | 1 day ago | Pomerium is an identity-aware access proxy. |
| [script](https://github.com/bitfield/script) | 1701 | 31 | 2019/04/20 | 3 days ago | Making it easy to write shell-like scripts in Go for DevOps and system administration tasks. |
| [fac](https://github.com/mkchoi212/fac) | 1694 | 31 | 2017/12/29 | 1 year ago | Command-line user interface to fix git merge conflicts. |
| [goxc](https://github.com/laher/goxc) | 1661 | 49 | 2013/02/11 | 1 year ago | build tool for Go, with a focus on cross-compiling and packaging. |
| [kala](https://github.com/ajvb/kala) | 1616 | 66 | 2015/03/19 | 3 weeks ago | Simplistic, modern, and performant job scheduler. |
| [statusok](https://github.com/sanathp/statusok) | 1454 | 52 | 2015/08/26 | 3 weeks ago | Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected. |
| [s3gof3r](https://github.com/rlmcpherson/s3gof3r) | 1083 | 32 | 2013/08/02 | 5 months ago | Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3. |
| [go-selfupdate](https://github.com/sanbornm/go-selfupdate) | 805 | 27 | 2013/11/13 | 5 days ago | Enable your Go applications to self update. |
| [skm](https://github.com/TimothyYe/skm) | 670 | 20 | 2017/10/11 | 2 months ago | SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily! |
| [scaleway-cli](https://github.com/scaleway/scaleway-cli) | 649 | 35 | 2015/03/20 | 2 days ago | Manage BareMetal Servers from Command Line (as easily as with Docker). |
| [s5cmd](https://github.com/peak/s5cmd) | 544 | 22 | 2016/11/16 | 2 days ago | Blazing fast S3 and local filesystem execution tool. |
| [aurora](https://github.com/xuri/aurora) | 502 | 29 | 2016/10/09 | 3 months ago | Cross-platform web-based Beanstalkd queue server console. |
| [cassowary](https://github.com/rogerwelin/cassowary) | 475 | 5 | 2019/08/25 | 5 days ago | Modern cross-platform HTTP load-testing tool written in Go. |
| [govvv](https://github.com/ahmetb/govvv) | 471 | 10 | 2016/08/02 | 1 year ago | “go build” wrapper to easily add version information into Go binaries. |
| [utask](https://github.com/ovh/utask) | 359 | 24 | 2019/11/05 | 2 days ago | Automation engine that models and executes business processes declared in yaml. |
| [gonative](https://github.com/inconshreveable/gonative) | 325 | 8 | 2014/05/01 | 4 years ago | Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages. |
| [trubka](https://github.com/xitonix/trubka) | 296 | 15 | 2019/07/05 | 2 months ago | A CLI tool to manage and troubleshoot Apache Kafka clusters with the ability of generically publishing/consuming protocol buffer and plain text events to/from Kafka. |
| [mora](https://github.com/emicklei/mora) | 282 | 24 | 2013/07/12 | 4 years ago | REST server for accessing MongoDB documents and meta data. |
| [lstags](https://github.com/ivanilves/lstags) | 275 | 13 | 2017/08/15 | 3 months ago | Tool and API to sync Docker images across different registries. |
| [pewpew](https://github.com/bengadbois/pewpew) | 265 | 9 | 2016/10/12 | 3 weeks ago | Flexible HTTP command line stress tester. |
| [jenkins-cli](https://github.com/jenkins-zh/jenkins-cli) | 239 | 11 | 2019/06/21 | 1 day ago | Jenkins CLI allows you manage your Jenkins as an easy way. |
| [dogo](https://github.com/liudng/dogo) | 233 | 18 | 2014/11/19 | 1 year ago | Monitoring changes in the source file and automatically compile and run (restart). |
| [godbg](https://github.com/sirnewton01/godbg) | 225 | 18 | 2013/08/09 | 2 years ago | Web-based gdb front-end application. |
| [manssh](https://github.com/xwjdsh/manssh) | 225 | 4 | 2017/10/08 | 2 years ago | manssh is a command line tool for managing your ssh alias config easily. |
| [blast](https://github.com/dave/blast) | 194 | 4 | 2017/10/21 | 3 years ago | A simple tool for API load testing and batch jobs. |
| [gobrew](https://github.com/cryptojuice/gobrew) | 178 | 5 | 2013/11/13 | 8 months ago | gobrew lets you easily switch between multiple versions of go. |
| [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) | 175 | 7 | 2017/03/03 | 3 weeks ago | Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`. |
| [ostent](https://github.com/ostrost/ostent) | 170 | 6 | 2014/03/31 | 2 years ago | collects and displays system metrics and optionally relays to Graphite and/or InfluxDB. |
| [abbreviate](https://github.com/dnnrly/abbreviate) | 160 | 6 | 2018/11/23 | 3 months ago | abbreviate is a tool turning long strings in to shorter ones with configurable seperaters, for example to embed branch names in to deployment stack IDs. |
| [grapes](https://github.com/yaronsumel/grapes) | 152 | 7 | 2016/09/01 | 1 month ago | Lightweight tool designed to distribute commands over ssh with ease. |
| [kcli](https://github.com/cswank/kcli) | 145 | 5 | 2017/03/25 | 1 year ago | Command line tool for inspecting kafka topics/partitions/messages. |
| [terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi) | 132 | 11 | 2017/10/17 | 3 days ago | Terraform provider plugin that dynamically configures itself at runtime based on an OpenAPI document (formerly known as swagger file) containing the definitions of the APIs exposed. |
| [winrm-cli](https://github.com/masterzen/winrm-cli) | 111 | 6 | 2016/05/23 | 9 months ago | Cli tool to remotely execute commands on Windows machines. |
| [dockerfile-generator](https://github.com/ozankasikci/dockerfile-generator) | 92 | 5 | 2019/08/14 | 1 year ago | A go library and an executable that produces valid Dockerfiles using various input channels. |
| [drone-scp](https://github.com/appleboy/drone-scp) | 80 | 3 | 2016/10/16 | 4 months ago | Copy files and artifacts via SSH using a binary, docker or Drone CI. |
| [go-furnace](https://github.com/go-furnace/go-furnace) | 79 | 1 | 2016/10/09 | 1 year ago | Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean. |
| [dropship](https://github.com/ChrisMcKenzie/dropship) | 54 | 3 | 2015/09/03 | 2 years ago | Tool for deploying code via cdn. |
| [s3-proxy](https://github.com/oxyno-zeta/s3-proxy) | 33 | 2 | 2019/09/22 | 1 day ago | S3 Proxy with GET, PUT and DELETE methods and authentication (OpenID Connect and Basic Auth). |
| [rodent](https://github.com/alouche/rodent) | 31 | 2 | 2014/06/01 | 3 years ago | Rodent helps you manage Go versions, projects and track dependencies. |
| [drone-jenkins](https://github.com/appleboy/drone-jenkins) | 30 | 2 | 2016/10/15 | 4 months ago | Trigger downstream Jenkins jobs using a binary, docker or Drone CI. |
| [awsenv](https://github.com/soniah/awsenv) | 26 | 2 | 2015/08/05 | 2 years ago | Small binary that loads Amazon (AWS) environment variables for a profile. |
| [lwc](https://github.com/timdp/lwc) | 24 | 3 | 2018/04/22 | 9 months ago | A live-updating version of the UNIX wc command. |
| [depcharge](https://github.com/centerorbit/depcharge) | 14 | 3 | 2018/07/25 | 1 year ago | Helps orchestrating the execution of commands across the many dependencies in larger projects. |
| [httpref](https://github.com/dnnrly/httpref) | 13 | 2 | 2020/01/10 | 3 weeks ago | httpref is a handy CLI reference for HTTP methods, status codes, headers, and TCP and UDP ports. |
| [sg](https://github.com/ChristopherRabotin/sg) | 6 | 1 | 2015/08/19 | 4 years ago | Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response. |
| [aptly-fork](https://github.com/smira/aptly-fork) | 3 | 0 | 2019/07/04 | 1 year ago | aptly is a Debian repository management tool. |

### Other Software
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [goreplay](https://github.com/buger/goreplay) | 13874 | 469 | 2013/05/30 | 6 days ago | Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time. |
| [restic](https://github.com/restic/restic) | 11893 | 228 | 2014/04/27 | 3 hours ago | De-duplicating backup program. |
| [croc](https://github.com/schollz/croc) | 11681 | 209 | 2017/10/17 | 2 days ago | Easily and securely send files or folders from one computer to another. |
| [seaweedfs](https://github.com/chrislusf/seaweedfs) | 11441 | 536 | 2014/07/14 | 9 minutes ago | Fast, Simple and Scalable Distributed File System with O(1) disk seek. |
| [confd](https://github.com/kelseyhightower/confd) | 7324 | 259 | 2013/10/01 | 2 months ago | Manage local application configuration files using templates and data from etcd or consul. |
| [comcast](https://github.com/tylertreat/comcast) | 6840 | 155 | 2014/11/12 | 9 months ago | Simulate bad network connections. |
| [liteide](https://github.com/visualfc/liteide) | 6338 | 373 | 2012/11/19 | 1 day ago | LiteIDE is a simple, open source, cross-platform Go IDE. |
| [drive](https://github.com/odeke-em/drive) | 5853 | 197 | 2014/11/03 | 6 days ago | Google Drive client for the commandline. |
| [toxiproxy](https://github.com/Shopify/toxiproxy) | 5154 | 299 | 2014/09/04 | 2 weeks ago | Proxy to simulate network and system conditions for automated tests. |
| [nes](https://github.com/fogleman/nes) | 4680 | 149 | 2015/03/02 | 2 weeks ago | Nintendo Entertainment System (NES) emulator written in Go. |
| [duplicacy](https://github.com/gilbertchen/duplicacy) | 3589 | 95 | 2016/02/23 | 2 weeks ago | A cross-platform network and cloud backup tool based on the idea of lock-free deduplication. |
| [mylg](https://github.com/mehrdadrad/mylg) | 2413 | 112 | 2016/06/21 | 11 months ago | Command Line Network Diagnostic tool written in Go. |
| [goboy](https://github.com/Humpheh/goboy) | 2315 | 42 | 2017/08/20 | 6 months ago | Nintendo Game Boy Color emulator written in Go. |
| [sup](https://github.com/pressly/sup) | 2235 | 74 | 2015/02/23 | 11 months ago | Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers. |
| [scc](https://github.com/boyter/scc) | 2121 | 20 | 2018/03/01 | 4 days ago | Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates. |
| [lgo](https://github.com/yunabe/lgo) | 2116 | 48 | 2017/10/05 | 2 months ago | Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility. |
| [circuit](https://github.com/gocircuit/circuit) | 1875 | 137 | 2014/04/10 | 9 months ago | Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications. |
| [snap](https://github.com/intelsdi-x/snap) | 1796 | 145 | 2014/08/13 | 2 years ago | Powerful telemetry framework. |
| [borg](https://github.com/ok-borg/borg) | 1493 | 40 | 2016/09/10 | 3 years ago | Terminal based search engine for bash snippets. |
| [community](https://github.com/documize/community) | 1275 | 49 | 2016/04/29 | 2 months ago | Modern wiki software that integrates data from SaaS tools. |
| [Go-Package-Store](https://github.com/shurcooL/Go-Package-Store) | 888 | 20 | 2014/01/24 | 11 months ago | App that displays updates for the Go packages in your GOPATH. |
| [vflow](https://github.com/VerizonDigital/vflow) | 768 | 81 | 2017/02/24 | 1 week ago | High-performance, scalable and reliable IPFIX, sFlow and Netflow collector. |
| [peg](https://github.com/pointlander/peg) | 761 | 29 | 2010/04/25 | 3 months ago | Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator. |
| [shell2http](https://github.com/msoap/shell2http) | 702 | 23 | 2015/03/11 | 1 month ago | Executing shell commands via http server (for prototyping or remote control). |
| [leaps](https://github.com/Jeffail/leaps) | 697 | 27 | 2014/06/19 | 1 year ago | Pair programming service using Operational Transforms. |
| [gfile](https://github.com/Antonito/gfile) | 596 | 13 | 2019/03/08 | 1 year ago | Securely transfer files between two computers, without any third party, over WebRTC. |
| [gebug](https://github.com/moshebe/gebug) | 528 | 5 | 2020/07/20 | 3 days ago | A tool that makes debugging of Dockerized Go applications super easy by enabling Debugger and Hot-Reload features, seamlessly. |
| [guora](https://github.com/meloalright/guora) | 515 | 16 | 2020/08/13 | 2 months ago | A self-hosted Quora like web application written in Go. |
| [mockingjay-server](https://github.com/quii/mockingjay-server) | 472 | 9 | 2015/04/04 | 1 month ago | Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests. |
| [gocc](https://github.com/goccmack/gocc) | 455 | 24 | 2015/06/05 | 1 hour ago | Gocc is a compiler kit for Go written in Go. |
| [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) | 423 | 19 | 2015/10/08 | 10 months ago | Video streaming torrent client. |
| [ipe](https://github.com/dimiro1/ipe) | 314 | 18 | 2015/01/13 | 11 months ago | Open source Pusher server implementation compatible with Pusher client libraries written in GO. |
| [IDE](https://github.com/thestrukture/IDE) | 308 | 16 | 2017/09/09 | 1 month ago | Browser accessible IDE. Designed for Go with Go. |
| [wellington](https://github.com/wellington/wellington) | 295 | 13 | 2014/12/08 | 3 months ago | Sass project management tool, extends the language with sprite functions (like Compass). |
| [cherry](https://github.com/rafael-santiago/cherry) | 237 | 12 | 2015/10/24 | 3 years ago | Tiny webchat server in Go. |
| [tcpprobe](https://github.com/mehrdadrad/tcpprobe) | 186 | 7 | 2020/10/26 | 1 month ago | TCP tool for network performance and path monitoring, including socket statistics. |
| [woke](https://github.com/get-woke/woke) | 183 | 6 | 2020/08/31 | 1 day ago | Detect non-inclusive language in your source code. |
| [orange-cat](https://github.com/utatti/orange-cat) | 182 | 5 | 2014/11/01 | 2 years ago | Markdown previewer written in Go. |
| [orange-cat](https://github.com/hatashiro/orange-cat) | 182 | 5 | 2014/11/01 | 2 years ago | Markdown previewer written in Go. |
| [joincap](https://github.com/assafmo/joincap) | 157 | 7 | 2018/05/31 | 10 months ago | Command-line utility for merging multiple pcap files together. |
| [orbit](https://github.com/gulien/orbit) | 147 | 9 | 2017/05/13 | 3 weeks ago | A simple tool for running commands and generating files from templates. |
| [vaku](https://github.com/lingrino/vaku) | 101 | 3 | 2018/04/24 | 5 days ago | CLI & API for folder-based functions in Vault like copy, move, and search. |
| [boxed](https://github.com/tejo/boxed) | 74 | 2 | 2015/04/18 | 2 years ago | Dropbox based blog engine. |
| [dp](https://github.com/scryinfo/dp) | 60 | 9 | 2018/12/12 | 2 months ago | Through SDK for data exchange with blockchain, developers can get easy access to DAPP development. |
| [naclpipe](https://github.com/unix4fun/naclpipe) | 21 | 6 | 2015/05/05 | 2 years ago | Simple NaCL EC25519 based crypto pipe tool written in Go. |
| [term-quiz](https://github.com/crazcalm/term-quiz) | 18 | 1 | 2017/12/26 | 2 years ago | Quizzes for your terminal. |
| [snitch](https://github.com/lucasgomide/snitch) | 14 | 1 | 2017/04/06 | 2 years ago | Simple way to notify your team and many tools when someone has deployed any application via Tsuru. |
| [GoDocTooltip](https://github.com/diankong/GoDocTooltip) | 13 | 3 | 2016/01/21 | 2 weeks ago | Chrome extension for Go Doc sites, which shows function description as tooltip at function list. |

# Resources
        
*Where to discover new Go libraries.*

## Benchmarks
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) | 1470 | 59 | 2013/12/16 | 5 months ago | Go HTTP request router benchmark and comparison. |
| [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) | 1406 | 82 | 2016/04/06 | 3 hours ago | Go web framework benchmark. |
| [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) | 1100 | 37 | 2013/01/18 | 1 month ago | Benchmarks of Go serialization methods. |
| [skynet](https://github.com/atemerev/skynet) | 976 | 49 | 2016/02/14 | 2 weeks ago | Skynet 1M threads microbenchmark. |
| [speedtest-resize](https://github.com/fawick/speedtest-resize) | 207 | 7 | 2013/09/16 | 3 months ago | Compare various Image resize algorithms for the Go language. |
| [go-benchmarks](https://github.com/tylertreat/go-benchmarks) | 137 | 11 | 2016/02/25 | 5 years ago | Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches. |
| [gospeed](https://github.com/feyeleanor/gospeed) | 106 | 8 | 2011/05/23 | 2 years ago | Go micro-benchmarks for calculating the speed of language constructs. |
| [autobench](https://github.com/davecheney/autobench) | 89 | 8 | 2013/03/28 | 6 years ago | Framework to compare the performance between different Go versions. |
| [gocostmodel](https://github.com/mna/gocostmodel) | 55 | 5 | 2014/12/19 | 6 years ago | Benchmarks of common basic operations for the Go language. |
| [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) | 53 | 5 | 2014/09/24 | 2 years ago | Collection of benchmarks for popular Go database/SQL utilities. |
| [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) | 21 | 1 | 2017/01/24 | 3 years ago | Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results. |
| [kvbench](https://github.com/jimrobinson/kvbench) | 19 | 1 | 2014/04/15 | 1 year ago | Key/Value database benchmark. |
| [go-json-benchmark](https://github.com/zerosnake0/go-json-benchmark) | 4 | 2 | 2019/11/10 | 4 months ago | Go JSON benchmark. |

## Conferences
        

## E-Books
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [GoBooks](https://github.com/dariubs/GoBooks) | 8909 | 570 | 2015/05/05 | 3 days ago | A curated list of Go books. |
| [The-Golang-Standard-Library-by-Example](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example) | 7537 | 589 | 2013/04/14 | 2 weeks ago | Golang标准库。对于程序员而言，标准库与语言本身同样重要，它好比一个百宝箱，能为各种常见的任务提供完美的解决方案。以示例驱动的方式讲解Golang的标准库。 |
| [gosuccinctly](https://github.com/thedevsir/gosuccinctly) | 18 | 3 | 2018/09/02 | 2 years ago | in Persian. |

## Gophers
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [gophers](https://github.com/ashleymcnamara/gophers) | 2375 | 98 | 2017/02/15 | 1 year ago | Gopher artworks by Ashley McNamara. |
| [gophers](https://github.com/egonelbre/gophers) | 2271 | 56 | 2015/06/03 | 8 months ago | Free gophers. |
| [free-gophers-pack](https://github.com/MariaLetta/free-gophers-pack) | 2157 | 56 | 2019/04/02 | 7 months ago | Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster. |
| [gophericons](https://github.com/shalakhin/gophericons) | 577 | 19 | 2015/08/22 | 2 years ago | 34 gopher images for Go developers community |
| [gopher-stickers](https://github.com/tenntenn/gopher-stickers) | 493 | 15 | 2014/11/09 | 1 year ago | gopher stickers |
| [gopherize.me](https://github.com/matryer/gopherize.me) | 478 | 7 | 2017/01/25 | 5 months ago | Gopherize yourself. |
| [gopher-vector](https://github.com/golang-samples/gopher-vector) | 372 | 13 | 2013/03/31 | 4 years ago | Vector data of gopher |
| [go-gopher](https://github.com/sillecelik/go-gopher) | 86 | 0 | 2018/03/28 | 1 month ago | Gopher amigurumi toy pattern. |
| [gopher-logos](https://github.com/GolangUA/gopher-logos) | 83 | 9 | 2017/07/27 | 2 years ago | adorable gopher logos. |
| [gophers](https://github.com/rogeralsing/gophers) | 54 | 2 | 2017/01/28 | 6 months ago | random gopher graphics. |
| [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) | 37 | 1 | 2014/09/03 | 2 years ago | Go gopher Vector Data [.ai, .svg]. |

## Meetups
        
*Add the group of your city/country here (send **PR**)*

## Twitter
        

## Websites
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) | 27194 | 1746 | 2014/07/08 | 2 weeks ago | List of other amazingly awesome lists. |
| [awesome-remote-job](https://github.com/lukasz-madon/awesome-remote-job) | 19946 | 1004 | 2015/01/02 | 1 day ago | Curated list of awesome remote jobs. A lot of them are looking for Go hackers. |
| [golang-graphics](https://github.com/mholt/golang-graphics) | 141 | 9 | 2014/03/24 | 5 years ago | Collection of Go images, graphics, and art. |
| [gocryforhelp](https://github.com/ninedraft/gocryforhelp) | 38 | 12 | 2016/05/09 | 3 years ago | Collection of Go projects that needs help. Good place to start your open-source way in Go. |

### Tutorials
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang) | 36915 | 2501 | 2012/08/02 | 1 week ago | Golang ebook intro how to build a web app with golang. |
| [go-patterns](https://github.com/tmrts/go-patterns) | 14572 | 608 | 2015/12/14 | 3 weeks ago | Curated list of Go design patterns, recipes and idioms. |
| [learn-go-with-tests](https://github.com/quii/learn-go-with-tests) | 12828 | 289 | 2018/03/02 | 5 hours ago | Learn Go with test-driven development. |
| [golang-cheat-sheet](https://github.com/a8m/golang-cheat-sheet) | 5144 | 187 | 2014/02/13 | 3 months ago | Go's reference card. |
| [golang-for-nodejs-developers](https://github.com/miguelmota/golang-for-nodejs-developers) | 1635 | 36 | 2019/01/03 | 2 months ago | Examples of Golang compared to Node.js for learning. |
| [working-with-go](https://github.com/mkaz/working-with-go) | 1162 | 48 | 2014/05/04 | 1 year ago | Intro to go for experienced programmers. |
| [ethereum-development-with-go-book](https://github.com/miguelmota/ethereum-development-with-go-book) | 734 | 44 | 2018/05/16 | 2 weeks ago | A little e-book on Ethereum Development with Go. |
| [goapp](https://github.com/bnkamalesh/goapp) | 223 | 9 | 2020/07/04 | 1 month ago | An opinionated guideline to structure & develop a Go web application/service. |
| [design-patterns](https://github.com/shubhamzanwar/design-patterns) | 47 | 4 | 2020/09/24 | 3 months ago | Collection of programming design patterns implemented in Go. |

## Style Guides
        

| Go_repository    | Stars      | Watchers   | Created_at | Latest_push | Description |
| :--------- | ---------:| ---------:|:---------:|:---------:|:--------- |
| [go-styleguide](https://github.com/bahlo/go-styleguide) | 1035 | 28 | 2017/07/29 | 3 weeks ago | 🏆 Opinionated Styleguide for the Go language |

> 该项目源码[Awesome Go Analysis](https://github.com/plholx/awesome-go-analysis)
> 更专业的go开源项目分析请移步 [Awesome Go](https://go.libhunt.com/)
