def main(ctx):
  before = testing()

  stages = [
    binaries([]),
  ]

  after = notification()

  for b in before:
    for s in stages:
      s['depends_on'].append(b['name'])

  for s in stages:
    for a in after:
      a['depends_on'].append(s['name'])

  return before + stages + after

def testing():
  return [{
    'kind': 'pipeline',
    'type': 'docker',
    'name': 'testing',
    'platform': {
      'os': 'linux',
      'arch': 'amd64',
    },
    'steps': [
      {
        'name': 'vet',
        'image': 'golang:1.12',
        'commands': [
          'go vet ./...'
        ],
      },
      {
        'name': 'test',
        'image': 'golang:1.12',
        'commands': [
          'go test -race -coverprofile=coverage.txt -covermode=atomic ./...'
        ],
      }
    ],
    'trigger': {
      'ref': [
        'refs/heads/master',
        'refs/tags/**',
        'refs/pull/**'
      ]
    }
  }]

def binaries(arch):
  return {
    'kind': 'pipeline',
    'type': 'docker',
    'name': 'build-binaries',
    'steps': [
      {
        'name': 'build',
        'image': 'techknowlogick/xgo:latest',
        'commands': [
          '[ -z "${DRONE_TAG}" ] && BUILD_VERSION=${DRONE_COMMIT_SHA:0:8} || BUILD_VERSION=${DRONE_TAG##v}',
          'mkdir -p release/',
          "cd cmd/url-parser && xgo -ldflags \"-s -w -X main.Version=$BUILD_VERSION\" -tags netgo -targets 'linux/amd64,linux/arm-6,linux/arm64' -out url-parser-$BUILD_VERSION .",
          'cp /build/* /drone/src/release/'
        ]
      },
      {
        'name': 'executable',
        'image': 'alpine',
        'commands': [
          '$(find release/ -executable -type f | grep url-parser-.*-linux-amd64) --help',
        ]
      },
      {
        'name': 'compress',
        'image': 'alpine',
        'commands': [
          'apk add upx',
          'find release/ -maxdepth 1 -executable -type f -exec upx {} \;',
          'ls -lh release/',
        ]
      },
      {
        'name': 'checksum',
        'image': 'alpine',
        'commands': [
            'cd release/ && sha256sum * > sha256sum.txt',
        ],
      },
      {
        'name': 'publish',
        'image': 'plugins/github-release',
        'settings': {
          'overwrite': True,
          'api_key': {
            'from_secret': 'github_token'
          },
          'files': [ "release/*" ],
          'title': '${DRONE_TAG}',
          'note': 'CHANGELOG.md',
        },
        'when': {
          'ref': [
            'refs/tags/**'
          ]
        }
      }
    ],
    'depends_on': [],
    'trigger': {
      'ref': [
        'refs/heads/master',
        'refs/tags/**',
        'refs/pull/**'
      ]
    }
  }

def notification():
  return [{
    'kind': 'pipeline',
    'type': 'docker',
    'name': 'notification',
    'steps': [
      {
        'name': 'coverage',
        'image': 'plugins/codevoc',
        'settings': {
          'token': {
            'from_secret': 'codecov_token',
          },
          'files':[
            'coverage.txt'
          ]
        },
        'when': {
          'status': [
            'success',
          ]
        }
      },
      {
        'name': 'matrix',
        'image': 'plugins/matrix',
        'settings': {
          'homeserver': {
            'from_secret': 'matrix_homeserver',
          },
          'password': {
            'from_secret': 'matrix_password',
          },
          'roomid': {
            'from_secret': 'matrix_roomid',
          },
          'template': 'Status: **{{ build.status }}**<br/> Build: [{{ repo.Owner }}/{{ repo.Name }}]({{ build.link }}) ({{ build.branch }}) by {{ build.author }}<br/> Message: {{ build.message }}',
          'username': {
            'from_secret': 'matrix_username',
          },
        },
        'when': {
          'ref': [
            'refs/heads/master',
            'refs/tags/**',
          ],
        }
      },
    ],
    'depends_on': [],
    'trigger': {
      'ref': [
        'refs/heads/master',
        'refs/tags/**',
        'refs/pull/**',
      ],
      'status': [
        'success',
        'failure'
      ]
    }
  }]
