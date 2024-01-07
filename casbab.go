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
	return casbabTitle(s, false)
}

// Pascal case is a variant of Camel case writing where
// the first letter of the first word is always capitalized.
//
// Example: "CamelSnakeKebab".
func Pascal(s string) string {
	return casbabTitle(s, true)
}

// Snake case is the practice of writing compound words
// or phrases in which the elements are separated with
// one underscore character (_) and no spaces, with all
// element letters lowercased within the compound.
//
// Example: "camel_snake_kebab".
func Snake(s string) string {
	return casbabWrap(s, '_', self)
}

// CamelSnake case is a variant of Camel case with
// each element's first letter uppercased.
//
// Example: "Camel_Snake_Kebab".
func CamelSnake(s string) string {
	return casbabWrap(s, '_', title)
}

// ScreamingSnake case is a variant of Camel case with
// all letters uppercased.
//
// Example: "CAMEL_SNAKE_KEBAB".
func ScreamingSnake(s string) string {
	return casbabWrap(s, '_', upper)
}

// Kebab case is the practice of writing compound words
// or phrases in which the elements are separated with
// one hyphen character (-) and no spaces, with all
// element letters lowercased within the compound.
//
// Example: "camel-snake-kebab".
func Kebab(s string) string {
	return casbabWrap(s, '-', self)
}

// CamelKebab case is a variant of Kebab case with
// each element's first letter uppercased.
//
// Example: "Camel-Snake-Kebab".
func CamelKebab(s string) string {
	return casbabWrap(s, '-', title)
}

// ScreamingKebab case is a variant of Kebab case with
// all letters uppercased.
//
// Example: "CAMEL-SNAKE-KEBAB".
func ScreamingKebab(s string) string {
	return casbabWrap(s, '-', upper)
}

// Lower is returning detected words, not in a compound
// form, but separated by one space character with all
// letters in lower case.
//
// Example: "camel snake kebab".
func Lower(s string) string {
	return casbabSeparate(s, ' ')
}

// Title is returning detected words, not in a compound
// form, but separated by one space character with first
// character in all letters in upper case and all other
// letters in lower case.
//
// Example: "Camel Snake Kebab".
func Title(s string) string {
	return casbabTransform(s, ' ', title)
}

// Screaming is returning detected words, not in a compound
// form, but separated by one space character with all
// letters in upper case.
//
// Example: "CAMEL SNAKE KEBAB".
func Screaming(s string) string {
	return casbabTransform(s, ' ', upper)
}

func casbabTitle(s string, capitalizeFirstWord bool) string {
	runes := []rune(s)

	var b strings.Builder
	b.Grow(len(s))

	w, runes := firstWord(runes)
	if capitalizeFirstWord {
		w = title(w)
	}
	for _, e := range w {
		b.WriteRune(e)
	}

	for {
		w, rest := firstWord(runes)
		if w == nil {
			break
		}
		for _, e := range title(w) {
			b.WriteRune(e)
		}
		if rest == nil {
			break
		}
		runes = rest
	}
	return b.String()
}

func casbabSeparate(s string, separator rune) string {
	runes := []rune(s)

	var b strings.Builder
	b.Grow(len(s))

	w, runes := firstWord(runes)
	for _, e := range w {
		b.WriteRune(e)
	}

	for {
		w, rest := firstWord(runes)
		if w == nil {
			break
		}
		b.WriteRune(separator)
		for _, e := range w {
			b.WriteRune(e)
		}
		if rest == nil {
			break
		}
		runes = rest
	}
	return b.String()
}

func casbabTransform(s string, separator rune, transformWord func([]rune) []rune) string {
	runes := []rune(s)

	var b strings.Builder
	b.Grow(len(s))

	w, runes := firstWord(runes)
	for _, e := range transformWord(w) {
		b.WriteRune(e)
	}

	for {
		w, rest := firstWord(runes)
		if w == nil {
			break
		}
		b.WriteRune(separator)
		for _, e := range transformWord(w) {
			b.WriteRune(e)
		}
		if rest == nil {
			break
		}
		runes = rest
	}
	return b.String()
}

func casbabWrap(s string, separator rune, transformWord func([]rune) []rune) string {
	runes := []rune(s)

	head, tail := headTailCount(runes, separator)

	var b strings.Builder
	b.Grow(len(s) + head + tail)

	for i := 0; i < head; i++ {
		b.WriteRune(separator)
	}

	w, runes := firstWord(runes)
	for _, e := range transformWord(w) {
		b.WriteRune(e)
	}

	for {
		w, rest := firstWord(runes)
		if w == nil {
			break
		}
		b.WriteRune(separator)
		for _, e := range transformWord(w) {
			b.WriteRune(e)
		}
		if rest == nil {
			break
		}
		runes = rest
	}

	for i := 0; i < tail; i++ {
		b.WriteRune(separator)
	}

	return b.String()
}

func firstWord(r []rune) (word, rest []rune) {
	start := 0
	l := len(r)
	var prevLower, prevUpper bool
Loop:
	for i, c := range r {
		if c == '-' || c == '_' || c == ' ' {
			if start != i {
				return lower(r[start:i]), r[i+1:]
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
					return lower(r[start:i]), r[i:]
				}
				start = i
				prevLower = false
			}
		} else {
			prevLower = true
			if prevUpper {
				if i-1 != start {
					return lower(r[start : i-1]), r[i-1:]
				}
				start = i - 1
				prevUpper = false
			}
		}
		if i == l-1 {
			return lower(r[start:]), nil
		}
	}
	return nil, nil
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

func upper(r []rune) []rune {
	for i, e := range r {
		r[i] = unicode.ToUpper(e)
	}
	return r
}

func lower(r []rune) []rune {
	for i, e := range r {
		r[i] = unicode.ToLower(e)
	}
	return r
}

func title(r []rune) []rune {
	if len(r) > 0 {
		r[0] = unicode.ToUpper(r[0])
	}
	return r
}

func self(r []rune) []rune {
	return r
}
