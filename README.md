checkers
=======

[![Build Status](https://travis-ci.org/mackerelio/checkers.svg?branch=master)][travis]
 [![CI](https://github.com/mackerelio/checkers/workflows/ci/badge.svg?branch=master)][actions]
[![Coverage Status](https://coveralls.io/repos/mackerelio/checkers/badge.svg?branch=master)][coveralls]
[![Apache License Version 2.0](http://img.shields.io/badge/license-APACHE2-blue.svg)][license]
[![GoDoc](https://godoc.org/github.com/mackerelio/checkers?status.svg)][godoc]

[travis]: https://travis-ci.org/mackerelio/checkers
[actions]: https://github.com/mackerelio/checkers/actions?workflow=ci
[coveralls]: https://coveralls.io/r/mackerelio/checkers?branch=master
[license]: https://github.com/mackerelio/checkers/blob/master/LICENSE
[godoc]: https://godoc.org/github.com/mackerelio/checkers

## Description

checkers is utility for writing check plugins.

## Synopsis

    chr := checkers.Ok("command ok!")
    chr.Exit()

# CONTRIBUTION

1. Fork ([https://github.com/mackerelio/checkers/fork](https://github.com/mackerelio/checkers/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create new Pull Request

# License

Copyright 2015 Hatena Co., Ltd.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
