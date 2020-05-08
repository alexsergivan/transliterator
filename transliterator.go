package transliterator

import (
	"strings"
	"unicode"

	"github.com/alexsergivan/transliterator/data"
	"github.com/alexsergivan/transliterator/languages"
)

// Transliterator structure.
type Transliterator struct {
	LanguageOverrides *languages.LanguageOverrides
	Data              map[rune][]string
}

// NewTransliterator creates Transliterator object.
func NewTransliterator(customLanguageOverrides *map[string]map[rune]string) *Transliterator {
	languageOverrides := languages.NewLanguageOverrides()
	if customLanguageOverrides != nil {
		languageOverrides.AddLanguageOverrides(customLanguageOverrides)
	}

	return &Transliterator{
		LanguageOverrides: languageOverrides,
		Data:              data.NewTransliterationData().Data,
	}
}

// Transliterate performs transliteration of the input text. If the langcode (ISO 639-1) is specified, it will use
// specific language transliteration rules.
func (t *Transliterator) Transliterate(text, langcode string) string {
	var replacement strings.Builder
	runes := []rune(text)
	for _, rune := range runes {
		if overrides, ok := t.LanguageOverrides.Overrides[langcode]; ok {
			if val, ok := overrides[rune]; ok {
				replacement.WriteString(val)
				continue
			}
		}
		// If the rune number less then maximum ASCII value, use it directly.
		if rune < unicode.MaxASCII {
			replacement.WriteString(string(rune))
			continue
		}

		// Example: "Ї" => in the hexadecimal - 0x407
		// bank: 0x4
		// code: 0x7

		// Shifting rune to the right by 8 bits.
		bank := rune >> 8
		// masks the variable so it leaves only the value in the last 8 bits, and ignores all the rest of the bits
		code := rune & 0xff
		if transliterationDataVal, ok := t.Data[bank]; ok {
			if len(transliterationDataVal) > int(code) {
				replacement.WriteString(transliterationDataVal[code])
			}
		}
	}

	return replacement.String()
}
