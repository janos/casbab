// Copyright (c) 2016, Janoš Guljaš <janos@resenje.org>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found s the LICENSE file.

// Package casbab is a Go library for converting
// representation style of compound words or phrases.
// Different writing styles of compound words are used
// for different purposes in computer code and variables
// to easily distinguish type, properties or meaning.
//
// Functions in this package are separating words from
// input string and constructing an appropriate phrase
// representation.
//
// Examples:
//
//	Kebab("camel_snake_kebab") == "camel-snake-kebab"
//	ScreamingSnake("camel_snake_kebab") == "CAMEL_SNAKE_KEBAB"
//	Camel("camel_snake_kebab") == "camelSnakeKebab"
//	Pascal("camel_snake_kebab") == "CamelSnakeKebab"
//	Snake("camelSNAKEKebab") == "camel_snake_kebab"
//
// Word separation works by detecting delimiters hyphen (-),
// underscore (_), space ( ) and letter case change.
//
// Note: Leading and trailing separators will be preserved
// only within the Snake family or within the Kebab family
// and not across them. This restriction is based on different
// semantics between different writings.
//
// Examples:
//
//	CamelSnake("__camel_snake_kebab__") == "__Camel_Snake_Kebab__"
//	Kebab("__camel_snake_kebab") == "camel-snake-kebab"
//	Screaming("__camel_snake_kebab") == "CAMEL SNAKE KEBAB"
//	CamelKebab("--camel-snake-kebab") == "--Camel-Snake-Kebab"
//	Snake("--camel-snake-kebab") == "camel_snake_kebab"
//	Screaming("--camel-snake-kebab") == "CAMEL SNAKE KEBAB"
package casbab // import "resenje.org/casbab"

import (
	"strings"
	"unicode"
)

// Camel case is the practice of writing compound words
// or phrases such that each word or abbreviation in the
// middle of the phrase begins with a capital letter,
// with no spaces or hyphens.
//
// Example: "camelSnakeKebab".
func Camel(s string) string {
	r := []rune(s)
	return join(camel(words(r), 1))
}

// Pascal case is a variant of Camel case writing where
// the first letter of the first word is always capitalized.
//
// Example: "CamelSnakeKebab".
func Pascal(s string) string {
	r := []rune(s)
	return join(camel(words(r), 0))
}

// Snake case is the practice of writing compound words
// or phrases in which the elements are separated with
// one underscore character (_) and no spaces, with all
// element letters lowercased within the compound.
//
// Example: "camel_snake_kebab".
func Snake(s string) string {
	r := []rune(s)
	head, tail := headTailCount(r, '_')
	return joinWrap(words(r), '_', head, tail)
}

// CamelSnake case is a variant of Camel case with
// each element's first letter uppercased.
//
// Example: "Camel_Snake_Kebab".
func CamelSnake(s string) string {
	r := []rune(s)
	head, tail := headTailCount(r, '_')
	return joinWrap(camel(words(r), 0), '_', head, tail)
}

// ScreamingSnake case is a variant of Camel case with
// all letters uppercased.
//
// Example: "CAMEL_SNAKE_KEBAB".
func ScreamingSnake(s string) string {
	r := []rune(s)
	head, tail := headTailCount(r, '_')
	return joinWrap(scream(words(r)), '_', head, tail)
}

// Kebab case is the practice of writing compound words
// or phrases in which the elements are separated with
// one hyphen character (-) and no spaces, with all
// element letters lowercased within the compound.
//
// Example: "camel-snake-kebab".
func Kebab(s string) string {
	r := []rune(s)
	head, tail := headTailCount(r, '-')
	return joinWrap(words(r), '-', head, tail)
}

// CamelKebab case is a variant of Kebab case with
// each element's first letter uppercased.
//
// Example: "Camel-Snake-Kebab".
func CamelKebab(s string) string {
	r := []rune(s)
	head, tail := headTailCount(r, '-')
	return joinWrap(camel(words(r), 0), '-', head, tail)
}

