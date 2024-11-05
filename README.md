# url-parser

Simple command-line URL parser

[![Build Status](https://ci.thegeeklab.de/api/badges/thegeeklab/url-parser/status.svg)](https://ci.thegeeklab.de/repos/thegeeklab/url-parser)
[![Go Report Card](https://goreportcard.com/badge/github.com/thegeeklab/url-parser)](https://goreportcard.com/report/github.com/thegeeklab/url-parser)
[![GitHub contributors](https://img.shields.io/github/contributors/thegeeklab/url-parser)](https://github.com/thegeeklab/url-parser/graphs/contributors)
[![License: MIT](https://img.shields.io/github/license/thegeeklab/url-parser)](https://github.com/thegeeklab/url-parser/blob/main/LICENSE)

Inspired by [herloct/url-parser](https://github.com/herloct/url-parser), a simple command-line utility for parsing URLs.

## Installation

Prebuilt multiarch binaries are available for Linux only.

```Shell
curl -SsfL https://github.com/thegeeklab/url-parser/releases/latest/download/url-parser-linux-amd64 -o /usr/local/bin/url-parser
chmod +x /usr/local/bin/url-parser
```

## Build

Build the binary from source with the following command:

```Shell
make build
```

## Usage

```Shell
$ url-parser --help
NAME:
   url-parser - Parse URL and shows the part of it.

USAGE:
   url-parser [global options] command [command options]

VERSION:
   devel

COMMANDS:
   all, a        Get all parts from url
   scheme, s     Get scheme from url
   user, u       Get username from url
   password, pw  Get password from url
   path, pt      Get path from url
   host, ht      Get hostname from url
   port, p       Get port from url
   query, q      Get query from url
   fragment, f   Get fragment from url
   help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --url value    source url to parse [$URL_PARSER_URL]
   --help, -h     show help
   --version, -v  print the version
```

## Examples

```Shell
$ url-parser --url https://somedomain.com host
somedomain.com

$ url-parser --url https://herloct@somedomain.com user
herloct

$ url-parser --url https://somedomain.com/path/to path
/path/to

$ url-parser --url https://somedomain.com/path/to path --path-index=1
to

$ url-parser --url https://somedomain.com/?some-key=somevalue query
some-key=somevalue

$ url-parser --url https://somedomain.com/?some-key=somevalue query --query-field=some-key
somevalue

# It is also possible to read the URL from stdin
$ echo "https://somedomain.com" | url-parser host
somedomain.com

# Get json output or all parsed parts
$ url-parser --url https://somedomain.com/?some-key=somevalue all --json
{"scheme":"https","hostname":"somedomain.com","port":"","path":"/","fragment":"","rawQuery":"some-key=somevalue","queryParams":[{"key":"some-key","value":"somevalue"}],"username":"","password":""}
```

## Contributors

Special thanks to all [contributors](https://github.com/thegeeklab/url-parser/graphs/contributors). If you would like to contribute, please see the [instructions](https://github.com/thegeeklab/url-parser/blob/main/CONTRIBUTING.md).

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/thegeeklab/url-parser/blob/main/LICENSE) file for details.
