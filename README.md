# hclog-zerolog

[![Go Report Card](https://goreportcard.com/badge/github.com/weastur/hclog-zerolog)](https://goreportcard.com/report/github.com/weastur/hclog-zerolog)
[![codecov](https://codecov.io/gh/weastur/hclog-zerolog/graph/badge.svg?token=V94BPHSLB0)](https://codecov.io/gh/weastur/hclog-zerolog)
[![test](https://github.com/weastur/hclog-zerolog/actions/workflows/test.yaml/badge.svg)](https://github.com/weastur/hclog-zerolog/actions/workflows/test.yaml)
[![lint](https://github.com/weastur/hclog-zerolog/actions/workflows/lint.yaml/badge.svg)](https://github.com/weastur/hclog-zerolog/actions/workflows/lint.yaml)
[![gitlint](https://github.com/weastur/hclog-zerolog/actions/workflows/gitlint.yaml/badge.svg)](https://github.com/weastur/hclog-zerolog/actions/workflows/gitlint.yaml)
[![pre-commit.ci status](https://results.pre-commit.ci/badge/github/weastur/hclog-zerolog/main.svg)](https://results.pre-commit.ci/latest/github/weastur/hclog-zerolog/main)</br>
![GitHub Release](https://img.shields.io/github/v/release/weastur/hclog-zerolog)
![GitHub commits since latest release](https://img.shields.io/github/commits-since/weastur/hclog-zerolog/latest)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/weastur/hclog-zerolog)
![GitHub License](https://img.shields.io/github/license/weastur/hclog-zerolog)

**hclog-zerolog** is a wrapper for [zerolog](https://github.com/rs/zerolog)
to use it as an implementation of
[hclog.Logger](https://pkg.go.dev/github.com/hashicorp/go-hclog#Logger) interface

## Why?

Working on project with [hashicorp raft](https://pkg.go.dev/github.com/hashicorp/raft) inside,
which is producing a lot of logs while running, I faced with the need to implement
[hclog.Logger](https://pkg.go.dev/github.com/hashicorp/go-hclog#Logger)
compatible [zerolog](https://github.com/rs/zerolog)
wrapper, to be able to organize logging in my app in homogeneous way.

Nothing complicated here, but implementing of 20+ methods just to normalize logging in your app could be
a bit tiring, so sharing my wrapper.

## Installation

```bash
go get -u github.com/weastur/hclog-zerolog
```

## Usage

Using [hashicorp raft](https://pkg.go.dev/github.com/hashicorp/raft) as an example of a library depending of
[hclog.Logger](https://pkg.go.dev/github.com/hashicorp/go-hclog#Logger). Of course, it could be anything else.

```go
import (
	"github.com/hashicorp/raft"
	"github.com/rs/zerolog/log"
	hclogzerolog "github.com/weastur/hclog-zerolog"
)

raftLogger := log.With().Str("component", "raft").Logger()
config := raft.DefaultConfig()
config.Logger = hclogzerolog.New(raftLogger)
```

Despite it's extremely simple, you can refer to the [example](./_example/) and
[godoc](https://pkg.go.dev/github.com/weastur/hclog-zerolog) to see a bit more.

## Security

Refer to the [SECURITY.md](SECURITY.md) file for more information.

## License

Mozilla Public License 2.0

Refer to the [LICENSE](LICENSE) file for more information.
