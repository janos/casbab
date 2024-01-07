// Copyright (c) 2016, Janoš Guljaš <janos@resenje.org>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package casbab

import "testing"

var (
	cases = []struct {
		In  []string
		Out map[string]string
	}{
		{
			In: []string{
				"camelSnakeKebab",
				"CamelSnakeKebab",
				"camel_snake_kebab",
				"Camel_Snake_Kebab",
				"CAMEL_SNAKE_KEBAB",
				"camel-snake-kebab",
				"Camel-Snake-Kebab",
				"CAMEL-SNAKE-KEBAB",
				"camel snake kebab",
				"Camel Snake Kebab",
				"CAMEL SNAKE KEBAB",
				"camel__snake_kebab",
				"camel___snake_kebab",
				"camel____snake_kebab",
				"camel_ snake_kebab",
				"camel_  snake_kebab",
				"camel_   snake_kebab",
				"camel_-snake_kebab",
				"camel_ -snake_kebab",
				"camel - snake_kebab",
				" camel - snake_kebab",
				"CAMELSnakeKebab",
				"camelSNAKEKebab   ",
				"   camelSnakeKEBAB",
			},
			Out: map[string]string{
				"Camel":          "camelSnakeKebab",
				"Pascal":         "CamelSnakeKebab",
				"Snake":          "camel_snake_kebab",
				"CamelSnake":     "Camel_Snake_Kebab",
				"ScreamingSnake": "CAMEL_SNAKE_KEBAB",
				"Kebab":          "camel-snake-kebab",
				"CamelKebab":     "Camel-Snake-Kebab",
				"ScreamingKebab": "CAMEL-SNAKE-KEBAB",
				"Lower":          "camel snake kebab",
				"Title":          "Camel Snake Kebab",
				"Screaming":      "CAMEL SNAKE KEBAB",
			},
		},
		{
			In: []string{
				"__camel_snake_kebab__",
				"__camel_snakeKEBAB__",
				"__ Camel-snakeKEBAB__",
			},
			Out: map[string]string{
				"Camel":          "camelSnakeKebab",
				"Pascal":         "CamelSnakeKebab",
				"Snake":          "__camel_snake_kebab__",
				"CamelSnake":     "__Camel_Snake_Kebab__",
				"ScreamingSnake": "__CAMEL_SNAKE_KEBAB__",
				"Kebab":          "camel-snake-kebab",
				"CamelKebab":     "Camel-Snake-Kebab",
				"ScreamingKebab": "CAMEL-SNAKE-KEBAB",
				"Lower":          "camel snake kebab",
				"Title":          "Camel Snake Kebab",
				"Screaming":      "CAMEL SNAKE KEBAB",
			},
		},
		{
			In: []string{
				"__ camel-snake_kebab__ _",
				"__ camelSnake_Kebab_",
				"__CamelSnake_Kebab_",
				"__CamelSNAKE_Kebab_",
			},
			Out: map[string]string{
				"Camel":          "camelSnakeKebab",
				"Pascal":         "CamelSnakeKebab",
				"Snake":          "__camel_snake_kebab_",
				"CamelSnake":     "__Camel_Snake_Kebab_",
				"ScreamingSnake": "__CAMEL_SNAKE_KEBAB_",
				"Kebab":          "camel-snake-kebab",
				"CamelKebab":     "Camel-Snake-Kebab",
				"ScreamingKebab": "CAMEL-SNAKE-KEBAB",
				"Lower":          "camel snake kebab",
				"Title":          "Camel Snake Kebab",
				"Screaming":      "CAMEL SNAKE KEBAB",
			},
		},
		{
			In: []string{
				"--camel-snake-kebab",
				"--CAMELSnake_kebab",
			},
			Out: map[string]string{
				"Camel":          "camelSnakeKebab",
				"Pascal":         "CamelSnakeKebab",
				"Snake":          "camel_snake_kebab",
				"CamelSnake":     "Camel_Snake_Kebab",
				"ScreamingSnake": "CAMEL_SNAKE_KEBAB",
				"Kebab":          "--camel-snake-kebab",
				"CamelKebab":     "--Camel-Snake-Kebab",
				"ScreamingKebab": "--CAMEL-SNAKE-KEBAB",
				"Lower":          "camel snake kebab",
				"Title":          "Camel Snake Kebab",
				"Screaming":      "CAMEL SNAKE KEBAB",
			},
		},
		{
			In: []string{
				"-camel-snake-kebab----",
				"-CAMEL   Snake_kebab ----",
			},
			Out: map[string]string{
				"Camel":          "camelSnakeKebab",
				"Pascal":         "CamelSnakeKebab",
				"Snake":          "camel_snake_kebab",
				"CamelSnake":     "Camel_Snake_Kebab",
				"ScreamingSnake": "CAMEL_SNAKE_KEBAB",
				"Kebab":          "-camel-snake-kebab----",
				"CamelKebab":     "-Camel-Snake-Kebab----",
				"ScreamingKebab": "-CAMEL-SNAKE-KEBAB----",
				"Lower":          "camel snake kebab",
				"Title":          "Camel Snake Kebab",
				"Screaming":      "CAMEL SNAKE KEBAB",
			},
		},
		{
			In: []string{
				"xCamelXXSnakeXXXKebab",
				"XCamelXXSnakeXXXKebab",
				"x_camel_xx_snake_xxx_kebab",
				"X_Camel_XX_Snake_XXX_Kebab",
				"X_CAMEL_XX_SNAKE_XXX_KEBAB",
				"x-camel-xx-snake-xxx-kebab",
				"X-Camel-XX_Snake-XXX-Kebab",
				"X-CAMEL-XX_SNAKE-XXX-KEBAB",
				"x camel xx snake xxx kebab",
				"X Camel XX Snake XXX Kebab",
				"X CAMEL XX SNAKE XXX KEBAB",
			},
			Out: map[string]string{
				"Camel":          "xCamelXxSnakeXxxKebab",
				"Pascal":         "XCamelXxSnakeXxxKebab",
				"Snake":          "x_camel_xx_snake_xxx_kebab",
				"CamelSnake":     "X_Camel_Xx_Snake_Xxx_Kebab",
				"ScreamingSnake": "X_CAMEL_XX_SNAKE_XXX_KEBAB",
				"Kebab":          "x-camel-xx-snake-xxx-kebab",
				"CamelKebab":     "X-Camel-Xx-Snake-Xxx-Kebab",
				"ScreamingKebab": "X-CAMEL-XX-SNAKE-XXX-KEBAB",
				"Lower":          "x camel xx snake xxx kebab",
				"Title":          "X Camel Xx Snake Xxx Kebab",
				"Screaming":      "X CAMEL XX SNAKE XXX KEBAB",
			},
		},
		{
			In: []string{
				"",
			},
			Out: map[string]string{
				"Camel":          "",
				"Pascal":         "",
				"Snake":          "",
				"CamelSnake":     "",
				"ScreamingSnake": "",
				"Kebab":          "",
				"CamelKebab":     "",
				"ScreamingKebab": "",
				"Lower":          "",
				"Title":          "",
				"Screaming":      "",
			},
		},
		{
			In: []string{
				"Ово је Brave NewСвет",
				" Ово је Brave NewСвет",
				" Ово је Brave NewСвет    ",
			},
			Out: map[string]string{
				"Camel":          "овоЈеBraveNewСвет",
				"Pascal":         "ОвоЈеBraveNewСвет",
				"Snake":          "ово_је_brave_new_свет",
				"CamelSnake":     "Ово_Је_Brave_New_Свет",
				"ScreamingSnake": "ОВО_ЈЕ_BRAVE_NEW_СВЕТ",
				"Kebab":          "ово-је-brave-new-свет",
				"CamelKebab":     "Ово-Је-Brave-New-Свет",
				"ScreamingKebab": "ОВО-ЈЕ-BRAVE-NEW-СВЕТ",
				"Lower":          "ово је brave new свет",
				"Title":          "Ово Је Brave New Свет",
				"Screaming":      "ОВО ЈЕ BRAVE NEW СВЕТ",
			},
		},
	}
	converters = map[string]func(string) string{
		"Camel":          Camel,
		"Pascal":         Pascal,
		"Snake":          Snake,
		"CamelSnake":     CamelSnake,
		"ScreamingSnake": ScreamingSnake,
		"Kebab":          Kebab,
		"CamelKebab":     CamelKebab,
		"ScreamingKebab": ScreamingKebab,
		"Lower":          Lower,
		"Title":          Title,
		"Screaming":      Screaming,
	}
)

