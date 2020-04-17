package languages

// LanguageOverrides structure.
type LanguageOverrides struct {
	Overrides map[string]map[rune]string
}

// NewLanguageOverrides creates new LanguageOverrides object.
func NewLanguageOverrides() *LanguageOverrides {
	return &LanguageOverrides{
		Overrides: initOverrides(),
	}
}

// AddLanguageOverrides adds custom transliterations overrides.
func (lo *LanguageOverrides) AddLanguageOverrides(overrides *map[string]map[rune]string) {
	for langcode, override := range *overrides {
		lo.AddLanguageOverride(langcode, override)
	}
}

// AddLanguageOverride adds custom transliteration overrides for specific language.
func (lo *LanguageOverrides) AddLanguageOverride(langcode string, override map[rune]string) {
	lo.Overrides[langcode] = override
}

// initOverrides adds some general language specific overrides.
func initOverrides() map[string]map[rune]string {
	// Language codes in ISO 639-1 format.
	return map[string]map[rune]string{
		"de": DE,
		"da": DA,
		"eo": EO,
		"ru": RU,
		"bg": BG,
		"sv": SV,
		"hu": HU,
		"hr": HR,
		"sl": SL,
		"sr": SR,
		"nb": NB,
		"uk": UK,
		"mk": MK,
		"ca": CA,
		"bs": BS,
	}
}
