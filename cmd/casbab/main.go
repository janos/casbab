// Copyright (c) 2016, Janoš Guljaš <janos@resenje.org>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found s the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	"resenje.org/casbab"
)

var (
	dialects = map[string]func(string) string{
		"camel":           casbab.Camel,
		"pascal":          casbab.Pascal,
		"snake":           casbab.Snake,
		"camel-snake":     casbab.CamelSnake,
		"screaming-snake": casbab.ScreamingSnake,
		"kebab":           casbab.Kebab,
		"camel-kebab":     casbab.CamelKebab,
		"screaming-kebab": casbab.ScreamingKebab,
		"lower":           casbab.Lower,
		"title":           casbab.Title,
		"screaming":       casbab.Screaming,
	}
)

func main() {
	switch len(os.Args) {
	case 1:
		fmt.Printf(`casbab: Camel Snake Kebab (https://github.com/janos/casbab)

USAGE: %s DIALECT [phrases..]

DIALECT can be one of the following:
  - camel            "camelSnakeKebab"
  - pascal           "CamelSnakeKebab"
  - snake            "camel_snake_kebab"
  - camel-snake      "Camel_Snake_Kebab"
  - screaming-snake  "CAMEL_SNAKE_KEBAB"
  - kebab            "camel-snake-kebab"
  - camel-kebab      "Camel-Snake-Kebab"
  - screaming-kebab  "CAMEL-SNAKE-KEBAB"
  - lower            "camel snake kebab"
  - title            "Camel Snake Kebab"
  - screaming        "CAMEL SNAKE KEBAB"

If no phrases are provided as arguments, arguments will be read from the
Stdin as the new-line separated list.

`, os.Args[0])
		os.Exit(1)
	case 2:
		dialect, ok := dialects[os.Args[1]]
		if !ok {
			fmt.Printf("unknown dialect %s\n", os.Args[1])
			os.Exit(1)
		}

		s, err := os.Stdin.Stat()
		if err != nil {
			fmt.Println("unable to acquire stdin")
			os.Exit(1)
		}

		if s.Mode()&os.ModeNamedPipe == 0 {
			return
		}

		r := bufio.NewReader(os.Stdin)
		var (
			phrase, line []byte
			isPrefix     bool
		)
		for {
			line, isPrefix, err = r.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Printf("reading stdin: %s\n", err)
				os.Exit(1)
			}
			if isPrefix {
				phrase = append(phrase, line...)
			} else {
				phrase = line
			}
			if len(bytes.TrimSpace(phrase)) == 0 {
				continue
			}
			fmt.Println(dialect(string(phrase)))
		}
	default:
		dialect, ok := dialects[os.Args[1]]
		if !ok {
			fmt.Printf("unknown dialect %s\n", os.Args[1])
			os.Exit(1)
		}
		for _, phrase := range os.Args[2:] {
			fmt.Println(dialect(string(phrase)))
		}
	}
}
