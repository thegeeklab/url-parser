local PipelineTest = {
  kind: 'pipeline',
  image_pull_secrets: ['docker_config'],
  name: 'test',
  platform: {
    os: 'linux',
    arch: 'amd64',
  },
  steps: [
    {
      name: 'staticcheck',
      image: 'golang:1.15',
      commands: [
        'go run honnef.co/go/tools/cmd/staticcheck ./...',
      ],
      volumes: [
        {
          name: 'gopath',
          path: '/go',
        },
      ],
    },
    {
      name: 'lint',
      image: 'golang:1.15',
      commands: [
        'go run golang.org/x/lint/golint -set_exit_status ./...',
      ],
      volumes: [
        {
          name: 'gopath',
          path: '/go',
        },
      ],
    },
    {
      name: 'vet',
      image: 'golang:1.15',
      commands: [
        'go vet ./...',
      ],
      volumes: [
        {
          name: 'gopath',
          path: '/go',
        },
      ],
    },
    {
      name: 'test',
      image: 'golang:1.15',
      commands: [
        'go test -race -coverprofile=coverage.txt -covermode=atomic ./...',
      ],
      volumes: [
        {
          name: 'gopath',
          path: '/go',
        },
      ],
    },
    {
      name: 'coverage',
      image: 'plugins/codecov',
      settings: {
        token: {
          from_secret: 'codecov_token',
        },
        files: [
          'coverage.txt',
        ],
      },
    },
  ],
  volumes: [
    {
      name: 'gopath',
      temp: {},
    },
  ],
  trigger: {
    ref: ['refs/heads/main', 'refs/tags/**', 'refs/pull/**'],
  },
};


local PipelineBuildBinaries = {
  kind: 'pipeline',
  image_pull_secrets: ['docker_config'],
  name: 'build-binaries',
  platform: {
    os: 'linux',
    arch: 'amd64',
  },
  steps: [
    {
      name: 'build',
      image: 'techknowlogick/xgo:go-1.15.x',
      commands: [
        '[ -z "${DRONE_TAG}" ] && BUILD_VERSION=${DRONE_COMMIT_SHA:0:8} || BUILD_VERSION=${DRONE_TAG##v}',
        'mkdir -p release/',
        "cd cmd/url-parser && xgo -ldflags \"-s -w -X main.version=$BUILD_VERSION\" -tags netgo -targets 'linux/amd64,linux/arm-6,linux/arm-7,linux/arm64' -out url-parser .",
        'mv /build/* /drone/src/release/',
        'ls -l /drone/src/release/',
      ],
    },
    {
      name: 'executable',
      image: 'alpine',
      commands: [
        '$(find release/ -executable -type f | grep url-parser-linux-amd64) --help',
      ],
    },
    {
      name: 'compress',
      image: 'alpine',
      commands: [
        'apk add upx',
        'find release/ -maxdepth 1 -executable -type f -exec upx {} \\;',
        'ls -lh release/',
      ],
    },
    {
      name: 'checksum',
      image: 'alpine',
      commands: [
        'cd release/ && sha256sum * > sha256sum.txt',
      ],
    },
    {
      name: 'changelog',
      image: 'thegeeklab/git-chglog',
      commands: [
        'git fetch -tq',
        'git-chglog --no-color --no-emoji ${DRONE_TAG:---next-tag unreleased unreleased}',
        'git-chglog --no-color --no-emoji -o CHANGELOG.md ${DRONE_TAG:---next-tag unreleased unreleased}',
      ],
    },
    {
      name: 'publish',
      image: 'plugins/github-release',
      settings: {
        overwrite: true,
        api_key: {
          from_secret: 'github_token',
        },
        files: ['release/*'],
        title: '${DRONE_TAG}',
        note: 'CHANGELOG.md',
      },
      when: {
        ref: [
          'refs/tags/**',
        ],
      },
    },
  ],
  depends_on: [
    'test',
  ],
  trigger: {
    ref: ['refs/heads/main', 'refs/tags/**', 'refs/pull/**'],
  },
};

local PipelineNotifications = {
  kind: 'pipeline',
  image_pull_secrets: ['docker_config'],
  name: 'notifications',
  platform: {
    os: 'linux',
    arch: 'amd64',
  },
  steps: [
    {
      name: 'matrix',
      image: 'plugins/matrix',
      settings: {
        homeserver: { from_secret: 'matrix_homeserver' },
        roomid: { from_secret: 'matrix_roomid' },
        template: 'Status: **{{ build.status }}**<br/> Build: [{{ repo.Owner }}/{{ repo.Name }}]({{ build.link }}) ({{ build.branch }}) by {{ build.author }}<br/> Message: {{ build.message }}',
        username: { from_secret: 'matrix_username' },
        password: { from_secret: 'matrix_password' },
      },
      when: {
        status: ['success', 'failure'],
      },
    },
  ],
  depends_on: [
    'build-binaries',
  ],
  trigger: {
    ref: ['refs/heads/main', 'refs/tags/**'],
    status: ['success', 'failure'],
  },
};

[
  PipelineTest,
  PipelineBuildBinaries,
  PipelineNotifications,
]
