# Camel Snake Kebab

CLI tool and a Go library for converting representation style of compound words or phrases.

[![GoDoc](https://godoc.org/resenje.org/casbab?status.svg)](https://godoc.org/resenje.org/casbab)

## Installation

```sh
go get resenje.org/casbab/cmd/casbab
```

```sh
docker pull janos/casbab
```

## Example usage

```sh
casbab screaming-snake "Camel Snake Kebab"
env | cut -d= -f1 | casbab kebab
```

```sh
docker run --rm janos/casbab screaming-snake "Camel Snake Kebab"
env | cut -d= -f1 | docker run --rm -i janos/casbab kebab
```
