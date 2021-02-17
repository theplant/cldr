package ksb

import "github.com/ContextLogic/cldr"

var Locale = &cldr.Locale{
	Locale: "ksb",
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
