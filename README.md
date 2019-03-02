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
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[EasyMIDI](https://github.com/algoGuy/EasyMIDI) | 17| 3| 2018-02-20 04:18:09| 2018-03-22 13:16:54|
|[flac](https://github.com/eaburns/flac) | 81| 5| 2013-08-21 01:48:58| 2017-10-04 04:06:22|
|[flac](https://github.com/mewkiz/flac) | 92| 6| 2012-11-02 04:13:58| 2019-02-22 23:13:57|
|[gaad](https://github.com/Comcast/gaad) | 50| 10| 2016-07-11 22:19:16| 2018-02-21 05:08:16|
|[go-sox](https://github.com/krig/go-sox) | 87| 8| 2013-10-08 22:11:04| 2018-06-21 02:14:36|
|[go_mediainfo](https://github.com/zhulik/go_mediainfo) | 21| 1| 2015-12-14 06:23:23| 2015-12-25 04:40:36|
|[gosamplerate](https://github.com/dh1tw/gosamplerate) | 8| 1| 2016-11-21 05:19:31| 2018-06-04 07:37:17|
|[id3v2](https://github.com/bogem/id3v2) | 90| 5| 2016-05-16 02:36:53| 2019-02-08 23:52:24|
|[malgo](https://github.com/gen2brain/malgo) | 55| 4| 2017-11-10 02:27:52| 2019-02-02 18:30:11|
|[minimp3](https://github.com/tosone/minimp3) | 19| 2| 2018-01-26 17:10:31| 2019-03-02 12:02:21|
|[mix](https://github.com/go-mix/mix) | 90| 3| 2016-01-03 23:53:57| 2017-06-25 02:22:11|
|[mp3](https://github.com/tcolgate/mp3) | 85| 1| 2015-02-27 04:37:37| 2017-04-27 03:41:34|
|[music-theory](https://github.com/go-music-theory/music-theory) | 233| 10| 2016-03-17 11:50:17| 2018-06-13 04:13:26|
|[portaudio](https://github.com/gordonklaus/portaudio) | 261| 9| 2015-09-16 15:59:23| 2018-08-20 06:23:52|
|[portmidi](https://github.com/rakyll/portmidi) | 187| 6| 2013-11-10 22:24:56| 2017-07-16 11:23:45|
|[go-taglib](https://github.com/wtolson/go-taglib) | 64| 6| 2012-11-20 09:03:40| 2018-07-18 08:00:47|
|[vorbis](https://github.com/mccoyst/vorbis) | 21| 3| 2013-07-12 10:45:39| 2017-10-14 23:28:32|
|[waveform](https://github.com/mdlayher/waveform) | 228| 13| 2014-09-14 02:07:36| 2016-07-07 23:25:25|

## Authentication and OAuth
        
*Libraries for implementing authentications schemes.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[authboss](https://github.com/volatiletech/authboss) | 1733| 38| 2015-01-03 13:12:02| 2019-02-07 12:11:42|
|[branca](https://github.com/hako/branca) | 49| 5| 2018-01-09 23:27:31| 2018-08-08 08:04:55|
|[casbin](https://github.com/casbin/casbin) | 3845| 140| 2017-04-08 15:51:23| 2019-02-12 21:48:40|
|[cookiestxt](https://github.com/mengzhuo/cookiestxt) | 2| 1| 2017-10-09 19:27:19| 2017-10-09 20:14:58|
|[go-jose](https://github.com/square/go-jose) | 1014| 61| 2014-11-15 02:27:31| 2019-02-27 08:14:27|
|[go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) | 1113| 71| 2015-11-01 21:30:09| 2019-02-21 21:39:31|
|[gologin](https://github.com/dghubble/gologin) | 960| 27| 2015-06-23 12:40:52| 2019-02-14 07:18:24|
|[gorbac](https://github.com/mikespook/gorbac) | 828| 55| 2013-12-26 18:00:41| 2019-01-31 14:21:04|
|[goth](https://github.com/markbates/goth) | 2081| 58| 2014-10-15 04:38:12| 2019-02-22 03:00:59|
|[httpauth](https://github.com/goji/httpauth) | 167| 5| 2014-05-27 06:53:57| 2016-06-01 21:53:05|
|[jwt](https://github.com/robbert229/jwt) | 60| 6| 2016-06-06 06:01:37| 2018-11-09 01:39:15|
|[jwt](https://github.com/pascaldekloe/jwt) | 38| 2| 2018-03-21 19:59:53| 2019-01-18 06:59:13|
|[jwt-auth](https://github.com/adam-hanna/jwt-auth) | 144| 9| 2016-07-06 07:31:43| 2019-01-30 23:42:05|
|[jwt-go](https://github.com/dgrijalva/jwt-go) | 5009| 130| 2012-04-18 09:41:49| 2019-02-22 01:13:13|
|[loginsrv](https://github.com/tarent/loginsrv) | 734| 45| 2016-11-11 20:11:21| 2019-02-19 04:45:04|
|[oauth2](https://github.com/golang/oauth2) | 2108| 92| 2014-04-14 23:07:35| 2019-02-27 04:54:24|
|[osin](https://github.com/openshift/osin) | 1506| 69| 2013-09-11 03:52:00| 2018-10-17 05:40:43|
|[paseto](https://github.com/o1egl/paseto) | 194| 9| 2018-01-23 13:27:39| 2018-11-08 23:32:44|
|[permissions2](https://github.com/xyproto/permissions2) | 297| 12| 2014-11-19 20:23:37| 2019-02-21 11:07:58|
|[rbac](https://github.com/zpatrick/rbac) | 20| 3| 2018-08-02 08:11:04| 2018-08-30 03:03:47|
|[securecookie](https://github.com/chmike/securecookie) | 26| 4| 2017-09-03 22:40:29| 2018-08-31 23:04:40|
|[session](https://github.com/icza/session) | 82| 6| 2016-02-08 17:07:07| 2018-09-10 17:22:58|
|[sessiongate-go](https://github.com/f0rmiga/sessiongate-go) | 6| 2| 2017-10-20 11:39:11| 2018-11-10 03:25:29|
|[sessions](https://github.com/adam-hanna/sessions) | 41| 3| 2017-04-29 09:09:28| 2017-11-29 02:34:24|
|[signedvalue](https://github.com/sashka/signedvalue) | 6| 0| 2018-01-06 13:57:09| 2019-01-28 19:42:41|

## Bot Building
        
*Libraries for building and working with bots.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[bot](https://github.com/go-chat-bot/bot) | 412| 34| 2015-09-23 00:41:13| 2019-02-19 19:36:50|
|[go-sarah](https://github.com/oklahomer/go-sarah) | 111| 3| 2016-11-06 18:04:43| 2018-07-31 21:21:57|
|[go-tgbot](https://github.com/olebedev/go-tgbot) | 80| 7| 2016-12-11 14:06:32| 2018-06-25 12:50:26|
|[golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot) | 147| 19| 2017-05-15 06:11:41| 2019-02-11 21:02:32|
|[govkbot](https://github.com/nikepan/govkbot) | 19| 2| 2016-07-12 06:09:54| 2019-03-02 02:39:58|
|[hanu](https://github.com/sbstjn/hanu) | 90| 6| 2016-09-16 15:10:42| 2018-09-04 22:42:13|
|[margelet](https://github.com/zhulik/margelet) | 51| 5| 2015-11-21 21:02:17| 2016-09-18 19:47:01|
|[micha](https://github.com/onrik/micha) | 9| 3| 2016-04-14 20:09:44| 2018-02-15 19:45:17|
|[slacker](https://github.com/shomali11/slacker) | 256| 13| 2017-05-20 09:41:20| 2019-02-22 10:40:29|
|[tbot](https://github.com/yanzay/tbot) | 185| 8| 2015-09-12 00:19:25| 2018-12-16 04:28:17|
|[telebot](https://github.com/tucnak/telebot) | 823| 34| 2015-06-26 03:27:50| 2019-02-27 06:51:36|
|[telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) | 1386| 56| 2015-06-25 13:33:57| 2019-03-02 00:58:51|
|[tenyks](https://github.com/kyleterry/tenyks) | 166| 14| 2012-08-26 10:02:24| 2017-03-05 16:10:44|

## Command Line
        
*Libraries for building Console Applications and Console User Interfaces.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[argparse](https://github.com/akamensky/argparse) | 74| 5| 2017-11-24 14:42:20| 2019-01-15 17:47:22|
|[argv](https://github.com/cosiner/argv) | 11| 1| 2017-01-22 18:37:21| 2019-01-15 20:05:58|
|[cli](https://github.com/mkideal/cli) | 427| 20| 2016-02-27 00:45:29| 2019-02-28 10:37:39|
|[cli](https://github.com/teris-io/cli) | 43| 1| 2017-05-25 07:07:07| 2018-07-06 04:41:43|
|[gcli](https://github.com/tcnksm/gcli) | 856| 25| 2014-06-19 17:10:15| 2017-11-20 12:20:18|
|[cobra](https://github.com/spf13/cobra) | 10748| 255| 2013-09-04 04:40:26| 2019-02-23 08:16:26|
|[commandeer](https://github.com/jaffee/commandeer) | 70| 4| 2017-10-12 10:51:05| 2019-03-02 01:02:58|
|[complete](https://github.com/posener/complete) | 540| 12| 2017-05-06 05:34:07| 2019-02-20 19:52:52|
|[docopt.go](https://github.com/docopt/docopt.go) | 1035| 36| 2013-08-26 07:05:52| 2018-09-25 07:40:37|
|[env](https://github.com/codingconcepts/env) | 37| 1| 2017-06-15 04:01:55| 2018-10-16 16:32:45|
|[flag](https://github.com/cosiner/flag) | 90| 5| 2016-10-06 00:49:41| 2019-02-20 23:50:24|
|[flaggy](https://github.com/integrii/flaggy) | 410| 11| 2018-03-05 13:55:05| 2019-02-20 13:21:20|
|[flagvar](https://github.com/sgreben/flagvar) | 26| 1| 2018-05-19 02:45:16| 2018-11-02 05:59:23|
|[go-arg](https://github.com/alexflint/go-arg) | 571| 15| 2015-11-01 09:30:06| 2018-12-28 04:00:43|
|[go-commander](https://github.com/yitsushi/go-commander) | 8| 1| 2016-10-10 18:09:41| 2019-02-26 16:31:26|
|[go-flags](https://github.com/jessevdk/go-flags) | 1290| 28| 2012-08-31 21:57:58| 2019-02-25 03:46:59|
|[go-getoptions](https://github.com/DavidGamba/go-getoptions) | 5| 1| 2015-12-18 10:21:14| 2019-02-23 07:57:43|
|[gocmd](https://github.com/devfacet/gocmd) | 28| 3| 2018-01-08 12:52:02| 2018-09-05 11:42:24|
|[hiboot](https://github.com/hidevopsio/hiboot) | 59| 6| 2018-03-16 19:21:46| 2019-02-26 12:03:41|
|[kingpin](https://github.com/alecthomas/kingpin) | 2296| 54| 2014-05-15 04:09:04| 2019-01-14 14:40:04|
|[liner](https://github.com/peterh/liner) | 520| 22| 2012-08-16 00:34:55| 2019-02-11 10:05:18|
|[cli](https://github.com/mitchellh/cli) | 925| 25| 2013-11-03 14:47:54| 2018-11-25 14:16:29|
|[mow.cli](https://github.com/jawher/mow.cli) | 590| 19| 2014-12-19 03:34:20| 2019-02-26 14:55:10|
|[pflag](https://github.com/spf13/pflag) | 623| 24| 2013-08-30 22:53:31| 2019-02-11 19:02:58|
|[readline](https://github.com/chzyer/readline) | 1305| 39| 2015-09-20 23:11:30| 2019-02-08 20:42:45|
|[sand](https://github.com/Zaba505/sand) | 2| 1| 2018-11-19 06:44:41| 2018-11-22 03:13:47|
|[sflags](https://github.com/octago/sflags) | 71| 5| 2016-12-04 22:49:27| 2019-02-03 07:42:37|
|[strumt](https://github.com/antham/strumt) | 22| 0| 2017-06-20 03:33:16| 2018-02-17 19:11:54|
|[clif](https://github.com/ukautz/clif) | 95| 2| 2015-05-31 02:30:08| 2019-02-18 22:43:25|
|[cli](https://github.com/urfave/cli) | 10066| 271| 2013-07-14 03:32:06| 2019-02-08 08:35:01|
|[wlog](https://github.com/dixonwille/wlog) | 29| 1| 2016-04-14 00:47:40| 2017-07-21 07:52:41|
|[wmenu](https://github.com/dixonwille/wmenu) | 67| 2| 2016-04-20 21:09:44| 2019-02-21 08:45:53|
|[asciigraph](https://github.com/guptarohit/asciigraph) | 1019| 23| 2018-06-17 18:37:16| 2019-01-12 21:09:29|
|[aurora](https://github.com/logrusorgru/aurora) | 425| 5| 2016-11-07 06:37:12| 2018-10-03 03:45:21|
|[cfmt](https://github.com/mingrammer/cfmt) | 53| 3| 2018-03-16 03:04:27| 2018-12-08 01:31:52|
|[chalk](https://github.com/ttacon/chalk) | 286| 11| 2014-07-19 03:38:58| 2016-06-27 04:24:23|
|[color](https://github.com/fatih/color) | 2861| 56| 2014-02-17 17:13:35| 2018-10-11 07:14:14|
|[colourize](https://github.com/TreyBastian/colourize) | 15| 3| 2015-05-11 19:49:39| 2016-05-10 17:50:02|
|[ctc](https://github.com/wzshiming/ctc) | 6| 1| 2018-04-28 02:07:42| 2018-10-31 17:40:47|
|[go-ataman](https://github.com/workanator/go-ataman) | 8| 2| 2017-05-18 03:04:57| 2017-09-25 14:09:08|
|[go-colorable](https://github.com/mattn/go-colorable) | 334| 16| 2014-07-30 10:38:06| 2019-02-23 00:32:50|
|[go-colortext](https://github.com/daviddengcn/go-colortext) | 188| 9| 2013-01-23 11:38:54| 2018-04-10 01:50:54|
|[go-isatty](https://github.com/mattn/go-isatty) | 298| 7| 2014-04-01 09:53:09| 2019-02-26 01:38:53|
|[go-prompt](https://github.com/c-bata/go-prompt) | 2078| 35| 2017-08-15 00:02:09| 2019-02-23 16:08:50|
|[gocui](https://github.com/jroimartin/gocui) | 3962| 103| 2014-01-04 10:50:20| 2019-02-20 23:24:10|
|[gommon](https://github.com/labstack/gommon) | 269| 19| 2015-03-13 06:35:57| 2019-01-26 02:56:11|
|[color](https://github.com/gookit/color) | 114| 3| 2018-07-01 15:28:17| 2019-02-22 13:35:31|
|[mpb](https://github.com/vbauerster/mpb) | 419| 8| 2016-12-14 19:56:29| 2019-02-26 17:24:45|
|[progressbar](https://github.com/schollz/progressbar) | 494| 13| 2017-10-27 02:28:10| 2019-01-31 01:37:54|
|[simpletable](https://github.com/alexeyco/simpletable) | 119| 2| 2017-03-29 15:27:23| 2019-02-23 00:50:45|
|[tabular](https://github.com/InVisionApp/tabular) | 23| 3| 2018-04-24 05:17:03| 2018-05-15 03:04:57|
|[termbox-go](https://github.com/nsf/termbox-go) | 3286| 103| 2012-01-13 05:03:03| 2019-02-13 09:09:28|
|[termdash](https://github.com/mum4k/termdash) | 78| 6| 2018-03-24 20:01:49| 2019-02-28 13:51:06|
|[termtables](https://github.com/apcera/termtables) | 203| 68| 2012-12-06 08:22:52| 2017-10-30 17:24:27|
|[termui](https://github.com/gizak/termui) | 8353| 275| 2015-02-03 22:09:27| 2019-03-02 06:13:46|
|[uilive](https://github.com/gosuri/uilive) | 723| 11| 2015-11-16 14:13:10| 2018-12-04 09:38:28|
|[uiprogress](https://github.com/gosuri/uiprogress) | 1228| 27| 2015-11-17 08:59:24| 2018-10-12 09:44:27|
|[uitable](https://github.com/gosuri/uitable) | 464| 15| 2015-11-14 05:59:21| 2017-08-23 15:41:33|

## Configuration
        
*Libraries for configuration parsing.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[config](https://github.com/olebedev/config) | 191| 8| 2014-04-21 23:09:39| 2018-09-10 23:57:18|
|[configure](https://github.com/paked/configure) | 43| 3| 2015-06-14 15:46:56| 2019-02-18 22:01:49|
|[confita](https://github.com/heetch/confita) | 208| 24| 2017-12-21 18:49:18| 2019-02-14 23:22:33|
|[conflate](https://github.com/the4thamigo-uk/conflate) | 3| 0| 2018-02-02 03:06:15| 2019-02-09 18:40:00|
|[env](https://github.com/caarlos0/env) | 697| 16| 2015-07-28 10:14:37| 2019-02-18 04:15:59|
|[envcfg](https://github.com/tomazk/envcfg) | 89| 1| 2014-11-29 19:43:53| 2017-06-19 23:53:22|
|[envconf](https://github.com/ian-kent/envconf) | 7| 1| 2014-10-26 20:12:26| 2014-10-26 20:12:40|
|[envconfig](https://github.com/vrischmann/envconfig) | 135| 4| 2015-04-22 07:37:17| 2018-12-28 23:48:21|
|[envh](https://github.com/antham/envh) | 91| 3| 2017-01-12 19:25:48| 2019-03-02 05:12:10|
|[gcfg](https://github.com/go-gcfg/gcfg) | 102| 4| 2015-08-17 22:40:55| 2018-05-18 05:53:51|
|[go-up](https://github.com/ufoscout/go-up) | 19| 1| 2018-02-18 17:50:00| 2018-07-19 15:23:56|
|[goconfig](https://github.com/crgimenes/goconfig) | 91| 11| 2016-12-18 19:22:41| 2019-02-06 17:17:32|
|[godotenv](https://github.com/joho/godotenv) | 1741| 29| 2013-07-30 15:45:19| 2019-02-26 16:05:57|
|[gofigure](https://github.com/ian-kent/gofigure) | 56| 6| 2014-11-25 08:12:40| 2017-05-03 03:22:48|
|[config](https://github.com/gookit/config) | 42| 3| 2018-07-07 16:11:39| 2019-01-19 09:54:56|
|[hjson-go](https://github.com/hjson/hjson-go) | 162| 7| 2016-08-06 06:59:18| 2019-02-09 10:37:18|
|[ingo](https://github.com/schachmat/ingo) | 22| 1| 2016-02-08 06:57:40| 2017-04-03 09:15:10|
|[ini](https://github.com/go-ini/ini) | 1324| 62| 2014-12-18 15:36:37| 2019-02-18 03:54:22|
|[config](https://github.com/joshbetz/config) | 193| 3| 2017-04-03 02:37:05| 2017-08-11 11:04:14|
|[envconfig](https://github.com/kelseyhightower/envconfig) | 2119| 37| 2013-11-07 01:01:55| 2019-02-17 16:18:49|
|[mini](https://github.com/sasbury/mini) | 19| 2| 2015-04-30 07:52:36| 2018-12-27 07:28:05|
|[sprbox](https://github.com/oblq/sprbox) | 3| 2| 2018-07-17 08:27:35| 2018-10-31 22:39:30|
|[store](https://github.com/tucnak/store) | 239| 5| 2015-10-04 03:17:28| 2017-09-05 19:38:35|
|[viper](https://github.com/spf13/viper) | 7652| 165| 2014-04-02 22:33:33| 2019-03-01 21:28:58|
|[xdg](https://github.com/OpenPeeDeeP/xdg) | 18| 3| 2017-07-20 23:58:29| 2019-02-21 08:38:40|

## Continuous Integration
        
*Tools for help with continuous integration.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[drone](https://github.com/drone/drone) | 17332| 584| 2014-02-07 15:54:44| 2019-03-01 22:27:32|
|[duci](https://github.com/duck8823/duci) | 30| 3| 2018-04-01 09:51:02| 2019-03-02 09:22:43|
|[gomason](https://github.com/nikogura/gomason) | 20| 0| 2017-11-18 08:59:11| 2019-02-09 04:31:35|
|[goveralls](https://github.com/mattn/goveralls) | 540| 16| 2013-04-17 18:58:40| 2019-02-03 00:53:12|
|[overalls](https://github.com/go-playground/overalls) | 94| 4| 2015-07-30 19:30:11| 2018-08-26 00:45:52|
|[roveralls](https://github.com/lawrencewoodman/roveralls) | 11| 1| 2016-06-26 15:45:32| 2017-11-20 03:39:13|

## CSS Preprocessors
        
*Libraries for preprocessing CSS files.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[gcss](https://github.com/yosssi/gcss) | 408| 16| 2014-09-04 22:38:20| 2014-10-12 22:07:10|
|[go-libsass](https://github.com/wellington/go-libsass) | 116| 4| 2015-04-19 23:09:47| 2019-02-12 20:41:27|

## Data Structures
        
*Generic datastructures and algorithms in Go.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[algorithms](https://github.com/shady831213/algorithms) | 154| 6| 2018-01-31 17:27:56| 2019-01-08 14:12:30|
|[binpacker](https://github.com/zhuangsirui/binpacker) | 110| 10| 2016-02-02 18:06:11| 2018-06-17 06:52:04|
|[bit](https://github.com/yourbasic/bit) | 29| 3| 2017-05-04 03:05:35| 2018-03-13 15:45:26|
|[bitset](https://github.com/willf/bitset) | 437| 27| 2011-05-11 11:33:44| 2019-03-01 05:25:26|
|[bloom](https://github.com/zentures/bloom) | 125| 8| 2013-09-03 10:27:35| 2018-04-16 15:52:10|
|[bloom](https://github.com/yourbasic/bloom) | 13| 1| 2017-05-07 03:57:47| 2017-06-20 01:00:50|
|[BoomFilters](https://github.com/tylertreat/BoomFilters) | 1080| 36| 2015-02-06 10:01:26| 2018-10-29 03:28:14|
|[concurrent-writer](https://github.com/free/concurrent-writer) | 15| 4| 2017-09-18 23:29:59| 2017-11-18 05:28:32|
|[conjungo](https://github.com/InVisionApp/conjungo) | 57| 14| 2016-12-30 07:50:38| 2019-02-12 21:44:17|
|[count-min-log](https://github.com/seiflotfy/count-min-log) | 40| 4| 2015-08-17 06:31:36| 2017-02-12 21:09:21|
|[cuckoofilter](https://github.com/seiflotfy/cuckoofilter) | 443| 17| 2015-06-29 07:22:09| 2019-02-22 16:40:01|
|[deque](https://github.com/edwingeng/deque) | 1| 1| 2019-02-01 11:32:28| 2019-02-13 10:15:57|
|[deque](https://github.com/gammazero/deque) | 32| 4| 2018-04-24 10:57:55| 2019-01-31 03:14:01|
|[encoding](https://github.com/zentures/encoding) | 92| 7| 2013-09-21 03:29:57| 2017-12-24 06:15:28|
|[go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) | 62| 5| 2016-04-01 09:40:40| 2018-12-06 19:29:01|
|[go-datastructures](https://github.com/Workiva/go-datastructures) | 4776| 292| 2014-10-29 21:55:17| 2018-11-09 01:26:27|
|[go-ef](https://github.com/amallia/go-ef) | 8| 1| 2017-09-22 09:47:16| 2017-09-26 04:07:11|
|[go-geoindex](https://github.com/hailocab/go-geoindex) | 297| 67| 2015-01-22 20:26:17| 2018-02-21 05:58:39|
|[go-mcache](https://github.com/OrlovEvgeny/go-mcache) | 23| 2| 2018-04-15 07:31:21| 2018-11-17 03:36:37|
|[go-rquad](https://github.com/arl/go-rquad) | 94| 3| 2016-09-13 05:46:37| 2018-06-18 03:19:46|
|[goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) | 2| 1| 2019-01-11 05:21:23| 2019-01-31 05:44:05|
|[gods](https://github.com/emirpasic/gods) | 5340| 254| 2015-03-04 22:19:52| 2019-01-24 20:07:05|
|[golang-set](https://github.com/deckarep/golang-set) | 974| 27| 2013-07-04 05:52:01| 2018-09-27 23:06:50|
|[goset](https://github.com/zoumo/goset) | 12| 1| 2017-08-25 17:21:30| 2017-08-25 18:23:15|
|[goskiplist](https://github.com/ryszard/goskiplist) | 178| 12| 2012-05-09 13:44:59| 2017-02-24 16:36:21|
|[gota](https://github.com/go-gota/gota) | 734| 57| 2016-02-07 01:23:25| 2019-02-27 00:15:43|
|[hide](https://github.com/emvi/hide) | 2| 2| 2019-01-16 21:54:17| 2019-02-22 23:06:34|
|[hilbert](https://github.com/google/hilbert) | 171| 19| 2015-08-06 23:50:00| 2018-11-22 14:15:33|
|[hyperloglog](https://github.com/axiomhq/hyperloglog) | 634| 18| 2017-06-18 19:18:12| 2018-12-23 19:14:21|
|[levenshtein](https://github.com/agext/levenshtein) | 25| 1| 2016-04-08 08:14:31| 2019-02-23 03:07:15|
|[levenshtein](https://github.com/agnivade/levenshtein) | 42| 3| 2014-07-30 22:03:55| 2019-02-28 01:56:32|
|[mafsa](https://github.com/smartystreets/mafsa) | 272| 18| 2014-11-08 05:07:57| 2017-03-07 23:18:16|
|[merkletree](https://github.com/cbergoon/merkletree) | 119| 4| 2017-04-12 10:50:11| 2018-12-26 09:42:24|
|[mspm](https://github.com/BlackRabbitt/mspm) | 4| 3| 2018-05-18 02:59:44| 2018-05-19 14:36:38|
|[null](https://github.com/emvi/null) | 2| 2| 2018-07-05 05:18:45| 2019-02-22 23:01:00|
|[pipeline](https://github.com/hyfather/pipeline) | 9| 1| 2018-04-25 08:11:36| 2018-08-31 11:09:33|
|[ring](https://github.com/TheTannerRyan/ring) | 83| 1| 2019-01-27 12:02:20| 2019-02-19 10:52:44|
|[roaring](https://github.com/RoaringBitmap/roaring) | 564| 34| 2014-07-11 04:14:34| 2019-01-10 22:41:21|
|[set](https://github.com/StudioSol/set) | 5| 17| 2018-07-21 05:53:37| 2018-10-10 02:02:01|
|[skiplist](https://github.com/MauriceGit/skiplist) | 80| 5| 2018-06-24 00:01:51| 2018-12-08 17:30:32|
|[skiplist](https://github.com/gansidui/skiplist) | 58| 7| 2014-11-19 00:29:53| 2014-11-21 13:13:52|
|[timedmap](https://github.com/zekroTJA/timedmap) | 0| 1| 2019-01-30 20:55:37| 2019-02-20 22:48:16|
|[trie](https://github.com/derekparker/trie) | 356| 20| 2014-03-07 06:01:49| 2018-05-11 06:29:42|
|[ttlcache](https://github.com/ReneKroon/ttlcache) | 78| 7| 2014-12-13 09:55:40| 2019-02-25 21:04:39|
|[bloom](https://github.com/willf/bloom) | 584| 30| 2011-05-21 22:18:41| 2019-03-01 05:23:55|

## Database
        
*SQL query builder, libraries for building and using SQL.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[badger](https://github.com/dgraph-io/badger) | 5374| 219| 2017-01-26 13:09:49| 2019-03-02 03:36:11|
|[bigcache](https://github.com/allegro/bigcache) | 1878| 64| 2016-03-23 15:18:52| 2019-02-18 14:46:06|
|[bolt](https://github.com/boltdb/bolt) | 9486| 337| 2013-12-21 02:26:14| 2018-03-03 02:00:53|
|[buntdb](https://github.com/tidwall/buntdb) | 2262| 93| 2016-07-20 06:11:40| 2019-02-03 03:56:57|
|[cache2go](https://github.com/muesli/cache2go) | 811| 57| 2013-11-11 11:45:02| 2018-10-29 06:37:51|
|[clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) | 25| 5| 2017-12-18 15:48:07| 2018-01-23 06:02:54|
|[cockroach](https://github.com/cockroachdb/cockroach) | 15515| 693| 2014-02-06 08:18:47| 2019-03-02 13:28:37|
|[couchcache](https://github.com/codingsince1985/couchcache) | 39| 3| 2015-04-05 15:13:05| 2017-10-25 09:28:55|
|[CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) | 481| 45| 2018-04-11 17:52:58| 2019-03-01 17:54:27|
|[dgraph](https://github.com/dgraph-io/dgraph) | 8682| 329| 2015-08-25 15:15:56| 2019-03-02 12:19:32|
|[diskv](https://github.com/peterbourgon/diskv) | 695| 29| 2012-03-22 00:44:32| 2018-09-04 20:15:47|
|[eliasdb](https://github.com/krotik/eliasdb) | 503| 23| 2016-08-13 21:53:28| 2019-02-16 01:18:02|
|[fastcache](https://github.com/VictoriaMetrics/fastcache) | 347| 12| 2018-11-23 06:50:13| 2019-02-25 05:17:14|
|[gcache](https://github.com/bluele/gcache) | 664| 29| 2015-01-25 02:17:07| 2019-03-01 12:41:16|
|[go-cache](https://github.com/patrickmn/go-cache) | 2415| 84| 2012-01-02 21:07:13| 2019-02-27 09:24:53|
|[goleveldb](https://github.com/syndtr/goleveldb) | 2787| 150| 2013-01-23 12:08:58| 2019-02-26 23:50:50|
|[gorocksdb](https://github.com/kapitan-k/gorocksdb) | 6| 1| 2017-12-28 18:28:48| 2018-01-11 04:23:36|
|[groupcache](https://github.com/golang/groupcache) | 7094| 443| 2013-07-23 05:55:07| 2019-02-13 09:06:53|
|[influxdb](https://github.com/influxdata/influxdb) | 15690| 750| 2013-09-26 22:31:10| 2019-03-02 11:25:15|
|[ledisdb](https://github.com/siddontang/ledisdb) | 2886| 184| 2014-04-30 08:43:09| 2019-02-05 13:49:34|
|[levigo](https://github.com/jmhodges/levigo) | 345| 24| 2012-01-17 16:17:54| 2019-02-28 18:43:15|
|[moss](https://github.com/couchbase/moss) | 682| 79| 2016-02-07 04:27:22| 2019-02-22 00:32:47|
|[piladb](https://github.com/fern4lvarez/piladb) | 167| 8| 2015-09-09 07:12:22| 2018-04-08 00:09:23|
|[prometheus](https://github.com/prometheus/prometheus) | 22352| 1004| 2012-11-24 19:14:12| 2019-03-02 11:28:07|
|[pudge](https://github.com/recoilme/pudge) | 181| 6| 2018-11-20 18:11:53| 2019-02-27 15:03:55|
|[rqlite](https://github.com/rqlite/rqlite) | 4253| 182| 2014-08-23 12:31:18| 2019-02-27 01:18:36|
|[golang-scribble](https://github.com/nanobox-io/golang-scribble) | 32| 3| 2018-06-22 06:13:33| 2018-10-18 23:48:25|
|[slowpoke](https://github.com/recoilme/slowpoke) | 74| 6| 2018-02-19 17:22:37| 2018-12-24 23:55:56|
|[tempdb](https://github.com/rafaeljesus/tempdb) | 13| 3| 2017-03-18 02:03:42| 2018-02-15 03:03:13|
|[tidb](https://github.com/pingcap/tidb) | 17531| 1198| 2015-09-06 12:01:52| 2019-03-02 13:16:30|
|[tiedot](https://github.com/HouzuoGuo/tiedot) | 2306| 161| 2013-05-26 18:03:49| 2019-01-26 01:43:32|
|[vasto](https://github.com/chrislusf/vasto) | 120| 16| 2018-01-16 13:16:57| 2018-10-23 17:02:13|
|[darwin](https://github.com/GuiaBolso/darwin) | 77| 2| 2016-04-05 23:57:59| 2018-11-28 00:12:25|
|[go-fixtures](https://github.com/RichardKnop/go-fixtures) | 17| 1| 2015-12-24 19:27:45| 2018-11-01 11:56:57|
|[go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) | 19| 1| 2018-08-11 15:00:13| 2018-12-23 00:30:09|
|[goose](https://github.com/steinbacher/goose) | 103| 4| 2016-03-05 04:33:14| 2016-09-24 15:03:49|
|[gormigrate](https://github.com/go-gormigrate/gormigrate) | 256| 3| 2016-08-31 19:46:23| 2019-02-04 03:31:20|
|[migrate](https://github.com/golang-migrate/migrate) | 1690| 37| 2018-01-19 17:30:58| 2019-03-02 10:52:21|
|[pravasan](https://github.com/pravasan/pravasan) | 23| 6| 2015-01-04 01:11:05| 2018-12-20 09:56:10|
|[pop](https://github.com/gobuffalo/pop) | 585| 24| 2018-02-08 05:13:46| 2019-03-01 00:49:58|
|[sql-migrate](https://github.com/rubenv/sql-migrate) | 1243| 29| 2014-09-09 15:31:41| 2019-02-15 05:27:34|
|[chproxy](https://github.com/Vertamedia/chproxy) | 240| 20| 2017-09-18 21:09:23| 2019-02-03 05:00:26|
|[clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) | 97| 6| 2017-04-29 18:38:41| 2019-02-20 03:05:55|
|[dbbench](https://github.com/sj14/dbbench) | 23| 2| 2018-11-24 21:21:18| 2019-02-02 16:57:42|
|[go-mysql](https://github.com/siddontang/go-mysql) | 1564| 133| 2014-02-21 09:56:45| 2019-02-28 17:00:56|
|[go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) | 2032| 150| 2015-01-15 17:54:18| 2019-02-27 10:54:41|
|[kingshard](https://github.com/flike/kingshard) | 4179| 388| 2015-07-04 10:22:32| 2019-02-19 18:33:01|
|[myreplication](https://github.com/2tvenom/myreplication) | 129| 17| 2015-02-05 04:59:49| 2018-10-05 15:34:57|
|[octillery](https://github.com/knocknote/octillery) | 42| 14| 2018-11-26 18:39:35| 2019-01-18 17:14:32|
|[orchestrator](https://github.com/github/orchestrator) | 2588| 181| 2016-11-30 21:44:24| 2019-03-01 18:17:54|
|[pgweb](https://github.com/sosedoff/pgweb) | 5750| 152| 2014-10-09 09:41:32| 2019-02-26 19:14:40|
|[prep](https://github.com/hexdigest/prep) | 22| 2| 2017-12-12 07:47:38| 2017-12-20 01:35:51|
|[prest](https://github.com/prest/prest) | 1963| 82| 2016-11-22 13:17:05| 2018-09-04 21:21:35|
|[rwdb](https://github.com/andizzle/rwdb) | 10| 2| 2017-10-04 11:55:29| 2017-11-08 17:10:17|
|[vitess](https://github.com/vitessio/vitess) | 7600| 464| 2013-06-28 05:20:28| 2019-03-02 11:09:12|
|[dotsql](https://github.com/gchaincl/dotsql) | 406| 21| 2014-11-20 20:14:39| 2018-10-22 14:31:23|
|[gendry](https://github.com/didi/gendry) | 570| 46| 2017-12-01 16:15:43| 2018-12-29 12:14:00|
|[godbal](https://github.com/xujiajun/godbal) | 48| 3| 2018-02-28 13:47:42| 2019-01-30 13:57:00|
|[goqu](https://github.com/doug-martin/goqu) | 491| 25| 2015-02-21 09:06:00| 2019-02-13 23:37:43|
|[igor](https://github.com/galeone/igor) | 73| 6| 2016-03-10 22:45:08| 2018-07-01 17:41:38|
|[ormlite](https://github.com/pupizoid/ormlite) | 0| 1| 2018-06-28 21:42:09| 2019-02-15 22:37:03|
|[ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) | 296| 25| 2015-12-11 06:39:26| 2018-08-23 14:45:48|
|[scaneo](https://github.com/variadico/scaneo) | 141| 8| 2015-04-30 08:36:27| 2018-04-03 20:49:37|
|[sqrl](https://github.com/elgris/sqrl) | 140| 5| 2014-06-25 18:03:06| 2019-01-30 01:38:07|
|[squirrel](https://github.com/Masterminds/squirrel) | 1930| 32| 2014-01-18 13:29:58| 2019-02-27 02:57:20|
|[xo](https://github.com/xo/xo) | 1944| 70| 2016-02-05 18:22:20| 2019-01-20 20:32:12|

## Database Drivers
        
*Libraries for connecting and operating databases.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|

### Relational Databases
        
|[calcite-avatica-go](https://github.com/apache/calcite-avatica-go) | 22| 16| 2017-08-08 15:00:08| 2019-01-08 07:50:24|
|[bgc](https://github.com/viant/bgc) | 10| 9| 2016-06-14 04:24:26| 2019-01-30 03:42:30|
|[firebirdsql](https://github.com/nakagami/firebirdsql) | 94| 14| 2013-08-27 21:09:14| 2019-03-02 10:01:54|
|[go-adodb](https://github.com/mattn/go-adodb) | 84| 12| 2011-11-14 12:32:50| 2019-02-26 18:19:11|
|[go-mssqldb](https://github.com/denisenkom/go-mssqldb) | 908| 59| 2013-12-16 08:10:47| 2019-02-19 13:38:43|
|[go-oci8](https://github.com/mattn/go-oci8) | 360| 34| 2012-02-29 20:19:16| 2019-02-05 18:29:31|
|[mysql](https://github.com/go-sql-driver/mysql) | 7034| 391| 2012-12-10 04:33:55| 2019-02-17 15:26:59|
|[go-sqlite3](https://github.com/mattn/go-sqlite3) | 3084| 125| 2011-11-11 20:36:50| 2019-02-18 01:40:34|
|[gofreetds](https://github.com/minus5/gofreetds) | 84| 21| 2012-12-07 01:29:26| 2019-02-20 00:37:09|
|[goracle](https://github.com/go-goracle/goracle) | 182| 28| 2015-03-25 13:58:16| 2019-02-20 13:20:21|
|[pgx](https://github.com/jackc/pgx) | 1645| 64| 2013-03-31 03:06:26| 2019-03-01 00:05:55|
|[pq](https://github.com/lib/pq) | 4669| 158| 2012-03-13 02:50:22| 2019-02-09 01:37:56|

### NoSQL Databases
        
|[aerospike-client-go](https://github.com/aerospike/aerospike-client-go) | 288| 39| 2014-07-26 10:56:21| 2019-02-21 21:56:25|
|[arangolite](https://github.com/solher/arangolite) | 64| 5| 2015-10-05 01:27:34| 2017-09-10 17:40:18|
|[asc](https://github.com/viant/asc) | 4| 10| 2016-06-14 04:22:31| 2018-11-13 12:21:38|
|[dynago](https://github.com/underarmour/dynago) | 63| 131| 2015-05-18 23:40:20| 2017-08-08 06:07:05|
|[goforestdb](https://github.com/couchbase/goforestdb) | 29| 44| 2014-05-14 23:36:12| 2016-12-16 06:01:01|
|[go-couchbase](https://github.com/couchbase/go-couchbase) | 285| 26| 2012-01-20 06:52:08| 2019-03-02 05:31:26|
|[go-couchdb](https://github.com/fjl/go-couchdb) | 51| 6| 2013-10-28 09:08:16| 2018-09-05 10:07:28|
|[go-pilosa](https://github.com/pilosa/go-pilosa) | 27| 13| 2016-10-01 05:37:10| 2019-03-01 23:01:52|
|[go-rejson](https://github.com/nitishm/go-rejson) | 67| 4| 2018-04-23 08:51:05| 2019-01-28 01:33:17|
|[gocb](https://github.com/couchbase/gocb) | 281| 71| 2015-01-16 04:01:32| 2019-02-19 12:00:42|
|[godscache](https://github.com/defcronyke/godscache) | 4| 2| 2018-05-09 04:19:39| 2019-02-08 15:04:54|
|[gomemcache](https://github.com/bradfitz/gomemcache) | 1011| 51| 2011-06-29 03:29:12| 2019-02-22 05:56:08|
|[rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1418| 45| 2013-09-12 21:56:27| 2019-01-22 02:10:18|
|[goriak](https://github.com/zegl/goriak) | 25| 2| 2016-10-06 00:48:17| 2018-12-28 07:06:55|
|[mgo](https://github.com/globalsign/mgo) | 1423| 77| 2017-04-13 19:14:04| 2019-03-01 00:42:21|
|[mongo-go-driver](https://github.com/mongodb/mongo-go-driver) | 2019| 136| 2017-02-09 01:18:02| 2019-03-01 06:50:01|
|[neo4j](https://github.com/cihangir/neo4j) | 24| 3| 2013-05-18 16:54:01| 2015-04-03 01:38:48|
|[Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) | 72| 6| 2011-06-05 00:08:35| 2018-06-20 20:15:38|
|[neoism](https://github.com/jmcvetta/neoism) | 350| 23| 2012-07-12 15:42:33| 2019-02-19 09:47:26|
|[redigo](https://github.com/gomodule/redigo) | 5538| 280| 2012-04-14 12:31:58| 2019-03-01 19:22:27|
|[redis](https://github.com/go-redis/redis) | 5135| 200| 2012-07-25 21:01:39| 2019-03-01 17:23:50|
|[redeo](https://github.com/bsm/redeo) | 240| 24| 2014-03-06 16:46:18| 2018-11-29 20:16:24|
|[xredis](https://github.com/shomali11/xredis) | 9| 1| 2017-06-14 08:19:26| 2018-06-07 08:59:07|
|[bleve](https://github.com/blevesearch/bleve) | 5000| 231| 2014-04-18 05:02:18| 2019-02-28 00:32:59|
|[elastic](https://github.com/olivere/elastic) | 3507| 160| 2012-12-07 01:15:33| 2019-02-26 18:19:22|
|[elasticsql](https://github.com/cch123/elasticsql) | 291| 20| 2016-08-24 15:29:43| 2019-02-22 15:10:51|
|[elastigo](https://github.com/mattbaird/elastigo) | 937| 49| 2012-10-12 12:19:59| 2019-02-06 02:17:02|
|[goes](https://github.com/OwnLocal/goes) | 25| 34| 2015-12-29 02:52:03| 2017-03-03 09:06:24|
|[riot](https://github.com/go-ego/riot) | 4295| 171| 2017-06-21 22:17:59| 2019-01-31 22:23:29|
|[skizze](https://github.com/seiflotfy/skizze) | 59| 6| 2016-01-17 20:10:40| 2016-05-10 02:15:30|
|[cachego](https://github.com/fabiorphp/cachego) | 103| 6| 2016-10-06 02:10:03| 2018-05-20 20:11:29|
|[cayley](https://github.com/cayleygraph/cayley) | 12246| 632| 2014-06-06 02:49:41| 2019-02-22 04:17:08|
|[dsc](https://github.com/viant/dsc) | 6| 17| 2016-06-14 04:18:10| 2019-02-24 07:34:49|
|[gokv](https://github.com/philippgille/gokv) | 44| 4| 2018-10-09 02:55:22| 2019-02-25 03:40:57|

## Date and Time
        
*Libraries for working with dates and times.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[carbon](https://github.com/uniplaces/carbon) | 305| 38| 2016-08-03 18:55:52| 2019-01-16 20:02:42|
|[date](https://github.com/rickb777/date) | 23| 2| 2015-11-24 06:58:07| 2018-11-27 22:32:14|
|[dateparse](https://github.com/araddon/dateparse) | 800| 14| 2014-04-21 10:55:48| 2019-02-23 09:01:43|
|[durafmt](https://github.com/hako/durafmt) | 208| 4| 2016-05-21 05:49:33| 2018-09-04 07:44:38|
|[feiertage](https://github.com/wlbr/feiertage) | 17| 1| 2015-11-04 22:19:27| 2019-02-14 08:19:35|
|[go-persian-calendar](https://github.com/yaa110/go-persian-calendar) | 48| 3| 2016-02-01 02:40:23| 2019-01-28 20:52:10|
|[go-sunrise](https://github.com/nathan-osman/go-sunrise) | 10| 3| 2017-06-16 04:49:41| 2017-11-22 04:50:07|
|[goweek](https://github.com/grsmv/goweek) | 18| 3| 2015-11-12 22:11:46| 2017-01-04 04:24:26|
|[iso8601](https://github.com/relvacode/iso8601) | 60| 2| 2017-04-25 23:54:18| 2018-12-21 23:13:36|
|[kair](https://github.com/GuilhermeCaruso/kair) | 9| 1| 2018-10-03 23:44:07| 2019-02-24 06:06:10|
|[now](https://github.com/jinzhu/now) | 1978| 51| 2013-11-18 18:55:30| 2019-02-17 08:24:44|
|[nulltime](https://github.com/kirillDanshin/nulltime) | 9| 1| 2016-03-06 09:44:58| 2017-03-22 12:30:28|
|[strftime](https://github.com/awoodbeck/strftime) | 5| 1| 2018-02-10 08:35:46| 2018-02-21 23:59:14|
|[timespan](https://github.com/SaidinWoT/timespan) | 60| 5| 2014-10-07 11:57:02| 2019-02-01 16:55:52|
|[timeutil](https://github.com/leekchan/timeutil) | 156| 6| 2015-08-02 09:32:06| 2019-02-03 21:14:43|
|[tuesday](https://github.com/osteele/tuesday) | 7| 1| 2017-08-11 04:46:26| 2017-08-30 21:28:26|

## Distributed Systems
        
*Packages that help with building Distributed Systems.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[celeriac.v1](https://github.com/svcavallar/celeriac.v1) | 47| 3| 2015-10-10 15:27:33| 2018-12-08 13:38:42|
|[consistent](https://github.com/buraksezer/consistent) | 151| 9| 2018-03-25 23:38:27| 2019-02-23 11:07:49|
|[digota](https://github.com/digota/digota) | 269| 24| 2017-08-14 20:01:37| 2018-10-15 06:52:00|
|[doublejump](https://github.com/edwingeng/doublejump) | 26| 3| 2018-06-27 00:04:50| 2019-01-02 18:37:11|
|[drmaa](https://github.com/dgruber/drmaa) | 23| 2| 2013-03-17 20:58:02| 2018-05-14 14:05:07|
|[dynatomic](https://github.com/tylfin/dynatomic) | 5| 0| 2019-02-09 01:45:14| 2019-02-21 20:40:01|
|[emitter](https://github.com/emitter-io/emitter) | 1651| 80| 2016-10-29 16:52:21| 2019-03-02 10:58:02|
|[flowgraph](https://github.com/vectaport/flowgraph) | 7| 1| 2018-08-30 05:45:26| 2019-03-02 02:28:02|
|[gleam](https://github.com/chrislusf/gleam) | 1829| 137| 2016-08-26 16:44:48| 2018-12-11 22:58:01|
|[glow](https://github.com/chrislusf/glow) | 2351| 141| 2015-06-14 08:33:48| 2018-11-02 14:09:14|
|[go-health](https://github.com/InVisionApp/go-health) | 437| 20| 2017-11-30 05:00:07| 2018-12-05 02:25:05|
|[go-jump](https://github.com/dgryski/go-jump) | 231| 13| 2014-06-16 06:12:04| 2018-02-12 22:36:50|
|[kit](https://github.com/go-kit/kit) | 12764| 641| 2015-02-03 08:01:19| 2019-02-25 09:14:30|
|[gorpc](https://github.com/valyala/gorpc) | 520| 26| 2014-11-21 01:02:37| 2017-04-07 09:55:22|
|[grpc-go](https://github.com/grpc/grpc-go) | 7575| 437| 2014-12-09 02:59:34| 2019-03-02 08:11:57|
|[hprose-golang](https://github.com/hprose/hprose-golang) | 918| 89| 2014-02-14 11:16:43| 2018-05-17 23:21:26|
|[jaeger](https://github.com/jaegertracing/jaeger) | 7328| 277| 2016-04-16 02:49:02| 2019-03-02 07:10:15|
|[jsonrpc](https://github.com/osamingo/jsonrpc) | 103| 6| 2016-10-28 21:36:59| 2019-01-04 00:39:07|
|[jsonrpc](https://github.com/ybbus/jsonrpc) | 78| 6| 2016-11-10 19:27:55| 2018-12-01 14:59:14|
|[krakend](https://github.com/devopsfaith/krakend) | 1273| 67| 2016-11-05 02:37:13| 2019-03-01 22:59:56|
|[micro](https://github.com/micro/micro) | 5582| 289| 2015-01-17 06:35:14| 2019-02-28 19:16:38|
|[gnatsd](https://github.com/nats-io/gnatsd) | 5307| 326| 2012-10-30 00:12:24| 2019-03-01 05:14:14|
|[outboxer](https://github.com/italolelis/outboxer) | 0| 0| 2019-02-01 17:50:13| 2019-02-22 05:34:43|
|[raft](https://github.com/hashicorp/raft) | 2483| 176| 2013-11-05 08:41:20| 2019-03-01 20:30:53|
|[etcd](https://github.com/etcd-io/etcd) | 23220| 1226| 2013-07-07 05:57:21| 2019-03-02 10:04:53|
|[redis-lock](https://github.com/bsm/redis-lock) | 108| 6| 2014-10-10 22:22:48| 2018-11-29 19:41:28|
|[ringpop-go](https://github.com/uber/ringpop-go) | 530| 2310| 2015-06-06 06:48:53| 2019-02-19 11:59:14|
|[rpcx](https://github.com/smallnest/rpcx) | 3233| 270| 2016-05-18 17:34:05| 2019-03-01 17:21:34|
|[sleuth](https://github.com/ursiform/sleuth) | 292| 7| 2016-04-23 22:21:41| 2018-03-21 23:59:30|
|[tendermint](https://github.com/tendermint/tendermint) | 2710| 241| 2014-05-15 07:21:35| 2019-03-02 14:06:58|
|[torrent](https://github.com/anacrolix/torrent) | 2608| 121| 2015-01-09 05:10:42| 2019-03-01 08:18:58|
|[go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) | 355| 18| 2015-10-09 03:44:47| 2019-02-23 23:51:07|

## Email
        
*Libraries and tools that implement email creation and sending.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[douceur](https://github.com/aymerick/douceur) | 141| 4| 2015-04-09 18:21:26| 2018-03-23 05:03:08|
|[email](https://github.com/jordan-wright/email) | 1007| 35| 2013-12-13 04:11:59| 2019-02-18 10:44:55|
|[go-dkim](https://github.com/toorop/go-dkim) | 47| 3| 2015-04-29 23:38:27| 2018-02-24 05:40:33|
|[go-imap](https://github.com/emersion/go-imap) | 625| 31| 2016-04-27 01:59:18| 2019-03-02 03:04:29|
|[go-message](https://github.com/emersion/go-message) | 79| 6| 2016-12-31 17:31:52| 2019-01-09 06:54:23|
|[gomail](https://github.com/go-gomail/gomail) | 2134| 65| 2014-10-15 23:47:48| 2019-02-13 09:16:07|
|[hectane](https://github.com/hectane/hectane) | 160| 12| 2015-08-28 09:36:47| 2018-04-15 04:06:48|
|[hermes](https://github.com/matcornic/hermes) | 1474| 23| 2017-03-26 02:25:36| 2019-02-27 02:41:35|
|[MailHog](https://github.com/mailhog/MailHog) | 4495| 125| 2014-04-17 06:28:49| 2019-03-02 01:48:44|
|[sendgrid-go](https://github.com/sendgrid/sendgrid-go) | 464| 181| 2013-09-12 11:31:13| 2019-02-23 06:04:09|
|[smtp](https://github.com/mailhog/smtp) | 49| 9| 2014-12-25 00:13:49| 2018-06-13 00:19:22|

## Embeddable Scripting Languages
        
*Embedding other languages inside your go code.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[agora](https://github.com/mna/agora) | 318| 22| 2013-06-19 10:16:30| 2015-01-26 01:12:45|
|[anko](https://github.com/mattn/anko) | 840| 47| 2014-03-28 15:29:40| 2019-02-26 19:35:12|
|[binder](https://github.com/alexeyco/binder) | 24| 2| 2017-04-03 01:14:52| 2018-07-30 06:00:27|
|[expr](https://github.com/antonmedv/expr) | 295| 16| 2018-07-14 23:57:34| 2019-01-19 17:14:27|
|[gisp](https://github.com/jcla1/gisp) | 415| 20| 2014-01-11 22:05:43| 2017-08-25 21:48:45|
|[go-duktape](https://github.com/olebedev/go-duktape) | 628| 23| 2015-01-08 13:09:05| 2019-02-14 07:42:58|
|[go-lua](https://github.com/Shopify/go-lua) | 1539| 253| 2013-12-21 01:29:43| 2018-11-07 02:40:36|
|[go-php](https://github.com/deuill/go-php) | 608| 39| 2015-09-18 05:23:52| 2018-10-07 23:22:34|
|[go-python](https://github.com/sbinet/go-python) | 807| 43| 2012-07-09 23:43:31| 2019-01-18 21:19:05|
|[golua](https://github.com/aarzilli/golua) | 422| 35| 2010-12-07 05:39:53| 2018-08-22 18:36:17|
|[gopher-lua](https://github.com/yuin/gopher-lua) | 2604| 136| 2015-02-15 21:23:37| 2019-02-07 22:37:53|
|[gval](https://github.com/PaesslerAG/gval) | 85| 12| 2017-09-27 16:32:49| 2019-02-19 20:43:41|
|[ngaro](https://github.com/db47h/ngaro) | 16| 1| 2016-08-09 23:23:50| 2018-06-03 18:57:43|
|[otto](https://github.com/robertkrimen/otto) | 4403| 187| 2012-10-06 09:48:39| 2019-02-25 15:56:09|
|[purl](https://github.com/ian-kent/purl) | 25| 2| 2014-11-30 03:06:01| 2014-12-08 01:45:34|
|[tengo](https://github.com/d5/tengo) | 1051| 24| 2019-01-09 15:17:17| 2019-03-02 07:56:04|

## Error Handling
        
*Libraries for handling errors.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[errors](https://github.com/pkg/errors) | 4159| 91| 2015-12-27 20:05:38| 2019-02-27 08:00:54|
|[errorx](https://github.com/joomcode/errorx) | 508| 37| 2018-08-17 16:02:10| 2019-02-15 20:27:52|
|[go-multierror](https://github.com/hashicorp/go-multierror) | 558| 81| 2014-12-16 04:12:26| 2018-12-14 18:18:49|
|[tracerr](https://github.com/ztrue/tracerr) | 172| 6| 2019-02-07 02:57:46| 2019-02-16 18:32:28|
|[werr](https://github.com/txgruppi/werr) | 10| 0| 2015-08-04 19:10:13| 2016-03-10 11:37:06|

## Files
        
*Libraries for  handling files and file systems.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[afero](https://github.com/spf13/afero) | 1935| 83| 2014-10-28 22:19:05| 2019-03-02 10:20:19|
|[go-csv-tag](https://github.com/artonge/go-csv-tag) | 39| 1| 2017-06-18 23:31:16| 2018-10-15 18:32:17|
|[go-decent-copy](https://github.com/hugocarreira/go-decent-copy) | 11| 1| 2018-10-16 15:08:24| 2018-10-18 19:24:25|
|[go-gtfs](https://github.com/artonge/go-gtfs) | 11| 2| 2017-07-09 17:30:31| 2019-01-10 21:24:12|
|[notify](https://github.com/rjeczalik/notify) | 450| 26| 2014-09-09 00:09:34| 2019-02-19 09:29:26|
|[opc](https://github.com/qmuntal/opc) | 47| 1| 2018-11-06 22:49:06| 2019-02-26 14:05:35|
|[pdfcpu](https://github.com/hhrutter/pdfcpu) | 796| 22| 2017-06-19 01:27:38| 2019-02-25 04:42:08|
|[skywalker](https://github.com/dixonwille/skywalker) | 41| 3| 2017-08-02 04:08:25| 2017-08-05 04:28:55|
|[tarfs](https://github.com/posener/tarfs) | 31| 2| 2017-03-11 06:13:11| 2017-04-02 16:33:57|

## Financial
        
*Packages for accounting and finance.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[accounting](https://github.com/leekchan/accounting) | 441| 12| 2015-08-10 21:23:56| 2019-01-31 14:20:31|
|[decimal](https://github.com/shopspring/decimal) | 1332| 55| 2015-02-26 04:12:57| 2019-02-19 17:12:29|
|[go-finance](https://github.com/FlashBoys/go-finance) | 538| 28| 2016-02-28 08:37:46| 2018-03-09 10:50:46|
|[go-finance](https://github.com/alpeb/go-finance) | 32| 2| 2017-06-01 23:58:33| 2018-05-06 00:52:06|
|[go-money](https://github.com/Rhymond/go-money) | 551| 15| 2017-03-21 00:23:54| 2019-01-30 18:29:45|
|[ofxgo](https://github.com/aclindsa/ofxgo) | 53| 5| 2015-11-08 21:56:53| 2019-03-02 11:43:10|
|[orderbook](https://github.com/i25959341/orderbook) | 21| 3| 2018-04-25 02:05:26| 2019-03-01 18:49:58|
|[techan](https://github.com/sdcoffey/techan) | 108| 17| 2017-03-08 11:04:08| 2018-12-13 23:35:39|
|[transaction](https://github.com/claygod/transaction) | 44| 8| 2017-10-11 21:50:30| 2018-09-06 02:12:01|
|[vat](https://github.com/dannyvankooten/vat) | 53| 1| 2016-06-19 00:10:09| 2018-09-10 22:01:40|

## Forms
        
*Libraries for working with forms.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[bind](https://github.com/robfig/bind) | 23| 2| 2014-08-06 08:13:10| 2014-08-17 01:03:51|
|[binding](https://github.com/mholt/binding) | 730| 30| 2014-05-21 07:35:14| 2018-03-29 07:47:34|
|[conform](https://github.com/leebenson/conform) | 163| 5| 2016-01-06 02:00:06| 2018-06-16 05:02:23|
|[form](https://github.com/go-playground/form) | 324| 8| 2016-05-26 21:26:40| 2019-02-03 09:53:44|
|[formam](https://github.com/monoculum/formam) | 111| 2| 2014-10-25 08:23:30| 2018-09-01 09:54:15|
|[forms](https://github.com/albrow/forms) | 96| 6| 2014-08-08 00:11:30| 2017-07-02 20:22:45|
|[csrf](https://github.com/gorilla/csrf) | 368| 19| 2015-08-03 08:35:16| 2018-12-20 18:56:45|
|[nosurf](https://github.com/justinas/nosurf) | 919| 35| 2013-08-23 01:47:34| 2019-01-19 00:37:50|

## Functional
        
*Packages to support functional programming in Go.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[fpGo](https://github.com/TeaEntityLab/fpGo) | 69| 1| 2018-05-24 17:08:45| 2018-07-19 13:51:53|
|[fuego](https://github.com/seborama/fuego) | 13| 1| 2018-11-06 06:24:09| 2019-03-02 01:59:16|
|[go-underscore](https://github.com/tobyhede/go-underscore) | 1009| 26| 2014-07-02 18:27:16| 2019-02-15 05:27:45|

## Game Development
        
*Awesome game development libraries.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[engine](https://github.com/azul3d/engine) | 406| 23| 2016-02-29 12:54:44| 2018-06-25 06:16:41|
|[ebiten](https://github.com/hajimehoshi/ebiten) | 1405| 65| 2013-06-16 23:13:01| 2019-03-02 03:44:05|
|[engo](https://github.com/EngoEngine/engo) | 951| 45| 2014-11-12 13:50:03| 2019-03-02 02:04:46|
|[engine](https://github.com/g3n/engine) | 590| 58| 2017-03-08 02:25:09| 2019-02-19 21:16:25|
|[GarageEngine](https://github.com/vova616/GarageEngine) | 303| 31| 2012-08-07 18:55:34| 2013-09-03 19:21:00|
|[glop](https://github.com/runningwild/glop) | 76| 4| 2011-04-21 06:48:34| 2015-09-24 11:00:09|
|[go-astar](https://github.com/beefsack/go-astar) | 307| 10| 2014-05-28 10:00:03| 2018-03-26 13:33:35|
|[go-collada](https://github.com/GlenKelley/go-collada) | 12| 3| 2013-09-19 12:19:15| 2013-09-27 05:39:30|
|[go-sdl2](https://github.com/veandco/go-sdl2) | 1051| 42| 2013-06-06 02:30:03| 2019-02-23 19:28:37|
|[go3d](https://github.com/ungerik/go3d) | 150| 10| 2011-06-27 21:02:26| 2019-03-01 14:15:22|
|[gonet](https://github.com/xtaci/gonet) | 997| 125| 2013-04-11 10:18:23| 2017-05-12 15:31:41|
|[goworld](https://github.com/xiaonanln/goworld) | 946| 99| 2017-06-03 23:02:46| 2019-01-01 21:45:25|
|[leaf](https://github.com/name5566/leaf) | 2714| 302| 2014-08-04 20:40:08| 2018-12-28 22:20:19|
|[nano](https://github.com/lonng/nano) | 844| 53| 2017-08-02 14:05:14| 2019-01-26 17:31:10|
|[oak](https://github.com/oakmound/oak) | 572| 35| 2017-07-16 00:24:27| 2019-02-26 10:52:54|
|[pitaya](https://github.com/topfreegames/pitaya) | 202| 26| 2018-03-20 03:40:36| 2019-03-01 05:29:28|
|[pixel](https://github.com/faiface/pixel) | 2111| 92| 2016-11-19 19:15:34| 2019-02-23 00:13:59|
|[raylib-go](https://github.com/gen2brain/raylib-go) | 340| 18| 2017-01-27 16:31:45| 2018-11-18 04:47:59|
|[termloop](https://github.com/JoelOtter/termloop) | 975| 35| 2015-05-24 01:12:34| 2018-12-02 01:22:10|

## Generation and Generics
        
*Tools to enhance the language with features like generics via code generation.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[efaceconv](https://github.com/t0pep0/efaceconv) | 39| 4| 2016-11-18 19:38:54| 2017-10-12 15:16:32|
|[gen](https://github.com/clipperhouse/gen) | 977| 36| 2013-10-14 04:26:36| 2018-06-12 03:54:36|
|[go-enum](https://github.com/abice/go-enum) | 65| 4| 2017-08-11 06:07:31| 2019-02-15 06:47:19|
|[go-linq](https://github.com/ahmetb/go-linq) | 1661| 72| 2013-12-19 11:05:00| 2018-11-09 16:43:21|
|[goderive](https://github.com/awalterschulze/goderive) | 670| 22| 2017-02-11 05:46:49| 2019-02-24 20:06:16|
|[gotype](https://github.com/wzshiming/gotype) | 17| 3| 2017-12-05 12:09:47| 2019-02-20 15:26:44|
|[gowrap](https://github.com/hexdigest/gowrap) | 203| 6| 2018-09-15 17:20:42| 2019-02-26 17:51:56|
|[interfaces](https://github.com/rjeczalik/interfaces) | 160| 5| 2015-12-06 08:04:50| 2018-12-21 17:45:12|
|[jennifer](https://github.com/dave/jennifer) | 1136| 21| 2016-12-05 04:57:38| 2019-01-30 21:02:10|
|[pkgreflect](https://github.com/ungerik/pkgreflect) | 77| 5| 2012-09-03 15:53:00| 2017-09-05 20:27:27|

## Geographic
        
*Geographic tools and servers*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[geocache](https://github.com/melihmucuk/geocache) | 95| 5| 2016-06-21 18:48:34| 2016-06-22 00:53:18|
|[geoserver](https://github.com/hishamkaram/geoserver) | 19| 2| 2018-03-27 05:36:49| 2018-11-24 11:20:18|
|[gismanager](https://github.com/hishamkaram/gismanager) | 9| 1| 2018-09-29 20:51:37| 2018-10-30 16:55:19|
|[osm](https://github.com/paulmach/osm) | 41| 5| 2016-02-02 08:59:03| 2018-12-15 04:48:15|
|[pbf](https://github.com/maguro/pbf) | 10| 1| 2017-09-19 07:13:18| 2018-10-15 00:12:41|
|[geo](https://github.com/golang/geo) | 800| 71| 2014-12-04 07:02:15| 2019-02-21 21:59:16|
|[tile38](https://github.com/tidwall/tile38) | 5314| 177| 2016-03-05 07:07:44| 2019-03-01 21:56:49|

## Go Compilers
        
*Tools for compiling Go to other languages.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[c4go](https://github.com/Konstantin8105/c4go) | 61| 7| 2018-03-28 14:24:57| 2019-03-01 18:52:53|
|[f4go](https://github.com/Konstantin8105/f4go) | 9| 2| 2018-07-09 01:05:43| 2018-12-21 20:53:18|
|[gopherjs](https://github.com/gopherjs/gopherjs) | 7984| 250| 2013-08-28 06:23:58| 2019-02-05 13:50:24|
|[llgo](https://github.com/go-llvm/llgo) | 951| 61| 2011-11-05 22:23:32| 2015-01-05 09:15:37|
|[tardisgo](https://github.com/tardisgo/tardisgo) | 382| 27| 2014-01-08 19:07:33| 2016-11-20 02:08:43|

## Goroutines
        
*Tools for managing and working with Goroutines.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[ants](https://github.com/panjf2000/ants) | 1357| 54| 2018-05-19 09:13:38| 2019-02-24 09:53:29|
|[artifex](https://github.com/borderstech/artifex) | 6| 2| 2018-11-01 03:34:31| 2018-11-03 06:43:19|
|[async](https://github.com/StudioSol/async) | 13| 9| 2017-07-01 01:08:33| 2018-09-18 21:53:08|
|[breaker](https://github.com/kamilsk/breaker) | 5| 1| 2019-02-15 23:08:24| 2019-03-02 13:45:37|
|[cyclicbarrier](https://github.com/marusama/cyclicbarrier) | 23| 3| 2018-01-11 18:38:46| 2018-10-27 18:23:13|
|[go-floc](https://github.com/workanator/go-floc) | 164| 8| 2017-07-03 15:34:06| 2018-01-26 13:46:48|
|[go-flow](https://github.com/kamildrazkiewicz/go-flow) | 94| 7| 2016-09-25 22:46:09| 2017-09-19 15:20:06|
|[go-trylock](https://github.com/subchen/go-trylock) | 4| 1| 2018-04-26 14:02:47| 2018-05-29 12:27:46|
|[GoSlaves](https://github.com/dgrr/GoSlaves) | 60| 4| 2017-09-18 03:05:14| 2018-08-02 07:50:33|
|[goworker](https://github.com/benmanns/goworker) | 2130| 70| 2013-07-23 01:04:27| 2019-02-06 01:37:25|
|[gpool](https://github.com/sherifabdlnaby/gpool) | 44| 1| 2018-12-03 12:23:35| 2019-01-17 17:05:52|
|[grpool](https://github.com/ivpusic/grpool) | 459| 29| 2015-07-22 08:15:04| 2019-01-28 07:07:22|
|[parallel-fn](https://github.com/rafaeljesus/parallel-fn) | 24| 3| 2017-06-18 17:47:54| 2018-01-02 04:34:49|
|[pool](https://github.com/go-playground/pool) | 451| 11| 2015-10-29 00:36:08| 2016-08-24 03:37:33|
|[semaphore](https://github.com/kamilsk/semaphore) | 59| 1| 2016-10-08 19:48:12| 2019-03-02 03:29:27|
|[semaphore](https://github.com/marusama/semaphore) | 62| 4| 2017-11-22 22:00:58| 2019-01-10 15:54:42|
|[stl](https://github.com/ssgreg/stl) | 5| 1| 2018-06-19 18:50:11| 2018-09-28 22:57:12|
|[threadpool](https://github.com/shettyh/threadpool) | 12| 2| 2017-09-07 02:45:39| 2017-12-12 00:31:30|
|[tunny](https://github.com/Jeffail/tunny) | 1134| 51| 2014-04-03 00:14:58| 2018-11-09 04:57:23|
|[worker-pool](https://github.com/vardius/worker-pool) | 29| 3| 2017-10-04 17:18:31| 2019-02-05 18:10:09|
|[workerpool](https://github.com/gammazero/workerpool) | 79| 5| 2016-05-17 22:32:06| 2018-12-31 04:30:50|

## GUI
        
*Interaction*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[app](https://github.com/maxence-charriere/app) | 2748| 106| 2016-10-12 08:31:33| 2019-03-01 10:09:46|
|[fyne](https://github.com/fyne-io/fyne) | 745| 40| 2018-02-05 06:07:16| 2019-03-01 07:45:40|
|[go-astilectron](https://github.com/asticode/go-astilectron) | 2353| 112| 2017-04-22 15:59:15| 2019-01-11 17:38:57|
|[go-sciter](https://github.com/sciter-sdk/go-sciter) | 1296| 113| 2015-10-15 20:41:06| 2018-11-01 03:25:52|
|[gotk3](https://github.com/gotk3/gotk3) | 617| 43| 2015-08-13 21:09:46| 2019-02-28 02:37:47|
|[gowd](https://github.com/dtylman/gowd) | 172| 19| 2017-03-29 22:50:53| 2018-10-25 17:08:40|
|[qt](https://github.com/therecipe/qt) | 5257| 274| 2014-11-19 08:03:08| 2019-01-22 06:05:19|
|[ui](https://github.com/andlabs/ui) | 6443| 355| 2014-02-18 07:44:00| 2019-02-07 00:27:53|
|[walk](https://github.com/lxn/walk) | 3274| 239| 2010-09-16 16:11:49| 2019-02-28 19:54:49|
|[webview](https://github.com/zserge/webview) | 3888| 173| 2017-08-19 16:26:00| 2019-02-20 08:46:14|
|[gosx-notifier](https://github.com/deckarep/gosx-notifier) | 476| 15| 2013-11-25 14:35:16| 2018-02-01 11:58:18|
|[robotgo](https://github.com/go-vgo/robotgo) | 3910| 176| 2016-09-27 00:26:56| 2019-03-01 19:57:51|
|[systray](https://github.com/getlantern/systray) | 634| 38| 2014-11-12 11:41:57| 2019-02-17 00:38:38|
|[trayhost](https://github.com/shurcooL/trayhost) | 143| 4| 2014-04-25 12:02:30| 2018-10-21 05:16:47|

## Hardware
        
*Libraries, tools, and tutorials for interacting with hardware.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|

## Images
        
*Libraries for programming devices of the IoT.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[bild](https://github.com/anthonynsimon/bild) | 1930| 55| 2016-08-01 23:54:29| 2018-10-05 06:01:15|
|[bimg](https://github.com/h2non/bimg) | 692| 35| 2015-03-17 22:14:02| 2019-02-23 05:48:24|
|[cameron](https://github.com/aofei/cameron) | 17| 1| 2018-05-06 06:13:11| 2019-02-26 16:49:18|
|[geopattern](https://github.com/pravj/geopattern) | 982| 20| 2014-10-23 01:26:30| 2019-01-09 04:17:57|
|[gg](https://github.com/fogleman/gg) | 1708| 68| 2016-02-19 05:05:08| 2019-02-21 06:12:58|
|[gift](https://github.com/disintegration/gift) | 1155| 49| 2014-07-13 02:47:40| 2018-10-06 06:10:33|
|[go-cairo](https://github.com/ungerik/go-cairo) | 83| 6| 2012-08-23 02:27:01| 2018-09-10 22:37:57|
|[go-gd](https://github.com/bolknote/go-gd) | 47| 4| 2011-05-12 14:33:54| 2018-05-08 03:29:26|
|[go-nude](https://github.com/koyachi/go-nude) | 271| 17| 2014-05-02 16:32:29| 2018-11-22 23:22:42|
|[go-opencv](https://github.com/go-opencv/go-opencv) | 1025| 68| 2013-12-09 17:43:26| 2019-01-15 15:07:43|
|[go-webcolors](https://github.com/jyotiska/go-webcolors) | 23| 1| 2014-04-24 22:41:22| 2015-08-21 12:56:56|
|[gocv](https://github.com/hybridgroup/gocv) | 1988| 102| 2017-09-19 05:54:17| 2019-02-20 18:16:57|
|[goimagehash](https://github.com/corona10/goimagehash) | 167| 9| 2017-07-29 01:15:58| 2019-02-08 17:02:26|
|[govatar](https://github.com/o1egl/govatar) | 280| 5| 2016-01-18 20:12:28| 2018-04-13 08:07:57|
|[image2ascii](https://github.com/qeesung/image2ascii) | 238| 7| 2018-10-20 13:06:25| 2018-11-07 23:15:43|
|[imagick](https://github.com/gographics/imagick) | 883| 54| 2013-05-01 01:31:48| 2019-01-15 10:11:26|
|[imaginary](https://github.com/h2non/imaginary) | 2307| 70| 2015-03-05 02:51:40| 2019-02-25 16:54:52|
|[imaging](https://github.com/disintegration/imaging) | 2181| 63| 2012-12-07 04:21:21| 2019-02-03 00:04:28|
|[img](https://github.com/hawx/img) | 126| 5| 2012-07-29 03:57:47| 2015-05-01 23:11:26|
|[ln](https://github.com/fogleman/ln) | 2377| 89| 2016-01-10 12:28:10| 2018-02-24 02:16:49|
|[mergi](https://github.com/noelyahan/mergi) | 42| 2| 2018-09-24 11:40:47| 2019-02-25 13:59:04|
|[mort](https://github.com/aldor007/mort) | 341| 17| 2017-11-19 21:37:58| 2019-02-19 04:07:09|
|[mpo](https://github.com/donatj/mpo) | 5| 1| 2015-04-15 06:37:59| 2019-01-26 12:06:30|
|[picfit](https://github.com/thoas/picfit) | 964| 41| 2014-12-07 01:30:45| 2019-01-29 23:39:30|
|[pt](https://github.com/fogleman/pt) | 1715| 62| 2015-01-24 03:39:29| 2018-03-09 10:44:07|
|[resize](https://github.com/nfnt/resize) | 2020| 73| 2012-08-03 03:48:26| 2018-02-22 03:10:15|
|[rez](https://github.com/bamiaux/rez) | 156| 9| 2014-01-17 05:16:15| 2017-08-01 02:51:31|
|[smartcrop](https://github.com/muesli/smartcrop) | 1203| 30| 2014-04-08 06:40:03| 2018-10-31 06:06:01|
|[steganography](https://github.com/auyer/steganography) | 12| 3| 2018-05-22 01:27:36| 2018-10-18 23:00:19|
|[stegify](https://github.com/DimitarPetrov/stegify) | 280| 13| 2018-11-30 00:45:58| 2019-02-12 23:23:43|
|[svgo](https://github.com/ajstarks/svgo) | 1220| 46| 2010-03-06 07:24:10| 2018-10-06 08:33:21|
|[tga](https://github.com/ftrvxmtrx/tga) | 21| 2| 2012-10-08 09:09:20| 2015-05-24 16:11:41|
|[connectordb](https://github.com/connectordb/connectordb) | 152| 17| 2015-01-17 03:44:21| 2019-03-01 07:33:49|
|[devices](https://github.com/goiot/devices) | 221| 17| 2016-05-30 16:07:02| 2016-07-10 08:46:08|
|[eywa](https://github.com/xcodersun/eywa) | 30| 8| 2016-02-21 01:02:16| 2017-04-12 15:41:51|
|[flogo](https://github.com/TIBCOSoftware/flogo) | 951| 112| 2016-07-10 10:57:43| 2019-03-01 08:52:46|
|[gatt](https://github.com/paypal/gatt) | 767| 57| 2014-04-23 21:45:27| 2018-11-23 19:01:00|
|[gobot](https://github.com/hybridgroup/gobot) | 5196| 298| 2013-09-21 22:09:19| 2019-02-14 02:17:11|
|[huego](https://github.com/amimof/huego) | 83| 1| 2017-05-16 13:31:45| 2019-02-27 02:14:03|
|[iot](https://github.com/vaelen/iot) | 28| 4| 2018-03-08 14:51:51| 2018-04-17 15:28:57|
|[mainflux](https://github.com/mainflux/mainflux) | 359| 53| 2015-07-07 04:31:50| 2019-03-02 13:22:09|
|[sensorbee](https://github.com/sensorbee/sensorbee) | 165| 19| 2016-02-19 15:49:56| 2017-12-31 08:14:29|

## IoT
        

## Job Scheduler
        
*Libraries for scheduling jobs.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[clockwork](https://github.com/whiteShtef/clockwork) | 48| 1| 2018-04-24 01:43:09| 2019-02-13 08:01:41|
|[go-cron](https://github.com/rk/go-cron) | 169| 9| 2011-04-15 22:50:49| 2015-05-15 08:30:45|
|[gron](https://github.com/roylee0704/gron) | 551| 16| 2016-06-04 16:02:22| 2019-02-24 20:15:10|
|[jobrunner](https://github.com/bamzi/jobrunner) | 516| 17| 2015-10-21 12:17:01| 2016-10-19 22:30:22|
|[jobs](https://github.com/albrow/jobs) | 438| 18| 2015-02-10 06:13:29| 2018-06-17 05:00:16|
|[leprechaun](https://github.com/kilgaloon/leprechaun) | 22| 4| 2018-04-08 21:44:04| 2019-02-28 19:03:38|
|[scheduler](https://github.com/carlescere/scheduler) | 266| 16| 2015-02-04 01:10:23| 2018-06-14 04:05:12|

## JSON
        
*Libraries for working with JSON.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[gjson](https://github.com/tidwall/gjson) | 4077| 111| 2016-08-11 11:08:47| 2019-02-17 10:21:13|
|[go-respond](https://github.com/nicklaw5/go-respond) | 18| 1| 2017-03-13 05:00:54| 2018-10-31 02:37:50|
|[gojq](https://github.com/elgs/gojq) | 132| 5| 2015-12-30 17:02:13| 2018-05-31 07:43:05|
|[gojson](https://github.com/ChimeraCoder/gojson) | 1917| 45| 2012-12-28 03:10:50| 2019-02-13 09:14:39|
|[jaydiff](https://github.com/yazgazan/jaydiff) | 33| 1| 2017-04-25 00:05:35| 2018-12-06 21:15:54|
|[jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) | 5| 1| 2016-07-08 18:08:58| 2016-11-18 00:02:12|
|[jsonf](https://github.com/miolini/jsonf) | 53| 3| 2015-05-25 12:53:32| 2016-07-08 08:43:10|
|[jsongo](https://github.com/ricardolonga/jsongo) | 85| 1| 2015-08-08 07:23:17| 2016-12-15 19:09:33|
|[jsonhal](https://github.com/RichardKnop/jsonhal) | 8| 2| 2016-01-15 19:38:40| 2018-11-01 11:57:04|
|[kazaam](https://github.com/qntfy/kazaam) | 106| 15| 2016-07-19 22:19:03| 2018-08-29 09:46:03|
|[mp](https://github.com/sanbornm/mp) | 29| 2| 2014-06-16 05:14:39| 2016-05-12 03:40:58|

## Logging
        
*Libraries for generating and working with log files.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[gone](https://github.com/One-com/gone) | 24| 7| 2016-09-05 17:39:11| 2019-02-08 20:52:36|
|[distillog](https://github.com/amoghe/distillog) | 15| 1| 2015-10-13 00:32:21| 2018-07-27 07:35:13|
|[glg](https://github.com/kpango/glg) | 29| 3| 2017-06-21 21:26:16| 2019-02-27 12:05:56|
|[glo](https://github.com/lajosbencz/glo) | 3| 1| 2019-01-20 06:10:42| 2019-01-23 19:35:10|
|[glog](https://github.com/golang/glog) | 2147| 84| 2013-07-16 12:33:04| 2018-11-10 17:44:40|
|[go-cronowriter](https://github.com/utahta/go-cronowriter) | 17| 1| 2017-02-04 17:02:55| 2018-09-01 09:11:16|
|[go-log](https://github.com/subchen/go-log) | 8| 1| 2017-05-07 16:09:24| 2018-05-19 16:03:37|
|[go-log](https://github.com/siddontang/go-log) | 20| 3| 2014-05-18 11:41:55| 2019-02-21 10:24:31|
|[go-log](https://github.com/ian-kent/go-log) | 33| 2| 2014-05-02 08:34:09| 2018-03-31 10:06:55|
|[go-logger](https://github.com/apsdehal/go-logger) | 212| 8| 2014-09-26 12:57:06| 2018-09-30 07:48:05|
|[gologger](https://github.com/sadlil/gologger) | 38| 6| 2015-09-02 16:52:26| 2018-01-31 11:17:58|
|[gomol](https://github.com/aphistic/gomol) | 12| 2| 2015-08-30 23:51:46| 2018-05-30 00:49:40|
|[journald](https://github.com/ssgreg/journald) | 17| 2| 2017-08-23 15:06:09| 2018-12-25 16:19:33|
|[log](https://github.com/apex/log) | 675| 35| 2015-12-22 04:27:48| 2018-10-12 00:46:42|
|[log](https://github.com/go-playground/log) | 257| 10| 2016-02-08 00:17:48| 2019-02-13 23:35:32|
|[log](https://github.com/teris-io/log) | 21| 1| 2017-10-29 03:57:55| 2017-12-05 02:53:45|
|[logvoyage](https://github.com/firstrow/logvoyage) | 83| 5| 2015-03-29 19:05:09| 2017-05-25 03:48:17|
|[log15](https://github.com/inconshreveable/log15) | 843| 23| 2014-05-20 08:11:52| 2018-10-15 22:41:02|
|[logdump](https://github.com/ewwwwwqm/logdump) | 7| 2| 2017-01-13 23:34:31| 2018-04-02 08:28:16|
|[logex](https://github.com/chzyer/logex) | 32| 7| 2014-10-10 14:38:39| 2017-03-29 14:49:08|
|[logger](https://github.com/azer/logger) | 130| 4| 2014-09-30 14:45:09| 2018-09-23 15:57:13|
|[logmatic](https://github.com/borderstech/logmatic) | 3| 1| 2018-11-07 09:52:45| 2019-01-13 22:47:44|
|[logo](https://github.com/mbndr/logo) | 4| 1| 2017-02-08 02:02:55| 2017-10-20 03:33:23|
|[logrus](https://github.com/sirupsen/logrus) | 9960| 268| 2013-10-17 03:08:55| 2019-02-27 20:42:50|
|[logrusly](https://github.com/sebest/logrusly) | 24| 5| 2014-09-12 07:27:11| 2018-03-16 03:02:19|
|[logutils](https://github.com/hashicorp/logutils) | 239| 87| 2013-10-09 15:31:15| 2018-08-29 00:26:51|
|[logxi](https://github.com/mgutz/logxi) | 328| 10| 2015-03-02 06:13:45| 2017-12-24 04:05:30|
|[lumberjack](https://github.com/natefinch/lumberjack) | 1187| 37| 2014-06-14 19:55:47| 2019-03-01 21:50:28|
|[mlog](https://github.com/jbrodriguez/mlog) | 16| 2| 2014-10-20 23:06:26| 2018-08-06 01:35:46|
|[onelog](https://github.com/francoispqt/onelog) | 303| 9| 2018-05-06 22:32:10| 2019-03-01 15:51:24|
|[ozzo-log](https://github.com/go-ozzo/ozzo-log) | 101| 11| 2015-10-23 06:29:02| 2018-04-06 17:37:37|
|[seelog](https://github.com/cihub/seelog) | 1269| 89| 2011-11-17 17:43:15| 2019-02-28 20:17:08|
|[go-spew](https://github.com/davecgh/go-spew) | 2976| 59| 2013-01-09 13:18:22| 2019-01-31 13:51:44|
|[log](https://github.com/alexcesaro/log) | 44| 4| 2014-04-19 22:31:56| 2015-09-16 06:13:22|
|[tail](https://github.com/hpcloud/tail) | 1358| 101| 2013-02-05 08:28:03| 2019-02-28 04:56:24|
|[xlog](https://github.com/xfxdev/xlog) | 4| 1| 2016-05-06 00:47:45| 2019-01-15 18:17:30|
|[xlog](https://github.com/rs/xlog) | 126| 7| 2015-10-22 17:26:45| 2018-03-20 02:49:19|
|[zap](https://github.com/uber-go/zap) | 6129| 199| 2016-02-19 03:52:56| 2019-02-15 09:25:42|
|[zerolog](https://github.com/rs/zerolog) | 1764| 47| 2017-05-12 13:24:39| 2019-03-02 11:16:32|

## Machine Learning
        
*Libraries for Machine Learning.*
| Go Repo    | Stars      | Watchers   | Created At | Pushed At |
| :--------- | ---------:| ---------:|:---------:|:---------:|
|[bayesian](https://github.com/jbrukh/bayesian) | 596| 30| 2011-11-23 12:17:00| 2019-02-18 12:38:19|
|[CloudForest](https://github.com/ryanbressler/CloudForest) | 627| 43| 2012-10-23 01:38:16| 2018-11-20 14:58:46|
|[eaopt](https://github.com/MaxHalford/eaopt) | 578| 28| 2016-01-31 08:04:52| 2019-02-20 03:56:00|
|[evoli](https://github.com/khezen/evoli) | 5| 3| 2015-06-12 14:58:30| 2019-01-28 20:40:57|
|[fonet](https://github.com/Fontinalis/fonet) | 26| 3| 2017-10-03 23:57:15| 2018-04-10 21:50:54|
|[go-cluster](https://github.com/e-XpertSolutions/go-cluster) | 18| 5| 2017-10-04 20:24:52| 2018-08-06 15:35:27|
|[go-deep](https://github.com/patrikeh/go-deep) | 201| 11| 2017-12-09 23:10:06| 2018-09-14 20:17:29|
|[go-fann](https://github.com/white-pony/go-fann) | 98| 8| 2011-03-11 05:08:27| 2015-02-04 05:53:31|
|[go-galib](https://github.com/thoj/go-galib) | 164| 14| 2009-11-30 18:46:58| 2015-12-29 00:27:45|
|[go-pr](https://github.com/daviddengcn/go-pr) | 56| 6| 2013-06-07 10:36:20| 2013-06-08 18:17:05|
|[gobrain](https://github.com/goml/gobrain) | 336| 23| 2014-04-29 21:32:36| 2018-12-09 21:17:54|
|[godist](https://github.com/e-dard/godist) | 22| 2| 2014-09-05 17:48:51| 2015-05-11 18:38:48|
|[goga](https://github.com/tomcraven/goga) | 72| 8| 2015-10-20 20:50:51| 2017-01-16 23:29:16|
|[golearn](https://github.com/sjwhitworth/golearn) | 6303| 429| 2013-12-26 21:06:14| 2019-02-07 04:42:36|
|[golinear](https://github.com/danieldk/golinear) | 38| 5| 2013-04-05 23:37:01| 2018-08-29 18:30:44|
|[gomind](https://github.com/surenderthakran/gomind) | 5| 2| 2017-10-19 11:48:51| 2018-07-31 20:57:31|
|[goml](https://github.com/cdipaolo/goml) | 960| 73| 2015-06-27 13:52:01| 2016-10-31 04:49:12|
|[goRecommend](https://github.com/timkaye11/goRecommend) | 136| 10| 2014-07-16 13:32:23| 2014-07-29 12:49:57|
|[gorgonia](https://github.com/gorgonia/gorgonia) | 2443| 161| 2016-09-15 07:19:43| 2019-03-02 00:30:55|
|[gorse](https://github.com/zhenghaoz/gorse) | 338| 16| 2018-08-14 19:01:09| 2019-02-28 16:00:27|
|[goscore](https://github.com/asafschers/goscore) | 27| 2| 2017-08-19 19:08:39| 2018-09-05 18:32:11|
|[gosseract](https://github.com/otiai10/gosseract) | 744| 31| 2013-10-11 15:27:53| 2019-02-25 17:01:49|
|[libsvm](https://github.com/datastream/libsvm) | 59| 10| 2012-07-31 15:57:47| 2016-05-09 11:47:11|
|[mlgo](https://github.com/NullHypothesis/mlgo) | 4| 2| 2015-12-07 23:41:25| 2015-12-08 00:08:20|
|[neat](https://github.com/jinyeom/neat) | 51| 11| 2016-11-17 12:23:14| 2018-07-05 04:45:55|

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
        