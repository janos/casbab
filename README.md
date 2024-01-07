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
BenchmarkCamel-10               1826664        642.2 ns/op      640 B/op        9 allocs/op
BenchmarkPascal
BenchmarkPascal-10              1830391        638.6 ns/op      640 B/op        9 allocs/op
BenchmarkSnake
BenchmarkSnake-10               1868508        653.8 ns/op      640 B/op        9 allocs/op
BenchmarkCamelSnake
BenchmarkCamelSnake-10          1782997        685.2 ns/op      640 B/op        9 allocs/op
BenchmarkScreamingSnake
BenchmarkScreamingSnake-10      1583528        740.1 ns/op      640 B/op        9 allocs/op
BenchmarkKebab
BenchmarkKebab-10               1859630        640.6 ns/op      640 B/op        9 allocs/op
BenchmarkCamelKebab
BenchmarkCamelKebab-10          1825042        658.5 ns/op      640 B/op        9 allocs/op
BenchmarkScreamingKebab
BenchmarkScreamingKebab-10      1685986        713.3 ns/op      640 B/op        9 allocs/op
BenchmarkLower
BenchmarkLower-10               1862577        640.1 ns/op      640 B/op        9 allocs/op
BenchmarkTitle
BenchmarkTitle-10               1818093        657.6 ns/op      640 B/op        9 allocs/op
BenchmarkScreaming
BenchmarkScreaming-10           1647198        725.4 ns/op      640 B/op        9 allocs/op
PASS
```
