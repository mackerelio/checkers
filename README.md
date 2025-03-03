checkers
=======

 [![Build Status](https://github.com/mackerelio/checkers/workflows/Build/badge.svg?branch=master)][actions]
[![Apache License Version 2.0](https://img.shields.io/badge/license-APACHE2-blue.svg)][license]
[![GoDev](https://pkg.go.dev/badge/github.com/mackerelio/checkers)][godev]

[actions]: https://github.com/mackerelio/checkers/actions?workflow=Build
[license]: https://github.com/mackerelio/checkers/blob/master/LICENSE
[godev]: https://pkg.go.dev/github.com/mackerelio/checkers

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
