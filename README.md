# Camel Snake Kebab

CLI tool and a Go library for converting representation style of compound words or phrases.

[![GoDoc](https://godoc.org/resenje.org/casbab?status.svg)](https://godoc.org/resenje.org/casbab)
[![Build Status](https://travis-ci.org/janos/casbab.svg?branch=master)](https://travis-ci.org/janos/casbab)

## Installation

```sh
go get resenje.org/casbab/cli/casbab
```

## Example usage

```sh
casbab screaming-snake "Camel Snake Kebab"
env | cut -d= -f1 | casbab kebab
```