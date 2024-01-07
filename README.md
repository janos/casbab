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
BenchmarkCamel-10               2106188        569.1 ns/op      640 B/op        9 allocs/op
BenchmarkPascal
BenchmarkPascal-10              2123118        567.5 ns/op      640 B/op        9 allocs/op
BenchmarkSnake
BenchmarkSnake-10               2091612        567.1 ns/op      640 B/op        9 allocs/op
BenchmarkCamelSnake
BenchmarkCamelSnake-10          2032400        586.8 ns/op      640 B/op        9 allocs/op
BenchmarkScreamingSnake
BenchmarkScreamingSnake-10      1753968        637.1 ns/op      640 B/op        9 allocs/op
BenchmarkKebab
BenchmarkKebab-10               2113995        571.8 ns/op      640 B/op        9 allocs/op
BenchmarkCamelKebab
BenchmarkCamelKebab-10          2047392        586.0 ns/op      640 B/op        9 allocs/op
BenchmarkScreamingKebab
BenchmarkScreamingKebab-10      1863102        635.0 ns/op      640 B/op        9 allocs/op
BenchmarkLower
BenchmarkLower-10               2107701        566.4 ns/op      640 B/op        9 allocs/op
BenchmarkTitle
BenchmarkTitle-10               2056044        584.0 ns/op      640 B/op        9 allocs/op
BenchmarkScreaming
BenchmarkScreaming-10           1887225        644.4 ns/op      640 B/op        9 allocs/op
PASS
```