func Test(t *testing.T) {
	for _, c := range cases {
		for _, in := range c.In {
			for converter, expected := range c.Out {
				got := converters[converter](in)
				if got != expected {
					t.Errorf("Converting %q to %s expected %q, but got %q", in, converter, expected, got)
				}
			}
		}
	}
}

var benchmarkPhrase = "xCAMELSnakeKebab_screaming pascal XXX"

func BenchmarkCamel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Camel(benchmarkPhrase)
	}
}

func BenchmarkPascal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pascal(benchmarkPhrase)
	}
}

func BenchmarkSnake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Snake(benchmarkPhrase)
	}
}
func BenchmarkCamelSnake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CamelSnake(benchmarkPhrase)
	}
}

func BenchmarkScreamingSnake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ScreamingSnake(benchmarkPhrase)
	}
}

func BenchmarkKebab(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Kebab(benchmarkPhrase)
	}
}

func BenchmarkCamelKebab(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CamelKebab(benchmarkPhrase)
	}
}

func BenchmarkScreamingKebab(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ScreamingKebab(benchmarkPhrase)
	}
}

func BenchmarkLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Lower(benchmarkPhrase)
	}
}

func BenchmarkTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Title(benchmarkPhrase)
	}
}

func BenchmarkScreaming(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Screaming(benchmarkPhrase)
	}
}
