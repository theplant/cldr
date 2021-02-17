package en_AI

import "github.com/ContextLogic/cldr"

var Locale = &cldr.Locale{
	Locale: "en_AI",
	Number: cldr.Number{
		Symbols:    symbols,
		Formats:    formats,
		Currencies: currencies,
	},
	Calendar:   calendar,
	PluralRule: pluralRule,
}

func init() {
	cldr.RegisterLocale(Locale)
}
