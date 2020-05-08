Golang text Transliterator
==============

[![Build Status](https://travis-ci.com/alexsergivan/transliterator.svg?branch=master)](https://travis-ci.com/github/alexsergivan/transliterator)
[![Coverage Status](https://coveralls.io/repos/github/alexsergivan/transliterator/badge.svg)](https://coveralls.io/github/alexsergivan/transliterator)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexsergivan/transliterator)](https://goreportcard.com/report/github.com/alexsergivan/transliterator)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/9b062cd8ba9f4f7f850e167d6966b75b)](https://www.codacy.com/manual/alexsergivan/transliterator?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=alexsergivan/transliterator&amp;utm_campaign=Badge_Grade)


Golang Transliterator provides one-way string transliteration. It takes Unicode text and converts to ASCII characters.
Example use-case: transliterate cyrilic city name to be able to use it in the url ("Київ" ==> "Куiv").

For now, only these languages have specific transliteration rules: DE, DA, EO, RU, BG, SV, HU, HR, SL, SR, NB, UK, MK, CA, BS. For other languages, general ASCII transliteration rules will be applied. Also, this package supports adding custom transliteration rules for your specific use-case. Please check the examples section below.


Installation
------------

```
go get -u github.com/alexsergivan/transliterator
```


Language specific transliteration example
------

```go
package main

import (
	"fmt"
	"github.com/alexsergivan/transliterator"
)

func main() {
	trans := transliterator.NewTransliterator(nil)
	text := "München"
	// Langcode should be provided accrding to ISO 639-1.
	fmt.Println(trans.Transliterate(text, "de")) // Result: Muenchen
	fmt.Println(trans.Transliterate(text, "en")) // Result: Munchen

	anotherText := "你好"
	fmt.Println(trans.Transliterate(anotherText, "")) // Result: Ni Hao

	oneMoreText := "Київ"
	fmt.Println(trans.Transliterate(oneMoreText, "uk")) // Result: Kyiv
	fmt.Println(trans.Transliterate(oneMoreText, "en")) // Result: Kiyiv
	fmt.Println(trans.Transliterate(oneMoreText, "")) // Result: Kiyiv
}
```

Adding of custom Language translitartion rules
------

```go
package main

import (
	"fmt"
	"github.com/alexsergivan/transliterator"
)

func main() {
	customLanguageOverrites := make(map[string]map[rune]string)

	customLanguageOverrites["myLangcode"] = map[rune]string{
		// Ї
		0x407: "CU",
		// и
		0x438: "y",
	}
	trans := transliterator.NewTransliterator(&customLanguageOverrites)
	text := "КиЇв"
	fmt.Println(trans.Transliterate(text, "myLangcode")) // Result: KyCUv

}
```
