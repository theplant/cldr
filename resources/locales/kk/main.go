package kk

import "github.com/ContextLogic/cldr"

var Locale = &cldr.Locale{
	Locale: "kk",
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