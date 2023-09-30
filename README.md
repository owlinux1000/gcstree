# gcstree

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![](https://github.com/owlinux1000/gcstree/actions/workflows/release.yaml/badge.svg)](https://github.com/owlinux1000/gcstree/actions)

`gcstree` is a CLI tool to list objects in Google Cloud Storage.

## Installation

You can install `gcstree` by `go install` or `brew`. In addition, you can download the binary from [releases](https://github.com/owlinux1000/gcstree/releases)

```
go install github.com/owlinux1000/gcstree@latest
```

```
brew install owlinux1000/tap/gcstree
```


## How to use

In advance, you might want to login to Google Cloud as follows:

```
gcloud auth application-default login
```

```
$ gcstree
A tree command for Google Cloud Storage

Usage:
  gcstree <bucket> [flags]

Flags:
  -h, --help      help for gcstree
  -v, --version   show the gcstree version
```

```
$ gcstree test
test
├── folder1
│   ├── folder1-1
│   │   └── hello.txt
│   └── folder1-2
└── folder2
    └── hello.txt

$ gcstree test/folder1
test
└── folder1
    ├── folder1-1
    │   └── hello.txt
    └── folder1-2
```