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

## Performance

Benchmarks run on MacBook Pro M1Pro yield these timings:

```
goos: darwin
goarch: arm64
pkg: resenje.org/casbab
BenchmarkCamel
BenchmarkCamel-10               3493452        336.7 ns/op      208 B/op        2 allocs/op
BenchmarkPascal
BenchmarkPascal-10              3532380        340.4 ns/op      208 B/op        2 allocs/op
BenchmarkSnake
BenchmarkSnake-10               3257454        366.4 ns/op      288 B/op        3 allocs/op
BenchmarkCamelSnake
BenchmarkCamelSnake-10          3164365        378.6 ns/op      288 B/op        3 allocs/op
BenchmarkScreamingSnake
BenchmarkScreamingSnake-10      2887975        415.8 ns/op      288 B/op        3 allocs/op
BenchmarkKebab
BenchmarkKebab-10               3232392        371.3 ns/op      288 B/op        3 allocs/op
BenchmarkCamelKebab
BenchmarkCamelKebab-10          3160936        382.2 ns/op      288 B/op        3 allocs/op
BenchmarkScreamingKebab
BenchmarkScreamingKebab-10      2873886        420.2 ns/op      288 B/op        3 allocs/op
BenchmarkLower
BenchmarkLower-10               3300728        369.9 ns/op      288 B/op        3 allocs/op
BenchmarkTitle
BenchmarkTitle-10               3079971        380.6 ns/op      288 B/op        3 allocs/op
BenchmarkScreaming
BenchmarkScreaming-10           2898652        417.5 ns/op      288 B/op        3 allocs/op
PASS
```
