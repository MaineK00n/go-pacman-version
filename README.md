# go-pacman-version
[![Test](https://github.com/MaineK00n/go-pacman-version/actions/workflows/test.yml/badge.svg)](https://github.com/MaineK00n/go-pacman-version/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/MaineK00n/go-pacman-version)](https://goreportcard.com/report/github.com/MaineK00n/go-pacman-version)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/MaineK00n/go-pacman-version/blob/main/LICENSE)

A Go library for parsing pacman package versions

go-pacman-version is a library for parsing and comparing versions

The implementation is based on [vercmp(8)](https://archlinux.org/pacman/vercmp.8.html), [this implementation](https://gitlab.archlinux.org/pacman/pacman/-/blob/master/lib/libalpm/version.c#L219-260)

OS: Arch Linux

# Installation and Usage

Installation can be done with a normal go get:

```
$ go get github.com/MaineK00n/go-pacman-version
```

## Version Parsing and Comparison

```
import version "github.com/MaineK00n/go-pacman-version"

v1, err := version.NewVersion("5.1.004-1")
v2, err := version.NewVersion("5.1.008-1")

// Comparison example. You can use GreaterThan and Equal as well.
if v1.LessThan(v2) {
    fmt.Printf("%s is less than %s", v1, v2)
}
```

## Version Sorting

```
raw := []string{"5.1.008-1", "5.1.0-2", "5.1.004-1", "5.1.0-1"}
vs := make([]version.Version, len(raw))
for i, r := range raw {
	v, _ := version.NewVersion(r)
	vs[i] = v
}

sort.Slice(vs, func(i, j int) bool {
	return vs[i].LessThan(vs[j])
})
```

# Contribute

1. fork a repository: github.com/MaineK00n/go-pacman-version to github.com/you/repo
2. get original code: `go get github.com/MaineK00n/go-pacman-version`
3. work on original code
4. add remote to your repo: git remote add myfork https://github.com/you/repo.git
5. push your changes: git push myfork
6. create a new Pull Request

- see [GitHub and Go: forking, pull requests, and go-getting](http://blog.campoy.cat/2014/03/github-and-go-forking-pull-requests-and.html)

----

# License
MIT

# Author
MaineK00n([@MaineK00n](https://twitter.com/MaineK00n))