# url-parser

Simple command-line URL parser

[![Build Status](https://img.shields.io/drone/build/thegeeklab/url-parser?logo=drone)](https://cloud.drone.io/thegeeklab/url-parser)
[![Go Report Card](https://goreportcard.com/badge/github.com/thegeeklab/url-parser)](https://goreportcard.com/report/github.com/thegeeklab/url-parser)
[![Codecov](https://img.shields.io/codecov/c/github/thegeeklab/url-parser)](https://codecov.io/gh/thegeeklab/url-parser)
[![GitHub contributors](https://img.shields.io/github/contributors/thegeeklab/url-parser)](https://github.com/thegeeklab/url-parser/graphs/contributors)
[![License: MIT](https://img.shields.io/github/license/thegeeklab/url-parser)](<[LICENSE](https://github.com/thegeeklab/url-parser/blob/master/LICENSE)>)

Inspired by [herloct/url-parser](https://github.com/herloct/url-parser), a simple command-line utility for parsing URLs.

## Instalation

Prebuild multiarch binaries are availabe for Linux only:

```Shell
curl -L https://github.com/thegeeklab/url-parser/releases/download/v0.1.0/url-parser-0.1.0-linux-amd64 > /usr/local/bin/url-parser
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
   devel

COMMANDS:
   all, a        Get all parts from url
   scheme, s     Get scheme from url
   user, u       Get username from url
   password, pw  Get password from url
   path, pt      Get path from url
   host, h       Get hostname from url
   port, p       Get port from url
   query, q      Get query from url
   fragment, f   Get fragment from url
   help, h       Shows a list of commands or help for one command

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

## Contributors

Special thanks goes to all [contributors](https://github.com/thegeeklab/url-parser/graphs/contributors).

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/thegeeklab/url-parser/blob/master/LICENSE) file for details.
