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
//     Kebab("camel_snake_kebab") == "camel-snake-kebab"
//     ScreamingSnake("camel_snake_kebab") == "CAMEL_SNAKE_KEBAB"
//     Camel("camel_snake_kebab") == "camelSnakeKebab"
//     Pascal("camel_snake_kebab") == "CamelSnakeKebab"
//     Snake("camelSNAKEKebab") == "camel_snake_kebab"
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
//     CamelSnake("__camel_snake_kebab__") == "__Camel_Snake_Kebab__"
//     Kebab("__camel_snake_kebab") == "camel-snake-kebab"
//     Screaming("__camel_snake_kebab") == "CAMEL SNAKE KEBAB"
//     CamelKebab("--camel-snake-kebab") == "--Camel-Snake-Kebab"
//     Snake("--camel-snake-kebab") == "camel_snake_kebab"
//     Screaming("--camel-snake-kebab") == "CAMEL SNAKE KEBAB"
package casbab // import "resenje.org/casbab"

import "strings"

// Camel case is the practice of writing compound words
// or phrases such that each word or abbreviation in the
// middle of the phrase begins with a capital letter,
// with no spaces or hyphens.
//
// Example: "camelSnakeKebab".
func Camel(s string) string {
	return strings.Join(camel(words(s), 1), "")
}

// Pascal case is a variant of Camel case writing where
// the first letter of the first word is always capitalized.
//
// Example: "CamelSnakeKebab".
func Pascal(s string) string {
	return strings.Join(camel(words(s), 0), "")
}

// Snake case is the practice of writing compound words
// or phrases in which the elements are separated with
// one underscore character (_) and no spaces, with all
// element letters lowercased within the compound.
//
// Example: "camel_snake_kebab".
func Snake(s string) string {
	head, tail := headTailCount(s, '_')
	return strings.Repeat("_", head) + strings.Join(words(s), "_") + strings.Repeat("_", tail)
}

// CamelSnake case is a variant of Camel case with
// each element's first letter uppercased.
//
// Example: "Camel_Snake_Kebab".
func CamelSnake(s string) string {
	head, tail := headTailCount(s, '_')
	return strings.Repeat("_", head) + strings.Join(camel(words(s), 0), "_") + strings.Repeat("_", tail)
}

// ScreamingSnake case is a variant of Camel case with
// all letters uppercased.
//
// Example: "CAMEL_SNAKE_KEBAB".
func ScreamingSnake(s string) string {
	head, tail := headTailCount(s, '_')
	return strings.Repeat("_", head) + strings.Join(scream(words(s)), "_") + strings.Repeat("_", tail)
}

// Kebab case is the practice of writing compound words
// or phrases in which the elements are separated with
// one hyphen character (-) and no spaces, with all
// element letters lowercased within the compound.
//
// Example: "camel-snake-kebab".
func Kebab(s string) string {
	head, tail := headTailCount(s, '-')
	return strings.Repeat("-", head) + strings.Join(words(s), "-") + strings.Repeat("-", tail)
}

// CamelKebab case is a variant of Kebab case with
// each element's first letter uppercased.
//
// Example: "Camel-Snake-Kebab".
func CamelKebab(s string) string {
	head, tail := headTailCount(s, '-')
	return strings.Repeat("-", head) + strings.Join(camel(words(s), 0), "-") + strings.Repeat("-", tail)
}

// ScreamingKebab case is a variant of Kebab case with
// all letters uppercased.
//
// Example: "CAMEL-SNAKE-KEBAB".
func ScreamingKebab(s string) string {
	head, tail := headTailCount(s, '-')
	return strings.Repeat("-", head) + strings.Join(scream(words(s)), "-") + strings.Repeat("-", tail)
}

// Lower is returning detected words, not in a compound
// form, but separated by one space character with all
// letters in lower case.
//
// Example: "camel snake kebab".
func Lower(s string) string {
	return strings.Join(words(s), " ")
}

// Title is returning detected words, not in a compound
// form, but separated by one space character with first
// character in all letters in upper case and all other
// letters in lower case.
//
// Example: "Camel Snake Kebab".
func Title(s string) string {
	return strings.Join(camel(words(s), 0), " ")
}

// Screaming is returning detected words, not in a compound
// form, but separated by one space character with all
// letters in upper case.
//
// Example: "CAMEL SNAKE KEBAB".
func Screaming(s string) string {
	return strings.Join(scream(words(s)), " ")
}

func words(s string) (w []string) {
	start := 0
	l := len(s)
	var prevLower, prevUpper bool
Loop:
	for i, c := range s {
		switch c {
		case '-', '_', ' ':
			if start != i {
				w = append(w, strings.ToLower(s[start:i]))
			}
			start = i + 1
			prevLower = false
			prevUpper = false
			continue Loop
		}
		cs := s[i : i+1]
		if strings.ToUpper(cs) == cs {
			prevUpper = true
			if prevLower {
				if i != start {
					w = append(w, strings.ToLower(s[start:i]))
				}
				start = i
				prevLower = false
			}
		} else {
			prevLower = true
			if prevUpper {
				if i-1 != start {
					w = append(w, strings.ToLower(s[start:i-1]))
				}
				start = i - 1
				prevUpper = false
			}
		}
		if i == l-1 {
			w = append(w, strings.ToLower(s[start:]))
		}
	}
	return
}

func scream(s []string) []string {
	for i := 0; i < len(s); i++ {
		s[i] = strings.ToUpper(s[i])
	}
	return s
}

func camel(s []string, start int) []string {
	for i := start; i < len(s); i++ {
		switch len(s[i]) {
		case 0:
		case 1:
			s[i] = strings.ToUpper(s[i][0:1])
		default:
			s[i] = strings.ToUpper(s[i][0:1]) + s[i][1:]
		}
	}
	return s
}

func headTailCount(s string, sub byte) (head, tail int) {
	for i := 0; i < len(s); i++ {
		if s[i] != sub {
			head = i
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != sub {
			tail = len(s) - i - 1
			break
		}
	}
	return
}
