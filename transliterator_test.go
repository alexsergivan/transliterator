package transliterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItShouldTransliterateGermanCorrectly(t *testing.T) {
	text := "München"
	expected := "Muenchen"

	transliterator := NewTransliterator(nil)
	actual := transliterator.Transliterate(text, "de")

	assert.Equal(t, expected, actual)

}

func TestItShouldTransliterateUkrainianCorrectly(t *testing.T) {
	text := "Київ"
	expected := "Kyiv"
	transliterator := NewTransliterator(nil)
	actual := transliterator.Transliterate(text, "uk")

	assert.Equal(t, expected, actual)
}

func TestItShouldTransliterateCorrectlyWithCustomOverrides(t *testing.T) {
	text := "КиЇв"
	expected := "KyCUv"

	customLanguageOverrites := make(map[string]map[rune]string)

	customLanguageOverrites["custom"] = map[rune]string{
		0x407: "CU",
		0x438: "y",
	}
	transliterator := NewTransliterator(&customLanguageOverrites)
	actual := transliterator.Transliterate(text, "custom")

	assert.Equal(t, expected, actual)
}

func TestItShouldTransliterateGeneral(t *testing.T) {
	cases := map[string]string{
		"北京":           "Bei Jing ",
		"80 km/h":      "80 km/h",
		"дом":          "dom",
		"ⓐⒶ⑳⒇⒛⓴⓾⓿":     "aA20(20)20.20100",
		"ch\u00e2teau": "chateau",
		"\u1eff":       "",
	}

	transliterator := NewTransliterator(nil)
	for text, expected := range cases {
		actual := transliterator.Transliterate(text, "")
		assert.Equal(t, expected, actual)
	}
}