// ScreamingKebab case is a variant of Kebab case with
// all letters uppercased.
//
// Example: "CAMEL-SNAKE-KEBAB".
func ScreamingKebab(s string) string {
	r := []rune(s)
	head, tail := headTailCount(r, '-')
	return joinWrap(scream(words(r)), '-', head, tail)
}

// Lower is returning detected words, not in a compound
// form, but separated by one space character with all
// letters in lower case.
//
// Example: "camel snake kebab".
func Lower(s string) string {
	r := []rune(s)
	return joinSep(words(r), ' ')
}

// Title is returning detected words, not in a compound
// form, but separated by one space character with first
// character in all letters in upper case and all other
// letters in lower case.
//
// Example: "Camel Snake Kebab".
func Title(s string) string {
	r := []rune(s)
	return joinSep(camel(words(r), 0), ' ')
}

// Screaming is returning detected words, not in a compound
// form, but separated by one space character with all
// letters in upper case.
//
// Example: "CAMEL SNAKE KEBAB".
func Screaming(s string) string {
	r := []rune(s)
	return joinSep(scream(words(r)), ' ')
}

func words(r []rune) (w [][]rune) {
	start := 0
	l := len(r)
	var prevLower, prevUpper bool
Loop:
	for i, c := range r {
		switch c {
		case '-', '_', ' ':
			if start != i {
				w = append(w, toLower(r[start:i]))
			}
			start = i + 1
			prevLower = false
			prevUpper = false
			continue Loop
		}
		cs := r[i]
		if unicode.ToUpper(cs) == cs {
			prevUpper = true
			if prevLower {
				if i != start {
					w = append(w, toLower(r[start:i]))
				}
				start = i
				prevLower = false
			}
		} else {
			prevLower = true
			if prevUpper {
				if i-1 != start {
					w = append(w, toLower(r[start:i-1]))
				}
				start = i - 1
				prevUpper = false
			}
		}
		if i == l-1 {
			w = append(w, toLower(r[start:]))
		}
	}
	return
}

func scream(s [][]rune) [][]rune {
	for i := 0; i < len(s); i++ {
		s[i] = toUpper(s[i])
	}
	return s
}

func camel(s [][]rune, start int) [][]rune {
	for i := start; i < len(s); i++ {
		r := []rune(s[i])
		if len(r) == 0 {
			continue
		}
		s[i][0] = unicode.ToUpper(r[0])
	}
	return s
}

func headTailCount(r []rune, sub rune) (head, tail int) {
	l := len(r)
	for i := 0; i < l; i++ {
		if r[i] != sub {
			head = i
			break
		}
	}
	for i := l - 1; i >= 0; i-- {
		if r[i] != sub {
			tail = l - i - 1
			break
		}
	}
	return
}

func toUpper(r []rune) []rune {
	for i, e := range r {
		r[i] = unicode.ToUpper(e)
	}
	return r
}

func toLower(r []rune) []rune {
	for i, e := range r {
		r[i] = unicode.ToLower(e)
	}
	return r
}

func join(elems [][]rune) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return string(elems[0])
	}

	var b strings.Builder
	b.WriteString(string(elems[0]))
	for _, s := range elems[1:] {
		b.WriteString(string(s))
	}
	return b.String()
}

func joinSep(elems [][]rune, sep rune) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return string(elems[0])
	}

	var b strings.Builder
	b.WriteString(string(elems[0]))
	for _, s := range elems[1:] {
		b.WriteRune(sep)
		b.WriteString(string(s))
	}
	return b.String()
}

func joinWrap(elems [][]rune, sep rune, prefixSize, suffixSize int) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return string(elems[0])
	}

	var b strings.Builder

	for i := 0; i < prefixSize; i++ {
		b.WriteRune(sep)
	}

	b.WriteString(string(elems[0]))
	for _, s := range elems[1:] {
		b.WriteRune(sep)
		b.WriteString(string(s))
	}

	for i := 0; i < suffixSize; i++ {
		b.WriteRune(sep)
	}
	return b.String()
}
