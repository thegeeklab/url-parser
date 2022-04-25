local PipelineTest = {
  kind: 'pipeline',
  name: 'test',
  platform: {
    os: 'linux',
    arch: 'amd64',
  },
  steps: [
    {
      name: 'deps',
      image: 'golang:1.18',
      commands: [
        'make deps',
      ],
      volumes: [
        {
          name: 'godeps',
          path: '/go',
        },
      ],
    },
    {
      name: 'lint',
      image: 'golang:1.18',
      commands: [
        'make lint',
      ],
      volumes: [
        {
          name: 'godeps',
          path: '/go',
        },
      ],
    },
    {
      name: 'test',
      image: 'golang:1.18',
      commands: [
        'make test',
      ],
      volumes: [
        {
          name: 'godeps',
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
      name: 'godeps',
      temp: {},
    },
  ],
  trigger: {
    ref: ['refs/heads/main', 'refs/tags/**', 'refs/pull/**'],
  },
};


local PipelineBuildBinaries = {
  kind: 'pipeline',
  name: 'build-binaries',
  platform: {
    os: 'linux',
    arch: 'amd64',
  },
  steps: [
    {
      name: 'build',
      image: 'techknowlogick/xgo:go-1.18.x',
      commands: [
        'make release',
      ],
    },
    {
      name: 'executable',
      image: 'alpine',
      commands: [
        '$(find dist/ -executable -type f -iname ${DRONE_REPO_NAME}-linux-amd64) --help',
      ],
    },
    {
      name: 'changelog-generate',
      image: 'thegeeklab/git-chglog',
      commands: [
        'git fetch -tq',
        'git-chglog --no-color --no-emoji -o CHANGELOG.md ${DRONE_TAG:---next-tag unreleased unreleased}',
      ],
    },
    {
      name: 'changelog-format',
      image: 'thegeeklab/alpine-tools',
      commands: [
        'prettier CHANGELOG.md',
        'prettier -w CHANGELOG.md',
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
        files: ['dist/*'],
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
  name: 'notifications',
  platform: {
    os: 'linux',
    arch: 'amd64',
  },
  steps: [
    {
      name: 'matrix',
      image: 'thegeeklab/drone-matrix',
      settings: {
        homeserver: { from_secret: 'matrix_homeserver' },
        roomid: { from_secret: 'matrix_roomid' },
        template: 'Status: **{{ build.Status }}**<br/> Build: [{{ repo.Owner }}/{{ repo.Name }}]({{ build.Link }}){{#if build.Branch}} ({{ build.Branch }}){{/if}} by {{ commit.Author }}<br/> Message: {{ commit.Message.Title }}',
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
