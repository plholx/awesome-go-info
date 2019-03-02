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
    - [IoT](#iot)
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
        
* [EasyMIDI](https://github.com/algoGuy/EasyMIDI)  - EasyMidi is a simple and reliable library for working with standard midi file (SMF).
* [flac](https://github.com/eaburns/flac)  - No-frills native Go FLAC decoder that decodes FLAC files into byte slices.
* [flac](https://github.com/mewkiz/flac)  - Native Go FLAC encoder/decoder with support for FLAC streams.
* [gaad](https://github.com/Comcast/gaad)  - Native Go AAC bitstream parser.
* [go-sox](https://github.com/krig/go-sox)  - libsox bindings for go.
* [go_mediainfo](https://github.com/zhulik/go_mediainfo)  - libmediainfo bindings for go.
* [gosamplerate](https://github.com/dh1tw/gosamplerate)  - libsamplerate bindings for go.
* [id3v2](https://github.com/bogem/id3v2)  - Fast and stable ID3 parsing and writing library for Go.
* [malgo](https://github.com/gen2brain/malgo)  - Mini audio library.
* [minimp3](https://github.com/tosone/minimp3)  - Lightweight MP3 decoder library.
* [mix](https://github.com/go-mix/mix)  - Sequence-based Go-native audio mixer for music apps.
* [mp3](https://github.com/tcolgate/mp3)  - Native Go MP3 decoder.
* [music-theory](https://github.com/go-music-theory/music-theory)  - Music theory models in Go.
* [portaudio](https://github.com/gordonklaus/portaudio)  - Go bindings for the PortAudio audio I/O library.
* [portmidi](https://github.com/rakyll/portmidi)  - Go bindings for PortMidi.
* [go-taglib](https://github.com/wtolson/go-taglib)  - Go bindings for taglib.
* [vorbis](https://github.com/mccoyst/vorbis)  - "Native" Go Vorbis decoder (uses CGO, but has no dependencies).
* [waveform](https://github.com/mdlayher/waveform)  - Go package capable of generating waveform images from audio streams.

## Authentication and OAuth
        
*Libraries for implementing authentications schemes.*
        
* [authboss](https://github.com/volatiletech/authboss)  - Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time.
* [branca](https://github.com/hako/branca)  - Golang implementation of Branca Tokens.
* [casbin](https://github.com/casbin/casbin)  - Authorization library that supports access control models like ACL, RBAC, ABAC.
* [cookiestxt](https://github.com/mengzhuo/cookiestxt)  - provides parser of cookies.txt file format.
* [go-jose](https://github.com/square/go-jose)  - Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs.
* [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server)  - Standalone, specification-compliant,  OAuth2 server written in Golang.
* [gologin](https://github.com/dghubble/gologin)  - chainable handlers for login with OAuth1 and OAuth2 authentication providers.
* [gorbac](https://github.com/mikespook/gorbac)  - provides a lightweight role-based access control (RBAC) implementation in Golang.
* [goth](https://github.com/markbates/goth)  - provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box.
* [httpauth](https://github.com/goji/httpauth)  - HTTP Authentication middleware.
* [jwt](https://github.com/robbert229/jwt)  - Clean and easy to use implementation of JSON Web Tokens (JWT).
* [jwt](https://github.com/pascaldekloe/jwt)  - Lightweight JSON Web Token (JWT) library.
* [jwt-auth](https://github.com/adam-hanna/jwt-auth)  - JWT middleware for Golang http servers with many configuration options.
* [jwt-go](https://github.com/dgrijalva/jwt-go)  - Golang implementation of JSON Web Tokens (JWT).
* [loginsrv](https://github.com/tarent/loginsrv)  - JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam.
* [oauth2](https://github.com/golang/oauth2)  - Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support.
* [osin](https://github.com/openshift/osin)  - Golang OAuth2 server library.
* [paseto](https://github.com/o1egl/paseto)  - Golang implementation of Platform-Agnostic Security Tokens (PASETO).
* [permissions2](https://github.com/xyproto/permissions2)  - Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt.
* [rbac](https://github.com/zpatrick/rbac)  - Minimalistic RBAC package for Go applications.
* [securecookie](https://github.com/chmike/securecookie)  - Efficient secure cookie encoding/decoding.
* [session](https://github.com/icza/session)  - Go session management for web servers (including support for Google App Engine - GAE).
* [sessiongate-go](https://github.com/f0rmiga/sessiongate-go)  - Go session management using the SessionGate Redis module.
* [sessions](https://github.com/adam-hanna/sessions)  - Dead simple, highly performant, highly customizable sessions service for go http servers.
* [signedvalue](https://github.com/sashka/signedvalue)  - Signed and timestamped strings compatible with [Tornado's](https://github.com/tornadoweb/tornado) `create_signed_value`, `decode_signed_value`, and therefore `set_secure_cookie` and `get_secure_cookie`.

## Bot Building
        
*Libraries for building and working with bots.*
        
* [bot](https://github.com/go-chat-bot/bot)  - IRC, Slack & Telegram bot written in Go.
* [go-sarah](https://github.com/oklahomer/go-sarah)  - Framework to build bot for desired chat services including LINE, Slack, Gitter and more.
* [go-tgbot](https://github.com/olebedev/go-tgbot)  - Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware.
* [golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot)  - A golang implementation of a console-based trading bot for cryptocurrency exchanges.
* [govkbot](https://github.com/nikepan/govkbot)  - Simple Go [VK](https://vk.com) bot library.
* [hanu](https://github.com/sbstjn/hanu)  - Framework for writing Slack bots.
* [margelet](https://github.com/zhulik/margelet)  - Framework for building Telegram bots.
* [micha](https://github.com/onrik/micha)  - Go Library for Telegram bot api.
* [slacker](https://github.com/shomali11/slacker)  - Easy to use framework to create Slack bots.
* [tbot](https://github.com/yanzay/tbot)  - Telegram bot server with API similar to net/http.
* [telebot](https://github.com/tucnak/telebot)  - Telegram bot framework written in Go.
* [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api)  - Simple and clean Telegram bot client.
* [tenyks](https://github.com/kyleterry/tenyks)  - Service oriented IRC bot using Redis and JSON for messaging.

## Command Line
        
*Libraries for building Console Applications and Console User Interfaces.*
        
* [argparse](https://github.com/akamensky/argparse)  - Command line argument parser inspired by Python's argparse module.
* [argv](https://github.com/cosiner/argv)  - Go library to split command line string as arguments array using the bash syntax.
* [cli](https://github.com/mkideal/cli)  - Feature-rich and easy to use command-line package based on golang struct tags.
* [cli](https://github.com/teris-io/cli)  - Simple and complete API for building command line interfaces in Go.
* [gcli](https://github.com/tcnksm/gcli)  - The easy way to start building Golang command line applications.
* [cobra](https://github.com/spf13/cobra)  - Commander for modern Go CLI interactions.
* [commandeer](https://github.com/jaffee/commandeer)  - Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags.
* [complete](https://github.com/posener/complete)  - Write bash completions in Go + Go command bash completion.
* [docopt.go](https://github.com/docopt/docopt.go)  - Command-line arguments parser that will make you smile.
* [env](https://github.com/codingconcepts/env)  - Tag-based environment configuration for structs.
* [flag](https://github.com/cosiner/flag)  - Simple but powerful command line option parsing library for Go supporting subcommand.
* [flaggy](https://github.com/integrii/flaggy)  - A robust and idiomatic flags package with excellent subcommand support.
* [flagvar](https://github.com/sgreben/flagvar)  - A collection of flag argument types for Go's standard `flag` package.
* [go-arg](https://github.com/alexflint/go-arg)  - Struct-based argument parsing in Go.
* [go-commander](https://github.com/yitsushi/go-commander)  - Go library to simplify CLI workflow.
* [go-flags](https://github.com/jessevdk/go-flags)  - go command line option parser.
* [go-getoptions](https://github.com/DavidGamba/go-getoptions)  - Go option parser inspired on the flexibility of Perl’s GetOpt::Long.
* [gocmd](https://github.com/devfacet/gocmd)  - Go library for building command line applications.
* [hiboot](https://github.com/hidevopsio/hiboot)  - cli application framework with auto configuration and dependency injection.
* [kingpin](https://github.com/alecthomas/kingpin)  - Command line and flag parser supporting sub commands.
* [liner](https://github.com/peterh/liner)  - Go readline-like library for command-line interfaces.
* [cli](https://github.com/mitchellh/cli)  - Go library for implementing command-line interfaces.
* [mow.cli](https://github.com/jawher/mow.cli)  - Go library for building CLI applications with sophisticated flag and argument parsing and validation.
* [pflag](https://github.com/spf13/pflag)  - Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags.
* [readline](https://github.com/chzyer/readline)  - Pure golang implementation that provides most features in GNU-Readline under MIT license.
* [sand](https://github.com/Zaba505/sand)  - Simple API for creating interpreters and so much more.
* [sflags](https://github.com/octago/sflags)  - Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries.
* [strumt](https://github.com/antham/strumt)  - Library to create prompt chain.
* [clif](https://github.com/ukautz/clif)  - Small command line interface framework.
* [cli](https://github.com/urfave/cli)  - Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli).
* [wlog](https://github.com/dixonwille/wlog)  - Simple logging interface that supports cross-platform color and concurrency.
* [wmenu](https://github.com/dixonwille/wmenu)  - Easy to use menu structure for cli applications that prompts users to make choices.
* [asciigraph](https://github.com/guptarohit/asciigraph)  - Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies.
* [aurora](https://github.com/logrusorgru/aurora)  - ANSI terminal colors that supports fmt.Printf/Sprintf.
* [cfmt](https://github.com/mingrammer/cfmt)  - Contextual fmt inspired by bootstrap color classes.
* [chalk](https://github.com/ttacon/chalk)  - Intuitive package for prettifying terminal/console output.
* [color](https://github.com/fatih/color)  - Versatile package for colored terminal output.
* [colourize](https://github.com/TreyBastian/colourize)  - Go library for ANSI colour text in terminals.
* [ctc](https://github.com/wzshiming/ctc)  - The non-invasive cross-platform terminal color library does not need to modify the Print method.
* [go-ataman](https://github.com/workanator/go-ataman)  - Go library for rendering ANSI colored text templates in terminals.
* [go-colorable](https://github.com/mattn/go-colorable)  - Colorable writer for windows.
* [go-colortext](https://github.com/daviddengcn/go-colortext)  - Go library for color output in terminals.
* [go-isatty](https://github.com/mattn/go-isatty)  - isatty for golang.
* [go-prompt](https://github.com/c-bata/go-prompt)  - Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit).
* [gocui](https://github.com/jroimartin/gocui)  - Minimalist Go library aimed at creating Console User Interfaces.
* [gommon](https://github.com/labstack/gommon)  - Style terminal text.
* [color](https://github.com/gookit/color)  - Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows.
* [mpb](https://github.com/vbauerster/mpb)  - Multi progress bar for terminal applications.
* [progressbar](https://github.com/schollz/progressbar)  - Basic thread-safe progress bar that works in every OS.
* [simpletable](https://github.com/alexeyco/simpletable)  - Simple tables in terminal with Go.
* [tabular](https://github.com/InVisionApp/tabular)  - Print ASCII tables from command line utilities without the need to pass large sets of data to the API.
* [termbox-go](https://github.com/nsf/termbox-go)  - Termbox is a library for creating cross-platform text-based interfaces.
* [termdash](https://github.com/mum4k/termdash)  - Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui).
* [termtables](https://github.com/apcera/termtables)  - Go port of the Ruby library [terminal-tables](https://github.com/tj/terminal-table) for simple ASCII table generation as well as providing markdown and HTML output.
* [termui](https://github.com/gizak/termui)  - Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib).
* [uilive](https://github.com/gosuri/uilive)  - Library for updating terminal output in realtime.
* [uiprogress](https://github.com/gosuri/uiprogress)  - Flexible library to render progress bars in terminal applications.
* [uitable](https://github.com/gosuri/uitable)  - Library to improve readability in terminal apps using tabular data.

## Configuration
        
*Libraries for configuration parsing.*
        
* [config](https://github.com/olebedev/config)  - JSON or YAML configuration wrapper with environment variables and flags parsing.
* [configure](https://github.com/paked/configure)  - Provides configuration through multiple sources, including JSON, flags and environment variables.
* [confita](https://github.com/heetch/confita)  - Load configuration in cascade from multiple backends into a struct.
* [conflate](https://github.com/the4thamigo-uk/conflate)  - Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema.
* [env](https://github.com/caarlos0/env)  - Parse environment variables to Go structs (with defaults).
* [envcfg](https://github.com/tomazk/envcfg)  - Un-marshaling environment variables to Go structs.
* [envconf](https://github.com/ian-kent/envconf)  - Configuration from environment.
* [envconfig](https://github.com/vrischmann/envconfig)  - Read your configuration from environment variables.
* [envh](https://github.com/antham/envh)  - Helpers to manage environment variables.
* [gcfg](https://github.com/go-gcfg/gcfg)  - read INI-style configuration files into Go structs; supports user-defined types and subsections.
* [go-up](https://github.com/ufoscout/go-up)  - A simple configuration library with recursive placeholders resolution and no magic.
* [goconfig](https://github.com/crgimenes/goconfig)  - Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file.
* [godotenv](https://github.com/joho/godotenv)  - Go port of Ruby's dotenv library (Loads environment variables from `.env`).
* [gofigure](https://github.com/ian-kent/gofigure)  - Go application configuration made easy.
* [config](https://github.com/gookit/config)  - application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge.
* [hjson-go](https://github.com/hjson/hjson-go)  - Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments.
* [ingo](https://github.com/schachmat/ingo)  - Flags persisted in an ini-like config file.
* [ini](https://github.com/go-ini/ini)  - Go package to read and write INI files.
* [config](https://github.com/joshbetz/config)  - Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP.
* [envconfig](https://github.com/kelseyhightower/envconfig)  - Go library for managing configuration data from environment variables.
* [mini](https://github.com/sasbury/mini)  - Golang package for parsing ini-style configuration files.
* [sprbox](https://github.com/oblq/sprbox)  - Build-environment aware toolbox factory and agnostic config parser (YAML, TOML, JSON and Environment vars).
* [store](https://github.com/tucnak/store)  - Lightweight configuration manager for Go.
* [viper](https://github.com/spf13/viper)  - Go configuration with fangs.
* [xdg](https://github.com/OpenPeeDeeP/xdg)  - Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html).

## Continuous Integration
        
*Tools for help with continuous integration.*
        
* [drone](https://github.com/drone/drone)  - Drone is a Continuous Integration platform built on Docker, written in Go.
* [duci](https://github.com/duck8823/duci)  - A simple ci server no needs domain specific languages.
* [gomason](https://github.com/nikogura/gomason)  - Test, Build, Sign, and Publish your go binaries from a clean workspace.
* [goveralls](https://github.com/mattn/goveralls)  - Go integration for Coveralls.io continuous code coverage tracking system.
* [overalls](https://github.com/go-playground/overalls)  - Multi-Package go project coverprofile for tools like goveralls.
* [roveralls](https://github.com/lawrencewoodman/roveralls)  - Recursive coverage testing tool.

## CSS Preprocessors
        
*Libraries for preprocessing CSS files.*
        
* [gcss](https://github.com/yosssi/gcss)  - Pure Go CSS Preprocessor.
* [go-libsass](https://github.com/wellington/go-libsass)  - Go wrapper to the 100% Sass compatible libsass project.

## Data Structures
        
*Generic datastructures and algorithms in Go.*
        
* [algorithms](https://github.com/shady831213/algorithms)  - Algorithms and data structures.CLRS study.
* [binpacker](https://github.com/zhuangsirui/binpacker)  - Binary packer and unpacker helps user build custom binary stream.
* [bit](https://github.com/yourbasic/bit)  - Golang set data structure with bonus bit-twiddling functions.
* [bitset](https://github.com/willf/bitset)  - Go package implementing bitsets.
* [bloom](https://github.com/zentures/bloom)  - Bloom filters implemented in Go.
* [bloom](https://github.com/yourbasic/bloom)  - Golang Bloom filter implementation.
* [BoomFilters](https://github.com/tylertreat/BoomFilters)  - Probabilistic data structures for processing continuous, unbounded streams.
* [concurrent-writer](https://github.com/free/concurrent-writer)  - Highly concurrent drop-in replacement for `bufio.Writer`.
* [conjungo](https://github.com/InVisionApp/conjungo)  - A small, powerful and flexible merge library.
* [count-min-log](https://github.com/seiflotfy/count-min-log)  - Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory).
* [cuckoofilter](https://github.com/seiflotfy/cuckoofilter)  - Cuckoo filter: a good alternative to a counting bloom filter implemented in Go.
* [deque](https://github.com/edwingeng/deque)  - A highly optimized double-ended queue.
* [deque](https://github.com/gammazero/deque)  - Fast ring-buffer deque (double-ended queue).
* [encoding](https://github.com/zentures/encoding)  - Integer Compression Libraries for Go.
* [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree)  - Go implementation of Adaptive Radix Tree.
* [go-datastructures](https://github.com/Workiva/go-datastructures)  - Collection of useful, performant, and thread-safe data structures.
* [go-ef](https://github.com/amallia/go-ef)  - A Go implementation of the Elias-Fano encoding.
* [go-geoindex](https://github.com/hailocab/go-geoindex)  - In-memory geo index.
* [go-mcache](https://github.com/OrlovEvgeny/go-mcache)  - Fast in-memory key:value store/cache library. Pointer caches.
* [go-rquad](https://github.com/arl/go-rquad)  - Region quadtrees with efficient point location and neighbour finding.
* [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue)  - Concurrent FIFO queue.
* [gods](https://github.com/emirpasic/gods)  - Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc.
* [golang-set](https://github.com/deckarep/golang-set)  - Thread-Safe and Non-Thread-Safe high-performance sets for Go.
* [goset](https://github.com/zoumo/goset)  - A useful Set collection implementation for Go.
* [goskiplist](https://github.com/ryszard/goskiplist)  - Skip list implementation in Go.
* [gota](https://github.com/go-gota/gota)  - Implementation of dataframes, series, and data wrangling methods for Go.
* [hide](https://github.com/emvi/hide)  - ID type with marshalling to/from hash to prevent sending IDs to clients.
* [hilbert](https://github.com/google/hilbert)  - Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves.
* [hyperloglog](https://github.com/axiomhq/hyperloglog)  - HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction.
* [levenshtein](https://github.com/agext/levenshtein)  - Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix.
* [levenshtein](https://github.com/agnivade/levenshtein)  - Implementation to calculate levenshtein distance in Go.
* [mafsa](https://github.com/smartystreets/mafsa)  - MA-FSA implementation with Minimal Perfect Hashing.
* [merkletree](https://github.com/cbergoon/merkletree)  - Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures.
* [mspm](https://github.com/BlackRabbitt/mspm)  - Multi-String Pattern Matching Algorithm for information retrieval.
* [null](https://github.com/emvi/null)  - Nullable Go types that can be marshalled/unmarshalled to/from JSON.
* [pipeline](https://github.com/hyfather/pipeline)  - An implementation of pipelines with fan-in and fan-out.
* [ring](https://github.com/TheTannerRyan/ring)  - Go implementation of a high performance, thread safe bloom filter.
* [roaring](https://github.com/RoaringBitmap/roaring)  - Go package implementing compressed bitsets.
* [set](https://github.com/StudioSol/set)  - Simple set data structure implementation in Go using LinkedHashMap.
* [skiplist](https://github.com/MauriceGit/skiplist)  - Very fast Go Skiplist implementation.
* [skiplist](https://github.com/gansidui/skiplist)  - Skiplist implementation in Go.
* [timedmap](https://github.com/zekroTJA/timedmap)  - Map with expiring key-value pairs.
* [trie](https://github.com/derekparker/trie)  - Trie implementation in Go.
* [ttlcache](https://github.com/ReneKroon/ttlcache)  - In-memory LRU string-interface{} map with expiration for golang.
* [bloom](https://github.com/willf/bloom)  - Go package implementing Bloom filters.

## Database
        
*SQL query builder, libraries for building and using SQL.*
        
* [badger](https://github.com/dgraph-io/badger)  - Fast key-value store in Go.
* [bigcache](https://github.com/allegro/bigcache)  - Efficient key/value cache for gigabytes of data.
* [bolt](https://github.com/boltdb/bolt)  - Low-level key/value database for Go.
* [buntdb](https://github.com/tidwall/buntdb)  - Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support.
* [cache2go](https://github.com/muesli/cache2go)  - In-memory key:value cache which supports automatic invalidation based on timeouts.
* [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache)  - BigCache with clustering support and individual item expiration.
* [cockroach](https://github.com/cockroachdb/cockroach)  - Scalable, Geo-Replicated, Transactional Datastore.
* [couchcache](https://github.com/codingsince1985/couchcache)  - RESTful caching micro-service backed by Couchbase server.
* [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL)  - CovenantSQL is a SQL database on blockchain.
* [dgraph](https://github.com/dgraph-io/dgraph)  - Scalable, Distributed, Low Latency, High Throughput Graph Database.
* [diskv](https://github.com/peterbourgon/diskv)  - Home-grown disk-backed key-value store.
* [eliasdb](https://github.com/krotik/eliasdb)  - Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language.
* [fastcache](https://github.com/VictoriaMetrics/fastcache)  - fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead.
* [gcache](https://github.com/bluele/gcache)  - Cache library with support for expirable Cache, LFU, LRU and ARC.
* [go-cache](https://github.com/patrickmn/go-cache)  - In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.
* [goleveldb](https://github.com/syndtr/goleveldb)  - Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go.
* [gorocksdb](https://github.com/kapitan-k/gorocksdb)  - Gorocksdb is a wrapper for [RocksDB](https://rocksdb.org) written in Go.
* [groupcache](https://github.com/golang/groupcache)  - Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.
* [influxdb](https://github.com/influxdata/influxdb)  - Scalable datastore for metrics, events, and real-time analytics.
* [ledisdb](https://github.com/siddontang/ledisdb)  - Ledisdb is a high performance NoSQL like Redis based on LevelDB.
* [levigo](https://github.com/jmhodges/levigo)  - Levigo is a Go wrapper for LevelDB.
* [moss](https://github.com/couchbase/moss)  - Moss is a simple LSM key-value storage engine written in 100% Go.
* [piladb](https://github.com/fern4lvarez/piladb)  - Lightweight RESTful database engine based on stack data structures.
* [prometheus](https://github.com/prometheus/prometheus)  - Monitoring system and time series database.
* [pudge](https://github.com/recoilme/pudge)  - Fast and simple  key/value store written using Go's standard library.
* [rqlite](https://github.com/rqlite/rqlite)  - The lightweight, distributed, relational database built on SQLite.
* [golang-scribble](https://github.com/nanobox-io/golang-scribble)  - Tiny flat file JSON store.
* [slowpoke](https://github.com/recoilme/slowpoke)  - Key-value store with persistence.
* [tempdb](https://github.com/rafaeljesus/tempdb)  - Key-value store for temporary items.
* [tidb](https://github.com/pingcap/tidb)  - TiDB is a distributed SQL database. Inspired by the design of Google F1.
* [tiedot](https://github.com/HouzuoGuo/tiedot)  - Your NoSQL database powered by Golang.
* [vasto](https://github.com/chrislusf/vasto)  - A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption.
* [darwin](https://github.com/GuiaBolso/darwin)  - Database schema evolution library for Go.
* [go-fixtures](https://github.com/RichardKnop/go-fixtures)  - Django style fixtures for Golang's excellent built-in database/sql library.
* [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations)  - A Go package to help write migrations with go-pg/pg.
* [goose](https://github.com/steinbacher/goose)  - Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts.
* [gormigrate](https://github.com/go-gormigrate/gormigrate)  - Database schema migration helper for Gorm ORM.
* [migrate](https://github.com/golang-migrate/migrate)  - Database migrations. CLI and Golang library.
* [pravasan](https://github.com/pravasan/pravasan)  - Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc.
* [pop](https://github.com/gobuffalo/pop)  - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.
* [sql-migrate](https://github.com/rubenv/sql-migrate)  - Database migration tool. Allows embedding migrations into the application using go-bindata.
* [chproxy](https://github.com/Vertamedia/chproxy)  - HTTP proxy for ClickHouse database.
* [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk)  - Collects small insterts and sends big requests to ClickHouse servers.
* [dbbench](https://github.com/sj14/dbbench)  - Database benchmarking tool with support for several databases and scripts.
* [go-mysql](https://github.com/siddontang/go-mysql)  - Go toolset to handle MySQL protocol and replication.
* [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch)  - Sync your MySQL data into Elasticsearch automatically.
* [kingshard](https://github.com/flike/kingshard)  - kingshard is a high performance proxy for MySQL powered by Golang.
* [myreplication](https://github.com/2tvenom/myreplication)  - MySql binary log replication listener. Supports statement and row based replication.
* [octillery](https://github.com/knocknote/octillery)  - Go package for sharding databases ( Supports every ORM or raw SQL ).
* [orchestrator](https://github.com/github/orchestrator)  - MySQL replication topology manager & visualizer.
* [pgweb](https://github.com/sosedoff/pgweb)  - Web-based PostgreSQL database browser.
* [prep](https://github.com/hexdigest/prep)  - Use prepared SQL statements without changing your code.
* [prest](https://github.com/prest/prest)  - Serve a RESTful API from any PostgreSQL database.
* [rwdb](https://github.com/andizzle/rwdb)  - rwdb provides read replica capability for multiple database servers setup.
* [vitess](https://github.com/vitessio/vitess)  - vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services.
* [dotsql](https://github.com/gchaincl/dotsql)  - Go library that helps you keep sql files in one place and use them with ease.
* [gendry](https://github.com/didi/gendry)  - Non-invasive SQL builder and powerful data binder.
* [godbal](https://github.com/xujiajun/godbal)  - Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily.
* [goqu](https://github.com/doug-martin/goqu)  - Idiomatic SQL builder and query library.
* [igor](https://github.com/galeone/igor)  - Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax.
* [ormlite](https://github.com/pupizoid/ormlite)  - Lightweight package containing some ORM-like features and helpers for sqlite databases.
* [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx)  - Powerful data retrieval methods as well as DB-agnostic query building capabilities.
* [scaneo](https://github.com/variadico/scaneo)  - Generate Go code to convert database rows into arbitrary structs.
* [sqrl](https://github.com/elgris/sqrl)  - SQL query builder, fork of Squirrel with improved performance.
* [squirrel](https://github.com/Masterminds/squirrel)  - Go library that helps you build SQL queries.
* [xo](https://github.com/xo/xo)  - Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server.

## Database Drivers
        
*Libraries for connecting and operating databases.*
        

### Relational Databases
        
* [calcite-avatica-go](https://github.com/apache/calcite-avatica-go)  - Apache Avatica/Phoenix SQL driver for database/sql.
* [bgc](https://github.com/viant/bgc)  - Datastore Connectivity for BigQuery for go.
* [firebirdsql](https://github.com/nakagami/firebirdsql)  - Firebird RDBMS SQL driver for Go.
* [go-adodb](https://github.com/mattn/go-adodb)  - Microsoft ActiveX Object DataBase driver for go that uses database/sql.
* [go-mssqldb](https://github.com/denisenkom/go-mssqldb)  - Microsoft MSSQL driver for Go.
* [go-oci8](https://github.com/mattn/go-oci8)  - Oracle driver for go that uses database/sql.
* [mysql](https://github.com/go-sql-driver/mysql)  - MySQL driver for Go.
* [go-sqlite3](https://github.com/mattn/go-sqlite3)  - SQLite3 driver for go that uses database/sql.
* [gofreetds](https://github.com/minus5/gofreetds)  - Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org).
* [goracle](https://github.com/go-goracle/goracle)  - Oracle driver for Go, using the ODPI-C driver.
* [pgx](https://github.com/jackc/pgx)  - PostgreSQL driver supporting features beyond those exposed by database/sql.
* [pq](https://github.com/lib/pq)  - Pure Go Postgres driver for database/sql.

### NoSQL Databases
        
* [aerospike-client-go](https://github.com/aerospike/aerospike-client-go)  - Aerospike client in Go language.
* [arangolite](https://github.com/solher/arangolite)  - Lightweight golang driver for ArangoDB.
* [asc](https://github.com/viant/asc)  - Datastore Connectivity for Aerospike for go.
* [dynago](https://github.com/underarmour/dynago)  - Dynago is a principle of least surprise client for DynamoDB.
* [goforestdb](https://github.com/couchbase/goforestdb)  - Go bindings for ForestDB.
* [go-couchbase](https://github.com/couchbase/go-couchbase)  - Couchbase client in Go.
* [go-couchdb](https://github.com/fjl/go-couchdb)  - Yet another CouchDB HTTP API wrapper for Go.
* [go-pilosa](https://github.com/pilosa/go-pilosa)  - Go client library for Pilosa.
* [go-rejson](https://github.com/nitishm/go-rejson)  - Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease.
* [gocb](https://github.com/couchbase/gocb)  - Official Couchbase Go SDK.
* [godscache](https://github.com/defcronyke/godscache)  - A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached.
* [gomemcache](https://github.com/bradfitz/gomemcache)  - memcache client library for the Go programming language.
* [rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go)  - Go language driver for RethinkDB.
* [goriak](https://github.com/zegl/goriak)  - Go language driver for Riak KV.
* [mgo](https://github.com/globalsign/mgo)  - MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms.
* [mongo-go-driver](https://github.com/mongodb/mongo-go-driver)  - Official MongoDB driver for the Go language.
* [neo4j](https://github.com/cihangir/neo4j)  - Neo4j Rest API Bindings for Golang.
* [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO)  - Neo4j REST Client in golang.
* [neoism](https://github.com/jmcvetta/neoism)  - Neo4j client for Golang.
* [redigo](https://github.com/gomodule/redigo)  - Redigo is a Go client for the Redis database.
* [redis](https://github.com/go-redis/redis)  - Redis client for Golang.
* [redeo](https://github.com/bsm/redeo)  - Redis-protocol compatible TCP servers/services.
* [xredis](https://github.com/shomali11/xredis)  - Typesafe, customizable, clean & easy to use Redis client.
* [bleve](https://github.com/blevesearch/bleve)  - Modern text indexing library for go.
* [elastic](https://github.com/olivere/elastic)  - Elasticsearch client for Go.
* [elasticsql](https://github.com/cch123/elasticsql)  - Convert sql to elasticsearch dsl in Go.
* [elastigo](https://github.com/mattbaird/elastigo)  - Elasticsearch client library.
* [goes](https://github.com/OwnLocal/goes)  - Library to interact with Elasticsearch.
* [riot](https://github.com/go-ego/riot)  - Go Open Source, Distributed, Simple and efficient Search Engine.
* [skizze](https://github.com/seiflotfy/skizze)  - probabilistic data-structures service and storage.
* [cachego](https://github.com/fabiorphp/cachego)  - Golang Cache component for multiple drivers.
* [cayley](https://github.com/cayleygraph/cayley)  - Graph database with support for multiple backends.
* [dsc](https://github.com/viant/dsc)  - Datastore connectivity for SQL, NoSQL, structured files.
* [gokv](https://github.com/philippgille/gokv)  - Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more)

## Date and Time
        
*Libraries for working with dates and times.*
        
* [carbon](https://github.com/uniplaces/carbon)  - Simple Time extension with a lot of util methods, ported from PHP Carbon library.
* [date](https://github.com/rickb777/date)  - Augments Time for working with dates, date ranges, time spans, periods, and time-of-day.
* [dateparse](https://github.com/araddon/dateparse)  - Parse date's without knowing format in advance.
* [durafmt](https://github.com/hako/durafmt)  - Time duration formatting library for Go.
* [feiertage](https://github.com/wlbr/feiertage)  - Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving...
* [go-persian-calendar](https://github.com/yaa110/go-persian-calendar)  - The implementation of the Persian (Solar Hijri) Calendar in Go (golang).
* [go-sunrise](https://github.com/nathan-osman/go-sunrise)  - Calculate the sunrise and sunset times for a given location.
* [goweek](https://github.com/grsmv/goweek)  - Library for working with week entity in golang.
* [iso8601](https://github.com/relvacode/iso8601)  - Efficiently parse ISO8601 date-times without regex.
* [kair](https://github.com/GuilhermeCaruso/kair)  - Date and Time - Golang Formatting Library.
* [now](https://github.com/jinzhu/now)  - Now is a time toolkit for golang.
* [nulltime](https://github.com/kirillDanshin/nulltime)  - Nullable `time.Time`.
* [strftime](https://github.com/awoodbeck/strftime)  - C99-compatible strftime formatter.
* [timespan](https://github.com/SaidinWoT/timespan)  - For interacting with intervals of time, defined as a start time and a duration.
* [timeutil](https://github.com/leekchan/timeutil)  - Useful extensions (Timedelta, Strftime, ...) to the golang's time package.
* [tuesday](https://github.com/osteele/tuesday)  - Ruby-compatible Strftime function.

## Distributed Systems
        
*Packages that help with building Distributed Systems.*
        
* [celeriac.v1](https://github.com/svcavallar/celeriac.v1)  - Library for adding support for interacting and monitoring Celery workers, tasks and events in Go.
* [consistent](https://github.com/buraksezer/consistent)  - Consistent hashing with bounded loads.
* [digota](https://github.com/digota/digota)  - grpc ecommerce microservice.
* [doublejump](https://github.com/edwingeng/doublejump)  - A revamped Google's jump consistent hash.
* [drmaa](https://github.com/dgruber/drmaa)  - Job submission library for cluster schedulers based on the DRMAA standard.
* [dynatomic](https://github.com/tylfin/dynatomic)  - A library for using DynamoDB as an atomic counter.
* [emitter](https://github.com/emitter-io/emitter)  - High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love.
* [flowgraph](https://github.com/vectaport/flowgraph)  - flow-based programming package.
* [gleam](https://github.com/chrislusf/gleam)  - Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed.
* [glow](https://github.com/chrislusf/glow)  - Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go.
* [go-health](https://github.com/InVisionApp/go-health)  - Library for enabling asynchronous dependency health checks in your service.
* [go-jump](https://github.com/dgryski/go-jump)  - Port of Google's "Jump" Consistent Hash function.
* [kit](https://github.com/go-kit/kit)  - Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc.
* [gorpc](https://github.com/valyala/gorpc)  - Simple, fast and scalable RPC library for high load.
* [grpc-go](https://github.com/grpc/grpc-go)  - The Go language implementation of gRPC. HTTP/2 based RPC.
* [hprose-golang](https://github.com/hprose/hprose-golang)  - Very newbility RPC Library, support 25+ languages now.
* [jaeger](https://github.com/jaegertracing/jaeger)  - A distributed tracing system.
* [jsonrpc](https://github.com/osamingo/jsonrpc)  - The jsonrpc package helps implement of JSON-RPC 2.0.
* [jsonrpc](https://github.com/ybbus/jsonrpc)  - JSON-RPC 2.0 HTTP client implementation.
* [krakend](https://github.com/devopsfaith/krakend)  - Ultra performant API Gateway framework with middlewares.
* [micro](https://github.com/micro/micro)  - Pluggable microservice toolkit and distributed systems platform.
* [gnatsd](https://github.com/nats-io/gnatsd)  - Lightweight, high performance messaging system for microservices, IoT, and cloud native systems.
* [outboxer](https://github.com/italolelis/outboxer)  - Outboxer is a go library that implements the outbox pattern.
* [raft](https://github.com/hashicorp/raft)  - Golang implementation of the Raft consensus protocol, by HashiCorp.
* [etcd](https://github.com/etcd-io/etcd)  - Go implementation of the Raft consensus protocol, by CoreOS.
* [redis-lock](https://github.com/bsm/redis-lock)  - Simplified distributed locking implementation using Redis.
* [ringpop-go](https://github.com/uber/ringpop-go)  - Scalable, fault-tolerant application-layer sharding for Go applications.
* [rpcx](https://github.com/smallnest/rpcx)  - Distributed pluggable RPC service framework like alibaba Dubbo.
* [sleuth](https://github.com/ursiform/sleuth)  - Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)).
* [tendermint](https://github.com/tendermint/tendermint)  - High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols.
* [torrent](https://github.com/anacrolix/torrent)  - BitTorrent client package.
* [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix)  - Video streaming torrent client.

## Email
        
*Libraries and tools that implement email creation and sending.*
        
* [douceur](https://github.com/aymerick/douceur)  - CSS inliner for your HTML emails.
* [email](https://github.com/jordan-wright/email)  - A robust and flexible email library for Go.
* [go-dkim](https://github.com/toorop/go-dkim)  - DKIM library, to sign & verify email.
* [go-imap](https://github.com/emersion/go-imap)  - IMAP library for clients and servers.
* [go-message](https://github.com/emersion/go-message)  - Streaming library for the Internet Message Format and mail messages.
* [gomail](https://github.com/go-gomail/gomail)  - Gomail is a very simple and powerful package to send emails.
* [hectane](https://github.com/hectane/hectane)  - Lightweight SMTP client providing an HTTP API.
* [hermes](https://github.com/matcornic/hermes)  - Golang package that generates clean, responsive HTML e-mails.
* [MailHog](https://github.com/mailhog/MailHog)  - Email and SMTP testing with web and API interface.
* [sendgrid-go](https://github.com/sendgrid/sendgrid-go)  - SendGrid's Go library for sending email.
* [smtp](https://github.com/mailhog/smtp)  - SMTP server protocol state machine.

## Embeddable Scripting Languages
        
*Embedding other languages inside your go code.*
        
* [agora](https://github.com/mna/agora)  - Dynamically typed, embeddable programming language in Go.
* [anko](https://github.com/mattn/anko)  - Scriptable interpreter written in Go.
* [binder](https://github.com/alexeyco/binder)  - Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua).
* [expr](https://github.com/antonmedv/expr)  - an engine that can evaluate expressions.
* [gisp](https://github.com/jcla1/gisp)  - Simple LISP in Go.
* [go-duktape](https://github.com/olebedev/go-duktape)  - Duktape JavaScript engine bindings for Go.
* [go-lua](https://github.com/Shopify/go-lua)  - Port of the Lua 5.2 VM to pure Go.
* [go-php](https://github.com/deuill/go-php)  - PHP bindings for Go.
* [go-python](https://github.com/sbinet/go-python)  - naive go bindings to the CPython C-API.
* [golua](https://github.com/aarzilli/golua)  - Go bindings for Lua C API.
* [gopher-lua](https://github.com/yuin/gopher-lua)  - Lua 5.1 VM and compiler written in Go.
* [gval](https://github.com/PaesslerAG/gval)  - A highly customizable expression language written in Go.
* [ngaro](https://github.com/db47h/ngaro)  - Embeddable Ngaro VM implementation enabling scripting in Retro.
* [otto](https://github.com/robertkrimen/otto)  - JavaScript interpreter written in Go.
* [purl](https://github.com/ian-kent/purl)  - Perl 5.18.2 embedded in Go.
* [tengo](https://github.com/d5/tengo)  - Bytecode compiled script language for Go.

## Error Handling
        
*Libraries for handling errors.*
        
* [errors](https://github.com/pkg/errors)  - Package that provides simple error handling primitives.
* [errorx](https://github.com/joomcode/errorx)  - A feature rich error package with stack traces, composition of errors and more.
* [go-multierror](https://github.com/hashicorp/go-multierror)  - Go (golang) package for representing a list of errors as a single error.
* [tracerr](https://github.com/ztrue/tracerr)  - Golang errors with stack trace and source fragments.
* [werr](https://github.com/txgruppi/werr)  - Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called.

## Files
        
*Libraries for  handling files and file systems.*
        
* [afero](https://github.com/spf13/afero)  - FileSystem Abstraction System for Go.
* [go-csv-tag](https://github.com/artonge/go-csv-tag)  - Load csv file using tag.
* [go-decent-copy](https://github.com/hugocarreira/go-decent-copy)  - Copy files for humans.
* [go-gtfs](https://github.com/artonge/go-gtfs)  - Load gtfs files in go.
* [notify](https://github.com/rjeczalik/notify)  - File system event notification library with simple API, similar to os/signal.
* [opc](https://github.com/qmuntal/opc)  - Load Open Packaging Conventions (OPC) files for Go.
* [pdfcpu](https://github.com/hhrutter/pdfcpu)  - PDF processor.
* [skywalker](https://github.com/dixonwille/skywalker)  - Package to allow one to concurrently go through a filesystem with ease.
* [tarfs](https://github.com/posener/tarfs)  - Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files.

## Financial
        
*Packages for accounting and finance.*
        
* [accounting](https://github.com/leekchan/accounting)  - money and currency formatting for golang.
* [decimal](https://github.com/shopspring/decimal)  - Arbitrary-precision fixed-point decimal numbers.
* [go-finance](https://github.com/FlashBoys/go-finance)  - Comprehensive financial markets data in Go.
* [go-finance](https://github.com/alpeb/go-finance)  - Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations.
* [go-money](https://github.com/Rhymond/go-money)  - Implementation of Fowler's Money pattern.
* [ofxgo](https://github.com/aclindsa/ofxgo)  - Query OFX servers and/or parse the responses (with example command-line client).
* [orderbook](https://github.com/i25959341/orderbook)  - Matching Engine for Limit Order Book in Golang.
* [techan](https://github.com/sdcoffey/techan)  - Technical analysis library with advanced market analysis and trading strategies.
* [transaction](https://github.com/claygod/transaction)  - Embedded transactional database of accounts, running in multithreaded mode.
* [vat](https://github.com/dannyvankooten/vat)  - VAT number validation & EU VAT rates.

## Forms
        
*Libraries for working with forms.*
        
* [bind](https://github.com/robfig/bind)  - Bind form data to any Go values.
* [binding](https://github.com/mholt/binding)  - Binds form and JSON data from net/http Request to struct.
* [conform](https://github.com/leebenson/conform)  - Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags.
* [form](https://github.com/go-playground/form)  - Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support.
* [formam](https://github.com/monoculum/formam)  - decode form's values into a struct.
* [forms](https://github.com/albrow/forms)  - Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files.
* [csrf](https://github.com/gorilla/csrf)  - CSRF protection for Go web applications & services.
* [nosurf](https://github.com/justinas/nosurf)  - CSRF protection middleware for Go.

## Functional
        
*Packages to support functional programming in Go.*
        
* [fpGo](https://github.com/TeaEntityLab/fpGo)  - Monad, Functional Programming features for Golang.
* [fuego](https://github.com/seborama/fuego)  - Functional Experiment in Go.
* [go-underscore](https://github.com/tobyhede/go-underscore)  - Useful collection of helpfully functional Go collection utilities.

## Game Development
        
*Awesome game development libraries.*
        
* [engine](https://github.com/azul3d/engine)  - 3D game engine written in Go.
* [ebiten](https://github.com/hajimehoshi/ebiten)  - dead simple 2D game library in Go.
* [engo](https://github.com/EngoEngine/engo)  - Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm.
* [engine](https://github.com/g3n/engine)  - Go 3D Game Engine.
* [GarageEngine](https://github.com/vova616/GarageEngine)  - 2d game engine written in Go working on OpenGL.
* [glop](https://github.com/runningwild/glop)  - Glop (Game Library Of Power) is a fairly simple cross-platform game library.
* [go-astar](https://github.com/beefsack/go-astar)  - Go implementation of the A\* path finding algorithm.
* [go-collada](https://github.com/GlenKelley/go-collada)  - Go package for working with the Collada file format.
* [go-sdl2](https://github.com/veandco/go-sdl2)  - Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/).
* [go3d](https://github.com/ungerik/go3d)  - Performance oriented 2D/3D math package for Go.
* [gonet](https://github.com/xtaci/gonet)  - Game server skeleton implemented with golang.
* [goworld](https://github.com/xiaonanln/goworld)  - Scalable game server engine, featuring space-entity framework and hot-swapping.
* [leaf](https://github.com/name5566/leaf)  - Lightweight game server framework.
* [nano](https://github.com/lonng/nano)  - Lightweight, facility, high performance golang based game server framework.
* [oak](https://github.com/oakmound/oak)  - Pure Go game engine.
* [pitaya](https://github.com/topfreegames/pitaya)  - Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK.
* [pixel](https://github.com/faiface/pixel)  - Hand-crafted 2D game library in Go.
* [raylib-go](https://github.com/gen2brain/raylib-go)  - Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming.
* [termloop](https://github.com/JoelOtter/termloop)  - Terminal-based game engine for Go, built on top of Termbox.

## Generation and Generics
        
*Tools to enhance the language with features like generics via code generation.*
        
* [efaceconv](https://github.com/t0pep0/efaceconv)  - Code generation tool for high performance conversion from interface{} to immutable type without allocations.
* [gen](https://github.com/clipperhouse/gen)  - Code generation tool for ‘generics’-like functionality.
* [go-enum](https://github.com/abice/go-enum)  - Code generation for enums from code comments.
* [go-linq](https://github.com/ahmetb/go-linq)  - .NET LINQ-like query methods for Go.
* [goderive](https://github.com/awalterschulze/goderive)  - Derives functions from input types.
* [gotype](https://github.com/wzshiming/gotype)  - Golang source code parsing, usage like reflect package.
* [gowrap](https://github.com/hexdigest/gowrap)  - Generate decorators for Go interfaces using simple templates.
* [interfaces](https://github.com/rjeczalik/interfaces)  - Command line tool for generating interface definitions.
* [jennifer](https://github.com/dave/jennifer)  - Generate arbitrary Go code without templates.
* [pkgreflect](https://github.com/ungerik/pkgreflect)  - Go preprocessor for package scoped reflection.

## Geographic
        
*Geographic tools and servers*
        
* [geocache](https://github.com/melihmucuk/geocache)  - In-memory cache that is suitable for geolocation based applications.
* [geoserver](https://github.com/hishamkaram/geoserver)  - geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API.
* [gismanager](https://github.com/hishamkaram/gismanager)  - Publish Your GIS Data(Vector Data) to PostGIS and Geoserver.
* [osm](https://github.com/paulmach/osm)  - Library for reading, writing and working with OpenStreetMap data and APIs.
* [pbf](https://github.com/maguro/pbf)  - OpenStreetMap PBF golang encoder/decoder.
* [geo](https://github.com/golang/geo)  - S2 geometry library in Go.
* [tile38](https://github.com/tidwall/tile38)  - Geolocation DB with spatial index and realtime geofencing.

## Go Compilers
        
*Tools for compiling Go to other languages.*
        
* [c4go](https://github.com/Konstantin8105/c4go)  - Transpile C code to Go code.
* [f4go](https://github.com/Konstantin8105/f4go)  - Transpile FORTRAN 77 code to Go code.
* [gopherjs](https://github.com/gopherjs/gopherjs)  - Compiler from Go to JavaScript.
* [llgo](https://github.com/go-llvm/llgo)  - LLVM-based compiler for Go.
* [tardisgo](https://github.com/tardisgo/tardisgo)  - Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler.

## Goroutines
        
*Tools for managing and working with Goroutines.*
        
* [ants](https://github.com/panjf2000/ants)  - A high-performance goroutine pool for golang.
* [artifex](https://github.com/borderstech/artifex)  - Simple in-memory job queue for Golang using worker-based dispatching.
* [async](https://github.com/StudioSol/async)  - A safe way to execute functions asynchronously, recovering them in case of panic.
* [breaker](https://github.com/kamilsk/breaker)  - 🚧 Flexible mechanism to make your code breakable.
* [cyclicbarrier](https://github.com/marusama/cyclicbarrier)  - CyclicBarrier for golang.
* [go-floc](https://github.com/workanator/go-floc)  - Orchestrate goroutines with ease.
* [go-flow](https://github.com/kamildrazkiewicz/go-flow)  - Control goroutines execution order.
* [go-trylock](https://github.com/subchen/go-trylock)  - TryLock support on read-write lock for Golang.
* [GoSlaves](https://github.com/dgrr/GoSlaves)  - Simple and Asynchronous Goroutine pool library.
* [goworker](https://github.com/benmanns/goworker)  - goworker is a Go-based background worker.
* [gpool](https://github.com/sherifabdlnaby/gpool)  - manages a resizeable pool of context-aware goroutines to bound concurrency.
* [grpool](https://github.com/ivpusic/grpool)  - Lightweight Goroutine pool.
* [parallel-fn](https://github.com/rafaeljesus/parallel-fn)  - Run functions in parallel.
* [pool](https://github.com/go-playground/pool)  - Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation.
* [semaphore](https://github.com/kamilsk/semaphore)  - Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context.
* [semaphore](https://github.com/marusama/semaphore)  - Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations).
* [stl](https://github.com/ssgreg/stl)  - Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism.
* [threadpool](https://github.com/shettyh/threadpool)  - Golang threadpool implementation.
* [tunny](https://github.com/Jeffail/tunny)  - Goroutine pool for golang.
* [worker-pool](https://github.com/vardius/worker-pool)  - goworker is a Go simple async worker pool.
* [workerpool](https://github.com/gammazero/workerpool)  - Goroutine pool that limits the concurrency of task execution, not the number of tasks queued.

## GUI
        
*Interaction*
        
* [app](https://github.com/maxence-charriere/app)  - Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress.
* [fyne](https://github.com/fyne-io/fyne)  - Cross platform native GUIs designed for Go, rendered using EFL. Supports: Linux, macOS, Windows.
* [go-astilectron](https://github.com/asticode/go-astilectron)  - Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron).
* [go-sciter](https://github.com/sciter-sdk/go-sciter)  - Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform.
* [gotk3](https://github.com/gotk3/gotk3)  - Go bindings for GTK3.
* [gowd](https://github.com/dtylman/gowd)  - Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform.
* [qt](https://github.com/therecipe/qt)  - Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi).
* [ui](https://github.com/andlabs/ui)  - Platform-native GUI library for Go. Cross platform.
* [walk](https://github.com/lxn/walk)  - Windows application library kit for Go.
* [webview](https://github.com/zserge/webview)  - Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux).
* [gosx-notifier](https://github.com/deckarep/gosx-notifier)  - OSX Desktop Notifications library for Go.
* [robotgo](https://github.com/go-vgo/robotgo)  - Go Native cross-platform GUI system automation. Control the mouse, keyboard and other.
* [systray](https://github.com/getlantern/systray)  - Cross platform Go library to place an icon and menu in the notification area.
* [trayhost](https://github.com/shurcooL/trayhost)  - Cross-platform Go library to place an icon in the host operating system's taskbar.

## Hardware
        
*Libraries, tools, and tutorials for interacting with hardware.*
        

## Images
        
*Libraries for programming devices of the IoT.*
        
* [bild](https://github.com/anthonynsimon/bild)  - Collection of image processing algorithms in pure Go.
* [bimg](https://github.com/h2non/bimg)  - Small package for fast and efficient image processing using libvips.
* [cameron](https://github.com/aofei/cameron)  - An avatar generator for Go.
* [geopattern](https://github.com/pravj/geopattern)  - Create beautiful generative image patterns from a string.
* [gg](https://github.com/fogleman/gg)  - 2D rendering in pure Go.
* [gift](https://github.com/disintegration/gift)  - Package of image processing filters.
* [go-cairo](https://github.com/ungerik/go-cairo)  - Go binding for the cairo graphics library.
* [go-gd](https://github.com/bolknote/go-gd)  - Go binding for GD library.
* [go-nude](https://github.com/koyachi/go-nude)  - Nudity detection with Go.
* [go-opencv](https://github.com/go-opencv/go-opencv)  - Go bindings for OpenCV.
* [go-webcolors](https://github.com/jyotiska/go-webcolors)  - Port of webcolors library from Python to Go.
* [gocv](https://github.com/hybridgroup/gocv)  - Go package for computer vision using OpenCV 3.3+.
* [goimagehash](https://github.com/corona10/goimagehash)  - Go Perceptual image hashing package.
* [govatar](https://github.com/o1egl/govatar)  - Library and CMD tool for generating funny avatars.
* [image2ascii](https://github.com/qeesung/image2ascii)  - Convert image to ASCII.
* [imagick](https://github.com/gographics/imagick)  - Go binding to ImageMagick's MagickWand C API.
* [imaginary](https://github.com/h2non/imaginary)  - Fast and simple HTTP microservice for image resizing.
* [imaging](https://github.com/disintegration/imaging)  - Simple Go image processing package.
* [img](https://github.com/hawx/img)  - Selection of image manipulation tools.
* [ln](https://github.com/fogleman/ln)  - 3D line art rendering in Go.
* [mergi](https://github.com/noelyahan/mergi)  - Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate).
* [mort](https://github.com/aldor007/mort)  - Storage and image processing server written in Go.
* [mpo](https://github.com/donatj/mpo)  - Decoder and conversion tool for MPO 3D Photos.
* [picfit](https://github.com/thoas/picfit)  - An image resizing server written in Go.
* [pt](https://github.com/fogleman/pt)  - Path tracing engine written in Go.
* [resize](https://github.com/nfnt/resize)  - Image resizing for Go with common interpolation methods.
* [rez](https://github.com/bamiaux/rez)  - Image resizing in pure Go and SIMD.
* [smartcrop](https://github.com/muesli/smartcrop)  - Finds good crops for arbitrary images and crop sizes.
* [steganography](https://github.com/auyer/steganography)  - Pure Go Library for LSB steganography.
* [stegify](https://github.com/DimitarPetrov/stegify)  - Go tool for LSB steganography, capable of hiding any file within an image.
* [svgo](https://github.com/ajstarks/svgo)  - Go Language Library for SVG generation.
* [tga](https://github.com/ftrvxmtrx/tga)  - Package tga is a TARGA image format decoder/encoder.
* [connectordb](https://github.com/connectordb/connectordb)  - Open-Source Platform for Quantified Self & IoT.
* [devices](https://github.com/goiot/devices)  - Suite of libraries for IoT devices, experimental for x/exp/io.
* [eywa](https://github.com/xcodersun/eywa)  - Project Eywa is essentially a connection manager that keeps track of connected devices.
* [flogo](https://github.com/TIBCOSoftware/flogo)  - Project Flogo is an Open Source Framework for IoT Edge Apps & Integration.
* [gatt](https://github.com/paypal/gatt)  - Gatt is a Go package for building Bluetooth Low Energy peripherals.
* [gobot](https://github.com/hybridgroup/gobot)  - Gobot is a framework for robotics, physical computing, and the Internet of Things.
* [huego](https://github.com/amimof/huego)  - An extensive Philips Hue client library for Go.
* [iot](https://github.com/vaelen/iot)  - IoT is a simple framework for implementing a Google IoT Core device.
* [mainflux](https://github.com/mainflux/mainflux)  - Industrial IoT Messaging and Device Management Server.
* [sensorbee](https://github.com/sensorbee/sensorbee)  - Lightweight stream processing engine for IoT.

## IoT
        

## Job Scheduler
        
*Libraries for scheduling jobs.*
        
* [clockwork](https://github.com/whiteShtef/clockwork)  - Simple and intuitive job scheduling library in Go.
* [go-cron](https://github.com/rk/go-cron)  - Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons.
* [gron](https://github.com/roylee0704/gron)  - Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly.
* [jobrunner](https://github.com/bamzi/jobrunner)  - Smart and featureful cron job scheduler with job queuing and live monitoring built in.
* [jobs](https://github.com/albrow/jobs)  - Persistent and flexible background jobs library.
* [leprechaun](https://github.com/kilgaloon/leprechaun)  - Job scheduler that supports webhooks, crons and classic scheduling.
* [scheduler](https://github.com/carlescere/scheduler)  - Cronjobs scheduling made easy.

## JSON
        
*Libraries for working with JSON.*
        
* [gjson](https://github.com/tidwall/gjson)  - Get a JSON value with one line of code.
* [go-respond](https://github.com/nicklaw5/go-respond)  - Go package for handling common HTTP JSON responses.
* [gojq](https://github.com/elgs/gojq)  - JSON query in Golang.
* [gojson](https://github.com/ChimeraCoder/gojson)  - Automatically generate Go (golang) struct definitions from example JSON.
* [jaydiff](https://github.com/yazgazan/jaydiff)  - JSON diff utility written in Go.
* [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors)  - Go bindings based on the JSON API errors reference.
* [jsonf](https://github.com/miolini/jsonf)  - Console tool for highlighted formatting and struct query fetching JSON.
* [jsongo](https://github.com/ricardolonga/jsongo)  - Fluent API to make it easier to create Json objects.
* [jsonhal](https://github.com/RichardKnop/jsonhal)  - Simple Go package to make custom structs marshal into HAL compatible JSON responses.
* [kazaam](https://github.com/qntfy/kazaam)  - API for arbitrary transformation of JSON documents.
* [mp](https://github.com/sanbornm/mp)  - Simple cli email parser. It currently takes stdin and outputs JSON.

## Logging
        
*Libraries for generating and working with log files.*
        
* [gone](https://github.com/One-com/gone)  - Golang packages for writing small daemons and servers.
* [distillog](https://github.com/amoghe/distillog)  - distilled levelled logging (think of it as stdlib + log levels).
* [glg](https://github.com/kpango/glg)  - glg is simple and fast leveled logging library for Go.
* [glo](https://github.com/lajosbencz/glo)  - PHP Monolog inspired logging facility with identical severity levels.
* [glog](https://github.com/golang/glog)  - Leveled execution logs for Go.
* [go-cronowriter](https://github.com/utahta/go-cronowriter)  - Simple writer that rotate log files automatically based on current date and time, like cronolog.
* [go-log](https://github.com/subchen/go-log)  - Simple and configurable Logging in Go, with level, formatters and writers.
* [go-log](https://github.com/siddontang/go-log)  - Log lib supports level and multi handlers.
* [go-log](https://github.com/ian-kent/go-log)  - Log4j implementation in Go.
* [go-logger](https://github.com/apsdehal/go-logger)  - Simple logger of Go Programs, with level handlers.
* [gologger](https://github.com/sadlil/gologger)  - Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch.
* [gomol](https://github.com/aphistic/gomol)  - Multiple-output, structured logging for Go with extensible logging outputs.
* [journald](https://github.com/ssgreg/journald)  - Go implementation of systemd Journal's native API for logging.
* [log](https://github.com/apex/log)  - Structured logging package for Go.
* [log](https://github.com/go-playground/log)  - Simple, configurable and scalable Structured Logging for Go.
* [log](https://github.com/teris-io/log)  - Structured log interface for Go cleanly separates logging facade from its implementation.
* [logvoyage](https://github.com/firstrow/logvoyage)  - Full-featured logging saas written in golang.
* [log15](https://github.com/inconshreveable/log15)  - Simple, powerful logging for Go.
* [logdump](https://github.com/ewwwwwqm/logdump)  - Package for multi-level logging.
* [logex](https://github.com/chzyer/logex)  - Golang log lib, supports tracking and level, wrap by standard log lib.
* [logger](https://github.com/azer/logger)  - Minimalistic logging library for Go.
* [logmatic](https://github.com/borderstech/logmatic)  - Colorized logger for Golang with dynamic log level configuration.
* [logo](https://github.com/mbndr/logo)  - Golang logger to different configurable writers.
* [logrus](https://github.com/sirupsen/logrus)  - Structured logger for Go.
* [logrusly](https://github.com/sebest/logrusly)  - [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/).
* [logutils](https://github.com/hashicorp/logutils)  - Utilities for slightly better logging in Go (Golang) extending the standard logger.
* [logxi](https://github.com/mgutz/logxi)  - 12-factor app logger that is fast and makes you happy.
* [lumberjack](https://github.com/natefinch/lumberjack)  - Simple rolling logger, implements io.WriteCloser.
* [mlog](https://github.com/jbrodriguez/mlog)  - Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output.
* [onelog](https://github.com/francoispqt/onelog)  - Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenario. Also, it is one of the logger with the lowest allocation.
* [ozzo-log](https://github.com/go-ozzo/ozzo-log)  - High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail).
* [seelog](https://github.com/cihub/seelog)  - Logging functionality with flexible dispatching, filtering, and formatting.
* [go-spew](https://github.com/davecgh/go-spew)  - Implements a deep pretty printer for Go data structures to aid in debugging.
* [log](https://github.com/alexcesaro/log)  - Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs.
* [tail](https://github.com/hpcloud/tail)  - Go package striving to emulate the features of the BSD tail program.
* [xlog](https://github.com/xfxdev/xlog)  - Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format.
* [xlog](https://github.com/rs/xlog)  - Structured logger for `net/context` aware HTTP handlers with flexible dispatching.
* [zap](https://github.com/uber-go/zap)  - Fast, structured, leveled logging in Go.
* [zerolog](https://github.com/rs/zerolog)  - Zero-allocation JSON logger.

## Machine Learning
        
*Libraries for Machine Learning.*
        
* [bayesian](https://github.com/jbrukh/bayesian)  - Naive Bayesian Classification for Golang.
* [CloudForest](https://github.com/ryanbressler/CloudForest)  - Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go.
* [eaopt](https://github.com/MaxHalford/eaopt)  - An evolutionary optimization library.
* [evoli](https://github.com/khezen/evoli)  - Genetic Algorithm and Particle Swarm Optimization library.
* [fonet](https://github.com/Fontinalis/fonet)  - A Deep Neural Network library written in Go.
* [go-cluster](https://github.com/e-XpertSolutions/go-cluster)  - Go implementation of the k-modes and k-prototypes clustering algorithms.
* [go-deep](https://github.com/patrikeh/go-deep)  - A feature-rich neural network library in Go.
* [go-fann](https://github.com/white-pony/go-fann)  - Go bindings for Fast Artificial Neural Networks(FANN) library.
* [go-galib](https://github.com/thoj/go-galib)  - Genetic Algorithms library written in Go / golang.
* [go-pr](https://github.com/daviddengcn/go-pr)  - Pattern recognition package in Go lang.
* [gobrain](https://github.com/goml/gobrain)  - Neural Networks written in go.
* [godist](https://github.com/e-dard/godist)  - Various probability distributions, and associated methods.
* [goga](https://github.com/tomcraven/goga)  - Genetic algorithm library for Go.
* [golearn](https://github.com/sjwhitworth/golearn)  - General Machine Learning library for Go.
* [golinear](https://github.com/danieldk/golinear)  - liblinear bindings for Go.
* [gomind](https://github.com/surenderthakran/gomind)  - A simplistic Neural Network Library in Go.
* [goml](https://github.com/cdipaolo/goml)  - On-line Machine Learning in Go.
* [goRecommend](https://github.com/timkaye11/goRecommend)  - Recommendation Algorithms library written in Go.
* [gorgonia](https://github.com/gorgonia/gorgonia)  - graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms.
* [gorse](https://github.com/zhenghaoz/gorse)  - A High Performance Recommender System Package based on Collaborative Filtering for Go.
* [goscore](https://github.com/asafschers/goscore)  - Go Scoring API for PMML.
* [gosseract](https://github.com/otiai10/gosseract)  - Go package for OCR (Optical Character Recognition), by using Tesseract C++ library.
* [libsvm](https://github.com/datastream/libsvm)  - libsvm golang version derived work based on LIBSVM 3.14.
* [mlgo](https://github.com/NullHypothesis/mlgo)  - This project aims to provide minimalistic machine learning algorithms in Go.
* [neat](https://github.com/jinyeom/neat)  - Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT).

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
        



# Awesome Go
        

## Audio and Music
        
*Libraries for manipulating audio.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[EasyMIDI](https://github.com/algoGuy/EasyMIDI) | 17| 3| 2018-02-20 04:18:09 +0800 CST| 2018-03-22 13:16:54 +0800 CST|
|[flac](https://github.com/eaburns/flac) | 81| 5| 2013-08-21 01:48:58 +0800 CST| 2017-10-04 04:06:22 +0800 CST|
|[flac](https://github.com/mewkiz/flac) | 92| 6| 2012-11-02 04:13:58 +0800 CST| 2019-02-22 23:13:57 +0800 CST|
|[gaad](https://github.com/Comcast/gaad) | 50| 10| 2016-07-11 22:19:16 +0800 CST| 2018-02-21 05:08:16 +0800 CST|
|[go-sox](https://github.com/krig/go-sox) | 87| 8| 2013-10-08 22:11:04 +0800 CST| 2018-06-21 02:14:36 +0800 CST|
|[go_mediainfo](https://github.com/zhulik/go_mediainfo) | 21| 1| 2015-12-14 06:23:23 +0800 CST| 2015-12-25 04:40:36 +0800 CST|
|[gosamplerate](https://github.com/dh1tw/gosamplerate) | 8| 1| 2016-11-21 05:19:31 +0800 CST| 2018-06-04 07:37:17 +0800 CST|
|[id3v2](https://github.com/bogem/id3v2) | 90| 5| 2016-05-16 02:36:53 +0800 CST| 2019-02-08 23:52:24 +0800 CST|
|[malgo](https://github.com/gen2brain/malgo) | 55| 4| 2017-11-10 02:27:52 +0800 CST| 2019-02-02 18:30:11 +0800 CST|
|[minimp3](https://github.com/tosone/minimp3) | 19| 2| 2018-01-26 17:10:31 +0800 CST| 2019-03-02 12:02:21 +0800 CST|
|[mix](https://github.com/go-mix/mix) | 90| 3| 2016-01-03 23:53:57 +0800 CST| 2017-06-25 02:22:11 +0800 CST|
|[mp3](https://github.com/tcolgate/mp3) | 85| 1| 2015-02-27 04:37:37 +0800 CST| 2017-04-27 03:41:34 +0800 CST|
|[music-theory](https://github.com/go-music-theory/music-theory) | 233| 10| 2016-03-17 11:50:17 +0800 CST| 2018-06-13 04:13:26 +0800 CST|
|[portaudio](https://github.com/gordonklaus/portaudio) | 261| 9| 2015-09-16 15:59:23 +0800 CST| 2018-08-20 06:23:52 +0800 CST|
|[portmidi](https://github.com/rakyll/portmidi) | 187| 6| 2013-11-10 22:24:56 +0800 CST| 2017-07-16 11:23:45 +0800 CST|
|[go-taglib](https://github.com/wtolson/go-taglib) | 64| 6| 2012-11-20 09:03:40 +0800 CST| 2018-07-18 08:00:47 +0800 CST|
|[vorbis](https://github.com/mccoyst/vorbis) | 21| 3| 2013-07-12 10:45:39 +0800 CST| 2017-10-14 23:28:32 +0800 CST|
|[waveform](https://github.com/mdlayher/waveform) | 228| 13| 2014-09-14 02:07:36 +0800 CST| 2016-07-07 23:25:25 +0800 CST|

## Authentication and OAuth
        
*Libraries for implementing authentications schemes.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[authboss](https://github.com/volatiletech/authboss) | 1733| 38| 2015-01-03 13:12:02 +0800 CST| 2019-02-07 12:11:42 +0800 CST|
|[branca](https://github.com/hako/branca) | 49| 5| 2018-01-09 23:27:31 +0800 CST| 2018-08-08 08:04:55 +0800 CST|
|[casbin](https://github.com/casbin/casbin) | 3845| 140| 2017-04-08 15:51:23 +0800 CST| 2019-02-12 21:48:40 +0800 CST|
|[cookiestxt](https://github.com/mengzhuo/cookiestxt) | 2| 1| 2017-10-09 19:27:19 +0800 CST| 2017-10-09 20:14:58 +0800 CST|
|[go-jose](https://github.com/square/go-jose) | 1014| 61| 2014-11-15 02:27:31 +0800 CST| 2019-02-27 08:14:27 +0800 CST|
|[go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) | 1113| 71| 2015-11-01 21:30:09 +0800 CST| 2019-02-21 21:39:31 +0800 CST|
|[gologin](https://github.com/dghubble/gologin) | 960| 27| 2015-06-23 12:40:52 +0800 CST| 2019-02-14 07:18:24 +0800 CST|
|[gorbac](https://github.com/mikespook/gorbac) | 828| 55| 2013-12-26 18:00:41 +0800 CST| 2019-01-31 14:21:04 +0800 CST|
|[goth](https://github.com/markbates/goth) | 2081| 58| 2014-10-15 04:38:12 +0800 CST| 2019-02-22 03:00:59 +0800 CST|
|[httpauth](https://github.com/goji/httpauth) | 167| 5| 2014-05-27 06:53:57 +0800 CST| 2016-06-01 21:53:05 +0800 CST|
|[jwt](https://github.com/robbert229/jwt) | 60| 6| 2016-06-06 06:01:37 +0800 CST| 2018-11-09 01:39:15 +0800 CST|
|[jwt](https://github.com/pascaldekloe/jwt) | 38| 2| 2018-03-21 19:59:53 +0800 CST| 2019-01-18 06:59:13 +0800 CST|
|[jwt-auth](https://github.com/adam-hanna/jwt-auth) | 144| 9| 2016-07-06 07:31:43 +0800 CST| 2019-01-30 23:42:05 +0800 CST|
|[jwt-go](https://github.com/dgrijalva/jwt-go) | 5009| 130| 2012-04-18 09:41:49 +0800 CST| 2019-02-22 01:13:13 +0800 CST|
|[loginsrv](https://github.com/tarent/loginsrv) | 734| 45| 2016-11-11 20:11:21 +0800 CST| 2019-02-19 04:45:04 +0800 CST|
|[oauth2](https://github.com/golang/oauth2) | 2108| 92| 2014-04-14 23:07:35 +0800 CST| 2019-02-27 04:54:24 +0800 CST|
|[osin](https://github.com/openshift/osin) | 1506| 69| 2013-09-11 03:52:00 +0800 CST| 2018-10-17 05:40:43 +0800 CST|
|[paseto](https://github.com/o1egl/paseto) | 194| 9| 2018-01-23 13:27:39 +0800 CST| 2018-11-08 23:32:44 +0800 CST|
|[permissions2](https://github.com/xyproto/permissions2) | 297| 12| 2014-11-19 20:23:37 +0800 CST| 2019-02-21 11:07:58 +0800 CST|
|[rbac](https://github.com/zpatrick/rbac) | 20| 3| 2018-08-02 08:11:04 +0800 CST| 2018-08-30 03:03:47 +0800 CST|
|[securecookie](https://github.com/chmike/securecookie) | 26| 4| 2017-09-03 22:40:29 +0800 CST| 2018-08-31 23:04:40 +0800 CST|
|[session](https://github.com/icza/session) | 82| 6| 2016-02-08 17:07:07 +0800 CST| 2018-09-10 17:22:58 +0800 CST|
|[sessiongate-go](https://github.com/f0rmiga/sessiongate-go) | 6| 2| 2017-10-20 11:39:11 +0800 CST| 2018-11-10 03:25:29 +0800 CST|
|[sessions](https://github.com/adam-hanna/sessions) | 41| 3| 2017-04-29 09:09:28 +0800 CST| 2017-11-29 02:34:24 +0800 CST|
|[signedvalue](https://github.com/sashka/signedvalue) | 6| 0| 2018-01-06 13:57:09 +0800 CST| 2019-01-28 19:42:41 +0800 CST|

## Bot Building
        
*Libraries for building and working with bots.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[bot](https://github.com/go-chat-bot/bot) | 412| 34| 2015-09-23 00:41:13 +0800 CST| 2019-02-19 19:36:50 +0800 CST|
|[go-sarah](https://github.com/oklahomer/go-sarah) | 111| 3| 2016-11-06 18:04:43 +0800 CST| 2018-07-31 21:21:57 +0800 CST|
|[go-tgbot](https://github.com/olebedev/go-tgbot) | 80| 7| 2016-12-11 14:06:32 +0800 CST| 2018-06-25 12:50:26 +0800 CST|
|[golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot) | 147| 19| 2017-05-15 06:11:41 +0800 CST| 2019-02-11 21:02:32 +0800 CST|
|[govkbot](https://github.com/nikepan/govkbot) | 19| 2| 2016-07-12 06:09:54 +0800 CST| 2019-03-02 02:39:58 +0800 CST|
|[hanu](https://github.com/sbstjn/hanu) | 90| 6| 2016-09-16 15:10:42 +0800 CST| 2018-09-04 22:42:13 +0800 CST|
|[margelet](https://github.com/zhulik/margelet) | 51| 5| 2015-11-21 21:02:17 +0800 CST| 2016-09-18 19:47:01 +0800 CST|
|[micha](https://github.com/onrik/micha) | 9| 3| 2016-04-14 20:09:44 +0800 CST| 2018-02-15 19:45:17 +0800 CST|
|[slacker](https://github.com/shomali11/slacker) | 256| 13| 2017-05-20 09:41:20 +0800 CST| 2019-02-22 10:40:29 +0800 CST|
|[tbot](https://github.com/yanzay/tbot) | 185| 8| 2015-09-12 00:19:25 +0800 CST| 2018-12-16 04:28:17 +0800 CST|
|[telebot](https://github.com/tucnak/telebot) | 823| 34| 2015-06-26 03:27:50 +0800 CST| 2019-02-27 06:51:36 +0800 CST|
|[telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) | 1386| 56| 2015-06-25 13:33:57 +0800 CST| 2019-03-02 00:58:51 +0800 CST|
|[tenyks](https://github.com/kyleterry/tenyks) | 166| 14| 2012-08-26 10:02:24 +0800 CST| 2017-03-05 16:10:44 +0800 CST|

## Command Line
        
*Libraries for building Console Applications and Console User Interfaces.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[argparse](https://github.com/akamensky/argparse) | 74| 5| 2017-11-24 14:42:20 +0800 CST| 2019-01-15 17:47:22 +0800 CST|
|[argv](https://github.com/cosiner/argv) | 11| 1| 2017-01-22 18:37:21 +0800 CST| 2019-01-15 20:05:58 +0800 CST|
|[cli](https://github.com/mkideal/cli) | 427| 20| 2016-02-27 00:45:29 +0800 CST| 2019-02-28 10:37:39 +0800 CST|
|[cli](https://github.com/teris-io/cli) | 43| 1| 2017-05-25 07:07:07 +0800 CST| 2018-07-06 04:41:43 +0800 CST|
|[gcli](https://github.com/tcnksm/gcli) | 856| 25| 2014-06-19 17:10:15 +0800 CST| 2017-11-20 12:20:18 +0800 CST|
|[cobra](https://github.com/spf13/cobra) | 10748| 255| 2013-09-04 04:40:26 +0800 CST| 2019-02-23 08:16:26 +0800 CST|
|[commandeer](https://github.com/jaffee/commandeer) | 70| 4| 2017-10-12 10:51:05 +0800 CST| 2019-03-02 01:02:58 +0800 CST|
|[complete](https://github.com/posener/complete) | 540| 12| 2017-05-06 05:34:07 +0800 CST| 2019-02-20 19:52:52 +0800 CST|
|[docopt.go](https://github.com/docopt/docopt.go) | 1035| 36| 2013-08-26 07:05:52 +0800 CST| 2018-09-25 07:40:37 +0800 CST|
|[env](https://github.com/codingconcepts/env) | 37| 1| 2017-06-15 04:01:55 +0800 CST| 2018-10-16 16:32:45 +0800 CST|
|[flag](https://github.com/cosiner/flag) | 90| 5| 2016-10-06 00:49:41 +0800 CST| 2019-02-20 23:50:24 +0800 CST|
|[flaggy](https://github.com/integrii/flaggy) | 410| 11| 2018-03-05 13:55:05 +0800 CST| 2019-02-20 13:21:20 +0800 CST|
|[flagvar](https://github.com/sgreben/flagvar) | 26| 1| 2018-05-19 02:45:16 +0800 CST| 2018-11-02 05:59:23 +0800 CST|
|[go-arg](https://github.com/alexflint/go-arg) | 571| 15| 2015-11-01 09:30:06 +0800 CST| 2018-12-28 04:00:43 +0800 CST|
|[go-commander](https://github.com/yitsushi/go-commander) | 8| 1| 2016-10-10 18:09:41 +0800 CST| 2019-02-26 16:31:26 +0800 CST|
|[go-flags](https://github.com/jessevdk/go-flags) | 1290| 28| 2012-08-31 21:57:58 +0800 CST| 2019-02-25 03:46:59 +0800 CST|
|[go-getoptions](https://github.com/DavidGamba/go-getoptions) | 5| 1| 2015-12-18 10:21:14 +0800 CST| 2019-02-23 07:57:43 +0800 CST|
|[gocmd](https://github.com/devfacet/gocmd) | 28| 3| 2018-01-08 12:52:02 +0800 CST| 2018-09-05 11:42:24 +0800 CST|
|[hiboot](https://github.com/hidevopsio/hiboot) | 59| 6| 2018-03-16 19:21:46 +0800 CST| 2019-02-26 12:03:41 +0800 CST|
|[kingpin](https://github.com/alecthomas/kingpin) | 2296| 54| 2014-05-15 04:09:04 +0800 CST| 2019-01-14 14:40:04 +0800 CST|
|[liner](https://github.com/peterh/liner) | 520| 22| 2012-08-16 00:34:55 +0800 CST| 2019-02-11 10:05:18 +0800 CST|
|[cli](https://github.com/mitchellh/cli) | 925| 25| 2013-11-03 14:47:54 +0800 CST| 2018-11-25 14:16:29 +0800 CST|
|[mow.cli](https://github.com/jawher/mow.cli) | 590| 19| 2014-12-19 03:34:20 +0800 CST| 2019-02-26 14:55:10 +0800 CST|
|[pflag](https://github.com/spf13/pflag) | 623| 24| 2013-08-30 22:53:31 +0800 CST| 2019-02-11 19:02:58 +0800 CST|
|[readline](https://github.com/chzyer/readline) | 1305| 39| 2015-09-20 23:11:30 +0800 CST| 2019-02-08 20:42:45 +0800 CST|
|[sand](https://github.com/Zaba505/sand) | 2| 1| 2018-11-19 06:44:41 +0800 CST| 2018-11-22 03:13:47 +0800 CST|
|[sflags](https://github.com/octago/sflags) | 71| 5| 2016-12-04 22:49:27 +0800 CST| 2019-02-03 07:42:37 +0800 CST|
|[strumt](https://github.com/antham/strumt) | 22| 0| 2017-06-20 03:33:16 +0800 CST| 2018-02-17 19:11:54 +0800 CST|
|[clif](https://github.com/ukautz/clif) | 95| 2| 2015-05-31 02:30:08 +0800 CST| 2019-02-18 22:43:25 +0800 CST|
|[cli](https://github.com/urfave/cli) | 10066| 271| 2013-07-14 03:32:06 +0800 CST| 2019-02-08 08:35:01 +0800 CST|
|[wlog](https://github.com/dixonwille/wlog) | 29| 1| 2016-04-14 00:47:40 +0800 CST| 2017-07-21 07:52:41 +0800 CST|
|[wmenu](https://github.com/dixonwille/wmenu) | 67| 2| 2016-04-20 21:09:44 +0800 CST| 2019-02-21 08:45:53 +0800 CST|
|[asciigraph](https://github.com/guptarohit/asciigraph) | 1019| 23| 2018-06-17 18:37:16 +0800 CST| 2019-01-12 21:09:29 +0800 CST|
|[aurora](https://github.com/logrusorgru/aurora) | 425| 5| 2016-11-07 06:37:12 +0800 CST| 2018-10-03 03:45:21 +0800 CST|
|[cfmt](https://github.com/mingrammer/cfmt) | 53| 3| 2018-03-16 03:04:27 +0800 CST| 2018-12-08 01:31:52 +0800 CST|
|[chalk](https://github.com/ttacon/chalk) | 286| 11| 2014-07-19 03:38:58 +0800 CST| 2016-06-27 04:24:23 +0800 CST|
|[color](https://github.com/fatih/color) | 2861| 56| 2014-02-17 17:13:35 +0800 CST| 2018-10-11 07:14:14 +0800 CST|
|[colourize](https://github.com/TreyBastian/colourize) | 15| 3| 2015-05-11 19:49:39 +0800 CST| 2016-05-10 17:50:02 +0800 CST|
|[ctc](https://github.com/wzshiming/ctc) | 6| 1| 2018-04-28 02:07:42 +0800 CST| 2018-10-31 17:40:47 +0800 CST|
|[go-ataman](https://github.com/workanator/go-ataman) | 8| 2| 2017-05-18 03:04:57 +0800 CST| 2017-09-25 14:09:08 +0800 CST|
|[go-colorable](https://github.com/mattn/go-colorable) | 334| 16| 2014-07-30 10:38:06 +0800 CST| 2019-02-23 00:32:50 +0800 CST|
|[go-colortext](https://github.com/daviddengcn/go-colortext) | 188| 9| 2013-01-23 11:38:54 +0800 CST| 2018-04-10 01:50:54 +0800 CST|
|[go-isatty](https://github.com/mattn/go-isatty) | 298| 7| 2014-04-01 09:53:09 +0800 CST| 2019-02-26 01:38:53 +0800 CST|
|[go-prompt](https://github.com/c-bata/go-prompt) | 2078| 35| 2017-08-15 00:02:09 +0800 CST| 2019-02-23 16:08:50 +0800 CST|
|[gocui](https://github.com/jroimartin/gocui) | 3962| 103| 2014-01-04 10:50:20 +0800 CST| 2019-02-20 23:24:10 +0800 CST|
|[gommon](https://github.com/labstack/gommon) | 269| 19| 2015-03-13 06:35:57 +0800 CST| 2019-01-26 02:56:11 +0800 CST|
|[color](https://github.com/gookit/color) | 114| 3| 2018-07-01 15:28:17 +0800 CST| 2019-02-22 13:35:31 +0800 CST|
|[mpb](https://github.com/vbauerster/mpb) | 419| 8| 2016-12-14 19:56:29 +0800 CST| 2019-02-26 17:24:45 +0800 CST|
|[progressbar](https://github.com/schollz/progressbar) | 494| 13| 2017-10-27 02:28:10 +0800 CST| 2019-01-31 01:37:54 +0800 CST|
|[simpletable](https://github.com/alexeyco/simpletable) | 119| 2| 2017-03-29 15:27:23 +0800 CST| 2019-02-23 00:50:45 +0800 CST|
|[tabular](https://github.com/InVisionApp/tabular) | 23| 3| 2018-04-24 05:17:03 +0800 CST| 2018-05-15 03:04:57 +0800 CST|
|[termbox-go](https://github.com/nsf/termbox-go) | 3286| 103| 2012-01-13 05:03:03 +0800 CST| 2019-02-13 09:09:28 +0800 CST|
|[termdash](https://github.com/mum4k/termdash) | 78| 6| 2018-03-24 20:01:49 +0800 CST| 2019-02-28 13:51:06 +0800 CST|
|[termtables](https://github.com/apcera/termtables) | 203| 68| 2012-12-06 08:22:52 +0800 CST| 2017-10-30 17:24:27 +0800 CST|
|[termui](https://github.com/gizak/termui) | 8353| 275| 2015-02-03 22:09:27 +0800 CST| 2019-03-02 06:13:46 +0800 CST|
|[uilive](https://github.com/gosuri/uilive) | 723| 11| 2015-11-16 14:13:10 +0800 CST| 2018-12-04 09:38:28 +0800 CST|
|[uiprogress](https://github.com/gosuri/uiprogress) | 1228| 27| 2015-11-17 08:59:24 +0800 CST| 2018-10-12 09:44:27 +0800 CST|
|[uitable](https://github.com/gosuri/uitable) | 464| 15| 2015-11-14 05:59:21 +0800 CST| 2017-08-23 15:41:33 +0800 CST|

## Configuration
        
*Libraries for configuration parsing.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[config](https://github.com/olebedev/config) | 191| 8| 2014-04-21 23:09:39 +0800 CST| 2018-09-10 23:57:18 +0800 CST|
|[configure](https://github.com/paked/configure) | 43| 3| 2015-06-14 15:46:56 +0800 CST| 2019-02-18 22:01:49 +0800 CST|
|[confita](https://github.com/heetch/confita) | 208| 24| 2017-12-21 18:49:18 +0800 CST| 2019-02-14 23:22:33 +0800 CST|
|[conflate](https://github.com/the4thamigo-uk/conflate) | 3| 0| 2018-02-02 03:06:15 +0800 CST| 2019-02-09 18:40:00 +0800 CST|
|[env](https://github.com/caarlos0/env) | 697| 16| 2015-07-28 10:14:37 +0800 CST| 2019-02-18 04:15:59 +0800 CST|
|[envcfg](https://github.com/tomazk/envcfg) | 89| 1| 2014-11-29 19:43:53 +0800 CST| 2017-06-19 23:53:22 +0800 CST|
|[envconf](https://github.com/ian-kent/envconf) | 7| 1| 2014-10-26 20:12:26 +0800 CST| 2014-10-26 20:12:40 +0800 CST|
|[envconfig](https://github.com/vrischmann/envconfig) | 135| 4| 2015-04-22 07:37:17 +0800 CST| 2018-12-28 23:48:21 +0800 CST|
|[envh](https://github.com/antham/envh) | 91| 3| 2017-01-12 19:25:48 +0800 CST| 2019-03-02 05:12:10 +0800 CST|
|[gcfg](https://github.com/go-gcfg/gcfg) | 102| 4| 2015-08-17 22:40:55 +0800 CST| 2018-05-18 05:53:51 +0800 CST|
|[go-up](https://github.com/ufoscout/go-up) | 19| 1| 2018-02-18 17:50:00 +0800 CST| 2018-07-19 15:23:56 +0800 CST|
|[goconfig](https://github.com/crgimenes/goconfig) | 91| 11| 2016-12-18 19:22:41 +0800 CST| 2019-02-06 17:17:32 +0800 CST|
|[godotenv](https://github.com/joho/godotenv) | 1741| 29| 2013-07-30 15:45:19 +0800 CST| 2019-02-26 16:05:57 +0800 CST|
|[gofigure](https://github.com/ian-kent/gofigure) | 56| 6| 2014-11-25 08:12:40 +0800 CST| 2017-05-03 03:22:48 +0800 CST|
|[config](https://github.com/gookit/config) | 42| 3| 2018-07-07 16:11:39 +0800 CST| 2019-01-19 09:54:56 +0800 CST|
|[hjson-go](https://github.com/hjson/hjson-go) | 162| 7| 2016-08-06 06:59:18 +0800 CST| 2019-02-09 10:37:18 +0800 CST|
|[ingo](https://github.com/schachmat/ingo) | 22| 1| 2016-02-08 06:57:40 +0800 CST| 2017-04-03 09:15:10 +0800 CST|
|[ini](https://github.com/go-ini/ini) | 1324| 62| 2014-12-18 15:36:37 +0800 CST| 2019-02-18 03:54:22 +0800 CST|
|[config](https://github.com/joshbetz/config) | 193| 3| 2017-04-03 02:37:05 +0800 CST| 2017-08-11 11:04:14 +0800 CST|
|[envconfig](https://github.com/kelseyhightower/envconfig) | 2119| 37| 2013-11-07 01:01:55 +0800 CST| 2019-02-17 16:18:49 +0800 CST|
|[mini](https://github.com/sasbury/mini) | 19| 2| 2015-04-30 07:52:36 +0800 CST| 2018-12-27 07:28:05 +0800 CST|
|[sprbox](https://github.com/oblq/sprbox) | 3| 2| 2018-07-17 08:27:35 +0800 CST| 2018-10-31 22:39:30 +0800 CST|
|[store](https://github.com/tucnak/store) | 239| 5| 2015-10-04 03:17:28 +0800 CST| 2017-09-05 19:38:35 +0800 CST|
|[viper](https://github.com/spf13/viper) | 7652| 165| 2014-04-02 22:33:33 +0800 CST| 2019-03-01 21:28:58 +0800 CST|
|[xdg](https://github.com/OpenPeeDeeP/xdg) | 18| 3| 2017-07-20 23:58:29 +0800 CST| 2019-02-21 08:38:40 +0800 CST|

## Continuous Integration
        
*Tools for help with continuous integration.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[drone](https://github.com/drone/drone) | 17332| 584| 2014-02-07 15:54:44 +0800 CST| 2019-03-01 22:27:32 +0800 CST|
|[duci](https://github.com/duck8823/duci) | 30| 3| 2018-04-01 09:51:02 +0800 CST| 2019-03-02 09:22:43 +0800 CST|
|[gomason](https://github.com/nikogura/gomason) | 20| 0| 2017-11-18 08:59:11 +0800 CST| 2019-02-09 04:31:35 +0800 CST|
|[goveralls](https://github.com/mattn/goveralls) | 540| 16| 2013-04-17 18:58:40 +0800 CST| 2019-02-03 00:53:12 +0800 CST|
|[overalls](https://github.com/go-playground/overalls) | 94| 4| 2015-07-30 19:30:11 +0800 CST| 2018-08-26 00:45:52 +0800 CST|
|[roveralls](https://github.com/lawrencewoodman/roveralls) | 11| 1| 2016-06-26 15:45:32 +0800 CST| 2017-11-20 03:39:13 +0800 CST|

## CSS Preprocessors
        
*Libraries for preprocessing CSS files.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[gcss](https://github.com/yosssi/gcss) | 408| 16| 2014-09-04 22:38:20 +0800 CST| 2014-10-12 22:07:10 +0800 CST|
|[go-libsass](https://github.com/wellington/go-libsass) | 116| 4| 2015-04-19 23:09:47 +0800 CST| 2019-02-12 20:41:27 +0800 CST|

## Data Structures
        
*Generic datastructures and algorithms in Go.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[algorithms](https://github.com/shady831213/algorithms) | 154| 6| 2018-01-31 17:27:56 +0800 CST| 2019-01-08 14:12:30 +0800 CST|
|[binpacker](https://github.com/zhuangsirui/binpacker) | 110| 10| 2016-02-02 18:06:11 +0800 CST| 2018-06-17 06:52:04 +0800 CST|
|[bit](https://github.com/yourbasic/bit) | 29| 3| 2017-05-04 03:05:35 +0800 CST| 2018-03-13 15:45:26 +0800 CST|
|[bitset](https://github.com/willf/bitset) | 437| 27| 2011-05-11 11:33:44 +0800 CST| 2019-03-01 05:25:26 +0800 CST|
|[bloom](https://github.com/zentures/bloom) | 125| 8| 2013-09-03 10:27:35 +0800 CST| 2018-04-16 15:52:10 +0800 CST|
|[bloom](https://github.com/yourbasic/bloom) | 13| 1| 2017-05-07 03:57:47 +0800 CST| 2017-06-20 01:00:50 +0800 CST|
|[BoomFilters](https://github.com/tylertreat/BoomFilters) | 1080| 36| 2015-02-06 10:01:26 +0800 CST| 2018-10-29 03:28:14 +0800 CST|
|[concurrent-writer](https://github.com/free/concurrent-writer) | 15| 4| 2017-09-18 23:29:59 +0800 CST| 2017-11-18 05:28:32 +0800 CST|
|[conjungo](https://github.com/InVisionApp/conjungo) | 57| 14| 2016-12-30 07:50:38 +0800 CST| 2019-02-12 21:44:17 +0800 CST|
|[count-min-log](https://github.com/seiflotfy/count-min-log) | 40| 4| 2015-08-17 06:31:36 +0800 CST| 2017-02-12 21:09:21 +0800 CST|
|[cuckoofilter](https://github.com/seiflotfy/cuckoofilter) | 443| 17| 2015-06-29 07:22:09 +0800 CST| 2019-02-22 16:40:01 +0800 CST|
|[deque](https://github.com/edwingeng/deque) | 1| 1| 2019-02-01 11:32:28 +0800 CST| 2019-02-13 10:15:57 +0800 CST|
|[deque](https://github.com/gammazero/deque) | 32| 4| 2018-04-24 10:57:55 +0800 CST| 2019-01-31 03:14:01 +0800 CST|
|[encoding](https://github.com/zentures/encoding) | 92| 7| 2013-09-21 03:29:57 +0800 CST| 2017-12-24 06:15:28 +0800 CST|
|[go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) | 62| 5| 2016-04-01 09:40:40 +0800 CST| 2018-12-06 19:29:01 +0800 CST|
|[go-datastructures](https://github.com/Workiva/go-datastructures) | 4776| 292| 2014-10-29 21:55:17 +0800 CST| 2018-11-09 01:26:27 +0800 CST|
|[go-ef](https://github.com/amallia/go-ef) | 8| 1| 2017-09-22 09:47:16 +0800 CST| 2017-09-26 04:07:11 +0800 CST|
|[go-geoindex](https://github.com/hailocab/go-geoindex) | 297| 67| 2015-01-22 20:26:17 +0800 CST| 2018-02-21 05:58:39 +0800 CST|
|[go-mcache](https://github.com/OrlovEvgeny/go-mcache) | 23| 2| 2018-04-15 07:31:21 +0800 CST| 2018-11-17 03:36:37 +0800 CST|
|[go-rquad](https://github.com/arl/go-rquad) | 94| 3| 2016-09-13 05:46:37 +0800 CST| 2018-06-18 03:19:46 +0800 CST|
|[goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) | 2| 1| 2019-01-11 05:21:23 +0800 CST| 2019-01-31 05:44:05 +0800 CST|
|[gods](https://github.com/emirpasic/gods) | 5340| 254| 2015-03-04 22:19:52 +0800 CST| 2019-01-24 20:07:05 +0800 CST|
|[golang-set](https://github.com/deckarep/golang-set) | 974| 27| 2013-07-04 05:52:01 +0800 CST| 2018-09-27 23:06:50 +0800 CST|
|[goset](https://github.com/zoumo/goset) | 12| 1| 2017-08-25 17:21:30 +0800 CST| 2017-08-25 18:23:15 +0800 CST|
|[goskiplist](https://github.com/ryszard/goskiplist) | 178| 12| 2012-05-09 13:44:59 +0800 CST| 2017-02-24 16:36:21 +0800 CST|
|[gota](https://github.com/go-gota/gota) | 734| 57| 2016-02-07 01:23:25 +0800 CST| 2019-02-27 00:15:43 +0800 CST|
|[hide](https://github.com/emvi/hide) | 2| 2| 2019-01-16 21:54:17 +0800 CST| 2019-02-22 23:06:34 +0800 CST|
|[hilbert](https://github.com/google/hilbert) | 171| 19| 2015-08-06 23:50:00 +0800 CST| 2018-11-22 14:15:33 +0800 CST|
|[hyperloglog](https://github.com/axiomhq/hyperloglog) | 634| 18| 2017-06-18 19:18:12 +0800 CST| 2018-12-23 19:14:21 +0800 CST|
|[levenshtein](https://github.com/agext/levenshtein) | 25| 1| 2016-04-08 08:14:31 +0800 CST| 2019-02-23 03:07:15 +0800 CST|
|[levenshtein](https://github.com/agnivade/levenshtein) | 42| 3| 2014-07-30 22:03:55 +0800 CST| 2019-02-28 01:56:32 +0800 CST|
|[mafsa](https://github.com/smartystreets/mafsa) | 272| 18| 2014-11-08 05:07:57 +0800 CST| 2017-03-07 23:18:16 +0800 CST|
|[merkletree](https://github.com/cbergoon/merkletree) | 119| 4| 2017-04-12 10:50:11 +0800 CST| 2018-12-26 09:42:24 +0800 CST|
|[mspm](https://github.com/BlackRabbitt/mspm) | 4| 3| 2018-05-18 02:59:44 +0800 CST| 2018-05-19 14:36:38 +0800 CST|
|[null](https://github.com/emvi/null) | 2| 2| 2018-07-05 05:18:45 +0800 CST| 2019-02-22 23:01:00 +0800 CST|
|[pipeline](https://github.com/hyfather/pipeline) | 9| 1| 2018-04-25 08:11:36 +0800 CST| 2018-08-31 11:09:33 +0800 CST|
|[ring](https://github.com/TheTannerRyan/ring) | 83| 1| 2019-01-27 12:02:20 +0800 CST| 2019-02-19 10:52:44 +0800 CST|
|[roaring](https://github.com/RoaringBitmap/roaring) | 564| 34| 2014-07-11 04:14:34 +0800 CST| 2019-01-10 22:41:21 +0800 CST|
|[set](https://github.com/StudioSol/set) | 5| 17| 2018-07-21 05:53:37 +0800 CST| 2018-10-10 02:02:01 +0800 CST|
|[skiplist](https://github.com/MauriceGit/skiplist) | 80| 5| 2018-06-24 00:01:51 +0800 CST| 2018-12-08 17:30:32 +0800 CST|
|[skiplist](https://github.com/gansidui/skiplist) | 58| 7| 2014-11-19 00:29:53 +0800 CST| 2014-11-21 13:13:52 +0800 CST|
|[timedmap](https://github.com/zekroTJA/timedmap) | 0| 1| 2019-01-30 20:55:37 +0800 CST| 2019-02-20 22:48:16 +0800 CST|
|[trie](https://github.com/derekparker/trie) | 356| 20| 2014-03-07 06:01:49 +0800 CST| 2018-05-11 06:29:42 +0800 CST|
|[ttlcache](https://github.com/ReneKroon/ttlcache) | 78| 7| 2014-12-13 09:55:40 +0800 CST| 2019-02-25 21:04:39 +0800 CST|
|[bloom](https://github.com/willf/bloom) | 584| 30| 2011-05-21 22:18:41 +0800 CST| 2019-03-01 05:23:55 +0800 CST|

## Database
        
*SQL query builder, libraries for building and using SQL.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[badger](https://github.com/dgraph-io/badger) | 5374| 219| 2017-01-26 13:09:49 +0800 CST| 2019-03-02 03:36:11 +0800 CST|
|[bigcache](https://github.com/allegro/bigcache) | 1878| 64| 2016-03-23 15:18:52 +0800 CST| 2019-02-18 14:46:06 +0800 CST|
|[bolt](https://github.com/boltdb/bolt) | 9486| 337| 2013-12-21 02:26:14 +0800 CST| 2018-03-03 02:00:53 +0800 CST|
|[buntdb](https://github.com/tidwall/buntdb) | 2262| 93| 2016-07-20 06:11:40 +0800 CST| 2019-02-03 03:56:57 +0800 CST|
|[cache2go](https://github.com/muesli/cache2go) | 811| 57| 2013-11-11 11:45:02 +0800 CST| 2018-10-29 06:37:51 +0800 CST|
|[clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) | 25| 5| 2017-12-18 15:48:07 +0800 CST| 2018-01-23 06:02:54 +0800 CST|
|[cockroach](https://github.com/cockroachdb/cockroach) | 15515| 693| 2014-02-06 08:18:47 +0800 CST| 2019-03-02 13:28:37 +0800 CST|
|[couchcache](https://github.com/codingsince1985/couchcache) | 39| 3| 2015-04-05 15:13:05 +0800 CST| 2017-10-25 09:28:55 +0800 CST|
|[CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) | 481| 45| 2018-04-11 17:52:58 +0800 CST| 2019-03-01 17:54:27 +0800 CST|
|[dgraph](https://github.com/dgraph-io/dgraph) | 8682| 329| 2015-08-25 15:15:56 +0800 CST| 2019-03-02 12:19:32 +0800 CST|
|[diskv](https://github.com/peterbourgon/diskv) | 695| 29| 2012-03-22 00:44:32 +0800 CST| 2018-09-04 20:15:47 +0800 CST|
|[eliasdb](https://github.com/krotik/eliasdb) | 503| 23| 2016-08-13 21:53:28 +0800 CST| 2019-02-16 01:18:02 +0800 CST|
|[fastcache](https://github.com/VictoriaMetrics/fastcache) | 347| 12| 2018-11-23 06:50:13 +0800 CST| 2019-02-25 05:17:14 +0800 CST|
|[gcache](https://github.com/bluele/gcache) | 664| 29| 2015-01-25 02:17:07 +0800 CST| 2019-03-01 12:41:16 +0800 CST|
|[go-cache](https://github.com/patrickmn/go-cache) | 2415| 84| 2012-01-02 21:07:13 +0800 CST| 2019-02-27 09:24:53 +0800 CST|
|[goleveldb](https://github.com/syndtr/goleveldb) | 2787| 150| 2013-01-23 12:08:58 +0800 CST| 2019-02-26 23:50:50 +0800 CST|
|[gorocksdb](https://github.com/kapitan-k/gorocksdb) | 6| 1| 2017-12-28 18:28:48 +0800 CST| 2018-01-11 04:23:36 +0800 CST|
|[groupcache](https://github.com/golang/groupcache) | 7094| 443| 2013-07-23 05:55:07 +0800 CST| 2019-02-13 09:06:53 +0800 CST|
|[influxdb](https://github.com/influxdata/influxdb) | 15690| 750| 2013-09-26 22:31:10 +0800 CST| 2019-03-02 11:25:15 +0800 CST|
|[ledisdb](https://github.com/siddontang/ledisdb) | 2886| 184| 2014-04-30 08:43:09 +0800 CST| 2019-02-05 13:49:34 +0800 CST|
|[levigo](https://github.com/jmhodges/levigo) | 345| 24| 2012-01-17 16:17:54 +0800 CST| 2019-02-28 18:43:15 +0800 CST|
|[moss](https://github.com/couchbase/moss) | 682| 79| 2016-02-07 04:27:22 +0800 CST| 2019-02-22 00:32:47 +0800 CST|
|[piladb](https://github.com/fern4lvarez/piladb) | 167| 8| 2015-09-09 07:12:22 +0800 CST| 2018-04-08 00:09:23 +0800 CST|
|[prometheus](https://github.com/prometheus/prometheus) | 22352| 1004| 2012-11-24 19:14:12 +0800 CST| 2019-03-02 11:28:07 +0800 CST|
|[pudge](https://github.com/recoilme/pudge) | 181| 6| 2018-11-20 18:11:53 +0800 CST| 2019-02-27 15:03:55 +0800 CST|
|[rqlite](https://github.com/rqlite/rqlite) | 4253| 182| 2014-08-23 12:31:18 +0800 CST| 2019-02-27 01:18:36 +0800 CST|
|[golang-scribble](https://github.com/nanobox-io/golang-scribble) | 32| 3| 2018-06-22 06:13:33 +0800 CST| 2018-10-18 23:48:25 +0800 CST|
|[slowpoke](https://github.com/recoilme/slowpoke) | 74| 6| 2018-02-19 17:22:37 +0800 CST| 2018-12-24 23:55:56 +0800 CST|
|[tempdb](https://github.com/rafaeljesus/tempdb) | 13| 3| 2017-03-18 02:03:42 +0800 CST| 2018-02-15 03:03:13 +0800 CST|
|[tidb](https://github.com/pingcap/tidb) | 17531| 1198| 2015-09-06 12:01:52 +0800 CST| 2019-03-02 13:16:30 +0800 CST|
|[tiedot](https://github.com/HouzuoGuo/tiedot) | 2306| 161| 2013-05-26 18:03:49 +0800 CST| 2019-01-26 01:43:32 +0800 CST|
|[vasto](https://github.com/chrislusf/vasto) | 120| 16| 2018-01-16 13:16:57 +0800 CST| 2018-10-23 17:02:13 +0800 CST|
|[darwin](https://github.com/GuiaBolso/darwin) | 77| 2| 2016-04-05 23:57:59 +0800 CST| 2018-11-28 00:12:25 +0800 CST|
|[go-fixtures](https://github.com/RichardKnop/go-fixtures) | 17| 1| 2015-12-24 19:27:45 +0800 CST| 2018-11-01 11:56:57 +0800 CST|
|[go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) | 19| 1| 2018-08-11 15:00:13 +0800 CST| 2018-12-23 00:30:09 +0800 CST|
|[goose](https://github.com/steinbacher/goose) | 103| 4| 2016-03-05 04:33:14 +0800 CST| 2016-09-24 15:03:49 +0800 CST|
|[gormigrate](https://github.com/go-gormigrate/gormigrate) | 256| 3| 2016-08-31 19:46:23 +0800 CST| 2019-02-04 03:31:20 +0800 CST|
|[migrate](https://github.com/golang-migrate/migrate) | 1690| 37| 2018-01-19 17:30:58 +0800 CST| 2019-03-02 10:52:21 +0800 CST|
|[pravasan](https://github.com/pravasan/pravasan) | 23| 6| 2015-01-04 01:11:05 +0800 CST| 2018-12-20 09:56:10 +0800 CST|
|[pop](https://github.com/gobuffalo/pop) | 585| 24| 2018-02-08 05:13:46 +0800 CST| 2019-03-01 00:49:58 +0800 CST|
|[sql-migrate](https://github.com/rubenv/sql-migrate) | 1243| 29| 2014-09-09 15:31:41 +0800 CST| 2019-02-15 05:27:34 +0800 CST|
|[chproxy](https://github.com/Vertamedia/chproxy) | 240| 20| 2017-09-18 21:09:23 +0800 CST| 2019-02-03 05:00:26 +0800 CST|
|[clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) | 97| 6| 2017-04-29 18:38:41 +0800 CST| 2019-02-20 03:05:55 +0800 CST|
|[dbbench](https://github.com/sj14/dbbench) | 23| 2| 2018-11-24 21:21:18 +0800 CST| 2019-02-02 16:57:42 +0800 CST|
|[go-mysql](https://github.com/siddontang/go-mysql) | 1564| 133| 2014-02-21 09:56:45 +0800 CST| 2019-02-28 17:00:56 +0800 CST|
|[go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) | 2032| 150| 2015-01-15 17:54:18 +0800 CST| 2019-02-27 10:54:41 +0800 CST|
|[kingshard](https://github.com/flike/kingshard) | 4179| 388| 2015-07-04 10:22:32 +0800 CST| 2019-02-19 18:33:01 +0800 CST|
|[myreplication](https://github.com/2tvenom/myreplication) | 129| 17| 2015-02-05 04:59:49 +0800 CST| 2018-10-05 15:34:57 +0800 CST|
|[octillery](https://github.com/knocknote/octillery) | 42| 14| 2018-11-26 18:39:35 +0800 CST| 2019-01-18 17:14:32 +0800 CST|
|[orchestrator](https://github.com/github/orchestrator) | 2588| 181| 2016-11-30 21:44:24 +0800 CST| 2019-03-01 18:17:54 +0800 CST|
|[pgweb](https://github.com/sosedoff/pgweb) | 5750| 152| 2014-10-09 09:41:32 +0800 CST| 2019-02-26 19:14:40 +0800 CST|
|[prep](https://github.com/hexdigest/prep) | 22| 2| 2017-12-12 07:47:38 +0800 CST| 2017-12-20 01:35:51 +0800 CST|
|[prest](https://github.com/prest/prest) | 1963| 82| 2016-11-22 13:17:05 +0800 CST| 2018-09-04 21:21:35 +0800 CST|
|[rwdb](https://github.com/andizzle/rwdb) | 10| 2| 2017-10-04 11:55:29 +0800 CST| 2017-11-08 17:10:17 +0800 CST|
|[vitess](https://github.com/vitessio/vitess) | 7600| 464| 2013-06-28 05:20:28 +0800 CST| 2019-03-02 11:09:12 +0800 CST|
|[dotsql](https://github.com/gchaincl/dotsql) | 406| 21| 2014-11-20 20:14:39 +0800 CST| 2018-10-22 14:31:23 +0800 CST|
|[gendry](https://github.com/didi/gendry) | 570| 46| 2017-12-01 16:15:43 +0800 CST| 2018-12-29 12:14:00 +0800 CST|
|[godbal](https://github.com/xujiajun/godbal) | 48| 3| 2018-02-28 13:47:42 +0800 CST| 2019-01-30 13:57:00 +0800 CST|
|[goqu](https://github.com/doug-martin/goqu) | 491| 25| 2015-02-21 09:06:00 +0800 CST| 2019-02-13 23:37:43 +0800 CST|
|[igor](https://github.com/galeone/igor) | 73| 6| 2016-03-10 22:45:08 +0800 CST| 2018-07-01 17:41:38 +0800 CST|
|[ormlite](https://github.com/pupizoid/ormlite) | 0| 1| 2018-06-28 21:42:09 +0800 CST| 2019-02-15 22:37:03 +0800 CST|
|[ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) | 296| 25| 2015-12-11 06:39:26 +0800 CST| 2018-08-23 14:45:48 +0800 CST|
|[scaneo](https://github.com/variadico/scaneo) | 141| 8| 2015-04-30 08:36:27 +0800 CST| 2018-04-03 20:49:37 +0800 CST|
|[sqrl](https://github.com/elgris/sqrl) | 140| 5| 2014-06-25 18:03:06 +0800 CST| 2019-01-30 01:38:07 +0800 CST|
|[squirrel](https://github.com/Masterminds/squirrel) | 1930| 32| 2014-01-18 13:29:58 +0800 CST| 2019-02-27 02:57:20 +0800 CST|
|[xo](https://github.com/xo/xo) | 1944| 70| 2016-02-05 18:22:20 +0800 CST| 2019-01-20 20:32:12 +0800 CST|

## Database Drivers
        
*Libraries for connecting and operating databases.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|

### Relational Databases
        
|[calcite-avatica-go](https://github.com/apache/calcite-avatica-go) | 22| 16| 2017-08-08 15:00:08 +0800 CST| 2019-01-08 07:50:24 +0800 CST|
|[bgc](https://github.com/viant/bgc) | 10| 9| 2016-06-14 04:24:26 +0800 CST| 2019-01-30 03:42:30 +0800 CST|
|[firebirdsql](https://github.com/nakagami/firebirdsql) | 94| 14| 2013-08-27 21:09:14 +0800 CST| 2019-03-02 10:01:54 +0800 CST|
|[go-adodb](https://github.com/mattn/go-adodb) | 84| 12| 2011-11-14 12:32:50 +0800 CST| 2019-02-26 18:19:11 +0800 CST|
|[go-mssqldb](https://github.com/denisenkom/go-mssqldb) | 908| 59| 2013-12-16 08:10:47 +0800 CST| 2019-02-19 13:38:43 +0800 CST|
|[go-oci8](https://github.com/mattn/go-oci8) | 360| 34| 2012-02-29 20:19:16 +0800 CST| 2019-02-05 18:29:31 +0800 CST|
|[mysql](https://github.com/go-sql-driver/mysql) | 7034| 391| 2012-12-10 04:33:55 +0800 CST| 2019-02-17 15:26:59 +0800 CST|
|[go-sqlite3](https://github.com/mattn/go-sqlite3) | 3084| 125| 2011-11-11 20:36:50 +0800 CST| 2019-02-18 01:40:34 +0800 CST|
|[gofreetds](https://github.com/minus5/gofreetds) | 84| 21| 2012-12-07 01:29:26 +0800 CST| 2019-02-20 00:37:09 +0800 CST|
|[goracle](https://github.com/go-goracle/goracle) | 182| 28| 2015-03-25 13:58:16 +0800 CST| 2019-02-20 13:20:21 +0800 CST|
|[pgx](https://github.com/jackc/pgx) | 1645| 64| 2013-03-31 03:06:26 +0800 CST| 2019-03-01 00:05:55 +0800 CST|
|[pq](https://github.com/lib/pq) | 4669| 158| 2012-03-13 02:50:22 +0800 CST| 2019-02-09 01:37:56 +0800 CST|

### NoSQL Databases
        
|[aerospike-client-go](https://github.com/aerospike/aerospike-client-go) | 288| 39| 2014-07-26 10:56:21 +0800 CST| 2019-02-21 21:56:25 +0800 CST|
|[arangolite](https://github.com/solher/arangolite) | 64| 5| 2015-10-05 01:27:34 +0800 CST| 2017-09-10 17:40:18 +0800 CST|
|[asc](https://github.com/viant/asc) | 4| 10| 2016-06-14 04:22:31 +0800 CST| 2018-11-13 12:21:38 +0800 CST|
|[dynago](https://github.com/underarmour/dynago) | 63| 131| 2015-05-18 23:40:20 +0800 CST| 2017-08-08 06:07:05 +0800 CST|
|[goforestdb](https://github.com/couchbase/goforestdb) | 29| 44| 2014-05-14 23:36:12 +0800 CST| 2016-12-16 06:01:01 +0800 CST|
|[go-couchbase](https://github.com/couchbase/go-couchbase) | 285| 26| 2012-01-20 06:52:08 +0800 CST| 2019-03-02 05:31:26 +0800 CST|
|[go-couchdb](https://github.com/fjl/go-couchdb) | 51| 6| 2013-10-28 09:08:16 +0800 CST| 2018-09-05 10:07:28 +0800 CST|
|[go-pilosa](https://github.com/pilosa/go-pilosa) | 27| 13| 2016-10-01 05:37:10 +0800 CST| 2019-03-01 23:01:52 +0800 CST|
|[go-rejson](https://github.com/nitishm/go-rejson) | 67| 4| 2018-04-23 08:51:05 +0800 CST| 2019-01-28 01:33:17 +0800 CST|
|[gocb](https://github.com/couchbase/gocb) | 281| 71| 2015-01-16 04:01:32 +0800 CST| 2019-02-19 12:00:42 +0800 CST|
|[godscache](https://github.com/defcronyke/godscache) | 4| 2| 2018-05-09 04:19:39 +0800 CST| 2019-02-08 15:04:54 +0800 CST|
|[gomemcache](https://github.com/bradfitz/gomemcache) | 1011| 51| 2011-06-29 03:29:12 +0800 CST| 2019-02-22 05:56:08 +0800 CST|
|[rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1418| 45| 2013-09-12 21:56:27 +0800 CST| 2019-01-22 02:10:18 +0800 CST|
|[goriak](https://github.com/zegl/goriak) | 25| 2| 2016-10-06 00:48:17 +0800 CST| 2018-12-28 07:06:55 +0800 CST|
|[mgo](https://github.com/globalsign/mgo) | 1423| 77| 2017-04-13 19:14:04 +0800 CST| 2019-03-01 00:42:21 +0800 CST|
|[mongo-go-driver](https://github.com/mongodb/mongo-go-driver) | 2019| 136| 2017-02-09 01:18:02 +0800 CST| 2019-03-01 06:50:01 +0800 CST|
|[neo4j](https://github.com/cihangir/neo4j) | 24| 3| 2013-05-18 16:54:01 +0800 CST| 2015-04-03 01:38:48 +0800 CST|
|[Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) | 72| 6| 2011-06-05 00:08:35 +0800 CST| 2018-06-20 20:15:38 +0800 CST|
|[neoism](https://github.com/jmcvetta/neoism) | 350| 23| 2012-07-12 15:42:33 +0800 CST| 2019-02-19 09:47:26 +0800 CST|
|[redigo](https://github.com/gomodule/redigo) | 5538| 280| 2012-04-14 12:31:58 +0800 CST| 2019-03-01 19:22:27 +0800 CST|
|[redis](https://github.com/go-redis/redis) | 5135| 200| 2012-07-25 21:01:39 +0800 CST| 2019-03-01 17:23:50 +0800 CST|
|[redeo](https://github.com/bsm/redeo) | 240| 24| 2014-03-06 16:46:18 +0800 CST| 2018-11-29 20:16:24 +0800 CST|
|[xredis](https://github.com/shomali11/xredis) | 9| 1| 2017-06-14 08:19:26 +0800 CST| 2018-06-07 08:59:07 +0800 CST|
|[bleve](https://github.com/blevesearch/bleve) | 5000| 231| 2014-04-18 05:02:18 +0800 CST| 2019-02-28 00:32:59 +0800 CST|
|[elastic](https://github.com/olivere/elastic) | 3507| 160| 2012-12-07 01:15:33 +0800 CST| 2019-02-26 18:19:22 +0800 CST|
|[elasticsql](https://github.com/cch123/elasticsql) | 291| 20| 2016-08-24 15:29:43 +0800 CST| 2019-02-22 15:10:51 +0800 CST|
|[elastigo](https://github.com/mattbaird/elastigo) | 937| 49| 2012-10-12 12:19:59 +0800 CST| 2019-02-06 02:17:02 +0800 CST|
|[goes](https://github.com/OwnLocal/goes) | 25| 34| 2015-12-29 02:52:03 +0800 CST| 2017-03-03 09:06:24 +0800 CST|
|[riot](https://github.com/go-ego/riot) | 4295| 171| 2017-06-21 22:17:59 +0800 CST| 2019-01-31 22:23:29 +0800 CST|
|[skizze](https://github.com/seiflotfy/skizze) | 59| 6| 2016-01-17 20:10:40 +0800 CST| 2016-05-10 02:15:30 +0800 CST|
|[cachego](https://github.com/fabiorphp/cachego) | 103| 6| 2016-10-06 02:10:03 +0800 CST| 2018-05-20 20:11:29 +0800 CST|
|[cayley](https://github.com/cayleygraph/cayley) | 12246| 632| 2014-06-06 02:49:41 +0800 CST| 2019-02-22 04:17:08 +0800 CST|
|[dsc](https://github.com/viant/dsc) | 6| 17| 2016-06-14 04:18:10 +0800 CST| 2019-02-24 07:34:49 +0800 CST|
|[gokv](https://github.com/philippgille/gokv) | 44| 4| 2018-10-09 02:55:22 +0800 CST| 2019-02-25 03:40:57 +0800 CST|

## Date and Time
        
*Libraries for working with dates and times.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[carbon](https://github.com/uniplaces/carbon) | 305| 38| 2016-08-03 18:55:52 +0800 CST| 2019-01-16 20:02:42 +0800 CST|
|[date](https://github.com/rickb777/date) | 23| 2| 2015-11-24 06:58:07 +0800 CST| 2018-11-27 22:32:14 +0800 CST|
|[dateparse](https://github.com/araddon/dateparse) | 800| 14| 2014-04-21 10:55:48 +0800 CST| 2019-02-23 09:01:43 +0800 CST|
|[durafmt](https://github.com/hako/durafmt) | 208| 4| 2016-05-21 05:49:33 +0800 CST| 2018-09-04 07:44:38 +0800 CST|
|[feiertage](https://github.com/wlbr/feiertage) | 17| 1| 2015-11-04 22:19:27 +0800 CST| 2019-02-14 08:19:35 +0800 CST|
|[go-persian-calendar](https://github.com/yaa110/go-persian-calendar) | 48| 3| 2016-02-01 02:40:23 +0800 CST| 2019-01-28 20:52:10 +0800 CST|
|[go-sunrise](https://github.com/nathan-osman/go-sunrise) | 10| 3| 2017-06-16 04:49:41 +0800 CST| 2017-11-22 04:50:07 +0800 CST|
|[goweek](https://github.com/grsmv/goweek) | 18| 3| 2015-11-12 22:11:46 +0800 CST| 2017-01-04 04:24:26 +0800 CST|
|[iso8601](https://github.com/relvacode/iso8601) | 60| 2| 2017-04-25 23:54:18 +0800 CST| 2018-12-21 23:13:36 +0800 CST|
|[kair](https://github.com/GuilhermeCaruso/kair) | 9| 1| 2018-10-03 23:44:07 +0800 CST| 2019-02-24 06:06:10 +0800 CST|
|[now](https://github.com/jinzhu/now) | 1978| 51| 2013-11-18 18:55:30 +0800 CST| 2019-02-17 08:24:44 +0800 CST|
|[nulltime](https://github.com/kirillDanshin/nulltime) | 9| 1| 2016-03-06 09:44:58 +0800 CST| 2017-03-22 12:30:28 +0800 CST|
|[strftime](https://github.com/awoodbeck/strftime) | 5| 1| 2018-02-10 08:35:46 +0800 CST| 2018-02-21 23:59:14 +0800 CST|
|[timespan](https://github.com/SaidinWoT/timespan) | 60| 5| 2014-10-07 11:57:02 +0800 CST| 2019-02-01 16:55:52 +0800 CST|
|[timeutil](https://github.com/leekchan/timeutil) | 156| 6| 2015-08-02 09:32:06 +0800 CST| 2019-02-03 21:14:43 +0800 CST|
|[tuesday](https://github.com/osteele/tuesday) | 7| 1| 2017-08-11 04:46:26 +0800 CST| 2017-08-30 21:28:26 +0800 CST|

## Distributed Systems
        
*Packages that help with building Distributed Systems.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[celeriac.v1](https://github.com/svcavallar/celeriac.v1) | 47| 3| 2015-10-10 15:27:33 +0800 CST| 2018-12-08 13:38:42 +0800 CST|
|[consistent](https://github.com/buraksezer/consistent) | 151| 9| 2018-03-25 23:38:27 +0800 CST| 2019-02-23 11:07:49 +0800 CST|
|[digota](https://github.com/digota/digota) | 269| 24| 2017-08-14 20:01:37 +0800 CST| 2018-10-15 06:52:00 +0800 CST|
|[doublejump](https://github.com/edwingeng/doublejump) | 26| 3| 2018-06-27 00:04:50 +0800 CST| 2019-01-02 18:37:11 +0800 CST|
|[drmaa](https://github.com/dgruber/drmaa) | 23| 2| 2013-03-17 20:58:02 +0800 CST| 2018-05-14 14:05:07 +0800 CST|
|[dynatomic](https://github.com/tylfin/dynatomic) | 5| 0| 2019-02-09 01:45:14 +0800 CST| 2019-02-21 20:40:01 +0800 CST|
|[emitter](https://github.com/emitter-io/emitter) | 1651| 80| 2016-10-29 16:52:21 +0800 CST| 2019-03-02 10:58:02 +0800 CST|
|[flowgraph](https://github.com/vectaport/flowgraph) | 7| 1| 2018-08-30 05:45:26 +0800 CST| 2019-03-02 02:28:02 +0800 CST|
|[gleam](https://github.com/chrislusf/gleam) | 1829| 137| 2016-08-26 16:44:48 +0800 CST| 2018-12-11 22:58:01 +0800 CST|
|[glow](https://github.com/chrislusf/glow) | 2351| 141| 2015-06-14 08:33:48 +0800 CST| 2018-11-02 14:09:14 +0800 CST|
|[go-health](https://github.com/InVisionApp/go-health) | 437| 20| 2017-11-30 05:00:07 +0800 CST| 2018-12-05 02:25:05 +0800 CST|
|[go-jump](https://github.com/dgryski/go-jump) | 231| 13| 2014-06-16 06:12:04 +0800 CST| 2018-02-12 22:36:50 +0800 CST|
|[kit](https://github.com/go-kit/kit) | 12764| 641| 2015-02-03 08:01:19 +0800 CST| 2019-02-25 09:14:30 +0800 CST|
|[gorpc](https://github.com/valyala/gorpc) | 520| 26| 2014-11-21 01:02:37 +0800 CST| 2017-04-07 09:55:22 +0800 CST|
|[grpc-go](https://github.com/grpc/grpc-go) | 7575| 437| 2014-12-09 02:59:34 +0800 CST| 2019-03-02 08:11:57 +0800 CST|
|[hprose-golang](https://github.com/hprose/hprose-golang) | 918| 89| 2014-02-14 11:16:43 +0800 CST| 2018-05-17 23:21:26 +0800 CST|
|[jaeger](https://github.com/jaegertracing/jaeger) | 7328| 277| 2016-04-16 02:49:02 +0800 CST| 2019-03-02 07:10:15 +0800 CST|
|[jsonrpc](https://github.com/osamingo/jsonrpc) | 103| 6| 2016-10-28 21:36:59 +0800 CST| 2019-01-04 00:39:07 +0800 CST|
|[jsonrpc](https://github.com/ybbus/jsonrpc) | 78| 6| 2016-11-10 19:27:55 +0800 CST| 2018-12-01 14:59:14 +0800 CST|
|[krakend](https://github.com/devopsfaith/krakend) | 1273| 67| 2016-11-05 02:37:13 +0800 CST| 2019-03-01 22:59:56 +0800 CST|
|[micro](https://github.com/micro/micro) | 5582| 289| 2015-01-17 06:35:14 +0800 CST| 2019-02-28 19:16:38 +0800 CST|
|[gnatsd](https://github.com/nats-io/gnatsd) | 5307| 326| 2012-10-30 00:12:24 +0800 CST| 2019-03-01 05:14:14 +0800 CST|
|[outboxer](https://github.com/italolelis/outboxer) | 0| 0| 2019-02-01 17:50:13 +0800 CST| 2019-02-22 05:34:43 +0800 CST|
|[raft](https://github.com/hashicorp/raft) | 2483| 176| 2013-11-05 08:41:20 +0800 CST| 2019-03-01 20:30:53 +0800 CST|
|[etcd](https://github.com/etcd-io/etcd) | 23220| 1226| 2013-07-07 05:57:21 +0800 CST| 2019-03-02 10:04:53 +0800 CST|
|[redis-lock](https://github.com/bsm/redis-lock) | 108| 6| 2014-10-10 22:22:48 +0800 CST| 2018-11-29 19:41:28 +0800 CST|
|[ringpop-go](https://github.com/uber/ringpop-go) | 530| 2310| 2015-06-06 06:48:53 +0800 CST| 2019-02-19 11:59:14 +0800 CST|
|[rpcx](https://github.com/smallnest/rpcx) | 3233| 270| 2016-05-18 17:34:05 +0800 CST| 2019-03-01 17:21:34 +0800 CST|
|[sleuth](https://github.com/ursiform/sleuth) | 292| 7| 2016-04-23 22:21:41 +0800 CST| 2018-03-21 23:59:30 +0800 CST|
|[tendermint](https://github.com/tendermint/tendermint) | 2710| 241| 2014-05-15 07:21:35 +0800 CST| 2019-03-02 14:06:58 +0800 CST|
|[torrent](https://github.com/anacrolix/torrent) | 2608| 121| 2015-01-09 05:10:42 +0800 CST| 2019-03-01 08:18:58 +0800 CST|
|[go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) | 355| 18| 2015-10-09 03:44:47 +0800 CST| 2019-02-23 23:51:07 +0800 CST|

## Email
        
*Libraries and tools that implement email creation and sending.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[douceur](https://github.com/aymerick/douceur) | 141| 4| 2015-04-09 18:21:26 +0800 CST| 2018-03-23 05:03:08 +0800 CST|
|[email](https://github.com/jordan-wright/email) | 1007| 35| 2013-12-13 04:11:59 +0800 CST| 2019-02-18 10:44:55 +0800 CST|
|[go-dkim](https://github.com/toorop/go-dkim) | 47| 3| 2015-04-29 23:38:27 +0800 CST| 2018-02-24 05:40:33 +0800 CST|
|[go-imap](https://github.com/emersion/go-imap) | 625| 31| 2016-04-27 01:59:18 +0800 CST| 2019-03-02 03:04:29 +0800 CST|
|[go-message](https://github.com/emersion/go-message) | 79| 6| 2016-12-31 17:31:52 +0800 CST| 2019-01-09 06:54:23 +0800 CST|
|[gomail](https://github.com/go-gomail/gomail) | 2134| 65| 2014-10-15 23:47:48 +0800 CST| 2019-02-13 09:16:07 +0800 CST|
|[hectane](https://github.com/hectane/hectane) | 160| 12| 2015-08-28 09:36:47 +0800 CST| 2018-04-15 04:06:48 +0800 CST|
|[hermes](https://github.com/matcornic/hermes) | 1474| 23| 2017-03-26 02:25:36 +0800 CST| 2019-02-27 02:41:35 +0800 CST|
|[MailHog](https://github.com/mailhog/MailHog) | 4495| 125| 2014-04-17 06:28:49 +0800 CST| 2019-03-02 01:48:44 +0800 CST|
|[sendgrid-go](https://github.com/sendgrid/sendgrid-go) | 464| 181| 2013-09-12 11:31:13 +0800 CST| 2019-02-23 06:04:09 +0800 CST|
|[smtp](https://github.com/mailhog/smtp) | 49| 9| 2014-12-25 00:13:49 +0800 CST| 2018-06-13 00:19:22 +0800 CST|

## Embeddable Scripting Languages
        
*Embedding other languages inside your go code.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[agora](https://github.com/mna/agora) | 318| 22| 2013-06-19 10:16:30 +0800 CST| 2015-01-26 01:12:45 +0800 CST|
|[anko](https://github.com/mattn/anko) | 840| 47| 2014-03-28 15:29:40 +0800 CST| 2019-02-26 19:35:12 +0800 CST|
|[binder](https://github.com/alexeyco/binder) | 24| 2| 2017-04-03 01:14:52 +0800 CST| 2018-07-30 06:00:27 +0800 CST|
|[expr](https://github.com/antonmedv/expr) | 295| 16| 2018-07-14 23:57:34 +0800 CST| 2019-01-19 17:14:27 +0800 CST|
|[gisp](https://github.com/jcla1/gisp) | 415| 20| 2014-01-11 22:05:43 +0800 CST| 2017-08-25 21:48:45 +0800 CST|
|[go-duktape](https://github.com/olebedev/go-duktape) | 628| 23| 2015-01-08 13:09:05 +0800 CST| 2019-02-14 07:42:58 +0800 CST|
|[go-lua](https://github.com/Shopify/go-lua) | 1539| 253| 2013-12-21 01:29:43 +0800 CST| 2018-11-07 02:40:36 +0800 CST|
|[go-php](https://github.com/deuill/go-php) | 608| 39| 2015-09-18 05:23:52 +0800 CST| 2018-10-07 23:22:34 +0800 CST|
|[go-python](https://github.com/sbinet/go-python) | 807| 43| 2012-07-09 23:43:31 +0800 CST| 2019-01-18 21:19:05 +0800 CST|
|[golua](https://github.com/aarzilli/golua) | 422| 35| 2010-12-07 05:39:53 +0800 CST| 2018-08-22 18:36:17 +0800 CST|
|[gopher-lua](https://github.com/yuin/gopher-lua) | 2604| 136| 2015-02-15 21:23:37 +0800 CST| 2019-02-07 22:37:53 +0800 CST|
|[gval](https://github.com/PaesslerAG/gval) | 85| 12| 2017-09-27 16:32:49 +0800 CST| 2019-02-19 20:43:41 +0800 CST|
|[ngaro](https://github.com/db47h/ngaro) | 16| 1| 2016-08-09 23:23:50 +0800 CST| 2018-06-03 18:57:43 +0800 CST|
|[otto](https://github.com/robertkrimen/otto) | 4403| 187| 2012-10-06 09:48:39 +0800 CST| 2019-02-25 15:56:09 +0800 CST|
|[purl](https://github.com/ian-kent/purl) | 25| 2| 2014-11-30 03:06:01 +0800 CST| 2014-12-08 01:45:34 +0800 CST|
|[tengo](https://github.com/d5/tengo) | 1051| 24| 2019-01-09 15:17:17 +0800 CST| 2019-03-02 07:56:04 +0800 CST|

## Error Handling
        
*Libraries for handling errors.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[errors](https://github.com/pkg/errors) | 4159| 91| 2015-12-27 20:05:38 +0800 CST| 2019-02-27 08:00:54 +0800 CST|
|[errorx](https://github.com/joomcode/errorx) | 508| 37| 2018-08-17 16:02:10 +0800 CST| 2019-02-15 20:27:52 +0800 CST|
|[go-multierror](https://github.com/hashicorp/go-multierror) | 558| 81| 2014-12-16 04:12:26 +0800 CST| 2018-12-14 18:18:49 +0800 CST|
|[tracerr](https://github.com/ztrue/tracerr) | 172| 6| 2019-02-07 02:57:46 +0800 CST| 2019-02-16 18:32:28 +0800 CST|
|[werr](https://github.com/txgruppi/werr) | 10| 0| 2015-08-04 19:10:13 +0800 CST| 2016-03-10 11:37:06 +0800 CST|

## Files
        
*Libraries for  handling files and file systems.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[afero](https://github.com/spf13/afero) | 1935| 83| 2014-10-28 22:19:05 +0800 CST| 2019-03-02 10:20:19 +0800 CST|
|[go-csv-tag](https://github.com/artonge/go-csv-tag) | 39| 1| 2017-06-18 23:31:16 +0800 CST| 2018-10-15 18:32:17 +0800 CST|
|[go-decent-copy](https://github.com/hugocarreira/go-decent-copy) | 11| 1| 2018-10-16 15:08:24 +0800 CST| 2018-10-18 19:24:25 +0800 CST|
|[go-gtfs](https://github.com/artonge/go-gtfs) | 11| 2| 2017-07-09 17:30:31 +0800 CST| 2019-01-10 21:24:12 +0800 CST|
|[notify](https://github.com/rjeczalik/notify) | 450| 26| 2014-09-09 00:09:34 +0800 CST| 2019-02-19 09:29:26 +0800 CST|
|[opc](https://github.com/qmuntal/opc) | 47| 1| 2018-11-06 22:49:06 +0800 CST| 2019-02-26 14:05:35 +0800 CST|
|[pdfcpu](https://github.com/hhrutter/pdfcpu) | 796| 22| 2017-06-19 01:27:38 +0800 CST| 2019-02-25 04:42:08 +0800 CST|
|[skywalker](https://github.com/dixonwille/skywalker) | 41| 3| 2017-08-02 04:08:25 +0800 CST| 2017-08-05 04:28:55 +0800 CST|
|[tarfs](https://github.com/posener/tarfs) | 31| 2| 2017-03-11 06:13:11 +0800 CST| 2017-04-02 16:33:57 +0800 CST|

## Financial
        
*Packages for accounting and finance.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[accounting](https://github.com/leekchan/accounting) | 441| 12| 2015-08-10 21:23:56 +0800 CST| 2019-01-31 14:20:31 +0800 CST|
|[decimal](https://github.com/shopspring/decimal) | 1332| 55| 2015-02-26 04:12:57 +0800 CST| 2019-02-19 17:12:29 +0800 CST|
|[go-finance](https://github.com/FlashBoys/go-finance) | 538| 28| 2016-02-28 08:37:46 +0800 CST| 2018-03-09 10:50:46 +0800 CST|
|[go-finance](https://github.com/alpeb/go-finance) | 32| 2| 2017-06-01 23:58:33 +0800 CST| 2018-05-06 00:52:06 +0800 CST|
|[go-money](https://github.com/Rhymond/go-money) | 551| 15| 2017-03-21 00:23:54 +0800 CST| 2019-01-30 18:29:45 +0800 CST|
|[ofxgo](https://github.com/aclindsa/ofxgo) | 53| 5| 2015-11-08 21:56:53 +0800 CST| 2019-03-02 11:43:10 +0800 CST|
|[orderbook](https://github.com/i25959341/orderbook) | 21| 3| 2018-04-25 02:05:26 +0800 CST| 2019-03-01 18:49:58 +0800 CST|
|[techan](https://github.com/sdcoffey/techan) | 108| 17| 2017-03-08 11:04:08 +0800 CST| 2018-12-13 23:35:39 +0800 CST|
|[transaction](https://github.com/claygod/transaction) | 44| 8| 2017-10-11 21:50:30 +0800 CST| 2018-09-06 02:12:01 +0800 CST|
|[vat](https://github.com/dannyvankooten/vat) | 53| 1| 2016-06-19 00:10:09 +0800 CST| 2018-09-10 22:01:40 +0800 CST|

## Forms
        
*Libraries for working with forms.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[bind](https://github.com/robfig/bind) | 23| 2| 2014-08-06 08:13:10 +0800 CST| 2014-08-17 01:03:51 +0800 CST|
|[binding](https://github.com/mholt/binding) | 730| 30| 2014-05-21 07:35:14 +0800 CST| 2018-03-29 07:47:34 +0800 CST|
|[conform](https://github.com/leebenson/conform) | 163| 5| 2016-01-06 02:00:06 +0800 CST| 2018-06-16 05:02:23 +0800 CST|
|[form](https://github.com/go-playground/form) | 324| 8| 2016-05-26 21:26:40 +0800 CST| 2019-02-03 09:53:44 +0800 CST|
|[formam](https://github.com/monoculum/formam) | 111| 2| 2014-10-25 08:23:30 +0800 CST| 2018-09-01 09:54:15 +0800 CST|
|[forms](https://github.com/albrow/forms) | 96| 6| 2014-08-08 00:11:30 +0800 CST| 2017-07-02 20:22:45 +0800 CST|
|[csrf](https://github.com/gorilla/csrf) | 368| 19| 2015-08-03 08:35:16 +0800 CST| 2018-12-20 18:56:45 +0800 CST|
|[nosurf](https://github.com/justinas/nosurf) | 919| 35| 2013-08-23 01:47:34 +0800 CST| 2019-01-19 00:37:50 +0800 CST|

## Functional
        
*Packages to support functional programming in Go.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[fpGo](https://github.com/TeaEntityLab/fpGo) | 69| 1| 2018-05-24 17:08:45 +0800 CST| 2018-07-19 13:51:53 +0800 CST|
|[fuego](https://github.com/seborama/fuego) | 13| 1| 2018-11-06 06:24:09 +0800 CST| 2019-03-02 01:59:16 +0800 CST|
|[go-underscore](https://github.com/tobyhede/go-underscore) | 1009| 26| 2014-07-02 18:27:16 +0800 CST| 2019-02-15 05:27:45 +0800 CST|

## Game Development
        
*Awesome game development libraries.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[engine](https://github.com/azul3d/engine) | 406| 23| 2016-02-29 12:54:44 +0800 CST| 2018-06-25 06:16:41 +0800 CST|
|[ebiten](https://github.com/hajimehoshi/ebiten) | 1405| 65| 2013-06-16 23:13:01 +0800 CST| 2019-03-02 03:44:05 +0800 CST|
|[engo](https://github.com/EngoEngine/engo) | 951| 45| 2014-11-12 13:50:03 +0800 CST| 2019-03-02 02:04:46 +0800 CST|
|[engine](https://github.com/g3n/engine) | 590| 58| 2017-03-08 02:25:09 +0800 CST| 2019-02-19 21:16:25 +0800 CST|
|[GarageEngine](https://github.com/vova616/GarageEngine) | 303| 31| 2012-08-07 18:55:34 +0800 CST| 2013-09-03 19:21:00 +0800 CST|
|[glop](https://github.com/runningwild/glop) | 76| 4| 2011-04-21 06:48:34 +0800 CST| 2015-09-24 11:00:09 +0800 CST|
|[go-astar](https://github.com/beefsack/go-astar) | 307| 10| 2014-05-28 10:00:03 +0800 CST| 2018-03-26 13:33:35 +0800 CST|
|[go-collada](https://github.com/GlenKelley/go-collada) | 12| 3| 2013-09-19 12:19:15 +0800 CST| 2013-09-27 05:39:30 +0800 CST|
|[go-sdl2](https://github.com/veandco/go-sdl2) | 1051| 42| 2013-06-06 02:30:03 +0800 CST| 2019-02-23 19:28:37 +0800 CST|
|[go3d](https://github.com/ungerik/go3d) | 150| 10| 2011-06-27 21:02:26 +0800 CST| 2019-03-01 14:15:22 +0800 CST|
|[gonet](https://github.com/xtaci/gonet) | 997| 125| 2013-04-11 10:18:23 +0800 CST| 2017-05-12 15:31:41 +0800 CST|
|[goworld](https://github.com/xiaonanln/goworld) | 946| 99| 2017-06-03 23:02:46 +0800 CST| 2019-01-01 21:45:25 +0800 CST|
|[leaf](https://github.com/name5566/leaf) | 2714| 302| 2014-08-04 20:40:08 +0800 CST| 2018-12-28 22:20:19 +0800 CST|
|[nano](https://github.com/lonng/nano) | 844| 53| 2017-08-02 14:05:14 +0800 CST| 2019-01-26 17:31:10 +0800 CST|
|[oak](https://github.com/oakmound/oak) | 572| 35| 2017-07-16 00:24:27 +0800 CST| 2019-02-26 10:52:54 +0800 CST|
|[pitaya](https://github.com/topfreegames/pitaya) | 202| 26| 2018-03-20 03:40:36 +0800 CST| 2019-03-01 05:29:28 +0800 CST|
|[pixel](https://github.com/faiface/pixel) | 2111| 92| 2016-11-19 19:15:34 +0800 CST| 2019-02-23 00:13:59 +0800 CST|
|[raylib-go](https://github.com/gen2brain/raylib-go) | 340| 18| 2017-01-27 16:31:45 +0800 CST| 2018-11-18 04:47:59 +0800 CST|
|[termloop](https://github.com/JoelOtter/termloop) | 975| 35| 2015-05-24 01:12:34 +0800 CST| 2018-12-02 01:22:10 +0800 CST|

## Generation and Generics
        
*Tools to enhance the language with features like generics via code generation.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[efaceconv](https://github.com/t0pep0/efaceconv) | 39| 4| 2016-11-18 19:38:54 +0800 CST| 2017-10-12 15:16:32 +0800 CST|
|[gen](https://github.com/clipperhouse/gen) | 977| 36| 2013-10-14 04:26:36 +0800 CST| 2018-06-12 03:54:36 +0800 CST|
|[go-enum](https://github.com/abice/go-enum) | 65| 4| 2017-08-11 06:07:31 +0800 CST| 2019-02-15 06:47:19 +0800 CST|
|[go-linq](https://github.com/ahmetb/go-linq) | 1661| 72| 2013-12-19 11:05:00 +0800 CST| 2018-11-09 16:43:21 +0800 CST|
|[goderive](https://github.com/awalterschulze/goderive) | 670| 22| 2017-02-11 05:46:49 +0800 CST| 2019-02-24 20:06:16 +0800 CST|
|[gotype](https://github.com/wzshiming/gotype) | 17| 3| 2017-12-05 12:09:47 +0800 CST| 2019-02-20 15:26:44 +0800 CST|
|[gowrap](https://github.com/hexdigest/gowrap) | 203| 6| 2018-09-15 17:20:42 +0800 CST| 2019-02-26 17:51:56 +0800 CST|
|[interfaces](https://github.com/rjeczalik/interfaces) | 160| 5| 2015-12-06 08:04:50 +0800 CST| 2018-12-21 17:45:12 +0800 CST|
|[jennifer](https://github.com/dave/jennifer) | 1136| 21| 2016-12-05 04:57:38 +0800 CST| 2019-01-30 21:02:10 +0800 CST|
|[pkgreflect](https://github.com/ungerik/pkgreflect) | 77| 5| 2012-09-03 15:53:00 +0800 CST| 2017-09-05 20:27:27 +0800 CST|

## Geographic
        
*Geographic tools and servers*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[geocache](https://github.com/melihmucuk/geocache) | 95| 5| 2016-06-21 18:48:34 +0800 CST| 2016-06-22 00:53:18 +0800 CST|
|[geoserver](https://github.com/hishamkaram/geoserver) | 19| 2| 2018-03-27 05:36:49 +0800 CST| 2018-11-24 11:20:18 +0800 CST|
|[gismanager](https://github.com/hishamkaram/gismanager) | 9| 1| 2018-09-29 20:51:37 +0800 CST| 2018-10-30 16:55:19 +0800 CST|
|[osm](https://github.com/paulmach/osm) | 41| 5| 2016-02-02 08:59:03 +0800 CST| 2018-12-15 04:48:15 +0800 CST|
|[pbf](https://github.com/maguro/pbf) | 10| 1| 2017-09-19 07:13:18 +0800 CST| 2018-10-15 00:12:41 +0800 CST|
|[geo](https://github.com/golang/geo) | 800| 71| 2014-12-04 07:02:15 +0800 CST| 2019-02-21 21:59:16 +0800 CST|
|[tile38](https://github.com/tidwall/tile38) | 5314| 177| 2016-03-05 07:07:44 +0800 CST| 2019-03-01 21:56:49 +0800 CST|

## Go Compilers
        
*Tools for compiling Go to other languages.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[c4go](https://github.com/Konstantin8105/c4go) | 61| 7| 2018-03-28 14:24:57 +0800 CST| 2019-03-01 18:52:53 +0800 CST|
|[f4go](https://github.com/Konstantin8105/f4go) | 9| 2| 2018-07-09 01:05:43 +0800 CST| 2018-12-21 20:53:18 +0800 CST|
|[gopherjs](https://github.com/gopherjs/gopherjs) | 7984| 250| 2013-08-28 06:23:58 +0800 CST| 2019-02-05 13:50:24 +0800 CST|
|[llgo](https://github.com/go-llvm/llgo) | 951| 61| 2011-11-05 22:23:32 +0800 CST| 2015-01-05 09:15:37 +0800 CST|
|[tardisgo](https://github.com/tardisgo/tardisgo) | 382| 27| 2014-01-08 19:07:33 +0800 CST| 2016-11-20 02:08:43 +0800 CST|

## Goroutines
        
*Tools for managing and working with Goroutines.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[ants](https://github.com/panjf2000/ants) | 1357| 54| 2018-05-19 09:13:38 +0800 CST| 2019-02-24 09:53:29 +0800 CST|
|[artifex](https://github.com/borderstech/artifex) | 6| 2| 2018-11-01 03:34:31 +0800 CST| 2018-11-03 06:43:19 +0800 CST|
|[async](https://github.com/StudioSol/async) | 13| 9| 2017-07-01 01:08:33 +0800 CST| 2018-09-18 21:53:08 +0800 CST|
|[breaker](https://github.com/kamilsk/breaker) | 5| 1| 2019-02-15 23:08:24 +0800 CST| 2019-03-02 13:45:37 +0800 CST|
|[cyclicbarrier](https://github.com/marusama/cyclicbarrier) | 23| 3| 2018-01-11 18:38:46 +0800 CST| 2018-10-27 18:23:13 +0800 CST|
|[go-floc](https://github.com/workanator/go-floc) | 164| 8| 2017-07-03 15:34:06 +0800 CST| 2018-01-26 13:46:48 +0800 CST|
|[go-flow](https://github.com/kamildrazkiewicz/go-flow) | 94| 7| 2016-09-25 22:46:09 +0800 CST| 2017-09-19 15:20:06 +0800 CST|
|[go-trylock](https://github.com/subchen/go-trylock) | 4| 1| 2018-04-26 14:02:47 +0800 CST| 2018-05-29 12:27:46 +0800 CST|
|[GoSlaves](https://github.com/dgrr/GoSlaves) | 60| 4| 2017-09-18 03:05:14 +0800 CST| 2018-08-02 07:50:33 +0800 CST|
|[goworker](https://github.com/benmanns/goworker) | 2130| 70| 2013-07-23 01:04:27 +0800 CST| 2019-02-06 01:37:25 +0800 CST|
|[gpool](https://github.com/sherifabdlnaby/gpool) | 44| 1| 2018-12-03 12:23:35 +0800 CST| 2019-01-17 17:05:52 +0800 CST|
|[grpool](https://github.com/ivpusic/grpool) | 459| 29| 2015-07-22 08:15:04 +0800 CST| 2019-01-28 07:07:22 +0800 CST|
|[parallel-fn](https://github.com/rafaeljesus/parallel-fn) | 24| 3| 2017-06-18 17:47:54 +0800 CST| 2018-01-02 04:34:49 +0800 CST|
|[pool](https://github.com/go-playground/pool) | 451| 11| 2015-10-29 00:36:08 +0800 CST| 2016-08-24 03:37:33 +0800 CST|
|[semaphore](https://github.com/kamilsk/semaphore) | 59| 1| 2016-10-08 19:48:12 +0800 CST| 2019-03-02 03:29:27 +0800 CST|
|[semaphore](https://github.com/marusama/semaphore) | 62| 4| 2017-11-22 22:00:58 +0800 CST| 2019-01-10 15:54:42 +0800 CST|
|[stl](https://github.com/ssgreg/stl) | 5| 1| 2018-06-19 18:50:11 +0800 CST| 2018-09-28 22:57:12 +0800 CST|
|[threadpool](https://github.com/shettyh/threadpool) | 12| 2| 2017-09-07 02:45:39 +0800 CST| 2017-12-12 00:31:30 +0800 CST|
|[tunny](https://github.com/Jeffail/tunny) | 1134| 51| 2014-04-03 00:14:58 +0800 CST| 2018-11-09 04:57:23 +0800 CST|
|[worker-pool](https://github.com/vardius/worker-pool) | 29| 3| 2017-10-04 17:18:31 +0800 CST| 2019-02-05 18:10:09 +0800 CST|
|[workerpool](https://github.com/gammazero/workerpool) | 79| 5| 2016-05-17 22:32:06 +0800 CST| 2018-12-31 04:30:50 +0800 CST|

## GUI
        
*Interaction*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[app](https://github.com/maxence-charriere/app) | 2748| 106| 2016-10-12 08:31:33 +0800 CST| 2019-03-01 10:09:46 +0800 CST|
|[fyne](https://github.com/fyne-io/fyne) | 745| 40| 2018-02-05 06:07:16 +0800 CST| 2019-03-01 07:45:40 +0800 CST|
|[go-astilectron](https://github.com/asticode/go-astilectron) | 2353| 112| 2017-04-22 15:59:15 +0800 CST| 2019-01-11 17:38:57 +0800 CST|
|[go-sciter](https://github.com/sciter-sdk/go-sciter) | 1296| 113| 2015-10-15 20:41:06 +0800 CST| 2018-11-01 03:25:52 +0800 CST|
|[gotk3](https://github.com/gotk3/gotk3) | 617| 43| 2015-08-13 21:09:46 +0800 CST| 2019-02-28 02:37:47 +0800 CST|
|[gowd](https://github.com/dtylman/gowd) | 172| 19| 2017-03-29 22:50:53 +0800 CST| 2018-10-25 17:08:40 +0800 CST|
|[qt](https://github.com/therecipe/qt) | 5257| 274| 2014-11-19 08:03:08 +0800 CST| 2019-01-22 06:05:19 +0800 CST|
|[ui](https://github.com/andlabs/ui) | 6443| 355| 2014-02-18 07:44:00 +0800 CST| 2019-02-07 00:27:53 +0800 CST|
|[walk](https://github.com/lxn/walk) | 3274| 239| 2010-09-16 16:11:49 +0800 CST| 2019-02-28 19:54:49 +0800 CST|
|[webview](https://github.com/zserge/webview) | 3888| 173| 2017-08-19 16:26:00 +0800 CST| 2019-02-20 08:46:14 +0800 CST|
|[gosx-notifier](https://github.com/deckarep/gosx-notifier) | 476| 15| 2013-11-25 14:35:16 +0800 CST| 2018-02-01 11:58:18 +0800 CST|
|[robotgo](https://github.com/go-vgo/robotgo) | 3910| 176| 2016-09-27 00:26:56 +0800 CST| 2019-03-01 19:57:51 +0800 CST|
|[systray](https://github.com/getlantern/systray) | 634| 38| 2014-11-12 11:41:57 +0800 CST| 2019-02-17 00:38:38 +0800 CST|
|[trayhost](https://github.com/shurcooL/trayhost) | 143| 4| 2014-04-25 12:02:30 +0800 CST| 2018-10-21 05:16:47 +0800 CST|

## Hardware
        
*Libraries, tools, and tutorials for interacting with hardware.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|

## Images
        
*Libraries for programming devices of the IoT.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[bild](https://github.com/anthonynsimon/bild) | 1930| 55| 2016-08-01 23:54:29 +0800 CST| 2018-10-05 06:01:15 +0800 CST|
|[bimg](https://github.com/h2non/bimg) | 692| 35| 2015-03-17 22:14:02 +0800 CST| 2019-02-23 05:48:24 +0800 CST|
|[cameron](https://github.com/aofei/cameron) | 17| 1| 2018-05-06 06:13:11 +0800 CST| 2019-02-26 16:49:18 +0800 CST|
|[geopattern](https://github.com/pravj/geopattern) | 982| 20| 2014-10-23 01:26:30 +0800 CST| 2019-01-09 04:17:57 +0800 CST|
|[gg](https://github.com/fogleman/gg) | 1708| 68| 2016-02-19 05:05:08 +0800 CST| 2019-02-21 06:12:58 +0800 CST|
|[gift](https://github.com/disintegration/gift) | 1155| 49| 2014-07-13 02:47:40 +0800 CST| 2018-10-06 06:10:33 +0800 CST|
|[go-cairo](https://github.com/ungerik/go-cairo) | 83| 6| 2012-08-23 02:27:01 +0800 CST| 2018-09-10 22:37:57 +0800 CST|
|[go-gd](https://github.com/bolknote/go-gd) | 47| 4| 2011-05-12 14:33:54 +0800 CST| 2018-05-08 03:29:26 +0800 CST|
|[go-nude](https://github.com/koyachi/go-nude) | 271| 17| 2014-05-02 16:32:29 +0800 CST| 2018-11-22 23:22:42 +0800 CST|
|[go-opencv](https://github.com/go-opencv/go-opencv) | 1025| 68| 2013-12-09 17:43:26 +0800 CST| 2019-01-15 15:07:43 +0800 CST|
|[go-webcolors](https://github.com/jyotiska/go-webcolors) | 23| 1| 2014-04-24 22:41:22 +0800 CST| 2015-08-21 12:56:56 +0800 CST|
|[gocv](https://github.com/hybridgroup/gocv) | 1988| 102| 2017-09-19 05:54:17 +0800 CST| 2019-02-20 18:16:57 +0800 CST|
|[goimagehash](https://github.com/corona10/goimagehash) | 167| 9| 2017-07-29 01:15:58 +0800 CST| 2019-02-08 17:02:26 +0800 CST|
|[govatar](https://github.com/o1egl/govatar) | 280| 5| 2016-01-18 20:12:28 +0800 CST| 2018-04-13 08:07:57 +0800 CST|
|[image2ascii](https://github.com/qeesung/image2ascii) | 238| 7| 2018-10-20 13:06:25 +0800 CST| 2018-11-07 23:15:43 +0800 CST|
|[imagick](https://github.com/gographics/imagick) | 883| 54| 2013-05-01 01:31:48 +0800 CST| 2019-01-15 10:11:26 +0800 CST|
|[imaginary](https://github.com/h2non/imaginary) | 2307| 70| 2015-03-05 02:51:40 +0800 CST| 2019-02-25 16:54:52 +0800 CST|
|[imaging](https://github.com/disintegration/imaging) | 2181| 63| 2012-12-07 04:21:21 +0800 CST| 2019-02-03 00:04:28 +0800 CST|
|[img](https://github.com/hawx/img) | 126| 5| 2012-07-29 03:57:47 +0800 CST| 2015-05-01 23:11:26 +0800 CST|
|[ln](https://github.com/fogleman/ln) | 2377| 89| 2016-01-10 12:28:10 +0800 CST| 2018-02-24 02:16:49 +0800 CST|
|[mergi](https://github.com/noelyahan/mergi) | 42| 2| 2018-09-24 11:40:47 +0800 CST| 2019-02-25 13:59:04 +0800 CST|
|[mort](https://github.com/aldor007/mort) | 341| 17| 2017-11-19 21:37:58 +0800 CST| 2019-02-19 04:07:09 +0800 CST|
|[mpo](https://github.com/donatj/mpo) | 5| 1| 2015-04-15 06:37:59 +0800 CST| 2019-01-26 12:06:30 +0800 CST|
|[picfit](https://github.com/thoas/picfit) | 964| 41| 2014-12-07 01:30:45 +0800 CST| 2019-01-29 23:39:30 +0800 CST|
|[pt](https://github.com/fogleman/pt) | 1715| 62| 2015-01-24 03:39:29 +0800 CST| 2018-03-09 10:44:07 +0800 CST|
|[resize](https://github.com/nfnt/resize) | 2020| 73| 2012-08-03 03:48:26 +0800 CST| 2018-02-22 03:10:15 +0800 CST|
|[rez](https://github.com/bamiaux/rez) | 156| 9| 2014-01-17 05:16:15 +0800 CST| 2017-08-01 02:51:31 +0800 CST|
|[smartcrop](https://github.com/muesli/smartcrop) | 1203| 30| 2014-04-08 06:40:03 +0800 CST| 2018-10-31 06:06:01 +0800 CST|
|[steganography](https://github.com/auyer/steganography) | 12| 3| 2018-05-22 01:27:36 +0800 CST| 2018-10-18 23:00:19 +0800 CST|
|[stegify](https://github.com/DimitarPetrov/stegify) | 280| 13| 2018-11-30 00:45:58 +0800 CST| 2019-02-12 23:23:43 +0800 CST|
|[svgo](https://github.com/ajstarks/svgo) | 1220| 46| 2010-03-06 07:24:10 +0800 CST| 2018-10-06 08:33:21 +0800 CST|
|[tga](https://github.com/ftrvxmtrx/tga) | 21| 2| 2012-10-08 09:09:20 +0800 CST| 2015-05-24 16:11:41 +0800 CST|
|[connectordb](https://github.com/connectordb/connectordb) | 152| 17| 2015-01-17 03:44:21 +0800 CST| 2019-03-01 07:33:49 +0800 CST|
|[devices](https://github.com/goiot/devices) | 221| 17| 2016-05-30 16:07:02 +0800 CST| 2016-07-10 08:46:08 +0800 CST|
|[eywa](https://github.com/xcodersun/eywa) | 30| 8| 2016-02-21 01:02:16 +0800 CST| 2017-04-12 15:41:51 +0800 CST|
|[flogo](https://github.com/TIBCOSoftware/flogo) | 951| 112| 2016-07-10 10:57:43 +0800 CST| 2019-03-01 08:52:46 +0800 CST|
|[gatt](https://github.com/paypal/gatt) | 767| 57| 2014-04-23 21:45:27 +0800 CST| 2018-11-23 19:01:00 +0800 CST|
|[gobot](https://github.com/hybridgroup/gobot) | 5196| 298| 2013-09-21 22:09:19 +0800 CST| 2019-02-14 02:17:11 +0800 CST|
|[huego](https://github.com/amimof/huego) | 83| 1| 2017-05-16 13:31:45 +0800 CST| 2019-02-27 02:14:03 +0800 CST|
|[iot](https://github.com/vaelen/iot) | 28| 4| 2018-03-08 14:51:51 +0800 CST| 2018-04-17 15:28:57 +0800 CST|
|[mainflux](https://github.com/mainflux/mainflux) | 359| 53| 2015-07-07 04:31:50 +0800 CST| 2019-03-02 13:22:09 +0800 CST|
|[sensorbee](https://github.com/sensorbee/sensorbee) | 165| 19| 2016-02-19 15:49:56 +0800 CST| 2017-12-31 08:14:29 +0800 CST|

## IoT
        

## Job Scheduler
        
*Libraries for scheduling jobs.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[clockwork](https://github.com/whiteShtef/clockwork) | 48| 1| 2018-04-24 01:43:09 +0800 CST| 2019-02-13 08:01:41 +0800 CST|
|[go-cron](https://github.com/rk/go-cron) | 169| 9| 2011-04-15 22:50:49 +0800 CST| 2015-05-15 08:30:45 +0800 CST|
|[gron](https://github.com/roylee0704/gron) | 551| 16| 2016-06-04 16:02:22 +0800 CST| 2019-02-24 20:15:10 +0800 CST|
|[jobrunner](https://github.com/bamzi/jobrunner) | 516| 17| 2015-10-21 12:17:01 +0800 CST| 2016-10-19 22:30:22 +0800 CST|
|[jobs](https://github.com/albrow/jobs) | 438| 18| 2015-02-10 06:13:29 +0800 CST| 2018-06-17 05:00:16 +0800 CST|
|[leprechaun](https://github.com/kilgaloon/leprechaun) | 22| 4| 2018-04-08 21:44:04 +0800 CST| 2019-02-28 19:03:38 +0800 CST|
|[scheduler](https://github.com/carlescere/scheduler) | 266| 16| 2015-02-04 01:10:23 +0800 CST| 2018-06-14 04:05:12 +0800 CST|

## JSON
        
*Libraries for working with JSON.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[gjson](https://github.com/tidwall/gjson) | 4077| 111| 2016-08-11 11:08:47 +0800 CST| 2019-02-17 10:21:13 +0800 CST|
|[go-respond](https://github.com/nicklaw5/go-respond) | 18| 1| 2017-03-13 05:00:54 +0800 CST| 2018-10-31 02:37:50 +0800 CST|
|[gojq](https://github.com/elgs/gojq) | 132| 5| 2015-12-30 17:02:13 +0800 CST| 2018-05-31 07:43:05 +0800 CST|
|[gojson](https://github.com/ChimeraCoder/gojson) | 1917| 45| 2012-12-28 03:10:50 +0800 CST| 2019-02-13 09:14:39 +0800 CST|
|[jaydiff](https://github.com/yazgazan/jaydiff) | 33| 1| 2017-04-25 00:05:35 +0800 CST| 2018-12-06 21:15:54 +0800 CST|
|[jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) | 5| 1| 2016-07-08 18:08:58 +0800 CST| 2016-11-18 00:02:12 +0800 CST|
|[jsonf](https://github.com/miolini/jsonf) | 53| 3| 2015-05-25 12:53:32 +0800 CST| 2016-07-08 08:43:10 +0800 CST|
|[jsongo](https://github.com/ricardolonga/jsongo) | 85| 1| 2015-08-08 07:23:17 +0800 CST| 2016-12-15 19:09:33 +0800 CST|
|[jsonhal](https://github.com/RichardKnop/jsonhal) | 8| 2| 2016-01-15 19:38:40 +0800 CST| 2018-11-01 11:57:04 +0800 CST|
|[kazaam](https://github.com/qntfy/kazaam) | 106| 15| 2016-07-19 22:19:03 +0800 CST| 2018-08-29 09:46:03 +0800 CST|
|[mp](https://github.com/sanbornm/mp) | 29| 2| 2014-06-16 05:14:39 +0800 CST| 2016-05-12 03:40:58 +0800 CST|

## Logging
        
*Libraries for generating and working with log files.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[gone](https://github.com/One-com/gone) | 24| 7| 2016-09-05 17:39:11 +0800 CST| 2019-02-08 20:52:36 +0800 CST|
|[distillog](https://github.com/amoghe/distillog) | 15| 1| 2015-10-13 00:32:21 +0800 CST| 2018-07-27 07:35:13 +0800 CST|
|[glg](https://github.com/kpango/glg) | 29| 3| 2017-06-21 21:26:16 +0800 CST| 2019-02-27 12:05:56 +0800 CST|
|[glo](https://github.com/lajosbencz/glo) | 3| 1| 2019-01-20 06:10:42 +0800 CST| 2019-01-23 19:35:10 +0800 CST|
|[glog](https://github.com/golang/glog) | 2147| 84| 2013-07-16 12:33:04 +0800 CST| 2018-11-10 17:44:40 +0800 CST|
|[go-cronowriter](https://github.com/utahta/go-cronowriter) | 17| 1| 2017-02-04 17:02:55 +0800 CST| 2018-09-01 09:11:16 +0800 CST|
|[go-log](https://github.com/subchen/go-log) | 8| 1| 2017-05-07 16:09:24 +0800 CST| 2018-05-19 16:03:37 +0800 CST|
|[go-log](https://github.com/siddontang/go-log) | 20| 3| 2014-05-18 11:41:55 +0800 CST| 2019-02-21 10:24:31 +0800 CST|
|[go-log](https://github.com/ian-kent/go-log) | 33| 2| 2014-05-02 08:34:09 +0800 CST| 2018-03-31 10:06:55 +0800 CST|
|[go-logger](https://github.com/apsdehal/go-logger) | 212| 8| 2014-09-26 12:57:06 +0800 CST| 2018-09-30 07:48:05 +0800 CST|
|[gologger](https://github.com/sadlil/gologger) | 38| 6| 2015-09-02 16:52:26 +0800 CST| 2018-01-31 11:17:58 +0800 CST|
|[gomol](https://github.com/aphistic/gomol) | 12| 2| 2015-08-30 23:51:46 +0800 CST| 2018-05-30 00:49:40 +0800 CST|
|[journald](https://github.com/ssgreg/journald) | 17| 2| 2017-08-23 15:06:09 +0800 CST| 2018-12-25 16:19:33 +0800 CST|
|[log](https://github.com/apex/log) | 675| 35| 2015-12-22 04:27:48 +0800 CST| 2018-10-12 00:46:42 +0800 CST|
|[log](https://github.com/go-playground/log) | 257| 10| 2016-02-08 00:17:48 +0800 CST| 2019-02-13 23:35:32 +0800 CST|
|[log](https://github.com/teris-io/log) | 21| 1| 2017-10-29 03:57:55 +0800 CST| 2017-12-05 02:53:45 +0800 CST|
|[logvoyage](https://github.com/firstrow/logvoyage) | 83| 5| 2015-03-29 19:05:09 +0800 CST| 2017-05-25 03:48:17 +0800 CST|
|[log15](https://github.com/inconshreveable/log15) | 843| 23| 2014-05-20 08:11:52 +0800 CST| 2018-10-15 22:41:02 +0800 CST|
|[logdump](https://github.com/ewwwwwqm/logdump) | 7| 2| 2017-01-13 23:34:31 +0800 CST| 2018-04-02 08:28:16 +0800 CST|
|[logex](https://github.com/chzyer/logex) | 32| 7| 2014-10-10 14:38:39 +0800 CST| 2017-03-29 14:49:08 +0800 CST|
|[logger](https://github.com/azer/logger) | 130| 4| 2014-09-30 14:45:09 +0800 CST| 2018-09-23 15:57:13 +0800 CST|
|[logmatic](https://github.com/borderstech/logmatic) | 3| 1| 2018-11-07 09:52:45 +0800 CST| 2019-01-13 22:47:44 +0800 CST|
|[logo](https://github.com/mbndr/logo) | 4| 1| 2017-02-08 02:02:55 +0800 CST| 2017-10-20 03:33:23 +0800 CST|
|[logrus](https://github.com/sirupsen/logrus) | 9960| 268| 2013-10-17 03:08:55 +0800 CST| 2019-02-27 20:42:50 +0800 CST|
|[logrusly](https://github.com/sebest/logrusly) | 24| 5| 2014-09-12 07:27:11 +0800 CST| 2018-03-16 03:02:19 +0800 CST|
|[logutils](https://github.com/hashicorp/logutils) | 239| 87| 2013-10-09 15:31:15 +0800 CST| 2018-08-29 00:26:51 +0800 CST|
|[logxi](https://github.com/mgutz/logxi) | 328| 10| 2015-03-02 06:13:45 +0800 CST| 2017-12-24 04:05:30 +0800 CST|
|[lumberjack](https://github.com/natefinch/lumberjack) | 1187| 37| 2014-06-14 19:55:47 +0800 CST| 2019-03-01 21:50:28 +0800 CST|
|[mlog](https://github.com/jbrodriguez/mlog) | 16| 2| 2014-10-20 23:06:26 +0800 CST| 2018-08-06 01:35:46 +0800 CST|
|[onelog](https://github.com/francoispqt/onelog) | 303| 9| 2018-05-06 22:32:10 +0800 CST| 2019-03-01 15:51:24 +0800 CST|
|[ozzo-log](https://github.com/go-ozzo/ozzo-log) | 101| 11| 2015-10-23 06:29:02 +0800 CST| 2018-04-06 17:37:37 +0800 CST|
|[seelog](https://github.com/cihub/seelog) | 1269| 89| 2011-11-17 17:43:15 +0800 CST| 2019-02-28 20:17:08 +0800 CST|
|[go-spew](https://github.com/davecgh/go-spew) | 2976| 59| 2013-01-09 13:18:22 +0800 CST| 2019-01-31 13:51:44 +0800 CST|
|[log](https://github.com/alexcesaro/log) | 44| 4| 2014-04-19 22:31:56 +0800 CST| 2015-09-16 06:13:22 +0800 CST|
|[tail](https://github.com/hpcloud/tail) | 1358| 101| 2013-02-05 08:28:03 +0800 CST| 2019-02-28 04:56:24 +0800 CST|
|[xlog](https://github.com/xfxdev/xlog) | 4| 1| 2016-05-06 00:47:45 +0800 CST| 2019-01-15 18:17:30 +0800 CST|
|[xlog](https://github.com/rs/xlog) | 126| 7| 2015-10-22 17:26:45 +0800 CST| 2018-03-20 02:49:19 +0800 CST|
|[zap](https://github.com/uber-go/zap) | 6129| 199| 2016-02-19 03:52:56 +0800 CST| 2019-02-15 09:25:42 +0800 CST|
|[zerolog](https://github.com/rs/zerolog) | 1764| 47| 2017-05-12 13:24:39 +0800 CST| 2019-03-02 11:16:32 +0800 CST|

## Machine Learning
        
*Libraries for Machine Learning.*
| Go Repo    |Stars      |Watchers   |Created At | Pushed At |
| :--------- |:---------:| ---------:|:---------:|:---------:|
|[bayesian](https://github.com/jbrukh/bayesian) | 596| 30| 2011-11-23 12:17:00 +0800 CST| 2019-02-18 12:38:19 +0800 CST|
|[CloudForest](https://github.com/ryanbressler/CloudForest) | 627| 43| 2012-10-23 01:38:16 +0800 CST| 2018-11-20 14:58:46 +0800 CST|
|[eaopt](https://github.com/MaxHalford/eaopt) | 578| 28| 2016-01-31 08:04:52 +0800 CST| 2019-02-20 03:56:00 +0800 CST|
|[evoli](https://github.com/khezen/evoli) | 5| 3| 2015-06-12 14:58:30 +0800 CST| 2019-01-28 20:40:57 +0800 CST|
|[fonet](https://github.com/Fontinalis/fonet) | 26| 3| 2017-10-03 23:57:15 +0800 CST| 2018-04-10 21:50:54 +0800 CST|
|[go-cluster](https://github.com/e-XpertSolutions/go-cluster) | 18| 5| 2017-10-04 20:24:52 +0800 CST| 2018-08-06 15:35:27 +0800 CST|
|[go-deep](https://github.com/patrikeh/go-deep) | 201| 11| 2017-12-09 23:10:06 +0800 CST| 2018-09-14 20:17:29 +0800 CST|
|[go-fann](https://github.com/white-pony/go-fann) | 98| 8| 2011-03-11 05:08:27 +0800 CST| 2015-02-04 05:53:31 +0800 CST|
|[go-galib](https://github.com/thoj/go-galib) | 164| 14| 2009-11-30 18:46:58 +0800 CST| 2015-12-29 00:27:45 +0800 CST|
|[go-pr](https://github.com/daviddengcn/go-pr) | 56| 6| 2013-06-07 10:36:20 +0800 CST| 2013-06-08 18:17:05 +0800 CST|
|[gobrain](https://github.com/goml/gobrain) | 336| 23| 2014-04-29 21:32:36 +0800 CST| 2018-12-09 21:17:54 +0800 CST|
|[godist](https://github.com/e-dard/godist) | 22| 2| 2014-09-05 17:48:51 +0800 CST| 2015-05-11 18:38:48 +0800 CST|
|[goga](https://github.com/tomcraven/goga) | 72| 8| 2015-10-20 20:50:51 +0800 CST| 2017-01-16 23:29:16 +0800 CST|
|[golearn](https://github.com/sjwhitworth/golearn) | 6303| 429| 2013-12-26 21:06:14 +0800 CST| 2019-02-07 04:42:36 +0800 CST|
|[golinear](https://github.com/danieldk/golinear) | 38| 5| 2013-04-05 23:37:01 +0800 CST| 2018-08-29 18:30:44 +0800 CST|
|[gomind](https://github.com/surenderthakran/gomind) | 5| 2| 2017-10-19 11:48:51 +0800 CST| 2018-07-31 20:57:31 +0800 CST|
|[goml](https://github.com/cdipaolo/goml) | 960| 73| 2015-06-27 13:52:01 +0800 CST| 2016-10-31 04:49:12 +0800 CST|
|[goRecommend](https://github.com/timkaye11/goRecommend) | 136| 10| 2014-07-16 13:32:23 +0800 CST| 2014-07-29 12:49:57 +0800 CST|
|[gorgonia](https://github.com/gorgonia/gorgonia) | 2443| 161| 2016-09-15 07:19:43 +0800 CST| 2019-03-02 00:30:55 +0800 CST|
|[gorse](https://github.com/zhenghaoz/gorse) | 338| 16| 2018-08-14 19:01:09 +0800 CST| 2019-02-28 16:00:27 +0800 CST|
|[goscore](https://github.com/asafschers/goscore) | 27| 2| 2017-08-19 19:08:39 +0800 CST| 2018-09-05 18:32:11 +0800 CST|
|[gosseract](https://github.com/otiai10/gosseract) | 744| 31| 2013-10-11 15:27:53 +0800 CST| 2019-02-25 17:01:49 +0800 CST|
|[libsvm](https://github.com/datastream/libsvm) | 59| 10| 2012-07-31 15:57:47 +0800 CST| 2016-05-09 11:47:11 +0800 CST|
|[mlgo](https://github.com/NullHypothesis/mlgo) | 4| 2| 2015-12-07 23:41:25 +0800 CST| 2015-12-08 00:08:20 +0800 CST|
|[neat](https://github.com/jinyeom/neat) | 51| 11| 2016-11-17 12:23:14 +0800 CST| 2018-07-05 04:45:55 +0800 CST|

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
        