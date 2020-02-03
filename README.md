# url-parser

[![Build Status](https://img.shields.io/drone/build/xoxys/url-parser?logo=drone)](https://cloud.drone.io/xoxys/url-parser)
[![Codecov](https://img.shields.io/codecov/c/github/xoxys/url-parser)](https://codecov.io/gh/xoxys/url-parser)
[![Go Report Card](https://goreportcard.com/badge/github.com/xoxys/url-parser)](https://goreportcard.com/report/github.com/xoxys/url-parser)
[![License: MIT](https://img.shields.io/github/license/xoxys/url-parser)](LICENSE)

Inspired by [herloct/url-parser](https://github.com/herloct/url-parser), a simple command-line utility for parsing URLs.

## Instalation

Prebuild multiarch binaries are availabe for Linux only:

```Shell
curl -L https://github.com/xoxys/url-parser/releases/download/v0.1.0/url-parser-0.1.0-linux-amd64 > /usr/local/bin/url-parser
chmod +x /usr/local/bin/url-parser
url-parser --help
```

## Usage

```Shell
$ url-parser --help
NAME:
   url-parser - Parse URL and shows the part of it.

USAGE:
   url-parser [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   all, a       print out all parts from url
   scheme, s    print out scheme from url
   user, u      print out username from url
   password, pw print out password from url
   path, p      print out the path from url
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --url value    source url to parse [$URL_PARSER_URL]
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Examples

```Shell
$ url-parser host --url https://somedomain.com
somedomain.com

$ url-parser user --url https://herloct@somedomain.com
herloct

$ url-parser path --url https://somedomain.com/path/to
/path/to

$ url-parser path --path-index=1 --url https://somedomain.com/path/to
to

$ url-parser query --url https://somedomain.com/?some-key=somevalue
some-key=somevalue

$ url-parser query --query-field=some-key --url https://somedomain.com/?some-key=somevalue
somevalue
```

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Maintainers and Contributors

[Robert Kaussow](https://github.com/xoxys)
