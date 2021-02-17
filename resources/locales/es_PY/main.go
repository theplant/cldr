package es_PY

import "github.com/ContextLogic/cldr"

var Locale = &cldr.Locale{
	Locale: "es_PY",
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
