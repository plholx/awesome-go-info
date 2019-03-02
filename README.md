# Awesome Go Data

go相关的开源项目列表，项目分类及项目数据完全来自于[awesome-go](https://github.com/avelino/awesome-go) 的[README.md](https://github.com/avelino/awesome-go/blob/master/README.md)文件，通过调用GitHub的API获取仓库信息，展示项目的star数、watch数...

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
        