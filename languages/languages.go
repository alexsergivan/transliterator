package languages

// LanguageOverwrites structure.
type LanguageOverwrites struct {
	Overwrites map[string]map[rune]string
}

// NewLanguageOverwrites creates new LanguageOverwrites object.
func NewLanguageOverwrites() *LanguageOverwrites {
	return &LanguageOverwrites{
		Overwrites: initOverwrites(),
	}
}

// AddLanguageOverwrites adds custom transliterations overwrites.
func (lo *LanguageOverwrites) AddLanguageOverwrites(overwrites *map[string]map[rune]string) {
	for langcode, override := range *overwrites {
		lo.AddLanguageOverride(langcode, override)
	}
}

// AddLanguageOverride adds custom transliteration overwrites for specific language.
func (lo *LanguageOverwrites) AddLanguageOverride(langcode string, override map[rune]string) {
	lo.Overwrites[langcode] = override
}

// initOverwrites adds some general language specific overwrites.
func initOverwrites() map[string]map[rune]string {
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
