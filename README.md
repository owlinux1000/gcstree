# gcstree

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`gcstree` is a CLI tool to list objects in Google Cloud Storage.

## How to install

```
go install github.com/owlinux1000/gcstree@latest
```

## How to use

```
$ gcstree
Error: requires at least 1 arg(s), only received 0
Usage:
  gcstree <bucket> [flags]

Flags:
  -h, --help   help for gcstree
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